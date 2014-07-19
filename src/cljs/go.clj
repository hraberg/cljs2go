(ns cljs.go
  (:require [cljs.closure]
            [cljs.env]
            [cljs.js-deps]))

(defn cljs->go [in]
  (cljs.env/with-compiler-env (cljs.env/default-compiler-env)
    (doall (cljs.closure/-compile in {}))))

(defn go->str [in]
  (if (seq? in)
    (apply str (map go->str in))
    (cljs.js-deps/-source in)))

(defn -main [& args]
  (println "ClojureScript to Go [clojure]"))

(comment
  (println (cljs->go
            '[(ns test.app (:require [goog.array :as array]))
              (defn plus-one [x] (inc x))]))

  (println (cljs->go
            (str "checkouts/clojurescript/" "samples/hello/src/hello/core.cljs")))

  (println (cljs->go
            (str "checkouts/clojurescript/" "samples/hello/src"))))
