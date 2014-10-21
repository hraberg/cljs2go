// Compiled by ClojureScript to Go 0.0-2371
// cljs.core

// Go overrides.
package core

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/hraberg/cljs.go/goog"
	goog_array "github.com/hraberg/cljs.go/goog/array"
	"github.com/hraberg/cljs.go/js"
	"github.com/hraberg/cljs.go/js/Math"
)

func init() {
	X_STAR_clojurescript_version_STAR_ = "0.0-2371"

	X_STAR_print_length_STAR_ = js.NaN

	X_STAR_print_level_STAR_ = js.NaN

	Set_print_fn_BANG_ = func(set_print_fn_BANG_ *AFn) *AFn {
		return Fn(set_print_fn_BANG_, 1, func(f interface{}) interface{} {
			return func() interface{} {
				var return__7959 = f.(*AFn)
				X_STAR_print_fn_STAR_ = return__7959
				return return__7959
			}()
		})
	}(&AFn{})

	Symbol_QMARK_ = func(symbol_QMARK_ *AFn) *AFn {
		return Fn(symbol_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().AssignableTo(reflect.TypeOf((**CljsCoreSymbol)(nil)).Elem())
		})
	}(&AFn{})

	Complement = func(complement *AFn) *AFn {
		return Fn(complement, 1, func(f interface{}) interface{} {
			return func(complement_fn *AFn) *AFn {
				return Fn(complement_fn, 2, func() interface{} {
					return Not.Arity1IB(func() interface{} {
						return f.(CljsCoreIFn).X_invoke_Arity0()
					}())
				}, func(x interface{}) interface{} {
					return Not.Arity1IB(func() interface{} {
						var G__7969 = x
						_ = G__7969
						return f.(CljsCoreIFn).X_invoke_Arity1(G__7969)
					}())
				}, func(x interface{}, y interface{}) interface{} {
					return Not.Arity1IB(func() interface{} {
						var G__7970 = x
						var G__7971 = y
						_, _ = G__7970, G__7971
						return f.(CljsCoreIFn).X_invoke_Arity2(G__7970, G__7971)
					}())
				}, func(x_y_zs__ ...interface{}) interface{} {
					var x = x_y_zs__[0]
					var y = x_y_zs__[1]
					var zs = Seq.Arity1IQ(x_y_zs__[2])
					_, _, _ = x, y, zs
					return Not.Arity1IB(Apply.X_invoke_Arity4(f, x, y, zs))
				})
			}(&AFn{})
		})
	}(&AFn{})

	Remove = func(remove *AFn) *AFn {
		return Fn(remove, 2, func(pred interface{}) interface{} {
			return Filter.X_invoke_Arity1(Complement.X_invoke_Arity1(pred).(CljsCoreIFn)).(CljsCoreIFn)
		}, func(pred interface{}, coll interface{}) interface{} {
			return Filter.X_invoke_Arity2(Complement.X_invoke_Arity1(pred).(CljsCoreIFn), coll).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Identity = func(identity *AFn) *AFn {
		return Fn(identity, 1, func(x interface{}) interface{} {
			return x
		}, func(x____ ...interface{}) interface{} {
			var x = x____[0]
			var ___ = Seq.Arity1IQ(x____[1])
			_, _ = x, ___
			return x
		})
	}(&AFn{})

	Rand = func(rand *AFn) *AFn {
		return Fn(rand, 1, func() float64 {
			return rand.Arity1IF(float64(1))
		}, func(n interface{}) float64 {
			return (func() interface{} {
				return Native_invoke_func.X_invoke_Arity2(Math.Random, []interface{}{})
			}().(float64) * n.(float64))
		})
	}(&AFn{})

	X_EQ_ = func(_EQ_ *AFn) *AFn {
		return Fn(_EQ_, 2, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			if Nil_(x) {
				return Nil_(y)
			} else {
				{
					var or__177__auto__ = reflect.DeepEqual(x, y)
					_ = or__177__auto__
					if Truth_(or__177__auto__) {
						return or__177__auto__
					} else {
						if DecoratedValue_(x).Type().Implements(reflect.TypeOf((*CljsCoreIEquiv)(nil)).Elem()) {
							return Decorate_(x).(CljsCoreIEquiv).X_equiv_Arity2(y)
						} else {
							return false
						}
					}
				}
			}
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			for {
				if _EQ_.Arity2IIB(x, y) {
					if Truth_(Next.Arity1IQ(more)) {
						x, y, more = y, First.X_invoke_Arity1(more), Next.Arity1IQ(more)
						continue
					} else {
						return _EQ_.Arity2IIB(y, First.X_invoke_Arity1(more))
					}
				} else {
					return false
				}
			}
		})
	}(&AFn{})

	Sort = func(sort *AFn) *AFn {
		return Fn(sort, 2, func(coll interface{}) interface{} {
			return sort.X_invoke_Arity2(Compare, coll)
		}, func(comp interface{}, coll interface{}) interface{} {
			if Truth_(Seq.Arity1IQ(coll)) {
				{
					var a = To_array.X_invoke_Arity1(coll).([]interface{})
					var comp___1 = Fn__GT_comparator.X_invoke_Arity1(comp).(CljsCoreIFn)
					_, _ = a, comp___1
					{
						var G__7981_7983 = a
						var G__7982_7984 = func(x, y interface{}) interface{} { return comp___1.X_invoke_Arity2(x, y) }
						_, _ = G__7981_7983, G__7982_7984
						Native_invoke_func.X_invoke_Arity2(goog_array.StableSort, []interface{}{G__7981_7983, G__7982_7984})
					}
					return Seq.Arity1IQ(a)
				}
			} else {
				return CljsCoreIEmptyList(CljsCoreList_EMPTY)
			}
		})
	}(&AFn{})

	Get = func(get *AFn) *AFn {
		return Fn(get, 3, func(o interface{}, k interface{}) interface{} {
			if Nil_(o) {
				return nil
			} else {
				if DecoratedValue_(o).Type().Implements(reflect.TypeOf((*CljsCoreILookup)(nil)).Elem()) {
					return Decorate_(o).(CljsCoreILookup).X_lookup_Arity2(k)
				} else {
					if Value_(o).Kind() == reflect.Slice {
						if (Value_(k).Kind() == reflect.Float64) && (k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64)) {
							return (o.([]interface{})[int(k.(float64))])
						} else {
							return nil
						}
					} else {
						if Value_(o).Kind() == reflect.String {
							if (Value_(k).Kind() == reflect.Float64) && (k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64)) {
								return string((o.(string)[int(k.(float64))]))
							} else {
								return nil
							}
						} else {
							if DecoratedValue_(o).Type().Implements(reflect.TypeOf((*CljsCoreILookup)(nil)).Elem()) {
								return Decorate_(o).(CljsCoreILookup).X_lookup_Arity2(k)
							} else {
								return nil

							}
						}
					}
				}
			}
		}, func(o interface{}, k interface{}, not_found interface{}) interface{} {
			if !(Nil_(o)) {
				if DecoratedValue_(o).Type().Implements(reflect.TypeOf((*CljsCoreILookup)(nil)).Elem()) {
					return Decorate_(o).(CljsCoreILookup).X_lookup_Arity3(k, not_found)
				} else {
					if Value_(o).Kind() == reflect.Slice {
						if (Value_(k).Kind() == reflect.Float64) && (k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64)) {
							return (o.([]interface{})[int(k.(float64))])
						} else {
							return not_found
						}
					} else {
						if Value_(o).Kind() == reflect.String {
							if (Value_(k).Kind() == reflect.Float64) && (k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64)) {
								return string((o.(string)[int(k.(float64))]))
							} else {
								return not_found
							}
						} else {
							if DecoratedValue_(o).Type().Implements(reflect.TypeOf((*CljsCoreILookup)(nil)).Elem()) {
								return Decorate_(o).(CljsCoreILookup).X_lookup_Arity3(k, not_found)
							} else {
								return not_found

							}
						}
					}
				}
			} else {
				return not_found
			}
		})
	}(&AFn{})

	Quote_string = func(quote_string *AFn) *AFn {
		return Fn(quote_string, 1, func(s interface{}) interface{} {
			return strconv.Quote(s.(string))
		})
	}(&AFn{})

	Pr_writer = func(pr_writer *AFn) *AFn {
		return Fn(pr_writer, 3, func(obj interface{}, writer interface{}, opts interface{}) interface{} {
			if Nil_(obj) {
				return Decorate_(writer).(CljsCoreIWriter).X_write_Arity2("nil")
			} else {
				if Truth_(func() interface{} {
					var and__165__auto__ = Get.X_invoke_Arity2(opts, (&CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}))
					_ = and__165__auto__
					if Truth_(and__165__auto__) {
						{
							var and__165__auto_____1 = DecoratedValue_(obj).Type().Implements(reflect.TypeOf((*CljsCoreIMeta)(nil)).Elem())
							_ = and__165__auto_____1
							if Truth_(and__165__auto_____1) {
								return Meta.X_invoke_Arity1(obj)
							} else {
								return and__165__auto_____1
							}
						}
					} else {
						return and__165__auto__
					}
				}()) {
					Decorate_(writer).(CljsCoreIWriter).X_write_Arity2("^")
					pr_writer.X_invoke_Arity3(Meta.X_invoke_Arity1(obj), writer, opts)
					Decorate_(writer).(CljsCoreIWriter).X_write_Arity2(" ")
				} else {
				}
				if Nil_(obj) {
					return Decorate_(writer).(CljsCoreIWriter).X_write_Arity2("nil")
				} else {
					if DecoratedValue_(obj).Type().Implements(reflect.TypeOf((*CljsCoreIPrintWithWriter)(nil)).Elem()) {
						return Decorate_(obj).(CljsCoreIPrintWithWriter).X_pr_writer_Arity3(writer, opts)
					} else {
						if (Value_(obj).Kind() == reflect.Bool) || (Value_(obj).Kind() == reflect.Float64) {
							return Decorate_(writer).(CljsCoreIWriter).X_write_Arity2((`` + Str.X_invoke_Arity1(obj).(string)))
						} else {
							if Value_(obj).Kind() == reflect.Slice {
								return Pr_sequential_writer.X_invoke_Arity7(writer, pr_writer, "#js [", " ", "]", opts, obj)
							} else {
								if Truth_(func() interface{} {
									var G__7989 = obj
									_ = G__7989
									return Native_invoke_func.X_invoke_Arity2(goog.IsString, []interface{}{G__7989})
								}()) {
									if Truth_((&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}).X_invoke_Arity1(opts)) {
										return Decorate_(writer).(CljsCoreIWriter).X_write_Arity2(Quote_string.X_invoke_Arity1(obj).(string))
									} else {
										return Decorate_(writer).(CljsCoreIWriter).X_write_Arity2(obj)
									}
								} else {
									if Fn_QMARK_.Arity1IB(obj) {
										return Write_all.X_invoke_ArityVariadic(writer, Array_seq.X_invoke_Arity1([]interface{}{"#<", (`` + Str.X_invoke_Arity1(obj).(string)), ">"}))
									} else {
										if Value_(obj).Type().AssignableTo(reflect.TypeOf((**js.Date)(nil)).Elem()) {
											{
												var normalize = func(G__7990 *AFn) *AFn {
													return Fn(G__7990, 2, func(n interface{}, len interface{}) interface{} {
														{
															var ns interface{} = (`` + Str.X_invoke_Arity1(n).(string))
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
												return Write_all.X_invoke_ArityVariadic(writer, Array_seq.X_invoke_Arity1([]interface{}{"#inst \"", (`` + Str.X_invoke_Arity1(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCFullYear", []interface{}{})).(string)), "-", normalize.X_invoke_Arity2((Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCMonth", []interface{}{}).(float64) + float64(1)), float64(2)).(string), "-", normalize.X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCDate", []interface{}{}), float64(2)).(string), "T", normalize.X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCHours", []interface{}{}), float64(2)).(string), ":", normalize.X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCMinutes", []interface{}{}), float64(2)).(string), ":", normalize.X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCSeconds", []interface{}{}), float64(2)).(string), ".", normalize.X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCMilliseconds", []interface{}{}), float64(3)).(string), "-", "00:00\""}))
											}
										} else {
											if Truth_(Regexp_QMARK_.X_invoke_Arity1(obj)) {
												return Write_all.X_invoke_ArityVariadic(writer, Array_seq.X_invoke_Arity1([]interface{}{"#\"", Native_get_instance_field.X_invoke_Arity2(obj, "Pattern"), "\""}))
											} else {
												if DecoratedValue_(obj).Type().Implements(reflect.TypeOf((*CljsCoreIPrintWithWriter)(nil)).Elem()) {
													return Decorate_(obj).(CljsCoreIPrintWithWriter).X_pr_writer_Arity3(writer, opts)
												} else {
													return Write_all.X_invoke_ArityVariadic(writer, Array_seq.X_invoke_Arity1([]interface{}{"#<", (`` + Str.X_invoke_Arity1(obj).(string)), ">"}))

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

	Pr_sequential_writer = func(pr_sequential_writer *AFn) *AFn {
		return Fn(pr_sequential_writer, 7, func(writer interface{}, print_one interface{}, begin interface{}, sep interface{}, end interface{}, opts interface{}, coll interface{}) interface{} {
			{
				var _STAR_print_level_STAR_7998 = X_STAR_print_level_STAR_
				_ = _STAR_print_level_STAR_7998
				return func() interface{} {
					defer func() {
						X_STAR_print_level_STAR_ = _STAR_print_level_STAR_7998

					}()
					{
						X_STAR_print_level_STAR_ = func() float64 {
							if js.NaN == X_STAR_print_level_STAR_ {
								return X_STAR_print_level_STAR_
							} else {
								return (X_STAR_print_level_STAR_ - float64(1))
							}
						}()

						if X_STAR_print_level_STAR_ < float64(0) {
							return Decorate_(writer).(CljsCoreIWriter).X_write_Arity2("#")
						} else {
							Decorate_(writer).(CljsCoreIWriter).X_write_Arity2(begin)
							if Truth_(Seq.Arity1IQ(coll)) {
								{
									var G__7999_8005 = First.X_invoke_Arity1(coll)
									var G__8000_8006 = writer
									var G__8001_8007 = opts
									_, _, _ = G__7999_8005, G__8000_8006, G__8001_8007
									print_one.(CljsCoreIFn).X_invoke_Arity3(G__7999_8005, G__8000_8006, G__8001_8007)
								}
							} else {
							}
							{
								var coll_8008___1 interface{} = Next.Arity1IQ(coll)
								var n_8009 = ((&CljsCoreKeyword{Ns: nil, Name: "print-length", Fqn: "print-length", X_hash: float64(1931866356)}).X_invoke_Arity1(opts).(float64) - float64(1))
								_, _ = coll_8008___1, n_8009
								for {
									if Truth_(func() interface{} {
										var and__165__auto__ = coll_8008___1
										_ = and__165__auto__
										if Truth_(and__165__auto__) {
											return (Nil_(n_8009)) || (!(n_8009 == float64(0)))
										} else {
											return and__165__auto__
										}
									}()) {
										Decorate_(writer).(CljsCoreIWriter).X_write_Arity2(sep)
										{
											var G__8002_8010 = First.X_invoke_Arity1(coll_8008___1)
											var G__8003_8011 = writer
											var G__8004_8012 = opts
											_, _, _ = G__8002_8010, G__8003_8011, G__8004_8012
											print_one.(CljsCoreIFn).X_invoke_Arity3(G__8002_8010, G__8003_8011, G__8004_8012)
										}
										coll_8008___1, n_8009 = Next.Arity1IQ(coll_8008___1), (n_8009 - float64(1))
										continue
									} else {
										if Truth_(func() interface{} {
											var and__165__auto__ = Seq.Arity1IQ(coll_8008___1)
											_ = and__165__auto__
											if Truth_(and__165__auto__) {
												return (n_8009 == float64(0))
											} else {
												return and__165__auto__
											}
										}()) {
											Decorate_(writer).(CljsCoreIWriter).X_write_Arity2(sep)
											Decorate_(writer).(CljsCoreIWriter).X_write_Arity2("...")
										} else {
										}
									}
									break
								}
							}
							return Decorate_(writer).(CljsCoreIWriter).X_write_Arity2(end)
						}
					}
				}()
			}
		})
	}(&AFn{})

	Type_ = func(type_ *AFn) *AFn {
		return Fn(type_, 1, func(x interface{}) interface{} {
			if Nil_(x) {
				return nil
			} else {
				return reflect.TypeOf(x)
			}
		})
	}(&AFn{})

	Type__GT_str = func(type__GT_str *AFn) *AFn {
		return Fn(type__GT_str, 1, func(ty interface{}) interface{} {
			return (`` + Str.X_invoke_Arity1(ty).(string))
		})
	}(&AFn{})

	Integer_QMARK_ = func(integer_QMARK_ *AFn) *AFn {
		return Fn(integer_QMARK_, 1, func(n interface{}) bool {
			return (Value_(n).Kind() == reflect.Float64) && (Not.Arity1IB(func() interface{} {
				var G__8016 = n
				_ = G__8016
				return Native_invoke_func.X_invoke_Arity2(js.IsNaN, []interface{}{G__8016})
			}())) && (!(reflect.DeepEqual(n, js.Infinity))) && (n.(float64) == float64(int(n.(float64))))
		})
	}(&AFn{})

	Array = func(array *AFn) *AFn {
		return Fn(array, 0, func(items__ ...interface{}) interface{} {
			var items = Seq.Arity1IQ(items__[0])
			_ = items
			return Into_array.Arity1IA(items)
		})
	}(&AFn{})

	Make_array = func(make_array *AFn) *AFn {
		return Fn(make_array, 2, func(size interface{}) []interface{} {
			return make_array.Arity2IIA(nil, size)
		}, func(type_ interface{}, size interface{}) []interface{} {
			return make([]interface{}, int(size.(float64)))
		})
	}(&AFn{})

	Char = func(char *AFn) *AFn {
		return Fn(char, 1, func(x interface{}) interface{} {
			if Value_(x).Kind() == reflect.Float64 {
				return js.String.FromCharCode(x.(float64))
			} else {
				if (Value_(x).Kind() == reflect.String) && (Native_get_instance_field.X_invoke_Arity2(x, "Length").(float64) == float64(1)) {
					return x
				} else {
					panic((&js.Error{"Argument to char must be a character or number"}))

				}
			}
		})
	}(&AFn{})

	String_hash_cache = map[interface{}]interface{}{}

	Add_to_string_hash_cache = func(add_to_string_hash_cache *AFn) *AFn {
		return Fn(add_to_string_hash_cache, 1, func(k interface{}) interface{} {
			{
				var h = Hash_string_STAR_.X_invoke_Arity1(k).(float64)
				_ = h
				String_hash_cache[k] = h
				String_hash_cache_count = (String_hash_cache_count + float64(1))

				return h
			}
		})
	}(&AFn{})

	Hash_string = func(hash_string *AFn) *AFn {
		return Fn(hash_string, 1, func(k interface{}) interface{} {
			if String_hash_cache_count > float64(255) {
				String_hash_cache = map[interface{}]interface{}{}

				String_hash_cache_count = float64(0)

			} else {
			}
			{
				var h = String_hash_cache[k]
				_ = h
				if Value_(h).Kind() == reflect.Float64 {
					return h
				} else {
					return Add_to_string_hash_cache.X_invoke_Arity1(k).(float64)
				}
			}
		})
	}(&AFn{})

	Enable_console_print_BANG_ = func(enable_console_print_BANG_ *AFn) *AFn {
		return Fn(enable_console_print_BANG_, 0, func() interface{} {
			X_STAR_print_newline_STAR_ = false

			return func() interface{} {
				var return__8019 = func(fmt_println *AFn) *AFn {
					return Fn(fmt_println, 1, func(x interface{}) interface{} {
						fmt.Println(x)
						return nil
					})
				}(&AFn{})
				X_STAR_print_fn_STAR_ = return__8019
				return return__8019
			}()
		})
	}(&AFn{})

	Apply = func(apply *AFn) *AFn {
		return Fn(apply, 5, func(f interface{}, args interface{}) interface{} {
			{
				var fixed_arity = MaxFixedArity_(f)
				_ = fixed_arity
				if (X_EQ_.Arity2IIB(float64(-1), fixed_arity)) || (Bounded_count.X_invoke_Arity2(args, (fixed_arity+float64(1))).(float64) <= fixed_arity) {
					return Call_(f.(CljsCoreIFn), Into_array.Arity1IA(args)...)
				} else {
					if Empty_QMARK_.Arity1IB(args) {
						return f.(*AFn).X_invoke_ArityVariadic(args)
					} else {
						return f.(*AFn).X_invoke_ArityVariadic(append(Into_array.Arity1IA(Take.X_invoke_Arity2(fixed_arity, args).(*CljsCoreLazySeq)), Drop.X_invoke_Arity2(fixed_arity, args).(*CljsCoreLazySeq))...)
					}
				}
			}
		}, func(f interface{}, x interface{}, args interface{}) interface{} {
			{
				var arglist = List_STAR_.X_invoke_Arity2(x, args).(*CljsCoreCons)
				var fixed_arity = MaxFixedArity_(f)
				_, _ = arglist, fixed_arity
				if (X_EQ_.Arity2IIB(float64(-1), fixed_arity)) || (Bounded_count.X_invoke_Arity2(arglist, (fixed_arity+float64(1))).(float64) <= fixed_arity) {
					return Call_(f.(CljsCoreIFn), Into_array.Arity1IA(arglist)...)
				} else {
					return f.(*AFn).X_invoke_ArityVariadic(append(Into_array.Arity1IA(Take.X_invoke_Arity2(fixed_arity, arglist).(*CljsCoreLazySeq)), Drop.X_invoke_Arity2(fixed_arity, arglist).(*CljsCoreLazySeq))...)
				}
			}
		}, func(f interface{}, x interface{}, y interface{}, args interface{}) interface{} {
			{
				var arglist = List_STAR_.X_invoke_Arity3(x, y, args).(*CljsCoreCons)
				var fixed_arity = MaxFixedArity_(f)
				_, _ = arglist, fixed_arity
				if (X_EQ_.Arity2IIB(float64(-1), fixed_arity)) || (Bounded_count.X_invoke_Arity2(arglist, (fixed_arity+float64(1))).(float64) <= fixed_arity) {
					return Call_(f.(CljsCoreIFn), Into_array.Arity1IA(arglist)...)
				} else {
					return f.(*AFn).X_invoke_ArityVariadic(append(Into_array.Arity1IA(Take.X_invoke_Arity2(fixed_arity, arglist).(*CljsCoreLazySeq)), Drop.X_invoke_Arity2(fixed_arity, arglist).(*CljsCoreLazySeq))...)
				}
			}
		}, func(f interface{}, x interface{}, y interface{}, z interface{}, args interface{}) interface{} {
			{
				var arglist = List_STAR_.X_invoke_Arity4(x, y, z, args).(*CljsCoreCons)
				var fixed_arity = MaxFixedArity_(f)
				_, _ = arglist, fixed_arity
				if (X_EQ_.Arity2IIB(float64(-1), fixed_arity)) || (Bounded_count.X_invoke_Arity2(arglist, (fixed_arity+float64(1))).(float64) <= fixed_arity) {
					return Call_(f.(CljsCoreIFn), Into_array.Arity1IA(arglist)...)
				} else {
					return f.(*AFn).X_invoke_ArityVariadic(append(Into_array.Arity1IA(Take.X_invoke_Arity2(fixed_arity, arglist).(*CljsCoreLazySeq)), Drop.X_invoke_Arity2(fixed_arity, arglist).(*CljsCoreLazySeq))...)
				}
			}
		}, func(f_a_b_c_d_args__ ...interface{}) interface{} {
			var f = f_a_b_c_d_args__[0]
			var a = f_a_b_c_d_args__[1]
			var b = f_a_b_c_d_args__[2]
			var c = f_a_b_c_d_args__[3]
			var d = f_a_b_c_d_args__[4]
			var args = Seq.Arity1IQ(f_a_b_c_d_args__[5])
			_, _, _, _, _, _ = f, a, b, c, d, args
			{
				var arglist = Cons.X_invoke_Arity2(a, Cons.X_invoke_Arity2(b, Cons.X_invoke_Arity2(c, Cons.X_invoke_Arity2(d, Spread.X_invoke_Arity1(args)).(*CljsCoreCons)).(*CljsCoreCons)).(*CljsCoreCons)).(*CljsCoreCons)
				var fixed_arity = MaxFixedArity_(f)
				_, _ = arglist, fixed_arity
				if (X_EQ_.Arity2IIB(float64(-1), fixed_arity)) || (Bounded_count.X_invoke_Arity2(arglist, (fixed_arity+float64(1))).(float64) <= fixed_arity) {
					return Call_(f.(CljsCoreIFn), append(Into_array.Arity1IA(Seq_(Butlast.X_invoke_Arity1(arglist))), Into_array.Arity1IA(Last.X_invoke_Arity1(arglist))...))
				} else {
					return f.(*AFn).X_invoke_ArityVariadic(append(Into_array.Arity1IA(Take.X_invoke_Arity2(fixed_arity, arglist).(*CljsCoreLazySeq)), Drop.X_invoke_Arity2(fixed_arity, arglist).(*CljsCoreLazySeq))...)
				}
			}
		})
	}(&AFn{})

	Native_satisfies_QMARK_ = func(native_satisfies_QMARK_ *AFn) *AFn {
		return Fn(native_satisfies_QMARK_, 2, func(p interface{}, x interface{}) bool {
			return DecoratedValue_(x).Type().Implements(p.(reflect.Type))
		})
	}(&AFn{})

}

var X_STAR_clojurescript_version_STAR_ string

var X_STAR_print_length_STAR_ float64

var X_STAR_print_level_STAR_ float64

// Set *print-fn* to f.
var Set_print_fn_BANG_ *AFn

var Symbol_QMARK_ *AFn

// Takes a fn f and returns a fn that takes the same arguments as f,
// has the same effects, if any, and returns the opposite truth value.
var Complement *AFn

// Returns a lazy sequence of the items in coll for which
// (pred item) returns false. pred must be free of side-effects.
// Returns a transducer when no collection is provided.
var Remove *AFn

// @param {...*} var_args
var Identity *AFn

// Returns a random floating point number between 0 (inclusive) and
// n (default 1) (exclusive).
var Rand *AFn

// Equality. Returns true if x equals y, false if not. Compares
// numbers and collections in a type-independent manner.  Clojure's immutable data
// structures define -equiv (and thus =) as a value, not an identity,
// comparison.
// @param {...*} var_args
var X_EQ_ *AFn

// Returns a sorted sequence of the items in coll. Comp can be
// boolean-valued comparison funcion, or a -/0/+ valued comparator.
// Comp defaults to compare.
var Sort *AFn

// Returns the value mapped to key, not-found or nil if key not present.
var Get *AFn

var Quote_string *AFn

// Prefer this to pr-seq, because it makes the printing function
// configurable, allowing efficient implementations such as appending
// to a StringBuffer.
var Pr_writer *AFn

var Pr_sequential_writer *AFn

var Type_ *AFn

var Type__GT_str *AFn

// Returns true if n is an integer.
var Integer_QMARK_ *AFn

// Creates a new javascript array.
// @param {...*} var_args
// @param {...*} var_args
var Array *AFn

var Make_array *AFn

// Coerce to char
var Char *AFn

var String_hash_cache map[interface{}]interface{}

var Add_to_string_hash_cache *AFn

var Hash_string *AFn

// Set *print-fn* to console.log
var Enable_console_print_BANG_ *AFn

// Applies fn f to the argument list formed by prepending intervening arguments to args.
// First cut.  Not lazy.  Needs to use emitted toApply.
// @param {...*} var_args
var Apply *AFn

// Internal - do not use!
var Native_satisfies_QMARK_ *AFn

func (_ *CljsCoreTransientArrayMap) CljsCoreITransientMap__() {}

func (tcoll *CljsCoreTransientArrayMap) X_dissoc_BANG__Arity2(key interface{}) interface{} {
	if Truth_(tcoll.Editable_QMARK_) {
		{
			var idx = Array_map_index_of.X_invoke_Arity2(tcoll, key).(float64)
			_ = idx
			if idx >= float64(0) {
				tcoll.Arr.([]interface{})[int(idx)] = Aget_(tcoll.Arr, (tcoll.Len.(float64) - float64(2)))
				tcoll.Arr.([]interface{})[int((idx + float64(1)))] = Aget_(tcoll.Arr, (tcoll.Len.(float64) - float64(1)))
				js.JSArray_(&tcoll.Arr).Pop()
				js.JSArray_(&tcoll.Arr).Pop()
				tcoll.Len = (tcoll.Len.(float64) - float64(2))

			} else {
			}
			return tcoll
		}
	} else {
		panic((&js.Error{"dissoc! after persistent!"}))
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreISorted__() {}

func (coll *CljsCorePersistentTreeSet) X_sorted_seq_Arity2(ascending_QMARK_ interface{}) interface{} {
	return Seq.Arity1IQ(Map_.X_invoke_Arity2(Key, Decorate_(coll.Tree_map).(CljsCoreISorted).X_sorted_seq_Arity2(ascending_QMARK_)).(*CljsCoreLazySeq))
}

func (coll *CljsCorePersistentTreeSet) X_sorted_seq_from_Arity3(k interface{}, ascending_QMARK_ interface{}) interface{} {
	return Seq.Arity1IQ(Map_.X_invoke_Arity2(Key, Decorate_(coll.Tree_map).(CljsCoreISorted).X_sorted_seq_from_Arity3(k, ascending_QMARK_)).(*CljsCoreLazySeq))
}

func (coll *CljsCorePersistentTreeSet) X_entry_key_Arity2(entry interface{}) interface{} {
	return entry
}

func (coll *CljsCorePersistentTreeSet) X_comparator_Arity1() interface{} {
	return Decorate_(coll.Tree_map).(CljsCoreISorted).X_comparator_Arity1()
}
