// Compiled by ClojureScript to Go 0.0-2371
// cljs.ns-test

package ns_test

import (
	"testing"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	cljs_ns_test_bar "github.com/hraberg/cljs2go/cljs/ns_test/bar"
	cljs_ns_test_foo "github.com/hraberg/cljs2go/cljs/ns_test/foo"
	clojure_set "github.com/hraberg/cljs2go/clojure/set"
	"github.com/hraberg/cljs2go/js"
	"github.com/stretchr/testify/assert"
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
				var G__13 = float64(2)
				var G__14 = float64(1)
				var G__15 = float64(1)
				_, _, _ = G__13, G__14, G__15
				return X_PLUS_.(cljs_core.CljsCoreIFn).X_invoke_Arity3(G__13, G__14, G__15)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (cljs.ns-test/+ 2 1 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), func() interface{} {
				var G__16 = float64(2)
				var G__17 = float64(1)
				var G__18 = float64(1)
				_, _, _ = G__16, G__17, G__18
				return X_PLUS_.(cljs_core.CljsCoreIFn).X_invoke_Arity3(G__16, G__17, G__18)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (+ 2 1 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(123), cljs_ns_test_foo.Baz.X_invoke_Arity0().(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= 123 (baz))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(123), cljs_ns_test_bar.Quux.X_invoke_Arity0().(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= 123 (quux))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity1(float64(5)).(*cljs_core.CljsCoreRange), func() *cljs_core.CljsCoreLazySeq {
				var iter__4951__auto__ = func(iter__19 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(iter__19, 1, func(s__20 interface{}) interface{} {
						return (&cljs_core.CljsCoreLazySeq{nil, func(G__25 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__25, 0, func() interface{} {
								{
									var s__20___1 interface{} = s__20
									_ = s__20___1
									for {
										{
											var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__20___1)
											_ = temp__4388__auto__
											if cljs_core.Truth_(temp__4388__auto__) {
												{
													var s__20___2 = temp__4388__auto__
													_ = s__20___2
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__20___2) {
														{
															var c__4949__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__20___2)
															var size__4950__auto__ = float64(cljs_core.Int32_(cljs_core.Count.X_invoke_Arity1(c__4949__auto__).(float64)))
															var b__22 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__4950__auto__).(*cljs_core.CljsCoreChunkBuffer)
															_, _, _ = c__4949__auto__, size__4950__auto__, b__22
															if func() bool {
																var i__21 = float64(cljs_core.Int32_(float64(0)))
																_ = i__21
																for {
																	if i__21 < size__4950__auto__ {
																		{
																			var x = cljs_core.Native_invoke_instance_method.X_invoke_Arity3(c__4949__auto__, "Nth", []interface{}{i__21})
																			_ = x
																			cljs_core.Chunk_append.X_invoke_Arity2(b__22, x)
																			i__21 = (i__21 + float64(1))
																			continue
																		}
																	} else {
																		return true
																	}
																}
															}() {
																return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__22), iter__19.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__20___2)).(*cljs_core.CljsCoreLazySeq))
															} else {
																return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__22), nil)
															}
														}
													} else {
														{
															var x = cljs_core.First.X_invoke_Arity1(s__20___2)
															_ = x
															return cljs_core.Cons.X_invoke_Arity2(x, iter__19.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__20___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
				_ = iter__4951__auto__
				return iter__4951__auto__.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(5)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
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

var X_PLUS_ interface{}

var Test_ns *cljs_core.AFn

func Test_runner(t *testing.T) {
	assert.Equal(t, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), Test_ns.X_invoke_Arity0().(*cljs_core.CljsCoreKeyword))
}
