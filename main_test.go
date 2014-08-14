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

	var date = js.Date{1407962432671}
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

	var is = []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}
	garray.Shuffle(is)
	garray.StableSort(is, func(a, b interface{}) interface{} { return a.(float64) - b.(float64) })
	assert.Equal(t, []interface{}{5.0, 4.0, 3.0, 2.0, 1.0}, is)
	garray.Shuffle(is)
	garray.StableSort(is, garray.DefaultCompare)
	assert.Equal(t, []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}, is)

	var ss = []interface{}{"foo", "bar"}
	garray.StableSort(ss, garray.DefaultCompare)
	assert.Equal(t, []interface{}{"bar", "foo"}, ss)

	var obj = gobject.Create("foo", 2, "bar", 3)
	var copy = make(map[string]interface{})
	gobject.ForEach(obj, func(k, v, o interface{}) interface{} {
		assert.Equal(t, obj, o)
		assert.Equal(t, v, o.(map[string]interface{})[k.(string)])
		copy[k.(string)] = v
		return nil
	})
	assert.Equal(t, obj, copy)

	var sb = gstring.StringBuffer{}
	assert.Equal(t, "Hello JavaScript World", sb.Append("Hello Java").Append("Script World").ToString())
	assert.Equal(t, "Hello JavaScript World", sb.String())

	assert.Equal(t, "l", (js.JSString("Hello").CharAt(2)))
	assert.Equal(t, 108, (js.JSString("Hello").CharCodeAt(2)))
	assert.Equal(t, 3.012568359e+09, (gstring.HashCode("Hello World")))
}

func Test_Main(t *testing.T) {
	var mainWasCalled = false
	STAR_main_cli_fn_STAR_ = func(args ...interface{}) interface{} {
		mainWasCalled = true
		return nil
	}
	Main()
	assert.True(t, mainWasCalled)
}

func init() {
	Foo_cljs__core__IFn___invoke__arity__1 = func(x interface{}) interface{} {
		return fmt.Sprint("Hello ", x)
	}
	Foo_cljs__core__IFn___invoke__arity__variadic = func(x interface{}, xs ...interface{}) interface{} {
		return fmt.Sprint("Hello ", x, xs)
	}

	Foo = func(arguments ...interface{}) interface{} {
		var l = len(arguments)
		switch {
		case l > 1:
			return Foo_cljs__core__IFn___invoke__arity__variadic(arguments[0], arguments[1:]...)
		case l == 1:
			return Foo_cljs__core__IFn___invoke__arity__1(arguments[0])
		}
		panic(js.Error{fmt.Sprint("Invalid arity: ", len(arguments))})
	}
}

var Foo_cljs__lang__maxFixedArity = 1
var Foo_cljs__core__IFn___invoke__arity__1 func(interface{}) interface{}
var Foo_cljs__core__IFn___invoke__arity__variadic func(interface{}, ...interface{}) interface{}
var Foo func(...interface{}) interface{}

func Foo_cljs__lang__applyTo(xs []interface{}) interface{} {
	return Foo(xs...)
}

func Test_Dispatch(t *testing.T) {
	assert.Equal(t, "Hello Space", Foo("Space"))
	assert.Equal(t, "Hello Space", Foo_cljs__core__IFn___invoke__arity__1("Space"))
	assert.Equal(t, "Hello Space[Hyper]", Foo("Space", "Hyper"))
	assert.Equal(t, "Hello Space[Hyper]", Foo_cljs__core__IFn___invoke__arity__variadic("Space", "Hyper"))
	assert.Equal(t, "Hello foo[bar]", Foo_cljs__lang__applyTo([]interface{}{"foo", "bar"}))
	assert.Panics(t, func() { Foo() })
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
