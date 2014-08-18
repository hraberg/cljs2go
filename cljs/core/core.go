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
	Name_Arity1() string
	Namespace_Arity1() string
}

var Name = AFn{
	Arity1: func(this interface{}) interface{} {
		return this.(INamed).Name_Arity1()
	},
}

var Namespace = AFn{
	Arity1: func(this interface{}) interface{} {
		return this.(INamed).Namespace_Arity1()
	},
}

func init() {
	protocols[Symbol.Arity2("cljs.core", "INamed")] = reflect.TypeOf((*INamed)(nil)).Elem()
}

// Note the arity difference here, starting with 0 unlike other protocols
type IFn interface {
	Invoke_Arity0() interface{}
	Invoke_Arity1(a interface{}) interface{}
	Invoke_Arity2(a, b interface{}) interface{}
	Invoke_Arity3(a, b, c interface{}) interface{}
	Invoke_Arity4(a, b, c, d interface{}) interface{}
	Invoke_Arity5(a, b, c, d, e interface{}) interface{}
	Invoke_ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{}
}

var Invoke IFn = AFn{
	Arity1: func(this interface{}) interface{} {
		return this.(IFn).Invoke_Arity0()
	},
	Arity2: func(this, a interface{}) interface{} {
		return this.(IFn).Invoke_Arity1(a)
	},
	Arity3: func(this, a, b interface{}) interface{} {
		return this.(IFn).Invoke_Arity2(a, b)
	},
	Arity4: func(this, a, b, c interface{}) interface{} {
		return this.(IFn).Invoke_Arity3(a, b, c)
	},
	Arity5: func(this, a, b, c, d interface{}) interface{} {
		return this.(IFn).Invoke_Arity4(a, b, c, d)
	},
	ArityVariadic: func(this_a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
		return this_a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest[0].(IFn).
			Invoke_ArityVariadic(this_a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest[1:]...)
	},
}

func init() {
	protocols[Symbol.Arity2("cljs.core", "IFn")] = reflect.TypeOf((*IFn)(nil)).Elem()
}

type ILookup interface {
	Lookup_Arity2(k interface{}) interface{}
	Lookup_Arity3(k, notFound interface{}) interface{}
}

var Lookup = AFn{
	Arity2: func(this, k interface{}) interface{} {
		return this.(ILookup).Lookup_Arity2(k)
	},
	Arity3: func(this, k, notFound interface{}) interface{} {
		return this.(ILookup).Lookup_Arity3(k, notFound)
	},
}

func init() {
	protocols[Symbol.Arity2("cljs.core", "ILookup")] = reflect.TypeOf((*ILookup)(nil)).Elem()
}

// Potential alternative, each protocol is a named struct with Arity0 etc fields:
// type ILookup struct { Arity2, Arity3 }
// A type has an ILookup field, with it's implementations of Arity1 etc.
// Looks like this in CLJS: o.cljs$core$ILookup$_lookup$arity$2;
// The owning namespace has a normal AFn IFn wrapper for Lookup which delegates Arity1 etc. to the receiver.

// These naming conventions stem from the IFn protocol, but currently there's no real connection here.
// Protocols are currently implemented as real Go interfaces with bridge methods hiding the AFn.
// Implementations need to set up actual named locals as argument, and create an array-seq from the varargs.
type ArityVariadic func(...interface{}) interface{}
type Arity0 func() interface{}
type Arity1 func(interface{}) interface{}
type Arity2 func(_, _ interface{}) interface{}
type Arity3 func(_, _, _ interface{}) interface{}
type Arity4 func(_, _, _, _ interface{}) interface{}
type Arity5 func(_, _, _, _, _ interface{}) interface{}

type Arity0F func() float64
type Arity1IF func(interface{}) float64
type Arity1FF func(float64) float64
type Arity1FI func(float64) interface{}
type Arity2IIF func(_, _ interface{}) float64
type Arity2IFI func(interface{}, float64) interface{}
type Arity2IFF func(interface{}, float64) float64
type Arity2FII func(float64, interface{}) interface{}
type Arity2FIF func(float64, interface{}) float64
type Arity2FFI func(_, _ float64) interface{}
type Arity2FFF func(_, _ float64) float64

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
	MaxFixedArity int
	ArityVariadic
	Arity0
	Arity1
	Arity2
	Arity3
	Arity4
	Arity5
	Arity0F
	Arity1IF
	Arity1FF
	Arity1FI
	Arity2IIF
	Arity2IFI
	Arity2IFF
	Arity2FII
	Arity2FIF
	Arity2FFI
	Arity2FFF
}

func ThrowArity(f interface{}, arity int) interface{} {
	if f == nil {
		panic(js.Error{fmt.Sprint("Invalid arity: ", arity)})
	}
	return f
}

func (this AFn) Call(args ...interface{}) interface{} {
	argc := len(args)
	if argc > this.MaxFixedArity && this.ArityVariadic != nil {
		return this.ArityVariadic(args...)
	}
	switch argc {
	case 0:
		switch {
		case this.Arity0 != nil:
			return this.Arity0()
		case this.Arity0F != nil:
			return this.Arity0F()
		}
	case 1:
		switch {
		case this.Arity1 != nil:
			return this.Arity1(args[0])
		case this.Arity1IF != nil:
			return this.Arity1IF(args[0])
		case this.Arity1FI != nil:
			return this.Arity1FI(args[0].(float64))
		case this.Arity1FF != nil:
			return this.Arity1FF(args[0].(float64))
		}
	case 2:
		switch {
		case this.Arity2 != nil:
			return this.Arity2(args[0], args[1])
		case this.Arity2IIF != nil:
			return this.Arity2IIF(args[0], args[1])
		case this.Arity2IFI != nil:
			return this.Arity2IFI(args[0], args[1].(float64))
		case this.Arity2IFF != nil:
			return this.Arity2IFI(args[0], args[1].(float64))
		case this.Arity2FII != nil:
			return this.Arity2FII(args[0].(float64), args[1])
		case this.Arity2FIF != nil:
			return this.Arity2FIF(args[0].(float64), args[1])
		case this.Arity2FFI != nil:
			return this.Arity2FFI(args[0].(float64), args[1].(float64))
		case this.Arity2FFF != nil:
			return this.Arity2FFF(args[0].(float64), args[1].(float64))
		}
	case 3:
		if this.Arity3 != nil {
			return this.Arity3(args[0], args[1], args[3])
		}
	case 4:
		if this.Arity4 != nil {
			return this.Arity4(args[0], args[1], args[3], args[4])
		}
	case 5:
		if this.Arity5 != nil {
			return this.Arity5(args[0], args[1], args[3], args[4], args[5])
		}

	}
	return ThrowArity(nil, argc)
}

func (this AFn) Invoke_ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
	ThrowArity(this.ArityVariadic, len(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest))
	return this.ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest...)
}

func (this AFn) Invoke_Arity0() interface{} {
	ThrowArity(this.Arity0, 0)
	return this.Arity0()
}

func (this AFn) Invoke_Arity1(a interface{}) interface{} {
	ThrowArity(this.Arity1, 1)
	return this.Arity1(a)
}

func (this AFn) Invoke_Arity2(a, b interface{}) interface{} {
	ThrowArity(this.Arity2, 2)
	return this.Arity2(a, b)
}

func (this AFn) Invoke_Arity3(a, b, c interface{}) interface{} {
	ThrowArity(this.Arity3, 3)
	return this.Arity3(a, b, c)
}

func (this AFn) Invoke_Arity4(a, b, c, d interface{}) interface{} {
	ThrowArity(this.Arity4, 4)
	return this.Arity4(a, b, c, d)
}

func (this AFn) Invoke_Arity5(a, b, c, d, e interface{}) interface{} {
	ThrowArity(this.Arity4, 4)
	return this.Arity4(a, b, c, d)
}

type AbstractIFn struct{}

func (this AbstractIFn) Invoke_ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
	return ThrowArity(nil, len(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest))
}

func (this AbstractIFn) Invoke_Arity0() interface{} {
	return ThrowArity(nil, 0)
}

func (this AbstractIFn) Invoke_Arity1(a interface{}) interface{} {
	return ThrowArity(nil, 1)
}

func (this AbstractIFn) Invoke_Arity2(a, b interface{}) interface{} {
	return ThrowArity(nil, 2)
}

func (this AbstractIFn) Invoke_Arity3(a, b, c interface{}) interface{} {
	return ThrowArity(nil, 3)
}

func (this AbstractIFn) Invoke_Arity4(a, b, c, d interface{}) interface{} {
	return ThrowArity(nil, 4)
}

func (this AbstractIFn) Invoke_Arity5(a, b, c, d, e interface{}) interface{} {
	return ThrowArity(nil, 4)
}

func Apply(f interface{}, args ...interface{}) interface{} {
	argc := len(args)
	if argc < 1 {
		ThrowArity(nil, argc)
	}
	var spread = args[argc-1].([]interface{}) // This will be a seq in real life.
	return f.(AFn).Call(append(args[:argc-1], spread...)...)
}

// This isn't really needed unless we actually make it faster. It's only used for variadic fns in CLJS.
func (this AFn) ApplyTo(args ...interface{}) interface{} {
	return Apply(this, args...)
}

type CljsCoreSymbol struct {
	ns, name, str, _hash, _meta interface{}
	AbstractIFn
}

var Symbol = func() AFn {
	Symbol := AFn{}
	Symbol.Arity1 = func(name interface{}) interface{} {
		return Symbol.Arity2(nil, name)
	}
	Symbol.Arity2 = func(ns, name interface{}) interface{} {
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

func (this CljsCoreSymbol) Name_Arity1() string {
	return this.name.(string)
}

func (this CljsCoreSymbol) Namespace_Arity1() string {
	return this.ns.(string)
}

func (this CljsCoreSymbol) String() string {
	return this.str.(string)
}

func (this CljsCoreSymbol) Invoke_Arity1(coll interface{}) interface{} {
	return coll.(ILookup).Lookup_Arity2(this)
}

func (this CljsCoreSymbol) Invoke_Arity2(coll, notFound interface{}) interface{} {
	return coll.(ILookup).Lookup_Arity3(this, notFound)
}

// Doesn't try to match cljs.core, just a temporary hack
type ObjMap map[interface{}]interface{}

func (coll ObjMap) Lookup_Arity2(k interface{}) interface{} {
	return coll.Lookup_Arity3(k, nil)
}

func (coll ObjMap) Lookup_Arity3(k, notFound interface{}) interface{} {
	val := coll[k]
	if val == nil {
		return notFound
	} else {
		return val
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
