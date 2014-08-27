(ns cljs.go
  (:require [cljs.analyzer :as ana]
            [cljs.env :as env]
            [cljs.go.compiler]
            [clojure.string :as s]
            [clojure.pprint :as pp]
            [clojure.java.io :as io]
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

(defn cljs->ast
  ([in] (cljs->ast in (ana/empty-env)))
  ([in env]
     (env/ensure
      (binding [ana/*cljs-ns* (or ana/*cljs-ns* 'cljs.user)
                ana/*passes* [elide-children simplify-env ana/infer-type]]
        (when-not (ana/get-namespace ana/*cljs-ns*)
          (ana/analyze env (list 'ns ana/*cljs-ns*)))
        (doall (map #(ana/analyze (assoc env :ns (ana/get-namespace ana/*cljs-ns*)) %) in))))))

(defn ast->go [in]
  (try
    (env/ensure
     (with-out-str
       (binding [ana/*cljs-static-fns* true]
         (dorun (map cljs.compiler/emit in)))))
    (catch Throwable t
      (.printStackTrace t)
      (throw t))))

(defn go-get [package]
  (let [{:keys [exit err]} (sh/sh "go" "get" package)]
    (assert (zero? exit) err)))

(defn goimports [in]
  (go-get "code.google.com/p/go.tools/cmd/goimports")
  (let [{:keys [exit out err]} (sh/sh "goimports" :in in)]
    (if (zero? exit)
      out
      (do (println err) in))))

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

(defn compile-file
  ([] (compile-file "target" (io/resource "cljs/core.cljs")))
  ([target-dir src]
     (let [src (io/file src)
           target (cljs.compiler/to-target-file target-dir src)]
       (env/ensure
        (binding [ana/*passes* [elide-children simplify-env ana/infer-type]]
          (with-fresh-ids
            (cljs.compiler/compile-file src target)))))))

(defn -main [& args]
  (println "ClojureScript to Go [clojure]"))
