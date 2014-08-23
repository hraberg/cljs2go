# cljs.go

*Way to Go!* -- [Siobhan Sharpe](http://londonist.com/2012/07/did-boris-write-twenty-twelve-way-to-go-gag.php)


ClojureScript to Go. Implemented as an overlay onto ClojureScript, instead of a fork. Go is emitted from `cljs.go.compiler`, which is a patched version of `cljs.compiler`. It also does some AST rewriting to deal with JS corner cases. At run time, a thin JS compatibility is provided to avoid having to touch `cljs.core`, these `js` and `goog` packages aren't intended for end-user usage.

Once the compiler starts working, the plan is to provide a second monkey patch in spirit of the first, but this time for the Java dependencies, like `java.io.File` and their Clojure counterpart, `clojure.java.io`. This patch is to be able to compile the compiler itself to Go. This won't change the run-time characteristics of the compiled programs, nor will it introduce `eval`, it's simply done to leverage the Go build pipeline. ClojureScript has a reader in `cljs.reader` that potentially can replace `tools.reader`. Macros used via `:require-macros` needs to be compiled to Go first, and an ad-hoc wrapper for the compiler depending on them will be created to compile the actual source (the basic compiler has the `cljs.core` macros built in).

While not introducing `eval`, it does open up for adding it as a third layer, the problem lies on the Go side, as you need a way to build and link code into the running process. This is certainly doable in various ways, but not yet sure what the right way to do so is. One potential way is to use Go's own `cgo` to both call and expose Go packages as C combined with `dlopen`, potentially using [go-ffi](https://bitbucket.org/binet/go-ffi). Another is via `net/rpc`. A third interesting option is to model the namespaces using [go9p](http://code.google.com/p/go9p/). Regardless the approach, it will need a way to keep or migrate the state of vars.

### Why?

Mainly for fun and to learn Go. But it is also trying to address issues around fast compilation and start-up time (a current concern in the Clojure world, see below). The goal isn't to compete with the JVM or V8 on performance. Another obvious reason would be to leverage the Go ecosystem using ClojureScript. It's also meant to simplify creating new emitters for ClojureScript. I'm aware of `tools.analyzer` etc, but as far as I've seen, they don't help you implement `clojure.core` and the actual language run-time. That said, building on a fork of `tools.analyzer.js` instead of `cljs.analyzer` is a likely future direction.

### See Also

http://martintrojer.github.io/clojure/2014/04/05/the-clojure-repl-a-blessing-and-a-curse/

#### Lisp's in Go:

https://github.com/tcard/gojure
https://github.com/bjeanes/gojure
https://github.com/zhemao/glisp

#### Dialects of Clojure

http://dev.clojure.org/display/design/Release.Next+Planning
https://github.com/clojure/tools.analyzer/
http://clojure-android.info/blog/2014/08/12/gsoc-2014-skummet-alpha1/
https://github.com/arrdem/oxcart
https://github.com/mikera/kiss/

### Ports of ClojureScript

https://github.com/schani/clojurec
https://github.com/takeoutweight/clojure-scheme
https://github.com/raph-amiard/clojurescript-lua
https://github.com/kanaka/clojurescript

## License

**cljs.go**

Copyright © 2014 Håkan Råberg

Distributed under the Eclipse Public License either version 1.0 or (at
your option) any later version.

**ClojureScript**

    Copyright (c) Rich Hickey. All rights reserved. The use and
    distribution terms for this software are covered by the Eclipse
    Public License 1.0 (http://opensource.org/licenses/eclipse-1.0.php)
    which can be found in the file epl-v10.html at the root of this
    distribution. By using this software in any fashion, you are
    agreeing to be bound by the terms of this license. You must
    not remove this notice, or any other, from this software.
