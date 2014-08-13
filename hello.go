// hello
package main

import (
	"fmt"
	"strings"
)
import (
	garray "github.com/hraberg/cljs.go/goog/array"
	gobject "github.com/hraberg/cljs.go/goog/object"
	gstring "github.com/hraberg/cljs.go/goog/string"
	"github.com/hraberg/cljs.go/js"
)
import . "github.com/hraberg/cljs.go/cljs/core"

func init() {
	Foo_cljs__core__IFn___invoke__arity__0 = func() interface{} {
		return Foo("World")
	}
	Foo_cljs__core__IFn___invoke__arity__1 = func(x interface{}) interface{} {
		return Println("Hello ", x)
	}
	Foo = func(arguments ...interface{}) interface{} {
		switch len(arguments) {
		case 0:
			return Foo_cljs__core__IFn___invoke__arity__0()
		case 1:
			return Foo_cljs__core__IFn___invoke__arity__1(arguments[0])
		}
		panic(js.Error{fmt.Sprint("Invalid arity: ", len(arguments))})
	}
}

var Foo_cljs__core__IFn___invoke__arity__0 func() interface{}
var Foo_cljs__core__IFn___invoke__arity__1 func(interface{}) interface{}
var Foo func(...interface{}) interface{}

func main() {
	Enable_console_print_BANG_()
	Foo()
	Foo("Space")
	Foo_cljs__core__IFn___invoke__arity__0()
	Foo_cljs__core__IFn___invoke__arity__1("Space")
	Foo("Space", "Hyper")
}

// JS smoke tests
func init() {
	js.Console.Log("Javascript", "Rules", js.Math.Random(), js.Infinity,
		js.Math.Ceil(2.6), js.Math.Imul(2.3, 6.7), js.String.FromCharCode(65, 66, 67))
	js.Console.Log(js.RegExp{"hello", "i"}.Exec("World Hello Hello"), js.RegExp{"Hello", ""}.Exec("World") == nil)
	js.Console.Log(js.JSString("Hello World").Replace(js.RegExp{"hello", "i"},
		func(match string) string {
			return strings.ToUpper(match)
		},
	))
	js.Console.Log(js.JSString("Hello World").Search(js.RegExp{"world", "i"}))
	var date = js.Date{1407962432671}
	js.Console.Log(date.GetUTCFullYear())
	js.Console.Log(date.GetUTCMonth())
	js.Console.Log(date.GetUTCDate())
	js.Console.Log(date.GetUTCHours())
	js.Console.Log(date.GetUTCMinutes())
	js.Console.Log(date.GetUTCSeconds())
	js.Console.Log(date.GetUTCMilliseconds())
	js.Console.Log(date)
	js.Console.Log(js.RegExp{"Hello", "i"})

	var xs = []int{1, 2, 3, 4, 5}
	var is = make([]interface{}, len(xs))
	for i, x := range xs {
		is[i] = float64(x)
	}
	garray.Shuffle(is)
	js.Console.Log(is)
	garray.StableSort(is, func(a, b interface{}) interface{} { return a.(float64) - b.(float64) })
	js.Console.Log(is)
	garray.Shuffle(is)
	js.Console.Log(is)
	garray.StableSort(is, garray.DefaultCompare)
	js.Console.Log(is)

	var ss = []string{"foo", "bar"}
	is = make([]interface{}, len(ss))
	for i, s := range ss {
		is[i] = s
	}
	garray.StableSort(is, garray.DefaultCompare)
	js.Console.Log(is)

	var obj = gobject.Create("foo", 2, "bar", 3)
	js.Console.Log(obj)
	gobject.ForEach(obj, func(k, v, obj interface{}) interface{} {
		js.Console.Log("k:", k, "v:", v, "in", obj)
		return nil
	})

	var sb = gstring.StringBuffer{}
	sb = sb.Append("Hello Java").Append("Script World")
	js.Console.Log(sb.ToString())

	js.Console.Log(gstring.HashCode("Hello World"))
}
