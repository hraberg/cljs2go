package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
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

	assert.Equal(t, "Hello JavaScript World", sb.Append("Hello Java").Append("Script World").String())
	assert.Equal(t, "Hello JavaScript World", sb.String())

	assert.Equal(t, "l", (js.JSString("Hello").CharAt(2)))
	assert.Equal(t, 108, (js.JSString("Hello").CharCodeAt(2)))
	assert.Equal(t, 3.012568359e+09, (gstring.HashCode("Hello World")))
}

func Test_Main(t *testing.T) {
	mainWasCalled := false
	STAR_main_cli_fn_STAR_ = NewAFn(func(args ...interface{}) interface{} {
		mainWasCalled = true
		return nil
	})
	Main()
	assert.True(t, mainWasCalled)
}

var Baz = NewAFn(func(args ...interface{}) interface{} {
	x := args[0]
	xs := args[1:] // this should be an array-seq (an IndexedSeq backed by slices or arrays)
	_ = x
	return xs
}, func(x interface{}) interface{} {
	return x
})

func Test_Invoke(t *testing.T) {
	assert.True(t, NativeSatisifes_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "IFn"), Baz).(bool))
	assert.Equal(t, "Invalid arity: 0", func() (message string) {
		defer func() { message = fmt.Sprint(recover()) }()
		Baz.Call()
		assert.Fail(t, fmt.Sprintf("%#v should panic", Baz.Call))
		return
	}())
	assert.Panics(t, func() { Baz.Invoke_Arity0() })
	assert.Equal(t, 1, Baz.MaxFixedArity)
	assert.Equal(t, "Hello", Baz.Invoke_Arity1("Hello"))
	assert.Equal(t, "Hello", Invoke_.Invoke_Arity2(Baz, "Hello"))
	assert.Panics(t, func() { Baz.Invoke_Arity2("Hello", "World") })
	assert.Equal(t, []interface{}{"World"}, Baz.Call("Hello", "World"))
	assert.Equal(t, []interface{}{"World"}, Baz.Invoke_ArityVariadic("Hello", "World"))
	assert.Equal(t, []interface{}{"World"}, Invoke_.Invoke_ArityVariadic(Baz, "Hello", "World"))
	assert.Equal(t, []interface{}{"World"}, Invoke_.Call(Baz, "Hello", "World"))
	assert.Equal(t, []interface{}{"World"}, Apply.Invoke_ArityVariadic(Baz, "Hello", []interface{}{"World"}))

	assert.Equal(t, "", Str.Invoke_Arity0())
	assert.Equal(t, "Hello", Str.Invoke_Arity1("Hello"))
	assert.Equal(t, "1", Str.Invoke_Arity1(1))
	assert.Equal(t, "HelloClojureWorld", Str.Invoke_ArityVariadic("Hello", "Clojure", "World"))
}

func Test_Protocols(t *testing.T) {
	symbol := Symbol.Invoke_Arity2("foo", "bar")

	assert.True(t, NativeSatisifes_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "INamed"), symbol).(bool))
	assert.True(t, NativeSatisifes_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "IFn"), symbol).(bool))
	assert.Equal(t, "foo", symbol.(INamed).Namespace_Arity1())
	assert.Equal(t, "foo", Namespace_.Invoke_Arity1(symbol))

	assert.Equal(t, "foo", Namespace.Invoke_Arity1(symbol))
	assert.Panics(t, func() { Namespace.Invoke_Arity1(2) })

	assert.Equal(t, "bar", Name.Invoke_Arity1(symbol))
	assert.Panics(t, func() { Name.Invoke_Arity1(2) })
	assert.Equal(t, "baz", Name.Invoke_Arity1("baz"))

	assert.Equal(t, "foo", Invoke_.Invoke_Arity2(Namespace_, symbol))
	assert.Equal(t, "bar", symbol.(INamed).Name_Arity1())
	assert.Equal(t, "foo/bar", symbol.(fmt.Stringer).String())

	foo, bar := Symbol.Invoke_Arity1("foo"), Symbol.Invoke_Arity1("bar")
	m := ObjMap(map[interface{}]interface{}{foo: "bar"})

	assert.True(t, NativeSatisifes_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "ILookup"), m).(bool))
	assert.Equal(t, "bar", m.Lookup_Arity2(foo))
	assert.Nil(t, Lookup_.Call(m, bar))
	assert.Equal(t, "baz", m.Lookup_Arity3(bar, "baz"))

	assert.Equal(t, "bar", foo.(IFn).Invoke_Arity1(m))
	assert.Equal(t, "bar", Invoke_.Invoke_Arity2(foo, m))
	assert.Nil(t, bar.(IFn).Invoke_Arity1(m))
	assert.Equal(t, "baz", bar.(IFn).Invoke_Arity2(m, "baz"))
}

func Test_InteropViaReflection(t *testing.T) {
	sb := &gstring.StringBuffer{"foo"}
	assert.Equal(t, "foo", NativeGetInstanceField.Invoke_Arity2(sb, "Buffer"),
		"(.-buffer sb)")
	NativeSetInstanceField.Invoke_Arity3(sb, "Buffer", "bar")
	assert.Equal(t, "bar", sb.Buffer,
		"(set! (.-buffer sb) \"bar\")")
	assert.Equal(t, sb, NativeInvokeInstanceMethod.Invoke_Arity3(sb, "Append", []interface{}{"baz"}))
	assert.Equal(t, "barbaz", sb.Buffer,
		"(.append sb \"baz\")")

	assert.Equal(t, js.Number.MAX_VALUE, NativeGetInstanceField.Invoke_Arity2(js.Number, "MAX_VALUE"),
		"(.-MAX-VALUE js/Number)")

	assert.Equal(t, 3, NativeInvokeFunc.Invoke_Arity2(Math.Floor, []interface{}{3.14}),
		"(Math/floor 3.14)")
	assert.Equal(t, "3.14", NativeInvokeFunc.Invoke_Arity2(fmt.Sprint, []interface{}{3.14}),
		"(fmt/Sprint 3.14)")
	assert.Equal(t, 3.14, NativeInvokeFunc.Invoke_Arity2(js.ParseFloat, []interface{}{"3.14"}),
		"(js/parseFloat \"3.14\")")

	assert.Equal(t, "ABC", NativeInvokeFunc.Invoke_Arity2(js.String.FromCharCode, []interface{}{65, 66, 67}),
		"(.fromCharCode js/String 65 66 67)")
}

/*
These are the special-forms of ClojureScript.
We need two levels of tests, one in Clojure testing the emitter, and one in Go testing and defining the semantics.
For now, the first test will be string based, and the expected Go source will be duplicated and tested here.
Eventually we want to actually generate the Go tests from Clojure, or at least an example package they use instead.

:no-op
:var
:meta
:map
:list
:vector
:set
:js-value
:constant
:if
:case*
:throw
:def
:fn
:do
:try
:let
:loop
:recur
:letfn
:invoke
:new
:set!
:ns
:deftype*
:defrecord*
:dot
:js
*/

func Benchmark_RecursiveDirectCall(t *testing.B) {
	fib := func(this IFn) IFn {
		return NewAFn(this, func(n interface{}) interface{} {
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.Invoke_Arity1(n.(float64)-1.0).(float64) +
					this.Invoke_Arity1(n.(float64)-2.0).(float64)
			}
		})
	}(&AFn{})
	assert.Equal(t, 832040, fib.Invoke_Arity1(30.0))
}

func Benchmark_RecursiveDirectPrimitiveCall(t *testing.B) {
	fib := func() IFn {
		var this = &AFnPrimtive{}
		this.Arity1 = func(n interface{}) interface{} {
			return this.Arity1FF(n.(float64))
		}
		this.Arity1FF = func(n float64) float64 {
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.Arity1FF(n-1.0) + this.Arity1FF(n-2.0)
			}
		}
		return this
	}()
	assert.Equal(t, 832040, fib.Invoke_Arity1(30.0))
}

func Benchmark_RecursiveDispatch(t *testing.B) {
	fib := func() *AFn {
		var this = &AFn{}
		this.Arity1 = func(a interface{}) interface{} {
			var n = a.(float64)
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.Call(n-1.0).(float64) + this.Call(n-2.0).(float64)
			}
		}
		return this
	}()
	assert.Equal(t, 832040, fib.Call(30.0))
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
