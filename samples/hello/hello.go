// Compiled by ClojureScript to Go 0.0-2322
// hello

package hello

import cljs_core "github.com/hraberg/cljs.go/cljs/core"

// @param {...*} var_args
var X_main *cljs_core.AFn

func init() {
	X_main = func(_main *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(_main, func(args__ ...interface{}) interface{} {
			var args = cljs_core.Array_seq.X_invoke_Arity1(args__[0:])
			_ = args
			return cljs_core.Println.X_invoke_ArityVariadic("Hello World")
		})
	}(&cljs_core.AFn{})
}
