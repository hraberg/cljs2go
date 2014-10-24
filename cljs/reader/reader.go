// Compiled by ClojureScript to Go 0.0-2371
// cljs.reader

package reader

import (
	"reflect"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	goog_string "github.com/hraberg/cljs2go/goog/string"
	"github.com/hraberg/cljs2go/js"
)

func init() {
	Read_char = func(read_char *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_char, 1, func(reader interface{}) interface{} {
			return cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
		})
	}(&cljs_core.AFn{})

	Unread = func(unread *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(unread, 2, func(reader interface{}, ch interface{}) interface{} {
			return cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Unread_Arity2(ch)
		})
	}(&cljs_core.AFn{})

	X__GT_StringPushbackReader = func(__GT_StringPushbackReader *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(__GT_StringPushbackReader, 3, func(s interface{}, buffer interface{}, idx interface{}) interface{} {
			return (&CljsReaderStringPushbackReader{s, buffer, idx})
		})
	}(&cljs_core.AFn{})

	Push_back_reader = func(push_back_reader *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(push_back_reader, 1, func(s interface{}) interface{} {
			return (&CljsReaderStringPushbackReader{s, []interface{}{}, float64(-1)})
		})
	}(&cljs_core.AFn{})

	Whitespace_QMARK_ = func(whitespace_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(whitespace_QMARK_, 1, func(ch interface{}) bool {
			{
				var or__182__auto__ = func() interface{} {
					var G__4 = ch
					_ = G__4
					return cljs_core.Truth_(cljs_core.Native_invoke_func.X_invoke_Arity2(goog_string.IsBreakingWhitespace, []interface{}{G__4}))
				}()
				_ = or__182__auto__
				if cljs_core.Truth_(or__182__auto__) {
					return or__182__auto__.(bool)
				} else {
					return reflect.DeepEqual(",", ch)
				}
			}
		})
	}(&cljs_core.AFn{})

	Numeric_QMARK_ = func(numeric_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(numeric_QMARK_, 1, func(ch interface{}) bool {
			{
				var G__6 = ch
				_ = G__6
				return cljs_core.Truth_(cljs_core.Native_invoke_func.X_invoke_Arity2(goog_string.IsNumeric, []interface{}{G__6}))
			}
		})
	}(&cljs_core.AFn{})

	Comment_prefix_QMARK_ = func(comment_prefix_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(comment_prefix_QMARK_, 1, func(ch interface{}) bool {
			return reflect.DeepEqual(";", ch)
		})
	}(&cljs_core.AFn{})

	Number_literal_QMARK_ = func(number_literal_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(number_literal_QMARK_, 2, func(reader interface{}, initch interface{}) bool {
			return (Numeric_QMARK_.Arity1IB(initch)) || (((reflect.DeepEqual("+", initch)) || (reflect.DeepEqual("-", initch))) && (Numeric_QMARK_.Arity1IB(func() interface{} {
				var next_ch = cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
				_ = next_ch
				cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Unread_Arity2(next_ch)
				return next_ch
			}())))
		})
	}(&cljs_core.AFn{})

	Reader_error = func(reader_error *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(reader_error, 1, func(rdr_msg__ ...interface{}) interface{} {
			var rdr = rdr_msg__[0]
			var msg = cljs_core.Seq.Arity1IQ(rdr_msg__[1])
			_, _ = rdr, msg
			panic((&js.Error{cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, msg)}))
		})
	}(&cljs_core.AFn{})

	Macro_terminating_QMARK_ = func(macro_terminating_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(macro_terminating_QMARK_, 1, func(ch interface{}) bool {
			{
				var and__170__auto__ = !(reflect.DeepEqual(ch, "#"))
				_ = and__170__auto__
				if cljs_core.Truth_(and__170__auto__) {
					{
						var and__170__auto_____1 = !(reflect.DeepEqual(ch, "'"))
						_ = and__170__auto_____1
						if cljs_core.Truth_(and__170__auto_____1) {
							{
								var and__170__auto_____2 = !(reflect.DeepEqual(ch, ":"))
								_ = and__170__auto_____2
								if cljs_core.Truth_(and__170__auto_____2) {
									{
										var G__14 = ch
										_ = G__14
										return cljs_core.Truth_(Macros.X_invoke_Arity1(G__14))
									}
								} else {
									return and__170__auto_____2
								}
							}
						} else {
							return and__170__auto_____1
						}
					}
				} else {
					return and__170__auto__
				}
			}
		})
	}(&cljs_core.AFn{})

	Skip_line = func(skip_line *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(skip_line, 2, func(reader interface{}, ___ interface{}) interface{} {
			{
				for {
					{
						var ch = cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
						_ = ch
						if (reflect.DeepEqual(ch, "\n")) || (reflect.DeepEqual(ch, "\r")) || (cljs_core.Nil_(ch)) {
							return reader
						} else {
							continue
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Int_pattern = cljs_core.Re_pattern.X_invoke_Arity1("^([-+]?)(?:(0)|([1-9][0-9]*)|0[xX]([0-9A-Fa-f]+)|0([0-7]+)|([1-9][0-9]?)[rR]([0-9A-Za-z]+))(N)?$")

	Ratio_pattern = cljs_core.Re_pattern.X_invoke_Arity1("^([-+]?[0-9]+)/([0-9]+)$")

	Float_pattern = cljs_core.Re_pattern.X_invoke_Arity1("^([-+]?[0-9]+(\\.[0-9]*)?([eE][-+]?[0-9]+)?)(M)?$")

	Symbol_pattern = cljs_core.Re_pattern.X_invoke_Arity1("^[:]?([^0-9/].*/)?([^0-9/][^/]*)$")

	Re_matches_STAR_ = func(re_matches_STAR_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(re_matches_STAR_, 2, func(re interface{}, s interface{}) interface{} {
			{
				var matches = cljs_core.Native_invoke_instance_method.X_invoke_Arity3(re, "Exec", []interface{}{s})
				_ = matches
				if (!(cljs_core.Nil_(matches))) && (reflect.DeepEqual(cljs_core.Aget_(matches, float64(0)), s)) {
					if cljs_core.Alength_(matches) == float64(1) {
						return cljs_core.Aget_(matches, float64(0))
					} else {
						return matches
					}
				} else {
					return nil
				}
			}
		})
	}(&cljs_core.AFn{})

	Match_int = func(match_int *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(match_int, 1, func(s interface{}) interface{} {
			{
				var groups = Re_matches_STAR_.X_invoke_Arity2(Int_pattern, s)
				var ie8_fix = cljs_core.Aget_(groups, float64(2))
				var zero = func() interface{} {
					if cljs_core.X_EQ_.Arity2IIB(ie8_fix, "") {
						return nil
					} else {
						return ie8_fix
					}
				}()
				_, _, _ = groups, ie8_fix, zero
				if !(cljs_core.Nil_(zero)) {
					return float64(0)
				} else {
					{
						var a = func() []interface{} {
							if cljs_core.Truth_(cljs_core.Aget_(groups, float64(3))) {
								return []interface{}{cljs_core.Aget_(groups, float64(3)), float64(10)}
							} else {
								return func() []interface{} {
									if cljs_core.Truth_(cljs_core.Aget_(groups, float64(4))) {
										return []interface{}{cljs_core.Aget_(groups, float64(4)), float64(16)}
									} else {
										return func() []interface{} {
											if cljs_core.Truth_(cljs_core.Aget_(groups, float64(5))) {
												return []interface{}{cljs_core.Aget_(groups, float64(5)), float64(8)}
											} else {
												return func() []interface{} {
													if cljs_core.Truth_(cljs_core.Aget_(groups, float64(6))) {
														return []interface{}{cljs_core.Aget_(groups, float64(7)), func() interface{} {
															var G__19 = cljs_core.Aget_(groups, float64(6))
															var G__20 = float64(10)
															_, _ = G__19, G__20
															return cljs_core.Native_invoke_func.X_invoke_Arity2(js.ParseInt, []interface{}{G__19, G__20})
														}()}
													} else {
														return []interface{}{nil, nil}
													}
												}()
											}
										}()
									}
								}()
							}
						}()
						var n = (a[int(float64(0))])
						var radix = (a[int(float64(1))])
						_, _, _ = a, n, radix
						if cljs_core.Nil_(n) {
							return nil
						} else {
							{
								var parsed = func() interface{} {
									var G__21 = n
									var G__22 = radix
									_, _ = G__21, G__22
									return cljs_core.Native_invoke_func.X_invoke_Arity2(js.ParseInt, []interface{}{G__21, G__22})
								}()
								_ = parsed
								if reflect.DeepEqual("-", cljs_core.Aget_(groups, float64(1))) {
									return (-parsed.(float64))
								} else {
									return parsed
								}
							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Match_ratio = func(match_ratio *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(match_ratio, 1, func(s interface{}) interface{} {
			{
				var groups = Re_matches_STAR_.X_invoke_Arity2(Ratio_pattern, s)
				var numinator = cljs_core.Aget_(groups, float64(1))
				var denominator = cljs_core.Aget_(groups, float64(2))
				_, _, _ = groups, numinator, denominator
				return (func() interface{} {
					var G__27 = numinator
					var G__28 = float64(10)
					_, _ = G__27, G__28
					return cljs_core.Native_invoke_func.X_invoke_Arity2(js.ParseInt, []interface{}{G__27, G__28})
				}().(float64) / func() interface{} {
					var G__29 = denominator
					var G__30 = float64(10)
					_, _ = G__29, G__30
					return cljs_core.Native_invoke_func.X_invoke_Arity2(js.ParseInt, []interface{}{G__29, G__30})
				}().(float64))
			}
		})
	}(&cljs_core.AFn{})

	Match_float = func(match_float *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(match_float, 1, func(s interface{}) interface{} {
			{
				var G__32 = s
				_ = G__32
				return cljs_core.Native_invoke_func.X_invoke_Arity2(js.ParseFloat, []interface{}{G__32})
			}
		})
	}(&cljs_core.AFn{})

	Match_number = func(match_number *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(match_number, 1, func(s interface{}) interface{} {
			if cljs_core.Truth_(Re_matches_STAR_.X_invoke_Arity2(Int_pattern, s)) {
				return Match_int.X_invoke_Arity1(s)
			} else {
				if cljs_core.Truth_(Re_matches_STAR_.X_invoke_Arity2(Ratio_pattern, s)) {
					return Match_ratio.X_invoke_Arity1(s).(float64)
				} else {
					if cljs_core.Truth_(Re_matches_STAR_.X_invoke_Arity2(Float_pattern, s)) {
						return Match_float.X_invoke_Arity1(s)
					} else {
						return nil
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Escape_char_map = func(escape_char_map *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(escape_char_map, 1, func(c interface{}) interface{} {
			if reflect.DeepEqual(c, "t") {
				return "\t"
			} else {
				if reflect.DeepEqual(c, "r") {
					return "\r"
				} else {
					if reflect.DeepEqual(c, "n") {
						return "\n"
					} else {
						if reflect.DeepEqual(c, "\\") {
							return "\\"
						} else {
							if reflect.DeepEqual(c, "\"") {
								return "\""
							} else {
								if reflect.DeepEqual(c, "b") {
									return "\b"
								} else {
									if reflect.DeepEqual(c, "f") {
										return "\f"
									} else {
										return nil

									}
								}
							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Unicode_2_pattern = cljs_core.Re_pattern.X_invoke_Arity1("^[0-9A-Fa-f]{2}$")

	Unicode_4_pattern = cljs_core.Re_pattern.X_invoke_Arity1("^[0-9A-Fa-f]{4}$")

	Validate_unicode_escape = func(validate_unicode_escape *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(validate_unicode_escape, 4, func(unicode_pattern interface{}, reader interface{}, escape_char interface{}, unicode_str interface{}) interface{} {
			if cljs_core.Truth_(cljs_core.Re_matches.X_invoke_Arity2(unicode_pattern, unicode_str)) {
				return unicode_str
			} else {
				return Reader_error.X_invoke_ArityVariadic(reader, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Unexpected unicode escape \\", escape_char, unicode_str}))
			}
		})
	}(&cljs_core.AFn{})

	Make_unicode_char = func(make_unicode_char *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(make_unicode_char, 1, func(code_str interface{}) interface{} {
			{
				var code = func() interface{} {
					var G__35 = code_str
					var G__36 = float64(16)
					_, _ = G__35, G__36
					return cljs_core.Native_invoke_func.X_invoke_Arity2(js.ParseInt, []interface{}{G__35, G__36})
				}()
				_ = code
				return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.String, "FromCharCode", []interface{}{code})
			}
		})
	}(&cljs_core.AFn{})

	Escape_char = func(escape_char *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(escape_char, 2, func(buffer interface{}, reader interface{}) interface{} {
			{
				var ch = cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
				var mapresult = Escape_char_map.X_invoke_Arity1(ch)
				_, _ = ch, mapresult
				if cljs_core.Truth_(mapresult) {
					return mapresult
				} else {
					if reflect.DeepEqual(ch, "x") {
						return Make_unicode_char.X_invoke_Arity1(Validate_unicode_escape.X_invoke_Arity4(Unicode_2_pattern, reader, ch, Read_2_chars.X_invoke_Arity1(reader)))
					} else {
						if reflect.DeepEqual(ch, "u") {
							return Make_unicode_char.X_invoke_Arity1(Validate_unicode_escape.X_invoke_Arity4(Unicode_4_pattern, reader, ch, Read_4_chars.X_invoke_Arity1(reader)))
						} else {
							if Numeric_QMARK_.Arity1IB(ch) {
								return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.String, "FromCharCode", []interface{}{ch})
							} else {
								return Reader_error.X_invoke_ArityVariadic(reader, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Unexpected unicode escape \\", ch}))

							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Read_past = func(read_past *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_past, 2, func(pred interface{}, rdr interface{}) interface{} {
			{
				var ch interface{} = cljs_core.Decorate_(rdr).(CljsReaderPushbackReader).Read_char_Arity1()
				_ = ch
				for {
					if cljs_core.Truth_(func() interface{} {
						var G__38 = ch
						_ = G__38
						return pred.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__38)
					}()) {
						ch = cljs_core.Decorate_(rdr).(CljsReaderPushbackReader).Read_char_Arity1()
						continue
					} else {
						return ch
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Read_delimited_list = func(read_delimited_list *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_delimited_list, 3, func(delim interface{}, rdr interface{}, recursive_QMARK_ interface{}) interface{} {
			{
				var a interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)
				_ = a
				for {
					{
						var ch = Read_past.X_invoke_Arity2(Whitespace_QMARK_, rdr)
						_ = ch
						if cljs_core.Truth_(ch) {
						} else {
							Reader_error.X_invoke_ArityVariadic(rdr, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"EOF while reading"}))
						}
						if reflect.DeepEqual(delim, ch) {
							return cljs_core.Persistent_BANG_.X_invoke_Arity1(a)
						} else {
							{
								var temp__4220__auto__ = func() interface{} {
									var G__46 = ch
									_ = G__46
									return Macros.X_invoke_Arity1(G__46)
								}()
								_ = temp__4220__auto__
								if cljs_core.Truth_(temp__4220__auto__) {
									{
										var macrofn = temp__4220__auto__
										_ = macrofn
										{
											var mret = func() interface{} {
												var G__47 = rdr
												var G__48 = ch
												_, _ = G__47, G__48
												return macrofn.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__47, G__48)
											}()
											_ = mret
											a = func() interface{} {
												if reflect.DeepEqual(mret, rdr) {
													return a
												} else {
													return cljs_core.Conj_BANG_.X_invoke_Arity2(a, mret)
												}
											}()
											continue
										}
									}
								} else {
									cljs_core.Decorate_(rdr).(CljsReaderPushbackReader).Unread_Arity2(ch)
									{
										var o = func() interface{} {
											var G__49 = rdr
											var G__50 = true
											var G__51 interface{} = nil
											var G__52 = recursive_QMARK_
											_, _, _, _ = G__49, G__50, G__51, G__52
											return Read.X_invoke_Arity4(G__49, G__50, G__51, G__52)
										}()
										_ = o
										a = func() interface{} {
											if reflect.DeepEqual(o, rdr) {
												return a
											} else {
												return cljs_core.Conj_BANG_.X_invoke_Arity2(a, o)
											}
										}()
										continue
									}
								}
							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Not_implemented = func(not_implemented *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(not_implemented, 2, func(rdr interface{}, ch interface{}) interface{} {
			return Reader_error.X_invoke_ArityVariadic(rdr, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Reader for ", ch, " not implemented yet"}))
		})
	}(&cljs_core.AFn{})

	Read_dispatch = func(read_dispatch *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_dispatch, 2, func(rdr interface{}, ___ interface{}) interface{} {
			{
				var ch = cljs_core.Decorate_(rdr).(CljsReaderPushbackReader).Read_char_Arity1()
				var dm = func() interface{} {
					var G__58 = ch
					_ = G__58
					return Dispatch_macros.X_invoke_Arity1(G__58)
				}()
				_, _ = ch, dm
				if cljs_core.Truth_(dm) {
					{
						var G__59 = rdr
						var G__60 = ___
						_, _ = G__59, G__60
						return dm.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__59, G__60)
					}
				} else {
					{
						var temp__4220__auto__ = func() interface{} {
							var G__61 = rdr
							var G__62 = ch
							_, _ = G__61, G__62
							return Maybe_read_tagged_type.X_invoke_Arity2(G__61, G__62)
						}()
						_ = temp__4220__auto__
						if cljs_core.Truth_(temp__4220__auto__) {
							{
								var obj = temp__4220__auto__
								_ = obj
								return obj
							}
						} else {
							return Reader_error.X_invoke_ArityVariadic(rdr, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"No dispatch macro for ", ch}))
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Read_unmatched_delimiter = func(read_unmatched_delimiter *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_unmatched_delimiter, 2, func(rdr interface{}, ch interface{}) interface{} {
			return Reader_error.X_invoke_ArityVariadic(rdr, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Unmached delimiter ", ch}))
		})
	}(&cljs_core.AFn{})

	Read_list = func(read_list *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_list, 2, func(rdr interface{}, ___ interface{}) interface{} {
			return cljs_core.Apply.X_invoke_Arity2(cljs_core.List, Read_delimited_list.X_invoke_Arity3(")", rdr, true))
		})
	}(&cljs_core.AFn{})

	Read_comment = Skip_line

	Read_vector = func(read_vector *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_vector, 2, func(rdr interface{}, ___ interface{}) interface{} {
			return Read_delimited_list.X_invoke_Arity3("]", rdr, true)
		})
	}(&cljs_core.AFn{})

	Read_map = func(read_map *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_map, 2, func(rdr interface{}, ___ interface{}) interface{} {
			{
				var l = Read_delimited_list.X_invoke_Arity3("}", rdr, true)
				_ = l
				if cljs_core.Odd_QMARK_.Arity1IB(cljs_core.Count.X_invoke_Arity1(l).(float64)) {
					Reader_error.X_invoke_ArityVariadic(rdr, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Map literal must contain an even number of forms"}))
				} else {
				}
				return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, l)
			}
		})
	}(&cljs_core.AFn{})

	Read_number = func(read_number *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_number, 2, func(reader interface{}, initch interface{}) interface{} {
			{
				var buffer = (&goog_string.StringBuffer{initch})
				var ch interface{} = cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
				_, _ = buffer, ch
				for {
					if cljs_core.Truth_(func() interface{} {
						var or__182__auto__ = cljs_core.Nil_(ch)
						_ = or__182__auto__
						if cljs_core.Truth_(or__182__auto__) {
							return or__182__auto__
						} else {
							{
								var or__182__auto_____1 = Whitespace_QMARK_.Arity1IB(ch)
								_ = or__182__auto_____1
								if cljs_core.Truth_(or__182__auto_____1) {
									return or__182__auto_____1
								} else {
									{
										var G__68 = ch
										_ = G__68
										return Macros.X_invoke_Arity1(G__68)
									}
								}
							}
						}
					}()) {
						cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Unread_Arity2(ch)
						{
							var s = cljs_core.Native_invoke_instance_method.X_invoke_Arity3(buffer, "ToString", []interface{}{})
							_ = s
							{
								var or__182__auto__ = Match_number.X_invoke_Arity1(s)
								_ = or__182__auto__
								if cljs_core.Truth_(or__182__auto__) {
									return or__182__auto__
								} else {
									return Reader_error.X_invoke_ArityVariadic(reader, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Invalid number format [", s, "]"}))
								}
							}
						}
					} else {
						buffer, ch = func() *goog_string.StringBuffer {
							cljs_core.Native_invoke_instance_method.X_invoke_Arity3(buffer, "Append", []interface{}{ch})
							return buffer
						}(), cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
						continue
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Read_string_STAR_ = func(read_string_STAR_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_string_STAR_, 2, func(reader interface{}, ___ interface{}) interface{} {
			{
				var buffer = (&goog_string.StringBuffer{})
				var ch interface{} = cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
				_, _ = buffer, ch
				for {
					if cljs_core.Nil_(ch) {
						return Reader_error.X_invoke_ArityVariadic(reader, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"EOF while reading"}))
					} else {
						if reflect.DeepEqual("\\", ch) {
							buffer, ch = func() *goog_string.StringBuffer {
								cljs_core.Native_invoke_instance_method.X_invoke_Arity3(buffer, "Append", []interface{}{Escape_char.X_invoke_Arity2(buffer, reader)})
								return buffer
							}(), cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
							continue
						} else {
							if reflect.DeepEqual("\"", ch) {
								return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(buffer, "ToString", []interface{}{})
							} else {
								buffer, ch = func() *goog_string.StringBuffer {
									cljs_core.Native_invoke_instance_method.X_invoke_Arity3(buffer, "Append", []interface{}{ch})
									return buffer
								}(), cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
								continue

							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Read_raw_string_STAR_ = func(read_raw_string_STAR_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_raw_string_STAR_, 2, func(reader interface{}, ___ interface{}) interface{} {
			{
				var buffer = (&goog_string.StringBuffer{})
				var ch interface{} = cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
				_, _ = buffer, ch
				for {
					if cljs_core.Nil_(ch) {
						return Reader_error.X_invoke_ArityVariadic(reader, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"EOF while reading"}))
					} else {
						if reflect.DeepEqual("\\", ch) {
							cljs_core.Native_invoke_instance_method.X_invoke_Arity3(buffer, "Append", []interface{}{ch})
							{
								var nch = cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
								_ = nch
								if cljs_core.Nil_(nch) {
									return Reader_error.X_invoke_ArityVariadic(reader, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"EOF while reading"}))
								} else {
									buffer, ch = func() *goog_string.StringBuffer {
										var G__71 = buffer
										_ = G__71
										G__71.Append(nch)
										return G__71
									}(), cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
									continue
								}
							}
						} else {
							if reflect.DeepEqual("\"", ch) {
								return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(buffer, "ToString", []interface{}{})
							} else {
								buffer, ch = func() *goog_string.StringBuffer {
									var G__72 = buffer
									_ = G__72
									G__72.Append(ch)
									return G__72
								}(), cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
								continue

							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Special_symbols = func(special_symbols *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(special_symbols, 2, func(t interface{}, not_found interface{}) interface{} {
			if reflect.DeepEqual(t, "nil") {
				return nil
			} else {
				if reflect.DeepEqual(t, "true") {
					return true
				} else {
					if reflect.DeepEqual(t, "false") {
						return false
					} else {
						return not_found

					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Read_symbol = func(read_symbol *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_symbol, 2, func(reader interface{}, initch interface{}) interface{} {
			{
				var token = Read_token.X_invoke_Arity2(reader, initch)
				_ = token
				if cljs_core.Truth_(func() interface{} {
					var G__75 = token
					var G__76 = "/"
					_, _ = G__75, G__76
					return cljs_core.Native_invoke_func.X_invoke_Arity2(goog_string.Contains, []interface{}{G__75, G__76})
				}()) {
					return cljs_core.Symbol.X_invoke_Arity2(cljs_core.Subs.X_invoke_Arity3(token, float64(0), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(token, "IndexOf", []interface{}{"/"})), cljs_core.Subs.X_invoke_Arity3(token, (cljs_core.Native_invoke_instance_method.X_invoke_Arity3(token, "IndexOf", []interface{}{"/"}).(float64)+float64(1)), cljs_core.Native_get_instance_field.X_invoke_Arity2(token, "Length"))).(*cljs_core.CljsCoreSymbol)
				} else {
					return Special_symbols.X_invoke_Arity2(token, cljs_core.Symbol.X_invoke_Arity1(token))
				}
			}
		})
	}(&cljs_core.AFn{})

	Read_keyword = func(read_keyword *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_keyword, 2, func(reader interface{}, initch interface{}) interface{} {
			{
				var token = Read_token.X_invoke_Arity2(reader, cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1())
				var a = Re_matches_STAR_.X_invoke_Arity2(Symbol_pattern, token)
				var token___1 = cljs_core.Aget_(a, float64(0))
				var ns = cljs_core.Aget_(a, float64(1))
				var name = cljs_core.Aget_(a, float64(2))
				_, _, _, _, _ = token, a, token___1, ns, name
				if ((!(cljs_core.Nil_(ns))) && (reflect.DeepEqual(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(ns, "Substring", []interface{}{(cljs_core.Native_get_instance_field.X_invoke_Arity2(ns, "Length").(float64) - float64(2)), cljs_core.Native_get_instance_field.X_invoke_Arity2(ns, "Length")}), ":/"))) || (reflect.DeepEqual(cljs_core.Aget_(name, (cljs_core.Native_get_instance_field.X_invoke_Arity2(name, "Length").(float64)-float64(1))), ":")) || (!(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(token___1, "IndexOf", []interface{}{"::", float64(1)}).(float64) == float64(-1))) {
					return Reader_error.X_invoke_ArityVariadic(reader, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Invalid token: ", token___1}))
				} else {
					if (!(cljs_core.Nil_(ns))) && (cljs_core.Native_get_instance_field.X_invoke_Arity2(ns, "Length").(float64) > float64(0)) {
						return cljs_core.Keyword.X_invoke_Arity2(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(ns, "Substring", []interface{}{float64(0), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(ns, "IndexOf", []interface{}{"/"})}), name).(*cljs_core.CljsCoreKeyword)
					} else {
						return cljs_core.Keyword.X_invoke_Arity1(token___1)
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Desugar_meta = func(desugar_meta *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(desugar_meta, 1, func(f interface{}) interface{} {
			if cljs_core.Value_(f).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreSymbol)(nil)).Elem()) {
				return (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "tag", Fqn: "tag", X_hash: float64(-1290361223)}), f}, nil})
			} else {
				if cljs_core.Value_(f).Kind() == reflect.String {
					return (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "tag", Fqn: "tag", X_hash: float64(-1290361223)}), f}, nil})
				} else {
					if cljs_core.Value_(f).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
						return cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{f, true}, true, false).(*cljs_core.CljsCorePersistentArrayMap)
					} else {
						return f

					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Wrapping_reader = func(wrapping_reader *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(wrapping_reader, 1, func(sym interface{}) interface{} {
			return func(G__85 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__85, 2, func(rdr interface{}, ___ interface{}) interface{} {
					return cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(func() interface{} {
						var G__81 = rdr
						var G__82 = true
						var G__83 interface{} = nil
						var G__84 = true
						_, _, _, _ = G__81, G__82, G__83, G__84
						return Read.X_invoke_Arity4(G__81, G__82, G__83, G__84)
					}())).(cljs_core.CljsCoreICollection).X_conj_Arity2(sym)
				})
			}(&cljs_core.AFn{})
		})
	}(&cljs_core.AFn{})

	Throwing_reader = func(throwing_reader *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(throwing_reader, 1, func(msg interface{}) interface{} {
			return func(G__86 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__86, 2, func(rdr interface{}, ___ interface{}) interface{} {
					return Reader_error.X_invoke_ArityVariadic(rdr, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{msg}))
				})
			}(&cljs_core.AFn{})
		})
	}(&cljs_core.AFn{})

	Read_meta = func(read_meta *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_meta, 2, func(rdr interface{}, ___ interface{}) interface{} {
			{
				var m = Desugar_meta.X_invoke_Arity1(func() interface{} {
					var G__95 = rdr
					var G__96 = true
					var G__97 interface{} = nil
					var G__98 = true
					_, _, _, _ = G__95, G__96, G__97, G__98
					return Read.X_invoke_Arity4(G__95, G__96, G__97, G__98)
				}())
				_ = m
				if cljs_core.Map_QMARK_.Arity1IB(m) {
				} else {
					Reader_error.X_invoke_ArityVariadic(rdr, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Metadata must be Symbol,Keyword,String or Map"}))
				}
				{
					var o = func() interface{} {
						var G__99 = rdr
						var G__100 = true
						var G__101 interface{} = nil
						var G__102 = true
						_, _, _, _ = G__99, G__100, G__101, G__102
						return Read.X_invoke_Arity4(G__99, G__100, G__101, G__102)
					}()
					_ = o
					if cljs_core.DecoratedValue_(o).Type().Implements(reflect.TypeOf((*cljs_core.CljsCoreIWithMeta)(nil)).Elem()) {
						return cljs_core.With_meta.X_invoke_Arity2(o, cljs_core.Merge.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Meta.X_invoke_Arity1(o), m})))
					} else {
						return Reader_error.X_invoke_ArityVariadic(rdr, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Metadata can only be applied to IWithMetas"}))
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Read_set = func(read_set *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_set, 2, func(rdr interface{}, ___ interface{}) interface{} {
			return cljs_core.Set.X_invoke_Arity1(Read_delimited_list.X_invoke_Arity3("}", rdr, true))
		})
	}(&cljs_core.AFn{})

	Read_regex = func(read_regex *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_regex, 2, func(rdr interface{}, ch interface{}) interface{} {
			return cljs_core.Re_pattern.X_invoke_Arity1(Read_raw_string_STAR_.X_invoke_Arity2(rdr, ch))
		})
	}(&cljs_core.AFn{})

	Read_discard = func(read_discard *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_discard, 2, func(rdr interface{}, ___ interface{}) interface{} {
			{
				var G__107_111 = rdr
				var G__108_112 = true
				var G__109_113 interface{} = nil
				var G__110_114 = true
				_, _, _, _ = G__107_111, G__108_112, G__109_113, G__110_114
				Read.X_invoke_Arity4(G__107_111, G__108_112, G__109_113, G__110_114)
			}
			return rdr
		})
	}(&cljs_core.AFn{})

	Dispatch_macros = func(dispatch_macros *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(dispatch_macros, 1, func(s interface{}) interface{} {
			if reflect.DeepEqual(s, "{") {
				return Read_set
			} else {
				if reflect.DeepEqual(s, "<") {
					return Throwing_reader.X_invoke_Arity1("Unreadable form").(cljs_core.CljsCoreIFn)
				} else {
					if reflect.DeepEqual(s, "\"") {
						return Read_regex
					} else {
						if reflect.DeepEqual(s, "!") {
							return Read_comment
						} else {
							if reflect.DeepEqual(s, "_") {
								return Read_discard
							} else {
								return nil

							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Read = func(read *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read, 4, func(reader interface{}, eof_is_error interface{}, sentinel interface{}, is_recursive interface{}) interface{} {
			for {
				{
					var ch = cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()
					_ = ch
					if cljs_core.Nil_(ch) {
						if cljs_core.Truth_(eof_is_error) {
							return Reader_error.X_invoke_ArityVariadic(reader, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"EOF while reading"}))
						} else {
							return sentinel
						}
					} else {
						if Whitespace_QMARK_.Arity1IB(ch) {
							reader, eof_is_error, sentinel, is_recursive = reader, eof_is_error, sentinel, is_recursive
							continue
						} else {
							if Comment_prefix_QMARK_.Arity1IB(ch) {
								reader, eof_is_error, sentinel, is_recursive = func() interface{} {
									var G__119 = reader
									var G__120 = ch
									_, _ = G__119, G__120
									return Read_comment.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__119, G__120)
								}(), eof_is_error, sentinel, is_recursive
								continue
							} else {
								{
									var f = Macros.X_invoke_Arity1(ch)
									var res = func() interface{} {
										if cljs_core.Truth_(f) {
											return func() interface{} {
												var G__121 = reader
												var G__122 = ch
												_, _ = G__121, G__122
												return f.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__121, G__122)
											}()
										} else {
											return func() interface{} {
												if Number_literal_QMARK_.Arity2IIB(reader, ch) {
													return Read_number.X_invoke_Arity2(reader, ch)
												} else {
													return Read_symbol.X_invoke_Arity2(reader, ch)
												}
											}()
										}
									}()
									_, _ = f, res
									if reflect.DeepEqual(res, reader) {
										reader, eof_is_error, sentinel, is_recursive = reader, eof_is_error, sentinel, is_recursive
										continue
									} else {
										return res
									}
								}

							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Read_string = func(read_string *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_string, 1, func(s interface{}) interface{} {
			{
				var r = Push_back_reader.X_invoke_Arity1(s).(*CljsReaderStringPushbackReader)
				_ = r
				return Read.X_invoke_Arity4(r, false, nil, false)
			}
		})
	}(&cljs_core.AFn{})

	Zero_fill_right_and_truncate = func(zero_fill_right_and_truncate *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(zero_fill_right_and_truncate, 2, func(s interface{}, width interface{}) interface{} {
			if cljs_core.X_EQ_.Arity2IIB(width, cljs_core.Count.X_invoke_Arity1(s).(float64)) {
				return s
			} else {
				if width.(float64) < cljs_core.Count.X_invoke_Arity1(s).(float64) {
					return cljs_core.Subs.X_invoke_Arity3(s, float64(0), width)
				} else {
					{
						var b interface{} = (&goog_string.StringBuffer{s})
						_ = b
						for {
							if cljs_core.Native_invoke_instance_method.X_invoke_Arity3(b, "GetLength", []interface{}{}).(float64) < width.(float64) {
								b = cljs_core.Native_invoke_instance_method.X_invoke_Arity3(b, "Append", []interface{}{"0"})
								continue
							} else {
								return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(b, "ToString", []interface{}{})
							}
						}
					}

				}
			}
		})
	}(&cljs_core.AFn{})

	Divisible_QMARK_ = func(divisible_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(divisible_QMARK_, 2, func(num interface{}, div interface{}) interface{} {
			return (cljs_core.Mod.X_invoke_Arity2(num, div).(float64) == float64(0))
		})
	}(&cljs_core.AFn{})

	Indivisible_QMARK_ = func(indivisible_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(indivisible_QMARK_, 2, func(num interface{}, div interface{}) interface{} {
			return !(cljs_core.Truth_(Divisible_QMARK_.X_invoke_Arity2(num, div)))
		})
	}(&cljs_core.AFn{})

	Leap_year_QMARK_ = func(leap_year_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(leap_year_QMARK_, 1, func(year interface{}) interface{} {
			return (cljs_core.Truth_(Divisible_QMARK_.X_invoke_Arity2(year, float64(4)))) && ((cljs_core.Truth_(Indivisible_QMARK_.X_invoke_Arity2(year, float64(100)))) || (cljs_core.Truth_(Divisible_QMARK_.X_invoke_Arity2(year, float64(400)))))
		})
	}(&cljs_core.AFn{})

	Timestamp_regex = (&js.RegExp{Pattern: `(\d\d\d\d)(?:-(\d\d)(?:-(\d\d)(?:[T](\d\d)(?::(\d\d)(?::(\d\d)(?:[.](\d+))?)?)?)?)?)?(?:[Z]|([-+])(\d\d):(\d\d))?`, Flags: ``})

	Parse_int = func(parse_int *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(parse_int, 1, func(s interface{}) interface{} {
			{
				var n = func() interface{} {
					var G__126 = s
					var G__127 = float64(10)
					_, _ = G__126, G__127
					return cljs_core.Native_invoke_func.X_invoke_Arity2(js.ParseInt, []interface{}{G__126, G__127})
				}()
				_ = n
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__128 = n
					_ = G__128
					return cljs_core.Native_invoke_func.X_invoke_Arity2(js.IsNaN, []interface{}{G__128})
				}()) {
					return n
				} else {
					return nil
				}
			}
		})
	}(&cljs_core.AFn{})

	Check = func(check *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(check, 4, func(low interface{}, n interface{}, high interface{}, msg interface{}) interface{} {
			if (low.(float64) <= n.(float64)) && (n.(float64) <= high.(float64)) {
			} else {
				Reader_error.X_invoke_ArityVariadic(nil, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(`` + cljs_core.Str.X_invoke_Arity1(msg).(string) + " Failed:  " + cljs_core.Str.X_invoke_Arity1(low).(string) + "<=" + cljs_core.Str.X_invoke_Arity1(n).(string) + "<=" + cljs_core.Str.X_invoke_Arity1(high).(string))}))
			}
			return n
		})
	}(&cljs_core.AFn{})

	Parse_and_validate_timestamp = func(parse_and_validate_timestamp *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(parse_and_validate_timestamp, 1, func(s interface{}) interface{} {
			{
				var vec__132 = cljs_core.Re_matches.X_invoke_Arity2(Timestamp_regex, s)
				var ___ = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(0), nil)
				var years = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(1), nil)
				var months = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(2), nil)
				var days = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(3), nil)
				var hours = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(4), nil)
				var minutes = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(5), nil)
				var seconds = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(6), nil)
				var fraction = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(7), nil)
				var offset_sign = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(8), nil)
				var offset_hours = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(9), nil)
				var offset_minutes = cljs_core.Nth.X_invoke_Arity3(vec__132, float64(10), nil)
				var v = vec__132
				_, _, _, _, _, _, _, _, _, _, _, _, _ = vec__132, ___, years, months, days, hours, minutes, seconds, fraction, offset_sign, offset_hours, offset_minutes, v
				if cljs_core.Not.Arity1IB(v) {
					return Reader_error.X_invoke_ArityVariadic(nil, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{("Unrecognized date/time syntax: " + cljs_core.Str.X_invoke_Arity1(s).(string))}))
				} else {
					{
						var years___1 = Parse_int.X_invoke_Arity1(years)
						var months___1 = func() interface{} {
							var or__182__auto__ = Parse_int.X_invoke_Arity1(months)
							_ = or__182__auto__
							if cljs_core.Truth_(or__182__auto__) {
								return or__182__auto__
							} else {
								return float64(1)
							}
						}()
						var days___1 = func() interface{} {
							var or__182__auto__ = Parse_int.X_invoke_Arity1(days)
							_ = or__182__auto__
							if cljs_core.Truth_(or__182__auto__) {
								return or__182__auto__
							} else {
								return float64(1)
							}
						}()
						var hours___1 = func() interface{} {
							var or__182__auto__ = Parse_int.X_invoke_Arity1(hours)
							_ = or__182__auto__
							if cljs_core.Truth_(or__182__auto__) {
								return or__182__auto__
							} else {
								return float64(0)
							}
						}()
						var minutes___1 = func() interface{} {
							var or__182__auto__ = Parse_int.X_invoke_Arity1(minutes)
							_ = or__182__auto__
							if cljs_core.Truth_(or__182__auto__) {
								return or__182__auto__
							} else {
								return float64(0)
							}
						}()
						var seconds___1 = func() interface{} {
							var or__182__auto__ = Parse_int.X_invoke_Arity1(seconds)
							_ = or__182__auto__
							if cljs_core.Truth_(or__182__auto__) {
								return or__182__auto__
							} else {
								return float64(0)
							}
						}()
						var fraction___1 = func() interface{} {
							var or__182__auto__ = Parse_int.X_invoke_Arity1(Zero_fill_right_and_truncate.X_invoke_Arity2(fraction, float64(3)))
							_ = or__182__auto__
							if cljs_core.Truth_(or__182__auto__) {
								return or__182__auto__
							} else {
								return float64(0)
							}
						}()
						var offset_sign___1 = func() float64 {
							if cljs_core.X_EQ_.Arity2IIB(offset_sign, "-") {
								return float64(-1)
							} else {
								return float64(1)
							}
						}()
						var offset_hours___1 = func() interface{} {
							var or__182__auto__ = Parse_int.X_invoke_Arity1(offset_hours)
							_ = or__182__auto__
							if cljs_core.Truth_(or__182__auto__) {
								return or__182__auto__
							} else {
								return float64(0)
							}
						}()
						var offset_minutes___1 = func() interface{} {
							var or__182__auto__ = Parse_int.X_invoke_Arity1(offset_minutes)
							_ = or__182__auto__
							if cljs_core.Truth_(or__182__auto__) {
								return or__182__auto__
							} else {
								return float64(0)
							}
						}()
						var offset = (offset_sign___1 * ((offset_hours___1.(float64) * float64(60)) + offset_minutes___1.(float64)))
						_, _, _, _, _, _, _, _, _, _, _ = years___1, months___1, days___1, hours___1, minutes___1, seconds___1, fraction___1, offset_sign___1, offset_hours___1, offset_minutes___1, offset
						return (&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{years___1, Check.X_invoke_Arity4(float64(1), months___1, float64(12), "timestamp month field must be in range 1..12"), Check.X_invoke_Arity4(float64(1), days___1, func() interface{} {
							var G__133 = months___1
							var G__134 = cljs_core.Truth_(Leap_year_QMARK_.X_invoke_Arity1(years___1))
							_, _ = G__133, G__134
							return Days_in_month.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__133, G__134)
						}(), "timestamp day field must be in range 1..last day in month"), Check.X_invoke_Arity4(float64(0), hours___1, float64(23), "timestamp hour field must be in range 0..23"), Check.X_invoke_Arity4(float64(0), minutes___1, float64(59), "timestamp minute field must be in range 0..59"), Check.X_invoke_Arity4(float64(0), seconds___1, func() float64 {
							if cljs_core.X_EQ_.Arity2IIB(minutes___1, float64(59)) {
								return float64(60)
							} else {
								return float64(59)
							}
						}(), "timestamp second field must be in range 0..60"), Check.X_invoke_Arity4(float64(0), fraction___1, float64(999), "timestamp millisecond field must be in range 0..999"), offset}, nil})
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	X_STAR_default_data_reader_fn_STAR_ = cljs_core.Atom.X_invoke_Arity1(nil).(*cljs_core.CljsCoreAtom)

	Maybe_read_tagged_type = func(maybe_read_tagged_type *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(maybe_read_tagged_type, 2, func(rdr interface{}, initch interface{}) interface{} {
			{
				var tag = Read_symbol.X_invoke_Arity2(rdr, initch)
				var pfn = cljs_core.Get.X_invoke_Arity2(cljs_core.Deref.X_invoke_Arity1(X_STAR_tag_table_STAR_), (`` + cljs_core.Str.X_invoke_Arity1(tag).(string)))
				var dfn = cljs_core.Deref.X_invoke_Arity1(X_STAR_default_data_reader_fn_STAR_)
				_, _, _ = tag, pfn, dfn
				if cljs_core.Truth_(pfn) {
					{
						var G__164 = Read.X_invoke_Arity4(rdr, true, nil, false)
						_ = G__164
						return pfn.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__164)
					}
				} else {
					if cljs_core.Truth_(dfn) {
						{
							var G__165 = tag
							var G__166 = Read.X_invoke_Arity4(rdr, true, nil, false)
							_, _ = G__165, G__166
							return dfn.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__165, G__166)
						}
					} else {
						return Reader_error.X_invoke_ArityVariadic(rdr, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Could not find tag parser for ", (`` + cljs_core.Str.X_invoke_Arity1(tag).(string)), " in ", cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Keys.X_invoke_Arity1(cljs_core.Deref.X_invoke_Arity1(X_STAR_tag_table_STAR_))})).(string)}))

					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Register_tag_parser_BANG_ = func(register_tag_parser_BANG_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(register_tag_parser_BANG_, 2, func(tag interface{}, f interface{}) interface{} {
			{
				var tag___1 = (`` + cljs_core.Str.X_invoke_Arity1(tag).(string))
				var old_parser = cljs_core.Get.X_invoke_Arity2(cljs_core.Deref.X_invoke_Arity1(X_STAR_tag_table_STAR_), tag___1)
				_, _ = tag___1, old_parser
				cljs_core.Swap_BANG_.X_invoke_Arity4(X_STAR_tag_table_STAR_, cljs_core.Assoc, tag___1, f)
				return old_parser
			}
		})
	}(&cljs_core.AFn{})

	Deregister_tag_parser_BANG_ = func(deregister_tag_parser_BANG_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(deregister_tag_parser_BANG_, 1, func(tag interface{}) interface{} {
			{
				var tag___1 = (`` + cljs_core.Str.X_invoke_Arity1(tag).(string))
				var old_parser = cljs_core.Get.X_invoke_Arity2(cljs_core.Deref.X_invoke_Arity1(X_STAR_tag_table_STAR_), tag___1)
				_, _ = tag___1, old_parser
				cljs_core.Swap_BANG_.X_invoke_Arity3(X_STAR_tag_table_STAR_, cljs_core.Dissoc, tag___1)
				return old_parser
			}
		})
	}(&cljs_core.AFn{})

	Register_default_tag_parser_BANG_ = func(register_default_tag_parser_BANG_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(register_default_tag_parser_BANG_, 1, func(f interface{}) interface{} {
			{
				var old_parser = cljs_core.Deref.X_invoke_Arity1(X_STAR_default_data_reader_fn_STAR_)
				_ = old_parser
				cljs_core.Swap_BANG_.X_invoke_Arity2(X_STAR_default_data_reader_fn_STAR_, func(G__167 *cljs_core.AFn, old_parser interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__167, 1, func(___ interface{}) interface{} {
						return f
					})
				}(&cljs_core.AFn{}, old_parser))
				return old_parser
			}
		})
	}(&cljs_core.AFn{})

	Deregister_default_tag_parser_BANG_ = func(deregister_default_tag_parser_BANG_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(deregister_default_tag_parser_BANG_, 0, func() interface{} {
			{
				var old_parser = cljs_core.Deref.X_invoke_Arity1(X_STAR_default_data_reader_fn_STAR_)
				_ = old_parser
				cljs_core.Swap_BANG_.X_invoke_Arity2(X_STAR_default_data_reader_fn_STAR_, func(G__168 *cljs_core.AFn, old_parser interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__168, 1, func(___ interface{}) interface{} {
						return nil
					})
				}(&cljs_core.AFn{}, old_parser))
				return old_parser
			}
		})
	}(&cljs_core.AFn{})

}

type CljsReaderPushbackReader interface {
	CljsReaderPushbackReader__()
	Read_char_Arity1() interface{}
	Unread_Arity2(ch interface{}) interface{}
}

var Read_char *cljs_core.AFn

var Unread *cljs_core.AFn

type CljsReaderStringPushbackReader struct {
	S      interface{}
	Buffer interface{}
	Idx    interface{}
}

func (_ *CljsReaderStringPushbackReader) CljsReaderPushbackReader__() {}
func (reader *CljsReaderStringPushbackReader) Read_char_Arity1() interface{} {
	if cljs_core.Alength_(reader.Buffer) == float64(0) {
		reader.Idx = (reader.Idx.(float64) + float64(1))

		return cljs_core.Aget_(reader.S, reader.Idx.(float64))
	} else {
		return js.JSArray_(&reader.Buffer).Pop()
	}
}

func (reader *CljsReaderStringPushbackReader) Unread_Arity2(ch interface{}) interface{} {
	return js.JSArray_(&reader.Buffer).Push(ch)
}

var X__GT_StringPushbackReader *cljs_core.AFn

var Push_back_reader *cljs_core.AFn

// Checks whether a given character is whitespace
var Whitespace_QMARK_ *cljs_core.AFn

// Checks whether a given character is numeric
var Numeric_QMARK_ *cljs_core.AFn

// Checks whether the character begins a comment.
var Comment_prefix_QMARK_ *cljs_core.AFn

// Checks whether the reader is at the start of a number literal
var Number_literal_QMARK_ *cljs_core.AFn

// @param {...*} var_args
var Reader_error *cljs_core.AFn

var Macro_terminating_QMARK_ *cljs_core.AFn

// Advances the reader to the end of a line. Returns the reader
var Skip_line *cljs_core.AFn

var Int_pattern interface{}

var Ratio_pattern interface{}

var Float_pattern interface{}

var Symbol_pattern interface{}

var Re_matches_STAR_ *cljs_core.AFn

var Match_int *cljs_core.AFn

var Match_ratio *cljs_core.AFn

var Match_float *cljs_core.AFn

var Match_number *cljs_core.AFn

var Escape_char_map *cljs_core.AFn

var Unicode_2_pattern interface{}

var Unicode_4_pattern interface{}

var Validate_unicode_escape *cljs_core.AFn

var Make_unicode_char *cljs_core.AFn

var Escape_char *cljs_core.AFn

// Read until first character that doesn't match pred, returning
// char.
var Read_past *cljs_core.AFn

var Read_delimited_list *cljs_core.AFn

var Not_implemented *cljs_core.AFn

var Read_dispatch *cljs_core.AFn

var Read_unmatched_delimiter *cljs_core.AFn

var Read_list *cljs_core.AFn

var Read_comment interface{}

var Read_vector *cljs_core.AFn

var Read_map *cljs_core.AFn

var Read_number *cljs_core.AFn

var Read_string_STAR_ *cljs_core.AFn

var Read_raw_string_STAR_ *cljs_core.AFn

var Special_symbols *cljs_core.AFn

var Read_symbol *cljs_core.AFn

var Read_keyword *cljs_core.AFn

var Desugar_meta *cljs_core.AFn

var Wrapping_reader *cljs_core.AFn

var Throwing_reader *cljs_core.AFn

var Read_meta *cljs_core.AFn

var Read_set *cljs_core.AFn

var Read_regex *cljs_core.AFn

var Read_discard *cljs_core.AFn

var Dispatch_macros *cljs_core.AFn

// Reads the first object from a PushbackReader. Returns the object read.
// If EOF, throws if eof-is-error is true. Otherwise returns sentinel.
var Read *cljs_core.AFn

// Reads one object from the string s
var Read_string *cljs_core.AFn

var Zero_fill_right_and_truncate *cljs_core.AFn

var Divisible_QMARK_ *cljs_core.AFn

var Indivisible_QMARK_ *cljs_core.AFn

var Leap_year_QMARK_ *cljs_core.AFn

var Timestamp_regex interface{}

var Parse_int *cljs_core.AFn

var Check *cljs_core.AFn

var Parse_and_validate_timestamp *cljs_core.AFn

var X_STAR_default_data_reader_fn_STAR_ *cljs_core.CljsCoreAtom

var Maybe_read_tagged_type *cljs_core.AFn

var Register_tag_parser_BANG_ *cljs_core.AFn

var Deregister_tag_parser_BANG_ *cljs_core.AFn

var Register_default_tag_parser_BANG_ *cljs_core.AFn

var Deregister_default_tag_parser_BANG_ *cljs_core.AFn
