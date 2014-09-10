// Compiled by ClojureScript to Go 0.0-2322
// cljs.core

package core

import (
	"fmt"
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
		return Fn(_STAR_print_fn_STAR_, func(___ interface{}) interface{} {
			panic((&js.Error{"No *print-fn* fn set for evaluation environment"}))
		})
	}(&AFn{})

	X_STAR_flush_on_newline_STAR_ = true

	X_STAR_print_newline_STAR_ = true

	X_STAR_print_readably_STAR_ = true

	X_STAR_print_meta_STAR_ = false

	X_STAR_print_dup_STAR_ = false

	Pr_opts = func(pr_opts *AFn) *AFn {
		return Fn(pr_opts, func() interface{} {
			return (&CljsCorePersistentArrayMap{nil, float64(5), []interface{}{(&CljsCoreKeyword{Ns: nil, Name: "flush-on-newline", Fqn: "flush-on-newline", X_hash: float64(-151457939)}), X_STAR_flush_on_newline_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), X_STAR_print_readably_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}), X_STAR_print_meta_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "dup", Fqn: "dup", X_hash: float64(556298533)}), X_STAR_print_dup_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "print-length", Fqn: "print-length", X_hash: float64(1931866356)}), X_STAR_print_length_STAR_}, nil})
		})
	}(&AFn{})

	Not_native = nil

	Identical_QMARK_ = func(identical_QMARK_ *AFn) *AFn {
		return Fn(identical_QMARK_, func(x interface{}, y interface{}) bool {
			return reflect.DeepEqual(x, y)
		})
	}(&AFn{})

	Nil_QMARK_ = func(nil_QMARK_ *AFn) *AFn {
		return Fn(nil_QMARK_, func(x interface{}) bool {
			return reflect.DeepEqual(x, nil)
		})
	}(&AFn{})

	Array_QMARK_ = func(array_QMARK_ *AFn) *AFn {
		return Fn(array_QMARK_, func(x interface{}) bool {
			return reflect.ValueOf(x).Kind() == reflect.Slice
		})
	}(&AFn{})

	Number_QMARK_ = func(number_QMARK_ *AFn) *AFn {
		return Fn(number_QMARK_, func(n interface{}) bool {
			return reflect.ValueOf(n).Kind() == reflect.Float64
		})
	}(&AFn{})

	Not = func(not *AFn) *AFn {
		return Fn(not, func(x interface{}) bool {
			if Truth_(x) {
				return false
			} else {
				return true
			}
		})
	}(&AFn{})

	Some_QMARK_ = func(some_QMARK_ *AFn) *AFn {
		return Fn(some_QMARK_, func(x interface{}) bool {
			return !(Nil_(x))
		})
	}(&AFn{})

	String_QMARK_ = func(string_QMARK_ *AFn) *AFn {
		return Fn(string_QMARK_, func(x interface{}) bool {
			return Truth_(Native_invoke_func.X_invoke_Arity2(goog.IsString, []interface{}{x}))
		})
	}(&AFn{})

	X_STAR_main_cli_fn_STAR_ = nil

	Aclone = func(aclone *AFn) *AFn {
		return Fn(aclone, func(arr interface{}) interface{} {
			{
				var len = float64(len(arr.([]interface{})))
				var new_arr = make([]interface{}, int(len))
				_, _ = len, new_arr
				{
					var n__1042__auto___287 = len
					_ = n__1042__auto___287
					{
						var i_288 = float64(0)
						_ = i_288
						for {
							if i_288 < n__1042__auto___287 {
								new_arr[int(i_288)] = (arr.([]interface{})[int(i_288)])
								i_288 = (i_288 + float64(1))
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
		return Fn(aget, func(array interface{}, i interface{}) interface{} {
			return (array.([]interface{})[int(i.(float64))])
		}, func(array_i_idxs__ ...interface{}) interface{} {
			var array = array_i_idxs__[0]
			var i = array_i_idxs__[1]
			var idxs = Array_seq.X_invoke_Arity1(array_i_idxs__[2:])
			_, _, _ = array, i, idxs
			return Apply.X_invoke_Arity3(aget, aget.X_invoke_Arity2(array, i), idxs)
		})
	}(&AFn{})

	Aset = func(aset *AFn) *AFn {
		return Fn(aset, func(array interface{}, i interface{}, val interface{}) interface{} {
			return func() interface{} {
				array.([]interface{})[int(i.(float64))] = val
				return array.([]interface{})[int(i.(float64))]
			}()
		}, func(array_idx_idx2_idxv__ ...interface{}) interface{} {
			var array = array_idx_idx2_idxv__[0]
			var idx = array_idx_idx2_idxv__[1]
			var idx2 = array_idx_idx2_idxv__[2]
			var idxv = Array_seq.X_invoke_Arity1(array_idx_idx2_idxv__[3:])
			_, _, _, _ = array, idx, idx2, idxv
			return Apply.X_invoke_Arity4(aset, (array.([]interface{})[int(idx.(float64))]), idx2, idxv)
		})
	}(&AFn{})

	Alength = func(alength *AFn) *AFn {
		return Fn(alength, func(array interface{}) float64 {
			return float64(len(array.([]interface{})))
		})
	}(&AFn{})

	Into_array = func(into_array *AFn) *AFn {
		return Fn(into_array, func(aseq interface{}) []interface{} {
			return into_array.Arity2IIA(nil, aseq)
		}, func(type_ interface{}, aseq interface{}) []interface{} {
			return Reduce.X_invoke_Arity3(func(G__289 *AFn) *AFn {
				return Fn(G__289, func(a interface{}, x interface{}) interface{} {
					js.JSArray_(&a).Push(x)
					return a
				})
			}(&AFn{}), []interface{}{}, aseq).([]interface{})
		})
	}(&AFn{})

	X_invoke = func(_invoke *AFn) *AFn {
		return Fn(_invoke, func(this interface{}) interface{} {
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
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity18(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity19(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s, t)
		}, func(this interface{}, a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
			return this.(CljsCoreIFn).X_invoke_Arity20(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, s, t, rest)
		})
	}(&AFn{})

	X_clone = func(_clone *AFn) *AFn {
		return Fn(_clone, func(value interface{}) interface{} {
			return value.(CljsCoreICloneable).X_clone_Arity1()
		})
	}(&AFn{})

	X_count = func(_count *AFn) *AFn {
		return Fn(_count, func(coll interface{}) float64 {
			return coll.(CljsCoreICounted).X_count_Arity1()
		})
	}(&AFn{})

	X_empty = func(_empty *AFn) *AFn {
		return Fn(_empty, func(coll interface{}) interface{} {
			return coll.(CljsCoreIEmptyableCollection).X_empty_Arity1()
		})
	}(&AFn{})

	X_conj = func(_conj *AFn) *AFn {
		return Fn(_conj, func(coll interface{}, o interface{}) interface{} {
			return coll.(CljsCoreICollection).X_conj_Arity2(o)
		})
	}(&AFn{})

	X_nth = func(_nth *AFn) *AFn {
		return Fn(_nth, func(coll interface{}, n interface{}) interface{} {
			return coll.(CljsCoreIIndexed).X_nth_Arity2(n)
		}, func(coll interface{}, n interface{}, not_found interface{}) interface{} {
			return coll.(CljsCoreIIndexed).X_nth_Arity3(n, not_found)
		})
	}(&AFn{})

	X_first = func(_first *AFn) *AFn {
		return Fn(_first, func(coll interface{}) interface{} {
			return coll.(CljsCoreISeq).X_first_Arity1()
		})
	}(&AFn{})

	X_rest = func(_rest *AFn) *AFn {
		return Fn(_rest, func(coll interface{}) interface{} {
			return coll.(CljsCoreISeq).X_rest_Arity1()
		})
	}(&AFn{})

	X_next = func(_next *AFn) *AFn {
		return Fn(_next, func(coll interface{}) interface{} {
			return coll.(CljsCoreINext).X_next_Arity1()
		})
	}(&AFn{})

	X_lookup = func(_lookup *AFn) *AFn {
		return Fn(_lookup, func(o interface{}, k interface{}) interface{} {
			return o.(CljsCoreILookup).X_lookup_Arity2(k)
		}, func(o interface{}, k interface{}, not_found interface{}) interface{} {
			return o.(CljsCoreILookup).X_lookup_Arity3(k, not_found)
		})
	}(&AFn{})

	X_contains_key_QMARK_ = func(_contains_key_QMARK_ *AFn) *AFn {
		return Fn(_contains_key_QMARK_, func(coll interface{}, k interface{}) bool {
			return coll.(CljsCoreIAssociative).X_contains_key_QMARK__Arity2(k)
		})
	}(&AFn{})

	X_assoc = func(_assoc *AFn) *AFn {
		return Fn(_assoc, func(coll interface{}, k interface{}, v interface{}) interface{} {
			return coll.(CljsCoreIAssociative).X_assoc_Arity3(k, v)
		})
	}(&AFn{})

	X_dissoc = func(_dissoc *AFn) *AFn {
		return Fn(_dissoc, func(coll interface{}, k interface{}) interface{} {
			return coll.(CljsCoreIMap).X_dissoc_Arity2(k)
		})
	}(&AFn{})

	X_key = func(_key *AFn) *AFn {
		return Fn(_key, func(coll interface{}) interface{} {
			return coll.(CljsCoreIMapEntry).X_key_Arity1()
		})
	}(&AFn{})

	X_val = func(_val *AFn) *AFn {
		return Fn(_val, func(coll interface{}) interface{} {
			return coll.(CljsCoreIMapEntry).X_val_Arity1()
		})
	}(&AFn{})

	X_disjoin = func(_disjoin *AFn) *AFn {
		return Fn(_disjoin, func(coll interface{}, v interface{}) interface{} {
			return coll.(CljsCoreISet).X_disjoin_Arity2(v)
		})
	}(&AFn{})

	X_peek = func(_peek *AFn) *AFn {
		return Fn(_peek, func(coll interface{}) interface{} {
			return coll.(CljsCoreIStack).X_peek_Arity1()
		})
	}(&AFn{})

	X_pop = func(_pop *AFn) *AFn {
		return Fn(_pop, func(coll interface{}) interface{} {
			return coll.(CljsCoreIStack).X_pop_Arity1()
		})
	}(&AFn{})

	X_assoc_n = func(_assoc_n *AFn) *AFn {
		return Fn(_assoc_n, func(coll interface{}, n interface{}, val interface{}) interface{} {
			return coll.(CljsCoreIVector).X_assoc_n_Arity3(n, val)
		})
	}(&AFn{})

	X_deref = func(_deref *AFn) *AFn {
		return Fn(_deref, func(o interface{}) interface{} {
			return o.(CljsCoreIDeref).X_deref_Arity1()
		})
	}(&AFn{})

	X_deref_with_timeout = func(_deref_with_timeout *AFn) *AFn {
		return Fn(_deref_with_timeout, func(o interface{}, msec interface{}, timeout_val interface{}) interface{} {
			return o.(CljsCoreIDerefWithTimeout).X_deref_with_timeout_Arity3(msec, timeout_val)
		})
	}(&AFn{})

	X_meta = func(_meta *AFn) *AFn {
		return Fn(_meta, func(o interface{}) interface{} {
			return o.(CljsCoreIMeta).X_meta_Arity1()
		})
	}(&AFn{})

	X_with_meta = func(_with_meta *AFn) *AFn {
		return Fn(_with_meta, func(o interface{}, meta interface{}) interface{} {
			return o.(CljsCoreIWithMeta).X_with_meta_Arity2(meta)
		})
	}(&AFn{})

	X_reduce = func(_reduce *AFn) *AFn {
		return Fn(_reduce, func(coll interface{}, f interface{}) interface{} {
			return coll.(CljsCoreIReduce).X_reduce_Arity2(f)
		}, func(coll interface{}, f interface{}, start interface{}) interface{} {
			return coll.(CljsCoreIReduce).X_reduce_Arity3(f, start)
		})
	}(&AFn{})

	X_kv_reduce = func(_kv_reduce *AFn) *AFn {
		return Fn(_kv_reduce, func(coll interface{}, f interface{}, init interface{}) interface{} {
			return coll.(CljsCoreIKVReduce).X_kv_reduce_Arity3(f, init)
		})
	}(&AFn{})

	X_equiv = func(_equiv *AFn) *AFn {
		return Fn(_equiv, func(o interface{}, other interface{}) bool {
			return o.(CljsCoreIEquiv).X_equiv_Arity2(other)
		})
	}(&AFn{})

	X_hash = func(_hash *AFn) *AFn {
		return Fn(_hash, func(o interface{}) interface{} {
			return o.(CljsCoreIHash).X_hash_Arity1()
		})
	}(&AFn{})

	X_seq = func(_seq *AFn) *AFn {
		return Fn(_seq, func(o interface{}) interface{} {
			return o.(CljsCoreISeqable).X_seq_Arity1()
		})
	}(&AFn{})

	X_rseq = func(_rseq *AFn) *AFn {
		return Fn(_rseq, func(coll interface{}) interface{} {
			return coll.(CljsCoreIReversible).X_rseq_Arity1()
		})
	}(&AFn{})

	X_sorted_seq = func(_sorted_seq *AFn) *AFn {
		return Fn(_sorted_seq, func(coll interface{}, ascending_QMARK_ interface{}) interface{} {
			return coll.(CljsCoreISorted).X_sorted_seq_Arity2(ascending_QMARK_)
		})
	}(&AFn{})

	X_sorted_seq_from = func(_sorted_seq_from *AFn) *AFn {
		return Fn(_sorted_seq_from, func(coll interface{}, k interface{}, ascending_QMARK_ interface{}) interface{} {
			return coll.(CljsCoreISorted).X_sorted_seq_from_Arity3(k, ascending_QMARK_)
		})
	}(&AFn{})

	X_entry_key = func(_entry_key *AFn) *AFn {
		return Fn(_entry_key, func(coll interface{}, entry interface{}) interface{} {
			return coll.(CljsCoreISorted).X_entry_key_Arity2(entry)
		})
	}(&AFn{})

	X_comparator = func(_comparator *AFn) *AFn {
		return Fn(_comparator, func(coll interface{}) interface{} {
			return coll.(CljsCoreISorted).X_comparator_Arity1()
		})
	}(&AFn{})

	X_write = func(_write *AFn) *AFn {
		return Fn(_write, func(writer interface{}, s interface{}) interface{} {
			return writer.(CljsCoreIWriter).X_write_Arity2(s)
		})
	}(&AFn{})

	X_flush = func(_flush *AFn) *AFn {
		return Fn(_flush, func(writer interface{}) interface{} {
			return writer.(CljsCoreIWriter).X_flush_Arity1()
		})
	}(&AFn{})

	X_pr_writer = func(_pr_writer *AFn) *AFn {
		return Fn(_pr_writer, func(o interface{}, writer interface{}, opts interface{}) interface{} {
			return o.(CljsCoreIPrintWithWriter).X_pr_writer_Arity3(writer, opts)
		})
	}(&AFn{})

	X_realized_QMARK_ = func(_realized_QMARK_ *AFn) *AFn {
		return Fn(_realized_QMARK_, func(d interface{}) bool {
			return d.(CljsCoreIPending).X_realized_QMARK__Arity1()
		})
	}(&AFn{})

	X_notify_watches = func(_notify_watches *AFn) *AFn {
		return Fn(_notify_watches, func(this interface{}, oldval interface{}, newval interface{}) interface{} {
			return this.(CljsCoreIWatchable).X_notify_watches_Arity3(oldval, newval)
		})
	}(&AFn{})

	X_add_watch = func(_add_watch *AFn) *AFn {
		return Fn(_add_watch, func(this interface{}, key interface{}, f interface{}) interface{} {
			return this.(CljsCoreIWatchable).X_add_watch_Arity3(key, f)
		})
	}(&AFn{})

	X_remove_watch = func(_remove_watch *AFn) *AFn {
		return Fn(_remove_watch, func(this interface{}, key interface{}) interface{} {
			return this.(CljsCoreIWatchable).X_remove_watch_Arity2(key)
		})
	}(&AFn{})

	X_as_transient = func(_as_transient *AFn) *AFn {
		return Fn(_as_transient, func(coll interface{}) interface{} {
			return coll.(CljsCoreIEditableCollection).X_as_transient_Arity1()
		})
	}(&AFn{})

	X_conj_BANG_ = func(_conj_BANG_ *AFn) *AFn {
		return Fn(_conj_BANG_, func(tcoll interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientCollection).X_conj_BANG__Arity2(val)
		})
	}(&AFn{})

	X_persistent_BANG_ = func(_persistent_BANG_ *AFn) *AFn {
		return Fn(_persistent_BANG_, func(tcoll interface{}) interface{} {
			return tcoll.(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
		})
	}(&AFn{})

	X_assoc_BANG_ = func(_assoc_BANG_ *AFn) *AFn {
		return Fn(_assoc_BANG_, func(tcoll interface{}, key interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3(key, val)
		})
	}(&AFn{})

	X_dissoc_BANG_ = func(_dissoc_BANG_ *AFn) *AFn {
		return Fn(_dissoc_BANG_, func(tcoll interface{}, key interface{}) interface{} {
			return tcoll.(CljsCoreITransientMap).X_dissoc_BANG__Arity2(key)
		})
	}(&AFn{})

	X_assoc_n_BANG_ = func(_assoc_n_BANG_ *AFn) *AFn {
		return Fn(_assoc_n_BANG_, func(tcoll interface{}, n interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientVector).X_assoc_n_BANG__Arity3(n, val)
		})
	}(&AFn{})

	X_pop_BANG_ = func(_pop_BANG_ *AFn) *AFn {
		return Fn(_pop_BANG_, func(tcoll interface{}) interface{} {
			return tcoll.(CljsCoreITransientVector).X_pop_BANG__Arity1()
		})
	}(&AFn{})

	X_disjoin_BANG_ = func(_disjoin_BANG_ *AFn) *AFn {
		return Fn(_disjoin_BANG_, func(tcoll interface{}, v interface{}) interface{} {
			return tcoll.(CljsCoreITransientSet).X_disjoin_BANG__Arity2(v)
		})
	}(&AFn{})

	X_compare = func(_compare *AFn) *AFn {
		return Fn(_compare, func(x interface{}, y interface{}) float64 {
			return x.(CljsCoreIComparable).X_compare_Arity2(y)
		})
	}(&AFn{})

	X_drop_first = func(_drop_first *AFn) *AFn {
		return Fn(_drop_first, func(coll interface{}) interface{} {
			return coll.(CljsCoreIChunk).X_drop_first_Arity1()
		})
	}(&AFn{})

	X_chunked_first = func(_chunked_first *AFn) *AFn {
		return Fn(_chunked_first, func(coll interface{}) interface{} {
			return coll.(CljsCoreIChunkedSeq).X_chunked_first_Arity1()
		})
	}(&AFn{})

	X_chunked_rest = func(_chunked_rest *AFn) *AFn {
		return Fn(_chunked_rest, func(coll interface{}) interface{} {
			return coll.(CljsCoreIChunkedSeq).X_chunked_rest_Arity1()
		})
	}(&AFn{})

	X_chunked_next = func(_chunked_next *AFn) *AFn {
		return Fn(_chunked_next, func(coll interface{}) interface{} {
			return coll.(CljsCoreIChunkedNext).X_chunked_next_Arity1()
		})
	}(&AFn{})

	X_name = func(_name *AFn) *AFn {
		return Fn(_name, func(x interface{}) string {
			return x.(CljsCoreINamed).X_name_Arity1()
		})
	}(&AFn{})

	X_namespace = func(_namespace *AFn) *AFn {
		return Fn(_namespace, func(x interface{}) string {
			return x.(CljsCoreINamed).X_namespace_Arity1()
		})
	}(&AFn{})

	X_reset_BANG_ = func(_reset_BANG_ *AFn) *AFn {
		return Fn(_reset_BANG_, func(o interface{}, new_value interface{}) interface{} {
			return o.(CljsCoreIReset).X_reset_BANG__Arity2(new_value)
		})
	}(&AFn{})

	X_swap_BANG_ = func(_swap_BANG_ *AFn) *AFn {
		return Fn(_swap_BANG_, func(o interface{}, f interface{}) interface{} {
			return o.(CljsCoreISwap).X_swap_BANG__Arity2(f)
		}, func(o interface{}, f interface{}, a interface{}) interface{} {
			return o.(CljsCoreISwap).X_swap_BANG__Arity3(f, a)
		}, func(o interface{}, f interface{}, a interface{}, b interface{}) interface{} {
			return o.(CljsCoreISwap).X_swap_BANG__Arity4(f, a, b)
		}, func(o interface{}, f interface{}, a interface{}, b interface{}, xs interface{}) interface{} {
			return o.(CljsCoreISwap).X_swap_BANG__Arity5(f, a, b, xs)
		})
	}(&AFn{})

	X__GT_StringBufferWriter = func(__GT_StringBufferWriter *AFn) *AFn {
		return Fn(__GT_StringBufferWriter, func(sb interface{}) interface{} {
			return (&CljsCoreStringBufferWriter{sb})
		})
	}(&AFn{})

	Pr_str_STAR_ = func(pr_str_STAR_ *AFn) *AFn {
		return Fn(pr_str_STAR_, func(obj interface{}) interface{} {
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
		return Fn(int_rotate_left, func(x interface{}, n interface{}) float64 {
			return float64(int(float64(int(x.(float64))<<uint(n.(float64)))) | int(float64(uint(x.(float64))>>uint((-n.(float64))))))
		})
	}(&AFn{})

	if (reflect.ValueOf(Math.Imul).Kind() != reflect.Invalid) && (!(Native_invoke_func.X_invoke_Arity2(Math.Imul, []interface{}{float64(4294967295), float64(5)}).(float64) == float64(0))) {
		Imul = func(imul *AFn) *AFn {
			return Fn(imul, func(a interface{}, b interface{}) float64 {
				return Native_invoke_func.X_invoke_Arity2(Math.Imul, []interface{}{a, b}).(float64)
			})
		}(&AFn{})

	} else {
		Imul = func(imul *AFn) *AFn {
			return Fn(imul, func(a interface{}, b interface{}) float64 {
				{
					var ah = float64(int(float64(uint(a.(float64))>>uint(float64(16)))) & int(float64(65535)))
					var al = float64(int(a.(float64)) & int(float64(65535)))
					var bh = float64(int(float64(uint(b.(float64))>>uint(float64(16)))) & int(float64(65535)))
					var bl = float64(int(b.(float64)) & int(float64(65535)))
					_, _, _, _ = ah, al, bh, bl
					return float64(int(((al * bl) + float64(uint(float64(int(((ah*bl)+(al*bh)))<<uint(float64(16))))>>uint(float64(0))))) | int(float64(0)))
				}
			})
		}(&AFn{})

	}
	M3_seed = float64(0)

	M3_C1 = float64(3432918353)

	M3_C2 = float64(461845907)

	M3_mix_K1 = func(m3_mix_K1 *AFn) *AFn {
		return Fn(m3_mix_K1, func(k1 interface{}) float64 {
			return Imul.Arity2IIF(Int_rotate_left.Arity2IIF(Imul.Arity2IIF(k1, M3_C1), float64(15)), M3_C2)
		})
	}(&AFn{})

	M3_mix_H1 = func(m3_mix_H1 *AFn) *AFn {
		return Fn(m3_mix_H1, func(h1 interface{}, k1 interface{}) float64 {
			return (Imul.Arity2IIF(Int_rotate_left.Arity2IIF(float64(int(h1.(float64))^int(k1.(float64))), float64(13)), float64(5)) + float64(3864292196))
		})
	}(&AFn{})

	M3_fmix = func(m3_fmix *AFn) *AFn {
		return Fn(m3_fmix, func(h1 interface{}, len interface{}) float64 {
			{
				var h1___1 = h1
				var h1___2 = float64(int(h1___1.(float64)) ^ int(len.(float64)))
				var h1___3 = float64(int(h1___2) ^ int(float64(uint(h1___2)>>uint(float64(16)))))
				var h1___4 = Imul.Arity2IIF(h1___3, float64(2246822507))
				var h1___5 = float64(int(h1___4) ^ int(float64(uint(h1___4)>>uint(float64(13)))))
				var h1___6 = Imul.Arity2IIF(h1___5, float64(3266489909))
				var h1___7 = float64(int(h1___6) ^ int(float64(uint(h1___6)>>uint(float64(16)))))
				_, _, _, _, _, _, _ = h1___1, h1___2, h1___3, h1___4, h1___5, h1___6, h1___7
				return h1___7
			}
		})
	}(&AFn{})

	M3_hash_int = func(m3_hash_int *AFn) *AFn {
		return Fn(m3_hash_int, func(in interface{}) float64 {
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
		return Fn(m3_hash_unencoded_chars, func(in interface{}) float64 {
			{
				var h1 = func() interface{} {
					var i = float64(1)
					var h1 = M3_seed
					_, _ = i, h1
					for {
						if i < float64(len(in.([]interface{}))) {
							i, h1 = (i + float64(2)), M3_mix_H1.Arity2IIF(h1, M3_mix_K1.Arity1IF(float64(int(Native_invoke_instance_method.X_invoke_Arity3(in, "CharCodeAt", []interface{}{(i - float64(1))}).(float64))|int(float64(int(Native_invoke_instance_method.X_invoke_Arity3(in, "CharCodeAt", []interface{}{i}).(float64))<<uint(float64(16)))))))
							continue
						} else {
							return h1
						}
					}
				}()
				var h1___1 = func() interface{} {
					if float64(int(float64(len(in.([]interface{}))))&int(float64(1))) == float64(1) {
						return float64(int(h1.(float64)) ^ int(M3_mix_K1.Arity1IF(Native_invoke_instance_method.X_invoke_Arity3(in, "CharCodeAt", []interface{}{(float64(len(in.([]interface{}))) - float64(1))}))))
					} else {
						return h1
					}
				}()
				_, _ = h1, h1___1
				return M3_fmix.Arity2IIF(h1___1, Imul.Arity2IIF(float64(2), float64(len(in.([]interface{})))))
			}
		})
	}(&AFn{})

	String_hash_cache_count = float64(0)

	Hash_string_STAR_ = func(hash_string_STAR_ *AFn) *AFn {
		return Fn(hash_string_STAR_, func(s interface{}) interface{} {
			if !(Nil_(s)) {
				{
					var len = float64(len(s.([]interface{})))
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
		return Fn(hash, func(o interface{}) interface{} {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IHash", Str: "cljs.core/IHash", X_hash: float64(-32453786), X_meta: nil}), o)) {
				return o.(CljsCoreIHash).X_hash_Arity1()
			} else {
				if reflect.ValueOf(o).Kind() == reflect.Float64 {
					return float64(int(Native_invoke_func.X_invoke_Arity2(Math.Floor, []interface{}{o}).(float64)) % int(float64(2147483647)))
				} else {
					if o == true {
						return float64(1)
					} else {
						if o == false {
							return float64(0)
						} else {
							if reflect.ValueOf(o).Kind() == reflect.String {
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
		return Fn(hash_combine, func(seed interface{}, hash interface{}) interface{} {
			return float64(int(seed.(float64)) ^ int((((hash.(float64) + float64(2654435769)) + float64(int(seed.(float64))<<uint(float64(6)))) + float64(int(seed.(float64))>>uint(float64(2))))))
		})
	}(&AFn{})

	Hash_symbol = func(hash_symbol *AFn) *AFn {
		return Fn(hash_symbol, func(sym interface{}) interface{} {
			return Hash_combine.X_invoke_Arity2(M3_hash_unencoded_chars.Arity1IF(Native_get_instance_field.X_invoke_Arity2(sym, "Name")), Hash_string.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(sym, "Ns"))).(float64)
		})
	}(&AFn{})

	Compare_symbols = func(compare_symbols *AFn) *AFn {
		return Fn(compare_symbols, func(a interface{}, b interface{}) interface{} {
			if X_EQ_.Arity2IIB(a, b) {
				return float64(0)
			} else {
				if Truth_(func() interface{} {
					var and__163__auto__ = Not.Arity1IB(Native_get_instance_field.X_invoke_Arity2(a, "Ns"))
					_ = and__163__auto__
					if Truth_(and__163__auto__) {
						return Native_get_instance_field.X_invoke_Arity2(b, "Ns")
					} else {
						return and__163__auto__
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
		return Fn(__GT_Symbol, func(ns interface{}, name interface{}, str interface{}, _hash interface{}, _meta interface{}) interface{} {
			return (&CljsCoreSymbol{ns, name, str, _hash, _meta})
		})
	}(&AFn{})

	Symbol = func(symbol *AFn) *AFn {
		return Fn(symbol, func(name interface{}) interface{} {
			if func() bool { _, instanceof := name.(*CljsCoreSymbol); return instanceof }() {
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

	Clone = func(clone *AFn) *AFn {
		return Fn(clone, func(value interface{}) interface{} {
			return value.(CljsCoreICloneable).X_clone_Arity1()
		})
	}(&AFn{})

	Cloneable_QMARK_ = func(cloneable_QMARK_ *AFn) *AFn {
		return Fn(cloneable_QMARK_, func(value interface{}) interface{} {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ICloneable", Str: "cljs.core/ICloneable", X_hash: float64(-1712216461), X_meta: nil}), value)
		})
	}(&AFn{})

	Seq = func(seq *AFn) *AFn {
		return Fn(seq, func(coll interface{}) CljsCoreISeq {
			if Nil_(coll) {
				return nil
			} else {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeqable", Str: "cljs.core/ISeqable", X_hash: float64(137437203), X_meta: nil}), coll)) {
					return Seq_(coll.(CljsCoreISeqable).X_seq_Arity1())
				} else {
					if reflect.ValueOf(coll).Kind() == reflect.Slice {
						if float64(len(coll.([]interface{}))) == float64(0) {
							return nil
						} else {
							return (&CljsCoreIndexedSeq{coll, float64(0)})
						}
					} else {
						if reflect.ValueOf(coll).Kind() == reflect.String {
							if float64(len(coll.([]interface{}))) == float64(0) {
								return nil
							} else {
								return (&CljsCoreIndexedSeq{coll, float64(0)})
							}
						} else {
							if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeqable", Str: "cljs.core/ISeqable", X_hash: float64(137437203), X_meta: nil}), coll)) {
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
		return Fn(first, func(coll interface{}) interface{} {
			if Nil_(coll) {
				return nil
			} else {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeq", Str: "cljs.core/ISeq", X_hash: float64(230133392), X_meta: nil}), coll)) {
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
		return Fn(rest, func(coll interface{}) CljsCoreISeq {
			if !(Nil_(coll)) {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeq", Str: "cljs.core/ISeq", X_hash: float64(230133392), X_meta: nil}), coll)) {
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
		return Fn(next, func(coll interface{}) CljsCoreISeq {
			if Nil_(coll) {
				return nil
			} else {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INext", Str: "cljs.core/INext", X_hash: float64(-113000046), X_meta: nil}), coll)) {
					return Seq_(coll.(CljsCoreINext).X_next_Arity1())
				} else {
					return Seq.Arity1IQ(Rest.Arity1IQ(coll))
				}
			}
		})
	}(&AFn{})

	Mix_collection_hash = func(mix_collection_hash *AFn) *AFn {
		return Fn(mix_collection_hash, func(hash_basis interface{}, count interface{}) float64 {
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
		return Fn(hash_ordered_coll, func(coll interface{}) float64 {
			{
				var n = float64(0)
				var hash_code = float64(1)
				var coll___1 = Seq.Arity1IQ(coll)
				_, _, _ = n, hash_code, coll___1
				for {
					if !(Nil_(coll___1)) {
						n, hash_code, coll___1 = (n + float64(1)), float64(int((Imul.Arity2IIF(float64(31), hash_code)+Hash.X_invoke_Arity1(First.X_invoke_Arity1(coll___1)).(float64)))|int(float64(0))), Next.Arity1IQ(coll___1)
						continue
					} else {
						return Mix_collection_hash.Arity2IIF(hash_code, n)
					}
				}
			}
		})
	}(&AFn{})

	Hash_unordered_coll = func(hash_unordered_coll *AFn) *AFn {
		return Fn(hash_unordered_coll, func(coll interface{}) float64 {
			{
				var n = float64(0)
				var hash_code = float64(0)
				var coll___1 = Seq.Arity1IQ(coll)
				_, _, _ = n, hash_code, coll___1
				for {
					if !(Nil_(coll___1)) {
						n, hash_code, coll___1 = (n + float64(1)), float64(int((hash_code+Hash.X_invoke_Arity1(First.X_invoke_Arity1(coll___1)).(float64)))|int(float64(0))), Next.Arity1IQ(coll___1)
						continue
					} else {
						return Mix_collection_hash.Arity2IIF(hash_code, n)
					}
				}
			}
		})
	}(&AFn{})

	Inc = func(inc *AFn) *AFn {
		return Fn(inc, func(x interface{}) interface{} {
			return (x.(float64) + float64(1))
		})
	}(&AFn{})

	X__GT_Reduced = func(__GT_Reduced *AFn) *AFn {
		return Fn(__GT_Reduced, func(val interface{}) interface{} {
			return (&CljsCoreReduced{val})
		})
	}(&AFn{})

	Reduced = func(reduced *AFn) *AFn {
		return Fn(reduced, func(x interface{}) interface{} {
			return (&CljsCoreReduced{x})
		})
	}(&AFn{})

	Reduced_QMARK_ = func(reduced_QMARK_ *AFn) *AFn {
		return Fn(reduced_QMARK_, func(r interface{}) bool {
			return func() bool { _, instanceof := r.(*CljsCoreReduced); return instanceof }()
		})
	}(&AFn{})

	Deref = func(deref *AFn) *AFn {
		return Fn(deref, func(o interface{}) interface{} {
			return o.(CljsCoreIDeref).X_deref_Arity1()
		})
	}(&AFn{})

	Ci_reduce = func(ci_reduce *AFn) *AFn {
		return Fn(ci_reduce, func(cicoll interface{}, f interface{}) interface{} {
			{
				var cnt = cicoll.(CljsCoreICounted).X_count_Arity1()
				_ = cnt
				if cnt == float64(0) {
					return f.(CljsCoreIFn).X_invoke_Arity0()
				} else {
					{
						var val = cicoll.(CljsCoreIIndexed).X_nth_Arity2(float64(0))
						var n = float64(1)
						_, _ = val, n
						for {
							if n < cnt {
								{
									var nval = f.(CljsCoreIFn).X_invoke_Arity2(val, cicoll.(CljsCoreIIndexed).X_nth_Arity2(n))
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
					var val___1 = val
					var n = float64(0)
					_, _ = val___1, n
					for {
						if n < cnt {
							{
								var nval = f.(CljsCoreIFn).X_invoke_Arity2(val___1, cicoll.(CljsCoreIIndexed).X_nth_Arity2(n))
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
					var val___1 = val
					var n = idx
					_, _ = val___1, n
					for {
						if n.(float64) < cnt {
							{
								var nval = f.(CljsCoreIFn).X_invoke_Arity2(val___1, cicoll.(CljsCoreIIndexed).X_nth_Arity2(n))
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
		return Fn(array_reduce, func(arr interface{}, f interface{}) interface{} {
			{
				var cnt = float64(len(arr.([]interface{})))
				_ = cnt
				if float64(len(arr.([]interface{}))) == float64(0) {
					return f.(CljsCoreIFn).X_invoke_Arity0()
				} else {
					{
						var val = (arr.([]interface{})[int(float64(0))])
						var n = float64(1)
						_, _ = val, n
						for {
							if n < cnt {
								{
									var nval = f.(CljsCoreIFn).X_invoke_Arity2(val, (arr.([]interface{})[int(n)]))
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
				var cnt = float64(len(arr.([]interface{})))
				_ = cnt
				{
					var val___1 = val
					var n = float64(0)
					_, _ = val___1, n
					for {
						if n < cnt {
							{
								var nval = f.(CljsCoreIFn).X_invoke_Arity2(val___1, (arr.([]interface{})[int(n)]))
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
				var cnt = float64(len(arr.([]interface{})))
				_ = cnt
				{
					var val___1 = val
					var n = idx
					_, _ = val___1, n
					for {
						if n.(float64) < cnt {
							{
								var nval = f.(CljsCoreIFn).X_invoke_Arity2(val___1, (arr.([]interface{})[int(n.(float64))]))
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
		return Fn(counted_QMARK_, func(x interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ICounted", Str: "cljs.core/ICounted", X_hash: float64(-1299011378), X_meta: nil}), x).(bool)
		})
	}(&AFn{})

	Indexed_QMARK_ = func(indexed_QMARK_ *AFn) *AFn {
		return Fn(indexed_QMARK_, func(x interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IIndexed", Str: "cljs.core/IIndexed", X_hash: float64(-436490749), X_meta: nil}), x).(bool)
		})
	}(&AFn{})

	X__GT_IndexedSeq = func(__GT_IndexedSeq *AFn) *AFn {
		return Fn(__GT_IndexedSeq, func(arr interface{}, i interface{}) interface{} {
			return (&CljsCoreIndexedSeq{arr, i})
		})
	}(&AFn{})

	Prim_seq = func(prim_seq *AFn) *AFn {
		return Fn(prim_seq, func(prim interface{}) interface{} {
			return prim_seq.X_invoke_Arity2(prim, float64(0))
		}, func(prim interface{}, i interface{}) interface{} {
			if i.(float64) < float64(len(prim.([]interface{}))) {
				return (&CljsCoreIndexedSeq{prim, i})
			} else {
				return nil
			}
		})
	}(&AFn{})

	Array_seq = func(array_seq *AFn) *AFn {
		return Fn(array_seq, func(array interface{}) interface{} {
			return Prim_seq.X_invoke_Arity2(array, float64(0))
		}, func(array interface{}, i interface{}) interface{} {
			return Prim_seq.X_invoke_Arity2(array, i)
		})
	}(&AFn{})

	X__GT_RSeq = func(__GT_RSeq *AFn) *AFn {
		return Fn(__GT_RSeq, func(ci interface{}, i interface{}, meta interface{}) interface{} {
			return (&CljsCoreRSeq{ci, i, meta})
		})
	}(&AFn{})

	Second = func(second *AFn) *AFn {
		return Fn(second, func(coll interface{}) interface{} {
			return First.X_invoke_Arity1(Next.Arity1IQ(coll))
		})
	}(&AFn{})

	Ffirst = func(ffirst *AFn) *AFn {
		return Fn(ffirst, func(coll interface{}) interface{} {
			return First.X_invoke_Arity1(First.X_invoke_Arity1(coll))
		})
	}(&AFn{})

	Nfirst = func(nfirst *AFn) *AFn {
		return Fn(nfirst, func(coll interface{}) interface{} {
			return Next.Arity1IQ(First.X_invoke_Arity1(coll))
		})
	}(&AFn{})

	Fnext = func(fnext *AFn) *AFn {
		return Fn(fnext, func(coll interface{}) interface{} {
			return First.X_invoke_Arity1(Next.Arity1IQ(coll))
		})
	}(&AFn{})

	Nnext = func(nnext *AFn) *AFn {
		return Fn(nnext, func(coll interface{}) interface{} {
			return Next.Arity1IQ(Next.Arity1IQ(coll))
		})
	}(&AFn{})

	Last = func(last *AFn) *AFn {
		return Fn(last, func(s interface{}) interface{} {
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
		return Fn(conj, func() interface{} {
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
			var xs = Array_seq.X_invoke_Arity1(coll_x_xs__[2:])
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
		return Fn(empty, func(coll interface{}) interface{} {
			if Nil_(coll) {
				return nil
			} else {
				return coll.(CljsCoreIEmptyableCollection).X_empty_Arity1()
			}
		})
	}(&AFn{})

	Accumulating_seq_count = func(accumulating_seq_count *AFn) *AFn {
		return Fn(accumulating_seq_count, func(coll interface{}) interface{} {
			{
				var s = Seq.Arity1IQ(coll)
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
		return Fn(count, func(coll interface{}) interface{} {
			if !(Nil_(coll)) {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ICounted", Str: "cljs.core/ICounted", X_hash: float64(-1299011378), X_meta: nil}), coll)) {
					return coll.(CljsCoreICounted).X_count_Arity1()
				} else {
					if reflect.ValueOf(coll).Kind() == reflect.Slice {
						return float64(len(coll.([]interface{})))
					} else {
						if reflect.ValueOf(coll).Kind() == reflect.String {
							return float64(len(coll.([]interface{})))
						} else {
							if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ICounted", Str: "cljs.core/ICounted", X_hash: float64(-1299011378), X_meta: nil}), coll)) {
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
		return Fn(linear_traversal_nth, func(coll interface{}, n interface{}) interface{} {
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
		return Fn(nth, func(coll interface{}, n interface{}) interface{} {
			if !(reflect.ValueOf(n).Kind() == reflect.Float64) {
				panic((&js.Error{"index argument to nth must be a number"}))
			} else {
				if Nil_(coll) {
					return coll
				} else {
					if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IIndexed", Str: "cljs.core/IIndexed", X_hash: float64(-436490749), X_meta: nil}), coll)) {
						return coll.(CljsCoreIIndexed).X_nth_Arity2(n)
					} else {
						if reflect.ValueOf(coll).Kind() == reflect.Slice {
							if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
								return (coll.([]interface{})[int(n.(float64))])
							} else {
								return nil
							}
						} else {
							if reflect.ValueOf(coll).Kind() == reflect.String {
								if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
									return (coll.([]interface{})[int(n.(float64))])
								} else {
									return nil
								}
							} else {
								if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IIndexed", Str: "cljs.core/IIndexed", X_hash: float64(-436490749), X_meta: nil}), coll)) {
									return coll.(CljsCoreIIndexed).X_nth_Arity2(n)
								} else {
									if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeq", Str: "cljs.core/ISeq", X_hash: float64(230133392), X_meta: nil}), coll)) {
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
			if !(reflect.ValueOf(n).Kind() == reflect.Float64) {
				panic((&js.Error{"index argument to nth must be a number."}))
			} else {
				if Nil_(coll) {
					return not_found
				} else {
					if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IIndexed", Str: "cljs.core/IIndexed", X_hash: float64(-436490749), X_meta: nil}), coll)) {
						return coll.(CljsCoreIIndexed).X_nth_Arity3(n, not_found)
					} else {
						if reflect.ValueOf(coll).Kind() == reflect.Slice {
							if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
								return (coll.([]interface{})[int(n.(float64))])
							} else {
								return not_found
							}
						} else {
							if reflect.ValueOf(coll).Kind() == reflect.String {
								if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
									return (coll.([]interface{})[int(n.(float64))])
								} else {
									return not_found
								}
							} else {
								if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IIndexed", Str: "cljs.core/IIndexed", X_hash: float64(-436490749), X_meta: nil}), coll)) {
									return coll.(CljsCoreIIndexed).X_nth_Arity2(n)
								} else {
									if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeq", Str: "cljs.core/ISeq", X_hash: float64(230133392), X_meta: nil}), coll)) {
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

	Get = func(get *AFn) *AFn {
		return Fn(get, func(o interface{}, k interface{}) interface{} {
			if Nil_(o) {
				return nil
			} else {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ILookup", Str: "cljs.core/ILookup", X_hash: float64(-150575073), X_meta: nil}), o)) {
					return o.(CljsCoreILookup).X_lookup_Arity2(k)
				} else {
					if reflect.ValueOf(o).Kind() == reflect.Slice {
						if k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64) {
							return (o.([]interface{})[int(k.(float64))])
						} else {
							return nil
						}
					} else {
						if reflect.ValueOf(o).Kind() == reflect.String {
							if k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64) {
								return (o.([]interface{})[int(k.(float64))])
							} else {
								return nil
							}
						} else {
							if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ILookup", Str: "cljs.core/ILookup", X_hash: float64(-150575073), X_meta: nil}), o)) {
								return o.(CljsCoreILookup).X_lookup_Arity2(k)
							} else {
								return nil

							}
						}
					}
				}
			}
		}, func(o interface{}, k interface{}, not_found interface{}) interface{} {
			if !(Nil_(o)) {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ILookup", Str: "cljs.core/ILookup", X_hash: float64(-150575073), X_meta: nil}), o)) {
					return o.(CljsCoreILookup).X_lookup_Arity3(k, not_found)
				} else {
					if reflect.ValueOf(o).Kind() == reflect.Slice {
						if k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64) {
							return (o.([]interface{})[int(k.(float64))])
						} else {
							return not_found
						}
					} else {
						if reflect.ValueOf(o).Kind() == reflect.String {
							if k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64) {
								return (o.([]interface{})[int(k.(float64))])
							} else {
								return not_found
							}
						} else {
							if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ILookup", Str: "cljs.core/ILookup", X_hash: float64(-150575073), X_meta: nil}), o)) {
								return o.(CljsCoreILookup).X_lookup_Arity3(k, not_found)
							} else {
								return not_found

							}
						}
					}
				}
			} else {
				return not_found
			}
		})
	}(&AFn{})

	Assoc = func(assoc *AFn) *AFn {
		return Fn(assoc, func(coll interface{}, k interface{}, v interface{}) interface{} {
			if !(Nil_(coll)) {
				return coll.(CljsCoreIAssociative).X_assoc_Arity3(k, v)
			} else {
				return CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{k}, []interface{}{v})
			}
		}, func(coll_k_v_kvs__ ...interface{}) interface{} {
			var coll = coll_k_v_kvs__[0]
			var k = coll_k_v_kvs__[1]
			var v = coll_k_v_kvs__[2]
			var kvs = Array_seq.X_invoke_Arity1(coll_k_v_kvs__[3:])
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
		return Fn(dissoc, func(coll interface{}) interface{} {
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
			var ks = Array_seq.X_invoke_Arity1(coll_k_ks__[2:])
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
		return Fn(fn_QMARK_, func(f interface{}) bool {
			{
				var or__175__auto__ = Truth_(Native_invoke_func.X_invoke_Arity2(goog.IsFunction, []interface{}{f}))
				_ = or__175__auto__
				if Truth_(or__175__auto__) {
					return or__175__auto__
				} else {
					return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "Fn", Str: "cljs.core/Fn", X_hash: float64(-695281833), X_meta: nil}), f).(bool)
				}
			}
		})
	}(&AFn{})

	X__GT_MetaFn = func(__GT_MetaFn *AFn) *AFn {
		return Fn(__GT_MetaFn, func(afn interface{}, meta interface{}) interface{} {
			return (&CljsCoreMetaFn{afn, meta})
		})
	}(&AFn{})

	With_meta = func(with_meta *AFn) *AFn {
		return Fn(with_meta, func(o interface{}, meta interface{}) interface{} {
			if (Fn_QMARK_.Arity1IB(o)) && (Not.Arity1IB(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IWithMeta", Str: "cljs.core/IWithMeta", X_hash: float64(-1981666051), X_meta: nil}), o))) {
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
		return Fn(meta, func(o interface{}) interface{} {
			if Truth_(func() interface{} {
				var and__163__auto__ = !(Nil_(o))
				_ = and__163__auto__
				if Truth_(and__163__auto__) {
					return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IMeta", Str: "cljs.core/IMeta", X_hash: float64(-1459057517), X_meta: nil}), o)
				} else {
					return and__163__auto__
				}
			}()) {
				return o.(CljsCoreIMeta).X_meta_Arity1()
			} else {
				return nil
			}
		})
	}(&AFn{})

	Peek = func(peek *AFn) *AFn {
		return Fn(peek, func(coll interface{}) interface{} {
			if Nil_(coll) {
				return nil
			} else {
				return coll.(CljsCoreIStack).X_peek_Arity1()
			}
		})
	}(&AFn{})

	Pop = func(pop *AFn) *AFn {
		return Fn(pop, func(coll interface{}) interface{} {
			if Nil_(coll) {
				return nil
			} else {
				return coll.(CljsCoreIStack).X_pop_Arity1()
			}
		})
	}(&AFn{})

	Disj = func(disj *AFn) *AFn {
		return Fn(disj, func(coll interface{}) interface{} {
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
			var ks = Array_seq.X_invoke_Arity1(coll_k_ks__[2:])
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
		return Fn(empty_QMARK_, func(coll interface{}) bool {
			return (Nil_(coll)) || (Not.Arity1IB(Seq.Arity1IQ(coll)))
		})
	}(&AFn{})

	Coll_QMARK_ = func(coll_QMARK_ *AFn) *AFn {
		return Fn(coll_QMARK_, func(x interface{}) bool {
			if Nil_(x) {
				return false
			} else {
				return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ICollection", Str: "cljs.core/ICollection", X_hash: float64(802638471), X_meta: nil}), x).(bool)
			}
		})
	}(&AFn{})

	Set_QMARK_ = func(set_QMARK_ *AFn) *AFn {
		return Fn(set_QMARK_, func(x interface{}) bool {
			if Nil_(x) {
				return false
			} else {
				return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISet", Str: "cljs.core/ISet", X_hash: float64(2003412810), X_meta: nil}), x).(bool)
			}
		})
	}(&AFn{})

	Associative_QMARK_ = func(associative_QMARK_ *AFn) *AFn {
		return Fn(associative_QMARK_, func(x interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IAssociative", Str: "cljs.core/IAssociative", X_hash: float64(-1700920611), X_meta: nil}), x).(bool)
		})
	}(&AFn{})

	Sequential_QMARK_ = func(sequential_QMARK_ *AFn) *AFn {
		return Fn(sequential_QMARK_, func(x interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISequential", Str: "cljs.core/ISequential", X_hash: float64(-950981796), X_meta: nil}), x).(bool)
		})
	}(&AFn{})

	Sorted_QMARK_ = func(sorted_QMARK_ *AFn) *AFn {
		return Fn(sorted_QMARK_, func(x interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISorted", Str: "cljs.core/ISorted", X_hash: float64(-1734125647), X_meta: nil}), x).(bool)
		})
	}(&AFn{})

	Reduceable_QMARK_ = func(reduceable_QMARK_ *AFn) *AFn {
		return Fn(reduceable_QMARK_, func(x interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), x).(bool)
		})
	}(&AFn{})

	Map_QMARK_ = func(map_QMARK_ *AFn) *AFn {
		return Fn(map_QMARK_, func(x interface{}) bool {
			if Nil_(x) {
				return false
			} else {
				return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IMap", Str: "cljs.core/IMap", X_hash: float64(1407777598), X_meta: nil}), x).(bool)
			}
		})
	}(&AFn{})

	Vector_QMARK_ = func(vector_QMARK_ *AFn) *AFn {
		return Fn(vector_QMARK_, func(x interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IVector", Str: "cljs.core/IVector", X_hash: float64(1711112835), X_meta: nil}), x).(bool)
		})
	}(&AFn{})

	Chunked_seq_QMARK_ = func(chunked_seq_QMARK_ *AFn) *AFn {
		return Fn(chunked_seq_QMARK_, func(x interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IChunkedSeq", Str: "cljs.core/IChunkedSeq", X_hash: float64(-892943716), X_meta: nil}), x).(bool)
		})
	}(&AFn{})

	Js_delete = func(js_delete *AFn) *AFn {
		return Fn(js_delete, func(obj interface{}, key interface{}) interface{} {
			return func(obj, key interface{}) interface{} { delete(obj.(map[string]interface{}), key.(string)); return obj }(obj, key)
		})
	}(&AFn{})

	Array_copy = func(array_copy *AFn) *AFn {
		return Fn(array_copy, func(from interface{}, i interface{}, to interface{}, j interface{}, len interface{}) interface{} {
			{
				var i___1 = i
				var j___1 = j
				var len___1 = len
				_, _, _ = i___1, j___1, len___1
				for {
					if len___1.(float64) == float64(0) {
						return to
					} else {
						to.([]interface{})[int(j___1.(float64))] = (from.([]interface{})[int(i___1.(float64))])
						i___1, j___1, len___1 = (i___1.(float64) + float64(1)), (j___1.(float64) + float64(1)), (len___1.(float64) - float64(1))
						continue
					}
				}
			}
		})
	}(&AFn{})

	Array_copy_downward = func(array_copy_downward *AFn) *AFn {
		return Fn(array_copy_downward, func(from interface{}, i interface{}, to interface{}, j interface{}, len interface{}) interface{} {
			{
				var i___1 = (i.(float64) + (len.(float64) - float64(1)))
				var j___1 = (j.(float64) + (len.(float64) - float64(1)))
				var len___1 = len
				_, _, _ = i___1, j___1, len___1
				for {
					if len___1.(float64) == float64(0) {
						return to
					} else {
						to.([]interface{})[int(j___1)] = (from.([]interface{})[int(i___1)])
						i___1, j___1, len___1 = (i___1 - float64(1)), (j___1 - float64(1)), (len___1.(float64) - float64(1))
						continue
					}
				}
			}
		})
	}(&AFn{})

	Lookup_sentinel = func() interface{} {
		var obj301 = map[string]interface{}{}
		_ = obj301
		return obj301
	}()

	False_QMARK_ = func(false_QMARK_ *AFn) *AFn {
		return Fn(false_QMARK_, func(x interface{}) bool {
			return x == false
		})
	}(&AFn{})

	True_QMARK_ = func(true_QMARK_ *AFn) *AFn {
		return Fn(true_QMARK_, func(x interface{}) bool {
			return x == true
		})
	}(&AFn{})

	Undefined_QMARK_ = func(undefined_QMARK_ *AFn) *AFn {
		return Fn(undefined_QMARK_, func(x interface{}) bool {
			return (nil == x)
		})
	}(&AFn{})

	Seq_QMARK_ = func(seq_QMARK_ *AFn) *AFn {
		return Fn(seq_QMARK_, func(s interface{}) bool {
			if Nil_(s) {
				return false
			} else {
				return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeq", Str: "cljs.core/ISeq", X_hash: float64(230133392), X_meta: nil}), s).(bool)
			}
		})
	}(&AFn{})

	Seqable_QMARK_ = func(seqable_QMARK_ *AFn) *AFn {
		return Fn(seqable_QMARK_, func(s interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeqable", Str: "cljs.core/ISeqable", X_hash: float64(137437203), X_meta: nil}), s).(bool)
		})
	}(&AFn{})

	Boolean = func(boolean *AFn) *AFn {
		return Fn(boolean, func(x interface{}) bool {
			if Truth_(x) {
				return true
			} else {
				return false
			}
		})
	}(&AFn{})

	Ifn_QMARK_ = func(ifn_QMARK_ *AFn) *AFn {
		return Fn(ifn_QMARK_, func(f interface{}) bool {
			{
				var or__175__auto__ = Fn_QMARK_.Arity1IB(f)
				_ = or__175__auto__
				if Truth_(or__175__auto__) {
					return or__175__auto__
				} else {
					return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IFn", Str: "cljs.core/IFn", X_hash: float64(-920223129), X_meta: nil}), f).(bool)
				}
			}
		})
	}(&AFn{})

	Contains_QMARK_ = func(contains_QMARK_ *AFn) *AFn {
		return Fn(contains_QMARK_, func(coll interface{}, v interface{}) bool {
			if reflect.DeepEqual(Get.X_invoke_Arity3(coll, v, Lookup_sentinel), Lookup_sentinel) {
				return false
			} else {
				return true
			}
		})
	}(&AFn{})

	Find = func(find *AFn) *AFn {
		return Fn(find, func(coll interface{}, k interface{}) interface{} {
			if (!(Nil_(coll))) && (Associative_QMARK_.Arity1IB(coll)) && (Contains_QMARK_.Arity2IIB(coll, k)) {
				return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{k, Get.X_invoke_Arity2(coll, k)}, nil})
			} else {
				return nil
			}
		})
	}(&AFn{})

	Distinct_QMARK_ = func(distinct_QMARK_ *AFn) *AFn {
		return Fn(distinct_QMARK_, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return !(X_EQ_.Arity2IIB(x, y))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			if !(X_EQ_.Arity2IIB(x, y)) {
				{
					var s = CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{x, y}, true).(*CljsCorePersistentHashSet)
					var xs = more
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
									s, xs = Conj.X_invoke_Arity2(s, x___1).(*CljsCorePersistentHashSet), etc
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
		return Fn(sequence, func(coll interface{}) CljsCoreISeq {
			if Seq_QMARK_.Arity1IB(coll) {
				return coll.(CljsCoreISeq)
			} else {
				{
					var or__175__auto__ = Seq.Arity1IQ(coll)
					_ = or__175__auto__
					if Truth_(or__175__auto__) {
						return or__175__auto__
					} else {
						return CljsCoreIEmptyList(CljsCoreList_EMPTY)
					}
				}
			}
		})
	}(&AFn{})

	Compare = func(compare *AFn) *AFn {
		return Fn(compare, func(x interface{}, y interface{}) float64 {
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
							if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IComparable", Str: "cljs.core/IComparable", X_hash: float64(1166626940), X_meta: nil}), x)) {
								return x.(CljsCoreIComparable).X_compare_Arity2(y)
							} else {
								return Native_invoke_func.X_invoke_Arity2(goog_array.DefaultCompare, []interface{}{x, y}).(float64)
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
		return Fn(compare_indexed, func(xs interface{}, ys interface{}) interface{} {
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
		return Fn(fn__GT_comparator, func(f interface{}) interface{} {
			if X_EQ_.Arity2IIB(f, Compare) {
				return Compare
			} else {
				return func(G__302 *AFn) *AFn {
					return Fn(G__302, func(x interface{}, y interface{}) interface{} {
						{
							var r = f.(CljsCoreIFn).X_invoke_Arity2(x, y)
							_ = r
							if reflect.ValueOf(r).Kind() == reflect.Float64 {
								return r
							} else {
								if Truth_(r) {
									return float64(-1)
								} else {
									if Truth_(f.(CljsCoreIFn).X_invoke_Arity2(y, x)) {
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

	Sort = func(sort *AFn) *AFn {
		return Fn(sort, func(coll interface{}) interface{} {
			return sort.X_invoke_Arity2(Compare, coll)
		}, func(comp interface{}, coll interface{}) interface{} {
			if Truth_(Seq.Arity1IQ(coll)) {
				{
					var a = To_array.X_invoke_Arity1(coll).([]interface{})
					_ = a
					Native_invoke_func.X_invoke_Arity2(goog_array.StableSort, []interface{}{a, Fn__GT_comparator.X_invoke_Arity1(comp)})
					return Seq.Arity1IQ(a)
				}
			} else {
				return CljsCoreIEmptyList(CljsCoreList_EMPTY)
			}
		})
	}(&AFn{})

	Sort_by = func(sort_by *AFn) *AFn {
		return Fn(sort_by, func(keyfn interface{}, coll interface{}) interface{} {
			return sort_by.X_invoke_Arity3(keyfn, Compare, coll)
		}, func(keyfn interface{}, comp interface{}, coll interface{}) interface{} {
			return Sort.X_invoke_Arity2(func(G__303 *AFn) *AFn {
				return Fn(G__303, func(x interface{}, y interface{}) interface{} {
					return Fn__GT_comparator.X_invoke_Arity1(comp).(CljsCoreIFn).X_invoke_Arity2(keyfn.(CljsCoreIFn).X_invoke_Arity1(x), keyfn.(CljsCoreIFn).X_invoke_Arity1(y))
				})
			}(&AFn{}), coll)
		})
	}(&AFn{})

	Seq_reduce = func(seq_reduce *AFn) *AFn {
		return Fn(seq_reduce, func(f interface{}, coll interface{}) interface{} {
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
					return f.(CljsCoreIFn).X_invoke_Arity0()
				}
			}
		}, func(f interface{}, val interface{}, coll interface{}) interface{} {
			{
				var val___1 = val
				var coll___1 = Seq.Arity1IQ(coll)
				_, _ = val___1, coll___1
				for {
					if Truth_(coll___1) {
						{
							var nval = f.(CljsCoreIFn).X_invoke_Arity2(val___1, First.X_invoke_Arity1(coll___1))
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
		return Fn(shuffle, func(coll interface{}) interface{} {
			{
				var a = To_array.X_invoke_Arity1(coll).([]interface{})
				_ = a
				Native_invoke_func.X_invoke_Arity2(goog_array.Shuffle, []interface{}{a})
				return Vec.X_invoke_Arity1(a)
			}
		})
	}(&AFn{})

	Reduce = func(reduce *AFn) *AFn {
		return Fn(reduce, func(f interface{}, coll interface{}) interface{} {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), coll)) {
				return coll.(CljsCoreIReduce).X_reduce_Arity2(f)
			} else {
				if reflect.ValueOf(coll).Kind() == reflect.Slice {
					return Array_reduce.X_invoke_Arity2(coll, f)
				} else {
					if reflect.ValueOf(coll).Kind() == reflect.String {
						return Array_reduce.X_invoke_Arity2(coll, f)
					} else {
						if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), coll)) {
							return coll.(CljsCoreIReduce).X_reduce_Arity2(f)
						} else {
							return Seq_reduce.X_invoke_Arity2(f, coll)

						}
					}
				}
			}
		}, func(f interface{}, val interface{}, coll interface{}) interface{} {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), coll)) {
				return coll.(CljsCoreIReduce).X_reduce_Arity3(f, val)
			} else {
				if reflect.ValueOf(coll).Kind() == reflect.Slice {
					return Array_reduce.X_invoke_Arity3(coll, f, val)
				} else {
					if reflect.ValueOf(coll).Kind() == reflect.String {
						return Array_reduce.X_invoke_Arity3(coll, f, val)
					} else {
						if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), coll)) {
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
		return Fn(reduce_kv, func(f interface{}, init interface{}, coll interface{}) interface{} {
			if !(Nil_(coll)) {
				return coll.(CljsCoreIKVReduce).X_kv_reduce_Arity3(f, init)
			} else {
				return init
			}
		})
	}(&AFn{})

	Identity = func(identity *AFn) *AFn {
		return Fn(identity, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Completing = func(completing *AFn) *AFn {
		return Fn(completing, func(f interface{}) interface{} {
			return completing.X_invoke_Arity2(f, Identity).(CljsCoreIFn)
		}, func(f interface{}, cf interface{}) interface{} {
			return func(G__304 *AFn) *AFn {
				return Fn(G__304, func() interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity0()
				}, func(x interface{}) interface{} {
					return cf.(CljsCoreIFn).X_invoke_Arity1(x)
				}, func(x interface{}, y interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity2(x, y)
				})
			}(&AFn{})
		})
	}(&AFn{})

	Transduce = func(transduce *AFn) *AFn {
		return Fn(transduce, func(xform interface{}, f interface{}, coll interface{}) interface{} {
			return transduce.X_invoke_Arity4(xform, f, f.(CljsCoreIFn).X_invoke_Arity0(), coll)
		}, func(xform interface{}, f interface{}, init interface{}, coll interface{}) interface{} {
			{
				var f___1 = xform.(CljsCoreIFn).X_invoke_Arity1(f)
				var ret = Reduce.X_invoke_Arity3(f___1, init, coll)
				_, _ = f___1, ret
				return f___1.(CljsCoreIFn).X_invoke_Arity1(ret)
			}
		})
	}(&AFn{})

	X_PLUS_ = func(_PLUS_ *AFn) *AFn {
		return Fn(_PLUS_, func() float64 {
			return float64(0)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) + y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(_PLUS_, (x.(float64) + y.(float64)), more)
		})
	}(&AFn{})

	X_ = func(___ *AFn) *AFn {
		return Fn(___, func(x interface{}) float64 {
			return (-x.(float64))
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) - y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(___, (x.(float64) - y.(float64)), more)
		})
	}(&AFn{})

	X_STAR_ = func(_STAR_ *AFn) *AFn {
		return Fn(_STAR_, func() float64 {
			return float64(1)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) * y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(_STAR_, (x.(float64) * y.(float64)), more)
		})
	}(&AFn{})

	X_SLASH_ = func(_SLASH_ *AFn) *AFn {
		return Fn(_SLASH_, func(x interface{}) float64 {
			return _SLASH_.Arity2IIF(float64(1), x)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) / y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(_SLASH_, _SLASH_.Arity2IIF(x, y), more)
		})
	}(&AFn{})

	X_LT_ = func(_LT_ *AFn) *AFn {
		return Fn(_LT_, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return (x.(float64) < y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
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
		return Fn(_LT__EQ_, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return (x.(float64) <= y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
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
		return Fn(_GT_, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return (x.(float64) > y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
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
		return Fn(_GT__EQ_, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return (x.(float64) >= y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
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
		return Fn(dec, func(x interface{}) interface{} {
			return (x.(float64) - float64(1))
		})
	}(&AFn{})

	Max = func(max *AFn) *AFn {
		return Fn(max, func(x interface{}) float64 {
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
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
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
		return Fn(min, func(x interface{}) float64 {
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
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
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
		return Fn(byte_, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Short = func(short *AFn) *AFn {
		return Fn(short, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Float_ = func(float_ *AFn) *AFn {
		return Fn(float_, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Double = func(double *AFn) *AFn {
		return Fn(double, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_byte = func(unchecked_byte *AFn) *AFn {
		return Fn(unchecked_byte, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_char = func(unchecked_char *AFn) *AFn {
		return Fn(unchecked_char, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_short = func(unchecked_short *AFn) *AFn {
		return Fn(unchecked_short, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_float = func(unchecked_float *AFn) *AFn {
		return Fn(unchecked_float, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_double = func(unchecked_double *AFn) *AFn {
		return Fn(unchecked_double, func(x interface{}) float64 {
			return x.(float64)
		})
	}(&AFn{})

	Unchecked_add = func(unchecked_add *AFn) *AFn {
		return Fn(unchecked_add, func() float64 {
			return float64(0)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) + y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_add, (x.(float64) + y.(float64)), more)
		})
	}(&AFn{})

	Unchecked_add_int = func(unchecked_add_int *AFn) *AFn {
		return Fn(unchecked_add_int, func() float64 {
			return float64(0)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) + y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_add_int, (x.(float64) + y.(float64)), more)
		})
	}(&AFn{})

	Unchecked_dec = func(unchecked_dec *AFn) *AFn {
		return Fn(unchecked_dec, func(x interface{}) interface{} {
			return (x.(float64) - float64(1))
		})
	}(&AFn{})

	Unchecked_dec_int = func(unchecked_dec_int *AFn) *AFn {
		return Fn(unchecked_dec_int, func(x interface{}) interface{} {
			return (x.(float64) - float64(1))
		})
	}(&AFn{})

	Unchecked_divide_int = func(unchecked_divide_int *AFn) *AFn {
		return Fn(unchecked_divide_int, func(x interface{}) float64 {
			return unchecked_divide_int.Arity2IIF(float64(1), x)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) / y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_divide_int, unchecked_divide_int.Arity2IIF(x, y), more)
		})
	}(&AFn{})

	Unchecked_inc = func(unchecked_inc *AFn) *AFn {
		return Fn(unchecked_inc, func(x interface{}) interface{} {
			return (x.(float64) + float64(1))
		})
	}(&AFn{})

	Unchecked_inc_int = func(unchecked_inc_int *AFn) *AFn {
		return Fn(unchecked_inc_int, func(x interface{}) interface{} {
			return (x.(float64) + float64(1))
		})
	}(&AFn{})

	Unchecked_multiply = func(unchecked_multiply *AFn) *AFn {
		return Fn(unchecked_multiply, func() float64 {
			return float64(1)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) * y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_multiply, (x.(float64) * y.(float64)), more)
		})
	}(&AFn{})

	Unchecked_multiply_int = func(unchecked_multiply_int *AFn) *AFn {
		return Fn(unchecked_multiply_int, func() float64 {
			return float64(1)
		}, func(x interface{}) float64 {
			return x.(float64)
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) * y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_multiply_int, (x.(float64) * y.(float64)), more)
		})
	}(&AFn{})

	Unchecked_negate = func(unchecked_negate *AFn) *AFn {
		return Fn(unchecked_negate, func(x interface{}) interface{} {
			return (-x.(float64))
		})
	}(&AFn{})

	Unchecked_negate_int = func(unchecked_negate_int *AFn) *AFn {
		return Fn(unchecked_negate_int, func(x interface{}) interface{} {
			return (-x.(float64))
		})
	}(&AFn{})

	Unchecked_remainder_int = func(unchecked_remainder_int *AFn) *AFn {
		return Fn(unchecked_remainder_int, func(x interface{}, n interface{}) interface{} {
			return Mod.X_invoke_Arity2(x, n).(float64)
		})
	}(&AFn{})

	Unchecked_subtract = func(unchecked_subtract *AFn) *AFn {
		return Fn(unchecked_subtract, func(x interface{}) float64 {
			return (-x.(float64))
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) - y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_subtract, (x.(float64) - y.(float64)), more)
		})
	}(&AFn{})

	Unchecked_subtract_int = func(unchecked_subtract_int *AFn) *AFn {
		return Fn(unchecked_subtract_int, func(x interface{}) float64 {
			return (-x.(float64))
		}, func(x interface{}, y interface{}) float64 {
			return (x.(float64) - y.(float64))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Reduce.X_invoke_Arity3(unchecked_subtract_int, (x.(float64) - y.(float64)), more)
		})
	}(&AFn{})

	Fix = func(fix *AFn) *AFn {
		return Fn(fix, func(q interface{}) float64 {
			if q.(float64) >= float64(0) {
				return Native_invoke_func.X_invoke_Arity2(Math.Floor, []interface{}{q}).(float64)
			} else {
				return Native_invoke_func.X_invoke_Arity2(Math.Ceil, []interface{}{q}).(float64)
			}
		})
	}(&AFn{})

	Int_ = func(int_ *AFn) *AFn {
		return Fn(int_, func(x interface{}) interface{} {
			return float64(int(x.(float64)) | int(float64(0)))
		})
	}(&AFn{})

	Unchecked_int = func(unchecked_int *AFn) *AFn {
		return Fn(unchecked_int, func(x interface{}) interface{} {
			return Fix.Arity1IF(x)
		})
	}(&AFn{})

	Long = func(long *AFn) *AFn {
		return Fn(long, func(x interface{}) interface{} {
			return Fix.Arity1IF(x)
		})
	}(&AFn{})

	Unchecked_long = func(unchecked_long *AFn) *AFn {
		return Fn(unchecked_long, func(x interface{}) interface{} {
			return Fix.Arity1IF(x)
		})
	}(&AFn{})

	Booleans = func(booleans *AFn) *AFn {
		return Fn(booleans, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Bytes = func(bytes *AFn) *AFn {
		return Fn(bytes, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Chars = func(chars *AFn) *AFn {
		return Fn(chars, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Shorts = func(shorts *AFn) *AFn {
		return Fn(shorts, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Ints = func(ints *AFn) *AFn {
		return Fn(ints, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Floats = func(floats *AFn) *AFn {
		return Fn(floats, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Doubles = func(doubles *AFn) *AFn {
		return Fn(doubles, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Longs = func(longs *AFn) *AFn {
		return Fn(longs, func(x interface{}) interface{} {
			return x
		})
	}(&AFn{})

	Js_mod = func(js_mod *AFn) *AFn {
		return Fn(js_mod, func(n interface{}, d interface{}) interface{} {
			return float64(int(n.(float64)) % int(d.(float64)))
		})
	}(&AFn{})

	Mod = func(mod *AFn) *AFn {
		return Fn(mod, func(n interface{}, d interface{}) interface{} {
			return float64(int((float64(int(n.(float64))%int(d.(float64))) + d.(float64))) % int(d.(float64)))
		})
	}(&AFn{})

	Quot = func(quot *AFn) *AFn {
		return Fn(quot, func(n interface{}, d interface{}) interface{} {
			{
				var rem = float64(int(n.(float64)) % int(d.(float64)))
				_ = rem
				return Fix.Arity1IF(((n.(float64) - rem) / d.(float64)))
			}
		})
	}(&AFn{})

	Rem = func(rem *AFn) *AFn {
		return Fn(rem, func(n interface{}, d interface{}) interface{} {
			{
				var q = Quot.X_invoke_Arity2(n, d).(float64)
				_ = q
				return (n.(float64) - (d.(float64) * q))
			}
		})
	}(&AFn{})

	Rand = func(rand *AFn) *AFn {
		return Fn(rand, func() float64 {
			return Native_invoke_func.X_invoke_Arity2(Math.Random, []interface{}{}).(float64)
		}, func(n interface{}) float64 {
			return (n.(float64) * rand.Arity0F())
		})
	}(&AFn{})

	Rand_int = func(rand_int *AFn) *AFn {
		return Fn(rand_int, func(n interface{}) interface{} {
			return Fix.Arity1IF(Rand.Arity1IF(n))
		})
	}(&AFn{})

	Bit_xor = func(bit_xor *AFn) *AFn {
		return Fn(bit_xor, func(x interface{}, y interface{}) interface{} {
			return float64(int(x.(float64)) ^ int(y.(float64)))
		})
	}(&AFn{})

	Bit_and = func(bit_and *AFn) *AFn {
		return Fn(bit_and, func(x interface{}, y interface{}) interface{} {
			return float64(int(x.(float64)) & int(y.(float64)))
		})
	}(&AFn{})

	Bit_or = func(bit_or *AFn) *AFn {
		return Fn(bit_or, func(x interface{}, y interface{}) interface{} {
			return float64(int(x.(float64)) | int(y.(float64)))
		})
	}(&AFn{})

	Bit_and_not = func(bit_and_not *AFn) *AFn {
		return Fn(bit_and_not, func(x interface{}, y interface{}) interface{} {
			return float64(int(x.(float64)) &^ int(y.(float64)))
		})
	}(&AFn{})

	Bit_clear = func(bit_clear *AFn) *AFn {
		return Fn(bit_clear, func(x interface{}, n interface{}) interface{} {
			return float64(int(x.(float64)) &^ (int(1) << uint(n.(float64))))
		})
	}(&AFn{})

	Bit_flip = func(bit_flip *AFn) *AFn {
		return Fn(bit_flip, func(x interface{}, n interface{}) interface{} {
			return float64(int(x.(float64)) ^ (int(1) << uint(n.(float64))))
		})
	}(&AFn{})

	Bit_not = func(bit_not *AFn) *AFn {
		return Fn(bit_not, func(x interface{}) interface{} {
			return float64(^int(x.(float64)))
		})
	}(&AFn{})

	Bit_set = func(bit_set *AFn) *AFn {
		return Fn(bit_set, func(x interface{}, n interface{}) interface{} {
			return float64(int(x.(float64)) | (int(1) << uint(n.(float64))))
		})
	}(&AFn{})

	Bit_test = func(bit_test *AFn) *AFn {
		return Fn(bit_test, func(x interface{}, n interface{}) interface{} {
			return float64((int(x.(float64)) & (int(1) << uint(n.(float64))))) != 0
		})
	}(&AFn{})

	Bit_shift_left = func(bit_shift_left *AFn) *AFn {
		return Fn(bit_shift_left, func(x interface{}, n interface{}) interface{} {
			return float64(int(x.(float64)) << uint(n.(float64)))
		})
	}(&AFn{})

	Bit_shift_right = func(bit_shift_right *AFn) *AFn {
		return Fn(bit_shift_right, func(x interface{}, n interface{}) interface{} {
			return float64(int(x.(float64)) >> uint(n.(float64)))
		})
	}(&AFn{})

	Bit_shift_right_zero_fill = func(bit_shift_right_zero_fill *AFn) *AFn {
		return Fn(bit_shift_right_zero_fill, func(x interface{}, n interface{}) interface{} {
			return float64(uint(x.(float64)) >> uint(n.(float64)))
		})
	}(&AFn{})

	Unsigned_bit_shift_right = func(unsigned_bit_shift_right *AFn) *AFn {
		return Fn(unsigned_bit_shift_right, func(x interface{}, n interface{}) interface{} {
			return float64(uint(x.(float64)) >> uint(n.(float64)))
		})
	}(&AFn{})

	Bit_count = func(bit_count *AFn) *AFn {
		return Fn(bit_count, func(v interface{}) interface{} {
			{
				var v___1 = (v.(float64) - float64(int(float64(int(v.(float64))>>uint(float64(1))))&int(float64(1431655765))))
				var v___2 = (float64(int(v___1)&int(float64(858993459))) + float64(int(float64(int(v___1)>>uint(float64(2))))&int(float64(858993459))))
				_, _ = v___1, v___2
				return float64(int((float64(int((v___2+float64(int(v___2)>>uint(float64(4)))))&int(float64(252645135))) * float64(16843009))) >> uint(float64(24)))
			}
		})
	}(&AFn{})

	X_EQ__EQ_ = func(_EQ__EQ_ *AFn) *AFn {
		return Fn(_EQ__EQ_, func(x interface{}) bool {
			return true
		}, func(x interface{}, y interface{}) bool {
			return x.(CljsCoreIEquiv).X_equiv_Arity2(y)
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
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
		return Fn(pos_QMARK_, func(n interface{}) bool {
			return (n.(float64) > float64(0))
		})
	}(&AFn{})

	Zero_QMARK_ = func(zero_QMARK_ *AFn) *AFn {
		return Fn(zero_QMARK_, func(n interface{}) bool {
			return (n.(float64) == float64(0))
		})
	}(&AFn{})

	Neg_QMARK_ = func(neg_QMARK_ *AFn) *AFn {
		return Fn(neg_QMARK_, func(x interface{}) bool {
			return (x.(float64) < float64(0))
		})
	}(&AFn{})

	Nthnext = func(nthnext *AFn) *AFn {
		return Fn(nthnext, func(coll interface{}, n interface{}) interface{} {
			{
				var n___1 = n
				var xs = Seq.Arity1IQ(coll)
				_, _ = n___1, xs
				for {
					if Truth_(func() interface{} {
						var and__163__auto__ = xs
						_ = and__163__auto__
						if Truth_(and__163__auto__) {
							return (n___1.(float64) > float64(0))
						} else {
							return and__163__auto__
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
		return Fn(str, func() interface{} {
			return ""
		}, func(x interface{}) interface{} {
			if Nil_(x) {
				return ""
			} else {
				return fmt.Sprint(x)
			}
		}, func(x_ys__ ...interface{}) interface{} {
			var x = x_ys__[0]
			var ys = Array_seq.X_invoke_Arity1(x_ys__[1:])
			_, _ = x, ys
			{
				var sb = (&goog_string.StringBuffer{str.X_invoke_Arity1(x)})
				var more = ys
				_, _ = sb, more
				for {
					if Truth_(more) {
						sb, more = sb.Append(str.X_invoke_Arity1(First.X_invoke_Arity1(more))), Next.Arity1IQ(more)
						continue
					} else {
						return sb.ToString()
					}
				}
			}
		})
	}(&AFn{})

	Subs = func(subs *AFn) *AFn {
		return Fn(subs, func(s interface{}, start interface{}) interface{} {
			return Native_invoke_instance_method.X_invoke_Arity3(s, "Substring", []interface{}{start})
		}, func(s interface{}, start interface{}, end interface{}) interface{} {
			return Native_invoke_instance_method.X_invoke_Arity3(s, "Substring", []interface{}{start, end})
		})
	}(&AFn{})

	Equiv_sequential = func(equiv_sequential *AFn) *AFn {
		return Fn(equiv_sequential, func(x interface{}, y interface{}) interface{} {
			return Boolean.Arity1IB(func() interface{} {
				if Sequential_QMARK_.Arity1IB(y) {
					return func() bool {
						if (Counted_QMARK_.Arity1IB(x)) && (Counted_QMARK_.Arity1IB(y)) && (!(Count.X_invoke_Arity1(x).(float64) == Count.X_invoke_Arity1(y).(float64))) {
							return false
						} else {
							return func() bool {
								var xs = Seq.Arity1IQ(x)
								var ys = Seq.Arity1IQ(y)
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
		return Fn(hash_coll, func(coll interface{}) interface{} {
			if Truth_(Seq.Arity1IQ(coll)) {
				{
					var res = Hash.X_invoke_Arity1(First.X_invoke_Arity1(coll))
					var s = Next.Arity1IQ(coll)
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
		return Fn(hash_imap, func(m interface{}) interface{} {
			{
				var h = float64(0)
				var s = Seq.Arity1IQ(m)
				_, _ = h, s
				for {
					if Truth_(s) {
						{
							var e = First.X_invoke_Arity1(s)
							_ = e
							h, s = float64(int((h+float64(int(Hash.X_invoke_Arity1(Key.X_invoke_Arity1(e)).(float64))^int(Hash.X_invoke_Arity1(Val.X_invoke_Arity1(e)).(float64)))))%int(float64(4503599627370496))), Next.Arity1IQ(s)
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
		return Fn(hash_iset, func(s interface{}) interface{} {
			{
				var h = float64(0)
				var s___1 = Seq.Arity1IQ(s)
				_, _ = h, s___1
				for {
					if Truth_(s___1) {
						{
							var e = First.X_invoke_Arity1(s___1)
							_ = e
							h, s___1 = float64(int((h+Hash.X_invoke_Arity1(e).(float64)))%int(float64(4503599627370496))), Next.Arity1IQ(s___1)
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
		return Fn(extend_object_BANG_, func(obj interface{}, fn_map interface{}) interface{} {
			{
				var seq__311_317 = Seq.Arity1IQ(fn_map)
				var chunk__312_318 interface{} = nil
				var count__313_319 = float64(0)
				var i__314_320 = float64(0)
				_, _, _, _ = seq__311_317, chunk__312_318, count__313_319, i__314_320
				for {
					if i__314_320 < count__313_319 {
						{
							var vec__315_321 = chunk__312_318.(CljsCoreIIndexed).X_nth_Arity2(i__314_320)
							var key_name_322 = Nth.X_invoke_Arity3(vec__315_321, float64(0), nil)
							var f_323 = Nth.X_invoke_Arity3(vec__315_321, float64(1), nil)
							_, _, _ = vec__315_321, key_name_322, f_323
							{
								var str_name_324 = Name.X_invoke_Arity1(key_name_322)
								_ = str_name_324
								obj.([]interface{})[int(str_name_324.(float64))] = f_323
							}
							seq__311_317, chunk__312_318, count__313_319, i__314_320 = seq__311_317, chunk__312_318, count__313_319, (i__314_320 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___325 = Seq.Arity1IQ(seq__311_317)
							_ = temp__4222__auto___325
							if Truth_(temp__4222__auto___325) {
								{
									var seq__311_326___1 = temp__4222__auto___325
									_ = seq__311_326___1
									if Chunked_seq_QMARK_.Arity1IB(seq__311_326___1) {
										{
											var c__942__auto___327 = Chunk_first.X_invoke_Arity1(seq__311_326___1)
											_ = c__942__auto___327
											seq__311_317, chunk__312_318, count__313_319, i__314_320 = Chunk_rest.X_invoke_Arity1(seq__311_326___1).(CljsCoreISeq), c__942__auto___327, Count.X_invoke_Arity1(c__942__auto___327).(float64), float64(0)
											continue
										}
									} else {
										{
											var vec__316_328 = First.X_invoke_Arity1(seq__311_326___1)
											var key_name_329 = Nth.X_invoke_Arity3(vec__316_328, float64(0), nil)
											var f_330 = Nth.X_invoke_Arity3(vec__316_328, float64(1), nil)
											_, _, _ = vec__316_328, key_name_329, f_330
											{
												var str_name_331 = Name.X_invoke_Arity1(key_name_329)
												_ = str_name_331
												obj.([]interface{})[int(str_name_331.(float64))] = f_330
											}
											seq__311_317, chunk__312_318, count__313_319, i__314_320 = Next.Arity1IQ(seq__311_326___1), nil, float64(0), float64(0)
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
		return Fn(__GT_List, func(meta interface{}, first interface{}, rest interface{}, count interface{}, __hash interface{}) interface{} {
			return (&CljsCoreList{meta, first, rest, count, __hash})
		})
	}(&AFn{})

	X__GT_EmptyList = func(__GT_EmptyList *AFn) *AFn {
		return Fn(__GT_EmptyList, func(meta interface{}) interface{} {
			return (&CljsCoreEmptyList{meta})
		})
	}(&AFn{})

	Reversible_QMARK_ = func(reversible_QMARK_ *AFn) *AFn {
		return Fn(reversible_QMARK_, func(coll interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReversible", Str: "cljs.core/IReversible", X_hash: float64(-1422278012), X_meta: nil}), coll).(bool)
		})
	}(&AFn{})

	Rseq = func(rseq *AFn) *AFn {
		return Fn(rseq, func(coll interface{}) CljsCoreISeq {
			return Seq_(coll.(CljsCoreIReversible).X_rseq_Arity1())
		})
	}(&AFn{})

	Reverse = func(reverse *AFn) *AFn {
		return Fn(reverse, func(coll interface{}) interface{} {
			if Reversible_QMARK_.Arity1IB(coll) {
				return Rseq.Arity1IQ(coll)
			} else {
				return Reduce.X_invoke_Arity3(Conj, CljsCoreIEmptyList(CljsCoreList_EMPTY), coll)
			}
		})
	}(&AFn{})

	List = func(list *AFn) *AFn {
		return Fn(list, func(xs__ ...interface{}) interface{} {
			var xs = Array_seq.X_invoke_Arity1(xs__[0:])
			_ = xs
			{
				var arr = func() interface{} {
					if (func() bool { _, instanceof := xs.(*CljsCoreIndexedSeq); return instanceof }()) && (Native_get_instance_field.X_invoke_Arity2(xs, "I").(float64) == float64(0)) {
						return Native_get_instance_field.X_invoke_Arity2(xs, "Arr")
					} else {
						return func() []interface{} {
							var arr = []interface{}{}
							_ = arr
							{
								var xs___1 = xs
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
					var i = float64(len(arr.([]interface{})))
					var r = CljsCoreIEmptyList(CljsCoreList_EMPTY)
					_, _ = i, r
					for {
						if i > float64(0) {
							i, r = (i - float64(1)), r.X_conj_Arity2((arr.([]interface{})[int((i - float64(1)))])).(*CljsCoreList)
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
		return Fn(__GT_Cons, func(meta interface{}, first interface{}, rest interface{}, __hash interface{}) interface{} {
			return (&CljsCoreCons{meta, first, rest, __hash})
		})
	}(&AFn{})

	Cons = func(cons *AFn) *AFn {
		return Fn(cons, func(x interface{}, coll interface{}) interface{} {
			if Truth_(func() interface{} {
				var or__175__auto__ = Nil_(coll)
				_ = or__175__auto__
				if Truth_(or__175__auto__) {
					return or__175__auto__
				} else {
					return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeq", Str: "cljs.core/ISeq", X_hash: float64(230133392), X_meta: nil}), coll)
				}
			}()) {
				return (&CljsCoreCons{nil, x, coll, nil})
			} else {
				return (&CljsCoreCons{nil, x, Seq.Arity1IQ(coll), nil})
			}
		})
	}(&AFn{})

	List_QMARK_ = func(list_QMARK_ *AFn) *AFn {
		return Fn(list_QMARK_, func(x interface{}) bool {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IList", Str: "cljs.core/IList", X_hash: float64(1015168964), X_meta: nil}), x).(bool)
		})
	}(&AFn{})

	Hash_keyword = func(hash_keyword *AFn) *AFn {
		return Fn(hash_keyword, func(k interface{}) interface{} {
			return float64(int((Hash_symbol.X_invoke_Arity1(k).(float64) + float64(2654435769))) | int(float64(0)))
		})
	}(&AFn{})

	X__GT_Keyword = func(__GT_Keyword *AFn) *AFn {
		return Fn(__GT_Keyword, func(ns interface{}, name interface{}, fqn interface{}, _hash interface{}) interface{} {
			return (&CljsCoreKeyword{ns, name, fqn, _hash})
		})
	}(&AFn{})

	Keyword_QMARK_ = func(keyword_QMARK_ *AFn) *AFn {
		return Fn(keyword_QMARK_, func(x interface{}) bool {
			return func() bool { _, instanceof := x.(*CljsCoreKeyword); return instanceof }()
		})
	}(&AFn{})

	Keyword_identical_QMARK_ = func(keyword_identical_QMARK_ *AFn) *AFn {
		return Fn(keyword_identical_QMARK_, func(x interface{}, y interface{}) bool {
			if reflect.DeepEqual(x, y) {
				return true
			} else {
				if (func() bool { _, instanceof := x.(*CljsCoreKeyword); return instanceof }()) && (func() bool { _, instanceof := y.(*CljsCoreKeyword); return instanceof }()) {
					return reflect.DeepEqual(Native_get_instance_field.X_invoke_Arity2(x, "Fqn"), Native_get_instance_field.X_invoke_Arity2(y, "Fqn"))
				} else {
					return false
				}
			}
		})
	}(&AFn{})

	Namespace = func(namespace *AFn) *AFn {
		return Fn(namespace, func(x interface{}) interface{} {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INamed", Str: "cljs.core/INamed", X_hash: float64(-857199025), X_meta: nil}), x)) {
				return x.(CljsCoreINamed).X_namespace_Arity1()
			} else {
				panic((&js.Error{("Doesn't support namespace: " + Str.X_invoke_Arity1(x).(string))}))
			}
		})
	}(&AFn{})

	Keyword = func(keyword *AFn) *AFn {
		return Fn(keyword, func(name interface{}) interface{} {
			if func() bool { _, instanceof := name.(*CljsCoreKeyword); return instanceof }() {
				return name
			} else {
				if func() bool { _, instanceof := name.(*CljsCoreSymbol); return instanceof }() {
					return (&CljsCoreKeyword{Namespace.X_invoke_Arity1(name).(string), Name.X_invoke_Arity1(name), Native_get_instance_field.X_invoke_Arity2(name, "Str"), nil})
				} else {
					if reflect.ValueOf(name).Kind() == reflect.String {
						{
							var parts = Native_invoke_instance_method.X_invoke_Arity3(name, "Split", []interface{}{"/"})
							_ = parts
							if float64(len(parts.([]interface{}))) == float64(2) {
								return (&CljsCoreKeyword{(parts.([]interface{})[int(float64(0))]), (parts.([]interface{})[int(float64(1))]), name, nil})
							} else {
								return (&CljsCoreKeyword{nil, (parts.([]interface{})[int(float64(0))]), name, nil})
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
		return Fn(__GT_LazySeq, func(meta interface{}, fn interface{}, s interface{}, __hash interface{}) interface{} {
			return (&CljsCoreLazySeq{meta, fn, s, __hash})
		})
	}(&AFn{})

	X__GT_ChunkBuffer = func(__GT_ChunkBuffer *AFn) *AFn {
		return Fn(__GT_ChunkBuffer, func(buf interface{}, end interface{}) interface{} {
			return (&CljsCoreChunkBuffer{buf, end})
		})
	}(&AFn{})

	Chunk_buffer = func(chunk_buffer *AFn) *AFn {
		return Fn(chunk_buffer, func(capacity interface{}) interface{} {
			return (&CljsCoreChunkBuffer{make([]interface{}, int(capacity.(float64))), float64(0)})
		})
	}(&AFn{})

	X__GT_ArrayChunk = func(__GT_ArrayChunk *AFn) *AFn {
		return Fn(__GT_ArrayChunk, func(arr interface{}, off interface{}, end interface{}) interface{} {
			return (&CljsCoreArrayChunk{arr, off, end})
		})
	}(&AFn{})

	Array_chunk = func(array_chunk *AFn) *AFn {
		return Fn(array_chunk, func(arr interface{}) interface{} {
			return (&CljsCoreArrayChunk{arr, float64(0), float64(len(arr.([]interface{})))})
		}, func(arr interface{}, off interface{}) interface{} {
			return (&CljsCoreArrayChunk{arr, off, float64(len(arr.([]interface{})))})
		}, func(arr interface{}, off interface{}, end interface{}) interface{} {
			return (&CljsCoreArrayChunk{arr, off, end})
		})
	}(&AFn{})

	X__GT_ChunkedCons = func(__GT_ChunkedCons *AFn) *AFn {
		return Fn(__GT_ChunkedCons, func(chunk interface{}, more interface{}, meta interface{}, __hash interface{}) interface{} {
			return (&CljsCoreChunkedCons{chunk, more, meta, __hash})
		})
	}(&AFn{})

	Chunk_cons = func(chunk_cons *AFn) *AFn {
		return Fn(chunk_cons, func(chunk interface{}, rest interface{}) interface{} {
			if chunk.(CljsCoreICounted).X_count_Arity1() == float64(0) {
				return rest
			} else {
				return (&CljsCoreChunkedCons{chunk, rest, nil, nil})
			}
		})
	}(&AFn{})

	Chunk_append = func(chunk_append *AFn) *AFn {
		return Fn(chunk_append, func(b interface{}, x interface{}) interface{} {
			return Native_invoke_instance_method.X_invoke_Arity3(b, "Add", []interface{}{x})
		})
	}(&AFn{})

	Chunk = func(chunk *AFn) *AFn {
		return Fn(chunk, func(b interface{}) interface{} {
			return Native_invoke_instance_method.X_invoke_Arity3(b, "Chunk", []interface{}{})
		})
	}(&AFn{})

	Chunk_first = func(chunk_first *AFn) *AFn {
		return Fn(chunk_first, func(s interface{}) interface{} {
			return s.(CljsCoreIChunkedSeq).X_chunked_first_Arity1()
		})
	}(&AFn{})

	Chunk_rest = func(chunk_rest *AFn) *AFn {
		return Fn(chunk_rest, func(s interface{}) interface{} {
			return s.(CljsCoreIChunkedSeq).X_chunked_rest_Arity1()
		})
	}(&AFn{})

	Chunk_next = func(chunk_next *AFn) *AFn {
		return Fn(chunk_next, func(s interface{}) interface{} {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IChunkedNext", Str: "cljs.core/IChunkedNext", X_hash: float64(-1343796569), X_meta: nil}), s)) {
				return s.(CljsCoreIChunkedNext).X_chunked_next_Arity1()
			} else {
				return Seq.Arity1IQ(s.(CljsCoreIChunkedSeq).X_chunked_rest_Arity1())
			}
		})
	}(&AFn{})

	To_array = func(to_array *AFn) *AFn {
		return Fn(to_array, func(s interface{}) interface{} {
			{
				var ary = []interface{}{}
				_ = ary
				{
					var s___1 = s
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
		return Fn(to_array_2d, func(coll interface{}) interface{} {
			{
				var ret = make([]interface{}, int(Count.X_invoke_Arity1(coll).(float64)))
				_ = ret
				{
					var i_332 = float64(0)
					var xs_333 = Seq.Arity1IQ(coll)
					_, _ = i_332, xs_333
					for {
						if Truth_(xs_333) {
							ret[int(i_332)] = To_array.X_invoke_Arity1(First.X_invoke_Arity1(xs_333)).([]interface{})
							i_332, xs_333 = (i_332 + float64(1)), Next.Arity1IQ(xs_333)
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
		return Fn(int_array, func(size_or_seq interface{}) interface{} {
			if reflect.ValueOf(size_or_seq).Kind() == reflect.Float64 {
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
							var s___1 = s
							_, _ = i, s___1
							for {
								if Truth_(func() interface{} {
									var and__163__auto__ = s___1
									_ = and__163__auto__
									if Truth_(and__163__auto__) {
										return (i < size.(float64))
									} else {
										return and__163__auto__
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
						var n__1042__auto___334 = size
						_ = n__1042__auto___334
						{
							var i_335 = float64(0)
							_ = i_335
							for {
								if i_335 < n__1042__auto___334.(float64) {
									a[int(i_335)] = init_val_or_seq
									i_335 = (i_335 + float64(1))
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
		return Fn(long_array, func(size_or_seq interface{}) interface{} {
			if reflect.ValueOf(size_or_seq).Kind() == reflect.Float64 {
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
							var s___1 = s
							_, _ = i, s___1
							for {
								if Truth_(func() interface{} {
									var and__163__auto__ = s___1
									_ = and__163__auto__
									if Truth_(and__163__auto__) {
										return (i < size.(float64))
									} else {
										return and__163__auto__
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
						var n__1042__auto___336 = size
						_ = n__1042__auto___336
						{
							var i_337 = float64(0)
							_ = i_337
							for {
								if i_337 < n__1042__auto___336.(float64) {
									a[int(i_337)] = init_val_or_seq
									i_337 = (i_337 + float64(1))
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
		return Fn(double_array, func(size_or_seq interface{}) interface{} {
			if reflect.ValueOf(size_or_seq).Kind() == reflect.Float64 {
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
							var s___1 = s
							_, _ = i, s___1
							for {
								if Truth_(func() interface{} {
									var and__163__auto__ = s___1
									_ = and__163__auto__
									if Truth_(and__163__auto__) {
										return (i < size.(float64))
									} else {
										return and__163__auto__
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
						var n__1042__auto___338 = size
						_ = n__1042__auto___338
						{
							var i_339 = float64(0)
							_ = i_339
							for {
								if i_339 < n__1042__auto___338.(float64) {
									a[int(i_339)] = init_val_or_seq
									i_339 = (i_339 + float64(1))
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
		return Fn(object_array, func(size_or_seq interface{}) interface{} {
			if reflect.ValueOf(size_or_seq).Kind() == reflect.Float64 {
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
							var s___1 = s
							_, _ = i, s___1
							for {
								if Truth_(func() interface{} {
									var and__163__auto__ = s___1
									_ = and__163__auto__
									if Truth_(and__163__auto__) {
										return (i < size.(float64))
									} else {
										return and__163__auto__
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
						var n__1042__auto___340 = size
						_ = n__1042__auto___340
						{
							var i_341 = float64(0)
							_ = i_341
							for {
								if i_341 < n__1042__auto___340.(float64) {
									a[int(i_341)] = init_val_or_seq
									i_341 = (i_341 + float64(1))
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
		return Fn(bounded_count, func(s interface{}, n interface{}) interface{} {
			if Counted_QMARK_.Arity1IB(s) {
				return Count.X_invoke_Arity1(s).(float64)
			} else {
				{
					var s___1 = s
					var i = n
					var sum = float64(0)
					_, _, _ = s___1, i, sum
					for {
						if Truth_(func() interface{} {
							var and__163__auto__ = (i.(float64) > float64(0))
							_ = and__163__auto__
							if Truth_(and__163__auto__) {
								return Seq.Arity1IQ(s___1)
							} else {
								return and__163__auto__
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
		return Fn(spread, func(arglist interface{}) interface{} {
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
		return Fn(concat, func() interface{} {
			return (&CljsCoreLazySeq{nil, func(G__342 *AFn) *AFn {
				return Fn(G__342, func() interface{} {
					return nil
				})
			}(&AFn{}), nil, nil})
		}, func(x interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__343 *AFn) *AFn {
				return Fn(G__343, func() interface{} {
					return x
				})
			}(&AFn{}), nil, nil})
		}, func(x interface{}, y interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__344 *AFn) *AFn {
				return Fn(G__344, func() interface{} {
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
			var zs = Array_seq.X_invoke_Arity1(x_y_zs__[2:])
			_, _, _ = x, y, zs
			{
				var cat = func(cat *AFn) *AFn {
					return Fn(cat, func(xys interface{}, zs___1 interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__345 *AFn) *AFn {
							return Fn(G__345, func() interface{} {
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
		return Fn(list_STAR_, func(args interface{}) interface{} {
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
			var more = Array_seq.X_invoke_Arity1(a_b_c_d_more__[4:])
			_, _, _, _, _ = a, b, c, d, more
			return Cons.X_invoke_Arity2(a, Cons.X_invoke_Arity2(b, Cons.X_invoke_Arity2(c, Cons.X_invoke_Arity2(d, Spread.X_invoke_Arity1(more)).(*CljsCoreCons)).(*CljsCoreCons)).(*CljsCoreCons)).(*CljsCoreCons)
		})
	}(&AFn{})

	Transient = func(transient *AFn) *AFn {
		return Fn(transient, func(coll interface{}) interface{} {
			return coll.(CljsCoreIEditableCollection).X_as_transient_Arity1()
		})
	}(&AFn{})

	Persistent_BANG_ = func(persistent_BANG_ *AFn) *AFn {
		return Fn(persistent_BANG_, func(tcoll interface{}) interface{} {
			return tcoll.(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
		})
	}(&AFn{})

	Conj_BANG_ = func(conj_BANG_ *AFn) *AFn {
		return Fn(conj_BANG_, func() interface{} {
			return Transient.X_invoke_Arity1(CljsCorePersistentVector_EMPTY)
		}, func(coll interface{}) interface{} {
			return coll
		}, func(tcoll interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientCollection).X_conj_BANG__Arity2(val)
		}, func(tcoll_val_vals__ ...interface{}) interface{} {
			var tcoll = tcoll_val_vals__[0]
			var val = tcoll_val_vals__[1]
			var vals = Array_seq.X_invoke_Arity1(tcoll_val_vals__[2:])
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
		return Fn(assoc_BANG_, func(tcoll interface{}, key interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3(key, val)
		}, func(tcoll_key_val_kvs__ ...interface{}) interface{} {
			var tcoll = tcoll_key_val_kvs__[0]
			var key = tcoll_key_val_kvs__[1]
			var val = tcoll_key_val_kvs__[2]
			var kvs = Array_seq.X_invoke_Arity1(tcoll_key_val_kvs__[3:])
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
		return Fn(dissoc_BANG_, func(tcoll interface{}, key interface{}) interface{} {
			return tcoll.(CljsCoreITransientMap).X_dissoc_BANG__Arity2(key)
		}, func(tcoll_key_ks__ ...interface{}) interface{} {
			var tcoll = tcoll_key_ks__[0]
			var key = tcoll_key_ks__[1]
			var ks = Array_seq.X_invoke_Arity1(tcoll_key_ks__[2:])
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
		return Fn(pop_BANG_, func(tcoll interface{}) interface{} {
			return tcoll.(CljsCoreITransientVector).X_pop_BANG__Arity1()
		})
	}(&AFn{})

	Disj_BANG_ = func(disj_BANG_ *AFn) *AFn {
		return Fn(disj_BANG_, func(tcoll interface{}, val interface{}) interface{} {
			return tcoll.(CljsCoreITransientSet).X_disjoin_BANG__Arity2(val)
		}, func(tcoll_val_vals__ ...interface{}) interface{} {
			var tcoll = tcoll_val_vals__[0]
			var val = tcoll_val_vals__[1]
			var vals = Array_seq.X_invoke_Arity1(tcoll_val_vals__[2:])
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
		return Fn(vary_meta, func(obj interface{}, f interface{}) interface{} {
			return With_meta.X_invoke_Arity2(obj, f.(CljsCoreIFn).X_invoke_Arity1(Meta.X_invoke_Arity1(obj)))
		}, func(obj interface{}, f interface{}, a interface{}) interface{} {
			return With_meta.X_invoke_Arity2(obj, f.(CljsCoreIFn).X_invoke_Arity2(Meta.X_invoke_Arity1(obj), a))
		}, func(obj interface{}, f interface{}, a interface{}, b interface{}) interface{} {
			return With_meta.X_invoke_Arity2(obj, f.(CljsCoreIFn).X_invoke_Arity3(Meta.X_invoke_Arity1(obj), a, b))
		}, func(obj interface{}, f interface{}, a interface{}, b interface{}, c interface{}) interface{} {
			return With_meta.X_invoke_Arity2(obj, f.(CljsCoreIFn).X_invoke_Arity4(Meta.X_invoke_Arity1(obj), a, b, c))
		}, func(obj interface{}, f interface{}, a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
			return With_meta.X_invoke_Arity2(obj, f.(CljsCoreIFn).X_invoke_Arity5(Meta.X_invoke_Arity1(obj), a, b, c, d))
		}, func(obj_f_a_b_c_d_args__ ...interface{}) interface{} {
			var obj = obj_f_a_b_c_d_args__[0]
			var f = obj_f_a_b_c_d_args__[1]
			var a = obj_f_a_b_c_d_args__[2]
			var b = obj_f_a_b_c_d_args__[3]
			var c = obj_f_a_b_c_d_args__[4]
			var d = obj_f_a_b_c_d_args__[5]
			var args = Array_seq.X_invoke_Arity1(obj_f_a_b_c_d_args__[6:])
			_, _, _, _, _, _, _ = obj, f, a, b, c, d, args
			return With_meta.X_invoke_Arity2(obj, Apply.X_invoke_ArityVariadic(f, Meta.X_invoke_Arity1(obj), a, b, c, d, args))
		})
	}(&AFn{})

	Not_EQ_ = func(not_EQ_ *AFn) *AFn {
		return Fn(not_EQ_, func(x interface{}) bool {
			return false
		}, func(x interface{}, y interface{}) bool {
			return !(X_EQ_.Arity2IIB(x, y))
		}, func(x_y_more__ ...interface{}) interface{} {
			var x = x_y_more__[0]
			var y = x_y_more__[1]
			var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
			_, _, _ = x, y, more
			return Not.Arity1IB(Apply.X_invoke_Arity4(X_EQ_, x, y, more))
		})
	}(&AFn{})

	Not_empty = func(not_empty *AFn) *AFn {
		return Fn(not_empty, func(coll interface{}) interface{} {
			if Truth_(Seq.Arity1IQ(coll)) {
				return coll
			} else {
				return nil
			}
		})
	}(&AFn{})

	Nil_iter = func(nil_iter *AFn) *AFn {
		return Fn(nil_iter, func() interface{} {
			X__GT_t349 = func(__GT_t349 *AFn) *AFn {
				return Fn(__GT_t349, func(nil_iter___1 interface{}, meta350 interface{}) interface{} {
					return (&CljsCoreT349{nil_iter___1, meta350})
				})
			}(&AFn{})

			return (&CljsCoreT349{nil_iter, nil})
		})
	}(&AFn{})

	X__GT_StringIter = func(__GT_StringIter *AFn) *AFn {
		return Fn(__GT_StringIter, func(s interface{}, i interface{}) interface{} {
			return (&CljsCoreStringIter{s, i})
		})
	}(&AFn{})

	String_iter = func(string_iter *AFn) *AFn {
		return Fn(string_iter, func(x interface{}) interface{} {
			return (&CljsCoreStringIter{x, float64(0)})
		})
	}(&AFn{})

	X__GT_ArrayIter = func(__GT_ArrayIter *AFn) *AFn {
		return Fn(__GT_ArrayIter, func(arr interface{}, i interface{}) interface{} {
			return (&CljsCoreArrayIter{arr, i})
		})
	}(&AFn{})

	Array_iter = func(array_iter *AFn) *AFn {
		return Fn(array_iter, func(x interface{}) interface{} {
			return (&CljsCoreArrayIter{x, float64(0)})
		})
	}(&AFn{})

	INIT = map[string]interface{}{}

	START = map[string]interface{}{}

	X__GT_SeqIter = func(__GT_SeqIter *AFn) *AFn {
		return Fn(__GT_SeqIter, func(_seq interface{}, _next interface{}) interface{} {
			return (&CljsCoreSeqIter{_seq, _next})
		})
	}(&AFn{})

	Seq_iter = func(seq_iter *AFn) *AFn {
		return Fn(seq_iter, func(coll interface{}) interface{} {
			return (&CljsCoreSeqIter{INIT, coll})
		})
	}(&AFn{})

	Iter = func(iter *AFn) *AFn {
		return Fn(iter, func(coll interface{}) interface{} {
			if Nil_(coll) {
				return Nil_iter.X_invoke_Arity0().(*CljsCoreT349)
			} else {
				if reflect.ValueOf(coll).Kind() == reflect.String {
					return String_iter.X_invoke_Arity1(coll).(*CljsCoreStringIter)
				} else {
					if reflect.ValueOf(coll).Kind() == reflect.Slice {
						return Array_iter.X_invoke_Arity1(coll).(*CljsCoreArrayIter)
					} else {
						return Seq_iter.X_invoke_Arity1(coll).(*CljsCoreSeqIter)

					}
				}
			}
		})
	}(&AFn{})

	Lazy_transformer = func(lazy_transformer *AFn) *AFn {
		return Fn(lazy_transformer, func(stepper interface{}) interface{} {
			return (&CljsCoreLazyTransformer{stepper, nil, nil, nil})
		})
	}(&AFn{})

	X__GT_Stepper = func(__GT_Stepper *AFn) *AFn {
		return Fn(__GT_Stepper, func(xform interface{}, iter interface{}) interface{} {
			return (&CljsCoreStepper{xform, iter})
		})
	}(&AFn{})

	Stepper = func(stepper *AFn) *AFn {
		return Fn(stepper, func(xform interface{}, iter interface{}) interface{} {
			{
				var stepfn *AFn
				stepfn = func(stepfn *AFn) *AFn {
					return Fn(stepfn, func(result interface{}) interface{} {
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
				return (&CljsCoreStepper{xform.(CljsCoreIFn).X_invoke_Arity1(stepfn), iter})
			}
		})
	}(&AFn{})

	X__GT_MultiStepper = func(__GT_MultiStepper *AFn) *AFn {
		return Fn(__GT_MultiStepper, func(xform interface{}, iters interface{}, nexts interface{}) interface{} {
			return (&CljsCoreMultiStepper{xform, iters, nexts})
		})
	}(&AFn{})

	Multi_stepper = func(multi_stepper *AFn) *AFn {
		return Fn(multi_stepper, func(xform interface{}, iters interface{}) interface{} {
			return multi_stepper.X_invoke_Arity3(xform, iters, make([]interface{}, int(float64(len(iters.([]interface{}))))))
		}, func(xform interface{}, iters interface{}, nexts interface{}) interface{} {
			{
				var stepfn *AFn
				stepfn = func(stepfn *AFn) *AFn {
					return Fn(stepfn, func(result interface{}) interface{} {
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
				return (&CljsCoreMultiStepper{xform.(CljsCoreIFn).X_invoke_Arity1(stepfn), iters, nexts})
			}
		})
	}(&AFn{})

	X__GT_LazyTransformer = func(__GT_LazyTransformer *AFn) *AFn {
		return Fn(__GT_LazyTransformer, func(stepper interface{}, first interface{}, rest interface{}, meta interface{}) interface{} {
			return (&CljsCoreLazyTransformer{stepper, first, rest, meta})
		})
	}(&AFn{})

	Sequence = func(sequence *AFn) *AFn {
		return Fn(sequence, func(coll interface{}) interface{} {
			if Seq_QMARK_.Arity1IB(coll) {
				return coll
			} else {
				{
					var or__175__auto__ = Seq.Arity1IQ(coll)
					_ = or__175__auto__
					if Truth_(or__175__auto__) {
						return or__175__auto__
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
			var colls = Array_seq.X_invoke_Arity1(xform_coll_colls__[2:])
			_, _, _ = xform, coll, colls
			return CljsCoreLazyTransformer_CreateMulti.X_invoke_Arity2(xform, To_array.X_invoke_Arity1(Cons.X_invoke_Arity2(coll, colls).(*CljsCoreCons)).([]interface{}))
		})
	}(&AFn{})

	Every_QMARK_ = func(every_QMARK_ *AFn) *AFn {
		return Fn(every_QMARK_, func(pred interface{}, coll interface{}) bool {
			for {
				if Nil_(Seq.Arity1IQ(coll)) {
					return true
				} else {
					if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(coll))) {
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
		return Fn(not_every_QMARK_, func(pred interface{}, coll interface{}) bool {
			return !(Every_QMARK_.Arity2IIB(pred, coll))
		})
	}(&AFn{})

	Some = func(some *AFn) *AFn {
		return Fn(some, func(pred interface{}, coll interface{}) interface{} {
			for {
				if Truth_(Seq.Arity1IQ(coll)) {
					{
						var or__175__auto__ = pred.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(coll))
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
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
		return Fn(not_any_QMARK_, func(pred interface{}, coll interface{}) bool {
			return Not.Arity1IB(Some.X_invoke_Arity2(pred, coll))
		})
	}(&AFn{})

	Even_QMARK_ = func(even_QMARK_ *AFn) *AFn {
		return Fn(even_QMARK_, func(n interface{}) bool {
			if Integer_QMARK_.Arity1IB(n) {
				return (float64(int(n.(float64))&int(float64(1))) == float64(0))
			} else {
				panic((&js.Error{("Argument must be an integer: " + Str.X_invoke_Arity1(n).(string))}))
			}
		})
	}(&AFn{})

	Odd_QMARK_ = func(odd_QMARK_ *AFn) *AFn {
		return Fn(odd_QMARK_, func(n interface{}) bool {
			return !(Even_QMARK_.Arity1IB(n))
		})
	}(&AFn{})

	Constantly = func(constantly *AFn) *AFn {
		return Fn(constantly, func(x interface{}) interface{} {
			return func(G__367 *AFn) *AFn {
				return Fn(G__367, func(args__ ...interface{}) interface{} {
					var args = Array_seq.X_invoke_Arity1(args__[0:])
					_ = args
					return x
				})
			}(&AFn{})
		})
	}(&AFn{})

	Comp = func(comp *AFn) *AFn {
		return Fn(comp, func() interface{} {
			return Identity
		}, func(f interface{}) interface{} {
			return f
		}, func(f interface{}, g interface{}) interface{} {
			return func(G__368 *AFn) *AFn {
				return Fn(G__368, func() interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity0())
				}, func(x interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(x))
				}, func(x interface{}, y interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity2(x, y))
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity3(x, y, z))
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					return f.(CljsCoreIFn).X_invoke_Arity1(Apply.X_invoke_Arity5(g, x, y, z, args))
				})
			}(&AFn{})
		}, func(f interface{}, g interface{}, h interface{}) interface{} {
			return func(G__369 *AFn) *AFn {
				return Fn(G__369, func() interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(h.(CljsCoreIFn).X_invoke_Arity0()))
				}, func(x interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(h.(CljsCoreIFn).X_invoke_Arity1(x)))
				}, func(x interface{}, y interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(h.(CljsCoreIFn).X_invoke_Arity2(x, y)))
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(h.(CljsCoreIFn).X_invoke_Arity3(x, y, z)))
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(Apply.X_invoke_Arity5(h, x, y, z, args)))
				})
			}(&AFn{})
		}, func(f1_f2_f3_fs__ ...interface{}) interface{} {
			var f1 = f1_f2_f3_fs__[0]
			var f2 = f1_f2_f3_fs__[1]
			var f3 = f1_f2_f3_fs__[2]
			var fs = Array_seq.X_invoke_Arity1(f1_f2_f3_fs__[3:])
			_, _, _, _ = f1, f2, f3, fs
			{
				var fs___1 = Reverse.X_invoke_Arity1(List_STAR_.X_invoke_Arity4(f1, f2, f3, fs).(*CljsCoreCons))
				_ = fs___1
				return func(G__370 *AFn, fs___1 interface{}) *AFn {
					return Fn(G__370, func(args__ ...interface{}) interface{} {
						var args = Array_seq.X_invoke_Arity1(args__[0:])
						_ = args
						{
							var ret = Apply.X_invoke_Arity2(First.X_invoke_Arity1(fs___1), args)
							var fs___2 = Next.Arity1IQ(fs___1)
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
		return Fn(partial, func(f interface{}) interface{} {
			return f
		}, func(f interface{}, arg1 interface{}) interface{} {
			return func(G__371 *AFn) *AFn {
				return Fn(G__371, func(args__ ...interface{}) interface{} {
					var args = Array_seq.X_invoke_Arity1(args__[0:])
					_ = args
					return Apply.X_invoke_Arity3(f, arg1, args)
				})
			}(&AFn{})
		}, func(f interface{}, arg1 interface{}, arg2 interface{}) interface{} {
			return func(G__372 *AFn) *AFn {
				return Fn(G__372, func(args__ ...interface{}) interface{} {
					var args = Array_seq.X_invoke_Arity1(args__[0:])
					_ = args
					return Apply.X_invoke_Arity4(f, arg1, arg2, args)
				})
			}(&AFn{})
		}, func(f interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} {
			return func(G__373 *AFn) *AFn {
				return Fn(G__373, func(args__ ...interface{}) interface{} {
					var args = Array_seq.X_invoke_Arity1(args__[0:])
					_ = args
					return Apply.X_invoke_Arity5(f, arg1, arg2, arg3, args)
				})
			}(&AFn{})
		}, func(f_arg1_arg2_arg3_more__ ...interface{}) interface{} {
			var f = f_arg1_arg2_arg3_more__[0]
			var arg1 = f_arg1_arg2_arg3_more__[1]
			var arg2 = f_arg1_arg2_arg3_more__[2]
			var arg3 = f_arg1_arg2_arg3_more__[3]
			var more = Array_seq.X_invoke_Arity1(f_arg1_arg2_arg3_more__[4:])
			_, _, _, _, _ = f, arg1, arg2, arg3, more
			return func(G__374 *AFn) *AFn {
				return Fn(G__374, func(args__ ...interface{}) interface{} {
					var args = Array_seq.X_invoke_Arity1(args__[0:])
					_ = args
					return Apply.X_invoke_Arity5(f, arg1, arg2, arg3, Concat.X_invoke_Arity2(more, args).(*CljsCoreLazySeq))
				})
			}(&AFn{})
		})
	}(&AFn{})

	Fnil = func(fnil *AFn) *AFn {
		return Fn(fnil, func(f interface{}, x interface{}) interface{} {
			return func(G__375 *AFn) *AFn {
				return Fn(G__375, func(a interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity1(func() interface{} {
						if Nil_(a) {
							return x
						} else {
							return a
						}
					}())
				}, func(a interface{}, b interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity2(func() interface{} {
						if Nil_(a) {
							return x
						} else {
							return a
						}
					}(), b)
				}, func(a interface{}, b interface{}, c interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity3(func() interface{} {
						if Nil_(a) {
							return x
						} else {
							return a
						}
					}(), b, c)
				}, func(a_b_c_ds__ ...interface{}) interface{} {
					var a = a_b_c_ds__[0]
					var b = a_b_c_ds__[1]
					var c = a_b_c_ds__[2]
					var ds = Array_seq.X_invoke_Arity1(a_b_c_ds__[3:])
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
			return func(G__376 *AFn) *AFn {
				return Fn(G__376, func(a interface{}, b interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity2(func() interface{} {
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
					}())
				}, func(a interface{}, b interface{}, c interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity3(func() interface{} {
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
					}(), c)
				}, func(a_b_c_ds__ ...interface{}) interface{} {
					var a = a_b_c_ds__[0]
					var b = a_b_c_ds__[1]
					var c = a_b_c_ds__[2]
					var ds = Array_seq.X_invoke_Arity1(a_b_c_ds__[3:])
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
			return func(G__377 *AFn) *AFn {
				return Fn(G__377, func(a interface{}, b interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity2(func() interface{} {
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
					}())
				}, func(a interface{}, b interface{}, c interface{}) interface{} {
					return f.(CljsCoreIFn).X_invoke_Arity3(func() interface{} {
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
					}())
				}, func(a_b_c_ds__ ...interface{}) interface{} {
					var a = a_b_c_ds__[0]
					var b = a_b_c_ds__[1]
					var c = a_b_c_ds__[2]
					var ds = Array_seq.X_invoke_Arity1(a_b_c_ds__[3:])
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
		return Fn(map_indexed, func(f interface{}, coll interface{}) interface{} {
			{
				var mapi *AFn
				mapi = func(mapi *AFn) *AFn {
					return Fn(mapi, func(idx interface{}, coll___1 interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__378 *AFn) *AFn {
							return Fn(G__378, func() interface{} {
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
														var n__1042__auto___379 = size
														_ = n__1042__auto___379
														{
															var i_380 = float64(0)
															_ = i_380
															for {
																if i_380 < n__1042__auto___379 {
																	Chunk_append.X_invoke_Arity2(b, f.(CljsCoreIFn).X_invoke_Arity2((idx.(float64)+i_380), c.(CljsCoreIIndexed).X_nth_Arity2(i_380)))
																	i_380 = (i_380 + float64(1))
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
												return Cons.X_invoke_Arity2(f.(CljsCoreIFn).X_invoke_Arity2(idx, First.X_invoke_Arity1(s)), mapi.X_invoke_Arity2((idx.(float64)+float64(1)), Rest.Arity1IQ(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
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
		return Fn(keep, func(f interface{}) interface{} {
			return func(G__381 *AFn) *AFn {
				return Fn(G__381, func(f1 interface{}) interface{} {
					return func(G__382 *AFn) *AFn {
						return Fn(G__382, func() interface{} {
							return f1.(CljsCoreIFn).X_invoke_Arity0()
						}, func(result interface{}) interface{} {
							return f1.(CljsCoreIFn).X_invoke_Arity1(result)
						}, func(result interface{}, input interface{}) interface{} {
							{
								var v = f.(CljsCoreIFn).X_invoke_Arity1(input)
								_ = v
								if Nil_(v) {
									return result
								} else {
									return f1.(CljsCoreIFn).X_invoke_Arity2(result, v)
								}
							}
						})
					}(&AFn{})
				})
			}(&AFn{})
		}, func(f interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__383 *AFn) *AFn {
				return Fn(G__383, func() interface{} {
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
											var n__1042__auto___384 = size
											_ = n__1042__auto___384
											{
												var i_385 = float64(0)
												_ = i_385
												for {
													if i_385 < n__1042__auto___384 {
														{
															var x_386 = f.(CljsCoreIFn).X_invoke_Arity1(c.(CljsCoreIIndexed).X_nth_Arity2(i_385))
															_ = x_386
															if Nil_(x_386) {
															} else {
																Chunk_append.X_invoke_Arity2(b, x_386)
															}
														}
														i_385 = (i_385 + float64(1))
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
										var x = f.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(s))
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
		return Fn(__GT_Atom, func(state interface{}, meta interface{}, validator interface{}, watches interface{}) interface{} {
			return (&CljsCoreAtom{state, meta, validator, watches})
		})
	}(&AFn{})

	Atom = func(atom *AFn) *AFn {
		return Fn(atom, func(x interface{}) interface{} {
			return (&CljsCoreAtom{x, nil, nil, nil})
		}, func(x_p__399__ ...interface{}) interface{} {
			var x = x_p__399__[0]
			var p__399 = Array_seq.X_invoke_Arity1(x_p__399__[1:])
			_, _ = x, p__399
			{
				var map__401 = p__399
				var map__401___1 = func() interface{} {
					if Seq_QMARK_.Arity1IB(map__401) {
						return Apply.X_invoke_Arity2(Hash_map, map__401)
					} else {
						return map__401
					}
				}()
				var validator = Get.X_invoke_Arity2(map__401___1, (&CljsCoreKeyword{Ns: nil, Name: "validator", Fqn: "validator", X_hash: float64(-1966190681)}))
				var meta = Get.X_invoke_Arity2(map__401___1, (&CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}))
				_, _, _, _ = map__401, map__401___1, validator, meta
				return (&CljsCoreAtom{x, meta, validator, nil})
			}
		})
	}(&AFn{})

	Reset_BANG_ = func(reset_BANG_ *AFn) *AFn {
		return Fn(reset_BANG_, func(a interface{}, new_value interface{}) interface{} {
			if func() bool { _, instanceof := a.(*CljsCoreAtom); return instanceof }() {
				{
					var validate = Native_get_instance_field.X_invoke_Arity2(a, "Validator")
					_ = validate
					if Nil_(validate) {
					} else {
						if Truth_(validate.(CljsCoreIFn).X_invoke_Arity1(new_value)) {
						} else {
							panic((&js.Error{("Assert failed: Validator rejected reference state\n" + Str.X_invoke_Arity1(Pr_str.X_invoke_ArityVariadic(List.X_invoke_ArityVariadic((&CljsCoreSymbol{Ns: nil, Name: "validate", Str: "validate", X_hash: float64(1439230700), X_meta: nil}), (&CljsCoreSymbol{Ns: nil, Name: "new-value", Str: "new-value", X_hash: float64(-1567397401), X_meta: nil})).(*CljsCoreList)).(string)).(string))}))
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
		return Fn(swap_BANG_, func(a interface{}, f interface{}) interface{} {
			if func() bool { _, instanceof := a.(*CljsCoreAtom); return instanceof }() {
				return Reset_BANG_.X_invoke_Arity2(a, f.(CljsCoreIFn).X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(a, "State")))
			} else {
				return a.(CljsCoreISwap).X_swap_BANG__Arity2(f)
			}
		}, func(a interface{}, f interface{}, x interface{}) interface{} {
			if func() bool { _, instanceof := a.(*CljsCoreAtom); return instanceof }() {
				return Reset_BANG_.X_invoke_Arity2(a, f.(CljsCoreIFn).X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(a, "State"), x))
			} else {
				return a.(CljsCoreISwap).X_swap_BANG__Arity3(f, x)
			}
		}, func(a interface{}, f interface{}, x interface{}, y interface{}) interface{} {
			if func() bool { _, instanceof := a.(*CljsCoreAtom); return instanceof }() {
				return Reset_BANG_.X_invoke_Arity2(a, f.(CljsCoreIFn).X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(a, "State"), x, y))
			} else {
				return a.(CljsCoreISwap).X_swap_BANG__Arity4(f, x, y)
			}
		}, func(a_f_x_y_more__ ...interface{}) interface{} {
			var a = a_f_x_y_more__[0]
			var f = a_f_x_y_more__[1]
			var x = a_f_x_y_more__[2]
			var y = a_f_x_y_more__[3]
			var more = Array_seq.X_invoke_Arity1(a_f_x_y_more__[4:])
			_, _, _, _, _ = a, f, x, y, more
			if func() bool { _, instanceof := a.(*CljsCoreAtom); return instanceof }() {
				return Reset_BANG_.X_invoke_Arity2(a, Apply.X_invoke_Arity5(f, Native_get_instance_field.X_invoke_Arity2(a, "State"), x, y, more))
			} else {
				return a.(CljsCoreISwap).X_swap_BANG__Arity5(f, x, y, more)
			}
		})
	}(&AFn{})

	Compare_and_set_BANG_ = func(compare_and_set_BANG_ *AFn) *AFn {
		return Fn(compare_and_set_BANG_, func(a interface{}, oldval interface{}, newval interface{}) interface{} {
			if X_EQ_.Arity2IIB(Native_get_instance_field.X_invoke_Arity2(a, "State"), oldval) {
				Reset_BANG_.X_invoke_Arity2(a, newval)
				return true
			} else {
				return false
			}
		})
	}(&AFn{})

	Set_validator_BANG_ = func(set_validator_BANG_ *AFn) *AFn {
		return Fn(set_validator_BANG_, func(iref interface{}, val interface{}) interface{} {
			return func() interface{} {
				var return__402 = val
				Native_set_instance_field.X_invoke_Arity3(iref, "Validator", return__402)
				return return__402
			}()
		})
	}(&AFn{})

	Get_validator = func(get_validator *AFn) *AFn {
		return Fn(get_validator, func(iref interface{}) interface{} {
			return Native_get_instance_field.X_invoke_Arity2(iref, "Validator")
		})
	}(&AFn{})

	Keep_indexed = func(keep_indexed *AFn) *AFn {
		return Fn(keep_indexed, func(f interface{}) interface{} {
			return func(G__403 *AFn) *AFn {
				return Fn(G__403, func(f1 interface{}) interface{} {
					{
						var ia = Atom.X_invoke_Arity1(float64(-1)).(*CljsCoreAtom)
						_ = ia
						return func(G__404 *AFn, ia *CljsCoreAtom) *AFn {
							return Fn(G__404, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0()
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result)
							}, func(result interface{}, input interface{}) interface{} {
								{
									var i = Swap_BANG_.X_invoke_Arity2(ia, Inc)
									var v = f.(CljsCoreIFn).X_invoke_Arity2(i, input)
									_, _ = i, v
									if Nil_(v) {
										return result
									} else {
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, v)
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
					return Fn(keepi, func(idx interface{}, coll___1 interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__405 *AFn) *AFn {
							return Fn(G__405, func() interface{} {
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
														var n__1042__auto___406 = size
														_ = n__1042__auto___406
														{
															var i_407 = float64(0)
															_ = i_407
															for {
																if i_407 < n__1042__auto___406 {
																	{
																		var x_408 = f.(CljsCoreIFn).X_invoke_Arity2((idx.(float64) + i_407), c.(CljsCoreIIndexed).X_nth_Arity2(i_407))
																		_ = x_408
																		if Nil_(x_408) {
																		} else {
																			Chunk_append.X_invoke_Arity2(b, x_408)
																		}
																	}
																	i_407 = (i_407 + float64(1))
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
													var x = f.(CljsCoreIFn).X_invoke_Arity2(idx, First.X_invoke_Arity1(s))
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
		return Fn(every_pred, func(p interface{}) interface{} {
			return func(ep1 *AFn) *AFn {
				return Fn(ep1, func() interface{} {
					return true
				}, func(x interface{}) interface{} {
					return Boolean.Arity1IB(p.(CljsCoreIFn).X_invoke_Arity1(x))
				}, func(x interface{}, y interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__163__auto__ = p.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = and__163__auto__
						if Truth_(and__163__auto__) {
							return p.(CljsCoreIFn).X_invoke_Arity1(y)
						} else {
							return and__163__auto__
						}
					}())
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__163__auto__ = p.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = and__163__auto__
						if Truth_(and__163__auto__) {
							{
								var and__163__auto_____1 = p.(CljsCoreIFn).X_invoke_Arity1(y)
								_ = and__163__auto_____1
								if Truth_(and__163__auto_____1) {
									return p.(CljsCoreIFn).X_invoke_Arity1(z)
								} else {
									return and__163__auto_____1
								}
							}
						} else {
							return and__163__auto__
						}
					}())
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					return Boolean.Arity1IB((Truth_(ep1.X_invoke_Arity3(x, y, z))) && (Every_QMARK_.Arity2IIB(p, args)))
				})
			}(&AFn{})
		}, func(p1 interface{}, p2 interface{}) interface{} {
			return func(ep2 *AFn) *AFn {
				return Fn(ep2, func() interface{} {
					return true
				}, func(x interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__163__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = and__163__auto__
						if Truth_(and__163__auto__) {
							return p2.(CljsCoreIFn).X_invoke_Arity1(x)
						} else {
							return and__163__auto__
						}
					}())
				}, func(x interface{}, y interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__163__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = and__163__auto__
						if Truth_(and__163__auto__) {
							{
								var and__163__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(y)
								_ = and__163__auto_____1
								if Truth_(and__163__auto_____1) {
									{
										var and__163__auto_____2 = p2.(CljsCoreIFn).X_invoke_Arity1(x)
										_ = and__163__auto_____2
										if Truth_(and__163__auto_____2) {
											return p2.(CljsCoreIFn).X_invoke_Arity1(y)
										} else {
											return and__163__auto_____2
										}
									}
								} else {
									return and__163__auto_____1
								}
							}
						} else {
							return and__163__auto__
						}
					}())
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__163__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = and__163__auto__
						if Truth_(and__163__auto__) {
							{
								var and__163__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(y)
								_ = and__163__auto_____1
								if Truth_(and__163__auto_____1) {
									{
										var and__163__auto_____2 = p1.(CljsCoreIFn).X_invoke_Arity1(z)
										_ = and__163__auto_____2
										if Truth_(and__163__auto_____2) {
											{
												var and__163__auto_____3 = p2.(CljsCoreIFn).X_invoke_Arity1(x)
												_ = and__163__auto_____3
												if Truth_(and__163__auto_____3) {
													{
														var and__163__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y)
														_ = and__163__auto_____4
														if Truth_(and__163__auto_____4) {
															return p2.(CljsCoreIFn).X_invoke_Arity1(z)
														} else {
															return and__163__auto_____4
														}
													}
												} else {
													return and__163__auto_____3
												}
											}
										} else {
											return and__163__auto_____2
										}
									}
								} else {
									return and__163__auto_____1
								}
							}
						} else {
							return and__163__auto__
						}
					}())
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					return Boolean.Arity1IB((Truth_(ep2.X_invoke_Arity3(x, y, z))) && (Every_QMARK_.Arity2IIB(func(G__415 *AFn) *AFn {
						return Fn(G__415, func(p1__409_SHARP_ interface{}) interface{} {
							{
								var and__163__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(p1__409_SHARP_)
								_ = and__163__auto__
								if Truth_(and__163__auto__) {
									return p2.(CljsCoreIFn).X_invoke_Arity1(p1__409_SHARP_)
								} else {
									return and__163__auto__
								}
							}
						})
					}(&AFn{}), args)))
				})
			}(&AFn{})
		}, func(p1 interface{}, p2 interface{}, p3 interface{}) interface{} {
			return func(ep3 *AFn) *AFn {
				return Fn(ep3, func() interface{} {
					return true
				}, func(x interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__163__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = and__163__auto__
						if Truth_(and__163__auto__) {
							{
								var and__163__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x)
								_ = and__163__auto_____1
								if Truth_(and__163__auto_____1) {
									return p3.(CljsCoreIFn).X_invoke_Arity1(x)
								} else {
									return and__163__auto_____1
								}
							}
						} else {
							return and__163__auto__
						}
					}())
				}, func(x interface{}, y interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__163__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = and__163__auto__
						if Truth_(and__163__auto__) {
							{
								var and__163__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x)
								_ = and__163__auto_____1
								if Truth_(and__163__auto_____1) {
									{
										var and__163__auto_____2 = p3.(CljsCoreIFn).X_invoke_Arity1(x)
										_ = and__163__auto_____2
										if Truth_(and__163__auto_____2) {
											{
												var and__163__auto_____3 = p1.(CljsCoreIFn).X_invoke_Arity1(y)
												_ = and__163__auto_____3
												if Truth_(and__163__auto_____3) {
													{
														var and__163__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y)
														_ = and__163__auto_____4
														if Truth_(and__163__auto_____4) {
															return p3.(CljsCoreIFn).X_invoke_Arity1(y)
														} else {
															return and__163__auto_____4
														}
													}
												} else {
													return and__163__auto_____3
												}
											}
										} else {
											return and__163__auto_____2
										}
									}
								} else {
									return and__163__auto_____1
								}
							}
						} else {
							return and__163__auto__
						}
					}())
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return Boolean.Arity1IB(func() interface{} {
						var and__163__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = and__163__auto__
						if Truth_(and__163__auto__) {
							{
								var and__163__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x)
								_ = and__163__auto_____1
								if Truth_(and__163__auto_____1) {
									{
										var and__163__auto_____2 = p3.(CljsCoreIFn).X_invoke_Arity1(x)
										_ = and__163__auto_____2
										if Truth_(and__163__auto_____2) {
											{
												var and__163__auto_____3 = p1.(CljsCoreIFn).X_invoke_Arity1(y)
												_ = and__163__auto_____3
												if Truth_(and__163__auto_____3) {
													{
														var and__163__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y)
														_ = and__163__auto_____4
														if Truth_(and__163__auto_____4) {
															{
																var and__163__auto_____5 = p3.(CljsCoreIFn).X_invoke_Arity1(y)
																_ = and__163__auto_____5
																if Truth_(and__163__auto_____5) {
																	{
																		var and__163__auto_____6 = p1.(CljsCoreIFn).X_invoke_Arity1(z)
																		_ = and__163__auto_____6
																		if Truth_(and__163__auto_____6) {
																			{
																				var and__163__auto_____7 = p2.(CljsCoreIFn).X_invoke_Arity1(z)
																				_ = and__163__auto_____7
																				if Truth_(and__163__auto_____7) {
																					return p3.(CljsCoreIFn).X_invoke_Arity1(z)
																				} else {
																					return and__163__auto_____7
																				}
																			}
																		} else {
																			return and__163__auto_____6
																		}
																	}
																} else {
																	return and__163__auto_____5
																}
															}
														} else {
															return and__163__auto_____4
														}
													}
												} else {
													return and__163__auto_____3
												}
											}
										} else {
											return and__163__auto_____2
										}
									}
								} else {
									return and__163__auto_____1
								}
							}
						} else {
							return and__163__auto__
						}
					}())
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					return Boolean.Arity1IB((Truth_(ep3.X_invoke_Arity3(x, y, z))) && (Every_QMARK_.Arity2IIB(func(G__416 *AFn) *AFn {
						return Fn(G__416, func(p1__410_SHARP_ interface{}) interface{} {
							{
								var and__163__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(p1__410_SHARP_)
								_ = and__163__auto__
								if Truth_(and__163__auto__) {
									{
										var and__163__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(p1__410_SHARP_)
										_ = and__163__auto_____1
										if Truth_(and__163__auto_____1) {
											return p3.(CljsCoreIFn).X_invoke_Arity1(p1__410_SHARP_)
										} else {
											return and__163__auto_____1
										}
									}
								} else {
									return and__163__auto__
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
			var ps = Array_seq.X_invoke_Arity1(p1_p2_p3_ps__[3:])
			_, _, _, _ = p1, p2, p3, ps
			{
				var ps___1 = List_STAR_.X_invoke_Arity4(p1, p2, p3, ps).(*CljsCoreCons)
				_ = ps___1
				return func(epn *AFn, ps___1 *CljsCoreCons) *AFn {
					return Fn(epn, func() interface{} {
						return true
					}, func(x interface{}) interface{} {
						return Every_QMARK_.Arity2IIB(func(G__417 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__417, func(p1__411_SHARP_ interface{}) interface{} {
								return p1__411_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x)
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x interface{}, y interface{}) interface{} {
						return Every_QMARK_.Arity2IIB(func(G__418 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__418, func(p1__412_SHARP_ interface{}) interface{} {
								{
									var and__163__auto__ = p1__412_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x)
									_ = and__163__auto__
									if Truth_(and__163__auto__) {
										return p1__412_SHARP_.(CljsCoreIFn).X_invoke_Arity1(y)
									} else {
										return and__163__auto__
									}
								}
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return Every_QMARK_.Arity2IIB(func(G__419 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__419, func(p1__413_SHARP_ interface{}) interface{} {
								{
									var and__163__auto__ = p1__413_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x)
									_ = and__163__auto__
									if Truth_(and__163__auto__) {
										{
											var and__163__auto_____1 = p1__413_SHARP_.(CljsCoreIFn).X_invoke_Arity1(y)
											_ = and__163__auto_____1
											if Truth_(and__163__auto_____1) {
												return p1__413_SHARP_.(CljsCoreIFn).X_invoke_Arity1(z)
											} else {
												return and__163__auto_____1
											}
										}
									} else {
										return and__163__auto__
									}
								}
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x_y_z_args__ ...interface{}) interface{} {
						var x = x_y_z_args__[0]
						var y = x_y_z_args__[1]
						var z = x_y_z_args__[2]
						var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
						_, _, _, _ = x, y, z, args
						return Boolean.Arity1IB((Truth_(epn.X_invoke_Arity3(x, y, z))) && (Every_QMARK_.Arity2IIB(func(G__420 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__420, func(p1__414_SHARP_ interface{}) interface{} {
								return Every_QMARK_.Arity2IIB(p1__414_SHARP_, args)
							})
						}(&AFn{}, ps___1), ps___1)))
					})
				}(&AFn{}, ps___1)
			}
		})
	}(&AFn{})

	Some_fn = func(some_fn *AFn) *AFn {
		return Fn(some_fn, func(p interface{}) interface{} {
			return func(sp1 *AFn) *AFn {
				return Fn(sp1, func() interface{} {
					return nil
				}, func(x interface{}) interface{} {
					return p.(CljsCoreIFn).X_invoke_Arity1(x)
				}, func(x interface{}, y interface{}) interface{} {
					{
						var or__175__auto__ = p.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							return p.(CljsCoreIFn).X_invoke_Arity1(y)
						}
					}
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					{
						var or__175__auto__ = p.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							{
								var or__175__auto_____1 = p.(CljsCoreIFn).X_invoke_Arity1(y)
								_ = or__175__auto_____1
								if Truth_(or__175__auto_____1) {
									return or__175__auto_____1
								} else {
									return p.(CljsCoreIFn).X_invoke_Arity1(z)
								}
							}
						}
					}
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					{
						var or__175__auto__ = sp1.X_invoke_Arity3(x, y, z)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							return Some.X_invoke_Arity2(p, args)
						}
					}
				})
			}(&AFn{})
		}, func(p1 interface{}, p2 interface{}) interface{} {
			return func(sp2 *AFn) *AFn {
				return Fn(sp2, func() interface{} {
					return nil
				}, func(x interface{}) interface{} {
					{
						var or__175__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							return p2.(CljsCoreIFn).X_invoke_Arity1(x)
						}
					}
				}, func(x interface{}, y interface{}) interface{} {
					{
						var or__175__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							{
								var or__175__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(y)
								_ = or__175__auto_____1
								if Truth_(or__175__auto_____1) {
									return or__175__auto_____1
								} else {
									{
										var or__175__auto_____2 = p2.(CljsCoreIFn).X_invoke_Arity1(x)
										_ = or__175__auto_____2
										if Truth_(or__175__auto_____2) {
											return or__175__auto_____2
										} else {
											return p2.(CljsCoreIFn).X_invoke_Arity1(y)
										}
									}
								}
							}
						}
					}
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					{
						var or__175__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							{
								var or__175__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(y)
								_ = or__175__auto_____1
								if Truth_(or__175__auto_____1) {
									return or__175__auto_____1
								} else {
									{
										var or__175__auto_____2 = p1.(CljsCoreIFn).X_invoke_Arity1(z)
										_ = or__175__auto_____2
										if Truth_(or__175__auto_____2) {
											return or__175__auto_____2
										} else {
											{
												var or__175__auto_____3 = p2.(CljsCoreIFn).X_invoke_Arity1(x)
												_ = or__175__auto_____3
												if Truth_(or__175__auto_____3) {
													return or__175__auto_____3
												} else {
													{
														var or__175__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y)
														_ = or__175__auto_____4
														if Truth_(or__175__auto_____4) {
															return or__175__auto_____4
														} else {
															return p2.(CljsCoreIFn).X_invoke_Arity1(z)
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
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					{
						var or__175__auto__ = sp2.X_invoke_Arity3(x, y, z)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							return Some.X_invoke_Arity2(func(G__427 *AFn, or__175__auto__ interface{}) *AFn {
								return Fn(G__427, func(p1__421_SHARP_ interface{}) interface{} {
									{
										var or__175__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(p1__421_SHARP_)
										_ = or__175__auto_____1
										if Truth_(or__175__auto_____1) {
											return or__175__auto_____1
										} else {
											return p2.(CljsCoreIFn).X_invoke_Arity1(p1__421_SHARP_)
										}
									}
								})
							}(&AFn{}, or__175__auto__), args)
						}
					}
				})
			}(&AFn{})
		}, func(p1 interface{}, p2 interface{}, p3 interface{}) interface{} {
			return func(sp3 *AFn) *AFn {
				return Fn(sp3, func() interface{} {
					return nil
				}, func(x interface{}) interface{} {
					{
						var or__175__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							{
								var or__175__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x)
								_ = or__175__auto_____1
								if Truth_(or__175__auto_____1) {
									return or__175__auto_____1
								} else {
									return p3.(CljsCoreIFn).X_invoke_Arity1(x)
								}
							}
						}
					}
				}, func(x interface{}, y interface{}) interface{} {
					{
						var or__175__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							{
								var or__175__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x)
								_ = or__175__auto_____1
								if Truth_(or__175__auto_____1) {
									return or__175__auto_____1
								} else {
									{
										var or__175__auto_____2 = p3.(CljsCoreIFn).X_invoke_Arity1(x)
										_ = or__175__auto_____2
										if Truth_(or__175__auto_____2) {
											return or__175__auto_____2
										} else {
											{
												var or__175__auto_____3 = p1.(CljsCoreIFn).X_invoke_Arity1(y)
												_ = or__175__auto_____3
												if Truth_(or__175__auto_____3) {
													return or__175__auto_____3
												} else {
													{
														var or__175__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y)
														_ = or__175__auto_____4
														if Truth_(or__175__auto_____4) {
															return or__175__auto_____4
														} else {
															return p3.(CljsCoreIFn).X_invoke_Arity1(y)
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
						var or__175__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							{
								var or__175__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x)
								_ = or__175__auto_____1
								if Truth_(or__175__auto_____1) {
									return or__175__auto_____1
								} else {
									{
										var or__175__auto_____2 = p3.(CljsCoreIFn).X_invoke_Arity1(x)
										_ = or__175__auto_____2
										if Truth_(or__175__auto_____2) {
											return or__175__auto_____2
										} else {
											{
												var or__175__auto_____3 = p1.(CljsCoreIFn).X_invoke_Arity1(y)
												_ = or__175__auto_____3
												if Truth_(or__175__auto_____3) {
													return or__175__auto_____3
												} else {
													{
														var or__175__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y)
														_ = or__175__auto_____4
														if Truth_(or__175__auto_____4) {
															return or__175__auto_____4
														} else {
															{
																var or__175__auto_____5 = p3.(CljsCoreIFn).X_invoke_Arity1(y)
																_ = or__175__auto_____5
																if Truth_(or__175__auto_____5) {
																	return or__175__auto_____5
																} else {
																	{
																		var or__175__auto_____6 = p1.(CljsCoreIFn).X_invoke_Arity1(z)
																		_ = or__175__auto_____6
																		if Truth_(or__175__auto_____6) {
																			return or__175__auto_____6
																		} else {
																			{
																				var or__175__auto_____7 = p2.(CljsCoreIFn).X_invoke_Arity1(z)
																				_ = or__175__auto_____7
																				if Truth_(or__175__auto_____7) {
																					return or__175__auto_____7
																				} else {
																					return p3.(CljsCoreIFn).X_invoke_Arity1(z)
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
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					{
						var or__175__auto__ = sp3.X_invoke_Arity3(x, y, z)
						_ = or__175__auto__
						if Truth_(or__175__auto__) {
							return or__175__auto__
						} else {
							return Some.X_invoke_Arity2(func(G__428 *AFn, or__175__auto__ interface{}) *AFn {
								return Fn(G__428, func(p1__422_SHARP_ interface{}) interface{} {
									{
										var or__175__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(p1__422_SHARP_)
										_ = or__175__auto_____1
										if Truth_(or__175__auto_____1) {
											return or__175__auto_____1
										} else {
											{
												var or__175__auto_____2 = p2.(CljsCoreIFn).X_invoke_Arity1(p1__422_SHARP_)
												_ = or__175__auto_____2
												if Truth_(or__175__auto_____2) {
													return or__175__auto_____2
												} else {
													return p3.(CljsCoreIFn).X_invoke_Arity1(p1__422_SHARP_)
												}
											}
										}
									}
								})
							}(&AFn{}, or__175__auto__), args)
						}
					}
				})
			}(&AFn{})
		}, func(p1_p2_p3_ps__ ...interface{}) interface{} {
			var p1 = p1_p2_p3_ps__[0]
			var p2 = p1_p2_p3_ps__[1]
			var p3 = p1_p2_p3_ps__[2]
			var ps = Array_seq.X_invoke_Arity1(p1_p2_p3_ps__[3:])
			_, _, _, _ = p1, p2, p3, ps
			{
				var ps___1 = List_STAR_.X_invoke_Arity4(p1, p2, p3, ps).(*CljsCoreCons)
				_ = ps___1
				return func(spn *AFn, ps___1 *CljsCoreCons) *AFn {
					return Fn(spn, func() interface{} {
						return nil
					}, func(x interface{}) interface{} {
						return Some.X_invoke_Arity2(func(G__429 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__429, func(p1__423_SHARP_ interface{}) interface{} {
								return p1__423_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x)
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x interface{}, y interface{}) interface{} {
						return Some.X_invoke_Arity2(func(G__430 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__430, func(p1__424_SHARP_ interface{}) interface{} {
								{
									var or__175__auto__ = p1__424_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x)
									_ = or__175__auto__
									if Truth_(or__175__auto__) {
										return or__175__auto__
									} else {
										return p1__424_SHARP_.(CljsCoreIFn).X_invoke_Arity1(y)
									}
								}
							})
						}(&AFn{}, ps___1), ps___1)
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return Some.X_invoke_Arity2(func(G__431 *AFn, ps___1 *CljsCoreCons) *AFn {
							return Fn(G__431, func(p1__425_SHARP_ interface{}) interface{} {
								{
									var or__175__auto__ = p1__425_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x)
									_ = or__175__auto__
									if Truth_(or__175__auto__) {
										return or__175__auto__
									} else {
										{
											var or__175__auto_____1 = p1__425_SHARP_.(CljsCoreIFn).X_invoke_Arity1(y)
											_ = or__175__auto_____1
											if Truth_(or__175__auto_____1) {
												return or__175__auto_____1
											} else {
												return p1__425_SHARP_.(CljsCoreIFn).X_invoke_Arity1(z)
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
						var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
						_, _, _, _ = x, y, z, args
						{
							var or__175__auto__ = spn.X_invoke_Arity3(x, y, z)
							_ = or__175__auto__
							if Truth_(or__175__auto__) {
								return or__175__auto__
							} else {
								return Some.X_invoke_Arity2(func(G__432 *AFn, or__175__auto__ interface{}, ps___1 *CljsCoreCons) *AFn {
									return Fn(G__432, func(p1__426_SHARP_ interface{}) interface{} {
										return Some.X_invoke_Arity2(p1__426_SHARP_, args)
									})
								}(&AFn{}, or__175__auto__, ps___1), ps___1)
							}
						}
					})
				}(&AFn{}, ps___1)
			}
		})
	}(&AFn{})

	Map_ = func(map_ *AFn) *AFn {
		return Fn(map_, func(f interface{}) interface{} {
			return func(G__434 *AFn) *AFn {
				return Fn(G__434, func(f1 interface{}) interface{} {
					return func(G__435 *AFn) *AFn {
						return Fn(G__435, func() interface{} {
							return f1.(CljsCoreIFn).X_invoke_Arity0()
						}, func(result interface{}) interface{} {
							return f1.(CljsCoreIFn).X_invoke_Arity1(result)
						}, func(result interface{}, input interface{}) interface{} {
							return f1.(CljsCoreIFn).X_invoke_Arity2(result, f.(CljsCoreIFn).X_invoke_Arity1(input))
						}, func(result_input_inputs__ ...interface{}) interface{} {
							var result = result_input_inputs__[0]
							var input = result_input_inputs__[1]
							var inputs = Array_seq.X_invoke_Arity1(result_input_inputs__[2:])
							_, _, _ = result, input, inputs
							return f1.(CljsCoreIFn).X_invoke_Arity2(result, Apply.X_invoke_Arity3(f, input, inputs))
						})
					}(&AFn{})
				})
			}(&AFn{})
		}, func(f interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__436 *AFn) *AFn {
				return Fn(G__436, func() interface{} {
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
											var n__1042__auto___437 = size
											_ = n__1042__auto___437
											{
												var i_438 = float64(0)
												_ = i_438
												for {
													if i_438 < n__1042__auto___437 {
														Chunk_append.X_invoke_Arity2(b, f.(CljsCoreIFn).X_invoke_Arity1(c.(CljsCoreIIndexed).X_nth_Arity2(i_438)))
														i_438 = (i_438 + float64(1))
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
									return Cons.X_invoke_Arity2(f.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(s)), map_.X_invoke_Arity2(f, Rest.Arity1IQ(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
								}
							}
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		}, func(f interface{}, c1 interface{}, c2 interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__439 *AFn) *AFn {
				return Fn(G__439, func() interface{} {
					{
						var s1 = Seq.Arity1IQ(c1)
						var s2 = Seq.Arity1IQ(c2)
						_, _ = s1, s2
						if Truth_(func() interface{} {
							var and__163__auto__ = s1
							_ = and__163__auto__
							if Truth_(and__163__auto__) {
								return s2
							} else {
								return and__163__auto__
							}
						}()) {
							return Cons.X_invoke_Arity2(f.(CljsCoreIFn).X_invoke_Arity2(First.X_invoke_Arity1(s1), First.X_invoke_Arity1(s2)), map_.X_invoke_Arity3(f, Rest.Arity1IQ(s1), Rest.Arity1IQ(s2)).(*CljsCoreLazySeq)).(*CljsCoreCons)
						} else {
							return nil
						}
					}
				})
			}(&AFn{}), nil, nil})
		}, func(f interface{}, c1 interface{}, c2 interface{}, c3 interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__440 *AFn) *AFn {
				return Fn(G__440, func() interface{} {
					{
						var s1 = Seq.Arity1IQ(c1)
						var s2 = Seq.Arity1IQ(c2)
						var s3 = Seq.Arity1IQ(c3)
						_, _, _ = s1, s2, s3
						if Truth_(func() interface{} {
							var and__163__auto__ = s1
							_ = and__163__auto__
							if Truth_(and__163__auto__) {
								{
									var and__163__auto_____1 = s2
									_ = and__163__auto_____1
									if Truth_(and__163__auto_____1) {
										return s3
									} else {
										return and__163__auto_____1
									}
								}
							} else {
								return and__163__auto__
							}
						}()) {
							return Cons.X_invoke_Arity2(f.(CljsCoreIFn).X_invoke_Arity3(First.X_invoke_Arity1(s1), First.X_invoke_Arity1(s2), First.X_invoke_Arity1(s3)), map_.X_invoke_Arity4(f, Rest.Arity1IQ(s1), Rest.Arity1IQ(s2), Rest.Arity1IQ(s3)).(*CljsCoreLazySeq)).(*CljsCoreCons)
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
			var colls = Array_seq.X_invoke_Arity1(f_c1_c2_c3_colls__[4:])
			_, _, _, _, _ = f, c1, c2, c3, colls
			{
				var step = func(step *AFn) *AFn {
					return Fn(step, func(cs interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__441 *AFn) *AFn {
							return Fn(G__441, func() interface{} {
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
				return map_.X_invoke_Arity2(func(G__442 *AFn, step CljsCoreIFn) *AFn {
					return Fn(G__442, func(p1__433_SHARP_ interface{}) interface{} {
						return Apply.X_invoke_Arity2(f, p1__433_SHARP_)
					})
				}(&AFn{}, step), step.X_invoke_Arity1(Conj.X_invoke_ArityVariadic(colls, c3, c2, c1)).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
			}
		})
	}(&AFn{})

	Take = func(take *AFn) *AFn {
		return Fn(take, func(n interface{}) interface{} {
			return func(G__443 *AFn) *AFn {
				return Fn(G__443, func(f1 interface{}) interface{} {
					{
						var na = Atom.X_invoke_Arity1(n).(*CljsCoreAtom)
						_ = na
						return func(G__444 *AFn, na *CljsCoreAtom) *AFn {
							return Fn(G__444, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0()
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result)
							}, func(result interface{}, input interface{}) interface{} {
								{
									var n___1 = Deref.X_invoke_Arity1(na)
									var nn = Swap_BANG_.X_invoke_Arity2(na, Dec)
									var result___1 = func() interface{} {
										if n___1.(float64) > float64(0) {
											return f1.(CljsCoreIFn).X_invoke_Arity2(result, input)
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
			return (&CljsCoreLazySeq{nil, func(G__445 *AFn) *AFn {
				return Fn(G__445, func() interface{} {
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
		return Fn(drop, func(n interface{}) interface{} {
			return func(G__446 *AFn) *AFn {
				return Fn(G__446, func(f1 interface{}) interface{} {
					{
						var na = Atom.X_invoke_Arity1(n).(*CljsCoreAtom)
						_ = na
						return func(G__447 *AFn, na *CljsCoreAtom) *AFn {
							return Fn(G__447, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0()
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result)
							}, func(result interface{}, input interface{}) interface{} {
								{
									var n___1 = Deref.X_invoke_Arity1(na)
									_ = n___1
									Swap_BANG_.X_invoke_Arity2(na, Dec)
									if n___1.(float64) > float64(0) {
										return result
									} else {
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, input)
									}
								}
							})
						}(&AFn{}, na)
					}
				})
			}(&AFn{})
		}, func(n interface{}, coll interface{}) interface{} {
			{
				var step = func(G__448 *AFn) *AFn {
					return Fn(G__448, func(n___1 interface{}, coll___1 interface{}) interface{} {
						for {
							{
								var s = Seq.Arity1IQ(coll___1)
								_ = s
								if Truth_(func() interface{} {
									var and__163__auto__ = (n___1.(float64) > float64(0))
									_ = and__163__auto__
									if Truth_(and__163__auto__) {
										return s
									} else {
										return and__163__auto__
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
				return (&CljsCoreLazySeq{nil, func(G__449 *AFn, step CljsCoreIFn) *AFn {
					return Fn(G__449, func() interface{} {
						return Seq_(step.X_invoke_Arity2(n, coll))
					})
				}(&AFn{}, step), nil, nil})
			}
		})
	}(&AFn{})

	Drop_last = func(drop_last *AFn) *AFn {
		return Fn(drop_last, func(s interface{}) interface{} {
			return drop_last.X_invoke_Arity2(float64(1), s).(*CljsCoreLazySeq)
		}, func(n interface{}, s interface{}) interface{} {
			return Map_.X_invoke_Arity3(func(G__450 *AFn) *AFn {
				return Fn(G__450, func(x interface{}, ___ interface{}) interface{} {
					return x
				})
			}(&AFn{}), s, Drop.X_invoke_Arity2(n, s).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Take_last = func(take_last *AFn) *AFn {
		return Fn(take_last, func(n interface{}, coll interface{}) interface{} {
			{
				var s = Seq.Arity1IQ(coll)
				var lead = Seq.Arity1IQ(Drop.X_invoke_Arity2(n, coll).(*CljsCoreLazySeq))
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
		return Fn(drop_while, func(pred interface{}) interface{} {
			return func(G__451 *AFn) *AFn {
				return Fn(G__451, func(f1 interface{}) interface{} {
					{
						var da = Atom.X_invoke_Arity1(true).(*CljsCoreAtom)
						_ = da
						return func(G__452 *AFn, da *CljsCoreAtom) *AFn {
							return Fn(G__452, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0()
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result)
							}, func(result interface{}, input interface{}) interface{} {
								{
									var drop_QMARK_ = Deref.X_invoke_Arity1(da)
									_ = drop_QMARK_
									if Truth_(func() interface{} {
										var and__163__auto__ = drop_QMARK_
										_ = and__163__auto__
										if Truth_(and__163__auto__) {
											return pred.(CljsCoreIFn).X_invoke_Arity1(input)
										} else {
											return and__163__auto__
										}
									}()) {
										return result
									} else {
										Reset_BANG_.X_invoke_Arity2(da, nil)
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, input)
									}
								}
							})
						}(&AFn{}, da)
					}
				})
			}(&AFn{})
		}, func(pred interface{}, coll interface{}) interface{} {
			{
				var step = func(G__453 *AFn) *AFn {
					return Fn(G__453, func(pred___1 interface{}, coll___1 interface{}) interface{} {
						for {
							{
								var s = Seq.Arity1IQ(coll___1)
								_ = s
								if Truth_(func() interface{} {
									var and__163__auto__ = s
									_ = and__163__auto__
									if Truth_(and__163__auto__) {
										return pred___1.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(s))
									} else {
										return and__163__auto__
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
				return (&CljsCoreLazySeq{nil, func(G__454 *AFn, step CljsCoreIFn) *AFn {
					return Fn(G__454, func() interface{} {
						return Seq_(step.X_invoke_Arity2(pred, coll))
					})
				}(&AFn{}, step), nil, nil})
			}
		})
	}(&AFn{})

	Cycle = func(cycle *AFn) *AFn {
		return Fn(cycle, func(coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__455 *AFn) *AFn {
				return Fn(G__455, func() interface{} {
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
		return Fn(split_at, func(n interface{}, coll interface{}) interface{} {
			return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Take.X_invoke_Arity2(n, coll).(*CljsCoreLazySeq), Drop.X_invoke_Arity2(n, coll).(*CljsCoreLazySeq)}, nil})
		})
	}(&AFn{})

	Repeat = func(repeat *AFn) *AFn {
		return Fn(repeat, func(x interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__456 *AFn) *AFn {
				return Fn(G__456, func() interface{} {
					return Cons.X_invoke_Arity2(x, repeat.X_invoke_Arity1(x).(*CljsCoreLazySeq)).(*CljsCoreCons)
				})
			}(&AFn{}), nil, nil})
		}, func(n interface{}, x interface{}) interface{} {
			return Take.X_invoke_Arity2(n, repeat.X_invoke_Arity1(x).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Replicate = func(replicate *AFn) *AFn {
		return Fn(replicate, func(n interface{}, x interface{}) interface{} {
			return Take.X_invoke_Arity2(n, Repeat.X_invoke_Arity1(x).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Repeatedly = func(repeatedly *AFn) *AFn {
		return Fn(repeatedly, func(f interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__457 *AFn) *AFn {
				return Fn(G__457, func() interface{} {
					return Cons.X_invoke_Arity2(f.(CljsCoreIFn).X_invoke_Arity0(), repeatedly.X_invoke_Arity1(f).(*CljsCoreLazySeq)).(*CljsCoreCons)
				})
			}(&AFn{}), nil, nil})
		}, func(n interface{}, f interface{}) interface{} {
			return Take.X_invoke_Arity2(n, repeatedly.X_invoke_Arity1(f).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Iterate = func(iterate *AFn) *AFn {
		return Fn(iterate, func(f interface{}, x interface{}) interface{} {
			return Cons.X_invoke_Arity2(x, (&CljsCoreLazySeq{nil, func(G__458 *AFn) *AFn {
				return Fn(G__458, func() interface{} {
					return iterate.X_invoke_Arity2(f, f.(CljsCoreIFn).X_invoke_Arity1(x)).(*CljsCoreCons)
				})
			}(&AFn{}), nil, nil})).(*CljsCoreCons)
		})
	}(&AFn{})

	Interleave = func(interleave *AFn) *AFn {
		return Fn(interleave, func(c1 interface{}, c2 interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__459 *AFn) *AFn {
				return Fn(G__459, func() interface{} {
					{
						var s1 = Seq.Arity1IQ(c1)
						var s2 = Seq.Arity1IQ(c2)
						_, _ = s1, s2
						if Truth_(func() interface{} {
							var and__163__auto__ = s1
							_ = and__163__auto__
							if Truth_(and__163__auto__) {
								return s2
							} else {
								return and__163__auto__
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
			var colls = Array_seq.X_invoke_Arity1(c1_c2_colls__[2:])
			_, _, _ = c1, c2, colls
			return (&CljsCoreLazySeq{nil, func(G__460 *AFn) *AFn {
				return Fn(G__460, func() interface{} {
					{
						var ss = Map_.X_invoke_Arity2(Seq, Conj.X_invoke_ArityVariadic(colls, c2, c1)).(*CljsCoreLazySeq)
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
		return Fn(interpose, func(sep interface{}, coll interface{}) interface{} {
			return Drop.X_invoke_Arity2(float64(1), Interleave.X_invoke_Arity2(Repeat.X_invoke_Arity1(sep).(*CljsCoreLazySeq), coll).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Flatten1 = func(flatten1 *AFn) *AFn {
		return Fn(flatten1, func(colls interface{}) interface{} {
			{
				var cat = func(cat *AFn) *AFn {
					return Fn(cat, func(coll interface{}, colls___1 interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__461 *AFn) *AFn {
							return Fn(G__461, func() interface{} {
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
		return Fn(mapcat, func(f interface{}) interface{} {
			return Comp.X_invoke_Arity2(Map_.X_invoke_Arity1(f).(CljsCoreIFn), Cat).(CljsCoreIFn)
		}, func(f_colls__ ...interface{}) interface{} {
			var f = f_colls__[0]
			var colls = Array_seq.X_invoke_Arity1(f_colls__[1:])
			_, _ = f, colls
			return Apply.X_invoke_Arity2(Concat, Apply.X_invoke_Arity3(Map_, f, colls))
		})
	}(&AFn{})

	Filter = func(filter *AFn) *AFn {
		return Fn(filter, func(pred interface{}) interface{} {
			return func(G__462 *AFn) *AFn {
				return Fn(G__462, func(f1 interface{}) interface{} {
					return func(G__463 *AFn) *AFn {
						return Fn(G__463, func() interface{} {
							return f1.(CljsCoreIFn).X_invoke_Arity0()
						}, func(result interface{}) interface{} {
							return f1.(CljsCoreIFn).X_invoke_Arity1(result)
						}, func(result interface{}, input interface{}) interface{} {
							if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(input)) {
								return f1.(CljsCoreIFn).X_invoke_Arity2(result, input)
							} else {
								return result
							}
						})
					}(&AFn{})
				})
			}(&AFn{})
		}, func(pred interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__464 *AFn) *AFn {
				return Fn(G__464, func() interface{} {
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
											var n__1042__auto___465 = size
											_ = n__1042__auto___465
											{
												var i_466 = float64(0)
												_ = i_466
												for {
													if i_466 < n__1042__auto___465 {
														if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(c.(CljsCoreIIndexed).X_nth_Arity2(i_466))) {
															Chunk_append.X_invoke_Arity2(b, c.(CljsCoreIIndexed).X_nth_Arity2(i_466))
														} else {
														}
														i_466 = (i_466 + float64(1))
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
										if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(f)) {
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

	Remove = func(remove *AFn) *AFn {
		return Fn(remove, func(pred interface{}) interface{} {
			return Filter.X_invoke_Arity1(Complement.Arity1IB(pred)).(CljsCoreIFn)
		}, func(pred interface{}, coll interface{}) interface{} {
			return Filter.X_invoke_Arity2(Complement.Arity1IB(pred), coll).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Tree_seq = func(tree_seq *AFn) *AFn {
		return Fn(tree_seq, func(branch_QMARK_ interface{}, children interface{}, root interface{}) interface{} {
			{
				var walk = func(walk *AFn) *AFn {
					return Fn(walk, func(node interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__467 *AFn) *AFn {
							return Fn(G__467, func() interface{} {
								return Cons.X_invoke_Arity2(node, func() interface{} {
									if Truth_(branch_QMARK_.(CljsCoreIFn).X_invoke_Arity1(node)) {
										return Mapcat.X_invoke_ArityVariadic(walk, children.(CljsCoreIFn).X_invoke_Arity1(node))
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
		return Fn(flatten, func(x interface{}) interface{} {
			return Filter.X_invoke_Arity2(func(G__469 *AFn) *AFn {
				return Fn(G__469, func(p1__468_SHARP_ interface{}) interface{} {
					return !(Sequential_QMARK_.Arity1IB(p1__468_SHARP_))
				})
			}(&AFn{}), Rest.Arity1IQ(Tree_seq.X_invoke_Arity3(Sequential_QMARK_, Seq, x).(*CljsCoreLazySeq))).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	Into = func(into *AFn) *AFn {
		return Fn(into, func(to interface{}, from interface{}) interface{} {
			if !(Nil_(to)) {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IEditableCollection", Str: "cljs.core/IEditableCollection", X_hash: float64(297050504), X_meta: nil}), to)) {
					return With_meta.X_invoke_Arity2(Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(X_conj_BANG_, Transient.X_invoke_Arity1(to), from)), Meta.X_invoke_Arity1(to))
				} else {
					return Reduce.X_invoke_Arity3(X_conj, to, from)
				}
			} else {
				return Reduce.X_invoke_Arity3(Conj, CljsCoreIEmptyList(CljsCoreList_EMPTY), from)
			}
		}, func(to interface{}, xform interface{}, from interface{}) interface{} {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IEditableCollection", Str: "cljs.core/IEditableCollection", X_hash: float64(297050504), X_meta: nil}), to)) {
				return With_meta.X_invoke_Arity2(Persistent_BANG_.X_invoke_Arity1(Transduce.X_invoke_Arity4(xform, Conj_BANG_, Transient.X_invoke_Arity1(to), from)), Meta.X_invoke_Arity1(to))
			} else {
				return Transduce.X_invoke_Arity4(xform, Conj, to, from)
			}
		})
	}(&AFn{})

	Mapv = func(mapv *AFn) *AFn {
		return Fn(mapv, func(f interface{}, coll interface{}) interface{} {
			return Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(func(G__470 *AFn) *AFn {
				return Fn(G__470, func(v interface{}, o interface{}) interface{} {
					return Conj_BANG_.X_invoke_Arity2(v, f.(CljsCoreIFn).X_invoke_Arity1(o))
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
			var colls = Array_seq.X_invoke_Arity1(f_c1_c2_c3_colls__[4:])
			_, _, _, _, _ = f, c1, c2, c3, colls
			return Into.X_invoke_Arity2(CljsCorePersistentVector_EMPTY, Apply.X_invoke_ArityVariadic(Map_, f, c1, c2, c3, colls))
		})
	}(&AFn{})

	Filterv = func(filterv *AFn) *AFn {
		return Fn(filterv, func(pred interface{}, coll interface{}) interface{} {
			return Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(func(G__471 *AFn) *AFn {
				return Fn(G__471, func(v interface{}, o interface{}) interface{} {
					if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(o)) {
						return Conj_BANG_.X_invoke_Arity2(v, o)
					} else {
						return v
					}
				})
			}(&AFn{}), Transient.X_invoke_Arity1(CljsCorePersistentVector_EMPTY), coll))
		})
	}(&AFn{})

	Partition = func(partition *AFn) *AFn {
		return Fn(partition, func(n interface{}, coll interface{}) interface{} {
			return partition.X_invoke_Arity3(n, n, coll).(*CljsCoreLazySeq)
		}, func(n interface{}, step interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__472 *AFn) *AFn {
				return Fn(G__472, func() interface{} {
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
			return (&CljsCoreLazySeq{nil, func(G__473 *AFn) *AFn {
				return Fn(G__473, func() interface{} {
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
		return Fn(get_in, func(m interface{}, ks interface{}) interface{} {
			return get_in.X_invoke_Arity3(m, ks, nil)
		}, func(m interface{}, ks interface{}, not_found interface{}) interface{} {
			{
				var sentinel = Lookup_sentinel
				var m___1 = m
				var ks___1 = Seq.Arity1IQ(ks)
				_, _, _ = sentinel, m___1, ks___1
				for {
					if Truth_(ks___1) {
						if Not.Arity1IB(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ILookup", Str: "cljs.core/ILookup", X_hash: float64(-150575073), X_meta: nil}), m___1)) {
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
		return Fn(assoc_in, func(m interface{}, p__474 interface{}, v interface{}) interface{} {
			{
				var vec__476 = p__474
				var k = Nth.X_invoke_Arity3(vec__476, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__476, float64(1)))
				_, _, _ = vec__476, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, assoc_in.X_invoke_Arity3(Get.X_invoke_Arity2(m, k), ks, v))
				} else {
					return Assoc.X_invoke_Arity3(m, k, v)
				}
			}
		})
	}(&AFn{})

	Update_in = func(update_in *AFn) *AFn {
		return Fn(update_in, func(m interface{}, p__477 interface{}, f interface{}) interface{} {
			{
				var vec__487 = p__477
				var k = Nth.X_invoke_Arity3(vec__487, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__487, float64(1)))
				_, _, _ = vec__487, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity3(Get.X_invoke_Arity2(m, k), ks, f))
				} else {
					return Assoc.X_invoke_Arity3(m, k, f.(CljsCoreIFn).X_invoke_Arity1(Get.X_invoke_Arity2(m, k)))
				}
			}
		}, func(m interface{}, p__478 interface{}, f interface{}, a interface{}) interface{} {
			{
				var vec__488 = p__478
				var k = Nth.X_invoke_Arity3(vec__488, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__488, float64(1)))
				_, _, _ = vec__488, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity4(Get.X_invoke_Arity2(m, k), ks, f, a))
				} else {
					return Assoc.X_invoke_Arity3(m, k, f.(CljsCoreIFn).X_invoke_Arity2(Get.X_invoke_Arity2(m, k), a))
				}
			}
		}, func(m interface{}, p__479 interface{}, f interface{}, a interface{}, b interface{}) interface{} {
			{
				var vec__489 = p__479
				var k = Nth.X_invoke_Arity3(vec__489, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__489, float64(1)))
				_, _, _ = vec__489, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity5(Get.X_invoke_Arity2(m, k), ks, f, a, b))
				} else {
					return Assoc.X_invoke_Arity3(m, k, f.(CljsCoreIFn).X_invoke_Arity3(Get.X_invoke_Arity2(m, k), a, b))
				}
			}
		}, func(m interface{}, p__480 interface{}, f interface{}, a interface{}, b interface{}, c interface{}) interface{} {
			{
				var vec__490 = p__480
				var k = Nth.X_invoke_Arity3(vec__490, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__490, float64(1)))
				_, _, _ = vec__490, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity6(Get.X_invoke_Arity2(m, k), ks, f, a, b, c))
				} else {
					return Assoc.X_invoke_Arity3(m, k, f.(CljsCoreIFn).X_invoke_Arity4(Get.X_invoke_Arity2(m, k), a, b, c))
				}
			}
		}, func(m_p__481_f_a_b_c_args__ ...interface{}) interface{} {
			var m = m_p__481_f_a_b_c_args__[0]
			var p__481 = m_p__481_f_a_b_c_args__[1]
			var f = m_p__481_f_a_b_c_args__[2]
			var a = m_p__481_f_a_b_c_args__[3]
			var b = m_p__481_f_a_b_c_args__[4]
			var c = m_p__481_f_a_b_c_args__[5]
			var args = Array_seq.X_invoke_Arity1(m_p__481_f_a_b_c_args__[6:])
			_, _, _, _, _, _, _ = m, p__481, f, a, b, c, args
			{
				var vec__491 = p__481
				var k = Nth.X_invoke_Arity3(vec__491, float64(0), nil)
				var ks = Seq_(Nthnext.X_invoke_Arity2(vec__491, float64(1)))
				_, _, _ = vec__491, k, ks
				if Truth_(ks) {
					return Assoc.X_invoke_Arity3(m, k, Apply.X_invoke_ArityVariadic(update_in, Get.X_invoke_Arity2(m, k), ks, f, a, b, c, args))
				} else {
					return Assoc.X_invoke_Arity3(m, k, Apply.X_invoke_ArityVariadic(f, Get.X_invoke_Arity2(m, k), a, b, c, args))
				}
			}
		})
	}(&AFn{})

	X__GT_VectorNode = func(__GT_VectorNode *AFn) *AFn {
		return Fn(__GT_VectorNode, func(edit interface{}, arr interface{}) interface{} {
			return (&CljsCoreVectorNode{edit, arr})
		})
	}(&AFn{})

	Pv_fresh_node = func(pv_fresh_node *AFn) *AFn {
		return Fn(pv_fresh_node, func(edit interface{}) interface{} {
			return (&CljsCoreVectorNode{edit, make([]interface{}, int(float64(32)))})
		})
	}(&AFn{})

	Pv_aget = func(pv_aget *AFn) *AFn {
		return Fn(pv_aget, func(node interface{}, idx interface{}) interface{} {
			return (Native_get_instance_field.X_invoke_Arity2(node, "Arr").([]interface{})[int(idx.(float64))])
		})
	}(&AFn{})

	Pv_aset = func(pv_aset *AFn) *AFn {
		return Fn(pv_aset, func(node interface{}, idx interface{}, val interface{}) interface{} {
			return func() interface{} {
				Native_get_instance_field.X_invoke_Arity2(node, "Arr").([]interface{})[int(idx.(float64))] = val
				return Native_get_instance_field.X_invoke_Arity2(node, "Arr").([]interface{})[int(idx.(float64))]
			}()
		})
	}(&AFn{})

	Pv_clone_node = func(pv_clone_node *AFn) *AFn {
		return Fn(pv_clone_node, func(node interface{}) interface{} {
			return (&CljsCoreVectorNode{Native_get_instance_field.X_invoke_Arity2(node, "Edit"), Aclone.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(node, "Arr")).([]interface{})})
		})
	}(&AFn{})

	Tail_off = func(tail_off *AFn) *AFn {
		return Fn(tail_off, func(pv interface{}) interface{} {
			{
				var cnt = Native_get_instance_field.X_invoke_Arity2(pv, "Cnt")
				_ = cnt
				if cnt.(float64) < float64(32) {
					return float64(0)
				} else {
					return float64(int(float64(uint((cnt.(float64)-float64(1)))>>uint(float64(5)))) << uint(float64(5)))
				}
			}
		})
	}(&AFn{})

	New_path = func(new_path *AFn) *AFn {
		return Fn(new_path, func(edit interface{}, level interface{}, node interface{}) interface{} {
			{
				var ll = level
				var ret = node
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
		return Fn(push_tail, func(pv interface{}, level interface{}, parent interface{}, tailnode interface{}) interface{} {
			{
				var ret = Pv_clone_node.X_invoke_Arity1(parent).(*CljsCoreVectorNode)
				var subidx = float64(int(float64(uint((Native_get_instance_field.X_invoke_Arity2(pv, "Cnt").(float64)-float64(1)))>>uint(level.(float64)))) & int(float64(31)))
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
		return Fn(vector_index_out_of_bounds, func(i interface{}, cnt interface{}) interface{} {
			panic((&js.Error{("No item " + Str.X_invoke_Arity1(i).(string) + " in vector of length " + Str.X_invoke_Arity1(cnt).(string))}))
		})
	}(&AFn{})

	First_array_for_longvec = func(first_array_for_longvec *AFn) *AFn {
		return Fn(first_array_for_longvec, func(pv interface{}) interface{} {
			{
				var node = Native_get_instance_field.X_invoke_Arity2(pv, "Root")
				var level = Native_get_instance_field.X_invoke_Arity2(pv, "Shift")
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
		return Fn(unchecked_array_for, func(pv interface{}, i interface{}) interface{} {
			if i.(float64) >= Tail_off.X_invoke_Arity1(pv).(float64) {
				return Native_get_instance_field.X_invoke_Arity2(pv, "Tail")
			} else {
				{
					var node = Native_get_instance_field.X_invoke_Arity2(pv, "Root")
					var level = Native_get_instance_field.X_invoke_Arity2(pv, "Shift")
					_, _ = node, level
					for {
						if level.(float64) > float64(0) {
							node, level = Pv_aget.X_invoke_Arity2(node, float64(int(float64(uint(i.(float64))>>uint(level.(float64))))&int(float64(31)))), (level.(float64) - float64(5))
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
		return Fn(array_for, func(pv interface{}, i interface{}) interface{} {
			if (float64(0) <= i.(float64)) && (i.(float64) < Native_get_instance_field.X_invoke_Arity2(pv, "Cnt").(float64)) {
				return Unchecked_array_for.X_invoke_Arity2(pv, i)
			} else {
				return Vector_index_out_of_bounds.X_invoke_Arity2(i, Native_get_instance_field.X_invoke_Arity2(pv, "Cnt"))
			}
		})
	}(&AFn{})

	Do_assoc = func(do_assoc *AFn) *AFn {
		return Fn(do_assoc, func(pv interface{}, level interface{}, node interface{}, i interface{}, val interface{}) interface{} {
			{
				var ret = Pv_clone_node.X_invoke_Arity1(node).(*CljsCoreVectorNode)
				_ = ret
				if level.(float64) == float64(0) {
					Pv_aset.X_invoke_Arity3(ret, float64(int(i.(float64))&int(float64(31))), val)
					return ret
				} else {
					{
						var subidx = float64(int(float64(uint(i.(float64))>>uint(level.(float64)))) & int(float64(31)))
						_ = subidx
						Pv_aset.X_invoke_Arity3(ret, subidx, do_assoc.X_invoke_Arity5(pv, (level.(float64)-float64(5)), Pv_aget.X_invoke_Arity2(node, subidx), i, val).(*CljsCoreVectorNode))
						return ret
					}
				}
			}
		})
	}(&AFn{})

	Pop_tail = func(pop_tail *AFn) *AFn {
		return Fn(pop_tail, func(pv interface{}, level interface{}, node interface{}) interface{} {
			{
				var subidx = float64(int(float64(uint((Native_get_instance_field.X_invoke_Arity2(pv, "Cnt").(float64)-float64(2)))>>uint(level.(float64)))) & int(float64(31)))
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
		return Fn(__GT_PersistentVector, func(meta interface{}, cnt interface{}, shift interface{}, root interface{}, tail interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentVector{meta, cnt, shift, root, tail, __hash})
		})
	}(&AFn{})

	Vec = func(vec *AFn) *AFn {
		return Fn(vec, func(coll interface{}) interface{} {
			return Reduce.X_invoke_Arity3(X_conj_BANG_, CljsCorePersistentVector_EMPTY.X_as_transient_Arity1(), coll).(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
		})
	}(&AFn{})

	Vector = func(vector *AFn) *AFn {
		return Fn(vector, func(args__ ...interface{}) interface{} {
			var args = Array_seq.X_invoke_Arity1(args__[0:])
			_ = args
			if (func() bool { _, instanceof := args.(*CljsCoreIndexedSeq); return instanceof }()) && (Native_get_instance_field.X_invoke_Arity2(args, "I").(float64) == float64(0)) {
				return CljsCorePersistentVector_FromArray.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(args, "Arr"), true)
			} else {
				return Vec.X_invoke_Arity1(args)
			}
		})
	}(&AFn{})

	X__GT_ChunkedSeq = func(__GT_ChunkedSeq *AFn) *AFn {
		return Fn(__GT_ChunkedSeq, func(vec interface{}, node interface{}, i interface{}, off interface{}, meta interface{}, __hash interface{}) interface{} {
			return (&CljsCoreChunkedSeq{vec, node, i, off, meta, __hash})
		})
	}(&AFn{})

	Chunked_seq = func(chunked_seq *AFn) *AFn {
		return Fn(chunked_seq, func(vec interface{}, i interface{}, off interface{}) interface{} {
			return (&CljsCoreChunkedSeq{vec, Array_for.X_invoke_Arity2(vec, i), i, off, nil, nil})
		}, func(vec interface{}, node interface{}, i interface{}, off interface{}) interface{} {
			return (&CljsCoreChunkedSeq{vec, node, i, off, nil, nil})
		}, func(vec interface{}, node interface{}, i interface{}, off interface{}, meta interface{}) interface{} {
			return (&CljsCoreChunkedSeq{vec, node, i, off, meta, nil})
		})
	}(&AFn{})

	X__GT_Subvec = func(__GT_Subvec *AFn) *AFn {
		return Fn(__GT_Subvec, func(meta interface{}, v interface{}, start interface{}, end interface{}, __hash interface{}) interface{} {
			return (&CljsCoreSubvec{meta, v, start, end, __hash})
		})
	}(&AFn{})

	Build_subvec = func(build_subvec *AFn) *AFn {
		return Fn(build_subvec, func(meta interface{}, v interface{}, start interface{}, end interface{}, __hash interface{}) interface{} {
			for {
				if func() bool { _, instanceof := v.(*CljsCoreSubvec); return instanceof }() {
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
		return Fn(subvec, func(v interface{}, start interface{}) interface{} {
			return subvec.X_invoke_Arity3(v, start, Count.X_invoke_Arity1(v).(float64)).(*CljsCoreSubvec)
		}, func(v interface{}, start interface{}, end interface{}) interface{} {
			return Build_subvec.X_invoke_Arity5(nil, v, start, end, nil).(*CljsCoreSubvec)
		})
	}(&AFn{})

	Tv_ensure_editable = func(tv_ensure_editable *AFn) *AFn {
		return Fn(tv_ensure_editable, func(edit interface{}, node interface{}) interface{} {
			if reflect.DeepEqual(edit, Native_get_instance_field.X_invoke_Arity2(node, "Edit")) {
				return node
			} else {
				return (&CljsCoreVectorNode{edit, Aclone.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(node, "Arr")).([]interface{})})
			}
		})
	}(&AFn{})

	Tv_editable_root = func(tv_editable_root *AFn) *AFn {
		return Fn(tv_editable_root, func(node interface{}) interface{} {
			return (&CljsCoreVectorNode{func() interface{} {
				var obj496 = map[string]interface{}{}
				_ = obj496
				return obj496
			}(), Aclone.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(node, "Arr")).([]interface{})})
		})
	}(&AFn{})

	Tv_editable_tail = func(tv_editable_tail *AFn) *AFn {
		return Fn(tv_editable_tail, func(tl interface{}) interface{} {
			{
				var ret = make([]interface{}, int(float64(32)))
				_ = ret
				Array_copy.X_invoke_Arity5(tl, float64(0), ret, float64(0), float64(len(tl.([]interface{}))))
				return ret
			}
		})
	}(&AFn{})

	Tv_push_tail = func(tv_push_tail *AFn) *AFn {
		return Fn(tv_push_tail, func(tv interface{}, level interface{}, parent interface{}, tail_node interface{}) interface{} {
			{
				var ret = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tv, "Root"), "Edit"), parent)
				var subidx = float64(int(float64(uint((Native_get_instance_field.X_invoke_Arity2(tv, "Cnt").(float64)-float64(1)))>>uint(level.(float64)))) & int(float64(31)))
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
		return Fn(tv_pop_tail, func(tv interface{}, level interface{}, node interface{}) interface{} {
			{
				var node___1 = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tv, "Root"), "Edit"), node)
				var subidx = float64(int(float64(uint((Native_get_instance_field.X_invoke_Arity2(tv, "Cnt").(float64)-float64(2)))>>uint(level.(float64)))) & int(float64(31)))
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
		return Fn(unchecked_editable_array_for, func(tv interface{}, i interface{}) interface{} {
			if i.(float64) >= Tail_off.X_invoke_Arity1(tv).(float64) {
				return Native_get_instance_field.X_invoke_Arity2(tv, "Tail")
			} else {
				{
					var root = Native_get_instance_field.X_invoke_Arity2(tv, "Root")
					_ = root
					{
						var node = root
						var level = Native_get_instance_field.X_invoke_Arity2(tv, "Shift")
						_, _ = node, level
						for {
							if level.(float64) > float64(0) {
								node, level = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(root, "Edit"), Pv_aget.X_invoke_Arity2(node, float64(int(float64(uint(i.(float64))>>uint(level.(float64))))&int(float64(31))))), (level.(float64) - float64(5))
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
		return Fn(__GT_TransientVector, func(cnt interface{}, shift interface{}, root interface{}, tail interface{}) interface{} {
			return (&CljsCoreTransientVector{cnt, shift, root, tail})
		})
	}(&AFn{})

	X__GT_PersistentQueueSeq = func(__GT_PersistentQueueSeq *AFn) *AFn {
		return Fn(__GT_PersistentQueueSeq, func(meta interface{}, front interface{}, rear interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentQueueSeq{meta, front, rear, __hash})
		})
	}(&AFn{})

	X__GT_PersistentQueue = func(__GT_PersistentQueue *AFn) *AFn {
		return Fn(__GT_PersistentQueue, func(meta interface{}, count interface{}, front interface{}, rear interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentQueue{meta, count, front, rear, __hash})
		})
	}(&AFn{})

	X__GT_NeverEquiv = func(__GT_NeverEquiv *AFn) *AFn {
		return Fn(__GT_NeverEquiv, func() interface{} {
			return (&CljsCoreNeverEquiv{})
		})
	}(&AFn{})

	Never_equiv = (&CljsCoreNeverEquiv{})

	Equiv_map = func(equiv_map *AFn) *AFn {
		return Fn(equiv_map, func(x interface{}, y interface{}) interface{} {
			return Boolean.Arity1IB(func() interface{} {
				if Map_QMARK_.Arity1IB(y) {
					return func() interface{} {
						if Count.X_invoke_Arity1(x).(float64) == Count.X_invoke_Arity1(y).(float64) {
							return Every_QMARK_.Arity2IIB(Identity, Map_.X_invoke_Arity2(func(G__497 *AFn) *AFn {
								return Fn(G__497, func(xkv interface{}) interface{} {
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
		return Fn(scan_array, func(incr interface{}, k interface{}, array interface{}) interface{} {
			{
				var len = float64(len(array.([]interface{})))
				_ = len
				{
					var i = float64(0)
					_ = i
					for {
						if i < len {
							if reflect.DeepEqual(k, (array.([]interface{})[int(i)])) {
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
		return Fn(obj_map_compare_keys, func(a interface{}, b interface{}) interface{} {
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
		return Fn(obj_map__GT_hash_map, func(m interface{}, k interface{}, v interface{}) interface{} {
			{
				var ks = Native_get_instance_field.X_invoke_Arity2(m, "Keys")
				var len = float64(len(ks.([]interface{})))
				var so = Native_get_instance_field.X_invoke_Arity2(m, "Strobj")
				var mm = Meta.X_invoke_Arity1(m)
				_, _, _, _ = ks, len, so, mm
				{
					var i = float64(0)
					var out = Transient.X_invoke_Arity1(CljsCorePersistentHashMap_EMPTY)
					_, _ = i, out
					for {
						if i < len {
							{
								var k___1 = (ks.([]interface{})[int(i)])
								_ = k___1
								i, out = (i + float64(1)), Assoc_BANG_.X_invoke_Arity3(out, k___1, (so.([]interface{})[int(k___1.(float64))]))
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

	Obj_clone = func(obj_clone *AFn) *AFn {
		return Fn(obj_clone, func(obj interface{}, ks interface{}) interface{} {
			{
				var new_obj = func() interface{} {
					var obj501 = map[string]interface{}{}
					_ = obj501
					return obj501
				}()
				var l = float64(len(ks.([]interface{})))
				_, _ = new_obj, l
				{
					var i_502 = float64(0)
					_ = i_502
					for {
						if i_502 < l {
							{
								var k_503 = (ks.([]interface{})[int(i_502)])
								_ = k_503
								new_obj.([]interface{})[int(k_503.(float64))] = (obj.([]interface{})[int(k_503.(float64))])
								i_502 = (i_502 + float64(1))
								continue
							}
						} else {
						}
						break
					}
				}
				return new_obj
			}
		})
	}(&AFn{})

	X__GT_ObjMap = func(__GT_ObjMap *AFn) *AFn {
		return Fn(__GT_ObjMap, func(meta interface{}, keys interface{}, strobj interface{}, update_count interface{}, __hash interface{}) interface{} {
			return (&CljsCoreObjMap{meta, keys, strobj, update_count, __hash})
		})
	}(&AFn{})

	X__GT_Iterator = func(__GT_Iterator *AFn) *AFn {
		return Fn(__GT_Iterator, func(s interface{}) interface{} {
			return (&CljsCoreIterator{s})
		})
	}(&AFn{})

	Iterator = func(iterator *AFn) *AFn {
		return Fn(iterator, func(coll interface{}) interface{} {
			return (&CljsCoreIterator{Seq.Arity1IQ(coll)})
		})
	}(&AFn{})

	X__GT_EntriesIterator = func(__GT_EntriesIterator *AFn) *AFn {
		return Fn(__GT_EntriesIterator, func(s interface{}) interface{} {
			return (&CljsCoreEntriesIterator{s})
		})
	}(&AFn{})

	Entries_iterator = func(entries_iterator *AFn) *AFn {
		return Fn(entries_iterator, func(coll interface{}) interface{} {
			return (&CljsCoreEntriesIterator{Seq.Arity1IQ(coll)})
		})
	}(&AFn{})

	X__GT_SetEntriesIterator = func(__GT_SetEntriesIterator *AFn) *AFn {
		return Fn(__GT_SetEntriesIterator, func(s interface{}) interface{} {
			return (&CljsCoreSetEntriesIterator{s})
		})
	}(&AFn{})

	Set_entries_iterator = func(set_entries_iterator *AFn) *AFn {
		return Fn(set_entries_iterator, func(coll interface{}) interface{} {
			return (&CljsCoreSetEntriesIterator{Seq.Arity1IQ(coll)})
		})
	}(&AFn{})

	Array_map_index_of_nil_QMARK_ = func(array_map_index_of_nil_QMARK_ *AFn) *AFn {
		return Fn(array_map_index_of_nil_QMARK_, func(arr interface{}, m interface{}, k interface{}) interface{} {
			{
				var len = float64(len(arr.([]interface{})))
				_ = len
				{
					var i = float64(0)
					_ = i
					for {
						if len <= i {
							return float64(-1)
						} else {
							if Nil_((arr.([]interface{})[int(i)])) {
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
		return Fn(array_map_index_of_keyword_QMARK_, func(arr interface{}, m interface{}, k interface{}) interface{} {
			{
				var len = float64(len(arr.([]interface{})))
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
								var k_SINGLEQUOTE_ = (arr.([]interface{})[int(i)])
								_ = k_SINGLEQUOTE_
								return (func() bool { _, instanceof := k_SINGLEQUOTE_.(*CljsCoreKeyword); return instanceof }()) && (reflect.DeepEqual(kstr, Native_get_instance_field.X_invoke_Arity2(k_SINGLEQUOTE_, "Fqn")))
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
		return Fn(array_map_index_of_symbol_QMARK_, func(arr interface{}, m interface{}, k interface{}) interface{} {
			{
				var len = float64(len(arr.([]interface{})))
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
								var k_SINGLEQUOTE_ = (arr.([]interface{})[int(i)])
								_ = k_SINGLEQUOTE_
								return (func() bool { _, instanceof := k_SINGLEQUOTE_.(*CljsCoreSymbol); return instanceof }()) && (reflect.DeepEqual(kstr, Native_get_instance_field.X_invoke_Arity2(k_SINGLEQUOTE_, "Str")))
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
		return Fn(array_map_index_of_identical_QMARK_, func(arr interface{}, m interface{}, k interface{}) interface{} {
			{
				var len = float64(len(arr.([]interface{})))
				_ = len
				{
					var i = float64(0)
					_ = i
					for {
						if len <= i {
							return float64(-1)
						} else {
							if reflect.DeepEqual(k, (arr.([]interface{})[int(i)])) {
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
		return Fn(array_map_index_of_equiv_QMARK_, func(arr interface{}, m interface{}, k interface{}) interface{} {
			{
				var len = float64(len(arr.([]interface{})))
				_ = len
				{
					var i = float64(0)
					_ = i
					for {
						if len <= i {
							return float64(-1)
						} else {
							if X_EQ_.Arity2IIB(k, (arr.([]interface{})[int(i)])) {
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
		return Fn(array_map_index_of, func(m interface{}, k interface{}) interface{} {
			{
				var arr = Native_get_instance_field.X_invoke_Arity2(m, "Arr")
				_ = arr
				if func() bool { _, instanceof := k.(*CljsCoreKeyword); return instanceof }() {
					return Array_map_index_of_keyword_QMARK_.X_invoke_Arity3(arr, m, k).(float64)
				} else {
					if (Truth_(Native_invoke_func.X_invoke_Arity2(goog.IsString, []interface{}{k}))) || (reflect.ValueOf(k).Kind() == reflect.Float64) {
						return Array_map_index_of_identical_QMARK_.X_invoke_Arity3(arr, m, k).(float64)
					} else {
						if func() bool { _, instanceof := k.(*CljsCoreSymbol); return instanceof }() {
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
		return Fn(array_map_extend_kv, func(m interface{}, k interface{}, v interface{}) interface{} {
			{
				var arr = Native_get_instance_field.X_invoke_Arity2(m, "Arr")
				var l = float64(len(arr.([]interface{})))
				var narr = make([]interface{}, int((l + float64(2))))
				_, _, _ = arr, l, narr
				{
					var i_510 = float64(0)
					_ = i_510
					for {
						if i_510 < l {
							narr[int(i_510)] = (arr.([]interface{})[int(i_510)])
							i_510 = (i_510 + float64(1))
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
		return Fn(__GT_PersistentArrayMapSeq, func(arr interface{}, i interface{}, _meta interface{}) interface{} {
			return (&CljsCorePersistentArrayMapSeq{arr, i, _meta})
		})
	}(&AFn{})

	Persistent_array_map_seq = func(persistent_array_map_seq *AFn) *AFn {
		return Fn(persistent_array_map_seq, func(arr interface{}, i interface{}, _meta interface{}) interface{} {
			if i.(float64) <= (float64(len(arr.([]interface{}))) - float64(2)) {
				return (&CljsCorePersistentArrayMapSeq{arr, i, _meta})
			} else {
				return nil
			}
		})
	}(&AFn{})

	X__GT_PersistentArrayMap = func(__GT_PersistentArrayMap *AFn) *AFn {
		return Fn(__GT_PersistentArrayMap, func(meta interface{}, cnt interface{}, arr interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentArrayMap{meta, cnt, arr, __hash})
		})
	}(&AFn{})

	X__GT_TransientArrayMap = func(__GT_TransientArrayMap *AFn) *AFn {
		return Fn(__GT_TransientArrayMap, func(editable_QMARK_ interface{}, len interface{}, arr interface{}) interface{} {
			return (&CljsCoreTransientArrayMap{editable_QMARK_, len, arr})
		})
	}(&AFn{})

	Array__GT_transient_hash_map = func(array__GT_transient_hash_map *AFn) *AFn {
		return Fn(array__GT_transient_hash_map, func(len interface{}, arr interface{}) interface{} {
			{
				var out = Transient.X_invoke_Arity1(CljsCorePersistentHashMap_EMPTY)
				var i = float64(0)
				_, _ = out, i
				for {
					if i < len.(float64) {
						out, i = Assoc_BANG_.X_invoke_Arity3(out, (arr.([]interface{})[int(i)]), (arr.([]interface{})[int((i+float64(1)))])), (i + float64(2))
						continue
					} else {
						return out
					}
				}
			}
		})
	}(&AFn{})

	X__GT_Box = func(__GT_Box *AFn) *AFn {
		return Fn(__GT_Box, func(val interface{}) interface{} {
			return (&CljsCoreBox{val})
		})
	}(&AFn{})

	Key_test = func(key_test *AFn) *AFn {
		return Fn(key_test, func(key interface{}, other interface{}) bool {
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
		return Fn(mask, func(hash interface{}, shift interface{}) interface{} {
			return float64(int(float64(uint(hash.(float64))>>uint(shift.(float64)))) & int(float64(31)))
		})
	}(&AFn{})

	Clone_and_set = func(clone_and_set *AFn) *AFn {
		return Fn(clone_and_set, func(arr interface{}, i interface{}, a interface{}) interface{} {
			{
				var G__534 = Aclone.X_invoke_Arity1(arr).([]interface{})
				_ = G__534
				G__534[int(i.(float64))] = a
				return G__534
			}
		}, func(arr interface{}, i interface{}, a interface{}, j interface{}, b interface{}) interface{} {
			{
				var G__535 = Aclone.X_invoke_Arity1(arr).([]interface{})
				_ = G__535
				G__535[int(i.(float64))] = a
				G__535[int(j.(float64))] = b
				return G__535
			}
		})
	}(&AFn{})

	Remove_pair = func(remove_pair *AFn) *AFn {
		return Fn(remove_pair, func(arr interface{}, i interface{}) interface{} {
			{
				var new_arr = make([]interface{}, int((float64(len(arr.([]interface{}))) - float64(2))))
				_ = new_arr
				Array_copy.X_invoke_Arity5(arr, float64(0), new_arr, float64(0), (float64(2) * i.(float64)))
				Array_copy.X_invoke_Arity5(arr, (float64(2) * (i.(float64) + float64(1))), new_arr, (float64(2) * i.(float64)), (float64(len(new_arr)) - (float64(2) * i.(float64))))
				return new_arr
			}
		})
	}(&AFn{})

	Bitmap_indexed_node_index = func(bitmap_indexed_node_index *AFn) *AFn {
		return Fn(bitmap_indexed_node_index, func(bitmap interface{}, bit interface{}) interface{} {
			return Bit_count.X_invoke_Arity1(float64(int(bitmap.(float64)) & int((bit.(float64) - float64(1))))).(float64)
		})
	}(&AFn{})

	Bitpos = func(bitpos *AFn) *AFn {
		return Fn(bitpos, func(hash interface{}, shift interface{}) interface{} {
			return float64(int(float64(1)) << uint(float64((uint(hash.(float64))>>uint(shift.(float64)))&0x01f)))
		})
	}(&AFn{})

	Edit_and_set = func(edit_and_set *AFn) *AFn {
		return Fn(edit_and_set, func(inode interface{}, edit interface{}, i interface{}, a interface{}) interface{} {
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
		return Fn(inode_kv_reduce, func(arr interface{}, f interface{}, init interface{}) interface{} {
			{
				var len = float64(len(arr.([]interface{})))
				_ = len
				{
					var i = float64(0)
					var init___1 = init
					_, _ = i, init___1
					for {
						if i < len {
							{
								var init___2 = func() interface{} {
									var k = (arr.([]interface{})[int(i)])
									_ = k
									if !(Nil_(k)) {
										return f.(CljsCoreIFn).X_invoke_Arity3(init___1, k, (arr.([]interface{})[int((i + float64(1)))]))
									} else {
										{
											var node = (arr.([]interface{})[int((i + float64(1)))])
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
		return Fn(__GT_BitmapIndexedNode, func(edit interface{}, bitmap interface{}, arr interface{}) interface{} {
			return (&CljsCoreBitmapIndexedNode{edit, bitmap, arr})
		})
	}(&AFn{})

	Pack_array_node = func(pack_array_node *AFn) *AFn {
		return Fn(pack_array_node, func(array_node interface{}, edit interface{}, idx interface{}) interface{} {
			{
				var arr = Native_get_instance_field.X_invoke_Arity2(array_node, "Arr")
				var len = float64(len(arr.([]interface{})))
				var new_arr = make([]interface{}, int((float64(2) * (Native_get_instance_field.X_invoke_Arity2(array_node, "Cnt").(float64) - float64(1)))))
				_, _, _ = arr, len, new_arr
				{
					var i = float64(0)
					var j = float64(1)
					var bitmap = float64(0)
					_, _, _ = i, j, bitmap
					for {
						if i < len {
							if (!(i == idx.(float64))) && (!(Nil_((arr.([]interface{})[int(i)])))) {
								new_arr[int(j)] = (arr.([]interface{})[int(i)])
								i, j, bitmap = (i + float64(1)), (j + float64(2)), float64(int(bitmap)|int(float64(int(float64(1))<<uint(i))))
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
		return Fn(__GT_ArrayNode, func(edit interface{}, cnt interface{}, arr interface{}) interface{} {
			return (&CljsCoreArrayNode{edit, cnt, arr})
		})
	}(&AFn{})

	Hash_collision_node_find_index = func(hash_collision_node_find_index *AFn) *AFn {
		return Fn(hash_collision_node_find_index, func(arr interface{}, cnt interface{}, key interface{}) interface{} {
			{
				var lim = (float64(2) * cnt.(float64))
				_ = lim
				{
					var i = float64(0)
					_ = i
					for {
						if i < lim {
							if Key_test.Arity2IIB(key, (arr.([]interface{})[int(i)])) {
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
		return Fn(__GT_HashCollisionNode, func(edit interface{}, collision_hash interface{}, cnt interface{}, arr interface{}) interface{} {
			return (&CljsCoreHashCollisionNode{edit, collision_hash, cnt, arr})
		})
	}(&AFn{})

	Create_node = func(create_node *AFn) *AFn {
		return Fn(create_node, func(shift interface{}, key1 interface{}, val1 interface{}, key2hash interface{}, key2 interface{}, val2 interface{}) interface{} {
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
		return Fn(__GT_NodeSeq, func(meta interface{}, nodes interface{}, i interface{}, s interface{}, __hash interface{}) interface{} {
			return (&CljsCoreNodeSeq{meta, nodes, i, s, __hash})
		})
	}(&AFn{})

	Create_inode_seq = func(create_inode_seq *AFn) *AFn {
		return Fn(create_inode_seq, func(nodes interface{}) interface{} {
			return create_inode_seq.X_invoke_Arity3(nodes, float64(0), nil)
		}, func(nodes interface{}, i interface{}, s interface{}) interface{} {
			if Nil_(s) {
				{
					var len = float64(len(nodes.([]interface{})))
					_ = len
					{
						var j = i
						_ = j
						for {
							if j.(float64) < len {
								if !(Nil_((nodes.([]interface{})[int(j.(float64))]))) {
									return (&CljsCoreNodeSeq{nil, nodes, j, nil, nil})
								} else {
									{
										var temp__4220__auto__ = (nodes.([]interface{})[int((j.(float64) + float64(1)))])
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
		return Fn(__GT_ArrayNodeSeq, func(meta interface{}, nodes interface{}, i interface{}, s interface{}, __hash interface{}) interface{} {
			return (&CljsCoreArrayNodeSeq{meta, nodes, i, s, __hash})
		})
	}(&AFn{})

	Create_array_node_seq = func(create_array_node_seq *AFn) *AFn {
		return Fn(create_array_node_seq, func(nodes interface{}) interface{} {
			return create_array_node_seq.X_invoke_Arity4(nil, nodes, float64(0), nil)
		}, func(meta interface{}, nodes interface{}, i interface{}, s interface{}) interface{} {
			if Nil_(s) {
				{
					var len = float64(len(nodes.([]interface{})))
					_ = len
					{
						var j = i
						_ = j
						for {
							if j.(float64) < len {
								{
									var temp__4220__auto__ = (nodes.([]interface{})[int(j.(float64))])
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
		return Fn(__GT_PersistentHashMap, func(meta interface{}, cnt interface{}, root interface{}, has_nil_QMARK_ bool, nil_val interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentHashMap{meta, cnt, root, has_nil_QMARK_, nil_val, __hash})
		})
	}(&AFn{})

	X__GT_TransientHashMap = func(__GT_TransientHashMap *AFn) *AFn {
		return Fn(__GT_TransientHashMap, func(edit bool, root interface{}, count interface{}, has_nil_QMARK_ bool, nil_val interface{}) interface{} {
			return (&CljsCoreTransientHashMap{edit, root, count, has_nil_QMARK_, nil_val})
		})
	}(&AFn{})

	Tree_map_seq_push = func(tree_map_seq_push *AFn) *AFn {
		return Fn(tree_map_seq_push, func(node interface{}, stack interface{}, ascending_QMARK_ bool) interface{} {
			{
				var t = node
				var stack___1 = stack
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
		return Fn(__GT_PersistentTreeMapSeq, func(meta interface{}, stack interface{}, ascending_QMARK_ bool, cnt interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentTreeMapSeq{meta, stack, ascending_QMARK_, cnt, __hash})
		})
	}(&AFn{})

	Create_tree_map_seq = func(create_tree_map_seq *AFn) *AFn {
		return Fn(create_tree_map_seq, func(tree interface{}, ascending_QMARK_ interface{}, cnt interface{}) interface{} {
			return (&CljsCorePersistentTreeMapSeq{nil, Tree_map_seq_push.X_invoke_Arity3(tree, nil, ascending_QMARK_), ascending_QMARK_.(bool), cnt, nil})
		})
	}(&AFn{})

	Balance_left = func(balance_left *AFn) *AFn {
		return Fn(balance_left, func(key interface{}, val interface{}, ins interface{}, right interface{}) interface{} {
			if func() bool { _, instanceof := ins.(*CljsCoreRedNode); return instanceof }() {
				if func() bool {
					_, instanceof := Native_get_instance_field.X_invoke_Arity2(ins, "Left").(*CljsCoreRedNode)
					return instanceof
				}() {
					return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(ins, "Key"), Native_get_instance_field.X_invoke_Arity2(ins, "Val"), Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(ins, "Left"), "Blacken", []interface{}{}), (&CljsCoreBlackNode{key, val, Native_get_instance_field.X_invoke_Arity2(ins, "Right"), right, nil}), nil})
				} else {
					if func() bool {
						_, instanceof := Native_get_instance_field.X_invoke_Arity2(ins, "Right").(*CljsCoreRedNode)
						return instanceof
					}() {
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
		return Fn(balance_right, func(key interface{}, val interface{}, left interface{}, ins interface{}) interface{} {
			if func() bool { _, instanceof := ins.(*CljsCoreRedNode); return instanceof }() {
				if func() bool {
					_, instanceof := Native_get_instance_field.X_invoke_Arity2(ins, "Right").(*CljsCoreRedNode)
					return instanceof
				}() {
					return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(ins, "Key"), Native_get_instance_field.X_invoke_Arity2(ins, "Val"), (&CljsCoreBlackNode{key, val, left, Native_get_instance_field.X_invoke_Arity2(ins, "Left"), nil}), Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(ins, "Right"), "Blacken", []interface{}{}), nil})
				} else {
					if func() bool {
						_, instanceof := Native_get_instance_field.X_invoke_Arity2(ins, "Left").(*CljsCoreRedNode)
						return instanceof
					}() {
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
		return Fn(balance_left_del, func(key interface{}, val interface{}, del interface{}, right interface{}) interface{} {
			if func() bool { _, instanceof := del.(*CljsCoreRedNode); return instanceof }() {
				return (&CljsCoreRedNode{key, val, Native_invoke_instance_method.X_invoke_Arity3(del, "Blacken", []interface{}{}), right, nil})
			} else {
				if func() bool { _, instanceof := right.(*CljsCoreBlackNode); return instanceof }() {
					return Balance_right.X_invoke_Arity4(key, val, del, Native_invoke_instance_method.X_invoke_Arity3(right, "Redden", []interface{}{}))
				} else {
					if (func() bool { _, instanceof := right.(*CljsCoreRedNode); return instanceof }()) && (func() bool {
						_, instanceof := Native_get_instance_field.X_invoke_Arity2(right, "Left").(*CljsCoreBlackNode)
						return instanceof
					}()) {
						return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(right, "Left"), "Key"), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(right, "Left"), "Val"), (&CljsCoreBlackNode{key, val, del, Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(right, "Left"), "Left"), nil}), Balance_right.X_invoke_Arity4(Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(right, "Left"), "Right"), Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(right, "Right"), "Redden", []interface{}{})), nil})
					} else {
						panic((&js.Error{"red-black tree invariant violation"}))

					}
				}
			}
		})
	}(&AFn{})

	Balance_right_del = func(balance_right_del *AFn) *AFn {
		return Fn(balance_right_del, func(key interface{}, val interface{}, left interface{}, del interface{}) interface{} {
			if func() bool { _, instanceof := del.(*CljsCoreRedNode); return instanceof }() {
				return (&CljsCoreRedNode{key, val, left, Native_invoke_instance_method.X_invoke_Arity3(del, "Blacken", []interface{}{}), nil})
			} else {
				if func() bool { _, instanceof := left.(*CljsCoreBlackNode); return instanceof }() {
					return Balance_left.X_invoke_Arity4(key, val, Native_invoke_instance_method.X_invoke_Arity3(left, "Redden", []interface{}{}), del)
				} else {
					if (func() bool { _, instanceof := left.(*CljsCoreRedNode); return instanceof }()) && (func() bool {
						_, instanceof := Native_get_instance_field.X_invoke_Arity2(left, "Right").(*CljsCoreBlackNode)
						return instanceof
					}()) {
						return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), "Key"), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), "Val"), Balance_left.X_invoke_Arity4(Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(left, "Left"), "Redden", []interface{}{}), Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), "Left")), (&CljsCoreBlackNode{key, val, Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), "Right"), del, nil}), nil})
					} else {
						panic((&js.Error{"red-black tree invariant violation"}))

					}
				}
			}
		})
	}(&AFn{})

	Tree_map_kv_reduce = func(tree_map_kv_reduce *AFn) *AFn {
		return Fn(tree_map_kv_reduce, func(node interface{}, f interface{}, init interface{}) interface{} {
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
						var init___2 = f.(CljsCoreIFn).X_invoke_Arity3(init___1, Native_get_instance_field.X_invoke_Arity2(node, "Key"), Native_get_instance_field.X_invoke_Arity2(node, "Val"))
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
		return Fn(__GT_BlackNode, func(key interface{}, val interface{}, left interface{}, right interface{}, __hash interface{}) interface{} {
			return (&CljsCoreBlackNode{key, val, left, right, __hash})
		})
	}(&AFn{})

	X__GT_RedNode = func(__GT_RedNode *AFn) *AFn {
		return Fn(__GT_RedNode, func(key interface{}, val interface{}, left interface{}, right interface{}, __hash interface{}) interface{} {
			return (&CljsCoreRedNode{key, val, left, right, __hash})
		})
	}(&AFn{})

	Tree_map_add = func(tree_map_add *AFn) *AFn {
		return Fn(tree_map_add, func(comp interface{}, tree interface{}, k interface{}, v interface{}, found interface{}) interface{} {
			if Nil_(tree) {
				return (&CljsCoreRedNode{k, v, nil, nil, nil})
			} else {
				{
					var c = comp.(CljsCoreIFn).X_invoke_Arity2(k, Native_get_instance_field.X_invoke_Arity2(tree, "Key"))
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
		return Fn(tree_map_append, func(left interface{}, right interface{}) interface{} {
			if Nil_(left) {
				return right
			} else {
				if Nil_(right) {
					return left
				} else {
					if func() bool { _, instanceof := left.(*CljsCoreRedNode); return instanceof }() {
						if func() bool { _, instanceof := right.(*CljsCoreRedNode); return instanceof }() {
							{
								var app = tree_map_append.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), Native_get_instance_field.X_invoke_Arity2(right, "Left"))
								_ = app
								if func() bool { _, instanceof := app.(*CljsCoreRedNode); return instanceof }() {
									return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(app, "Key"), Native_get_instance_field.X_invoke_Arity2(app, "Val"), (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_get_instance_field.X_invoke_Arity2(left, "Left"), Native_get_instance_field.X_invoke_Arity2(app, "Left"), nil}), (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), Native_get_instance_field.X_invoke_Arity2(app, "Right"), Native_get_instance_field.X_invoke_Arity2(right, "Right"), nil}), nil})
								} else {
									return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_get_instance_field.X_invoke_Arity2(left, "Left"), (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), app, Native_get_instance_field.X_invoke_Arity2(right, "Right"), nil}), nil})
								}
							}
						} else {
							return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_get_instance_field.X_invoke_Arity2(left, "Left"), tree_map_append.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), right), nil})
						}
					} else {
						if func() bool { _, instanceof := right.(*CljsCoreRedNode); return instanceof }() {
							return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), tree_map_append.X_invoke_Arity2(left, Native_get_instance_field.X_invoke_Arity2(right, "Left")), Native_get_instance_field.X_invoke_Arity2(right, "Right"), nil})
						} else {
							{
								var app = tree_map_append.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), Native_get_instance_field.X_invoke_Arity2(right, "Left"))
								_ = app
								if func() bool { _, instanceof := app.(*CljsCoreRedNode); return instanceof }() {
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
		return Fn(tree_map_remove, func(comp interface{}, tree interface{}, k interface{}, found interface{}) interface{} {
			if !(Nil_(tree)) {
				{
					var c = comp.(CljsCoreIFn).X_invoke_Arity2(k, Native_get_instance_field.X_invoke_Arity2(tree, "Key"))
					_ = c
					if c.(float64) == float64(0) {
						found.([]interface{})[int(float64(0))] = tree
						return Tree_map_append.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tree, "Left"), Native_get_instance_field.X_invoke_Arity2(tree, "Right"))
					} else {
						if c.(float64) < float64(0) {
							{
								var del = tree_map_remove.X_invoke_Arity4(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Left"), k, found)
								_ = del
								if (!(Nil_(del))) || (!(Nil_((found.([]interface{})[int(float64(0))])))) {
									if func() bool {
										_, instanceof := Native_get_instance_field.X_invoke_Arity2(tree, "Left").(*CljsCoreBlackNode)
										return instanceof
									}() {
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
								if (!(Nil_(del))) || (!(Nil_((found.([]interface{})[int(float64(0))])))) {
									if func() bool {
										_, instanceof := Native_get_instance_field.X_invoke_Arity2(tree, "Right").(*CljsCoreBlackNode)
										return instanceof
									}() {
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
		return Fn(tree_map_replace, func(comp interface{}, tree interface{}, k interface{}, v interface{}) interface{} {
			{
				var tk = Native_get_instance_field.X_invoke_Arity2(tree, "Key")
				var c = comp.(CljsCoreIFn).X_invoke_Arity2(k, tk)
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
		return Fn(__GT_PersistentTreeMap, func(comp interface{}, tree interface{}, cnt interface{}, meta interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentTreeMap{comp, tree, cnt, meta, __hash})
		})
	}(&AFn{})

	Hash_map = func(hash_map *AFn) *AFn {
		return Fn(hash_map, func(keyvals__ ...interface{}) interface{} {
			var keyvals = Array_seq.X_invoke_Arity1(keyvals__[0:])
			_ = keyvals
			{
				var in = Seq.Arity1IQ(keyvals)
				var out = Transient.X_invoke_Arity1(CljsCorePersistentHashMap_EMPTY)
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
		return Fn(array_map, func(keyvals__ ...interface{}) interface{} {
			var keyvals = Array_seq.X_invoke_Arity1(keyvals__[0:])
			_ = keyvals
			return (&CljsCorePersistentArrayMap{nil, Quot.X_invoke_Arity2(Count.X_invoke_Arity1(keyvals).(float64), float64(2)).(float64), Apply.X_invoke_Arity2(Array, keyvals), nil})
		})
	}(&AFn{})

	Obj_map = func(obj_map *AFn) *AFn {
		return Fn(obj_map, func(keyvals__ ...interface{}) interface{} {
			var keyvals = Array_seq.X_invoke_Arity1(keyvals__[0:])
			_ = keyvals
			{
				var ks = []interface{}{}
				var obj = func() interface{} {
					var obj568 = map[string]interface{}{}
					_ = obj568
					return obj568
				}()
				_, _ = ks, obj
				{
					var kvs = Seq.Arity1IQ(keyvals)
					_ = kvs
					for {
						if Truth_(kvs) {
							js.JSArray_(&ks).Push(First.X_invoke_Arity1(kvs))
							obj.([]interface{})[int(First.X_invoke_Arity1(kvs).(float64))] = Second.X_invoke_Arity1(kvs)
							kvs = Seq_(Nnext.X_invoke_Arity1(kvs))
							continue
						} else {
							return CljsCoreObjMap_FromObject.X_invoke_Arity2(ks, obj)
						}
					}
				}
			}
		})
	}(&AFn{})

	Sorted_map = func(sorted_map *AFn) *AFn {
		return Fn(sorted_map, func(keyvals__ ...interface{}) interface{} {
			var keyvals = Array_seq.X_invoke_Arity1(keyvals__[0:])
			_ = keyvals
			{
				var in = Seq.Arity1IQ(keyvals)
				var out = CljsCorePersistentTreeMap_EMPTY
				_, _ = in, out
				for {
					if Truth_(in) {
						in, out = Seq_(Nnext.X_invoke_Arity1(in)), Assoc.X_invoke_Arity3(out, First.X_invoke_Arity1(in), Second.X_invoke_Arity1(in)).(*CljsCorePersistentTreeMap)
						continue
					} else {
						return out
					}
				}
			}
		})
	}(&AFn{})

	Sorted_map_by = func(sorted_map_by *AFn) *AFn {
		return Fn(sorted_map_by, func(comparator_keyvals__ ...interface{}) interface{} {
			var comparator = comparator_keyvals__[0]
			var keyvals = Array_seq.X_invoke_Arity1(comparator_keyvals__[1:])
			_, _ = comparator, keyvals
			{
				var in = Seq.Arity1IQ(keyvals)
				var out = (&CljsCorePersistentTreeMap{Fn__GT_comparator.X_invoke_Arity1(comparator), nil, float64(0), nil, float64(0)})
				_, _ = in, out
				for {
					if Truth_(in) {
						in, out = Seq_(Nnext.X_invoke_Arity1(in)), Assoc.X_invoke_Arity3(out, First.X_invoke_Arity1(in), Second.X_invoke_Arity1(in)).(*CljsCorePersistentTreeMap)
						continue
					} else {
						return out
					}
				}
			}
		})
	}(&AFn{})

	X__GT_KeySeq = func(__GT_KeySeq *AFn) *AFn {
		return Fn(__GT_KeySeq, func(mseq interface{}, _meta interface{}) interface{} {
			return (&CljsCoreKeySeq{mseq, _meta})
		})
	}(&AFn{})

	Keys = func(keys *AFn) *AFn {
		return Fn(keys, func(hash_map interface{}) interface{} {
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
		return Fn(key, func(map_entry interface{}) interface{} {
			return map_entry.(CljsCoreIMapEntry).X_key_Arity1()
		})
	}(&AFn{})

	X__GT_ValSeq = func(__GT_ValSeq *AFn) *AFn {
		return Fn(__GT_ValSeq, func(mseq interface{}, _meta interface{}) interface{} {
			return (&CljsCoreValSeq{mseq, _meta})
		})
	}(&AFn{})

	Vals = func(vals *AFn) *AFn {
		return Fn(vals, func(hash_map interface{}) interface{} {
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
		return Fn(val, func(map_entry interface{}) interface{} {
			return map_entry.(CljsCoreIMapEntry).X_val_Arity1()
		})
	}(&AFn{})

	Merge = func(merge *AFn) *AFn {
		return Fn(merge, func(maps__ ...interface{}) interface{} {
			var maps = Array_seq.X_invoke_Arity1(maps__[0:])
			_ = maps
			if Truth_(Some.X_invoke_Arity2(Identity, maps)) {
				return Reduce.X_invoke_Arity2(func(G__571 *AFn) *AFn {
					return Fn(G__571, func(p1__569_SHARP_ interface{}, p2__570_SHARP_ interface{}) interface{} {
						return Conj.X_invoke_Arity2(func() interface{} {
							var or__175__auto__ = p1__569_SHARP_
							_ = or__175__auto__
							if Truth_(or__175__auto__) {
								return or__175__auto__
							} else {
								return CljsCorePersistentArrayMap_EMPTY
							}
						}(), p2__570_SHARP_)
					})
				}(&AFn{}), maps)
			} else {
				return nil
			}
		})
	}(&AFn{})

	Merge_with = func(merge_with *AFn) *AFn {
		return Fn(merge_with, func(f_maps__ ...interface{}) interface{} {
			var f = f_maps__[0]
			var maps = Array_seq.X_invoke_Arity1(f_maps__[1:])
			_, _ = f, maps
			if Truth_(Some.X_invoke_Arity2(Identity, maps)) {
				{
					var merge_entry = func(G__572 *AFn) *AFn {
						return Fn(G__572, func(m interface{}, e interface{}) interface{} {
							{
								var k = First.X_invoke_Arity1(e)
								var v = Second.X_invoke_Arity1(e)
								_, _ = k, v
								if Contains_QMARK_.Arity2IIB(m, k) {
									return Assoc.X_invoke_Arity3(m, k, f.(CljsCoreIFn).X_invoke_Arity2(Get.X_invoke_Arity2(m, k), v))
								} else {
									return Assoc.X_invoke_Arity3(m, k, v)
								}
							}
						})
					}(&AFn{})
					var merge2 = func(G__573 *AFn, merge_entry CljsCoreIFn) *AFn {
						return Fn(G__573, func(m1 interface{}, m2 interface{}) interface{} {
							return Reduce.X_invoke_Arity3(merge_entry, func() interface{} {
								var or__175__auto__ = m1
								_ = or__175__auto__
								if Truth_(or__175__auto__) {
									return or__175__auto__
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
		return Fn(select_keys, func(map_ interface{}, keyseq interface{}) interface{} {
			{
				var ret = CljsCorePersistentArrayMap_EMPTY
				var keys = Seq.Arity1IQ(keyseq)
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
							}().(*CljsCorePersistentArrayMap), Next.Arity1IQ(keys)
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
		return Fn(__GT_PersistentHashSet, func(meta interface{}, hash_map interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentHashSet{meta, hash_map, __hash})
		})
	}(&AFn{})

	X__GT_TransientHashSet = func(__GT_TransientHashSet *AFn) *AFn {
		return Fn(__GT_TransientHashSet, func(transient_map interface{}) interface{} {
			return (&CljsCoreTransientHashSet{transient_map})
		})
	}(&AFn{})

	X__GT_PersistentTreeSet = func(__GT_PersistentTreeSet *AFn) *AFn {
		return Fn(__GT_PersistentTreeSet, func(meta interface{}, tree_map interface{}, __hash interface{}) interface{} {
			return (&CljsCorePersistentTreeSet{meta, tree_map, __hash})
		})
	}(&AFn{})

	Set_from_indexed_seq = func(set_from_indexed_seq *AFn) *AFn {
		return Fn(set_from_indexed_seq, func(iseq interface{}) interface{} {
			{
				var arr = Native_get_instance_field.X_invoke_Arity2(iseq, "Arr")
				var ret = func() interface{} {
					var a__1036__auto__ = arr
					_ = a__1036__auto__
					{
						var i = float64(0)
						var res = CljsCorePersistentHashSet_EMPTY.X_as_transient_Arity1()
						_, _ = i, res
						for {
							if i < float64(len(a__1036__auto__.([]interface{}))) {
								i, res = (i + float64(1)), res.(CljsCoreITransientCollection).X_conj_BANG__Arity2((arr.([]interface{})[int(i)]))
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
		return Fn(set, func(coll interface{}) interface{} {
			{
				var in = Seq.Arity1IQ(coll)
				_ = in
				if Nil_(in) {
					return CljsCorePersistentHashSet_EMPTY
				} else {
					if (func() bool { _, instanceof := in.(*CljsCoreIndexedSeq); return instanceof }()) && (Native_get_instance_field.X_invoke_Arity2(in, "I").(float64) == float64(0)) {
						return Set_from_indexed_seq.X_invoke_Arity1(in)
					} else {
						{
							var in___1 = in
							var out = CljsCorePersistentHashSet_EMPTY.X_as_transient_Arity1()
							_, _ = in___1, out
							for {
								if !(Nil_(in___1)) {
									in___1, out = in___1.(CljsCoreINext).X_next_Arity1().(CljsCoreISeq), out.(CljsCoreITransientCollection).X_conj_BANG__Arity2(in___1.(CljsCoreISeq).X_first_Arity1())
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
		return Fn(hash_set, func() interface{} {
			return CljsCorePersistentHashSet_EMPTY
		}, func(keys__ ...interface{}) interface{} {
			var keys = Array_seq.X_invoke_Arity1(keys__[0:])
			_ = keys
			return Set.X_invoke_Arity1(keys)
		})
	}(&AFn{})

	Sorted_set = func(sorted_set *AFn) *AFn {
		return Fn(sorted_set, func(keys__ ...interface{}) interface{} {
			var keys = Array_seq.X_invoke_Arity1(keys__[0:])
			_ = keys
			return Reduce.X_invoke_Arity3(X_conj, CljsCorePersistentTreeSet_EMPTY, keys)
		})
	}(&AFn{})

	Sorted_set_by = func(sorted_set_by *AFn) *AFn {
		return Fn(sorted_set_by, func(comparator_keys__ ...interface{}) interface{} {
			var comparator = comparator_keys__[0]
			var keys = Array_seq.X_invoke_Arity1(comparator_keys__[1:])
			_, _ = comparator, keys
			return Reduce.X_invoke_Arity3(X_conj, (&CljsCorePersistentTreeSet{nil, Sorted_map_by.X_invoke_Arity1(comparator).(*CljsCorePersistentTreeMap), float64(0)}), keys)
		})
	}(&AFn{})

	Replace = func(replace *AFn) *AFn {
		return Fn(replace, func(smap interface{}) interface{} {
			return Map_.X_invoke_Arity1(func(G__603 *AFn) *AFn {
				return Fn(G__603, func(p1__601_SHARP_ interface{}) interface{} {
					{
						var temp__4220__auto__ = Find.X_invoke_Arity2(smap, p1__601_SHARP_)
						_ = temp__4220__auto__
						if Truth_(temp__4220__auto__) {
							{
								var e = temp__4220__auto__
								_ = e
								return Val.X_invoke_Arity1(e)
							}
						} else {
							return p1__601_SHARP_
						}
					}
				})
			}(&AFn{})).(CljsCoreIFn)
		}, func(smap interface{}, coll interface{}) interface{} {
			if Vector_QMARK_.Arity1IB(coll) {
				{
					var n = Count.X_invoke_Arity1(coll).(float64)
					_ = n
					return Reduce.X_invoke_Arity3(func(G__604 *AFn, n float64) *AFn {
						return Fn(G__604, func(v interface{}, i interface{}) interface{} {
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
				return Map_.X_invoke_Arity2(func(G__605 *AFn) *AFn {
					return Fn(G__605, func(p1__602_SHARP_ interface{}) interface{} {
						{
							var temp__4220__auto__ = Find.X_invoke_Arity2(smap, p1__602_SHARP_)
							_ = temp__4220__auto__
							if Truth_(temp__4220__auto__) {
								{
									var e = temp__4220__auto__
									_ = e
									return Second.X_invoke_Arity1(e)
								}
							} else {
								return p1__602_SHARP_
							}
						}
					})
				}(&AFn{}), coll).(*CljsCoreLazySeq)
			}
		})
	}(&AFn{})

	Distinct = func(distinct *AFn) *AFn {
		return Fn(distinct, func(coll interface{}) interface{} {
			{
				var step = func(step *AFn) *AFn {
					return Fn(step, func(xs interface{}, seen interface{}) interface{} {
						return (&CljsCoreLazySeq{nil, func(G__614 *AFn) *AFn {
							return Fn(G__614, func() interface{} {
								return func(G__615 *AFn) *AFn {
									return Fn(G__615, func(p__612 interface{}, seen___1 interface{}) interface{} {
										for {
											{
												var vec__613 = p__612
												var f = Nth.X_invoke_Arity3(vec__613, float64(0), nil)
												var xs___1 = vec__613
												_, _, _ = vec__613, f, xs___1
												{
													var temp__4222__auto__ = Seq.Arity1IQ(xs___1)
													_ = temp__4222__auto__
													if Truth_(temp__4222__auto__) {
														{
															var s = temp__4222__auto__
															_ = s
															if Contains_QMARK_.Arity2IIB(seen___1, f) {
																p__612, seen___1 = Rest.Arity1IQ(s), seen___1
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
		return Fn(butlast, func(s interface{}) interface{} {
			{
				var ret = CljsCorePersistentVector_EMPTY
				var s___1 = s
				_, _ = ret, s___1
				for {
					if Truth_(Next.Arity1IQ(s___1)) {
						ret, s___1 = Conj.X_invoke_Arity2(ret, First.X_invoke_Arity1(s___1)).(*CljsCorePersistentVector), Next.Arity1IQ(s___1)
						continue
					} else {
						return Seq.Arity1IQ(ret)
					}
				}
			}
		})
	}(&AFn{})

	Name = func(name *AFn) *AFn {
		return Fn(name, func(x interface{}) interface{} {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INamed", Str: "cljs.core/INamed", X_hash: float64(-857199025), X_meta: nil}), x)) {
				return x.(CljsCoreINamed).X_name_Arity1()
			} else {
				if reflect.ValueOf(x).Kind() == reflect.String {
					return x
				} else {
					panic((&js.Error{("Doesn't support name: " + Str.X_invoke_Arity1(x).(string))}))
				}
			}
		})
	}(&AFn{})

	Zipmap = func(zipmap *AFn) *AFn {
		return Fn(zipmap, func(keys interface{}, vals interface{}) interface{} {
			{
				var map_ = Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY)
				var ks = Seq.Arity1IQ(keys)
				var vs = Seq.Arity1IQ(vals)
				_, _, _ = map_, ks, vs
				for {
					if Truth_(func() interface{} {
						var and__163__auto__ = ks
						_ = and__163__auto__
						if Truth_(and__163__auto__) {
							return vs
						} else {
							return and__163__auto__
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
		return Fn(max_key, func(k interface{}, x interface{}) interface{} {
			return x
		}, func(k interface{}, x interface{}, y interface{}) interface{} {
			if k.(CljsCoreIFn).X_invoke_Arity1(x).(float64) > k.(CljsCoreIFn).X_invoke_Arity1(y).(float64) {
				return x
			} else {
				return y
			}
		}, func(k_x_y_more__ ...interface{}) interface{} {
			var k = k_x_y_more__[0]
			var x = k_x_y_more__[1]
			var y = k_x_y_more__[2]
			var more = Array_seq.X_invoke_Arity1(k_x_y_more__[3:])
			_, _, _, _ = k, x, y, more
			return Reduce.X_invoke_Arity3(func(G__618 *AFn) *AFn {
				return Fn(G__618, func(p1__616_SHARP_ interface{}, p2__617_SHARP_ interface{}) interface{} {
					return max_key.X_invoke_Arity3(k, p1__616_SHARP_, p2__617_SHARP_)
				})
			}(&AFn{}), max_key.X_invoke_Arity3(k, x, y), more)
		})
	}(&AFn{})

	Min_key = func(min_key *AFn) *AFn {
		return Fn(min_key, func(k interface{}, x interface{}) interface{} {
			return x
		}, func(k interface{}, x interface{}, y interface{}) interface{} {
			if k.(CljsCoreIFn).X_invoke_Arity1(x).(float64) < k.(CljsCoreIFn).X_invoke_Arity1(y).(float64) {
				return x
			} else {
				return y
			}
		}, func(k_x_y_more__ ...interface{}) interface{} {
			var k = k_x_y_more__[0]
			var x = k_x_y_more__[1]
			var y = k_x_y_more__[2]
			var more = Array_seq.X_invoke_Arity1(k_x_y_more__[3:])
			_, _, _, _ = k, x, y, more
			return Reduce.X_invoke_Arity3(func(G__621 *AFn) *AFn {
				return Fn(G__621, func(p1__619_SHARP_ interface{}, p2__620_SHARP_ interface{}) interface{} {
					return min_key.X_invoke_Arity3(k, p1__619_SHARP_, p2__620_SHARP_)
				})
			}(&AFn{}), min_key.X_invoke_Arity3(k, x, y), more)
		})
	}(&AFn{})

	X__GT_ArrayList = func(__GT_ArrayList *AFn) *AFn {
		return Fn(__GT_ArrayList, func(arr interface{}) interface{} {
			return (&CljsCoreArrayList{arr})
		})
	}(&AFn{})

	Array_list = func(array_list *AFn) *AFn {
		return Fn(array_list, func() interface{} {
			return (&CljsCoreArrayList{[]interface{}{}})
		})
	}(&AFn{})

	Partition_all = func(partition_all *AFn) *AFn {
		return Fn(partition_all, func(n interface{}) interface{} {
			return func(G__622 *AFn) *AFn {
				return Fn(G__622, func(f1 interface{}) interface{} {
					{
						var a = Array_list.X_invoke_Arity0().(*CljsCoreArrayList)
						_ = a
						return func(G__623 *AFn, a *CljsCoreArrayList) *AFn {
							return Fn(G__623, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0()
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
												return f1.(CljsCoreIFn).X_invoke_Arity2(result, v)
											}()
										}
									}()
									_ = result___1
									return f1.(CljsCoreIFn).X_invoke_Arity1(result___1)
								}
							}, func(result interface{}, input interface{}) interface{} {
								a.Add(input)
								if n.(float64) == a.Size().(float64) {
									{
										var v = Vec.X_invoke_Arity1(a.ToArray())
										_ = v
										a.Clear()
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, v)
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
			return (&CljsCoreLazySeq{nil, func(G__624 *AFn) *AFn {
				return Fn(G__624, func() interface{} {
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
		return Fn(take_while, func(pred interface{}) interface{} {
			return func(G__625 *AFn) *AFn {
				return Fn(G__625, func(f1 interface{}) interface{} {
					return func(G__626 *AFn) *AFn {
						return Fn(G__626, func() interface{} {
							return f1.(CljsCoreIFn).X_invoke_Arity0()
						}, func(result interface{}) interface{} {
							return f1.(CljsCoreIFn).X_invoke_Arity1(result)
						}, func(result interface{}, input interface{}) interface{} {
							if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(input)) {
								return f1.(CljsCoreIFn).X_invoke_Arity2(result, input)
							} else {
								return Reduced.X_invoke_Arity1(result).(*CljsCoreReduced)
							}
						})
					}(&AFn{})
				})
			}(&AFn{})
		}, func(pred interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__627 *AFn) *AFn {
				return Fn(G__627, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(s))) {
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
		return Fn(mk_bound_fn, func(sc interface{}, test interface{}, key interface{}) interface{} {
			return func(G__628 *AFn) *AFn {
				return Fn(G__628, func(e interface{}) interface{} {
					{
						var comp = sc.(CljsCoreISorted).X_comparator_Arity1()
						_ = comp
						return test.(CljsCoreIFn).X_invoke_Arity2(comp.(CljsCoreIFn).X_invoke_Arity2(sc.(CljsCoreISorted).X_entry_key_Arity2(e), key), float64(0))
					}
				})
			}(&AFn{})
		})
	}(&AFn{})

	Subseq = func(subseq *AFn) *AFn {
		return Fn(subseq, func(sc interface{}, test interface{}, key interface{}) interface{} {
			{
				var include = Mk_bound_fn.X_invoke_Arity3(sc, test, key).(CljsCoreIFn)
				_ = include
				if Truth_(CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{X_GT_, X_GT__EQ_}, true).(*CljsCorePersistentHashSet).X_invoke_Arity1(test)) {
					{
						var temp__4222__auto__ = sc.(CljsCoreISorted).X_sorted_seq_from_Arity3(key, true)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var vec__631 = temp__4222__auto__
								var e = Nth.X_invoke_Arity3(vec__631, float64(0), nil)
								var s = vec__631
								_, _, _ = vec__631, e, s
								if Truth_(include.X_invoke_Arity1(e)) {
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
						var vec__632 = temp__4222__auto__
						var e = Nth.X_invoke_Arity3(vec__632, float64(0), nil)
						var s = vec__632
						_, _, _ = vec__632, e, s
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
		return Fn(rsubseq, func(sc interface{}, test interface{}, key interface{}) interface{} {
			{
				var include = Mk_bound_fn.X_invoke_Arity3(sc, test, key).(CljsCoreIFn)
				_ = include
				if Truth_(CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{X_LT_, X_LT__EQ_}, true).(*CljsCorePersistentHashSet).X_invoke_Arity1(test)) {
					{
						var temp__4222__auto__ = sc.(CljsCoreISorted).X_sorted_seq_from_Arity3(key, false)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var vec__635 = temp__4222__auto__
								var e = Nth.X_invoke_Arity3(vec__635, float64(0), nil)
								var s = vec__635
								_, _, _ = vec__635, e, s
								if Truth_(include.X_invoke_Arity1(e)) {
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
						var vec__636 = temp__4222__auto__
						var e = Nth.X_invoke_Arity3(vec__636, float64(0), nil)
						var s = vec__636
						_, _, _ = vec__636, e, s
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
		return Fn(__GT_Range, func(meta interface{}, start interface{}, end interface{}, step interface{}, __hash interface{}) interface{} {
			return (&CljsCoreRange{meta, start, end, step, __hash})
		})
	}(&AFn{})

	Range_ = func(range_ *AFn) *AFn {
		return Fn(range_, func() interface{} {
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
		return Fn(take_nth, func(n interface{}) interface{} {
			return func(G__637 *AFn) *AFn {
				return Fn(G__637, func(f1 interface{}) interface{} {
					{
						var ia = Atom.X_invoke_Arity1(float64(-1)).(*CljsCoreAtom)
						_ = ia
						return func(G__638 *AFn, ia *CljsCoreAtom) *AFn {
							return Fn(G__638, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0()
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result)
							}, func(result interface{}, input interface{}) interface{} {
								{
									var i = Swap_BANG_.X_invoke_Arity2(ia, Inc)
									_ = i
									if Rem.X_invoke_Arity2(i, n).(float64) == float64(0) {
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, input)
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
			return (&CljsCoreLazySeq{nil, func(G__639 *AFn) *AFn {
				return Fn(G__639, func() interface{} {
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
		return Fn(split_with, func(pred interface{}, coll interface{}) interface{} {
			return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Take_while.X_invoke_Arity2(pred, coll).(*CljsCoreLazySeq), Drop_while.X_invoke_Arity2(pred, coll).(*CljsCoreLazySeq)}, nil})
		})
	}(&AFn{})

	Partition_by = func(partition_by *AFn) *AFn {
		return Fn(partition_by, func(f interface{}) interface{} {
			return func(G__641 *AFn) *AFn {
				return Fn(G__641, func(f1 interface{}) interface{} {
					{
						var a = Array_list.X_invoke_Arity0().(*CljsCoreArrayList)
						var pa = Atom.X_invoke_Arity1((&CljsCoreKeyword{Ns: "cljs.core", Name: "none", Fqn: "cljs.core/none", X_hash: float64(926646439)})).(*CljsCoreAtom)
						_, _ = a, pa
						return func(G__642 *AFn, a *CljsCoreArrayList, pa *CljsCoreAtom) *AFn {
							return Fn(G__642, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0()
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
												return f1.(CljsCoreIFn).X_invoke_Arity2(result, v)
											}()
										}
									}()
									_ = result___1
									return f1.(CljsCoreIFn).X_invoke_Arity1(result___1)
								}
							}, func(result interface{}, input interface{}) interface{} {
								{
									var pval = Deref.X_invoke_Arity1(pa)
									var val = f.(CljsCoreIFn).X_invoke_Arity1(input)
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
												var ret = f1.(CljsCoreIFn).X_invoke_Arity2(result, v)
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
			return (&CljsCoreLazySeq{nil, func(G__643 *AFn) *AFn {
				return Fn(G__643, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								{
									var fst = First.X_invoke_Arity1(s)
									var fv = f.(CljsCoreIFn).X_invoke_Arity1(fst)
									var run = Cons.X_invoke_Arity2(fst, Take_while.X_invoke_Arity2(func(G__644 *AFn, fst interface{}, fv interface{}, s CljsCoreISeq, temp__4222__auto__ CljsCoreISeq) *AFn {
										return Fn(G__644, func(p1__640_SHARP_ interface{}) interface{} {
											return X_EQ_.Arity2IIB(fv, f.(CljsCoreIFn).X_invoke_Arity1(p1__640_SHARP_))
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
		return Fn(frequencies, func(coll interface{}) interface{} {
			return Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(func(G__645 *AFn) *AFn {
				return Fn(G__645, func(counts interface{}, x interface{}) interface{} {
					return Assoc_BANG_.X_invoke_Arity3(counts, x, (Get.X_invoke_Arity3(counts, x, float64(0)).(float64) + float64(1)))
				})
			}(&AFn{}), Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY), coll))
		})
	}(&AFn{})

	Reductions = func(reductions *AFn) *AFn {
		return Fn(reductions, func(f interface{}, coll interface{}) interface{} {
			return (&CljsCoreLazySeq{nil, func(G__646 *AFn) *AFn {
				return Fn(G__646, func() interface{} {
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
							return CljsCoreList_EMPTY.X_conj_Arity2(f.(CljsCoreIFn).X_invoke_Arity0())
						}
					}
				})
			}(&AFn{}), nil, nil})
		}, func(f interface{}, init interface{}, coll interface{}) interface{} {
			return Cons.X_invoke_Arity2(init, (&CljsCoreLazySeq{nil, func(G__647 *AFn) *AFn {
				return Fn(G__647, func() interface{} {
					{
						var temp__4222__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4222__auto__
						if Truth_(temp__4222__auto__) {
							{
								var s = temp__4222__auto__
								_ = s
								return reductions.X_invoke_Arity3(f, f.(CljsCoreIFn).X_invoke_Arity2(init, First.X_invoke_Arity1(s)), Rest.Arity1IQ(s)).(*CljsCoreCons)
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
		return Fn(juxt, func(f interface{}) interface{} {
			return func(G__658 *AFn) *AFn {
				return Fn(G__658, func() interface{} {
					return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity0()}, nil})
				}, func(x interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity1(x)}, nil})
				}, func(x interface{}, y interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity2(x, y)}, nil})
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity3(x, y, z)}, nil})
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Apply.X_invoke_Arity5(f, x, y, z, args)}, nil})
				})
			}(&AFn{})
		}, func(f interface{}, g interface{}) interface{} {
			return func(G__659 *AFn) *AFn {
				return Fn(G__659, func() interface{} {
					return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity0(), g.(CljsCoreIFn).X_invoke_Arity0()}, nil})
				}, func(x interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity1(x), g.(CljsCoreIFn).X_invoke_Arity1(x)}, nil})
				}, func(x interface{}, y interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity2(x, y), g.(CljsCoreIFn).X_invoke_Arity2(x, y)}, nil})
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity3(x, y, z), g.(CljsCoreIFn).X_invoke_Arity3(x, y, z)}, nil})
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Apply.X_invoke_Arity5(f, x, y, z, args), Apply.X_invoke_Arity5(g, x, y, z, args)}, nil})
				})
			}(&AFn{})
		}, func(f interface{}, g interface{}, h interface{}) interface{} {
			return func(G__660 *AFn) *AFn {
				return Fn(G__660, func() interface{} {
					return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity0(), g.(CljsCoreIFn).X_invoke_Arity0(), h.(CljsCoreIFn).X_invoke_Arity0()}, nil})
				}, func(x interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity1(x), g.(CljsCoreIFn).X_invoke_Arity1(x), h.(CljsCoreIFn).X_invoke_Arity1(x)}, nil})
				}, func(x interface{}, y interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity2(x, y), g.(CljsCoreIFn).X_invoke_Arity2(x, y), h.(CljsCoreIFn).X_invoke_Arity2(x, y)}, nil})
				}, func(x interface{}, y interface{}, z interface{}) interface{} {
					return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity3(x, y, z), g.(CljsCoreIFn).X_invoke_Arity3(x, y, z), h.(CljsCoreIFn).X_invoke_Arity3(x, y, z)}, nil})
				}, func(x_y_z_args__ ...interface{}) interface{} {
					var x = x_y_z_args__[0]
					var y = x_y_z_args__[1]
					var z = x_y_z_args__[2]
					var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
					_, _, _, _ = x, y, z, args
					return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{Apply.X_invoke_Arity5(f, x, y, z, args), Apply.X_invoke_Arity5(g, x, y, z, args), Apply.X_invoke_Arity5(h, x, y, z, args)}, nil})
				})
			}(&AFn{})
		}, func(f_g_h_fs__ ...interface{}) interface{} {
			var f = f_g_h_fs__[0]
			var g = f_g_h_fs__[1]
			var h = f_g_h_fs__[2]
			var fs = Array_seq.X_invoke_Arity1(f_g_h_fs__[3:])
			_, _, _, _ = f, g, h, fs
			{
				var fs___1 = List_STAR_.X_invoke_Arity4(f, g, h, fs).(*CljsCoreCons)
				_ = fs___1
				return func(G__661 *AFn, fs___1 *CljsCoreCons) *AFn {
					return Fn(G__661, func() interface{} {
						return Reduce.X_invoke_Arity3(func(G__662 *AFn, fs___1 *CljsCoreCons) *AFn {
							return Fn(G__662, func(p1__648_SHARP_ interface{}, p2__649_SHARP_ interface{}) interface{} {
								return Conj.X_invoke_Arity2(p1__648_SHARP_, p2__649_SHARP_.(CljsCoreIFn).X_invoke_Arity0())
							})
						}(&AFn{}, fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(G__663 *AFn, fs___1 *CljsCoreCons) *AFn {
							return Fn(G__663, func(p1__650_SHARP_ interface{}, p2__651_SHARP_ interface{}) interface{} {
								return Conj.X_invoke_Arity2(p1__650_SHARP_, p2__651_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x))
							})
						}(&AFn{}, fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x interface{}, y interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(G__664 *AFn, fs___1 *CljsCoreCons) *AFn {
							return Fn(G__664, func(p1__652_SHARP_ interface{}, p2__653_SHARP_ interface{}) interface{} {
								return Conj.X_invoke_Arity2(p1__652_SHARP_, p2__653_SHARP_.(CljsCoreIFn).X_invoke_Arity2(x, y))
							})
						}(&AFn{}, fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(G__665 *AFn, fs___1 *CljsCoreCons) *AFn {
							return Fn(G__665, func(p1__654_SHARP_ interface{}, p2__655_SHARP_ interface{}) interface{} {
								return Conj.X_invoke_Arity2(p1__654_SHARP_, p2__655_SHARP_.(CljsCoreIFn).X_invoke_Arity3(x, y, z))
							})
						}(&AFn{}, fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x_y_z_args__ ...interface{}) interface{} {
						var x = x_y_z_args__[0]
						var y = x_y_z_args__[1]
						var z = x_y_z_args__[2]
						var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
						_, _, _, _ = x, y, z, args
						return Reduce.X_invoke_Arity3(func(G__666 *AFn, fs___1 *CljsCoreCons) *AFn {
							return Fn(G__666, func(p1__656_SHARP_ interface{}, p2__657_SHARP_ interface{}) interface{} {
								return Conj.X_invoke_Arity2(p1__656_SHARP_, Apply.X_invoke_Arity5(p2__657_SHARP_, x, y, z, args))
							})
						}(&AFn{}, fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					})
				}(&AFn{}, fs___1)
			}
		})
	}(&AFn{})

	Dorun = func(dorun *AFn) *AFn {
		return Fn(dorun, func(coll interface{}) interface{} {
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
					var and__163__auto__ = Seq.Arity1IQ(coll)
					_ = and__163__auto__
					if Truth_(and__163__auto__) {
						return (n.(float64) > float64(0))
					} else {
						return and__163__auto__
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
		return Fn(doall, func(coll interface{}) interface{} {
			Dorun.X_invoke_Arity1(coll)
			return coll
		}, func(n interface{}, coll interface{}) interface{} {
			Dorun.X_invoke_Arity2(n, coll)
			return coll
		})
	}(&AFn{})

	Regexp_QMARK_ = func(regexp_QMARK_ *AFn) *AFn {
		return Fn(regexp_QMARK_, func(o interface{}) interface{} {
			return func() bool { _, instanceof := o.(*js.RegExp); return instanceof }()
		})
	}(&AFn{})

	Re_matches = func(re_matches *AFn) *AFn {
		return Fn(re_matches, func(re interface{}, s interface{}) interface{} {
			if reflect.ValueOf(s).Kind() == reflect.String {
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
		return Fn(re_find, func(re interface{}, s interface{}) interface{} {
			if reflect.ValueOf(s).Kind() == reflect.String {
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
		return Fn(re_seq, func(re interface{}, s interface{}) interface{} {
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
					return (&CljsCoreLazySeq{nil, func(G__667 *AFn, match_data interface{}, match_idx interface{}, match_str interface{}, post_match interface{}) *AFn {
						return Fn(G__667, func() interface{} {
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
		return Fn(re_pattern, func(s interface{}) interface{} {
			{
				var vec__669 = Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `^(?:\(\?([idmsux]*)\))?(.*)`, Flags: ``}), s)
				var ___ = Nth.X_invoke_Arity3(vec__669, float64(0), nil)
				var flags = Nth.X_invoke_Arity3(vec__669, float64(1), nil)
				var pattern = Nth.X_invoke_Arity3(vec__669, float64(2), nil)
				_, _, _, _ = vec__669, ___, flags, pattern
				return (&js.RegExp{pattern, flags})
			}
		})
	}(&AFn{})

	Write_all = func(write_all *AFn) *AFn {
		return Fn(write_all, func(writer_ss__ ...interface{}) interface{} {
			var writer = writer_ss__[0]
			var ss = Array_seq.X_invoke_Arity1(writer_ss__[1:])
			_, _ = writer, ss
			{
				var seq__676 = Seq.Arity1IQ(ss)
				var chunk__677 interface{} = nil
				var count__678 = float64(0)
				var i__679 = float64(0)
				_, _, _, _ = seq__676, chunk__677, count__678, i__679
				for {
					if i__679 < count__678 {
						{
							var s = chunk__677.(CljsCoreIIndexed).X_nth_Arity2(i__679)
							_ = s
							writer.(CljsCoreIWriter).X_write_Arity2(s)
							seq__676, chunk__677, count__678, i__679 = seq__676, chunk__677, count__678, (i__679 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto__ = Seq.Arity1IQ(seq__676)
							_ = temp__4222__auto__
							if Truth_(temp__4222__auto__) {
								{
									var seq__676___1 = temp__4222__auto__
									_ = seq__676___1
									if Chunked_seq_QMARK_.Arity1IB(seq__676___1) {
										{
											var c__942__auto__ = Chunk_first.X_invoke_Arity1(seq__676___1)
											_ = c__942__auto__
											seq__676, chunk__677, count__678, i__679 = Chunk_rest.X_invoke_Arity1(seq__676___1).(CljsCoreISeq), c__942__auto__, Count.X_invoke_Arity1(c__942__auto__).(float64), float64(0)
											continue
										}
									} else {
										{
											var s = First.X_invoke_Arity1(seq__676___1)
											_ = s
											writer.(CljsCoreIWriter).X_write_Arity2(s)
											seq__676, chunk__677, count__678, i__679 = Next.Arity1IQ(seq__676___1), nil, float64(0), float64(0)
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
		return Fn(string_print, func(x interface{}) interface{} {
			X_STAR_print_fn_STAR_.X_invoke_Arity1(x)
			return nil
		})
	}(&AFn{})

	Flush = func(flush *AFn) *AFn {
		return Fn(flush, func() interface{} {
			return nil
		})
	}(&AFn{})

	Char_escapes = func() interface{} {
		var obj681 = map[string]interface{}{`"\""`: "\\\"", `"\\"`: "\\\\", `"\b"`: "\\b", `"\f"`: "\\f", `"\n"`: "\\n", `"\r"`: "\\r", `"\t"`: "\\t"}
		_ = obj681
		return obj681
	}()

	Pr_seq_writer = func(pr_seq_writer *AFn) *AFn {
		return Fn(pr_seq_writer, func(objs interface{}, writer interface{}, opts interface{}) interface{} {
			Pr_writer.X_invoke_Arity3(First.X_invoke_Arity1(objs), writer, opts)
			{
				var seq__686 = Seq.Arity1IQ(Next.Arity1IQ(objs))
				var chunk__687 interface{} = nil
				var count__688 = float64(0)
				var i__689 = float64(0)
				_, _, _, _ = seq__686, chunk__687, count__688, i__689
				for {
					if i__689 < count__688 {
						{
							var obj = chunk__687.(CljsCoreIIndexed).X_nth_Arity2(i__689)
							_ = obj
							writer.(CljsCoreIWriter).X_write_Arity2(" ")
							Pr_writer.X_invoke_Arity3(obj, writer, opts)
							seq__686, chunk__687, count__688, i__689 = seq__686, chunk__687, count__688, (i__689 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto__ = Seq.Arity1IQ(seq__686)
							_ = temp__4222__auto__
							if Truth_(temp__4222__auto__) {
								{
									var seq__686___1 = temp__4222__auto__
									_ = seq__686___1
									if Chunked_seq_QMARK_.Arity1IB(seq__686___1) {
										{
											var c__942__auto__ = Chunk_first.X_invoke_Arity1(seq__686___1)
											_ = c__942__auto__
											seq__686, chunk__687, count__688, i__689 = Chunk_rest.X_invoke_Arity1(seq__686___1).(CljsCoreISeq), c__942__auto__, Count.X_invoke_Arity1(c__942__auto__).(float64), float64(0)
											continue
										}
									} else {
										{
											var obj = First.X_invoke_Arity1(seq__686___1)
											_ = obj
											writer.(CljsCoreIWriter).X_write_Arity2(" ")
											Pr_writer.X_invoke_Arity3(obj, writer, opts)
											seq__686, chunk__687, count__688, i__689 = Next.Arity1IQ(seq__686___1), nil, float64(0), float64(0)
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
		return Fn(pr_sb_with_opts, func(objs interface{}, opts interface{}) interface{} {
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
		return Fn(pr_str_with_opts, func(objs interface{}, opts interface{}) interface{} {
			if Empty_QMARK_.Arity1IB(objs) {
				return ""
			} else {
				return (`` + Str.X_invoke_Arity1(Pr_sb_with_opts.X_invoke_Arity2(objs, opts).(*goog_string.StringBuffer)).(string))
			}
		})
	}(&AFn{})

	Prn_str_with_opts = func(prn_str_with_opts *AFn) *AFn {
		return Fn(prn_str_with_opts, func(objs interface{}, opts interface{}) interface{} {
			if Empty_QMARK_.Arity1IB(objs) {
				return "\n"
			} else {
				{
					var sb = Pr_sb_with_opts.X_invoke_Arity2(objs, opts).(*goog_string.StringBuffer)
					_ = sb
					sb.Append('\n')
					return (`` + Str.X_invoke_Arity1(sb).(string))
				}
			}
		})
	}(&AFn{})

	Pr_with_opts = func(pr_with_opts *AFn) *AFn {
		return Fn(pr_with_opts, func(objs interface{}, opts interface{}) interface{} {
			return String_print.X_invoke_Arity1(Pr_str_with_opts.X_invoke_Arity2(objs, opts).(string))
		})
	}(&AFn{})

	Newline = func(newline *AFn) *AFn {
		return Fn(newline, func(opts interface{}) interface{} {
			String_print.X_invoke_Arity1("\n")
			if Truth_(Get.X_invoke_Arity2(opts, (&CljsCoreKeyword{Ns: nil, Name: "flush-on-newline", Fqn: "flush-on-newline", X_hash: float64(-151457939)}))) {
				return Flush.X_invoke_Arity0()
			} else {
				return nil
			}
		})
	}(&AFn{})

	Pr_str = func(pr_str *AFn) *AFn {
		return Fn(pr_str, func(objs__ ...interface{}) interface{} {
			var objs = Array_seq.X_invoke_Arity1(objs__[0:])
			_ = objs
			return Pr_str_with_opts.X_invoke_Arity2(objs, Pr_opts.X_invoke_Arity0().(CljsCoreIMap)).(string)
		})
	}(&AFn{})

	Prn_str = func(prn_str *AFn) *AFn {
		return Fn(prn_str, func(objs__ ...interface{}) interface{} {
			var objs = Array_seq.X_invoke_Arity1(objs__[0:])
			_ = objs
			return Prn_str_with_opts.X_invoke_Arity2(objs, Pr_opts.X_invoke_Arity0().(CljsCoreIMap)).(string)
		})
	}(&AFn{})

	Pr = func(pr *AFn) *AFn {
		return Fn(pr, func(objs__ ...interface{}) interface{} {
			var objs = Array_seq.X_invoke_Arity1(objs__[0:])
			_ = objs
			return Pr_with_opts.X_invoke_Arity2(objs, Pr_opts.X_invoke_Arity0().(CljsCoreIMap))
		})
	}(&AFn{})

	Print = func(cljs_core_print *AFn) *AFn {
		return Fn(cljs_core_print, func(objs__ ...interface{}) interface{} {
			var objs = Array_seq.X_invoke_Arity1(objs__[0:])
			_ = objs
			return Pr_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false))
		})
	}(&AFn{})

	Print_str = func(print_str *AFn) *AFn {
		return Fn(print_str, func(objs__ ...interface{}) interface{} {
			var objs = Array_seq.X_invoke_Arity1(objs__[0:])
			_ = objs
			return Pr_str_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false)).(string)
		})
	}(&AFn{})

	Println = func(println *AFn) *AFn {
		return Fn(println, func(objs__ ...interface{}) interface{} {
			var objs = Array_seq.X_invoke_Arity1(objs__[0:])
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
		return Fn(println_str, func(objs__ ...interface{}) interface{} {
			var objs = Array_seq.X_invoke_Arity1(objs__[0:])
			_ = objs
			return Prn_str_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false)).(string)
		})
	}(&AFn{})

	Prn = func(prn *AFn) *AFn {
		return Fn(prn, func(objs__ ...interface{}) interface{} {
			var objs = Array_seq.X_invoke_Arity1(objs__[0:])
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
		return Fn(print_map, func(m interface{}, print_one interface{}, writer interface{}, opts interface{}) interface{} {
			return Pr_sequential_writer.X_invoke_Arity7(writer, func(G__690 *AFn) *AFn {
				return Fn(G__690, func(e interface{}, w interface{}, opts___1 interface{}) interface{} {
					print_one.(CljsCoreIFn).X_invoke_Arity3(Key.X_invoke_Arity1(e), w, opts___1)
					w.(CljsCoreIWriter).X_write_Arity2(' ')
					return print_one.(CljsCoreIFn).X_invoke_Arity3(Val.X_invoke_Arity1(e), w, opts___1)
				})
			}(&AFn{}), "{", ", ", "}", opts, Seq.Arity1IQ(m))
		})
	}(&AFn{})

	Alter_meta_BANG_ = func(alter_meta_BANG_ *AFn) *AFn {
		return Fn(alter_meta_BANG_, func(iref_f_args__ ...interface{}) interface{} {
			var iref = iref_f_args__[0]
			var f = iref_f_args__[1]
			var args = Array_seq.X_invoke_Arity1(iref_f_args__[2:])
			_, _, _ = iref, f, args
			return func() interface{} {
				var return__691 = Apply.X_invoke_Arity3(f, Native_get_instance_field.X_invoke_Arity2(iref, "Meta"), args)
				Native_set_instance_field.X_invoke_Arity3(iref, "Meta", return__691)
				return return__691
			}()
		})
	}(&AFn{})

	Reset_meta_BANG_ = func(reset_meta_BANG_ *AFn) *AFn {
		return Fn(reset_meta_BANG_, func(iref interface{}, m interface{}) interface{} {
			return func() interface{} {
				var return__692 = m
				Native_set_instance_field.X_invoke_Arity3(iref, "Meta", return__692)
				return return__692
			}()
		})
	}(&AFn{})

	Add_watch = func(add_watch *AFn) *AFn {
		return Fn(add_watch, func(iref interface{}, key interface{}, f interface{}) interface{} {
			return iref.(CljsCoreIWatchable).X_add_watch_Arity3(key, f)
		})
	}(&AFn{})

	Remove_watch = func(remove_watch *AFn) *AFn {
		return Fn(remove_watch, func(iref interface{}, key interface{}) interface{} {
			return iref.(CljsCoreIWatchable).X_remove_watch_Arity2(key)
		})
	}(&AFn{})

	Gensym_counter = nil

	Gensym = func(gensym *AFn) *AFn {
		return Fn(gensym, func() interface{} {
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
		return Fn(__GT_Delay, func(f interface{}, value interface{}) interface{} {
			return (&CljsCoreDelay{f, value})
		})
	}(&AFn{})

	Delay_QMARK_ = func(delay_QMARK_ *AFn) *AFn {
		return Fn(delay_QMARK_, func(x interface{}) bool {
			return func() bool { _, instanceof := x.(*CljsCoreDelay); return instanceof }()
		})
	}(&AFn{})

	Force = func(force *AFn) *AFn {
		return Fn(force, func(x interface{}) interface{} {
			if Delay_QMARK_.Arity1IB(x) {
				return Deref.X_invoke_Arity1(x)
			} else {
				return x
			}
		})
	}(&AFn{})

	Realized_QMARK_ = func(realized_QMARK_ *AFn) *AFn {
		return Fn(realized_QMARK_, func(d interface{}) bool {
			return d.(CljsCoreIPending).X_realized_QMARK__Arity1()
		})
	}(&AFn{})

	Preserving_reduced = func(preserving_reduced *AFn) *AFn {
		return Fn(preserving_reduced, func(f1 interface{}) interface{} {
			return func(G__695 *AFn) *AFn {
				return Fn(G__695, func(p1__693_SHARP_ interface{}, p2__694_SHARP_ interface{}) interface{} {
					{
						var ret = f1.(CljsCoreIFn).X_invoke_Arity2(p1__693_SHARP_, p2__694_SHARP_)
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
		return Fn(cat, func(f1 interface{}) interface{} {
			{
				var rf1 = Preserving_reduced.X_invoke_Arity1(f1).(CljsCoreIFn)
				_ = rf1
				return func(G__696 *AFn, rf1 CljsCoreIFn) *AFn {
					return Fn(G__696, func() interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity0()
					}, func(result interface{}) interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity1(result)
					}, func(result interface{}, input interface{}) interface{} {
						return Reduce.X_invoke_Arity3(rf1, result, input)
					})
				}(&AFn{}, rf1)
			}
		})
	}(&AFn{})

	Dedupe = func(dedupe *AFn) *AFn {
		return Fn(dedupe, func() interface{} {
			return func(G__697 *AFn) *AFn {
				return Fn(G__697, func(f1 interface{}) interface{} {
					{
						var pa = Atom.X_invoke_Arity1((&CljsCoreKeyword{Ns: "cljs.core", Name: "none", Fqn: "cljs.core/none", X_hash: float64(926646439)})).(*CljsCoreAtom)
						_ = pa
						return func(G__698 *AFn, pa *CljsCoreAtom) *AFn {
							return Fn(G__698, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0()
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result)
							}, func(result interface{}, input interface{}) interface{} {
								{
									var prior = Deref.X_invoke_Arity1(pa)
									_ = prior
									Reset_BANG_.X_invoke_Arity2(pa, input)
									if X_EQ_.Arity2IIB(prior, input) {
										return result
									} else {
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, input)
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
		return Fn(random_sample, func(prob interface{}) interface{} {
			return Filter.X_invoke_Arity1(func(G__699 *AFn) *AFn {
				return Fn(G__699, func(___ interface{}) interface{} {
					return (Rand.Arity0F() < prob.(float64))
				})
			}(&AFn{})).(CljsCoreIFn)
		}, func(prob interface{}, coll interface{}) interface{} {
			return Filter.X_invoke_Arity2(func(G__700 *AFn) *AFn {
				return Fn(G__700, func(___ interface{}) interface{} {
					return (Rand.Arity0F() < prob.(float64))
				})
			}(&AFn{}), coll).(*CljsCoreLazySeq)
		})
	}(&AFn{})

	X__GT_Iteration = func(__GT_Iteration *AFn) *AFn {
		return Fn(__GT_Iteration, func(xform interface{}, coll interface{}) interface{} {
			return (&CljsCoreIteration{xform, coll})
		})
	}(&AFn{})

	Iteration = func(iteration *AFn) *AFn {
		return Fn(iteration, func(xform interface{}, coll interface{}) interface{} {
			return (&CljsCoreIteration{xform, coll})
		})
	}(&AFn{})

	Run_BANG_ = func(run_BANG_ *AFn) *AFn {
		return Fn(run_BANG_, func(proc interface{}, coll interface{}) interface{} {
			return Reduce.X_invoke_Arity3(func(G__703 *AFn) *AFn {
				return Fn(G__703, func(p1__702_SHARP_ interface{}, p2__701_SHARP_ interface{}) interface{} {
					return proc.(CljsCoreIFn).X_invoke_Arity1(p2__701_SHARP_)
				})
			}(&AFn{}), nil, coll)
		})
	}(&AFn{})

	X_clj__GT_js = func(_clj__GT_js *AFn) *AFn {
		return Fn(_clj__GT_js, func(x interface{}) interface{} {
			return x.(CljsCoreIEncodeJS).X_clj__GT_js_Arity1()
		})
	}(&AFn{})

	X_key__GT_js = func(_key__GT_js *AFn) *AFn {
		return Fn(_key__GT_js, func(x interface{}) interface{} {
			return x.(CljsCoreIEncodeJS).X_key__GT_js_Arity1()
		})
	}(&AFn{})

	Key__GT_js = func(key__GT_js *AFn) *AFn {
		return Fn(key__GT_js, func(k interface{}) interface{} {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IEncodeJS", Str: "cljs.core/IEncodeJS", X_hash: float64(-2022651465), X_meta: nil}), k)) {
				return k.(CljsCoreIEncodeJS).X_clj__GT_js_Arity1()
			} else {
				if (reflect.ValueOf(k).Kind() == reflect.String) || (reflect.ValueOf(k).Kind() == reflect.Float64) || (func() bool { _, instanceof := k.(*CljsCoreKeyword); return instanceof }()) || (func() bool { _, instanceof := k.(*CljsCoreSymbol); return instanceof }()) {
					return Clj__GT_js.X_invoke_Arity1(k)
				} else {
					return Pr_str.X_invoke_ArityVariadic(k).(string)
				}
			}
		})
	}(&AFn{})

	Clj__GT_js = func(clj__GT_js *AFn) *AFn {
		return Fn(clj__GT_js, func(x interface{}) interface{} {
			if Nil_(x) {
				return nil
			} else {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IEncodeJS", Str: "cljs.core/IEncodeJS", X_hash: float64(-2022651465), X_meta: nil}), x)) {
					return x.(CljsCoreIEncodeJS).X_clj__GT_js_Arity1()
				} else {
					if func() bool { _, instanceof := x.(*CljsCoreKeyword); return instanceof }() {
						return Name.X_invoke_Arity1(x)
					} else {
						if func() bool { _, instanceof := x.(*CljsCoreSymbol); return instanceof }() {
							return (`` + Str.X_invoke_Arity1(x).(string))
						} else {
							if Map_QMARK_.Arity1IB(x) {
								{
									var m = func() interface{} {
										var obj717 = map[string]interface{}{}
										_ = obj717
										return obj717
									}()
									_ = m
									{
										var seq__718_728 = Seq.Arity1IQ(x)
										var chunk__719_729 interface{} = nil
										var count__720_730 = float64(0)
										var i__721_731 = float64(0)
										_, _, _, _ = seq__718_728, chunk__719_729, count__720_730, i__721_731
										for {
											if i__721_731 < count__720_730 {
												{
													var vec__722_732 = chunk__719_729.(CljsCoreIIndexed).X_nth_Arity2(i__721_731)
													var k_733 = Nth.X_invoke_Arity3(vec__722_732, float64(0), nil)
													var v_734 = Nth.X_invoke_Arity3(vec__722_732, float64(1), nil)
													_, _, _ = vec__722_732, k_733, v_734
													m.([]interface{})[int(Key__GT_js.X_invoke_Arity1(k_733).(float64))] = clj__GT_js.X_invoke_Arity1(v_734)
													seq__718_728, chunk__719_729, count__720_730, i__721_731 = seq__718_728, chunk__719_729, count__720_730, (i__721_731 + float64(1))
													continue
												}
											} else {
												{
													var temp__4222__auto___735 = Seq.Arity1IQ(seq__718_728)
													_ = temp__4222__auto___735
													if Truth_(temp__4222__auto___735) {
														{
															var seq__718_736___1 = temp__4222__auto___735
															_ = seq__718_736___1
															if Chunked_seq_QMARK_.Arity1IB(seq__718_736___1) {
																{
																	var c__942__auto___737 = Chunk_first.X_invoke_Arity1(seq__718_736___1)
																	_ = c__942__auto___737
																	seq__718_728, chunk__719_729, count__720_730, i__721_731 = Chunk_rest.X_invoke_Arity1(seq__718_736___1).(CljsCoreISeq), c__942__auto___737, Count.X_invoke_Arity1(c__942__auto___737).(float64), float64(0)
																	continue
																}
															} else {
																{
																	var vec__723_738 = First.X_invoke_Arity1(seq__718_736___1)
																	var k_739 = Nth.X_invoke_Arity3(vec__723_738, float64(0), nil)
																	var v_740 = Nth.X_invoke_Arity3(vec__723_738, float64(1), nil)
																	_, _, _ = vec__723_738, k_739, v_740
																	m.([]interface{})[int(Key__GT_js.X_invoke_Arity1(k_739).(float64))] = clj__GT_js.X_invoke_Arity1(v_740)
																	seq__718_728, chunk__719_729, count__720_730, i__721_731 = Next.Arity1IQ(seq__718_736___1), nil, float64(0), float64(0)
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
									return m
								}
							} else {
								if Coll_QMARK_.Arity1IB(x) {
									{
										var arr = []interface{}{}
										_ = arr
										{
											var seq__724_741 = Seq.Arity1IQ(Map_.X_invoke_Arity2(clj__GT_js, x).(*CljsCoreLazySeq))
											var chunk__725_742 interface{} = nil
											var count__726_743 = float64(0)
											var i__727_744 = float64(0)
											_, _, _, _ = seq__724_741, chunk__725_742, count__726_743, i__727_744
											for {
												if i__727_744 < count__726_743 {
													{
														var x_745___1 = chunk__725_742.(CljsCoreIIndexed).X_nth_Arity2(i__727_744)
														_ = x_745___1
														js.JSArray_(&arr).Push(x_745___1)
														seq__724_741, chunk__725_742, count__726_743, i__727_744 = seq__724_741, chunk__725_742, count__726_743, (i__727_744 + float64(1))
														continue
													}
												} else {
													{
														var temp__4222__auto___746 = Seq.Arity1IQ(seq__724_741)
														_ = temp__4222__auto___746
														if Truth_(temp__4222__auto___746) {
															{
																var seq__724_747___1 = temp__4222__auto___746
																_ = seq__724_747___1
																if Chunked_seq_QMARK_.Arity1IB(seq__724_747___1) {
																	{
																		var c__942__auto___748 = Chunk_first.X_invoke_Arity1(seq__724_747___1)
																		_ = c__942__auto___748
																		seq__724_741, chunk__725_742, count__726_743, i__727_744 = Chunk_rest.X_invoke_Arity1(seq__724_747___1).(CljsCoreISeq), c__942__auto___748, Count.X_invoke_Arity1(c__942__auto___748).(float64), float64(0)
																		continue
																	}
																} else {
																	{
																		var x_749___1 = First.X_invoke_Arity1(seq__724_747___1)
																		_ = x_749___1
																		js.JSArray_(&arr).Push(x_749___1)
																		seq__724_741, chunk__725_742, count__726_743, i__727_744 = Next.Arity1IQ(seq__724_747___1), nil, float64(0), float64(0)
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
										return arr
									}
								} else {
									return x

								}
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	X_js__GT_clj = func(_js__GT_clj *AFn) *AFn {
		return Fn(_js__GT_clj, func(x interface{}, options interface{}) interface{} {
			return x.(CljsCoreIEncodeClojure).X_js__GT_clj_Arity2(options)
		})
	}(&AFn{})

	Memoize = func(memoize *AFn) *AFn {
		return Fn(memoize, func(f interface{}) interface{} {
			{
				var mem = Atom.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY).(*CljsCoreAtom)
				_ = mem
				return func(G__768 *AFn, mem *CljsCoreAtom) *AFn {
					return Fn(G__768, func(args__ ...interface{}) interface{} {
						var args = Array_seq.X_invoke_Arity1(args__[0:])
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
		return Fn(trampoline, func(f interface{}) interface{} {
			for {
				{
					var ret = f.(CljsCoreIFn).X_invoke_Arity0()
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
			var args = Array_seq.X_invoke_Arity1(f_args__[1:])
			_, _ = f, args
			return trampoline.X_invoke_Arity1(func(G__769 *AFn) *AFn {
				return Fn(G__769, func() interface{} {
					return Apply.X_invoke_Arity2(f, args)
				})
			}(&AFn{}))
		})
	}(&AFn{})

	Rand = func(rand *AFn) *AFn {
		return Fn(rand, func() interface{} {
			return rand.X_invoke_Arity1(float64(1)).(float64)
		}, func(n interface{}) interface{} {
			return (Native_invoke_func.X_invoke_Arity2(Math.Random, []interface{}{}).(float64) * n.(float64))
		})
	}(&AFn{})

	Rand_int = func(rand_int *AFn) *AFn {
		return Fn(rand_int, func(n interface{}) interface{} {
			return Native_invoke_func.X_invoke_Arity2(Math.Floor, []interface{}{(Native_invoke_func.X_invoke_Arity2(Math.Random, []interface{}{}).(float64) * n.(float64))})
		})
	}(&AFn{})

	Rand_nth = func(rand_nth *AFn) *AFn {
		return Fn(rand_nth, func(coll interface{}) interface{} {
			return Nth.X_invoke_Arity2(coll, Rand_int.X_invoke_Arity1(Count.X_invoke_Arity1(coll).(float64)))
		})
	}(&AFn{})

	Group_by = func(group_by *AFn) *AFn {
		return Fn(group_by, func(f interface{}, coll interface{}) interface{} {
			return Reduce.X_invoke_Arity3(func(G__770 *AFn) *AFn {
				return Fn(G__770, func(ret interface{}, x interface{}) interface{} {
					{
						var k = f.(CljsCoreIFn).X_invoke_Arity1(x)
						_ = k
						return Assoc.X_invoke_Arity3(ret, k, Conj.X_invoke_Arity2(Get.X_invoke_Arity3(ret, k, CljsCorePersistentVector_EMPTY), x))
					}
				})
			}(&AFn{}), CljsCorePersistentArrayMap_EMPTY, coll)
		})
	}(&AFn{})

	Make_hierarchy = func(make_hierarchy *AFn) *AFn {
		return Fn(make_hierarchy, func() interface{} {
			return (&CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}), CljsCorePersistentArrayMap_EMPTY, (&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}), CljsCorePersistentArrayMap_EMPTY, (&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}), CljsCorePersistentArrayMap_EMPTY}, nil})
		})
	}(&AFn{})

	X_global_hierarchy = nil

	Get_global_hierarchy = func(get_global_hierarchy *AFn) *AFn {
		return Fn(get_global_hierarchy, func() interface{} {
			if Nil_(X_global_hierarchy) {
				X_global_hierarchy = Atom.X_invoke_Arity1(Make_hierarchy.X_invoke_Arity0().(CljsCoreIMap)).(*CljsCoreAtom)

			} else {
			}
			return X_global_hierarchy
		})
	}(&AFn{})

	Swap_global_hierarchy_BANG_ = func(swap_global_hierarchy_BANG_ *AFn) *AFn {
		return Fn(swap_global_hierarchy_BANG_, func(f_args__ ...interface{}) interface{} {
			var f = f_args__[0]
			var args = Array_seq.X_invoke_Arity1(f_args__[1:])
			_, _ = f, args
			return Apply.X_invoke_Arity4(Swap_BANG_, Get_global_hierarchy.X_invoke_Arity0(), f, args)
		})
	}(&AFn{})

	Isa_QMARK_ = func(isa_QMARK_ *AFn) *AFn {
		return Fn(isa_QMARK_, func(child interface{}, parent interface{}) bool {
			return isa_QMARK_.Arity3IIIB(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0()), child, parent)
		}, func(h interface{}, child interface{}, parent interface{}) bool {
			{
				var or__175__auto__ = X_EQ_.Arity2IIB(child, parent)
				_ = or__175__auto__
				if Truth_(or__175__auto__) {
					return or__175__auto__
				} else {
					{
						var or__175__auto_____1 = Contains_QMARK_.Arity2IIB((&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h).(CljsCoreIFn).X_invoke_Arity1(child), parent)
						_ = or__175__auto_____1
						if Truth_(or__175__auto_____1) {
							return or__175__auto_____1
						} else {
							{
								var and__163__auto__ = Vector_QMARK_.Arity1IB(parent)
								_ = and__163__auto__
								if Truth_(and__163__auto__) {
									{
										var and__163__auto_____1 = Vector_QMARK_.Arity1IB(child)
										_ = and__163__auto_____1
										if Truth_(and__163__auto_____1) {
											{
												var and__163__auto_____2 = (Count.X_invoke_Arity1(parent).(float64) == Count.X_invoke_Arity1(child).(float64))
												_ = and__163__auto_____2
												if Truth_(and__163__auto_____2) {
													{
														var ret = true
														var i = float64(0)
														_, _ = ret, i
														for {
															if (!(ret)) || (i == Count.X_invoke_Arity1(parent).(float64)) {
																return ret
															} else {
																ret, i = isa_QMARK_.Arity3IIIB(h, child.(CljsCoreIFn).X_invoke_Arity1(i), parent.(CljsCoreIFn).X_invoke_Arity1(i)), (i + float64(1))
																continue
															}
														}
													}
												} else {
													return and__163__auto_____2
												}
											}
										} else {
											return and__163__auto_____1
										}
									}
								} else {
									return and__163__auto__
								}
							}
						}
					}
				}
			}
		})
	}(&AFn{})

	Parents = func(parents *AFn) *AFn {
		return Fn(parents, func(tag interface{}) interface{} {
			return parents.X_invoke_Arity2(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0()), tag)
		}, func(h interface{}, tag interface{}) interface{} {
			return Not_empty.X_invoke_Arity1(Get.X_invoke_Arity2((&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h), tag))
		})
	}(&AFn{})

	Ancestors = func(ancestors *AFn) *AFn {
		return Fn(ancestors, func(tag interface{}) interface{} {
			return ancestors.X_invoke_Arity2(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0()), tag)
		}, func(h interface{}, tag interface{}) interface{} {
			return Not_empty.X_invoke_Arity1(Get.X_invoke_Arity2((&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h), tag))
		})
	}(&AFn{})

	Descendants = func(descendants *AFn) *AFn {
		return Fn(descendants, func(tag interface{}) interface{} {
			return descendants.X_invoke_Arity2(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0()), tag)
		}, func(h interface{}, tag interface{}) interface{} {
			return Not_empty.X_invoke_Arity1(Get.X_invoke_Arity2((&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}).X_invoke_Arity1(h), tag))
		})
	}(&AFn{})

	Derive = func(derive *AFn) *AFn {
		return Fn(derive, func(tag interface{}, parent interface{}) interface{} {
			if Truth_(Namespace.X_invoke_Arity1(parent)) {
			} else {
				panic((&js.Error{("Assert failed: " + Str.X_invoke_Arity1(Pr_str.X_invoke_ArityVariadic(List.X_invoke_ArityVariadic((&CljsCoreSymbol{Ns: nil, Name: "namespace", Str: "namespace", X_hash: float64(1263021155), X_meta: nil}), (&CljsCoreSymbol{Ns: nil, Name: "parent", Str: "parent", X_hash: float64(761652748), X_meta: nil})).(*CljsCoreList)).(string)).(string))}))
			}
			Swap_global_hierarchy_BANG_.X_invoke_ArityVariadic(derive, tag, parent)
			return nil
		}, func(h interface{}, tag interface{}, parent interface{}) interface{} {
			if Not_EQ_.Arity2IIB(tag, parent) {
			} else {
				panic((&js.Error{("Assert failed: " + Str.X_invoke_Arity1(Pr_str.X_invoke_ArityVariadic(List.X_invoke_ArityVariadic((&CljsCoreSymbol{Ns: nil, Name: "not=", Str: "not=", X_hash: float64(1466536204), X_meta: nil}), (&CljsCoreSymbol{Ns: nil, Name: "tag", Str: "tag", X_hash: float64(350170304), X_meta: nil}), (&CljsCoreSymbol{Ns: nil, Name: "parent", Str: "parent", X_hash: float64(761652748), X_meta: nil})).(*CljsCoreList)).(string)).(string))}))
			}
			{
				var tp = (&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h)
				var td = (&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}).X_invoke_Arity1(h)
				var ta = (&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h)
				var tf = func(G__771 *AFn, tp interface{}, td interface{}, ta interface{}) *AFn {
					return Fn(G__771, func(m interface{}, source interface{}, sources interface{}, target interface{}, targets interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(G__772 *AFn, tp interface{}, td interface{}, ta interface{}) *AFn {
							return Fn(G__772, func(ret interface{}, k interface{}) interface{} {
								return Assoc.X_invoke_Arity3(ret, k, Reduce.X_invoke_Arity3(Conj, Get.X_invoke_Arity3(targets, k, CljsCorePersistentHashSet_EMPTY), Cons.X_invoke_Arity2(target, targets.(CljsCoreIFn).X_invoke_Arity1(target)).(*CljsCoreCons)))
							})
						}(&AFn{}, tp, td, ta), m, Cons.X_invoke_Arity2(source, sources.(CljsCoreIFn).X_invoke_Arity1(source)).(*CljsCoreCons))
					})
				}(&AFn{}, tp, td, ta)
				_, _, _, _ = tp, td, ta, tf
				{
					var or__175__auto__ = func() interface{} {
						if Contains_QMARK_.Arity2IIB(tp.(CljsCoreIFn).X_invoke_Arity1(tag), parent) {
							return nil
						} else {
							return func() CljsCoreIMap {
								if Contains_QMARK_.Arity2IIB(ta.(CljsCoreIFn).X_invoke_Arity1(tag), parent) {
									panic((&js.Error{(`` + Str.X_invoke_Arity1(tag).(string) + "already has" + Str.X_invoke_Arity1(parent).(string) + "as ancestor")}))
								} else {
								}
								if Contains_QMARK_.Arity2IIB(ta.(CljsCoreIFn).X_invoke_Arity1(parent), tag) {
									panic((&js.Error{("Cyclic derivation:" + Str.X_invoke_Arity1(parent).(string) + "has" + Str.X_invoke_Arity1(tag).(string) + "as ancestor")}))
								} else {
								}
								return (&CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}), Assoc.X_invoke_Arity3((&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h), tag, Conj.X_invoke_Arity2(Get.X_invoke_Arity3(tp, tag, CljsCorePersistentHashSet_EMPTY), parent)), (&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}), tf.X_invoke_Arity5((&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h), tag, td, parent, ta), (&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}), tf.X_invoke_Arity5((&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}).X_invoke_Arity1(h), parent, ta, tag, td)}, nil})
							}()
						}
					}()
					_ = or__175__auto__
					if Truth_(or__175__auto__) {
						return or__175__auto__
					} else {
						return h
					}
				}
			}
		})
	}(&AFn{})

	Underive = func(underive *AFn) *AFn {
		return Fn(underive, func(tag interface{}, parent interface{}) interface{} {
			Swap_global_hierarchy_BANG_.X_invoke_ArityVariadic(underive, tag, parent)
			return nil
		}, func(h interface{}, tag interface{}, parent interface{}) interface{} {
			{
				var parentMap = (&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h)
				var childsParents = func() interface{} {
					if Truth_(parentMap.(CljsCoreIFn).X_invoke_Arity1(tag)) {
						return Disj.X_invoke_Arity2(parentMap.(CljsCoreIFn).X_invoke_Arity1(tag), parent)
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
				var deriv_seq = Flatten.X_invoke_Arity1(Map_.X_invoke_Arity2(func(G__776 *AFn, parentMap interface{}, childsParents interface{}, newParents interface{}) *AFn {
					return Fn(G__776, func(p1__773_SHARP_ interface{}) interface{} {
						return Cons.X_invoke_Arity2(First.X_invoke_Arity1(p1__773_SHARP_), Interpose.X_invoke_Arity2(First.X_invoke_Arity1(p1__773_SHARP_), Second.X_invoke_Arity1(p1__773_SHARP_)).(*CljsCoreLazySeq)).(*CljsCoreCons)
					})
				}(&AFn{}, parentMap, childsParents, newParents), Seq.Arity1IQ(newParents)).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
				_, _, _, _ = parentMap, childsParents, newParents, deriv_seq
				if Contains_QMARK_.Arity2IIB(parentMap.(CljsCoreIFn).X_invoke_Arity1(tag), parent) {
					return Reduce.X_invoke_Arity3(func(G__777 *AFn, parentMap interface{}, childsParents interface{}, newParents interface{}, deriv_seq *CljsCoreLazySeq) *AFn {
						return Fn(G__777, func(p1__774_SHARP_ interface{}, p2__775_SHARP_ interface{}) interface{} {
							return Apply.X_invoke_Arity3(Derive, p1__774_SHARP_, p2__775_SHARP_)
						})
					}(&AFn{}, parentMap, childsParents, newParents, deriv_seq), Make_hierarchy.X_invoke_Arity0().(CljsCoreIMap), Partition.X_invoke_Arity2(float64(2), deriv_seq).(*CljsCoreLazySeq))
				} else {
					return h
				}
			}
		})
	}(&AFn{})

	Reset_cache = func(reset_cache *AFn) *AFn {
		return Fn(reset_cache, func(method_cache interface{}, method_table interface{}, cached_hierarchy interface{}, hierarchy interface{}) interface{} {
			Swap_BANG_.X_invoke_Arity2(method_cache, func(G__778 *AFn) *AFn {
				return Fn(G__778, func(___ interface{}) interface{} {
					return Deref.X_invoke_Arity1(method_table)
				})
			}(&AFn{}))
			return Swap_BANG_.X_invoke_Arity2(cached_hierarchy, func(G__779 *AFn) *AFn {
				return Fn(G__779, func(___ interface{}) interface{} {
					return Deref.X_invoke_Arity1(hierarchy)
				})
			}(&AFn{}))
		})
	}(&AFn{})

	Prefers_STAR_ = func(prefers_STAR_ *AFn) *AFn {
		return Fn(prefers_STAR_, func(x interface{}, y interface{}, prefer_table interface{}) interface{} {
			{
				var xprefs = Deref.X_invoke_Arity1(prefer_table).(CljsCoreIFn).X_invoke_Arity1(x)
				_ = xprefs
				{
					var or__175__auto__ = func() interface{} {
						if Truth_(func() interface{} {
							var and__163__auto__ = xprefs
							_ = and__163__auto__
							if Truth_(and__163__auto__) {
								return xprefs.(CljsCoreIFn).X_invoke_Arity1(y)
							} else {
								return and__163__auto__
							}
						}()) {
							return true
						} else {
							return nil
						}
					}()
					_ = or__175__auto__
					if Truth_(or__175__auto__) {
						return or__175__auto__
					} else {
						{
							var or__175__auto_____1 interface{} = func() interface{} {
								var ps = Parents.X_invoke_Arity1(y)
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
							_ = or__175__auto_____1
							if Truth_(or__175__auto_____1) {
								return or__175__auto_____1
							} else {
								{
									var or__175__auto_____2 interface{} = func() interface{} {
										var ps = Parents.X_invoke_Arity1(x)
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
									_ = or__175__auto_____2
									if Truth_(or__175__auto_____2) {
										return or__175__auto_____2
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
		return Fn(dominates, func(x interface{}, y interface{}, prefer_table interface{}) interface{} {
			{
				var or__175__auto__ = Prefers_STAR_.X_invoke_Arity3(x, y, prefer_table)
				_ = or__175__auto__
				if Truth_(or__175__auto__) {
					return or__175__auto__
				} else {
					return Isa_QMARK_.Arity2IIB(x, y)
				}
			}
		})
	}(&AFn{})

	Find_and_cache_best_method = func(find_and_cache_best_method *AFn) *AFn {
		return Fn(find_and_cache_best_method, func(name interface{}, dispatch_val interface{}, hierarchy interface{}, method_table interface{}, prefer_table interface{}, method_cache interface{}, cached_hierarchy interface{}) interface{} {
			{
				var best_entry = Reduce.X_invoke_Arity3(func(G__784 *AFn) *AFn {
					return Fn(G__784, func(be interface{}, p__782 interface{}) interface{} {
						{
							var vec__783 = p__782
							var k = Nth.X_invoke_Arity3(vec__783, float64(0), nil)
							var ___ = Nth.X_invoke_Arity3(vec__783, float64(1), nil)
							var e = vec__783
							_, _, _, _ = vec__783, k, ___, e
							if Isa_QMARK_.Arity3IIIB(Deref.X_invoke_Arity1(hierarchy), dispatch_val, k) {
								{
									var be2 = func() interface{} {
										if Truth_(func() interface{} {
											var or__175__auto__ = Nil_(be)
											_ = or__175__auto__
											if Truth_(or__175__auto__) {
												return or__175__auto__
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
		return Fn(_reset, func(mf interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_reset_Arity1()
		})
	}(&AFn{})

	X_add_method = func(_add_method *AFn) *AFn {
		return Fn(_add_method, func(mf interface{}, dispatch_val interface{}, method interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_add_method_Arity3(dispatch_val, method)
		})
	}(&AFn{})

	X_remove_method = func(_remove_method *AFn) *AFn {
		return Fn(_remove_method, func(mf interface{}, dispatch_val interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_remove_method_Arity2(dispatch_val)
		})
	}(&AFn{})

	X_prefer_method = func(_prefer_method *AFn) *AFn {
		return Fn(_prefer_method, func(mf interface{}, dispatch_val interface{}, dispatch_val_y interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_prefer_method_Arity3(dispatch_val, dispatch_val_y)
		})
	}(&AFn{})

	X_get_method = func(_get_method *AFn) *AFn {
		return Fn(_get_method, func(mf interface{}, dispatch_val interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_get_method_Arity2(dispatch_val)
		})
	}(&AFn{})

	X_methods = func(_methods *AFn) *AFn {
		return Fn(_methods, func(mf interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_methods_Arity1()
		})
	}(&AFn{})

	X_prefers = func(_prefers *AFn) *AFn {
		return Fn(_prefers, func(mf interface{}) interface{} {
			return mf.(CljsCoreIMultiFn).X_prefers_Arity1()
		})
	}(&AFn{})

	Throw_no_method_error = func(throw_no_method_error *AFn) *AFn {
		return Fn(throw_no_method_error, func(name interface{}, dispatch_val interface{}) interface{} {
			panic((&js.Error{("No method in multimethod '" + Str.X_invoke_Arity1(name).(string) + "' for dispatch value: " + Str.X_invoke_Arity1(dispatch_val).(string))}))
		})
	}(&AFn{})

	X__GT_MultiFn = func(__GT_MultiFn *AFn) *AFn {
		return Fn(__GT_MultiFn, func(name interface{}, dispatch_fn interface{}, default_dispatch_val interface{}, hierarchy interface{}, method_table interface{}, prefer_table interface{}, method_cache interface{}, cached_hierarchy interface{}) interface{} {
			return (&CljsCoreMultiFn{name, dispatch_fn, default_dispatch_val, hierarchy, method_table, prefer_table, method_cache, cached_hierarchy})
		})
	}(&AFn{})

	Remove_all_methods = func(remove_all_methods *AFn) *AFn {
		return Fn(remove_all_methods, func(multifn interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_reset_Arity1()
		})
	}(&AFn{})

	Remove_method = func(remove_method *AFn) *AFn {
		return Fn(remove_method, func(multifn interface{}, dispatch_val interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_remove_method_Arity2(dispatch_val)
		})
	}(&AFn{})

	Prefer_method = func(prefer_method *AFn) *AFn {
		return Fn(prefer_method, func(multifn interface{}, dispatch_val_x interface{}, dispatch_val_y interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_prefer_method_Arity3(dispatch_val_x, dispatch_val_y)
		})
	}(&AFn{})

	Methods = func(methods *AFn) *AFn {
		return Fn(methods, func(multifn interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_methods_Arity1()
		})
	}(&AFn{})

	Get_method = func(get_method *AFn) *AFn {
		return Fn(get_method, func(multifn interface{}, dispatch_val interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_get_method_Arity2(dispatch_val)
		})
	}(&AFn{})

	Prefers = func(prefers *AFn) *AFn {
		return Fn(prefers, func(multifn interface{}) interface{} {
			return multifn.(CljsCoreIMultiFn).X_prefers_Arity1()
		})
	}(&AFn{})

	X__GT_UUID = func(__GT_UUID *AFn) *AFn {
		return Fn(__GT_UUID, func(uuid interface{}) interface{} {
			return (&CljsCoreUUID{uuid})
		})
	}(&AFn{})

	X__GT_ExceptionInfo = func(__GT_ExceptionInfo *AFn) *AFn {
		return Fn(__GT_ExceptionInfo, func(message interface{}, data interface{}, cause interface{}) interface{} {
			return (&CljsCoreExceptionInfo{message, data, cause})
		})
	}(&AFn{})

	Ex_info = func(ex_info *AFn) *AFn {
		return Fn(ex_info, func(msg interface{}, map_ interface{}) interface{} {
			return (&CljsCoreExceptionInfo{msg, map_, nil})
		}, func(msg interface{}, map_ interface{}, cause interface{}) interface{} {
			return (&CljsCoreExceptionInfo{msg, map_, cause})
		})
	}(&AFn{})

	Ex_data = func(ex_data *AFn) *AFn {
		return Fn(ex_data, func(ex interface{}) interface{} {
			if func() bool { _, instanceof := ex.(*CljsCoreExceptionInfo); return instanceof }() {
				return Native_get_instance_field.X_invoke_Arity2(ex, "Data")
			} else {
				return nil
			}
		})
	}(&AFn{})

	Ex_message = func(ex_message *AFn) *AFn {
		return Fn(ex_message, func(ex interface{}) interface{} {
			if func() bool { _, instanceof := ex.(*js.Error); return instanceof }() {
				return Native_get_instance_field.X_invoke_Arity2(ex, "Message")
			} else {
				return nil
			}
		})
	}(&AFn{})

	Ex_cause = func(ex_cause *AFn) *AFn {
		return Fn(ex_cause, func(ex interface{}) interface{} {
			if func() bool { _, instanceof := ex.(*CljsCoreExceptionInfo); return instanceof }() {
				return Native_get_instance_field.X_invoke_Arity2(ex, "Cause")
			} else {
				return nil
			}
		})
	}(&AFn{})

	Comparator = func(comparator *AFn) *AFn {
		return Fn(comparator, func(pred interface{}) interface{} {
			return func(G__785 *AFn) *AFn {
				return Fn(G__785, func(x interface{}, y interface{}) interface{} {
					if Truth_(pred.(CljsCoreIFn).X_invoke_Arity2(x, y)) {
						return float64(-1)
					} else {
						if Truth_(pred.(CljsCoreIFn).X_invoke_Arity2(y, x)) {
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
		return Fn(special_symbol_QMARK_, func(x interface{}) bool {
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

func init() {
	RegisterProtocol_("cljs.core/Fn", (*CljsCoreFn)(nil))
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
	X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{}
	X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{}
	X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IFn", (*CljsCoreIFn)(nil))
}

var X_invoke *AFn

type CljsCoreICloneable interface {
	CljsCoreICloneable__()
	X_clone_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ICloneable", (*CljsCoreICloneable)(nil))
}

var X_clone *AFn

type CljsCoreICounted interface {
	CljsCoreICounted__()
	X_count_Arity1() float64
}

func init() {
	RegisterProtocol_("cljs.core/ICounted", (*CljsCoreICounted)(nil))
}

var X_count *AFn

type CljsCoreIEmptyableCollection interface {
	CljsCoreIEmptyableCollection__()
	X_empty_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IEmptyableCollection", (*CljsCoreIEmptyableCollection)(nil))
}

var X_empty *AFn

type CljsCoreICollection interface {
	CljsCoreICollection__()
	X_conj_Arity2(o interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ICollection", (*CljsCoreICollection)(nil))
}

var X_conj *AFn

type CljsCoreIIndexed interface {
	CljsCoreIIndexed__()
	X_nth_Arity2(n interface{}) interface{}
	X_nth_Arity3(n interface{}, not_found interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IIndexed", (*CljsCoreIIndexed)(nil))
}

var X_nth *AFn

type CljsCoreASeq interface {
	CljsCoreASeq__()
}

func init() {
	RegisterProtocol_("cljs.core/ASeq", (*CljsCoreASeq)(nil))
}

type CljsCoreISeq interface {
	CljsCoreISeq__()
	X_first_Arity1() interface{}
	X_rest_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ISeq", (*CljsCoreISeq)(nil))
}

var X_first *AFn

var X_rest *AFn

type CljsCoreINext interface {
	CljsCoreINext__()
	X_next_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/INext", (*CljsCoreINext)(nil))
}

var X_next *AFn

type CljsCoreILookup interface {
	CljsCoreILookup__()
	X_lookup_Arity2(k interface{}) interface{}
	X_lookup_Arity3(k interface{}, not_found interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ILookup", (*CljsCoreILookup)(nil))
}

var X_lookup *AFn

type CljsCoreIAssociative interface {
	CljsCoreIAssociative__()
	X_contains_key_QMARK__Arity2(k interface{}) bool
	X_assoc_Arity3(k interface{}, v interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IAssociative", (*CljsCoreIAssociative)(nil))
}

var X_contains_key_QMARK_ *AFn

var X_assoc *AFn

type CljsCoreIMap interface {
	CljsCoreIMap__()
	X_dissoc_Arity2(k interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IMap", (*CljsCoreIMap)(nil))
}

var X_dissoc *AFn

type CljsCoreIMapEntry interface {
	CljsCoreIMapEntry__()
	X_key_Arity1() interface{}
	X_val_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IMapEntry", (*CljsCoreIMapEntry)(nil))
}

var X_key *AFn

var X_val *AFn

type CljsCoreISet interface {
	CljsCoreISet__()
	X_disjoin_Arity2(v interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ISet", (*CljsCoreISet)(nil))
}

var X_disjoin *AFn

type CljsCoreIStack interface {
	CljsCoreIStack__()
	X_peek_Arity1() interface{}
	X_pop_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IStack", (*CljsCoreIStack)(nil))
}

var X_peek *AFn

var X_pop *AFn

type CljsCoreIVector interface {
	CljsCoreIVector__()
	X_assoc_n_Arity3(n interface{}, val interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IVector", (*CljsCoreIVector)(nil))
}

var X_assoc_n *AFn

type CljsCoreIDeref interface {
	CljsCoreIDeref__()
	X_deref_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IDeref", (*CljsCoreIDeref)(nil))
}

var X_deref *AFn

type CljsCoreIDerefWithTimeout interface {
	CljsCoreIDerefWithTimeout__()
	X_deref_with_timeout_Arity3(msec interface{}, timeout_val interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IDerefWithTimeout", (*CljsCoreIDerefWithTimeout)(nil))
}

var X_deref_with_timeout *AFn

type CljsCoreIMeta interface {
	CljsCoreIMeta__()
	X_meta_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IMeta", (*CljsCoreIMeta)(nil))
}

var X_meta *AFn

type CljsCoreIWithMeta interface {
	CljsCoreIWithMeta__()
	X_with_meta_Arity2(meta interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IWithMeta", (*CljsCoreIWithMeta)(nil))
}

var X_with_meta *AFn

type CljsCoreIReduce interface {
	CljsCoreIReduce__()
	X_reduce_Arity2(f interface{}) interface{}
	X_reduce_Arity3(f interface{}, start interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IReduce", (*CljsCoreIReduce)(nil))
}

var X_reduce *AFn

type CljsCoreIKVReduce interface {
	CljsCoreIKVReduce__()
	X_kv_reduce_Arity3(f interface{}, init interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IKVReduce", (*CljsCoreIKVReduce)(nil))
}

var X_kv_reduce *AFn

type CljsCoreIEquiv interface {
	CljsCoreIEquiv__()
	X_equiv_Arity2(other interface{}) bool
}

func init() {
	RegisterProtocol_("cljs.core/IEquiv", (*CljsCoreIEquiv)(nil))
}

var X_equiv *AFn

type CljsCoreIHash interface {
	CljsCoreIHash__()
	X_hash_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IHash", (*CljsCoreIHash)(nil))
}

var X_hash *AFn

type CljsCoreISeqable interface {
	CljsCoreISeqable__()
	X_seq_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ISeqable", (*CljsCoreISeqable)(nil))
}

var X_seq *AFn

type CljsCoreISequential interface {
	CljsCoreISequential__()
}

func init() {
	RegisterProtocol_("cljs.core/ISequential", (*CljsCoreISequential)(nil))
}

type CljsCoreIList interface {
	CljsCoreIList__()
}

func init() {
	RegisterProtocol_("cljs.core/IList", (*CljsCoreIList)(nil))
}

type CljsCoreIRecord interface {
	CljsCoreIRecord__()
}

func init() {
	RegisterProtocol_("cljs.core/IRecord", (*CljsCoreIRecord)(nil))
}

type CljsCoreIReversible interface {
	CljsCoreIReversible__()
	X_rseq_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IReversible", (*CljsCoreIReversible)(nil))
}

var X_rseq *AFn

type CljsCoreISorted interface {
	CljsCoreISorted__()
	X_sorted_seq_Arity2(ascending_QMARK_ interface{}) interface{}
	X_sorted_seq_from_Arity3(k interface{}, ascending_QMARK_ interface{}) interface{}
	X_entry_key_Arity2(entry interface{}) interface{}
	X_comparator_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ISorted", (*CljsCoreISorted)(nil))
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

func init() {
	RegisterProtocol_("cljs.core/IWriter", (*CljsCoreIWriter)(nil))
}

var X_write *AFn

var X_flush *AFn

type CljsCoreIPrintWithWriter interface {
	CljsCoreIPrintWithWriter__()
	X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IPrintWithWriter", (*CljsCoreIPrintWithWriter)(nil))
}

var X_pr_writer *AFn

type CljsCoreIPending interface {
	CljsCoreIPending__()
	X_realized_QMARK__Arity1() bool
}

func init() {
	RegisterProtocol_("cljs.core/IPending", (*CljsCoreIPending)(nil))
}

var X_realized_QMARK_ *AFn

type CljsCoreIWatchable interface {
	CljsCoreIWatchable__()
	X_notify_watches_Arity3(oldval interface{}, newval interface{}) interface{}
	X_add_watch_Arity3(key interface{}, f interface{}) interface{}
	X_remove_watch_Arity2(key interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IWatchable", (*CljsCoreIWatchable)(nil))
}

var X_notify_watches *AFn

var X_add_watch *AFn

var X_remove_watch *AFn

type CljsCoreIEditableCollection interface {
	CljsCoreIEditableCollection__()
	X_as_transient_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IEditableCollection", (*CljsCoreIEditableCollection)(nil))
}

var X_as_transient *AFn

type CljsCoreITransientCollection interface {
	CljsCoreITransientCollection__()
	X_conj_BANG__Arity2(val interface{}) interface{}
	X_persistent_BANG__Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ITransientCollection", (*CljsCoreITransientCollection)(nil))
}

var X_conj_BANG_ *AFn

var X_persistent_BANG_ *AFn

type CljsCoreITransientAssociative interface {
	CljsCoreITransientAssociative__()
	X_assoc_BANG__Arity3(key interface{}, val interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ITransientAssociative", (*CljsCoreITransientAssociative)(nil))
}

var X_assoc_BANG_ *AFn

type CljsCoreITransientMap interface {
	CljsCoreITransientMap__()
	X_dissoc_BANG__Arity2(key interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ITransientMap", (*CljsCoreITransientMap)(nil))
}

var X_dissoc_BANG_ *AFn

type CljsCoreITransientVector interface {
	CljsCoreITransientVector__()
	X_assoc_n_BANG__Arity3(n interface{}, val interface{}) interface{}
	X_pop_BANG__Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ITransientVector", (*CljsCoreITransientVector)(nil))
}

var X_assoc_n_BANG_ *AFn

var X_pop_BANG_ *AFn

type CljsCoreITransientSet interface {
	CljsCoreITransientSet__()
	X_disjoin_BANG__Arity2(v interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ITransientSet", (*CljsCoreITransientSet)(nil))
}

var X_disjoin_BANG_ *AFn

type CljsCoreIComparable interface {
	CljsCoreIComparable__()
	X_compare_Arity2(y interface{}) float64
}

func init() {
	RegisterProtocol_("cljs.core/IComparable", (*CljsCoreIComparable)(nil))
}

var X_compare *AFn

type CljsCoreIChunk interface {
	CljsCoreIChunk__()
	X_drop_first_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IChunk", (*CljsCoreIChunk)(nil))
}

var X_drop_first *AFn

type CljsCoreIChunkedSeq interface {
	CljsCoreIChunkedSeq__()
	X_chunked_first_Arity1() interface{}
	X_chunked_rest_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IChunkedSeq", (*CljsCoreIChunkedSeq)(nil))
}

var X_chunked_first *AFn

var X_chunked_rest *AFn

type CljsCoreIChunkedNext interface {
	CljsCoreIChunkedNext__()
	X_chunked_next_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IChunkedNext", (*CljsCoreIChunkedNext)(nil))
}

var X_chunked_next *AFn

type CljsCoreINamed interface {
	CljsCoreINamed__()
	X_name_Arity1() string
	X_namespace_Arity1() string
}

func init() {
	RegisterProtocol_("cljs.core/INamed", (*CljsCoreINamed)(nil))
}

var X_name *AFn

var X_namespace *AFn

type CljsCoreIAtom interface {
	CljsCoreIAtom__()
}

func init() {
	RegisterProtocol_("cljs.core/IAtom", (*CljsCoreIAtom)(nil))
}

type CljsCoreIReset interface {
	CljsCoreIReset__()
	X_reset_BANG__Arity2(new_value interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IReset", (*CljsCoreIReset)(nil))
}

var X_reset_BANG_ *AFn

type CljsCoreISwap interface {
	CljsCoreISwap__()
	X_swap_BANG__Arity2(f interface{}) interface{}
	X_swap_BANG__Arity3(f interface{}, a interface{}) interface{}
	X_swap_BANG__Arity4(f interface{}, a interface{}, b interface{}) interface{}
	X_swap_BANG__Arity5(f interface{}, a interface{}, b interface{}, xs interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ISwap", (*CljsCoreISwap)(nil))
}

var X_swap_BANG_ *AFn

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

func (___ *CljsCoreSymbol) X_name_Arity1() string {
	return ___.Name.(string)
}

func (___ *CljsCoreSymbol) X_namespace_Arity1() string {
	return ___.Ns.(string)
}

func (_ *CljsCoreSymbol) CljsCoreIHash__() {}

func (sym *CljsCoreSymbol) X_hash_Arity1() interface{} {
	{
		var h__582__auto__ = sym.X_hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_symbol.X_invoke_Arity1(sym).(float64)
				_ = h__582__auto_____1
				sym.X_hash = h__582__auto_____1

				return h__582__auto_____1
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

func (this *CljsCoreSymbol) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreSymbol) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

func (_ *CljsCoreSymbol) CljsCoreIEquiv__() {}

func (___ *CljsCoreSymbol) X_equiv_Arity2(other interface{}) bool {
	if func() bool { _, instanceof := other.(*CljsCoreSymbol); return instanceof }() {
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
		if i___1 < float64(len(coll.Arr.([]interface{}))) {
			return (coll.Arr.([]interface{})[int(i___1)])
		} else {
			return nil
		}
	}
}

func (coll *CljsCoreIndexedSeq) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	{
		var i___1 = (n.(float64) + coll.I.(float64))
		_ = i___1
		if i___1 < float64(len(coll.Arr.([]interface{}))) {
			return (coll.Arr.([]interface{})[int(i___1)])
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
	if (___.I.(float64) + float64(1)) < float64(len(___.Arr.([]interface{}))) {
		return (&CljsCoreIndexedSeq{___.Arr, (___.I.(float64) + float64(1))})
	} else {
		return nil
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreICounted__() {}

func (___ *CljsCoreIndexedSeq) X_count_Arity1() float64 {
	return (float64(len(___.Arr.([]interface{}))) - ___.I.(float64))
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
	return Array_reduce.X_invoke_Arity4(coll.Arr, f, (coll.Arr.([]interface{})[int(coll.I.(float64))]), (coll.I.(float64) + float64(1)))
}

func (coll *CljsCoreIndexedSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Array_reduce.X_invoke_Arity4(coll.Arr, f, start, coll.I)
}

func (_ *CljsCoreIndexedSeq) CljsCoreISeq__() {}

func (___ *CljsCoreIndexedSeq) X_first_Arity1() interface{} {
	return (___.Arr.([]interface{})[int(___.I.(float64))])
}

func (___ *CljsCoreIndexedSeq) X_rest_Arity1() interface{} {
	if (___.I.(float64) + float64(1)) < float64(len(___.Arr.([]interface{}))) {
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

// Returns the value mapped to key, not-found or nil if key not present.
var Get *AFn

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
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity0()
}

func (___ *CljsCoreMetaFn) X_invoke_Arity1(a interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity1(a)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity2(a interface{}, b interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity2(a, b)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity3(a, b, c)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity4(a, b, c, d)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity5(a, b, c, d, e)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity6(a, b, c, d, e, f)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity7(a, b, c, d, e, f, g)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity8(a, b, c, d, e, f, g, h)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity9(a, b, c, d, e, f, g, h, i)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity10(a, b, c, d, e, f, g, h, i, j)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity11(a, b, c, d, e, f, g, h, i, j, k)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity12(a, b, c, d, e, f, g, h, i, j, k, l)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity13(a, b, c, d, e, f, g, h, i, j, k, l, m)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity14(a, b, c, d, e, f, g, h, i, j, k, l, m, n)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity15(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity16(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity17(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity18(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity19(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
}

func (___ *CljsCoreMetaFn) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	return ___.Afn.(CljsCoreIFn).X_invoke_Arity20(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
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

var Lookup_sentinel interface{}

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

// Returns a sorted sequence of the items in coll. Comp can be
// boolean-valued comparison funcion, or a -/0/+ valued comparator.
// Comp defaults to compare.
var Sort *AFn

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

var Identity *AFn

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

// Returns a random floating point number between 0 (inclusive) and n (default 1) (exclusive).
var Rand *AFn

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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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

func (___ *CljsCoreKeyword) X_name_Arity1() string {
	return ___.Name.(string)
}

func (___ *CljsCoreKeyword) X_namespace_Arity1() string {
	return ___.Ns.(string)
}

func (_ *CljsCoreKeyword) CljsCoreIHash__() {}

func (this *CljsCoreKeyword) X_hash_Arity1() interface{} {
	{
		var h__582__auto__ = this.X_hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_keyword.X_invoke_Arity1(this).(float64)
				_ = h__582__auto_____1
				this.X_hash = h__582__auto_____1

				return h__582__auto_____1
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

func (this *CljsCoreKeyword) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreKeyword) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

func (_ *CljsCoreKeyword) CljsCoreIEquiv__() {}

func (___ *CljsCoreKeyword) X_equiv_Arity2(other interface{}) bool {
	if func() bool { _, instanceof := other.(*CljsCoreKeyword); return instanceof }() {
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
		coll.S = coll.Fn.(CljsCoreIFn).X_invoke_Arity0()

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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
			var ls = coll.S
			_ = ls
			for {
				if func() bool { _, instanceof := ls.(*CljsCoreLazySeq); return instanceof }() {
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
		var return__786 = (___.End.(float64) + float64(1))
		___.End = return__786
		return return__786
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
	return Array_reduce.X_invoke_Arity4(coll.Arr, f, (coll.Arr.([]interface{})[int(coll.Off.(float64))]), (coll.Off.(float64) + float64(1)))
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
	return (coll.Arr.([]interface{})[int((coll.Off.(float64) + i.(float64)))])
}

func (coll *CljsCoreArrayChunk) X_nth_Arity3(i interface{}, not_found interface{}) interface{} {
	if (i.(float64) >= float64(0)) && (i.(float64) < (coll.End.(float64) - coll.Off.(float64))) {
		return (coll.Arr.([]interface{})[int((coll.Off.(float64) + i.(float64)))])
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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

type CljsCoreT349 struct {
	Nil_iter interface{}
	Meta350  interface{}
}

func (_ *CljsCoreT349) CljsCoreObject__() {}

func (___ *CljsCoreT349) HasNext() interface{} {
	return false
}

func (___ *CljsCoreT349) Next() interface{} {
	return (&js.Error{"No such element"})
}

func (___ *CljsCoreT349) Remove() interface{} {
	return (&js.Error{"Unsupported operation"})
}

func (_ *CljsCoreT349) CljsCoreIMeta__() {}

func (_ *CljsCoreT349) CljsCoreIWithMeta__() {}

var X__GT_t349 *AFn

type CljsCoreStringIter struct {
	S interface{}
	I interface{}
}

func (_ *CljsCoreStringIter) CljsCoreObject__() {}

func (___ *CljsCoreStringIter) HasNext() interface{} {
	return (___.I.(float64) < float64(len(___.S.([]interface{}))))
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
	return (___.I.(float64) < float64(len(___.Arr.([]interface{}))))
}

func (___ *CljsCoreArrayIter) Next() interface{} {
	{
		var ret = (___.Arr.([]interface{})[int(___.I.(float64))])
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
				var and__163__auto__ = !(Nil_(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper")))
				_ = and__163__auto__
				if Truth_(and__163__auto__) {
					return Native_invoke_instance_method.X_invoke_Arity3(this.Iter, "HasNext", []interface{}{})
				} else {
					return and__163__auto__
				}
			}()) {
				if Reduced_QMARK_.Arity1IB(this.Xform.(CljsCoreIFn).X_invoke_Arity2(lt, Native_invoke_instance_method.X_invoke_Arity3(this.Iter, "Next", []interface{}{}))) {
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
		return this.Xform.(CljsCoreIFn).X_invoke_Arity1(lt)
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
		var iters___1 = Seq.Arity1IQ(___.Iters)
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
		var n__1042__auto___787 = float64(len(___.Iters.([]interface{})))
		_ = n__1042__auto___787
		{
			var i_788 = float64(0)
			_ = i_788
			for {
				if i_788 < n__1042__auto___787 {
					___.Nexts.([]interface{})[int(i_788)] = Native_invoke_instance_method.X_invoke_Arity3((___.Iters.([]interface{})[int(i_788)]), "Next", []interface{}{})
					i_788 = (i_788 + float64(1))
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
				var and__163__auto__ = !(Nil_(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper")))
				_ = and__163__auto__
				if Truth_(and__163__auto__) {
					return this.HasNext()
				} else {
					return and__163__auto__
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
		return this.Xform.(CljsCoreIFn).X_invoke_Arity1(lt)
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

var CljsCoreLazyTransformer_Create = func(G__789 *AFn) *AFn {
	return Fn(G__789, func(xform interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazyTransformer{Stepper.X_invoke_Arity2(xform, Iter.X_invoke_Arity1(coll)), nil, nil, nil})
	})
}(&AFn{})

var CljsCoreLazyTransformer_CreateMulti = func(G__790 *AFn) *AFn {
	return Fn(G__790, func(xform interface{}, colls interface{}) interface{} {
		{
			var iters = []interface{}{}
			_ = iters
			{
				var seq__353_791 = Seq.Arity1IQ(colls)
				var chunk__354_792 interface{} = nil
				var count__355_793 = float64(0)
				var i__356_794 = float64(0)
				_, _, _, _ = seq__353_791, chunk__354_792, count__355_793, i__356_794
				for {
					if i__356_794 < count__355_793 {
						{
							var coll_795 = chunk__354_792.(CljsCoreIIndexed).X_nth_Arity2(i__356_794)
							_ = coll_795
							js.JSArray_(&iters).Push(Iter.X_invoke_Arity1(coll_795))
							seq__353_791, chunk__354_792, count__355_793, i__356_794 = seq__353_791, chunk__354_792, count__355_793, (i__356_794 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___796 = Seq.Arity1IQ(seq__353_791)
							_ = temp__4222__auto___796
							if Truth_(temp__4222__auto___796) {
								{
									var seq__353_797___1 = temp__4222__auto___796
									_ = seq__353_797___1
									if Chunked_seq_QMARK_.Arity1IB(seq__353_797___1) {
										{
											var c__942__auto___798 = Chunk_first.X_invoke_Arity1(seq__353_797___1)
											_ = c__942__auto___798
											seq__353_791, chunk__354_792, count__355_793, i__356_794 = Chunk_rest.X_invoke_Arity1(seq__353_797___1).(CljsCoreISeq), c__942__auto___798, Count.X_invoke_Arity1(c__942__auto___798).(float64), float64(0)
											continue
										}
									} else {
										{
											var coll_799 = First.X_invoke_Arity1(seq__353_797___1)
											_ = coll_799
											js.JSArray_(&iters).Push(Iter.X_invoke_Arity1(coll_799))
											seq__353_791, chunk__354_792, count__355_793, i__356_794 = Next.Arity1IQ(seq__353_797___1), nil, float64(0), float64(0)
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
	return Native_invoke_func.X_invoke_Arity2(goog.GetUid, []interface{}{this})
}

func (_ *CljsCoreAtom) CljsCoreIWatchable__() {}

func (this *CljsCoreAtom) X_notify_watches_Arity3(oldval interface{}, newval interface{}) interface{} {
	{
		var seq__393 = Seq.Arity1IQ(this.Watches)
		var chunk__394 interface{} = nil
		var count__395 = float64(0)
		var i__396 = float64(0)
		_, _, _, _ = seq__393, chunk__394, count__395, i__396
		for {
			if i__396 < count__395 {
				{
					var vec__397 = chunk__394.(CljsCoreIIndexed).X_nth_Arity2(i__396)
					var key = Nth.X_invoke_Arity3(vec__397, float64(0), nil)
					var f = Nth.X_invoke_Arity3(vec__397, float64(1), nil)
					_, _, _ = vec__397, key, f
					f.(CljsCoreIFn).X_invoke_Arity4(key, this, oldval, newval)
					seq__393, chunk__394, count__395, i__396 = seq__393, chunk__394, count__395, (i__396 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__393)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__393___1 = temp__4222__auto__
							_ = seq__393___1
							if Chunked_seq_QMARK_.Arity1IB(seq__393___1) {
								{
									var c__942__auto__ = Chunk_first.X_invoke_Arity1(seq__393___1)
									_ = c__942__auto__
									seq__393, chunk__394, count__395, i__396 = Chunk_rest.X_invoke_Arity1(seq__393___1).(CljsCoreISeq), c__942__auto__, Count.X_invoke_Arity1(c__942__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__398 = First.X_invoke_Arity1(seq__393___1)
									var key = Nth.X_invoke_Arity3(vec__398, float64(0), nil)
									var f = Nth.X_invoke_Arity3(vec__398, float64(1), nil)
									_, _, _ = vec__398, key, f
									f.(CljsCoreIFn).X_invoke_Arity4(key, this, oldval, newval)
									seq__393, chunk__394, count__395, i__396 = Next.Arity1IQ(seq__393___1), nil, float64(0), float64(0)
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
		var return__800 = Dissoc.X_invoke_Arity2(this.Watches, key)
		this.Watches = return__800
		return return__800
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

// Returns a lazy sequence of the items in coll for which
// (pred item) returns false. pred must be free of side-effects.
// Returns a transducer when no collection is provided.
var Remove *AFn

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
	if reflect.ValueOf(k).Kind() == reflect.Float64 {
		return coll.X_nth_Arity3(k, not_found)
	} else {
		return not_found
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIKVReduce__() {}

func (v *CljsCorePersistentVector) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var step_init = []interface{}{float64(0), init}
		_ = step_init
		{
			var i = float64(0)
			_ = i
			for {
				if i < v.Cnt.(float64) {
					{
						var arr = Unchecked_array_for.X_invoke_Arity2(v, i)
						var len = float64(len(arr.([]interface{})))
						_, _ = arr, len
						{
							var init___1 = func() interface{} {
								var j = float64(0)
								var init___1 = (step_init[int(float64(1))])
								_, _ = j, init___1
								for {
									if j < len {
										{
											var init___2 = f.(CljsCoreIFn).X_invoke_Arity3(init___1, (j + i), (arr.([]interface{})[int(j)]))
											_ = init___2
											if Reduced_QMARK_.Arity1IB(init___2) {
												return init___2
											} else {
												j, init___1 = (j + float64(1)), init___2
												continue
											}
										}
									} else {
										step_init[int(float64(0))] = len
										step_init[int(float64(1))] = init___1
										return init___1
									}
								}
							}()
							_ = init___1
							if Reduced_QMARK_.Arity1IB(init___1) {
								return Deref.X_invoke_Arity1(init___1)
							} else {
								i = (i + (step_init[int(float64(0))]).(float64))
								continue
							}
						}
					}
				} else {
					return (step_init[int(float64(1))])
				}
			}
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIIndexed__() {}

func (coll *CljsCorePersistentVector) X_nth_Arity2(n interface{}) interface{} {
	return (Array_for.X_invoke_Arity2(coll, n).([]interface{})[int(float64(int(n.(float64))&int(float64(31))))])
}

func (coll *CljsCorePersistentVector) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	if (float64(0) <= n.(float64)) && (n.(float64) < coll.Cnt.(float64)) {
		return (Unchecked_array_for.X_invoke_Arity2(coll, n).([]interface{})[int(float64(int(n.(float64))&int(float64(31))))])
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
				new_tail[int(float64(int(n.(float64))&int(float64(31))))] = val
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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

func (v *CljsCorePersistentVector) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	return Ci_reduce.X_invoke_Arity3(v, f, start)
}

func (_ *CljsCorePersistentVector) CljsCoreIAssociative__() {}

func (coll *CljsCorePersistentVector) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	if reflect.ValueOf(k).Kind() == reflect.Float64 {
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
			var len = float64(len(coll.Tail.([]interface{})))
			var new_tail = make([]interface{}, int((len + float64(1))))
			_, _ = len, new_tail
			{
				var n__1042__auto___801 = len
				_ = n__1042__auto___801
				{
					var i_802 = float64(0)
					_ = i_802
					for {
						if i_802 < n__1042__auto___801 {
							new_tail[int(i_802)] = (coll.Tail.([]interface{})[int(i_802)])
							i_802 = (i_802 + float64(1))
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
			var root_overflow_QMARK_ = (float64(uint(coll.Cnt.(float64))>>uint(float64(5))) > float64(int(float64(1))<<uint(coll.Shift.(float64))))
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

func (this *CljsCorePersistentVector) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentVector) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_PersistentVector *AFn

var CljsCorePersistentVector_EMPTY_NODE = (&CljsCoreVectorNode{nil, make([]interface{}, int(float64(32)))})

var CljsCorePersistentVector_EMPTY = (&CljsCorePersistentVector{nil, float64(0), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{}, float64(0)})

var CljsCorePersistentVector_FromArray = func(G__803 *AFn) *AFn {
	return Fn(G__803, func(xs interface{}, no_clone bool) interface{} {
		{
			var l = float64(len(xs.([]interface{})))
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
						var out = v.X_as_transient_Arity1()
						_, _ = i, out
						for {
							if i < l {
								i, out = (i + float64(1)), Conj_BANG_.X_invoke_Arity2(out, (xs___1.([]interface{})[int(i)]))
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

func (_ *CljsCoreChunkedSeq) CljsCoreINext__() {}

func (coll *CljsCoreChunkedSeq) X_next_Arity1() interface{} {
	if (coll.Off.(float64) + float64(1)) < float64(len(coll.Node.([]interface{}))) {
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
	return (coll.Node.([]interface{})[int(coll.Off.(float64))])
}

func (coll *CljsCoreChunkedSeq) X_rest_Arity1() interface{} {
	if (coll.Off.(float64) + float64(1)) < float64(len(coll.Node.([]interface{}))) {
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
		var end = (coll.I.(float64) + float64(len(coll.Node.([]interface{}))))
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

func (coll *CljsCoreChunkedSeq) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreChunkedSeq) CljsCoreICollection__() {}

func (coll *CljsCoreChunkedSeq) X_conj_Arity2(o interface{}) interface{} {
	return Cons.X_invoke_Arity2(o, coll).(*CljsCoreCons)
}

func (_ *CljsCoreChunkedSeq) CljsCoreASeq__() {}

func (_ *CljsCoreChunkedSeq) CljsCoreIChunkedNext__() {}

func (coll *CljsCoreChunkedSeq) X_chunked_next_Arity1() interface{} {
	{
		var end = (coll.I.(float64) + float64(len(coll.Node.([]interface{}))))
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
	if reflect.ValueOf(k).Kind() == reflect.Float64 {
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
	if reflect.ValueOf(key).Kind() == reflect.Float64 {
		return coll.X_assoc_n_Arity3(key, val)
	} else {
		panic((&js.Error{"Subvec's key for assoc must be a number."}))
	}
}

func (_ *CljsCoreSubvec) CljsCoreISeqable__() {}

func (coll *CljsCoreSubvec) X_seq_Arity1() interface{} {
	{
		var subvec_seq = func(subvec_seq *AFn) *AFn {
			return Fn(subvec_seq, func(i interface{}) interface{} {
				if i.(float64) == coll.End.(float64) {
					return nil
				} else {
					return Cons.X_invoke_Arity2(coll.V.(CljsCoreIIndexed).X_nth_Arity2(i), (&CljsCoreLazySeq{nil, func(G__804 *AFn) *AFn {
						return Fn(G__804, func() interface{} {
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

func (this *CljsCoreSubvec) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreSubvec) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
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

func (this *CljsCoreTransientVector) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreTransientVector) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

func (_ *CljsCoreTransientVector) CljsCoreILookup__() {}

func (coll *CljsCoreTransientVector) X_lookup_Arity2(k interface{}) interface{} {
	return coll.X_lookup_Arity3(k, nil)
}

func (coll *CljsCoreTransientVector) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	if reflect.ValueOf(k).Kind() == reflect.Float64 {
		return coll.X_nth_Arity3(k, not_found)
	} else {
		return not_found
	}
}

func (_ *CljsCoreTransientVector) CljsCoreIIndexed__() {}

func (coll *CljsCoreTransientVector) X_nth_Arity2(n interface{}) interface{} {
	if Truth_(Native_get_instance_field.X_invoke_Arity2(coll.Root, "Edit")) {
		return (Array_for.X_invoke_Arity2(coll, n).([]interface{})[int(float64(int(n.(float64))&int(float64(31))))])
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
				tcoll.Tail.([]interface{})[int(float64(int(n.(float64))&int(float64(31))))] = val
				return tcoll
			} else {
				{
					var new_root = func(go_ *AFn) *AFn {
						return Fn(go_, func(level interface{}, node interface{}) interface{} {
							{
								var node___1 = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit"), node)
								_ = node___1
								if level.(float64) == float64(0) {
									Pv_aset.X_invoke_Arity3(node___1, float64(int(n.(float64))&int(float64(31))), val)
									return node___1
								} else {
									{
										var subidx = float64(int(float64(uint(n.(float64))>>uint(level.(float64)))) & int(float64(31)))
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
				if float64(int((tcoll.Cnt.(float64)-float64(1)))&int(float64(31))) > float64(0) {
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
	if reflect.ValueOf(key).Kind() == reflect.Float64 {
		return tcoll.X_assoc_n_BANG__Arity3(key, val)
	} else {
		panic((&js.Error{"TransientVector's key for assoc! must be a number."}))
	}
}

func (_ *CljsCoreTransientVector) CljsCoreITransientCollection__() {}

func (tcoll *CljsCoreTransientVector) X_conj_BANG__Arity2(o interface{}) interface{} {
	if Truth_(Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit")) {
		if (tcoll.Cnt.(float64) - Tail_off.X_invoke_Arity1(tcoll).(float64)) < float64(32) {
			tcoll.Tail.([]interface{})[int(float64(int(tcoll.Cnt.(float64))&int(float64(31))))] = o
			tcoll.Cnt = (tcoll.Cnt.(float64) + float64(1))

			return tcoll
		} else {
			{
				var tail_node = (&CljsCoreVectorNode{Native_get_instance_field.X_invoke_Arity2(tcoll.Root, "Edit"), tcoll.Tail})
				var new_tail = make([]interface{}, int(float64(32)))
				_, _ = tail_node, new_tail
				new_tail[int(float64(0))] = o
				tcoll.Tail = new_tail

				if float64(uint(tcoll.Cnt.(float64))>>uint(float64(5))) > float64(int(float64(1))<<uint(tcoll.Shift.(float64))) {
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
			var or__175__auto__ = coll.Front
			_ = or__175__auto__
			if Truth_(or__175__auto__) {
				return or__175__auto__
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
			var or__175__auto__ = coll.Rear
			_ = or__175__auto__
			if Truth_(or__175__auto__) {
				return or__175__auto__
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

var Obj_clone *AFn

type CljsCoreObjMap struct {
	Meta         interface{}
	Keys         interface{}
	Strobj       interface{}
	Update_count interface{}
	X__hash      interface{}
}

func (_ *CljsCoreObjMap) CljsCoreObject__() {}

func (coll *CljsCoreObjMap) ToString() string {
	return Pr_str_STAR_.X_invoke_Arity1(coll).(string)
}

func (this *CljsCoreObjMap) String() string {
	return this.ToString()
}

func (this *CljsCoreObjMap) Equiv(other interface{}) bool {
	return this.X_equiv_Arity2(other)
}

func (_ *CljsCoreObjMap) CljsCoreILookup__() {}

func (coll *CljsCoreObjMap) X_lookup_Arity2(k interface{}) interface{} {
	return coll.X_lookup_Arity3(k, nil)
}

func (coll *CljsCoreObjMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	if (Truth_(Native_invoke_func.X_invoke_Arity2(goog.IsString, []interface{}{k}))) && (!(Nil_(Scan_array.X_invoke_Arity3(float64(1), k, coll.Keys)))) {
		return (coll.Strobj.([]interface{})[int(k.(float64))])
	} else {
		return not_found
	}
}

func (_ *CljsCoreObjMap) CljsCoreIKVReduce__() {}

func (coll *CljsCoreObjMap) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var len = float64(len(coll.Keys.([]interface{})))
		_ = len
		{
			var keys___1 = Native_invoke_instance_method.X_invoke_Arity3(coll.Keys, "Sort", []interface{}{Obj_map_compare_keys})
			var init___1 = init
			_, _ = keys___1, init___1
			for {
				if Truth_(Seq.Arity1IQ(keys___1)) {
					{
						var k = First.X_invoke_Arity1(keys___1)
						var init___2 = f.(CljsCoreIFn).X_invoke_Arity3(init___1, k, (coll.Strobj.([]interface{})[int(k.(float64))]))
						_, _ = k, init___2
						if Reduced_QMARK_.Arity1IB(init___2) {
							return Deref.X_invoke_Arity1(init___2)
						} else {
							keys___1, init___1 = Rest.Arity1IQ(keys___1), init___2
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

func (_ *CljsCoreObjMap) CljsCoreIMeta__() {}

func (coll *CljsCoreObjMap) X_meta_Arity1() interface{} {
	return coll.Meta
}

func (_ *CljsCoreObjMap) CljsCoreICounted__() {}

func (coll *CljsCoreObjMap) X_count_Arity1() float64 {
	return float64(len(coll.Keys.([]interface{})))
}

func (_ *CljsCoreObjMap) CljsCoreIHash__() {}

func (coll *CljsCoreObjMap) X_hash_Arity1() interface{} {
	{
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
			}
		}
	}
}

func (_ *CljsCoreObjMap) CljsCoreIEquiv__() {}

func (coll *CljsCoreObjMap) X_equiv_Arity2(other interface{}) bool {
	return Truth_(Equiv_map.X_invoke_Arity2(coll, other))
}

func (_ *CljsCoreObjMap) CljsCoreIEditableCollection__() {}

func (coll *CljsCoreObjMap) X_as_transient_Arity1() interface{} {
	return Transient.X_invoke_Arity1(Into.X_invoke_Arity2(CljsCorePersistentHashMap_EMPTY, coll))
}

func (_ *CljsCoreObjMap) CljsCoreIEmptyableCollection__() {}

func (coll *CljsCoreObjMap) X_empty_Arity1() interface{} {
	return With_meta.X_invoke_Arity2(CljsCoreObjMap_EMPTY, coll.Meta)
}

func (_ *CljsCoreObjMap) CljsCoreIMap__() {}

func (coll *CljsCoreObjMap) X_dissoc_Arity2(k interface{}) interface{} {
	if (Truth_(Native_invoke_func.X_invoke_Arity2(goog.IsString, []interface{}{k}))) && (!(Nil_(Scan_array.X_invoke_Arity3(float64(1), k, coll.Keys)))) {
		{
			var new_keys = Aclone.X_invoke_Arity1(coll.Keys).([]interface{})
			var new_strobj = Obj_clone.X_invoke_Arity2(coll.Strobj, coll.Keys)
			_, _ = new_keys, new_strobj
			js.JSArray_(&new_keys).Splice(Scan_array.X_invoke_Arity3(float64(1), k, new_keys), float64(1))
			func(obj, key interface{}) interface{} { delete(obj.(map[string]interface{}), key.(string)); return obj }(new_strobj, k)
			return (&CljsCoreObjMap{coll.Meta, new_keys, new_strobj, (coll.Update_count.(float64) + float64(1)), nil})
		}
	} else {
		return coll
	}
}

func (_ *CljsCoreObjMap) CljsCoreIAssociative__() {}

func (coll *CljsCoreObjMap) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	if Truth_(Native_invoke_func.X_invoke_Arity2(goog.IsString, []interface{}{k})) {
		if (coll.Update_count.(float64) > CljsCoreObjMap_HASHMAP_THRESHOLD) || (float64(len(coll.Keys.([]interface{}))) >= CljsCoreObjMap_HASHMAP_THRESHOLD) {
			return Obj_map__GT_hash_map.X_invoke_Arity3(coll, k, v)
		} else {
			if !(Nil_(Scan_array.X_invoke_Arity3(float64(1), k, coll.Keys))) {
				{
					var new_strobj = Obj_clone.X_invoke_Arity2(coll.Strobj, coll.Keys)
					_ = new_strobj
					new_strobj.([]interface{})[int(k.(float64))] = v
					return (&CljsCoreObjMap{coll.Meta, coll.Keys, new_strobj, (coll.Update_count.(float64) + float64(1)), nil})
				}
			} else {
				{
					var new_strobj = Obj_clone.X_invoke_Arity2(coll.Strobj, coll.Keys)
					var new_keys = Aclone.X_invoke_Arity1(coll.Keys).([]interface{})
					_, _ = new_strobj, new_keys
					new_strobj.([]interface{})[int(k.(float64))] = v
					js.JSArray_(&new_keys).Push(k)
					return (&CljsCoreObjMap{coll.Meta, new_keys, new_strobj, (coll.Update_count.(float64) + float64(1)), nil})
				}
			}
		}
	} else {
		return Obj_map__GT_hash_map.X_invoke_Arity3(coll, k, v)
	}
}

func (coll *CljsCoreObjMap) X_contains_key_QMARK__Arity2(k interface{}) bool {
	if (Truth_(Native_invoke_func.X_invoke_Arity2(goog.IsString, []interface{}{k}))) && (!(Nil_(Scan_array.X_invoke_Arity3(float64(1), k, coll.Keys)))) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCoreObjMap) CljsCoreISeqable__() {}

func (coll *CljsCoreObjMap) X_seq_Arity1() interface{} {
	if float64(len(coll.Keys.([]interface{}))) > float64(0) {
		return Map_.X_invoke_Arity2(func(G__805 *AFn) *AFn {
			return Fn(G__805, func(p1__504_SHARP_ interface{}) interface{} {
				return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__504_SHARP_, (coll.Strobj.([]interface{})[int(p1__504_SHARP_.(float64))])}, nil})
			})
		}(&AFn{}), Native_invoke_instance_method.X_invoke_Arity3(coll.Keys, "Sort", []interface{}{Obj_map_compare_keys})).(*CljsCoreLazySeq)
	} else {
		return nil
	}
}

func (_ *CljsCoreObjMap) CljsCoreIWithMeta__() {}

func (coll *CljsCoreObjMap) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	return (&CljsCoreObjMap{meta___1, coll.Keys, coll.Strobj, coll.Update_count, coll.X__hash})
}

func (_ *CljsCoreObjMap) CljsCoreICollection__() {}

func (coll *CljsCoreObjMap) X_conj_Arity2(entry interface{}) interface{} {
	if Vector_QMARK_.Arity1IB(entry) {
		return coll.X_assoc_Arity3(entry.(CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry.(CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return Reduce.X_invoke_Arity3(X_conj, coll, entry)
	}
}

func (_ *CljsCoreObjMap) CljsCoreIFn__() {}

func (this *CljsCoreObjMap) X_invoke_Arity0() interface{} {
	panic((&js.Error{"Invalid arity: 0"}))
}

func (coll *CljsCoreObjMap) X_invoke_Arity1(k interface{}) interface{} {
	return coll.X_lookup_Arity2(k)
}

func (coll *CljsCoreObjMap) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	return coll.X_lookup_Arity3(k, not_found)
}

func (this *CljsCoreObjMap) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreObjMap) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_ObjMap *AFn

var CljsCoreObjMap_EMPTY = (&CljsCoreObjMap{nil, []interface{}{}, func() interface{} {
	var obj506 = map[string]interface{}{}
	_ = obj506
	return obj506
}(), float64(0), float64(0)})

var CljsCoreObjMap_HASHMAP_THRESHOLD = float64(8)

var CljsCoreObjMap_FromObject = func(G__806 *AFn) *AFn {
	return Fn(G__806, func(ks interface{}, obj interface{}) interface{} {
		return (&CljsCoreObjMap{nil, ks, obj, float64(0), nil})
	})
}(&AFn{})

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
			var vec__509 = First.X_invoke_Arity1(___.S)
			var k = Nth.X_invoke_Arity3(vec__509, float64(0), nil)
			var v = Nth.X_invoke_Arity3(vec__509, float64(1), nil)
			_, _, _ = vec__509, k, v
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
	if coll.I.(float64) < (float64(len(coll.Arr.([]interface{}))) - float64(2)) {
		return (&CljsCorePersistentArrayMapSeq{coll.Arr, (coll.I.(float64) + float64(2)), coll.X_meta})
	} else {
		return nil
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreICounted__() {}

func (coll *CljsCorePersistentArrayMapSeq) X_count_Arity1() float64 {
	return ((float64(len(coll.Arr.([]interface{}))) - coll.I.(float64)) / float64(2))
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
	return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{(coll.Arr.([]interface{})[int(coll.I.(float64))]), (coll.Arr.([]interface{})[int((coll.I.(float64) + float64(1)))])}, nil})
}

func (coll *CljsCorePersistentArrayMapSeq) X_rest_Arity1() interface{} {
	if coll.I.(float64) < (float64(len(coll.Arr.([]interface{}))) - float64(2)) {
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
		var seq__517 = Seq.Arity1IQ(coll)
		var chunk__518 interface{} = nil
		var count__519 = float64(0)
		var i__520 = float64(0)
		_, _, _, _ = seq__517, chunk__518, count__519, i__520
		for {
			if i__520 < count__519 {
				{
					var vec__521 = chunk__518.(CljsCoreIIndexed).X_nth_Arity2(i__520)
					var k = Nth.X_invoke_Arity3(vec__521, float64(0), nil)
					var v = Nth.X_invoke_Arity3(vec__521, float64(1), nil)
					_, _, _ = vec__521, k, v
					f.(CljsCoreIFn).X_invoke_Arity2(v, k)
					seq__517, chunk__518, count__519, i__520 = seq__517, chunk__518, count__519, (i__520 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__517)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__517___1 = temp__4222__auto__
							_ = seq__517___1
							if Chunked_seq_QMARK_.Arity1IB(seq__517___1) {
								{
									var c__942__auto__ = Chunk_first.X_invoke_Arity1(seq__517___1)
									_ = c__942__auto__
									seq__517, chunk__518, count__519, i__520 = Chunk_rest.X_invoke_Arity1(seq__517___1).(CljsCoreISeq), c__942__auto__, Count.X_invoke_Arity1(c__942__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__522 = First.X_invoke_Arity1(seq__517___1)
									var k = Nth.X_invoke_Arity3(vec__522, float64(0), nil)
									var v = Nth.X_invoke_Arity3(vec__522, float64(1), nil)
									_, _, _ = vec__522, k, v
									f.(CljsCoreIFn).X_invoke_Arity2(v, k)
									seq__517, chunk__518, count__519, i__520 = Next.Arity1IQ(seq__517___1), nil, float64(0), float64(0)
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
			return (coll.Arr.([]interface{})[int((idx + float64(1)))])
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIKVReduce__() {}

func (coll *CljsCorePersistentArrayMap) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var len = float64(len(coll.Arr.([]interface{})))
		_ = len
		{
			var i = float64(0)
			var init___1 = init
			_, _ = i, init___1
			for {
				if i < len {
					{
						var init___2 = f.(CljsCoreIFn).X_invoke_Arity3(init___1, (coll.Arr.([]interface{})[int(i)]), (coll.Arr.([]interface{})[int((i + float64(1)))]))
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
	return (&CljsCoreTransientArrayMap{func() interface{} {
		var obj526 = map[string]interface{}{}
		_ = obj526
		return obj526
	}(), float64(len(coll.Arr.([]interface{}))), Aclone.X_invoke_Arity1(coll.Arr).([]interface{})})
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
				var len = float64(len(coll.Arr.([]interface{})))
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
									if X_EQ_.Arity2IIB(k, (coll.Arr.([]interface{})[int(s)])) {
										s, d = (s + float64(2)), d
										continue
									} else {
										new_arr[int(d)] = (coll.Arr.([]interface{})[int(s)])
										new_arr[int((d + float64(1)))] = (coll.Arr.([]interface{})[int((s + float64(1)))])
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
			if reflect.DeepEqual(v, (coll.Arr.([]interface{})[int((idx + float64(1)))])) {
				return coll
			} else {
				{
					var arr___1 = func() []interface{} {
						var G__528 = Aclone.X_invoke_Arity1(coll.Arr).([]interface{})
						_ = G__528
						G__528[int((idx + float64(1)))] = v
						return G__528
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
			var ret = coll
			var es = Seq.Arity1IQ(entry)
			_, _ = ret, es
			for {
				if Nil_(es) {
					return ret
				} else {
					{
						var e = First.X_invoke_Arity1(es)
						_ = e
						if Vector_QMARK_.Arity1IB(e) {
							ret, es = ret.X_assoc_Arity3(e.(CljsCoreIIndexed).X_nth_Arity2(float64(0)), e.(CljsCoreIIndexed).X_nth_Arity2(float64(1))).(*CljsCorePersistentArrayMap), Next.Arity1IQ(es)
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

func (this *CljsCorePersistentArrayMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentArrayMap) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_PersistentArrayMap *AFn

var CljsCorePersistentArrayMap_EMPTY = (&CljsCorePersistentArrayMap{nil, float64(0), []interface{}{}, nil})

var CljsCorePersistentArrayMap_HASHMAP_THRESHOLD = float64(8)

var CljsCorePersistentArrayMap_FromArray = func(G__807 *AFn) *AFn {
	return Fn(G__807, func(arr interface{}, no_clone bool, no_check bool) interface{} {
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
					var cnt = (float64(len(arr___1.([]interface{}))) / float64(2))
					_ = cnt
					return (&CljsCorePersistentArrayMap{nil, cnt, arr___1, nil})
				}
			} else {
				{
					var len = float64(len(arr___1.([]interface{})))
					_ = len
					{
						var i = float64(0)
						var ret = Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY)
						_, _ = i, ret
						for {
							if i < len {
								i, ret = (i + float64(2)), ret.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3((arr___1.([]interface{})[int(i)]), (arr___1.([]interface{})[int((i+float64(1)))]))
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

func (_ *CljsCoreTransientArrayMap) CljsCoreITransientMap__() {}

func (tcoll *CljsCoreTransientArrayMap) X_dissoc_BANG__Arity2(key interface{}) interface{} {
	if Truth_(tcoll.Editable_QMARK_) {
		{
			var idx = Array_map_index_of.X_invoke_Arity2(tcoll, key).(float64)
			_ = idx
			if idx >= float64(0) {
				tcoll.Arr.([]interface{})[int(idx)] = (tcoll.Arr.([]interface{})[int((tcoll.Len.(float64) - float64(2)))])
				tcoll.Arr.([]interface{})[int((idx + float64(1)))] = (tcoll.Arr.([]interface{})[int((tcoll.Len.(float64) - float64(1)))])
				{
					var G__531_808 = tcoll.Arr
					_ = G__531_808
					js.JSArray_(&G__531_808).Pop()
					js.JSArray_(&G__531_808).Pop()
				}
				tcoll.Len = (tcoll.Len.(float64) - float64(2))

			} else {
			}
			return tcoll
		}
	} else {
		panic((&js.Error{"dissoc! after persistent!"}))
	}
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
				if reflect.DeepEqual(val, (tcoll.Arr.([]interface{})[int((idx + float64(1)))])) {
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
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IMapEntry", Str: "cljs.core/IMapEntry", X_hash: float64(535941300), X_meta: nil}), o)) {
			return tcoll.X_assoc_BANG__Arity3(Key.X_invoke_Arity1(o), Val.X_invoke_Arity1(o))
		} else {
			{
				var es = Seq.Arity1IQ(o)
				var tcoll___1 = tcoll
				_, _ = es, tcoll___1
				for {
					{
						var temp__4220__auto__ = First.X_invoke_Arity1(es)
						_ = temp__4220__auto__
						if Truth_(temp__4220__auto__) {
							{
								var e = temp__4220__auto__
								_ = e
								es, tcoll___1 = Next.Arity1IQ(es), tcoll___1.X_assoc_BANG__Arity3(Key.X_invoke_Arity1(e), Val.X_invoke_Arity1(e)).(*CljsCoreTransientArrayMap)
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
				return (tcoll.Arr.([]interface{})[int((idx + float64(1)))])
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
		var bit = float64(int(1) << uint(float64((uint(hash.(float64))>>uint(shift.(float64)))&0x01f)))
		_ = bit
		if float64(int(inode.Bitmap.(float64))&int(bit)) == float64(0) {
			return inode
		} else {
			{
				var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
				var key_or_nil = (inode.Arr.([]interface{})[int((float64(2) * idx))])
				var val_or_node = (inode.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
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
			var len = float64(len(earr.([]interface{})))
			_, _, _ = editable, earr, len
			Native_set_instance_field.X_invoke_Arity3(editable, "Bitmap", float64(int(bit.(float64))^int(Native_get_instance_field.X_invoke_Arity2(editable, "Bitmap").(float64))))
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
		var bit = float64(int(1) << uint(float64((uint(hash.(float64))>>uint(shift.(float64)))&0x01f)))
		_ = bit
		if float64(int(inode.Bitmap.(float64))&int(bit)) == float64(0) {
			return not_found
		} else {
			{
				var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
				var key_or_nil = (inode.Arr.([]interface{})[int((float64(2) * idx))])
				var val_or_node = (inode.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
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
		var bit = float64(int(1) << uint(float64((uint(hash.(float64))>>uint(shift.(float64)))&0x01f)))
		var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
		_, _ = bit, idx
		if float64(int(inode.Bitmap.(float64))&int(bit)) == float64(0) {
			{
				var n = Bit_count.X_invoke_Arity1(inode.Bitmap).(float64)
				_ = n
				if (float64(2) * n) < float64(len(inode.Arr.([]interface{}))) {
					{
						var editable = inode.Ensure_editable(edit___1)
						var earr = Native_get_instance_field.X_invoke_Arity2(editable, "Arr")
						_, _ = editable, earr
						Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)
						Array_copy_downward.X_invoke_Arity5(earr, (float64(2) * idx), earr, (float64(2) * (idx + float64(1))), (float64(2) * (n - idx)))
						earr.([]interface{})[int((float64(2) * idx))] = key
						earr.([]interface{})[int(((float64(2) * idx) + float64(1)))] = val
						Native_set_instance_field.X_invoke_Arity3(editable, "Bitmap", float64(int(Native_get_instance_field.X_invoke_Arity2(editable, "Bitmap").(float64))|int(bit)))
						return editable
					}
				} else {
					if n >= float64(16) {
						{
							var nodes = make([]interface{}, int(float64(32)))
							var jdx = float64((uint(hash.(float64)) >> uint(shift.(float64))) & 0x01f)
							_, _ = nodes, jdx
							nodes[int(jdx)] = Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
							{
								var i_809 = float64(0)
								var j_810 = float64(0)
								_, _ = i_809, j_810
								for {
									if i_809 < float64(32) {
										if float64(int(float64(uint(inode.Bitmap.(float64))>>uint(i_809)))&int(float64(1))) == float64(0) {
											i_809, j_810 = (i_809 + float64(1)), j_810
											continue
										} else {
											nodes[int(i_809)] = func() interface{} {
												if !(Nil_((inode.Arr.([]interface{})[int(j_810)]))) {
													return Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), Hash.X_invoke_Arity1((inode.Arr.([]interface{})[int(j_810)])), (inode.Arr.([]interface{})[int(j_810)]), (inode.Arr.([]interface{})[int((j_810 + float64(1)))]), added_leaf_QMARK_})
												} else {
													return (inode.Arr.([]interface{})[int((j_810 + float64(1)))])
												}
											}()
											i_809, j_810 = (i_809 + float64(1)), (j_810 + float64(2))
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
								Native_set_instance_field.X_invoke_Arity3(editable, "Bitmap", float64(int(Native_get_instance_field.X_invoke_Arity2(editable, "Bitmap").(float64))|int(bit)))
								return editable
							}
						}

					}
				}
			}
		} else {
			{
				var key_or_nil = (inode.Arr.([]interface{})[int((float64(2) * idx))])
				var val_or_node = (inode.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
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
		var bit = float64(int(1) << uint(float64((uint(hash.(float64))>>uint(shift.(float64)))&0x01f)))
		var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
		_, _ = bit, idx
		if float64(int(inode.Bitmap.(float64))&int(bit)) == float64(0) {
			{
				var n = Bit_count.X_invoke_Arity1(inode.Bitmap).(float64)
				_ = n
				if n >= float64(16) {
					{
						var nodes = make([]interface{}, int(float64(32)))
						var jdx = float64((uint(hash.(float64)) >> uint(shift.(float64))) & 0x01f)
						_, _ = nodes, jdx
						nodes[int(jdx)] = Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
						{
							var i_811 = float64(0)
							var j_812 = float64(0)
							_, _ = i_811, j_812
							for {
								if i_811 < float64(32) {
									if float64(int(float64(uint(inode.Bitmap.(float64))>>uint(i_811)))&int(float64(1))) == float64(0) {
										i_811, j_812 = (i_811 + float64(1)), j_812
										continue
									} else {
										nodes[int(i_811)] = func() interface{} {
											if !(Nil_((inode.Arr.([]interface{})[int(j_812)]))) {
												return Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), Hash.X_invoke_Arity1((inode.Arr.([]interface{})[int(j_812)])), (inode.Arr.([]interface{})[int(j_812)]), (inode.Arr.([]interface{})[int((j_812 + float64(1)))]), added_leaf_QMARK_})
											} else {
												return (inode.Arr.([]interface{})[int((j_812 + float64(1)))])
											}
										}()
										i_811, j_812 = (i_811 + float64(1)), (j_812 + float64(2))
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
						return (&CljsCoreBitmapIndexedNode{nil, float64(int(inode.Bitmap.(float64)) | int(bit)), new_arr})
					}
				}
			}
		} else {
			{
				var key_or_nil = (inode.Arr.([]interface{})[int((float64(2) * idx))])
				var val_or_node = (inode.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
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
		var bit = float64(int(1) << uint(float64((uint(hash.(float64))>>uint(shift.(float64)))&0x01f)))
		_ = bit
		if float64(int(inode.Bitmap.(float64))&int(bit)) == float64(0) {
			return not_found
		} else {
			{
				var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
				var key_or_nil = (inode.Arr.([]interface{})[int((float64(2) * idx))])
				var val_or_node = (inode.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
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
		var bit = float64(int(1) << uint(float64((uint(hash.(float64))>>uint(shift.(float64)))&0x01f)))
		_ = bit
		if float64(int(inode.Bitmap.(float64))&int(bit)) == float64(0) {
			return inode
		} else {
			{
				var idx = Bitmap_indexed_node_index.X_invoke_Arity2(inode.Bitmap, bit).(float64)
				var key_or_nil = (inode.Arr.([]interface{})[int((float64(2) * idx))])
				var val_or_node = (inode.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
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
									return (&CljsCoreBitmapIndexedNode{nil, float64(int(inode.Bitmap.(float64)) ^ int(bit)), Remove_pair.X_invoke_Arity2(inode.Arr, idx).([]interface{})})

								}
							}
						}
					}
				} else {
					if Key_test.Arity2IIB(key, key_or_nil) {
						return (&CljsCoreBitmapIndexedNode{nil, float64(int(inode.Bitmap.(float64)) ^ int(bit)), Remove_pair.X_invoke_Arity2(inode.Arr, idx).([]interface{})})
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
		var idx = float64((uint(hash.(float64)) >> uint(shift.(float64))) & 0x01f)
		var node = (inode.Arr.([]interface{})[int(idx)])
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
		var len = float64(len(inode.Arr.([]interface{})))
		_ = len
		{
			var i = float64(0)
			var init___1 = init
			_, _ = i, init___1
			for {
				if i < len {
					{
						var node = (inode.Arr.([]interface{})[int(i)])
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
		var idx = float64((uint(hash.(float64)) >> uint(shift.(float64))) & 0x01f)
		var node = (inode.Arr.([]interface{})[int(idx)])
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
		var idx = float64((uint(hash.(float64)) >> uint(shift.(float64))) & 0x01f)
		var node = (inode.Arr.([]interface{})[int(idx)])
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
		var idx = float64((uint(hash.(float64)) >> uint(shift.(float64))) & 0x01f)
		var node = (inode.Arr.([]interface{})[int(idx)])
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
		var idx = float64((uint(hash.(float64)) >> uint(shift.(float64))) & 0x01f)
		var node = (inode.Arr.([]interface{})[int(idx)])
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
		var idx = float64((uint(hash.(float64)) >> uint(shift.(float64))) & 0x01f)
		var node = (inode.Arr.([]interface{})[int(idx)])
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
					earr.([]interface{})[int(idx)] = (earr.([]interface{})[int(((float64(2) * inode.Cnt.(float64)) - float64(2)))])
					earr.([]interface{})[int((idx + float64(1)))] = (earr.([]interface{})[int(((float64(2) * inode.Cnt.(float64)) - float64(1)))])
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
			if Key_test.Arity2IIB(key, (inode.Arr.([]interface{})[int(idx)])) {
				return (inode.Arr.([]interface{})[int((idx + float64(1)))])
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
				if float64(len(inode.Arr.([]interface{}))) > (float64(2) * inode.Cnt.(float64)) {
					{
						var editable = Edit_and_set.X_invoke_Arity6(inode, edit___1, (float64(2) * inode.Cnt.(float64)), key, ((float64(2) * inode.Cnt.(float64)) + float64(1)), val)
						_ = editable
						Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)
						Native_set_instance_field.X_invoke_Arity3(editable, "Cnt", (Native_get_instance_field.X_invoke_Arity2(editable, "Cnt").(float64) + float64(1)))
						return editable
					}
				} else {
					{
						var len = float64(len(inode.Arr.([]interface{})))
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
				if reflect.DeepEqual((inode.Arr.([]interface{})[int((idx + float64(1)))]), val) {
					return inode
				} else {
					return Edit_and_set.X_invoke_Arity4(inode, edit___1, (idx + float64(1)), val)
				}
			}
		}
	} else {
		return (&CljsCoreBitmapIndexedNode{edit___1, float64(int(1) << uint(float64((uint(inode.Collision_hash.(float64))>>uint(shift.(float64)))&0x01f))), []interface{}{nil, inode, nil, nil}}).Inode_assoc_BANG_(edit___1, shift, hash, key, val, added_leaf_QMARK_)
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
				if X_EQ_.Arity2IIB((inode.Arr.([]interface{})[int(idx)]), val) {
					return inode
				} else {
					return (&CljsCoreHashCollisionNode{nil, inode.Collision_hash, inode.Cnt, Clone_and_set.X_invoke_Arity3(inode.Arr, (idx + float64(1)), val).([]interface{})})
				}
			}
		}
	} else {
		return (&CljsCoreBitmapIndexedNode{nil, float64(int(1) << uint(float64((uint(inode.Collision_hash.(float64))>>uint(shift.(float64)))&0x01f))), []interface{}{nil, inode}}).Inode_assoc(shift, hash, key, val, added_leaf_QMARK_)
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
			if Key_test.Arity2IIB(key, (inode.Arr.([]interface{})[int(idx)])) {
				return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{(inode.Arr.([]interface{})[int(idx)]), (inode.Arr.([]interface{})[int((idx + float64(1)))])}, nil})
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
		return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{(coll.Nodes.([]interface{})[int(coll.I.(float64))]), (coll.Nodes.([]interface{})[int((coll.I.(float64) + float64(1)))])}, nil})
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
		var seq__542 = Seq.Arity1IQ(coll)
		var chunk__543 interface{} = nil
		var count__544 = float64(0)
		var i__545 = float64(0)
		_, _, _, _ = seq__542, chunk__543, count__544, i__545
		for {
			if i__545 < count__544 {
				{
					var vec__546 = chunk__543.(CljsCoreIIndexed).X_nth_Arity2(i__545)
					var k = Nth.X_invoke_Arity3(vec__546, float64(0), nil)
					var v = Nth.X_invoke_Arity3(vec__546, float64(1), nil)
					_, _, _ = vec__546, k, v
					f.(CljsCoreIFn).X_invoke_Arity2(v, k)
					seq__542, chunk__543, count__544, i__545 = seq__542, chunk__543, count__544, (i__545 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__542)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__542___1 = temp__4222__auto__
							_ = seq__542___1
							if Chunked_seq_QMARK_.Arity1IB(seq__542___1) {
								{
									var c__942__auto__ = Chunk_first.X_invoke_Arity1(seq__542___1)
									_ = c__942__auto__
									seq__542, chunk__543, count__544, i__545 = Chunk_rest.X_invoke_Arity1(seq__542___1).(CljsCoreISeq), c__942__auto__, Count.X_invoke_Arity1(c__942__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__547 = First.X_invoke_Arity1(seq__542___1)
									var k = Nth.X_invoke_Arity3(vec__547, float64(0), nil)
									var v = Nth.X_invoke_Arity3(vec__547, float64(1), nil)
									_, _, _ = vec__547, k, v
									f.(CljsCoreIFn).X_invoke_Arity2(v, k)
									seq__542, chunk__543, count__544, i__545 = Next.Arity1IQ(seq__542___1), nil, float64(0), float64(0)
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
				return f.(CljsCoreIFn).X_invoke_Arity3(init, nil, coll.Nil_val)
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
	return (&CljsCoreTransientHashMap{func() interface{} {
		var obj551 = map[string]interface{}{}
		_ = obj551
		return obj551
	}().(bool), coll.Root, coll.Cnt, coll.Has_nil_QMARK_, coll.Nil_val})
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
			var ret = coll
			var es = Seq.Arity1IQ(entry)
			_, _ = ret, es
			for {
				if Nil_(es) {
					return ret
				} else {
					{
						var e = First.X_invoke_Arity1(es)
						_ = e
						if Vector_QMARK_.Arity1IB(e) {
							ret, es = ret.X_assoc_Arity3(e.(CljsCoreIIndexed).X_nth_Arity2(float64(0)), e.(CljsCoreIIndexed).X_nth_Arity2(float64(1))).(*CljsCorePersistentHashMap), Next.Arity1IQ(es)
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

func (this *CljsCorePersistentHashMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentHashMap) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_PersistentHashMap *AFn

var CljsCorePersistentHashMap_EMPTY = (&CljsCorePersistentHashMap{nil, float64(0), nil, false, nil, float64(0)})

var CljsCorePersistentHashMap_FromArrays = func(G__813 *AFn) *AFn {
	return Fn(G__813, func(ks interface{}, vs interface{}) interface{} {
		{
			var len = float64(len(ks.([]interface{})))
			_ = len
			{
				var i = float64(0)
				var out = Transient.X_invoke_Arity1(CljsCorePersistentHashMap_EMPTY)
				_, _ = i, out
				for {
					if i < len {
						i, out = (i + float64(1)), out.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3((ks.([]interface{})[int(i)]), (vs.([]interface{})[int(i)]))
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
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IMapEntry", Str: "cljs.core/IMapEntry", X_hash: float64(535941300), X_meta: nil}), o)) {
			return tcoll.Assoc_BANG_(Key.X_invoke_Arity1(o), Val.X_invoke_Arity1(o))
		} else {
			{
				var es = Seq.Arity1IQ(o)
				var tcoll___1 = tcoll
				_, _ = es, tcoll___1
				for {
					{
						var temp__4220__auto__ = First.X_invoke_Arity1(es)
						_ = temp__4220__auto__
						if Truth_(temp__4220__auto__) {
							{
								var e = temp__4220__auto__
								_ = e
								es, tcoll___1 = Next.Arity1IQ(es), tcoll___1.Assoc_BANG_(Key.X_invoke_Arity1(e), Val.X_invoke_Arity1(e)).(*CljsCoreTransientHashMap)
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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

func (this *CljsCoreBlackNode) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreBlackNode) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
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
	if func() bool { _, instanceof := node.Left.(*CljsCoreRedNode); return instanceof }() {
		return (&CljsCoreRedNode{node.Key, node.Val, Native_invoke_instance_method.X_invoke_Arity3(node.Left, "Blacken", []interface{}{}), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), node.Right, Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil}), nil})
	} else {
		if func() bool { _, instanceof := node.Right.(*CljsCoreRedNode); return instanceof }() {
			return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(node.Right, "Key"), Native_get_instance_field.X_invoke_Arity2(node.Right, "Val"), (&CljsCoreBlackNode{node.Key, node.Val, node.Left, Native_get_instance_field.X_invoke_Arity2(node.Right, "Left"), nil}), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(node.Right, "Right"), Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil}), nil})
		} else {
			return (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), node, Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil})

		}
	}
}

func (node *CljsCoreRedNode) Balance_right(parent interface{}) interface{} {
	if func() bool { _, instanceof := node.Right.(*CljsCoreRedNode); return instanceof }() {
		return (&CljsCoreRedNode{node.Key, node.Val, (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(parent, "Left"), node.Left, nil}), Native_invoke_instance_method.X_invoke_Arity3(node.Right, "Blacken", []interface{}{}), nil})
	} else {
		if func() bool { _, instanceof := node.Left.(*CljsCoreRedNode); return instanceof }() {
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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

func (this *CljsCoreRedNode) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreRedNode) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
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
		var seq__559 = Seq.Arity1IQ(coll)
		var chunk__560 interface{} = nil
		var count__561 = float64(0)
		var i__562 = float64(0)
		_, _, _, _ = seq__559, chunk__560, count__561, i__562
		for {
			if i__562 < count__561 {
				{
					var vec__563 = chunk__560.(CljsCoreIIndexed).X_nth_Arity2(i__562)
					var k = Nth.X_invoke_Arity3(vec__563, float64(0), nil)
					var v = Nth.X_invoke_Arity3(vec__563, float64(1), nil)
					_, _, _ = vec__563, k, v
					f.(CljsCoreIFn).X_invoke_Arity2(v, k)
					seq__559, chunk__560, count__561, i__562 = seq__559, chunk__560, count__561, (i__562 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__559)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__559___1 = temp__4222__auto__
							_ = seq__559___1
							if Chunked_seq_QMARK_.Arity1IB(seq__559___1) {
								{
									var c__942__auto__ = Chunk_first.X_invoke_Arity1(seq__559___1)
									_ = c__942__auto__
									seq__559, chunk__560, count__561, i__562 = Chunk_rest.X_invoke_Arity1(seq__559___1).(CljsCoreISeq), c__942__auto__, Count.X_invoke_Arity1(c__942__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__564 = First.X_invoke_Arity1(seq__559___1)
									var k = Nth.X_invoke_Arity3(vec__564, float64(0), nil)
									var v = Nth.X_invoke_Arity3(vec__564, float64(1), nil)
									_, _, _ = vec__564, k, v
									f.(CljsCoreIFn).X_invoke_Arity2(v, k)
									seq__559, chunk__560, count__561, i__562 = Next.Arity1IQ(seq__559___1), nil, float64(0), float64(0)
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
		var t = coll.Tree
		_ = t
		for {
			if !(Nil_(t)) {
				{
					var c = coll.Comp.(CljsCoreIFn).X_invoke_Arity2(k, Native_get_instance_field.X_invoke_Arity2(t, "Key"))
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
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
			var ret = coll
			var es = Seq.Arity1IQ(entry)
			_, _ = ret, es
			for {
				if Nil_(es) {
					return ret
				} else {
					{
						var e = First.X_invoke_Arity1(es)
						_ = e
						if Vector_QMARK_.Arity1IB(e) {
							ret, es = ret.X_assoc_Arity3(e.(CljsCoreIIndexed).X_nth_Arity2(float64(0)), e.(CljsCoreIIndexed).X_nth_Arity2(float64(1))).(*CljsCorePersistentTreeMap), Next.Arity1IQ(es)
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

func (this *CljsCorePersistentTreeMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentTreeMap) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
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
			var t = coll.Tree
			_, _ = stack, t
			for {
				if !(Nil_(t)) {
					{
						var c = coll.Comp.(CljsCoreIFn).X_invoke_Arity2(k, Native_get_instance_field.X_invoke_Arity2(t, "Key"))
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
// Returns a new object map with supplied mappings.
// @param {...*} var_args
var Obj_map *AFn

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
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INext", Str: "cljs.core/INext", X_hash: float64(-113000046), X_meta: nil}), coll.Mseq)) {
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
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INext", Str: "cljs.core/INext", X_hash: float64(-113000046), X_meta: nil}), coll.Mseq)) {
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
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INext", Str: "cljs.core/INext", X_hash: float64(-113000046), X_meta: nil}), coll.Mseq)) {
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
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INext", Str: "cljs.core/INext", X_hash: float64(-113000046), X_meta: nil}), coll.Mseq)) {
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
		var seq__581 = Seq.Arity1IQ(coll)
		var chunk__582 interface{} = nil
		var count__583 = float64(0)
		var i__584 = float64(0)
		_, _, _, _ = seq__581, chunk__582, count__583, i__584
		for {
			if i__584 < count__583 {
				{
					var vec__585 = chunk__582.(CljsCoreIIndexed).X_nth_Arity2(i__584)
					var k = Nth.X_invoke_Arity3(vec__585, float64(0), nil)
					var v = Nth.X_invoke_Arity3(vec__585, float64(1), nil)
					_, _, _ = vec__585, k, v
					f.(CljsCoreIFn).X_invoke_Arity2(v, k)
					seq__581, chunk__582, count__583, i__584 = seq__581, chunk__582, count__583, (i__584 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__581)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__581___1 = temp__4222__auto__
							_ = seq__581___1
							if Chunked_seq_QMARK_.Arity1IB(seq__581___1) {
								{
									var c__942__auto__ = Chunk_first.X_invoke_Arity1(seq__581___1)
									_ = c__942__auto__
									seq__581, chunk__582, count__583, i__584 = Chunk_rest.X_invoke_Arity1(seq__581___1).(CljsCoreISeq), c__942__auto__, Count.X_invoke_Arity1(c__942__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__586 = First.X_invoke_Arity1(seq__581___1)
									var k = Nth.X_invoke_Arity3(vec__586, float64(0), nil)
									var v = Nth.X_invoke_Arity3(vec__586, float64(1), nil)
									_, _, _ = vec__586, k, v
									f.(CljsCoreIFn).X_invoke_Arity2(v, k)
									seq__581, chunk__582, count__583, i__584 = Next.Arity1IQ(seq__581___1), nil, float64(0), float64(0)
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentHashSet) X_equiv_Arity2(other interface{}) bool {
	return (Set_QMARK_.Arity1IB(other)) && (Count.X_invoke_Arity1(coll).(float64) == Count.X_invoke_Arity1(other).(float64)) && (Every_QMARK_.Arity2IIB(func(G__814 *AFn) *AFn {
		return Fn(G__814, func(p1__574_SHARP_ interface{}) interface{} {
			return Contains_QMARK_.Arity2IIB(coll, p1__574_SHARP_)
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

func (this *CljsCorePersistentHashSet) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentHashSet) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_PersistentHashSet *AFn

var CljsCorePersistentHashSet_EMPTY = (&CljsCorePersistentHashSet{nil, CljsCorePersistentArrayMap_EMPTY, float64(0)})

var CljsCorePersistentHashSet_FromArray = func(G__815 *AFn) *AFn {
	return Fn(G__815, func(items interface{}, no_clone bool) interface{} {
		{
			var len = float64(len(items.([]interface{})))
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
						var out = Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY)
						_, _ = i, out
						for {
							if i < len {
								i, out = (i + float64(1)), out.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3((items.([]interface{})[int(i)]), nil)
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
					var out = Transient.X_invoke_Arity1(CljsCorePersistentHashSet_EMPTY)
					_, _ = i, out
					for {
						if i < len {
							i, out = (i + float64(1)), out.(CljsCoreITransientCollection).X_conj_BANG__Arity2((items.([]interface{})[int(i)]))
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

func (this *CljsCoreTransientHashSet) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCoreTransientHashSet) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
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
		var seq__595 = Seq.Arity1IQ(coll)
		var chunk__596 interface{} = nil
		var count__597 = float64(0)
		var i__598 = float64(0)
		_, _, _, _ = seq__595, chunk__596, count__597, i__598
		for {
			if i__598 < count__597 {
				{
					var vec__599 = chunk__596.(CljsCoreIIndexed).X_nth_Arity2(i__598)
					var k = Nth.X_invoke_Arity3(vec__599, float64(0), nil)
					var v = Nth.X_invoke_Arity3(vec__599, float64(1), nil)
					_, _, _ = vec__599, k, v
					f.(CljsCoreIFn).X_invoke_Arity2(v, k)
					seq__595, chunk__596, count__597, i__598 = seq__595, chunk__596, count__597, (i__598 + float64(1))
					continue
				}
			} else {
				{
					var temp__4222__auto__ = Seq.Arity1IQ(seq__595)
					_ = temp__4222__auto__
					if Truth_(temp__4222__auto__) {
						{
							var seq__595___1 = temp__4222__auto__
							_ = seq__595___1
							if Chunked_seq_QMARK_.Arity1IB(seq__595___1) {
								{
									var c__942__auto__ = Chunk_first.X_invoke_Arity1(seq__595___1)
									_ = c__942__auto__
									seq__595, chunk__596, count__597, i__598 = Chunk_rest.X_invoke_Arity1(seq__595___1).(CljsCoreISeq), c__942__auto__, Count.X_invoke_Arity1(c__942__auto__).(float64), float64(0)
									continue
								}
							} else {
								{
									var vec__600 = First.X_invoke_Arity1(seq__595___1)
									var k = Nth.X_invoke_Arity3(vec__600, float64(0), nil)
									var v = Nth.X_invoke_Arity3(vec__600, float64(1), nil)
									_, _, _ = vec__600, k, v
									f.(CljsCoreIFn).X_invoke_Arity2(v, k)
									seq__595, chunk__596, count__597, i__598 = Next.Arity1IQ(seq__595___1), nil, float64(0), float64(0)
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
		var h__582__auto__ = coll.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_unordered_coll.Arity1IF(coll)
				_ = h__582__auto_____1
				coll.X__hash = h__582__auto_____1

				return h__582__auto_____1
			}
		}
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIEquiv__() {}

func (coll *CljsCorePersistentTreeSet) X_equiv_Arity2(other interface{}) bool {
	return (Set_QMARK_.Arity1IB(other)) && (Count.X_invoke_Arity1(coll).(float64) == Count.X_invoke_Arity1(other).(float64)) && (Every_QMARK_.Arity2IIB(func(G__816 *AFn) *AFn {
		return Fn(G__816, func(p1__588_SHARP_ interface{}) interface{} {
			return Contains_QMARK_.Arity2IIB(coll, p1__588_SHARP_)
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

func (this *CljsCorePersistentTreeSet) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCorePersistentTreeSet) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
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
	return float64(len(___.Arr.([]interface{})))
}

func (___ *CljsCoreArrayList) Clear() interface{} {
	return func() interface{} {
		var return__817 = []interface{}{}
		___.Arr = return__817
		return return__817
	}()
}

func (___ *CljsCoreArrayList) IsEmpty() interface{} {
	return (float64(len(___.Arr.([]interface{}))) == float64(0))
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
		return Native_invoke_func.X_invoke_Arity2(Math.Ceil, []interface{}{((rng.End.(float64) - rng.Start.(float64)) / rng.Step.(float64))}).(float64)
	}
}

func (_ *CljsCoreRange) CljsCoreIHash__() {}

func (rng *CljsCoreRange) X_hash_Arity1() interface{} {
	{
		var h__582__auto__ = rng.X__hash
		_ = h__582__auto__
		if !(Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = Hash_ordered_coll.Arity1IF(rng)
				_ = h__582__auto_____1
				rng.X__hash = h__582__auto_____1

				return h__582__auto_____1
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

var Char_escapes interface{}

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

func (_ *CljsCoreObjMap) CljsCoreIPrintWithWriter__() {}

func (coll *CljsCoreObjMap) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	return Print_map.X_invoke_Arity4(coll, Pr_writer, writer, opts)
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
		___.Value = ___.F.(CljsCoreIFn).X_invoke_Arity0()

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

func init() {
	RegisterProtocol_("cljs.core/IEncodeJS", (*CljsCoreIEncodeJS)(nil))
}

var X_clj__GT_js *AFn

var X_key__GT_js *AFn

var Key__GT_js *AFn

// Recursively transforms ClojureScript values to JavaScript.
// sets/vectors/lists become Arrays, Keywords and Symbol become Strings,
// Maps become Objects. Arbitrary keys are encoded to by key->js.
var Clj__GT_js *AFn

type CljsCoreIEncodeClojure interface {
	CljsCoreIEncodeClojure__()
	X_js__GT_clj_Arity2(options interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IEncodeClojure", (*CljsCoreIEncodeClojure)(nil))
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

func init() {
	RegisterProtocol_("cljs.core/IMultiFn", (*CljsCoreIMultiFn)(nil))
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
	return Native_invoke_func.X_invoke_Arity2(goog.GetUid, []interface{}{this})
}

func (_ *CljsCoreMultiFn) CljsCoreIMultiFn__() {}

func (mf *CljsCoreMultiFn) X_reset_Arity1() interface{} {
	Swap_BANG_.X_invoke_Arity2(mf.Method_table, func(G__818 *AFn) *AFn {
		return Fn(G__818, func(mf___1 interface{}) interface{} {
			return CljsCorePersistentArrayMap_EMPTY
		})
	}(&AFn{}))
	Swap_BANG_.X_invoke_Arity2(mf.Method_cache, func(G__819 *AFn) *AFn {
		return Fn(G__819, func(mf___1 interface{}) interface{} {
			return CljsCorePersistentArrayMap_EMPTY
		})
	}(&AFn{}))
	Swap_BANG_.X_invoke_Arity2(mf.Prefer_table, func(G__820 *AFn) *AFn {
		return Fn(G__820, func(mf___1 interface{}) interface{} {
			return CljsCorePersistentArrayMap_EMPTY
		})
	}(&AFn{}))
	Swap_BANG_.X_invoke_Arity2(mf.Cached_hierarchy, func(G__821 *AFn) *AFn {
		return Fn(G__821, func(mf___1 interface{}) interface{} {
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
	Swap_BANG_.X_invoke_Arity2(mf.Prefer_table, func(G__822 *AFn) *AFn {
		return Fn(G__822, func(old interface{}) interface{} {
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
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity1(a)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity1(a)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity2(a interface{}, b interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity2(a, b)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity2(a, b)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity3(a, b, c)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity3(a, b, c)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity4(a, b, c, d)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity4(a, b, c, d)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity5(a, b, c, d, e)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity5(a, b, c, d, e)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity6(a, b, c, d, e, f)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity6(a, b, c, d, e, f)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity7(a, b, c, d, e, f, g)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity7(a, b, c, d, e, f, g)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity8(a, b, c, d, e, f, g, h)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity8(a, b, c, d, e, f, g, h)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity9(a, b, c, d, e, f, g, h, i)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity9(a, b, c, d, e, f, g, h, i)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity10(a, b, c, d, e, f, g, h, i, j)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity10(a, b, c, d, e, f, g, h, i, j)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity11(a, b, c, d, e, f, g, h, i, j, k)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity11(a, b, c, d, e, f, g, h, i, j, k)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity12(a, b, c, d, e, f, g, h, i, j, k, l)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity12(a, b, c, d, e, f, g, h, i, j, k, l)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity13(a, b, c, d, e, f, g, h, i, j, k, l, m)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity13(a, b, c, d, e, f, g, h, i, j, k, l, m)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity14(a, b, c, d, e, f, g, h, i, j, k, l, m, n)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity14(a, b, c, d, e, f, g, h, i, j, k, l, m, n)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity15(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity15(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity16(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity16(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity17(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity17(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity18(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity18(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity19(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity19(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s)
	}
}

func (mf *CljsCoreMultiFn) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	{
		var dispatch_val = mf.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity20(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
		var target_fn = mf.X_get_method_Arity2(dispatch_val)
		_, _ = dispatch_val, target_fn
		if Truth_(target_fn) {
		} else {
			Throw_no_method_error.X_invoke_Arity2(mf.Name, dispatch_val)
		}
		return target_fn.(CljsCoreIFn).X_invoke_Arity20(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t)
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
	return Native_invoke_func.X_invoke_Arity2(goog_string.HashCode, []interface{}{Pr_str.X_invoke_ArityVariadic(this).(string)})
}

func (_ *CljsCoreUUID) CljsCoreIPrintWithWriter__() {}

func (___ *CljsCoreUUID) X_pr_writer_Arity3(writer interface{}, ______1 interface{}) interface{} {
	return writer.(CljsCoreIWriter).X_write_Arity2(("#uuid \"" + Str.X_invoke_Arity1(___.Uuid).(string) + "\""))
}

func (_ *CljsCoreUUID) CljsCoreIEquiv__() {}

func (___ *CljsCoreUUID) X_equiv_Arity2(other interface{}) bool {
	return (func() bool { _, instanceof := other.(*CljsCoreUUID); return instanceof }()) && (reflect.DeepEqual(___.Uuid, Native_get_instance_field.X_invoke_Arity2(other, "Uuid")))
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
