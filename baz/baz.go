// Compiled by ClojureScript to Go 0.0-2356
// baz

package baz

import cljs_core "github.com/hraberg/cljs.go/cljs/core"

func init() {
	F = func(f *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(f, 1, func(x interface{}) interface{} {
			return x
		})
	}(&cljs_core.AFn{})

}

var F *cljs_core.AFn
