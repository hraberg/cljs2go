(ns cljs.go
  (:require [cljs.closure]
            [cljs.env]
            [cljs.js-deps]
            [cljs.go.compiler]))

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
  (->> '[(ns test.app (:require [goog.array :as array]))
         (defn plus-one [x] (inc x))]
       cljs->go
       go->str
       println)

;; func plus_one(x interface{}) interface{} {
;; 	return (x.(int) + (1))
;; }


  (->> '[(ns hello)
         (defn world [world]
           "World")
         (defn ^:export greet [n]
           (str "Hello " n (world)))]
       cljs->go
       go->str
       println)

  (cljs.closure/build '[(ns hello.core)
                        (defn ^{:export greet} greet [n] (str "Hola " n))
                        (defn ^:export sum [xs] 42)]
                      {:optimizations :none})

  (cljs.closure/build (str "checkouts/clojurescript/" "samples/hello/src")
                      {:optimizations :none}))
