// Compiled by ClojureScript to Go 0.0-2411
// cljs.letfn-test

package letfn_test

import (
	"strings"
	"testing"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	"github.com/hraberg/cljs2go/js"
	"github.com/stretchr/testify/assert"
)

func init() {
	Test_letfn = func(test_letfn *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test_letfn, 0, func() interface{} {
			{
				var ev_QMARK_, od_QMARK_ *cljs_core.AFn
				ev_QMARK_ = func(ev_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(ev_QMARK_, 1, func(x interface{}) interface{} {
						if x.(float64) == float64(0) {
							return true
						} else {
							return od_QMARK_.X_invoke_Arity1((x.(float64) - float64(1)))
						}
					})
				}(&cljs_core.AFn{})
				od_QMARK_ = func(od_QMARK_ *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(od_QMARK_, 1, func(x interface{}) interface{} {
						if x.(float64) == float64(0) {
							return false
						} else {
							return ev_QMARK_.X_invoke_Arity1((x.(float64) - float64(1)))
						}
					})
				}(&cljs_core.AFn{})
				_, _ = ev_QMARK_, od_QMARK_
				if cljs_core.Truth_(ev_QMARK_.X_invoke_Arity1(float64(0))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(ev? 0)").(string)}, ``)}))
				}
				if cljs_core.Truth_(ev_QMARK_.X_invoke_Arity1(float64(10))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(ev? 10)").(string)}, ``)}))
				}
				if cljs_core.Not.Arity1IB(ev_QMARK_.X_invoke_Arity1(float64(1))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (ev? 1))").(string)}, ``)}))
				}
				if cljs_core.Not.Arity1IB(ev_QMARK_.X_invoke_Arity1(float64(11))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (ev? 11))").(string)}, ``)}))
				}
				if cljs_core.Not.Arity1IB(od_QMARK_.X_invoke_Arity1(float64(0))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (od? 0))").(string)}, ``)}))
				}
				if cljs_core.Not.Arity1IB(od_QMARK_.X_invoke_Arity1(float64(10))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (od? 10))").(string)}, ``)}))
				}
				if cljs_core.Truth_(od_QMARK_.X_invoke_Arity1(float64(1))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(od? 1)").(string)}, ``)}))
				}
				if cljs_core.Truth_(od_QMARK_.X_invoke_Arity1(float64(11))) {
					return nil
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(od? 11)").(string)}, ``)}))
				}
			}
		})
	}(&cljs_core.AFn{})

}

var Test_letfn *cljs_core.AFn

func Test_runner(t *testing.T) {
	Test_letfn.X_invoke_Arity0()
	assert.True(t, true)
}
