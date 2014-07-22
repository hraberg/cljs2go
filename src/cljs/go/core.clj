;   Copyright (c) Rich Hickey. All rights reserved.
;   The use and distribution terms for this software are covered by the
;   Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
;   which can be found in the file epl-v10.html at the root of this distribution.
;   By using this software in any fashion, you are agreeing to be bound by
;   the terms of this license.
;   You must not remove this notice, or any other, from this software.

;; Go emitter. Forked from 5ca535759f8e2490b42dfa25df0458a6c3376d8c
;; This file only overlays cljs.core macros that use js*
;; Functions which use the internal representation of IFn, related to get-apply-to and protocols probabaly also needs to be here.
;; Some of this leaks out into cljs/core.cljs

(require 'cljs.core)
(in-ns 'cljs.core)

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
                  (map #(if (core/or (core/string? %) (core/integer? %)
                                     (core/float? %) (core/true? %)
                                     (core/false? %))
                          "~{}"
                          "cljs.core.str.cljs$core$IFn$_invoke$arity$1(~{})"))
                  (interpose "+")
                  (apply core/str))]
    ;; Google closure advanced compile will stringify and concat strings and
    ;; numbers at compilation time.
    (list* 'js* (core/str (if (core/string? (first clean-xs)) "(" "(''+")
                          strs ")")
           clean-xs)))

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

;; internal - do not use.
(defmacro coercive-not [x]
  (bool-expr (core/list 'js* "(!~{})" x)))

;; internal - do not use.
(defmacro coercive-not= [x y]
  (bool-expr (core/list 'js* "(~{} != ~{})" x y)))

;; internal - do not use.
(defmacro coercive-= [x y]
  (bool-expr (core/list 'js* "(~{} == ~{})" x y)))

;; internal - do not use.
(defmacro coercive-boolean [x]
  (with-meta (core/list 'js* "~{}" x)
    {:tag 'boolean}))

;; internal - do not use.
(defmacro truth_ [x]
  (core/assert (clojure.core/symbol? x) "x is substituted twice")
  (core/list 'js* "(~{} != nil && ~{} != false)" x x))

(defmacro js-delete [obj key]
  (core/list 'js* "delete(~{}, ~{})" obj key))

(defmacro true? [x]
  (bool-expr (core/list 'js* "~{} == true" x)))

(defmacro false? [x]
  (bool-expr (core/list 'js* "~{} == false" x)))

(defmacro array? [x]
  (if (= :nodejs (:target @env/*compiler*))
    (bool-expr `(.isArray js/Array ~x))
    (bool-expr (core/list 'js* "~{}.(type) == \"interface{}\"" x))))

(defmacro string? [x]
  (bool-expr (core/list 'js* "~{}.(type) == \"string\"" x)))

;; TODO: x must be a symbol, not an arbitrary expression
(defmacro exists? [x]
  (bool-expr
   (core/list 'js* "~{} != nil"
              (vary-meta x assoc :cljs.analyzer/no-resolve true))))

(defmacro undefined? [x]
  (bool-expr (core/list 'js* "(nil == ~{})" x)))

(defmacro identical? [a b]
  (bool-expr (core/list 'js* "(&~{} == &~{})" a b)))

(defmacro instance? [t o]
  ;; Google Closure warns about some references to RegExp, so
  ;; (instance? RegExp ...) needs to be inlined, but the expansion
  ;; should preserve the order of argument evaluation.
  (bool-expr (if (clojure.core/symbol? t)
               (core/list 'js* "(~{}.(type) == ~{})" o t)
               `(let [t# ~t o# ~o]
                  (~'js* "(~{}.(type) == ~{})" o# t#)))))

(defmacro number? [x]
  (bool-expr (core/list 'js* "~{}.(type) == 'float64'" x)))

(defmacro aget
  ([a i]
     (core/list 'js* "(~{}[~{}])" a i))
  ([a i & idxs]
     (let [astr (apply core/str (repeat (count idxs) "[~{}]"))]
       `(~'js* ~(core/str "(~{}[~{}]" astr ")") ~a ~i ~@idxs))))

(defmacro aset
  ([a i v]
     (core/list 'js* "(~{}[~{}] = ~{})" a i v))
  ([a idx idx2 & idxv]
     (let [n    (core/dec (count idxv))
           astr (apply core/str (repeat n "[~{}]"))]
       `(~'js* ~(core/str "(~{}[~{}][~{}]" astr " = ~{})") ~a ~idx ~idx2 ~@idxv))))

(defmacro byte [x] (core/list 'js* "(~{}).(byte)" x))
(defmacro short [x] (core/list 'js* "(~{}).(int16)" x))
(defmacro float [x] (core/list 'js* "(~{}).(float32)" x))
(defmacro double [x]
  (if (core/number? x)
    x
    (core/list 'js* "(~{}).(float64)" x)))

(defmacro unchecked-byte [x] `(byte ~x))
(defmacro unchecked-char [x] (core/list 'js* "(~{}).(uint16)" x))
(defmacro unchecked-short [x] `(short ~x))
(defmacro unchecked-float [x] `(float ~x))
(defmacro unchecked-double [x] `(double ~x))

(defmacro ^::ana/numeric +
  ([] 0)
  ([x] `(double ~x))
  ([x y] (core/list 'js* "(~{} + ~{})" `(double ~x) `(double ~y)))
  ([x y & more] `(+ (+ ~x ~y) ~@more)))

(defmacro ^::ana/numeric -
  ([x] (core/list 'js* "(- ~{})" `(double ~x)))
  ([x y] (core/list 'js* "(~{} - ~{})" `(double ~x) `(double ~y)))
  ([x y & more] `(- (- ~x ~y) ~@more)))

(defmacro ^::ana/numeric *
  ([] 1)
  ([x] `(double ~x))
  ([x y] (core/list 'js* "(~{} * ~{})" `(double ~x) `(double ~y)))
  ([x y & more] `(* (* ~x ~y) ~@more)))

(defmacro ^::ana/numeric /
  ([x] `(/ 1 ~x))
  ([x y] (core/list 'js* "(~{} / ~{})" `(double ~x) `(double ~y)))
  ([x y & more] `(/ (/ ~x ~y) ~@more)))

(defmacro ^::ana/numeric divide
  ([x] `(/ 1 ~x))
  ([x y] (core/list 'js* "(~{} / ~{})" `(double ~x) `(double ~y)))
  ([x y & more] `(/ (/ ~x ~y) ~@more)))

(defmacro ^::ana/numeric <
  ([x] true)
  ([x y] (bool-expr (core/list 'js* "(~{} < ~{})" `(double ~x) `(double ~y))))
  ([x y & more] `(and (< ~x ~y) (< ~y ~@more))))

(defmacro ^::ana/numeric <=
  ([x] true)
  ([x y] (bool-expr (core/list 'js* "(~{} <= ~{})" `(double ~x) `(double ~y))))
  ([x y & more] `(and (<= ~x ~y) (<= ~y ~@more))))

(defmacro ^::ana/numeric >
  ([x] true)
  ([x y] (bool-expr (core/list 'js* "(~{} > ~{})" `(double ~x) `(double ~y))))
  ([x y & more] `(and (> ~x ~y) (> ~y ~@more))))

(defmacro ^::ana/numeric >=
  ([x] true)
  ([x y] (bool-expr (core/list 'js* "(~{} >= ~{})" `(double ~x) `(double ~y))))
  ([x y & more] `(and (>= ~x ~y) (>= ~y ~@more))))

(defmacro ^::ana/numeric ==
  ([x] true)
  ([x y] (bool-expr (core/list 'js* "(~{} == ~{})" `(double ~x) `(double ~y))))
  ([x y & more] `(and (== ~x ~y) (== ~y ~@more))))

(defmacro ^::ana/numeric max
  ([x] x)
  ([x y] `(let [x# ~x, y# ~y]
            (if (> x# y#) x# y#)))
  ([x y & more] `(max (max ~x ~y) ~@more)))

(defmacro ^::ana/numeric min
  ([x] x)
  ([x y] `(let [x# ~x, y# ~y]
            (if (< x# y#) x# y#)))
  ([x y & more] `(min (min ~x ~y) ~@more)))

(defmacro ^::ana/numeric js-mod [num div]
  (core/list 'js* "(~{} % ~{})" `(double ~num) `(double ~div)))

(defmacro ^::ana/numeric bit-not [x]
  (core/list 'js* "(~ ~{})" `(int ~x)))

(defmacro ^::ana/numeric bit-and
  ([x y] (core/list 'js* "(~{} & ~{})" x y))
  ([x y & more] `(bit-and (bit-and ~x ~y) ~@more)))

;; internal do not use
(defmacro ^::ana/numeric unsafe-bit-and
  ([x y] (bool-expr (core/list 'js* "(~{} & ~{})" x y)))
  ([x y & more] `(unsafe-bit-and (unsafe-bit-and ~x ~y) ~@more)))

(defmacro ^::ana/numeric bit-or
  ([x y] (core/list 'js* "(~{} | ~{})" x y))
  ([x y & more] `(bit-or (bit-or ~x ~y) ~@more)))

(defmacro ^::ana/numeric int [x]
  `(bit-or ~x 0))

(defmacro ^::ana/numeric bit-xor
  ([x y] (core/list 'js* "(~{} ^ ~{})" x y))
  ([x y & more] `(bit-xor (bit-xor ~x ~y) ~@more)))

(defmacro ^::ana/numeric bit-and-not
  ([x y] (core/list 'js* "(~{} & ~~{})" x y))
  ([x y & more] `(bit-and-not (bit-and-not ~x ~y) ~@more)))

(defmacro ^::ana/numeric bit-clear [x n]
  (core/list 'js* "(~{} & ~(1 << ~{}))" x n))

(defmacro ^::ana/numeric bit-flip [x n]
  (core/list 'js* "(~{} ^ (1 << ~{}))" x n))

(defmacro ^::ana/numeric bit-test [x n]
  (core/list 'js* "((~{} & (1 << ~{})) != 0)" x n))

(defmacro ^::ana/numeric bit-shift-left [x n]
  (core/list 'js* "(~{} << ~{})" x n))

(defmacro ^::ana/numeric bit-shift-right [x n]
  (core/list 'js* "(~{} >> ~{})" x n))

(defmacro ^::ana/numeric bit-shift-right-zero-fill [x n]
  (core/list 'js* "(~{} >>> ~{})" x n))

(defmacro ^::ana/numeric unsigned-bit-shift-right [x n]
  (core/list 'js* "(~{} >>> ~{})" x n))

(defmacro ^::ana/numeric bit-set [x n]
  (core/list 'js* "(~{} | (1 << ~{}))" x n))

;; internal
(defmacro mask [hash shift]
  (core/list 'js* "((~{} >>> ~{}) & 0x01f)" hash shift))

;; internal
(defmacro bitpos [hash shift]
  (core/list 'js* "(1 << ~{})" `(mask ~hash ~shift)))

(def #^:private base-type
  {nil "nil"
   'object "interface{}"
   'string "string"
   'number "float64"
   'array "[]interface{}"
   'function "func"
   'boolean "bool"
   'default "_"})

(def #^:private js-base-type
  {'js/Boolean "bool"
   'js/String "string"
   'js/Array "[]interface{}"
   'js/Object "interface{}"
   'js/Number "float64"
   'js/Function "func"})

(defn js-obj* [kvs]
  (let [kvs-str (->> (repeat "\"~{}\":~{}")
                     (take (count kvs))
                     (interpose ",")
                     (apply core/str))]
    (vary-meta
     (list* 'js* (core/str "map[string]interface{}{" kvs-str "}") (apply concat kvs))
     assoc :tag 'object)))

(defmacro alength [a]
  (vary-meta
   (core/list 'js* "len(~{})" a)
   assoc :tag 'number))
