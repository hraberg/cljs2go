// Compiled by ClojureScript to Go 0.0-2371
// clojure.data

// Go overrides.
package data

import (
	"reflect"

	cljs_core "github.com/hraberg/cljs.go/cljs/core"
	"github.com/hraberg/cljs.go/js"
)

func init() {
	Ep = func(ep *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(ep, 1, func(x interface{}) interface{} {
			if cljs_core.Nil_(x) {
				return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "atom", Fqn: "atom", X_hash: float64(-397043653)})
			} else {
				if cljs_core.Value_(x).Kind() == reflect.String {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "atom", Fqn: "atom", X_hash: float64(-397043653)})
				} else {
					if cljs_core.Value_(x).Kind() == reflect.Float64 {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "atom", Fqn: "atom", X_hash: float64(-397043653)})
					} else {
						if cljs_core.Value_(x).Kind() == reflect.Slice {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sequential", Fqn: "sequential", X_hash: float64(-1082983960)})
						} else {
							if cljs_core.Fn_QMARK_.Arity1IB(x) {
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "atom", Fqn: "atom", X_hash: float64(-397043653)})
							} else {
								if cljs_core.Value_(x).Kind() == reflect.Bool {
									return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "atom", Fqn: "atom", X_hash: float64(-397043653)})
								} else {
									if cljs_core.DecoratedValue_(x).Type().Implements(reflect.TypeOf((*cljs_core.CljsCoreIMap)(nil)).Elem()) {
										return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "map", Fqn: "map", X_hash: float64(1371690461)})
									} else {
										if cljs_core.DecoratedValue_(x).Type().Implements(reflect.TypeOf((*cljs_core.CljsCoreISet)(nil)).Elem()) {
											return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "set", Fqn: "set", X_hash: float64(304602554)})
										} else {
											if cljs_core.DecoratedValue_(x).Type().Implements(reflect.TypeOf((*cljs_core.CljsCoreISequential)(nil)).Elem()) {
												return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sequential", Fqn: "sequential", X_hash: float64(-1082983960)})
											} else {
												return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "atom", Fqn: "atom", X_hash: float64(-397043653)})

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
	}(&cljs_core.AFn{})

	Ds = func(ds *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(ds, 2, func(a interface{}, b interface{}) interface{} {
			return func() interface{} {
				var G__17 = func() interface{} {
					if cljs_core.Value_(Ep.X_invoke_Arity1(a).(*cljs_core.CljsCoreKeyword)).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
						return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(Ep.X_invoke_Arity1(a).(*cljs_core.CljsCoreKeyword)), "Fqn")
					} else {
						return nil
					}
				}()
				_ = G__17
				switch G__17 {
				case "map":
					return Diff_associative

				case "sequential":
					return Diff_sequential

				case "set":
					return Diff_set

				case "atom":
					return Atom_diff

				default:
					panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(Ep.X_invoke_Arity1(a).(*cljs_core.CljsCoreKeyword)).(string))}))

				}
			}().(cljs_core.CljsCoreIFn).X_invoke_Arity2(a, b)
		})
	}(&cljs_core.AFn{})

	Diff = func(diff *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(diff, 2, func(a interface{}, b interface{}) interface{} {
			if cljs_core.X_EQ_.Arity2IIB(a, b) {
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, nil, a}, nil})
			} else {
				if cljs_core.X_EQ_.Arity2IIB(Ep.X_invoke_Arity1(a).(*cljs_core.CljsCoreKeyword), Ep.X_invoke_Arity1(b).(*cljs_core.CljsCoreKeyword)) {
					return Ds.X_invoke_Arity2(a, b)
				} else {
					return Atom_diff.X_invoke_Arity2(a, b).(cljs_core.CljsCoreIVector)
				}
			}
		})
	}(&cljs_core.AFn{})

}

var Ep *cljs_core.AFn

var Ds *cljs_core.AFn

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
