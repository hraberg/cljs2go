// Compiled by ClojureScript to Go 0.0-2356
// cljs.top-level-test

package top_level_test

import (
	"testing"

	cljs_core "github.com/hraberg/cljs.go/cljs/core"
	"github.com/hraberg/cljs.go/js"
	"github.com/stretchr/testify/assert"
)

func init() {
	{
		var foo_1___1 = float64(1)
		_ = foo_1___1
		Bar = func(bar *cljs_core.AFn, foo_1___1 float64) *cljs_core.AFn {
			return cljs_core.Fn(bar, 0, func() interface{} {
				return foo_1___1
			})
		}(&cljs_core.AFn{}, foo_1___1)

	}
	{
		var foo_2___1 = float64(2)
		_ = foo_2___1
		Baz = func(baz___1 *cljs_core.AFn, foo_2___1 float64) *cljs_core.AFn {
			return cljs_core.Fn(baz___1, 0, func() interface{} {
				return foo_2___1
			})
		}(&cljs_core.AFn{}, foo_2___1)

	}
	Test = func(test *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test, 0, func() interface{} {
			if cljs_core.X_EQ_.Arity2IIB(Bar.X_invoke_Arity0().(float64), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (bar) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(Baz.X_invoke_Arity0().(float64), float64(2)) {
				return nil
			} else {
				panic((&js.Error{("Assert failed: (= (baz) 2)")}))
			}
		})
	}(&cljs_core.AFn{})

}

var Bar *cljs_core.AFn

var Baz *cljs_core.AFn

var Test *cljs_core.AFn

func Test_runner(t *testing.T) {
	Test.X_invoke_Arity0()
	assert.True(t, true)
}
