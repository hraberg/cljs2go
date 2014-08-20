(ns cljs.go-test
  (:require [clojure.test :refer :all]
            [cljs.go :refer :all]
            [clojure.string :as s]
            [clojure.java.io :as io]
            [clojure.java.shell :as sh]
            [clojure.pprint :as pp])
  (:import [java.io File]))

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

(defn go-test [package]
  (let [{:keys [out err exit]} (sh/sh "go" "test" package)
        out (s/replace (s/replace (str err out) "\r" "\n") "\n\t\t" "")]
    (is (zero? exit) out)))

(defn testify-header []
  (with-out-str
    (println "package go_test")
    (println "import (")
    (doseq [import ["math" "testing"
                    "github.com/hraberg/cljs.go/js"
                    "github.com/hraberg/cljs.go/js/Math"
                    "github.com/stretchr/testify/assert"]]
      (println "\t" (pr-str import)))
    (println "\t" "." (pr-str "github.com/hraberg/cljs.go/cljs/core"))
    (println ")")))

(defn testify [test & assertions]
  (with-out-str
    (printf "func Test_%s(t *testing.T) {\n" (name test))
    (doseq [[cljs assertion expected] assertions]
      (printf "\tassert.%s(t%s, %s)\n"
              assertion
              (if expected (str ", " expected) "")
              (if (string? cljs)
                cljs
                (-> cljs cljs->ast ast->go))))
    (printf "}\n")))

(defn emit-and-test [package src]
  (let [^File f (io/file "target/generated/" (str package ".go"))]
    (io/make-parents f)
    (spit f (gofmt (apply str (testify-header) src)))
    (go-test (str (io/file "." (.getParent f)))))  )

(deftest go-emitter-tests
  (->> [(testify "IncludeCljsCore"
                 ["Truth_(true)" "True"]
                 ["Truth_(false)" "False"])]
       (emit-and-test "go_test")))

(deftest go-main-test
  (go-test "."))
