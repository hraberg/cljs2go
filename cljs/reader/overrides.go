// Compiled by ClojureScript to Go 0.0-2371
// cljs.reader

// Go overrides.
package reader

import (
	"reflect"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	goog_string "github.com/hraberg/cljs2go/goog/string"
	"github.com/hraberg/cljs2go/js"
)

func init() {
	Read_2_chars = func(read_2_chars *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_2_chars, 1, func(reader interface{}) interface{} {
			return (`` + cljs_core.Str.X_invoke_Arity1(cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()).(string) + cljs_core.Str.X_invoke_Arity1(cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()).(string))
		})
	}(&cljs_core.AFn{})

	Read_4_chars = func(read_4_chars *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_4_chars, 1, func(reader interface{}) interface{} {
			return (`` + cljs_core.Str.X_invoke_Arity1(cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()).(string) + cljs_core.Str.X_invoke_Arity1(cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()).(string) + cljs_core.Str.X_invoke_Arity1(cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()).(string) + cljs_core.Str.X_invoke_Arity1(cljs_core.Decorate_(reader).(CljsReaderPushbackReader).Read_char_Arity1()).(string))
		})
	}(&cljs_core.AFn{})

	Read_token = func(read_token *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_token, 2, func(rdr interface{}, initch interface{}) interface{} {
			{
				var sb = (&goog_string.StringBuffer{initch})
				var ch interface{} = cljs_core.Decorate_(rdr).(CljsReaderPushbackReader).Read_char_Arity1()
				_, _ = sb, ch
				for {
					if (cljs_core.Nil_(ch)) || (Whitespace_QMARK_.Arity1IB(ch)) || (Macro_terminating_QMARK_.Arity1IB(ch)) {
						cljs_core.Decorate_(rdr).(CljsReaderPushbackReader).Unread_Arity2(ch)
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(sb, "ToString", []interface{}{})
					} else {
						sb, ch = func() *goog_string.StringBuffer {
							cljs_core.Native_invoke_instance_method.X_invoke_Arity3(sb, "Append", []interface{}{ch})
							return sb
						}(), cljs_core.Decorate_(rdr).(CljsReaderPushbackReader).Read_char_Arity1()
						continue
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Macros = func(macros *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(macros, 1, func(c interface{}) interface{} {
			if reflect.DeepEqual(c, "\"") {
				return Read_string_STAR_
			} else {
				if reflect.DeepEqual(c, ":") {
					return Read_keyword
				} else {
					if reflect.DeepEqual(c, ";") {
						return Read_comment
					} else {
						if reflect.DeepEqual(c, "'") {
							return Wrapping_reader.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "quote", Str: "quote", X_hash: float64(1377916282), X_meta: nil})).(cljs_core.CljsCoreIFn)
						} else {
							if reflect.DeepEqual(c, "@") {
								return Wrapping_reader.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "deref", Str: "deref", X_hash: float64(1494944732), X_meta: nil})).(cljs_core.CljsCoreIFn)
							} else {
								if reflect.DeepEqual(c, "^") {
									return Read_meta
								} else {
									if reflect.DeepEqual(c, "`") {
										return Not_implemented
									} else {
										if reflect.DeepEqual(c, "~") {
											return Not_implemented
										} else {
											if reflect.DeepEqual(c, "(") {
												return Read_list
											} else {
												if reflect.DeepEqual(c, ")") {
													return Read_unmatched_delimiter
												} else {
													if reflect.DeepEqual(c, "[") {
														return Read_vector
													} else {
														if reflect.DeepEqual(c, "]") {
															return Read_unmatched_delimiter
														} else {
															if reflect.DeepEqual(c, "{") {
																return Read_map
															} else {
																if reflect.DeepEqual(c, "}") {
																	return Read_unmatched_delimiter
																} else {
																	if reflect.DeepEqual(c, "\\") {
																		return func(G__169 *cljs_core.AFn) *cljs_core.AFn {
																			return cljs_core.Fn(G__169, 2, func(rdr interface{}, ___ interface{}) interface{} {
																				return cljs_core.Decorate_(rdr).(CljsReaderPushbackReader).Read_char_Arity1()
																			})
																		}(&cljs_core.AFn{})
																	} else {
																		if reflect.DeepEqual(c, "#") {
																			return Read_dispatch
																		} else {
																			return nil

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
								}
							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Days_in_month = func() interface{} {
		var dim_norm = (&cljs_core.CljsCorePersistentVector{nil, float64(13), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, float64(31), float64(28), float64(31), float64(30), float64(31), float64(30), float64(31), float64(31), float64(30), float64(31), float64(30), float64(31)}, nil})
		var dim_leap = (&cljs_core.CljsCorePersistentVector{nil, float64(13), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, float64(31), float64(29), float64(31), float64(30), float64(31), float64(30), float64(31), float64(31), float64(30), float64(31), float64(30), float64(31)}, nil})
		_, _ = dim_norm, dim_leap
		return cljs_core.Identity.X_invoke_Arity1(func(G__170 *cljs_core.AFn, dim_norm cljs_core.CljsCoreIVector, dim_leap cljs_core.CljsCoreIVector) *cljs_core.AFn {
			return cljs_core.Fn(G__170, 2, func(month interface{}, leap_year_QMARK_ interface{}) interface{} {
				return cljs_core.Get.X_invoke_Arity2(func() cljs_core.CljsCoreIVector {
					if cljs_core.Truth_(leap_year_QMARK_) {
						return dim_leap
					} else {
						return dim_norm
					}
				}(), month)
			})
		}(&cljs_core.AFn{}, dim_norm, dim_leap))
	}()

	Parse_timestamp = func(parse_timestamp *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(parse_timestamp, 1, func(ts interface{}) interface{} {
			{
				var temp__4386__auto__ = Parse_and_validate_timestamp.X_invoke_Arity1(ts).(cljs_core.CljsCoreIVector)
				_ = temp__4386__auto__
				if cljs_core.Truth_(temp__4386__auto__) {
					{
						var vec__172 = temp__4386__auto__
						var years = cljs_core.Nth.X_invoke_Arity3(vec__172, float64(0), nil)
						var months = cljs_core.Nth.X_invoke_Arity3(vec__172, float64(1), nil)
						var days = cljs_core.Nth.X_invoke_Arity3(vec__172, float64(2), nil)
						var hours = cljs_core.Nth.X_invoke_Arity3(vec__172, float64(3), nil)
						var minutes = cljs_core.Nth.X_invoke_Arity3(vec__172, float64(4), nil)
						var seconds = cljs_core.Nth.X_invoke_Arity3(vec__172, float64(5), nil)
						var ms = cljs_core.Nth.X_invoke_Arity3(vec__172, float64(6), nil)
						var offset = cljs_core.Nth.X_invoke_Arity3(vec__172, float64(7), nil)
						_, _, _, _, _, _, _, _, _ = vec__172, years, months, days, hours, minutes, seconds, ms, offset
						return (&js.Date{ts})
					}
				} else {
					return Reader_error.X_invoke_ArityVariadic(nil, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{("Unrecognized date/time syntax: " + cljs_core.Str.X_invoke_Arity1(ts).(string))}))
				}
			}
		})
	}(&cljs_core.AFn{})

	Read_queue = func(read_queue *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_queue, 1, func(elems interface{}) interface{} {
			if cljs_core.Vector_QMARK_.Arity1IB(elems) {
				return cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, elems)
			} else {
				return Reader_error.X_invoke_ArityVariadic(nil, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Queue literal expects a vector for its elements."}))
			}
		})
	}(&cljs_core.AFn{})

	Read_date = func(read_date *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_date, 1, func(s interface{}) interface{} {
			if cljs_core.Value_(s).Kind() == reflect.String {
				return Parse_timestamp.X_invoke_Arity1(s).(*js.Date)
			} else {
				return Reader_error.X_invoke_ArityVariadic(nil, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Instance literal expects a string for its timestamp."}))
			}
		})
	}(&cljs_core.AFn{})

	Read_uuid = func(read_uuid *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(read_uuid, 1, func(uuid interface{}) interface{} {
			if cljs_core.Value_(uuid).Kind() == reflect.String {
				return (&cljs_core.CljsCoreUUID{uuid})
			} else {
				return Reader_error.X_invoke_ArityVariadic(nil, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"UUID literal expects a string as its representation."}))
			}
		})
	}(&cljs_core.AFn{})

	X_STAR_tag_table_STAR_ = cljs_core.Atom.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{"inst", Read_date, "uuid", Read_uuid, "queue", Read_queue}, nil})).(*cljs_core.CljsCoreAtom)

}

var Read_2_chars *cljs_core.AFn

var Read_4_chars *cljs_core.AFn

var Read_token *cljs_core.AFn

var Macros *cljs_core.AFn

var Days_in_month interface{}

var Parse_timestamp *cljs_core.AFn

var Read_queue *cljs_core.AFn

var Read_date *cljs_core.AFn

var Read_uuid *cljs_core.AFn

var X_STAR_tag_table_STAR_ *cljs_core.CljsCoreAtom
