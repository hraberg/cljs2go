(ns cljs.go
  (:require [cljs.analyzer :as ana]
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

(defn cljs->ast
  ([in] (cljs->ast in (ana/empty-env)))
  ([in env]
     (cljs.env/ensure
      (binding [ana/*cljs-ns* (or ana/*cljs-ns* 'cljs.user)
                ana/*passes* [elide-children simplify-env ana/infer-type]]
        (when-not (ana/get-namespace ana/*cljs-ns*)
          (ana/analyze env (list 'ns ana/*cljs-ns*)))
        (doall (map #(ana/analyze (assoc env :ns (ana/get-namespace ana/*cljs-ns*)) %) in))))))

(defn ast->go [in]
  (with-out-str
    (cljs.env/with-compiler-env (cljs.env/default-compiler-env)
      (dorun (map cljs.compiler/emit in)))))

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

(defn -main [& args]
  (println "ClojureScript to Go [clojure]"))
