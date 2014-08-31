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
		return js.Console.Log.Apply(js.Console, []interface{}{x})
	})
	return nil
})

var X_STAR_main_cli_fn_STAR_ interface{}

var pr_opts = Fn(func() interface{} {
	return nil
})

var Newline = Fn(func(opts interface{}) interface{} {
	if Truth_(X_STAR_print_newline_STAR_) {
		X_STAR_print_fn_STAR_.(CljsCoreIFn).X_invoke_Arity1("\n")
	}
	return nil
})

var Println = Fn(func(objs ...interface{}) interface{} {
	X_STAR_print_fn_STAR_.(CljsCoreIFn).X_invoke_Arity1(fmt.Sprint(objs...))
	Newline.X_invoke_Arity1(pr_opts.X_invoke_Arity0())
	return nil
})

// core protocols

// Protocols in ClojureScript don't seem to support vargs.
// In cljs.core, only IFn, IReduce, IIndexed, ILookup, and ISwap have overloaded arities.
// IFn is a special case which drops the receiver arg.

type CljsCoreINamed interface {
	X_name_Arity1() string
	X_namespace_Arity1() string
}

var X_name = Fn(func(this interface{}) interface{} {
	return this.(CljsCoreINamed).X_name_Arity1()
})

var X_namespace = Fn(func(this interface{}) interface{} {
	return this.(CljsCoreINamed).X_namespace_Arity1()
})

func init() {
	RegisterProtocol_("cljs.core/INamed", (*CljsCoreINamed)(nil))
}

// Note the arity difference here, starting with 0 unlike other protocols
type CljsCoreIFn interface {
	X_invoke_Arity0() interface{}
	X_invoke_Arity1(a interface{}) interface{}
	X_invoke_Arity2(a interface{}, b interface{}) interface{}
	X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{}
	X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{}
	X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{}
	X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{}
	X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{}
	X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{}
	X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{}
	X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{}
	X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{}
	X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{}
	X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{}
	X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{}
	X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{}
	X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{}
	X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{}
	X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{}
	X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{}
	X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IFn", (*CljsCoreIFn)(nil))
}

var X_invoke = func(_invoke *AFn) *AFn {
	return Fn(_invoke, func(this interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity0()
	}, func(this interface{}, a interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity1(a)
	}, func(this interface{}, a interface{}, b interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity2(a, b)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity3(a, b, c)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity4(a, b, c, d)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity5(a, b, c, d, e)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity6(a, b, c, d, e, f)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity7(a, b, c, d, e, f, g)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity8(a, b, c, d, e, f, g, h)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity9(a, b, c, d, e, f, g, h, i)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity10(a, b, c, d, e, f, g, h, i, j)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity11(a, b, c, d, e, f, g, h, i, j, k)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity12(a, b, c, d, e, f, g, h, i, j, k, l)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity13(a, b, c, d, e, f, g, h, i, j, k, l, m)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity14(a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity15(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity16(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity17(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity18(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s)
	}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
		return this.(CljsCoreIFn).X_invoke_Arity19(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s, t)
	})
}(&AFn{})

type CljsCoreILookup interface {
	X_lookup_Arity2(k interface{}) interface{}
	X_lookup_Arity3(k, notFound interface{}) interface{}
}

var X_Lookup = Fn(func(this, k interface{}) interface{} {
	return this.(CljsCoreILookup).X_lookup_Arity2(k)
}, func(this, k, notFound interface{}) interface{} {
	return this.(CljsCoreILookup).X_lookup_Arity3(k, notFound)
})

func init() {
	RegisterProtocol_("cljs.core/ILookup", (*CljsCoreILookup)(nil))
}

type CljsCoreUUID struct {
	Uuid interface{}
}

type CljsCoreKeyword struct {
	Ns, Name, Fqn, X_hash interface{}
}

type CljsCoreSymbol struct {
	Ns, Name, Str, X_hash, X_meta interface{}
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

func (this *CljsCoreSymbol) X_name_Arity1() string {
	return this.Name.(string)
}

func (this *CljsCoreSymbol) X_namespace_Arity1() string {
	return this.Ns.(string)
}

func (this *CljsCoreSymbol) X_invoke_Arity1(coll interface{}) interface{} {
	return X_Lookup.X_invoke_Arity2(coll, this)
}

func (this *CljsCoreSymbol) X_invoke_Arity2(coll, notFound interface{}) interface{} {
	return X_Lookup.X_invoke_Arity3(coll, this, notFound)
}

// Doesn't try to match cljs.core, just a temporary hack
type ObjMap map[interface{}]interface{}

func (coll ObjMap) X_lookup_Arity2(k interface{}) interface{} {
	return coll.X_lookup_Arity3(k, nil)
}

func (coll ObjMap) X_lookup_Arity3(k, notFound interface{}) interface{} {
	val := coll[k]
	if val == nil {
		return notFound
	} else {
		return val
	}
}

var Symbol = func(Symbol *AFn) *AFn {
	return Fn(Symbol, func(name interface{}) interface{} {
		return Symbol.X_invoke_Arity2(nil, name)
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

var String_QMARK_ = Fn(&AFn{}, func(x interface{}) bool {
	return goog.IsString(x)
})

var Namespace = Fn(func(x interface{}) interface{} {
	if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2(Symbol.X_invoke_Arity2("cljs.core", "INamed"), x)) {
		return X_namespace.X_invoke_Arity1(x)
	} else {
		panic(&js.Error{Str.X_invoke_ArityVariadic("Doesn't support namespace: ", x)})
	}
})

var Name = Fn(func(x interface{}) interface{} {
	if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2(Symbol.X_invoke_Arity2("cljs.core", "INamed"), x)) {
		return X_name.X_invoke_Arity1(x)
	} else {
		if Truth_(String_QMARK_.X_invoke_Arity1(x)) {
			return x
		} else {
			panic(&js.Error{Str.X_invoke_ArityVariadic("Doesn't support name: ", x)})
		}
	}
})

var Array_seq = func(Array_seq *AFn) *AFn {
	return Fn(Array_seq, func(array interface{}) interface{} {
		return Array_seq.X_invoke_Arity2(array, 0)
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

var Str = func(Str *AFn) *AFn {
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

		var sb, more interface{} = &goog_string.StringBuffer{Str.X_invoke_Arity1(x)}, ys
		for {
			if Truth_(more) {
				sb = Native_invoke_instance_method.X_invoke_Arity3(sb, "Append",
					[]interface{}{Str.X_invoke_Arity1(First.X_invoke_Arity1(more))})
				more = Next.X_invoke_Arity1(more)
				continue
			} else {
				return fmt.Sprint(sb)
			}
		}
	})
}(&AFn{})

var Not = Fn(func(x interface{}) bool {
	return !Truth_(x)
})

type CljsCoreISeq interface{}
