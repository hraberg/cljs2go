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

var NativeGetInstanceField = Fn(func(target, fieldName interface{}) interface{} {
	return element(target).FieldByName(fieldName.(string)).Interface()
})

var NativeSetInstanceField = Fn(func(target, fieldName, val interface{}) interface{} {
	element(target).FieldByName(fieldName.(string)).Set(reflect.ValueOf(val))
	return val
})

func value(x interface{}) reflect.Value {
	if v, ok := x.(reflect.Value); ok {
		return v
	}
	return reflect.ValueOf(x)
}

var NativeInvokeFunc = Fn(func(f, args interface{}) interface{} {
	in := make([]reflect.Value, len(args.([]interface{})))
	for i, a := range args.([]interface{}) {
		in[i] = reflect.ValueOf(a)
	}
	return value(f).Call(in)[0].Interface()
})

func decorate(target interface{}) interface{} {
	switch object := target.(type) {
	case string:
		return js.JSString(object)
	case []interface{}:
		return js.JSArray(object)
	case map[string]interface{}:
		return js.JSObject(object)
	default:
		return object
	}
}

var NativeInvokeInstanceMethod = Fn(func(target, methodName, args interface{}) interface{} {
	return NativeInvokeFunc.Invoke_Arity2(reflect.ValueOf(decorate(target)).MethodByName(methodName.(string)), args)
})

var protocols = map[string]reflect.Type{}

func RegisterProtocol(name string, p interface{}) {
	protocols[name] = reflect.TypeOf(p).Elem()
}

type ArityVariadic func(...interface{}) interface{}
type Arity0 func() interface{}
type Arity1 func(interface{}) interface{}
type Arity2 func(_, _ interface{}) interface{}
type Arity3 func(_, _, _ interface{}) interface{}
type Arity4 func(_, _, _, _ interface{}) interface{}
type Arity5 func(_, _, _, _, _ interface{}) interface{}

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
type Arity1BI func(bool) interface{}
type Arity1BB func(bool) bool
type Arity2IIB func(_, _ interface{}) bool
type Arity2IBI func(interface{}, bool) interface{}
type Arity2IBB func(interface{}, bool) bool
type Arity2BII func(bool, interface{}) interface{}
type Arity2BIB func(bool, interface{}) bool
type Arity2BBI func(_, _ bool) interface{}
type Arity2BBB func(_, _ bool) bool

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
	Arity1IB
	Arity1BI
	Arity1BB
	Arity2IIB
	Arity2IBI
	Arity2IBB
	Arity2BII
	Arity2BIB
	Arity2BBI
	Arity2BBB
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
	if this == X_Invoke {
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

var type2sig = map[reflect.Kind]string{reflect.Interface: "I", reflect.Slice: "I", reflect.Float64: "F", reflect.Bool: "B"}

func primtiveSignature(t reflect.Type) string {
	sig := ""
	for i := 0; i < t.NumIn(); i++ {
		sig += type2sig[t.In(i).Kind()]
	}
	for i := 0; i < t.NumOut(); i++ {
		sig += type2sig[t.Out(i).Kind()]
	}
	if strings.Replace(sig, "I", "", -1) == "" {
		return ""
	}
	return sig
}

func makePrimtiveBridge(f reflect.Value, from reflect.Type) reflect.Value {
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
				af.Set(makePrimtiveBridge(av, af.Type()))
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
		if vf.Kind() == reflect.Func && vf.IsNil() {
			if variadic && vf.Type().NumIn() > maxFixedArity {
				vf.Set(makeVarargsBridge(reflect.ValueOf(f.ArityVariadic), vf.Type()))
			} else {
				vf.Set(makeInvalidArity(vf.Type()))
			}
		}
	}
	if fp != nil {
		return fp
	}
	return f
}

func Truth_(x interface{}) bool {
	return x != nil && x != false
}

func init() {
	RegisterProtocol("js/Object", (*js.Object)(nil))

	Apply = Fn(func(f_args ...interface{}) interface{} {
		f, args := f_args[0], f_args[1:]
		argc := len(args)
		if argc < 1 {
			throwArity(nil, argc)
		}
		var spread = args[argc-1].([]interface{}) // This will be a seq in real life.
		if fp, ok := f.(*AFnPrimtive); ok {
			return fp.AFn.Call(append(args[:argc-1], spread...)...)
		}
		return f.(*AFn).Call(append(args[:argc-1], spread...)...)
	})

	NativeSatisifes_QMARK_ = Fn(&AFnPrimtive{}, func(p, x interface{}) bool {
		return reflect.ValueOf(x).Type().Implements(protocols[fmt.Sprint(p)])
	}).(*AFnPrimtive)
	Implements_QMARK_ = NativeSatisifes_QMARK_
}

func Main() {
	Enable_console_print_BANG_.Invoke_Arity0()
	args := make([]interface{}, len(os.Args[1:]))
	for i, a := range os.Args[1:] {
		args[i] = a
	}
	X_STAR_main_cli_fn_STAR_.(IFn).Invoke_ArityVariadic(args...)
}
