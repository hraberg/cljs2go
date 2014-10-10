// Compiled by ClojureScript to Go 0.0-2371
// cljs.reader-test

package reader_test

import (
	"reflect"
	"testing"

	cljs_core "github.com/hraberg/cljs.go/cljs/core"
	cljs_reader "github.com/hraberg/cljs.go/cljs/reader"
	"github.com/hraberg/cljs.go/js"
	"github.com/stretchr/testify/assert"
)

func init() {
	X__GT_T = func(__GT_T *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(__GT_T, 2, func(a interface{}, b interface{}) interface{} {
			return (&CljsReader_testT{a, b})
		})
	}(&cljs_core.AFn{})

	{
		X__GT_R = func(__GT_R *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(__GT_R, 2, func(a interface{}, b interface{}) interface{} {
				return (&CljsReader_testR{a, b, nil, nil, nil})
			})
		}(&cljs_core.AFn{})

		Map__GT_R = func(map__GT_R *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(map__GT_R, 1, func(G__163 interface{}) interface{} {
				return (&CljsReader_testR{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__163), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__163), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__163, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})})), nil})
			})
		}(&cljs_core.AFn{})

	}
	Test_reader = func(test_reader *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test_reader, 0, func() interface{} {
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_reader.Read_string.X_invoke_Arity1("1")) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (reader/read-string \"1\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_reader.Read_string.X_invoke_Arity1("#_nope 2")) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (reader/read-string \"#_nope 2\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_reader.Read_string.X_invoke_Arity1("-1")) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (reader/read-string \"-1\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(-1.5, cljs_reader.Read_string.X_invoke_Arity1("-1.5")) {
			} else {
				panic((&js.Error{("Assert failed: (= -1.5 (reader/read-string \"-1.5\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4)}, nil}), cljs_reader.Read_string.X_invoke_Arity1("[3 4]")) {
			} else {
				panic((&js.Error{("Assert failed: (= [3 4] (reader/read-string \"[3 4]\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("foo", cljs_reader.Read_string.X_invoke_Arity1("\"foo\"")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"foo\" (reader/read-string \"\\\"foo\\\"\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hello", Fqn: "hello", X_hash: float64(-245025397)}), cljs_reader.Read_string.X_invoke_Arity1(":hello")) {
			} else {
				panic((&js.Error{("Assert failed: (= :hello (reader/read-string \":hello\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "goodbye", Str: "goodbye", X_hash: float64(-1229580373), X_meta: nil}), cljs_reader.Read_string.X_invoke_Arity1("goodbye")) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote goodbye) (reader/read-string \"goodbye\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "%", Str: "%", X_hash: float64(-950237169), X_meta: nil}), cljs_reader.Read_string.X_invoke_Arity1("%")) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote %) (reader/read-string \"%\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), cljs_reader.Read_string.X_invoke_Arity1("#{1 2 3}")) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 3 2} (reader/read-string \"#{1 2 3}\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList), cljs_reader.Read_string.X_invoke_Arity1("(7 8 9)")) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (7 8 9)) (reader/read-string \"(7 8 9)\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "deref", Str: "deref", X_hash: float64(1494944732), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})})).(*cljs_core.CljsCoreList), cljs_reader.Read_string.X_invoke_Arity1("@foo")) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (deref foo)) (reader/read-string \"@foo\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "quote", Str: "quote", X_hash: float64(1377916282), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})})).(*cljs_core.CljsCoreList), cljs_reader.Read_string.X_invoke_Arity1("'bar")) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (quote bar)) (reader/read-string \"'bar\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: "foo", Name: "bar", Str: "foo/bar", X_hash: float64(254379989), X_meta: nil}), cljs_reader.Read_string.X_invoke_Arity1("foo/bar")) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote foo/bar) (reader/read-string \"foo/bar\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("a", cljs_reader.Read_string.X_invoke_Arity1("\\a")) {
			} else {
				panic((&js.Error{("Assert failed: (= \\a (reader/read-string \"\\\\a\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "tag", Fqn: "tag", X_hash: float64(-1290361223)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "String", Str: "String", X_hash: float64(-2070057435), X_meta: nil})}, nil}), cljs_core.Meta.X_invoke_Arity1(cljs_reader.Read_string.X_invoke_Arity1("^String {:a 1}"))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:tag (quote String)} (meta (reader/read-string \"^String {:a 1}\")))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)})}, nil})}, nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil})}, true).(*cljs_core.CljsCorePersistentHashSet)}, nil}), cljs_reader.Read_string.X_invoke_Arity1("[:a b #{c {:d [:e :f :g]}}]")) {
			} else {
				panic((&js.Error{("Assert failed: (= [:a (quote b) #{{:d [:e :f :g]} (quote c)}] (reader/read-string \"[:a b #{c {:d [:e :f :g]}}]\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_reader.Read_string.X_invoke_Arity1("nil")) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (reader/read-string \"nil\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(true, cljs_reader.Read_string.X_invoke_Arity1("true")) {
			} else {
				panic((&js.Error{("Assert failed: (= true (reader/read-string \"true\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(false, cljs_reader.Read_string.X_invoke_Arity1("false")) {
			} else {
				panic((&js.Error{("Assert failed: (= false (reader/read-string \"false\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("string", cljs_reader.Read_string.X_invoke_Arity1("\"string\"")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"string\" (reader/read-string \"\\\"string\\\"\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("escape chars \t \r \n \\ \" \b \f", cljs_reader.Read_string.X_invoke_Arity1("\"escape chars \\t \\r \\n \\\\ \\\" \\b \\f\"")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"escape chars \\t \\r \\n \\\\ \\\" \\b \\f\" (reader/read-string \"\\\"escape chars \\\\t \\\\r \\\\n \\\\\\\\ \\\\\\\" \\\\b \\\\f\\\"\"))")}))
			}
			if cljs_core.Truth_(cljs_core.Apply.X_invoke_Arity3(cljs_core.X_EQ_, float64(0), cljs_core.Map_.X_invoke_Arity2(cljs_reader.Read_string, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"0", "+0", "-0", " 0 "}, nil})).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{("Assert failed: (apply = 0 (map reader/read-string [\"0\" \"+0\" \"-0\" \" 0 \"]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentQueue_EMPTY, cljs_reader.Read_string.X_invoke_Arity1("#queue []")) {
			} else {
				panic((&js.Error{("Assert failed: (= (.-EMPTY cljs.core/PersistentQueue) (reader/read-string \"#queue []\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, float64(1)), cljs_reader.Read_string.X_invoke_Arity1("#queue [1]")) {
			} else {
				panic((&js.Error{("Assert failed: (= (-> (.-EMPTY cljs.core/PersistentQueue) (conj 1)) (reader/read-string \"#queue [1]\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), cljs_reader.Read_string.X_invoke_Arity1("#queue [1 2]")) {
			} else {
				panic((&js.Error{("Assert failed: (= (into (.-EMPTY cljs.core/PersistentQueue) [1 2]) (reader/read-string \"#queue [1 2]\"))")}))
			}
			if cljs_core.Nil_(cljs_reader.Read_string.X_invoke_Arity1(";foo")) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (reader/read-string \";foo\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__253 interface{}) {
				defer func() {
					if e217 := recover(); e217 != nil {
						if cljs_core.Value_(e217).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
							{
								var e = e217
								_ = e
								return__253 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "threw", Fqn: "threw", X_hash: float64(17630075)})
							}
						} else {
							panic(e217)

						}
					}
				}()
				{
					return cljs_reader.Read_string.X_invoke_Arity1(";foo\n3")
				}
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 3 (try (reader/read-string \";foo\\n3\") (catch js/Error e :threw)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__254 interface{}) {
				defer func() {
					if e218 := recover(); e218 != nil {
						if cljs_core.Value_(e218).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
							{
								var e = e218
								_ = e
								return__254 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "threw", Fqn: "threw", X_hash: float64(17630075)})
							}
						} else {
							panic(e218)

						}
					}
				}()
				{
					return cljs_reader.Read_string.X_invoke_Arity1(";foo\n3\n5")
				}
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 3 (try (reader/read-string \";foo\\n3\\n5\") (catch js/Error e :threw)))")}))
			}
			{
				var est_inst_255 = cljs_reader.Read_string.X_invoke_Arity1("#inst \"2010-11-12T13:14:15.666-05:00\"")
				var utc_inst_256 = cljs_reader.Read_string.X_invoke_Arity1("#inst \"2010-11-12T18:14:15.666-00:00\"")
				var pad_257 = func(G__258 *cljs_core.AFn, est_inst_255 interface{}, utc_inst_256 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__258, 1, func(n interface{}) interface{} {
						if n.(float64) < float64(10) {
							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
						} else {
							return n
						}
					})
				}(&cljs_core.AFn{}, est_inst_255, utc_inst_256)
				_, _, _ = est_inst_255, utc_inst_256, pad_257
				if cljs_core.X_EQ_.Arity2IIB((&js.Date{"2010-11-12T13:14:15.666-05:00"}).ValueOf(), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_255, "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf (js/Date. \"2010-11-12T13:14:15.666-05:00\")) (.valueOf est-inst))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_255, "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{est_inst_255})).(string)), "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf est-inst) (.valueOf (reader/read-string (pr-str est-inst))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_255, "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(utc_inst_256, "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf est-inst) (.valueOf utc-inst))")}))
				}
				{
					var seq__219_259 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(13)).(*cljs_core.CljsCoreRange))
					var chunk__232_260 interface{} = nil
					var count__233_261 = float64(0)
					var i__234_262 = float64(0)
					_, _, _, _ = seq__219_259, chunk__232_260, count__233_261, i__234_262
					for {
						if i__234_262 < count__233_261 {
							{
								var month_263 = cljs_core.Decorate_(chunk__232_260).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__234_262)
								_ = month_263
								{
									var seq__235_264 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
									var chunk__240_265 interface{} = nil
									var count__241_266 = float64(0)
									var i__242_267 = float64(0)
									_, _, _, _ = seq__235_264, chunk__240_265, count__241_266, i__242_267
									for {
										if i__242_267 < count__241_266 {
											{
												var day_268 = cljs_core.Decorate_(chunk__240_265).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__242_267)
												_ = day_268
												{
													var seq__243_269 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
													var chunk__244_270 interface{} = nil
													var count__245_271 = float64(0)
													var i__246_272 = float64(0)
													_, _, _, _ = seq__243_269, chunk__244_270, count__245_271, i__246_272
													for {
														if i__246_272 < count__245_271 {
															{
																var hour_273 = cljs_core.Decorate_(chunk__244_270).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__246_272)
																_ = hour_273
																{
																	var s_274 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(month_263)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(day_268)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(hour_273)).(string) + ":14:15.666-06:00\"")
																	_ = s_274
																	if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_274), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_274)})).(string)), "ValueOf", []interface{}{})) {
																	} else {
																		panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																	}
																}
																seq__243_269, chunk__244_270, count__245_271, i__246_272 = seq__243_269, chunk__244_270, count__245_271, (i__246_272 + float64(1))
																continue
															}
														} else {
															{
																var temp__4222__auto___275 = cljs_core.Seq.Arity1IQ(seq__243_269)
																_ = temp__4222__auto___275
																if cljs_core.Truth_(temp__4222__auto___275) {
																	{
																		var seq__243_276___1 = temp__4222__auto___275
																		_ = seq__243_276___1
																		if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__243_276___1) {
																			{
																				var c__966__auto___277 = cljs_core.Chunk_first.X_invoke_Arity1(seq__243_276___1)
																				_ = c__966__auto___277
																				seq__243_269, chunk__244_270, count__245_271, i__246_272 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__243_276___1), c__966__auto___277, cljs_core.Count.X_invoke_Arity1(c__966__auto___277).(float64), float64(0)
																				continue
																			}
																		} else {
																			{
																				var hour_278 = cljs_core.First.X_invoke_Arity1(seq__243_276___1)
																				_ = hour_278
																				{
																					var s_279 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(month_263)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(day_268)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(hour_278)).(string) + ":14:15.666-06:00\"")
																					_ = s_279
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_279), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_279)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__243_269, chunk__244_270, count__245_271, i__246_272 = cljs_core.Next.Arity1IQ(seq__243_276___1), nil, float64(0), float64(0)
																				continue
																			}
																		}
																	}
																} else {
																}
															}
														}
														break
													}
												}
												seq__235_264, chunk__240_265, count__241_266, i__242_267 = seq__235_264, chunk__240_265, count__241_266, (i__242_267 + float64(1))
												continue
											}
										} else {
											{
												var temp__4222__auto___280 = cljs_core.Seq.Arity1IQ(seq__235_264)
												_ = temp__4222__auto___280
												if cljs_core.Truth_(temp__4222__auto___280) {
													{
														var seq__235_281___1 = temp__4222__auto___280
														_ = seq__235_281___1
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__235_281___1) {
															{
																var c__966__auto___282 = cljs_core.Chunk_first.X_invoke_Arity1(seq__235_281___1)
																_ = c__966__auto___282
																seq__235_264, chunk__240_265, count__241_266, i__242_267 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__235_281___1), c__966__auto___282, cljs_core.Count.X_invoke_Arity1(c__966__auto___282).(float64), float64(0)
																continue
															}
														} else {
															{
																var day_283 = cljs_core.First.X_invoke_Arity1(seq__235_281___1)
																_ = day_283
																{
																	var seq__236_284 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																	var chunk__237_285 interface{} = nil
																	var count__238_286 = float64(0)
																	var i__239_287 = float64(0)
																	_, _, _, _ = seq__236_284, chunk__237_285, count__238_286, i__239_287
																	for {
																		if i__239_287 < count__238_286 {
																			{
																				var hour_288 = cljs_core.Decorate_(chunk__237_285).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__239_287)
																				_ = hour_288
																				{
																					var s_289 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(month_263)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(day_283)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(hour_288)).(string) + ":14:15.666-06:00\"")
																					_ = s_289
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_289), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_289)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__236_284, chunk__237_285, count__238_286, i__239_287 = seq__236_284, chunk__237_285, count__238_286, (i__239_287 + float64(1))
																				continue
																			}
																		} else {
																			{
																				var temp__4222__auto___290___1 = cljs_core.Seq.Arity1IQ(seq__236_284)
																				_ = temp__4222__auto___290___1
																				if cljs_core.Truth_(temp__4222__auto___290___1) {
																					{
																						var seq__236_291___1 = temp__4222__auto___290___1
																						_ = seq__236_291___1
																						if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__236_291___1) {
																							{
																								var c__966__auto___292 = cljs_core.Chunk_first.X_invoke_Arity1(seq__236_291___1)
																								_ = c__966__auto___292
																								seq__236_284, chunk__237_285, count__238_286, i__239_287 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__236_291___1), c__966__auto___292, cljs_core.Count.X_invoke_Arity1(c__966__auto___292).(float64), float64(0)
																								continue
																							}
																						} else {
																							{
																								var hour_293 = cljs_core.First.X_invoke_Arity1(seq__236_291___1)
																								_ = hour_293
																								{
																									var s_294 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(month_263)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(day_283)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(hour_293)).(string) + ":14:15.666-06:00\"")
																									_ = s_294
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_294), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_294)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__236_284, chunk__237_285, count__238_286, i__239_287 = cljs_core.Next.Arity1IQ(seq__236_291___1), nil, float64(0), float64(0)
																								continue
																							}
																						}
																					}
																				} else {
																				}
																			}
																		}
																		break
																	}
																}
																seq__235_264, chunk__240_265, count__241_266, i__242_267 = cljs_core.Next.Arity1IQ(seq__235_281___1), nil, float64(0), float64(0)
																continue
															}
														}
													}
												} else {
												}
											}
										}
										break
									}
								}
								seq__219_259, chunk__232_260, count__233_261, i__234_262 = seq__219_259, chunk__232_260, count__233_261, (i__234_262 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___295 = cljs_core.Seq.Arity1IQ(seq__219_259)
								_ = temp__4222__auto___295
								if cljs_core.Truth_(temp__4222__auto___295) {
									{
										var seq__219_296___1 = temp__4222__auto___295
										_ = seq__219_296___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__219_296___1) {
											{
												var c__966__auto___297 = cljs_core.Chunk_first.X_invoke_Arity1(seq__219_296___1)
												_ = c__966__auto___297
												seq__219_259, chunk__232_260, count__233_261, i__234_262 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__219_296___1), c__966__auto___297, cljs_core.Count.X_invoke_Arity1(c__966__auto___297).(float64), float64(0)
												continue
											}
										} else {
											{
												var month_298 = cljs_core.First.X_invoke_Arity1(seq__219_296___1)
												_ = month_298
												{
													var seq__220_299 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
													var chunk__225_300 interface{} = nil
													var count__226_301 = float64(0)
													var i__227_302 = float64(0)
													_, _, _, _ = seq__220_299, chunk__225_300, count__226_301, i__227_302
													for {
														if i__227_302 < count__226_301 {
															{
																var day_303 = cljs_core.Decorate_(chunk__225_300).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__227_302)
																_ = day_303
																{
																	var seq__228_304 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																	var chunk__229_305 interface{} = nil
																	var count__230_306 = float64(0)
																	var i__231_307 = float64(0)
																	_, _, _, _ = seq__228_304, chunk__229_305, count__230_306, i__231_307
																	for {
																		if i__231_307 < count__230_306 {
																			{
																				var hour_308 = cljs_core.Decorate_(chunk__229_305).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__231_307)
																				_ = hour_308
																				{
																					var s_309 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(month_298)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(day_303)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(hour_308)).(string) + ":14:15.666-06:00\"")
																					_ = s_309
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_309), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_309)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__228_304, chunk__229_305, count__230_306, i__231_307 = seq__228_304, chunk__229_305, count__230_306, (i__231_307 + float64(1))
																				continue
																			}
																		} else {
																			{
																				var temp__4222__auto___310___1 = cljs_core.Seq.Arity1IQ(seq__228_304)
																				_ = temp__4222__auto___310___1
																				if cljs_core.Truth_(temp__4222__auto___310___1) {
																					{
																						var seq__228_311___1 = temp__4222__auto___310___1
																						_ = seq__228_311___1
																						if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__228_311___1) {
																							{
																								var c__966__auto___312 = cljs_core.Chunk_first.X_invoke_Arity1(seq__228_311___1)
																								_ = c__966__auto___312
																								seq__228_304, chunk__229_305, count__230_306, i__231_307 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__228_311___1), c__966__auto___312, cljs_core.Count.X_invoke_Arity1(c__966__auto___312).(float64), float64(0)
																								continue
																							}
																						} else {
																							{
																								var hour_313 = cljs_core.First.X_invoke_Arity1(seq__228_311___1)
																								_ = hour_313
																								{
																									var s_314 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(month_298)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(day_303)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(hour_313)).(string) + ":14:15.666-06:00\"")
																									_ = s_314
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_314), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_314)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__228_304, chunk__229_305, count__230_306, i__231_307 = cljs_core.Next.Arity1IQ(seq__228_311___1), nil, float64(0), float64(0)
																								continue
																							}
																						}
																					}
																				} else {
																				}
																			}
																		}
																		break
																	}
																}
																seq__220_299, chunk__225_300, count__226_301, i__227_302 = seq__220_299, chunk__225_300, count__226_301, (i__227_302 + float64(1))
																continue
															}
														} else {
															{
																var temp__4222__auto___315___1 = cljs_core.Seq.Arity1IQ(seq__220_299)
																_ = temp__4222__auto___315___1
																if cljs_core.Truth_(temp__4222__auto___315___1) {
																	{
																		var seq__220_316___1 = temp__4222__auto___315___1
																		_ = seq__220_316___1
																		if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__220_316___1) {
																			{
																				var c__966__auto___317 = cljs_core.Chunk_first.X_invoke_Arity1(seq__220_316___1)
																				_ = c__966__auto___317
																				seq__220_299, chunk__225_300, count__226_301, i__227_302 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__220_316___1), c__966__auto___317, cljs_core.Count.X_invoke_Arity1(c__966__auto___317).(float64), float64(0)
																				continue
																			}
																		} else {
																			{
																				var day_318 = cljs_core.First.X_invoke_Arity1(seq__220_316___1)
																				_ = day_318
																				{
																					var seq__221_319 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																					var chunk__222_320 interface{} = nil
																					var count__223_321 = float64(0)
																					var i__224_322 = float64(0)
																					_, _, _, _ = seq__221_319, chunk__222_320, count__223_321, i__224_322
																					for {
																						if i__224_322 < count__223_321 {
																							{
																								var hour_323 = cljs_core.Decorate_(chunk__222_320).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__224_322)
																								_ = hour_323
																								{
																									var s_324 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(month_298)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(day_318)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(hour_323)).(string) + ":14:15.666-06:00\"")
																									_ = s_324
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_324), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_324)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__221_319, chunk__222_320, count__223_321, i__224_322 = seq__221_319, chunk__222_320, count__223_321, (i__224_322 + float64(1))
																								continue
																							}
																						} else {
																							{
																								var temp__4222__auto___325___2 = cljs_core.Seq.Arity1IQ(seq__221_319)
																								_ = temp__4222__auto___325___2
																								if cljs_core.Truth_(temp__4222__auto___325___2) {
																									{
																										var seq__221_326___1 = temp__4222__auto___325___2
																										_ = seq__221_326___1
																										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__221_326___1) {
																											{
																												var c__966__auto___327 = cljs_core.Chunk_first.X_invoke_Arity1(seq__221_326___1)
																												_ = c__966__auto___327
																												seq__221_319, chunk__222_320, count__223_321, i__224_322 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__221_326___1), c__966__auto___327, cljs_core.Count.X_invoke_Arity1(c__966__auto___327).(float64), float64(0)
																												continue
																											}
																										} else {
																											{
																												var hour_328 = cljs_core.First.X_invoke_Arity1(seq__221_326___1)
																												_ = hour_328
																												{
																													var s_329 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(month_298)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(day_318)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_257.X_invoke_Arity1(hour_328)).(string) + ":14:15.666-06:00\"")
																													_ = s_329
																													if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_329), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_329)})).(string)), "ValueOf", []interface{}{})) {
																													} else {
																														panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																													}
																												}
																												seq__221_319, chunk__222_320, count__223_321, i__224_322 = cljs_core.Next.Arity1IQ(seq__221_326___1), nil, float64(0), float64(0)
																												continue
																											}
																										}
																									}
																								} else {
																								}
																							}
																						}
																						break
																					}
																				}
																				seq__220_299, chunk__225_300, count__226_301, i__227_302 = cljs_core.Next.Arity1IQ(seq__220_316___1), nil, float64(0), float64(0)
																				continue
																			}
																		}
																	}
																} else {
																}
															}
														}
														break
													}
												}
												seq__219_259, chunk__232_260, count__233_261, i__234_262 = cljs_core.Next.Arity1IQ(seq__219_296___1), nil, float64(0), float64(0)
												continue
											}
										}
									}
								} else {
								}
							}
						}
						break
					}
				}
			}
			{
				var insts_330 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01T00\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01T00:00:00.000\"")}, nil})
				_ = insts_330
				if cljs_core.Truth_(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_EQ_, cljs_core.Map_.X_invoke_Arity2(func(G__331 *cljs_core.AFn, insts_330 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__331, 1, func(p1__180_SHARP_ interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(p1__180_SHARP_, "ValueOf", []interface{}{})
					})
				}(&cljs_core.AFn{}, insts_330), insts_330).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (apply = (map (fn* [p1__180#] (.valueOf p1__180#)) insts))")}))
				}
			}
			{
				var u_332 = cljs_reader.Read_string.X_invoke_Arity1("#uuid \"550e8400-e29b-41d4-a716-446655440000\"")
				_ = u_332
				if cljs_core.X_EQ_.Arity2IIB(u_332, cljs_reader.Read_string.X_invoke_Arity1("#uuid \"550e8400-e29b-41d4-a716-446655440000\"")) {
				} else {
					panic((&js.Error{("Assert failed: (= u (reader/read-string \"#uuid \\\"550e8400-e29b-41d4-a716-446655440000\\\"\"))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(u_332, cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{u_332})).(string))) {
				} else {
					panic((&js.Error{("Assert failed: (= u (-> u pr-str reader/read-string))")}))
				}
			}
			cljs_reader.Register_tag_parser_BANG_.X_invoke_Arity2((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil}), cljs_core.Identity)
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), cljs_reader.Read_string.X_invoke_Arity1("#foo [1 2]")) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (reader/read-string \"#foo [1 2]\"))")}))
			}
			cljs_reader.Register_tag_parser_BANG_.X_invoke_Arity2((&cljs_core.CljsCoreSymbol{Ns: "foo.bar", Name: "baz", Str: "foo.bar/baz", X_hash: float64(-145013548), X_meta: nil}), cljs_core.Identity)
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), cljs_reader.Read_string.X_invoke_Arity1("#foo.bar/baz [1 2]")) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (reader/read-string \"#foo.bar/baz [1 2]\"))")}))
			}
			cljs_reader.Register_default_tag_parser_BANG_.X_invoke_Arity1(func(G__333 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__333, 2, func(tag interface{}, val interface{}) interface{} {
					return val
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), cljs_reader.Read_string.X_invoke_Arity1("#a.b/c [1 2]")) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (reader/read-string \"#a.b/c [1 2]\"))")}))
			}
			{
				var seq__247_334 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"\"abc \\ua\"", "\"abc \\x0z  ...etc\"", "\"abc \\u0g00 ..etc\""}, nil}))
				var chunk__248_335 interface{} = nil
				var count__249_336 = float64(0)
				var i__250_337 = float64(0)
				_, _, _, _ = seq__247_334, chunk__248_335, count__249_336, i__250_337
				for {
					if i__250_337 < count__249_336 {
						{
							var unicode_error_338 = cljs_core.Decorate_(chunk__248_335).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__250_337)
							_ = unicode_error_338
							{
								var r_339 = func() (return__340 *cljs_core.CljsCoreKeyword) {
									defer func() {
										if e251 := recover(); e251 != nil {
											if cljs_core.Value_(e251).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
												{
													var e = e251
													_ = e
													return__340 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
												}
											} else {
												panic(e251)

											}
										}
									}()
									{
										cljs_reader.Read_string.X_invoke_Arity1(unicode_error_338)
										return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed-to-throw", Fqn: "failed-to-throw", X_hash: float64(-1638969220)})
									}
								}()
								_ = r_339
								if cljs_core.X_EQ_.Arity2IIB(r_339, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
								} else {
									panic((&js.Error{("Assert failed: " + cljs_core.Str.X_invoke_Arity1(("Failed to throw reader error for: " + cljs_core.Str.X_invoke_Arity1(unicode_error_338).(string))).(string) + "\n(= r :ok)")}))
								}
							}
							seq__247_334, chunk__248_335, count__249_336, i__250_337 = seq__247_334, chunk__248_335, count__249_336, (i__250_337 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___341 = cljs_core.Seq.Arity1IQ(seq__247_334)
							_ = temp__4222__auto___341
							if cljs_core.Truth_(temp__4222__auto___341) {
								{
									var seq__247_342___1 = temp__4222__auto___341
									_ = seq__247_342___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__247_342___1) {
										{
											var c__966__auto___343 = cljs_core.Chunk_first.X_invoke_Arity1(seq__247_342___1)
											_ = c__966__auto___343
											seq__247_334, chunk__248_335, count__249_336, i__250_337 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__247_342___1), c__966__auto___343, cljs_core.Count.X_invoke_Arity1(c__966__auto___343).(float64), float64(0)
											continue
										}
									} else {
										{
											var unicode_error_344 = cljs_core.First.X_invoke_Arity1(seq__247_342___1)
											_ = unicode_error_344
											{
												var r_345 = func() (return__346 *cljs_core.CljsCoreKeyword) {
													defer func() {
														if e252 := recover(); e252 != nil {
															if cljs_core.Value_(e252).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
																{
																	var e = e252
																	_ = e
																	return__346 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
																}
															} else {
																panic(e252)

															}
														}
													}()
													{
														cljs_reader.Read_string.X_invoke_Arity1(unicode_error_344)
														return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed-to-throw", Fqn: "failed-to-throw", X_hash: float64(-1638969220)})
													}
												}()
												_ = r_345
												if cljs_core.X_EQ_.Arity2IIB(r_345, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
												} else {
													panic((&js.Error{("Assert failed: " + cljs_core.Str.X_invoke_Arity1(("Failed to throw reader error for: " + cljs_core.Str.X_invoke_Arity1(unicode_error_344).(string))).(string) + "\n(= r :ok)")}))
												}
											}
											seq__247_334, chunk__248_335, count__249_336, i__250_337 = cljs_core.Next.Arity1IQ(seq__247_342___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
							}
						}
					}
					break
				}
			}
			if cljs_core.Nil_(cljs_reader.Read_string.X_invoke_Arity1("")) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (reader/read-string \"\"))")}))
			}
			return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
		})
	}(&cljs_core.AFn{})

}

type CljsReader_testT struct {
	A interface{}
	B interface{}
}

var X__GT_T *cljs_core.AFn

type CljsReader_testR struct {
	A         interface{}
	B         interface{}
	X__meta   interface{}
	X__extmap interface{}
	X__hash   interface{}
}

func (_ *CljsReader_testR) CljsCoreILookup__() {}

func (this__770__auto__ *CljsReader_testR) X_lookup_Arity2(k__771__auto__ interface{}) interface{} {
	return this__770__auto__.X_lookup_Arity3(k__771__auto__, nil)
}

func (this__772__auto__ *CljsReader_testR) X_lookup_Arity3(k162 interface{}, else__773__auto__ interface{}) interface{} {
	{
		var G__165 = func() interface{} {
			if cljs_core.Value_(k162).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k162), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__165
		switch G__165 {
		case "b":
			return this__772__auto__.B

		case "a":
			return this__772__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__772__auto__.X__extmap, k162, else__773__auto__)

		}
	}
}

func (_ *CljsReader_testR) CljsCoreIPrintWithWriter__() {}

func (this__786__auto__ *CljsReader_testR) X_pr_writer_Arity3(writer__787__auto__ interface{}, opts__788__auto__ interface{}) interface{} {
	{
		var pr_pair__789__auto__ = func(G__348 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__348, 3, func(keyval__790__auto__ interface{}, ___791__auto__ interface{}, ___791__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__787__auto__, cljs_core.Pr_writer, "", " ", "", opts__788__auto__, keyval__790__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__789__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__787__auto__, pr_pair__789__auto__, "#cljs.reader-test.R{", ", ", "}", opts__788__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__786__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__786__auto__.B}, nil})}, nil}), this__786__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsReader_testR) CljsCoreIMeta__() {}

func (this__768__auto__ *CljsReader_testR) X_meta_Arity1() interface{} {
	return this__768__auto__.X__meta
}

func (_ *CljsReader_testR) CljsCoreICloneable__() {}

func (this__764__auto__ *CljsReader_testR) X_clone_Arity1() interface{} {
	return (&CljsReader_testR{this__764__auto__.A, this__764__auto__.B, this__764__auto__.X__meta, this__764__auto__.X__extmap, this__764__auto__.X__hash})
}

func (_ *CljsReader_testR) CljsCoreICounted__() {}

func (this__774__auto__ *CljsReader_testR) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__774__auto__.X__extmap).(float64))
}

func (_ *CljsReader_testR) CljsCoreIHash__() {}

func (this__765__auto__ *CljsReader_testR) X_hash_Arity1() interface{} {
	{
		var h__588__auto__ = this__765__auto__.X__hash
		_ = h__588__auto__
		if !(cljs_core.Nil_(h__588__auto__)) {
			return h__588__auto__
		} else {
			{
				var h__588__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__765__auto__).(float64)
				_ = h__588__auto_____1
				this__765__auto__.X__hash = h__588__auto_____1

				return h__588__auto_____1
			}
		}
	}
}

func (_ *CljsReader_testR) CljsCoreIEquiv__() {}

func (this__766__auto__ *CljsReader_testR) X_equiv_Arity2(other__767__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__165__auto__ = other__767__auto__
		_ = and__165__auto__
		if cljs_core.Truth_(and__165__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__766__auto__), cljs_core.Type_.X_invoke_Arity1(other__767__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__766__auto__, other__767__auto__)))
		} else {
			return and__165__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsReader_testR) CljsCoreIRecord__() {}

func (_ *CljsReader_testR) CljsCoreIMap__() {}

func (this__781__auto__ *CljsReader_testR) X_dissoc_Arity2(k__782__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__782__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__781__auto__), this__781__auto__.X__meta), k__782__auto__)
	} else {
		return (&CljsReader_testR{this__781__auto__.A, this__781__auto__.B, this__781__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__781__auto__.X__extmap, k__782__auto__)), nil})
	}
}

func (_ *CljsReader_testR) CljsCoreIAssociative__() {}

func (this__777__auto__ *CljsReader_testR) X_assoc_Arity3(k__778__auto__ interface{}, G__161 interface{}) interface{} {
	{
		var pred__173 = cljs_core.Keyword_identical_QMARK_
		var expr__174 = k__778__auto__
		_, _ = pred__173, expr__174
		if cljs_core.Truth_(func() interface{} {
			var G__176 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__177 = expr__174
			_, _ = G__176, G__177
			return pred__173.X_invoke_Arity2(G__176, G__177)
		}()) {
			return (&CljsReader_testR{G__161, this__777__auto__.B, this__777__auto__.X__meta, this__777__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__178 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__179 = expr__174
				_, _ = G__178, G__179
				return pred__173.X_invoke_Arity2(G__178, G__179)
			}()) {
				return (&CljsReader_testR{this__777__auto__.A, G__161, this__777__auto__.X__meta, this__777__auto__.X__extmap, nil})
			} else {
				return (&CljsReader_testR{this__777__auto__.A, this__777__auto__.B, this__777__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__777__auto__.X__extmap, k__778__auto__, G__161), nil})
			}
		}
	}
}

func (this__779__auto__ *CljsReader_testR) X_contains_key_QMARK__Arity2(k__780__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsReader_testR) CljsCoreISeqable__() {}

func (this__784__auto__ *CljsReader_testR) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__784__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__784__auto__.B}, nil})}, nil}), this__784__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsReader_testR) CljsCoreIWithMeta__() {}

func (this__769__auto__ *CljsReader_testR) X_with_meta_Arity2(G__161 interface{}) interface{} {
	return (&CljsReader_testR{this__769__auto__.A, this__769__auto__.B, G__161, this__769__auto__.X__extmap, this__769__auto__.X__hash})
}

func (_ *CljsReader_testR) CljsCoreICollection__() {}

func (this__775__auto__ *CljsReader_testR) X_conj_Arity2(entry__776__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__776__auto__) {
		return this__775__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__776__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__776__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__775__auto__, entry__776__auto__)
	}
}

var X__GT_R *cljs_core.AFn

var Map__GT_R *cljs_core.AFn

var Test_reader *cljs_core.AFn

func Test_runner(t *testing.T) {
	assert.Equal(t, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), Test_reader.X_invoke_Arity0().(*cljs_core.CljsCoreKeyword))
}
