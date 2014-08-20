(ns cljs.go
  (:require [cljs.analyzer]
            [cljs.closure]
            [cljs.env]
            [cljs.js-deps]
            [cljs.go.compiler]
            [clojure.string :as s]
            [clojure.pprint :as pp]
            [clojure.java.shell :as sh]))

(defn elide-children [_ ast]
  (dissoc ast :children))

(defn simplify-env [_ {:keys [op] :as ast}]
  (let [env (:env ast)
        ast (if (= op :fn)
              (assoc ast :methods
                (map #(simplify-env nil %) (:methods ast)))
              ast)]
    (update-in ast [:env] #(select-keys % [:context :column :line]))))

(defn cljs->ast [in]
  (binding [cljs.analyzer/*cljs-ns* 'cljs.user
            cljs.analyzer/*passes* [elide-children simplify-env cljs.analyzer/infer-type]]
    (doall (map #(cljs.analyzer/analyze (cljs.analyzer/empty-env) %) in))))

(defn ast->go [in]
  (with-out-str
    (cljs.env/with-compiler-env (cljs.env/default-compiler-env)
      (dorun (map cljs.compiler/emit in)))))

(defn -main [& args]
  (println "ClojureScript to Go [clojure]"))

(defn go-get [package]
  (let [{:keys [exit err]} (sh/sh "go" "get" package)]
    (assert (zero? exit) err)))

(defn goimports [in]
  (go-get "code.google.com/p/go.tools/cmd/goimports")
  (sh/sh "goimports" :in in))

(defn gofmt [in]
  (let [{:keys [exit out err]} (goimports in)]
    (if (zero? exit)
      out
      (do (println err) in))))

(defn with-line-numbers [in]
  (->> (for [[n l] (map-indexed vector (s/split-lines in))]
         (format "%3d  %s\n" (inc n) l))
       (apply str)
       s/trim-newline))

(defn print-ast-and-emitted-go
  ([cljs ast go] (print-ast-and-emitted-go cljs ast go nil))
  ([cljs ast go gofmt]
      (println "==================== ClojureScript")
      (doseq [f cljs]
        (pp/pprint f))
      (println "==================== AST")
      (pp/pprint ast)
      (println "==================== Go")
      (println (with-line-numbers go))
      (when (seq gofmt)
        (println "==================== gofmt")
        (println gofmt))))

(defn emit-go*
  ([cljs] (emit-go* false cljs))
  ([debug? cljs]
     (let [ast (cljs->ast cljs)
           go (ast->go ast)
           {:keys [exit err out] :as result} (goimports go)
           error? (pos? exit)
           go (if error? go out)]
       (when (or error? debug?)
         (print-ast-and-emitted-go cljs ast go err))
       result)))

(defn emit-go [cljs]
  (let [{:keys [out err exit]} (emit-go* cljs)]
    (assert (zero? exit) err)
    out))

(comment
  (cljs.closure/build '[(ns hello.core)
                        (defn ^{:export greet} greet [n] (str "Hola " n))
                        (defn ^:export sum [xs] 42)]
                      {:optimizations :none})

  (cljs.closure/build (str "checkouts/clojurescript/" "samples/hello/src")
                      {:optimizations :none}))
