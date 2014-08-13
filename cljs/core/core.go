package core

import "github.com/hraberg/cljs.go/js"

// This file will eventually be generated from cljs.core.

var _STAR_print_fn_STAR_ = func(objs ...interface{}) interface{} {
	panic(js.Error{"No *print-fn* fn set for evaluation environment"})
}

func Println(objs ...interface{}) interface{} {
	_STAR_print_fn_STAR_(objs...)
	return nil
}

func Set_print_fn_BANG_(f func(objs ...interface{}) interface{}) {
	_STAR_print_fn_STAR_ = f
}
