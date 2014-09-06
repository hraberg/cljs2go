// Compiled by ClojureScript to Go 0.0-2322
// clojure.set

// Set operations such as union/intersection.
// Author: Rich Hickey
package set

import (
	"reflect"

	cljs_core "github.com/hraberg/cljs.go/cljs/core"
)

var Bubble_max_key *cljs_core.AFn

func init() {
	Bubble_max_key = func(bubble_max_key *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(bubble_max_key, func(k interface{}, coll interface{}) interface{} {
			{
				var max = cljs_core.Apply.X_invoke_Arity3(cljs_core.Max_key, k, coll)
				_ = max
				return cljs_core.Cons.X_invoke_Arity2(max, cljs_core.Remove.X_invoke_Arity2(func(max interface{}) *cljs_core.AFn {
					return func(G__2 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__2, func(p1__1_SHARP_ interface{}) interface{} {
							return reflect.DeepEqual(max, p1__1_SHARP_)
						})
					}(&cljs_core.AFn{})
				}(max), coll).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
			}
		})
	}(&cljs_core.AFn{})
}

// Return a set that is the union of the input sets
// @param {...*} var_args
var Union *cljs_core.AFn

func init() {
	Union = func(union *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(union, func() interface{} {
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
			var sets = cljs_core.Array_seq.X_invoke_Arity1(s1_s2_sets__[2:])
			_, _, _ = s1, s2, sets
			{
				var bubbled_sets = Bubble_max_key.X_invoke_Arity2(cljs_core.Count, cljs_core.Conj.X_invoke_ArityVariadic(sets, s2, s1)).(*cljs_core.CljsCoreCons)
				_ = bubbled_sets
				return cljs_core.Reduce.X_invoke_Arity3(cljs_core.Into, cljs_core.First.X_invoke_Arity1(bubbled_sets), cljs_core.Rest.Arity1IQ(bubbled_sets))
			}
		})
	}(&cljs_core.AFn{})
}

// Return a set that is the intersection of the input sets
// @param {...*} var_args
var Intersection *cljs_core.AFn

func init() {
	Intersection = func(intersection *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(intersection, func(s1 interface{}) interface{} {
			return s1
		}, func(s1 interface{}, s2 interface{}) interface{} {
			for {
				if cljs_core.Count.X_invoke_Arity1(s2).(float64) < cljs_core.Count.X_invoke_Arity1(s1).(float64) {
					s1, s2 = s2, s1
					continue
				} else {
					return cljs_core.Reduce.X_invoke_Arity3(func(s1 interface{}, s2 interface{}) *cljs_core.AFn {
						return func(G__4 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__4, func(result interface{}, item interface{}) interface{} {
								if cljs_core.Contains_QMARK_.Arity2IIB(s2, item) {
									return result
								} else {
									return cljs_core.Disj.X_invoke_Arity2(result, item)
								}
							})
						}(&cljs_core.AFn{})
					}(s1, s2), s1, s1)
				}
			}
		}, func(s1_s2_sets__ ...interface{}) interface{} {
			var s1 = s1_s2_sets__[0]
			var s2 = s1_s2_sets__[1]
			var sets = cljs_core.Array_seq.X_invoke_Arity1(s1_s2_sets__[2:])
			_, _, _ = s1, s2, sets
			{
				var bubbled_sets = Bubble_max_key.X_invoke_Arity2(func(G__5 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5, func(p1__3_SHARP_ interface{}) interface{} {
						return (-cljs_core.Count.X_invoke_Arity1(p1__3_SHARP_).(float64))
					})
				}(&cljs_core.AFn{}), cljs_core.Conj.X_invoke_ArityVariadic(sets, s2, s1)).(*cljs_core.CljsCoreCons)
				_ = bubbled_sets
				return cljs_core.Reduce.X_invoke_Arity3(intersection, cljs_core.First.X_invoke_Arity1(bubbled_sets), cljs_core.Rest.Arity1IQ(bubbled_sets))
			}
		})
	}(&cljs_core.AFn{})
}

// Return a set that is the first set without elements of the remaining sets
// @param {...*} var_args
var Difference *cljs_core.AFn

func init() {
	Difference = func(difference *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(difference, func(s1 interface{}) interface{} {
			return s1
		}, func(s1 interface{}, s2 interface{}) interface{} {
			if cljs_core.Count.X_invoke_Arity1(s1).(float64) < cljs_core.Count.X_invoke_Arity1(s2).(float64) {
				return cljs_core.Reduce.X_invoke_Arity3(func(G__6 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__6, func(result interface{}, item interface{}) interface{} {
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
			var sets = cljs_core.Array_seq.X_invoke_Arity1(s1_s2_sets__[2:])
			_, _, _ = s1, s2, sets
			return cljs_core.Reduce.X_invoke_Arity3(difference, s1, cljs_core.Conj.X_invoke_Arity2(sets, s2))
		})
	}(&cljs_core.AFn{})
}

// Returns a set of the elements for which pred is true
var Select_ *cljs_core.AFn

func init() {
	Select_ = func(select_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(select_, func(pred interface{}, xset interface{}) interface{} {
			return cljs_core.Reduce.X_invoke_Arity3(func(G__7 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__7, func(s interface{}, k interface{}) interface{} {
					if cljs_core.Truth_(pred.(cljs_core.CljsCoreIFn).X_invoke_Arity1(k)) {
						return s
					} else {
						return cljs_core.Disj.X_invoke_Arity2(s, k)
					}
				})
			}(&cljs_core.AFn{}), xset, xset)
		})
	}(&cljs_core.AFn{})
}

// Returns a rel of the elements of xrel with only the keys in ks
var Project *cljs_core.AFn

func init() {
	Project = func(project *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(project, func(xrel interface{}, ks interface{}) interface{} {
			return cljs_core.Set.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(func(G__9 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__9, func(p1__8_SHARP_ interface{}) interface{} {
					return cljs_core.Select_keys.X_invoke_Arity2(p1__8_SHARP_, ks).(cljs_core.CljsCoreIMap)
				})
			}(&cljs_core.AFn{}), xrel).(*cljs_core.CljsCoreLazySeq))
		})
	}(&cljs_core.AFn{})
}

// Returns the map with the keys in kmap renamed to the vals in kmap
var Rename_keys *cljs_core.AFn

func init() {
	Rename_keys = func(rename_keys *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(rename_keys, func(map_ interface{}, kmap interface{}) interface{} {
			return cljs_core.Reduce.X_invoke_Arity3(func(G__14 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__14, func(m interface{}, p__12 interface{}) interface{} {
					{
						var vec__13 = p__12
						var old = cljs_core.Nth.X_invoke_Arity3(vec__13, float64(0), nil)
						var new = cljs_core.Nth.X_invoke_Arity3(vec__13, float64(1), nil)
						_, _, _ = vec__13, old, new
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
}

// Returns a rel of the maps in xrel with the keys in kmap renamed to the vals in kmap
var Rename *cljs_core.AFn

func init() {
	Rename = func(rename *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(rename, func(xrel interface{}, kmap interface{}) interface{} {
			return cljs_core.Set.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(func(G__16 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__16, func(p1__15_SHARP_ interface{}) interface{} {
					return Rename_keys.X_invoke_Arity2(p1__15_SHARP_, kmap)
				})
			}(&cljs_core.AFn{}), xrel).(*cljs_core.CljsCoreLazySeq))
		})
	}(&cljs_core.AFn{})
}

// Returns a map of the distinct values of ks in the xrel mapped to a
// set of the maps in xrel with the corresponding values of ks.
var Index *cljs_core.AFn

func init() {
	Index = func(index *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(index, func(xrel interface{}, ks interface{}) interface{} {
			return cljs_core.Reduce.X_invoke_Arity3(func(G__17 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__17, func(m interface{}, x interface{}) interface{} {
					{
						var ik = cljs_core.Select_keys.X_invoke_Arity2(x, ks).(cljs_core.CljsCoreIMap)
						_ = ik
						return cljs_core.Assoc.X_invoke_Arity3(m, ik, cljs_core.Conj.X_invoke_Arity2(cljs_core.Get.X_invoke_Arity3(m, ik, cljs_core.CljsCorePersistentHashSet_EMPTY), x))
					}
				})
			}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentArrayMap_EMPTY, xrel)
		})
	}(&cljs_core.AFn{})
}

// Returns the map with the vals mapped to the keys.
var Map_invert *cljs_core.AFn

func init() {
	Map_invert = func(map_invert *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(map_invert, func(m interface{}) interface{} {
			return cljs_core.Reduce.X_invoke_Arity3(func(G__22 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__22, func(m___1 interface{}, p__20 interface{}) interface{} {
					{
						var vec__21 = p__20
						var k = cljs_core.Nth.X_invoke_Arity3(vec__21, float64(0), nil)
						var v = cljs_core.Nth.X_invoke_Arity3(vec__21, float64(1), nil)
						_, _, _ = vec__21, k, v
						return cljs_core.Assoc.X_invoke_Arity3(m___1, v, k)
					}
				})
			}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentArrayMap_EMPTY, m)
		})
	}(&cljs_core.AFn{})
}

// When passed 2 rels, returns the rel corresponding to the natural
// join. When passed an additional keymap, joins on the corresponding
// keys.
var Join *cljs_core.AFn

func init() {
	Join = func(join *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(join, func(xrel interface{}, yrel interface{}) interface{} {
			if cljs_core.Truth_(func() interface{} {
				var and__172__auto__ = cljs_core.Seq.Arity1IQ(xrel)
				_ = and__172__auto__
				if cljs_core.Truth_(and__172__auto__) {
					return cljs_core.Seq.Arity1IQ(yrel)
				} else {
					return and__172__auto__
				}
			}()) {
				{
					var ks = Intersection.X_invoke_Arity2(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(cljs_core.First.X_invoke_Arity1(xrel))), cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(cljs_core.First.X_invoke_Arity1(yrel))))
					var vec__29 = func() cljs_core.CljsCoreIVector {
						if cljs_core.Count.X_invoke_Arity1(xrel).(float64) <= cljs_core.Count.X_invoke_Arity1(yrel).(float64) {
							return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{xrel, yrel}, nil})
						} else {
							return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{yrel, xrel}, nil})
						}
					}()
					var r = cljs_core.Nth.X_invoke_Arity3(vec__29, float64(0), nil)
					var s = cljs_core.Nth.X_invoke_Arity3(vec__29, float64(1), nil)
					var idx = Index.X_invoke_Arity2(r, ks)
					_, _, _, _, _ = ks, vec__29, r, s, idx
					return cljs_core.Reduce.X_invoke_Arity3(func(ks interface{}, vec__29 cljs_core.CljsCoreIVector, r interface{}, s interface{}, idx interface{}) *cljs_core.AFn {
						return func(G__31 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__31, func(ret interface{}, x interface{}) interface{} {
								{
									var found = idx.(cljs_core.CljsCoreIFn).X_invoke_Arity1(cljs_core.Select_keys.X_invoke_Arity2(x, ks).(cljs_core.CljsCoreIMap))
									_ = found
									if cljs_core.Truth_(found) {
										return cljs_core.Reduce.X_invoke_Arity3(func(found interface{}, ks interface{}, vec__29 cljs_core.CljsCoreIVector, r interface{}, s interface{}, idx interface{}) *cljs_core.AFn {
											return func(G__32 *cljs_core.AFn) *cljs_core.AFn {
												return cljs_core.Fn(G__32, func(p1__23_SHARP_ interface{}, p2__24_SHARP_ interface{}) interface{} {
													return cljs_core.Conj.X_invoke_Arity2(p1__23_SHARP_, cljs_core.Merge.X_invoke_ArityVariadic(p2__24_SHARP_, x))
												})
											}(&cljs_core.AFn{})
										}(found, ks, vec__29, r, s, idx), ret, found)
									} else {
										return ret
									}
								}
							})
						}(&cljs_core.AFn{})
					}(ks, vec__29, r, s, idx), cljs_core.CljsCorePersistentHashSet_EMPTY, s)
				}
			} else {
				return cljs_core.CljsCorePersistentHashSet_EMPTY
			}
		}, func(xrel interface{}, yrel interface{}, km interface{}) interface{} {
			{
				var vec__30 = func() cljs_core.CljsCoreIVector {
					if cljs_core.Count.X_invoke_Arity1(xrel).(float64) <= cljs_core.Count.X_invoke_Arity1(yrel).(float64) {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{xrel, yrel, Map_invert.X_invoke_Arity1(km)}, nil})
					} else {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{yrel, xrel, km}, nil})
					}
				}()
				var r = cljs_core.Nth.X_invoke_Arity3(vec__30, float64(0), nil)
				var s = cljs_core.Nth.X_invoke_Arity3(vec__30, float64(1), nil)
				var k = cljs_core.Nth.X_invoke_Arity3(vec__30, float64(2), nil)
				var idx = Index.X_invoke_Arity2(r, cljs_core.Vals.X_invoke_Arity1(k))
				_, _, _, _, _ = vec__30, r, s, k, idx
				return cljs_core.Reduce.X_invoke_Arity3(func(vec__30 cljs_core.CljsCoreIVector, r interface{}, s interface{}, k interface{}, idx interface{}) *cljs_core.AFn {
					return func(G__33 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__33, func(ret interface{}, x interface{}) interface{} {
							{
								var found = idx.(cljs_core.CljsCoreIFn).X_invoke_Arity1(Rename_keys.X_invoke_Arity2(cljs_core.Select_keys.X_invoke_Arity2(x, cljs_core.Keys.X_invoke_Arity1(k)).(cljs_core.CljsCoreIMap), k))
								_ = found
								if cljs_core.Truth_(found) {
									return cljs_core.Reduce.X_invoke_Arity3(func(found interface{}, vec__30 cljs_core.CljsCoreIVector, r interface{}, s interface{}, k interface{}, idx interface{}) *cljs_core.AFn {
										return func(G__34 *cljs_core.AFn) *cljs_core.AFn {
											return cljs_core.Fn(G__34, func(p1__25_SHARP_ interface{}, p2__26_SHARP_ interface{}) interface{} {
												return cljs_core.Conj.X_invoke_Arity2(p1__25_SHARP_, cljs_core.Merge.X_invoke_ArityVariadic(p2__26_SHARP_, x))
											})
										}(&cljs_core.AFn{})
									}(found, vec__30, r, s, k, idx), ret, found)
								} else {
									return ret
								}
							}
						})
					}(&cljs_core.AFn{})
				}(vec__30, r, s, k, idx), cljs_core.CljsCorePersistentHashSet_EMPTY, s)
			}
		})
	}(&cljs_core.AFn{})
}

// Is set1 a subset of set2?
var Subset_QMARK_ *cljs_core.AFn

func init() {
	Subset_QMARK_ = func(subset_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(subset_QMARK_, func(set1 interface{}, set2 interface{}) interface{} {
			return (cljs_core.Count.X_invoke_Arity1(set1).(float64) <= cljs_core.Count.X_invoke_Arity1(set2).(float64)) && (cljs_core.Every_QMARK_.Arity2IIB(func(G__36 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__36, func(p1__35_SHARP_ interface{}) interface{} {
					return cljs_core.Contains_QMARK_.Arity2IIB(set2, p1__35_SHARP_)
				})
			}(&cljs_core.AFn{}), set1))
		})
	}(&cljs_core.AFn{})
}

// Is set1 a superset of set2?
var Superset_QMARK_ *cljs_core.AFn

func init() {
	Superset_QMARK_ = func(superset_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(superset_QMARK_, func(set1 interface{}, set2 interface{}) interface{} {
			return (cljs_core.Count.X_invoke_Arity1(set1).(float64) >= cljs_core.Count.X_invoke_Arity1(set2).(float64)) && (cljs_core.Every_QMARK_.Arity2IIB(func(G__38 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__38, func(p1__37_SHARP_ interface{}) interface{} {
					return cljs_core.Contains_QMARK_.Arity2IIB(set1, p1__37_SHARP_)
				})
			}(&cljs_core.AFn{}), set2))
		})
	}(&cljs_core.AFn{})
}
