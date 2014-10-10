// Compiled by ClojureScript to Go 0.0-2371
// cljs.ns-test.foo

package foo

import (
	cljs_core "github.com/hraberg/cljs.go/cljs/core"
	"github.com/hraberg/cljs.go/js"
)

func init() {
	Baz = func(baz *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(baz, 0, func() interface{} {
			return float64(123)
		})
	}(&cljs_core.AFn{})

	Kw = (&cljs_core.CljsCoreKeyword{Ns: "cljs.ns-test.foo", Name: "foo", Fqn: "cljs.ns-test.foo/foo", X_hash: float64(-1760770742)})

	Qkw = (&cljs_core.CljsCoreKeyword{Ns: "cljs.ns-test.foo", Name: "foo", Fqn: "cljs.ns-test.foo/foo", X_hash: float64(-1760770742)})

	if cljs_core.X_EQ_.Arity2IIB((`` + cljs_core.Str.X_invoke_Arity1(Kw).(string)), ":cljs.ns-test.foo/foo") {
	} else {
		panic((&js.Error{("Assert failed: (= (str kw) \":cljs.ns-test.foo/foo\")")}))
	}
	if cljs_core.X_EQ_.Arity2IIB((`` + cljs_core.Str.X_invoke_Arity1(Qkw).(string)), ":cljs.ns-test.foo/foo") {
	} else {
		panic((&js.Error{("Assert failed: (= (str qkw) \":cljs.ns-test.foo/foo\")")}))
	}
}

var Baz *cljs_core.AFn

var Kw *cljs_core.CljsCoreKeyword

var Qkw *cljs_core.CljsCoreKeyword
