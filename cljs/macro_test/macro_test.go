// Compiled by ClojureScript to Go 0.0-2356
// cljs.macro-test

package macro_test

import (
	"testing"

	cljs_core "github.com/hraberg/cljs.go/cljs/core"
	"github.com/hraberg/cljs.go/js"
	"github.com/stretchr/testify/assert"
)

func init() {
	Test_macros = func(test_macros *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test_macros, 0, func() interface{} {
			if cljs_core.X_EQ_.Arity2IIB((float64(1) + float64(1)), float64(2)) {
				return nil
			} else {
				panic((&js.Error{("Assert failed: (= (== 1 1) 2)")}))
			}
		})
	}(&cljs_core.AFn{})

}

var Test_macros *cljs_core.AFn

func Test_runner(t *testing.T) {
	Test_macros.X_invoke_Arity0()
	assert.True(t, true)
}
