(ns cljs.go
  (:require [cljs.analyzer :as ana]
            [cljs.env :as env]
            [cljs.go.compiler]
            [clojure.string :as s]
            [clojure.java.io :as io]
            [clojure.java.shell :as sh]))

(set! *warn-on-reflection* true)

(defn elide-children [_ ast]
  (dissoc ast :children))

(defn simplify-env [_ {:keys [op] :as ast}]
  (cond-> ast
          (= op :fn) (update-in [:methods] #(map (partial simplify-env nil) %))
          (= op :set!) (update-in [:target] (partial simplify-env nil))
          :then (update-in [:env] #(select-keys % [:context :column :line]))))

(defn cljs->ast
  ([in] (cljs->ast in (ana/empty-env)))
  ([in env]
     (env/ensure
      (binding [ana/*passes* [elide-children simplify-env ana/infer-type]
                ana/*cljs-ns* (or ana/*cljs-ns* 'cljs.user)]
        (cljs.compiler/setup-native-defs env)
        (doall (map #(ana/analyze (cljs.compiler/ensure-ns-exist env) %) in))))))

(defn ast->go [in]
  (try
    (env/ensure
     (with-out-str
       (binding [ana/*cljs-static-fns* true
                 cljs.compiler/*go-defs* (or cljs.compiler/*go-defs* (atom []))]
         (dorun (map cljs.compiler/emit in)))))
    (catch Throwable t
      (.printStackTrace t)
      (throw t))))

(defn godef
  ([var] (godef nil var))
  ([package var]
     (let [in (str "package _;"
                   (when (seq (str package))
                     (str "import \"" package "\";"))
                   "var _ = " (cond->> var package (str (last (s/split package #"\.")) ".")))
           {:keys [exit err out]} (sh/sh "godef" "-i" "-t" "-o" (str (count in)) :in in)]
       (if (zero? exit)
         (let [[loc & def] (s/split-lines out)
               [file line col] (s/split loc #":")]
           (merge
            {:file file :def (s/trim (s/join "\n" (remove (comp empty? s/trim) def)))}
            (when (and line col)
              {:line (biginteger line) :col (biginteger col)}) ))
         (println err)))))

(defn go-signature [package var]
  (when-let [def (:def (godef package var))]
    (let [ret-tag (symbol (last (s/split def #"\s+")))
          [type method] (s/split var #"\.")
          params (s/split (last (re-find #"\((.+)\)" def)) #",")
          params (map #(s/split (s/trim %) #"\s+") params)
          params (loop [last-t nil [[v t] & params] (reverse params) acc []]
                   (if v
                     (let [t (symbol (or t last-t))]
                       (recur t params (conj acc (with-meta (symbol v) {:tag t}))))
                     (vec (reverse acc))))]
      (cond->
       {:ns package :ret-tag ret-tag :method-params (list params)}
       method (assoc :type (symbol type) :method (symbol method))
       (not method) (assoc :func (symbol var))))))

(defn godoc [package name]
  (let [{:keys [exit err out]} (sh/sh "godoc" package name)
        out (s/trim out)]
    (if (and (zero? exit) (not= "No match found." out))
      (->> (s/split out #"\n\n+")
           (group-by #(let [[what name s-or-i method ] (s/split % #" ")
                            name (-> (if (= \( (first name)) method name)
                                     (s/split #"\(")
                                     first
                                     symbol)]
                        (or
                         (when (= "type" what)
                           (some->> s-or-i keyword #{:struct :interface} (vector name)))
                         [name (keyword what)]))))
      (println err))))

(defn go-get [package]
  (let [{:keys [exit err]} (sh/sh "go" "get" package)]
    (assert (zero? exit) err)))

(defn get-goimports []
  (go-get "code.google.com/p/go.tools/cmd/goimports"))

(defn goimports [in]
  (get-goimports)
  (let [{:keys [exit out err]} (sh/sh "goimports" :in in)]
    (if (zero? exit)
      out
      (do (println err) in))))

(defn goimports-file [file]
  (get-goimports)
  (let [{:keys [exit out err]} (sh/sh "goimports" "-e=true" "-w=true" (str file))]
    (when-not (zero? exit)
      (println err)
      {:file file :errors (count (s/split-lines err))})))

(defn error-summary [dir {:keys [exit out err]}]
  (when-not (zero? exit)
    (println err)
    {:dir dir :errors (count (filter #(re-find #"^.+.go:\d+: " %) (s/split-lines err)))}))

(defn go-test-compile [dir]
  (->> (sh/sh "go" "test" "-c" "--gcflags" "-e" :dir dir) (error-summary dir)))

(defn go-install [dir]
  (->> (sh/sh "go" "install" "--gcflags" "-e" :dir dir) (error-summary dir)))

(defmacro with-fresh-ids [& body]
  `(let [^java.util.concurrent.atomic.AtomicInteger id# (.get (doto (.getDeclaredField clojure.lang.RT "id")
                                                                (.setAccessible true))
                                                              nil)
         current-id# (.get id#)]
     (try
       (.set id# 1)
       ~@body
       (finally
         (.set id# current-id#) ))))

(defn defs-in-file [f]
  (let [clj (read-string (str \[ (slurp f) \]))
        [[_ ns]] (filter (comp #{'ns} first) clj)]
    (for [[def? name] clj :when ('#{def defn} def?)]
      (symbol (str ns) (str name)))))

(defn maybe-add-overrides [dir]
  (when-let [overrides (io/resource (str (s/replace (str ana/*cljs-ns*) "." "/") "/overrides.cljs"))]
    (let [target (io/file dir "overrides.go")]
      (binding [cljs.compiler/*go-defs* (atom [])]
        (cljs.compiler/compile-file (io/file overrides) target {:overrides? true})
        (goimports-file target)))))

(defn compile-file
  ([] (compile-file "." (io/resource "cljs/core.cljs")))
  ([target-dir src]
     (let [src (io/file src)
           target (cljs.compiler/to-target-file target-dir src)
           dir (.getParentFile target)]
       (env/ensure
        (binding [ana/*passes* [elide-children simplify-env ana/infer-type]
                  ana/*cljs-static-fns* true
                  cljs.compiler/*go-def-vars* false
                  cljs.compiler/*go-defs* (or cljs.compiler/*go-defs* (atom []))]
          (with-fresh-ids
            (when-let [compiled-ns (:ns (cljs.compiler/compile-file src target))]
              (binding [ana/*cljs-ns* compiled-ns
                        cljs.compiler/*go-skip-def* #{}]
                (maybe-add-overrides dir))
              (goimports-file target)
              (if (re-find #"-test$" (str compiled-ns))
                (go-test-compile dir)
                (go-install dir)))))))))

(defn compile-clojurescript
  ([] (compile-clojurescript "."))
  ([target-dir]
     (env/ensure
      (with-fresh-ids
        (reset! ana/-cljs-macros-loaded false)
        (ana/load-core))
      (doseq [ns '[cljs.core
                   ; cljs.reader ;; doesn't work yet, some ns / import issue.
                   ; clojure.core.reducers ;; fails to macroexpand a destructure during analyzing.
                   clojure.set clojure.data clojure.string clojure.walk clojure.zip]
              :let [resource (io/resource (str (s/replace (str ns) "." "/") ".cljs"))]]
        (time
         (do
           (println "compiling" (str resource))
           (compile-file target-dir resource)))))))

(defn -main [& args]
  (println "ClojureScript to Go [clojure]")
  (compile-clojurescript))
