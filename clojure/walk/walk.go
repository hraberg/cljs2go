// Compiled by ClojureScript to Go 0.0-2322
// clojure.walk

// This file defines a generic tree walker for Clojure data
// structures.  It takes any data structure (list, vector, map, set,
// seq), calls a function on every element, and uses the return value
// of the function in place of the original.  This makes it fairly
// easy to write recursive search-and-replace functions, as shown in
// the examples.
//
// Note: "walk" supports all Clojure data structures EXCEPT maps
// created with sorted-map-by.  There is no (obvious) way to retrieve
// the sorting function.
// Author: Stuart Sierra
package walk

import (
	"reflect"

	cljs_core "github.com/hraberg/cljs.go/cljs/core"
)

func init() {
	Walk = func(walk *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(walk, 3, func(inner interface{}, outer interface{}, form interface{}) interface{} {
			if cljs_core.Seq_QMARK_.Arity1IB(form) {
				{
					var G__4 = cljs_core.Doall.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(inner, form).(*cljs_core.CljsCoreLazySeq))
					_ = G__4
					return outer.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4)
				}
			} else {
				if cljs_core.Coll_QMARK_.Arity1IB(form) {
					{
						var G__5 = cljs_core.Into.X_invoke_Arity2(cljs_core.Empty.X_invoke_Arity1(form), cljs_core.Map_.X_invoke_Arity2(inner, form).(*cljs_core.CljsCoreLazySeq))
						_ = G__5
						return outer.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__5)
					}
				} else {
					{
						var G__6 = form
						_ = G__6
						return outer.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__6)
					}

				}
			}
		})
	}(&cljs_core.AFn{})

	Postwalk = func(postwalk *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(postwalk, 2, func(f interface{}, form interface{}) interface{} {
			return Walk.X_invoke_Arity3(cljs_core.Partial.X_invoke_Arity2(postwalk, f).(cljs_core.CljsCoreIFn), f, form)
		})
	}(&cljs_core.AFn{})

	Prewalk = func(prewalk *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(prewalk, 2, func(f interface{}, form interface{}) interface{} {
			return Walk.X_invoke_Arity3(cljs_core.Partial.X_invoke_Arity2(prewalk, f).(cljs_core.CljsCoreIFn), cljs_core.Identity, func() interface{} {
				var G__8 = form
				_ = G__8
				return f.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__8)
			}())
		})
	}(&cljs_core.AFn{})

	Keywordize_keys = func(keywordize_keys *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(keywordize_keys, 1, func(m interface{}) interface{} {
			{
				var f = func(G__13 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__13, 1, func(p__11 interface{}) interface{} {
						{
							var vec__12 = p__11
							var k = cljs_core.Nth.X_invoke_Arity3(vec__12, float64(0), nil)
							var v = cljs_core.Nth.X_invoke_Arity3(vec__12, float64(1), nil)
							_, _, _ = vec__12, k, v
							if cljs_core.Value_(k).Kind() == reflect.String {
								return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Keyword.X_invoke_Arity1(k), v}, nil})
							} else {
								return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{k, v}, nil})
							}
						}
					})
				}(&cljs_core.AFn{})
				_ = f
				return Postwalk.X_invoke_Arity2(func(G__14 *cljs_core.AFn, f cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__14, 1, func(x interface{}) interface{} {
						if cljs_core.Map_QMARK_.Arity1IB(x) {
							return cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Map_.X_invoke_Arity2(f, x).(*cljs_core.CljsCoreLazySeq))
						} else {
							return x
						}
					})
				}(&cljs_core.AFn{}, f), m)
			}
		})
	}(&cljs_core.AFn{})

	Stringify_keys = func(stringify_keys *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(stringify_keys, 1, func(m interface{}) interface{} {
			{
				var f = func(G__19 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__19, 1, func(p__17 interface{}) interface{} {
						{
							var vec__18 = p__17
							var k = cljs_core.Nth.X_invoke_Arity3(vec__18, float64(0), nil)
							var v = cljs_core.Nth.X_invoke_Arity3(vec__18, float64(1), nil)
							_, _, _ = vec__18, k, v
							if cljs_core.Value_(k).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
								return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Name.X_invoke_Arity1(k), v}, nil})
							} else {
								return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{k, v}, nil})
							}
						}
					})
				}(&cljs_core.AFn{})
				_ = f
				return Postwalk.X_invoke_Arity2(func(G__20 *cljs_core.AFn, f cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__20, 1, func(x interface{}) interface{} {
						if cljs_core.Map_QMARK_.Arity1IB(x) {
							return cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Map_.X_invoke_Arity2(f, x).(*cljs_core.CljsCoreLazySeq))
						} else {
							return x
						}
					})
				}(&cljs_core.AFn{}, f), m)
			}
		})
	}(&cljs_core.AFn{})

	Prewalk_replace = func(prewalk_replace *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(prewalk_replace, 2, func(smap interface{}, form interface{}) interface{} {
			return Prewalk.X_invoke_Arity2(func(G__23 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__23, 1, func(x interface{}) interface{} {
					if cljs_core.Contains_QMARK_.Arity2IIB(smap, x) {
						{
							var G__22 = x
							_ = G__22
							return smap.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__22)
						}
					} else {
						return x
					}
				})
			}(&cljs_core.AFn{}), form)
		})
	}(&cljs_core.AFn{})

	Postwalk_replace = func(postwalk_replace *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(postwalk_replace, 2, func(smap interface{}, form interface{}) interface{} {
			return Postwalk.X_invoke_Arity2(func(G__26 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__26, 1, func(x interface{}) interface{} {
					if cljs_core.Contains_QMARK_.Arity2IIB(smap, x) {
						{
							var G__25 = x
							_ = G__25
							return smap.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__25)
						}
					} else {
						return x
					}
				})
			}(&cljs_core.AFn{}), form)
		})
	}(&cljs_core.AFn{})

}

// Traverses form, an arbitrary data structure.  inner and outer are
// functions.  Applies inner to each element of form, building up a
// data structure of the same type, then applies outer to the result.
// Recognizes all Clojure data structures. Consumes seqs as with doall.
var Walk *cljs_core.AFn

// Performs a depth-first, post-order traversal of form.  Calls f on
// each sub-form, uses f's return value in place of the original.
// Recognizes all Clojure data structures. Consumes seqs as with doall.
var Postwalk *cljs_core.AFn

// Like postwalk, but does pre-order traversal.
var Prewalk *cljs_core.AFn

// Recursively transforms all map keys from strings to keywords.
var Keywordize_keys *cljs_core.AFn

// Recursively transforms all map keys from keywords to strings.
var Stringify_keys *cljs_core.AFn

// Recursively transforms form by replacing keys in smap with their
// values.  Like clojure/replace but works on any data structure.  Does
// replacement at the root of the tree first.
var Prewalk_replace *cljs_core.AFn

// Recursively transforms form by replacing keys in smap with their
// values.  Like clojure/replace but works on any data structure.  Does
// replacement at the leaves of the tree first.
var Postwalk_replace *cljs_core.AFn
