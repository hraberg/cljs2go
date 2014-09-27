// Compiled by ClojureScript to Go 0.0-2356
// cljs.core

package core

import (
	"fmt"
	"math"
	"reflect"

	"github.com/hraberg/cljs.go/goog"
	goog_array "github.com/hraberg/cljs.go/goog/array"
	goog_string "github.com/hraberg/cljs.go/goog/string"
	"github.com/hraberg/cljs.go/js"
	"github.com/hraberg/cljs.go/js/Math"
)

func init() {
	X_STAR_unchecked_if_STAR_ = false

	X_STAR_print_fn_STAR_ = func(_STAR_print_fn_STAR_ *AFn) *AFn {
		return Fn(_STAR_print_fn_STAR_, 1, func(___ interface{}) interface{} {
			panic((&js.Error{"No *print-fn* fn set for evaluation environment"}))
		})
	}(&AFn{})

	X_STAR_flush_on_newline_STAR_ = true

	X_STAR_print_newline_STAR_ = true

	X_STAR_print_readably_STAR_ = true

	X_STAR_print_meta_STAR_ = false

	X_STAR_print_dup_STAR_ = false

	Pr_opts = func(pr_opts *AFn) *AFn {
		return Fn(pr_opts, 0, func() interface{} {
			return (&CljsCorePersistentArrayMap{nil, float64(5), []interface{}{(&CljsCoreKeyword{Ns: nil, Name: "flush-on-newline", Fqn: "flush-on-newline", X_hash: float64(-151457939)}), X_STAR_flush_on_newline_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), X_STAR_print_readably_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}), X_STAR_print_meta_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "dup", Fqn: "dup", X_hash: float64(556298533)}), X_STAR_print_dup_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "print-length", Fqn: "print-length", X_hash: float64(1931866356)}), X_STAR_print_length_STAR_}, nil})
		})
	}(&AFn{})

	Not_native = nil

	Identical_QMARK_ = func(identical_QMARK_ *AFn) *AFn {
		return Fn(identical_QMARK_, 2, func(x interface{}, y interface{}) bool {
			return reflect.DeepEqual(x, y)
		})
	}(&AFn{})

	Nil_QMARK_ = func(nil_QMARK_ *AFn) *AFn {
		return Fn(nil_QMARK_, 1, func(x interface{}) bool {
			return reflect.DeepEqual(x, nil)
		})
	}(&AFn{})

	Array_QMARK_ = func(array_QMARK_ *AFn) *AFn {
		return Fn(array_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Kind() == reflect.Slice
		})
	}(&AFn{})

	Number_QMARK_ = func(number_QMARK_ *AFn) *AFn {
		return Fn(number_QMARK_, 1, func(n interface{}) bool {
			return Value_(n).Kind() == reflect.Float64
		})
	}(&AFn{})

	Not = func(not *AFn) *AFn {
		return Fn(not, 1, func(x interface{}) bool {
			if Truth_(x) {
				return false
			} else {
				return true
			}
		})
	}(&AFn{})

	Some_QMARK_ = func(some_QMARK_ *AFn) *AFn {
		return Fn(some_QMARK_, 1, func(x interface{}) bool {
			return !(Nil_(x))
		})
	}(&AFn{})

	String_QMARK_ = func(string_QMARK_ *AFn) *AFn {
		return Fn(string_QMARK_, 1, func(x interface{}) bool {
			{
				var G__4028 = x
				_ = G__4028
				return Truth_(Native_invoke_func.X_invoke_Arity2(goog.IsString, []interface{}{G__4028}))
			}
		})
	}(&AFn{})

	X_STAR_main_cli_fn_STAR_ = nil

	Aclone = func(aclone *AFn) *AFn {
		return Fn(aclone, 1, func(arr interface{}) interface{} {
			{
				var len = Alength_(arr)
				var new_arr = make([]interface{}, int(len))
				_, _ = len, new_arr
				{
					var n__1054__auto___4034 = len
					_ = n__1054__auto___4034
					{
						var i_4035 = float64(0)
						_ = i_4035
						for {
							if i_4035 < n__1054__auto___4034 {
								new_arr[int(i_4035)] = Aget_(arr, i_4035)
								i_4035 = (i_4035 + float64(1))
								continue
							} else {
							}
							break
						}
					}
				}
				return new_arr
			}
		})
	}(&AFn{})

	Aget = func(aget *AFn) *AFn {
		return Fn(aget, 2, func(array interface{}, i interface{}) interface{} {
			return Aget_(array, i.(float64))
		}, func(array_i_idxs__ ...interface{}) interface{} {
			var array = array_i_idxs__[0]
			var i = array_i_idxs__[1]
			var idxs = Seq.Arity1IQ(array_i_idxs__[2])
			_, _, _ = array, i, idxs
			return Apply.X_invoke_Arity3(aget, aget.X_invoke_Arity2(array, i), idxs)
		})
	}(&AFn{})

	Aset = func(aset *AFn) *AFn {
		return Fn(aset, 3, func(array interface{}, i interface{}, val interface{}) interface{} {
			return func() interface{} {
				array.([]interface{})[int(i.(float64))] = val
				return array.([]interface{})[int(i.(float64))]
			}()
		}, func(array_idx_idx2_idxv__ ...interface{}) interface{} {
			var array = array_idx_idx2_idxv__[0]
			var idx = array_idx_idx2_idxv__[1]
			var idx2 = array_idx_idx2_idxv__[2]
			var idxv = Seq.Arity1IQ(array_idx_idx2_idxv__[3])
			_, _, _, _ = array, idx, idx2, idxv
			return Apply.X_invoke_Arity4(aset, Aget_(array, idx.(float64)), idx2, idxv)
		})
	}(&AFn{})

	Alength = func(alength *AFn) *AFn {
		return Fn(alength, 1, func(array interface{}) float64 {
			return Alength_(array)
		})
	}(&AFn{})

	Into_array = func(into_array *AFn) *AFn {
		return Fn(into_array, 2, func(aseq interface{}) []interface{} {
			return into_array.Arity2IIA(nil, aseq)
		}, func(type_ interface{}, aseq interface{}) []interface{} {
			return Reduce.X_invoke_Arity3(func(G__4040 *AFn) *AFn {
				return Fn(G__4040, 2, func(a interface{}, x interface{}) interface{} {
					js.JSArray_(&a).Push(x)
					return a
				})
			}(&AFn{}), []interface{}{}, aseq).([]interface{})
		})
	}(&AFn{})

	X_invoke = func(_invoke *AFn) *AFn {
		return Fn(_invoke, 21, func(this interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity0()
		}, func(this interface{}, a interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity1(a)
		}, func(this interface{}, a interface{}, b interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity2(a, b)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity3(a, b, c)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity4(a, b, c, d)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity5(a, b, c, d, e)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity6(a, b, c, d, e, f)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity7(a, b, c, d, e, f, g)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity8(a, b, c, d, e, f, g, h)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity9(a, b, c, d, e, f, g, h, i)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity10(a, b, c, d, e, f, g, h, i, j)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity11(a, b, c, d, e, f, g, h, i, j, k)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity12(a, b, c, d, e, f, g, h, i, j, k, l)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity13(a, b, c, d, e, f, g, h, i, j, k, l, m)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity14(a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity15(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity16(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity17(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity18(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity19(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity20(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
		})
	}(&AFn{})

	X_clone = func(_clone *AFn) *AFn {
		return Fn(_clone, 1, func(value interface{}) interface{} {
			return value.(CljsCoreICloneable).X_clone_Arity1()
		})
	}(&AFn{})

	X_count = func(_count *AFn) *AFn {
		return Fn(_count, 1, func(coll interface{}) float64 {
			return coll.(CljsCoreICounted).X_count_Arity1()
		})
	}(&AFn{})

	X_empty = func(_empty *AFn) *AFn {
		return Fn(_empty, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIEmptyableCollection).X_empty_Arity1()
		})
	}(&AFn{})

	X_conj = func(_conj *AFn) *AFn {
		return Fn(_conj, 2, func(coll interface{}, o interface{}) interface{} {
			return coll.(CljsCoreICollection).X_conj_Arity2(o)
		})
	}(&AFn{})

	X_nth = func(_nth *AFn) *AFn {
		return Fn(_nth, 3, func(coll interface{}, n interface{}) interface{} {
			return coll.(CljsCoreIIndexed).X_nth_Arity2(n)
		}, func(coll interface{}, n interface{}, not_found interface{}) interface{} {
			return coll.(CljsCoreIIndexed).X_nth_Arity3(n, not_found)
		})
	}(&AFn{})

	X_first = func(_first *AFn) *AFn {
		return Fn(_first, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreISeq).X_first_Arity1()
		})
	}(&AFn{})

	X_rest = func(_rest *AFn) *AFn {
		return Fn(_rest, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreISeq).X_rest_Arity1()
		})
	}(&AFn{})

	X_next = func(_next *AFn) *AFn {
		return Fn(_next, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreINext).X_next_Arity1()
		})
	}(&AFn{})

	X_lookup = func(_lookup *AFn) *AFn {
		return Fn(_lookup, 3, func(o interface{}, k interface{}) interface{} {
			return o.(CljsCoreILookup).X_lookup_Arity2(k)
		}, func(o interface{}, k interface{}, not_found interface{}) interface{} {
			return o.(CljsCoreILookup).X_lookup_Arity3(k, not_found)
		})
	}(&AFn{})

	X_contains_key_QMARK_ = func(_contains_key_QMARK_ *AFn) *AFn {
		return Fn(_contains_key_QMARK_, 2, func(coll interface{}, k interface{}) bool {
			return coll.(CljsCoreIAssociative).X_contains_key_QMARK__Arity2(k)
		})
	}(&AFn{})

	X_assoc = func(_assoc *AFn) *AFn {
		return Fn(_assoc, 3, func(coll interface{}, k interface{}, v interface{}) interface{} {
			return coll.(CljsCoreIAssociative).X_assoc_Arity3(k, v)
		})
	}(&AFn{})

	X_dissoc = func(_dissoc *AFn) *AFn {
		return Fn(_dissoc, 2, func(coll interface{}, k interface{}) interface{} {
			return coll.(CljsCoreIMap).X_dissoc_Arity2(k)
		})
	}(&AFn{})

	X_key = func(_key *AFn) *AFn {
		return Fn(_key, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIMapEntry).X_key_Arity1()
		})
	}(&AFn{})

	X_val = func(_val *AFn) *AFn {
		return Fn(_val, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIMapEntry).X_val_Arity1()
		})
	}(&AFn{})

	X_disjoin = func(_disjoin *AFn) *AFn {
		return Fn(_disjoin, 2, func(coll interface{}, v interface{}) interface{} {
			return coll.(CljsCoreISet).X_disjoin_Arity2(v)
		})
	}(&AFn{})

	X_peek = func(_peek *AFn) *AFn {
		return Fn(_peek, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIStack).X_peek_Arity1()
		})
	}(&AFn{})

	X_pop = func(_pop *AFn) *AFn {
		return Fn(_pop, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIStack).X_pop_Arity1()
		})
	}(&AFn{})

	X_assoc_n = func(_assoc_n *AFn) *AFn {
		return Fn(_assoc_n, 3, func(coll interface{}, n interface{}, val interface{}) interface{} {
			return coll.(CljsCoreIVector).X_assoc_n_Arity3(n, val)
		})
	}(&AFn{})

	X_deref = func(_deref *AFn) *AFn {
		return Fn(_deref, 1, func(o interface{}) interface{} {
			return o.(CljsCoreIDeref).X_deref_Arity1()
		})
	}(&AFn{})

	X_deref_with_timeout = func(_deref_with_timeout *AFn) *AFn {
		return Fn(_deref_with_timeout, 3, func(o interface{}, msec interface{}, timeout_val interface{}) interface{} {
			return o.(CljsCoreIDerefWithTimeout).X_deref_with_timeout_Arity3(msec, timeout_val)
		})
	}(&AFn{})

	X_meta = func(_meta *AFn) *AFn {
		return Fn(_meta, 1, func(o interface{}) interface{} {
			return o.(CljsCoreIMeta).X_meta_Arity1()
		})
	}(&AFn{})

	X_with_meta = func(_with_meta *AFn) *AFn {
		return Fn(_with_meta, 2, func(o interface{}, meta interface{}) interface{} {
			return o.(CljsCoreIWithMeta).X_with_meta_Arity2(meta)
		})
	}(&AFn{})

	X_reduce = func(_reduce *AFn) *AFn {
		return Fn(_reduce, 3, func(coll interface{}, f interface{}) interface{} {
			return coll.(CljsCoreIReduce).X_reduce_Arity2(f)
		}, func(coll interface{}, f interface{}, start interface{}) interface{} {
			return coll.(CljsCoreIReduce).X_reduce_Arity3(f, start)
		})
	}(&AFn{})

	X_kv_reduce = func(_kv_reduce *AFn) *AFn {
		return Fn(_kv_reduce, 3, func(coll interface{}, f interface{}, init interface{}) interface{} {
			return coll.(CljsCoreIKVReduce).X_kv_reduce_Arity3(f, init)
		})
	}(&AFn{})

	X_equiv = func(_equiv *AFn) *AFn {
		return Fn(_equiv, 2, func(o interface{}, other interface{}) bool {
			return o.(CljsCoreIEquiv).X_equiv_Arity2(other)
		})
	}(&AFn{})

	X_hash = func(_hash *AFn) *AFn {
		return Fn(_hash, 1, func(o interface{}) interface{} {
			return o.(CljsCoreIHash).X_hash_Arity1()
		})
	}(&AFn{})

	X_seq = func(_seq *AFn) *AFn {
		return Fn(_seq, 1, func(o interface{}) interface{} {
			return o.(CljsCoreISeqable).X_seq_Arity1()
		})
	}(&AFn{})

	X_rseq = func(_rseq *AFn) *AFn {
		return Fn(_rseq, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIReversible).X_rseq_Arity1()
		})
	}(&AFn{})

	X_sorted_seq = func(_sorted_seq *AFn) *AFn {
		return Fn(_sorted_seq, 2, func(coll interface{}, ascending_QMARK_ interface{}) interface{} {
			return coll.(CljsCoreISorted).X_sorted_seq_Arity2(ascending_QMARK_)
		})
	}(&AFn{})

	X_sorted_seq_from = func(_sorted_seq_from *AFn) *AFn {
		return Fn(_sorted_seq_from, 3, func(coll interface{}, k interface{}, ascending_QMARK_ interface{}) interface{} {
			return coll.(CljsCoreISorted).X_sorted_seq_from_Arity3(k, ascending_QMARK_)
		})
	}(&AFn{})

	X_entry_key = func(_entry_key *AFn) *AFn {
		return Fn(_entry_key, 2, func(coll interface{}, entry interface{}) interface{} {
			return coll.(CljsCoreISorted).X_entry_key_Arity2(entry)
		})
	}(&AFn{})

	X_comparator = func(_comparator *AFn) *AFn {
		return Fn(_comparator, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreISorted).X_comparator_Arity1()
		})
	}(&AFn{})

	X_write = func(_write *AFn) *AFn {
		return Fn(_write, 2, func(writer interface{}, s interface{}) interface{} {
			return writer.(CljsCoreIWriter).X_write_Arity2(s)
		})
	}(&AFn{})

	X_flush = func(_flush *AFn) *AFn {
		return Fn(_flush, 1, func(writer interface{}) interface{} {
			return writer.(CljsCoreIWriter).X_flush_Arity1()
		})
	}(&AFn{})

	X_pr_writer = func(_pr_writer *AFn) *AFn {
		return Fn(_pr_writer, 3, func(o interface{}, writer interface{}, opts interface{}) interface{} {
			return o.(CljsCoreIPrintWithWriter).X_pr_writer_Arity3(writer, opts)
		})
	}(&AFn{})

	X_realized_QMARK_ = func(_realized_QMARK_ *AFn) *AFn {
		return Fn(_realized_QMARK_, 1, func(d interface{}) bool {
			return d.(CljsCoreIPending).X_realized_QMARK__Arity1()
		})
	}(&AFn{})

	X_notify_watches = func(_notify_watches *AFn) *AFn {
		return Fn(_notify_watches, 3, func(this interface{}, oldval interface{}, newval interface{}) interface{} {
			return this.(CljsCoreIWatchable).X_notify_watches_Arity3(oldval, newval)
		})
	}(&AFn{})

	X_add_watch = func(_add_watch *AFn) *AFn {
		return Fn(_add_watch, 3, func(this interface{}, key interface{}, f interface{}) interface{} {
			return this.(CljsCoreIWatchable).X_add_watch_Arity3(key, f)
		})
	}(&AFn{})

	X_remove_watch = func(_remove_watch *AFn) *AFn {
		return Fn(_remove_watch, 2, func(this interface{}, key interface{}) interface{} {
			return this.(CljsCoreIWatchable).X_remove_watch_Arity2(key)
		})
	}(&AFn{})

	X_as_transient = func(_as_transient *AFn) *AFn {
		return Fn(_as_transient, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIEditableCollection).X_as_transient_Arity1()
		})
	}(&AFn{})

	X_conj_BANG_ = func(_conj_BANG_ *AFn) *AFn {
		return Fn(_conj_BANG_, 2, func(tcoll interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientCollection).X_conj_BANG__Arity2(val)
		})
	}(&AFn{})

	X_persistent_BANG_ = func(_persistent_BANG_ *AFn) *AFn {
		return Fn(_persistent_BANG_, 1, func(tcoll interface{}) interface{} {
			return tcoll.(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
		})
	}(&AFn{})

	X_assoc_BANG_ = func(_assoc_BANG_ *AFn) *AFn {
		return Fn(_assoc_BANG_, 3, func(tcoll interface{}, key interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3(key, val)
		})
	}(&AFn{})

	X_dissoc_BANG_ = func(_dissoc_BANG_ *AFn) *AFn {
		return Fn(_dissoc_BANG_, 2, func(tcoll interface{}, key interface{}) interface{} {
			return tcoll.(CljsCoreITransientMap).X_dissoc_BANG__Arity2(key)
		})
	}(&AFn{})

	X_assoc_n_BANG_ = func(_assoc_n_BANG_ *AFn) *AFn {
		return Fn(_assoc_n_BANG_, 3, func(tcoll interface{}, n interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientVector).X_assoc_n_BANG__Arity3(n, val)
		})
	}(&AFn{})

	X_pop_BANG_ = func(_pop_BANG_ *AFn) *AFn {
		return Fn(_pop_BANG_, 1, func(tcoll interface{}) interface{} {
			return tcoll.(CljsCoreITransientVector).X_pop_BANG__Arity1()
		})
	}(&AFn{})

	X_disjoin_BANG_ = func(_disjoin_BANG_ *AFn) *AFn {
		return Fn(_disjoin_BANG_, 2, func(tcoll interface{}, v interface{}) interface{} {
			return tcoll.(CljsCoreITransientSet).X_disjoin_BANG__Arity2(v)
		})
	}(&AFn{})

	X_compare = func(_compare *AFn) *AFn {
		return Fn(_compare, 2, func(x interface{}, y interface{}) float64 {
			return x.(CljsCoreIComparable).X_compare_Arity2(y)
		})
	}(&AFn{})

	X_drop_first = func(_drop_first *AFn) *AFn {
		return Fn(_drop_first, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIChunk).X_drop_first_Arity1()
		})
	}(&AFn{})

	X_chunked_first = func(_chunked_first *AFn) *AFn {
		return Fn(_chunked_first, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIChunkedSeq).X_chunked_first_Arity1()
		})
	}(&AFn{})

	X_chunked_rest = func(_chunked_rest *AFn) *AFn {
		return Fn(_chunked_rest, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIChunkedSeq).X_chunked_rest_Arity1()
		})
	}(&AFn{})

	X_chunked_next = func(_chunked_next *AFn) *AFn {
		return Fn(_chunked_next, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIChunkedNext).X_chunked_next_Arity1()
		})
	}(&AFn{})

	X_name = func(_name *AFn) *AFn {
		return Fn(_name, 1, func(x interface{}) interface{} {
			return x.(CljsCoreINamed).X_name_Arity1()
		})
	}(&AFn{})

	X_namespace = func(_namespace *AFn) *AFn {
		return Fn(_namespace, 1, func(x interface{}) interface{} {
			return x.(CljsCoreINamed).X_namespace_Arity1()
		})
	}(&AFn{})

	X_reset_BANG_ = func(_reset_BANG_ *AFn) *AFn {
		return Fn(_reset_BANG_, 2, func(o interface{}, new_value interface{}) interface{} {
			return o.(CljsCoreIReset).X_reset_BANG__Arity2(new_value)
		})
	}(&AFn{})

	X_swap_BANG_ = func(_swap_BANG_ *AFn) *AFn {
		return Fn(_swap_BANG_, 5, func(o interface{}, f interface{}) interface{} {
			return o.(CljsCoreISwap).X_swap_BANG__Arity2(f)
		}, func(o interface{}, f interface{}, a interface{}) interface{} {
			return o.(CljsCoreISwap).X_swap_BANG__Arity3(f, a)
		}, func(o interface{}, f interface{}, a interface{}, b interface{}) interface{} {
			return o.(CljsCoreISwap).X_swap_BANG__Arity4(f, a, b)
		}, func(o interface{}, f interface{}, a interface{}, b interface{}, xs interface{}) interface{} {
			return o.(CljsCoreISwap).X_swap_BANG__Arity5(f, a, b, xs)
		})
	}(&AFn{})

	X_iterator = func(_iterator *AFn) *AFn {
		return Fn(_iterator, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIIterable).X_iterator_Arity1()
		})
	}(&AFn{})

	X__GT_StringBufferWriter = func(__GT_StringBufferWriter *AFn) *AFn {
		return Fn(__GT_StringBufferWriter, 1, func(sb interface{}) interface{} {
			return (&CljsCoreStringBufferWriter{sb})
		})
	}(&AFn{})

	Pr_str_STAR_ = func(pr_str_STAR_ *AFn) *AFn {
		return Fn(pr_str_STAR_, 1, func(obj interface{}) interface{} {
			{
				var sb = (&goog_string.StringBuffer{})
				var writer = (&CljsCoreStringBufferWriter{sb})
				_, _ = sb, writer
				obj.(CljsCoreIPrintWithWriter).X_pr_writer_Arity3(writer, Pr_opts.X_invoke_Arity0().(CljsCoreIMap))
				writer.X_flush_Arity1()
				return (`` + Str.X_invoke_Arity1(sb).(string))
			}
		})
	}(&AFn{})

	Int_rotate_left = func(int_rotate_left *AFn) *AFn {
		return Fn(int_rotate_left, 2, func(x interface{}, n interface{}) float64 {
			return float64((Int32_(float64((Int32_(x.(float64)) << UInt32_(n.(float64))))) | Int32_(float64((UInt32_(x.(float64)) >> UInt32_(float64((32+Int32_((-n.(float64))))%32)))))))
		})
	}(&AFn{})

	if (Value_(Math.Imul).Kind() != reflect.Invalid) && (!(func() interface{} {
		var G__4043 = float64(4294967295)
		var G__4044 = float64(5)
		_, _ = G__4043, G__4044
		return Native_invoke_func.X_invoke_Arity2(Math.Imul, []interface{}{G__4043, G__4044})
	}().(float64) == float64(0))) {
		Imul = func(imul *AFn) *AFn {
			return Fn(imul, 2, func(a interface{}, b interface{}) float64 {
				{
					var G__4047 = a
					var G__4048 = b
					_, _ = G__4047, G__4048
					return Native_invoke_func.X_invoke_Arity2(Math.Imul, []interface{}{G__4047, G__4048}).(float64)
				}
			})
		}(&AFn{})

	} else {
		Imul = func(imul *AFn) *AFn {
			return Fn(imul, 2, func(a interface{}, b interface{}) float64 {
				{
					var ah = float64((Int32_(float64((UInt32_(a.(float64)) >> UInt32_(float64((32+Int32_(float64(16)))%32))))) & Int32_(float64(65535))))
					var al = float64((Int32_(a.(float64)) & Int32_(float64(65535))))
					var bh = float64((Int32_(float64((UInt32_(b.(float64)) >> UInt32_(float64((32+Int32_(float64(16)))%32))))) & Int32_(float64(65535))))
					var bl = float64((Int32_(b.(float64)) & Int32_(float64(65535))))
					_, _, _, _ = ah, al, bh, bl
					return float64((Int32_(((al * bl) + float64((UInt32_(float64((Int32_(((ah * bl) + (al * bh))) << UInt32_(float64(16))))) >> UInt32_(float64((32+Int32_(float64(0)))%32)))))) | Int32_(float64(0))))
				}
			})
		}(&AFn{})

	}
	M3_seed = float64(0)

	M3_C1 = float64(3432918353)

	M3_C2 = float64(461845907)

	M3_mix_K1 = func(m3_mix_K1 *AFn) *AFn {
		return Fn(m3_mix_K1, 1, func(k1 interface{}) float64 {
			return Imul.Arity2IIF(Int_rotate_left.Arity2IIF(Imul.Arity2IIF(k1, M3_C1), float64(15)), M3_C2)
		})
	}(&AFn{})

	M3_mix_H1 = func(m3_mix_H1 *AFn) *AFn {
		return Fn(m3_mix_H1, 2, func(h1 interface{}, k1 interface{}) float64 {
			return (Imul.Arity2IIF(Int_rotate_left.Arity2IIF(float64((Int32_(h1.(float64))^Int32_(k1.(float64)))), float64(13)), float64(5)) + float64(3864292196))
		})
	}(&AFn{})

	M3_fmix = func(m3_fmix *AFn) *AFn {
		return Fn(m3_fmix, 2, func(h1 interface{}, len interface{}) float64 {
			{
				var h1___1 = h1
				var h1___2 = float64((Int32_(h1___1.(float64)) ^ Int32_(len.(float64))))
				var h1___3 = float64((Int32_(h1___2) ^ Int32_(float64((UInt32_(h1___2) >> UInt32_(float64((32+Int32_(float64(16)))%32)))))))
				var h1___4 = Imul.Arity2IIF(h1___3, float64(2246822507))
				var h1___5 = float64((Int32_(h1___4) ^ Int32_(float64((UInt32_(h1___4) >> UInt32_(float64((32+Int32_(float64(13)))%32)))))))
				var h1___6 = Imul.Arity2IIF(h1___5, float64(3266489909))
				var h1___7 = float64((Int32_(h1___6) ^ Int32_(float64((UInt32_(h1___6) >> UInt32_(float64((32+Int32_(float64(16)))%32)))))))
				_, _, _, _, _, _, _ = h1___1, h1___2, h1___3, h1___4, h1___5, h1___6, h1___7
				return h1___7
			}
		})
	}(&AFn{})

	M3_hash_int = func(m3_hash_int *AFn) *AFn {
		return Fn(m3_hash_int, 1, func(in interface{}) float64 {
			if in.(float64) == float64(0) {
				return in.(float64)
			} else {
				{
					var k1 = M3_mix_K1.Arity1IF(in)
					var h1 = M3_mix_H1.Arity2IIF(M3_seed, k1)
					_, _ = k1, h1
					return M3_fmix.Arity2IIF(h1, float64(4))
				}
			}
		})
	}(&AFn{})

	M3_hash_unencoded_chars = func(m3_hash_unencoded_chars *AFn) *AFn {
		return Fn(m3_hash_unencoded_chars, 1, func(in interface{}) float64 {
			{
				var h1 = func() interface{} {
					var i = float64(1)
					var h1 interface{} = M3_seed
					_, _ = i, h1
					for {
						if i < Alength_(in) {
							i, h1 = (i + float64(2)), M3_mix_H1.Arity2IIF(h1, M3_mix_K1.Arity1IF(float64((Int32_(Native_invoke_instance_method.X_invoke_Arity3(in, "CharCodeAt", []interface{}{(i - float64(1))}).(float64))|Int32_(float64((Int32_(Native_invoke_instance_method.X_invoke_Arity3(in, "CharCodeAt", []interface{}{i}).(float64))<<UInt32_(float64(16)))))))))
							continue
						} else {
							return h1
						}
					}
				}()
				var h1___1 = func() interface{} {
					if float64((Int32_(Alength_(in)) & Int32_(float64(1)))) == float64(1) {
						return float64((Int32_(h1.(float64)) ^ Int32_(M3_mix_K1.Arity1IF(Native_invoke_instance_method.X_invoke_Arity3(in, "CharCodeAt", []interface{}{(Alength_(in) - float64(1))})))))
					} else {
						return h1
					}
				}()
				_, _ = h1, h1___1
				return M3_fmix.Arity2IIF(h1___1, Imul.Arity2IIF(float64(2), Alength_(in)))
			}
		})
	}(&AFn{})

	String_hash_cache_count = float64(0)

	Hash_string_STAR_ = func(hash_string_STAR_ *AFn) *AFn {
		return Fn(hash_string_STAR_, 1, func(s interface{}) interface{} {
			if !(Nil_(s)) {
				{
					var len = Alength_(s)
					_ = len
					if len > float64(0) {
						{
							var i = float64(0)
							var hash = float64(0)
							_, _ = i, hash
							for {
								if i < len {
									i, hash = (i + float64(1)), (Imul.Arity2IIF(float64(31), hash) + Native_invoke_instance_method.X_invoke_Arity3(s, "CharCodeAt", []interface{}{i}).(float64))
									continue
								} else {
									return hash
								}
							}
						}
					} else {
						return float64(0)
					}
				}
			} else {
				return float64(0)
			}
		})
	}(&AFn{})

	Hash = func(hash *AFn) *AFn {
		return Fn(hash, 1, func(o interface{}) interface{} {
			if Value_(o).Type().Implements(reflect.TypeOf((*CljsCoreIHash)(nil)).Elem()) {
				return o.(CljsCoreIHash).X_hash_Arity1()
			} else {
				if Value_(o).Kind() == reflect.Float64 {
					return math.Mod(func() interface{} {
						var G__4056 = o
						_ = G__4056
						return Native_invoke_func.X_invoke_Arity2(Math.Floor, []interface{}{G__4056})
					}().(float64), float64(2147483647))
				} else {
					if o == true {
						return float64(1)
					} else {
						if o == false {
							return float64(0)
						} else {
							if Value_(o).Kind() == reflect.String {
								return M3_hash_int.Arity1IF(Hash_string.X_invoke_Arity1(o))
							} else {
								if Nil_(o) {
									return float64(0)
								} else {
									return o.(CljsCoreIHash).X_hash_Arity1()

								}
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Hash_combine = func(hash_combine *AFn) *AFn {
		return Fn(hash_combine, 2, func(seed interface{}, hash interface{}) interface{} {
			return float64((Int32_(seed.(float64)) ^ Int32_((((hash.(float64) + float64(2654435769)) + float64((Int32_(seed.(float64)) << UInt32_(float64(6))))) + float64((Int32_(seed.(float64)) >> UInt32_(float64(2))))))))
		})
	}(&AFn{})

	Hash_symbol = func(hash_symbol *AFn) *AFn {
		return Fn(hash_symbol, 1, func(sym interface{}) interface{} {
			return Hash_combine.X_invoke_Arity2(M3_hash_unencoded_chars.Arity1IF(Native_get_instance_field.X_invoke_Arity2(sym, "Name")), Hash_string.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(sym, "Ns"))).(float64)
		})
	}(&AFn{})

	Compare_symbols = func(compare_symbols *AFn) *AFn {
		return Fn(compare_symbols, 2, func(a interface{}, b interface{}) interface{} {
			if X_EQ_.Arity2IIB(a, b) {
				return float64(0)
			} else {
				if Truth_(func() interface{} {
					var and__159__auto__ = Not.Arity1IB(Native_get_instance_field.X_invoke_Arity2(a, "Ns"))
					_ = and__159__auto__
					if Truth_(and__159__auto__) {
						return Native_get_instance_field.X_invoke_Arity2(b, "Ns")
					} else {
						return and__159__auto__
					}
				}()) {
					return float64(-1)
				} else {
					if Truth_(Native_get_instance_field.X_invoke_Arity2(a, "Ns")) {
						if Not.Arity1IB(Native_get_instance_field.X_invoke_Arity2(b, "Ns")) {
							return float64(1)
						} else {
							{
								var nsc = Compare.Arity2IIF(Native_get_instance_field.X_invoke_Arity2(a, "Ns"), Native_get_instance_field.X_invoke_Arity2(b, "Ns"))
								_ = nsc
								if nsc == float64(0) {
									return Compare.Arity2IIF(Native_get_instance_field.X_invoke_Arity2(a, "Name"), Native_get_instance_field.X_invoke_Arity2(b, "Name"))
								} else {
									return nsc
								}
							}
						}
					} else {
						return Compare.Arity2IIF(Native_get_instance_field.X_invoke_Arity2(a, "Name"), Native_get_instance_field.X_invoke_Arity2(b, "Name"))

					}
				}
			}
		})
	}(&AFn{})

	X__GT_Symbol = func(__GT_Symbol *AFn) *AFn {
		return Fn(__GT_Symbol, 5, func(ns interface{}, name interface{}, str interface{}, _hash interface{}, _meta interface{}) interface{} {
			return (&CljsCoreSymbol{ns, name, str, _hash, _meta})
		})
	}(&AFn{})

	Symbol = func(symbol *AFn) *AFn {
		return Fn(symbol, 2, func(name interface{}) interface{} {
			if Value_(name).Type().AssignableTo(reflect.TypeOf((**CljsCoreSymbol)(nil)).Elem()) {
				return name
			} else {
				return symbol.X_invoke_Arity2(nil, name).(*CljsCoreSymbol)
			}
		}, func(ns interface{}, name interface{}) interface{} {
			{
				var sym_str = func() interface{} {
					if !(Nil_(ns)) {
						return (`` + Str.X_invoke_Arity1(ns).(string) + "/" + Str.X_invoke_Arity1(name).(string))
					} else {
						return name
					}
				}()
				_ = sym_str
				return (&CljsCoreSymbol{ns, name, sym_str, nil, nil})
			}
		})
	}(&AFn{})

	Iterable_QMARK_ = func(iterable_QMARK_ *AFn) *AFn {
		return Fn(iterable_QMARK_, 1, func(x interface{}) interface{} {
			return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreIIterable)(nil)).Elem())
		})
	}(&AFn{})

	Clone = func(clone *AFn) *AFn {
		return Fn(clone, 1, func(value interface{}) interface{} {
			return value.(CljsCoreICloneable).X_clone_Arity1()
		})
	}(&AFn{})

	Cloneable_QMARK_ = func(cloneable_QMARK_ *AFn) *AFn {
		return Fn(cloneable_QMARK_, 1, func(value interface{}) interface{} {
			return Value_(value).Type().Implements(reflect.TypeOf((*CljsCoreICloneable)(nil)).Elem())
		})
	}(&AFn{})

	Seq = func(seq *AFn) *AFn {
		return Fn(seq, 1, func(coll interface{}) CljsCoreISeq {
			if Nil_(coll) {
				return nil
			} else {
				if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreISeqable)(nil)).Elem()) {
					return Seq_(coll.(CljsCoreISeqable).X_seq_Arity1())
				} else {
					if Value_(coll).Kind() == reflect.Slice {
						if Alength_(coll) == float64(0) {
							return nil
						} else {
							return (&CljsCoreIndexedSeq{coll, float64(0)})
						}
					} else {
						if Value_(coll).Kind() == reflect.String {
							if Alength_(coll) == float64(0) {
								return nil
							} else {
								return (&CljsCoreIndexedSeq{coll, float64(0)})
							}
						} else {
							if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreISeqable)(nil)).Elem()) {
								return Seq_(coll.(CljsCoreISeqable).X_seq_Arity1())
							} else {
								panic((&js.Error{(`` + Str.X_invoke_Arity1(coll).(string) + " is not ISeqable")}))

							}
						}
					}
				}
			}
		})
	}(&AFn{})

	First = func(first *AFn) *AFn {
		return Fn(first, 1, func(coll interface{}) interface{} {
			if Nil_(coll) {
				return nil
			} else {
				if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreISeq)(nil)).Elem()) {
					return coll.(CljsCoreISeq).X_first_Arity1()
				} else {
					{
						var s = Seq.Arity1IQ(coll)
						_ = s
						if Nil_(s) {
							return nil
						} else {
							return s.X_first_Arity1()
						}
					}
				}
			}
		})
	}(&AFn{})

	Rest = func(rest *AFn) *AFn {
		return Fn(rest, 1, func(coll interface{}) CljsCoreISeq {
			if !(Nil_(coll)) {
				if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreISeq)(nil)).Elem()) {
					return Seq_(coll.(CljsCoreISeq).X_rest_Arity1())
				} else {
					{
						var s = Seq.Arity1IQ(coll)
						_ = s
						if Truth_(s) {
							return Seq_(s.X_rest_Arity1())
						} else {
							return CljsCoreIEmptyList(CljsCoreList_EMPTY)
						}
					}
				}
			} else {
				return CljsCoreIEmptyList(CljsCoreList_EMPTY)
			}
		})
	}(&AFn{})

	Next = func(next *AFn) *AFn {
		return Fn(next, 1, func(coll interface{}) CljsCoreISeq {
			if Nil_(coll) {
				return nil
			} else {
				if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreINext)(nil)).Elem()) {
					return Seq_(coll.(CljsCoreINext).X_next_Arity1())
				} else {
					return Seq.Arity1IQ(Rest.Arity1IQ(coll))
				}
			}
		})
	}(&AFn{})

	Mix_collection_hash = func(mix_collection_hash *AFn) *AFn {
		return Fn(mix_collection_hash, 2, func(hash_basis interface{}, count interface{}) float64 {
			{
				var h1 = M3_seed
				var k1 = M3_mix_K1.Arity1IF(hash_basis)
				var h1___1 = M3_mix_H1.Arity2IIF(h1, k1)
				_, _, _ = h1, k1, h1___1
				return M3_fmix.Arity2IIF(h1___1, count)
			}
		})
	}(&AFn{})

	Hash_ordered_coll = func(hash_ordered_coll *AFn) *AFn {
		return Fn(hash_ordered_coll, 1, func(coll interface{}) float64 {
			{
				var n = float64(0)
				var hash_code = float64(1)
				var coll___1 interface{} = Seq.Arity1IQ(coll)
				_, _, _ = n, hash_code, coll___1
				for {
					if !(Nil_(coll___1)) {
						n, hash_code, coll___1 = (n + float64(1)), float64((Int32_((Imul.Arity2IIF(float64(31), hash_code) + Hash.X_invoke_Arity1(First.X_invoke_Arity1(coll___1)).(float64))) | Int32_(float64(0)))), Next.Arity1IQ(coll___1)
						continue
					} else {
						return Mix_collection_hash.Arity2IIF(hash_code, n)
					}
				}
			}
		})
	}(&AFn{})

	Hash_unordered_coll = func(hash_unordered_coll *AFn) *AFn {
		return Fn(hash_unordered_coll, 1, func(coll interface{}) float64 {
			{
				var n = float64(0)
				var hash_code = float64(0)
				var coll___1 interface{} = Seq.Arity1IQ(coll)
				_, _, _ = n, hash_code, coll___1
				for {
					if !(Nil_(coll___1)) {
						n, hash_code, coll___1 = (n + float64(1)), float64((Int32_((hash_code + Hash.X_invoke_Arity1(First.X_invoke_Arity1(coll___1)).(float64))) | Int32_(float64(0)))), Next.Arity1IQ(coll___1)
						continue
					} else {
						return Mix_collection_hash.Arity2IIF(hash_code, n)
					}
				}
			}
		})
	}(&AFn{})

	Inc = func(inc *AFn) *AFn {
		return Fn(inc, 1, func(x interface{}) interface{} {
			return (x.(float64) + float64(1))
		})
	}(&AFn{})

	X__GT_Reduced = func(__GT_Reduced *AFn) *AFn {
		return Fn(__GT_Reduced, 1, func(val interface{}) interface{} {
			return (&CljsCoreReduced{val})
		})
	}(&AFn{})

	Reduced = func(reduced *AFn) *AFn {
		return Fn(reduced, 1, func(x interface{}) interface{} {
			return (&CljsCoreReduced{x})
		})
	}(&AFn{})

	Reduced_QMARK_ = func(reduced_QMARK_ *AFn) *AFn {
		return Fn(reduced_QMARK_, 1, func(r interface{}) bool {
			return Value_(r).Type().AssignableTo(reflect.TypeOf((**CljsCoreReduced)(nil)).Elem())
		})
	}(&AFn{})

	Deref = func(deref *AFn) *AFn {
		return Fn(deref, 1, func(o interface{}) interface{} {
			return o.(CljsCoreIDeref).X_deref_Arity1()
		})
	}(&AFn{})

	Ci_reduce = func(ci_reduce *AFn) *AFn {
		return Fn(ci_reduce, 4, func(cicoll interface{}, f interface{}) interface{} {
			{
				var cnt = cicoll.(CljsCoreICounted).X_count_Arity1()
				_ = cnt
				if cnt == float64(0) {
					{
						return f.(CljsCoreIFn).X_invoke_Arity0()
					}
				} else {
					{
						var val interface{} = cicoll.(CljsCoreIIndexed).X_nth_Arity2(float64(0))
						var n = float64(1)
						_, _ = val, n
						for {
							if n < cnt {
								{
									var nval = func() interface{} {
										var G__4069 = val
										var G__4070 = cicoll.(CljsCoreIIndexed).X_nth_Arity2(n)
										_, _ = G__4069, G__4070
										return f.(CljsCoreIFn).X_invoke_Arity2(G__4069, G__4070)
									}()
									_ = nval
									if Reduced_QMARK_.Arity1IB(nval) {
										return Deref.X_invoke_Arity1(nval)
									} else {
										val, n = nval, (n + float64(1))
										continue
									}
								}
							} else {
								return val
							}
						}
					}
				}
			}
		}, func(cicoll interface{}, f interface{}, val interface{}) interface{} {
			{
				var cnt = cicoll.(CljsCoreICounted).X_count_Arity1()
				_ = cnt
				{
					var val___1 interface{} = val
					var n = float64(0)
					_, _ = val___1, n
					for {
						if n < cnt {
							{
								var nval = func() interface{} {
									var G__4071 = val___1
									var G__4072 = cicoll.(CljsCoreIIndexed).X_nth_Arity2(n)
									_, _ = G__4071, G__4072
									return f.(CljsCoreIFn).X_invoke_Arity2(G__4071, G__4072)
								}()
								_ = nval
								if Reduced_QMARK_.Arity1IB(nval) {
									return Deref.X_invoke_Arity1(nval)
								} else {
									val___1, n = nval, (n + float64(1))
									continue
								}
							}
						} else {
							return val___1
						}
					}
				}
			}
		}, func(cicoll interface{}, f interface{}, val interface{}, idx interface{}) interface{} {
			{
				var cnt = cicoll.(CljsCoreICounted).X_count_Arity1()
				_ = cnt
				{
					var val___1 interface{} = val
					var n interface{} = idx
					_, _ = val___1, n
					for {
						if n.(float64) < cnt {
							{
								var nval = func() interface{} {
									var G__4073 = val___1
									var G__4074 = cicoll.(CljsCoreIIndexed).X_nth_Arity2(n)
									_, _ = G__4073, G__4074
									return f.(CljsCoreIFn).X_invoke_Arity2(G__4073, G__4074)
								}()
								_ = nval
								if Reduced_QMARK_.Arity1IB(nval) {
									return Deref.X_invoke_Arity1(nval)
								} else {
									val___1, n = nval, (n.(float64) + float64(1))
									continue
								}
							}
						} else {
							return val___1
						}
					}
				}
			}
		})
	}(&AFn{})

	Array_reduce = func(array_reduce *AFn) *AFn {
		return Fn(array_reduce, 4, func(arr interface{}, f interface{}) interface{} {
			{
				var cnt = Alength_(arr)
				_ = cnt
				if Alength_(arr) == float64(0) {
					{
						return f.(CljsCoreIFn).X_invoke_Arity0()
					}
				} else {
					{
						var val interface{} = Aget_(arr, float64(0))
						var n = float64(1)
						_, _ = val, n
						for {
							if n < cnt {
								{
									var nval = func() interface{} {
										var G__4081 = val
										var G__4082 = Aget_(arr, n)
										_, _ = G__4081, G__4082
										return f.(CljsCoreIFn).X_invoke_Arity2(G__4081, G__4082)
									}()
									_ = nval
									if Reduced_QMARK_.Arity1IB(nval) {
										return Deref.X_invoke_Arity1(nval)
									} else {
										val, n = nval, (n + float64(1))
										continue
									}
								}
							} else {
								return val
							}
						}
					}
				}
			}
		}, func(arr interface{}, f interface{}, val interface{}) interface{} {
			{
				var cnt = Alength_(arr)
				_ = cnt
				{
					var val___1 interface{} = val
					var n = float64(0)
					_, _ = val___1, n
					for {
						if n < cnt {
							{
								var nval = func() interface{} {
									var G__4083 = val___1
									var G__4084 = Aget_(arr, n)
									_, _ = G__4083, G__4084
									return f.(CljsCoreIFn).X_invoke_Arity2(G__4083, G__4084)
								}()
								_ = nval
								if Reduced_QMARK_.Arity1IB(nval) {
									return Deref.X_invoke_Arity1(nval)
								} else {
									val___1, n = nval, (n + float64(1))
									continue
								}
							}
						} else {
							return val___1
						}
					}
				}
			}
		}, func(arr interface{}, f interface{}, val interface{}, idx interface{}) interface{} {
			{
				var cnt = Alength_(arr)
				_ = cnt
				{
					var val___1 interface{} = val
					var n interface{} = idx
					_, _ = val___1, n
					for {
						if n.(float64) < cnt {
							{
								var nval = func() interface{} {
									var G__4085 = val___1
									var G__4086 = Aget_(arr, n.(float64))
									_, _ = G__4085, G__4086
									return f.(CljsCoreIFn).X_invoke_Arity2(G__4085, G__4086)
								}()
								_ = nval
								if Reduced_QMARK_.Arity1IB(nval) {
									return Deref.X_invoke_Arity1(nval)
								} else {
									val___1, n = nval, (n.(float64) + float64(1))
									continue
								}
							}
						} else {
							return val___1
						}
					}
				}
			}
		})
	}(&AFn{})

	Counted_QMARK_ = func(counted_QMARK_ *AFn) *AFn {
		return Fn(counted_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreICounted)(nil)).Elem())
		})
	}(&AFn{})

	Indexed_QMARK_ = func(indexed_QMARK_ *AFn) *AFn {
		return Fn(indexed_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreIIndexed)(nil)).Elem())
		})
	}(&AFn{})

	X__GT_IndexedSeq = func(__GT_IndexedSeq *AFn) *AFn {
		return Fn(__GT_IndexedSeq, 2, func(arr interface{}, i interface{}) interface{} {
			return (&CljsCoreIndexedSeq{arr, i})
		})
	}(&AFn{})

	Prim_seq = func(prim_seq *AFn) *AFn {
		return Fn(prim_seq, 2, func(prim interface{}) interface{} {
			return prim_seq.X_invoke_Arity2(prim, float64(0))
		}, func(prim interface{}, i interface{}) interface{} {
			if i.(float64) < Alength_(prim) {
				return (&CljsCoreIndexedSeq{prim, i})
			} else {
				return nil
			}
		})
	}(&AFn{})

	Array_seq = func(array_seq *AFn) *AFn {
		return Fn(array_seq, 2, func(array interface{}) interface{} {
			return Prim_seq.X_invoke_Arity2(array, float64(0))
		}, func(array interface{}, i interface{}) interface{} {
			return Prim_seq.X_invoke_Arity2(array, i)
		})
	}(&AFn{})

	X__GT_RSeq = func(__GT_RSeq *AFn) *AFn {
		return Fn(__GT_RSeq, 3, func(ci interface{}, i interface{}, meta interface{}) interface{} {
			return (&CljsCoreRSeq{ci, i, meta})
		})
	}(&AFn{})

	Second = func(second *AFn) *AFn {
		return Fn(second, 1, func(coll interface{}) interface{} {
			return First.X_invoke_Arity1(Next.Arity1IQ(coll))
		})
	}(&AFn{})

	Ffirst = func(ffirst *AFn) *AFn {
		return Fn(ffirst, 1, func(coll interface{}) interface{} {
			return First.X_invoke_Arity1(First.X_invoke_Arity1(coll))
		})
	}(&AFn{})

	Nfirst = func(nfirst *AFn) *AFn {
		return Fn(nfirst, 1, func(coll interface{}) interface{} {
			return Next.Arity1IQ(First.X_invoke_Arity1(coll))
		})
	}(&AFn{})

	Fnext = func(fnext *AFn) *AFn {
		return Fn(fnext, 1, func(coll interface{}) interface{} {
			return First.X_invoke_Arity1(Next.Arity1IQ(coll))
		})
	}(&AFn{})

	Nnext = func(nnext *AFn) *AFn {
		return Fn(nnext, 1, func(coll interface{}) interface{} {
			return Next.Arity1IQ(Next.Arity1IQ(coll))
		})
	}(&AFn{})

	Last = func(last *AFn) *AFn {
		return Fn(last, 1, func(s interface{}) interface{} {
			for {
				{
					var sn = Next.Arity1IQ(s)
					_ = sn
					if !(Nil_(sn)) {
						s = sn
						continue
					} else {
						return First.X_invoke_Arity1(s)
					}
				}
			}
		})
	}(&AFn{})

	Conj = func(conj *AFn) *AFn {
		return Fn(conj, 2, func() interface{} {
			return CljsCorePersistentVector_EMPTY
		}, func(coll interface{}) interface{} {
			return coll
		}, func(coll interface{}, x interface{}) interface{} {
			if !(Nil_(coll)) {
				return coll.(CljsCoreICollection).X_conj_Arity2(x)
			} else {
				return CljsCoreList_EMPTY.X_conj_Arity2(x)
			}
		}, func(coll_x_xs__ ...interface{}) interface{} {
			var coll = coll_x_xs__[0]
			var x = coll_x_xs__[1]
			var xs = Seq.Arity1IQ(coll_x_xs__[2])
			_, _, _ = coll, x, xs
			for {
				if Truth_(xs) {
					coll, x, xs = conj.X_invoke_Arity2(coll, x), First.X_invoke_Arity1(xs), Next.Arity1IQ(xs)
					continue
				} else {
					return conj.X_invoke_Arity2(coll, x)
				}
			}
		})
	}(&AFn{})

	Empty = func(empty *AFn) *AFn {
		return Fn(empty, 1, func(coll interface{}) interface{} {
			if Nil_(coll) {
				return nil
			} else {
				return coll.(CljsCoreIEmptyableCollection).X_empty_Arity1()
			}
		})
	}(&AFn{})

	Accumulating_seq_count = func(accumulating_seq_count *AFn) *AFn {
		return Fn(accumulating_seq_count, 1, func(coll interface{}) interface{} {
			{
				var s interface{} = Seq.Arity1IQ(coll)
				var acc = float64(0)
				_, _ = s, acc
				for {
					if Counted_QMARK_.Arity1IB(s) {
						return (acc + s.(CljsCoreICounted).X_count_Arity1())
					} else {
						s, acc = Next.Arity1IQ(s), (acc + float64(1))
						continue
					}
				}
			}
		})
	}(&AFn{})

	Count = func(count *AFn) *AFn {
		return Fn(count, 1, func(coll interface{}) interface{} {
			if !(Nil_(coll)) {
				if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreICounted)(nil)).Elem()) {
					return coll.(CljsCoreICounted).X_count_Arity1()
				} else {
					if Value_(coll).Kind() == reflect.Slice {
						return Alength_(coll)
					} else {
						if Value_(coll).Kind() == reflect.String {
							return Alength_(coll)
						} else {
							if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreICounted)(nil)).Elem()) {
								return coll.(CljsCoreICounted).X_count_Arity1()
							} else {
								return Accumulating_seq_count.X_invoke_Arity1(coll).(float64)

							}
						}
					}
				}
			} else {
				return float64(0)
			}
		})
	}(&AFn{})

	Linear_traversal_nth = func(linear_traversal_nth *AFn) *AFn {
		return Fn(linear_traversal_nth, 3, func(coll interface{}, n interface{}) interface{} {
			for {
				if Nil_(coll) {
					panic((&js.Error{"Index out of bounds"}))
				} else {
					if n.(float64) == float64(0) {
						if Truth_(Seq.Arity1IQ(coll)) {
							return First.X_invoke_Arity1(coll)
						} else {
							panic((&js.Error{"Index out of bounds"}))
						}
					} else {
						if Indexed_QMARK_.Arity1IB(coll) {
							return coll.(CljsCoreIIndexed).X_nth_Arity2(n)
						} else {
							if Truth_(Seq.Arity1IQ(coll)) {
								coll, n = Next.Arity1IQ(coll), (n.(float64) - float64(1))
								continue
							} else {
								panic((&js.Error{"Index out of bounds"}))

							}
						}
					}
				}
			}
		}, func(coll interface{}, n interface{}, not_found interface{}) interface{} {
			for {
				if Nil_(coll) {
					return not_found
				} else {
					if n.(float64) == float64(0) {
						if Truth_(Seq.Arity1IQ(coll)) {
							return First.X_invoke_Arity1(coll)
						} else {
							return not_found
						}
					} else {
						if Indexed_QMARK_.Arity1IB(coll) {
							return coll.(CljsCoreIIndexed).X_nth_Arity3(n, not_found)
						} else {
							if Truth_(Seq.Arity1IQ(coll)) {
								coll, n, not_found = Next.Arity1IQ(coll), (n.(float64) - float64(1)), not_found
								continue
							} else {
								return not_found

							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Nth = func(nth *AFn) *AFn {
		return Fn(nth, 3, func(coll interface{}, n interface{}) interface{} {
			if !(Value_(n).Kind() == reflect.Float64) {
				panic((&js.Error{"index argument to nth must be a number"}))
			} else {
				if Nil_(coll) {
					return coll
				} else {
					if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreIIndexed)(nil)).Elem()) {
						return coll.(CljsCoreIIndexed).X_nth_Arity2(n)
					} else {
						if Value_(coll).Kind() == reflect.Slice {
							if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
								return Aget_(coll, n.(float64))
							} else {
								return nil
							}
						} else {
							if Value_(coll).Kind() == reflect.String {
								if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
									return Aget_(coll, n.(float64))
								} else {
									return nil
								}
							} else {
								if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreIIndexed)(nil)).Elem()) {
									return coll.(CljsCoreIIndexed).X_nth_Arity2(n)
								} else {
									if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreISeq)(nil)).Elem()) {
										return Linear_traversal_nth.X_invoke_Arity2(coll, n)
									} else {
										panic((&js.Error{("nth not supported on this type " + Str.X_invoke_Arity1(Type__GT_str.X_invoke_Arity1(Type_.X_invoke_Arity1(coll))).(string))}))

									}
								}
							}
						}
					}
				}
			}
		}, func(coll interface{}, n interface{}, not_found interface{}) interface{} {
			if !(Value_(n).Kind() == reflect.Float64) {
				panic((&js.Error{"index argument to nth must be a number."}))
			} else {
				if Nil_(coll) {
					return not_found
				} else {
					if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreIIndexed)(nil)).Elem()) {
						return coll.(CljsCoreIIndexed).X_nth_Arity3(n, not_found)
					} else {
						if Value_(coll).Kind() == reflect.Slice {
							if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
								return Aget_(coll, n.(float64))
							} else {
								return not_found
							}
						} else {
							if Value_(coll).Kind() == reflect.String {
								if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
									return Aget_(coll, n.(float64))
								} else {
									return not_found
								}
							} else {
								if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreIIndexed)(nil)).Elem()) {
									return coll.(CljsCoreIIndexed).X_nth_Arity2(n)
								} else {
									if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreISeq)(nil)).Elem()) {
										return Linear_traversal_nth.X_invoke_Arity3(coll, n, not_found)
									} else {
										panic((&js.Error{("nth not supported on this type " + Str.X_invoke_Arity1(Type__GT_str.X_invoke_Arity1(Type_.X_invoke_Arity1(coll))).(string))}))

									}
								}
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Assoc = func(assoc *AFn) *AFn {
		return Fn(assoc, 3, func(coll interface{}, k interface{}, v interface{}) interface{} {
			if !(Nil_(coll)) {
				return coll.(CljsCoreIAssociative).X_assoc_Arity3(k, v)
			} else {
				return CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{k}, []interface{}{v})
			}
		}, func(coll_k_v_kvs__ ...interface{}) interface{} {
			var coll = coll_k_v_kvs__[0]
			var k = coll_k_v_kvs__[1]
			var v = coll_k_v_kvs__[2]
			var kvs = Seq.Arity1IQ(coll_k_v_kvs__[3])
			_, _, _, _ = coll, k, v, kvs
			for {
				{
					var ret = assoc.X_invoke_Arity3(coll, k, v)
					_ = ret
					if Truth_(kvs) {
						coll, k, v, kvs = ret, First.X_invoke_Arity1(kvs), Second.X_invoke_Arity1(kvs), Seq_(Nnext.X_invoke_Arity1(kvs))
						continue
					} else {
						return ret
					}
				}
			}
		})
	}(&AFn{})

	Dissoc = func(dissoc *AFn) *AFn {
		return Fn(dissoc, 2, func(coll interface{}) interface{} {
			return coll
		}, func(coll interface{}, k interface{}) interface{} {
			if Nil_(coll) {
				return nil
			} else {
				return coll.(CljsCoreIMap).X_dissoc_Arity2(k)
			}
		}, func(coll_k_ks__ ...interface{}) interface{} {
			var coll = coll_k_ks__[0]
			var k = coll_k_ks__[1]
			var ks = Seq.Arity1IQ(coll_k_ks__[2])
			_, _, _ = coll, k, ks
			for {
				if Nil_(coll) {
					return nil
				} else {
					{
						var ret = dissoc.X_invoke_Arity2(coll, k)
						_ = ret
						if Truth_(ks) {
							coll, k, ks = ret, First.X_invoke_Arity1(ks), Next.Arity1IQ(ks)
							continue
						} else {
							return ret
						}
					}
				}
			}
		})
	}(&AFn{})

	Fn_QMARK_ = func(fn_QMARK_ *AFn) *AFn {
		return Fn(fn_QMARK_, 1, func(f interface{}) bool {
			{
				var or__171__auto__ = func() interface{} {
					var G__4101 = f
					_ = G__4101
					return Truth_(Native_invoke_func.X_invoke_Arity2(goog.IsFunction, []interface{}{G__4101}))
				}()
				_ = or__171__auto__
				if Truth_(or__171__auto__) {
					return or__171__auto__.(bool)
				} else {
					return Value_(f).Type().Implements(reflect.TypeOf((*CljsCoreFn)(nil)).Elem())
				}
			}
		})
	}(&AFn{})

	X__GT_MetaFn = func(__GT_MetaFn *AFn) *AFn {
		return Fn(__GT_MetaFn, 2, func(afn interface{}, meta interface{}) interface{} {
			return (&CljsCoreMetaFn{afn, meta})
		})
	}(&AFn{})

	With_meta = func(with_meta *AFn) *AFn {
		return Fn(with_meta, 2, func(o interface{}, meta interface{}) interface{} {
			if (Fn_QMARK_.Arity1IB(o)) && (!(Value_(o).Type().Implements(reflect.TypeOf((*CljsCoreIWithMeta)(nil)).Elem()))) {
				return (&CljsCoreMetaFn{o, meta})
			} else {
				if Nil_(o) {
					return nil
				} else {
					return o.(CljsCoreIWithMeta).X_with_meta_Arity2(meta)
				}
			}
		})
	}(&AFn{})

	Meta = func(meta *AFn) *AFn {
		return Fn(meta, 1, func(o interface{}) interface{} {
			if (!(Nil_(o))) && (Value_(o).Type().Implements(reflect.TypeOf((*CljsCoreIMeta)(nil)).Elem())) {
				return o.(CljsCoreIMeta).X_meta_Arity1()
			} else {
				return nil
			}
		})
	}(&AFn{})

	Peek = func(peek *AFn) *AFn {
		return Fn(peek, 1, func(coll interface{}) interface{} {
			if Nil_(coll) {
				return nil
			} else {
				return coll.(CljsCoreIStack).X_peek_Arity1()
			}
		})
	}(&AFn{})

	Pop = func(pop *AFn) *AFn {
		return Fn(pop, 1, func(coll interface{}) interface{} {
			if Nil_(coll) {
				return nil
			} else {
				return coll.(CljsCoreIStack).X_pop_Arity1()
			}
		})
	}(&AFn{})

	Disj = func(disj *AFn) *AFn {
		return Fn(disj, 2, func(coll interface{}) interface{} {
			return coll
		}, func(coll interface{}, k interface{}) interface{} {
			if Nil_(coll) {
				return nil
			} else {
				return coll.(CljsCoreISet).X_disjoin_Arity2(k)
			}
		}, func(coll_k_ks__ ...interface{}) interface{} {
			var coll = coll_k_ks__[0]
			var k = coll_k_ks__[1]
			var ks = Seq.Arity1IQ(coll_k_ks__[2])
			_, _, _ = coll, k, ks
			for {
				if Nil_(coll) {
					return nil
				} else {
					{
						var ret = disj.X_invoke_Arity2(coll, k)
						_ = ret
						if Truth_(ks) {
							coll, k, ks = ret, First.X_invoke_Arity1(ks), Next.Arity1IQ(ks)
							continue
						} else {
							return ret
						}
					}
				}
			}
		})
	}(&AFn{})

	Empty_QMARK_ = func(empty_QMARK_ *AFn) *AFn {
		return Fn(empty_QMARK_, 1, func(coll interface{}) bool {
			return (Nil_(coll)) || (Not.Arity1IB(Seq.Arity1IQ(coll)))
		})
	}(&AFn{})

	Coll_QMARK_ = func(coll_QMARK_ *AFn) *AFn {
		return Fn(coll_QMARK_, 1, func(x interface{}) bool {
			if Nil_(x) {
				return false
			} else {
				return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreICollection)(nil)).Elem())
			}
		})
	}(&AFn{})

	Set_QMARK_ = func(set_QMARK_ *AFn) *AFn {
		return Fn(set_QMARK_, 1, func(x interface{}) bool {
			if Nil_(x) {
				return false
			} else {
				return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreISet)(nil)).Elem())
			}
		})
	}(&AFn{})

	Associative_QMARK_ = func(associative_QMARK_ *AFn) *AFn {
		return Fn(associative_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreIAssociative)(nil)).Elem())
		})
	}(&AFn{})

	Sequential_QMARK_ = func(sequential_QMARK_ *AFn) *AFn {
		return Fn(sequential_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreISequential)(nil)).Elem())
		})
	}(&AFn{})

	Sorted_QMARK_ = func(sorted_QMARK_ *AFn) *AFn {
		return Fn(sorted_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreISorted)(nil)).Elem())
		})
	}(&AFn{})

	Reduceable_QMARK_ = func(reduceable_QMARK_ *AFn) *AFn {
		return Fn(reduceable_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreIReduce)(nil)).Elem())
		})
	}(&AFn{})

	Map_QMARK_ = func(map_QMARK_ *AFn) *AFn {
		return Fn(map_QMARK_, 1, func(x interface{}) bool {
			if Nil_(x) {
				return false
			} else {
				return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreIMap)(nil)).Elem())
			}
		})
	}(&AFn{})

	Vector_QMARK_ = func(vector_QMARK_ *AFn) *AFn {
		return Fn(vector_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreIVector)(nil)).Elem())
		})
	}(&AFn{})

	Chunked_seq_QMARK_ = func(chunked_seq_QMARK_ *AFn) *AFn {
		return Fn(chunked_seq_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreIChunkedSeq)(nil)).Elem())
		})
	}(&AFn{})

	Js_delete = func(js_delete *AFn) *AFn {
		return Fn(js_delete, 2, func(obj interface{}, key interface{}) interface{} {
			return func(obj, key interface{}) interface{} { delete(obj.(map[string]interface{}), key.(string)); return obj }(obj, key)
		})
	}(&AFn{})

	Array_copy = func(array_copy *AFn) *AFn {
		return Fn(array_copy, 5, func(from interface{}, i interface{}, to interface{}, j interface{}, len interface{}) interface{} {
			{
				var i___1 interface{} = i
				var j___1 interface{} = j
				var len___1 interface{} = len
				_, _, _ = i___1, j___1, len___1
				for {
					if len___1.(float64) == float64(0) {
						return to
					} else {
						to.([]interface{})[int(j___1.(float64))] = Aget_(from, i___1.(float64))
						i___1, j___1, len___1 = (i___1.(float64) + float64(1)), (j___1.(float64) + float64(1)), (len___1.(float64) - float64(1))
						continue
					}
				}
			}
		})
	}(&AFn{})

	Array_copy_downward = func(array_copy_downward *AFn) *AFn {
		return Fn(array_copy_downward, 5, func(from interface{}, i interface{}, to interface{}, j interface{}, len interface{}) interface{} {
			{
				var i___1 = (i.(float64) + (len.(float64) - float64(1)))
				var j___1 = (j.(float64) + (len.(float64) - float64(1)))
				var len___1 interface{} = len
				_, _, _ = i___1, j___1, len___1
				for {
					if len___1.(float64) == float64(0) {
						return to
					} else {
						to.([]interface{})[int(j___1)] = Aget_(from, i___1)
						i___1, j___1, len___1 = (i___1 - float64(1)), (j___1 - float64(1)), (len___1.(float64) - float64(1))
						continue
					}
				}
			}
		})
	}(&AFn{})

	Lookup_sentinel = true

	False_QMARK_ = func(false_QMARK_ *AFn) *AFn {
		return Fn(false_QMARK_, 1, func(x interface{}) bool {
			return x == false
		})
	}(&AFn{})

	True_QMARK_ = func(true_QMARK_ *AFn) *AFn {
		return Fn(true_QMARK_, 1, func(x interface{}) bool {
			return x == true
		})
	}(&AFn{})

	Undefined_QMARK_ = func(undefined_QMARK_ *AFn) *AFn {
		return Fn(undefined_QMARK_, 1, func(x interface{}) bool {
			return (nil == x)
		})
	}(&AFn{})

	Seq_QMARK_ = func(seq_QMARK_ *AFn) *AFn {
		return Fn(seq_QMARK_, 1, func(s interface{}) bool {
			if Nil_(s) {
				return false
			} else {
				return Value_(s).Type().Implements(reflect.TypeOf((*CljsCoreISeq)(nil)).Elem())
			}
		})
	}(&AFn{})

	Seqable_QMARK_ = func(seqable_QMARK_ *AFn) *AFn {
		return Fn(seqable_QMARK_, 1, func(s interface{}) bool {
			return Value_(s).Type().Implements(reflect.TypeOf((*CljsCoreISeqable)(nil)).Elem())
		})
	}(&AFn{})

	Boolean = func(boolean *AFn) *AFn {
		return Fn(boolean, 1, func(x interface{}) bool {
			if Truth_(x) {
				return true
			} else {
				return false
			}
		})
	}(&AFn{})

	Ifn_QMARK_ = func(ifn_QMARK_ *AFn) *AFn {
		return Fn(ifn_QMARK_, 1, func(f interface{}) bool {
			return (Fn_QMARK_.Arity1IB(f)) || (Value_(f).Type().Implements(reflect.TypeOf((*CljsCoreIFn)(nil)).Elem()))
		})
	}(&AFn{})

	Contains_QMARK_ = func(contains_QMARK_ *AFn) *AFn {
		return Fn(contains_QMARK_, 2, func(coll interface{}, v interface{}) bool {
			if reflect.DeepEqual(Get.X_invoke_Arity3(coll, v, Lookup_sentinel), Lookup_sentinel) {
				return false
			} else {
				return true
			}
		})
	}(&AFn{})

	Find = func(find *AFn) *AFn {
		return Fn(find, 2, func(coll interface{}, k interface{}) interface{} {
			if (!(Nil_(coll))) && (Associative_QMARK_.Arity1IB(coll)) && (Contains_QMARK_.Arity2IIB(coll, k)) {
				return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{k, Get.X_invoke_Arity2(coll, k)}, nil})
			} else {
				return nil
			}
		})
	}(&AFn{})

	Distinct_QMARK_ = func(distinct_QMARK_ *AFn) *AFn {
		return Fn(distinct_QMARK_, 2, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return !(X_EQ_.Arity2IIB(x, y))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			if !(X_EQ_.Arity2IIB(x, y)) {
				{
					var s interface{} = CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{x, y}, true).(*CljsCorePersistentHashSet)
					var xs interface{} = more
					_, _ = s, xs
					for {
						{
							var x___1 = First.X_invoke_Arity1(xs)
							var etc = Next.Arity1IQ(xs)
							_, _ = x___1, etc
							if Truth_(xs) {
								if Contains_QMARK_.Arity2IIB(s, x___1) {
									return false
								} else {
									s, xs = Conj.X_invoke_Arity2(s, x___1), etc
									continue
								}
							} else {
								return true
							}
						}
					}
				}
			} else {
				return false
			}
		})
	}(&AFn{})

	Sequence = func(sequence *AFn) *AFn {
		return Fn(sequence, 1, func(coll interface{}) CljsCoreISeq {
			if Seq_QMARK_.Arity1IB(coll) {
				return coll.(CljsCoreISeq)
			} else {
				{
					var or__171__auto__ = Seq.Arity1IQ(coll)
					_ = or__171__auto__
					if Truth_(or__171__auto__) {
						return or__171__auto__
					} else {
						return CljsCoreIEmptyList(CljsCoreList_EMPTY)
					}
				}
			}
		})
	}(&AFn{})

	Compare = func(compare *AFn) *AFn {
		return Fn(compare, 2, func(x interface{}, y interface{}) float64 {
			if reflect.DeepEqual(x, y) {
				return float64(0)
			} else {
				if Nil_(x) {
					return float64(-1)
				} else {
					if Nil_(y) {
						return float64(1)
					} else {
						if reflect.DeepEqual(Type_.X_invoke_Arity1(x), Type_.X_invoke_Arity1(y)) {
							if Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreIComparable)(nil)).Elem()) {
								return x.(CljsCoreIComparable).X_compare_Arity2(y)
							} else {
								{
									var G__4552 = x
									var G__4553 = y
									_, _ = G__4552, G__4553
									return Native_invoke_func.X_invoke_Arity2(goog_array.DefaultCompare, []interface{}{G__4552, G__4553}).(float64)
								}
							}
						} else {
							panic((&js.Error{"compare on non-nil objects of different types"}))

						}
					}
				}
			}
		})
	}(&AFn{})

	Compare_indexed = func(compare_indexed *AFn) *AFn {
		return Fn(compare_indexed, 4, func(xs interface{}, ys interface{}) interface{} {
			{
				var xl = Count.X_invoke_Arity1(xs).(float64)
				var yl = Count.X_invoke_Arity1(ys).(float64)
				_, _ = xl, yl
				if xl < yl {
					return float64(-1)
				} else {
					if xl > yl {
						return float64(1)
					} else {
						return compare_indexed.X_invoke_Arity4(xs, ys, xl, float64(0)).(float64)

					}
				}
			}
		}, func(xs interface{}, ys interface{}, len interface{}, n interface{}) interface{} {
			for {
				{
					var d = Compare.Arity2IIF(Nth.X_invoke_Arity2(xs, n), Nth.X_invoke_Arity2(ys, n))
					_ = d
					if (d == float64(0)) && ((n.(float64) + float64(1)) < len.(float64)) {
						xs, ys, len, n = xs, ys, len, (n.(float64) + float64(1))
						continue
					} else {
						return d
					}
				}
			}
		})
	}(&AFn{})

	Fn__GT_comparator = func(fn__GT_comparator *AFn) *AFn {
		return Fn(fn__GT_comparator, 1, func(f interface{}) interface{} {
			if X_EQ_.Arity2IIB(f, Compare) {
				return Compare
			} else {
				return func(G__4566 *AFn) *AFn {
					return Fn(G__4566, 2, func(x interface{}, y interface{}) interface{} {
						{
							var r = func() interface{} {
								var G__4562 = x
								var G__4563 = y
								_, _ = G__4562, G__4563
								return f.(CljsCoreIFn).X_invoke_Arity2(G__4562, G__4563)
							}()
							_ = r
							if Value_(r).Kind() == reflect.Float64 {
								return r
							} else {
								if Truth_(r) {
									return float64(-1)
								} else {
									if Truth_(func() interface{} {
										var G__4564 = y
										var G__4565 = x
										_, _ = G__4564, G__4565
										return f.(CljsCoreIFn).X_invoke_Arity2(G__4564, G__4565)
									}()) {
										return float64(1)
									} else {
										return float64(0)
									}
								}
							}
						}
					})
				}(&AFn{})
			}
		})
	}(&AFn{})

	Sort_by = func(sort_by *AFn) *AFn {
		return Fn(sort_by, 3, func(keyfn interface{}, coll interface{}) interface{} {
			return sort_by.X_invoke_Arity3(keyfn, Compare, coll)
		}, func(keyfn interface{}, comp interface{}, coll interface{}) interface{} {
			return Sort.X_invoke_Arity2(func(G__4580 *AFn) *AFn {
				return Fn(G__4580, 2, func(x interface{}, y interface{}) interface{} {
					return Fn__GT_comparator.X_invoke_Arity1(comp).(CljsCoreIFn).X_invoke_Arity2(func() interface{} {
						var G__4578 = x
						_ = G__4578
						return keyfn.(CljsCoreIFn).X_invoke_Arity1(G__4578)
					}(), func() interface{} {
						var G__4579 = y
						_ = G__4579
						return keyfn.(CljsCoreIFn).X_invoke_Arity1(G__4579)
					}())
				})
			}(&AFn{}), coll)
		})
	}(&AFn{})

	Seq_reduce = func(seq_reduce *AFn) *AFn {
		return Fn(seq_reduce, 3, func(f interface{}, coll interface{}) interface{} {
			{
				var temp__4220__auto__ = Seq.Arity1IQ(coll)
				_ = temp__4220__auto__
				if Truth_(temp__4220__auto__) {
					{
						var s = temp__4220__auto__
						_ = s
						return Reduce.X_invoke_Arity3(f, First.X_invoke_Arity1(s), Next.Arity1IQ(s))
					}
				} else {
					{
						return f.(CljsCoreIFn).X_invoke_Arity0()
					}
				}
			}
		}, func(f interface{}, val interface{}, coll interface{}) interface{} {
			{
				var val___1 interface{} = val
				var coll___1 interface{} = Seq.Arity1IQ(coll)
				_, _ = val___1, coll___1
				for {
					if Truth_(coll___1) {
						{
							var nval = func() interface{} {
								var G__4583 = val___1
								var G__4584 = First.X_invoke_Arity1(coll___1)
								_, _ = G__4583, G__4584
								return f.(CljsCoreIFn).X_invoke_Arity2(G__4583, G__4584)
							}()
							_ = nval
							if Reduced_QMARK_.Arity1IB(nval) {
								return Deref.X_invoke_Arity1(nval)
							} else {
								val___1, coll___1 = nval, Next.Arity1IQ(coll___1)
								continue
							}
						}
					} else {
						return val___1
					}
				}
			}
		})
	}(&AFn{})

	Shuffle = func(shuffle *AFn) *AFn {
		return Fn(shuffle, 1, func(coll interface{}) interface{} {
			{
				var a = To_array.X_invoke_Arity1(coll).([]interface{})
				_ = a
				{
					var G__4586_4587 = a
					_ = G__4586_4587
					Native_invoke_func.X_invoke_Arity2(goog_array.Shuffle, []interface{}{G__4586_4587})
				}
				return Vec.X_invoke_Arity1(a)
			}
		})
	}(&AFn{})

	Reduce = func(reduce *AFn) *AFn {
		return Fn(reduce, 3, func(f interface{}, coll interface{}) interface{} {
			if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreIReduce)(nil)).Elem()) {
				return coll.(CljsCoreIReduce).X_reduce_Arity2(f)
			} else {
				if Value_(coll).Kind() == reflect.Slice {
					return Array_reduce.X_invoke_Arity2(coll, f)
				} else {
					if Value_(coll).Kind() == reflect.String {
						return Array_reduce.X_invoke_Arity2(coll, f)
					} else {
						if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreIReduce)(nil)).Elem()) {
							return coll.(CljsCoreIReduce).X_reduce_Arity2(f)
						} else {
							return Seq_reduce.X_invoke_Arity2(f, coll)

						}
					}
				}
			}
		}, func(f interface{}, val interface{}, coll interface{}) interface{} {
			if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreIReduce)(nil)).Elem()) {
				return coll.(CljsCoreIReduce).X_reduce_Arity3(f, val)
			} else {
				if Value_(coll).Kind() == reflect.Slice {
					return Array_reduce.X_invoke_Arity3(coll, f, val)
				} else {
					if Value_(coll).Kind() == reflect.String {
						return Array_reduce.X_invoke_Arity3(coll, f, val)
					} else {
						if Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreIReduce)(nil)).Elem()) {
							return coll.(CljsCoreIReduce).X_reduce_Arity3(f, val)
						} else {
							return Seq_reduce.X_invoke_Arity3(f, val, coll)

						}
					}
				}
			}
		})
	}(&AFn{})

	Reduce_kv = func(reduce_kv *AFn) *AFn {
		return Fn(reduce_kv, 3, func(f interface{}, init interface{}, coll interface{}) interface{} {
			if !(Nil_(coll)) {
				return coll.(CljsCoreIKVReduce).X_kv_reduce_Arity3(f, init)
			} else {
				return init
			}
		})
	}(&AFn{})

	Completing = func(completing *AFn) *AFn {
		return Fn(completing, 2, func(f interface{}) interface{} {
			return completing.X_invoke_Arity2(f, Identity).(CljsCoreIFn)
		}, func(f interface{}, cf interface{}) interface{} {
			return func(G__4596 *AFn) *AFn {
				return Fn(G__4596, 2, func() interface{} {
					{
						return f.(CljsCoreIFn).X_invoke_Arity0()
					}
				}, func(x interface{}) interface{} {
					{
						var G__4593 = x
						_ = G__4593
						return cf.(CljsCoreIFn).X_invoke_Arity1(G__4593)
					}
				}, func(x interface{}, y interface{}) interface{} {
					{
						var G__4594 = x
						var G__4595 = y
						_, _ = G__4594, G__4595
						return f.(CljsCoreIFn).X_invoke_Arity2(G__4594, G__4595)
					}
				})
			}(&AFn{})
		})
	}(&AFn{})

	Transduce = func(transduce *AFn) *AFn {
		return Fn(transduce, 4, func(xform interface{}, f interface{}, coll interface{}) interface{} {
			return transduce.X_invoke_Arity4(xform, f, func() interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity0()
			}(), coll)
		}, func(xform interface{}, f interface{}, init interface{}, coll interface{}) interface{} {
			{
				var f___1 = func() interface{} {
					var G__4603 = f
					_ = G__4603
					return xform.(CljsCoreIFn).X_invoke_Arity1(G__4603)
				}()
				var ret = Reduce.X_invoke_Arity3(f___1, init, coll)
				_, _ = f___1, ret
				{
					var G__4604 = ret
					_ = G__4604
					return f___1.(CljsCoreIFn).X_invoke_Arity1(G__4604)
				}
			}
		})
	}(&AFn{})

	X_PLUS_ = func(_PLUS_ *AFn) *AFn {
		return Fn(_PLUS_, 2, func() float64 {
			return float64(0)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) + y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(_PLUS_, (x.(float64) + y.(float64)), more)
		})
	}(&AFn{})

	X_ = func(___ *AFn) *AFn {
		return Fn(___, 2, func(x interface{}) float64 {
			return (-x.(float64))
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) - y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(___, (x.(float64) - y.(float64)), more)
		})
	}(&AFn{})

	X_STAR_ = func(_STAR_ *AFn) *AFn {
		return Fn(_STAR_, 2, func() float64 {
			return float64(1)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) * y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(_STAR_, (x.(float64) * y.(float64)), more)
		})
	}(&AFn{})

	X_SLASH_ = func(_SLASH_ *AFn) *AFn {
		return Fn(_SLASH_, 2, func(x interface{}) float64 {
			return _SLASH_.Arity2IIF(float64(1), x)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) / y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(_SLASH_, _SLASH_.Arity2IIF(x, y), more)
		})
	}(&AFn{})

	X_LT_ = func(_LT_ *AFn) *AFn {
		return Fn(_LT_, 2, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return (x.(float64) < y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			for {
				if x.(float64) < y.(float64) {
					if Truth_(Next.Arity1IQ(more)) {
						x, y, more = y, First.X_invoke_Arity1(more), Next.Arity1IQ(more)
						continue
					} else {
						return (y.(float64) < First.X_invoke_Arity1(more).(float64))
					}
				} else {
					return false
				}
			}
		})
	}(&AFn{})

	X_LT__EQ_ = func(_LT__EQ_ *AFn) *AFn {
		return Fn(_LT__EQ_, 2, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return (x.(float64) <= y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			for {
				if x.(float64) <= y.(float64) {
					if Truth_(Next.Arity1IQ(more)) {
						x, y, more = y, First.X_invoke_Arity1(more), Next.Arity1IQ(more)
						continue
					} else {
						return (y.(float64) <= First.X_invoke_Arity1(more).(float64))
					}
				} else {
					return false
				}
			}
		})
	}(&AFn{})

	X_GT_ = func(_GT_ *AFn) *AFn {
		return Fn(_GT_, 2, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return (x.(float64) > y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			for {
				if x.(float64) > y.(float64) {
					if Truth_(Next.Arity1IQ(more)) {
						x, y, more = y, First.X_invoke_Arity1(more), Next.Arity1IQ(more)
						continue
					} else {
						return (y.(float64) > First.X_invoke_Arity1(more).(float64))
					}
				} else {
					return false
				}
			}
		})
	}(&AFn{})

	X_GT__EQ_ = func(_GT__EQ_ *AFn) *AFn {
		return Fn(_GT__EQ_, 2, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return (x.(float64) >= y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			for {
				if x.(float64) >= y.(float64) {
					if Truth_(Next.Arity1IQ(more)) {
						x, y, more = y, First.X_invoke_Arity1(more), Next.Arity1IQ(more)
						continue
					} else {
						return (y.(float64) >= First.X_invoke_Arity1(more).(float64))
					}
				} else {
					return false
				}
			}
		})
	}(&AFn{})

	Dec = func(dec *AFn) *AFn {
		return Fn(dec, 1, func(x interface{}) interface{} {
			return (x.(float64) - float64(1))
		})
	}(&AFn{})

	Max = func(max *AFn) *AFn {
		return Fn(max, 2, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(x.(float64), y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(max, func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(x.(float64), y.(float64)), more)
		})
	}(&AFn{})

	Min = func(min *AFn) *AFn {
		return Fn(min, 2, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(x.(float64), y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(min, func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(x.(float64), y.(float64)), more)
		})
	}(&AFn{})

	Byte_ = func(byte_ *AFn) *AFn {
		return Fn(byte_, 1, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Short = func(short *AFn) *AFn {
		return Fn(short, 1, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Float_ = func(float_ *AFn) *AFn {
		return Fn(float_, 1, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Double = func(double *AFn) *AFn {
		return Fn(double, 1, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_byte = func(unchecked_byte *AFn) *AFn {
		return Fn(unchecked_byte, 1, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_char = func(unchecked_char *AFn) *AFn {
		return Fn(unchecked_char, 1, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_short = func(unchecked_short *AFn) *AFn {
		return Fn(unchecked_short, 1, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_float = func(unchecked_float *AFn) *AFn {
		return Fn(unchecked_float, 1, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_double = func(unchecked_double *AFn) *AFn {
		return Fn(unchecked_double, 1, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_add = func(unchecked_add *AFn) *AFn {
		return Fn(unchecked_add, 2, func() float64 {
			return float64(0)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) + y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_add, (x.(float64) + y.(float64)), more)
		})
	}(&AFn{})

	Unchecked_add_int = func(unchecked_add_int *AFn) *AFn {
		return Fn(unchecked_add_int, 2, func() float64 {
			return float64(0)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) + y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_add_int, (x.(float64) + y.(float64)), more)
		})
	}(&AFn{})

	Unchecked_dec = func(unchecked_dec *AFn) *AFn {
		return Fn(unchecked_dec, 1, func(x interface{}) interface{} {
			return (x.(float64) - float64(1))
		})
	}(&AFn{})

	Unchecked_dec_int = func(unchecked_dec_int *AFn) *AFn {
		return Fn(unchecked_dec_int, 1, func(x interface{}) interface{} {
			return (x.(float64) - float64(1))
		})
	}(&AFn{})

	Unchecked_divide_int = func(unchecked_divide_int *AFn) *AFn {
		return Fn(unchecked_divide_int, 2, func(x interface{}) float64 {
			return unchecked_divide_int.Arity2IIF(float64(1), x)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) / y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_divide_int, unchecked_divide_int.Arity2IIF(x, y), more)
		})
	}(&AFn{})

	Unchecked_inc = func(unchecked_inc *AFn) *AFn {
		return Fn(unchecked_inc, 1, func(x interface{}) interface{} {
			return (x.(float64) + float64(1))
		})
	}(&AFn{})

	Unchecked_inc_int = func(unchecked_inc_int *AFn) *AFn {
		return Fn(unchecked_inc_int, 1, func(x interface{}) interface{} {
			return (x.(float64) + float64(1))
		})
	}(&AFn{})

	Unchecked_multiply = func(unchecked_multiply *AFn) *AFn {
		return Fn(unchecked_multiply, 2, func() float64 {
			return float64(1)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) * y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_multiply, (x.(float64) * y.(float64)), more)
		})
	}(&AFn{})

	Unchecked_multiply_int = func(unchecked_multiply_int *AFn) *AFn {
		return Fn(unchecked_multiply_int, 2, func() float64 {
			return float64(1)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) * y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_multiply_int, (x.(float64) * y.(float64)), more)
		})
	}(&AFn{})

	Unchecked_negate = func(unchecked_negate *AFn) *AFn {
		return Fn(unchecked_negate, 1, func(x interface{}) interface{} {
			return (-x.(float64))
		})
	}(&AFn{})

	Unchecked_negate_int = func(unchecked_negate_int *AFn) *AFn {
		return Fn(unchecked_negate_int, 1, func(x interface{}) interface{} {
			return (-x.(float64))
		})
	}(&AFn{})

	Unchecked_remainder_int = func(unchecked_remainder_int *AFn) *AFn {
		return Fn(unchecked_remainder_int, 2, func(x interface{}, n interface{}) interface{} {
			return Mod.X_invoke_Arity2(x, n).(float64)
		})
	}(&AFn{})

	Unchecked_subtract = func(unchecked_subtract *AFn) *AFn {
		return Fn(unchecked_subtract, 2, func(x interface{}) float64 {
			return (-x.(float64))
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) - y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_subtract, (x.(float64) - y.(float64)), more)
		})
	}(&AFn{})

	Unchecked_subtract_int = func(unchecked_subtract_int *AFn) *AFn {
		return Fn(unchecked_subtract_int, 2, func(x interface{}) float64 {
			return (-x.(float64))
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) - y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_subtract_int, (x.(float64) - y.(float64)), more)
		})
	}(&AFn{})

	Fix = func(fix *AFn) *AFn {
		return Fn(fix, 1, func(q interface{}) float64 {
			if q.(float64) >= float64(0) {
				{
					var G__4615 = q
					_ = G__4615
					return Native_invoke_func.X_invoke_Arity2(Math.Floor, []interface{}{G__4615}).(float64)
				}
			} else {
				{
					var G__4616 = q
					_ = G__4616
					return Native_invoke_func.X_invoke_Arity2(Math.Ceil, []interface{}{G__4616}).(float64)
				}
			}
		})
	}(&AFn{})

	Int_ = func(int_ *AFn) *AFn {
		return Fn(int_, 1, func(x interface{}) interface{} {
			return float64((Int32_(x.(float64)) | Int32_(float64(0))))
		})
	}(&AFn{})

	Unchecked_int = func(unchecked_int *AFn) *AFn {
		return Fn(unchecked_int, 1, func(x interface{}) interface{} {
			return Fix.Arity1IF(x)
		})
	}(&AFn{})

	Long = func(long *AFn) *AFn {
		return Fn(long, 1, func(x interface{}) interface{} {
			return Fix.Arity1IF(x)
		})
	}(&AFn{})

	Unchecked_long = func(unchecked_long *AFn) *AFn {
		return Fn(unchecked_long, 1, func(x interface{}) interface{} {
			return Fix.Arity1IF(x)
		})
	}(&AFn{})

	Booleans = func(booleans *AFn) *AFn {
		return Fn(booleans, 1, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Bytes = func(bytes *AFn) *AFn {
		return Fn(bytes, 1, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Chars = func(chars *AFn) *AFn {
		return Fn(chars, 1, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Shorts = func(shorts *AFn) *AFn {
		return Fn(shorts, 1, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Ints = func(ints *AFn) *AFn {
		return Fn(ints, 1, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Floats = func(floats *AFn) *AFn {
		return Fn(floats, 1, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Doubles = func(doubles *AFn) *AFn {
		return Fn(doubles, 1, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Longs = func(longs *AFn) *AFn {
		return Fn(longs, 1, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Js_mod = func(js_mod *AFn) *AFn {
		return Fn(js_mod, 2, func(n interface{}, d interface{}) interface{} {
			return math.Mod(n.(float64), d.(float64))
		})
	}(&AFn{})

	Mod = func(mod *AFn) *AFn {
		return Fn(mod, 2, func(n interface{}, d interface{}) interface{} {
			return math.Mod((math.Mod(n.(float64), d.(float64)) + d.(float64)), d.(float64))
		})
	}(&AFn{})

	Quot = func(quot *AFn) *AFn {
		return Fn(quot, 2, func(n interface{}, d interface{}) interface{} {
			{
				var rem = math.Mod(n.(float64), d.(float64))
				_ = rem
				return Fix.Arity1IF(((n.(float64) - rem) / d.(float64)))
			}
		})
	}(&AFn{})

	Rem = func(rem *AFn) *AFn {
		return Fn(rem, 2, func(n interface{}, d interface{}) interface{} {
			{
				var q = Quot.X_invoke_Arity2(n, d).(float64)
				_ = q
				return (n.(float64) - (d.(float64) * q))
			}
		})
	}(&AFn{})

	Rand_int = func(rand_int *AFn) *AFn {
		return Fn(rand_int, 1, func(n interface{}) interface{} {
			return Fix.Arity1IF(Rand.Arity1IF(n))
		})
	}(&AFn{})

	Bit_xor = func(bit_xor *AFn) *AFn {
		return Fn(bit_xor, 2, func(x interface{}, y interface{}) interface{} {
			return float64((Int32_(x.(float64)) ^ Int32_(y.(float64))))
		})
	}(&AFn{})

	Bit_and = func(bit_and *AFn) *AFn {
		return Fn(bit_and, 2, func(x interface{}, y interface{}) interface{} {
			return float64((Int32_(x.(float64)) & Int32_(y.(float64))))
		})
	}(&AFn{})

	Bit_or = func(bit_or *AFn) *AFn {
		return Fn(bit_or, 2, func(x interface{}, y interface{}) interface{} {
			return float64((Int32_(x.(float64)) | Int32_(y.(float64))))
		})
	}(&AFn{})

	Bit_and_not = func(bit_and_not *AFn) *AFn {
		return Fn(bit_and_not, 2, func(x interface{}, y interface{}) interface{} {
			return float64((Int32_(x.(float64)) &^ Int32_(y.(float64))))
		})
	}(&AFn{})

	Bit_clear = func(bit_clear *AFn) *AFn {
		return Fn(bit_clear, 2, func(x interface{}, n interface{}) interface{} {
			return float64((Int32_(x.(float64)) &^ (Int32_(1) << UInt32_(n.(float64)))))
		})
	}(&AFn{})

	Bit_flip = func(bit_flip *AFn) *AFn {
		return Fn(bit_flip, 2, func(x interface{}, n interface{}) interface{} {
			return float64((Int32_(x.(float64)) ^ (Int32_(1) << UInt32_(n.(float64)))))
		})
	}(&AFn{})

	Bit_not = func(bit_not *AFn) *AFn {
		return Fn(bit_not, 1, func(x interface{}) interface{} {
			return float64((^Int32_(x.(float64))))
		})
	}(&AFn{})

	Bit_set = func(bit_set *AFn) *AFn {
		return Fn(bit_set, 2, func(x interface{}, n interface{}) interface{} {
			return float64((Int32_(x.(float64)) | (Int32_(1) << UInt32_(n.(float64)))))
		})
	}(&AFn{})

	Bit_test = func(bit_test *AFn) *AFn {
		return Fn(bit_test, 2, func(x interface{}, n interface{}) interface{} {
			return (Int32_(x.(float64)) & (Int32_(1) << UInt32_(n.(float64)))) != 0
		})
	}(&AFn{})

	Bit_shift_left = func(bit_shift_left *AFn) *AFn {
		return Fn(bit_shift_left, 2, func(x interface{}, n interface{}) interface{} {
			return float64((Int32_(x.(float64)) << UInt32_(n.(float64))))
		})
	}(&AFn{})

	Bit_shift_right = func(bit_shift_right *AFn) *AFn {
		return Fn(bit_shift_right, 2, func(x interface{}, n interface{}) interface{} {
			return float64((Int32_(x.(float64)) >> UInt32_(n.(float64))))
		})
	}(&AFn{})

	Bit_shift_right_zero_fill = func(bit_shift_right_zero_fill *AFn) *AFn {
		return Fn(bit_shift_right_zero_fill, 2, func(x interface{}, n interface{}) interface{} {
			return float64((UInt32_(x.(float64)) >> UInt32_(float64((32+Int32_(n.(float64)))%32))))
		})
	}(&AFn{})

	Unsigned_bit_shift_right = func(unsigned_bit_shift_right *AFn) *AFn {
		return Fn(unsigned_bit_shift_right, 2, func(x interface{}, n interface{}) interface{} {
			return float64((UInt32_(x.(float64)) >> UInt32_(float64((32+Int32_(n.(float64)))%32))))
		})
	}(&AFn{})

	Bit_count = func(bit_count *AFn) *AFn {
		return Fn(bit_count, 1, func(v interface{}) interface{} {
			{
				var v___1 = (v.(float64) - float64((Int32_(float64((Int32_(v.(float64)) >> UInt32_(float64(1))))) & Int32_(float64(1431655765)))))
				var v___2 = (float64((Int32_(v___1) & Int32_(float64(858993459)))) + float64((Int32_(float64((Int32_(v___1) >> UInt32_(float64(2))))) & Int32_(float64(858993459)))))
				_, _ = v___1, v___2
				return float64((Int32_((float64((Int32_((v___2 + float64((Int32_(v___2) >> UInt32_(float64(4)))))) & Int32_(float64(252645135)))) * float64(16843009))) >> UInt32_(float64(24))))
			}
		})
	}(&AFn{})

	X_EQ__EQ_ = func(_EQ__EQ_ *AFn) *AFn {
		return Fn(_EQ__EQ_, 2, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return x.(CljsCoreIEquiv).X_equiv_Arity2(y)
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			for {
				if _EQ__EQ_.Arity2IIB(x, y) {
					if Truth_(Next.Arity1IQ(more)) {
						x, y, more = y, First.X_invoke_Arity1(more), Next.Arity1IQ(more)
						continue
					} else {
						return _EQ__EQ_.Arity2IIB(y, First.X_invoke_Arity1(more))
					}
				} else {
					return false
				}
			}
		})
	}(&AFn{})

	Pos_QMARK_ = func(pos_QMARK_ *AFn) *AFn {
		return Fn(pos_QMARK_, 1, func(n interface{}) bool {
			return (n.(float64) > float64(0))
		})
	}(&AFn{})

	Zero_QMARK_ = func(zero_QMARK_ *AFn) *AFn {
		return Fn(zero_QMARK_, 1, func(n interface{}) bool {
			return (n.(float64) == float64(0))
		})
	}(&AFn{})

	Neg_QMARK_ = func(neg_QMARK_ *AFn) *AFn {
		return Fn(neg_QMARK_, 1, func(x interface{}) bool {
			return (x.(float64) < float64(0))
		})
	}(&AFn{})

	Nthnext = func(nthnext *AFn) *AFn {
		return Fn(nthnext, 2, func(coll interface{}, n interface{}) interface{} {
			{
				var n___1 interface{} = n
				var xs interface{} = Seq.Arity1IQ(coll)
				_, _ = n___1, xs
				for {
					if Truth_(func() interface{} {
						var and__159__auto__ = xs
						_ = and__159__auto__
						if Truth_(and__159__auto__) {
							return (n___1.(float64) > float64(0))
						} else {
							return and__159__auto__
						}
					}()) {
						n___1, xs = (n___1.(float64) - float64(1)), Next.Arity1IQ(xs)
						continue
					} else {
						return xs
					}
				}
			}
		})
	}(&AFn{})

	Str = func(str *AFn) *AFn {
		return Fn(str, 1, func() interface{} {
			return ""
		}, func(x interface{}) interface{} {
			if Nil_(x) {
				return ""
			} else {
				return fmt.Sprint(x)
			}
		}, func(x_ys__ ...interface{}) interface{} {
			var x = x_ys__[0]
			var ys = Seq.Arity1IQ(x_ys__[1])
			_, _ = x, ys
			{
				var sb interface{} = (&goog_string.StringBuffer{str.X_invoke_Arity1(x)})
				var more interface{} = ys
				_, _ = sb, more
				for {
					if Truth_(more) {
						sb, more = Native_invoke_instance_method.X_invoke_Arity3(sb, "Append", []interface{}{str.X_invoke_Arity1(First.X_invoke_Arity1(more))}), Next.Arity1IQ(more)
						continue
					} else {
						return Native_invoke_instance_method.X_invoke_Arity3(sb, "ToString", []interface{}{})
					}
				}
			}
		})
	}(&AFn{})

	Subs = func(subs *AFn) *AFn {
		return Fn(subs, 3, func(s interface{}, start interface{}) interface{} {
			return Native_invoke_instance_method.X_invoke_Arity3(s, "Substring", []interface{}{start})
		}, func(s interface{}, start interface{}, end interface{}) interface{} {
			return Native_invoke_instance_method.X_invoke_Arity3(s, "Substring", []interface{}{start, end})
		})
	}(&AFn{})

	Equiv_sequential = func(equiv_sequential *AFn) *AFn {
		return Fn(equiv_sequential, 2, func(x interface{}, y interface{}) interface{} {
			return Boolean.Arity1IB(func() interface{} {
				if Sequential_QMARK_.Arity1IB(y) {
					return func() bool {
						if (Counted_QMARK_.Arity1IB(x)) && (Counted_QMARK_.Arity1IB(y)) && (!(Count.X_invoke_Arity1(x).(float64) == Count.X_invoke_Arity1(y).(float64))) {
							return false
						} else {
							return func() bool {
								var xs interface{} = Seq.Arity1IQ(x)
								var ys interface{} = Seq.Arity1IQ(y)
								_, _ = xs, ys
								for {
									if Nil_(xs) {
										return Nil_(ys)
									} else {
										if Nil_(ys) {
											return false
										} else {
											if X_EQ_.Arity2IIB(First.X_invoke_Arity1(xs), First.X_invoke_Arity1(ys)) {
												xs, ys = Next.Arity1IQ(xs), Next.Arity1IQ(ys)
												continue
											} else {
												return false

											}
										}
									}
								}
							}()
						}
					}()
				} else {
					return nil
				}
			}())
		})
	}(&AFn{})

	Hash_coll = func(hash_coll *AFn) *AFn {
		return Fn(hash_coll, 1, func(coll interface{}) interface{} {
			if Truth_(Seq.Arity1IQ(coll)) {
				{
					var res interface{} = Hash.X_invoke_Arity1(First.X_invoke_Arity1(coll))
					var s interface{} = Next.Arity1IQ(coll)
					_, _ = res, s
					for {
						if Nil_(s) {
							return res
						} else {
							res, s = Hash_combine.X_invoke_Arity2(res, Hash.X_invoke_Arity1(First.X_invoke_Arity1(s))).(float64), Next.Arity1IQ(s)
							continue
						}
					}
				}
			} else {
				return float64(0)
			}
		})
	}(&AFn{})

	Hash_imap = func(hash_imap *AFn) *AFn {
		return Fn(hash_imap, 1, func(m interface{}) interface{} {
			{
				var h = float64(0)
				var s interface{} = Seq.Arity1IQ(m)
				_, _ = h, s
				for {
					if Truth_(s) {
						{
							var e = First.X_invoke_Arity1(s)
							_ = e
							h, s = math.Mod((h+float64((Int32_(Hash.X_invoke_Arity1(Key.X_invoke_Arity1(e)).(float64))^Int32_(Hash.X_invoke_Arity1(Val.X_invoke_Arity1(e)).(float64))))), float64(4503599627370496)), Next.Arity1IQ(s)
							continue
						}
					} else {
						return h
					}
				}
			}
		})
	}(&AFn{})

	Hash_iset = func(hash_iset *AFn) *AFn {
		return Fn(hash_iset, 1, func(s interface{}) interface{} {
			{
				var h = float64(0)
				var s___1 interface{} = Seq.Arity1IQ(s)
				_, _ = h, s___1
				for {
					if Truth_(s___1) {
						{
							var e = First.X_invoke_Arity1(s___1)
							_ = e
							h, s___1 = math.Mod((h+Hash.X_invoke_Arity1(e).(float64)), float64(4503599627370496)), Next.Arity1IQ(s___1)
							continue
						}
					} else {
						return h
					}
				}
			}
		})
	}(&AFn{})

	Extend_object_BANG_ = func(extend_object_BANG_ *AFn) *AFn {
		return Fn(extend_object_BANG_, 2, func(obj interface{}, fn_map interface{}) interface{} {
			{
				var seq__4629_4635 interface{} = Seq.Arity1IQ(fn_map)
				var chunk__4630_4636 interface{} = nil
				var count__4631_4637 = float64(0)
				var i__4632_4638 = float64(0)
				_, _, _, _ = seq__4629_4635, chunk__4630_4636, count__4631_4637, i__4632_4638
				for {
					if i__4632_4638 < count__4631_4637 {
						{
							var vec__4633_4639 = chunk__4630_4636.(CljsCoreIIndexed).X_nth_Arity2(i__4632_4638)
							var key_name_4640 = Nth.X_invoke_Arity3(vec__4633_4639, float64(0), nil)
							var f_4641 = Nth.X_invoke_Arity3(vec__4633_4639, float64(1), nil)
							_, _, _ = vec__4633_4639, key_name_4640, f_4641
							{
								var str_name_4642 = Name.X_invoke_Arity1(key_name_4640)
								_ = str_name_4642
								obj.([]interface{})[int(str_name_4642.(float64))] = f_4641
							}
							seq__4629_4635, chunk__4630_4636, count__4631_4637, i__4632_4638 = seq__4629_4635, chunk__4630_4636, count__4631_4637, (i__4632_4638 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___4643 = Seq.Arity1IQ(seq__4629_4635)
							_ = temp__4222__auto___4643
							if Truth_(temp__4222__auto___4643) {
								{
									var seq__4629_4644___1 = temp__4222__auto___4643
									_ = seq__4629_4644___1
									if Chunked_seq_QMARK_.Arity1IB(seq__4629_4644___1) {
										{
											var c__954__auto___4645 = Chunk_first.X_invoke_Arity1(seq__4629_4644___1)
											_ = c__954__auto___4645
											seq__4629_4635, chunk__4630_4636, count__4631_4637, i__4632_4638 = Chunk_rest.X_invoke_Arity1(seq__4629_4644___1), c__954__auto___4645, Count.X_invoke_Arity1(c__954__auto___4645).(float64), float64(0)
											continue
										}
									} else {
										{
											var vec__4634_4646 = First.X_invoke_Arity1(seq__4629_4644___1)
											var key_name_4647 = Nth.X_invoke_Arity3(vec__4634_4646, float64(0), nil)
											var f_4648 = Nth.X_invoke_Arity3(vec__4634_4646, float64(1), nil)
											_, _, _ = vec__4634_4646, key_name_4647, f_4648
											{
												var str_name_4649 = Name.X_invoke_Arity1(key_name_4647)
												_ = str_name_4649
												obj.([]interface{})[int(str_name_4649.(float64))] = f_4648
											}
											seq__4629_4635, chunk__4630_4636, count__4631_4637, i__4632_4638 = Next.Arity1IQ(seq__4629_4644___1), nil, float64(0), float64(0)
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
			return obj
		})
	}(&AFn{})

	X__GT_List = func(__GT_List *AFn) *AFn {
		return Fn(__GT_List, 5, func(meta interface{}, first interface{}, rest interface{}, count interface{}, __hash interface{}) interface{} {
			return (&CljsCoreList{meta, first, rest, count, __hash})
		})
	}(&AFn{})

	X__GT_EmptyList = func(__GT_EmptyList *AFn) *AFn {
		return Fn(__GT_EmptyList, 1, func(meta interface{}) interface{} {
			return (&CljsCoreEmptyList{meta})
		})
	}(&AFn{})

	Reversible_QMARK_ = func(reversible_QMARK_ *AFn) *AFn {
		return Fn(reversible_QMARK_, 1, func(coll interface{}) bool {
			return Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreIReversible)(nil)).Elem())
		})
	}(&AFn{})

	Rseq = func(rseq *AFn) *AFn {
		return Fn(rseq, 1, func(coll interface{}) CljsCoreISeq {
			return Seq_(coll.(CljsCoreIReversible).X_rseq_Arity1())
		})
	}(&AFn{})

	Reverse = func(reverse *AFn) *AFn {
		return Fn(reverse, 1, func(coll interface{}) interface{} {
			if Reversible_QMARK_.Arity1IB(coll) {
				return Rseq.Arity1IQ(coll)
			} else {
				return Reduce.X_invoke_Arity3(Conj, CljsCoreIEmptyList(CljsCoreList_EMPTY), coll)
			}
		})
	}(&AFn{})

	List = func(list *AFn) *AFn {
		return Fn(list, 0, func(xs__ ...interface{}) interface{} {
			var xs = Seq.Arity1IQ(xs__[0])
			_ = xs
			{
				var arr = func() interface{} {
					if (Value_(xs).Type().AssignableTo(reflect.TypeOf((**CljsCoreIndexedSeq)(nil)).Elem())) && (Native_get_instance_field.X_invoke_Arity2(xs, "I").(float64) == float64(0)) {
						return Native_get_instance_field.X_invoke_Arity2(xs, "Arr")
					} else {
						return func() []interface{} {
							var arr = []interface{}{}
							_ = arr
							{
								var xs___1 interface{} = xs
								_ = xs___1
								for {
									if !(Nil_(xs___1)) {
										js.JSArray_(&arr).Push(xs___1.(CljsCoreISeq).X_first_Arity1())
										xs___1 = xs___1.(CljsCoreINext).X_next_Arity1()
										continue
									} else {
										return arr
									}
								}
							}
						}()
					}
				}()
				_ = arr
				{
					var i = Alength_(arr)
					var r interface{} = CljsCoreIEmptyList(CljsCoreList_EMPTY)
					_, _ = i, r
					for {
						if i > float64(0) {
							i, r = (i - float64(1)), r.(CljsCoreICollection).X_conj_Arity2(Aget_(arr, (i-float64(1))))
							continue
						} else {
							return r
						}
					}
				}
			}
		})
	}(&AFn{})

	X__GT_Cons = func(__GT_Cons *AFn) *AFn {
		return Fn(__GT_Cons, 4, func(meta interface{}, first interface{}, rest interface{}, __hash interface{}) interface{} {
			return (&CljsCoreCons{meta, first, rest, __hash})
		})
	}(&AFn{})

	Cons = func(cons *AFn) *AFn {
		return Fn(cons, 2, func(x interface{}, coll interface{}) interface{} {
			if (Nil_(coll)) || (Value_(coll).Type().Implements(reflect.TypeOf((*CljsCoreISeq)(nil)).Elem())) {
				return (&CljsCoreCons{nil, x, coll, nil})
			} else {
				return (&CljsCoreCons{nil, x, Seq.Arity1IQ(coll), nil})
			}
		})
	}(&AFn{})

	List_QMARK_ = func(list_QMARK_ *AFn) *AFn {
		return Fn(list_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreIList)(nil)).Elem())
		})
	}(&AFn{})

	Hash_keyword = func(hash_keyword *AFn) *AFn {
		return Fn(hash_keyword, 1, func(k interface{}) interface{} {
			return float64(Int32_((Hash_symbol.X_invoke_Arity1(k).(float64) + float64(2654435769))))
		})
	}(&AFn{})

	X__GT_Keyword = func(__GT_Keyword *AFn) *AFn {
		return Fn(__GT_Keyword, 4, func(ns interface{}, name interface{}, fqn interface{}, _hash interface{}) interface{} {
			return (&CljsCoreKeyword{ns, name, fqn, _hash})
		})
	}(&AFn{})

	Keyword_QMARK_ = func(keyword_QMARK_ *AFn) *AFn {
		return Fn(keyword_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().AssignableTo(reflect.TypeOf((**CljsCoreKeyword)(nil)).Elem())
		})
	}(&AFn{})

	Keyword_identical_QMARK_ = func(keyword_identical_QMARK_ *AFn) *AFn {
		return Fn(keyword_identical_QMARK_, 2, func(x interface{}, y interface{}) bool {
			if reflect.DeepEqual(x, y) {
				return true
			} else {
				if (Value_(x).Type().AssignableTo(reflect.TypeOf((**CljsCoreKeyword)(nil)).Elem())) && (Value_(y).Type().AssignableTo(reflect.TypeOf((**CljsCoreKeyword)(nil)).Elem())) {
					return reflect.DeepEqual(Native_get_instance_field.X_invoke_Arity2(x, "Fqn"), Native_get_instance_field.X_invoke_Arity2(y, "Fqn"))
				} else {
					return false
				}
			}
		})
	}(&AFn{})

	Namespace = func(namespace *AFn) *AFn {
		return Fn(namespace, 1, func(x interface{}) interface{} {
			if Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreINamed)(nil)).Elem()) {
				return x.(CljsCoreINamed).X_namespace_Arity1()
			} else {
				panic((&js.Error{("Doesn't support namespace: " + Str.X_invoke_Arity1(x).(string))}))
			}
		})
	}(&AFn{})

	Keyword = func(keyword *AFn) *AFn {
		return Fn(keyword, 2, func(name interface{}) interface{} {
			if Value_(name).Type().AssignableTo(reflect.TypeOf((**CljsCoreKeyword)(nil)).Elem()) {
				return name
			} else {
				if Value_(name).Type().AssignableTo(reflect.TypeOf((**CljsCoreSymbol)(nil)).Elem()) {
					return (&CljsCoreKeyword{Namespace.X_invoke_Arity1(name), Name.X_invoke_Arity1(name), Native_get_instance_field.X_invoke_Arity2(name, "Str"), nil})
				} else {
					if Value_(name).Kind() == reflect.String {
						{
							var parts = Native_invoke_instance_method.X_invoke_Arity3(name, "Split", []interface{}{"/"})
							_ = parts
							if Alength_(parts) == float64(2) {
								return (&CljsCoreKeyword{Aget_(parts, float64(0)), Aget_(parts, float64(1)), name, nil})
							} else {
								return (&CljsCoreKeyword{nil, Aget_(parts, float64(0)), name, nil})
							}
						}
					} else {
						return nil
					}
				}
			}
		}, func(ns interface{}, name interface{}) interface{} {
			return (&CljsCoreKeyword{ns, name, (`` + Str.X_invoke_Arity1(func() interface{} {
				if Truth_(ns) {
					return (`` + Str.X_invoke_Arity1(ns).(string) + "/")
				} else {
					return nil
				}
			}()).(string) + Str.X_invoke_Arity1(name).(string)), nil})
		})
	}(&AFn{})

	X__GT_LazySeq = func(__GT_LazySeq *AFn) *AFn {
		return Fn(__GT_LazySeq, 4, func(meta interface{}, fn interface{}, s interface{}, __hash interface{}) interface{} {
			return (&CljsCoreLazySeq{meta, fn, s, __hash})
		})
	}(&AFn{})

	X__GT_ChunkBuffer = func(__GT_ChunkBuffer *AFn) *AFn {
		return Fn(__GT_ChunkBuffer, 2, func(buf interface{}, end interface{}) interface{} {
			return (&CljsCoreChunkBuffer{buf, end})
		})
	}(&AFn{})

	Chunk_buffer = func(chunk_buffer *AFn) *AFn {
		return Fn(chunk_buffer, 1, func(capacity interface{}) interface{} {
			return (&CljsCoreChunkBuffer{make([]interface{}, int(capacity.(float64))), float64(0)})
		})
	}(&AFn{})

	X__GT_ArrayChunk = func(__GT_ArrayChunk *AFn) *AFn {
		return Fn(__GT_ArrayChunk, 3, func(arr interface{}, off interface{}, end interface{}) interface{} {
			return (&CljsCoreArrayChunk{arr, off, end})
		})
	}(&AFn{})

	Array_chunk = func(array_chunk *AFn) *AFn {
		return Fn(array_chunk, 3, func(arr interface{}) interface{} {
			return (&CljsCoreArrayChunk{arr, float64(0), Alength_(arr)})
		}, func(arr interface{}, off interface{}) interface{} {
			return (&CljsCoreArrayChunk{arr, off, Alength_(arr)})
		}, func(arr interface{}, off interface{}, end interface{}) interface{} {
			return (&CljsCoreArrayChunk{arr, off, end})
		})
	}(&AFn{})

	X__GT_ChunkedCons = func(__GT_ChunkedCons *AFn) *AFn {
		return Fn(__GT_ChunkedCons, 4, func(chunk interface{}, more interface{}, meta interface{}, __hash interface{}) interface{} {
			return (&CljsCoreChunkedCons{chunk, more, meta, __hash})
		})
	}(&AFn{})

	Chunk_cons = func(chunk_cons *AFn) *AFn {
		return Fn(chunk_cons, 2, func(chunk interface{}, rest interface{}) interface{} {
			if chunk.(CljsCoreICounted).X_count_Arity1() == float64(0) {
				return rest
			} else {
				return (&CljsCoreChunkedCons{chunk, rest, nil, nil})
			}
		})
	}(&AFn{})

	Chunk_append = func(chunk_append *AFn) *AFn {
		return Fn(chunk_append, 2, func(b interface{}, x interface{}) interface{} {
			return Native_invoke_instance_method.X_invoke_Arity3(b, "Add", []interface{}{x})
		})
	}(&AFn{})

	Chunk = func(chunk *AFn) *AFn {
		return Fn(chunk, 1, func(b interface{}) interface{} {
			return Native_invoke_instance_method.X_invoke_Arity3(b, "Chunk", []interface{}{})
		})
	}(&AFn{})

	Chunk_first = func(chunk_first *AFn) *AFn {
		return Fn(chunk_first, 1, func(s interface{}) interface{} {
			return s.(CljsCoreIChunkedSeq).X_chunked_first_Arity1()
		})
	}(&AFn{})

	Chunk_rest = func(chunk_rest *AFn) *AFn {
		return Fn(chunk_rest, 1, func(s interface{}) interface{} {
			return s.(CljsCoreIChunkedSeq).X_chunked_rest_Arity1()
		})
	}(&AFn{})

	Chunk_next = func(chunk_next *AFn) *AFn {
		return Fn(chunk_next, 1, func(s interface{}) interface{} {
			if Value_(s).Type().Implements(reflect.TypeOf((*CljsCoreIChunkedNext)(nil)).Elem()) {
				return s.(CljsCoreIChunkedNext).X_chunked_next_Arity1()
			} else {
				return Seq.Arity1IQ(s.(CljsCoreIChunkedSeq).X_chunked_rest_Arity1())
			}
		})
	}(&AFn{})

	To_array = func(to_array *AFn) *AFn {
		return Fn(to_array, 1, func(s interface{}) interface{} {
			{
				var ary = []interface{}{}
				_ = ary
				{
					var s___1 interface{} = s
					_ = s___1
					for {
						if Truth_(Seq.Arity1IQ(s___1)) {
							js.JSArray_(&ary).Push(First.X_invoke_Arity1(s___1))
							s___1 = Next.Arity1IQ(s___1)
							continue
						} else {
							return ary
						}
					}
				}
			}
		})
	}(&AFn{})

	To_array_2d = func(to_array_2d *AFn) *AFn {
		return Fn(to_array_2d, 1, func(coll interface{}) interface{} {
			{
				var ret = make([]interface{}, int(Count.X_invoke_Arity1(coll).(float64)))
				_ = ret
				{
					var i_4650 = float64(0)
					var xs_4651 interface{} = Seq.Arity1IQ(coll)
					_, _ = i_4650, xs_4651
					for {
						if Truth_(xs_4651) {
							ret[int(i_4650)] = To_array.X_invoke_Arity1(First.X_invoke_Arity1(xs_4651)).([]interface{})
							i_4650, xs_4651 = (i_4650 + float64(1)), Next.Arity1IQ(xs_4651)
							continue
						} else {
						}
						break
					}
				}
				return ret
			}
		})
	}(&AFn{})

	Int_array = func(int_array *AFn) *AFn {
		return Fn(int_array, 2, func(size_or_seq interface{}) interface{} {
			if Value_(size_or_seq).Kind() == reflect.Float64 {
				return int_array.X_invoke_Arity2(size_or_seq, nil).([]interface{})
			} else {
				return Into_array.Arity1IA(size_or_seq)
			}
		}, func(size interface{}, init_val_or_seq interface{}) interface{} {
			{
				var a = make([]interface{}, int(size.(float64)))
				_ = a
				if Seq_QMARK_.Arity1IB(init_val_or_seq) {
					{
						var s = Seq.Arity1IQ(init_val_or_seq)
						_ = s
						{
							var i = float64(0)
							var s___1 interface{} = s
							_, _ = i, s___1
							for {
								if Truth_(func() interface{} {
									var and__159__auto__ = s___1
									_ = and__159__auto__
									if Truth_(and__159__auto__) {
										return (i < size.(float64))
									} else {
										return and__159__auto__
									}
								}()) {
									a[int(i)] = First.X_invoke_Arity1(s___1)
									i, s___1 = (i + float64(1)), Next.Arity1IQ(s___1)
									continue
								} else {
									return a
								}
							}
						}
					}
				} else {
					{
						var n__1054__auto___4654 = size
						_ = n__1054__auto___4654
						{
							var i_4655 = float64(0)
							_ = i_4655
							for {
								if i_4655 < n__1054__auto___4654.(float64) {
									a[int(i_4655)] = init_val_or_seq
									i_4655 = (i_4655 + float64(1))
									continue
								} else {
								}
								break
							}
						}
					}
					return a
				}
			}
		})
	}(&AFn{})

	Long_array = func(long_array *AFn) *AFn {
		return Fn(long_array, 2, func(size_or_seq interface{}) interface{} {
			if Value_(size_or_seq).Kind() == reflect.Float64 {
				return long_array.X_invoke_Arity2(size_or_seq, nil).([]interface{})
			} else {
				return Into_array.Arity1IA(size_or_seq)
			}
		}, func(size interface{}, init_val_or_seq interface{}) interface{} {
			{
				var a = make([]interface{}, int(size.(float64)))
				_ = a
				if Seq_QMARK_.Arity1IB(init_val_or_seq) {
					{
						var s = Seq.Arity1IQ(init_val_or_seq)
						_ = s
						{
							var i = float64(0)
							var s___1 interface{} = s
							_, _ = i, s___1
							for {
								if Truth_(func() interface{} {
									var and__159__auto__ = s___1
									_ = and__159__auto__
									if Truth_(and__159__auto__) {
										return (i < size.(float64))
									} else {
										return and__159__auto__
									}
								}()) {
									a[int(i)] = First.X_invoke_Arity1(s___1)
									i, s___1 = (i + float64(1)), Next.Arity1IQ(s___1)
									continue
								} else {
									return a
								}
							}
						}
					}
				} else {
					{
						var n__1054__auto___4658 = size
						_ = n__1054__auto___4658
						{
							var i_4659 = float64(0)
							_ = i_4659
							for {
								if i_4659 < n__1054__auto___4658.(float64) {
									a[int(i_4659)] = init_val_or_seq
									i_4659 = (i_4659 + float64(1))
									continue
								} else {
								}
								break
							}
						}
					}
					return a
				}
			}
		})
	}(&AFn{})

	Double_array = func(double_array *AFn) *AFn {
		return Fn(double_array, 2, func(size_or_seq interface{}) interface{} {
			if Value_(size_or_seq).Kind() == reflect.Float64 {
				return double_array.X_invoke_Arity2(size_or_seq, nil).([]interface{})
			} else {
				return Into_array.Arity1IA(size_or_seq)
			}
		}, func(size interface{}, init_val_or_seq interface{}) interface{} {
			{
				var a = make([]interface{}, int(size.(float64)))
				_ = a
				if Seq_QMARK_.Arity1IB(init_val_or_seq) {
					{
						var s = Seq.Arity1IQ(init_val_or_seq)
						_ = s
						{
							var i = float64(0)
							var s___1 interface{} = s
							_, _ = i, s___1
							for {
								if Truth_(func() interface{} {
									var and__159__auto__ = s___1
									_ = and__159__auto__
									if Truth_(and__159__auto__) {
										return (i < size.(float64))
									} else {
										return and__159__auto__
									}
								}()) {
									a[int(i)] = First.X_invoke_Arity1(s___1)
									i, s___1 = (i + float64(1)), Next.Arity1IQ(s___1)
									continue
								} else {
									return a
								}
							}
						}
					}
				} else {
					{
						var n__1054__auto___4662 = size
						_ = n__1054__auto___4662
						{
							var i_4663 = float64(0)
							_ = i_4663
							for {
								if i_4663 < n__1054__auto___4662.(float64) {
									a[int(i_4663)] = init_val_or_seq
									i_4663 = (i_4663 + float64(1))
									continue
								} else {
								}
								break
							}
						}
					}
					return a
				}
			}
		})
	}(&AFn{})

	Object_array = func(object_array *AFn) *AFn {
		return Fn(object_array, 2, func(size_or_seq interface{}) interface{} {
			if Value_(size_or_seq).Kind() == reflect.Float64 {
				return object_array.X_invoke_Arity2(size_or_seq, nil).([]interface{})
			} else {
				return Into_array.Arity1IA(size_or_seq)
			}
		}, func(size interface{}, init_val_or_seq interface{}) interface{} {
			{
				var a = make([]interface{}, int(size.(float64)))
				_ = a
				if Seq_QMARK_.Arity1IB(init_val_or_seq) {
					{
						var s = Seq.Arity1IQ(init_val_or_seq)
						_ = s
						{
							var i = float64(0)
							var s___1 interface{} = s
							_, _ = i, s___1
							for {
								if Truth_(func() interface{} {
									var and__159__auto__ = s___1
									_ = and__159__auto__
									if Truth_(and__159__auto__) {
										return (i < size.(float64))
									} else {
										return and__159__auto__
									}
								}()) {
									a[int(i)] = First.X_invoke_Arity1(s___1)
									i, s___1 = (i + float64(1)), Next.Arity1IQ(s___1)
									continue
								} else {
									return a
								}
							}
						}
					}
				} else {
					{
						var n__1054__auto___4666 = size
						_ = n__1054__auto___4666
						{
							var i_4667 = float64(0)
							_ = i_4667
							for {
								if i_4667 < n__1054__auto___4666.(float64) {
									a[int(i_4667)] = init_val_or_seq
									i_4667 = (i_4667 + float64(1))
									continue
								} else {
								}
								break
							}
						}
					}
					return a
				}
			}
		})
	}(&AFn{})

	Bounded_count = func(bounded_count *AFn) *AFn {
		return Fn(bounded_count, 2, func(s interface{}, n interface{}) interface{} {
			if Counted_QMARK_.Arity1IB(s) {
				return Count.X_invoke_Arity1(s).(float64)
			} else {
				{
					var s___1 interface{} = s
					var i interface{} = n
					var sum = float64(0)
					_, _, _ = s___1, i, sum
					for {
						if Truth_(func() interface{} {
							var and__159__auto__ = (i.(float64) > float64(0))
							_ = and__159__auto__
							if Truth_(and__159__auto__) {
								return Seq.Arity1IQ(s___1)
							} else {
								return and__159__auto__
							}
						}()) {
							s___1, i, sum = Next.Arity1IQ(s___1), (i.(float64) - float64(1)), (sum + float64(1))
							continue
						} else {
							return sum
						}
					}
				}
			}
		})
	}(&AFn{})

	Spread = func(spread *AFn) *AFn {
		return Fn(spread, 1, func(arglist interface{}) interface{} {
			if Nil_(arglist) {
				return nil
			} else {
				if Nil_(Next.Arity1IQ(arglist)) {
					return Seq.Arity1IQ(First.X_invoke_Arity1(arglist))
				} else {
					return Cons.X_invoke_Arity2(First.X_invoke_Arity1(arglist), spread.X_invoke_Arity1(Next.Arity1IQ(arglist))).(*CljsCoreCons)

				}
			}
		})
	}(&AFn{})

	Concat = func(concat *AFn) *AFn {
		return Fn(concat, 2, func() interface{} {
			return (&CljsCoreLazySeq{nil, func(G__4687 *AFn) *AFn {
				return Fn(G__4687, 0, func() interface{} {
					return nil
				})
			}(&AFn{}), nil, nil})
		}, func(x interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__4688 *AFn) *AFn {
				return Fn(G__4688, 0, func() interface{} {
					return x
				})
			}(&AFn{}), nil, nil})
		}, func(x interface{}, y interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__4689 *AFn) *AFn {
				return Fn(G__4689, 0, func() interface{} {
					{
						var s = Seq.Arity1IQ(x)
						_ = s
						if Truth_(s) {
							if Chunked_seq_QMARK_.Arity1IB(s) {
								return Chunk_cons.X_invoke_Arity2(Chunk_first.X_invoke_Arity1(s), concat.X_invoke_Arity2(Chunk_rest.X_invoke_Arity1(s), y).(*CljsCoreLazySeq))
							} else {
								return Cons.X_invoke_Arity2(First.X_invoke_Arity1(s), concat.X_invoke_Arity2(Rest.Arity1IQ(s), y).(*CljsCoreLazySeq)).(*CljsCoreCons)
							}
						} else {
							return y
						}
					}
				})
			}(&AFn{}), nil, nil})
		}, func(x_y_zs__ ...interface{}) interface{} {
			var x = x_y_zs__[0]
			var y = x_y_zs__[1]
			var zs = Seq.Arity1IQ(x_y_zs__[2])
			_, _, _ = x, y, zs
			{
				var cat = func(cat *AFn) *AFn {
					return Fn(cat, 2, func(xys interface{}, zs___1 interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__4690 *AFn) *AFn {
							return Fn(G__4690, 0, func() interface{} {
								{
									var xys___1 = Seq.Arity1IQ(xys)
									_ = xys___1
									if Truth_(xys___1) {
										if Chunked_seq_QMARK_.Arity1IB(xys___1) {
											return Chunk_cons.X_invoke_Arity2(Chunk_first.X_invoke_Arity1(xys___1), cat.X_invoke_Arity2(Chunk_rest.X_invoke_Arity1(xys___1), zs___1).(*CljsCoreLazySeq))
										} else {
											return Cons.X_invoke_Arity2(First.X_invoke_Arity1(xys___1), cat.X_invoke_Arity2(Rest.Arity1IQ(xys___1), zs___1).(*CljsCoreLazySeq)).(*CljsCoreCons)
										}
									} else {
										if Truth_(zs___1) {
											return cat.X_invoke_Arity2(First.X_invoke_Arity1(zs___1), Next.Arity1IQ(zs___1)).(*CljsCoreLazySeq)
										} else {
											return nil
										}
									}
								}
							})
						}(&AFn{}), nil, nil})
					})
				}(&AFn{})
				_ = cat
				return cat.X_invoke_Arity2(concat.X_invoke_Arity2(x, y).(*CljsCoreLazySeq), zs).(*CljsCoreLazySeq)
			}
		})
	}(&AFn{})

	List_STAR_ = func(list_STAR_ *AFn) *AFn {
		return Fn(list_STAR_, 4, func(args interface{}) interface{} {
			return Seq.Arity1IQ(args)
		}, func(a interface{}, args interface{}) interface{} {
			return Cons.X_invoke_Arity2(a, args).(*CljsCoreCons)
		}, func(a interface{}, b interface{}, args interface{}) interface{} {
			return Cons.X_invoke_Arity2(a, Cons.X_invoke_Arity2(b, args).(*CljsCoreCons)).(*CljsCoreCons)
		}, func(a interface{}, b interface{}, c interface{}, args interface{}) interface{} {
			return Cons.X_invoke_Arity2(a, Cons.X_invoke_Arity2(b, Cons.X_invoke_Arity2(c, args).(*CljsCoreCons)).(*CljsCoreCons)).(*CljsCoreCons)
		}, func(a_b_c_d_more__ ...interface{}) interface{} {
			var a = a_b_c_d_more__[0]
			var b = a_b_c_d_more__[1]
			var c = a_b_c_d_more__[2]
			var d = a_b_c_d_more__[3]
			var more = Seq.Arity1IQ(a_b_c_d_more__[4])
			_, _, _, _, _ = a, b, c, d, more
			return Cons.X_invoke_Arity2(a, Cons.X_invoke_Arity2(b, Cons.X_invoke_Arity2(c, Cons.X_invoke_Arity2(d, Spread.X_invoke_Arity1(more)).(*CljsCoreCons)).(*CljsCoreCons)).(*CljsCoreCons)).(*CljsCoreCons)
		})
	}(&AFn{})

	Transient = func(transient *AFn) *AFn {
		return Fn(transient, 1, func(coll interface{}) interface{} {
			return coll.(CljsCoreIEditableCollection).X_as_transient_Arity1()
		})
	}(&AFn{})

	Persistent_BANG_ = func(persistent_BANG_ *AFn) *AFn {
		return Fn(persistent_BANG_, 1, func(tcoll interface{}) interface{} {
			return tcoll.(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
		})
	}(&AFn{})

	Conj_BANG_ = func(conj_BANG_ *AFn) *AFn {
		return Fn(conj_BANG_, 2, func() interface{} {
			return Transient.X_invoke_Arity1(CljsCorePersistentVector_EMPTY)
		}, func(coll interface{}) interface{} {
			return coll
		}, func(tcoll interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientCollection).X_conj_BANG__Arity2(val)
		}, func(tcoll_val_vals__ ...interface{}) interface{} {
			var tcoll = tcoll_val_vals__[0]
			var val = tcoll_val_vals__[1]
			var vals = Seq.Arity1IQ(tcoll_val_vals__[2])
			_, _, _ = tcoll, val, vals
			for {
				{
					var ntcoll = tcoll.(CljsCoreITransientCollection).X_conj_BANG__Arity2(val)
					_ = ntcoll
					if Truth_(vals) {
						tcoll, val, vals = ntcoll, First.X_invoke_Arity1(vals), Next.Arity1IQ(vals)
						continue
					} else {
						return ntcoll
					}
				}
			}
		})
	}(&AFn{})

	Assoc_BANG_ = func(assoc_BANG_ *AFn) *AFn {
		return Fn(assoc_BANG_, 3, func(tcoll interface{}, key interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3(key, val)
		}, func(tcoll_key_val_kvs__ ...interface{}) interface{} {
			var tcoll = tcoll_key_val_kvs__[0]
			var key = tcoll_key_val_kvs__[1]
			var val = tcoll_key_val_kvs__[2]
			var kvs = Seq.Arity1IQ(tcoll_key_val_kvs__[3])
			_, _, _, _ = tcoll, key, val, kvs
			for {
				{
					var ntcoll = tcoll.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3(key, val)
					_ = ntcoll
					if Truth_(kvs) {
						tcoll, key, val, kvs = ntcoll, First.X_invoke_Arity1(kvs), Second.X_invoke_Arity1(kvs), Seq_(Nnext.X_invoke_Arity1(kvs))
						continue
					} else {
						return ntcoll
					}
				}
			}
		})
	}(&AFn{})

	Dissoc_BANG_ = func(dissoc_BANG_ *AFn) *AFn {
		return Fn(dissoc_BANG_, 2, func(tcoll interface{}, key interface{}) interface{} {
			return tcoll.(CljsCoreITransientMap).X_dissoc_BANG__Arity2(key)
		}, func(tcoll_key_ks__ ...interface{}) interface{} {
			var tcoll = tcoll_key_ks__[0]
			var key = tcoll_key_ks__[1]
			var ks = Seq.Arity1IQ(tcoll_key_ks__[2])
			_, _, _ = tcoll, key, ks
			for {
				{
					var ntcoll = tcoll.(CljsCoreITransientMap).X_dissoc_BANG__Arity2(key)
					_ = ntcoll
					if Truth_(ks) {
						tcoll, key, ks = ntcoll, First.X_invoke_Arity1(ks), Next.Arity1IQ(ks)
						continue
					} else {
						return ntcoll
					}
				}
			}
		})
	}(&AFn{})

	Pop_BANG_ = func(pop_BANG_ *AFn) *AFn {
		return Fn(pop_BANG_, 1, func(tcoll interface{}) interface{} {
			return tcoll.(CljsCoreITransientVector).X_pop_BANG__Arity1()
		})
	}(&AFn{})

	Disj_BANG_ = func(disj_BANG_ *AFn) *AFn {
		return Fn(disj_BANG_, 2, func(tcoll interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientSet).X_disjoin_BANG__Arity2(val)
		}, func(tcoll_val_vals__ ...interface{}) interface{} {
			var tcoll = tcoll_val_vals__[0]
			var val = tcoll_val_vals__[1]
			var vals = Seq.Arity1IQ(tcoll_val_vals__[2])
			_, _, _ = tcoll, val, vals
			for {
				{
					var ntcoll = tcoll.(CljsCoreITransientSet).X_disjoin_BANG__Arity2(val)
					_ = ntcoll
					if Truth_(vals) {
						tcoll, val, vals = ntcoll, First.X_invoke_Arity1(vals), Next.Arity1IQ(vals)
						continue
					} else {
						return ntcoll
					}
				}
			}
		})
	}(&AFn{})

	Vary_meta = func(vary_meta *AFn) *AFn {
		return Fn(vary_meta, 6, func(obj interface{}, f interface{}) interface{} {
			return With_meta.X_invoke_Arity2(obj, func() interface{} {
				var G__4736 = Meta.X_invoke_Arity1(obj)
				_ = G__4736
				return f.(CljsCoreIFn).X_invoke_Arity1(G__4736)
			}())
		}, func(obj interface{}, f interface{}, a interface{}) interface{} {
			return With_meta.X_invoke_Arity2(obj, func() interface{} {
				var G__4737 = Meta.X_invoke_Arity1(obj)
				var G__4738 = a
				_, _ = G__4737, G__4738
				return f.(CljsCoreIFn).X_invoke_Arity2(G__4737, G__4738)
			}())
		}, func(obj interface{}, f interface{}, a interface{}, b interface{}) interface{} {
			return With_meta.X_invoke_Arity2(obj, func() interface{} {
				var G__4739 = Meta.X_invoke_Arity1(obj)
				var G__4740 = a
				var G__4741 = b
				_, _, _ = G__4739, G__4740, G__4741
				return f.(CljsCoreIFn).X_invoke_Arity3(G__4739, G__4740, G__4741)
			}())
		}, func(obj interface{}, f interface{}, a interface{}, b interface{}, c interface{}) interface{} {
			return With_meta.X_invoke_Arity2(obj, func() interface{} {
				var G__4742 = Meta.X_invoke_Arity1(obj)
				var G__4743 = a
				var G__4744 = b
				var G__4745 = c
				_, _, _, _ = G__4742, G__4743, G__4744, G__4745
				return f.(CljsCoreIFn).X_invoke_Arity4(G__4742, G__4743, G__4744, G__4745)
			}())
		}, func(obj interface{}, f interface{}, a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
			return With_meta.X_invoke_Arity2(obj, func() interface{} {
				var G__4746 = Meta.X_invoke_Arity1(obj)
				var G__4747 = a
				var G__4748 = b
				var G__4749 = c
				var G__4750 = d
				_, _, _, _, _ = G__4746, G__4747, G__4748, G__4749, G__4750
				return f.(CljsCoreIFn).X_invoke_Arity5(G__4746, G__4747, G__4748, G__4749, G__4750)
			}())
		}, func(obj_f_a_b_c_d_args__ ...interface{}) interface{} {
			var obj = obj_f_a_b_c_d_args__[0]
			var f = obj_f_a_b_c_d_args__[1]
			var a = obj_f_a_b_c_d_args__[2]
			var b = obj_f_a_b_c_d_args__[3]
			var c = obj_f_a_b_c_d_args__[4]
			var d = obj_f_a_b_c_d_args__[5]
			var args = Seq.Arity1IQ(obj_f_a_b_c_d_args__[6])
			_, _, _, _, _, _, _ = obj, f, a, b, c, d, args
			return With_meta.X_invoke_Arity2(obj, Apply.X_invoke_ArityVariadic(f, Meta.X_invoke_Arity1(obj), a, b, c, Array_seq.X_invoke_Arity1([]interface{}{d, args})))
		})
	}(&AFn{})

	Not_EQ_ = func(not_EQ_ *AFn) *AFn {
		return Fn(not_EQ_, 2, func(x interface{}) bool {
			return false
		}, func(x interface{}, y interface{}) bool {
			return !(X_EQ_.Arity2IIB(x, y))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Seq.Arity1IQ(x_y_more__[2])
			_, _, _ = x, y, more
			return Not.Arity1IB(Apply.X_invoke_Arity4(X_EQ_, x, y, more))
		})
	}(&AFn{})

	Not_empty = func(not_empty *AFn) *AFn {
		return Fn(not_empty, 1, func(coll interface{}) interface{} {
			if Truth_(Seq.Arity1IQ(coll)) {
				return coll
			} else {
				return nil
			}
		})
	}(&AFn{})

	Nil_iter = func(nil_iter *AFn) *AFn {
		return Fn(nil_iter, 0, func() interface{} {
			X__GT_t4754 = func(__GT_t4754 *AFn) *AFn {
				return Fn(__GT_t4754, 2, func(nil_iter___1 interface{}, meta4755 interface{}) interface{} {
					return (&CljsCoreT4754{nil_iter___1, meta4755})
				})
			}(&AFn{})

			return (&CljsCoreT4754{nil_iter, nil})
		})
	}(&AFn{})

	X__GT_StringIter = func(__GT_StringIter *AFn) *AFn {
		return Fn(__GT_StringIter, 2, func(s interface{}, i interface{}) interface{} {
			return (&CljsCoreStringIter{s, i})
		})
	}(&AFn{})

	String_iter = func(string_iter *AFn) *AFn {
		return Fn(string_iter, 1, func(x interface{}) interface{} {
			return (&CljsCoreStringIter{x, float64(0)})
		})
	}(&AFn{})

	X__GT_ArrayIter = func(__GT_ArrayIter *AFn) *AFn {
		return Fn(__GT_ArrayIter, 2, func(arr interface{}, i interface{}) interface{} {
			return (&CljsCoreArrayIter{arr, i})
		})
	}(&AFn{})

	Array_iter = func(array_iter *AFn) *AFn {
		return Fn(array_iter, 1, func(x interface{}) interface{} {
			return (&CljsCoreArrayIter{x, float64(0)})
		})
	}(&AFn{})

	INIT = map[string]interface{}{}

	START = map[string]interface{}{}

	X__GT_SeqIter = func(__GT_SeqIter *AFn) *AFn {
		return Fn(__GT_SeqIter, 2, func(_seq interface{}, _next interface{}) interface{} {
			return (&CljsCoreSeqIter{_seq, _next})
		})
	}(&AFn{})

	Seq_iter = func(seq_iter *AFn) *AFn {
		return Fn(seq_iter, 1, func(coll interface{}) interface{} {
			return (&CljsCoreSeqIter{INIT, coll})
		})
	}(&AFn{})

	Iter = func(iter *AFn) *AFn {
		return Fn(iter, 1, func(coll interface{}) interface{} {
			if Nil_(coll) {
				return Nil_iter.X_invoke_Arity0().(*CljsCoreT4754)
			} else {
				if Value_(coll).Kind() == reflect.String {
					return String_iter.X_invoke_Arity1(coll).(*CljsCoreStringIter)
				} else {
					if Value_(coll).Kind() == reflect.Slice {
						return Array_iter.X_invoke_Arity1(coll).(*CljsCoreArrayIter)
					} else {
						if Truth_(Iterable_QMARK_.X_invoke_Arity1(coll)) {
							return coll.(CljsCoreIIterable).X_iterator_Arity1()
						} else {
							if Seqable_QMARK_.Arity1IB(coll) {
								return Seq_iter.X_invoke_Arity1(coll).(*CljsCoreSeqIter)
							} else {
								panic((&js.Error{("Cannot create iterator from " + Str.X_invoke_Arity1(coll).(string))}))

							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Lazy_transformer = func(lazy_transformer *AFn) *AFn {
		return Fn(lazy_transformer, 1, func(stepper interface{}) interface{} {
			return (&CljsCoreLazyTransformer{stepper, nil, nil, nil})
		})
	}(&AFn{})

	X__GT_Stepper = func(__GT_Stepper *AFn) *AFn {
		return Fn(__GT_Stepper, 2, func(xform interface{}, iter interface{}) interface{} {
			return (&CljsCoreStepper{xform, iter})
		})
	}(&AFn{})

	Stepper = func(stepper *AFn) *AFn {
		return Fn(stepper, 2, func(xform interface{}, iter interface{}) interface{} {
			{
				var stepfn *AFn
				stepfn = func(stepfn *AFn) *AFn {
					return Fn(stepfn, 2, func(result interface{}) interface{} {
						{
							var lt = func() interface{} {
								if Reduced_QMARK_.Arity1IB(result) {
									return Deref.X_invoke_Arity1(result)
								} else {
									return result
								}
							}()
							_ = lt
							Native_set_instance_field.X_invoke_Arity3(lt, "Stepper", nil)
							return result
						}
					}, func(result interface{}, input interface{}) interface{} {
						{
							var lt = result
							_ = lt
							Native_set_instance_field.X_invoke_Arity3(lt, "First", input)
							Native_set_instance_field.X_invoke_Arity3(lt, "Rest", Lazy_transformer.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper")).(*CljsCoreLazyTransformer))
							Native_set_instance_field.X_invoke_Arity3(lt, "Stepper", nil)
							return Native_get_instance_field.X_invoke_Arity2(lt, "Rest")
						}
					})
				}(&AFn{})
				_ = stepfn
				return (&CljsCoreStepper{func() interface{} {
					var G__4764 = stepfn
					_ = G__4764
					return xform.(CljsCoreIFn).X_invoke_Arity1(G__4764)
				}(), iter})
			}
		})
	}(&AFn{})

	X__GT_MultiStepper = func(__GT_MultiStepper *AFn) *AFn {
		return Fn(__GT_MultiStepper, 3, func(xform interface{}, iters interface{}, nexts interface{}) interface{} {
			return (&CljsCoreMultiStepper{xform, iters, nexts})
		})
	}(&AFn{})

	Multi_stepper = func(multi_stepper *AFn) *AFn {
		return Fn(multi_stepper, 3, func(xform interface{}, iters interface{}) interface{} {
			return multi_stepper.X_invoke_Arity3(xform, iters, make([]interface{}, int(Alength_(iters))))
		}, func(xform interface{}, iters interface{}, nexts interface{}) interface{} {
			{
				var stepfn *AFn
				stepfn = func(stepfn *AFn) *AFn {
					return Fn(stepfn, 2, func(result interface{}) interface{} {
						{
							var lt = func() interface{} {
								if Reduced_QMARK_.Arity1IB(result) {
									return Deref.X_invoke_Arity1(result)
								} else {
									return result
								}
							}()
							_ = lt
							Native_set_instance_field.X_invoke_Arity3(lt, "Stepper", nil)
							return lt
						}
					}, func(result interface{}, input interface{}) interface{} {
						{
							var lt = result
							_ = lt
							Native_set_instance_field.X_invoke_Arity3(lt, "First", input)
							Native_set_instance_field.X_invoke_Arity3(lt, "Rest", Lazy_transformer.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper")).(*CljsCoreLazyTransformer))
							Native_set_instance_field.X_invoke_Arity3(lt, "Stepper", nil)
							return Native_get_instance_field.X_invoke_Arity2(lt, "Rest")
						}
					})
				}(&AFn{})
				_ = stepfn
				return (&CljsCoreMultiStepper{func() interface{} {
					var G__4771 = stepfn
					_ = G__4771
					return xform.(CljsCoreIFn).X_invoke_Arity1(G__4771)
				}(), iters, nexts})
			}
		})
	}(&AFn{})

	X__GT_LazyTransformer = func(__GT_LazyTransformer *AFn) *AFn {
		return Fn(__GT_LazyTransformer, 4, func(stepper interface{}, first interface{}, rest interface{}, meta interface{}) interface{} {
			return (&CljsCoreLazyTransformer{stepper, first, rest, meta})
		})
	}(&AFn{})

	Sequence = func(sequence *AFn) *AFn {
		return Fn(sequence, 2, func(coll interface{}) interface{} {
			if Seq_QMARK_.Arity1IB(coll) {
				return coll
			} else {
				{
					var or__171__auto__ = Seq.Arity1IQ(coll)
					_ = or__171__auto__
					if Truth_(or__171__auto__) {
						return or__171__auto__
					} else {
						return CljsCoreIEmptyList(CljsCoreList_EMPTY)
					}
				}
			}
		}, func(xform interface{}, coll interface{}) interface{} {
			return CljsCoreLazyTransformer_Create.X_invoke_Arity2(xform, coll)
		}, func(xform_coll_colls__ ...interface{}) interface{} {
			var xform = xform_coll_colls__[0]
			var coll = xform_coll_colls__[1]
			var colls = Seq.Arity1IQ(xform_coll_colls__[2])
			_, _, _ = xform, coll, colls
			return CljsCoreLazyTransformer_CreateMulti.X_invoke_Arity2(xform, To_array.X_invoke_Arity1(Cons.X_invoke_Arity2(coll, colls).(*CljsCoreCons)).([]interface{}))
		})
	}(&AFn{})

	Every_QMARK_ = func(every_QMARK_ *AFn) *AFn {
		return Fn(every_QMARK_, 2, func(pred interface{}, coll interface{}) bool {
			for {
				if Nil_(Seq.Arity1IQ(coll)) {
					return true
				} else {
					if Truth_(func() interface{} {
						var G__4788 = First.X_invoke_Arity1(coll)
						_ = G__4788
						return pred.(CljsCoreIFn).X_invoke_Arity1(G__4788)
					}()) {
						pred, coll = pred, Next.Arity1IQ(coll)
						continue
					} else {
						return false

					}
				}
			}
		})
	}(&AFn{})

	Not_every_QMARK_ = func(not_every_QMARK_ *AFn) *AFn {
		return Fn(not_every_QMARK_, 2, func(pred interface{}, coll interface{}) bool {
			return !(Every_QMARK_.Arity2IIB(pred, coll))
		})
	}(&AFn{})

	Some = func(some *AFn) *AFn {
		return Fn(some, 2, func(pred interface{}, coll interface{}) interface{} {
			for {
				if Truth_(Seq.Arity1IQ(coll)) {
					{
						var or__171__auto__ = func() interface{} {
							var G__4792 = First.X_invoke_Arity1(coll)
							_ = G__4792
							return pred.(CljsCoreIFn).X_invoke_Arity1(G__4792)
						}()
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							pred, coll = pred, Next.Arity1IQ(coll)
							continue
						}
					}
				} else {
					return nil
				}
			}
		})
	}(&AFn{})

	Not_any_QMARK_ = func(not_any_QMARK_ *AFn) *AFn {
		return Fn(not_any_QMARK_, 2, func(pred interface{}, coll interface{}) bool {
			return Not.Arity1IB(Some.X_invoke_Arity2(pred, coll))
		})
	}(&AFn{})

	Even_QMARK_ = func(even_QMARK_ *AFn) *AFn {
		return Fn(even_QMARK_, 1, func(n interface{}) bool {
			if Integer_QMARK_.Arity1IB(n) {
				return (float64((Int32_(n.(float64)) & Int32_(float64(1)))) == float64(0))
			} else {
				panic((&js.Error{("Argument must be an integer: " + Str.X_invoke_Arity1(n).(string))}))
			}
		})
	}(&AFn{})

	Odd_QMARK_ = func(odd_QMARK_ *AFn) *AFn {
		return Fn(odd_QMARK_, 1, func(n interface{}) bool {
			return !(Even_QMARK_.Arity1IB(n))
		})
	}(&AFn{})

	Constantly = func(constantly *AFn) *AFn {
		return Fn(constantly, 1, func(x interface{}) interface{} {
			return func(G__4799 *AFn) *AFn {
				return Fn(G__4799, 0, func(args__ ...interface{}) interface{} {
					var args = Seq.Arity1IQ(args__[0])
					_ = args
					return x
				})
			}(&AFn{})
		})
	}(&AFn{})

	Comp = func(comp *AFn) *AFn {
		return Fn(comp, 3, func() interface{} {
			return Identity
		}, func(f interface{}) interface{} {
			return f
		}, func(f interface{}, g interface{}) interface{} {
			return func(G__4854 *AFn) *AFn {
				return Fn(G__4854, 3, func() interface{} {
					{
						var G__4827 = func() interface{} {
							return g.(CljsCoreIFn).X_invoke_Arity0()
						}()
						_ = G__4827
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4827)
					}
				}, func(x interface{}) interface{} {
					{
						var G__4828 = func() interface{} {
							var G__4829 = x
							_ = G__4829
							return g.(CljsCoreIFn).X_invoke_Arity1(G__4829)
						}()
						_ = G__4828
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4828)
					}
				}, func(x interface{}, y interface{}) interface{} {
					{
						var G__4830 = func() interface{} {
							var G__4831 = x
							var G__4832 = y
							_, _ = G__4831, G__4832
							return g.(CljsCoreIFn).X_invoke_Arity2(G__4831, G__4832)
						}()
						_ = G__4830
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4830)
					}
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					{
						var G__4833 = func() interface{} {
							var G__4834 = x
							var G__4835 = y
							var G__4836 = z
							_, _, _ = G__4834, G__4835, G__4836
							return g.(CljsCoreIFn).X_invoke_Arity3(G__4834, G__4835, G__4836)
						}()
						_ = G__4833
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4833)
					}
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					{
						var G__4837 = Apply.X_invoke_Arity5(g, x, y, z, args)
						_ = G__4837
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4837)
					}
				})
			}(&AFn{})
		}, func(f interface{}, g interface{}, h interface{}) interface{} {
			return func(G__4855 *AFn) *AFn {
				return Fn(G__4855, 3, func() interface{} {
					{
						var G__4838 = func() interface{} {
							var G__4839 = func() interface{} {
								return h.(CljsCoreIFn).X_invoke_Arity0()
							}()
							_ = G__4839
							return g.(CljsCoreIFn).X_invoke_Arity1(G__4839)
						}()
						_ = G__4838
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4838)
					}
				}, func(x interface{}) interface{} {
					{
						var G__4840 = func() interface{} {
							var G__4841 = func() interface{} {
								var G__4842 = x
								_ = G__4842
								return h.(CljsCoreIFn).X_invoke_Arity1(G__4842)
							}()
							_ = G__4841
							return g.(CljsCoreIFn).X_invoke_Arity1(G__4841)
						}()
						_ = G__4840
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4840)
					}
				}, func(x interface{}, y interface{}) interface{} {
					{
						var G__4843 = func() interface{} {
							var G__4844 = func() interface{} {
								var G__4845 = x
								var G__4846 = y
								_, _ = G__4845, G__4846
								return h.(CljsCoreIFn).X_invoke_Arity2(G__4845, G__4846)
							}()
							_ = G__4844
							return g.(CljsCoreIFn).X_invoke_Arity1(G__4844)
						}()
						_ = G__4843
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4843)
					}
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					{
						var G__4847 = func() interface{} {
							var G__4848 = func() interface{} {
								var G__4849 = x
								var G__4850 = y
								var G__4851 = z
								_, _, _ = G__4849, G__4850, G__4851
								return h.(CljsCoreIFn).X_invoke_Arity3(G__4849, G__4850, G__4851)
							}()
							_ = G__4848
							return g.(CljsCoreIFn).X_invoke_Arity1(G__4848)
						}()
						_ = G__4847
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4847)
					}
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					{
						var G__4852 = func() interface{} {
							var G__4853 = Apply.X_invoke_Arity5(h, x, y, z, args)
							_ = G__4853
							return g.(CljsCoreIFn).X_invoke_Arity1(G__4853)
						}()
						_ = G__4852
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4852)
					}
				})
			}(&AFn{})
		}, func(f1_f2_f3_fs__ ...interface{}) interface{} {
			var f1 = f1_f2_f3_fs__[0]
			var f2 = f1_f2_f3_fs__[1]
			var f3 = f1_f2_f3_fs__[2]
			var fs = Seq.Arity1IQ(f1_f2_f3_fs__[3])
			_, _, _, _ = f1, f2, f3, fs
			{
				var fs___1 = Reverse.X_invoke_Arity1(List_STAR_.X_invoke_Arity4(f1, f2, f3, fs).(*CljsCoreCons))
				_ = fs___1
				return func(G__4856 *AFn, fs___1 interface{}) *AFn {
					return Fn(G__4856, 0, func(args__ ...interface{}) interface{} {
						var args = Seq.Arity1IQ(args__[0])
						_ = args
						{
							var ret interface{} = Apply.X_invoke_Arity2(First.X_invoke_Arity1(fs___1), args)
							var fs___2 interface{} = Next.Arity1IQ(fs___1)
							_, _ = ret, fs___2
							for {
								if Truth_(fs___2) {
									ret, fs___2 = First.X_invoke_Arity1(fs___2).(CljsCoreIFn).X_invoke_Arity1(ret), Next.Arity1IQ(fs___2)
									continue
								} else {
									return ret
								}
							}
						}
					})
				}(&AFn{}, fs___1)
			}
		})
	}(&AFn{})

	Partial = func(partial *AFn) *AFn {
		return Fn(partial, 4, func(f interface{}) interface{} {
			return f
		}, func(f interface{}, arg1 interface{}) interface{} {
			return func(G__4857 *AFn) *AFn {
				return Fn(G__4857, 0, func(args__ ...interface{}) interface{} {
					var args = Seq.Arity1IQ(args__[0])
					_ = args
					return Apply.X_invoke_Arity3(f, arg1, args)
				})
			}(&AFn{})
		}, func(f interface{}, arg1 interface{}, arg2 interface{}) interface{} {
			return func(G__4858 *AFn) *AFn {
				return Fn(G__4858, 0, func(args__ ...interface{}) interface{} {
					var args = Seq.Arity1IQ(args__[0])
					_ = args
					return Apply.X_invoke_Arity4(f, arg1, arg2, args)
				})
			}(&AFn{})
		}, func(f interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} {
			return func(G__4859 *AFn) *AFn {
				return Fn(G__4859, 0, func(args__ ...interface{}) interface{} {
					var args = Seq.Arity1IQ(args__[0])
					_ = args
					return Apply.X_invoke_Arity5(f, arg1, arg2, arg3, args)
				})
			}(&AFn{})
		}, func(f_arg1_arg2_arg3_more__ ...interface{}) interface{} {
			var f = f_arg1_arg2_arg3_more__[0]
			var arg1 = f_arg1_arg2_arg3_more__[1]
			var arg2 = f_arg1_arg2_arg3_more__[2]
			var arg3 = f_arg1_arg2_arg3_more__[3]
			var more = Seq.Arity1IQ(f_arg1_arg2_arg3_more__[4])
			_, _, _, _, _ = f, arg1, arg2, arg3, more
			return func(G__4860 *AFn) *AFn {
				return Fn(G__4860, 0, func(args__ ...interface{}) interface{} {
					var args = Seq.Arity1IQ(args__[0])
					_ = args
					return Apply.X_invoke_Arity5(f, arg1, arg2, arg3, Concat.X_invoke_Arity2(more, args).(*CljsCoreLazySeq))
				})
			}(&AFn{})
		})
	}(&AFn{})

	Fnil = func(fnil *AFn) *AFn {
		return Fn(fnil, 4, func(f interface{}, x interface{}) interface{} {
			return func(G__4893 *AFn) *AFn {
				return Fn(G__4893, 3, func(a interface{}) interface{} {
					{
						var G__4877 = func() interface{} {
							if Nil_(a) {
								return x
							} else {
								return a
							}
						}()
						_ = G__4877
						return f.(CljsCoreIFn).X_invoke_Arity1(G__4877)
					}
				}, func(a interface{}, b interface{}) interface{} {
					{
						var G__4878 = func() interface{} {
							if Nil_(a) {
								return x
							} else {
								return a
							}
						}()
						var G__4879 = b
						_, _ = G__4878, G__4879
						return f.(CljsCoreIFn).X_invoke_Arity2(G__4878, G__4879)
					}
				}, func(a interface{}, b interface{}, c interface{}) interface{} {
					{
						var G__4880 = func() interface{} {
							if Nil_(a) {
								return x
							} else {
								return a
							}
						}()
						var G__4881 = b
						var G__4882 = c
						_, _, _ = G__4880, G__4881, G__4882
						return f.(CljsCoreIFn).X_invoke_Arity3(G__4880, G__4881, G__4882)
					}
				}, func(a_b_c_ds__ ...interface{}) interface{} {
					var a = a_b_c_ds__[0]
					var b = a_b_c_ds__[1]
					var c = a_b_c_ds__[2]
					var ds = Seq.Arity1IQ(a_b_c_ds__[3])
					_, _, _, _ = a, b, c, ds
					return Apply.X_invoke_Arity5(f, func() interface{} {
						if Nil_(a) {
							return x
						} else {
							return a
						}
					}(), b, c, ds)
				})
			}(&AFn{})
		}, func(f interface{}, x interface{}, y interface{}) interface{} {
			return func(G__4894 *AFn) *AFn {
				return Fn(G__4894, 3, func(a interface{}, b interface{}) interface{} {
					{
						var G__4883 = func() interface{} {
							if Nil_(a) {
								return x
							} else {
								return a
							}
						}()
						var G__4884 = func() interface{} {
							if Nil_(b) {
								return y
							} else {
								return b
							}
						}()
						_, _ = G__4883, G__4884
						return f.(CljsCoreIFn).X_invoke_Arity2(G__4883, G__4884)
					}
				}, func(a interface{}, b interface{}, c interface{}) interface{} {
					{
						var G__4885 = func() interface{} {
							if Nil_(a) {
								return x
							} else {
								return a
							}
						}()
						var G__4886 = func() interface{} {
							if Nil_(b) {
								return y
							} else {
								return b
							}
						}()
						var G__4887 = c
						_, _, _ = G__4885, G__4886, G__4887
						return f.(CljsCoreIFn).X_invoke_Arity3(G__4885, G__4886, G__4887)
					}
				}, func(a_b_c_ds__ ...interface{}) interface{} {
					var a = a_b_c_ds__[0]
					var b = a_b_c_ds__[1]
					var c = a_b_c_ds__[2]
					var ds = Seq.Arity1IQ(a_b_c_ds__[3])
					_, _, _, _ = a, b, c, ds
					return Apply.X_invoke_Arity5(f, func() interface{} {
						if Nil_(a) {
							return x
						} else {
							return a
						}
					}(), func() interface{} {
						if Nil_(b) {
							return y
						} else {
							return b
						}
					}(), c, ds)
				})
			}(&AFn{})
		}, func(f interface{}, x interface{}, y interface{}, z interface{}) interface{} {
			return func(G__4895 *AFn) *AFn {
				return Fn(G__4895, 3, func(a interface{}, b interface{}) interface{} {
					{
						var G__4888 = func() interface{} {
							if Nil_(a) {
								return x
							} else {
								return a
							}
						}()
						var G__4889 = func() interface{} {
							if Nil_(b) {
								return y
							} else {
								return b
							}
						}()
						_, _ = G__4888, G__4889
						return f.(CljsCoreIFn).X_invoke_Arity2(G__4888, G__4889)
					}
				}, func(a interface{}, b interface{}, c interface{}) interface{} {
					{
						var G__4890 = func() interface{} {
							if Nil_(a) {
								return x
							} else {
								return a
							}
						}()
						var G__4891 = func() interface{} {
							if Nil_(b) {
								return y
							} else {
								return b
							}
						}()
						var G__4892 = func() interface{} {
							if Nil_(c) {
								return z
							} else {
								return c
							}
						}()
						_, _, _ = G__4890, G__4891, G__4892
						return f.(CljsCoreIFn).X_invoke_Arity3(G__4890, G__4891, G__4892)
					}
				}, func(a_b_c_ds__ ...interface{}) interface{} {
					var a = a_b_c_ds__[0]
					var b = a_b_c_ds__[1]
					var c = a_b_c_ds__[2]
					var ds = Seq.Arity1IQ(a_b_c_ds__[3])
					_, _, _, _ = a, b, c, ds
					return Apply.X_invoke_Arity5(f, func() interface{} {
						if Nil_(a) {
							return x
						} else {
							return a
						}
					}(), func() interface{} {
						if Nil_(b) {
							return y
						} else {
							return b
						}
					}(), func() interface{} {
						if Nil_(c) {
							return z
						} else {
							return c
						}
					}(), ds)
				})
			}(&AFn{})
		})
	}(&AFn{})

	Map_indexed = func(map_indexed *AFn) *AFn {
		return Fn(map_indexed, 2, func(f interface{}, coll interface{}) interface{} {
			{
				var mapi *AFn
				mapi = func(mapi *AFn) *AFn {
					return Fn(mapi, 2, func(idx interface{}, coll___1 interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__4944 *AFn) *AFn {
							return Fn(G__4944, 0, func() interface{} {
								{
									var temp__4222__auto__ = Seq.Arity1IQ(coll___1)
									_ = temp__4222__auto__
									if Truth_(temp__4222__auto__) {
										{
											var s = temp__4222__auto__
											_ = s
											if Chunked_seq_QMARK_.Arity1IB(s) {
												{
													var c = Chunk_first.X_invoke_Arity1(s)
													var size = Count.X_invoke_Arity1(c).(float64)
													var b = Chunk_buffer.X_invoke_Arity1(size).(*CljsCoreChunkBuffer)
													_, _, _ = c, size, b
													{
														var n__1054__auto___4945 = size
														_ = n__1054__auto___4945
														{
															var i_4946 = float64(0)
															_ = i_4946
															for {
																if i_4946 < n__1054__auto___4945 {
																	Chunk_append.X_invoke_Arity2(b, func() interface{} {
																		var G__4940 = (idx.(float64) + i_4946)
																		var G__4941 = c.(CljsCoreIIndexed).X_nth_Arity2(i_4946)
																		_, _ = G__4940, G__4941
																		return f.(CljsCoreIFn).X_invoke_Arity2(G__4940, G__4941)
																	}())
																	i_4946 = (i_4946 + float64(1))
																	continue
																} else {
																}
																break
															}
														}
													}
													return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b), mapi.X_invoke_Arity2((idx.(float64)+size), Chunk_rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq))
												}
											} else {
												return Cons.X_invoke_Arity2(func() interface{} {
													var G__4942 = idx
													var G__4943 = First.X_invoke_Arity1(s)
													_, _ = G__4942, G__4943
													return f.(CljsCoreIFn).X_invoke_Arity2(G__4942, G__4943)
												}(), mapi.X_invoke_Arity2((idx.(float64)+float64(1)), Rest.Arity1IQ(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
											}
										}
									} else {
										return nil
									}
								}
							})
						}(&AFn{}), nil, nil})
					})
				}(&AFn{})
				_ = mapi
				return mapi.X_invoke_Arity2(float64(0), coll).(*CljsCoreLazySeq)
			}
		})
	}(&AFn{})

	Keep = func(keep *AFn) *AFn {
		return Fn(keep, 2, func(f interface{}) interface{} {
			return func(G__4965 *AFn) *AFn {
				return Fn(G__4965, 1, func(f1 interface{}) interface{} {
					return func(G__4966 *AFn) *AFn {
						return Fn(G__4966, 2, func() interface{} {
							{
								return f1.(CljsCoreIFn).X_invoke_Arity0()
							}
						}, func(result interface{}) interface{} {
							{
								var G__4959 = result
								_ = G__4959
								return f1.(CljsCoreIFn).X_invoke_Arity1(G__4959)
							}
						}, func(result interface{}, input interface{}) interface{} {
							{
								var v = func() interface{} {
									var G__4960 = input
									_ = G__4960
									return f.(CljsCoreIFn).X_invoke_Arity1(G__4960)
								}()
								_ = v
								if Nil_(v) {
									return result
								} else {
									{
										var G__4961 = result
										var G__4962 = v
										_, _ = G__4961, G__4962
										return f1.(CljsCoreIFn).X_invoke_Arity2(G__4961, G__4962)
									}
								}
							}
						})
					}(&AFn{})
				})
			}(&AFn{})
		}, func(f interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__4967 *AFn) *AFn {
				return Fn(G__4967, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								if Chunked_seq_QMARK_.Arity1IB(s) {
									{
										var c = Chunk_first.X_invoke_Arity1(s)
										var size = Count.X_invoke_Arity1(c).(float64)
										var b = Chunk_buffer.X_invoke_Arity1(size).(*CljsCoreChunkBuffer)
										_, _, _ = c, size, b
										{
											var n__1054__auto___4968 = size
											_ = n__1054__auto___4968
											{
												var i_4969 = float64(0)
												_ = i_4969
												for {
													if i_4969 < n__1054__auto___4968 {
														{
															var x_4970 = func() interface{} {
																var G__4963 = c.(CljsCoreIIndexed).X_nth_Arity2(i_4969)
																_ = G__4963
																return f.(CljsCoreIFn).X_invoke_Arity1(G__4963)
															}()
															_ = x_4970
															if Nil_(x_4970) {
															} else {
																Chunk_append.X_invoke_Arity2(b, x_4970)
															}
														}
														i_4969 = (i_4969 + float64(1))
														continue
													} else {
													}
													break
												}
											}
										}
										return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b), keep.X_invoke_Arity2(f, Chunk_rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq))
									}
								} else {
									{
										var x = func() interface{} {
											var G__4964 = First.X_invoke_Arity1(s)
											_ = G__4964
											return f.(CljsCoreIFn).X_invoke_Arity1(G__4964)
										}()
										_ = x
										if Nil_(x) {
											return keep.X_invoke_Arity2(f, Rest.Arity1IQ(s)).(*CljsCoreLazySeq)
										} else {
											return Cons.X_invoke_Arity2(x, keep.X_invoke_Arity2(f, Rest.Arity1IQ(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
										}
									}
								}
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		})
	}(&AFn{})

	X__GT_Atom = func(__GT_Atom *AFn) *AFn {
		return Fn(__GT_Atom, 4, func(state interface{}, meta interface{}, validator interface{}, watches interface{}) interface{} {
			return (&CljsCoreAtom{state, meta, validator, watches})
		})
	}(&AFn{})

	Atom = func(atom *AFn) *AFn {
		return Fn(atom, 1, func(x interface{}) interface{} {
			return (&CljsCoreAtom{x, nil, nil, nil})
		}, func(x_p__5001__ ...interface{}) interface{} {
			var x = x_p__5001__[0]
			var p__5001 = Seq.Arity1IQ(x_p__5001__[1])
			_, _ = x, p__5001
			{
				var map__5003 = p__5001
				var map__5003___1 = func() interface{} {
					if Seq_QMARK_.Arity1IB(map__5003) {
						return Apply.X_invoke_Arity2(Hash_map, map__5003)
					} else {
						return map__5003
					}
				}()
				var validator = Get.X_invoke_Arity2(map__5003___1, (&CljsCoreKeyword{Ns: nil, Name: "validator", Fqn: "validator", X_hash: float64(-1966190681)}))
				var meta = Get.X_invoke_Arity2(map__5003___1, (&CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}))
				_, _, _, _ = map__5003, map__5003___1, validator, meta
				return (&CljsCoreAtom{x, meta, validator, nil})
			}
		})
	}(&AFn{})

	Reset_BANG_ = func(reset_BANG_ *AFn) *AFn {
		return Fn(reset_BANG_, 2, func(a interface{}, new_value interface{}) interface{} {
			if Value_(a).Type().AssignableTo(reflect.TypeOf((**CljsCoreAtom)(nil)).Elem()) {
				{
					var validate = Native_get_instance_field.X_invoke_Arity2(a, "Validator")
					_ = validate
					if Nil_(validate) {
					} else {
						if Truth_(func() interface{} {
							var G__5005 = new_value
							_ = G__5005
							return validate.(CljsCoreIFn).X_invoke_Arity1(G__5005)
						}()) {
						} else {
							panic((&js.Error{("Assert failed: Validator rejected reference state\n(validate new-value)")}))
						}
					}
					{
						var old_value = Native_get_instance_field.X_invoke_Arity2(a, "State")
						_ = old_value
						Native_set_instance_field.X_invoke_Arity3(a, "State", new_value)
						if Nil_(Native_get_instance_field.X_invoke_Arity2(a, "Watches")) {
						} else {
							a.(CljsCoreIWatchable).X_notify_watches_Arity3(old_value, new_value)
						}
						return new_value
					}
				}
			} else {
				return a.(CljsCoreIReset).X_reset_BANG__Arity2(new_value)
			}
		})
	}(&AFn{})

	Swap_BANG_ = func(swap_BANG_ *AFn) *AFn {
		return Fn(swap_BANG_, 4, func(a interface{}, f interface{}) interface{} {
			if Value_(a).Type().AssignableTo(reflect.TypeOf((**CljsCoreAtom)(nil)).Elem()) {
				return Reset_BANG_.X_invoke_Arity2(a, func() interface{} {
					var G__5012 = Native_get_instance_field.X_invoke_Arity2(a, "State")
					_ = G__5012
					return f.(CljsCoreIFn).X_invoke_Arity1(G__5012)
				}())
			} else {
				return a.(CljsCoreISwap).X_swap_BANG__Arity2(f)
			}
		}, func(a interface{}, f interface{}, x interface{}) interface{} {
			if Value_(a).Type().AssignableTo(reflect.TypeOf((**CljsCoreAtom)(nil)).Elem()) {
				return Reset_BANG_.X_invoke_Arity2(a, func() interface{} {
					var G__5013 = Native_get_instance_field.X_invoke_Arity2(a, "State")
					var G__5014 = x
					_, _ = G__5013, G__5014
					return f.(CljsCoreIFn).X_invoke_Arity2(G__5013, G__5014)
				}())
			} else {
				return a.(CljsCoreISwap).X_swap_BANG__Arity3(f, x)
			}
		}, func(a interface{}, f interface{}, x interface{}, y interface{}) interface{} {
			if Value_(a).Type().AssignableTo(reflect.TypeOf((**CljsCoreAtom)(nil)).Elem()) {
				return Reset_BANG_.X_invoke_Arity2(a, func() interface{} {
					var G__5015 = Native_get_instance_field.X_invoke_Arity2(a, "State")
					var G__5016 = x
					var G__5017 = y
					_, _, _ = G__5015, G__5016, G__5017
					return f.(CljsCoreIFn).X_invoke_Arity3(G__5015, G__5016, G__5017)
				}())
			} else {
				return a.(CljsCoreISwap).X_swap_BANG__Arity4(f, x, y)
			}
		}, func(a_f_x_y_more__ ...interface{}) interface{} {
			var a = a_f_x_y_more__[0]
			var f = a_f_x_y_more__[1]
			var x = a_f_x_y_more__[2]
			var y = a_f_x_y_more__[3]
			var more = Seq.Arity1IQ(a_f_x_y_more__[4])
			_, _, _, _, _ = a, f, x, y, more
			if Value_(a).Type().AssignableTo(reflect.TypeOf((**CljsCoreAtom)(nil)).Elem()) {
				return Reset_BANG_.X_invoke_Arity2(a, Apply.X_invoke_Arity5(f, Native_get_instance_field.X_invoke_Arity2(a, "State"), x, y, more))
			} else {
				return a.(CljsCoreISwap).X_swap_BANG__Arity5(f, x, y, more)
			}
		})
	}(&AFn{})

	Compare_and_set_BANG_ = func(compare_and_set_BANG_ *AFn) *AFn {
		return Fn(compare_and_set_BANG_, 3, func(a interface{}, oldval interface{}, newval interface{}) interface{} {
			if X_EQ_.Arity2IIB(Native_get_instance_field.X_invoke_Arity2(a, "State"), oldval) {
				Reset_BANG_.X_invoke_Arity2(a, newval)
				return true
			} else {
				return false
			}
		})
	}(&AFn{})

	Set_validator_BANG_ = func(set_validator_BANG_ *AFn) *AFn {
		return Fn(set_validator_BANG_, 2, func(iref interface{}, val interface{}) interface{} {
			return func() interface{} {
				var return__5018 = val
				Native_set_instance_field.X_invoke_Arity3(iref, "Validator", return__5018)
				return return__5018
			}()
		})
	}(&AFn{})

	Get_validator = func(get_validator *AFn) *AFn {
		return Fn(get_validator, 1, func(iref interface{}) interface{} {
			return Native_get_instance_field.X_invoke_Arity2(iref, "Validator")
		})
	}(&AFn{})

	Keep_indexed = func(keep_indexed *AFn) *AFn {
		return Fn(keep_indexed, 2, func(f interface{}) interface{} {
			return func(G__5085 *AFn) *AFn {
				return Fn(G__5085, 1, func(f1 interface{}) interface{} {
					{
						var ia = Atom.X_invoke_Arity1(float64(-1)).(*CljsCoreAtom)
						_ = ia
						return func(G__5086 *AFn, ia *CljsCoreAtom) *AFn {
							return Fn(G__5086, 2, func() interface{} {
								{
									return f1.(CljsCoreIFn).X_invoke_Arity0()
								}
							}, func(result interface{}) interface{} {
								{
									var G__5052 = result
									_ = G__5052
									return f1.(CljsCoreIFn).X_invoke_Arity1(G__5052)
								}
							}, func(result interface{}, input interface{}) interface{} {
								{
									var i = Swap_BANG_.X_invoke_Arity2(ia, Inc)
									var v = func() interface{} {
										var G__5053 = i
										var G__5054 = input
										_, _ = G__5053, G__5054
										return f.(CljsCoreIFn).X_invoke_Arity2(G__5053, G__5054)
									}()
									_, _ = i, v
									if Nil_(v) {
										return result
									} else {
										{
											var G__5055 = result
											var G__5056 = v
											_, _ = G__5055, G__5056
											return f1.(CljsCoreIFn).X_invoke_Arity2(G__5055, G__5056)
										}
									}
								}
							})
						}(&AFn{}, ia)
					}
				})
			}(&AFn{})
		}, func(f interface{}, coll interface{}) interface{} {
			{
				var keepi *AFn
				keepi = func(keepi *AFn) *AFn {
					return Fn(keepi, 2, func(idx interface{}, coll___1 interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__5087 *AFn) *AFn {
							return Fn(G__5087, 0, func() interface{} {
								{
									var temp__4222__auto__ = Seq.Arity1IQ(coll___1)
									_ = temp__4222__auto__
									if Truth_(temp__4222__auto__) {
										{
											var s = temp__4222__auto__
											_ = s
											if Chunked_seq_QMARK_.Arity1IB(s) {
												{
													var c = Chunk_first.X_invoke_Arity1(s)
													var size = Count.X_invoke_Arity1(c).(float64)
													var b = Chunk_buffer.X_invoke_Arity1(size).(*CljsCoreChunkBuffer)
													_, _, _ = c, size, b
													{
														var n__1054__auto___5088 = size
														_ = n__1054__auto___5088
														{
															var i_5089 = float64(0)
															_ = i_5089
															for {
																if i_5089 < n__1054__auto___5088 {
																	{
																		var x_5090 = func() interface{} {
																			var G__5081 = (idx.(float64) + i_5089)
																			var G__5082 = c.(CljsCoreIIndexed).X_nth_Arity2(i_5089)
																			_, _ = G__5081, G__5082
																			return f.(CljsCoreIFn).X_invoke_Arity2(G__5081, G__5082)
																		}()
																		_ = x_5090
																		if Nil_(x_5090) {
																		} else {
																			Chunk_append.X_invoke_Arity2(b, x_5090)
																		}
																	}
																	i_5089 = (i_5089 + float64(1))
																	continue
																} else {
																}
																break
															}
														}
													}
													return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b), keepi.X_invoke_Arity2((idx.(float64)+size), Chunk_rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq))
												}
											} else {
												{
													var x = func() interface{} {
														var G__5083 = idx
														var G__5084 = First.X_invoke_Arity1(s)
														_, _ = G__5083, G__5084
														return f.(CljsCoreIFn).X_invoke_Arity2(G__5083, G__5084)
													}()
													_ = x
													if Nil_(x) {
														return keepi.X_invoke_Arity2((idx.(float64) + float64(1)), Rest.Arity1IQ(s)).(*CljsCoreLazySeq)
													} else {
														return Cons.X_invoke_Arity2(x, keepi.X_invoke_Arity2((idx.(float64)+float64(1)), Rest.Arity1IQ(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
													}
												}
											}
										}
									} else {
										return nil
									}
								}
							})
						}(&AFn{}), nil, nil})
					})
				}(&AFn{})
				_ = keepi
				return keepi.X_invoke_Arity2(float64(0), coll).(*CljsCoreLazySeq)
			}
		})
	}(&AFn{})

	Every_pred = func(every_pred *AFn) *AFn {
		return Fn(every_pred, 3, func(p interface{}) interface{} {
			return func(ep1 *AFn) *AFn {
				return Fn(ep1, 3, func() interface{} {
					return true
				}, func(x interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var G__5304 = x
						_ = G__5304
						return p.(CljsCoreIFn).X_invoke_Arity1(G__5304)
					}())
				}, func(x interface{}, y interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__159__auto__ = func() interface{} {
							var G__5306 = x
							_ = G__5306
							return p.(CljsCoreIFn).X_invoke_Arity1(G__5306)
						}()
						_ = and__159__auto__
						if Truth_(and__159__auto__) {
							{
								var G__5307 = y
								_ = G__5307
								return p.(CljsCoreIFn).X_invoke_Arity1(G__5307)
							}
						} else {
							return and__159__auto__
						}
					}())
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__159__auto__ = func() interface{} {
							var G__5309 = x
							_ = G__5309
							return p.(CljsCoreIFn).X_invoke_Arity1(G__5309)
						}()
						_ = and__159__auto__
						if Truth_(and__159__auto__) {
							{
								var and__159__auto_____1 = func() interface{} {
									var G__5311 = y
									_ = G__5311
									return p.(CljsCoreIFn).X_invoke_Arity1(G__5311)
								}()
								_ = and__159__auto_____1
								if Truth_(and__159__auto_____1) {
									{
										var G__5312 = z
										_ = G__5312
										return p.(CljsCoreIFn).X_invoke_Arity1(G__5312)
									}
								} else {
									return and__159__auto_____1
								}
							}
						} else {
							return and__159__auto__
						}
					}())
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					return Boolean.Arity1IB((Truth_(ep1.X_invoke_Arity3(x, y, z))) && (Every_QMARK_.Arity2IIB(p, args)))
				})
			}(&AFn{})
		}, func(p1 interface{}, p2 interface{}) interface{} {
			return func(ep2 *AFn) *AFn {
				return Fn(ep2, 3, func() interface{} {
					return true
				}, func(x interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__159__auto__ = func() interface{} {
							var G__5344 = x
							_ = G__5344
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5344)
						}()
						_ = and__159__auto__
						if Truth_(and__159__auto__) {
							{
								var G__5345 = x
								_ = G__5345
								return p2.(CljsCoreIFn).X_invoke_Arity1(G__5345)
							}
						} else {
							return and__159__auto__
						}
					}())
				}, func(x interface{}, y interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__159__auto__ = func() interface{} {
							var G__5347 = x
							_ = G__5347
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5347)
						}()
						_ = and__159__auto__
						if Truth_(and__159__auto__) {
							{
								var and__159__auto_____1 = func() interface{} {
									var G__5349 = y
									_ = G__5349
									return p1.(CljsCoreIFn).X_invoke_Arity1(G__5349)
								}()
								_ = and__159__auto_____1
								if Truth_(and__159__auto_____1) {
									{
										var and__159__auto_____2 = func() interface{} {
											var G__5351 = x
											_ = G__5351
											return p2.(CljsCoreIFn).X_invoke_Arity1(G__5351)
										}()
										_ = and__159__auto_____2
										if Truth_(and__159__auto_____2) {
											{
												var G__5352 = y
												_ = G__5352
												return p2.(CljsCoreIFn).X_invoke_Arity1(G__5352)
											}
										} else {
											return and__159__auto_____2
										}
									}
								} else {
									return and__159__auto_____1
								}
							}
						} else {
							return and__159__auto__
						}
					}())
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__159__auto__ = func() interface{} {
							var G__5354 = x
							_ = G__5354
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5354)
						}()
						_ = and__159__auto__
						if Truth_(and__159__auto__) {
							{
								var and__159__auto_____1 = func() interface{} {
									var G__5356 = y
									_ = G__5356
									return p1.(CljsCoreIFn).X_invoke_Arity1(G__5356)
								}()
								_ = and__159__auto_____1
								if Truth_(and__159__auto_____1) {
									{
										var and__159__auto_____2 = func() interface{} {
											var G__5358 = z
											_ = G__5358
											return p1.(CljsCoreIFn).X_invoke_Arity1(G__5358)
										}()
										_ = and__159__auto_____2
										if Truth_(and__159__auto_____2) {
											{
												var and__159__auto_____3 = func() interface{} {
													var G__5360 = x
													_ = G__5360
													return p2.(CljsCoreIFn).X_invoke_Arity1(G__5360)
												}()
												_ = and__159__auto_____3
												if Truth_(and__159__auto_____3) {
													{
														var and__159__auto_____4 = func() interface{} {
															var G__5362 = y
															_ = G__5362
															return p2.(CljsCoreIFn).X_invoke_Arity1(G__5362)
														}()
														_ = and__159__auto_____4
														if Truth_(and__159__auto_____4) {
															{
																var G__5363 = z
																_ = G__5363
																return p2.(CljsCoreIFn).X_invoke_Arity1(G__5363)
															}
														} else {
															return and__159__auto_____4
														}
													}
												} else {
													return and__159__auto_____3
												}
											}
										} else {
											return and__159__auto_____2
										}
									}
								} else {
									return and__159__auto_____1
								}
							}
						} else {
							return and__159__auto__
						}
					}())
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					return Boolean.Arity1IB((Truth_(ep2.X_invoke_Arity3(x, y, z))) && (Every_QMARK_.Arity2IIB(func(G__5481 *AFn) *AFn {
						return Fn(G__5481, 1, func(p1__5091_SHARP_ interface{}) interface{} {
							{
								var and__159__auto__ = func() interface{} {
									var G__5368 = p1__5091_SHARP_
									_ = G__5368
									return p1.(CljsCoreIFn).X_invoke_Arity1(G__5368)
								}()
								_ = and__159__auto__
								if Truth_(and__159__auto__) {
									{
										var G__5369 = p1__5091_SHARP_
										_ = G__5369
										return p2.(CljsCoreIFn).X_invoke_Arity1(G__5369)
									}
								} else {
									return and__159__auto__
								}
							}
						})
					}(&AFn{}), args)))
				})
			}(&AFn{})
		}, func(p1 interface{}, p2 interface{}, p3 interface{}) interface{} {
			return func(ep3 *AFn) *AFn {
				return Fn(ep3, 3, func() interface{} {
					return true
				}, func(x interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__159__auto__ = func() interface{} {
							var G__5415 = x
							_ = G__5415
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5415)
						}()
						_ = and__159__auto__
						if Truth_(and__159__auto__) {
							{
								var and__159__auto_____1 = func() interface{} {
									var G__5417 = x
									_ = G__5417
									return p2.(CljsCoreIFn).X_invoke_Arity1(G__5417)
								}()
								_ = and__159__auto_____1
								if Truth_(and__159__auto_____1) {
									{
										var G__5418 = x
										_ = G__5418
										return p3.(CljsCoreIFn).X_invoke_Arity1(G__5418)
									}
								} else {
									return and__159__auto_____1
								}
							}
						} else {
							return and__159__auto__
						}
					}())
				}, func(x interface{}, y interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__159__auto__ = func() interface{} {
							var G__5420 = x
							_ = G__5420
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5420)
						}()
						_ = and__159__auto__
						if Truth_(and__159__auto__) {
							{
								var and__159__auto_____1 = func() interface{} {
									var G__5422 = x
									_ = G__5422
									return p2.(CljsCoreIFn).X_invoke_Arity1(G__5422)
								}()
								_ = and__159__auto_____1
								if Truth_(and__159__auto_____1) {
									{
										var and__159__auto_____2 = func() interface{} {
											var G__5424 = x
											_ = G__5424
											return p3.(CljsCoreIFn).X_invoke_Arity1(G__5424)
										}()
										_ = and__159__auto_____2
										if Truth_(and__159__auto_____2) {
											{
												var and__159__auto_____3 = func() interface{} {
													var G__5426 = y
													_ = G__5426
													return p1.(CljsCoreIFn).X_invoke_Arity1(G__5426)
												}()
												_ = and__159__auto_____3
												if Truth_(and__159__auto_____3) {
													{
														var and__159__auto_____4 = func() interface{} {
															var G__5428 = y
															_ = G__5428
															return p2.(CljsCoreIFn).X_invoke_Arity1(G__5428)
														}()
														_ = and__159__auto_____4
														if Truth_(and__159__auto_____4) {
															{
																var G__5429 = y
																_ = G__5429
																return p3.(CljsCoreIFn).X_invoke_Arity1(G__5429)
															}
														} else {
															return and__159__auto_____4
														}
													}
												} else {
													return and__159__auto_____3
												}
											}
										} else {
											return and__159__auto_____2
										}
									}
								} else {
									return and__159__auto_____1
								}
							}
						} else {
							return and__159__auto__
						}
					}())
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__159__auto__ = func() interface{} {
							var G__5431 = x
							_ = G__5431
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5431)
						}()
						_ = and__159__auto__
						if Truth_(and__159__auto__) {
							{
								var and__159__auto_____1 = func() interface{} {
									var G__5433 = x
									_ = G__5433
									return p2.(CljsCoreIFn).X_invoke_Arity1(G__5433)
								}()
								_ = and__159__auto_____1
								if Truth_(and__159__auto_____1) {
									{
										var and__159__auto_____2 = func() interface{} {
											var G__5435 = x
											_ = G__5435
											return p3.(CljsCoreIFn).X_invoke_Arity1(G__5435)
										}()
										_ = and__159__auto_____2
										if Truth_(and__159__auto_____2) {
											{
												var and__159__auto_____3 = func() interface{} {
													var G__5437 = y
													_ = G__5437
													return p1.(CljsCoreIFn).X_invoke_Arity1(G__5437)
												}()
												_ = and__159__auto_____3
												if Truth_(and__159__auto_____3) {
													{
														var and__159__auto_____4 = func() interface{} {
															var G__5439 = y
															_ = G__5439
															return p2.(CljsCoreIFn).X_invoke_Arity1(G__5439)
														}()
														_ = and__159__auto_____4
														if Truth_(and__159__auto_____4) {
															{
																var and__159__auto_____5 = func() interface{} {
																	var G__5441 = y
																	_ = G__5441
																	return p3.(CljsCoreIFn).X_invoke_Arity1(G__5441)
																}()
																_ = and__159__auto_____5
																if Truth_(and__159__auto_____5) {
																	{
																		var and__159__auto_____6 = func() interface{} {
																			var G__5443 = z
																			_ = G__5443
																			return p1.(CljsCoreIFn).X_invoke_Arity1(G__5443)
																		}()
																		_ = and__159__auto_____6
																		if Truth_(and__159__auto_____6) {
																			{
																				var and__159__auto_____7 = func() interface{} {
																					var G__5445 = z
																					_ = G__5445
																					return p2.(CljsCoreIFn).X_invoke_Arity1(G__5445)
																				}()
																				_ = and__159__auto_____7
																				if Truth_(and__159__auto_____7) {
																					{
																						var G__5446 = z
																						_ = G__5446
																						return p3.(CljsCoreIFn).X_invoke_Arity1(G__5446)
																					}
																				} else {
																					return and__159__auto_____7
																				}
																			}
																		} else {
																			return and__159__auto_____6
																		}
																	}
																} else {
																	return and__159__auto_____5
																}
															}
														} else {
															return and__159__auto_____4
														}
													}
												} else {
													return and__159__auto_____3
												}
											}
										} else {
											return and__159__auto_____2
										}
									}
								} else {
									return and__159__auto_____1
								}
							}
						} else {
							return and__159__auto__
						}
					}())
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					return Boolean.Arity1IB((Truth_(ep3.X_invoke_Arity3(x, y, z))) && (Every_QMARK_.Arity2IIB(func(G__5482 *AFn) *AFn {
						return Fn(G__5482, 1, func(p1__5092_SHARP_ interface{}) interface{} {
							{
								var and__159__auto__ = func() interface{} {
									var G__5453 = p1__5092_SHARP_
									_ = G__5453
									return p1.(CljsCoreIFn).X_invoke_Arity1(G__5453)
								}()
								_ = and__159__auto__
								if Truth_(and__159__auto__) {
									{
										var and__159__auto_____1 = func() interface{} {
											var G__5455 = p1__5092_SHARP_
											_ = G__5455
											return p2.(CljsCoreIFn).X_invoke_Arity1(G__5455)
										}()
										_ = and__159__auto_____1
										if Truth_(and__159__auto_____1) {
											{
												var G__5456 = p1__5092_SHARP_
												_ = G__5456
												return p3.(CljsCoreIFn).X_invoke_Arity1(G__5456)
											}
										} else {
											return and__159__auto_____1
										}
									}
								} else {
									return and__159__auto__
								}
							}
						})
					}(&AFn{}), args)))
				})
			}(&AFn{})
		}, func(p1_p2_p3_ps__ ...interface{}) interface{} {
			var p1 = p1_p2_p3_ps__[0]
			var p2 = p1_p2_p3_ps__[1]
			var p3 = p1_p2_p3_ps__[2]
			var ps = Seq.Arity1IQ(p1_p2_p3_ps__[3])
			_, _, _, _ = p1, p2, p3, ps
			{
				var ps___1 = List_STAR_.X_invoke_Arity4(p1, p2, p3, ps).(*CljsCoreCons)
				_ = ps___1
				return func(epn *AFn, ps___1 *CljsCoreCons) *AFn {
					return Fn(epn, 3, func() interface{} {
						return true
					}, func(x interface{}) interface{} {
						return Every_QMARK_.Arity2IIB(func(G__5483 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__5483, 1, func(p1__5093_SHARP_ interface{}) interface{} {
								{
									var G__5472 = x
									_ = G__5472
									return p1__5093_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5472)
								}
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x interface{}, y interface{}) interface{} {
						return Every_QMARK_.Arity2IIB(func(G__5484 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__5484, 1, func(p1__5094_SHARP_ interface{}) interface{} {
								{
									var and__159__auto__ = func() interface{} {
										var G__5474 = x
										_ = G__5474
										return p1__5094_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5474)
									}()
									_ = and__159__auto__
									if Truth_(and__159__auto__) {
										{
											var G__5475 = y
											_ = G__5475
											return p1__5094_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5475)
										}
									} else {
										return and__159__auto__
									}
								}
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return Every_QMARK_.Arity2IIB(func(G__5485 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__5485, 1, func(p1__5095_SHARP_ interface{}) interface{} {
								{
									var and__159__auto__ = func() interface{} {
										var G__5477 = x
										_ = G__5477
										return p1__5095_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5477)
									}()
									_ = and__159__auto__
									if Truth_(and__159__auto__) {
										{
											var and__159__auto_____1 = func() interface{} {
												var G__5479 = y
												_ = G__5479
												return p1__5095_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5479)
											}()
											_ = and__159__auto_____1
											if Truth_(and__159__auto_____1) {
												{
													var G__5480 = z
													_ = G__5480
													return p1__5095_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5480)
												}
											} else {
												return and__159__auto_____1
											}
										}
									} else {
										return and__159__auto__
									}
								}
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x_y_z_args__ ...interface{}) interface{} {
						var x = x_y_z_args__[0]
						var y = x_y_z_args__[1]
						var z = x_y_z_args__[2]
						var args = Seq.Arity1IQ(x_y_z_args__[3])
						_, _, _, _ = x, y, z, args
						return Boolean.Arity1IB((Truth_(epn.X_invoke_Arity3(x, y, z))) && (Every_QMARK_.Arity2IIB(func(G__5486 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__5486, 1, func(p1__5096_SHARP_ interface{}) interface{} {
								return Every_QMARK_.Arity2IIB(p1__5096_SHARP_, args)
							})
						}(&AFn{}, ps___1), ps___1)))
					})
				}(&AFn{}, ps___1)
			}
		})
	}(&AFn{})

	Some_fn = func(some_fn *AFn) *AFn {
		return Fn(some_fn, 3, func(p interface{}) interface{} {
			return func(sp1 *AFn) *AFn {
				return Fn(sp1, 3, func() interface{} {
					return nil
				}, func(x interface{}) interface{} {
					{
						var G__5692 = x
						_ = G__5692
						return p.(CljsCoreIFn).X_invoke_Arity1(G__5692)
					}
				}, func(x interface{}, y interface{}) interface{} {
					{
						var or__171__auto__ = func() interface{} {
							var G__5694 = x
							_ = G__5694
							return p.(CljsCoreIFn).X_invoke_Arity1(G__5694)
						}()
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							{
								var G__5695 = y
								_ = G__5695
								return p.(CljsCoreIFn).X_invoke_Arity1(G__5695)
							}
						}
					}
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					{
						var or__171__auto__ = func() interface{} {
							var G__5697 = x
							_ = G__5697
							return p.(CljsCoreIFn).X_invoke_Arity1(G__5697)
						}()
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							{
								var or__171__auto_____1 = func() interface{} {
									var G__5699 = y
									_ = G__5699
									return p.(CljsCoreIFn).X_invoke_Arity1(G__5699)
								}()
								_ = or__171__auto_____1
								if Truth_(or__171__auto_____1) {
									return or__171__auto_____1
								} else {
									{
										var G__5700 = z
										_ = G__5700
										return p.(CljsCoreIFn).X_invoke_Arity1(G__5700)
									}
								}
							}
						}
					}
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					{
						var or__171__auto__ = sp1.X_invoke_Arity3(x, y, z)
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							return Some.X_invoke_Arity2(p, args)
						}
					}
				})
			}(&AFn{})
		}, func(p1 interface{}, p2 interface{}) interface{} {
			return func(sp2 *AFn) *AFn {
				return Fn(sp2, 3, func() interface{} {
					return nil
				}, func(x interface{}) interface{} {
					{
						var or__171__auto__ = func() interface{} {
							var G__5732 = x
							_ = G__5732
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5732)
						}()
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							{
								var G__5733 = x
								_ = G__5733
								return p2.(CljsCoreIFn).X_invoke_Arity1(G__5733)
							}
						}
					}
				}, func(x interface{}, y interface{}) interface{} {
					{
						var or__171__auto__ = func() interface{} {
							var G__5735 = x
							_ = G__5735
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5735)
						}()
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							{
								var or__171__auto_____1 = func() interface{} {
									var G__5737 = y
									_ = G__5737
									return p1.(CljsCoreIFn).X_invoke_Arity1(G__5737)
								}()
								_ = or__171__auto_____1
								if Truth_(or__171__auto_____1) {
									return or__171__auto_____1
								} else {
									{
										var or__171__auto_____2 = func() interface{} {
											var G__5739 = x
											_ = G__5739
											return p2.(CljsCoreIFn).X_invoke_Arity1(G__5739)
										}()
										_ = or__171__auto_____2
										if Truth_(or__171__auto_____2) {
											return or__171__auto_____2
										} else {
											{
												var G__5740 = y
												_ = G__5740
												return p2.(CljsCoreIFn).X_invoke_Arity1(G__5740)
											}
										}
									}
								}
							}
						}
					}
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					{
						var or__171__auto__ = func() interface{} {
							var G__5742 = x
							_ = G__5742
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5742)
						}()
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							{
								var or__171__auto_____1 = func() interface{} {
									var G__5744 = y
									_ = G__5744
									return p1.(CljsCoreIFn).X_invoke_Arity1(G__5744)
								}()
								_ = or__171__auto_____1
								if Truth_(or__171__auto_____1) {
									return or__171__auto_____1
								} else {
									{
										var or__171__auto_____2 = func() interface{} {
											var G__5746 = z
											_ = G__5746
											return p1.(CljsCoreIFn).X_invoke_Arity1(G__5746)
										}()
										_ = or__171__auto_____2
										if Truth_(or__171__auto_____2) {
											return or__171__auto_____2
										} else {
											{
												var or__171__auto_____3 = func() interface{} {
													var G__5748 = x
													_ = G__5748
													return p2.(CljsCoreIFn).X_invoke_Arity1(G__5748)
												}()
												_ = or__171__auto_____3
												if Truth_(or__171__auto_____3) {
													return or__171__auto_____3
												} else {
													{
														var or__171__auto_____4 = func() interface{} {
															var G__5750 = y
															_ = G__5750
															return p2.(CljsCoreIFn).X_invoke_Arity1(G__5750)
														}()
														_ = or__171__auto_____4
														if Truth_(or__171__auto_____4) {
															return or__171__auto_____4
														} else {
															{
																var G__5751 = z
																_ = G__5751
																return p2.(CljsCoreIFn).X_invoke_Arity1(G__5751)
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					{
						var or__171__auto__ = sp2.X_invoke_Arity3(x, y, z)
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							return Some.X_invoke_Arity2(func(G__5861 *AFn, or__171__auto__ interface{}) *AFn {
								return Fn(G__5861, 1, func(p1__5487_SHARP_ interface{}) interface{} {
									{
										var or__171__auto_____1 = func() interface{} {
											var G__5753 = p1__5487_SHARP_
											_ = G__5753
											return p1.(CljsCoreIFn).X_invoke_Arity1(G__5753)
										}()
										_ = or__171__auto_____1
										if Truth_(or__171__auto_____1) {
											return or__171__auto_____1
										} else {
											{
												var G__5754 = p1__5487_SHARP_
												_ = G__5754
												return p2.(CljsCoreIFn).X_invoke_Arity1(G__5754)
											}
										}
									}
								})
							}(&AFn{}, or__171__auto__), args)
						}
					}
				})
			}(&AFn{})
		}, func(p1 interface{}, p2 interface{}, p3 interface{}) interface{} {
			return func(sp3 *AFn) *AFn {
				return Fn(sp3, 3, func() interface{} {
					return nil
				}, func(x interface{}) interface{} {
					{
						var or__171__auto__ = func() interface{} {
							var G__5800 = x
							_ = G__5800
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5800)
						}()
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							{
								var or__171__auto_____1 = func() interface{} {
									var G__5802 = x
									_ = G__5802
									return p2.(CljsCoreIFn).X_invoke_Arity1(G__5802)
								}()
								_ = or__171__auto_____1
								if Truth_(or__171__auto_____1) {
									return or__171__auto_____1
								} else {
									{
										var G__5803 = x
										_ = G__5803
										return p3.(CljsCoreIFn).X_invoke_Arity1(G__5803)
									}
								}
							}
						}
					}
				}, func(x interface{}, y interface{}) interface{} {
					{
						var or__171__auto__ = func() interface{} {
							var G__5805 = x
							_ = G__5805
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5805)
						}()
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							{
								var or__171__auto_____1 = func() interface{} {
									var G__5807 = x
									_ = G__5807
									return p2.(CljsCoreIFn).X_invoke_Arity1(G__5807)
								}()
								_ = or__171__auto_____1
								if Truth_(or__171__auto_____1) {
									return or__171__auto_____1
								} else {
									{
										var or__171__auto_____2 = func() interface{} {
											var G__5809 = x
											_ = G__5809
											return p3.(CljsCoreIFn).X_invoke_Arity1(G__5809)
										}()
										_ = or__171__auto_____2
										if Truth_(or__171__auto_____2) {
											return or__171__auto_____2
										} else {
											{
												var or__171__auto_____3 = func() interface{} {
													var G__5811 = y
													_ = G__5811
													return p1.(CljsCoreIFn).X_invoke_Arity1(G__5811)
												}()
												_ = or__171__auto_____3
												if Truth_(or__171__auto_____3) {
													return or__171__auto_____3
												} else {
													{
														var or__171__auto_____4 = func() interface{} {
															var G__5813 = y
															_ = G__5813
															return p2.(CljsCoreIFn).X_invoke_Arity1(G__5813)
														}()
														_ = or__171__auto_____4
														if Truth_(or__171__auto_____4) {
															return or__171__auto_____4
														} else {
															{
																var G__5814 = y
																_ = G__5814
																return p3.(CljsCoreIFn).X_invoke_Arity1(G__5814)
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					{
						var or__171__auto__ = func() interface{} {
							var G__5816 = x
							_ = G__5816
							return p1.(CljsCoreIFn).X_invoke_Arity1(G__5816)
						}()
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							{
								var or__171__auto_____1 = func() interface{} {
									var G__5818 = x
									_ = G__5818
									return p2.(CljsCoreIFn).X_invoke_Arity1(G__5818)
								}()
								_ = or__171__auto_____1
								if Truth_(or__171__auto_____1) {
									return or__171__auto_____1
								} else {
									{
										var or__171__auto_____2 = func() interface{} {
											var G__5820 = x
											_ = G__5820
											return p3.(CljsCoreIFn).X_invoke_Arity1(G__5820)
										}()
										_ = or__171__auto_____2
										if Truth_(or__171__auto_____2) {
											return or__171__auto_____2
										} else {
											{
												var or__171__auto_____3 = func() interface{} {
													var G__5822 = y
													_ = G__5822
													return p1.(CljsCoreIFn).X_invoke_Arity1(G__5822)
												}()
												_ = or__171__auto_____3
												if Truth_(or__171__auto_____3) {
													return or__171__auto_____3
												} else {
													{
														var or__171__auto_____4 = func() interface{} {
															var G__5824 = y
															_ = G__5824
															return p2.(CljsCoreIFn).X_invoke_Arity1(G__5824)
														}()
														_ = or__171__auto_____4
														if Truth_(or__171__auto_____4) {
															return or__171__auto_____4
														} else {
															{
																var or__171__auto_____5 = func() interface{} {
																	var G__5826 = y
																	_ = G__5826
																	return p3.(CljsCoreIFn).X_invoke_Arity1(G__5826)
																}()
																_ = or__171__auto_____5
																if Truth_(or__171__auto_____5) {
																	return or__171__auto_____5
																} else {
																	{
																		var or__171__auto_____6 = func() interface{} {
																			var G__5828 = z
																			_ = G__5828
																			return p1.(CljsCoreIFn).X_invoke_Arity1(G__5828)
																		}()
																		_ = or__171__auto_____6
																		if Truth_(or__171__auto_____6) {
																			return or__171__auto_____6
																		} else {
																			{
																				var or__171__auto_____7 = func() interface{} {
																					var G__5830 = z
																					_ = G__5830
																					return p2.(CljsCoreIFn).X_invoke_Arity1(G__5830)
																				}()
																				_ = or__171__auto_____7
																				if Truth_(or__171__auto_____7) {
																					return or__171__auto_____7
																				} else {
																					{
																						var G__5831 = z
																						_ = G__5831
																						return p3.(CljsCoreIFn).X_invoke_Arity1(G__5831)
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					{
						var or__171__auto__ = sp3.X_invoke_Arity3(x, y, z)
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							return Some.X_invoke_Arity2(func(G__5862 *AFn, or__171__auto__ interface{}) *AFn {
								return Fn(G__5862, 1, func(p1__5488_SHARP_ interface{}) interface{} {
									{
										var or__171__auto_____1 = func() interface{} {
											var G__5833 = p1__5488_SHARP_
											_ = G__5833
											return p1.(CljsCoreIFn).X_invoke_Arity1(G__5833)
										}()
										_ = or__171__auto_____1
										if Truth_(or__171__auto_____1) {
											return or__171__auto_____1
										} else {
											{
												var or__171__auto_____2 = func() interface{} {
													var G__5835 = p1__5488_SHARP_
													_ = G__5835
													return p2.(CljsCoreIFn).X_invoke_Arity1(G__5835)
												}()
												_ = or__171__auto_____2
												if Truth_(or__171__auto_____2) {
													return or__171__auto_____2
												} else {
													{
														var G__5836 = p1__5488_SHARP_
														_ = G__5836
														return p3.(CljsCoreIFn).X_invoke_Arity1(G__5836)
													}
												}
											}
										}
									}
								})
							}(&AFn{}, or__171__auto__), args)
						}
					}
				})
			}(&AFn{})
		}, func(p1_p2_p3_ps__ ...interface{}) interface{} {
			var p1 = p1_p2_p3_ps__[0]
			var p2 = p1_p2_p3_ps__[1]
			var p3 = p1_p2_p3_ps__[2]
			var ps = Seq.Arity1IQ(p1_p2_p3_ps__[3])
			_, _, _, _ = p1, p2, p3, ps
			{
				var ps___1 = List_STAR_.X_invoke_Arity4(p1, p2, p3, ps).(*CljsCoreCons)
				_ = ps___1
				return func(spn *AFn, ps___1 *CljsCoreCons) *AFn {
					return Fn(spn, 3, func() interface{} {
						return nil
					}, func(x interface{}) interface{} {
						return Some.X_invoke_Arity2(func(G__5863 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__5863, 1, func(p1__5489_SHARP_ interface{}) interface{} {
								{
									var G__5852 = x
									_ = G__5852
									return p1__5489_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5852)
								}
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x interface{}, y interface{}) interface{} {
						return Some.X_invoke_Arity2(func(G__5864 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__5864, 1, func(p1__5490_SHARP_ interface{}) interface{} {
								{
									var or__171__auto__ = func() interface{} {
										var G__5854 = x
										_ = G__5854
										return p1__5490_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5854)
									}()
									_ = or__171__auto__
									if Truth_(or__171__auto__) {
										return or__171__auto__
									} else {
										{
											var G__5855 = y
											_ = G__5855
											return p1__5490_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5855)
										}
									}
								}
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return Some.X_invoke_Arity2(func(G__5865 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__5865, 1, func(p1__5491_SHARP_ interface{}) interface{} {
								{
									var or__171__auto__ = func() interface{} {
										var G__5857 = x
										_ = G__5857
										return p1__5491_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5857)
									}()
									_ = or__171__auto__
									if Truth_(or__171__auto__) {
										return or__171__auto__
									} else {
										{
											var or__171__auto_____1 = func() interface{} {
												var G__5859 = y
												_ = G__5859
												return p1__5491_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5859)
											}()
											_ = or__171__auto_____1
											if Truth_(or__171__auto_____1) {
												return or__171__auto_____1
											} else {
												{
													var G__5860 = z
													_ = G__5860
													return p1__5491_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__5860)
												}
											}
										}
									}
								}
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x_y_z_args__ ...interface{}) interface{} {
						var x = x_y_z_args__[0]
						var y = x_y_z_args__[1]
						var z = x_y_z_args__[2]
						var args = Seq.Arity1IQ(x_y_z_args__[3])
						_, _, _, _ = x, y, z, args
						{
							var or__171__auto__ = spn.X_invoke_Arity3(x, y, z)
							_ = or__171__auto__
							if Truth_(or__171__auto__) {
								return or__171__auto__
							} else {
								return Some.X_invoke_Arity2(func(G__5866 *AFn, or__171__auto__ interface{}, ps___1 *CljsCoreCons) *AFn {
									return Fn(G__5866, 1, func(p1__5492_SHARP_ interface{}) interface{} {
										return Some.X_invoke_Arity2(p1__5492_SHARP_, args)
									})
								}(&AFn{}, or__171__auto__, ps___1), ps___1)
							}
						}
					})
				}(&AFn{}, ps___1)
			}
		})
	}(&AFn{})

	Map_ = func(map_ *AFn) *AFn {
		return Fn(map_, 4, func(f interface{}) interface{} {
			return func(G__5921 *AFn) *AFn {
				return Fn(G__5921, 1, func(f1 interface{}) interface{} {
					return func(G__5922 *AFn) *AFn {
						return Fn(G__5922, 2, func() interface{} {
							{
								return f1.(CljsCoreIFn).X_invoke_Arity0()
							}
						}, func(result interface{}) interface{} {
							{
								var G__5907 = result
								_ = G__5907
								return f1.(CljsCoreIFn).X_invoke_Arity1(G__5907)
							}
						}, func(result interface{}, input interface{}) interface{} {
							{
								var G__5908 = result
								var G__5909 = func() interface{} {
									var G__5910 = input
									_ = G__5910
									return f.(CljsCoreIFn).X_invoke_Arity1(G__5910)
								}()
								_, _ = G__5908, G__5909
								return f1.(CljsCoreIFn).X_invoke_Arity2(G__5908, G__5909)
							}
						}, func(result_input_inputs__ ...interface{}) interface{} {
							var result = result_input_inputs__[0]
							var input = result_input_inputs__[1]
							var inputs = Seq.Arity1IQ(result_input_inputs__[2])
							_, _, _ = result, input, inputs
							{
								var G__5911 = result
								var G__5912 = Apply.X_invoke_Arity3(f, input, inputs)
								_, _ = G__5911, G__5912
								return f1.(CljsCoreIFn).X_invoke_Arity2(G__5911, G__5912)
							}
						})
					}(&AFn{})
				})
			}(&AFn{})
		}, func(f interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__5923 *AFn) *AFn {
				return Fn(G__5923, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								if Chunked_seq_QMARK_.Arity1IB(s) {
									{
										var c = Chunk_first.X_invoke_Arity1(s)
										var size = Count.X_invoke_Arity1(c).(float64)
										var b = Chunk_buffer.X_invoke_Arity1(size).(*CljsCoreChunkBuffer)
										_, _, _ = c, size, b
										{
											var n__1054__auto___5924 = size
											_ = n__1054__auto___5924
											{
												var i_5925 = float64(0)
												_ = i_5925
												for {
													if i_5925 < n__1054__auto___5924 {
														Chunk_append.X_invoke_Arity2(b, func() interface{} {
															var G__5913 = c.(CljsCoreIIndexed).X_nth_Arity2(i_5925)
															_ = G__5913
															return f.(CljsCoreIFn).X_invoke_Arity1(G__5913)
														}())
														i_5925 = (i_5925 + float64(1))
														continue
													} else {
													}
													break
												}
											}
										}
										return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b), map_.X_invoke_Arity2(f, Chunk_rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq))
									}
								} else {
									return Cons.X_invoke_Arity2(func() interface{} {
										var G__5914 = First.X_invoke_Arity1(s)
										_ = G__5914
										return f.(CljsCoreIFn).X_invoke_Arity1(G__5914)
									}(), map_.X_invoke_Arity2(f, Rest.Arity1IQ(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
								}
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		}, func(f interface{}, c1 interface{}, c2 interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__5926 *AFn) *AFn {
				return Fn(G__5926, 0, func() interface{} {
					{
						var s1 = Seq.Arity1IQ(c1)
						var s2 = Seq.Arity1IQ(c2)
						_, _ = s1, s2
						if Truth_(func() interface{} {
							var and__159__auto__ = s1
							_ = and__159__auto__
							if Truth_(and__159__auto__) {
								return s2
							} else {
								return and__159__auto__
							}
						}()) {
							return Cons.X_invoke_Arity2(func() interface{} {
								var G__5915 = First.X_invoke_Arity1(s1)
								var G__5916 = First.X_invoke_Arity1(s2)
								_, _ = G__5915, G__5916
								return f.(CljsCoreIFn).X_invoke_Arity2(G__5915, G__5916)
							}(), map_.X_invoke_Arity3(f, Rest.Arity1IQ(s1), Rest.Arity1IQ(s2)).(*CljsCoreLazySeq)).(*CljsCoreCons)
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		}, func(f interface{}, c1 interface{}, c2 interface{}, c3 interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__5927 *AFn) *AFn {
				return Fn(G__5927, 0, func() interface{} {
					{
						var s1 = Seq.Arity1IQ(c1)
						var s2 = Seq.Arity1IQ(c2)
						var s3 = Seq.Arity1IQ(c3)
						_, _, _ = s1, s2, s3
						if Truth_(func() interface{} {
							var and__159__auto__ = s1
							_ = and__159__auto__
							if Truth_(and__159__auto__) {
								{
									var and__159__auto_____1 = s2
									_ = and__159__auto_____1
									if Truth_(and__159__auto_____1) {
										return s3
									} else {
										return and__159__auto_____1
									}
								}
							} else {
								return and__159__auto__
							}
						}()) {
							return Cons.X_invoke_Arity2(func() interface{} {
								var G__5917 = First.X_invoke_Arity1(s1)
								var G__5918 = First.X_invoke_Arity1(s2)
								var G__5919 = First.X_invoke_Arity1(s3)
								_, _, _ = G__5917, G__5918, G__5919
								return f.(CljsCoreIFn).X_invoke_Arity3(G__5917, G__5918, G__5919)
							}(), map_.X_invoke_Arity4(f, Rest.Arity1IQ(s1), Rest.Arity1IQ(s2), Rest.Arity1IQ(s3)).(*CljsCoreLazySeq)).(*CljsCoreCons)
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		}, func(f_c1_c2_c3_colls__ ...interface{}) interface{} {
			var f = f_c1_c2_c3_colls__[0]
			var c1 = f_c1_c2_c3_colls__[1]
			var c2 = f_c1_c2_c3_colls__[2]
			var c3 = f_c1_c2_c3_colls__[3]
			var colls = Seq.Arity1IQ(f_c1_c2_c3_colls__[4])
			_, _, _, _, _ = f, c1, c2, c3, colls
			{
				var step = func(step *AFn) *AFn {
					return Fn(step, 1, func(cs interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__5928 *AFn) *AFn {
							return Fn(G__5928, 0, func() interface{} {
								{
									var ss = map_.X_invoke_Arity2(Seq, cs).(*CljsCoreLazySeq)
									_ = ss
									if Every_QMARK_.Arity2IIB(Identity, ss) {
										return Cons.X_invoke_Arity2(map_.X_invoke_Arity2(First, ss).(*CljsCoreLazySeq), step.X_invoke_Arity1(map_.X_invoke_Arity2(Rest, ss).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)).(*CljsCoreCons)
									} else {
										return nil
									}
								}
							})
						}(&AFn{}), nil, nil})
					})
				}(&AFn{})
				_ = step
				return map_.X_invoke_Arity2(func(G__5929 *AFn, step CljsCoreIFn) *AFn {
					return Fn(G__5929, 1, func(p1__5867_SHARP_ interface{}) interface{} {
						return Apply.X_invoke_Arity2(f, p1__5867_SHARP_)
					})
				}(&AFn{}, step), step.X_invoke_Arity1(Conj.X_invoke_ArityVariadic(colls, c3, Array_seq.X_invoke_Arity1([]interface{}{c2, c1}))).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
			}
		})
	}(&AFn{})

	Take = func(take *AFn) *AFn {
		return Fn(take, 2, func(n interface{}) interface{} {
			return func(G__5938 *AFn) *AFn {
				return Fn(G__5938, 1, func(f1 interface{}) interface{} {
					{
						var na = Atom.X_invoke_Arity1(n).(*CljsCoreAtom)
						_ = na
						return func(G__5939 *AFn, na *CljsCoreAtom) *AFn {
							return Fn(G__5939, 2, func() interface{} {
								{
									return f1.(CljsCoreIFn).X_invoke_Arity0()
								}
							}, func(result interface{}) interface{} {
								{
									var G__5935 = result
									_ = G__5935
									return f1.(CljsCoreIFn).X_invoke_Arity1(G__5935)
								}
							}, func(result interface{}, input interface{}) interface{} {
								{
									var n___1 = Deref.X_invoke_Arity1(na)
									var nn = Swap_BANG_.X_invoke_Arity2(na, Dec)
									var result___1 = func() interface{} {
										if n___1.(float64) > float64(0) {
											return func() interface{} {
												var G__5936 = result
												var G__5937 = input
												_, _ = G__5936, G__5937
												return f1.(CljsCoreIFn).X_invoke_Arity2(G__5936, G__5937)
											}()
										} else {
											return result
										}
									}()
									_, _, _ = n___1, nn, result___1
									if !(nn.(float64) > float64(0)) {
										return Reduced.X_invoke_Arity1(result___1).(*CljsCoreReduced)
									} else {
										return result___1
									}
								}
							})
						}(&AFn{}, na)
					}
				})
			}(&AFn{})
		}, func(n interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__5940 *AFn) *AFn {
				return Fn(G__5940, 0, func() interface{} {
					if n.(float64) > float64(0) {
						{
							var temp__4222__auto__ = Seq.Arity1IQ(coll)
							_ = temp__4222__auto__
							if Truth_(temp__4222__auto__) {
								{
									var s = temp__4222__auto__
									_ = s
									return Cons.X_invoke_Arity2(First.X_invoke_Arity1(s), take.X_invoke_Arity2((n.(float64)-float64(1)), Rest.Arity1IQ(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
								}
							} else {
								return nil
							}
						}
					} else {
						return nil
					}
				})
			}(&AFn{}), nil, nil})
		})
	}(&AFn{})

	Drop = func(drop *AFn) *AFn {
		return Fn(drop, 2, func(n interface{}) interface{} {
			return func(G__5947 *AFn) *AFn {
				return Fn(G__5947, 1, func(f1 interface{}) interface{} {
					{
						var na = Atom.X_invoke_Arity1(n).(*CljsCoreAtom)
						_ = na
						return func(G__5948 *AFn, na *CljsCoreAtom) *AFn {
							return Fn(G__5948, 2, func() interface{} {
								{
									return f1.(CljsCoreIFn).X_invoke_Arity0()
								}
							}, func(result interface{}) interface{} {
								{
									var G__5944 = result
									_ = G__5944
									return f1.(CljsCoreIFn).X_invoke_Arity1(G__5944)
								}
							}, func(result interface{}, input interface{}) interface{} {
								{
									var n___1 = Deref.X_invoke_Arity1(na)
									_ = n___1
									Swap_BANG_.X_invoke_Arity2(na, Dec)
									if n___1.(float64) > float64(0) {
										return result
									} else {
										{
											var G__5945 = result
											var G__5946 = input
											_, _ = G__5945, G__5946
											return f1.(CljsCoreIFn).X_invoke_Arity2(G__5945, G__5946)
										}
									}
								}
							})
						}(&AFn{}, na)
					}
				})
			}(&AFn{})
		}, func(n interface{}, coll interface{}) interface{} {
			{
				var step = func(G__5949 *AFn) *AFn {
					return Fn(G__5949, 2, func(n___1 interface{}, coll___1 interface{}) interface{} {
						for {
							{
								var s = Seq.Arity1IQ(coll___1)
								_ = s
								if Truth_(func() interface{} {
									var and__159__auto__ = (n___1.(float64) > float64(0))
									_ = and__159__auto__
									if Truth_(and__159__auto__) {
										return s
									} else {
										return and__159__auto__
									}
								}()) {
									n___1, coll___1 = (n___1.(float64) - float64(1)), Rest.Arity1IQ(s)
									continue
								} else {
									return s
								}
							}
						}
					})
				}(&AFn{})
				_ = step
				return (&CljsCoreLazySeq{nil, func(G__5950 *AFn, step CljsCoreIFn) *AFn {
					return Fn(G__5950, 0, func() interface{} {
						return Seq_(step.X_invoke_Arity2(n, coll))
					})
				}(&AFn{}, step), nil, nil})
			}
		})
	}(&AFn{})

	Drop_last = func(drop_last *AFn) *AFn {
		return Fn(drop_last, 2, func(s interface{}) interface{} {
			return drop_last.X_invoke_Arity2(float64(1), s).(*CljsCoreLazySeq)
		}, func(n interface{}, s interface{}) interface{} {
			return Map_.X_invoke_Arity3(func(G__5953 *AFn) *AFn {
				return Fn(G__5953, 2, func(x interface{}, ___ interface{}) interface{} {
					return x
				})
			}(&AFn{}), s, Drop.X_invoke_Arity2(n, s).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Take_last = func(take_last *AFn) *AFn {
		return Fn(take_last, 2, func(n interface{}, coll interface{}) interface{} {
			{
				var s interface{} = Seq.Arity1IQ(coll)
				var lead interface{} = Seq.Arity1IQ(Drop.X_invoke_Arity2(n, coll).(*CljsCoreLazySeq))
				_, _ = s, lead
				for {
					if Truth_(lead) {
						s, lead = Next.Arity1IQ(s), Next.Arity1IQ(lead)
						continue
					} else {
						return s
					}
				}
			}
		})
	}(&AFn{})

	Drop_while = func(drop_while *AFn) *AFn {
		return Fn(drop_while, 2, func(pred interface{}) interface{} {
			return func(G__5964 *AFn) *AFn {
				return Fn(G__5964, 1, func(f1 interface{}) interface{} {
					{
						var da = Atom.X_invoke_Arity1(true).(*CljsCoreAtom)
						_ = da
						return func(G__5965 *AFn, da *CljsCoreAtom) *AFn {
							return Fn(G__5965, 2, func() interface{} {
								{
									return f1.(CljsCoreIFn).X_invoke_Arity0()
								}
							}, func(result interface{}) interface{} {
								{
									var G__5959 = result
									_ = G__5959
									return f1.(CljsCoreIFn).X_invoke_Arity1(G__5959)
								}
							}, func(result interface{}, input interface{}) interface{} {
								{
									var drop_QMARK_ = Deref.X_invoke_Arity1(da)
									_ = drop_QMARK_
									if Truth_(func() interface{} {
										var and__159__auto__ = drop_QMARK_
										_ = and__159__auto__
										if Truth_(and__159__auto__) {
											{
												var G__5960 = input
												_ = G__5960
												return pred.(CljsCoreIFn).X_invoke_Arity1(G__5960)
											}
										} else {
											return and__159__auto__
										}
									}()) {
										return result
									} else {
										Reset_BANG_.X_invoke_Arity2(da, nil)
										{
											var G__5961 = result
											var G__5962 = input
											_, _ = G__5961, G__5962
											return f1.(CljsCoreIFn).X_invoke_Arity2(G__5961, G__5962)
										}
									}
								}
							})
						}(&AFn{}, da)
					}
				})
			}(&AFn{})
		}, func(pred interface{}, coll interface{}) interface{} {
			{
				var step = func(G__5966 *AFn) *AFn {
					return Fn(G__5966, 2, func(pred___1 interface{}, coll___1 interface{}) interface{} {
						for {
							{
								var s = Seq.Arity1IQ(coll___1)
								_ = s
								if Truth_(func() interface{} {
									var and__159__auto__ = s
									_ = and__159__auto__
									if Truth_(and__159__auto__) {
										{
											var G__5963 = First.X_invoke_Arity1(s)
											_ = G__5963
											return pred___1.(CljsCoreIFn).X_invoke_Arity1(G__5963)
										}
									} else {
										return and__159__auto__
									}
								}()) {
									pred___1, coll___1 = pred___1, Rest.Arity1IQ(s)
									continue
								} else {
									return s
								}
							}
						}
					})
				}(&AFn{})
				_ = step
				return (&CljsCoreLazySeq{nil, func(G__5967 *AFn, step CljsCoreIFn) *AFn {
					return Fn(G__5967, 0, func() interface{} {
						return Seq_(step.X_invoke_Arity2(pred, coll))
					})
				}(&AFn{}, step), nil, nil})
			}
		})
	}(&AFn{})

	Cycle = func(cycle *AFn) *AFn {
		return Fn(cycle, 1, func(coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__5969 *AFn) *AFn {
				return Fn(G__5969, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								return Concat.X_invoke_Arity2(s, cycle.X_invoke_Arity1(s).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		})
	}(&AFn{})

	Split_at = func(split_at *AFn) *AFn {
		return Fn(split_at, 2, func(n interface{}, coll interface{}) interface{} {
			return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Take.X_invoke_Arity2(n, coll).(*CljsCoreLazySeq), Drop.X_invoke_Arity2(n, coll).(*CljsCoreLazySeq)}, nil})
		})
	}(&AFn{})

	Repeat = func(repeat *AFn) *AFn {
		return Fn(repeat, 2, func(x interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__5972 *AFn) *AFn {
				return Fn(G__5972, 0, func() interface{} {
					return Cons.X_invoke_Arity2(x, repeat.X_invoke_Arity1(x).(*CljsCoreLazySeq)).(*CljsCoreCons)
				})
			}(&AFn{}), nil, nil})
		}, func(n interface{}, x interface{}) interface{} {
			return Take.X_invoke_Arity2(n, repeat.X_invoke_Arity1(x).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Replicate = func(replicate *AFn) *AFn {
		return Fn(replicate, 2, func(n interface{}, x interface{}) interface{} {
			return Take.X_invoke_Arity2(n, Repeat.X_invoke_Arity1(x).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Repeatedly = func(repeatedly *AFn) *AFn {
		return Fn(repeatedly, 2, func(f interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__5975 *AFn) *AFn {
				return Fn(G__5975, 0, func() interface{} {
					return Cons.X_invoke_Arity2(func() interface{} {
						return f.(CljsCoreIFn).X_invoke_Arity0()
					}(), repeatedly.X_invoke_Arity1(f).(*CljsCoreLazySeq)).(*CljsCoreCons)
				})
			}(&AFn{}), nil, nil})
		}, func(n interface{}, f interface{}) interface{} {
			return Take.X_invoke_Arity2(n, repeatedly.X_invoke_Arity1(f).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Iterate = func(iterate *AFn) *AFn {
		return Fn(iterate, 2, func(f interface{}, x interface{}) interface{} {
			return Cons.X_invoke_Arity2(x, (&CljsCoreLazySeq{nil, func(G__5980 *AFn) *AFn {
				return Fn(G__5980, 0, func() interface{} {
					return iterate.X_invoke_Arity2(f, func() interface{} {
						var G__5979 = x
						_ = G__5979
						return f.(CljsCoreIFn).X_invoke_Arity1(G__5979)
					}()).(*CljsCoreCons)
				})
			}(&AFn{}), nil, nil})).(*CljsCoreCons)
		})
	}(&AFn{})

	Interleave = func(interleave *AFn) *AFn {
		return Fn(interleave, 2, func(c1 interface{}, c2 interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__5983 *AFn) *AFn {
				return Fn(G__5983, 0, func() interface{} {
					{
						var s1 = Seq.Arity1IQ(c1)
						var s2 = Seq.Arity1IQ(c2)
						_, _ = s1, s2
						if Truth_(func() interface{} {
							var and__159__auto__ = s1
							_ = and__159__auto__
							if Truth_(and__159__auto__) {
								return s2
							} else {
								return and__159__auto__
							}
						}()) {
							return Cons.X_invoke_Arity2(First.X_invoke_Arity1(s1), Cons.X_invoke_Arity2(First.X_invoke_Arity1(s2), interleave.X_invoke_Arity2(Rest.Arity1IQ(s1), Rest.Arity1IQ(s2)).(*CljsCoreLazySeq)).(*CljsCoreCons)).(*CljsCoreCons)
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		}, func(c1_c2_colls__ ...interface{}) interface{} {
			var c1 = c1_c2_colls__[0]
			var c2 = c1_c2_colls__[1]
			var colls = Seq.Arity1IQ(c1_c2_colls__[2])
			_, _, _ = c1, c2, colls
			return (&CljsCoreLazySeq{nil, func(G__5984 *AFn) *AFn {
				return Fn(G__5984, 0, func() interface{} {
					{
						var ss = Map_.X_invoke_Arity2(Seq, Conj.X_invoke_ArityVariadic(colls, c2, Array_seq.X_invoke_Arity1([]interface{}{c1}))).(*CljsCoreLazySeq)
						_ = ss
						if Every_QMARK_.Arity2IIB(Identity, ss) {
							return Concat.X_invoke_Arity2(Map_.X_invoke_Arity2(First, ss).(*CljsCoreLazySeq), Apply.X_invoke_Arity2(interleave, Map_.X_invoke_Arity2(Rest, ss).(*CljsCoreLazySeq))).(*CljsCoreLazySeq)
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		})
	}(&AFn{})

	Interpose = func(interpose *AFn) *AFn {
		return Fn(interpose, 2, func(sep interface{}, coll interface{}) interface{} {
			return Drop.X_invoke_Arity2(float64(1), Interleave.X_invoke_Arity2(Repeat.X_invoke_Arity1(sep).(*CljsCoreLazySeq), coll).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Flatten1 = func(flatten1 *AFn) *AFn {
		return Fn(flatten1, 1, func(colls interface{}) interface{} {
			{
				var cat = func(cat *AFn) *AFn {
					return Fn(cat, 2, func(coll interface{}, colls___1 interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__5993 *AFn) *AFn {
							return Fn(G__5993, 0, func() interface{} {
								{
									var temp__4220__auto__ = Seq.Arity1IQ(coll)
									_ = temp__4220__auto__
									if Truth_(temp__4220__auto__) {
										{
											var coll___1 = temp__4220__auto__
											_ = coll___1
											return Cons.X_invoke_Arity2(First.X_invoke_Arity1(coll___1), cat.X_invoke_Arity2(Rest.Arity1IQ(coll___1), colls___1).(*CljsCoreLazySeq)).(*CljsCoreCons)
										}
									} else {
										if Truth_(Seq.Arity1IQ(colls___1)) {
											return cat.X_invoke_Arity2(First.X_invoke_Arity1(colls___1), Rest.Arity1IQ(colls___1)).(*CljsCoreLazySeq)
										} else {
											return nil
										}
									}
								}
							})
						}(&AFn{}), nil, nil})
					})
				}(&AFn{})
				_ = cat
				return cat.X_invoke_Arity2(nil, colls).(*CljsCoreLazySeq)
			}
		})
	}(&AFn{})

	Mapcat = func(mapcat *AFn) *AFn {
		return Fn(mapcat, 1, func(f interface{}) interface{} {
			return Comp.X_invoke_Arity2(Map_.X_invoke_Arity1(f).(CljsCoreIFn), Cat).(CljsCoreIFn)
		}, func(f_colls__ ...interface{}) interface{} {
			var f = f_colls__[0]
			var colls = Seq.Arity1IQ(f_colls__[1])
			_, _ = f, colls
			return Apply.X_invoke_Arity2(Concat, Apply.X_invoke_Arity3(Map_, f, colls))
		})
	}(&AFn{})

	Filter = func(filter *AFn) *AFn {
		return Fn(filter, 2, func(pred interface{}) interface{} {
			return func(G__6012 *AFn) *AFn {
				return Fn(G__6012, 1, func(f1 interface{}) interface{} {
					return func(G__6013 *AFn) *AFn {
						return Fn(G__6013, 2, func() interface{} {
							{
								return f1.(CljsCoreIFn).X_invoke_Arity0()
							}
						}, func(result interface{}) interface{} {
							{
								var G__6006 = result
								_ = G__6006
								return f1.(CljsCoreIFn).X_invoke_Arity1(G__6006)
							}
						}, func(result interface{}, input interface{}) interface{} {
							if Truth_(func() interface{} {
								var G__6007 = input
								_ = G__6007
								return pred.(CljsCoreIFn).X_invoke_Arity1(G__6007)
							}()) {
								{
									var G__6008 = result
									var G__6009 = input
									_, _ = G__6008, G__6009
									return f1.(CljsCoreIFn).X_invoke_Arity2(G__6008, G__6009)
								}
							} else {
								return result
							}
						})
					}(&AFn{})
				})
			}(&AFn{})
		}, func(pred interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__6014 *AFn) *AFn {
				return Fn(G__6014, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								if Chunked_seq_QMARK_.Arity1IB(s) {
									{
										var c = Chunk_first.X_invoke_Arity1(s)
										var size = Count.X_invoke_Arity1(c).(float64)
										var b = Chunk_buffer.X_invoke_Arity1(size).(*CljsCoreChunkBuffer)
										_, _, _ = c, size, b
										{
											var n__1054__auto___6015 = size
											_ = n__1054__auto___6015
											{
												var i_6016 = float64(0)
												_ = i_6016
												for {
													if i_6016 < n__1054__auto___6015 {
														if Truth_(func() interface{} {
															var G__6010 = c.(CljsCoreIIndexed).X_nth_Arity2(i_6016)
															_ = G__6010
															return pred.(CljsCoreIFn).X_invoke_Arity1(G__6010)
														}()) {
															Chunk_append.X_invoke_Arity2(b, c.(CljsCoreIIndexed).X_nth_Arity2(i_6016))
														} else {
														}
														i_6016 = (i_6016 + float64(1))
														continue
													} else {
													}
													break
												}
											}
										}
										return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b), filter.X_invoke_Arity2(pred, Chunk_rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq))
									}
								} else {
									{
										var f = First.X_invoke_Arity1(s)
										var r = Rest.Arity1IQ(s)
										_, _ = f, r
										if Truth_(func() interface{} {
											var G__6011 = f
											_ = G__6011
											return pred.(CljsCoreIFn).X_invoke_Arity1(G__6011)
										}()) {
											return Cons.X_invoke_Arity2(f, filter.X_invoke_Arity2(pred, r).(*CljsCoreLazySeq)).(*CljsCoreCons)
										} else {
											return filter.X_invoke_Arity2(pred, r).(*CljsCoreLazySeq)
										}
									}
								}
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		})
	}(&AFn{})

	Tree_seq = func(tree_seq *AFn) *AFn {
		return Fn(tree_seq, 3, func(branch_QMARK_ interface{}, children interface{}, root interface{}) interface{} {
			{
				var walk = func(walk *AFn) *AFn {
					return Fn(walk, 1, func(node interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__6025 *AFn) *AFn {
							return Fn(G__6025, 0, func() interface{} {
								return Cons.X_invoke_Arity2(node, func() interface{} {
									if Truth_(func() interface{} {
										var G__6023 = node
										_ = G__6023
										return branch_QMARK_.(CljsCoreIFn).X_invoke_Arity1(G__6023)
									}()) {
										return Mapcat.X_invoke_ArityVariadic(walk, Array_seq.X_invoke_Arity1([]interface{}{func() interface{} {
											var G__6024 = node
											_ = G__6024
											return children.(CljsCoreIFn).X_invoke_Arity1(G__6024)
										}()}))
									} else {
										return nil
									}
								}()).(*CljsCoreCons)
							})
						}(&AFn{}), nil, nil})
					})
				}(&AFn{})
				_ = walk
				return walk.X_invoke_Arity1(root).(*CljsCoreLazySeq)
			}
		})
	}(&AFn{})

	Flatten = func(flatten *AFn) *AFn {
		return Fn(flatten, 1, func(x interface{}) interface{} {
			return Filter.X_invoke_Arity2(func(G__6027 *AFn) *AFn {
				return Fn(G__6027, 1, func(p1__6026_SHARP_ interface{}) interface{} {
					return !(Sequential_QMARK_.Arity1IB(p1__6026_SHARP_))
				})
			}(&AFn{}), Rest.Arity1IQ(Tree_seq.X_invoke_Arity3(Sequential_QMARK_, Seq, x).(*CljsCoreLazySeq))).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Into = func(into *AFn) *AFn {
		return Fn(into, 3, func(to interface{}, from interface{}) interface{} {
			if !(Nil_(to)) {
				if Value_(to).Type().Implements(reflect.TypeOf((*CljsCoreIEditableCollection)(nil)).Elem()) {
					return With_meta.X_invoke_Arity2(Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(X_conj_BANG_, Transient.X_invoke_Arity1(to), from)), Meta.X_invoke_Arity1(to))
				} else {
					return Reduce.X_invoke_Arity3(X_conj, to, from)
				}
			} else {
				return Reduce.X_invoke_Arity3(Conj, CljsCoreIEmptyList(CljsCoreList_EMPTY), from)
			}
		}, func(to interface{}, xform interface{}, from interface{}) interface{} {
			if Value_(to).Type().Implements(reflect.TypeOf((*CljsCoreIEditableCollection)(nil)).Elem()) {
				return With_meta.X_invoke_Arity2(Persistent_BANG_.X_invoke_Arity1(Transduce.X_invoke_Arity4(xform, Conj_BANG_, Transient.X_invoke_Arity1(to), from)), Meta.X_invoke_Arity1(to))
			} else {
				return Transduce.X_invoke_Arity4(xform, Conj, to, from)
			}
		})
	}(&AFn{})

	Mapv = func(mapv *AFn) *AFn {
		return Fn(mapv, 4, func(f interface{}, coll interface{}) interface{} {
			return Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(func(G__6030 *AFn) *AFn {
				return Fn(G__6030, 2, func(v interface{}, o interface{}) interface{} {
					return Conj_BANG_.X_invoke_Arity2(v, func() interface{} {
						var G__6029 = o
						_ = G__6029
						return f.(CljsCoreIFn).X_invoke_Arity1(G__6029)
					}())
				})
			}(&AFn{}), Transient.X_invoke_Arity1(CljsCorePersistentVector_EMPTY), coll))
		}, func(f interface{}, c1 interface{}, c2 interface{}) interface{} {
			return Into.X_invoke_Arity2(CljsCorePersistentVector_EMPTY, Map_.X_invoke_Arity3(f, c1, c2).(*CljsCoreLazySeq))
		}, func(f interface{}, c1 interface{}, c2 interface{}, c3 interface{}) interface{} {
			return Into.X_invoke_Arity2(CljsCorePersistentVector_EMPTY, Map_.X_invoke_Arity4(f, c1, c2, c3).(*CljsCoreLazySeq))
		}, func(f_c1_c2_c3_colls__ ...interface{}) interface{} {
			var f = f_c1_c2_c3_colls__[0]
			var c1 = f_c1_c2_c3_colls__[1]
			var c2 = f_c1_c2_c3_colls__[2]
			var c3 = f_c1_c2_c3_colls__[3]
			var colls = Seq.Arity1IQ(f_c1_c2_c3_colls__[4])
			_, _, _, _, _ = f, c1, c2, c3, colls
			return Into.X_invoke_Arity2(CljsCorePersistentVector_EMPTY, Apply.X_invoke_ArityVariadic(Map_, f, c1, c2, c3, Array_seq.X_invoke_Arity1([]interface{}{colls})))
		})
	}(&AFn{})

	Filterv = func(filterv *AFn) *AFn {
		return Fn(filterv, 2, func(pred interface{}, coll interface{}) interface{} {
			return Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(func(G__6033 *AFn) *AFn {
				return Fn(G__6033, 2, func(v interface{}, o interface{}) interface{} {
					if Truth_(func() interface{} {
						var G__6032 = o
						_ = G__6032
						return pred.(CljsCoreIFn).X_invoke_Arity1(G__6032)
					}()) {
						return Conj_BANG_.X_invoke_Arity2(v, o)
					} else {
						return v
					}
				})
			}(&AFn{}), Transient.X_invoke_Arity1(CljsCorePersistentVector_EMPTY), coll))
		})
	}(&AFn{})

	Partition = func(partition *AFn) *AFn {
		return Fn(partition, 4, func(n interface{}, coll interface{}) interface{} {
			return partition.X_invoke_Arity3(n, n, coll).(*CljsCoreLazySeq)
		}, func(n interface{}, step interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__6044 *AFn) *AFn {
				return Fn(G__6044, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								{
									var p = Take.X_invoke_Arity2(n, s).(*CljsCoreLazySeq)
									_ = p
									if n.(float64) == Count.X_invoke_Arity1(p).(float64) {
										return Cons.X_invoke_Arity2(p, partition.X_invoke_Arity3(n, step, Drop.X_invoke_Arity2(step, s).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)).(*CljsCoreCons)
									} else {
										return nil
									}
								}
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		}, func(n interface{}, step interface{}, pad interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__6045 *AFn) *AFn {
				return Fn(G__6045, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								{
									var p = Take.X_invoke_Arity2(n, s).(*CljsCoreLazySeq)
									_ = p
									if n.(float64) == Count.X_invoke_Arity1(p).(float64) {
										return Cons.X_invoke_Arity2(p, partition.X_invoke_Arity4(n, step, pad, Drop.X_invoke_Arity2(step, s).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)).(*CljsCoreCons)
									} else {
										return CljsCoreList_EMPTY.X_conj_Arity2(Take.X_invoke_Arity2(n, Concat.X_invoke_Arity2(p, pad).(*CljsCoreLazySeq)).(*CljsCoreLazySeq))
									}
								}
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		})
	}(&AFn{})

	Get_in = func(get_in *AFn) *AFn {
		return Fn(get_in, 3, func(m interface{}, ks interface{}) interface{} {
			return get_in.X_invoke_Arity3(m, ks, nil)
		}, func(m interface{}, ks interface{}, not_found interface{}) interface{} {
			{
				var sentinel interface{} = Lookup_sentinel
				var m___1 interface{} = m
				var ks___1 interface{} = Seq.Arity1IQ(ks)
				_, _, _ = sentinel, m___1, ks___1
				for {
					if Truth_(ks___1) {
						if !(Value_(m___1).Type().Implements(reflect.TypeOf((*CljsCoreILookup)(nil)).Elem())) {
							return not_found
						} else {
							{
								var m___2 = Get.X_invoke_Arity3(m___1, First.X_invoke_Arity1(ks___1), sentinel)
								_ = m___2
								if reflect.DeepEqual(sentinel, m___2) {
									return not_found
								} else {
									sentinel, m___1, ks___1 = sentinel, m___2, Next.Arity1IQ(ks___1)
									continue
								}
							}
						}
					} else {
						return m___1
					}
				}
			}
		})
	}(&AFn{})

	Assoc_in = func(assoc_in *AFn) *AFn {
		return Fn(assoc_in, 3, func(m interface{}, p__6049 interface{}, v interface{}) interface{} {
			{
				var vec__6054 = p__6049
				var k = Nth.X_invoke_Arity3(vec__6054, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__6054, float64(1)))
				_, _, _ = vec__6054, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, assoc_in.X_invoke_Arity3(Get.X_invoke_Arity2(m, k), ks, v))
				} else {
					return Assoc.X_invoke_Arity3(m, k, v)
				}
			}
		})
	}(&AFn{})

	Update_in = func(update_in *AFn) *AFn {
		return Fn(update_in, 6, func(m interface{}, p__6055 interface{}, f interface{}) interface{} {
			{
				var vec__6093 = p__6055
				var k = Nth.X_invoke_Arity3(vec__6093, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__6093, float64(1)))
				_, _, _ = vec__6093, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity3(Get.X_invoke_Arity2(m, k), ks, f))
				} else {
					return Assoc.X_invoke_Arity3(m, k, func() interface{} {
						var G__6094 = Get.X_invoke_Arity2(m, k)
						_ = G__6094
						return f.(CljsCoreIFn).X_invoke_Arity1(G__6094)
					}())
				}
			}
		}, func(m interface{}, p__6056 interface{}, f interface{}, a interface{}) interface{} {
			{
				var vec__6095 = p__6056
				var k = Nth.X_invoke_Arity3(vec__6095, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__6095, float64(1)))
				_, _, _ = vec__6095, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity4(Get.X_invoke_Arity2(m, k), ks, f, a))
				} else {
					return Assoc.X_invoke_Arity3(m, k, func() interface{} {
						var G__6096 = Get.X_invoke_Arity2(m, k)
						var G__6097 = a
						_, _ = G__6096, G__6097
						return f.(CljsCoreIFn).X_invoke_Arity2(G__6096, G__6097)
					}())
				}
			}
		}, func(m interface{}, p__6057 interface{}, f interface{}, a interface{}, b interface{}) interface{} {
			{
				var vec__6098 = p__6057
				var k = Nth.X_invoke_Arity3(vec__6098, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__6098, float64(1)))
				_, _, _ = vec__6098, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity5(Get.X_invoke_Arity2(m, k), ks, f, a, b))
				} else {
					return Assoc.X_invoke_Arity3(m, k, func() interface{} {
						var G__6099 = Get.X_invoke_Arity2(m, k)
						var G__6100 = a
						var G__6101 = b
						_, _, _ = G__6099, G__6100, G__6101
						return f.(CljsCoreIFn).X_invoke_Arity3(G__6099, G__6100, G__6101)
					}())
				}
			}
		}, func(m interface{}, p__6058 interface{}, f interface{}, a interface{}, b interface{}, c interface{}) interface{} {
			{
				var vec__6102 = p__6058
				var k = Nth.X_invoke_Arity3(vec__6102, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__6102, float64(1)))
				_, _, _ = vec__6102, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity6(Get.X_invoke_Arity2(m, k), ks, f, a, b, c))
				} else {
					return Assoc.X_invoke_Arity3(m, k, func() interface{} {
						var G__6103 = Get.X_invoke_Arity2(m, k)
						var G__6104 = a
						var G__6105 = b
						var G__6106 = c
						_, _, _, _ = G__6103, G__6104, G__6105, G__6106
						return f.(CljsCoreIFn).X_invoke_Arity4(G__6103, G__6104, G__6105, G__6106)
					}())
				}
			}
		}, func(m_p__6059_f_a_b_c_args__ ...interface{}) interface{} {
			var m = m_p__6059_f_a_b_c_args__[0]
			var p__6059 = m_p__6059_f_a_b_c_args__[1]
			var f = m_p__6059_f_a_b_c_args__[2]
			var a = m_p__6059_f_a_b_c_args__[3]
			var b = m_p__6059_f_a_b_c_args__[4]
			var c = m_p__6059_f_a_b_c_args__[5]
			var args = Seq.Arity1IQ(m_p__6059_f_a_b_c_args__[6])
			_, _, _, _, _, _, _ = m, p__6059, f, a, b, c, args
			{
				var vec__6107 = p__6059
				var k = Nth.X_invoke_Arity3(vec__6107, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__6107, float64(1)))
				_, _, _ = vec__6107, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, Apply.X_invoke_ArityVariadic(update_in, Get.X_invoke_Arity2(m, k), ks, f, a, Array_seq.X_invoke_Arity1([]interface{}{b, c, args})))
				} else {
					return Assoc.X_invoke_Arity3(m, k, Apply.X_invoke_ArityVariadic(f, Get.X_invoke_Arity2(m, k), a, b, c, Array_seq.X_invoke_Arity1([]interface{}{args})))
				}
			}
		})
	}(&AFn{})

	X__GT_VectorNode = func(__GT_VectorNode *AFn) *AFn {
		return Fn(__GT_VectorNode, 2, func(edit interface{}, arr interface{}) interface{} {
			return (&CljsCoreVectorNode{edit, arr})
		})
	}(&AFn{})

	Pv_fresh_node = func(pv_fresh_node *AFn) *AFn {
		return Fn(pv_fresh_node, 1, func(edit interface{}) interface{} {
			return (&CljsCoreVectorNode{edit, make([]interface{}, int(float64(32)))})
		})
	}(&AFn{})

	Pv_aget = func(pv_aget *AFn) *AFn {
		return Fn(pv_aget, 2, func(node interface{}, idx interface{}) interface{} {
			return Aget_(Native_get_instance_field.X_invoke_Arity2(node, "Arr"), idx.(float64))
		})
	}(&AFn{})

	Pv_aset = func(pv_aset *AFn) *AFn {
		return Fn(pv_aset, 3, func(node interface{}, idx interface{}, val interface{}) interface{} {
			return func() interface{} {
				Native_get_instance_field.X_invoke_Arity2(node, "Arr").([]interface{})[int(idx.(float64))] = val
				return Native_get_instance_field.X_invoke_Arity2(node, "Arr").([]interface{})[int(idx.(float64))]
			}()
		})
	}(&AFn{})

	Pv_clone_node = func(pv_clone_node *AFn) *AFn {
		return Fn(pv_clone_node, 1, func(node interface{}) interface{} {
			return (&CljsCoreVectorNode{Native_get_instance_field.X_invoke_Arity2(node, "Edit"), Aclone.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(node, "Arr")).([]interface{})})
		})
	}(&AFn{})

	Tail_off = func(tail_off *AFn) *AFn {
		return Fn(tail_off, 1, func(pv interface{}) interface{} {
			{
				var cnt = Native_get_instance_field.X_invoke_Arity2(pv, "Cnt")
				_ = cnt
				if cnt.(float64) < float64(32) {
					return float64(0)
				} else {
					return float64((Int32_(float64((UInt32_((cnt.(float64) - float64(1))) >> UInt32_(float64((32+Int32_(float64(5)))%32))))) << UInt32_(float64(5))))
				}
			}
		})
	}(&AFn{})

	New_path = func(new_path *AFn) *AFn {
		return Fn(new_path, 3, func(edit interface{}, level interface{}, node interface{}) interface{} {
			{
				var ll interface{} = level
				var ret interface{} = node
				_, _ = ll, ret
				for {
					if ll.(float64) == float64(0) {
						return ret
					} else {
						{
							var embed = ret
							var r = Pv_fresh_node.X_invoke_Arity1(edit).(*CljsCoreVectorNode)
							var ___ = Pv_aset.X_invoke_Arity3(r, float64(0), embed)
							_, _, _ = embed, r, ___
							ll, ret = (ll.(float64) - float64(5)), r
							continue
						}
					}
				}
			}
		})
	}(&AFn{})

	Push_tail = func(push_tail *AFn) *AFn {
		return Fn(push_tail, 4, func(pv interface{}, level interface{}, parent interface{}, tailnode interface{}) interface{} {
			{
				var ret = Pv_clone_node.X_invoke_Arity1(parent).(*CljsCoreVectorNode)
				var subidx = float64((Int32_(float64((UInt32_((Native_get_instance_field.X_invoke_Arity2(pv, "Cnt").(float64) - float64(1))) >> UInt32_(float64((32+Int32_(level.(float64)))%32))))) & Int32_(float64(31))))
				_, _ = ret, subidx
				if float64(5) == level.(float64) {
					Pv_aset.X_invoke_Arity3(ret, subidx, tailnode)
					return ret
				} else {
					{
						var child = Pv_aget.X_invoke_Arity2(parent, subidx)
						_ = child
						if !(Nil_(child)) {
							{
								var node_to_insert = push_tail.X_invoke_Arity4(pv, (level.(float64) - float64(5)), child, tailnode).(*CljsCoreVectorNode)
								_ = node_to_insert
								Pv_aset.X_invoke_Arity3(ret, subidx, node_to_insert)
								return ret
							}
						} else {
							{
								var node_to_insert = New_path.X_invoke_Arity3(nil, (level.(float64) - float64(5)), tailnode)
								_ = node_to_insert
								Pv_aset.X_invoke_Arity3(ret, subidx, node_to_insert)
								return ret
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Vector_index_out_of_bounds = func(vector_index_out_of_bounds *AFn) *AFn {
		return Fn(vector_index_out_of_bounds, 2, func(i interface{}, cnt interface{}) interface{} {
			panic((&js.Error{("No item " + Str.X_invoke_Arity1(i).(string) + " in vector of length " + Str.X_invoke_Arity1(cnt).(string))}))
		})
	}(&AFn{})

	First_array_for_longvec = func(first_array_for_longvec *AFn) *AFn {
		return Fn(first_array_for_longvec, 1, func(pv interface{}) interface{} {
			{
				var node interface{} = Native_get_instance_field.X_invoke_Arity2(pv, "Root")
				var level interface{} = Native_get_instance_field.X_invoke_Arity2(pv, "Shift")
				_, _ = node, level
				for {
					if level.(float64) > float64(0) {
						node, level = Pv_aget.X_invoke_Arity2(node, float64(0)), (level.(float64) - float64(5))
						continue
					} else {
						return Native_get_instance_field.X_invoke_Arity2(node, "Arr")
					}
				}
			}
		})
	}(&AFn{})

	Unchecked_array_for = func(unchecked_array_for *AFn) *AFn {
		return Fn(unchecked_array_for, 2, func(pv interface{}, i interface{}) interface{} {
			if i.(float64) >= Tail_off.X_invoke_Arity1(pv).(float64) {
				return Native_get_instance_field.X_invoke_Arity2(pv, "Tail")
			} else {
				{
					var node interface{} = Native_get_instance_field.X_invoke_Arity2(pv, "Root")
					var level interface{} = Native_get_instance_field.X_invoke_Arity2(pv, "Shift")
					_, _ = node, level
					for {
						if level.(float64) > float64(0) {
							node, level = Pv_aget.X_invoke_Arity2(node, float64((Int32_(float64((UInt32_(i.(float64))>>UInt32_(float64((32+Int32_(level.(float64)))%32)))))&Int32_(float64(31))))), (level.(float64) - float64(5))
							continue
						} else {
							return Native_get_instance_field.X_invoke_Arity2(node, "Arr")
						}
					}
				}
			}
		})
	}(&AFn{})

	Array_for = func(array_for *AFn) *AFn {
		return Fn(array_for, 2, func(pv interface{}, i interface{}) interface{} {
			if (float64(0) <= i.(float64)) && (i.(float64) < Native_get_instance_field.X_invoke_Arity2(pv, "Cnt").(float64)) {
				return Unchecked_array_for.X_invoke_Arity2(pv, i)
			} else {
				return Vector_index_out_of_bounds.X_invoke_Arity2(i, Native_get_instance_field.X_invoke_Arity2(pv, "Cnt"))
			}
		})
	}(&AFn{})

	Do_assoc = func(do_assoc *AFn) *AFn {
		return Fn(do_assoc, 5, func(pv interface{}, level interface{}, node interface{}, i interface{}, val interface{}) interface{} {
			{
				var ret = Pv_clone_node.X_invoke_Arity1(node).(*CljsCoreVectorNode)
				_ = ret
				if level.(float64) == float64(0) {
					Pv_aset.X_invoke_Arity3(ret, float64((Int32_(i.(float64)) & Int32_(float64(31)))), val)
					return ret
				} else {
					{
						var subidx = float64((Int32_(float64((UInt32_(i.(float64)) >> UInt32_(float64((32+Int32_(level.(float64)))%32))))) & Int32_(float64(31))))
						_ = subidx
						Pv_aset.X_invoke_Arity3(ret, subidx, do_assoc.X_invoke_Arity5(pv, (level.(float64)-float64(5)), Pv_aget.X_invoke_Arity2(node, subidx), i, val).(*CljsCoreVectorNode))
						return ret
					}
				}
			}
		})
	}(&AFn{})

	Pop_tail = func(pop_tail *AFn) *AFn {
		return Fn(pop_tail, 3, func(pv interface{}, level interface{}, node interface{}) interface{} {
			{
				var subidx = float64((Int32_(float64((UInt32_((Native_get_instance_field.X_invoke_Arity2(pv, "Cnt").(float64) - float64(2))) >> UInt32_(float64((32+Int32_(level.(float64)))%32))))) & Int32_(float64(31))))
				_ = subidx
				if level.(float64) > float64(5) {
					{
						var new_child = pop_tail.X_invoke_Arity3(pv, (level.(float64) - float64(5)), Pv_aget.X_invoke_Arity2(node, subidx))
						_ = new_child
						if (Nil_(new_child)) && (subidx == float64(0)) {
							return nil
						} else {
							{
								var ret = Pv_clone_node.X_invoke_Arity1(node).(*CljsCoreVectorNode)
								_ = ret
								Pv_aset.X_invoke_Arity3(ret, subidx, new_child)
								return ret
							}
						}
					}
				} else {
					if subidx == float64(0) {
						return nil
					} else {
						{
							var ret = Pv_clone_node.X_invoke_Arity1(node).(*CljsCoreVectorNode)
							_ = ret
							Pv_aset.X_invoke_Arity3(ret, subidx, nil)
							return ret
						}

					}
				}
			}
		})
	}(&AFn{})

	X__GT_PersistentVector = func(__GT_PersistentVector *AFn) *AFn {
		return Fn(__GT_PersistentVector, 6, func(meta interface{}, cnt interface{}, shift interface{}, root interface{}, tail interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentVector{meta, cnt, shift, root, tail, __hash})
		})
	}(&AFn{})

	Vec = func(vec *AFn) *AFn {
		return Fn(vec, 1, func(coll interface{}) interface{} {
			return Reduce.X_invoke_Arity3(X_conj_BANG_, CljsCorePersistentVector_EMPTY.X_as_transient_Arity1(), coll).(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
		})
	}(&AFn{})

	Vector = func(vector *AFn) *AFn {
		return Fn(vector, 0, func(args__ ...interface{}) interface{} {
			var args = Seq.Arity1IQ(args__[0])
			_ = args
			if (Value_(args).Type().AssignableTo(reflect.TypeOf((**CljsCoreIndexedSeq)(nil)).Elem())) && (Native_get_instance_field.X_invoke_Arity2(args, "I").(float64) == float64(0)) {
				return CljsCorePersistentVector_FromArray.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(args, "Arr"), true)
			} else {
				return Vec.X_invoke_Arity1(args)
			}
		})
	}(&AFn{})

	X__GT_ChunkedSeq = func(__GT_ChunkedSeq *AFn) *AFn {
		return Fn(__GT_ChunkedSeq, 6, func(vec interface{}, node interface{}, i interface{}, off interface{}, meta interface{}, __hash interface{}) interface{} {
			return (&CljsCoreChunkedSeq{vec, node, i, off, meta, __hash})
		})
	}(&AFn{})

	Chunked_seq = func(chunked_seq *AFn) *AFn {
		return Fn(chunked_seq, 5, func(vec interface{}, i interface{}, off interface{}) interface{} {
			return (&CljsCoreChunkedSeq{vec, Array_for.X_invoke_Arity2(vec, i), i, off, nil, nil})
		}, func(vec interface{}, node interface{}, i interface{}, off interface{}) interface{} {
			return (&CljsCoreChunkedSeq{vec, node, i, off, nil, nil})
		}, func(vec interface{}, node interface{}, i interface{}, off interface{}, meta interface{}) interface{} {
			return (&CljsCoreChunkedSeq{vec, node, i, off, meta, nil})
		})
	}(&AFn{})

	X__GT_Subvec = func(__GT_Subvec *AFn) *AFn {
		return Fn(__GT_Subvec, 5, func(meta interface{}, v interface{}, start interface{}, end interface{}, __hash interface{}) interface{} {
			return (&CljsCoreSubvec{meta, v, start, end, __hash})
		})
	}(&AFn{})

	Build_subvec = func(build_subvec *AFn) *AFn {
		return Fn(build_subvec, 5, func(meta interface{}, v interface{}, start interface{}, end interface{}, __hash interface{}) interface{} {
			for {
				if Value_(v).Type().AssignableTo(reflect.TypeOf((**CljsCoreSubvec)(nil)).Elem()) {
					meta, v, start, end, __hash = meta, Native_get_instance_field.X_invoke_Arity2(v, "V"), (Native_get_instance_field.X_invoke_Arity2(v, "Start").(float64) + start.(float64)), (Native_get_instance_field.X_invoke_Arity2(v, "Start").(float64) + end.(float64)), __hash
					continue
				} else {
					{
						var c = Count.X_invoke_Arity1(v).(float64)
						_ = c
						if (start.(float64) < float64(0)) || (end.(float64) < float64(0)) || (start.(float64) > c) || (end.(float64) > c) {
							panic((&js.Error{"Index out of bounds"}))
						} else {
						}
						return (&CljsCoreSubvec{meta, v, start, end, __hash})
					}
				}
			}
		})
	}(&AFn{})

	Subvec = func(subvec *AFn) *AFn {
		return Fn(subvec, 3, func(v interface{}, start interface{}) interface{} {
			return subvec.X_invoke_Arity3(v, start, Count.X_invoke_Arity1(v).(float64)).(*CljsCoreSubvec)
		}, func(v interface{}, start interface{}, end interface{}) interface{} {
			return Build_subvec.X_invoke_Arity5(nil, v, start, end, nil).(*CljsCoreSubvec)
		})
	}(&AFn{})

	Tv_ensure_editable = func(tv_ensure_editable *AFn) *AFn {
		return Fn(tv_ensure_editable, 2, func(edit interface{}, node interface{}) interface{} {
			if reflect.DeepEqual(edit, Native_get_instance_field.X_invoke_Arity2(node, "Edit")) {
				return node
			} else {
				return (&CljsCoreVectorNode{edit, Aclone.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(node, "Arr")).([]interface{})})
			}
		})
	}(&AFn{})

	Tv_editable_root = func(tv_editable_root *AFn) *AFn {
		return Fn(tv_editable_root, 1, func(node interface{}) interface{} {
			return (&CljsCoreVectorNode{true, Aclone.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(node, "Arr")).([]interface{})})
		})
	}(&AFn{})

	Tv_editable_tail = func(tv_editable_tail *AFn) *AFn {
		return Fn(tv_editable_tail, 1, func(tl interface{}) interface{} {
			{
				var ret = make([]interface{}, int(float64(32)))
				_ = ret
				Array_copy.X_invoke_Arity5(tl, float64(0), ret, float64(0), Alength_(tl))
				return ret
			}
		})
	}(&AFn{})

	Tv_push_tail = func(tv_push_tail *AFn) *AFn {
		return Fn(tv_push_tail, 4, func(tv interface{}, level interface{}, parent interface{}, tail_node interface{}) interface{} {
			{
				var ret = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tv, "Root"), "Edit"), parent)
				var subidx = float64((Int32_(float64((UInt32_((Native_get_instance_field.X_invoke_Arity2(tv, "Cnt").(float64) - float64(1))) >> UInt32_(float64((32+Int32_(level.(float64)))%32))))) & Int32_(float64(31))))
				_, _ = ret, subidx
				Pv_aset.X_invoke_Arity3(ret, subidx, func() interface{} {
					if level.(float64) == float64(5) {
						return tail_node
					} else {
						return func() interface{} {
							var child = Pv_aget.X_invoke_Arity2(ret, subidx)
							_ = child
							if !(Nil_(child)) {
								return tv_push_tail.X_invoke_Arity4(tv, (level.(float64) - float64(5)), child, tail_node)
							} else {
								return New_path.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tv, "Root"), "Edit"), (level.(float64) - float64(5)), tail_node)
							}
						}()
					}
				}())
				return ret
			}
		})
	}(&AFn{})

	Tv_pop_tail = func(tv_pop_tail *AFn) *AFn {
		return Fn(tv_pop_tail, 3, func(tv interface{}, level interface{}, node interface{}) interface{} {
			{
				var node___1 = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tv, "Root"), "Edit"), node)
				var subidx = float64((Int32_(float64((UInt32_((Native_get_instance_field.X_invoke_Arity2(tv, "Cnt").(float64) - float64(2))) >> UInt32_(float64((32+Int32_(level.(float64)))%32))))) & Int32_(float64(31))))
				_, _ = node___1, subidx
				if level.(float64) > float64(5) {
					{
						var new_child = tv_pop_tail.X_invoke_Arity3(tv, (level.(float64) - float64(5)), Pv_aget.X_invoke_Arity2(node___1, subidx))
						_ = new_child
						if (Nil_(new_child)) && (subidx == float64(0)) {
							return nil
						} else {
							Pv_aset.X_invoke_Arity3(node___1, subidx, new_child)
							return node___1
						}
					}
				} else {
					if subidx == float64(0) {
						return nil
					} else {
						Pv_aset.X_invoke_Arity3(node___1, subidx, nil)
						return node___1

					}
				}
			}
		})
	}(&AFn{})

	Unchecked_editable_array_for = func(unchecked_editable_array_for *AFn) *AFn {
		return Fn(unchecked_editable_array_for, 2, func(tv interface{}, i interface{}) interface{} {
			if i.(float64) >= Tail_off.X_invoke_Arity1(tv).(float64) {
				return Native_get_instance_field.X_invoke_Arity2(tv, "Tail")
			} else {
				{
					var root = Native_get_instance_field.X_invoke_Arity2(tv, "Root")
					_ = root
					{
						var node interface{} = root
						var level interface{} = Native_get_instance_field.X_invoke_Arity2(tv, "Shift")
						_, _ = node, level
						for {
							if level.(float64) > float64(0) {
								node, level = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(root, "Edit"), Pv_aget.X_invoke_Arity2(node, float64((Int32_(float64((UInt32_(i.(float64))>>UInt32_(float64((32+Int32_(level.(float64)))%32)))))&Int32_(float64(31)))))), (level.(float64) - float64(5))
								continue
							} else {
								return Native_get_instance_field.X_invoke_Arity2(node, "Arr")
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	X__GT_TransientVector = func(__GT_TransientVector *AFn) *AFn {
		return Fn(__GT_TransientVector, 4, func(cnt interface{}, shift interface{}, root interface{}, tail interface{}) interface{} {
			return (&CljsCoreTransientVector{cnt, shift, root, tail})
		})
	}(&AFn{})

	X__GT_PersistentQueueSeq = func(__GT_PersistentQueueSeq *AFn) *AFn {
		return Fn(__GT_PersistentQueueSeq, 4, func(meta interface{}, front interface{}, rear interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentQueueSeq{meta, front, rear, __hash})
		})
	}(&AFn{})

	X__GT_PersistentQueue = func(__GT_PersistentQueue *AFn) *AFn {
		return Fn(__GT_PersistentQueue, 5, func(meta interface{}, count interface{}, front interface{}, rear interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentQueue{meta, count, front, rear, __hash})
		})
	}(&AFn{})

	X__GT_NeverEquiv = func(__GT_NeverEquiv *AFn) *AFn {
		return Fn(__GT_NeverEquiv, 0, func() interface{} {
			return (&CljsCoreNeverEquiv{})
		})
	}(&AFn{})

	Never_equiv = (&CljsCoreNeverEquiv{})

	Equiv_map = func(equiv_map *AFn) *AFn {
		return Fn(equiv_map, 2, func(x interface{}, y interface{}) interface{} {
			return Boolean.Arity1IB(func() interface{} {
				if Map_QMARK_.Arity1IB(y) {
					return func() interface{} {
						if Count.X_invoke_Arity1(x).(float64) == Count.X_invoke_Arity1(y).(float64) {
							return Every_QMARK_.Arity2IIB(Identity, Map_.X_invoke_Arity2(func(G__6167 *AFn) *AFn {
								return Fn(G__6167, 1, func(xkv interface{}) interface{} {
									return X_EQ_.Arity2IIB(Get.X_invoke_Arity3(y, First.X_invoke_Arity1(xkv), Never_equiv), Second.X_invoke_Arity1(xkv))
								})
							}(&AFn{}), x).(*CljsCoreLazySeq))
						} else {
							return nil
						}
					}()
				} else {
					return nil
				}
			}())
		})
	}(&AFn{})

	Scan_array = func(scan_array *AFn) *AFn {
		return Fn(scan_array, 3, func(incr interface{}, k interface{}, array interface{}) interface{} {
			{
				var len = Alength_(array)
				_ = len
				{
					var i = float64(0)
					_ = i
					for {
						if i < len {
							if reflect.DeepEqual(k, Aget_(array, i)) {
								return i
							} else {
								i = (i + incr.(float64))
								continue
							}
						} else {
							return nil
						}
					}
				}
			}
		})
	}(&AFn{})

	Obj_map_compare_keys = func(obj_map_compare_keys *AFn) *AFn {
		return Fn(obj_map_compare_keys, 2, func(a interface{}, b interface{}) interface{} {
			{
				var a___1 = Hash.X_invoke_Arity1(a)
				var b___1 = Hash.X_invoke_Arity1(b)
				_, _ = a___1, b___1
				if a___1.(float64) < b___1.(float64) {
					return float64(-1)
				} else {
					if a___1.(float64) > b___1.(float64) {
						return float64(1)
					} else {
						return float64(0)

					}
				}
			}
		})
	}(&AFn{})

	Obj_map__GT_hash_map = func(obj_map__GT_hash_map *AFn) *AFn {
		return Fn(obj_map__GT_hash_map, 3, func(m interface{}, k interface{}, v interface{}) interface{} {
			{
				var ks = Native_get_instance_field.X_invoke_Arity2(m, "Keys")
				var len = Alength_(ks)
				var so = Native_get_instance_field.X_invoke_Arity2(m, "Strobj")
				var mm = Meta.X_invoke_Arity1(m)
				_, _, _, _ = ks, len, so, mm
				{
					var i = float64(0)
					var out interface{} = Transient.X_invoke_Arity1(CljsCorePersistentHashMap_EMPTY)
					_, _ = i, out
					for {
						if i < len {
							{
								var k___1 = Aget_(ks, i)
								_ = k___1
								i, out = (i + float64(1)), Assoc_BANG_.X_invoke_Arity3(out, k___1, Aget_(so, k___1.(float64)))
								continue
							}
						} else {
							return With_meta.X_invoke_Arity2(Persistent_BANG_.X_invoke_Arity1(Assoc_BANG_.X_invoke_Arity3(out, k, v)), mm)
						}
					}
				}
			}
		})
	}(&AFn{})

	X__GT_Iterator = func(__GT_Iterator *AFn) *AFn {
		return Fn(__GT_Iterator, 1, func(s interface{}) interface{} {
			return (&CljsCoreIterator{s})
		})
	}(&AFn{})

	Iterator = func(iterator *AFn) *AFn {
		return Fn(iterator, 1, func(coll interface{}) interface{} {
			return (&CljsCoreIterator{Seq.Arity1IQ(coll)})
		})
	}(&AFn{})

	X__GT_EntriesIterator = func(__GT_EntriesIterator *AFn) *AFn {
		return Fn(__GT_EntriesIterator, 1, func(s interface{}) interface{} {
			return (&CljsCoreEntriesIterator{s})
		})
	}(&AFn{})

	Entries_iterator = func(entries_iterator *AFn) *AFn {
		return Fn(entries_iterator, 1, func(coll interface{}) interface{} {
			return (&CljsCoreEntriesIterator{Seq.Arity1IQ(coll)})
		})
	}(&AFn{})

	X__GT_SetEntriesIterator = func(__GT_SetEntriesIterator *AFn) *AFn {
		return Fn(__GT_SetEntriesIterator, 1, func(s interface{}) interface{} {
			return (&CljsCoreSetEntriesIterator{s})
		})
	}(&AFn{})

	Set_entries_iterator = func(set_entries_iterator *AFn) *AFn {
		return Fn(set_entries_iterator, 1, func(coll interface{}) interface{} {
			return (&CljsCoreSetEntriesIterator{Seq.Arity1IQ(coll)})
		})
	}(&AFn{})

	Array_map_index_of_nil_QMARK_ = func(array_map_index_of_nil_QMARK_ *AFn) *AFn {
		return Fn(array_map_index_of_nil_QMARK_, 3, func(arr interface{}, m interface{}, k interface{}) interface{} {
			{
				var len = Alength_(arr)
				_ = len
				{
					var i = float64(0)
					_ = i
					for {
						if len <= i {
							return float64(-1)
						} else {
							if Nil_(Aget_(arr, i)) {
								return i
							} else {
								i = (i + float64(2))
								continue

							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Array_map_index_of_keyword_QMARK_ = func(array_map_index_of_keyword_QMARK_ *AFn) *AFn {
		return Fn(array_map_index_of_keyword_QMARK_, 3, func(arr interface{}, m interface{}, k interface{}) interface{} {
			{
				var len = Alength_(arr)
				var kstr = Native_get_instance_field.X_invoke_Arity2(k, "Fqn")
				_, _ = len, kstr
				{
					var i = float64(0)
					_ = i
					for {
						if len <= i {
							return float64(-1)
						} else {
							if func() bool {
								var k_SINGLEQUOTE_ = Aget_(arr, i)
								_ = k_SINGLEQUOTE_
								return (Value_(k_SINGLEQUOTE_).Type().AssignableTo(reflect.TypeOf((**CljsCoreKeyword)(nil)).Elem())) && (reflect.DeepEqual(kstr, Native_get_instance_field.X_invoke_Arity2(k_SINGLEQUOTE_, "Fqn")))
							}() {
								return i
							} else {
								i = (i + float64(2))
								continue

							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Array_map_index_of_symbol_QMARK_ = func(array_map_index_of_symbol_QMARK_ *AFn) *AFn {
		return Fn(array_map_index_of_symbol_QMARK_, 3, func(arr interface{}, m interface{}, k interface{}) interface{} {
			{
				var len = Alength_(arr)
				var kstr = Native_get_instance_field.X_invoke_Arity2(k, "Str")
				_, _ = len, kstr
				{
					var i = float64(0)
					_ = i
					for {
						if len <= i {
							return float64(-1)
						} else {
							if func() bool {
								var k_SINGLEQUOTE_ = Aget_(arr, i)
								_ = k_SINGLEQUOTE_
								return (Value_(k_SINGLEQUOTE_).Type().AssignableTo(reflect.TypeOf((**CljsCoreSymbol)(nil)).Elem())) && (reflect.DeepEqual(kstr, Native_get_instance_field.X_invoke_Arity2(k_SINGLEQUOTE_, "Str")))
							}() {
								return i
							} else {
								i = (i + float64(2))
								continue

							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Array_map_index_of_identical_QMARK_ = func(array_map_index_of_identical_QMARK_ *AFn) *AFn {
		return Fn(array_map_index_of_identical_QMARK_, 3, func(arr interface{}, m interface{}, k interface{}) interface{} {
			{
				var len = Alength_(arr)
				_ = len
				{
					var i = float64(0)
					_ = i
					for {
						if len <= i {
							return float64(-1)
						} else {
							if reflect.DeepEqual(k, Aget_(arr, i)) {
								return i
							} else {
								i = (i + float64(2))
								continue

							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Array_map_index_of_equiv_QMARK_ = func(array_map_index_of_equiv_QMARK_ *AFn) *AFn {
		return Fn(array_map_index_of_equiv_QMARK_, 3, func(arr interface{}, m interface{}, k interface{}) interface{} {
			{
				var len = Alength_(arr)
				_ = len
				{
					var i = float64(0)
					_ = i
					for {
						if len <= i {
							return float64(-1)
						} else {
							if X_EQ_.Arity2IIB(k, Aget_(arr, i)) {
								return i
							} else {
								i = (i + float64(2))
								continue

							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Array_map_index_of = func(array_map_index_of *AFn) *AFn {
		return Fn(array_map_index_of, 2, func(m interface{}, k interface{}) interface{} {
			{
				var arr = Native_get_instance_field.X_invoke_Arity2(m, "Arr")
				_ = arr
				if Value_(k).Type().AssignableTo(reflect.TypeOf((**CljsCoreKeyword)(nil)).Elem()) {
					return Array_map_index_of_keyword_QMARK_.X_invoke_Arity3(arr, m, k).(float64)
				} else {
					if Truth_(func() interface{} {
						var or__171__auto__ = func() interface{} {
							var G__6183 = k
							_ = G__6183
							return Native_invoke_func.X_invoke_Arity2(goog.IsString, []interface{}{G__6183})
						}()
						_ = or__171__auto__
						if Truth_(or__171__auto__) {
							return or__171__auto__
						} else {
							return Value_(k).Kind() == reflect.Float64
						}
					}()) {
						return Array_map_index_of_identical_QMARK_.X_invoke_Arity3(arr, m, k).(float64)
					} else {
						if Value_(k).Type().AssignableTo(reflect.TypeOf((**CljsCoreSymbol)(nil)).Elem()) {
							return Array_map_index_of_symbol_QMARK_.X_invoke_Arity3(arr, m, k).(float64)
						} else {
							if Nil_(k) {
								return Array_map_index_of_nil_QMARK_.X_invoke_Arity3(arr, m, k).(float64)
							} else {
								return Array_map_index_of_equiv_QMARK_.X_invoke_Arity3(arr, m, k).(float64)

							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Array_map_extend_kv = func(array_map_extend_kv *AFn) *AFn {
		return Fn(array_map_extend_kv, 3, func(m interface{}, k interface{}, v interface{}) interface{} {
			{
				var arr = Native_get_instance_field.X_invoke_Arity2(m, "Arr")
				var l = Alength_(arr)
				var narr = make([]interface{}, int((l + float64(2))))
				_, _, _ = arr, l, narr
				{
					var i_6184 = float64(0)
					_ = i_6184
					for {
						if i_6184 < l {
							narr[int(i_6184)] = Aget_(arr, i_6184)
							i_6184 = (i_6184 + float64(1))
							continue
						} else {
						}
						break
					}
				}
				narr[int(l)] = k
				narr[int((l + float64(1)))] = v
				return narr
			}
		})
	}(&AFn{})

	X__GT_PersistentArrayMapSeq = func(__GT_PersistentArrayMapSeq *AFn) *AFn {
		return Fn(__GT_PersistentArrayMapSeq, 3, func(arr interface{}, i interface{}, _meta interface{}) interface{} {
			return (&CljsCorePersistentArrayMapSeq{arr, i, _meta})
		})
	}(&AFn{})

	Persistent_array_map_seq = func(persistent_array_map_seq *AFn) *AFn {
		return Fn(persistent_array_map_seq, 3, func(arr interface{}, i interface{}, _meta interface{}) interface{} {
			if i.(float64) <= (Alength_(arr) - float64(2)) {
				return (&CljsCorePersistentArrayMapSeq{arr, i, _meta})
			} else {
				return nil
			}
		})
	}(&AFn{})

	X__GT_PersistentArrayMap = func(__GT_PersistentArrayMap *AFn) *AFn {
		return Fn(__GT_PersistentArrayMap, 4, func(meta interface{}, cnt interface{}, arr interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentArrayMap{meta, cnt, arr, __hash})
		})
	}(&AFn{})

	X__GT_TransientArrayMap = func(__GT_TransientArrayMap *AFn) *AFn {
		return Fn(__GT_TransientArrayMap, 3, func(editable_QMARK_ interface{}, len interface{}, arr interface{}) interface{} {
			return (&CljsCoreTransientArrayMap{editable_QMARK_, len, arr})
		})
	}(&AFn{})

	Array__GT_transient_hash_map = func(array__GT_transient_hash_map *AFn) *AFn {
		return Fn(array__GT_transient_hash_map, 2, func(len interface{}, arr interface{}) interface{} {
			{
				var out interface{} = Transient.X_invoke_Arity1(CljsCorePersistentHashMap_EMPTY)
				var i = float64(0)
				_, _ = out, i
				for {
					if i < len.(float64) {
						out, i = Assoc_BANG_.X_invoke_Arity3(out, Aget_(arr, i), Aget_(arr, (i+float64(1)))), (i + float64(2))
						continue
					} else {
						return out
					}
				}
			}
		})
	}(&AFn{})

	X__GT_Box = func(__GT_Box *AFn) *AFn {
		return Fn(__GT_Box, 1, func(val interface{}) interface{} {
			return (&CljsCoreBox{val})
		})
	}(&AFn{})

	Key_test = func(key_test *AFn) *AFn {
		return Fn(key_test, 2, func(key interface{}, other interface{}) bool {
			if reflect.DeepEqual(key, other) {
				return true
			} else {
				if Keyword_identical_QMARK_.Arity2IIB(key, other) {
					return true
				} else {
					return X_EQ_.Arity2IIB(key, other)

				}
			}
		})
	}(&AFn{})

	Mask = func(mask *AFn) *AFn {
		return Fn(mask, 2, func(hash interface{}, shift interface{}) interface{} {
			return float64((Int32_(float64((UInt32_(hash.(float64)) >> UInt32_(float64((32+Int32_(shift.(float64)))%32))))) & Int32_(float64(31))))
		})
	}(&AFn{})

	Clone_and_set = func(clone_and_set *AFn) *AFn {
		return Fn(clone_and_set, 5, func(arr interface{}, i interface{}, a interface{}) interface{} {
			{
				var G__6230 = Aclone.X_invoke_Arity1(arr).([]interface{})
				_ = G__6230
				G__6230[int(i.(float64))] = a
				return G__6230
			}
		}, func(arr interface{}, i interface{}, a interface{}, j interface{}, b interface{}) interface{} {
			{
				var G__6231 = Aclone.X_invoke_Arity1(arr).([]interface{})
				_ = G__6231
				G__6231[int(i.(float64))] = a
				G__6231[int(j.(float64))] = b
				return G__6231
			}
		})
	}(&AFn{})

	Remove_pair = func(remove_pair *AFn) *AFn {
		return Fn(remove_pair, 2, func(arr interface{}, i interface{}) interface{} {
			{
				var new_arr = make([]interface{}, int((Alength_(arr) - float64(2))))
				_ = new_arr
				Array_copy.X_invoke_Arity5(arr, float64(0), new_arr, float64(0), (float64(2) * i.(float64)))
				Array_copy.X_invoke_Arity5(arr, (float64(2) * (i.(float64) + float64(1))), new_arr, (float64(2) * i.(float64)), (float64(len(new_arr)) - (float64(2) * i.(float64))))
				return new_arr
			}
		})
	}(&AFn{})

	Bitmap_indexed_node_index = func(bitmap_indexed_node_index *AFn) *AFn {
		return Fn(bitmap_indexed_node_index, 2, func(bitmap interface{}, bit interface{}) interface{} {
			return Bit_count.X_invoke_Arity1(float64((Int32_(bitmap.(float64)) & Int32_((bit.(float64) - float64(1)))))).(float64)
		})
	}(&AFn{})

	Bitpos = func(bitpos *AFn) *AFn {
		return Fn(bitpos, 2, func(hash interface{}, shift interface{}) interface{} {
			return float64((Int32_(float64(1)) << UInt32_(float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f)))))
		})
	}(&AFn{})

	Edit_and_set = func(edit_and_set *AFn) *AFn {
		return Fn(edit_and_set, 6, func(inode interface{}, edit interface{}, i interface{}, a interface{}) interface{} {
			{
				var editable = Native_invoke_instance_method.X_invoke_Arity3(inode, "Ensure_editable", []interface{}{edit})
				_ = editable
				Native_get_instance_field.X_invoke_Arity2(editable, "Arr").([]interface{})[int(i.(float64))] = a
				return editable
			}
		}, func(inode interface{}, edit interface{}, i interface{}, a interface{}, j interface{}, b interface{}) interface{} {
			{
				var editable = Native_invoke_instance_method.X_invoke_Arity3(inode, "Ensure_editable", []interface{}{edit})
				_ = editable
				Native_get_instance_field.X_invoke_Arity2(editable, "Arr").([]interface{})[int(i.(float64))] = a
				Native_get_instance_field.X_invoke_Arity2(editable, "Arr").([]interface{})[int(j.(float64))] = b
				return editable
			}
		})
	}(&AFn{})

	Inode_kv_reduce = func(inode_kv_reduce *AFn) *AFn {
		return Fn(inode_kv_reduce, 3, func(arr interface{}, f interface{}, init interface{}) interface{} {
			{
				var len = Alength_(arr)
				_ = len
				{
					var i = float64(0)
					var init___1 interface{} = init
					_, _ = i, init___1
					for {
						if i < len {
							{
								var init___2 = func() interface{} {
									var k = Aget_(arr, i)
									_ = k
									if !(Nil_(k)) {
										{
											var G__6235 = init___1
											var G__6236 = k
											var G__6237 = Aget_(arr, (i + float64(1)))
											_, _, _ = G__6235, G__6236, G__6237
											return f.(CljsCoreIFn).X_invoke_Arity3(G__6235, G__6236, G__6237)
										}
									} else {
										{
											var node = Aget_(arr, (i + float64(1)))
											_ = node
											if !(Nil_(node)) {
												return Native_invoke_instance_method.X_invoke_Arity3(node, "Kv_reduce", []interface{}{f, init___1})
											} else {
												return init___1
											}
										}
									}
								}()
								_ = init___2
								if Reduced_QMARK_.Arity1IB(init___2) {
									return Deref.X_invoke_Arity1(init___2)
								} else {
									i, init___1 = (i + float64(2)), init___2
									continue
								}
							}
						} else {
							return init___1
						}
					}
				}
			}
		})
	}(&AFn{})

	X__GT_BitmapIndexedNode = func(__GT_BitmapIndexedNode *AFn) *AFn {
		return Fn(__GT_BitmapIndexedNode, 3, func(edit interface{}, bitmap interface{}, arr interface{}) interface{} {
			return (&CljsCoreBitmapIndexedNode{edit, bitmap, arr})
		})
	}(&AFn{})

	Pack_array_node = func(pack_array_node *AFn) *AFn {
		return Fn(pack_array_node, 3, func(array_node interface{}, edit interface{}, idx interface{}) interface{} {
			{
				var arr = Native_get_instance_field.X_invoke_Arity2(array_node, "Arr")
				var len = Alength_(arr)
				var new_arr = make([]interface{}, int((float64(2) * (Native_get_instance_field.X_invoke_Arity2(array_node, "Cnt").(float64) - float64(1)))))
				_, _, _ = arr, len, new_arr
				{
					var i = float64(0)
					var j = float64(1)
					var bitmap = float64(0)
					_, _, _ = i, j, bitmap
					for {
						if i < len {
							if (!(i == idx.(float64))) && (!(Nil_(Aget_(arr, i)))) {
								new_arr[int(j)] = Aget_(arr, i)
								i, j, bitmap = (i + float64(1)), (j + float64(2)), float64((Int32_(bitmap) | Int32_(float64((Int32_(float64(1)) << UInt32_(i))))))
								continue
							} else {
								i, j, bitmap = (i + float64(1)), j, bitmap
								continue
							}
						} else {
							return (&CljsCoreBitmapIndexedNode{edit, bitmap, new_arr})
						}
					}
				}
			}
		})
	}(&AFn{})

	X__GT_ArrayNode = func(__GT_ArrayNode *AFn) *AFn {
		return Fn(__GT_ArrayNode, 3, func(edit interface{}, cnt interface{}, arr interface{}) interface{} {
			return (&CljsCoreArrayNode{edit, cnt, arr})
		})
	}(&AFn{})

	Hash_collision_node_find_index = func(hash_collision_node_find_index *AFn) *AFn {
		return Fn(hash_collision_node_find_index, 3, func(arr interface{}, cnt interface{}, key interface{}) interface{} {
			{
				var lim = (float64(2) * cnt.(float64))
				_ = lim
				{
					var i = float64(0)
					_ = i
					for {
						if i < lim {
							if Key_test.Arity2IIB(key, Aget_(arr, i)) {
								return i
							} else {
								i = (i + float64(2))
								continue
							}
						} else {
							return float64(-1)
						}
					}
				}
			}
		})
	}(&AFn{})

	X__GT_HashCollisionNode = func(__GT_HashCollisionNode *AFn) *AFn {
		return Fn(__GT_HashCollisionNode, 4, func(edit interface{}, collision_hash interface{}, cnt interface{}, arr interface{}) interface{} {
			return (&CljsCoreHashCollisionNode{edit, collision_hash, cnt, arr})
		})
	}(&AFn{})

	Create_node = func(create_node *AFn) *AFn {
		return Fn(create_node, 7, func(shift interface{}, key1 interface{}, val1 interface{}, key2hash interface{}, key2 interface{}, val2 interface{}) interface{} {
			{
				var key1hash = Hash.X_invoke_Arity1(key1)
				_ = key1hash
				if key1hash.(float64) == key2hash.(float64) {
					return (&CljsCoreHashCollisionNode{nil, key1hash, float64(2), []interface{}{key1, val1, key2, val2}})
				} else {
					{
						var added_leaf_QMARK_ = (&CljsCoreBox{false})
						_ = added_leaf_QMARK_
						return Native_invoke_instance_method.X_invoke_Arity3(Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc", []interface{}{shift, key1hash, key1, val1, added_leaf_QMARK_}), "Inode_assoc", []interface{}{shift, key2hash, key2, val2, added_leaf_QMARK_})
					}
				}
			}
		}, func(edit interface{}, shift interface{}, key1 interface{}, val1 interface{}, key2hash interface{}, key2 interface{}, val2 interface{}) interface{} {
			{
				var key1hash = Hash.X_invoke_Arity1(key1)
				_ = key1hash
				if key1hash.(float64) == key2hash.(float64) {
					return (&CljsCoreHashCollisionNode{nil, key1hash, float64(2), []interface{}{key1, val1, key2, val2}})
				} else {
					{
						var added_leaf_QMARK_ = (&CljsCoreBox{false})
						_ = added_leaf_QMARK_
						return Native_invoke_instance_method.X_invoke_Arity3(Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc_BANG_", []interface{}{edit, shift, key1hash, key1, val1, added_leaf_QMARK_}), "Inode_assoc_BANG_", []interface{}{edit, shift, key2hash, key2, val2, added_leaf_QMARK_})
					}
				}
			}
		})
	}(&AFn{})

	X__GT_NodeSeq = func(__GT_NodeSeq *AFn) *AFn {
		return Fn(__GT_NodeSeq, 5, func(meta interface{}, nodes interface{}, i interface{}, s interface{}, __hash interface{}) interface{} {
			return (&CljsCoreNodeSeq{meta, nodes, i, s, __hash})
		})
	}(&AFn{})

	Create_inode_seq = func(create_inode_seq *AFn) *AFn {
		return Fn(create_inode_seq, 3, func(nodes interface{}) interface{} {
			return create_inode_seq.X_invoke_Arity3(nodes, float64(0), nil)
		}, func(nodes interface{}, i interface{}, s interface{}) interface{} {
			if Nil_(s) {
				{
					var len = Alength_(nodes)
					_ = len
					{
						var j interface{} = i
						_ = j
						for {
							if j.(float64) < len {
								if !(Nil_(Aget_(nodes, j.(float64)))) {
									return (&CljsCoreNodeSeq{nil, nodes, j, nil, nil})
								} else {
									{
										var temp__4220__auto__ = Aget_(nodes, (j.(float64) + float64(1)))
										_ = temp__4220__auto__
										if Truth_(temp__4220__auto__) {
											{
												var node = temp__4220__auto__
												_ = node
												{
													var temp__4220__auto_____1 = Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_seq", []interface{}{})
													_ = temp__4220__auto_____1
													if Truth_(temp__4220__auto_____1) {
														{
															var node_seq = temp__4220__auto_____1
															_ = node_seq
															return (&CljsCoreNodeSeq{nil, nodes, (j.(float64) + float64(2)), node_seq, nil})
														}
													} else {
														j = (j.(float64) + float64(2))
														continue
													}
												}
											}
										} else {
											j = (j.(float64) + float64(2))
											continue
										}
									}
								}
							} else {
								return nil
							}
						}
					}
				}
			} else {
				return (&CljsCoreNodeSeq{nil, nodes, i, s, nil})
			}
		})
	}(&AFn{})

	X__GT_ArrayNodeSeq = func(__GT_ArrayNodeSeq *AFn) *AFn {
		return Fn(__GT_ArrayNodeSeq, 5, func(meta interface{}, nodes interface{}, i interface{}, s interface{}, __hash interface{}) interface{} {
			return (&CljsCoreArrayNodeSeq{meta, nodes, i, s, __hash})
		})
	}(&AFn{})

	Create_array_node_seq = func(create_array_node_seq *AFn) *AFn {
		return Fn(create_array_node_seq, 4, func(nodes interface{}) interface{} {
			return create_array_node_seq.X_invoke_Arity4(nil, nodes, float64(0), nil)
		}, func(meta interface{}, nodes interface{}, i interface{}, s interface{}) interface{} {
			if Nil_(s) {
				{
					var len = Alength_(nodes)
					_ = len
					{
						var j interface{} = i
						_ = j
						for {
							if j.(float64) < len {
								{
									var temp__4220__auto__ = Aget_(nodes, j.(float64))
									_ = temp__4220__auto__
									if Truth_(temp__4220__auto__) {
										{
											var nj = temp__4220__auto__
											_ = nj
											{
												var temp__4220__auto_____1 = Native_invoke_instance_method.X_invoke_Arity3(nj, "Inode_seq", []interface{}{})
												_ = temp__4220__auto_____1
												if Truth_(temp__4220__auto_____1) {
													{
														var ns = temp__4220__auto_____1
														_ = ns
														return (&CljsCoreArrayNodeSeq{meta, nodes, (j.(float64) + float64(1)), ns, nil})
													}
												} else {
													j = (j.(float64) + float64(1))
													continue
												}
											}
										}
									} else {
										j = (j.(float64) + float64(1))
										continue
									}
								}
							} else {
								return nil
							}
						}
					}
				}
			} else {
				return (&CljsCoreArrayNodeSeq{meta, nodes, i, s, nil})
			}
		})
	}(&AFn{})

	X__GT_PersistentHashMap = func(__GT_PersistentHashMap *AFn) *AFn {
		return Fn(__GT_PersistentHashMap, 6, func(meta interface{}, cnt interface{}, root interface{}, has_nil_QMARK_ bool, nil_val interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentHashMap{meta, cnt, root, has_nil_QMARK_, nil_val, __hash})
		})
	}(&AFn{})

	X__GT_TransientHashMap = func(__GT_TransientHashMap *AFn) *AFn {
		return Fn(__GT_TransientHashMap, 5, func(edit bool, root interface{}, count interface{}, has_nil_QMARK_ bool, nil_val interface{}) interface{} {
			return (&CljsCoreTransientHashMap{edit, root, count, has_nil_QMARK_, nil_val})
		})
	}(&AFn{})

	Tree_map_seq_push = func(tree_map_seq_push *AFn) *AFn {
		return Fn(tree_map_seq_push, 3, func(node interface{}, stack interface{}, ascending_QMARK_ bool) interface{} {
			{
				var t interface{} = node
				var stack___1 interface{} = stack
				_, _ = t, stack___1
				for {
					if !(Nil_(t)) {
						t, stack___1 = func() interface{} {
							if Truth_(ascending_QMARK_) {
								return Native_get_instance_field.X_invoke_Arity2(t, "Left")
							} else {
								return Native_get_instance_field.X_invoke_Arity2(t, "Right")
							}
						}(), Conj.X_invoke_Arity2(stack___1, t)
						continue
					} else {
						return stack___1
					}
				}
			}
		})
	}(&AFn{})

	X__GT_PersistentTreeMapSeq = func(__GT_PersistentTreeMapSeq *AFn) *AFn {
		return Fn(__GT_PersistentTreeMapSeq, 5, func(meta interface{}, stack interface{}, ascending_QMARK_ bool, cnt interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentTreeMapSeq{meta, stack, ascending_QMARK_, cnt, __hash})
		})
	}(&AFn{})

	Create_tree_map_seq = func(create_tree_map_seq *AFn) *AFn {
		return Fn(create_tree_map_seq, 3, func(tree interface{}, ascending_QMARK_ interface{}, cnt interface{}) interface{} {
			return (&CljsCorePersistentTreeMapSeq{nil, Tree_map_seq_push.X_invoke_Arity3(tree, nil, ascending_QMARK_), ascending_QMARK_.(bool), cnt, nil})
		})
	}(&AFn{})

	Balance_left = func(balance_left *AFn) *AFn {
		return Fn(balance_left, 4, func(key interface{}, val interface{}, ins interface{}, right interface{}) interface{} {
			if Value_(ins).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
				if Value_(Native_get_instance_field.X_invoke_Arity2(ins, "Left")).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
					return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(ins, "Key"), Native_get_instance_field.X_invoke_Arity2(ins, "Val"), Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(ins, "Left"), "Blacken", []interface{}{}), (&CljsCoreBlackNode{key, val, Native_get_instance_field.X_invoke_Arity2(ins, "Right"), right, nil}), nil})
				} else {
					if Value_(Native_get_instance_field.X_invoke_Arity2(ins, "Right")).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
						return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(ins, "Right"), "Key"), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(ins, "Right"), "Val"), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(ins, "Key"), Native_get_instance_field.X_invoke_Arity2(ins, "Val"), Native_get_instance_field.X_invoke_Arity2(ins, "Left"), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(ins, "Right"), "Left"), nil}), (&CljsCoreBlackNode{key, val, Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(ins, "Right"), "Right"), right, nil}), nil})
					} else {
						return (&CljsCoreBlackNode{key, val, ins, right, nil})

					}
				}
			} else {
				return (&CljsCoreBlackNode{key, val, ins, right, nil})
			}
		})
	}(&AFn{})

	Balance_right = func(balance_right *AFn) *AFn {
		return Fn(balance_right, 4, func(key interface{}, val interface{}, left interface{}, ins interface{}) interface{} {
			if Value_(ins).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
				if Value_(Native_get_instance_field.X_invoke_Arity2(ins, "Right")).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
					return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(ins, "Key"), Native_get_instance_field.X_invoke_Arity2(ins, "Val"), (&CljsCoreBlackNode{key, val, left, Native_get_instance_field.X_invoke_Arity2(ins, "Left"), nil}), Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(ins, "Right"), "Blacken", []interface{}{}), nil})
				} else {
					if Value_(Native_get_instance_field.X_invoke_Arity2(ins, "Left")).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
						return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(ins, "Left"), "Key"), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(ins, "Left"), "Val"), (&CljsCoreBlackNode{key, val, left, Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(ins, "Left"), "Left"), nil}), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(ins, "Key"), Native_get_instance_field.X_invoke_Arity2(ins, "Val"), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(ins, "Left"), "Right"), Native_get_instance_field.X_invoke_Arity2(ins, "Right"), nil}), nil})
					} else {
						return (&CljsCoreBlackNode{key, val, left, ins, nil})

					}
				}
			} else {
				return (&CljsCoreBlackNode{key, val, left, ins, nil})
			}
		})
	}(&AFn{})

	Balance_left_del = func(balance_left_del *AFn) *AFn {
		return Fn(balance_left_del, 4, func(key interface{}, val interface{}, del interface{}, right interface{}) interface{} {
			if Value_(del).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
				return (&CljsCoreRedNode{key, val, Native_invoke_instance_method.X_invoke_Arity3(del, "Blacken", []interface{}{}), right, nil})
			} else {
				if Value_(right).Type().AssignableTo(reflect.TypeOf((**CljsCoreBlackNode)(nil)).Elem()) {
					return Balance_right.X_invoke_Arity4(key, val, del, Native_invoke_instance_method.X_invoke_Arity3(right, "Redden", []interface{}{}))
				} else {
					if (Value_(right).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem())) && (Value_(Native_get_instance_field.X_invoke_Arity2(right, "Left")).Type().AssignableTo(reflect.TypeOf((**CljsCoreBlackNode)(nil)).Elem())) {
						return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(right, "Left"), "Key"), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(right, "Left"), "Val"), (&CljsCoreBlackNode{key, val, del, Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(right, "Left"), "Left"), nil}), Balance_right.X_invoke_Arity4(Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(right, "Left"), "Right"), Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(right, "Right"), "Redden", []interface{}{})), nil})
					} else {
						panic((&js.Error{"red-black tree invariant violation"}))

					}
				}
			}
		})
	}(&AFn{})

	Balance_right_del = func(balance_right_del *AFn) *AFn {
		return Fn(balance_right_del, 4, func(key interface{}, val interface{}, left interface{}, del interface{}) interface{} {
			if Value_(del).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
				return (&CljsCoreRedNode{key, val, left, Native_invoke_instance_method.X_invoke_Arity3(del, "Blacken", []interface{}{}), nil})
			} else {
				if Value_(left).Type().AssignableTo(reflect.TypeOf((**CljsCoreBlackNode)(nil)).Elem()) {
					return Balance_left.X_invoke_Arity4(key, val, Native_invoke_instance_method.X_invoke_Arity3(left, "Redden", []interface{}{}), del)
				} else {
					if (Value_(left).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem())) && (Value_(Native_get_instance_field.X_invoke_Arity2(left, "Right")).Type().AssignableTo(reflect.TypeOf((**CljsCoreBlackNode)(nil)).Elem())) {
						return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), "Key"), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), "Val"), Balance_left.X_invoke_Arity4(Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(left, "Left"), "Redden", []interface{}{}), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), "Left")), (&CljsCoreBlackNode{key, val, Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), "Right"), del, nil}), nil})
					} else {
						panic((&js.Error{"red-black tree invariant violation"}))

					}
				}
			}
		})
	}(&AFn{})

	Tree_map_kv_reduce = func(tree_map_kv_reduce *AFn) *AFn {
		return Fn(tree_map_kv_reduce, 3, func(node interface{}, f interface{}, init interface{}) interface{} {
			{
				var init___1 = func() interface{} {
					if !(Nil_(Native_get_instance_field.X_invoke_Arity2(node, "Left"))) {
						return tree_map_kv_reduce.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(node, "Left"), f, init)
					} else {
						return init
					}
				}()
				_ = init___1
				if Reduced_QMARK_.Arity1IB(init___1) {
					return Deref.X_invoke_Arity1(init___1)
				} else {
					{
						var init___2 = func() interface{} {
							var G__6289 = init___1
							var G__6290 = Native_get_instance_field.X_invoke_Arity2(node, "Key")
							var G__6291 = Native_get_instance_field.X_invoke_Arity2(node, "Val")
							_, _, _ = G__6289, G__6290, G__6291
							return f.(CljsCoreIFn).X_invoke_Arity3(G__6289, G__6290, G__6291)
						}()
						_ = init___2
						if Reduced_QMARK_.Arity1IB(init___2) {
							return Deref.X_invoke_Arity1(init___2)
						} else {
							{
								var init___3 = func() interface{} {
									if !(Nil_(Native_get_instance_field.X_invoke_Arity2(node, "Right"))) {
										return tree_map_kv_reduce.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(node, "Right"), f, init___2)
									} else {
										return init___2
									}
								}()
								_ = init___3
								if Reduced_QMARK_.Arity1IB(init___3) {
									return Deref.X_invoke_Arity1(init___3)
								} else {
									return init___3
								}
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	X__GT_BlackNode = func(__GT_BlackNode *AFn) *AFn {
		return Fn(__GT_BlackNode, 5, func(key interface{}, val interface{}, left interface{}, right interface{}, __hash interface{}) interface{} {
			return (&CljsCoreBlackNode{key, val, left, right, __hash})
		})
	}(&AFn{})

	X__GT_RedNode = func(__GT_RedNode *AFn) *AFn {
		return Fn(__GT_RedNode, 5, func(key interface{}, val interface{}, left interface{}, right interface{}, __hash interface{}) interface{} {
			return (&CljsCoreRedNode{key, val, left, right, __hash})
		})
	}(&AFn{})

	Tree_map_add = func(tree_map_add *AFn) *AFn {
		return Fn(tree_map_add, 5, func(comp interface{}, tree interface{}, k interface{}, v interface{}, found interface{}) interface{} {
			if Nil_(tree) {
				return (&CljsCoreRedNode{k, v, nil, nil, nil})
			} else {
				{
					var c = func() interface{} {
						var G__6310 = k
						var G__6311 = Native_get_instance_field.X_invoke_Arity2(tree, "Key")
						_, _ = G__6310, G__6311
						return comp.(CljsCoreIFn).X_invoke_Arity2(G__6310, G__6311)
					}()
					_ = c
					if c.(float64) == float64(0) {
						found.([]interface{})[int(float64(0))] = tree
						return nil
					} else {
						if c.(float64) < float64(0) {
							{
								var ins = tree_map_add.X_invoke_Arity5(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Left"), k, v, found)
								_ = ins
								if !(Nil_(ins)) {
									return Native_invoke_instance_method.X_invoke_Arity3(tree, "Add_left", []interface{}{ins})
								} else {
									return nil
								}
							}
						} else {
							{
								var ins = tree_map_add.X_invoke_Arity5(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Right"), k, v, found)
								_ = ins
								if !(Nil_(ins)) {
									return Native_invoke_instance_method.X_invoke_Arity3(tree, "Add_right", []interface{}{ins})
								} else {
									return nil
								}
							}

						}
					}
				}
			}
		})
	}(&AFn{})

	Tree_map_append = func(tree_map_append *AFn) *AFn {
		return Fn(tree_map_append, 2, func(left interface{}, right interface{}) interface{} {
			if Nil_(left) {
				return right
			} else {
				if Nil_(right) {
					return left
				} else {
					if Value_(left).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
						if Value_(right).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
							{
								var app = tree_map_append.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), Native_get_instance_field.X_invoke_Arity2(right, "Left"))
								_ = app
								if Value_(app).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
									return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(app, "Key"), Native_get_instance_field.X_invoke_Arity2(app, "Val"), (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_get_instance_field.X_invoke_Arity2(left, "Left"), Native_get_instance_field.X_invoke_Arity2(app, "Left"), nil}), (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), Native_get_instance_field.X_invoke_Arity2(app, "Right"), Native_get_instance_field.X_invoke_Arity2(right, "Right"), nil}), nil})
								} else {
									return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_get_instance_field.X_invoke_Arity2(left, "Left"), (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), app, Native_get_instance_field.X_invoke_Arity2(right, "Right"), nil}), nil})
								}
							}
						} else {
							return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_get_instance_field.X_invoke_Arity2(left, "Left"), tree_map_append.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), right), nil})
						}
					} else {
						if Value_(right).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
							return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), tree_map_append.X_invoke_Arity2(left, Native_get_instance_field.X_invoke_Arity2(right, "Left")), Native_get_instance_field.X_invoke_Arity2(right, "Right"), nil})
						} else {
							{
								var app = tree_map_append.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), Native_get_instance_field.X_invoke_Arity2(right, "Left"))
								_ = app
								if Value_(app).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
									return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(app, "Key"), Native_get_instance_field.X_invoke_Arity2(app, "Val"), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_get_instance_field.X_invoke_Arity2(left, "Left"), Native_get_instance_field.X_invoke_Arity2(app, "Left"), nil}), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), Native_get_instance_field.X_invoke_Arity2(app, "Right"), Native_get_instance_field.X_invoke_Arity2(right, "Right"), nil}), nil})
								} else {
									return Balance_left_del.X_invoke_Arity4(Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_get_instance_field.X_invoke_Arity2(left, "Left"), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), app, Native_get_instance_field.X_invoke_Arity2(right, "Right"), nil}))
								}
							}

						}
					}
				}
			}
		})
	}(&AFn{})

	Tree_map_remove = func(tree_map_remove *AFn) *AFn {
		return Fn(tree_map_remove, 4, func(comp interface{}, tree interface{}, k interface{}, found interface{}) interface{} {
			if !(Nil_(tree)) {
				{
					var c = func() interface{} {
						var G__6330 = k
						var G__6331 = Native_get_instance_field.X_invoke_Arity2(tree, "Key")
						_, _ = G__6330, G__6331
						return comp.(CljsCoreIFn).X_invoke_Arity2(G__6330, G__6331)
					}()
					_ = c
					if c.(float64) == float64(0) {
						found.([]interface{})[int(float64(0))] = tree
						return Tree_map_append.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tree, "Left"), Native_get_instance_field.X_invoke_Arity2(tree, "Right"))
					} else {
						if c.(float64) < float64(0) {
							{
								var del = tree_map_remove.X_invoke_Arity4(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Left"), k, found)
								_ = del
								if (!(Nil_(del))) || (!(Nil_(Aget_(found, float64(0))))) {
									if Value_(Native_get_instance_field.X_invoke_Arity2(tree, "Left")).Type().AssignableTo(reflect.TypeOf((**CljsCoreBlackNode)(nil)).Elem()) {
										return Balance_left_del.X_invoke_Arity4(Native_get_instance_field.X_invoke_Arity2(tree, "Key"), Native_get_instance_field.X_invoke_Arity2(tree, "Val"), del, Native_get_instance_field.X_invoke_Arity2(tree, "Right"))
									} else {
										return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(tree, "Key"), Native_get_instance_field.X_invoke_Arity2(tree, "Val"), del, Native_get_instance_field.X_invoke_Arity2(tree, "Right"), nil})
									}
								} else {
									return nil
								}
							}
						} else {
							{
								var del = tree_map_remove.X_invoke_Arity4(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Right"), k, found)
								_ = del
								if (!(Nil_(del))) || (!(Nil_(Aget_(found, float64(0))))) {
									if Value_(Native_get_instance_field.X_invoke_Arity2(tree, "Right")).Type().AssignableTo(reflect.TypeOf((**CljsCoreBlackNode)(nil)).Elem()) {
										return Balance_right_del.X_invoke_Arity4(Native_get_instance_field.X_invoke_Arity2(tree, "Key"), Native_get_instance_field.X_invoke_Arity2(tree, "Val"), Native_get_instance_field.X_invoke_Arity2(tree, "Left"), del)
									} else {
										return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(tree, "Key"), Native_get_instance_field.X_invoke_Arity2(tree, "Val"), Native_get_instance_field.X_invoke_Arity2(tree, "Left"), del, nil})
									}
								} else {
									return nil
								}
							}

						}
					}
				}
			} else {
				return nil
			}
		})
	}(&AFn{})

	Tree_map_replace = func(tree_map_replace *AFn) *AFn {
		return Fn(tree_map_replace, 4, func(comp interface{}, tree interface{}, k interface{}, v interface{}) interface{} {
			{
				var tk = Native_get_instance_field.X_invoke_Arity2(tree, "Key")
				var c = func() interface{} {
					var G__6342 = k
					var G__6343 = tk
					_, _ = G__6342, G__6343
					return comp.(CljsCoreIFn).X_invoke_Arity2(G__6342, G__6343)
				}()
				_, _ = tk, c
				if c.(float64) == float64(0) {
					return Native_invoke_instance_method.X_invoke_Arity3(tree, "Replace", []interface{}{tk, v, Native_get_instance_field.X_invoke_Arity2(tree, "Left"), Native_get_instance_field.X_invoke_Arity2(tree, "Right")})
				} else {
					if c.(float64) < float64(0) {
						return Native_invoke_instance_method.X_invoke_Arity3(tree, "Replace", []interface{}{tk, Native_get_instance_field.X_invoke_Arity2(tree, "Val"), tree_map_replace.X_invoke_Arity4(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Left"), k, v), Native_get_instance_field.X_invoke_Arity2(tree, "Right")})
					} else {
						return Native_invoke_instance_method.X_invoke_Arity3(tree, "Replace", []interface{}{tk, Native_get_instance_field.X_invoke_Arity2(tree, "Val"), Native_get_instance_field.X_invoke_Arity2(tree, "Left"), tree_map_replace.X_invoke_Arity4(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Right"), k, v)})

					}
				}
			}
		})
	}(&AFn{})

	X__GT_PersistentTreeMap = func(__GT_PersistentTreeMap *AFn) *AFn {
		return Fn(__GT_PersistentTreeMap, 5, func(comp interface{}, tree interface{}, cnt interface{}, meta interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentTreeMap{comp, tree, cnt, meta, __hash})
		})
	}(&AFn{})

	Hash_map = func(hash_map *AFn) *AFn {
		return Fn(hash_map, 0, func(keyvals__ ...interface{}) interface{} {
			var keyvals = Seq.Arity1IQ(keyvals__[0])
			_ = keyvals
			{
				var in interface{} = Seq.Arity1IQ(keyvals)
				var out interface{} = Transient.X_invoke_Arity1(CljsCorePersistentHashMap_EMPTY)
				_, _ = in, out
				for {
					if Truth_(in) {
						in, out = Seq_(Nnext.X_invoke_Arity1(in)), Assoc_BANG_.X_invoke_Arity3(out, First.X_invoke_Arity1(in), Second.X_invoke_Arity1(in))
						continue
					} else {
						return Persistent_BANG_.X_invoke_Arity1(out)
					}
				}
			}
		})
	}(&AFn{})

	Array_map = func(array_map *AFn) *AFn {
		return Fn(array_map, 0, func(keyvals__ ...interface{}) interface{} {
			var keyvals = Seq.Arity1IQ(keyvals__[0])
			_ = keyvals
			return (&CljsCorePersistentArrayMap{nil, Quot.X_invoke_Arity2(Count.X_invoke_Arity1(keyvals).(float64), float64(2)).(float64), Apply.X_invoke_Arity2(Array, keyvals), nil})
		})
	}(&AFn{})

	Sorted_map = func(sorted_map *AFn) *AFn {
		return Fn(sorted_map, 0, func(keyvals__ ...interface{}) interface{} {
			var keyvals = Seq.Arity1IQ(keyvals__[0])
			_ = keyvals
			{
				var in interface{} = Seq.Arity1IQ(keyvals)
				var out interface{} = CljsCorePersistentTreeMap_EMPTY
				_, _ = in, out
				for {
					if Truth_(in) {
						in, out = Seq_(Nnext.X_invoke_Arity1(in)), Assoc.X_invoke_Arity3(out, First.X_invoke_Arity1(in), Second.X_invoke_Arity1(in))
						continue
					} else {
						return out
					}
				}
			}
		})
	}(&AFn{})

	Sorted_map_by = func(sorted_map_by *AFn) *AFn {
		return Fn(sorted_map_by, 1, func(comparator_keyvals__ ...interface{}) interface{} {
			var comparator = comparator_keyvals__[0]
			var keyvals = Seq.Arity1IQ(comparator_keyvals__[1])
			_, _ = comparator, keyvals
			{
				var in interface{} = Seq.Arity1IQ(keyvals)
				var out interface{} = (&CljsCorePersistentTreeMap{Fn__GT_comparator.X_invoke_Arity1(comparator), nil, float64(0), nil, float64(0)})
				_, _ = in, out
				for {
					if Truth_(in) {
						in, out = Seq_(Nnext.X_invoke_Arity1(in)), Assoc.X_invoke_Arity3(out, First.X_invoke_Arity1(in), Second.X_invoke_Arity1(in))
						continue
					} else {
						return out
					}
				}
			}
		})
	}(&AFn{})

	X__GT_KeySeq = func(__GT_KeySeq *AFn) *AFn {
		return Fn(__GT_KeySeq, 2, func(mseq interface{}, _meta interface{}) interface{} {
			return (&CljsCoreKeySeq{mseq, _meta})
		})
	}(&AFn{})

	Keys = func(keys *AFn) *AFn {
		return Fn(keys, 1, func(hash_map interface{}) interface{} {
			{
				var temp__4222__auto__ = Seq.Arity1IQ(hash_map)
				_ = temp__4222__auto__
				if Truth_(temp__4222__auto__) {
					{
						var mseq = temp__4222__auto__
						_ = mseq
						return (&CljsCoreKeySeq{mseq, nil})
					}
				} else {
					return nil
				}
			}
		})
	}(&AFn{})

	Key = func(key *AFn) *AFn {
		return Fn(key, 1, func(map_entry interface{}) interface{} {
			return map_entry.(CljsCoreIMapEntry).X_key_Arity1()
		})
	}(&AFn{})

	X__GT_ValSeq = func(__GT_ValSeq *AFn) *AFn {
		return Fn(__GT_ValSeq, 2, func(mseq interface{}, _meta interface{}) interface{} {
			return (&CljsCoreValSeq{mseq, _meta})
		})
	}(&AFn{})

	Vals = func(vals *AFn) *AFn {
		return Fn(vals, 1, func(hash_map interface{}) interface{} {
			{
				var temp__4222__auto__ = Seq.Arity1IQ(hash_map)
				_ = temp__4222__auto__
				if Truth_(temp__4222__auto__) {
					{
						var mseq = temp__4222__auto__
						_ = mseq
						return (&CljsCoreValSeq{mseq, nil})
					}
				} else {
					return nil
				}
			}
		})
	}(&AFn{})

	Val = func(val *AFn) *AFn {
		return Fn(val, 1, func(map_entry interface{}) interface{} {
			return map_entry.(CljsCoreIMapEntry).X_val_Arity1()
		})
	}(&AFn{})

	Merge = func(merge *AFn) *AFn {
		return Fn(merge, 0, func(maps__ ...interface{}) interface{} {
			var maps = Seq.Arity1IQ(maps__[0])
			_ = maps
			if Truth_(Some.X_invoke_Arity2(Identity, maps)) {
				return Reduce.X_invoke_Arity2(func(G__6386 *AFn) *AFn {
					return Fn(G__6386, 2, func(p1__6384_SHARP_ interface{}, p2__6385_SHARP_ interface{}) interface{} {
						return Conj.X_invoke_Arity2(func() interface{} {
							var or__171__auto__ = p1__6384_SHARP_
							_ = or__171__auto__
							if Truth_(or__171__auto__) {
								return or__171__auto__
							} else {
								return CljsCorePersistentArrayMap_EMPTY
							}
						}(), p2__6385_SHARP_)
					})
				}(&AFn{}), maps)
			} else {
				return nil
			}
		})
	}(&AFn{})

	Merge_with = func(merge_with *AFn) *AFn {
		return Fn(merge_with, 1, func(f_maps__ ...interface{}) interface{} {
			var f = f_maps__[0]
			var maps = Seq.Arity1IQ(f_maps__[1])
			_, _ = f, maps
			if Truth_(Some.X_invoke_Arity2(Identity, maps)) {
				{
					var merge_entry = func(G__6391 *AFn) *AFn {
						return Fn(G__6391, 2, func(m interface{}, e interface{}) interface{} {
							{
								var k = First.X_invoke_Arity1(e)
								var v = Second.X_invoke_Arity1(e)
								_, _ = k, v
								if Contains_QMARK_.Arity2IIB(m, k) {
									return Assoc.X_invoke_Arity3(m, k, func() interface{} {
										var G__6389 = Get.X_invoke_Arity2(m, k)
										var G__6390 = v
										_, _ = G__6389, G__6390
										return f.(CljsCoreIFn).X_invoke_Arity2(G__6389, G__6390)
									}())
								} else {
									return Assoc.X_invoke_Arity3(m, k, v)
								}
							}
						})
					}(&AFn{})
					var merge2 = func(G__6392 *AFn, merge_entry CljsCoreIFn) *AFn {
						return Fn(G__6392, 2, func(m1 interface{}, m2 interface{}) interface{} {
							return Reduce.X_invoke_Arity3(merge_entry, func() interface{} {
								var or__171__auto__ = m1
								_ = or__171__auto__
								if Truth_(or__171__auto__) {
									return or__171__auto__
								} else {
									return CljsCorePersistentArrayMap_EMPTY
								}
							}(), Seq.Arity1IQ(m2))
						})
					}(&AFn{}, merge_entry)
					_, _ = merge_entry, merge2
					return Reduce.X_invoke_Arity2(merge2, maps)
				}
			} else {
				return nil
			}
		})
	}(&AFn{})

	Select_keys = func(select_keys *AFn) *AFn {
		return Fn(select_keys, 2, func(map_ interface{}, keyseq interface{}) interface{} {
			{
				var ret interface{} = CljsCorePersistentArrayMap_EMPTY
				var keys interface{} = Seq.Arity1IQ(keyseq)
				_, _ = ret, keys
				for {
					if Truth_(keys) {
						{
							var key = First.X_invoke_Arity1(keys)
							var entry = Get.X_invoke_Arity3(map_, key, (&CljsCoreKeyword{Ns: "cljs.core", Name: "not-found", Fqn: "cljs.core/not-found", X_hash: float64(-1572889185)}))
							_, _ = key, entry
							ret, keys = func() interface{} {
								if Not_EQ_.Arity2IIB(entry, (&CljsCoreKeyword{Ns: "cljs.core", Name: "not-found", Fqn: "cljs.core/not-found", X_hash: float64(-1572889185)})) {
									return Assoc.X_invoke_Arity3(ret, key, entry)
								} else {
									return ret
								}
							}(), Next.Arity1IQ(keys)
							continue
						}
					} else {
						return ret
					}
				}
			}
		})
	}(&AFn{})

	X__GT_PersistentHashSet = func(__GT_PersistentHashSet *AFn) *AFn {
		return Fn(__GT_PersistentHashSet, 3, func(meta interface{}, hash_map interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentHashSet{meta, hash_map, __hash})
		})
	}(&AFn{})

	X__GT_TransientHashSet = func(__GT_TransientHashSet *AFn) *AFn {
		return Fn(__GT_TransientHashSet, 1, func(transient_map interface{}) interface{} {
			return (&CljsCoreTransientHashSet{transient_map})
		})
	}(&AFn{})

	X__GT_PersistentTreeSet = func(__GT_PersistentTreeSet *AFn) *AFn {
		return Fn(__GT_PersistentTreeSet, 3, func(meta interface{}, tree_map interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentTreeSet{meta, tree_map, __hash})
		})
	}(&AFn{})

	Set_from_indexed_seq = func(set_from_indexed_seq *AFn) *AFn {
		return Fn(set_from_indexed_seq, 1, func(iseq interface{}) interface{} {
			{
				var arr = Native_get_instance_field.X_invoke_Arity2(iseq, "Arr")
				var ret = func() interface{} {
					var a__1048__auto__ = arr
					_ = a__1048__auto__
					{
						var i = float64(0)
						var res interface{} = CljsCorePersistentHashSet_EMPTY.X_as_transient_Arity1()
						_, _ = i, res
						for {
							if i < Alength_(a__1048__auto__) {
								i, res = (i + float64(1)), res.(CljsCoreITransientCollection).X_conj_BANG__Arity2(Aget_(arr, i))
								continue
							} else {
								return res
							}
						}
					}
				}()
				_, _ = arr, ret
				return ret.(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
			}
		})
	}(&AFn{})

	Set = func(set *AFn) *AFn {
		return Fn(set, 1, func(coll interface{}) interface{} {
			{
				var in = Seq.Arity1IQ(coll)
				_ = in
				if Nil_(in) {
					return CljsCorePersistentHashSet_EMPTY
				} else {
					if (Value_(in).Type().AssignableTo(reflect.TypeOf((**CljsCoreIndexedSeq)(nil)).Elem())) && (Native_get_instance_field.X_invoke_Arity2(in, "I").(float64) == float64(0)) {
						return Set_from_indexed_seq.X_invoke_Arity1(in)
					} else {
						{
							var in___1 interface{} = in
							var out interface{} = CljsCorePersistentHashSet_EMPTY.X_as_transient_Arity1()
							_, _ = in___1, out
							for {
								if !(Nil_(in___1)) {
									in___1, out = in___1.(CljsCoreINext).X_next_Arity1(), out.(CljsCoreITransientCollection).X_conj_BANG__Arity2(in___1.(CljsCoreISeq).X_first_Arity1())
									continue
								} else {
									return out.(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
								}
							}
						}

					}
				}
			}
		})
	}(&AFn{})

	Hash_set = func(hash_set *AFn) *AFn {
		return Fn(hash_set, 0, func() interface{} {
			return CljsCorePersistentHashSet_EMPTY
		}, func(keys__ ...interface{}) interface{} {
			var keys = Seq.Arity1IQ(keys__[0])
			_ = keys
			return Set.X_invoke_Arity1(keys)
		})
	}(&AFn{})

	Sorted_set = func(sorted_set *AFn) *AFn {
		return Fn(sorted_set, 0, func(keys__ ...interface{}) interface{} {
			var keys = Seq.Arity1IQ(keys__[0])
			_ = keys
			return Reduce.X_invoke_Arity3(X_conj, CljsCorePersistentTreeSet_EMPTY, keys)
		})
	}(&AFn{})

	Sorted_set_by = func(sorted_set_by *AFn) *AFn {
		return Fn(sorted_set_by, 1, func(comparator_keys__ ...interface{}) interface{} {
			var comparator = comparator_keys__[0]
			var keys = Seq.Arity1IQ(comparator_keys__[1])
			_, _ = comparator, keys
			return Reduce.X_invoke_Arity3(X_conj, (&CljsCorePersistentTreeSet{nil, Sorted_map_by.X_invoke_ArityVariadic(comparator, Array_seq.X_invoke_Arity1([]interface{}{})).(*CljsCorePersistentTreeMap), float64(0)}), keys)
		})
	}(&AFn{})

	Replace = func(replace *AFn) *AFn {
		return Fn(replace, 2, func(smap interface{}) interface{} {
			return Map_.X_invoke_Arity1(func(G__6458 *AFn) *AFn {
				return Fn(G__6458, 1, func(p1__6456_SHARP_ interface{}) interface{} {
					{
						var temp__4220__auto__ = Find.X_invoke_Arity2(smap, p1__6456_SHARP_)
						_ = temp__4220__auto__
						if Truth_(temp__4220__auto__) {
							{
								var e = temp__4220__auto__
								_ = e
								return Val.X_invoke_Arity1(e)
							}
						} else {
							return p1__6456_SHARP_
						}
					}
				})
			}(&AFn{})).(CljsCoreIFn)
		}, func(smap interface{}, coll interface{}) interface{} {
			if Vector_QMARK_.Arity1IB(coll) {
				{
					var n = Count.X_invoke_Arity1(coll).(float64)
					_ = n
					return Reduce.X_invoke_Arity3(func(G__6459 *AFn, n float64) *AFn {
						return Fn(G__6459, 2, func(v interface{}, i interface{}) interface{} {
							{
								var temp__4220__auto__ = Find.X_invoke_Arity2(smap, Nth.X_invoke_Arity2(v, i))
								_ = temp__4220__auto__
								if Truth_(temp__4220__auto__) {
									{
										var e = temp__4220__auto__
										_ = e
										return Assoc.X_invoke_Arity3(v, i, Second.X_invoke_Arity1(e))
									}
								} else {
									return v
								}
							}
						})
					}(&AFn{}, n), coll, Take.X_invoke_Arity2(n, Iterate.X_invoke_Arity2(Inc, float64(0)).(*CljsCoreCons)).(*CljsCoreLazySeq))
				}
			} else {
				return Map_.X_invoke_Arity2(func(G__6460 *AFn) *AFn {
					return Fn(G__6460, 1, func(p1__6457_SHARP_ interface{}) interface{} {
						{
							var temp__4220__auto__ = Find.X_invoke_Arity2(smap, p1__6457_SHARP_)
							_ = temp__4220__auto__
							if Truth_(temp__4220__auto__) {
								{
									var e = temp__4220__auto__
									_ = e
									return Second.X_invoke_Arity1(e)
								}
							} else {
								return p1__6457_SHARP_
							}
						}
					})
				}(&AFn{}), coll).(*CljsCoreLazySeq)
			}
		})
	}(&AFn{})

	Distinct = func(distinct *AFn) *AFn {
		return Fn(distinct, 1, func(coll interface{}) interface{} {
			{
				var step = func(step *AFn) *AFn {
					return Fn(step, 2, func(xs interface{}, seen interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__6473 *AFn) *AFn {
							return Fn(G__6473, 0, func() interface{} {
								return func(G__6474 *AFn) *AFn {
									return Fn(G__6474, 2, func(p__6471 interface{}, seen___1 interface{}) interface{} {
										for {
											{
												var vec__6472 = p__6471
												var f = Nth.X_invoke_Arity3(vec__6472, float64(0), nil)
												var xs___1 = vec__6472
												_, _, _ = vec__6472, f, xs___1
												{
													var temp__4222__auto__ = Seq.Arity1IQ(xs___1)
													_ = temp__4222__auto__
													if Truth_(temp__4222__auto__) {
														{
															var s = temp__4222__auto__
															_ = s
															if Contains_QMARK_.Arity2IIB(seen___1, f) {
																p__6471, seen___1 = Rest.Arity1IQ(s), seen___1
																continue
															} else {
																return Cons.X_invoke_Arity2(f, step.X_invoke_Arity2(Rest.Arity1IQ(s), Conj.X_invoke_Arity2(seen___1, f)).(*CljsCoreLazySeq)).(*CljsCoreCons)
															}
														}
													} else {
														return nil
													}
												}
											}
										}
									})
								}(&AFn{}).X_invoke_Arity2(xs, seen)
							})
						}(&AFn{}), nil, nil})
					})
				}(&AFn{})
				_ = step
				return step.X_invoke_Arity2(coll, CljsCorePersistentHashSet_EMPTY).(*CljsCoreLazySeq)
			}
		})
	}(&AFn{})

	Butlast = func(butlast *AFn) *AFn {
		return Fn(butlast, 1, func(s interface{}) interface{} {
			{
				var ret interface{} = CljsCorePersistentVector_EMPTY
				var s___1 interface{} = s
				_, _ = ret, s___1
				for {
					if Truth_(Next.Arity1IQ(s___1)) {
						ret, s___1 = Conj.X_invoke_Arity2(ret, First.X_invoke_Arity1(s___1)), Next.Arity1IQ(s___1)
						continue
					} else {
						return Seq.Arity1IQ(ret)
					}
				}
			}
		})
	}(&AFn{})

	Name = func(name *AFn) *AFn {
		return Fn(name, 1, func(x interface{}) interface{} {
			if Value_(x).Type().Implements(reflect.TypeOf((*CljsCoreINamed)(nil)).Elem()) {
				return x.(CljsCoreINamed).X_name_Arity1()
			} else {
				if Value_(x).Kind() == reflect.String {
					return x
				} else {
					panic((&js.Error{("Doesn't support name: " + Str.X_invoke_Arity1(x).(string))}))
				}
			}
		})
	}(&AFn{})

	Zipmap = func(zipmap *AFn) *AFn {
		return Fn(zipmap, 2, func(keys interface{}, vals interface{}) interface{} {
			{
				var map_ interface{} = Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY)
				var ks interface{} = Seq.Arity1IQ(keys)
				var vs interface{} = Seq.Arity1IQ(vals)
				_, _, _ = map_, ks, vs
				for {
					if Truth_(func() interface{} {
						var and__159__auto__ = ks
						_ = and__159__auto__
						if Truth_(and__159__auto__) {
							return vs
						} else {
							return and__159__auto__
						}
					}()) {
						map_, ks, vs = Assoc_BANG_.X_invoke_Arity3(map_, First.X_invoke_Arity1(ks), First.X_invoke_Arity1(vs)), Next.Arity1IQ(ks), Next.Arity1IQ(vs)
						continue
					} else {
						return Persistent_BANG_.X_invoke_Arity1(map_)
					}
				}
			}
		})
	}(&AFn{})

	Max_key = func(max_key *AFn) *AFn {
		return Fn(max_key, 3, func(k interface{}, x interface{}) interface{} {
			return x
		}, func(k interface{}, x interface{}, y interface{}) interface{} {
			if func() interface{} {
				var G__6485 = x
				_ = G__6485
				return k.(CljsCoreIFn).X_invoke_Arity1(G__6485)
			}().(float64) > func() interface{} {
				var G__6486 = y
				_ = G__6486
				return k.(CljsCoreIFn).X_invoke_Arity1(G__6486)
			}().(float64) {
				return x
			} else {
				return y
			}
		}, func(k_x_y_more__ ...interface{}) interface{} {
			var k = k_x_y_more__[0]
			var x = k_x_y_more__[1]
			var y = k_x_y_more__[2]
			var more = Seq.Arity1IQ(k_x_y_more__[3])
			_, _, _, _ = k, x, y, more
			return Reduce.X_invoke_Arity3(func(G__6487 *AFn) *AFn {
				return Fn(G__6487, 2, func(p1__6475_SHARP_ interface{}, p2__6476_SHARP_ interface{}) interface{} {
					return max_key.X_invoke_Arity3(k, p1__6475_SHARP_, p2__6476_SHARP_)
				})
			}(&AFn{}), max_key.X_invoke_Arity3(k, x, y), more)
		})
	}(&AFn{})

	Min_key = func(min_key *AFn) *AFn {
		return Fn(min_key, 3, func(k interface{}, x interface{}) interface{} {
			return x
		}, func(k interface{}, x interface{}, y interface{}) interface{} {
			if func() interface{} {
				var G__6498 = x
				_ = G__6498
				return k.(CljsCoreIFn).X_invoke_Arity1(G__6498)
			}().(float64) < func() interface{} {
				var G__6499 = y
				_ = G__6499
				return k.(CljsCoreIFn).X_invoke_Arity1(G__6499)
			}().(float64) {
				return x
			} else {
				return y
			}
		}, func(k_x_y_more__ ...interface{}) interface{} {
			var k = k_x_y_more__[0]
			var x = k_x_y_more__[1]
			var y = k_x_y_more__[2]
			var more = Seq.Arity1IQ(k_x_y_more__[3])
			_, _, _, _ = k, x, y, more
			return Reduce.X_invoke_Arity3(func(G__6500 *AFn) *AFn {
				return Fn(G__6500, 2, func(p1__6488_SHARP_ interface{}, p2__6489_SHARP_ interface{}) interface{} {
					return min_key.X_invoke_Arity3(k, p1__6488_SHARP_, p2__6489_SHARP_)
				})
			}(&AFn{}), min_key.X_invoke_Arity3(k, x, y), more)
		})
	}(&AFn{})

	X__GT_ArrayList = func(__GT_ArrayList *AFn) *AFn {
		return Fn(__GT_ArrayList, 1, func(arr interface{}) interface{} {
			return (&CljsCoreArrayList{arr})
		})
	}(&AFn{})

	Array_list = func(array_list *AFn) *AFn {
		return Fn(array_list, 0, func() interface{} {
			return (&CljsCoreArrayList{[]interface{}{}})
		})
	}(&AFn{})

	Partition_all = func(partition_all *AFn) *AFn {
		return Fn(partition_all, 3, func(n interface{}) interface{} {
			return func(G__6517 *AFn) *AFn {
				return Fn(G__6517, 1, func(f1 interface{}) interface{} {
					{
						var a = Array_list.X_invoke_Arity0().(*CljsCoreArrayList)
						_ = a
						return func(G__6518 *AFn, a *CljsCoreArrayList) *AFn {
							return Fn(G__6518, 2, func() interface{} {
								{
									return f1.(CljsCoreIFn).X_invoke_Arity0()
								}
							}, func(result interface{}) interface{} {
								{
									var result___1 = func() interface{} {
										if Truth_(a.IsEmpty()) {
											return result
										} else {
											return func() interface{} {
												var v = Vec.X_invoke_Arity1(a.ToArray())
												_ = v
												a.Clear()
												{
													var G__6512 = result
													var G__6513 = v
													_, _ = G__6512, G__6513
													return f1.(CljsCoreIFn).X_invoke_Arity2(G__6512, G__6513)
												}
											}()
										}
									}()
									_ = result___1
									{
										var G__6514 = result___1
										_ = G__6514
										return f1.(CljsCoreIFn).X_invoke_Arity1(G__6514)
									}
								}
							}, func(result interface{}, input interface{}) interface{} {
								a.Add(input)
								if n.(float64) == a.Size().(float64) {
									{
										var v = Vec.X_invoke_Arity1(a.ToArray())
										_ = v
										a.Clear()
										{
											var G__6515 = result
											var G__6516 = v
											_, _ = G__6515, G__6516
											return f1.(CljsCoreIFn).X_invoke_Arity2(G__6515, G__6516)
										}
									}
								} else {
									return result
								}
							})
						}(&AFn{}, a)
					}
				})
			}(&AFn{})
		}, func(n interface{}, coll interface{}) interface{} {
			return partition_all.X_invoke_Arity3(n, n, coll).(*CljsCoreLazySeq)
		}, func(n interface{}, step interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__6519 *AFn) *AFn {
				return Fn(G__6519, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								return Cons.X_invoke_Arity2(Take.X_invoke_Arity2(n, s).(*CljsCoreLazySeq), partition_all.X_invoke_Arity3(n, step, Drop.X_invoke_Arity2(step, s).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)).(*CljsCoreCons)
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		})
	}(&AFn{})

	Take_while = func(take_while *AFn) *AFn {
		return Fn(take_while, 2, func(pred interface{}) interface{} {
			return func(G__6532 *AFn) *AFn {
				return Fn(G__6532, 1, func(f1 interface{}) interface{} {
					return func(G__6533 *AFn) *AFn {
						return Fn(G__6533, 2, func() interface{} {
							{
								return f1.(CljsCoreIFn).X_invoke_Arity0()
							}
						}, func(result interface{}) interface{} {
							{
								var G__6527 = result
								_ = G__6527
								return f1.(CljsCoreIFn).X_invoke_Arity1(G__6527)
							}
						}, func(result interface{}, input interface{}) interface{} {
							if Truth_(func() interface{} {
								var G__6528 = input
								_ = G__6528
								return pred.(CljsCoreIFn).X_invoke_Arity1(G__6528)
							}()) {
								{
									var G__6529 = result
									var G__6530 = input
									_, _ = G__6529, G__6530
									return f1.(CljsCoreIFn).X_invoke_Arity2(G__6529, G__6530)
								}
							} else {
								return Reduced.X_invoke_Arity1(result).(*CljsCoreReduced)
							}
						})
					}(&AFn{})
				})
			}(&AFn{})
		}, func(pred interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__6534 *AFn) *AFn {
				return Fn(G__6534, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								if Truth_(func() interface{} {
									var G__6531 = First.X_invoke_Arity1(s)
									_ = G__6531
									return pred.(CljsCoreIFn).X_invoke_Arity1(G__6531)
								}()) {
									return Cons.X_invoke_Arity2(First.X_invoke_Arity1(s), take_while.X_invoke_Arity2(pred, Rest.Arity1IQ(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
								} else {
									return nil
								}
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		})
	}(&AFn{})

	Mk_bound_fn = func(mk_bound_fn *AFn) *AFn {
		return Fn(mk_bound_fn, 3, func(sc interface{}, test interface{}, key interface{}) interface{} {
			return func(G__6543 *AFn) *AFn {
				return Fn(G__6543, 1, func(e interface{}) interface{} {
					{
						var comp = sc.(CljsCoreISorted).X_comparator_Arity1()
						_ = comp
						{
							var G__6539 = func() interface{} {
								var G__6541 = sc.(CljsCoreISorted).X_entry_key_Arity2(e)
								var G__6542 = key
								_, _ = G__6541, G__6542
								return comp.(CljsCoreIFn).X_invoke_Arity2(G__6541, G__6542)
							}()
							var G__6540 = float64(0)
							_, _ = G__6539, G__6540
							return test.(CljsCoreIFn).X_invoke_Arity2(G__6539, G__6540)
						}
					}
				})
			}(&AFn{})
		})
	}(&AFn{})

	Subseq = func(subseq *AFn) *AFn {
		return Fn(subseq, 5, func(sc interface{}, test interface{}, key interface{}) interface{} {
			{
				var include = Mk_bound_fn.X_invoke_Arity3(sc, test, key).(CljsCoreIFn)
				_ = include
				if Truth_(CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{X_GT_, X_GT__EQ_}, true).(*CljsCorePersistentHashSet).X_invoke_Arity1(test)) {
					{
						var temp__4222__auto__ = sc.(CljsCoreISorted).X_sorted_seq_from_Arity3(key, true)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var vec__6547 = temp__4222__auto__
								var e = Nth.X_invoke_Arity3(vec__6547, float64(0), nil)
								var s = vec__6547
								_, _, _ = vec__6547, e, s
								if Truth_(func() interface{} {
									var G__6548 = e
									_ = G__6548
									return include.X_invoke_Arity1(G__6548)
								}()) {
									return s
								} else {
									return Next.Arity1IQ(s)
								}
							}
						} else {
							return nil
						}
					}
				} else {
					return Take_while.X_invoke_Arity2(include, sc.(CljsCoreISorted).X_sorted_seq_Arity2(true)).(*CljsCoreLazySeq)
				}
			}
		}, func(sc interface{}, start_test interface{}, start_key interface{}, end_test interface{}, end_key interface{}) interface{} {
			{
				var temp__4222__auto__ = sc.(CljsCoreISorted).X_sorted_seq_from_Arity3(start_key, true)
				_ = temp__4222__auto__
				if Truth_(temp__4222__auto__) {
					{
						var vec__6549 = temp__4222__auto__
						var e = Nth.X_invoke_Arity3(vec__6549, float64(0), nil)
						var s = vec__6549
						_, _, _ = vec__6549, e, s
						return Take_while.X_invoke_Arity2(Mk_bound_fn.X_invoke_Arity3(sc, end_test, end_key).(CljsCoreIFn), func() interface{} {
							if Truth_(Mk_bound_fn.X_invoke_Arity3(sc, start_test, start_key).(CljsCoreIFn).(CljsCoreIFn).X_invoke_Arity1(e)) {
								return s
							} else {
								return Next.Arity1IQ(s)
							}
						}()).(*CljsCoreLazySeq)
					}
				} else {
					return nil
				}
			}
		})
	}(&AFn{})

	Rsubseq = func(rsubseq *AFn) *AFn {
		return Fn(rsubseq, 5, func(sc interface{}, test interface{}, key interface{}) interface{} {
			{
				var include = Mk_bound_fn.X_invoke_Arity3(sc, test, key).(CljsCoreIFn)
				_ = include
				if Truth_(CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{X_LT_, X_LT__EQ_}, true).(*CljsCorePersistentHashSet).X_invoke_Arity1(test)) {
					{
						var temp__4222__auto__ = sc.(CljsCoreISorted).X_sorted_seq_from_Arity3(key, false)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var vec__6553 = temp__4222__auto__
								var e = Nth.X_invoke_Arity3(vec__6553, float64(0), nil)
								var s = vec__6553
								_, _, _ = vec__6553, e, s
								if Truth_(func() interface{} {
									var G__6554 = e
									_ = G__6554
									return include.X_invoke_Arity1(G__6554)
								}()) {
									return s
								} else {
									return Next.Arity1IQ(s)
								}
							}
						} else {
							return nil
						}
					}
				} else {
					return Take_while.X_invoke_Arity2(include, sc.(CljsCoreISorted).X_sorted_seq_Arity2(false)).(*CljsCoreLazySeq)
				}
			}
		}, func(sc interface{}, start_test interface{}, start_key interface{}, end_test interface{}, end_key interface{}) interface{} {
			{
				var temp__4222__auto__ = sc.(CljsCoreISorted).X_sorted_seq_from_Arity3(end_key, false)
				_ = temp__4222__auto__
				if Truth_(temp__4222__auto__) {
					{
						var vec__6555 = temp__4222__auto__
						var e = Nth.X_invoke_Arity3(vec__6555, float64(0), nil)
						var s = vec__6555
						_, _, _ = vec__6555, e, s
						return Take_while.X_invoke_Arity2(Mk_bound_fn.X_invoke_Arity3(sc, start_test, start_key).(CljsCoreIFn), func() interface{} {
							if Truth_(Mk_bound_fn.X_invoke_Arity3(sc, end_test, end_key).(CljsCoreIFn).(CljsCoreIFn).X_invoke_Arity1(e)) {
								return s
							} else {
								return Next.Arity1IQ(s)
							}
						}()).(*CljsCoreLazySeq)
					}
				} else {
					return nil
				}
			}
		})
	}(&AFn{})

	X__GT_Range = func(__GT_Range *AFn) *AFn {
		return Fn(__GT_Range, 5, func(meta interface{}, start interface{}, end interface{}, step interface{}, __hash interface{}) interface{} {
			return (&CljsCoreRange{meta, start, end, step, __hash})
		})
	}(&AFn{})

	Range_ = func(range_ *AFn) *AFn {
		return Fn(range_, 3, func() interface{} {
			return range_.X_invoke_Arity3(float64(0), Native_get_instance_field.X_invoke_Arity2(js.Number, "MAX_VALUE"), float64(1)).(*CljsCoreRange)
		}, func(end interface{}) interface{} {
			return range_.X_invoke_Arity3(float64(0), end, float64(1)).(*CljsCoreRange)
		}, func(start interface{}, end interface{}) interface{} {
			return range_.X_invoke_Arity3(start, end, float64(1)).(*CljsCoreRange)
		}, func(start interface{}, end interface{}, step interface{}) interface{} {
			return (&CljsCoreRange{nil, start, end, step, nil})
		})
	}(&AFn{})

	Take_nth = func(take_nth *AFn) *AFn {
		return Fn(take_nth, 2, func(n interface{}) interface{} {
			return func(G__6575 *AFn) *AFn {
				return Fn(G__6575, 1, func(f1 interface{}) interface{} {
					{
						var ia = Atom.X_invoke_Arity1(float64(-1)).(*CljsCoreAtom)
						_ = ia
						return func(G__6576 *AFn, ia *CljsCoreAtom) *AFn {
							return Fn(G__6576, 2, func() interface{} {
								{
									return f1.(CljsCoreIFn).X_invoke_Arity0()
								}
							}, func(result interface{}) interface{} {
								{
									var G__6572 = result
									_ = G__6572
									return f1.(CljsCoreIFn).X_invoke_Arity1(G__6572)
								}
							}, func(result interface{}, input interface{}) interface{} {
								{
									var i = Swap_BANG_.X_invoke_Arity2(ia, Inc)
									_ = i
									if Rem.X_invoke_Arity2(i, n).(float64) == float64(0) {
										{
											var G__6573 = result
											var G__6574 = input
											_, _ = G__6573, G__6574
											return f1.(CljsCoreIFn).X_invoke_Arity2(G__6573, G__6574)
										}
									} else {
										return result
									}
								}
							})
						}(&AFn{}, ia)
					}
				})
			}(&AFn{})
		}, func(n interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__6577 *AFn) *AFn {
				return Fn(G__6577, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								return Cons.X_invoke_Arity2(First.X_invoke_Arity1(s), take_nth.X_invoke_Arity2(n, Drop.X_invoke_Arity2(n, s).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)).(*CljsCoreCons)
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		})
	}(&AFn{})

	Split_with = func(split_with *AFn) *AFn {
		return Fn(split_with, 2, func(pred interface{}, coll interface{}) interface{} {
			return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Take_while.X_invoke_Arity2(pred, coll).(*CljsCoreLazySeq), Drop_while.X_invoke_Arity2(pred, coll).(*CljsCoreLazySeq)}, nil})
		})
	}(&AFn{})

	Partition_by = func(partition_by *AFn) *AFn {
		return Fn(partition_by, 2, func(f interface{}) interface{} {
			return func(G__6597 *AFn) *AFn {
				return Fn(G__6597, 1, func(f1 interface{}) interface{} {
					{
						var a = Array_list.X_invoke_Arity0().(*CljsCoreArrayList)
						var pa = Atom.X_invoke_Arity1((&CljsCoreKeyword{Ns: "cljs.core", Name: "none", Fqn: "cljs.core/none", X_hash: float64(926646439)})).(*CljsCoreAtom)
						_, _ = a, pa
						return func(G__6598 *AFn, a *CljsCoreArrayList, pa *CljsCoreAtom) *AFn {
							return Fn(G__6598, 2, func() interface{} {
								{
									return f1.(CljsCoreIFn).X_invoke_Arity0()
								}
							}, func(result interface{}) interface{} {
								{
									var result___1 = func() interface{} {
										if Truth_(a.IsEmpty()) {
											return result
										} else {
											return func() interface{} {
												var v = Vec.X_invoke_Arity1(a.ToArray())
												_ = v
												a.Clear()
												{
													var G__6589 = result
													var G__6590 = v
													_, _ = G__6589, G__6590
													return f1.(CljsCoreIFn).X_invoke_Arity2(G__6589, G__6590)
												}
											}()
										}
									}()
									_ = result___1
									{
										var G__6591 = result___1
										_ = G__6591
										return f1.(CljsCoreIFn).X_invoke_Arity1(G__6591)
									}
								}
							}, func(result interface{}, input interface{}) interface{} {
								{
									var pval = Deref.X_invoke_Arity1(pa)
									var val = func() interface{} {
										var G__6592 = input
										_ = G__6592
										return f.(CljsCoreIFn).X_invoke_Arity1(G__6592)
									}()
									_, _ = pval, val
									Reset_BANG_.X_invoke_Arity2(pa, val)
									if (Keyword_identical_QMARK_.Arity2IIB(pval, (&CljsCoreKeyword{Ns: "cljs.core", Name: "none", Fqn: "cljs.core/none", X_hash: float64(926646439)}))) || (X_EQ_.Arity2IIB(val, pval)) {
										a.Add(input)
										return result
									} else {
										{
											var v = Vec.X_invoke_Arity1(a.ToArray())
											_ = v
											a.Clear()
											{
												var ret = func() interface{} {
													var G__6593 = result
													var G__6594 = v
													_, _ = G__6593, G__6594
													return f1.(CljsCoreIFn).X_invoke_Arity2(G__6593, G__6594)
												}()
												_ = ret
												if Reduced_QMARK_.Arity1IB(ret) {
												} else {
													a.Add(input)
												}
												return ret
											}
										}
									}
								}
							})
						}(&AFn{}, a, pa)
					}
				})
			}(&AFn{})
		}, func(f interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__6599 *AFn) *AFn {
				return Fn(G__6599, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								{
									var fst = First.X_invoke_Arity1(s)
									var fv = func() interface{} {
										var G__6595 = fst
										_ = G__6595
										return f.(CljsCoreIFn).X_invoke_Arity1(G__6595)
									}()
									var run = Cons.X_invoke_Arity2(fst, Take_while.X_invoke_Arity2(func(G__6600 *AFn, fst interface{}, fv interface{}, s CljsCoreISeq, temp__4222__auto__ CljsCoreISeq) *AFn {
										return Fn(G__6600, 1, func(p1__6578_SHARP_ interface{}) interface{} {
											return X_EQ_.Arity2IIB(fv, func() interface{} {
												var G__6596 = p1__6578_SHARP_
												_ = G__6596
												return f.(CljsCoreIFn).X_invoke_Arity1(G__6596)
											}())
										})
									}(&AFn{}, fst, fv, s, temp__4222__auto__), Next.Arity1IQ(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
									_, _, _ = fst, fv, run
									return Cons.X_invoke_Arity2(run, partition_by.X_invoke_Arity2(f, Seq.Arity1IQ(Drop.X_invoke_Arity2(Count.X_invoke_Arity1(run).(float64), s).(*CljsCoreLazySeq))).(*CljsCoreLazySeq)).(*CljsCoreCons)
								}
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		})
	}(&AFn{})

	Frequencies = func(frequencies *AFn) *AFn {
		return Fn(frequencies, 1, func(coll interface{}) interface{} {
			return Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(func(G__6601 *AFn) *AFn {
				return Fn(G__6601, 2, func(counts interface{}, x interface{}) interface{} {
					return Assoc_BANG_.X_invoke_Arity3(counts, x, (Get.X_invoke_Arity3(counts, x, float64(0)).(float64) + float64(1)))
				})
			}(&AFn{}), Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY), coll))
		})
	}(&AFn{})

	Reductions = func(reductions *AFn) *AFn {
		return Fn(reductions, 3, func(f interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__6612 *AFn) *AFn {
				return Fn(G__6612, 0, func() interface{} {
					{
						var temp__4220__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4220__auto__
						if Truth_(temp__4220__auto__) {
							{
								var s = temp__4220__auto__
								_ = s
								return reductions.X_invoke_Arity3(f, First.X_invoke_Arity1(s), Rest.Arity1IQ(s)).(*CljsCoreCons)
							}
						} else {
							return CljsCoreList_EMPTY.X_conj_Arity2(func() interface{} {
								return f.(CljsCoreIFn).X_invoke_Arity0()
							}())
						}
					}
				})
			}(&AFn{}), nil, nil})
		}, func(f interface{}, init interface{}, coll interface{}) interface{} {
			return Cons.X_invoke_Arity2(init, (&CljsCoreLazySeq{nil, func(G__6613 *AFn) *AFn {
				return Fn(G__6613, 0, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								return reductions.X_invoke_Arity3(f, func() interface{} {
									var G__6610 = init
									var G__6611 = First.X_invoke_Arity1(s)
									_, _ = G__6610, G__6611
									return f.(CljsCoreIFn).X_invoke_Arity2(G__6610, G__6611)
								}(), Rest.Arity1IQ(s)).(*CljsCoreCons)
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})).(*CljsCoreCons)
		})
	}(&AFn{})

	Juxt = func(juxt *AFn) *AFn {
		return Fn(juxt, 3, func(f interface{}) interface{} {
			return func(G__6708 *AFn) *AFn {
				return Fn(G__6708, 3, func() interface{} {
					return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						return f.(CljsCoreIFn).X_invoke_Arity0()
					}()}, nil})
				}, func(x interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						var G__6666 = x
						_ = G__6666
						return f.(CljsCoreIFn).X_invoke_Arity1(G__6666)
					}()}, nil})
				}, func(x interface{}, y interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						var G__6667 = x
						var G__6668 = y
						_, _ = G__6667, G__6668
						return f.(CljsCoreIFn).X_invoke_Arity2(G__6667, G__6668)
					}()}, nil})
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						var G__6669 = x
						var G__6670 = y
						var G__6671 = z
						_, _, _ = G__6669, G__6670, G__6671
						return f.(CljsCoreIFn).X_invoke_Arity3(G__6669, G__6670, G__6671)
					}()}, nil})
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Apply.X_invoke_Arity5(f, x, y, z, args)}, nil})
				})
			}(&AFn{})
		}, func(f interface{}, g interface{}) interface{} {
			return func(G__6709 *AFn) *AFn {
				return Fn(G__6709, 3, func() interface{} {
					return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						return f.(CljsCoreIFn).X_invoke_Arity0()
					}(), func() interface{} {
						return g.(CljsCoreIFn).X_invoke_Arity0()
					}()}, nil})
				}, func(x interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						var G__6672 = x
						_ = G__6672
						return f.(CljsCoreIFn).X_invoke_Arity1(G__6672)
					}(), func() interface{} {
						var G__6673 = x
						_ = G__6673
						return g.(CljsCoreIFn).X_invoke_Arity1(G__6673)
					}()}, nil})
				}, func(x interface{}, y interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						var G__6674 = x
						var G__6675 = y
						_, _ = G__6674, G__6675
						return f.(CljsCoreIFn).X_invoke_Arity2(G__6674, G__6675)
					}(), func() interface{} {
						var G__6676 = x
						var G__6677 = y
						_, _ = G__6676, G__6677
						return g.(CljsCoreIFn).X_invoke_Arity2(G__6676, G__6677)
					}()}, nil})
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						var G__6678 = x
						var G__6679 = y
						var G__6680 = z
						_, _, _ = G__6678, G__6679, G__6680
						return f.(CljsCoreIFn).X_invoke_Arity3(G__6678, G__6679, G__6680)
					}(), func() interface{} {
						var G__6681 = x
						var G__6682 = y
						var G__6683 = z
						_, _, _ = G__6681, G__6682, G__6683
						return g.(CljsCoreIFn).X_invoke_Arity3(G__6681, G__6682, G__6683)
					}()}, nil})
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Apply.X_invoke_Arity5(f, x, y, z, args), Apply.X_invoke_Arity5(g, x, y, z, args)}, nil})
				})
			}(&AFn{})
		}, func(f interface{}, g interface{}, h interface{}) interface{} {
			return func(G__6710 *AFn) *AFn {
				return Fn(G__6710, 3, func() interface{} {
					return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						return f.(CljsCoreIFn).X_invoke_Arity0()
					}(), func() interface{} {
						return g.(CljsCoreIFn).X_invoke_Arity0()
					}(), func() interface{} {
						return h.(CljsCoreIFn).X_invoke_Arity0()
					}()}, nil})
				}, func(x interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						var G__6684 = x
						_ = G__6684
						return f.(CljsCoreIFn).X_invoke_Arity1(G__6684)
					}(), func() interface{} {
						var G__6685 = x
						_ = G__6685
						return g.(CljsCoreIFn).X_invoke_Arity1(G__6685)
					}(), func() interface{} {
						var G__6686 = x
						_ = G__6686
						return h.(CljsCoreIFn).X_invoke_Arity1(G__6686)
					}()}, nil})
				}, func(x interface{}, y interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						var G__6687 = x
						var G__6688 = y
						_, _ = G__6687, G__6688
						return f.(CljsCoreIFn).X_invoke_Arity2(G__6687, G__6688)
					}(), func() interface{} {
						var G__6689 = x
						var G__6690 = y
						_, _ = G__6689, G__6690
						return g.(CljsCoreIFn).X_invoke_Arity2(G__6689, G__6690)
					}(), func() interface{} {
						var G__6691 = x
						var G__6692 = y
						_, _ = G__6691, G__6692
						return h.(CljsCoreIFn).X_invoke_Arity2(G__6691, G__6692)
					}()}, nil})
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{func() interface{} {
						var G__6693 = x
						var G__6694 = y
						var G__6695 = z
						_, _, _ = G__6693, G__6694, G__6695
						return f.(CljsCoreIFn).X_invoke_Arity3(G__6693, G__6694, G__6695)
					}(), func() interface{} {
						var G__6696 = x
						var G__6697 = y
						var G__6698 = z
						_, _, _ = G__6696, G__6697, G__6698
						return g.(CljsCoreIFn).X_invoke_Arity3(G__6696, G__6697, G__6698)
					}(), func() interface{} {
						var G__6699 = x
						var G__6700 = y
						var G__6701 = z
						_, _, _ = G__6699, G__6700, G__6701
						return h.(CljsCoreIFn).X_invoke_Arity3(G__6699, G__6700, G__6701)
					}()}, nil})
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Seq.Arity1IQ(x_y_z_args__[3])
					_, _, _, _ = x, y, z, args
					return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Apply.X_invoke_Arity5(f, x, y, z, args), Apply.X_invoke_Arity5(g, x, y, z, args), Apply.X_invoke_Arity5(h, x, y, z, args)}, nil})
				})
			}(&AFn{})
		}, func(f_g_h_fs__ ...interface{}) interface{} {
			var f = f_g_h_fs__[0]
			var g = f_g_h_fs__[1]
			var h = f_g_h_fs__[2]
			var fs = Seq.Arity1IQ(f_g_h_fs__[3])
			_, _, _, _ = f, g, h, fs
			{
				var fs___1 = List_STAR_.X_invoke_Arity4(f, g, h, fs).(*CljsCoreCons)
				_ = fs___1
				return func(G__6711 *AFn, fs___1 *CljsCoreCons) *AFn {
					return Fn(G__6711, 3, func() interface{} {
						return Reduce.X_invoke_Arity3(func(G__6712 *AFn, fs___1 *CljsCoreCons) *AFn {
							return Fn(G__6712, 2, func(p1__6614_SHARP_ interface{}, p2__6615_SHARP_ interface{}) interface{} {
								return Conj.X_invoke_Arity2(p1__6614_SHARP_, func() interface{} {
									return p2__6615_SHARP_.(CljsCoreIFn).X_invoke_Arity0()
								}())
							})
						}(&AFn{}, fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(G__6713 *AFn, fs___1 *CljsCoreCons) *AFn {
							return Fn(G__6713, 2, func(p1__6616_SHARP_ interface{}, p2__6617_SHARP_ interface{}) interface{} {
								return Conj.X_invoke_Arity2(p1__6616_SHARP_, func() interface{} {
									var G__6702 = x
									_ = G__6702
									return p2__6617_SHARP_.(CljsCoreIFn).X_invoke_Arity1(G__6702)
								}())
							})
						}(&AFn{}, fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x interface{}, y interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(G__6714 *AFn, fs___1 *CljsCoreCons) *AFn {
							return Fn(G__6714, 2, func(p1__6618_SHARP_ interface{}, p2__6619_SHARP_ interface{}) interface{} {
								return Conj.X_invoke_Arity2(p1__6618_SHARP_, func() interface{} {
									var G__6703 = x
									var G__6704 = y
									_, _ = G__6703, G__6704
									return p2__6619_SHARP_.(CljsCoreIFn).X_invoke_Arity2(G__6703, G__6704)
								}())
							})
						}(&AFn{}, fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(G__6715 *AFn, fs___1 *CljsCoreCons) *AFn {
							return Fn(G__6715, 2, func(p1__6620_SHARP_ interface{}, p2__6621_SHARP_ interface{}) interface{} {
								return Conj.X_invoke_Arity2(p1__6620_SHARP_, func() interface{} {
									var G__6705 = x
									var G__6706 = y
									var G__6707 = z
									_, _, _ = G__6705, G__6706, G__6707
									return p2__6621_SHARP_.(CljsCoreIFn).X_invoke_Arity3(G__6705, G__6706, G__6707)
								}())
							})
						}(&AFn{}, fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x_y_z_args__ ...interface{}) interface{} {
						var x = x_y_z_args__[0]
						var y = x_y_z_args__[1]
						var z = x_y_z_args__[2]
						var args = Seq.Arity1IQ(x_y_z_args__[3])
						_, _, _, _ = x, y, z, args
						return Reduce.X_invoke_Arity3(func(G__6716 *AFn, fs___1 *CljsCoreCons) *AFn {
							return Fn(G__6716, 2, func(p1__6622_SHARP_ interface{}, p2__6623_SHARP_ interface{}) interface{} {
								return Conj.X_invoke_Arity2(p1__6622_SHARP_, Apply.X_invoke_Arity5(p2__6623_SHARP_, x, y, z, args))
							})
						}(&AFn{}, fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					})
				}(&AFn{}, fs___1)
			}
		})
	}(&AFn{})

	Dorun = func(dorun *AFn) *AFn {
		return Fn(dorun, 2, func(coll interface{}) interface{} {
			for {
				if Truth_(Seq.Arity1IQ(coll)) {
					coll = Next.Arity1IQ(coll)
					continue
				} else {
					return nil
				}
			}
		}, func(n interface{}, coll interface{}) interface{} {
			for {
				if Truth_(func() interface{} {
					var and__159__auto__ = Seq.Arity1IQ(coll)
					_ = and__159__auto__
					if Truth_(and__159__auto__) {
						return (n.(float64) > float64(0))
					} else {
						return and__159__auto__
					}
				}()) {
					n, coll = (n.(float64) - float64(1)), Next.Arity1IQ(coll)
					continue
				} else {
					return nil
				}
			}
		})
	}(&AFn{})

	Doall = func(doall *AFn) *AFn {
		return Fn(doall, 2, func(coll interface{}) interface{} {
			Dorun.X_invoke_Arity1(coll)
			return coll
		}, func(n interface{}, coll interface{}) interface{} {
			Dorun.X_invoke_Arity2(n, coll)
			return coll
		})
	}(&AFn{})

	Regexp_QMARK_ = func(regexp_QMARK_ *AFn) *AFn {
		return Fn(regexp_QMARK_, 1, func(o interface{}) interface{} {
			return Value_(o).Type().AssignableTo(reflect.TypeOf((**js.RegExp)(nil)).Elem())
		})
	}(&AFn{})

	Re_matches = func(re_matches *AFn) *AFn {
		return Fn(re_matches, 2, func(re interface{}, s interface{}) interface{} {
			if Value_(s).Kind() == reflect.String {
				{
					var matches = Native_invoke_instance_method.X_invoke_Arity3(re, "Exec", []interface{}{s})
					_ = matches
					if X_EQ_.Arity2IIB(First.X_invoke_Arity1(matches), s) {
						if Count.X_invoke_Arity1(matches).(float64) == float64(1) {
							return First.X_invoke_Arity1(matches)
						} else {
							return Vec.X_invoke_Arity1(matches)
						}
					} else {
						return nil
					}
				}
			} else {
				panic((&js.TypeError{"re-matches must match against a string."}))
			}
		})
	}(&AFn{})

	Re_find = func(re_find *AFn) *AFn {
		return Fn(re_find, 2, func(re interface{}, s interface{}) interface{} {
			if Value_(s).Kind() == reflect.String {
				{
					var matches = Native_invoke_instance_method.X_invoke_Arity3(re, "Exec", []interface{}{s})
					_ = matches
					if Nil_(matches) {
						return nil
					} else {
						if Count.X_invoke_Arity1(matches).(float64) == float64(1) {
							return First.X_invoke_Arity1(matches)
						} else {
							return Vec.X_invoke_Arity1(matches)
						}
					}
				}
			} else {
				panic((&js.TypeError{"re-find must match against a string."}))
			}
		})
	}(&AFn{})

	Re_seq = func(re_seq *AFn) *AFn {
		return Fn(re_seq, 2, func(re interface{}, s interface{}) interface{} {
			{
				var match_data = Re_find.X_invoke_Arity2(re, s)
				var match_idx = Native_invoke_instance_method.X_invoke_Arity3(s, "Search", []interface{}{re})
				var match_str = func() interface{} {
					if Coll_QMARK_.Arity1IB(match_data) {
						return First.X_invoke_Arity1(match_data)
					} else {
						return match_data
					}
				}()
				var post_match = Subs.X_invoke_Arity2(s, (match_idx.(float64) + Count.X_invoke_Arity1(match_str).(float64)))
				_, _, _, _ = match_data, match_idx, match_str, post_match
				if Truth_(match_data) {
					return (&CljsCoreLazySeq{nil, func(G__6719 *AFn, match_data interface{}, match_idx interface{}, match_str interface{}, post_match interface{}) *AFn {
						return Fn(G__6719, 0, func() interface{} {
							return Cons.X_invoke_Arity2(match_data, func() interface{} {
								if Truth_(Seq.Arity1IQ(post_match)) {
									return re_seq.X_invoke_Arity2(re, post_match)
								} else {
									return nil
								}
							}()).(*CljsCoreCons)
						})
					}(&AFn{}, match_data, match_idx, match_str, post_match), nil, nil})
				} else {
					return nil
				}
			}
		})
	}(&AFn{})

	Re_pattern = func(re_pattern *AFn) *AFn {
		return Fn(re_pattern, 1, func(s interface{}) interface{} {
			{
				var vec__6721 = Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `^(?:\(\?([idmsux]*)\))?(.*)`, Flags: ``}), s)
				var ___ = Nth.X_invoke_Arity3(vec__6721, float64(0), nil)
				var flags = Nth.X_invoke_Arity3(vec__6721, float64(1), nil)
				var pattern = Nth.X_invoke_Arity3(vec__6721, float64(2), nil)
				_, _, _, _ = vec__6721, ___, flags, pattern
				return (&js.RegExp{pattern, flags})
			}
		})
	}(&AFn{})

	Write_all = func(write_all *AFn) *AFn {
		return Fn(write_all, 1, func(writer_ss__ ...interface{}) interface{} {
			var writer = writer_ss__[0]
			var ss = Seq.Arity1IQ(writer_ss__[1])
			_, _ = writer, ss
			{
				var seq__6740 interface{} = Seq.Arity1IQ(ss)
				var chunk__6741 interface{} = nil
				var count__6742 = float64(0)
				var i__6743 = float64(0)
				_, _, _, _ = seq__6740, chunk__6741, count__6742, i__6743
				for {
					if i__6743 < count__6742 {
						{
							var s = chunk__6741.(CljsCoreIIndexed).X_nth_Arity2(i__6743)
							_ = s
							writer.(CljsCoreIWriter).X_write_Arity2(s)
							seq__6740, chunk__6741, count__6742, i__6743 = seq__6740, chunk__6741, count__6742, (i__6743 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto__ = Seq.Arity1IQ(seq__6740)
							_ = temp__4222__auto__
							if Truth_(temp__4222__auto__) {
								{
									var seq__6740___1 = temp__4222__auto__
									_ = seq__6740___1
									if Chunked_seq_QMARK_.Arity1IB(seq__6740___1) {
										{
											var c__954__auto__ = Chunk_first.X_invoke_Arity1(seq__6740___1)
											_ = c__954__auto__
											seq__6740, chunk__6741, count__6742, i__6743 = Chunk_rest.X_invoke_Arity1(seq__6740___1), c__954__auto__, Count.X_invoke_Arity1(c__954__auto__).(float64), float64(0)
											continue
										}
									} else {
										{
											var s = First.X_invoke_Arity1(seq__6740___1)
											_ = s
											writer.(CljsCoreIWriter).X_write_Arity2(s)
											seq__6740, chunk__6741, count__6742, i__6743 = Next.Arity1IQ(seq__6740___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
								return nil
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	String_print = func(string_print *AFn) *AFn {
		return Fn(string_print, 1, func(x interface{}) interface{} {
			X_STAR_print_fn_STAR_.X_invoke_Arity1(x)
			return nil
		})
	}(&AFn{})

	Flush = func(flush *AFn) *AFn {
		return Fn(flush, 0, func() interface{} {
			return nil
		})
	}(&AFn{})

	Pr_seq_writer = func(pr_seq_writer *AFn) *AFn {
		return Fn(pr_seq_writer, 3, func(objs interface{}, writer interface{}, opts interface{}) interface{} {
			Pr_writer.X_invoke_Arity3(First.X_invoke_Arity1(objs), writer, opts)
			{
				var seq__6759 interface{} = Seq.Arity1IQ(Next.Arity1IQ(objs))
				var chunk__6760 interface{} = nil
				var count__6761 = float64(0)
				var i__6762 = float64(0)
				_, _, _, _ = seq__6759, chunk__6760, count__6761, i__6762
				for {
					if i__6762 < count__6761 {
						{
							var obj = chunk__6760.(CljsCoreIIndexed).X_nth_Arity2(i__6762)
							_ = obj
							writer.(CljsCoreIWriter).X_write_Arity2(" ")
							Pr_writer.X_invoke_Arity3(obj, writer, opts)
							seq__6759, chunk__6760, count__6761, i__6762 = seq__6759, chunk__6760, count__6761, (i__6762 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto__ = Seq.Arity1IQ(seq__6759)
							_ = temp__4222__auto__
							if Truth_(temp__4222__auto__) {
								{
									var seq__6759___1 = temp__4222__auto__
									_ = seq__6759___1
									if Chunked_seq_QMARK_.Arity1IB(seq__6759___1) {
										{
											var c__954__auto__ = Chunk_first.X_invoke_Arity1(seq__6759___1)
											_ = c__954__auto__
											seq__6759, chunk__6760, count__6761, i__6762 = Chunk_rest.X_invoke_Arity1(seq__6759___1), c__954__auto__, Count.X_invoke_Arity1(c__954__auto__).(float64), float64(0)
											continue
										}
									} else {
										{
											var obj = First.X_invoke_Arity1(seq__6759___1)
											_ = obj
											writer.(CljsCoreIWriter).X_write_Arity2(" ")
											Pr_writer.X_invoke_Arity3(obj, writer, opts)
											seq__6759, chunk__6760, count__6761, i__6762 = Next.Arity1IQ(seq__6759___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
								return nil
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Pr_sb_with_opts = func(pr_sb_with_opts *AFn) *AFn {
		return Fn(pr_sb_with_opts, 2, func(objs interface{}, opts interface{}) interface{} {
			{
				var sb = (&goog_string.StringBuffer{})
				var writer = (&CljsCoreStringBufferWriter{sb})
				_, _ = sb, writer
				Pr_seq_writer.X_invoke_Arity3(objs, writer, opts)
				writer.X_flush_Arity1()
				return sb
			}
		})
	}(&AFn{})

	Pr_str_with_opts = func(pr_str_with_opts *AFn) *AFn {
		return Fn(pr_str_with_opts, 2, func(objs interface{}, opts interface{}) interface{} {
			if Empty_QMARK_.Arity1IB(objs) {
				return ""
			} else {
				return (`` + Str.X_invoke_Arity1(Pr_sb_with_opts.X_invoke_Arity2(objs, opts).(*goog_string.StringBuffer)).(string))
			}
		})
	}(&AFn{})

	Prn_str_with_opts = func(prn_str_with_opts *AFn) *AFn {
		return Fn(prn_str_with_opts, 2, func(objs interface{}, opts interface{}) interface{} {
			if Empty_QMARK_.Arity1IB(objs) {
				return "\n"
			} else {
				{
					var sb = Pr_sb_with_opts.X_invoke_Arity2(objs, opts).(*goog_string.StringBuffer)
					_ = sb
					sb.Append("\n")
					return (`` + Str.X_invoke_Arity1(sb).(string))
				}
			}
		})
	}(&AFn{})

	Pr_with_opts = func(pr_with_opts *AFn) *AFn {
		return Fn(pr_with_opts, 2, func(objs interface{}, opts interface{}) interface{} {
			return String_print.X_invoke_Arity1(Pr_str_with_opts.X_invoke_Arity2(objs, opts).(string))
		})
	}(&AFn{})

	Newline = func(newline *AFn) *AFn {
		return Fn(newline, 1, func(opts interface{}) interface{} {
			String_print.X_invoke_Arity1("\n")
			if Truth_(Get.X_invoke_Arity2(opts, (&CljsCoreKeyword{Ns: nil, Name: "flush-on-newline", Fqn: "flush-on-newline", X_hash: float64(-151457939)}))) {
				return Flush.X_invoke_Arity0()
			} else {
				return nil
			}
		})
	}(&AFn{})

	Pr_str = func(pr_str *AFn) *AFn {
		return Fn(pr_str, 0, func(objs__ ...interface{}) interface{} {
			var objs = Seq.Arity1IQ(objs__[0])
			_ = objs
			return Pr_str_with_opts.X_invoke_Arity2(objs, Pr_opts.X_invoke_Arity0().(CljsCoreIMap)).(string)
		})
	}(&AFn{})

	Prn_str = func(prn_str *AFn) *AFn {
		return Fn(prn_str, 0, func(objs__ ...interface{}) interface{} {
			var objs = Seq.Arity1IQ(objs__[0])
			_ = objs
			return Prn_str_with_opts.X_invoke_Arity2(objs, Pr_opts.X_invoke_Arity0().(CljsCoreIMap)).(string)
		})
	}(&AFn{})

	Pr = func(pr *AFn) *AFn {
		return Fn(pr, 0, func(objs__ ...interface{}) interface{} {
			var objs = Seq.Arity1IQ(objs__[0])
			_ = objs
			return Pr_with_opts.X_invoke_Arity2(objs, Pr_opts.X_invoke_Arity0().(CljsCoreIMap))
		})
	}(&AFn{})

	Print = func(cljs_core_print *AFn) *AFn {
		return Fn(cljs_core_print, 0, func(objs__ ...interface{}) interface{} {
			var objs = Seq.Arity1IQ(objs__[0])
			_ = objs
			return Pr_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false))
		})
	}(&AFn{})

	Print_str = func(print_str *AFn) *AFn {
		return Fn(print_str, 0, func(objs__ ...interface{}) interface{} {
			var objs = Seq.Arity1IQ(objs__[0])
			_ = objs
			return Pr_str_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false)).(string)
		})
	}(&AFn{})

	Println = func(println *AFn) *AFn {
		return Fn(println, 0, func(objs__ ...interface{}) interface{} {
			var objs = Seq.Arity1IQ(objs__[0])
			_ = objs
			Pr_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false))
			if Truth_(X_STAR_print_newline_STAR_) {
				return Newline.X_invoke_Arity1(Pr_opts.X_invoke_Arity0().(CljsCoreIMap))
			} else {
				return nil
			}
		})
	}(&AFn{})

	Println_str = func(println_str *AFn) *AFn {
		return Fn(println_str, 0, func(objs__ ...interface{}) interface{} {
			var objs = Seq.Arity1IQ(objs__[0])
			_ = objs
			return Prn_str_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false)).(string)
		})
	}(&AFn{})

	Prn = func(prn *AFn) *AFn {
		return Fn(prn, 0, func(objs__ ...interface{}) interface{} {
			var objs = Seq.Arity1IQ(objs__[0])
			_ = objs
			Pr_with_opts.X_invoke_Arity2(objs, Pr_opts.X_invoke_Arity0().(CljsCoreIMap))
			if Truth_(X_STAR_print_newline_STAR_) {
				return Newline.X_invoke_Arity1(Pr_opts.X_invoke_Arity0().(CljsCoreIMap))
			} else {
				return nil
			}
		})
	}(&AFn{})

	Print_map = func(print_map *AFn) *AFn {
		return Fn(print_map, 4, func(m interface{}, print_one interface{}, writer interface{}, opts interface{}) interface{} {
			return Pr_sequential_writer.X_invoke_Arity7(writer, func(G__6775 *AFn) *AFn {
				return Fn(G__6775, 3, func(e interface{}, w interface{}, opts___1 interface{}) interface{} {
					{
						var G__6769_6776 = Key.X_invoke_Arity1(e)
						var G__6770_6777 = w
						var G__6771_6778 = opts___1
						_, _, _ = G__6769_6776, G__6770_6777, G__6771_6778
						print_one.(CljsCoreIFn).X_invoke_Arity3(G__6769_6776, G__6770_6777, G__6771_6778)
					}
					w.(CljsCoreIWriter).X_write_Arity2(" ")
					{
						var G__6772 = Val.X_invoke_Arity1(e)
						var G__6773 = w
						var G__6774 = opts___1
						_, _, _ = G__6772, G__6773, G__6774
						return print_one.(CljsCoreIFn).X_invoke_Arity3(G__6772, G__6773, G__6774)
					}
				})
			}(&AFn{}), "{", ", ", "}", opts, Seq.Arity1IQ(m))
		})
	}(&AFn{})

	Alter_meta_BANG_ = func(alter_meta_BANG_ *AFn) *AFn {
		return Fn(alter_meta_BANG_, 2, func(iref_f_args__ ...interface{}) interface{} {
			var iref = iref_f_args__[0]
			var f = iref_f_args__[1]
			var args = Seq.Arity1IQ(iref_f_args__[2])
			_, _, _ = iref, f, args
			return func() interface{} {
				var return__6779 = Apply.X_invoke_Arity3(f, Native_get_instance_field.X_invoke_Arity2(iref, "Meta"), args)
				Native_set_instance_field.X_invoke_Arity3(iref, "Meta", return__6779)
				return return__6779
			}()
		})
	}(&AFn{})

	Reset_meta_BANG_ = func(reset_meta_BANG_ *AFn) *AFn {
		return Fn(reset_meta_BANG_, 2, func(iref interface{}, m interface{}) interface{} {
			return func() interface{} {
				var return__6780 = m
				Native_set_instance_field.X_invoke_Arity3(iref, "Meta", return__6780)
				return return__6780
			}()
		})
	}(&AFn{})

	Add_watch = func(add_watch *AFn) *AFn {
		return Fn(add_watch, 3, func(iref interface{}, key interface{}, f interface{}) interface{} {
			return iref.(CljsCoreIWatchable).X_add_watch_Arity3(key, f)
		})
	}(&AFn{})

	Remove_watch = func(remove_watch *AFn) *AFn {
		return Fn(remove_watch, 2, func(iref interface{}, key interface{}) interface{} {
			return iref.(CljsCoreIWatchable).X_remove_watch_Arity2(key)
		})
	}(&AFn{})

	Gensym_counter = nil

	Gensym = func(gensym *AFn) *AFn {
		return Fn(gensym, 1, func() interface{} {
			return gensym.X_invoke_Arity1("G__")
		}, func(prefix_string interface{}) interface{} {
			if Nil_(Gensym_counter) {
				Gensym_counter = Atom.X_invoke_Arity1(float64(0)).(*CljsCoreAtom)

			} else {
			}
			return Symbol.X_invoke_Arity1((`` + Str.X_invoke_Arity1(prefix_string).(string) + Str.X_invoke_Arity1(Swap_BANG_.X_invoke_Arity2(Gensym_counter, Inc)).(string)))
		})
	}(&AFn{})

	Fixture1 = float64(1)

	Fixture2 = float64(2)

	X__GT_Delay = func(__GT_Delay *AFn) *AFn {
		return Fn(__GT_Delay, 2, func(f interface{}, value interface{}) interface{} {
			return (&CljsCoreDelay{f, value})
		})
	}(&AFn{})

	Delay_QMARK_ = func(delay_QMARK_ *AFn) *AFn {
		return Fn(delay_QMARK_, 1, func(x interface{}) bool {
			return Value_(x).Type().AssignableTo(reflect.TypeOf((**CljsCoreDelay)(nil)).Elem())
		})
	}(&AFn{})

	Force = func(force *AFn) *AFn {
		return Fn(force, 1, func(x interface{}) interface{} {
			if Delay_QMARK_.Arity1IB(x) {
				return Deref.X_invoke_Arity1(x)
			} else {
				return x
			}
		})
	}(&AFn{})

	Realized_QMARK_ = func(realized_QMARK_ *AFn) *AFn {
		return Fn(realized_QMARK_, 1, func(d interface{}) bool {
			return d.(CljsCoreIPending).X_realized_QMARK__Arity1()
		})
	}(&AFn{})

	Preserving_reduced = func(preserving_reduced *AFn) *AFn {
		return Fn(preserving_reduced, 1, func(f1 interface{}) interface{} {
			return func(G__6788 *AFn) *AFn {
				return Fn(G__6788, 2, func(p1__6782_SHARP_ interface{}, p2__6783_SHARP_ interface{}) interface{} {
					{
						var ret = func() interface{} {
							var G__6786 = p1__6782_SHARP_
							var G__6787 = p2__6783_SHARP_
							_, _ = G__6786, G__6787
							return f1.(CljsCoreIFn).X_invoke_Arity2(G__6786, G__6787)
						}()
						_ = ret
						if Reduced_QMARK_.Arity1IB(ret) {
							return Reduced.X_invoke_Arity1(ret).(*CljsCoreReduced)
						} else {
							return ret
						}
					}
				})
			}(&AFn{})
		})
	}(&AFn{})

	Cat = func(cat *AFn) *AFn {
		return Fn(cat, 1, func(f1 interface{}) interface{} {
			{
				var rf1 = Preserving_reduced.X_invoke_Arity1(f1).(CljsCoreIFn)
				_ = rf1
				return func(G__6791 *AFn, rf1 CljsCoreIFn) *AFn {
					return Fn(G__6791, 2, func() interface{} {
						{
							return f1.(CljsCoreIFn).X_invoke_Arity0()
						}
					}, func(result interface{}) interface{} {
						{
							var G__6790 = result
							_ = G__6790
							return f1.(CljsCoreIFn).X_invoke_Arity1(G__6790)
						}
					}, func(result interface{}, input interface{}) interface{} {
						return Reduce.X_invoke_Arity3(rf1, result, input)
					})
				}(&AFn{}, rf1)
			}
		})
	}(&AFn{})

	Dedupe = func(dedupe *AFn) *AFn {
		return Fn(dedupe, 1, func() interface{} {
			return func(G__6798 *AFn) *AFn {
				return Fn(G__6798, 1, func(f1 interface{}) interface{} {
					{
						var pa = Atom.X_invoke_Arity1((&CljsCoreKeyword{Ns: "cljs.core", Name: "none", Fqn: "cljs.core/none", X_hash: float64(926646439)})).(*CljsCoreAtom)
						_ = pa
						return func(G__6799 *AFn, pa *CljsCoreAtom) *AFn {
							return Fn(G__6799, 2, func() interface{} {
								{
									return f1.(CljsCoreIFn).X_invoke_Arity0()
								}
							}, func(result interface{}) interface{} {
								{
									var G__6795 = result
									_ = G__6795
									return f1.(CljsCoreIFn).X_invoke_Arity1(G__6795)
								}
							}, func(result interface{}, input interface{}) interface{} {
								{
									var prior = Deref.X_invoke_Arity1(pa)
									_ = prior
									Reset_BANG_.X_invoke_Arity2(pa, input)
									if X_EQ_.Arity2IIB(prior, input) {
										return result
									} else {
										{
											var G__6796 = result
											var G__6797 = input
											_, _ = G__6796, G__6797
											return f1.(CljsCoreIFn).X_invoke_Arity2(G__6796, G__6797)
										}
									}
								}
							})
						}(&AFn{}, pa)
					}
				})
			}(&AFn{})
		}, func(coll interface{}) interface{} {
			return Sequence.X_invoke_Arity2(dedupe.X_invoke_Arity0().(CljsCoreIFn), coll)
		})
	}(&AFn{})

	Random_sample = func(random_sample *AFn) *AFn {
		return Fn(random_sample, 2, func(prob interface{}) interface{} {
			return Filter.X_invoke_Arity1(func(G__6800 *AFn) *AFn {
				return Fn(G__6800, 1, func(___ interface{}) interface{} {
					return (Rand.Arity0F() < prob.(float64))
				})
			}(&AFn{})).(CljsCoreIFn)
		}, func(prob interface{}, coll interface{}) interface{} {
			return Filter.X_invoke_Arity2(func(G__6801 *AFn) *AFn {
				return Fn(G__6801, 1, func(___ interface{}) interface{} {
					return (Rand.Arity0F() < prob.(float64))
				})
			}(&AFn{}), coll).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	X__GT_Iteration = func(__GT_Iteration *AFn) *AFn {
		return Fn(__GT_Iteration, 2, func(xform interface{}, coll interface{}) interface{} {
			return (&CljsCoreIteration{xform, coll})
		})
	}(&AFn{})

	Iteration = func(iteration *AFn) *AFn {
		return Fn(iteration, 2, func(xform interface{}, coll interface{}) interface{} {
			return (&CljsCoreIteration{xform, coll})
		})
	}(&AFn{})

	Run_BANG_ = func(run_BANG_ *AFn) *AFn {
		return Fn(run_BANG_, 2, func(proc interface{}, coll interface{}) interface{} {
			return Reduce.X_invoke_Arity3(func(G__6806 *AFn) *AFn {
				return Fn(G__6806, 2, func(p1__6803_SHARP_ interface{}, p2__6802_SHARP_ interface{}) interface{} {
					{
						var G__6805 = p2__6802_SHARP_
						_ = G__6805
						return proc.(CljsCoreIFn).X_invoke_Arity1(G__6805)
					}
				})
			}(&AFn{}), nil, coll)
		})
	}(&AFn{})

	X_clj__GT_js = func(_clj__GT_js *AFn) *AFn {
		return Fn(_clj__GT_js, 1, func(x interface{}) interface{} {
			return x.(CljsCoreIEncodeJS).X_clj__GT_js_Arity1()
		})
	}(&AFn{})

	X_key__GT_js = func(_key__GT_js *AFn) *AFn {
		return Fn(_key__GT_js, 1, func(x interface{}) interface{} {
			return x.(CljsCoreIEncodeJS).X_key__GT_js_Arity1()
		})
	}(&AFn{})

	X_js__GT_clj = func(_js__GT_clj *AFn) *AFn {
		return Fn(_js__GT_clj, 2, func(x interface{}, options interface{}) interface{} {
			return x.(CljsCoreIEncodeClojure).X_js__GT_clj_Arity2(options)
		})
	}(&AFn{})

	Memoize = func(memoize *AFn) *AFn {
		return Fn(memoize, 1, func(f interface{}) interface{} {
			{
				var mem = Atom.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY).(*CljsCoreAtom)
				_ = mem
				return func(G__6885 *AFn, mem *CljsCoreAtom) *AFn {
					return Fn(G__6885, 0, func(args__ ...interface{}) interface{} {
						var args = Seq.Arity1IQ(args__[0])
						_ = args
						{
							var v = Get.X_invoke_Arity3(Deref.X_invoke_Arity1(mem), args, Lookup_sentinel)
							_ = v
							if reflect.DeepEqual(v, Lookup_sentinel) {
								{
									var ret = Apply.X_invoke_Arity2(f, args)
									_ = ret
									Swap_BANG_.X_invoke_Arity4(mem, Assoc, args, ret)
									return ret
								}
							} else {
								return v
							}
						}
					})
				}(&AFn{}, mem)
			}
		})
	}(&AFn{})

	Trampoline = func(trampoline *AFn) *AFn {
		return Fn(trampoline, 1, func(f interface{}) interface{} {
			for {
				{
					var ret = func() interface{} {
						return f.(CljsCoreIFn).X_invoke_Arity0()
					}()
					_ = ret
					if Fn_QMARK_.Arity1IB(ret) {
						f = ret
						continue
					} else {
						return ret
					}
				}
			}
		}, func(f_args__ ...interface{}) interface{} {
			var f = f_args__[0]
			var args = Seq.Arity1IQ(f_args__[1])
			_, _ = f, args
			return trampoline.X_invoke_Arity1(func(G__6887 *AFn) *AFn {
				return Fn(G__6887, 0, func() interface{} {
					return Apply.X_invoke_Arity2(f, args)
				})
			}(&AFn{}))
		})
	}(&AFn{})

	Rand_int = func(rand_int *AFn) *AFn {
		return Fn(rand_int, 1, func(n interface{}) interface{} {
			{
				var G__6890 = (func() interface{} {
					return Native_invoke_func.X_invoke_Arity2(Math.Random, []interface{}{})
				}().(float64) * n.(float64))
				_ = G__6890
				return Native_invoke_func.X_invoke_Arity2(Math.Floor, []interface{}{G__6890})
			}
		})
	}(&AFn{})

	Rand_nth = func(rand_nth *AFn) *AFn {
		return Fn(rand_nth, 1, func(coll interface{}) interface{} {
			return Nth.X_invoke_Arity2(coll, Rand_int.X_invoke_Arity1(Count.X_invoke_Arity1(coll).(float64)))
		})
	}(&AFn{})

	Group_by = func(group_by *AFn) *AFn {
		return Fn(group_by, 2, func(f interface{}, coll interface{}) interface{} {
			return Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(func(G__6893 *AFn) *AFn {
				return Fn(G__6893, 2, func(ret interface{}, x interface{}) interface{} {
					{
						var k = func() interface{} {
							var G__6892 = x
							_ = G__6892
							return f.(CljsCoreIFn).X_invoke_Arity1(G__6892)
						}()
						_ = k
						return Assoc_BANG_.X_invoke_Arity3(ret, k, Conj.X_invoke_Arity2(Get.X_invoke_Arity3(ret, k, CljsCorePersistentVector_EMPTY), x))
					}
				})
			}(&AFn{}), Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY), coll))
		})
	}(&AFn{})

	Make_hierarchy = func(make_hierarchy *AFn) *AFn {
		return Fn(make_hierarchy, 0, func() interface{} {
			return (&CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}), CljsCorePersistentArrayMap_EMPTY, (&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}), CljsCorePersistentArrayMap_EMPTY, (&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}), CljsCorePersistentArrayMap_EMPTY}, nil})
		})
	}(&AFn{})

	X_global_hierarchy = nil

	Get_global_hierarchy = func(get_global_hierarchy *AFn) *AFn {
		return Fn(get_global_hierarchy, 0, func() interface{} {
			if Nil_(X_global_hierarchy) {
				X_global_hierarchy = Atom.X_invoke_Arity1(Make_hierarchy.X_invoke_Arity0().(CljsCoreIMap)).(*CljsCoreAtom)

			} else {
			}
			return X_global_hierarchy
		})
	}(&AFn{})

	Swap_global_hierarchy_BANG_ = func(swap_global_hierarchy_BANG_ *AFn) *AFn {
		return Fn(swap_global_hierarchy_BANG_, 1, func(f_args__ ...interface{}) interface{} {
			var f = f_args__[0]
			var args = Seq.Arity1IQ(f_args__[1])
			_, _ = f, args
			return Apply.X_invoke_Arity4(Swap_BANG_, Get_global_hierarchy.X_invoke_Arity0(), f, args)
		})
	}(&AFn{})

	Isa_QMARK_ = func(isa_QMARK_ *AFn) *AFn {
		return Fn(isa_QMARK_, 3, func(child interface{}, parent interface{}) bool {
			return isa_QMARK_.Arity3IIIB(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0()), child, parent)
		}, func(h interface{}, child interface{}, parent interface{}) bool {
			{
				var or__171__auto__ = X_EQ_.Arity2IIB(child, parent)
				_ = or__171__auto__
				if Truth_(or__171__auto__) {
					return or__171__auto__
				} else {
					{
						var or__171__auto_____1 = Contains_QMARK_.Arity2IIB((&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h).(CljsCoreIFn).X_invoke_Arity1(child), parent)
						_ = or__171__auto_____1
						if Truth_(or__171__auto_____1) {
							return or__171__auto_____1
						} else {
							{
								var and__159__auto__ = Vector_QMARK_.Arity1IB(parent)
								_ = and__159__auto__
								if Truth_(and__159__auto__) {
									{
										var and__159__auto_____1 = Vector_QMARK_.Arity1IB(child)
										_ = and__159__auto_____1
										if Truth_(and__159__auto_____1) {
											{
												var and__159__auto_____2 = (Count.X_invoke_Arity1(parent).(float64) == Count.X_invoke_Arity1(child).(float64))
												_ = and__159__auto_____2
												if Truth_(and__159__auto_____2) {
													{
														var ret = true
														var i = float64(0)
														_, _ = ret, i
														for {
															if (!(ret)) || (i == Count.X_invoke_Arity1(parent).(float64)) {
																return ret
															} else {
																ret, i = isa_QMARK_.Arity3IIIB(h, func() interface{} {
																	var G__6979 = i
																	_ = G__6979
																	return child.(CljsCoreIFn).X_invoke_Arity1(G__6979)
																}(), func() interface{} {
																	var G__6980 = i
																	_ = G__6980
																	return parent.(CljsCoreIFn).X_invoke_Arity1(G__6980)
																}()), (i + float64(1))
																continue
															}
														}
													}
												} else {
													return and__159__auto_____2
												}
											}
										} else {
											return and__159__auto_____1
										}
									}
								} else {
									return and__159__auto__
								}
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Parents = func(parents *AFn) *AFn {
		return Fn(parents, 2, func(tag interface{}) interface{} {
			return parents.X_invoke_Arity2(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0()), tag)
		}, func(h interface{}, tag interface{}) interface{} {
			return Not_empty.X_invoke_Arity1(Get.X_invoke_Arity2((&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h), tag))
		})
	}(&AFn{})

	Ancestors = func(ancestors *AFn) *AFn {
		return Fn(ancestors, 2, func(tag interface{}) interface{} {
			return ancestors.X_invoke_Arity2(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0()), tag)
		}, func(h interface{}, tag interface{}) interface{} {
			return Not_empty.X_invoke_Arity1(Get.X_invoke_Arity2((&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h), tag))
		})
	}(&AFn{})

	Descendants = func(descendants *AFn) *AFn {
		return Fn(descendants, 2, func(tag interface{}) interface{} {
			return descendants.X_invoke_Arity2(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0()), tag)
		}, func(h interface{}, tag interface{}) interface{} {
			return Not_empty.X_invoke_Arity1(Get.X_invoke_Arity2((&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}).X_invoke_Arity1(h), tag))
		})
	}(&AFn{})

	Derive = func(derive *AFn) *AFn {
		return Fn(derive, 3, func(tag interface{}, parent interface{}) interface{} {
			if Truth_(Namespace.X_invoke_Arity1(parent)) {
			} else {
				panic((&js.Error{("Assert failed: (namespace parent)")}))
			}
			Swap_global_hierarchy_BANG_.X_invoke_ArityVariadic(derive, Array_seq.X_invoke_Arity1([]interface{}{tag, parent}))
			return nil
		}, func(h interface{}, tag interface{}, parent interface{}) interface{} {
			if Not_EQ_.Arity2IIB(tag, parent) {
			} else {
				panic((&js.Error{("Assert failed: (not= tag parent)")}))
			}
			{
				var tp = (&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h)
				var td = (&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}).X_invoke_Arity1(h)
				var ta = (&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h)
				var tf = func(G__7003 *AFn, tp interface{}, td interface{}, ta interface{}) *AFn {
					return Fn(G__7003, 5, func(m interface{}, source interface{}, sources interface{}, target interface{}, targets interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(G__7004 *AFn, tp interface{}, td interface{}, ta interface{}) *AFn {
							return Fn(G__7004, 2, func(ret interface{}, k interface{}) interface{} {
								return Assoc.X_invoke_Arity3(ret, k, Reduce.X_invoke_Arity3(Conj, Get.X_invoke_Arity3(targets, k, CljsCorePersistentHashSet_EMPTY), Cons.X_invoke_Arity2(target, func() interface{} {
									var G__6995 = target
									_ = G__6995
									return targets.(CljsCoreIFn).X_invoke_Arity1(G__6995)
								}()).(*CljsCoreCons)))
							})
						}(&AFn{}, tp, td, ta), m, Cons.X_invoke_Arity2(source, func() interface{} {
							var G__6996 = source
							_ = G__6996
							return sources.(CljsCoreIFn).X_invoke_Arity1(G__6996)
						}()).(*CljsCoreCons))
					})
				}(&AFn{}, tp, td, ta)
				_, _, _, _ = tp, td, ta, tf
				{
					var or__171__auto__ = func() interface{} {
						if Contains_QMARK_.Arity2IIB(func() interface{} {
							var G__7000 = tag
							_ = G__7000
							return tp.(CljsCoreIFn).X_invoke_Arity1(G__7000)
						}(), parent) {
							return nil
						} else {
							return func() CljsCoreIMap {
								if Contains_QMARK_.Arity2IIB(func() interface{} {
									var G__7001 = tag
									_ = G__7001
									return ta.(CljsCoreIFn).X_invoke_Arity1(G__7001)
								}(), parent) {
									panic((&js.Error{(`` + Str.X_invoke_Arity1(tag).(string) + "already has" + Str.X_invoke_Arity1(parent).(string) + "as ancestor")}))
								} else {
								}
								if Contains_QMARK_.Arity2IIB(func() interface{} {
									var G__7002 = parent
									_ = G__7002
									return ta.(CljsCoreIFn).X_invoke_Arity1(G__7002)
								}(), tag) {
									panic((&js.Error{("Cyclic derivation:" + Str.X_invoke_Arity1(parent).(string) + "has" + Str.X_invoke_Arity1(tag).(string) + "as ancestor")}))
								} else {
								}
								return (&CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}), Assoc.X_invoke_Arity3((&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h), tag, Conj.X_invoke_Arity2(Get.X_invoke_Arity3(tp, tag, CljsCorePersistentHashSet_EMPTY), parent)), (&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}), tf.X_invoke_Arity5((&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h), tag, td, parent, ta), (&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}), tf.X_invoke_Arity5((&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}).X_invoke_Arity1(h), parent, ta, tag, td)}, nil})
							}()
						}
					}()
					_ = or__171__auto__
					if Truth_(or__171__auto__) {
						return or__171__auto__
					} else {
						return h
					}
				}
			}
		})
	}(&AFn{})

	Underive = func(underive *AFn) *AFn {
		return Fn(underive, 3, func(tag interface{}, parent interface{}) interface{} {
			Swap_global_hierarchy_BANG_.X_invoke_ArityVariadic(underive, Array_seq.X_invoke_Arity1([]interface{}{tag, parent}))
			return nil
		}, func(h interface{}, tag interface{}, parent interface{}) interface{} {
			{
				var parentMap = (&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h)
				var childsParents = func() interface{} {
					if Truth_(func() interface{} {
						var G__7011 = tag
						_ = G__7011
						return parentMap.(CljsCoreIFn).X_invoke_Arity1(G__7011)
					}()) {
						return Disj.X_invoke_Arity2(func() interface{} {
							var G__7012 = tag
							_ = G__7012
							return parentMap.(CljsCoreIFn).X_invoke_Arity1(G__7012)
						}(), parent)
					} else {
						return CljsCorePersistentHashSet_EMPTY
					}
				}()
				var newParents = func() interface{} {
					if Truth_(Not_empty.X_invoke_Arity1(childsParents)) {
						return Assoc.X_invoke_Arity3(parentMap, tag, childsParents)
					} else {
						return Dissoc.X_invoke_Arity2(parentMap, tag)
					}
				}()
				var deriv_seq = Flatten.X_invoke_Arity1(Map_.X_invoke_Arity2(func(G__7014 *AFn, parentMap interface{}, childsParents interface{}, newParents interface{}) *AFn {
					return Fn(G__7014, 1, func(p1__7005_SHARP_ interface{}) interface{} {
						return Cons.X_invoke_Arity2(First.X_invoke_Arity1(p1__7005_SHARP_), Interpose.X_invoke_Arity2(First.X_invoke_Arity1(p1__7005_SHARP_), Second.X_invoke_Arity1(p1__7005_SHARP_)).(*CljsCoreLazySeq)).(*CljsCoreCons)
					})
				}(&AFn{}, parentMap, childsParents, newParents), Seq.Arity1IQ(newParents)).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
				_, _, _, _ = parentMap, childsParents, newParents, deriv_seq
				if Contains_QMARK_.Arity2IIB(func() interface{} {
					var G__7013 = tag
					_ = G__7013
					return parentMap.(CljsCoreIFn).X_invoke_Arity1(G__7013)
				}(), parent) {
					return Reduce.X_invoke_Arity3(func(G__7015 *AFn, parentMap interface{}, childsParents interface{}, newParents interface{}, deriv_seq *CljsCoreLazySeq) *AFn {
						return Fn(G__7015, 2, func(p1__7006_SHARP_ interface{}, p2__7007_SHARP_ interface{}) interface{} {
							return Apply.X_invoke_Arity3(Derive, p1__7006_SHARP_, p2__7007_SHARP_)
						})
					}(&AFn{}, parentMap, childsParents, newParents, deriv_seq), Make_hierarchy.X_invoke_Arity0().(CljsCoreIMap), Partition.X_invoke_Arity2(float64(2), deriv_seq).(*CljsCoreLazySeq))
				} else {
					return h
				}
			}
		})
	}(&AFn{})

	Reset_cache = func(reset_cache *AFn) *AFn {
		return Fn(reset_cache, 4, func(method_cache interface{}, method_table interface{}, cached_hierarchy interface{}, hierarchy interface{}) interface{} {
			Swap_BANG_.X_invoke_Arity2(method_cache, func(G__7016 *AFn) *AFn {
				return Fn(G__7016, 1, func(___ interface{}) interface{} {
					return Deref.X_invoke_Arity1(method_table)
				})
			}(&AFn{}))
			return Swap_BANG_.X_invoke_Arity2(cached_hierarchy, func(G__7017 *AFn) *AFn {
				return Fn(G__7017, 1, func(___ interface{}) interface{} {
					return Deref.X_invoke_Arity1(hierarchy)
				})
			}(&AFn{}))
		})
	}(&AFn{})

	Prefers_STAR_ = func(prefers_STAR_ *AFn) *AFn {
		return Fn(prefers_STAR_, 3, func(x interface{}, y interface{}, prefer_table interface{}) interface{} {
			{
				var xprefs = Deref.X_invoke_Arity1(prefer_table).(CljsCoreIFn).X_invoke_Arity1(x)
				_ = xprefs
				{
					var or__171__auto__ = func() interface{} {
						if Truth_(func() interface{} {
							var and__159__auto__ = xprefs
							_ = and__159__auto__
							if Truth_(and__159__auto__) {
								{
									var G__7033 = y
									_ = G__7033
									return xprefs.(CljsCoreIFn).X_invoke_Arity1(G__7033)
								}
							} else {
								return and__159__auto__
							}
						}()) {
							return true
						} else {
							return nil
						}
					}()
					_ = or__171__auto__
					if Truth_(or__171__auto__) {
						return or__171__auto__
					} else {
						{
							var or__171__auto_____1 interface{} = func() interface{} {
								var ps interface{} = Parents.X_invoke_Arity1(y)
								_ = ps
								for {
									if Count.X_invoke_Arity1(ps).(float64) > float64(0) {
										if Truth_(prefers_STAR_.X_invoke_Arity3(x, First.X_invoke_Arity1(ps), prefer_table)) {
										} else {
										}
										ps = Rest.Arity1IQ(ps)
										continue
									} else {
										return nil
									}
								}
							}()
							_ = or__171__auto_____1
							if Truth_(or__171__auto_____1) {
								return or__171__auto_____1
							} else {
								{
									var or__171__auto_____2 interface{} = func() interface{} {
										var ps interface{} = Parents.X_invoke_Arity1(x)
										_ = ps
										for {
											if Count.X_invoke_Arity1(ps).(float64) > float64(0) {
												if Truth_(prefers_STAR_.X_invoke_Arity3(First.X_invoke_Arity1(ps), y, prefer_table)) {
												} else {
												}
												ps = Rest.Arity1IQ(ps)
												continue
											} else {
												return nil
											}
										}
									}()
									_ = or__171__auto_____2
									if Truth_(or__171__auto_____2) {
										return or__171__auto_____2
									} else {
										return false
									}
								}
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Dominates = func(dominates *AFn) *AFn {
		return Fn(dominates, 3, func(x interface{}, y interface{}, prefer_table interface{}) interface{} {
			{
				var or__171__auto__ = Prefers_STAR_.X_invoke_Arity3(x, y, prefer_table)
				_ = or__171__auto__
				if Truth_(or__171__auto__) {
					return or__171__auto__
				} else {
					return Isa_QMARK_.Arity2IIB(x, y)
				}
			}
		})
	}(&AFn{})

	Find_and_cache_best_method = func(find_and_cache_best_method *AFn) *AFn {
		return Fn(find_and_cache_best_method, 7, func(name interface{}, dispatch_val interface{}, hierarchy interface{}, method_table interface{}, prefer_table interface{}, method_cache interface{}, cached_hierarchy interface{}) interface{} {
			{
				var best_entry = Reduce.X_invoke_Arity3(func(G__7045 *AFn) *AFn {
					return Fn(G__7045, 2, func(be interface{}, p__7043 interface{}) interface{} {
						{
							var vec__7044 = p__7043
							var k = Nth.X_invoke_Arity3(vec__7044, float64(0), nil)
							var ___ = Nth.X_invoke_Arity3(vec__7044, float64(1), nil)
							var e = vec__7044
							_, _, _, _ = vec__7044, k, ___, e
							if Isa_QMARK_.Arity3IIIB(Deref.X_invoke_Arity1(hierarchy), dispatch_val, k) {
								{
									var be2 = func() interface{} {
										if Truth_(func() interface{} {
											var or__171__auto__ = Nil_(be)
											_ = or__171__auto__
											if Truth_(or__171__auto__) {
												return or__171__auto__
											} else {
												return Dominates.X_invoke_Arity3(k, First.X_invoke_Arity1(be), prefer_table)
											}
										}()) {
											return e
										} else {
											return be
										}
									}()
									_ = be2
									if Truth_(Dominates.X_invoke_Arity3(First.X_invoke_Arity1(be2), k, prefer_table)) {
									} else {
										panic((&js.Error{("Multiple methods in multimethod '" + Str.X_invoke_Arity1(name).(string) + "' match dispatch value: " + Str.X_invoke_Arity1(dispatch_val).(string) + " -> " + Str.X_invoke_Arity1(k).(string) + " and " + Str.X_invoke_Arity1(First.X_invoke_Arity1(be2)).(string) + ", and neither is preferred")}))
									}
									return be2
								}
							} else {
								return be
							}
						}
					})
				}(&AFn{}), nil, Deref.X_invoke_Arity1(method_table))
				_ = best_entry
				if Truth_(best_entry) {
					if X_EQ_.Arity2IIB(Deref.X_invoke_Arity1(cached_hierarchy), Deref.X_invoke_Arity1(hierarchy)) {
						Swap_BANG_.X_invoke_Arity4(method_cache, Assoc, dispatch_val, Second.X_invoke_Arity1(best_entry))
						return Second.X_invoke_Arity1(best_entry)
					} else {
						Reset_cache.X_invoke_Arity4(method_cache, method_table, cached_hierarchy, hierarchy)
						return find_and_cache_best_method.X_invoke_Arity7(name, dispatch_val, hierarchy, method_table, prefer_table, method_cache, cached_hierarchy)
					}
				} else {
					return nil
				}
			}
		})
	}(&AFn{})

	X_reset = func(_reset *AFn) *AFn {
		return Fn(_reset, 1, func(mf interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_reset_Arity1()
		})
	}(&AFn{})

	X_add_method = func(_add_method *AFn) *AFn {
		return Fn(_add_method, 3, func(mf interface{}, dispatch_val interface{}, method interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_add_method_Arity3(dispatch_val, method)
		})
	}(&AFn{})

	X_remove_method = func(_remove_method *AFn) *AFn {
		return Fn(_remove_method, 2, func(mf interface{}, dispatch_val interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_remove_method_Arity2(dispatch_val)
		})
	}(&AFn{})

	X_prefer_method = func(_prefer_method *AFn) *AFn {
		return Fn(_prefer_method, 3, func(mf interface{}, dispatch_val interface{}, dispatch_val_y interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_prefer_method_Arity3(dispatch_val, dispatch_val_y)
		})
	}(&AFn{})

	X_get_method = func(_get_method *AFn) *AFn {
		return Fn(_get_method, 2, func(mf interface{}, dispatch_val interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_get_method_Arity2(dispatch_val)
		})
	}(&AFn{})

	X_methods = func(_methods *AFn) *AFn {
		return Fn(_methods, 1, func(mf interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_methods_Arity1()
		})
	}(&AFn{})

	X_prefers = func(_prefers *AFn) *AFn {
		return Fn(_prefers, 1, func(mf interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_prefers_Arity1()
		})
	}(&AFn{})

	Throw_no_method_error = func(throw_no_method_error *AFn) *AFn {
		return Fn(throw_no_method_error, 2, func(name interface{}, dispatch_val interface{}) interface{} {
			panic((&js.Error{("No method in multimethod '" + Str.X_invoke_Arity1(name).(string) + "' for dispatch value: " + Str.X_invoke_Arity1(dispatch_val).(string))}))
		})
	}(&AFn{})

	X__GT_MultiFn = func(__GT_MultiFn *AFn) *AFn {
		return Fn(__GT_MultiFn, 8, func(name interface{}, dispatch_fn interface{}, default_dispatch_val interface{}, hierarchy interface{}, method_table interface{}, prefer_table interface{}, method_cache interface{}, cached_hierarchy interface{}) interface{} {
			return (&CljsCoreMultiFn{name, dispatch_fn, default_dispatch_val, hierarchy, method_table, prefer_table, method_cache, cached_hierarchy})
		})
	}(&AFn{})

	Remove_all_methods = func(remove_all_methods *AFn) *AFn {
		return Fn(remove_all_methods, 1, func(multifn interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_reset_Arity1()
		})
	}(&AFn{})

	Remove_method = func(remove_method *AFn) *AFn {
		return Fn(remove_method, 2, func(multifn interface{}, dispatch_val interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_remove_method_Arity2(dispatch_val)
		})
	}(&AFn{})

	Prefer_method = func(prefer_method *AFn) *AFn {
		return Fn(prefer_method, 3, func(multifn interface{}, dispatch_val_x interface{}, dispatch_val_y interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_prefer_method_Arity3(dispatch_val_x, dispatch_val_y)
		})
	}(&AFn{})

	Methods = func(methods *AFn) *AFn {
		return Fn(methods, 1, func(multifn interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_methods_Arity1()
		})
	}(&AFn{})

	Get_method = func(get_method *AFn) *AFn {
		return Fn(get_method, 2, func(multifn interface{}, dispatch_val interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_get_method_Arity2(dispatch_val)
		})
	}(&AFn{})

	Prefers = func(prefers *AFn) *AFn {
		return Fn(prefers, 1, func(multifn interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_prefers_Arity1()
		})
	}(&AFn{})

	X__GT_UUID = func(__GT_UUID *AFn) *AFn {
		return Fn(__GT_UUID, 1, func(uuid interface{}) interface{} {
			return (&CljsCoreUUID{uuid})
		})
	}(&AFn{})

	X__GT_ExceptionInfo = func(__GT_ExceptionInfo *AFn) *AFn {
		return Fn(__GT_ExceptionInfo, 3, func(message interface{}, data interface{}, cause interface{}) interface{} {
			return (&CljsCoreExceptionInfo{message, data, cause})
		})
	}(&AFn{})

	Ex_info = func(ex_info *AFn) *AFn {
		return Fn(ex_info, 3, func(msg interface{}, map_ interface{}) interface{} {
			return (&CljsCoreExceptionInfo{msg, map_, nil})
		}, func(msg interface{}, map_ interface{}, cause interface{}) interface{} {
			return (&CljsCoreExceptionInfo{msg, map_, cause})
		})
	}(&AFn{})

	Ex_data = func(ex_data *AFn) *AFn {
		return Fn(ex_data, 1, func(ex interface{}) interface{} {
			if Value_(ex).Type().AssignableTo(reflect.TypeOf((**CljsCoreExceptionInfo)(nil)).Elem()) {
				return Native_get_instance_field.X_invoke_Arity2(ex, "Data")
			} else {
				return nil
			}
		})
	}(&AFn{})

	Ex_message = func(ex_message *AFn) *AFn {
		return Fn(ex_message, 1, func(ex interface{}) interface{} {
			if Value_(ex).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
				return Native_get_instance_field.X_invoke_Arity2(ex, "Message")
			} else {
				return nil
			}
		})
	}(&AFn{})

	Ex_cause = func(ex_cause *AFn) *AFn {
		return Fn(ex_cause, 1, func(ex interface{}) interface{} {
			if Value_(ex).Type().AssignableTo(reflect.TypeOf((**CljsCoreExceptionInfo)(nil)).Elem()) {
				return Native_get_instance_field.X_invoke_Arity2(ex, "Cause")
			} else {
				return nil
			}
		})
	}(&AFn{})

	Comparator = func(comparator *AFn) *AFn {
		return Fn(comparator, 1, func(pred interface{}) interface{} {
			return func(G__7898 *AFn) *AFn {
				return Fn(G__7898, 2, func(x interface{}, y interface{}) interface{} {
					if Truth_(func() interface{} {
						var G__7894 = x
						var G__7895 = y
						_, _ = G__7894, G__7895
						return pred.(CljsCoreIFn).X_invoke_Arity2(G__7894, G__7895)
					}()) {
						return float64(-1)
					} else {
						if Truth_(func() interface{} {
							var G__7896 = y
							var G__7897 = x
							_, _ = G__7896, G__7897
							return pred.(CljsCoreIFn).X_invoke_Arity2(G__7896, G__7897)
						}()) {
							return float64(1)
						} else {
							return float64(0)

						}
					}
				})
			}(&AFn{})
		})
	}(&AFn{})

	Special_symbol_QMARK_ = func(special_symbol_QMARK_ *AFn) *AFn {
		return Fn(special_symbol_QMARK_, 1, func(x interface{}) bool {
			return Contains_QMARK_.Arity2IIB((&CljsCorePersistentHashSet{nil, &CljsCorePersistentArrayMap{nil, float64(19), []interface{}{(&CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "defrecord*", Str: "defrecord*", X_hash: float64(-1936366207), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "try", Str: "try", X_hash: float64(-1273693247), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "loop*", Str: "loop*", X_hash: float64(615029416), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "do", Str: "do", X_hash: float64(1686842252), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "letfn*", Str: "letfn*", X_hash: float64(-110097810), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "if", Str: "if", X_hash: float64(1181717262), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "new", Str: "new", X_hash: float64(-444906321), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "ns", Str: "ns", X_hash: float64(2082130287), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "deftype*", Str: "deftype*", X_hash: float64(962659890), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "let*", Str: "let*", X_hash: float64(1920721458), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "js*", Str: "js*", X_hash: float64(-1134233646), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "fn*", Str: "fn*", X_hash: float64(-752876845), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "recur", Str: "recur", X_hash: float64(1202958259), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "set!", Str: "set!", X_hash: float64(250714521), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: ".", Str: ".", X_hash: float64(1975675962), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "quote", Str: "quote", X_hash: float64(1377916282), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "throw", Str: "throw", X_hash: float64(595905694), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "def", Str: "def", X_hash: float64(597100991), X_meta: nil}), nil}, nil}, nil}), x)
		})
	}(&AFn{})

}

var X_STAR_unchecked_if_STAR_ bool

// Each runtime environment provides a different way to print output.
// Whatever function *print-fn* is bound to will be passed any
// Strings which should be printed.
var X_STAR_print_fn_STAR_ *AFn

var X_STAR_flush_on_newline_STAR_ bool

var X_STAR_print_newline_STAR_ bool

var X_STAR_print_readably_STAR_ bool

var X_STAR_print_meta_STAR_ bool

var X_STAR_print_dup_STAR_ bool

var Pr_opts *AFn

// bound in a repl thread to the most recent value printed
var X_STAR_1 interface{}

// bound in a repl thread to the second most recent value printed
var X_STAR_2 interface{}

// bound in a repl thread to the third most recent value printed
var X_STAR_3 interface{}

var Not_native interface{}

// Tests if 2 arguments are the same object
var Identical_QMARK_ *AFn

// Returns true if x is nil, false otherwise.
var Nil_QMARK_ *AFn

var Array_QMARK_ *AFn

var Number_QMARK_ *AFn

// Returns true if x is logical false, false otherwise.
var Not *AFn

// Returns true if x is not nil, false otherwise.
var Some_QMARK_ *AFn

var String_QMARK_ *AFn

// When compiled for a command-line target, whatever
// function *main-fn* is set to will be called with the command-line
// argv as arguments
var X_STAR_main_cli_fn_STAR_ interface{}

// Returns a javascript array, cloned from the passed in array
var Aclone *AFn

// Returns the value at the index.
// @param {...*} var_args
var Aget *AFn

// Sets the value at the index.
// @param {...*} var_args
var Aset *AFn

// Returns the length of the array. Works on arrays of all types.
var Alength *AFn

var Into_array *AFn

type CljsCoreFn interface {
	CljsCoreFn__()
}

type CljsCoreIFn interface {
	CljsCoreIFn__()
	X_invoke_Arity0() interface{}
	X_invoke_Arity1(a interface{}) interface{}
	X_invoke_Arity2(a interface{}, b interface{}) interface{}
	X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{}
	X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{}
	X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{}
	X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{}
	X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{}
	X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{}
	X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{}
	X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{}
	X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{}
	X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{}
	X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{}
	X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{}
	X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{}
	X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{}
	X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{}
	X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{}
	X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{}
	X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{}
}

var X_invoke *AFn

type CljsCoreICloneable interface {
	CljsCoreICloneable__()
	X_clone_Arity1() interface{}
}

var X_clone *AFn

type CljsCoreICounted interface {
	CljsCoreICounted__()
	X_count_Arity1() float64
}

var X_count *AFn

type CljsCoreIEmptyableCollection interface {
	CljsCoreIEmptyableCollection__()
	X_empty_Arity1() interface{}
}

var X_empty *AFn

type CljsCoreICollection interface {
	CljsCoreICollection__()
	X_conj_Arity2(o interface{}) interface{}
}

var X_conj *AFn

type CljsCoreIIndexed interface {
	CljsCoreIIndexed__()
	X_nth_Arity2(n interface{}) interface{}
	X_nth_Arity3(n interface{}, not_found interface{}) interface{}
}

var X_nth *AFn

type CljsCoreASeq interface {
	CljsCoreASeq__()
}

type CljsCoreISeq interface {
	CljsCoreISeq__()
	X_first_Arity1() interface{}
	X_rest_Arity1() interface{}
}

var X_first *AFn

var X_rest *AFn

type CljsCoreINext interface {
	CljsCoreINext__()
	X_next_Arity1() interface{}
}

var X_next *AFn

type CljsCoreILookup interface {
	CljsCoreILookup__()
	X_lookup_Arity2(k interface{}) interface{}
	X_lookup_Arity3(k interface{}, not_found interface{}) interface{}
}

var X_lookup *AFn

type CljsCoreIAssociative interface {
	CljsCoreIAssociative__()
	X_contains_key_QMARK__Arity2(k interface{}) bool
	X_assoc_Arity3(k interface{}, v interface{}) interface{}
}

var X_contains_key_QMARK_ *AFn

var X_assoc *AFn

type CljsCoreIMap interface {
	CljsCoreIMap__()
	X_dissoc_Arity2(k interface{}) interface{}
}

var X_dissoc *AFn

type CljsCoreIMapEntry interface {
	CljsCoreIMapEntry__()
	X_key_Arity1() interface{}
	X_val_Arity1() interface{}
}

var X_key *AFn

var X_val *AFn

type CljsCoreISet interface {
	CljsCoreISet__()
	X_disjoin_Arity2(v interface{}) interface{}
}

var X_disjoin *AFn

type CljsCoreIStack interface {
	CljsCoreIStack__()
	X_peek_Arity1() interface{}
	X_pop_Arity1() interface{}
}

var X_peek *AFn

var X_pop *AFn

type CljsCoreIVector interface {
	CljsCoreIVector__()
	X_assoc_n_Arity3(n interface{}, val interface{}) interface{}
}

var X_assoc_n *AFn

type CljsCoreIDeref interface {
	CljsCoreIDeref__()
	X_deref_Arity1() interface{}
}

var X_deref *AFn

type CljsCoreIDerefWithTimeout interface {
	CljsCoreIDerefWithTimeout__()
	X_deref_with_timeout_Arity3(msec interface{}, timeout_val interface{}) interface{}
}

var X_deref_with_timeout *AFn

type CljsCoreIMeta interface {
	CljsCoreIMeta__()
	X_meta_Arity1() interface{}
}

var X_meta *AFn

type CljsCoreIWithMeta interface {
	CljsCoreIWithMeta__()
	X_with_meta_Arity2(meta interface{}) interface{}
}

var X_with_meta *AFn

type CljsCoreIReduce interface {
	CljsCoreIReduce__()
	X_reduce_Arity2(f interface{}) interface{}
	X_reduce_Arity3(f interface{}, start interface{}) interface{}
}

var X_reduce *AFn

type CljsCoreIKVReduce interface {
	CljsCoreIKVReduce__()
	X_kv_reduce_Arity3(f interface{}, init interface{}) interface{}
}

var X_kv_reduce *AFn

type CljsCoreIEquiv interface {
	CljsCoreIEquiv__()
	X_equiv_Arity2(other interface{}) bool
}

var X_equiv *AFn

type CljsCoreIHash interface {
	CljsCoreIHash__()
	X_hash_Arity1() interface{}
}

var X_hash *AFn

type CljsCoreISeqable interface {
	CljsCoreISeqable__()
	X_seq_Arity1() interface{}
}

var X_seq *AFn

type CljsCoreISequential interface {
	CljsCoreISequential__()
}

type CljsCoreIList interface {
	CljsCoreIList__()
}

type CljsCoreIRecord interface {
	CljsCoreIRecord__()
}

type CljsCoreIReversible interface {
	CljsCoreIReversible__()
	X_rseq_Arity1() interface{}
}

var X_rseq *AFn

type CljsCoreISorted interface {
	CljsCoreISorted__()
	X_sorted_seq_Arity2(ascending_QMARK_ interface{}) interface{}
	X_sorted_seq_from_Arity3(k interface{}, ascending_QMARK_ interface{}) interface{}
	X_entry_key_Arity2(entry interface{}) interface{}
	X_comparator_Arity1() interface{}
}

var X_sorted_seq *AFn

var X_sorted_seq_from *AFn

var X_entry_key *AFn

var X_comparator *AFn

type CljsCoreIWriter interface {
	CljsCoreIWriter__()
	X_write_Arity2(s interface{}) interface{}
	X_flush_Arity1() interface{}
}

var X_write *AFn

var X_flush *AFn

type CljsCoreIPrintWithWriter interface {
	CljsCoreIPrintWithWriter__()
	X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{}
}

var X_pr_writer *AFn

type CljsCoreIPending interface {
	CljsCoreIPending__()
	X_realized_QMARK__Arity1() bool
}

var X_realized_QMARK_ *AFn

type CljsCoreIWatchable interface {
	CljsCoreIWatchable__()
	X_notify_watches_Arity3(oldval interface{}, newval interface{}) interface{}
	X_add_watch_Arity3(key interface{}, f interface{}) interface{}
	X_remove_watch_Arity2(key interface{}) interface{}
}

var X_notify_watches *AFn

var X_add_watch *AFn

var X_remove_watch *AFn

type CljsCoreIEditableCollection interface {
	CljsCoreIEditableCollection__()
	X_as_transient_Arity1() interface{}
}

var X_as_transient *AFn

type CljsCoreITransientCollection interface {
	CljsCoreITransientCollection__()
	X_conj_BANG__Arity2(val interface{}) interface{}
	X_persistent_BANG__Arity1() interface{}
}

var X_conj_BANG_ *AFn

var X_persistent_BANG_ *AFn

type CljsCoreITransientAssociative interface {
	CljsCoreITransientAssociative__()
	X_assoc_BANG__Arity3(key interface{}, val interface{}) interface{}
}

var X_assoc_BANG_ *AFn

type CljsCoreITransientMap interface {
	CljsCoreITransientMap__()
	X_dissoc_BANG__Arity2(key interface{}) interface{}
}

var X_dissoc_BANG_ *AFn

type CljsCoreITransientVector interface {
	CljsCoreITransientVector__()
	X_assoc_n_BANG__Arity3(n interface{}, val interface{}) interface{}
	X_pop_BANG__Arity1() interface{}
}

var X_assoc_n_BANG_ *AFn

var X_pop_BANG_ *AFn

type CljsCoreITransientSet interface {
	CljsCoreITransientSet__()
	X_disjoin_BANG__Arity2(v interface{}) interface{}
}

var X_disjoin_BANG_ *AFn

type CljsCoreIComparable interface {
	CljsCoreIComparable__()
	X_compare_Arity2(y interface{}) float64
}

var X_compare *AFn

type CljsCoreIChunk interface {
	CljsCoreIChunk__()
	X_drop_first_Arity1() interface{}
}

var X_drop_first *AFn

type CljsCoreIChunkedSeq interface {
	CljsCoreIChunkedSeq__()
	X_chunked_first_Arity1() interface{}
	X_chunked_rest_Arity1() interface{}
}

var X_chunked_first *AFn

var X_chunked_rest *AFn

type CljsCoreIChunkedNext interface {
	CljsCoreIChunkedNext__()
	X_chunked_next_Arity1() interface{}
}

var X_chunked_next *AFn

type CljsCoreINamed interface {
	CljsCoreINamed__()
	X_name_Arity1() interface{}
	X_namespace_Arity1() interface{}
}

var X_name *AFn

var X_namespace *AFn

type CljsCoreIAtom interface {
	CljsCoreIAtom__()
}

type CljsCoreIReset interface {
	CljsCoreIReset__()
	X_reset_BANG__Arity2(new_value interface{}) interface{}
}

var X_reset_BANG_ *AFn

type CljsCoreISwap interface {
	CljsCoreISwap__()
	X_swap_BANG__Arity2(f interface{}) interface{}
	X_swap_BANG__Arity3(f interface{}, a interface{}) interface{}
	X_swap_BANG__Arity4(f interface{}, a interface{}, b interface{}) interface{}
	X_swap_BANG__Arity5(f interface{}, a interface{}, b interface{}, xs interface{}) interface{}
}

var X_swap_BANG_ *AFn

type CljsCoreIIterable interface {
	CljsCoreIIterable__()
	X_iterator_Arity1() interface{}
}

var X_iterator *AFn

type CljsCoreStringBufferWriter struct{ Sb interface{} }

func (_ *CljsCoreStringBufferWriter) CljsCoreIWriter__() {}

func (___ *CljsCoreStringBufferWriter) X_write_Arity2(s interface{}) interface{} {
	return Native_invoke_instance_method.X_invoke_Arity3(___.Sb, "Append", []interface{}{s})
}

func (___ *CljsCoreStringBufferWriter) X_flush_Arity1() interface{} {
	return nil
}

var X__GT_StringBufferWriter *AFn

// Support so that collections can implement toString without
// loading all the printing machinery.
var Pr_str_STAR_ *AFn

var Int_rotate_left *AFn

var Imul *AFn

var M3_seed float64

var M3_C1 float64

var M3_C2 float64

var M3_mix_K1 *AFn

var M3_mix_H1 *AFn

var M3_fmix *AFn

var M3_hash_int *AFn

var M3_hash_unencoded_chars *AFn

var String_hash_cache_count float64

var Hash_string_STAR_ *AFn

var Hash *AFn

var Hash_combine *AFn

var Hash_symbol *AFn

var Compare_symbols *AFn

type CljsCoreSymbol struct {
	Ns     interface{}
	Name   interface{}
	Str    interface{}
	X_hash interface{}
	X_meta interface{}
}

func (_ *CljsCoreSymbol) CljsCoreIPrintWithWriter__() {}

func (o *CljsCoreSymbol) X_pr_writer_Arity3(writer interface{}, ___ interface{}) interface{} {
	return writer.(CljsCoreIWriter).X_write_Arity2(o.Str)
}

func (_ *CljsCoreSymbol) CljsCoreINamed__() {}

func (___ *CljsCoreSymbol) X_name_Arity1() interface{} {
	return ___.Name
}

func (___ *CljsCoreSymbol) X_namespace_Arity1() interface{} {
	return ___.Ns
}

func (_ *CljsCoreSymbol) CljsCoreIHash__() {}

func (sym *CljsCoreSymbol) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = sym.X_hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_symbol.X_invoke_Arity1(sym).(float64)
				_ = h__577__auto_____1
				sym.X_hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreSymbol) CljsCoreIWithMeta__() {}

func (___ *CljsCoreSymbol) X_with_meta_Arity2(new_meta interface{}) interface{} {
	return (&CljsCoreSymbol{___.Ns, ___.Name, ___.Str, ___.X_hash, new_meta})
}

func (_ *CljsCoreSymbol) CljsCoreIMeta__() {}

func (___ *CljsCoreSymbol) X_meta_Arity1() interface{} {
	return ___.X_meta
}

func (_ *CljsCoreSymbol) CljsCoreIFn__() {}

func (this *CljsCoreSymbol) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (sym *CljsCoreSymbol) X_invoke_Arity1(coll interface{}) interface{} {
	return coll.(CljsCoreILookup).X_lookup_Arity3(sym, nil)
}

func (sym *CljsCoreSymbol) X_invoke_Arity2(coll interface{}, not_found interface{}) interface{} {
	return coll.(CljsCoreILookup).X_lookup_Arity3(sym, not_found)
}

func (this *CljsCoreSymbol) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

func (_ *CljsCoreSymbol) CljsCoreIEquiv__() {}

func (___ *CljsCoreSymbol) X_equiv_Arity2(other interface{}) bool {
	if Value_(other).Type().AssignableTo(reflect.TypeOf((**CljsCoreSymbol)(nil)).Elem()) {
		return reflect.DeepEqual(___.Str, Native_get_instance_field.X_invoke_Arity2(other, "Str"))
	} else {
		return false
	}
}

func (_ *CljsCoreSymbol) CljsCoreObject__() {}

func (___ *CljsCoreSymbol) ToString() string {
	return ___.Str.(string)
}

func (this *CljsCoreSymbol) String() string {
	return this.ToString()
}

func (this *CljsCoreSymbol) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

var X__GT_Symbol *AFn

var Symbol *AFn

var Iterable_QMARK_ *AFn

var Clone *AFn

var Cloneable_QMARK_ *AFn

// Returns a seq on the collection. If the collection is
// empty, returns nil.  (seq nil) returns nil. seq also works on
// Strings.
var Seq *AFn

// Returns the first item in the collection. Calls seq on its
// argument. If coll is nil, returns nil.
var First *AFn

// Returns a possibly empty seq of the items after the first. Calls seq on its
// argument.
var Rest *AFn

// Returns a seq of the items after the first. Calls seq on its
// argument.  If there are no more items, returns nil
var Next *AFn

// Mix final collection hash for ordered or unordered collections.
// hash-basis is the combined collection hash, count is the number
// of elements included in the basis. Note this is the hash code
// consistent with =, different from .hashCode.
// See http://clojure.org/data_structures#hash for full algorithms.
var Mix_collection_hash *AFn

// Returns the hash code, consistent with =, for an external ordered
// collection implementing Iterable.
// See http://clojure.org/data_structures#hash for full algorithms.
var Hash_ordered_coll *AFn

// Returns the hash code, consistent with =, for an external unordered
// collection implementing Iterable. For maps, the iterator should
// return map entries whose hash is computed as
// (hash-ordered-coll [k v]).
// See http://clojure.org/data_structures#hash for full algorithms.
var Hash_unordered_coll *AFn

// Returns a number one greater than num.
var Inc *AFn

type CljsCoreReduced struct{ Val interface{} }

func (_ *CljsCoreReduced) CljsCoreIDeref__() {}

func (o *CljsCoreReduced) X_deref_Arity1() interface{} {
	return o.Val
}

var X__GT_Reduced *AFn

// Wraps x in a way such that a reduce will terminate with the value x
var Reduced *AFn

// Returns true if x is the result of a call to reduced
var Reduced_QMARK_ *AFn

var Deref *AFn

// Accepts any collection which satisfies the ICount and IIndexed protocols and
// reduces them without incurring seq initialization
var Ci_reduce *AFn

var Array_reduce *AFn

// Returns true if coll implements count in constant time
var Counted_QMARK_ *AFn

// Returns true if coll implements nth in constant time
var Indexed_QMARK_ *AFn

type CljsCoreIndexedSeq struct {
	Arr interface{}
	I   interface{}
}

func (_ *CljsCoreIndexedSeq) CljsCoreObject__() {}

func (coll *CljsCoreIndexedSeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreIndexedSeq) String() string {
	return this.ToString()
}

func (this *CljsCoreIndexedSeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreIndexedSeq) CljsCoreIIndexed__() {}

func (coll *CljsCoreIndexedSeq) X_nth_Arity2(n interface{}) interface{} {
	{
		var i___1 = (n.(float64) + coll.I.(float64))
		_ = i___1
		if i___1 < Alength_(coll.Arr) {
			return Aget_(coll.Arr, i___1)
		} else {
			return nil
		}
	}
}

func (coll *CljsCoreIndexedSeq) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	{
		var i___1 = (n.(float64) + coll.I.(float64))
		_ = i___1
		if i___1 < Alength_(coll.Arr) {
			return Aget_(coll.Arr, i___1)
		} else {
			return not_found
		}
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreICloneable__() {}

func (___ *CljsCoreIndexedSeq) X_clone_Arity1() interface{} {
	return (&CljsCoreIndexedSeq{___.Arr, ___.I})
}

func (_ *CljsCoreIndexedSeq) CljsCoreINext__() {}

func (___ *CljsCoreIndexedSeq) X_next_Arity1() interface{} {
	if (___.I.(float64) + float64(1)) < Alength_(___.Arr) {
		return (&CljsCoreIndexedSeq{___.Arr, (___.I.(float64) + float64(1))})
	} else {
		return nil
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreICounted__() {}

func (___ *CljsCoreIndexedSeq) X_count_Arity1() float64 {
	return (Alength_(___.Arr) - ___.I.(float64))
}

func (_ *CljsCoreIndexedSeq) CljsCoreIReversible__() {}

func (coll *CljsCoreIndexedSeq) X_rseq_Arity1() interface{} {
	{
		var c = coll.X_count_Arity1()
		_ = c
		if c > float64(0) {
			return (&CljsCoreRSeq{coll, (c - float64(1)), nil})
		} else {
			return nil
		}
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreIHash__() {}

func (coll *CljsCoreIndexedSeq) X_hash_Arity1() interface{} {
	return Hash_ordered_coll.Arity1IF(coll)
}

func (_ *CljsCoreIndexedSeq) CljsCoreIEquiv__() {}

func (coll *CljsCoreIndexedSeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreIndexedSeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreIndexedSeq) X_empty_Arity1() interface{} {
	return CljsCoreList_EMPTY
}

func (_ *CljsCoreIndexedSeq) CljsCoreIReduce__() {}

func (coll *CljsCoreIndexedSeq) X_reduce_Arity2(f interface{}) interface{} {
	return Array_reduce.X_invoke_Arity4(coll.Arr, f, Aget_(coll.Arr, coll.I.(float64)), (coll.I.(float64) + float64(1)))
}

func (coll *CljsCoreIndexedSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Array_reduce.X_invoke_Arity4(coll.Arr, f, start, coll.I)
}

func (_ *CljsCoreIndexedSeq) CljsCoreISeq__() {}

func (___ *CljsCoreIndexedSeq) X_first_Arity1() interface{} {
	return Aget_(___.Arr, ___.I.(float64))
}

func (___ *CljsCoreIndexedSeq) X_rest_Arity1() interface{} {
	if (___.I.(float64) + float64(1)) < Alength_(___.Arr) {
		return (&CljsCoreIndexedSeq{___.Arr, (___.I.(float64) + float64(1))})
	} else {
		return CljsCoreList_EMPTY
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreISeqable__() {}

func (this *CljsCoreIndexedSeq) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCoreIndexedSeq) CljsCoreISequential__() {}

func (_ *CljsCoreIndexedSeq) CljsCoreICollection__() {}

func (coll *CljsCoreIndexedSeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

func (_ *CljsCoreIndexedSeq) CljsCoreASeq__() {}

var X__GT_IndexedSeq *AFn

var Prim_seq *AFn

var Array_seq *AFn

type CljsCoreRSeq struct {
	Ci   interface{}
	I    interface{}
	Meta interface{}
}

func (_ *CljsCoreRSeq) CljsCoreObject__() {}

func (coll *CljsCoreRSeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreRSeq) String() string {
	return this.ToString()
}

func (this *CljsCoreRSeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreRSeq) CljsCoreIMeta__() {}

func (coll *CljsCoreRSeq) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreRSeq) CljsCoreICloneable__() {}

func (___ *CljsCoreRSeq) X_clone_Arity1() interface{} {
	return (&CljsCoreRSeq{___.Ci, ___.I, ___.Meta})
}

func (_ *CljsCoreRSeq) CljsCoreINext__() {}

func (coll *CljsCoreRSeq) X_next_Arity1() interface{} {
	if coll.I.(float64) > float64(0) {
		return (&CljsCoreRSeq{coll.Ci, (coll.I.(float64) - float64(1)), nil})
	} else {
		return nil
	}
}

func (_ *CljsCoreRSeq) CljsCoreICounted__() {}

func (coll *CljsCoreRSeq) X_count_Arity1() float64 {
	return (coll.I.(float64) + float64(1))
}

func (_ *CljsCoreRSeq) CljsCoreIHash__() {}

func (coll *CljsCoreRSeq) X_hash_Arity1() interface{} {
	return Hash_ordered_coll.Arity1IF(coll)
}

func (_ *CljsCoreRSeq) CljsCoreIEquiv__() {}

func (coll *CljsCoreRSeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreRSeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreRSeq) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.Meta)
}

func (_ *CljsCoreRSeq) CljsCoreIReduce__() {}

func (col *CljsCoreRSeq) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, col)
}

func (col *CljsCoreRSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, col)
}

func (_ *CljsCoreRSeq) CljsCoreISeq__() {}

func (coll *CljsCoreRSeq) X_first_Arity1() interface{} {
	return coll.Ci.(CljsCoreIIndexed).X_nth_Arity2(coll.I)
}

func (coll *CljsCoreRSeq) X_rest_Arity1() interface{} {
	if coll.I.(float64) > float64(0) {
		return (&CljsCoreRSeq{coll.Ci, (coll.I.(float64) - float64(1)), nil})
	} else {
		return CljsCoreIEmptyList(CljsCoreList_EMPTY)
	}
}

func (_ *CljsCoreRSeq) CljsCoreISeqable__() {}

func (coll *CljsCoreRSeq) X_seq_Arity1() interface{} {
	return coll
}

func (_ *CljsCoreRSeq) CljsCoreISequential__() {}

func (_ *CljsCoreRSeq) CljsCoreIWithMeta__() {}

func (coll *CljsCoreRSeq) X_with_meta_Arity2(new_meta interface{}) interface{} {
	return (&CljsCoreRSeq{coll.Ci, coll.I, new_meta})
}

func (_ *CljsCoreRSeq) CljsCoreICollection__() {}

func (coll *CljsCoreRSeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

var X__GT_RSeq *AFn

// Same as (first (next x))
var Second *AFn

// Same as (first (first x))
var Ffirst *AFn

// Same as (next (first x))
var Nfirst *AFn

// Same as (first (next x))
var Fnext *AFn

// Same as (next (next x))
var Nnext *AFn

// Return the last item in coll, in linear time
var Last *AFn

// conj[oin]. Returns a new collection with the xs
// 'added'. (conj nil item) returns (item).  The 'addition' may
// happen at different 'places' depending on the concrete type.
// @param {...*} var_args
var Conj *AFn

// Returns an empty collection of the same category as coll, or nil
var Empty *AFn

var Accumulating_seq_count *AFn

// Returns the number of items in the collection. (count nil) returns
// 0.  Also works on strings, arrays, and Maps
var Count *AFn

var Linear_traversal_nth *AFn

// Returns the value at the index. get returns nil if index out of
// bounds, nth throws an exception unless not-found is supplied.  nth
// also works for strings, arrays, regex Matchers and Lists, and,
// in O(n) time, for sequences.
var Nth *AFn

// assoc[iate]. When applied to a map, returns a new map of the
// same (hashed/sorted) type, that contains the mapping of key(s) to
// val(s). When applied to a vector, returns a new vector that
// contains val at index.
// @param {...*} var_args
var Assoc *AFn

// dissoc[iate]. Returns a new map of the same (hashed/sorted) type,
// that does not contain a mapping for key(s).
// @param {...*} var_args
var Dissoc *AFn

var Fn_QMARK_ *AFn

type CljsCoreMetaFn struct {
	Afn  interface{}
	Meta interface{}
}

func (_ *CljsCoreMetaFn) CljsCoreIFn__() {}

func (___ *CljsCoreMetaFn) X_invoke_Arity0() interface{} {
	{
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity0()
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity1(a interface{}) interface{} {
	{
		var G__4103 = a
		_ = G__4103
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity1(G__4103)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity2(a interface{}, b interface{}) interface{} {
	{
		var G__4106 = a
		var G__4107 = b
		_, _ = G__4106, G__4107
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity2(G__4106, G__4107)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var G__4111 = a
		var G__4112 = b
		var G__4113 = c
		_, _, _ = G__4111, G__4112, G__4113
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity3(G__4111, G__4112, G__4113)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var G__4118 = a
		var G__4119 = b
		var G__4120 = c
		var G__4121 = d
		_, _, _, _ = G__4118, G__4119, G__4120, G__4121
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity4(G__4118, G__4119, G__4120, G__4121)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var G__4127 = a
		var G__4128 = b
		var G__4129 = c
		var G__4130 = d
		var G__4131 = e
		_, _, _, _, _ = G__4127, G__4128, G__4129, G__4130, G__4131
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity5(G__4127, G__4128, G__4129, G__4130, G__4131)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var G__4138 = a
		var G__4139 = b
		var G__4140 = c
		var G__4141 = d
		var G__4142 = e
		var G__4143 = f
		_, _, _, _, _, _ = G__4138, G__4139, G__4140, G__4141, G__4142, G__4143
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity6(G__4138, G__4139, G__4140, G__4141, G__4142, G__4143)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var G__4151 = a
		var G__4152 = b
		var G__4153 = c
		var G__4154 = d
		var G__4155 = e
		var G__4156 = f
		var G__4157 = g
		_, _, _, _, _, _, _ = G__4151, G__4152, G__4153, G__4154, G__4155, G__4156, G__4157
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity7(G__4151, G__4152, G__4153, G__4154, G__4155, G__4156, G__4157)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var G__4166 = a
		var G__4167 = b
		var G__4168 = c
		var G__4169 = d
		var G__4170 = e
		var G__4171 = f
		var G__4172 = g
		var G__4173 = h
		_, _, _, _, _, _, _, _ = G__4166, G__4167, G__4168, G__4169, G__4170, G__4171, G__4172, G__4173
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity8(G__4166, G__4167, G__4168, G__4169, G__4170, G__4171, G__4172, G__4173)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var G__4183 = a
		var G__4184 = b
		var G__4185 = c
		var G__4186 = d
		var G__4187 = e
		var G__4188 = f
		var G__4189 = g
		var G__4190 = h
		var G__4191 = i
		_, _, _, _, _, _, _, _, _ = G__4183, G__4184, G__4185, G__4186, G__4187, G__4188, G__4189, G__4190, G__4191
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity9(G__4183, G__4184, G__4185, G__4186, G__4187, G__4188, G__4189, G__4190, G__4191)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var G__4202 = a
		var G__4203 = b
		var G__4204 = c
		var G__4205 = d
		var G__4206 = e
		var G__4207 = f
		var G__4208 = g
		var G__4209 = h
		var G__4210 = i
		var G__4211 = j
		_, _, _, _, _, _, _, _, _, _ = G__4202, G__4203, G__4204, G__4205, G__4206, G__4207, G__4208, G__4209, G__4210, G__4211
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity10(G__4202, G__4203, G__4204, G__4205, G__4206, G__4207, G__4208, G__4209, G__4210, G__4211)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var G__4223 = a
		var G__4224 = b
		var G__4225 = c
		var G__4226 = d
		var G__4227 = e
		var G__4228 = f
		var G__4229 = g
		var G__4230 = h
		var G__4231 = i
		var G__4232 = j
		var G__4233 = k
		_, _, _, _, _, _, _, _, _, _, _ = G__4223, G__4224, G__4225, G__4226, G__4227, G__4228, G__4229, G__4230, G__4231, G__4232, G__4233
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity11(G__4223, G__4224, G__4225, G__4226, G__4227, G__4228, G__4229, G__4230, G__4231, G__4232, G__4233)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var G__4246 = a
		var G__4247 = b
		var G__4248 = c
		var G__4249 = d
		var G__4250 = e
		var G__4251 = f
		var G__4252 = g
		var G__4253 = h
		var G__4254 = i
		var G__4255 = j
		var G__4256 = k
		var G__4257 = l
		_, _, _, _, _, _, _, _, _, _, _, _ = G__4246, G__4247, G__4248, G__4249, G__4250, G__4251, G__4252, G__4253, G__4254, G__4255, G__4256, G__4257
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity12(G__4246, G__4247, G__4248, G__4249, G__4250, G__4251, G__4252, G__4253, G__4254, G__4255, G__4256, G__4257)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var G__4271 = a
		var G__4272 = b
		var G__4273 = c
		var G__4274 = d
		var G__4275 = e
		var G__4276 = f
		var G__4277 = g
		var G__4278 = h
		var G__4279 = i
		var G__4280 = j
		var G__4281 = k
		var G__4282 = l
		var G__4283 = m
		_, _, _, _, _, _, _, _, _, _, _, _, _ = G__4271, G__4272, G__4273, G__4274, G__4275, G__4276, G__4277, G__4278, G__4279, G__4280, G__4281, G__4282, G__4283
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity13(G__4271, G__4272, G__4273, G__4274, G__4275, G__4276, G__4277, G__4278, G__4279, G__4280, G__4281, G__4282, G__4283)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var G__4298 = a
		var G__4299 = b
		var G__4300 = c
		var G__4301 = d
		var G__4302 = e
		var G__4303 = f
		var G__4304 = g
		var G__4305 = h
		var G__4306 = i
		var G__4307 = j
		var G__4308 = k
		var G__4309 = l
		var G__4310 = m
		var G__4311 = n
		_, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__4298, G__4299, G__4300, G__4301, G__4302, G__4303, G__4304, G__4305, G__4306, G__4307, G__4308, G__4309, G__4310, G__4311
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity14(G__4298, G__4299, G__4300, G__4301, G__4302, G__4303, G__4304, G__4305, G__4306, G__4307, G__4308, G__4309, G__4310, G__4311)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var G__4327 = a
		var G__4328 = b
		var G__4329 = c
		var G__4330 = d
		var G__4331 = e
		var G__4332 = f
		var G__4333 = g
		var G__4334 = h
		var G__4335 = i
		var G__4336 = j
		var G__4337 = k
		var G__4338 = l
		var G__4339 = m
		var G__4340 = n
		var G__4341 = o
		_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__4327, G__4328, G__4329, G__4330, G__4331, G__4332, G__4333, G__4334, G__4335, G__4336, G__4337, G__4338, G__4339, G__4340, G__4341
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity15(G__4327, G__4328, G__4329, G__4330, G__4331, G__4332, G__4333, G__4334, G__4335, G__4336, G__4337, G__4338, G__4339, G__4340, G__4341)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var G__4358 = a
		var G__4359 = b
		var G__4360 = c
		var G__4361 = d
		var G__4362 = e
		var G__4363 = f
		var G__4364 = g
		var G__4365 = h
		var G__4366 = i
		var G__4367 = j
		var G__4368 = k
		var G__4369 = l
		var G__4370 = m
		var G__4371 = n
		var G__4372 = o
		var G__4373 = p
		_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__4358, G__4359, G__4360, G__4361, G__4362, G__4363, G__4364, G__4365, G__4366, G__4367, G__4368, G__4369, G__4370, G__4371, G__4372, G__4373
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity16(G__4358, G__4359, G__4360, G__4361, G__4362, G__4363, G__4364, G__4365, G__4366, G__4367, G__4368, G__4369, G__4370, G__4371, G__4372, G__4373)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var G__4391 = a
		var G__4392 = b
		var G__4393 = c
		var G__4394 = d
		var G__4395 = e
		var G__4396 = f
		var G__4397 = g
		var G__4398 = h
		var G__4399 = i
		var G__4400 = j
		var G__4401 = k
		var G__4402 = l
		var G__4403 = m
		var G__4404 = n
		var G__4405 = o
		var G__4406 = p
		var G__4407 = q
		_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__4391, G__4392, G__4393, G__4394, G__4395, G__4396, G__4397, G__4398, G__4399, G__4400, G__4401, G__4402, G__4403, G__4404, G__4405, G__4406, G__4407
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity17(G__4391, G__4392, G__4393, G__4394, G__4395, G__4396, G__4397, G__4398, G__4399, G__4400, G__4401, G__4402, G__4403, G__4404, G__4405, G__4406, G__4407)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	{
		var G__4426 = a
		var G__4427 = b
		var G__4428 = c
		var G__4429 = d
		var G__4430 = e
		var G__4431 = f
		var G__4432 = g
		var G__4433 = h
		var G__4434 = i
		var G__4435 = j
		var G__4436 = k
		var G__4437 = l
		var G__4438 = m
		var G__4439 = n
		var G__4440 = o
		var G__4441 = p
		var G__4442 = q
		var G__4443 = r
		_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__4426, G__4427, G__4428, G__4429, G__4430, G__4431, G__4432, G__4433, G__4434, G__4435, G__4436, G__4437, G__4438, G__4439, G__4440, G__4441, G__4442, G__4443
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity18(G__4426, G__4427, G__4428, G__4429, G__4430, G__4431, G__4432, G__4433, G__4434, G__4435, G__4436, G__4437, G__4438, G__4439, G__4440, G__4441, G__4442, G__4443)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	{
		var G__4463 = a
		var G__4464 = b
		var G__4465 = c
		var G__4466 = d
		var G__4467 = e
		var G__4468 = f
		var G__4469 = g
		var G__4470 = h
		var G__4471 = i
		var G__4472 = j
		var G__4473 = k
		var G__4474 = l
		var G__4475 = m
		var G__4476 = n
		var G__4477 = o
		var G__4478 = p
		var G__4479 = q
		var G__4480 = r
		var G__4481 = s
		_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__4463, G__4464, G__4465, G__4466, G__4467, G__4468, G__4469, G__4470, G__4471, G__4472, G__4473, G__4474, G__4475, G__4476, G__4477, G__4478, G__4479, G__4480, G__4481
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity19(G__4463, G__4464, G__4465, G__4466, G__4467, G__4468, G__4469, G__4470, G__4471, G__4472, G__4473, G__4474, G__4475, G__4476, G__4477, G__4478, G__4479, G__4480, G__4481)
	}
}

func (___ *CljsCoreMetaFn) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	{
		var G__4502 = a
		var G__4503 = b
		var G__4504 = c
		var G__4505 = d
		var G__4506 = e
		var G__4507 = f
		var G__4508 = g
		var G__4509 = h
		var G__4510 = i
		var G__4511 = j
		var G__4512 = k
		var G__4513 = l
		var G__4514 = m
		var G__4515 = n
		var G__4516 = o
		var G__4517 = p
		var G__4518 = q
		var G__4519 = r
		var G__4520 = s
		var G__4521 = t
		_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__4502, G__4503, G__4504, G__4505, G__4506, G__4507, G__4508, G__4509, G__4510, G__4511, G__4512, G__4513, G__4514, G__4515, G__4516, G__4517, G__4518, G__4519, G__4520, G__4521
		return ___.Afn.(CljsCoreIFn).X_invoke_Arity20(G__4502, G__4503, G__4504, G__4505, G__4506, G__4507, G__4508, G__4509, G__4510, G__4511, G__4512, G__4513, G__4514, G__4515, G__4516, G__4517, G__4518, G__4519, G__4520, G__4521)
	}
}

func (_ *CljsCoreMetaFn) CljsCoreFn__() {}

func (_ *CljsCoreMetaFn) CljsCoreIWithMeta__() {}

func (___ *CljsCoreMetaFn) X_with_meta_Arity2(new_meta interface{}) interface{} {
	return (&CljsCoreMetaFn{___.Afn, new_meta})
}

func (_ *CljsCoreMetaFn) CljsCoreIMeta__() {}

func (___ *CljsCoreMetaFn) X_meta_Arity1() interface{} {
	return ___.Meta
}

var X__GT_MetaFn *AFn

// Returns an object of the same type and value as obj, with
// map m as its metadata.
var With_meta *AFn

// Returns the metadata of obj, returns nil if there is no metadata.
var Meta *AFn

// For a list or queue, same as first, for a vector, same as, but much
// more efficient than, last. If the collection is empty, returns nil.
var Peek *AFn

// For a list or queue, returns a new list/queue without the first
// item, for a vector, returns a new vector without the last item.
// Note - not the same as next/butlast.
var Pop *AFn

// disj[oin]. Returns a new set of the same (hashed/sorted) type, that
// does not contain key(s).
// @param {...*} var_args
var Disj *AFn

// Returns true if coll has no items - same as (not (seq coll)).
// Please use the idiom (seq x) rather than (not (empty? x))
var Empty_QMARK_ *AFn

// Returns true if x satisfies ICollection
var Coll_QMARK_ *AFn

// Returns true if x satisfies ISet
var Set_QMARK_ *AFn

// Returns true if coll implements Associative
var Associative_QMARK_ *AFn

// Returns true if coll satisfies ISequential
var Sequential_QMARK_ *AFn

// Returns true if coll satisfies ISorted
var Sorted_QMARK_ *AFn

// Returns true if coll satisfies IReduce
var Reduceable_QMARK_ *AFn

// Return true if x satisfies IMap
var Map_QMARK_ *AFn

// Return true if x satisfies IVector
var Vector_QMARK_ *AFn

var Chunked_seq_QMARK_ *AFn

var Js_delete *AFn

var Array_copy *AFn

var Array_copy_downward *AFn

var Lookup_sentinel bool

// Returns true if x is the value false, false otherwise.
var False_QMARK_ *AFn

// Returns true if x is the value true, false otherwise.
var True_QMARK_ *AFn

var Undefined_QMARK_ *AFn

// Return true if s satisfies ISeq
var Seq_QMARK_ *AFn

// Return true if s satisfies ISeqable
var Seqable_QMARK_ *AFn

var Boolean *AFn

var Ifn_QMARK_ *AFn

// Returns true if key is present in the given collection, otherwise
// returns false.  Note that for numerically indexed collections like
// vectors and arrays, this tests if the numeric key is within the
// range of indexes. 'contains?' operates constant or logarithmic time;
// it will not perform a linear search for a value.  See also 'some'.
var Contains_QMARK_ *AFn

// Returns the map entry for key, or nil if key not present.
var Find *AFn

// Returns true if no two of the arguments are =
// @param {...*} var_args
var Distinct_QMARK_ *AFn

// Coerces coll to a (possibly empty) sequence, if it is not already
// one. Will not force a lazy seq. (sequence nil) yields ()
var Sequence *AFn

// Comparator. Returns a negative number, zero, or a positive number
// when x is logically 'less than', 'equal to', or 'greater than'
// y. Uses IComparable if available and google.array.defaultCompare for objects
// of the same type and special-cases nil to be less than any other object.
var Compare *AFn

// Compare indexed collection.
var Compare_indexed *AFn

// Given a fn that might be boolean valued or a comparator,
// return a fn that is a comparator.
var Fn__GT_comparator *AFn

// Returns a sorted sequence of the items in coll, where the sort
// order is determined by comparing (keyfn item).  Comp can be
// boolean-valued comparison funcion, or a -/0/+ valued comparator.
// Comp defaults to compare.
var Sort_by *AFn

var Seq_reduce *AFn

// Return a random permutation of coll
var Shuffle *AFn

// f should be a function of 2 arguments. If val is not supplied,
// returns the result of applying f to the first 2 items in coll, then
// applying f to that result and the 3rd item, etc. If coll contains no
// items, f must accept no arguments as well, and reduce returns the
// result of calling f with no arguments.  If coll has only 1 item, it
// is returned and f is not called.  If val is supplied, returns the
// result of applying f to val and the first item in coll, then
// applying f to that result and the 2nd item, etc. If coll contains no
// items, returns val and f is not called.
var Reduce *AFn

// Reduces an associative collection. f should be a function of 3
// arguments. Returns the result of applying f to init, the first key
// and the first value in coll, then applying f to that result and the
// 2nd key and value, etc. If coll contains no entries, returns init
// and f is not called. Note that reduce-kv is supported on vectors,
// where the keys will be the ordinals.
var Reduce_kv *AFn

var Completing *AFn

// reduce with a transformation of f (xf). If init is not
// supplied, (f) will be called to produce it. f should be a reducing
// step function that accepts both 1 and 2 arguments, if it accepts
// only 2 you can add the arity-1 with 'completing'. Returns the result
// of applying (the transformed) xf to init and the first item in coll,
// then applying xf to that result and the 2nd item, etc. If coll
// contains no items, returns init and f is not called. Note that
// certain transforms may inject or skip items.
var Transduce *AFn

// Returns the sum of nums. (+) returns 0.
// @param {...*} var_args
var X_PLUS_ *AFn

// If no ys are supplied, returns the negation of x, else subtracts
// the ys from x and returns the result.
// @param {...*} var_args
var X_ *AFn

// Returns the product of nums. (*) returns 1.
// @param {...*} var_args
var X_STAR_ *AFn

// If no denominators are supplied, returns 1/numerator,
// else returns numerator divided by all of the denominators.
// @param {...*} var_args
var X_SLASH_ *AFn

// Returns non-nil if nums are in monotonically increasing order,
// otherwise false.
// @param {...*} var_args
var X_LT_ *AFn

// Returns non-nil if nums are in monotonically non-decreasing order,
// otherwise false.
// @param {...*} var_args
var X_LT__EQ_ *AFn

// Returns non-nil if nums are in monotonically decreasing order,
// otherwise false.
// @param {...*} var_args
var X_GT_ *AFn

// Returns non-nil if nums are in monotonically non-increasing order,
// otherwise false.
// @param {...*} var_args
var X_GT__EQ_ *AFn

// Returns a number one less than num.
var Dec *AFn

// Returns the greatest of the nums.
// @param {...*} var_args
var Max *AFn

// Returns the least of the nums.
// @param {...*} var_args
var Min *AFn

var Byte_ *AFn

var Short *AFn

var Float_ *AFn

var Double *AFn

var Unchecked_byte *AFn

var Unchecked_char *AFn

var Unchecked_short *AFn

var Unchecked_float *AFn

var Unchecked_double *AFn

// Returns the sum of nums. (+) returns 0.
// @param {...*} var_args
var Unchecked_add *AFn

// Returns the sum of nums. (+) returns 0.
// @param {...*} var_args
var Unchecked_add_int *AFn

var Unchecked_dec *AFn

var Unchecked_dec_int *AFn

// If no denominators are supplied, returns 1/numerator,
// else returns numerator divided by all of the denominators.
// @param {...*} var_args
var Unchecked_divide_int *AFn

var Unchecked_inc *AFn

var Unchecked_inc_int *AFn

// Returns the product of nums. (*) returns 1.
// @param {...*} var_args
var Unchecked_multiply *AFn

// Returns the product of nums. (*) returns 1.
// @param {...*} var_args
var Unchecked_multiply_int *AFn

var Unchecked_negate *AFn

var Unchecked_negate_int *AFn

var Unchecked_remainder_int *AFn

// If no ys are supplied, returns the negation of x, else subtracts
// the ys from x and returns the result.
// @param {...*} var_args
var Unchecked_subtract *AFn

// If no ys are supplied, returns the negation of x, else subtracts
// the ys from x and returns the result.
// @param {...*} var_args
var Unchecked_subtract_int *AFn

var Fix *AFn

// Coerce to int by stripping decimal places.
var Int_ *AFn

// Coerce to int by stripping decimal places.
var Unchecked_int *AFn

// Coerce to long by stripping decimal places. Identical to `int'.
var Long *AFn

// Coerce to long by stripping decimal places. Identical to `int'.
var Unchecked_long *AFn

var Booleans *AFn

var Bytes *AFn

var Chars *AFn

var Shorts *AFn

var Ints *AFn

var Floats *AFn

var Doubles *AFn

var Longs *AFn

// Modulus of num and div with original javascript behavior. i.e. bug for negative numbers
var Js_mod *AFn

// Modulus of num and div. Truncates toward negative infinity.
var Mod *AFn

// quot[ient] of dividing numerator by denominator.
var Quot *AFn

// remainder of dividing numerator by denominator.
var Rem *AFn

// Returns a random integer between 0 (inclusive) and n (exclusive).
var Rand_int *AFn

// Bitwise exclusive or
var Bit_xor *AFn

// Bitwise and
var Bit_and *AFn

// Bitwise or
var Bit_or *AFn

// Bitwise and
var Bit_and_not *AFn

// Clear bit at index n
var Bit_clear *AFn

// Flip bit at index n
var Bit_flip *AFn

// Bitwise complement
var Bit_not *AFn

// Set bit at index n
var Bit_set *AFn

// Test bit at index n
var Bit_test *AFn

// Bitwise shift left
var Bit_shift_left *AFn

// Bitwise shift right
var Bit_shift_right *AFn

// DEPRECATED: Bitwise shift right with zero fill
var Bit_shift_right_zero_fill *AFn

// Bitwise shift right with zero fill
var Unsigned_bit_shift_right *AFn

// Counts the number of bits set in n
var Bit_count *AFn

// Returns non-nil if nums all have the equivalent
// value, otherwise false. Behavior on non nums is
// undefined.
// @param {...*} var_args
var X_EQ__EQ_ *AFn

// Returns true if num is greater than zero, else false
var Pos_QMARK_ *AFn

var Zero_QMARK_ *AFn

// Returns true if num is less than zero, else false
var Neg_QMARK_ *AFn

// Returns the nth next of coll, (seq coll) when n is 0.
var Nthnext *AFn

// With no args, returns the empty string. With one arg x, returns
// x.toString().  (str nil) returns the empty string. With more than
// one arg, returns the concatenation of the str values of the args.
// @param {...*} var_args
var Str *AFn

// Returns the substring of s beginning at start inclusive, and ending
// at end (defaults to length of string), exclusive.
var Subs *AFn

// Assumes x is sequential. Returns true if x equals y, otherwise
// returns false.
var Equiv_sequential *AFn

var Hash_coll *AFn

var Hash_imap *AFn

var Hash_iset *AFn

// Takes a JavaScript object and a map of names to functions and
// attaches said functions as methods on the object.  Any references to
// JavaScript's implict this (via the this-as macro) will resolve to the
// object that the function is attached.
var Extend_object_BANG_ *AFn

type CljsCoreList struct {
	Meta    interface{}
	First   interface{}
	Rest    interface{}
	Count   interface{}
	X__hash interface{}
}

func (_ *CljsCoreList) CljsCoreObject__() {}

func (coll *CljsCoreList) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreList) String() string {
	return this.ToString()
}

func (this *CljsCoreList) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreList) CljsCoreIMeta__() {}

func (coll *CljsCoreList) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreList) CljsCoreICloneable__() {}

func (___ *CljsCoreList) X_clone_Arity1() interface{} {
	return (&CljsCoreList{___.Meta, ___.First, ___.Rest, ___.Count, ___.X__hash})
}

func (_ *CljsCoreList) CljsCoreINext__() {}

func (coll *CljsCoreList) X_next_Arity1() interface{} {
	if coll.Count.(float64) == float64(1) {
		return nil
	} else {
		return coll.Rest
	}
}

func (_ *CljsCoreList) CljsCoreICounted__() {}

func (coll *CljsCoreList) X_count_Arity1() float64 {
	return coll.Count.(float64)
}

func (_ *CljsCoreList) CljsCoreIStack__() {}

func (coll *CljsCoreList) X_peek_Arity1() interface{} {
	return coll.First
}

func (coll *CljsCoreList) X_pop_Arity1() interface{} {
	return coll.X_rest_Arity1()
}

func (_ *CljsCoreList) CljsCoreIHash__() {}

func (coll *CljsCoreList) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreList) CljsCoreIEquiv__() {}

func (coll *CljsCoreList) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreList) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreList) X_empty_Arity1() interface{} {
	return CljsCoreList_EMPTY
}

func (_ *CljsCoreList) CljsCoreIReduce__() {}

func (coll *CljsCoreList) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCoreList) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCoreList) CljsCoreISeq__() {}

func (coll *CljsCoreList) X_first_Arity1() interface{} {
	return coll.First
}

func (coll *CljsCoreList) X_rest_Arity1() interface{} {
	if coll.Count.(float64) == float64(1) {
		return CljsCoreIEmptyList(CljsCoreList_EMPTY)
	} else {
		return coll.Rest
	}
}

func (_ *CljsCoreList) CljsCoreISeqable__() {}

func (coll *CljsCoreList) X_seq_Arity1() interface{} {
	return coll
}

func (_ *CljsCoreList) CljsCoreISequential__() {}

func (_ *CljsCoreList) CljsCoreIWithMeta__() {}

func (coll *CljsCoreList) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCoreList{meta___1, coll.First, coll.Rest, coll.Count, coll.X__hash})
}

func (_ *CljsCoreList) CljsCoreICollection__() {}

func (coll *CljsCoreList) X_conj_Arity2(o interface{}) interface{} {
	return (&CljsCoreList{coll.Meta, o, coll, (coll.Count.(float64) + float64(1)), nil})
}

func (_ *CljsCoreList) CljsCoreASeq__() {}

func (_ *CljsCoreList) CljsCoreIList__() {}

var X__GT_List *AFn

type CljsCoreEmptyList struct{ Meta interface{} }

func (_ *CljsCoreEmptyList) CljsCoreObject__() {}

func (coll *CljsCoreEmptyList) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreEmptyList) String() string {
	return this.ToString()
}

func (this *CljsCoreEmptyList) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreEmptyList) CljsCoreIMeta__() {}

func (coll *CljsCoreEmptyList) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreEmptyList) CljsCoreICloneable__() {}

func (___ *CljsCoreEmptyList) X_clone_Arity1() interface{} {
	return (&CljsCoreEmptyList{___.Meta})
}

func (_ *CljsCoreEmptyList) CljsCoreINext__() {}

func (coll *CljsCoreEmptyList) X_next_Arity1() interface{} {
	return nil
}

func (_ *CljsCoreEmptyList) CljsCoreICounted__() {}

func (coll *CljsCoreEmptyList) X_count_Arity1() float64 {
	return float64(0)
}

func (_ *CljsCoreEmptyList) CljsCoreIStack__() {}

func (coll *CljsCoreEmptyList) X_peek_Arity1() interface{} {
	return nil
}

func (coll *CljsCoreEmptyList) X_pop_Arity1() interface{} {
	panic((&js.Error{"Can't pop empty list"}))
}

func (_ *CljsCoreEmptyList) CljsCoreIHash__() {}

func (coll *CljsCoreEmptyList) X_hash_Arity1() interface{} {
	return float64(0)
}

func (_ *CljsCoreEmptyList) CljsCoreIEquiv__() {}

func (coll *CljsCoreEmptyList) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreEmptyList) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreEmptyList) X_empty_Arity1() interface{} {
	return coll
}

func (_ *CljsCoreEmptyList) CljsCoreIReduce__() {}

func (coll *CljsCoreEmptyList) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCoreEmptyList) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCoreEmptyList) CljsCoreISeq__() {}

func (coll *CljsCoreEmptyList) X_first_Arity1() interface{} {
	return nil
}

func (coll *CljsCoreEmptyList) X_rest_Arity1() interface{} {
	return CljsCoreIEmptyList(CljsCoreList_EMPTY)
}

func (_ *CljsCoreEmptyList) CljsCoreISeqable__() {}

func (coll *CljsCoreEmptyList) X_seq_Arity1() interface{} {
	return nil
}

func (_ *CljsCoreEmptyList) CljsCoreISequential__() {}

func (_ *CljsCoreEmptyList) CljsCoreIWithMeta__() {}

func (coll *CljsCoreEmptyList) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCoreEmptyList{meta___1})
}

func (_ *CljsCoreEmptyList) CljsCoreICollection__() {}

func (coll *CljsCoreEmptyList) X_conj_Arity2(o interface{}) interface{} {
	return (&CljsCoreList{coll.Meta, o, nil, float64(1), nil})
}

func (_ *CljsCoreEmptyList) CljsCoreIList__() {}

var X__GT_EmptyList *AFn

var CljsCoreList_EMPTY = (&CljsCoreEmptyList{nil})

var Reversible_QMARK_ *AFn

var Rseq *AFn

// Returns a seq of the items in coll in reverse order. Not lazy.
var Reverse *AFn

// @param {...*} var_args
var List *AFn

type CljsCoreCons struct {
	Meta    interface{}
	First   interface{}
	Rest    interface{}
	X__hash interface{}
}

func (_ *CljsCoreCons) CljsCoreObject__() {}

func (coll *CljsCoreCons) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreCons) String() string {
	return this.ToString()
}

func (this *CljsCoreCons) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreCons) CljsCoreIMeta__() {}

func (coll *CljsCoreCons) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreCons) CljsCoreICloneable__() {}

func (___ *CljsCoreCons) X_clone_Arity1() interface{} {
	return (&CljsCoreCons{___.Meta, ___.First, ___.Rest, ___.X__hash})
}

func (_ *CljsCoreCons) CljsCoreINext__() {}

func (coll *CljsCoreCons) X_next_Arity1() interface{} {
	if Nil_(coll.Rest) {
		return nil
	} else {
		return Seq.Arity1IQ(coll.Rest)
	}
}

func (_ *CljsCoreCons) CljsCoreIHash__() {}

func (coll *CljsCoreCons) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreCons) CljsCoreIEquiv__() {}

func (coll *CljsCoreCons) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreCons) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreCons) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.Meta)
}

func (_ *CljsCoreCons) CljsCoreIReduce__() {}

func (coll *CljsCoreCons) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCoreCons) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCoreCons) CljsCoreISeq__() {}

func (coll *CljsCoreCons) X_first_Arity1() interface{} {
	return coll.First
}

func (coll *CljsCoreCons) X_rest_Arity1() interface{} {
	if Nil_(coll.Rest) {
		return CljsCoreIEmptyList(CljsCoreList_EMPTY)
	} else {
		return coll.Rest
	}
}

func (_ *CljsCoreCons) CljsCoreISeqable__() {}

func (coll *CljsCoreCons) X_seq_Arity1() interface{} {
	return coll
}

func (_ *CljsCoreCons) CljsCoreISequential__() {}

func (_ *CljsCoreCons) CljsCoreIWithMeta__() {}

func (coll *CljsCoreCons) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCoreCons{meta___1, coll.First, coll.Rest, coll.X__hash})
}

func (_ *CljsCoreCons) CljsCoreICollection__() {}

func (coll *CljsCoreCons) X_conj_Arity2(o interface{}) interface{} {
	return (&CljsCoreCons{nil, o, coll, coll.X__hash})
}

func (_ *CljsCoreCons) CljsCoreASeq__() {}

func (_ *CljsCoreCons) CljsCoreIList__() {}

var X__GT_Cons *AFn

// Returns a new seq where x is the first element and seq is the rest.
var Cons *AFn

var List_QMARK_ *AFn

var Hash_keyword *AFn

type CljsCoreKeyword struct {
	Ns     interface{}
	Name   interface{}
	Fqn    interface{}
	X_hash interface{}
}

func (_ *CljsCoreKeyword) CljsCoreIPrintWithWriter__() {}

func (o *CljsCoreKeyword) X_pr_writer_Arity3(writer interface{}, ___ interface{}) interface{} {
	return writer.(CljsCoreIWriter).X_write_Arity2((":" + Str.X_invoke_Arity1(o.Fqn).(string)))
}

func (_ *CljsCoreKeyword) CljsCoreINamed__() {}

func (___ *CljsCoreKeyword) X_name_Arity1() interface{} {
	return ___.Name
}

func (___ *CljsCoreKeyword) X_namespace_Arity1() interface{} {
	return ___.Ns
}

func (_ *CljsCoreKeyword) CljsCoreIHash__() {}

func (this *CljsCoreKeyword) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = this.X_hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_keyword.X_invoke_Arity1(this).(float64)
				_ = h__577__auto_____1
				this.X_hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreKeyword) CljsCoreIFn__() {}

func (this *CljsCoreKeyword) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (kw *CljsCoreKeyword) X_invoke_Arity1(coll interface{}) interface{} {
	return Get.X_invoke_Arity2(coll, kw)
}

func (kw *CljsCoreKeyword) X_invoke_Arity2(coll interface{}, not_found interface{}) interface{} {
	return Get.X_invoke_Arity3(coll, kw, not_found)
}

func (this *CljsCoreKeyword) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

func (_ *CljsCoreKeyword) CljsCoreIEquiv__() {}

func (___ *CljsCoreKeyword) X_equiv_Arity2(other interface{}) bool {
	if Value_(other).Type().AssignableTo(reflect.TypeOf((**CljsCoreKeyword)(nil)).Elem()) {
		return reflect.DeepEqual(___.Fqn, Native_get_instance_field.X_invoke_Arity2(other, "Fqn"))
	} else {
		return false
	}
}

func (_ *CljsCoreKeyword) CljsCoreObject__() {}

func (___ *CljsCoreKeyword) ToString() string {
	return (":" + Str.X_invoke_Arity1(___.Fqn).(string))
}

func (this *CljsCoreKeyword) String() string {
	return this.ToString()
}

func (this *CljsCoreKeyword) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

var X__GT_Keyword *AFn

var Keyword_QMARK_ *AFn

var Keyword_identical_QMARK_ *AFn

// Returns the namespace String of a symbol or keyword, or nil if not present.
var Namespace *AFn

// Returns a Keyword with the given namespace and name.  Do not use :
// in the keyword strings, it will be added automatically.
var Keyword *AFn

type CljsCoreLazySeq struct {
	Meta    interface{}
	Fn      interface{}
	S       interface{}
	X__hash interface{}
}

func (_ *CljsCoreLazySeq) CljsCoreObject__() {}

func (coll *CljsCoreLazySeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreLazySeq) String() string {
	return this.ToString()
}

func (this *CljsCoreLazySeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (coll *CljsCoreLazySeq) Sval() interface{} {
	if Nil_(coll.Fn) {
		return coll.S
	} else {
		coll.S = func() interface{} {
			return coll.Fn.(CljsCoreIFn).X_invoke_Arity0()
		}()

		coll.Fn = nil

		return coll.S
	}
}

func (_ *CljsCoreLazySeq) CljsCoreIMeta__() {}

func (coll *CljsCoreLazySeq) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreLazySeq) CljsCoreINext__() {}

func (coll *CljsCoreLazySeq) X_next_Arity1() interface{} {
	coll.X_seq_Arity1()
	if Nil_(coll.S) {
		return nil
	} else {
		return Next.Arity1IQ(coll.S)
	}
}

func (_ *CljsCoreLazySeq) CljsCoreIHash__() {}

func (coll *CljsCoreLazySeq) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreLazySeq) CljsCoreIEquiv__() {}

func (coll *CljsCoreLazySeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreLazySeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreLazySeq) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.Meta)
}

func (_ *CljsCoreLazySeq) CljsCoreIReduce__() {}

func (coll *CljsCoreLazySeq) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCoreLazySeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCoreLazySeq) CljsCoreISeq__() {}

func (coll *CljsCoreLazySeq) X_first_Arity1() interface{} {
	coll.X_seq_Arity1()
	if Nil_(coll.S) {
		return nil
	} else {
		return First.X_invoke_Arity1(coll.S)
	}
}

func (coll *CljsCoreLazySeq) X_rest_Arity1() interface{} {
	coll.X_seq_Arity1()
	if !(Nil_(coll.S)) {
		return Rest.Arity1IQ(coll.S)
	} else {
		return CljsCoreIEmptyList(CljsCoreList_EMPTY)
	}
}

func (_ *CljsCoreLazySeq) CljsCoreISeqable__() {}

func (coll *CljsCoreLazySeq) X_seq_Arity1() interface{} {
	coll.Sval()
	if Nil_(coll.S) {
		return nil
	} else {
		{
			var ls interface{} = coll.S
			_ = ls
			for {
				if Value_(ls).Type().AssignableTo(reflect.TypeOf((**CljsCoreLazySeq)(nil)).Elem()) {
					ls = Native_invoke_instance_method.X_invoke_Arity3(ls, "Sval", []interface{}{})
					continue
				} else {
					coll.S = ls

					return Seq.Arity1IQ(coll.S)
				}
			}
		}
	}
}

func (_ *CljsCoreLazySeq) CljsCoreISequential__() {}

func (_ *CljsCoreLazySeq) CljsCoreIWithMeta__() {}

func (coll *CljsCoreLazySeq) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCoreLazySeq{meta___1, coll.Fn, coll.S, coll.X__hash})
}

func (_ *CljsCoreLazySeq) CljsCoreICollection__() {}

func (coll *CljsCoreLazySeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

var X__GT_LazySeq *AFn

type CljsCoreChunkBuffer struct {
	Buf interface{}
	End interface{}
}

func (_ *CljsCoreChunkBuffer) CljsCoreICounted__() {}

func (___ *CljsCoreChunkBuffer) X_count_Arity1() float64 {
	return ___.End.(float64)
}

func (_ *CljsCoreChunkBuffer) CljsCoreObject__() {}

func (___ *CljsCoreChunkBuffer) Add(o interface{}) interface{} {
	___.Buf.([]interface{})[int(___.End.(float64))] = o
	return func() interface{} {
		var return__7899 = (___.End.(float64) + float64(1))
		___.End = return__7899
		return return__7899
	}()
}

func (___ *CljsCoreChunkBuffer) Chunk(o interface{}) interface{} {
	{
		var ret = (&CljsCoreArrayChunk{___.Buf, float64(0), ___.End})
		_ = ret
		___.Buf = nil

		return ret
	}
}

var X__GT_ChunkBuffer *AFn

var Chunk_buffer *AFn

type CljsCoreArrayChunk struct {
	Arr interface{}
	Off interface{}
	End interface{}
}

func (_ *CljsCoreArrayChunk) CljsCoreIReduce__() {}

func (coll *CljsCoreArrayChunk) X_reduce_Arity2(f interface{}) interface{} {
	return Array_reduce.X_invoke_Arity4(coll.Arr, f, Aget_(coll.Arr, coll.Off.(float64)), (coll.Off.(float64) + float64(1)))
}

func (coll *CljsCoreArrayChunk) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Array_reduce.X_invoke_Arity4(coll.Arr, f, start, coll.Off)
}

func (_ *CljsCoreArrayChunk) CljsCoreIChunk__() {}

func (coll *CljsCoreArrayChunk) X_drop_first_Arity1() interface{} {
	if coll.Off.(float64) == coll.End.(float64) {
		panic((&js.Error{"-drop-first of empty chunk"}))
	} else {
		return (&CljsCoreArrayChunk{coll.Arr, (coll.Off.(float64) + float64(1)), coll.End})
	}
}

func (_ *CljsCoreArrayChunk) CljsCoreIIndexed__() {}

func (coll *CljsCoreArrayChunk) X_nth_Arity2(i interface{}) interface{} {
	return Aget_(coll.Arr, (coll.Off.(float64) + i.(float64)))
}

func (coll *CljsCoreArrayChunk) X_nth_Arity3(i interface{}, not_found interface{}) interface{} {
	if (i.(float64) >= float64(0)) && (i.(float64) < (coll.End.(float64) - coll.Off.(float64))) {
		return Aget_(coll.Arr, (coll.Off.(float64) + i.(float64)))
	} else {
		return not_found
	}
}

func (_ *CljsCoreArrayChunk) CljsCoreICounted__() {}

func (___ *CljsCoreArrayChunk) X_count_Arity1() float64 {
	return (___.End.(float64) - ___.Off.(float64))
}

var X__GT_ArrayChunk *AFn

var Array_chunk *AFn

type CljsCoreChunkedCons struct {
	Chunk   interface{}
	More    interface{}
	Meta    interface{}
	X__hash interface{}
}

func (_ *CljsCoreChunkedCons) CljsCoreObject__() {}

func (coll *CljsCoreChunkedCons) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreChunkedCons) String() string {
	return this.ToString()
}

func (this *CljsCoreChunkedCons) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreChunkedCons) CljsCoreIMeta__() {}

func (coll *CljsCoreChunkedCons) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreChunkedCons) CljsCoreINext__() {}

func (coll *CljsCoreChunkedCons) X_next_Arity1() interface{} {
	if coll.Chunk.(CljsCoreICounted).X_count_Arity1() > float64(1) {
		return (&CljsCoreChunkedCons{coll.Chunk.(CljsCoreIChunk).X_drop_first_Arity1(), coll.More, coll.Meta, nil})
	} else {
		{
			var more___1 = coll.More.(CljsCoreISeqable).X_seq_Arity1()
			_ = more___1
			if Nil_(more___1) {
				return nil
			} else {
				return more___1
			}
		}
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreIHash__() {}

func (coll *CljsCoreChunkedCons) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreIEquiv__() {}

func (coll *CljsCoreChunkedCons) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreChunkedCons) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreChunkedCons) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.Meta)
}

func (_ *CljsCoreChunkedCons) CljsCoreISeq__() {}

func (coll *CljsCoreChunkedCons) X_first_Arity1() interface{} {
	return coll.Chunk.(CljsCoreIIndexed).X_nth_Arity2(float64(0))
}

func (coll *CljsCoreChunkedCons) X_rest_Arity1() interface{} {
	if coll.Chunk.(CljsCoreICounted).X_count_Arity1() > float64(1) {
		return (&CljsCoreChunkedCons{coll.Chunk.(CljsCoreIChunk).X_drop_first_Arity1(), coll.More, coll.Meta, nil})
	} else {
		if Nil_(coll.More) {
			return CljsCoreIEmptyList(CljsCoreList_EMPTY)
		} else {
			return coll.More
		}
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreISeqable__() {}

func (coll *CljsCoreChunkedCons) X_seq_Arity1() interface{} {
	return coll
}

func (_ *CljsCoreChunkedCons) CljsCoreISequential__() {}

func (_ *CljsCoreChunkedCons) CljsCoreIChunkedSeq__() {}

func (coll *CljsCoreChunkedCons) X_chunked_first_Arity1() interface{} {
	return coll.Chunk
}

func (coll *CljsCoreChunkedCons) X_chunked_rest_Arity1() interface{} {
	if Nil_(coll.More) {
		return CljsCoreIEmptyList(CljsCoreList_EMPTY)
	} else {
		return coll.More
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreIWithMeta__() {}

func (coll *CljsCoreChunkedCons) X_with_meta_Arity2(m interface{}) interface{} {
	return (&CljsCoreChunkedCons{coll.Chunk, coll.More, m, coll.X__hash})
}

func (_ *CljsCoreChunkedCons) CljsCoreICollection__() {}

func (this *CljsCoreChunkedCons) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, this).(*CljsCoreCons)
}

func (_ *CljsCoreChunkedCons) CljsCoreASeq__() {}

func (_ *CljsCoreChunkedCons) CljsCoreIChunkedNext__() {}

func (coll *CljsCoreChunkedCons) X_chunked_next_Arity1() interface{} {
	if Nil_(coll.More) {
		return nil
	} else {
		return coll.More
	}
}

var X__GT_ChunkedCons *AFn

var Chunk_cons *AFn

var Chunk_append *AFn

var Chunk *AFn

var Chunk_first *AFn

var Chunk_rest *AFn

var Chunk_next *AFn

// Naive impl of to-array as a start.
var To_array *AFn

// Returns a (potentially-ragged) 2-dimensional array
// containing the contents of coll.
var To_array_2d *AFn

var Int_array *AFn

var Long_array *AFn

var Double_array *AFn

var Object_array *AFn

var Bounded_count *AFn

var Spread *AFn

// Returns a lazy seq representing the concatenation of the elements in the supplied colls.
// @param {...*} var_args
var Concat *AFn

// Creates a new list containing the items prepended to the rest, the
// last of which will be treated as a sequence.
// @param {...*} var_args
var List_STAR_ *AFn

// Returns a new, transient version of the collection, in constant time.
var Transient *AFn

// Returns a new, persistent version of the transient collection, in
// constant time. The transient collection cannot be used after this
// call, any such use will throw an exception.
var Persistent_BANG_ *AFn

// Adds x to the transient collection, and return coll. The 'addition'
// may happen at different 'places' depending on the concrete type.
// @param {...*} var_args
var Conj_BANG_ *AFn

// When applied to a transient map, adds mapping of key(s) to
// val(s). When applied to a transient vector, sets the val at index.
// Note - index must be <= (count vector). Returns coll.
// @param {...*} var_args
var Assoc_BANG_ *AFn

// Returns a transient map that doesn't contain a mapping for key(s).
// @param {...*} var_args
var Dissoc_BANG_ *AFn

// Removes the last item from a transient vector. If
// the collection is empty, throws an exception. Returns coll
var Pop_BANG_ *AFn

// disj[oin]. Returns a transient set of the same (hashed/sorted) type, that
// does not contain key(s).
// @param {...*} var_args
var Disj_BANG_ *AFn

// Returns an object of the same type and value as obj, with
// (apply f (meta obj) args) as its metadata.
// @param {...*} var_args
var Vary_meta *AFn

// Same as (not (= obj1 obj2))
// @param {...*} var_args
var Not_EQ_ *AFn

// If coll is empty, returns nil, else coll
var Not_empty *AFn

var Nil_iter *AFn

type CljsCoreT4754 struct {
	Nil_iter interface{}
	Meta4755 interface{}
}

func (_ *CljsCoreT4754) CljsCoreObject__() {}

func (___ *CljsCoreT4754) HasNext() interface{} {
	return false
}

func (___ *CljsCoreT4754) Next() interface{} {
	return (&js.Error{"No such element"})
}

func (___ *CljsCoreT4754) Remove() interface{} {
	return (&js.Error{"Unsupported operation"})
}

func (_ *CljsCoreT4754) CljsCoreIMeta__() {}

func (_4756 *CljsCoreT4754) X_meta_Arity1() interface{} {
	return _4756.Meta4755
}

func (_ *CljsCoreT4754) CljsCoreIWithMeta__() {}

func (_4756 *CljsCoreT4754) X_with_meta_Arity2(meta4755___1 interface{}) interface{} {
	return (&CljsCoreT4754{_4756.Nil_iter, meta4755___1})
}

var X__GT_t4754 *AFn

type CljsCoreStringIter struct {
	S interface{}
	I interface{}
}

func (_ *CljsCoreStringIter) CljsCoreObject__() {}

func (___ *CljsCoreStringIter) HasNext() interface{} {
	return (___.I.(float64) < Alength_(___.S))
}

func (___ *CljsCoreStringIter) Next() interface{} {
	{
		var ret = Native_invoke_instance_method.X_invoke_Arity3(___.S, "CharAt", []interface{}{___.I})
		_ = ret
		___.I = (___.I.(float64) + float64(1))

		return ret
	}
}

func (___ *CljsCoreStringIter) Remove() interface{} {
	return (&js.Error{"Unsupported operation"})
}

var X__GT_StringIter *AFn

var String_iter *AFn

type CljsCoreArrayIter struct {
	Arr interface{}
	I   interface{}
}

func (_ *CljsCoreArrayIter) CljsCoreObject__() {}

func (___ *CljsCoreArrayIter) HasNext() interface{} {
	return (___.I.(float64) < Alength_(___.Arr))
}

func (___ *CljsCoreArrayIter) Next() interface{} {
	{
		var ret = Aget_(___.Arr, ___.I.(float64))
		_ = ret
		___.I = (___.I.(float64) + float64(1))

		return ret
	}
}

func (___ *CljsCoreArrayIter) Remove() interface{} {
	return (&js.Error{"Unsupported operation"})
}

var X__GT_ArrayIter *AFn

var Array_iter *AFn

var INIT interface{}

var START interface{}

type CljsCoreSeqIter struct {
	X_seq  interface{}
	X_next interface{}
}

func (_ *CljsCoreSeqIter) CljsCoreObject__() {}

func (___ *CljsCoreSeqIter) HasNext() interface{} {
	if reflect.DeepEqual(___.X_seq, INIT) {
		___.X_seq = START

		___.X_next = Seq.Arity1IQ(___.X_next)

	} else {
		if reflect.DeepEqual(___.X_seq, ___.X_next) {
			___.X_next = Next.Arity1IQ(___.X_seq)

		} else {
		}
	}
	return !(Nil_(___.X_next))
}

func (this *CljsCoreSeqIter) Next() interface{} {
	if Not.Arity1IB(this.HasNext()) {
		panic((&js.Error{"No such element"}))
	} else {
		this.X_seq = this.X_next

		return First.X_invoke_Arity1(this.X_next)
	}
}

func (___ *CljsCoreSeqIter) Remove() interface{} {
	return (&js.Error{"Unsupported operation"})
}

var X__GT_SeqIter *AFn

var Seq_iter *AFn

var Iter *AFn

var Lazy_transformer *AFn

type CljsCoreStepper struct {
	Xform interface{}
	Iter  interface{}
}

func (_ *CljsCoreStepper) CljsCoreObject__() {}

func (this *CljsCoreStepper) Step(lt interface{}) interface{} {
	{
		for {
			if Truth_(func() interface{} {
				var and__159__auto__ = !(Nil_(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper")))
				_ = and__159__auto__
				if Truth_(and__159__auto__) {
					return Native_invoke_instance_method.X_invoke_Arity3(this.Iter, "HasNext", []interface{}{})
				} else {
					return and__159__auto__
				}
			}()) {
				if Reduced_QMARK_.Arity1IB(func() interface{} {
					var G__4760 = lt
					var G__4761 = Native_invoke_instance_method.X_invoke_Arity3(this.Iter, "Next", []interface{}{})
					_, _ = G__4760, G__4761
					return this.Xform.(CljsCoreIFn).X_invoke_Arity2(G__4760, G__4761)
				}()) {
					if Nil_(Native_get_instance_field.X_invoke_Arity2(lt, "Rest")) {
					} else {
						Native_set_instance_field.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(lt, "Rest"), "Stepper", nil)
					}
				} else {
					continue
				}
			} else {
			}
			break
		}
	}
	if Nil_(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper")) {
		return nil
	} else {
		{
			var G__4762 = lt
			_ = G__4762
			return this.Xform.(CljsCoreIFn).X_invoke_Arity1(G__4762)
		}
	}
}

var X__GT_Stepper *AFn

var Stepper *AFn

type CljsCoreMultiStepper struct {
	Xform interface{}
	Iters interface{}
	Nexts interface{}
}

func (_ *CljsCoreMultiStepper) CljsCoreObject__() {}

func (___ *CljsCoreMultiStepper) HasNext() interface{} {
	{
		var iters___1 interface{} = Seq.Arity1IQ(___.Iters)
		_ = iters___1
		for {
			if !(Nil_(iters___1)) {
				{
					var iter = First.X_invoke_Arity1(iters___1)
					_ = iter
					if Not.Arity1IB(Native_invoke_instance_method.X_invoke_Arity3(iter, "HasNext", []interface{}{})) {
						return false
					} else {
						iters___1 = Next.Arity1IQ(iters___1)
						continue
					}
				}
			} else {
				return true
			}
		}
	}
}

func (___ *CljsCoreMultiStepper) Next() interface{} {
	{
		var n__1054__auto___7900 = Alength_(___.Iters)
		_ = n__1054__auto___7900
		{
			var i_7901 = float64(0)
			_ = i_7901
			for {
				if i_7901 < n__1054__auto___7900 {
					___.Nexts.([]interface{})[int(i_7901)] = Native_invoke_instance_method.X_invoke_Arity3(Aget_(___.Iters, i_7901), "Next", []interface{}{})
					i_7901 = (i_7901 + float64(1))
					continue
				} else {
				}
				break
			}
		}
	}
	return Prim_seq.X_invoke_Arity2(___.Nexts, float64(0))
}

func (this *CljsCoreMultiStepper) Step(lt interface{}) interface{} {
	{
		for {
			if Truth_(func() interface{} {
				var and__159__auto__ = !(Nil_(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper")))
				_ = and__159__auto__
				if Truth_(and__159__auto__) {
					return this.HasNext()
				} else {
					return and__159__auto__
				}
			}()) {
				if Reduced_QMARK_.Arity1IB(Apply.X_invoke_Arity2(this.Xform, Cons.X_invoke_Arity2(lt, this.Next()).(*CljsCoreCons))) {
					if Nil_(Native_get_instance_field.X_invoke_Arity2(lt, "Rest")) {
					} else {
						Native_set_instance_field.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(lt, "Rest"), "Stepper", nil)
					}
				} else {
					continue
				}
			} else {
			}
			break
		}
	}
	if Nil_(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper")) {
		return nil
	} else {
		{
			var G__4766 = lt
			_ = G__4766
			return this.Xform.(CljsCoreIFn).X_invoke_Arity1(G__4766)
		}
	}
}

var X__GT_MultiStepper *AFn

var Multi_stepper *AFn

type CljsCoreLazyTransformer struct {
	Stepper interface{}
	First   interface{}
	Rest    interface{}
	Meta    interface{}
}

func (_ *CljsCoreLazyTransformer) CljsCoreINext__() {}

func (this *CljsCoreLazyTransformer) X_next_Arity1() interface{} {
	if Nil_(this.Stepper) {
	} else {
		this.X_seq_Arity1()
	}
	if Nil_(this.Rest) {
		return nil
	} else {
		return this.Rest.(CljsCoreISeqable).X_seq_Arity1()
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreISeq__() {}

func (this *CljsCoreLazyTransformer) X_first_Arity1() interface{} {
	if Nil_(this.Stepper) {
	} else {
		this.X_seq_Arity1()
	}
	if Nil_(this.Rest) {
		return nil
	} else {
		return this.First
	}
}

func (this *CljsCoreLazyTransformer) X_rest_Arity1() interface{} {
	if Nil_(this.Stepper) {
	} else {
		this.X_seq_Arity1()
	}
	if Nil_(this.Rest) {
		return CljsCoreIEmptyList(CljsCoreList_EMPTY)
	} else {
		return this.Rest
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreISeqable__() {}

func (this *CljsCoreLazyTransformer) X_seq_Arity1() interface{} {
	if Nil_(this.Stepper) {
	} else {
		Native_invoke_instance_method.X_invoke_Arity3(this.Stepper, "Step", []interface{}{this})
	}
	if Nil_(this.Rest) {
		return nil
	} else {
		return this
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreIHash__() {}

func (this *CljsCoreLazyTransformer) X_hash_Arity1() interface{} {
	return Hash_ordered_coll.Arity1IF(this)
}

func (_ *CljsCoreLazyTransformer) CljsCoreIEquiv__() {}

func (this *CljsCoreLazyTransformer) X_equiv_Arity2(other interface{}) bool {
	{
		var s = this.X_seq_Arity1()
		_ = s
		if !(Nil_(s)) {
			return Truth_(Equiv_sequential.X_invoke_Arity2(this, other))
		} else {
			return (Sequential_QMARK_.Arity1IB(other)) && (Nil_(Seq.Arity1IQ(other)))
		}
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreISequential__() {}

func (_ *CljsCoreLazyTransformer) CljsCoreIEmptyableCollection__() {}

func (this *CljsCoreLazyTransformer) X_empty_Arity1() interface{} {
	return CljsCoreIEmptyList(CljsCoreList_EMPTY)
}

func (_ *CljsCoreLazyTransformer) CljsCoreICollection__() {}

func (this *CljsCoreLazyTransformer) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, this.X_seq_Arity1()).(*CljsCoreCons)
}

func (_ *CljsCoreLazyTransformer) CljsCoreIWithMeta__() {}

func (this *CljsCoreLazyTransformer) X_with_meta_Arity2(new_meta interface{}) interface{} {
	return (&CljsCoreLazyTransformer{this.Stepper, this.First, this.Rest, new_meta})
}

var X__GT_LazyTransformer *AFn

var CljsCoreLazyTransformer_Create = func(G__7902 *AFn) *AFn {
	return Fn(G__7902, 2, func(xform interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazyTransformer{Stepper.X_invoke_Arity2(xform, Iter.X_invoke_Arity1(coll)), nil, nil, nil})
	})
}(&AFn{})

var CljsCoreLazyTransformer_CreateMulti = func(G__7903 *AFn) *AFn {
	return Fn(G__7903, 2, func(xform interface{}, colls interface{}) interface{} {
		{
			var iters = []interface{}{}
			_ = iters
			{
				var seq__4773_7904 interface{} = Seq.Arity1IQ(colls)
				var chunk__4774_7905 interface{} = nil
				var count__4775_7906 = float64(0)
				var i__4776_7907 = float64(0)
				_, _, _, _ = seq__4773_7904, chunk__4774_7905, count__4775_7906, i__4776_7907
				for {
					if i__4776_7907 < count__4775_7906 {
						{
							var coll_7908 = chunk__4774_7905.(CljsCoreIIndexed).X_nth_Arity2(i__4776_7907)
							_ = coll_7908
							js.JSArray_(&iters).Push(Iter.X_invoke_Arity1(coll_7908))
							seq__4773_7904, chunk__4774_7905, count__4775_7906, i__4776_7907 = seq__4773_7904, chunk__4774_7905, count__4775_7906, (i__4776_7907 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___7909 = Seq.Arity1IQ(seq__4773_7904)
							_ = temp__4222__auto___7909
							if Truth_(temp__4222__auto___7909) {
								{
									var seq__4773_7910___1 = temp__4222__auto___7909
									_ = seq__4773_7910___1
									if Chunked_seq_QMARK_.Arity1IB(seq__4773_7910___1) {
										{
											var c__954__auto___7911 = Chunk_first.X_invoke_Arity1(seq__4773_7910___1)
											_ = c__954__auto___7911
											seq__4773_7904, chunk__4774_7905, count__4775_7906, i__4776_7907 = Chunk_rest.X_invoke_Arity1(seq__4773_7910___1), c__954__auto___7911, Count.X_invoke_Arity1(c__954__auto___7911).(float64), float64(0)
											continue
										}
									} else {
										{
											var coll_7912 = First.X_invoke_Arity1(seq__4773_7910___1)
											_ = coll_7912
											js.JSArray_(&iters).Push(Iter.X_invoke_Arity1(coll_7912))
											seq__4773_7904, chunk__4774_7905, count__4775_7906, i__4776_7907 = Next.Arity1IQ(seq__4773_7910___1), nil, float64(0), float64(0)
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
			return (&CljsCoreLazyTransformer{Multi_stepper.X_invoke_Arity3(xform, iters, make([]interface{}, int(float64(len(iters))))), nil, nil, nil})
		}
	})
}(&AFn{})

// Returns true if (pred x) is logical true for every x in coll, else
// false.
var Every_QMARK_ *AFn

// Returns false if (pred x) is logical true for every x in
// coll, else true.
var Not_every_QMARK_ *AFn

// Returns the first logical true value of (pred x) for any x in coll,
// else nil.  One common idiom is to use a set as pred, for example
// this will return :fred if :fred is in the sequence, otherwise nil:
// (some #{:fred} coll)
var Some *AFn

// Returns false if (pred x) is logical true for any x in coll,
// else true.
var Not_any_QMARK_ *AFn

// Returns true if n is even, throws an exception if n is not an integer
var Even_QMARK_ *AFn

// Returns true if n is odd, throws an exception if n is not an integer
var Odd_QMARK_ *AFn

// Returns a function that takes any number of arguments and returns x.
var Constantly *AFn

// Takes a set of functions and returns a fn that is the composition
// of those fns.  The returned fn takes a variable number of args,
// applies the rightmost of fns to the args, the next
// fn (right-to-left) to the result, etc.
// @param {...*} var_args
var Comp *AFn

// Takes a function f and fewer than the normal arguments to f, and
// returns a fn that takes a variable number of additional args. When
// called, the returned function calls f with args + additional args.
// @param {...*} var_args
var Partial *AFn

// Takes a function f, and returns a function that calls f, replacing
// a nil first argument to f with the supplied value x. Higher arity
// versions can replace arguments in the second and third
// positions (y, z). Note that the function f can take any number of
// arguments, not just the one(s) being nil-patched.
var Fnil *AFn

// Returns a lazy sequence consisting of the result of applying f to 0
// and the first item of coll, followed by applying f to 1 and the second
// item in coll, etc, until coll is exhausted. Thus function f should
// accept 2 arguments, index and item.
var Map_indexed *AFn

// Returns a lazy sequence of the non-nil results of (f item). Note,
// this means false return values will be included.  f must be free of
// side-effects.  Returns a transducer when no collection is provided.
var Keep *AFn

type CljsCoreAtom struct {
	State     interface{}
	Meta      interface{}
	Validator interface{}
	Watches   interface{}
}

func (_ *CljsCoreAtom) CljsCoreIHash__() {}

func (this *CljsCoreAtom) X_hash_Arity1() interface{} {
	{
		var G__4972 = this
		_ = G__4972
		return Native_invoke_func.X_invoke_Arity2(goog.GetUid, []interface{}{G__4972})
	}
}

func (_ *CljsCoreAtom) CljsCoreIWatchable__() {}

func (this *CljsCoreAtom) X_notify_watches_Arity3(oldval interface{}, newval interface{}) interface{} {
	{
		var seq__4987 interface{} = Seq.Arity1IQ(this.Watches)
		var chunk__4988 interface{} = nil
		var count__4989 = float64(0)
		var i__4990 = float64(0)
		_, _, _, _ = seq__4987, chunk__4988, count__4989, i__4990
		for {
			if i__4990 < count__4989 {
				{
					var vec__4991 = chunk__4988.(CljsCoreIIndexed).X_nth_Arity2(i__4990)
					var key = Nth.X_invoke_Arity3(vec__4991, float64(0), nil)
					var f = Nth.X_invoke_Arity3(vec__4991, float64(1), nil)
					_, _, _ = vec__4991, key, f
					{
						var G__4992_7913 = key
						var G__4993_7914 = this
						var G__4994_7915 = oldval
						var G__4995_7916 = newval
						_, _, _, _ = G__4992_7913, G__4993_7914, G__4994_7915, G__4995_7916
						f.(CljsCoreIFn).X_invoke_Arity4(G__4992_7913, G__4993_7914, G__4994_7915, G__4995_7916)
					}
					seq__4987, chunk__4988, count__4989, i__4990 = seq__4987, chunk__4988, count__4989, (i__4990 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__4987)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__4987___1 = temp__4222__auto__
							_ = seq__4987___1
							if Chunked_seq_QMARK_.Arity1IB(seq__4987___1) {
								{
									var c__954__auto__ = Chunk_first.X_invoke_Arity1(seq__4987___1)
									_ = c__954__auto__
									seq__4987, chunk__4988, count__4989, i__4990 = Chunk_rest.X_invoke_Arity1(seq__4987___1), c__954__auto__, Count.X_invoke_Arity1(c__954__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__4996 = First.X_invoke_Arity1(seq__4987___1)
									var key = Nth.X_invoke_Arity3(vec__4996, float64(0), nil)
									var f = Nth.X_invoke_Arity3(vec__4996, float64(1), nil)
									_, _, _ = vec__4996, key, f
									{
										var G__4997_7917 = key
										var G__4998_7918 = this
										var G__4999_7919 = oldval
										var G__5000_7920 = newval
										_, _, _, _ = G__4997_7917, G__4998_7918, G__4999_7919, G__5000_7920
										f.(CljsCoreIFn).X_invoke_Arity4(G__4997_7917, G__4998_7918, G__4999_7919, G__5000_7920)
									}
									seq__4987, chunk__4988, count__4989, i__4990 = Next.Arity1IQ(seq__4987___1), nil, float64(0), float64(0)
									continue
								}
							}
						}
					} else {
						return nil
					}
				}
			}
		}
	}
}

func (this *CljsCoreAtom) X_add_watch_Arity3(key interface{}, f interface{}) interface{} {
	this.Watches = Assoc.X_invoke_Arity3(this.Watches, key, f)

	return this
}

func (this *CljsCoreAtom) X_remove_watch_Arity2(key interface{}) interface{} {
	return func() interface{} {
		var return__7921 = Dissoc.X_invoke_Arity2(this.Watches, key)
		this.Watches = return__7921
		return return__7921
	}()
}

func (_ *CljsCoreAtom) CljsCoreIMeta__() {}

func (___ *CljsCoreAtom) X_meta_Arity1() interface{} {
	return ___.Meta
}

func (_ *CljsCoreAtom) CljsCoreIDeref__() {}

func (___ *CljsCoreAtom) X_deref_Arity1() interface{} {
	return ___.State
}

func (_ *CljsCoreAtom) CljsCoreIEquiv__() {}

func (o *CljsCoreAtom) X_equiv_Arity2(other interface{}) bool {
	return reflect.DeepEqual(o, other)
}

func (_ *CljsCoreAtom) CljsCoreIAtom__() {}

func (_ *CljsCoreAtom) CljsCoreObject__() {}

func (this *CljsCoreAtom) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

var X__GT_Atom *AFn

// Creates and returns an Atom with an initial value of x and zero or
// more options (in any order):
//
// :meta metadata-map
//
// :validator validate-fn
//
// If metadata-map is supplied, it will be come the metadata on the
// atom. validate-fn must be nil or a side-effect-free fn of one
// argument, which will be passed the intended new state on any state
// change. If the new state is unacceptable, the validate-fn should
// return false or throw an Error.  If either of these error conditions
// occur, then the value of the atom will not change.
// @param {...*} var_args
var Atom *AFn

// Sets the value of atom to newval without regard for the
// current value. Returns newval.
var Reset_BANG_ *AFn

// Atomically swaps the value of atom to be:
// (apply f current-value-of-atom args). Note that f may be called
// multiple times, and thus should be free of side effects.  Returns
// the value that was swapped in.
// @param {...*} var_args
var Swap_BANG_ *AFn

// Atomically sets the value of atom to newval if and only if the
// current value of the atom is identical to oldval. Returns true if
// set happened, else false.
var Compare_and_set_BANG_ *AFn

// Sets the validator-fn for an atom. validator-fn must be nil or a
// side-effect-free fn of one argument, which will be passed the intended
// new state on any state change. If the new state is unacceptable, the
// validator-fn should return false or throw an Error. If the current state
// is not acceptable to the new validator, an Error will be thrown and the
// validator will not be changed.
var Set_validator_BANG_ *AFn

// Gets the validator-fn for a var/ref/agent/atom.
var Get_validator *AFn

// Returns a lazy sequence of the non-nil results of (f index item). Note,
// this means false return values will be included.  f must be free of
// side-effects.  Returns a stateful transducer when no collection is
// provided.
var Keep_indexed *AFn

// Takes a set of predicates and returns a function f that returns true if all of its
// composing predicates return a logical true value against all of its arguments, else it returns
// false. Note that f is short-circuiting in that it will stop execution on the first
// argument that triggers a logical false result against the original predicates.
// @param {...*} var_args
var Every_pred *AFn

// Takes a set of predicates and returns a function f that returns the first logical true value
// returned by one of its composing predicates against any of its arguments, else it returns
// logical false. Note that f is short-circuiting in that it will stop execution on the first
// argument that triggers a logical true result against the original predicates.
// @param {...*} var_args
var Some_fn *AFn

// Returns a lazy sequence consisting of the result of applying f to
// the set of first items of each coll, followed by applying f to the
// set of second items in each coll, until any one of the colls is
// exhausted.  Any remaining items in other colls are ignored. Function
// f should accept number-of-colls arguments. Returns a transducer when
// no collection is provided.
// @param {...*} var_args
var Map_ *AFn

// Returns a lazy sequence of the first n items in coll, or all items if
// there are fewer than n.  Returns a stateful transducer when
// no collection is provided.
var Take *AFn

// Returns a lazy sequence of all but the first n items in coll.
// Returns a stateful transducer when no collection is provided.
var Drop *AFn

// Return a lazy sequence of all but the last n (default 1) items in coll
var Drop_last *AFn

// Returns a seq of the last n items in coll.  Depending on the type
// of coll may be no better than linear time.  For vectors, see also subvec.
var Take_last *AFn

// Returns a lazy sequence of the items in coll starting from the
// first item for which (pred item) returns logical false.  Returns a
// stateful transducer when no collection is provided.
var Drop_while *AFn

// Returns a lazy (infinite!) sequence of repetitions of the items in coll.
var Cycle *AFn

// Returns a vector of [(take n coll) (drop n coll)]
var Split_at *AFn

// Returns a lazy (infinite!, or length n if supplied) sequence of xs.
var Repeat *AFn

// Returns a lazy seq of n xs.
var Replicate *AFn

// Takes a function of no args, presumably with side effects, and
// returns an infinite (or length n if supplied) lazy sequence of calls
// to it
var Repeatedly *AFn

// Returns a lazy sequence of x, (f x), (f (f x)) etc. f must be free of side-effects
var Iterate *AFn

// Returns a lazy seq of the first item in each coll, then the second etc.
// @param {...*} var_args
var Interleave *AFn

// Returns a lazy seq of the elements of coll separated by sep
var Interpose *AFn

// Take a collection of collections, and return a lazy seq
// of items from the inner collection
var Flatten1 *AFn

// Returns the result of applying concat to the result of applying map
// to f and colls.  Thus function f should return a collection. Returns
// a transducer when no collections are provided
// @param {...*} var_args
var Mapcat *AFn

// Returns a lazy sequence of the items in coll for which
// (pred item) returns true. pred must be free of side-effects.
// Returns a transducer when no collection is provided.
var Filter *AFn

// Returns a lazy sequence of the nodes in a tree, via a depth-first walk.
// branch? must be a fn of one arg that returns true if passed a node
// that can have children (but may not).  children must be a fn of one
// arg that returns a sequence of the children. Will only be called on
// nodes for which branch? returns true. Root is the root node of the
// tree.
var Tree_seq *AFn

// Takes any nested combination of sequential things (lists, vectors,
// etc.) and returns their contents as a single, flat sequence.
// (flatten nil) returns nil.
var Flatten *AFn

// Returns a new coll consisting of to-coll with all of the items of
// from-coll conjoined. A transducer may be supplied.
var Into *AFn

// Returns a vector consisting of the result of applying f to the
// set of first items of each coll, followed by applying f to the set
// of second items in each coll, until any one of the colls is
// exhausted.  Any remaining items in other colls are ignored. Function
// f should accept number-of-colls arguments.
// @param {...*} var_args
var Mapv *AFn

// Returns a vector of the items in coll for which
// (pred item) returns true. pred must be free of side-effects.
var Filterv *AFn

// Returns a lazy sequence of lists of n items each, at offsets step
// apart. If step is not supplied, defaults to n, i.e. the partitions
// do not overlap. If a pad collection is supplied, use its elements as
// necessary to complete last partition upto n items. In case there are
// not enough padding elements, return a partition with less than n items.
var Partition *AFn

// Returns the value in a nested associative structure,
// where ks is a sequence of keys. Returns nil if the key is not present,
// or the not-found value if supplied.
var Get_in *AFn

// Associates a value in a nested associative structure, where ks is a
// sequence of keys and v is the new value and returns a new nested structure.
// If any levels do not exist, hash-maps will be created.
var Assoc_in *AFn

// 'Updates' a value in a nested associative structure, where ks is a
// sequence of keys and f is a function that will take the old value
// and any supplied args and return the new value, and returns a new
// nested structure.  If any levels do not exist, hash-maps will be
// created.
// @param {...*} var_args
var Update_in *AFn

type CljsCoreVectorNode struct {
	Edit interface{}
	Arr  interface{}
}

var X__GT_VectorNode *AFn

var Pv_fresh_node *AFn

var Pv_aget *AFn

var Pv_aset *AFn

var Pv_clone_node *AFn

var Tail_off *AFn

var New_path *AFn

var Push_tail *AFn

var Vector_index_out_of_bounds *AFn

var First_array_for_longvec *AFn

var Unchecked_array_for *AFn

var Array_for *AFn

var Do_assoc *AFn

var Pop_tail *AFn

type CljsCorePersistentVector struct {
	Meta    interface{}
	Cnt     interface{}
	Shift   interface{}
	Root    interface{}
	Tail    interface{}
	X__hash interface{}
}

func (_ *CljsCorePersistentVector) CljsCoreObject__() {}

func (coll *CljsCorePersistentVector) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCorePersistentVector) String() string {
	return this.ToString()
}

func (this *CljsCorePersistentVector) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCorePersistentVector) CljsCoreILookup__() {}

func (coll *CljsCorePersistentVector) X_lookup_Arity2(k interface{}) interface{} {
	return coll.X_lookup_Arity3(k, nil)
}

func (coll *CljsCorePersistentVector) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	if Value_(k).Kind() == reflect.Float64 {
		return coll.X_nth_Arity3(k, not_found)
	} else {
		return not_found
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIKVReduce__() {}

func (v *CljsCorePersistentVector) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var i = float64(0)
		var init___1 interface{} = init
		_, _ = i, init___1
		for {
			if i < v.Cnt.(float64) {
				{
					var arr = Unchecked_array_for.X_invoke_Arity2(v, i)
					var len = Alength_(arr)
					var init___2 = func() interface{} {
						var j = float64(0)
						var init___2 interface{} = init___1
						_, _ = j, init___2
						for {
							if j < len {
								{
									var init___3 = func() interface{} {
										var G__6126 = init___2
										var G__6127 = (j + i)
										var G__6128 = Aget_(arr, j)
										_, _, _ = G__6126, G__6127, G__6128
										return f.(CljsCoreIFn).X_invoke_Arity3(G__6126, G__6127, G__6128)
									}()
									_ = init___3
									if Reduced_QMARK_.Arity1IB(init___3) {
										return init___3
									} else {
										j, init___2 = (j + float64(1)), init___3
										continue
									}
								}
							} else {
								return init___2
							}
						}
					}()
					_, _, _ = arr, len, init___2
					if Reduced_QMARK_.Arity1IB(init___2) {
						return Deref.X_invoke_Arity1(init___2)
					} else {
						i, init___1 = (i + len), init___2
						continue
					}
				}
			} else {
				return init___1
			}
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIIndexed__() {}

func (coll *CljsCorePersistentVector) X_nth_Arity2(n interface{}) interface{} {
	return Aget_(Array_for.X_invoke_Arity2(coll, n), float64((Int32_(n.(float64)) & Int32_(float64(31)))))
}

func (coll *CljsCorePersistentVector) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	if (float64(0) <= n.(float64)) && (n.(float64) < coll.Cnt.(float64)) {
		return Aget_(Unchecked_array_for.X_invoke_Arity2(coll, n), float64((Int32_(n.(float64)) & Int32_(float64(31)))))
	} else {
		return not_found
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIVector__() {}

func (coll *CljsCorePersistentVector) X_assoc_n_Arity3(n interface{}, val interface{}) interface{} {
	if (float64(0) <= n.(float64)) && (n.(float64) < coll.Cnt.(float64)) {
		if Tail_off.X_invoke_Arity1(coll).(float64) <= n.(float64) {
			{
				var new_tail = Aclone.X_invoke_Arity1(coll.Tail).([]interface{})
				_ = new_tail
				new_tail[int(float64((Int32_(n.(float64)) & Int32_(float64(31)))))] = val
				return (&CljsCorePersistentVector{coll.Meta, coll.Cnt, coll.Shift, coll.Root, new_tail, nil})
			}
		} else {
			return (&CljsCorePersistentVector{coll.Meta, coll.Cnt, coll.Shift, Do_assoc.X_invoke_Arity5(coll, coll.Shift, coll.Root, n, val).(*CljsCoreVectorNode), coll.Tail, nil})
		}
	} else {
		if n.(float64) == coll.Cnt.(float64) {
			return coll.X_conj_Arity2(val)
		} else {
			panic((&js.Error{("Index " + Str.X_invoke_Arity1(n).(string) + " out of bounds  [0," + Str.X_invoke_Arity1(coll.Cnt).(string) + "]")}))

		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIMeta__() {}

func (coll *CljsCorePersistentVector) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCorePersistentVector) CljsCoreICloneable__() {}

func (___ *CljsCorePersistentVector) X_clone_Arity1() interface{} {
	return (&CljsCorePersistentVector{___.Meta, ___.Cnt, ___.Shift, ___.Root, ___.Tail, ___.X__hash})
}

func (_ *CljsCorePersistentVector) CljsCoreICounted__() {}

func (coll *CljsCorePersistentVector) X_count_Arity1() float64 {
	return coll.Cnt.(float64)
}

func (_ *CljsCorePersistentVector) CljsCoreIMapEntry__() {}

func (coll *CljsCorePersistentVector) X_key_Arity1() interface{} {
	return coll.X_nth_Arity2(float64(0))
}

func (coll *CljsCorePersistentVector) X_val_Arity1() interface{} {
	return coll.X_nth_Arity2(float64(1))
}

func (_ *CljsCorePersistentVector) CljsCoreIStack__() {}

func (coll *CljsCorePersistentVector) X_peek_Arity1() interface{} {
	if coll.Cnt.(float64) > float64(0) {
		return coll.X_nth_Arity2((coll.Cnt.(float64) - float64(1)))
	} else {
		return nil
	}
}

func (coll *CljsCorePersistentVector) X_pop_Arity1() interface{} {
	if coll.Cnt.(float64) == float64(0) {
		panic((&js.Error{"Can't pop empty vector"}))
	} else {
		if float64(1) == coll.Cnt.(float64) {
			return CljsCorePersistentVector_EMPTY.X_with_meta_Arity2(coll.Meta)
		} else {
			if float64(1) < (coll.Cnt.(float64) - Tail_off.X_invoke_Arity1(coll).(float64)) {
				return (&CljsCorePersistentVector{coll.Meta, (coll.Cnt.(float64) - float64(1)), coll.Shift, coll.Root, js.JSArray_(&coll.Tail).Slice(float64(0), float64(-1)), nil})
			} else {
				{
					var new_tail = Unchecked_array_for.X_invoke_Arity2(coll, (coll.Cnt.(float64) - float64(2)))
					var nr = Pop_tail.X_invoke_Arity3(coll, coll.Shift, coll.Root)
					var new_root = func() interface{} {
						if Nil_(nr) {
							return CljsCorePersistentVector_EMPTY_NODE
						} else {
							return nr
						}
					}()
					var cnt_1 = (coll.Cnt.(float64) - float64(1))
					_, _, _, _ = new_tail, nr, new_root, cnt_1
					if (float64(5) < coll.Shift.(float64)) && (Nil_(Pv_aget.X_invoke_Arity2(new_root, float64(1)))) {
						return (&CljsCorePersistentVector{coll.Meta, cnt_1, (coll.Shift.(float64) - float64(5)), Pv_aget.X_invoke_Arity2(new_root, float64(0)), new_tail, nil})
					} else {
						return (&CljsCorePersistentVector{coll.Meta, cnt_1, coll.Shift, new_root, new_tail, nil})
					}
				}

			}
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIReversible__() {}

func (coll *CljsCorePersistentVector) X_rseq_Arity1() interface{} {
	if coll.Cnt.(float64) > float64(0) {
		return (&CljsCoreRSeq{coll, (coll.Cnt.(float64) - float64(1)), nil})
	} else {
		return nil
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIHash__() {}

func (coll *CljsCorePersistentVector) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentVector) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCorePersistentVector) CljsCoreIEditableCollection__() {}

func (coll *CljsCorePersistentVector) X_as_transient_Arity1() interface{} {
	return (&CljsCoreTransientVector{coll.Cnt, coll.Shift, Tv_editable_root.X_invoke_Arity1(coll.Root).(*CljsCoreVectorNode), Tv_editable_tail.X_invoke_Arity1(coll.Tail).([]interface{})})
}

func (_ *CljsCorePersistentVector) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCorePersistentVector) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCorePersistentVector_EMPTY, coll.Meta)
}

func (_ *CljsCorePersistentVector) CljsCoreIReduce__() {}

func (v *CljsCorePersistentVector) X_reduce_Arity2(f interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity2(v, f)
}

func (v *CljsCorePersistentVector) X_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var i = float64(0)
		var init___1 interface{} = init
		_, _ = i, init___1
		for {
			if i < v.Cnt.(float64) {
				{
					var arr = Unchecked_array_for.X_invoke_Arity2(v, i)
					var len = Alength_(arr)
					var init___2 = func() interface{} {
						var j = float64(0)
						var init___2 interface{} = init___1
						_, _ = j, init___2
						for {
							if j < len {
								{
									var init___3 = func() interface{} {
										var G__6131 = init___2
										var G__6132 = Aget_(arr, j)
										_, _ = G__6131, G__6132
										return f.(CljsCoreIFn).X_invoke_Arity2(G__6131, G__6132)
									}()
									_ = init___3
									if Reduced_QMARK_.Arity1IB(init___3) {
										return init___3
									} else {
										j, init___2 = (j + float64(1)), init___3
										continue
									}
								}
							} else {
								return init___2
							}
						}
					}()
					_, _, _ = arr, len, init___2
					if Reduced_QMARK_.Arity1IB(init___2) {
						return Deref.X_invoke_Arity1(init___2)
					} else {
						i, init___1 = (i + len), init___2
						continue
					}
				}
			} else {
				return init___1
			}
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIAssociative__() {}

func (coll *CljsCorePersistentVector) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	if Value_(k).Kind() == reflect.Float64 {
		return coll.X_assoc_n_Arity3(k, v)
	} else {
		panic((&js.Error{"Vector's key for assoc must be a number."}))
	}
}

func (_ *CljsCorePersistentVector) CljsCoreISeqable__() {}

func (coll *CljsCorePersistentVector) X_seq_Arity1() interface{} {
	if coll.Cnt.(float64) == float64(0) {
		return nil
	} else {
		if coll.Cnt.(float64) <= float64(32) {
			return (&CljsCoreIndexedSeq{coll.Tail, float64(0)})
		} else {
			return Chunked_seq.X_invoke_Arity4(coll, First_array_for_longvec.X_invoke_Arity1(coll), float64(0), float64(0)).(*CljsCoreChunkedSeq)

		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreISequential__() {}

func (_ *CljsCorePersistentVector) CljsCoreIWithMeta__() {}

func (coll *CljsCorePersistentVector) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCorePersistentVector{meta___1, coll.Cnt, coll.Shift, coll.Root, coll.Tail, coll.X__hash})
}

func (_ *CljsCorePersistentVector) CljsCoreICollection__() {}

func (coll *CljsCorePersistentVector) X_conj_Arity2(o interface{}) interface{} {
	if (coll.Cnt.(float64) - Tail_off.X_invoke_Arity1(coll).(float64)) < float64(32) {
		{
			var len = Alength_(coll.Tail)
			var new_tail = make([]interface{}, int((len + float64(1))))
			_, _ = len, new_tail
			{
				var n__1054__auto___7922 = len
				_ = n__1054__auto___7922
				{
					var i_7923 = float64(0)
					_ = i_7923
					for {
						if i_7923 < n__1054__auto___7922 {
							new_tail[int(i_7923)] = Aget_(coll.Tail, i_7923)
							i_7923 = (i_7923 + float64(1))
							continue
						} else {
						}
						break
					}
				}
			}
			new_tail[int(len)] = o
			return (&CljsCorePersistentVector{coll.Meta, (coll.Cnt.(float64) + float64(1)), coll.Shift, coll.Root, new_tail, nil})
		}
	} else {
		{
			var root_overflow_QMARK_ = (float64((UInt32_(coll.Cnt.(float64)) >> UInt32_(float64((32+Int32_(float64(5)))%32)))) > float64((Int32_(float64(1)) << UInt32_(coll.Shift.(float64)))))
			var new_shift = func() interface{} {
				if Truth_(root_overflow_QMARK_) {
					return (coll.Shift.(float64) + float64(5))
				} else {
					return coll.Shift
				}
			}()
			var new_root = func() *CljsCoreVectorNode {
				if Truth_(root_overflow_QMARK_) {
					return func() *CljsCoreVectorNode {
						var n_r = Pv_fresh_node.X_invoke_Arity1(nil).(*CljsCoreVectorNode)
						_ = n_r
						Pv_aset.X_invoke_Arity3(n_r, float64(0), coll.Root)
						Pv_aset.X_invoke_Arity3(n_r, float64(1), New_path.X_invoke_Arity3(nil, coll.Shift, (&CljsCoreVectorNode{nil, coll.Tail})))
						return n_r
					}()
				} else {
					return Push_tail.X_invoke_Arity4(coll, coll.Shift, coll.Root, (&CljsCoreVectorNode{nil, coll.Tail})).(*CljsCoreVectorNode)
				}
			}()
			_, _, _ = root_overflow_QMARK_, new_shift, new_root
			return (&CljsCorePersistentVector{coll.Meta, (coll.Cnt.(float64) + float64(1)), new_shift, new_root, []interface{}{o}, nil})
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIFn__() {}

func (this *CljsCorePersistentVector) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (coll *CljsCorePersistentVector) X_invoke_Arity1(k interface{}) interface{} {
	return coll.X_nth_Arity2(k)
}

func (coll *CljsCorePersistentVector) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return coll.X_nth_Arity3(k, not_found)
}

func (this *CljsCorePersistentVector) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_PersistentVector *AFn

var CljsCorePersistentVector_EMPTY_NODE = (&CljsCoreVectorNode{nil, make([]interface{}, int(float64(32)))})

var CljsCorePersistentVector_EMPTY = (&CljsCorePersistentVector{nil, float64(0), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{}, float64(0)})

var CljsCorePersistentVector_FromArray = func(G__7924 *AFn) *AFn {
	return Fn(G__7924, 2, func(xs interface{}, no_clone bool) interface{} {
		{
			var l = Alength_(xs)
			var xs___1 = func() interface{} {
				if Truth_(no_clone) {
					return xs
				} else {
					return Aclone.X_invoke_Arity1(xs).([]interface{})
				}
			}()
			_, _ = l, xs___1
			if l < float64(32) {
				return (&CljsCorePersistentVector{nil, l, float64(5), CljsCorePersistentVector_EMPTY_NODE, xs___1, nil})
			} else {
				{
					var node = js.JSArray_(&xs___1).Slice(float64(0), float64(32))
					var v = (&CljsCorePersistentVector{nil, float64(32), float64(5), CljsCorePersistentVector_EMPTY_NODE, node, nil})
					_, _ = node, v
					{
						var i = float64(32)
						var out interface{} = v.X_as_transient_Arity1()
						_, _ = i, out
						for {
							if i < l {
								i, out = (i + float64(1)), Conj_BANG_.X_invoke_Arity2(out, Aget_(xs___1, i))
								continue
							} else {
								return Persistent_BANG_.X_invoke_Arity1(out)
							}
						}
					}
				}
			}
		}
	})
}(&AFn{})

var Vec *AFn

// @param {...*} var_args
var Vector *AFn

type CljsCoreChunkedSeq struct {
	Vec     interface{}
	Node    interface{}
	I       interface{}
	Off     interface{}
	Meta    interface{}
	X__hash interface{}
}

func (_ *CljsCoreChunkedSeq) CljsCoreObject__() {}

func (coll *CljsCoreChunkedSeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreChunkedSeq) String() string {
	return this.ToString()
}

func (this *CljsCoreChunkedSeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreChunkedSeq) CljsCoreIMeta__() {}

func (coll *CljsCoreChunkedSeq) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreChunkedSeq) CljsCoreINext__() {}

func (coll *CljsCoreChunkedSeq) X_next_Arity1() interface{} {
	if (coll.Off.(float64) + float64(1)) < Alength_(coll.Node) {
		{
			var s = Chunked_seq.X_invoke_Arity4(coll.Vec, coll.Node, coll.I, (coll.Off.(float64) + float64(1))).(*CljsCoreChunkedSeq)
			_ = s
			if Nil_(s) {
				return nil
			} else {
				return s
			}
		}
	} else {
		return coll.X_chunked_next_Arity1()
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreIHash__() {}

func (coll *CljsCoreChunkedSeq) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreIEquiv__() {}

func (coll *CljsCoreChunkedSeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreChunkedSeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreChunkedSeq) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCorePersistentVector_EMPTY, coll.Meta)
}

func (_ *CljsCoreChunkedSeq) CljsCoreIReduce__() {}

func (coll *CljsCoreChunkedSeq) X_reduce_Arity2(f interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity2(Subvec.X_invoke_Arity3(coll.Vec, (coll.I.(float64)+coll.Off.(float64)), Count.X_invoke_Arity1(coll.Vec).(float64)).(*CljsCoreSubvec), f)
}

func (coll *CljsCoreChunkedSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity3(Subvec.X_invoke_Arity3(coll.Vec, (coll.I.(float64)+coll.Off.(float64)), Count.X_invoke_Arity1(coll.Vec).(float64)).(*CljsCoreSubvec), f, start)
}

func (_ *CljsCoreChunkedSeq) CljsCoreISeq__() {}

func (coll *CljsCoreChunkedSeq) X_first_Arity1() interface{} {
	return Aget_(coll.Node, coll.Off.(float64))
}

func (coll *CljsCoreChunkedSeq) X_rest_Arity1() interface{} {
	if (coll.Off.(float64) + float64(1)) < Alength_(coll.Node) {
		{
			var s = Chunked_seq.X_invoke_Arity4(coll.Vec, coll.Node, coll.I, (coll.Off.(float64) + float64(1))).(*CljsCoreChunkedSeq)
			_ = s
			if Nil_(s) {
				return CljsCoreIEmptyList(CljsCoreList_EMPTY)
			} else {
				return s
			}
		}
	} else {
		return coll.X_chunked_rest_Arity1()
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreISeqable__() {}

func (coll *CljsCoreChunkedSeq) X_seq_Arity1() interface{} {
	return coll
}

func (_ *CljsCoreChunkedSeq) CljsCoreISequential__() {}

func (_ *CljsCoreChunkedSeq) CljsCoreIChunkedSeq__() {}

func (coll *CljsCoreChunkedSeq) X_chunked_first_Arity1() interface{} {
	return Array_chunk.X_invoke_Arity2(coll.Node, coll.Off).(*CljsCoreArrayChunk)
}

func (coll *CljsCoreChunkedSeq) X_chunked_rest_Arity1() interface{} {
	{
		var end = (coll.I.(float64) + Alength_(coll.Node))
		_ = end
		if end < coll.Vec.(CljsCoreICounted).X_count_Arity1() {
			return Chunked_seq.X_invoke_Arity4(coll.Vec, Unchecked_array_for.X_invoke_Arity2(coll.Vec, end), end, float64(0)).(*CljsCoreChunkedSeq)
		} else {
			return CljsCoreIEmptyList(CljsCoreList_EMPTY)
		}
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreIWithMeta__() {}

func (coll *CljsCoreChunkedSeq) X_with_meta_Arity2(m interface{}) interface{} {
	return Chunked_seq.X_invoke_Arity5(coll.Vec, coll.Node, coll.I, coll.Off, m).(*CljsCoreChunkedSeq)
}

func (_ *CljsCoreChunkedSeq) CljsCoreICollection__() {}

func (coll *CljsCoreChunkedSeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

func (_ *CljsCoreChunkedSeq) CljsCoreASeq__() {}

func (_ *CljsCoreChunkedSeq) CljsCoreIChunkedNext__() {}

func (coll *CljsCoreChunkedSeq) X_chunked_next_Arity1() interface{} {
	{
		var end = (coll.I.(float64) + Alength_(coll.Node))
		_ = end
		if end < coll.Vec.(CljsCoreICounted).X_count_Arity1() {
			return Chunked_seq.X_invoke_Arity4(coll.Vec, Unchecked_array_for.X_invoke_Arity2(coll.Vec, end), end, float64(0)).(*CljsCoreChunkedSeq)
		} else {
			return nil
		}
	}
}

var X__GT_ChunkedSeq *AFn

var Chunked_seq *AFn

type CljsCoreSubvec struct {
	Meta    interface{}
	V       interface{}
	Start   interface{}
	End     interface{}
	X__hash interface{}
}

func (_ *CljsCoreSubvec) CljsCoreObject__() {}

func (coll *CljsCoreSubvec) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreSubvec) String() string {
	return this.ToString()
}

func (this *CljsCoreSubvec) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreSubvec) CljsCoreILookup__() {}

func (coll *CljsCoreSubvec) X_lookup_Arity2(k interface{}) interface{} {
	return coll.X_lookup_Arity3(k, nil)
}

func (coll *CljsCoreSubvec) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	if Value_(k).Kind() == reflect.Float64 {
		return coll.X_nth_Arity3(k, not_found)
	} else {
		return not_found
	}
}

func (_ *CljsCoreSubvec) CljsCoreIIndexed__() {}

func (coll *CljsCoreSubvec) X_nth_Arity2(n interface{}) interface{} {
	if (n.(float64) < float64(0)) || (coll.End.(float64) <= (coll.Start.(float64) + n.(float64))) {
		return Vector_index_out_of_bounds.X_invoke_Arity2(n, (coll.End.(float64) - coll.Start.(float64)))
	} else {
		return coll.V.(CljsCoreIIndexed).X_nth_Arity2((coll.Start.(float64) + n.(float64)))
	}
}

func (coll *CljsCoreSubvec) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	if (n.(float64) < float64(0)) || (coll.End.(float64) <= (coll.Start.(float64) + n.(float64))) {
		return not_found
	} else {
		return coll.V.(CljsCoreIIndexed).X_nth_Arity3((coll.Start.(float64) + n.(float64)), not_found)
	}
}

func (_ *CljsCoreSubvec) CljsCoreIVector__() {}

func (coll *CljsCoreSubvec) X_assoc_n_Arity3(n interface{}, val interface{}) interface{} {
	{
		var v_pos = (coll.Start.(float64) + n.(float64))
		_ = v_pos
		return Build_subvec.X_invoke_Arity5(coll.Meta, Assoc.X_invoke_Arity3(coll.V, v_pos, val), coll.Start, func(x, y float64) float64 {
			if x > y {
				return x
			} else {
				return y
			}
		}(coll.End.(float64), (v_pos+float64(1))), nil).(*CljsCoreSubvec)
	}
}

func (_ *CljsCoreSubvec) CljsCoreIMeta__() {}

func (coll *CljsCoreSubvec) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreSubvec) CljsCoreICloneable__() {}

func (___ *CljsCoreSubvec) X_clone_Arity1() interface{} {
	return (&CljsCoreSubvec{___.Meta, ___.V, ___.Start, ___.End, ___.X__hash})
}

func (_ *CljsCoreSubvec) CljsCoreICounted__() {}

func (coll *CljsCoreSubvec) X_count_Arity1() float64 {
	return (coll.End.(float64) - coll.Start.(float64))
}

func (_ *CljsCoreSubvec) CljsCoreIStack__() {}

func (coll *CljsCoreSubvec) X_peek_Arity1() interface{} {
	return coll.V.(CljsCoreIIndexed).X_nth_Arity2((coll.End.(float64) - float64(1)))
}

func (coll *CljsCoreSubvec) X_pop_Arity1() interface{} {
	if coll.Start.(float64) == coll.End.(float64) {
		panic((&js.Error{"Can't pop empty vector"}))
	} else {
		return Build_subvec.X_invoke_Arity5(coll.Meta, coll.V, coll.Start, (coll.End.(float64) - float64(1)), nil).(*CljsCoreSubvec)
	}
}

func (_ *CljsCoreSubvec) CljsCoreIReversible__() {}

func (coll *CljsCoreSubvec) X_rseq_Arity1() interface{} {
	if !(coll.Start.(float64) == coll.End.(float64)) {
		return (&CljsCoreRSeq{coll, ((coll.End.(float64) - coll.Start.(float64)) - float64(1)), nil})
	} else {
		return nil
	}
}

func (_ *CljsCoreSubvec) CljsCoreIHash__() {}

func (coll *CljsCoreSubvec) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreSubvec) CljsCoreIEquiv__() {}

func (coll *CljsCoreSubvec) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreSubvec) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreSubvec) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCorePersistentVector_EMPTY, coll.Meta)
}

func (_ *CljsCoreSubvec) CljsCoreIReduce__() {}

func (coll *CljsCoreSubvec) X_reduce_Arity2(f interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity2(coll, f)
}

func (coll *CljsCoreSubvec) X_reduce_Arity3(f interface{}, start___1 interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity3(coll, f, start___1)
}

func (_ *CljsCoreSubvec) CljsCoreIAssociative__() {}

func (coll *CljsCoreSubvec) X_assoc_Arity3(key interface{}, val interface{}) interface{} {
	if Value_(key).Kind() == reflect.Float64 {
		return coll.X_assoc_n_Arity3(key, val)
	} else {
		panic((&js.Error{"Subvec's key for assoc must be a number."}))
	}
}

func (_ *CljsCoreSubvec) CljsCoreISeqable__() {}

func (coll *CljsCoreSubvec) X_seq_Arity1() interface{} {
	{
		var subvec_seq = func(subvec_seq *AFn) *AFn {
			return Fn(subvec_seq, 1, func(i interface{}) interface{} {
				if i.(float64) == coll.End.(float64) {
					return nil
				} else {
					return Cons.X_invoke_Arity2(coll.V.(CljsCoreIIndexed).X_nth_Arity2(i), (&CljsCoreLazySeq{nil, func(G__7925 *AFn) *AFn {
						return Fn(G__7925, 0, func() interface{} {
							return subvec_seq.X_invoke_Arity1((i.(float64) + float64(1)))
						})
					}(&AFn{}), nil, nil})).(*CljsCoreCons)
				}
			})
		}(&AFn{})
		_ = subvec_seq
		return subvec_seq.X_invoke_Arity1(coll.Start)
	}
}

func (_ *CljsCoreSubvec) CljsCoreISequential__() {}

func (_ *CljsCoreSubvec) CljsCoreIWithMeta__() {}

func (coll *CljsCoreSubvec) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return Build_subvec.X_invoke_Arity5(meta___1, coll.V, coll.Start, coll.End, coll.X__hash).(*CljsCoreSubvec)
}

func (_ *CljsCoreSubvec) CljsCoreICollection__() {}

func (coll *CljsCoreSubvec) X_conj_Arity2(o interface{}) interface{} {
	return Build_subvec.X_invoke_Arity5(coll.Meta, coll.V.(CljsCoreIVector).X_assoc_n_Arity3(coll.End, o), coll.Start, (coll.End.(float64) + float64(1)), nil).(*CljsCoreSubvec)
}

func (_ *CljsCoreSubvec) CljsCoreIFn__() {}

func (this *CljsCoreSubvec) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (coll *CljsCoreSubvec) X_invoke_Arity1(k interface{}) interface{} {
	return coll.X_nth_Arity2(k)
}

func (coll *CljsCoreSubvec) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return coll.X_nth_Arity3(k, not_found)
}

func (this *CljsCoreSubvec) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_Subvec *AFn

var Build_subvec *AFn

// Returns a persistent vector of the items in vector from
// start (inclusive) to end (exclusive).  If end is not supplied,
// defaults to (count vector). This operation is O(1) and very fast, as
// the resulting vector shares structure with the original and no
// trimming is done.
var Subvec *AFn

var Tv_ensure_editable *AFn

var Tv_editable_root *AFn

var Tv_editable_tail *AFn

var Tv_push_tail *AFn

var Tv_pop_tail *AFn

var Unchecked_editable_array_for *AFn

type CljsCoreTransientVector struct {
	Cnt   interface{}
	Shift interface{}
	Root  interface{}
	Tail  interface{}
}

func (_ *CljsCoreTransientVector) CljsCoreIFn__() {}

func (this *CljsCoreTransientVector) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (coll *CljsCoreTransientVector) X_invoke_Arity1(k interface{}) interface{} {
	return coll.X_lookup_Arity2(k)
}

func (coll *CljsCoreTransientVector) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return coll.X_lookup_Arity3(k, not_found)
}

func (this *CljsCoreTransientVector) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

func (_ *CljsCoreTransientVector) CljsCoreILookup__() {}

func (coll *CljsCoreTransientVector) X_lookup_Arity2(k interface{}) interface{} {
	return coll.X_lookup_Arity3(k, nil)
}

func (coll *CljsCoreTransientVector) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	if Value_(k).Kind() == reflect.Float64 {
		return coll.X_nth_Arity3(k, not_found)
	} else {
		return not_found
	}
}

func (_ *CljsCoreTransientVector) CljsCoreIIndexed__() {}

func (coll *CljsCoreTransientVector) X_nth_Arity2(n interface{}) interface{} {
	if Truth_(Native_get_instance_field.X_invoke_Arity2(coll.Root, "Edit")) {
		return Aget_(Array_for.X_invoke_Arity2(coll, n), float64((Int32_(n.(float64)) & Int32_(float64(31)))))
	} else {
		panic((&js.Error{"nth after persistent!"}))
	}
}

func (coll *CljsCoreTransientVector) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	if (float64(0) <= n.(float64)) && (n.(float64) < coll.Cnt.(float64)) {
		return coll.X_nth_Arity2(n)
	} else {
		return not_found
	}
}

func (_ *CljsCoreTransientVector) CljsCoreICounted__() {}

func (coll *CljsCoreTransientVector) X_count_Arity1() float64 {
	if Truth_(Native_get_instance_field.X_invoke_Arity2(coll.Root, "Edit")) {
		return coll.Cnt.(float64)
	} else {
		panic((&js.Error{"count after persistent!"}))
	}
}

func (_ *CljsCoreTransientVector) CljsCoreITransientVector__() {}

func (tcoll *CljsCoreTransientVector) X_assoc_n_BANG__Arity3(n interface{}, val interface{}) interface{} {
	if Truth_(Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit")) {
		if (float64(0) <= n.(float64)) && (n.(float64) < tcoll.Cnt.(float64)) {
			if Tail_off.X_invoke_Arity1(tcoll).(float64) <= n.(float64) {
				tcoll.Tail.([]interface{})[int(float64((Int32_(n.(float64)) & Int32_(float64(31)))))] = val
				return tcoll
			} else {
				{
					var new_root = func(go_ *AFn) *AFn {
						return Fn(go_, 2, func(level interface{}, node interface{}) interface{} {
							{
								var node___1 = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit"), node)
								_ = node___1
								if level.(float64) == float64(0) {
									Pv_aset.X_invoke_Arity3(node___1, float64((Int32_(n.(float64)) & Int32_(float64(31)))), val)
									return node___1
								} else {
									{
										var subidx = float64((Int32_(float64((UInt32_(n.(float64)) >> UInt32_(float64((32+Int32_(level.(float64)))%32))))) & Int32_(float64(31))))
										_ = subidx
										Pv_aset.X_invoke_Arity3(node___1, subidx, go_.X_invoke_Arity2((level.(float64)-float64(5)), Pv_aget.X_invoke_Arity2(node___1, subidx)))
										return node___1
									}
								}
							}
						})
					}(&AFn{}).X_invoke_Arity2(tcoll.Shift, tcoll.Root)
					_ = new_root
					tcoll.Root = new_root

					return tcoll
				}
			}
		} else {
			if n.(float64) == tcoll.Cnt.(float64) {
				return tcoll.X_conj_BANG__Arity2(val)
			} else {
				panic((&js.Error{("Index " + Str.X_invoke_Arity1(n).(string) + " out of bounds for TransientVector of length" + Str.X_invoke_Arity1(tcoll.Cnt).(string))}))

			}
		}
	} else {
		panic((&js.Error{"assoc! after persistent!"}))
	}
}

func (tcoll *CljsCoreTransientVector) X_pop_BANG__Arity1() interface{} {
	if Truth_(Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit")) {
		if tcoll.Cnt.(float64) == float64(0) {
			panic((&js.Error{"Can't pop empty vector"}))
		} else {
			if float64(1) == tcoll.Cnt.(float64) {
				tcoll.Cnt = float64(0)

				return tcoll
			} else {
				if float64((Int32_((tcoll.Cnt.(float64) - float64(1))) & Int32_(float64(31)))) > float64(0) {
					tcoll.Cnt = (tcoll.Cnt.(float64) - float64(1))

					return tcoll
				} else {
					{
						var new_tail = Unchecked_editable_array_for.X_invoke_Arity2(tcoll, (tcoll.Cnt.(float64) - float64(2)))
						var new_root = func() interface{} {
							var nr = Tv_pop_tail.X_invoke_Arity3(tcoll, tcoll.Shift, tcoll.Root)
							_ = nr
							if !(Nil_(nr)) {
								return nr
							} else {
								return (&CljsCoreVectorNode{Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit"), make([]interface{}, int(float64(32)))})
							}
						}()
						_, _ = new_tail, new_root
						if (float64(5) < tcoll.Shift.(float64)) && (Nil_(Pv_aget.X_invoke_Arity2(new_root, float64(1)))) {
							{
								var new_root___1 = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit"), Pv_aget.X_invoke_Arity2(new_root, float64(0)))
								_ = new_root___1
								tcoll.Root = new_root___1

								tcoll.Shift = (tcoll.Shift.(float64) - float64(5))

								tcoll.Cnt = (tcoll.Cnt.(float64) - float64(1))

								tcoll.Tail = new_tail

								return tcoll
							}
						} else {
							tcoll.Root = new_root

							tcoll.Cnt = (tcoll.Cnt.(float64) - float64(1))

							tcoll.Tail = new_tail

							return tcoll
						}
					}

				}
			}
		}
	} else {
		panic((&js.Error{"pop! after persistent!"}))
	}
}

func (_ *CljsCoreTransientVector) CljsCoreITransientAssociative__() {}

func (tcoll *CljsCoreTransientVector) X_assoc_BANG__Arity3(key interface{}, val interface{}) interface{} {
	if Value_(key).Kind() == reflect.Float64 {
		return tcoll.X_assoc_n_BANG__Arity3(key, val)
	} else {
		panic((&js.Error{"TransientVector's key for assoc! must be a number."}))
	}
}

func (_ *CljsCoreTransientVector) CljsCoreITransientCollection__() {}

func (tcoll *CljsCoreTransientVector) X_conj_BANG__Arity2(o interface{}) interface{} {
	if Truth_(Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit")) {
		if (tcoll.Cnt.(float64) - Tail_off.X_invoke_Arity1(tcoll).(float64)) < float64(32) {
			tcoll.Tail.([]interface{})[int(float64((Int32_(tcoll.Cnt.(float64)) & Int32_(float64(31)))))] = o
			tcoll.Cnt = (tcoll.Cnt.(float64) + float64(1))

			return tcoll
		} else {
			{
				var tail_node = (&CljsCoreVectorNode{Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit"), tcoll.Tail})
				var new_tail = make([]interface{}, int(float64(32)))
				_, _ = tail_node, new_tail
				new_tail[int(float64(0))] = o
				tcoll.Tail = new_tail

				if float64((UInt32_(tcoll.Cnt.(float64)) >> UInt32_(float64((32+Int32_(float64(5)))%32)))) > float64((Int32_(float64(1)) << UInt32_(tcoll.Shift.(float64)))) {
					{
						var new_root_array = make([]interface{}, int(float64(32)))
						var new_shift = (tcoll.Shift.(float64) + float64(5))
						_, _ = new_root_array, new_shift
						new_root_array[int(float64(0))] = tcoll.Root
						new_root_array[int(float64(1))] = New_path.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit"), tcoll.Shift, tail_node)
						tcoll.Root = (&CljsCoreVectorNode{Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit"), new_root_array})

						tcoll.Shift = new_shift

						tcoll.Cnt = (tcoll.Cnt.(float64) + float64(1))

						return tcoll
					}
				} else {
					{
						var new_root = Tv_push_tail.X_invoke_Arity4(tcoll, tcoll.Shift, tcoll.Root, tail_node)
						_ = new_root
						tcoll.Root = new_root

						tcoll.Cnt = (tcoll.Cnt.(float64) + float64(1))

						return tcoll
					}
				}
			}
		}
	} else {
		panic((&js.Error{"conj! after persistent!"}))
	}
}

func (tcoll *CljsCoreTransientVector) X_persistent_BANG__Arity1() interface{} {
	if Truth_(Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit")) {
		Native_set_instance_field.X_invoke_Arity3(tcoll.Root, "Edit", nil)
		{
			var len = (tcoll.Cnt.(float64) - Tail_off.X_invoke_Arity1(tcoll).(float64))
			var trimmed_tail = make([]interface{}, int(len))
			_, _ = len, trimmed_tail
			Array_copy.X_invoke_Arity5(tcoll.Tail, float64(0), trimmed_tail, float64(0), len)
			return (&CljsCorePersistentVector{nil, tcoll.Cnt, tcoll.Shift, tcoll.Root, trimmed_tail, nil})
		}
	} else {
		panic((&js.Error{"persistent! called twice"}))
	}
}

var X__GT_TransientVector *AFn

type CljsCorePersistentQueueSeq struct {
	Meta    interface{}
	Front   interface{}
	Rear    interface{}
	X__hash interface{}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreObject__() {}

func (coll *CljsCorePersistentQueueSeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCorePersistentQueueSeq) String() string {
	return this.ToString()
}

func (this *CljsCorePersistentQueueSeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIMeta__() {}

func (coll *CljsCorePersistentQueueSeq) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIHash__() {}

func (coll *CljsCorePersistentQueueSeq) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentQueueSeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCorePersistentQueueSeq) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.Meta)
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreISeq__() {}

func (coll *CljsCorePersistentQueueSeq) X_first_Arity1() interface{} {
	return First.X_invoke_Arity1(coll.Front)
}

func (coll *CljsCorePersistentQueueSeq) X_rest_Arity1() interface{} {
	{
		var temp__4220__auto__ = Next.Arity1IQ(coll.Front)
		_ = temp__4220__auto__
		if Truth_(temp__4220__auto__) {
			{
				var f1 = temp__4220__auto__
				_ = f1
				return (&CljsCorePersistentQueueSeq{coll.Meta, f1, coll.Rear, nil})
			}
		} else {
			if Nil_(coll.Rear) {
				return coll.X_empty_Arity1()
			} else {
				return (&CljsCorePersistentQueueSeq{coll.Meta, coll.Rear, nil, nil})
			}
		}
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreISeqable__() {}

func (coll *CljsCorePersistentQueueSeq) X_seq_Arity1() interface{} {
	return coll
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreISequential__() {}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIWithMeta__() {}

func (coll *CljsCorePersistentQueueSeq) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCorePersistentQueueSeq{meta___1, coll.Front, coll.Rear, coll.X__hash})
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreICollection__() {}

func (coll *CljsCorePersistentQueueSeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

var X__GT_PersistentQueueSeq *AFn

type CljsCorePersistentQueue struct {
	Meta    interface{}
	Count   interface{}
	Front   interface{}
	Rear    interface{}
	X__hash interface{}
}

func (_ *CljsCorePersistentQueue) CljsCoreObject__() {}

func (coll *CljsCorePersistentQueue) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCorePersistentQueue) String() string {
	return this.ToString()
}

func (this *CljsCorePersistentQueue) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCorePersistentQueue) CljsCoreIMeta__() {}

func (coll *CljsCorePersistentQueue) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCorePersistentQueue) CljsCoreICloneable__() {}

func (coll *CljsCorePersistentQueue) X_clone_Arity1() interface{} {
	return (&CljsCorePersistentQueue{coll.Meta, coll.Count, coll.Front, coll.Rear, coll.X__hash})
}

func (_ *CljsCorePersistentQueue) CljsCoreICounted__() {}

func (coll *CljsCorePersistentQueue) X_count_Arity1() float64 {
	return coll.Count.(float64)
}

func (_ *CljsCorePersistentQueue) CljsCoreIStack__() {}

func (coll *CljsCorePersistentQueue) X_peek_Arity1() interface{} {
	return First.X_invoke_Arity1(coll.Front)
}

func (coll *CljsCorePersistentQueue) X_pop_Arity1() interface{} {
	if Truth_(coll.Front) {
		{
			var temp__4220__auto__ = Next.Arity1IQ(coll.Front)
			_ = temp__4220__auto__
			if Truth_(temp__4220__auto__) {
				{
					var f1 = temp__4220__auto__
					_ = f1
					return (&CljsCorePersistentQueue{coll.Meta, (coll.Count.(float64) - float64(1)), f1, coll.Rear, nil})
				}
			} else {
				return (&CljsCorePersistentQueue{coll.Meta, (coll.Count.(float64) - float64(1)), Seq.Arity1IQ(coll.Rear), CljsCorePersistentVector_EMPTY, nil})
			}
		}
	} else {
		return coll
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreIHash__() {}

func (coll *CljsCorePersistentQueue) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentQueue) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCorePersistentQueue) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCorePersistentQueue) X_empty_Arity1() interface{} {
	return CljsCorePersistentQueue_EMPTY
}

func (_ *CljsCorePersistentQueue) CljsCoreISeq__() {}

func (coll *CljsCorePersistentQueue) X_first_Arity1() interface{} {
	return First.X_invoke_Arity1(coll.Front)
}

func (coll *CljsCorePersistentQueue) X_rest_Arity1() interface{} {
	return Rest.Arity1IQ(Seq.Arity1IQ(coll))
}

func (_ *CljsCorePersistentQueue) CljsCoreISeqable__() {}

func (coll *CljsCorePersistentQueue) X_seq_Arity1() interface{} {
	{
		var rear___1 = Seq.Arity1IQ(coll.Rear)
		_ = rear___1
		if Truth_(func() interface{} {
			var or__171__auto__ = coll.Front
			_ = or__171__auto__
			if Truth_(or__171__auto__) {
				return or__171__auto__
			} else {
				return rear___1
			}
		}()) {
			return (&CljsCorePersistentQueueSeq{nil, coll.Front, Seq.Arity1IQ(rear___1), nil})
		} else {
			return nil
		}
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreISequential__() {}

func (_ *CljsCorePersistentQueue) CljsCoreIWithMeta__() {}

func (coll *CljsCorePersistentQueue) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCorePersistentQueue{meta___1, coll.Count, coll.Front, coll.Rear, coll.X__hash})
}

func (_ *CljsCorePersistentQueue) CljsCoreICollection__() {}

func (coll *CljsCorePersistentQueue) X_conj_Arity2(o interface{}) interface{} {
	if Truth_(coll.Front) {
		return (&CljsCorePersistentQueue{coll.Meta, (coll.Count.(float64) + float64(1)), coll.Front, Conj.X_invoke_Arity2(func() interface{} {
			var or__171__auto__ = coll.Rear
			_ = or__171__auto__
			if Truth_(or__171__auto__) {
				return or__171__auto__
			} else {
				return CljsCorePersistentVector_EMPTY
			}
		}(), o), nil})
	} else {
		return (&CljsCorePersistentQueue{coll.Meta, (coll.Count.(float64) + float64(1)), Conj.X_invoke_Arity2(coll.Front, o), CljsCorePersistentVector_EMPTY, nil})
	}
}

var X__GT_PersistentQueue *AFn

var CljsCorePersistentQueue_EMPTY = (&CljsCorePersistentQueue{nil, float64(0), nil, CljsCorePersistentVector_EMPTY, float64(0)})

type CljsCoreNeverEquiv struct{}

func (_ *CljsCoreNeverEquiv) CljsCoreIEquiv__() {}

func (o *CljsCoreNeverEquiv) X_equiv_Arity2(other interface{}) bool {
	return false
}

func (_ *CljsCoreNeverEquiv) CljsCoreObject__() {}

func (this *CljsCoreNeverEquiv) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

var X__GT_NeverEquiv *AFn

var Never_equiv *CljsCoreNeverEquiv

// Assumes y is a map. Returns true if x equals y, otherwise returns
// false.
var Equiv_map *AFn

var Scan_array *AFn

var Obj_map_compare_keys *AFn

var Obj_map__GT_hash_map *AFn

type CljsCoreIterator struct{ S interface{} }

func (_ *CljsCoreIterator) CljsCoreObject__() {}

func (___ *CljsCoreIterator) Next() interface{} {
	if !(Nil_(___.S)) {
		{
			var x = First.X_invoke_Arity1(___.S)
			_ = x
			___.S = Native_invoke_func.X_invoke_Arity2((*CljsCoreIterator).Next, []interface{}{___.S})

			return map[string]interface{}{"done": false, "value": x}
		}
	} else {
		return map[string]interface{}{"done": true, "value": nil}
	}
}

var X__GT_Iterator *AFn

var Iterator *AFn

type CljsCoreEntriesIterator struct{ S interface{} }

func (_ *CljsCoreEntriesIterator) CljsCoreObject__() {}

func (___ *CljsCoreEntriesIterator) Next() interface{} {
	if !(Nil_(___.S)) {
		{
			var vec__6178 = First.X_invoke_Arity1(___.S)
			var k = Nth.X_invoke_Arity3(vec__6178, float64(0), nil)
			var v = Nth.X_invoke_Arity3(vec__6178, float64(1), nil)
			_, _, _ = vec__6178, k, v
			___.S = Native_invoke_func.X_invoke_Arity2((*CljsCoreEntriesIterator).Next, []interface{}{___.S})

			return map[string]interface{}{"done": false, "value": []interface{}{k, v}}
		}
	} else {
		return map[string]interface{}{"done": true, "value": nil}
	}
}

var X__GT_EntriesIterator *AFn

var Entries_iterator *AFn

type CljsCoreSetEntriesIterator struct{ S interface{} }

func (_ *CljsCoreSetEntriesIterator) CljsCoreObject__() {}

func (___ *CljsCoreSetEntriesIterator) Next() interface{} {
	if !(Nil_(___.S)) {
		{
			var x = First.X_invoke_Arity1(___.S)
			_ = x
			___.S = Native_invoke_func.X_invoke_Arity2((*CljsCoreSetEntriesIterator).Next, []interface{}{___.S})

			return map[string]interface{}{"done": false, "value": []interface{}{x, x}}
		}
	} else {
		return map[string]interface{}{"done": true, "value": nil}
	}
}

var X__GT_SetEntriesIterator *AFn

var Set_entries_iterator *AFn

var Array_map_index_of_nil_QMARK_ *AFn

var Array_map_index_of_keyword_QMARK_ *AFn

var Array_map_index_of_symbol_QMARK_ *AFn

var Array_map_index_of_identical_QMARK_ *AFn

var Array_map_index_of_equiv_QMARK_ *AFn

var Array_map_index_of *AFn

var Array_map_extend_kv *AFn

type CljsCorePersistentArrayMapSeq struct {
	Arr    interface{}
	I      interface{}
	X_meta interface{}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreObject__() {}

func (coll *CljsCorePersistentArrayMapSeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCorePersistentArrayMapSeq) String() string {
	return this.ToString()
}

func (this *CljsCorePersistentArrayMapSeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIMeta__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_meta_Arity1() interface{} {
	return coll.X_meta
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreINext__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_next_Arity1() interface{} {
	if coll.I.(float64) < (Alength_(coll.Arr) - float64(2)) {
		return (&CljsCorePersistentArrayMapSeq{coll.Arr, (coll.I.(float64) + float64(2)), coll.X_meta})
	} else {
		return nil
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreICounted__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_count_Arity1() float64 {
	return ((Alength_(coll.Arr) - coll.I.(float64)) / float64(2))
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIHash__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_hash_Arity1() interface{} {
	return Hash_ordered_coll.Arity1IF(coll)
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.X_meta)
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIReduce__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCorePersistentArrayMapSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreISeq__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_first_Arity1() interface{} {
	return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Aget_(coll.Arr, coll.I.(float64)), Aget_(coll.Arr, (coll.I.(float64) + float64(1)))}, nil})
}

func (coll *CljsCorePersistentArrayMapSeq) X_rest_Arity1() interface{} {
	if coll.I.(float64) < (Alength_(coll.Arr) - float64(2)) {
		return (&CljsCorePersistentArrayMapSeq{coll.Arr, (coll.I.(float64) + float64(2)), coll.X_meta})
	} else {
		return CljsCoreIEmptyList(CljsCoreList_EMPTY)
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreISeqable__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_seq_Arity1() interface{} {
	return coll
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreISequential__() {}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIWithMeta__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_with_meta_Arity2(new_meta interface{}) interface{} {
	return (&CljsCorePersistentArrayMapSeq{coll.Arr, coll.I, new_meta})
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreICollection__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

var X__GT_PersistentArrayMapSeq *AFn

var Persistent_array_map_seq *AFn

type CljsCorePersistentArrayMap struct {
	Meta    interface{}
	Cnt     interface{}
	Arr     interface{}
	X__hash interface{}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreObject__() {}

func (coll *CljsCorePersistentArrayMap) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCorePersistentArrayMap) String() string {
	return this.ToString()
}

func (this *CljsCorePersistentArrayMap) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (coll *CljsCorePersistentArrayMap) Keys() interface{} {
	return Iterator.X_invoke_Arity1(Native_invoke_func.X_invoke_Arity2((*CljsCorePersistentArrayMap).Keys, []interface{}{coll}).(*CljsCoreIterator)).(*CljsCoreIterator)
}

func (coll *CljsCorePersistentArrayMap) Entries() interface{} {
	return Entries_iterator.X_invoke_Arity1(Seq.Arity1IQ(coll)).(*CljsCoreEntriesIterator)
}

func (coll *CljsCorePersistentArrayMap) Values() interface{} {
	return Iterator.X_invoke_Arity1(Vals.X_invoke_Arity1(coll)).(*CljsCoreIterator)
}

func (coll *CljsCorePersistentArrayMap) Has(k interface{}) interface{} {
	return Contains_QMARK_.Arity2IIB(coll, k)
}

func (coll *CljsCorePersistentArrayMap) Get(k interface{}) interface{} {
	return coll.X_lookup_Arity2(k)
}

func (coll *CljsCorePersistentArrayMap) ForEach(f interface{}) interface{} {
	{
		var seq__6196 interface{} = Seq.Arity1IQ(coll)
		var chunk__6197 interface{} = nil
		var count__6198 = float64(0)
		var i__6199 = float64(0)
		_, _, _, _ = seq__6196, chunk__6197, count__6198, i__6199
		for {
			if i__6199 < count__6198 {
				{
					var vec__6200 = chunk__6197.(CljsCoreIIndexed).X_nth_Arity2(i__6199)
					var k = Nth.X_invoke_Arity3(vec__6200, float64(0), nil)
					var v = Nth.X_invoke_Arity3(vec__6200, float64(1), nil)
					_, _, _ = vec__6200, k, v
					{
						var G__6201_7926 = v
						var G__6202_7927 = k
						_, _ = G__6201_7926, G__6202_7927
						f.(CljsCoreIFn).X_invoke_Arity2(G__6201_7926, G__6202_7927)
					}
					seq__6196, chunk__6197, count__6198, i__6199 = seq__6196, chunk__6197, count__6198, (i__6199 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__6196)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__6196___1 = temp__4222__auto__
							_ = seq__6196___1
							if Chunked_seq_QMARK_.Arity1IB(seq__6196___1) {
								{
									var c__954__auto__ = Chunk_first.X_invoke_Arity1(seq__6196___1)
									_ = c__954__auto__
									seq__6196, chunk__6197, count__6198, i__6199 = Chunk_rest.X_invoke_Arity1(seq__6196___1), c__954__auto__, Count.X_invoke_Arity1(c__954__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__6203 = First.X_invoke_Arity1(seq__6196___1)
									var k = Nth.X_invoke_Arity3(vec__6203, float64(0), nil)
									var v = Nth.X_invoke_Arity3(vec__6203, float64(1), nil)
									_, _, _ = vec__6203, k, v
									{
										var G__6204_7928 = v
										var G__6205_7929 = k
										_, _ = G__6204_7928, G__6205_7929
										f.(CljsCoreIFn).X_invoke_Arity2(G__6204_7928, G__6205_7929)
									}
									seq__6196, chunk__6197, count__6198, i__6199 = Next.Arity1IQ(seq__6196___1), nil, float64(0), float64(0)
									continue
								}
							}
						}
					} else {
						return nil
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreILookup__() {}

func (coll *CljsCorePersistentArrayMap) X_lookup_Arity2(k interface{}) interface{} {
	return coll.X_lookup_Arity3(k, nil)
}

func (coll *CljsCorePersistentArrayMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var idx = Array_map_index_of.X_invoke_Arity2(coll, k).(float64)
		_ = idx
		if idx == float64(-1) {
			return not_found
		} else {
			return Aget_(coll.Arr, (idx + float64(1)))
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIKVReduce__() {}

func (coll *CljsCorePersistentArrayMap) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var len = Alength_(coll.Arr)
		_ = len
		{
			var i = float64(0)
			var init___1 interface{} = init
			_, _ = i, init___1
			for {
				if i < len {
					{
						var init___2 = func() interface{} {
							var G__6212 = init___1
							var G__6213 = Aget_(coll.Arr, i)
							var G__6214 = Aget_(coll.Arr, (i + float64(1)))
							_, _, _ = G__6212, G__6213, G__6214
							return f.(CljsCoreIFn).X_invoke_Arity3(G__6212, G__6213, G__6214)
						}()
						_ = init___2
						if Reduced_QMARK_.Arity1IB(init___2) {
							return Deref.X_invoke_Arity1(init___2)
						} else {
							i, init___1 = (i + float64(2)), init___2
							continue
						}
					}
				} else {
					return init___1
				}
			}
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIMeta__() {}

func (coll *CljsCorePersistentArrayMap) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCorePersistentArrayMap) CljsCoreICloneable__() {}

func (___ *CljsCorePersistentArrayMap) X_clone_Arity1() interface{} {
	return (&CljsCorePersistentArrayMap{___.Meta, ___.Cnt, ___.Arr, ___.X__hash})
}

func (_ *CljsCorePersistentArrayMap) CljsCoreICounted__() {}

func (coll *CljsCorePersistentArrayMap) X_count_Arity1() float64 {
	return coll.Cnt.(float64)
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIHash__() {}

func (coll *CljsCorePersistentArrayMap) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentArrayMap) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_map.X_invoke_Arity2(coll, other))
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIEditableCollection__() {}

func (coll *CljsCorePersistentArrayMap) X_as_transient_Arity1() interface{} {
	return (&CljsCoreTransientArrayMap{true, Alength_(coll.Arr), Aclone.X_invoke_Arity1(coll.Arr).([]interface{})})
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCorePersistentArrayMap) X_empty_Arity1() interface{} {
	return CljsCorePersistentArrayMap_EMPTY.X_with_meta_Arity2(coll.Meta)
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIReduce__() {}

func (coll *CljsCorePersistentArrayMap) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCorePersistentArrayMap) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIMap__() {}

func (coll *CljsCorePersistentArrayMap) X_dissoc_Arity2(k interface{}) interface{} {
	{
		var idx = Array_map_index_of.X_invoke_Arity2(coll, k).(float64)
		_ = idx
		if idx >= float64(0) {
			{
				var len = Alength_(coll.Arr)
				var new_len = (len - float64(2))
				_, _ = len, new_len
				if new_len == float64(0) {
					return coll.X_empty_Arity1()
				} else {
					{
						var new_arr = make([]interface{}, int(new_len))
						_ = new_arr
						{
							var s = float64(0)
							var d = float64(0)
							_, _ = s, d
							for {
								if s >= len {
									return (&CljsCorePersistentArrayMap{coll.Meta, (coll.Cnt.(float64) - float64(1)), new_arr, nil})
								} else {
									if X_EQ_.Arity2IIB(k, Aget_(coll.Arr, s)) {
										s, d = (s + float64(2)), d
										continue
									} else {
										new_arr[int(d)] = Aget_(coll.Arr, s)
										new_arr[int((d + float64(1)))] = Aget_(coll.Arr, (s + float64(1)))
										s, d = (s + float64(2)), (d + float64(2))
										continue

									}
								}
							}
						}
					}
				}
			}
		} else {
			return coll
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIAssociative__() {}

func (coll *CljsCorePersistentArrayMap) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	{
		var idx = Array_map_index_of.X_invoke_Arity2(coll, k).(float64)
		_ = idx
		if idx == float64(-1) {
			if coll.Cnt.(float64) < CljsCorePersistentArrayMap_HASHMAP_THRESHOLD {
				{
					var arr___1 = Array_map_extend_kv.X_invoke_Arity3(coll, k, v).([]interface{})
					_ = arr___1
					return (&CljsCorePersistentArrayMap{coll.Meta, (coll.Cnt.(float64) + float64(1)), arr___1, nil})
				}
			} else {
				return Into.X_invoke_Arity2(CljsCorePersistentHashMap_EMPTY, coll).(CljsCoreIAssociative).X_assoc_Arity3(k, v).(CljsCoreIWithMeta).X_with_meta_Arity2(coll.Meta)
			}
		} else {
			if reflect.DeepEqual(v, Aget_(coll.Arr, (idx+float64(1)))) {
				return coll
			} else {
				{
					var arr___1 = func() []interface{} {
						var G__6223 = Aclone.X_invoke_Arity1(coll.Arr).([]interface{})
						_ = G__6223
						G__6223[int((idx + float64(1)))] = v
						return G__6223
					}()
					_ = arr___1
					return (&CljsCorePersistentArrayMap{coll.Meta, coll.Cnt, arr___1, nil})
				}

			}
		}
	}
}

func (coll *CljsCorePersistentArrayMap) X_contains_key_QMARK__Arity2(k interface{}) bool {
	return !(Array_map_index_of.X_invoke_Arity2(coll, k).(float64) == float64(-1))
}

func (_ *CljsCorePersistentArrayMap) CljsCoreISeqable__() {}

func (coll *CljsCorePersistentArrayMap) X_seq_Arity1() interface{} {
	return Persistent_array_map_seq.X_invoke_Arity3(coll.Arr, float64(0), nil)
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIWithMeta__() {}

func (coll *CljsCorePersistentArrayMap) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCorePersistentArrayMap{meta___1, coll.Cnt, coll.Arr, coll.X__hash})
}

func (_ *CljsCorePersistentArrayMap) CljsCoreICollection__() {}

func (coll *CljsCorePersistentArrayMap) X_conj_Arity2(entry interface{}) interface{} {
	if Vector_QMARK_.Arity1IB(entry) {
		return coll.X_assoc_Arity3(entry.(CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry.(CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		{
			var ret interface{} = coll
			var es interface{} = Seq.Arity1IQ(entry)
			_, _ = ret, es
			for {
				if Nil_(es) {
					return ret
				} else {
					{
						var e = First.X_invoke_Arity1(es)
						_ = e
						if Vector_QMARK_.Arity1IB(e) {
							ret, es = ret.(CljsCoreIAssociative).X_assoc_Arity3(e.(CljsCoreIIndexed).X_nth_Arity2(float64(0)), e.(CljsCoreIIndexed).X_nth_Arity2(float64(1))), Next.Arity1IQ(es)
							continue
						} else {
							panic((&js.Error{"conj on a map takes map entries or seqables of map entries"}))
						}
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIFn__() {}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (coll *CljsCorePersistentArrayMap) X_invoke_Arity1(k interface{}) interface{} {
	return coll.X_lookup_Arity2(k)
}

func (coll *CljsCorePersistentArrayMap) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return coll.X_lookup_Arity3(k, not_found)
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_PersistentArrayMap *AFn

var CljsCorePersistentArrayMap_EMPTY = (&CljsCorePersistentArrayMap{nil, float64(0), []interface{}{}, nil})

var CljsCorePersistentArrayMap_HASHMAP_THRESHOLD = float64(8)

var CljsCorePersistentArrayMap_FromArray = func(G__7930 *AFn) *AFn {
	return Fn(G__7930, 3, func(arr interface{}, no_clone bool, no_check bool) interface{} {
		{
			var arr___1 = func() interface{} {
				if Truth_(no_clone) {
					return arr
				} else {
					return Aclone.X_invoke_Arity1(arr).([]interface{})
				}
			}()
			_ = arr___1
			if Truth_(no_check) {
				{
					var cnt = (Alength_(arr___1) / float64(2))
					_ = cnt
					return (&CljsCorePersistentArrayMap{nil, cnt, arr___1, nil})
				}
			} else {
				{
					var len = Alength_(arr___1)
					_ = len
					{
						var i = float64(0)
						var ret interface{} = Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY)
						_, _ = i, ret
						for {
							if i < len {
								i, ret = (i + float64(2)), ret.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3(Aget_(arr___1, i), Aget_(arr___1, (i+float64(1))))
								continue
							} else {
								return ret.(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
							}
						}
					}
				}
			}
		}
	})
}(&AFn{})

type CljsCoreTransientArrayMap struct {
	Editable_QMARK_ interface{}
	Len             interface{}
	Arr             interface{}
}

func (_ *CljsCoreTransientArrayMap) CljsCoreITransientAssociative__() {}

func (tcoll *CljsCoreTransientArrayMap) X_assoc_BANG__Arity3(key interface{}, val interface{}) interface{} {
	if Truth_(tcoll.Editable_QMARK_) {
		{
			var idx = Array_map_index_of.X_invoke_Arity2(tcoll, key).(float64)
			_ = idx
			if idx == float64(-1) {
				if (tcoll.Len.(float64) + float64(2)) <= (float64(2) * CljsCorePersistentArrayMap_HASHMAP_THRESHOLD) {
					tcoll.Len = (tcoll.Len.(float64) + float64(2))

					js.JSArray_(&tcoll.Arr).Push(key)
					js.JSArray_(&tcoll.Arr).Push(val)
					return tcoll
				} else {
					return Assoc_BANG_.X_invoke_Arity3(Array__GT_transient_hash_map.X_invoke_Arity2(tcoll.Len, tcoll.Arr), key, val)
				}
			} else {
				if reflect.DeepEqual(val, Aget_(tcoll.Arr, (idx+float64(1)))) {
					return tcoll
				} else {
					tcoll.Arr.([]interface{})[int((idx + float64(1)))] = val
					return tcoll
				}
			}
		}
	} else {
		panic((&js.Error{"assoc! after persistent!"}))
	}
}

func (_ *CljsCoreTransientArrayMap) CljsCoreITransientCollection__() {}

func (tcoll *CljsCoreTransientArrayMap) X_conj_BANG__Arity2(o interface{}) interface{} {
	if Truth_(tcoll.Editable_QMARK_) {
		if Value_(o).Type().Implements(reflect.TypeOf((*CljsCoreIMapEntry)(nil)).Elem()) {
			return tcoll.X_assoc_BANG__Arity3(Key.X_invoke_Arity1(o), Val.X_invoke_Arity1(o))
		} else {
			{
				var es interface{} = Seq.Arity1IQ(o)
				var tcoll___1 interface{} = tcoll
				_, _ = es, tcoll___1
				for {
					{
						var temp__4220__auto__ = First.X_invoke_Arity1(es)
						_ = temp__4220__auto__
						if Truth_(temp__4220__auto__) {
							{
								var e = temp__4220__auto__
								_ = e
								es, tcoll___1 = Next.Arity1IQ(es), tcoll___1.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3(Key.X_invoke_Arity1(e), Val.X_invoke_Arity1(e))
								continue
							}
						} else {
							return tcoll___1
						}
					}
				}
			}
		}
	} else {
		panic((&js.Error{"conj! after persistent!"}))
	}
}

func (tcoll *CljsCoreTransientArrayMap) X_persistent_BANG__Arity1() interface{} {
	if Truth_(tcoll.Editable_QMARK_) {
		tcoll.Editable_QMARK_ = false

		return (&CljsCorePersistentArrayMap{nil, Quot.X_invoke_Arity2(tcoll.Len, float64(2)).(float64), tcoll.Arr, nil})
	} else {
		panic((&js.Error{"persistent! called twice"}))
	}
}

func (_ *CljsCoreTransientArrayMap) CljsCoreILookup__() {}

func (tcoll *CljsCoreTransientArrayMap) X_lookup_Arity2(k interface{}) interface{} {
	return tcoll.X_lookup_Arity3(k, nil)
}

func (tcoll *CljsCoreTransientArrayMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	if Truth_(tcoll.Editable_QMARK_) {
		{
			var idx = Array_map_index_of.X_invoke_Arity2(tcoll, k).(float64)
			_ = idx
			if idx == float64(-1) {
				return not_found
			} else {
				return Aget_(tcoll.Arr, (idx + float64(1)))
			}
		}
	} else {
		panic((&js.Error{"lookup after persistent!"}))
	}
}

func (_ *CljsCoreTransientArrayMap) CljsCoreICounted__() {}

func (tcoll *CljsCoreTransientArrayMap) X_count_Arity1() float64 {
	if Truth_(tcoll.Editable_QMARK_) {
		return Quot.X_invoke_Arity2(tcoll.Len, float64(2)).(float64)
	} else {
		panic((&js.Error{"count after persistent!"}))
	}
}

var X__GT_TransientArrayMap *AFn

var Array__GT_transient_hash_map *AFn

type CljsCoreBox struct{ Val interface{} }

var X__GT_Box *AFn

var Key_test *AFn

var Mask *AFn

var Clone_and_set *AFn

var Remove_pair *AFn

var Bitmap_indexed_node_index *AFn

var Bitpos *AFn

var Edit_and_set *AFn

var Inode_kv_reduce *AFn

type CljsCoreBitmapIndexedNode struct {
	Edit   interface{}
	Bitmap interface{}
	Arr    interface{}
}

func (_ *CljsCoreBitmapIndexedNode) CljsCoreObject__() {}

func (inode *CljsCoreBitmapIndexedNode) Ensure_editable(e interface{}) interface{} {
	if reflect.DeepEqual(e, inode.Edit) {
		return inode
	} else {
		{
			var n = Bit_count.X_invoke_Arity1(inode.Bitmap).(float64)
			var new_arr = make([]interface{}, int(func() float64 {
				if n < float64(0) {
					return float64(4)
				} else {
					return (float64(2) * (n + float64(1)))
				}
			}()))
			_, _ = n, new_arr
			Array_copy.X_invoke_Arity5(inode.Arr, float64(0), new_arr, float64(0), (float64(2) * n))
			return (&CljsCoreBitmapIndexedNode{e, inode.Bitmap, new_arr})
		}
	}
}

func (inode *CljsCoreBitmapIndexedNode) Inode_without_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, removed_leaf_QMARK_ interface{}) interface{} {
	{
		var bit = float64((Int32_(1) << UInt32_(float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f)))))
		_ = bit
		if float64((Int32_(inode.Bitmap.(float64)) & Int32_(bit))) == float64(0) {
			return inode
		} else {
			{
				var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
				var key_or_nil = Aget_(inode.Arr, (float64(2) * idx))
				var val_or_node = Aget_(inode.Arr, ((float64(2) * idx) + float64(1)))
				_, _, _ = idx, key_or_nil, val_or_node
				if Nil_(key_or_nil) {
					{
						var n = Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_without_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, removed_leaf_QMARK_})
						_ = n
						if reflect.DeepEqual(n, val_or_node) {
							return inode
						} else {
							if !(Nil_(n)) {
								return Edit_and_set.X_invoke_Arity4(inode, edit___1, ((float64(2) * idx) + float64(1)), n)
							} else {
								if inode.Bitmap.(float64) == bit {
									return nil
								} else {
									return inode.Edit_and_remove_pair(edit___1, bit, idx)

								}
							}
						}
					}
				} else {
					if Key_test.Arity2IIB(key, key_or_nil) {
						removed_leaf_QMARK_.(*CljsCoreBox).Val = float64(0)
						return inode.Edit_and_remove_pair(edit___1, bit, idx)
					} else {
						return inode

					}
				}
			}
		}
	}
}

func (inode *CljsCoreBitmapIndexedNode) Edit_and_remove_pair(e interface{}, bit interface{}, i interface{}) interface{} {
	if inode.Bitmap.(float64) == bit.(float64) {
		return nil
	} else {
		{
			var editable = inode.Ensure_editable(e)
			var earr = Native_get_instance_field.X_invoke_Arity2(editable, "Arr")
			var len = Alength_(earr)
			_, _, _ = editable, earr, len
			Native_set_instance_field.X_invoke_Arity3(editable, "Bitmap", float64((Int32_(bit.(float64)) ^ Int32_(Native_get_instance_field.X_invoke_Arity2(editable, "Bitmap").(float64)))))
			Array_copy.X_invoke_Arity5(earr, (float64(2) * (i.(float64) + float64(1))), earr, (float64(2) * i.(float64)), (len - (float64(2) * (i.(float64) + float64(1)))))
			earr.([]interface{})[int((len - float64(2)))] = nil
			earr.([]interface{})[int((len - float64(1)))] = nil
			return editable
		}
	}
}

func (inode *CljsCoreBitmapIndexedNode) Inode_seq() interface{} {
	return Create_inode_seq.X_invoke_Arity1(inode.Arr)
}

func (inode *CljsCoreBitmapIndexedNode) Kv_reduce(f interface{}, init interface{}) interface{} {
	return Inode_kv_reduce.X_invoke_Arity3(inode.Arr, f, init)
}

func (inode *CljsCoreBitmapIndexedNode) Inode_lookup(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var bit = float64((Int32_(1) << UInt32_(float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f)))))
		_ = bit
		if float64((Int32_(inode.Bitmap.(float64)) & Int32_(bit))) == float64(0) {
			return not_found
		} else {
			{
				var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
				var key_or_nil = Aget_(inode.Arr, (float64(2) * idx))
				var val_or_node = Aget_(inode.Arr, ((float64(2) * idx) + float64(1)))
				_, _, _ = idx, key_or_nil, val_or_node
				if Nil_(key_or_nil) {
					return Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_lookup", []interface{}{(shift.(float64) + float64(5)), hash, key, not_found})
				} else {
					if Key_test.Arity2IIB(key, key_or_nil) {
						return val_or_node
					} else {
						return not_found

					}
				}
			}
		}
	}
}

func (inode *CljsCoreBitmapIndexedNode) Inode_assoc_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	{
		var bit = float64((Int32_(1) << UInt32_(float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f)))))
		var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
		_, _ = bit, idx
		if float64((Int32_(inode.Bitmap.(float64)) & Int32_(bit))) == float64(0) {
			{
				var n = Bit_count.X_invoke_Arity1(inode.Bitmap).(float64)
				_ = n
				if (float64(2) * n) < Alength_(inode.Arr) {
					{
						var editable = inode.Ensure_editable(edit___1)
						var earr = Native_get_instance_field.X_invoke_Arity2(editable, "Arr")
						_, _ = editable, earr
						Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)
						Array_copy_downward.X_invoke_Arity5(earr, (float64(2) * idx), earr, (float64(2) * (idx + float64(1))), (float64(2) * (n - idx)))
						earr.([]interface{})[int((float64(2) * idx))] = key
						earr.([]interface{})[int(((float64(2) * idx) + float64(1)))] = val
						Native_set_instance_field.X_invoke_Arity3(editable, "Bitmap", float64((Int32_(Native_get_instance_field.X_invoke_Arity2(editable, "Bitmap").(float64)) | Int32_(bit))))
						return editable
					}
				} else {
					if n >= float64(16) {
						{
							var nodes = make([]interface{}, int(float64(32)))
							var jdx = float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f))
							_, _ = nodes, jdx
							nodes[int(jdx)] = Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
							{
								var i_7931 = float64(0)
								var j_7932 = float64(0)
								_, _ = i_7931, j_7932
								for {
									if i_7931 < float64(32) {
										if float64((Int32_(float64((UInt32_(inode.Bitmap.(float64)) >> UInt32_(float64((32+Int32_(i_7931))%32))))) & Int32_(float64(1)))) == float64(0) {
											i_7931, j_7932 = (i_7931 + float64(1)), j_7932
											continue
										} else {
											nodes[int(i_7931)] = func() interface{} {
												if !(Nil_(Aget_(inode.Arr, j_7932))) {
													return Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), Hash.X_invoke_Arity1(Aget_(inode.Arr, j_7932)), Aget_(inode.Arr, j_7932), Aget_(inode.Arr, (j_7932 + float64(1))), added_leaf_QMARK_})
												} else {
													return Aget_(inode.Arr, (j_7932 + float64(1)))
												}
											}()
											i_7931, j_7932 = (i_7931 + float64(1)), (j_7932 + float64(2))
											continue
										}
									} else {
									}
									break
								}
							}
							return (&CljsCoreArrayNode{edit___1, (n + float64(1)), nodes})
						}
					} else {
						{
							var new_arr = make([]interface{}, int((float64(2) * (n + float64(4)))))
							_ = new_arr
							Array_copy.X_invoke_Arity5(inode.Arr, float64(0), new_arr, float64(0), (float64(2) * idx))
							new_arr[int((float64(2) * idx))] = key
							new_arr[int(((float64(2) * idx) + float64(1)))] = val
							Array_copy.X_invoke_Arity5(inode.Arr, (float64(2) * idx), new_arr, (float64(2) * (idx + float64(1))), (float64(2) * (n - idx)))
							Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)
							{
								var editable = inode.Ensure_editable(edit___1)
								_ = editable
								Native_set_instance_field.X_invoke_Arity3(editable, "Arr", new_arr)
								Native_set_instance_field.X_invoke_Arity3(editable, "Bitmap", float64((Int32_(Native_get_instance_field.X_invoke_Arity2(editable, "Bitmap").(float64)) | Int32_(bit))))
								return editable
							}
						}

					}
				}
			}
		} else {
			{
				var key_or_nil = Aget_(inode.Arr, (float64(2) * idx))
				var val_or_node = Aget_(inode.Arr, ((float64(2) * idx) + float64(1)))
				_, _ = key_or_nil, val_or_node
				if Nil_(key_or_nil) {
					{
						var n = Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
						_ = n
						if reflect.DeepEqual(n, val_or_node) {
							return inode
						} else {
							return Edit_and_set.X_invoke_Arity4(inode, edit___1, ((float64(2) * idx) + float64(1)), n)
						}
					}
				} else {
					if Key_test.Arity2IIB(key, key_or_nil) {
						if reflect.DeepEqual(val, val_or_node) {
							return inode
						} else {
							return Edit_and_set.X_invoke_Arity4(inode, edit___1, ((float64(2) * idx) + float64(1)), val)
						}
					} else {
						Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)
						return Edit_and_set.X_invoke_Arity6(inode, edit___1, (float64(2) * idx), nil, ((float64(2) * idx) + float64(1)), Create_node.X_invoke_Arity7(edit___1, (shift.(float64)+float64(5)), key_or_nil, val_or_node, hash, key, val))

					}
				}
			}
		}
	}
}

func (inode *CljsCoreBitmapIndexedNode) Inode_assoc(shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	{
		var bit = float64((Int32_(1) << UInt32_(float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f)))))
		var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
		_, _ = bit, idx
		if float64((Int32_(inode.Bitmap.(float64)) & Int32_(bit))) == float64(0) {
			{
				var n = Bit_count.X_invoke_Arity1(inode.Bitmap).(float64)
				_ = n
				if n >= float64(16) {
					{
						var nodes = make([]interface{}, int(float64(32)))
						var jdx = float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f))
						_, _ = nodes, jdx
						nodes[int(jdx)] = Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
						{
							var i_7933 = float64(0)
							var j_7934 = float64(0)
							_, _ = i_7933, j_7934
							for {
								if i_7933 < float64(32) {
									if float64((Int32_(float64((UInt32_(inode.Bitmap.(float64)) >> UInt32_(float64((32+Int32_(i_7933))%32))))) & Int32_(float64(1)))) == float64(0) {
										i_7933, j_7934 = (i_7933 + float64(1)), j_7934
										continue
									} else {
										nodes[int(i_7933)] = func() interface{} {
											if !(Nil_(Aget_(inode.Arr, j_7934))) {
												return Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), Hash.X_invoke_Arity1(Aget_(inode.Arr, j_7934)), Aget_(inode.Arr, j_7934), Aget_(inode.Arr, (j_7934 + float64(1))), added_leaf_QMARK_})
											} else {
												return Aget_(inode.Arr, (j_7934 + float64(1)))
											}
										}()
										i_7933, j_7934 = (i_7933 + float64(1)), (j_7934 + float64(2))
										continue
									}
								} else {
								}
								break
							}
						}
						return (&CljsCoreArrayNode{nil, (n + float64(1)), nodes})
					}
				} else {
					{
						var new_arr = make([]interface{}, int((float64(2) * (n + float64(1)))))
						_ = new_arr
						Array_copy.X_invoke_Arity5(inode.Arr, float64(0), new_arr, float64(0), (float64(2) * idx))
						new_arr[int((float64(2) * idx))] = key
						new_arr[int(((float64(2) * idx) + float64(1)))] = val
						Array_copy.X_invoke_Arity5(inode.Arr, (float64(2) * idx), new_arr, (float64(2) * (idx + float64(1))), (float64(2) * (n - idx)))
						Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)
						return (&CljsCoreBitmapIndexedNode{nil, float64((Int32_(inode.Bitmap.(float64)) | Int32_(bit))), new_arr})
					}
				}
			}
		} else {
			{
				var key_or_nil = Aget_(inode.Arr, (float64(2) * idx))
				var val_or_node = Aget_(inode.Arr, ((float64(2) * idx) + float64(1)))
				_, _ = key_or_nil, val_or_node
				if Nil_(key_or_nil) {
					{
						var n = Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
						_ = n
						if reflect.DeepEqual(n, val_or_node) {
							return inode
						} else {
							return (&CljsCoreBitmapIndexedNode{nil, inode.Bitmap, Clone_and_set.X_invoke_Arity3(inode.Arr, ((float64(2) * idx) + float64(1)), n).([]interface{})})
						}
					}
				} else {
					if Key_test.Arity2IIB(key, key_or_nil) {
						if reflect.DeepEqual(val, val_or_node) {
							return inode
						} else {
							return (&CljsCoreBitmapIndexedNode{nil, inode.Bitmap, Clone_and_set.X_invoke_Arity3(inode.Arr, ((float64(2) * idx) + float64(1)), val).([]interface{})})
						}
					} else {
						Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)
						return (&CljsCoreBitmapIndexedNode{nil, inode.Bitmap, Clone_and_set.X_invoke_Arity5(inode.Arr, (float64(2) * idx), nil, ((float64(2) * idx) + float64(1)), Create_node.X_invoke_Arity6((shift.(float64)+float64(5)), key_or_nil, val_or_node, hash, key, val)).([]interface{})})

					}
				}
			}
		}
	}
}

func (inode *CljsCoreBitmapIndexedNode) Inode_find(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var bit = float64((Int32_(1) << UInt32_(float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f)))))
		_ = bit
		if float64((Int32_(inode.Bitmap.(float64)) & Int32_(bit))) == float64(0) {
			return not_found
		} else {
			{
				var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
				var key_or_nil = Aget_(inode.Arr, (float64(2) * idx))
				var val_or_node = Aget_(inode.Arr, ((float64(2) * idx) + float64(1)))
				_, _, _ = idx, key_or_nil, val_or_node
				if Nil_(key_or_nil) {
					return Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_find", []interface{}{(shift.(float64) + float64(5)), hash, key, not_found})
				} else {
					if Key_test.Arity2IIB(key, key_or_nil) {
						return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{key_or_nil, val_or_node}, nil})
					} else {
						return not_found

					}
				}
			}
		}
	}
}

func (inode *CljsCoreBitmapIndexedNode) Inode_without(shift interface{}, hash interface{}, key interface{}) interface{} {
	{
		var bit = float64((Int32_(1) << UInt32_(float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f)))))
		_ = bit
		if float64((Int32_(inode.Bitmap.(float64)) & Int32_(bit))) == float64(0) {
			return inode
		} else {
			{
				var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
				var key_or_nil = Aget_(inode.Arr, (float64(2) * idx))
				var val_or_node = Aget_(inode.Arr, ((float64(2) * idx) + float64(1)))
				_, _, _ = idx, key_or_nil, val_or_node
				if Nil_(key_or_nil) {
					{
						var n = Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_without", []interface{}{(shift.(float64) + float64(5)), hash, key})
						_ = n
						if reflect.DeepEqual(n, val_or_node) {
							return inode
						} else {
							if !(Nil_(n)) {
								return (&CljsCoreBitmapIndexedNode{nil, inode.Bitmap, Clone_and_set.X_invoke_Arity3(inode.Arr, ((float64(2) * idx) + float64(1)), n).([]interface{})})
							} else {
								if inode.Bitmap.(float64) == bit {
									return nil
								} else {
									return (&CljsCoreBitmapIndexedNode{nil, float64((Int32_(inode.Bitmap.(float64)) ^ Int32_(bit))), Remove_pair.X_invoke_Arity2(inode.Arr, idx).([]interface{})})

								}
							}
						}
					}
				} else {
					if Key_test.Arity2IIB(key, key_or_nil) {
						return (&CljsCoreBitmapIndexedNode{nil, float64((Int32_(inode.Bitmap.(float64)) ^ Int32_(bit))), Remove_pair.X_invoke_Arity2(inode.Arr, idx).([]interface{})})
					} else {
						return inode

					}
				}
			}
		}
	}
}

var X__GT_BitmapIndexedNode *AFn

var CljsCoreBitmapIndexedNode_EMPTY = (&CljsCoreBitmapIndexedNode{nil, float64(0), make([]interface{}, int(float64(0)))})

var Pack_array_node *AFn

type CljsCoreArrayNode struct {
	Edit interface{}
	Cnt  interface{}
	Arr  interface{}
}

func (_ *CljsCoreArrayNode) CljsCoreObject__() {}

func (inode *CljsCoreArrayNode) Ensure_editable(e interface{}) interface{} {
	if reflect.DeepEqual(e, inode.Edit) {
		return inode
	} else {
		return (&CljsCoreArrayNode{e, inode.Cnt, Aclone.X_invoke_Arity1(inode.Arr).([]interface{})})
	}
}

func (inode *CljsCoreArrayNode) Inode_without_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, removed_leaf_QMARK_ interface{}) interface{} {
	{
		var idx = float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f))
		var node = Aget_(inode.Arr, idx)
		_, _ = idx, node
		if Nil_(node) {
			return inode
		} else {
			{
				var n = Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_without_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, removed_leaf_QMARK_})
				_ = n
				if reflect.DeepEqual(n, node) {
					return inode
				} else {
					if Nil_(n) {
						if inode.Cnt.(float64) <= float64(8) {
							return Pack_array_node.X_invoke_Arity3(inode, edit___1, idx).(*CljsCoreBitmapIndexedNode)
						} else {
							{
								var editable = Edit_and_set.X_invoke_Arity4(inode, edit___1, idx, n)
								_ = editable
								Native_set_instance_field.X_invoke_Arity3(editable, "Cnt", (Native_get_instance_field.X_invoke_Arity2(editable, "Cnt").(float64) - float64(1)))
								return editable
							}
						}
					} else {
						return Edit_and_set.X_invoke_Arity4(inode, edit___1, idx, n)

					}
				}
			}
		}
	}
}

func (inode *CljsCoreArrayNode) Inode_seq() interface{} {
	return Create_array_node_seq.X_invoke_Arity1(inode.Arr)
}

func (inode *CljsCoreArrayNode) Kv_reduce(f interface{}, init interface{}) interface{} {
	{
		var len = Alength_(inode.Arr)
		_ = len
		{
			var i = float64(0)
			var init___1 interface{} = init
			_, _ = i, init___1
			for {
				if i < len {
					{
						var node = Aget_(inode.Arr, i)
						_ = node
						if !(Nil_(node)) {
							{
								var init___2 = Native_invoke_instance_method.X_invoke_Arity3(node, "Kv_reduce", []interface{}{f, init___1})
								_ = init___2
								if Reduced_QMARK_.Arity1IB(init___2) {
									return Deref.X_invoke_Arity1(init___2)
								} else {
									i, init___1 = (i + float64(1)), init___2
									continue
								}
							}
						} else {
							i, init___1 = (i + float64(1)), init___1
							continue
						}
					}
				} else {
					return init___1
				}
			}
		}
	}
}

func (inode *CljsCoreArrayNode) Inode_lookup(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var idx = float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f))
		var node = Aget_(inode.Arr, idx)
		_, _ = idx, node
		if !(Nil_(node)) {
			return Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_lookup", []interface{}{(shift.(float64) + float64(5)), hash, key, not_found})
		} else {
			return not_found
		}
	}
}

func (inode *CljsCoreArrayNode) Inode_assoc_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	{
		var idx = float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f))
		var node = Aget_(inode.Arr, idx)
		_, _ = idx, node
		if Nil_(node) {
			{
				var editable = Edit_and_set.X_invoke_Arity4(inode, edit___1, idx, Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_}))
				_ = editable
				Native_set_instance_field.X_invoke_Arity3(editable, "Cnt", (Native_get_instance_field.X_invoke_Arity2(editable, "Cnt").(float64) + float64(1)))
				return editable
			}
		} else {
			{
				var n = Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
				_ = n
				if reflect.DeepEqual(n, node) {
					return inode
				} else {
					return Edit_and_set.X_invoke_Arity4(inode, edit___1, idx, n)
				}
			}
		}
	}
}

func (inode *CljsCoreArrayNode) Inode_assoc(shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	{
		var idx = float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f))
		var node = Aget_(inode.Arr, idx)
		_, _ = idx, node
		if Nil_(node) {
			return (&CljsCoreArrayNode{nil, (inode.Cnt.(float64) + float64(1)), Clone_and_set.X_invoke_Arity3(inode.Arr, idx, Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})).([]interface{})})
		} else {
			{
				var n = Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
				_ = n
				if reflect.DeepEqual(n, node) {
					return inode
				} else {
					return (&CljsCoreArrayNode{nil, inode.Cnt, Clone_and_set.X_invoke_Arity3(inode.Arr, idx, n).([]interface{})})
				}
			}
		}
	}
}

func (inode *CljsCoreArrayNode) Inode_find(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var idx = float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f))
		var node = Aget_(inode.Arr, idx)
		_, _ = idx, node
		if !(Nil_(node)) {
			return Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_find", []interface{}{(shift.(float64) + float64(5)), hash, key, not_found})
		} else {
			return not_found
		}
	}
}

func (inode *CljsCoreArrayNode) Inode_without(shift interface{}, hash interface{}, key interface{}) interface{} {
	{
		var idx = float64(((UInt32_(hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f))
		var node = Aget_(inode.Arr, idx)
		_, _ = idx, node
		if !(Nil_(node)) {
			{
				var n = Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_without", []interface{}{(shift.(float64) + float64(5)), hash, key})
				_ = n
				if reflect.DeepEqual(n, node) {
					return inode
				} else {
					if Nil_(n) {
						if inode.Cnt.(float64) <= float64(8) {
							return Pack_array_node.X_invoke_Arity3(inode, nil, idx).(*CljsCoreBitmapIndexedNode)
						} else {
							return (&CljsCoreArrayNode{nil, (inode.Cnt.(float64) - float64(1)), Clone_and_set.X_invoke_Arity3(inode.Arr, idx, n).([]interface{})})
						}
					} else {
						return (&CljsCoreArrayNode{nil, inode.Cnt, Clone_and_set.X_invoke_Arity3(inode.Arr, idx, n).([]interface{})})

					}
				}
			}
		} else {
			return inode
		}
	}
}

var X__GT_ArrayNode *AFn

var Hash_collision_node_find_index *AFn

type CljsCoreHashCollisionNode struct {
	Edit           interface{}
	Collision_hash interface{}
	Cnt            interface{}
	Arr            interface{}
}

func (_ *CljsCoreHashCollisionNode) CljsCoreObject__() {}

func (inode *CljsCoreHashCollisionNode) Ensure_editable(e interface{}) interface{} {
	if reflect.DeepEqual(e, inode.Edit) {
		return inode
	} else {
		{
			var new_arr = make([]interface{}, int((float64(2) * (inode.Cnt.(float64) + float64(1)))))
			_ = new_arr
			Array_copy.X_invoke_Arity5(inode.Arr, float64(0), new_arr, float64(0), (float64(2) * inode.Cnt.(float64)))
			return (&CljsCoreHashCollisionNode{e, inode.Collision_hash, inode.Cnt, new_arr})
		}
	}
}

func (inode *CljsCoreHashCollisionNode) Inode_without_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, removed_leaf_QMARK_ interface{}) interface{} {
	{
		var idx = Hash_collision_node_find_index.X_invoke_Arity3(inode.Arr, inode.Cnt, key).(float64)
		_ = idx
		if idx == float64(-1) {
			return inode
		} else {
			removed_leaf_QMARK_.(*CljsCoreBox).Val = float64(0)
			if inode.Cnt.(float64) == float64(1) {
				return nil
			} else {
				{
					var editable = inode.Ensure_editable(edit___1)
					var earr = Native_get_instance_field.X_invoke_Arity2(editable, "Arr")
					_, _ = editable, earr
					earr.([]interface{})[int(idx)] = Aget_(earr, ((float64(2) * inode.Cnt.(float64)) - float64(2)))
					earr.([]interface{})[int((idx + float64(1)))] = Aget_(earr, ((float64(2) * inode.Cnt.(float64)) - float64(1)))
					earr.([]interface{})[int(((float64(2) * inode.Cnt.(float64)) - float64(1)))] = nil
					earr.([]interface{})[int(((float64(2) * inode.Cnt.(float64)) - float64(2)))] = nil
					Native_set_instance_field.X_invoke_Arity3(editable, "Cnt", (Native_get_instance_field.X_invoke_Arity2(editable, "Cnt").(float64) - float64(1)))
					return editable
				}
			}
		}
	}
}

func (inode *CljsCoreHashCollisionNode) Inode_seq() interface{} {
	return Create_inode_seq.X_invoke_Arity1(inode.Arr)
}

func (inode *CljsCoreHashCollisionNode) Kv_reduce(f interface{}, init interface{}) interface{} {
	return Inode_kv_reduce.X_invoke_Arity3(inode.Arr, f, init)
}

func (inode *CljsCoreHashCollisionNode) Inode_lookup(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var idx = Hash_collision_node_find_index.X_invoke_Arity3(inode.Arr, inode.Cnt, key).(float64)
		_ = idx
		if idx < float64(0) {
			return not_found
		} else {
			if Key_test.Arity2IIB(key, Aget_(inode.Arr, idx)) {
				return Aget_(inode.Arr, (idx + float64(1)))
			} else {
				return not_found

			}
		}
	}
}

func (inode *CljsCoreHashCollisionNode) Inode_assoc_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	if hash.(float64) == inode.Collision_hash.(float64) {
		{
			var idx = Hash_collision_node_find_index.X_invoke_Arity3(inode.Arr, inode.Cnt, key).(float64)
			_ = idx
			if idx == float64(-1) {
				if Alength_(inode.Arr) > (float64(2) * inode.Cnt.(float64)) {
					{
						var editable = Edit_and_set.X_invoke_Arity6(inode, edit___1, (float64(2) * inode.Cnt.(float64)), key, ((float64(2) * inode.Cnt.(float64)) + float64(1)), val)
						_ = editable
						Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)
						Native_set_instance_field.X_invoke_Arity3(editable, "Cnt", (Native_get_instance_field.X_invoke_Arity2(editable, "Cnt").(float64) + float64(1)))
						return editable
					}
				} else {
					{
						var len = Alength_(inode.Arr)
						var new_arr = make([]interface{}, int((len + float64(2))))
						_, _ = len, new_arr
						Array_copy.X_invoke_Arity5(inode.Arr, float64(0), new_arr, float64(0), len)
						new_arr[int(len)] = key
						new_arr[int((len + float64(1)))] = val
						Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)
						return inode.Ensure_editable_array(edit___1, (inode.Cnt.(float64) + float64(1)), new_arr)
					}
				}
			} else {
				if reflect.DeepEqual(Aget_(inode.Arr, (idx+float64(1))), val) {
					return inode
				} else {
					return Edit_and_set.X_invoke_Arity4(inode, edit___1, (idx + float64(1)), val)
				}
			}
		}
	} else {
		return (&CljsCoreBitmapIndexedNode{edit___1, float64((Int32_(1) << UInt32_(float64(((UInt32_(inode.Collision_hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f))))), []interface{}{nil, inode, nil, nil}}).Inode_assoc_BANG_(edit___1, shift, hash, key, val, added_leaf_QMARK_)
	}
}

func (inode *CljsCoreHashCollisionNode) Inode_assoc(shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	if hash.(float64) == inode.Collision_hash.(float64) {
		{
			var idx = Hash_collision_node_find_index.X_invoke_Arity3(inode.Arr, inode.Cnt, key).(float64)
			_ = idx
			if idx == float64(-1) {
				{
					var len = (float64(2) * inode.Cnt.(float64))
					var new_arr = make([]interface{}, int((len + float64(2))))
					_, _ = len, new_arr
					Array_copy.X_invoke_Arity5(inode.Arr, float64(0), new_arr, float64(0), len)
					new_arr[int(len)] = key
					new_arr[int((len + float64(1)))] = val
					Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)
					return (&CljsCoreHashCollisionNode{nil, inode.Collision_hash, (inode.Cnt.(float64) + float64(1)), new_arr})
				}
			} else {
				if X_EQ_.Arity2IIB(Aget_(inode.Arr, idx), val) {
					return inode
				} else {
					return (&CljsCoreHashCollisionNode{nil, inode.Collision_hash, inode.Cnt, Clone_and_set.X_invoke_Arity3(inode.Arr, (idx + float64(1)), val).([]interface{})})
				}
			}
		}
	} else {
		return (&CljsCoreBitmapIndexedNode{nil, float64((Int32_(1) << UInt32_(float64(((UInt32_(inode.Collision_hash.(float64)) >> UInt32_(shift.(float64))) & 0x01f))))), []interface{}{nil, inode}}).Inode_assoc(shift, hash, key, val, added_leaf_QMARK_)
	}
}

func (inode *CljsCoreHashCollisionNode) Ensure_editable_array(e interface{}, count interface{}, array interface{}) interface{} {
	if reflect.DeepEqual(e, inode.Edit) {
		inode.Arr = array

		inode.Cnt = count

		return inode
	} else {
		return (&CljsCoreHashCollisionNode{inode.Edit, inode.Collision_hash, count, array})
	}
}

func (inode *CljsCoreHashCollisionNode) Inode_find(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var idx = Hash_collision_node_find_index.X_invoke_Arity3(inode.Arr, inode.Cnt, key).(float64)
		_ = idx
		if idx < float64(0) {
			return not_found
		} else {
			if Key_test.Arity2IIB(key, Aget_(inode.Arr, idx)) {
				return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Aget_(inode.Arr, idx), Aget_(inode.Arr, (idx + float64(1)))}, nil})
			} else {
				return not_found

			}
		}
	}
}

func (inode *CljsCoreHashCollisionNode) Inode_without(shift interface{}, hash interface{}, key interface{}) interface{} {
	{
		var idx = Hash_collision_node_find_index.X_invoke_Arity3(inode.Arr, inode.Cnt, key).(float64)
		_ = idx
		if idx == float64(-1) {
			return inode
		} else {
			if inode.Cnt.(float64) == float64(1) {
				return nil
			} else {
				return (&CljsCoreHashCollisionNode{nil, inode.Collision_hash, (inode.Cnt.(float64) - float64(1)), Remove_pair.X_invoke_Arity2(inode.Arr, Quot.X_invoke_Arity2(idx, float64(2)).(float64)).([]interface{})})

			}
		}
	}
}

var X__GT_HashCollisionNode *AFn

var Create_node *AFn

type CljsCoreNodeSeq struct {
	Meta    interface{}
	Nodes   interface{}
	I       interface{}
	S       interface{}
	X__hash interface{}
}

func (_ *CljsCoreNodeSeq) CljsCoreObject__() {}

func (coll *CljsCoreNodeSeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreNodeSeq) String() string {
	return this.ToString()
}

func (this *CljsCoreNodeSeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreNodeSeq) CljsCoreIMeta__() {}

func (coll *CljsCoreNodeSeq) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreNodeSeq) CljsCoreIHash__() {}

func (coll *CljsCoreNodeSeq) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreIEquiv__() {}

func (coll *CljsCoreNodeSeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreNodeSeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreNodeSeq) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.Meta)
}

func (_ *CljsCoreNodeSeq) CljsCoreIReduce__() {}

func (coll *CljsCoreNodeSeq) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCoreNodeSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCoreNodeSeq) CljsCoreISeq__() {}

func (coll *CljsCoreNodeSeq) X_first_Arity1() interface{} {
	if Nil_(coll.S) {
		return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Aget_(coll.Nodes, coll.I.(float64)), Aget_(coll.Nodes, (coll.I.(float64) + float64(1)))}, nil})
	} else {
		return First.X_invoke_Arity1(coll.S)
	}
}

func (coll *CljsCoreNodeSeq) X_rest_Arity1() interface{} {
	if Nil_(coll.S) {
		return Create_inode_seq.X_invoke_Arity3(coll.Nodes, (coll.I.(float64) + float64(2)), nil)
	} else {
		return Create_inode_seq.X_invoke_Arity3(coll.Nodes, coll.I, Next.Arity1IQ(coll.S))
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreISeqable__() {}

func (this *CljsCoreNodeSeq) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCoreNodeSeq) CljsCoreISequential__() {}

func (_ *CljsCoreNodeSeq) CljsCoreIWithMeta__() {}

func (coll *CljsCoreNodeSeq) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCoreNodeSeq{meta___1, coll.Nodes, coll.I, coll.S, coll.X__hash})
}

func (_ *CljsCoreNodeSeq) CljsCoreICollection__() {}

func (coll *CljsCoreNodeSeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

var X__GT_NodeSeq *AFn

var Create_inode_seq *AFn

type CljsCoreArrayNodeSeq struct {
	Meta    interface{}
	Nodes   interface{}
	I       interface{}
	S       interface{}
	X__hash interface{}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreObject__() {}

func (coll *CljsCoreArrayNodeSeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreArrayNodeSeq) String() string {
	return this.ToString()
}

func (this *CljsCoreArrayNodeSeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIMeta__() {}

func (coll *CljsCoreArrayNodeSeq) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIHash__() {}

func (coll *CljsCoreArrayNodeSeq) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIEquiv__() {}

func (coll *CljsCoreArrayNodeSeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreArrayNodeSeq) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.Meta)
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIReduce__() {}

func (coll *CljsCoreArrayNodeSeq) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCoreArrayNodeSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreISeq__() {}

func (coll *CljsCoreArrayNodeSeq) X_first_Arity1() interface{} {
	return First.X_invoke_Arity1(coll.S)
}

func (coll *CljsCoreArrayNodeSeq) X_rest_Arity1() interface{} {
	return Create_array_node_seq.X_invoke_Arity4(nil, coll.Nodes, coll.I, Next.Arity1IQ(coll.S))
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreISeqable__() {}

func (this *CljsCoreArrayNodeSeq) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreISequential__() {}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIWithMeta__() {}

func (coll *CljsCoreArrayNodeSeq) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCoreArrayNodeSeq{meta___1, coll.Nodes, coll.I, coll.S, coll.X__hash})
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreICollection__() {}

func (coll *CljsCoreArrayNodeSeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

var X__GT_ArrayNodeSeq *AFn

var Create_array_node_seq *AFn

type CljsCorePersistentHashMap struct {
	Meta           interface{}
	Cnt            interface{}
	Root           interface{}
	Has_nil_QMARK_ bool
	Nil_val        interface{}
	X__hash        interface{}
}

func (_ *CljsCorePersistentHashMap) CljsCoreObject__() {}

func (coll *CljsCorePersistentHashMap) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCorePersistentHashMap) String() string {
	return this.ToString()
}

func (this *CljsCorePersistentHashMap) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (coll *CljsCorePersistentHashMap) Keys() interface{} {
	return Iterator.X_invoke_Arity1(Native_invoke_func.X_invoke_Arity2((*CljsCorePersistentHashMap).Keys, []interface{}{coll}).(*CljsCoreIterator)).(*CljsCoreIterator)
}

func (coll *CljsCorePersistentHashMap) Entries() interface{} {
	return Entries_iterator.X_invoke_Arity1(Seq.Arity1IQ(coll)).(*CljsCoreEntriesIterator)
}

func (coll *CljsCorePersistentHashMap) Values() interface{} {
	return Iterator.X_invoke_Arity1(Vals.X_invoke_Arity1(coll)).(*CljsCoreIterator)
}

func (coll *CljsCorePersistentHashMap) Has(k interface{}) interface{} {
	return Contains_QMARK_.Arity2IIB(coll, k)
}

func (coll *CljsCorePersistentHashMap) Get(k interface{}) interface{} {
	return coll.X_lookup_Arity2(k)
}

func (coll *CljsCorePersistentHashMap) ForEach(f interface{}) interface{} {
	{
		var seq__6256 interface{} = Seq.Arity1IQ(coll)
		var chunk__6257 interface{} = nil
		var count__6258 = float64(0)
		var i__6259 = float64(0)
		_, _, _, _ = seq__6256, chunk__6257, count__6258, i__6259
		for {
			if i__6259 < count__6258 {
				{
					var vec__6260 = chunk__6257.(CljsCoreIIndexed).X_nth_Arity2(i__6259)
					var k = Nth.X_invoke_Arity3(vec__6260, float64(0), nil)
					var v = Nth.X_invoke_Arity3(vec__6260, float64(1), nil)
					_, _, _ = vec__6260, k, v
					{
						var G__6261_7935 = v
						var G__6262_7936 = k
						_, _ = G__6261_7935, G__6262_7936
						f.(CljsCoreIFn).X_invoke_Arity2(G__6261_7935, G__6262_7936)
					}
					seq__6256, chunk__6257, count__6258, i__6259 = seq__6256, chunk__6257, count__6258, (i__6259 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__6256)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__6256___1 = temp__4222__auto__
							_ = seq__6256___1
							if Chunked_seq_QMARK_.Arity1IB(seq__6256___1) {
								{
									var c__954__auto__ = Chunk_first.X_invoke_Arity1(seq__6256___1)
									_ = c__954__auto__
									seq__6256, chunk__6257, count__6258, i__6259 = Chunk_rest.X_invoke_Arity1(seq__6256___1), c__954__auto__, Count.X_invoke_Arity1(c__954__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__6263 = First.X_invoke_Arity1(seq__6256___1)
									var k = Nth.X_invoke_Arity3(vec__6263, float64(0), nil)
									var v = Nth.X_invoke_Arity3(vec__6263, float64(1), nil)
									_, _, _ = vec__6263, k, v
									{
										var G__6264_7937 = v
										var G__6265_7938 = k
										_, _ = G__6264_7937, G__6265_7938
										f.(CljsCoreIFn).X_invoke_Arity2(G__6264_7937, G__6265_7938)
									}
									seq__6256, chunk__6257, count__6258, i__6259 = Next.Arity1IQ(seq__6256___1), nil, float64(0), float64(0)
									continue
								}
							}
						}
					} else {
						return nil
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreILookup__() {}

func (coll *CljsCorePersistentHashMap) X_lookup_Arity2(k interface{}) interface{} {
	return coll.X_lookup_Arity3(k, nil)
}

func (coll *CljsCorePersistentHashMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	if Nil_(k) {
		if coll.Has_nil_QMARK_ {
			return coll.Nil_val
		} else {
			return not_found
		}
	} else {
		if Nil_(coll.Root) {
			return not_found
		} else {
			return Native_invoke_instance_method.X_invoke_Arity3(coll.Root, "Inode_lookup", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k, not_found})

		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIKVReduce__() {}

func (coll *CljsCorePersistentHashMap) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var init___1 = func() interface{} {
			if coll.Has_nil_QMARK_ {
				return func() interface{} {
					var G__6272 = init
					var G__6273 interface{} = nil
					var G__6274 = coll.Nil_val
					_, _, _ = G__6272, G__6273, G__6274
					return f.(CljsCoreIFn).X_invoke_Arity3(G__6272, G__6273, G__6274)
				}()
			} else {
				return init
			}
		}()
		_ = init___1
		if Reduced_QMARK_.Arity1IB(init___1) {
			return Deref.X_invoke_Arity1(init___1)
		} else {
			if !(Nil_(coll.Root)) {
				return Native_invoke_instance_method.X_invoke_Arity3(coll.Root, "Kv_reduce", []interface{}{f, init___1})
			} else {
				return init___1

			}
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIMeta__() {}

func (coll *CljsCorePersistentHashMap) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCorePersistentHashMap) CljsCoreICloneable__() {}

func (___ *CljsCorePersistentHashMap) X_clone_Arity1() interface{} {
	return (&CljsCorePersistentHashMap{___.Meta, ___.Cnt, ___.Root, ___.Has_nil_QMARK_, ___.Nil_val, ___.X__hash})
}

func (_ *CljsCorePersistentHashMap) CljsCoreICounted__() {}

func (coll *CljsCorePersistentHashMap) X_count_Arity1() float64 {
	return coll.Cnt.(float64)
}

func (_ *CljsCorePersistentHashMap) CljsCoreIHash__() {}

func (coll *CljsCorePersistentHashMap) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentHashMap) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_map.X_invoke_Arity2(coll, other))
}

func (_ *CljsCorePersistentHashMap) CljsCoreIEditableCollection__() {}

func (coll *CljsCorePersistentHashMap) X_as_transient_Arity1() interface{} {
	return (&CljsCoreTransientHashMap{true, coll.Root, coll.Cnt, coll.Has_nil_QMARK_, coll.Nil_val})
}

func (_ *CljsCorePersistentHashMap) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCorePersistentHashMap) X_empty_Arity1() interface{} {
	return CljsCorePersistentHashMap_EMPTY.X_with_meta_Arity2(coll.Meta)
}

func (_ *CljsCorePersistentHashMap) CljsCoreIMap__() {}

func (coll *CljsCorePersistentHashMap) X_dissoc_Arity2(k interface{}) interface{} {
	if Nil_(k) {
		if coll.Has_nil_QMARK_ {
			return (&CljsCorePersistentHashMap{coll.Meta, (coll.Cnt.(float64) - float64(1)), coll.Root, false, nil, nil})
		} else {
			return coll
		}
	} else {
		if Nil_(coll.Root) {
			return coll
		} else {
			{
				var new_root = Native_invoke_instance_method.X_invoke_Arity3(coll.Root, "Inode_without", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k})
				_ = new_root
				if reflect.DeepEqual(new_root, coll.Root) {
					return coll
				} else {
					return (&CljsCorePersistentHashMap{coll.Meta, (coll.Cnt.(float64) - float64(1)), new_root, coll.Has_nil_QMARK_, coll.Nil_val, nil})
				}
			}

		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIAssociative__() {}

func (coll *CljsCorePersistentHashMap) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	if Nil_(k) {
		if (coll.Has_nil_QMARK_) && (reflect.DeepEqual(v, coll.Nil_val)) {
			return coll
		} else {
			return (&CljsCorePersistentHashMap{coll.Meta, func() interface{} {
				if coll.Has_nil_QMARK_ {
					return coll.Cnt
				} else {
					return (coll.Cnt.(float64) + float64(1))
				}
			}(), coll.Root, true, v, nil})
		}
	} else {
		{
			var added_leaf_QMARK_ = (&CljsCoreBox{false})
			var new_root = Native_invoke_instance_method.X_invoke_Arity3(func() interface{} {
				if Nil_(coll.Root) {
					return CljsCoreBitmapIndexedNode_EMPTY
				} else {
					return coll.Root
				}
			}(), "Inode_assoc", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k, v, added_leaf_QMARK_})
			_, _ = added_leaf_QMARK_, new_root
			if reflect.DeepEqual(new_root, coll.Root) {
				return coll
			} else {
				return (&CljsCorePersistentHashMap{coll.Meta, func() interface{} {
					if Truth_(added_leaf_QMARK_.Val) {
						return (coll.Cnt.(float64) + float64(1))
					} else {
						return coll.Cnt
					}
				}(), new_root, coll.Has_nil_QMARK_, coll.Nil_val, nil})
			}
		}
	}
}

func (coll *CljsCorePersistentHashMap) X_contains_key_QMARK__Arity2(k interface{}) bool {
	if Nil_(k) {
		return coll.Has_nil_QMARK_
	} else {
		if Nil_(coll.Root) {
			return false
		} else {
			return !(reflect.DeepEqual(Native_invoke_instance_method.X_invoke_Arity3(coll.Root, "Inode_lookup", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k, Lookup_sentinel}), Lookup_sentinel))

		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreISeqable__() {}

func (coll *CljsCorePersistentHashMap) X_seq_Arity1() interface{} {
	if coll.Cnt.(float64) > float64(0) {
		{
			var s = func() interface{} {
				if !(Nil_(coll.Root)) {
					return Native_invoke_instance_method.X_invoke_Arity3(coll.Root, "Inode_seq", []interface{}{})
				} else {
					return nil
				}
			}()
			_ = s
			if coll.Has_nil_QMARK_ {
				return Cons.X_invoke_Arity2((&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, coll.Nil_val}, nil}), s).(*CljsCoreCons)
			} else {
				return s
			}
		}
	} else {
		return nil
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIWithMeta__() {}

func (coll *CljsCorePersistentHashMap) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCorePersistentHashMap{meta___1, coll.Cnt, coll.Root, coll.Has_nil_QMARK_, coll.Nil_val, coll.X__hash})
}

func (_ *CljsCorePersistentHashMap) CljsCoreICollection__() {}

func (coll *CljsCorePersistentHashMap) X_conj_Arity2(entry interface{}) interface{} {
	if Vector_QMARK_.Arity1IB(entry) {
		return coll.X_assoc_Arity3(entry.(CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry.(CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		{
			var ret interface{} = coll
			var es interface{} = Seq.Arity1IQ(entry)
			_, _ = ret, es
			for {
				if Nil_(es) {
					return ret
				} else {
					{
						var e = First.X_invoke_Arity1(es)
						_ = e
						if Vector_QMARK_.Arity1IB(e) {
							ret, es = ret.(CljsCoreIAssociative).X_assoc_Arity3(e.(CljsCoreIIndexed).X_nth_Arity2(float64(0)), e.(CljsCoreIIndexed).X_nth_Arity2(float64(1))), Next.Arity1IQ(es)
							continue
						} else {
							panic((&js.Error{"conj on a map takes map entries or seqables of map entries"}))
						}
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIFn__() {}

func (this *CljsCorePersistentHashMap) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (coll *CljsCorePersistentHashMap) X_invoke_Arity1(k interface{}) interface{} {
	return coll.X_lookup_Arity2(k)
}

func (coll *CljsCorePersistentHashMap) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return coll.X_lookup_Arity3(k, not_found)
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_PersistentHashMap *AFn

var CljsCorePersistentHashMap_EMPTY = (&CljsCorePersistentHashMap{nil, float64(0), nil, false, nil, float64(0)})

var CljsCorePersistentHashMap_FromArrays = func(G__7939 *AFn) *AFn {
	return Fn(G__7939, 2, func(ks interface{}, vs interface{}) interface{} {
		{
			var len = Alength_(ks)
			_ = len
			{
				var i = float64(0)
				var out interface{} = Transient.X_invoke_Arity1(CljsCorePersistentHashMap_EMPTY)
				_, _ = i, out
				for {
					if i < len {
						i, out = (i + float64(1)), out.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3(Aget_(ks, i), Aget_(vs, i))
						continue
					} else {
						return Persistent_BANG_.X_invoke_Arity1(out)
					}
				}
			}
		}
	})
}(&AFn{})

type CljsCoreTransientHashMap struct {
	Edit           bool
	Root           interface{}
	Count          interface{}
	Has_nil_QMARK_ bool
	Nil_val        interface{}
}

func (_ *CljsCoreTransientHashMap) CljsCoreITransientMap__() {}

func (tcoll *CljsCoreTransientHashMap) X_dissoc_BANG__Arity2(key interface{}) interface{} {
	return tcoll.Without_BANG_(key)
}

func (_ *CljsCoreTransientHashMap) CljsCoreITransientAssociative__() {}

func (tcoll *CljsCoreTransientHashMap) X_assoc_BANG__Arity3(key interface{}, val interface{}) interface{} {
	return tcoll.Assoc_BANG_(key, val)
}

func (_ *CljsCoreTransientHashMap) CljsCoreITransientCollection__() {}

func (tcoll *CljsCoreTransientHashMap) X_conj_BANG__Arity2(val interface{}) interface{} {
	return tcoll.Conj_BANG_(val)
}

func (tcoll *CljsCoreTransientHashMap) X_persistent_BANG__Arity1() interface{} {
	return tcoll.Persistent_BANG_()
}

func (_ *CljsCoreTransientHashMap) CljsCoreILookup__() {}

func (tcoll *CljsCoreTransientHashMap) X_lookup_Arity2(k interface{}) interface{} {
	if Nil_(k) {
		if tcoll.Has_nil_QMARK_ {
			return tcoll.Nil_val
		} else {
			return nil
		}
	} else {
		if Nil_(tcoll.Root) {
			return nil
		} else {
			return Native_invoke_instance_method.X_invoke_Arity3(tcoll.Root, "Inode_lookup", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k})
		}
	}
}

func (tcoll *CljsCoreTransientHashMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	if Nil_(k) {
		if tcoll.Has_nil_QMARK_ {
			return tcoll.Nil_val
		} else {
			return not_found
		}
	} else {
		if Nil_(tcoll.Root) {
			return not_found
		} else {
			return Native_invoke_instance_method.X_invoke_Arity3(tcoll.Root, "Inode_lookup", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k, not_found})
		}
	}
}

func (_ *CljsCoreTransientHashMap) CljsCoreICounted__() {}

func (coll *CljsCoreTransientHashMap) X_count_Arity1() float64 {
	if coll.Edit {
		return coll.Count.(float64)
	} else {
		panic((&js.Error{"count after persistent!"}))
	}
}

func (_ *CljsCoreTransientHashMap) CljsCoreObject__() {}

func (tcoll *CljsCoreTransientHashMap) Conj_BANG_(o interface{}) interface{} {
	if tcoll.Edit {
		if Value_(o).Type().Implements(reflect.TypeOf((*CljsCoreIMapEntry)(nil)).Elem()) {
			return tcoll.Assoc_BANG_(Key.X_invoke_Arity1(o), Val.X_invoke_Arity1(o))
		} else {
			{
				var es interface{} = Seq.Arity1IQ(o)
				var tcoll___1 interface{} = tcoll
				_, _ = es, tcoll___1
				for {
					{
						var temp__4220__auto__ = First.X_invoke_Arity1(es)
						_ = temp__4220__auto__
						if Truth_(temp__4220__auto__) {
							{
								var e = temp__4220__auto__
								_ = e
								es, tcoll___1 = Next.Arity1IQ(es), Native_invoke_instance_method.X_invoke_Arity3(tcoll___1, "Assoc_BANG_", []interface{}{Key.X_invoke_Arity1(e), Val.X_invoke_Arity1(e)})
								continue
							}
						} else {
							return tcoll___1
						}
					}
				}
			}
		}
	} else {
		panic((&js.Error{"conj! after persistent"}))
	}
}

func (tcoll *CljsCoreTransientHashMap) Assoc_BANG_(k interface{}, v interface{}) interface{} {
	if tcoll.Edit {
		if Nil_(k) {
			if reflect.DeepEqual(tcoll.Nil_val, v) {
			} else {
				tcoll.Nil_val = v

			}
			if tcoll.Has_nil_QMARK_ {
			} else {
				tcoll.Count = (tcoll.Count.(float64) + float64(1))

				tcoll.Has_nil_QMARK_ = true

			}
			return tcoll
		} else {
			{
				var added_leaf_QMARK_ = (&CljsCoreBox{false})
				var node = Native_invoke_instance_method.X_invoke_Arity3(func() interface{} {
					if Nil_(tcoll.Root) {
						return CljsCoreBitmapIndexedNode_EMPTY
					} else {
						return tcoll.Root
					}
				}(), "Inode_assoc_BANG_", []interface{}{tcoll.Edit, float64(0), Hash.X_invoke_Arity1(k), k, v, added_leaf_QMARK_})
				_, _ = added_leaf_QMARK_, node
				if reflect.DeepEqual(node, tcoll.Root) {
				} else {
					tcoll.Root = node

				}
				if Truth_(added_leaf_QMARK_.Val) {
					tcoll.Count = (tcoll.Count.(float64) + float64(1))

				} else {
				}
				return tcoll
			}
		}
	} else {
		panic((&js.Error{"assoc! after persistent!"}))
	}
}

func (tcoll *CljsCoreTransientHashMap) Without_BANG_(k interface{}) interface{} {
	if tcoll.Edit {
		if Nil_(k) {
			if tcoll.Has_nil_QMARK_ {
				tcoll.Has_nil_QMARK_ = false

				tcoll.Nil_val = nil

				tcoll.Count = (tcoll.Count.(float64) - float64(1))

				return tcoll
			} else {
				return tcoll
			}
		} else {
			if Nil_(tcoll.Root) {
				return tcoll
			} else {
				{
					var removed_leaf_QMARK_ = (&CljsCoreBox{false})
					var node = Native_invoke_instance_method.X_invoke_Arity3(tcoll.Root, "Inode_without_BANG_", []interface{}{tcoll.Edit, float64(0), Hash.X_invoke_Arity1(k), k, removed_leaf_QMARK_})
					_, _ = removed_leaf_QMARK_, node
					if reflect.DeepEqual(node, tcoll.Root) {
					} else {
						tcoll.Root = node

					}
					if Truth_(removed_leaf_QMARK_.Val) {
						tcoll.Count = (tcoll.Count.(float64) - float64(1))

					} else {
					}
					return tcoll
				}
			}
		}
	} else {
		panic((&js.Error{"dissoc! after persistent!"}))
	}
}

func (tcoll *CljsCoreTransientHashMap) Persistent_BANG_() interface{} {
	if tcoll.Edit {
		tcoll.Edit = Truth_(nil)

		return (&CljsCorePersistentHashMap{nil, tcoll.Count, tcoll.Root, tcoll.Has_nil_QMARK_, tcoll.Nil_val, nil})
	} else {
		panic((&js.Error{"persistent! called twice"}))
	}
}

var X__GT_TransientHashMap *AFn

var Tree_map_seq_push *AFn

type CljsCorePersistentTreeMapSeq struct {
	Meta             interface{}
	Stack            interface{}
	Ascending_QMARK_ bool
	Cnt              interface{}
	X__hash          interface{}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreObject__() {}

func (coll *CljsCorePersistentTreeMapSeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCorePersistentTreeMapSeq) String() string {
	return this.ToString()
}

func (this *CljsCorePersistentTreeMapSeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIMeta__() {}

func (coll *CljsCorePersistentTreeMapSeq) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreICounted__() {}

func (coll *CljsCorePersistentTreeMapSeq) X_count_Arity1() float64 {
	if coll.Cnt.(float64) < float64(0) {
		return (Count.X_invoke_Arity1(Next.Arity1IQ(coll)).(float64) + float64(1))
	} else {
		return coll.Cnt.(float64)
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIHash__() {}

func (coll *CljsCorePersistentTreeMapSeq) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentTreeMapSeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCorePersistentTreeMapSeq) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.Meta)
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIReduce__() {}

func (coll *CljsCorePersistentTreeMapSeq) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCorePersistentTreeMapSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreISeq__() {}

func (this *CljsCorePersistentTreeMapSeq) X_first_Arity1() interface{} {
	return Peek.X_invoke_Arity1(this.Stack)
}

func (this *CljsCorePersistentTreeMapSeq) X_rest_Arity1() interface{} {
	{
		var t = First.X_invoke_Arity1(this.Stack)
		var next_stack = Tree_map_seq_push.X_invoke_Arity3(func() interface{} {
			if this.Ascending_QMARK_ {
				return Native_get_instance_field.X_invoke_Arity2(t, "Right")
			} else {
				return Native_get_instance_field.X_invoke_Arity2(t, "Left")
			}
		}(), Next.Arity1IQ(this.Stack), this.Ascending_QMARK_)
		_, _ = t, next_stack
		if !(Nil_(next_stack)) {
			return (&CljsCorePersistentTreeMapSeq{nil, next_stack, this.Ascending_QMARK_, (this.Cnt.(float64) - float64(1)), nil})
		} else {
			return CljsCoreIEmptyList(CljsCoreList_EMPTY)
		}
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreISeqable__() {}

func (this *CljsCorePersistentTreeMapSeq) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreISequential__() {}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIWithMeta__() {}

func (coll *CljsCorePersistentTreeMapSeq) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCorePersistentTreeMapSeq{meta___1, coll.Stack, coll.Ascending_QMARK_, coll.Cnt, coll.X__hash})
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreICollection__() {}

func (coll *CljsCorePersistentTreeMapSeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

var X__GT_PersistentTreeMapSeq *AFn

var Create_tree_map_seq *AFn

var Balance_left *AFn

var Balance_right *AFn

var Balance_left_del *AFn

var Balance_right_del *AFn

var Tree_map_kv_reduce *AFn

type CljsCoreBlackNode struct {
	Key     interface{}
	Val     interface{}
	Left    interface{}
	Right   interface{}
	X__hash interface{}
}

func (_ *CljsCoreBlackNode) CljsCoreObject__() {}

func (node *CljsCoreBlackNode) Add_right(ins interface{}) interface{} {
	return Native_invoke_instance_method.X_invoke_Arity3(ins, "Balance_right", []interface{}{node})
}

func (node *CljsCoreBlackNode) Redden() interface{} {
	return (&CljsCoreRedNode{node.Key, node.Val, node.Left, node.Right, nil})
}

func (node *CljsCoreBlackNode) Blacken() interface{} {
	return node
}

func (node *CljsCoreBlackNode) Add_left(ins interface{}) interface{} {
	return Native_invoke_instance_method.X_invoke_Arity3(ins, "Balance_left", []interface{}{node})
}

func (node *CljsCoreBlackNode) Replace(key___1 interface{}, val___1 interface{}, left___1 interface{}, right___1 interface{}) interface{} {
	return (&CljsCoreBlackNode{key___1, val___1, left___1, right___1, nil})
}

func (node *CljsCoreBlackNode) Balance_left(parent interface{}) interface{} {
	return (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), node, Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil})
}

func (node *CljsCoreBlackNode) Balance_right(parent interface{}) interface{} {
	return (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(parent, "Left"), node, nil})
}

func (node *CljsCoreBlackNode) Remove_left(del interface{}) interface{} {
	return Balance_left_del.X_invoke_Arity4(node.Key, node.Val, del, node.Right)
}

func (node *CljsCoreBlackNode) Kv_reduce(f interface{}, init interface{}) interface{} {
	return Tree_map_kv_reduce.X_invoke_Arity3(node, f, init)
}

func (node *CljsCoreBlackNode) Remove_right(del interface{}) interface{} {
	return Balance_right_del.X_invoke_Arity4(node.Key, node.Val, node.Left, del)
}

func (_ *CljsCoreBlackNode) CljsCoreILookup__() {}

func (node *CljsCoreBlackNode) X_lookup_Arity2(k interface{}) interface{} {
	return node.X_nth_Arity3(k, nil)
}

func (node *CljsCoreBlackNode) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	return node.X_nth_Arity3(k, not_found)
}

func (_ *CljsCoreBlackNode) CljsCoreIIndexed__() {}

func (node *CljsCoreBlackNode) X_nth_Arity2(n interface{}) interface{} {
	if n.(float64) == float64(0) {
		return node.Key
	} else {
		if n.(float64) == float64(1) {
			return node.Val
		} else {
			return nil

		}
	}
}

func (node *CljsCoreBlackNode) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	if n.(float64) == float64(0) {
		return node.Key
	} else {
		if n.(float64) == float64(1) {
			return node.Val
		} else {
			return not_found

		}
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIVector__() {}

func (node *CljsCoreBlackNode) X_assoc_n_Arity3(n interface{}, v interface{}) interface{} {
	return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{node.Key, node.Val}, nil}).X_assoc_n_Arity3(n, v)
}

func (_ *CljsCoreBlackNode) CljsCoreIMeta__() {}

func (node *CljsCoreBlackNode) X_meta_Arity1() interface{} {
	return nil
}

func (_ *CljsCoreBlackNode) CljsCoreICounted__() {}

func (node *CljsCoreBlackNode) X_count_Arity1() float64 {
	return float64(2)
}

func (_ *CljsCoreBlackNode) CljsCoreIMapEntry__() {}

func (node *CljsCoreBlackNode) X_key_Arity1() interface{} {
	return node.Key
}

func (node *CljsCoreBlackNode) X_val_Arity1() interface{} {
	return node.Val
}

func (_ *CljsCoreBlackNode) CljsCoreIStack__() {}

func (node *CljsCoreBlackNode) X_peek_Arity1() interface{} {
	return node.Val
}

func (node *CljsCoreBlackNode) X_pop_Arity1() interface{} {
	return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{node.Key}, nil})
}

func (_ *CljsCoreBlackNode) CljsCoreIHash__() {}

func (coll *CljsCoreBlackNode) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIEquiv__() {}

func (coll *CljsCoreBlackNode) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreBlackNode) CljsCoreIEmptyableCollection__() {}

func (node *CljsCoreBlackNode) X_empty_Arity1() interface{} {
	return CljsCorePersistentVector_EMPTY
}

func (_ *CljsCoreBlackNode) CljsCoreIReduce__() {}

func (node *CljsCoreBlackNode) X_reduce_Arity2(f interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity2(node, f)
}

func (node *CljsCoreBlackNode) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity3(node, f, start)
}

func (_ *CljsCoreBlackNode) CljsCoreIAssociative__() {}

func (node *CljsCoreBlackNode) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	return Assoc.X_invoke_Arity3((&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{node.Key, node.Val}, nil}), k, v)
}

func (_ *CljsCoreBlackNode) CljsCoreISeqable__() {}

func (node *CljsCoreBlackNode) X_seq_Arity1() interface{} {
	return CljsCoreList_EMPTY.X_conj_Arity2(node.Val).(CljsCoreICollection).X_conj_Arity2(node.Key)
}

func (_ *CljsCoreBlackNode) CljsCoreISequential__() {}

func (_ *CljsCoreBlackNode) CljsCoreIWithMeta__() {}

func (node *CljsCoreBlackNode) X_with_meta_Arity2(meta interface{}) interface{} {
	return With_meta.X_invoke_Arity2((&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{node.Key, node.Val}, nil}), meta)
}

func (_ *CljsCoreBlackNode) CljsCoreICollection__() {}

func (node *CljsCoreBlackNode) X_conj_Arity2(o interface{}) interface{} {
	return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{node.Key, node.Val, o}, nil})
}

func (_ *CljsCoreBlackNode) CljsCoreIFn__() {}

func (this *CljsCoreBlackNode) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (node *CljsCoreBlackNode) X_invoke_Arity1(k interface{}) interface{} {
	return node.X_lookup_Arity2(k)
}

func (node *CljsCoreBlackNode) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return node.X_lookup_Arity3(k, not_found)
}

func (this *CljsCoreBlackNode) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_BlackNode *AFn

type CljsCoreRedNode struct {
	Key     interface{}
	Val     interface{}
	Left    interface{}
	Right   interface{}
	X__hash interface{}
}

func (_ *CljsCoreRedNode) CljsCoreObject__() {}

func (node *CljsCoreRedNode) Add_right(ins interface{}) interface{} {
	return (&CljsCoreRedNode{node.Key, node.Val, node.Left, ins, nil})
}

func (node *CljsCoreRedNode) Redden() interface{} {
	panic((&js.Error{"red-black tree invariant violation"}))
}

func (node *CljsCoreRedNode) Blacken() interface{} {
	return (&CljsCoreBlackNode{node.Key, node.Val, node.Left, node.Right, nil})
}

func (node *CljsCoreRedNode) Add_left(ins interface{}) interface{} {
	return (&CljsCoreRedNode{node.Key, node.Val, ins, node.Right, nil})
}

func (node *CljsCoreRedNode) Replace(key___1 interface{}, val___1 interface{}, left___1 interface{}, right___1 interface{}) interface{} {
	return (&CljsCoreRedNode{key___1, val___1, left___1, right___1, nil})
}

func (node *CljsCoreRedNode) Balance_left(parent interface{}) interface{} {
	if Value_(node.Left).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
		return (&CljsCoreRedNode{node.Key, node.Val, Native_invoke_instance_method.X_invoke_Arity3(node.Left, "Blacken", []interface{}{}), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), node.Right, Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil}), nil})
	} else {
		if Value_(node.Right).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
			return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(node.Right, "Key"), Native_get_instance_field.X_invoke_Arity2(node.Right, "Val"), (&CljsCoreBlackNode{node.Key, node.Val, node.Left, Native_get_instance_field.X_invoke_Arity2(node.Right, "Left"), nil}), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(node.Right, "Right"), Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil}), nil})
		} else {
			return (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), node, Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil})

		}
	}
}

func (node *CljsCoreRedNode) Balance_right(parent interface{}) interface{} {
	if Value_(node.Right).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
		return (&CljsCoreRedNode{node.Key, node.Val, (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(parent, "Left"), node.Left, nil}), Native_invoke_instance_method.X_invoke_Arity3(node.Right, "Blacken", []interface{}{}), nil})
	} else {
		if Value_(node.Left).Type().AssignableTo(reflect.TypeOf((**CljsCoreRedNode)(nil)).Elem()) {
			return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(node.Left, "Key"), Native_get_instance_field.X_invoke_Arity2(node.Left, "Val"), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(parent, "Left"), Native_get_instance_field.X_invoke_Arity2(node.Left, "Left"), nil}), (&CljsCoreBlackNode{node.Key, node.Val, Native_get_instance_field.X_invoke_Arity2(node.Left, "Right"), node.Right, nil}), nil})
		} else {
			return (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(parent, "Left"), node, nil})

		}
	}
}

func (node *CljsCoreRedNode) Remove_left(del interface{}) interface{} {
	return (&CljsCoreRedNode{node.Key, node.Val, del, node.Right, nil})
}

func (node *CljsCoreRedNode) Kv_reduce(f interface{}, init interface{}) interface{} {
	return Tree_map_kv_reduce.X_invoke_Arity3(node, f, init)
}

func (node *CljsCoreRedNode) Remove_right(del interface{}) interface{} {
	return (&CljsCoreRedNode{node.Key, node.Val, node.Left, del, nil})
}

func (_ *CljsCoreRedNode) CljsCoreILookup__() {}

func (node *CljsCoreRedNode) X_lookup_Arity2(k interface{}) interface{} {
	return node.X_nth_Arity3(k, nil)
}

func (node *CljsCoreRedNode) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	return node.X_nth_Arity3(k, not_found)
}

func (_ *CljsCoreRedNode) CljsCoreIIndexed__() {}

func (node *CljsCoreRedNode) X_nth_Arity2(n interface{}) interface{} {
	if n.(float64) == float64(0) {
		return node.Key
	} else {
		if n.(float64) == float64(1) {
			return node.Val
		} else {
			return nil

		}
	}
}

func (node *CljsCoreRedNode) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	if n.(float64) == float64(0) {
		return node.Key
	} else {
		if n.(float64) == float64(1) {
			return node.Val
		} else {
			return not_found

		}
	}
}

func (_ *CljsCoreRedNode) CljsCoreIVector__() {}

func (node *CljsCoreRedNode) X_assoc_n_Arity3(n interface{}, v interface{}) interface{} {
	return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{node.Key, node.Val}, nil}).X_assoc_n_Arity3(n, v)
}

func (_ *CljsCoreRedNode) CljsCoreIMeta__() {}

func (node *CljsCoreRedNode) X_meta_Arity1() interface{} {
	return nil
}

func (_ *CljsCoreRedNode) CljsCoreICounted__() {}

func (node *CljsCoreRedNode) X_count_Arity1() float64 {
	return float64(2)
}

func (_ *CljsCoreRedNode) CljsCoreIMapEntry__() {}

func (node *CljsCoreRedNode) X_key_Arity1() interface{} {
	return node.Key
}

func (node *CljsCoreRedNode) X_val_Arity1() interface{} {
	return node.Val
}

func (_ *CljsCoreRedNode) CljsCoreIStack__() {}

func (node *CljsCoreRedNode) X_peek_Arity1() interface{} {
	return node.Val
}

func (node *CljsCoreRedNode) X_pop_Arity1() interface{} {
	return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{node.Key}, nil})
}

func (_ *CljsCoreRedNode) CljsCoreIHash__() {}

func (coll *CljsCoreRedNode) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreRedNode) CljsCoreIEquiv__() {}

func (coll *CljsCoreRedNode) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreRedNode) CljsCoreIEmptyableCollection__() {}

func (node *CljsCoreRedNode) X_empty_Arity1() interface{} {
	return CljsCorePersistentVector_EMPTY
}

func (_ *CljsCoreRedNode) CljsCoreIReduce__() {}

func (node *CljsCoreRedNode) X_reduce_Arity2(f interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity2(node, f)
}

func (node *CljsCoreRedNode) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity3(node, f, start)
}

func (_ *CljsCoreRedNode) CljsCoreIAssociative__() {}

func (node *CljsCoreRedNode) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	return Assoc.X_invoke_Arity3((&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{node.Key, node.Val}, nil}), k, v)
}

func (_ *CljsCoreRedNode) CljsCoreISeqable__() {}

func (node *CljsCoreRedNode) X_seq_Arity1() interface{} {
	return CljsCoreList_EMPTY.X_conj_Arity2(node.Val).(CljsCoreICollection).X_conj_Arity2(node.Key)
}

func (_ *CljsCoreRedNode) CljsCoreISequential__() {}

func (_ *CljsCoreRedNode) CljsCoreIWithMeta__() {}

func (node *CljsCoreRedNode) X_with_meta_Arity2(meta interface{}) interface{} {
	return With_meta.X_invoke_Arity2((&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{node.Key, node.Val}, nil}), meta)
}

func (_ *CljsCoreRedNode) CljsCoreICollection__() {}

func (node *CljsCoreRedNode) X_conj_Arity2(o interface{}) interface{} {
	return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{node.Key, node.Val, o}, nil})
}

func (_ *CljsCoreRedNode) CljsCoreIFn__() {}

func (this *CljsCoreRedNode) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (node *CljsCoreRedNode) X_invoke_Arity1(k interface{}) interface{} {
	return node.X_lookup_Arity2(k)
}

func (node *CljsCoreRedNode) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return node.X_lookup_Arity3(k, not_found)
}

func (this *CljsCoreRedNode) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_RedNode *AFn

var Tree_map_add *AFn

var Tree_map_append *AFn

var Tree_map_remove *AFn

var Tree_map_replace *AFn

type CljsCorePersistentTreeMap struct {
	Comp    interface{}
	Tree    interface{}
	Cnt     interface{}
	Meta    interface{}
	X__hash interface{}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreObject__() {}

func (coll *CljsCorePersistentTreeMap) ForEach(f interface{}) interface{} {
	{
		var seq__6354 interface{} = Seq.Arity1IQ(coll)
		var chunk__6355 interface{} = nil
		var count__6356 = float64(0)
		var i__6357 = float64(0)
		_, _, _, _ = seq__6354, chunk__6355, count__6356, i__6357
		for {
			if i__6357 < count__6356 {
				{
					var vec__6358 = chunk__6355.(CljsCoreIIndexed).X_nth_Arity2(i__6357)
					var k = Nth.X_invoke_Arity3(vec__6358, float64(0), nil)
					var v = Nth.X_invoke_Arity3(vec__6358, float64(1), nil)
					_, _, _ = vec__6358, k, v
					{
						var G__6359_7940 = v
						var G__6360_7941 = k
						_, _ = G__6359_7940, G__6360_7941
						f.(CljsCoreIFn).X_invoke_Arity2(G__6359_7940, G__6360_7941)
					}
					seq__6354, chunk__6355, count__6356, i__6357 = seq__6354, chunk__6355, count__6356, (i__6357 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__6354)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__6354___1 = temp__4222__auto__
							_ = seq__6354___1
							if Chunked_seq_QMARK_.Arity1IB(seq__6354___1) {
								{
									var c__954__auto__ = Chunk_first.X_invoke_Arity1(seq__6354___1)
									_ = c__954__auto__
									seq__6354, chunk__6355, count__6356, i__6357 = Chunk_rest.X_invoke_Arity1(seq__6354___1), c__954__auto__, Count.X_invoke_Arity1(c__954__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__6361 = First.X_invoke_Arity1(seq__6354___1)
									var k = Nth.X_invoke_Arity3(vec__6361, float64(0), nil)
									var v = Nth.X_invoke_Arity3(vec__6361, float64(1), nil)
									_, _, _ = vec__6361, k, v
									{
										var G__6362_7942 = v
										var G__6363_7943 = k
										_, _ = G__6362_7942, G__6363_7943
										f.(CljsCoreIFn).X_invoke_Arity2(G__6362_7942, G__6363_7943)
									}
									seq__6354, chunk__6355, count__6356, i__6357 = Next.Arity1IQ(seq__6354___1), nil, float64(0), float64(0)
									continue
								}
							}
						}
					} else {
						return nil
					}
				}
			}
		}
	}
}

func (coll *CljsCorePersistentTreeMap) Get(k interface{}) interface{} {
	return coll.X_lookup_Arity2(k)
}

func (coll *CljsCorePersistentTreeMap) Entries() interface{} {
	return Entries_iterator.X_invoke_Arity1(Seq.Arity1IQ(coll)).(*CljsCoreEntriesIterator)
}

func (coll *CljsCorePersistentTreeMap) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCorePersistentTreeMap) String() string {
	return this.ToString()
}

func (coll *CljsCorePersistentTreeMap) Keys() interface{} {
	return Iterator.X_invoke_Arity1(Native_invoke_func.X_invoke_Arity2((*CljsCorePersistentTreeMap).Keys, []interface{}{coll}).(*CljsCoreIterator)).(*CljsCoreIterator)
}

func (coll *CljsCorePersistentTreeMap) Values() interface{} {
	return Iterator.X_invoke_Arity1(Vals.X_invoke_Arity1(coll)).(*CljsCoreIterator)
}

func (this *CljsCorePersistentTreeMap) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (coll *CljsCorePersistentTreeMap) Entry_at(k interface{}) interface{} {
	{
		var t interface{} = coll.Tree
		_ = t
		for {
			if !(Nil_(t)) {
				{
					var c = func() interface{} {
						var G__6367 = k
						var G__6368 = Native_get_instance_field.X_invoke_Arity2(t, "Key")
						_, _ = G__6367, G__6368
						return coll.Comp.(CljsCoreIFn).X_invoke_Arity2(G__6367, G__6368)
					}()
					_ = c
					if c.(float64) == float64(0) {
						return t
					} else {
						if c.(float64) < float64(0) {
							t = Native_get_instance_field.X_invoke_Arity2(t, "Left")
							continue
						} else {
							t = Native_get_instance_field.X_invoke_Arity2(t, "Right")
							continue

						}
					}
				}
			} else {
				return nil
			}
		}
	}
}

func (coll *CljsCorePersistentTreeMap) Has(k interface{}) interface{} {
	return Contains_QMARK_.Arity2IIB(coll, k)
}

func (_ *CljsCorePersistentTreeMap) CljsCoreILookup__() {}

func (coll *CljsCorePersistentTreeMap) X_lookup_Arity2(k interface{}) interface{} {
	return coll.X_lookup_Arity3(k, nil)
}

func (coll *CljsCorePersistentTreeMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var n = coll.Entry_at(k)
		_ = n
		if !(Nil_(n)) {
			return Native_get_instance_field.X_invoke_Arity2(n, "Val")
		} else {
			return not_found
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIKVReduce__() {}

func (coll *CljsCorePersistentTreeMap) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	if !(Nil_(coll.Tree)) {
		return Tree_map_kv_reduce.X_invoke_Arity3(coll.Tree, f, init)
	} else {
		return init
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIMeta__() {}

func (coll *CljsCorePersistentTreeMap) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCorePersistentTreeMap) CljsCoreICloneable__() {}

func (___ *CljsCorePersistentTreeMap) X_clone_Arity1() interface{} {
	return (&CljsCorePersistentTreeMap{___.Comp, ___.Tree, ___.Cnt, ___.Meta, ___.X__hash})
}

func (_ *CljsCorePersistentTreeMap) CljsCoreICounted__() {}

func (coll *CljsCorePersistentTreeMap) X_count_Arity1() float64 {
	return coll.Cnt.(float64)
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIReversible__() {}

func (coll *CljsCorePersistentTreeMap) X_rseq_Arity1() interface{} {
	if coll.Cnt.(float64) > float64(0) {
		return Create_tree_map_seq.X_invoke_Arity3(coll.Tree, false, coll.Cnt).(*CljsCorePersistentTreeMapSeq)
	} else {
		return nil
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIHash__() {}

func (coll *CljsCorePersistentTreeMap) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentTreeMap) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_map.X_invoke_Arity2(coll, other))
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCorePersistentTreeMap) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCorePersistentTreeMap_EMPTY, coll.Meta)
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIMap__() {}

func (coll *CljsCorePersistentTreeMap) X_dissoc_Arity2(k interface{}) interface{} {
	{
		var found = []interface{}{nil}
		var t = Tree_map_remove.X_invoke_Arity4(coll.Comp, coll.Tree, k, found)
		_, _ = found, t
		if Nil_(t) {
			if Nil_(Nth.X_invoke_Arity2(found, float64(0))) {
				return coll
			} else {
				return (&CljsCorePersistentTreeMap{coll.Comp, nil, float64(0), coll.Meta, nil})
			}
		} else {
			return (&CljsCorePersistentTreeMap{coll.Comp, Native_invoke_instance_method.X_invoke_Arity3(t, "Blacken", []interface{}{}), (coll.Cnt.(float64) - float64(1)), coll.Meta, nil})
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIAssociative__() {}

func (coll *CljsCorePersistentTreeMap) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	{
		var found = []interface{}{nil}
		var t = Tree_map_add.X_invoke_Arity5(coll.Comp, coll.Tree, k, v, found)
		_, _ = found, t
		if Nil_(t) {
			{
				var found_node = Nth.X_invoke_Arity2(found, float64(0))
				_ = found_node
				if X_EQ_.Arity2IIB(v, Native_get_instance_field.X_invoke_Arity2(found_node, "Val")) {
					return coll
				} else {
					return (&CljsCorePersistentTreeMap{coll.Comp, Tree_map_replace.X_invoke_Arity4(coll.Comp, coll.Tree, k, v), coll.Cnt, coll.Meta, nil})
				}
			}
		} else {
			return (&CljsCorePersistentTreeMap{coll.Comp, Native_invoke_instance_method.X_invoke_Arity3(t, "Blacken", []interface{}{}), (coll.Cnt.(float64) + float64(1)), coll.Meta, nil})
		}
	}
}

func (coll *CljsCorePersistentTreeMap) X_contains_key_QMARK__Arity2(k interface{}) bool {
	return !(Nil_(coll.Entry_at(k)))
}

func (_ *CljsCorePersistentTreeMap) CljsCoreISeqable__() {}

func (coll *CljsCorePersistentTreeMap) X_seq_Arity1() interface{} {
	if coll.Cnt.(float64) > float64(0) {
		return Create_tree_map_seq.X_invoke_Arity3(coll.Tree, true, coll.Cnt).(*CljsCorePersistentTreeMapSeq)
	} else {
		return nil
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIWithMeta__() {}

func (coll *CljsCorePersistentTreeMap) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCorePersistentTreeMap{coll.Comp, coll.Tree, coll.Cnt, meta___1, coll.X__hash})
}

func (_ *CljsCorePersistentTreeMap) CljsCoreICollection__() {}

func (coll *CljsCorePersistentTreeMap) X_conj_Arity2(entry interface{}) interface{} {
	if Vector_QMARK_.Arity1IB(entry) {
		return coll.X_assoc_Arity3(entry.(CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry.(CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		{
			var ret interface{} = coll
			var es interface{} = Seq.Arity1IQ(entry)
			_, _ = ret, es
			for {
				if Nil_(es) {
					return ret
				} else {
					{
						var e = First.X_invoke_Arity1(es)
						_ = e
						if Vector_QMARK_.Arity1IB(e) {
							ret, es = ret.(CljsCoreIAssociative).X_assoc_Arity3(e.(CljsCoreIIndexed).X_nth_Arity2(float64(0)), e.(CljsCoreIIndexed).X_nth_Arity2(float64(1))), Next.Arity1IQ(es)
							continue
						} else {
							panic((&js.Error{"conj on a map takes map entries or seqables of map entries"}))
						}
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIFn__() {}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (coll *CljsCorePersistentTreeMap) X_invoke_Arity1(k interface{}) interface{} {
	return coll.X_lookup_Arity2(k)
}

func (coll *CljsCorePersistentTreeMap) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return coll.X_lookup_Arity3(k, not_found)
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

func (_ *CljsCorePersistentTreeMap) CljsCoreISorted__() {}

func (coll *CljsCorePersistentTreeMap) X_sorted_seq_Arity2(ascending_QMARK_ interface{}) interface{} {
	if coll.Cnt.(float64) > float64(0) {
		return Create_tree_map_seq.X_invoke_Arity3(coll.Tree, ascending_QMARK_, coll.Cnt).(*CljsCorePersistentTreeMapSeq)
	} else {
		return nil
	}
}

func (coll *CljsCorePersistentTreeMap) X_sorted_seq_from_Arity3(k interface{}, ascending_QMARK_ interface{}) interface{} {
	if coll.Cnt.(float64) > float64(0) {
		{
			var stack interface{} = nil
			var t interface{} = coll.Tree
			_, _ = stack, t
			for {
				if !(Nil_(t)) {
					{
						var c = func() interface{} {
							var G__6374 = k
							var G__6375 = Native_get_instance_field.X_invoke_Arity2(t, "Key")
							_, _ = G__6374, G__6375
							return coll.Comp.(CljsCoreIFn).X_invoke_Arity2(G__6374, G__6375)
						}()
						_ = c
						if c.(float64) == float64(0) {
							return (&CljsCorePersistentTreeMapSeq{nil, Conj.X_invoke_Arity2(stack, t), ascending_QMARK_.(bool), float64(-1), nil})
						} else {
							if Truth_(ascending_QMARK_) {
								if c.(float64) < float64(0) {
									stack, t = Conj.X_invoke_Arity2(stack, t), Native_get_instance_field.X_invoke_Arity2(t, "Left")
									continue
								} else {
									stack, t = stack, Native_get_instance_field.X_invoke_Arity2(t, "Right")
									continue
								}
							} else {
								if c.(float64) > float64(0) {
									stack, t = Conj.X_invoke_Arity2(stack, t), Native_get_instance_field.X_invoke_Arity2(t, "Right")
									continue
								} else {
									stack, t = stack, Native_get_instance_field.X_invoke_Arity2(t, "Left")
									continue
								}

							}
						}
					}
				} else {
					if Nil_(stack) {
						return nil
					} else {
						return (&CljsCorePersistentTreeMapSeq{nil, stack, ascending_QMARK_.(bool), float64(-1), nil})
					}
				}
			}
		}
	} else {
		return nil
	}
}

func (coll *CljsCorePersistentTreeMap) X_entry_key_Arity2(entry interface{}) interface{} {
	return Key.X_invoke_Arity1(entry)
}

func (coll *CljsCorePersistentTreeMap) X_comparator_Arity1() interface{} {
	return coll.Comp
}

var X__GT_PersistentTreeMap *AFn

var CljsCorePersistentTreeMap_EMPTY = (&CljsCorePersistentTreeMap{Compare, nil, float64(0), nil, float64(0)})

// keyval => key val
// Returns a new hash map with supplied mappings.
// @param {...*} var_args
var Hash_map *AFn

// keyval => key val
// Returns a new array map with supplied mappings.
// @param {...*} var_args
var Array_map *AFn

// keyval => key val
// Returns a new sorted map with supplied mappings.
// @param {...*} var_args
var Sorted_map *AFn

// keyval => key val
// Returns a new sorted map with supplied mappings, using the supplied comparator.
// @param {...*} var_args
var Sorted_map_by *AFn

type CljsCoreKeySeq struct {
	Mseq   interface{}
	X_meta interface{}
}

func (_ *CljsCoreKeySeq) CljsCoreObject__() {}

func (coll *CljsCoreKeySeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreKeySeq) String() string {
	return this.ToString()
}

func (this *CljsCoreKeySeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreKeySeq) CljsCoreIMeta__() {}

func (coll *CljsCoreKeySeq) X_meta_Arity1() interface{} {
	return coll.X_meta
}

func (_ *CljsCoreKeySeq) CljsCoreINext__() {}

func (coll *CljsCoreKeySeq) X_next_Arity1() interface{} {
	{
		var nseq = func() interface{} {
			if Value_(coll.Mseq).Type().Implements(reflect.TypeOf((*CljsCoreINext)(nil)).Elem()) {
				return coll.Mseq.(CljsCoreINext).X_next_Arity1()
			} else {
				return Next.Arity1IQ(coll.Mseq)
			}
		}()
		_ = nseq
		if Nil_(nseq) {
			return nil
		} else {
			return (&CljsCoreKeySeq{nseq, coll.X_meta})
		}
	}
}

func (_ *CljsCoreKeySeq) CljsCoreIHash__() {}

func (coll *CljsCoreKeySeq) X_hash_Arity1() interface{} {
	return Hash_ordered_coll.Arity1IF(coll)
}

func (_ *CljsCoreKeySeq) CljsCoreIEquiv__() {}

func (coll *CljsCoreKeySeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreKeySeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreKeySeq) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.X_meta)
}

func (_ *CljsCoreKeySeq) CljsCoreIReduce__() {}

func (coll *CljsCoreKeySeq) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCoreKeySeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCoreKeySeq) CljsCoreISeq__() {}

func (coll *CljsCoreKeySeq) X_first_Arity1() interface{} {
	{
		var me = coll.Mseq.(CljsCoreISeq).X_first_Arity1()
		_ = me
		return me.(CljsCoreIMapEntry).X_key_Arity1()
	}
}

func (coll *CljsCoreKeySeq) X_rest_Arity1() interface{} {
	{
		var nseq = func() interface{} {
			if Value_(coll.Mseq).Type().Implements(reflect.TypeOf((*CljsCoreINext)(nil)).Elem()) {
				return coll.Mseq.(CljsCoreINext).X_next_Arity1()
			} else {
				return Next.Arity1IQ(coll.Mseq)
			}
		}()
		_ = nseq
		if !(Nil_(nseq)) {
			return (&CljsCoreKeySeq{nseq, coll.X_meta})
		} else {
			return CljsCoreIEmptyList(CljsCoreList_EMPTY)
		}
	}
}

func (_ *CljsCoreKeySeq) CljsCoreISeqable__() {}

func (coll *CljsCoreKeySeq) X_seq_Arity1() interface{} {
	return coll
}

func (_ *CljsCoreKeySeq) CljsCoreISequential__() {}

func (_ *CljsCoreKeySeq) CljsCoreIWithMeta__() {}

func (coll *CljsCoreKeySeq) X_with_meta_Arity2(new_meta interface{}) interface{} {
	return (&CljsCoreKeySeq{coll.Mseq, new_meta})
}

func (_ *CljsCoreKeySeq) CljsCoreICollection__() {}

func (coll *CljsCoreKeySeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

var X__GT_KeySeq *AFn

// Returns a sequence of the map's keys.
var Keys *AFn

// Returns the key of the map entry.
var Key *AFn

type CljsCoreValSeq struct {
	Mseq   interface{}
	X_meta interface{}
}

func (_ *CljsCoreValSeq) CljsCoreObject__() {}

func (coll *CljsCoreValSeq) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreValSeq) String() string {
	return this.ToString()
}

func (this *CljsCoreValSeq) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreValSeq) CljsCoreIMeta__() {}

func (coll *CljsCoreValSeq) X_meta_Arity1() interface{} {
	return coll.X_meta
}

func (_ *CljsCoreValSeq) CljsCoreINext__() {}

func (coll *CljsCoreValSeq) X_next_Arity1() interface{} {
	{
		var nseq = func() interface{} {
			if Value_(coll.Mseq).Type().Implements(reflect.TypeOf((*CljsCoreINext)(nil)).Elem()) {
				return coll.Mseq.(CljsCoreINext).X_next_Arity1()
			} else {
				return Next.Arity1IQ(coll.Mseq)
			}
		}()
		_ = nseq
		if Nil_(nseq) {
			return nil
		} else {
			return (&CljsCoreValSeq{nseq, coll.X_meta})
		}
	}
}

func (_ *CljsCoreValSeq) CljsCoreIHash__() {}

func (coll *CljsCoreValSeq) X_hash_Arity1() interface{} {
	return Hash_ordered_coll.Arity1IF(coll)
}

func (_ *CljsCoreValSeq) CljsCoreIEquiv__() {}

func (coll *CljsCoreValSeq) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreValSeq) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreValSeq) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, coll.X_meta)
}

func (_ *CljsCoreValSeq) CljsCoreIReduce__() {}

func (coll *CljsCoreValSeq) X_reduce_Arity2(f interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity2(f, coll)
}

func (coll *CljsCoreValSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Seq_reduce.X_invoke_Arity3(f, start, coll)
}

func (_ *CljsCoreValSeq) CljsCoreISeq__() {}

func (coll *CljsCoreValSeq) X_first_Arity1() interface{} {
	{
		var me = coll.Mseq.(CljsCoreISeq).X_first_Arity1()
		_ = me
		return me.(CljsCoreIMapEntry).X_val_Arity1()
	}
}

func (coll *CljsCoreValSeq) X_rest_Arity1() interface{} {
	{
		var nseq = func() interface{} {
			if Value_(coll.Mseq).Type().Implements(reflect.TypeOf((*CljsCoreINext)(nil)).Elem()) {
				return coll.Mseq.(CljsCoreINext).X_next_Arity1()
			} else {
				return Next.Arity1IQ(coll.Mseq)
			}
		}()
		_ = nseq
		if !(Nil_(nseq)) {
			return (&CljsCoreValSeq{nseq, coll.X_meta})
		} else {
			return CljsCoreIEmptyList(CljsCoreList_EMPTY)
		}
	}
}

func (_ *CljsCoreValSeq) CljsCoreISeqable__() {}

func (coll *CljsCoreValSeq) X_seq_Arity1() interface{} {
	return coll
}

func (_ *CljsCoreValSeq) CljsCoreISequential__() {}

func (_ *CljsCoreValSeq) CljsCoreIWithMeta__() {}

func (coll *CljsCoreValSeq) X_with_meta_Arity2(new_meta interface{}) interface{} {
	return (&CljsCoreValSeq{coll.Mseq, new_meta})
}

func (_ *CljsCoreValSeq) CljsCoreICollection__() {}

func (coll *CljsCoreValSeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

var X__GT_ValSeq *AFn

// Returns a sequence of the map's values.
var Vals *AFn

// Returns the value in the map entry.
var Val *AFn

// Returns a map that consists of the rest of the maps conj-ed onto
// the first.  If a key occurs in more than one map, the mapping from
// the latter (left-to-right) will be the mapping in the result.
// @param {...*} var_args
var Merge *AFn

// Returns a map that consists of the rest of the maps conj-ed onto
// the first.  If a key occurs in more than one map, the mapping(s)
// from the latter (left-to-right) will be combined with the mapping in
// the result by calling (f val-in-result val-in-latter).
// @param {...*} var_args
var Merge_with *AFn

// Returns a map containing only those entries in map whose key is in keys
var Select_keys *AFn

type CljsCorePersistentHashSet struct {
	Meta     interface{}
	Hash_map interface{}
	X__hash  interface{}
}

func (_ *CljsCorePersistentHashSet) CljsCoreObject__() {}

func (coll *CljsCorePersistentHashSet) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCorePersistentHashSet) String() string {
	return this.ToString()
}

func (this *CljsCorePersistentHashSet) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (coll *CljsCorePersistentHashSet) Keys() interface{} {
	return Iterator.X_invoke_Arity1(Seq.Arity1IQ(coll)).(*CljsCoreIterator)
}

func (coll *CljsCorePersistentHashSet) Entries() interface{} {
	return Set_entries_iterator.X_invoke_Arity1(Seq.Arity1IQ(coll)).(*CljsCoreSetEntriesIterator)
}

func (coll *CljsCorePersistentHashSet) Values() interface{} {
	return Iterator.X_invoke_Arity1(Seq.Arity1IQ(coll)).(*CljsCoreIterator)
}

func (coll *CljsCorePersistentHashSet) Has(k interface{}) interface{} {
	return Contains_QMARK_.Arity2IIB(coll, k)
}

func (coll *CljsCorePersistentHashSet) ForEach(f interface{}) interface{} {
	{
		var seq__6404 interface{} = Seq.Arity1IQ(coll)
		var chunk__6405 interface{} = nil
		var count__6406 = float64(0)
		var i__6407 = float64(0)
		_, _, _, _ = seq__6404, chunk__6405, count__6406, i__6407
		for {
			if i__6407 < count__6406 {
				{
					var vec__6408 = chunk__6405.(CljsCoreIIndexed).X_nth_Arity2(i__6407)
					var k = Nth.X_invoke_Arity3(vec__6408, float64(0), nil)
					var v = Nth.X_invoke_Arity3(vec__6408, float64(1), nil)
					_, _, _ = vec__6408, k, v
					{
						var G__6409_7944 = v
						var G__6410_7945 = k
						_, _ = G__6409_7944, G__6410_7945
						f.(CljsCoreIFn).X_invoke_Arity2(G__6409_7944, G__6410_7945)
					}
					seq__6404, chunk__6405, count__6406, i__6407 = seq__6404, chunk__6405, count__6406, (i__6407 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__6404)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__6404___1 = temp__4222__auto__
							_ = seq__6404___1
							if Chunked_seq_QMARK_.Arity1IB(seq__6404___1) {
								{
									var c__954__auto__ = Chunk_first.X_invoke_Arity1(seq__6404___1)
									_ = c__954__auto__
									seq__6404, chunk__6405, count__6406, i__6407 = Chunk_rest.X_invoke_Arity1(seq__6404___1), c__954__auto__, Count.X_invoke_Arity1(c__954__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__6411 = First.X_invoke_Arity1(seq__6404___1)
									var k = Nth.X_invoke_Arity3(vec__6411, float64(0), nil)
									var v = Nth.X_invoke_Arity3(vec__6411, float64(1), nil)
									_, _, _ = vec__6411, k, v
									{
										var G__6412_7946 = v
										var G__6413_7947 = k
										_, _ = G__6412_7946, G__6413_7947
										f.(CljsCoreIFn).X_invoke_Arity2(G__6412_7946, G__6413_7947)
									}
									seq__6404, chunk__6405, count__6406, i__6407 = Next.Arity1IQ(seq__6404___1), nil, float64(0), float64(0)
									continue
								}
							}
						}
					} else {
						return nil
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreILookup__() {}

func (coll *CljsCorePersistentHashSet) X_lookup_Arity2(v interface{}) interface{} {
	return coll.X_lookup_Arity3(v, nil)
}

func (coll *CljsCorePersistentHashSet) X_lookup_Arity3(v interface{}, not_found interface{}) interface{} {
	if coll.Hash_map.(CljsCoreIAssociative).X_contains_key_QMARK__Arity2(v) {
		return v
	} else {
		return not_found
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIMeta__() {}

func (coll *CljsCorePersistentHashSet) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCorePersistentHashSet) CljsCoreICloneable__() {}

func (___ *CljsCorePersistentHashSet) X_clone_Arity1() interface{} {
	return (&CljsCorePersistentHashSet{___.Meta, ___.Hash_map, ___.X__hash})
}

func (_ *CljsCorePersistentHashSet) CljsCoreICounted__() {}

func (coll *CljsCorePersistentHashSet) X_count_Arity1() float64 {
	return coll.Hash_map.(CljsCoreICounted).X_count_Arity1()
}

func (_ *CljsCorePersistentHashSet) CljsCoreIHash__() {}

func (coll *CljsCorePersistentHashSet) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentHashSet) X_equiv_Arity2(other interface{}) bool {
	return (Set_QMARK_.Arity1IB(other)) && (Count.X_invoke_Arity1(coll).(float64) == Count.X_invoke_Arity1(other).(float64)) && (Every_QMARK_.Arity2IIB(func(G__7948 *AFn) *AFn {
		return Fn(G__7948, 1, func(p1__6393_SHARP_ interface{}) interface{} {
			return Contains_QMARK_.Arity2IIB(coll, p1__6393_SHARP_)
		})
	}(&AFn{}), other))
}

func (_ *CljsCorePersistentHashSet) CljsCoreIEditableCollection__() {}

func (coll *CljsCorePersistentHashSet) X_as_transient_Arity1() interface{} {
	return (&CljsCoreTransientHashSet{coll.Hash_map.(CljsCoreIEditableCollection).X_as_transient_Arity1()})
}

func (_ *CljsCorePersistentHashSet) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCorePersistentHashSet) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCorePersistentHashSet_EMPTY, coll.Meta)
}

func (_ *CljsCorePersistentHashSet) CljsCoreISet__() {}

func (coll *CljsCorePersistentHashSet) X_disjoin_Arity2(v interface{}) interface{} {
	return (&CljsCorePersistentHashSet{coll.Meta, coll.Hash_map.(CljsCoreIMap).X_dissoc_Arity2(v), nil})
}

func (_ *CljsCorePersistentHashSet) CljsCoreISeqable__() {}

func (coll *CljsCorePersistentHashSet) X_seq_Arity1() interface{} {
	return Keys.X_invoke_Arity1(coll.Hash_map)
}

func (_ *CljsCorePersistentHashSet) CljsCoreIWithMeta__() {}

func (coll *CljsCorePersistentHashSet) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCorePersistentHashSet{meta___1, coll.Hash_map, coll.X__hash})
}

func (_ *CljsCorePersistentHashSet) CljsCoreICollection__() {}

func (coll *CljsCorePersistentHashSet) X_conj_Arity2(o interface{}) interface{} {
	return (&CljsCorePersistentHashSet{coll.Meta, Assoc.X_invoke_Arity3(coll.Hash_map, o, nil), nil})
}

func (_ *CljsCorePersistentHashSet) CljsCoreIFn__() {}

func (this *CljsCorePersistentHashSet) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (coll *CljsCorePersistentHashSet) X_invoke_Arity1(k interface{}) interface{} {
	return coll.X_lookup_Arity2(k)
}

func (coll *CljsCorePersistentHashSet) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return coll.X_lookup_Arity3(k, not_found)
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_PersistentHashSet *AFn

var CljsCorePersistentHashSet_EMPTY = (&CljsCorePersistentHashSet{nil, CljsCorePersistentArrayMap_EMPTY, float64(0)})

var CljsCorePersistentHashSet_FromArray = func(G__7949 *AFn) *AFn {
	return Fn(G__7949, 2, func(items interface{}, no_clone bool) interface{} {
		{
			var len = Alength_(items)
			_ = len
			if len <= CljsCorePersistentArrayMap_HASHMAP_THRESHOLD {
				{
					var arr = func() interface{} {
						if Truth_(no_clone) {
							return items
						} else {
							return Aclone.X_invoke_Arity1(items).([]interface{})
						}
					}()
					_ = arr
					{
						var i = float64(0)
						var out interface{} = Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY)
						_, _ = i, out
						for {
							if i < len {
								i, out = (i + float64(1)), out.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3(Aget_(items, i), nil)
								continue
							} else {
								return (&CljsCorePersistentHashSet{nil, out.(CljsCoreITransientCollection).X_persistent_BANG__Arity1(), nil})
							}
						}
					}
				}
			} else {
				{
					var i = float64(0)
					var out interface{} = Transient.X_invoke_Arity1(CljsCorePersistentHashSet_EMPTY)
					_, _ = i, out
					for {
						if i < len {
							i, out = (i + float64(1)), out.(CljsCoreITransientCollection).X_conj_BANG__Arity2(Aget_(items, i))
							continue
						} else {
							return out.(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
						}
					}
				}
			}
		}
	})
}(&AFn{})

type CljsCoreTransientHashSet struct{ Transient_map interface{} }

func (_ *CljsCoreTransientHashSet) CljsCoreIFn__() {}

func (this *CljsCoreTransientHashSet) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (tcoll *CljsCoreTransientHashSet) X_invoke_Arity1(k interface{}) interface{} {
	if reflect.DeepEqual(tcoll.Transient_map.(CljsCoreILookup).X_lookup_Arity3(k, Lookup_sentinel), Lookup_sentinel) {
		return nil
	} else {
		return k
	}
}

func (tcoll *CljsCoreTransientHashSet) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	if reflect.DeepEqual(tcoll.Transient_map.(CljsCoreILookup).X_lookup_Arity3(k, Lookup_sentinel), Lookup_sentinel) {
		return not_found
	} else {
		return k
	}
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

func (_ *CljsCoreTransientHashSet) CljsCoreILookup__() {}

func (tcoll *CljsCoreTransientHashSet) X_lookup_Arity2(v interface{}) interface{} {
	return tcoll.X_lookup_Arity3(v, nil)
}

func (tcoll *CljsCoreTransientHashSet) X_lookup_Arity3(v interface{}, not_found interface{}) interface{} {
	if reflect.DeepEqual(tcoll.Transient_map.(CljsCoreILookup).X_lookup_Arity3(v, Lookup_sentinel), Lookup_sentinel) {
		return not_found
	} else {
		return v
	}
}

func (_ *CljsCoreTransientHashSet) CljsCoreICounted__() {}

func (tcoll *CljsCoreTransientHashSet) X_count_Arity1() float64 {
	return Count.X_invoke_Arity1(tcoll.Transient_map).(float64)
}

func (_ *CljsCoreTransientHashSet) CljsCoreITransientSet__() {}

func (tcoll *CljsCoreTransientHashSet) X_disjoin_BANG__Arity2(v interface{}) interface{} {
	tcoll.Transient_map = Dissoc_BANG_.X_invoke_Arity2(tcoll.Transient_map, v)

	return tcoll
}

func (_ *CljsCoreTransientHashSet) CljsCoreITransientCollection__() {}

func (tcoll *CljsCoreTransientHashSet) X_conj_BANG__Arity2(o interface{}) interface{} {
	tcoll.Transient_map = Assoc_BANG_.X_invoke_Arity3(tcoll.Transient_map, o, nil)

	return tcoll
}

func (tcoll *CljsCoreTransientHashSet) X_persistent_BANG__Arity1() interface{} {
	return (&CljsCorePersistentHashSet{nil, Persistent_BANG_.X_invoke_Arity1(tcoll.Transient_map), nil})
}

var X__GT_TransientHashSet *AFn

type CljsCorePersistentTreeSet struct {
	Meta     interface{}
	Tree_map interface{}
	X__hash  interface{}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreObject__() {}

func (coll *CljsCorePersistentTreeSet) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCorePersistentTreeSet) String() string {
	return this.ToString()
}

func (this *CljsCorePersistentTreeSet) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (coll *CljsCorePersistentTreeSet) Keys() interface{} {
	return Iterator.X_invoke_Arity1(Seq.Arity1IQ(coll)).(*CljsCoreIterator)
}

func (coll *CljsCorePersistentTreeSet) Entries() interface{} {
	return Set_entries_iterator.X_invoke_Arity1(Seq.Arity1IQ(coll)).(*CljsCoreSetEntriesIterator)
}

func (coll *CljsCorePersistentTreeSet) Values() interface{} {
	return Iterator.X_invoke_Arity1(Seq.Arity1IQ(coll)).(*CljsCoreIterator)
}

func (coll *CljsCorePersistentTreeSet) Has(k interface{}) interface{} {
	return Contains_QMARK_.Arity2IIB(coll, k)
}

func (coll *CljsCorePersistentTreeSet) ForEach(f interface{}) interface{} {
	{
		var seq__6437 interface{} = Seq.Arity1IQ(coll)
		var chunk__6438 interface{} = nil
		var count__6439 = float64(0)
		var i__6440 = float64(0)
		_, _, _, _ = seq__6437, chunk__6438, count__6439, i__6440
		for {
			if i__6440 < count__6439 {
				{
					var vec__6441 = chunk__6438.(CljsCoreIIndexed).X_nth_Arity2(i__6440)
					var k = Nth.X_invoke_Arity3(vec__6441, float64(0), nil)
					var v = Nth.X_invoke_Arity3(vec__6441, float64(1), nil)
					_, _, _ = vec__6441, k, v
					{
						var G__6442_7950 = v
						var G__6443_7951 = k
						_, _ = G__6442_7950, G__6443_7951
						f.(CljsCoreIFn).X_invoke_Arity2(G__6442_7950, G__6443_7951)
					}
					seq__6437, chunk__6438, count__6439, i__6440 = seq__6437, chunk__6438, count__6439, (i__6440 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__6437)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__6437___1 = temp__4222__auto__
							_ = seq__6437___1
							if Chunked_seq_QMARK_.Arity1IB(seq__6437___1) {
								{
									var c__954__auto__ = Chunk_first.X_invoke_Arity1(seq__6437___1)
									_ = c__954__auto__
									seq__6437, chunk__6438, count__6439, i__6440 = Chunk_rest.X_invoke_Arity1(seq__6437___1), c__954__auto__, Count.X_invoke_Arity1(c__954__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__6444 = First.X_invoke_Arity1(seq__6437___1)
									var k = Nth.X_invoke_Arity3(vec__6444, float64(0), nil)
									var v = Nth.X_invoke_Arity3(vec__6444, float64(1), nil)
									_, _, _ = vec__6444, k, v
									{
										var G__6445_7952 = v
										var G__6446_7953 = k
										_, _ = G__6445_7952, G__6446_7953
										f.(CljsCoreIFn).X_invoke_Arity2(G__6445_7952, G__6446_7953)
									}
									seq__6437, chunk__6438, count__6439, i__6440 = Next.Arity1IQ(seq__6437___1), nil, float64(0), float64(0)
									continue
								}
							}
						}
					} else {
						return nil
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreILookup__() {}

func (coll *CljsCorePersistentTreeSet) X_lookup_Arity2(v interface{}) interface{} {
	return coll.X_lookup_Arity3(v, nil)
}

func (coll *CljsCorePersistentTreeSet) X_lookup_Arity3(v interface{}, not_found interface{}) interface{} {
	{
		var n = Native_invoke_instance_method.X_invoke_Arity3(coll.Tree_map, "Entry_at", []interface{}{v})
		_ = n
		if !(Nil_(n)) {
			return Native_get_instance_field.X_invoke_Arity2(n, "Key")
		} else {
			return not_found
		}
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIMeta__() {}

func (coll *CljsCorePersistentTreeSet) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCorePersistentTreeSet) CljsCoreICloneable__() {}

func (___ *CljsCorePersistentTreeSet) X_clone_Arity1() interface{} {
	return (&CljsCorePersistentTreeSet{___.Meta, ___.Tree_map, ___.X__hash})
}

func (_ *CljsCorePersistentTreeSet) CljsCoreICounted__() {}

func (coll *CljsCorePersistentTreeSet) X_count_Arity1() float64 {
	return Count.X_invoke_Arity1(coll.Tree_map).(float64)
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIReversible__() {}

func (coll *CljsCorePersistentTreeSet) X_rseq_Arity1() interface{} {
	if Count.X_invoke_Arity1(coll.Tree_map).(float64) > float64(0) {
		return Map_.X_invoke_Arity2(Key, Rseq.Arity1IQ(coll.Tree_map)).(*CljsCoreLazySeq)
	} else {
		return nil
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIHash__() {}

func (coll *CljsCorePersistentTreeSet) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = coll.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__577__auto_____1
				coll.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentTreeSet) X_equiv_Arity2(other interface{}) bool {
	return (Set_QMARK_.Arity1IB(other)) && (Count.X_invoke_Arity1(coll).(float64) == Count.X_invoke_Arity1(other).(float64)) && (Every_QMARK_.Arity2IIB(func(G__7954 *AFn) *AFn {
		return Fn(G__7954, 1, func(p1__6426_SHARP_ interface{}) interface{} {
			return Contains_QMARK_.Arity2IIB(coll, p1__6426_SHARP_)
		})
	}(&AFn{}), other))
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCorePersistentTreeSet) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCorePersistentTreeSet_EMPTY, coll.Meta)
}

func (_ *CljsCorePersistentTreeSet) CljsCoreISet__() {}

func (coll *CljsCorePersistentTreeSet) X_disjoin_Arity2(v interface{}) interface{} {
	return (&CljsCorePersistentTreeSet{coll.Meta, Dissoc.X_invoke_Arity2(coll.Tree_map, v), nil})
}

func (_ *CljsCorePersistentTreeSet) CljsCoreISeqable__() {}

func (coll *CljsCorePersistentTreeSet) X_seq_Arity1() interface{} {
	return Keys.X_invoke_Arity1(coll.Tree_map)
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIWithMeta__() {}

func (coll *CljsCorePersistentTreeSet) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCorePersistentTreeSet{meta___1, coll.Tree_map, coll.X__hash})
}

func (_ *CljsCorePersistentTreeSet) CljsCoreICollection__() {}

func (coll *CljsCorePersistentTreeSet) X_conj_Arity2(o interface{}) interface{} {
	return (&CljsCorePersistentTreeSet{coll.Meta, Assoc.X_invoke_Arity3(coll.Tree_map, o, nil), nil})
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIFn__() {}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (coll *CljsCorePersistentTreeSet) X_invoke_Arity1(k interface{}) interface{} {
	return coll.X_lookup_Arity2(k)
}

func (coll *CljsCorePersistentTreeSet) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return coll.X_lookup_Arity3(k, not_found)
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

func (_ *CljsCorePersistentTreeSet) CljsCoreISorted__() {}

func (coll *CljsCorePersistentTreeSet) X_sorted_seq_Arity2(ascending_QMARK_ interface{}) interface{} {
	return Map_.X_invoke_Arity2(Key, coll.Tree_map.(CljsCoreISorted).X_sorted_seq_Arity2(ascending_QMARK_)).(*CljsCoreLazySeq)
}

func (coll *CljsCorePersistentTreeSet) X_sorted_seq_from_Arity3(k interface{}, ascending_QMARK_ interface{}) interface{} {
	return Map_.X_invoke_Arity2(Key, coll.Tree_map.(CljsCoreISorted).X_sorted_seq_from_Arity3(k, ascending_QMARK_)).(*CljsCoreLazySeq)
}

func (coll *CljsCorePersistentTreeSet) X_entry_key_Arity2(entry interface{}) interface{} {
	return entry
}

func (coll *CljsCorePersistentTreeSet) X_comparator_Arity1() interface{} {
	return coll.Tree_map.(CljsCoreISorted).X_comparator_Arity1()
}

var X__GT_PersistentTreeSet *AFn

var CljsCorePersistentTreeSet_EMPTY = (&CljsCorePersistentTreeSet{nil, CljsCorePersistentTreeMap_EMPTY, float64(0)})

var Set_from_indexed_seq *AFn

// Returns a set of the distinct elements of coll.
var Set *AFn

// @param {...*} var_args
var Hash_set *AFn

// Returns a new sorted set with supplied keys.
// @param {...*} var_args
var Sorted_set *AFn

// Returns a new sorted set with supplied keys, using the supplied comparator.
// @param {...*} var_args
var Sorted_set_by *AFn

// Given a map of replacement pairs and a vector/collection, returns a
// vector/seq with any elements = a key in smap replaced with the
// corresponding val in smap.  Returns a transducer when no collection
// is provided.
var Replace *AFn

// Returns a lazy sequence of the elements of coll with duplicates removed
var Distinct *AFn

var Butlast *AFn

// Returns the name String of a string, symbol or keyword.
var Name *AFn

// Returns a map with the keys mapped to the corresponding vals.
var Zipmap *AFn

// Returns the x for which (k x), a number, is greatest.
// @param {...*} var_args
var Max_key *AFn

// Returns the x for which (k x), a number, is least.
// @param {...*} var_args
var Min_key *AFn

type CljsCoreArrayList struct{ Arr interface{} }

func (_ *CljsCoreArrayList) CljsCoreObject__() {}

func (___ *CljsCoreArrayList) Add(x interface{}) interface{} {
	return js.JSArray_(&___.Arr).Push(x)
}

func (___ *CljsCoreArrayList) Size() interface{} {
	return Alength_(___.Arr)
}

func (___ *CljsCoreArrayList) Clear() interface{} {
	return func() interface{} {
		var return__7955 = []interface{}{}
		___.Arr = return__7955
		return return__7955
	}()
}

func (___ *CljsCoreArrayList) IsEmpty() interface{} {
	return (Alength_(___.Arr) == float64(0))
}

func (___ *CljsCoreArrayList) ToArray() interface{} {
	return ___.Arr
}

var X__GT_ArrayList *AFn

var Array_list *AFn

// Returns a lazy sequence of lists like partition, but may include
// partitions with fewer than n items at the end.  Returns a stateful
// transducer when no collection is provided.
var Partition_all *AFn

// Returns a lazy sequence of successive items from coll while
// (pred item) returns true. pred must be free of side-effects.
// Returns a transducer when no collection is provided.
var Take_while *AFn

var Mk_bound_fn *AFn

// sc must be a sorted collection, test(s) one of <, <=, > or
// >=. Returns a seq of those entries with keys ek for
// which (test (.. sc comparator (compare ek key)) 0) is true
var Subseq *AFn

// sc must be a sorted collection, test(s) one of <, <=, > or
// >=. Returns a reverse seq of those entries with keys ek for
// which (test (.. sc comparator (compare ek key)) 0) is true
var Rsubseq *AFn

type CljsCoreRange struct {
	Meta    interface{}
	Start   interface{}
	End     interface{}
	Step    interface{}
	X__hash interface{}
}

func (_ *CljsCoreRange) CljsCoreObject__() {}

func (coll *CljsCoreRange) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreRange) String() string {
	return this.ToString()
}

func (this *CljsCoreRange) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreRange) CljsCoreIIndexed__() {}

func (rng *CljsCoreRange) X_nth_Arity2(n interface{}) interface{} {
	if n.(float64) < rng.X_count_Arity1() {
		return (rng.Start.(float64) + (n.(float64) * rng.Step.(float64)))
	} else {
		if (rng.Start.(float64) > rng.End.(float64)) && (rng.Step.(float64) == float64(0)) {
			return rng.Start
		} else {
			panic((&js.Error{"Index out of bounds"}))
		}
	}
}

func (rng *CljsCoreRange) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	if n.(float64) < rng.X_count_Arity1() {
		return (rng.Start.(float64) + (n.(float64) * rng.Step.(float64)))
	} else {
		if (rng.Start.(float64) > rng.End.(float64)) && (rng.Step.(float64) == float64(0)) {
			return rng.Start
		} else {
			return not_found
		}
	}
}

func (_ *CljsCoreRange) CljsCoreIMeta__() {}

func (rng *CljsCoreRange) X_meta_Arity1() interface{} {
	return rng.Meta
}

func (_ *CljsCoreRange) CljsCoreICloneable__() {}

func (___ *CljsCoreRange) X_clone_Arity1() interface{} {
	return (&CljsCoreRange{___.Meta, ___.Start, ___.End, ___.Step, ___.X__hash})
}

func (_ *CljsCoreRange) CljsCoreINext__() {}

func (rng *CljsCoreRange) X_next_Arity1() interface{} {
	if rng.Step.(float64) > float64(0) {
		if (rng.Start.(float64) + rng.Step.(float64)) < rng.End.(float64) {
			return (&CljsCoreRange{rng.Meta, (rng.Start.(float64) + rng.Step.(float64)), rng.End, rng.Step, nil})
		} else {
			return nil
		}
	} else {
		if (rng.Start.(float64) + rng.Step.(float64)) > rng.End.(float64) {
			return (&CljsCoreRange{rng.Meta, (rng.Start.(float64) + rng.Step.(float64)), rng.End, rng.Step, nil})
		} else {
			return nil
		}
	}
}

func (_ *CljsCoreRange) CljsCoreICounted__() {}

func (rng *CljsCoreRange) X_count_Arity1() float64 {
	if Not.Arity1IB(rng.X_seq_Arity1()) {
		return float64(0)
	} else {
		{
			var G__6557 = ((rng.End.(float64) - rng.Start.(float64)) / rng.Step.(float64))
			_ = G__6557
			return Native_invoke_func.X_invoke_Arity2(Math.Ceil, []interface{}{G__6557}).(float64)
		}
	}
}

func (_ *CljsCoreRange) CljsCoreIHash__() {}

func (rng *CljsCoreRange) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = rng.X__hash
		_ = h__577__auto__
		if !(Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = Hash_ordered_coll.Arity1IF(rng)
				_ = h__577__auto_____1
				rng.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCoreRange) CljsCoreIEquiv__() {}

func (rng *CljsCoreRange) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_sequential.X_invoke_Arity2(rng, other))
}

func (_ *CljsCoreRange) CljsCoreIEmptyableCollection__() {}

func (rng *CljsCoreRange) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, rng.Meta)
}

func (_ *CljsCoreRange) CljsCoreIReduce__() {}

func (rng *CljsCoreRange) X_reduce_Arity2(f interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity2(rng, f)
}

func (rng *CljsCoreRange) X_reduce_Arity3(f interface{}, s interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity3(rng, f, s)
}

func (_ *CljsCoreRange) CljsCoreISeq__() {}

func (rng *CljsCoreRange) X_first_Arity1() interface{} {
	if Nil_(rng.X_seq_Arity1()) {
		return nil
	} else {
		return rng.Start
	}
}

func (rng *CljsCoreRange) X_rest_Arity1() interface{} {
	if !(Nil_(rng.X_seq_Arity1())) {
		return (&CljsCoreRange{rng.Meta, (rng.Start.(float64) + rng.Step.(float64)), rng.End, rng.Step, nil})
	} else {
		return CljsCoreIEmptyList(CljsCoreList_EMPTY)
	}
}

func (_ *CljsCoreRange) CljsCoreISeqable__() {}

func (rng *CljsCoreRange) X_seq_Arity1() interface{} {
	if rng.Step.(float64) > float64(0) {
		if rng.Start.(float64) < rng.End.(float64) {
			return rng
		} else {
			return nil
		}
	} else {
		if rng.Start.(float64) > rng.End.(float64) {
			return rng
		} else {
			return nil
		}
	}
}

func (_ *CljsCoreRange) CljsCoreISequential__() {}

func (_ *CljsCoreRange) CljsCoreIWithMeta__() {}

func (rng *CljsCoreRange) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCoreRange{meta___1, rng.Start, rng.End, rng.Step, rng.X__hash})
}

func (_ *CljsCoreRange) CljsCoreICollection__() {}

func (rng *CljsCoreRange) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, rng).(*CljsCoreCons)
}

var X__GT_Range *AFn

// Returns a lazy seq of nums from start (inclusive) to end
// (exclusive), by step, where start defaults to 0, step to 1,
// and end to infinity.
var Range_ *AFn

// Returns a lazy seq of every nth item in coll.  Returns a stateful
// transducer when no collection is provided.
var Take_nth *AFn

// Returns a vector of [(take-while pred coll) (drop-while pred coll)]
var Split_with *AFn

// Applies f to each value in coll, splitting it each time f returns a
// new value.  Returns a lazy seq of partitions.  Returns a stateful
// transducer when no collection is provided.
var Partition_by *AFn

// Returns a map from distinct items in coll to the number of times
// they appear.
var Frequencies *AFn

// Returns a lazy seq of the intermediate values of the reduction (as
// per reduce) of coll by f, starting with init.
var Reductions *AFn

// Takes a set of functions and returns a fn that is the juxtaposition
// of those fns.  The returned fn takes a variable number of args, and
// returns a vector containing the result of applying each fn to the
// args (left-to-right).
// ((juxt a b c) x) => [(a x) (b x) (c x)]
// @param {...*} var_args
var Juxt *AFn

// When lazy sequences are produced via functions that have side
// effects, any effects other than those needed to produce the first
// element in the seq do not occur until the seq is consumed. dorun can
// be used to force any effects. Walks through the successive nexts of
// the seq, does not retain the head and returns nil.
var Dorun *AFn

// When lazy sequences are produced via functions that have side
// effects, any effects other than those needed to produce the first
// element in the seq do not occur until the seq is consumed. doall can
// be used to force any effects. Walks through the successive nexts of
// the seq, retains the head and returns it, thus causing the entire
// seq to reside in memory at one time.
var Doall *AFn

var Regexp_QMARK_ *AFn

// Returns the result of (re-find re s) if re fully matches s.
var Re_matches *AFn

// Returns the first regex match, if any, of s to re, using
// re.exec(s). Returns a vector, containing first the matching
// substring, then any capturing groups if the regular expression contains
// capturing groups.
var Re_find *AFn

// Returns a lazy sequence of successive matches of re in s.
var Re_seq *AFn

// Returns an instance of RegExp which has compiled the provided string.
var Re_pattern *AFn

// @param {...*} var_args
var Write_all *AFn

var String_print *AFn

var Flush *AFn

var Pr_seq_writer *AFn

var Pr_sb_with_opts *AFn

// Prints a sequence of objects to a string, observing all the
// options given in opts
var Pr_str_with_opts *AFn

// Same as pr-str-with-opts followed by (newline)
var Prn_str_with_opts *AFn

// Prints a sequence of objects using string-print, observing all
// the options given in opts
var Pr_with_opts *AFn

var Newline *AFn

// pr to a string, returning it. Fundamental entrypoint to IPrintWithWriter.
// @param {...*} var_args
var Pr_str *AFn

// Same as pr-str followed by (newline)
// @param {...*} var_args
var Prn_str *AFn

// Prints the object(s) using string-print.  Prints the
// object(s), separated by spaces if there is more than one.
// By default, pr and prn print in a way that objects can be
// read by the reader
// @param {...*} var_args
var Pr *AFn

// Prints the object(s) using string-print.
// print and println produce output for human consumption.
// @param {...*} var_args
var Print *AFn

// print to a string, returning it
// @param {...*} var_args
var Print_str *AFn

// Same as print followed by (newline)
// @param {...*} var_args
var Println *AFn

// println to a string, returning it
// @param {...*} var_args
var Println_str *AFn

// Same as pr followed by (newline).
// @param {...*} var_args
var Prn *AFn

var Print_map *AFn

func (_ *CljsCoreIndexedSeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreIndexedSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreLazySeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreLazySeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCorePersistentTreeMapSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreNodeSeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreNodeSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreBlackNode) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreBlackNode) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "[", " ", "]", opts, coll)
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCorePersistentTreeSet) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "#{", " ", "}", opts, coll)
}

func (_ *CljsCoreChunkedSeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreChunkedSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreCons) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreCons) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreRSeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreRSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCorePersistentHashMap) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCorePersistentHashMap) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Print_map.X_invoke_Arity4(coll, Pr_writer, writer, opts)
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreArrayNodeSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreSubvec) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreSubvec) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "[", " ", "]", opts, coll)
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCorePersistentTreeMap) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Print_map.X_invoke_Arity4(coll, Pr_writer, writer, opts)
}

func (_ *CljsCorePersistentHashSet) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCorePersistentHashSet) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "#{", " ", "}", opts, coll)
}

func (_ *CljsCoreChunkedCons) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreChunkedCons) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreAtom) CljsCoreIPrintWithWriter__() {}

func (a *CljsCoreAtom) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	writer.(CljsCoreIWriter).X_write_Arity2("#<Atom: ")
	Pr_writer.X_invoke_Arity3(a.State, writer, opts)
	return writer.(CljsCoreIWriter).X_write_Arity2(">")
}

func (_ *CljsCoreValSeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreValSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreRedNode) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreRedNode) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "[", " ", "]", opts, coll)
}

func (_ *CljsCorePersistentVector) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCorePersistentVector) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "[", " ", "]", opts, coll)
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCorePersistentQueueSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreEmptyList) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreEmptyList) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return writer.(CljsCoreIWriter).X_write_Arity2("()")
}

func (_ *CljsCoreLazyTransformer) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreLazyTransformer) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCorePersistentQueue) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCorePersistentQueue) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "#queue [", " ", "]", opts, Seq.Arity1IQ(coll))
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCorePersistentArrayMap) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Print_map.X_invoke_Arity4(coll, Pr_writer, writer, opts)
}

func (_ *CljsCoreRange) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreRange) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreKeySeq) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreKeySeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCoreList) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreList) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll)
}

func (_ *CljsCorePersistentVector) CljsCoreIComparable__() {}

func (x *CljsCorePersistentVector) X_compare_Arity2(y interface{}) float64 {
	return Compare_indexed.X_invoke_Arity2(x, y).(float64)
}

func (_ *CljsCoreSubvec) CljsCoreIComparable__() {}

func (x *CljsCoreSubvec) X_compare_Arity2(y interface{}) float64 {
	return Compare_indexed.X_invoke_Arity2(x, y).(float64)
}

func (_ *CljsCoreKeyword) CljsCoreIComparable__() {}

func (x *CljsCoreKeyword) X_compare_Arity2(y interface{}) float64 {
	return Compare_symbols.X_invoke_Arity2(x, y).(float64)
}

func (_ *CljsCoreSymbol) CljsCoreIComparable__() {}

func (x *CljsCoreSymbol) X_compare_Arity2(y interface{}) float64 {
	return Compare_symbols.X_invoke_Arity2(x, y).(float64)
}

// Atomically sets the metadata for a namespace/var/ref/agent/atom to be:
//
// (apply f its-current-meta args)
//
// f must be free of side-effects
// @param {...*} var_args
var Alter_meta_BANG_ *AFn

// Atomically resets the metadata for an atom
var Reset_meta_BANG_ *AFn

// Alpha - subject to change.
//
// Adds a watch function to an atom reference. The watch fn must be a
// fn of 4 args: a key, the reference, its old-state, its
// new-state. Whenever the reference's state might have been changed,
// any registered watches will have their functions called. The watch
// fn will be called synchronously. Note that an atom's state
// may have changed again prior to the fn call, so use old/new-state
// rather than derefing the reference. Keys must be unique per
// reference, and can be used to remove the watch with remove-watch,
// but are otherwise considered opaque by the watch mechanism.  Bear in
// mind that regardless of the result or action of the watch fns the
// atom's value will change.  Example:
//
// (def a (atom 0))
// (add-watch a :inc (fn [k r o n] (assert (== 0 n))))
// (swap! a inc)
// ;; Assertion Error
// (deref a)
// ;=> 1
var Add_watch *AFn

// Alpha - subject to change.
//
// Removes a watch (set by add-watch) from a reference
var Remove_watch *AFn

var Gensym_counter interface{}

// Returns a new symbol with a unique name. If a prefix string is
// supplied, the name is prefix# where # is some unique number. If
// prefix is not supplied, the prefix is 'G__'.
var Gensym *AFn

var Fixture1 float64

var Fixture2 float64

type CljsCoreDelay struct {
	F     interface{}
	Value interface{}
}

func (_ *CljsCoreDelay) CljsCoreIPending__() {}

func (d *CljsCoreDelay) X_realized_QMARK__Arity1() bool {
	return Not.Arity1IB(d.F)
}

func (_ *CljsCoreDelay) CljsCoreIDeref__() {}

func (___ *CljsCoreDelay) X_deref_Arity1() interface{} {
	if Truth_(___.F) {
		___.Value = func() interface{} {
			return ___.F.(CljsCoreIFn).X_invoke_Arity0()
		}()

		___.F = nil

	} else {
	}
	return ___.Value
}

var X__GT_Delay *AFn

// returns true if x is a Delay created with delay
var Delay_QMARK_ *AFn

// If x is a Delay, returns the (possibly cached) value of its expression, else returns x
var Force *AFn

// Returns true if a value has been produced for a promise, delay, future or lazy sequence.
var Realized_QMARK_ *AFn

var Preserving_reduced *AFn

// A transducer which concatenates the contents of each input, which must be a
// collection, into the reduction.
var Cat *AFn

// Returns a lazy sequence removing consecutive duplicates in coll.
// Returns a transducer when no collection is provided.
var Dedupe *AFn

// Returns items from coll with random probability of prob (0.0 -
// 1.0).  Returns a transducer when no collection is provided.
var Random_sample *AFn

type CljsCoreIteration struct {
	Xform interface{}
	Coll  interface{}
}

func (_ *CljsCoreIteration) CljsCoreIPrintWithWriter__() {}

func (coll___1 *CljsCoreIteration) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1)
}

func (_ *CljsCoreIteration) CljsCoreIReduce__() {}

func (___ *CljsCoreIteration) X_reduce_Arity3(f interface{}, init interface{}) interface{} {
	return Transduce.X_invoke_Arity4(___.Xform, f, init, ___.Coll)
}

func (_ *CljsCoreIteration) CljsCoreISeqable__() {}

func (___ *CljsCoreIteration) X_seq_Arity1() interface{} {
	return Seq.Arity1IQ(Sequence.X_invoke_Arity2(___.Xform, ___.Coll))
}

func (_ *CljsCoreIteration) CljsCoreISequential__() {}

var X__GT_Iteration *AFn

// Returns an iterable/seqable/reducible sequence of applications of
// the transducer to the items in coll. Note that these applications
// will be performed every time iterator/seq/reduce is called.
var Iteration *AFn

// Runs the supplied procedure (via reduce), for purposes of side
// effects, on successive items in the collection. Returns nil
var Run_BANG_ *AFn

type CljsCoreIEncodeJS interface {
	CljsCoreIEncodeJS__()
	X_clj__GT_js_Arity1() interface{}
	X_key__GT_js_Arity1() interface{}
}

var X_clj__GT_js *AFn

var X_key__GT_js *AFn

type CljsCoreIEncodeClojure interface {
	CljsCoreIEncodeClojure__()
	X_js__GT_clj_Arity2(options interface{}) interface{}
}

var X_js__GT_clj *AFn

// Returns a memoized version of a referentially transparent function. The
// memoized version of the function keeps a cache of the mapping from arguments
// to results and, when calls with the same arguments are repeated often, has
// higher performance at the expense of higher memory use.
var Memoize *AFn

// trampoline can be used to convert algorithms requiring mutual
// recursion without stack consumption. Calls f with supplied args, if
// any. If f returns a fn, calls that fn with no arguments, and
// continues to repeat, until the return value is not a fn, then
// returns that non-fn value. Note that if you want to return a fn as a
// final value, you must wrap it in some data structure and unpack it
// after trampoline returns.
// @param {...*} var_args
var Trampoline *AFn

// Return a random element of the (sequential) collection. Will have
// the same performance characteristics as nth for the given
// collection.
var Rand_nth *AFn

// Returns a map of the elements of coll keyed by the result of
// f on each element. The value at each key will be a vector of the
// corresponding elements, in the order they appeared in coll.
var Group_by *AFn

// Creates a hierarchy object for use with derive, isa? etc.
var Make_hierarchy *AFn

var X_global_hierarchy interface{}

var Get_global_hierarchy *AFn

// @param {...*} var_args
var Swap_global_hierarchy_BANG_ *AFn

// Returns true if (= child parent), or child is directly or indirectly derived from
// parent, either via a JavaScript type inheritance relationship or a
// relationship established via derive. h must be a hierarchy obtained
// from make-hierarchy, if not supplied defaults to the global
// hierarchy
var Isa_QMARK_ *AFn

// Returns the immediate parents of tag, either via a JavaScript type
// inheritance relationship or a relationship established via derive. h
// must be a hierarchy obtained from make-hierarchy, if not supplied
// defaults to the global hierarchy
var Parents *AFn

// Returns the immediate and indirect parents of tag, either via a JavaScript type
// inheritance relationship or a relationship established via derive. h
// must be a hierarchy obtained from make-hierarchy, if not supplied
// defaults to the global hierarchy
var Ancestors *AFn

// Returns the immediate and indirect children of tag, through a
// relationship established via derive. h must be a hierarchy obtained
// from make-hierarchy, if not supplied defaults to the global
// hierarchy. Note: does not work on JavaScript type inheritance
// relationships.
var Descendants *AFn

// Establishes a parent/child relationship between parent and
// tag. Parent must be a namespace-qualified symbol or keyword and
// child can be either a namespace-qualified symbol or keyword or a
// class. h must be a hierarchy obtained from make-hierarchy, if not
// supplied defaults to, and modifies, the global hierarchy.
var Derive *AFn

// Removes a parent/child relationship between parent and
// tag. h must be a hierarchy obtained from make-hierarchy, if not
// supplied defaults to, and modifies, the global hierarchy.
var Underive *AFn

var Reset_cache *AFn

var Prefers_STAR_ *AFn

var Dominates *AFn

var Find_and_cache_best_method *AFn

type CljsCoreIMultiFn interface {
	CljsCoreIMultiFn__()
	X_reset_Arity1() interface{}
	X_add_method_Arity3(dispatch_val interface{}, method interface{}) interface{}
	X_remove_method_Arity2(dispatch_val interface{}) interface{}
	X_prefer_method_Arity3(dispatch_val interface{}, dispatch_val_y interface{}) interface{}
	X_get_method_Arity2(dispatch_val interface{}) interface{}
	X_methods_Arity1() interface{}
	X_prefers_Arity1() interface{}
}

var X_reset *AFn

var X_add_method *AFn

var X_remove_method *AFn

var X_prefer_method *AFn

var X_get_method *AFn

var X_methods *AFn

var X_prefers *AFn

var Throw_no_method_error *AFn

type CljsCoreMultiFn struct {
	Name                 interface{}
	Dispatch_fn          interface{}
	Default_dispatch_val interface{}
	Hierarchy            interface{}
	Method_table         interface{}
	Prefer_table         interface{}
	Method_cache         interface{}
	Cached_hierarchy     interface{}
}

func (_ *CljsCoreMultiFn) CljsCoreIHash__() {}

func (this *CljsCoreMultiFn) X_hash_Arity1() interface{} {
	{
		var G__7047 = this
		_ = G__7047
		return Native_invoke_func.X_invoke_Arity2(goog.GetUid, []interface{}{G__7047})
	}
}

func (_ *CljsCoreMultiFn) CljsCoreIMultiFn__() {}

func (mf *CljsCoreMultiFn) X_reset_Arity1() interface{} {
	Swap_BANG_.X_invoke_Arity2(mf.Method_table, func(G__7956 *AFn) *AFn {
		return Fn(G__7956, 1, func(mf___1 interface{}) interface{} {
			return CljsCorePersistentArrayMap_EMPTY
		})
	}(&AFn{}))
	Swap_BANG_.X_invoke_Arity2(mf.Method_cache, func(G__7957 *AFn) *AFn {
		return Fn(G__7957, 1, func(mf___1 interface{}) interface{} {
			return CljsCorePersistentArrayMap_EMPTY
		})
	}(&AFn{}))
	Swap_BANG_.X_invoke_Arity2(mf.Prefer_table, func(G__7958 *AFn) *AFn {
		return Fn(G__7958, 1, func(mf___1 interface{}) interface{} {
			return CljsCorePersistentArrayMap_EMPTY
		})
	}(&AFn{}))
	Swap_BANG_.X_invoke_Arity2(mf.Cached_hierarchy, func(G__7959 *AFn) *AFn {
		return Fn(G__7959, 1, func(mf___1 interface{}) interface{} {
			return nil
		})
	}(&AFn{}))
	return mf
}

func (mf *CljsCoreMultiFn) X_add_method_Arity3(dispatch_val interface{}, method interface{}) interface{} {
	Swap_BANG_.X_invoke_Arity4(mf.Method_table, Assoc, dispatch_val, method)
	Reset_cache.X_invoke_Arity4(mf.Method_cache, mf.Method_table, mf.Cached_hierarchy, mf.Hierarchy)
	return mf
}

func (mf *CljsCoreMultiFn) X_remove_method_Arity2(dispatch_val interface{}) interface{} {
	Swap_BANG_.X_invoke_Arity3(mf.Method_table, Dissoc, dispatch_val)
	Reset_cache.X_invoke_Arity4(mf.Method_cache, mf.Method_table, mf.Cached_hierarchy, mf.Hierarchy)
	return mf
}

func (mf *CljsCoreMultiFn) X_get_method_Arity2(dispatch_val interface{}) interface{} {
	if X_EQ_.Arity2IIB(Deref.X_invoke_Arity1(mf.Cached_hierarchy), Deref.X_invoke_Arity1(mf.Hierarchy)) {
	} else {
		Reset_cache.X_invoke_Arity4(mf.Method_cache, mf.Method_table, mf.Cached_hierarchy, mf.Hierarchy)
	}
	{
		var temp__4220__auto__ = Deref.X_invoke_Arity1(mf.Method_cache).(CljsCoreIFn).X_invoke_Arity1(dispatch_val)
		_ = temp__4220__auto__
		if Truth_(temp__4220__auto__) {
			{
				var target_fn = temp__4220__auto__
				_ = target_fn
				return target_fn
			}
		} else {
			{
				var temp__4220__auto_____1 = Find_and_cache_best_method.X_invoke_Arity7(mf.Name, dispatch_val, mf.Hierarchy, mf.Method_table, mf.Prefer_table, mf.Method_cache, mf.Cached_hierarchy)
				_ = temp__4220__auto_____1
				if Truth_(temp__4220__auto_____1) {
					{
						var target_fn = temp__4220__auto_____1
						_ = target_fn
						return target_fn
					}
				} else {
					return Deref.X_invoke_Arity1(mf.Method_table).(CljsCoreIFn).X_invoke_Arity1(mf.Default_dispatch_val)
				}
			}
		}
	}
}

func (mf *CljsCoreMultiFn) X_prefer_method_Arity3(dispatch_val_x interface{}, dispatch_val_y interface{}) interface{} {
	if Truth_(Prefers_STAR_.X_invoke_Arity3(dispatch_val_x, dispatch_val_y, mf.Prefer_table)) {
		panic((&js.Error{("Preference conflict in multimethod '" + Str.X_invoke_Arity1(mf.Name).(string) + "': " + Str.X_invoke_Arity1(dispatch_val_y).(string) + " is already preferred to " + Str.X_invoke_Arity1(dispatch_val_x).(string))}))
	} else {
	}
	Swap_BANG_.X_invoke_Arity2(mf.Prefer_table, func(G__7960 *AFn) *AFn {
		return Fn(G__7960, 1, func(old interface{}) interface{} {
			return Assoc.X_invoke_Arity3(old, dispatch_val_x, Conj.X_invoke_Arity2(Get.X_invoke_Arity3(old, dispatch_val_x, CljsCorePersistentHashSet_EMPTY), dispatch_val_y))
		})
	}(&AFn{}))
	return Reset_cache.X_invoke_Arity4(mf.Method_cache, mf.Method_table, mf.Cached_hierarchy, mf.Hierarchy)
}

func (mf *CljsCoreMultiFn) X_methods_Arity1() interface{} {
	return Deref.X_invoke_Arity1(mf.Method_table)
}

func (mf *CljsCoreMultiFn) X_prefers_Arity1() interface{} {
	return Deref.X_invoke_Arity1(mf.Prefer_table)
}

func (_ *CljsCoreMultiFn) CljsCoreIFn__() {}

func (this *CljsCoreMultiFn) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (mf *CljsCoreMultiFn) X_invoke_Arity1(a interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7050 = a
			_ = G__7050
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity1(G__7050)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7051 = a
			_ = G__7051
			return target_fn.(CljsCoreIFn).X_invoke_Arity1(G__7051)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity2(a interface{}, b interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7056 = a
			var G__7057 = b
			_, _ = G__7056, G__7057
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity2(G__7056, G__7057)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7058 = a
			var G__7059 = b
			_, _ = G__7058, G__7059
			return target_fn.(CljsCoreIFn).X_invoke_Arity2(G__7058, G__7059)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7066 = a
			var G__7067 = b
			var G__7068 = c
			_, _, _ = G__7066, G__7067, G__7068
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity3(G__7066, G__7067, G__7068)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7069 = a
			var G__7070 = b
			var G__7071 = c
			_, _, _ = G__7069, G__7070, G__7071
			return target_fn.(CljsCoreIFn).X_invoke_Arity3(G__7069, G__7070, G__7071)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7080 = a
			var G__7081 = b
			var G__7082 = c
			var G__7083 = d
			_, _, _, _ = G__7080, G__7081, G__7082, G__7083
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity4(G__7080, G__7081, G__7082, G__7083)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7084 = a
			var G__7085 = b
			var G__7086 = c
			var G__7087 = d
			_, _, _, _ = G__7084, G__7085, G__7086, G__7087
			return target_fn.(CljsCoreIFn).X_invoke_Arity4(G__7084, G__7085, G__7086, G__7087)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7098 = a
			var G__7099 = b
			var G__7100 = c
			var G__7101 = d
			var G__7102 = e
			_, _, _, _, _ = G__7098, G__7099, G__7100, G__7101, G__7102
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity5(G__7098, G__7099, G__7100, G__7101, G__7102)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7103 = a
			var G__7104 = b
			var G__7105 = c
			var G__7106 = d
			var G__7107 = e
			_, _, _, _, _ = G__7103, G__7104, G__7105, G__7106, G__7107
			return target_fn.(CljsCoreIFn).X_invoke_Arity5(G__7103, G__7104, G__7105, G__7106, G__7107)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7120 = a
			var G__7121 = b
			var G__7122 = c
			var G__7123 = d
			var G__7124 = e
			var G__7125 = f
			_, _, _, _, _, _ = G__7120, G__7121, G__7122, G__7123, G__7124, G__7125
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity6(G__7120, G__7121, G__7122, G__7123, G__7124, G__7125)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7126 = a
			var G__7127 = b
			var G__7128 = c
			var G__7129 = d
			var G__7130 = e
			var G__7131 = f
			_, _, _, _, _, _ = G__7126, G__7127, G__7128, G__7129, G__7130, G__7131
			return target_fn.(CljsCoreIFn).X_invoke_Arity6(G__7126, G__7127, G__7128, G__7129, G__7130, G__7131)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7146 = a
			var G__7147 = b
			var G__7148 = c
			var G__7149 = d
			var G__7150 = e
			var G__7151 = f
			var G__7152 = g
			_, _, _, _, _, _, _ = G__7146, G__7147, G__7148, G__7149, G__7150, G__7151, G__7152
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity7(G__7146, G__7147, G__7148, G__7149, G__7150, G__7151, G__7152)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7153 = a
			var G__7154 = b
			var G__7155 = c
			var G__7156 = d
			var G__7157 = e
			var G__7158 = f
			var G__7159 = g
			_, _, _, _, _, _, _ = G__7153, G__7154, G__7155, G__7156, G__7157, G__7158, G__7159
			return target_fn.(CljsCoreIFn).X_invoke_Arity7(G__7153, G__7154, G__7155, G__7156, G__7157, G__7158, G__7159)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7176 = a
			var G__7177 = b
			var G__7178 = c
			var G__7179 = d
			var G__7180 = e
			var G__7181 = f
			var G__7182 = g
			var G__7183 = h
			_, _, _, _, _, _, _, _ = G__7176, G__7177, G__7178, G__7179, G__7180, G__7181, G__7182, G__7183
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity8(G__7176, G__7177, G__7178, G__7179, G__7180, G__7181, G__7182, G__7183)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7184 = a
			var G__7185 = b
			var G__7186 = c
			var G__7187 = d
			var G__7188 = e
			var G__7189 = f
			var G__7190 = g
			var G__7191 = h
			_, _, _, _, _, _, _, _ = G__7184, G__7185, G__7186, G__7187, G__7188, G__7189, G__7190, G__7191
			return target_fn.(CljsCoreIFn).X_invoke_Arity8(G__7184, G__7185, G__7186, G__7187, G__7188, G__7189, G__7190, G__7191)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7210 = a
			var G__7211 = b
			var G__7212 = c
			var G__7213 = d
			var G__7214 = e
			var G__7215 = f
			var G__7216 = g
			var G__7217 = h
			var G__7218 = i
			_, _, _, _, _, _, _, _, _ = G__7210, G__7211, G__7212, G__7213, G__7214, G__7215, G__7216, G__7217, G__7218
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity9(G__7210, G__7211, G__7212, G__7213, G__7214, G__7215, G__7216, G__7217, G__7218)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7219 = a
			var G__7220 = b
			var G__7221 = c
			var G__7222 = d
			var G__7223 = e
			var G__7224 = f
			var G__7225 = g
			var G__7226 = h
			var G__7227 = i
			_, _, _, _, _, _, _, _, _ = G__7219, G__7220, G__7221, G__7222, G__7223, G__7224, G__7225, G__7226, G__7227
			return target_fn.(CljsCoreIFn).X_invoke_Arity9(G__7219, G__7220, G__7221, G__7222, G__7223, G__7224, G__7225, G__7226, G__7227)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7248 = a
			var G__7249 = b
			var G__7250 = c
			var G__7251 = d
			var G__7252 = e
			var G__7253 = f
			var G__7254 = g
			var G__7255 = h
			var G__7256 = i
			var G__7257 = j
			_, _, _, _, _, _, _, _, _, _ = G__7248, G__7249, G__7250, G__7251, G__7252, G__7253, G__7254, G__7255, G__7256, G__7257
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity10(G__7248, G__7249, G__7250, G__7251, G__7252, G__7253, G__7254, G__7255, G__7256, G__7257)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7258 = a
			var G__7259 = b
			var G__7260 = c
			var G__7261 = d
			var G__7262 = e
			var G__7263 = f
			var G__7264 = g
			var G__7265 = h
			var G__7266 = i
			var G__7267 = j
			_, _, _, _, _, _, _, _, _, _ = G__7258, G__7259, G__7260, G__7261, G__7262, G__7263, G__7264, G__7265, G__7266, G__7267
			return target_fn.(CljsCoreIFn).X_invoke_Arity10(G__7258, G__7259, G__7260, G__7261, G__7262, G__7263, G__7264, G__7265, G__7266, G__7267)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7290 = a
			var G__7291 = b
			var G__7292 = c
			var G__7293 = d
			var G__7294 = e
			var G__7295 = f
			var G__7296 = g
			var G__7297 = h
			var G__7298 = i
			var G__7299 = j
			var G__7300 = k
			_, _, _, _, _, _, _, _, _, _, _ = G__7290, G__7291, G__7292, G__7293, G__7294, G__7295, G__7296, G__7297, G__7298, G__7299, G__7300
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity11(G__7290, G__7291, G__7292, G__7293, G__7294, G__7295, G__7296, G__7297, G__7298, G__7299, G__7300)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7301 = a
			var G__7302 = b
			var G__7303 = c
			var G__7304 = d
			var G__7305 = e
			var G__7306 = f
			var G__7307 = g
			var G__7308 = h
			var G__7309 = i
			var G__7310 = j
			var G__7311 = k
			_, _, _, _, _, _, _, _, _, _, _ = G__7301, G__7302, G__7303, G__7304, G__7305, G__7306, G__7307, G__7308, G__7309, G__7310, G__7311
			return target_fn.(CljsCoreIFn).X_invoke_Arity11(G__7301, G__7302, G__7303, G__7304, G__7305, G__7306, G__7307, G__7308, G__7309, G__7310, G__7311)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7336 = a
			var G__7337 = b
			var G__7338 = c
			var G__7339 = d
			var G__7340 = e
			var G__7341 = f
			var G__7342 = g
			var G__7343 = h
			var G__7344 = i
			var G__7345 = j
			var G__7346 = k
			var G__7347 = l
			_, _, _, _, _, _, _, _, _, _, _, _ = G__7336, G__7337, G__7338, G__7339, G__7340, G__7341, G__7342, G__7343, G__7344, G__7345, G__7346, G__7347
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity12(G__7336, G__7337, G__7338, G__7339, G__7340, G__7341, G__7342, G__7343, G__7344, G__7345, G__7346, G__7347)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7348 = a
			var G__7349 = b
			var G__7350 = c
			var G__7351 = d
			var G__7352 = e
			var G__7353 = f
			var G__7354 = g
			var G__7355 = h
			var G__7356 = i
			var G__7357 = j
			var G__7358 = k
			var G__7359 = l
			_, _, _, _, _, _, _, _, _, _, _, _ = G__7348, G__7349, G__7350, G__7351, G__7352, G__7353, G__7354, G__7355, G__7356, G__7357, G__7358, G__7359
			return target_fn.(CljsCoreIFn).X_invoke_Arity12(G__7348, G__7349, G__7350, G__7351, G__7352, G__7353, G__7354, G__7355, G__7356, G__7357, G__7358, G__7359)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7386 = a
			var G__7387 = b
			var G__7388 = c
			var G__7389 = d
			var G__7390 = e
			var G__7391 = f
			var G__7392 = g
			var G__7393 = h
			var G__7394 = i
			var G__7395 = j
			var G__7396 = k
			var G__7397 = l
			var G__7398 = m
			_, _, _, _, _, _, _, _, _, _, _, _, _ = G__7386, G__7387, G__7388, G__7389, G__7390, G__7391, G__7392, G__7393, G__7394, G__7395, G__7396, G__7397, G__7398
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity13(G__7386, G__7387, G__7388, G__7389, G__7390, G__7391, G__7392, G__7393, G__7394, G__7395, G__7396, G__7397, G__7398)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7399 = a
			var G__7400 = b
			var G__7401 = c
			var G__7402 = d
			var G__7403 = e
			var G__7404 = f
			var G__7405 = g
			var G__7406 = h
			var G__7407 = i
			var G__7408 = j
			var G__7409 = k
			var G__7410 = l
			var G__7411 = m
			_, _, _, _, _, _, _, _, _, _, _, _, _ = G__7399, G__7400, G__7401, G__7402, G__7403, G__7404, G__7405, G__7406, G__7407, G__7408, G__7409, G__7410, G__7411
			return target_fn.(CljsCoreIFn).X_invoke_Arity13(G__7399, G__7400, G__7401, G__7402, G__7403, G__7404, G__7405, G__7406, G__7407, G__7408, G__7409, G__7410, G__7411)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7440 = a
			var G__7441 = b
			var G__7442 = c
			var G__7443 = d
			var G__7444 = e
			var G__7445 = f
			var G__7446 = g
			var G__7447 = h
			var G__7448 = i
			var G__7449 = j
			var G__7450 = k
			var G__7451 = l
			var G__7452 = m
			var G__7453 = n
			_, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7440, G__7441, G__7442, G__7443, G__7444, G__7445, G__7446, G__7447, G__7448, G__7449, G__7450, G__7451, G__7452, G__7453
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity14(G__7440, G__7441, G__7442, G__7443, G__7444, G__7445, G__7446, G__7447, G__7448, G__7449, G__7450, G__7451, G__7452, G__7453)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7454 = a
			var G__7455 = b
			var G__7456 = c
			var G__7457 = d
			var G__7458 = e
			var G__7459 = f
			var G__7460 = g
			var G__7461 = h
			var G__7462 = i
			var G__7463 = j
			var G__7464 = k
			var G__7465 = l
			var G__7466 = m
			var G__7467 = n
			_, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7454, G__7455, G__7456, G__7457, G__7458, G__7459, G__7460, G__7461, G__7462, G__7463, G__7464, G__7465, G__7466, G__7467
			return target_fn.(CljsCoreIFn).X_invoke_Arity14(G__7454, G__7455, G__7456, G__7457, G__7458, G__7459, G__7460, G__7461, G__7462, G__7463, G__7464, G__7465, G__7466, G__7467)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7498 = a
			var G__7499 = b
			var G__7500 = c
			var G__7501 = d
			var G__7502 = e
			var G__7503 = f
			var G__7504 = g
			var G__7505 = h
			var G__7506 = i
			var G__7507 = j
			var G__7508 = k
			var G__7509 = l
			var G__7510 = m
			var G__7511 = n
			var G__7512 = o
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7498, G__7499, G__7500, G__7501, G__7502, G__7503, G__7504, G__7505, G__7506, G__7507, G__7508, G__7509, G__7510, G__7511, G__7512
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity15(G__7498, G__7499, G__7500, G__7501, G__7502, G__7503, G__7504, G__7505, G__7506, G__7507, G__7508, G__7509, G__7510, G__7511, G__7512)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7513 = a
			var G__7514 = b
			var G__7515 = c
			var G__7516 = d
			var G__7517 = e
			var G__7518 = f
			var G__7519 = g
			var G__7520 = h
			var G__7521 = i
			var G__7522 = j
			var G__7523 = k
			var G__7524 = l
			var G__7525 = m
			var G__7526 = n
			var G__7527 = o
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7513, G__7514, G__7515, G__7516, G__7517, G__7518, G__7519, G__7520, G__7521, G__7522, G__7523, G__7524, G__7525, G__7526, G__7527
			return target_fn.(CljsCoreIFn).X_invoke_Arity15(G__7513, G__7514, G__7515, G__7516, G__7517, G__7518, G__7519, G__7520, G__7521, G__7522, G__7523, G__7524, G__7525, G__7526, G__7527)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7560 = a
			var G__7561 = b
			var G__7562 = c
			var G__7563 = d
			var G__7564 = e
			var G__7565 = f
			var G__7566 = g
			var G__7567 = h
			var G__7568 = i
			var G__7569 = j
			var G__7570 = k
			var G__7571 = l
			var G__7572 = m
			var G__7573 = n
			var G__7574 = o
			var G__7575 = p
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7560, G__7561, G__7562, G__7563, G__7564, G__7565, G__7566, G__7567, G__7568, G__7569, G__7570, G__7571, G__7572, G__7573, G__7574, G__7575
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity16(G__7560, G__7561, G__7562, G__7563, G__7564, G__7565, G__7566, G__7567, G__7568, G__7569, G__7570, G__7571, G__7572, G__7573, G__7574, G__7575)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7576 = a
			var G__7577 = b
			var G__7578 = c
			var G__7579 = d
			var G__7580 = e
			var G__7581 = f
			var G__7582 = g
			var G__7583 = h
			var G__7584 = i
			var G__7585 = j
			var G__7586 = k
			var G__7587 = l
			var G__7588 = m
			var G__7589 = n
			var G__7590 = o
			var G__7591 = p
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7576, G__7577, G__7578, G__7579, G__7580, G__7581, G__7582, G__7583, G__7584, G__7585, G__7586, G__7587, G__7588, G__7589, G__7590, G__7591
			return target_fn.(CljsCoreIFn).X_invoke_Arity16(G__7576, G__7577, G__7578, G__7579, G__7580, G__7581, G__7582, G__7583, G__7584, G__7585, G__7586, G__7587, G__7588, G__7589, G__7590, G__7591)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7626 = a
			var G__7627 = b
			var G__7628 = c
			var G__7629 = d
			var G__7630 = e
			var G__7631 = f
			var G__7632 = g
			var G__7633 = h
			var G__7634 = i
			var G__7635 = j
			var G__7636 = k
			var G__7637 = l
			var G__7638 = m
			var G__7639 = n
			var G__7640 = o
			var G__7641 = p
			var G__7642 = q
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7626, G__7627, G__7628, G__7629, G__7630, G__7631, G__7632, G__7633, G__7634, G__7635, G__7636, G__7637, G__7638, G__7639, G__7640, G__7641, G__7642
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity17(G__7626, G__7627, G__7628, G__7629, G__7630, G__7631, G__7632, G__7633, G__7634, G__7635, G__7636, G__7637, G__7638, G__7639, G__7640, G__7641, G__7642)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7643 = a
			var G__7644 = b
			var G__7645 = c
			var G__7646 = d
			var G__7647 = e
			var G__7648 = f
			var G__7649 = g
			var G__7650 = h
			var G__7651 = i
			var G__7652 = j
			var G__7653 = k
			var G__7654 = l
			var G__7655 = m
			var G__7656 = n
			var G__7657 = o
			var G__7658 = p
			var G__7659 = q
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7643, G__7644, G__7645, G__7646, G__7647, G__7648, G__7649, G__7650, G__7651, G__7652, G__7653, G__7654, G__7655, G__7656, G__7657, G__7658, G__7659
			return target_fn.(CljsCoreIFn).X_invoke_Arity17(G__7643, G__7644, G__7645, G__7646, G__7647, G__7648, G__7649, G__7650, G__7651, G__7652, G__7653, G__7654, G__7655, G__7656, G__7657, G__7658, G__7659)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7696 = a
			var G__7697 = b
			var G__7698 = c
			var G__7699 = d
			var G__7700 = e
			var G__7701 = f
			var G__7702 = g
			var G__7703 = h
			var G__7704 = i
			var G__7705 = j
			var G__7706 = k
			var G__7707 = l
			var G__7708 = m
			var G__7709 = n
			var G__7710 = o
			var G__7711 = p
			var G__7712 = q
			var G__7713 = r
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7696, G__7697, G__7698, G__7699, G__7700, G__7701, G__7702, G__7703, G__7704, G__7705, G__7706, G__7707, G__7708, G__7709, G__7710, G__7711, G__7712, G__7713
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity18(G__7696, G__7697, G__7698, G__7699, G__7700, G__7701, G__7702, G__7703, G__7704, G__7705, G__7706, G__7707, G__7708, G__7709, G__7710, G__7711, G__7712, G__7713)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7714 = a
			var G__7715 = b
			var G__7716 = c
			var G__7717 = d
			var G__7718 = e
			var G__7719 = f
			var G__7720 = g
			var G__7721 = h
			var G__7722 = i
			var G__7723 = j
			var G__7724 = k
			var G__7725 = l
			var G__7726 = m
			var G__7727 = n
			var G__7728 = o
			var G__7729 = p
			var G__7730 = q
			var G__7731 = r
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7714, G__7715, G__7716, G__7717, G__7718, G__7719, G__7720, G__7721, G__7722, G__7723, G__7724, G__7725, G__7726, G__7727, G__7728, G__7729, G__7730, G__7731
			return target_fn.(CljsCoreIFn).X_invoke_Arity18(G__7714, G__7715, G__7716, G__7717, G__7718, G__7719, G__7720, G__7721, G__7722, G__7723, G__7724, G__7725, G__7726, G__7727, G__7728, G__7729, G__7730, G__7731)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7770 = a
			var G__7771 = b
			var G__7772 = c
			var G__7773 = d
			var G__7774 = e
			var G__7775 = f
			var G__7776 = g
			var G__7777 = h
			var G__7778 = i
			var G__7779 = j
			var G__7780 = k
			var G__7781 = l
			var G__7782 = m
			var G__7783 = n
			var G__7784 = o
			var G__7785 = p
			var G__7786 = q
			var G__7787 = r
			var G__7788 = s
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7770, G__7771, G__7772, G__7773, G__7774, G__7775, G__7776, G__7777, G__7778, G__7779, G__7780, G__7781, G__7782, G__7783, G__7784, G__7785, G__7786, G__7787, G__7788
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity19(G__7770, G__7771, G__7772, G__7773, G__7774, G__7775, G__7776, G__7777, G__7778, G__7779, G__7780, G__7781, G__7782, G__7783, G__7784, G__7785, G__7786, G__7787, G__7788)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7789 = a
			var G__7790 = b
			var G__7791 = c
			var G__7792 = d
			var G__7793 = e
			var G__7794 = f
			var G__7795 = g
			var G__7796 = h
			var G__7797 = i
			var G__7798 = j
			var G__7799 = k
			var G__7800 = l
			var G__7801 = m
			var G__7802 = n
			var G__7803 = o
			var G__7804 = p
			var G__7805 = q
			var G__7806 = r
			var G__7807 = s
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7789, G__7790, G__7791, G__7792, G__7793, G__7794, G__7795, G__7796, G__7797, G__7798, G__7799, G__7800, G__7801, G__7802, G__7803, G__7804, G__7805, G__7806, G__7807
			return target_fn.(CljsCoreIFn).X_invoke_Arity19(G__7789, G__7790, G__7791, G__7792, G__7793, G__7794, G__7795, G__7796, G__7797, G__7798, G__7799, G__7800, G__7801, G__7802, G__7803, G__7804, G__7805, G__7806, G__7807)
		}
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	{
		var dispatch_val = func() interface{} {
			var G__7848 = a
			var G__7849 = b
			var G__7850 = c
			var G__7851 = d
			var G__7852 = e
			var G__7853 = f
			var G__7854 = g
			var G__7855 = h
			var G__7856 = i
			var G__7857 = j
			var G__7858 = k
			var G__7859 = l
			var G__7860 = m
			var G__7861 = n
			var G__7862 = o
			var G__7863 = p
			var G__7864 = q
			var G__7865 = r
			var G__7866 = s
			var G__7867 = t
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7848, G__7849, G__7850, G__7851, G__7852, G__7853, G__7854, G__7855, G__7856, G__7857, G__7858, G__7859, G__7860, G__7861, G__7862, G__7863, G__7864, G__7865, G__7866, G__7867
			return mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity20(G__7848, G__7849, G__7850, G__7851, G__7852, G__7853, G__7854, G__7855, G__7856, G__7857, G__7858, G__7859, G__7860, G__7861, G__7862, G__7863, G__7864, G__7865, G__7866, G__7867)
		}()
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		{
			var G__7868 = a
			var G__7869 = b
			var G__7870 = c
			var G__7871 = d
			var G__7872 = e
			var G__7873 = f
			var G__7874 = g
			var G__7875 = h
			var G__7876 = i
			var G__7877 = j
			var G__7878 = k
			var G__7879 = l
			var G__7880 = m
			var G__7881 = n
			var G__7882 = o
			var G__7883 = p
			var G__7884 = q
			var G__7885 = r
			var G__7886 = s
			var G__7887 = t
			_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = G__7868, G__7869, G__7870, G__7871, G__7872, G__7873, G__7874, G__7875, G__7876, G__7877, G__7878, G__7879, G__7880, G__7881, G__7882, G__7883, G__7884, G__7885, G__7886, G__7887
			return target_fn.(CljsCoreIFn).X_invoke_Arity20(G__7868, G__7869, G__7870, G__7871, G__7872, G__7873, G__7874, G__7875, G__7876, G__7877, G__7878, G__7879, G__7880, G__7881, G__7882, G__7883, G__7884, G__7885, G__7886, G__7887)
		}
	}
}

var X__GT_MultiFn *AFn

// Removes all of the methods of multimethod.
var Remove_all_methods *AFn

// Removes the method of multimethod associated with dispatch-value.
var Remove_method *AFn

// Causes the multimethod to prefer matches of dispatch-val-x over dispatch-val-y
// when there is a conflict
var Prefer_method *AFn

// Given a multimethod, returns a map of dispatch values -> dispatch fns
var Methods *AFn

// Given a multimethod and a dispatch value, returns the dispatch fn
// that would apply to that value, or nil if none apply and no default
var Get_method *AFn

// Given a multimethod, returns a map of preferred value -> set of other values
var Prefers *AFn

type CljsCoreUUID struct{ Uuid interface{} }

func (_ *CljsCoreUUID) CljsCoreIHash__() {}

func (this *CljsCoreUUID) X_hash_Arity1() interface{} {
	{
		var G__7889 = Pr_str.X_invoke_ArityVariadic(Array_seq.X_invoke_Arity1([]interface{}{this})).(string)
		_ = G__7889
		return Native_invoke_func.X_invoke_Arity2(goog_string.HashCode, []interface{}{G__7889})
	}
}

func (_ *CljsCoreUUID) CljsCoreIPrintWithWriter__() {}

func (___ *CljsCoreUUID) X_pr_writer_Arity3(writer interface{}, ______1 interface{}) interface{} {
	return writer.(CljsCoreIWriter).X_write_Arity2(("#uuid \"" + Str.X_invoke_Arity1(___.Uuid).(string) + "\""))
}

func (_ *CljsCoreUUID) CljsCoreIEquiv__() {}

func (___ *CljsCoreUUID) X_equiv_Arity2(other interface{}) bool {
	return (Value_(other).Type().AssignableTo(reflect.TypeOf((**CljsCoreUUID)(nil)).Elem())) && (reflect.DeepEqual(___.Uuid, Native_get_instance_field.X_invoke_Arity2(other, "Uuid")))
}

func (_ *CljsCoreUUID) CljsCoreObject__() {}

func (___ *CljsCoreUUID) ToString() string {
	return ___.Uuid.(string)
}

func (this *CljsCoreUUID) String() string {
	return this.ToString()
}

func (this *CljsCoreUUID) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

var X__GT_UUID *AFn

type CljsCoreExceptionInfo struct {
	Message interface{}
	Data    interface{}
	Cause   interface{}
}

var X__GT_ExceptionInfo *AFn

// Alpha - subject to change.
// Create an instance of ExceptionInfo, an Error type that carries a
// map of additional data.
var Ex_info *AFn

// Alpha - subject to change.
// Returns exception data (a map) if ex is an ExceptionInfo.
// Otherwise returns nil.
var Ex_data *AFn

// Alpha - subject to change.
// Returns the message attached to the given Error / ExceptionInfo object.
// For non-Errors returns nil.
var Ex_message *AFn

// Alpha - subject to change.
// Returns exception cause (an Error / ExceptionInfo) if ex is an
// ExceptionInfo.
// Otherwise returns nil.
var Ex_cause *AFn

// Returns an JavaScript compatible comparator based upon pred.
var Comparator *AFn

var Special_symbol_QMARK_ *AFn
