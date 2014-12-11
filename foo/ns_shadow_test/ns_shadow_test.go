// Compiled by ClojureScript to Go 0.0-2411
// foo.ns-shadow-test

package ns_shadow_test

import (
	"strings"
	"testing"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	"github.com/hraberg/cljs2go/js"
	"github.com/stretchr/testify/assert"
)

func init() {
	Bar = func(bar *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(bar, 0, func() interface{} {
			return float64(1)
		})
	}(&cljs_core.AFn{})

	Quux = func(quux *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(quux, 1, func(foo___1 interface{}) interface{} {
			return (Bar.X_invoke_Arity0().(float64) + foo___1.(float64))
		})
	}(&cljs_core.AFn{})

	Id = func(id *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(id, 1, func(x interface{}) interface{} {
			return x
		})
	}(&cljs_core.AFn{})

	Foo = func(foo___1 *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(foo___1, 0, func() interface{} {
			return Id.X_invoke_Arity1(float64(42))
		})
	}(&cljs_core.AFn{})

	Baz = func(baz___1 *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(baz___1, 1, func() interface{} {
			return baz___1.X_invoke_Arity1(float64(2)).(float64)
		}, func(x interface{}) interface{} {
			return Quux.X_invoke_Arity1(float64(2)).(float64)
		})
	}(&cljs_core.AFn{})

	Test_shadow = func(test_shadow *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test_shadow, 0, func() interface{} {
			if cljs_core.X_EQ_.Arity2IIB(Quux.X_invoke_Arity1(float64(2)).(float64), float64(3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quux 2) 3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(Foo.X_invoke_Arity0(), float64(42)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (foo) 42)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(Baz.X_invoke_Arity0().(float64), float64(3)) {
				return nil
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (baz) 3)").(string)}, ``)}))
			}
		})
	}(&cljs_core.AFn{})

}

var Bar *cljs_core.AFn

var Quux *cljs_core.AFn

var Id *cljs_core.AFn

var Foo *cljs_core.AFn

var Baz *cljs_core.AFn

var Test_shadow *cljs_core.AFn

func Test_runner(t *testing.T) {
	Test_shadow.X_invoke_Arity0()
	assert.True(t, true)
}
