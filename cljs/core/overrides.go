package core

import (
	"fmt"
	"reflect"
)

var Enable_console_print_BANG_ = Fn(func() interface{} {
	X_STAR_print_fn_STAR_ = Fn(func(x interface{}) interface{} {
		fmt.Print(x)
		return nil
	})
	return nil
})

var Apply = Fn(func(f_args ...interface{}) interface{} {
	f, args := f_args[0], f_args[1:]
	argc := len(args)
	if argc < 1 {
		throwArity(nil, argc)
	}
	var spread = args[argc-1].([]interface{}) // This will be a seq in real life.
	return f.(*AFn).Call(append(args[:argc-1], spread...)...)
})

var Native_satisfies_QMARK_ = Fn(func(p, x interface{}) bool {
	return reflect.ValueOf(x).Type().Implements(protocols[fmt.Sprint(p)])
})

// unimplemented
var Set_print_fn_BANG_, Symbol_QMARK_, Complement, Iter, Pr_writer, Type_, Type__GT_str,
	Add_to_string_hash_cache, Hash_string, Fn_QMARK_, Integer_QMARK_, Array, Re_pattern, Nil_iter *AFn

var X_STAR_print_length_STAR_ = -1.0
var X_STAR_print_level_STAR_ = -1.0

var String_hash_cache = map[string]interface{}{}
