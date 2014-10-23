// Compiled by ClojureScript to Go 0.0-2371
// cljs.ns-test.bar

package bar

import cljs_core "github.com/hraberg/cljs2go/cljs/core"

func init() {
	Quux = func(quux *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(quux, 0, func() interface{} {
			return float64(123)
		})
	}(&cljs_core.AFn{})

}

var Quux *cljs_core.AFn
