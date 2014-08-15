package core

import (
	"os"

	"github.com/hraberg/cljs.go/js"
)

import "fmt"

// This file will eventually be generated from cljs.core.

var STAR_print_fn_STAR_ = func(_ interface{}) interface{} {
	panic(js.Error{"No *print-fn* fn set for evaluation environment"})
}

var STAR_print_newline_STAR_ = true

func Enable_console_print_BANG_() interface{} {
	STAR_print_newline_STAR_ = false
	STAR_print_fn_STAR_ = func(x interface{}) interface{} {
		js.Console.Log(x)
		return nil
	}
	return nil
}

var STAR_main_cli_fn_STAR_ interface{}

func pr_opts() interface{} {
	return nil
}

func Newline(opts interface{}) interface{} {
	if STAR_print_newline_STAR_ {
		STAR_print_fn_STAR_("\n")
	}
	return nil
}

func Println(objs ...interface{}) interface{} {
	STAR_print_fn_STAR_(fmt.Sprint(objs...))
	Newline(pr_opts())
	return nil
}

// Main preamble

func Main() {
	Enable_console_print_BANG_()
	var args = make([]interface{}, len(os.Args[1:]))
	for i, a := range os.Args[1:] {
		args[i] = a
	}
	STAR_main_cli_fn_STAR_.(func(...interface{}) interface{})(args...)
}

// IFn

type CljsCoreIFn_InvokeArityVariadic func(...interface{}) interface{}
type CljsCoreIFn_InvokeArity0 func() interface{}
type CljsCoreIFn_InvokeArity1 func(interface{}) interface{}
type CljsCoreIFn_InvokeArity2 func(_, _ interface{}) interface{}
type CljsCoreIFn_InvokeArity3 func(_, _, _ interface{}) interface{}
type CljsCoreIFn_InvokeArity4 func(_, _, _, _ interface{}) interface{}

// There's a protocol called cljs.core.IFn we need to cooperate with, so we use AFn for now.

type AFn struct {
	CljsLangMaxFixedArity int
	CljsCoreIFn_InvokeArityVariadic
	CljsCoreIFn_InvokeArity0
	CljsCoreIFn_InvokeArity1
	CljsCoreIFn_InvokeArity2
	CljsCoreIFn_InvokeArity3
	CljsCoreIFn_InvokeArity4
}

func throwArity(arity int) interface{} {
	panic(js.Error{fmt.Sprint("Invalid arity: ", arity)})
}

func (this AFn) CljsCoreIFn_Invoke(args ...interface{}) interface{} {
	var argc = len(args)
	switch {
	case argc == 0 && this.CljsCoreIFn_InvokeArity0 != nil:
		return this.CljsCoreIFn_InvokeArity0()
	case argc == 1 && this.CljsCoreIFn_InvokeArity1 != nil:
		return this.CljsCoreIFn_InvokeArity1(args[0])
	case argc == 2 && this.CljsCoreIFn_InvokeArity2 != nil:
		return this.CljsCoreIFn_InvokeArity2(args[0], args[1])
	case argc == 3 && this.CljsCoreIFn_InvokeArity3 != nil:
		return this.CljsCoreIFn_InvokeArity3(args[0], args[1], args[3])
	case argc == 4 && this.CljsCoreIFn_InvokeArity4 != nil:
		return this.CljsCoreIFn_InvokeArity4(args[0], args[1], args[3], args[4])
	case argc > this.CljsLangMaxFixedArity && this.CljsCoreIFn_InvokeArityVariadic != nil:
		return this.CljsCoreIFn_InvokeArityVariadic(args...)
	}
	return throwArity(argc)
}

func (this AFn) CljsLangApplyTo(args ...interface{}) interface{} {
	var argc = len(args)
	if argc < 1 {
		throwArity(argc)
	}
	var spread = args[argc-1].([]interface{})
	return this.CljsCoreIFn_Invoke(append(args[:argc-1], spread...)...)
}
