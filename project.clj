(defproject cljs2go "0.1.0-SNAPSHOT"
  :description "ClojureScript to Go"
  :url "http://github.com/hraberg/cljs.go"
  :license {:name "Eclipse Public License"
            :url "http://www.eclipse.org/legal/epl-v10.html"}
  :main cljs.go
  :aliases {"compile-clojurescript" ["run" "-m" "cljs.go/compile-clojurescript"]
            "compile-clojurescript-tests" ["run" "-m" "cljs.go-test/clojurescript-tests"]}
  :dependencies [[org.clojure/clojure "1.7.0-alpha4"]
                 [org.clojure/clojurescript "0.0-2411"]
                 [org.clojure/tools.analyzer.js "0.1.0-beta5"]
                 [org.clojure/core.async "0.1.346.0-17112a-alpha"]])
