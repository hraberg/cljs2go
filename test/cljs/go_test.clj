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
    (doseq [[actual assertion expected setup message] assertions]
      (printf "{\n")
      (when setup
        (printf "/*\n")
        (printf "%s\n" (with-out-str (clojure.pprint/pprint setup)))
        (printf " */\n")
        (printf "\t%s\n" (-> [setup] cljs->ast ast->go)))
      (printf "\tassert.%s(t%s, %s%s)\n"
              (s/capitalize (name assertion))
              (if (nil? expected) "" (str ", " expected))
              actual
              (if (nil? message) "" (str ", `" message "`")))
      (printf "}\n"))
    (printf "}\n")))

(defn emit-test [package tests]
  (let [^File f (io/file "target/generated" (str package ".go"))
        src (gofmt (apply str (testify-header package) tests))]
    (io/make-parents f)
    (spit f src)
    (go-test (io/file "." (.getParent f))))  )

(defn constant [expected actual]
  ["X" "Equal" expected (list 'def 'x actual)])

(defn expr [expected actual op]
  ["X" "Equal" expected (list 'def 'x actual) op])

(deftest go-emitter-tests
  (->>
   [(testify "Constants"
              (constant "nil" nil)
              (constant  true true)
              (constant false false)
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
              (constant "&CljsCoreSymbol{Ns: `user`, Name: `x`, Str: `user/x`, Hash: float64(-568535109), Meta: nil}"
                        ''user/x)
              (constant "&CljsCoreKeyword{Ns: `user`, Name: `x`, Fqn: `user/x`, Hash: float64(2085900660)}"
                        :user/x))
    (testify "SpecialForms"
             (expr 1 '(let [y 1]
                        y) "let")
             (expr true '(let [y :foo]
                           (if y true false)) "if")
             (expr 5 '(loop [y 0]
                        (if (>= y 5)
                          y
                          (recur (inc y)))) "loop")
             (expr 3 '(do 1 2 3) "do")
             (expr "&js.Date{Millis: 0}" '(js/Date. 0) "new")
             (expr 1970 '(.getUTCFullYear (js/Date. 0)) "dot")
             (expr "math.Inf(1)" 'js/Infinity "var")
             (expr true '(let [x 2]
                           (case x
                             2 true
                             1 false
                             0))
                   "case*")
             (expr "reflect.Float64" '(let [x 1]
                                        (js* "reflect.ValueOf(x).Kind()"))
                   "js*")
             (expr "reflect.Int" '(let [v (js* "reflect.ValueOf(1)")]
                                    (-> v .Type .Kind))
                   "js*")
             (expr "&js.Error{`Foo`}" '(try
                                         (throw (js/Error. "Foo"))
                                         (catch js/Error e
                                           e))
                   "try")
             (expr "`TypeError`" '(try
                                    (throw (js/TypeError. "Foo"))
                                    (catch js/Error _
                                      "Error")
                                    (catch js/TypeError _
                                      "TypeError"))
                   "try")
             (expr "`Bar`" '(try
                              "Bar"
                              (catch js/Error e
                                e)
                              (finally
                                "Baz"))
                   "try")
             (expr "map[string]interface{}{`finally`: true}"
                   '(let [x (js* "map[string]interface{}{}")]
                      (try
                        x
                        (finally
                          (js* "x[`finally`] = true"))))
                   "try")
             (expr "map[string]interface{}{`catch`: true, `finally`: true, `last`: `finally`}"
                   '(let [x (js* "map[string]interface{}{}")]
                      (try
                        (throw (js/Error. "Foo"))
                        (catch js/Error _
                          (js* "x[`catch`] = true")
                          (js* "x[`last`] = `catch`")
                          x)
                        (finally
                          (js* "x[`finally`] = true")
                          (js* "x[`last`] = `finally`"))))
                   "try"))]
   (emit-test "go_test")))

(deftest go-main-test
  (go-test "."))
