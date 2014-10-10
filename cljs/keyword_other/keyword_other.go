// Compiled by ClojureScript to Go 0.0-2371
// cljs.keyword-other

package keyword_other

import cljs_core "github.com/hraberg/cljs.go/cljs/core"

func init() {
	Foo = func(foo___1 *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(foo___1, 2, func(a interface{}, b interface{}) interface{} {
			return (a.(float64) + b.(float64))
		})
	}(&cljs_core.AFn{})

}

var Foo *cljs_core.AFn
