package main

import (
	"fmt"
	"reflect"
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

type Foo struct{}

type Bar interface {
	Bar(...interface{}) interface{}
}

var Bar_cljs__lang__maxFixedArity = 2

func (f Foo) Bar_cljs__core__IFn___invoke__arity__1(x interface{}) interface{} {
	fmt.Printf("%d\n", x)
	return "Bar_cljs__core__IFn___invoke__arity__1"
}

func (f Foo) Bar_cljs__core__IFn___invoke__arity__2(x interface{}, y interface{}) interface{} {
	fmt.Printf("%d %d\n", x, y)
	return "Bar_cljs__core__IFn___invoke__arity__2"
}

func (f Foo) Bar_cljs__core__IFn___invoke__arity__variadic(x, y interface{}, xs ...interface{}) interface{} {
	fmt.Printf("%d %d %d\n", x, y, xs)
	return "Bar_cljs__core__IFn___invoke__arity__variadic"
}

func (f Foo) Bar_cljs__lang__applyTo(xs []interface{}) interface{} {
	var l = len(xs)
	switch {
	case l > Bar_cljs__lang__maxFixedArity:
		var v = reflect.ValueOf(f)
		var vs = make([]reflect.Value, len(xs))
		for i, x := range xs {
			vs[i] = reflect.ValueOf(x)
		}
		return v.MethodByName("Bar_cljs__core__IFn___invoke__arity__variadic").Call(vs)[0].Interface()
	case l == 1:
		return f.Bar_cljs__core__IFn___invoke__arity__1(xs[0])
	case l == 2:
		return f.Bar_cljs__core__IFn___invoke__arity__2(xs[0], xs[1])
	}
	panic(fmt.Sprintf("Invalid arity: %d", l))
}

func (f Foo) Bar(xs ...interface{}) interface{} {
	var l = len(xs)
	switch {
	case l > Bar_cljs__lang__maxFixedArity:
		return f.Bar_cljs__lang__applyTo(xs)
	case l == 1:
		return f.Bar_cljs__core__IFn___invoke__arity__1(xs[0])
	case l == 2:
		return f.Bar_cljs__core__IFn___invoke__arity__2(xs[0], xs[1])
	}
	panic(fmt.Sprintf("Invalid arity: %d", l))
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

func plus_one(x interface{}) interface{} {
	return (double(x) + (1))
}

func main() {
	fmt.Printf("ClojureScript to Go [go] %v\n", plus_one(1))

	var foo = Foo{}
	fmt.Printf("%v\n", foo.Bar(8))

	var xs = []int{4, 5, 6, 7}
	var is = make([]interface{}, len(xs))
	for i, x := range xs {
		is[i] = x
	}
	fmt.Printf("%v\n", foo.Bar_cljs__lang__applyTo(is))
	xs = []int{4}
	is = make([]interface{}, len(xs))
	for i, x := range xs {
		is[i] = x
	}
	fmt.Printf("%v\n", foo.Bar_cljs__lang__applyTo(is))

	xs = []int{8, 9}
	is = make([]interface{}, len(xs))
	for i, x := range xs {
		is[i] = x
	}
	fmt.Printf("%v\n", foo.Bar_cljs__lang__applyTo(is))

	fmt.Printf("%v\n", foo.Bar_cljs__core__IFn___invoke__arity__1(2))
	fmt.Printf("%v\n", foo.Bar_cljs__core__IFn___invoke__arity__2(2, 3))
	fmt.Printf("%v\n", foo.Bar_cljs__core__IFn___invoke__arity__variadic(2, 3, 4))

	fmt.Printf("%v\n", foo.Bar(2))
	fmt.Printf("%v\n", foo.Bar(2, 3))
	fmt.Printf("%v\n", foo.Bar(2, 3, 4))
	fmt.Printf("%v\n", foo.Bar())
}
