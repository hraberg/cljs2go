package core

import (
	"os"
	"reflect"
	"regexp"

	goog_string "github.com/hraberg/cljs.go/goog/string"
	"github.com/hraberg/cljs.go/js"
)

import "fmt"

// This file will eventually (partly) be generated from cljs.core. Some of this will serve as overrides over cljs.core.

var STAR_print_fn_STAR_ interface{} = Fn(func(_ interface{}) interface{} {
	panic(&js.Error{"No *print-fn* fn set for evaluation environment"})
})

var STAR_print_newline_STAR_ interface{} = true

var Enable_console_print_BANG_ = Fn(func() interface{} {
	STAR_print_newline_STAR_ = false
	STAR_print_fn_STAR_ = Fn(func(x interface{}) interface{} {
		js.Console.Log(x)
		return nil
	})
	return nil
})

var STAR_main_cli_fn_STAR_ interface{}

var pr_opts = Fn(func() interface{} {
	return nil
})

var Newline = Fn(func(opts interface{}) interface{} {
	if Truth_(STAR_print_newline_STAR_) {
		STAR_print_fn_STAR_.(IFn).Invoke_Arity1("\n")
	}
	return nil
})

var Println = Fn(func(objs ...interface{}) interface{} {
	STAR_print_fn_STAR_.(IFn).Invoke_Arity1(fmt.Sprint(objs...))
	Newline.Invoke_Arity1(pr_opts.Invoke_Arity0())
	return nil
})

// core protocols

// Protocols in ClojureScript don't seem to support vargs.
// In cljs.core, only IFn, IReduce, IIndexed, ILookup, and ISwap have overloaded arities.
// IFn is a special case which drops the receiver arg.

var protocols = map[string]reflect.Type{}

var NativeSatisifes_QMARK_ = Fn(func(p, x interface{}) interface{} {
	return reflect.ValueOf(x).Type().Implements(protocols[fmt.Sprint(p)])
})

type INamed interface {
	Name_Arity1() string
	Namespace_Arity1() string
}

var Name_ = Fn(func(this interface{}) interface{} {
	return this.(INamed).Name_Arity1()
})

var Namespace_ = Fn(func(this interface{}) interface{} {
	return this.(INamed).Namespace_Arity1()
})

func init() {
	protocols["cljs.core/INamed"] = reflect.TypeOf((*INamed)(nil)).Elem()
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

var Invoke_ = Fn(func(this interface{}) interface{} {
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
	protocols["cljs.core/IFn"] = reflect.TypeOf((*IFn)(nil)).Elem()
}

type ILookup interface {
	Lookup_Arity2(k interface{}) interface{}
	Lookup_Arity3(k, notFound interface{}) interface{}
}

var Lookup_ = Fn(func(this, k interface{}) interface{} {
	return this.(ILookup).Lookup_Arity2(k)
}, func(this, k, notFound interface{}) interface{} {
	return this.(ILookup).Lookup_Arity3(k, notFound)
})

func init() {
	protocols["cljs.core/ILookup"] = reflect.TypeOf((*ILookup)(nil)).Elem()
}

// This will likley be generated / overriding the gen-apply-to stuff.
var Apply = Fn(func(f_args ...interface{}) interface{} {
	f, args := f_args[0], f_args[1:]
	argc := len(args)
	if argc < 1 {
		throwArity(nil, argc)
	}
	var spread = args[argc-1].([]interface{}) // This will be a seq in real life.
	return f.(*AFn).Call(append(args[:argc-1], spread...)...)
})

type CljsCoreSymbol struct {
	ns, name, str, _hash, _meta interface{}
	AbstractIFn
}

func (this *CljsCoreSymbol) Name_Arity1() string {
	return this.name.(string)
}

func (this *CljsCoreSymbol) Namespace_Arity1() string {
	return this.ns.(string)
}

func (this *CljsCoreSymbol) String() string {
	return this.str.(string)
}

func (this *CljsCoreSymbol) Invoke_Arity1(coll interface{}) interface{} {
	return Lookup_.Invoke_Arity2(coll, this)
}

func (this *CljsCoreSymbol) Invoke_Arity2(coll, notFound interface{}) interface{} {
	return Lookup_.Invoke_Arity3(coll, this, notFound)
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
		return &CljsCoreSymbol{ns: ns, name: name, str: symStr}
	})
}(&AFn{})

var Implements_QMARK_ = NativeSatisifes_QMARK_

var String_QMARK_ = Fn(func(x interface{}) interface{} {
	return reflect.ValueOf(x).Kind() == reflect.String
})

var Namespace = Fn(func(x interface{}) interface{} {
	if Truth_(Implements_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "INamed"), x)) {
		return Namespace_.Invoke_Arity1(x)
	} else {
		panic(&js.Error{Str.Invoke_ArityVariadic("Doesn't support namespace: ", x)})
	}
})

var Name = Fn(func(x interface{}) interface{} {
	if Truth_(Implements_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "INamed"), x)) {
		return Name_.Invoke_Arity1(x)
	} else {
		if Truth_(String_QMARK_.Invoke_Arity1(x)) {
			return x
		} else {
			panic(&js.Error{Str.Invoke_ArityVariadic("Doesn't support name: ", x)})
		}
	}
})

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
				sb = NativeInvokeInstanceMethod.Invoke_Arity3(sb, "Append",
					[]interface{}{Str.Invoke_Arity1(First.Invoke_Arity1(more))})
				more = Next.Invoke_Arity1(more)
				continue
			} else {
				return fmt.Sprint(sb)
			}
		}
	})
}(&AFn{})

// cljs.reflect / clojure.lang.Reflector
var NativeGetInstanceField = Fn(func(target, fieldName interface{}) interface{} {
	tv := reflect.ValueOf(target)
	if tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}
	return tv.FieldByName(fieldName.(string)).Interface()
})

var NativeSetInstanceField = Fn(func(target, fieldName, val interface{}) interface{} {
	tv := reflect.ValueOf(target)
	if tv.Kind() == reflect.Ptr {
		tv = tv.Elem()
	}
	tv.FieldByName(fieldName.(string)).Set(reflect.ValueOf(val))
	return val
})

var NativeInvokeFunc = Fn(func(f, args interface{}) interface{} {
	var fv reflect.Value
	switch f.(type) {
	case reflect.Value:
		fv = f.(reflect.Value)
	default:
		fv = reflect.ValueOf(f)
	}
	in := make([]reflect.Value, len(args.([]interface{})))
	for i, a := range args.([]interface{}) {
		in[i] = reflect.ValueOf(a)
	}
	return fv.Call(in)[0].Interface()
})

var NativeInvokeInstanceMethod = Fn(func(target, methodName, args interface{}) interface{} {
	return NativeInvokeFunc.Invoke_Arity2(reflect.ValueOf(target).MethodByName(methodName.(string)), args)
})

// cljs.rt

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

type Arity0B func() bool
type Arity1IB func(interface{}) bool
type Arity2IIB func(_, _ interface{}) bool

// CLJS also (among other things) adds .call and .apply when implementing the IFn protocol, see cljs.core/add-ifn-methods, clj
// The easiest way to acheive this is renaming (and hide) the fields, and make CljsCoreIFn_InvokeArity1 an interface method.
// Then there's the issue of other protocols and how to represent them, as we prefer to keep them as Go interfaces.
// As can be seen above regarding IFn, it would be nice to ensure that IFn is just a special case that emit* :invoke knows about.
// IFn -invoke is like any other protocol in JS. There's a dispatch fn setup which looks for an implementation on the receiver.
// That is, cljs.core._invoke will call  receiver.cljs$core$IFn$_invoke$arity$1(receiver, x)
// My impression is that this is how CLJS does it. To reiterate - the reason this becomes extra messy in Go is lack of overloading.
type AFn struct {
	MaxFixedArity int
	ArityVariadic
	Arity0
	Arity1
	Arity2
	Arity3
	Arity4
	Arity5
}

type AFnPrimtive struct {
	AFn
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
	Arity0B
	Arity1IB
	Arity2IIB
}

func throwArity(f, arity interface{}) interface{} {
	if f == nil || reflect.ValueOf(f).IsNil() {
		panic(&js.Error{fmt.Sprint("Invalid arity: ", arity)})
	}
	return f
}

func (this *AFn) IsVariadic() bool {
	return this.MaxFixedArity >= 0 && this.ArityVariadic != nil
}

func (this *AFn) Call(args ...interface{}) interface{} {
	if this == Invoke_ {
		return args[0].(*AFn).Call(args[1:]...)
	}
	argc := len(args)
	if argc > this.MaxFixedArity && this.IsVariadic() {
		return this.Invoke_ArityVariadic(args...)
	}
	switch argc {
	case 0:
		return this.Invoke_Arity0()
	case 1:
		return this.Invoke_Arity1(args[0])
	case 2:
		return this.Invoke_Arity2(args[0], args[1])
	case 3:
		return this.Invoke_Arity3(args[0], args[1], args[2])
	case 4:
		return this.Invoke_Arity4(args[0], args[1], args[2], args[3])
	case 5:
		return this.Invoke_Arity5(args[0], args[1], args[2], args[3], args[4])
	}
	return throwArity(nil, argc)
}

func (this *AFnPrimtive) Call(args ...interface{}) interface{} {
	switch len(args) {
	case 0:
		switch {
		case this.Arity0F != nil:
			return this.Arity0F()
		}
	case 1:
		switch {
		case this.Arity1IF != nil:
			return this.Arity1IF(args[0])
		case this.Arity1FI != nil:
			return this.Arity1FI(args[0].(float64))
		case this.Arity1FF != nil:
			return this.Arity1FF(args[0].(float64))
		}
	case 2:
		switch {
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
	}
	return this.AFn.Call(args...)
}

func (this ArityVariadic) Invoke_ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
	return this(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest...)
}

func (this Arity0) Invoke_Arity0() interface{} {
	return this()
}

func (this Arity1) Invoke_Arity1(a interface{}) interface{} {
	return this(a)
}

func (this Arity2) Invoke_Arity2(a, b interface{}) interface{} {
	return this(a, b)
}

func (this Arity3) Invoke_Arity3(a, b, c interface{}) interface{} {
	return this(a, b, c)
}

func (this Arity4) Invoke_Arity4(a, b, c, d interface{}) interface{} {
	return this(a, b, c, d)
}

func (this Arity5) Invoke_Arity5(a, b, c, d, e interface{}) interface{} {
	return this(a, b, c, d, e)
}

type AbstractIFn struct{}

func (this AbstractIFn) Invoke_ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
	return throwArity(nil, len(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest))
}

func (this AbstractIFn) Invoke_Arity0() interface{} {
	return throwArity(nil, 0)
}

func (this AbstractIFn) Invoke_Arity1(a interface{}) interface{} {
	return throwArity(nil, 1)
}

func (this AbstractIFn) Invoke_Arity2(a, b interface{}) interface{} {
	return throwArity(nil, 2)
}

func (this AbstractIFn) Invoke_Arity3(a, b, c interface{}) interface{} {
	return throwArity(nil, 3)
}

func (this AbstractIFn) Invoke_Arity4(a, b, c, d interface{}) interface{} {
	return throwArity(nil, 4)
}

func (this AbstractIFn) Invoke_Arity5(a, b, c, d, e interface{}) interface{} {
	return throwArity(nil, 4)
}

var type2sig = map[string]string{"interface": "I", "slice": "I", "float64": "F", "bool": "B"}

func primtiveSignature(t reflect.Type) string {
	sig := ""
	for i := 0; i < t.NumIn(); i++ {
		sig += type2sig[t.In(i).Kind().String()]
	}
	for i := 0; i < t.NumOut(); i++ {
		sig += type2sig[t.Out(i).Kind().String()]
	}
	if regexp.MustCompile("I+").MatchString(sig) {
		sig = ""
	}
	return sig
}

func makeBridge(f reflect.Value, from reflect.Type) reflect.Value {
	return reflect.MakeFunc(from, func(in []reflect.Value) []reflect.Value {
		for i, v := range in {
			in[i] = reflect.ValueOf(v.Interface())
		}
		out := f.Call(in)
		for i, v := range out {
			switch v.Kind() {
			case reflect.Float64:
				var o interface{} = v.Float()
				out[i] = reflect.ValueOf(&o).Elem()
			}
		}
		return out
	})
}

func makeInvalidArity(from reflect.Type) reflect.Value {
	return reflect.MakeFunc(from, func(in []reflect.Value) []reflect.Value {
		throwArity(nil, len(in))
		return nil
	})
}

func Fn(fns ...interface{}) IFn {
	var f *AFn = &AFn{}
	var fp *AFnPrimtive
	var vp reflect.Value
	if len(fns) > 0 {
		switch reflect.ValueOf(fns[0]).Type() {
		case reflect.TypeOf(f):
			f = fns[0].(*AFn)
			fns = fns[1:]
		case reflect.TypeOf(fp):
			fp = fns[0].(*AFnPrimtive)
			f = &fp.AFn
			fns = fns[1:]
			vp = reflect.ValueOf(fp).Elem()
		}
	}
	v := reflect.ValueOf(f).Elem()

	variadic := false
	maxFixedArity := 0
	for _, a := range fns {
		av := reflect.ValueOf(a)
		at := av.Type()

		if at.IsVariadic() {
			variadic = true
			f.ArityVariadic = a.(func(...interface{}) interface{})
		} else {
			if maxFixedArity < at.NumIn() {
				maxFixedArity = at.NumIn()
			}
			af := v.FieldByName(fmt.Sprint("Arity", at.NumIn()))
			if sig := primtiveSignature(at); sig != "" {
				vp.FieldByName(fmt.Sprintf("Arity%d%s", at.NumIn(), sig)).Set(av)
				af.Set(makeBridge(av, af.Type()))
			} else {
				af.Set(av)
			}
		}
	}
	if variadic {
		f.MaxFixedArity = maxFixedArity
	} else {
		f.MaxFixedArity = -1
	}
	vt := v.Type()
	for i := 0; i < vt.NumField(); i++ {
		vf := v.Field(i)
		if vf.Kind() == reflect.Func && vf.IsNil() {
			vf.Set(makeInvalidArity(vf.Type()))
		}
	}
	if fp != nil {
		return fp
	} else {
		return f
	}
}

func Truth_(x interface{}) bool {
	return x != nil && x != false
}

func Main() {
	Enable_console_print_BANG_.Invoke_Arity0()
	args := make([]interface{}, len(os.Args[1:]))
	for i, a := range os.Args[1:] {
		args[i] = a
	}
	STAR_main_cli_fn_STAR_.(IFn).Invoke_ArityVariadic(args...)
}
