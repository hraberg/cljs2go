// Compiled by ClojureScript to Go 0.0-2371
// clojure.set

// Set operations such as union/intersection.
// Author: Rich Hickey
package set

import (
	"reflect"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
)

func init() {
	Bubble_max_key = func(bubble_max_key *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(bubble_max_key, 2, func(k interface{}, coll interface{}) interface{} {
			{
				var max = cljs_core.Apply.X_invoke_Arity3(cljs_core.Max_key, k, coll)
				_ = max
				return cljs_core.Cons.X_invoke_Arity2(max, cljs_core.Remove.X_invoke_Arity2(func(G__2 *cljs_core.AFn, max interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__2, 1, func(p1__1_SHARP_ interface{}) interface{} {
						return reflect.DeepEqual(max, p1__1_SHARP_)
					})
				}(&cljs_core.AFn{}, max), coll).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
			}
		})
	}(&cljs_core.AFn{})

	Union = func(union *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(union, 2, func() interface{} {
			return cljs_core.CljsCorePersistentHashSet_EMPTY
		}, func(s1 interface{}) interface{} {
			return s1
		}, func(s1 interface{}, s2 interface{}) interface{} {
			if cljs_core.Count.X_invoke_Arity1(s1).(float64) < cljs_core.Count.X_invoke_Arity1(s2).(float64) {
				return cljs_core.Reduce.X_invoke_Arity3(cljs_core.Conj, s2, s1)
			} else {
				return cljs_core.Reduce.X_invoke_Arity3(cljs_core.Conj, s1, s2)
			}
		}, func(s1_s2_sets__ ...interface{}) interface{} {
			var s1 = s1_s2_sets__[0]
			var s2 = s1_s2_sets__[1]
			var sets = cljs_core.Seq.Arity1IQ(s1_s2_sets__[2])
			_, _, _ = s1, s2, sets
			{
				var bubbled_sets = Bubble_max_key.X_invoke_Arity2(cljs_core.Count, cljs_core.Conj.X_invoke_ArityVariadic(sets, s2, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{s1}))).(*cljs_core.CljsCoreCons)
				_ = bubbled_sets
				return cljs_core.Reduce.X_invoke_Arity3(cljs_core.Into, cljs_core.First.X_invoke_Arity1(bubbled_sets), cljs_core.Rest.Arity1IQ(bubbled_sets))
			}
		})
	}(&cljs_core.AFn{})

	Intersection = func(intersection *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(intersection, 2, func(s1 interface{}) interface{} {
			return s1
		}, func(s1 interface{}, s2 interface{}) interface{} {
			for {
				if cljs_core.Count.X_invoke_Arity1(s2).(float64) < cljs_core.Count.X_invoke_Arity1(s1).(float64) {
					s1, s2 = s2, s1
					continue
				} else {
					return cljs_core.Reduce.X_invoke_Arity3(func(G__4 *cljs_core.AFn, s1 interface{}, s2 interface{}) *cljs_core.AFn {
						return cljs_core.Fn(G__4, 2, func(result interface{}, item interface{}) interface{} {
							if cljs_core.Contains_QMARK_.Arity2IIB(s2, item) {
								return result
							} else {
								return cljs_core.Disj.X_invoke_Arity2(result, item)
							}
						})
					}(&cljs_core.AFn{}, s1, s2), s1, s1)
				}
			}
		}, func(s1_s2_sets__ ...interface{}) interface{} {
			var s1 = s1_s2_sets__[0]
			var s2 = s1_s2_sets__[1]
			var sets = cljs_core.Seq.Arity1IQ(s1_s2_sets__[2])
			_, _, _ = s1, s2, sets
			{
				var bubbled_sets = Bubble_max_key.X_invoke_Arity2(func(G__5 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5, 1, func(p1__3_SHARP_ interface{}) interface{} {
						return (-cljs_core.Count.X_invoke_Arity1(p1__3_SHARP_).(float64))
					})
				}(&cljs_core.AFn{}), cljs_core.Conj.X_invoke_ArityVariadic(sets, s2, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{s1}))).(*cljs_core.CljsCoreCons)
				_ = bubbled_sets
				return cljs_core.Reduce.X_invoke_Arity3(intersection, cljs_core.First.X_invoke_Arity1(bubbled_sets), cljs_core.Rest.Arity1IQ(bubbled_sets))
			}
		})
	}(&cljs_core.AFn{})

	Difference = func(difference *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(difference, 2, func(s1 interface{}) interface{} {
			return s1
		}, func(s1 interface{}, s2 interface{}) interface{} {
			if cljs_core.Count.X_invoke_Arity1(s1).(float64) < cljs_core.Count.X_invoke_Arity1(s2).(float64) {
				return cljs_core.Reduce.X_invoke_Arity3(func(G__6 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__6, 2, func(result interface{}, item interface{}) interface{} {
						if cljs_core.Contains_QMARK_.Arity2IIB(s2, item) {
							return cljs_core.Disj.X_invoke_Arity2(result, item)
						} else {
							return result
						}
					})
				}(&cljs_core.AFn{}), s1, s1)
			} else {
				return cljs_core.Reduce.X_invoke_Arity3(cljs_core.Disj, s1, s2)
			}
		}, func(s1_s2_sets__ ...interface{}) interface{} {
			var s1 = s1_s2_sets__[0]
			var s2 = s1_s2_sets__[1]
			var sets = cljs_core.Seq.Arity1IQ(s1_s2_sets__[2])
			_, _, _ = s1, s2, sets
			return cljs_core.Reduce.X_invoke_Arity3(difference, s1, cljs_core.Conj.X_invoke_Arity2(sets, s2))
		})
	}(&cljs_core.AFn{})

	Select_ = func(select_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(select_, 2, func(pred interface{}, xset interface{}) interface{} {
			return cljs_core.Reduce.X_invoke_Arity3(func(G__9 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__9, 2, func(s interface{}, k interface{}) interface{} {
					if cljs_core.Truth_(func() interface{} {
						var G__8 = k
						_ = G__8
						return pred.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__8)
					}()) {
						return s
					} else {
						return cljs_core.Disj.X_invoke_Arity2(s, k)
					}
				})
			}(&cljs_core.AFn{}), xset, xset)
		})
	}(&cljs_core.AFn{})

	Project = func(project *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(project, 2, func(xrel interface{}, ks interface{}) interface{} {
			return cljs_core.Set.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(func(G__11 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__11, 1, func(p1__10_SHARP_ interface{}) interface{} {
					return cljs_core.Select_keys.X_invoke_Arity2(p1__10_SHARP_, ks)
				})
			}(&cljs_core.AFn{}), xrel).(*cljs_core.CljsCoreLazySeq))
		})
	}(&cljs_core.AFn{})

	Rename_keys = func(rename_keys *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(rename_keys, 2, func(map_ interface{}, kmap interface{}) interface{} {
			return cljs_core.Reduce.X_invoke_Arity3(func(G__16 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__16, 2, func(m interface{}, p__14 interface{}) interface{} {
					{
						var vec__15 = p__14
						var old = cljs_core.Nth.X_invoke_Arity3(vec__15, float64(0), nil)
						var new = cljs_core.Nth.X_invoke_Arity3(vec__15, float64(1), nil)
						_, _, _ = vec__15, old, new
						if cljs_core.Contains_QMARK_.Arity2IIB(map_, old) {
							return cljs_core.Assoc.X_invoke_Arity3(m, new, cljs_core.Get.X_invoke_Arity2(map_, old))
						} else {
							return m
						}
					}
				})
			}(&cljs_core.AFn{}), cljs_core.Apply.X_invoke_Arity3(cljs_core.Dissoc, map_, cljs_core.Keys.X_invoke_Arity1(kmap)), kmap)
		})
	}(&cljs_core.AFn{})

	Rename = func(rename *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(rename, 2, func(xrel interface{}, kmap interface{}) interface{} {
			return cljs_core.Set.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(func(G__18 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__18, 1, func(p1__17_SHARP_ interface{}) interface{} {
					return Rename_keys.X_invoke_Arity2(p1__17_SHARP_, kmap)
				})
			}(&cljs_core.AFn{}), xrel).(*cljs_core.CljsCoreLazySeq))
		})
	}(&cljs_core.AFn{})

	Index = func(index *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(index, 2, func(xrel interface{}, ks interface{}) interface{} {
			return cljs_core.Reduce.X_invoke_Arity3(func(G__19 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__19, 2, func(m interface{}, x interface{}) interface{} {
					{
						var ik = cljs_core.Select_keys.X_invoke_Arity2(x, ks)
						_ = ik
						return cljs_core.Assoc.X_invoke_Arity3(m, ik, cljs_core.Conj.X_invoke_Arity2(cljs_core.Get.X_invoke_Arity3(m, ik, cljs_core.CljsCorePersistentHashSet_EMPTY), x))
					}
				})
			}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentArrayMap_EMPTY, xrel)
		})
	}(&cljs_core.AFn{})

	Map_invert = func(map_invert *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(map_invert, 1, func(m interface{}) interface{} {
			return cljs_core.Reduce.X_invoke_Arity3(func(G__24 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__24, 2, func(m___1 interface{}, p__22 interface{}) interface{} {
					{
						var vec__23 = p__22
						var k = cljs_core.Nth.X_invoke_Arity3(vec__23, float64(0), nil)
						var v = cljs_core.Nth.X_invoke_Arity3(vec__23, float64(1), nil)
						_, _, _ = vec__23, k, v
						return cljs_core.Assoc.X_invoke_Arity3(m___1, v, k)
					}
				})
			}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentArrayMap_EMPTY, m)
		})
	}(&cljs_core.AFn{})

	Join = func(join *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(join, 3, func(xrel interface{}, yrel interface{}) interface{} {
			if cljs_core.Truth_(func() interface{} {
				var and__165__auto__ = cljs_core.Seq.Arity1IQ(xrel)
				_ = and__165__auto__
				if cljs_core.Truth_(and__165__auto__) {
					return cljs_core.Seq.Arity1IQ(yrel)
				} else {
					return and__165__auto__
				}
			}()) {
				{
					var ks = Intersection.X_invoke_Arity2(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(cljs_core.First.X_invoke_Arity1(xrel))), cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(cljs_core.First.X_invoke_Arity1(yrel))))
					var vec__33 = func() cljs_core.CljsCoreIVector {
						if cljs_core.Count.X_invoke_Arity1(xrel).(float64) <= cljs_core.Count.X_invoke_Arity1(yrel).(float64) {
							return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{xrel, yrel}, nil})
						} else {
							return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{yrel, xrel}, nil})
						}
					}()
					var r = cljs_core.Nth.X_invoke_Arity3(vec__33, float64(0), nil)
					var s = cljs_core.Nth.X_invoke_Arity3(vec__33, float64(1), nil)
					var idx = Index.X_invoke_Arity2(r, ks)
					_, _, _, _, _ = ks, vec__33, r, s, idx
					return cljs_core.Reduce.X_invoke_Arity3(func(G__37 *cljs_core.AFn, ks interface{}, vec__33 cljs_core.CljsCoreIVector, r interface{}, s interface{}, idx interface{}) *cljs_core.AFn {
						return cljs_core.Fn(G__37, 2, func(ret interface{}, x interface{}) interface{} {
							{
								var found = func() interface{} {
									var G__34 = cljs_core.Select_keys.X_invoke_Arity2(x, ks)
									_ = G__34
									return idx.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__34)
								}()
								_ = found
								if cljs_core.Truth_(found) {
									return cljs_core.Reduce.X_invoke_Arity3(func(G__38 *cljs_core.AFn, found interface{}, ks interface{}, vec__33 cljs_core.CljsCoreIVector, r interface{}, s interface{}, idx interface{}) *cljs_core.AFn {
										return cljs_core.Fn(G__38, 2, func(p1__25_SHARP_ interface{}, p2__26_SHARP_ interface{}) interface{} {
											return cljs_core.Conj.X_invoke_Arity2(p1__25_SHARP_, cljs_core.Merge.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{p2__26_SHARP_, x})))
										})
									}(&cljs_core.AFn{}, found, ks, vec__33, r, s, idx), ret, found)
								} else {
									return ret
								}
							}
						})
					}(&cljs_core.AFn{}, ks, vec__33, r, s, idx), cljs_core.CljsCorePersistentHashSet_EMPTY, s)
				}
			} else {
				return cljs_core.CljsCorePersistentHashSet_EMPTY
			}
		}, func(xrel interface{}, yrel interface{}, km interface{}) interface{} {
			{
				var vec__35 = func() cljs_core.CljsCoreIVector {
					if cljs_core.Count.X_invoke_Arity1(xrel).(float64) <= cljs_core.Count.X_invoke_Arity1(yrel).(float64) {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{xrel, yrel, Map_invert.X_invoke_Arity1(km)}, nil})
					} else {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{yrel, xrel, km}, nil})
					}
				}()
				var r = cljs_core.Nth.X_invoke_Arity3(vec__35, float64(0), nil)
				var s = cljs_core.Nth.X_invoke_Arity3(vec__35, float64(1), nil)
				var k = cljs_core.Nth.X_invoke_Arity3(vec__35, float64(2), nil)
				var idx = Index.X_invoke_Arity2(r, cljs_core.Vals.X_invoke_Arity1(k))
				_, _, _, _, _ = vec__35, r, s, k, idx
				return cljs_core.Reduce.X_invoke_Arity3(func(G__39 *cljs_core.AFn, vec__35 cljs_core.CljsCoreIVector, r interface{}, s interface{}, k interface{}, idx interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__39, 2, func(ret interface{}, x interface{}) interface{} {
						{
							var found = func() interface{} {
								var G__36 = Rename_keys.X_invoke_Arity2(cljs_core.Select_keys.X_invoke_Arity2(x, cljs_core.Keys.X_invoke_Arity1(k)), k)
								_ = G__36
								return idx.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__36)
							}()
							_ = found
							if cljs_core.Truth_(found) {
								return cljs_core.Reduce.X_invoke_Arity3(func(G__40 *cljs_core.AFn, found interface{}, vec__35 cljs_core.CljsCoreIVector, r interface{}, s interface{}, k interface{}, idx interface{}) *cljs_core.AFn {
									return cljs_core.Fn(G__40, 2, func(p1__27_SHARP_ interface{}, p2__28_SHARP_ interface{}) interface{} {
										return cljs_core.Conj.X_invoke_Arity2(p1__27_SHARP_, cljs_core.Merge.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{p2__28_SHARP_, x})))
									})
								}(&cljs_core.AFn{}, found, vec__35, r, s, k, idx), ret, found)
							} else {
								return ret
							}
						}
					})
				}(&cljs_core.AFn{}, vec__35, r, s, k, idx), cljs_core.CljsCorePersistentHashSet_EMPTY, s)
			}
		})
	}(&cljs_core.AFn{})

	Subset_QMARK_ = func(subset_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(subset_QMARK_, 2, func(set1 interface{}, set2 interface{}) interface{} {
			return (cljs_core.Count.X_invoke_Arity1(set1).(float64) <= cljs_core.Count.X_invoke_Arity1(set2).(float64)) && (cljs_core.Every_QMARK_.Arity2IIB(func(G__42 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__42, 1, func(p1__41_SHARP_ interface{}) interface{} {
					return cljs_core.Contains_QMARK_.Arity2IIB(set2, p1__41_SHARP_)
				})
			}(&cljs_core.AFn{}), set1))
		})
	}(&cljs_core.AFn{})

	Superset_QMARK_ = func(superset_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(superset_QMARK_, 2, func(set1 interface{}, set2 interface{}) interface{} {
			return (cljs_core.Count.X_invoke_Arity1(set1).(float64) >= cljs_core.Count.X_invoke_Arity1(set2).(float64)) && (cljs_core.Every_QMARK_.Arity2IIB(func(G__44 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__44, 1, func(p1__43_SHARP_ interface{}) interface{} {
					return cljs_core.Contains_QMARK_.Arity2IIB(set1, p1__43_SHARP_)
				})
			}(&cljs_core.AFn{}), set2))
		})
	}(&cljs_core.AFn{})

}

var Bubble_max_key *cljs_core.AFn

// Return a set that is the union of the input sets
// @param {...*} var_args
var Union *cljs_core.AFn

// Return a set that is the intersection of the input sets
// @param {...*} var_args
var Intersection *cljs_core.AFn

// Return a set that is the first set without elements of the remaining sets
// @param {...*} var_args
var Difference *cljs_core.AFn

// Returns a set of the elements for which pred is true
var Select_ *cljs_core.AFn

// Returns a rel of the elements of xrel with only the keys in ks
var Project *cljs_core.AFn

// Returns the map with the keys in kmap renamed to the vals in kmap
var Rename_keys *cljs_core.AFn

// Returns a rel of the maps in xrel with the keys in kmap renamed to the vals in kmap
var Rename *cljs_core.AFn

// Returns a map of the distinct values of ks in the xrel mapped to a
// set of the maps in xrel with the corresponding values of ks.
var Index *cljs_core.AFn

// Returns the map with the vals mapped to the keys.
var Map_invert *cljs_core.AFn

// When passed 2 rels, returns the rel corresponding to the natural
// join. When passed an additional keymap, joins on the corresponding
// keys.
var Join *cljs_core.AFn

// Is set1 a subset of set2?
var Subset_QMARK_ *cljs_core.AFn

// Is set1 a superset of set2?
var Superset_QMARK_ *cljs_core.AFn
