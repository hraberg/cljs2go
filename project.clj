(defproject cljs2go "0.1.0-SNAPSHOT"
  :description "ClojureScript to Go"
  :url "http://github.com/hraberg/cljs.go"
  :license {:name "Eclipse Public License"
            :url "http://www.eclipse.org/legal/epl-v10.html"}
  :main cljs.go
  :aliases {"compile-clojurescript" ["run" "-m" "cljs.go/compile-clojurescript"]}
  :dependencies [[org.clojure/clojure "1.7.0-alpha3"]
                 [org.clojure/clojurescript "0.0-2371"]
                 [org.clojure/tools.analyzer.js "0.1.0-beta4"]
                 [org.clojure/core.async "0.1.346.0-17112a-alpha"]])
