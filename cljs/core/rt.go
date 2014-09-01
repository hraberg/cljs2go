package core

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/hraberg/cljs.go/js"
)

func element(x interface{}) reflect.Value {
	if e := reflect.ValueOf(x); e.Kind() == reflect.Ptr {
		return e.Elem()
	} else {
		return e
	}
}

var Native_get_instance_field = Fn(func(target, fieldName interface{}) interface{} {
	return element(target).FieldByName(fieldName.(string)).Interface()
})

var Native_set_instance_field = Fn(func(target, fieldName, val interface{}) interface{} {
	element(target).FieldByName(fieldName.(string)).Set(reflect.ValueOf(val))
	return val
})

func value(x interface{}) reflect.Value {
	if v, ok := x.(reflect.Value); ok {
		return v
	}
	return reflect.ValueOf(x)
}

var Native_invoke_func = Fn(func(f, args interface{}) interface{} {
	in := make([]reflect.Value, len(args.([]interface{})))
	for i, a := range args.([]interface{}) {
		in[i] = reflect.ValueOf(a)
	}
	return value(f).Call(in)[0].Interface()
})

func decorate(target interface{}) interface{} {
	switch object := target.(type) {
	case string:
		return js.JSString_(object)
	case []interface{}:
		return js.JSArray_(&object)
	case map[string]interface{}:
		return js.JSObject(object)
	case nil:
		return js.JSNil(object)
	default:
		return object
	}
}

var Native_invoke_instance_method = Fn(func(target, methodName, args interface{}) interface{} {
	return Native_invoke_func.X_invoke_Arity2(reflect.ValueOf(decorate(target)).MethodByName(methodName.(string)), args)
})

var protocols = map[string]reflect.Type{}

func RegisterProtocol_(name string, p interface{}) {
	protocols[name] = reflect.TypeOf(p).Elem()
}

type ArityVariadic func(...interface{}) interface{}
type Arity0 func() interface{}
type Arity1 func(_ interface{}) interface{}
type Arity2 func(_, _ interface{}) interface{}
type Arity3 func(_, _, _ interface{}) interface{}
type Arity4 func(_, _, _, _ interface{}) interface{}
type Arity5 func(_, _, _, _, _ interface{}) interface{}
type Arity6 func(_, _, _, _, _, _ interface{}) interface{}
type Arity7 func(_, _, _, _, _, _, _ interface{}) interface{}
type Arity8 func(_, _, _, _, _, _, _, _ interface{}) interface{}
type Arity9 func(_, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity10 func(_, _, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity11 func(_, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity12 func(_, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity13 func(_, _, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity14 func(_, _, _, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity15 func(_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity16 func(_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity17 func(_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity18 func(_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity19 func(_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}
type Arity20 func(_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}

/*
 Primitive signatures used by ClojureScript itself, but it only ever type hints the return tag.
Arity0F
Arity1IA
Arity1IB
Arity1IF
Arity2IIB
Arity2IIF

 We might want to propagate :cljs.analyzer/numeric from the macros to the fns for arity 1 and 2.
 These would be:
 Arity1FF
 Arity2FFF
 Arity1FB
 Arity2FFB

 The full list of tags used for returns on protocols are:
 ^array
 ^boolean
 ^clj
 ^clj-or-nil
 ^not-native
 ^number
 ^seq
 ^string

 ^seq is also used for normal fns, we can add ClsjCoreISeq, with short type Q:
 Arity1IQ
*/

type Arity1IA func(interface{}) []interface{}

type Arity1IQ func(interface{}) CljsCoreISeq

type Arity0F func() float64
type Arity1IF func(interface{}) float64
type Arity1FI func(float64) interface{}
type Arity1FF func(float64) float64
type Arity2IIF func(_, _ interface{}) float64
type Arity2IFI func(interface{}, float64) interface{}
type Arity2IFF func(interface{}, float64) float64
type Arity2FII func(float64, interface{}) interface{}
type Arity2FIF func(float64, interface{}) float64
type Arity2FFI func(_, _ float64) interface{}
type Arity2FFF func(_, _ float64) float64

type Arity0B func() bool
type Arity1IB func(interface{}) bool
type Arity1FB func(float64) bool
type Arity2IIB func(_, _ interface{}) bool
type Arity2FFB func(_, _ float64) bool

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
	Arity6
	Arity7
	Arity8
	Arity9
	Arity10
	Arity11
	Arity12
	Arity13
	Arity14
	Arity15
	Arity16
	Arity17
	Arity18
	Arity19
	Arity20

	Arity0F
	Arity1IF
	Arity1FI
	Arity1FF
	Arity2IIF
	Arity2IFI
	Arity2IFF
	Arity2FII
	Arity2FIF
	Arity2FFI
	Arity2FFF

	Arity0B
	Arity1FB
	Arity1IB
	Arity2IIB
	Arity2FFB

	Arity1IA
	Arity1IQ
}

func throwArity(f, arity interface{}) interface{} {
	if f == nil || reflect.ValueOf(f).IsNil() {
		panic(&js.Error{fmt.Sprint("Invalid arity: ", arity)})
	}
	return f
}

func (this *AFn) isVariadic() bool {
	return this.MaxFixedArity >= 0 && this.ArityVariadic != nil
}

func (this *AFn) Call(args ...interface{}) interface{} {
	if this == X_invoke {
		return args[0].(*AFn).Call(args[1:]...)
	}
	argc := len(args)
	if argc > this.MaxFixedArity && this.isVariadic() {
		return this.X_invoke_ArityVariadic(args...)
	}
	switch argc {
	case 0:
		switch {
		case this.Arity0F != nil:
			return this.Arity0F()
		case this.Arity0B != nil:
			return this.Arity0B()
		default:
			return this.Arity0()
		}
	case 1:
		switch {
		case this.Arity1IF != nil:
			return this.Arity1IF(args[0])
		case this.Arity1FI != nil:
			return this.Arity1FI(args[0].(float64))
		case this.Arity1FF != nil:
			return this.Arity1FF(args[0].(float64))
		case this.Arity1IB != nil:
			return this.Arity1IB(args[0])
		case this.Arity1FB != nil:
			return this.Arity1FB(args[0].(float64))
		case this.Arity1IA != nil:
			return this.Arity1IA(args[0])
		case this.Arity1IQ != nil:
			return this.Arity1IQ(args[0])
		default:
			return this.Arity1(args[0])
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
		case this.Arity2IIB != nil:
			return this.Arity2IIB(args[0], args[1])
		case this.Arity2FFB != nil:
			return this.Arity2FFB(args[0].(float64), args[1].(float64))
		default:
			return this.Arity2(args[0], args[1])
		}
	case 3:
		return this.Arity3(args[0], args[1], args[2])
	case 4:
		return this.Arity4(args[0], args[1], args[2], args[3])
	case 5:
		return this.Arity5(args[0], args[1], args[2], args[3], args[4])
	case 6:
		return this.Arity6(args[0], args[1], args[2], args[3], args[4], args[5])
	case 7:
		return this.Arity7(args[0], args[1], args[2], args[3], args[4], args[5], args[6])
	case 8:
		return this.Arity8(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
	case 9:
		return this.Arity9(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8])
	case 10:
		return this.Arity10(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9])
	case 11:
		return this.Arity11(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10])
	case 12:
		return this.Arity12(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11])
	case 13:
		return this.Arity13(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12])
	case 14:
		return this.Arity14(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13])
	case 15:
		return this.Arity15(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14])
	case 16:
		return this.Arity16(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15])
	case 17:
		return this.Arity17(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16])
	case 18:
		return this.Arity18(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17])
	case 19:
		return this.Arity19(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17], args[18])
	case 20:
		return this.Arity20(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17], args[18], args[19])
	}
	return throwArity(nil, argc)
}

func (this Arity0) X_invoke_Arity0() interface{} {
	return this()
}

func (this Arity1) X_invoke_Arity1(a interface{}) interface{} {
	return this(a)
}

func (this Arity2) X_invoke_Arity2(a, b interface{}) interface{} {
	return this(a, b)
}

func (this Arity3) X_invoke_Arity3(a, b, c interface{}) interface{} {
	return this(a, b, c)
}

func (this Arity4) X_invoke_Arity4(a, b, c, d interface{}) interface{} {
	return this(a, b, c, d)
}

func (this Arity5) X_invoke_Arity5(a, b, c, d, e interface{}) interface{} {
	return this(a, b, c, d, e)
}

func (this Arity6) X_invoke_Arity6(a, b, c, d, e, f interface{}) interface{} {
	return this(a, b, c, d, e, f)
}

func (this Arity7) X_invoke_Arity7(a, b, c, d, e, f, g interface{}) interface{} {
	return this(a, b, c, d, e, f, g)
}

func (this Arity8) X_invoke_Arity8(a, b, c, d, e, f, g, h interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h)
}

func (this Arity9) X_invoke_Arity9(a, b, c, d, e, f, g, h, i interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i)
}

func (this Arity10) X_invoke_Arity10(a, b, c, d, e, f, g, h, i, j interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j)
}

func (this Arity11) X_invoke_Arity11(a, b, c, d, e, f, g, h, i, j, k interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j, k)
}

func (this Arity12) X_invoke_Arity12(a, b, c, d, e, f, g, h, i, j, k, l interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j, k, l)
}

func (this Arity13) X_invoke_Arity13(a, b, c, d, e, f, g, h, i, j, k, l, m interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j, k, l, m)
}

func (this Arity14) X_invoke_Arity14(a, b, c, d, e, f, g, h, i, j, k, l, m, n interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j, k, l, m, n)
}

func (this Arity15) X_invoke_Arity15(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
}

func (this Arity16) X_invoke_Arity16(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
}

func (this Arity17) X_invoke_Arity17(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
}

// IFn skips r
func (this Arity18) X_invoke_Arity18(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s)
}

func (this Arity19) X_invoke_Arity19(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s, t interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s, t)
}

func (this Arity20) X_invoke_Arity20(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s, t, rest interface{}) interface{} {
	return this(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s, t, rest)
}

func (this ArityVariadic) X_invoke_ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
	return this(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest...)
}

var type2sig = map[string]string{
	"interface {}":      "I",
	"[]interface {}":    "A",
	"float64":           "F",
	"bool":              "B",
	"string":            "S",
	"core.CljsCoreISeq": "Q"}

func typedSignature(t reflect.Type) string {
	sig := ""
	if t.IsVariadic() {
		return sig
	}
	for i := 0; i < t.NumIn(); i++ {
		sig += type2sig[t.In(i).String()]
	}
	for i := 0; i < t.NumOut(); i++ {
		sig += type2sig[t.Out(i).String()]
	}
	if strings.Replace(sig, "I", "", -1) == "" {
		return ""
	}
	return sig
}

func makeTypedBridge(f reflect.Value, from reflect.Type) reflect.Value {
	return reflect.MakeFunc(from, func(in []reflect.Value) []reflect.Value {
		for i, v := range in {
			in[i] = reflect.ValueOf(v.Interface())
		}
		out := f.Call(in)
		for i, v := range out {
			var o interface{}
			switch v.Kind() {
			case reflect.Float64:
				o = v.Float()
				out[i] = reflect.ValueOf(&o).Elem()
			case reflect.Bool:
				o = v.Bool()
				out[i] = reflect.ValueOf(&o).Elem()
			}
		}
		return out
	})
}

func makeVarargsBridge(varargs reflect.Value, from reflect.Type) reflect.Value {
	return reflect.MakeFunc(from, func(in []reflect.Value) []reflect.Value {
		return varargs.Call(in)
	})
}

func makeInvalidArity(from reflect.Type) reflect.Value {
	return reflect.MakeFunc(from, func(in []reflect.Value) []reflect.Value {
		throwArity(nil, len(in))
		return nil
	})
}

func Fn(fns ...interface{}) *AFn {
	var f *AFn = &AFn{}
	if len(fns) > 0 {
		switch reflect.ValueOf(fns[0]).Type() {
		case reflect.TypeOf(f):
			f = fns[0].(*AFn)
			fns = fns[1:]
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
			if sig := typedSignature(at); sig != "" {
				v.FieldByName(fmt.Sprintf("Arity%d%s", at.NumIn(), sig)).Set(av)
				af.Set(makeTypedBridge(av, af.Type()))
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
	for i := 0; i < v.Type().NumField(); i++ {
		vf := v.Field(i)
		vt := vf.Type()
		if vf.Kind() == reflect.Func && vf.IsNil() {
			if variadic && vt.NumIn() > maxFixedArity {
				vf.Set(makeVarargsBridge(reflect.ValueOf(f.ArityVariadic), vf.Type()))
			} else if sig := typedSignature(vt); sig == "" {
				vf.Set(makeInvalidArity(vf.Type()))
			}
		}
	}
	return f
}

func Truth_(x interface{}) bool {
	return x != nil && x != false
}

type Object interface {
	ToString() string
	Equiv(other interface{}) bool
}

var Apply = Fn(func(f_args ...interface{}) interface{} {
	f, args := f_args[0], f_args[1:]
	argc := len(args)
	if argc < 1 {
		throwArity(nil, argc)
	}
	var spread = args[argc-1].([]interface{}) // This will be a seq in real life.
	return f.(*AFn).Call(append(args[:argc-1], spread...)...)
})

var Native_satisfies_QMARK_ = Fn(func(p, x interface{}) bool {
	return reflect.ValueOf(x).Type().Implements(protocols[fmt.Sprint(p)])
})

func init() {
	RegisterProtocol_("cljs.core/Object", (*Object)(nil))

	X_invoke.ArityVariadic = func(f_args ...interface{}) interface{} {
		f, args := f_args[0], f_args[1:]
		if f, ok := f.(*AFn); ok {
			return f.X_invoke_ArityVariadic(args...)
		}
		return throwArity(nil, len(args)-1)
	}
}

func Main_() {
	Enable_console_print_BANG_.X_invoke_Arity0()
	args := make([]interface{}, len(os.Args[1:]))
	for i, a := range os.Args[1:] {
		args[i] = a
	}
	X_STAR_main_cli_fn_STAR_.(*AFn).X_invoke_ArityVariadic(args...)
}
