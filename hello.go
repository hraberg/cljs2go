// hello
package main

import "fmt"
import "github.com/hraberg/cljs.go/js"
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

func init() {
	Set_print_fn_BANG_(func(objs ...interface{}) interface{} {
		fmt.Println(objs...)
		return nil
	})
}

func main() {
	Foo()
	Foo("Space")
	Foo_cljs__core__IFn___invoke__arity__0()
	Foo_cljs__core__IFn___invoke__arity__1("Space")
	Foo("Space", "Hyper")
}
