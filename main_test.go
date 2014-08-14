package main

import (
	"fmt"
	"os"
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

func Main(args ...interface{}) interface{} {
	return nil
}

func MainPreamble() {
	Enable_console_print_BANG_()
	var args = make([]interface{}, len(os.Args[1:]))
	for i, a := range os.Args[1:] {
		args[i] = a
	}
	Main(args...)
}

func Test_JS(t *testing.T) {
	js.Console.Log("Javascript", "Rules", Math.Random(), js.Infinity, js.Number.MAX_VALUE,
		Math.Ceil(2.6), Math.Imul(2.3, 6.7), js.String.FromCharCode(65, 66, 67))
	js.Console.Log(js.RegExp{"hello", "i"}.Exec("World Hello Hello"), js.RegExp{"Hello", ""}.Exec("World") == nil)
	js.Console.Log(js.JSString("Hello World").Replace(js.RegExp{"hello", "i"},
		func(match string) string {
			return strings.ToUpper(match)
		},
	))
	js.Console.Log(js.JSString("Hello World").Search(js.RegExp{"world", "i"}))
	js.Console.Log(js.RegExp{"Hello", "i"})

	var date = js.Date{1407962432671}
	js.Console.Log(date.GetUTCFullYear())
	js.Console.Log(date.GetUTCMonth())
	js.Console.Log(date.GetUTCDate())
	js.Console.Log(date.GetUTCHours())
	js.Console.Log(date.GetUTCMinutes())
	js.Console.Log(date.GetUTCSeconds())
	js.Console.Log(date.GetUTCMilliseconds())
	js.Console.Log(date)

	js.Console.Log(js.ParseFloat("3.14"))
	js.Console.Log(js.ParseFloat(""))
	js.Console.Log(js.ParseInt("314", 10))
	js.Console.Log(js.ParseInt("x", 10))

	var is = []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}
	garray.Shuffle(is)
	js.Console.Log(is)
	garray.StableSort(is, func(a, b interface{}) interface{} { return a.(float64) - b.(float64) })
	js.Console.Log(is)
	garray.Shuffle(is)
	js.Console.Log(is)
	garray.StableSort(is, garray.DefaultCompare)
	js.Console.Log(is)

	var ss = []interface{}{"foo", "bar"}
	garray.StableSort(ss, garray.DefaultCompare)
	js.Console.Log(ss)

	var obj = gobject.Create("foo", 2, "bar", 3)
	js.Console.Log(obj)
	gobject.ForEach(obj, func(k, v, obj interface{}) interface{} {
		js.Console.Log("k:", k, "v:", v, "in", obj)
		return nil
	})

	var sb = gstring.StringBuffer{}
	js.Console.Log(sb.Append("Hello Java").Append("Script World").ToString())
	js.Console.Log(sb.String())

	js.Console.Log(js.JSString("Hello").CharAt(2))
	js.Console.Log(js.JSString("Hello").CharCodeAt(2))

	js.Console.Log(gstring.HashCode("Hello World"))
}

func init() {
	MainPreamble()
}

var Foo_cljs__lang__maxFixedArity = 1

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
