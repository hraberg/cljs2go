// Compiled by ClojureScript to Go 0.0-2322
// cljs.core
package core

import (
	"reflect"

	goog_array "github.com/hraberg/cljs.go/goog/array"
	goog_object "github.com/hraberg/cljs.go/goog/object"
	goog_string "github.com/hraberg/cljs.go/goog/string"
	"github.com/hraberg/cljs.go/js"
	"github.com/hraberg/cljs.go/js/Math"
)

var X_STAR_clojurescript_version_STAR_ interface{} = nil

var X_STAR_unchecked_if_STAR_ = false

/**
* Each runtime environment provides a different way to print output.
* Whatever function *print-fn* is bound to will be passed any
* Strings which should be printed.
 */
var X_STAR_print_fn_STAR_ = func(_STAR_print_fn_STAR_ *AFn) *AFn {
	return Fn(_STAR_print_fn_STAR_, func(___ interface{}) interface{} {
		panic((&js.Error{"No *print-fn* fn set for evaluation environment"}))
	})
}(&AFn{})

/**
* Set *print-fn* to f.
 */
var Set_print_fn_BANG_ = func(set_print_fn_BANG_ *AFn) *AFn {
	return Fn(set_print_fn_BANG_, func(f interface{}) interface{} {
		return func() interface{} {
			var return__287 = f
			X_STAR_print_fn_STAR_ = return__287
			return return__287
		}()
	})
}(&AFn{})

var X_STAR_flush_on_newline_STAR_ = true

var X_STAR_print_newline_STAR_ = true

var X_STAR_print_readably_STAR_ = true

var X_STAR_print_meta_STAR_ = false

var X_STAR_print_dup_STAR_ = false

var X_STAR_print_length_STAR_ interface{} = nil

var X_STAR_print_level_STAR_ interface{} = nil

var Pr_opts = func(pr_opts *AFn) *AFn {
	return Fn(pr_opts, func() interface{} {
		return &CljsCorePersistentArrayMap{nil, 5, []interface{}{(&CljsCoreKeyword{Ns: nil, Name: "flush-on-newline", Fqn: "flush-on-newline", X_hash: float64(-151457939)}), X_STAR_flush_on_newline_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), X_STAR_print_readably_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}), X_STAR_print_meta_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "dup", Fqn: "dup", X_hash: float64(556298533)}), X_STAR_print_dup_STAR_, (&CljsCoreKeyword{Ns: nil, Name: "print-length", Fqn: "print-length", X_hash: float64(1931866356)}), X_STAR_print_length_STAR_}, nil}
	})
}(&AFn{})

/**
* Set *print-fn* to console.log
 */
var Enable_console_print_BANG_ = func(enable_console_print_BANG_ *AFn) *AFn {
	return Fn(enable_console_print_BANG_, func() interface{} {
		X_STAR_print_newline_STAR_ = false

		return func() interface{} {
			var return__288 = func(G__289 *AFn) *AFn {
				return Fn(G__289, func(args__ ...interface{}) interface{} {
					var args = Array_seq.X_invoke_Arity1(args__[0:])
					_ = args
					return Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(js.Console, "Log"), "Apply", []interface{}{js.Console, Into_array.X_invoke_Arity1(args).(interface{})})
				})
			}(&AFn{})
			X_STAR_print_fn_STAR_ = return__288
			return return__288
		}()
	})
}(&AFn{})

/**
* bound in a repl thread to the most recent value printed
 */
var X_STAR_1 interface{} = nil

/**
* bound in a repl thread to the second most recent value printed
 */
var X_STAR_2 interface{} = nil

/**
* bound in a repl thread to the third most recent value printed
 */
var X_STAR_3 interface{} = nil

var Not_native interface{} = nil

/**
* Tests if 2 arguments are the same object
 */
var Identical_QMARK_ = func(identical_QMARK_ *AFn) *AFn {
	return Fn(identical_QMARK_, func(x interface{}, y interface{}) bool {
		return (x == y)
	})
}(&AFn{})

/**
* Returns true if x is nil, false otherwise.
 */
var Nil_QMARK_ = func(nil_QMARK_ *AFn) *AFn {
	return Fn(nil_QMARK_, func(x interface{}) bool {
		return (x == nil)
	})
}(&AFn{})

var Array_QMARK_ = func(array_QMARK_ *AFn) *AFn {
	return Fn(array_QMARK_, func(x interface{}) bool {
		return reflect.TypeOf(x).Kind() == reflect.Slice
	})
}(&AFn{})

var Number_QMARK_ = func(number_QMARK_ *AFn) *AFn {
	return Fn(number_QMARK_, func(n interface{}) bool {
		return reflect.TypeOf(n).Kind() == reflect.Float64
	})
}(&AFn{})

/**
* Returns true if x is logical false, false otherwise.
 */
var Not = func(not *AFn) *AFn {
	return Fn(not, func(x interface{}) bool {
		if Truth_(x) {
			return false
		} else {
			return true
		}
	})
}(&AFn{})

/**
* Returns true if x is not nil, false otherwise.
 */
var Some_QMARK_ = func(some_QMARK_ *AFn) *AFn {
	return Fn(some_QMARK_, func(x interface{}) bool {
		return !(x == nil)
	})
}(&AFn{})

var String_QMARK_ = func(string_QMARK_ *AFn) *AFn {
	return Fn(string_QMARK_, func(x interface{}) bool {
		return goog.IsString(x).(bool)
	})
}(&AFn{})

/**
* When compiled for a command-line target, whatever
* function *main-fn* is set to will be called with the command-line
* argv as arguments
 */
var X_STAR_main_cli_fn_STAR_ interface{} = nil

var Type__GT_str = func(type__GT_str *AFn) *AFn {
	return Fn(type__GT_str, func(ty interface{}) interface{} {
		{
			var temp__4217__auto__ = Native_get_instance_field.X_invoke_Arity2(ty, "Cljs_lang_ctorStr")
			_ = temp__4217__auto__
			if Truth_(temp__4217__auto__) {
				{
					var s = temp__4217__auto__
					_ = s
					return s
				}
			} else {
				return (`` + Str.X_invoke_Arity1(ty).(string))
			}
		}
	})
}(&AFn{})

/**
* Returns a javascript array, cloned from the passed in array
 */
var Aclone = func(aclone *AFn) *AFn {
	return Fn(aclone, func(arr interface{}) interface{} {
		{
			var len = float64(len(arr.([]interface{})))
			var new_arr = make([]interface{}, int(len))
			_, _ = len, new_arr
			{
				var n__77528__auto___290 = len
				_ = n__77528__auto___290
				{
					var i_291 = float64(0)
					_ = i_291
					for {
						if i_291 < n__77528__auto___290 {
							new_arr[int(i_291)] = (arr.([]interface{})[int(i_291)])
							{
								var G__292 = (i_291 + float64(1))
								i_291 = G__292
								continue
							}
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

/**
* Creates a new javascript array.
* @param {...*} var_args
 */
var Array = func(array *AFn) *AFn {
	return Fn(array, func(var_args interface{}) []interface{} {
		return Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(js.Array, "Prototype"), "Slice"), "Call", []interface{}{arguments}).([]interface{})
	})
}(&AFn{})

/**
* Returns the value at the index.
* @param {...*} var_args
 */
var Aget = func(aget *AFn) *AFn {
	return Fn(aget, func(array interface{}, i interface{}) interface{} {
		return (array.([]interface{})[int(i.(float64))])
	}, func(array_i_idxs__ ...interface{}) interface{} {
		var array = array_i_idxs__[0]
		var i = array_i_idxs__[1]
		var idxs = Array_seq.X_invoke_Arity1(array_i_idxs__[2:])
		_, _, _ = array, i, idxs
		return Apply.X_invoke_Arity3(aget, aget.X_invoke_Arity2(array, i).(interface{}), idxs).(interface{})
	})
}(&AFn{})

/**
* Sets the value at the index.
* @param {...*} var_args
 */
var Aset = func(aset *AFn) *AFn {
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
		return Apply.X_invoke_Arity4(aset, (array.([]interface{})[int(idx.(float64))]), idx2, idxv).(interface{})
	})
}(&AFn{})

/**
* Returns the length of the array. Works on arrays of all types.
 */
var Alength = func(alength *AFn) *AFn {
	return Fn(alength, func(array interface{}) float64 {
		return float64(len(array.([]interface{})))
	})
}(&AFn{})

var Into_array = func(into_array *AFn) *AFn {
	return Fn(into_array, func(aseq interface{}) []interface{} {
		return into_array.X_invoke_Arity2(nil, aseq)
	}, func(type_ interface{}, aseq interface{}) []interface{} {
		return Reduce.X_invoke_Arity3(func(G__293 *AFn) *AFn {
			return Fn(G__293, func(a interface{}, x interface{}) interface{} {
				Native_invoke_instance_method.X_invoke_Arity3(a, "Push", []interface{}{x})
				return a
			})
		}(&AFn{}), []interface{}{}, aseq).(interface{}).([]interface{})
	})
}(&AFn{})

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

var X_invoke = func(_invoke *AFn) *AFn {
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
	})
}(&AFn{})

type CljsCoreICloneable interface {
	CljsCoreICloneable__()
	X_clone_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ICloneable", (*CljsCoreICloneable)(nil))
}

var X_clone = func(_clone *AFn) *AFn {
	return Fn(_clone, func(value interface{}) interface{} {
		return value.(CljsCoreICloneable).X_clone_Arity1()
	})
}(&AFn{})

type CljsCoreICounted interface {
	CljsCoreICounted__()
	X_count_Arity1() float64
}

func init() {
	RegisterProtocol_("cljs.core/ICounted", (*CljsCoreICounted)(nil))
}

var X_count = func(_count *AFn) *AFn {
	return Fn(_count, func(coll interface{}) float64 {
		return coll.(CljsCoreICounted).X_count_Arity1()
	})
}(&AFn{})

type CljsCoreIEmptyableCollection interface {
	CljsCoreIEmptyableCollection__()
	X_empty_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IEmptyableCollection", (*CljsCoreIEmptyableCollection)(nil))
}

var X_empty = func(_empty *AFn) *AFn {
	return Fn(_empty, func(coll interface{}) interface{} {
		return coll.(CljsCoreIEmptyableCollection).X_empty_Arity1()
	})
}(&AFn{})

type CljsCoreICollection interface {
	CljsCoreICollection__()
	X_conj_Arity2(o interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ICollection", (*CljsCoreICollection)(nil))
}

var X_conj = func(_conj *AFn) *AFn {
	return Fn(_conj, func(coll interface{}, o interface{}) interface{} {
		return coll.(CljsCoreICollection).X_conj_Arity2(o)
	})
}(&AFn{})

type CljsCoreIIndexed interface {
	CljsCoreIIndexed__()
	X_nth_Arity2(n interface{}) interface{}
	X_nth_Arity3(n interface{}, not_found interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IIndexed", (*CljsCoreIIndexed)(nil))
}

var X_nth = func(_nth *AFn) *AFn {
	return Fn(_nth, func(coll interface{}, n interface{}) interface{} {
		return coll.(CljsCoreIIndexed).X_nth_Arity2(n)
	}, func(coll interface{}, n interface{}, not_found interface{}) interface{} {
		return coll.(CljsCoreIIndexed).X_nth_Arity3(n, not_found)
	})
}(&AFn{})

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

var X_first = func(_first *AFn) *AFn {
	return Fn(_first, func(coll interface{}) interface{} {
		return coll.(CljsCoreISeq).X_first_Arity1()
	})
}(&AFn{})

var X_rest = func(_rest *AFn) *AFn {
	return Fn(_rest, func(coll interface{}) interface{} {
		return coll.(CljsCoreISeq).X_rest_Arity1()
	})
}(&AFn{})

type CljsCoreINext interface {
	CljsCoreINext__()
	X_next_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/INext", (*CljsCoreINext)(nil))
}

var X_next = func(_next *AFn) *AFn {
	return Fn(_next, func(coll interface{}) interface{} {
		return coll.(CljsCoreINext).X_next_Arity1()
	})
}(&AFn{})

type CljsCoreILookup interface {
	CljsCoreILookup__()
	X_lookup_Arity2(k interface{}) interface{}
	X_lookup_Arity3(k interface{}, not_found interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ILookup", (*CljsCoreILookup)(nil))
}

var X_lookup = func(_lookup *AFn) *AFn {
	return Fn(_lookup, func(o interface{}, k interface{}) interface{} {
		return o.(CljsCoreILookup).X_lookup_Arity2(k)
	}, func(o interface{}, k interface{}, not_found interface{}) interface{} {
		return o.(CljsCoreILookup).X_lookup_Arity3(k, not_found)
	})
}(&AFn{})

type CljsCoreIAssociative interface {
	CljsCoreIAssociative__()
	X_contains_key_QMARK__Arity2(k interface{}) bool
	X_assoc_Arity3(k interface{}, v interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IAssociative", (*CljsCoreIAssociative)(nil))
}

var X_contains_key_QMARK_ = func(_contains_key_QMARK_ *AFn) *AFn {
	return Fn(_contains_key_QMARK_, func(coll interface{}, k interface{}) bool {
		return coll.(CljsCoreIAssociative).X_contains_key_QMARK__Arity2(k)
	})
}(&AFn{})

var X_assoc = func(_assoc *AFn) *AFn {
	return Fn(_assoc, func(coll interface{}, k interface{}, v interface{}) interface{} {
		return coll.(CljsCoreIAssociative).X_assoc_Arity3(k, v)
	})
}(&AFn{})

type CljsCoreIMap interface {
	CljsCoreIMap__()
	X_dissoc_Arity2(k interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IMap", (*CljsCoreIMap)(nil))
}

var X_dissoc = func(_dissoc *AFn) *AFn {
	return Fn(_dissoc, func(coll interface{}, k interface{}) interface{} {
		return coll.(CljsCoreIMap).X_dissoc_Arity2(k)
	})
}(&AFn{})

type CljsCoreIMapEntry interface {
	CljsCoreIMapEntry__()
	X_key_Arity1() interface{}
	X_val_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IMapEntry", (*CljsCoreIMapEntry)(nil))
}

var X_key = func(_key *AFn) *AFn {
	return Fn(_key, func(coll interface{}) interface{} {
		return coll.(CljsCoreIMapEntry).X_key_Arity1()
	})
}(&AFn{})

var X_val = func(_val *AFn) *AFn {
	return Fn(_val, func(coll interface{}) interface{} {
		return coll.(CljsCoreIMapEntry).X_val_Arity1()
	})
}(&AFn{})

type CljsCoreISet interface {
	CljsCoreISet__()
	X_disjoin_Arity2(v interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ISet", (*CljsCoreISet)(nil))
}

var X_disjoin = func(_disjoin *AFn) *AFn {
	return Fn(_disjoin, func(coll interface{}, v interface{}) interface{} {
		return coll.(CljsCoreISet).X_disjoin_Arity2(v)
	})
}(&AFn{})

type CljsCoreIStack interface {
	CljsCoreIStack__()
	X_peek_Arity1() interface{}
	X_pop_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IStack", (*CljsCoreIStack)(nil))
}

var X_peek = func(_peek *AFn) *AFn {
	return Fn(_peek, func(coll interface{}) interface{} {
		return coll.(CljsCoreIStack).X_peek_Arity1()
	})
}(&AFn{})

var X_pop = func(_pop *AFn) *AFn {
	return Fn(_pop, func(coll interface{}) interface{} {
		return coll.(CljsCoreIStack).X_pop_Arity1()
	})
}(&AFn{})

type CljsCoreIVector interface {
	CljsCoreIVector__()
	X_assoc_n_Arity3(n interface{}, val interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IVector", (*CljsCoreIVector)(nil))
}

var X_assoc_n = func(_assoc_n *AFn) *AFn {
	return Fn(_assoc_n, func(coll interface{}, n interface{}, val interface{}) interface{} {
		return coll.(CljsCoreIVector).X_assoc_n_Arity3(n, val)
	})
}(&AFn{})

type CljsCoreIDeref interface {
	CljsCoreIDeref__()
	X_deref_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IDeref", (*CljsCoreIDeref)(nil))
}

var X_deref = func(_deref *AFn) *AFn {
	return Fn(_deref, func(o interface{}) interface{} {
		return o.(CljsCoreIDeref).X_deref_Arity1()
	})
}(&AFn{})

type CljsCoreIDerefWithTimeout interface {
	CljsCoreIDerefWithTimeout__()
	X_deref_with_timeout_Arity3(msec interface{}, timeout_val interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IDerefWithTimeout", (*CljsCoreIDerefWithTimeout)(nil))
}

var X_deref_with_timeout = func(_deref_with_timeout *AFn) *AFn {
	return Fn(_deref_with_timeout, func(o interface{}, msec interface{}, timeout_val interface{}) interface{} {
		return o.(CljsCoreIDerefWithTimeout).X_deref_with_timeout_Arity3(msec, timeout_val)
	})
}(&AFn{})

type CljsCoreIMeta interface {
	CljsCoreIMeta__()
	X_meta_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IMeta", (*CljsCoreIMeta)(nil))
}

var X_meta = func(_meta *AFn) *AFn {
	return Fn(_meta, func(o interface{}) interface{} {
		return o.(CljsCoreIMeta).X_meta_Arity1()
	})
}(&AFn{})

type CljsCoreIWithMeta interface {
	CljsCoreIWithMeta__()
	X_with_meta_Arity2(meta interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IWithMeta", (*CljsCoreIWithMeta)(nil))
}

var X_with_meta = func(_with_meta *AFn) *AFn {
	return Fn(_with_meta, func(o interface{}, meta interface{}) interface{} {
		return o.(CljsCoreIWithMeta).X_with_meta_Arity2(meta)
	})
}(&AFn{})

type CljsCoreIReduce interface {
	CljsCoreIReduce__()
	X_reduce_Arity2(f interface{}) interface{}
	X_reduce_Arity3(f interface{}, start interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IReduce", (*CljsCoreIReduce)(nil))
}

var X_reduce = func(_reduce *AFn) *AFn {
	return Fn(_reduce, func(coll interface{}, f interface{}) interface{} {
		return coll.(CljsCoreIReduce).X_reduce_Arity2(f)
	}, func(coll interface{}, f interface{}, start interface{}) interface{} {
		return coll.(CljsCoreIReduce).X_reduce_Arity3(f, start)
	})
}(&AFn{})

type CljsCoreIKVReduce interface {
	CljsCoreIKVReduce__()
	X_kv_reduce_Arity3(f interface{}, init interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IKVReduce", (*CljsCoreIKVReduce)(nil))
}

var X_kv_reduce = func(_kv_reduce *AFn) *AFn {
	return Fn(_kv_reduce, func(coll interface{}, f interface{}, init interface{}) interface{} {
		return coll.(CljsCoreIKVReduce).X_kv_reduce_Arity3(f, init)
	})
}(&AFn{})

type CljsCoreIEquiv interface {
	CljsCoreIEquiv__()
	X_equiv_Arity2(other interface{}) bool
}

func init() {
	RegisterProtocol_("cljs.core/IEquiv", (*CljsCoreIEquiv)(nil))
}

var X_equiv = func(_equiv *AFn) *AFn {
	return Fn(_equiv, func(o interface{}, other interface{}) bool {
		return o.(CljsCoreIEquiv).X_equiv_Arity2(other)
	})
}(&AFn{})

type CljsCoreIHash interface {
	CljsCoreIHash__()
	X_hash_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IHash", (*CljsCoreIHash)(nil))
}

var X_hash = func(_hash *AFn) *AFn {
	return Fn(_hash, func(o interface{}) interface{} {
		return o.(CljsCoreIHash).X_hash_Arity1()
	})
}(&AFn{})

type CljsCoreISeqable interface {
	CljsCoreISeqable__()
	X_seq_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ISeqable", (*CljsCoreISeqable)(nil))
}

var X_seq = func(_seq *AFn) *AFn {
	return Fn(_seq, func(o interface{}) interface{} {
		return o.(CljsCoreISeqable).X_seq_Arity1()
	})
}(&AFn{})

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

var X_rseq = func(_rseq *AFn) *AFn {
	return Fn(_rseq, func(coll interface{}) interface{} {
		return coll.(CljsCoreIReversible).X_rseq_Arity1()
	})
}(&AFn{})

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

var X_sorted_seq = func(_sorted_seq *AFn) *AFn {
	return Fn(_sorted_seq, func(coll interface{}, ascending_QMARK_ interface{}) interface{} {
		return coll.(CljsCoreISorted).X_sorted_seq_Arity2(ascending_QMARK_)
	})
}(&AFn{})

var X_sorted_seq_from = func(_sorted_seq_from *AFn) *AFn {
	return Fn(_sorted_seq_from, func(coll interface{}, k interface{}, ascending_QMARK_ interface{}) interface{} {
		return coll.(CljsCoreISorted).X_sorted_seq_from_Arity3(k, ascending_QMARK_)
	})
}(&AFn{})

var X_entry_key = func(_entry_key *AFn) *AFn {
	return Fn(_entry_key, func(coll interface{}, entry interface{}) interface{} {
		return coll.(CljsCoreISorted).X_entry_key_Arity2(entry)
	})
}(&AFn{})

var X_comparator = func(_comparator *AFn) *AFn {
	return Fn(_comparator, func(coll interface{}) interface{} {
		return coll.(CljsCoreISorted).X_comparator_Arity1()
	})
}(&AFn{})

type CljsCoreIWriter interface {
	CljsCoreIWriter__()
	X_write_Arity2(s interface{}) interface{}
	X_flush_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IWriter", (*CljsCoreIWriter)(nil))
}

var X_write = func(_write *AFn) *AFn {
	return Fn(_write, func(writer interface{}, s interface{}) interface{} {
		return writer.(CljsCoreIWriter).X_write_Arity2(s)
	})
}(&AFn{})

var X_flush = func(_flush *AFn) *AFn {
	return Fn(_flush, func(writer interface{}) interface{} {
		return writer.(CljsCoreIWriter).X_flush_Arity1()
	})
}(&AFn{})

type CljsCoreIPrintWithWriter interface {
	CljsCoreIPrintWithWriter__()
	X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IPrintWithWriter", (*CljsCoreIPrintWithWriter)(nil))
}

var X_pr_writer = func(_pr_writer *AFn) *AFn {
	return Fn(_pr_writer, func(o interface{}, writer interface{}, opts interface{}) interface{} {
		return o.(CljsCoreIPrintWithWriter).X_pr_writer_Arity3(writer, opts)
	})
}(&AFn{})

type CljsCoreIPending interface {
	CljsCoreIPending__()
	X_realized_QMARK__Arity1() bool
}

func init() {
	RegisterProtocol_("cljs.core/IPending", (*CljsCoreIPending)(nil))
}

var X_realized_QMARK_ = func(_realized_QMARK_ *AFn) *AFn {
	return Fn(_realized_QMARK_, func(d interface{}) bool {
		return d.(CljsCoreIPending).X_realized_QMARK__Arity1()
	})
}(&AFn{})

type CljsCoreIWatchable interface {
	CljsCoreIWatchable__()
	X_notify_watches_Arity3(oldval interface{}, newval interface{}) interface{}
	X_add_watch_Arity3(key interface{}, f interface{}) interface{}
	X_remove_watch_Arity2(key interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IWatchable", (*CljsCoreIWatchable)(nil))
}

var X_notify_watches = func(_notify_watches *AFn) *AFn {
	return Fn(_notify_watches, func(this interface{}, oldval interface{}, newval interface{}) interface{} {
		return this.(CljsCoreIWatchable).X_notify_watches_Arity3(oldval, newval)
	})
}(&AFn{})

var X_add_watch = func(_add_watch *AFn) *AFn {
	return Fn(_add_watch, func(this interface{}, key interface{}, f interface{}) interface{} {
		return this.(CljsCoreIWatchable).X_add_watch_Arity3(key, f)
	})
}(&AFn{})

var X_remove_watch = func(_remove_watch *AFn) *AFn {
	return Fn(_remove_watch, func(this interface{}, key interface{}) interface{} {
		return this.(CljsCoreIWatchable).X_remove_watch_Arity2(key)
	})
}(&AFn{})

type CljsCoreIEditableCollection interface {
	CljsCoreIEditableCollection__()
	X_as_transient_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IEditableCollection", (*CljsCoreIEditableCollection)(nil))
}

var X_as_transient = func(_as_transient *AFn) *AFn {
	return Fn(_as_transient, func(coll interface{}) interface{} {
		return coll.(CljsCoreIEditableCollection).X_as_transient_Arity1()
	})
}(&AFn{})

type CljsCoreITransientCollection interface {
	CljsCoreITransientCollection__()
	X_conj_BANG__Arity2(val interface{}) interface{}
	X_persistent_BANG__Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ITransientCollection", (*CljsCoreITransientCollection)(nil))
}

var X_conj_BANG_ = func(_conj_BANG_ *AFn) *AFn {
	return Fn(_conj_BANG_, func(tcoll interface{}, val interface{}) interface{} {
		return tcoll.(CljsCoreITransientCollection).X_conj_BANG__Arity2(val)
	})
}(&AFn{})

var X_persistent_BANG_ = func(_persistent_BANG_ *AFn) *AFn {
	return Fn(_persistent_BANG_, func(tcoll interface{}) interface{} {
		return tcoll.(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
	})
}(&AFn{})

type CljsCoreITransientAssociative interface {
	CljsCoreITransientAssociative__()
	X_assoc_BANG__Arity3(key interface{}, val interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ITransientAssociative", (*CljsCoreITransientAssociative)(nil))
}

var X_assoc_BANG_ = func(_assoc_BANG_ *AFn) *AFn {
	return Fn(_assoc_BANG_, func(tcoll interface{}, key interface{}, val interface{}) interface{} {
		return tcoll.(CljsCoreITransientAssociative).X_assoc_BANG__Arity3(key, val)
	})
}(&AFn{})

type CljsCoreITransientMap interface {
	CljsCoreITransientMap__()
	X_dissoc_BANG__Arity2(key interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ITransientMap", (*CljsCoreITransientMap)(nil))
}

var X_dissoc_BANG_ = func(_dissoc_BANG_ *AFn) *AFn {
	return Fn(_dissoc_BANG_, func(tcoll interface{}, key interface{}) interface{} {
		return tcoll.(CljsCoreITransientMap).X_dissoc_BANG__Arity2(key)
	})
}(&AFn{})

type CljsCoreITransientVector interface {
	CljsCoreITransientVector__()
	X_assoc_n_BANG__Arity3(n interface{}, val interface{}) interface{}
	X_pop_BANG__Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ITransientVector", (*CljsCoreITransientVector)(nil))
}

var X_assoc_n_BANG_ = func(_assoc_n_BANG_ *AFn) *AFn {
	return Fn(_assoc_n_BANG_, func(tcoll interface{}, n interface{}, val interface{}) interface{} {
		return tcoll.(CljsCoreITransientVector).X_assoc_n_BANG__Arity3(n, val)
	})
}(&AFn{})

var X_pop_BANG_ = func(_pop_BANG_ *AFn) *AFn {
	return Fn(_pop_BANG_, func(tcoll interface{}) interface{} {
		return tcoll.(CljsCoreITransientVector).X_pop_BANG__Arity1()
	})
}(&AFn{})

type CljsCoreITransientSet interface {
	CljsCoreITransientSet__()
	X_disjoin_BANG__Arity2(v interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/ITransientSet", (*CljsCoreITransientSet)(nil))
}

var X_disjoin_BANG_ = func(_disjoin_BANG_ *AFn) *AFn {
	return Fn(_disjoin_BANG_, func(tcoll interface{}, v interface{}) interface{} {
		return tcoll.(CljsCoreITransientSet).X_disjoin_BANG__Arity2(v)
	})
}(&AFn{})

type CljsCoreIComparable interface {
	CljsCoreIComparable__()
	X_compare_Arity2(y interface{}) float64
}

func init() {
	RegisterProtocol_("cljs.core/IComparable", (*CljsCoreIComparable)(nil))
}

var X_compare = func(_compare *AFn) *AFn {
	return Fn(_compare, func(x interface{}, y interface{}) float64 {
		return x.(CljsCoreIComparable).X_compare_Arity2(y)
	})
}(&AFn{})

type CljsCoreIChunk interface {
	CljsCoreIChunk__()
	X_drop_first_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IChunk", (*CljsCoreIChunk)(nil))
}

var X_drop_first = func(_drop_first *AFn) *AFn {
	return Fn(_drop_first, func(coll interface{}) interface{} {
		return coll.(CljsCoreIChunk).X_drop_first_Arity1()
	})
}(&AFn{})

type CljsCoreIChunkedSeq interface {
	CljsCoreIChunkedSeq__()
	X_chunked_first_Arity1() interface{}
	X_chunked_rest_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IChunkedSeq", (*CljsCoreIChunkedSeq)(nil))
}

var X_chunked_first = func(_chunked_first *AFn) *AFn {
	return Fn(_chunked_first, func(coll interface{}) interface{} {
		return coll.(CljsCoreIChunkedSeq).X_chunked_first_Arity1()
	})
}(&AFn{})

var X_chunked_rest = func(_chunked_rest *AFn) *AFn {
	return Fn(_chunked_rest, func(coll interface{}) interface{} {
		return coll.(CljsCoreIChunkedSeq).X_chunked_rest_Arity1()
	})
}(&AFn{})

type CljsCoreIChunkedNext interface {
	CljsCoreIChunkedNext__()
	X_chunked_next_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IChunkedNext", (*CljsCoreIChunkedNext)(nil))
}

var X_chunked_next = func(_chunked_next *AFn) *AFn {
	return Fn(_chunked_next, func(coll interface{}) interface{} {
		return coll.(CljsCoreIChunkedNext).X_chunked_next_Arity1()
	})
}(&AFn{})

type CljsCoreINamed interface {
	CljsCoreINamed__()
	X_name_Arity1() string
	X_namespace_Arity1() string
}

func init() {
	RegisterProtocol_("cljs.core/INamed", (*CljsCoreINamed)(nil))
}

var X_name = func(_name *AFn) *AFn {
	return Fn(_name, func(x interface{}) string {
		return x.(CljsCoreINamed).X_name_Arity1()
	})
}(&AFn{})

var X_namespace = func(_namespace *AFn) *AFn {
	return Fn(_namespace, func(x interface{}) string {
		return x.(CljsCoreINamed).X_namespace_Arity1()
	})
}(&AFn{})

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

var X_reset_BANG_ = func(_reset_BANG_ *AFn) *AFn {
	return Fn(_reset_BANG_, func(o interface{}, new_value interface{}) interface{} {
		return o.(CljsCoreIReset).X_reset_BANG__Arity2(new_value)
	})
}(&AFn{})

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

var X_swap_BANG_ = func(_swap_BANG_ *AFn) *AFn {
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

type CljsCoreStringBufferWriter struct{ Sb interface{} }

func (_ *CljsCoreStringBufferWriter) CljsCoreIWriter__() {}
func (self__ *CljsCoreStringBufferWriter) X_write_Arity2(s interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return Native_invoke_instance_method.X_invoke_Arity3(self__.Sb, "Append", []interface{}{s})
	}
}

func (self__ *CljsCoreStringBufferWriter) X_flush_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return nil
	}
}

var X__GT_StringBufferWriter = func(__GT_StringBufferWriter *AFn) *AFn {
	return Fn(__GT_StringBufferWriter, func(sb interface{}) interface{} {
		return (&CljsCoreStringBufferWriter{sb})
	})
}(&AFn{})

/**
* Support so that collections can implement toString without
* loading all the printing machinery.
 */
var Pr_str_STAR_ = func(pr_str_STAR_ *AFn) *AFn {
	return Fn(pr_str_STAR_, func(obj interface{}) interface{} {
		{
			var sb = (&goog_string.StringBuffer{})
			var writer = (&CljsCoreStringBufferWriter{sb})
			_, _ = sb, writer
			obj.X_pr_writer_Arity3(writer, Pr_opts.X_invoke_Arity0().(CljsCoreIMap)).(interface{})
			writer.X_flush_Arity1().(interface{})
			return (`` + Str.X_invoke_Arity1(sb).(string))
		}
	})
}(&AFn{})

var Int_rotate_left = func(int_rotate_left *AFn) *AFn {
	return Fn(int_rotate_left, func(x interface{}, n interface{}) float64 {
		return float64(int(float64(int(x.(float64))<<uint(n.(float64)))) | int(float64(uint(x.(float64))>>uint((-n.(float64))))))
	})
}(&AFn{})

var Imul = func(imul *AFn) *AFn {
	return Fn(imul, func(a interface{}, b interface{}) float64 {
		return Math.Imul(a, b).(float64)
	})
}(&AFn{})

var M3_seed = float64(0)

var M3_C1 = float64(3432918353)

var M3_C2 = float64(461845907)

var M3_mix_K1 = func(m3_mix_K1 *AFn) *AFn {
	return Fn(m3_mix_K1, func(k1 interface{}) float64 {
		return Imul.X_invoke_Arity2(Int_rotate_left.X_invoke_Arity2(Imul.Arity2IIF(k1, M3_C1), float64(15)), M3_C2)
	})
}(&AFn{})

var M3_mix_H1 = func(m3_mix_H1 *AFn) *AFn {
	return Fn(m3_mix_H1, func(h1 interface{}, k1 interface{}) float64 {
		return (Imul.X_invoke_Arity2(Int_rotate_left.X_invoke_Arity2(float64(int(h1.(float64))^int(k1.(float64))), float64(13)), float64(5)) + float64(3864292196))
	})
}(&AFn{})

var M3_fmix = func(m3_fmix *AFn) *AFn {
	return Fn(m3_fmix, func(h1 interface{}, len interface{}) float64 {
		{
			var h1___1 = h1
			var h1___2 = float64(int(h1___1.(float64)) ^ int(len.(float64)))
			var h1___3 = float64(int(h1___2) ^ int(float64(uint(h1___2)>>uint(float64(16)))))
			var h1___4 = Imul.X_invoke_Arity2(h1___3, float64(2246822507))
			var h1___5 = float64(int(h1___4) ^ int(float64(uint(h1___4)>>uint(float64(13)))))
			var h1___6 = Imul.X_invoke_Arity2(h1___5, float64(3266489909))
			var h1___7 = float64(int(h1___6) ^ int(float64(uint(h1___6)>>uint(float64(16)))))
			_, _, _, _, _, _, _ = h1___1, h1___2, h1___3, h1___4, h1___5, h1___6, h1___7
			return h1___7
		}
	})
}(&AFn{})

var M3_hash_int = func(m3_hash_int *AFn) *AFn {
	return Fn(m3_hash_int, func(in interface{}) float64 {
		if in.(float64) == float64(0) {
			return in
		} else {
			{
				var k1 = M3_mix_K1.Arity1IF(in)
				var h1 = M3_mix_H1.X_invoke_Arity2(M3_seed, k1)
				_, _ = k1, h1
				return M3_fmix.X_invoke_Arity2(h1, float64(4))
			}
		}
	})
}(&AFn{})

var M3_hash_unencoded_chars = func(m3_hash_unencoded_chars *AFn) *AFn {
	return Fn(m3_hash_unencoded_chars, func(in interface{}) float64 {
		{
			var h1 = func() interface{} {
				var i = float64(1)
				var h1 = M3_seed
				_, _ = i, h1
				for {
					if i < float64(len(in.([]interface{}))) {
						{
							var G__294 = (i + float64(2))
							var G__295 = M3_mix_H1.X_invoke_Arity2(h1, M3_mix_K1.X_invoke_Arity1(float64(int(Native_invoke_instance_method.X_invoke_Arity3(in, "CharCodeAt", []interface{}{(i - float64(1))}).(float64))|int(float64(int(Native_invoke_instance_method.X_invoke_Arity3(in, "CharCodeAt", []interface{}{i}).(float64))<<uint(float64(16)))))))
							i = G__294
							h1 = G__295
							continue
						}
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
			return M3_fmix.X_invoke_Arity2(h1___1, Imul.X_invoke_Arity2(float64(2), float64(len(in.([]interface{})))))
		}
	})
}(&AFn{})

var String_hash_cache = func() interface{} {
	var obj297 = map[string]interface{}{}
	_ = obj297
	return obj297
}()

var String_hash_cache_count = float64(0)

var Hash_string_STAR_ = func(hash_string_STAR_ *AFn) *AFn {
	return Fn(hash_string_STAR_, func(s interface{}) interface{} {
		if !(s == nil) {
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
								{
									var G__298 = (i + float64(1))
									var G__299 = (Imul.X_invoke_Arity2(float64(31), hash) + Native_invoke_instance_method.X_invoke_Arity3(s, "CharCodeAt", []interface{}{i}).(float64))
									i = G__298
									hash = G__299
									continue
								}
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

var Add_to_string_hash_cache = func(add_to_string_hash_cache *AFn) *AFn {
	return Fn(add_to_string_hash_cache, func(k interface{}) interface{} {
		{
			var h = Hash_string_STAR_.X_invoke_Arity1(k).(float64)
			_ = h
			String_hash_cache.([]interface{})[int(k.(float64))] = h
			String_hash_cache_count = (String_hash_cache_count.(float64) + float64(1))

			return h
		}
	})
}(&AFn{})

var Hash_string = func(hash_string *AFn) *AFn {
	return Fn(hash_string, func(k interface{}) interface{} {
		if String_hash_cache_count.(float64) > float64(255) {
			String_hash_cache = func() interface{} {
				var obj303 = map[string]interface{}{}
				_ = obj303
				return obj303
			}()

			String_hash_cache_count = float64(0)

		} else {
		}
		{
			var h = (String_hash_cache.([]interface{})[int(k.(float64))])
			_ = h
			if reflect.TypeOf(h).Kind() == reflect.Float64 {
				return h
			} else {
				return Add_to_string_hash_cache.X_invoke_Arity1(k).(float64)
			}
		}
	})
}(&AFn{})

var Hash = func(hash *AFn) *AFn {
	return Fn(hash, func(o interface{}) interface{} {
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IHash", Str: "cljs.core/IHash", X_hash: float64(-32453786), X_meta: nil}), o)) {
			return o.X_hash_Arity1().(interface{})
		} else {
			if reflect.TypeOf(o).Kind() == reflect.Float64 {
				return (Math.Floor(o).(float64) % float64(2147483647))
			} else {
				if o == true {
					return float64(1)
				} else {
					if o == false {
						return float64(0)
					} else {
						if reflect.TypeOf(o).Kind() == reflect.String {
							return M3_hash_int.X_invoke_Arity1(Hash_string.X_invoke_Arity1(o))
						} else {
							if o == nil {
								return float64(0)
							} else {
								return o.(CljsCoreIHash).X_hash_Arity1().(interface{})

							}
						}
					}
				}
			}
		}
	})
}(&AFn{})

var Hash_combine = func(hash_combine *AFn) *AFn {
	return Fn(hash_combine, func(seed interface{}, hash interface{}) interface{} {
		return float64(int(seed.(float64)) ^ int((((hash.(float64) + float64(2654435769)) + float64(int(seed.(float64))<<uint(float64(6)))) + float64(int(seed.(float64))>>uint(float64(2))))))
	})
}(&AFn{})

var Instance_QMARK_ = func(instance_QMARK_ *AFn) *AFn {
	return Fn(instance_QMARK_, func(t interface{}, o interface{}) bool {
		return func() bool { _, instanceof := o.(*t); return instanceof }()
	})
}(&AFn{})

var Symbol_QMARK_ = func(symbol_QMARK_ *AFn) *AFn {
	return Fn(symbol_QMARK_, func(x interface{}) bool {
		return func() bool { _, instanceof := x.(*Symbol); return instanceof }()
	})
}(&AFn{})

var Hash_symbol = func(hash_symbol *AFn) *AFn {
	return Fn(hash_symbol, func(sym interface{}) interface{} {
		return Hash_combine.X_invoke_Arity2(M3_hash_unencoded_chars.Arity1IF(Native_get_instance_field.X_invoke_Arity2(sym, "Name")), Hash_string.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(sym, "Ns"))).(float64)
	})
}(&AFn{})

var Compare_symbols = func(compare_symbols *AFn) *AFn {
	return Fn(compare_symbols, func(a interface{}, b interface{}) interface{} {
		if Truth_(X_EQ_.X_invoke_Arity2(a, b).(interface{})) {
			return float64(0)
		} else {
			if Truth_(func() interface{} {
				var and__76620__auto__ = Not.Arity1IB(Native_get_instance_field.X_invoke_Arity2(a, "Ns"))
				_ = and__76620__auto__
				if and__76620__auto__ {
					return Native_get_instance_field.X_invoke_Arity2(b, "Ns")
				} else {
					return and__76620__auto__
				}
			}()) {
				return float64(-1)
			} else {
				if Truth_(Native_get_instance_field.X_invoke_Arity2(a, "Ns")) {
					if Not.Arity1IB(Native_get_instance_field.X_invoke_Arity2(b, "Ns")) {
						return float64(1)
					} else {
						{
							var nsc = Compare.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(a, "Ns"), Native_get_instance_field.X_invoke_Arity2(b, "Ns")).(interface{})
							_ = nsc
							if nsc.(float64) == float64(0) {
								return Compare.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(a, "Name"), Native_get_instance_field.X_invoke_Arity2(b, "Name")).(interface{})
							} else {
								return nsc
							}
						}
					}
				} else {
					return Compare.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(a, "Name"), Native_get_instance_field.X_invoke_Arity2(b, "Name")).(interface{})

				}
			}
		}
	})
}(&AFn{})

type CljsCoreSymbol struct {
	Ns     interface{}
	Name   interface{}
	Str    interface{}
	X_hash interface{}
	X_meta interface{}
}

func (_ *CljsCoreSymbol) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreSymbol) X_pr_writer_Arity3(writer interface{}, ___ interface{}) interface{} {
	{
		var o___1 = self__
		_ = o___1
		return writer.(CljsCoreIWriter).X_write_Arity2(self__.Str).(interface{})
	}
}

func (_ *CljsCoreSymbol) CljsCoreINamed__() {}
func (self__ *CljsCoreSymbol) X_name_Arity1() string {
	{
		var ______1 = self__
		_ = ______1
		return self__.Name.(string)
	}
}

func (self__ *CljsCoreSymbol) X_namespace_Arity1() string {
	{
		var ______1 = self__
		_ = ______1
		return self__.Ns.(string)
	}
}

func (_ *CljsCoreSymbol) CljsCoreIHash__() {}
func (self__ *CljsCoreSymbol) X_hash_Arity1() interface{} {
	{
		var sym___1 = self__
		_ = sym___1
		{
			var h__77043__auto__ = self__.X_hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_symbol.X_invoke_Arity1(sym___1).(float64)
					_ = h__77043__auto_____1
					self__.X_hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreSymbol) CljsCoreIWithMeta__() {}
func (self__ *CljsCoreSymbol) X_with_meta_Arity2(new_meta interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCoreSymbol{self__.Ns, self__.Name, self__.Str, self__.X_hash, new_meta})
	}
}

func (_ *CljsCoreSymbol) CljsCoreIMeta__() {}
func (self__ *CljsCoreSymbol) X_meta_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.X_meta
	}
}

func (_ *CljsCoreSymbol) CljsCoreIFn__() {}
func (self__ *CljsCoreSymbol) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity1(coll interface{}) interface{} {
	{
		var sym___1 = self__
		_ = sym___1
		return coll.(CljsCoreILookup).X_lookup_Arity3(sym___1, nil).(interface{})
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity2(coll interface{}, not_found interface{}) interface{} {
	{
		var sym___1 = self__
		_ = sym___1
		return coll.(CljsCoreILookup).X_lookup_Arity3(sym___1, not_found).(interface{})
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCoreSymbol) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

func (_ *CljsCoreSymbol) CljsCoreIEquiv__() {}
func (self__ *CljsCoreSymbol) X_equiv_Arity2(other interface{}) bool {
	{
		var ______1 = self__
		_ = ______1
		if func() bool { _, instanceof := other.(*CljsCoreSymbol); return instanceof }() {
			return (self__.Str == Native_get_instance_field.X_invoke_Arity2(other, "Str"))
		} else {
			return false
		}
	}
}

func (_ *CljsCoreSymbol) CljsCoreObject__() {}
func (self__ *CljsCoreSymbol) ToString() string {
	{
		var ______1 = self__
		_ = ______1
		return self__.Str.(string)
	}
}

func (self__ *CljsCoreSymbol) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreSymbol) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

var X__GT_Symbol = func(__GT_Symbol *AFn) *AFn {
	return Fn(__GT_Symbol, func(ns interface{}, name interface{}, str interface{}, _hash interface{}, _meta interface{}) interface{} {
		return (&CljsCoreSymbol{ns, name, str, _hash, _meta})
	})
}(&AFn{})

var Symbol = func(symbol *AFn) *AFn {
	return Fn(symbol, func(name interface{}) interface{} {
		if func() bool { _, instanceof := name.(*CljsCoreSymbol); return instanceof }() {
			return name
		} else {
			return symbol.X_invoke_Arity2(nil, name).(*CljsCoreSymbol)
		}
	}, func(ns interface{}, name interface{}) interface{} {
		{
			var sym_str = func() interface{} {
				if !(ns == nil) {
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

var Clone = func(clone *AFn) *AFn {
	return Fn(clone, func(value interface{}) interface{} {
		return value.(CljsCoreICloneable).X_clone_Arity1()
	})
}(&AFn{})

var Cloneable_QMARK_ = func(cloneable_QMARK_ *AFn) *AFn {
	return Fn(cloneable_QMARK_, func(value interface{}) interface{} {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ICloneable", Str: "cljs.core/ICloneable", X_hash: float64(-1712216461), X_meta: nil}), value)
	})
}(&AFn{})

/**
* Returns a seq on the collection. If the collection is
* empty, returns nil.  (seq nil) returns nil. seq also works on
* Strings.
 */
var Seq = func(seq *AFn) *AFn {
	return Fn(seq, func(coll interface{}) CljsCoreISeq {
		if coll == nil {
			return nil
		} else {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeqable", Str: "cljs.core/ISeqable", X_hash: float64(137437203), X_meta: nil}), coll)) {
				return coll.X_seq_Arity1()
			} else {
				if reflect.TypeOf(coll).Kind() == reflect.Slice {
					if float64(len(coll.([]interface{}))) == float64(0) {
						return nil
					} else {
						return (&CljsCoreIndexedSeq{coll, float64(0)})
					}
				} else {
					if reflect.TypeOf(coll).Kind() == reflect.String {
						if float64(len(coll.([]interface{}))) == float64(0) {
							return nil
						} else {
							return (&CljsCoreIndexedSeq{coll, float64(0)})
						}
					} else {
						if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeqable", Str: "cljs.core/ISeqable", X_hash: float64(137437203), X_meta: nil}), coll)) {
							return coll.(CljsCoreISeqable).X_seq_Arity1()
						} else {
							panic((&js.Error{(`` + Str.X_invoke_Arity1(coll).(string) + " is not ISeqable")}))

						}
					}
				}
			}
		}
	})
}(&AFn{})

/**
* Returns the first item in the collection. Calls seq on its
* argument. If coll is nil, returns nil.
 */
var First = func(first *AFn) *AFn {
	return Fn(first, func(coll interface{}) interface{} {
		if coll == nil {
			return nil
		} else {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeq", Str: "cljs.core/ISeq", X_hash: float64(230133392), X_meta: nil}), coll)) {
				return coll.X_first_Arity1().(interface{})
			} else {
				{
					var s = Seq.Arity1IQ(coll)
					_ = s
					if s == nil {
						return nil
					} else {
						return s.X_first_Arity1().(interface{})
					}
				}
			}
		}
	})
}(&AFn{})

/**
* Returns a possibly empty seq of the items after the first. Calls seq on its
* argument.
 */
var Rest = func(rest *AFn) *AFn {
	return Fn(rest, func(coll interface{}) CljsCoreISeq {
		if !(coll == nil) {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeq", Str: "cljs.core/ISeq", X_hash: float64(230133392), X_meta: nil}), coll)) {
				return coll.X_rest_Arity1().(CljsCoreISeq)
			} else {
				{
					var s = Seq.Arity1IQ(coll)
					_ = s
					if Truth_(s) {
						return s.X_rest_Arity1().(CljsCoreISeq)
					} else {
						return CljsCoreList_EMPTY.(CljsCoreISeq)
					}
				}
			}
		} else {
			return CljsCoreList_EMPTY.(CljsCoreISeq)
		}
	})
}(&AFn{})

/**
* Returns a seq of the items after the first. Calls seq on its
* argument.  If there are no more items, returns nil
 */
var Next = func(next *AFn) *AFn {
	return Fn(next, func(coll interface{}) CljsCoreISeq {
		if coll == nil {
			return nil
		} else {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INext", Str: "cljs.core/INext", X_hash: float64(-113000046), X_meta: nil}), coll)) {
				return coll.X_next_Arity1()
			} else {
				return Seq.X_invoke_Arity1(Rest.Arity1IQ(coll))
			}
		}
	})
}(&AFn{})

/**
* Equality. Returns true if x equals y, false if not. Compares
* numbers and collections in a type-independent manner.  Clojure's immutable data
* structures define -equiv (and thus =) as a value, not an identity,
* comparison.
* @param {...*} var_args
 */
var X_EQ_ = func(_EQ_ *AFn) *AFn {
	return Fn(_EQ_, func(x interface{}) bool {
		return true
	}, func(x interface{}, y interface{}) bool {
		if x == nil {
			return (y == nil)
		} else {
			return (x == y) || (x.(CljsCoreIEquiv).X_equiv_Arity2(y))
		}
	}, func(x_y_more__ ...interface{}) interface{} {
		var x = x_y_more__[0]
		var y = x_y_more__[1]
		var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
		_, _, _ = x, y, more
		for {
			if _EQ_.Arity2IIB(x, y) {
				if Truth_(Next.Arity1IQ(more)) {
					{
						var G__304 = y
						var G__305 = First.X_invoke_Arity1(more)
						var G__306 = Next.Arity1IQ(more)
						x = G__304
						y = G__305
						more = G__306
						continue
					}
				} else {
					return _EQ_.X_invoke_Arity2(y, First.X_invoke_Arity1(more))
				}
			} else {
				return false
			}
		}
	})
}(&AFn{})

/**
* Mix final collection hash for ordered or unordered collections.
* hash-basis is the combined collection hash, count is the number
* of elements included in the basis. Note this is the hash code
* consistent with =, different from .hashCode.
* See http://clojure.org/data_structures#hash for full algorithms.
 */
var Mix_collection_hash = func(mix_collection_hash *AFn) *AFn {
	return Fn(mix_collection_hash, func(hash_basis interface{}, count interface{}) float64 {
		{
			var h1 = M3_seed
			var k1 = M3_mix_K1.Arity1IF(hash_basis)
			var h1___1 = M3_mix_H1.X_invoke_Arity2(h1, k1)
			_, _, _ = h1, k1, h1___1
			return M3_fmix.X_invoke_Arity2(h1___1, count)
		}
	})
}(&AFn{})

/**
* Returns the hash code, consistent with =, for an external ordered
* collection implementing Iterable.
* See http://clojure.org/data_structures#hash for full algorithms.
 */
var Hash_ordered_coll = func(hash_ordered_coll *AFn) *AFn {
	return Fn(hash_ordered_coll, func(coll interface{}) float64 {
		{
			var n = float64(0)
			var hash_code = float64(1)
			var coll___1 = Seq.Arity1IQ(coll)
			_, _, _ = n, hash_code, coll___1
			for {
				if !(coll___1 == nil) {
					{
						var G__307 = (n + float64(1))
						var G__308 = float64(int((Imul.X_invoke_Arity2(float64(31), hash_code) + Hash.X_invoke_Arity1(First.X_invoke_Arity1(coll___1)).(float64))) | int(float64(0)))
						var G__309 = Next.X_invoke_Arity1(coll___1)
						n = G__307
						hash_code = G__308
						coll___1 = G__309
						continue
					}
				} else {
					return Mix_collection_hash.X_invoke_Arity2(hash_code, n)
				}
			}
		}
	})
}(&AFn{})

/**
* Returns the hash code, consistent with =, for an external unordered
* collection implementing Iterable. For maps, the iterator should
* return map entries whose hash is computed as
* (hash-ordered-coll [k v]).
* See http://clojure.org/data_structures#hash for full algorithms.
 */
var Hash_unordered_coll = func(hash_unordered_coll *AFn) *AFn {
	return Fn(hash_unordered_coll, func(coll interface{}) float64 {
		{
			var n = float64(0)
			var hash_code = float64(0)
			var coll___1 = Seq.Arity1IQ(coll)
			_, _, _ = n, hash_code, coll___1
			for {
				if !(coll___1 == nil) {
					{
						var G__310 = (n + float64(1))
						var G__311 = float64(int((hash_code + Hash.X_invoke_Arity1(First.X_invoke_Arity1(coll___1)).(float64))) | int(float64(0)))
						var G__312 = Next.X_invoke_Arity1(coll___1)
						n = G__310
						hash_code = G__311
						coll___1 = G__312
						continue
					}
				} else {
					return Mix_collection_hash.X_invoke_Arity2(hash_code, n)
				}
			}
		}
	})
}(&AFn{})

/**
* Returns a number one greater than num.
 */
var Inc = func(inc *AFn) *AFn {
	return Fn(inc, func(x interface{}) interface{} {
		return (x.(float64) + float64(1))
	})
}(&AFn{})

type CljsCoreReduced struct{ Val interface{} }

func (_ *CljsCoreReduced) CljsCoreIDeref__() {}
func (self__ *CljsCoreReduced) X_deref_Arity1() interface{} {
	{
		var o___1 = self__
		_ = o___1
		return self__.Val
	}
}

var X__GT_Reduced = func(__GT_Reduced *AFn) *AFn {
	return Fn(__GT_Reduced, func(val interface{}) interface{} {
		return (&CljsCoreReduced{val})
	})
}(&AFn{})

/**
* Wraps x in a way such that a reduce will terminate with the value x
 */
var Reduced = func(reduced *AFn) *AFn {
	return Fn(reduced, func(x interface{}) interface{} {
		return (&CljsCoreReduced{x})
	})
}(&AFn{})

/**
* Returns true if x is the result of a call to reduced
 */
var Reduced_QMARK_ = func(reduced_QMARK_ *AFn) *AFn {
	return Fn(reduced_QMARK_, func(r interface{}) bool {
		return func() bool { _, instanceof := r.(*CljsCoreReduced); return instanceof }()
	})
}(&AFn{})

var Deref = func(deref *AFn) *AFn {
	return Fn(deref, func(o interface{}) interface{} {
		return o.(CljsCoreIDeref).X_deref_Arity1().(interface{})
	})
}(&AFn{})

/**
* Accepts any collection which satisfies the ICount and IIndexed protocols and
* reduces them without incurring seq initialization
 */
var Ci_reduce = func(ci_reduce *AFn) *AFn {
	return Fn(ci_reduce, func(cicoll interface{}, f interface{}) interface{} {
		{
			var cnt = cicoll.(CljsCoreICounted).X_count_Arity1()
			_ = cnt
			if cnt == float64(0) {
				return f.(CljsCoreIFn).X_invoke_Arity0().(interface{})
			} else {
				{
					var val = cicoll.(CljsCoreIIndexed).X_nth_Arity2(float64(0)).(interface{})
					var n = float64(1)
					_, _ = val, n
					for {
						if n < cnt {
							{
								var nval = f.(CljsCoreIFn).X_invoke_Arity2(val, cicoll.(CljsCoreIIndexed).X_nth_Arity2(n).(interface{})).(interface{})
								_ = nval
								if Reduced_QMARK_.X_invoke_Arity1(nval) {
									return Deref.X_invoke_Arity1(nval).(interface{})
								} else {
									{
										var G__313 = nval
										var G__314 = (n + float64(1))
										val = G__313
										n = G__314
										continue
									}
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
							var nval = f.(CljsCoreIFn).X_invoke_Arity2(val___1, cicoll.(CljsCoreIIndexed).X_nth_Arity2(n).(interface{})).(interface{})
							_ = nval
							if Reduced_QMARK_.X_invoke_Arity1(nval) {
								return Deref.X_invoke_Arity1(nval).(interface{})
							} else {
								{
									var G__315 = nval
									var G__316 = (n + float64(1))
									val___1 = G__315
									n = G__316
									continue
								}
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
							var nval = f.(CljsCoreIFn).X_invoke_Arity2(val___1, cicoll.(CljsCoreIIndexed).X_nth_Arity2(n).(interface{})).(interface{})
							_ = nval
							if Reduced_QMARK_.X_invoke_Arity1(nval) {
								return Deref.X_invoke_Arity1(nval).(interface{})
							} else {
								{
									var G__317 = nval
									var G__318 = (n.(float64) + float64(1))
									val___1 = G__317
									n = G__318
									continue
								}
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

var Array_reduce = func(array_reduce *AFn) *AFn {
	return Fn(array_reduce, func(arr interface{}, f interface{}) interface{} {
		{
			var cnt = float64(len(arr.([]interface{})))
			_ = cnt
			if float64(len(arr.([]interface{}))) == float64(0) {
				return f.(CljsCoreIFn).X_invoke_Arity0().(interface{})
			} else {
				{
					var val = (arr.([]interface{})[int(float64(0))])
					var n = float64(1)
					_, _ = val, n
					for {
						if n < cnt {
							{
								var nval = f.(CljsCoreIFn).X_invoke_Arity2(val, (arr.([]interface{})[int(n)])).(interface{})
								_ = nval
								if Reduced_QMARK_.X_invoke_Arity1(nval) {
									return Deref.X_invoke_Arity1(nval).(interface{})
								} else {
									{
										var G__319 = nval
										var G__320 = (n + float64(1))
										val = G__319
										n = G__320
										continue
									}
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
							var nval = f.(CljsCoreIFn).X_invoke_Arity2(val___1, (arr.([]interface{})[int(n)])).(interface{})
							_ = nval
							if Reduced_QMARK_.X_invoke_Arity1(nval) {
								return Deref.X_invoke_Arity1(nval).(interface{})
							} else {
								{
									var G__321 = nval
									var G__322 = (n + float64(1))
									val___1 = G__321
									n = G__322
									continue
								}
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
							var nval = f.(CljsCoreIFn).X_invoke_Arity2(val___1, (arr.([]interface{})[int(n.(float64))])).(interface{})
							_ = nval
							if Reduced_QMARK_.X_invoke_Arity1(nval) {
								return Deref.X_invoke_Arity1(nval).(interface{})
							} else {
								{
									var G__323 = nval
									var G__324 = (n.(float64) + float64(1))
									val___1 = G__323
									n = G__324
									continue
								}
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

/**
* Returns true if coll implements count in constant time
 */
var Counted_QMARK_ = func(counted_QMARK_ *AFn) *AFn {
	return Fn(counted_QMARK_, func(x interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ICounted", Str: "cljs.core/ICounted", X_hash: float64(-1299011378), X_meta: nil}), x).(bool)
	})
}(&AFn{})

/**
* Returns true if coll implements nth in constant time
 */
var Indexed_QMARK_ = func(indexed_QMARK_ *AFn) *AFn {
	return Fn(indexed_QMARK_, func(x interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IIndexed", Str: "cljs.core/IIndexed", X_hash: float64(-436490749), X_meta: nil}), x).(bool)
	})
}(&AFn{})

type CljsCoreIndexedSeq struct {
	Arr interface{}
	I   interface{}
}

func (_ *CljsCoreIndexedSeq) CljsCoreObject__() {}
func (self__ *CljsCoreIndexedSeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreIndexedSeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreIndexedSeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreIIndexed__() {}
func (self__ *CljsCoreIndexedSeq) X_nth_Arity2(n interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var i___1 = (n.(float64) + self__.I.(float64))
			_ = i___1
			if i___1 < float64(len(self__.Arr.([]interface{}))) {
				return (self__.Arr.([]interface{})[int(i___1)])
			} else {
				return nil
			}
		}
	}
}

func (self__ *CljsCoreIndexedSeq) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var i___1 = (n.(float64) + self__.I.(float64))
			_ = i___1
			if i___1 < float64(len(self__.Arr.([]interface{}))) {
				return (self__.Arr.([]interface{})[int(i___1)])
			} else {
				return not_found
			}
		}
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreICloneable__() {}
func (self__ *CljsCoreIndexedSeq) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCoreIndexedSeq{self__.Arr, self__.I})
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreINext__() {}
func (self__ *CljsCoreIndexedSeq) X_next_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		if (self__.I.(float64) + float64(1)) < float64(len(self__.Arr.([]interface{}))) {
			return (&CljsCoreIndexedSeq{self__.Arr, (self__.I.(float64) + float64(1))})
		} else {
			return nil
		}
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreICounted__() {}
func (self__ *CljsCoreIndexedSeq) X_count_Arity1() float64 {
	{
		var ______1 = self__
		_ = ______1
		return (float64(len(self__.Arr.([]interface{}))) - self__.I.(float64))
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreIReversible__() {}
func (self__ *CljsCoreIndexedSeq) X_rseq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var c = coll___1.X_count_Arity1()
			_ = c
			if c > float64(0) {
				return (&CljsCoreRSeq{coll___1, (c - float64(1)), nil})
			} else {
				return nil
			}
		}
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreIHash__() {}
func (self__ *CljsCoreIndexedSeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Hash_ordered_coll.X_invoke_Arity1(coll___1)
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreIEquiv__() {}
func (self__ *CljsCoreIndexedSeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(interface{}).(bool)
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreIndexedSeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Native_get_instance_field.X_invoke_Arity2(List, "EMPTY")
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreIReduce__() {}
func (self__ *CljsCoreIndexedSeq) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Array_reduce.X_invoke_Arity4(self__.Arr, f, (self__.Arr.([]interface{})[int(self__.I.(float64))]), (self__.I.(float64) + float64(1)))
	}
}

func (self__ *CljsCoreIndexedSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Array_reduce.X_invoke_Arity4(self__.Arr, f, start, self__.I)
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreISeq__() {}
func (self__ *CljsCoreIndexedSeq) X_first_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (self__.Arr.([]interface{})[int(self__.I.(float64))])
	}
}

func (self__ *CljsCoreIndexedSeq) X_rest_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		if (self__.I.(float64) + float64(1)) < float64(len(self__.Arr.([]interface{}))) {
			return (&CljsCoreIndexedSeq{self__.Arr, (self__.I.(float64) + float64(1))})
		} else {
			return Native_get_instance_field.X_invoke_Arity2(List, "EMPTY")
		}
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreISeqable__() {}
func (self__ *CljsCoreIndexedSeq) X_seq_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		return this___1
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreISequential__() {}
func (_ *CljsCoreIndexedSeq) CljsCoreICollection__() {}
func (self__ *CljsCoreIndexedSeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(interface{})
	}
}

func (_ *CljsCoreIndexedSeq) CljsCoreASeq__() {}

var X__GT_IndexedSeq = func(__GT_IndexedSeq *AFn) *AFn {
	return Fn(__GT_IndexedSeq, func(arr interface{}, i interface{}) interface{} {
		return (&CljsCoreIndexedSeq{arr, i})
	})
}(&AFn{})

var Prim_seq = func(prim_seq *AFn) *AFn {
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

var Array_seq = func(array_seq *AFn) *AFn {
	return Fn(array_seq, func(array interface{}) interface{} {
		return Prim_seq.X_invoke_Arity2(array, float64(0))
	}, func(array interface{}, i interface{}) interface{} {
		return Prim_seq.X_invoke_Arity2(array, i)
	})
}(&AFn{})

type CljsCoreRSeq struct {
	Ci   interface{}
	I    interface{}
	Meta interface{}
}

func (_ *CljsCoreRSeq) CljsCoreObject__() {}
func (self__ *CljsCoreRSeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreRSeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreRSeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreRSeq) CljsCoreIMeta__() {}
func (self__ *CljsCoreRSeq) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreRSeq) CljsCoreICloneable__() {}
func (self__ *CljsCoreRSeq) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCoreRSeq{self__.Ci, self__.I, self__.Meta})
	}
}

func (_ *CljsCoreRSeq) CljsCoreINext__() {}
func (self__ *CljsCoreRSeq) X_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.I.(float64) > float64(0) {
			return (&CljsCoreRSeq{self__.Ci, (self__.I.(float64) - float64(1)), nil})
		} else {
			return nil
		}
	}
}

func (_ *CljsCoreRSeq) CljsCoreICounted__() {}
func (self__ *CljsCoreRSeq) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return (self__.I.(float64) + float64(1))
	}
}

func (_ *CljsCoreRSeq) CljsCoreIHash__() {}
func (self__ *CljsCoreRSeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Hash_ordered_coll.X_invoke_Arity1(coll___1)
	}
}

func (_ *CljsCoreRSeq) CljsCoreIEquiv__() {}
func (self__ *CljsCoreRSeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(interface{}).(bool)
	}
}

func (_ *CljsCoreRSeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreRSeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(List, "EMPTY"), self__.Meta).(interface{})
	}
}

func (_ *CljsCoreRSeq) CljsCoreIReduce__() {}
func (self__ *CljsCoreRSeq) X_reduce_Arity2(f interface{}) interface{} {
	{
		var col___1 = self__
		_ = col___1
		return Seq_reduce.X_invoke_Arity2(f, col___1).(interface{})
	}
}

func (self__ *CljsCoreRSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var col___1 = self__
		_ = col___1
		return Seq_reduce.X_invoke_Arity3(f, start, col___1).(interface{})
	}
}

func (_ *CljsCoreRSeq) CljsCoreISeq__() {}
func (self__ *CljsCoreRSeq) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Ci.(CljsCoreIIndexed).X_nth_Arity2(self__.I).(interface{})
	}
}

func (self__ *CljsCoreRSeq) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.I.(float64) > float64(0) {
			return (&CljsCoreRSeq{self__.Ci, (self__.I.(float64) - float64(1)), nil})
		} else {
			return CljsCoreList_EMPTY
		}
	}
}

func (_ *CljsCoreRSeq) CljsCoreISeqable__() {}
func (self__ *CljsCoreRSeq) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1
	}
}

func (_ *CljsCoreRSeq) CljsCoreISequential__() {}
func (_ *CljsCoreRSeq) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreRSeq) X_with_meta_Arity2(new_meta interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreRSeq{self__.Ci, self__.I, new_meta})
	}
}

func (_ *CljsCoreRSeq) CljsCoreICollection__() {}
func (self__ *CljsCoreRSeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(interface{})
	}
}

var X__GT_RSeq = func(__GT_RSeq *AFn) *AFn {
	return Fn(__GT_RSeq, func(ci interface{}, i interface{}, meta interface{}) interface{} {
		return (&CljsCoreRSeq{ci, i, meta})
	})
}(&AFn{})

/**
* Same as (first (next x))
 */
var Second = func(second *AFn) *AFn {
	return Fn(second, func(coll interface{}) interface{} {
		return First.X_invoke_Arity1(Next.Arity1IQ(coll))
	})
}(&AFn{})

/**
* Same as (first (first x))
 */
var Ffirst = func(ffirst *AFn) *AFn {
	return Fn(ffirst, func(coll interface{}) interface{} {
		return First.X_invoke_Arity1(First.X_invoke_Arity1(coll))
	})
}(&AFn{})

/**
* Same as (next (first x))
 */
var Nfirst = func(nfirst *AFn) *AFn {
	return Fn(nfirst, func(coll interface{}) interface{} {
		return Next.X_invoke_Arity1(First.X_invoke_Arity1(coll))
	})
}(&AFn{})

/**
* Same as (first (next x))
 */
var Fnext = func(fnext *AFn) *AFn {
	return Fn(fnext, func(coll interface{}) interface{} {
		return First.X_invoke_Arity1(Next.Arity1IQ(coll))
	})
}(&AFn{})

/**
* Same as (next (next x))
 */
var Nnext = func(nnext *AFn) *AFn {
	return Fn(nnext, func(coll interface{}) interface{} {
		return Next.X_invoke_Arity1(Next.Arity1IQ(coll))
	})
}(&AFn{})

/**
* Return the last item in coll, in linear time
 */
var Last = func(last *AFn) *AFn {
	return Fn(last, func(s interface{}) interface{} {
		for {
			{
				var sn = Next.Arity1IQ(s)
				_ = sn
				if !(sn == nil) {
					{
						var G__325 = sn
						s = G__325
						continue
					}
				} else {
					return First.X_invoke_Arity1(s)
				}
			}
		}
	})
}(&AFn{})

/**
* conj[oin]. Returns a new collection with the xs
* 'added'. (conj nil item) returns (item).  The 'addition' may
* happen at different 'places' depending on the concrete type.
* @param {...*} var_args
 */
var Conj = func(conj *AFn) *AFn {
	return Fn(conj, func() interface{} {
		return CljsCorePersistentVector_EMPTY
	}, func(coll interface{}) interface{} {
		return coll
	}, func(coll interface{}, x interface{}) interface{} {
		if !(coll == nil) {
			return coll.(CljsCoreICollection).X_conj_Arity2(x)
		} else {
			return Native_get_instance_field.X_invoke_Arity2(List, "EMPTY").(CljsCoreICollection).X_conj_Arity2(x)
		}
	}, func(coll_x_xs__ ...interface{}) interface{} {
		var coll = coll_x_xs__[0]
		var x = coll_x_xs__[1]
		var xs = Array_seq.X_invoke_Arity1(coll_x_xs__[2:])
		_, _, _ = coll, x, xs
		for {
			if Truth_(xs) {
				{
					var G__326 = conj.X_invoke_Arity2(coll, x).(interface{})
					var G__327 = First.X_invoke_Arity1(xs)
					var G__328 = Next.Arity1IQ(xs)
					coll = G__326
					x = G__327
					xs = G__328
					continue
				}
			} else {
				return conj.X_invoke_Arity2(coll, x).(interface{})
			}
		}
	})
}(&AFn{})

/**
* Returns an empty collection of the same category as coll, or nil
 */
var Empty = func(empty *AFn) *AFn {
	return Fn(empty, func(coll interface{}) interface{} {
		if coll == nil {
			return nil
		} else {
			return coll.(CljsCoreIEmptyableCollection).X_empty_Arity1().(interface{})
		}
	})
}(&AFn{})

var Accumulating_seq_count = func(accumulating_seq_count *AFn) *AFn {
	return Fn(accumulating_seq_count, func(coll interface{}) interface{} {
		{
			var s = Seq.Arity1IQ(coll)
			var acc = float64(0)
			_, _ = s, acc
			for {
				if Counted_QMARK_.X_invoke_Arity1(s) {
					return (acc + s.X_count_Arity1())
				} else {
					{
						var G__329 = Next.X_invoke_Arity1(s)
						var G__330 = (acc + float64(1))
						s = G__329
						acc = G__330
						continue
					}
				}
			}
		}
	})
}(&AFn{})

/**
* Returns the number of items in the collection. (count nil) returns
* 0.  Also works on strings, arrays, and Maps
 */
var Count = func(count *AFn) *AFn {
	return Fn(count, func(coll interface{}) interface{} {
		if !(coll == nil) {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ICounted", Str: "cljs.core/ICounted", X_hash: float64(-1299011378), X_meta: nil}), coll)) {
				return coll.X_count_Arity1()
			} else {
				if reflect.TypeOf(coll).Kind() == reflect.Slice {
					return float64(len(coll.([]interface{})))
				} else {
					if reflect.TypeOf(coll).Kind() == reflect.String {
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

var Linear_traversal_nth = func(linear_traversal_nth *AFn) *AFn {
	return Fn(linear_traversal_nth, func(coll interface{}, n interface{}) interface{} {
		for {
			if coll == nil {
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
						return coll.(CljsCoreIIndexed).X_nth_Arity2(n).(interface{})
					} else {
						if Truth_(Seq.Arity1IQ(coll)) {
							{
								var G__331 = Next.Arity1IQ(coll)
								var G__332 = (n.(float64) - float64(1))
								coll = G__331
								n = G__332
								continue
							}
						} else {
							panic((&js.Error{"Index out of bounds"}))

						}
					}
				}
			}
		}
	}, func(coll interface{}, n interface{}, not_found interface{}) interface{} {
		for {
			if coll == nil {
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
						return coll.(CljsCoreIIndexed).X_nth_Arity3(n, not_found).(interface{})
					} else {
						if Truth_(Seq.Arity1IQ(coll)) {
							{
								var G__333 = Next.Arity1IQ(coll)
								var G__334 = (n.(float64) - float64(1))
								var G__335 = not_found
								coll = G__333
								n = G__334
								not_found = G__335
								continue
							}
						} else {
							return not_found

						}
					}
				}
			}
		}
	})
}(&AFn{})

/**
* Returns the value at the index. get returns nil if index out of
* bounds, nth throws an exception unless not-found is supplied.  nth
* also works for strings, arrays, regex Matchers and Lists, and,
* in O(n) time, for sequences.
 */
var Nth = func(nth *AFn) *AFn {
	return Fn(nth, func(coll interface{}, n interface{}) interface{} {
		if !(reflect.TypeOf(n).Kind() == reflect.Float64) {
			panic((&js.Error{"index argument to nth must be a number"}))
		} else {
			if coll == nil {
				return coll
			} else {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IIndexed", Str: "cljs.core/IIndexed", X_hash: float64(-436490749), X_meta: nil}), coll)) {
					return coll.X_nth_Arity2(n).(interface{})
				} else {
					if reflect.TypeOf(coll).Kind() == reflect.Slice {
						if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
							return (coll.([]interface{})[int(n.(float64))])
						} else {
							return nil
						}
					} else {
						if reflect.TypeOf(coll).Kind() == reflect.String {
							if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
								return (coll.([]interface{})[int(n.(float64))])
							} else {
								return nil
							}
						} else {
							if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IIndexed", Str: "cljs.core/IIndexed", X_hash: float64(-436490749), X_meta: nil}), coll)) {
								return coll.(CljsCoreIIndexed).X_nth_Arity2(n).(interface{})
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
		if !(reflect.TypeOf(n).Kind() == reflect.Float64) {
			panic((&js.Error{"index argument to nth must be a number."}))
		} else {
			if coll == nil {
				return not_found
			} else {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IIndexed", Str: "cljs.core/IIndexed", X_hash: float64(-436490749), X_meta: nil}), coll)) {
					return coll.X_nth_Arity3(n, not_found).(interface{})
				} else {
					if reflect.TypeOf(coll).Kind() == reflect.Slice {
						if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
							return (coll.([]interface{})[int(n.(float64))])
						} else {
							return not_found
						}
					} else {
						if reflect.TypeOf(coll).Kind() == reflect.String {
							if n.(float64) < Native_get_instance_field.X_invoke_Arity2(coll, "Length").(float64) {
								return (coll.([]interface{})[int(n.(float64))])
							} else {
								return not_found
							}
						} else {
							if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IIndexed", Str: "cljs.core/IIndexed", X_hash: float64(-436490749), X_meta: nil}), coll)) {
								return coll.(CljsCoreIIndexed).X_nth_Arity2(n).(interface{})
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

/**
* Returns the value mapped to key, not-found or nil if key not present.
 */
var Get = func(get *AFn) *AFn {
	return Fn(get, func(o interface{}, k interface{}) interface{} {
		if o == nil {
			return nil
		} else {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ILookup", Str: "cljs.core/ILookup", X_hash: float64(-150575073), X_meta: nil}), o)) {
				return o.X_lookup_Arity2(k).(interface{})
			} else {
				if reflect.TypeOf(o).Kind() == reflect.Slice {
					if k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64) {
						return (o.([]interface{})[int(k.(float64))])
					} else {
						return nil
					}
				} else {
					if reflect.TypeOf(o).Kind() == reflect.String {
						if k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64) {
							return (o.([]interface{})[int(k.(float64))])
						} else {
							return nil
						}
					} else {
						if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ILookup", Str: "cljs.core/ILookup", X_hash: float64(-150575073), X_meta: nil}), o)) {
							return o.(CljsCoreILookup).X_lookup_Arity2(k).(interface{})
						} else {
							return nil

						}
					}
				}
			}
		}
	}, func(o interface{}, k interface{}, not_found interface{}) interface{} {
		if !(o == nil) {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ILookup", Str: "cljs.core/ILookup", X_hash: float64(-150575073), X_meta: nil}), o)) {
				return o.X_lookup_Arity3(k, not_found).(interface{})
			} else {
				if reflect.TypeOf(o).Kind() == reflect.Slice {
					if k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64) {
						return (o.([]interface{})[int(k.(float64))])
					} else {
						return not_found
					}
				} else {
					if reflect.TypeOf(o).Kind() == reflect.String {
						if k.(float64) < Native_get_instance_field.X_invoke_Arity2(o, "Length").(float64) {
							return (o.([]interface{})[int(k.(float64))])
						} else {
							return not_found
						}
					} else {
						if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ILookup", Str: "cljs.core/ILookup", X_hash: float64(-150575073), X_meta: nil}), o)) {
							return o.(CljsCoreILookup).X_lookup_Arity3(k, not_found).(interface{})
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

/**
* assoc[iate]. When applied to a map, returns a new map of the
* same (hashed/sorted) type, that contains the mapping of key(s) to
* val(s). When applied to a vector, returns a new vector that
* contains val at index.
* @param {...*} var_args
 */
var Assoc = func(assoc *AFn) *AFn {
	return Fn(assoc, func(coll interface{}, k interface{}, v interface{}) interface{} {
		if !(coll == nil) {
			return coll.(CljsCoreIAssociative).X_assoc_Arity3(k, v)
		} else {
			return Native_invoke_instance_method.X_invoke_Arity3(PersistentHashMap, "FromArrays", []interface{}{[]interface{}{k}, []interface{}{v}})
		}
	}, func(coll_k_v_kvs__ ...interface{}) interface{} {
		var coll = coll_k_v_kvs__[0]
		var k = coll_k_v_kvs__[1]
		var v = coll_k_v_kvs__[2]
		var kvs = Array_seq.X_invoke_Arity1(coll_k_v_kvs__[3:])
		_, _, _, _ = coll, k, v, kvs
		for {
			{
				var ret = assoc.X_invoke_Arity3(coll, k, v).(interface{})
				_ = ret
				if Truth_(kvs) {
					{
						var G__336 = ret
						var G__337 = First.X_invoke_Arity1(kvs)
						var G__338 = Second.X_invoke_Arity1(kvs)
						var G__339 = Nnext.X_invoke_Arity1(kvs).(CljsCoreISeq)
						coll = G__336
						k = G__337
						v = G__338
						kvs = G__339
						continue
					}
				} else {
					return ret
				}
			}
		}
	})
}(&AFn{})

/**
* dissoc[iate]. Returns a new map of the same (hashed/sorted) type,
* that does not contain a mapping for key(s).
* @param {...*} var_args
 */
var Dissoc = func(dissoc *AFn) *AFn {
	return Fn(dissoc, func(coll interface{}) interface{} {
		return coll
	}, func(coll interface{}, k interface{}) interface{} {
		if coll == nil {
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
			if coll == nil {
				return nil
			} else {
				{
					var ret = dissoc.X_invoke_Arity2(coll, k)
					_ = ret
					if Truth_(ks) {
						{
							var G__340 = ret
							var G__341 = First.X_invoke_Arity1(ks)
							var G__342 = Next.Arity1IQ(ks)
							coll = G__340
							k = G__341
							ks = G__342
							continue
						}
					} else {
						return ret
					}
				}
			}
		}
	})
}(&AFn{})

var Fn_QMARK_ = func(fn_QMARK_ *AFn) *AFn {
	return Fn(fn_QMARK_, func(f interface{}) bool {
		{
			var or__76632__auto__ = goog.IsFunction(f)
			_ = or__76632__auto__
			if or__76632__auto__ {
				return or__76632__auto__
			} else {
				return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "Fn", Str: "cljs.core/Fn", X_hash: float64(-695281833), X_meta: nil}), f)
			}
		}
	})
}(&AFn{})

type CljsCoreMetaFn struct {
	Afn  interface{}
	Meta interface{}
}

func (_ *CljsCoreMetaFn) CljsCoreIFn__() {}
func (self__ *CljsCoreMetaFn) X_invoke_Arity0() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity0().(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity1(a interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity1(a).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity2(a interface{}, b interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity2(a, b).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity3(a, b, c).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity4(a, b, c, d).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity5(a, b, c, d, e).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity6(a, b, c, d, e, f).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity7(a, b, c, d, e, f, g).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity8(a, b, c, d, e, f, g, h).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity9(a, b, c, d, e, f, g, h, i).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity10(a, b, c, d, e, f, g, h, i, j).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity11(a, b, c, d, e, f, g, h, i, j, k).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity12(a, b, c, d, e, f, g, h, i, j, k, l).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity13(a, b, c, d, e, f, g, h, i, j, k, l, m).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity14(a, b, c, d, e, f, g, h, i, j, k, l, m, n).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity15(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity16(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity17(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity18(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity19(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s).(interface{})
	}
}

func (self__ *CljsCoreMetaFn) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Afn.(CljsCoreIFn).X_invoke_Arity20(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t).(interface{})
	}
}

func (_ *CljsCoreMetaFn) CljsCoreFn__()        {}
func (_ *CljsCoreMetaFn) CljsCoreIWithMeta__() {}
func (self__ *CljsCoreMetaFn) X_with_meta_Arity2(new_meta interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCoreMetaFn{self__.Afn, new_meta})
	}
}

func (_ *CljsCoreMetaFn) CljsCoreIMeta__() {}
func (self__ *CljsCoreMetaFn) X_meta_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Meta
	}
}

var X__GT_MetaFn = func(__GT_MetaFn *AFn) *AFn {
	return Fn(__GT_MetaFn, func(afn interface{}, meta interface{}) interface{} {
		return (&CljsCoreMetaFn{afn, meta})
	})
}(&AFn{})

/**
* Returns an object of the same type and value as obj, with
* map m as its metadata.
 */
var With_meta = func(with_meta *AFn) *AFn {
	return Fn(with_meta, func(o interface{}, meta interface{}) interface{} {
		if (Fn_QMARK_.Arity1IB(o)) && (Not.Arity1IB(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IWithMeta", Str: "cljs.core/IWithMeta", X_hash: float64(-1981666051), X_meta: nil}), o))) {
			return (&CljsCoreMetaFn{o, meta})
		} else {
			if o == nil {
				return nil
			} else {
				return o.(CljsCoreIWithMeta).X_with_meta_Arity2(meta)
			}
		}
	})
}(&AFn{})

/**
* Returns the metadata of obj, returns nil if there is no metadata.
 */
var Meta = func(meta *AFn) *AFn {
	return Fn(meta, func(o interface{}) interface{} {
		if Truth_(func() interface{} {
			var and__76620__auto__ = !(o == nil)
			_ = and__76620__auto__
			if and__76620__auto__ {
				return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IMeta", Str: "cljs.core/IMeta", X_hash: float64(-1459057517), X_meta: nil}), o)
			} else {
				return and__76620__auto__
			}
		}()) {
			return o.(CljsCoreIMeta).X_meta_Arity1()
		} else {
			return nil
		}
	})
}(&AFn{})

/**
* For a list or queue, same as first, for a vector, same as, but much
* more efficient than, last. If the collection is empty, returns nil.
 */
var Peek = func(peek *AFn) *AFn {
	return Fn(peek, func(coll interface{}) interface{} {
		if coll == nil {
			return nil
		} else {
			return coll.(CljsCoreIStack).X_peek_Arity1().(interface{})
		}
	})
}(&AFn{})

/**
* For a list or queue, returns a new list/queue without the first
* item, for a vector, returns a new vector without the last item.
* Note - not the same as next/butlast.
 */
var Pop = func(pop *AFn) *AFn {
	return Fn(pop, func(coll interface{}) interface{} {
		if coll == nil {
			return nil
		} else {
			return coll.(CljsCoreIStack).X_pop_Arity1()
		}
	})
}(&AFn{})

/**
* disj[oin]. Returns a new set of the same (hashed/sorted) type, that
* does not contain key(s).
* @param {...*} var_args
 */
var Disj = func(disj *AFn) *AFn {
	return Fn(disj, func(coll interface{}) interface{} {
		return coll
	}, func(coll interface{}, k interface{}) interface{} {
		if coll == nil {
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
			if coll == nil {
				return nil
			} else {
				{
					var ret = disj.X_invoke_Arity2(coll, k)
					_ = ret
					if Truth_(ks) {
						{
							var G__343 = ret
							var G__344 = First.X_invoke_Arity1(ks)
							var G__345 = Next.Arity1IQ(ks)
							coll = G__343
							k = G__344
							ks = G__345
							continue
						}
					} else {
						return ret
					}
				}
			}
		}
	})
}(&AFn{})

/**
* Returns true if coll has no items - same as (not (seq coll)).
* Please use the idiom (seq x) rather than (not (empty? x))
 */
var Empty_QMARK_ = func(empty_QMARK_ *AFn) *AFn {
	return Fn(empty_QMARK_, func(coll interface{}) bool {
		return (coll == nil) || (Not.X_invoke_Arity1(Seq.Arity1IQ(coll)))
	})
}(&AFn{})

/**
* Returns true if x satisfies ICollection
 */
var Coll_QMARK_ = func(coll_QMARK_ *AFn) *AFn {
	return Fn(coll_QMARK_, func(x interface{}) bool {
		if x == nil {
			return false
		} else {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ICollection", Str: "cljs.core/ICollection", X_hash: float64(802638471), X_meta: nil}), x)
		}
	})
}(&AFn{})

/**
* Returns true if x satisfies ISet
 */
var Set_QMARK_ = func(set_QMARK_ *AFn) *AFn {
	return Fn(set_QMARK_, func(x interface{}) bool {
		if x == nil {
			return false
		} else {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISet", Str: "cljs.core/ISet", X_hash: float64(2003412810), X_meta: nil}), x)
		}
	})
}(&AFn{})

/**
* Returns true if coll implements Associative
 */
var Associative_QMARK_ = func(associative_QMARK_ *AFn) *AFn {
	return Fn(associative_QMARK_, func(x interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IAssociative", Str: "cljs.core/IAssociative", X_hash: float64(-1700920611), X_meta: nil}), x).(bool)
	})
}(&AFn{})

/**
* Returns true if coll satisfies ISequential
 */
var Sequential_QMARK_ = func(sequential_QMARK_ *AFn) *AFn {
	return Fn(sequential_QMARK_, func(x interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISequential", Str: "cljs.core/ISequential", X_hash: float64(-950981796), X_meta: nil}), x).(bool)
	})
}(&AFn{})

/**
* Returns true if coll satisfies ISorted
 */
var Sorted_QMARK_ = func(sorted_QMARK_ *AFn) *AFn {
	return Fn(sorted_QMARK_, func(x interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISorted", Str: "cljs.core/ISorted", X_hash: float64(-1734125647), X_meta: nil}), x).(bool)
	})
}(&AFn{})

/**
* Returns true if coll satisfies IReduce
 */
var Reduceable_QMARK_ = func(reduceable_QMARK_ *AFn) *AFn {
	return Fn(reduceable_QMARK_, func(x interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), x).(bool)
	})
}(&AFn{})

/**
* Return true if x satisfies IMap
 */
var Map_QMARK_ = func(map_QMARK_ *AFn) *AFn {
	return Fn(map_QMARK_, func(x interface{}) bool {
		if x == nil {
			return false
		} else {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IMap", Str: "cljs.core/IMap", X_hash: float64(1407777598), X_meta: nil}), x)
		}
	})
}(&AFn{})

/**
* Return true if x satisfies IVector
 */
var Vector_QMARK_ = func(vector_QMARK_ *AFn) *AFn {
	return Fn(vector_QMARK_, func(x interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IVector", Str: "cljs.core/IVector", X_hash: float64(1711112835), X_meta: nil}), x).(bool)
	})
}(&AFn{})

var Chunked_seq_QMARK_ = func(chunked_seq_QMARK_ *AFn) *AFn {
	return Fn(chunked_seq_QMARK_, func(x interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IChunkedSeq", Str: "cljs.core/IChunkedSeq", X_hash: float64(-892943716), X_meta: nil}), x).(bool)
	})
}(&AFn{})

/**
* @param {...*} var_args
 */
var Js_obj = func(js_obj *AFn) *AFn {
	return Fn(js_obj, func() interface{} {
		{
			var obj349 = map[string]interface{}{}
			_ = obj349
			return obj349
		}
	}, func(keyvals__ ...interface{}) interface{} {
		var keyvals = Array_seq.X_invoke_Arity1(keyvals__[0:])
		_ = keyvals
		return Apply.X_invoke_Arity2(goog_object.Create, keyvals).(interface{})
	})
}(&AFn{})

var Js_keys = func(js_keys *AFn) *AFn {
	return Fn(js_keys, func(obj interface{}) interface{} {
		{
			var keys = []interface{}{}
			_ = keys
			goog_object.ForEach(obj, func(keys []interface{}) *AFn {
				return func(G__350 *AFn) *AFn {
					return Fn(G__350, func(val interface{}, key interface{}, obj___1 interface{}) interface{} {
						return keys.Push(key)
					})
				}(&AFn{})
			}(keys))
			return keys
		}
	})
}(&AFn{})

var Js_delete = func(js_delete *AFn) *AFn {
	return Fn(js_delete, func(obj interface{}, key interface{}) interface{} {
		return delete(obj, key)
	})
}(&AFn{})

var Array_copy = func(array_copy *AFn) *AFn {
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
					{
						var G__351 = (i___1.(float64) + float64(1))
						var G__352 = (j___1.(float64) + float64(1))
						var G__353 = (len___1.(float64) - float64(1))
						i___1 = G__351
						j___1 = G__352
						len___1 = G__353
						continue
					}
				}
			}
		}
	})
}(&AFn{})

var Array_copy_downward = func(array_copy_downward *AFn) *AFn {
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
					{
						var G__354 = (i___1 - float64(1))
						var G__355 = (j___1 - float64(1))
						var G__356 = (len___1.(float64) - float64(1))
						i___1 = G__354
						j___1 = G__355
						len___1 = G__356
						continue
					}
				}
			}
		}
	})
}(&AFn{})

var Lookup_sentinel = func() interface{} {
	var obj358 = map[string]interface{}{}
	_ = obj358
	return obj358
}()

/**
* Returns true if x is the value false, false otherwise.
 */
var False_QMARK_ = func(false_QMARK_ *AFn) *AFn {
	return Fn(false_QMARK_, func(x interface{}) bool {
		return x == false
	})
}(&AFn{})

/**
* Returns true if x is the value true, false otherwise.
 */
var True_QMARK_ = func(true_QMARK_ *AFn) *AFn {
	return Fn(true_QMARK_, func(x interface{}) bool {
		return x == true
	})
}(&AFn{})

var Undefined_QMARK_ = func(undefined_QMARK_ *AFn) *AFn {
	return Fn(undefined_QMARK_, func(x interface{}) bool {
		return (nil == x)
	})
}(&AFn{})

/**
* Return true if s satisfies ISeq
 */
var Seq_QMARK_ = func(seq_QMARK_ *AFn) *AFn {
	return Fn(seq_QMARK_, func(s interface{}) bool {
		if s == nil {
			return false
		} else {
			return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeq", Str: "cljs.core/ISeq", X_hash: float64(230133392), X_meta: nil}), s)
		}
	})
}(&AFn{})

/**
* Return true if s satisfies ISeqable
 */
var Seqable_QMARK_ = func(seqable_QMARK_ *AFn) *AFn {
	return Fn(seqable_QMARK_, func(s interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "ISeqable", Str: "cljs.core/ISeqable", X_hash: float64(137437203), X_meta: nil}), s).(bool)
	})
}(&AFn{})

var Boolean = func(boolean *AFn) *AFn {
	return Fn(boolean, func(x interface{}) bool {
		if Truth_(x) {
			return true
		} else {
			return false
		}
	})
}(&AFn{})

var Ifn_QMARK_ = func(ifn_QMARK_ *AFn) *AFn {
	return Fn(ifn_QMARK_, func(f interface{}) bool {
		{
			var or__76632__auto__ = Fn_QMARK_.Arity1IB(f)
			_ = or__76632__auto__
			if or__76632__auto__ {
				return or__76632__auto__
			} else {
				return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IFn", Str: "cljs.core/IFn", X_hash: float64(-920223129), X_meta: nil}), f)
			}
		}
	})
}(&AFn{})

/**
* Returns true if n is an integer.
 */
var Integer_QMARK_ = func(integer_QMARK_ *AFn) *AFn {
	return Fn(integer_QMARK_, func(n interface{}) bool {
		return (reflect.TypeOf(n).Kind() == reflect.Float64) && (!(js.IsNaN(n))) && (!(n == js.Infinity)) && (js.ParseFloat(n).(float64) == js.ParseInt(n, float64(10)).(float64))
	})
}(&AFn{})

/**
* Returns true if key is present in the given collection, otherwise
* returns false.  Note that for numerically indexed collections like
* vectors and arrays, this tests if the numeric key is within the
* range of indexes. 'contains?' operates constant or logarithmic time;
* it will not perform a linear search for a value.  See also 'some'.
 */
var Contains_QMARK_ = func(contains_QMARK_ *AFn) *AFn {
	return Fn(contains_QMARK_, func(coll interface{}, v interface{}) bool {
		if Get.X_invoke_Arity3(coll, v, Lookup_sentinel) == Lookup_sentinel {
			return false
		} else {
			return true
		}
	})
}(&AFn{})

/**
* Returns the map entry for key, or nil if key not present.
 */
var Find = func(find *AFn) *AFn {
	return Fn(find, func(coll interface{}, k interface{}) interface{} {
		if (!(coll == nil)) && (Associative_QMARK_.Arity1IB(coll)) && (Contains_QMARK_.Arity2IIB(coll, k)) {
			return &CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{k, Get.X_invoke_Arity2(coll, k)}, nil}
		} else {
			return nil
		}
	})
}(&AFn{})

/**
* Returns true if no two of the arguments are =
* @param {...*} var_args
 */
var Distinct_QMARK_ = func(distinct_QMARK_ *AFn) *AFn {
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
							if Contains_QMARK_.X_invoke_Arity2(s, x___1) {
								return false
							} else {
								{
									var G__359 = Conj.X_invoke_Arity2(s, x___1).(interface{})
									var G__360 = etc
									s = G__359
									xs = G__360
									continue
								}
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

/**
* Coerces coll to a (possibly empty) sequence, if it is not already
* one. Will not force a lazy seq. (sequence nil) yields ()
 */
var Sequence = func(sequence *AFn) *AFn {
	return Fn(sequence, func(coll interface{}) CljsCoreISeq {
		if Seq_QMARK_.Arity1IB(coll) {
			return coll
		} else {
			{
				var or__76632__auto__ = Seq.Arity1IQ(coll)
				_ = or__76632__auto__
				if Truth_(or__76632__auto__) {
					return or__76632__auto__
				} else {
					return CljsCoreList_EMPTY
				}
			}
		}
	})
}(&AFn{})

/**
* Comparator. Returns a negative number, zero, or a positive number
* when x is logically 'less than', 'equal to', or 'greater than'
* y. Uses IComparable if available and google.array.defaultCompare for objects
* of the same type and special-cases nil to be less than any other object.
 */
var Compare = func(compare *AFn) *AFn {
	return Fn(compare, func(x interface{}, y interface{}) float64 {
		if x == y {
			return float64(0)
		} else {
			if x == nil {
				return float64(-1)
			} else {
				if y == nil {
					return float64(1)
				} else {
					if Type_.X_invoke_Arity1(x) == Type_.X_invoke_Arity1(y) {
						if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IComparable", Str: "cljs.core/IComparable", X_hash: float64(1166626940), X_meta: nil}), x)) {
							return x.X_compare_Arity2(y)
						} else {
							return goog_array.DefaultCompare(x, y)
						}
					} else {
						panic((&js.Error{"compare on non-nil objects of different types"}))

					}
				}
			}
		}
	})
}(&AFn{})

/**
* Compare indexed collection.
 */
var Compare_indexed = func(compare_indexed *AFn) *AFn {
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
				var d = Compare.X_invoke_Arity2(Nth.X_invoke_Arity2(xs, n), Nth.X_invoke_Arity2(ys, n))
				_ = d
				if (d == float64(0)) && ((n.(float64) + float64(1)) < len.(float64)) {
					{
						var G__361 = xs
						var G__362 = ys
						var G__363 = len
						var G__364 = (n.(float64) + float64(1))
						xs = G__361
						ys = G__362
						len = G__363
						n = G__364
						continue
					}
				} else {
					return d
				}
			}
		}
	})
}(&AFn{})

/**
* Given a fn that might be boolean valued or a comparator,
* return a fn that is a comparator.
 */
var Fn__GT_comparator = func(fn__GT_comparator *AFn) *AFn {
	return Fn(fn__GT_comparator, func(f interface{}) interface{} {
		if X_EQ_.X_invoke_Arity2(f, Compare) {
			return Compare
		} else {
			return func(G__365 *AFn) *AFn {
				return Fn(G__365, func(x interface{}, y interface{}) interface{} {
					{
						var r = f.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{})
						_ = r
						if reflect.TypeOf(r).Kind() == reflect.Float64 {
							return r
						} else {
							if Truth_(r) {
								return float64(-1)
							} else {
								if Truth_(f.(CljsCoreIFn).X_invoke_Arity2(y, x).(interface{})) {
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

/**
* Returns a sorted sequence of the items in coll. Comp can be
* boolean-valued comparison funcion, or a -/0/+ valued comparator.
* Comp defaults to compare.
 */
var Sort = func(sort *AFn) *AFn {
	return Fn(sort, func(coll interface{}) interface{} {
		return sort.X_invoke_Arity2(Compare, coll)
	}, func(comp interface{}, coll interface{}) interface{} {
		if Truth_(Seq.Arity1IQ(coll)) {
			{
				var a = To_array.X_invoke_Arity1(coll).(interface{})
				_ = a
				goog_array.StableSort(a, Fn__GT_comparator.X_invoke_Arity1(comp))
				return Seq.X_invoke_Arity1(a)
			}
		} else {
			return CljsCoreList_EMPTY
		}
	})
}(&AFn{})

/**
* Returns a sorted sequence of the items in coll, where the sort
* order is determined by comparing (keyfn item).  Comp can be
* boolean-valued comparison funcion, or a -/0/+ valued comparator.
* Comp defaults to compare.
 */
var Sort_by = func(sort_by *AFn) *AFn {
	return Fn(sort_by, func(keyfn interface{}, coll interface{}) interface{} {
		return sort_by.X_invoke_Arity3(keyfn, Compare, coll)
	}, func(keyfn interface{}, comp interface{}, coll interface{}) interface{} {
		return Sort.X_invoke_Arity2(func(G__366 *AFn) *AFn {
			return Fn(G__366, func(x interface{}, y interface{}) interface{} {
				return Fn__GT_comparator.X_invoke_Arity1(comp).X_invoke_Arity2(keyfn.(CljsCoreIFn).X_invoke_Arity1(x).(interface{}), keyfn.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})).(interface{})
			})
		}(&AFn{}), coll)
	})
}(&AFn{})

var Seq_reduce = func(seq_reduce *AFn) *AFn {
	return Fn(seq_reduce, func(f interface{}, coll interface{}) interface{} {
		{
			var temp__4217__auto__ = Seq.Arity1IQ(coll)
			_ = temp__4217__auto__
			if Truth_(temp__4217__auto__) {
				{
					var s = temp__4217__auto__
					_ = s
					return Reduce.X_invoke_Arity3(f, First.X_invoke_Arity1(s), Next.X_invoke_Arity1(s)).(interface{})
				}
			} else {
				return f.(CljsCoreIFn).X_invoke_Arity0().(interface{})
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
						var nval = f.(CljsCoreIFn).X_invoke_Arity2(val___1, First.X_invoke_Arity1(coll___1)).(interface{})
						_ = nval
						if Reduced_QMARK_.X_invoke_Arity1(nval) {
							return Deref.X_invoke_Arity1(nval).(interface{})
						} else {
							{
								var G__367 = nval
								var G__368 = Next.X_invoke_Arity1(coll___1)
								val___1 = G__367
								coll___1 = G__368
								continue
							}
						}
					}
				} else {
					return val___1
				}
			}
		}
	})
}(&AFn{})

/**
* Return a random permutation of coll
 */
var Shuffle = func(shuffle *AFn) *AFn {
	return Fn(shuffle, func(coll interface{}) interface{} {
		{
			var a = To_array.X_invoke_Arity1(coll).(interface{})
			_ = a
			goog_array.Shuffle(a)
			return Vec.X_invoke_Arity1(a).(interface{})
		}
	})
}(&AFn{})

/**
* f should be a function of 2 arguments. If val is not supplied,
* returns the result of applying f to the first 2 items in coll, then
* applying f to that result and the 3rd item, etc. If coll contains no
* items, f must accept no arguments as well, and reduce returns the
* result of calling f with no arguments.  If coll has only 1 item, it
* is returned and f is not called.  If val is supplied, returns the
* result of applying f to val and the first item in coll, then
* applying f to that result and the 2nd item, etc. If coll contains no
* items, returns val and f is not called.
 */
var Reduce = func(reduce *AFn) *AFn {
	return Fn(reduce, func(f interface{}, coll interface{}) interface{} {
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), coll)) {
			return coll.X_reduce_Arity2(f).(interface{})
		} else {
			if reflect.TypeOf(coll).Kind() == reflect.Slice {
				return Array_reduce.X_invoke_Arity2(coll, f)
			} else {
				if reflect.TypeOf(coll).Kind() == reflect.String {
					return Array_reduce.X_invoke_Arity2(coll, f)
				} else {
					if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), coll)) {
						return coll.(CljsCoreIReduce).X_reduce_Arity2(f).(interface{})
					} else {
						return Seq_reduce.X_invoke_Arity2(f, coll).(interface{})

					}
				}
			}
		}
	}, func(f interface{}, val interface{}, coll interface{}) interface{} {
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), coll)) {
			return coll.X_reduce_Arity3(f, val).(interface{})
		} else {
			if reflect.TypeOf(coll).Kind() == reflect.Slice {
				return Array_reduce.X_invoke_Arity3(coll, f, val)
			} else {
				if reflect.TypeOf(coll).Kind() == reflect.String {
					return Array_reduce.X_invoke_Arity3(coll, f, val)
				} else {
					if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), coll)) {
						return coll.(CljsCoreIReduce).X_reduce_Arity3(f, val).(interface{})
					} else {
						return Seq_reduce.X_invoke_Arity3(f, val, coll)

					}
				}
			}
		}
	})
}(&AFn{})

/**
* Reduces an associative collection. f should be a function of 3
* arguments. Returns the result of applying f to init, the first key
* and the first value in coll, then applying f to that result and the
* 2nd key and value, etc. If coll contains no entries, returns init
* and f is not called. Note that reduce-kv is supported on vectors,
* where the keys will be the ordinals.
 */
var Reduce_kv = func(reduce_kv *AFn) *AFn {
	return Fn(reduce_kv, func(f interface{}, init interface{}, coll interface{}) interface{} {
		if !(coll == nil) {
			return coll.(CljsCoreIKVReduce).X_kv_reduce_Arity3(f, init).(interface{})
		} else {
			return init
		}
	})
}(&AFn{})

var Completing = func(completing *AFn) *AFn {
	return Fn(completing, func(f interface{}) interface{} {
		return func(G__369 *AFn) *AFn {
			return Fn(G__369, func() interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity0().(interface{})
			}, func(x interface{}) interface{} {
				return x
			}, func(x interface{}, y interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{})
			})
		}(&AFn{})
	})
}(&AFn{})

/**
* reduce with a transformation of f (xf). If init is not
* supplied, (f) will be called to produce it. Returns the result of
* applying (the transformed) xf to init and the first item in coll,
* then applying xf to that result and the 2nd item, etc. If coll
* contains no items, returns init and f is not called. Note that
* certain transforms may inject or skip items.
 */
var Transduce = func(transduce *AFn) *AFn {
	return Fn(transduce, func(xform interface{}, f interface{}, coll interface{}) interface{} {
		return transduce.X_invoke_Arity4(xform, f, f.(CljsCoreIFn).X_invoke_Arity0().(interface{}), coll).(interface{})
	}, func(xform interface{}, f interface{}, init interface{}, coll interface{}) interface{} {
		{
			var f___1 = xform.(CljsCoreIFn).X_invoke_Arity1(Completing.X_invoke_Arity1(f).(interface{})).(interface{})
			var ret = Reduce.X_invoke_Arity3(f___1, init, coll)
			var ret___1 = f___1.(CljsCoreIFn).X_invoke_Arity1(func() interface{} {
				if Reduced_QMARK_.X_invoke_Arity1(ret) {
					return Deref.X_invoke_Arity1(ret).(interface{})
				} else {
					return ret
				}
			}()).(interface{})
			_, _, _ = f___1, ret, ret___1
			if Reduced_QMARK_.X_invoke_Arity1(ret___1) {
				return Deref.X_invoke_Arity1(ret___1).(interface{})
			} else {
				return ret___1
			}
		}
	})
}(&AFn{})

/**
* Returns the sum of nums. (+) returns 0.
* @param {...*} var_args
 */
var X_PLUS_ = func(_PLUS_ *AFn) *AFn {
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

/**
* If no ys are supplied, returns the negation of x, else subtracts
* the ys from x and returns the result.
* @param {...*} var_args
 */
var X_ = func(___ *AFn) *AFn {
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

/**
* Returns the product of nums. (*) returns 1.
* @param {...*} var_args
 */
var X_STAR_ = func(_STAR_ *AFn) *AFn {
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

/**
* If no denominators are supplied, returns 1/numerator,
* else returns numerator divided by all of the denominators.
* @param {...*} var_args
 */
var X_SLASH_ = func(_SLASH_ *AFn) *AFn {
	return Fn(_SLASH_, func(x interface{}) float64 {
		return _SLASH_.X_invoke_Arity2(float64(1), x)
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

/**
* Returns non-nil if nums are in monotonically increasing order,
* otherwise false.
* @param {...*} var_args
 */
var X_LT_ = func(_LT_ *AFn) *AFn {
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
					{
						var G__370 = y
						var G__371 = First.X_invoke_Arity1(more)
						var G__372 = Next.Arity1IQ(more)
						x = G__370
						y = G__371
						more = G__372
						continue
					}
				} else {
					return (y.(float64) < First.X_invoke_Arity1(more).(float64))
				}
			} else {
				return false
			}
		}
	})
}(&AFn{})

/**
* Returns non-nil if nums are in monotonically non-decreasing order,
* otherwise false.
* @param {...*} var_args
 */
var X_LT__EQ_ = func(_LT__EQ_ *AFn) *AFn {
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
					{
						var G__373 = y
						var G__374 = First.X_invoke_Arity1(more)
						var G__375 = Next.Arity1IQ(more)
						x = G__373
						y = G__374
						more = G__375
						continue
					}
				} else {
					return (y.(float64) <= First.X_invoke_Arity1(more).(float64))
				}
			} else {
				return false
			}
		}
	})
}(&AFn{})

/**
* Returns non-nil if nums are in monotonically decreasing order,
* otherwise false.
* @param {...*} var_args
 */
var X_GT_ = func(_GT_ *AFn) *AFn {
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
					{
						var G__376 = y
						var G__377 = First.X_invoke_Arity1(more)
						var G__378 = Next.Arity1IQ(more)
						x = G__376
						y = G__377
						more = G__378
						continue
					}
				} else {
					return (y.(float64) > First.X_invoke_Arity1(more).(float64))
				}
			} else {
				return false
			}
		}
	})
}(&AFn{})

/**
* Returns non-nil if nums are in monotonically non-increasing order,
* otherwise false.
* @param {...*} var_args
 */
var X_GT__EQ_ = func(_GT__EQ_ *AFn) *AFn {
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
					{
						var G__379 = y
						var G__380 = First.X_invoke_Arity1(more)
						var G__381 = Next.Arity1IQ(more)
						x = G__379
						y = G__380
						more = G__381
						continue
					}
				} else {
					return (y.(float64) >= First.X_invoke_Arity1(more).(float64))
				}
			} else {
				return false
			}
		}
	})
}(&AFn{})

/**
* Returns a number one less than num.
 */
var Dec = func(dec *AFn) *AFn {
	return Fn(dec, func(x interface{}) interface{} {
		return (x.(float64) - float64(1))
	})
}(&AFn{})

/**
* Returns the greatest of the nums.
* @param {...*} var_args
 */
var Max = func(max *AFn) *AFn {
	return Fn(max, func(x interface{}) float64 {
		return x.(float64)
	}, func(x interface{}, y interface{}) float64 {
		{
			var x__76939__auto__ = x
			var y__76940__auto__ = y
			_, _ = x__76939__auto__, y__76940__auto__
			return func() float64 {
				if x__76939__auto__ > y__76940__auto__ {
					return x__76939__auto__
				} else {
					return y__76940__auto__
				}
			}().(float64)
		}
	}, func(x_y_more__ ...interface{}) interface{} {
		var x = x_y_more__[0]
		var y = x_y_more__[1]
		var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
		_, _, _ = x, y, more
		return Reduce.X_invoke_Arity3(max, func() interface{} {
			var x__76939__auto__ = x
			var y__76940__auto__ = y
			_, _ = x__76939__auto__, y__76940__auto__
			return func() float64 {
				if x__76939__auto__ > y__76940__auto__ {
					return x__76939__auto__
				} else {
					return y__76940__auto__
				}
			}()
		}(), more)
	})
}(&AFn{})

/**
* Returns the least of the nums.
* @param {...*} var_args
 */
var Min = func(min *AFn) *AFn {
	return Fn(min, func(x interface{}) float64 {
		return x.(float64)
	}, func(x interface{}, y interface{}) float64 {
		{
			var x__76946__auto__ = x
			var y__76947__auto__ = y
			_, _ = x__76946__auto__, y__76947__auto__
			return func() float64 {
				if x__76946__auto__ < y__76947__auto__ {
					return x__76946__auto__
				} else {
					return y__76947__auto__
				}
			}().(float64)
		}
	}, func(x_y_more__ ...interface{}) interface{} {
		var x = x_y_more__[0]
		var y = x_y_more__[1]
		var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
		_, _, _ = x, y, more
		return Reduce.X_invoke_Arity3(min, func() interface{} {
			var x__76946__auto__ = x
			var y__76947__auto__ = y
			_, _ = x__76946__auto__, y__76947__auto__
			return func() float64 {
				if x__76946__auto__ < y__76947__auto__ {
					return x__76946__auto__
				} else {
					return y__76947__auto__
				}
			}()
		}(), more)
	})
}(&AFn{})

var Byte = func(byte *AFn) *AFn {
	return Fn(byte, func(x interface{}) float64 {
		return x.(float64)
	})
}(&AFn{})

/**
* Coerce to char
 */
var Char = func(char *AFn) *AFn {
	return Fn(char, func(x interface{}) interface{} {
		if reflect.TypeOf(x).Kind() == reflect.Float64 {
			return Native_invoke_instance_method.X_invoke_Arity3(js.String, "FromCharCode", []interface{}{x})
		} else {
			if (reflect.TypeOf(x).Kind() == reflect.String) && (Native_get_instance_field.X_invoke_Arity2(x, "Length").(float64) == float64(1)) {
				return x
			} else {
				panic((&js.Error{"Argument to char must be a character or number"}))

			}
		}
	})
}(&AFn{})

var Short = func(short *AFn) *AFn {
	return Fn(short, func(x interface{}) float64 {
		return x.(float64)
	})
}(&AFn{})

var Float = func(float *AFn) *AFn {
	return Fn(float, func(x interface{}) float64 {
		return x.(float64)
	})
}(&AFn{})

var Double = func(double *AFn) *AFn {
	return Fn(double, func(x interface{}) float64 {
		return x.(float64)
	})
}(&AFn{})

var Unchecked_byte = func(unchecked_byte *AFn) *AFn {
	return Fn(unchecked_byte, func(x interface{}) float64 {
		return x.(float64)
	})
}(&AFn{})

var Unchecked_char = func(unchecked_char *AFn) *AFn {
	return Fn(unchecked_char, func(x interface{}) float64 {
		return x.(float64)
	})
}(&AFn{})

var Unchecked_short = func(unchecked_short *AFn) *AFn {
	return Fn(unchecked_short, func(x interface{}) float64 {
		return x.(float64)
	})
}(&AFn{})

var Unchecked_float = func(unchecked_float *AFn) *AFn {
	return Fn(unchecked_float, func(x interface{}) float64 {
		return x.(float64)
	})
}(&AFn{})

var Unchecked_double = func(unchecked_double *AFn) *AFn {
	return Fn(unchecked_double, func(x interface{}) float64 {
		return x.(float64)
	})
}(&AFn{})

/**
* Returns the sum of nums. (+) returns 0.
* @param {...*} var_args
 */
var Unchecked_add = func(unchecked_add *AFn) *AFn {
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

/**
* Returns the sum of nums. (+) returns 0.
* @param {...*} var_args
 */
var Unchecked_add_int = func(unchecked_add_int *AFn) *AFn {
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

var Unchecked_dec = func(unchecked_dec *AFn) *AFn {
	return Fn(unchecked_dec, func(x interface{}) interface{} {
		return (x.(float64) - float64(1))
	})
}(&AFn{})

var Unchecked_dec_int = func(unchecked_dec_int *AFn) *AFn {
	return Fn(unchecked_dec_int, func(x interface{}) interface{} {
		return (x.(float64) - float64(1))
	})
}(&AFn{})

/**
* If no denominators are supplied, returns 1/numerator,
* else returns numerator divided by all of the denominators.
* @param {...*} var_args
 */
var Unchecked_divide_int = func(unchecked_divide_int *AFn) *AFn {
	return Fn(unchecked_divide_int, func(x interface{}) float64 {
		return unchecked_divide_int.X_invoke_Arity2(float64(1), x)
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

var Unchecked_inc = func(unchecked_inc *AFn) *AFn {
	return Fn(unchecked_inc, func(x interface{}) interface{} {
		return (x.(float64) + float64(1))
	})
}(&AFn{})

var Unchecked_inc_int = func(unchecked_inc_int *AFn) *AFn {
	return Fn(unchecked_inc_int, func(x interface{}) interface{} {
		return (x.(float64) + float64(1))
	})
}(&AFn{})

/**
* Returns the product of nums. (*) returns 1.
* @param {...*} var_args
 */
var Unchecked_multiply = func(unchecked_multiply *AFn) *AFn {
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

/**
* Returns the product of nums. (*) returns 1.
* @param {...*} var_args
 */
var Unchecked_multiply_int = func(unchecked_multiply_int *AFn) *AFn {
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

var Unchecked_negate = func(unchecked_negate *AFn) *AFn {
	return Fn(unchecked_negate, func(x interface{}) interface{} {
		return (-x.(float64))
	})
}(&AFn{})

var Unchecked_negate_int = func(unchecked_negate_int *AFn) *AFn {
	return Fn(unchecked_negate_int, func(x interface{}) interface{} {
		return (-x.(float64))
	})
}(&AFn{})

var Unchecked_remainder_int = func(unchecked_remainder_int *AFn) *AFn {
	return Fn(unchecked_remainder_int, func(x interface{}, n interface{}) interface{} {
		return Mod.X_invoke_Arity2(x, n).(interface{})
	})
}(&AFn{})

/**
* If no ys are supplied, returns the negation of x, else subtracts
* the ys from x and returns the result.
* @param {...*} var_args
 */
var Unchecked_subtract = func(unchecked_subtract *AFn) *AFn {
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

/**
* If no ys are supplied, returns the negation of x, else subtracts
* the ys from x and returns the result.
* @param {...*} var_args
 */
var Unchecked_subtract_int = func(unchecked_subtract_int *AFn) *AFn {
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

var Fix = func(fix *AFn) *AFn {
	return Fn(fix, func(q interface{}) float64 {
		if q.(float64) >= float64(0) {
			return Math.Floor(q).(float64)
		} else {
			return Math.Ceil(q).(float64)
		}
	})
}(&AFn{})

/**
* Coerce to int by stripping decimal places.
 */
var Int = func(int *AFn) *AFn {
	return Fn(int, func(x interface{}) interface{} {
		return float64(int(x.(float64)) | int(float64(0)))
	})
}(&AFn{})

/**
* Coerce to int by stripping decimal places.
 */
var Unchecked_int = func(unchecked_int *AFn) *AFn {
	return Fn(unchecked_int, func(x interface{}) interface{} {
		return Fix.Arity1IF(x)
	})
}(&AFn{})

/**
* Coerce to long by stripping decimal places. Identical to `int'.
 */
var Long = func(long *AFn) *AFn {
	return Fn(long, func(x interface{}) interface{} {
		return Fix.Arity1IF(x)
	})
}(&AFn{})

/**
* Coerce to long by stripping decimal places. Identical to `int'.
 */
var Unchecked_long = func(unchecked_long *AFn) *AFn {
	return Fn(unchecked_long, func(x interface{}) interface{} {
		return Fix.Arity1IF(x)
	})
}(&AFn{})

var Booleans = func(booleans *AFn) *AFn {
	return Fn(booleans, func(x interface{}) interface{} {
		return x
	})
}(&AFn{})

var Bytes = func(bytes *AFn) *AFn {
	return Fn(bytes, func(x interface{}) interface{} {
		return x
	})
}(&AFn{})

var Chars = func(chars *AFn) *AFn {
	return Fn(chars, func(x interface{}) interface{} {
		return x
	})
}(&AFn{})

var Shorts = func(shorts *AFn) *AFn {
	return Fn(shorts, func(x interface{}) interface{} {
		return x
	})
}(&AFn{})

var Ints = func(ints *AFn) *AFn {
	return Fn(ints, func(x interface{}) interface{} {
		return x
	})
}(&AFn{})

var Floats = func(floats *AFn) *AFn {
	return Fn(floats, func(x interface{}) interface{} {
		return x
	})
}(&AFn{})

var Doubles = func(doubles *AFn) *AFn {
	return Fn(doubles, func(x interface{}) interface{} {
		return x
	})
}(&AFn{})

var Longs = func(longs *AFn) *AFn {
	return Fn(longs, func(x interface{}) interface{} {
		return x
	})
}(&AFn{})

/**
* Modulus of num and div with original javascript behavior. i.e. bug for negative numbers
 */
var Js_mod = func(js_mod *AFn) *AFn {
	return Fn(js_mod, func(n interface{}, d interface{}) interface{} {
		return (n.(float64) % d.(float64))
	})
}(&AFn{})

/**
* Modulus of num and div. Truncates toward negative infinity.
 */
var Mod = func(mod *AFn) *AFn {
	return Fn(mod, func(n interface{}, d interface{}) interface{} {
		return (((n.(float64) % d.(float64)) + d.(float64)) % d.(float64))
	})
}(&AFn{})

/**
* quot[ient] of dividing numerator by denominator.
 */
var Quot = func(quot *AFn) *AFn {
	return Fn(quot, func(n interface{}, d interface{}) interface{} {
		{
			var rem = (n.(float64) % d.(float64))
			_ = rem
			return Fix.X_invoke_Arity1(((n.(float64) - rem) / d.(float64)))
		}
	})
}(&AFn{})

/**
* remainder of dividing numerator by denominator.
 */
var Rem = func(rem *AFn) *AFn {
	return Fn(rem, func(n interface{}, d interface{}) interface{} {
		{
			var q = Quot.X_invoke_Arity2(n, d).(float64)
			_ = q
			return (n.(float64) - (d.(float64) * q))
		}
	})
}(&AFn{})

/**
* Returns a random floating point number between 0 (inclusive) and n (default 1) (exclusive).
 */
var Rand = func(rand *AFn) *AFn {
	return Fn(rand, func() float64 {
		return Math.Random().(float64)
	}, func(n interface{}) float64 {
		return (n.(float64) * rand.Arity0F())
	})
}(&AFn{})

/**
* Returns a random integer between 0 (inclusive) and n (exclusive).
 */
var Rand_int = func(rand_int *AFn) *AFn {
	return Fn(rand_int, func(n interface{}) interface{} {
		return Fix.X_invoke_Arity1(Rand.Arity1IF(n))
	})
}(&AFn{})

/**
* Bitwise exclusive or
 */
var Bit_xor = func(bit_xor *AFn) *AFn {
	return Fn(bit_xor, func(x interface{}, y interface{}) interface{} {
		return float64(int(x.(float64)) ^ int(y.(float64)))
	})
}(&AFn{})

/**
* Bitwise and
 */
var Bit_and = func(bit_and *AFn) *AFn {
	return Fn(bit_and, func(x interface{}, y interface{}) interface{} {
		return float64(int(x.(float64)) & int(y.(float64)))
	})
}(&AFn{})

/**
* Bitwise or
 */
var Bit_or = func(bit_or *AFn) *AFn {
	return Fn(bit_or, func(x interface{}, y interface{}) interface{} {
		return float64(int(x.(float64)) | int(y.(float64)))
	})
}(&AFn{})

/**
* Bitwise and
 */
var Bit_and_not = func(bit_and_not *AFn) *AFn {
	return Fn(bit_and_not, func(x interface{}, y interface{}) interface{} {
		return float64(int(x.(float64)) &^ int(y.(float64)))
	})
}(&AFn{})

/**
* Clear bit at index n
 */
var Bit_clear = func(bit_clear *AFn) *AFn {
	return Fn(bit_clear, func(x interface{}, n interface{}) interface{} {
		return float64(int(x.(float64)) &^ (1 << uint(n.(float64))))
	})
}(&AFn{})

/**
* Flip bit at index n
 */
var Bit_flip = func(bit_flip *AFn) *AFn {
	return Fn(bit_flip, func(x interface{}, n interface{}) interface{} {
		return float64(int(x.(float64)) ^ (1 << uint(n.(float64))))
	})
}(&AFn{})

/**
* Bitwise complement
 */
var Bit_not = func(bit_not *AFn) *AFn {
	return Fn(bit_not, func(x interface{}) interface{} {
		return float64(^int(x.(float64)))
	})
}(&AFn{})

/**
* Set bit at index n
 */
var Bit_set = func(bit_set *AFn) *AFn {
	return Fn(bit_set, func(x interface{}, n interface{}) interface{} {
		return float64(int(x.(float64)) | (1 << iint(n.(float64))))
	})
}(&AFn{})

/**
* Test bit at index n
 */
var Bit_test = func(bit_test *AFn) *AFn {
	return Fn(bit_test, func(x interface{}, n interface{}) interface{} {
		return float64((int(x.(float64)) & (1 << uint(n.(float64)))) != 0)
	})
}(&AFn{})

/**
* Bitwise shift left
 */
var Bit_shift_left = func(bit_shift_left *AFn) *AFn {
	return Fn(bit_shift_left, func(x interface{}, n interface{}) interface{} {
		return float64(int(x.(float64)) << uint(n.(float64)))
	})
}(&AFn{})

/**
* Bitwise shift right
 */
var Bit_shift_right = func(bit_shift_right *AFn) *AFn {
	return Fn(bit_shift_right, func(x interface{}, n interface{}) interface{} {
		return float64(int(x.(float64)) >> uint(n.(float64)))
	})
}(&AFn{})

/**
* DEPRECATED: Bitwise shift right with zero fill
 */
var Bit_shift_right_zero_fill = func(bit_shift_right_zero_fill *AFn) *AFn {
	return Fn(bit_shift_right_zero_fill, func(x interface{}, n interface{}) interface{} {
		return float64(uint(x.(float64)) >> uint(n.(float64)))
	})
}(&AFn{})

/**
* Bitwise shift right with zero fill
 */
var Unsigned_bit_shift_right = func(unsigned_bit_shift_right *AFn) *AFn {
	return Fn(unsigned_bit_shift_right, func(x interface{}, n interface{}) interface{} {
		return float64(uint(x.(float64)) >> uint(n.(float64)))
	})
}(&AFn{})

/**
* Counts the number of bits set in n
 */
var Bit_count = func(bit_count *AFn) *AFn {
	return Fn(bit_count, func(v interface{}) interface{} {
		{
			var v___1 = (v.(float64) - float64(int(float64(int(v.(float64))>>uint(float64(1))))&int(float64(1431655765))))
			var v___2 = (float64(int(v___1)&int(float64(858993459))) + float64(int(float64(int(v___1)>>uint(float64(2))))&int(float64(858993459))))
			_, _ = v___1, v___2
			return float64(int((float64(int((v___2+float64(int(v___2)>>uint(float64(4)))))&int(float64(252645135))) * float64(16843009))) >> uint(float64(24)))
		}
	})
}(&AFn{})

/**
* Returns non-nil if nums all have the equivalent
* value, otherwise false. Behavior on non nums is
* undefined.
* @param {...*} var_args
 */
var X_EQ__EQ_ = func(_EQ__EQ_ *AFn) *AFn {
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
					{
						var G__382 = y
						var G__383 = First.X_invoke_Arity1(more)
						var G__384 = Next.Arity1IQ(more)
						x = G__382
						y = G__383
						more = G__384
						continue
					}
				} else {
					return _EQ__EQ_.X_invoke_Arity2(y, First.X_invoke_Arity1(more))
				}
			} else {
				return false
			}
		}
	})
}(&AFn{})

/**
* Returns true if num is greater than zero, else false
 */
var Pos_QMARK_ = func(pos_QMARK_ *AFn) *AFn {
	return Fn(pos_QMARK_, func(n interface{}) bool {
		return (n.(float64) > float64(0))
	})
}(&AFn{})

var Zero_QMARK_ = func(zero_QMARK_ *AFn) *AFn {
	return Fn(zero_QMARK_, func(n interface{}) bool {
		return (n.(float64) == float64(0))
	})
}(&AFn{})

/**
* Returns true if num is less than zero, else false
 */
var Neg_QMARK_ = func(neg_QMARK_ *AFn) *AFn {
	return Fn(neg_QMARK_, func(x interface{}) bool {
		return (x.(float64) < float64(0))
	})
}(&AFn{})

/**
* Returns the nth next of coll, (seq coll) when n is 0.
 */
var Nthnext = func(nthnext *AFn) *AFn {
	return Fn(nthnext, func(coll interface{}, n interface{}) interface{} {
		{
			var n___1 = n
			var xs = Seq.Arity1IQ(coll)
			_, _ = n___1, xs
			for {
				if (xs) && (n___1.(float64) > float64(0)) {
					{
						var G__385 = (n___1.(float64) - float64(1))
						var G__386 = Next.X_invoke_Arity1(xs)
						n___1 = G__385
						xs = G__386
						continue
					}
				} else {
					return xs
				}
			}
		}
	})
}(&AFn{})

/**
* With no args, returns the empty string. With one arg x, returns
* x.toString().  (str nil) returns the empty string. With more than
* one arg, returns the concatenation of the str values of the args.
* @param {...*} var_args
 */
var Str = func(str *AFn) *AFn {
	return Fn(str, func() interface{} {
		return ""
	}, func(x interface{}) interface{} {
		if x == nil {
			return ""
		} else {
			return Native_invoke_instance_method.X_invoke_Arity3(x, "ToString", []interface{}{})
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
					{
						var G__387 = sb.Append(str.X_invoke_Arity1(First.X_invoke_Arity1(more)))
						var G__388 = Next.Arity1IQ(more)
						sb = G__387
						more = G__388
						continue
					}
				} else {
					return sb.ToString()
				}
			}
		}
	})
}(&AFn{})

/**
* Returns the substring of s beginning at start inclusive, and ending
* at end (defaults to length of string), exclusive.
 */
var Subs = func(subs *AFn) *AFn {
	return Fn(subs, func(s interface{}, start interface{}) interface{} {
		return Native_invoke_instance_method.X_invoke_Arity3(s, "Substring", []interface{}{start})
	}, func(s interface{}, start interface{}, end interface{}) interface{} {
		return Native_invoke_instance_method.X_invoke_Arity3(s, "Substring", []interface{}{start, end})
	})
}(&AFn{})

/**
* Assumes x is sequential. Returns true if x equals y, otherwise
* returns false.
 */
var Equiv_sequential = func(equiv_sequential *AFn) *AFn {
	return Fn(equiv_sequential, func(x interface{}, y interface{}) interface{} {
		return Boolean.X_invoke_Arity1(func() interface{} {
			if Sequential_QMARK_.Arity1IB(y) {
				return func() interface{} {
					if (Counted_QMARK_.Arity1IB(x)) && (Counted_QMARK_.Arity1IB(y)) && (!(Count.X_invoke_Arity1(x).(float64) == Count.X_invoke_Arity1(y).(float64))) {
						return false
					} else {
						return func() interface{} {
							var xs = Seq.Arity1IQ(x)
							var ys = Seq.Arity1IQ(y)
							_, _ = xs, ys
							for {
								if xs == nil {
									return (ys == nil)
								} else {
									if ys == nil {
										return false
									} else {
										if X_EQ_.X_invoke_Arity2(First.X_invoke_Arity1(xs), First.X_invoke_Arity1(ys)) {
											{
												var G__389 = Next.X_invoke_Arity1(xs)
												var G__390 = Next.X_invoke_Arity1(ys)
												xs = G__389
												ys = G__390
												continue
											}
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

var Hash_coll = func(hash_coll *AFn) *AFn {
	return Fn(hash_coll, func(coll interface{}) interface{} {
		if Truth_(Seq.Arity1IQ(coll)) {
			{
				var res = Hash.X_invoke_Arity1(First.X_invoke_Arity1(coll))
				var s = Next.Arity1IQ(coll)
				_, _ = res, s
				for {
					if s == nil {
						return res
					} else {
						{
							var G__391 = Hash_combine.X_invoke_Arity2(res, Hash.X_invoke_Arity1(First.X_invoke_Arity1(s))).(float64)
							var G__392 = Next.X_invoke_Arity1(s)
							res = G__391
							s = G__392
							continue
						}
					}
				}
			}
		} else {
			return float64(0)
		}
	})
}(&AFn{})

var Hash_imap = func(hash_imap *AFn) *AFn {
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
						{
							var G__393 = ((h + float64(int(Hash.X_invoke_Arity1(Key.X_invoke_Arity1(e).(interface{})).(float64))^int(Hash.X_invoke_Arity1(Val.X_invoke_Arity1(e).(interface{})).(float64)))) % float64(4503599627370496))
							var G__394 = Next.X_invoke_Arity1(s)
							h = G__393
							s = G__394
							continue
						}
					}
				} else {
					return h
				}
			}
		}
	})
}(&AFn{})

var Hash_iset = func(hash_iset *AFn) *AFn {
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
						{
							var G__395 = ((h + Hash.X_invoke_Arity1(e).(float64)) % float64(4503599627370496))
							var G__396 = Next.X_invoke_Arity1(s___1)
							h = G__395
							s___1 = G__396
							continue
						}
					}
				} else {
					return h
				}
			}
		}
	})
}(&AFn{})

/**
* Takes a JavaScript object and a map of names to functions and
* attaches said functions as methods on the object.  Any references to
* JavaScript's implict this (via the this-as macro) will resolve to the
* object that the function is attached.
 */
var Extend_object_BANG_ = func(extend_object_BANG_ *AFn) *AFn {
	return Fn(extend_object_BANG_, func(obj interface{}, fn_map interface{}) interface{} {
		{
			var seq__403_409 = Seq.Arity1IQ(fn_map)
			var chunk__404_410 = nil
			var count__405_411 = float64(0)
			var i__406_412 = float64(0)
			_, _, _, _ = seq__403_409, chunk__404_410, count__405_411, i__406_412
			for {
				if i__406_412 < count__405_411 {
					{
						var vec__407_413 = chunk__404_410.X_nth_Arity2(i__406_412).(interface{})
						var key_name_414 = Nth.X_invoke_Arity3(vec__407_413, float64(0), nil)
						var f_415 = Nth.X_invoke_Arity3(vec__407_413, float64(1), nil)
						_, _, _ = vec__407_413, key_name_414, f_415
						{
							var str_name_416 = Name.X_invoke_Arity1(key_name_414).(interface{})
							_ = str_name_416
							obj.([]interface{})[int(str_name_416.(float64))] = f_415
						}
						{
							var G__417 = seq__403_409
							var G__418 = chunk__404_410
							var G__419 = count__405_411
							var G__420 = (i__406_412 + float64(1))
							seq__403_409 = G__417
							chunk__404_410 = G__418
							count__405_411 = G__419
							i__406_412 = G__420
							continue
						}
					}
				} else {
					{
						var temp__4219__auto___421 = Seq.X_invoke_Arity1(seq__403_409)
						_ = temp__4219__auto___421
						if Truth_(temp__4219__auto___421) {
							{
								var seq__403_422___1 = temp__4219__auto___421
								_ = seq__403_422___1
								if Chunked_seq_QMARK_.X_invoke_Arity1(seq__403_422___1) {
									{
										var c__77428__auto___423 = Chunk_first.X_invoke_Arity1(seq__403_422___1).(interface{})
										_ = c__77428__auto___423
										{
											var G__424 = Chunk_rest.X_invoke_Arity1(seq__403_422___1).(interface{})
											var G__425 = c__77428__auto___423
											var G__426 = Count.X_invoke_Arity1(c__77428__auto___423).(float64)
											var G__427 = float64(0)
											seq__403_409 = G__424
											chunk__404_410 = G__425
											count__405_411 = G__426
											i__406_412 = G__427
											continue
										}
									}
								} else {
									{
										var vec__408_428 = First.X_invoke_Arity1(seq__403_422___1)
										var key_name_429 = Nth.X_invoke_Arity3(vec__408_428, float64(0), nil)
										var f_430 = Nth.X_invoke_Arity3(vec__408_428, float64(1), nil)
										_, _, _ = vec__408_428, key_name_429, f_430
										{
											var str_name_431 = Name.X_invoke_Arity1(key_name_429).(interface{})
											_ = str_name_431
											obj.([]interface{})[int(str_name_431.(float64))] = f_430
										}
										{
											var G__432 = Next.X_invoke_Arity1(seq__403_422___1)
											var G__433 = nil
											var G__434 = float64(0)
											var G__435 = float64(0)
											seq__403_409 = G__432
											chunk__404_410 = G__433
											count__405_411 = G__434
											i__406_412 = G__435
											continue
										}
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

type CljsCoreList struct {
	Meta    interface{}
	First   interface{}
	Rest    interface{}
	Count   interface{}
	X__hash interface{}
}

func (_ *CljsCoreList) CljsCoreObject__() {}
func (self__ *CljsCoreList) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreList) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreList) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreList) CljsCoreIMeta__() {}
func (self__ *CljsCoreList) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreList) CljsCoreICloneable__() {}
func (self__ *CljsCoreList) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCoreList{self__.Meta, self__.First, self__.Rest, self__.Count, self__.X__hash})
	}
}

func (_ *CljsCoreList) CljsCoreINext__() {}
func (self__ *CljsCoreList) X_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Count.(float64) == float64(1) {
			return nil
		} else {
			return self__.Rest
		}
	}
}

func (_ *CljsCoreList) CljsCoreICounted__() {}
func (self__ *CljsCoreList) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Count.(float64)
	}
}

func (_ *CljsCoreList) CljsCoreIStack__() {}
func (self__ *CljsCoreList) X_peek_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.First
	}
}

func (self__ *CljsCoreList) X_pop_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_rest_Arity1()
	}
}

func (_ *CljsCoreList) CljsCoreIHash__() {}
func (self__ *CljsCoreList) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreList) CljsCoreIEquiv__() {}
func (self__ *CljsCoreList) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreList) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreList) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return CljsCoreList_EMPTY
	}
}

func (_ *CljsCoreList) CljsCoreIReduce__() {}
func (self__ *CljsCoreList) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCoreList) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCoreList) CljsCoreISeq__() {}
func (self__ *CljsCoreList) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.First
	}
}

func (self__ *CljsCoreList) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Count.(float64) == float64(1) {
			return CljsCoreList_EMPTY
		} else {
			return self__.Rest
		}
	}
}

func (_ *CljsCoreList) CljsCoreISeqable__() {}
func (self__ *CljsCoreList) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1
	}
}

func (_ *CljsCoreList) CljsCoreISequential__() {}
func (_ *CljsCoreList) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreList) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreList{meta___1, self__.First, self__.Rest, self__.Count, self__.X__hash})
	}
}

func (_ *CljsCoreList) CljsCoreICollection__() {}
func (self__ *CljsCoreList) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreList{self__.Meta, o, coll___1, (self__.Count.(float64) + float64(1)), nil})
	}
}

func (_ *CljsCoreList) CljsCoreASeq__()  {}
func (_ *CljsCoreList) CljsCoreIList__() {}

var X__GT_List = func(__GT_List *AFn) *AFn {
	return Fn(__GT_List, func(meta interface{}, first interface{}, rest interface{}, count interface{}, __hash interface{}) interface{} {
		return (&CljsCoreList{meta, first, rest, count, __hash})
	})
}(&AFn{})

type CljsCoreEmptyList struct{ Meta interface{} }

func (_ *CljsCoreEmptyList) CljsCoreObject__() {}
func (self__ *CljsCoreEmptyList) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreEmptyList) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreEmptyList) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreEmptyList) CljsCoreIMeta__() {}
func (self__ *CljsCoreEmptyList) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreEmptyList) CljsCoreICloneable__() {}
func (self__ *CljsCoreEmptyList) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCoreEmptyList{self__.Meta})
	}
}

func (_ *CljsCoreEmptyList) CljsCoreINext__() {}
func (self__ *CljsCoreEmptyList) X_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return nil
	}
}

func (_ *CljsCoreEmptyList) CljsCoreICounted__() {}
func (self__ *CljsCoreEmptyList) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return float64(0)
	}
}

func (_ *CljsCoreEmptyList) CljsCoreIStack__() {}
func (self__ *CljsCoreEmptyList) X_peek_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return nil
	}
}

func (self__ *CljsCoreEmptyList) X_pop_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		panic((&js.Error{"Can't pop empty list"}))
	}
}

func (_ *CljsCoreEmptyList) CljsCoreIHash__() {}
func (self__ *CljsCoreEmptyList) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return float64(0)
	}
}

func (_ *CljsCoreEmptyList) CljsCoreIEquiv__() {}
func (self__ *CljsCoreEmptyList) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreEmptyList) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreEmptyList) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1
	}
}

func (_ *CljsCoreEmptyList) CljsCoreIReduce__() {}
func (self__ *CljsCoreEmptyList) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCoreEmptyList) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCoreEmptyList) CljsCoreISeq__() {}
func (self__ *CljsCoreEmptyList) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return nil
	}
}

func (self__ *CljsCoreEmptyList) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return CljsCoreList_EMPTY
	}
}

func (_ *CljsCoreEmptyList) CljsCoreISeqable__() {}
func (self__ *CljsCoreEmptyList) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return nil
	}
}

func (_ *CljsCoreEmptyList) CljsCoreISequential__() {}
func (_ *CljsCoreEmptyList) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreEmptyList) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreEmptyList{meta___1})
	}
}

func (_ *CljsCoreEmptyList) CljsCoreICollection__() {}
func (self__ *CljsCoreEmptyList) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreList{self__.Meta, o, nil, float64(1), nil})
	}
}

func (_ *CljsCoreEmptyList) CljsCoreIList__() {}

var X__GT_EmptyList = func(__GT_EmptyList *AFn) *AFn {
	return Fn(__GT_EmptyList, func(meta interface{}) interface{} {
		return (&CljsCoreEmptyList{meta})
	})
}(&AFn{})

var CljsCoreList_EMPTY = (&CljsCoreEmptyList{nil})

var Reversible_QMARK_ = func(reversible_QMARK_ *AFn) *AFn {
	return Fn(reversible_QMARK_, func(coll interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IReversible", Str: "cljs.core/IReversible", X_hash: float64(-1422278012), X_meta: nil}), coll).(bool)
	})
}(&AFn{})

var Rseq = func(rseq *AFn) *AFn {
	return Fn(rseq, func(coll interface{}) CljsCoreISeq {
		return coll.(CljsCoreIReversible).X_rseq_Arity1().(CljsCoreISeq)
	})
}(&AFn{})

/**
* Returns a seq of the items in coll in reverse order. Not lazy.
 */
var Reverse = func(reverse *AFn) *AFn {
	return Fn(reverse, func(coll interface{}) interface{} {
		if Reversible_QMARK_.Arity1IB(coll) {
			return Rseq.Arity1IQ(coll)
		} else {
			return Reduce.X_invoke_Arity3(Conj, CljsCoreList_EMPTY, coll)
		}
	})
}(&AFn{})

/**
* @param {...*} var_args
 */
var List = func(list *AFn) *AFn {
	return Fn(list, func(xs__ ...interface{}) interface{} {
		var xs = Array_seq.X_invoke_Arity1(xs__[0:])
		_ = xs
		{
			var arr = func() interface{} {
				if (func() bool { _, instanceof := xs.(*CljsCoreIndexedSeq); return instanceof }()) && (Native_get_instance_field.X_invoke_Arity2(xs, "I").(float64) == float64(0)) {
					return Native_get_instance_field.X_invoke_Arity2(xs, "Arr")
				} else {
					return func() interface{} {
						var arr = []interface{}{}
						_ = arr
						{
							var xs___1 = xs
							_ = xs___1
							for {
								if !(xs___1 == nil) {
									arr.Push(xs___1.X_first_Arity1().(interface{}))
									{
										var G__436 = xs___1.X_next_Arity1()
										xs___1 = G__436
										continue
									}
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
				var r = CljsCoreList_EMPTY
				_, _ = i, r
				for {
					if i > float64(0) {
						{
							var G__437 = (i - float64(1))
							var G__438 = r.X_conj_Arity2((arr.([]interface{})[int((i - float64(1)))]))
							i = G__437
							r = G__438
							continue
						}
					} else {
						return r
					}
				}
			}
		}
	})
}(&AFn{})

type CljsCoreCons struct {
	Meta    interface{}
	First   interface{}
	Rest    interface{}
	X__hash interface{}
}

func (_ *CljsCoreCons) CljsCoreObject__() {}
func (self__ *CljsCoreCons) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreCons) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreCons) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreCons) CljsCoreIMeta__() {}
func (self__ *CljsCoreCons) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreCons) CljsCoreICloneable__() {}
func (self__ *CljsCoreCons) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCoreCons{self__.Meta, self__.First, self__.Rest, self__.X__hash})
	}
}

func (_ *CljsCoreCons) CljsCoreINext__() {}
func (self__ *CljsCoreCons) X_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Rest == nil {
			return nil
		} else {
			return Seq.Arity1IQ(self__.Rest)
		}
	}
}

func (_ *CljsCoreCons) CljsCoreIHash__() {}
func (self__ *CljsCoreCons) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreCons) CljsCoreIEquiv__() {}
func (self__ *CljsCoreCons) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreCons) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreCons) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.Meta)
	}
}

func (_ *CljsCoreCons) CljsCoreIReduce__() {}
func (self__ *CljsCoreCons) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCoreCons) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCoreCons) CljsCoreISeq__() {}
func (self__ *CljsCoreCons) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.First
	}
}

func (self__ *CljsCoreCons) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Rest == nil {
			return CljsCoreList_EMPTY
		} else {
			return self__.Rest
		}
	}
}

func (_ *CljsCoreCons) CljsCoreISeqable__() {}
func (self__ *CljsCoreCons) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1
	}
}

func (_ *CljsCoreCons) CljsCoreISequential__() {}
func (_ *CljsCoreCons) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreCons) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreCons{meta___1, self__.First, self__.Rest, self__.X__hash})
	}
}

func (_ *CljsCoreCons) CljsCoreICollection__() {}
func (self__ *CljsCoreCons) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreCons{nil, o, coll___1, self__.X__hash})
	}
}

func (_ *CljsCoreCons) CljsCoreASeq__()  {}
func (_ *CljsCoreCons) CljsCoreIList__() {}

var X__GT_Cons = func(__GT_Cons *AFn) *AFn {
	return Fn(__GT_Cons, func(meta interface{}, first interface{}, rest interface{}, __hash interface{}) interface{} {
		return (&CljsCoreCons{meta, first, rest, __hash})
	})
}(&AFn{})

/**
* Returns a new seq where x is the first element and seq is the rest.
 */
var Cons = func(cons *AFn) *AFn {
	return Fn(cons, func(x interface{}, coll interface{}) interface{} {
		if Truth_(func() interface{} {
			var or__76632__auto__ = (coll == nil)
			_ = or__76632__auto__
			if or__76632__auto__ {
				return or__76632__auto__
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

var List_QMARK_ = func(list_QMARK_ *AFn) *AFn {
	return Fn(list_QMARK_, func(x interface{}) bool {
		return Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IList", Str: "cljs.core/IList", X_hash: float64(1015168964), X_meta: nil}), x).(bool)
	})
}(&AFn{})

var Hash_keyword = func(hash_keyword *AFn) *AFn {
	return Fn(hash_keyword, func(k interface{}) interface{} {
		return float64(int((Hash_symbol.X_invoke_Arity1(k).(float64) + float64(2654435769))) | int(float64(0)))
	})
}(&AFn{})

type CljsCoreKeyword struct {
	Ns     interface{}
	Name   interface{}
	Fqn    interface{}
	X_hash interface{}
}

func (_ *CljsCoreKeyword) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreKeyword) X_pr_writer_Arity3(writer interface{}, ___ interface{}) interface{} {
	{
		var o___1 = self__
		_ = o___1
		return writer.(CljsCoreIWriter).X_write_Arity2((":" + Str.X_invoke_Arity1(self__.Fqn).(string))).(interface{})
	}
}

func (_ *CljsCoreKeyword) CljsCoreINamed__() {}
func (self__ *CljsCoreKeyword) X_name_Arity1() string {
	{
		var ______1 = self__
		_ = ______1
		return self__.Name.(string)
	}
}

func (self__ *CljsCoreKeyword) X_namespace_Arity1() string {
	{
		var ______1 = self__
		_ = ______1
		return self__.Ns.(string)
	}
}

func (_ *CljsCoreKeyword) CljsCoreIHash__() {}
func (self__ *CljsCoreKeyword) X_hash_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		{
			var h__77043__auto__ = self__.X_hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_keyword.X_invoke_Arity1(this___1).(float64)
					_ = h__77043__auto_____1
					self__.X_hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreKeyword) CljsCoreIFn__() {}
func (self__ *CljsCoreKeyword) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity1(coll interface{}) interface{} {
	{
		var kw___1 = self__
		_ = kw___1
		return Get.X_invoke_Arity2(coll, kw___1)
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity2(coll interface{}, not_found interface{}) interface{} {
	{
		var kw___1 = self__
		_ = kw___1
		return Get.X_invoke_Arity3(coll, kw___1, not_found)
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCoreKeyword) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

func (_ *CljsCoreKeyword) CljsCoreIEquiv__() {}
func (self__ *CljsCoreKeyword) X_equiv_Arity2(other interface{}) bool {
	{
		var ______1 = self__
		_ = ______1
		if func() bool { _, instanceof := other.(*CljsCoreKeyword); return instanceof }() {
			return (self__.Fqn == Native_get_instance_field.X_invoke_Arity2(other, "Fqn"))
		} else {
			return false
		}
	}
}

func (_ *CljsCoreKeyword) CljsCoreObject__() {}
func (self__ *CljsCoreKeyword) ToString() string {
	{
		var ______1 = self__
		_ = ______1
		return (":" + Str.X_invoke_Arity1(self__.Fqn).(string))
	}
}

func (self__ *CljsCoreKeyword) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreKeyword) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

var X__GT_Keyword = func(__GT_Keyword *AFn) *AFn {
	return Fn(__GT_Keyword, func(ns interface{}, name interface{}, fqn interface{}, _hash interface{}) interface{} {
		return (&CljsCoreKeyword{ns, name, fqn, _hash})
	})
}(&AFn{})

var Keyword_QMARK_ = func(keyword_QMARK_ *AFn) *AFn {
	return Fn(keyword_QMARK_, func(x interface{}) bool {
		return func() bool { _, instanceof := x.(*CljsCoreKeyword); return instanceof }()
	})
}(&AFn{})

var Keyword_identical_QMARK_ = func(keyword_identical_QMARK_ *AFn) *AFn {
	return Fn(keyword_identical_QMARK_, func(x interface{}, y interface{}) bool {
		if x == y {
			return true
		} else {
			if (func() bool { _, instanceof := x.(*CljsCoreKeyword); return instanceof }()) && (func() bool { _, instanceof := y.(*CljsCoreKeyword); return instanceof }()) {
				return (Native_get_instance_field.X_invoke_Arity2(x, "Fqn") == Native_get_instance_field.X_invoke_Arity2(y, "Fqn"))
			} else {
				return false
			}
		}
	})
}(&AFn{})

/**
* Returns the namespace String of a symbol or keyword, or nil if not present.
 */
var Namespace = func(namespace *AFn) *AFn {
	return Fn(namespace, func(x interface{}) interface{} {
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INamed", Str: "cljs.core/INamed", X_hash: float64(-857199025), X_meta: nil}), x)) {
			return x.X_namespace_Arity1()
		} else {
			panic((&js.Error{("Doesn't support namespace: " + Str.X_invoke_Arity1(x).(string))}))
		}
	})
}(&AFn{})

/**
* Returns a Keyword with the given namespace and name.  Do not use :
* in the keyword strings, it will be added automatically.
 */
var Keyword = func(keyword *AFn) *AFn {
	return Fn(keyword, func(name interface{}) interface{} {
		if func() bool { _, instanceof := name.(*CljsCoreKeyword); return instanceof }() {
			return name
		} else {
			if func() bool { _, instanceof := name.(*CljsCoreSymbol); return instanceof }() {
				return (&CljsCoreKeyword{Namespace.X_invoke_Arity1(name).(string), Name.X_invoke_Arity1(name).(interface{}), Native_get_instance_field.X_invoke_Arity2(name, "Str"), nil})
			} else {
				if reflect.TypeOf(name).Kind() == reflect.String {
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
			if Truth_(ns.(bool)) {
				return (`` + Str.X_invoke_Arity1(ns).(string) + "/")
			} else {
				return nil
			}
		}()).(string) + Str.X_invoke_Arity1(name).(string)), nil})
	})
}(&AFn{})

type CljsCoreLazySeq struct {
	Meta    interface{}
	Fn      interface{}
	S       interface{}
	X__hash interface{}
}

func (_ *CljsCoreLazySeq) CljsCoreObject__() {}
func (self__ *CljsCoreLazySeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreLazySeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreLazySeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (self__ *CljsCoreLazySeq) Sval() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Fn == nil {
			return self__.S
		} else {
			self__.S = self__.Fn.(CljsCoreIFn).X_invoke_Arity0().(interface{})

			self__.Fn = nil

			return self__.S
		}
	}
}

func (_ *CljsCoreLazySeq) CljsCoreIMeta__() {}
func (self__ *CljsCoreLazySeq) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreLazySeq) CljsCoreINext__() {}
func (self__ *CljsCoreLazySeq) X_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		coll___1.X_seq_Arity1()
		if self__.S == nil {
			return nil
		} else {
			return Next.Arity1IQ(self__.S)
		}
	}
}

func (_ *CljsCoreLazySeq) CljsCoreIHash__() {}
func (self__ *CljsCoreLazySeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreLazySeq) CljsCoreIEquiv__() {}
func (self__ *CljsCoreLazySeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreLazySeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreLazySeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.Meta)
	}
}

func (_ *CljsCoreLazySeq) CljsCoreIReduce__() {}
func (self__ *CljsCoreLazySeq) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCoreLazySeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCoreLazySeq) CljsCoreISeq__() {}
func (self__ *CljsCoreLazySeq) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		coll___1.X_seq_Arity1()
		if self__.S == nil {
			return nil
		} else {
			return First.X_invoke_Arity1(self__.S)
		}
	}
}

func (self__ *CljsCoreLazySeq) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		coll___1.X_seq_Arity1()
		if !(self__.S == nil) {
			return Rest.Arity1IQ(self__.S)
		} else {
			return CljsCoreList_EMPTY
		}
	}
}

func (_ *CljsCoreLazySeq) CljsCoreISeqable__() {}
func (self__ *CljsCoreLazySeq) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		coll___1.Sval()
		if self__.S == nil {
			return nil
		} else {
			{
				var ls = self__.S
				_ = ls
				for {
					if func() bool { _, instanceof := ls.(*CljsCoreLazySeq); return instanceof }() {
						{
							var G__439 = Native_invoke_instance_method.X_invoke_Arity3(ls, "Sval", []interface{}{})
							ls = G__439
							continue
						}
					} else {
						self__.S = ls

						return Seq.Arity1IQ(self__.S)
					}
				}
			}
		}
	}
}

func (_ *CljsCoreLazySeq) CljsCoreISequential__() {}
func (_ *CljsCoreLazySeq) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreLazySeq) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreLazySeq{meta___1, self__.Fn, self__.S, self__.X__hash})
	}
}

func (_ *CljsCoreLazySeq) CljsCoreICollection__() {}
func (self__ *CljsCoreLazySeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(*CljsCoreCons)
	}
}

var X__GT_LazySeq = func(__GT_LazySeq *AFn) *AFn {
	return Fn(__GT_LazySeq, func(meta interface{}, fn interface{}, s interface{}, __hash interface{}) interface{} {
		return (&CljsCoreLazySeq{meta, fn, s, __hash})
	})
}(&AFn{})

type CljsCoreChunkBuffer struct {
	Buf interface{}
	End interface{}
}

func (_ *CljsCoreChunkBuffer) CljsCoreICounted__() {}
func (self__ *CljsCoreChunkBuffer) X_count_Arity1() float64 {
	{
		var ______1 = self__
		_ = ______1
		return self__.End.(float64)
	}
}

func (_ *CljsCoreChunkBuffer) CljsCoreObject__() {}
func (self__ *CljsCoreChunkBuffer) Add(o interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		self__.Buf.([]interface{})[int(self__.End.(float64))] = o
		return func() float64 {
			var return__440 = (self__.End.(float64) + float64(1))
			self__.End = return__440
			return return__440
		}()
	}
}

func (self__ *CljsCoreChunkBuffer) Chunk(o interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		{
			var ret = (&CljsCoreArrayChunk{self__.Buf, float64(0), self__.End})
			_ = ret
			self__.Buf = nil

			return ret
		}
	}
}

var X__GT_ChunkBuffer = func(__GT_ChunkBuffer *AFn) *AFn {
	return Fn(__GT_ChunkBuffer, func(buf interface{}, end interface{}) interface{} {
		return (&CljsCoreChunkBuffer{buf, end})
	})
}(&AFn{})

var Chunk_buffer = func(chunk_buffer *AFn) *AFn {
	return Fn(chunk_buffer, func(capacity interface{}) interface{} {
		return (&CljsCoreChunkBuffer{make([]interface{}, int(capacity.(float64))), float64(0)})
	})
}(&AFn{})

type CljsCoreArrayChunk struct {
	Arr interface{}
	Off interface{}
	End interface{}
}

func (_ *CljsCoreArrayChunk) CljsCoreIReduce__() {}
func (self__ *CljsCoreArrayChunk) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Array_reduce.X_invoke_Arity4(self__.Arr, f, (self__.Arr.([]interface{})[int(self__.Off.(float64))]), (self__.Off.(float64) + float64(1)))
	}
}

func (self__ *CljsCoreArrayChunk) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Array_reduce.X_invoke_Arity4(self__.Arr, f, start, self__.Off)
	}
}

func (_ *CljsCoreArrayChunk) CljsCoreIChunk__() {}
func (self__ *CljsCoreArrayChunk) X_drop_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Off.(float64) == self__.End.(float64) {
			panic((&js.Error{"-drop-first of empty chunk"}))
		} else {
			return (&CljsCoreArrayChunk{self__.Arr, (self__.Off.(float64) + float64(1)), self__.End})
		}
	}
}

func (_ *CljsCoreArrayChunk) CljsCoreIIndexed__() {}
func (self__ *CljsCoreArrayChunk) X_nth_Arity2(i interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (self__.Arr.([]interface{})[int((self__.Off.(float64) + i.(float64)))])
	}
}

func (self__ *CljsCoreArrayChunk) X_nth_Arity3(i interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (i.(float64) >= float64(0)) && (i.(float64) < (self__.End.(float64) - self__.Off.(float64))) {
			return (self__.Arr.([]interface{})[int((self__.Off.(float64) + i.(float64)))])
		} else {
			return not_found
		}
	}
}

func (_ *CljsCoreArrayChunk) CljsCoreICounted__() {}
func (self__ *CljsCoreArrayChunk) X_count_Arity1() float64 {
	{
		var ______1 = self__
		_ = ______1
		return (self__.End.(float64) - self__.Off.(float64))
	}
}

var X__GT_ArrayChunk = func(__GT_ArrayChunk *AFn) *AFn {
	return Fn(__GT_ArrayChunk, func(arr interface{}, off interface{}, end interface{}) interface{} {
		return (&CljsCoreArrayChunk{arr, off, end})
	})
}(&AFn{})

var Array_chunk = func(array_chunk *AFn) *AFn {
	return Fn(array_chunk, func(arr interface{}) interface{} {
		return (&CljsCoreArrayChunk{arr, float64(0), float64(len(arr.([]interface{})))})
	}, func(arr interface{}, off interface{}) interface{} {
		return (&CljsCoreArrayChunk{arr, off, float64(len(arr.([]interface{})))})
	}, func(arr interface{}, off interface{}, end interface{}) interface{} {
		return (&CljsCoreArrayChunk{arr, off, end})
	})
}(&AFn{})

type CljsCoreChunkedCons struct {
	Chunk   interface{}
	More    interface{}
	Meta    interface{}
	X__hash interface{}
}

func (_ *CljsCoreChunkedCons) CljsCoreObject__() {}
func (self__ *CljsCoreChunkedCons) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreChunkedCons) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreChunkedCons) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreIMeta__() {}
func (self__ *CljsCoreChunkedCons) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreINext__() {}
func (self__ *CljsCoreChunkedCons) X_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Chunk.(CljsCoreICounted).X_count_Arity1() > float64(1) {
			return (&CljsCoreChunkedCons{self__.Chunk.(CljsCoreIChunk).X_drop_first_Arity1().(interface{}), self__.More, self__.Meta, nil})
		} else {
			{
				var more___1 = self__.More.(CljsCoreISeqable).X_seq_Arity1()
				_ = more___1
				if more___1 == nil {
					return nil
				} else {
					return more___1
				}
			}
		}
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreIHash__() {}
func (self__ *CljsCoreChunkedCons) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreIEquiv__() {}
func (self__ *CljsCoreChunkedCons) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreChunkedCons) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.Meta)
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreISeq__() {}
func (self__ *CljsCoreChunkedCons) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Chunk.(CljsCoreIIndexed).X_nth_Arity2(float64(0)).(interface{})
	}
}

func (self__ *CljsCoreChunkedCons) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Chunk.(CljsCoreICounted).X_count_Arity1() > float64(1) {
			return (&CljsCoreChunkedCons{self__.Chunk.(CljsCoreIChunk).X_drop_first_Arity1().(interface{}), self__.More, self__.Meta, nil})
		} else {
			if self__.More == nil {
				return CljsCoreList_EMPTY
			} else {
				return self__.More
			}
		}
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreISeqable__() {}
func (self__ *CljsCoreChunkedCons) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreISequential__() {}
func (_ *CljsCoreChunkedCons) CljsCoreIChunkedSeq__() {}
func (self__ *CljsCoreChunkedCons) X_chunked_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Chunk
	}
}

func (self__ *CljsCoreChunkedCons) X_chunked_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.More == nil {
			return CljsCoreList_EMPTY
		} else {
			return self__.More
		}
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreIWithMeta__() {}
func (self__ *CljsCoreChunkedCons) X_with_meta_Arity2(m interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreChunkedCons{self__.Chunk, self__.More, m, self__.X__hash})
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreICollection__() {}
func (self__ *CljsCoreChunkedCons) X_conj_Arity2(o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		return Cons.X_invoke_Arity2(o, this___1).(*CljsCoreCons)
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreASeq__()         {}
func (_ *CljsCoreChunkedCons) CljsCoreIChunkedNext__() {}
func (self__ *CljsCoreChunkedCons) X_chunked_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.More == nil {
			return nil
		} else {
			return self__.More
		}
	}
}

var X__GT_ChunkedCons = func(__GT_ChunkedCons *AFn) *AFn {
	return Fn(__GT_ChunkedCons, func(chunk interface{}, more interface{}, meta interface{}, __hash interface{}) interface{} {
		return (&CljsCoreChunkedCons{chunk, more, meta, __hash})
	})
}(&AFn{})

var Chunk_cons = func(chunk_cons *AFn) *AFn {
	return Fn(chunk_cons, func(chunk interface{}, rest interface{}) interface{} {
		if chunk.(CljsCoreICounted).X_count_Arity1() == float64(0) {
			return rest
		} else {
			return (&CljsCoreChunkedCons{chunk, rest, nil, nil})
		}
	})
}(&AFn{})

var Chunk_append = func(chunk_append *AFn) *AFn {
	return Fn(chunk_append, func(b interface{}, x interface{}) interface{} {
		return Native_invoke_instance_method.X_invoke_Arity3(b, "Add", []interface{}{x})
	})
}(&AFn{})

var Chunk = func(chunk *AFn) *AFn {
	return Fn(chunk, func(b interface{}) interface{} {
		return Native_invoke_instance_method.X_invoke_Arity3(b, "Chunk", []interface{}{})
	})
}(&AFn{})

var Chunk_first = func(chunk_first *AFn) *AFn {
	return Fn(chunk_first, func(s interface{}) interface{} {
		return s.(CljsCoreIChunkedSeq).X_chunked_first_Arity1().(interface{})
	})
}(&AFn{})

var Chunk_rest = func(chunk_rest *AFn) *AFn {
	return Fn(chunk_rest, func(s interface{}) interface{} {
		return s.(CljsCoreIChunkedSeq).X_chunked_rest_Arity1().(interface{})
	})
}(&AFn{})

var Chunk_next = func(chunk_next *AFn) *AFn {
	return Fn(chunk_next, func(s interface{}) interface{} {
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IChunkedNext", Str: "cljs.core/IChunkedNext", X_hash: float64(-1343796569), X_meta: nil}), s)) {
			return s.(CljsCoreIChunkedNext).X_chunked_next_Arity1().(interface{})
		} else {
			return Seq.X_invoke_Arity1(s.(CljsCoreIChunkedSeq).X_chunked_rest_Arity1().(interface{}))
		}
	})
}(&AFn{})

/**
* Naive impl of to-array as a start.
 */
var To_array = func(to_array *AFn) *AFn {
	return Fn(to_array, func(s interface{}) interface{} {
		{
			var ary = []interface{}{}
			_ = ary
			{
				var s___1 = s
				_ = s___1
				for {
					if Truth_(Seq.Arity1IQ(s___1)) {
						ary.Push(First.X_invoke_Arity1(s___1))
						{
							var G__441 = Next.Arity1IQ(s___1)
							s___1 = G__441
							continue
						}
					} else {
						return ary
					}
				}
			}
		}
	})
}(&AFn{})

/**
* Returns a (potentially-ragged) 2-dimensional array
* containing the contents of coll.
 */
var To_array_2d = func(to_array_2d *AFn) *AFn {
	return Fn(to_array_2d, func(coll interface{}) interface{} {
		{
			var ret = make([]interface{}, int(Count.X_invoke_Arity1(coll).(float64)))
			_ = ret
			{
				var i_442 = float64(0)
				var xs_443 = Seq.Arity1IQ(coll)
				_, _ = i_442, xs_443
				for {
					if Truth_(xs_443) {
						ret[int(i_442)] = To_array.X_invoke_Arity1(First.X_invoke_Arity1(xs_443)).([]interface{})
						{
							var G__444 = (i_442 + float64(1))
							var G__445 = Next.X_invoke_Arity1(xs_443)
							i_442 = G__444
							xs_443 = G__445
							continue
						}
					} else {
					}
					break
				}
			}
			return ret
		}
	})
}(&AFn{})

var Int_array = func(int_array *AFn) *AFn {
	return Fn(int_array, func(size_or_seq interface{}) interface{} {
		if reflect.TypeOf(size_or_seq).Kind() == reflect.Float64 {
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
							if (s___1) && (i < size.(float64)) {
								a[int(i)] = First.X_invoke_Arity1(s___1)
								{
									var G__446 = (i + float64(1))
									var G__447 = Next.X_invoke_Arity1(s___1)
									i = G__446
									s___1 = G__447
									continue
								}
							} else {
								return a
							}
						}
					}
				}
			} else {
				{
					var n__77528__auto___448 = size
					_ = n__77528__auto___448
					{
						var i_449 = float64(0)
						_ = i_449
						for {
							if i_449 < n__77528__auto___448.(float64) {
								a[int(i_449)] = init_val_or_seq
								{
									var G__450 = (i_449 + float64(1))
									i_449 = G__450
									continue
								}
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

var Long_array = func(long_array *AFn) *AFn {
	return Fn(long_array, func(size_or_seq interface{}) interface{} {
		if reflect.TypeOf(size_or_seq).Kind() == reflect.Float64 {
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
							if (s___1) && (i < size.(float64)) {
								a[int(i)] = First.X_invoke_Arity1(s___1)
								{
									var G__451 = (i + float64(1))
									var G__452 = Next.X_invoke_Arity1(s___1)
									i = G__451
									s___1 = G__452
									continue
								}
							} else {
								return a
							}
						}
					}
				}
			} else {
				{
					var n__77528__auto___453 = size
					_ = n__77528__auto___453
					{
						var i_454 = float64(0)
						_ = i_454
						for {
							if i_454 < n__77528__auto___453.(float64) {
								a[int(i_454)] = init_val_or_seq
								{
									var G__455 = (i_454 + float64(1))
									i_454 = G__455
									continue
								}
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

var Double_array = func(double_array *AFn) *AFn {
	return Fn(double_array, func(size_or_seq interface{}) interface{} {
		if reflect.TypeOf(size_or_seq).Kind() == reflect.Float64 {
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
							if (s___1) && (i < size.(float64)) {
								a[int(i)] = First.X_invoke_Arity1(s___1)
								{
									var G__456 = (i + float64(1))
									var G__457 = Next.X_invoke_Arity1(s___1)
									i = G__456
									s___1 = G__457
									continue
								}
							} else {
								return a
							}
						}
					}
				}
			} else {
				{
					var n__77528__auto___458 = size
					_ = n__77528__auto___458
					{
						var i_459 = float64(0)
						_ = i_459
						for {
							if i_459 < n__77528__auto___458.(float64) {
								a[int(i_459)] = init_val_or_seq
								{
									var G__460 = (i_459 + float64(1))
									i_459 = G__460
									continue
								}
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

var Object_array = func(object_array *AFn) *AFn {
	return Fn(object_array, func(size_or_seq interface{}) interface{} {
		if reflect.TypeOf(size_or_seq).Kind() == reflect.Float64 {
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
							if (s___1) && (i < size.(float64)) {
								a[int(i)] = First.X_invoke_Arity1(s___1)
								{
									var G__461 = (i + float64(1))
									var G__462 = Next.X_invoke_Arity1(s___1)
									i = G__461
									s___1 = G__462
									continue
								}
							} else {
								return a
							}
						}
					}
				}
			} else {
				{
					var n__77528__auto___463 = size
					_ = n__77528__auto___463
					{
						var i_464 = float64(0)
						_ = i_464
						for {
							if i_464 < n__77528__auto___463.(float64) {
								a[int(i_464)] = init_val_or_seq
								{
									var G__465 = (i_464 + float64(1))
									i_464 = G__465
									continue
								}
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

var Bounded_count = func(bounded_count *AFn) *AFn {
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
					if (i.(float64) > float64(0)) && (Seq.Arity1IQ(s___1)) {
						{
							var G__466 = Next.Arity1IQ(s___1)
							var G__467 = (i.(float64) - float64(1))
							var G__468 = (sum + float64(1))
							s___1 = G__466
							i = G__467
							sum = G__468
							continue
						}
					} else {
						return sum
					}
				}
			}
		}
	})
}(&AFn{})

var Spread = func(spread *AFn) *AFn {
	return Fn(spread, func(arglist interface{}) interface{} {
		if arglist == nil {
			return nil
		} else {
			if Next.Arity1IQ(arglist) == nil {
				return Seq.X_invoke_Arity1(First.X_invoke_Arity1(arglist))
			} else {
				return Cons.X_invoke_Arity2(First.X_invoke_Arity1(arglist), spread.X_invoke_Arity1(Next.Arity1IQ(arglist))).(*CljsCoreCons)

			}
		}
	})
}(&AFn{})

/**
* Returns a lazy seq representing the concatenation of the elements in the supplied colls.
* @param {...*} var_args
 */
var Concat = func(concat *AFn) *AFn {
	return Fn(concat, func() interface{} {
		return (&CljsCoreLazySeq{nil, func(G__469 *AFn) *AFn {
			return Fn(G__469, func() interface{} {
				return nil
			})
		}(&AFn{}), nil, nil})
	}, func(x interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__470 *AFn) *AFn {
			return Fn(G__470, func() interface{} {
				return x
			})
		}(&AFn{}), nil, nil})
	}, func(x interface{}, y interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__471 *AFn) *AFn {
			return Fn(G__471, func() interface{} {
				{
					var s = Seq.Arity1IQ(x)
					_ = s
					if Truth_(s) {
						if Chunked_seq_QMARK_.X_invoke_Arity1(s) {
							return Chunk_cons.X_invoke_Arity2(Chunk_first.X_invoke_Arity1(s).(interface{}), concat.X_invoke_Arity2(Chunk_rest.X_invoke_Arity1(s).(interface{}), y).(*CljsCoreLazySeq))
						} else {
							return Cons.X_invoke_Arity2(First.X_invoke_Arity1(s), concat.X_invoke_Arity2(Rest.X_invoke_Arity1(s), y).(*CljsCoreLazySeq)).(*CljsCoreCons)
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
					return (&CljsCoreLazySeq{nil, func(G__472 *AFn) *AFn {
						return Fn(G__472, func() interface{} {
							{
								var xys___1 = Seq.Arity1IQ(xys)
								_ = xys___1
								if Truth_(xys___1) {
									if Chunked_seq_QMARK_.X_invoke_Arity1(xys___1) {
										return Chunk_cons.X_invoke_Arity2(Chunk_first.X_invoke_Arity1(xys___1).(interface{}), cat.X_invoke_Arity2(Chunk_rest.X_invoke_Arity1(xys___1).(interface{}), zs___1).(*CljsCoreLazySeq))
									} else {
										return Cons.X_invoke_Arity2(First.X_invoke_Arity1(xys___1), cat.X_invoke_Arity2(Rest.X_invoke_Arity1(xys___1), zs___1).(*CljsCoreLazySeq)).(*CljsCoreCons)
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
			return cat.(CljsCoreIFn).X_invoke_Arity2(concat.X_invoke_Arity2(x, y).(*CljsCoreLazySeq), zs).(*CljsCoreLazySeq)
		}
	})
}(&AFn{})

/**
* Creates a new list containing the items prepended to the rest, the
* last of which will be treated as a sequence.
* @param {...*} var_args
 */
var List_STAR_ = func(list_STAR_ *AFn) *AFn {
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

/**
* Returns a new, transient version of the collection, in constant time.
 */
var Transient = func(transient *AFn) *AFn {
	return Fn(transient, func(coll interface{}) interface{} {
		return coll.(CljsCoreIEditableCollection).X_as_transient_Arity1()
	})
}(&AFn{})

/**
* Returns a new, persistent version of the transient collection, in
* constant time. The transient collection cannot be used after this
* call, any such use will throw an exception.
 */
var Persistent_BANG_ = func(persistent_BANG_ *AFn) *AFn {
	return Fn(persistent_BANG_, func(tcoll interface{}) interface{} {
		return tcoll.(CljsCoreITransientCollection).X_persistent_BANG__Arity1()
	})
}(&AFn{})

/**
* Adds x to the transient collection, and return coll. The 'addition'
* may happen at different 'places' depending on the concrete type.
* @param {...*} var_args
 */
var Conj_BANG_ = func(conj_BANG_ *AFn) *AFn {
	return Fn(conj_BANG_, func() interface{} {
		return Transient.X_invoke_Arity1(CljsCorePersistentVector_EMPTY).(interface{})
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
					{
						var G__473 = ntcoll
						var G__474 = First.X_invoke_Arity1(vals)
						var G__475 = Next.Arity1IQ(vals)
						tcoll = G__473
						val = G__474
						vals = G__475
						continue
					}
				} else {
					return ntcoll
				}
			}
		}
	})
}(&AFn{})

/**
* When applied to a transient map, adds mapping of key(s) to
* val(s). When applied to a transient vector, sets the val at index.
* Note - index must be <= (count vector). Returns coll.
* @param {...*} var_args
 */
var Assoc_BANG_ = func(assoc_BANG_ *AFn) *AFn {
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
					{
						var G__476 = ntcoll
						var G__477 = First.X_invoke_Arity1(kvs)
						var G__478 = Second.X_invoke_Arity1(kvs)
						var G__479 = Nnext.X_invoke_Arity1(kvs).(CljsCoreISeq)
						tcoll = G__476
						key = G__477
						val = G__478
						kvs = G__479
						continue
					}
				} else {
					return ntcoll
				}
			}
		}
	})
}(&AFn{})

/**
* Returns a transient map that doesn't contain a mapping for key(s).
* @param {...*} var_args
 */
var Dissoc_BANG_ = func(dissoc_BANG_ *AFn) *AFn {
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
					{
						var G__480 = ntcoll
						var G__481 = First.X_invoke_Arity1(ks)
						var G__482 = Next.Arity1IQ(ks)
						tcoll = G__480
						key = G__481
						ks = G__482
						continue
					}
				} else {
					return ntcoll
				}
			}
		}
	})
}(&AFn{})

/**
* Removes the last item from a transient vector. If
* the collection is empty, throws an exception. Returns coll
 */
var Pop_BANG_ = func(pop_BANG_ *AFn) *AFn {
	return Fn(pop_BANG_, func(tcoll interface{}) interface{} {
		return tcoll.(CljsCoreITransientVector).X_pop_BANG__Arity1()
	})
}(&AFn{})

/**
* disj[oin]. Returns a transient set of the same (hashed/sorted) type, that
* does not contain key(s).
* @param {...*} var_args
 */
var Disj_BANG_ = func(disj_BANG_ *AFn) *AFn {
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
					{
						var G__483 = ntcoll
						var G__484 = First.X_invoke_Arity1(vals)
						var G__485 = Next.Arity1IQ(vals)
						tcoll = G__483
						val = G__484
						vals = G__485
						continue
					}
				} else {
					return ntcoll
				}
			}
		}
	})
}(&AFn{})

/**
* Returns an object of the same type and value as obj, with
* (apply f (meta obj) args) as its metadata.
* @param {...*} var_args
 */
var Vary_meta = func(vary_meta *AFn) *AFn {
	return Fn(vary_meta, func(obj interface{}, f interface{}) interface{} {
		return With_meta.X_invoke_Arity2(obj, f.(CljsCoreIFn).X_invoke_Arity1(Meta.X_invoke_Arity1(obj)).(interface{}))
	}, func(obj interface{}, f interface{}, a interface{}) interface{} {
		return With_meta.X_invoke_Arity2(obj, f.(CljsCoreIFn).X_invoke_Arity2(Meta.X_invoke_Arity1(obj), a).(interface{}))
	}, func(obj interface{}, f interface{}, a interface{}, b interface{}) interface{} {
		return With_meta.X_invoke_Arity2(obj, f.(CljsCoreIFn).X_invoke_Arity3(Meta.X_invoke_Arity1(obj), a, b).(interface{}))
	}, func(obj interface{}, f interface{}, a interface{}, b interface{}, c interface{}) interface{} {
		return With_meta.X_invoke_Arity2(obj, f.(CljsCoreIFn).X_invoke_Arity4(Meta.X_invoke_Arity1(obj), a, b, c).(interface{}))
	}, func(obj interface{}, f interface{}, a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
		return With_meta.X_invoke_Arity2(obj, f.(CljsCoreIFn).X_invoke_Arity5(Meta.X_invoke_Arity1(obj), a, b, c, d).(interface{}))
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

/**
* Same as (not (= obj1 obj2))
* @param {...*} var_args
 */
var Not_EQ_ = func(not_EQ_ *AFn) *AFn {
	return Fn(not_EQ_, func(x interface{}) bool {
		return false
	}, func(x interface{}, y interface{}) bool {
		return !(X_EQ_.Arity2IIB(x, y))
	}, func(x_y_more__ ...interface{}) interface{} {
		var x = x_y_more__[0]
		var y = x_y_more__[1]
		var more = Array_seq.X_invoke_Arity1(x_y_more__[2:])
		_, _, _ = x, y, more
		return Not.X_invoke_Arity1(Apply.X_invoke_Arity4(X_EQ_, x, y, more))
	})
}(&AFn{})

/**
* If coll is empty, returns nil, else coll
 */
var Not_empty = func(not_empty *AFn) *AFn {
	return Fn(not_empty, func(coll interface{}) interface{} {
		if Truth_(Seq.Arity1IQ(coll)) {
			return coll
		} else {
			return nil
		}
	})
}(&AFn{})

type CljsCoreStringIter struct {
	S interface{}
	I interface{}
}

func (_ *CljsCoreStringIter) CljsCoreObject__() {}
func (self__ *CljsCoreStringIter) HasNext() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (self__.I.(float64) < float64(len(self__.S.([]interface{}))))
	}
}

func (self__ *CljsCoreStringIter) Next() CljsCoreISeq {
	{
		var ______1 = self__
		_ = ______1
		{
			var ret = Native_invoke_instance_method.X_invoke_Arity3(self__.S, "CharAt", []interface{}{self__.I})
			_ = ret
			self__.I = (self__.I.(float64) + float64(1))

			return ret.(CljsCoreISeq)
		}
	}
}

func (self__ *CljsCoreStringIter) Remove() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&js.Error{"Unsupported operation"})
	}
}

var X__GT_StringIter = func(__GT_StringIter *AFn) *AFn {
	return Fn(__GT_StringIter, func(s interface{}, i interface{}) interface{} {
		return (&CljsCoreStringIter{s, i})
	})
}(&AFn{})

var String_iter = func(string_iter *AFn) *AFn {
	return Fn(string_iter, func(x interface{}) interface{} {
		return (&CljsCoreStringIter{x, float64(0)})
	})
}(&AFn{})

type CljsCoreArrayIter struct {
	Arr interface{}
	I   interface{}
}

func (_ *CljsCoreArrayIter) CljsCoreObject__() {}
func (self__ *CljsCoreArrayIter) HasNext() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (self__.I.(float64) < float64(len(self__.Arr.([]interface{}))))
	}
}

func (self__ *CljsCoreArrayIter) Next() CljsCoreISeq {
	{
		var ______1 = self__
		_ = ______1
		{
			var ret = (self__.Arr.([]interface{})[int(self__.I.(float64))])
			_ = ret
			self__.I = (self__.I.(float64) + float64(1))

			return ret.(CljsCoreISeq)
		}
	}
}

func (self__ *CljsCoreArrayIter) Remove() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&js.Error{"Unsupported operation"})
	}
}

var X__GT_ArrayIter = func(__GT_ArrayIter *AFn) *AFn {
	return Fn(__GT_ArrayIter, func(arr interface{}, i interface{}) interface{} {
		return (&CljsCoreArrayIter{arr, i})
	})
}(&AFn{})

var Array_iter = func(array_iter *AFn) *AFn {
	return Fn(array_iter, func(x interface{}) interface{} {
		return (&CljsCoreArrayIter{x, float64(0)})
	})
}(&AFn{})

var INIT = map[string]interface{}{}

var START = map[string]interface{}{}

type CljsCoreSeqIter struct {
	X_seq  interface{}
	X_next interface{}
}

func (_ *CljsCoreSeqIter) CljsCoreObject__() {}
func (self__ *CljsCoreSeqIter) HasNext() interface{} {
	{
		var ______1 = self__
		_ = ______1
		if self__.X_seq == INIT {
			self__.X_seq = START

			self__.X_next = Seq.Arity1IQ(self__.X_next)

		} else {
			if self__.X_seq == self__.X_next {
				self__.X_next = Next.Arity1IQ(self__.X_seq)

			} else {
			}
		}
		return !(self__.X_next == nil)
	}
}

func (self__ *CljsCoreSeqIter) Next() CljsCoreISeq {
	{
		var this___1 = self__
		_ = this___1
		if Not.Arity1IB(this___1.HasNext()) {
			panic((&js.Error{"No such element"}))
		} else {
			self__.X_seq = self__.X_next

			return First.X_invoke_Arity1(self__.X_next)
		}
	}
}

func (self__ *CljsCoreSeqIter) Remove() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&js.Error{"Unsupported operation"})
	}
}

var X__GT_SeqIter = func(__GT_SeqIter *AFn) *AFn {
	return Fn(__GT_SeqIter, func(_seq interface{}, _next interface{}) interface{} {
		return (&CljsCoreSeqIter{_seq, _next})
	})
}(&AFn{})

var Seq_iter = func(seq_iter *AFn) *AFn {
	return Fn(seq_iter, func(coll interface{}) interface{} {
		return (&CljsCoreSeqIter{INIT, coll})
	})
}(&AFn{})

var Iter = func(iter *AFn) *AFn {
	return Fn(iter, func(coll interface{}) interface{} {
		if coll == nil {
			return Nil_iter.X_invoke_Arity0().(*CljsCoreT489)
		} else {
			if reflect.TypeOf(coll).Kind() == reflect.String {
				return String_iter.X_invoke_Arity1(coll).(*CljsCoreStringIter)
			} else {
				if reflect.TypeOf(coll).Kind() == reflect.Slice {
					return Array_iter.X_invoke_Arity1(coll).(*CljsCoreArrayIter)
				} else {
					return Seq_iter.X_invoke_Arity1(coll).(*CljsCoreSeqIter)

				}
			}
		}
	})
}(&AFn{})

var Lazy_transformer = func(lazy_transformer *AFn) *AFn {
	return Fn(lazy_transformer, func(stepper interface{}) interface{} {
		return (&CljsCoreLazyTransformer{stepper, nil, nil, nil})
	})
}(&AFn{})

type CljsCoreStepper struct {
	Xform interface{}
	Iter  interface{}
}

func (_ *CljsCoreStepper) CljsCoreObject__() {}
func (self__ *CljsCoreStepper) Step(lt interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		{
			for {
				if Truth_(func() interface{} {
					var and__76620__auto__ = !(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper") == nil)
					_ = and__76620__auto__
					if and__76620__auto__ {
						return Native_invoke_instance_method.X_invoke_Arity3(self__.Iter, "HasNext", []interface{}{})
					} else {
						return and__76620__auto__
					}
				}()) {
					if Reduced_QMARK_.X_invoke_Arity1(self__.Xform.(CljsCoreIFn).X_invoke_Arity2(lt, Native_invoke_instance_method.X_invoke_Arity3(self__.Iter, "Next", []interface{}{})).(interface{})) {
						if Native_get_instance_field.X_invoke_Arity2(lt, "Rest") == nil {
						} else {
							Native_set_instance_field.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(lt, "Rest"), "Stepper", nil)

						}
					} else {
						{
							continue
						}
					}
				} else {
				}
				break
			}
		}
		if Native_get_instance_field.X_invoke_Arity2(lt, "Stepper") == nil {
			return nil
		} else {
			return self__.Xform.(CljsCoreIFn).X_invoke_Arity1(lt).(interface{})
		}
	}
}

var X__GT_Stepper = func(__GT_Stepper *AFn) *AFn {
	return Fn(__GT_Stepper, func(xform interface{}, iter interface{}) interface{} {
		return (&CljsCoreStepper{xform, iter})
	})
}(&AFn{})

var Stepper = func(stepper *AFn) *AFn {
	return Fn(stepper, func(xform interface{}, iter interface{}) interface{} {
		{
			var stepfn *AFn
			stepfn = func(stepfn *AFn) *AFn {
				return Fn(stepfn, func(result interface{}) interface{} {
					{
						var lt = func() interface{} {
							if Reduced_QMARK_.Arity1IB(result) {
								return Deref.X_invoke_Arity1(result).(interface{})
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

						Native_set_instance_field.X_invoke_Arity3(lt, "Rest", Lazy_transformer.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper")).(CljsCoreLazyTransformer))

						Native_set_instance_field.X_invoke_Arity3(lt, "Stepper", nil)

						return Native_get_instance_field.X_invoke_Arity2(lt, "Rest")
					}
				})
			}(&AFn{})
			_ = stepfn
			return (&CljsCoreStepper{xform.(CljsCoreIFn).X_invoke_Arity1(stepfn).(interface{}), iter})
		}
	})
}(&AFn{})

type CljsCoreMultiStepper struct {
	Xform interface{}
	Iters interface{}
	Nexts interface{}
}

func (_ *CljsCoreMultiStepper) CljsCoreObject__() {}
func (self__ *CljsCoreMultiStepper) HasNext() interface{} {
	{
		var ______1 = self__
		_ = ______1
		{
			var iters___1 = Seq.Arity1IQ(self__.Iters)
			_ = iters___1
			for {
				if !(iters___1 == nil) {
					{
						var iter = First.X_invoke_Arity1(iters___1)
						_ = iter
						if Not.Arity1IB(iter.HasNext()) {
							return false
						} else {
							{
								var G__492 = Next.X_invoke_Arity1(iters___1)
								iters___1 = G__492
								continue
							}
						}
					}
				} else {
					return true
				}
			}
		}
	}
}

func (self__ *CljsCoreMultiStepper) Next() CljsCoreISeq {
	{
		var ______1 = self__
		_ = ______1
		{
			var n__77528__auto___493 = float64(len(self__.Iters.([]interface{})))
			_ = n__77528__auto___493
			{
				var i_494 = float64(0)
				_ = i_494
				for {
					if i_494 < n__77528__auto___493 {
						self__.Nexts.([]interface{})[int(i_494)] = Native_invoke_instance_method.X_invoke_Arity3((self__.Iters.([]interface{})[int(i_494)]), "Next", []interface{}{})
						{
							var G__495 = (i_494 + float64(1))
							i_494 = G__495
							continue
						}
					} else {
					}
					break
				}
			}
		}
		return Prim_seq.X_invoke_Arity2(self__.Nexts, float64(0))
	}
}

func (self__ *CljsCoreMultiStepper) Step(lt interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		{
			for {
				if Truth_(func() interface{} {
					var and__76620__auto__ = !(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper") == nil)
					_ = and__76620__auto__
					if and__76620__auto__ {
						return this___1.HasNext()
					} else {
						return and__76620__auto__
					}
				}()) {
					if Reduced_QMARK_.X_invoke_Arity1(Apply.X_invoke_Arity2(self__.Xform, Cons.X_invoke_Arity2(lt, this___1.Next()).(*CljsCoreCons))) {
						if Native_get_instance_field.X_invoke_Arity2(lt, "Rest") == nil {
						} else {
							Native_set_instance_field.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(lt, "Rest"), "Stepper", nil)

						}
					} else {
						{
							continue
						}
					}
				} else {
				}
				break
			}
		}
		if Native_get_instance_field.X_invoke_Arity2(lt, "Stepper") == nil {
			return nil
		} else {
			return self__.Xform.(CljsCoreIFn).X_invoke_Arity1(lt).(interface{})
		}
	}
}

var X__GT_MultiStepper = func(__GT_MultiStepper *AFn) *AFn {
	return Fn(__GT_MultiStepper, func(xform interface{}, iters interface{}, nexts interface{}) interface{} {
		return (&CljsCoreMultiStepper{xform, iters, nexts})
	})
}(&AFn{})

var Multi_stepper = func(multi_stepper *AFn) *AFn {
	return Fn(multi_stepper, func(xform interface{}, iters interface{}) interface{} {
		return multi_stepper.X_invoke_Arity3(xform, iters, make([]interface{}, int(float64(len(iters.([]interface{})))))).(interface{})
	}, func(xform interface{}, iters interface{}, nexts interface{}) interface{} {
		{
			var stepfn *AFn
			stepfn = func(stepfn *AFn) *AFn {
				return Fn(stepfn, func(result interface{}) interface{} {
					{
						var lt = func() interface{} {
							if Reduced_QMARK_.Arity1IB(result) {
								return Deref.X_invoke_Arity1(result).(interface{})
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

						Native_set_instance_field.X_invoke_Arity3(lt, "Rest", Lazy_transformer.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(lt, "Stepper")).(CljsCoreLazyTransformer))

						Native_set_instance_field.X_invoke_Arity3(lt, "Stepper", nil)

						return Native_get_instance_field.X_invoke_Arity2(lt, "Rest")
					}
				})
			}(&AFn{})
			_ = stepfn
			return (&CljsCoreMultiStepper{xform.(CljsCoreIFn).X_invoke_Arity1(stepfn).(interface{}), iters, nexts})
		}
	})
}(&AFn{})

type CljsCoreLazyTransformer struct {
	Stepper interface{}
	First   interface{}
	Rest    interface{}
	Meta    interface{}
}

func (_ *CljsCoreLazyTransformer) CljsCoreINext__() {}
func (self__ *CljsCoreLazyTransformer) X_next_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		if self__.Stepper == nil {
		} else {
			this___1.X_seq_Arity1()
		}
		if self__.Rest == nil {
			return nil
		} else {
			return self__.Rest.(CljsCoreISeqable).X_seq_Arity1()
		}
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreISeq__() {}
func (self__ *CljsCoreLazyTransformer) X_first_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		if self__.Stepper == nil {
		} else {
			this___1.X_seq_Arity1()
		}
		if self__.Rest == nil {
			return nil
		} else {
			return self__.First
		}
	}
}

func (self__ *CljsCoreLazyTransformer) X_rest_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		if self__.Stepper == nil {
		} else {
			this___1.X_seq_Arity1()
		}
		if self__.Rest == nil {
			return CljsCoreList_EMPTY
		} else {
			return self__.Rest
		}
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreISeqable__() {}
func (self__ *CljsCoreLazyTransformer) X_seq_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		if self__.Stepper == nil {
		} else {
			Native_invoke_instance_method.X_invoke_Arity3(self__.Stepper, "Step", []interface{}{this___1})
		}
		if self__.Rest == nil {
			return nil
		} else {
			return this___1
		}
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreIHash__() {}
func (self__ *CljsCoreLazyTransformer) X_hash_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		return Hash_ordered_coll.X_invoke_Arity1(this___1)
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreIEquiv__() {}
func (self__ *CljsCoreLazyTransformer) X_equiv_Arity2(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		{
			var s = this___1.X_seq_Arity1()
			_ = s
			if !(s == nil) {
				return Equiv_sequential.X_invoke_Arity2(this___1, other).(bool)
			} else {
				return (Sequential_QMARK_.Arity1IB(other)) && (Seq.Arity1IQ(other) == nil)
			}
		}
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreISequential__()          {}
func (_ *CljsCoreLazyTransformer) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreLazyTransformer) X_empty_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		return CljsCoreList_EMPTY
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreICollection__() {}
func (self__ *CljsCoreLazyTransformer) X_conj_Arity2(o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		return Cons.X_invoke_Arity2(o, this___1.X_seq_Arity1()).(*CljsCoreCons)
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreIWithMeta__() {}
func (self__ *CljsCoreLazyTransformer) X_with_meta_Arity2(new_meta interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		return (&CljsCoreLazyTransformer{self__.Stepper, self__.First, self__.Rest, new_meta})
	}
}

var X__GT_LazyTransformer = func(__GT_LazyTransformer *AFn) *AFn {
	return Fn(__GT_LazyTransformer, func(stepper interface{}, first interface{}, rest interface{}, meta interface{}) interface{} {
		return (&CljsCoreLazyTransformer{stepper, first, rest, meta})
	})
}(&AFn{})

var CljsCoreLazyTransformer_Create = func(G__496 *AFn) *AFn {
	return Fn(G__496, func(xform interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazyTransformer{Stepper.X_invoke_Arity2(xform, Iter.X_invoke_Arity1(coll).(interface{})).(interface{}), nil, nil, nil})
	})
}(&AFn{})

var CljsCoreLazyTransformer_CreateMulti = func(G__501 *AFn) *AFn {
	return Fn(G__501, func(xform interface{}, colls interface{}) interface{} {
		{
			var iters = []interface{}{}
			_ = iters
			{
				var seq__497_502 = Seq.Arity1IQ(colls)
				var chunk__498_503 = nil
				var count__499_504 = float64(0)
				var i__500_505 = float64(0)
				_, _, _, _ = seq__497_502, chunk__498_503, count__499_504, i__500_505
				for {
					if i__500_505 < count__499_504 {
						{
							var coll_506 = chunk__498_503.X_nth_Arity2(i__500_505).(interface{})
							_ = coll_506
							iters.Push(Iter.X_invoke_Arity1(coll_506).(interface{}))
							{
								var G__507 = seq__497_502
								var G__508 = chunk__498_503
								var G__509 = count__499_504
								var G__510 = (i__500_505 + float64(1))
								seq__497_502 = G__507
								chunk__498_503 = G__508
								count__499_504 = G__509
								i__500_505 = G__510
								continue
							}
						}
					} else {
						{
							var temp__4219__auto___511 = Seq.X_invoke_Arity1(seq__497_502)
							_ = temp__4219__auto___511
							if Truth_(temp__4219__auto___511) {
								{
									var seq__497_512___1 = temp__4219__auto___511
									_ = seq__497_512___1
									if Chunked_seq_QMARK_.X_invoke_Arity1(seq__497_512___1) {
										{
											var c__77428__auto___513 = Chunk_first.X_invoke_Arity1(seq__497_512___1).(interface{})
											_ = c__77428__auto___513
											{
												var G__514 = Chunk_rest.X_invoke_Arity1(seq__497_512___1).(interface{})
												var G__515 = c__77428__auto___513
												var G__516 = Count.X_invoke_Arity1(c__77428__auto___513).(float64)
												var G__517 = float64(0)
												seq__497_502 = G__514
												chunk__498_503 = G__515
												count__499_504 = G__516
												i__500_505 = G__517
												continue
											}
										}
									} else {
										{
											var coll_518 = First.X_invoke_Arity1(seq__497_512___1)
											_ = coll_518
											iters.Push(Iter.X_invoke_Arity1(coll_518).(interface{}))
											{
												var G__519 = Next.X_invoke_Arity1(seq__497_512___1)
												var G__520 = nil
												var G__521 = float64(0)
												var G__522 = float64(0)
												seq__497_502 = G__519
												chunk__498_503 = G__520
												count__499_504 = G__521
												i__500_505 = G__522
												continue
											}
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
			return (&CljsCoreLazyTransformer{Multi_stepper.X_invoke_Arity3(xform, iters, make([]interface{}, int(float64(len(iters))))).(interface{}), nil, nil, nil})
		}
	})
}(&AFn{})

func init() {
	/**
	 * Coerces coll to a (possibly empty) sequence, if it is not already
	 * one. Will not force a lazy seq. (sequence nil) yields (), When a
	 * transducer is supplied, returns a lazy sequence of applications of
	 * the transform to the items in coll(s), i.e. to the set of first
	 * items of each coll, followed by the set of second
	 * items in each coll, until any one of the colls is exhausted.  Any
	 * remaining items in other colls are ignored. The transform should accept
	 * number-of-colls arguments
	 * @param {...*} var_args
	 */
	Sequence = func(sequence *AFn) *AFn {
		return Fn(sequence, func(coll interface{}) interface{} {
			if Seq_QMARK_.Arity1IB(coll) {
				return coll
			} else {
				{
					var or__76632__auto__ = Seq.Arity1IQ(coll)
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						return CljsCoreList_EMPTY
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

}

/**
* Returns true if (pred x) is logical true for every x in coll, else
* false.
 */
var Every_QMARK_ = func(every_QMARK_ *AFn) *AFn {
	return Fn(every_QMARK_, func(pred interface{}, coll interface{}) bool {
		for {
			if Seq.Arity1IQ(coll) == nil {
				return true
			} else {
				if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(coll)).(interface{})) {
					{
						var G__523 = pred
						var G__524 = Next.Arity1IQ(coll)
						pred = G__523
						coll = G__524
						continue
					}
				} else {
					return false

				}
			}
		}
	})
}(&AFn{})

/**
* Returns false if (pred x) is logical true for every x in
* coll, else true.
 */
var Not_every_QMARK_ = func(not_every_QMARK_ *AFn) *AFn {
	return Fn(not_every_QMARK_, func(pred interface{}, coll interface{}) bool {
		return !(Every_QMARK_.Arity2IIB(pred, coll))
	})
}(&AFn{})

/**
* Returns the first logical true value of (pred x) for any x in coll,
* else nil.  One common idiom is to use a set as pred, for example
* this will return :fred if :fred is in the sequence, otherwise nil:
* (some #{:fred} coll)
 */
var Some = func(some *AFn) *AFn {
	return Fn(some, func(pred interface{}, coll interface{}) interface{} {
		for {
			if Truth_(Seq.Arity1IQ(coll)) {
				{
					var or__76632__auto__ = pred.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(coll)).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						{
							var G__525 = pred
							var G__526 = Next.Arity1IQ(coll)
							pred = G__525
							coll = G__526
							continue
						}
					}
				}
			} else {
				return nil
			}
		}
	})
}(&AFn{})

/**
* Returns false if (pred x) is logical true for any x in coll,
* else true.
 */
var Not_any_QMARK_ = func(not_any_QMARK_ *AFn) *AFn {
	return Fn(not_any_QMARK_, func(pred interface{}, coll interface{}) bool {
		return Not.X_invoke_Arity1(Some.X_invoke_Arity2(pred, coll))
	})
}(&AFn{})

/**
* Returns true if n is even, throws an exception if n is not an integer
 */
var Even_QMARK_ = func(even_QMARK_ *AFn) *AFn {
	return Fn(even_QMARK_, func(n interface{}) bool {
		if Integer_QMARK_.Arity1IB(n) {
			return (float64(int(n.(float64))&int(float64(1))) == float64(0))
		} else {
			panic((&js.Error{("Argument must be an integer: " + Str.X_invoke_Arity1(n).(string))}))
		}
	})
}(&AFn{})

/**
* Returns true if n is odd, throws an exception if n is not an integer
 */
var Odd_QMARK_ = func(odd_QMARK_ *AFn) *AFn {
	return Fn(odd_QMARK_, func(n interface{}) bool {
		return !(Even_QMARK_.Arity1IB(n))
	})
}(&AFn{})

var Identity = func(identity *AFn) *AFn {
	return Fn(identity, func(x interface{}) interface{} {
		return x
	})
}(&AFn{})

/**
* Takes a fn f and returns a fn that takes the same arguments as f,
* has the same effects, if any, and returns the opposite truth value.
 */
var Complement = func(complement *AFn) *AFn {
	return Fn(complement, func(f interface{}) bool {
		return func(G__527 *AFn) *AFn {
			return Fn(G__527, func() interface{} {
				return Not.X_invoke_Arity1(f.(CljsCoreIFn).X_invoke_Arity0().(interface{}))
			}, func(x interface{}) interface{} {
				return Not.X_invoke_Arity1(f.(CljsCoreIFn).X_invoke_Arity1(x).(interface{}))
			}, func(x interface{}, y interface{}) interface{} {
				return Not.X_invoke_Arity1(f.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{}))
			}, func(x_y_zs__ ...interface{}) interface{} {
				var x = x_y_zs__[0]
				var y = x_y_zs__[1]
				var zs = Array_seq.X_invoke_Arity1(x_y_zs__[2:])
				_, _, _ = x, y, zs
				return Not.X_invoke_Arity1(Apply.X_invoke_Arity4(f, x, y, zs)).(bool)
			})
		}(&AFn{})
	})
}(&AFn{})

/**
* Returns a function that takes any number of arguments and returns x.
 */
var Constantly = func(constantly *AFn) *AFn {
	return Fn(constantly, func(x interface{}) interface{} {
		return func(G__528 *AFn) *AFn {
			return Fn(G__528, func(args__ ...interface{}) interface{} {
				var args = Array_seq.X_invoke_Arity1(args__[0:])
				_ = args
				return x
			})
		}(&AFn{})
	})
}(&AFn{})

/**
* Takes a set of functions and returns a fn that is the composition
* of those fns.  The returned fn takes a variable number of args,
* applies the rightmost of fns to the args, the next
* fn (right-to-left) to the result, etc.
* @param {...*} var_args
 */
var Comp = func(comp *AFn) *AFn {
	return Fn(comp, func() interface{} {
		return Identity
	}, func(f interface{}) interface{} {
		return f
	}, func(f interface{}, g interface{}) interface{} {
		return func(G__529 *AFn) *AFn {
			return Fn(G__529, func() interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity0().(interface{})).(interface{})
			}, func(x interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})).(interface{})
			}, func(x interface{}, y interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{})).(interface{})
			}, func(x interface{}, y interface{}, z interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity3(x, y, z).(interface{})).(interface{})
			}, func(x_y_z_args__ ...interface{}) interface{} {
				var x = x_y_z_args__[0]
				var y = x_y_z_args__[1]
				var z = x_y_z_args__[2]
				var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
				_, _, _, _ = x, y, z, args
				return f.(CljsCoreIFn).X_invoke_Arity1(Apply.X_invoke_Arity5(g, x, y, z, args)).(interface{})
			})
		}(&AFn{})
	}, func(f interface{}, g interface{}, h interface{}) interface{} {
		return func(G__530 *AFn) *AFn {
			return Fn(G__530, func() interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(h.(CljsCoreIFn).X_invoke_Arity0().(interface{})).(interface{})).(interface{})
			}, func(x interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(h.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})).(interface{})).(interface{})
			}, func(x interface{}, y interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(h.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{})).(interface{})).(interface{})
			}, func(x interface{}, y interface{}, z interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(h.(CljsCoreIFn).X_invoke_Arity3(x, y, z).(interface{})).(interface{})).(interface{})
			}, func(x_y_z_args__ ...interface{}) interface{} {
				var x = x_y_z_args__[0]
				var y = x_y_z_args__[1]
				var z = x_y_z_args__[2]
				var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
				_, _, _, _ = x, y, z, args
				return f.(CljsCoreIFn).X_invoke_Arity1(g.(CljsCoreIFn).X_invoke_Arity1(Apply.X_invoke_Arity5(h, x, y, z, args)).(interface{})).(interface{})
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
			return func(fs___1 interface{}) *AFn {
				return func(G__531 *AFn) *AFn {
					return Fn(G__531, func(args__ ...interface{}) interface{} {
						var args = Array_seq.X_invoke_Arity1(args__[0:])
						_ = args
						{
							var ret = Apply.X_invoke_Arity2(First.X_invoke_Arity1(fs___1), args)
							var fs___2 = Next.X_invoke_Arity1(fs___1)
							_, _ = ret, fs___2
							for {
								if Truth_(fs___2) {
									{
										var G__532 = First.X_invoke_Arity1(fs___2).X_invoke_Arity1(ret).(interface{})
										var G__533 = Next.X_invoke_Arity1(fs___2)
										ret = G__532
										fs___2 = G__533
										continue
									}
								} else {
									return ret
								}
							}
						}
					})
				}(&AFn{})
			}(fs___1)
		}
	})
}(&AFn{})

/**
* Takes a function f and fewer than the normal arguments to f, and
* returns a fn that takes a variable number of additional args. When
* called, the returned function calls f with args + additional args.
* @param {...*} var_args
 */
var Partial = func(partial *AFn) *AFn {
	return Fn(partial, func(f interface{}) interface{} {
		return f
	}, func(f interface{}, arg1 interface{}) interface{} {
		return func(G__534 *AFn) *AFn {
			return Fn(G__534, func(args__ ...interface{}) interface{} {
				var args = Array_seq.X_invoke_Arity1(args__[0:])
				_ = args
				return Apply.X_invoke_Arity3(f, arg1, args)
			})
		}(&AFn{})
	}, func(f interface{}, arg1 interface{}, arg2 interface{}) interface{} {
		return func(G__535 *AFn) *AFn {
			return Fn(G__535, func(args__ ...interface{}) interface{} {
				var args = Array_seq.X_invoke_Arity1(args__[0:])
				_ = args
				return Apply.X_invoke_Arity4(f, arg1, arg2, args)
			})
		}(&AFn{})
	}, func(f interface{}, arg1 interface{}, arg2 interface{}, arg3 interface{}) interface{} {
		return func(G__536 *AFn) *AFn {
			return Fn(G__536, func(args__ ...interface{}) interface{} {
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
		return func(G__537 *AFn) *AFn {
			return Fn(G__537, func(args__ ...interface{}) interface{} {
				var args = Array_seq.X_invoke_Arity1(args__[0:])
				_ = args
				return Apply.X_invoke_Arity5(f, arg1, arg2, arg3, Concat.X_invoke_Arity2(more, args).(*CljsCoreLazySeq))
			})
		}(&AFn{})
	})
}(&AFn{})

/**
* Takes a function f, and returns a function that calls f, replacing
* a nil first argument to f with the supplied value x. Higher arity
* versions can replace arguments in the second and third
* positions (y, z). Note that the function f can take any number of
* arguments, not just the one(s) being nil-patched.
 */
var Fnil = func(fnil *AFn) *AFn {
	return Fn(fnil, func(f interface{}, x interface{}) interface{} {
		return func(G__538 *AFn) *AFn {
			return Fn(G__538, func(a interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity1(func() interface{} {
					if a == nil {
						return x
					} else {
						return a
					}
				}()).(interface{})
			}, func(a interface{}, b interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity2(func() interface{} {
					if a == nil {
						return x
					} else {
						return a
					}
				}(), b).(interface{})
			}, func(a interface{}, b interface{}, c interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity3(func() interface{} {
					if a == nil {
						return x
					} else {
						return a
					}
				}(), b, c).(interface{})
			}, func(a_b_c_ds__ ...interface{}) interface{} {
				var a = a_b_c_ds__[0]
				var b = a_b_c_ds__[1]
				var c = a_b_c_ds__[2]
				var ds = Array_seq.X_invoke_Arity1(a_b_c_ds__[3:])
				_, _, _, _ = a, b, c, ds
				return Apply.X_invoke_Arity5(f, func() interface{} {
					if a == nil {
						return x
					} else {
						return a
					}
				}(), b, c, ds)
			})
		}(&AFn{})
	}, func(f interface{}, x interface{}, y interface{}) interface{} {
		return func(G__539 *AFn) *AFn {
			return Fn(G__539, func(a interface{}, b interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity2(func() interface{} {
					if a == nil {
						return x
					} else {
						return a
					}
				}(), func() interface{} {
					if b == nil {
						return y
					} else {
						return b
					}
				}()).(interface{})
			}, func(a interface{}, b interface{}, c interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity3(func() interface{} {
					if a == nil {
						return x
					} else {
						return a
					}
				}(), func() interface{} {
					if b == nil {
						return y
					} else {
						return b
					}
				}(), c).(interface{})
			}, func(a_b_c_ds__ ...interface{}) interface{} {
				var a = a_b_c_ds__[0]
				var b = a_b_c_ds__[1]
				var c = a_b_c_ds__[2]
				var ds = Array_seq.X_invoke_Arity1(a_b_c_ds__[3:])
				_, _, _, _ = a, b, c, ds
				return Apply.X_invoke_Arity5(f, func() interface{} {
					if a == nil {
						return x
					} else {
						return a
					}
				}(), func() interface{} {
					if b == nil {
						return y
					} else {
						return b
					}
				}(), c, ds)
			})
		}(&AFn{})
	}, func(f interface{}, x interface{}, y interface{}, z interface{}) interface{} {
		return func(G__540 *AFn) *AFn {
			return Fn(G__540, func(a interface{}, b interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity2(func() interface{} {
					if a == nil {
						return x
					} else {
						return a
					}
				}(), func() interface{} {
					if b == nil {
						return y
					} else {
						return b
					}
				}()).(interface{})
			}, func(a interface{}, b interface{}, c interface{}) interface{} {
				return f.(CljsCoreIFn).X_invoke_Arity3(func() interface{} {
					if a == nil {
						return x
					} else {
						return a
					}
				}(), func() interface{} {
					if b == nil {
						return y
					} else {
						return b
					}
				}(), func() interface{} {
					if c == nil {
						return z
					} else {
						return c
					}
				}()).(interface{})
			}, func(a_b_c_ds__ ...interface{}) interface{} {
				var a = a_b_c_ds__[0]
				var b = a_b_c_ds__[1]
				var c = a_b_c_ds__[2]
				var ds = Array_seq.X_invoke_Arity1(a_b_c_ds__[3:])
				_, _, _, _ = a, b, c, ds
				return Apply.X_invoke_Arity5(f, func() interface{} {
					if a == nil {
						return x
					} else {
						return a
					}
				}(), func() interface{} {
					if b == nil {
						return y
					} else {
						return b
					}
				}(), func() interface{} {
					if c == nil {
						return z
					} else {
						return c
					}
				}(), ds)
			})
		}(&AFn{})
	})
}(&AFn{})

/**
* Returns a lazy sequence consisting of the result of applying f to 0
* and the first item of coll, followed by applying f to 1 and the second
* item in coll, etc, until coll is exhausted. Thus function f should
* accept 2 arguments, index and item.
 */
var Map_indexed = func(map_indexed *AFn) *AFn {
	return Fn(map_indexed, func(f interface{}, coll interface{}) interface{} {
		{
			var mapi *AFn
			mapi = func(mapi *AFn) *AFn {
				return Fn(mapi, func(idx interface{}, coll___1 interface{}) interface{} {
					return (&CljsCoreLazySeq{nil, func(G__541 *AFn) *AFn {
						return Fn(G__541, func() interface{} {
							{
								var temp__4219__auto__ = Seq.Arity1IQ(coll___1)
								_ = temp__4219__auto__
								if Truth_(temp__4219__auto__) {
									{
										var s = temp__4219__auto__
										_ = s
										if Chunked_seq_QMARK_.X_invoke_Arity1(s) {
											{
												var c = Chunk_first.X_invoke_Arity1(s).(interface{})
												var size = Count.X_invoke_Arity1(c).(float64)
												var b = Chunk_buffer.X_invoke_Arity1(size).(*CljsCoreChunkBuffer)
												_, _, _ = c, size, b
												{
													var n__77528__auto___542 = size
													_ = n__77528__auto___542
													{
														var i_543 = float64(0)
														_ = i_543
														for {
															if i_543 < n__77528__auto___542 {
																Chunk_append.X_invoke_Arity2(b, f.(CljsCoreIFn).X_invoke_Arity2((idx.(float64)+i_543), c.X_nth_Arity2(i_543).(interface{})).(interface{})).(interface{})
																{
																	var G__544 = (i_543 + float64(1))
																	i_543 = G__544
																	continue
																}
															} else {
															}
															break
														}
													}
												}
												return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b).(interface{}), mapi.X_invoke_Arity2((idx.(float64)+size), Chunk_rest.X_invoke_Arity1(s).(interface{})).(*CljsCoreLazySeq))
											}
										} else {
											return Cons.X_invoke_Arity2(f.(CljsCoreIFn).X_invoke_Arity2(idx, First.X_invoke_Arity1(s)).(interface{}), mapi.X_invoke_Arity2((idx.(float64)+float64(1)), Rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
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

/**
* Returns a lazy sequence of the non-nil results of (f item). Note,
* this means false return values will be included.  f must be free of
* side-effects.  Returns a transducer when no collection is provided.
 */
var Keep = func(keep *AFn) *AFn {
	return Fn(keep, func(f interface{}) interface{} {
		return func(G__545 *AFn) *AFn {
			return Fn(G__545, func(f1 interface{}) interface{} {
				return func(G__546 *AFn) *AFn {
					return Fn(G__546, func() interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
					}, func(result interface{}) interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
					}, func(result interface{}, input interface{}) interface{} {
						{
							var v = f.(CljsCoreIFn).X_invoke_Arity1(input).(interface{})
							_ = v
							if v == nil {
								return result
							} else {
								return f1.(CljsCoreIFn).X_invoke_Arity2(result, v).(interface{})
							}
						}
					})
				}(&AFn{})
			})
		}(&AFn{})
	}, func(f interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__547 *AFn) *AFn {
			return Fn(G__547, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
							_ = s
							if Chunked_seq_QMARK_.X_invoke_Arity1(s) {
								{
									var c = Chunk_first.X_invoke_Arity1(s).(interface{})
									var size = Count.X_invoke_Arity1(c).(float64)
									var b = Chunk_buffer.X_invoke_Arity1(size).(*CljsCoreChunkBuffer)
									_, _, _ = c, size, b
									{
										var n__77528__auto___548 = size
										_ = n__77528__auto___548
										{
											var i_549 = float64(0)
											_ = i_549
											for {
												if i_549 < n__77528__auto___548 {
													{
														var x_550 = f.(CljsCoreIFn).X_invoke_Arity1(c.X_nth_Arity2(i_549).(interface{})).(interface{})
														_ = x_550
														if x_550 == nil {
														} else {
															Chunk_append.X_invoke_Arity2(b, x_550).(interface{})
														}
													}
													{
														var G__551 = (i_549 + float64(1))
														i_549 = G__551
														continue
													}
												} else {
												}
												break
											}
										}
									}
									return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b).(interface{}), keep.X_invoke_Arity2(f, Chunk_rest.X_invoke_Arity1(s).(interface{})).(*CljsCoreLazySeq))
								}
							} else {
								{
									var x = f.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(s)).(interface{})
									_ = x
									if x == nil {
										return keep.X_invoke_Arity2(f, Rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq)
									} else {
										return Cons.X_invoke_Arity2(x, keep.X_invoke_Arity2(f, Rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
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

type CljsCoreAtom struct {
	State     interface{}
	Meta      interface{}
	Validator interface{}
	Watches   interface{}
}

func (_ *CljsCoreAtom) CljsCoreIHash__() {}
func (self__ *CljsCoreAtom) X_hash_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		return goog.GetUid(this___1)
	}
}

func (_ *CljsCoreAtom) CljsCoreIWatchable__() {}
func (self__ *CljsCoreAtom) X_notify_watches_Arity3(oldval interface{}, newval interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		{
			var seq__558 = Seq.Arity1IQ(self__.Watches)
			var chunk__559 = nil
			var count__560 = float64(0)
			var i__561 = float64(0)
			_, _, _, _ = seq__558, chunk__559, count__560, i__561
			for {
				if i__561 < count__560 {
					{
						var vec__562 = chunk__559.X_nth_Arity2(i__561).(interface{})
						var key = Nth.X_invoke_Arity3(vec__562, float64(0), nil)
						var f = Nth.X_invoke_Arity3(vec__562, float64(1), nil)
						_, _, _ = vec__562, key, f
						f.(CljsCoreIFn).X_invoke_Arity4(key, this___1, oldval, newval).(interface{})
						{
							var G__564 = seq__558
							var G__565 = chunk__559
							var G__566 = count__560
							var G__567 = (i__561 + float64(1))
							seq__558 = G__564
							chunk__559 = G__565
							count__560 = G__566
							i__561 = G__567
							continue
						}
					}
				} else {
					{
						var temp__4219__auto__ = Seq.X_invoke_Arity1(seq__558)
						_ = temp__4219__auto__
						if Truth_(temp__4219__auto__) {
							{
								var seq__558___1 = temp__4219__auto__
								_ = seq__558___1
								if Chunked_seq_QMARK_.X_invoke_Arity1(seq__558___1) {
									{
										var c__77428__auto__ = Chunk_first.X_invoke_Arity1(seq__558___1).(interface{})
										_ = c__77428__auto__
										{
											var G__568 = Chunk_rest.X_invoke_Arity1(seq__558___1).(interface{})
											var G__569 = c__77428__auto__
											var G__570 = Count.X_invoke_Arity1(c__77428__auto__).(float64)
											var G__571 = float64(0)
											seq__558 = G__568
											chunk__559 = G__569
											count__560 = G__570
											i__561 = G__571
											continue
										}
									}
								} else {
									{
										var vec__563 = First.X_invoke_Arity1(seq__558___1)
										var key = Nth.X_invoke_Arity3(vec__563, float64(0), nil)
										var f = Nth.X_invoke_Arity3(vec__563, float64(1), nil)
										_, _, _ = vec__563, key, f
										f.(CljsCoreIFn).X_invoke_Arity4(key, this___1, oldval, newval).(interface{})
										{
											var G__572 = Next.X_invoke_Arity1(seq__558___1)
											var G__573 = nil
											var G__574 = float64(0)
											var G__575 = float64(0)
											seq__558 = G__572
											chunk__559 = G__573
											count__560 = G__574
											i__561 = G__575
											continue
										}
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
}

func (self__ *CljsCoreAtom) X_add_watch_Arity3(key interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		Native_set_instance_field.X_invoke_Arity3(this___1, "Watches", Assoc.X_invoke_Arity3(self__.Watches, key, f).(interface{}))

		return this___1
	}
}

func (self__ *CljsCoreAtom) X_remove_watch_Arity2(key interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		return func() interface{} {
			var return__576 = Dissoc.X_invoke_Arity2(self__.Watches, key)
			Native_set_instance_field.X_invoke_Arity3(this___1, "Watches", return__576)
			return return__576
		}()
	}
}

func (_ *CljsCoreAtom) CljsCoreIMeta__() {}
func (self__ *CljsCoreAtom) X_meta_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Meta
	}
}

func (_ *CljsCoreAtom) CljsCoreIDeref__() {}
func (self__ *CljsCoreAtom) X_deref_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.State
	}
}

func (_ *CljsCoreAtom) CljsCoreIEquiv__() {}
func (self__ *CljsCoreAtom) X_equiv_Arity2(other interface{}) bool {
	{
		var o___1 = self__
		_ = o___1
		return (o___1 == other)
	}
}

func (_ *CljsCoreAtom) CljsCoreIAtom__()  {}
func (_ *CljsCoreAtom) CljsCoreObject__() {}
func (self__ *CljsCoreAtom) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

var X__GT_Atom = func(__GT_Atom *AFn) *AFn {
	return Fn(__GT_Atom, func(state interface{}, meta interface{}, validator interface{}, watches interface{}) interface{} {
		return (&CljsCoreAtom{state, meta, validator, watches})
	})
}(&AFn{})

/**
* Creates and returns an Atom with an initial value of x and zero or
* more options (in any order):
*
* :meta metadata-map
*
* :validator validate-fn
*
* If metadata-map is supplied, it will be come the metadata on the
* atom. validate-fn must be nil or a side-effect-free fn of one
* argument, which will be passed the intended new state on any state
* change. If the new state is unacceptable, the validate-fn should
* return false or throw an Error.  If either of these error conditions
* occur, then the value of the atom will not change.
* @param {...*} var_args
 */
var Atom = func(atom *AFn) *AFn {
	return Fn(atom, func(x interface{}) interface{} {
		return (&CljsCoreAtom{x, nil, nil, nil})
	}, func(x_p__577__ ...interface{}) interface{} {
		var x = x_p__577__[0]
		var p__577 = Array_seq.X_invoke_Arity1(x_p__577__[1:])
		_, _ = x, p__577
		{
			var map__579 = p__577
			var map__579___1 = func() interface{} {
				if Seq_QMARK_.Arity1IB(map__579) {
					return Apply.X_invoke_Arity2(Hash_map, map__579)
				} else {
					return map__579
				}
			}()
			var validator = Get.X_invoke_Arity2(map__579___1, (&CljsCoreKeyword{Ns: nil, Name: "validator", Fqn: "validator", X_hash: float64(-1966190681)}))
			var meta = Get.X_invoke_Arity2(map__579___1, (&CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}))
			_, _, _, _ = map__579, map__579___1, validator, meta
			return (&CljsCoreAtom{x, meta, validator, nil})
		}
	})
}(&AFn{})

/**
* Sets the value of atom to newval without regard for the
* current value. Returns newval.
 */
var Reset_BANG_ = func(reset_BANG_ *AFn) *AFn {
	return Fn(reset_BANG_, func(a interface{}, new_value interface{}) interface{} {
		if func() bool { _, instanceof := a.(*CljsCoreAtom); return instanceof }() {
			{
				var validate = Native_get_instance_field.X_invoke_Arity2(a, "Validator")
				_ = validate
				if validate == nil {
				} else {
					if Truth_(validate.(CljsCoreIFn).X_invoke_Arity1(new_value).(interface{})) {
					} else {
						panic((&js.Error{("Assert failed: Validator rejected reference state\n" + Str.X_invoke_Arity1(Pr_str.X_invoke_Arity1(List((&CljsCoreSymbol{Ns: nil, Name: "validate", Str: "validate", X_hash: float64(1439230700), X_meta: nil}), (&CljsCoreSymbol{Ns: nil, Name: "new-value", Str: "new-value", X_hash: float64(-1567397401), X_meta: nil}))).(interface{})).(string))}))
					}
				}
				{
					var old_value = Native_get_instance_field.X_invoke_Arity2(a, "State")
					_ = old_value
					Native_set_instance_field.X_invoke_Arity3(a, "State", new_value)

					if Native_get_instance_field.X_invoke_Arity2(a, "Watches") == nil {
					} else {
						a.(CljsCoreIWatchable).X_notify_watches_Arity3(old_value, new_value).(interface{})
					}
					return new_value
				}
			}
		} else {
			return a.(CljsCoreIReset).X_reset_BANG__Arity2(new_value).(interface{})
		}
	})
}(&AFn{})

/**
* Atomically swaps the value of atom to be:
* (apply f current-value-of-atom args). Note that f may be called
* multiple times, and thus should be free of side effects.  Returns
* the value that was swapped in.
* @param {...*} var_args
 */
var Swap_BANG_ = func(swap_BANG_ *AFn) *AFn {
	return Fn(swap_BANG_, func(a interface{}, f interface{}) interface{} {
		if func() bool { _, instanceof := a.(*CljsCoreAtom); return instanceof }() {
			return Reset_BANG_.X_invoke_Arity2(a, f.(CljsCoreIFn).X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(a, "State")).(interface{}))
		} else {
			return a.(CljsCoreISwap).X_swap_BANG__Arity2(f).(interface{})
		}
	}, func(a interface{}, f interface{}, x interface{}) interface{} {
		if func() bool { _, instanceof := a.(*CljsCoreAtom); return instanceof }() {
			return Reset_BANG_.X_invoke_Arity2(a, f.(CljsCoreIFn).X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(a, "State"), x).(interface{}))
		} else {
			return a.(CljsCoreISwap).X_swap_BANG__Arity3(f, x).(interface{})
		}
	}, func(a interface{}, f interface{}, x interface{}, y interface{}) interface{} {
		if func() bool { _, instanceof := a.(*CljsCoreAtom); return instanceof }() {
			return Reset_BANG_.X_invoke_Arity2(a, f.(CljsCoreIFn).X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(a, "State"), x, y).(interface{}))
		} else {
			return a.(CljsCoreISwap).X_swap_BANG__Arity4(f, x, y).(interface{})
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
			return a.(CljsCoreISwap).X_swap_BANG__Arity5(f, x, y, more).(interface{})
		}
	})
}(&AFn{})

/**
* Atomically sets the value of atom to newval if and only if the
* current value of the atom is identical to oldval. Returns true if
* set happened, else false.
 */
var Compare_and_set_BANG_ = func(compare_and_set_BANG_ *AFn) *AFn {
	return Fn(compare_and_set_BANG_, func(a interface{}, oldval interface{}, newval interface{}) interface{} {
		if X_EQ_.Arity2IIB(Native_get_instance_field.X_invoke_Arity2(a, "State"), oldval) {
			Reset_BANG_.X_invoke_Arity2(a, newval)
			return true
		} else {
			return false
		}
	})
}(&AFn{})

/**
* Sets the validator-fn for an atom. validator-fn must be nil or a
* side-effect-free fn of one argument, which will be passed the intended
* new state on any state change. If the new state is unacceptable, the
* validator-fn should return false or throw an Error. If the current state
* is not acceptable to the new validator, an Error will be thrown and the
* validator will not be changed.
 */
var Set_validator_BANG_ = func(set_validator_BANG_ *AFn) *AFn {
	return Fn(set_validator_BANG_, func(iref interface{}, val interface{}) interface{} {
		return func() interface{} {
			var return__580 = val
			Native_set_instance_field.X_invoke_Arity3(iref, "Validator", return__580)
			return return__580
		}()
	})
}(&AFn{})

/**
* Gets the validator-fn for a var/ref/agent/atom.
 */
var Get_validator = func(get_validator *AFn) *AFn {
	return Fn(get_validator, func(iref interface{}) interface{} {
		return Native_get_instance_field.X_invoke_Arity2(iref, "Validator")
	})
}(&AFn{})

/**
* Returns a lazy sequence of the non-nil results of (f index item). Note,
* this means false return values will be included.  f must be free of
* side-effects.  Returns a stateful transducer when no collection is
* provided.
 */
var Keep_indexed = func(keep_indexed *AFn) *AFn {
	return Fn(keep_indexed, func(f interface{}) interface{} {
		return func(G__581 *AFn) *AFn {
			return Fn(G__581, func(f1 interface{}) interface{} {
				{
					var ia = Atom.X_invoke_Arity1(float64(-1)).(*CljsCoreAtom)
					_ = ia
					return func(ia *CljsCoreAtom) *AFn {
						return func(G__582 *AFn) *AFn {
							return Fn(G__582, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
							}, func(result interface{}, input interface{}) interface{} {
								{
									var i = Swap_BANG_.X_invoke_Arity2(ia, Inc)
									var v = f.(CljsCoreIFn).X_invoke_Arity2(i, input).(interface{})
									_, _ = i, v
									if v == nil {
										return result
									} else {
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, v).(interface{})
									}
								}
							})
						}(&AFn{})
					}(ia)
				}
			})
		}(&AFn{})
	}, func(f interface{}, coll interface{}) interface{} {
		{
			var keepi *AFn
			keepi = func(keepi *AFn) *AFn {
				return Fn(keepi, func(idx interface{}, coll___1 interface{}) interface{} {
					return (&CljsCoreLazySeq{nil, func(G__583 *AFn) *AFn {
						return Fn(G__583, func() interface{} {
							{
								var temp__4219__auto__ = Seq.Arity1IQ(coll___1)
								_ = temp__4219__auto__
								if Truth_(temp__4219__auto__) {
									{
										var s = temp__4219__auto__
										_ = s
										if Chunked_seq_QMARK_.X_invoke_Arity1(s) {
											{
												var c = Chunk_first.X_invoke_Arity1(s).(interface{})
												var size = Count.X_invoke_Arity1(c).(float64)
												var b = Chunk_buffer.X_invoke_Arity1(size).(*CljsCoreChunkBuffer)
												_, _, _ = c, size, b
												{
													var n__77528__auto___584 = size
													_ = n__77528__auto___584
													{
														var i_585 = float64(0)
														_ = i_585
														for {
															if i_585 < n__77528__auto___584 {
																{
																	var x_586 = f.(CljsCoreIFn).X_invoke_Arity2((idx.(float64) + i_585), c.X_nth_Arity2(i_585).(interface{})).(interface{})
																	_ = x_586
																	if x_586 == nil {
																	} else {
																		Chunk_append.X_invoke_Arity2(b, x_586).(interface{})
																	}
																}
																{
																	var G__587 = (i_585 + float64(1))
																	i_585 = G__587
																	continue
																}
															} else {
															}
															break
														}
													}
												}
												return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b).(interface{}), keepi.X_invoke_Arity2((idx.(float64)+size), Chunk_rest.X_invoke_Arity1(s).(interface{})).(*CljsCoreLazySeq))
											}
										} else {
											{
												var x = f.(CljsCoreIFn).X_invoke_Arity2(idx, First.X_invoke_Arity1(s)).(interface{})
												_ = x
												if x == nil {
													return keepi.X_invoke_Arity2((idx.(float64) + float64(1)), Rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq)
												} else {
													return Cons.X_invoke_Arity2(x, keepi.X_invoke_Arity2((idx.(float64)+float64(1)), Rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
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

/**
* Takes a set of predicates and returns a function f that returns true if all of its
* composing predicates return a logical true value against all of its arguments, else it returns
* false. Note that f is short-circuiting in that it will stop execution on the first
* argument that triggers a logical false result against the original predicates.
* @param {...*} var_args
 */
var Every_pred = func(every_pred *AFn) *AFn {
	return Fn(every_pred, func(p interface{}) interface{} {
		return func(ep1 *AFn) *AFn {
			return Fn(ep1, func() interface{} {
				return true
			}, func(x interface{}) interface{} {
				return Boolean.X_invoke_Arity1(p.(CljsCoreIFn).X_invoke_Arity1(x).(interface{}))
			}, func(x interface{}, y interface{}) interface{} {
				return Boolean.X_invoke_Arity1(func() interface{} {
					var and__76620__auto__ = p.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = and__76620__auto__
					if Truth_(and__76620__auto__) {
						return p.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
					} else {
						return and__76620__auto__
					}
				}())
			}, func(x interface{}, y interface{}, z interface{}) interface{} {
				return Boolean.X_invoke_Arity1(func() interface{} {
					var and__76620__auto__ = p.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = and__76620__auto__
					if Truth_(and__76620__auto__) {
						{
							var and__76620__auto_____1 = p.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
							_ = and__76620__auto_____1
							if Truth_(and__76620__auto_____1) {
								return p.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
							} else {
								return and__76620__auto_____1
							}
						}
					} else {
						return and__76620__auto__
					}
				}())
			}, func(x_y_z_args__ ...interface{}) interface{} {
				var x = x_y_z_args__[0]
				var y = x_y_z_args__[1]
				var z = x_y_z_args__[2]
				var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
				_, _, _, _ = x, y, z, args
				return Boolean.X_invoke_Arity1((ep1.X_invoke_Arity3(x, y, z).(bool)) && (Every_QMARK_.Arity2IIB(p, args)))
			})
		}(&AFn{})
	}, func(p1 interface{}, p2 interface{}) interface{} {
		return func(ep2 *AFn) *AFn {
			return Fn(ep2, func() interface{} {
				return true
			}, func(x interface{}) interface{} {
				return Boolean.X_invoke_Arity1(func() interface{} {
					var and__76620__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = and__76620__auto__
					if Truth_(and__76620__auto__) {
						return p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					} else {
						return and__76620__auto__
					}
				}())
			}, func(x interface{}, y interface{}) interface{} {
				return Boolean.X_invoke_Arity1(func() interface{} {
					var and__76620__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = and__76620__auto__
					if Truth_(and__76620__auto__) {
						{
							var and__76620__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
							_ = and__76620__auto_____1
							if Truth_(and__76620__auto_____1) {
								{
									var and__76620__auto_____2 = p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
									_ = and__76620__auto_____2
									if Truth_(and__76620__auto_____2) {
										return p2.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
									} else {
										return and__76620__auto_____2
									}
								}
							} else {
								return and__76620__auto_____1
							}
						}
					} else {
						return and__76620__auto__
					}
				}())
			}, func(x interface{}, y interface{}, z interface{}) interface{} {
				return Boolean.X_invoke_Arity1(func() interface{} {
					var and__76620__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = and__76620__auto__
					if Truth_(and__76620__auto__) {
						{
							var and__76620__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
							_ = and__76620__auto_____1
							if Truth_(and__76620__auto_____1) {
								{
									var and__76620__auto_____2 = p1.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
									_ = and__76620__auto_____2
									if Truth_(and__76620__auto_____2) {
										{
											var and__76620__auto_____3 = p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
											_ = and__76620__auto_____3
											if Truth_(and__76620__auto_____3) {
												{
													var and__76620__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
													_ = and__76620__auto_____4
													if Truth_(and__76620__auto_____4) {
														return p2.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
													} else {
														return and__76620__auto_____4
													}
												}
											} else {
												return and__76620__auto_____3
											}
										}
									} else {
										return and__76620__auto_____2
									}
								}
							} else {
								return and__76620__auto_____1
							}
						}
					} else {
						return and__76620__auto__
					}
				}())
			}, func(x_y_z_args__ ...interface{}) interface{} {
				var x = x_y_z_args__[0]
				var y = x_y_z_args__[1]
				var z = x_y_z_args__[2]
				var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
				_, _, _, _ = x, y, z, args
				return Boolean.X_invoke_Arity1((ep2.X_invoke_Arity3(x, y, z).(bool)) && (Every_QMARK_.X_invoke_Arity2(func(G__594 *AFn) *AFn {
					return Fn(G__594, func(p1__588_SHARP_ interface{}) interface{} {
						{
							var and__76620__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(p1__588_SHARP_).(interface{})
							_ = and__76620__auto__
							if Truth_(and__76620__auto__) {
								return p2.(CljsCoreIFn).X_invoke_Arity1(p1__588_SHARP_).(interface{})
							} else {
								return and__76620__auto__
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
				return Boolean.X_invoke_Arity1(func() interface{} {
					var and__76620__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = and__76620__auto__
					if Truth_(and__76620__auto__) {
						{
							var and__76620__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
							_ = and__76620__auto_____1
							if Truth_(and__76620__auto_____1) {
								return p3.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
							} else {
								return and__76620__auto_____1
							}
						}
					} else {
						return and__76620__auto__
					}
				}())
			}, func(x interface{}, y interface{}) interface{} {
				return Boolean.X_invoke_Arity1(func() interface{} {
					var and__76620__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = and__76620__auto__
					if Truth_(and__76620__auto__) {
						{
							var and__76620__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
							_ = and__76620__auto_____1
							if Truth_(and__76620__auto_____1) {
								{
									var and__76620__auto_____2 = p3.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
									_ = and__76620__auto_____2
									if Truth_(and__76620__auto_____2) {
										{
											var and__76620__auto_____3 = p1.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
											_ = and__76620__auto_____3
											if Truth_(and__76620__auto_____3) {
												{
													var and__76620__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
													_ = and__76620__auto_____4
													if Truth_(and__76620__auto_____4) {
														return p3.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
													} else {
														return and__76620__auto_____4
													}
												}
											} else {
												return and__76620__auto_____3
											}
										}
									} else {
										return and__76620__auto_____2
									}
								}
							} else {
								return and__76620__auto_____1
							}
						}
					} else {
						return and__76620__auto__
					}
				}())
			}, func(x interface{}, y interface{}, z interface{}) interface{} {
				return Boolean.X_invoke_Arity1(func() interface{} {
					var and__76620__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = and__76620__auto__
					if Truth_(and__76620__auto__) {
						{
							var and__76620__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
							_ = and__76620__auto_____1
							if Truth_(and__76620__auto_____1) {
								{
									var and__76620__auto_____2 = p3.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
									_ = and__76620__auto_____2
									if Truth_(and__76620__auto_____2) {
										{
											var and__76620__auto_____3 = p1.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
											_ = and__76620__auto_____3
											if Truth_(and__76620__auto_____3) {
												{
													var and__76620__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
													_ = and__76620__auto_____4
													if Truth_(and__76620__auto_____4) {
														{
															var and__76620__auto_____5 = p3.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
															_ = and__76620__auto_____5
															if Truth_(and__76620__auto_____5) {
																{
																	var and__76620__auto_____6 = p1.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
																	_ = and__76620__auto_____6
																	if Truth_(and__76620__auto_____6) {
																		{
																			var and__76620__auto_____7 = p2.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
																			_ = and__76620__auto_____7
																			if Truth_(and__76620__auto_____7) {
																				return p3.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
																			} else {
																				return and__76620__auto_____7
																			}
																		}
																	} else {
																		return and__76620__auto_____6
																	}
																}
															} else {
																return and__76620__auto_____5
															}
														}
													} else {
														return and__76620__auto_____4
													}
												}
											} else {
												return and__76620__auto_____3
											}
										}
									} else {
										return and__76620__auto_____2
									}
								}
							} else {
								return and__76620__auto_____1
							}
						}
					} else {
						return and__76620__auto__
					}
				}())
			}, func(x_y_z_args__ ...interface{}) interface{} {
				var x = x_y_z_args__[0]
				var y = x_y_z_args__[1]
				var z = x_y_z_args__[2]
				var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
				_, _, _, _ = x, y, z, args
				return Boolean.X_invoke_Arity1((ep3.X_invoke_Arity3(x, y, z).(bool)) && (Every_QMARK_.X_invoke_Arity2(func(G__595 *AFn) *AFn {
					return Fn(G__595, func(p1__589_SHARP_ interface{}) interface{} {
						{
							var and__76620__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(p1__589_SHARP_).(interface{})
							_ = and__76620__auto__
							if Truth_(and__76620__auto__) {
								{
									var and__76620__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(p1__589_SHARP_).(interface{})
									_ = and__76620__auto_____1
									if Truth_(and__76620__auto_____1) {
										return p3.(CljsCoreIFn).X_invoke_Arity1(p1__589_SHARP_).(interface{})
									} else {
										return and__76620__auto_____1
									}
								}
							} else {
								return and__76620__auto__
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
			return func(ps___1 *CljsCoreCons) *AFn {
				return func(epn *AFn) *AFn {
					return Fn(epn, func() interface{} {
						return true
					}, func(x interface{}) interface{} {
						return Every_QMARK_.X_invoke_Arity2(func(ps___1 *CljsCoreCons) *AFn {
							return func(G__596 *AFn) *AFn {
								return Fn(G__596, func(p1__590_SHARP_ interface{}) interface{} {
									return p1__590_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
								})
							}(&AFn{})
						}(ps___1), ps___1)
					}, func(x interface{}, y interface{}) interface{} {
						return Every_QMARK_.X_invoke_Arity2(func(ps___1 *CljsCoreCons) *AFn {
							return func(G__597 *AFn) *AFn {
								return Fn(G__597, func(p1__591_SHARP_ interface{}) interface{} {
									{
										var and__76620__auto__ = p1__591_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
										_ = and__76620__auto__
										if Truth_(and__76620__auto__) {
											return p1__591_SHARP_.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
										} else {
											return and__76620__auto__
										}
									}
								})
							}(&AFn{})
						}(ps___1), ps___1)
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return Every_QMARK_.X_invoke_Arity2(func(ps___1 *CljsCoreCons) *AFn {
							return func(G__598 *AFn) *AFn {
								return Fn(G__598, func(p1__592_SHARP_ interface{}) interface{} {
									{
										var and__76620__auto__ = p1__592_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
										_ = and__76620__auto__
										if Truth_(and__76620__auto__) {
											{
												var and__76620__auto_____1 = p1__592_SHARP_.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
												_ = and__76620__auto_____1
												if Truth_(and__76620__auto_____1) {
													return p1__592_SHARP_.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
												} else {
													return and__76620__auto_____1
												}
											}
										} else {
											return and__76620__auto__
										}
									}
								})
							}(&AFn{})
						}(ps___1), ps___1)
					}, func(x_y_z_args__ ...interface{}) interface{} {
						var x = x_y_z_args__[0]
						var y = x_y_z_args__[1]
						var z = x_y_z_args__[2]
						var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
						_, _, _, _ = x, y, z, args
						return Boolean.X_invoke_Arity1((epn.X_invoke_Arity3(x, y, z).(bool)) && (Every_QMARK_.X_invoke_Arity2(func(ps___1 *CljsCoreCons) *AFn {
							return func(G__599 *AFn) *AFn {
								return Fn(G__599, func(p1__593_SHARP_ interface{}) interface{} {
									return Every_QMARK_.Arity2IIB(p1__593_SHARP_, args)
								})
							}(&AFn{})
						}(ps___1), ps___1)))
					})
				}(&AFn{})
			}(ps___1)
		}
	})
}(&AFn{})

/**
* Takes a set of predicates and returns a function f that returns the first logical true value
* returned by one of its composing predicates against any of its arguments, else it returns
* logical false. Note that f is short-circuiting in that it will stop execution on the first
* argument that triggers a logical true result against the original predicates.
* @param {...*} var_args
 */
var Some_fn = func(some_fn *AFn) *AFn {
	return Fn(some_fn, func(p interface{}) interface{} {
		return func(sp1 *AFn) *AFn {
			return Fn(sp1, func() interface{} {
				return nil
			}, func(x interface{}) interface{} {
				return p.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
			}, func(x interface{}, y interface{}) interface{} {
				{
					var or__76632__auto__ = p.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						return p.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
					}
				}
			}, func(x interface{}, y interface{}, z interface{}) interface{} {
				{
					var or__76632__auto__ = p.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						{
							var or__76632__auto_____1 = p.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
							_ = or__76632__auto_____1
							if Truth_(or__76632__auto_____1) {
								return or__76632__auto_____1
							} else {
								return p.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
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
					var or__76632__auto__ = sp1.X_invoke_Arity3(x, y, z).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
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
					var or__76632__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						return p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					}
				}
			}, func(x interface{}, y interface{}) interface{} {
				{
					var or__76632__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						{
							var or__76632__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
							_ = or__76632__auto_____1
							if Truth_(or__76632__auto_____1) {
								return or__76632__auto_____1
							} else {
								{
									var or__76632__auto_____2 = p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
									_ = or__76632__auto_____2
									if Truth_(or__76632__auto_____2) {
										return or__76632__auto_____2
									} else {
										return p2.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
									}
								}
							}
						}
					}
				}
			}, func(x interface{}, y interface{}, z interface{}) interface{} {
				{
					var or__76632__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						{
							var or__76632__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
							_ = or__76632__auto_____1
							if Truth_(or__76632__auto_____1) {
								return or__76632__auto_____1
							} else {
								{
									var or__76632__auto_____2 = p1.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
									_ = or__76632__auto_____2
									if Truth_(or__76632__auto_____2) {
										return or__76632__auto_____2
									} else {
										{
											var or__76632__auto_____3 = p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
											_ = or__76632__auto_____3
											if Truth_(or__76632__auto_____3) {
												return or__76632__auto_____3
											} else {
												{
													var or__76632__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
													_ = or__76632__auto_____4
													if Truth_(or__76632__auto_____4) {
														return or__76632__auto_____4
													} else {
														return p2.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
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
					var or__76632__auto__ = sp2.X_invoke_Arity3(x, y, z).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						return Some.X_invoke_Arity2(func(or__76632__auto__ interface{}) *AFn {
							return func(G__606 *AFn) *AFn {
								return Fn(G__606, func(p1__600_SHARP_ interface{}) interface{} {
									{
										var or__76632__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(p1__600_SHARP_).(interface{})
										_ = or__76632__auto_____1
										if Truth_(or__76632__auto_____1) {
											return or__76632__auto_____1
										} else {
											return p2.(CljsCoreIFn).X_invoke_Arity1(p1__600_SHARP_).(interface{})
										}
									}
								})
							}(&AFn{})
						}(or__76632__auto__), args)
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
					var or__76632__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						{
							var or__76632__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
							_ = or__76632__auto_____1
							if Truth_(or__76632__auto_____1) {
								return or__76632__auto_____1
							} else {
								return p3.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
							}
						}
					}
				}
			}, func(x interface{}, y interface{}) interface{} {
				{
					var or__76632__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						{
							var or__76632__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
							_ = or__76632__auto_____1
							if Truth_(or__76632__auto_____1) {
								return or__76632__auto_____1
							} else {
								{
									var or__76632__auto_____2 = p3.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
									_ = or__76632__auto_____2
									if Truth_(or__76632__auto_____2) {
										return or__76632__auto_____2
									} else {
										{
											var or__76632__auto_____3 = p1.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
											_ = or__76632__auto_____3
											if Truth_(or__76632__auto_____3) {
												return or__76632__auto_____3
											} else {
												{
													var or__76632__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
													_ = or__76632__auto_____4
													if Truth_(or__76632__auto_____4) {
														return or__76632__auto_____4
													} else {
														return p3.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
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
					var or__76632__auto__ = p1.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						{
							var or__76632__auto_____1 = p2.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
							_ = or__76632__auto_____1
							if Truth_(or__76632__auto_____1) {
								return or__76632__auto_____1
							} else {
								{
									var or__76632__auto_____2 = p3.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
									_ = or__76632__auto_____2
									if Truth_(or__76632__auto_____2) {
										return or__76632__auto_____2
									} else {
										{
											var or__76632__auto_____3 = p1.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
											_ = or__76632__auto_____3
											if Truth_(or__76632__auto_____3) {
												return or__76632__auto_____3
											} else {
												{
													var or__76632__auto_____4 = p2.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
													_ = or__76632__auto_____4
													if Truth_(or__76632__auto_____4) {
														return or__76632__auto_____4
													} else {
														{
															var or__76632__auto_____5 = p3.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
															_ = or__76632__auto_____5
															if Truth_(or__76632__auto_____5) {
																return or__76632__auto_____5
															} else {
																{
																	var or__76632__auto_____6 = p1.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
																	_ = or__76632__auto_____6
																	if Truth_(or__76632__auto_____6) {
																		return or__76632__auto_____6
																	} else {
																		{
																			var or__76632__auto_____7 = p2.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
																			_ = or__76632__auto_____7
																			if Truth_(or__76632__auto_____7) {
																				return or__76632__auto_____7
																			} else {
																				return p3.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
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
					var or__76632__auto__ = sp3.X_invoke_Arity3(x, y, z).(interface{})
					_ = or__76632__auto__
					if Truth_(or__76632__auto__) {
						return or__76632__auto__
					} else {
						return Some.X_invoke_Arity2(func(or__76632__auto__ interface{}) *AFn {
							return func(G__607 *AFn) *AFn {
								return Fn(G__607, func(p1__601_SHARP_ interface{}) interface{} {
									{
										var or__76632__auto_____1 = p1.(CljsCoreIFn).X_invoke_Arity1(p1__601_SHARP_).(interface{})
										_ = or__76632__auto_____1
										if Truth_(or__76632__auto_____1) {
											return or__76632__auto_____1
										} else {
											{
												var or__76632__auto_____2 = p2.(CljsCoreIFn).X_invoke_Arity1(p1__601_SHARP_).(interface{})
												_ = or__76632__auto_____2
												if Truth_(or__76632__auto_____2) {
													return or__76632__auto_____2
												} else {
													return p3.(CljsCoreIFn).X_invoke_Arity1(p1__601_SHARP_).(interface{})
												}
											}
										}
									}
								})
							}(&AFn{})
						}(or__76632__auto__), args)
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
			return func(ps___1 *CljsCoreCons) *AFn {
				return func(spn *AFn) *AFn {
					return Fn(spn, func() interface{} {
						return nil
					}, func(x interface{}) interface{} {
						return Some.X_invoke_Arity2(func(ps___1 *CljsCoreCons) *AFn {
							return func(G__608 *AFn) *AFn {
								return Fn(G__608, func(p1__602_SHARP_ interface{}) interface{} {
									return p1__602_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
								})
							}(&AFn{})
						}(ps___1), ps___1)
					}, func(x interface{}, y interface{}) interface{} {
						return Some.X_invoke_Arity2(func(ps___1 *CljsCoreCons) *AFn {
							return func(G__609 *AFn) *AFn {
								return Fn(G__609, func(p1__603_SHARP_ interface{}) interface{} {
									{
										var or__76632__auto__ = p1__603_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
										_ = or__76632__auto__
										if Truth_(or__76632__auto__) {
											return or__76632__auto__
										} else {
											return p1__603_SHARP_.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
										}
									}
								})
							}(&AFn{})
						}(ps___1), ps___1)
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return Some.X_invoke_Arity2(func(ps___1 *CljsCoreCons) *AFn {
							return func(G__610 *AFn) *AFn {
								return Fn(G__610, func(p1__604_SHARP_ interface{}) interface{} {
									{
										var or__76632__auto__ = p1__604_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
										_ = or__76632__auto__
										if Truth_(or__76632__auto__) {
											return or__76632__auto__
										} else {
											{
												var or__76632__auto_____1 = p1__604_SHARP_.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
												_ = or__76632__auto_____1
												if Truth_(or__76632__auto_____1) {
													return or__76632__auto_____1
												} else {
													return p1__604_SHARP_.(CljsCoreIFn).X_invoke_Arity1(z).(interface{})
												}
											}
										}
									}
								})
							}(&AFn{})
						}(ps___1), ps___1)
					}, func(x_y_z_args__ ...interface{}) interface{} {
						var x = x_y_z_args__[0]
						var y = x_y_z_args__[1]
						var z = x_y_z_args__[2]
						var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
						_, _, _, _ = x, y, z, args
						{
							var or__76632__auto__ = spn.X_invoke_Arity3(x, y, z)
							_ = or__76632__auto__
							if Truth_(or__76632__auto__) {
								return or__76632__auto__
							} else {
								return Some.X_invoke_Arity2(func(or__76632__auto__ interface{}, ps___1 *CljsCoreCons) *AFn {
									return func(G__611 *AFn) *AFn {
										return Fn(G__611, func(p1__605_SHARP_ interface{}) interface{} {
											return Some.X_invoke_Arity2(p1__605_SHARP_, args)
										})
									}(&AFn{})
								}(or__76632__auto__, ps___1), ps___1)
							}
						}
					})
				}(&AFn{})
			}(ps___1)
		}
	})
}(&AFn{})

/**
* Returns a lazy sequence consisting of the result of applying f to
* the set of first items of each coll, followed by applying f to the
* set of second items in each coll, until any one of the colls is
* exhausted.  Any remaining items in other colls are ignored. Function
* f should accept number-of-colls arguments. Returns a transducer when
* no collection is provided.
* @param {...*} var_args
 */
var Map_ = func(map_ *AFn) *AFn {
	return Fn(map_, func(f interface{}) interface{} {
		return func(G__613 *AFn) *AFn {
			return Fn(G__613, func(f1 interface{}) interface{} {
				return func(G__614 *AFn) *AFn {
					return Fn(G__614, func() interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
					}, func(result interface{}) interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
					}, func(result interface{}, input interface{}) interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity2(result, f.(CljsCoreIFn).X_invoke_Arity1(input).(interface{})).(interface{})
					}, func(result_input_inputs__ ...interface{}) interface{} {
						var result = result_input_inputs__[0]
						var input = result_input_inputs__[1]
						var inputs = Array_seq.X_invoke_Arity1(result_input_inputs__[2:])
						_, _, _ = result, input, inputs
						return f1.(CljsCoreIFn).X_invoke_Arity2(result, Apply.X_invoke_Arity3(f, input, inputs)).(interface{})
					})
				}(&AFn{})
			})
		}(&AFn{})
	}, func(f interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__615 *AFn) *AFn {
			return Fn(G__615, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
							_ = s
							if Chunked_seq_QMARK_.X_invoke_Arity1(s) {
								{
									var c = Chunk_first.X_invoke_Arity1(s).(interface{})
									var size = Count.X_invoke_Arity1(c).(float64)
									var b = Chunk_buffer.X_invoke_Arity1(size).(*CljsCoreChunkBuffer)
									_, _, _ = c, size, b
									{
										var n__77528__auto___616 = size
										_ = n__77528__auto___616
										{
											var i_617 = float64(0)
											_ = i_617
											for {
												if i_617 < n__77528__auto___616 {
													Chunk_append.X_invoke_Arity2(b, f.(CljsCoreIFn).X_invoke_Arity1(c.X_nth_Arity2(i_617).(interface{})).(interface{})).(interface{})
													{
														var G__618 = (i_617 + float64(1))
														i_617 = G__618
														continue
													}
												} else {
												}
												break
											}
										}
									}
									return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b).(interface{}), map_.X_invoke_Arity2(f, Chunk_rest.X_invoke_Arity1(s).(interface{})).(*CljsCoreLazySeq))
								}
							} else {
								return Cons.X_invoke_Arity2(f.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(s)).(interface{}), map_.X_invoke_Arity2(f, Rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
							}
						}
					} else {
						return nil
					}
				}
			})
		}(&AFn{}), nil, nil})
	}, func(f interface{}, c1 interface{}, c2 interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__619 *AFn) *AFn {
			return Fn(G__619, func() interface{} {
				{
					var s1 = Seq.Arity1IQ(c1)
					var s2 = Seq.Arity1IQ(c2)
					_, _ = s1, s2
					if (s1) && (s2) {
						return Cons.X_invoke_Arity2(f.(CljsCoreIFn).X_invoke_Arity2(First.X_invoke_Arity1(s1), First.X_invoke_Arity1(s2)).(interface{}), map_.X_invoke_Arity3(f, Rest.X_invoke_Arity1(s1), Rest.X_invoke_Arity1(s2)).(*CljsCoreLazySeq)).(*CljsCoreCons)
					} else {
						return nil
					}
				}
			})
		}(&AFn{}), nil, nil})
	}, func(f interface{}, c1 interface{}, c2 interface{}, c3 interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__620 *AFn) *AFn {
			return Fn(G__620, func() interface{} {
				{
					var s1 = Seq.Arity1IQ(c1)
					var s2 = Seq.Arity1IQ(c2)
					var s3 = Seq.Arity1IQ(c3)
					_, _, _ = s1, s2, s3
					if (s1) && (s2) && (s3) {
						return Cons.X_invoke_Arity2(f.(CljsCoreIFn).X_invoke_Arity3(First.X_invoke_Arity1(s1), First.X_invoke_Arity1(s2), First.X_invoke_Arity1(s3)).(interface{}), map_.X_invoke_Arity4(f, Rest.X_invoke_Arity1(s1), Rest.X_invoke_Arity1(s2), Rest.X_invoke_Arity1(s3)).(*CljsCoreLazySeq)).(*CljsCoreCons)
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
					return (&CljsCoreLazySeq{nil, func(G__621 *AFn) *AFn {
						return Fn(G__621, func() interface{} {
							{
								var ss = map_.X_invoke_Arity2(Seq, cs).(*CljsCoreLazySeq)
								_ = ss
								if Every_QMARK_.X_invoke_Arity2(Identity, ss) {
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
			return map_.X_invoke_Arity2(func(step interface{}) *AFn {
				return func(G__622 *AFn) *AFn {
					return Fn(G__622, func(p1__612_SHARP_ interface{}) interface{} {
						return Apply.X_invoke_Arity2(f, p1__612_SHARP_)
					})
				}(&AFn{})
			}(step), step.(CljsCoreIFn).X_invoke_Arity1(Conj.X_invoke_ArityVariadic(colls, c3, c2, c1).(interface{})).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
		}
	})
}(&AFn{})

/**
* Returns a lazy sequence of the first n items in coll, or all items if
* there are fewer than n.  Returns a stateful transducer when
* no collection is provided.
 */
var Take = func(take *AFn) *AFn {
	return Fn(take, func(n interface{}) interface{} {
		return func(G__623 *AFn) *AFn {
			return Fn(G__623, func(f1 interface{}) interface{} {
				{
					var na = Atom.X_invoke_Arity1(n).(*CljsCoreAtom)
					_ = na
					return func(na *CljsCoreAtom) *AFn {
						return func(G__624 *AFn) *AFn {
							return Fn(G__624, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
							}, func(result interface{}, input interface{}) interface{} {
								{
									var n___1 = Deref.X_invoke_Arity1(na).(interface{})
									var nn = Swap_BANG_.X_invoke_Arity2(na, Dec)
									var result___1 = func() interface{} {
										if n___1.(float64) > float64(0) {
											return f1.(CljsCoreIFn).X_invoke_Arity2(result, input).(interface{})
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
						}(&AFn{})
					}(na)
				}
			})
		}(&AFn{})
	}, func(n interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__625 *AFn) *AFn {
			return Fn(G__625, func() interface{} {
				if n.(float64) > float64(0) {
					{
						var temp__4219__auto__ = Seq.Arity1IQ(coll)
						_ = temp__4219__auto__
						if Truth_(temp__4219__auto__) {
							{
								var s = temp__4219__auto__
								_ = s
								return Cons.X_invoke_Arity2(First.X_invoke_Arity1(s), take.X_invoke_Arity2((n.(float64)-float64(1)), Rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
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

/**
* Returns a lazy sequence of all but the first n items in coll.
* Returns a stateful transducer when no collection is provided.
 */
var Drop = func(drop *AFn) *AFn {
	return Fn(drop, func(n interface{}) interface{} {
		return func(G__626 *AFn) *AFn {
			return Fn(G__626, func(f1 interface{}) interface{} {
				{
					var na = Atom.X_invoke_Arity1(n).(*CljsCoreAtom)
					_ = na
					return func(na *CljsCoreAtom) *AFn {
						return func(G__627 *AFn) *AFn {
							return Fn(G__627, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
							}, func(result interface{}, input interface{}) interface{} {
								{
									var n___1 = Deref.X_invoke_Arity1(na).(interface{})
									_ = n___1
									Swap_BANG_.X_invoke_Arity2(na, Dec)
									if n___1.(float64) > float64(0) {
										return result
									} else {
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, input).(interface{})
									}
								}
							})
						}(&AFn{})
					}(na)
				}
			})
		}(&AFn{})
	}, func(n interface{}, coll interface{}) interface{} {
		{
			var step = func(G__628 *AFn) *AFn {
				return Fn(G__628, func(n___1 interface{}, coll___1 interface{}) interface{} {
					for {
						{
							var s = Seq.Arity1IQ(coll___1)
							_ = s
							if (n___1.(float64) > float64(0)) && (s) {
								{
									var G__629 = (n___1.(float64) - float64(1))
									var G__630 = Rest.X_invoke_Arity1(s)
									n___1 = G__629
									coll___1 = G__630
									continue
								}
							} else {
								return s
							}
						}
					}
				})
			}(&AFn{})
			_ = step
			return (&CljsCoreLazySeq{nil, func(step interface{}) *AFn {
				return func(G__631 *AFn) *AFn {
					return Fn(G__631, func() interface{} {
						return step.(CljsCoreIFn).X_invoke_Arity2(n, coll).(CljsCoreISeq)
					})
				}(&AFn{})
			}(step), nil, nil})
		}
	})
}(&AFn{})

/**
* Return a lazy sequence of all but the last n (default 1) items in coll
 */
var Drop_last = func(drop_last *AFn) *AFn {
	return Fn(drop_last, func(s interface{}) interface{} {
		return drop_last.X_invoke_Arity2(float64(1), s).(*CljsCoreLazySeq)
	}, func(n interface{}, s interface{}) interface{} {
		return Map_.X_invoke_Arity3(func(G__632 *AFn) *AFn {
			return Fn(G__632, func(x interface{}, ___ interface{}) interface{} {
				return x
			})
		}(&AFn{}), s, Drop.X_invoke_Arity2(n, s).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
	})
}(&AFn{})

/**
* Returns a seq of the last n items in coll.  Depending on the type
* of coll may be no better than linear time.  For vectors, see also subvec.
 */
var Take_last = func(take_last *AFn) *AFn {
	return Fn(take_last, func(n interface{}, coll interface{}) interface{} {
		{
			var s = Seq.Arity1IQ(coll)
			var lead = Seq.X_invoke_Arity1(Drop.X_invoke_Arity2(n, coll).(*CljsCoreLazySeq))
			_, _ = s, lead
			for {
				if Truth_(lead) {
					{
						var G__633 = Next.X_invoke_Arity1(s)
						var G__634 = Next.X_invoke_Arity1(lead)
						s = G__633
						lead = G__634
						continue
					}
				} else {
					return s
				}
			}
		}
	})
}(&AFn{})

/**
* Returns a lazy sequence of the items in coll starting from the
* first item for which (pred item) returns logical false.  Returns a
* stateful transducer when no collection is provided.
 */
var Drop_while = func(drop_while *AFn) *AFn {
	return Fn(drop_while, func(pred interface{}) interface{} {
		return func(G__635 *AFn) *AFn {
			return Fn(G__635, func(f1 interface{}) interface{} {
				{
					var da = Atom.X_invoke_Arity1(true).(*CljsCoreAtom)
					_ = da
					return func(da *CljsCoreAtom) *AFn {
						return func(G__636 *AFn) *AFn {
							return Fn(G__636, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
							}, func(result interface{}, input interface{}) interface{} {
								{
									var drop_QMARK_ = Deref.X_invoke_Arity1(da).(interface{})
									_ = drop_QMARK_
									if Truth_(func() interface{} {
										var and__76620__auto__ = drop_QMARK_
										_ = and__76620__auto__
										if Truth_(and__76620__auto__) {
											return pred.(CljsCoreIFn).X_invoke_Arity1(input).(interface{})
										} else {
											return and__76620__auto__
										}
									}()) {
										return result
									} else {
										Reset_BANG_.X_invoke_Arity2(da, nil)
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, input).(interface{})
									}
								}
							})
						}(&AFn{})
					}(da)
				}
			})
		}(&AFn{})
	}, func(pred interface{}, coll interface{}) interface{} {
		{
			var step = func(G__637 *AFn) *AFn {
				return Fn(G__637, func(pred___1 interface{}, coll___1 interface{}) interface{} {
					for {
						{
							var s = Seq.Arity1IQ(coll___1)
							_ = s
							if Truth_(func() interface{} {
								var and__76620__auto__ = s
								_ = and__76620__auto__
								if Truth_(and__76620__auto__) {
									return pred___1.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(s)).(interface{})
								} else {
									return and__76620__auto__
								}
							}()) {
								{
									var G__638 = pred___1
									var G__639 = Rest.X_invoke_Arity1(s)
									pred___1 = G__638
									coll___1 = G__639
									continue
								}
							} else {
								return s
							}
						}
					}
				})
			}(&AFn{})
			_ = step
			return (&CljsCoreLazySeq{nil, func(step interface{}) *AFn {
				return func(G__640 *AFn) *AFn {
					return Fn(G__640, func() interface{} {
						return step.(CljsCoreIFn).X_invoke_Arity2(pred, coll).(CljsCoreISeq)
					})
				}(&AFn{})
			}(step), nil, nil})
		}
	})
}(&AFn{})

/**
* Returns a lazy (infinite!) sequence of repetitions of the items in coll.
 */
var Cycle = func(cycle *AFn) *AFn {
	return Fn(cycle, func(coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__641 *AFn) *AFn {
			return Fn(G__641, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
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

/**
* Returns a vector of [(take n coll) (drop n coll)]
 */
var Split_at = func(split_at *AFn) *AFn {
	return Fn(split_at, func(n interface{}, coll interface{}) interface{} {
		return &CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{Take.X_invoke_Arity2(n, coll).(*CljsCoreLazySeq), Drop.X_invoke_Arity2(n, coll).(*CljsCoreLazySeq)}, nil}
	})
}(&AFn{})

/**
* Returns a lazy (infinite!, or length n if supplied) sequence of xs.
 */
var Repeat = func(repeat *AFn) *AFn {
	return Fn(repeat, func(x interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__642 *AFn) *AFn {
			return Fn(G__642, func() interface{} {
				return Cons.X_invoke_Arity2(x, repeat.X_invoke_Arity1(x).(*CljsCoreLazySeq)).(*CljsCoreCons)
			})
		}(&AFn{}), nil, nil})
	}, func(n interface{}, x interface{}) interface{} {
		return Take.X_invoke_Arity2(n, repeat.X_invoke_Arity1(x).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
	})
}(&AFn{})

/**
* Returns a lazy seq of n xs.
 */
var Replicate = func(replicate *AFn) *AFn {
	return Fn(replicate, func(n interface{}, x interface{}) interface{} {
		return Take.X_invoke_Arity2(n, Repeat.X_invoke_Arity1(x).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
	})
}(&AFn{})

/**
* Takes a function of no args, presumably with side effects, and
* returns an infinite (or length n if supplied) lazy sequence of calls
* to it
 */
var Repeatedly = func(repeatedly *AFn) *AFn {
	return Fn(repeatedly, func(f interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__643 *AFn) *AFn {
			return Fn(G__643, func() interface{} {
				return Cons.X_invoke_Arity2(f.(CljsCoreIFn).X_invoke_Arity0().(interface{}), repeatedly.X_invoke_Arity1(f).(*CljsCoreLazySeq)).(*CljsCoreCons)
			})
		}(&AFn{}), nil, nil})
	}, func(n interface{}, f interface{}) interface{} {
		return Take.X_invoke_Arity2(n, repeatedly.X_invoke_Arity1(f).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
	})
}(&AFn{})

/**
* Returns a lazy sequence of x, (f x), (f (f x)) etc. f must be free of side-effects
 */
var Iterate = func(iterate *AFn) *AFn {
	return Fn(iterate, func(f interface{}, x interface{}) interface{} {
		return Cons.X_invoke_Arity2(x, (&CljsCoreLazySeq{nil, func(G__644 *AFn) *AFn {
			return Fn(G__644, func() interface{} {
				return iterate.X_invoke_Arity2(f, f.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})).(*CljsCoreCons)
			})
		}(&AFn{}), nil, nil})).(*CljsCoreCons)
	})
}(&AFn{})

/**
* Returns a lazy seq of the first item in each coll, then the second etc.
* @param {...*} var_args
 */
var Interleave = func(interleave *AFn) *AFn {
	return Fn(interleave, func(c1 interface{}, c2 interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__645 *AFn) *AFn {
			return Fn(G__645, func() interface{} {
				{
					var s1 = Seq.Arity1IQ(c1)
					var s2 = Seq.Arity1IQ(c2)
					_, _ = s1, s2
					if (s1) && (s2) {
						return Cons.X_invoke_Arity2(First.X_invoke_Arity1(s1), Cons.X_invoke_Arity2(First.X_invoke_Arity1(s2), interleave.X_invoke_Arity2(Rest.X_invoke_Arity1(s1), Rest.X_invoke_Arity1(s2)).(*CljsCoreLazySeq)).(*CljsCoreCons)).(*CljsCoreCons)
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
		return (&CljsCoreLazySeq{nil, func(G__646 *AFn) *AFn {
			return Fn(G__646, func() interface{} {
				{
					var ss = Map_.X_invoke_Arity2(Seq, Conj.X_invoke_ArityVariadic(colls, c2, c1).(interface{})).(*CljsCoreLazySeq)
					_ = ss
					if Every_QMARK_.X_invoke_Arity2(Identity, ss) {
						return Concat.X_invoke_Arity2(Map_.X_invoke_Arity2(First, ss).(*CljsCoreLazySeq), Apply.X_invoke_Arity2(interleave, Map_.X_invoke_Arity2(Rest, ss).(*CljsCoreLazySeq))).(*CljsCoreLazySeq)
					} else {
						return nil
					}
				}
			})
		}(&AFn{}), nil, nil})
	})
}(&AFn{})

/**
* Returns a lazy seq of the elements of coll separated by sep
 */
var Interpose = func(interpose *AFn) *AFn {
	return Fn(interpose, func(sep interface{}, coll interface{}) interface{} {
		return Drop.X_invoke_Arity2(float64(1), Interleave.X_invoke_Arity2(Repeat.X_invoke_Arity1(sep).(*CljsCoreLazySeq), coll).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
	})
}(&AFn{})

/**
* Take a collection of collections, and return a lazy seq
* of items from the inner collection
 */
var Flatten1 = func(flatten1 *AFn) *AFn {
	return Fn(flatten1, func(colls interface{}) interface{} {
		{
			var cat = func(cat *AFn) *AFn {
				return Fn(cat, func(coll interface{}, colls___1 interface{}) interface{} {
					return (&CljsCoreLazySeq{nil, func(G__647 *AFn) *AFn {
						return Fn(G__647, func() interface{} {
							{
								var temp__4217__auto__ = Seq.Arity1IQ(coll)
								_ = temp__4217__auto__
								if Truth_(temp__4217__auto__) {
									{
										var coll___1 = temp__4217__auto__
										_ = coll___1
										return Cons.X_invoke_Arity2(First.X_invoke_Arity1(coll___1), cat.X_invoke_Arity2(Rest.X_invoke_Arity1(coll___1), colls___1).(*CljsCoreLazySeq)).(*CljsCoreCons)
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
			return cat.(CljsCoreIFn).X_invoke_Arity2(nil, colls).(*CljsCoreLazySeq)
		}
	})
}(&AFn{})

/**
* Returns the result of applying concat to the result of applying map
* to f and colls.  Thus function f should return a collection.
* @param {...*} var_args
 */
var Mapcat = func(mapcat *AFn) *AFn {
	return Fn(mapcat, func(f interface{}, coll interface{}) interface{} {
		return Flatten1.X_invoke_Arity1(Map_.X_invoke_Arity2(f, coll).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
	}, func(f_coll_colls__ ...interface{}) interface{} {
		var f = f_coll_colls__[0]
		var coll = f_coll_colls__[1]
		var colls = Array_seq.X_invoke_Arity1(f_coll_colls__[2:])
		_, _, _ = f, coll, colls
		return Flatten1.X_invoke_Arity1(Apply.X_invoke_Arity4(Map_, f, coll, colls)).(*CljsCoreLazySeq)
	})
}(&AFn{})

/**
* Returns a lazy sequence of the items in coll for which
* (pred item) returns true. pred must be free of side-effects.
* Returns a transducer when no collection is provided.
 */
var Filter = func(filter *AFn) *AFn {
	return Fn(filter, func(pred interface{}) interface{} {
		return func(G__648 *AFn) *AFn {
			return Fn(G__648, func(f1 interface{}) interface{} {
				return func(G__649 *AFn) *AFn {
					return Fn(G__649, func() interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
					}, func(result interface{}) interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
					}, func(result interface{}, input interface{}) interface{} {
						if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(input).(interface{})) {
							return f1.(CljsCoreIFn).X_invoke_Arity2(result, input).(interface{})
						} else {
							return result
						}
					})
				}(&AFn{})
			})
		}(&AFn{})
	}, func(pred interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__650 *AFn) *AFn {
			return Fn(G__650, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
							_ = s
							if Chunked_seq_QMARK_.X_invoke_Arity1(s) {
								{
									var c = Chunk_first.X_invoke_Arity1(s).(interface{})
									var size = Count.X_invoke_Arity1(c).(float64)
									var b = Chunk_buffer.X_invoke_Arity1(size).(*CljsCoreChunkBuffer)
									_, _, _ = c, size, b
									{
										var n__77528__auto___651 = size
										_ = n__77528__auto___651
										{
											var i_652 = float64(0)
											_ = i_652
											for {
												if i_652 < n__77528__auto___651 {
													if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(c.X_nth_Arity2(i_652).(interface{})).(interface{})) {
														Chunk_append.X_invoke_Arity2(b, c.X_nth_Arity2(i_652).(interface{})).(interface{})
													} else {
													}
													{
														var G__653 = (i_652 + float64(1))
														i_652 = G__653
														continue
													}
												} else {
												}
												break
											}
										}
									}
									return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b).(interface{}), filter.X_invoke_Arity2(pred, Chunk_rest.X_invoke_Arity1(s).(interface{})).(*CljsCoreLazySeq))
								}
							} else {
								{
									var f = First.X_invoke_Arity1(s)
									var r = Rest.X_invoke_Arity1(s)
									_, _ = f, r
									if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(f).(interface{})) {
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

/**
* Returns a lazy sequence of the items in coll for which
* (pred item) returns false. pred must be free of side-effects.
* Returns a transducer when no collection is provided.
 */
var Remove = func(remove *AFn) *AFn {
	return Fn(remove, func(pred interface{}) interface{} {
		return Filter.X_invoke_Arity1(Complement.Arity1IB(pred)).(interface{})
	}, func(pred interface{}, coll interface{}) interface{} {
		return Filter.X_invoke_Arity2(Complement.Arity1IB(pred), coll).(*CljsCoreLazySeq)
	})
}(&AFn{})

/**
* Returns a lazy sequence of the nodes in a tree, via a depth-first walk.
* branch? must be a fn of one arg that returns true if passed a node
* that can have children (but may not).  children must be a fn of one
* arg that returns a sequence of the children. Will only be called on
* nodes for which branch? returns true. Root is the root node of the
* tree.
 */
var Tree_seq = func(tree_seq *AFn) *AFn {
	return Fn(tree_seq, func(branch_QMARK_ interface{}, children interface{}, root interface{}) interface{} {
		{
			var walk = func(walk *AFn) *AFn {
				return Fn(walk, func(node interface{}) interface{} {
					return (&CljsCoreLazySeq{nil, func(G__654 *AFn) *AFn {
						return Fn(G__654, func() interface{} {
							return Cons.X_invoke_Arity2(node, func() interface{} {
								if Truth_(branch_QMARK_.(CljsCoreIFn).X_invoke_Arity1(node).(interface{}).(bool)) {
									return Mapcat.X_invoke_Arity2(walk, children.(CljsCoreIFn).X_invoke_Arity1(node).(interface{})).(*CljsCoreLazySeq)
								} else {
									return nil
								}
							}()).(*CljsCoreCons)
						})
					}(&AFn{}), nil, nil})
				})
			}(&AFn{})
			_ = walk
			return walk.(CljsCoreIFn).X_invoke_Arity1(root).(*CljsCoreLazySeq)
		}
	})
}(&AFn{})

/**
* Takes any nested combination of sequential things (lists, vectors,
* etc.) and returns their contents as a single, flat sequence.
* (flatten nil) returns nil.
 */
var Flatten = func(flatten *AFn) *AFn {
	return Fn(flatten, func(x interface{}) interface{} {
		return Filter.X_invoke_Arity2(func(G__656 *AFn) *AFn {
			return Fn(G__656, func(p1__655_SHARP_ interface{}) interface{} {
				return !(Sequential_QMARK_.Arity1IB(p1__655_SHARP_))
			})
		}(&AFn{}), Rest.X_invoke_Arity1(Tree_seq.X_invoke_Arity3(Sequential_QMARK_, Seq, x).(*CljsCoreLazySeq))).(*CljsCoreLazySeq)
	})
}(&AFn{})

/**
* Returns a new coll consisting of to-coll with all of the items of
* from-coll conjoined. A transducer may be supplied.
 */
var Into = func(into *AFn) *AFn {
	return Fn(into, func(to interface{}, from interface{}) interface{} {
		if !(to == nil) {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IEditableCollection", Str: "cljs.core/IEditableCollection", X_hash: float64(297050504), X_meta: nil}), to)) {
				return With_meta.X_invoke_Arity2(Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(X_conj_BANG_, Transient.X_invoke_Arity1(to).(interface{}), from)).(interface{}), Meta.X_invoke_Arity1(to))
			} else {
				return Reduce.X_invoke_Arity3(X_conj, to, from)
			}
		} else {
			return Reduce.X_invoke_Arity3(Conj, CljsCoreList_EMPTY, from)
		}
	}, func(to interface{}, xform interface{}, from interface{}) interface{} {
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IEditableCollection", Str: "cljs.core/IEditableCollection", X_hash: float64(297050504), X_meta: nil}), to)) {
			return With_meta.X_invoke_Arity2(Persistent_BANG_.X_invoke_Arity1(Transduce.X_invoke_Arity4(xform, X_conj_BANG_, Transient.X_invoke_Arity1(to).(interface{}), from).(interface{})).(interface{}), Meta.X_invoke_Arity1(to))
		} else {
			return Transduce.X_invoke_Arity4(xform, Conj, to, from).(interface{})
		}
	})
}(&AFn{})

/**
* Returns a vector consisting of the result of applying f to the
* set of first items of each coll, followed by applying f to the set
* of second items in each coll, until any one of the colls is
* exhausted.  Any remaining items in other colls are ignored. Function
* f should accept number-of-colls arguments.
* @param {...*} var_args
 */
var Mapv = func(mapv *AFn) *AFn {
	return Fn(mapv, func(f interface{}, coll interface{}) interface{} {
		return Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(func(G__657 *AFn) *AFn {
			return Fn(G__657, func(v interface{}, o interface{}) interface{} {
				return Conj_BANG_.X_invoke_Arity2(v, f.(CljsCoreIFn).X_invoke_Arity1(o).(interface{})).(interface{})
			})
		}(&AFn{}), Transient.X_invoke_Arity1(CljsCorePersistentVector_EMPTY).(interface{}), coll)).(interface{})
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

/**
* Returns a vector of the items in coll for which
* (pred item) returns true. pred must be free of side-effects.
 */
var Filterv = func(filterv *AFn) *AFn {
	return Fn(filterv, func(pred interface{}, coll interface{}) interface{} {
		return Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(func(G__658 *AFn) *AFn {
			return Fn(G__658, func(v interface{}, o interface{}) interface{} {
				if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(o).(interface{})) {
					return Conj_BANG_.X_invoke_Arity2(v, o).(interface{})
				} else {
					return v
				}
			})
		}(&AFn{}), Transient.X_invoke_Arity1(CljsCorePersistentVector_EMPTY).(interface{}), coll)).(interface{})
	})
}(&AFn{})

/**
* Returns a lazy sequence of lists of n items each, at offsets step
* apart. If step is not supplied, defaults to n, i.e. the partitions
* do not overlap. If a pad collection is supplied, use its elements as
* necessary to complete last partition upto n items. In case there are
* not enough padding elements, return a partition with less than n items.
 */
var Partition = func(partition *AFn) *AFn {
	return Fn(partition, func(n interface{}, coll interface{}) interface{} {
		return partition.X_invoke_Arity3(n, n, coll).(*CljsCoreLazySeq)
	}, func(n interface{}, step interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__659 *AFn) *AFn {
			return Fn(G__659, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
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
		return (&CljsCoreLazySeq{nil, func(G__660 *AFn) *AFn {
			return Fn(G__660, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
							_ = s
							{
								var p = Take.X_invoke_Arity2(n, s).(*CljsCoreLazySeq)
								_ = p
								if n.(float64) == Count.X_invoke_Arity1(p).(float64) {
									return Cons.X_invoke_Arity2(p, partition.X_invoke_Arity4(n, step, pad, Drop.X_invoke_Arity2(step, s).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)).(*CljsCoreCons)
								} else {
									return CljsCoreList_EMPTY.(CljsCoreICollection).X_conj_Arity2(Take.X_invoke_Arity2(n, Concat.X_invoke_Arity2(p, pad).(*CljsCoreLazySeq)).(*CljsCoreLazySeq))
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

/**
* Returns the value in a nested associative structure,
* where ks is a sequence of keys. Returns nil if the key is not present,
* or the not-found value if supplied.
 */
var Get_in = func(get_in *AFn) *AFn {
	return Fn(get_in, func(m interface{}, ks interface{}) interface{} {
		return get_in.X_invoke_Arity3(m, ks, nil).(interface{})
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
							if sentinel == m___2 {
								return not_found
							} else {
								{
									var G__661 = sentinel
									var G__662 = m___2
									var G__663 = Next.X_invoke_Arity1(ks___1)
									sentinel = G__661
									m___1 = G__662
									ks___1 = G__663
									continue
								}
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

/**
* Associates a value in a nested associative structure, where ks is a
* sequence of keys and v is the new value and returns a new nested structure.
* If any levels do not exist, hash-maps will be created.
 */
var Assoc_in = func(assoc_in *AFn) *AFn {
	return Fn(assoc_in, func(m interface{}, p__664 interface{}, v interface{}) interface{} {
		{
			var vec__666 = p__664
			var k = Nth.X_invoke_Arity3(vec__666, float64(0), nil)
			var ks = Nthnext.X_invoke_Arity2(vec__666, float64(1)).(CljsCoreISeq)
			_, _, _ = vec__666, k, ks
			if Truth_(ks) {
				return Assoc.X_invoke_Arity3(m, k, assoc_in.X_invoke_Arity3(Get.X_invoke_Arity2(m, k), ks, v).(interface{})).(interface{})
			} else {
				return Assoc.X_invoke_Arity3(m, k, v).(interface{})
			}
		}
	})
}(&AFn{})

/**
* 'Updates' a value in a nested associative structure, where ks is a
* sequence of keys and f is a function that will take the old value
* and any supplied args and return the new value, and returns a new
* nested structure.  If any levels do not exist, hash-maps will be
* created.
* @param {...*} var_args
 */
var Update_in = func(update_in *AFn) *AFn {
	return Fn(update_in, func(m interface{}, p__667 interface{}, f interface{}) interface{} {
		{
			var vec__677 = p__667
			var k = Nth.X_invoke_Arity3(vec__677, float64(0), nil)
			var ks = Nthnext.X_invoke_Arity2(vec__677, float64(1)).(CljsCoreISeq)
			_, _, _ = vec__677, k, ks
			if Truth_(ks) {
				return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity3(Get.X_invoke_Arity2(m, k), ks, f).(interface{})).(interface{})
			} else {
				return Assoc.X_invoke_Arity3(m, k, f.(CljsCoreIFn).X_invoke_Arity1(Get.X_invoke_Arity2(m, k)).(interface{})).(interface{})
			}
		}
	}, func(m interface{}, p__668 interface{}, f interface{}, a interface{}) interface{} {
		{
			var vec__678 = p__668
			var k = Nth.X_invoke_Arity3(vec__678, float64(0), nil)
			var ks = Nthnext.X_invoke_Arity2(vec__678, float64(1)).(CljsCoreISeq)
			_, _, _ = vec__678, k, ks
			if Truth_(ks) {
				return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity4(Get.X_invoke_Arity2(m, k), ks, f, a).(interface{})).(interface{})
			} else {
				return Assoc.X_invoke_Arity3(m, k, f.(CljsCoreIFn).X_invoke_Arity2(Get.X_invoke_Arity2(m, k), a).(interface{})).(interface{})
			}
		}
	}, func(m interface{}, p__669 interface{}, f interface{}, a interface{}, b interface{}) interface{} {
		{
			var vec__679 = p__669
			var k = Nth.X_invoke_Arity3(vec__679, float64(0), nil)
			var ks = Nthnext.X_invoke_Arity2(vec__679, float64(1)).(CljsCoreISeq)
			_, _, _ = vec__679, k, ks
			if Truth_(ks) {
				return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity5(Get.X_invoke_Arity2(m, k), ks, f, a, b).(interface{})).(interface{})
			} else {
				return Assoc.X_invoke_Arity3(m, k, f.(CljsCoreIFn).X_invoke_Arity3(Get.X_invoke_Arity2(m, k), a, b).(interface{})).(interface{})
			}
		}
	}, func(m interface{}, p__670 interface{}, f interface{}, a interface{}, b interface{}, c interface{}) interface{} {
		{
			var vec__680 = p__670
			var k = Nth.X_invoke_Arity3(vec__680, float64(0), nil)
			var ks = Nthnext.X_invoke_Arity2(vec__680, float64(1)).(CljsCoreISeq)
			_, _, _ = vec__680, k, ks
			if Truth_(ks) {
				return Assoc.X_invoke_Arity3(m, k, update_in.X_invoke_Arity6(Get.X_invoke_Arity2(m, k), ks, f, a, b, c).(interface{})).(interface{})
			} else {
				return Assoc.X_invoke_Arity3(m, k, f.(CljsCoreIFn).X_invoke_Arity4(Get.X_invoke_Arity2(m, k), a, b, c).(interface{})).(interface{})
			}
		}
	}, func(m_p__671_f_a_b_c_args__ ...interface{}) interface{} {
		var m = m_p__671_f_a_b_c_args__[0]
		var p__671 = m_p__671_f_a_b_c_args__[1]
		var f = m_p__671_f_a_b_c_args__[2]
		var a = m_p__671_f_a_b_c_args__[3]
		var b = m_p__671_f_a_b_c_args__[4]
		var c = m_p__671_f_a_b_c_args__[5]
		var args = Array_seq.X_invoke_Arity1(m_p__671_f_a_b_c_args__[6:])
		_, _, _, _, _, _, _ = m, p__671, f, a, b, c, args
		{
			var vec__681 = p__671
			var k = Nth.X_invoke_Arity3(vec__681, float64(0), nil)
			var ks = Nthnext.X_invoke_Arity2(vec__681, float64(1)).(CljsCoreISeq)
			_, _, _ = vec__681, k, ks
			if Truth_(ks) {
				return Assoc.X_invoke_Arity3(m, k, Apply.X_invoke_ArityVariadic(update_in, Get.X_invoke_Arity2(m, k), ks, f, a, b, c, args)).(interface{})
			} else {
				return Assoc.X_invoke_Arity3(m, k, Apply.X_invoke_ArityVariadic(f, Get.X_invoke_Arity2(m, k), a, b, c, args)).(interface{})
			}
		}
	})
}(&AFn{})

type CljsCoreVectorNode struct {
	Edit interface{}
	Arr  interface{}
}

var X__GT_VectorNode = func(__GT_VectorNode *AFn) *AFn {
	return Fn(__GT_VectorNode, func(edit interface{}, arr interface{}) interface{} {
		return (&CljsCoreVectorNode{edit, arr})
	})
}(&AFn{})

var Pv_fresh_node = func(pv_fresh_node *AFn) *AFn {
	return Fn(pv_fresh_node, func(edit interface{}) interface{} {
		return (&CljsCoreVectorNode{edit, make([]interface{}, int(float64(32)))})
	})
}(&AFn{})

var Pv_aget = func(pv_aget *AFn) *AFn {
	return Fn(pv_aget, func(node interface{}, idx interface{}) interface{} {
		return (Native_get_instance_field.X_invoke_Arity2(node, "Arr").([]interface{})[int(idx.(float64))])
	})
}(&AFn{})

var Pv_aset = func(pv_aset *AFn) *AFn {
	return Fn(pv_aset, func(node interface{}, idx interface{}, val interface{}) interface{} {
		return func() interface{} {
			Native_get_instance_field.X_invoke_Arity2(node, "Arr").([]interface{})[int(idx.(float64))] = val
			return Native_get_instance_field.X_invoke_Arity2(node, "Arr").([]interface{})[int(idx.(float64))]
		}()
	})
}(&AFn{})

var Pv_clone_node = func(pv_clone_node *AFn) *AFn {
	return Fn(pv_clone_node, func(node interface{}) interface{} {
		return (&CljsCoreVectorNode{Native_get_instance_field.X_invoke_Arity2(node, "Edit"), Aclone.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(node, "Arr")).([]interface{})})
	})
}(&AFn{})

var Tail_off = func(tail_off *AFn) *AFn {
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

var New_path = func(new_path *AFn) *AFn {
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
						var ___ = Pv_aset.X_invoke_Arity3(r, float64(0), embed).(interface{})
						_, _, _ = embed, r, ___
						{
							var G__682 = (ll.(float64) - float64(5))
							var G__683 = r
							ll = G__682
							ret = G__683
							continue
						}
					}
				}
			}
		}
	})
}(&AFn{})

var Push_tail = func(push_tail *AFn) *AFn {
	return Fn(push_tail, func(pv interface{}, level interface{}, parent interface{}, tailnode interface{}) interface{} {
		{
			var ret = Pv_clone_node.X_invoke_Arity1(parent).(*CljsCoreVectorNode)
			var subidx = float64(int(float64(uint((Native_get_instance_field.X_invoke_Arity2(pv, "Cnt").(float64)-float64(1)))>>uint(level.(float64)))) & int(float64(31)))
			_, _ = ret, subidx
			if float64(5) == level.(float64) {
				Pv_aset.X_invoke_Arity3(ret, subidx, tailnode).(interface{})
				return ret
			} else {
				{
					var child = Pv_aget.X_invoke_Arity2(parent, subidx).(interface{})
					_ = child
					if !(child == nil) {
						{
							var node_to_insert = push_tail.X_invoke_Arity4(pv, (level.(float64) - float64(5)), child, tailnode).(*CljsCoreVectorNode)
							_ = node_to_insert
							Pv_aset.X_invoke_Arity3(ret, subidx, node_to_insert).(interface{})
							return ret
						}
					} else {
						{
							var node_to_insert = New_path.X_invoke_Arity3(nil, (level.(float64) - float64(5)), tailnode).(interface{})
							_ = node_to_insert
							Pv_aset.X_invoke_Arity3(ret, subidx, node_to_insert).(interface{})
							return ret
						}
					}
				}
			}
		}
	})
}(&AFn{})

var Vector_index_out_of_bounds = func(vector_index_out_of_bounds *AFn) *AFn {
	return Fn(vector_index_out_of_bounds, func(i interface{}, cnt interface{}) interface{} {
		panic((&js.Error{("No item " + Str.X_invoke_Arity1(i).(string) + " in vector of length " + Str.X_invoke_Arity1(cnt).(string))}))
	})
}(&AFn{})

var First_array_for_longvec = func(first_array_for_longvec *AFn) *AFn {
	return Fn(first_array_for_longvec, func(pv interface{}) interface{} {
		{
			var node = Native_get_instance_field.X_invoke_Arity2(pv, "Root")
			var level = Native_get_instance_field.X_invoke_Arity2(pv, "Shift")
			_, _ = node, level
			for {
				if level.(float64) > float64(0) {
					{
						var G__684 = Pv_aget.X_invoke_Arity2(node, float64(0)).(interface{})
						var G__685 = (level.(float64) - float64(5))
						node = G__684
						level = G__685
						continue
					}
				} else {
					return Native_get_instance_field.X_invoke_Arity2(node, "Arr")
				}
			}
		}
	})
}(&AFn{})

var Unchecked_array_for = func(unchecked_array_for *AFn) *AFn {
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
						{
							var G__686 = Pv_aget.X_invoke_Arity2(node, float64(int(float64(uint(i.(float64))>>uint(level.(float64))))&int(float64(31)))).(interface{})
							var G__687 = (level.(float64) - float64(5))
							node = G__686
							level = G__687
							continue
						}
					} else {
						return Native_get_instance_field.X_invoke_Arity2(node, "Arr")
					}
				}
			}
		}
	})
}(&AFn{})

var Array_for = func(array_for *AFn) *AFn {
	return Fn(array_for, func(pv interface{}, i interface{}) interface{} {
		if (float64(0) <= i.(float64)) && (i.(float64) < Native_get_instance_field.X_invoke_Arity2(pv, "Cnt").(float64)) {
			return Unchecked_array_for.X_invoke_Arity2(pv, i).(interface{})
		} else {
			return Vector_index_out_of_bounds.X_invoke_Arity2(i, Native_get_instance_field.X_invoke_Arity2(pv, "Cnt")).(interface{})
		}
	})
}(&AFn{})

var Do_assoc = func(do_assoc *AFn) *AFn {
	return Fn(do_assoc, func(pv interface{}, level interface{}, node interface{}, i interface{}, val interface{}) interface{} {
		{
			var ret = Pv_clone_node.X_invoke_Arity1(node).(*CljsCoreVectorNode)
			_ = ret
			if level.(float64) == float64(0) {
				Pv_aset.X_invoke_Arity3(ret, float64(int(i.(float64))&int(float64(31))), val).(interface{})
				return ret
			} else {
				{
					var subidx = float64(int(float64(uint(i.(float64))>>uint(level.(float64)))) & int(float64(31)))
					_ = subidx
					Pv_aset.X_invoke_Arity3(ret, subidx, do_assoc.X_invoke_Arity5(pv, (level.(float64)-float64(5)), Pv_aget.X_invoke_Arity2(node, subidx).(interface{}), i, val).(*CljsCoreVectorNode)).(interface{})
					return ret
				}
			}
		}
	})
}(&AFn{})

var Pop_tail = func(pop_tail *AFn) *AFn {
	return Fn(pop_tail, func(pv interface{}, level interface{}, node interface{}) interface{} {
		{
			var subidx = float64(int(float64(uint((Native_get_instance_field.X_invoke_Arity2(pv, "Cnt").(float64)-float64(2)))>>uint(level.(float64)))) & int(float64(31)))
			_ = subidx
			if level.(float64) > float64(5) {
				{
					var new_child = pop_tail.X_invoke_Arity3(pv, (level.(float64) - float64(5)), Pv_aget.X_invoke_Arity2(node, subidx).(interface{}))
					_ = new_child
					if (new_child == nil) && (subidx == float64(0)) {
						return nil
					} else {
						{
							var ret = Pv_clone_node.X_invoke_Arity1(node).(*CljsCoreVectorNode)
							_ = ret
							Pv_aset.X_invoke_Arity3(ret, subidx, new_child).(interface{})
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
						Pv_aset.X_invoke_Arity3(ret, subidx, nil).(interface{})
						return ret
					}

				}
			}
		}
	})
}(&AFn{})

type CljsCorePersistentVector struct {
	Meta    interface{}
	Cnt     interface{}
	Shift   interface{}
	Root    interface{}
	Tail    interface{}
	X__hash interface{}
}

func (_ *CljsCorePersistentVector) CljsCoreObject__() {}
func (self__ *CljsCorePersistentVector) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCorePersistentVector) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCorePersistentVector) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCorePersistentVector) CljsCoreILookup__() {}
func (self__ *CljsCorePersistentVector) X_lookup_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return _lookup.X_invoke_Arity3(coll___1, k, nil).(interface{})
	}
}

func (self__ *CljsCorePersistentVector) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if reflect.TypeOf(k).Kind() == reflect.Float64 {
			return coll___1.X_nth_Arity3(k, not_found).(interface{})
		} else {
			return not_found
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIKVReduce__() {}
func (self__ *CljsCorePersistentVector) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var v___1 = self__
		_ = v___1
		{
			var step_init = []interface{}{float64(0), init}
			_ = step_init
			{
				var i = float64(0)
				_ = i
				for {
					if i < self__.Cnt.(float64) {
						{
							var arr = Unchecked_array_for.X_invoke_Arity2(v___1, i).(interface{})
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
												var init___2 = f.(CljsCoreIFn).X_invoke_Arity3(init___1, (j + i), (arr.([]interface{})[int(j)])).(interface{})
												_ = init___2
												if Reduced_QMARK_.X_invoke_Arity1(init___2) {
													return init___2
												} else {
													{
														var G__688 = (j + float64(1))
														var G__689 = init___2
														j = G__688
														init___1 = G__689
														continue
													}
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
								if Reduced_QMARK_.X_invoke_Arity1(init___1) {
									return Deref.X_invoke_Arity1(init___1).(interface{})
								} else {
									{
										var G__690 = (i + (step_init[int(float64(0))]).(float64))
										i = G__690
										continue
									}
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
}

func (_ *CljsCorePersistentVector) CljsCoreIIndexed__() {}
func (self__ *CljsCorePersistentVector) X_nth_Arity2(n interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (Array_for.X_invoke_Arity2(coll___1, n).(interface{}).([]interface{})[int(float64(int(n.(float64))&int(float64(31))))])
	}
}

func (self__ *CljsCorePersistentVector) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (float64(0) <= n.(float64)) && (n.(float64) < self__.Cnt.(float64)) {
			return (Unchecked_array_for.X_invoke_Arity2(coll___1, n).(interface{}).([]interface{})[int(float64(int(n.(float64))&int(float64(31))))])
		} else {
			return not_found
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIVector__() {}
func (self__ *CljsCorePersistentVector) X_assoc_n_Arity3(n interface{}, val interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (float64(0) <= n.(float64)) && (n.(float64) < self__.Cnt.(float64)) {
			if Tail_off.X_invoke_Arity1(coll___1).(float64) <= n.(float64) {
				{
					var new_tail = Aclone.X_invoke_Arity1(self__.Tail).([]interface{})
					_ = new_tail
					new_tail[int(float64(int(n.(float64))&int(float64(31))))] = val
					return (&CljsCorePersistentVector{self__.Meta, self__.Cnt, self__.Shift, self__.Root, new_tail, nil})
				}
			} else {
				return (&CljsCorePersistentVector{self__.Meta, self__.Cnt, self__.Shift, Do_assoc.X_invoke_Arity5(coll___1, self__.Shift, self__.Root, n, val).(*CljsCoreVectorNode), self__.Tail, nil})
			}
		} else {
			if n.(float64) == self__.Cnt.(float64) {
				return coll___1.X_conj_Arity2(val)
			} else {
				panic((&js.Error{("Index " + Str.X_invoke_Arity1(n).(string) + " out of bounds  [0," + Str.X_invoke_Arity1(self__.Cnt).(string) + "]")}))

			}
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIMeta__() {}
func (self__ *CljsCorePersistentVector) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCorePersistentVector) CljsCoreICloneable__() {}
func (self__ *CljsCorePersistentVector) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCorePersistentVector{self__.Meta, self__.Cnt, self__.Shift, self__.Root, self__.Tail, self__.X__hash})
	}
}

func (_ *CljsCorePersistentVector) CljsCoreICounted__() {}
func (self__ *CljsCorePersistentVector) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Cnt.(float64)
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIMapEntry__() {}
func (self__ *CljsCorePersistentVector) X_key_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_nth_Arity2(float64(0)).(interface{})
	}
}

func (self__ *CljsCorePersistentVector) X_val_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_nth_Arity2(float64(1)).(interface{})
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIStack__() {}
func (self__ *CljsCorePersistentVector) X_peek_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Cnt.(float64) > float64(0) {
			return coll___1.X_nth_Arity2((self__.Cnt.(float64) - float64(1))).(interface{})
		} else {
			return nil
		}
	}
}

func (self__ *CljsCorePersistentVector) X_pop_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Cnt.(float64) == float64(0) {
			panic((&js.Error{"Can't pop empty vector"}))
		} else {
			if float64(1) == self__.Cnt.(float64) {
				return CljsCorePersistentVector_EMPTY.(CljsCoreIWithMeta).X_with_meta_Arity2(self__.Meta)
			} else {
				if float64(1) < (self__.Cnt.(float64) - Tail_off.X_invoke_Arity1(coll___1).(float64)) {
					return (&CljsCorePersistentVector{self__.Meta, (self__.Cnt.(float64) - float64(1)), self__.Shift, self__.Root, Native_invoke_instance_method.X_invoke_Arity3(self__.Tail, "Slice", []interface{}{float64(0), float64(-1)}), nil})
				} else {
					{
						var new_tail = Unchecked_array_for.X_invoke_Arity2(coll___1, (self__.Cnt.(float64) - float64(2))).(interface{})
						var nr = Pop_tail.X_invoke_Arity3(coll___1, self__.Shift, self__.Root)
						var new_root = func() interface{} {
							if nr == nil {
								return CljsCorePersistentVector_EMPTY_NODE
							} else {
								return nr
							}
						}()
						var cnt_1 = (self__.Cnt.(float64) - float64(1))
						_, _, _, _ = new_tail, nr, new_root, cnt_1
						if (float64(5) < self__.Shift.(float64)) && (Pv_aget.X_invoke_Arity2(new_root, float64(1)).(interface{}) == nil) {
							return (&CljsCorePersistentVector{self__.Meta, cnt_1, (self__.Shift.(float64) - float64(5)), Pv_aget.X_invoke_Arity2(new_root, float64(0)).(interface{}), new_tail, nil})
						} else {
							return (&CljsCorePersistentVector{self__.Meta, cnt_1, self__.Shift, new_root, new_tail, nil})
						}
					}

				}
			}
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIReversible__() {}
func (self__ *CljsCorePersistentVector) X_rseq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Cnt.(float64) > float64(0) {
			return (&CljsCoreRSeq{coll___1, (self__.Cnt.(float64) - float64(1)), nil})
		} else {
			return nil
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIHash__() {}
func (self__ *CljsCorePersistentVector) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIEquiv__() {}
func (self__ *CljsCorePersistentVector) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIEditableCollection__() {}
func (self__ *CljsCorePersistentVector) X_as_transient_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreTransientVector{self__.Cnt, self__.Shift, Tv_editable_root.X_invoke_Arity1(self__.Root).(interface{}), Tv_editable_tail.X_invoke_Arity1(self__.Tail).(interface{})})
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCorePersistentVector) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCorePersistentVector_EMPTY, self__.Meta)
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIReduce__() {}
func (self__ *CljsCorePersistentVector) X_reduce_Arity2(f interface{}) interface{} {
	{
		var v___1 = self__
		_ = v___1
		return Ci_reduce.X_invoke_Arity2(v___1, f).(interface{})
	}
}

func (self__ *CljsCorePersistentVector) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var v___1 = self__
		_ = v___1
		return Ci_reduce.X_invoke_Arity3(v___1, f, start)
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIAssociative__() {}
func (self__ *CljsCorePersistentVector) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if reflect.TypeOf(k).Kind() == reflect.Float64 {
			return coll___1.X_assoc_n_Arity3(k, v)
		} else {
			panic((&js.Error{"Vector's key for assoc must be a number."}))
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreISeqable__() {}
func (self__ *CljsCorePersistentVector) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Cnt.(float64) == float64(0) {
			return nil
		} else {
			if self__.Cnt.(float64) <= float64(32) {
				return (&CljsCoreIndexedSeq{self__.Tail, float64(0)})
			} else {
				return Chunked_seq.X_invoke_Arity4(coll___1, First_array_for_longvec.X_invoke_Arity1(coll___1).(interface{}), float64(0), float64(0)).(interface{})

			}
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreISequential__() {}
func (_ *CljsCorePersistentVector) CljsCoreIWithMeta__()   {}
func (self__ *CljsCorePersistentVector) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentVector{meta___1, self__.Cnt, self__.Shift, self__.Root, self__.Tail, self__.X__hash})
	}
}

func (_ *CljsCorePersistentVector) CljsCoreICollection__() {}
func (self__ *CljsCorePersistentVector) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (self__.Cnt.(float64) - Tail_off.X_invoke_Arity1(coll___1).(float64)) < float64(32) {
			{
				var len = float64(len(self__.Tail.([]interface{})))
				var new_tail = make([]interface{}, int((len + float64(1))))
				_, _ = len, new_tail
				{
					var n__77528__auto___691 = len
					_ = n__77528__auto___691
					{
						var i_692 = float64(0)
						_ = i_692
						for {
							if i_692 < n__77528__auto___691 {
								new_tail[int(i_692)] = (self__.Tail.([]interface{})[int(i_692)])
								{
									var G__693 = (i_692 + float64(1))
									i_692 = G__693
									continue
								}
							} else {
							}
							break
						}
					}
				}
				new_tail[int(len)] = o
				return (&CljsCorePersistentVector{self__.Meta, (self__.Cnt.(float64) + float64(1)), self__.Shift, self__.Root, new_tail, nil})
			}
		} else {
			{
				var root_overflow_QMARK_ = (float64(uint(self__.Cnt.(float64))>>uint(float64(5))) > float64(int(float64(1))<<uint(self__.Shift.(float64))))
				var new_shift = func() interface{} {
					if root_overflow_QMARK_ {
						return (self__.Shift.(float64) + float64(5))
					} else {
						return self__.Shift
					}
				}()
				var new_root = func() interface{} {
					if root_overflow_QMARK_ {
						return func() interface{} {
							var n_r = Pv_fresh_node.X_invoke_Arity1(nil).(*CljsCoreVectorNode)
							_ = n_r
							Pv_aset.X_invoke_Arity3(n_r, float64(0), self__.Root).(interface{})
							Pv_aset.X_invoke_Arity3(n_r, float64(1), New_path.X_invoke_Arity3(nil, self__.Shift, (&CljsCoreVectorNode{nil, self__.Tail})).(interface{})).(interface{})
							return n_r
						}()
					} else {
						return Push_tail.X_invoke_Arity4(coll___1, self__.Shift, self__.Root, (&CljsCoreVectorNode{nil, self__.Tail})).(*CljsCoreVectorNode)
					}
				}()
				_, _, _ = root_overflow_QMARK_, new_shift, new_root
				return (&CljsCorePersistentVector{self__.Meta, (self__.Cnt.(float64) + float64(1)), new_shift, new_root, []interface{}{o}, nil})
			}
		}
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIFn__() {}
func (self__ *CljsCorePersistentVector) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity1(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_nth_Arity2(k).(interface{})
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_nth_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCorePersistentVector) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

var X__GT_PersistentVector = func(__GT_PersistentVector *AFn) *AFn {
	return Fn(__GT_PersistentVector, func(meta interface{}, cnt interface{}, shift interface{}, root interface{}, tail interface{}, __hash interface{}) interface{} {
		return (&CljsCorePersistentVector{meta, cnt, shift, root, tail, __hash})
	})
}(&AFn{})

var CljsCorePersistentVector_EMPTY_NODE = (&CljsCoreVectorNode{nil, make([]interface{}, int(float64(32)))})

var CljsCorePersistentVector_EMPTY = (&CljsCorePersistentVector{nil, float64(0), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{}, float64(0)})

var CljsCorePersistentVector_FromArray = func(G__694 *AFn) *AFn {
	return Fn(G__694, func(xs interface{}, no_clone bool) interface{} {
		{
			var l = float64(len(xs.([]interface{})))
			var xs___1 = func() interface{} {
				if no_clone {
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
					var node = xs___1.Slice(float64(0), float64(32))
					var v = (&CljsCorePersistentVector{nil, float64(32), float64(5), CljsCorePersistentVector_EMPTY_NODE, node, nil})
					_, _ = node, v
					{
						var i = float64(32)
						var out = v.X_as_transient_Arity1()
						_, _ = i, out
						for {
							if i < l {
								{
									var G__695 = (i + float64(1))
									var G__696 = Conj_BANG_.X_invoke_Arity2(out, (xs___1.([]interface{})[int(i)])).(interface{})
									i = G__695
									out = G__696
									continue
								}
							} else {
								return Persistent_BANG_.X_invoke_Arity1(out).(interface{})
							}
						}
					}
				}
			}
		}
	})
}(&AFn{})

var Vec = func(vec *AFn) *AFn {
	return Fn(vec, func(coll interface{}) interface{} {
		return Reduce.X_invoke_Arity3(X_conj_BANG_, CljsCorePersistentVector_EMPTY.(CljsCoreIEditableCollection).X_as_transient_Arity1(), coll).X_persistent_BANG__Arity1()
	})
}(&AFn{})

/**
* @param {...*} var_args
 */
var Vector = func(vector *AFn) *AFn {
	return Fn(vector, func(args__ ...interface{}) interface{} {
		var args = Array_seq.X_invoke_Arity1(args__[0:])
		_ = args
		if (func() bool { _, instanceof := args.(*CljsCoreIndexedSeq); return instanceof }()) && (Native_get_instance_field.X_invoke_Arity2(args, "I").(float64) == float64(0)) {
			return CljsCorePersistentVector_FromArray.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(args, "Arr"), true)
		} else {
			return Vec.X_invoke_Arity1(args).(interface{})
		}
	})
}(&AFn{})

type CljsCoreChunkedSeq struct {
	Vec     interface{}
	Node    interface{}
	I       interface{}
	Off     interface{}
	Meta    interface{}
	X__hash interface{}
}

func (_ *CljsCoreChunkedSeq) CljsCoreObject__() {}
func (self__ *CljsCoreChunkedSeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreChunkedSeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreChunkedSeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreINext__() {}
func (self__ *CljsCoreChunkedSeq) X_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (self__.Off.(float64) + float64(1)) < float64(len(self__.Node.([]interface{}))) {
			{
				var s = Chunked_seq.X_invoke_Arity4(self__.Vec, self__.Node, self__.I, (self__.Off.(float64) + float64(1))).(interface{})
				_ = s
				if s == nil {
					return nil
				} else {
					return s
				}
			}
		} else {
			return coll___1.X_chunked_next_Arity1().(interface{})
		}
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreIHash__() {}
func (self__ *CljsCoreChunkedSeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreIEquiv__() {}
func (self__ *CljsCoreChunkedSeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreChunkedSeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCorePersistentVector_EMPTY, self__.Meta)
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreIReduce__() {}
func (self__ *CljsCoreChunkedSeq) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Ci_reduce.X_invoke_Arity2(Subvec.X_invoke_Arity3(self__.Vec, (self__.I.(float64)+self__.Off.(float64)), Count.X_invoke_Arity1(self__.Vec).(float64)).(interface{}), f).(interface{})
	}
}

func (self__ *CljsCoreChunkedSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Ci_reduce.X_invoke_Arity3(Subvec.X_invoke_Arity3(self__.Vec, (self__.I.(float64)+self__.Off.(float64)), Count.X_invoke_Arity1(self__.Vec).(float64)).(interface{}), f, start)
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreISeq__() {}
func (self__ *CljsCoreChunkedSeq) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (self__.Node.([]interface{})[int(self__.Off.(float64))])
	}
}

func (self__ *CljsCoreChunkedSeq) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (self__.Off.(float64) + float64(1)) < float64(len(self__.Node.([]interface{}))) {
			{
				var s = Chunked_seq.X_invoke_Arity4(self__.Vec, self__.Node, self__.I, (self__.Off.(float64) + float64(1))).(interface{})
				_ = s
				if s == nil {
					return CljsCoreList_EMPTY
				} else {
					return s
				}
			}
		} else {
			return coll___1.X_chunked_rest_Arity1().(interface{})
		}
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreISeqable__() {}
func (self__ *CljsCoreChunkedSeq) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreISequential__() {}
func (_ *CljsCoreChunkedSeq) CljsCoreIChunkedSeq__() {}
func (self__ *CljsCoreChunkedSeq) X_chunked_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Array_chunk.X_invoke_Arity2(self__.Node, self__.Off).(*CljsCoreArrayChunk)
	}
}

func (self__ *CljsCoreChunkedSeq) X_chunked_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var end = (self__.I.(float64) + float64(len(self__.Node.([]interface{}))))
			_ = end
			if end < self__.Vec.(CljsCoreICounted).X_count_Arity1() {
				return Chunked_seq.X_invoke_Arity4(self__.Vec, Unchecked_array_for.X_invoke_Arity2(self__.Vec, end).(interface{}), end, float64(0)).(interface{})
			} else {
				return CljsCoreList_EMPTY
			}
		}
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreIWithMeta__() {}
func (self__ *CljsCoreChunkedSeq) X_with_meta_Arity2(m interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Chunked_seq.X_invoke_Arity5(self__.Vec, self__.Node, self__.I, self__.Off, m).(interface{})
	}
}

func (self__ *CljsCoreChunkedSeq) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreICollection__() {}
func (self__ *CljsCoreChunkedSeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(*CljsCoreCons)
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreASeq__()         {}
func (_ *CljsCoreChunkedSeq) CljsCoreIChunkedNext__() {}
func (self__ *CljsCoreChunkedSeq) X_chunked_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var end = (self__.I.(float64) + float64(len(self__.Node.([]interface{}))))
			_ = end
			if end < self__.Vec.(CljsCoreICounted).X_count_Arity1() {
				return Chunked_seq.X_invoke_Arity4(self__.Vec, Unchecked_array_for.X_invoke_Arity2(self__.Vec, end).(interface{}), end, float64(0)).(interface{})
			} else {
				return nil
			}
		}
	}
}

var X__GT_ChunkedSeq = func(__GT_ChunkedSeq *AFn) *AFn {
	return Fn(__GT_ChunkedSeq, func(vec interface{}, node interface{}, i interface{}, off interface{}, meta interface{}, __hash interface{}) interface{} {
		return (&CljsCoreChunkedSeq{vec, node, i, off, meta, __hash})
	})
}(&AFn{})

var Chunked_seq = func(chunked_seq *AFn) *AFn {
	return Fn(chunked_seq, func(vec interface{}, i interface{}, off interface{}) interface{} {
		return (&CljsCoreChunkedSeq{vec, Array_for.X_invoke_Arity2(vec, i).(interface{}), i, off, nil, nil})
	}, func(vec interface{}, node interface{}, i interface{}, off interface{}) interface{} {
		return (&CljsCoreChunkedSeq{vec, node, i, off, nil, nil})
	}, func(vec interface{}, node interface{}, i interface{}, off interface{}, meta interface{}) interface{} {
		return (&CljsCoreChunkedSeq{vec, node, i, off, meta, nil})
	})
}(&AFn{})

type CljsCoreSubvec struct {
	Meta    interface{}
	V       interface{}
	Start   interface{}
	End     interface{}
	X__hash interface{}
}

func (_ *CljsCoreSubvec) CljsCoreObject__() {}
func (self__ *CljsCoreSubvec) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreSubvec) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreSubvec) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreSubvec) CljsCoreILookup__() {}
func (self__ *CljsCoreSubvec) X_lookup_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return _lookup.X_invoke_Arity3(coll___1, k, nil).(interface{})
	}
}

func (self__ *CljsCoreSubvec) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if reflect.TypeOf(k).Kind() == reflect.Float64 {
			return coll___1.X_nth_Arity3(k, not_found).(interface{})
		} else {
			return not_found
		}
	}
}

func (_ *CljsCoreSubvec) CljsCoreIIndexed__() {}
func (self__ *CljsCoreSubvec) X_nth_Arity2(n interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (n.(float64) < float64(0)) || (self__.End.(float64) <= (self__.Start.(float64) + n.(float64))) {
			return Vector_index_out_of_bounds.X_invoke_Arity2(n, (self__.End.(float64) - self__.Start.(float64))).(interface{})
		} else {
			return _nth.X_invoke_Arity2(self__.V, (self__.Start.(float64) + n.(float64))).(interface{})
		}
	}
}

func (self__ *CljsCoreSubvec) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (n.(float64) < float64(0)) || (self__.End.(float64) <= (self__.Start.(float64) + n.(float64))) {
			return not_found
		} else {
			return _nth.X_invoke_Arity3(self__.V, (self__.Start.(float64) + n.(float64)), not_found)
		}
	}
}

func (_ *CljsCoreSubvec) CljsCoreIVector__() {}
func (self__ *CljsCoreSubvec) X_assoc_n_Arity3(n interface{}, val interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var v_pos = (self__.Start.(float64) + n.(float64))
			_ = v_pos
			return Build_subvec.X_invoke_Arity5(self__.Meta, Assoc.X_invoke_Arity3(self__.V, v_pos, val).(interface{}), self__.Start, func() interface{} {
				var x__76939__auto__ = self__.End
				var y__76940__auto__ = (v_pos + float64(1))
				_, _ = x__76939__auto__, y__76940__auto__
				return func() float64 {
					if x__76939__auto__ > y__76940__auto__ {
						return x__76939__auto__
					} else {
						return y__76940__auto__
					}
				}()
			}(), nil).(interface{})
		}
	}
}

func (_ *CljsCoreSubvec) CljsCoreIMeta__() {}
func (self__ *CljsCoreSubvec) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreSubvec) CljsCoreICloneable__() {}
func (self__ *CljsCoreSubvec) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCoreSubvec{self__.Meta, self__.V, self__.Start, self__.End, self__.X__hash})
	}
}

func (_ *CljsCoreSubvec) CljsCoreICounted__() {}
func (self__ *CljsCoreSubvec) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return (self__.End.(float64) - self__.Start.(float64))
	}
}

func (_ *CljsCoreSubvec) CljsCoreIStack__() {}
func (self__ *CljsCoreSubvec) X_peek_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.V.(CljsCoreIIndexed).X_nth_Arity2((self__.End.(float64) - float64(1))).(interface{})
	}
}

func (self__ *CljsCoreSubvec) X_pop_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Start.(float64) == self__.End.(float64) {
			panic((&js.Error{"Can't pop empty vector"}))
		} else {
			return Build_subvec.X_invoke_Arity5(self__.Meta, self__.V, self__.Start, (self__.End.(float64) - float64(1)), nil).(interface{})
		}
	}
}

func (_ *CljsCoreSubvec) CljsCoreIReversible__() {}
func (self__ *CljsCoreSubvec) X_rseq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if !(self__.Start.(float64) == self__.End.(float64)) {
			return (&CljsCoreRSeq{coll___1, ((self__.End.(float64) - self__.Start.(float64)) - float64(1)), nil})
		} else {
			return nil
		}
	}
}

func (_ *CljsCoreSubvec) CljsCoreIHash__() {}
func (self__ *CljsCoreSubvec) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreSubvec) CljsCoreIEquiv__() {}
func (self__ *CljsCoreSubvec) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreSubvec) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreSubvec) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCorePersistentVector_EMPTY, self__.Meta)
	}
}

func (_ *CljsCoreSubvec) CljsCoreIReduce__() {}
func (self__ *CljsCoreSubvec) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Ci_reduce.X_invoke_Arity2(coll___1, f).(interface{})
	}
}

func (self__ *CljsCoreSubvec) X_reduce_Arity3(f interface{}, start___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Ci_reduce.X_invoke_Arity3(coll___1, f, start___1)
	}
}

func (_ *CljsCoreSubvec) CljsCoreIAssociative__() {}
func (self__ *CljsCoreSubvec) X_assoc_Arity3(key interface{}, val interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if reflect.TypeOf(key).Kind() == reflect.Float64 {
			return coll___1.X_assoc_n_Arity3(key, val)
		} else {
			panic((&js.Error{"Subvec's key for assoc must be a number."}))
		}
	}
}

func (_ *CljsCoreSubvec) CljsCoreISeqable__() {}
func (self__ *CljsCoreSubvec) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var subvec_seq = func(coll___1 *CljsCoreSubvec) *AFn {
				return func(subvec_seq *AFn) *AFn {
					return Fn(subvec_seq, func(i interface{}) interface{} {
						if i.(float64) == self__.End.(float64) {
							return nil
						} else {
							return Cons.X_invoke_Arity2(self__.V.(CljsCoreIIndexed).X_nth_Arity2(i).(interface{}), (&CljsCoreLazySeq{nil, func(coll___1 *CljsCoreSubvec) *AFn {
								return func(G__697 *AFn) *AFn {
									return Fn(G__697, func() interface{} {
										return subvec_seq.X_invoke_Arity1((i.(float64) + float64(1)))
									})
								}(&AFn{})
							}(coll___1), nil, nil})).(*CljsCoreCons)
						}
					})
				}(&AFn{})
			}(coll___1)
			_ = subvec_seq
			return subvec_seq.(CljsCoreIFn).X_invoke_Arity1(self__.Start)
		}
	}
}

func (_ *CljsCoreSubvec) CljsCoreISequential__() {}
func (_ *CljsCoreSubvec) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreSubvec) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Build_subvec.X_invoke_Arity5(meta___1, self__.V, self__.Start, self__.End, self__.X__hash).(interface{})
	}
}

func (_ *CljsCoreSubvec) CljsCoreICollection__() {}
func (self__ *CljsCoreSubvec) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Build_subvec.X_invoke_Arity5(self__.Meta, self__.V.(CljsCoreIVector).X_assoc_n_Arity3(self__.End, o), self__.Start, (self__.End.(float64) + float64(1)), nil).(interface{})
	}
}

func (_ *CljsCoreSubvec) CljsCoreIFn__() {}
func (self__ *CljsCoreSubvec) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity1(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_nth_Arity2(k).(interface{})
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_nth_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCoreSubvec) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

var X__GT_Subvec = func(__GT_Subvec *AFn) *AFn {
	return Fn(__GT_Subvec, func(meta interface{}, v interface{}, start interface{}, end interface{}, __hash interface{}) interface{} {
		return (&CljsCoreSubvec{meta, v, start, end, __hash})
	})
}(&AFn{})

var Build_subvec = func(build_subvec *AFn) *AFn {
	return Fn(build_subvec, func(meta interface{}, v interface{}, start interface{}, end interface{}, __hash interface{}) interface{} {
		for {
			if func() bool { _, instanceof := v.(*CljsCoreSubvec); return instanceof }() {
				{
					var G__698 = meta
					var G__699 = Native_get_instance_field.X_invoke_Arity2(v, "V")
					var G__700 = (Native_get_instance_field.X_invoke_Arity2(v, "Start").(float64) + start.(float64))
					var G__701 = (Native_get_instance_field.X_invoke_Arity2(v, "Start").(float64) + end.(float64))
					var G__702 = __hash
					meta = G__698
					v = G__699
					start = G__700
					end = G__701
					__hash = G__702
					continue
				}
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

/**
* Returns a persistent vector of the items in vector from
* start (inclusive) to end (exclusive).  If end is not supplied,
* defaults to (count vector). This operation is O(1) and very fast, as
* the resulting vector shares structure with the original and no
* trimming is done.
 */
var Subvec = func(subvec *AFn) *AFn {
	return Fn(subvec, func(v interface{}, start interface{}) interface{} {
		return subvec.X_invoke_Arity3(v, start, Count.X_invoke_Arity1(v).(float64)).(*CljsCoreSubvec)
	}, func(v interface{}, start interface{}, end interface{}) interface{} {
		return Build_subvec.X_invoke_Arity5(nil, v, start, end, nil).(*CljsCoreSubvec)
	})
}(&AFn{})

var Tv_ensure_editable = func(tv_ensure_editable *AFn) *AFn {
	return Fn(tv_ensure_editable, func(edit interface{}, node interface{}) interface{} {
		if edit == Native_get_instance_field.X_invoke_Arity2(node, "Edit") {
			return node
		} else {
			return (&CljsCoreVectorNode{edit, Aclone.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(node, "Arr")).([]interface{})})
		}
	})
}(&AFn{})

var Tv_editable_root = func(tv_editable_root *AFn) *AFn {
	return Fn(tv_editable_root, func(node interface{}) interface{} {
		return (&CljsCoreVectorNode{func() interface{} {
			var obj706 = map[string]interface{}{}
			_ = obj706
			return obj706
		}(), Aclone.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(node, "Arr")).([]interface{})})
	})
}(&AFn{})

var Tv_editable_tail = func(tv_editable_tail *AFn) *AFn {
	return Fn(tv_editable_tail, func(tl interface{}) interface{} {
		{
			var ret = make([]interface{}, int(float64(32)))
			_ = ret
			Array_copy.X_invoke_Arity5(tl, float64(0), ret, float64(0), float64(len(tl.([]interface{})))).(interface{})
			return ret
		}
	})
}(&AFn{})

var Tv_push_tail = func(tv_push_tail *AFn) *AFn {
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
						var child = Pv_aget.X_invoke_Arity2(ret, subidx).(interface{})
						_ = child
						if !(child == nil) {
							return tv_push_tail.X_invoke_Arity4(tv, (level.(float64) - float64(5)), child, tail_node)
						} else {
							return New_path.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tv, "Root"), "Edit"), (level.(float64) - float64(5)), tail_node).(interface{})
						}
					}()
				}
			}()).(interface{})
			return ret
		}
	})
}(&AFn{})

var Tv_pop_tail = func(tv_pop_tail *AFn) *AFn {
	return Fn(tv_pop_tail, func(tv interface{}, level interface{}, node interface{}) interface{} {
		{
			var node___1 = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tv, "Root"), "Edit"), node)
			var subidx = float64(int(float64(uint((Native_get_instance_field.X_invoke_Arity2(tv, "Cnt").(float64)-float64(2)))>>uint(level.(float64)))) & int(float64(31)))
			_, _ = node___1, subidx
			if level.(float64) > float64(5) {
				{
					var new_child = tv_pop_tail.X_invoke_Arity3(tv, (level.(float64) - float64(5)), Pv_aget.X_invoke_Arity2(node___1, subidx).(interface{}))
					_ = new_child
					if (new_child == nil) && (subidx == float64(0)) {
						return nil
					} else {
						Pv_aset.X_invoke_Arity3(node___1, subidx, new_child).(interface{})
						return node___1
					}
				}
			} else {
				if subidx == float64(0) {
					return nil
				} else {
					Pv_aset.X_invoke_Arity3(node___1, subidx, nil).(interface{})
					return node___1

				}
			}
		}
	})
}(&AFn{})

var Unchecked_editable_array_for = func(unchecked_editable_array_for *AFn) *AFn {
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
							{
								var G__707 = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(root, "Edit"), Pv_aget.X_invoke_Arity2(node, float64(int(float64(uint(i.(float64))>>uint(level.(float64))))&int(float64(31)))).(interface{}))
								var G__708 = (level.(float64) - float64(5))
								node = G__707
								level = G__708
								continue
							}
						} else {
							return Native_get_instance_field.X_invoke_Arity2(node, "Arr")
						}
					}
				}
			}
		}
	})
}(&AFn{})

type CljsCoreTransientVector struct {
	Cnt   interface{}
	Shift interface{}
	Root  interface{}
	Tail  interface{}
}

func (_ *CljsCoreTransientVector) CljsCoreIFn__() {}
func (self__ *CljsCoreTransientVector) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity1(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCoreTransientVector) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

func (_ *CljsCoreTransientVector) CljsCoreILookup__() {}
func (self__ *CljsCoreTransientVector) X_lookup_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return _lookup.X_invoke_Arity3(coll___1, k, nil).(interface{})
	}
}

func (self__ *CljsCoreTransientVector) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if reflect.TypeOf(k).Kind() == reflect.Float64 {
			return coll___1.X_nth_Arity3(k, not_found).(interface{})
		} else {
			return not_found
		}
	}
}

func (_ *CljsCoreTransientVector) CljsCoreIIndexed__() {}
func (self__ *CljsCoreTransientVector) X_nth_Arity2(n interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit") {
			return (Array_for.X_invoke_Arity2(coll___1, n).(interface{}).([]interface{})[int(float64(int(n.(float64))&int(float64(31))))])
		} else {
			panic((&js.Error{"nth after persistent!"}))
		}
	}
}

func (self__ *CljsCoreTransientVector) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (float64(0) <= n.(float64)) && (n.(float64) < self__.Cnt.(float64)) {
			return _nth.X_invoke_Arity2(coll___1, n).(interface{})
		} else {
			return not_found
		}
	}
}

func (_ *CljsCoreTransientVector) CljsCoreICounted__() {}
func (self__ *CljsCoreTransientVector) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		if Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit") {
			return self__.Cnt.(float64)
		} else {
			panic((&js.Error{"count after persistent!"}))
		}
	}
}

func (_ *CljsCoreTransientVector) CljsCoreITransientVector__() {}
func (self__ *CljsCoreTransientVector) X_assoc_n_BANG__Arity3(n interface{}, val interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit") {
			if (float64(0) <= n.(float64)) && (n.(float64) < self__.Cnt.(float64)) {
				if Tail_off.X_invoke_Arity1(tcoll___1).(float64) <= n.(float64) {
					self__.Tail.([]interface{})[int(float64(int(n.(float64))&int(float64(31))))] = val
					return tcoll___1
				} else {
					{
						var new_root = func(tcoll___1 *CljsCoreTransientVector) *AFn {
							return func(go_ *AFn) *AFn {
								return Fn(go_, func(level interface{}, node interface{}) interface{} {
									{
										var node___1 = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit"), node)
										_ = node___1
										if level.(float64) == float64(0) {
											Pv_aset.X_invoke_Arity3(node___1, float64(int(n.(float64))&int(float64(31))), val).(interface{})
											return node___1
										} else {
											{
												var subidx = float64(int(float64(uint(n.(float64))>>uint(level.(float64)))) & int(float64(31)))
												_ = subidx
												Pv_aset.X_invoke_Arity3(node___1, subidx, go_.X_invoke_Arity2((level.(float64)-float64(5)), Pv_aget.X_invoke_Arity2(node___1, subidx).(interface{}))).(interface{})
												return node___1
											}
										}
									}
								})
							}(&AFn{})
						}(tcoll___1).X_invoke_Arity2(self__.Shift, self__.Root)
						_ = new_root
						self__.Root = new_root

						return tcoll___1
					}
				}
			} else {
				if n.(float64) == self__.Cnt.(float64) {
					return tcoll___1.X_conj_BANG__Arity2(val)
				} else {
					panic((&js.Error{("Index " + Str.X_invoke_Arity1(n).(string) + " out of bounds for TransientVector of length" + Str.X_invoke_Arity1(self__.Cnt).(string))}))

				}
			}
		} else {
			panic((&js.Error{"assoc! after persistent!"}))
		}
	}
}

func (self__ *CljsCoreTransientVector) X_pop_BANG__Arity1() interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit") {
			if self__.Cnt.(float64) == float64(0) {
				panic((&js.Error{"Can't pop empty vector"}))
			} else {
				if float64(1) == self__.Cnt.(float64) {
					self__.Cnt = float64(0)

					return tcoll___1
				} else {
					if float64(int((self__.Cnt.(float64)-float64(1)))&int(float64(31))) > float64(0) {
						self__.Cnt = (self__.Cnt.(float64) - float64(1))

						return tcoll___1
					} else {
						{
							var new_tail = Unchecked_editable_array_for.X_invoke_Arity2(tcoll___1, (self__.Cnt.(float64) - float64(2))).(interface{})
							var new_root = func() interface{} {
								var nr = Tv_pop_tail.X_invoke_Arity3(tcoll___1, self__.Shift, self__.Root)
								_ = nr
								if !(nr == nil) {
									return nr
								} else {
									return (&CljsCoreVectorNode{Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit"), make([]interface{}, int(float64(32)))})
								}
							}()
							_, _ = new_tail, new_root
							if (float64(5) < self__.Shift.(float64)) && (Pv_aget.X_invoke_Arity2(new_root, float64(1)).(interface{}) == nil) {
								{
									var new_root___1 = Tv_ensure_editable.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit"), Pv_aget.X_invoke_Arity2(new_root, float64(0)).(interface{}))
									_ = new_root___1
									self__.Root = new_root___1

									self__.Shift = (self__.Shift.(float64) - float64(5))

									self__.Cnt = (self__.Cnt.(float64) - float64(1))

									self__.Tail = new_tail

									return tcoll___1
								}
							} else {
								self__.Root = new_root

								self__.Cnt = (self__.Cnt.(float64) - float64(1))

								self__.Tail = new_tail

								return tcoll___1
							}
						}

					}
				}
			}
		} else {
			panic((&js.Error{"pop! after persistent!"}))
		}
	}
}

func (_ *CljsCoreTransientVector) CljsCoreITransientAssociative__() {}
func (self__ *CljsCoreTransientVector) X_assoc_BANG__Arity3(key interface{}, val interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if reflect.TypeOf(key).Kind() == reflect.Float64 {
			return tcoll___1.X_assoc_n_BANG__Arity3(key, val)
		} else {
			panic((&js.Error{"TransientVector's key for assoc! must be a number."}))
		}
	}
}

func (_ *CljsCoreTransientVector) CljsCoreITransientCollection__() {}
func (self__ *CljsCoreTransientVector) X_conj_BANG__Arity2(o interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit") {
			if (self__.Cnt.(float64) - Tail_off.X_invoke_Arity1(tcoll___1).(float64)) < float64(32) {
				self__.Tail.([]interface{})[int(float64(int(self__.Cnt.(float64))&int(float64(31))))] = o
				self__.Cnt = (self__.Cnt.(float64) + float64(1))

				return tcoll___1
			} else {
				{
					var tail_node = (&CljsCoreVectorNode{Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit"), self__.Tail})
					var new_tail = make([]interface{}, int(float64(32)))
					_, _ = tail_node, new_tail
					new_tail[int(float64(0))] = o
					self__.Tail = new_tail

					if float64(uint(self__.Cnt.(float64))>>uint(float64(5))) > float64(int(float64(1))<<uint(self__.Shift.(float64))) {
						{
							var new_root_array = make([]interface{}, int(float64(32)))
							var new_shift = (self__.Shift.(float64) + float64(5))
							_, _ = new_root_array, new_shift
							new_root_array[int(float64(0))] = self__.Root
							new_root_array[int(float64(1))] = New_path.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit"), self__.Shift, tail_node).(interface{})
							self__.Root = (&CljsCoreVectorNode{Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit"), new_root_array})

							self__.Shift = new_shift

							self__.Cnt = (self__.Cnt.(float64) + float64(1))

							return tcoll___1
						}
					} else {
						{
							var new_root = Tv_push_tail.X_invoke_Arity4(tcoll___1, self__.Shift, self__.Root, tail_node)
							_ = new_root
							self__.Root = new_root

							self__.Cnt = (self__.Cnt.(float64) + float64(1))

							return tcoll___1
						}
					}
				}
			}
		} else {
			panic((&js.Error{"conj! after persistent!"}))
		}
	}
}

func (self__ *CljsCoreTransientVector) X_persistent_BANG__Arity1() interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if Native_get_instance_field.X_invoke_Arity2(self__.Root, "Edit") {
			Native_set_instance_field.X_invoke_Arity3(self__.Root, "Edit", nil)

			{
				var len = (self__.Cnt.(float64) - Tail_off.X_invoke_Arity1(tcoll___1).(float64))
				var trimmed_tail = make([]interface{}, int(len))
				_, _ = len, trimmed_tail
				Array_copy.X_invoke_Arity5(self__.Tail, float64(0), trimmed_tail, float64(0), len).(interface{})
				return (&CljsCorePersistentVector{nil, self__.Cnt, self__.Shift, self__.Root, trimmed_tail, nil})
			}
		} else {
			panic((&js.Error{"persistent! called twice"}))
		}
	}
}

var X__GT_TransientVector = func(__GT_TransientVector *AFn) *AFn {
	return Fn(__GT_TransientVector, func(cnt interface{}, shift interface{}, root interface{}, tail interface{}) interface{} {
		return (&CljsCoreTransientVector{cnt, shift, root, tail})
	})
}(&AFn{})

type CljsCorePersistentQueueSeq struct {
	Meta    interface{}
	Front   interface{}
	Rear    interface{}
	X__hash interface{}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreObject__() {}
func (self__ *CljsCorePersistentQueueSeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCorePersistentQueueSeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCorePersistentQueueSeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIMeta__() {}
func (self__ *CljsCorePersistentQueueSeq) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIHash__() {}
func (self__ *CljsCorePersistentQueueSeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIEquiv__() {}
func (self__ *CljsCorePersistentQueueSeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCorePersistentQueueSeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.Meta)
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreISeq__() {}
func (self__ *CljsCorePersistentQueueSeq) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return First.X_invoke_Arity1(self__.Front)
	}
}

func (self__ *CljsCorePersistentQueueSeq) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var temp__4217__auto__ = Next.Arity1IQ(self__.Front)
			_ = temp__4217__auto__
			if Truth_(temp__4217__auto__) {
				{
					var f1 = temp__4217__auto__
					_ = f1
					return (&CljsCorePersistentQueueSeq{self__.Meta, f1, self__.Rear, nil})
				}
			} else {
				if self__.Rear == nil {
					return coll___1.X_empty_Arity1().(interface{})
				} else {
					return (&CljsCorePersistentQueueSeq{self__.Meta, self__.Rear, nil, nil})
				}
			}
		}
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreISeqable__() {}
func (self__ *CljsCorePersistentQueueSeq) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreISequential__() {}
func (_ *CljsCorePersistentQueueSeq) CljsCoreIWithMeta__()   {}
func (self__ *CljsCorePersistentQueueSeq) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentQueueSeq{meta___1, self__.Front, self__.Rear, self__.X__hash})
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreICollection__() {}
func (self__ *CljsCorePersistentQueueSeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(*CljsCoreCons)
	}
}

var X__GT_PersistentQueueSeq = func(__GT_PersistentQueueSeq *AFn) *AFn {
	return Fn(__GT_PersistentQueueSeq, func(meta interface{}, front interface{}, rear interface{}, __hash interface{}) interface{} {
		return (&CljsCorePersistentQueueSeq{meta, front, rear, __hash})
	})
}(&AFn{})

type CljsCorePersistentQueue struct {
	Meta    interface{}
	Count   interface{}
	Front   interface{}
	Rear    interface{}
	X__hash interface{}
}

func (_ *CljsCorePersistentQueue) CljsCoreObject__() {}
func (self__ *CljsCorePersistentQueue) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCorePersistentQueue) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCorePersistentQueue) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreIMeta__() {}
func (self__ *CljsCorePersistentQueue) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreICloneable__() {}
func (self__ *CljsCorePersistentQueue) X_clone_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentQueue{self__.Meta, self__.Count, self__.Front, self__.Rear, self__.X__hash})
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreICounted__() {}
func (self__ *CljsCorePersistentQueue) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Count.(float64)
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreIStack__() {}
func (self__ *CljsCorePersistentQueue) X_peek_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return First.X_invoke_Arity1(self__.Front)
	}
}

func (self__ *CljsCorePersistentQueue) X_pop_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if Truth_(self__.Front) {
			{
				var temp__4217__auto__ = Next.Arity1IQ(self__.Front)
				_ = temp__4217__auto__
				if Truth_(temp__4217__auto__) {
					{
						var f1 = temp__4217__auto__
						_ = f1
						return (&CljsCorePersistentQueue{self__.Meta, (self__.Count.(float64) - float64(1)), f1, self__.Rear, nil})
					}
				} else {
					return (&CljsCorePersistentQueue{self__.Meta, (self__.Count.(float64) - float64(1)), Seq.Arity1IQ(self__.Rear), CljsCorePersistentVector_EMPTY, nil})
				}
			}
		} else {
			return coll___1
		}
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreIHash__() {}
func (self__ *CljsCorePersistentQueue) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreIEquiv__() {}
func (self__ *CljsCorePersistentQueue) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCorePersistentQueue) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return CljsCorePersistentQueue_EMPTY
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreISeq__() {}
func (self__ *CljsCorePersistentQueue) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return First.X_invoke_Arity1(self__.Front)
	}
}

func (self__ *CljsCorePersistentQueue) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Rest.X_invoke_Arity1(Seq.X_invoke_Arity1(coll___1))
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreISeqable__() {}
func (self__ *CljsCorePersistentQueue) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var rear___1 = Seq.Arity1IQ(self__.Rear)
			_ = rear___1
			if Truth_(func() interface{} {
				var or__76632__auto__ = self__.Front
				_ = or__76632__auto__
				if Truth_(or__76632__auto__) {
					return or__76632__auto__
				} else {
					return rear___1
				}
			}()) {
				return (&CljsCorePersistentQueueSeq{nil, self__.Front, Seq.X_invoke_Arity1(rear___1), nil})
			} else {
				return nil
			}
		}
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreISequential__() {}
func (_ *CljsCorePersistentQueue) CljsCoreIWithMeta__()   {}
func (self__ *CljsCorePersistentQueue) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentQueue{meta___1, self__.Count, self__.Front, self__.Rear, self__.X__hash})
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreICollection__() {}
func (self__ *CljsCorePersistentQueue) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if Truth_(self__.Front) {
			return (&CljsCorePersistentQueue{self__.Meta, (self__.Count.(float64) + float64(1)), self__.Front, Conj.X_invoke_Arity2(func() interface{} {
				var or__76632__auto__ = self__.Rear
				_ = or__76632__auto__
				if Truth_(or__76632__auto__) {
					return or__76632__auto__
				} else {
					return CljsCorePersistentVector_EMPTY
				}
			}(), o).(interface{}), nil})
		} else {
			return (&CljsCorePersistentQueue{self__.Meta, (self__.Count.(float64) + float64(1)), Conj.X_invoke_Arity2(self__.Front, o).(interface{}), CljsCorePersistentVector_EMPTY, nil})
		}
	}
}

var X__GT_PersistentQueue = func(__GT_PersistentQueue *AFn) *AFn {
	return Fn(__GT_PersistentQueue, func(meta interface{}, count interface{}, front interface{}, rear interface{}, __hash interface{}) interface{} {
		return (&CljsCorePersistentQueue{meta, count, front, rear, __hash})
	})
}(&AFn{})

var CljsCorePersistentQueue_EMPTY = (&CljsCorePersistentQueue{nil, float64(0), nil, CljsCorePersistentVector_EMPTY, float64(0)})

type CljsCoreNeverEquiv struct{}

func (_ *CljsCoreNeverEquiv) CljsCoreIEquiv__() {}
func (self__ *CljsCoreNeverEquiv) X_equiv_Arity2(other interface{}) bool {
	{
		var o___1 = self__
		_ = o___1
		return false
	}
}

func (_ *CljsCoreNeverEquiv) CljsCoreObject__() {}
func (self__ *CljsCoreNeverEquiv) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

var X__GT_NeverEquiv = func(__GT_NeverEquiv *AFn) *AFn {
	return Fn(__GT_NeverEquiv, func() interface{} {
		return (&CljsCoreNeverEquiv{})
	})
}(&AFn{})

var Never_equiv = (&CljsCoreNeverEquiv{})

/**
* Assumes y is a map. Returns true if x equals y, otherwise returns
* false.
 */
var Equiv_map = func(equiv_map *AFn) *AFn {
	return Fn(equiv_map, func(x interface{}, y interface{}) interface{} {
		return Boolean.X_invoke_Arity1(func() interface{} {
			if Map_QMARK_.Arity1IB(y) {
				return func() interface{} {
					if Count.X_invoke_Arity1(x).(float64) == Count.X_invoke_Arity1(y).(float64) {
						return Every_QMARK_.X_invoke_Arity2(Identity, Map_.X_invoke_Arity2(func(G__709 *AFn) *AFn {
							return Fn(G__709, func(xkv interface{}) interface{} {
								return X_EQ_.X_invoke_Arity2(Get.X_invoke_Arity3(y, First.X_invoke_Arity1(xkv), Never_equiv), Second.X_invoke_Arity1(xkv))
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

var Scan_array = func(scan_array *AFn) *AFn {
	return Fn(scan_array, func(incr interface{}, k interface{}, array interface{}) interface{} {
		{
			var len = float64(len(array.([]interface{})))
			_ = len
			{
				var i = float64(0)
				_ = i
				for {
					if i < len {
						if k == (array.([]interface{})[int(i)]) {
							return i
						} else {
							{
								var G__710 = (i + incr.(float64))
								i = G__710
								continue
							}
						}
					} else {
						return nil
					}
				}
			}
		}
	})
}(&AFn{})

var Obj_map_compare_keys = func(obj_map_compare_keys *AFn) *AFn {
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

var Obj_map__GT_hash_map = func(obj_map__GT_hash_map *AFn) *AFn {
	return Fn(obj_map__GT_hash_map, func(m interface{}, k interface{}, v interface{}) interface{} {
		{
			var ks = Native_get_instance_field.X_invoke_Arity2(m, "Keys")
			var len = float64(len(ks.([]interface{})))
			var so = Native_get_instance_field.X_invoke_Arity2(m, "Strobj")
			var mm = Meta.X_invoke_Arity1(m)
			_, _, _, _ = ks, len, so, mm
			{
				var i = float64(0)
				var out = Transient.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(PersistentHashMap, "EMPTY")).(interface{})
				_, _ = i, out
				for {
					if i < len {
						{
							var k___1 = (ks.([]interface{})[int(i)])
							_ = k___1
							{
								var G__711 = (i + float64(1))
								var G__712 = Assoc_BANG_.X_invoke_Arity3(out, k___1, (so.([]interface{})[int(k___1.(float64))])).(interface{})
								i = G__711
								out = G__712
								continue
							}
						}
					} else {
						return With_meta.X_invoke_Arity2(Persistent_BANG_.X_invoke_Arity1(Assoc_BANG_.X_invoke_Arity3(out, k, v).(interface{})).(interface{}), mm)
					}
				}
			}
		}
	})
}(&AFn{})

var Obj_clone = func(obj_clone *AFn) *AFn {
	return Fn(obj_clone, func(obj interface{}, ks interface{}) interface{} {
		{
			var new_obj = func() interface{} {
				var obj716 = map[string]interface{}{}
				_ = obj716
				return obj716
			}()
			var l = float64(len(ks.([]interface{})))
			_, _ = new_obj, l
			{
				var i_717 = float64(0)
				_ = i_717
				for {
					if i_717 < l {
						{
							var k_718 = (ks.([]interface{})[int(i_717)])
							_ = k_718
							new_obj.([]interface{})[int(k_718.(float64))] = (obj.([]interface{})[int(k_718.(float64))])
							{
								var G__719 = (i_717 + float64(1))
								i_717 = G__719
								continue
							}
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

type CljsCoreObjMap struct {
	Meta         interface{}
	Keys         interface{}
	Strobj       interface{}
	Update_count interface{}
	X__hash      interface{}
}

func (_ *CljsCoreObjMap) CljsCoreObject__() {}
func (self__ *CljsCoreObjMap) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreObjMap) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreObjMap) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreObjMap) CljsCoreILookup__() {}
func (self__ *CljsCoreObjMap) X_lookup_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return _lookup.X_invoke_Arity3(coll___1, k, nil).(interface{})
	}
}

func (self__ *CljsCoreObjMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (goog.IsString(k)) && (!(Scan_array.X_invoke_Arity3(float64(1), k, self__.Keys) == nil)) {
			return (self__.Strobj.([]interface{})[int(k.(float64))])
		} else {
			return not_found
		}
	}
}

func (_ *CljsCoreObjMap) CljsCoreIKVReduce__() {}
func (self__ *CljsCoreObjMap) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var len = float64(len(self__.Keys.([]interface{})))
			_ = len
			{
				var keys___1 = Native_invoke_instance_method.X_invoke_Arity3(self__.Keys, "Sort", []interface{}{Obj_map_compare_keys})
				var init___1 = init
				_, _ = keys___1, init___1
				for {
					if Truth_(Seq.Arity1IQ(keys___1)) {
						{
							var k = First.X_invoke_Arity1(keys___1)
							var init___2 = f.(CljsCoreIFn).X_invoke_Arity3(init___1, k, (self__.Strobj.([]interface{})[int(k.(float64))])).(interface{})
							_, _ = k, init___2
							if Reduced_QMARK_.X_invoke_Arity1(init___2) {
								return Deref.X_invoke_Arity1(init___2).(interface{})
							} else {
								{
									var G__721 = Rest.Arity1IQ(keys___1)
									var G__722 = init___2
									keys___1 = G__721
									init___1 = G__722
									continue
								}
							}
						}
					} else {
						return init___1
					}
				}
			}
		}
	}
}

func (_ *CljsCoreObjMap) CljsCoreIMeta__() {}
func (self__ *CljsCoreObjMap) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreObjMap) CljsCoreICounted__() {}
func (self__ *CljsCoreObjMap) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return float64(len(self__.Keys.([]interface{})))
	}
}

func (_ *CljsCoreObjMap) CljsCoreIHash__() {}
func (self__ *CljsCoreObjMap) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_unordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreObjMap) CljsCoreIEquiv__() {}
func (self__ *CljsCoreObjMap) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_map.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreObjMap) CljsCoreIEditableCollection__() {}
func (self__ *CljsCoreObjMap) X_as_transient_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Transient.X_invoke_Arity1(Into.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(PersistentHashMap, "EMPTY"), coll___1)).(interface{})
	}
}

func (_ *CljsCoreObjMap) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreObjMap) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreObjMap_EMPTY, self__.Meta)
	}
}

func (_ *CljsCoreObjMap) CljsCoreIMap__() {}
func (self__ *CljsCoreObjMap) X_dissoc_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if (goog.IsString(k)) && (!(Scan_array.X_invoke_Arity3(float64(1), k, self__.Keys) == nil)) {
			{
				var new_keys = Aclone.X_invoke_Arity1(self__.Keys).([]interface{})
				var new_strobj = Obj_clone.X_invoke_Arity2(self__.Strobj, self__.Keys).(interface{})
				_, _ = new_keys, new_strobj
				new_keys.Splice(Scan_array.X_invoke_Arity3(float64(1), k, new_keys), float64(1))
				delete(new_strobj, k)
				return (&CljsCoreObjMap{self__.Meta, new_keys, new_strobj, (self__.Update_count.(float64) + float64(1)), nil})
			}
		} else {
			return coll___1
		}
	}
}

func (_ *CljsCoreObjMap) CljsCoreIAssociative__() {}
func (self__ *CljsCoreObjMap) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if goog.IsString(k) {
			if (self__.Update_count.(float64) > CljsCoreObjMap_HASHMAP_THRESHOLD.(float64)) || (float64(len(self__.Keys.([]interface{}))) >= CljsCoreObjMap_HASHMAP_THRESHOLD.(float64)) {
				return Obj_map__GT_hash_map.X_invoke_Arity3(coll___1, k, v)
			} else {
				if !(Scan_array.X_invoke_Arity3(float64(1), k, self__.Keys) == nil) {
					{
						var new_strobj = Obj_clone.X_invoke_Arity2(self__.Strobj, self__.Keys).(interface{})
						_ = new_strobj
						new_strobj.([]interface{})[int(k.(float64))] = v
						return (&CljsCoreObjMap{self__.Meta, self__.Keys, new_strobj, (self__.Update_count.(float64) + float64(1)), nil})
					}
				} else {
					{
						var new_strobj = Obj_clone.X_invoke_Arity2(self__.Strobj, self__.Keys).(interface{})
						var new_keys = Aclone.X_invoke_Arity1(self__.Keys).([]interface{})
						_, _ = new_strobj, new_keys
						new_strobj.([]interface{})[int(k.(float64))] = v
						new_keys.Push(k)
						return (&CljsCoreObjMap{self__.Meta, new_keys, new_strobj, (self__.Update_count.(float64) + float64(1)), nil})
					}
				}
			}
		} else {
			return Obj_map__GT_hash_map.X_invoke_Arity3(coll___1, k, v)
		}
	}
}

func (self__ *CljsCoreObjMap) X_contains_key_QMARK__Arity2(k interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		if (goog.IsString(k)) && (!(Scan_array.X_invoke_Arity3(float64(1), k, self__.Keys) == nil)) {
			return true
		} else {
			return false
		}
	}
}

func (_ *CljsCoreObjMap) CljsCoreISeqable__() {}
func (self__ *CljsCoreObjMap) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if float64(len(self__.Keys.([]interface{}))) > float64(0) {
			return Map_.X_invoke_Arity2(func(coll___1 *CljsCoreObjMap) *AFn {
				return func(G__723 *AFn) *AFn {
					return Fn(G__723, func(p1__720_SHARP_ interface{}) interface{} {
						return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__720_SHARP_, (self__.Strobj.([]interface{})[int(p1__720_SHARP_.(float64))])}, nil})
					})
				}(&AFn{})
			}(coll___1), Native_invoke_instance_method.X_invoke_Arity3(self__.Keys, "Sort", []interface{}{Obj_map_compare_keys})).(*CljsCoreLazySeq)
		} else {
			return nil
		}
	}
}

func (_ *CljsCoreObjMap) CljsCoreIWithMeta__() {}
func (self__ *CljsCoreObjMap) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreObjMap{meta___1, self__.Keys, self__.Strobj, self__.Update_count, self__.X__hash})
	}
}

func (_ *CljsCoreObjMap) CljsCoreICollection__() {}
func (self__ *CljsCoreObjMap) X_conj_Arity2(entry interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if Vector_QMARK_.Arity1IB(entry) {
			return coll___1.X_assoc_Arity3(entry.(CljsCoreIIndexed).X_nth_Arity2(float64(0)).(interface{}), entry.(CljsCoreIIndexed).X_nth_Arity2(float64(1)).(interface{}))
		} else {
			return Reduce.X_invoke_Arity3(_conj, coll___1, entry)
		}
	}
}

func (_ *CljsCoreObjMap) CljsCoreIFn__() {}
func (self__ *CljsCoreObjMap) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity1(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCoreObjMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

var X__GT_ObjMap = func(__GT_ObjMap *AFn) *AFn {
	return Fn(__GT_ObjMap, func(meta interface{}, keys interface{}, strobj interface{}, update_count interface{}, __hash interface{}) interface{} {
		return (&CljsCoreObjMap{meta, keys, strobj, update_count, __hash})
	})
}(&AFn{})

var CljsCoreObjMap_EMPTY = (&CljsCoreObjMap{nil, []interface{}{}, func() interface{} {
	var obj725 = map[string]interface{}{}
	_ = obj725
	return obj725
}(), float64(0), float64(0)})

var CljsCoreObjMap_HASHMAP_THRESHOLD = float64(8)

var CljsCoreObjMap_FromObject = func(G__726 *AFn) *AFn {
	return Fn(G__726, func(ks interface{}, obj interface{}) interface{} {
		return (&CljsCoreObjMap{nil, ks, obj, float64(0), nil})
	})
}(&AFn{})

type CljsCoreIterator struct{ S interface{} }

func (_ *CljsCoreIterator) CljsCoreObject__() {}
func (self__ *CljsCoreIterator) Next() CljsCoreISeq {
	{
		var ______1 = self__
		_ = ______1
		if !(self__.S == nil) {
			{
				var x = First.X_invoke_Arity1(self__.S)
				_ = x
				self__.S = next.X_invoke_Arity1(self__.S)

				return map[string]interface{}{"done": false, "value": x}.(CljsCoreISeq)
			}
		} else {
			return map[string]interface{}{"done": true, "value": nil}.(CljsCoreISeq)
		}
	}
}

var X__GT_Iterator = func(__GT_Iterator *AFn) *AFn {
	return Fn(__GT_Iterator, func(s interface{}) interface{} {
		return (&CljsCoreIterator{s})
	})
}(&AFn{})

var Iterator = func(iterator *AFn) *AFn {
	return Fn(iterator, func(coll interface{}) interface{} {
		return (&CljsCoreIterator{Seq.Arity1IQ(coll)})
	})
}(&AFn{})

type CljsCoreEntriesIterator struct{ S interface{} }

func (_ *CljsCoreEntriesIterator) CljsCoreObject__() {}
func (self__ *CljsCoreEntriesIterator) Next() CljsCoreISeq {
	{
		var ______1 = self__
		_ = ______1
		if !(self__.S == nil) {
			{
				var vec__728 = First.X_invoke_Arity1(self__.S)
				var k = Nth.X_invoke_Arity3(vec__728, float64(0), nil)
				var v = Nth.X_invoke_Arity3(vec__728, float64(1), nil)
				_, _, _ = vec__728, k, v
				self__.S = next.X_invoke_Arity1(self__.S)

				return map[string]interface{}{"done": false, "value": []interface{}{k, v}}.(CljsCoreISeq)
			}
		} else {
			return map[string]interface{}{"done": true, "value": nil}.(CljsCoreISeq)
		}
	}
}

var X__GT_EntriesIterator = func(__GT_EntriesIterator *AFn) *AFn {
	return Fn(__GT_EntriesIterator, func(s interface{}) interface{} {
		return (&CljsCoreEntriesIterator{s})
	})
}(&AFn{})

var Entries_iterator = func(entries_iterator *AFn) *AFn {
	return Fn(entries_iterator, func(coll interface{}) interface{} {
		return (&CljsCoreEntriesIterator{Seq.Arity1IQ(coll)})
	})
}(&AFn{})

type CljsCoreSetEntriesIterator struct{ S interface{} }

func (_ *CljsCoreSetEntriesIterator) CljsCoreObject__() {}
func (self__ *CljsCoreSetEntriesIterator) Next() CljsCoreISeq {
	{
		var ______1 = self__
		_ = ______1
		if !(self__.S == nil) {
			{
				var x = First.X_invoke_Arity1(self__.S)
				_ = x
				self__.S = next.X_invoke_Arity1(self__.S)

				return map[string]interface{}{"done": false, "value": []interface{}{x, x}}.(CljsCoreISeq)
			}
		} else {
			return map[string]interface{}{"done": true, "value": nil}.(CljsCoreISeq)
		}
	}
}

var X__GT_SetEntriesIterator = func(__GT_SetEntriesIterator *AFn) *AFn {
	return Fn(__GT_SetEntriesIterator, func(s interface{}) interface{} {
		return (&CljsCoreSetEntriesIterator{s})
	})
}(&AFn{})

var Set_entries_iterator = func(set_entries_iterator *AFn) *AFn {
	return Fn(set_entries_iterator, func(coll interface{}) interface{} {
		return (&CljsCoreSetEntriesIterator{Seq.Arity1IQ(coll)})
	})
}(&AFn{})

var Array_map_index_of_nil_QMARK_ = func(array_map_index_of_nil_QMARK_ *AFn) *AFn {
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
						if (arr.([]interface{})[int(i)]) == nil {
							return i
						} else {
							{
								var G__729 = (i + float64(2))
								i = G__729
								continue
							}

						}
					}
				}
			}
		}
	})
}(&AFn{})

var Array_map_index_of_keyword_QMARK_ = func(array_map_index_of_keyword_QMARK_ *AFn) *AFn {
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
						if func() interface{} {
							var k_SINGLEQUOTE_ = (arr.([]interface{})[int(i)])
							_ = k_SINGLEQUOTE_
							return (func() bool { _, instanceof := k_SINGLEQUOTE_.(*CljsCoreKeyword); return instanceof }()) && (kstr == Native_get_instance_field.X_invoke_Arity2(k_SINGLEQUOTE_, "Fqn"))
						}() {
							return i
						} else {
							{
								var G__730 = (i + float64(2))
								i = G__730
								continue
							}

						}
					}
				}
			}
		}
	})
}(&AFn{})

var Array_map_index_of_symbol_QMARK_ = func(array_map_index_of_symbol_QMARK_ *AFn) *AFn {
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
						if func() interface{} {
							var k_SINGLEQUOTE_ = (arr.([]interface{})[int(i)])
							_ = k_SINGLEQUOTE_
							return (func() bool { _, instanceof := k_SINGLEQUOTE_.(*CljsCoreSymbol); return instanceof }()) && (kstr == Native_get_instance_field.X_invoke_Arity2(k_SINGLEQUOTE_, "Str"))
						}() {
							return i
						} else {
							{
								var G__731 = (i + float64(2))
								i = G__731
								continue
							}

						}
					}
				}
			}
		}
	})
}(&AFn{})

var Array_map_index_of_identical_QMARK_ = func(array_map_index_of_identical_QMARK_ *AFn) *AFn {
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
						if k == (arr.([]interface{})[int(i)]) {
							return i
						} else {
							{
								var G__732 = (i + float64(2))
								i = G__732
								continue
							}

						}
					}
				}
			}
		}
	})
}(&AFn{})

var Array_map_index_of_equiv_QMARK_ = func(array_map_index_of_equiv_QMARK_ *AFn) *AFn {
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
							{
								var G__733 = (i + float64(2))
								i = G__733
								continue
							}

						}
					}
				}
			}
		}
	})
}(&AFn{})

var Array_map_index_of = func(array_map_index_of *AFn) *AFn {
	return Fn(array_map_index_of, func(m interface{}, k interface{}) interface{} {
		{
			var arr = Native_get_instance_field.X_invoke_Arity2(m, "Arr")
			_ = arr
			if func() bool { _, instanceof := k.(*CljsCoreKeyword); return instanceof }() {
				return Array_map_index_of_keyword_QMARK_.X_invoke_Arity3(arr, m, k).(float64)
			} else {
				if (goog.IsString(k)) || (reflect.TypeOf(k).Kind() == reflect.Float64) {
					return Array_map_index_of_identical_QMARK_.X_invoke_Arity3(arr, m, k).(float64)
				} else {
					if func() bool { _, instanceof := k.(*CljsCoreSymbol); return instanceof }() {
						return Array_map_index_of_symbol_QMARK_.X_invoke_Arity3(arr, m, k).(float64)
					} else {
						if k == nil {
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

var Array_map_extend_kv = func(array_map_extend_kv *AFn) *AFn {
	return Fn(array_map_extend_kv, func(m interface{}, k interface{}, v interface{}) interface{} {
		{
			var arr = Native_get_instance_field.X_invoke_Arity2(m, "Arr")
			var l = float64(len(arr.([]interface{})))
			var narr = make([]interface{}, int((l + float64(2))))
			_, _, _ = arr, l, narr
			{
				var i_734 = float64(0)
				_ = i_734
				for {
					if i_734 < l {
						narr[int(i_734)] = (arr.([]interface{})[int(i_734)])
						{
							var G__735 = (i_734 + float64(1))
							i_734 = G__735
							continue
						}
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

type CljsCorePersistentArrayMapSeq struct {
	Arr    interface{}
	I      interface{}
	X_meta interface{}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreObject__() {}
func (self__ *CljsCorePersistentArrayMapSeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCorePersistentArrayMapSeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCorePersistentArrayMapSeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIMeta__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.X_meta
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreINext__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.I.(float64) < (float64(len(self__.Arr.([]interface{}))) - float64(2)) {
			return (&CljsCorePersistentArrayMapSeq{self__.Arr, (self__.I.(float64) + float64(2)), self__.X_meta})
		} else {
			return nil
		}
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreICounted__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return ((float64(len(self__.Arr.([]interface{}))) - self__.I.(float64)) / float64(2))
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIHash__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Hash_ordered_coll.X_invoke_Arity1(coll___1)
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIEquiv__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.X_meta)
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIReduce__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCorePersistentArrayMapSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreISeq__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return &CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{(self__.Arr.([]interface{})[int(self__.I.(float64))]), (self__.Arr.([]interface{})[int((self__.I.(float64) + float64(1)))])}, nil}
	}
}

func (self__ *CljsCorePersistentArrayMapSeq) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.I.(float64) < (float64(len(self__.Arr.([]interface{}))) - float64(2)) {
			return (&CljsCorePersistentArrayMapSeq{self__.Arr, (self__.I.(float64) + float64(2)), self__.X_meta})
		} else {
			return CljsCoreList_EMPTY
		}
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreISeqable__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreISequential__() {}
func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIWithMeta__()   {}
func (self__ *CljsCorePersistentArrayMapSeq) X_with_meta_Arity2(new_meta interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentArrayMapSeq{self__.Arr, self__.I, new_meta})
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreICollection__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(*CljsCoreCons)
	}
}

var X__GT_PersistentArrayMapSeq = func(__GT_PersistentArrayMapSeq *AFn) *AFn {
	return Fn(__GT_PersistentArrayMapSeq, func(arr interface{}, i interface{}, _meta interface{}) interface{} {
		return (&CljsCorePersistentArrayMapSeq{arr, i, _meta})
	})
}(&AFn{})

var Persistent_array_map_seq = func(persistent_array_map_seq *AFn) *AFn {
	return Fn(persistent_array_map_seq, func(arr interface{}, i interface{}, _meta interface{}) interface{} {
		if i.(float64) <= (float64(len(arr.([]interface{}))) - float64(2)) {
			return (&CljsCorePersistentArrayMapSeq{arr, i, _meta})
		} else {
			return nil
		}
	})
}(&AFn{})

type CljsCorePersistentArrayMap struct {
	Meta    interface{}
	Cnt     interface{}
	Arr     interface{}
	X__hash interface{}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreObject__() {}
func (self__ *CljsCorePersistentArrayMap) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCorePersistentArrayMap) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCorePersistentArrayMap) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (self__ *CljsCorePersistentArrayMap) Keys() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Iterator.X_invoke_Arity1(keys.X_invoke_Arity1(coll___1).(*CljsCoreIterator)).(*CljsCoreIterator)
	}
}

func (self__ *CljsCorePersistentArrayMap) Entries() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Entries_iterator.X_invoke_Arity1(Seq.X_invoke_Arity1(coll___1)).(*CljsCoreEntriesIterator)
	}
}

func (self__ *CljsCorePersistentArrayMap) Values() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Iterator.X_invoke_Arity1(Vals.X_invoke_Arity1(coll___1).(interface{})).(*CljsCoreIterator)
	}
}

func (self__ *CljsCorePersistentArrayMap) Has(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Contains_QMARK_.X_invoke_Arity2(coll___1, k)
	}
}

func (self__ *CljsCorePersistentArrayMap) Get(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCorePersistentArrayMap) ForEach(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var seq__742 = Seq.X_invoke_Arity1(coll___1)
			var chunk__743 = nil
			var count__744 = float64(0)
			var i__745 = float64(0)
			_, _, _, _ = seq__742, chunk__743, count__744, i__745
			for {
				if i__745 < count__744 {
					{
						var vec__746 = chunk__743.X_nth_Arity2(i__745).(interface{})
						var k = Nth.X_invoke_Arity3(vec__746, float64(0), nil)
						var v = Nth.X_invoke_Arity3(vec__746, float64(1), nil)
						_, _, _ = vec__746, k, v
						f.(CljsCoreIFn).X_invoke_Arity2(v, k).(interface{})
						{
							var G__754 = seq__742
							var G__755 = chunk__743
							var G__756 = count__744
							var G__757 = (i__745 + float64(1))
							seq__742 = G__754
							chunk__743 = G__755
							count__744 = G__756
							i__745 = G__757
							continue
						}
					}
				} else {
					{
						var temp__4219__auto__ = Seq.X_invoke_Arity1(seq__742)
						_ = temp__4219__auto__
						if Truth_(temp__4219__auto__) {
							{
								var seq__742___1 = temp__4219__auto__
								_ = seq__742___1
								if Chunked_seq_QMARK_.X_invoke_Arity1(seq__742___1) {
									{
										var c__77428__auto__ = Chunk_first.X_invoke_Arity1(seq__742___1).(interface{})
										_ = c__77428__auto__
										{
											var G__758 = Chunk_rest.X_invoke_Arity1(seq__742___1).(interface{})
											var G__759 = c__77428__auto__
											var G__760 = Count.X_invoke_Arity1(c__77428__auto__).(float64)
											var G__761 = float64(0)
											seq__742 = G__758
											chunk__743 = G__759
											count__744 = G__760
											i__745 = G__761
											continue
										}
									}
								} else {
									{
										var vec__747 = First.X_invoke_Arity1(seq__742___1)
										var k = Nth.X_invoke_Arity3(vec__747, float64(0), nil)
										var v = Nth.X_invoke_Arity3(vec__747, float64(1), nil)
										_, _, _ = vec__747, k, v
										f.(CljsCoreIFn).X_invoke_Arity2(v, k).(interface{})
										{
											var G__762 = Next.X_invoke_Arity1(seq__742___1)
											var G__763 = nil
											var G__764 = float64(0)
											var G__765 = float64(0)
											seq__742 = G__762
											chunk__743 = G__763
											count__744 = G__764
											i__745 = G__765
											continue
										}
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
}

func (_ *CljsCorePersistentArrayMap) CljsCoreILookup__() {}
func (self__ *CljsCorePersistentArrayMap) X_lookup_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return _lookup.X_invoke_Arity3(coll___1, k, nil).(interface{})
	}
}

func (self__ *CljsCorePersistentArrayMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var idx = Array_map_index_of.X_invoke_Arity2(coll___1, k).(float64)
			_ = idx
			if idx == float64(-1) {
				return not_found
			} else {
				return (self__.Arr.([]interface{})[int((idx + float64(1)))])
			}
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIKVReduce__() {}
func (self__ *CljsCorePersistentArrayMap) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var len = float64(len(self__.Arr.([]interface{})))
			_ = len
			{
				var i = float64(0)
				var init___1 = init
				_, _ = i, init___1
				for {
					if i < len {
						{
							var init___2 = f.(CljsCoreIFn).X_invoke_Arity3(init___1, (self__.Arr.([]interface{})[int(i)]), (self__.Arr.([]interface{})[int((i + float64(1)))])).(interface{})
							_ = init___2
							if Reduced_QMARK_.X_invoke_Arity1(init___2) {
								return Deref.X_invoke_Arity1(init___2).(interface{})
							} else {
								{
									var G__766 = (i + float64(2))
									var G__767 = init___2
									i = G__766
									init___1 = G__767
									continue
								}
							}
						}
					} else {
						return init___1
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIMeta__() {}
func (self__ *CljsCorePersistentArrayMap) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreICloneable__() {}
func (self__ *CljsCorePersistentArrayMap) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCorePersistentArrayMap{self__.Meta, self__.Cnt, self__.Arr, self__.X__hash})
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreICounted__() {}
func (self__ *CljsCorePersistentArrayMap) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Cnt.(float64)
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIHash__() {}
func (self__ *CljsCorePersistentArrayMap) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_unordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIEquiv__() {}
func (self__ *CljsCorePersistentArrayMap) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_map.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIEditableCollection__() {}
func (self__ *CljsCorePersistentArrayMap) X_as_transient_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreTransientArrayMap{func() interface{} {
			var obj751 = map[string]interface{}{}
			_ = obj751
			return obj751
		}(), float64(len(self__.Arr.([]interface{}))), Aclone.X_invoke_Arity1(self__.Arr).([]interface{})})
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCorePersistentArrayMap) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return CljsCorePersistentArrayMap_EMPTY.(CljsCoreIWithMeta).X_with_meta_Arity2(self__.Meta)
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIReduce__() {}
func (self__ *CljsCorePersistentArrayMap) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCorePersistentArrayMap) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIMap__() {}
func (self__ *CljsCorePersistentArrayMap) X_dissoc_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var idx = Array_map_index_of.X_invoke_Arity2(coll___1, k).(float64)
			_ = idx
			if idx >= float64(0) {
				{
					var len = float64(len(self__.Arr.([]interface{})))
					var new_len = (len - float64(2))
					_, _ = len, new_len
					if new_len == float64(0) {
						return coll___1.X_empty_Arity1().(interface{})
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
										return (&CljsCorePersistentArrayMap{self__.Meta, (self__.Cnt.(float64) - float64(1)), new_arr, nil})
									} else {
										if X_EQ_.Arity2IIB(k, (self__.Arr.([]interface{})[int(s)])) {
											{
												var G__768 = (s + float64(2))
												var G__769 = d
												s = G__768
												d = G__769
												continue
											}
										} else {
											new_arr[int(d)] = (self__.Arr.([]interface{})[int(s)])
											new_arr[int((d + float64(1)))] = (self__.Arr.([]interface{})[int((s + float64(1)))])
											{
												var G__770 = (s + float64(2))
												var G__771 = (d + float64(2))
												s = G__770
												d = G__771
												continue
											}

										}
									}
								}
							}
						}
					}
				}
			} else {
				return coll___1
			}
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIAssociative__() {}
func (self__ *CljsCorePersistentArrayMap) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var idx = Array_map_index_of.X_invoke_Arity2(coll___1, k).(float64)
			_ = idx
			if idx == float64(-1) {
				if self__.Cnt.(float64) < CljsCorePersistentArrayMap_HASHMAP_THRESHOLD.(float64) {
					{
						var arr___1 = Array_map_extend_kv.X_invoke_Arity3(coll___1, k, v).([]interface{})
						_ = arr___1
						return (&CljsCorePersistentArrayMap{self__.Meta, (self__.Cnt.(float64) + float64(1)), arr___1, nil})
					}
				} else {
					return _assoc.X_invoke_Arity3(Into.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(PersistentHashMap, "EMPTY"), coll___1), k, v).X_with_meta_Arity2(self__.Meta)
				}
			} else {
				if v == (self__.Arr.([]interface{})[int((idx + float64(1)))]) {
					return coll___1
				} else {
					{
						var arr___1 = func() interface{} {
							var G__753 = Aclone.X_invoke_Arity1(self__.Arr).([]interface{})
							_ = G__753
							G__753[int((idx + float64(1)))] = v
							return G__753
						}()
						_ = arr___1
						return (&CljsCorePersistentArrayMap{self__.Meta, self__.Cnt, arr___1, nil})
					}

				}
			}
		}
	}
}

func (self__ *CljsCorePersistentArrayMap) X_contains_key_QMARK__Arity2(k interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return !(Array_map_index_of.X_invoke_Arity2(coll___1, k).(float64) == float64(-1))
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreISeqable__() {}
func (self__ *CljsCorePersistentArrayMap) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Persistent_array_map_seq.X_invoke_Arity3(self__.Arr, float64(0), nil)
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIWithMeta__() {}
func (self__ *CljsCorePersistentArrayMap) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentArrayMap{meta___1, self__.Cnt, self__.Arr, self__.X__hash})
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreICollection__() {}
func (self__ *CljsCorePersistentArrayMap) X_conj_Arity2(entry interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if Vector_QMARK_.Arity1IB(entry) {
			return coll___1.X_assoc_Arity3(entry.(CljsCoreIIndexed).X_nth_Arity2(float64(0)).(interface{}), entry.(CljsCoreIIndexed).X_nth_Arity2(float64(1)).(interface{}))
		} else {
			{
				var ret = coll___1
				var es = Seq.Arity1IQ(entry)
				_, _ = ret, es
				for {
					if es == nil {
						return ret
					} else {
						{
							var e = First.X_invoke_Arity1(es)
							_ = e
							if Vector_QMARK_.X_invoke_Arity1(e) {
								{
									var G__772 = ret.X_assoc_Arity3(e.X_nth_Arity2(float64(0)).(interface{}), e.X_nth_Arity2(float64(1)).(interface{}))
									var G__773 = Next.X_invoke_Arity1(es)
									ret = G__772
									es = G__773
									continue
								}
							} else {
								panic((&js.Error{"conj on a map takes map entries or seqables of map entries"}))
							}
						}
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIFn__() {}
func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity1(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCorePersistentArrayMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

var X__GT_PersistentArrayMap = func(__GT_PersistentArrayMap *AFn) *AFn {
	return Fn(__GT_PersistentArrayMap, func(meta interface{}, cnt interface{}, arr interface{}, __hash interface{}) interface{} {
		return (&CljsCorePersistentArrayMap{meta, cnt, arr, __hash})
	})
}(&AFn{})

var CljsCorePersistentArrayMap_EMPTY = (&CljsCorePersistentArrayMap{nil, float64(0), []interface{}{}, nil})

var CljsCorePersistentArrayMap_HASHMAP_THRESHOLD = float64(8)

var CljsCorePersistentArrayMap_FromArray = func(G__774 *AFn) *AFn {
	return Fn(G__774, func(arr interface{}, no_clone bool, no_check bool) interface{} {
		{
			var arr___1 = func() interface{} {
				if no_clone {
					return arr
				} else {
					return Aclone.X_invoke_Arity1(arr).([]interface{})
				}
			}()
			_ = arr___1
			if no_check {
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
						var ret = Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY).(interface{})
						_, _ = i, ret
						for {
							if i < len {
								{
									var G__775 = (i + float64(2))
									var G__776 = ret.X_assoc_BANG__Arity3((arr___1.([]interface{})[int(i)]), (arr___1.([]interface{})[int((i + float64(1)))]))
									i = G__775
									ret = G__776
									continue
								}
							} else {
								return ret.X_persistent_BANG__Arity1()
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
func (self__ *CljsCoreTransientArrayMap) X_dissoc_BANG__Arity2(key interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if Truth_(self__.Editable_QMARK_) {
			{
				var idx = Array_map_index_of.X_invoke_Arity2(tcoll___1, key).(float64)
				_ = idx
				if idx >= float64(0) {
					self__.Arr.([]interface{})[int(idx)] = (self__.Arr.([]interface{})[int((self__.Len.(float64) - float64(2)))])
					self__.Arr.([]interface{})[int((idx + float64(1)))] = (self__.Arr.([]interface{})[int((self__.Len.(float64) - float64(1)))])
					{
						var G__778_779 = self__.Arr
						_ = G__778_779
						Native_invoke_instance_method.X_invoke_Arity3(G__778_779, "Pop", []interface{}{})
						Native_invoke_instance_method.X_invoke_Arity3(G__778_779, "Pop", []interface{}{})
					}
					self__.Len = (self__.Len.(float64) - float64(2))

				} else {
				}
				return tcoll___1
			}
		} else {
			panic((&js.Error{"dissoc! after persistent!"}))
		}
	}
}

func (_ *CljsCoreTransientArrayMap) CljsCoreITransientAssociative__() {}
func (self__ *CljsCoreTransientArrayMap) X_assoc_BANG__Arity3(key interface{}, val interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if Truth_(self__.Editable_QMARK_) {
			{
				var idx = Array_map_index_of.X_invoke_Arity2(tcoll___1, key).(float64)
				_ = idx
				if idx == float64(-1) {
					if (self__.Len.(float64) + float64(2)) <= (float64(2) * CljsCorePersistentArrayMap_HASHMAP_THRESHOLD.(float64)) {
						self__.Len = (self__.Len.(float64) + float64(2))

						Native_invoke_instance_method.X_invoke_Arity3(self__.Arr, "Push", []interface{}{key})
						Native_invoke_instance_method.X_invoke_Arity3(self__.Arr, "Push", []interface{}{val})
						return tcoll___1
					} else {
						return Assoc_BANG_.X_invoke_Arity3(Array__GT_transient_hash_map.X_invoke_Arity2(self__.Len, self__.Arr).(interface{}), key, val).(interface{})
					}
				} else {
					if val == (self__.Arr.([]interface{})[int((idx + float64(1)))]) {
						return tcoll___1
					} else {
						self__.Arr.([]interface{})[int((idx + float64(1)))] = val
						return tcoll___1
					}
				}
			}
		} else {
			panic((&js.Error{"assoc! after persistent!"}))
		}
	}
}

func (_ *CljsCoreTransientArrayMap) CljsCoreITransientCollection__() {}
func (self__ *CljsCoreTransientArrayMap) X_conj_BANG__Arity2(o interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if Truth_(self__.Editable_QMARK_) {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IMapEntry", Str: "cljs.core/IMapEntry", X_hash: float64(535941300), X_meta: nil}), o)) {
				return tcoll___1.X_assoc_BANG__Arity3(Key.X_invoke_Arity1(o).(interface{}), Val.X_invoke_Arity1(o).(interface{}))
			} else {
				{
					var es = Seq.Arity1IQ(o)
					var tcoll___2 = tcoll___1
					_, _ = es, tcoll___2
					for {
						{
							var temp__4217__auto__ = First.X_invoke_Arity1(es)
							_ = temp__4217__auto__
							if Truth_(temp__4217__auto__) {
								{
									var e = temp__4217__auto__
									_ = e
									{
										var G__780 = Next.X_invoke_Arity1(es)
										var G__781 = tcoll___2.X_assoc_BANG__Arity3(Key.X_invoke_Arity1(e).(interface{}), Val.X_invoke_Arity1(e).(interface{}))
										es = G__780
										tcoll___2 = G__781
										continue
									}
								}
							} else {
								return tcoll___2
							}
						}
					}
				}
			}
		} else {
			panic((&js.Error{"conj! after persistent!"}))
		}
	}
}

func (self__ *CljsCoreTransientArrayMap) X_persistent_BANG__Arity1() interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if Truth_(self__.Editable_QMARK_) {
			self__.Editable_QMARK_ = false

			return (&CljsCorePersistentArrayMap{nil, Quot.X_invoke_Arity2(self__.Len, float64(2)).(float64), self__.Arr, nil})
		} else {
			panic((&js.Error{"persistent! called twice"}))
		}
	}
}

func (_ *CljsCoreTransientArrayMap) CljsCoreILookup__() {}
func (self__ *CljsCoreTransientArrayMap) X_lookup_Arity2(k interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		return _lookup.X_invoke_Arity3(tcoll___1, k, nil).(interface{})
	}
}

func (self__ *CljsCoreTransientArrayMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if Truth_(self__.Editable_QMARK_) {
			{
				var idx = Array_map_index_of.X_invoke_Arity2(tcoll___1, k).(float64)
				_ = idx
				if idx == float64(-1) {
					return not_found
				} else {
					return (self__.Arr.([]interface{})[int((idx + float64(1)))])
				}
			}
		} else {
			panic((&js.Error{"lookup after persistent!"}))
		}
	}
}

func (_ *CljsCoreTransientArrayMap) CljsCoreICounted__() {}
func (self__ *CljsCoreTransientArrayMap) X_count_Arity1() float64 {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if Truth_(self__.Editable_QMARK_) {
			return Quot.X_invoke_Arity2(self__.Len, float64(2)).(float64)
		} else {
			panic((&js.Error{"count after persistent!"}))
		}
	}
}

var X__GT_TransientArrayMap = func(__GT_TransientArrayMap *AFn) *AFn {
	return Fn(__GT_TransientArrayMap, func(editable_QMARK_ interface{}, len interface{}, arr interface{}) interface{} {
		return (&CljsCoreTransientArrayMap{editable_QMARK_, len, arr})
	})
}(&AFn{})

var Array__GT_transient_hash_map = func(array__GT_transient_hash_map *AFn) *AFn {
	return Fn(array__GT_transient_hash_map, func(len interface{}, arr interface{}) interface{} {
		{
			var out = Transient.X_invoke_Arity1(Native_get_instance_field.X_invoke_Arity2(PersistentHashMap, "EMPTY")).(interface{})
			var i = float64(0)
			_, _ = out, i
			for {
				if i < len.(float64) {
					{
						var G__782 = Assoc_BANG_.X_invoke_Arity3(out, (arr.([]interface{})[int(i)]), (arr.([]interface{})[int((i + float64(1)))])).(interface{})
						var G__783 = (i + float64(2))
						out = G__782
						i = G__783
						continue
					}
				} else {
					return out
				}
			}
		}
	})
}(&AFn{})

type CljsCoreBox struct{ Val interface{} }

var X__GT_Box = func(__GT_Box *AFn) *AFn {
	return Fn(__GT_Box, func(val interface{}) interface{} {
		return (&CljsCoreBox{val})
	})
}(&AFn{})

var Key_test = func(key_test *AFn) *AFn {
	return Fn(key_test, func(key interface{}, other interface{}) bool {
		if key == other {
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

var Mask = func(mask *AFn) *AFn {
	return Fn(mask, func(hash interface{}, shift interface{}) interface{} {
		return float64(int(float64(uint(hash.(float64))>>uint(shift.(float64)))) & int(float64(31)))
	})
}(&AFn{})

var Clone_and_set = func(clone_and_set *AFn) *AFn {
	return Fn(clone_and_set, func(arr interface{}, i interface{}, a interface{}) interface{} {
		{
			var G__786 = Aclone.X_invoke_Arity1(arr).([]interface{})
			_ = G__786
			G__786[int(i.(float64))] = a
			return G__786
		}
	}, func(arr interface{}, i interface{}, a interface{}, j interface{}, b interface{}) interface{} {
		{
			var G__787 = Aclone.X_invoke_Arity1(arr).([]interface{})
			_ = G__787
			G__787[int(i.(float64))] = a
			G__787[int(j.(float64))] = b
			return G__787
		}
	})
}(&AFn{})

var Remove_pair = func(remove_pair *AFn) *AFn {
	return Fn(remove_pair, func(arr interface{}, i interface{}) interface{} {
		{
			var new_arr = make([]interface{}, int((float64(len(arr.([]interface{}))) - float64(2))))
			_ = new_arr
			Array_copy.X_invoke_Arity5(arr, float64(0), new_arr, float64(0), (float64(2) * i.(float64))).(interface{})
			Array_copy.X_invoke_Arity5(arr, (float64(2) * (i.(float64) + float64(1))), new_arr, (float64(2) * i.(float64)), (float64(len(new_arr)) - (float64(2) * i.(float64)))).(interface{})
			return new_arr
		}
	})
}(&AFn{})

var Bitmap_indexed_node_index = func(bitmap_indexed_node_index *AFn) *AFn {
	return Fn(bitmap_indexed_node_index, func(bitmap interface{}, bit interface{}) interface{} {
		return Bit_count.X_invoke_Arity1(float64(int(bitmap.(float64)) & int((bit.(float64) - float64(1))))).(float64)
	})
}(&AFn{})

var Bitpos = func(bitpos *AFn) *AFn {
	return Fn(bitpos, func(hash interface{}, shift interface{}) interface{} {
		return float64(int(float64(1)) << uint(float64((uint(hash)>>uint(shift))&0x01f).(float64)))
	})
}(&AFn{})

var Edit_and_set = func(edit_and_set *AFn) *AFn {
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

var Inode_kv_reduce = func(inode_kv_reduce *AFn) *AFn {
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
								if !(k == nil) {
									return f.(CljsCoreIFn).X_invoke_Arity3(init___1, k, (arr.([]interface{})[int((i + float64(1)))])).(interface{})
								} else {
									{
										var node = (arr.([]interface{})[int((i + float64(1)))])
										_ = node
										if !(node == nil) {
											return Native_invoke_instance_method.X_invoke_Arity3(node, "Kv_reduce", []interface{}{f, init___1})
										} else {
											return init___1
										}
									}
								}
							}()
							_ = init___2
							if Reduced_QMARK_.X_invoke_Arity1(init___2) {
								return Deref.X_invoke_Arity1(init___2).(interface{})
							} else {
								{
									var G__788 = (i + float64(2))
									var G__789 = init___2
									i = G__788
									init___1 = G__789
									continue
								}
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

type CljsCoreBitmapIndexedNode struct {
	Edit   interface{}
	Bitmap interface{}
	Arr    interface{}
}

func (_ *CljsCoreBitmapIndexedNode) CljsCoreObject__() {}
func (self__ *CljsCoreBitmapIndexedNode) Ensure_editable(e interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		if e == self__.Edit {
			return inode___1
		} else {
			{
				var n = Bit_count.X_invoke_Arity1(self__.Bitmap).(float64)
				var new_arr = make([]interface{}, int(func() interface{} {
					if n < float64(0) {
						return float64(4)
					} else {
						return (float64(2) * (n + float64(1)))
					}
				}()))
				_, _ = n, new_arr
				Array_copy.X_invoke_Arity5(self__.Arr, float64(0), new_arr, float64(0), (float64(2) * n)).(interface{})
				return (&CljsCoreBitmapIndexedNode{e, self__.Bitmap, new_arr})
			}
		}
	}
}

func (self__ *CljsCoreBitmapIndexedNode) Inode_without_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, removed_leaf_QMARK_ interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var bit = float64(1 << uint(float64((uint(hash)>>uint(shift))&0x01f)))
			_ = bit
			if float64(int(self__.Bitmap.(float64))&int(bit.(float64))) == float64(0) {
				return inode___1
			} else {
				{
					var idx = Bitmap_indexed_node_index.X_invoke_Arity2(self__.Bitmap, bit).(float64)
					var key_or_nil = (self__.Arr.([]interface{})[int((float64(2) * idx))])
					var val_or_node = (self__.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
					_, _, _ = idx, key_or_nil, val_or_node
					if key_or_nil == nil {
						{
							var n = Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_without_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, removed_leaf_QMARK_})
							_ = n
							if n == val_or_node {
								return inode___1
							} else {
								if !(n == nil) {
									return Edit_and_set.X_invoke_Arity4(inode___1, edit___1, ((float64(2) * idx) + float64(1)), n).(interface{})
								} else {
									if self__.Bitmap.(float64) == bit.(float64) {
										return nil
									} else {
										return inode___1.Edit_and_remove_pair(edit___1, bit, idx)

									}
								}
							}
						}
					} else {
						if Key_test.Arity2IIB(key, key_or_nil) {
							removed_leaf_QMARK_.([]interface{})[int(float64(0))] = true
							return inode___1.Edit_and_remove_pair(edit___1, bit, idx)
						} else {
							return inode___1

						}
					}
				}
			}
		}
	}
}

func (self__ *CljsCoreBitmapIndexedNode) Edit_and_remove_pair(e interface{}, bit interface{}, i interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		if self__.Bitmap.(float64) == bit.(float64) {
			return nil
		} else {
			{
				var editable = inode___1.Ensure_editable(e)
				var earr = Native_get_instance_field.X_invoke_Arity2(editable, "Arr")
				var len = float64(len(earr.([]interface{})))
				_, _, _ = editable, earr, len
				Native_set_instance_field.X_invoke_Arity3(editable, "Bitmap", float64(int(bit.(float64))^int(Native_get_instance_field.X_invoke_Arity2(editable, "Bitmap").(float64))))

				Array_copy.X_invoke_Arity5(earr, (float64(2) * (i.(float64) + float64(1))), earr, (float64(2) * i.(float64)), (len - (float64(2) * (i.(float64) + float64(1))))).(interface{})
				earr.([]interface{})[int((len - float64(2)))] = nil
				earr.([]interface{})[int((len - float64(1)))] = nil
				return editable
			}
		}
	}
}

func (self__ *CljsCoreBitmapIndexedNode) Inode_seq() interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		return Create_inode_seq.X_invoke_Arity1(self__.Arr).(interface{})
	}
}

func (self__ *CljsCoreBitmapIndexedNode) Kv_reduce(f interface{}, init interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		return Inode_kv_reduce.X_invoke_Arity3(self__.Arr, f, init)
	}
}

func (self__ *CljsCoreBitmapIndexedNode) Inode_lookup(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var bit = float64(1 << uint(float64((uint(hash)>>uint(shift))&0x01f)))
			_ = bit
			if float64(int(self__.Bitmap.(float64))&int(bit.(float64))) == float64(0) {
				return not_found
			} else {
				{
					var idx = Bitmap_indexed_node_index.X_invoke_Arity2(self__.Bitmap, bit).(float64)
					var key_or_nil = (self__.Arr.([]interface{})[int((float64(2) * idx))])
					var val_or_node = (self__.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
					_, _, _ = idx, key_or_nil, val_or_node
					if key_or_nil == nil {
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
}

func (self__ *CljsCoreBitmapIndexedNode) Inode_assoc_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var bit = float64(1 << uint(float64((uint(hash)>>uint(shift))&0x01f)))
			var idx = Bitmap_indexed_node_index.X_invoke_Arity2(self__.Bitmap, bit).(float64)
			_, _ = bit, idx
			if float64(int(self__.Bitmap.(float64))&int(bit.(float64))) == float64(0) {
				{
					var n = Bit_count.X_invoke_Arity1(self__.Bitmap).(float64)
					_ = n
					if (float64(2) * n) < float64(len(self__.Arr.([]interface{}))) {
						{
							var editable = inode___1.Ensure_editable(edit___1)
							var earr = Native_get_instance_field.X_invoke_Arity2(editable, "Arr")
							_, _ = editable, earr
							Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)

							Array_copy_downward.X_invoke_Arity5(earr, (float64(2) * idx), earr, (float64(2) * (idx + float64(1))), (float64(2) * (n - idx))).(interface{})
							earr.([]interface{})[int((float64(2) * idx))] = key
							earr.([]interface{})[int(((float64(2) * idx) + float64(1)))] = val
							Native_set_instance_field.X_invoke_Arity3(editable, "Bitmap", float64(int(Native_get_instance_field.X_invoke_Arity2(editable, "Bitmap").(float64))|int(bit.(float64))))

							return editable
						}
					} else {
						if n >= float64(16) {
							{
								var nodes = make([]interface{}, int(float64(32)))
								var jdx = float64((uint(hash) >> uint(shift)) & 0x01f)
								_, _ = nodes, jdx
								nodes[int(jdx.(float64))] = Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
								{
									var i_790 = float64(0)
									var j_791 = float64(0)
									_, _ = i_790, j_791
									for {
										if i_790 < float64(32) {
											if float64(int(float64(uint(self__.Bitmap.(float64))>>uint(i_790)))&int(float64(1))) == float64(0) {
												{
													var G__792 = (i_790 + float64(1))
													var G__793 = j_791
													i_790 = G__792
													j_791 = G__793
													continue
												}
											} else {
												nodes[int(i_790)] = func() interface{} {
													if !((self__.Arr.([]interface{})[int(j_791)]) == nil) {
														return Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), Hash.X_invoke_Arity1((self__.Arr.([]interface{})[int(j_791)])), (self__.Arr.([]interface{})[int(j_791)]), (self__.Arr.([]interface{})[int((j_791 + float64(1)))]), added_leaf_QMARK_})
													} else {
														return (self__.Arr.([]interface{})[int((j_791 + float64(1)))])
													}
												}()
												{
													var G__794 = (i_790 + float64(1))
													var G__795 = (j_791 + float64(2))
													i_790 = G__794
													j_791 = G__795
													continue
												}
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
								Array_copy.X_invoke_Arity5(self__.Arr, float64(0), new_arr, float64(0), (float64(2) * idx)).(interface{})
								new_arr[int((float64(2) * idx))] = key
								new_arr[int(((float64(2) * idx) + float64(1)))] = val
								Array_copy.X_invoke_Arity5(self__.Arr, (float64(2) * idx), new_arr, (float64(2) * (idx + float64(1))), (float64(2) * (n - idx))).(interface{})
								Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)

								{
									var editable = inode___1.Ensure_editable(edit___1)
									_ = editable
									Native_set_instance_field.X_invoke_Arity3(editable, "Arr", new_arr)

									Native_set_instance_field.X_invoke_Arity3(editable, "Bitmap", float64(int(Native_get_instance_field.X_invoke_Arity2(editable, "Bitmap").(float64))|int(bit.(float64))))

									return editable
								}
							}

						}
					}
				}
			} else {
				{
					var key_or_nil = (self__.Arr.([]interface{})[int((float64(2) * idx))])
					var val_or_node = (self__.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
					_, _ = key_or_nil, val_or_node
					if key_or_nil == nil {
						{
							var n = Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
							_ = n
							if n == val_or_node {
								return inode___1
							} else {
								return Edit_and_set.X_invoke_Arity4(inode___1, edit___1, ((float64(2) * idx) + float64(1)), n).(interface{})
							}
						}
					} else {
						if Key_test.Arity2IIB(key, key_or_nil) {
							if val == val_or_node {
								return inode___1
							} else {
								return Edit_and_set.X_invoke_Arity4(inode___1, edit___1, ((float64(2) * idx) + float64(1)), val).(interface{})
							}
						} else {
							Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)

							return Edit_and_set.X_invoke_Arity6(inode___1, edit___1, (float64(2) * idx), nil, ((float64(2) * idx) + float64(1)), Create_node.X_invoke_Arity7(edit___1, (shift.(float64)+float64(5)), key_or_nil, val_or_node, hash, key, val).(interface{})).(interface{})

						}
					}
				}
			}
		}
	}
}

func (self__ *CljsCoreBitmapIndexedNode) Inode_assoc(shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var bit = float64(1 << uint(float64((uint(hash)>>uint(shift))&0x01f)))
			var idx = Bitmap_indexed_node_index.X_invoke_Arity2(self__.Bitmap, bit).(float64)
			_, _ = bit, idx
			if float64(int(self__.Bitmap.(float64))&int(bit.(float64))) == float64(0) {
				{
					var n = Bit_count.X_invoke_Arity1(self__.Bitmap).(float64)
					_ = n
					if n >= float64(16) {
						{
							var nodes = make([]interface{}, int(float64(32)))
							var jdx = float64((uint(hash) >> uint(shift)) & 0x01f)
							_, _ = nodes, jdx
							nodes[int(jdx.(float64))] = Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
							{
								var i_796 = float64(0)
								var j_797 = float64(0)
								_, _ = i_796, j_797
								for {
									if i_796 < float64(32) {
										if float64(int(float64(uint(self__.Bitmap.(float64))>>uint(i_796)))&int(float64(1))) == float64(0) {
											{
												var G__798 = (i_796 + float64(1))
												var G__799 = j_797
												i_796 = G__798
												j_797 = G__799
												continue
											}
										} else {
											nodes[int(i_796)] = func() interface{} {
												if !((self__.Arr.([]interface{})[int(j_797)]) == nil) {
													return Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), Hash.X_invoke_Arity1((self__.Arr.([]interface{})[int(j_797)])), (self__.Arr.([]interface{})[int(j_797)]), (self__.Arr.([]interface{})[int((j_797 + float64(1)))]), added_leaf_QMARK_})
												} else {
													return (self__.Arr.([]interface{})[int((j_797 + float64(1)))])
												}
											}()
											{
												var G__800 = (i_796 + float64(1))
												var G__801 = (j_797 + float64(2))
												i_796 = G__800
												j_797 = G__801
												continue
											}
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
							Array_copy.X_invoke_Arity5(self__.Arr, float64(0), new_arr, float64(0), (float64(2) * idx)).(interface{})
							new_arr[int((float64(2) * idx))] = key
							new_arr[int(((float64(2) * idx) + float64(1)))] = val
							Array_copy.X_invoke_Arity5(self__.Arr, (float64(2) * idx), new_arr, (float64(2) * (idx + float64(1))), (float64(2) * (n - idx))).(interface{})
							Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)

							return (&CljsCoreBitmapIndexedNode{nil, float64(int(self__.Bitmap.(float64)) | int(bit.(float64))), new_arr})
						}
					}
				}
			} else {
				{
					var key_or_nil = (self__.Arr.([]interface{})[int((float64(2) * idx))])
					var val_or_node = (self__.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
					_, _ = key_or_nil, val_or_node
					if key_or_nil == nil {
						{
							var n = Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
							_ = n
							if n == val_or_node {
								return inode___1
							} else {
								return (&CljsCoreBitmapIndexedNode{nil, self__.Bitmap, Clone_and_set.X_invoke_Arity3(self__.Arr, ((float64(2) * idx) + float64(1)), n).([]interface{})})
							}
						}
					} else {
						if Key_test.Arity2IIB(key, key_or_nil) {
							if val == val_or_node {
								return inode___1
							} else {
								return (&CljsCoreBitmapIndexedNode{nil, self__.Bitmap, Clone_and_set.X_invoke_Arity3(self__.Arr, ((float64(2) * idx) + float64(1)), val).([]interface{})})
							}
						} else {
							Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)

							return (&CljsCoreBitmapIndexedNode{nil, self__.Bitmap, Clone_and_set.X_invoke_Arity5(self__.Arr, (float64(2) * idx), nil, ((float64(2) * idx) + float64(1)), Create_node.X_invoke_Arity6((shift.(float64)+float64(5)), key_or_nil, val_or_node, hash, key, val).(interface{})).([]interface{})})

						}
					}
				}
			}
		}
	}
}

func (self__ *CljsCoreBitmapIndexedNode) Inode_find(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var bit = float64(1 << uint(float64((uint(hash)>>uint(shift))&0x01f)))
			_ = bit
			if float64(int(self__.Bitmap.(float64))&int(bit.(float64))) == float64(0) {
				return not_found
			} else {
				{
					var idx = Bitmap_indexed_node_index.X_invoke_Arity2(self__.Bitmap, bit).(float64)
					var key_or_nil = (self__.Arr.([]interface{})[int((float64(2) * idx))])
					var val_or_node = (self__.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
					_, _, _ = idx, key_or_nil, val_or_node
					if key_or_nil == nil {
						return Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_find", []interface{}{(shift.(float64) + float64(5)), hash, key, not_found})
					} else {
						if Key_test.Arity2IIB(key, key_or_nil) {
							return &CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{key_or_nil, val_or_node}, nil}
						} else {
							return not_found

						}
					}
				}
			}
		}
	}
}

func (self__ *CljsCoreBitmapIndexedNode) Inode_without(shift interface{}, hash interface{}, key interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var bit = float64(1 << uint(float64((uint(hash)>>uint(shift))&0x01f)))
			_ = bit
			if float64(int(self__.Bitmap.(float64))&int(bit.(float64))) == float64(0) {
				return inode___1
			} else {
				{
					var idx = Bitmap_indexed_node_index.X_invoke_Arity2(self__.Bitmap, bit).(float64)
					var key_or_nil = (self__.Arr.([]interface{})[int((float64(2) * idx))])
					var val_or_node = (self__.Arr.([]interface{})[int(((float64(2) * idx) + float64(1)))])
					_, _, _ = idx, key_or_nil, val_or_node
					if key_or_nil == nil {
						{
							var n = Native_invoke_instance_method.X_invoke_Arity3(val_or_node, "Inode_without", []interface{}{(shift.(float64) + float64(5)), hash, key})
							_ = n
							if n == val_or_node {
								return inode___1
							} else {
								if !(n == nil) {
									return (&CljsCoreBitmapIndexedNode{nil, self__.Bitmap, Clone_and_set.X_invoke_Arity3(self__.Arr, ((float64(2) * idx) + float64(1)), n).([]interface{})})
								} else {
									if self__.Bitmap.(float64) == bit.(float64) {
										return nil
									} else {
										return (&CljsCoreBitmapIndexedNode{nil, float64(int(self__.Bitmap.(float64)) ^ int(bit.(float64))), Remove_pair.X_invoke_Arity2(self__.Arr, idx).([]interface{})})

									}
								}
							}
						}
					} else {
						if Key_test.Arity2IIB(key, key_or_nil) {
							return (&CljsCoreBitmapIndexedNode{nil, float64(int(self__.Bitmap.(float64)) ^ int(bit.(float64))), Remove_pair.X_invoke_Arity2(self__.Arr, idx).([]interface{})})
						} else {
							return inode___1

						}
					}
				}
			}
		}
	}
}

var X__GT_BitmapIndexedNode = func(__GT_BitmapIndexedNode *AFn) *AFn {
	return Fn(__GT_BitmapIndexedNode, func(edit interface{}, bitmap interface{}, arr interface{}) interface{} {
		return (&CljsCoreBitmapIndexedNode{edit, bitmap, arr})
	})
}(&AFn{})

var CljsCoreBitmapIndexedNode_EMPTY = (&CljsCoreBitmapIndexedNode{nil, float64(0), make([]interface{}, int(float64(0)))})

var Pack_array_node = func(pack_array_node *AFn) *AFn {
	return Fn(pack_array_node, func(array_node interface{}, edit interface{}, idx interface{}) interface{} {
		{
			var arr = Native_get_instance_field.X_invoke_Arity2(array_node, "Arr")
			var len = (float64(2) * (Native_get_instance_field.X_invoke_Arity2(array_node, "Cnt").(float64) - float64(1)))
			var new_arr = make([]interface{}, int(len))
			_, _, _ = arr, len, new_arr
			{
				var i = float64(0)
				var j = float64(1)
				var bitmap = float64(0)
				_, _, _ = i, j, bitmap
				for {
					if i < len {
						if (!(i == idx.(float64))) && (!((arr.([]interface{})[int(i)]) == nil)) {
							new_arr[int(j)] = (arr.([]interface{})[int(i)])
							{
								var G__802 = (i + float64(1))
								var G__803 = (j + float64(2))
								var G__804 = float64(int(bitmap) | int(float64(int(float64(1))<<uint(i))))
								i = G__802
								j = G__803
								bitmap = G__804
								continue
							}
						} else {
							{
								var G__805 = (i + float64(1))
								var G__806 = j
								var G__807 = bitmap
								i = G__805
								j = G__806
								bitmap = G__807
								continue
							}
						}
					} else {
						return (&CljsCoreBitmapIndexedNode{edit, bitmap, new_arr})
					}
				}
			}
		}
	})
}(&AFn{})

type CljsCoreArrayNode struct {
	Edit interface{}
	Cnt  interface{}
	Arr  interface{}
}

func (_ *CljsCoreArrayNode) CljsCoreObject__() {}
func (self__ *CljsCoreArrayNode) Ensure_editable(e interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		if e == self__.Edit {
			return inode___1
		} else {
			return (&CljsCoreArrayNode{e, self__.Cnt, Aclone.X_invoke_Arity1(self__.Arr).([]interface{})})
		}
	}
}

func (self__ *CljsCoreArrayNode) Inode_without_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, removed_leaf_QMARK_ interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var idx = float64((uint(hash) >> uint(shift)) & 0x01f)
			var node = (self__.Arr.([]interface{})[int(idx.(float64))])
			_, _ = idx, node
			if node == nil {
				return inode___1
			} else {
				{
					var n = Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_without_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, removed_leaf_QMARK_})
					_ = n
					if n == node {
						return inode___1
					} else {
						if n == nil {
							if self__.Cnt.(float64) <= float64(8) {
								return Pack_array_node.X_invoke_Arity3(inode___1, edit___1, idx).(*CljsCoreBitmapIndexedNode)
							} else {
								{
									var editable = Edit_and_set.X_invoke_Arity4(inode___1, edit___1, idx, n).(interface{})
									_ = editable
									Native_set_instance_field.X_invoke_Arity3(editable, "Cnt", (editable.Cnt.(float64) - float64(1)))

									return editable
								}
							}
						} else {
							return Edit_and_set.X_invoke_Arity4(inode___1, edit___1, idx, n).(interface{})

						}
					}
				}
			}
		}
	}
}

func (self__ *CljsCoreArrayNode) Inode_seq() interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		return Create_array_node_seq.X_invoke_Arity1(self__.Arr).(interface{})
	}
}

func (self__ *CljsCoreArrayNode) Kv_reduce(f interface{}, init interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var len = float64(len(self__.Arr.([]interface{})))
			_ = len
			{
				var i = float64(0)
				var init___1 = init
				_, _ = i, init___1
				for {
					if i < len {
						{
							var node = (self__.Arr.([]interface{})[int(i)])
							_ = node
							if !(node == nil) {
								{
									var init___2 = Native_invoke_instance_method.X_invoke_Arity3(node, "Kv_reduce", []interface{}{f, init___1})
									_ = init___2
									if Reduced_QMARK_.Arity1IB(init___2) {
										return Deref.X_invoke_Arity1(init___2).(interface{})
									} else {
										{
											var G__808 = (i + float64(1))
											var G__809 = init___2
											i = G__808
											init___1 = G__809
											continue
										}
									}
								}
							} else {
								{
									var G__810 = (i + float64(1))
									var G__811 = init___1
									i = G__810
									init___1 = G__811
									continue
								}
							}
						}
					} else {
						return init___1
					}
				}
			}
		}
	}
}

func (self__ *CljsCoreArrayNode) Inode_lookup(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var idx = float64((uint(hash) >> uint(shift)) & 0x01f)
			var node = (self__.Arr.([]interface{})[int(idx.(float64))])
			_, _ = idx, node
			if !(node == nil) {
				return Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_lookup", []interface{}{(shift.(float64) + float64(5)), hash, key, not_found})
			} else {
				return not_found
			}
		}
	}
}

func (self__ *CljsCoreArrayNode) Inode_assoc_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var idx = float64((uint(hash) >> uint(shift)) & 0x01f)
			var node = (self__.Arr.([]interface{})[int(idx.(float64))])
			_, _ = idx, node
			if node == nil {
				{
					var editable = Edit_and_set.X_invoke_Arity4(inode___1, edit___1, idx, Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})).(interface{})
					_ = editable
					Native_set_instance_field.X_invoke_Arity3(editable, "Cnt", (editable.Cnt.(float64) + float64(1)))

					return editable
				}
			} else {
				{
					var n = Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_assoc_BANG_", []interface{}{edit___1, (shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
					_ = n
					if n == node {
						return inode___1
					} else {
						return Edit_and_set.X_invoke_Arity4(inode___1, edit___1, idx, n).(interface{})
					}
				}
			}
		}
	}
}

func (self__ *CljsCoreArrayNode) Inode_assoc(shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var idx = float64((uint(hash) >> uint(shift)) & 0x01f)
			var node = (self__.Arr.([]interface{})[int(idx.(float64))])
			_, _ = idx, node
			if node == nil {
				return (&CljsCoreArrayNode{nil, (self__.Cnt.(float64) + float64(1)), Clone_and_set.X_invoke_Arity3(self__.Arr, idx, Native_invoke_instance_method.X_invoke_Arity3(CljsCoreBitmapIndexedNode_EMPTY, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})).([]interface{})})
			} else {
				{
					var n = Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_assoc", []interface{}{(shift.(float64) + float64(5)), hash, key, val, added_leaf_QMARK_})
					_ = n
					if n == node {
						return inode___1
					} else {
						return (&CljsCoreArrayNode{nil, self__.Cnt, Clone_and_set.X_invoke_Arity3(self__.Arr, idx, n).([]interface{})})
					}
				}
			}
		}
	}
}

func (self__ *CljsCoreArrayNode) Inode_find(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var idx = float64((uint(hash) >> uint(shift)) & 0x01f)
			var node = (self__.Arr.([]interface{})[int(idx.(float64))])
			_, _ = idx, node
			if !(node == nil) {
				return Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_find", []interface{}{(shift.(float64) + float64(5)), hash, key, not_found})
			} else {
				return not_found
			}
		}
	}
}

func (self__ *CljsCoreArrayNode) Inode_without(shift interface{}, hash interface{}, key interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var idx = float64((uint(hash) >> uint(shift)) & 0x01f)
			var node = (self__.Arr.([]interface{})[int(idx.(float64))])
			_, _ = idx, node
			if !(node == nil) {
				{
					var n = Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_without", []interface{}{(shift.(float64) + float64(5)), hash, key})
					_ = n
					if n == node {
						return inode___1
					} else {
						if n == nil {
							if self__.Cnt.(float64) <= float64(8) {
								return Pack_array_node.X_invoke_Arity3(inode___1, nil, idx).(*CljsCoreBitmapIndexedNode)
							} else {
								return (&CljsCoreArrayNode{nil, (self__.Cnt.(float64) - float64(1)), Clone_and_set.X_invoke_Arity3(self__.Arr, idx, n).([]interface{})})
							}
						} else {
							return (&CljsCoreArrayNode{nil, self__.Cnt, Clone_and_set.X_invoke_Arity3(self__.Arr, idx, n).([]interface{})})

						}
					}
				}
			} else {
				return inode___1
			}
		}
	}
}

var X__GT_ArrayNode = func(__GT_ArrayNode *AFn) *AFn {
	return Fn(__GT_ArrayNode, func(edit interface{}, cnt interface{}, arr interface{}) interface{} {
		return (&CljsCoreArrayNode{edit, cnt, arr})
	})
}(&AFn{})

var Hash_collision_node_find_index = func(hash_collision_node_find_index *AFn) *AFn {
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
							{
								var G__812 = (i + float64(2))
								i = G__812
								continue
							}
						}
					} else {
						return float64(-1)
					}
				}
			}
		}
	})
}(&AFn{})

type CljsCoreHashCollisionNode struct {
	Edit           interface{}
	Collision_hash interface{}
	Cnt            interface{}
	Arr            interface{}
}

func (_ *CljsCoreHashCollisionNode) CljsCoreObject__() {}
func (self__ *CljsCoreHashCollisionNode) Ensure_editable(e interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		if e == self__.Edit {
			return inode___1
		} else {
			{
				var new_arr = make([]interface{}, int((float64(2) * (self__.Cnt.(float64) + float64(1)))))
				_ = new_arr
				Array_copy.X_invoke_Arity5(self__.Arr, float64(0), new_arr, float64(0), (float64(2) * self__.Cnt.(float64))).(interface{})
				return (&CljsCoreHashCollisionNode{e, self__.Collision_hash, self__.Cnt, new_arr})
			}
		}
	}
}

func (self__ *CljsCoreHashCollisionNode) Inode_without_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, removed_leaf_QMARK_ interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var idx = Hash_collision_node_find_index.X_invoke_Arity3(self__.Arr, self__.Cnt, key).(float64)
			_ = idx
			if idx == float64(-1) {
				return inode___1
			} else {
				removed_leaf_QMARK_.([]interface{})[int(float64(0))] = true
				if self__.Cnt.(float64) == float64(1) {
					return nil
				} else {
					{
						var editable = inode___1.Ensure_editable(edit___1)
						var earr = Native_get_instance_field.X_invoke_Arity2(editable, "Arr")
						_, _ = editable, earr
						earr.([]interface{})[int(idx)] = (earr.([]interface{})[int(((float64(2) * self__.Cnt.(float64)) - float64(2)))])
						earr.([]interface{})[int((idx + float64(1)))] = (earr.([]interface{})[int(((float64(2) * self__.Cnt.(float64)) - float64(1)))])
						earr.([]interface{})[int(((float64(2) * self__.Cnt.(float64)) - float64(1)))] = nil
						earr.([]interface{})[int(((float64(2) * self__.Cnt.(float64)) - float64(2)))] = nil
						Native_set_instance_field.X_invoke_Arity3(editable, "Cnt", (Native_get_instance_field.X_invoke_Arity2(editable, "Cnt").(float64) - float64(1)))

						return editable
					}
				}
			}
		}
	}
}

func (self__ *CljsCoreHashCollisionNode) Inode_seq() interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		return Create_inode_seq.X_invoke_Arity1(self__.Arr).(interface{})
	}
}

func (self__ *CljsCoreHashCollisionNode) Kv_reduce(f interface{}, init interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		return Inode_kv_reduce.X_invoke_Arity3(self__.Arr, f, init)
	}
}

func (self__ *CljsCoreHashCollisionNode) Inode_lookup(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var idx = Hash_collision_node_find_index.X_invoke_Arity3(self__.Arr, self__.Cnt, key).(float64)
			_ = idx
			if idx < float64(0) {
				return not_found
			} else {
				if Key_test.Arity2IIB(key, (self__.Arr.([]interface{})[int(idx)])) {
					return (self__.Arr.([]interface{})[int((idx + float64(1)))])
				} else {
					return not_found

				}
			}
		}
	}
}

func (self__ *CljsCoreHashCollisionNode) Inode_assoc_BANG_(edit___1 interface{}, shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		if hash.(float64) == self__.Collision_hash.(float64) {
			{
				var idx = Hash_collision_node_find_index.X_invoke_Arity3(self__.Arr, self__.Cnt, key).(float64)
				_ = idx
				if idx == float64(-1) {
					if float64(len(self__.Arr.([]interface{}))) > (float64(2) * self__.Cnt.(float64)) {
						{
							var editable = Edit_and_set.X_invoke_Arity6(inode___1, edit___1, (float64(2) * self__.Cnt.(float64)), key, ((float64(2) * self__.Cnt.(float64)) + float64(1)), val).(interface{})
							_ = editable
							Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)

							Native_set_instance_field.X_invoke_Arity3(editable, "Cnt", (editable.Cnt.(float64) + float64(1)))

							return editable
						}
					} else {
						{
							var len = float64(len(self__.Arr.([]interface{})))
							var new_arr = make([]interface{}, int((len + float64(2))))
							_, _ = len, new_arr
							Array_copy.X_invoke_Arity5(self__.Arr, float64(0), new_arr, float64(0), len).(interface{})
							new_arr[int(len)] = key
							new_arr[int((len + float64(1)))] = val
							Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)

							return inode___1.Ensure_editable_array(edit___1, (self__.Cnt.(float64) + float64(1)), new_arr)
						}
					}
				} else {
					if (self__.Arr.([]interface{})[int((idx + float64(1)))]) == val {
						return inode___1
					} else {
						return Edit_and_set.X_invoke_Arity4(inode___1, edit___1, (idx + float64(1)), val).(interface{})
					}
				}
			}
		} else {
			return (&CljsCoreBitmapIndexedNode{edit___1, float64(1 << uint(float64((uint(self__.Collision_hash)>>uint(shift))&0x01f))), []interface{}{nil, inode___1, nil, nil}}).Inode_assoc_BANG_(edit___1, shift, hash, key, val, added_leaf_QMARK_)
		}
	}
}

func (self__ *CljsCoreHashCollisionNode) Inode_assoc(shift interface{}, hash interface{}, key interface{}, val interface{}, added_leaf_QMARK_ interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		if hash.(float64) == self__.Collision_hash.(float64) {
			{
				var idx = Hash_collision_node_find_index.X_invoke_Arity3(self__.Arr, self__.Cnt, key).(float64)
				_ = idx
				if idx == float64(-1) {
					{
						var len = (float64(2) * self__.Cnt.(float64))
						var new_arr = make([]interface{}, int((len + float64(2))))
						_, _ = len, new_arr
						Array_copy.X_invoke_Arity5(self__.Arr, float64(0), new_arr, float64(0), len).(interface{})
						new_arr[int(len)] = key
						new_arr[int((len + float64(1)))] = val
						Native_set_instance_field.X_invoke_Arity3(added_leaf_QMARK_, "Val", true)

						return (&CljsCoreHashCollisionNode{nil, self__.Collision_hash, (self__.Cnt.(float64) + float64(1)), new_arr})
					}
				} else {
					if X_EQ_.Arity2IIB((self__.Arr.([]interface{})[int(idx)]), val) {
						return inode___1
					} else {
						return (&CljsCoreHashCollisionNode{nil, self__.Collision_hash, self__.Cnt, Clone_and_set.X_invoke_Arity3(self__.Arr, (idx + float64(1)), val).([]interface{})})
					}
				}
			}
		} else {
			return (&CljsCoreBitmapIndexedNode{nil, float64(1 << uint(float64((uint(self__.Collision_hash)>>uint(shift))&0x01f))), []interface{}{nil, inode___1}}).Inode_assoc(shift, hash, key, val, added_leaf_QMARK_)
		}
	}
}

func (self__ *CljsCoreHashCollisionNode) Ensure_editable_array(e interface{}, count interface{}, array interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		if e == self__.Edit {
			self__.Arr = array

			self__.Cnt = count

			return inode___1
		} else {
			return (&CljsCoreHashCollisionNode{self__.Edit, self__.Collision_hash, count, array})
		}
	}
}

func (self__ *CljsCoreHashCollisionNode) Inode_find(shift interface{}, hash interface{}, key interface{}, not_found interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var idx = Hash_collision_node_find_index.X_invoke_Arity3(self__.Arr, self__.Cnt, key).(float64)
			_ = idx
			if idx < float64(0) {
				return not_found
			} else {
				if Key_test.Arity2IIB(key, (self__.Arr.([]interface{})[int(idx)])) {
					return &CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{(self__.Arr.([]interface{})[int(idx)]), (self__.Arr.([]interface{})[int((idx + float64(1)))])}, nil}
				} else {
					return not_found

				}
			}
		}
	}
}

func (self__ *CljsCoreHashCollisionNode) Inode_without(shift interface{}, hash interface{}, key interface{}) interface{} {
	{
		var inode___1 = self__
		_ = inode___1
		{
			var idx = Hash_collision_node_find_index.X_invoke_Arity3(self__.Arr, self__.Cnt, key).(float64)
			_ = idx
			if idx == float64(-1) {
				return inode___1
			} else {
				if self__.Cnt.(float64) == float64(1) {
					return nil
				} else {
					return (&CljsCoreHashCollisionNode{nil, self__.Collision_hash, (self__.Cnt.(float64) - float64(1)), Remove_pair.X_invoke_Arity2(self__.Arr, Quot.X_invoke_Arity2(idx, float64(2)).(float64)).([]interface{})})

				}
			}
		}
	}
}

var X__GT_HashCollisionNode = func(__GT_HashCollisionNode *AFn) *AFn {
	return Fn(__GT_HashCollisionNode, func(edit interface{}, collision_hash interface{}, cnt interface{}, arr interface{}) interface{} {
		return (&CljsCoreHashCollisionNode{edit, collision_hash, cnt, arr})
	})
}(&AFn{})

var Create_node = func(create_node *AFn) *AFn {
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

type CljsCoreNodeSeq struct {
	Meta    interface{}
	Nodes   interface{}
	I       interface{}
	S       interface{}
	X__hash interface{}
}

func (_ *CljsCoreNodeSeq) CljsCoreObject__() {}
func (self__ *CljsCoreNodeSeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreNodeSeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreNodeSeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreIMeta__() {}
func (self__ *CljsCoreNodeSeq) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreIHash__() {}
func (self__ *CljsCoreNodeSeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreIEquiv__() {}
func (self__ *CljsCoreNodeSeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreNodeSeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.Meta)
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreIReduce__() {}
func (self__ *CljsCoreNodeSeq) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCoreNodeSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreISeq__() {}
func (self__ *CljsCoreNodeSeq) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.S == nil {
			return &CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{(self__.Nodes.([]interface{})[int(self__.I.(float64))]), (self__.Nodes.([]interface{})[int((self__.I.(float64) + float64(1)))])}, nil}
		} else {
			return First.X_invoke_Arity1(self__.S)
		}
	}
}

func (self__ *CljsCoreNodeSeq) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.S == nil {
			return Create_inode_seq.X_invoke_Arity3(self__.Nodes, (self__.I.(float64) + float64(2)), nil).(interface{})
		} else {
			return Create_inode_seq.X_invoke_Arity3(self__.Nodes, self__.I, Next.Arity1IQ(self__.S)).(interface{})
		}
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreISeqable__() {}
func (self__ *CljsCoreNodeSeq) X_seq_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		return this___1
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreISequential__() {}
func (_ *CljsCoreNodeSeq) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreNodeSeq) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreNodeSeq{meta___1, self__.Nodes, self__.I, self__.S, self__.X__hash})
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreICollection__() {}
func (self__ *CljsCoreNodeSeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(*CljsCoreCons)
	}
}

var X__GT_NodeSeq = func(__GT_NodeSeq *AFn) *AFn {
	return Fn(__GT_NodeSeq, func(meta interface{}, nodes interface{}, i interface{}, s interface{}, __hash interface{}) interface{} {
		return (&CljsCoreNodeSeq{meta, nodes, i, s, __hash})
	})
}(&AFn{})

var Create_inode_seq = func(create_inode_seq *AFn) *AFn {
	return Fn(create_inode_seq, func(nodes interface{}) interface{} {
		return create_inode_seq.X_invoke_Arity3(nodes, float64(0), nil)
	}, func(nodes interface{}, i interface{}, s interface{}) interface{} {
		if s == nil {
			{
				var len = float64(len(nodes.([]interface{})))
				_ = len
				{
					var j = i
					_ = j
					for {
						if j.(float64) < len {
							if !((nodes.([]interface{})[int(j.(float64))]) == nil) {
								return (&CljsCoreNodeSeq{nil, nodes, j, nil, nil})
							} else {
								{
									var temp__4217__auto__ = (nodes.([]interface{})[int((j.(float64) + float64(1)))])
									_ = temp__4217__auto__
									if Truth_(temp__4217__auto__) {
										{
											var node = temp__4217__auto__
											_ = node
											{
												var temp__4217__auto_____1 = Native_invoke_instance_method.X_invoke_Arity3(node, "Inode_seq", []interface{}{})
												_ = temp__4217__auto_____1
												if Truth_(temp__4217__auto_____1) {
													{
														var node_seq = temp__4217__auto_____1
														_ = node_seq
														return (&CljsCoreNodeSeq{nil, nodes, (j.(float64) + float64(2)), node_seq, nil})
													}
												} else {
													{
														var G__813 = (j.(float64) + float64(2))
														j = G__813
														continue
													}
												}
											}
										}
									} else {
										{
											var G__814 = (j.(float64) + float64(2))
											j = G__814
											continue
										}
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

type CljsCoreArrayNodeSeq struct {
	Meta    interface{}
	Nodes   interface{}
	I       interface{}
	S       interface{}
	X__hash interface{}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreObject__() {}
func (self__ *CljsCoreArrayNodeSeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreArrayNodeSeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreArrayNodeSeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIMeta__() {}
func (self__ *CljsCoreArrayNodeSeq) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIHash__() {}
func (self__ *CljsCoreArrayNodeSeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIEquiv__() {}
func (self__ *CljsCoreArrayNodeSeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreArrayNodeSeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.Meta)
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIReduce__() {}
func (self__ *CljsCoreArrayNodeSeq) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCoreArrayNodeSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreISeq__() {}
func (self__ *CljsCoreArrayNodeSeq) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return First.X_invoke_Arity1(self__.S)
	}
}

func (self__ *CljsCoreArrayNodeSeq) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Create_array_node_seq.X_invoke_Arity4(nil, self__.Nodes, self__.I, Next.Arity1IQ(self__.S)).(interface{})
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreISeqable__() {}
func (self__ *CljsCoreArrayNodeSeq) X_seq_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		return this___1
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreISequential__() {}
func (_ *CljsCoreArrayNodeSeq) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreArrayNodeSeq) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreArrayNodeSeq{meta___1, self__.Nodes, self__.I, self__.S, self__.X__hash})
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreICollection__() {}
func (self__ *CljsCoreArrayNodeSeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(*CljsCoreCons)
	}
}

var X__GT_ArrayNodeSeq = func(__GT_ArrayNodeSeq *AFn) *AFn {
	return Fn(__GT_ArrayNodeSeq, func(meta interface{}, nodes interface{}, i interface{}, s interface{}, __hash interface{}) interface{} {
		return (&CljsCoreArrayNodeSeq{meta, nodes, i, s, __hash})
	})
}(&AFn{})

var Create_array_node_seq = func(create_array_node_seq *AFn) *AFn {
	return Fn(create_array_node_seq, func(nodes interface{}) interface{} {
		return create_array_node_seq.X_invoke_Arity4(nil, nodes, float64(0), nil)
	}, func(meta interface{}, nodes interface{}, i interface{}, s interface{}) interface{} {
		if s == nil {
			{
				var len = float64(len(nodes.([]interface{})))
				_ = len
				{
					var j = i
					_ = j
					for {
						if j.(float64) < len {
							{
								var temp__4217__auto__ = (nodes.([]interface{})[int(j.(float64))])
								_ = temp__4217__auto__
								if Truth_(temp__4217__auto__) {
									{
										var nj = temp__4217__auto__
										_ = nj
										{
											var temp__4217__auto_____1 = Native_invoke_instance_method.X_invoke_Arity3(nj, "Inode_seq", []interface{}{})
											_ = temp__4217__auto_____1
											if Truth_(temp__4217__auto_____1) {
												{
													var ns = temp__4217__auto_____1
													_ = ns
													return (&CljsCoreArrayNodeSeq{meta, nodes, (j.(float64) + float64(1)), ns, nil})
												}
											} else {
												{
													var G__815 = (j.(float64) + float64(1))
													j = G__815
													continue
												}
											}
										}
									}
								} else {
									{
										var G__816 = (j.(float64) + float64(1))
										j = G__816
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
			return (&CljsCoreArrayNodeSeq{meta, nodes, i, s, nil})
		}
	})
}(&AFn{})

type CljsCorePersistentHashMap struct {
	Meta           interface{}
	Cnt            interface{}
	Root           interface{}
	Has_nil_QMARK_ bool
	Nil_val        interface{}
	X__hash        interface{}
}

func (_ *CljsCorePersistentHashMap) CljsCoreObject__() {}
func (self__ *CljsCorePersistentHashMap) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCorePersistentHashMap) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCorePersistentHashMap) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (self__ *CljsCorePersistentHashMap) Keys() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Iterator.X_invoke_Arity1(keys.X_invoke_Arity1(coll___1).(*CljsCoreIterator)).(*CljsCoreIterator)
	}
}

func (self__ *CljsCorePersistentHashMap) Entries() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Entries_iterator.X_invoke_Arity1(Seq.X_invoke_Arity1(coll___1)).(*CljsCoreEntriesIterator)
	}
}

func (self__ *CljsCorePersistentHashMap) Values() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Iterator.X_invoke_Arity1(Vals.X_invoke_Arity1(coll___1).(interface{})).(*CljsCoreIterator)
	}
}

func (self__ *CljsCorePersistentHashMap) Has(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Contains_QMARK_.X_invoke_Arity2(coll___1, k)
	}
}

func (self__ *CljsCorePersistentHashMap) Get(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCorePersistentHashMap) ForEach(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var seq__823 = Seq.X_invoke_Arity1(coll___1)
			var chunk__824 = nil
			var count__825 = float64(0)
			var i__826 = float64(0)
			_, _, _, _ = seq__823, chunk__824, count__825, i__826
			for {
				if i__826 < count__825 {
					{
						var vec__827 = chunk__824.X_nth_Arity2(i__826).(interface{})
						var k = Nth.X_invoke_Arity3(vec__827, float64(0), nil)
						var v = Nth.X_invoke_Arity3(vec__827, float64(1), nil)
						_, _, _ = vec__827, k, v
						f.(CljsCoreIFn).X_invoke_Arity2(v, k).(interface{})
						{
							var G__833 = seq__823
							var G__834 = chunk__824
							var G__835 = count__825
							var G__836 = (i__826 + float64(1))
							seq__823 = G__833
							chunk__824 = G__834
							count__825 = G__835
							i__826 = G__836
							continue
						}
					}
				} else {
					{
						var temp__4219__auto__ = Seq.X_invoke_Arity1(seq__823)
						_ = temp__4219__auto__
						if Truth_(temp__4219__auto__) {
							{
								var seq__823___1 = temp__4219__auto__
								_ = seq__823___1
								if Chunked_seq_QMARK_.X_invoke_Arity1(seq__823___1) {
									{
										var c__77428__auto__ = Chunk_first.X_invoke_Arity1(seq__823___1).(interface{})
										_ = c__77428__auto__
										{
											var G__837 = Chunk_rest.X_invoke_Arity1(seq__823___1).(interface{})
											var G__838 = c__77428__auto__
											var G__839 = Count.X_invoke_Arity1(c__77428__auto__).(float64)
											var G__840 = float64(0)
											seq__823 = G__837
											chunk__824 = G__838
											count__825 = G__839
											i__826 = G__840
											continue
										}
									}
								} else {
									{
										var vec__828 = First.X_invoke_Arity1(seq__823___1)
										var k = Nth.X_invoke_Arity3(vec__828, float64(0), nil)
										var v = Nth.X_invoke_Arity3(vec__828, float64(1), nil)
										_, _, _ = vec__828, k, v
										f.(CljsCoreIFn).X_invoke_Arity2(v, k).(interface{})
										{
											var G__841 = Next.X_invoke_Arity1(seq__823___1)
											var G__842 = nil
											var G__843 = float64(0)
											var G__844 = float64(0)
											seq__823 = G__841
											chunk__824 = G__842
											count__825 = G__843
											i__826 = G__844
											continue
										}
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
}

func (_ *CljsCorePersistentHashMap) CljsCoreILookup__() {}
func (self__ *CljsCorePersistentHashMap) X_lookup_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return _lookup.X_invoke_Arity3(coll___1, k, nil).(interface{})
	}
}

func (self__ *CljsCorePersistentHashMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if k == nil {
			if self__.Has_nil_QMARK_ {
				return self__.Nil_val
			} else {
				return not_found
			}
		} else {
			if self__.Root == nil {
				return not_found
			} else {
				return Native_invoke_instance_method.X_invoke_Arity3(self__.Root, "Inode_lookup", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k, not_found})

			}
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIKVReduce__() {}
func (self__ *CljsCorePersistentHashMap) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var init___1 = func() interface{} {
				if self__.Has_nil_QMARK_ {
					return f.(CljsCoreIFn).X_invoke_Arity3(init, nil, self__.Nil_val).(interface{})
				} else {
					return init
				}
			}()
			_ = init___1
			if Reduced_QMARK_.X_invoke_Arity1(init___1) {
				return Deref.X_invoke_Arity1(init___1).(interface{})
			} else {
				if !(self__.Root == nil) {
					return Native_invoke_instance_method.X_invoke_Arity3(self__.Root, "Kv_reduce", []interface{}{f, init___1})
				} else {
					return init___1

				}
			}
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIMeta__() {}
func (self__ *CljsCorePersistentHashMap) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreICloneable__() {}
func (self__ *CljsCorePersistentHashMap) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCorePersistentHashMap{self__.Meta, self__.Cnt, self__.Root, self__.Has_nil_QMARK_, self__.Nil_val, self__.X__hash})
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreICounted__() {}
func (self__ *CljsCorePersistentHashMap) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Cnt.(float64)
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIHash__() {}
func (self__ *CljsCorePersistentHashMap) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_unordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIEquiv__() {}
func (self__ *CljsCorePersistentHashMap) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_map.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIEditableCollection__() {}
func (self__ *CljsCorePersistentHashMap) X_as_transient_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreTransientHashMap{func() interface{} {
			var obj832 = map[string]interface{}{}
			_ = obj832
			return obj832
		}(), self__.Root, self__.Cnt, self__.Has_nil_QMARK_, self__.Nil_val})
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCorePersistentHashMap) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return CljsCorePersistentHashMap_EMPTY.(CljsCoreIWithMeta).X_with_meta_Arity2(self__.Meta)
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIMap__() {}
func (self__ *CljsCorePersistentHashMap) X_dissoc_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if k == nil {
			if self__.Has_nil_QMARK_ {
				return (&CljsCorePersistentHashMap{self__.Meta, (self__.Cnt.(float64) - float64(1)), self__.Root, false, nil, nil})
			} else {
				return coll___1
			}
		} else {
			if self__.Root == nil {
				return coll___1
			} else {
				{
					var new_root = Native_invoke_instance_method.X_invoke_Arity3(self__.Root, "Inode_without", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k})
					_ = new_root
					if new_root == self__.Root {
						return coll___1
					} else {
						return (&CljsCorePersistentHashMap{self__.Meta, (self__.Cnt.(float64) - float64(1)), new_root, self__.Has_nil_QMARK_, self__.Nil_val, nil})
					}
				}

			}
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIAssociative__() {}
func (self__ *CljsCorePersistentHashMap) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if k == nil {
			if (self__.Has_nil_QMARK_) && (v == self__.Nil_val) {
				return coll___1
			} else {
				return (&CljsCorePersistentHashMap{self__.Meta, func() interface{} {
					if self__.Has_nil_QMARK_ {
						return self__.Cnt
					} else {
						return (self__.Cnt.(float64) + float64(1))
					}
				}(), self__.Root, true, v, nil})
			}
		} else {
			{
				var added_leaf_QMARK_ = (&CljsCoreBox{false})
				var new_root = Native_invoke_instance_method.X_invoke_Arity3(func() interface{} {
					if self__.Root == nil {
						return CljsCoreBitmapIndexedNode_EMPTY
					} else {
						return self__.Root
					}
				}(), "Inode_assoc", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k, v, added_leaf_QMARK_})
				_, _ = added_leaf_QMARK_, new_root
				if new_root == self__.Root {
					return coll___1
				} else {
					return (&CljsCorePersistentHashMap{self__.Meta, func() interface{} {
						if added_leaf_QMARK_.Val {
							return (self__.Cnt.(float64) + float64(1))
						} else {
							return self__.Cnt
						}
					}(), new_root, self__.Has_nil_QMARK_, self__.Nil_val, nil})
				}
			}
		}
	}
}

func (self__ *CljsCorePersistentHashMap) X_contains_key_QMARK__Arity2(k interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		if k == nil {
			return self__.Has_nil_QMARK_
		} else {
			if self__.Root == nil {
				return false
			} else {
				return !(Native_invoke_instance_method.X_invoke_Arity3(self__.Root, "Inode_lookup", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k, Lookup_sentinel}) == Lookup_sentinel)

			}
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreISeqable__() {}
func (self__ *CljsCorePersistentHashMap) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Cnt.(float64) > float64(0) {
			{
				var s = func() interface{} {
					if !(self__.Root == nil) {
						return Native_invoke_instance_method.X_invoke_Arity3(self__.Root, "Inode_seq", []interface{}{})
					} else {
						return nil
					}
				}()
				_ = s
				if self__.Has_nil_QMARK_ {
					return Cons.X_invoke_Arity2(&CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, self__.Nil_val}, nil}, s).(*CljsCoreCons)
				} else {
					return s
				}
			}
		} else {
			return nil
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIWithMeta__() {}
func (self__ *CljsCorePersistentHashMap) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentHashMap{meta___1, self__.Cnt, self__.Root, self__.Has_nil_QMARK_, self__.Nil_val, self__.X__hash})
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreICollection__() {}
func (self__ *CljsCorePersistentHashMap) X_conj_Arity2(entry interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if Vector_QMARK_.Arity1IB(entry) {
			return coll___1.X_assoc_Arity3(entry.(CljsCoreIIndexed).X_nth_Arity2(float64(0)).(interface{}), entry.(CljsCoreIIndexed).X_nth_Arity2(float64(1)).(interface{}))
		} else {
			{
				var ret = coll___1
				var es = Seq.Arity1IQ(entry)
				_, _ = ret, es
				for {
					if es == nil {
						return ret
					} else {
						{
							var e = First.X_invoke_Arity1(es)
							_ = e
							if Vector_QMARK_.X_invoke_Arity1(e) {
								{
									var G__845 = ret.X_assoc_Arity3(e.X_nth_Arity2(float64(0)).(interface{}), e.X_nth_Arity2(float64(1)).(interface{}))
									var G__846 = Next.X_invoke_Arity1(es)
									ret = G__845
									es = G__846
									continue
								}
							} else {
								panic((&js.Error{"conj on a map takes map entries or seqables of map entries"}))
							}
						}
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIFn__() {}
func (self__ *CljsCorePersistentHashMap) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity1(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCorePersistentHashMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

var X__GT_PersistentHashMap = func(__GT_PersistentHashMap *AFn) *AFn {
	return Fn(__GT_PersistentHashMap, func(meta interface{}, cnt interface{}, root interface{}, has_nil_QMARK_ bool, nil_val interface{}, __hash interface{}) interface{} {
		return (&CljsCorePersistentHashMap{meta, cnt, root, has_nil_QMARK_, nil_val, __hash})
	})
}(&AFn{})

var CljsCorePersistentHashMap_EMPTY = (&CljsCorePersistentHashMap{nil, float64(0), nil, false, nil, float64(0)})

var CljsCorePersistentHashMap_FromArrays = func(G__847 *AFn) *AFn {
	return Fn(G__847, func(ks interface{}, vs interface{}) interface{} {
		{
			var len = float64(len(ks.([]interface{})))
			_ = len
			{
				var i = float64(0)
				var out = Transient.X_invoke_Arity1(CljsCorePersistentHashMap_EMPTY).(interface{})
				_, _ = i, out
				for {
					if i < len {
						{
							var G__848 = (i + float64(1))
							var G__849 = out.X_assoc_BANG__Arity3((ks.([]interface{})[int(i)]), (vs.([]interface{})[int(i)]))
							i = G__848
							out = G__849
							continue
						}
					} else {
						return Persistent_BANG_.X_invoke_Arity1(out).(interface{})
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
func (self__ *CljsCoreTransientHashMap) X_dissoc_BANG__Arity2(key interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		return tcoll___1.Without_BANG_(key)
	}
}

func (_ *CljsCoreTransientHashMap) CljsCoreITransientAssociative__() {}
func (self__ *CljsCoreTransientHashMap) X_assoc_BANG__Arity3(key interface{}, val interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		return tcoll___1.Assoc_BANG_(key, val)
	}
}

func (_ *CljsCoreTransientHashMap) CljsCoreITransientCollection__() {}
func (self__ *CljsCoreTransientHashMap) X_conj_BANG__Arity2(val interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		return tcoll___1.Conj_BANG_(val)
	}
}

func (self__ *CljsCoreTransientHashMap) X_persistent_BANG__Arity1() interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		return tcoll___1.Persistent_BANG_()
	}
}

func (_ *CljsCoreTransientHashMap) CljsCoreILookup__() {}
func (self__ *CljsCoreTransientHashMap) X_lookup_Arity2(k interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if k == nil {
			if self__.Has_nil_QMARK_ {
				return self__.Nil_val
			} else {
				return nil
			}
		} else {
			if self__.Root == nil {
				return nil
			} else {
				return Native_invoke_instance_method.X_invoke_Arity3(self__.Root, "Inode_lookup", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k})
			}
		}
	}
}

func (self__ *CljsCoreTransientHashMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if k == nil {
			if self__.Has_nil_QMARK_ {
				return self__.Nil_val
			} else {
				return not_found
			}
		} else {
			if self__.Root == nil {
				return not_found
			} else {
				return Native_invoke_instance_method.X_invoke_Arity3(self__.Root, "Inode_lookup", []interface{}{float64(0), Hash.X_invoke_Arity1(k), k, not_found})
			}
		}
	}
}

func (_ *CljsCoreTransientHashMap) CljsCoreICounted__() {}
func (self__ *CljsCoreTransientHashMap) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Edit {
			return self__.Count.(float64)
		} else {
			panic((&js.Error{"count after persistent!"}))
		}
	}
}

func (_ *CljsCoreTransientHashMap) CljsCoreObject__() {}
func (self__ *CljsCoreTransientHashMap) Conj_BANG_(o interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if self__.Edit {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IMapEntry", Str: "cljs.core/IMapEntry", X_hash: float64(535941300), X_meta: nil}), o)) {
				return tcoll___1.Assoc_BANG_(Key.X_invoke_Arity1(o).(interface{}), Val.X_invoke_Arity1(o).(interface{}))
			} else {
				{
					var es = Seq.Arity1IQ(o)
					var tcoll___2 = tcoll___1
					_, _ = es, tcoll___2
					for {
						{
							var temp__4217__auto__ = First.X_invoke_Arity1(es)
							_ = temp__4217__auto__
							if Truth_(temp__4217__auto__) {
								{
									var e = temp__4217__auto__
									_ = e
									{
										var G__850 = Next.X_invoke_Arity1(es)
										var G__851 = tcoll___2.Assoc_BANG_(Key.X_invoke_Arity1(e).(interface{}), Val.X_invoke_Arity1(e).(interface{}))
										es = G__850
										tcoll___2 = G__851
										continue
									}
								}
							} else {
								return tcoll___2
							}
						}
					}
				}
			}
		} else {
			panic((&js.Error{"conj! after persistent"}))
		}
	}
}

func (self__ *CljsCoreTransientHashMap) Assoc_BANG_(k interface{}, v interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if self__.Edit {
			if k == nil {
				if self__.Nil_val == v {
				} else {
					self__.Nil_val = v

				}
				if self__.Has_nil_QMARK_ {
				} else {
					self__.Count = (self__.Count.(float64) + float64(1))

					self__.Has_nil_QMARK_ = true

				}
				return tcoll___1
			} else {
				{
					var added_leaf_QMARK_ = (&CljsCoreBox{false})
					var node = Native_invoke_instance_method.X_invoke_Arity3(func() interface{} {
						if self__.Root == nil {
							return CljsCoreBitmapIndexedNode_EMPTY
						} else {
							return self__.Root
						}
					}(), "Inode_assoc_BANG_", []interface{}{self__.Edit, float64(0), Hash.X_invoke_Arity1(k), k, v, added_leaf_QMARK_})
					_, _ = added_leaf_QMARK_, node
					if node == self__.Root {
					} else {
						self__.Root = node

					}
					if added_leaf_QMARK_.Val {
						self__.Count = (self__.Count.(float64) + float64(1))

					} else {
					}
					return tcoll___1
				}
			}
		} else {
			panic((&js.Error{"assoc! after persistent!"}))
		}
	}
}

func (self__ *CljsCoreTransientHashMap) Without_BANG_(k interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if self__.Edit {
			if k == nil {
				if self__.Has_nil_QMARK_ {
					self__.Has_nil_QMARK_ = false

					self__.Nil_val = nil

					self__.Count = (self__.Count.(float64) - float64(1))

					return tcoll___1
				} else {
					return tcoll___1
				}
			} else {
				if self__.Root == nil {
					return tcoll___1
				} else {
					{
						var removed_leaf_QMARK_ = (&CljsCoreBox{false})
						var node = Native_invoke_instance_method.X_invoke_Arity3(self__.Root, "Inode_without_BANG_", []interface{}{self__.Edit, float64(0), Hash.X_invoke_Arity1(k), k, removed_leaf_QMARK_})
						_, _ = removed_leaf_QMARK_, node
						if node == self__.Root {
						} else {
							self__.Root = node

						}
						if Truth_((removed_leaf_QMARK_.([]interface{})[int(float64(0))])) {
							self__.Count = (self__.Count.(float64) - float64(1))

						} else {
						}
						return tcoll___1
					}
				}
			}
		} else {
			panic((&js.Error{"dissoc! after persistent!"}))
		}
	}
}

func (self__ *CljsCoreTransientHashMap) Persistent_BANG_() interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if self__.Edit {
			self__.Edit = nil

			return (&CljsCorePersistentHashMap{nil, self__.Count, self__.Root, self__.Has_nil_QMARK_, self__.Nil_val, nil})
		} else {
			panic((&js.Error{"persistent! called twice"}))
		}
	}
}

var X__GT_TransientHashMap = func(__GT_TransientHashMap *AFn) *AFn {
	return Fn(__GT_TransientHashMap, func(edit bool, root interface{}, count interface{}, has_nil_QMARK_ bool, nil_val interface{}) interface{} {
		return (&CljsCoreTransientHashMap{edit, root, count, has_nil_QMARK_, nil_val})
	})
}(&AFn{})

var Tree_map_seq_push = func(tree_map_seq_push *AFn) *AFn {
	return Fn(tree_map_seq_push, func(node interface{}, stack interface{}, ascending_QMARK_ bool) interface{} {
		{
			var t = node
			var stack___1 = stack
			_, _ = t, stack___1
			for {
				if !(t == nil) {
					{
						var G__852 = func() interface{} {
							if ascending_QMARK_ {
								return Native_get_instance_field.X_invoke_Arity2(t, "Left")
							} else {
								return Native_get_instance_field.X_invoke_Arity2(t, "Right")
							}
						}()
						var G__853 = Conj.X_invoke_Arity2(stack___1, t).(interface{})
						t = G__852
						stack___1 = G__853
						continue
					}
				} else {
					return stack___1
				}
			}
		}
	})
}(&AFn{})

type CljsCorePersistentTreeMapSeq struct {
	Meta             interface{}
	Stack            interface{}
	Ascending_QMARK_ bool
	Cnt              interface{}
	X__hash          interface{}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreObject__() {}
func (self__ *CljsCorePersistentTreeMapSeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCorePersistentTreeMapSeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCorePersistentTreeMapSeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIMeta__() {}
func (self__ *CljsCorePersistentTreeMapSeq) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreICounted__() {}
func (self__ *CljsCorePersistentTreeMapSeq) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Cnt.(float64) < float64(0) {
			return (Count.X_invoke_Arity1(Next.X_invoke_Arity1(coll___1)).(float64) + float64(1))
		} else {
			return self__.Cnt
		}
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIHash__() {}
func (self__ *CljsCorePersistentTreeMapSeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIEquiv__() {}
func (self__ *CljsCorePersistentTreeMapSeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCorePersistentTreeMapSeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.Meta)
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIReduce__() {}
func (self__ *CljsCorePersistentTreeMapSeq) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCorePersistentTreeMapSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreISeq__() {}
func (self__ *CljsCorePersistentTreeMapSeq) X_first_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		return Peek.X_invoke_Arity1(self__.Stack)
	}
}

func (self__ *CljsCorePersistentTreeMapSeq) X_rest_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		{
			var t = First.X_invoke_Arity1(self__.Stack)
			var next_stack = Tree_map_seq_push.X_invoke_Arity3(func() interface{} {
				if self__.Ascending_QMARK_ {
					return t.Right
				} else {
					return t.Left
				}
			}(), Next.Arity1IQ(self__.Stack), self__.Ascending_QMARK_).(interface{})
			_, _ = t, next_stack
			if !(next_stack == nil) {
				return (&CljsCorePersistentTreeMapSeq{nil, next_stack, self__.Ascending_QMARK_, (self__.Cnt.(float64) - float64(1)), nil})
			} else {
				return CljsCoreList_EMPTY
			}
		}
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreISeqable__() {}
func (self__ *CljsCorePersistentTreeMapSeq) X_seq_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		return this___1
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreISequential__() {}
func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIWithMeta__()   {}
func (self__ *CljsCorePersistentTreeMapSeq) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentTreeMapSeq{meta___1, self__.Stack, self__.Ascending_QMARK_, self__.Cnt, self__.X__hash})
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreICollection__() {}
func (self__ *CljsCorePersistentTreeMapSeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(*CljsCoreCons)
	}
}

var X__GT_PersistentTreeMapSeq = func(__GT_PersistentTreeMapSeq *AFn) *AFn {
	return Fn(__GT_PersistentTreeMapSeq, func(meta interface{}, stack interface{}, ascending_QMARK_ bool, cnt interface{}, __hash interface{}) interface{} {
		return (&CljsCorePersistentTreeMapSeq{meta, stack, ascending_QMARK_, cnt, __hash})
	})
}(&AFn{})

var Create_tree_map_seq = func(create_tree_map_seq *AFn) *AFn {
	return Fn(create_tree_map_seq, func(tree interface{}, ascending_QMARK_ interface{}, cnt interface{}) interface{} {
		return (&CljsCorePersistentTreeMapSeq{nil, Tree_map_seq_push.X_invoke_Arity3(tree, nil, ascending_QMARK_).(interface{}), ascending_QMARK_, cnt, nil})
	})
}(&AFn{})

var Balance_left = func(balance_left *AFn) *AFn {
	return Fn(balance_left, func(key interface{}, val interface{}, ins interface{}, right interface{}) interface{} {
		if func() bool { _, instanceof := ins.(*RedNode); return instanceof }() {
			if func() bool {
				_, instanceof := Native_get_instance_field.X_invoke_Arity2(ins, "Left").(*RedNode)
				return instanceof
			}() {
				return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(ins, "Key"), Native_get_instance_field.X_invoke_Arity2(ins, "Val"), Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(ins, "Left"), "Blacken", []interface{}{}), (&CljsCoreBlackNode{key, val, Native_get_instance_field.X_invoke_Arity2(ins, "Right"), right, nil}), nil})
			} else {
				if func() bool {
					_, instanceof := Native_get_instance_field.X_invoke_Arity2(ins, "Right").(*RedNode)
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

var Balance_right = func(balance_right *AFn) *AFn {
	return Fn(balance_right, func(key interface{}, val interface{}, left interface{}, ins interface{}) interface{} {
		if func() bool { _, instanceof := ins.(*RedNode); return instanceof }() {
			if func() bool {
				_, instanceof := Native_get_instance_field.X_invoke_Arity2(ins, "Right").(*RedNode)
				return instanceof
			}() {
				return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(ins, "Key"), Native_get_instance_field.X_invoke_Arity2(ins, "Val"), (&CljsCoreBlackNode{key, val, left, Native_get_instance_field.X_invoke_Arity2(ins, "Left"), nil}), Native_invoke_instance_method.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(ins, "Right"), "Blacken", []interface{}{}), nil})
			} else {
				if func() bool {
					_, instanceof := Native_get_instance_field.X_invoke_Arity2(ins, "Left").(*RedNode)
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

var Balance_left_del = func(balance_left_del *AFn) *AFn {
	return Fn(balance_left_del, func(key interface{}, val interface{}, del interface{}, right interface{}) interface{} {
		if func() bool { _, instanceof := del.(*RedNode); return instanceof }() {
			return (&CljsCoreRedNode{key, val, Native_invoke_instance_method.X_invoke_Arity3(del, "Blacken", []interface{}{}), right, nil})
		} else {
			if func() bool { _, instanceof := right.(*BlackNode); return instanceof }() {
				return Balance_right.X_invoke_Arity4(key, val, del, Native_invoke_instance_method.X_invoke_Arity3(right, "Redden", []interface{}{}))
			} else {
				if (func() bool { _, instanceof := right.(*RedNode); return instanceof }()) && (func() bool {
					_, instanceof := Native_get_instance_field.X_invoke_Arity2(right, "Left").(*BlackNode)
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

var Balance_right_del = func(balance_right_del *AFn) *AFn {
	return Fn(balance_right_del, func(key interface{}, val interface{}, left interface{}, del interface{}) interface{} {
		if func() bool { _, instanceof := del.(*RedNode); return instanceof }() {
			return (&CljsCoreRedNode{key, val, left, Native_invoke_instance_method.X_invoke_Arity3(del, "Blacken", []interface{}{}), nil})
		} else {
			if func() bool { _, instanceof := left.(*BlackNode); return instanceof }() {
				return Balance_left.X_invoke_Arity4(key, val, Native_invoke_instance_method.X_invoke_Arity3(left, "Redden", []interface{}{}), del)
			} else {
				if (func() bool { _, instanceof := left.(*RedNode); return instanceof }()) && (func() bool {
					_, instanceof := Native_get_instance_field.X_invoke_Arity2(left, "Right").(*BlackNode)
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

var Tree_map_kv_reduce = func(tree_map_kv_reduce *AFn) *AFn {
	return Fn(tree_map_kv_reduce, func(node interface{}, f interface{}, init interface{}) interface{} {
		{
			var init___1 = func() interface{} {
				if !(Native_get_instance_field.X_invoke_Arity2(node, "Left") == nil) {
					return tree_map_kv_reduce.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(node, "Left"), f, init).(interface{})
				} else {
					return init
				}
			}()
			_ = init___1
			if Reduced_QMARK_.X_invoke_Arity1(init___1) {
				return Deref.X_invoke_Arity1(init___1).(interface{})
			} else {
				{
					var init___2 = f.(CljsCoreIFn).X_invoke_Arity3(init___1, Native_get_instance_field.X_invoke_Arity2(node, "Key"), Native_get_instance_field.X_invoke_Arity2(node, "Val")).(interface{})
					_ = init___2
					if Reduced_QMARK_.X_invoke_Arity1(init___2) {
						return Deref.X_invoke_Arity1(init___2).(interface{})
					} else {
						{
							var init___3 = func() interface{} {
								if !(Native_get_instance_field.X_invoke_Arity2(node, "Right") == nil) {
									return tree_map_kv_reduce.X_invoke_Arity3(Native_get_instance_field.X_invoke_Arity2(node, "Right"), f, init___2).(interface{})
								} else {
									return init___2
								}
							}()
							_ = init___3
							if Reduced_QMARK_.X_invoke_Arity1(init___3) {
								return Deref.X_invoke_Arity1(init___3).(interface{})
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

type CljsCoreBlackNode struct {
	Key     interface{}
	Val     interface{}
	Left    interface{}
	Right   interface{}
	X__hash interface{}
}

func (_ *CljsCoreBlackNode) CljsCoreObject__() {}
func (self__ *CljsCoreBlackNode) Add_right(ins interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Native_invoke_instance_method.X_invoke_Arity3(ins, "Balance_right", []interface{}{node___1})
	}
}

func (self__ *CljsCoreBlackNode) Redden() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return (&CljsCoreRedNode{self__.Key, self__.Val, self__.Left, self__.Right, nil})
	}
}

func (self__ *CljsCoreBlackNode) Blacken() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return node___1
	}
}

func (self__ *CljsCoreBlackNode) Add_left(ins interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Native_invoke_instance_method.X_invoke_Arity3(ins, "Balance_left", []interface{}{node___1})
	}
}

func (self__ *CljsCoreBlackNode) Replace(key___1 interface{}, val___1 interface{}, left___1 interface{}, right___1 interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return (&CljsCoreBlackNode{key___1, val___1, left___1, right___1, nil})
	}
}

func (self__ *CljsCoreBlackNode) Balance_left(parent interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), node___1, Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil})
	}
}

func (self__ *CljsCoreBlackNode) Balance_right(parent interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(parent, "Left"), node___1, nil})
	}
}

func (self__ *CljsCoreBlackNode) Remove_left(del interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Balance_left_del.X_invoke_Arity4(self__.Key, self__.Val, del, self__.Right)
	}
}

func (self__ *CljsCoreBlackNode) Kv_reduce(f interface{}, init interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Tree_map_kv_reduce.X_invoke_Arity3(node___1, f, init).(interface{})
	}
}

func (self__ *CljsCoreBlackNode) Remove_right(del interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Balance_right_del.X_invoke_Arity4(self__.Key, self__.Val, self__.Left, del)
	}
}

func (_ *CljsCoreBlackNode) CljsCoreILookup__() {}
func (self__ *CljsCoreBlackNode) X_lookup_Arity2(k interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return node___1.X_nth_Arity3(k, nil).(interface{})
	}
}

func (self__ *CljsCoreBlackNode) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return node___1.X_nth_Arity3(k, not_found).(interface{})
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIIndexed__() {}
func (self__ *CljsCoreBlackNode) X_nth_Arity2(n interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		if n.(float64) == float64(0) {
			return self__.Key
		} else {
			if n.(float64) == float64(1) {
				return self__.Val
			} else {
				return nil

			}
		}
	}
}

func (self__ *CljsCoreBlackNode) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		if n.(float64) == float64(0) {
			return self__.Key
		} else {
			if n.(float64) == float64(1) {
				return self__.Val
			} else {
				return not_found

			}
		}
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIVector__() {}
func (self__ *CljsCoreBlackNode) X_assoc_n_Arity3(n interface{}, v interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return _assoc_n.X_invoke_Arity3(&CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{self__.Key, self__.Val}, nil}, n, v)
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIMeta__() {}
func (self__ *CljsCoreBlackNode) X_meta_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return nil
	}
}

func (_ *CljsCoreBlackNode) CljsCoreICounted__() {}
func (self__ *CljsCoreBlackNode) X_count_Arity1() float64 {
	{
		var node___1 = self__
		_ = node___1
		return float64(2)
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIMapEntry__() {}
func (self__ *CljsCoreBlackNode) X_key_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return self__.Key
	}
}

func (self__ *CljsCoreBlackNode) X_val_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return self__.Val
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIStack__() {}
func (self__ *CljsCoreBlackNode) X_peek_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return self__.Val
	}
}

func (self__ *CljsCoreBlackNode) X_pop_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return &CljsCorePersistentVector{nil, 1, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{self__.Key}, nil}
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIHash__() {}
func (self__ *CljsCoreBlackNode) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIEquiv__() {}
func (self__ *CljsCoreBlackNode) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreBlackNode) X_empty_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return CljsCorePersistentVector_EMPTY
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIReduce__() {}
func (self__ *CljsCoreBlackNode) X_reduce_Arity2(f interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Ci_reduce.X_invoke_Arity2(node___1, f).(interface{})
	}
}

func (self__ *CljsCoreBlackNode) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Ci_reduce.X_invoke_Arity3(node___1, f, start)
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIAssociative__() {}
func (self__ *CljsCoreBlackNode) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Assoc.X_invoke_Arity3(&CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{self__.Key, self__.Val}, nil}, k, v).(interface{})
	}
}

func (_ *CljsCoreBlackNode) CljsCoreISeqable__() {}
func (self__ *CljsCoreBlackNode) X_seq_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return CljsCoreList_EMPTY.(CljsCoreICollection).X_conj_Arity2(self__.Val).X_conj_Arity2(self__.Key)
	}
}

func (_ *CljsCoreBlackNode) CljsCoreISequential__() {}
func (_ *CljsCoreBlackNode) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreBlackNode) X_with_meta_Arity2(meta interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return With_meta.X_invoke_Arity2(&CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{self__.Key, self__.Val}, nil}, meta)
	}
}

func (_ *CljsCoreBlackNode) CljsCoreICollection__() {}
func (self__ *CljsCoreBlackNode) X_conj_Arity2(o interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return &CljsCorePersistentVector{nil, 3, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{self__.Key, self__.Val, o}, nil}
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIFn__() {}
func (self__ *CljsCoreBlackNode) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity1(k interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return node___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return node___1.X_lookup_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCoreBlackNode) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

var X__GT_BlackNode = func(__GT_BlackNode *AFn) *AFn {
	return Fn(__GT_BlackNode, func(key interface{}, val interface{}, left interface{}, right interface{}, __hash interface{}) interface{} {
		return (&CljsCoreBlackNode{key, val, left, right, __hash})
	})
}(&AFn{})

type CljsCoreRedNode struct {
	Key     interface{}
	Val     interface{}
	Left    interface{}
	Right   interface{}
	X__hash interface{}
}

func (_ *CljsCoreRedNode) CljsCoreObject__() {}
func (self__ *CljsCoreRedNode) Add_right(ins interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return (&CljsCoreRedNode{self__.Key, self__.Val, self__.Left, ins, nil})
	}
}

func (self__ *CljsCoreRedNode) Redden() interface{} {
	{
		var node___1 = self__
		_ = node___1
		panic((&js.Error{"red-black tree invariant violation"}))
	}
}

func (self__ *CljsCoreRedNode) Blacken() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return (&CljsCoreBlackNode{self__.Key, self__.Val, self__.Left, self__.Right, nil})
	}
}

func (self__ *CljsCoreRedNode) Add_left(ins interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return (&CljsCoreRedNode{self__.Key, self__.Val, ins, self__.Right, nil})
	}
}

func (self__ *CljsCoreRedNode) Replace(key___1 interface{}, val___1 interface{}, left___1 interface{}, right___1 interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return (&CljsCoreRedNode{key___1, val___1, left___1, right___1, nil})
	}
}

func (self__ *CljsCoreRedNode) Balance_left(parent interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		if func() bool { _, instanceof := self__.Left.(*CljsCoreRedNode); return instanceof }() {
			return (&CljsCoreRedNode{self__.Key, self__.Val, Native_invoke_instance_method.X_invoke_Arity3(self__.Left, "Blacken", []interface{}{}), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), self__.Right, Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil}), nil})
		} else {
			if func() bool { _, instanceof := self__.Right.(*CljsCoreRedNode); return instanceof }() {
				return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(self__.Right, "Key"), Native_get_instance_field.X_invoke_Arity2(self__.Right, "Val"), (&CljsCoreBlackNode{self__.Key, self__.Val, self__.Left, Native_get_instance_field.X_invoke_Arity2(self__.Right, "Left"), nil}), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(self__.Right, "Right"), Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil}), nil})
			} else {
				return (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), node___1, Native_get_instance_field.X_invoke_Arity2(parent, "Right"), nil})

			}
		}
	}
}

func (self__ *CljsCoreRedNode) Balance_right(parent interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		if func() bool { _, instanceof := self__.Right.(*CljsCoreRedNode); return instanceof }() {
			return (&CljsCoreRedNode{self__.Key, self__.Val, (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(parent, "Left"), self__.Left, nil}), Native_invoke_instance_method.X_invoke_Arity3(self__.Right, "Blacken", []interface{}{}), nil})
		} else {
			if func() bool { _, instanceof := self__.Left.(*CljsCoreRedNode); return instanceof }() {
				return (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(self__.Left, "Key"), Native_get_instance_field.X_invoke_Arity2(self__.Left, "Val"), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(parent, "Left"), Native_get_instance_field.X_invoke_Arity2(self__.Left, "Left"), nil}), (&CljsCoreBlackNode{self__.Key, self__.Val, Native_get_instance_field.X_invoke_Arity2(self__.Left, "Right"), self__.Right, nil}), nil})
			} else {
				return (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(parent, "Key"), Native_get_instance_field.X_invoke_Arity2(parent, "Val"), Native_get_instance_field.X_invoke_Arity2(parent, "Left"), node___1, nil})

			}
		}
	}
}

func (self__ *CljsCoreRedNode) Remove_left(del interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return (&CljsCoreRedNode{self__.Key, self__.Val, del, self__.Right, nil})
	}
}

func (self__ *CljsCoreRedNode) Kv_reduce(f interface{}, init interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Tree_map_kv_reduce.X_invoke_Arity3(node___1, f, init).(interface{})
	}
}

func (self__ *CljsCoreRedNode) Remove_right(del interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return (&CljsCoreRedNode{self__.Key, self__.Val, self__.Left, del, nil})
	}
}

func (_ *CljsCoreRedNode) CljsCoreILookup__() {}
func (self__ *CljsCoreRedNode) X_lookup_Arity2(k interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return node___1.X_nth_Arity3(k, nil).(interface{})
	}
}

func (self__ *CljsCoreRedNode) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return node___1.X_nth_Arity3(k, not_found).(interface{})
	}
}

func (_ *CljsCoreRedNode) CljsCoreIIndexed__() {}
func (self__ *CljsCoreRedNode) X_nth_Arity2(n interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		if n.(float64) == float64(0) {
			return self__.Key
		} else {
			if n.(float64) == float64(1) {
				return self__.Val
			} else {
				return nil

			}
		}
	}
}

func (self__ *CljsCoreRedNode) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		if n.(float64) == float64(0) {
			return self__.Key
		} else {
			if n.(float64) == float64(1) {
				return self__.Val
			} else {
				return not_found

			}
		}
	}
}

func (_ *CljsCoreRedNode) CljsCoreIVector__() {}
func (self__ *CljsCoreRedNode) X_assoc_n_Arity3(n interface{}, v interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return _assoc_n.X_invoke_Arity3(&CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{self__.Key, self__.Val}, nil}, n, v)
	}
}

func (_ *CljsCoreRedNode) CljsCoreIMeta__() {}
func (self__ *CljsCoreRedNode) X_meta_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return nil
	}
}

func (_ *CljsCoreRedNode) CljsCoreICounted__() {}
func (self__ *CljsCoreRedNode) X_count_Arity1() float64 {
	{
		var node___1 = self__
		_ = node___1
		return float64(2)
	}
}

func (_ *CljsCoreRedNode) CljsCoreIMapEntry__() {}
func (self__ *CljsCoreRedNode) X_key_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return self__.Key
	}
}

func (self__ *CljsCoreRedNode) X_val_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return self__.Val
	}
}

func (_ *CljsCoreRedNode) CljsCoreIStack__() {}
func (self__ *CljsCoreRedNode) X_peek_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return self__.Val
	}
}

func (self__ *CljsCoreRedNode) X_pop_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return &CljsCorePersistentVector{nil, 1, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{self__.Key}, nil}
	}
}

func (_ *CljsCoreRedNode) CljsCoreIHash__() {}
func (self__ *CljsCoreRedNode) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreRedNode) CljsCoreIEquiv__() {}
func (self__ *CljsCoreRedNode) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreRedNode) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreRedNode) X_empty_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return CljsCorePersistentVector_EMPTY
	}
}

func (_ *CljsCoreRedNode) CljsCoreIReduce__() {}
func (self__ *CljsCoreRedNode) X_reduce_Arity2(f interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Ci_reduce.X_invoke_Arity2(node___1, f).(interface{})
	}
}

func (self__ *CljsCoreRedNode) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Ci_reduce.X_invoke_Arity3(node___1, f, start)
	}
}

func (_ *CljsCoreRedNode) CljsCoreIAssociative__() {}
func (self__ *CljsCoreRedNode) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return Assoc.X_invoke_Arity3(&CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{self__.Key, self__.Val}, nil}, k, v).(interface{})
	}
}

func (_ *CljsCoreRedNode) CljsCoreISeqable__() {}
func (self__ *CljsCoreRedNode) X_seq_Arity1() interface{} {
	{
		var node___1 = self__
		_ = node___1
		return CljsCoreList_EMPTY.(CljsCoreICollection).X_conj_Arity2(self__.Val).X_conj_Arity2(self__.Key)
	}
}

func (_ *CljsCoreRedNode) CljsCoreISequential__() {}
func (_ *CljsCoreRedNode) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreRedNode) X_with_meta_Arity2(meta interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return With_meta.X_invoke_Arity2(&CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{self__.Key, self__.Val}, nil}, meta)
	}
}

func (_ *CljsCoreRedNode) CljsCoreICollection__() {}
func (self__ *CljsCoreRedNode) X_conj_Arity2(o interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return &CljsCorePersistentVector{nil, 3, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{self__.Key, self__.Val, o}, nil}
	}
}

func (_ *CljsCoreRedNode) CljsCoreIFn__() {}
func (self__ *CljsCoreRedNode) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity1(k interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return node___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var node___1 = self__
		_ = node___1
		return node___1.X_lookup_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCoreRedNode) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

var X__GT_RedNode = func(__GT_RedNode *AFn) *AFn {
	return Fn(__GT_RedNode, func(key interface{}, val interface{}, left interface{}, right interface{}, __hash interface{}) interface{} {
		return (&CljsCoreRedNode{key, val, left, right, __hash})
	})
}(&AFn{})

var Tree_map_add = func(tree_map_add *AFn) *AFn {
	return Fn(tree_map_add, func(comp interface{}, tree interface{}, k interface{}, v interface{}, found interface{}) interface{} {
		if tree == nil {
			return (&CljsCoreRedNode{k, v, nil, nil, nil})
		} else {
			{
				var c = comp.(CljsCoreIFn).X_invoke_Arity2(k, Native_get_instance_field.X_invoke_Arity2(tree, "Key")).(interface{})
				_ = c
				if c.(float64) == float64(0) {
					found.([]interface{})[int(float64(0))] = tree
					return nil
				} else {
					if c.(float64) < float64(0) {
						{
							var ins = tree_map_add.X_invoke_Arity5(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Left"), k, v, found)
							_ = ins
							if !(ins == nil) {
								return Native_invoke_instance_method.X_invoke_Arity3(tree, "Add_left", []interface{}{ins})
							} else {
								return nil
							}
						}
					} else {
						{
							var ins = tree_map_add.X_invoke_Arity5(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Right"), k, v, found)
							_ = ins
							if !(ins == nil) {
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

var Tree_map_append = func(tree_map_append *AFn) *AFn {
	return Fn(tree_map_append, func(left interface{}, right interface{}) interface{} {
		if left == nil {
			return right
		} else {
			if right == nil {
				return left
			} else {
				if func() bool { _, instanceof := left.(*CljsCoreRedNode); return instanceof }() {
					if func() bool { _, instanceof := right.(*CljsCoreRedNode); return instanceof }() {
						{
							var app = tree_map_append.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(left, "Right"), Native_get_instance_field.X_invoke_Arity2(right, "Left"))
							_ = app
							if func() bool { _, instanceof := app.(*CljsCoreRedNode); return instanceof }() {
								return (&CljsCoreRedNode{app.Key, app.Val, (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_get_instance_field.X_invoke_Arity2(left, "Left"), app.Left, nil}), (&CljsCoreRedNode{Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), app.Right, Native_get_instance_field.X_invoke_Arity2(right, "Right"), nil}), nil})
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
								return (&CljsCoreRedNode{app.Key, app.Val, (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(left, "Key"), Native_get_instance_field.X_invoke_Arity2(left, "Val"), Native_get_instance_field.X_invoke_Arity2(left, "Left"), app.Left, nil}), (&CljsCoreBlackNode{Native_get_instance_field.X_invoke_Arity2(right, "Key"), Native_get_instance_field.X_invoke_Arity2(right, "Val"), app.Right, Native_get_instance_field.X_invoke_Arity2(right, "Right"), nil}), nil})
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

var Tree_map_remove = func(tree_map_remove *AFn) *AFn {
	return Fn(tree_map_remove, func(comp interface{}, tree interface{}, k interface{}, found interface{}) interface{} {
		if !(tree == nil) {
			{
				var c = comp.(CljsCoreIFn).X_invoke_Arity2(k, Native_get_instance_field.X_invoke_Arity2(tree, "Key")).(interface{})
				_ = c
				if c.(float64) == float64(0) {
					found.([]interface{})[int(float64(0))] = tree
					return Tree_map_append.X_invoke_Arity2(Native_get_instance_field.X_invoke_Arity2(tree, "Left"), Native_get_instance_field.X_invoke_Arity2(tree, "Right"))
				} else {
					if c.(float64) < float64(0) {
						{
							var del = tree_map_remove.X_invoke_Arity4(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Left"), k, found)
							_ = del
							if (!(del == nil)) || (!((found.([]interface{})[int(float64(0))]) == nil)) {
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
							if (!(del == nil)) || (!((found.([]interface{})[int(float64(0))]) == nil)) {
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

var Tree_map_replace = func(tree_map_replace *AFn) *AFn {
	return Fn(tree_map_replace, func(comp interface{}, tree interface{}, k interface{}, v interface{}) interface{} {
		{
			var tk = Native_get_instance_field.X_invoke_Arity2(tree, "Key")
			var c = comp.(CljsCoreIFn).X_invoke_Arity2(k, tk).(interface{})
			_, _ = tk, c
			if c.(float64) == float64(0) {
				return Native_invoke_instance_method.X_invoke_Arity3(tree, "Replace", []interface{}{tk, v, Native_get_instance_field.X_invoke_Arity2(tree, "Left"), Native_get_instance_field.X_invoke_Arity2(tree, "Right")})
			} else {
				if c.(float64) < float64(0) {
					return Native_invoke_instance_method.X_invoke_Arity3(tree, "Replace", []interface{}{tk, Native_get_instance_field.X_invoke_Arity2(tree, "Val"), tree_map_replace.X_invoke_Arity4(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Left"), k, v).(interface{}), Native_get_instance_field.X_invoke_Arity2(tree, "Right")})
				} else {
					return Native_invoke_instance_method.X_invoke_Arity3(tree, "Replace", []interface{}{tk, Native_get_instance_field.X_invoke_Arity2(tree, "Val"), Native_get_instance_field.X_invoke_Arity2(tree, "Left"), tree_map_replace.X_invoke_Arity4(comp, Native_get_instance_field.X_invoke_Arity2(tree, "Right"), k, v).(interface{})})

				}
			}
		}
	})
}(&AFn{})

type CljsCorePersistentTreeMap struct {
	Comp    interface{}
	Tree    interface{}
	Cnt     interface{}
	Meta    interface{}
	X__hash interface{}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreObject__() {}
func (self__ *CljsCorePersistentTreeMap) ForEach(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var seq__860 = Seq.X_invoke_Arity1(coll___1)
			var chunk__861 = nil
			var count__862 = float64(0)
			var i__863 = float64(0)
			_, _, _, _ = seq__860, chunk__861, count__862, i__863
			for {
				if i__863 < count__862 {
					{
						var vec__864 = chunk__861.X_nth_Arity2(i__863).(interface{})
						var k = Nth.X_invoke_Arity3(vec__864, float64(0), nil)
						var v = Nth.X_invoke_Arity3(vec__864, float64(1), nil)
						_, _, _ = vec__864, k, v
						f.(CljsCoreIFn).X_invoke_Arity2(v, k).(interface{})
						{
							var G__866 = seq__860
							var G__867 = chunk__861
							var G__868 = count__862
							var G__869 = (i__863 + float64(1))
							seq__860 = G__866
							chunk__861 = G__867
							count__862 = G__868
							i__863 = G__869
							continue
						}
					}
				} else {
					{
						var temp__4219__auto__ = Seq.X_invoke_Arity1(seq__860)
						_ = temp__4219__auto__
						if Truth_(temp__4219__auto__) {
							{
								var seq__860___1 = temp__4219__auto__
								_ = seq__860___1
								if Chunked_seq_QMARK_.X_invoke_Arity1(seq__860___1) {
									{
										var c__77428__auto__ = Chunk_first.X_invoke_Arity1(seq__860___1).(interface{})
										_ = c__77428__auto__
										{
											var G__870 = Chunk_rest.X_invoke_Arity1(seq__860___1).(interface{})
											var G__871 = c__77428__auto__
											var G__872 = Count.X_invoke_Arity1(c__77428__auto__).(float64)
											var G__873 = float64(0)
											seq__860 = G__870
											chunk__861 = G__871
											count__862 = G__872
											i__863 = G__873
											continue
										}
									}
								} else {
									{
										var vec__865 = First.X_invoke_Arity1(seq__860___1)
										var k = Nth.X_invoke_Arity3(vec__865, float64(0), nil)
										var v = Nth.X_invoke_Arity3(vec__865, float64(1), nil)
										_, _, _ = vec__865, k, v
										f.(CljsCoreIFn).X_invoke_Arity2(v, k).(interface{})
										{
											var G__874 = Next.X_invoke_Arity1(seq__860___1)
											var G__875 = nil
											var G__876 = float64(0)
											var G__877 = float64(0)
											seq__860 = G__874
											chunk__861 = G__875
											count__862 = G__876
											i__863 = G__877
											continue
										}
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
}

func (self__ *CljsCorePersistentTreeMap) Get(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCorePersistentTreeMap) Entries() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Entries_iterator.X_invoke_Arity1(Seq.X_invoke_Arity1(coll___1)).(*CljsCoreEntriesIterator)
	}
}

func (self__ *CljsCorePersistentTreeMap) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCorePersistentTreeMap) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCorePersistentTreeMap) Keys() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Iterator.X_invoke_Arity1(keys.X_invoke_Arity1(coll___1).(*CljsCoreIterator)).(*CljsCoreIterator)
	}
}

func (self__ *CljsCorePersistentTreeMap) Values() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Iterator.X_invoke_Arity1(Vals.X_invoke_Arity1(coll___1).(interface{})).(*CljsCoreIterator)
	}
}

func (self__ *CljsCorePersistentTreeMap) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (self__ *CljsCorePersistentTreeMap) Entry_at(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var t = self__.Tree
			_ = t
			for {
				if !(t == nil) {
					{
						var c = self__.Comp.(CljsCoreIFn).X_invoke_Arity2(k, Native_get_instance_field.X_invoke_Arity2(t, "Key")).(interface{})
						_ = c
						if c.(float64) == float64(0) {
							return t
						} else {
							if c.(float64) < float64(0) {
								{
									var G__878 = Native_get_instance_field.X_invoke_Arity2(t, "Left")
									t = G__878
									continue
								}
							} else {
								{
									var G__879 = Native_get_instance_field.X_invoke_Arity2(t, "Right")
									t = G__879
									continue
								}

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

func (self__ *CljsCorePersistentTreeMap) Has(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Contains_QMARK_.X_invoke_Arity2(coll___1, k)
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreILookup__() {}
func (self__ *CljsCorePersistentTreeMap) X_lookup_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return _lookup.X_invoke_Arity3(coll___1, k, nil).(interface{})
	}
}

func (self__ *CljsCorePersistentTreeMap) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var n = coll___1.Entry_at(k)
			_ = n
			if !(n == nil) {
				return Native_get_instance_field.X_invoke_Arity2(n, "Val")
			} else {
				return not_found
			}
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIKVReduce__() {}
func (self__ *CljsCorePersistentTreeMap) X_kv_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if !(self__.Tree == nil) {
			return Tree_map_kv_reduce.X_invoke_Arity3(self__.Tree, f, init).(interface{})
		} else {
			return init
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIMeta__() {}
func (self__ *CljsCorePersistentTreeMap) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreICloneable__() {}
func (self__ *CljsCorePersistentTreeMap) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCorePersistentTreeMap{self__.Comp, self__.Tree, self__.Cnt, self__.Meta, self__.X__hash})
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreICounted__() {}
func (self__ *CljsCorePersistentTreeMap) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Cnt.(float64)
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIReversible__() {}
func (self__ *CljsCorePersistentTreeMap) X_rseq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Cnt.(float64) > float64(0) {
			return Create_tree_map_seq.X_invoke_Arity3(self__.Tree, false, self__.Cnt).(*CljsCorePersistentTreeMapSeq)
		} else {
			return nil
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIHash__() {}
func (self__ *CljsCorePersistentTreeMap) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_unordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIEquiv__() {}
func (self__ *CljsCorePersistentTreeMap) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_map.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCorePersistentTreeMap) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCorePersistentTreeMap_EMPTY, self__.Meta)
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIMap__() {}
func (self__ *CljsCorePersistentTreeMap) X_dissoc_Arity2(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var found = []interface{}{nil}
			var t = Tree_map_remove.X_invoke_Arity4(self__.Comp, self__.Tree, k, found)
			_, _ = found, t
			if t == nil {
				if Nth.X_invoke_Arity2(found, float64(0)) == nil {
					return coll___1
				} else {
					return (&CljsCorePersistentTreeMap{self__.Comp, nil, float64(0), self__.Meta, nil})
				}
			} else {
				return (&CljsCorePersistentTreeMap{self__.Comp, t.Blacken(), (self__.Cnt.(float64) - float64(1)), self__.Meta, nil})
			}
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIAssociative__() {}
func (self__ *CljsCorePersistentTreeMap) X_assoc_Arity3(k interface{}, v interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var found = []interface{}{nil}
			var t = Tree_map_add.X_invoke_Arity5(self__.Comp, self__.Tree, k, v, found)
			_, _ = found, t
			if t == nil {
				{
					var found_node = Nth.X_invoke_Arity2(found, float64(0))
					_ = found_node
					if X_EQ_.Arity2IIB(v, found_node.Val) {
						return coll___1
					} else {
						return (&CljsCorePersistentTreeMap{self__.Comp, Tree_map_replace.X_invoke_Arity4(self__.Comp, self__.Tree, k, v).(interface{}), self__.Cnt, self__.Meta, nil})
					}
				}
			} else {
				return (&CljsCorePersistentTreeMap{self__.Comp, t.Blacken(), (self__.Cnt.(float64) + float64(1)), self__.Meta, nil})
			}
		}
	}
}

func (self__ *CljsCorePersistentTreeMap) X_contains_key_QMARK__Arity2(k interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return !(coll___1.Entry_at(k) == nil)
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreISeqable__() {}
func (self__ *CljsCorePersistentTreeMap) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Cnt.(float64) > float64(0) {
			return Create_tree_map_seq.X_invoke_Arity3(self__.Tree, true, self__.Cnt).(*CljsCorePersistentTreeMapSeq)
		} else {
			return nil
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIWithMeta__() {}
func (self__ *CljsCorePersistentTreeMap) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentTreeMap{self__.Comp, self__.Tree, self__.Cnt, meta___1, self__.X__hash})
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreICollection__() {}
func (self__ *CljsCorePersistentTreeMap) X_conj_Arity2(entry interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if Vector_QMARK_.Arity1IB(entry) {
			return coll___1.X_assoc_Arity3(entry.(CljsCoreIIndexed).X_nth_Arity2(float64(0)).(interface{}), entry.(CljsCoreIIndexed).X_nth_Arity2(float64(1)).(interface{}))
		} else {
			{
				var ret = coll___1
				var es = Seq.Arity1IQ(entry)
				_, _ = ret, es
				for {
					if es == nil {
						return ret
					} else {
						{
							var e = First.X_invoke_Arity1(es)
							_ = e
							if Vector_QMARK_.X_invoke_Arity1(e) {
								{
									var G__880 = ret.X_assoc_Arity3(e.X_nth_Arity2(float64(0)).(interface{}), e.X_nth_Arity2(float64(1)).(interface{}))
									var G__881 = Next.X_invoke_Arity1(es)
									ret = G__880
									es = G__881
									continue
								}
							} else {
								panic((&js.Error{"conj on a map takes map entries or seqables of map entries"}))
							}
						}
					}
				}
			}
		}
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIFn__() {}
func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity1(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCorePersistentTreeMap) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreISorted__() {}
func (self__ *CljsCorePersistentTreeMap) X_sorted_seq_Arity2(ascending_QMARK_ interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Cnt.(float64) > float64(0) {
			return Create_tree_map_seq.X_invoke_Arity3(self__.Tree, ascending_QMARK_, self__.Cnt).(*CljsCorePersistentTreeMapSeq)
		} else {
			return nil
		}
	}
}

func (self__ *CljsCorePersistentTreeMap) X_sorted_seq_from_Arity3(k interface{}, ascending_QMARK_ interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Cnt.(float64) > float64(0) {
			{
				var stack = nil
				var t = self__.Tree
				_, _ = stack, t
				for {
					if !(t == nil) {
						{
							var c = self__.Comp.(CljsCoreIFn).X_invoke_Arity2(k, Native_get_instance_field.X_invoke_Arity2(t, "Key")).(interface{})
							_ = c
							if c.(float64) == float64(0) {
								return (&CljsCorePersistentTreeMapSeq{nil, Conj.X_invoke_Arity2(stack, t).(interface{}), ascending_QMARK_, float64(-1), nil})
							} else {
								if Truth_(ascending_QMARK_) {
									if c.(float64) < float64(0) {
										{
											var G__882 = Conj.X_invoke_Arity2(stack, t).(interface{})
											var G__883 = Native_get_instance_field.X_invoke_Arity2(t, "Left")
											stack = G__882
											t = G__883
											continue
										}
									} else {
										{
											var G__884 = stack
											var G__885 = Native_get_instance_field.X_invoke_Arity2(t, "Right")
											stack = G__884
											t = G__885
											continue
										}
									}
								} else {
									if c.(float64) > float64(0) {
										{
											var G__886 = Conj.X_invoke_Arity2(stack, t).(interface{})
											var G__887 = Native_get_instance_field.X_invoke_Arity2(t, "Right")
											stack = G__886
											t = G__887
											continue
										}
									} else {
										{
											var G__888 = stack
											var G__889 = Native_get_instance_field.X_invoke_Arity2(t, "Left")
											stack = G__888
											t = G__889
											continue
										}
									}

								}
							}
						}
					} else {
						if stack == nil {
							return nil
						} else {
							return (&CljsCorePersistentTreeMapSeq{nil, stack, ascending_QMARK_, float64(-1), nil})
						}
					}
				}
			}
		} else {
			return nil
		}
	}
}

func (self__ *CljsCorePersistentTreeMap) X_entry_key_Arity2(entry interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Key.X_invoke_Arity1(entry).(interface{})
	}
}

func (self__ *CljsCorePersistentTreeMap) X_comparator_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Comp
	}
}

var X__GT_PersistentTreeMap = func(__GT_PersistentTreeMap *AFn) *AFn {
	return Fn(__GT_PersistentTreeMap, func(comp interface{}, tree interface{}, cnt interface{}, meta interface{}, __hash interface{}) interface{} {
		return (&CljsCorePersistentTreeMap{comp, tree, cnt, meta, __hash})
	})
}(&AFn{})

var CljsCorePersistentTreeMap_EMPTY = (&CljsCorePersistentTreeMap{Compare, nil, float64(0), nil, float64(0)})

/**
* keyval => key val
* Returns a new hash map with supplied mappings.
* @param {...*} var_args
 */
var Hash_map = func(hash_map *AFn) *AFn {
	return Fn(hash_map, func(keyvals__ ...interface{}) interface{} {
		var keyvals = Array_seq.X_invoke_Arity1(keyvals__[0:])
		_ = keyvals
		{
			var in = Seq.Arity1IQ(keyvals)
			var out = Transient.X_invoke_Arity1(CljsCorePersistentHashMap_EMPTY).(interface{})
			_, _ = in, out
			for {
				if Truth_(in) {
					{
						var G__890 = Nnext.X_invoke_Arity1(in).(CljsCoreISeq)
						var G__891 = Assoc_BANG_.X_invoke_Arity3(out, First.X_invoke_Arity1(in), Second.X_invoke_Arity1(in)).(interface{})
						in = G__890
						out = G__891
						continue
					}
				} else {
					return Persistent_BANG_.X_invoke_Arity1(out).(interface{})
				}
			}
		}
	})
}(&AFn{})

/**
* keyval => key val
* Returns a new array map with supplied mappings.
* @param {...*} var_args
 */
var Array_map = func(array_map *AFn) *AFn {
	return Fn(array_map, func(keyvals__ ...interface{}) interface{} {
		var keyvals = Array_seq.X_invoke_Arity1(keyvals__[0:])
		_ = keyvals
		return (&CljsCorePersistentArrayMap{nil, Quot.X_invoke_Arity2(Count.X_invoke_Arity1(keyvals).(float64), float64(2)).(float64), Apply.X_invoke_Arity2(Array, keyvals), nil})
	})
}(&AFn{})

/**
* keyval => key val
* Returns a new object map with supplied mappings.
* @param {...*} var_args
 */
var Obj_map = func(obj_map *AFn) *AFn {
	return Fn(obj_map, func(keyvals__ ...interface{}) interface{} {
		var keyvals = Array_seq.X_invoke_Arity1(keyvals__[0:])
		_ = keyvals
		{
			var ks = []interface{}{}
			var obj = func() interface{} {
				var obj895 = map[string]interface{}{}
				_ = obj895
				return obj895
			}()
			_, _ = ks, obj
			{
				var kvs = Seq.Arity1IQ(keyvals)
				_ = kvs
				for {
					if Truth_(kvs) {
						ks.Push(First.X_invoke_Arity1(kvs))
						obj.([]interface{})[int(First.X_invoke_Arity1(kvs).(float64))] = Second.X_invoke_Arity1(kvs)
						{
							var G__896 = Nnext.X_invoke_Arity1(kvs).(CljsCoreISeq)
							kvs = G__896
							continue
						}
					} else {
						return CljsCoreObjMap_FromObject.X_invoke_Arity2(ks, obj)
					}
				}
			}
		}
	})
}(&AFn{})

/**
* keyval => key val
* Returns a new sorted map with supplied mappings.
* @param {...*} var_args
 */
var Sorted_map = func(sorted_map *AFn) *AFn {
	return Fn(sorted_map, func(keyvals__ ...interface{}) interface{} {
		var keyvals = Array_seq.X_invoke_Arity1(keyvals__[0:])
		_ = keyvals
		{
			var in = Seq.Arity1IQ(keyvals)
			var out = CljsCorePersistentTreeMap_EMPTY
			_, _ = in, out
			for {
				if Truth_(in) {
					{
						var G__897 = Nnext.X_invoke_Arity1(in).(CljsCoreISeq)
						var G__898 = Assoc.X_invoke_Arity3(out, First.X_invoke_Arity1(in), Second.X_invoke_Arity1(in)).(interface{})
						in = G__897
						out = G__898
						continue
					}
				} else {
					return out
				}
			}
		}
	})
}(&AFn{})

/**
* keyval => key val
* Returns a new sorted map with supplied mappings, using the supplied comparator.
* @param {...*} var_args
 */
var Sorted_map_by = func(sorted_map_by *AFn) *AFn {
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
					{
						var G__899 = Nnext.X_invoke_Arity1(in).(CljsCoreISeq)
						var G__900 = Assoc.X_invoke_Arity3(out, First.X_invoke_Arity1(in), Second.X_invoke_Arity1(in)).(interface{})
						in = G__899
						out = G__900
						continue
					}
				} else {
					return out
				}
			}
		}
	})
}(&AFn{})

type CljsCoreKeySeq struct {
	Mseq   interface{}
	X_meta interface{}
}

func (_ *CljsCoreKeySeq) CljsCoreObject__() {}
func (self__ *CljsCoreKeySeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreKeySeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreKeySeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreKeySeq) CljsCoreIMeta__() {}
func (self__ *CljsCoreKeySeq) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.X_meta
	}
}

func (_ *CljsCoreKeySeq) CljsCoreINext__() {}
func (self__ *CljsCoreKeySeq) X_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var nseq = func() interface{} {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INext", Str: "cljs.core/INext", X_hash: float64(-113000046), X_meta: nil}), self__.Mseq).(bool)) {
					return _next.X_invoke_Arity1(self__.Mseq)
				} else {
					return Next.X_invoke_Arity1(self__.Mseq)
				}
			}()
			_ = nseq
			if nseq == nil {
				return nil
			} else {
				return (&CljsCoreKeySeq{nseq, self__.X_meta})
			}
		}
	}
}

func (_ *CljsCoreKeySeq) CljsCoreIHash__() {}
func (self__ *CljsCoreKeySeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Hash_ordered_coll.X_invoke_Arity1(coll___1)
	}
}

func (_ *CljsCoreKeySeq) CljsCoreIEquiv__() {}
func (self__ *CljsCoreKeySeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreKeySeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreKeySeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.X_meta)
	}
}

func (_ *CljsCoreKeySeq) CljsCoreIReduce__() {}
func (self__ *CljsCoreKeySeq) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCoreKeySeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCoreKeySeq) CljsCoreISeq__() {}
func (self__ *CljsCoreKeySeq) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var me = _first.X_invoke_Arity1(self__.Mseq).(interface{})
			_ = me
			return me.X_key_Arity1().(interface{})
		}
	}
}

func (self__ *CljsCoreKeySeq) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var nseq = func() interface{} {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INext", Str: "cljs.core/INext", X_hash: float64(-113000046), X_meta: nil}), self__.Mseq).(bool)) {
					return self__.Mseq.X_next_Arity1()
				} else {
					return Next.X_invoke_Arity1(self__.Mseq)
				}
			}()
			_ = nseq
			if !(nseq == nil) {
				return (&CljsCoreKeySeq{nseq, self__.X_meta})
			} else {
				return CljsCoreList_EMPTY
			}
		}
	}
}

func (_ *CljsCoreKeySeq) CljsCoreISeqable__() {}
func (self__ *CljsCoreKeySeq) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1
	}
}

func (_ *CljsCoreKeySeq) CljsCoreISequential__() {}
func (_ *CljsCoreKeySeq) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreKeySeq) X_with_meta_Arity2(new_meta interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreKeySeq{self__.Mseq, new_meta})
	}
}

func (_ *CljsCoreKeySeq) CljsCoreICollection__() {}
func (self__ *CljsCoreKeySeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(*CljsCoreCons)
	}
}

var X__GT_KeySeq = func(__GT_KeySeq *AFn) *AFn {
	return Fn(__GT_KeySeq, func(mseq interface{}, _meta interface{}) interface{} {
		return (&CljsCoreKeySeq{mseq, _meta})
	})
}(&AFn{})

/**
* Returns a sequence of the map's keys.
 */
var Keys = func(keys *AFn) *AFn {
	return Fn(keys, func(hash_map interface{}) interface{} {
		{
			var temp__4219__auto__ = Seq.Arity1IQ(hash_map)
			_ = temp__4219__auto__
			if Truth_(temp__4219__auto__) {
				{
					var mseq = temp__4219__auto__
					_ = mseq
					return (&CljsCoreKeySeq{mseq, nil})
				}
			} else {
				return nil
			}
		}
	})
}(&AFn{})

/**
* Returns the key of the map entry.
 */
var Key = func(key *AFn) *AFn {
	return Fn(key, func(map_entry interface{}) interface{} {
		return map_entry.(CljsCoreIMapEntry).X_key_Arity1().(interface{})
	})
}(&AFn{})

type CljsCoreValSeq struct {
	Mseq   interface{}
	X_meta interface{}
}

func (_ *CljsCoreValSeq) CljsCoreObject__() {}
func (self__ *CljsCoreValSeq) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreValSeq) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreValSeq) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreValSeq) CljsCoreIMeta__() {}
func (self__ *CljsCoreValSeq) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.X_meta
	}
}

func (_ *CljsCoreValSeq) CljsCoreINext__() {}
func (self__ *CljsCoreValSeq) X_next_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var nseq = func() interface{} {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INext", Str: "cljs.core/INext", X_hash: float64(-113000046), X_meta: nil}), self__.Mseq).(bool)) {
					return _next.X_invoke_Arity1(self__.Mseq)
				} else {
					return Next.X_invoke_Arity1(self__.Mseq)
				}
			}()
			_ = nseq
			if nseq == nil {
				return nil
			} else {
				return (&CljsCoreValSeq{nseq, self__.X_meta})
			}
		}
	}
}

func (_ *CljsCoreValSeq) CljsCoreIHash__() {}
func (self__ *CljsCoreValSeq) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Hash_ordered_coll.X_invoke_Arity1(coll___1)
	}
}

func (_ *CljsCoreValSeq) CljsCoreIEquiv__() {}
func (self__ *CljsCoreValSeq) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return Equiv_sequential.X_invoke_Arity2(coll___1, other).(bool)
	}
}

func (_ *CljsCoreValSeq) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreValSeq) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.X_meta)
	}
}

func (_ *CljsCoreValSeq) CljsCoreIReduce__() {}
func (self__ *CljsCoreValSeq) X_reduce_Arity2(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity2(f, coll___1).(interface{})
	}
}

func (self__ *CljsCoreValSeq) X_reduce_Arity3(f interface{}, start interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Seq_reduce.X_invoke_Arity3(f, start, coll___1)
	}
}

func (_ *CljsCoreValSeq) CljsCoreISeq__() {}
func (self__ *CljsCoreValSeq) X_first_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var me = _first.X_invoke_Arity1(self__.Mseq).(interface{})
			_ = me
			return me.X_val_Arity1().(interface{})
		}
	}
}

func (self__ *CljsCoreValSeq) X_rest_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var nseq = func() interface{} {
				if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INext", Str: "cljs.core/INext", X_hash: float64(-113000046), X_meta: nil}), self__.Mseq).(bool)) {
					return self__.Mseq.X_next_Arity1()
				} else {
					return Next.X_invoke_Arity1(self__.Mseq)
				}
			}()
			_ = nseq
			if !(nseq == nil) {
				return (&CljsCoreValSeq{nseq, self__.X_meta})
			} else {
				return CljsCoreList_EMPTY
			}
		}
	}
}

func (_ *CljsCoreValSeq) CljsCoreISeqable__() {}
func (self__ *CljsCoreValSeq) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1
	}
}

func (_ *CljsCoreValSeq) CljsCoreISequential__() {}
func (_ *CljsCoreValSeq) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreValSeq) X_with_meta_Arity2(new_meta interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreValSeq{self__.Mseq, new_meta})
	}
}

func (_ *CljsCoreValSeq) CljsCoreICollection__() {}
func (self__ *CljsCoreValSeq) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Cons.X_invoke_Arity2(o, coll___1).(*CljsCoreCons)
	}
}

var X__GT_ValSeq = func(__GT_ValSeq *AFn) *AFn {
	return Fn(__GT_ValSeq, func(mseq interface{}, _meta interface{}) interface{} {
		return (&CljsCoreValSeq{mseq, _meta})
	})
}(&AFn{})

/**
* Returns a sequence of the map's values.
 */
var Vals = func(vals *AFn) *AFn {
	return Fn(vals, func(hash_map interface{}) interface{} {
		{
			var temp__4219__auto__ = Seq.Arity1IQ(hash_map)
			_ = temp__4219__auto__
			if Truth_(temp__4219__auto__) {
				{
					var mseq = temp__4219__auto__
					_ = mseq
					return (&CljsCoreValSeq{mseq, nil})
				}
			} else {
				return nil
			}
		}
	})
}(&AFn{})

/**
* Returns the value in the map entry.
 */
var Val = func(val *AFn) *AFn {
	return Fn(val, func(map_entry interface{}) interface{} {
		return map_entry.(CljsCoreIMapEntry).X_val_Arity1().(interface{})
	})
}(&AFn{})

/**
* Returns a map that consists of the rest of the maps conj-ed onto
* the first.  If a key occurs in more than one map, the mapping from
* the latter (left-to-right) will be the mapping in the result.
* @param {...*} var_args
 */
var Merge = func(merge *AFn) *AFn {
	return Fn(merge, func(maps__ ...interface{}) interface{} {
		var maps = Array_seq.X_invoke_Arity1(maps__[0:])
		_ = maps
		if Truth_(Some.X_invoke_Arity2(Identity, maps)) {
			return Reduce.X_invoke_Arity2(func(G__903 *AFn) *AFn {
				return Fn(G__903, func(p1__901_SHARP_ interface{}, p2__902_SHARP_ interface{}) interface{} {
					return Conj.X_invoke_Arity2(func() interface{} {
						var or__76632__auto__ = p1__901_SHARP_
						_ = or__76632__auto__
						if Truth_(or__76632__auto__) {
							return or__76632__auto__
						} else {
							return CljsCorePersistentArrayMap_EMPTY
						}
					}(), p2__902_SHARP_).(interface{})
				})
			}(&AFn{}), maps)
		} else {
			return nil
		}
	})
}(&AFn{})

/**
* Returns a map that consists of the rest of the maps conj-ed onto
* the first.  If a key occurs in more than one map, the mapping(s)
* from the latter (left-to-right) will be combined with the mapping in
* the result by calling (f val-in-result val-in-latter).
* @param {...*} var_args
 */
var Merge_with = func(merge_with *AFn) *AFn {
	return Fn(merge_with, func(f_maps__ ...interface{}) interface{} {
		var f = f_maps__[0]
		var maps = Array_seq.X_invoke_Arity1(f_maps__[1:])
		_, _ = f, maps
		if Truth_(Some.X_invoke_Arity2(Identity, maps)) {
			{
				var merge_entry = func(G__904 *AFn) *AFn {
					return Fn(G__904, func(m interface{}, e interface{}) interface{} {
						{
							var k = First.X_invoke_Arity1(e)
							var v = Second.X_invoke_Arity1(e)
							_, _ = k, v
							if Contains_QMARK_.X_invoke_Arity2(m, k) {
								return Assoc.X_invoke_Arity3(m, k, f.(CljsCoreIFn).X_invoke_Arity2(Get.X_invoke_Arity2(m, k), v).(interface{})).(interface{})
							} else {
								return Assoc.X_invoke_Arity3(m, k, v).(interface{})
							}
						}
					})
				}(&AFn{})
				var merge2 = func(merge_entry interface{}) *AFn {
					return func(G__905 *AFn) *AFn {
						return Fn(G__905, func(m1 interface{}, m2 interface{}) interface{} {
							return Reduce.X_invoke_Arity3(merge_entry, func() interface{} {
								var or__76632__auto__ = m1
								_ = or__76632__auto__
								if Truth_(or__76632__auto__) {
									return or__76632__auto__
								} else {
									return CljsCorePersistentArrayMap_EMPTY
								}
							}(), Seq.Arity1IQ(m2))
						})
					}(&AFn{})
				}(merge_entry)
				_, _ = merge_entry, merge2
				return Reduce.X_invoke_Arity2(merge2, maps)
			}
		} else {
			return nil
		}
	})
}(&AFn{})

/**
* Returns a map containing only those entries in map whose key is in keys
 */
var Select_keys = func(select_keys *AFn) *AFn {
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
						{
							var G__906 = func() interface{} {
								if Not_EQ_.X_invoke_Arity2(entry, (&CljsCoreKeyword{Ns: "cljs.core", Name: "not-found", Fqn: "cljs.core/not-found", X_hash: float64(-1572889185)})) {
									return Assoc.X_invoke_Arity3(ret, key, entry).(interface{})
								} else {
									return ret
								}
							}()
							var G__907 = Next.X_invoke_Arity1(keys)
							ret = G__906
							keys = G__907
							continue
						}
					}
				} else {
					return ret
				}
			}
		}
	})
}(&AFn{})

type CljsCorePersistentHashSet struct {
	Meta     interface{}
	Hash_map interface{}
	X__hash  interface{}
}

func (_ *CljsCorePersistentHashSet) CljsCoreObject__() {}
func (self__ *CljsCorePersistentHashSet) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCorePersistentHashSet) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCorePersistentHashSet) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (self__ *CljsCorePersistentHashSet) Keys() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Iterator.X_invoke_Arity1(Seq.X_invoke_Arity1(coll___1)).(*CljsCoreIterator)
	}
}

func (self__ *CljsCorePersistentHashSet) Entries() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Set_entries_iterator.X_invoke_Arity1(Seq.X_invoke_Arity1(coll___1)).(*CljsCoreSetEntriesIterator)
	}
}

func (self__ *CljsCorePersistentHashSet) Values() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Iterator.X_invoke_Arity1(Seq.X_invoke_Arity1(coll___1)).(*CljsCoreIterator)
	}
}

func (self__ *CljsCorePersistentHashSet) Has(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Contains_QMARK_.X_invoke_Arity2(coll___1, k)
	}
}

func (self__ *CljsCorePersistentHashSet) ForEach(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var seq__915 = Seq.X_invoke_Arity1(coll___1)
			var chunk__916 = nil
			var count__917 = float64(0)
			var i__918 = float64(0)
			_, _, _, _ = seq__915, chunk__916, count__917, i__918
			for {
				if i__918 < count__917 {
					{
						var vec__919 = chunk__916.X_nth_Arity2(i__918).(interface{})
						var k = Nth.X_invoke_Arity3(vec__919, float64(0), nil)
						var v = Nth.X_invoke_Arity3(vec__919, float64(1), nil)
						_, _, _ = vec__919, k, v
						f.(CljsCoreIFn).X_invoke_Arity2(v, k).(interface{})
						{
							var G__921 = seq__915
							var G__922 = chunk__916
							var G__923 = count__917
							var G__924 = (i__918 + float64(1))
							seq__915 = G__921
							chunk__916 = G__922
							count__917 = G__923
							i__918 = G__924
							continue
						}
					}
				} else {
					{
						var temp__4219__auto__ = Seq.X_invoke_Arity1(seq__915)
						_ = temp__4219__auto__
						if Truth_(temp__4219__auto__) {
							{
								var seq__915___1 = temp__4219__auto__
								_ = seq__915___1
								if Chunked_seq_QMARK_.X_invoke_Arity1(seq__915___1) {
									{
										var c__77428__auto__ = Chunk_first.X_invoke_Arity1(seq__915___1).(interface{})
										_ = c__77428__auto__
										{
											var G__925 = Chunk_rest.X_invoke_Arity1(seq__915___1).(interface{})
											var G__926 = c__77428__auto__
											var G__927 = Count.X_invoke_Arity1(c__77428__auto__).(float64)
											var G__928 = float64(0)
											seq__915 = G__925
											chunk__916 = G__926
											count__917 = G__927
											i__918 = G__928
											continue
										}
									}
								} else {
									{
										var vec__920 = First.X_invoke_Arity1(seq__915___1)
										var k = Nth.X_invoke_Arity3(vec__920, float64(0), nil)
										var v = Nth.X_invoke_Arity3(vec__920, float64(1), nil)
										_, _, _ = vec__920, k, v
										f.(CljsCoreIFn).X_invoke_Arity2(v, k).(interface{})
										{
											var G__929 = Next.X_invoke_Arity1(seq__915___1)
											var G__930 = nil
											var G__931 = float64(0)
											var G__932 = float64(0)
											seq__915 = G__929
											chunk__916 = G__930
											count__917 = G__931
											i__918 = G__932
											continue
										}
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
}

func (_ *CljsCorePersistentHashSet) CljsCoreILookup__() {}
func (self__ *CljsCorePersistentHashSet) X_lookup_Arity2(v interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return _lookup.X_invoke_Arity3(coll___1, v, nil).(interface{})
	}
}

func (self__ *CljsCorePersistentHashSet) X_lookup_Arity3(v interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if self__.Hash_map.(CljsCoreIAssociative).X_contains_key_QMARK__Arity2(v) {
			return v
		} else {
			return not_found
		}
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIMeta__() {}
func (self__ *CljsCorePersistentHashSet) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreICloneable__() {}
func (self__ *CljsCorePersistentHashSet) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCorePersistentHashSet{self__.Meta, self__.Hash_map, self__.X__hash})
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreICounted__() {}
func (self__ *CljsCorePersistentHashSet) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return _count.X_invoke_Arity1(self__.Hash_map)
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIHash__() {}
func (self__ *CljsCorePersistentHashSet) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_unordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIEquiv__() {}
func (self__ *CljsCorePersistentHashSet) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return (Set_QMARK_.Arity1IB(other)) && (Count.X_invoke_Arity1(coll___1).(float64) == Count.X_invoke_Arity1(other).(float64)) && (Every_QMARK_.X_invoke_Arity2(func(coll___1 *CljsCorePersistentHashSet) *AFn {
			return func(G__933 *AFn) *AFn {
				return Fn(G__933, func(p1__908_SHARP_ interface{}) interface{} {
					return Contains_QMARK_.X_invoke_Arity2(coll___1, p1__908_SHARP_)
				})
			}(&AFn{})
		}(coll___1), other))
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIEditableCollection__() {}
func (self__ *CljsCorePersistentHashSet) X_as_transient_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCoreTransientHashSet{_as_transient.X_invoke_Arity1(self__.Hash_map)})
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCorePersistentHashSet) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCorePersistentHashSet_EMPTY, self__.Meta)
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreISet__() {}
func (self__ *CljsCorePersistentHashSet) X_disjoin_Arity2(v interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentHashSet{self__.Meta, self__.Hash_map.(CljsCoreIMap).X_dissoc_Arity2(v), nil})
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreISeqable__() {}
func (self__ *CljsCorePersistentHashSet) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Keys.X_invoke_Arity1(self__.Hash_map)
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIWithMeta__() {}
func (self__ *CljsCorePersistentHashSet) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentHashSet{meta___1, self__.Hash_map, self__.X__hash})
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreICollection__() {}
func (self__ *CljsCorePersistentHashSet) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentHashSet{self__.Meta, Assoc.X_invoke_Arity3(self__.Hash_map, o, nil).(interface{}), nil})
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIFn__() {}
func (self__ *CljsCorePersistentHashSet) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity1(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCorePersistentHashSet) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

var X__GT_PersistentHashSet = func(__GT_PersistentHashSet *AFn) *AFn {
	return Fn(__GT_PersistentHashSet, func(meta interface{}, hash_map interface{}, __hash interface{}) interface{} {
		return (&CljsCorePersistentHashSet{meta, hash_map, __hash})
	})
}(&AFn{})

var CljsCorePersistentHashSet_EMPTY = (&CljsCorePersistentHashSet{nil, CljsCorePersistentArrayMap_EMPTY, float64(0)})

var CljsCorePersistentHashSet_FromArray = func(G__934 *AFn) *AFn {
	return Fn(G__934, func(items interface{}, no_clone bool) interface{} {
		{
			var len = float64(len(items.([]interface{})))
			_ = len
			if len <= CljsCorePersistentArrayMap_HASHMAP_THRESHOLD.(float64) {
				{
					var arr = func() interface{} {
						if no_clone {
							return items
						} else {
							return Aclone.X_invoke_Arity1(items).([]interface{})
						}
					}()
					_ = arr
					{
						var i = float64(0)
						var out = Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY).(interface{})
						_, _ = i, out
						for {
							if i < len {
								{
									var G__935 = (i + float64(1))
									var G__936 = out.X_assoc_BANG__Arity3((items.([]interface{})[int(i)]), nil)
									i = G__935
									out = G__936
									continue
								}
							} else {
								return (&CljsCorePersistentHashSet{nil, out.X_persistent_BANG__Arity1(), nil})
							}
						}
					}
				}
			} else {
				{
					var i = float64(0)
					var out = Transient.X_invoke_Arity1(CljsCorePersistentHashSet_EMPTY).(interface{})
					_, _ = i, out
					for {
						if i < len {
							{
								var G__937 = (i + float64(1))
								var G__938 = out.X_conj_BANG__Arity2((items.([]interface{})[int(i)]))
								i = G__937
								out = G__938
								continue
							}
						} else {
							return out.X_persistent_BANG__Arity1()
						}
					}
				}
			}
		}
	})
}(&AFn{})

type CljsCoreTransientHashSet struct{ Transient_map interface{} }

func (_ *CljsCoreTransientHashSet) CljsCoreIFn__() {}
func (self__ *CljsCoreTransientHashSet) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity1(k interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if self__.Transient_map.(CljsCoreILookup).X_lookup_Arity3(k, Lookup_sentinel).(interface{}) == Lookup_sentinel {
			return nil
		} else {
			return k
		}
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if self__.Transient_map.(CljsCoreILookup).X_lookup_Arity3(k, Lookup_sentinel).(interface{}) == Lookup_sentinel {
			return not_found
		} else {
			return k
		}
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCoreTransientHashSet) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

func (_ *CljsCoreTransientHashSet) CljsCoreILookup__() {}
func (self__ *CljsCoreTransientHashSet) X_lookup_Arity2(v interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		return _lookup.X_invoke_Arity3(tcoll___1, v, nil).(interface{})
	}
}

func (self__ *CljsCoreTransientHashSet) X_lookup_Arity3(v interface{}, not_found interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		if _lookup.X_invoke_Arity3(self__.Transient_map, v, Lookup_sentinel).(interface{}) == Lookup_sentinel {
			return not_found
		} else {
			return v
		}
	}
}

func (_ *CljsCoreTransientHashSet) CljsCoreICounted__() {}
func (self__ *CljsCoreTransientHashSet) X_count_Arity1() float64 {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		return Count.X_invoke_Arity1(self__.Transient_map).(float64)
	}
}

func (_ *CljsCoreTransientHashSet) CljsCoreITransientSet__() {}
func (self__ *CljsCoreTransientHashSet) X_disjoin_BANG__Arity2(v interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		self__.Transient_map = Dissoc_BANG_.X_invoke_Arity2(self__.Transient_map, v).(interface{})

		return tcoll___1
	}
}

func (_ *CljsCoreTransientHashSet) CljsCoreITransientCollection__() {}
func (self__ *CljsCoreTransientHashSet) X_conj_BANG__Arity2(o interface{}) interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		self__.Transient_map = Assoc_BANG_.X_invoke_Arity3(self__.Transient_map, o, nil).(interface{})

		return tcoll___1
	}
}

func (self__ *CljsCoreTransientHashSet) X_persistent_BANG__Arity1() interface{} {
	{
		var tcoll___1 = self__
		_ = tcoll___1
		return (&CljsCorePersistentHashSet{nil, Persistent_BANG_.X_invoke_Arity1(self__.Transient_map).(interface{}), nil})
	}
}

var X__GT_TransientHashSet = func(__GT_TransientHashSet *AFn) *AFn {
	return Fn(__GT_TransientHashSet, func(transient_map interface{}) interface{} {
		return (&CljsCoreTransientHashSet{transient_map})
	})
}(&AFn{})

type CljsCorePersistentTreeSet struct {
	Meta     interface{}
	Tree_map interface{}
	X__hash  interface{}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreObject__() {}
func (self__ *CljsCorePersistentTreeSet) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCorePersistentTreeSet) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCorePersistentTreeSet) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (self__ *CljsCorePersistentTreeSet) Keys() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Iterator.X_invoke_Arity1(Seq.X_invoke_Arity1(coll___1)).(*CljsCoreIterator)
	}
}

func (self__ *CljsCorePersistentTreeSet) Entries() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Set_entries_iterator.X_invoke_Arity1(Seq.X_invoke_Arity1(coll___1)).(*CljsCoreSetEntriesIterator)
	}
}

func (self__ *CljsCorePersistentTreeSet) Values() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Iterator.X_invoke_Arity1(Seq.X_invoke_Arity1(coll___1)).(*CljsCoreIterator)
	}
}

func (self__ *CljsCorePersistentTreeSet) Has(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Contains_QMARK_.X_invoke_Arity2(coll___1, k)
	}
}

func (self__ *CljsCorePersistentTreeSet) ForEach(f interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var seq__946 = Seq.X_invoke_Arity1(coll___1)
			var chunk__947 = nil
			var count__948 = float64(0)
			var i__949 = float64(0)
			_, _, _, _ = seq__946, chunk__947, count__948, i__949
			for {
				if i__949 < count__948 {
					{
						var vec__950 = chunk__947.X_nth_Arity2(i__949).(interface{})
						var k = Nth.X_invoke_Arity3(vec__950, float64(0), nil)
						var v = Nth.X_invoke_Arity3(vec__950, float64(1), nil)
						_, _, _ = vec__950, k, v
						f.(CljsCoreIFn).X_invoke_Arity2(v, k).(interface{})
						{
							var G__952 = seq__946
							var G__953 = chunk__947
							var G__954 = count__948
							var G__955 = (i__949 + float64(1))
							seq__946 = G__952
							chunk__947 = G__953
							count__948 = G__954
							i__949 = G__955
							continue
						}
					}
				} else {
					{
						var temp__4219__auto__ = Seq.X_invoke_Arity1(seq__946)
						_ = temp__4219__auto__
						if Truth_(temp__4219__auto__) {
							{
								var seq__946___1 = temp__4219__auto__
								_ = seq__946___1
								if Chunked_seq_QMARK_.X_invoke_Arity1(seq__946___1) {
									{
										var c__77428__auto__ = Chunk_first.X_invoke_Arity1(seq__946___1).(interface{})
										_ = c__77428__auto__
										{
											var G__956 = Chunk_rest.X_invoke_Arity1(seq__946___1).(interface{})
											var G__957 = c__77428__auto__
											var G__958 = Count.X_invoke_Arity1(c__77428__auto__).(float64)
											var G__959 = float64(0)
											seq__946 = G__956
											chunk__947 = G__957
											count__948 = G__958
											i__949 = G__959
											continue
										}
									}
								} else {
									{
										var vec__951 = First.X_invoke_Arity1(seq__946___1)
										var k = Nth.X_invoke_Arity3(vec__951, float64(0), nil)
										var v = Nth.X_invoke_Arity3(vec__951, float64(1), nil)
										_, _, _ = vec__951, k, v
										f.(CljsCoreIFn).X_invoke_Arity2(v, k).(interface{})
										{
											var G__960 = Next.X_invoke_Arity1(seq__946___1)
											var G__961 = nil
											var G__962 = float64(0)
											var G__963 = float64(0)
											seq__946 = G__960
											chunk__947 = G__961
											count__948 = G__962
											i__949 = G__963
											continue
										}
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
}

func (_ *CljsCorePersistentTreeSet) CljsCoreILookup__() {}
func (self__ *CljsCorePersistentTreeSet) X_lookup_Arity2(v interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return _lookup.X_invoke_Arity3(coll___1, v, nil).(interface{})
	}
}

func (self__ *CljsCorePersistentTreeSet) X_lookup_Arity3(v interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var n = Native_invoke_instance_method.X_invoke_Arity3(self__.Tree_map, "Entry_at", []interface{}{v})
			_ = n
			if !(n == nil) {
				return Native_get_instance_field.X_invoke_Arity2(n, "Key")
			} else {
				return not_found
			}
		}
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIMeta__() {}
func (self__ *CljsCorePersistentTreeSet) X_meta_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return self__.Meta
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreICloneable__() {}
func (self__ *CljsCorePersistentTreeSet) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCorePersistentTreeSet{self__.Meta, self__.Tree_map, self__.X__hash})
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreICounted__() {}
func (self__ *CljsCorePersistentTreeSet) X_count_Arity1() float64 {
	{
		var coll___1 = self__
		_ = coll___1
		return Count.X_invoke_Arity1(self__.Tree_map).(float64)
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIReversible__() {}
func (self__ *CljsCorePersistentTreeSet) X_rseq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		if Count.X_invoke_Arity1(self__.Tree_map).(float64) > float64(0) {
			return Map_.X_invoke_Arity2(Key, Rseq.Arity1IQ(self__.Tree_map)).(*CljsCoreLazySeq)
		} else {
			return nil
		}
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIHash__() {}
func (self__ *CljsCorePersistentTreeSet) X_hash_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_unordered_coll.X_invoke_Arity1(coll___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIEquiv__() {}
func (self__ *CljsCorePersistentTreeSet) X_equiv_Arity2(other interface{}) bool {
	{
		var coll___1 = self__
		_ = coll___1
		return (Set_QMARK_.Arity1IB(other)) && (Count.X_invoke_Arity1(coll___1).(float64) == Count.X_invoke_Arity1(other).(float64)) && (Every_QMARK_.X_invoke_Arity2(func(coll___1 *CljsCorePersistentTreeSet) *AFn {
			return func(G__964 *AFn) *AFn {
				return Fn(G__964, func(p1__939_SHARP_ interface{}) interface{} {
					return Contains_QMARK_.X_invoke_Arity2(coll___1, p1__939_SHARP_)
				})
			}(&AFn{})
		}(coll___1), other))
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCorePersistentTreeSet) X_empty_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return With_meta.X_invoke_Arity2(CljsCorePersistentTreeSet_EMPTY, self__.Meta)
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreISet__() {}
func (self__ *CljsCorePersistentTreeSet) X_disjoin_Arity2(v interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentTreeSet{self__.Meta, Dissoc.X_invoke_Arity2(self__.Tree_map, v), nil})
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreISeqable__() {}
func (self__ *CljsCorePersistentTreeSet) X_seq_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Keys.X_invoke_Arity1(self__.Tree_map)
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIWithMeta__() {}
func (self__ *CljsCorePersistentTreeSet) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentTreeSet{meta___1, self__.Tree_map, self__.X__hash})
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreICollection__() {}
func (self__ *CljsCorePersistentTreeSet) X_conj_Arity2(o interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return (&CljsCorePersistentTreeSet{self__.Meta, Assoc.X_invoke_Arity3(self__.Tree_map, o, nil).(interface{}), nil})
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIFn__() {}
func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity1(k interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity2(k).(interface{})
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity2(k interface{}, not_found interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return coll___1.X_lookup_Arity3(k, not_found).(interface{})
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 3"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 4"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 5"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 6"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 7"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 8"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 9"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 10"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 11"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 12"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 13"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 14"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 15"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 16"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 17"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 18"}))
	}
}

func (self__ *CljsCorePersistentTreeSet) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 19"}))
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreISorted__() {}
func (self__ *CljsCorePersistentTreeSet) X_sorted_seq_Arity2(ascending_QMARK_ interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Map_.X_invoke_Arity2(Key, _sorted_seq.X_invoke_Arity2(self__.Tree_map, ascending_QMARK_)).(*CljsCoreLazySeq)
	}
}

func (self__ *CljsCorePersistentTreeSet) X_sorted_seq_from_Arity3(k interface{}, ascending_QMARK_ interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Map_.X_invoke_Arity2(Key, _sorted_seq_from.X_invoke_Arity3(self__.Tree_map, k, ascending_QMARK_)).(*CljsCoreLazySeq)
	}
}

func (self__ *CljsCorePersistentTreeSet) X_entry_key_Arity2(entry interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return entry
	}
}

func (self__ *CljsCorePersistentTreeSet) X_comparator_Arity1() interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return _comparator.X_invoke_Arity1(self__.Tree_map).(interface{})
	}
}

var X__GT_PersistentTreeSet = func(__GT_PersistentTreeSet *AFn) *AFn {
	return Fn(__GT_PersistentTreeSet, func(meta interface{}, tree_map interface{}, __hash interface{}) interface{} {
		return (&CljsCorePersistentTreeSet{meta, tree_map, __hash})
	})
}(&AFn{})

var CljsCorePersistentTreeSet_EMPTY = (&CljsCorePersistentTreeSet{nil, CljsCorePersistentTreeMap_EMPTY, float64(0)})

var Set_from_indexed_seq = func(set_from_indexed_seq *AFn) *AFn {
	return Fn(set_from_indexed_seq, func(iseq interface{}) interface{} {
		{
			var arr = Native_get_instance_field.X_invoke_Arity2(iseq, "Arr")
			var ret = func() interface{} {
				var a__77522__auto__ = arr
				_ = a__77522__auto__
				{
					var i = float64(0)
					var res = CljsCorePersistentHashSet_EMPTY.X_as_transient_Arity1()
					_, _ = i, res
					for {
						if i < float64(len(a__77522__auto__.([]interface{}))) {
							{
								var G__965 = (i + float64(1))
								var G__966 = res.X_conj_BANG__Arity2((arr.([]interface{})[int(i)]))
								i = G__965
								res = G__966
								continue
							}
						} else {
							return res
						}
					}
				}
			}()
			_, _ = arr, ret
			return ret.X_persistent_BANG__Arity1()
		}
	})
}(&AFn{})

/**
* Returns a set of the distinct elements of coll.
 */
var Set = func(set *AFn) *AFn {
	return Fn(set, func(coll interface{}) interface{} {
		{
			var in = Seq.Arity1IQ(coll)
			_ = in
			if in == nil {
				return CljsCorePersistentHashSet_EMPTY
			} else {
				if (func() bool { _, instanceof := in.(*CljsCoreIndexedSeq); return instanceof }()) && (in.I.(float64) == float64(0)) {
					return Set_from_indexed_seq.X_invoke_Arity1(in).(interface{})
				} else {
					{
						var in___1 = in
						var out = CljsCorePersistentHashSet_EMPTY.X_as_transient_Arity1()
						_, _ = in___1, out
						for {
							if !(in___1 == nil) {
								{
									var G__967 = in___1.X_next_Arity1()
									var G__968 = out.X_conj_BANG__Arity2(in___1.X_first_Arity1().(interface{}))
									in___1 = G__967
									out = G__968
									continue
								}
							} else {
								return out.X_persistent_BANG__Arity1()
							}
						}
					}

				}
			}
		}
	})
}(&AFn{})

/**
* @param {...*} var_args
 */
var Hash_set = func(hash_set *AFn) *AFn {
	return Fn(hash_set, func() interface{} {
		return CljsCorePersistentHashSet_EMPTY
	}, func(keys__ ...interface{}) interface{} {
		var keys = Array_seq.X_invoke_Arity1(keys__[0:])
		_ = keys
		return Set.X_invoke_Arity1(keys).(interface{})
	})
}(&AFn{})

/**
* Returns a new sorted set with supplied keys.
* @param {...*} var_args
 */
var Sorted_set = func(sorted_set *AFn) *AFn {
	return Fn(sorted_set, func(keys__ ...interface{}) interface{} {
		var keys = Array_seq.X_invoke_Arity1(keys__[0:])
		_ = keys
		return Reduce.X_invoke_Arity3(X_conj, CljsCorePersistentTreeSet_EMPTY, keys)
	})
}(&AFn{})

/**
* Returns a new sorted set with supplied keys, using the supplied comparator.
* @param {...*} var_args
 */
var Sorted_set_by = func(sorted_set_by *AFn) *AFn {
	return Fn(sorted_set_by, func(comparator_keys__ ...interface{}) interface{} {
		var comparator = comparator_keys__[0]
		var keys = Array_seq.X_invoke_Arity1(comparator_keys__[1:])
		_, _ = comparator, keys
		return Reduce.X_invoke_Arity3(X_conj, (&CljsCorePersistentTreeSet{nil, Sorted_map_by.X_invoke_Arity1(comparator).(*CljsCorePersistentTreeMap), float64(0)}), keys)
	})
}(&AFn{})

/**
* Given a map of replacement pairs and a vector/collection, returns a
* vector/seq with any elements = a key in smap replaced with the
* corresponding val in smap.  Returns a transducer when no collection
* is provided.
 */
var Replace = func(replace *AFn) *AFn {
	return Fn(replace, func(smap interface{}) interface{} {
		return Map_.X_invoke_Arity1(func(G__971 *AFn) *AFn {
			return Fn(G__971, func(p1__969_SHARP_ interface{}) interface{} {
				{
					var temp__4217__auto__ = Find.X_invoke_Arity2(smap, p1__969_SHARP_)
					_ = temp__4217__auto__
					if Truth_(temp__4217__auto__) {
						{
							var e = temp__4217__auto__
							_ = e
							return Val.X_invoke_Arity1(e).(interface{})
						}
					} else {
						return p1__969_SHARP_
					}
				}
			})
		}(&AFn{})).(interface{})
	}, func(smap interface{}, coll interface{}) interface{} {
		if Vector_QMARK_.Arity1IB(coll) {
			{
				var n = Count.X_invoke_Arity1(coll).(float64)
				_ = n
				return Reduce.X_invoke_Arity3(func(n float64) *AFn {
					return func(G__972 *AFn) *AFn {
						return Fn(G__972, func(v interface{}, i interface{}) interface{} {
							{
								var temp__4217__auto__ = Find.X_invoke_Arity2(smap, Nth.X_invoke_Arity2(v, i))
								_ = temp__4217__auto__
								if Truth_(temp__4217__auto__) {
									{
										var e = temp__4217__auto__
										_ = e
										return Assoc.X_invoke_Arity3(v, i, Second.X_invoke_Arity1(e)).(interface{})
									}
								} else {
									return v
								}
							}
						})
					}(&AFn{})
				}(n), coll, Take.X_invoke_Arity2(n, Iterate.X_invoke_Arity2(Inc, float64(0)).(*CljsCoreCons)).(*CljsCoreLazySeq))
			}
		} else {
			return Map_.X_invoke_Arity2(func(G__973 *AFn) *AFn {
				return Fn(G__973, func(p1__970_SHARP_ interface{}) interface{} {
					{
						var temp__4217__auto__ = Find.X_invoke_Arity2(smap, p1__970_SHARP_)
						_ = temp__4217__auto__
						if Truth_(temp__4217__auto__) {
							{
								var e = temp__4217__auto__
								_ = e
								return Second.X_invoke_Arity1(e)
							}
						} else {
							return p1__970_SHARP_
						}
					}
				})
			}(&AFn{}), coll).(*CljsCoreLazySeq)
		}
	})
}(&AFn{})

/**
* Returns a lazy sequence of the elements of coll with duplicates removed
 */
var Distinct = func(distinct *AFn) *AFn {
	return Fn(distinct, func(coll interface{}) interface{} {
		{
			var step = func(step *AFn) *AFn {
				return Fn(step, func(xs interface{}, seen interface{}) interface{} {
					return (&CljsCoreLazySeq{nil, func(G__982 *AFn) *AFn {
						return Fn(G__982, func() interface{} {
							return func(G__983 *AFn) *AFn {
								return Fn(G__983, func(p__980 interface{}, seen___1 interface{}) interface{} {
									for {
										{
											var vec__981 = p__980
											var f = Nth.X_invoke_Arity3(vec__981, float64(0), nil)
											var xs___1 = vec__981
											_, _, _ = vec__981, f, xs___1
											{
												var temp__4219__auto__ = Seq.Arity1IQ(xs___1)
												_ = temp__4219__auto__
												if Truth_(temp__4219__auto__) {
													{
														var s = temp__4219__auto__
														_ = s
														if Contains_QMARK_.X_invoke_Arity2(seen___1, f) {
															{
																var G__984 = Rest.X_invoke_Arity1(s)
																var G__985 = seen___1
																p__980 = G__984
																seen___1 = G__985
																continue
															}
														} else {
															return Cons.X_invoke_Arity2(f, step.X_invoke_Arity2(Rest.X_invoke_Arity1(s), Conj.X_invoke_Arity2(seen___1, f).(interface{})).(*CljsCoreLazySeq)).(*CljsCoreCons)
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
			return step.(CljsCoreIFn).X_invoke_Arity2(coll, CljsCorePersistentHashSet_EMPTY).(*CljsCoreLazySeq)
		}
	})
}(&AFn{})

var Butlast = func(butlast *AFn) *AFn {
	return Fn(butlast, func(s interface{}) interface{} {
		{
			var ret = CljsCorePersistentVector_EMPTY
			var s___1 = s
			_, _ = ret, s___1
			for {
				if Truth_(Next.Arity1IQ(s___1)) {
					{
						var G__986 = Conj.X_invoke_Arity2(ret, First.X_invoke_Arity1(s___1)).(interface{})
						var G__987 = Next.Arity1IQ(s___1)
						ret = G__986
						s___1 = G__987
						continue
					}
				} else {
					return Seq.X_invoke_Arity1(ret)
				}
			}
		}
	})
}(&AFn{})

/**
* Returns the name String of a string, symbol or keyword.
 */
var Name = func(name *AFn) *AFn {
	return Fn(name, func(x interface{}) interface{} {
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "INamed", Str: "cljs.core/INamed", X_hash: float64(-857199025), X_meta: nil}), x)) {
			return x.X_name_Arity1()
		} else {
			if reflect.TypeOf(x).Kind() == reflect.String {
				return x
			} else {
				panic((&js.Error{("Doesn't support name: " + Str.X_invoke_Arity1(x).(string))}))
			}
		}
	})
}(&AFn{})

/**
* Returns a map with the keys mapped to the corresponding vals.
 */
var Zipmap = func(zipmap *AFn) *AFn {
	return Fn(zipmap, func(keys interface{}, vals interface{}) interface{} {
		{
			var map_ = Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY).(interface{})
			var ks = Seq.Arity1IQ(keys)
			var vs = Seq.Arity1IQ(vals)
			_, _, _ = map_, ks, vs
			for {
				if (ks) && (vs) {
					{
						var G__988 = Assoc_BANG_.X_invoke_Arity3(map_, First.X_invoke_Arity1(ks), First.X_invoke_Arity1(vs)).(interface{})
						var G__989 = Next.X_invoke_Arity1(ks)
						var G__990 = Next.X_invoke_Arity1(vs)
						map_ = G__988
						ks = G__989
						vs = G__990
						continue
					}
				} else {
					return Persistent_BANG_.X_invoke_Arity1(map_).(interface{})
				}
			}
		}
	})
}(&AFn{})

/**
* Returns the x for which (k x), a number, is greatest.
* @param {...*} var_args
 */
var Max_key = func(max_key *AFn) *AFn {
	return Fn(max_key, func(k interface{}, x interface{}) interface{} {
		return x
	}, func(k interface{}, x interface{}, y interface{}) interface{} {
		if k.(CljsCoreIFn).X_invoke_Arity1(x).(interface{}).(float64) > k.(CljsCoreIFn).X_invoke_Arity1(y).(interface{}).(float64) {
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
		return Reduce.X_invoke_Arity3(func(G__993 *AFn) *AFn {
			return Fn(G__993, func(p1__991_SHARP_ interface{}, p2__992_SHARP_ interface{}) interface{} {
				return max_key.X_invoke_Arity3(k, p1__991_SHARP_, p2__992_SHARP_).(interface{})
			})
		}(&AFn{}), max_key.X_invoke_Arity3(k, x, y).(interface{}), more)
	})
}(&AFn{})

/**
* Returns the x for which (k x), a number, is least.
* @param {...*} var_args
 */
var Min_key = func(min_key *AFn) *AFn {
	return Fn(min_key, func(k interface{}, x interface{}) interface{} {
		return x
	}, func(k interface{}, x interface{}, y interface{}) interface{} {
		if k.(CljsCoreIFn).X_invoke_Arity1(x).(interface{}).(float64) < k.(CljsCoreIFn).X_invoke_Arity1(y).(interface{}).(float64) {
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
		return Reduce.X_invoke_Arity3(func(G__996 *AFn) *AFn {
			return Fn(G__996, func(p1__994_SHARP_ interface{}, p2__995_SHARP_ interface{}) interface{} {
				return min_key.X_invoke_Arity3(k, p1__994_SHARP_, p2__995_SHARP_).(interface{})
			})
		}(&AFn{}), min_key.X_invoke_Arity3(k, x, y).(interface{}), more)
	})
}(&AFn{})

type CljsCoreArrayList struct{ Arr interface{} }

func (_ *CljsCoreArrayList) CljsCoreObject__() {}
func (self__ *CljsCoreArrayList) Add(x interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return Native_invoke_instance_method.X_invoke_Arity3(self__.Arr, "Push", []interface{}{x})
	}
}

func (self__ *CljsCoreArrayList) Size() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return float64(len(self__.Arr.([]interface{})))
	}
}

func (self__ *CljsCoreArrayList) Clear() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return func() []interface{} {
			var return__997 = []interface{}{}
			self__.Arr = return__997
			return return__997
		}()
	}
}

func (self__ *CljsCoreArrayList) IsEmpty() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (float64(len(self__.Arr.([]interface{}))) == float64(0))
	}
}

func (self__ *CljsCoreArrayList) ToArray() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return self__.Arr
	}
}

var X__GT_ArrayList = func(__GT_ArrayList *AFn) *AFn {
	return Fn(__GT_ArrayList, func(arr interface{}) interface{} {
		return (&CljsCoreArrayList{arr})
	})
}(&AFn{})

var Array_list = func(array_list *AFn) *AFn {
	return Fn(array_list, func() interface{} {
		return (&CljsCoreArrayList{[]interface{}{}})
	})
}(&AFn{})

/**
* Returns a lazy sequence of lists like partition, but may include
* partitions with fewer than n items at the end.  Returns a stateful
* transducer when no collection is provided.
 */
var Partition_all = func(partition_all *AFn) *AFn {
	return Fn(partition_all, func(n interface{}) interface{} {
		return func(G__998 *AFn) *AFn {
			return Fn(G__998, func(f1 interface{}) interface{} {
				{
					var a = Array_list.X_invoke_Arity0().(*CljsCoreArrayList)
					_ = a
					return func(a *CljsCoreArrayList) *AFn {
						return func(G__999 *AFn) *AFn {
							return Fn(G__999, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
							}, func(result interface{}) interface{} {
								{
									var result___1 = func() interface{} {
										if Truth_(a.IsEmpty().(bool)) {
											return result
										} else {
											return func() interface{} {
												var v = Vec.X_invoke_Arity1(a.ToArray()).(interface{})
												_ = v
												a.Clear()
												return f1.(CljsCoreIFn).X_invoke_Arity2(result, v).(interface{})
											}()
										}
									}()
									_ = result___1
									return f1.(CljsCoreIFn).X_invoke_Arity1(result___1).(interface{})
								}
							}, func(result interface{}, input interface{}) interface{} {
								a.Add(input)
								if n.(float64) == a.Size().(float64) {
									{
										var v = Vec.X_invoke_Arity1(a.ToArray()).(interface{})
										_ = v
										a.Clear()
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, v).(interface{})
									}
								} else {
									return result
								}
							})
						}(&AFn{})
					}(a)
				}
			})
		}(&AFn{})
	}, func(n interface{}, coll interface{}) interface{} {
		return partition_all.X_invoke_Arity3(n, n, coll).(*CljsCoreLazySeq)
	}, func(n interface{}, step interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__1000 *AFn) *AFn {
			return Fn(G__1000, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
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

/**
* Returns a lazy sequence of successive items from coll while
* (pred item) returns true. pred must be free of side-effects.
* Returns a transducer when no collection is provided.
 */
var Take_while = func(take_while *AFn) *AFn {
	return Fn(take_while, func(pred interface{}) interface{} {
		return func(G__1001 *AFn) *AFn {
			return Fn(G__1001, func(f1 interface{}) interface{} {
				return func(G__1002 *AFn) *AFn {
					return Fn(G__1002, func() interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
					}, func(result interface{}) interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
					}, func(result interface{}, input interface{}) interface{} {
						if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(input).(interface{})) {
							return f1.(CljsCoreIFn).X_invoke_Arity2(result, input).(interface{})
						} else {
							return Reduced.X_invoke_Arity1(result).(*CljsCoreReduced)
						}
					})
				}(&AFn{})
			})
		}(&AFn{})
	}, func(pred interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__1003 *AFn) *AFn {
			return Fn(G__1003, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
							_ = s
							if Truth_(pred.(CljsCoreIFn).X_invoke_Arity1(First.X_invoke_Arity1(s)).(interface{})) {
								return Cons.X_invoke_Arity2(First.X_invoke_Arity1(s), take_while.X_invoke_Arity2(pred, Rest.X_invoke_Arity1(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
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

var Mk_bound_fn = func(mk_bound_fn *AFn) *AFn {
	return Fn(mk_bound_fn, func(sc interface{}, test interface{}, key interface{}) interface{} {
		return func(G__1004 *AFn) *AFn {
			return Fn(G__1004, func(e interface{}) interface{} {
				{
					var comp = sc.(CljsCoreISorted).X_comparator_Arity1().(interface{})
					_ = comp
					return test.(CljsCoreIFn).X_invoke_Arity2(comp.(CljsCoreIFn).X_invoke_Arity2(sc.(CljsCoreISorted).X_entry_key_Arity2(e).(interface{}), key).(interface{}), float64(0)).(interface{})
				}
			})
		}(&AFn{})
	})
}(&AFn{})

/**
* sc must be a sorted collection, test(s) one of <, <=, > or
* >=. Returns a seq of those entries with keys ek for
* which (test (.. sc comparator (compare ek key)) 0) is true
 */
var Subseq = func(subseq *AFn) *AFn {
	return Fn(subseq, func(sc interface{}, test interface{}, key interface{}) interface{} {
		{
			var include = Mk_bound_fn.X_invoke_Arity3(sc, test, key).(interface{})
			_ = include
			if Truth_(CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{X_GT_, X_GT__EQ_}, true).(*CljsCorePersistentHashSet).X_invoke_Arity1(test).(interface{})) {
				{
					var temp__4219__auto__ = sc.(CljsCoreISorted).X_sorted_seq_from_Arity3(key, true)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var vec__1007 = temp__4219__auto__
							var e = Nth.X_invoke_Arity3(vec__1007, float64(0), nil)
							var s = vec__1007
							_, _, _ = vec__1007, e, s
							if Truth_(include.(CljsCoreIFn).X_invoke_Arity1(e).(interface{})) {
								return s
							} else {
								return Next.X_invoke_Arity1(s)
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
			var temp__4219__auto__ = sc.(CljsCoreISorted).X_sorted_seq_from_Arity3(start_key, true)
			_ = temp__4219__auto__
			if Truth_(temp__4219__auto__) {
				{
					var vec__1008 = temp__4219__auto__
					var e = Nth.X_invoke_Arity3(vec__1008, float64(0), nil)
					var s = vec__1008
					_, _, _ = vec__1008, e, s
					return Take_while.X_invoke_Arity2(Mk_bound_fn.X_invoke_Arity3(sc, end_test, end_key).(interface{}), func() interface{} {
						if Truth_(Mk_bound_fn.X_invoke_Arity3(sc, start_test, start_key).(interface{}).X_invoke_Arity1(e).(interface{}).(bool)) {
							return s
						} else {
							return Next.X_invoke_Arity1(s)
						}
					}()).(*CljsCoreLazySeq)
				}
			} else {
				return nil
			}
		}
	})
}(&AFn{})

/**
* sc must be a sorted collection, test(s) one of <, <=, > or
* >=. Returns a reverse seq of those entries with keys ek for
* which (test (.. sc comparator (compare ek key)) 0) is true
 */
var Rsubseq = func(rsubseq *AFn) *AFn {
	return Fn(rsubseq, func(sc interface{}, test interface{}, key interface{}) interface{} {
		{
			var include = Mk_bound_fn.X_invoke_Arity3(sc, test, key).(interface{})
			_ = include
			if Truth_(CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{X_LT_, X_LT__EQ_}, true).(*CljsCorePersistentHashSet).X_invoke_Arity1(test).(interface{})) {
				{
					var temp__4219__auto__ = sc.(CljsCoreISorted).X_sorted_seq_from_Arity3(key, false)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var vec__1011 = temp__4219__auto__
							var e = Nth.X_invoke_Arity3(vec__1011, float64(0), nil)
							var s = vec__1011
							_, _, _ = vec__1011, e, s
							if Truth_(include.(CljsCoreIFn).X_invoke_Arity1(e).(interface{})) {
								return s
							} else {
								return Next.X_invoke_Arity1(s)
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
			var temp__4219__auto__ = sc.(CljsCoreISorted).X_sorted_seq_from_Arity3(end_key, false)
			_ = temp__4219__auto__
			if Truth_(temp__4219__auto__) {
				{
					var vec__1012 = temp__4219__auto__
					var e = Nth.X_invoke_Arity3(vec__1012, float64(0), nil)
					var s = vec__1012
					_, _, _ = vec__1012, e, s
					return Take_while.X_invoke_Arity2(Mk_bound_fn.X_invoke_Arity3(sc, start_test, start_key).(interface{}), func() interface{} {
						if Truth_(Mk_bound_fn.X_invoke_Arity3(sc, end_test, end_key).(interface{}).X_invoke_Arity1(e).(interface{}).(bool)) {
							return s
						} else {
							return Next.X_invoke_Arity1(s)
						}
					}()).(*CljsCoreLazySeq)
				}
			} else {
				return nil
			}
		}
	})
}(&AFn{})

type CljsCoreRange struct {
	Meta    interface{}
	Start   interface{}
	End     interface{}
	Step    interface{}
	X__hash interface{}
}

func (_ *CljsCoreRange) CljsCoreObject__() {}
func (self__ *CljsCoreRange) ToString() string {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_str_STAR_.X_invoke_Arity1(coll___1).(string)
	}
}

func (self__ *CljsCoreRange) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreRange) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

func (_ *CljsCoreRange) CljsCoreIIndexed__() {}
func (self__ *CljsCoreRange) X_nth_Arity2(n interface{}) interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		if n.(float64) < rng___1.X_count_Arity1() {
			return (self__.Start.(float64) + (n.(float64) * self__.Step.(float64)))
		} else {
			if (self__.Start.(float64) > self__.End.(float64)) && (self__.Step.(float64) == float64(0)) {
				return self__.Start
			} else {
				panic((&js.Error{"Index out of bounds"}))
			}
		}
	}
}

func (self__ *CljsCoreRange) X_nth_Arity3(n interface{}, not_found interface{}) interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		if n.(float64) < rng___1.X_count_Arity1() {
			return (self__.Start.(float64) + (n.(float64) * self__.Step.(float64)))
		} else {
			if (self__.Start.(float64) > self__.End.(float64)) && (self__.Step.(float64) == float64(0)) {
				return self__.Start
			} else {
				return not_found
			}
		}
	}
}

func (_ *CljsCoreRange) CljsCoreIMeta__() {}
func (self__ *CljsCoreRange) X_meta_Arity1() interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		return self__.Meta
	}
}

func (_ *CljsCoreRange) CljsCoreICloneable__() {}
func (self__ *CljsCoreRange) X_clone_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return (&CljsCoreRange{self__.Meta, self__.Start, self__.End, self__.Step, self__.X__hash})
	}
}

func (_ *CljsCoreRange) CljsCoreINext__() {}
func (self__ *CljsCoreRange) X_next_Arity1() interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		if self__.Step.(float64) > float64(0) {
			if (self__.Start.(float64) + self__.Step.(float64)) < self__.End.(float64) {
				return (&CljsCoreRange{self__.Meta, (self__.Start.(float64) + self__.Step.(float64)), self__.End, self__.Step, nil})
			} else {
				return nil
			}
		} else {
			if (self__.Start.(float64) + self__.Step.(float64)) > self__.End.(float64) {
				return (&CljsCoreRange{self__.Meta, (self__.Start.(float64) + self__.Step.(float64)), self__.End, self__.Step, nil})
			} else {
				return nil
			}
		}
	}
}

func (_ *CljsCoreRange) CljsCoreICounted__() {}
func (self__ *CljsCoreRange) X_count_Arity1() float64 {
	{
		var rng___1 = self__
		_ = rng___1
		if Not.X_invoke_Arity1(rng___1.X_seq_Arity1()) {
			return float64(0)
		} else {
			return Math.Ceil(((self__.End.(float64) - self__.Start.(float64)) / self__.Step.(float64)))
		}
	}
}

func (_ *CljsCoreRange) CljsCoreIHash__() {}
func (self__ *CljsCoreRange) X_hash_Arity1() interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		{
			var h__77043__auto__ = self__.X__hash
			_ = h__77043__auto__
			if !(h__77043__auto__ == nil) {
				return h__77043__auto__
			} else {
				{
					var h__77043__auto_____1 = Hash_ordered_coll.X_invoke_Arity1(rng___1)
					_ = h__77043__auto_____1
					self__.X__hash = h__77043__auto_____1

					return h__77043__auto_____1
				}
			}
		}
	}
}

func (_ *CljsCoreRange) CljsCoreIEquiv__() {}
func (self__ *CljsCoreRange) X_equiv_Arity2(other interface{}) bool {
	{
		var rng___1 = self__
		_ = rng___1
		return Equiv_sequential.X_invoke_Arity2(rng___1, other).(bool)
	}
}

func (_ *CljsCoreRange) CljsCoreIEmptyableCollection__() {}
func (self__ *CljsCoreRange) X_empty_Arity1() interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		return With_meta.X_invoke_Arity2(CljsCoreList_EMPTY, self__.Meta)
	}
}

func (_ *CljsCoreRange) CljsCoreIReduce__() {}
func (self__ *CljsCoreRange) X_reduce_Arity2(f interface{}) interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		return Ci_reduce.X_invoke_Arity2(rng___1, f).(interface{})
	}
}

func (self__ *CljsCoreRange) X_reduce_Arity3(f interface{}, s interface{}) interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		return Ci_reduce.X_invoke_Arity3(rng___1, f, s)
	}
}

func (_ *CljsCoreRange) CljsCoreISeq__() {}
func (self__ *CljsCoreRange) X_first_Arity1() interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		if rng___1.X_seq_Arity1() == nil {
			return nil
		} else {
			return self__.Start
		}
	}
}

func (self__ *CljsCoreRange) X_rest_Arity1() interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		if !(rng___1.X_seq_Arity1() == nil) {
			return (&CljsCoreRange{self__.Meta, (self__.Start.(float64) + self__.Step.(float64)), self__.End, self__.Step, nil})
		} else {
			return CljsCoreList_EMPTY
		}
	}
}

func (_ *CljsCoreRange) CljsCoreISeqable__() {}
func (self__ *CljsCoreRange) X_seq_Arity1() interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		if self__.Step.(float64) > float64(0) {
			if self__.Start.(float64) < self__.End.(float64) {
				return rng___1
			} else {
				return nil
			}
		} else {
			if self__.Start.(float64) > self__.End.(float64) {
				return rng___1
			} else {
				return nil
			}
		}
	}
}

func (_ *CljsCoreRange) CljsCoreISequential__() {}
func (_ *CljsCoreRange) CljsCoreIWithMeta__()   {}
func (self__ *CljsCoreRange) X_with_meta_Arity2(meta___1 interface{}) interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		return (&CljsCoreRange{meta___1, self__.Start, self__.End, self__.Step, self__.X__hash})
	}
}

func (_ *CljsCoreRange) CljsCoreICollection__() {}
func (self__ *CljsCoreRange) X_conj_Arity2(o interface{}) interface{} {
	{
		var rng___1 = self__
		_ = rng___1
		return Cons.X_invoke_Arity2(o, rng___1).(*CljsCoreCons)
	}
}

var X__GT_Range = func(__GT_Range *AFn) *AFn {
	return Fn(__GT_Range, func(meta interface{}, start interface{}, end interface{}, step interface{}, __hash interface{}) interface{} {
		return (&CljsCoreRange{meta, start, end, step, __hash})
	})
}(&AFn{})

/**
* Returns a lazy seq of nums from start (inclusive) to end
* (exclusive), by step, where start defaults to 0, step to 1,
* and end to infinity.
 */
var Range_ = func(range_ *AFn) *AFn {
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

/**
* Returns a lazy seq of every nth item in coll.  Returns a stateful
* transducer when no collection is provided.
 */
var Take_nth = func(take_nth *AFn) *AFn {
	return Fn(take_nth, func(n interface{}) interface{} {
		return func(G__1013 *AFn) *AFn {
			return Fn(G__1013, func(f1 interface{}) interface{} {
				{
					var ia = Atom.X_invoke_Arity1(float64(-1)).(interface{})
					_ = ia
					return func(ia interface{}) *AFn {
						return func(G__1014 *AFn) *AFn {
							return Fn(G__1014, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
							}, func(result interface{}, input interface{}) interface{} {
								{
									var i = Swap_BANG_.X_invoke_Arity2(ia, Inc)
									_ = i
									if Rem.X_invoke_Arity2(i, n).(float64) == float64(0) {
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, input).(interface{})
									} else {
										return result
									}
								}
							})
						}(&AFn{})
					}(ia)
				}
			})
		}(&AFn{})
	}, func(n interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__1015 *AFn) *AFn {
			return Fn(G__1015, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
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

/**
* Returns a vector of [(take-while pred coll) (drop-while pred coll)]
 */
var Split_with = func(split_with *AFn) *AFn {
	return Fn(split_with, func(pred interface{}, coll interface{}) interface{} {
		return &CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{Take_while.X_invoke_Arity2(pred, coll).(*CljsCoreLazySeq), Drop_while.X_invoke_Arity2(pred, coll).(*CljsCoreLazySeq)}, nil}
	})
}(&AFn{})

/**
* Applies f to each value in coll, splitting it each time f returns a
* new value.  Returns a lazy seq of partitions.  Returns a stateful
* transducer when no collection is provided.
 */
var Partition_by = func(partition_by *AFn) *AFn {
	return Fn(partition_by, func(f interface{}) interface{} {
		return func(G__1017 *AFn) *AFn {
			return Fn(G__1017, func(f1 interface{}) interface{} {
				{
					var a = Array_list.X_invoke_Arity0().(*CljsCoreArrayList)
					var pa = Atom.X_invoke_Arity1((&CljsCoreKeyword{Ns: "cljs.core", Name: "none", Fqn: "cljs.core/none", X_hash: float64(926646439)})).(interface{})
					_, _ = a, pa
					return func(a *CljsCoreArrayList, pa interface{}) *AFn {
						return func(G__1018 *AFn) *AFn {
							return Fn(G__1018, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
							}, func(result interface{}) interface{} {
								{
									var result___1 = func() interface{} {
										if Truth_(a.IsEmpty().(bool)) {
											return result
										} else {
											return func() interface{} {
												var v = Vec.X_invoke_Arity1(a.ToArray()).(interface{})
												_ = v
												a.Clear()
												return f1.(CljsCoreIFn).X_invoke_Arity2(result, v).(interface{})
											}()
										}
									}()
									_ = result___1
									return f1.(CljsCoreIFn).X_invoke_Arity1(result___1).(interface{})
								}
							}, func(result interface{}, input interface{}) interface{} {
								{
									var pval = Deref.X_invoke_Arity1(pa).(interface{})
									var val = f.(CljsCoreIFn).X_invoke_Arity1(input).(interface{})
									_, _ = pval, val
									Reset_BANG_.X_invoke_Arity2(pa, val).(interface{})
									if (Keyword_identical_QMARK_.X_invoke_Arity2(pval, (&CljsCoreKeyword{Ns: "cljs.core", Name: "none", Fqn: "cljs.core/none", X_hash: float64(926646439)}))) || (X_EQ_.X_invoke_Arity2(val, pval)) {
										a.Add(input)
										return result
									} else {
										{
											var v = Vec.X_invoke_Arity1(a.ToArray()).(interface{})
											_ = v
											a.Clear()
											a.Add(input)
											return f1.(CljsCoreIFn).X_invoke_Arity2(result, v).(interface{})
										}
									}
								}
							})
						}(&AFn{})
					}(a, pa)
				}
			})
		}(&AFn{})
	}, func(f interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__1019 *AFn) *AFn {
			return Fn(G__1019, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
							_ = s
							{
								var fst = First.X_invoke_Arity1(s)
								var fv = f.(CljsCoreIFn).X_invoke_Arity1(fst).(interface{})
								var run = Cons.X_invoke_Arity2(fst, Take_while.X_invoke_Arity2(func(fst interface{}, fv interface{}, s CljsCoreISeq, temp__4219__auto__ CljsCoreISeq) *AFn {
									return func(G__1020 *AFn) *AFn {
										return Fn(G__1020, func(p1__1016_SHARP_ interface{}) interface{} {
											return X_EQ_.X_invoke_Arity2(fv, f.(CljsCoreIFn).X_invoke_Arity1(p1__1016_SHARP_).(interface{}))
										})
									}(&AFn{})
								}(fst, fv, s, temp__4219__auto__), Next.X_invoke_Arity1(s)).(*CljsCoreLazySeq)).(*CljsCoreCons)
								_, _, _ = fst, fv, run
								return Cons.X_invoke_Arity2(run, partition_by.X_invoke_Arity2(f, Seq.X_invoke_Arity1(Drop.X_invoke_Arity2(Count.X_invoke_Arity1(run).(float64), s).(*CljsCoreLazySeq))).(*CljsCoreLazySeq)).(*CljsCoreCons)
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

/**
* Returns a map from distinct items in coll to the number of times
* they appear.
 */
var Frequencies = func(frequencies *AFn) *AFn {
	return Fn(frequencies, func(coll interface{}) interface{} {
		return Persistent_BANG_.X_invoke_Arity1(Reduce.X_invoke_Arity3(func(G__1021 *AFn) *AFn {
			return Fn(G__1021, func(counts interface{}, x interface{}) interface{} {
				return Assoc_BANG_.X_invoke_Arity3(counts, x, (Get.X_invoke_Arity3(counts, x, float64(0)).(float64) + float64(1))).(interface{})
			})
		}(&AFn{}), Transient.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY).(interface{}), coll)).(interface{})
	})
}(&AFn{})

/**
* Returns a lazy seq of the intermediate values of the reduction (as
* per reduce) of coll by f, starting with init.
 */
var Reductions = func(reductions *AFn) *AFn {
	return Fn(reductions, func(f interface{}, coll interface{}) interface{} {
		return (&CljsCoreLazySeq{nil, func(G__1022 *AFn) *AFn {
			return Fn(G__1022, func() interface{} {
				{
					var temp__4217__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4217__auto__
					if Truth_(temp__4217__auto__) {
						{
							var s = temp__4217__auto__
							_ = s
							return reductions.X_invoke_Arity3(f, First.X_invoke_Arity1(s), Rest.X_invoke_Arity1(s)).(*CljsCoreCons)
						}
					} else {
						return CljsCoreList_EMPTY.(CljsCoreICollection).X_conj_Arity2(f.(CljsCoreIFn).X_invoke_Arity0().(interface{}))
					}
				}
			})
		}(&AFn{}), nil, nil})
	}, func(f interface{}, init interface{}, coll interface{}) interface{} {
		return Cons.X_invoke_Arity2(init, (&CljsCoreLazySeq{nil, func(G__1023 *AFn) *AFn {
			return Fn(G__1023, func() interface{} {
				{
					var temp__4219__auto__ = Seq.Arity1IQ(coll)
					_ = temp__4219__auto__
					if Truth_(temp__4219__auto__) {
						{
							var s = temp__4219__auto__
							_ = s
							return reductions.X_invoke_Arity3(f, f.(CljsCoreIFn).X_invoke_Arity2(init, First.X_invoke_Arity1(s)).(interface{}), Rest.X_invoke_Arity1(s)).(*CljsCoreCons)
						}
					} else {
						return nil
					}
				}
			})
		}(&AFn{}), nil, nil})).(*CljsCoreCons)
	})
}(&AFn{})

/**
* Takes a set of functions and returns a fn that is the juxtaposition
* of those fns.  The returned fn takes a variable number of args, and
* returns a vector containing the result of applying each fn to the
* args (left-to-right).
* ((juxt a b c) x) => [(a x) (b x) (c x)]
* @param {...*} var_args
 */
var Juxt = func(juxt *AFn) *AFn {
	return Fn(juxt, func(f interface{}) interface{} {
		return func(G__1034 *AFn) *AFn {
			return Fn(G__1034, func() interface{} {
				return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity0().(interface{})}, nil})
			}, func(x interface{}) interface{} {
				return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})}, nil})
			}, func(x interface{}, y interface{}) interface{} {
				return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{})}, nil})
			}, func(x interface{}, y interface{}, z interface{}) interface{} {
				return (&CljsCorePersistentVector{nil, float64(1), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity3(x, y, z).(interface{})}, nil})
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
		return func(G__1035 *AFn) *AFn {
			return Fn(G__1035, func() interface{} {
				return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity0().(interface{}), g.(CljsCoreIFn).X_invoke_Arity0().(interface{})}, nil})
			}, func(x interface{}) interface{} {
				return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity1(x).(interface{}), g.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})}, nil})
			}, func(x interface{}, y interface{}) interface{} {
				return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{}), g.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{})}, nil})
			}, func(x interface{}, y interface{}, z interface{}) interface{} {
				return (&CljsCorePersistentVector{nil, float64(2), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity3(x, y, z).(interface{}), g.(CljsCoreIFn).X_invoke_Arity3(x, y, z).(interface{})}, nil})
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
		return func(G__1036 *AFn) *AFn {
			return Fn(G__1036, func() interface{} {
				return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity0().(interface{}), g.(CljsCoreIFn).X_invoke_Arity0().(interface{}), h.(CljsCoreIFn).X_invoke_Arity0().(interface{})}, nil})
			}, func(x interface{}) interface{} {
				return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity1(x).(interface{}), g.(CljsCoreIFn).X_invoke_Arity1(x).(interface{}), h.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})}, nil})
			}, func(x interface{}, y interface{}) interface{} {
				return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{}), g.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{}), h.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{})}, nil})
			}, func(x interface{}, y interface{}, z interface{}) interface{} {
				return (&CljsCorePersistentVector{nil, float64(3), float64(5), CljsCorePersistentVector_EMPTY_NODE, []interface{}{f.(CljsCoreIFn).X_invoke_Arity3(x, y, z).(interface{}), g.(CljsCoreIFn).X_invoke_Arity3(x, y, z).(interface{}), h.(CljsCoreIFn).X_invoke_Arity3(x, y, z).(interface{})}, nil})
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
			return func(fs___1 *CljsCoreCons) *AFn {
				return func(G__1037 *AFn) *AFn {
					return Fn(G__1037, func() interface{} {
						return Reduce.X_invoke_Arity3(func(fs___1 *CljsCoreCons) *AFn {
							return func(G__1038 *AFn) *AFn {
								return Fn(G__1038, func(p1__1024_SHARP_ interface{}, p2__1025_SHARP_ interface{}) interface{} {
									return Conj.X_invoke_Arity2(p1__1024_SHARP_, p2__1025_SHARP_.(CljsCoreIFn).X_invoke_Arity0().(interface{})).(interface{})
								})
							}(&AFn{})
						}(fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(fs___1 *CljsCoreCons) *AFn {
							return func(G__1039 *AFn) *AFn {
								return Fn(G__1039, func(p1__1026_SHARP_ interface{}, p2__1027_SHARP_ interface{}) interface{} {
									return Conj.X_invoke_Arity2(p1__1026_SHARP_, p2__1027_SHARP_.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})).(interface{})
								})
							}(&AFn{})
						}(fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x interface{}, y interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(fs___1 *CljsCoreCons) *AFn {
							return func(G__1040 *AFn) *AFn {
								return Fn(G__1040, func(p1__1028_SHARP_ interface{}, p2__1029_SHARP_ interface{}) interface{} {
									return Conj.X_invoke_Arity2(p1__1028_SHARP_, p2__1029_SHARP_.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{})).(interface{})
								})
							}(&AFn{})
						}(fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(fs___1 *CljsCoreCons) *AFn {
							return func(G__1041 *AFn) *AFn {
								return Fn(G__1041, func(p1__1030_SHARP_ interface{}, p2__1031_SHARP_ interface{}) interface{} {
									return Conj.X_invoke_Arity2(p1__1030_SHARP_, p2__1031_SHARP_.(CljsCoreIFn).X_invoke_Arity3(x, y, z).(interface{})).(interface{})
								})
							}(&AFn{})
						}(fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					}, func(x_y_z_args__ ...interface{}) interface{} {
						var x = x_y_z_args__[0]
						var y = x_y_z_args__[1]
						var z = x_y_z_args__[2]
						var args = Array_seq.X_invoke_Arity1(x_y_z_args__[3:])
						_, _, _, _ = x, y, z, args
						return Reduce.X_invoke_Arity3(func(fs___1 *CljsCoreCons) *AFn {
							return func(G__1042 *AFn) *AFn {
								return Fn(G__1042, func(p1__1032_SHARP_ interface{}, p2__1033_SHARP_ interface{}) interface{} {
									return Conj.X_invoke_Arity2(p1__1032_SHARP_, Apply.X_invoke_Arity5(p2__1033_SHARP_, x, y, z, args)).(interface{})
								})
							}(&AFn{})
						}(fs___1), CljsCorePersistentVector_EMPTY, fs___1)
					})
				}(&AFn{})
			}(fs___1)
		}
	})
}(&AFn{})

/**
* When lazy sequences are produced via functions that have side
* effects, any effects other than those needed to produce the first
* element in the seq do not occur until the seq is consumed. dorun can
* be used to force any effects. Walks through the successive nexts of
* the seq, does not retain the head and returns nil.
 */
var Dorun = func(dorun *AFn) *AFn {
	return Fn(dorun, func(coll interface{}) interface{} {
		for {
			if Truth_(Seq.Arity1IQ(coll)) {
				{
					var G__1043 = Next.Arity1IQ(coll)
					coll = G__1043
					continue
				}
			} else {
				return nil
			}
		}
	}, func(n interface{}, coll interface{}) interface{} {
		for {
			if (Seq.Arity1IQ(coll)) && (n.(float64) > float64(0)) {
				{
					var G__1044 = (n.(float64) - float64(1))
					var G__1045 = Next.Arity1IQ(coll)
					n = G__1044
					coll = G__1045
					continue
				}
			} else {
				return nil
			}
		}
	})
}(&AFn{})

/**
* When lazy sequences are produced via functions that have side
* effects, any effects other than those needed to produce the first
* element in the seq do not occur until the seq is consumed. doall can
* be used to force any effects. Walks through the successive nexts of
* the seq, retains the head and returns it, thus causing the entire
* seq to reside in memory at one time.
 */
var Doall = func(doall *AFn) *AFn {
	return Fn(doall, func(coll interface{}) interface{} {
		Dorun.X_invoke_Arity1(coll)
		return coll
	}, func(n interface{}, coll interface{}) interface{} {
		Dorun.X_invoke_Arity2(n, coll)
		return coll
	})
}(&AFn{})

var Regexp_QMARK_ = func(regexp_QMARK_ *AFn) *AFn {
	return Fn(regexp_QMARK_, func(o interface{}) interface{} {
		return func() bool { _, instanceof := o.(*js.RegExp); return instanceof }()
	})
}(&AFn{})

/**
* Returns the result of (re-find re s) if re fully matches s.
 */
var Re_matches = func(re_matches *AFn) *AFn {
	return Fn(re_matches, func(re interface{}, s interface{}) interface{} {
		if reflect.TypeOf(s).Kind() == reflect.String {
			{
				var matches = Native_invoke_instance_method.X_invoke_Arity3(re, "Exec", []interface{}{s})
				_ = matches
				if X_EQ_.X_invoke_Arity2(First.X_invoke_Arity1(matches), s) {
					if Count.X_invoke_Arity1(matches).(float64) == float64(1) {
						return First.X_invoke_Arity1(matches)
					} else {
						return Vec.X_invoke_Arity1(matches).(interface{})
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

/**
* Returns the first regex match, if any, of s to re, using
* re.exec(s). Returns a vector, containing first the matching
* substring, then any capturing groups if the regular expression contains
* capturing groups.
 */
var Re_find = func(re_find *AFn) *AFn {
	return Fn(re_find, func(re interface{}, s interface{}) interface{} {
		if reflect.TypeOf(s).Kind() == reflect.String {
			{
				var matches = Native_invoke_instance_method.X_invoke_Arity3(re, "Exec", []interface{}{s})
				_ = matches
				if matches == nil {
					return nil
				} else {
					if Count.X_invoke_Arity1(matches).(float64) == float64(1) {
						return First.X_invoke_Arity1(matches)
					} else {
						return Vec.X_invoke_Arity1(matches).(interface{})
					}
				}
			}
		} else {
			panic((&js.TypeError{"re-find must match against a string."}))
		}
	})
}(&AFn{})

/**
* Returns a lazy sequence of successive matches of re in s.
 */
var Re_seq = func(re_seq *AFn) *AFn {
	return Fn(re_seq, func(re interface{}, s interface{}) interface{} {
		{
			var match_data = Re_find.X_invoke_Arity2(re, s)
			var match_idx = Native_invoke_instance_method.X_invoke_Arity3(s, "Search", []interface{}{re})
			var match_str = func() interface{} {
				if Coll_QMARK_.X_invoke_Arity1(match_data) {
					return First.X_invoke_Arity1(match_data)
				} else {
					return match_data
				}
			}()
			var post_match = Subs.X_invoke_Arity2(s, (match_idx.(float64) + Count.X_invoke_Arity1(match_str).(float64))).(interface{})
			_, _, _, _ = match_data, match_idx, match_str, post_match
			if Truth_(match_data) {
				return (&CljsCoreLazySeq{nil, func(match_data interface{}, match_idx interface{}, match_str interface{}, post_match interface{}) *AFn {
					return func(G__1046 *AFn) *AFn {
						return Fn(G__1046, func() interface{} {
							return Cons.X_invoke_Arity2(match_data, func() interface{} {
								if Truth_(Seq.X_invoke_Arity1(post_match).(bool)) {
									return re_seq.X_invoke_Arity2(re, post_match)
								} else {
									return nil
								}
							}()).(*CljsCoreCons)
						})
					}(&AFn{})
				}(match_data, match_idx, match_str, post_match), nil, nil})
			} else {
				return nil
			}
		}
	})
}(&AFn{})

/**
* Returns an instance of RegExp which has compiled the provided string.
 */
var Re_pattern = func(re_pattern *AFn) *AFn {
	return Fn(re_pattern, func(s interface{}) interface{} {
		{
			var vec__1048 = Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `^(?:\(\?([idmsux]*)\))?(.*)`, Flags: ``}), s)
			var ___ = Nth.X_invoke_Arity3(vec__1048, float64(0), nil)
			var flags = Nth.X_invoke_Arity3(vec__1048, float64(1), nil)
			var pattern = Nth.X_invoke_Arity3(vec__1048, float64(2), nil)
			_, _, _, _ = vec__1048, ___, flags, pattern
			return (&js.RegExp{pattern, flags})
		}
	})
}(&AFn{})

var Pr_sequential_writer = func(pr_sequential_writer *AFn) *AFn {
	return Fn(pr_sequential_writer, func(writer interface{}, print_one interface{}, begin interface{}, sep interface{}, end interface{}, opts interface{}, coll interface{}) interface{} {
		{
			var _STAR_print_level_STAR_1050 = X_STAR_print_level_STAR_
			_ = _STAR_print_level_STAR_1050
			return func() interface{} {
				defer func() {
					X_STAR_print_level_STAR_ = _STAR_print_level_STAR_1050

				}()
				{
					X_STAR_print_level_STAR_ = func() interface{} {
						if X_STAR_print_level_STAR_ == nil {
							return nil
						} else {
							return (X_STAR_print_level_STAR_.(float64) - float64(1))
						}
					}()

					if (!(X_STAR_print_level_STAR_ == nil)) && (X_STAR_print_level_STAR_.(float64) < float64(0)) {
						return writer.(CljsCoreIWriter).X_write_Arity2("#").(interface{})
					} else {
						writer.(CljsCoreIWriter).X_write_Arity2(begin).(interface{})
						if Truth_(Seq.Arity1IQ(coll)) {
							print_one.(CljsCoreIFn).X_invoke_Arity3(First.X_invoke_Arity1(coll), writer, opts).(interface{})
						} else {
						}
						{
							var coll_1051___1 = Next.Arity1IQ(coll)
							var n_1052 = ((&CljsCoreKeyword{Ns: nil, Name: "print-length", Fqn: "print-length", X_hash: float64(1931866356)}).X_invoke_Arity1(opts).(interface{}).(float64) - float64(1))
							_, _ = coll_1051___1, n_1052
							for {
								if (coll_1051___1) && ((n_1052 == nil) || (!(n_1052 == float64(0)))) {
									writer.(CljsCoreIWriter).X_write_Arity2(sep).(interface{})
									print_one.(CljsCoreIFn).X_invoke_Arity3(First.X_invoke_Arity1(coll_1051___1), writer, opts).(interface{})
									{
										var G__1053 = Next.X_invoke_Arity1(coll_1051___1)
										var G__1054 = (n_1052 - float64(1))
										coll_1051___1 = G__1053
										n_1052 = G__1054
										continue
									}
								} else {
									if (Seq.X_invoke_Arity1(coll_1051___1)) && (n_1052 == float64(0)) {
										writer.(CljsCoreIWriter).X_write_Arity2(sep).(interface{})
										writer.(CljsCoreIWriter).X_write_Arity2("...").(interface{})
									} else {
									}
								}
								break
							}
						}
						return writer.(CljsCoreIWriter).X_write_Arity2(end).(interface{})
					}
				}
			}()
		}
	})
}(&AFn{})

/**
* @param {...*} var_args
 */
var Write_all = func(write_all *AFn) *AFn {
	return Fn(write_all, func(writer_ss__ ...interface{}) interface{} {
		var writer = writer_ss__[0]
		var ss = Array_seq.X_invoke_Arity1(writer_ss__[1:])
		_, _ = writer, ss
		{
			var seq__1059 = Seq.Arity1IQ(ss)
			var chunk__1060 = nil
			var count__1061 = float64(0)
			var i__1062 = float64(0)
			_, _, _, _ = seq__1059, chunk__1060, count__1061, i__1062
			for {
				if i__1062 < count__1061 {
					{
						var s = chunk__1060.X_nth_Arity2(i__1062).(interface{})
						_ = s
						writer.(CljsCoreIWriter).X_write_Arity2(s).(interface{})
						{
							var G__1063 = seq__1059
							var G__1064 = chunk__1060
							var G__1065 = count__1061
							var G__1066 = (i__1062 + float64(1))
							seq__1059 = G__1063
							chunk__1060 = G__1064
							count__1061 = G__1065
							i__1062 = G__1066
							continue
						}
					}
				} else {
					{
						var temp__4219__auto__ = Seq.X_invoke_Arity1(seq__1059)
						_ = temp__4219__auto__
						if Truth_(temp__4219__auto__) {
							{
								var seq__1059___1 = temp__4219__auto__
								_ = seq__1059___1
								if Chunked_seq_QMARK_.X_invoke_Arity1(seq__1059___1) {
									{
										var c__77428__auto__ = Chunk_first.X_invoke_Arity1(seq__1059___1).(interface{})
										_ = c__77428__auto__
										{
											var G__1067 = Chunk_rest.X_invoke_Arity1(seq__1059___1).(interface{})
											var G__1068 = c__77428__auto__
											var G__1069 = Count.X_invoke_Arity1(c__77428__auto__).(float64)
											var G__1070 = float64(0)
											seq__1059 = G__1067
											chunk__1060 = G__1068
											count__1061 = G__1069
											i__1062 = G__1070
											continue
										}
									}
								} else {
									{
										var s = First.X_invoke_Arity1(seq__1059___1)
										_ = s
										writer.(CljsCoreIWriter).X_write_Arity2(s).(interface{})
										{
											var G__1071 = Next.X_invoke_Arity1(seq__1059___1)
											var G__1072 = nil
											var G__1073 = float64(0)
											var G__1074 = float64(0)
											seq__1059 = G__1071
											chunk__1060 = G__1072
											count__1061 = G__1073
											i__1062 = G__1074
											continue
										}
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

var String_print = func(string_print *AFn) *AFn {
	return Fn(string_print, func(x interface{}) interface{} {
		X_STAR_print_fn_STAR_.X_invoke_Arity1(x).(interface{})
		return nil
	})
}(&AFn{})

var Flush = func(flush *AFn) *AFn {
	return Fn(flush, func() interface{} {
		return nil
	})
}(&AFn{})

var Char_escapes = func() interface{} {
	var obj1076 = map[string]interface{}{`"\""`: "\\\"", `"\\"`: "\\\\", `"\b"`: "\\b", `"\f"`: "\\f", `"\n"`: "\\n", `"\r"`: "\\r", `"\t"`: "\\t"}
	_ = obj1076
	return obj1076
}()

var Quote_string = func(quote_string *AFn) *AFn {
	return Fn(quote_string, func(s interface{}) interface{} {
		return ("\"" + Str.X_invoke_Arity1(Native_invoke_instance_method.X_invoke_Arity3(s, "Replace", []interface{}{js.RegExp("[\\\\\"\b\f\n\r\t]", "g"), func(G__1077 *AFn) *AFn {
			return Fn(G__1077, func(match interface{}) interface{} {
				return (Char_escapes.([]interface{})[int(match.(float64))])
			})
		}(&AFn{})})).(string) + "\"")
	})
}(&AFn{})

/**
* Prefer this to pr-seq, because it makes the printing function
* configurable, allowing efficient implementations such as appending
* to a StringBuffer.
 */
var Pr_writer = func(pr_writer *AFn) *AFn {
	return Fn(pr_writer, func(obj interface{}, writer interface{}, opts interface{}) interface{} {
		if obj == nil {
			return writer.(CljsCoreIWriter).X_write_Arity2("nil").(interface{})
		} else {
			if nil == obj {
				return writer.(CljsCoreIWriter).X_write_Arity2("#<undefined>").(interface{})
			} else {
				if Truth_(func() interface{} {
					var and__76620__auto__ = Get.X_invoke_Arity2(opts, (&CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}))
					_ = and__76620__auto__
					if Truth_(and__76620__auto__) {
						{
							var and__76620__auto_____1 = Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IMeta", Str: "cljs.core/IMeta", X_hash: float64(-1459057517), X_meta: nil}), obj)
							_ = and__76620__auto_____1
							if Truth_(and__76620__auto_____1) {
								return Meta.X_invoke_Arity1(obj)
							} else {
								return and__76620__auto_____1
							}
						}
					} else {
						return and__76620__auto__
					}
				}()) {
					writer.(CljsCoreIWriter).X_write_Arity2("^").(interface{})
					pr_writer.X_invoke_Arity3(Meta.X_invoke_Arity1(obj), writer, opts)
					writer.(CljsCoreIWriter).X_write_Arity2(" ").(interface{})
				} else {
				}
				if obj == nil {
					return writer.(CljsCoreIWriter).X_write_Arity2("nil").(interface{})
				} else {
					if Native_get_instance_field.X_invoke_Arity2(obj, "Cljs_lang_type") {
						return Native_invoke_instance_method.X_invoke_Arity3(obj, "Cljs_lang_ctorPrWriter", []interface{}{obj, writer, opts})
					} else {
						if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IPrintWithWriter", Str: "cljs.core/IPrintWithWriter", X_hash: float64(1349251417), X_meta: nil}), obj)) {
							return obj.X_pr_writer_Arity3(writer, opts).(interface{})
						} else {
							if (Type_.X_invoke_Arity1(obj) == js.Boolean) || (reflect.TypeOf(obj).Kind() == reflect.Float64) {
								return writer.(CljsCoreIWriter).X_write_Arity2((`` + Str.X_invoke_Arity1(obj).(string))).(interface{})
							} else {
								if Object_QMARK_.Arity1IB(obj) {
									writer.(CljsCoreIWriter).X_write_Arity2("#js ").(interface{})
									return Print_map.X_invoke_Arity4(Map_.X_invoke_Arity2(func(G__1078 *AFn) *AFn {
										return Fn(G__1078, func(k interface{}) interface{} {
											return &CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{Keyword.X_invoke_Arity1(k), (obj.([]interface{})[int(k.(float64))])}, nil}
										})
									}(&AFn{}), Js_keys.X_invoke_Arity1(obj).([]interface{})).(*CljsCoreLazySeq), pr_writer, writer, opts).(interface{})
								} else {
									if reflect.TypeOf(obj).Kind() == reflect.Slice {
										return Pr_sequential_writer.X_invoke_Arity7(writer, pr_writer, "#js [", " ", "]", opts, obj).(interface{})
									} else {
										if goog.IsString(obj) {
											if Truth_((&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}).X_invoke_Arity1(opts).(interface{})) {
												return writer.(CljsCoreIWriter).X_write_Arity2(Quote_string.X_invoke_Arity1(obj).(string)).(interface{})
											} else {
												return writer.(CljsCoreIWriter).X_write_Arity2(obj).(interface{})
											}
										} else {
											if Fn_QMARK_.Arity1IB(obj) {
												return Write_all.X_invoke_ArityVariadic(writer, "#<", (`` + Str.X_invoke_Arity1(obj).(string)), ">")
											} else {
												if func() bool { _, instanceof := obj.(*js.Date); return instanceof }() {
													{
														var normalize = func(G__1079 *AFn) *AFn {
															return Fn(G__1079, func(n interface{}, len interface{}) interface{} {
																{
																	var ns = (`` + Str.X_invoke_Arity1(n).(string))
																	_ = ns
																	for {
																		if Count.X_invoke_Arity1(ns).(float64) < len.(float64) {
																			{
																				var G__1080 = ("0" + Str.X_invoke_Arity1(ns).(string))
																				ns = G__1080
																				continue
																			}
																		} else {
																			return ns
																		}
																	}
																}
															})
														}(&AFn{})
														_ = normalize
														return Write_all.X_invoke_ArityVariadic(writer, "#inst \"", (`` + Str.X_invoke_Arity1(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCFullYear", []interface{}{})).(string)), "-", normalize.(CljsCoreIFn).X_invoke_Arity2((Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCMonth", []interface{}{}).(float64)+float64(1)), float64(2)).(string), "-", normalize.(CljsCoreIFn).X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCDate", []interface{}{}), float64(2)).(string), "T", normalize.(CljsCoreIFn).X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCHours", []interface{}{}), float64(2)).(string), ":", normalize.(CljsCoreIFn).X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCMinutes", []interface{}{}), float64(2)).(string), ":", normalize.(CljsCoreIFn).X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCSeconds", []interface{}{}), float64(2)).(string), ".", normalize.(CljsCoreIFn).X_invoke_Arity2(Native_invoke_instance_method.X_invoke_Arity3(obj, "GetUTCMilliseconds", []interface{}{}), float64(3)).(string), "-", "00:00\"")
													}
												} else {
													if Regexp_QMARK_.X_invoke_Arity1(obj).(bool) {
														return Write_all.X_invoke_ArityVariadic(writer, "#\"", Native_get_instance_field.X_invoke_Arity2(obj, "Source"), "\"")
													} else {
														if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IPrintWithWriter", Str: "cljs.core/IPrintWithWriter", X_hash: float64(1349251417), X_meta: nil}), obj)) {
															return obj.(CljsCoreIPrintWithWriter).X_pr_writer_Arity3(writer, opts).(interface{})
														} else {
															return Write_all.X_invoke_ArityVariadic(writer, "#<", (`` + Str.X_invoke_Arity1(obj).(string)), ">")

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
	})
}(&AFn{})

var Pr_seq_writer = func(pr_seq_writer *AFn) *AFn {
	return Fn(pr_seq_writer, func(objs interface{}, writer interface{}, opts interface{}) interface{} {
		Pr_writer.X_invoke_Arity3(First.X_invoke_Arity1(objs), writer, opts)
		{
			var seq__1085 = Seq.X_invoke_Arity1(Next.Arity1IQ(objs))
			var chunk__1086 = nil
			var count__1087 = float64(0)
			var i__1088 = float64(0)
			_, _, _, _ = seq__1085, chunk__1086, count__1087, i__1088
			for {
				if i__1088 < count__1087 {
					{
						var obj = chunk__1086.X_nth_Arity2(i__1088).(interface{})
						_ = obj
						writer.(CljsCoreIWriter).X_write_Arity2(" ").(interface{})
						Pr_writer.X_invoke_Arity3(obj, writer, opts)
						{
							var G__1089 = seq__1085
							var G__1090 = chunk__1086
							var G__1091 = count__1087
							var G__1092 = (i__1088 + float64(1))
							seq__1085 = G__1089
							chunk__1086 = G__1090
							count__1087 = G__1091
							i__1088 = G__1092
							continue
						}
					}
				} else {
					{
						var temp__4219__auto__ = Seq.X_invoke_Arity1(seq__1085)
						_ = temp__4219__auto__
						if Truth_(temp__4219__auto__) {
							{
								var seq__1085___1 = temp__4219__auto__
								_ = seq__1085___1
								if Chunked_seq_QMARK_.X_invoke_Arity1(seq__1085___1) {
									{
										var c__77428__auto__ = Chunk_first.X_invoke_Arity1(seq__1085___1).(interface{})
										_ = c__77428__auto__
										{
											var G__1093 = Chunk_rest.X_invoke_Arity1(seq__1085___1).(interface{})
											var G__1094 = c__77428__auto__
											var G__1095 = Count.X_invoke_Arity1(c__77428__auto__).(float64)
											var G__1096 = float64(0)
											seq__1085 = G__1093
											chunk__1086 = G__1094
											count__1087 = G__1095
											i__1088 = G__1096
											continue
										}
									}
								} else {
									{
										var obj = First.X_invoke_Arity1(seq__1085___1)
										_ = obj
										writer.(CljsCoreIWriter).X_write_Arity2(" ").(interface{})
										Pr_writer.X_invoke_Arity3(obj, writer, opts)
										{
											var G__1097 = Next.X_invoke_Arity1(seq__1085___1)
											var G__1098 = nil
											var G__1099 = float64(0)
											var G__1100 = float64(0)
											seq__1085 = G__1097
											chunk__1086 = G__1098
											count__1087 = G__1099
											i__1088 = G__1100
											continue
										}
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

var Pr_sb_with_opts = func(pr_sb_with_opts *AFn) *AFn {
	return Fn(pr_sb_with_opts, func(objs interface{}, opts interface{}) interface{} {
		{
			var sb = (&goog_string.StringBuffer{})
			var writer = (&CljsCoreStringBufferWriter{sb})
			_, _ = sb, writer
			Pr_seq_writer.X_invoke_Arity3(objs, writer, opts)
			writer.X_flush_Arity1().(interface{})
			return sb
		}
	})
}(&AFn{})

/**
* Prints a sequence of objects to a string, observing all the
* options given in opts
 */
var Pr_str_with_opts = func(pr_str_with_opts *AFn) *AFn {
	return Fn(pr_str_with_opts, func(objs interface{}, opts interface{}) interface{} {
		if Empty_QMARK_.Arity1IB(objs) {
			return ""
		} else {
			return (`` + Str.X_invoke_Arity1(Pr_sb_with_opts.X_invoke_Arity2(objs, opts).(goog.GoogStringStringBuffer)).(string))
		}
	})
}(&AFn{})

/**
* Same as pr-str-with-opts followed by (newline)
 */
var Prn_str_with_opts = func(prn_str_with_opts *AFn) *AFn {
	return Fn(prn_str_with_opts, func(objs interface{}, opts interface{}) interface{} {
		if Empty_QMARK_.Arity1IB(objs) {
			return "\n"
		} else {
			{
				var sb = Pr_sb_with_opts.X_invoke_Arity2(objs, opts).(goog.GoogStringStringBuffer)
				_ = sb
				sb.Append('\n')
				return (`` + Str.X_invoke_Arity1(sb).(string))
			}
		}
	})
}(&AFn{})

/**
* Prints a sequence of objects using string-print, observing all
* the options given in opts
 */
var Pr_with_opts = func(pr_with_opts *AFn) *AFn {
	return Fn(pr_with_opts, func(objs interface{}, opts interface{}) interface{} {
		return String_print.X_invoke_Arity1(Pr_str_with_opts.X_invoke_Arity2(objs, opts).(string))
	})
}(&AFn{})

var Newline = func(newline *AFn) *AFn {
	return Fn(newline, func(opts interface{}) interface{} {
		String_print.X_invoke_Arity1("\n")
		if Truth_(Get.X_invoke_Arity2(opts, (&CljsCoreKeyword{Ns: nil, Name: "flush-on-newline", Fqn: "flush-on-newline", X_hash: float64(-151457939)}))) {
			return Flush.X_invoke_Arity0()
		} else {
			return nil
		}
	})
}(&AFn{})

/**
* pr to a string, returning it. Fundamental entrypoint to IPrintWithWriter.
* @param {...*} var_args
 */
var Pr_str = func(pr_str *AFn) *AFn {
	return Fn(pr_str, func(objs__ ...interface{}) interface{} {
		var objs = Array_seq.X_invoke_Arity1(objs__[0:])
		_ = objs
		return Pr_str_with_opts.X_invoke_Arity2(objs, Pr_opts.X_invoke_Arity0().(CljsCoreIMap)).(string)
	})
}(&AFn{})

/**
* Same as pr-str followed by (newline)
* @param {...*} var_args
 */
var Prn_str = func(prn_str *AFn) *AFn {
	return Fn(prn_str, func(objs__ ...interface{}) interface{} {
		var objs = Array_seq.X_invoke_Arity1(objs__[0:])
		_ = objs
		return Prn_str_with_opts.X_invoke_Arity2(objs, Pr_opts.X_invoke_Arity0().(CljsCoreIMap)).(string)
	})
}(&AFn{})

/**
* Prints the object(s) using string-print.  Prints the
* object(s), separated by spaces if there is more than one.
* By default, pr and prn print in a way that objects can be
* read by the reader
* @param {...*} var_args
 */
var Pr = func(pr *AFn) *AFn {
	return Fn(pr, func(objs__ ...interface{}) interface{} {
		var objs = Array_seq.X_invoke_Arity1(objs__[0:])
		_ = objs
		return Pr_with_opts.X_invoke_Arity2(objs, Pr_opts.X_invoke_Arity0().(CljsCoreIMap))
	})
}(&AFn{})

/**
* Prints the object(s) using string-print.
* print and println produce output for human consumption.
* @param {...*} var_args
 */
var Print = func(cljs_core_print *AFn) *AFn {
	return Fn(cljs_core_print, func(objs__ ...interface{}) interface{} {
		var objs = Array_seq.X_invoke_Arity1(objs__[0:])
		_ = objs
		return Pr_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false).(interface{}))
	})
}(&AFn{})

/**
* print to a string, returning it
* @param {...*} var_args
 */
var Print_str = func(print_str *AFn) *AFn {
	return Fn(print_str, func(objs__ ...interface{}) interface{} {
		var objs = Array_seq.X_invoke_Arity1(objs__[0:])
		_ = objs
		return Pr_str_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false).(interface{})).(string)
	})
}(&AFn{})

/**
* Same as print followed by (newline)
* @param {...*} var_args
 */
var Println = func(println *AFn) *AFn {
	return Fn(println, func(objs__ ...interface{}) interface{} {
		var objs = Array_seq.X_invoke_Arity1(objs__[0:])
		_ = objs
		Pr_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false).(interface{}))
		if Truth_(X_STAR_print_newline_STAR_) {
			return Newline.X_invoke_Arity1(Pr_opts.X_invoke_Arity0().(CljsCoreIMap))
		} else {
			return nil
		}
	})
}(&AFn{})

/**
* println to a string, returning it
* @param {...*} var_args
 */
var Println_str = func(println_str *AFn) *AFn {
	return Fn(println_str, func(objs__ ...interface{}) interface{} {
		var objs = Array_seq.X_invoke_Arity1(objs__[0:])
		_ = objs
		return Prn_str_with_opts.X_invoke_Arity2(objs, Assoc.X_invoke_Arity3(Pr_opts.X_invoke_Arity0().(CljsCoreIMap), (&CljsCoreKeyword{Ns: nil, Name: "readably", Fqn: "readably", X_hash: float64(1129599760)}), false).(interface{})).(string)
	})
}(&AFn{})

/**
* Same as pr followed by (newline).
* @param {...*} var_args
 */
var Prn = func(prn *AFn) *AFn {
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

var Print_map = func(print_map *AFn) *AFn {
	return Fn(print_map, func(m interface{}, print_one interface{}, writer interface{}, opts interface{}) interface{} {
		return Pr_sequential_writer.X_invoke_Arity7(writer, func(G__1101 *AFn) *AFn {
			return Fn(G__1101, func(e interface{}, w interface{}, opts___1 interface{}) interface{} {
				print_one.(CljsCoreIFn).X_invoke_Arity3(Key.X_invoke_Arity1(e).(interface{}), w, opts___1).(interface{})
				w.(CljsCoreIWriter).X_write_Arity2(' ').(interface{})
				return print_one.(CljsCoreIFn).X_invoke_Arity3(Val.X_invoke_Arity1(e).(interface{}), w, opts___1).(interface{})
			})
		}(&AFn{}), "{", ", ", "}", opts, Seq.Arity1IQ(m)).(interface{})
	})
}(&AFn{})

func (_ *CljsCoreIndexedSeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreIndexedSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreLazySeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreLazySeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCorePersistentTreeMapSeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCorePersistentTreeMapSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreNodeSeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreNodeSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreBlackNode) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreBlackNode) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "[", " ", "]", opts, coll___1).(interface{})
	}
}

func (_ *CljsCorePersistentArrayMapSeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCorePersistentArrayMapSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCorePersistentTreeSet) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCorePersistentTreeSet) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "#{", " ", "}", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreChunkedSeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreChunkedSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreObjMap) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreObjMap) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Print_map.X_invoke_Arity4(coll___1, Pr_writer, writer, opts).(interface{})
	}
}

func (_ *CljsCoreCons) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreCons) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreRSeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreRSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCorePersistentHashMap) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCorePersistentHashMap) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Print_map.X_invoke_Arity4(coll___1, Pr_writer, writer, opts).(interface{})
	}
}

func (_ *CljsCoreArrayNodeSeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreArrayNodeSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreSubvec) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreSubvec) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "[", " ", "]", opts, coll___1).(interface{})
	}
}

func (_ *CljsCorePersistentTreeMap) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCorePersistentTreeMap) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Print_map.X_invoke_Arity4(coll___1, Pr_writer, writer, opts).(interface{})
	}
}

func (_ *CljsCorePersistentHashSet) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCorePersistentHashSet) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "#{", " ", "}", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreChunkedCons) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreChunkedCons) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreAtom) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreAtom) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var a___1 = self__
		_ = a___1
		writer.(CljsCoreIWriter).X_write_Arity2("#<Atom: ").(interface{})
		Pr_writer.X_invoke_Arity3(a___1.State, writer, opts)
		return writer.(CljsCoreIWriter).X_write_Arity2(">").(interface{})
	}
}

func (_ *CljsCoreValSeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreValSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreRedNode) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreRedNode) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "[", " ", "]", opts, coll___1).(interface{})
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCorePersistentVector) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "[", " ", "]", opts, coll___1).(interface{})
	}
}

func (_ *CljsCorePersistentQueueSeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCorePersistentQueueSeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreEmptyList) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreEmptyList) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return writer.(CljsCoreIWriter).X_write_Arity2("()").(interface{})
	}
}

func (_ *CljsCoreLazyTransformer) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreLazyTransformer) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCorePersistentQueue) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCorePersistentQueue) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "#queue [", " ", "]", opts, Seq.X_invoke_Arity1(coll___1)).(interface{})
	}
}

func (_ *CljsCorePersistentArrayMap) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCorePersistentArrayMap) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Print_map.X_invoke_Arity4(coll___1, Pr_writer, writer, opts).(interface{})
	}
}

func (_ *CljsCoreRange) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreRange) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreKeySeq) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreKeySeq) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCoreList) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreList) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___1 = self__
		_ = coll___1
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___1).(interface{})
	}
}

func (_ *CljsCorePersistentVector) CljsCoreIComparable__() {}
func (self__ *CljsCorePersistentVector) X_compare_Arity2(y interface{}) float64 {
	{
		var x___1 = self__
		_ = x___1
		return Compare_indexed.X_invoke_Arity2(x___1, y).(float64)
	}
}

func (_ *CljsCoreSubvec) CljsCoreIComparable__() {}
func (self__ *CljsCoreSubvec) X_compare_Arity2(y interface{}) float64 {
	{
		var x___1 = self__
		_ = x___1
		return Compare_indexed.X_invoke_Arity2(x___1, y).(float64)
	}
}

func (_ *CljsCoreKeyword) CljsCoreIComparable__() {}
func (self__ *CljsCoreKeyword) X_compare_Arity2(y interface{}) float64 {
	{
		var x___1 = self__
		_ = x___1
		return Compare_symbols.X_invoke_Arity2(x___1, y)
	}
}

func (_ *CljsCoreSymbol) CljsCoreIComparable__() {}
func (self__ *CljsCoreSymbol) X_compare_Arity2(y interface{}) float64 {
	{
		var x___1 = self__
		_ = x___1
		return Compare_symbols.X_invoke_Arity2(x___1, y)
	}
}

/**
* Atomically sets the metadata for a namespace/var/ref/agent/atom to be:
*
* (apply f its-current-meta args)
*
* f must be free of side-effects
* @param {...*} var_args
 */
var Alter_meta_BANG_ = func(alter_meta_BANG_ *AFn) *AFn {
	return Fn(alter_meta_BANG_, func(iref_f_args__ ...interface{}) interface{} {
		var iref = iref_f_args__[0]
		var f = iref_f_args__[1]
		var args = Array_seq.X_invoke_Arity1(iref_f_args__[2:])
		_, _, _ = iref, f, args
		return func() interface{} {
			var return__1102 = Apply.X_invoke_Arity3(f, Native_get_instance_field.X_invoke_Arity2(iref, "Meta"), args)
			Native_set_instance_field.X_invoke_Arity3(iref, "Meta", return__1102)
			return return__1102
		}()
	})
}(&AFn{})

/**
* Atomically resets the metadata for an atom
 */
var Reset_meta_BANG_ = func(reset_meta_BANG_ *AFn) *AFn {
	return Fn(reset_meta_BANG_, func(iref interface{}, m interface{}) interface{} {
		return func() interface{} {
			var return__1103 = m
			Native_set_instance_field.X_invoke_Arity3(iref, "Meta", return__1103)
			return return__1103
		}()
	})
}(&AFn{})

/**
* Alpha - subject to change.
*
* Adds a watch function to an atom reference. The watch fn must be a
* fn of 4 args: a key, the reference, its old-state, its
* new-state. Whenever the reference's state might have been changed,
* any registered watches will have their functions called. The watch
* fn will be called synchronously. Note that an atom's state
* may have changed again prior to the fn call, so use old/new-state
* rather than derefing the reference. Keys must be unique per
* reference, and can be used to remove the watch with remove-watch,
* but are otherwise considered opaque by the watch mechanism.  Bear in
* mind that regardless of the result or action of the watch fns the
* atom's value will change.  Example:
*
* (def a (atom 0))
* (add-watch a :inc (fn [k r o n] (assert (== 0 n))))
* (swap! a inc)
* ;; Assertion Error
* (deref a)
* ;=> 1
 */
var Add_watch = func(add_watch *AFn) *AFn {
	return Fn(add_watch, func(iref interface{}, key interface{}, f interface{}) interface{} {
		return iref.(CljsCoreIWatchable).X_add_watch_Arity3(key, f).(interface{})
	})
}(&AFn{})

/**
* Alpha - subject to change.
*
* Removes a watch (set by add-watch) from a reference
 */
var Remove_watch = func(remove_watch *AFn) *AFn {
	return Fn(remove_watch, func(iref interface{}, key interface{}) interface{} {
		return iref.(CljsCoreIWatchable).X_remove_watch_Arity2(key).(interface{})
	})
}(&AFn{})

var Gensym_counter interface{} = nil

/**
* Returns a new symbol with a unique name. If a prefix string is
* supplied, the name is prefix# where # is some unique number. If
* prefix is not supplied, the prefix is 'G__'.
 */
var Gensym = func(gensym *AFn) *AFn {
	return Fn(gensym, func() interface{} {
		return gensym.X_invoke_Arity1("G__")
	}, func(prefix_string interface{}) interface{} {
		if Gensym_counter == nil {
			Gensym_counter = Atom.X_invoke_Arity1(float64(0)).(interface{})

		} else {
		}
		return Symbol.X_invoke_Arity1((`` + Str.X_invoke_Arity1(prefix_string).(string) + Str.X_invoke_Arity1(Swap_BANG_.X_invoke_Arity2(Gensym_counter, Inc)).(string)))
	})
}(&AFn{})

var Fixture1 = float64(1)

var Fixture2 = float64(2)

type CljsCoreDelay struct {
	F     interface{}
	Value interface{}
}

func (_ *CljsCoreDelay) CljsCoreIPending__() {}
func (self__ *CljsCoreDelay) X_realized_QMARK__Arity1() bool {
	{
		var d___1 = self__
		_ = d___1
		return Not.Arity1IB(self__.F)
	}
}

func (_ *CljsCoreDelay) CljsCoreIDeref__() {}
func (self__ *CljsCoreDelay) X_deref_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		if Truth_(self__.F) {
			self__.Value = self__.F.(CljsCoreIFn).X_invoke_Arity0().(interface{})

			self__.F = nil

		} else {
		}
		return self__.Value
	}
}

var X__GT_Delay = func(__GT_Delay *AFn) *AFn {
	return Fn(__GT_Delay, func(f interface{}, value interface{}) interface{} {
		return (&CljsCoreDelay{f, value})
	})
}(&AFn{})

/**
* returns true if x is a Delay created with delay
 */
var Delay_QMARK_ = func(delay_QMARK_ *AFn) *AFn {
	return Fn(delay_QMARK_, func(x interface{}) bool {
		return func() bool { _, instanceof := x.(*CljsCoreDelay); return instanceof }()
	})
}(&AFn{})

/**
* If x is a Delay, returns the (possibly cached) value of its expression, else returns x
 */
var Force = func(force *AFn) *AFn {
	return Fn(force, func(x interface{}) interface{} {
		if Delay_QMARK_.Arity1IB(x) {
			return Deref.X_invoke_Arity1(x).(interface{})
		} else {
			return x
		}
	})
}(&AFn{})

/**
* Returns true if a value has been produced for a promise, delay, future or lazy sequence.
 */
var Realized_QMARK_ = func(realized_QMARK_ *AFn) *AFn {
	return Fn(realized_QMARK_, func(d interface{}) bool {
		return d.(CljsCoreIPending).X_realized_QMARK__Arity1()
	})
}(&AFn{})

var Preserving_reduced = func(preserving_reduced *AFn) *AFn {
	return Fn(preserving_reduced, func(f1 interface{}) interface{} {
		return func(G__1106 *AFn) *AFn {
			return Fn(G__1106, func(p1__1104_SHARP_ interface{}, p2__1105_SHARP_ interface{}) interface{} {
				{
					var ret = f1.(CljsCoreIFn).X_invoke_Arity2(p1__1104_SHARP_, p2__1105_SHARP_).(interface{})
					_ = ret
					if Reduced_QMARK_.X_invoke_Arity1(ret) {
						return Reduced.X_invoke_Arity1(ret).(*CljsCoreReduced)
					} else {
						return ret
					}
				}
			})
		}(&AFn{})
	})
}(&AFn{})

/**
* maps f over coll and concatenates the results.  Thus function f
* should return a collection.  Returns a transducer when no collection
* is provided.
 */
var Flatmap = func(flatmap *AFn) *AFn {
	return Fn(flatmap, func(f interface{}) interface{} {
		return func(G__1107 *AFn) *AFn {
			return Fn(G__1107, func(f1 interface{}) interface{} {
				return func(G__1108 *AFn) *AFn {
					return Fn(G__1108, func() interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
					}, func(result interface{}) interface{} {
						return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
					}, func(result interface{}, input interface{}) interface{} {
						return Reduce.X_invoke_Arity3(Preserving_reduced.X_invoke_Arity1(f1).(interface{}), result, f.(CljsCoreIFn).X_invoke_Arity1(input).(interface{}))
					})
				}(&AFn{})
			})
		}(&AFn{})
	}, func(f interface{}, coll interface{}) interface{} {
		return Sequence.X_invoke_Arity2(flatmap.X_invoke_Arity1(f).(interface{}), coll).(interface{})
	})
}(&AFn{})

/**
* Returns a lazy sequence removing consecutive duplicates in coll.
* Returns a transducer when no collection is provided.
 */
var Dedupe = func(dedupe *AFn) *AFn {
	return Fn(dedupe, func() interface{} {
		return func(G__1109 *AFn) *AFn {
			return Fn(G__1109, func(f1 interface{}) interface{} {
				{
					var pa = Atom.X_invoke_Arity1((&CljsCoreKeyword{Ns: "cljs.core", Name: "none", Fqn: "cljs.core/none", X_hash: float64(926646439)})).(interface{})
					_ = pa
					return func(pa interface{}) *AFn {
						return func(G__1110 *AFn) *AFn {
							return Fn(G__1110, func() interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity0().(interface{})
							}, func(result interface{}) interface{} {
								return f1.(CljsCoreIFn).X_invoke_Arity1(result).(interface{})
							}, func(result interface{}, input interface{}) interface{} {
								{
									var prior = Deref.X_invoke_Arity1(pa).(interface{})
									_ = prior
									Reset_BANG_.X_invoke_Arity2(pa, input).(interface{})
									if X_EQ_.X_invoke_Arity2(prior, input) {
										return result
									} else {
										return f1.(CljsCoreIFn).X_invoke_Arity2(result, input).(interface{})
									}
								}
							})
						}(&AFn{})
					}(pa)
				}
			})
		}(&AFn{})
	}, func(coll interface{}) interface{} {
		return Sequence.X_invoke_Arity2(dedupe.X_invoke_Arity0().(interface{}), coll).(interface{})
	})
}(&AFn{})

/**
* Returns items from coll with random probability of prob (0.0 -
* 1.0).  Returns a transducer when no collection is provided.
 */
var Random_sample = func(random_sample *AFn) *AFn {
	return Fn(random_sample, func(prob interface{}) interface{} {
		return Filter.X_invoke_Arity1(func(G__1111 *AFn) *AFn {
			return Fn(G__1111, func(___ interface{}) interface{} {
				return (Rand.Arity0F() < prob.(float64))
			})
		}(&AFn{})).(interface{})
	}, func(prob interface{}, coll interface{}) interface{} {
		return Filter.X_invoke_Arity2(func(G__1112 *AFn) *AFn {
			return Fn(G__1112, func(___ interface{}) interface{} {
				return (Rand.Arity0F() < prob.(float64))
			})
		}(&AFn{}), coll).(*CljsCoreLazySeq)
	})
}(&AFn{})

type CljsCoreIteration struct {
	Xform interface{}
	Coll  interface{}
}

func (_ *CljsCoreIteration) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreIteration) X_pr_writer_Arity3(writer interface{}, opts interface{}) interface{} {
	{
		var coll___2 = self__
		_ = coll___2
		return Pr_sequential_writer.X_invoke_Arity7(writer, Pr_writer, "(", " ", ")", opts, coll___2).(interface{})
	}
}

func (_ *CljsCoreIteration) CljsCoreIReduce__() {}
func (self__ *CljsCoreIteration) X_reduce_Arity3(f interface{}, init interface{}) interface{} {
	{
		var ______1 = self__
		_ = ______1
		return Transduce.X_invoke_Arity4(self__.Xform, f, init, self__.Coll).(interface{})
	}
}

func (_ *CljsCoreIteration) CljsCoreISeqable__() {}
func (self__ *CljsCoreIteration) X_seq_Arity1() interface{} {
	{
		var ______1 = self__
		_ = ______1
		return Seq.X_invoke_Arity1(Sequence.X_invoke_Arity2(self__.Xform, self__.Coll).(interface{}))
	}
}

func (_ *CljsCoreIteration) CljsCoreISequential__() {}

var X__GT_Iteration = func(__GT_Iteration *AFn) *AFn {
	return Fn(__GT_Iteration, func(xform interface{}, coll interface{}) interface{} {
		return (&CljsCoreIteration{xform, coll})
	})
}(&AFn{})

/**
* Returns an iterable/seqable/reducible sequence of applications of
* the transducer to the items in coll. Note that these applications
* will be performed every time iterator/seq/reduce is called.
 */
var Iteration = func(iteration *AFn) *AFn {
	return Fn(iteration, func(xform interface{}, coll interface{}) interface{} {
		return (&CljsCoreIteration{xform, coll})
	})
}(&AFn{})

/**
* Runs the supplied procedure (via reduce), for purposes of side
* effects, on successive items in the collection. Returns nil
 */
var Run_BANG_ = func(run_BANG_ *AFn) *AFn {
	return Fn(run_BANG_, func(proc interface{}, coll interface{}) interface{} {
		return Reduce.X_invoke_Arity3(func(G__1115 *AFn) *AFn {
			return Fn(G__1115, func(p1__1114_SHARP_ interface{}, p2__1113_SHARP_ interface{}) interface{} {
				return proc.(CljsCoreIFn).X_invoke_Arity1(p2__1113_SHARP_).(interface{})
			})
		}(&AFn{}), nil, coll)
	})
}(&AFn{})

type CljsCoreIEncodeJS interface {
	CljsCoreIEncodeJS__()
	X_clj__GT_js_Arity1() interface{}
	X_key__GT_js_Arity1() interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IEncodeJS", (*CljsCoreIEncodeJS)(nil))
}

var X_clj__GT_js = func(_clj__GT_js *AFn) *AFn {
	return Fn(_clj__GT_js, func(x interface{}) interface{} {
		return x.(CljsCoreIEncodeJS).X_clj__GT_js_Arity1()
	})
}(&AFn{})

var X_key__GT_js = func(_key__GT_js *AFn) *AFn {
	return Fn(_key__GT_js, func(x interface{}) interface{} {
		return x.(CljsCoreIEncodeJS).X_key__GT_js_Arity1()
	})
}(&AFn{})

var Key__GT_js = func(key__GT_js *AFn) *AFn {
	return Fn(key__GT_js, func(k interface{}) interface{} {
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IEncodeJS", Str: "cljs.core/IEncodeJS", X_hash: float64(-2022651465), X_meta: nil}), k)) {
			return k.(CljsCoreIEncodeJS).X_clj__GT_js_Arity1().(interface{})
		} else {
			if (reflect.TypeOf(k).Kind() == reflect.String) || (reflect.TypeOf(k).Kind() == reflect.Float64) || (func() bool { _, instanceof := k.(*CljsCoreKeyword); return instanceof }()) || (func() bool { _, instanceof := k.(*CljsCoreSymbol); return instanceof }()) {
				return Clj__GT_js.X_invoke_Arity1(k).(interface{})
			} else {
				return Pr_str.X_invoke_ArityVariadic(k).(string)
			}
		}
	})
}(&AFn{})

/**
* Recursively transforms ClojureScript values to JavaScript.
* sets/vectors/lists become Arrays, Keywords and Symbol become Strings,
* Maps become Objects. Arbitrary keys are encoded to by key->js.
 */
var Clj__GT_js = func(clj__GT_js *AFn) *AFn {
	return Fn(clj__GT_js, func(x interface{}) interface{} {
		if x == nil {
			return nil
		} else {
			if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IEncodeJS", Str: "cljs.core/IEncodeJS", X_hash: float64(-2022651465), X_meta: nil}), x)) {
				return x.(CljsCoreIEncodeJS).X_clj__GT_js_Arity1().(interface{})
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
									var obj1129 = map[string]interface{}{}
									_ = obj1129
									return obj1129
								}()
								_ = m
								{
									var seq__1130_1140 = Seq.Arity1IQ(x)
									var chunk__1131_1141 = nil
									var count__1132_1142 = float64(0)
									var i__1133_1143 = float64(0)
									_, _, _, _ = seq__1130_1140, chunk__1131_1141, count__1132_1142, i__1133_1143
									for {
										if i__1133_1143 < count__1132_1142 {
											{
												var vec__1134_1144 = chunk__1131_1141.X_nth_Arity2(i__1133_1143).(interface{})
												var k_1145 = Nth.X_invoke_Arity3(vec__1134_1144, float64(0), nil)
												var v_1146 = Nth.X_invoke_Arity3(vec__1134_1144, float64(1), nil)
												_, _, _ = vec__1134_1144, k_1145, v_1146
												m.([]interface{})[int(Key__GT_js.X_invoke_Arity1(k_1145).(float64))] = clj__GT_js.X_invoke_Arity1(v_1146)
												{
													var G__1147 = seq__1130_1140
													var G__1148 = chunk__1131_1141
													var G__1149 = count__1132_1142
													var G__1150 = (i__1133_1143 + float64(1))
													seq__1130_1140 = G__1147
													chunk__1131_1141 = G__1148
													count__1132_1142 = G__1149
													i__1133_1143 = G__1150
													continue
												}
											}
										} else {
											{
												var temp__4219__auto___1151 = Seq.X_invoke_Arity1(seq__1130_1140)
												_ = temp__4219__auto___1151
												if Truth_(temp__4219__auto___1151) {
													{
														var seq__1130_1152___1 = temp__4219__auto___1151
														_ = seq__1130_1152___1
														if Chunked_seq_QMARK_.X_invoke_Arity1(seq__1130_1152___1) {
															{
																var c__77428__auto___1153 = Chunk_first.X_invoke_Arity1(seq__1130_1152___1).(interface{})
																_ = c__77428__auto___1153
																{
																	var G__1154 = Chunk_rest.X_invoke_Arity1(seq__1130_1152___1).(interface{})
																	var G__1155 = c__77428__auto___1153
																	var G__1156 = Count.X_invoke_Arity1(c__77428__auto___1153).(float64)
																	var G__1157 = float64(0)
																	seq__1130_1140 = G__1154
																	chunk__1131_1141 = G__1155
																	count__1132_1142 = G__1156
																	i__1133_1143 = G__1157
																	continue
																}
															}
														} else {
															{
																var vec__1135_1158 = First.X_invoke_Arity1(seq__1130_1152___1)
																var k_1159 = Nth.X_invoke_Arity3(vec__1135_1158, float64(0), nil)
																var v_1160 = Nth.X_invoke_Arity3(vec__1135_1158, float64(1), nil)
																_, _, _ = vec__1135_1158, k_1159, v_1160
																m.([]interface{})[int(Key__GT_js.X_invoke_Arity1(k_1159).(float64))] = clj__GT_js.X_invoke_Arity1(v_1160)
																{
																	var G__1161 = Next.X_invoke_Arity1(seq__1130_1152___1)
																	var G__1162 = nil
																	var G__1163 = float64(0)
																	var G__1164 = float64(0)
																	seq__1130_1140 = G__1161
																	chunk__1131_1141 = G__1162
																	count__1132_1142 = G__1163
																	i__1133_1143 = G__1164
																	continue
																}
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
										var seq__1136_1165 = Seq.X_invoke_Arity1(Map_.X_invoke_Arity2(clj__GT_js, x).(*CljsCoreLazySeq))
										var chunk__1137_1166 = nil
										var count__1138_1167 = float64(0)
										var i__1139_1168 = float64(0)
										_, _, _, _ = seq__1136_1165, chunk__1137_1166, count__1138_1167, i__1139_1168
										for {
											if i__1139_1168 < count__1138_1167 {
												{
													var x_1169___1 = chunk__1137_1166.X_nth_Arity2(i__1139_1168).(interface{})
													_ = x_1169___1
													arr.Push(x_1169___1)
													{
														var G__1170 = seq__1136_1165
														var G__1171 = chunk__1137_1166
														var G__1172 = count__1138_1167
														var G__1173 = (i__1139_1168 + float64(1))
														seq__1136_1165 = G__1170
														chunk__1137_1166 = G__1171
														count__1138_1167 = G__1172
														i__1139_1168 = G__1173
														continue
													}
												}
											} else {
												{
													var temp__4219__auto___1174 = Seq.X_invoke_Arity1(seq__1136_1165)
													_ = temp__4219__auto___1174
													if Truth_(temp__4219__auto___1174) {
														{
															var seq__1136_1175___1 = temp__4219__auto___1174
															_ = seq__1136_1175___1
															if Chunked_seq_QMARK_.X_invoke_Arity1(seq__1136_1175___1) {
																{
																	var c__77428__auto___1176 = Chunk_first.X_invoke_Arity1(seq__1136_1175___1).(interface{})
																	_ = c__77428__auto___1176
																	{
																		var G__1177 = Chunk_rest.X_invoke_Arity1(seq__1136_1175___1).(interface{})
																		var G__1178 = c__77428__auto___1176
																		var G__1179 = Count.X_invoke_Arity1(c__77428__auto___1176).(float64)
																		var G__1180 = float64(0)
																		seq__1136_1165 = G__1177
																		chunk__1137_1166 = G__1178
																		count__1138_1167 = G__1179
																		i__1139_1168 = G__1180
																		continue
																	}
																}
															} else {
																{
																	var x_1181___1 = First.X_invoke_Arity1(seq__1136_1175___1)
																	_ = x_1181___1
																	arr.Push(x_1181___1)
																	{
																		var G__1182 = Next.X_invoke_Arity1(seq__1136_1175___1)
																		var G__1183 = nil
																		var G__1184 = float64(0)
																		var G__1185 = float64(0)
																		seq__1136_1165 = G__1182
																		chunk__1137_1166 = G__1183
																		count__1138_1167 = G__1184
																		i__1139_1168 = G__1185
																		continue
																	}
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

type CljsCoreIEncodeClojure interface {
	CljsCoreIEncodeClojure__()
	X_js__GT_clj_Arity2(options interface{}) interface{}
}

func init() {
	RegisterProtocol_("cljs.core/IEncodeClojure", (*CljsCoreIEncodeClojure)(nil))
}

var X_js__GT_clj = func(_js__GT_clj *AFn) *AFn {
	return Fn(_js__GT_clj, func(x interface{}, options interface{}) interface{} {
		return x.(CljsCoreIEncodeClojure).X_js__GT_clj_Arity2(options)
	})
}(&AFn{})

/**
* Recursively transforms JavaScript arrays into ClojureScript
* vectors, and JavaScript objects into ClojureScript maps.  With
* option ':keywordize-keys true' will convert object fields from
* strings to keywords.
* @param {...*} var_args
 */
var Js__GT_clj = func(js__GT_clj *AFn) *AFn {
	return Fn(js__GT_clj, func(x interface{}) interface{} {
		return js__GT_clj.X_invoke_ArityVariadic(x, &CljsCorePersistentArrayMap{nil, 1, []interface{}{(&CljsCoreKeyword{Ns: nil, Name: "keywordize-keys", Fqn: "keywordize-keys", X_hash: float64(1310784252)}), false}, nil})
	}, func(x_opts__ ...interface{}) interface{} {
		var x = x_opts__[0]
		var opts = Array_seq.X_invoke_Arity1(x_opts__[1:])
		_, _ = x, opts
		if Truth_(Native_satisfies_QMARK_.X_invoke_Arity2((&CljsCoreSymbol{Ns: "cljs.core", Name: "IEncodeClojure", Str: "cljs.core/IEncodeClojure", X_hash: float64(1514086517), X_meta: nil}), x)) {
			return x.(CljsCoreIEncodeClojure).X_js__GT_clj_Arity2(Apply.X_invoke_Arity2(Array_map, opts)).(interface{})
		} else {
			if Truth_(Seq.Arity1IQ(opts)) {
				{
					var map__1195 = opts
					var map__1195___1 = func() interface{} {
						if Seq_QMARK_.Arity1IB(map__1195) {
							return Apply.X_invoke_Arity2(Hash_map, map__1195)
						} else {
							return map__1195
						}
					}()
					var keywordize_keys = Get.X_invoke_Arity2(map__1195___1, (&CljsCoreKeyword{Ns: nil, Name: "keywordize-keys", Fqn: "keywordize-keys", X_hash: float64(1310784252)}))
					var keyfn = func() interface{} {
						if Truth_(keywordize_keys.(bool)) {
							return Keyword
						} else {
							return Str
						}
					}()
					var f = func(map__1195 interface{}, map__1195___1 interface{}, keywordize_keys interface{}, keyfn interface{}) *AFn {
						return func(thisfn *AFn) *AFn {
							return Fn(thisfn, func(x___1 interface{}) interface{} {
								if Seq_QMARK_.Arity1IB(x___1) {
									return Doall.X_invoke_Arity1(Map_.X_invoke_Arity2(thisfn, x___1).(*CljsCoreLazySeq)).(interface{})
								} else {
									if Coll_QMARK_.Arity1IB(x___1) {
										return Into.X_invoke_Arity2(Empty.X_invoke_Arity1(x___1), Map_.X_invoke_Arity2(thisfn, x___1).(*CljsCoreLazySeq))
									} else {
										if reflect.TypeOf(x___1).Kind() == reflect.Slice {
											return Vec.X_invoke_Arity1(Map_.X_invoke_Arity2(thisfn, x___1).(*CljsCoreLazySeq)).(interface{})
										} else {
											if Type_.X_invoke_Arity1(x___1) == js.Object {
												return Into.X_invoke_Arity2(CljsCorePersistentArrayMap_EMPTY, func() interface{} {
													var iter__77397__auto__ = func(map__1195 interface{}, map__1195___1 interface{}, keywordize_keys interface{}, keyfn interface{}) *AFn {
														return func(iter__1200 *AFn) *AFn {
															return Fn(iter__1200, func(s__1201 interface{}) interface{} {
																return (&CljsCoreLazySeq{nil, func(map__1195 interface{}, map__1195___1 interface{}, keywordize_keys interface{}, keyfn interface{}) *AFn {
																	return func(G__1204 *AFn) *AFn {
																		return Fn(G__1204, func() interface{} {
																			{
																				var s__1201___1 = s__1201
																				_ = s__1201___1
																				for {
																					{
																						var temp__4219__auto__ = Seq.Arity1IQ(s__1201___1)
																						_ = temp__4219__auto__
																						if Truth_(temp__4219__auto__) {
																							{
																								var s__1201___2 = temp__4219__auto__
																								_ = s__1201___2
																								if Chunked_seq_QMARK_.X_invoke_Arity1(s__1201___2) {
																									{
																										var c__77395__auto__ = Chunk_first.X_invoke_Arity1(s__1201___2).(CljsCoreNot_native)
																										var size__77396__auto__ = Count.X_invoke_Arity1(c__77395__auto__).(float64)
																										var b__1203 = Chunk_buffer.X_invoke_Arity1(size__77396__auto__).(*CljsCoreChunkBuffer)
																										_, _, _ = c__77395__auto__, size__77396__auto__, b__1203
																										if func() interface{} {
																											var i__1202 = float64(0)
																											_ = i__1202
																											for {
																												if i__1202 < size__77396__auto__ {
																													{
																														var k = c__77395__auto__.X_nth_Arity2(i__1202).(interface{})
																														_ = k
																														Chunk_append.X_invoke_Arity2(b__1203, &CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{keyfn.(CljsCoreIFn).X_invoke_Arity1(k).(interface{}), thisfn.X_invoke_Arity1((x___1.([]interface{})[int(k.(float64))]))}, nil}).(interface{})
																														{
																															var G__1205 = (i__1202 + float64(1))
																															i__1202 = G__1205
																															continue
																														}
																													}
																												} else {
																													return true
																												}
																											}
																										}() {
																											return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b__1203).(interface{}), iter__1200.X_invoke_Arity1(Chunk_rest.X_invoke_Arity1(s__1201___2).(interface{})).(*CljsCoreLazySeq))
																										} else {
																											return Chunk_cons.X_invoke_Arity2(Chunk.X_invoke_Arity1(b__1203).(interface{}), nil)
																										}
																									}
																								} else {
																									{
																										var k = First.X_invoke_Arity1(s__1201___2)
																										_ = k
																										return Cons.X_invoke_Arity2(&CljsCorePersistentVector{nil, 2, 5, CljsCorePersistentVector_EMPTY_NODE, []interface{}{keyfn.(CljsCoreIFn).X_invoke_Arity1(k).(interface{}), thisfn.X_invoke_Arity1((x___1.([]interface{})[int(k.(float64))]))}, nil}, iter__1200.X_invoke_Arity1(Rest.X_invoke_Arity1(s__1201___2)).(*CljsCoreLazySeq)).(*CljsCoreCons)
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
																	}(&AFn{})
																}(map__1195, map__1195___1, keywordize_keys, keyfn), nil, nil})
															})
														}(&AFn{})
													}(map__1195, map__1195___1, keywordize_keys, keyfn)
													_ = iter__77397__auto__
													return iter__77397__auto__.(CljsCoreIFn).X_invoke_Arity1(Js_keys.X_invoke_Arity1(x___1).([]interface{})).(*CljsCoreLazySeq)
												}())
											} else {
												return x___1

											}
										}
									}
								}
							})
						}(&AFn{})
					}(map__1195, map__1195___1, keywordize_keys, keyfn)
					_, _, _, _, _ = map__1195, map__1195___1, keywordize_keys, keyfn, f
					return f.(CljsCoreIFn).X_invoke_Arity1(x)
				}
			} else {
				return nil
			}
		}
	})
}(&AFn{})

/**
* Returns a memoized version of a referentially transparent function. The
* memoized version of the function keeps a cache of the mapping from arguments
* to results and, when calls with the same arguments are repeated often, has
* higher performance at the expense of higher memory use.
 */
var Memoize = func(memoize *AFn) *AFn {
	return Fn(memoize, func(f interface{}) interface{} {
		{
			var mem = Atom.X_invoke_Arity1(CljsCorePersistentArrayMap_EMPTY).(interface{})
			_ = mem
			return func(mem interface{}) *AFn {
				return func(G__1206 *AFn) *AFn {
					return Fn(G__1206, func(args__ ...interface{}) interface{} {
						var args = Array_seq.X_invoke_Arity1(args__[0:])
						_ = args
						{
							var v = Get.X_invoke_Arity3(Deref.X_invoke_Arity1(mem).(interface{}), args, Lookup_sentinel)
							_ = v
							if v == Lookup_sentinel {
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
				}(&AFn{})
			}(mem)
		}
	})
}(&AFn{})

/**
* trampoline can be used to convert algorithms requiring mutual
* recursion without stack consumption. Calls f with supplied args, if
* any. If f returns a fn, calls that fn with no arguments, and
* continues to repeat, until the return value is not a fn, then
* returns that non-fn value. Note that if you want to return a fn as a
* final value, you must wrap it in some data structure and unpack it
* after trampoline returns.
* @param {...*} var_args
 */
var Trampoline = func(trampoline *AFn) *AFn {
	return Fn(trampoline, func(f interface{}) interface{} {
		for {
			{
				var ret = f.(CljsCoreIFn).X_invoke_Arity0().(interface{})
				_ = ret
				if Fn_QMARK_.X_invoke_Arity1(ret) {
					{
						var G__1207 = ret
						f = G__1207
						continue
					}
				} else {
					return ret
				}
			}
		}
	}, func(f_args__ ...interface{}) interface{} {
		var f = f_args__[0]
		var args = Array_seq.X_invoke_Arity1(f_args__[1:])
		_, _ = f, args
		return trampoline.X_invoke_Arity1(func(G__1208 *AFn) *AFn {
			return Fn(G__1208, func() interface{} {
				return Apply.X_invoke_Arity2(f, args)
			})
		}(&AFn{})).(interface{})
	})
}(&AFn{})

func init() {
	/**
	 * Returns a random floating point number between 0 (inclusive) and
	 * n (default 1) (exclusive).
	 */
	Rand = func(rand *AFn) *AFn {
		return Fn(rand, func() interface{} {
			return rand.X_invoke_Arity1(float64(1)).(float64)
		}, func(n interface{}) interface{} {
			return (Math.Random().(float64) * n.(float64))
		})
	}(&AFn{})

}
func init() {
	/**
	 * Returns a random integer between 0 (inclusive) and n (exclusive).
	 */
	Rand_int = func(rand_int *AFn) *AFn {
		return Fn(rand_int, func(n interface{}) interface{} {
			return Math.Floor((Math.Random().(float64) * n.(float64)))
		})
	}(&AFn{})

}

/**
* Return a random element of the (sequential) collection. Will have
* the same performance characteristics as nth for the given
* collection.
 */
var Rand_nth = func(rand_nth *AFn) *AFn {
	return Fn(rand_nth, func(coll interface{}) interface{} {
		return Nth.X_invoke_Arity2(coll, Rand_int.X_invoke_Arity1(Count.X_invoke_Arity1(coll).(float64)).(interface{}))
	})
}(&AFn{})

/**
* Returns a map of the elements of coll keyed by the result of
* f on each element. The value at each key will be a vector of the
* corresponding elements, in the order they appeared in coll.
 */
var Group_by = func(group_by *AFn) *AFn {
	return Fn(group_by, func(f interface{}, coll interface{}) interface{} {
		return Reduce.X_invoke_Arity3(func(G__1209 *AFn) *AFn {
			return Fn(G__1209, func(ret interface{}, x interface{}) interface{} {
				{
					var k = f.(CljsCoreIFn).X_invoke_Arity1(x).(interface{})
					_ = k
					return Assoc.X_invoke_Arity3(ret, k, Conj.X_invoke_Arity2(Get.X_invoke_Arity3(ret, k, CljsCorePersistentVector_EMPTY), x).(interface{})).(interface{})
				}
			})
		}(&AFn{}), CljsCorePersistentArrayMap_EMPTY, coll)
	})
}(&AFn{})

/**
* Creates a hierarchy object for use with derive, isa? etc.
 */
var Make_hierarchy = func(make_hierarchy *AFn) *AFn {
	return Fn(make_hierarchy, func() interface{} {
		return &CljsCorePersistentArrayMap{nil, 3, []interface{}{(&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}), CljsCorePersistentArrayMap_EMPTY, (&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}), CljsCorePersistentArrayMap_EMPTY, (&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}), CljsCorePersistentArrayMap_EMPTY}, nil}
	})
}(&AFn{})

var X_global_hierarchy interface{} = nil

var Get_global_hierarchy = func(get_global_hierarchy *AFn) *AFn {
	return Fn(get_global_hierarchy, func() interface{} {
		if X_global_hierarchy == nil {
			X_global_hierarchy = Atom.X_invoke_Arity1(Make_hierarchy.X_invoke_Arity0().(CljsCoreIMap)).(interface{})

		} else {
		}
		return X_global_hierarchy
	})
}(&AFn{})

/**
* @param {...*} var_args
 */
var Swap_global_hierarchy_BANG_ = func(swap_global_hierarchy_BANG_ *AFn) *AFn {
	return Fn(swap_global_hierarchy_BANG_, func(f_args__ ...interface{}) interface{} {
		var f = f_args__[0]
		var args = Array_seq.X_invoke_Arity1(f_args__[1:])
		_, _ = f, args
		return Apply.X_invoke_Arity4(Swap_BANG_, Get_global_hierarchy.X_invoke_Arity0().(interface{}), f, args)
	})
}(&AFn{})

/**
* Returns true if (= child parent), or child is directly or indirectly derived from
* parent, either via a JavaScript type inheritance relationship or a
* relationship established via derive. h must be a hierarchy obtained
* from make-hierarchy, if not supplied defaults to the global
* hierarchy
 */
var Isa_QMARK_ = func(isa_QMARK_ *AFn) *AFn {
	return Fn(isa_QMARK_, func(child interface{}, parent interface{}) bool {
		return isa_QMARK_.X_invoke_Arity3(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0().(interface{})).(interface{}), child, parent)
	}, func(h interface{}, child interface{}, parent interface{}) bool {
		{
			var or__76632__auto__ = X_EQ_.Arity2IIB(child, parent)
			_ = or__76632__auto__
			if or__76632__auto__ {
				return or__76632__auto__
			} else {
				{
					var or__76632__auto_____1 = Contains_QMARK_.X_invoke_Arity2((&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h).(interface{}).X_invoke_Arity1(child).(interface{}), parent)
					_ = or__76632__auto_____1
					if or__76632__auto_____1 {
						return or__76632__auto_____1
					} else {
						{
							var and__76620__auto__ = Vector_QMARK_.Arity1IB(parent)
							_ = and__76620__auto__
							if and__76620__auto__ {
								{
									var and__76620__auto_____1 = Vector_QMARK_.Arity1IB(child)
									_ = and__76620__auto_____1
									if and__76620__auto_____1 {
										{
											var and__76620__auto_____2 = (Count.X_invoke_Arity1(parent).(float64) == Count.X_invoke_Arity1(child).(float64))
											_ = and__76620__auto_____2
											if and__76620__auto_____2 {
												{
													var ret = true
													var i = float64(0)
													_, _ = ret, i
													for {
														if (!(ret)) || (i == Count.X_invoke_Arity1(parent).(float64)) {
															return ret
														} else {
															{
																var G__1210 = isa_QMARK_.X_invoke_Arity3(h, child.(CljsCoreIFn).X_invoke_Arity1(i).(interface{}), parent.(CljsCoreIFn).X_invoke_Arity1(i).(interface{}))
																var G__1211 = (i + float64(1))
																ret = G__1210
																i = G__1211
																continue
															}
														}
													}
												}
											} else {
												return and__76620__auto_____2
											}
										}
									} else {
										return and__76620__auto_____1
									}
								}
							} else {
								return and__76620__auto__
							}
						}
					}
				}
			}
		}
	})
}(&AFn{})

/**
* Returns the immediate parents of tag, either via a JavaScript type
* inheritance relationship or a relationship established via derive. h
* must be a hierarchy obtained from make-hierarchy, if not supplied
* defaults to the global hierarchy
 */
var Parents = func(parents *AFn) *AFn {
	return Fn(parents, func(tag interface{}) interface{} {
		return parents.X_invoke_Arity2(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0().(interface{})).(interface{}), tag)
	}, func(h interface{}, tag interface{}) interface{} {
		return Not_empty.X_invoke_Arity1(Get.X_invoke_Arity2((&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h).(interface{}), tag))
	})
}(&AFn{})

/**
* Returns the immediate and indirect parents of tag, either via a JavaScript type
* inheritance relationship or a relationship established via derive. h
* must be a hierarchy obtained from make-hierarchy, if not supplied
* defaults to the global hierarchy
 */
var Ancestors = func(ancestors *AFn) *AFn {
	return Fn(ancestors, func(tag interface{}) interface{} {
		return ancestors.X_invoke_Arity2(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0().(interface{})).(interface{}), tag)
	}, func(h interface{}, tag interface{}) interface{} {
		return Not_empty.X_invoke_Arity1(Get.X_invoke_Arity2((&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h).(interface{}), tag))
	})
}(&AFn{})

/**
* Returns the immediate and indirect children of tag, through a
* relationship established via derive. h must be a hierarchy obtained
* from make-hierarchy, if not supplied defaults to the global
* hierarchy. Note: does not work on JavaScript type inheritance
* relationships.
 */
var Descendants = func(descendants *AFn) *AFn {
	return Fn(descendants, func(tag interface{}) interface{} {
		return descendants.X_invoke_Arity2(Deref.X_invoke_Arity1(Get_global_hierarchy.X_invoke_Arity0().(interface{})).(interface{}), tag)
	}, func(h interface{}, tag interface{}) interface{} {
		return Not_empty.X_invoke_Arity1(Get.X_invoke_Arity2((&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}).X_invoke_Arity1(h).(interface{}), tag))
	})
}(&AFn{})

/**
* Establishes a parent/child relationship between parent and
* tag. Parent must be a namespace-qualified symbol or keyword and
* child can be either a namespace-qualified symbol or keyword or a
* class. h must be a hierarchy obtained from make-hierarchy, if not
* supplied defaults to, and modifies, the global hierarchy.
 */
var Derive = func(derive *AFn) *AFn {
	return Fn(derive, func(tag interface{}, parent interface{}) interface{} {
		if Truth_(Namespace.X_invoke_Arity1(parent).(string)) {
		} else {
			panic((&js.Error{("Assert failed: " + Str.X_invoke_Arity1(Pr_str.X_invoke_ArityVariadic(List((&CljsCoreSymbol{Ns: nil, Name: "namespace", Str: "namespace", X_hash: float64(1263021155), X_meta: nil}), (&CljsCoreSymbol{Ns: nil, Name: "parent", Str: "parent", X_hash: float64(761652748), X_meta: nil}))).(string)).(string))}))
		}
		Swap_global_hierarchy_BANG_.X_invoke_ArityVariadic(derive, tag, parent)
		return nil
	}, func(h interface{}, tag interface{}, parent interface{}) interface{} {
		if Not_EQ_.Arity2IIB(tag, parent) {
		} else {
			panic((&js.Error{("Assert failed: " + Str.X_invoke_Arity1(Pr_str.X_invoke_ArityVariadic(List((&CljsCoreSymbol{Ns: nil, Name: "not=", Str: "not=", X_hash: float64(1466536204), X_meta: nil}), (&CljsCoreSymbol{Ns: nil, Name: "tag", Str: "tag", X_hash: float64(350170304), X_meta: nil}), (&CljsCoreSymbol{Ns: nil, Name: "parent", Str: "parent", X_hash: float64(761652748), X_meta: nil}))).(string)).(string))}))
		}
		{
			var tp = (&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h).(interface{})
			var td = (&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}).X_invoke_Arity1(h).(interface{})
			var ta = (&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h).(interface{})
			var tf = func(tp interface{}, td interface{}, ta interface{}) *AFn {
				return func(G__1212 *AFn) *AFn {
					return Fn(G__1212, func(m interface{}, source interface{}, sources interface{}, target interface{}, targets interface{}) interface{} {
						return Reduce.X_invoke_Arity3(func(tp interface{}, td interface{}, ta interface{}) *AFn {
							return func(G__1213 *AFn) *AFn {
								return Fn(G__1213, func(ret interface{}, k interface{}) interface{} {
									return Assoc.X_invoke_Arity3(ret, k, Reduce.X_invoke_Arity3(Conj, Get.X_invoke_Arity3(targets, k, CljsCorePersistentHashSet_EMPTY), Cons.X_invoke_Arity2(target, targets.(CljsCoreIFn).X_invoke_Arity1(target).(interface{})).(*CljsCoreCons))).(interface{})
								})
							}(&AFn{})
						}(tp, td, ta), m, Cons.X_invoke_Arity2(source, sources.(CljsCoreIFn).X_invoke_Arity1(source).(interface{})).(*CljsCoreCons))
					})
				}(&AFn{})
			}(tp, td, ta)
			_, _, _, _ = tp, td, ta, tf
			{
				var or__76632__auto__ = func() interface{} {
					if Contains_QMARK_.X_invoke_Arity2(tp.(CljsCoreIFn).X_invoke_Arity1(tag).(interface{}), parent) {
						return nil
					} else {
						return func() interface{} {
							if Contains_QMARK_.X_invoke_Arity2(ta.(CljsCoreIFn).X_invoke_Arity1(tag).(interface{}), parent) {
								panic((&js.Error{(`` + Str.X_invoke_Arity1(tag).(string) + "already has" + Str.X_invoke_Arity1(parent).(string) + "as ancestor")}))
							} else {
							}
							if Contains_QMARK_.X_invoke_Arity2(ta.(CljsCoreIFn).X_invoke_Arity1(parent).(interface{}), tag) {
								panic((&js.Error{("Cyclic derivation:" + Str.X_invoke_Arity1(parent).(string) + "has" + Str.X_invoke_Arity1(tag).(string) + "as ancestor")}))
							} else {
							}
							return &CljsCorePersistentArrayMap{nil, 3, []interface{}{(&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}), Assoc.X_invoke_Arity3((&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h).(interface{}), tag, Conj.X_invoke_Arity2(Get.X_invoke_Arity3(tp, tag, CljsCorePersistentHashSet_EMPTY), parent).(interface{})).(interface{}), (&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}), tf.(CljsCoreIFn).X_invoke_Arity5((&CljsCoreKeyword{Ns: nil, Name: "ancestors", Fqn: "ancestors", X_hash: float64(-776045424)}).X_invoke_Arity1(h).(interface{}), tag, td, parent, ta), (&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}), tf.(CljsCoreIFn).X_invoke_Arity5((&CljsCoreKeyword{Ns: nil, Name: "descendants", Fqn: "descendants", X_hash: float64(1824886031)}).X_invoke_Arity1(h).(interface{}), parent, ta, tag, td)}, nil}
						}()
					}
				}()
				_ = or__76632__auto__
				if Truth_(or__76632__auto__) {
					return or__76632__auto__
				} else {
					return h
				}
			}
		}
	})
}(&AFn{})

/**
* Removes a parent/child relationship between parent and
* tag. h must be a hierarchy obtained from make-hierarchy, if not
* supplied defaults to, and modifies, the global hierarchy.
 */
var Underive = func(underive *AFn) *AFn {
	return Fn(underive, func(tag interface{}, parent interface{}) interface{} {
		Swap_global_hierarchy_BANG_.X_invoke_ArityVariadic(underive, tag, parent)
		return nil
	}, func(h interface{}, tag interface{}, parent interface{}) interface{} {
		{
			var parentMap = (&CljsCoreKeyword{Ns: nil, Name: "parents", Fqn: "parents", X_hash: float64(-2027538891)}).X_invoke_Arity1(h).(interface{})
			var childsParents = func() interface{} {
				if Truth_(parentMap.(CljsCoreIFn).X_invoke_Arity1(tag).(interface{}).(bool)) {
					return Disj.X_invoke_Arity2(parentMap.(CljsCoreIFn).X_invoke_Arity1(tag).(interface{}), parent)
				} else {
					return CljsCorePersistentHashSet_EMPTY
				}
			}()
			var newParents = func() interface{} {
				if Truth_(Not_empty.X_invoke_Arity1(childsParents).(bool)) {
					return Assoc.X_invoke_Arity3(parentMap, tag, childsParents).(interface{})
				} else {
					return Dissoc.X_invoke_Arity2(parentMap, tag)
				}
			}()
			var deriv_seq = Flatten.X_invoke_Arity1(Map_.X_invoke_Arity2(func(parentMap interface{}, childsParents interface{}, newParents interface{}) *AFn {
				return func(G__1217 *AFn) *AFn {
					return Fn(G__1217, func(p1__1214_SHARP_ interface{}) interface{} {
						return Cons.X_invoke_Arity2(First.X_invoke_Arity1(p1__1214_SHARP_), Interpose.X_invoke_Arity2(First.X_invoke_Arity1(p1__1214_SHARP_), Second.X_invoke_Arity1(p1__1214_SHARP_)).(*CljsCoreLazySeq)).(*CljsCoreCons)
					})
				}(&AFn{})
			}(parentMap, childsParents, newParents), Seq.X_invoke_Arity1(newParents)).(*CljsCoreLazySeq)).(*CljsCoreLazySeq)
			_, _, _, _ = parentMap, childsParents, newParents, deriv_seq
			if Contains_QMARK_.X_invoke_Arity2(parentMap.(CljsCoreIFn).X_invoke_Arity1(tag).(interface{}), parent) {
				return Reduce.X_invoke_Arity3(func(parentMap interface{}, childsParents interface{}, newParents interface{}, deriv_seq *CljsCoreLazySeq) *AFn {
					return func(G__1218 *AFn) *AFn {
						return Fn(G__1218, func(p1__1215_SHARP_ interface{}, p2__1216_SHARP_ interface{}) interface{} {
							return Apply.X_invoke_Arity3(Derive, p1__1215_SHARP_, p2__1216_SHARP_)
						})
					}(&AFn{})
				}(parentMap, childsParents, newParents, deriv_seq), Make_hierarchy.X_invoke_Arity0().(CljsCoreIMap), Partition.X_invoke_Arity2(float64(2), deriv_seq).(*CljsCoreLazySeq))
			} else {
				return h
			}
		}
	})
}(&AFn{})

var Reset_cache = func(reset_cache *AFn) *AFn {
	return Fn(reset_cache, func(method_cache interface{}, method_table interface{}, cached_hierarchy interface{}, hierarchy interface{}) interface{} {
		Swap_BANG_.X_invoke_Arity2(method_cache, func(G__1219 *AFn) *AFn {
			return Fn(G__1219, func(___ interface{}) interface{} {
				return Deref.X_invoke_Arity1(method_table).(interface{})
			})
		}(&AFn{}))
		return Swap_BANG_.X_invoke_Arity2(cached_hierarchy, func(G__1220 *AFn) *AFn {
			return Fn(G__1220, func(___ interface{}) interface{} {
				return Deref.X_invoke_Arity1(hierarchy).(interface{})
			})
		}(&AFn{}))
	})
}(&AFn{})

var Prefers_STAR_ = func(prefers_STAR_ *AFn) *AFn {
	return Fn(prefers_STAR_, func(x interface{}, y interface{}, prefer_table interface{}) interface{} {
		{
			var xprefs = Deref.X_invoke_Arity1(prefer_table).(interface{}).X_invoke_Arity1(x).(interface{})
			_ = xprefs
			{
				var or__76632__auto__ = func() interface{} {
					if Truth_(func() interface{} {
						var and__76620__auto__ = xprefs
						_ = and__76620__auto__
						if Truth_(and__76620__auto__) {
							return xprefs.(CljsCoreIFn).X_invoke_Arity1(y).(interface{})
						} else {
							return and__76620__auto__
						}
					}().(bool)) {
						return true
					} else {
						return nil
					}
				}()
				_ = or__76632__auto__
				if Truth_(or__76632__auto__) {
					return or__76632__auto__
				} else {
					{
						var or__76632__auto_____1 = func() interface{} {
							var ps = Parents.X_invoke_Arity1(y)
							_ = ps
							for {
								if Count.X_invoke_Arity1(ps).(float64) > float64(0) {
									if Truth_(prefers_STAR_.X_invoke_Arity3(x, First.X_invoke_Arity1(ps), prefer_table)) {
									} else {
									}
									{
										var G__1221 = Rest.X_invoke_Arity1(ps)
										ps = G__1221
										continue
									}
								} else {
									return nil
								}
							}
						}()
						_ = or__76632__auto_____1
						if Truth_(or__76632__auto_____1) {
							return or__76632__auto_____1
						} else {
							{
								var or__76632__auto_____2 = func() interface{} {
									var ps = Parents.X_invoke_Arity1(x)
									_ = ps
									for {
										if Count.X_invoke_Arity1(ps).(float64) > float64(0) {
											if Truth_(prefers_STAR_.X_invoke_Arity3(First.X_invoke_Arity1(ps), y, prefer_table)) {
											} else {
											}
											{
												var G__1222 = Rest.X_invoke_Arity1(ps)
												ps = G__1222
												continue
											}
										} else {
											return nil
										}
									}
								}()
								_ = or__76632__auto_____2
								if Truth_(or__76632__auto_____2) {
									return or__76632__auto_____2
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

var Dominates = func(dominates *AFn) *AFn {
	return Fn(dominates, func(x interface{}, y interface{}, prefer_table interface{}) interface{} {
		{
			var or__76632__auto__ = Prefers_STAR_.X_invoke_Arity3(x, y, prefer_table)
			_ = or__76632__auto__
			if Truth_(or__76632__auto__) {
				return or__76632__auto__
			} else {
				return Isa_QMARK_.Arity2IIB(x, y)
			}
		}
	})
}(&AFn{})

var Find_and_cache_best_method = func(find_and_cache_best_method *AFn) *AFn {
	return Fn(find_and_cache_best_method, func(name interface{}, dispatch_val interface{}, hierarchy interface{}, method_table interface{}, prefer_table interface{}, method_cache interface{}, cached_hierarchy interface{}) interface{} {
		{
			var best_entry = Reduce.X_invoke_Arity3(func(G__1227 *AFn) *AFn {
				return Fn(G__1227, func(be interface{}, p__1225 interface{}) interface{} {
					{
						var vec__1226 = p__1225
						var k = Nth.X_invoke_Arity3(vec__1226, float64(0), nil)
						var ___ = Nth.X_invoke_Arity3(vec__1226, float64(1), nil)
						var e = vec__1226
						_, _, _, _ = vec__1226, k, ___, e
						if Isa_QMARK_.X_invoke_Arity3(Deref.X_invoke_Arity1(hierarchy).(interface{}), dispatch_val, k) {
							{
								var be2 = func() interface{} {
									if Truth_(func() interface{} {
										var or__76632__auto__ = (be == nil)
										_ = or__76632__auto__
										if or__76632__auto__ {
											return or__76632__auto__
										} else {
											return Dominates.X_invoke_Arity3(k, First.X_invoke_Arity1(be), prefer_table)
										}
									}().(bool)) {
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
			}(&AFn{}), nil, Deref.X_invoke_Arity1(method_table).(interface{}))
			_ = best_entry
			if Truth_(best_entry) {
				if X_EQ_.X_invoke_Arity2(Deref.X_invoke_Arity1(cached_hierarchy).(interface{}), Deref.X_invoke_Arity1(hierarchy).(interface{})) {
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

var X_reset = func(_reset *AFn) *AFn {
	return Fn(_reset, func(mf interface{}) interface{} {
		return mf.(CljsCoreIMultiFn).X_reset_Arity1()
	})
}(&AFn{})

var X_add_method = func(_add_method *AFn) *AFn {
	return Fn(_add_method, func(mf interface{}, dispatch_val interface{}, method interface{}) interface{} {
		return mf.(CljsCoreIMultiFn).X_add_method_Arity3(dispatch_val, method)
	})
}(&AFn{})

var X_remove_method = func(_remove_method *AFn) *AFn {
	return Fn(_remove_method, func(mf interface{}, dispatch_val interface{}) interface{} {
		return mf.(CljsCoreIMultiFn).X_remove_method_Arity2(dispatch_val)
	})
}(&AFn{})

var X_prefer_method = func(_prefer_method *AFn) *AFn {
	return Fn(_prefer_method, func(mf interface{}, dispatch_val interface{}, dispatch_val_y interface{}) interface{} {
		return mf.(CljsCoreIMultiFn).X_prefer_method_Arity3(dispatch_val, dispatch_val_y)
	})
}(&AFn{})

var X_get_method = func(_get_method *AFn) *AFn {
	return Fn(_get_method, func(mf interface{}, dispatch_val interface{}) interface{} {
		return mf.(CljsCoreIMultiFn).X_get_method_Arity2(dispatch_val)
	})
}(&AFn{})

var X_methods = func(_methods *AFn) *AFn {
	return Fn(_methods, func(mf interface{}) interface{} {
		return mf.(CljsCoreIMultiFn).X_methods_Arity1()
	})
}(&AFn{})

var X_prefers = func(_prefers *AFn) *AFn {
	return Fn(_prefers, func(mf interface{}) interface{} {
		return mf.(CljsCoreIMultiFn).X_prefers_Arity1()
	})
}(&AFn{})

var Throw_no_method_error = func(throw_no_method_error *AFn) *AFn {
	return Fn(throw_no_method_error, func(name interface{}, dispatch_val interface{}) interface{} {
		panic((&js.Error{("No method in multimethod '" + Str.X_invoke_Arity1(name).(string) + "' for dispatch value: " + Str.X_invoke_Arity1(dispatch_val).(string))}))
	})
}(&AFn{})

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
func (self__ *CljsCoreMultiFn) X_hash_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		return goog.GetUid(this___1)
	}
}

func (_ *CljsCoreMultiFn) CljsCoreIMultiFn__() {}
func (self__ *CljsCoreMultiFn) X_reset_Arity1() interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		Swap_BANG_.X_invoke_Arity2(self__.Method_table, func(mf___1 *CljsCoreMultiFn) *AFn {
			return func(G__1228 *AFn) *AFn {
				return Fn(G__1228, func(mf___2 interface{}) interface{} {
					return CljsCorePersistentArrayMap_EMPTY
				})
			}(&AFn{})
		}(mf___1))
		Swap_BANG_.X_invoke_Arity2(self__.Method_cache, func(mf___1 *CljsCoreMultiFn) *AFn {
			return func(G__1229 *AFn) *AFn {
				return Fn(G__1229, func(mf___2 interface{}) interface{} {
					return CljsCorePersistentArrayMap_EMPTY
				})
			}(&AFn{})
		}(mf___1))
		Swap_BANG_.X_invoke_Arity2(self__.Prefer_table, func(mf___1 *CljsCoreMultiFn) *AFn {
			return func(G__1230 *AFn) *AFn {
				return Fn(G__1230, func(mf___2 interface{}) interface{} {
					return CljsCorePersistentArrayMap_EMPTY
				})
			}(&AFn{})
		}(mf___1))
		Swap_BANG_.X_invoke_Arity2(self__.Cached_hierarchy, func(mf___1 *CljsCoreMultiFn) *AFn {
			return func(G__1231 *AFn) *AFn {
				return Fn(G__1231, func(mf___2 interface{}) interface{} {
					return nil
				})
			}(&AFn{})
		}(mf___1))
		return mf___1
	}
}

func (self__ *CljsCoreMultiFn) X_add_method_Arity3(dispatch_val interface{}, method interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		Swap_BANG_.X_invoke_Arity4(self__.Method_table, Assoc, dispatch_val, method)
		Reset_cache.X_invoke_Arity4(self__.Method_cache, self__.Method_table, self__.Cached_hierarchy, self__.Hierarchy)
		return mf___1
	}
}

func (self__ *CljsCoreMultiFn) X_remove_method_Arity2(dispatch_val interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		Swap_BANG_.X_invoke_Arity3(self__.Method_table, Dissoc, dispatch_val)
		Reset_cache.X_invoke_Arity4(self__.Method_cache, self__.Method_table, self__.Cached_hierarchy, self__.Hierarchy)
		return mf___1
	}
}

func (self__ *CljsCoreMultiFn) X_get_method_Arity2(dispatch_val interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		if X_EQ_.X_invoke_Arity2(Deref.X_invoke_Arity1(self__.Cached_hierarchy).(interface{}), Deref.X_invoke_Arity1(self__.Hierarchy).(interface{})) {
		} else {
			Reset_cache.X_invoke_Arity4(self__.Method_cache, self__.Method_table, self__.Cached_hierarchy, self__.Hierarchy)
		}
		{
			var temp__4217__auto__ = Deref.X_invoke_Arity1(self__.Method_cache).(interface{}).X_invoke_Arity1(dispatch_val).(interface{})
			_ = temp__4217__auto__
			if Truth_(temp__4217__auto__) {
				{
					var target_fn = temp__4217__auto__
					_ = target_fn
					return target_fn
				}
			} else {
				{
					var temp__4217__auto_____1 = Find_and_cache_best_method.X_invoke_Arity7(self__.Name, dispatch_val, self__.Hierarchy, self__.Method_table, self__.Prefer_table, self__.Method_cache, self__.Cached_hierarchy)
					_ = temp__4217__auto_____1
					if Truth_(temp__4217__auto_____1) {
						{
							var target_fn = temp__4217__auto_____1
							_ = target_fn
							return target_fn
						}
					} else {
						return Deref.X_invoke_Arity1(self__.Method_table).(interface{}).X_invoke_Arity1(self__.Default_dispatch_val).(interface{})
					}
				}
			}
		}
	}
}

func (self__ *CljsCoreMultiFn) X_prefer_method_Arity3(dispatch_val_x interface{}, dispatch_val_y interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		if Truth_(Prefers_STAR_.X_invoke_Arity3(dispatch_val_x, dispatch_val_y, self__.Prefer_table)) {
			panic((&js.Error{("Preference conflict in multimethod '" + Str.X_invoke_Arity1(self__.Name).(string) + "': " + Str.X_invoke_Arity1(dispatch_val_y).(string) + " is already preferred to " + Str.X_invoke_Arity1(dispatch_val_x).(string))}))
		} else {
		}
		Swap_BANG_.X_invoke_Arity2(self__.Prefer_table, func(mf___1 *CljsCoreMultiFn) *AFn {
			return func(G__1232 *AFn) *AFn {
				return Fn(G__1232, func(old interface{}) interface{} {
					return Assoc.X_invoke_Arity3(old, dispatch_val_x, Conj.X_invoke_Arity2(Get.X_invoke_Arity3(old, dispatch_val_x, CljsCorePersistentHashSet_EMPTY), dispatch_val_y).(interface{})).(interface{})
				})
			}(&AFn{})
		}(mf___1))
		return Reset_cache.X_invoke_Arity4(self__.Method_cache, self__.Method_table, self__.Cached_hierarchy, self__.Hierarchy)
	}
}

func (self__ *CljsCoreMultiFn) X_methods_Arity1() interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		return Deref.X_invoke_Arity1(self__.Method_table).(interface{})
	}
}

func (self__ *CljsCoreMultiFn) X_prefers_Arity1() interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		return Deref.X_invoke_Arity1(self__.Prefer_table).(interface{})
	}
}

func (_ *CljsCoreMultiFn) CljsCoreIFn__() {}
func (self__ *CljsCoreMultiFn) X_invoke_Arity0() interface{} {
	{
		var this___1 = self__
		_ = this___1
		panic((&js.Error{"Invalid arity: 0"}))
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity1(a interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity1(a).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity1(a).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity2(a interface{}, b interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity2(a, b).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity2(a, b).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity3(a, b, c).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity3(a, b, c).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity4(a, b, c, d).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity4(a, b, c, d).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity5(a, b, c, d, e).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity5(a, b, c, d, e).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity6(a, b, c, d, e, f).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity6(a, b, c, d, e, f).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity7(a, b, c, d, e, f, g).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity7(a, b, c, d, e, f, g).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity8(a, b, c, d, e, f, g, h).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity8(a, b, c, d, e, f, g, h).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity9(a, b, c, d, e, f, g, h, i).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity9(a, b, c, d, e, f, g, h, i).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity10(a, b, c, d, e, f, g, h, i, j).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity10(a, b, c, d, e, f, g, h, i, j).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity11(a, b, c, d, e, f, g, h, i, j, k).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity11(a, b, c, d, e, f, g, h, i, j, k).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity12(a, b, c, d, e, f, g, h, i, j, k, l).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity12(a, b, c, d, e, f, g, h, i, j, k, l).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity13(a, b, c, d, e, f, g, h, i, j, k, l, m).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity13(a, b, c, d, e, f, g, h, i, j, k, l, m).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity14(a, b, c, d, e, f, g, h, i, j, k, l, m, n).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity14(a, b, c, d, e, f, g, h, i, j, k, l, m, n).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity15(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity15(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity16(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity16(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity17(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity17(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity18(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity18(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity19(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity19(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s).(interface{})
		}
	}
}

func (self__ *CljsCoreMultiFn) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	{
		var mf___1 = self__
		_ = mf___1
		{
			var dispatch_val = self__.Dispatch_fn.(CljsCoreIFn).X_invoke_Arity20(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t).(interface{})
			var target_fn = mf___1.X_get_method_Arity2(dispatch_val).(interface{})
			_, _ = dispatch_val, target_fn
			if Truth_(target_fn) {
			} else {
				Throw_no_method_error.X_invoke_Arity2(self__.Name, dispatch_val).(interface{})
			}
			return target_fn.(CljsCoreIFn).X_invoke_Arity20(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t).(interface{})
		}
	}
}

var X__GT_MultiFn = func(__GT_MultiFn *AFn) *AFn {
	return Fn(__GT_MultiFn, func(name interface{}, dispatch_fn interface{}, default_dispatch_val interface{}, hierarchy interface{}, method_table interface{}, prefer_table interface{}, method_cache interface{}, cached_hierarchy interface{}) interface{} {
		return (&CljsCoreMultiFn{name, dispatch_fn, default_dispatch_val, hierarchy, method_table, prefer_table, method_cache, cached_hierarchy})
	})
}(&AFn{})

/**
* Removes all of the methods of multimethod.
 */
var Remove_all_methods = func(remove_all_methods *AFn) *AFn {
	return Fn(remove_all_methods, func(multifn interface{}) interface{} {
		return multifn.(CljsCoreIMultiFn).X_reset_Arity1().(interface{})
	})
}(&AFn{})

/**
* Removes the method of multimethod associated with dispatch-value.
 */
var Remove_method = func(remove_method *AFn) *AFn {
	return Fn(remove_method, func(multifn interface{}, dispatch_val interface{}) interface{} {
		return multifn.(CljsCoreIMultiFn).X_remove_method_Arity2(dispatch_val).(interface{})
	})
}(&AFn{})

/**
* Causes the multimethod to prefer matches of dispatch-val-x over dispatch-val-y
* when there is a conflict
 */
var Prefer_method = func(prefer_method *AFn) *AFn {
	return Fn(prefer_method, func(multifn interface{}, dispatch_val_x interface{}, dispatch_val_y interface{}) interface{} {
		return multifn.(CljsCoreIMultiFn).X_prefer_method_Arity3(dispatch_val_x, dispatch_val_y).(interface{})
	})
}(&AFn{})

/**
* Given a multimethod, returns a map of dispatch values -> dispatch fns
 */
var Methods = func(methods *AFn) *AFn {
	return Fn(methods, func(multifn interface{}) interface{} {
		return multifn.(CljsCoreIMultiFn).X_methods_Arity1().(interface{})
	})
}(&AFn{})

/**
* Given a multimethod and a dispatch value, returns the dispatch fn
* that would apply to that value, or nil if none apply and no default
 */
var Get_method = func(get_method *AFn) *AFn {
	return Fn(get_method, func(multifn interface{}, dispatch_val interface{}) interface{} {
		return multifn.(CljsCoreIMultiFn).X_get_method_Arity2(dispatch_val).(interface{})
	})
}(&AFn{})

/**
* Given a multimethod, returns a map of preferred value -> set of other values
 */
var Prefers = func(prefers *AFn) *AFn {
	return Fn(prefers, func(multifn interface{}) interface{} {
		return multifn.(CljsCoreIMultiFn).X_prefers_Arity1().(interface{})
	})
}(&AFn{})

type CljsCoreUUID struct{ Uuid interface{} }

func (_ *CljsCoreUUID) CljsCoreIHash__() {}
func (self__ *CljsCoreUUID) X_hash_Arity1() interface{} {
	{
		var this___1 = self__
		_ = this___1
		return goog_string.HashCode(Pr_str.X_invoke_ArityVariadic(this___1).(string))
	}
}

func (_ *CljsCoreUUID) CljsCoreIPrintWithWriter__() {}
func (self__ *CljsCoreUUID) X_pr_writer_Arity3(writer interface{}, ______1 interface{}) interface{} {
	{
		var ______2 = self__
		_ = ______2
		return writer.(CljsCoreIWriter).X_write_Arity2(("#uuid \"" + Str.X_invoke_Arity1(self__.Uuid).(string) + "\"")).(interface{})
	}
}

func (_ *CljsCoreUUID) CljsCoreIEquiv__() {}
func (self__ *CljsCoreUUID) X_equiv_Arity2(other interface{}) bool {
	{
		var ______1 = self__
		_ = ______1
		return (func() bool { _, instanceof := other.(*CljsCoreUUID); return instanceof }()) && (self__.Uuid == Native_get_instance_field.X_invoke_Arity2(other, "Uuid"))
	}
}

func (_ *CljsCoreUUID) CljsCoreObject__() {}
func (self__ *CljsCoreUUID) ToString() string {
	{
		var ______1 = self__
		_ = ______1
		return self__.Uuid.(string)
	}
}

func (self__ *CljsCoreUUID) String() string {
	{
		var this___1 = self__
		_ = this___1
		return this___1.ToString()
	}
}

func (self__ *CljsCoreUUID) Equiv(other interface{}) bool {
	{
		var this___1 = self__
		_ = this___1
		return this___1.X_equiv_Arity2(other)
	}
}

var X__GT_UUID = func(__GT_UUID *AFn) *AFn {
	return Fn(__GT_UUID, func(uuid interface{}) interface{} {
		return (&CljsCoreUUID{uuid})
	})
}(&AFn{})

type CljsCoreExceptionInfo struct {
	Message interface{}
	Data    interface{}
	Cause   interface{}
}

var X__GT_ExceptionInfo = func(__GT_ExceptionInfo *AFn) *AFn {
	return Fn(__GT_ExceptionInfo, func(message interface{}, data interface{}, cause interface{}) interface{} {
		return (&CljsCoreExceptionInfo{message, data, cause})
	})
}(&AFn{})

/**
* Alpha - subject to change.
* Create an instance of ExceptionInfo, an Error type that carries a
* map of additional data.
 */
var Ex_info = func(ex_info *AFn) *AFn {
	return Fn(ex_info, func(msg interface{}, map_ interface{}) interface{} {
		return (&CljsCoreExceptionInfo{msg, map_, nil})
	}, func(msg interface{}, map_ interface{}, cause interface{}) interface{} {
		return (&CljsCoreExceptionInfo{msg, map_, cause})
	})
}(&AFn{})

/**
* Alpha - subject to change.
* Returns exception data (a map) if ex is an ExceptionInfo.
* Otherwise returns nil.
 */
var Ex_data = func(ex_data *AFn) *AFn {
	return Fn(ex_data, func(ex interface{}) interface{} {
		if func() bool { _, instanceof := ex.(*CljsCoreExceptionInfo); return instanceof }() {
			return Native_get_instance_field.X_invoke_Arity2(ex, "Data")
		} else {
			return nil
		}
	})
}(&AFn{})

/**
* Alpha - subject to change.
* Returns the message attached to the given Error / ExceptionInfo object.
* For non-Errors returns nil.
 */
var Ex_message = func(ex_message *AFn) *AFn {
	return Fn(ex_message, func(ex interface{}) interface{} {
		if func() bool { _, instanceof := ex.(*js.Error); return instanceof }() {
			return Native_get_instance_field.X_invoke_Arity2(ex, "Message")
		} else {
			return nil
		}
	})
}(&AFn{})

/**
* Alpha - subject to change.
* Returns exception cause (an Error / ExceptionInfo) if ex is an
* ExceptionInfo.
* Otherwise returns nil.
 */
var Ex_cause = func(ex_cause *AFn) *AFn {
	return Fn(ex_cause, func(ex interface{}) interface{} {
		if func() bool { _, instanceof := ex.(*CljsCoreExceptionInfo); return instanceof }() {
			return Native_get_instance_field.X_invoke_Arity2(ex, "Cause")
		} else {
			return nil
		}
	})
}(&AFn{})

/**
* Returns an JavaScript compatible comparator based upon pred.
 */
var Comparator = func(comparator *AFn) *AFn {
	return Fn(comparator, func(pred interface{}) interface{} {
		return func(G__1233 *AFn) *AFn {
			return Fn(G__1233, func(x interface{}, y interface{}) interface{} {
				if Truth_(pred.(CljsCoreIFn).X_invoke_Arity2(x, y).(interface{})) {
					return float64(-1)
				} else {
					if Truth_(pred.(CljsCoreIFn).X_invoke_Arity2(y, x).(interface{})) {
						return float64(1)
					} else {
						return float64(0)

					}
				}
			})
		}(&AFn{})
	})
}(&AFn{})

var Special_symbol_QMARK_ = func(special_symbol_QMARK_ *AFn) *AFn {
	return Fn(special_symbol_QMARK_, func(x interface{}) bool {
		return Contains_QMARK_.X_invoke_Arity2(&CljsCorePersistentHashSet{nil, &CljsCorePersistentArrayMap{nil, 19, []interface{}{(&CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "defrecord*", Str: "defrecord*", X_hash: float64(-1936366207), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "try", Str: "try", X_hash: float64(-1273693247), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "loop*", Str: "loop*", X_hash: float64(615029416), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "do", Str: "do", X_hash: float64(1686842252), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "letfn*", Str: "letfn*", X_hash: float64(-110097810), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "if", Str: "if", X_hash: float64(1181717262), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "new", Str: "new", X_hash: float64(-444906321), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "ns", Str: "ns", X_hash: float64(2082130287), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "deftype*", Str: "deftype*", X_hash: float64(962659890), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "let*", Str: "let*", X_hash: float64(1920721458), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "js*", Str: "js*", X_hash: float64(-1134233646), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "fn*", Str: "fn*", X_hash: float64(-752876845), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "recur", Str: "recur", X_hash: float64(1202958259), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "set!", Str: "set!", X_hash: float64(250714521), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: ".", Str: ".", X_hash: float64(1975675962), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "quote", Str: "quote", X_hash: float64(1377916282), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "throw", Str: "throw", X_hash: float64(595905694), X_meta: nil}), nil, (&CljsCoreSymbol{Ns: nil, Name: "def", Str: "def", X_hash: float64(597100991), X_meta: nil}), nil}, nil}, nil}, x)
	})
}(&AFn{})
