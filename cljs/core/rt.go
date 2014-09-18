package core

import (
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/hraberg/cljs.go/js"
)

func element(x interface{}) reflect.Value {
	if e := reflect.ValueOf(x); e.Kind() == reflect.Ptr {
		return e.Elem()
	} else {
		return e
	}
}

func Value_(x interface{}) reflect.Value {
	if v, ok := x.(reflect.Value); ok {
		return v
	}
	if x == nil {
		return reflect.ValueOf(&x).Elem()
	}
	return reflect.ValueOf(x)
}

func decorate(target interface{}) interface{} {
	switch object := target.(type) {
	case string:
		return js.JSString_(object)
	case []interface{}:
		return js.JSArray_(&object)
	case map[string]interface{}:
		return js.JSObject(object)
	case nil:
		return js.JSNil{}
	case bool:
		return js.JSBoolean(object)
	case float64:
		return js.JSNumber(object)
	default:
		return object
	}
}

var Native_get_instance_field = Fn(func(target, fieldName interface{}) interface{} {
	return element(decorate(target)).FieldByName(fieldName.(string)).Interface()
})

var Native_set_instance_field = Fn(func(target, fieldName, val interface{}) interface{} {
	element(decorate(target)).FieldByName(fieldName.(string)).Set(Value_(val))
	return val
})

var Native_invoke_func = Fn(func(f, args interface{}) interface{} {
	fv := Value_(f)
	argsArray := args.([]interface{}) // this should really take a seq
	argc := fv.Type().NumIn()
	if len(argsArray) > argc {
		argc = len(argsArray)
	}
	in := make([]reflect.Value, argc)
	for i, v := range argsArray {
		in[i] = Value_(v)
	}
	for i := len(argsArray); i < len(in); i++ {
		in[i] = Value_(nil)
	}
	return fv.Call(in)[0].Interface()
})

var Native_invoke_instance_method = Fn(func(target, methodName, args interface{}) interface{} {
	return Native_invoke_func.X_invoke_Arity2(reflect.ValueOf(decorate(target)).MethodByName(methodName.(string)), args)
})

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

// Unsupported
type Arity21 func(_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{}

type Arity1IA func(interface{}) []interface{}
type Arity2IIA func(_, _ interface{}) []interface{}

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
type Arity1BI func(bool) interface{}
type Arity1FB func(float64) bool
type Arity2IIB func(_, _ interface{}) bool
type Arity2IBI func(_ interface{}, _ bool) interface{}
type Arity2FFB func(_, _ float64) bool
type Arity3IIIB func(_, _, _ interface{}) bool
type Arity3IBBI func(_ interface{}, _, _ bool) interface{}
type Arity3IIBI func(_, _ interface{}, _ bool) interface{}

type Arity5IIBIII func(_, _ interface{}, _ bool, _, _ interface{}) interface{}
type Arity5BIIBII func(_ bool, _, _ interface{}, _ bool, _ interface{}) interface{}
type Arity6IIIBIII func(_, _, _ interface{}, _ bool, _, _ interface{}) interface{}

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
	Arity21

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
	Arity1BI

	Arity2IIB
	Arity2IBI
	Arity2FFB

	Arity3IIIB
	Arity3IBBI
	Arity3IIBI

	Arity1IA
	Arity2IIA

	Arity1IQ

	// For positonal factories, needs better solution:
	Arity5IIBIII
	Arity5BIIBII
	Arity6IIIBIII

	Meta CljsCoreIMap
}

func throwArity(f, arity interface{}) interface{} {
	if f == nil || reflect.ValueOf(f).IsNil() {
		panic(&js.Error{fmt.Sprint("Invalid arity: ", arity)})
	}
	return f
}

func (this *AFn) CljsCoreIFn__() {
}

func (this *AFn) CljsCoreFn__() {
}

func (_ *AFn) CljsCoreICloneable__() {}

func (this *AFn) X_clone_Arity1() interface{} {
	clone := *this
	return &clone
}

func (_ *AFn) CljsCoreIMeta__() {}

func (this *AFn) X_meta_Arity1() interface{} {
	return this.Meta
}

func (_ *AFn) CljsCoreIWithMeta__() {}

func (this *AFn) X_with_meta_Arity2(meta interface{}) interface{} {
	clone := this.X_clone_Arity1().(*AFn)
	clone.Meta = meta.(CljsCoreIMap)
	return clone
}

func (this *AFn) isVariadic() bool {
	return this.MaxFixedArity >= 0 && this.ArityVariadic != nil
}

func (this *AFn) Call(args ...interface{}) interface{} {
	if this == X_invoke {
		return Call_(args[0].(CljsCoreIFn), args[1:]...)
	}
	argc := len(args)
	if argc > this.MaxFixedArity && this.isVariadic() {
		return this.X_invoke_ArityVariadic(append(args, Array_seq.X_invoke_Arity1([]interface{}{}))...)
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
		case this.Arity1BI != nil:
			return this.Arity1IB(args[0].(bool))
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
		case this.Arity2IBI != nil:
			return this.Arity2IBI(args[0], args[1].(bool))
		case this.Arity2FFB != nil:
			return this.Arity2FFB(args[0].(float64), args[1].(float64))
		case this.Arity2IIA != nil:
			return this.Arity2IIA(args[0], args[1])
		default:
			return this.Arity2(args[0], args[1])
		}
	case 3:
		switch {
		case this.Arity3IIIB != nil:
			return this.Arity3IIIB(args[0], args[1], args[2])
		case this.Arity3IBBI != nil:
			return this.Arity3IBBI(args[0], args[1].(bool), args[2].(bool))
		case this.Arity3IIBI != nil:
			return this.Arity3IIBI(args[0], args[1], args[2].(bool))
		default:
			return this.Arity3(args[0], args[1], args[2])
		}
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

func Call_(this CljsCoreIFn, args ...interface{}) interface{} {
	if afn, ok := this.(*AFn); ok {
		return afn.Call(args...)
	}
	argc := len(args)
	switch argc {
	case 0:
		return this.X_invoke_Arity0()
	case 1:
		return this.X_invoke_Arity1(args[0])
	case 2:
		return this.X_invoke_Arity2(args[0], args[1])
	case 3:
		return this.X_invoke_Arity3(args[0], args[1], args[2])
	case 4:
		return this.X_invoke_Arity4(args[0], args[1], args[2], args[3])
	case 5:
		return this.X_invoke_Arity5(args[0], args[1], args[2], args[3], args[4])
	case 6:
		return this.X_invoke_Arity6(args[0], args[1], args[2], args[3], args[4], args[5])
	case 7:
		return this.X_invoke_Arity7(args[0], args[1], args[2], args[3], args[4], args[5], args[6])
	case 8:
		return this.X_invoke_Arity8(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7])
	case 9:
		return this.X_invoke_Arity9(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8])
	case 10:
		return this.X_invoke_Arity10(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9])
	case 11:
		return this.X_invoke_Arity11(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10])
	case 12:
		return this.X_invoke_Arity12(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11])
	case 13:
		return this.X_invoke_Arity13(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12])
	case 14:
		return this.X_invoke_Arity14(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13])
	case 15:
		return this.X_invoke_Arity15(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14])
	case 16:
		return this.X_invoke_Arity16(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15])
	case 17:
		return this.X_invoke_Arity17(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16])
	case 18:
		return this.X_invoke_Arity18(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17])
	case 19:
		return this.X_invoke_Arity19(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17], args[18])
	case 20:
		return this.X_invoke_Arity20(args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12], args[13], args[14], args[15], args[16], args[17], args[18], args[19])
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

func (this Arity21) X_invoke_Arity21(_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ interface{}) interface{} {
	panic(&js.Error{"Only up to 20 arguments supported on functions."})
}

func (this ArityVariadic) X_invoke_ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
	return this(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest...)
}

var type2sig = map[string]rune{
	"interface {}":      'I',
	"[]interface {}":    'A',
	"float64":           'F',
	"bool":              'B',
	"core.CljsCoreISeq": 'Q'}

var noPrimitivesRegexp = regexp.MustCompile("^I+$")

func typedSignature_(t reflect.Type) string {
	if t.IsVariadic() {
		return ""
	}
	in, out := t.NumIn(), t.NumOut()
	sig := make([]rune, in+out)
	for i := 0; i < in; i++ {
		sig[i] = type2sig[t.In(i).String()]
	}
	for i := 0; i < out; i++ {
		sig[i+in] = type2sig[t.Out(i).String()]
	}
	if noPrimitivesRegexp.MatchString(string(sig)) {
		return ""
	} else {
		return string(sig)
	}
}

var signaturesByType = map[string]string{}

func typedSignature(t reflect.Type) string {
	key := t.String()
	if sig, exists := signaturesByType[key]; exists {
		return sig
	} else {
		sig = typedSignature_(t)
		signaturesByType[key] = sig
		return sig
	}
}

func makeTypedBridge(f reflect.Value, from reflect.Type) reflect.Value {
	return reflect.MakeFunc(from, func(in []reflect.Value) []reflect.Value {
		for i, v := range in {
			in[i] = Value_(v.Interface())
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
			case reflect.String:
				o = v.String()
				out[i] = reflect.ValueOf(&o).Elem()
			default:
				o = v.Interface()
				out[i] = reflect.ValueOf(&o).Elem()
			}
		}
		return out
	})
}

func makeVarargsBridge(maxFixedArity int, varargs reflect.Value, from reflect.Type) reflect.Value {
	return reflect.MakeFunc(from, func(in []reflect.Value) []reflect.Value {
		fixedArgs, restArgs := in[:maxFixedArity], in[maxFixedArity:]
		unpackedRestArgs := make([]interface{}, len(restArgs))
		for i, v := range restArgs {
			unpackedRestArgs[i] = v.Interface()
		}
		if len(unpackedRestArgs) == 0 {
			return varargs.Call(append(fixedArgs, reflect.ValueOf(CljsCoreList_EMPTY)))
		}
		return varargs.Call(append(fixedArgs, reflect.ValueOf(Array_seq.X_invoke_Arity1(unpackedRestArgs))))
	})
}

func makeInvalidArity(from reflect.Type) reflect.Value {
	return reflect.MakeFunc(from, func(in []reflect.Value) []reflect.Value {
		throwArity(nil, len(in))
		return nil
	})
}

func Fn(fns ...interface{}) *AFn {
	var f *AFn
	maxFixedArity := -1

	if len(fns) > 0 {
		if afn, ok := fns[0].(*AFn); ok {
			f = afn
			fns = fns[1:]
		}
	}

	if len(fns) > 0 {
		if arity, ok := fns[0].(int); ok {
			maxFixedArity = arity
			fns = fns[1:]
		}
	}

	if f == nil {
		f = &AFn{}
	}
	v := reflect.ValueOf(f).Elem()

	variadic := false
	for _, a := range fns {
		av := reflect.ValueOf(a)
		at := av.Type()
		in := at.NumIn()

		if at.IsVariadic() {
			variadic = true
			f.ArityVariadic = a.(func(...interface{}) interface{})
		} else {
			af := v.FieldByName(fmt.Sprintf("Arity%d", in))
			if sig := typedSignature(at); sig != "" {
				v.FieldByName(fmt.Sprintf("Arity%d%s", in, sig)).Set(av)
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
	fields := v.Type().NumField()
	for i := 0; i < fields; i++ {
		vf := v.Field(i)
		vt := vf.Type()
		if vf.Kind() == reflect.Func && vf.IsNil() && typedSignature(vt) == "" {
			if variadic && vt.NumIn() >= maxFixedArity {
				vf.Set(makeVarargsBridge(maxFixedArity, reflect.ValueOf(f.ArityVariadic), vt))
			} else {
				vf.Set(makeInvalidArity(vt))
			}
		}
	}
	return f
}

func Truth_(x interface{}) bool {
	return x != nil && x != false
}

func Seq_(x interface{}) CljsCoreISeq {
	if x == nil {
		return nil
	}
	return x.(CljsCoreISeq)
}

func Nil_(x interface{}) bool {
	if x == nil {
		return true
	}
	v := Value_(x)
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}

func Int32_(x float64) int32 {
	return int32(int(x))
}

func UInt32_(x float64) uint32 {
	return uint32(uint(x))
}

func MaxFixedArity_(x interface{}) float64 {
	if afn, ok := x.(*AFn); ok {
		return float64(afn.MaxFixedArity)
	}
	return -1
}

func Alength_(x interface{}) float64 {
	switch x := x.(type) {
	case []interface{}:
		return float64(len(x))
	case string:
		return float64(len(x))
	default:
		return float64(Value_(x).Len())
	}
}

func Aget_(x interface{}, idx float64) interface{} {
	switch x := x.(type) {
	case []interface{}:
		return x[int(idx)]
	case string:
		return string(x[int(idx)])
	default:
		return Value_(x).Index(int(idx)).Interface()
	}
}

type CljsCoreObject interface {
	ToString() string
	Equiv(other interface{}) bool
}

type CljsCoreIEmptyList interface {
	CljsCoreIList
	CljsCoreISeq
	CljsCoreICollection
}

func (this *CljsCorePersistentVector) X_contains_key_QMARK__Arity2(k interface{}) bool {
	panic("Not implemented.")
}

func (this *CljsCoreSubvec) X_contains_key_QMARK__Arity2(k interface{}) bool {
	panic("Not implemented.")
}

func (e *CljsCoreExceptionInfo) Error() string {
	return fmt.Sprint(e.Message)
}

var X___ *AFn

func init() {
	X_invoke.ArityVariadic = func(f_args ...interface{}) interface{} {
		f, args := f_args[0], f_args[1:]
		if f, ok := f.(*AFn); ok {
			return f.X_invoke_ArityVariadic(args...)
		}
		return throwArity(nil, len(args)-1)
	}

	// Hack, calls to minus gets over-munged.
	X___ = X_

	CljsCorePersistentTreeMap_EMPTY.Comp = Compare
}

func Main_() {
	Enable_console_print_BANG_.X_invoke_Arity0()
	args := make([]interface{}, len(os.Args[1:]))
	for i, a := range os.Args[1:] {
		args[i] = a
	}
	X_STAR_main_cli_fn_STAR_.(*AFn).X_invoke_ArityVariadic(args...)
}
