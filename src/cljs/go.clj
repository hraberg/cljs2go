(ns cljs.go
  (:require [cljs.analyzer :as ana]
            [cljs.env :as env]
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

(defn cljs->ast
  ([in] (cljs->ast in (ana/empty-env)))
  ([in env]
     (env/ensure
      (binding [ana/*cljs-ns* (or ana/*cljs-ns* 'cljs.user)
                ana/*cljs-static-fns* true
                ana/*passes* [elide-children simplify-env ana/infer-type]]
        (when-not (ana/get-namespace ana/*cljs-ns*)
          (ana/analyze env (list 'ns ana/*cljs-ns*)))
        (doall (map #(ana/analyze (assoc env :ns (ana/get-namespace ana/*cljs-ns*)) %) in))))))

(defn ast->go [in]
  (env/ensure
   (with-out-str
     (dorun (map cljs.compiler/emit in)))))

(defn go-get [package]
  (let [{:keys [exit err]} (sh/sh "go" "get" package)]
    (assert (zero? exit) err)))

(defn goimports [in]
  (go-get "code.google.com/p/go.tools/cmd/goimports")
  (let [{:keys [exit out err]} (sh/sh "goimports" :in in)]
    (if (zero? exit)
      out
      (do (println err) in))))

(defn -main [& args]
  (println "ClojureScript to Go [clojure]"))
