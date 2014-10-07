(ns ^{:doc "Go overrides."}
  clojure.string
  (:refer-clojure :exclude [replace reverse])
  (:require [goog.string :as gstring]))

(defn replace
  "Replaces all instance of match with replacement in s.
   match/replacement can be:

   string / string
   pattern / (string or function of match)."
  [s match replacement]
  (cond (string? match)
        (.replace s (js/RegExp. (gstring/regExpEscape match) "g") replacement)
        (instance? js/RegExp match)
        (.replace s (js/RegExp. (.-pattern match) (str (.-flags match) "g"))
                  (if (fn? replacement)
                    (js* "func(x interface{}) interface{} { return ~{}.(cljs_core.CljsCoreIFn).X_invoke_Arity1(x) }"
                         replacement)
                    replacement))
        :else (throw (str "Invalid match arg: " match))))

(defn replace-first
  "Replaces the first instance of match with replacement in s.
   match/replacement can be:

   string / string
   pattern / (string or function of match)."
  [s match replacement]
  (.replace s match (if (fn? replacement)
                      (js* "func(x interface{}) interface{} { return ~{}.(cljs_core.CljsCoreIFn).X_invoke_Arity1(x) }"
                           replacement)
                      replacement)))
