(ns cljs.go-test
  (:refer-clojure :exclude [test])
  (:require [clojure.test :refer :all]
            [cljs.go :refer :all]
            [cljs.tagged-literals ]
            [clojure.pprint :as pp]
            [clojure.string :as s]
            [clojure.java.io :as io]
            [clojure.java.shell :as sh])
  (:import [java.io Writer]
           [cljs.tagged_literals JSValue]))

(defn go-test [package]
  (let [{:keys [out err exit]} (sh/sh "go" "test" (str package))
        out (s/replace (s/replace (str err out) "\r" "\n") "\n\t\t" "")]
    (is (zero? exit) out)))

(defn test-header [package & imports]
  (with-out-str
    (println "package" package)
    (println "import" "." (pr-str "github.com/hraberg/cljs.go/cljs/core"))))

(defn test-comment [code ast]
  (println "/*")
  (pp/pprint code)
  (println)
  (pp/pprint ast)
  (println"*/"))

(defn test-setup [setup]
  (with-out-str
    (let [ast (cljs->ast setup)]
      (test-comment setup ast)
      (printf "\t%s\n" (ast->go ast)))))

(defn test [test & assertions]
  (with-out-str
    (printf "func Test_%s(t *testing.T) {\n" (name test))
    (doseq [[expected actual] (partition 2 assertions)
            :let [ast (-> [actual] (cljs->ast (assoc (cljs.analyzer/empty-env) :context :expr)))]]
      (test-comment actual (first ast))
      (printf "\tassert.Equal(t,\n %s,\n %s)\n" expected (ast->go ast)))
    (printf "}\n\n")))

(defn emit-test [package file tests]
  (go-get "github.com/stretchr/testify/assert")
  (doto (io/file (str "target/generated/" package) (str file ".go"))
    io/make-parents
    (spit (goimports (apply str (test-header package) tests)))))

(defn constants []
  (->>
   [(test "Constants"
          "nil" nil
          true true
          false false
          1 1
          3.14 3.14
          2 '(inc 1)
          10 '(* 4 2.5)
          "`foo`" "foo"
          "'x'" \x
          "map[string]interface{}{`foo`: `bar`}" (read-string "#js {:foo \"bar\"}")
          "[]interface{}{\"foo\", \"bar\"}" (read-string "#js [\"foo\", \"bar\"])")
          "&js.Date{Millis: 1408642409602}" #inst "2014-08-21T17:33:29.602-00:00"
          "&CljsCoreUUID{Uuid: `15c52219-a8fd-4771-87e2-42ee33b79bca`}" #uuid "15c52219-a8fd-4771-87e2-42ee33b79bca"
          "&js.RegExp{Pattern: `x`, Flags: ``}" #"x"
          "&js.RegExp{Pattern: ``, Flags: ``}" #""
          "&CljsCoreSymbol{Ns: nil, Name: `x`, Str: `x`, Hash: float64(-555367584), Meta: nil}" ''x
          "&CljsCoreSymbol{Ns: `user`, Name: `x`, Str: `user/x`, Hash: float64(-568535109), Meta: nil}" ''user/x
          "&CljsCoreKeyword{Ns: nil, Name: `x`, Fqn: `x`, Hash: float64(2099068185)}" :x
          "&CljsCoreKeyword{Ns: `user`, Name: `x`, Fqn: `user/x`, Hash: float64(2085900660)}" :user/x)]
   (emit-test "go_test" "constants_test")))

(defn special-forms []
  (->>
   [(test "Let"
          1 '(let [y 1]
               y))
    (test "Letfn"
          "`bar`" '(letfn [(foo [] "bar")]
                     (foo))
          ;; Mutual recursion works here, but arithmetic needs unboxing or primtive fn.
          ;; true '(letfn [(even? [x]
          ;;                 (or (zero? x)
          ;;                     (odd? (dec x))))
          ;;               (odd? [x]
          ;;                 (and (not (zero? x))
          ;;                      (even? (dec x))))]
          ;;         (odd? 5))
          )
    (test "If"
          true '(let [y :foo]
                  (if y true false))
          1 '(let [y true
                   z (if y 1 0)]
               z))
    (test "Loop"
          5 '(loop [y 0]
               (if (== y 5)
                 y
                 (recur (inc y)))))
    (test "Do"
          3 '(do 1 2 3))
    (test-setup '[(def x 2)])
    (test "Def"
          2 'x)
    (test-setup '[(defn foo [] "bar")
                  (defn bar
                    ([x] x)
                    ([x y] y))
                  (defn baz
                    ([] (foo))
                    ([x] x)
                    ([x & ys] ys))])
    (test "Defn"
          "`bar`" '(foo)
          "`bar`" '(letfn [(bar [f] (f))]
                     (bar foo))
          1 '(bar 1)
          2 '(bar 1 2)
          "`bar`" '(baz)
          1 '(baz 1)
          "[]interface{}{2, 3}" '(baz 1 2 3)  ;; this should be a seq
          "[]interface{}{2, 3, 4}" (read-string "(apply baz 1 2 #js [3 4])")) ;; last arg should be a seq
    (test "New"
          "&js.Date{Millis: 0}" '(js/Date. 0))
    (test "Dot"
          1970 '(.getUTCFullYear (js/Date. 0)))
    (test "Var"
          "math.Inf(1)" 'js/Infinity)
    (test "Case"
          true '(let [x 2]
                  (case x
                    2 true
                    1 false
                    0)))
    (test "JS"
          "reflect.Float64" '(let [x 1]
                               (js* "reflect.ValueOf(x).Kind()"))
          "reflect.Int" '(let [v (js* "reflect.ValueOf(1)")]
                           (-> v .Type .Kind)))
    (test "Try"
          "&js.Error{`Foo`}" '(try
                                (throw (js/Error. "Foo"))
                                (catch js/Error e
                                  e))
          "`TypeError`" '(try
                           (throw (js/TypeError. "Foo"))
                           (catch js/Error _
                             "Error")
                           (catch js/TypeError _
                             "TypeError"))
          "`Bar`" '(try
                     "Bar"
                     (catch js/Error e
                       e)
                     (finally
                       "Baz"))
          "map[string]interface{}{`finally`: true}"
          '(let [x (js* "map[string]interface{}{}")]
             (try
               x
               (finally
                 (js* "x[`finally`] = true"))))
          "map[string]interface{}{`catch`: true, `finally`: true, `last`: `finally`}"
          '(let [x (js* "map[string]interface{}{}")]
             (try
               (throw (js/Error. "Foo"))
               (catch js/Error _
                 (js* "x[`catch`] = true")
                 (js* "x[`last`] = `catch`")
                 x)
               (finally
                 (js* "x[`finally`] = true")
                 (js* "x[`last`] = `finally`")))))]
   (emit-test "go_test" "special_forms_test")))

;; We need to maintain the same env for setup and tests.
(deftest go-all-tests
  (binding [cljs.analyzer/*cljs-file* (:file (meta #'go-test))
            ;; cljs.analyzer/*cljs-static-fns* true ;; we probably want this as default behavior
            cljs.compiler/*go-line-numbers* true
            *data-readers* cljs.tagged-literals/*cljs-data-readers*]
    (constants)
    (special-forms)
    (go-test "./...")))

(defmethod print-method cljs.tagged_literals.JSValue
  [^JSValue d ^Writer w]
  (.write w "#js ")
  (print-method (.val d) w))
