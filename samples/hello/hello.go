// Compiled by ClojureScript to Go 0.0-2356
// hello

package hello

import cljs_core "github.com/hraberg/cljs.go/cljs/core"

func init() {
	X_main = func(_main *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(_main, 0, func(args__ ...interface{}) interface{} {
			var args = cljs_core.Seq.Arity1IQ(args__[0])
			_ = args
			return cljs_core.Println.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"Hello World"}))
		})
	}(&cljs_core.AFn{})

}

// @param {...*} var_args
var X_main *cljs_core.AFn
