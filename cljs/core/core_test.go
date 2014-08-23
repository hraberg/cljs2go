package core

import (
	"fmt"

	"testing"
)
import (
	gstring "github.com/hraberg/cljs.go/goog/string"
	"github.com/hraberg/cljs.go/js"
	"github.com/hraberg/cljs.go/js/Math"
	"github.com/stretchr/testify/assert"
)

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

func Test_Main(t *testing.T) {
	mainWasCalled := false
	STAR_main_cli_fn_STAR_ = Fn(func(args ...interface{}) interface{} {
		mainWasCalled = true
		return nil
	})
	Main()
	assert.True(t, mainWasCalled)
}

var Baz = Fn(func(args ...interface{}) interface{} {
	x := args[0]
	xs := args[1:] // this should be an array-seq (an IndexedSeq backed by slices or arrays)
	_ = x
	return xs
}, func(x interface{}) interface{} {
	return x
})

func Test_Invoke(t *testing.T) {
	assert.True(t, NativeSatisifes_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "IFn"), Baz).(bool))
	PanicsWith(t, "Invalid arity: 0", func() { Baz.(*AFn).Call() })
	PanicsWith(t, "Invalid arity: 0", func() { Baz.Invoke_Arity0() })
	assert.Equal(t, 1, Baz.(*AFn).MaxFixedArity)
	assert.True(t, Baz.(*AFn).IsVariadic())
	assert.Equal(t, "Hello", Baz.Invoke_Arity1("Hello"))
	assert.Equal(t, "Hello", Invoke_.Invoke_Arity2(Baz, "Hello"))
	assert.Equal(t, []interface{}{"World"}, Baz.(*AFn).Call("Hello", "World"))
	assert.Equal(t, []interface{}{"World"}, Baz.Invoke_ArityVariadic("Hello", "World"))
	assert.Equal(t, []interface{}{"World"}, Invoke_.Invoke_ArityVariadic(Baz, "Hello", "World"))
	assert.Equal(t, []interface{}{"World"}, Invoke_.(*AFn).Call(Baz, "Hello", "World"))
	assert.Equal(t, []interface{}{"World"}, Apply.Invoke_ArityVariadic(Baz, "Hello", []interface{}{"World"}))

	assert.Equal(t, []interface{}{"World"}, Baz.Invoke_Arity2("Hello", "World"))
	assert.Equal(t, []interface{}{"World", "Space"}, Baz.Invoke_Arity3("Hello", "World", "Space"))

	assert.Equal(t, "", Str.Invoke_Arity0())
	assert.Equal(t, "Hello", Str.Invoke_Arity1("Hello"))
	assert.Equal(t, "1", Str.Invoke_Arity1(1))
	assert.Equal(t, "HelloClojureWorld", Str.Invoke_ArityVariadic("Hello", "Clojure", "World"))
}

func Test_PrimitiveFn(t *testing.T) {
	fib := func(this *AFnPrimtive) *AFnPrimtive {
		return Fn(this, func(n float64) float64 {
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.Arity1FF(n-1.0) + this.Arity1FF(n-2.0)
			}
		}).(*AFnPrimtive)
	}(&AFnPrimtive{})
	assert.NotNil(t, fib.Arity1FF)
	assert.NotNil(t, fib.Arity1)
	assert.Nil(t, fib.Arity0F)
	assert.Equal(t, 832040, fib.Invoke_Arity1(30.0))

	odd := Fn(&AFnPrimtive{}, func(n interface{}) bool {
		return int(n.(float64))%2 != 0
	}).(*AFnPrimtive)
	assert.NotNil(t, odd.Arity1IB)
	assert.NotNil(t, odd.Arity1)
	assert.True(t, odd.Invoke_Arity1(1.0).(bool))
	assert.False(t, odd.Arity1IB(2.0))

	assert.True(t, Not.Invoke_Arity1(false).(bool))
	assert.False(t, Not.Arity1IB(true))
}

func Test_Protocols(t *testing.T) {
	symbol := Symbol.Invoke_Arity2("foo", "bar")

	assert.True(t, symbol.(js.Object).Equiv(Symbol.Invoke_Arity2("foo", "bar")))
	assert.True(t, NativeSatisifes_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("js", "Object"), symbol).(bool))
	assert.True(t, NativeSatisifes_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "INamed"), symbol).(bool))
	assert.True(t, NativeSatisifes_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "IFn"), symbol).(bool))
	assert.Equal(t, "foo", symbol.(INamed).Namespace_Arity1())
	assert.Equal(t, "foo", Namespace_.Invoke_Arity1(symbol))
	PanicsWith(t, "Invalid arity: 0", func() { symbol.(*CljsCoreSymbol).Invoke_Arity0() })
	PanicsWith(t, "Invalid arity: 0", func() { Invoke_.Invoke_Arity1(symbol) })

	assert.Equal(t, "foo", Namespace.Invoke_Arity1(symbol))
	PanicsWith(t, "Doesn't support namespace: 2", func() { Namespace.Invoke_Arity1(2) })

	assert.Equal(t, "bar", Name.Invoke_Arity1(symbol))
	PanicsWith(t, "Doesn't support name: 2", func() { Name.Invoke_Arity1(2) })
	assert.Equal(t, "baz", Name.Invoke_Arity1("baz"))

	assert.Equal(t, "foo", Invoke_.Invoke_Arity2(Namespace_, symbol))
	assert.Equal(t, "bar", symbol.(INamed).Name_Arity1())
	assert.Equal(t, "foo/bar", symbol.(fmt.Stringer).String())

	foo, bar := Symbol.Invoke_Arity1("foo"), Symbol.Invoke_Arity1("bar")
	m := ObjMap(map[interface{}]interface{}{foo: "bar"})

	assert.True(t, NativeSatisifes_QMARK_.Invoke_Arity2(Symbol.Invoke_Arity2("cljs.core", "ILookup"), m).(bool))
	assert.Equal(t, "bar", m.Lookup_Arity2(foo))
	assert.Nil(t, Lookup_.(*AFn).Call(m, bar))
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
	assert.Equal(t, "barbaz", sb.String(),
		"(.append sb \"baz\")")

	assert.Equal(t, "H", NativeInvokeInstanceMethod.Invoke_Arity3("Hello", "CharAt", []interface{}{0.0}))

	assert.Equal(t, js.Number.MAX_VALUE, NativeGetInstanceField.Invoke_Arity2(js.Number, "MAX_VALUE"),
		"(.-MAX-VALUE js/Number)")

	assert.Equal(t, 3, NativeInvokeFunc.Invoke_Arity2(Math.Floor, []interface{}{3.14}),
		"(Math/floor 3.14)")
	assert.Equal(t, "3.14", NativeInvokeFunc.Invoke_Arity2(fmt.Sprint, []interface{}{3.14}),
		"(fmt/Sprint 3.14)")
	assert.Equal(t, 3.14, NativeInvokeFunc.Invoke_Arity2(js.ParseFloat, []interface{}{"3.14"}),
		"(js/parseFloat \"3.14\")")

	assert.Equal(t, "ABC", NativeInvokeFunc.Invoke_Arity2(js.String.FromCharCode, []interface{}{65.0, 66.0, 67.0}),
		"(.fromCharCode js/String 65 66 67)")
}

func PanicsWith(t *testing.T, message string, f assert.PanicTestFunc) {
	assert.Equal(t, message, func() (message string) {
		defer func() { message = fmt.Sprint(recover()) }()
		f()
		assert.Fail(t, "should panic")
		return
	}())
}
