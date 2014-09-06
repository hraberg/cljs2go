// Compiled by ClojureScript to Go 0.0-2322
// cljs.core
package core

import (
	"fmt"
	"reflect"

	"github.com/hraberg/cljs.go/goog"
	"github.com/hraberg/cljs.go/js"
)

var X_STAR_clojurescript_version_STAR_ = "0.0-2322"

var X_STAR_print_length_STAR_ = float64(-1)

var X_STAR_print_level_STAR_ = float64(-1)

// Set *print-fn* to f.
var Set_print_fn_BANG_ *AFn

func init() {
	Set_print_fn_BANG_ = func(set_print_fn_BANG_ *AFn) *AFn {
		return Fn(set_print_fn_BANG_, func(f interface{}) interface{} {
			return func() interface{} {
				var return__808 = f.(*AFn)
				X_STAR_print_fn_STAR_ = return__808
				return return__808
			}()
		})
	}(&AFn{})
}

var Symbol_QMARK_ *AFn

func init() {
	Symbol_QMARK_ = func(symbol_QMARK_ *AFn) *AFn {
		return Fn(symbol_QMARK_, func(x interface{}) bool {
			return func() bool { _, instanceof := x.(*CljsCoreSymbol); return instanceof }()
		})
	}(&AFn{})
}

// Takes a fn f and returns a fn that takes the same arguments as f,
// has the same effects, if any, and returns the opposite truth value.
var Complement *AFn

func init() {
	Complement = func(complement *AFn) *AFn {
		return Fn(complement, func(f interface{}) interface{} {
			return func(complement_fn *AFn) *AFn {
				return Fn(complement_fn, func() interface{} {
					return Not.Arity1IB(f.(CljsCoreIFn).X_invoke_Arity0())
				}, func(x interface{}) interface{} {
					return Not.Arity1IB(f.(CljsCoreIFn).X_invoke_Arity1(x))
				}, func(x interface{}, y interface{}) interface{} {
					return Not.Arity1IB(f.(CljsCoreIFn).X_invoke_Arity2(x, y))
				}, func(x_y_zs__ ...interface{}) interface{} {
					var x = x_y_zs__[0]
					var y = x_y_zs__[1]
					var zs = Array_seq.X_invoke_Arity1(x_y_zs__[2:])
					_, _, _ = x, y, zs
					return Not.Arity1IB(Apply.X_invoke_Arity4(f, x, y, zs))
				})
			}(&AFn{})
		})
	}(&AFn{})
}

var Quote_string *AFn

func init() {
	Quote_string = func(quote_string *AFn) *AFn {
		return Fn(quote_string, func(s interface{}) interface{} {
			return ("\"" + Str.X_invoke_Arity1(Native_invoke_instance_method.X_invoke_Arity3(s, "Replace", []interface{}{(&js.RegExp{"[\\\\\"\b\f\n\r\t]", "g"}), func(G__809 *AFn) *AFn {
				return Fn(G__809, func(match interface{}) interface{} {
					return Char_escapes.(map[string]interface{})[match.(string)]
				})
			}(&AFn{})})).(string) + "\"")
		})
	}(&AFn{})
}

// Prefer this to pr-seq, because it makes the printing function
// configurable, allowing efficient implementations such as appending
// to a StringBuffer.
var Pr_writer *AFn

func init() {
	Pr_writer = func(pr_writer *AFn) *AFn {
		return Fn(pr_writer, func(obj interface{}, writer interface{}, opts interface{}) interface{} {
			if Nil_(obj) {
				return writer.(CljsCoreIWriter).X_write_Arity2("nil")
			} else {
				if Truth_(func() interface{} {
					var and__6309__auto__ = Get.X_invoke_Arity2(opts, (&CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}))
					_ = and__6309__auto__
					if Truth_(and__6309__auto__) {
						{
							var and__6309__auto_____1 = Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IMeta", Str: "cljs.core/IMeta", X_hash: float64(-1459057517), X_meta: nil}), obj)
							_ = and__6309__auto_____1
							if Truth_(and__6309__auto_____1) {
								return Meta.X_invoke_Arity1(obj)
							} else {
								return and__6309__auto_____1
							}
						}
					} else {
						return and__6309__auto__
					}
				}()) {
					writer.(CljsCoreIWriter).X_write_Arity2("^")
					pr_writer.X_invoke_Arity3(Meta.X_invoke_Arity1(obj), writer, opts)
					writer.(CljsCoreIWriter).X_write_Arity2(" ")
				} else {
				}
				if Nil_(obj) {
					return writer.(CljsCoreIWriter).X_write_Arity2("nil")
				} else {
					if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IPrintWithWriter", Str: "cljs.core/IPrintWithWriter", X_hash: float64(1349251417), X_meta: nil}), obj)) {
						return obj.(CljsCoreIPrintWithWriter).X_pr_writer_Arity3(writer, opts)
					} else {
						if (reflect.ValueOf(obj).Kind() == reflect.Bool) || (reflect.ValueOf(obj).Kind() == reflect.Float64) {
							return writer.(CljsCoreIWriter).X_write_Arity2((`` + Str.X_invoke_Arity1(obj).(string)))
						} else {
							if reflect.ValueOf(obj).Kind() == reflect.Slice {
								return Pr_sequential_writer.X_invoke_Arity7(writer, pr_writer, "#js [", " ", "]", opts, obj)
							} else {
								if Truth_(Native_invoke_func.X_invoke_Arity2(goog.IsString, []interface{}{obj})) {
									if Truth_((&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}).X_invoke_Arity1(opts)) {
										return writer.(CljsCoreIWriter).X_write_Arity2(Quote_string.X_invoke_Arity1(obj).(string))
									} else {
										return writer.(CljsCoreIWriter).X_write_Arity2(obj)
									}
								} else {
									if Fn_QMARK_.Arity1IB(obj) {
										return Write_all.X_invoke_ArityVariadic(writer, "#<", (`` + Str.X_invoke_Arity1(obj).(string)), ">")
									} else {
										if func() bool { _, instanceof := obj.(*js.Date); return instanceof }() {
											{
												var normalize = func(G__810 *AFn) *AFn {
													return Fn(G__810, func(n interface{}, len interface{}) interface{} {
														{
															var ns = (`` + Str.X_invoke_Arity1(n).(string))
															_ = ns
															for {
																if Count.X_invoke_Arity1(ns).(float64) < len.(float64) {
																	ns = ("0" + Str.X_invoke_Arity1(ns).(string))
																	continue
																} else {
																	return ns
																}
															}
														}
													})
												}(&AFn{})
												_ = normalize
												return Write_all.X_invoke_ArityVariadic(writer, "#inst \"", (`` + Str.X_invoke_Arity1(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCFullYear", []interface{}{})).(string)), "-", normalize.X_invoke_Arity2((Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCMonth", []interface{}{}).(float64)+float64(1)), float64(2)).(string), "-", normalize.X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCDate", []interface{}{}), float64(2)).(string), "T", normalize.X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCHours", []interface{}{}), float64(2)).(string), ":", normalize.X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCMinutes", []interface{}{}), float64(2)).(string), ":", normalize.X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCSeconds", []interface{}{}), float64(2)).(string), ".", normalize.X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCMilliseconds", []interface{}{}), float64(3)).(string), "-", "00:00\"")
											}
										} else {
											if Truth_(Regexp_QMARK_.X_invoke_Arity1(obj)) {
												return Write_all.X_invoke_ArityVariadic(writer, "#\"", Native_get_instance_field.X_invoke_Arity2(obj, "Source"), "\"")
											} else {
												if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IPrintWithWriter", Str: "cljs.core/IPrintWithWriter", X_hash: float64(1349251417), X_meta: nil}), obj)) {
													return obj.(CljsCoreIPrintWithWriter).X_pr_writer_Arity3(writer, opts)
												} else {
													return Write_all.X_invoke_ArityVariadic(writer, "#<", (`` + Str.X_invoke_Arity1(obj).(string)), ">")

												}
											}
										}
									}
								}
							}
						}
					}
				}

			}
		})
	}(&AFn{})
}

var Pr_sequential_writer *AFn

func init() {
	Pr_sequential_writer = func(pr_sequential_writer *AFn) *AFn {
		return Fn(pr_sequential_writer, func(writer interface{}, print_one interface{}, begin interface{}, sep interface{}, end interface{}, opts interface{}, coll interface{}) interface{} {
			{
				var _STAR_print_level_STAR_812 = X_STAR_print_level_STAR_
				_ = _STAR_print_level_STAR_812
				return func() interface{} {
					defer func() {
						X_STAR_print_level_STAR_ = _STAR_print_level_STAR_812

					}()
					{
						X_STAR_print_level_STAR_ = func() float64 {
							if float64(-1) == X_STAR_print_level_STAR_ {
								return float64(-1)
							} else {
								return (X_STAR_print_level_STAR_ - float64(1))
							}
						}()

						if (!(Nil_(X_STAR_print_level_STAR_))) && (X_STAR_print_level_STAR_ < float64(0)) {
							return writer.(CljsCoreIWriter).X_write_Arity2("#")
						} else {
							writer.(CljsCoreIWriter).X_write_Arity2(begin)
							if Truth_(Seq.Arity1IQ(coll)) {
								print_one.(CljsCoreIFn).X_invoke_Arity3(First.X_invoke_Arity1(coll), writer, opts)
							} else {
							}
							{
								var coll_813___1 = Next.Arity1IQ(coll)
								var n_814 = ((&CljsCoreKeyword{Ns: nil, Name: "print-length", Fqn: "print-length", X_hash: float64(1931866356)}).X_invoke_Arity1(opts).(float64) - float64(1))
								_, _ = coll_813___1, n_814
								for {
									if Truth_(func() interface{} {
										var and__6309__auto__ = coll_813___1
										_ = and__6309__auto__
										if Truth_(and__6309__auto__) {
											return (Nil_(n_814)) || (!(n_814 == float64(0)))
										} else {
											return and__6309__auto__
										}
									}()) {
										writer.(CljsCoreIWriter).X_write_Arity2(sep)
										print_one.(CljsCoreIFn).X_invoke_Arity3(First.X_invoke_Arity1(coll_813___1), writer, opts)
										coll_813___1, n_814 = Next.Arity1IQ(coll_813___1), (n_814 - float64(1))
										continue
									} else {
										if Truth_(func() interface{} {
											var and__6309__auto__ = Seq.Arity1IQ(coll_813___1)
											_ = and__6309__auto__
											if Truth_(and__6309__auto__) {
												return (n_814 == float64(0))
											} else {
												return and__6309__auto__
											}
										}()) {
											writer.(CljsCoreIWriter).X_write_Arity2(sep)
											writer.(CljsCoreIWriter).X_write_Arity2("...")
										} else {
										}
									}
									break
								}
							}
							return writer.(CljsCoreIWriter).X_write_Arity2(end)
						}
					}
				}()
			}
		})
	}(&AFn{})
}

var Type_ *AFn

func init() {
	Type_ = func(type_ *AFn) *AFn {
		return Fn(type_, func(x interface{}) interface{} {
			if Nil_(x) {
				return nil
			} else {
				return reflect.TypeOf(x)
			}
		})
	}(&AFn{})
}

var Type__GT_str *AFn

func init() {
	Type__GT_str = func(type__GT_str *AFn) *AFn {
		return Fn(type__GT_str, func(ty interface{}) interface{} {
			return (`` + Str.X_invoke_Arity1(ty).(string))
		})
	}(&AFn{})
}

// Returns true if n is an integer.
var Integer_QMARK_ *AFn

func init() {
	Integer_QMARK_ = func(integer_QMARK_ *AFn) *AFn {
		return Fn(integer_QMARK_, func(n interface{}) bool {
			return (reflect.ValueOf(n).Kind() == reflect.Float64) && (!(Truth_(Native_invoke_func.X_invoke_Arity2(js.IsNaN, []interface{}{n})))) && (!(reflect.DeepEqual(n, js.Infinity))) && (n.(float64) == float64(int(n.(float64))))
		})
	}(&AFn{})
}

// Creates a new javascript array.
// @param {...*} var_args
// @param {...*} var_args
var Array *AFn

func init() {
	Array = func(array *AFn) *AFn {
		return Fn(array, func(items__ ...interface{}) interface{} {
			var items = Array_seq.X_invoke_Arity1(items__[0:])
			_ = items
			return Into_array.Arity1IA(items)
		})
	}(&AFn{})
}

var Make_array *AFn

func init() {
	Make_array = func(make_array *AFn) *AFn {
		return Fn(make_array, func(size interface{}) []interface{} {
			return make_array.Arity2IIA(nil, size)
		}, func(type_ interface{}, size interface{}) []interface{} {
			return make([]interface{}, int(size.(float64)))
		})
	}(&AFn{})
}

func init() {
	Char = func(char *AFn) *AFn {
		return Fn(char, func(x interface{}) interface{} {
			if reflect.ValueOf(x).Kind() == reflect.Float64 {
				return js.String.FromCharCode(x.(float64))
			} else {
				if (reflect.ValueOf(x).Kind() == reflect.String) && (Native_get_instance_field.X_invoke_Arity2(x, "Length").(float64) == float64(1)) {
					return x
				} else {
					panic((&js.Error{"Argument to char must be a character or number"}))

				}
			}
		})
	}(&AFn{})
}

var String_hash_cache = map[string]interface{}{}
var Add_to_string_hash_cache *AFn

func init() {
	Add_to_string_hash_cache = func(add_to_string_hash_cache *AFn) *AFn {
		return Fn(add_to_string_hash_cache, func(k interface{}) interface{} {
			{
				var h = Hash_string_STAR_.X_invoke_Arity1(k).(float64)
				_ = h
				String_hash_cache[k.(string)] = h
				String_hash_cache_count = (String_hash_cache_count + float64(1))

				return h
			}
		})
	}(&AFn{})
}

var Hash_string *AFn

func init() {
	Hash_string = func(hash_string *AFn) *AFn {
		return Fn(hash_string, func(k interface{}) interface{} {
			if String_hash_cache_count > float64(255) {
				String_hash_cache = map[string]interface{}{}

				String_hash_cache_count = float64(0)

			} else {
			}
			{
				var h = String_hash_cache[k.(string)]
				_ = h
				if reflect.ValueOf(h).Kind() == reflect.Float64 {
					return h
				} else {
					return Add_to_string_hash_cache.X_invoke_Arity1(k).(float64)
				}
			}
		})
	}(&AFn{})
}

type CljsCoreT350 struct{}

func (_ *CljsCoreT350) CljsCoreObject__() {}
func (self__ *CljsCoreT350) HasNext() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return false
	}
}

func (self__ *CljsCoreT350) Next() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&js.Error{"No such element"})
	}
}

func (self__ *CljsCoreT350) Remove() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&js.Error{"Unsupported operation"})
	}
}

var X__GT_T350 *AFn

func init() {
	X__GT_T350 = func(__GT_T350 *AFn) *AFn {
		return Fn(__GT_T350, func() interface{} {
			return (&CljsCoreT350{})
		})
	}(&AFn{})
}

var Nil_iter *AFn

func init() {
	Nil_iter = func(nil_iter *AFn) *AFn {
		return Fn(nil_iter, func() interface{} {
			return (&CljsCoreT350{})
		})
	}(&AFn{})
}

// Set *print-fn* to console.log
var Enable_console_print_BANG_ *AFn

func init() {
	Enable_console_print_BANG_ = func(enable_console_print_BANG_ *AFn) *AFn {
		return Fn(enable_console_print_BANG_, func() interface{} {
			X_STAR_print_newline_STAR_ = false

			return func() interface{} {
				var return__815 = func(fmt_println *AFn) *AFn {
					return Fn(fmt_println, func(x interface{}) interface{} {
						fmt.Println(x)
						return nil
					})
				}(&AFn{})
				X_STAR_print_fn_STAR_ = return__815
				return return__815
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
			return f.(*AFn).Call(Into_array.Arity1IA(args)...)
		}, func(f interface{}, x interface{}, args interface{}) interface{} {
			return f.(*AFn).Call(append([]interface{}{x}, Into_array.Arity1IA(args)...)...)
		}, func(f interface{}, x interface{}, y interface{}, args interface{}) interface{} {
			return f.(*AFn).Call(append([]interface{}{x, y}, Into_array.Arity1IA(args)...)...)
		}, func(f interface{}, x interface{}, y interface{}, z interface{}, args interface{}) interface{} {
			return f.(*AFn).Call(append([]interface{}{x, y, z}, Into_array.Arity1IA(args)...)...)
		}, func(f_a_b_c_d_args__ ...interface{}) interface{} {
			var f = f_a_b_c_d_args__[0]
			var a = f_a_b_c_d_args__[1]
			var b = f_a_b_c_d_args__[2]
			var c = f_a_b_c_d_args__[3]
			var d = f_a_b_c_d_args__[4]
			var args = Array_seq.X_invoke_Arity1(f_a_b_c_d_args__[5:])
			_, _, _, _, _, _ = f, a, b, c, d, args
			{
				var arr = Into_array.Arity1IA(Seq_(Butlast.X_invoke_Arity1(args)))
				var varargs = Into_array.Arity1IA(Last.X_invoke_Arity1(args))
				_, _ = arr, varargs
				return f.(*AFn).Call(append([]interface{}{a, b, c, d}, append(arr, varargs...)...)...)
			}
		})
	}(&AFn{})
}

// Internal - do not use!
var Native_satisfies_QMARK_ *AFn

func init() {
	Native_satisfies_QMARK_ = func(native_satisfies_QMARK_ *AFn) *AFn {
		return Fn(native_satisfies_QMARK_, func(p interface{}, x interface{}) bool {
			return value(decorate(x)).Type().Implements(protocols[(`` + Str.X_invoke_Arity1(p).(string))])
		})
	}(&AFn{})
}
