// Compiled by ClojureScript to Go 0.0-2371
// cljs.binding-test

package binding_test

import (
	"testing"

	cljs_binding_test_other_ns "github.com/hraberg/cljs2go/cljs/binding_test_other_ns"
	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	"github.com/hraberg/cljs2go/js"
	"github.com/stretchr/testify/assert"
)

func init() {
	Test_binding = func(test_binding *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test_binding, 0, func() interface{} {
			{
				var _STAR_foo_STAR_2_3 = cljs_binding_test_other_ns.X_STAR_foo_STAR_
				_ = _STAR_foo_STAR_2_3
				func() {
					defer func() {
						cljs_binding_test_other_ns.X_STAR_foo_STAR_ = _STAR_foo_STAR_2_3

					}()
					{
						cljs_binding_test_other_ns.X_STAR_foo_STAR_ = float64(2)

						if cljs_core.X_EQ_.Arity2IIB(cljs_binding_test_other_ns.X_STAR_foo_STAR_, float64(2)) {
						} else {
							panic((&js.Error{("Assert failed: (= o/*foo* 2)")}))
						}
					}
				}()
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_binding_test_other_ns.X_STAR_foo_STAR_, float64(1)) {
				return nil
			} else {
				panic((&js.Error{("Assert failed: (= o/*foo* 1)")}))
			}
		})
	}(&cljs_core.AFn{})

	Test_with_redefs = func(test_with_redefs *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test_with_redefs, 0, func() interface{} {
			{
				var bar5_6 = cljs_binding_test_other_ns.Bar
				_ = bar5_6
				func() {
					defer func() {
						cljs_binding_test_other_ns.Bar = bar5_6

					}()
					{
						cljs_binding_test_other_ns.Bar = float64(2)

						if cljs_core.X_EQ_.Arity2IIB(cljs_binding_test_other_ns.Bar, float64(2)) {
						} else {
							panic((&js.Error{("Assert failed: (= o/bar 2)")}))
						}
					}
				}()
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_binding_test_other_ns.Bar, float64(10)) {
				return nil
			} else {
				panic((&js.Error{("Assert failed: (= o/bar 10)")}))
			}
		})
	}(&cljs_core.AFn{})

}

var Test_binding *cljs_core.AFn

var Test_with_redefs *cljs_core.AFn

func Test_runner(t *testing.T) {
	Test_binding.X_invoke_Arity0()
	Test_with_redefs.X_invoke_Arity0()
	assert.True(t, true)
}
