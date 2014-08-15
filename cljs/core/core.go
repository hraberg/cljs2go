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

type IFn interface {
	Invoke(...interface{}) interface{}
}

func assertArity(n int, args []interface{}) {
	var argc = len(args)
	if n != argc {
		panic(js.Error{fmt.Sprint("Invalid arity: ", argc)})
	}
}

type InvokeArityVariadic func(...interface{}) interface{}

func (this InvokeArityVariadic) Invoke(args ...interface{}) interface{} {
	return this(args...)
}

type InvokeArity0 func() interface{}

func (this InvokeArity0) Invoke(args ...interface{}) interface{} {
	assertArity(0, args)
	return this()
}

type InvokeArity1 func(interface{}) interface{}

func (this InvokeArity1) Invoke(args ...interface{}) interface{} {
	assertArity(1, args)
	return this(args[0])
}

type InvokeArity2 func(_, _ interface{}) interface{}

func (this InvokeArity2) Invoke(args ...interface{}) interface{} {
	assertArity(2, args)
	return this(args[0], args[1])
}

type InvokeArity3 func(_, _, _ interface{}) interface{}

func (this InvokeArity3) Invoke(args ...interface{}) interface{} {
	assertArity(3, args)
	return this(args[0], args[1], args[2])
}

type InvokeArity4 func(_, _, _, _ interface{}) interface{}

func (this InvokeArity4) Invoke(args ...interface{}) interface{} {
	assertArity(4, args)
	return this(args[0], args[1], args[2], args[3])
}

type CljsCoreIFn struct {
	MaxFixedArity int
	InvokeArityVariadic
	InvokeArity0
	InvokeArity1
	InvokeArity2
	InvokeArity3
	InvokeArity4
}

func (this CljsCoreIFn) Invoke(args ...interface{}) interface{} {
	var argc = len(args)
	switch {
	case argc == 0 && this.InvokeArity0 != nil:
		return this.InvokeArity0()
	case argc == 1 && this.InvokeArity1 != nil:
		return this.InvokeArity1(args[0])
	case argc == 2 && this.InvokeArity2 != nil:
		return this.InvokeArity2(args[0], args[1])
	case argc == 3 && this.InvokeArity3 != nil:
		return this.InvokeArity3(args[0], args[1], args[3])
	case argc == 4 && this.InvokeArity4 != nil:
		return this.InvokeArity4(args[0], args[1], args[3], args[4])
	case argc > this.MaxFixedArity && this.InvokeArityVariadic != nil:
		return this.InvokeArityVariadic(args...)
	}
	panic(js.Error{fmt.Sprint("Invalid arity: ", argc)})
}

func (this CljsCoreIFn) ApplyTo(args ...interface{}) interface{} {
	var argc = len(args)
	if argc < 1 {
		panic(js.Error{fmt.Sprint("Invalid arity: ", argc)})
	}
	args = append(args[:argc-1], args[argc-1].([]interface{})...)
	return this.Invoke(args...)
}
