(ns ^{:doc "Go overrides."}
  cljs.core
  (:require [goog.array :as garray]))

(def *clojurescript-version* (clojurescript-version))

(def ^:dynamic *print-length* -1)
(def ^:dynamic *print-level* -1)

(defn set-print-fn!
  "Set *print-fn* to f."
  [f] (set! *print-fn* (js* "~{}.(*AFn)" f)))

(defn ^boolean symbol? [x]
  (cljs.core/instance? Symbol x))

(defn complement
  "Takes a fn f and returns a fn that takes the same arguments as f,
  has the same effects, if any, and returns the opposite truth value."
  [f]
  (fn complement-fn
    ([] (not (f)))
    ([x] (not (f x)))
    ([x y] (not (f x y)))
    ([x y & zs] (not (apply f x y zs)))))

(defn ^boolean =
  "Equality. Returns true if x equals y, false if not. Compares
  numbers and collections in a type-independent manner.  Clojure's immutable data
  structures define -equiv (and thus =) as a value, not an identity,
  comparison."
  ([x] true)
  ([x y]
    (if (nil? x)
      (nil? y)
      (or (identical? x y)
          (if (implements? IEquiv x)
            ^boolean (-equiv x y)
            false))))
  ([x y & more]
     (if (= x y)
       (if (next more)
         (recur y (first more) (next more))
         (= y (first more)))
       false)))

(defn sort
  "Returns a sorted sequence of the items in coll. Comp can be
   boolean-valued comparison funcion, or a -/0/+ valued comparator.
   Comp defaults to compare."
  ([coll]
   (sort compare coll))
  ([comp coll]
   (if (seq coll)
     (let [a (to-array coll)
           comp ^function (fn->comparator comp)]
       ;; matching Clojure's stable sort, though docs don't promise it
       (garray/stableSort a (js* "func(x, y interface{}) interface{} { return ~{}.X_invoke_Arity2(x, y) }" comp))
       (seq a))
     ())))

(defn get
  "Returns the value mapped to key, not-found or nil if key not present."
  ([o k]
    (when-not (nil? o)
      (cond
        (implements? ILookup o)
        (-lookup ^not-native o k)

        (array? o)
        (when (and (number? k) (< k (.-length o)))
          (aget ^array (js* "~{}.([]interface{})" o) k))

        (string? o)
        (when (and (number? k) (< k (.-length o)))
          (js* "string(~{})" (aget ^string (js* "~{}.(string)" o) k)))

        (native-satisfies? ILookup o)
        (-lookup o k)

        :else nil)))
  ([o k not-found]
    (if-not (nil? o)
      (cond
        (implements? ILookup o)
        (-lookup ^not-native o k not-found)

        (array? o)
        (if (and (number? k) (< k (.-length o)))
          (aget ^array (js* "~{}.([]interface{})" o) k)
          not-found)

        (string? o)
        (if (and (number? k) (< k (.-length o)))
          (js* "string(~{})" (aget ^string (js* "~{}.(string)" o) k))
          not-found)

        (native-satisfies? ILookup o)
        (-lookup o k not-found)

        :else not-found)
      not-found)))

(defn ^:private quote-string
  [s]
  (str \"
       (.replace s (js/RegExp. "[\\\\\"\b\f\n\r\t]" "")
         (js* "func(match interface{}) interface{} { return ~{}.(map[string]interface{})[match.(string)] }" char-escapes))
       \"))

(defn- pr-writer
  "Prefer this to pr-seq, because it makes the printing function
   configurable, allowing efficient implementations such as appending
   to a StringBuffer."
  [obj writer opts]
  (cond
    (nil? obj) (-write writer "nil")
    :else (do
            (when (and (get opts :meta)
                       (satisfies? IMeta obj)
                       (meta obj))
              (-write writer "^")
              (pr-writer (meta obj) writer opts)
              (-write writer " "))
            (cond
              (nil? obj) (-write writer "nil")

              ;; handle CLJS ctors
              ;; ^boolean (.-cljs$lang$type obj)
              ;; (.cljs$lang$ctorPrWriter obj obj writer opts)

              ; Use the new, more efficient, IPrintWithWriter interface when possible.
              (implements? IPrintWithWriter obj)
              (-pr-writer ^not-native obj writer opts)

              (or ^boolean (js* "reflect.ValueOf(~{}).Kind() == reflect.Bool" obj) (number? obj))
              (-write writer (str obj))

              (array? obj)
              (pr-sequential-writer writer pr-writer "#js [" " " "]" opts obj)

              ^boolean (goog/isString obj)
              (if (:readably opts)
                (-write writer (quote-string obj))
                (-write writer obj))

              (fn? obj)
              (write-all writer "#<" (str obj) ">")

              (instance? js/Date obj)
              (let [normalize (fn [n len]
                                (loop [ns (str n)]
                                  (if (< (count ns) len)
                                    (recur (str "0" ns))
                                    ns)))]
                (write-all writer
                  "#inst \""
                  (str (.getUTCFullYear obj))             "-"
                  (normalize (inc (.getUTCMonth obj)) 2)  "-"
                  (normalize (.getUTCDate obj) 2)         "T"
                  (normalize (.getUTCHours obj) 2)        ":"
                  (normalize (.getUTCMinutes obj) 2)      ":"
                  (normalize (.getUTCSeconds obj) 2)      "."
                  (normalize (.getUTCMilliseconds obj) 3) "-"
                  "00:00\""))

              (regexp? obj) (write-all writer "#\"" (.-source obj) "\"")

              (satisfies? IPrintWithWriter obj)
              (-pr-writer obj writer opts)

              :else (write-all writer "#<" (str obj) ">")))))

(defn pr-sequential-writer [writer print-one begin sep end opts coll]
  (binding [*print-level* (if (== -1 *print-level*) -1 (dec *print-level*))]
    (if (and (not (nil? *print-level*)) (neg? *print-level*))
      (-write writer "#")
      (do
        (-write writer begin)
        (when (seq coll)
          (print-one (first coll) writer opts))
        (loop [coll (next coll) n (dec (:print-length opts))]
          (if (and coll (or (nil? n) (not (zero? n))))
            (do
              (-write writer sep)
              (print-one (first coll) writer opts)
              (recur (next coll) (dec n)))
            (when (and (seq coll) (zero? n))
              (-write writer sep)
              (-write writer "..."))))
        (-write writer end)))))

(defn type [x]
  (when-not (nil? x)
    (js* "reflect.TypeOf(~{})" x)))

(defn type->str [ty]
  (str ty))

(defn ^boolean integer?
  "Returns true if n is an integer."
  [n]
  (and (number? n)
       (not ^boolean (js/isNaN n))
       (not (identical? n js/Infinity))
       (== n ^number (js* "float64(int(~{}.(float64)))" n))))

(defn ^array array
  "Creates a new javascript array.
@param {...*} var_args" ;;array is a special case, don't emulate this doc string
  [& items]
  (into-array items))

(defn ^array make-array
  ([size]
     (make-array nil size))
  ([type size]
     (cljs.core/make-array size)))

(defn char
  "Coerce to char"
  [x]
  (cond
    (number? x) (js* "js.String.FromCharCode(~{}.(float64))" x)
    (and (string? x) (== (.-length x) 1)) x
    :else (throw (js/Error. "Argument to char must be a character or number"))))

;; Simple caching of string hashcode
(def string-hash-cache ^"map[interface{}]interface{}" (js* "map[interface{}]interface{}{}"))

(defn add-to-string-hash-cache [k]
  (let [h (hash-string* k)]
    (js* "~{}[~{}] = ~{}" string-hash-cache k h)
    (set! string-hash-cache-count (inc string-hash-cache-count))
    h))

(defn hash-string [k]
  (when (> string-hash-cache-count 255)
    (set! string-hash-cache (js* "map[interface{}]interface{}{}"))
    (set! string-hash-cache-count 0))
  (let [h (js* "~{}[~{}]" string-hash-cache k)]
    (if (number? h)
      h
      (add-to-string-hash-cache k))))

(defn enable-console-print!
  "Set *print-fn* to console.log"
  []
  (set! *print-newline* false)
  (set! *print-fn*
        (fn fmt-println [x]
          (js* "fmt.Println(~{})" x)
          nil)))

(defn apply
  "Applies fn f to the argument list formed by prepending intervening arguments to args.
  First cut.  Not lazy.  Needs to use emitted toApply."
  ([f args]
     (js* "Call_(~{}.(CljsCoreIFn), ~{}...)" f (into-array args)))
  ([f x args]
     (js* "Call_(~{}.(CljsCoreIFn), append([]interface{}{~{}}, ~{}...)...)" f x (into-array args)))
  ([f x y args]
     (js* "Call_(~{}.(CljsCoreIFn), append([]interface{}{~{}, ~{}}, ~{}...)...)" f x y (into-array args)))
  ([f x y z args]
     (js* "Call_(~{}.(CljsCoreIFn), append([]interface{}{~{}, ~{}, ~{}}, ~{}...)...)" f x y z (into-array args)))
  ([f a b c d & args]
     (let [arr (into-array (butlast args))
           varargs (into-array (last args))]
       (js* "Call_(~{}.(CljsCoreIFn), append([]interface{}{~{}, ~{}, ~{}, ~{}}, append(~{}, ~{}...)...)...)" f a b c d arr varargs))))

(defn ^boolean native-satisfies?
  "Internal - do not use!"
  [p x]
  ^boolean (js* "value(decorate(~{})).Type().Implements(protocols[~{}])" x (str p)))

;; There are two protocols in clojure.data, EqualityPartition and Diff which extend the base types and default which gets skipped.
;; There are also the extent-type calls at start of core.cljs we aren't dealing with.
