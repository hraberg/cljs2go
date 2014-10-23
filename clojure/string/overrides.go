// Compiled by ClojureScript to Go 0.0-2371
// clojure.string

// Go overrides.
package string

import (
	"reflect"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	goog_string "github.com/hraberg/cljs2go/goog/string"
	"github.com/hraberg/cljs2go/js"
)

func init() {
	Replace = func(replace *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(replace, 3, func(s interface{}, match interface{}, replacement interface{}) interface{} {
			if cljs_core.Value_(match).Kind() == reflect.String {
				return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "Replace", []interface{}{(&js.RegExp{func() interface{} {
					var G__31 = match
					_ = G__31
					return cljs_core.Native_invoke_func.X_invoke_Arity2(goog_string.RegExpEscape, []interface{}{G__31})
				}(), "g"}), replacement})
			} else {
				if cljs_core.Value_(match).Type().AssignableTo(reflect.TypeOf((**js.RegExp)(nil)).Elem()) {
					return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "Replace", []interface{}{(&js.RegExp{cljs_core.Native_get_instance_field.X_invoke_Arity2(match, "Pattern"), (`` + cljs_core.Str.X_invoke_Arity1(cljs_core.Native_get_instance_field.X_invoke_Arity2(match, "Flags")).(string) + "g")}), func() interface{} {
						if cljs_core.Fn_QMARK_.Arity1IB(replacement) {
							return func(x interface{}) interface{} { return replacement.(cljs_core.CljsCoreIFn).X_invoke_Arity1(x) }
						} else {
							return replacement
						}
					}()})
				} else {
					panic(("Invalid match arg: " + cljs_core.Str.X_invoke_Arity1(match).(string)))

				}
			}
		})
	}(&cljs_core.AFn{})

	Replace_first = func(replace_first *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(replace_first, 3, func(s interface{}, match interface{}, replacement interface{}) interface{} {
			return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(s, "Replace", []interface{}{match, func() interface{} {
				if cljs_core.Fn_QMARK_.Arity1IB(replacement) {
					return func(x interface{}) interface{} { return replacement.(cljs_core.CljsCoreIFn).X_invoke_Arity1(x) }
				} else {
					return replacement
				}
			}()})
		})
	}(&cljs_core.AFn{})

}

// Replaces all instance of match with replacement in s.
// match/replacement can be:
//
// string / string
// pattern / (string or function of match).
var Replace *cljs_core.AFn

// Replaces the first instance of match with replacement in s.
// match/replacement can be:
//
// string / string
// pattern / (string or function of match).
var Replace_first *cljs_core.AFn
