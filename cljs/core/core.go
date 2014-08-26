package core

import (
	"github.com/hraberg/cljs.go/goog"
	goog_string "github.com/hraberg/cljs.go/goog/string"
	"github.com/hraberg/cljs.go/js"
)

import "fmt"

// This file will eventually (partly) be generated from cljs.core. Some of this will serve as overrides over cljs.core.
// Note that "source maps" are possible using //line cljs/core.cljs:1 directives (absolute or relative to package)
// Also: packages can be imported relative to each other.

var X_STAR_print_fn_STAR_ interface{} = Fn(func(_ interface{}) interface{} {
	panic(&js.Error{"No *print-fn* fn set for evaluation environment"})
})

var X_STAR_print_newline_STAR_ interface{} = true

var Enable_console_print_BANG_ = Fn(func() interface{} {
	X_STAR_print_newline_STAR_ = false
	X_STAR_print_fn_STAR_ = Fn(func(x interface{}) interface{} {
		js.Console.Log(x)
		return nil
	})
	return nil
})

var X_STAR_main_cli_fn_STAR_ interface{}

var pr_opts = Fn(func() interface{} {
	return nil
})

var Newline = Fn(func(opts interface{}) interface{} {
	if Truth_(X_STAR_print_newline_STAR_) {
		X_STAR_print_fn_STAR_.(IFn).Invoke_Arity1("\n")
	}
	return nil
})

var Println = Fn(func(objs ...interface{}) interface{} {
	X_STAR_print_fn_STAR_.(IFn).Invoke_Arity1(fmt.Sprint(objs...))
	Newline.Invoke_Arity1(pr_opts.Invoke_Arity0())
	return nil
})

// core protocols

// Protocols in ClojureScript don't seem to support vargs.
// In cljs.core, only IFn, IReduce, IIndexed, ILookup, and ISwap have overloaded arities.
// IFn is a special case which drops the receiver arg.

type INamed interface {
	Name_Arity1() string
	Namespace_Arity1() string
}

var X_Name = Fn(func(this interface{}) interface{} {
	return this.(INamed).Name_Arity1()
})

var X_Namespace = Fn(func(this interface{}) interface{} {
	return this.(INamed).Namespace_Arity1()
})

func init() {
	Native_register_protocol("cljs.core/INamed", (*INamed)(nil))
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

var X_Invoke = Fn(func(this interface{}) interface{} {
	return this.(IFn).Invoke_Arity0()
}, func(this, a interface{}) interface{} {
	return this.(IFn).Invoke_Arity1(a)
}, func(this, a, b interface{}) interface{} {
	return this.(IFn).Invoke_Arity2(a, b)
}, func(this, a, b, c interface{}) interface{} {
	return this.(IFn).Invoke_Arity3(a, b, c)
}, func(this, a, b, c, d interface{}) interface{} {
	return this.(IFn).Invoke_Arity4(a, b, c, d)
}, func(this_a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
	return this_a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest[0].(IFn).
		Invoke_ArityVariadic(this_a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest[1:]...)
})

func init() {
	Native_register_protocol("cljs.core/IFn", (*IFn)(nil))
}

type ILookup interface {
	Lookup_Arity2(k interface{}) interface{}
	Lookup_Arity3(k, notFound interface{}) interface{}
}

var X_Lookup = Fn(func(this, k interface{}) interface{} {
	return this.(ILookup).Lookup_Arity2(k)
}, func(this, k, notFound interface{}) interface{} {
	return this.(ILookup).Lookup_Arity3(k, notFound)
})

func init() {
	Native_register_protocol("cljs.core/ILookup", (*ILookup)(nil))
}

// These are overridden in rt.go
var Apply IFn
var Native_satisfies_QMARK_ *AFnPrimitive
var Implements_QMARK_ *AFnPrimitive

type CljsCoreUUID struct {
	Uuid interface{}
}

type CljsCoreKeyword struct {
	Ns, Name, Fqn, Hash interface{}
	AbstractIFn
}

type CljsCoreSymbol struct {
	Ns, Name, Str, Hash, Meta interface{}
	AbstractIFn
}

// This is a "Go" protocol which compiles to real type hinted methods without IFn dispatch.
// While Object is provided by rt, the same technique applies for interop with normal Go interfaces.
// The naming convetion is (-equiv ) for IEquiv ClojureScript dispatch, and (equiv ) for Go dispatch.

func (this *CljsCoreSymbol) ToString() string {
	return (this.Str).(string)
}

// Will be added automatically when implementing ToString
func (this *CljsCoreSymbol) String() string {
	return this.ToString()
}

func (this *CljsCoreSymbol) Equiv(other interface{}) bool {
	if ptr, instanceof := other.(*CljsCoreSymbol); instanceof {
		other = *ptr
	}
	return *this == other
}

func (this *CljsCoreSymbol) Name_Arity1() string {
	return this.Name.(string)
}

func (this *CljsCoreSymbol) Namespace_Arity1() string {
	return this.Ns.(string)
}

func (this *CljsCoreSymbol) Invoke_Arity1(coll interface{}) interface{} {
	return X_Lookup.Invoke_Arity2(coll, this)
}

func (this *CljsCoreSymbol) Invoke_Arity2(coll, notFound interface{}) interface{} {
	return X_Lookup.Invoke_Arity3(coll, this, notFound)
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

var Symbol = func(Symbol IFn) IFn {
	return Fn(Symbol, func(name interface{}) interface{} {
		return Symbol.Invoke_Arity2(nil, name)
	}, func(ns, name interface{}) interface{} {
		symStr := func() interface{} {
			if ns != nil {
				return ns.(string) + "/" + name.(string)
			} else {
				return name
			}
		}()
		return &CljsCoreSymbol{Ns: ns, Name: name, Str: symStr}
	})
}(&AFn{})

var String_QMARK_ = Fn(&AFnPrimitive{}, func(x interface{}) bool {
	return goog.IsString(x)
}).(*AFnPrimitive)

var Namespace = Fn(func(x interface{}) interface{} {
	if Truth_(Implements_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "INamed"), x)) {
		return X_Namespace.Invoke_Arity1(x)
	} else {
		panic(&js.Error{Str.Invoke_ArityVariadic("Doesn't support namespace: ", x)})
	}
})

var Name = Fn(func(x interface{}) interface{} {
	if Truth_(Implements_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "INamed"), x)) {
		return X_Name.Invoke_Arity1(x)
	} else {
		if Truth_(String_QMARK_.Invoke_Arity1(x)) {
			return x
		} else {
			panic(&js.Error{Str.Invoke_ArityVariadic("Doesn't support name: ", x)})
		}
	}
})

var Array_seq = func(Array_seq IFn) IFn {
	return Fn(Array_seq, func(array interface{}) interface{} {
		return Array_seq.Invoke_Arity2(array, 0)
	}, func(array, i interface{}) interface{} {
		return array
	})
}(&AFn{})

var First = Fn(func(coll interface{}) interface{} {
	seq := coll.([]interface{})
	if seq != nil && len(seq) != 0 {
		return seq[0]
	}
	return nil
})

var Next = Fn(func(coll interface{}) interface{} {
	seq := coll.([]interface{})
	if seq != nil && len(seq) != 0 {
		return seq[1:]
	}
	return nil
})

var Str = func(Str IFn) IFn {
	return Fn(Str, func() interface{} {
		return ""
	}, func(x interface{}) interface{} {
		if x == nil {
			return ""
		} else {
			return fmt.Sprint(x)
		}
	}, func(x_ys ...interface{}) interface{} {
		var x, ys interface{} = x_ys[0], x_ys[1:]

		var sb, more interface{} = &goog_string.StringBuffer{Str.Invoke_Arity1(x)}, ys
		for {
			if Truth_(more) {
				sb = Native_invoke_instance_method.Invoke_Arity3(sb, "Append",
					[]interface{}{Str.Invoke_Arity1(First.Invoke_Arity1(more))})
				more = Next.Invoke_Arity1(more)
				continue
			} else {
				return fmt.Sprint(sb)
			}
		}
	})
}(&AFn{})

var Not = Fn(&AFnPrimitive{}, func(x interface{}) bool {
	return !Truth_(x)
}).(*AFnPrimitive)
