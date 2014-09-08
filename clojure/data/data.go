// Compiled by ClojureScript to Go 0.0-2322
// clojure.data

// Non-core data functions.
// Author: Stuart Halloway
package data

import (
	cljs_core "github.com/hraberg/cljs.go/cljs/core"

	clojure_set "github.com/hraberg/cljs.go/clojure/set"
)

func init() {
	Atom_diff = func(atom_diff *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(atom_diff, func(a interface{}, b interface{}) interface{} {
			if cljs_core.X_EQ_.Arity2IIB(a, b) {
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, nil, a}, nil})
			} else {
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b, nil}, nil})
			}
		})
	}(&cljs_core.AFn{})

	Vectorize = func(vectorize *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(vectorize, func(m interface{}) interface{} {
			if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(m)) {
				return cljs_core.Reduce.X_invoke_Arity3(func(G__5 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5, func(result interface{}, p__3 interface{}) interface{} {
						{
							var vec__4 = p__3
							var k = cljs_core.Nth.X_invoke_Arity3(vec__4, float64(0), nil)
							var v = cljs_core.Nth.X_invoke_Arity3(vec__4, float64(1), nil)
							_, _, _ = vec__4, k, v
							return cljs_core.Assoc.X_invoke_Arity3(result, k, v)
						}
					})
				}(&cljs_core.AFn{}), cljs_core.Vec.X_invoke_Arity1(cljs_core.Repeat.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, cljs_core.Keys.X_invoke_Arity1(m)), nil).(*cljs_core.CljsCoreLazySeq)), m)
			} else {
				return nil
			}
		})
	}(&cljs_core.AFn{})

	Diff_associative_key = func(diff_associative_key *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(diff_associative_key, func(a interface{}, b interface{}, k interface{}) interface{} {
			{
				var va = cljs_core.Get.X_invoke_Arity2(a, k)
				var vb = cljs_core.Get.X_invoke_Arity2(b, k)
				var vec__7 = Diff.X_invoke_Arity2(va, vb)
				var a_STAR_ = cljs_core.Nth.X_invoke_Arity3(vec__7, float64(0), nil)
				var b_STAR_ = cljs_core.Nth.X_invoke_Arity3(vec__7, float64(1), nil)
				var ab = cljs_core.Nth.X_invoke_Arity3(vec__7, float64(2), nil)
				var in_a = cljs_core.Contains_QMARK_.Arity2IIB(a, k)
				var in_b = cljs_core.Contains_QMARK_.Arity2IIB(b, k)
				var same = (in_a) && (in_b) && ((!(cljs_core.Nil_(ab))) || ((cljs_core.Nil_(va)) && (cljs_core.Nil_(vb))))
				_, _, _, _, _, _, _, _, _ = va, vb, vec__7, a_STAR_, b_STAR_, ab, in_a, in_b, same
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
					if (in_a) && ((!(cljs_core.Nil_(a_STAR_))) || (!(same))) {
						return cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{k, a_STAR_}, true, false).(*cljs_core.CljsCorePersistentArrayMap)
					} else {
						return nil
					}
				}(), func() interface{} {
					if (in_b) && ((!(cljs_core.Nil_(b_STAR_))) || (!(same))) {
						return cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{k, b_STAR_}, true, false).(*cljs_core.CljsCorePersistentArrayMap)
					} else {
						return nil
					}
				}(), func() interface{} {
					if cljs_core.Truth_(same) {
						return cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{k, ab}, true, false).(*cljs_core.CljsCorePersistentArrayMap)
					} else {
						return nil
					}
				}()}, nil})
			}
		})
	}(&cljs_core.AFn{})

	Diff_associative = func(diff_associative *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(diff_associative, func(a interface{}, b interface{}) interface{} {
			return diff_associative.X_invoke_Arity3(a, b, clojure_set.Union.X_invoke_Arity2(cljs_core.Keys.X_invoke_Arity1(a), cljs_core.Keys.X_invoke_Arity1(b)))
		}, func(a interface{}, b interface{}, ks interface{}) interface{} {
			return cljs_core.Reduce.X_invoke_Arity3(func(G__8 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__8, func(diff1 interface{}, diff2 interface{}) interface{} {
					return cljs_core.Doall.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity3(cljs_core.Merge, diff1, diff2).(*cljs_core.CljsCoreLazySeq))
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, nil, nil}, nil}), cljs_core.Map_.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity3(Diff_associative_key, a, b).(cljs_core.CljsCoreIFn), ks).(*cljs_core.CljsCoreLazySeq))
		})
	}(&cljs_core.AFn{})

	Diff_sequential = func(diff_sequential *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(diff_sequential, func(a interface{}, b interface{}) interface{} {
			return cljs_core.Vec.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(Vectorize, Diff_associative.X_invoke_Arity3(func() interface{} {
				if cljs_core.Vector_QMARK_.Arity1IB(a) {
					return a
				} else {
					return cljs_core.Vec.X_invoke_Arity1(a)
				}
			}(), func() interface{} {
				if cljs_core.Vector_QMARK_.Arity1IB(b) {
					return b
				} else {
					return cljs_core.Vec.X_invoke_Arity1(b)
				}
			}(), cljs_core.Range_.X_invoke_Arity1(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(cljs_core.Count.X_invoke_Arity1(a).(float64), cljs_core.Count.X_invoke_Arity1(b).(float64))).(*cljs_core.CljsCoreRange))).(*cljs_core.CljsCoreLazySeq))
		})
	}(&cljs_core.AFn{})

	Diff_set = func(diff_set *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(diff_set, func(a interface{}, b interface{}) interface{} {
			return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Not_empty.X_invoke_Arity1(clojure_set.Difference.X_invoke_Arity2(a, b)), cljs_core.Not_empty.X_invoke_Arity1(clojure_set.Difference.X_invoke_Arity2(b, a)), cljs_core.Not_empty.X_invoke_Arity1(clojure_set.Intersection.X_invoke_Arity2(a, b))}, nil})
		})
	}(&cljs_core.AFn{})

	Equality_partition = func(equality_partition *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(equality_partition, func(x interface{}) interface{} {
			return x.(ClojureDataEqualityPartition).Equality_partition_Arity1()
		})
	}(&cljs_core.AFn{})

	Diff_similar = func(diff_similar *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(diff_similar, func(a interface{}, b interface{}) interface{} {
			return a.(ClojureDataDiff).Diff_similar_Arity2(b)
		})
	}(&cljs_core.AFn{})

	Diff = func(diff *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(diff, func(a interface{}, b interface{}) interface{} {
			if cljs_core.X_EQ_.Arity2IIB(a, b) {
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, nil, a}, nil})
			} else {
				if cljs_core.X_EQ_.Arity2IIB(a.(ClojureDataEqualityPartition).Equality_partition_Arity1(), b.(ClojureDataEqualityPartition).Equality_partition_Arity1()) {
					return a.(ClojureDataDiff).Diff_similar_Arity2(b)
				} else {
					return Atom_diff.X_invoke_Arity2(a, b).(cljs_core.CljsCoreIVector)
				}
			}
		})
	}(&cljs_core.AFn{})

}

// Internal helper for diff.
var Atom_diff *cljs_core.AFn

// Convert an associative-by-numeric-index collection into
// an equivalent vector, with nil for any missing keys
var Vectorize *cljs_core.AFn

// Diff associative things a and b, comparing only the key k.
var Diff_associative_key *cljs_core.AFn

// Diff associative things a and b, comparing only keys in ks (if supplied).
var Diff_associative *cljs_core.AFn

var Diff_sequential *cljs_core.AFn

var Diff_set *cljs_core.AFn

var Equality_partition *cljs_core.AFn

var Diff_similar *cljs_core.AFn

// Recursively compares a and b, returning a tuple of
// [things-only-in-a things-only-in-b things-in-both].
// Comparison rules:
//
// * For equal a and b, return [nil nil a].
// * Maps are subdiffed where keys match and values differ.
// * Sets are never subdiffed.
// * All sequential things are treated as associative collections
// by their indexes, with results returned as vectors.
// * Everything else (including strings!) is treated as
// an atom and compared for equality.
var Diff *cljs_core.AFn

type ClojureDataEqualityPartition interface {
	ClojureDataEqualityPartition__()
	Equality_partition_Arity1() interface{}
}

type ClojureDataDiff interface {
	ClojureDataDiff__()
	Diff_similar_Arity2(b interface{}) interface{}
}

func init() {
	cljs_core.RegisterProtocol_("clojure.data/EqualityPartition", (*ClojureDataEqualityPartition)(nil))
}

func init() {
	cljs_core.RegisterProtocol_("clojure.data/Diff", (*ClojureDataDiff)(nil))
}
