(ns cljs.go
  (:require [cljs.analyzer]
            [cljs.closure]
            [cljs.env]
            [cljs.js-deps]
            [cljs.go.compiler]
            [clojure.string :as s]
            [clojure.walk :as w]
            [clojure.java.shell :as sh]))

(defn empty-env []
  (dissoc (cljs.analyzer/empty-env) :js-globals))

(defn cljs->ast [in]
  (binding [cljs.analyzer/*cljs-ns* 'cljs.user]
    (doall (map #(cljs.analyzer/analyze (empty-env) %) in))))

(defn ast->go [in]
  (with-out-str
    (cljs.env/with-compiler-env (cljs.env/default-compiler-env)
      (dorun (map cljs.compiler/emit in)))))

(defn ast->simple-ast [in]
  (->> in
       (w/prewalk
        #(cond-> %
                 (map? %) (dissoc :children :ns :column :shadow
                                  :protocol-inline :protocol-impl
                                  :doc :js-globals :jsdoc)
                 (:locals %) (update-in [:locals] keys)))))

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

(defn print-simple-ast-and-emitted-go [in]
  (let [ast (cljs->ast in)]
    (println "==================== AST")
    (->> ast
         ast->simple-ast
         clojure.pprint/pprint)
    (println "==================== Go")
    (->> ast
         ast->go
         gofmt
         println)))

(comment
  (cljs.closure/build '[(ns hello.core)
                        (defn ^{:export greet} greet [n] (str "Hola " n))
                        (defn ^:export sum [xs] 42)]
                      {:optimizations :none})

  (cljs.closure/build (str "checkouts/clojurescript/" "samples/hello/src")
                      {:optimizations :none}))
