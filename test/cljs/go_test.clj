(ns cljs.go-test
  (:require [clojure.test :refer :all]
            [cljs.go :refer :all]
            [cljs.tagged-literals :refer [*cljs-data-readers*]]
            [clojure.string :as s]
            [clojure.java.io :as io]
            [clojure.java.shell :as sh])
  (:import [java.io File Writer]
           [cljs.tagged_literals JSValue]))

(defmethod print-method cljs.tagged_literals.JSValue
  [^JSValue d ^Writer w]
  (.write w "#js ")
  (print-method (.val d) w))

(defmacro emit [& body]
  `(let [res# (emit-go* '~body)]
     (is (zero? (:exit res#)) (:err res#))))

(comment
  (deftest cljs->go
    (emit (ns test.go (:require [goog.array :as array]))
          (defn plus-one [x] (inc x)))

    (emit (ns hello)
          (defn -main []
            (println "Hello World")))

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

(defn test-comment [code ast]
  (println "/*")
  (clojure.pprint/pprint code)
  (println)
  (clojure.pprint/pprint ast)
  (println"*/"))

(defn test-setup [setup]
  (let [ast (-> [setup] cljs->ast)]
    (test-comment setup ast)
    (printf "\t%s\n" (ast->go ast))))

(defn test-setup-toplevel [setup]
  (binding [cljs.analyzer/*cljs-file* (.getAbsolutePath (io/file (io/resource "cljs/go_test.clj")))
            cljs.compiler/*go-line-numbers* true]
    (with-out-str
      (test-setup setup))))

(defn testify [test & assertions]
  (binding [cljs.analyzer/*cljs-file* (.getAbsolutePath (io/file (io/resource "cljs/go_test.clj")))
            cljs.compiler/*go-line-numbers* true]
    (with-out-str
      (printf "func Test_%s(t *testing.T) {\n" (name test))
      (doseq [[expected actual setup message] assertions
              :let [ast (-> [actual] (cljs->ast (assoc (cljs.analyzer/empty-env) :context :expr)))]]
        (when setup
          (printf "{\n")
          (test-setup setup))
        (test-comment actual (first ast))
        (printf "\tassert.Equal(t%s, %s%s)\n"
                (if (nil? expected) "" (str ", \n" expected))
                (str "\n" (ast->go ast))
                (if (nil? message) "" (str ", \n`" message "`")))
        (when setup
          (printf "}\n")))
      (printf "}\n"))))

(defn emit-test [package file tests]
  (let [^File f (io/file (str "target/generated/" package) (str file ".go"))
        src (gofmt (apply str (testify-header package) tests))]
    (io/make-parents f)
    (spit f src)
    (go-test (io/file "." (.getParent f)))))

(deftest constants
  (->>
   [(testify "Constants"
             ["nil" nil]
             [true true]
             [false false]
             [1 1]
             [3.14 3.14]
             [2 '(inc 1)]
             [10 '(* 4 2.5)]
             ["`foo`" "foo"]
             ["'x'" \x]
             (binding [*data-readers* *cljs-data-readers*]
               ["js.JSObject{`foo`: `bar`}"
                (read-string "#js {:foo \"bar\"}")])
             (binding [*data-readers* *cljs-data-readers*]
               ["js.JSArray{\"foo\", \"bar\"}"
                (read-string "#js [\"foo\", \"bar\"])")])
             ["&js.Date{Millis: 1408642409602}"
              #inst "2014-08-21T17:33:29.602-00:00"]
             ["&CljsCoreUUID{Uuid: `15c52219-a8fd-4771-87e2-42ee33b79bca`}"
              #uuid "15c52219-a8fd-4771-87e2-42ee33b79bca"]
             ["&js.RegExp{Pattern: `x`, Flags: ``}" #"x"]
             ["&js.RegExp{Pattern: ``, Flags: ``}" #""]
             ["&CljsCoreSymbol{Ns: nil, Name: `x`, Str: `x`, Hash: float64(-555367584), Meta: nil}"
              ''x]
             ["&CljsCoreSymbol{Ns: `user`, Name: `x`, Str: `user/x`, Hash: float64(-568535109), Meta: nil}"
              ''user/x]
             ["&CljsCoreKeyword{Ns: nil, Name: `x`, Fqn: `x`, Hash: float64(2099068185)}"
              :x]
             ["&CljsCoreKeyword{Ns: `user`, Name: `x`, Fqn: `user/x`, Hash: float64(2085900660)}"
              :user/x])]
   (emit-test "go_test" "constants_test")))

(deftest special-forms
  (->>
   [(testify "Let"
             [1 '(let [y 1]
                   y)])
    ;; (testify "Letfn"
    ;;          ["`bar`" '(letfn [(foo [] "bar")]
    ;;                      (foo))])

    ;; func() interface{} {
    ;;  var foo IFn
    ;; 	foo = Fn(func() interface{} {
    ;; 		return js.JSString("bar")
    ;; 	})
    ;; 	return foo.Invoke_Arity0()
    ;; }()

    ;; (testify "Letfn"
    ;;          [true '(letfn [(even? [x]
    ;;                           (or (zero? x)
    ;;                               (odd? (dec x))))
    ;;                         (odd? [x]
    ;;                           (and (not (zero? x))
    ;;                                (even? (dec x))))]
    ;;                   (odd? 5))])

    ;; func() interface{} {
    ;;  var even_QMARK_, odd_QMARK_ IFn
    ;;  even_QMARK_ = Fn(func(x interface{}) interface{} {
    ;;  	var or__22923__auto__ = x == float64(0)
    ;;  	if or__22923__auto__ {
    ;;  		return or__22923__auto__
    ;;  	} else {
    ;;  		return odd_QMARK_.Invoke_Arity1((x.(float64) - float64(1)))
    ;; 		}
    ;; 	})
    ;; 	odd_QMARK_ = Fn(func(x interface{}) interface{} {
    ;; 		var and__22911__auto__ = Not.Invoke_Arity1(x == float64(0))
    ;; 		if Truth_(and__22911__auto__) {
    ;; 			return even_QMARK_.Invoke_Arity1((x.(float64) - float64(1)))
    ;; 		} else {
    ;; 			return and__22911__auto__
    ;; 		}
    ;; 	})
    ;; 	return odd_QMARK_.Invoke_Arity1(float64(5))
    ;; }()

    (testify "If"
             [true '(let [y :foo]
                      (if y true false))]
             [1 '(let [y true
                       z (if y 1 0)]
                   z)])
    (testify "Loop"
             [5 '(loop [y 0]
                   (if (== y 5)
                     y
                     (recur (inc y))))])
    (testify "Do"
             [3 '(do 1 2 3)])
    (test-setup-toplevel
     '(def x 2))
    (testify "Def"
             [2 'x])
    (testify "New"
             ["&js.Date{Millis: 0}" '(js/Date. 0)])
    (testify "Dot"
             [1970 '(.getUTCFullYear (js/Date. 0))])
    (testify "Var"
             ["math.Inf(1)" 'js/Infinity])
    (testify "Case"
             [true '(let [x 2]
                      (case x
                        2 true
                        1 false
                        0))])
    (testify "JS"
             ["reflect.Float64" '(let [x 1]
                                   (js* "reflect.ValueOf(x).Kind()"))]
             ["reflect.Int" '(let [v (js* "reflect.ValueOf(1)")]
                               (-> v .Type .Kind))])
    (testify "Try"
             ["&js.Error{`Foo`}" '(try
                                    (throw (js/Error. "Foo"))
                                    (catch js/Error e
                                      e))]
             ["`TypeError`" '(try
                               (throw (js/TypeError. "Foo"))
                               (catch js/Error _
                                 "Error")
                               (catch js/TypeError _
                                 "TypeError"))]
             ["`Bar`" '(try
                         "Bar"
                         (catch js/Error e
                           e)
                         (finally
                           "Baz"))]
             ["js.JSObject{`finally`: true}"
              '(let [x (js* "js.JSObject{}")]
                 (try
                    x
                    (finally
                      (js* "x[`finally`] = true"))))]
             ["js.JSObject{`catch`: true, `finally`: true, `last`: `finally`}"
              '(let [x (js* "js.JSObject{}")]
                 (try
                   (throw (js/Error. "Foo"))
                   (catch js/Error _
                     (js* "x[`catch`] = true")
                     (js* "x[`last`] = `catch`")
                     x)
                   (finally
                     (js* "x[`finally`] = true")
                      (js* "x[`last`] = `finally`"))))])]
   (emit-test "go_test" "special_forms_test")))

(deftest go-all-tests
  (go-test "./..."))
