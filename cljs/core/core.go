package core

import "github.com/hraberg/cljs.go/js"

import "fmt"

// This file will eventually be generated from cljs.core.

var _STAR_print_fn_STAR_ = func(_ interface{}) interface{} {
	panic(js.Error{"No *print-fn* fn set for evaluation environment"})
}

var _STAR_print_newline_STAR_ = true

func Enable_console_print_BANG_() interface{} {
	_STAR_print_newline_STAR_ = false
	_STAR_print_fn_STAR_ = func(x interface{}) interface{} {
		js.Console.Log(x)
		return nil
	}
	return nil
}

func pr_opts() interface{} {
	return nil
}

func Newline(opts interface{}) interface{} {
	if _STAR_print_newline_STAR_ {
		_STAR_print_fn_STAR_("\n")
	}
	return nil
}

func Println(objs ...interface{}) interface{} {
	_STAR_print_fn_STAR_(fmt.Sprint(objs...))
	Newline(pr_opts())
	return nil
}
