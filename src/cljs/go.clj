(ns cljs.go
  (:require [cljs.closure]
            [cljs.env]
            [cljs.js-deps]
            [cljs.go.compiler]
            [clojure.string :as s]
            [clojure.java.shell :as sh]))

(defn cljs->go [in]
  (cljs.env/with-compiler-env (cljs.env/default-compiler-env)
    (doall (cljs.closure/-compile in {}))))

(defn go->str [in]
  (if (seq? in)
    (apply str (map go->str in))
    (cljs.js-deps/-source in)))

(defn -main [& args]
  (println "ClojureScript to Go [clojure]"))

(defn go-get [package]
  (let [{:keys [exit err]} (sh/sh "go" "get" package)]
    (assert (zero? exit) err)))

(defn gofmt [in]
  (go-get "code.google.com/p/go.tools/cmd/goimports")
  (let [{:keys [exit out err]} (sh/sh "goimports" :in in)]
    (if (zero? exit)
      out
      (do (println err) in))))

(defn with-line-numbers [in]
  (->> (for [[n l] (map-indexed vector (s/split-lines in))]
         (format "%3d  %s\n" (inc n) l))
       (apply str)
       s/trim-newline))

(comment
  (->> '[(ns test.app (:require [goog.array :as array]))
         (defn plus-one [x] (inc x))]
       cljs->go
       go->str
       gofmt
       with-line-numbers
       println)

;; func plus_one(x interface{}) interface{} {
;; 	return (x.(int) + (1))
;; }


  (->> '[(ns hello)
         (defn world []
           "World")
         (defn ^:export greet [n]
           (str "Hello " n (world)))]
       cljs->go
       go->str
       gofmt
       with-line-numbers
       println)

  (->> '[(ns hello)
         (defn main []
           (println "Hello World"))]
       cljs->go
       go->str
       gofmt
       with-line-numbers
       println)

  (->> '[(ns hello)
         (defn foo
           ([] (foo "World"))
           ([x] (println "Hello " x))
;           ([x & ys] (println "Hello " x ys))
           )]
       cljs->go
       go->str
       gofmt
;       with-line-numbers
       println)

  (cljs.closure/build '[(ns hello.core)
                        (defn ^{:export greet} greet [n] (str "Hola " n))
                        (defn ^:export sum [xs] 42)]
                      {:optimizations :none})

  (cljs.closure/build (str "checkouts/clojurescript/" "samples/hello/src")
                      {:optimizations :none}))
