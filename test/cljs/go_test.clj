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
    (doseq [[actual assertion expected setup] assertions]
      (printf "{\n")
      (when setup
        (printf "\t%s\n" (-> setup cljs->ast ast->go)))
      (printf "\tassert.%s(t%s, %s)\n"
              (s/capitalize (name assertion))
              (if expected (str ", " expected) "")
              actual)
      (printf "}\n"))
    (printf "}\n")))

(defn emit-test [package & tests]
  (let [^File f (io/file "target/generated" (str package ".go"))
        src (gofmt (apply str (testify-header package) tests))]
    (io/make-parents f)
    (spit f src)
    (go-test (io/file "." (.getParent f))))  )

(defn constant [expected actual]
  ["X" "Equal" expected [(list 'def 'x actual)]])

(deftest go-emitter-tests
  (->>
   (testify "Constants"
            (constant "nil" nil)
            (constant true true)
            (constant 1 1)
            (constant 3.14 3.14)
            (constant 2 '(inc 1))
            (constant 10 '(* 4 2.5))
            (constant "`foo`" "foo")
            (constant "'x'" \x)
            (constant "map[string]interface{}{`foo`: `bar`}",
                      (binding [*data-readers* cljs.tagged-literals/*cljs-data-readers*]
                        (read-string "#js {:foo \"bar\"}")))
            (constant "[]interface{}{\"foo\", \"bar\"}"
                      (binding [*data-readers* cljs.tagged-literals/*cljs-data-readers*]
                        (read-string "#js [\"foo\", \"bar\"])")))
            (constant "&js.Date{Millis: 1408642409602}" #inst "2014-08-21T17:33:29.602-00:00")
            (constant "&CljsCoreUUID{Uuid: `15c52219-a8fd-4771-87e2-42ee33b79bca`}"
                      #uuid "15c52219-a8fd-4771-87e2-42ee33b79bca")
            (constant "&js.RegExp{Pattern: `x`, Flags: ``}" #"x")
            (constant "&js.RegExp{Pattern: ``, Flags: ``}" #"")
            (constant "&CljsCoreSymbol{Ns: `user`, Name: `x`, Str: `user/x`, Hash: -568535109, Meta: nil}"
                      ''user/x)
            (constant "&CljsCoreKeyword{Ns: `user`, Name: `x`, Fqn: `user/x`, Hash: 2085900660}"
                      :user/x))
   (emit-test "go_test")))

(deftest go-main-test
  (go-test "."))
