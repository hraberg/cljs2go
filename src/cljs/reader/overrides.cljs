(ns ^{:doc "Go overrides."}
  cljs.reader
  (:require [goog.string :as gstring]))

(defn read-2-chars [reader]
  (str
   (read-char reader)
   (read-char reader)))

(defn read-4-chars [reader]
  (str
   (read-char reader)
   (read-char reader)
   (read-char reader)
   (read-char reader)))

(defn read-token
  [rdr initch]
  (loop [sb (gstring/StringBuffer. initch)
         ch (read-char rdr)]
    (if (or (nil? ch)
            (whitespace? ch)
            (macro-terminating? ch))
      (do (unread rdr ch) (.toString sb))
      (recur (do (.append sb ch) sb) (read-char rdr)))))

(defn macros [c]
  (cond
   (identical? c \") read-string*
   (identical? c \:) read-keyword
   (identical? c \;) read-comment
   (identical? c \') (wrapping-reader 'quote)
   (identical? c \@) (wrapping-reader 'deref)
   (identical? c \^) read-meta
   (identical? c \`) not-implemented
   (identical? c \~) not-implemented
   (identical? c \() read-list
   (identical? c \)) read-unmatched-delimiter
   (identical? c \[) read-vector
   (identical? c \]) read-unmatched-delimiter
   (identical? c \{) read-map
   (identical? c \}) read-unmatched-delimiter
   (identical? c \\) (fn [rdr _] (read-char rdr))
   (identical? c \#) read-dispatch
   :else nil))

(def ^:private days-in-month
  (let [dim-norm [nil 31 28 31 30 31 30 31 31 30 31 30 31]
        dim-leap [nil 31 29 31 30 31 30 31 31 30 31 30 31]]
    (identity (fn [month leap-year?]
                (get (if leap-year? dim-leap dim-norm) month)))))

(defn parse-timestamp
  [ts]
  (if-let [[years months days hours minutes seconds ms offset]
           (parse-and-validate-timestamp ts)]
    (js/Date. ts)
    (reader-error nil (str "Unrecognized date/time syntax: " ts))))

(defn ^:private read-queue
  [elems]
  (if (vector? elems)
    (into (.-EMPTY cljs.core/PersistentQueue) elems)
    (reader-error nil "Queue literal expects a vector for its elements.")))

(defn ^:private read-date
  [s]
  (if (string? s)
    (parse-timestamp s)
    (reader-error nil "Instance literal expects a string for its timestamp.")))

(defn ^:private read-uuid
  [uuid]
  (if (string? uuid)
    (UUID. uuid)
    (reader-error nil "UUID literal expects a string as its representation.")))

(def ^:dynamic *tag-table*
  (atom {"inst"  read-date
         "uuid"  read-uuid
         "queue" read-queue}))
