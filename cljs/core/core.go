package core

import (
	"os"

	"github.com/hraberg/cljs.go/js"
)

import "fmt"

// This file will eventually be generated from cljs.core.

var STAR_print_fn_STAR_ = func(_ interface{}) interface{} {
	panic(js.Error{"No *print-fn* fn set for evaluation environment"})
}

var STAR_print_newline_STAR_ = true

func Enable_console_print_BANG_() interface{} {
	STAR_print_newline_STAR_ = false
	STAR_print_fn_STAR_ = func(x interface{}) interface{} {
		js.Console.Log(x)
		return nil
	}
	return nil
}

var STAR_main_cli_fn_STAR_ interface{}

func pr_opts() interface{} {
	return nil
}

func Newline(opts interface{}) interface{} {
	if STAR_print_newline_STAR_ {
		STAR_print_fn_STAR_("\n")
	}
	return nil
}

func Println(objs ...interface{}) interface{} {
	STAR_print_fn_STAR_(fmt.Sprint(objs...))
	Newline(pr_opts())
	return nil
}

// Main preamble

func Main() {
	Enable_console_print_BANG_()
	args := make([]interface{}, len(os.Args[1:]))
	for i, a := range os.Args[1:] {
		args[i] = a
	}
	STAR_main_cli_fn_STAR_.(func(...interface{}) interface{})(args...)
}

// IFn

// Implementations need to set up actual named locals as argument, and create an array-seq from the varargs.
type CljsCoreIFn_InvokeArityVariadic func(...interface{}) interface{}
type CljsCoreIFn_InvokeArity0 func() interface{}
type CljsCoreIFn_InvokeArity1 func(interface{}) interface{}
type CljsCoreIFn_InvokeArity2 func(_, _ interface{}) interface{}
type CljsCoreIFn_InvokeArity3 func(_, _, _ interface{}) interface{}
type CljsCoreIFn_InvokeArity4 func(_, _, _, _ interface{}) interface{}

// There's a protocol called cljs.core.IFn we need to cooperate with, so we use AFn for now.
// So we might need to complicate this for various reasons, Keywords for example are IFns by protocol.
// That is, they implement CljsCoreIFn_InvokeArity1 (and 2), so we might need to re-add the interfaces.
// The easiest way to acheive this is renaming (and hide) the fields, and make CljsCoreIFn_InvokeArity1 an interface method.
// Then there's the issue of other protocols and how to represent them, as we prefer to keep them as Go interfaces.
// As can be seen above regarding IFn, it would be nice to ensure that IFn is just a special case that emit* :invoke knows about.
// My impression is that this is how CLJS does it. To reiterate - the reason this becomes extra messy in Go is lack of overloading.
// The main issue of making all invocations methods is that there's no way to create an anonymous type.
// The anonymous function bodies need to stay in the context where they're defined for closures to work.
// Interfaces might solve this though, as per earlier spikes.
// Another reason a function cannot (only) be an empty struct, is that in a protocol it's an interface method in Go.
// Most of this should be solvable with some indirection and interfaces.

// We also have the issue of functions that refer to themselves and how to ensure the var exists.
// This is extra interesting for anonymous "named" functions, ie. (fn foo [] (foo))

type AFn struct {
	CljsLangMaxFixedArity int
	CljsCoreIFn_InvokeArityVariadic
	CljsCoreIFn_InvokeArity0
	CljsCoreIFn_InvokeArity1
	CljsCoreIFn_InvokeArity2
	CljsCoreIFn_InvokeArity3
	CljsCoreIFn_InvokeArity4
}

func throwArity(arity int) interface{} {
	panic(js.Error{fmt.Sprint("Invalid arity: ", arity)})
}

func (this AFn) CljsCoreIFn_Invoke(args ...interface{}) interface{} {
	argc := len(args)
	switch {
	case argc == 0 && this.CljsCoreIFn_InvokeArity0 != nil:
		return this.CljsCoreIFn_InvokeArity0()
	case argc == 1 && this.CljsCoreIFn_InvokeArity1 != nil:
		return this.CljsCoreIFn_InvokeArity1(args[0])
	case argc == 2 && this.CljsCoreIFn_InvokeArity2 != nil:
		return this.CljsCoreIFn_InvokeArity2(args[0], args[1])
	case argc == 3 && this.CljsCoreIFn_InvokeArity3 != nil:
		return this.CljsCoreIFn_InvokeArity3(args[0], args[1], args[3])
	case argc == 4 && this.CljsCoreIFn_InvokeArity4 != nil:
		return this.CljsCoreIFn_InvokeArity4(args[0], args[1], args[3], args[4])
	case argc > this.CljsLangMaxFixedArity && this.CljsCoreIFn_InvokeArityVariadic != nil:
		return this.CljsCoreIFn_InvokeArityVariadic(args...)
	}
	return throwArity(argc)
}

func (this AFn) CljsLangApplyTo(args ...interface{}) interface{} {
	argc := len(args)
	if argc < 1 {
		throwArity(argc)
	}
	var spread = args[argc-1].([]interface{}) // This will be a seq in real life.
	return this.CljsCoreIFn_Invoke(append(args[:argc-1], spread...)...)
}

/*

There's a cljs.core, clj macro that eagerly turns things into strings in cljs.core.clj, but when not:
(str x) compiles to cljs.core.str.cljs$core$IFn$_invoke$arity$1(x)
cljs.core.str is the fn var
cljs$core$IFn is the protocol
$_invoke is the protocol fn
$arity$1 is the overload, ie [this a] in IFn, [x] in str. (How does the variadic [x & ys] map to the IFn arities?)

;; in cljs.core, cljs
(defn str
  "With no args, returns the empty string. With one arg x, returns
  x.toString().  (str nil) returns the empty string. With more than
  one arg, returns the concatenation of the str values of the args."
  ([] "")
  ([x] (if (nil? x)
         ""
         (.toString x)))
  ([x & ys]
    (loop [sb (StringBuffer. (str x)) more ys]
      (if more
        (recur (. sb  (append (str (first more)))) (next more))
        (.toString sb)))))

(defprotocol IFn
  (-invoke
    [this]
    [this a]
    [this a b]
    [this a b c]
    [this a b c d]
    [this a b c d e]
    [this a b c d e f]
    [this a b c d e f g]
    [this a b c d e f g h]
    [this a b c d e f g h i]
    [this a b c d e f g h i j]
    [this a b c d e f g h i j k]
    [this a b c d e f g h i j k l]
    [this a b c d e f g h i j k l m]
    [this a b c d e f g h i j k l m n]
    [this a b c d e f g h i j k l m n o]
    [this a b c d e f g h i j k l m n o p]
    [this a b c d e f g h i j k l m n o p q]
    [this a b c d e f g h i j k l m n o p q s]
    [this a b c d e f g h i j k l m n o p q s t]
    [this a b c d e f g h i j k l m n o p q s t rest]))
*/
