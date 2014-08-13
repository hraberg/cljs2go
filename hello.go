// hello
package main

import (
	"fmt"
	"strings"
)
import (
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

	js.Console.Log(js.RegExp{"hello", "i"}.Exec("World Hello Hello"), js.RegExp{"Hello", ""}.Exec("World") == nil)

	js.Console.Log(js.JSString("Hello World").Replace(js.RegExp{"hello", "i"},
		func(match string) string {
			return strings.ToUpper(match)
		},
	))
	js.Console.Log(js.JSString("Hello World").Search(js.RegExp{"world", "i"}))

	var sb = gstring.StringBuffer{}
	sb = sb.Append("Hello Java").Append("Script World")
	js.Console.Log(sb.ToString())
}
