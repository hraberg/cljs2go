(ns cljs.go-test
  (:require [clojure.test :refer :all]
            [cljs.go :refer :all]
            [clojure.string :as s]
            [clojure.java.shell :as sh]
            [clojure.pprint :as pp]))

(defn emit* [cljs]
  (let [ast (cljs->ast cljs)
        go (ast->go ast)
        {:keys [exit err]} (goimports go)]
    (when-not (zero? exit)
      (println)
      (println "==================== ClojureScript")
      (doseq [f cljs]
        (pp/pprint f))
      (println "==================== AST")
      (pp/pprint (ast->simple-ast ast))
      (println "==================== Go")
      (println (with-line-numbers go))
      (println "==================== gofmt")
      (println err))
    (is (zero? exit) err)))

(defmacro emit [& body]
  `(emit* '~body))

(deftest cljs->go
  (emit (ns test.go (:require [goog.array :as array]))
        (defn plus-one [x] (inc x)))

  (emit (ns hello)
        (defn -main []
          (println "Hello World")))

  (comment
    (emit (ns hello)
          (defn world []
            "World")
          (defn ^:export greet [n]
            (str "Hello " n (world))))

    (emit (ns hello)
          (defn foo
            ([] (foo "World"))
            ([x] (println "Hello " x))
            ([x & ys] (println "Hello " x ys))))))

(deftest go-test
  (let [{:keys [out err exit]} (sh/sh "go" "test")
        out (s/replace (s/replace (str err out) "\r" "\n") "\n\t\t" "")]
    (is (zero? exit) out)))
