(ns ^{:doc "Go overrides."}
  clojure.data)


(defn ep [x]
  (cond
   (nil? x) :atom
   (string? x) :atom
   (number? x) :atom
   (array? x) :sequential
   (fn? x) :atom
   (boolean? x) :atom
   (satisfies? IMap x) :map
   (satisfies? ISet x) :set
   (satisfies? ISequential x) :sequential
   :default :atom))

(defn ds [a b]
  ((case (ep a)
     :atom atom-diff
     :set diff-set
     :sequential diff-sequential
     :map diff-associative)
   a b))

(defn diff
  "Recursively compares a and b, returning a tuple of
  [things-only-in-a things-only-in-b things-in-both].
  Comparison rules:

  * For equal a and b, return [nil nil a].
  * Maps are subdiffed where keys match and values differ.
  * Sets are never subdiffed.
  * All sequential things are treated as associative collections
    by their indexes, with results returned as vectors.
  * Everything else (including strings!) is treated as
    an atom and compared for equality."
  [a b]
  (if (= a b)
    [nil nil a]
    (if (= (ep a) (ep b))
      (ds a b)
      (atom-diff a b))))
