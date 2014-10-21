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
			return cljs_core.Fn(map__GT_R, 1, func(G__4189 interface{}) interface{} {
				return (&CljsReader_testR{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__4189), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__4189), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__4189, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})})), nil})
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
			if cljs_core.Truth_(cljs_core.Apply.X_invoke_Arity3(cljs_core.X_EQ_, float64(42), cljs_core.Map_.X_invoke_Arity2(cljs_reader.Read_string, (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"052", "0x2a", "2r101010", "8R52", "16r2a", "36r16"}, nil})).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{("Assert failed: (apply = 42 (map reader/read-string [\"052\" \"0x2a\" \"2r101010\" \"8R52\" \"16r2a\" \"36r16\"]))")}))
			}
			if cljs_core.Truth_(cljs_core.Apply.X_invoke_Arity3(cljs_core.X_EQ_, float64(42), cljs_core.Map_.X_invoke_Arity2(cljs_reader.Read_string, (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"+052", "+0x2a", "+2r101010", "+8r52", "+16R2a", "+36r16"}, nil})).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{("Assert failed: (apply = 42 (map reader/read-string [\"+052\" \"+0x2a\" \"+2r101010\" \"+8r52\" \"+16R2a\" \"+36r16\"]))")}))
			}
			if cljs_core.Truth_(cljs_core.Apply.X_invoke_Arity3(cljs_core.X_EQ_, float64(-42), cljs_core.Map_.X_invoke_Arity2(cljs_reader.Read_string, (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"-052", "-0X2a", "-2r101010", "-8r52", "-16r2a", "-36R16"}, nil})).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{("Assert failed: (apply = -42 (map reader/read-string [\"-052\" \"-0X2a\" \"-2r101010\" \"-8r52\" \"-16r2a\" \"-36R16\"]))")}))
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
			if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__4279 interface{}) {
				defer func() {
					if e4243 := recover(); e4243 != nil {
						if cljs_core.Value_(e4243).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
							{
								var e = e4243
								_ = e
								return__4279 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "threw", Fqn: "threw", X_hash: float64(17630075)})
							}
						} else {
							panic(e4243)

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
			if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__4280 interface{}) {
				defer func() {
					if e4244 := recover(); e4244 != nil {
						if cljs_core.Value_(e4244).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
							{
								var e = e4244
								_ = e
								return__4280 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "threw", Fqn: "threw", X_hash: float64(17630075)})
							}
						} else {
							panic(e4244)

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
				var est_inst_4281 = cljs_reader.Read_string.X_invoke_Arity1("#inst \"2010-11-12T13:14:15.666-05:00\"")
				var utc_inst_4282 = cljs_reader.Read_string.X_invoke_Arity1("#inst \"2010-11-12T18:14:15.666-00:00\"")
				var pad_4283 = func(G__4284 *cljs_core.AFn, est_inst_4281 interface{}, utc_inst_4282 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__4284, 1, func(n interface{}) interface{} {
						if n.(float64) < float64(10) {
							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
						} else {
							return n
						}
					})
				}(&cljs_core.AFn{}, est_inst_4281, utc_inst_4282)
				_, _, _ = est_inst_4281, utc_inst_4282, pad_4283
				if cljs_core.X_EQ_.Arity2IIB((&js.Date{"2010-11-12T13:14:15.666-05:00"}).ValueOf(), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_4281, "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf (js/Date. \"2010-11-12T13:14:15.666-05:00\")) (.valueOf est-inst))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_4281, "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{est_inst_4281})).(string)), "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf est-inst) (.valueOf (reader/read-string (pr-str est-inst))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_4281, "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(utc_inst_4282, "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf est-inst) (.valueOf utc-inst))")}))
				}
				{
					var seq__4245_4285 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(13)).(*cljs_core.CljsCoreRange))
					var chunk__4258_4286 interface{} = nil
					var count__4259_4287 = float64(0)
					var i__4260_4288 = float64(0)
					_, _, _, _ = seq__4245_4285, chunk__4258_4286, count__4259_4287, i__4260_4288
					for {
						if i__4260_4288 < count__4259_4287 {
							{
								var month_4289 = cljs_core.Decorate_(chunk__4258_4286).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4260_4288)
								_ = month_4289
								{
									var seq__4261_4290 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
									var chunk__4266_4291 interface{} = nil
									var count__4267_4292 = float64(0)
									var i__4268_4293 = float64(0)
									_, _, _, _ = seq__4261_4290, chunk__4266_4291, count__4267_4292, i__4268_4293
									for {
										if i__4268_4293 < count__4267_4292 {
											{
												var day_4294 = cljs_core.Decorate_(chunk__4266_4291).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4268_4293)
												_ = day_4294
												{
													var seq__4269_4295 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
													var chunk__4270_4296 interface{} = nil
													var count__4271_4297 = float64(0)
													var i__4272_4298 = float64(0)
													_, _, _, _ = seq__4269_4295, chunk__4270_4296, count__4271_4297, i__4272_4298
													for {
														if i__4272_4298 < count__4271_4297 {
															{
																var hour_4299 = cljs_core.Decorate_(chunk__4270_4296).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4272_4298)
																_ = hour_4299
																{
																	var s_4300 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(month_4289)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(day_4294)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(hour_4299)).(string) + ":14:15.666-06:00\"")
																	_ = s_4300
																	if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4300), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4300)})).(string)), "ValueOf", []interface{}{})) {
																	} else {
																		panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																	}
																}
																seq__4269_4295, chunk__4270_4296, count__4271_4297, i__4272_4298 = seq__4269_4295, chunk__4270_4296, count__4271_4297, (i__4272_4298 + float64(1))
																continue
															}
														} else {
															{
																var temp__4222__auto___4301 = cljs_core.Seq.Arity1IQ(seq__4269_4295)
																_ = temp__4222__auto___4301
																if cljs_core.Truth_(temp__4222__auto___4301) {
																	{
																		var seq__4269_4302___1 = temp__4222__auto___4301
																		_ = seq__4269_4302___1
																		if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4269_4302___1) {
																			{
																				var c__966__auto___4303 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4269_4302___1)
																				_ = c__966__auto___4303
																				seq__4269_4295, chunk__4270_4296, count__4271_4297, i__4272_4298 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4269_4302___1), c__966__auto___4303, cljs_core.Count.X_invoke_Arity1(c__966__auto___4303).(float64), float64(0)
																				continue
																			}
																		} else {
																			{
																				var hour_4304 = cljs_core.First.X_invoke_Arity1(seq__4269_4302___1)
																				_ = hour_4304
																				{
																					var s_4305 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(month_4289)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(day_4294)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(hour_4304)).(string) + ":14:15.666-06:00\"")
																					_ = s_4305
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4305), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4305)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__4269_4295, chunk__4270_4296, count__4271_4297, i__4272_4298 = cljs_core.Next.Arity1IQ(seq__4269_4302___1), nil, float64(0), float64(0)
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
												seq__4261_4290, chunk__4266_4291, count__4267_4292, i__4268_4293 = seq__4261_4290, chunk__4266_4291, count__4267_4292, (i__4268_4293 + float64(1))
												continue
											}
										} else {
											{
												var temp__4222__auto___4306 = cljs_core.Seq.Arity1IQ(seq__4261_4290)
												_ = temp__4222__auto___4306
												if cljs_core.Truth_(temp__4222__auto___4306) {
													{
														var seq__4261_4307___1 = temp__4222__auto___4306
														_ = seq__4261_4307___1
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4261_4307___1) {
															{
																var c__966__auto___4308 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4261_4307___1)
																_ = c__966__auto___4308
																seq__4261_4290, chunk__4266_4291, count__4267_4292, i__4268_4293 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4261_4307___1), c__966__auto___4308, cljs_core.Count.X_invoke_Arity1(c__966__auto___4308).(float64), float64(0)
																continue
															}
														} else {
															{
																var day_4309 = cljs_core.First.X_invoke_Arity1(seq__4261_4307___1)
																_ = day_4309
																{
																	var seq__4262_4310 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																	var chunk__4263_4311 interface{} = nil
																	var count__4264_4312 = float64(0)
																	var i__4265_4313 = float64(0)
																	_, _, _, _ = seq__4262_4310, chunk__4263_4311, count__4264_4312, i__4265_4313
																	for {
																		if i__4265_4313 < count__4264_4312 {
																			{
																				var hour_4314 = cljs_core.Decorate_(chunk__4263_4311).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4265_4313)
																				_ = hour_4314
																				{
																					var s_4315 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(month_4289)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(day_4309)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(hour_4314)).(string) + ":14:15.666-06:00\"")
																					_ = s_4315
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4315), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4315)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__4262_4310, chunk__4263_4311, count__4264_4312, i__4265_4313 = seq__4262_4310, chunk__4263_4311, count__4264_4312, (i__4265_4313 + float64(1))
																				continue
																			}
																		} else {
																			{
																				var temp__4222__auto___4316___1 = cljs_core.Seq.Arity1IQ(seq__4262_4310)
																				_ = temp__4222__auto___4316___1
																				if cljs_core.Truth_(temp__4222__auto___4316___1) {
																					{
																						var seq__4262_4317___1 = temp__4222__auto___4316___1
																						_ = seq__4262_4317___1
																						if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4262_4317___1) {
																							{
																								var c__966__auto___4318 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4262_4317___1)
																								_ = c__966__auto___4318
																								seq__4262_4310, chunk__4263_4311, count__4264_4312, i__4265_4313 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4262_4317___1), c__966__auto___4318, cljs_core.Count.X_invoke_Arity1(c__966__auto___4318).(float64), float64(0)
																								continue
																							}
																						} else {
																							{
																								var hour_4319 = cljs_core.First.X_invoke_Arity1(seq__4262_4317___1)
																								_ = hour_4319
																								{
																									var s_4320 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(month_4289)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(day_4309)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(hour_4319)).(string) + ":14:15.666-06:00\"")
																									_ = s_4320
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4320), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4320)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__4262_4310, chunk__4263_4311, count__4264_4312, i__4265_4313 = cljs_core.Next.Arity1IQ(seq__4262_4317___1), nil, float64(0), float64(0)
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
																seq__4261_4290, chunk__4266_4291, count__4267_4292, i__4268_4293 = cljs_core.Next.Arity1IQ(seq__4261_4307___1), nil, float64(0), float64(0)
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
								seq__4245_4285, chunk__4258_4286, count__4259_4287, i__4260_4288 = seq__4245_4285, chunk__4258_4286, count__4259_4287, (i__4260_4288 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___4321 = cljs_core.Seq.Arity1IQ(seq__4245_4285)
								_ = temp__4222__auto___4321
								if cljs_core.Truth_(temp__4222__auto___4321) {
									{
										var seq__4245_4322___1 = temp__4222__auto___4321
										_ = seq__4245_4322___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4245_4322___1) {
											{
												var c__966__auto___4323 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4245_4322___1)
												_ = c__966__auto___4323
												seq__4245_4285, chunk__4258_4286, count__4259_4287, i__4260_4288 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4245_4322___1), c__966__auto___4323, cljs_core.Count.X_invoke_Arity1(c__966__auto___4323).(float64), float64(0)
												continue
											}
										} else {
											{
												var month_4324 = cljs_core.First.X_invoke_Arity1(seq__4245_4322___1)
												_ = month_4324
												{
													var seq__4246_4325 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
													var chunk__4251_4326 interface{} = nil
													var count__4252_4327 = float64(0)
													var i__4253_4328 = float64(0)
													_, _, _, _ = seq__4246_4325, chunk__4251_4326, count__4252_4327, i__4253_4328
													for {
														if i__4253_4328 < count__4252_4327 {
															{
																var day_4329 = cljs_core.Decorate_(chunk__4251_4326).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4253_4328)
																_ = day_4329
																{
																	var seq__4254_4330 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																	var chunk__4255_4331 interface{} = nil
																	var count__4256_4332 = float64(0)
																	var i__4257_4333 = float64(0)
																	_, _, _, _ = seq__4254_4330, chunk__4255_4331, count__4256_4332, i__4257_4333
																	for {
																		if i__4257_4333 < count__4256_4332 {
																			{
																				var hour_4334 = cljs_core.Decorate_(chunk__4255_4331).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4257_4333)
																				_ = hour_4334
																				{
																					var s_4335 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(month_4324)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(day_4329)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(hour_4334)).(string) + ":14:15.666-06:00\"")
																					_ = s_4335
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4335), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4335)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__4254_4330, chunk__4255_4331, count__4256_4332, i__4257_4333 = seq__4254_4330, chunk__4255_4331, count__4256_4332, (i__4257_4333 + float64(1))
																				continue
																			}
																		} else {
																			{
																				var temp__4222__auto___4336___1 = cljs_core.Seq.Arity1IQ(seq__4254_4330)
																				_ = temp__4222__auto___4336___1
																				if cljs_core.Truth_(temp__4222__auto___4336___1) {
																					{
																						var seq__4254_4337___1 = temp__4222__auto___4336___1
																						_ = seq__4254_4337___1
																						if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4254_4337___1) {
																							{
																								var c__966__auto___4338 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4254_4337___1)
																								_ = c__966__auto___4338
																								seq__4254_4330, chunk__4255_4331, count__4256_4332, i__4257_4333 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4254_4337___1), c__966__auto___4338, cljs_core.Count.X_invoke_Arity1(c__966__auto___4338).(float64), float64(0)
																								continue
																							}
																						} else {
																							{
																								var hour_4339 = cljs_core.First.X_invoke_Arity1(seq__4254_4337___1)
																								_ = hour_4339
																								{
																									var s_4340 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(month_4324)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(day_4329)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(hour_4339)).(string) + ":14:15.666-06:00\"")
																									_ = s_4340
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4340), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4340)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__4254_4330, chunk__4255_4331, count__4256_4332, i__4257_4333 = cljs_core.Next.Arity1IQ(seq__4254_4337___1), nil, float64(0), float64(0)
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
																seq__4246_4325, chunk__4251_4326, count__4252_4327, i__4253_4328 = seq__4246_4325, chunk__4251_4326, count__4252_4327, (i__4253_4328 + float64(1))
																continue
															}
														} else {
															{
																var temp__4222__auto___4341___1 = cljs_core.Seq.Arity1IQ(seq__4246_4325)
																_ = temp__4222__auto___4341___1
																if cljs_core.Truth_(temp__4222__auto___4341___1) {
																	{
																		var seq__4246_4342___1 = temp__4222__auto___4341___1
																		_ = seq__4246_4342___1
																		if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4246_4342___1) {
																			{
																				var c__966__auto___4343 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4246_4342___1)
																				_ = c__966__auto___4343
																				seq__4246_4325, chunk__4251_4326, count__4252_4327, i__4253_4328 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4246_4342___1), c__966__auto___4343, cljs_core.Count.X_invoke_Arity1(c__966__auto___4343).(float64), float64(0)
																				continue
																			}
																		} else {
																			{
																				var day_4344 = cljs_core.First.X_invoke_Arity1(seq__4246_4342___1)
																				_ = day_4344
																				{
																					var seq__4247_4345 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																					var chunk__4248_4346 interface{} = nil
																					var count__4249_4347 = float64(0)
																					var i__4250_4348 = float64(0)
																					_, _, _, _ = seq__4247_4345, chunk__4248_4346, count__4249_4347, i__4250_4348
																					for {
																						if i__4250_4348 < count__4249_4347 {
																							{
																								var hour_4349 = cljs_core.Decorate_(chunk__4248_4346).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4250_4348)
																								_ = hour_4349
																								{
																									var s_4350 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(month_4324)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(day_4344)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(hour_4349)).(string) + ":14:15.666-06:00\"")
																									_ = s_4350
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4350), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4350)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__4247_4345, chunk__4248_4346, count__4249_4347, i__4250_4348 = seq__4247_4345, chunk__4248_4346, count__4249_4347, (i__4250_4348 + float64(1))
																								continue
																							}
																						} else {
																							{
																								var temp__4222__auto___4351___2 = cljs_core.Seq.Arity1IQ(seq__4247_4345)
																								_ = temp__4222__auto___4351___2
																								if cljs_core.Truth_(temp__4222__auto___4351___2) {
																									{
																										var seq__4247_4352___1 = temp__4222__auto___4351___2
																										_ = seq__4247_4352___1
																										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4247_4352___1) {
																											{
																												var c__966__auto___4353 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4247_4352___1)
																												_ = c__966__auto___4353
																												seq__4247_4345, chunk__4248_4346, count__4249_4347, i__4250_4348 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4247_4352___1), c__966__auto___4353, cljs_core.Count.X_invoke_Arity1(c__966__auto___4353).(float64), float64(0)
																												continue
																											}
																										} else {
																											{
																												var hour_4354 = cljs_core.First.X_invoke_Arity1(seq__4247_4352___1)
																												_ = hour_4354
																												{
																													var s_4355 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(month_4324)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(day_4344)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_4283.X_invoke_Arity1(hour_4354)).(string) + ":14:15.666-06:00\"")
																													_ = s_4355
																													if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_4355), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_4355)})).(string)), "ValueOf", []interface{}{})) {
																													} else {
																														panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																													}
																												}
																												seq__4247_4345, chunk__4248_4346, count__4249_4347, i__4250_4348 = cljs_core.Next.Arity1IQ(seq__4247_4352___1), nil, float64(0), float64(0)
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
																				seq__4246_4325, chunk__4251_4326, count__4252_4327, i__4253_4328 = cljs_core.Next.Arity1IQ(seq__4246_4342___1), nil, float64(0), float64(0)
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
												seq__4245_4285, chunk__4258_4286, count__4259_4287, i__4260_4288 = cljs_core.Next.Arity1IQ(seq__4245_4322___1), nil, float64(0), float64(0)
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
				var insts_4356 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01T00\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01T00:00:00.000\"")}, nil})
				_ = insts_4356
				if cljs_core.Truth_(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_EQ_, cljs_core.Map_.X_invoke_Arity2(func(G__4357 *cljs_core.AFn, insts_4356 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__4357, 1, func(p1__4206_SHARP_ interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(p1__4206_SHARP_, "ValueOf", []interface{}{})
					})
				}(&cljs_core.AFn{}, insts_4356), insts_4356).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (apply = (map (fn* [p1__4206#] (.valueOf p1__4206#)) insts))")}))
				}
			}
			{
				var u_4358 = cljs_reader.Read_string.X_invoke_Arity1("#uuid \"550e8400-e29b-41d4-a716-446655440000\"")
				_ = u_4358
				if cljs_core.X_EQ_.Arity2IIB(u_4358, cljs_reader.Read_string.X_invoke_Arity1("#uuid \"550e8400-e29b-41d4-a716-446655440000\"")) {
				} else {
					panic((&js.Error{("Assert failed: (= u (reader/read-string \"#uuid \\\"550e8400-e29b-41d4-a716-446655440000\\\"\"))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(u_4358, cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{u_4358})).(string))) {
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
			cljs_reader.Register_default_tag_parser_BANG_.X_invoke_Arity1(func(G__4359 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4359, 2, func(tag interface{}, val interface{}) interface{} {
					return val
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), cljs_reader.Read_string.X_invoke_Arity1("#a.b/c [1 2]")) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (reader/read-string \"#a.b/c [1 2]\"))")}))
			}
			{
				var seq__4273_4360 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"\"abc \\ua\"", "\"abc \\x0z  ...etc\"", "\"abc \\u0g00 ..etc\""}, nil}))
				var chunk__4274_4361 interface{} = nil
				var count__4275_4362 = float64(0)
				var i__4276_4363 = float64(0)
				_, _, _, _ = seq__4273_4360, chunk__4274_4361, count__4275_4362, i__4276_4363
				for {
					if i__4276_4363 < count__4275_4362 {
						{
							var unicode_error_4364 = cljs_core.Decorate_(chunk__4274_4361).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4276_4363)
							_ = unicode_error_4364
							{
								var r_4365 = func() (return__4366 *cljs_core.CljsCoreKeyword) {
									defer func() {
										if e4277 := recover(); e4277 != nil {
											if cljs_core.Value_(e4277).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
												{
													var e = e4277
													_ = e
													return__4366 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
												}
											} else {
												panic(e4277)

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
							seq__4273_4360, chunk__4274_4361, count__4275_4362, i__4276_4363 = seq__4273_4360, chunk__4274_4361, count__4275_4362, (i__4276_4363 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___4367 = cljs_core.Seq.Arity1IQ(seq__4273_4360)
							_ = temp__4222__auto___4367
							if cljs_core.Truth_(temp__4222__auto___4367) {
								{
									var seq__4273_4368___1 = temp__4222__auto___4367
									_ = seq__4273_4368___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4273_4368___1) {
										{
											var c__966__auto___4369 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4273_4368___1)
											_ = c__966__auto___4369
											seq__4273_4360, chunk__4274_4361, count__4275_4362, i__4276_4363 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4273_4368___1), c__966__auto___4369, cljs_core.Count.X_invoke_Arity1(c__966__auto___4369).(float64), float64(0)
											continue
										}
									} else {
										{
											var unicode_error_4370 = cljs_core.First.X_invoke_Arity1(seq__4273_4368___1)
											_ = unicode_error_4370
											{
												var r_4371 = func() (return__4372 *cljs_core.CljsCoreKeyword) {
													defer func() {
														if e4278 := recover(); e4278 != nil {
															if cljs_core.Value_(e4278).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
																{
																	var e = e4278
																	_ = e
																	return__4372 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
																}
															} else {
																panic(e4278)

															}
														}
													}()
													{
														cljs_reader.Read_string.X_invoke_Arity1(unicode_error_4370)
														return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed-to-throw", Fqn: "failed-to-throw", X_hash: float64(-1638969220)})
													}
												}()
												_ = r_4371
												if cljs_core.X_EQ_.Arity2IIB(r_4371, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
												} else {
													panic((&js.Error{("Assert failed: " + cljs_core.Str.X_invoke_Arity1(("Failed to throw reader error for: " + cljs_core.Str.X_invoke_Arity1(unicode_error_4370).(string))).(string) + "\n(= r :ok)")}))
												}
											}
											seq__4273_4360, chunk__4274_4361, count__4275_4362, i__4276_4363 = cljs_core.Next.Arity1IQ(seq__4273_4368___1), nil, float64(0), float64(0)
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

func (this__772__auto__ *CljsReader_testR) X_lookup_Arity3(k4188 interface{}, else__773__auto__ interface{}) interface{} {
	{
		var G__4191 = func() interface{} {
			if cljs_core.Value_(k4188).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k4188), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4191
		switch G__4191 {
		case "b":
			return this__772__auto__.B

		case "a":
			return this__772__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__772__auto__.X__extmap, k4188, else__773__auto__)

		}
	}
}

func (_ *CljsReader_testR) CljsCoreIPrintWithWriter__() {}
func (this__786__auto__ *CljsReader_testR) X_pr_writer_Arity3(writer__787__auto__ interface{}, opts__788__auto__ interface{}) interface{} {
	{
		var pr_pair__789__auto__ = func(G__4374 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__4374, 3, func(keyval__790__auto__ interface{}, ___791__auto__ interface{}, ___791__auto_____1 interface{}) interface{} {
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
func (_ *CljsReader_testR) CljsCoreIMap__()    {}
func (this__781__auto__ *CljsReader_testR) X_dissoc_Arity2(k__782__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__782__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__781__auto__), this__781__auto__.X__meta), k__782__auto__)
	} else {
		return (&CljsReader_testR{this__781__auto__.A, this__781__auto__.B, this__781__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__781__auto__.X__extmap, k__782__auto__)), nil})
	}
}

func (_ *CljsReader_testR) CljsCoreIAssociative__() {}
func (this__777__auto__ *CljsReader_testR) X_assoc_Arity3(k__778__auto__ interface{}, G__4187 interface{}) interface{} {
	{
		var pred__4199 = cljs_core.Keyword_identical_QMARK_
		var expr__4200 = k__778__auto__
		_, _ = pred__4199, expr__4200
		if cljs_core.Truth_(func() interface{} {
			var G__4202 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__4203 = expr__4200
			_, _ = G__4202, G__4203
			return pred__4199.X_invoke_Arity2(G__4202, G__4203)
		}()) {
			return (&CljsReader_testR{G__4187, this__777__auto__.B, this__777__auto__.X__meta, this__777__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__4204 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__4205 = expr__4200
				_, _ = G__4204, G__4205
				return pred__4199.X_invoke_Arity2(G__4204, G__4205)
			}()) {
				return (&CljsReader_testR{this__777__auto__.A, G__4187, this__777__auto__.X__meta, this__777__auto__.X__extmap, nil})
			} else {
				return (&CljsReader_testR{this__777__auto__.A, this__777__auto__.B, this__777__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__777__auto__.X__extmap, k__778__auto__, G__4187), nil})
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
func (this__769__auto__ *CljsReader_testR) X_with_meta_Arity2(G__4187 interface{}) interface{} {
	return (&CljsReader_testR{this__769__auto__.A, this__769__auto__.B, G__4187, this__769__auto__.X__extmap, this__769__auto__.X__hash})
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
