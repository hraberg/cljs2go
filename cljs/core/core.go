package core

import (
	"os"
	"reflect"

	"github.com/hraberg/cljs.go/js"
)

import "fmt"

// This file will eventually (partly) be generated from cljs.core. Some of this will serve as overrides over cljs.core.

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

// core protocols

var protocols = map[interface{}]reflect.Type{}

func NativeSatisifes_QMARK_(p, x interface{}) interface{} {
	return reflect.ValueOf(x).Type().Implements(protocols[p])
}

type INamed interface {
	Name() string
	Namespace() string
}

func init() {
	protocols[Symbol.CljsCoreIFn_InvokeArity2("cljs.core", "INamed")] = reflect.TypeOf((*INamed)(nil)).Elem()
}

// How to properly consolidate this with the internal CljsCoreIFn_InvokeArityN stuff?
type IFn interface {
	Invoke(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{}
}

func init() {
	protocols[Symbol.CljsCoreIFn_InvokeArity2("cljs.core", "IFn")] = reflect.TypeOf((*IFn)(nil)).Elem()
}

type ILookup interface {
	Lookup(k interface{}, notFound ...interface{}) interface{}
}

func init() {
	protocols[Symbol.CljsCoreIFn_InvokeArity2("cljs.core", "ILookup")] = reflect.TypeOf((*ILookup)(nil)).Elem()
}

// These naming conventions stem from the IFn protocol, but currently there's no real connection here.
// Protocols are currently implemented as real Go interfaces with bridge methods hiding the AFn.
// Implementations need to set up actual named locals as argument, and create an array-seq from the varargs.
type CljsCoreIFn_InvokeArityVariadic func(...interface{}) interface{}
type CljsCoreIFn_InvokeArity0 func() interface{}
type CljsCoreIFn_InvokeArity1 func(interface{}) interface{}
type CljsCoreIFn_InvokeArity2 func(_, _ interface{}) interface{}
type CljsCoreIFn_InvokeArity3 func(_, _, _ interface{}) interface{}
type CljsCoreIFn_InvokeArity4 func(_, _, _, _ interface{}) interface{}

type CljsCoreIFn_InvokeArity0D func() float64
type CljsCoreIFn_InvokeArity1OD func(interface{}) float64
type CljsCoreIFn_InvokeArity1DD func(float64) float64
type CljsCoreIFn_InvokeArity1DO func(float64) interface{}
type CljsCoreIFn_InvokeArity2OOD func(_, _ interface{}) float64
type CljsCoreIFn_InvokeArity2ODO func(interface{}, float64) interface{}
type CljsCoreIFn_InvokeArity2ODD func(interface{}, float64) float64
type CljsCoreIFn_InvokeArity2DOO func(float64, interface{}) interface{}
type CljsCoreIFn_InvokeArity2DOD func(float64, interface{}) float64
type CljsCoreIFn_InvokeArity2DDO func(_, _ float64) interface{}
type CljsCoreIFn_InvokeArity2DDD func(_, _ float64) float64

// There's a protocol called cljs.core.IFn we need to cooperate with, so we use AFn for now.
// So we might need to complicate this for various reasons, Keywords for example are IFns by protocol.
// That is, they implement CljsCoreIFn_InvokeArity1 (and 2), so we might need to re-add the interfaces.
// CLJS also (among other things) adds .call and .apply when implementing the IFn protocol, see cljs.core/add-ifn-methods, clj
// The easiest way to acheive this is renaming (and hide) the fields, and make CljsCoreIFn_InvokeArity1 an interface method.
// Then there's the issue of other protocols and how to represent them, as we prefer to keep them as Go interfaces.
// As can be seen above regarding IFn, it would be nice to ensure that IFn is just a special case that emit* :invoke knows about.
// IFn -invoke is like any other protocol in JS. There's a dispatch fn setup which looks for an implementation on the receiver.
// That is, cljs.core._invoke will call  receiver.cljs$core$IFn$_invoke$arity$1(receiver, x)
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
	CljsCoreIFn_InvokeArity0D
	CljsCoreIFn_InvokeArity1OD
	CljsCoreIFn_InvokeArity1DO
	CljsCoreIFn_InvokeArity1DD
	CljsCoreIFn_InvokeArity2OOD
	CljsCoreIFn_InvokeArity2ODO
	CljsCoreIFn_InvokeArity2ODD
	CljsCoreIFn_InvokeArity2DOO
	CljsCoreIFn_InvokeArity2DOD
	CljsCoreIFn_InvokeArity2DDO
	CljsCoreIFn_InvokeArity2DDD
}

func ThrowArity(arity int) interface{} {
	panic(js.Error{fmt.Sprint("Invalid arity: ", arity)})
}

func (this AFn) Invoke(args ...interface{}) interface{} {
	argc := len(args)
	if argc > this.CljsLangMaxFixedArity && this.CljsCoreIFn_InvokeArityVariadic != nil {
		return this.CljsCoreIFn_InvokeArityVariadic(args...)
	}
	switch argc {
	case 0:
		switch {
		case this.CljsCoreIFn_InvokeArity0 != nil:
			return this.CljsCoreIFn_InvokeArity0()
		case this.CljsCoreIFn_InvokeArity0D != nil:
			return this.CljsCoreIFn_InvokeArity0D()
		}
	case 1:
		switch {
		case this.CljsCoreIFn_InvokeArity1 != nil:
			return this.CljsCoreIFn_InvokeArity1(args[0])
		case this.CljsCoreIFn_InvokeArity1OD != nil:
			return this.CljsCoreIFn_InvokeArity1OD(args[0])
		case this.CljsCoreIFn_InvokeArity1DO != nil:
			return this.CljsCoreIFn_InvokeArity1DO(args[0].(float64))
		case this.CljsCoreIFn_InvokeArity1DD != nil:
			return this.CljsCoreIFn_InvokeArity1DD(args[0].(float64))
		}
	case 2:
		switch {
		case this.CljsCoreIFn_InvokeArity2 != nil:
			return this.CljsCoreIFn_InvokeArity2(args[0], args[1])
		case this.CljsCoreIFn_InvokeArity2OOD != nil:
			return this.CljsCoreIFn_InvokeArity2OOD(args[0], args[1])
		case this.CljsCoreIFn_InvokeArity2ODO != nil:
			return this.CljsCoreIFn_InvokeArity2OOD(args[0], args[1].(float64))
		case this.CljsCoreIFn_InvokeArity2ODD != nil:
			return this.CljsCoreIFn_InvokeArity2ODO(args[0], args[1].(float64))
		case this.CljsCoreIFn_InvokeArity2DOO != nil:
			return this.CljsCoreIFn_InvokeArity2DOO(args[0].(float64), args[1])
		case this.CljsCoreIFn_InvokeArity2DOD != nil:
			return this.CljsCoreIFn_InvokeArity2DOD(args[0].(float64), args[1])
		case this.CljsCoreIFn_InvokeArity2DDO != nil:
			return this.CljsCoreIFn_InvokeArity2DDO(args[0].(float64), args[1].(float64))
		case this.CljsCoreIFn_InvokeArity2DDD != nil:
			return this.CljsCoreIFn_InvokeArity2DDD(args[0].(float64), args[1].(float64))
		}
	case 3:
		if this.CljsCoreIFn_InvokeArity3 != nil {
			return this.CljsCoreIFn_InvokeArity3(args[0], args[1], args[3])
		}
	case 4:
		if this.CljsCoreIFn_InvokeArity4 != nil {
			return this.CljsCoreIFn_InvokeArity4(args[0], args[1], args[3], args[4])
		}
	}
	return ThrowArity(argc)
}

func Apply(f interface{}, args ...interface{}) interface{} {
	argc := len(args)
	if argc < 1 {
		ThrowArity(argc)
	}
	var spread = args[argc-1].([]interface{}) // This will be a seq in real life.
	return f.(IFn).Invoke(append(args[:argc-1], spread...)...)
}

// This isn't really needed unless we actually make it faster. It's only used for variadic fns in CLJS.
func (this AFn) CljsLangApplyTo(args ...interface{}) interface{} {
	return Apply(this, args...)
}

type CljsCoreSymbol struct {
	ns, name, str, _hash, _meta interface{}
}

var Symbol = func() AFn {
	Symbol := AFn{}
	Symbol.CljsCoreIFn_InvokeArity1 = func(name interface{}) interface{} {
		return Symbol.CljsCoreIFn_InvokeArity2(nil, name)
	}
	Symbol.CljsCoreIFn_InvokeArity2 = func(ns, name interface{}) interface{} {
		symStr := func() interface{} {
			if ns != nil {
				return ns.(string) + "/" + name.(string)
			} else {
				return name
			}
		}()
		return CljsCoreSymbol{ns: ns, name: name, str: symStr}
	}
	return Symbol
}()

func (this CljsCoreSymbol) Name() string {
	return this.name.(string)
}

func (this CljsCoreSymbol) Namespace() string {
	return this.ns.(string)
}

func (this CljsCoreSymbol) String() string {
	return this.str.(string)
}

func (this CljsCoreSymbol) Invoke(coll_notFound ...interface{}) interface{} {
	Invoke := AFn{}
	Invoke.CljsCoreIFn_InvokeArity1 = func(coll interface{}) interface{} {
		return coll.(ILookup).Lookup(this, nil)
	}
	Invoke.CljsCoreIFn_InvokeArity2 = func(coll, notFound interface{}) interface{} {
		return coll.(ILookup).Lookup(this, notFound)
	}
	argc := len(coll_notFound)
	switch argc {
	case 1:
		return Invoke.CljsCoreIFn_InvokeArity1(coll_notFound[0])
	case 2:
		return Invoke.CljsCoreIFn_InvokeArity2(coll_notFound[0], coll_notFound[1])
	default:
		return ThrowArity(argc)
	}
}

// Doesn't try to match cljs.core, just a temporary hack
type ObjMap map[interface{}]interface{}

func (coll ObjMap) Lookup(k interface{}, notFound ...interface{}) interface{} {
	Lookup := AFn{}
	Lookup.CljsCoreIFn_InvokeArity1 = func(k interface{}) interface{} {
		return coll.Lookup(k, nil)
	}
	Lookup.CljsCoreIFn_InvokeArity2 = func(k, notFound interface{}) interface{} {
		val := coll[k]
		if val == nil {
			return notFound
		} else {
			return val
		}
	}
	argc := len(notFound)
	switch argc {
	case 0:
		return Lookup.CljsCoreIFn_InvokeArity1(k)
	case 1:
		return Lookup.CljsCoreIFn_InvokeArity2(k, notFound[0])
	default:
		return ThrowArity(argc)
	}
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
