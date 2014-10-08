// Compiled by ClojureScript to Go 0.0-2356
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
			return cljs_core.Fn(map__GT_R, 1, func(G__4183 interface{}) interface{} {
				return (&CljsReader_testR{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__4183), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__4183), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__4183, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})})), nil})
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
			if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__4273 interface{}) {
				defer func() {
					if e4237 := recover(); e4237 != nil {
						if cljs_core.Value_(e4237).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
							{
								var e = e4237
								_ = e
								return__4273 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "threw", Fqn: "threw", X_hash: float64(17630075)})
							}
						} else {
							panic(e4237)

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
			if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__4274 interface{}) {
				defer func() {
					if e4238 := recover(); e4238 != nil {
						if cljs_core.Value_(e4238).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
							{
								var e = e4238
								_ = e
								return__4274 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "threw", Fqn: "threw", X_hash: float64(17630075)})
							}
						} else {
							panic(e4238)

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
				var est_inst_4275 = cljs_reader.Read_string.X_invoke_Arity1("#inst \"2010-11-12T13:14:15.666-05:00\"")
				var utc_inst_4276 = cljs_reader.Read_string.X_invoke_Arity1("#inst \"2010-11-12T18:14:15.666-00:00\"")
				var pad_4277 = func(G__4278 *cljs_core.AFn, est_inst_4275 interface{}, utc_inst_4276 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__4278, 1, func(n interface{}) interface{} {
						if n.(float64) < float64(10) {
							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
						} else {
							return n
						}
					})
				}(&cljs_core.AFn{}, est_inst_4275, utc_inst_4276)
				_, _, _ = est_inst_4275, utc_inst_4276, pad_4277
				if cljs_core.X_EQ_.Arity2IIB((&js.Date{"2010-11-12T13:14:15.666-05:00"}).ValueOf(), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_4275, "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf (js/Date. \"2010-11-12T13:14:15.666-05:00\")) (.valueOf est-inst))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_4275, "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{est_inst_4275})).(string)), "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf est-inst) (.valueOf (reader/read-string (pr-str est-inst))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_4275, "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(utc_inst_4276, "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf est-inst) (.valueOf utc-inst))")}))
				}
				{
					var seq__4239_4279 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(13)).(*cljs_core.CljsCoreRange))
					var chunk__4252_4280 interface{} = nil
					var count__4253_4281 = float64(0)
					var i__4254_4282 = float64(0)
					_, _, _, _ = seq__4239_4279, chunk__4252_4280, count__4253_4281, i__4254_4282
					for {
						if i__4254_4282 < count__4253_4281 {
							{
								var month_4283 = cljs_core.Decorate_(chunk__4252_4280).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4254_4282)
								_ = month_4283
								{
									var seq__4255_4284 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
									var chunk__4260_4285 interface{} = nil
									var count__4261_4286 = float64(0)
									var i__4262_4287 = float64(0)
									_, _, _, _ = seq__4255_4284, chunk__4260_4285, count__4261_4286, i__4262_4287
									for {
										if i__4262_4287 < count__4261_4286 {
											{
												var day_4288 = cljs_core.Decorate_(chunk__4260_4285).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4262_4287)
												_ = day_4288
												{
													var seq__4263_4289 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
													var chunk__4264_4290 interface{} = nil
													var count__4265_4291 = float64(0)
													var i__4266_4292 = float64(0)
													_, _, _, _ = seq__4263_4289, chunk__4264_4290, count__4265_4291, i__4266_4292
													for {
														if i__4266_4292 < count__4265_4291 {
															{
																var hour_4293 = cljs_core.Decorate_(chunk__4264_4290).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4266_4292)
																_ = hour_4293
																{
																	var s_4294 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(month_4283)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(day_4288)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(hour_4293)).(string) + ":14:15.666-06:00\"")
																	_ = s_4294
																	if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4294), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4294)})).(string)), "ValueOf", []interface{}{})) {
																	} else {
																		panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																	}
																}
																seq__4263_4289, chunk__4264_4290, count__4265_4291, i__4266_4292 = seq__4263_4289, chunk__4264_4290, count__4265_4291, (i__4266_4292 + float64(1))
																continue
															}
														} else {
															{
																var temp__4222__auto___4295 = cljs_core.Seq.Arity1IQ(seq__4263_4289)
																_ = temp__4222__auto___4295
																if cljs_core.Truth_(temp__4222__auto___4295) {
																	{
																		var seq__4263_4296___1 = temp__4222__auto___4295
																		_ = seq__4263_4296___1
																		if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4263_4296___1) {
																			{
																				var c__965__auto___4297 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4263_4296___1)
																				_ = c__965__auto___4297
																				seq__4263_4289, chunk__4264_4290, count__4265_4291, i__4266_4292 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4263_4296___1), c__965__auto___4297, cljs_core.Count.X_invoke_Arity1(c__965__auto___4297).(float64), float64(0)
																				continue
																			}
																		} else {
																			{
																				var hour_4298 = cljs_core.First.X_invoke_Arity1(seq__4263_4296___1)
																				_ = hour_4298
																				{
																					var s_4299 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(month_4283)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(day_4288)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(hour_4298)).(string) + ":14:15.666-06:00\"")
																					_ = s_4299
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4299), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4299)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__4263_4289, chunk__4264_4290, count__4265_4291, i__4266_4292 = cljs_core.Next.Arity1IQ(seq__4263_4296___1), nil, float64(0), float64(0)
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
												seq__4255_4284, chunk__4260_4285, count__4261_4286, i__4262_4287 = seq__4255_4284, chunk__4260_4285, count__4261_4286, (i__4262_4287 + float64(1))
												continue
											}
										} else {
											{
												var temp__4222__auto___4300 = cljs_core.Seq.Arity1IQ(seq__4255_4284)
												_ = temp__4222__auto___4300
												if cljs_core.Truth_(temp__4222__auto___4300) {
													{
														var seq__4255_4301___1 = temp__4222__auto___4300
														_ = seq__4255_4301___1
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4255_4301___1) {
															{
																var c__965__auto___4302 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4255_4301___1)
																_ = c__965__auto___4302
																seq__4255_4284, chunk__4260_4285, count__4261_4286, i__4262_4287 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4255_4301___1), c__965__auto___4302, cljs_core.Count.X_invoke_Arity1(c__965__auto___4302).(float64), float64(0)
																continue
															}
														} else {
															{
																var day_4303 = cljs_core.First.X_invoke_Arity1(seq__4255_4301___1)
																_ = day_4303
																{
																	var seq__4256_4304 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																	var chunk__4257_4305 interface{} = nil
																	var count__4258_4306 = float64(0)
																	var i__4259_4307 = float64(0)
																	_, _, _, _ = seq__4256_4304, chunk__4257_4305, count__4258_4306, i__4259_4307
																	for {
																		if i__4259_4307 < count__4258_4306 {
																			{
																				var hour_4308 = cljs_core.Decorate_(chunk__4257_4305).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4259_4307)
																				_ = hour_4308
																				{
																					var s_4309 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(month_4283)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(day_4303)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(hour_4308)).(string) + ":14:15.666-06:00\"")
																					_ = s_4309
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4309), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4309)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__4256_4304, chunk__4257_4305, count__4258_4306, i__4259_4307 = seq__4256_4304, chunk__4257_4305, count__4258_4306, (i__4259_4307 + float64(1))
																				continue
																			}
																		} else {
																			{
																				var temp__4222__auto___4310___1 = cljs_core.Seq.Arity1IQ(seq__4256_4304)
																				_ = temp__4222__auto___4310___1
																				if cljs_core.Truth_(temp__4222__auto___4310___1) {
																					{
																						var seq__4256_4311___1 = temp__4222__auto___4310___1
																						_ = seq__4256_4311___1
																						if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4256_4311___1) {
																							{
																								var c__965__auto___4312 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4256_4311___1)
																								_ = c__965__auto___4312
																								seq__4256_4304, chunk__4257_4305, count__4258_4306, i__4259_4307 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4256_4311___1), c__965__auto___4312, cljs_core.Count.X_invoke_Arity1(c__965__auto___4312).(float64), float64(0)
																								continue
																							}
																						} else {
																							{
																								var hour_4313 = cljs_core.First.X_invoke_Arity1(seq__4256_4311___1)
																								_ = hour_4313
																								{
																									var s_4314 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(month_4283)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(day_4303)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(hour_4313)).(string) + ":14:15.666-06:00\"")
																									_ = s_4314
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4314), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4314)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__4256_4304, chunk__4257_4305, count__4258_4306, i__4259_4307 = cljs_core.Next.Arity1IQ(seq__4256_4311___1), nil, float64(0), float64(0)
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
																seq__4255_4284, chunk__4260_4285, count__4261_4286, i__4262_4287 = cljs_core.Next.Arity1IQ(seq__4255_4301___1), nil, float64(0), float64(0)
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
								seq__4239_4279, chunk__4252_4280, count__4253_4281, i__4254_4282 = seq__4239_4279, chunk__4252_4280, count__4253_4281, (i__4254_4282 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___4315 = cljs_core.Seq.Arity1IQ(seq__4239_4279)
								_ = temp__4222__auto___4315
								if cljs_core.Truth_(temp__4222__auto___4315) {
									{
										var seq__4239_4316___1 = temp__4222__auto___4315
										_ = seq__4239_4316___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4239_4316___1) {
											{
												var c__965__auto___4317 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4239_4316___1)
												_ = c__965__auto___4317
												seq__4239_4279, chunk__4252_4280, count__4253_4281, i__4254_4282 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4239_4316___1), c__965__auto___4317, cljs_core.Count.X_invoke_Arity1(c__965__auto___4317).(float64), float64(0)
												continue
											}
										} else {
											{
												var month_4318 = cljs_core.First.X_invoke_Arity1(seq__4239_4316___1)
												_ = month_4318
												{
													var seq__4240_4319 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
													var chunk__4245_4320 interface{} = nil
													var count__4246_4321 = float64(0)
													var i__4247_4322 = float64(0)
													_, _, _, _ = seq__4240_4319, chunk__4245_4320, count__4246_4321, i__4247_4322
													for {
														if i__4247_4322 < count__4246_4321 {
															{
																var day_4323 = cljs_core.Decorate_(chunk__4245_4320).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4247_4322)
																_ = day_4323
																{
																	var seq__4248_4324 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																	var chunk__4249_4325 interface{} = nil
																	var count__4250_4326 = float64(0)
																	var i__4251_4327 = float64(0)
																	_, _, _, _ = seq__4248_4324, chunk__4249_4325, count__4250_4326, i__4251_4327
																	for {
																		if i__4251_4327 < count__4250_4326 {
																			{
																				var hour_4328 = cljs_core.Decorate_(chunk__4249_4325).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4251_4327)
																				_ = hour_4328
																				{
																					var s_4329 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(month_4318)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(day_4323)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(hour_4328)).(string) + ":14:15.666-06:00\"")
																					_ = s_4329
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4329), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4329)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__4248_4324, chunk__4249_4325, count__4250_4326, i__4251_4327 = seq__4248_4324, chunk__4249_4325, count__4250_4326, (i__4251_4327 + float64(1))
																				continue
																			}
																		} else {
																			{
																				var temp__4222__auto___4330___1 = cljs_core.Seq.Arity1IQ(seq__4248_4324)
																				_ = temp__4222__auto___4330___1
																				if cljs_core.Truth_(temp__4222__auto___4330___1) {
																					{
																						var seq__4248_4331___1 = temp__4222__auto___4330___1
																						_ = seq__4248_4331___1
																						if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4248_4331___1) {
																							{
																								var c__965__auto___4332 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4248_4331___1)
																								_ = c__965__auto___4332
																								seq__4248_4324, chunk__4249_4325, count__4250_4326, i__4251_4327 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4248_4331___1), c__965__auto___4332, cljs_core.Count.X_invoke_Arity1(c__965__auto___4332).(float64), float64(0)
																								continue
																							}
																						} else {
																							{
																								var hour_4333 = cljs_core.First.X_invoke_Arity1(seq__4248_4331___1)
																								_ = hour_4333
																								{
																									var s_4334 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(month_4318)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(day_4323)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(hour_4333)).(string) + ":14:15.666-06:00\"")
																									_ = s_4334
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4334), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4334)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__4248_4324, chunk__4249_4325, count__4250_4326, i__4251_4327 = cljs_core.Next.Arity1IQ(seq__4248_4331___1), nil, float64(0), float64(0)
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
																seq__4240_4319, chunk__4245_4320, count__4246_4321, i__4247_4322 = seq__4240_4319, chunk__4245_4320, count__4246_4321, (i__4247_4322 + float64(1))
																continue
															}
														} else {
															{
																var temp__4222__auto___4335___1 = cljs_core.Seq.Arity1IQ(seq__4240_4319)
																_ = temp__4222__auto___4335___1
																if cljs_core.Truth_(temp__4222__auto___4335___1) {
																	{
																		var seq__4240_4336___1 = temp__4222__auto___4335___1
																		_ = seq__4240_4336___1
																		if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4240_4336___1) {
																			{
																				var c__965__auto___4337 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4240_4336___1)
																				_ = c__965__auto___4337
																				seq__4240_4319, chunk__4245_4320, count__4246_4321, i__4247_4322 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4240_4336___1), c__965__auto___4337, cljs_core.Count.X_invoke_Arity1(c__965__auto___4337).(float64), float64(0)
																				continue
																			}
																		} else {
																			{
																				var day_4338 = cljs_core.First.X_invoke_Arity1(seq__4240_4336___1)
																				_ = day_4338
																				{
																					var seq__4241_4339 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																					var chunk__4242_4340 interface{} = nil
																					var count__4243_4341 = float64(0)
																					var i__4244_4342 = float64(0)
																					_, _, _, _ = seq__4241_4339, chunk__4242_4340, count__4243_4341, i__4244_4342
																					for {
																						if i__4244_4342 < count__4243_4341 {
																							{
																								var hour_4343 = cljs_core.Decorate_(chunk__4242_4340).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4244_4342)
																								_ = hour_4343
																								{
																									var s_4344 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(month_4318)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(day_4338)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(hour_4343)).(string) + ":14:15.666-06:00\"")
																									_ = s_4344
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4344), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4344)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__4241_4339, chunk__4242_4340, count__4243_4341, i__4244_4342 = seq__4241_4339, chunk__4242_4340, count__4243_4341, (i__4244_4342 + float64(1))
																								continue
																							}
																						} else {
																							{
																								var temp__4222__auto___4345___2 = cljs_core.Seq.Arity1IQ(seq__4241_4339)
																								_ = temp__4222__auto___4345___2
																								if cljs_core.Truth_(temp__4222__auto___4345___2) {
																									{
																										var seq__4241_4346___1 = temp__4222__auto___4345___2
																										_ = seq__4241_4346___1
																										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4241_4346___1) {
																											{
																												var c__965__auto___4347 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4241_4346___1)
																												_ = c__965__auto___4347
																												seq__4241_4339, chunk__4242_4340, count__4243_4341, i__4244_4342 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4241_4346___1), c__965__auto___4347, cljs_core.Count.X_invoke_Arity1(c__965__auto___4347).(float64), float64(0)
																												continue
																											}
																										} else {
																											{
																												var hour_4348 = cljs_core.First.X_invoke_Arity1(seq__4241_4346___1)
																												_ = hour_4348
																												{
																													var s_4349 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(month_4318)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(day_4338)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4277.X_invoke_Arity1(hour_4348)).(string) + ":14:15.666-06:00\"")
																													_ = s_4349
																													if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4349), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4349)})).(string)), "ValueOf", []interface{}{})) {
																													} else {
																														panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																													}
																												}
																												seq__4241_4339, chunk__4242_4340, count__4243_4341, i__4244_4342 = cljs_core.Next.Arity1IQ(seq__4241_4346___1), nil, float64(0), float64(0)
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
																				seq__4240_4319, chunk__4245_4320, count__4246_4321, i__4247_4322 = cljs_core.Next.Arity1IQ(seq__4240_4336___1), nil, float64(0), float64(0)
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
												seq__4239_4279, chunk__4252_4280, count__4253_4281, i__4254_4282 = cljs_core.Next.Arity1IQ(seq__4239_4316___1), nil, float64(0), float64(0)
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
				var insts_4350 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01T00\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01T00:00:00.000\"")}, nil})
				_ = insts_4350
				if cljs_core.Truth_(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_EQ_, cljs_core.Map_.X_invoke_Arity2(func(G__4351 *cljs_core.AFn, insts_4350 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__4351, 1, func(p1__4200_SHARP_ interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(p1__4200_SHARP_, "ValueOf", []interface{}{})
					})
				}(&cljs_core.AFn{}, insts_4350), insts_4350).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (apply = (map (fn* [p1__4200#] (.valueOf p1__4200#)) insts))")}))
				}
			}
			{
				var u_4352 = cljs_reader.Read_string.X_invoke_Arity1("#uuid \"550e8400-e29b-41d4-a716-446655440000\"")
				_ = u_4352
				if cljs_core.X_EQ_.Arity2IIB(u_4352, cljs_reader.Read_string.X_invoke_Arity1("#uuid \"550e8400-e29b-41d4-a716-446655440000\"")) {
				} else {
					panic((&js.Error{("Assert failed: (= u (reader/read-string \"#uuid \\\"550e8400-e29b-41d4-a716-446655440000\\\"\"))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(u_4352, cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{u_4352})).(string))) {
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
			cljs_reader.Register_default_tag_parser_BANG_.X_invoke_Arity1(func(G__4353 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4353, 2, func(tag interface{}, val interface{}) interface{} {
					return val
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), cljs_reader.Read_string.X_invoke_Arity1("#a.b/c [1 2]")) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (reader/read-string \"#a.b/c [1 2]\"))")}))
			}
			{
				var seq__4267_4354 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"\"abc \\ua\"", "\"abc \\x0z  ...etc\"", "\"abc \\u0g00 ..etc\""}, nil}))
				var chunk__4268_4355 interface{} = nil
				var count__4269_4356 = float64(0)
				var i__4270_4357 = float64(0)
				_, _, _, _ = seq__4267_4354, chunk__4268_4355, count__4269_4356, i__4270_4357
				for {
					if i__4270_4357 < count__4269_4356 {
						{
							var unicode_error_4358 = cljs_core.Decorate_(chunk__4268_4355).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4270_4357)
							_ = unicode_error_4358
							{
								var r_4359 = func() (return__4360 *cljs_core.CljsCoreKeyword) {
									defer func() {
										if e4271 := recover(); e4271 != nil {
											if cljs_core.Value_(e4271).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
												{
													var e = e4271
													_ = e
													return__4360 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
												}
											} else {
												panic(e4271)

											}
										}
									}()
									{
										cljs_reader.Read_string.X_invoke_Arity1(unicode_error_4358)
										return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed-to-throw", Fqn: "failed-to-throw", X_hash: float64(-1638969220)})
									}
								}()
								_ = r_4359
								if cljs_core.X_EQ_.Arity2IIB(r_4359, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
								} else {
									panic((&js.Error{("Assert failed: " + cljs_core.Str.X_invoke_Arity1(("Failed to throw reader error for: " + cljs_core.Str.X_invoke_Arity1(unicode_error_4358).(string))).(string) + "\n(= r :ok)")}))
								}
							}
							seq__4267_4354, chunk__4268_4355, count__4269_4356, i__4270_4357 = seq__4267_4354, chunk__4268_4355, count__4269_4356, (i__4270_4357 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___4361 = cljs_core.Seq.Arity1IQ(seq__4267_4354)
							_ = temp__4222__auto___4361
							if cljs_core.Truth_(temp__4222__auto___4361) {
								{
									var seq__4267_4362___1 = temp__4222__auto___4361
									_ = seq__4267_4362___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4267_4362___1) {
										{
											var c__965__auto___4363 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4267_4362___1)
											_ = c__965__auto___4363
											seq__4267_4354, chunk__4268_4355, count__4269_4356, i__4270_4357 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4267_4362___1), c__965__auto___4363, cljs_core.Count.X_invoke_Arity1(c__965__auto___4363).(float64), float64(0)
											continue
										}
									} else {
										{
											var unicode_error_4364 = cljs_core.First.X_invoke_Arity1(seq__4267_4362___1)
											_ = unicode_error_4364
											{
												var r_4365 = func() (return__4366 *cljs_core.CljsCoreKeyword) {
													defer func() {
														if e4272 := recover(); e4272 != nil {
															if cljs_core.Value_(e4272).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
																{
																	var e = e4272
																	_ = e
																	return__4366 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
																}
															} else {
																panic(e4272)

															}
														}
													}()
													{
														cljs_reader.Read_string.X_invoke_Arity1(unicode_error_4364)
														return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed-to-throw", Fqn: "failed-to-throw", X_hash: float64(-1638969220)})
													}
												}()
												_ = r_4365
												if cljs_core.X_EQ_.Arity2IIB(r_4365, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
												} else {
													panic((&js.Error{("Assert failed: " + cljs_core.Str.X_invoke_Arity1(("Failed to throw reader error for: " + cljs_core.Str.X_invoke_Arity1(unicode_error_4364).(string))).(string) + "\n(= r :ok)")}))
												}
											}
											seq__4267_4354, chunk__4268_4355, count__4269_4356, i__4270_4357 = cljs_core.Next.Arity1IQ(seq__4267_4362___1), nil, float64(0), float64(0)
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

func (this__769__auto__ *CljsReader_testR) X_lookup_Arity2(k__770__auto__ interface{}) interface{} {
	return this__769__auto__.X_lookup_Arity3(k__770__auto__, nil)
}

func (this__771__auto__ *CljsReader_testR) X_lookup_Arity3(k4182 interface{}, else__772__auto__ interface{}) interface{} {
	{
		var G__4185 = func() interface{} {
			if cljs_core.Value_(k4182).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k4182), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4185
		switch G__4185 {
		case "b":
			return this__771__auto__.B

		case "a":
			return this__771__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__771__auto__.X__extmap, k4182, else__772__auto__)

		}
	}
}

func (_ *CljsReader_testR) CljsCoreIPrintWithWriter__() {}

func (this__785__auto__ *CljsReader_testR) X_pr_writer_Arity3(writer__786__auto__ interface{}, opts__787__auto__ interface{}) interface{} {
	{
		var pr_pair__788__auto__ = func(G__4368 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__4368, 3, func(keyval__789__auto__ interface{}, ___790__auto__ interface{}, ___790__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__786__auto__, cljs_core.Pr_writer, "", " ", "", opts__787__auto__, keyval__789__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__788__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__786__auto__, pr_pair__788__auto__, "#cljs.reader-test.R{", ", ", "}", opts__787__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__785__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__785__auto__.B}, nil})}, nil}), this__785__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsReader_testR) CljsCoreIMeta__() {}

func (this__767__auto__ *CljsReader_testR) X_meta_Arity1() interface{} {
	return this__767__auto__.X__meta
}

func (_ *CljsReader_testR) CljsCoreICloneable__() {}

func (this__763__auto__ *CljsReader_testR) X_clone_Arity1() interface{} {
	return (&CljsReader_testR{this__763__auto__.A, this__763__auto__.B, this__763__auto__.X__meta, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsReader_testR) CljsCoreICounted__() {}

func (this__773__auto__ *CljsReader_testR) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__773__auto__.X__extmap).(float64))
}

func (_ *CljsReader_testR) CljsCoreIHash__() {}

func (this__764__auto__ *CljsReader_testR) X_hash_Arity1() interface{} {
	{
		var h__588__auto__ = this__764__auto__.X__hash
		_ = h__588__auto__
		if !(cljs_core.Nil_(h__588__auto__)) {
			return h__588__auto__
		} else {
			{
				var h__588__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__764__auto__).(float64)
				_ = h__588__auto_____1
				this__764__auto__.X__hash = h__588__auto_____1

				return h__588__auto_____1
			}
		}
	}
}

func (_ *CljsReader_testR) CljsCoreIEquiv__() {}

func (this__765__auto__ *CljsReader_testR) X_equiv_Arity2(other__766__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__165__auto__ = other__766__auto__
		_ = and__165__auto__
		if cljs_core.Truth_(and__165__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__765__auto__), cljs_core.Type_.X_invoke_Arity1(other__766__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__765__auto__, other__766__auto__)))
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

func (this__780__auto__ *CljsReader_testR) X_dissoc_Arity2(k__781__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__781__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__780__auto__), this__780__auto__.X__meta), k__781__auto__)
	} else {
		return (&CljsReader_testR{this__780__auto__.A, this__780__auto__.B, this__780__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__780__auto__.X__extmap, k__781__auto__)), nil})
	}
}

func (_ *CljsReader_testR) CljsCoreIAssociative__() {}

func (this__776__auto__ *CljsReader_testR) X_assoc_Arity3(k__777__auto__ interface{}, G__4181 interface{}) interface{} {
	{
		var pred__4193 = cljs_core.Keyword_identical_QMARK_
		var expr__4194 = k__777__auto__
		_, _ = pred__4193, expr__4194
		if cljs_core.Truth_(func() interface{} {
			var G__4196 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__4197 = expr__4194
			_, _ = G__4196, G__4197
			return pred__4193.X_invoke_Arity2(G__4196, G__4197)
		}()) {
			return (&CljsReader_testR{G__4181, this__776__auto__.B, this__776__auto__.X__meta, this__776__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__4198 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__4199 = expr__4194
				_, _ = G__4198, G__4199
				return pred__4193.X_invoke_Arity2(G__4198, G__4199)
			}()) {
				return (&CljsReader_testR{this__776__auto__.A, G__4181, this__776__auto__.X__meta, this__776__auto__.X__extmap, nil})
			} else {
				return (&CljsReader_testR{this__776__auto__.A, this__776__auto__.B, this__776__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__776__auto__.X__extmap, k__777__auto__, G__4181), nil})
			}
		}
	}
}

func (this__778__auto__ *CljsReader_testR) X_contains_key_QMARK__Arity2(k__779__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsReader_testR) CljsCoreISeqable__() {}

func (this__783__auto__ *CljsReader_testR) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__783__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__783__auto__.B}, nil})}, nil}), this__783__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsReader_testR) CljsCoreIWithMeta__() {}

func (this__768__auto__ *CljsReader_testR) X_with_meta_Arity2(G__4181 interface{}) interface{} {
	return (&CljsReader_testR{this__768__auto__.A, this__768__auto__.B, G__4181, this__768__auto__.X__extmap, this__768__auto__.X__hash})
}

func (_ *CljsReader_testR) CljsCoreICollection__() {}

func (this__774__auto__ *CljsReader_testR) X_conj_Arity2(entry__775__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__775__auto__) {
		return this__774__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__775__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__775__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__774__auto__, entry__775__auto__)
	}
}

var X__GT_R *cljs_core.AFn

var Map__GT_R *cljs_core.AFn

var Test_reader *cljs_core.AFn

func Test_runner(t *testing.T) {
	assert.Equal(t, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), Test_reader.X_invoke_Arity0().(*cljs_core.CljsCoreKeyword))
}
