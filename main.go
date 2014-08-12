package main

import (
	"fmt"
	"log"
	"os/exec"
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

var Bar_cljs__lang__maxFixedArity = 2

func Bar_cljs__core__IFn___invoke__arity__1(x interface{}) interface{} {
	fmt.Printf("%d\n", x)
	return "Bar_cljs__core__IFn___invoke__arity__1"
}

func Bar_cljs__core__IFn___invoke__arity__2(x interface{}, y interface{}) interface{} {
	fmt.Printf("%d %d\n", x, y)
	return "Bar_cljs__core__IFn___invoke__arity__2"
}

func Bar_cljs__core__IFn___invoke__arity__variadic(x, y interface{}, xs ...interface{}) interface{} {
	fmt.Printf("%d %d %d\n", x, y, xs)
	return "Bar_cljs__core__IFn___invoke__arity__variadic"
}

func Bar(xs ...interface{}) interface{} {
	var l = len(xs)
	switch {
	case l > 2:
		return Bar_cljs__core__IFn___invoke__arity__variadic(xs[0], xs[1], xs[2:]...)
	case l == 1:
		return Bar_cljs__core__IFn___invoke__arity__1(xs[0])
	case l == 2:
		return Bar_cljs__core__IFn___invoke__arity__2(xs[0], xs[1])
	}
	panic(fmt.Sprintf("Invalid arity: %d", l))
}

func Bar_cljs__lang__applyTo(xs []interface{}) interface{} {
	return Bar(xs...)
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
	fmt.Printf("ClojureScript to Go [go]\n")
	out, err := exec.Command("go", "get", "code.google.com/p/go.tools/cmd/goimports").CombinedOutput()
	if err != nil {
		log.Fatal(string(out[:]))
	}

	var xs = []int{4, 5, 6, 7}
	var is = make([]interface{}, len(xs))
	for i, x := range xs {
		is[i] = x
	}
	fmt.Printf("%v\n", Bar_cljs__lang__applyTo(is))
	xs = []int{4}
	is = make([]interface{}, len(xs))
	for i, x := range xs {
		is[i] = x
	}
	fmt.Printf("%v\n", Bar_cljs__lang__applyTo(is))

	xs = []int{8, 9}
	is = make([]interface{}, len(xs))
	for i, x := range xs {
		is[i] = x
	}
	fmt.Printf("%v\n", Bar_cljs__lang__applyTo(is))

	fmt.Printf("%v\n", Bar_cljs__core__IFn___invoke__arity__1(2))
	fmt.Printf("%v\n", Bar_cljs__core__IFn___invoke__arity__2(2, 3))
	fmt.Printf("%v\n", Bar_cljs__core__IFn___invoke__arity__variadic(2, 3, 4))

	fmt.Printf("%v\n", Bar(2))
	fmt.Printf("%v\n", Bar(2, 3))
	fmt.Printf("%v\n", Bar(2, 3, 4))
	fmt.Printf("%v\n", Bar())
}
