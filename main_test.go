package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"unsafe"
)
import (
	garray "github.com/hraberg/cljs.go/goog/array"
	gobject "github.com/hraberg/cljs.go/goog/object"
	gstring "github.com/hraberg/cljs.go/goog/string"
	"github.com/hraberg/cljs.go/js"
	"github.com/hraberg/cljs.go/js/Math"
	"github.com/stretchr/testify/assert"
)
import . "github.com/hraberg/cljs.go/cljs/core"

/*
;; IFn

;; Used in core.cljs apply, which uses apply-to built by core.clj gen-apply-to.
;; It falls back to JS .apply it applyTo doesn't exist on the passed in fn.
.-cljs$lang$maxFixedArity ;; a field on the fn
.-cljs$lang$applyTo ;; accessed as field to see if it's there
.cljs$lang$applyTo  ;; then called as fn if it exists.

;; Used by the dispatch fn to actually invoke the various overloaded fns, uses JS arguments in a switch.
;; See emit* :fn and emit* :invoke
;; emit* :fn will emit a single fn or a dispatch fn around the real overlaoded fns.
;; At times these are called directly, like cljs.core.str.cljs$core$IFn$_invoke$arity$1(~{}) in cljs.core/str
.cljs$core$IFn$_invoke$arity$variadic
.cljs$core$IFn$_invoke$arity$N

;; defprotocol

.-cljs$lang$protocol_mask$partitionN$

;; deftype

.-cljs$lang$type
.-cljs$lang$ctorStr
.-cljs$lang$ctorPrWriter

*/

func Test_JS(t *testing.T) {
	assert.Equal(t, math.Inf(1), js.Infinity)
	assert.Equal(t, math.MaxFloat64, js.Number.MAX_VALUE)
	assert.Equal(t, 0.6046602879796196, Math.Random())
	assert.Equal(t, 3, Math.Ceil(2.6))
	assert.Equal(t, 2, Math.Floor(2.6))
	assert.Equal(t, 12, Math.Imul(2.3, 6.7))
	assert.Equal(t, "ABC", js.String.FromCharCode(65, 66, 67))
	assert.Nil(t, js.RegExp{"Hello", ""}.Exec("World"))
	assert.Equal(t, []string{"Hello", "Hello"}, js.RegExp{"hello", "i"}.Exec("World Hello Hello"))
	assert.Equal(t, "HELLO World", (js.JSString("Hello World").Replace(js.RegExp{"hello", "i"},
		func(match string) string {
			return strings.ToUpper(match)
		},
	)))
	assert.Equal(t, 6, (js.JSString("Hello World").Search(js.RegExp{"world", "i"})))
	assert.Equal(t, "(?i)Hello", (js.RegExp{"Hello", "i"}).String())

	date := js.Date{1407962432671}
	assert.Equal(t, 2014, date.GetUTCFullYear())
	assert.Equal(t, 7, date.GetUTCMonth())
	assert.Equal(t, 13, date.GetUTCDate())
	assert.Equal(t, 20, date.GetUTCHours())
	assert.Equal(t, 40, date.GetUTCMinutes())
	assert.Equal(t, 32, date.GetUTCSeconds())
	assert.Equal(t, 671, date.GetUTCMilliseconds())
	assert.Equal(t, "2014-08-13 21:40:32.000671 +0100 BST", date.String())

	assert.Equal(t, 3.14, js.ParseFloat("3.14"))
	assert.Equal(t, math.NaN(), js.ParseFloat(""))
	assert.Equal(t, 3, js.ParseInt("3", 10))
	assert.Equal(t, 10, js.ParseInt("a", 16))
	assert.Equal(t, math.NaN(), js.ParseInt("3.14", 10))
	assert.Equal(t, math.NaN(), js.ParseInt("x", 10))

	is := []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}
	garray.Shuffle(is)
	garray.StableSort(is, func(a, b interface{}) interface{} { return a.(float64) - b.(float64) })
	assert.Equal(t, []interface{}{5.0, 4.0, 3.0, 2.0, 1.0}, is)
	garray.Shuffle(is)
	garray.StableSort(is, garray.DefaultCompare)
	assert.Equal(t, []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}, is)

	ss := []interface{}{"foo", "bar"}
	garray.StableSort(ss, garray.DefaultCompare)
	assert.Equal(t, []interface{}{"bar", "foo"}, ss)

	obj := gobject.Create("foo", 2, "bar", 3)
	copy := make(map[string]interface{})
	gobject.ForEach(obj, func(k, v, o interface{}) interface{} {
		assert.Equal(t, obj, o)
		assert.Equal(t, v, o.(map[string]interface{})[k.(string)])
		copy[k.(string)] = v
		return nil
	})
	assert.Equal(t, obj, copy)

	sb := gstring.StringBuffer{}
	assert.Equal(t, "Hello JavaScript World", sb.Append("Hello Java").Append("Script World").ToString())
	assert.Equal(t, "Hello JavaScript World", sb.String())

	assert.Equal(t, "l", (js.JSString("Hello").CharAt(2)))
	assert.Equal(t, 108, (js.JSString("Hello").CharCodeAt(2)))
	assert.Equal(t, 3.012568359e+09, (gstring.HashCode("Hello World")))
}

func Test_Main(t *testing.T) {
	mainWasCalled := false
	STAR_main_cli_fn_STAR_ = func(args ...interface{}) interface{} {
		mainWasCalled = true
		return nil
	}
	Main()
	assert.True(t, mainWasCalled)
}

var Baz = AFn{
	CljsLangMaxFixedArity: 1,
	CljsCoreIFn_InvokeArityVariadic: func(args ...interface{}) interface{} {
		x := args[0]
		xs := args[1:] // this should be an array-seq (an IndexedSeq backed by slices or arrays)
		_ = x
		return xs
	},
	CljsCoreIFn_InvokeArity1: func(x interface{}) interface{} {
		return x
	},
}

func Test_Invoke(t *testing.T) {
	assert.Panics(t, func() { Baz.Invoke() })
	assert.Equal(t, "Hello", Baz.Invoke("Hello"))
	assert.Panics(t, func() { Baz.CljsCoreIFn_InvokeArity0() })
	assert.Equal(t, "Hello", Baz.CljsCoreIFn_InvokeArity1("Hello"))
	assert.Panics(t, func() { Baz.CljsCoreIFn_InvokeArity2("Hello", "World") })
	assert.Equal(t, []interface{}{"World"}, Baz.Invoke("Hello", "World"))
	assert.Equal(t, []interface{}{"World"}, Baz.CljsCoreIFn_InvokeArityVariadic("Hello", "World"))
	assert.Equal(t, []interface{}{"World"}, Baz.CljsLangApplyTo("Hello", []interface{}{"World"}))
}

// Protocols in ClojureScript don't seem to support vargs.
// In cljs.core, only IFn, IReduce, IIndexed, ILookup, and ISwap have overloaded arities.
// IFn is a special case which drops the receiver arg.

// (defprotocol INamed
//   (^string -name [x])
//   (^string -namespace [x]))

// (defprotocol ILookup
//   (-lookup [o k] [o k not-found]))

func Test_Protocols(t *testing.T) {
	symbol := Symbol.CljsCoreIFn_InvokeArity2("foo", "bar")

	assert.True(t, NativeSatisifes_QMARK_(Symbol.CljsCoreIFn_InvokeArity2("cljs.core", "INamed"), symbol).(bool))
	assert.True(t, NativeSatisifes_QMARK_(Symbol.CljsCoreIFn_InvokeArity2("cljs.core", "IFn"), symbol).(bool))
	assert.Equal(t, "foo", symbol.(INamed).Namespace())
	assert.Equal(t, "bar", symbol.(INamed).Name())
	assert.Equal(t, "foo/bar", symbol.(fmt.Stringer).String())

	foo, bar := Symbol.CljsCoreIFn_InvokeArity1("foo"), Symbol.CljsCoreIFn_InvokeArity1("bar")
	m := ObjMap(map[interface{}]interface{}{foo: "bar"})

	assert.True(t, NativeSatisifes_QMARK_(Symbol.CljsCoreIFn_InvokeArity2("cljs.core", "ILookup"), m).(bool))
	assert.Equal(t, "bar", m.Lookup(foo))
	assert.Nil(t, m.Lookup(bar))
	assert.Equal(t, "baz", m.Lookup(bar, "baz"))

	assert.Equal(t, "bar", foo.(IFn).Invoke(m))
	assert.Nil(t, bar.(IFn).Invoke(m))
	assert.Equal(t, "baz", bar.(IFn).Invoke(m, "baz"))
}

type ArityVariadic func(...interface{}) interface{}
type Arity0 func() interface{}
type Arity1 func(interface{}) interface{}
type Arity2 func(_, _ interface{}) interface{}

type INamed2 interface {
	Name() interface{}
	Name_Arity1() interface{}
	Namespace() interface{}
	Namespace_Arity1() interface{}
}

var Name2 = AFn2{
	Arity1: func(this interface{}) interface{} {
		return this.(INamed2).Name_Arity1()
	},
}

var Namespace2 = AFn2{
	Arity1: func(this interface{}) interface{} {
		return this.(INamed2).Namespace_Arity1()
	},
}

type Symbol2 struct {
	ns, name, str, _hash, _meta interface{}
}

func (this Symbol2) Name_Arity1() interface{} {
	return this.name
}

func (this Symbol2) Name() interface{} {
	return this.Name_Arity1()
}

func (this Symbol2) Namespace_Arity1() interface{} {
	return this.ns
}

func (this Symbol2) Namespace() interface{} {
	return this.Namespace_Arity1()
}

func (this Symbol2) ToString_Arity1() interface{} {
	return this.str
}

func (this Symbol2) ToString() string {
	return this.ToString_Arity1().(string)
}

var SymbolF = func() AFn2 {
	SymbolF := AFn2{}
	SymbolF.Arity1 = func(name interface{}) interface{} {
		return SymbolF.Invoke_Arity2(nil, name)
	}
	SymbolF.Arity2 = func(ns, name interface{}) interface{} {
		symStr := func() interface{} {
			if ns != nil {
				return ns.(string) + "/" + name.(string)
			} else {
				return name
			}
		}()
		return Symbol2{ns: ns, name: name, str: symStr}
	}
	return SymbolF
}()

type IFn2 interface {
	Invoke(...interface{}) interface{}
	Invoke_Arity0() interface{}
	Invoke_Arity1(interface{}) interface{}
	Invoke_Arity2(_, _ interface{}) interface{}
	Invoke_ArityVariadic(...interface{}) interface{}
}

// IFn is a special case, should have all arities up to 20 (they are - 1 as the receiver is the fn)
var Invoke2 = AFn2{
	Arity1: func(this interface{}) interface{} {
		return this.(IFn2).Invoke_Arity0()
	},
	Arity2: func(this, a interface{}) interface{} {
		return this.(IFn2).Invoke_Arity1(a)
	},
	ArityVariadic: func(this_a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
		return this_a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest[0].(IFn2).
			Invoke_ArityVariadic(this_a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest[1:]...)
	},
}

type AFn2 struct {
	MaxFixedArity int
	ArityVariadic
	Arity0
	Arity1
	Arity2
}

func ThrowArity2(f interface{}, arity int) interface{} {
	if f == nil {
		panic(js.Error{fmt.Sprint("Invalid arity: ", arity)})
	}
	return f
}

// This one should really be the dispatch fn
func (this AFn2) Invoke(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
	ThrowArity2(this.ArityVariadic, len(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest))
	return this.ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest...)
}

func (this AFn2) Invoke_ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest ...interface{}) interface{} {
	ThrowArity2(this.ArityVariadic, len(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest))
	return this.ArityVariadic(a_b_c_d_e_f_g_h_i_j_k_l_m_n_o_p_q_t_rest...)
}

func (this AFn2) Invoke_Arity0() interface{} {
	ThrowArity2(this.Arity0, 0)
	return this.Arity0()
}

func (this AFn2) Invoke_Arity1(a interface{}) interface{} {
	ThrowArity2(this.Arity1, 1)
	return this.Arity1(a)
}

func (this AFn2) Invoke_Arity2(a, b interface{}) interface{} {
	ThrowArity2(this.Arity2, 2)
	return this.Arity2(a, b)
}

func Test_ProtocolsFnStyle(t *testing.T) {
	symbol := SymbolF.Invoke_Arity2("foo", "bar")
	assert.Equal(t, "foo/bar", symbol.(Symbol2).ToString())
	assert.Equal(t, "bar", Name2.Invoke_Arity1(symbol))
	assert.Equal(t, "bar", Name2.Invoke_Arity1(symbol))
	assert.Equal(t, "foo", Namespace2.Invoke_Arity1(symbol))
	assert.Equal(t, "bar", INamed2.Name_Arity1(symbol.(INamed2)))
	assert.Equal(t, "bar", INamed2.Name(symbol.(INamed2)))
	assert.Equal(t, "bar", symbol.(INamed2).Name_Arity1())
	assert.Equal(t, "bar", symbol.(INamed2).Name())
	assert.Equal(t, "foo", INamed2.Namespace_Arity1(symbol.(INamed2)))
	assert.Equal(t, "foo", INamed2.Namespace(symbol.(INamed2)))
	assert.Equal(t, "foo", symbol.(INamed2).Namespace_Arity1())
	assert.Equal(t, "foo", symbol.(INamed2).Namespace())

	baz := SymbolF.Invoke_Arity1("baz")
	assert.Equal(t, "baz", Name2.Invoke_Arity1(baz))
	assert.Nil(t, Namespace2.Invoke_Arity1(baz))
}

func Benchmark_RecursiveDirectCallPrimitiveLocal(t *testing.B) {
	fib := func() AFn {
		var this = AFn{}
		this.CljsCoreIFn_InvokeArity1 = func(a interface{}) interface{} {
			var n = a.(float64)
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.CljsCoreIFn_InvokeArity1(n-1.0).(float64) +
					this.CljsCoreIFn_InvokeArity1(n-2.0).(float64)
			}
		}
		return this
	}()
	assert.Equal(t, 832040, fib.Invoke(30.0))
}

func Benchmark_RecursiveDirectCall(t *testing.B) {
	fib := func() AFn {
		var this = AFn{}
		this.CljsCoreIFn_InvokeArity1 = func(n interface{}) interface{} {
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.CljsCoreIFn_InvokeArity1(n.(float64)-1.0).(float64) +
					this.CljsCoreIFn_InvokeArity1(n.(float64)-2.0).(float64)
			}
		}
		return this
	}()
	assert.Equal(t, 832040, fib.Invoke(30.0))
}

func Benchmark_RecursiveDirectPrimitiveCall(t *testing.B) {
	fib := func() AFn {
		var this = AFn{}
		this.CljsCoreIFn_InvokeArity1DD = func(n float64) float64 {
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.CljsCoreIFn_InvokeArity1DD(n-1.0) +
					this.CljsCoreIFn_InvokeArity1DD(n-2.0)
			}
		}
		return this
	}()
	assert.Equal(t, 832040, fib.Invoke(30.0))
}

func Benchmark_RecursiveDispatch(t *testing.B) {
	fib := func() AFn {
		var this = AFn{}
		this.CljsCoreIFn_InvokeArity1 = func(a interface{}) interface{} {
			var n = a.(float64)
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.Invoke(n-1.0).(float64) +
					this.Invoke(n-2.0).(float64)
			}
		}
		return this
	}()
	assert.Equal(t, 832040, fib.Invoke(30.0))
}

func Benchmark_RecursiveGo(t *testing.B) {
	fib := func() func(float64) float64 {
		var this func(float64) float64
		this = func(n float64) float64 {
			if n == 0 {
				return 0
			} else if n == 1 {
				return 1
			} else {
				return this(n-1) + this(n-2)
			}
		}
		return this
	}()
	assert.Equal(t, 832040, fib(30))
}

func Benchmark_RecursiveGoUintptr(t *testing.B) {
	fib := func() func(uintptr) uintptr {
		var this func(uintptr) uintptr
		this = func(n uintptr) uintptr {
			if math.Float64frombits(uint64(n)) == 0.0 {
				return uintptr(math.Float64bits(0))
			} else if math.Float64frombits(uint64(n)) == 1.0 {
				return uintptr(math.Float64bits(1.0))
			} else {
				return uintptr(math.Float64bits(math.Float64frombits(uint64(this(uintptr(math.Float64bits(math.Float64frombits(uint64(n))-1))))) +
					math.Float64frombits(uint64(this(uintptr(math.Float64bits(math.Float64frombits(uint64(n))-2)))))))
			}
		}
		return this
	}()
	assert.Equal(t, 832040, math.Float64frombits(uint64(fib(uintptr(math.Float64bits(30.0))))))
}

func Benchmark_RecursiveGoUnsafePointer(t *testing.B) {
	fib := func() func(unsafe.Pointer) unsafe.Pointer {
		var this func(unsafe.Pointer) unsafe.Pointer
		this = func(n unsafe.Pointer) unsafe.Pointer {
			if math.Float64frombits(uint64(uintptr(n))) == 0 {
				return unsafe.Pointer(uintptr(math.Float64bits(0)))
			} else if math.Float64frombits(uint64(uintptr(n))) == 1 {
				return unsafe.Pointer(uintptr(math.Float64bits(1)))
			} else {
				return unsafe.Pointer(uintptr(math.Float64bits(math.Float64frombits(uint64(uintptr(this(unsafe.Pointer(uintptr(math.Float64bits(math.Float64frombits(uint64(uintptr(n)))-1))))))) + math.Float64frombits(uint64(uintptr(this(unsafe.Pointer(uintptr(math.Float64bits(math.Float64frombits(uint64(uintptr(n)))-2))))))))))
			}
		}
		return this
	}()
	assert.Equal(t, 832040, math.Float64frombits(uint64(uintptr(fib(unsafe.Pointer(uintptr(math.Float64bits(30))))))))
}

func BoxFloat64(f float64) unsafe.Pointer {
	return unsafe.Pointer(uintptr(math.Float64bits(f)))
}

func UnboxFloat64(p unsafe.Pointer) float64 {
	return math.Float64frombits(uint64(uintptr(p)))
}

func Benchmark_RecursiveGoUnsafePointerBox(t *testing.B) {
	fib := func() func(unsafe.Pointer) unsafe.Pointer {
		var this func(unsafe.Pointer) unsafe.Pointer
		this = func(n unsafe.Pointer) unsafe.Pointer {
			if UnboxFloat64(n) == 0 {
				return BoxFloat64(0)
			} else if UnboxFloat64(n) == 1 {
				return BoxFloat64(1)
			} else {
				return BoxFloat64(UnboxFloat64(this(BoxFloat64(UnboxFloat64(n)-1))) +
					UnboxFloat64(this(BoxFloat64(UnboxFloat64(n)-2))))
			}
		}
		return this
	}()
	assert.Equal(t, 832040, UnboxFloat64(fib(BoxFloat64(30))))
}

func IBoxFloat64(f float64) interface{} {
	var i interface{}
	// v := (*[2]uint64)(unsafe.Pointer(&i))
	//v[0] = 5262720
	// v[1] = math.Float64bits(f)
	(*[2]uint64)(unsafe.Pointer(&i))[1] = math.Float64bits(f)
	return i
}

func IUnboxFloat64(p interface{}) float64 {
	return math.Float64frombits((*[2]uint64)(unsafe.Pointer(&p))[1])
}

func Benchmark_RecursiveGoEmbedInInterface(t *testing.B) {
	fib := func() func(interface{}) interface{} {
		var this func(interface{}) interface{}
		this = func(n interface{}) interface{} {
			if IUnboxFloat64(n) == 0 {
				return IBoxFloat64(0)
			} else if IUnboxFloat64(n) == 1 {
				return IBoxFloat64(1)
			} else {
				return IBoxFloat64(IUnboxFloat64(this(IBoxFloat64(IUnboxFloat64(n)-1))) +
					IUnboxFloat64(this(IBoxFloat64(IUnboxFloat64(n)-2))))
			}
		}
		return this
	}()
	assert.Equal(t, 832040, IUnboxFloat64(fib(IBoxFloat64(30))))
}

func Test_Pointers(t *testing.T) {
	var x interface{} = "Hello"
	x_ptr := &x
	assert.Equal(t, "Hello", *x_ptr)
	u_ptr := uintptr(unsafe.Pointer(x_ptr))
	assert.Equal(t, "Hello", *(*interface{})(unsafe.Pointer(u_ptr)))
	//	var x_ptr uint64 = unsafe.Pointer(&x).(unit64)

	pi := 3.14
	var pi_ptr interface{} = uintptr(math.Float64bits(pi))

	// func(i interface{}) {
	// 	e := unsafe.Pointer(&i)
	// 	val := (*[2]uintptr)(e)
	// 	fmt.Println(*val)
	// 	var x interface{} = (unsafe.Pointer(uintptr(0)))
	// 	fmt.Println(reflect.TypeOf(x))
	// 	fmt.Println(*(*[2]uintptr)(unsafe.Pointer(&x)))
	// 	var y interface{}
	// 	fmt.Println(reflect.TypeOf(y))
	// 	fmt.Println(*(*[2]uintptr)(unsafe.Pointer(&y)))

	// 	fmt.Println(math.Float64frombits(uint64(val[1])))
	// 	val[0] = uintptr(math.Float64bits(2.0))
	// 	val[1] = 0
	// 	fmt.Println(uintptr(math.Float64bits(2.0)))
	// 	fmt.Println(math.Float64bits(2.0))
	// 	// fmt.Println(reflect.TypeOf(i))
	// 	//		fmt.Println(i)
	// }(1.0)

	// var pi_interface interface{} = pi_ptr
	//	assert.Equal(t, pi, math.Float64frombits(uint64(uintptr(pi_ptr))))
	assert.Equal(t, pi, math.Float64frombits(uint64(pi_ptr.(uintptr))))
	assert.Equal(t, pi, *(*float64)(unsafe.Pointer(&pi)))

	// var y = *(*interface{})(unsafe.Pointer(p))
	// assert.Equal(t, "Hello", y)
	assert.Equal(t, "Hello", *(*interface{})(unsafe.Pointer(&x)))
}

func double(x interface{}) float64 {
	switch x.(type) {
	case int:
		return float64(x.(int))
	case int64:
		return float64(x.(int64))
	default:
		return x.(float64)
	}
}

func long(x interface{}) int64 {
	switch x.(type) {
	case int:
		return int64(x.(int))
	case float64:
		return int64(x.(float64))
	default:
		return x.(int64)
	}
}
