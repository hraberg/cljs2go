// Compiled by ClojureScript to Go 0.0-2322
// clojure.string

package string

import (
	"reflect"

	cljs_core "github.com/hraberg/cljs.go/cljs/core"
	"github.com/hraberg/cljs.go/js"

	goog_string "github.com/hraberg/cljs.go/goog/string"
)

func init() {
	Seq_reverse = func(seq_reverse *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(seq_reverse, 1, func(coll interface{}) interface{} {
			return cljs_core.Reduce.X_invoke_Arity3(cljs_core.Conj, cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), coll)
		})
	}(&cljs_core.AFn{})

	Reverse = func(reverse *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(reverse, 1, func(s interface{}) interface{} {
			return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "Split", []interface{}{""}), "Reverse", []interface{}{}), "Join", []interface{}{""})
		})
	}(&cljs_core.AFn{})

	Replace = func(replace *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(replace, 3, func(s interface{}, match interface{}, replacement interface{}) interface{} {
			if reflect.ValueOf(match).Kind() == reflect.String {
				return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "Replace", []interface{}{(&js.RegExp{func() interface{} {
					var G__2 = match
					_ = G__2
					return cljs_core.Native_invoke_func.X_invoke_Arity2(goog_string.RegExpEscape, []interface{}{G__2})
				}(), "g"}), replacement})
			} else {
				if cljs_core.Truth_(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(match, "HasOwnProperty", []interface{}{"source"})) {
					return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "Replace", []interface{}{(&js.RegExp{cljs_core.Native_get_instance_field.X_invoke_Arity2(match, "Source"), "g"}), replacement})
				} else {
					panic(("Invalid match arg: " + cljs_core.Str.X_invoke_Arity1(match).(string)))

				}
			}
		})
	}(&cljs_core.AFn{})

	Replace_first = func(replace_first *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(replace_first, 3, func(s interface{}, match interface{}, replacement interface{}) interface{} {
			return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "Replace", []interface{}{match, replacement})
		})
	}(&cljs_core.AFn{})

	Join = func(join *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(join, 2, func(coll interface{}) interface{} {
			return cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, coll)
		}, func(separator interface{}, coll interface{}) interface{} {
			return cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Interpose.X_invoke_Arity2(separator, coll).(*cljs_core.CljsCoreLazySeq))
		})
	}(&cljs_core.AFn{})

	Upper_case = func(upper_case *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(upper_case, 1, func(s interface{}) interface{} {
			return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "ToUpperCase", []interface{}{})
		})
	}(&cljs_core.AFn{})

	Lower_case = func(lower_case *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(lower_case, 1, func(s interface{}) interface{} {
			return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "ToLowerCase", []interface{}{})
		})
	}(&cljs_core.AFn{})

	Capitalize = func(capitalize *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(capitalize, 1, func(s interface{}) interface{} {
			if cljs_core.Count.X_invoke_Arity1(s).(float64) < float64(2) {
				return Upper_case.X_invoke_Arity1(s)
			} else {
				return (`` + cljs_core.Str.X_invoke_Arity1(Upper_case.X_invoke_Arity1(cljs_core.Subs.X_invoke_Arity3(s, float64(0), float64(1)))).(string) + cljs_core.Str.X_invoke_Arity1(Lower_case.X_invoke_Arity1(cljs_core.Subs.X_invoke_Arity2(s, float64(1)))).(string))
			}
		})
	}(&cljs_core.AFn{})

	Pop_last_while_empty = func(pop_last_while_empty *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(pop_last_while_empty, 1, func(v interface{}) interface{} {
			{
				var v___1 interface{} = v
				_ = v___1
				for {
					if cljs_core.X_EQ_.Arity2IIB("", cljs_core.Peek.X_invoke_Arity1(v___1)) {
						v___1 = cljs_core.Pop.X_invoke_Arity1(v___1)
						continue
					} else {
						return v___1
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Discard_trailing_if_needed = func(discard_trailing_if_needed *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(discard_trailing_if_needed, 2, func(limit interface{}, v interface{}) interface{} {
			if cljs_core.X_EQ_.Arity2IIB(float64(0), limit) {
				return Pop_last_while_empty.X_invoke_Arity1(v)
			} else {
				return v
			}
		})
	}(&cljs_core.AFn{})

	Split_with_empty_regex = func(split_with_empty_regex *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(split_with_empty_regex, 2, func(s interface{}, limit interface{}) interface{} {
			if (limit.(float64) <= float64(0)) || (limit.(float64) >= (float64(2) + cljs_core.Count.X_invoke_Arity1(s).(float64))) {
				return cljs_core.Conj.X_invoke_Arity2(cljs_core.Vec.X_invoke_Arity1(cljs_core.Cons.X_invoke_Arity2("", cljs_core.Map_.X_invoke_Arity2(cljs_core.Str, cljs_core.Seq.Arity1IQ(s)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)), "")
			} else {
				{
					var pred__10 = cljs_core.X_EQ_
					var expr__11 = limit
					_, _ = pred__10, expr__11
					if cljs_core.Truth_(func() interface{} {
						var G__13 = float64(1)
						var G__14 = expr__11
						_, _ = G__13, G__14
						return pred__10.X_invoke_Arity2(G__13, G__14)
					}()) {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{s}, nil})
					} else {
						if cljs_core.Truth_(func() interface{} {
							var G__15 = float64(2)
							var G__16 = expr__11
							_, _ = G__15, G__16
							return pred__10.X_invoke_Arity2(G__15, G__16)
						}()) {
							return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"", s}, nil})
						} else {
							{
								var c = (limit.(float64) - float64(2))
								_ = c
								return cljs_core.Conj.X_invoke_Arity2(cljs_core.Vec.X_invoke_Arity1(cljs_core.Cons.X_invoke_Arity2("", cljs_core.Subvec.X_invoke_Arity3(cljs_core.Vec.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(cljs_core.Str, cljs_core.Seq.Arity1IQ(s)).(*cljs_core.CljsCoreLazySeq)), float64(0), c).(*cljs_core.CljsCoreSubvec)).(*cljs_core.CljsCoreCons)), cljs_core.Subs.X_invoke_Arity2(s, c))
							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Split = func(split *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(split, 3, func(s interface{}, re interface{}) interface{} {
			return split.X_invoke_Arity3(s, re, float64(0))
		}, func(s interface{}, re interface{}, limit interface{}) interface{} {
			return Discard_trailing_if_needed.X_invoke_Arity2(limit, func() interface{} {
				if cljs_core.X_EQ_.Arity2IIB((`` + cljs_core.Str.X_invoke_Arity1(re).(string)), "/(?:)/") {
					return Split_with_empty_regex.X_invoke_Arity2(s, limit)
				} else {
					return func() interface{} {
						if limit.(float64) < float64(1) {
							return cljs_core.Vec.X_invoke_Arity1(js.JSString_((`` + cljs_core.Str.X_invoke_Arity1(s).(string))).Split(re))
						} else {
							return func() interface{} {
								var s___1 interface{} = s
								var limit___1 interface{} = limit
								var parts interface{} = cljs_core.CljsCorePersistentVector_EMPTY
								_, _, _ = s___1, limit___1, parts
								for {
									if cljs_core.X_EQ_.Arity2IIB(limit___1, float64(1)) {
										return cljs_core.Conj.X_invoke_Arity2(parts, s___1)
									} else {
										{
											var temp__4220__auto__ = cljs_core.Re_find.X_invoke_Arity2(re, s___1)
											_ = temp__4220__auto__
											if cljs_core.Truth_(temp__4220__auto__) {
												{
													var m = temp__4220__auto__
													_ = m
													{
														var index = cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s___1, "IndexOf", []interface{}{m})
														_ = index
														s___1, limit___1, parts = cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s___1, "Substring", []interface{}{(index.(float64) + cljs_core.Count.X_invoke_Arity1(m).(float64))}), (limit___1.(float64) - float64(1)), cljs_core.Conj.X_invoke_Arity2(parts, cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s___1, "Substring", []interface{}{float64(0), index}))
														continue
													}
												}
											} else {
												return cljs_core.Conj.X_invoke_Arity2(parts, s___1)
											}
										}
									}
								}
							}()
						}
					}()
				}
			}())
		})
	}(&cljs_core.AFn{})

	Split_lines = func(split_lines *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(split_lines, 1, func(s interface{}) interface{} {
			return Split.X_invoke_Arity2(s, (&js.RegExp{Pattern: `\n|\r\n`, Flags: ``}))
		})
	}(&cljs_core.AFn{})

	Trim = func(trim *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(trim, 1, func(s interface{}) interface{} {
			{
				var G__21 = s
				_ = G__21
				return cljs_core.Native_invoke_func.X_invoke_Arity2(goog_string.Trim, []interface{}{G__21})
			}
		})
	}(&cljs_core.AFn{})

	Triml = func(triml *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(triml, 1, func(s interface{}) interface{} {
			{
				var G__23 = s
				_ = G__23
				return cljs_core.Native_invoke_func.X_invoke_Arity2(goog_string.TrimLeft, []interface{}{G__23})
			}
		})
	}(&cljs_core.AFn{})

	Trimr = func(trimr *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(trimr, 1, func(s interface{}) interface{} {
			{
				var G__25 = s
				_ = G__25
				return cljs_core.Native_invoke_func.X_invoke_Arity2(goog_string.TrimRight, []interface{}{G__25})
			}
		})
	}(&cljs_core.AFn{})

	Trim_newline = func(trim_newline *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(trim_newline, 1, func(s interface{}) interface{} {
			{
				var index interface{} = cljs_core.Native_get_instance_field.X_invoke_Arity2(s, "Length")
				_ = index
				for {
					if index.(float64) == float64(0) {
						return ""
					} else {
						{
							var ch = cljs_core.Get.X_invoke_Arity2(s, (index.(float64) - float64(1)))
							_ = ch
							if (cljs_core.X_EQ_.Arity2IIB(ch, "\n")) || (cljs_core.X_EQ_.Arity2IIB(ch, "\r")) {
								index = (index.(float64) - float64(1))
								continue
							} else {
								return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "Substring", []interface{}{float64(0), index})
							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

	Blank_QMARK_ = func(blank_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(blank_QMARK_, 1, func(s interface{}) interface{} {
			{
				var G__27 = s
				_ = G__27
				return cljs_core.Native_invoke_func.X_invoke_Arity2(goog_string.IsEmptySafe, []interface{}{G__27})
			}
		})
	}(&cljs_core.AFn{})

	Escape = func(escape___1 *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(escape___1, 2, func(s interface{}, cmap interface{}) interface{} {
			{
				var buffer = (&goog_string.StringBuffer{})
				var length = cljs_core.Native_get_instance_field.X_invoke_Arity2(s, "Length")
				_, _ = buffer, length
				{
					var index = float64(0)
					_ = index
					for {
						if cljs_core.X_EQ_.Arity2IIB(length, index) {
							return buffer.ToString()
						} else {
							{
								var ch = cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "CharAt", []interface{}{index})
								_ = ch
								{
									var temp__4220__auto___28 = cljs_core.Get.X_invoke_Arity2(cmap, ch)
									_ = temp__4220__auto___28
									if cljs_core.Truth_(temp__4220__auto___28) {
										{
											var replacement_29 = temp__4220__auto___28
											_ = replacement_29
											buffer.Append((`` + cljs_core.Str.X_invoke_Arity1(replacement_29).(string)))
										}
									} else {
										buffer.Append(ch)
									}
								}
								index = (index + float64(1))
								continue
							}
						}
					}
				}
			}
		})
	}(&cljs_core.AFn{})

}

var Seq_reverse *cljs_core.AFn

// Returns s with its characters reversed.
var Reverse *cljs_core.AFn

// Replaces all instance of match with replacement in s.
// match/replacement can be:
//
// string / string
// pattern / (string or function of match).
var Replace *cljs_core.AFn

// Replaces the first instance of match with replacement in s.
// match/replacement can be:
//
// string / string
// pattern / (string or function of match).
var Replace_first *cljs_core.AFn

// Returns a string of all elements in coll, as returned by (seq coll),
// separated by an optional separator.
var Join *cljs_core.AFn

// Converts string to all upper-case.
var Upper_case *cljs_core.AFn

// Converts string to all lower-case.
var Lower_case *cljs_core.AFn

// Converts first character of the string to upper-case, all other
// characters to lower-case.
var Capitalize *cljs_core.AFn

var Pop_last_while_empty *cljs_core.AFn

var Discard_trailing_if_needed *cljs_core.AFn

var Split_with_empty_regex *cljs_core.AFn

// Splits string on a regular expression. Optional argument limit is
// the maximum number of splits. Not lazy. Returns vector of the splits.
var Split *cljs_core.AFn

// Splits s on
// or
// .
var Split_lines *cljs_core.AFn

// Removes whitespace from both ends of string.
var Trim *cljs_core.AFn

// Removes whitespace from the left side of string.
var Triml *cljs_core.AFn

// Removes whitespace from the right side of string.
var Trimr *cljs_core.AFn

// Removes all trailing newline \n or return \r characters from
// string.  Similar to Perl's chomp.
var Trim_newline *cljs_core.AFn

// True is s is nil, empty, or contains only whitespace.
var Blank_QMARK_ *cljs_core.AFn

// Return a new string, using cmap to escape each character ch
// from s as follows:
//
// If (cmap ch) is nil, append ch to the new string.
// If (cmap ch) is non-nil, append (str (cmap ch)) instead.
var Escape *cljs_core.AFn
