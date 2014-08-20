(ns cljs.go-test
  (:require [clojure.test :refer :all]
            [cljs.go :refer :all]
            [clojure.string :as s]
            [clojure.java.io :as io]
            [clojure.java.shell :as sh])
  (:import [java.io File]))

(defmacro emit [& body]
  `(let [res# (emit-go* '~body)]
     (is (zero? (:exit res#)) (:err res#))))

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
  (let [{:keys [out err exit]} (sh/sh "go" "test" (str package))
        out (s/replace (s/replace (str err out) "\r" "\n") "\n\t\t" "")]
    (is (zero? exit) out)))

(defn testify-header [package & imports]
  (with-out-str
    (println "package" package)
    (println "import (")
    (doseq [import (concat ["testing"
                            "github.com/stretchr/testify/assert"
                            "github.com/hraberg/cljs.go/js"
                            "github.com/hraberg/cljs.go/js/Math"]
                           imports)]
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

(defn emit-test [package & tests]
  (let [^File f (io/file "target/generated" (str package ".go"))
        src (gofmt (apply str (testify-header package) tests))]
    (io/make-parents f)
    (spit f src)
    (go-test (io/file "." (.getParent f))))  )

(deftest go-emitter-tests
  (emit-test "go_test"
             (testify "IncludeCljsCore"
                      ["Truth_(true)" "True"]
                      ["Truth_(false)" "False"])))

(deftest go-main-test
  (go-test "."))
