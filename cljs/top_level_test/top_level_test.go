// Compiled by ClojureScript to Go 0.0-2411
// cljs.top-level-test

package top_level_test

import (
	"strings"
	"testing"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	"github.com/hraberg/cljs2go/js"
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (bar) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(Baz.X_invoke_Arity0().(float64), float64(2)) {
				return nil
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (baz) 2)").(string)}, ``)}))
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
