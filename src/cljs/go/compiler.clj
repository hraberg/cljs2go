;   Copyright (c) Rich Hickey. All rights reserved.
;   The use and distribution terms for this software are covered by the
;   Eclipse Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
;   which can be found in the file epl-v10.html at the root of this distribution.
;   By using this software in any fashion, you are agreeing to be bound by
;   the terms of this license.
;   You must not remove this notice, or any other, from this software.

;; Go emitter. Forked from f0dcc75573a42758f8c39b57d1747a2b4967327e
;; References to js in the public API are retained.

(require 'cljs.compiler)
(in-ns 'cljs.compiler)

;; js* overlays, loaded by core.analyzer at the first run.
(alter-var-root #'ana/*cljs-macros-path* (constantly "/cljs/go/core"))

(def js-reserved
  #{"break" "case" "chan" "const" "continue" "default"
    "defer" "else" "fallthrough" "for" "func" "go" "goto"
    "if" "import" "interface" "map" "package" "range"
    "return" "select" "struct" "switch" "type" "var"})

(defn munge
  ([s] (munge s js-reserved))
  ([s reserved]
    (if (map? s)
      ; Unshadowing
      (let [{:keys [name field] :as info} s
            depth (shadow-depth s)
            renamed (*lexical-renames* (System/identityHashCode s))
            munged-name (munge (cond field (str "self__." name)
                                     renamed renamed
                                     :else name)
                               reserved)]
        (if (or field (zero? depth))
          munged-name
          (symbol (str munged-name "___" depth))))
      ; String munging
      (let [ss (string/replace (str s) #"\/(.)" ".$1") ; Division is special
            ss (apply str (map #(if (reserved %) (str % "_") %)
                               (string/split ss #"(?<=\.)|(?=\.)")))
            ms (string/split (clojure.lang.Compiler/munge ss) #"\.")
            ms (if (butlast ms)
                 (str (string/join "_"(butlast ms)) "." (last ms))
                 (str (last ms)))]
        (if (symbol? s)
          (symbol ms)
          ms)))))

(defmethod emit-constant nil [x] (emits "nil"))
(defmethod emit-constant Character [x]
  (emits (str "'" (escape-char x) "'")))

(defmethod emit-constant java.util.regex.Pattern [x]
  (if (= "" (str x))
    (emits "(js.RegExp{\"\"})")
    (let [[_ flags pattern] (re-find #"^(?:\(\?([idmsux]*)\))?(.*)" (str x))]
      (emits "(js.RegExp{\"" (.replaceAll (re-matcher #"/" pattern) "\\\\/") ", " flags "\"}"))))

(defmethod emit-constant clojure.lang.Keyword [x]
  (if (-> @env/*compiler* :opts :emit-constants)
    (let [value (-> @env/*compiler* ::ana/constant-table x)]
      (emits "cljs_core." value))
    (let [ns   (namespace x)
          name (name x)]
      (emits "cljs_core.Keyword{")
      (emit-constant ns)
      (emits ",")
      (emit-constant name)
      (emits ",")
      (emit-constant (if ns
                       (str ns "/" name)
                       name))
      (emits ",")
      (emit-constant (hash x))
      (emits "}"))))

(defmethod emit-constant clojure.lang.Symbol [x]
  (let [ns     (namespace x)
        name   (name x)
        symstr (if-not (nil? ns)
                 (str ns "/" name)
                 name)]
    (emits "cljs_core.Symbol{")
    (emit-constant ns)
    (emits ",")
    (emit-constant name)
    (emits ",")
    (emit-constant symstr)
    (emits ",")
    (emit-constant (hash x))
    (emits ",")
    (emit-constant nil)
    (emits "}")))

;; tagged literal support

(defmethod emit-constant java.util.Date [^java.util.Date date]
  (emits "js.Date{" (.getTime date) "}"))

(defmethod emit-constant java.util.UUID [^java.util.UUID uuid]
  (emits "cljs_core.UUID{\"" (.toString uuid) "\"}"))

(defmacro emit-wrap [env & body]
  `(let [env# ~env]
     (when (= :return (:context env#)) (emits "return "))
     ~@body
     (when-not (= :expr (:context env#)) (emitln))))

(defmethod emit* :no-op [m])

(defmethod emit* :var
  [{:keys [info env] :as arg}]
  (let [var-name (:name info)]
    ; We need a way to write bindings out to source maps and javascript
    ; without getting wrapped in an emit-wrap calls, otherwise we get
    ; e.g. (function greet(return x, return y) {}).
    (if (:binding-form? arg)
      ; Emit the arg map so shadowing is properly handled when munging
      ; (prevents duplicate fn-param-names)
      (emits (munge arg))
      (when-not (= :statement (:context env))
        (emit-wrap env (emits (munge (if (= ana/*cljs-ns* (:ns info)) (update-in info [:name] name) info))))))))

(defmethod emit* :meta
  [{:keys [expr meta env]}]
  (emit-wrap env
    (emits "cljs_core.with_meta(" expr "," meta ")")))

(defmethod emit* :map
  [{:keys [env keys vals]}]
  (let [simple-keys? (every? #(or (string? %) (keyword? %)) keys)]
    (emit-wrap env
      (cond
        (zero? (count keys))
        (emits "cljs_core.PersistentArrayMap.EMPTY")

        (<= (count keys) array-map-threshold)
        (if (distinct-keys? keys)
          (emits "cljs_core.PersistentArrayMap{nil, " (count keys) ", []interface{}{"
            (comma-sep (interleave keys vals))
            "}, nil}")
          (emits "cljs_core.PersistentArrayMap.fromArray([]interface{}{"
            (comma-sep (interleave keys vals))
            "}, true, false)"))

        :else
        (emits "cljs_core.PersistentHashMap.fromArrays([]interface{}{"
               (comma-sep keys)
               "},[]interface{}{"
               (comma-sep vals)
               "})")))))

(defmethod emit* :list
  [{:keys [items env]}]
  (emit-wrap env
    (if (empty? items)
      (emits "cljs_core.List.EMPTY")
      (emits "cljs_core.list(" (comma-sep items) ")"))))

(defmethod emit* :vector
  [{:keys [items env]}]
  (emit-wrap env
    (if (empty? items)
      (emits "cljs_core.PersistentVector.EMPTY")
      (let [cnt (count items)]
        (if (< cnt 32)
          (emits "cljs_core.PersistentVector{nil, " cnt
            ", 5, cljs_core.PersistentVector.EMPTY_NODE, []interface{}{"  (comma-sep items) "}, nil}")
          (emits "cljs_core.PersistentVector.fromArray([]interface{}{" (comma-sep items) "}, true)"))))))

(defmethod emit* :set
  [{:keys [items env]}]
  (emit-wrap env
    (cond
      (empty? items)
      (emits "cljs_core.PersistentHashSet.EMPTY")

      (distinct-constants? items)
      (emits "cljs_core.PersistentHashSet{nil, cljs_core.PersistentArrayMap{nil, " (count items) ", []interface{}{"
        (comma-sep (interleave items (repeat "nil"))) "}, nil}, nil}")

      :else (emits "cljs_core.PersistentHashSet.fromArray([]interface{}{" (comma-sep items) "}, true)"))))

(defmethod emit* :js-value
  [{:keys [items js-type env]}]
  (emit-wrap env
    (if (= js-type :object)
      (do
        (emits "map[string]interface{}{")
        (when-let [items (seq items)]
          (let [[[k v] & r] items]
            (emits "\"" (name k) "\": " v)
            (doseq [[k v] r]
              (emits ", \"" (name k) "\": " v))))
        (emits "}"))
      (emits "[]interface{}{" (comma-sep items) "}"))))

(defmethod emit* :if
  [{:keys [test then else env unchecked]}]
  (let [context (:context env)
        checked (not (or unchecked (safe-test? env test)))]
    (if (= :expr context)
      (emits "(func() { if " (when checked "cljs_core.truth_") "(" test ") { return " then "} else { return " else "} )()")
      (do
        (if checked
          (emitln "if(cljs_core.truth_(" test "))")
          (emitln "if(" test ")"))
        (emitln "{" then "} else")
        (emitln "{" else "}")))))

(defmethod emit* :case*
  [{:keys [v tests thens default env]}]
  (when (= (:context env) :expr)
    (emitln "(func(){"))
  (let [gs (gensym "caseval__")]
    (when (= :expr (:context env))
      (emitln "var " gs ""))
    (emitln "switch " v " {")
    (doseq [[ts then] (partition 2 (interleave tests thens))]
      (doseq [test ts]
        (emitln "case " test ":"))
      (if (= :expr (:context env))
        (emitln gs "=" then)
        (emitln then))
      (emitln "break"))
    (when default
      (emitln "default:")
      (if (= :expr (:context env))
        (emitln gs "=" default)
        (emitln default)))
    (emitln "}")
    (when (= :expr (:context env))
      (emitln "return " gs "})()"))))

(defmethod emit* :throw
  [{:keys [throw env]}]
  (if (= :expr (:context env))
    (emits "(func(){panic(" throw ")})()")
    (emitln "panic(" throw ")")))

(defmethod emit* :def
  [{:keys [name var init env doc export]}]
  (let [mname (munge name)]
    (when init
      (emit-comment doc (:jsdoc init))
      (if (= :fn (:op init))
        (emits init)
        (do
          (emits "var " (last (string/split (str mname) #"\.")))
          (emitln " = " init)))
      ;; NOTE: JavaScriptCore does not like this under advanced compilation
      ;; this change was primarily for REPL interactions - David
      ;(emits " = (typeof " mname " != 'undefined') ? " mname " : undefined")
      (when-not (= :expr (:context env)) (emitln))

      (when export
        (let [export (last (string/split (str (munge export)) #"\."))]
          (emitln "var "(str (string/upper-case (subs (str (munge export)) 0 1)) (subs (str (munge export)) 1)) " = " (last (string/split (str mname) #"\."))))))))

(defn emit-apply-to
  [{:keys [name params env]}]
  (let [arglist (gensym "arglist__")
        delegate-name (str (munge name) "__delegate")]
    (emitln "(func (" arglist " interface{}) interface{} {")
    (doseq [[i param] (map-indexed vector (drop-last 2 params))]
      (emits "var ")
      (emit param)
      (emits " = cljs_core.first(")
      (emitln arglist ")")
      (emitln arglist " = cljs_core.next(" arglist ")"))
    (if (< 1 (count params))
      (do
        (emits "var ")
        (emit (last (butlast params)))
        (emitln " = cljs_core.first(" arglist ")")
        (emits "var ")
        (emit (last params))
        (emitln " = cljs_core.rest(" arglist ")")
        (emits "return " delegate-name "(")
        (doseq [param params]
          (emit param)
          (when-not (= param (last params)) (emits ",")))
        (emitln ")"))
      (do
        (emits "var ")
        (emit (last params))
        (emitln " = cljs_core.seq(" arglist ")")
        (emits "return " delegate-name "(")
        (doseq [param params]
          (emit param)
          (when-not (= param (last params)) (emits ",")))
        (emitln ")")))
    (emits "})")))

(defn emit-fn-method
  [{:keys [type name variadic params expr env recurs max-fixed-arity]}]
  (emit-wrap env
    (emits "func " (munge name) "(")
    (emit-fn-params params)
    (emits (when (seq params) " interface{}") ") interface{} {")
    (when type
      (emitln "var self__ = this"))
    (when recurs (emitln "for {"))
    (emits expr)
    (when recurs
      (emitln "break")
      (emitln "}"))
    (emits "}")))

(defn emit-variadic-fn-method
  [{:keys [type name variadic params expr env recurs max-fixed-arity] :as f}]
  (emit-wrap env
             (let [name (or name (gensym))
                   mname (munge name)
                   delegate-name (str mname "__delegate")]
               (emitln "(func() { ")
               (emits "var " delegate-name " = func (")
               (doseq [param params]
                 (emit param)
                 (when-not (= param (last params)) (emits ",")))
               (emits " interface{}) interface{} {")
               (when recurs (emitln "for {"))
               (emits expr)
               (when recurs
                 (emitln "break")
                 (emitln "}"))
               (emitln "}")

               (emitln "var " mname " = func (" (comma-sep
                                                 (if variadic
                                                   (butlast params)
                                                   params))
                       (when (seq (if variadic
                                    (butlast params)
                                    params)) " interface{},")
                       (when variadic " arguments ...interface{}") ") interface{} {")
               (when type
                 (emitln "var self__ = this"))
               (when variadic
                 (emits "var ")
                 (emit (last params))
                 (emitln " = nil")
                 (emitln "if (len(arguments) > 0) {")
                 (emits "  ")
                 (emit (last params))
                 (emits " = cljs_core.array_seq(arguments,0)")
                 (emitln "} "))
               (emits "return " delegate-name "(this,")
               (doseq [param params]
                 (emit param)
                 (when-not (= param (last params)) (emits ",")))
               (emits ")")
               (emitln "}")

               (emitln mname ".cljs__lang__maxFixedArity = " max-fixed-arity)
               (emits mname ".cljs__lang__applyTo = ")
               (emit-apply-to (assoc f :name name))
               (emitln)
               (emitln mname ".cljs__core__IFn___invoke__arity__variadic = " delegate-name)
               (emitln "return " mname)
               (emitln "})()"))))

(defmethod emit* :fn
  [{:keys [name env methods max-fixed-arity variadic recur-frames loop-lets]}]
  ;;fn statements get erased, serve no purpose and can pollute scope if named
  (when-not (= :statement (:context env))
    (let [loop-locals (->> (concat (mapcat :params (filter #(and % @(:flag %)) recur-frames))
                                   (mapcat :params loop-lets))
                           (map munge)
                           seq)]
      (when loop-locals
        (when (= :return (:context env))
            (emits "return "))
        (emitln "((func (" (comma-sep (map munge loop-locals)) " interface{}){")
        (when-not (= :return (:context env))
            (emits "return ")))
      (if (= 1 (count methods))
        (if variadic
          (emit-variadic-fn-method (assoc (first methods) :name name))
          (emit-fn-method (assoc (first methods) :name name)))
        (let [has-name? (and name true)
              name (or name (gensym))
              mname (munge name)
              maxparams (apply max-key count (map :params methods))
              mmap (into {}
                     (map (fn [method]
                            [(munge (symbol (str mname "__" (count (:params method)))))
                             method])
                          methods))
              ms (sort-by #(-> % second :params count) (seq mmap))]
          (when (= :return (:context env))
            (emits "return "))
          (emitln "(func() {")
          (emitln "var " mname " = nil")
          (doseq [[n meth] ms]
            (emits "var " n " = ")
            (if (:variadic meth)
              (emit-variadic-fn-method meth)
              (emit-fn-method meth))
            (emitln))
            (emitln mname " = func(" (comma-sep
                                      (if variadic
                                        (butlast maxparams)
                                        maxparams))
                    (when (seq (if variadic
                                 (butlast maxparams)
                                 maxparams)) " interface{},")
                    (when variadic " arguments ...interface{}") ") interface{} {")
          (when variadic
            (emits "var ")
            (emit (last maxparams))
            (emitln " = var_args"))
          (emitln "switch(len(arguments)){")
          (doseq [[n meth] ms]
            (if (:variadic meth)
              (do (emitln "default:")
                  (emitln "return " n ".cljs__core__IFn___invoke__arity__variadic("
                          (comma-sep (butlast maxparams))
                          (when (> (count maxparams) 1) ", ")
                          "cljs_core.array_seq(arguments, " max-fixed-arity "))"))
              (let [pcnt (count (:params meth))]
                (emitln "case " pcnt ":")
                (emitln "return " n "(this" (if (zero? pcnt) nil
                                                (list "," (comma-sep (take pcnt maxparams)))) ")"))))
          (emitln "}")
          (emitln "panic(js.Error{\"Invalid arity: \" + len(arguments)})")
          (emitln "}")
          (when variadic
            (emitln mname ".cljs__lang__maxFixedArity = " max-fixed-arity)
            (emitln mname ".cljs__lang__applyTo = " (some #(let [[n m] %] (when (:variadic m) n)) ms) ".cljs__lang__applyTo"))
          (when has-name?
            (doseq [[n meth] ms]
              (let [c (count (:params meth))]
                (if (:variadic meth)
                  (emitln mname ".cljs__core__IFn___invoke__arity__variadic = " n ".cljs__core__IFn___invoke__arity__variadic")
                  (emitln mname ".cljs__core__IFn___invoke__arity__" c " = " n)))))
          (emitln "return " mname)
          (emitln "})()")))
      (when loop-locals
        (emitln "})(" (comma-sep loop-locals) "))")))))

(defmethod emit* :do
  [{:keys [statements ret env]}]
  (let [context (:context env)]
    (when (and statements (= :expr context)) (emits "(func (){"))
    (when statements
      (emits statements))
    (emit ret)
    (when (and statements (= :expr context)) (emits "})()"))))

(defmethod emit* :try
  [{:keys [env try catch name finally]}]
  (let [context (:context env)]
    (if (or name finally)
      (do
        (emits "(func (){")
        (when name
          (emits "defer func() { if " (munge name) " := recover(); " (munge name) " != nil {" catch "}}()"))
        (when finally
          (assert (not= :constant (:op finally)) "finally block cannot contain constant")
          (emits "defer func() {" finally "}()"))
        (emits "{" try "}")
        (emits "})()"))
      (emits try))))

(defn emit-let
  [{:keys [bindings expr env]} is-loop]
  (let [context (:context env)]
    (when (= :expr context) (emits "(func (){"))
    (binding [*lexical-renames* (into *lexical-renames*
                                      (when (= :statement context)
                                        (map #(vector (System/identityHashCode %)
                                                      (gensym (str (:name %) "-")))
                                             bindings)))]
      (doseq [{:keys [init] :as binding} bindings]
        (emits "var ")
        (emit binding) ; Binding will be treated as a var
        (emitln " = " init))
      (when is-loop (emitln "for {"))
      (emits expr)
      (when is-loop
        (emitln "break")
        (emitln "}")))
    (when (= :expr context) (emits "})()"))))

(defmethod emit* :let [ast]
  (emit-let ast false))

(defmethod emit* :loop [ast]
  (emit-let ast true))

(defmethod emit* :recur
  [{:keys [frame exprs env]}]
  (let [temps (vec (take (count exprs) (repeatedly gensym)))
        params (:params frame)]
    (emitln "{")
    (dotimes [i (count exprs)]
      (emitln "var " (temps i) " = " (exprs i)))
    (dotimes [i (count exprs)]
      (emitln (munge (params i)) " = " (temps i)))
    (emitln "continue")
    (emitln "}")))

(defmethod emit* :letfn
  [{:keys [bindings expr env]}]
  (let [context (:context env)]
    (when (= :expr context) (emits "(func (){"))
    (doseq [{:keys [init] :as binding} bindings]
      (emitln "var " (munge binding) " = " init))
    (emits expr)
    (when (= :expr context) (emits "})()"))))

(defmethod emit* :invoke
  [{:keys [f args env] :as expr}]
  (let [info (:info f)
        fn? (and ana/*cljs-static-fns*
                 (not (:dynamic info))
                 (:fn-var info))
        protocol (:protocol info)
        tag      (ana/infer-tag env (first (:args expr)))
        proto? (and protocol tag
                 (or (and ana/*cljs-static-fns* protocol (= tag 'not-native))
                     (and
                       (or ana/*cljs-static-fns*
                           (:protocol-inline env))
                       (or (= protocol tag)
                           ;; ignore new type hints for now - David
                           (and (not (set? tag))
                                (not ('#{any clj clj-or-nil} tag))
                                (when-let [ps (:protocols (ana/resolve-existing-var (dissoc env :locals) tag))]
                                  (ps protocol)))))))
        opt-not? (and (= (:name info) 'cljs.core/not)
                      (= (ana/infer-tag env (first (:args expr))) 'boolean))
        ns (:ns info)
        js? (= ns 'js)
        goog? (when ns
                (or (= ns 'goog)
                    (when-let [ns-str (str ns)]
                      (= (get (string/split ns-str #"\.") 0 nil) "goog"))))
        keyword? (and (= (-> f :op) :constant)
                      (keyword? (-> f :form)))
        [f variadic-invoke]
        (if fn?
          (let [arity (count args)
                variadic? (:variadic info)
                mps (:method-params info)
                mfa (:max-fixed-arity info)]
            (cond
             ;; if only one method, no renaming needed
             (and (not variadic?)
                  (= (count mps) 1))
             [f nil]

             ;; direct dispatch to variadic case
             (and variadic? (> arity mfa))
             [(update-in f [:info :name]
                             (fn [name] (symbol (str (munge info) ".cljs__core__IFn___invoke__arity__variadic"))))
              {:max-fixed-arity mfa}]

             ;; direct dispatch to specific arity case
             :else
             (let [arities (map count mps)]
               (if (some #{arity} arities)
                 [(update-in f [:info :name]
                             (fn [name] (symbol (str (munge info) ".cljs__core__IFn___invoke__arity__" arity)))) nil]
                 [f nil]))))
          [f nil])]
    (emit-wrap env
      (cond
       opt-not?
       (emits "!(" (first args) ")")
       proto?
       (let [pimpl (str (munge (protocol-prefix protocol))
                        (munge (name (:name info))) "__arity__" (count args))]
         (emits (first args) "." pimpl "(" (comma-sep (cons "nil" (rest args))) ")"))

       keyword?
       (emits f ".cljs__core__IFn___invoke__arity__" (count args) "(" (comma-sep args) ")")

       variadic-invoke
       (let [mfa (:max-fixed-arity variadic-invoke)]
        (emits f "(" (comma-sep (take mfa args))
               (when-not (zero? mfa) ",")
               "cljs_core.array_seq([]interface{}{" (comma-sep (drop mfa args)) "}, 0))"))

       (or fn? js? goog?)
       (emits f "(" (comma-sep args)  ")")

       :else
       (if (and ana/*cljs-static-fns* (= (:op f) :var))
         (let [fprop (str ".cljs__core__IFn___invoke__arity__" (count args))]
           (emits "(func() { if " f fprop " { return " f fprop "(" (comma-sep args) ") } else { return " f "(" (comma-sep (cons "nil" args)) ")}})()"))
         (emits f "(" (comma-sep (cons "nil" args)) ")"))))))

(defmethod emit* :new
  [{:keys [ctor args env]}]
  (emit-wrap env
             (emits "(" ctor "{"
                    (comma-sep args)
                    "})")))

(defmethod emit* :set!
  [{:keys [target val env]}]
  (emit-wrap env (emits target " = " val)))

(defmethod emit* :ns
  [{:keys [name requires uses require-macros env]}]
  (emitln "// " name)
  (emitln "package " (last (string/split (str (munge name)) #"\.")))
  (emitln)
  (emitln "import (")
  (emitln "\t" (wrap-in-double-quotes "js"))
  (when-not (= name 'cljs.core)
    (emitln "\t" "cljs_core" " " (wrap-in-double-quotes "cljs/core")))
  (doseq [lib (distinct (into (vals requires) (vals uses)))]
    (emitln "\t" (string/replace (munge lib) "." "_") " " (wrap-in-double-quotes (string/replace (munge lib) #"[._]" "/"))))
  (emitln ")")
  (emitln))

(defmethod emit* :deftype*
  [{:keys [t fields pmasks]}]
  (let [fields (map munge fields)]
    (emitln "")
    (emitln "/**")
    (emitln "* @constructor")
    (emitln "*/")
    (emitln "type " (munge t) " struct {" (interleave fields (repeat  " interface{};")) "}")
    (emitln (munge t) " = (func (" (comma-sep fields) "  interface{}){")
    (emitln "var this = new(" t ")")
    (doseq [fld fields]
      (emitln "this." fld " = " fld))
    (doseq [[pno pmask] pmasks]
      (emitln "this.cljs__lang__protocol_mask__partition" pno "__ = " pmask))
    (emitln "})")))

(defmethod emit* :defrecord*
  [{:keys [t fields pmasks]}]
  (let [fields (map munge fields)]
    (emitln "")
    (emitln "/**")
    (emitln "* @constructor")
    (doseq [fld fields]
      (emitln "* @param {*} " fld))
    (emitln "* @param {*=} __meta ")
    (emitln "* @param {*=} __extmap")
    (emitln "*/")
    (emitln "type " (munge t) " struct {" (interleave fields (repeat  " interface{};")) "}")
    (emitln (munge t) " = (func (" (comma-sep fields) (when (seq fields)  " interface{},") " arguments ...interface{}) " (munge t) "{")
    (emitln "var this = new(" t ")")
    (doseq [fld fields]
      (emitln "this." fld " = " fld))
    (doseq [[pno pmask] pmasks]
      (emitln "this.cljs__lang__protocol_mask__partition" pno "__ = " pmask))
    (emitln "switch len(arguments) {" )
    (emitln "case 1:")
    (emitln "this.__meta = arguments[0]")
    (emitln "break")
    (emitln "case 2:")
    (emitln "this.__meta = arguments[0]")
    (emitln "this.__extmap = arguments[1]")
    (emitln "}")
    (emitln "}")
    (emitln "})")))

(defn rename-to-js
  "Change the file extension from .cljs to .js. Takes a File or a
  String. Always returns a String."
  [file-str]
  (clojure.string/replace file-str #"\.cljs$" ".go"))

(defn ^File to-target-file
  [target cljs-file]
  (let [relative-path (string/split
                        (ana/munge-path
                          (str (:ns (parse-ns cljs-file)))) #"\.")
        parents (butlast relative-path)]
    (io/file
      (io/file (to-path (cons target parents)))
      (str (last relative-path) ".go"))))


;; TODO: needs fixing, table will include other things than keywords - David

(defn emit-constants-table [table]
  (doseq [[keyword value] table]
    (let [ns   (namespace keyword)
          name (name keyword)]
      (emits "cljs_core." value " = cljs_core.Keyword{")
      (emit-constant ns)
      (emits ",")
      (emit-constant name)
      (emits ",")
      (emit-constant (if ns
                       (str ns "/" name)
                       name))
      (emitln "}"))))
