// Compiled by ClojureScript to Go 0.0-2371
// cljs.reader-test

package reader_test

import (
	"reflect"
	"testing"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	cljs_reader "github.com/hraberg/cljs2go/cljs/reader"
	"github.com/hraberg/cljs2go/js"
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: "foo", Name: "bar", Fqn: "foo/bar", X_hash: float64(-1386151538)}), cljs_reader.Read_string.X_invoke_Arity1(":foo/bar")) {
			} else {
				panic((&js.Error{("Assert failed: (= :foo/bar (reader/read-string \":foo/bar\"))")}))
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
			if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__261 interface{}) {
				defer func() {
					if e221 := recover(); e221 != nil {
						if cljs_core.Value_(e221).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
							{
								var e = e221
								_ = e
								return__261 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "threw", Fqn: "threw", X_hash: float64(17630075)})
							}
						} else {
							panic(e221)

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
			if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__262 interface{}) {
				defer func() {
					if e222 := recover(); e222 != nil {
						if cljs_core.Value_(e222).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
							{
								var e = e222
								_ = e
								return__262 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "threw", Fqn: "threw", X_hash: float64(17630075)})
							}
						} else {
							panic(e222)

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
				var est_inst_263 = cljs_reader.Read_string.X_invoke_Arity1("#inst \"2010-11-12T13:14:15.666-05:00\"")
				var utc_inst_264 = cljs_reader.Read_string.X_invoke_Arity1("#inst \"2010-11-12T18:14:15.666-00:00\"")
				var pad_265 = func(G__266 *cljs_core.AFn, est_inst_263 interface{}, utc_inst_264 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__266, 1, func(n interface{}) interface{} {
						if n.(float64) < float64(10) {
							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
						} else {
							return n
						}
					})
				}(&cljs_core.AFn{}, est_inst_263, utc_inst_264)
				_, _, _ = est_inst_263, utc_inst_264, pad_265
				if cljs_core.X_EQ_.Arity2IIB((&js.Date{"2010-11-12T13:14:15.666-05:00"}).ValueOf(), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_263, "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf (js/Date. \"2010-11-12T13:14:15.666-05:00\")) (.valueOf est-inst))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_263, "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{est_inst_263})).(string)), "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf est-inst) (.valueOf (reader/read-string (pr-str est-inst))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(est_inst_263, "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(utc_inst_264, "ValueOf", []interface{}{})) {
				} else {
					panic((&js.Error{("Assert failed: (= (.valueOf est-inst) (.valueOf utc-inst))")}))
				}
				{
					var seq__223_267 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(13)).(*cljs_core.CljsCoreRange))
					var chunk__236_268 interface{} = nil
					var count__237_269 = float64(0)
					var i__238_270 = float64(0)
					_, _, _, _ = seq__223_267, chunk__236_268, count__237_269, i__238_270
					for {
						if i__238_270 < count__237_269 {
							{
								var month_271 = cljs_core.Decorate_(chunk__236_268).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__238_270)
								_ = month_271
								{
									var seq__239_272 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
									var chunk__244_273 interface{} = nil
									var count__245_274 = float64(0)
									var i__246_275 = float64(0)
									_, _, _, _ = seq__239_272, chunk__244_273, count__245_274, i__246_275
									for {
										if i__246_275 < count__245_274 {
											{
												var day_276 = cljs_core.Decorate_(chunk__244_273).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__246_275)
												_ = day_276
												{
													var seq__247_277 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
													var chunk__248_278 interface{} = nil
													var count__249_279 = float64(0)
													var i__250_280 = float64(0)
													_, _, _, _ = seq__247_277, chunk__248_278, count__249_279, i__250_280
													for {
														if i__250_280 < count__249_279 {
															{
																var hour_281 = cljs_core.Decorate_(chunk__248_278).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__250_280)
																_ = hour_281
																{
																	var s_282 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(month_271)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(day_276)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(hour_281)).(string) + ":14:15.666-06:00\"")
																	_ = s_282
																	if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_282), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_282)})).(string)), "ValueOf", []interface{}{})) {
																	} else {
																		panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																	}
																}
																seq__247_277, chunk__248_278, count__249_279, i__250_280 = seq__247_277, chunk__248_278, count__249_279, (i__250_280 + float64(1))
																continue
															}
														} else {
															{
																var temp__4388__auto___283 = cljs_core.Seq.Arity1IQ(seq__247_277)
																_ = temp__4388__auto___283
																if cljs_core.Truth_(temp__4388__auto___283) {
																	{
																		var seq__247_284___1 = temp__4388__auto___283
																		_ = seq__247_284___1
																		if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__247_284___1) {
																			{
																				var c__971__auto___285 = cljs_core.Chunk_first.X_invoke_Arity1(seq__247_284___1)
																				_ = c__971__auto___285
																				seq__247_277, chunk__248_278, count__249_279, i__250_280 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__247_284___1), c__971__auto___285, cljs_core.Count.X_invoke_Arity1(c__971__auto___285).(float64), float64(0)
																				continue
																			}
																		} else {
																			{
																				var hour_286 = cljs_core.First.X_invoke_Arity1(seq__247_284___1)
																				_ = hour_286
																				{
																					var s_287 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(month_271)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(day_276)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(hour_286)).(string) + ":14:15.666-06:00\"")
																					_ = s_287
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_287), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_287)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__247_277, chunk__248_278, count__249_279, i__250_280 = cljs_core.Next.Arity1IQ(seq__247_284___1), nil, float64(0), float64(0)
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
												seq__239_272, chunk__244_273, count__245_274, i__246_275 = seq__239_272, chunk__244_273, count__245_274, (i__246_275 + float64(1))
												continue
											}
										} else {
											{
												var temp__4388__auto___288 = cljs_core.Seq.Arity1IQ(seq__239_272)
												_ = temp__4388__auto___288
												if cljs_core.Truth_(temp__4388__auto___288) {
													{
														var seq__239_289___1 = temp__4388__auto___288
														_ = seq__239_289___1
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__239_289___1) {
															{
																var c__971__auto___290 = cljs_core.Chunk_first.X_invoke_Arity1(seq__239_289___1)
																_ = c__971__auto___290
																seq__239_272, chunk__244_273, count__245_274, i__246_275 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__239_289___1), c__971__auto___290, cljs_core.Count.X_invoke_Arity1(c__971__auto___290).(float64), float64(0)
																continue
															}
														} else {
															{
																var day_291 = cljs_core.First.X_invoke_Arity1(seq__239_289___1)
																_ = day_291
																{
																	var seq__240_292 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																	var chunk__241_293 interface{} = nil
																	var count__242_294 = float64(0)
																	var i__243_295 = float64(0)
																	_, _, _, _ = seq__240_292, chunk__241_293, count__242_294, i__243_295
																	for {
																		if i__243_295 < count__242_294 {
																			{
																				var hour_296 = cljs_core.Decorate_(chunk__241_293).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__243_295)
																				_ = hour_296
																				{
																					var s_297 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(month_271)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(day_291)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(hour_296)).(string) + ":14:15.666-06:00\"")
																					_ = s_297
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_297), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_297)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__240_292, chunk__241_293, count__242_294, i__243_295 = seq__240_292, chunk__241_293, count__242_294, (i__243_295 + float64(1))
																				continue
																			}
																		} else {
																			{
																				var temp__4388__auto___298___1 = cljs_core.Seq.Arity1IQ(seq__240_292)
																				_ = temp__4388__auto___298___1
																				if cljs_core.Truth_(temp__4388__auto___298___1) {
																					{
																						var seq__240_299___1 = temp__4388__auto___298___1
																						_ = seq__240_299___1
																						if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__240_299___1) {
																							{
																								var c__971__auto___300 = cljs_core.Chunk_first.X_invoke_Arity1(seq__240_299___1)
																								_ = c__971__auto___300
																								seq__240_292, chunk__241_293, count__242_294, i__243_295 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__240_299___1), c__971__auto___300, cljs_core.Count.X_invoke_Arity1(c__971__auto___300).(float64), float64(0)
																								continue
																							}
																						} else {
																							{
																								var hour_301 = cljs_core.First.X_invoke_Arity1(seq__240_299___1)
																								_ = hour_301
																								{
																									var s_302 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(month_271)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(day_291)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(hour_301)).(string) + ":14:15.666-06:00\"")
																									_ = s_302
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_302), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_302)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__240_292, chunk__241_293, count__242_294, i__243_295 = cljs_core.Next.Arity1IQ(seq__240_299___1), nil, float64(0), float64(0)
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
																seq__239_272, chunk__244_273, count__245_274, i__246_275 = cljs_core.Next.Arity1IQ(seq__239_289___1), nil, float64(0), float64(0)
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
								seq__223_267, chunk__236_268, count__237_269, i__238_270 = seq__223_267, chunk__236_268, count__237_269, (i__238_270 + float64(1))
								continue
							}
						} else {
							{
								var temp__4388__auto___303 = cljs_core.Seq.Arity1IQ(seq__223_267)
								_ = temp__4388__auto___303
								if cljs_core.Truth_(temp__4388__auto___303) {
									{
										var seq__223_304___1 = temp__4388__auto___303
										_ = seq__223_304___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__223_304___1) {
											{
												var c__971__auto___305 = cljs_core.Chunk_first.X_invoke_Arity1(seq__223_304___1)
												_ = c__971__auto___305
												seq__223_267, chunk__236_268, count__237_269, i__238_270 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__223_304___1), c__971__auto___305, cljs_core.Count.X_invoke_Arity1(c__971__auto___305).(float64), float64(0)
												continue
											}
										} else {
											{
												var month_306 = cljs_core.First.X_invoke_Arity1(seq__223_304___1)
												_ = month_306
												{
													var seq__224_307 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
													var chunk__229_308 interface{} = nil
													var count__230_309 = float64(0)
													var i__231_310 = float64(0)
													_, _, _, _ = seq__224_307, chunk__229_308, count__230_309, i__231_310
													for {
														if i__231_310 < count__230_309 {
															{
																var day_311 = cljs_core.Decorate_(chunk__229_308).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__231_310)
																_ = day_311
																{
																	var seq__232_312 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																	var chunk__233_313 interface{} = nil
																	var count__234_314 = float64(0)
																	var i__235_315 = float64(0)
																	_, _, _, _ = seq__232_312, chunk__233_313, count__234_314, i__235_315
																	for {
																		if i__235_315 < count__234_314 {
																			{
																				var hour_316 = cljs_core.Decorate_(chunk__233_313).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__235_315)
																				_ = hour_316
																				{
																					var s_317 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(month_306)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(day_311)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(hour_316)).(string) + ":14:15.666-06:00\"")
																					_ = s_317
																					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_317), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_317)})).(string)), "ValueOf", []interface{}{})) {
																					} else {
																						panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																					}
																				}
																				seq__232_312, chunk__233_313, count__234_314, i__235_315 = seq__232_312, chunk__233_313, count__234_314, (i__235_315 + float64(1))
																				continue
																			}
																		} else {
																			{
																				var temp__4388__auto___318___1 = cljs_core.Seq.Arity1IQ(seq__232_312)
																				_ = temp__4388__auto___318___1
																				if cljs_core.Truth_(temp__4388__auto___318___1) {
																					{
																						var seq__232_319___1 = temp__4388__auto___318___1
																						_ = seq__232_319___1
																						if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__232_319___1) {
																							{
																								var c__971__auto___320 = cljs_core.Chunk_first.X_invoke_Arity1(seq__232_319___1)
																								_ = c__971__auto___320
																								seq__232_312, chunk__233_313, count__234_314, i__235_315 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__232_319___1), c__971__auto___320, cljs_core.Count.X_invoke_Arity1(c__971__auto___320).(float64), float64(0)
																								continue
																							}
																						} else {
																							{
																								var hour_321 = cljs_core.First.X_invoke_Arity1(seq__232_319___1)
																								_ = hour_321
																								{
																									var s_322 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(month_306)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(day_311)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(hour_321)).(string) + ":14:15.666-06:00\"")
																									_ = s_322
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_322), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_322)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__232_312, chunk__233_313, count__234_314, i__235_315 = cljs_core.Next.Arity1IQ(seq__232_319___1), nil, float64(0), float64(0)
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
																seq__224_307, chunk__229_308, count__230_309, i__231_310 = seq__224_307, chunk__229_308, count__230_309, (i__231_310 + float64(1))
																continue
															}
														} else {
															{
																var temp__4388__auto___323___1 = cljs_core.Seq.Arity1IQ(seq__224_307)
																_ = temp__4388__auto___323___1
																if cljs_core.Truth_(temp__4388__auto___323___1) {
																	{
																		var seq__224_324___1 = temp__4388__auto___323___1
																		_ = seq__224_324___1
																		if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__224_324___1) {
																			{
																				var c__971__auto___325 = cljs_core.Chunk_first.X_invoke_Arity1(seq__224_324___1)
																				_ = c__971__auto___325
																				seq__224_307, chunk__229_308, count__230_309, i__231_310 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__224_324___1), c__971__auto___325, cljs_core.Count.X_invoke_Arity1(c__971__auto___325).(float64), float64(0)
																				continue
																			}
																		} else {
																			{
																				var day_326 = cljs_core.First.X_invoke_Arity1(seq__224_324___1)
																				_ = day_326
																				{
																					var seq__225_327 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																					var chunk__226_328 interface{} = nil
																					var count__227_329 = float64(0)
																					var i__228_330 = float64(0)
																					_, _, _, _ = seq__225_327, chunk__226_328, count__227_329, i__228_330
																					for {
																						if i__228_330 < count__227_329 {
																							{
																								var hour_331 = cljs_core.Decorate_(chunk__226_328).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__228_330)
																								_ = hour_331
																								{
																									var s_332 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(month_306)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(day_326)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(hour_331)).(string) + ":14:15.666-06:00\"")
																									_ = s_332
																									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_332), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_332)})).(string)), "ValueOf", []interface{}{})) {
																									} else {
																										panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																									}
																								}
																								seq__225_327, chunk__226_328, count__227_329, i__228_330 = seq__225_327, chunk__226_328, count__227_329, (i__228_330 + float64(1))
																								continue
																							}
																						} else {
																							{
																								var temp__4388__auto___333___2 = cljs_core.Seq.Arity1IQ(seq__225_327)
																								_ = temp__4388__auto___333___2
																								if cljs_core.Truth_(temp__4388__auto___333___2) {
																									{
																										var seq__225_334___1 = temp__4388__auto___333___2
																										_ = seq__225_334___1
																										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__225_334___1) {
																											{
																												var c__971__auto___335 = cljs_core.Chunk_first.X_invoke_Arity1(seq__225_334___1)
																												_ = c__971__auto___335
																												seq__225_327, chunk__226_328, count__227_329, i__228_330 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__225_334___1), c__971__auto___335, cljs_core.Count.X_invoke_Arity1(c__971__auto___335).(float64), float64(0)
																												continue
																											}
																										} else {
																											{
																												var hour_336 = cljs_core.First.X_invoke_Arity1(seq__225_334___1)
																												_ = hour_336
																												{
																													var s_337 = ("#inst \"2010-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(month_306)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(day_326)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_265.X_invoke_Arity1(hour_336)).(string) + ":14:15.666-06:00\"")
																													_ = s_337
																													if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(s_337), "ValueOf", []interface{}{}), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_reader.Read_string.X_invoke_Arity1(s_337)})).(string)), "ValueOf", []interface{}{})) {
																													} else {
																														panic((&js.Error{("Assert failed: (= (-> s reader/read-string .valueOf) (-> s reader/read-string pr-str reader/read-string .valueOf))")}))
																													}
																												}
																												seq__225_327, chunk__226_328, count__227_329, i__228_330 = cljs_core.Next.Arity1IQ(seq__225_334___1), nil, float64(0), float64(0)
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
																				seq__224_307, chunk__229_308, count__230_309, i__231_310 = cljs_core.Next.Arity1IQ(seq__224_324___1), nil, float64(0), float64(0)
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
												seq__223_267, chunk__236_268, count__237_269, i__238_270 = cljs_core.Next.Arity1IQ(seq__223_304___1), nil, float64(0), float64(0)
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
				var insts_338 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01T00\""), cljs_reader.Read_string.X_invoke_Arity1("#inst \"2012-01-01T00:00:00.000\"")}, nil})
				_ = insts_338
				if cljs_core.Truth_(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_EQ_, cljs_core.Map_.X_invoke_Arity2(func(G__339 *cljs_core.AFn, insts_338 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__339, 1, func(p1__180_SHARP_ interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(p1__180_SHARP_, "ValueOf", []interface{}{})
					})
				}(&cljs_core.AFn{}, insts_338), insts_338).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (apply = (map (fn* [p1__180#] (.valueOf p1__180#)) insts))")}))
				}
			}
			{
				var u_340 = cljs_reader.Read_string.X_invoke_Arity1("#uuid \"550e8400-e29b-41d4-a716-446655440000\"")
				_ = u_340
				if cljs_core.X_EQ_.Arity2IIB(u_340, cljs_reader.Read_string.X_invoke_Arity1("#uuid \"550e8400-e29b-41d4-a716-446655440000\"")) {
				} else {
					panic((&js.Error{("Assert failed: (= u (reader/read-string \"#uuid \\\"550e8400-e29b-41d4-a716-446655440000\\\"\"))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(u_340, cljs_reader.Read_string.X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{u_340})).(string))) {
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
			cljs_reader.Register_default_tag_parser_BANG_.X_invoke_Arity1(func(G__341 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__341, 2, func(tag interface{}, val interface{}) interface{} {
					return val
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), cljs_reader.Read_string.X_invoke_Arity1("#a.b/c [1 2]")) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (reader/read-string \"#a.b/c [1 2]\"))")}))
			}
			{
				var seq__251_342 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(18), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"\u0627\u062E\u062A\u0628\u0627\u0631", "\u0E17\u0E14\u0E2A\u0E2D\u0E1A", "\u3053\u3093\u306B\u3061\u306F", "\u4F60\u597D", "\u05D0\u05B7 \u05D2\u05D5\u05D8 \u05D9\u05D0\u05B8\u05E8", "cze\u015B\u0107", "\u043F\u0440\u0438\u0432\u0435\u0442", (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "\u0E17\u0E14\u0E2A\u0E2D\u0E1A", Str: "\u0E17\u0E14\u0E2A\u0E2D\u0E1A", X_hash: float64(593353297), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "\u3053\u3093\u306B\u3061\u306F", Str: "\u3053\u3093\u306B\u3061\u306F", X_hash: float64(1271843387), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "\u4F60\u597D", Str: "\u4F60\u597D", X_hash: float64(1250039035), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "cze\u015B\u0107", Str: "cze\u015B\u0107", X_hash: float64(1605608982), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "\u043F\u0440\u0438\u0432\u0435\u0442", Str: "\u043F\u0440\u0438\u0432\u0435\u0442", X_hash: float64(-93177211), X_meta: nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "\u0E17\u0E14\u0E2A\u0E2D\u0E1A", Fqn: "\u0E17\u0E14\u0E2A\u0E2D\u0E1A", X_hash: float64(-1047178230)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "\u3053\u3093\u306B\u3061\u306F", Fqn: "\u3053\u3093\u306B\u3061\u306F", X_hash: float64(-368688140)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "\u4F60\u597D", Fqn: "\u4F60\u597D", X_hash: float64(-390492492)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "cze\u015B\u0107", Fqn: "cze\u015B\u0107", X_hash: float64(-34922545)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "\u043F\u0440\u0438\u0432\u0435\u0442", Fqn: "\u043F\u0440\u0438\u0432\u0435\u0442", X_hash: float64(-1733708738)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "\u043F\u0440\u0438\u0432\u0435\u0442", Fqn: "\u043F\u0440\u0438\u0432\u0435\u0442", X_hash: float64(-1733708738)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ru", Fqn: "ru", X_hash: float64(-1755311210)}), "\u4F60\u597D", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "cn", Fqn: "cn", X_hash: float64(457269822)})}, nil})}, nil}))
				var chunk__252_343 interface{} = nil
				var count__253_344 = float64(0)
				var i__254_345 = float64(0)
				_, _, _, _ = seq__251_342, chunk__252_343, count__253_344, i__254_345
				for {
					if i__254_345 < count__253_344 {
						{
							var unicode_346 = cljs_core.Decorate_(chunk__252_343).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__254_345)
							_ = unicode_346
							{
								var input_347 = cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{unicode_346})).(string)
								var read_348 = cljs_reader.Read_string.X_invoke_Arity1(input_347)
								_, _ = input_347, read_348
								if cljs_core.X_EQ_.Arity2IIB(unicode_346, read_348) {
								} else {
									panic((&js.Error{("Assert failed: " + cljs_core.Str.X_invoke_Arity1(("Failed to read-string \"" + cljs_core.Str.X_invoke_Arity1(unicode_346).(string) + "\" from: " + cljs_core.Str.X_invoke_Arity1(input_347).(string))).(string) + "\n(= unicode read)")}))
								}
							}
							seq__251_342, chunk__252_343, count__253_344, i__254_345 = seq__251_342, chunk__252_343, count__253_344, (i__254_345 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___349 = cljs_core.Seq.Arity1IQ(seq__251_342)
							_ = temp__4388__auto___349
							if cljs_core.Truth_(temp__4388__auto___349) {
								{
									var seq__251_350___1 = temp__4388__auto___349
									_ = seq__251_350___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__251_350___1) {
										{
											var c__971__auto___351 = cljs_core.Chunk_first.X_invoke_Arity1(seq__251_350___1)
											_ = c__971__auto___351
											seq__251_342, chunk__252_343, count__253_344, i__254_345 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__251_350___1), c__971__auto___351, cljs_core.Count.X_invoke_Arity1(c__971__auto___351).(float64), float64(0)
											continue
										}
									} else {
										{
											var unicode_352 = cljs_core.First.X_invoke_Arity1(seq__251_350___1)
											_ = unicode_352
											{
												var input_353 = cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{unicode_352})).(string)
												var read_354 = cljs_reader.Read_string.X_invoke_Arity1(input_353)
												_, _ = input_353, read_354
												if cljs_core.X_EQ_.Arity2IIB(unicode_352, read_354) {
												} else {
													panic((&js.Error{("Assert failed: " + cljs_core.Str.X_invoke_Arity1(("Failed to read-string \"" + cljs_core.Str.X_invoke_Arity1(unicode_352).(string) + "\" from: " + cljs_core.Str.X_invoke_Arity1(input_353).(string))).(string) + "\n(= unicode read)")}))
												}
											}
											seq__251_342, chunk__252_343, count__253_344, i__254_345 = cljs_core.Next.Arity1IQ(seq__251_350___1), nil, float64(0), float64(0)
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
			{
				var seq__255_355 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"\"abc \\ua\"", "\"abc \\x0z  ...etc\"", "\"abc \\u0g00 ..etc\""}, nil}))
				var chunk__256_356 interface{} = nil
				var count__257_357 = float64(0)
				var i__258_358 = float64(0)
				_, _, _, _ = seq__255_355, chunk__256_356, count__257_357, i__258_358
				for {
					if i__258_358 < count__257_357 {
						{
							var unicode_error_359 = cljs_core.Decorate_(chunk__256_356).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__258_358)
							_ = unicode_error_359
							{
								var r_360 = func() (return__361 *cljs_core.CljsCoreKeyword) {
									defer func() {
										if e259 := recover(); e259 != nil {
											if cljs_core.Value_(e259).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
												{
													var e = e259
													_ = e
													return__361 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
												}
											} else {
												panic(e259)

											}
										}
									}()
									{
										cljs_reader.Read_string.X_invoke_Arity1(unicode_error_359)
										return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed-to-throw", Fqn: "failed-to-throw", X_hash: float64(-1638969220)})
									}
								}()
								_ = r_360
								if cljs_core.X_EQ_.Arity2IIB(r_360, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
								} else {
									panic((&js.Error{("Assert failed: " + cljs_core.Str.X_invoke_Arity1(("Failed to throw reader error for: " + cljs_core.Str.X_invoke_Arity1(unicode_error_359).(string))).(string) + "\n(= r :ok)")}))
								}
							}
							seq__255_355, chunk__256_356, count__257_357, i__258_358 = seq__255_355, chunk__256_356, count__257_357, (i__258_358 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___362 = cljs_core.Seq.Arity1IQ(seq__255_355)
							_ = temp__4388__auto___362
							if cljs_core.Truth_(temp__4388__auto___362) {
								{
									var seq__255_363___1 = temp__4388__auto___362
									_ = seq__255_363___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__255_363___1) {
										{
											var c__971__auto___364 = cljs_core.Chunk_first.X_invoke_Arity1(seq__255_363___1)
											_ = c__971__auto___364
											seq__255_355, chunk__256_356, count__257_357, i__258_358 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__255_363___1), c__971__auto___364, cljs_core.Count.X_invoke_Arity1(c__971__auto___364).(float64), float64(0)
											continue
										}
									} else {
										{
											var unicode_error_365 = cljs_core.First.X_invoke_Arity1(seq__255_363___1)
											_ = unicode_error_365
											{
												var r_366 = func() (return__367 *cljs_core.CljsCoreKeyword) {
													defer func() {
														if e260 := recover(); e260 != nil {
															if cljs_core.Value_(e260).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
																{
																	var e = e260
																	_ = e
																	return__367 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
																}
															} else {
																panic(e260)

															}
														}
													}()
													{
														cljs_reader.Read_string.X_invoke_Arity1(unicode_error_365)
														return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed-to-throw", Fqn: "failed-to-throw", X_hash: float64(-1638969220)})
													}
												}()
												_ = r_366
												if cljs_core.X_EQ_.Arity2IIB(r_366, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
												} else {
													panic((&js.Error{("Assert failed: " + cljs_core.Str.X_invoke_Arity1(("Failed to throw reader error for: " + cljs_core.Str.X_invoke_Arity1(unicode_error_365).(string))).(string) + "\n(= r :ok)")}))
												}
											}
											seq__255_355, chunk__256_356, count__257_357, i__258_358 = cljs_core.Next.Arity1IQ(seq__255_363___1), nil, float64(0), float64(0)
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
			{
				var re_368 = cljs_reader.Read_string.X_invoke_Arity1("#\"\\s\\u00a1\"")
				var m_369 = cljs_core.Re_find.X_invoke_Arity2(re_368, " \u00A1   ")
				_, _ = re_368, m_369
				if cljs_core.X_EQ_.Arity2IIB(m_369, " \u00A1") {
				} else {
					panic((&js.Error{("Assert failed: (= m \" \u00A1\")")}))
				}
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
func (this__775__auto__ *CljsReader_testR) X_lookup_Arity2(k__776__auto__ interface{}) interface{} {
	return this__775__auto__.X_lookup_Arity3(k__776__auto__, nil)
}

func (this__777__auto__ *CljsReader_testR) X_lookup_Arity3(k162 interface{}, else__778__auto__ interface{}) interface{} {
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
			return this__777__auto__.B

		case "a":
			return this__777__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__777__auto__.X__extmap, k162, else__778__auto__)

		}
	}
}

func (_ *CljsReader_testR) CljsCoreIPrintWithWriter__() {}
func (this__791__auto__ *CljsReader_testR) X_pr_writer_Arity3(writer__792__auto__ interface{}, opts__793__auto__ interface{}) interface{} {
	{
		var pr_pair__794__auto__ = func(G__371 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__371, 3, func(keyval__795__auto__ interface{}, ___796__auto__ interface{}, ___796__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, cljs_core.Pr_writer, "", " ", "", opts__793__auto__, keyval__795__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__794__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, pr_pair__794__auto__, "#cljs.reader-test.R{", ", ", "}", opts__793__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__791__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__791__auto__.B}, nil})}, nil}), this__791__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsReader_testR) CljsCoreIMeta__() {}
func (this__773__auto__ *CljsReader_testR) X_meta_Arity1() interface{} {
	return this__773__auto__.X__meta
}

func (_ *CljsReader_testR) CljsCoreICloneable__() {}
func (this__769__auto__ *CljsReader_testR) X_clone_Arity1() interface{} {
	return (&CljsReader_testR{this__769__auto__.A, this__769__auto__.B, this__769__auto__.X__meta, this__769__auto__.X__extmap, this__769__auto__.X__hash})
}

func (_ *CljsReader_testR) CljsCoreICounted__() {}
func (this__779__auto__ *CljsReader_testR) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__779__auto__.X__extmap).(float64))
}

func (_ *CljsReader_testR) CljsCoreIHash__() {}
func (this__770__auto__ *CljsReader_testR) X_hash_Arity1() interface{} {
	{
		var h__593__auto__ = this__770__auto__.X__hash
		_ = h__593__auto__
		if !(cljs_core.Nil_(h__593__auto__)) {
			return h__593__auto__
		} else {
			{
				var h__593__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__770__auto__).(float64)
				_ = h__593__auto_____1
				this__770__auto__.X__hash = h__593__auto_____1

				return h__593__auto_____1
			}
		}
	}
}

func (_ *CljsReader_testR) CljsCoreIEquiv__() {}
func (this__771__auto__ *CljsReader_testR) X_equiv_Arity2(other__772__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__170__auto__ = other__772__auto__
		_ = and__170__auto__
		if cljs_core.Truth_(and__170__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__771__auto__), cljs_core.Type_.X_invoke_Arity1(other__772__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__771__auto__, other__772__auto__)))
		} else {
			return and__170__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsReader_testR) CljsCoreIRecord__() {}
func (_ *CljsReader_testR) CljsCoreIMap__()    {}
func (this__786__auto__ *CljsReader_testR) X_dissoc_Arity2(k__787__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__787__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__786__auto__), this__786__auto__.X__meta), k__787__auto__)
	} else {
		return (&CljsReader_testR{this__786__auto__.A, this__786__auto__.B, this__786__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__786__auto__.X__extmap, k__787__auto__)), nil})
	}
}

func (_ *CljsReader_testR) CljsCoreIAssociative__() {}
func (this__782__auto__ *CljsReader_testR) X_assoc_Arity3(k__783__auto__ interface{}, G__161 interface{}) interface{} {
	{
		var pred__173 = cljs_core.Keyword_identical_QMARK_
		var expr__174 = k__783__auto__
		_, _ = pred__173, expr__174
		if cljs_core.Truth_(func() interface{} {
			var G__176 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__177 = expr__174
			_, _ = G__176, G__177
			return pred__173.X_invoke_Arity2(G__176, G__177)
		}()) {
			return (&CljsReader_testR{G__161, this__782__auto__.B, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__178 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__179 = expr__174
				_, _ = G__178, G__179
				return pred__173.X_invoke_Arity2(G__178, G__179)
			}()) {
				return (&CljsReader_testR{this__782__auto__.A, G__161, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
			} else {
				return (&CljsReader_testR{this__782__auto__.A, this__782__auto__.B, this__782__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__782__auto__.X__extmap, k__783__auto__, G__161), nil})
			}
		}
	}
}

func (this__784__auto__ *CljsReader_testR) X_contains_key_QMARK__Arity2(k__785__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsReader_testR) CljsCoreISeqable__() {}
func (this__789__auto__ *CljsReader_testR) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__789__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__789__auto__.B}, nil})}, nil}), this__789__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsReader_testR) CljsCoreIWithMeta__() {}
func (this__774__auto__ *CljsReader_testR) X_with_meta_Arity2(G__161 interface{}) interface{} {
	return (&CljsReader_testR{this__774__auto__.A, this__774__auto__.B, G__161, this__774__auto__.X__extmap, this__774__auto__.X__hash})
}

func (_ *CljsReader_testR) CljsCoreICollection__() {}
func (this__780__auto__ *CljsReader_testR) X_conj_Arity2(entry__781__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__781__auto__) {
		return this__780__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__780__auto__, entry__781__auto__)
	}
}

var X__GT_R *cljs_core.AFn

var Map__GT_R *cljs_core.AFn

var Test_reader *cljs_core.AFn

func Test_runner(t *testing.T) {
	assert.Equal(t, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), Test_reader.X_invoke_Arity0().(*cljs_core.CljsCoreKeyword))
}
