// Compiled by ClojureScript to Go 0.0-2356
// cljs.ns-test

package ns_test

import (
	cljs_core "github.com/hraberg/cljs.go/cljs/core"
	"github.com/hraberg/cljs.go/js"

	clojure_set "github.com/hraberg/cljs.go/clojure/set"
)

func init() {
	X_PLUS_ = cljs_core.X___

	Test_ns = func(test_ns *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test_ns, 0, func() interface{} {
			if cljs_core.X_EQ_.Arity2IIB(float64(4), ((float64(2) + float64(1)) + float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 4 (clojure.core/+ 2 1 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), func() interface{} {
				var G__4061 = float64(2)
				var G__4062 = float64(1)
				var G__4063 = float64(1)
				_, _, _ = G__4061, G__4062, G__4063
				return X_PLUS_.(cljs_core.CljsCoreIFn).X_invoke_Arity3(G__4061, G__4062, G__4063)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (cljs.ns-test/+ 2 1 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), func() interface{} {
				var G__4064 = float64(2)
				var G__4065 = float64(1)
				var G__4066 = float64(1)
				_, _, _ = G__4064, G__4065, G__4066
				return X_PLUS_.(cljs_core.CljsCoreIFn).X_invoke_Arity3(G__4064, G__4065, G__4066)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (+ 2 1 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(123), Baz.X_invoke_Arity0().(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= 123 (baz))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(123), Quux.X_invoke_Arity0().(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= 123 (quux))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity1(float64(5)).(*cljs_core.CljsCoreRange), func() *cljs_core.CljsCoreLazySeq {
				var iter__4763__auto__ = func(iter__4067 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(iter__4067, 1, func(s__4068 interface{}) interface{} {
						return (&cljs_core.CljsCoreLazySeq{nil, func(G__4073 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__4073, 0, func() interface{} {
								{
									var s__4068___1 interface{} = s__4068
									_ = s__4068___1
									for {
										{
											var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4068___1)
											_ = temp__4222__auto__
											if cljs_core.Truth_(temp__4222__auto__) {
												{
													var s__4068___2 = temp__4222__auto__
													_ = s__4068___2
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4068___2) {
														{
															var c__4761__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4068___2)
															var size__4762__auto__ = float64(cljs_core.Int32_(cljs_core.Count.X_invoke_Arity1(c__4761__auto__).(float64)))
															var b__4070 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__4762__auto__).(*cljs_core.CljsCoreChunkBuffer)
															_, _, _ = c__4761__auto__, size__4762__auto__, b__4070
															if func() bool {
																var i__4069 = float64(cljs_core.Int32_(float64(0)))
																_ = i__4069
																for {
																	if i__4069 < size__4762__auto__ {
																		{
																			var x = cljs_core.Native_invoke_instance_method.X_invoke_Arity3(c__4761__auto__, "Nth", []interface{}{i__4069})
																			_ = x
																			cljs_core.Chunk_append.X_invoke_Arity2(b__4070, x)
																			i__4069 = (i__4069 + float64(1))
																			continue
																		}
																	} else {
																		return true
																	}
																}
															}() {
																return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4070), iter__4067.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4068___2)).(*cljs_core.CljsCoreLazySeq))
															} else {
																return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4070), nil)
															}
														}
													} else {
														{
															var x = cljs_core.First.X_invoke_Arity1(s__4068___2)
															_ = x
															return cljs_core.Cons.X_invoke_Arity2(x, iter__4067.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4068___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
														}
													}
												}
											} else {
												return nil
											}
										}
									}
								}
							})
						}(&cljs_core.AFn{}), nil, nil})
					})
				}(&cljs_core.AFn{})
				_ = iter__4763__auto__
				return iter__4763__auto__.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(5)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 5) (lang/for [x (range 5)] x))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), clojure_set.Union.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(3), nil, float64(2), nil}, nil}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 3 2} (s/union #{1} #{3 2}))")}))
			}
			return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
		})
	}(&cljs_core.AFn{})

}

var X_PLUS_ float64

var Test_ns *cljs_core.AFn
