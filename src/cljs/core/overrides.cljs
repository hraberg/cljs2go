;; Go overrides
(ns cljs.core)

(def ^:dynamic *print-length* -1)
(def ^:dynamic *print-level* -1)

(def ^function set-print-fn!)
(def ^function symbol?)
(def ^function complement)
(def ^function iter)
(def ^function pr-writer)
(def ^function pr-sequential-writer)
(def ^function type_)
(def ^function type->str)
(def ^function add-string-to-hash-cache)
(def ^function hash-string)
(def ^function interger?)
(def ^function array)
(def ^function nil-iter)

(defn enable-console-print!
  "Set *print-fn* to console.log"
  []
  (set! *print-fn*
        (fn fmt-println [x]
          (js* "fmt.Println(~{})" x)
          nil)))

(defn apply
  "Applies fn f to the argument list formed by prepending intervening arguments to args.
  First cut.  Not lazy.  Needs to use emitted toApply."
  ([f args]
     (js* "f.(*AFn).Call(~{}.([]interface{})...)" (into-array args)))
  ([f x args]
     (js* "f.(*AFn).Call(append([]interface{}{~{}}, ~{}.([]interface{})...)...)" x (into-array args)))
  ([f x y args]
     (js* "f.(*AFn).Call(append([]interface{}{~{}, ~{}}, ~{}.([]interface{})...)...)" x y (into-array args)))
  ([f x y z args]
     (js* "f.(*AFn).Call(append([]interface{}{~{}, ~{}, ~{}}, ~{}.([]interface{})...)...)" x y z (into-array args)))
  ([f a b c d & args]
     (let [arr (into-array (butlast args))
           varargs (into-array (last args))]
       (js* "f.(*AFn).Call(append([]interface{}{~{}, ~{}, ~{}, ~{}}, append(~{}, ~{}...))...)" a b c d arr varargs))))

(defn ^boolean native-satisfies?
  "Internal - do not use!"
  [p x]
  ^boolean  (js* "value(decorate(~{})).Type().Implements(protocols[fmt.Sprint(~{})])" p x))

;; Simple caching of string hashcode
(def string-hash-cache (js* "map[string]interface{}{}"))
