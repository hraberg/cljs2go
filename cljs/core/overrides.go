// Compiled by ClojureScript to Go 0.0-2322
// cljs.core
package core

import "fmt"

var X_STAR_clojurescript_version_STAR_ = "0.0-2322"

var X_STAR_print_length_STAR_ = float64(-1)

var X_STAR_print_level_STAR_ = float64(-1)
var Set_print_fn_BANG_ *AFn

var Symbol_QMARK_ *AFn

var Complement *AFn

var Iter *AFn

var Pr_writer *AFn

var Pr_sequential_writer *AFn

var Type_ *AFn

var Type__GT_str *AFn

var Add_string_to_hash_cache *AFn

var Hash_string *AFn

var Integer_QMARK_ *AFn

var Array *AFn

var Nil_iter *AFn

// Set *print-fn* to console.log
var Enable_console_print_BANG_ *AFn

func init() {
	Enable_console_print_BANG_ = func(enable_console_print_BANG_ *AFn) *AFn {
		return Fn(enable_console_print_BANG_, func() interface{} {
			return func() interface{} {
				var return__808 = func(fmt_println *AFn) *AFn {
					return Fn(fmt_println, func(x interface{}) interface{} {
						fmt.Println(x)
						return nil
					})
				}(&AFn{})
				X_STAR_print_fn_STAR_ = return__808
				return return__808
			}()
		})
	}(&AFn{})
}

// Applies fn f to the argument list formed by prepending intervening arguments to args.
// First cut.  Not lazy.  Needs to use emitted toApply.
// @param {...*} var_args
var Apply *AFn

func init() {
	Apply = func(apply *AFn) *AFn {
		return Fn(apply, func(f interface{}, args interface{}) interface{} {
			return f.(*AFn).Call(Into_array.X_invoke_Arity1(args).([]interface{})...)
		}, func(f interface{}, x interface{}, args interface{}) interface{} {
			return f.(*AFn).Call(append([]interface{}{x}, Into_array.X_invoke_Arity1(args).([]interface{})...)...)
		}, func(f interface{}, x interface{}, y interface{}, args interface{}) interface{} {
			return f.(*AFn).Call(append([]interface{}{x, y}, Into_array.X_invoke_Arity1(args).([]interface{})...)...)
		}, func(f interface{}, x interface{}, y interface{}, z interface{}, args interface{}) interface{} {
			return f.(*AFn).Call(append([]interface{}{x, y, z}, Into_array.X_invoke_Arity1(args).([]interface{})...)...)
		}, func(f_a_b_c_d_args__ ...interface{}) interface{} {
			var f = f_a_b_c_d_args__[0]
			var a = f_a_b_c_d_args__[1]
			var b = f_a_b_c_d_args__[2]
			var c = f_a_b_c_d_args__[3]
			var d = f_a_b_c_d_args__[4]
			var args = Array_seq.X_invoke_Arity1(f_a_b_c_d_args__[5:])
			_, _, _, _, _, _ = f, a, b, c, d, args
			{
				var arr = Into_array.X_invoke_Arity1(Butlast.X_invoke_Arity1(args)).([]interface{})
				var varargs = Into_array.X_invoke_Arity1(Last.X_invoke_Arity1(args)).([]interface{})
				_, _ = arr, varargs
				return f.(*AFn).Call(append([]interface{}{a, b, c, d}, append(arr, varargs...))...)
			}
		})
	}(&AFn{})
}

// Internal - do not use!
var Native_satisfies_QMARK_ *AFn

func init() {
	Native_satisfies_QMARK_ = func(native_satisfies_QMARK_ *AFn) *AFn {
		return Fn(native_satisfies_QMARK_, func(p interface{}, x interface{}) bool {
			return value(decorate(x)).Type().Implements(protocols[fmt.Sprint(p)])
		})
	}(&AFn{})
}

var String_hash_cache = map[string]interface{}{}
