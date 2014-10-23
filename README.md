# cljs2go

*Way to Go!* -- [Siobhan Sharpe](http://londonist.com/2012/07/did-boris-write-twenty-twelve-way-to-go-gag.php)


**ClojureScript to Go**. Implemented as an overlay onto ClojureScript, instead of a fork. Go is emitted from `cljs.go.compiler`, which is a patched version of `cljs.compiler`. At run time, a thin JS compatibility is provided to avoid having to touch `cljs.core`, these `js` and `goog` packages aren't intended for end-user usage.

Once the compiler starts working, the plan is to provide a second monkey patch in spirit of the first, but this time for the Java dependencies, like `java.io.File` and their Clojure counterpart, `clojure.java.io`. This patch is to be able to compile the compiler itself to Go. This won't change the run-time characteristics of the compiled programs, nor will it introduce `eval`, it's simply done to leverage the Go build pipeline. ClojureScript has a reader in `cljs.reader` that potentially can replace `tools.reader`. Macros used via `:require-macros` needs to be compiled to Go first, and an ad-hoc wrapper for the compiler depending on them will be created to compile the actual source (the basic compiler has the `cljs.core` macros built in).

While not introducing `eval`, it does open up for adding it as a third layer, the problem lies on the Go side, as you need a way to build and link code into the running process. This is certainly doable in various ways, but not yet sure what the right way to do so is. One potential way is to use Go's own `cgo` to both call and expose Go packages as C combined with `dlopen`, potentially using [go-ffi](https://bitbucket.org/binet/go-ffi). Another is via `net/rpc`. A third interesting option is to model the namespaces using [go9p](http://code.google.com/p/go9p/). Regardless the approach, it will need a way to keep or migrate the state of vars. There's also [Go Execution Modes](https://docs.google.com/document/d/1nr-TQHw*er6GOQRsF6T43GGhFDelrAP0NqSS*00RgZQ/preview?sle=true) which hints at a future `-buildmode=plugin`, potentially in Go 1.5.

### Why?

Mainly for fun and to learn Go. But it is also trying to address issues around fast compilation and start-up time (a current concern in the Clojure world, see below). On the other side, see Fernando's [motivation](https://github.com/eudoxia0/corvus#why) for his corvus LLVM Lisp: "The era of dynamic languages is over. There is currently a 'race to the bottom': Languages like Rust and Nimrod are merging the world of low-level programming with high-level compiler features." The reason to start this exploration with Go itself is most pragmatically explained in Rob's motivating [talk](https://talks.golang.org/2012/splash.article). And, Go has momentum.

The goal isn't to compete with the JVM or V8 on performance. Another obvious reason would be to leverage the Go ecosystem using ClojureScript. It's also meant to simplify creating new emitters (at least for myself) for ClojureScript. I'm aware of `tools.analyzer` etc., but they don't help you implement `clojure.core` and the actual language run-time. That said, building on a fork of `tools.analyzer.js` instead of `cljs.analyzer` is a likely future direction. Currently, ClojureScript's analyzer is not modified, and a lot of extra mess happens in the Go compiler. Some of this could / should be moved to analyzer passes.

### How?

The easiest way to start understanding how the compiler works is to look at the emitted Go code. `core.go` is `core.cljs` compiled to Go. `overrides.go` are Go specific overrides compiled from `overrides.cljs`. `rt.go` is a handwritten Go file providing the implementation needed, mainly `AFn` and wrappers around Go `reflect` capabilities and some coercing functions (like `Truth_`) used by the emitted code. ClojureScript functions are represented by this `AFn` struct, which bundles up potentially more than one Go function into a ClojureScript one. `deftype` compiles to Go structs, `defprotocol` to Go interfaces. Protocol methods are real Go methods, not `AFn`s. `number`, `boolean`, `array`, `string` and `seq` are recognized and compiled to `float64`, `bool`, `[]interface{}`, `string` and `CljsCoreISeq`. Like in JavaScript there are no real integers, all numbers are `float64` for simplicity, but this is likely to change.

As can be seen on `ISeq` above, types and protocols are ns-prefixed to not clash with functions (like `Symbol` vs. `symbol`). Public functions in Go must start with an uppercase character. Functions starting with a `_` (munged from `-`) have an `X` in front of them. The arity is appended, so a full compiled name will look like this: `X_invoke_Arity1`. `ArityVariadic` is a special case which regardless of how many fixed parameters take a single varargs parameter which is then unpacked inside the generated body. This is to simplify dispatch and avoid having 20+ different varargs signatures (this might change). Functions with primitives (up to 3 arguments) are compiled into something like `Arity1FF` (takes one `float64` and returns a `float64`). These functions live beneath the normal protocol `IFn` dispatch. A normal (without primitives) bridge `IFn` function is also generated (using [MakeFunc](http://golang.org/pkg/reflect/#MakeFunc)) for the matching arity. Like in ClojureScript, there's a special protocol `Object` that allow creating of methods that look like real Go methods (no `_ArityN`). These methods, like any host methods or functions, are invoked by the dot notation, like `(.toString x)`. There's currently no way to create a plain Go `func`, but this will likely become a macro.

When compiling, the unaltered `cljs.analyzer` from ClojureScript is used to build the AST (see Nicola's [AST Quickref](http://clojure.github.io/tools.analyzer/spec/quickref.html)). The analyzer depends on the `cljs.core` macros (`core.clj`), which (like all ClojureScript macros) are written in Clojure. This file is replaced by `cljs.go.core`, and heavily uses the `js*` macro to emit literal Go code. Once the AST has been generated, it's fed into the `cljs.go.compiler`, which emits the Go source code. This is in turn (usually) fed into [`goimports`](http://godoc.org/code.google.com/p/go.tools/cmd/goimports) (a version of `gofmt` that also fixes the imports) and then finally `go build` (or `go install`).

### Usage

*Not ready yet.*

While the compiler more or less works, and passes most of ClojureScript's test suite, it's not packaged for actual use, if you want to play with it, it's easiest to fire up a REPL and look at the `cljs.go` namespace. You can also run the tests using `go test -v ./...`. Note that you need to check this out into `$GOPATH/src/github.com/hraberg/cljs2go` for that to work. `go get` will eventually be supported.

The Hello World sample can be built like this:

```clj
;; Generate the Go source from a Clojure REPL
cljs.go> (compile-file "samples" "samples/hello.cljs")
```

```bash
# go run
$ cd samples/hello
$ go run main.go
Hello World

# go build
$ go build -o hello main.go
$ ./hello
Hello World

```

### Road Map

*We're currently in phase 1. `cljs.core-tests` now pass.*

#### Phase 0

* Generate compiling ClojureScript with mapping for all basic constructs to Go.
* Minimal JavaScript host libraries assumed by ClojureScript.

#### Phase 1

* Getting the relevant parts of ClojureScript's test suite to run.
* Ability to build an app with more than one namespace and a main fn which depends on the core.
* Go interop, `func` macro, coercions functions, ability to write moderately useful programs.

#### Phase 2

* Ability transform `core.clj`, `analyzer.clj` and `compiler.clj` to compiling ClojureScript.
* Compiling one reader, either `cljs.reader` (which might need extensions) or `tools.reader` to ClojureScript.
* Bootstrap the compiler using the above to Go.
* Ability to build new wrapping compilers for `require-macros` to work. Assume the macros are valid ClojureScript.

#### Phase 3

* Performance. Revisit all the basic assumptions. Generate cleaner code, remove redundant constructs.
* Move to the `tools.analyzer` stack.
* Full bootstrap, porting the hacks that allow the compiler to compile itself.
* `core.async` using real Go blocks and channels.
* ClojureScript to Go REPL

#### Phase 4

* Porting to other targets.

### See Also

* http://martintrojer.github.io/clojure/2014/04/05/the-clojure-repl-a-blessing-and-a-curse/
* https://al3x.net/2014/09/16/thoughts-on-five-years-of-emerging-languages.html
* http://blog.cognitect.com/blog/2014/10/20/results-of-2014-state-of-clojure-and-clojurescript-survey

#### Lisp's in Go:

* https://github.com/tcard/gojure
* https://github.com/bjeanes/gojure
* https://github.com/zhemao/glisp

#### Dialects of Clojure

* http://dev.clojure.org/display/design/Release.Next+Planning
* https://github.com/clojure/tools.analyzer/
* http://clojure-android.info/blog/2014/08/12/gsoc-2014-skummet-alpha1/
* https://github.com/arrdem/oxcart
* https://github.com/mikera/kiss/

### Ports of ClojureScript

* https://github.com/schani/clojurec
* https://github.com/takeoutweight/clojure-scheme
* https://github.com/raph-amiard/clojurescript-lua
* https://github.com/kanaka/clojurescript (in ClojureScript)

### Lower-level Lisps

* https://github.com/pixie-lang/pixie
* https://github.com/artagnon/rhine
* https://github.com/eudoxia0/corvus
* http://piumarta.com/software/maru/

### Other potential targets

* http://nimrod-lang.org
* http://www.gnu.org/software/libjit/
* http://www.rust-lang.org/
* http://julialang.org/
* http://llvm.org/
* http://openjdk.java.net/projects/graal/ / Indy JVM
* https://software.intel.com/en-us/articles/introduction-to-x64-assembly

## License

**cljs2go**

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
