;   Copyright (c) Rich Hickey. All rights reserved.
;   The use and distribution terms for this software are covered by the
;   Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
;   which can be found in the file epl-v10.html at the root of this distribution.
;   By using this software in any fashion, you are agreeing to be bound by
;   the terms of this license.
;   You must not remove this notice, or any other, from this software.

-;; Go emitter macros. Forked from 562e6ccc9265979c586c2afc2f50a388e6c3c03e

(ns cljs.core
  (:refer-clojure :exclude [-> ->> .. amap and areduce alength aclone assert binding bound-fn case comment cond condp
                            declare definline definterface defmethod defmulti defn defn- defonce
                            defprotocol defrecord defstruct deftype delay destructure doseq dosync dotimes doto
                            extend-protocol extend-type fn for future gen-class gen-interface
                            if-let if-not import io! lazy-cat lazy-seq let letfn locking loop
                            memfn ns or proxy proxy-super pvalues refer-clojure reify sync time
                            when when-first when-let when-not while with-bindings with-in-str
                            with-loading-context with-local-vars with-open with-out-str with-precision with-redefs
                            satisfies? identical? true? false? number? nil? instance? symbol? keyword? string? str get
                            make-array vector list hash-map array-map hash-set

                            aget aset
                            + - * / < <= > >= == zero? pos? neg? inc dec max min mod
                            byte char short int long float double
                            unchecked-byte unchecked-char unchecked-short unchecked-int
                            unchecked-long unchecked-float unchecked-double
                            unchecked-add unchecked-add-int unchecked-dec unchecked-dec-int
                            unchecked-divide unchecked-divide-int unchecked-inc unchecked-inc-int
                            unchecked-multiply unchecked-multiply-int unchecked-negate unchecked-negate-int
                            unchecked-subtract unchecked-subtract-int unchecked-remainder-int
                            unsigned-bit-shift-right

                            bit-and bit-and-not bit-clear bit-flip bit-not bit-or bit-set
                            bit-test bit-shift-left bit-shift-right bit-xor

                            cond-> cond->> as-> some-> some->>

                            if-some when-some])
  (:require clojure.walk
            clojure.set
            cljs.go.compiler
            clojure.string
            [cljs.env :as env]))

(alias 'core 'clojure.core)
(alias 'ana 'cljs.analyzer)

(defmacro import-macros [ns [& vars]]
  (core/let [ns (find-ns ns)
             vars (map #(ns-resolve ns %) vars)
             syms (map (core/fn [^clojure.lang.Var v] (core/-> v .sym (with-meta {:macro true}))) vars)
             defs (map (core/fn [sym var]
                                `(do (def ~sym (deref ~var))
                                     ;for AOT compilation
                                     (alter-meta! (var ~sym) assoc :macro true)))
                       syms vars)]
            `(do ~@defs
                 :imported)))

(import-macros clojure.core
 [-> ->> .. assert comment cond
  defn defn-
  doto
  extend-protocol fn for
  if-let if-not letfn
  memfn
  when when-first when-let when-not while
  cond-> cond->> as-> some-> some->>
  if-some when-some])

(defmacro declare [& names]
  (core/doseq [n names]
    (swap! env/*compiler* assoc-in [:cljs.analyzer/namespaces ana/*cljs-ns* :defs n :declared-var] true)))

(defmacro defonce [x init]
  `(when-not (exists? ~x)
     (def ~x ~init)))

(defmacro ^{:private true} assert-args [fnname & pairs]
  `(do (when-not ~(first pairs)
         (throw (IllegalArgumentException.
                  ~(core/str fnname " requires " (second pairs)))))
     ~(core/let [more (nnext pairs)]
        (when more
          (list* `assert-args fnname more)))))

(defn destructure [bindings]
  (core/let [bents (partition 2 bindings)
         pb (fn pb [bvec b v]
              (core/let [pvec
                     (fn [bvec b val]
                       (core/let [gvec (gensym "vec__")]
                         (core/loop [ret (-> bvec (conj gvec) (conj val))
                                     n 0
                                     bs b
                                     seen-rest? false]
                           (if (seq bs)
                             (core/let [firstb (first bs)]
                               (core/cond
                                 (= firstb '&) (recur (pb ret (second bs) (core/list `nthnext gvec n))
                                                      n
                                                      (nnext bs)
                                                      true)
                                 (= firstb :as) (pb ret (second bs) gvec)
                                 :else (if seen-rest?
                                         (throw (new Exception "Unsupported binding form, only :as can follow & parameter"))
                                         (recur (pb ret firstb (core/list `nth gvec n nil))
                                                (core/inc n)
                                                (next bs)
                                                seen-rest?))))
                             ret))))
                     pmap
                     (fn [bvec b v]
                       (core/let [gmap (gensym "map__")
                                  defaults (:or b)]
                         (core/loop [ret (-> bvec (conj gmap) (conj v)
                                             (conj gmap) (conj `(if (seq? ~gmap) (apply core/hash-map ~gmap) ~gmap))
                                             ((fn [ret]
                                                (if (:as b)
                                                  (conj ret (:as b) gmap)
                                                  ret))))
                                     bes (reduce
                                          (fn [bes entry]
                                            (reduce #(assoc %1 %2 ((val entry) %2))
                                                    (dissoc bes (key entry))
                                                    ((key entry) bes)))
                                          (dissoc b :as :or)
                                          {:keys #(if (core/keyword? %) % (keyword (core/str %))),
                                           :strs core/str, :syms #(core/list `quote %)})]
                           (if (seq bes)
                             (core/let [bb (key (first bes))
                                        bk (val (first bes))
                                        has-default (contains? defaults bb)]
                               (recur (pb ret bb (if has-default
                                                   (core/list `get gmap bk (defaults bb))
                                                   (core/list `get gmap bk)))
                                      (next bes)))
                             ret))))]
                    (core/cond
                      (core/symbol? b) (-> bvec (conj (if (namespace b) (symbol (name b)) b)) (conj v))
                      (core/keyword? b) (-> bvec (conj (symbol (name b))) (conj v))
                      (vector? b) (pvec bvec b v)
                      (map? b) (pmap bvec b v)
                      :else (throw (new Exception (core/str "Unsupported binding form: " b))))))
         process-entry (fn [bvec b] (pb bvec (first b) (second b)))]
        (if (every? core/symbol? (map first bents))
          bindings
          (if-let [kwbs (seq (filter #(core/keyword? (first %)) bents))]
            (throw (new Exception (core/str "Unsupported binding key: " (ffirst kwbs))))
            (reduce process-entry [] bents)))))

(defmacro let
  "binding => binding-form init-expr

  Evaluates the exprs in a lexical context in which the symbols in
  the binding-forms are bound to their respective init-exprs or parts
  therein."
  [bindings & body]
  (assert-args
     (vector? bindings) "a vector for its binding"
     (even? (count bindings)) "an even number of forms in binding vector")
  `(let* ~(destructure bindings) ~@body))

(defmacro loop
  "Evaluates the exprs in a lexical context in which the symbols in
  the binding-forms are bound to their respective init-exprs or parts
  therein. Acts as a recur target."
  [bindings & body]
    (assert-args
      (vector? bindings) "a vector for its binding"
      (even? (count bindings)) "an even number of forms in binding vector")
    (let [db (destructure bindings)]
      (if (= db bindings)
        `(loop* ~bindings ~@body)
        (let [vs (take-nth 2 (drop 1 bindings))
              bs (take-nth 2 bindings)
              gs (map (fn [b] (if (core/symbol? b) b (gensym))) bs)
              bfs (reduce (fn [ret [b v g]]
                            (if (core/symbol? b)
                              (conj ret g v)
                              (conj ret g v b g)))
                          [] (map core/vector bs vs gs))]
          `(let ~bfs
             (loop* ~(vec (interleave gs gs))
               (let ~(vec (interleave bs gs))
                 ~@body)))))))

(defmacro str [& xs]
  ;; Eagerly stringify any string or char literals.
  (let [clean-xs (reduce (fn [acc x]
                     (core/cond
                       (core/or (core/string? x) (core/char? x))
                       (if (core/string? (peek acc))
                         (conj (pop acc) (core/str (peek acc) x))
                         (conj acc (core/str x)))
                       (core/nil? x) acc
                       :else (conj acc x)))
                   [] xs)
        ;; clean-xs now has no nils, chars, or string-adjoining-string. bools,
        ;; ints and floats will be emitted literally to allow JS string coersion.
        strs (->> clean-xs
                  (map #(if (core/string? %)
                            "~{}"
                            (core/str (cljs.compiler/go-core "Str") ".X_invoke_Arity1(~{}).(string)")))
                  (interpose "+")
                  (apply core/str))]
    ;; Google closure advanced compile will stringify and concat strings and
    ;; numbers at compilation time.
    (with-meta
     (list* 'js* (core/str (if (core/string? (first clean-xs)) "(" "(``+")
                           strs ")")
            clean-xs)
     {:tag 'string})))

(defn bool-expr [e]
  (vary-meta e assoc :tag 'boolean))

(defn simple-test-expr? [env ast]
  (core/and
    (#{:var :invoke :constant :dot :js} (:op ast))
    (= 'boolean (cljs.analyzer/infer-tag env ast))))

(defmacro and
  "Evaluates exprs one at a time, from left to right. If a form
  returns logical false (nil or false), and returns that value and
  doesn't evaluate any of the other expressions, otherwise it returns
  the value of the last expr. (and) returns true."
  ([] true)
  ([x] x)
  ([x & next]
    (let [forms (concat [x] next)]
      (if (every? #(simple-test-expr? &env %)
            (map #(cljs.analyzer/analyze &env %) forms))
        (let [and-str (->> (repeat (count forms) "(~{})")
                        (interpose " && ")
                        (apply core/str))]
          (bool-expr `(~'js* ~and-str ~@forms)))
        `(let [and# ~x]
           (if and# (and ~@next) and#))))))

(defmacro or
  "Evaluates exprs one at a time, from left to right. If a form
  returns a logical true value, or returns that value and doesn't
  evaluate any of the other expressions, otherwise it returns the
  value of the last expression. (or) returns nil."
  ([] nil)
  ([x] x)
  ([x & next]
    (let [forms (concat [x] next)]
      (if (every? #(simple-test-expr? &env %)
            (map #(cljs.analyzer/analyze &env %) forms))
        (let [or-str (->> (repeat (count forms) "(~{})")
                        (interpose " || ")
                        (apply core/str))]
          (bool-expr `(~'js* ~or-str ~@forms)))
        `(let [or# ~x]
           (if or# or# (or ~@next)))))))

(defmacro nil? [x]
  (bool-expr (core/list 'js* (core/str (cljs.compiler/go-core "Nil_") "(~{})") x)))

;; internal - do not use.
(defmacro coercive-not [x]
  (bool-expr (core/list 'js* "(!~{})" x)))

;; internal - do not use.
(defmacro coercive-not= [x y]
  (bool-expr (core/list 'js* "(~{} != ~{})" x y)))

;; internal - do not use.
(defmacro coercive-= [x y]
  (bool-expr (core/list 'js* "reflect.DeepEqual(~{}, ~{})" x y)))

;; internal - do not use.
(defmacro coercive-boolean [x]
  (with-meta (core/list 'js* "~{}" x)
    {:tag 'boolean}))

;; internal - do not use.
(defmacro truth_ [x]
  (assert (clojure.core/symbol? x) "x is substituted twice")
  (core/list 'js* "(~{} != nil && ~{} != false)" x x))

;; internal - do not use
(defmacro js-arguments []
  (core/list 'js* "arguments"))

(defmacro js-delete [obj key]
  (core/list 'js* "func(obj, key interface{}) interface{} { delete(obj.(map[string]interface{}), key.(string)); return obj }(~{}, ~{})"
             obj key))

(defmacro true? [x]
  (bool-expr (core/list 'js* "~{} == true" x)))

(defmacro false? [x]
  (bool-expr (core/list 'js* "~{} == false" x)))

(defmacro array? [x]
  (bool-expr (core/list 'js* (core/str (cljs.compiler/go-core "Value_") "(~{}).Kind() == reflect.Slice") x)))

(defmacro string? [x]
  (bool-expr (core/list 'js* (core/str (cljs.compiler/go-core "Value_") "(~{}).Kind() == reflect.String") x)))

;; TODO: x must be a symbol, not an arbitrary expression
(defmacro exists? [x]
  (bool-expr
   (core/list 'js* (core/str (cljs.compiler/go-core "Value_") "(~{}).Kind() != reflect.Invalid")
      (vary-meta x assoc :cljs.analyzer/no-resolve true))))

(defmacro undefined? [x]
  (bool-expr (core/list 'js* (core/str (cljs.compiler/go-core "Nil_") "(~{})") x)))

(defmacro identical? [a b]
  (bool-expr (core/list 'js* "reflect.DeepEqual(~{}, ~{})" a b))) ;; this is wrong, but kind of works

(defmacro instance? [t o]
  ;; Google Closure warns about some references to RegExp, so
  ;; (instance? RegExp ...) needs to be inlined, but the expansion
  ;; should preserve the order of argument evaluation.
  (bool-expr `(~'js* ~(core/str (cljs.compiler/go-core "Value_") "(~{}).Type().AssignableTo(~{})") ~o ~t)))

(defmacro number? [x]
  (bool-expr (core/list 'js* (core/str (cljs.compiler/go-core "Value_") "(~{}).Kind() == reflect.Float64") x)))

(defmacro boolean? [x]
  (bool-expr (core/list 'js* (core/str (cljs.compiler/go-core "Value_") "(~{}).Kind() == reflect.Bool") x)))

(defmacro symbol? [x]
  (bool-expr `(instance? Symbol ~x)))

(defmacro keyword? [x]
  (bool-expr `(instance? Keyword ~x)))

(defmacro aget
  ([a i]
     (core/list 'js* "(~{}[int(~{})])" a i))
  ([a i & idxs]
     (let [astr (apply core/str (repeat (count idxs) ".([]interface{})[int(~{})]"))]
      `(~'js* ~(core/str "(~{}[int(~{})]" astr ")") ~a ~i ~@idxs))))

(defmacro aset
  ([a i v]
    (core/list 'js* "~{}[int(~{})] = ~{}" a i v))
  ([a idx idx2 & idxv]
    (let [n    (core/dec (count idxv))
          astr (apply core/str (repeat n ".([]interface{})[int(~{})]"))]
      `(~'js* ~(core/str "~{}[int(~{})].([]interface{})[int(~{})]" astr " = ~{}") ~a ~idx ~idx2 ~@idxv))))

(defmacro ^::ana/numeric +
  ([] 0)
  ([x] x)
  ([x y] (core/list 'js* "(~{} + ~{})" x y))
  ([x y & more] `(+ (+ ~x ~y) ~@more)))

(defmacro byte [x] x)
(defmacro short [x] x)
(defmacro float [x] x)
(defmacro double [x] x)

(defmacro unchecked-byte [x] x)
(defmacro unchecked-char [x] x)
(defmacro unchecked-short [x] x)
(defmacro unchecked-float [x] x)
(defmacro unchecked-double [x] x)

(defmacro ^::ana/numeric unchecked-add
  ([& xs] `(+ ~@xs)))

(defmacro ^::ana/numeric unchecked-add-int
  ([& xs] `(+ ~@xs)))

(defmacro ^::ana/numeric unchecked-dec
  ([x] `(dec ~x)))

(defmacro ^::ana/numeric unchecked-dec-int
  ([x] `(dec ~x)))

(defmacro ^::ana/numeric unchecked-divide-int
  ([& xs] `(/ ~@xs)))

(defmacro ^::ana/numeric unchecked-inc
  ([x] `(inc ~x)))

(defmacro ^::ana/numeric unchecked-inc-int
  ([x] `(inc ~x)))

(defmacro ^::ana/numeric unchecked-multiply
  ([& xs] `(* ~@xs)))

(defmacro ^::ana/numeric unchecked-multiply-int
  ([& xs] `(* ~@xs)))

(defmacro ^::ana/numeric unchecked-negate
  ([x] `(- ~x)))

(defmacro ^::ana/numeric unchecked-negate-int
  ([x] `(- ~x)))

(defmacro ^::ana/numeric unchecked-remainder-int
  ([x n] `(mod ~x ~n)))

(defmacro ^::ana/numeric unchecked-subtract
  ([& xs] `(- ~@xs)))

(defmacro ^::ana/numeric unchecked-subtract-int
  ([& xs] `(- ~@xs)))

(defmacro ^::ana/numeric -
  ([x] (core/list 'js* "(- ~{})" x))
  ([x y] (core/list 'js* "(~{} - ~{})" x y))
  ([x y & more] `(- (- ~x ~y) ~@more)))

(defmacro ^::ana/numeric *
  ([] 1)
  ([x] x)
  ([x y] (core/list 'js* "(~{} * ~{})" x y))
  ([x y & more] `(* (* ~x ~y) ~@more)))

(defmacro ^::ana/numeric /
  ([x] `(/ 1 ~x))
  ([x y] (core/list 'js* "(~{} / ~{})" x y))
  ([x y & more] `(/ (/ ~x ~y) ~@more)))

(defmacro ^::ana/numeric divide
  ([x] `(/ 1 ~x))
  ([x y] (core/list 'js* "(~{} / ~{})" x y))
  ([x y & more] `(/ (/ ~x ~y) ~@more)))

(defmacro ^::ana/numeric <
  ([x] true)
  ([x y] (bool-expr (core/list 'js* "(~{} < ~{})" x y)))
  ([x y & more] `(and (< ~x ~y) (< ~y ~@more))))

(defmacro ^::ana/numeric <=
  ([x] true)
  ([x y] (bool-expr (core/list 'js* "(~{} <= ~{})" x y)))
  ([x y & more] `(and (<= ~x ~y) (<= ~y ~@more))))

(defmacro ^::ana/numeric >
  ([x] true)
  ([x y] (bool-expr (core/list 'js* "(~{} > ~{})" x y)))
  ([x y & more] `(and (> ~x ~y) (> ~y ~@more))))

(defmacro ^::ana/numeric >=
  ([x] true)
  ([x y] (bool-expr (core/list 'js* "(~{} >= ~{})" x y)))
  ([x y & more] `(and (>= ~x ~y) (>= ~y ~@more))))

(defmacro ^::ana/numeric ==
  ([x] true)
  ([x y] (bool-expr (core/list 'js* "(~{} == ~{})" x y)))
  ([x y & more] `(and (== ~x ~y) (== ~y ~@more))))

(defmacro ^::ana/numeric dec [x]
  `(- ~x 1))

(defmacro ^::ana/numeric inc [x]
  `(+ ~x 1))

(defmacro ^::ana/numeric zero? [x]
  `(== ~x 0))

(defmacro ^::ana/numeric pos? [x]
  `(> ~x 0))

(defmacro ^::ana/numeric neg? [x]
  `(< ~x 0))

(defmacro ^::ana/numeric max
  ([x] x)
  ([x y] `(~'js* "func(x, y float64) float64 { if x > y { return x } else { return y } }(~{}, ~{})" ~x ~y))
  ([x y & more] `(max (max ~x ~y) ~@more)))

(defmacro ^::ana/numeric min
  ([x] x)
  ([x y] `(~'js* "func(x, y float64) float64 { if x < y { return x } else { return y } }(~{}, ~{})" ~x ~y))
  ([x y & more] `(min (min ~x ~y) ~@more)))

(defmacro ^::ana/numeric js-mod [num div]
  (core/list 'js* "math.Mod(~{}, ~{})" num div))

(defn with-32bit-ints
  ([s] (with-32bit-ints s true))
  ([s float64]
     (let [s (-> s
                 (clojure.string/replace "uint" (cljs.compiler/go-core "UInt32_"))
                 (clojure.string/replace "int" (cljs.compiler/go-core "Int32_")))]
       (if float64
         (core/str "float64(" s ")")
         s))))

(defmacro ^::ana/numeric bit-not [x]
  (core/list 'js* (with-32bit-ints "(^ int(~{}))") x))

(defmacro ^::ana/numeric bit-and
  ([x y] (core/list 'js* (with-32bit-ints "(int(~{}) & int(~{}))") x y))
  ([x y & more] `(bit-and (bit-and ~x ~y) ~@more)))

;; internal do not use
(defmacro ^::ana/numeric unsafe-bit-and
  ([x y] (bool-expr (core/list 'js* (with-32bit-ints "(int(~{}) & int(~{}))") x y)))
  ([x y & more] `(unsafe-bit-and (unsafe-bit-and ~x ~y) ~@more)))

(defmacro ^::ana/numeric bit-or
  ([x y] (core/list 'js* (with-32bit-ints "(int(~{}) | int(~{}))") x y))
  ([x y & more] `(bit-or (bit-or ~x ~y) ~@more)))

(defmacro ^::ana/numeric int [x]
  `(~'js* ~(with-32bit-ints "int(~{})") ~x))

(defmacro ^::ana/numeric bit-xor
  ([x y] (core/list 'js* (with-32bit-ints "(int(~{}) ^ int(~{}))") x y))
  ([x y & more] `(bit-xor (bit-xor ~x ~y) ~@more)))

(defmacro ^::ana/numeric bit-and-not
  ([x y] (core/list 'js* (with-32bit-ints "(int(~{}) &^ int(~{}))") x y))
  ([x y & more] `(bit-and-not (bit-and-not ~x ~y) ~@more)))

(defmacro ^::ana/numeric bit-clear [x n]
  (core/list 'js* (with-32bit-ints "(int(~{}) &^ (int(1) << uint(~{})))") x n))

(defmacro ^::ana/numeric bit-flip [x n]
  (core/list 'js* (with-32bit-ints "(int(~{}) ^ (int(1) << uint(~{})))") x n))

(defmacro ^::ana/numeric bit-test [x n]
  (bool-expr (core/list 'js* (with-32bit-ints "((int(~{}) & (int(1) << uint(~{})))) != 0" false) x n)))

(defmacro ^::ana/numeric bit-shift-left [x n]
  (core/list 'js* (with-32bit-ints "(int(~{}) << uint(~{}))") x n))

(defmacro ^::ana/numeric bit-shift-right [x n]
  (core/list 'js* (with-32bit-ints "(int(~{}) >> uint(~{}))") x n))

(defmacro ^::ana/numeric bit-shift-right-zero-fill [x n]
  (core/list 'js* (with-32bit-ints "(uint(~{}) >> uint(float64((32 + int(~{})) % 32)))") x n))

(defmacro ^::ana/numeric unsigned-bit-shift-right [x n]
  (core/list 'js* (with-32bit-ints "(uint(~{}) >> uint(float64((32 + int(~{})) % 32)))") x n))

(defmacro ^::ana/numeric bit-set [x n]
  (core/list 'js* (with-32bit-ints "(int(~{}) | (int(1) << uint(~{})))") x n))

;; internal
(defmacro ^::ana/numeric mask [hash shift]
  (core/list 'js* (with-32bit-ints "((uint(~{}) >> uint(~{})) & 0x01f)") hash shift))

;; internal
(defmacro ^::ana/numeric bitpos [hash shift]
  (core/list 'js* (with-32bit-ints "(int(1) << uint(~{}))") `(mask ~hash ~shift)))

;; internal
(defmacro caching-hash [coll hash-fn hash-key]
  (assert (clojure.core/symbol? hash-key) "hash-key is substituted twice")
  `(let [h# ~hash-key]
     (if-not (nil? h#)
       h#
       (let [h# (~hash-fn ~coll)]
         (set! ~hash-key h#)
         h#))))

;;; internal -- reducers-related macros

(defn- do-curried
  [name doc meta args body]
  (let [cargs (vec (butlast args))]
    `(defn ~name ~doc ~meta
       (~cargs (fn [x#] (~name ~@cargs x#)))
       (~args ~@body))))

(defmacro ^:private defcurried
  "Builds another arity of the fn that returns a fn awaiting the last
  param"
  [name doc meta args & body]
  (do-curried name doc meta args body))

(defn- do-rfn [f1 k fkv]
  `(fn
     ([] (~f1))
     ~(clojure.walk/postwalk
       #(if (sequential? %)
          ((if (vector? %) vec identity)
           (core/remove #{k} %))
          %)
       fkv)
     ~fkv))

(defmacro ^:private rfn
  "Builds 3-arity reducing fn given names of wrapped fn and key, and k/v impl."
  [[f1 k] fkv]
  (do-rfn f1 k fkv))

;;; end of reducers macros

(defn protocol-prefix [psym]
  (core/str (-> (core/str psym) (.replace \. \$) (.replace \/ \$)) "$"))

(def #^:private base-type
     {nil "null"
      'object "object"
      'string "string"
      'number "number"
      'array "array"
      'function "function"
      'boolean "boolean"
      'default "_"
      'js/Date "date" ;; hack
      })

(def #^:private js-base-type
     {'js/Boolean "boolean"
      'js/String "string"
      'js/Array "array"
      'js/Object "object"
      'js/Number "number"
      'js/Function "function"})

(defmacro reify [& impls]
  (let [t        (with-meta (gensym "t") {:anonymous true})
        meta-sym (gensym "meta")
        this-sym (gensym "_")
        locals   (keys (:locals &env))
        ns       (-> &env :ns :name)
        munge    cljs.compiler/munge]
    `(do
       (deftype ~t [~@locals ~meta-sym]
         ~'IWithMeta
         (~'-with-meta [~this-sym ~meta-sym]
           (new ~t ~@locals ~meta-sym))
         ~'IMeta
         (~'-meta [~this-sym] ~meta-sym)
         ~@impls)
       (new ~t ~@locals nil))))

(defmacro specify! [expr & impls]
  (let [x (with-meta (gensym "x") {:extend :instance})]
    `(let [~x ~expr]
       (extend-type ~x ~@impls)
       ~x)))

(defmacro specify [expr & impls]
  `(cljs.core/specify! (cljs.core/clone ~expr)
     ~@impls))

(defn to-property [sym]
  (symbol (core/str "-" sym)))

(defn warn-and-update-protocol [p type env]
  (when-not (= 'Object p)
    (if-let [var (cljs.analyzer/resolve-existing-var (dissoc env :locals) p)]
      (do
        (when-not (:protocol-symbol var)
          (cljs.analyzer/warning :invalid-protocol-symbol env {:protocol p}))
        (when (core/and (:protocol-deprecated cljs.analyzer/*cljs-warnings*)
                (-> var :deprecated)
                (not (-> p meta :deprecation-nowarn)))
          (cljs.analyzer/warning :protocol-deprecated env {:protocol p}))
        (when (:protocol-symbol var)
          (swap! env/*compiler* update-in [:cljs.analyzer/namespaces]
            (fn [ns]
              (update-in ns [(:ns var) :defs (symbol (name p)) :impls]
                conj type)))))
      (when (:undeclared cljs.analyzer/*cljs-warnings*)
        (cljs.analyzer/warning :undeclared-protocol-symbol env {:protocol p})))))

(defn resolve-var [env sym]
  (let [ret (-> (dissoc env :locals)
              (cljs.analyzer/resolve-var sym)
              :name)]
    (assert ret (core/str "Can't resolve: " sym))
    ret))

(defn ->impl-map [impls]
  (loop [ret {} s impls]
    (if (seq s)
      (recur (assoc ret (first s) (take-while seq? (next s)))
        (drop-while seq? (next s)))
      ret)))

;; For core.cljs we can do this manually as native overrides for now
(defn base-assign-impls [env resolve tsym type [p sigs]]
  (warn-and-update-protocol p tsym env)
  (let [psym       (resolve p)
        pfn-prefix (subs (core/str psym) 0
                     (clojure.core/inc (.indexOf (core/str psym) "/")))]
    ;; (cons `(aset ~psym ~type true)
    ;;   (map (fn [[f & meths :as form]]
    ;;          `(aset ~(symbol (core/str pfn-prefix f))
    ;;             ~type ~(with-meta `(fn ~@meths) (meta form))))
    ;;     sigs))
    ))

(core/defmulti extend-prefix (fn [tsym sym] (-> tsym meta :extend)))

(core/defmethod extend-prefix :instance
  [tsym sym] `(.. ~tsym ~(to-property sym)))

(core/defmethod extend-prefix :default
  [tsym sym] `(.. ~tsym -prototype ~(to-property sym)))

(defn real-tag [tag object?]
  (if object?
    tag
    ('{string nil} tag tag))  )

(defn adapt-proto-params [type proto-fn [[this & args :as sig] & body]]
  (let [[proto-params] ((group-by count (:method-params proto-fn)) (count sig))
        this (vary-meta this assoc :tag type)
        object? (= (:protocol proto-fn) 'cljs.core/Object)]
    `(~(vec (cons this (map #(vary-meta % assoc :tag (real-tag (-> %2 meta :tag) object?))
                            args (core/or (next proto-params) args))))
      ~@body)))

;; In ClojureScript, Object isn't really a protocol per-se, but a way of attaching native fns to a type.
(defn add-obj-methods [type type-sym sigs]
  (map (fn [[f & meths :as form]]
         (let [proto-fn (get-in (ana/get-namespace 'cljs.core) [:defs f])
               proto-fn (when (= (:protocol proto-fn) 'cljs.core/Object)
                          proto-fn)]
           `(do
              ~(with-meta `(fn ~(vary-meta f assoc :tag (or (:ret-tag proto-fn) (-> f meta :tag)))
                             ~@(map #(adapt-proto-params type proto-fn %) meths)) (meta form))
              ~(when (= 'toString f)
                 (with-meta `(fn ~(with-meta 'string {:tag (:ret-tag proto-fn)})
                               ~(adapt-proto-params type proto-fn '[[this] ^string (.toString this)])) (meta form))))))
       sigs))

(def ^:private arity-limit 20)

(defn add-ifn-methods [type type-sym [f & meths :as form]]
  (let [proto-fn (get-in (ana/get-namespace 'cljs.core) [:defs f])
        ;; we don't support provided varargs here for now, so we drop the last rest fn.
        meths (remove #(core/> (core/dec (count (first %))) arity-limit) meths)
        f (with-meta f {:tag (:ret-tag proto-fn)})
        missing (for [params (:method-params proto-fn)
                      :let [body `(throw (js/Error. ~(core/str "Invalid arity: " (core/dec (count params)))))]]
                  [(count params)
                   (with-meta `(fn ~f ~(adapt-proto-params type proto-fn `[~params ~body])) (meta form))])
        provided (for [meth meths]
                   [(count (first meth))
                    (with-meta `(fn ~f ~(adapt-proto-params type proto-fn meth)) (meta form))])]
    (vals (merge (into (sorted-map) missing) (into {} provided)))))

(defn proto-slot-name [fname sig]
  (core/str (-> fname cljs.compiler/munge cljs.compiler/go-public)
            "_Arity" (cond-> (count sig) (= '-invoke fname) core/dec)))

(defn add-proto-methods* [psym type type-sym [f & meths :as form]]
  (let [proto-fn (get-in (ana/get-namespace (symbol (namespace psym))) [:defs f])]
    (map (fn [[sig & body :as meth]]
           `(do
              ~(with-meta `(fn ~(vary-meta f assoc :tag (real-tag (:ret-tag proto-fn) false))
                             ~(adapt-proto-params type proto-fn meth)) (meta form))))
         (cond-> meths (vector? (first meths)) core/vector))))

(defn proto-assign-impls [env resolve type-sym type [p sigs]]
  (warn-and-update-protocol p type env)
  (let [psym      (resolve p)]
    (when-not (get-in cljs.compiler/*go-skip-protocol* [type psym])
      (cons
       `(do ^:top-level (~'js* ~(core/str "func (_ " (cljs.compiler/go-type type) ")~{}__() {}")
                               ~(symbol (cljs.compiler/go-type-fqn psym))))
       (if (= p 'Object)
         (add-obj-methods type type-sym sigs)
         (concat
          (mapcat
           (fn [sig]
             (if (= psym 'cljs.core/IFn)
               (add-ifn-methods type type-sym sig)
               (add-proto-methods* psym type type-sym sig)))
           sigs)))))))

(core/declare dt->et collect-protocols)

(def go-skip-type '#{cljs.core/ObjMap})

(defn validate-impl-sigs [env p method]
  (when-not (= p 'Object)
    (let [var (ana/resolve-var (dissoc env :locals) p)
          minfo (-> var :protocol-info :methods)
          [fname sigs] (if (core/vector? (second method))
                         [(first method) [(second method)]]
                         [(first method) (map first (rest method))])
          decmeths (core/get minfo fname ::not-found)]
      (when (= decmeths ::not-found)
        (ana/warning :protocol-invalid-method env {:protocol p :fname fname :no-such-method true}))
      (loop [sigs sigs seen #{}]
        (when (seq sigs)
          (let [sig (first sigs)
                c   (count sig)]
            (when (contains? seen c)
              (ana/warning :protocol-duped-method env {:protocol p :fname fname}))
            (when (core/and (not= decmeths ::not-found) (not (some #{c} (map count decmeths))))
              (ana/warning :protocol-invalid-method env {:protocol p :fname fname :invalid-arity c}))
            (recur (next sigs) (conj seen c))))))))

(defn validate-impls [env impls]
  (loop [protos #{} impls impls]
    (when (seq impls)
      (let [proto   (first impls)
            methods (take-while seq? (next impls))
            impls   (drop-while seq? (next impls))]
        (when (contains? protos proto)
          (ana/warning :protocol-multiple-impls env {:protocol proto}))
        (core/doseq [method methods]
          (validate-impl-sigs env proto method))
        (recur (conj protos proto) impls)))))

(defmacro extend-type [type-sym & impls]
  (let [env &env
        _ (validate-impls env impls)
        resolve (partial resolve-var env)
        [type assign-impls] (if-let [type (base-type type-sym)]
                              [type base-assign-impls]
                              [(resolve type-sym) proto-assign-impls])
        fq-type-sym (symbol type)
        extending-existing-type? (not (some (comp :protocol-impl meta) impls))
        real-type-in-current-ns? (= ana/*cljs-ns* (some-> fq-type-sym namespace symbol))
        impls (if (and extending-existing-type? real-type-in-current-ns?)
                ;; we need to enrich here, but JS ClojureScript doesn't.
                (dt->et (vary-meta type-sym assoc :protocols (collect-protocols impls env))
                        impls (cljs.compiler/go-fields-of-type fq-type-sym) true)
                impls)
        impl-map (->impl-map impls)]
    (when-not (go-skip-type fq-type-sym)
      (when (core/and (:extending-base-js-type cljs.analyzer/*cljs-warnings*)
                      (js-base-type type-sym))
        (cljs.analyzer/warning :extending-base-js-type env
                               {:current-symbol type-sym :suggested-symbol (js-base-type type-sym)}))
      `(do ~@(mapcat #(assign-impls env resolve type-sym type %) impl-map)))))

(defn annotate-specs [annots v [f sigs]]
  (conj v
    (vary-meta (cons f (map #(cons (second %) (nnext %)) sigs))
      merge annots)))

(defn dt->et
  ([type specs fields]
    (dt->et type specs fields false))
  ([type specs fields inline]
    (let [annots {:cljs.analyzer/type type
                  :protocol-impl true
                  :protocol-inline inline}
          protocols (group-by (comp symbol name) (:protocols (meta type)))]
      (loop [ret [] specs specs]
        (if (seq specs)
          (let [p (first specs)
                ret (-> (conj ret p)
                        (into (reduce (partial annotate-specs
                                               (assoc annots :protocol-impl
                                                      (first (protocols p)))) []
                              (group-by first (take-while seq? (next specs))))))
                specs (drop-while seq? (next specs))]
            (recur ret specs))
          ret)))))

(defn collect-protocols [impls env]
  (->> impls
       (filter core/symbol?)
       (map #(:name (cljs.analyzer/resolve-var (dissoc env :locals) %)))
       (into #{})))

(defn- build-positional-factory
  [rsym rname fields]
  (let [fn-name (with-meta (symbol (core/str '-> rsym))
                  (assoc (meta rsym) :factory :positional))]
    `(defn ~fn-name
       [~@fields]
       (new ~rname ~@fields))))

(defmacro deftype [t fields & impls]
  (let [env &env
        r (:name (cljs.analyzer/resolve-var (dissoc env :locals) t))
        protocols (collect-protocols impls env)
        t (vary-meta t assoc
            :protocols protocols) ]
    (when-not (go-skip-type r)
      `(do
         (deftype* ~t ~fields nil
           ~(if (seq impls)
              `(extend-type ~t ~@(dt->et t impls fields true))))
         ~(build-positional-factory t r fields)
         ~t))))

(defn- emit-defrecord
  "Do not use this directly - use defrecord"
  [env tagname rname fields impls]
  (let [hinted-fields fields
        fields (vec (map #(with-meta % nil) fields))
        base-fields fields
        pr-open (core/str "#" (.getNamespace rname) "." (.getName rname) "{")
        fields (conj fields '__meta '__extmap (with-meta '__hash {:mutable true}))]
    (let [gs (gensym)
          ksym (gensym "k")
          impls (concat
                 impls
                 ['IRecord
                  'ICloneable
                  `(~'-clone [this#] (new ~tagname ~@fields))
                  'IHash
                  `(~'-hash [this#] (caching-hash this# ~'hash-imap ~'__hash))
                  'IEquiv
                  `(~'-equiv [this# other#]
                             (if (and other#
                                      (identical? (type this#)
                                                  (type other#))
                                      (equiv-map this# other#))
                               true
                               false))
                  'IMeta
                  `(~'-meta [this#] ~'__meta)
                  'IWithMeta
                  `(~'-with-meta [this# ~gs] (new ~tagname ~@(replace {'__meta gs} fields)))
                  'ILookup
                  `(~'-lookup [this# k#] (-lookup this# k# nil))
                  `(~'-lookup [this# ~ksym else#]
                              (case ~ksym
                                ~@(mapcat (fn [f] [(keyword f) f]) base-fields)
                                (get ~'__extmap ~ksym else#)))
                  'ICounted
                  `(~'-count [this#] (+ ~(count base-fields) (count ~'__extmap)))
                  'ICollection
                  `(~'-conj [this# entry#]
                            (if (vector? entry#)
                              (-assoc this# (-nth entry# 0) (-nth entry# 1))
                              (reduce -conj
                                      this#
                                      entry#)))
                  'IAssociative
                  `(~'-assoc [this# k# ~gs]
                             (condp keyword-identical? k#
                               ~@(mapcat (fn [fld]
                                           [(keyword fld) (list* `new tagname (replace {fld gs '__hash nil} fields))])
                                         base-fields)
                               (new ~tagname ~@(remove #{'__extmap '__hash} fields) (assoc ~'__extmap k# ~gs) nil)))
                  `(~'-contains-key? [this# k#]
                                     ^boolean (throw (js/Error. "Not implemented.")))
                  'IMap
                  `(~'-dissoc [this# k#] (if (contains? #{~@(map keyword base-fields)} k#)
                                           (dissoc (with-meta (into {} this#) ~'__meta) k#)
                                           (new ~tagname ~@(remove #{'__extmap '__hash} fields)
                                                (not-empty (dissoc ~'__extmap k#))
                                                nil)))
                  'ISeqable
                  `(~'-seq [this#] (seq (concat [~@(map #(core/list `vector (keyword %) %) base-fields)]
                                                ~'__extmap)))

                  'IPrintWithWriter
                  `(~'-pr-writer [this# writer# opts#]
                                 (let [pr-pair# (fn [keyval# _# _#] (pr-sequential-writer writer# pr-writer "" " " "" opts# keyval#))]
                                   (pr-sequential-writer
                                    writer# pr-pair# ~pr-open ", " "}" opts#
                                    (concat [~@(map #(core/list `vector (keyword %) %) base-fields)]
                                            ~'__extmap))))
                  ])
          protocols (collect-protocols impls env)
          tagname (vary-meta tagname assoc
                    :protocols protocols)]
      `(do
         (~'defrecord* ~tagname ~hinted-fields nil
           (extend-type ~tagname ~@(dt->et tagname impls fields true)))))))

(defn- build-map-factory [rsym rname fields]
  (let [fn-name (with-meta (symbol (core/str 'map-> rsym))
                  (assoc (meta rsym) :factory :map))
        ms (gensym)
        ks (map keyword fields)
        getters (map (fn [k] `(~k ~ms)) ks)]
    `(defn ~fn-name [~ms]
       (new ~rname ~@getters nil (dissoc ~ms ~@ks)))))

(defmacro defrecord [rsym fields & impls]
  (let [rsym (vary-meta rsym assoc :internal-ctor true)
        r    (vary-meta
               (:name (cljs.analyzer/resolve-var (dissoc &env :locals) rsym))
               assoc :internal-ctor true)]
    `(let []
       ~(emit-defrecord &env rsym r fields impls)
       ~(build-positional-factory rsym r fields)
       ~(build-map-factory rsym r fields)
       ~r)))

(defmacro defprotocol [psym & doc+methods]
  (let [p (:name (cljs.analyzer/resolve-var (dissoc &env :locals) psym))
        psym (vary-meta psym assoc :protocol-symbol true)
        ns-name (-> &env :ns :name)
        fq-psym (symbol (some-> ns-name name) (name psym))
        object? (= 'cljs.core/Object p)
        go-psym (symbol (cljs.compiler/go-type fq-psym))
        methods (if (core/string? (first doc+methods)) (next doc+methods) doc+methods)
        expand-sig (fn [fname sig]
                     `(~sig
                       ;; this should really just be a protocol call emitted by :invoke,
                       ;; (~fname ~(vary-meta (first sig) assoc :tag fq-psym) ~@(rest sig))
                       ~(vary-meta
                         `(~'js* ~(core/str (cljs.compiler/go-core "Decorate_")
                                            "(" (cljs.compiler/munge (first sig)) ").(" go-psym ")." (proto-slot-name fname sig)
                                            "(" (apply core/str (interpose ", " (map cljs.compiler/munge (rest sig)))) ")"))
                         assoc :tag (real-tag (-> fname meta :tag) object?))))
        psym   (vary-meta psym assoc-in [:protocol-info :methods]
                 (into {}
                   (map
                     (fn [[fname & sigs]]
                       (let [sigs (take-while vector? sigs)]
                         [fname (vec sigs)]))
                     methods)))
        method-sigs (fn [arity-limit sigs] (take arity-limit (sort-by count (take-while vector? sigs))))
        method (fn [[fname & sigs]]
                 (let [fname (vary-meta fname assoc :protocol p :tag (real-tag (-> fname meta :tag) object?))]
                   `(defn ~fname ~@(map (partial expand-sig fname) (method-sigs (core/inc arity-limit) sigs)))))
        method-decl (fn [[fname & sigs]]
                      (->>
                       (for [sig (method-sigs (core/inc arity-limit) sigs)]
                         (core/str "\t"
                                   (proto-slot-name fname sig)
                                   "("
                                   (->> (rest sig)
                                        (map #(core/str (cljs.compiler/munge %) " "
                                                        (cljs.compiler/go-type (-> % meta :tag) object?)))
                                        (interpose ", ")
                                        (apply core/str))
                                   ")"
                                   " " (cljs.compiler/go-type (-> fname meta :tag) object?)
                                   "\n"))
                       (apply core/str)))]
    `(do
       (def ~psym) ;; Empty init gets dropped by the compiler, but registered by the analyzer.
       ^:top-level (~'js* ~(core/str "type ~{} interface{\n" (core/str go-psym "__()\n")
                                     (apply core/str (map method-decl methods)) "}\n")
                          ~go-psym)
       ~@(map method methods))))

(defmacro native-satisfies?
  "EXPERIMENTAL"
  [psym x]
  (let [p (:name
           (cljs.analyzer/resolve-var
            (dissoc &env :locals) psym))]
    (bool-expr `(~'js* ~(core/str (cljs.compiler/go-core "DecoratedValue_") "(~{}).Type().Implements(~{})") ~x ~p))))

(defmacro implements?
  "EXPERIMENTAL"
  [psym x]
  `(cljs.core/native-satisfies? ~psym ~x))

(defmacro satisfies?
  "Returns true if x satisfies the protocol"
  [psym x]
  `(cljs.core/native-satisfies? ~psym ~x))

(defmacro lazy-seq [& body]
  `(new cljs.core/LazySeq nil (fn [] ~@body) nil nil))

(defmacro delay [& body]
  "Takes a body of expressions and yields a Delay object that will
  invoke the body only the first time it is forced (with force or deref/@), and
  will cache the result and return it on all subsequent force
  calls."
  `(new cljs.core/Delay (fn [] ~@body) nil))

(defmacro with-redefs
  "binding => var-symbol temp-value-expr

  Temporarily redefines vars while executing the body.  The
  temp-value-exprs will be evaluated and each resulting value will
  replace in parallel the root value of its var.  After the body is
  executed, the root values of all the vars will be set back to their
  old values. Useful for mocking out functions during testing."
  [bindings & body]
  (let [names (take-nth 2 bindings)
        vals (take-nth 2 (drop 1 bindings))
        tempnames (map (comp gensym name) names)
        binds (map core/vector names vals)
        resets (reverse (map core/vector names tempnames))
        bind-value (fn [[k v]] (core/list 'set! k v))]
    `(let [~@(interleave tempnames names)]
       (try
        ~@(map bind-value binds)
        ~@body
        (finally
         ~@(map bind-value resets))))))

(defmacro binding
  "binding => var-symbol init-expr

  Creates new bindings for the (already-existing) vars, with the
  supplied initial values, executes the exprs in an implicit do, then
  re-establishes the bindings that existed before.  The new bindings
  are made in parallel (unlike let); all init-exprs are evaluated
  before the vars are bound to their new values."
  [bindings & body]
  (let [names (take-nth 2 bindings)]
    (cljs.analyzer/confirm-bindings &env names)
    `(with-redefs ~bindings ~@body)))

(defmacro condp
  "Takes a binary predicate, an expression, and a set of clauses.
  Each clause can take the form of either:

  test-expr result-expr

  test-expr :>> result-fn

  Note :>> is an ordinary keyword.

  For each clause, (pred test-expr expr) is evaluated. If it returns
  logical true, the clause is a match. If a binary clause matches, the
  result-expr is returned, if a ternary clause matches, its result-fn,
  which must be a unary function, is called with the result of the
  predicate as its argument, the result of that call being the return
  value of condp. A single default expression can follow the clauses,
  and its value will be returned if no clause matches. If no default
  expression is provided and no clause matches, an
  IllegalArgumentException is thrown."
  {:added "1.0"}

  [pred expr & clauses]
  (let [gpred (gensym "pred__")
        gexpr (gensym "expr__")
        emit (fn emit [pred expr args]
               (let [[[a b c :as clause] more]
                       (split-at (if (= :>> (second args)) 3 2) args)
                       n (count clause)]
                 (core/cond
                  (= 0 n) `(throw (js/Error. (core/str "No matching clause: " ~expr)))
                  (= 1 n) a
                  (= 2 n) `(if (~pred ~a ~expr)
                             ~b
                             ~(emit pred expr more))
                  :else `(if-let [p# (~pred ~a ~expr)]
                           (~c p#)
                           ~(emit pred expr more)))))
        gres (gensym "res__")]
    `(let [~gpred ~pred
           ~gexpr ~expr]
       ~(emit gpred gexpr clauses))))

(defn- assoc-test [m test expr env]
  (if (contains? m test)
    (throw
      (clojure.core/IllegalArgumentException.
        (core/str "Duplicate case test constant '"
          test "'"
          (when (:line env)
            (core/str " on line " (:line env) " "
              cljs.analyzer/*cljs-file*)))))
    (assoc m test expr)))

(defmacro case [e & clauses]
  (core/let [default (if (odd? (count clauses))
                       (last clauses)
                       `(throw
                          (js/Error.
                            (core/str "No matching clause: " ~e))))
             env     &env
             pairs   (reduce
                       (fn [m [test expr]]
                         (core/cond
                           (seq? test)
                           (reduce
                             (fn [m test]
                               (let [test (if (core/symbol? test)
                                            (core/list 'quote test)
                                            test)]
                                 (assoc-test m test expr env)))
                             m test)
                           (core/symbol? test)
                           (assoc-test m (core/list 'quote test) expr env)
                           :else
                           (assoc-test m test expr env)))
                     {} (partition 2 clauses))
             esym    (gensym)
             tests   (keys pairs)]
    (cond
      (every? (some-fn core/number? core/string? core/char?) tests)
      (core/let [no-default (if (odd? (count clauses)) (butlast clauses) clauses)
                 tests      (mapv #(if (seq? %) (vec %) [%]) (take-nth 2 no-default))
                 thens      (vec (take-nth 2 (drop 1 no-default)))]
        `(let [~esym ~e] (case* ~esym ~tests ~thens ~default)))

      (every? core/keyword? tests)
      (let [tests (->> tests
                       (map #(.substring (core/str %) 1))
                       vec
                       (mapv #(if (seq? %) (vec %) [%])))
            thens (vec (vals pairs))]
        `(let [~esym (if (keyword? ~e) (.-fqn (keyword ~e)) nil)]
           (case* ~esym ~tests ~thens ~default)))

      ;; equality
      :else
      `(let [~esym ~e]
         (cond
           ~@(mapcat (fn [[m c]] `((cljs.core/= ~m ~esym) ~c)) pairs)
           :else ~default)))))

(defmacro assert
  "Evaluates expr and throws an exception if it does not evaluate to
  logical true."
  ([x]
     (when *assert*
       `(when-not ~x
          (throw (js/Error.
                  (cljs.core/str "Assert failed: " ~(pr-str x)))))))
  ([x message]
     (when *assert*
       `(when-not ~x
          (throw (js/Error.
                  (cljs.core/str "Assert failed: " ~message "\n" ~(pr-str x))))))))

(defmacro for
  "List comprehension. Takes a vector of one or more
   binding-form/collection-expr pairs, each followed by zero or more
   modifiers, and yields a lazy sequence of evaluations of expr.
   Collections are iterated in a nested fashion, rightmost fastest,
   and nested coll-exprs can refer to bindings created in prior
   binding-forms.  Supported modifiers are: :let [binding-form expr ...],
   :while test, :when test.

  (take 100 (for [x (range 100000000) y (range 1000000) :while (< y x)]  [x y]))"
  [seq-exprs body-expr]
  (assert-args for
     (vector? seq-exprs) "a vector for its binding"
     (even? (count seq-exprs)) "an even number of forms in binding vector")
  (let [to-groups (fn [seq-exprs]
                    (reduce (fn [groups [k v]]
                              (if (core/keyword? k)
                                (conj (pop groups) (conj (peek groups) [k v]))
                                (conj groups [k v])))
                            [] (partition 2 seq-exprs)))
        err (fn [& msg] (throw (ex-info (apply core/str msg) {})))
        emit-bind (fn emit-bind [[[bind expr & mod-pairs]
                                  & [[_ next-expr] :as next-groups]]]
                    (let [giter (gensym "iter__")
                          gxs (gensym "s__")
                          do-mod (fn do-mod [[[k v :as pair] & etc]]
                                   (core/cond
                                     (= k :let) `(let ~v ~(do-mod etc))
                                     (= k :while) `(when ~v ~(do-mod etc))
                                     (= k :when) `(if ~v
                                                    ~(do-mod etc)
                                                    (recur (rest ~gxs)))
                                     (core/keyword? k) (err "Invalid 'for' keyword " k)
                                     next-groups
                                      `(let [iterys# ~(emit-bind next-groups)
                                             fs# (seq (iterys# ~next-expr))]
                                         (if fs#
                                           (concat fs# (~giter (rest ~gxs)))
                                           (recur (rest ~gxs))))
                                     :else `(cons ~body-expr
                                                  (~giter (rest ~gxs)))))]
                      (if next-groups
                        #_ "not the inner-most loop"
                        `(fn ~giter [~gxs]
                           (lazy-seq
                            (loop [~gxs ~gxs]
                              (when-first [~bind ~gxs]
                                ~(do-mod mod-pairs)))))
                        #_"inner-most loop"
                        (let [gi (gensym "i__")
                              gb (gensym "b__")
                              do-cmod (fn do-cmod [[[k v :as pair] & etc]]
                                        (core/cond
                                          (= k :let) `(let ~v ~(do-cmod etc))
                                          (= k :while) `(when ~v ~(do-cmod etc))
                                          (= k :when) `(if ~v
                                                         ~(do-cmod etc)
                                                         (recur
                                                           (unchecked-inc ~gi)))
                                          (core/keyword? k)
                                            (err "Invalid 'for' keyword " k)
                                          :else
                                            `(do (chunk-append ~gb ~body-expr)
                                                 (recur (unchecked-inc ~gi)))))]
                          `(fn ~giter [~gxs]
                             (lazy-seq
                               (loop [~gxs ~gxs]
                                 (when-let [~gxs (seq ~gxs)]
                                   (if (chunked-seq? ~gxs)
                                     (let [c# ^not-native (chunk-first ~gxs)
                                           size# (count c#)
                                           ~gb (chunk-buffer size#)]
                                       (if (loop [~gi 0]
                                             (if (< ~gi size#)
                                               (let [~bind (-nth c# ~gi)]
                                                 ~(do-cmod mod-pairs))
                                               true))
                                         (chunk-cons
                                           (chunk ~gb)
                                           (~giter (chunk-rest ~gxs)))
                                         (chunk-cons (chunk ~gb) nil)))
                                     (let [~bind (first ~gxs)]
                                       ~(do-mod mod-pairs)))))))))))]
    `(let [iter# ~(emit-bind (to-groups seq-exprs))]
       (iter# ~(second seq-exprs)))))

(defmacro doseq
  "Repeatedly executes body (presumably for side-effects) with
  bindings and filtering as provided by \"for\".  Does not retain
  the head of the sequence. Returns nil."
  [seq-exprs & body]
  (assert-args doseq
     (vector? seq-exprs) "a vector for its binding"
     (even? (count seq-exprs)) "an even number of forms in binding vector")
  (let [err (fn [& msg] (throw (ex-info (apply core/str msg) {})))
        step (fn step [recform exprs]
               (if-not exprs
                 [true `(do ~@body)]
                 (let [k (first exprs)
                       v (second exprs)

                       seqsym (gensym "seq__")
                       recform (if (core/keyword? k) recform `(recur (next ~seqsym) nil 0 0))
                       steppair (step recform (nnext exprs))
                       needrec (steppair 0)
                       subform (steppair 1)]
                   (core/cond
                     (= k :let) [needrec `(let ~v ~subform)]
                     (= k :while) [false `(when ~v
                                            ~subform
                                            ~@(when needrec [recform]))]
                     (= k :when) [false `(if ~v
                                           (do
                                             ~subform
                                             ~@(when needrec [recform]))
                                           ~recform)]
                     (core/keyword? k) (err "Invalid 'doseq' keyword" k)
                     :else (let [chunksym (with-meta (gensym "chunk__")
                                            {:tag 'not-native})
                                 countsym (gensym "count__")
                                 isym     (gensym "i__")
                                 recform-chunk  `(recur ~seqsym ~chunksym ~countsym (unchecked-inc ~isym))
                                 steppair-chunk (step recform-chunk (nnext exprs))
                                 subform-chunk  (steppair-chunk 1)]
                             [true `(loop [~seqsym   (seq ~v)
                                           ~chunksym nil
                                           ~countsym 0
                                           ~isym     0]
                                      (if (coercive-boolean (< ~isym ~countsym))
                                        (let [~k (-nth ~chunksym ~isym)]
                                          ~subform-chunk
                                          ~@(when needrec [recform-chunk]))
                                        (when-let [~seqsym (seq ~seqsym)]
                                          (if (chunked-seq? ~seqsym)
                                            (let [c# (chunk-first ~seqsym)]
                                              (recur (chunk-rest ~seqsym) c#
                                                     (count c#) 0))
                                            (let [~k (first ~seqsym)]
                                              ~subform
                                              ~@(when needrec [recform]))))))])))))]
    (nth (step nil (seq seq-exprs)) 1)))

(defmacro array [& rest]
  (let [xs-str (->> (repeat "~{}")
                    (take (count rest))
                    (interpose ",")
                    (apply core/str))]
    (vary-meta
      (list* 'js* (core/str "[]interface{}{" xs-str "}") rest)
      assoc :tag 'array)))

(defmacro make-array
  ([size] `(make-array nil ~size))
  ([type size]
      (vary-meta
       `(~'js* "make([]interface{}, int(~{}))" ~size)
       assoc :tag 'array)))

(defmacro list
  ([] '(.-EMPTY cljs.core/List))
  ([x & xs]
    `(-conj (list ~@xs) ~x)))

(defmacro vector
  ([] '(.-EMPTY cljs.core/PersistentVector))
  ([& xs]
    (let [cnt (count xs)]
      (if (core/< cnt 32)
        `(cljs.core/PersistentVector. nil ~cnt 5
           (.-EMPTY-NODE cljs.core/PersistentVector) (array ~@xs) nil)
        (vary-meta
          `(.fromArray cljs.core/PersistentVector (array ~@xs) true)
          assoc :tag 'cljs.core/PersistentVector)))))

(defmacro array-map
  ([] '(.-EMPTY cljs.core/PersistentArrayMap))
  ([& kvs]
    (core/cond
      (core/> (count kvs) 16)
      `(hash-map ~@kvs)

      (let [keys (map first (partition 2 kvs))]
        (core/and (every? #(= (:op %) :constant)
                    (map #(cljs.analyzer/analyze &env %) keys))
                  (= (count (into #{} keys)) (count keys))))
      `(cljs.core/PersistentArrayMap. nil ~(clojure.core// (count kvs) 2) (array ~@kvs) nil)

      :else
      `(.fromArray cljs.core/PersistentArrayMap (array ~@kvs) true false))))

(defmacro hash-map
  ([] `(.-EMPTY cljs.core/PersistentHashMap))
  ([& kvs]
    (let [pairs (partition 2 kvs)
          ks    (map first pairs)
          vs    (map second pairs)]
      (vary-meta
        `(.fromArrays cljs.core/PersistentHashMap (array ~@ks) (array ~@vs))
        assoc :tag 'cljs.core/PersistentHashMap))))

(defmacro hash-set
  ([] `(.-EMPTY cljs.core/PersistentHashSet))
  ([& xs]
    (if (core/and (core/<= (count xs) 8)
                  (every? #(= (:op %) :constant)
                    (map #(cljs.analyzer/analyze &env %) xs))
                  (= (count (into #{} xs)) (count xs)))
      `(cljs.core/PersistentHashSet. nil
         (cljs.core/PersistentArrayMap. nil ~(count xs) (array ~@(interleave xs (repeat nil))) nil)
         nil)
      (vary-meta
        `(.fromArray cljs.core/PersistentHashSet (array ~@xs) true)
        assoc :tag 'cljs.core/PersistentHashSet))))

(defn js-obj* [kvs]
  (let [kvs-str (->> (repeat "`~{}`:~{}")
                     (take (count kvs))
                     (interpose ",")
                     (apply core/str))]
    (vary-meta
      (list* 'js* (core/str "map[string]interface{}{" kvs-str "}") (apply concat kvs))
      assoc :tag 'object)))

(defmacro js-obj [& rest]
  (let [sym-or-str? (fn [x] (core/or (core/symbol? x) (core/string? x)))
        filter-on-keys (fn [f coll]
                         (->> coll
                              (filter (fn [[k _]] (f k)))
                              (into {})))
        kvs (into {} (map vec (partition 2 rest)))
        sym-pairs (filter-on-keys core/symbol? kvs)
        expr->local (zipmap
                     (filter (complement sym-or-str?) (keys kvs))
                     (repeatedly gensym))
        obj (gensym "obj")]
    (if (seq rest)
      `(let [~@(apply concat (clojure.set/map-invert expr->local))
             ~obj ~(js-obj* (filter-on-keys core/string? kvs))]
         ~@(map (fn [[k v]] `(aset ~obj ~k ~v)) sym-pairs)
         ~@(map (fn [[k v]] `(aset ~obj ~v ~(core/get kvs k))) expr->local)
         ~obj)
      `(do true)))) ;; hack as transients use {} as edit marker.

(defmacro alength [a]
  (vary-meta
    (core/list 'js* "float64(len(~{}))" a)
    assoc :tag 'number))

(defmacro amap
  "Maps an expression across an array a, using an index named idx, and
  return value named ret, initialized to a clone of a, then setting
  each element of ret to the evaluation of expr, returning the new
  array ret."
  [a idx ret expr]
  `(let [a# ~a
         ~ret (aclone a#)]
     (loop  [~idx 0]
       (if (< ~idx  (alength a#))
         (do
           (aset ~ret ~idx ~expr)
           (recur (inc ~idx)))
         ~ret))))

(defmacro areduce
  "Reduces an expression across an array a, using an index named idx,
  and return value named ret, initialized to init, setting ret to the
  evaluation of expr at each step, returning ret."
  [a idx ret init expr]
  `(let [a# ~a]
     (loop  [~idx 0 ~ret ~init]
       (if (< ~idx  (alength a#))
         (recur (inc ~idx) ~expr)
         ~ret))))

(defmacro dotimes
  "bindings => name n

  Repeatedly executes body (presumably for side-effects) with name
  bound to integers from 0 through n-1."
  [bindings & body]
  (let [i (first bindings)
        n (second bindings)]
    `(let [n# ~n]
       (loop [~i 0]
         (when (< ~i n#)
           ~@body
           (recur (inc ~i)))))))

(defn ^:private check-valid-options
  "Throws an exception if the given option map contains keys not listed
  as valid, else returns nil."
  [options & valid-keys]
  (when (seq (apply disj (apply core/hash-set (keys options)) valid-keys))
    (throw
     (apply core/str "Only these options are valid: "
	    (first valid-keys)
	    (map #(core/str ", " %) (rest valid-keys))))))

(defmacro defmulti
  "Creates a new multimethod with the associated dispatch function.
  The docstring and attribute-map are optional.

  Options are key-value pairs and may be one of:
    :default    the default dispatch value, defaults to :default
    :hierarchy  the isa? hierarchy to use for dispatching
                defaults to the global hierarchy"
  [mm-name & options]
  (let [docstring   (if (core/string? (first options))
                      (first options)
                      nil)
        options     (if (core/string? (first options))
                      (next options)
                      options)
        m           (if (map? (first options))
                      (first options)
                      {})
        options     (if (map? (first options))
                      (next options)
                      options)
        dispatch-fn (first options)
        options     (next options)
        m           (if docstring
                      (assoc m :doc docstring)
                      m)
        m           (if (meta mm-name)
                      (conj (meta mm-name) m)
                      m)
        m (assoc m :tag 'cljs.core/MultiFn)]
    (when (= (count options) 1)
      (throw (Exception. "The syntax for defmulti has changed. Example: (defmulti name dispatch-fn :default dispatch-value)")))
    (let [options   (apply core/hash-map options)
          default   (core/get options :default :default)]
      (check-valid-options options :default :hierarchy)
      `(def ~(with-meta mm-name m)
         (let [method-table# (atom {})
               prefer-table# (atom {})
               method-cache# (atom {})
               cached-hierarchy# (atom {})
               hierarchy# (get ~options :hierarchy (cljs.core/get-global-hierarchy))]
           (cljs.core/MultiFn. ~(name mm-name) ~dispatch-fn ~default hierarchy#
                               method-table# prefer-table# method-cache# cached-hierarchy#))))))

(defmacro defmethod
  "Creates and installs a new method of multimethod associated with dispatch-value. "
  [multifn dispatch-val & fn-tail]
  `(-add-method ~(with-meta multifn {:tag 'cljs.core/MultiFn}) ~dispatch-val (fn ~@fn-tail)))

(defmacro time
  "Evaluates expr and prints the time it took. Returns the value of expr."
  [expr]
  `(let [start# (.getTime (js/Date.))
         ret# ~expr]
     (prn (core/str "Elapsed time: " (- (.getTime (js/Date.)) start#) " msecs"))
     ret#))

(defmacro simple-benchmark
  "Runs expr iterations times in the context of a let expression with
  the given bindings, then prints out the bindings and the expr
  followed by number of iterations and total time. The optional
  argument print-fn, defaulting to println, sets function used to
  print the result. expr's string representation will be produced
  using pr-str in any case."
  [bindings expr iterations & {:keys [print-fn] :or {print-fn 'println}}]
  (let [bs-str   (pr-str bindings)
        expr-str (pr-str expr)]
    `(let ~bindings
       (let [start#   (.getTime (js/Date.))
             ret#     (dotimes [_# ~iterations] ~expr)
             end#     (.getTime (js/Date.))
             elapsed# (- end# start#)]
         (~print-fn (str ~bs-str ", " ~expr-str ", "
                         ~iterations " runs, " elapsed# " msecs"))))))

(def cs (into [] (map (comp gensym core/str core/char) (range 97 118))))

(defmacro gen-apply-to []
  `(do ))

(defmacro with-out-str
  "Evaluates exprs in a context in which *print-fn* is bound to .append
  on a fresh StringBuffer.  Returns the string created by any nested
  printing calls."
  [& body]
  `(let [sb# (goog.string.StringBuffer.)]
     (binding [cljs.core/*print-fn* (fn [x#] (.append sb# x#))]
       ~@body)
     (cljs.core/str sb#)))

(defmacro lazy-cat
  "Expands to code which yields a lazy sequence of the concatenation
  of the supplied colls.  Each coll expr is not evaluated until it is
  needed.

  (lazy-cat xs ys zs) === (concat (lazy-seq xs) (lazy-seq ys) (lazy-seq zs))"
  [& colls]
  `(concat ~@(map #(core/list `lazy-seq %) colls)))

(defmacro js-str [s]
  (core/list 'js* "fmt.Sprint(~{})" s))

(defmacro clojurescript-version []
  (cljs.compiler/clojurescript-to-go-version))
