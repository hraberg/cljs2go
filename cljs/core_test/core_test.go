// Compiled by ClojureScript to Go 0.0-2356
// cljs.core-test

package core_test

import (
	"reflect"
	"testing"

	cljs_core "github.com/hraberg/cljs.go/cljs/core"
	clojure_set "github.com/hraberg/cljs.go/clojure/set"
	goog_string "github.com/hraberg/cljs.go/goog/string"
	"github.com/hraberg/cljs.go/js"
	"github.com/hraberg/cljs.go/js/Math"
	"github.com/stretchr/testify/assert"
)

func init() {
	Test_stuff = func(test_stuff *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test_stuff, 0, func() interface{} {
			if cljs_core.X_EQ_.Arity1IB(float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 1)")}))
			}
			if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 1 1)")}))
			}
			if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(1)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 1 1 1)")}))
			}
			if !(cljs_core.X_EQ_.Arity2IIB(float64(1), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (not (= 1 2))")}))
			}
			if !(cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})))) {
			} else {
				panic((&js.Error{("Assert failed: (not (= 1 2 1))")}))
			}
			if !(cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)})))) {
			} else {
				panic((&js.Error{("Assert failed: (not (= 1 1 2))")}))
			}
			if !(cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(1)})))) {
			} else {
				panic((&js.Error{("Assert failed: (not (= 1 1 2 1))")}))
			}
			if !(cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})))) {
			} else {
				panic((&js.Error{("Assert failed: (not (= 1 1 1 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (+) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.CljsCorePersistentVector_EMPTY), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply + []) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (+ 1) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply + [1]) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) + float64(1)), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (+ 1 1) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply + [1 1]) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(((float64(1) + float64(2)) + float64(3)), float64(6)) {
			} else {
				panic((&js.Error{("Assert failed: (= (+ 1 2 3) 6)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})), float64(6)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply + [1 2 3]) 6)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((-float64(1)), float64(-1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (- 1) -1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X___, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(-1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply - [1]) -1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) - float64(1)), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (- 1 1) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X___, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply - [1 1]) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(((float64(3) - float64(2)) - float64(1)), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (- 3 2 1) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X___, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil})), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply - [3 2 1]) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (*) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_STAR_, cljs_core.CljsCorePersistentVector_EMPTY), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply * []) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (* 2) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_STAR_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil})), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply * [2]) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(2) * float64(3)), float64(6)) {
			} else {
				panic((&js.Error{("Assert failed: (= (* 2 3) 6)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_STAR_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})), float64(6)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply * [2 3]) 6)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) / float64(2)), 0.5) {
			} else {
				panic((&js.Error{("Assert failed: (= (/ 2) 0.5)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_SLASH_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil})), 0.5) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply / [2]) 0.5)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(6) / float64(2)), float64(3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (/ 6 2) 3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_SLASH_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(6), float64(2)}, nil})), float64(3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply / [6 2]) 3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(((float64(6) / float64(3)) / float64(2)), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (/ 6 3 2) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_SLASH_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(6), float64(3), float64(2)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply / [6 3 2]) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(true, true) {
			} else {
				panic((&js.Error{("Assert failed: (= (< 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply < [1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) < float64(2)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (< 1 2) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply < [1 2]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) < float64(1)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (< 1 1) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply < [1 1]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(2) < float64(1)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (< 2 1) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply < [2 1]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) < float64(2)) && (float64(2) < float64(3)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (< 1 2 3) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply < [1 2 3]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) < float64(1)) && (float64(1) < float64(3)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (< 1 1 3) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(3)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply < [1 1 3]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) < float64(1)) && (float64(1) < float64(1)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (< 3 1 1) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply < [3 1 1]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(true, true) {
			} else {
				panic((&js.Error{("Assert failed: (= (<= 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply <= [1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) <= float64(1)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (<= 1 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply <= [1 1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) <= float64(2)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (<= 1 2) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply <= [1 2]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(2) <= float64(1)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (<= 2 1) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply <= [2 1]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) <= float64(2)) && (float64(2) <= float64(3)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (<= 1 2 3) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply <= [1 2 3]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) <= float64(1)) && (float64(1) <= float64(3)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (<= 1 1 3) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(3)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply <= [1 1 3]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) <= float64(1)) && (float64(1) <= float64(1)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (<= 3 1 1) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply <= [3 1 1]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(true, true) {
			} else {
				panic((&js.Error{("Assert failed: (= (> 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply > [1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(2) > float64(1)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (> 2 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply > [2 1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) > float64(1)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (> 1 1) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply > [1 1]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) > float64(2)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (> 1 2) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply > [1 2]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) > float64(2)) && (float64(2) > float64(1)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (> 3 2 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply > [3 2 1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) > float64(1)) && (float64(1) > float64(1)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (> 3 1 1) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply > [3 1 1]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) > float64(1)) && (float64(1) > float64(3)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (> 1 1 3) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(3)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply > [1 1 3]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(true, true) {
			} else {
				panic((&js.Error{("Assert failed: (= (>= 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply >= [1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(2) >= float64(1)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (>= 2 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply >= [2 1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) >= float64(1)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (>= 1 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply >= [1 1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) >= float64(2)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (>= 1 2) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply >= [1 2]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) >= float64(2)) && (float64(2) >= float64(1)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (>= 3 2 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply >= [3 2 1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) >= float64(1)) && (float64(1) >= float64(1)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (>= 3 1 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply >= [3 1 1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) >= float64(1)) && (float64(1) >= float64(2)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (>= 3 1 2) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(2)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply >= [3 1 2]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) >= float64(1)) && (float64(1) >= float64(3)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (>= 1 1 3) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(3)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply >= [1 1 3]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) - float64(1)), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (dec 1) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Dec, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply dec [1]) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(0) + float64(1)), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (inc 0) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Inc, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply inc [0]) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(0) == float64(0)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (zero? 0) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Zero_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply zero? [0]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) == float64(0)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (zero? 1) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Zero_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply zero? [1]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(-11) == float64(0)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (zero? -11) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Zero_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-11)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply zero? [-11]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(0) > float64(0)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (pos? 0) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Pos_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply pos? [0]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) > float64(0)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (pos? 1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Pos_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply pos? [1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(-1) > float64(0)), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (pos? -1) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Pos_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-1)}, nil})), false) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply pos? [-1]) false)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(-1) < float64(0)), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (neg? -1) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Neg_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-1)}, nil})), true) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply neg? [-1]) true)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (max 1) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply max [1]) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(float64(1), float64(2)), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (max 1 2) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply max [1 2]) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(float64(2), float64(1)), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (max 2 1) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply max [2 1]) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(float64(1), float64(2)), float64(3)), float64(3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (max 1 2 3) 3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})), float64(3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply max [1 2 3]) 3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(float64(1), float64(3)), float64(2)), float64(3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (max 1 3 2) 3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(2)}, nil})), float64(3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply max [1 3 2]) 3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (min 1) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Min, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply min [1]) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(float64(1), float64(2)), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (min 1 2) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Min, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply min [1 2]) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(float64(2), float64(1)), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (min 2 1) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Min, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply min [2 1]) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(float64(1), float64(2)), float64(3)), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (min 1 2 3) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Min, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply min [1 2 3]) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(float64(2), float64(1)), float64(3)), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (min 2 1 3) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Min, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(3)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply min [3 1 3]) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(4), float64(2)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 4 2) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Mod, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(2)}, nil})), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply mod [4 2]) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(3), float64(2)).(float64), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 3 2) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Mod, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply mod [3 2]) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(-2), float64(5)).(float64), float64(3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod -2 5) 3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(3), float64(2), float64(1), float64(0)}, nil}), func() interface{} {
				var i = float64(0)
				var j interface{} = cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)
				_, _ = i, j
				for {
					if i < float64(5) {
						i, j = (i + float64(1)), cljs_core.Conj.X_invoke_Arity2(j, func(G__4967 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__4967, 0, func() interface{} {
								return i
							})
						}(&cljs_core.AFn{}, i, j))
						continue
					} else {
						return cljs_core.Map_.X_invoke_Arity2(func(G__4968 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__4968, 1, func(p1__4076_SHARP_ interface{}) interface{} {
								{
									return p1__4076_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
								}
							})
						}(&cljs_core.AFn{}, i, j), j).(*cljs_core.CljsCoreLazySeq)
					}
				}
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [4 3 2 1 0] (loop [i 0 j ()] (if (< i 5) (recur (inc i) (conj j (fn [] i))) (map (fn* [p1__4076#] (p1__4076#)) j))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})}, nil}), cljs_core.Map_.X_invoke_Arity2(func(G__4969 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4969, 1, func(p1__4077_SHARP_ interface{}) interface{} {
					{
						return p1__4077_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
					}
				})
			}(&cljs_core.AFn{}), func() *cljs_core.CljsCoreLazySeq {
				var iter__929__auto__ = func(iter__4541 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(iter__4541, 1, func(s__4542 interface{}) interface{} {
						return (&cljs_core.CljsCoreLazySeq{nil, func(G__4970 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__4970, 0, func() interface{} {
								{
									var s__4542___1 interface{} = s__4542
									_ = s__4542___1
									for {
										{
											var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4542___1)
											_ = temp__4222__auto__
											if cljs_core.Truth_(temp__4222__auto__) {
												{
													var xs__4752__auto__ = temp__4222__auto__
													_ = xs__4752__auto__
													{
														var i = cljs_core.First.X_invoke_Arity1(xs__4752__auto__)
														_ = i
														{
															var iterys__925__auto__ = func(iter__4543 *cljs_core.AFn, s__4542___1 interface{}, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																return cljs_core.Fn(iter__4543, 1, func(s__4544 interface{}) interface{} {
																	return (&cljs_core.CljsCoreLazySeq{nil, func(G__4971 *cljs_core.AFn, s__4542___1 interface{}, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																		return cljs_core.Fn(G__4971, 0, func() interface{} {
																			{
																				var s__4544___1 interface{} = s__4544
																				_ = s__4544___1
																				for {
																					{
																						var temp__4222__auto_____1 = cljs_core.Seq.Arity1IQ(s__4544___1)
																						_ = temp__4222__auto_____1
																						if cljs_core.Truth_(temp__4222__auto_____1) {
																							{
																								var s__4544___2 = temp__4222__auto_____1
																								_ = s__4544___2
																								if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4544___2) {
																									{
																										var c__927__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4544___2)
																										var size__928__auto__ = cljs_core.Count.X_invoke_Arity1(c__927__auto__).(float64)
																										var b__4546 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__928__auto__).(*cljs_core.CljsCoreChunkBuffer)
																										_, _, _ = c__927__auto__, size__928__auto__, b__4546
																										if func() bool {
																											var i__4545 = float64(0)
																											_ = i__4545
																											for {
																												if i__4545 < size__928__auto__ {
																													{
																														var j = c__927__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4545)
																														_ = j
																														cljs_core.Chunk_append.X_invoke_Arity2(b__4546, func(G__4972 *cljs_core.AFn, i__4545 float64, s__4542___1 interface{}, j interface{}, c__927__auto__ interface{}, size__928__auto__ float64, b__4546 *cljs_core.CljsCoreChunkBuffer, s__4544___2 interface{}, temp__4222__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																															return cljs_core.Fn(G__4972, 0, func() interface{} {
																																return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																															})
																														}(&cljs_core.AFn{}, i__4545, s__4542___1, j, c__927__auto__, size__928__auto__, b__4546, s__4544___2, temp__4222__auto_____1, i, xs__4752__auto__, temp__4222__auto__))
																														i__4545 = (i__4545 + float64(1))
																														continue
																													}
																												} else {
																													return true
																												}
																											}
																										}() {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4546), iter__4543.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4544___2)).(*cljs_core.CljsCoreLazySeq))
																										} else {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4546), nil)
																										}
																									}
																								} else {
																									{
																										var j = cljs_core.First.X_invoke_Arity1(s__4544___2)
																										_ = j
																										return cljs_core.Cons.X_invoke_Arity2(func(G__4973 *cljs_core.AFn, s__4542___1 interface{}, j interface{}, s__4544___2 interface{}, temp__4222__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																											return cljs_core.Fn(G__4973, 0, func() interface{} {
																												return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																											})
																										}(&cljs_core.AFn{}, s__4542___1, j, s__4544___2, temp__4222__auto_____1, i, xs__4752__auto__, temp__4222__auto__), iter__4543.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4544___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
																									}
																								}
																							}
																						} else {
																							return nil
																						}
																					}
																				}
																			}
																		})
																	}(&cljs_core.AFn{}, s__4542___1, i, xs__4752__auto__, temp__4222__auto__), nil, nil})
																})
															}(&cljs_core.AFn{}, s__4542___1, i, xs__4752__auto__, temp__4222__auto__)
															var fs__926__auto__ = cljs_core.Seq.Arity1IQ(iterys__925__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq))
															_, _ = iterys__925__auto__, fs__926__auto__
															if cljs_core.Truth_(fs__926__auto__) {
																return cljs_core.Concat.X_invoke_Arity2(fs__926__auto__, iter__4541.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4542___1)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
															} else {
																s__4542___1 = cljs_core.Rest.Arity1IQ(s__4542___1)
																continue
															}
														}
													}
												}
											} else {
												return nil
											}
										}
									}
								}
							})
						}(&cljs_core.AFn{}), nil, nil})
					})
				}(&cljs_core.AFn{})
				_ = iter__929__auto__
				return iter__929__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})).(*cljs_core.CljsCoreLazySeq)
			}()).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [[1 1] [1 2] [1 3] [2 1] [2 2] [2 3]] (map (fn* [p1__4077#] (p1__4077#)) (for [i [1 2] j [1 2 3]] (fn [] [i j]))))")}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (integer? 0)")}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(float64(42)) {
			} else {
				panic((&js.Error{("Assert failed: (integer? 42)")}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(float64(-42)) {
			} else {
				panic((&js.Error{("Assert failed: (integer? -42)")}))
			}
			if !(cljs_core.Integer_QMARK_.Arity1IB("")) {
			} else {
				panic((&js.Error{("Assert failed: (not (integer? \"\"))")}))
			}
			if !(cljs_core.Integer_QMARK_.Arity1IB(1.0E308)) {
			} else {
				panic((&js.Error{("Assert failed: (not (integer? 1.0E308))")}))
			}
			if !(cljs_core.Integer_QMARK_.Arity1IB(js.Infinity)) {
			} else {
				panic((&js.Error{("Assert failed: (not (integer? js/Infinity))")}))
			}
			if !(cljs_core.Integer_QMARK_.Arity1IB((-js.Infinity))) {
			} else {
				panic((&js.Error{("Assert failed: (not (integer? (- js/Infinity)))")}))
			}
			if !(cljs_core.Integer_QMARK_.Arity1IB(js.NaN)) {
			} else {
				panic((&js.Error{("Assert failed: (not (integer? js/NaN))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(42), float64(cljs_core.Int32_(42.5))) {
			} else {
				panic((&js.Error{("Assert failed: (= 42 (int 42.5))")}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(float64(cljs_core.Int32_(42.5))) {
			} else {
				panic((&js.Error{("Assert failed: (integer? (int 42.5))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Long.X_invoke_Arity1(42.5).(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= 42 (long 42.5))")}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(cljs_core.Long.X_invoke_Arity1(42.5).(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (integer? (long 42.5))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), float64(cljs_core.Int32_(-1.5))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (int -1.5))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-9), cljs_core.Long.X_invoke_Arity1(-9.8).(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= -9 (long -9.8))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (:b {:a 1, :b 2}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}).X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 ((quote b) (quote {:a 1, b 2})))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}).X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 ({:a 1, :b 2} :b))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), float64(1), float64(2), float64(2)}, nil}).X_invoke_Arity1(float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 ({1 1, 2 2} 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(1)}, nil}), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (:a {:b 1} 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (:a {} 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(1)}, nil}).X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 ({:b 1} :a 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.CljsCorePersistentArrayMap_EMPTY.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 ({} :a 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (:a {}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1("")) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (:a \"\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity2("", float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (:a \"\" 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}).X_invoke_Arity1(float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (#{1 3 2} 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Apply.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), float64(2)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (apply :a (quote [{:a 1, a 2}])))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Apply.X_invoke_Arity2((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (apply (quote a) (quote [{a 1, :b 2}])))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Apply.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (apply {:a 1} [:a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Apply.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (apply {:a 1} [:b 2]))")}))
			}
			if cljs_core.Nil_(cljs_core.Namespace.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "/", Str: "/", X_hash: float64(-1371932971), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (namespace (quote /)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("/", cljs_core.Name.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "/", Str: "/", X_hash: float64(-1371932971), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"/\" (name (quote /)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("keyword", cljs_core.Name.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "keyword", Fqn: "keyword", X_hash: float64(811389747)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"keyword\" (name :keyword))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(":hello", (`` + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hello", Fqn: "hello", X_hash: float64(-245025397)})).(string))) {
			} else {
				panic((&js.Error{("Assert failed: (= \":hello\" (str :hello))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("hello", (`` + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "hello", Str: "hello", X_hash: float64(1395506130), X_meta: nil})).(string))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"hello\" (str (quote hello)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("hello:world", ("hello" + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "world", Fqn: "world", X_hash: float64(-418292623)})).(string))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"hello:world\" (str \"hello\" :world))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(":helloworld", (`` + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hello", Fqn: "hello", X_hash: float64(-245025397)})).(string) + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "world", Str: "world", X_hash: float64(1222238904), X_meta: nil})).(string))) {
			} else {
				panic((&js.Error{("Assert failed: (= \":helloworld\" (str :hello (quote world)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), cljs_core.Symbol.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote a) (symbol (quote a)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Keyword.X_invoke_Arity1("a")) {
			} else {
				panic((&js.Error{("Assert failed: (= :a (keyword \"a\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Keyword.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= :a (keyword (quote a)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), cljs_core.Keyword.X_invoke_Arity2((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil})).(*cljs_core.CljsCoreKeyword)) {
			} else {
				panic((&js.Error{("Assert failed: (= :a/b (keyword (quote a) (quote b)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Keyword.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= :a (keyword :a))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), cljs_core.Get.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), float64(4), float64(5)}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:a :b} (get {[1 2 3] {:a :b}, 4 5} [1 2 3]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil}), float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= :a (nth [:a :b :c :d] 0))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil}), 0.1)) {
			} else {
				panic((&js.Error{("Assert failed: (= :a (nth [:a :b :c :d] 0.1))")}))
			}
			if !(cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), nil}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (not (= {:a :b, :c nil} {:a :b, :d nil}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3)), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (list 3 2 1) [3 2 1])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Seq.Arity1IQ([]interface{}{float64(3), float64(2), float64(1)})) {
			} else {
				panic((&js.Error{("Assert failed: (= [3 2 1] (seq (array 3 2 1)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(9), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Next.Arity1IQ(cljs_core.Seq.Arity1IQ([]interface{}{float64(1), float64(2), float64(3), float64(4)})))) {
			} else {
				panic((&js.Error{("Assert failed: (= 9 (reduce + (next (seq (array 1 2 3 4)))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Rest.Arity1IQ(nil)) {
			} else {
				panic((&js.Error{("Assert failed: (= () (rest nil))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Seq.Arity1IQ([]interface{}{})) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (seq (array)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Seq.Arity1IQ("")) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (seq \"\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Seq.Arity1IQ(cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (seq []))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Seq.Arity1IQ(cljs_core.CljsCorePersistentArrayMap_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (seq {}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Rest.Arity1IQ(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY))) {
			} else {
				panic((&js.Error{("Assert failed: (= () (rest ()))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Rest.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= () (rest [1]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Rest.Arity1IQ([]interface{}{float64(1)})) {
			} else {
				panic((&js.Error{("Assert failed: (= () (rest (array 1)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"x", "y"}, nil}), cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"x", "y"}, nil})))) {
			} else {
				panic((&js.Error{("Assert failed: (= {\"x\" \"y\"} (meta []))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), cljs_core.Dissoc.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:a :b} (dissoc {:a :b, :c :d} :c))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("\"asdf\" \"asdf\"", cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"asdf", "asdf"})).(string)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\\\"asdf\\\" \\\"asdf\\\"\" (pr-str \"asdf\" \"asdf\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("[1 true {:a 2, :b #\"x\\\"y\"} #js [3 4]]", cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), true, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&js.RegExp{Pattern: `x\"y`, Flags: ``})}, nil}), []interface{}{float64(3), float64(4)}}, nil})})).(string)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"[1 true {:a 2, :b #\\\"x\\\\\\\"y\\\"} #js [3 4]]\" (pr-str [1 true {:a 2, :b #\"x\\\"y\"} (array 3 4)]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("\"asdf\"\n", cljs_core.Prn_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"asdf"})).(string)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\\\"asdf\\\"\\n\" (prn-str \"asdf\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("[1 true {:a 2, :b 42} #js [3 4]]\n", cljs_core.Prn_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), true, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(42)}, nil}), []interface{}{float64(3), float64(4)}}, nil})})).(string)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"[1 true {:a 2, :b 42} #js [3 4]]\\n\" (prn-str [1 true {:a 2, :b 42} (array 3 4)]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("asdf", cljs_core.Print_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"asdf"})).(string)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"asdf\" (print-str \"asdf\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("asdf\n", cljs_core.Println_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"asdf"})).(string)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"asdf\\n\" (println-str \"asdf\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("", cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{})).(string)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\" (pr-str))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("\n", cljs_core.Prn_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{})).(string)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\\n\" (prn-str))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("12", func() string {
				var sb__1130__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1130__auto__
				{
					var _STAR_print_fn_STAR_4552_4974 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_4552_4974
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_4552_4974

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__4975 *cljs_core.AFn, _STAR_print_fn_STAR_4552_4974 interface{}, sb__1130__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__4975, 1, func(x__1131__auto__ interface{}) interface{} {
									return sb__1130__auto__.Append(x__1131__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_4552_4974, sb__1130__auto__)

							cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)}))
							cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)}))
						}
					}()
				}
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1130__auto__).(string))
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= \"12\" (with-out-str (print 1) (print 2)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("12", func() string {
				var sb__1130__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1130__auto__
				{
					var _STAR_print_fn_STAR_4553_4976 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_4553_4976
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_4553_4976

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__4977 *cljs_core.AFn, _STAR_print_fn_STAR_4553_4976 interface{}, sb__1130__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__4977, 1, func(x__1131__auto__ interface{}) interface{} {
									return sb__1130__auto__.Append(x__1131__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_4553_4976, sb__1130__auto__)

							cljs_core.X_STAR_print_fn_STAR_.X_invoke_Arity1(float64(1))
							cljs_core.X_STAR_print_fn_STAR_.X_invoke_Arity1(float64(2))
						}
					}()
				}
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1130__auto__).(string))
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= \"12\" (with-out-str (*print-fn* 1) (*print-fn* 2)))")}))
			}
			if !(cljs_core.X_EQ_.Arity2IIB("one", "two")) {
			} else {
				panic((&js.Error{("Assert failed: (not (= \"one\" \"two\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Count.X_invoke_Arity1("abc").(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= 3 (count \"abc\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Count.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(4)}).(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= 4 (count (array 1 2 3 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("c", cljs_core.Nth.X_invoke_Arity2("abc", float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"c\" (nth \"abc\" 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("quux", cljs_core.Nth.X_invoke_Arity3("abc", float64(3), "quux")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"quux\" (nth \"abc\" 3 \"quux\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Nth.X_invoke_Arity2([]interface{}{float64(1), float64(2), float64(3), float64(4)}, float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (nth (array 1 2 3 4) 0))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("val", cljs_core.Nth.X_invoke_Arity3([]interface{}{float64(1), float64(2), float64(3), float64(4)}, float64(4), "val")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"val\" (nth (array 1 2 3 4) 4 \"val\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("b", cljs_core.Get.X_invoke_Arity2("abc", float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"b\" (get \"abc\" 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("harriet", cljs_core.Get.X_invoke_Arity3("abcd", float64(4), "harriet")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"harriet\" (get \"abcd\" 4 \"harriet\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity2([]interface{}{float64(1), float64(2), float64(3), float64(4)}, float64(3))) {
			} else {
				panic((&js.Error{("Assert failed: (= 4 (get (array 1 2 3 4) 3))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("zot", cljs_core.Get.X_invoke_Arity3([]interface{}{float64(1), float64(2), float64(3), float64(4)}, float64(4), "zot")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"zot\" (get (array 1 2 3 4) 4 \"zot\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(10), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, []interface{}{float64(1), float64(2), float64(3), float64(4)})) {
			} else {
				panic((&js.Error{("Assert failed: (= 10 (reduce + (array 1 2 3 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(20), cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), []interface{}{float64(1), float64(2), float64(3), float64(4)})) {
			} else {
				panic((&js.Error{("Assert failed: (= 20 (reduce + 10 (array 1 2 3 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("cabd", func() interface{} {
				var jumble = func(G__4978 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4978, 2, func(a interface{}, b interface{}) interface{} {
						return (`` + cljs_core.Str.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Reverse.X_invoke_Arity1((``+cljs_core.Str.X_invoke_Arity1(a).(string))))).(string) + cljs_core.Str.X_invoke_Arity1(b).(string))
					})
				}(&cljs_core.AFn{})
				_ = jumble
				return cljs_core.Reduce.X_invoke_Arity2(jumble, "abcd")
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= \"cabd\" (let [jumble (fn [a b] (str (apply str (reverse (str a))) b))] (reduce jumble \"abcd\")))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("cafrogbd", func() interface{} {
				var jumble = func(G__4979 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4979, 2, func(a interface{}, b interface{}) interface{} {
						return (`` + cljs_core.Str.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Reverse.X_invoke_Arity1((``+cljs_core.Str.X_invoke_Arity1(a).(string))))).(string) + cljs_core.Str.X_invoke_Arity1(b).(string))
					})
				}(&cljs_core.AFn{})
				_ = jumble
				return cljs_core.Reduce.X_invoke_Arity3(jumble, "frog", "abcd")
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= \"cafrogbd\" (let [jumble (fn [a b] (str (apply str (reverse (str a))) b))] (reduce jumble \"frog\" \"abcd\")))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0), float64(1), float64(0), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) & cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(0)) & cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(1)) & cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(42)) & cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(41)) & cljs_core.Int32_(float64(1))))}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= [0 0 1 0 1] [(bit-and 1 0) (bit-and 0 0) (bit-and 1 1) (bit-and 42 1) (bit-and 41 1)])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(0), float64(1), float64(43), float64(41)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) | cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(0)) | cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(1)) | cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(42)) | cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(41)) | cljs_core.Int32_(float64(1))))}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 0 1 43 41] [(bit-or 1 0) (bit-or 0 0) (bit-or 1 1) (bit-or 42 1) (bit-or 41 1)])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(0), float64(0), float64(42), float64(40)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) &^ cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(0)) &^ cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(1)) &^ cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(42)) &^ cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(41)) &^ cljs_core.Int32_(float64(1))))}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 0 0 42 40] [(bit-and-not 1 0) (bit-and-not 0 0) (bit-and-not 1 1) (bit-and-not 42 1) (bit-and-not 41 1)])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2), float64(968), float64(16649), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) &^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(0))))), float64((cljs_core.Int32_(float64(2)) &^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(0))))), float64((cljs_core.Int32_(float64(1000)) &^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(5))))), float64((cljs_core.Int32_(float64(16713)) &^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(6))))), float64((cljs_core.Int32_(float64(1024)) &^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(10)))))}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= [0 2 968 16649 0] [(bit-clear 1 0) (bit-clear 2 0) (bit-clear 1000 5) (bit-clear 16713 6) (bit-clear 1024 10)])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0), float64(992), float64(18761), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) ^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(0))))), float64((cljs_core.Int32_(float64(2)) ^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(1))))), float64((cljs_core.Int32_(float64(1000)) ^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(3))))), float64((cljs_core.Int32_(float64(16713)) ^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(11))))), float64((cljs_core.Int32_(float64(1024)) ^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(10)))))}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= [0 0 992 18761 0] [(bit-flip 1 0) (bit-flip 2 1) (bit-flip 1000 3) (bit-flip 16713 11) (bit-flip 1024 10)])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-2), float64(-3), float64(999), float64(-16714), float64(-1025)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((^cljs_core.Int32_(float64(1)))), float64((^cljs_core.Int32_(float64(2)))), float64((^cljs_core.Int32_(float64(-1000)))), float64((^cljs_core.Int32_(float64(16713)))), float64((^cljs_core.Int32_(float64(1024))))}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= [-2 -3 999 -16714 -1025] [(bit-not 1) (bit-not 2) (bit-not -1000) (bit-not 16713) (bit-not 1024)])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(1000), float64(18761), float64(1024)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) | (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(0))))), float64((cljs_core.Int32_(float64(2)) | (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(1))))), float64((cljs_core.Int32_(float64(1000)) | (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(3))))), float64((cljs_core.Int32_(float64(16713)) | (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(11))))), float64((cljs_core.Int32_(float64(1024)) | (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(10)))))}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 1000 18761 1024] [(bit-set 1 0) (bit-set 2 1) (bit-set 1000 3) (bit-set 16713 11) (bit-set 1024 10)])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true, true, false, true}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(cljs_core.Int32_(float64(1)) & (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(0)))) != 0, (cljs_core.Int32_(float64(2)) & (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(1)))) != 0, (cljs_core.Int32_(float64(1000)) & (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(3)))) != 0, (cljs_core.Int32_(float64(16713)) & (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(11)))) != 0, (cljs_core.Int32_(float64(1024)) & (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(10)))) != 0}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= [true true true false true] [(bit-test 1 0) (bit-test 2 1) (bit-test 1000 3) (bit-test 16713 11) (bit-test 1024 10)])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, true, false, false, false}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true == true, false == true, false == false, true == false, js.Undefined == true, js.Undefined == false}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= [true false true false false false] [(true? true) (true? false) (false? false) (false? true) (true? js/undefined) (false? js/undefined)])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, nil)) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (apply + nil))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.CljsCoreList_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (apply + (list)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (apply + (list 1)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Apply.X_invoke_Arity3(cljs_core.X_PLUS_, float64(1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2)))) {
			} else {
				panic((&js.Error{("Assert failed: (= 3 (apply + 1 (list 2)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(7), cljs_core.Apply.X_invoke_Arity4(cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(4)))) {
			} else {
				panic((&js.Error{("Assert failed: (= 7 (apply + 1 2 (list 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(15), cljs_core.Apply.X_invoke_Arity5(cljs_core.X_PLUS_, float64(1), float64(2), float64(4), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)))) {
			} else {
				panic((&js.Error{("Assert failed: (= 15 (apply + 1 2 4 (list 8)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(31), cljs_core.Apply.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(1), float64(2), float64(4), float64(8), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(16))}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 31 (apply + 1 2 4 8 (list 16)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(63), cljs_core.Apply.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(1), float64(2), float64(4), float64(8), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(16), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(32))}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 63 (apply + 1 2 4 8 16 (list 32)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(127), cljs_core.Apply.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(1), float64(2), float64(4), float64(8), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(16), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(64)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(32))}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 127 (apply + 1 2 4 8 16 (list 32 64)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(4950), cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Take.X_invoke_Arity2(float64(100), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{("Assert failed: (= 4950 (apply + (take 100 (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Apply.X_invoke_Arity2(cljs_core.List, cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (= () (apply list []))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), cljs_core.Apply.X_invoke_Arity2(cljs_core.List, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 3] (apply list [1 2 3]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Apply.X_invoke_Arity2(cljs_core.Apply, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 6 (apply apply [+ [1 2 3]]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Apply.X_invoke_Arity2(func(G__4980 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4980, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return ((cljs_core.Nth.X_invoke_Arity2(args, float64(0)).(float64) + cljs_core.Nth.X_invoke_Arity2(args, float64(1)).(float64)) + cljs_core.Nth.X_invoke_Arity2(args, float64(2)).(float64))
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))) {
			} else {
				panic((&js.Error{("Assert failed: (= 3 (apply (fn [& args] (+ (nth args 0) (nth args 1) (nth args 2))) (iterate inc 0)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4981 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4981, 0, func(m__ ...interface{}) interface{} {
					var m = cljs_core.Seq.Arity1IQ(m__[0])
					_ = m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [0 1 2 3 4] (take 5 (apply (fn [& m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4982 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4982, 1, func(x_m__ ...interface{}) interface{} {
					var x = x_m__[0]
					var m = cljs_core.Seq.Arity1IQ(x_m__[1])
					_, _ = x, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (take 5 (apply (fn [x & m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4983 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4983, 2, func(x_y_m__ ...interface{}) interface{} {
					var x = x_y_m__[0]
					var y = x_y_m__[1]
					var m = cljs_core.Seq.Arity1IQ(x_y_m__[2])
					_, _, _ = x, y, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 3 4 5 6] (take 5 (apply (fn [x y & m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5), float64(6), float64(7)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4984 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4984, 3, func(x_y_z_m__ ...interface{}) interface{} {
					var x = x_y_z_m__[0]
					var y = x_y_z_m__[1]
					var z = x_y_z_m__[2]
					var m = cljs_core.Seq.Arity1IQ(x_y_z_m__[3])
					_, _, _, _ = x, y, z, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [3 4 5 6 7] (take 5 (apply (fn [x y z & m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6), float64(7), float64(8)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4985 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4985, 4, func(x_y_z_a_m__ ...interface{}) interface{} {
					var x = x_y_z_a_m__[0]
					var y = x_y_z_a_m__[1]
					var z = x_y_z_a_m__[2]
					var a = x_y_z_a_m__[3]
					var m = cljs_core.Seq.Arity1IQ(x_y_z_a_m__[4])
					_, _, _, _, _ = x, y, z, a, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [4 5 6 7 8] (take 5 (apply (fn [x y z a & m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4986 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4986, 5, func(x_y_z_a_b_m__ ...interface{}) interface{} {
					var x = x_y_z_a_b_m__[0]
					var y = x_y_z_a_b_m__[1]
					var z = x_y_z_a_b_m__[2]
					var a = x_y_z_a_b_m__[3]
					var b = x_y_z_a_b_m__[4]
					var m = cljs_core.Seq.Arity1IQ(x_y_z_a_b_m__[5])
					_, _, _, _, _, _ = x, y, z, a, b, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [5 6 7 8 9] (take 5 (apply (fn [x y z a b & m] m) (iterate inc 0))))")}))
			}
			{
				var single_arity_non_variadic_4987 = func(G__4992 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4992, 3, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{})
				var multiple_arity_non_variadic_4988 = func(G__4993 *cljs_core.AFn, single_arity_non_variadic_4987 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4993, 3, func(x interface{}) interface{} {
						return x
					}, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{y, x}, nil})
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_4987)
				var single_arity_variadic_fixedargs_4989 = func(G__4994 *cljs_core.AFn, single_arity_non_variadic_4987 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_4988 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4994, 2, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{more, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_4987, multiple_arity_non_variadic_4988)
				var single_arity_variadic_nofixedargs_4990 = func(G__4995 *cljs_core.AFn, single_arity_non_variadic_4987 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_4988 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_4989 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4995, 0, func(more__ ...interface{}) interface{} {
						var more = cljs_core.Seq.Arity1IQ(more__[0])
						_ = more
						return more
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_4987, multiple_arity_non_variadic_4988, single_arity_variadic_fixedargs_4989)
				var multiple_arity_variadic_4991 = func(G__4996 *cljs_core.AFn, single_arity_non_variadic_4987 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_4988 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_4989 cljs_core.CljsCoreIFn, single_arity_variadic_nofixedargs_4990 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4996, 2, func(x interface{}) interface{} {
						return x
					}, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{y, x}, nil})
					}, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{more, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_4987, multiple_arity_non_variadic_4988, single_arity_variadic_fixedargs_4989, single_arity_variadic_nofixedargs_4990)
				_, _, _, _, _ = single_arity_non_variadic_4987, multiple_arity_non_variadic_4988, single_arity_variadic_fixedargs_4989, single_arity_variadic_nofixedargs_4990, multiple_arity_variadic_4991
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_non_variadic_4987, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic [1 2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_non_variadic_4987, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic 1 [2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_non_variadic_4987, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic 1 2 [3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_4988, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (apply multiple-arity-non-variadic [42]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_4988, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-non-variadic [1 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_4988, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-non-variadic 1 [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_4988, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic [1 2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_4988, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic 1 [2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_non_variadic_4988, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic 1 2 [3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_4989, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_fixedargs_4989, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_fixedargs_4989, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_fixedargs_4989, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_fixedargs_4989, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_4989, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 4 5] (take 3 (first (apply single-arity-variadic-fixedargs (iterate inc 1)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_4989, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (rest (apply single-arity-variadic-fixedargs (iterate inc 1))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_4990, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_nofixedargs_4990, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_nofixedargs_4990, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_nofixedargs_4990, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_nofixedargs_4990, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_4990, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (take 5 (apply single-arity-variadic-nofixedargs (iterate inc 1))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_4991, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (apply multiple-arity-variadic [42]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_4991, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-variadic [1 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_4991, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-variadic 1 [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_4991, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_4991, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_variadic_4991, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(multiple_arity_variadic_4991, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(multiple_arity_variadic_4991, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_4991, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 4 5] (take 3 (first (apply multiple-arity-variadic (iterate inc 1)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_4991, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (rest (apply multiple-arity-variadic (iterate inc 1))))")}))
				}
			}
			{
				var f1_4997 = func(f1 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(f1, 3, func() interface{} {
						return float64(0)
					}, func(a interface{}) interface{} {
						return float64(1)
					}, func(a interface{}, b interface{}) interface{} {
						return float64(2)
					}, func(a_b_c_more__ ...interface{}) interface{} {
						var a = a_b_c_more__[0]
						var b = a_b_c_more__[1]
						var c = a_b_c_more__[2]
						var more = cljs_core.Seq.Arity1IQ(a_b_c_more__[3])
						_, _, _, _ = a, b, c, more
						return float64(3)
					})
				}(&cljs_core.AFn{})
				var f2_4998 = func(f2 *cljs_core.AFn, f1_4997 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(f2, 2, func(x interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					}, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return cljs_core.Apply.X_invoke_Arity3(f1_4997, y, more)
					})
				}(&cljs_core.AFn{}, f1_4997)
				_, _ = f1_4997, f2_4998
				if cljs_core.X_EQ_.Arity2IIB(float64(1), f2_4998.X_invoke_Arity2(float64(1), float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (f2 1 2))")}))
				}
			}
			{
				var f_4999 = func(G__5000 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5000, 1, func() interface{} {
						return nil
					}, func(a_more__ ...interface{}) interface{} {
						var a = a_more__[0]
						var more = cljs_core.Seq.Arity1IQ(a_more__[1])
						_, _ = a, more
						return more
					})
				}(&cljs_core.AFn{})
				_ = f_4999
				if cljs_core.Nil_(f_4999.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (f :foo))")}))
				}
			}
			if cljs_core.Nil_(cljs_core.Array_seq.X_invoke_Arity2([]interface{}{float64(1)}, float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (array-seq (array 1) 1))")}))
			}
			{
				var f_5001 = func(G__5004 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5004, 1, func(x interface{}) interface{} {
						return (x.(float64) * float64(2))
					})
				}(&cljs_core.AFn{})
				var m_5002 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar"}, nil})
				var mf_5003 = cljs_core.With_meta.X_invoke_Arity2(f_5001, m_5002)
				_, _, _ = f_5001, m_5002, mf_5003
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(f_5001)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (meta f))")}))
				}
				if cljs_core.Fn_QMARK_.Arity1IB(mf_5003) {
				} else {
					panic((&js.Error{("Assert failed: (fn? mf)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__4554 = float64(2)
					_ = G__4554
					return mf_5003.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4554)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (mf 2))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Apply.X_invoke_Arity2(mf_5003, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (apply mf [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(mf_5003), m_5002) {
				} else {
					panic((&js.Error{("Assert failed: (= (meta mf) m)")}))
				}
			}
			{
				var a_5005 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_5005
				if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Deref.X_invoke_Arity1(a_5005)) {
				} else {
					panic((&js.Error{("Assert failed: (= 0 (deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity2(a_5005, cljs_core.Inc)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (swap! a inc))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(false, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_5005, float64(0), float64(42)))) {
				} else {
					panic((&js.Error{("Assert failed: (= false (compare-and-set! a 0 42))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(true, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_5005, float64(1), float64(7)))) {
				} else {
					panic((&js.Error{("Assert failed: (= true (compare-and-set! a 1 7))")}))
				}
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(a_5005)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (meta a))")}))
				}
				if cljs_core.Nil_(cljs_core.Get_validator.X_invoke_Arity1(a_5005)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (get-validator a))")}))
				}
			}
			{
				var a_5006 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_5006
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity3(a_5006, cljs_core.X_PLUS_, float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (swap! a + 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Swap_BANG_.X_invoke_Arity4(a_5006, cljs_core.X_PLUS_, float64(1), float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (swap! a + 1 2))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(10), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_5006, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 10 (swap! a + 1 2 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(20), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_5006, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(4)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 20 (swap! a + 1 2 3 4))")}))
				}
			}
			{
				var a_5007 = cljs_core.Atom.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "validator", Fqn: "validator", X_hash: float64(-1966190681)}), cljs_core.Coll_QMARK_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})})).(*cljs_core.CljsCoreAtom)
				_ = a_5007
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Coll_QMARK_, cljs_core.Get_validator.X_invoke_Arity1(a_5007)) {
				} else {
					panic((&js.Error{("Assert failed: (= coll? (get-validator a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_5007)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:a 1} (meta a))")}))
				}
				cljs_core.Alter_meta_BANG_.X_invoke_ArityVariadic(a_5007, cljs_core.Assoc, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}))
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_5007)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:a 1, :b 2} (meta a))")}))
				}
			}
			if cljs_core.Nil_(cljs_core.Empty.X_invoke_Arity1(nil)) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (empty nil))")}))
			}
			{
				var e_lazy_seq_5008 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCoreLazySeq{nil, func(G__5009 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5009, 0, func() interface{} {
						return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons)
					})
				}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_lazy_seq_5008
				if cljs_core.Seq_QMARK_.Arity1IB(e_lazy_seq_5008) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-lazy-seq)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_lazy_seq_5008) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-lazy-seq)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_lazy_seq_5008)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-lazy-seq))")}))
				}
			}
			{
				var e_list_5010 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_list_5010
				if cljs_core.Seq_QMARK_.Arity1IB(e_list_5010) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-list)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_list_5010) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-list)")}))
				}
			}
			{
				var e_elist_5011 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_elist_5011
				if cljs_core.Seq_QMARK_.Arity1IB(e_elist_5011) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-elist)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_elist_5011) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-elist)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), cljs_core.Get.X_invoke_Arity2(cljs_core.Meta.X_invoke_Arity1(e_elist_5011), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= :c (get (meta e-elist) :b))")}))
				}
			}
			{
				var e_cons_5012 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_cons_5012
				if cljs_core.Seq_QMARK_.Arity1IB(e_cons_5012) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-cons)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_cons_5012) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-cons)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_cons_5012)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-cons))")}))
				}
			}
			{
				var e_vec_5013 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_vec_5013
				if cljs_core.Vector_QMARK_.Arity1IB(e_vec_5013) {
				} else {
					panic((&js.Error{("Assert failed: (vector? e-vec)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_vec_5013) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-vec)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_vec_5013)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-vec))")}))
				}
			}
			{
				var e_omap_5014 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_omap_5014
				if cljs_core.Map_QMARK_.Arity1IB(e_omap_5014) {
				} else {
					panic((&js.Error{("Assert failed: (map? e-omap)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_omap_5014) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-omap)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_omap_5014)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-omap))")}))
				}
			}
			{
				var e_hmap_5015 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_hmap_5015
				if cljs_core.Map_QMARK_.Arity1IB(e_hmap_5015) {
				} else {
					panic((&js.Error{("Assert failed: (map? e-hmap)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_hmap_5015) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-hmap)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_hmap_5015)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-hmap))")}))
				}
			}
			{
				var a_5016 = cljs_core.Atom.X_invoke_Arity1(nil).(*cljs_core.CljsCoreAtom)
				_ = a_5016
				if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (try 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5017 interface{}) {
					defer func() {
						if e4555 := recover(); e4555 != nil {
							if cljs_core.Value_(e4555).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4555
									_ = e
									return__5017 = float64(2)
								}
							} else {
								panic(e4555)

							}
						}
					}()
					{
						panic((&js.Error{}))
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (try 1 (throw (js/Error.)) (catch js/Error e 2)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5018 interface{}) {
					defer func() {
						if e4556 := recover(); e4556 != nil {
							if cljs_core.Value_(e4556).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4556
									_ = e
									return__5018 = float64(2)
								}
							} else {
								panic(e4556)

							}
						}
					}()
					{
						panic((&js.Error{}))
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (try 1 (throw (js/Error.)) (catch js/Error e 1 2)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5019 interface{}) {
					defer func() {
						if e4557 := recover(); e4557 != nil {
							if cljs_core.Value_(e4557).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4557
									_ = e
									return__5019 = float64(2)
								}
							} else {
								{
									var e = e4557
									_ = e
									return__5019 = float64(3)
								}

							}
						}
					}()
					{
						panic((&js.Error{}))
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (try 1 (throw (js/Error.)) (catch js/Error e 2) (catch :default e 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__5020 interface{}) {
					defer func() {
						if e4558 := recover(); e4558 != nil {
							if cljs_core.Value_(e4558).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4558
									_ = e
									return__5020 = float64(2)
								}
							} else {
								{
									var e = e4558
									_ = e
									return__5020 = float64(3)
								}

							}
						}
					}()
					{
						panic(true)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (try 1 (throw true) (catch js/Error e 2) (catch :default e 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5021 interface{}) {
					defer func() {
						if e4559 := recover(); e4559 != nil {
							if cljs_core.Value_(e4559).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4559
									_ = e
									return__5021 = float64(3)
								}
							} else {
								{
									var e = e4559
									_ = e
									return__5021 = e
								}

							}
						}
					}()
					{
						panic(float64(2))
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (try 1 (throw 2) (catch js/Error e 3) (catch :default e e)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), func() float64 {
					defer func() {
						cljs_core.Reset_BANG_.X_invoke_Arity2(a_5016, float64(42))
					}()
					{
						return float64(1)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (try 1 (finally (reset! a 42))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Deref.X_invoke_Arity1(a_5016)) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (deref a))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), cljs_core.Seq_(cljs_core.Nthnext.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)))) {
			} else {
				panic((&js.Error{("Assert failed: (= [3] (nthnext [1 2 3] 2))")}))
			}
			{
				var v_5022 = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})
				_ = v_5022
				if cljs_core.X_EQ_.Arity2IIB(v_5022, func() *cljs_core.CljsCoreLazySeq {
					var iter__929__auto__ = func(iter__4560 *cljs_core.AFn, v_5022 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__4560, 1, func(s__4561 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__5023 *cljs_core.AFn, v_5022 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__5023, 0, func() interface{} {
									{
										var s__4561___1 interface{} = s__4561
										_ = s__4561___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4561___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__4561___2 = temp__4222__auto__
														_ = s__4561___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4561___2) {
															{
																var c__927__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4561___2)
																var size__928__auto__ = cljs_core.Count.X_invoke_Arity1(c__927__auto__).(float64)
																var b__4563 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__928__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__927__auto__, size__928__auto__, b__4563
																if func() bool {
																	var i__4562 = float64(0)
																	_ = i__4562
																	for {
																		if i__4562 < size__928__auto__ {
																			{
																				var e = c__927__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4562)
																				_ = e
																				cljs_core.Chunk_append.X_invoke_Arity2(b__4563, e)
																				i__4562 = (i__4562 + float64(1))
																				continue
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4563), iter__4560.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4561___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4563), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__4561___2)
																_ = e
																return cljs_core.Cons.X_invoke_Arity2(e, iter__4560.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4561___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
															}
														}
													}
												} else {
													return nil
												}
											}
										}
									}
								})
							}(&cljs_core.AFn{}, v_5022), nil, nil})
						})
					}(&cljs_core.AFn{}, v_5022)
					_ = iter__929__auto__
					return iter__929__auto__.X_invoke_Arity1(v_5022).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= v (for [e v] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(9)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__929__auto__ = func(iter__4566 *cljs_core.AFn, v_5022 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__4566, 1, func(s__4567 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__5024 *cljs_core.AFn, v_5022 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__5024, 0, func() interface{} {
									{
										var s__4567___1 interface{} = s__4567
										_ = s__4567___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4567___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__4567___2 = temp__4222__auto__
														_ = s__4567___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4567___2) {
															{
																var c__927__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4567___2)
																var size__928__auto__ = cljs_core.Count.X_invoke_Arity1(c__927__auto__).(float64)
																var b__4569 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__928__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__927__auto__, size__928__auto__, b__4569
																if func() bool {
																	var i__4568 = float64(0)
																	_ = i__4568
																	for {
																		if i__4568 < size__928__auto__ {
																			{
																				var e = c__927__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4568)
																				_ = e
																				{
																					var m = (e.(float64) * e.(float64))
																					_ = m
																					cljs_core.Chunk_append.X_invoke_Arity2(b__4569, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																					i__4568 = (i__4568 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4569), iter__4566.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4567___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4569), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__4567___2)
																_ = e
																{
																	var m = (e.(float64) * e.(float64))
																	_ = m
																	return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__4566.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4567___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
																}
															}
														}
													}
												} else {
													return nil
												}
											}
										}
									}
								})
							}(&cljs_core.AFn{}, v_5022), nil, nil})
						})
					}(&cljs_core.AFn{}, v_5022)
					_ = iter__929__auto__
					return iter__929__auto__.X_invoke_Arity1(v_5022).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [[1 1] [2 4] [3 9]] (for [e v :let [m (* e e)]] [e m]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__929__auto__ = func(iter__4572 *cljs_core.AFn, v_5022 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__4572, 1, func(s__4573 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__5025 *cljs_core.AFn, v_5022 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__5025, 0, func() interface{} {
									{
										var s__4573___1 interface{} = s__4573
										_ = s__4573___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4573___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__4573___2 = temp__4222__auto__
														_ = s__4573___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4573___2) {
															{
																var c__927__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4573___2)
																var size__928__auto__ = cljs_core.Count.X_invoke_Arity1(c__927__auto__).(float64)
																var b__4575 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__928__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__927__auto__, size__928__auto__, b__4575
																if cljs_core.Truth_(func() interface{} {
																	var i__4574 = float64(0)
																	_ = i__4574
																	for {
																		if i__4574 < size__928__auto__ {
																			{
																				var e = c__927__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4574)
																				_ = e
																				if e.(float64) < float64(3) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__4575, e)
																					i__4574 = (i__4574 + float64(1))
																					continue
																				} else {
																					return nil
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}()) {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4575), iter__4572.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4573___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4575), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__4573___2)
																_ = e
																if e.(float64) < float64(3) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__4572.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4573___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
																} else {
																	return nil
																}
															}
														}
													}
												} else {
													return nil
												}
											}
										}
									}
								})
							}(&cljs_core.AFn{}, v_5022), nil, nil})
						})
					}(&cljs_core.AFn{}, v_5022)
					_ = iter__929__auto__
					return iter__929__auto__.X_invoke_Arity1(v_5022).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2] (for [e v :while (< e 3)] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__929__auto__ = func(iter__4578 *cljs_core.AFn, v_5022 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__4578, 1, func(s__4579 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__5026 *cljs_core.AFn, v_5022 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__5026, 0, func() interface{} {
									{
										var s__4579___1 interface{} = s__4579
										_ = s__4579___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4579___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__4579___2 = temp__4222__auto__
														_ = s__4579___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4579___2) {
															{
																var c__927__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4579___2)
																var size__928__auto__ = cljs_core.Count.X_invoke_Arity1(c__927__auto__).(float64)
																var b__4581 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__928__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__927__auto__, size__928__auto__, b__4581
																if func() bool {
																	var i__4580 = float64(0)
																	_ = i__4580
																	for {
																		if i__4580 < size__928__auto__ {
																			{
																				var e = c__927__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4580)
																				_ = e
																				if e.(float64) > float64(2) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__4581, e)
																					i__4580 = (i__4580 + float64(1))
																					continue
																				} else {
																					i__4580 = (i__4580 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4581), iter__4578.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4579___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4581), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__4579___2)
																_ = e
																if e.(float64) > float64(2) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__4578.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4579___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
																} else {
																	s__4579___1 = cljs_core.Rest.Arity1IQ(s__4579___2)
																	continue
																}
															}
														}
													}
												} else {
													return nil
												}
											}
										}
									}
								})
							}(&cljs_core.AFn{}, v_5022), nil, nil})
						})
					}(&cljs_core.AFn{}, v_5022)
					_ = iter__929__auto__
					return iter__929__auto__.X_invoke_Arity1(v_5022).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [3] (for [e v :when (> e 2)] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__929__auto__ = func(iter__4584 *cljs_core.AFn, v_5022 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__4584, 1, func(s__4585 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__5027 *cljs_core.AFn, v_5022 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__5027, 0, func() interface{} {
									{
										var s__4585___1 interface{} = s__4585
										_ = s__4585___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4585___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__4585___2 = temp__4222__auto__
														_ = s__4585___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4585___2) {
															{
																var c__927__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4585___2)
																var size__928__auto__ = cljs_core.Count.X_invoke_Arity1(c__927__auto__).(float64)
																var b__4587 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__928__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__927__auto__, size__928__auto__, b__4587
																if cljs_core.Truth_(func() interface{} {
																	var i__4586 = float64(0)
																	_ = i__4586
																	for {
																		if i__4586 < size__928__auto__ {
																			{
																				var e = c__927__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4586)
																				_ = e
																				if e.(float64) < float64(3) {
																					{
																						var m = (e.(float64) * e.(float64))
																						_ = m
																						cljs_core.Chunk_append.X_invoke_Arity2(b__4587, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																						i__4586 = (i__4586 + float64(1))
																						continue
																					}
																				} else {
																					return nil
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}()) {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4587), iter__4584.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4585___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4587), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__4585___2)
																_ = e
																if e.(float64) < float64(3) {
																	{
																		var m = (e.(float64) * e.(float64))
																		_ = m
																		return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__4584.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4585___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
																	}
																} else {
																	return nil
																}
															}
														}
													}
												} else {
													return nil
												}
											}
										}
									}
								})
							}(&cljs_core.AFn{}, v_5022), nil, nil})
						})
					}(&cljs_core.AFn{}, v_5022)
					_ = iter__929__auto__
					return iter__929__auto__.X_invoke_Arity1(v_5022).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [[1 1] [2 4]] (for [e v :while (< e 3) :let [m (* e e)]] [e m]))")}))
				}
			}
			if cljs_core.Not_EQ_.Arity2IIB(float64(1), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (not= 1 2)")}))
			}
			if !(cljs_core.Not_EQ_.Arity2IIB(float64(1), float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (not (not= 1 1))")}))
			}
			if cljs_core.Not.Arity1IB(cljs_core.Not_empty.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (not (not-empty []))")}))
			}
			if cljs_core.Boolean.Arity1IB(cljs_core.Not_empty.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (boolean (not-empty [1 2 3]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("joel", cljs_core.Min_key.X_invoke_ArityVariadic(cljs_core.Count, "joel", "tom servo", cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"crooooooooow"}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"joel\" (min-key count \"joel\" \"tom servo\" \"crooooooooow\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("crooooooooow", cljs_core.Max_key.X_invoke_ArityVariadic(cljs_core.Count, "joel", "tom servo", cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"crooooooooow"}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"crooooooooow\" (max-key count \"joel\" \"tom servo\" \"crooooooooow\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Partition_all.X_invoke_Arity2(float64(4), (&cljs_core.CljsCorePersistentVector{nil, float64(9), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9)}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7), float64(8)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(9)}, nil})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (partition-all 4 [1 2 3 4 5 6 7 8 9]) [[1 2 3 4] [5 6 7 8] [9]])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Partition_all.X_invoke_Arity3(float64(4), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(9), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9)}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5), float64(6)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7), float64(8)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(7), float64(8), float64(9)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(9)}, nil})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (partition-all 4 2 [1 2 3 4 5 6 7 8 9]) [[1 2 3 4] [3 4 5 6] [5 6 7 8] [7 8 9] [9]])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true}, nil}), cljs_core.Take_while.X_invoke_Arity2(cljs_core.True_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true, float64(2), float64(3), float64(4)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [true true] (take-while true? [true true 2 3 4]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2), float64(4), float64(6), float64(8), float64(10)}, nil}), cljs_core.Take_nth.X_invoke_Arity2(float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(11), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [0 2 4 6 8 10] (take-nth 2 [0 1 2 3 4 5 6 7 8 9 10]))")}))
			}
			{
				var a10_5028 = cljs_core.Partial.X_invoke_Arity2(cljs_core.X_PLUS_, float64(10)).(cljs_core.CljsCoreIFn)
				var a20_5029 = cljs_core.Partial.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), float64(10)).(cljs_core.CljsCoreIFn)
				var a21_5030 = cljs_core.Partial.X_invoke_Arity4(cljs_core.X_PLUS_, float64(10), float64(10), float64(1)).(cljs_core.CljsCoreIFn)
				var a22_5031 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)})).(cljs_core.CljsCoreIFn)
				var a23_5032 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(1)})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = a10_5028, a20_5029, a21_5030, a22_5031, a23_5032
				if cljs_core.X_EQ_.Arity2IIB(float64(110), func() interface{} {
					var G__4590 = float64(100)
					_ = G__4590
					return a10_5028.X_invoke_Arity1(G__4590)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 110 (a10 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(120), func() interface{} {
					var G__4591 = float64(100)
					_ = G__4591
					return a20_5029.X_invoke_Arity1(G__4591)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 120 (a20 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(121), func() interface{} {
					var G__4592 = float64(100)
					_ = G__4592
					return a21_5030.X_invoke_Arity1(G__4592)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 121 (a21 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(122), func() interface{} {
					var G__4593 = float64(100)
					_ = G__4593
					return a22_5031.X_invoke_Arity1(G__4593)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 122 (a22 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(123), func() interface{} {
					var G__4594 = float64(100)
					_ = G__4594
					return a23_5032.X_invoke_Arity1(G__4594)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 123 (a23 100))")}))
				}
			}
			{
				var n2_5033 = cljs_core.Comp.X_invoke_Arity2(cljs_core.First, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n3_5034 = cljs_core.Comp.X_invoke_Arity3(cljs_core.First, cljs_core.Rest, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n4_5035 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n5_5036 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n6_5037 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = n2_5033, n3_5034, n4_5035, n5_5036, n6_5037
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() interface{} {
					var G__4595 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__4595
					return n2_5033.X_invoke_Arity1(G__4595)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (n2 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() interface{} {
					var G__4596 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__4596
					return n3_5034.X_invoke_Arity1(G__4596)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (n3 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__4597 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__4597
					return n4_5035.X_invoke_Arity1(G__4597)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (n4 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(5), func() interface{} {
					var G__4598 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__4598
					return n5_5036.X_invoke_Arity1(G__4598)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 5 (n5 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), func() interface{} {
					var G__4599 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__4599
					return n6_5037.X_invoke_Arity1(G__4599)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (n6 [1 2 3 4 5 6 7]))")}))
				}
			}
			{
				var sf_5038 = cljs_core.Some_fn.X_invoke_Arity3(cljs_core.Number_QMARK_, cljs_core.Keyword_QMARK_, cljs_core.Symbol_QMARK_).(cljs_core.CljsCoreIFn)
				_ = sf_5038
				if cljs_core.Truth_(func() interface{} {
					var G__4600 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					var G__4601 = float64(1)
					_, _ = G__4600, G__4601
					return sf_5038.X_invoke_Arity2(G__4600, G__4601)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf :foo 1)")}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__4602 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					_ = G__4602
					return sf_5038.X_invoke_Arity1(G__4602)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf :foo)")}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__4603 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})
					var G__4604 = float64(1)
					_, _ = G__4603, G__4604
					return sf_5038.X_invoke_Arity2(G__4603, G__4604)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf (quote bar) 1)")}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__4605 = cljs_core.CljsCorePersistentVector_EMPTY
					var G__4606 = cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)
					_, _ = G__4605, G__4606
					return sf_5038.X_invoke_Arity2(G__4605, G__4606)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (not (sf [] ()))")}))
				}
			}
			{
				var ep_5039 = cljs_core.Every_pred.X_invoke_Arity2(cljs_core.Number_QMARK_, cljs_core.Zero_QMARK_).(cljs_core.CljsCoreIFn)
				_ = ep_5039
				if cljs_core.Truth_(func() interface{} {
					var G__4607 = float64(0)
					var G__4608 = float64(0)
					var G__4609 = float64(0)
					_, _, _ = G__4607, G__4608, G__4609
					return ep_5039.X_invoke_Arity3(G__4607, G__4608, G__4609)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (ep 0 0 0)")}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__4610 = float64(1)
					var G__4611 = float64(2)
					var G__4612 = float64(3)
					var G__4613 = float64(0)
					_, _, _, _ = G__4610, G__4611, G__4612, G__4613
					return ep_5039.X_invoke_Arity4(G__4610, G__4611, G__4612, G__4613)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (not (ep 1 2 3 0))")}))
				}
			}
			if cljs_core.Truth_(cljs_core.Complement.X_invoke_Arity1(cljs_core.Number_QMARK_).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
			} else {
				panic((&js.Error{("Assert failed: ((complement number?) :foo)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})}, nil}), cljs_core.Juxt.X_invoke_Arity3(cljs_core.First, cljs_core.Rest, cljs_core.Seq).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 [2 3] [1 2 3]] ((juxt first rest seq) [1 2 3]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(5), func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(float64(1), float64(2)), float64(3)), float64(4)), float64(5))) {
			} else {
				panic((&js.Error{("Assert failed: (= 5 (max 1 2 3 4 5))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(5), func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(float64(5), float64(4)), float64(3)), float64(2)), float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 5 (max 5 4 3 2 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(5.5, func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(float64(1), float64(2)), float64(3)), float64(4)), float64(5)), 5.5)) {
			} else {
				panic((&js.Error{("Assert failed: (= 5.5 (max 1 2 3 4 5 5.5))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(float64(5), float64(4)), float64(3)), float64(2)), float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (min 5 4 3 2 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(float64(1), float64(2)), float64(3)), float64(4)), float64(5))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (min 1 2 3 4 5))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(0.5, func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(float64(5), float64(4)), float64(3)), 0.5), float64(2)), float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 0.5 (min 5 4 3 0.5 2 1))")}))
			}
			if cljs_core.Truth_(cljs_core.Set.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (set [])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, cljs_core.Set.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (= #{} (set []))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, cljs_core.CljsCorePersistentHashSet_EMPTY) {
			} else {
				panic((&js.Error{("Assert failed: (= #{} (hash-set))")}))
			}
			if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentHashSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentHashSet (type (hash-set)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"foo", nil}, nil}, nil}), cljs_core.Set.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo"}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= #{\"foo\"} (set [\"foo\"]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"foo", nil}, nil}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"foo", nil}, nil}), nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= #{\"foo\"} (hash-set \"foo\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 3 2} #{1 3 2})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(7), float64(8)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6)}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(9), float64(10)}, true).(*cljs_core.CljsCorePersistentHashSet), cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(7), float64(8)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6)}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(9), float64(10)}, true).(*cljs_core.CljsCorePersistentHashSet)) {
			} else {
				panic((&js.Error{("Assert failed: (= #{{7 8} [4 5 6] #{1 3 2} 9 10} #{{7 8} [4 5 6] #{1 3 2} 9 10})")}))
			}
			if !(cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{nil, float64(0), cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashSet_EMPTY}, true).(*cljs_core.CljsCorePersistentHashSet), cljs_core.CljsCorePersistentHashSet_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (not (= #{nil 0 [] {} #{}} #{}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{nil, float64(0), cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashSet_EMPTY}, true).(*cljs_core.CljsCorePersistentHashSet)).(float64), float64(5)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count #{nil 0 [] {} #{}}) 5)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), float64(1)), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (conj #{1} 1) #{1})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), float64(2)), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (conj #{1} 2) #{1 2})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{float64(1), nil, float64(4), nil, float64(3), nil, float64(2), nil}, nil}, nil}).X_empty_Arity1()) {
			} else {
				panic((&js.Error{("Assert failed: (= #{} (-empty #{1 4 3 2}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(5), []interface{}{float64(1), nil, float64(4), nil, float64(3), nil, float64(2), nil, float64(5), nil}, nil}, nil})), float64(15)) {
			} else {
				panic((&js.Error{("Assert failed: (= (reduce + #{1 4 3 2 5}) 15)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{float64(1), nil, float64(4), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(4))) {
			} else {
				panic((&js.Error{("Assert failed: (= 4 (get #{1 4 3 2} 4))")}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{float64(1), nil, float64(4), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(4)) {
			} else {
				panic((&js.Error{("Assert failed: (contains? #{1 4 3 2} 4)")}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{nil, float64(0), cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashSet_EMPTY}, true).(*cljs_core.CljsCorePersistentHashSet), cljs_core.CljsCorePersistentArrayMap_EMPTY) {
			} else {
				panic((&js.Error{("Assert failed: (contains? #{nil 0 [] {} #{}} {})")}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})}, true).(*cljs_core.CljsCorePersistentHashSet), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (contains? #{[1 2 3]} [1 2 3])")}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}).X_disjoin_Arity2(float64(3)), float64(3))) {
			} else {
				panic((&js.Error{("Assert failed: (not (contains? (-disjoin #{1 3 2} 3) 3))")}))
			}
			if float64(-1) < float64(0) {
			} else {
				panic((&js.Error{("Assert failed: (neg? -1)")}))
			}
			if !(float64(1) < float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (not (neg? 1))")}))
			}
			if -1.765 < float64(0) {
			} else {
				panic((&js.Error{("Assert failed: (neg? -1.765)")}))
			}
			if !(float64(0) < float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (not (neg? 0))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, true, false, true, false, true, false}, nil}), cljs_core.Map_.X_invoke_Arity2(cljs_core.Integer_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), 1.00001, float64(2023), cljs_core.CljsCorePersistentVector_EMPTY, (float64(88) - float64(1001991881)), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(0), "0"}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [true false true false true false true false] (map integer? [1 1.00001 2023 [] (- 88 1001991881) :foo 0 \"0\"]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, true, false, true, false}, nil}), cljs_core.Map_.X_invoke_Arity2(cljs_core.Odd_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(-1), float64(0)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [true false true false true false] (map odd? [1 2 3 4 -1 0]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, true, false, true, true}, nil}), cljs_core.Map_.X_invoke_Arity2(cljs_core.Even_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(-2), float64(0)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [true false true false true true] (map even? [2 3 4 5 -2 0]))")}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})) {
			} else {
				panic((&js.Error{("Assert failed: (contains? {:a 1, :b 2} :a)")}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "z", Fqn: "z", X_hash: float64(-789527183)}))) {
			} else {
				panic((&js.Error{("Assert failed: (not (contains? {:a 1, :b 2} :z))")}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil}), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (contains? [5 6 7] 1)")}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil}), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (contains? [5 6 7] 2)")}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil}), float64(3))) {
			} else {
				panic((&js.Error{("Assert failed: (not (contains? [5 6 7] 3))")}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil})).([]interface{}), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (contains? (to-array [5 6 7]) 1)")}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil})).([]interface{}), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (contains? (to-array [5 6 7]) 2)")}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil})).([]interface{}), float64(3))) {
			} else {
				panic((&js.Error{("Assert failed: (not (contains? (to-array [5 6 7]) 3))")}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB(nil, float64(42))) {
			} else {
				panic((&js.Error{("Assert failed: (not (contains? nil 42))")}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB("f", float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (contains? \"f\" 0)")}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB("f", float64(55))) {
			} else {
				panic((&js.Error{("Assert failed: (not (contains? \"f\" 55))")}))
			}
			if cljs_core.Truth_(cljs_core.Distinct_QMARK_.X_invoke_ArityVariadic(float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
			} else {
				panic((&js.Error{("Assert failed: (distinct? 1 2 3)")}))
			}
			if !(cljs_core.Truth_(cljs_core.Distinct_QMARK_.X_invoke_ArityVariadic(float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(1)})))) {
			} else {
				panic((&js.Error{("Assert failed: (not (distinct? 1 2 3 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct ()) ())")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(*cljs_core.CljsCoreList)).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct (quote (1))) (quote (1)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(1), float64(1), float64(1)})).(*cljs_core.CljsCoreList)).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct (quote (1 2 3 1 1 1))) (quote (1 2 3)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(1), float64(2)}, nil})).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [1 1 1 2]) (quote (1 2)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(1), float64(2)}, nil})).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [1 2 1 2]) (quote (1 2)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1("a").(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct \"a\") [\"a\"])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1("abcabab").(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a", "b", "c"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct \"abcabab\") [\"a\" \"b\" \"c\"])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"abc", "abc"}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"abc"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [\"abc\" \"abc\"]) [\"abc\"])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, nil}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [nil nil]) [nil])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{0.0, 0.0}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{0.0}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [0.0 0.0]) [0.0])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "sym", Str: "sym", X_hash: float64(195671222), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "sym", Str: "sym", X_hash: float64(195671222), X_meta: nil})}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "sym", Str: "sym", X_hash: float64(195671222), X_meta: nil})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [(quote sym) (quote sym)]) (quote [sym]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "kw", Fqn: "kw", X_hash: float64(1158308175)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "kw", Fqn: "kw", X_hash: float64(1158308175)})}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "kw", Fqn: "kw", X_hash: float64(1158308175)})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [:kw :kw]) [:kw])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42), float64(42)}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [42 42]) [42])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.CljsCorePersistentVector_EMPTY}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentVector_EMPTY}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [[] []]) [[]])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [(quote (1 2)) (quote (1 2))]) (quote [(1 2)]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [() ()]) [()])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [[1 2] [1 2]]) [[1 2]])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [{:a 1, :b 2} {:a 1, :b 2}]) [{:a 1, :b 2}])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentArrayMap_EMPTY}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [{} {}]) [{}])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil})}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [#{1 2} #{1 2}]) [#{1 2}])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentHashSet_EMPTY, cljs_core.CljsCorePersistentHashSet_EMPTY}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentHashSet_EMPTY}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (distinct [#{} #{}]) [#{}])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_find.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo").(*js.RegExp), "foo bar foo baz foo zot"), "foo") {
			} else {
				panic((&js.Error{("Assert failed: (= (re-find (re-pattern \"foo\") \"foo bar foo baz foo zot\") \"foo\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_find.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("f(.)o").(*js.RegExp), "foo bar foo baz foo zot"), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (re-find (re-pattern \"f(.)o\") \"foo bar foo baz foo zot\") [\"foo\" \"o\"])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo").(*js.RegExp), "foo"), "foo") {
			} else {
				panic((&js.Error{("Assert failed: (= (re-matches (re-pattern \"foo\") \"foo\") \"foo\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo").(*js.RegExp), "foo bar foo baz foo zot"), nil) {
			} else {
				panic((&js.Error{("Assert failed: (= (re-matches (re-pattern \"foo\") \"foo bar foo baz foo zot\") nil)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo.*").(*js.RegExp), "foo bar foo baz foo zot"), "foo bar foo baz foo zot") {
			} else {
				panic((&js.Error{("Assert failed: (= (re-matches (re-pattern \"foo.*\") \"foo bar foo baz foo zot\") \"foo bar foo baz foo zot\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_seq.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo").(*js.RegExp), "foo bar foo baz foo zot"), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2("foo").(cljs_core.CljsCoreICollection).X_conj_Arity2("foo").(cljs_core.CljsCoreICollection).X_conj_Arity2("foo")) {
			} else {
				panic((&js.Error{("Assert failed: (= (re-seq (re-pattern \"foo\") \"foo bar foo baz foo zot\") (list \"foo\" \"foo\" \"foo\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_seq.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("f(.)o").(*js.RegExp), "foo bar foo baz foo zot"), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (re-seq (re-pattern \"f(.)o\") \"foo bar foo baz foo zot\") (list [\"foo\" \"o\"] [\"foo\" \"o\"] [\"foo\" \"o\"]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("(?i)foo").(*js.RegExp), "Foo"), "Foo") {
			} else {
				panic((&js.Error{("Assert failed: (= (re-matches (re-pattern \"(?i)foo\") \"Foo\") \"Foo\")")}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{"#\"(?:)\"", nil, "#\"\"", nil}, nil}, nil}).X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.RegExp{Pattern: ``, Flags: ``})})).(string))) {
			} else {
				panic((&js.Error{("Assert failed: (#{\"#\\\"(?:)\\\"\" \"#\\\"\\\"\"} (pr-str #\"\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__4614 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4614, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4614, float64(1), nil)
				_, _, _ = vec__4614, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] [1 2]] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), func() cljs_core.CljsCoreISet {
				var vec__4615 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4615, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4615, float64(1), nil)
				_, _, _ = vec__4615, a, b
				return cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{a, b}, true).(*cljs_core.CljsCorePersistentHashSet)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 2} (let [[a b] [1 2]] #{a b}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__4616 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__4616___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4616) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4616)
					} else {
						return map__4616
					}
				}()
				var a = cljs_core.Get.X_invoke_Arity2(map__4616___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				var b = cljs_core.Get.X_invoke_Arity2(map__4616___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				_, _, _, _ = map__4616, map__4616___1, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (let [{a :a, b :b} {:a 1, :b 2}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__4617 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__4617___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4617) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4617)
					} else {
						return map__4617
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__4617___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__4617___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__4617, map__4617___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (let [{:keys [a b]} {:a 1, :b 2}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__4618 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4618, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4618, float64(1), nil)
				var v = vec__4618
				_, _, _, _ = vec__4618, a, b, v
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b, v}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 [1 2]] (let [[a b :as v] [1 2]] [a b v]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(42)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__4619 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__4619___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4619) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4619)
					} else {
						return map__4619
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity3(map__4619___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(42))
				var a = cljs_core.Get.X_invoke_Arity2(map__4619___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__4619, map__4619___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 42] (let [{:keys [a b], :or {b 42}} {:a 1}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), nil}, nil}), func() cljs_core.CljsCoreIVector {
				var map__4620 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__4620___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4620) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4620)
					} else {
						return map__4620
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__4620___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__4620___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__4620, map__4620___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 nil] (let [{:keys [a b], :or {c 42}} {:a 1}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__4621 = cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4621, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4621, float64(1), nil)
				_, _, _ = vec__4621, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] (quote (1 2))] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIMap {
				var vec__4622 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4622, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4622, float64(1), nil)
				_, _, _ = vec__4622, a, b
				return cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{a, b}, true, false).(*cljs_core.CljsCorePersistentArrayMap)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= {1 2} (let [[a b] [1 2]] {a b}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__4623 = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4623, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4623, float64(1), nil)
				_, _, _ = vec__4623, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] (seq [1 2])] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(1)}, nil})}, nil})}, nil}), cljs_core.Update_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(0)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}), cljs_core.Inc)) {
			} else {
				panic((&js.Error{("Assert failed: (= {:foo {:bar {:baz 1}}} (update-in {:foo {:bar {:baz 0}}} [:foo :bar :baz] inc))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(10)}, nil}), cljs_core.Update_in.X_invoke_Arity4((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}), cljs_core.X_PLUS_, float64(7))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 10} (update-in {:foo 1, :bar 2, :baz 3} [:baz] + 7))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(3)}, nil})}, nil}), cljs_core.Update_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}), cljs_core.Inc)) {
			} else {
				panic((&js.Error{("Assert failed: (= [{:foo 1, :bar 2} {:foo 1, :bar 3}] (update-in [{:foo 1, :bar 2} {:foo 1, :bar 2}] [1 :bar] inc))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(3)}, nil})}, nil})}, nil}), cljs_core.Update_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}), cljs_core.Inc)) {
			} else {
				panic((&js.Error{("Assert failed: (= [{:foo {:bar 2}} {:foo {:bar 3}}] (update-in [{:foo {:bar 2}} {:foo {:bar 2}}] [1 :foo :bar] inc))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(100)}, nil})}, nil})}, nil}), cljs_core.Assoc_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(0)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}), float64(100))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:foo {:bar {:baz 100}}} (assoc-in {:foo {:bar {:baz 0}}} [:foo :bar :baz] 100))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(100)}, nil}), cljs_core.Assoc_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}), float64(100))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 100} (assoc-in {:foo 1, :bar 2, :baz 3} [:baz] 100))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(100)}, nil})}, nil})}, nil})}, nil}), cljs_core.Assoc_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}), float64(100))) {
			} else {
				panic((&js.Error{("Assert failed: (= [{:foo [{:bar 2} {:baz 3}]} {:foo [{:bar 2} {:baz 100}]}] (assoc-in [{:foo [{:bar 2} {:baz 3}]} {:foo [{:bar 2} {:baz 3}]}] [1 :foo 1 :baz] 100))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(100)}, nil})}, nil}), cljs_core.Assoc_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}), float64(100))) {
			} else {
				panic((&js.Error{("Assert failed: (= [{:foo 1, :bar 2} {:foo 1, :bar 100}] (assoc-in [{:foo 1, :bar 2} {:foo 1, :bar 2}] [1 :bar] 100))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Get_in.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (get-in {:foo 1, :bar 2} [:foo]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Get_in.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (get-in {:foo {:bar 2}} [:foo :bar]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Get_in.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (get-in [{:foo 1} {:foo 2}] [0 :foo]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get_in.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(1)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "buzz", Fqn: "buzz", X_hash: float64(298530064)}), float64(2)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "buzz", Fqn: "buzz", X_hash: float64(298530064)}), float64(4)}, nil})}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "buzz", Fqn: "buzz", X_hash: float64(298530064)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 4 (get-in [{:foo 1, :bar [{:baz 1} {:buzz 2}]} {:foo 3, :bar [{:baz 3} {:buzz 4}]}] [1 :bar 1 :buzz]))")}))
			}
			{
				var a_5040 = cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{})
				_ = a_5040
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20), float64(30)}, nil}), cljs_core.Seq.Arity1IQ(func() []interface{} {
					var a__1048__auto__ = a_5040
					var ret = cljs_core.Aclone.X_invoke_Arity1(a__1048__auto__).([]interface{})
					_, _ = a__1048__auto__, ret
					{
						var i = float64(0)
						_ = i
						for {
							if i < float64(len(a__1048__auto__)) {
								ret[int(i)] = (float64(10) * (a_5040[int(i)]).(float64))
								i = (i + float64(1))
								continue
							} else {
								return ret
							}
						}
					}
				}())) {
				} else {
					panic((&js.Error{("Assert failed: (= [10 20 30] (seq (amap a i ret (* 10 (aget a i)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), func() float64 {
					var a__1054__auto__ = a_5040
					_ = a__1054__auto__
					{
						var i = float64(0)
						var ret = float64(0)
						_, _ = i, ret
						for {
							if i < float64(len(a__1054__auto__)) {
								i, ret = (i + float64(1)), (ret + (a_5040[int(i)]).(float64))
								continue
							} else {
								return ret
							}
						}
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (areduce a i ret 0 (+ ret (aget a i))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_5040), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (seq a) (seq (to-array [1 2 3])))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), func() float64 { a_5040[int(float64(0))] = float64(42); return a_5040[int(float64(0))].(float64) }()) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (aset a 0 42))")}))
				}
				if cljs_core.Not_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_5040), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
				} else {
					panic((&js.Error{("Assert failed: (not= (seq a) (seq (to-array [1 2 3])))")}))
				}
			}
			{
				var a_5041 = []interface{}{[]interface{}{float64(1), float64(2), float64(3)}, []interface{}{float64(4), float64(5), float64(6)}}
				_ = a_5041
				if cljs_core.X_EQ_.Arity2IIB((a_5041[int(float64(0))].([]interface{})[int(float64(1))]), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (aget a 0 1) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity3(cljs_core.Aget, a_5041, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil})), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (apply aget a [0 1]) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((a_5041[int(float64(1))].([]interface{})[int(float64(1))]), float64(5)) {
				} else {
					panic((&js.Error{("Assert failed: (= (aget a 1 1) 5)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity3(cljs_core.Aget, a_5041, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), float64(5)) {
				} else {
					panic((&js.Error{("Assert failed: (= (apply aget a [1 1]) 5)")}))
				}
				a_5041[int(float64(0))].([]interface{})[int(float64(0))] = "foo"
				if cljs_core.X_EQ_.Arity2IIB((a_5041[int(float64(0))].([]interface{})[int(float64(0))]), "foo") {
				} else {
					panic((&js.Error{("Assert failed: (= (aget a 0 0) \"foo\")")}))
				}
				cljs_core.Apply.X_invoke_Arity3(cljs_core.Aset, a_5041, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0), "bar"}, nil}))
				if cljs_core.X_EQ_.Arity2IIB((a_5041[int(float64(0))].([]interface{})[int(float64(0))]), "bar") {
				} else {
					panic((&js.Error{("Assert failed: (= (aget a 0 0) \"bar\")")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Sort.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(3), float64(1), float64(4), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (sort [5 3 1 4 2]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Sort.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(3), float64(1), float64(4), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (sort < [5 3 1 4 2]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(4), float64(3), float64(2), float64(1)}, nil}), cljs_core.Sort.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(3), float64(1), float64(4), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [5 4 3 2 1] (sort > [5 3 1 4 2]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), "foo"}, nil}), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.Count, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "a", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"a\" [1 2] \"foo\"] (sort-by count [\"foo\" \"a\" [1 2]]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), "a"}, nil}), cljs_core.Sort_by.X_invoke_Arity3(cljs_core.Count, cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "a", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"foo\" [1 2] \"a\"] (sort-by count > [\"foo\" \"a\" [1 2]]))")}))
			}
			{
				var coll_5042 = (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})
				var shuffles_5043 = cljs_core.Filter.X_invoke_Arity2(func(G__5044 *cljs_core.AFn, coll_5042 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__5044, 1, func(p1__4078_SHARP_ interface{}) interface{} {
						return cljs_core.Not_EQ_.Arity2IIB(coll_5042, p1__4078_SHARP_)
					})
				}(&cljs_core.AFn{}, coll_5042), cljs_core.Take.X_invoke_Arity2(float64(100), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Shuffle, coll_5042).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
				_, _ = coll_5042, shuffles_5043
				if !(cljs_core.Empty_QMARK_.Arity1IB(shuffles_5043)) {
				} else {
					panic((&js.Error{("Assert failed: (not (empty? shuffles))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Last.X_invoke_Arity1(nil)) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (last nil))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Last.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 3 (last [1 2 3]))")}))
			}
			{
				var s_5045 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = s_5045
				{
					var n__1060__auto___5046 = float64(5)
					_ = n__1060__auto___5046
					{
						var n_5047 = float64(0)
						_ = n_5047
						for {
							if n_5047 < n__1060__auto___5046 {
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_5045, cljs_core.Conj, n_5047)
								n_5047 = (n_5047 + float64(1))
								continue
							} else {
							}
							break
						}
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Deref.X_invoke_Arity1(s_5045)) {
				} else {
					panic((&js.Error{("Assert failed: (= [0 1 2 3 4] (clojure.core/deref s))")}))
				}
			}
			{
				var v_5048 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})
				var s_5049 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)).(*cljs_core.CljsCoreAtom)
				_, _ = v_5048, s_5049
				{
					var seq__4624_5050 interface{} = cljs_core.Seq.Arity1IQ(v_5048)
					var chunk__4625_5051 interface{} = nil
					var count__4626_5052 = float64(0)
					var i__4627_5053 = float64(0)
					_, _, _, _ = seq__4624_5050, chunk__4625_5051, count__4626_5052, i__4627_5053
					for {
						if i__4627_5053 < count__4626_5052 {
							{
								var n_5054 = chunk__4625_5051.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4627_5053)
								_ = n_5054
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_5049, cljs_core.Conj, n_5054)
								seq__4624_5050, chunk__4625_5051, count__4626_5052, i__4627_5053 = seq__4624_5050, chunk__4625_5051, count__4626_5052, (i__4627_5053 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___5055 = cljs_core.Seq.Arity1IQ(seq__4624_5050)
								_ = temp__4222__auto___5055
								if cljs_core.Truth_(temp__4222__auto___5055) {
									{
										var seq__4624_5056___1 = temp__4222__auto___5055
										_ = seq__4624_5056___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4624_5056___1) {
											{
												var c__960__auto___5057 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4624_5056___1)
												_ = c__960__auto___5057
												seq__4624_5050, chunk__4625_5051, count__4626_5052, i__4627_5053 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4624_5056___1), c__960__auto___5057, cljs_core.Count.X_invoke_Arity1(c__960__auto___5057).(float64), float64(0)
												continue
											}
										} else {
											{
												var n_5058 = cljs_core.First.X_invoke_Arity1(seq__4624_5056___1)
												_ = n_5058
												cljs_core.Swap_BANG_.X_invoke_Arity3(s_5049, cljs_core.Conj, n_5058)
												seq__4624_5050, chunk__4625_5051, count__4626_5052, i__4627_5053 = cljs_core.Next.Arity1IQ(seq__4624_5056___1), nil, float64(0), float64(0)
												continue
											}
										}
									}
								} else {
								}
							}
						}
						break
					}
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Deref.X_invoke_Arity1(s_5049), cljs_core.Reverse.X_invoke_Arity1(v_5048)) {
				} else {
					panic((&js.Error{("Assert failed: (= (clojure.core/deref s) (reverse v))")}))
				}
			}
			{
				var a_5059 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				var d_5060 = (&cljs_core.CljsCoreDelay{func(G__5061 *cljs_core.AFn, a_5059 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__5061, 0, func() interface{} {
						return cljs_core.Swap_BANG_.X_invoke_Arity2(a_5059, cljs_core.Inc)
					})
				}(&cljs_core.AFn{}, a_5059), nil})
				_, _ = a_5059, d_5060
				if cljs_core.Realized_QMARK_.Arity1IB(d_5060) == false {
				} else {
					panic((&js.Error{("Assert failed: (false? (realized? d))")}))
				}
				if cljs_core.Deref.X_invoke_Arity1(a_5059).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (clojure.core/deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_5060)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_5059)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref a))")}))
				}
				if cljs_core.Realized_QMARK_.Arity1IB(d_5060) == true {
				} else {
					panic((&js.Error{("Assert failed: (true? (realized? d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_5060)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_5059)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Force.X_invoke_Arity1(d_5060), cljs_core.Deref.X_invoke_Arity1(d_5060)) {
				} else {
					panic((&js.Error{("Assert failed: (= (force d) (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Force.X_invoke_Arity1(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (force 1))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Assoc.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentArrayMap_EMPTY, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(4)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= {1 2, 3 4} (assoc {} 1 2 3 4))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), cljs_core.Assoc.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, float64(1), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= {1 2} (assoc {} 1 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42), float64(2)}, nil}), cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(0), float64(42))) {
			} else {
				panic((&js.Error{("Assert failed: (= [42 2] (assoc [1 2] 0 42))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Dissoc.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= {} (dissoc {1 2, 3 4} 1 3))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), cljs_core.Dissoc.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), float64(3))) {
			} else {
				panic((&js.Error{("Assert failed: (= {1 2} (dissoc {1 2, 3 4} 3))")}))
			}
			if cljs_core.Nil_(cljs_core.Dissoc.X_invoke_Arity2(nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (dissoc nil :foo))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), cljs_core.Disj.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 3 2} (disj #{1 3 2}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), cljs_core.Disj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(3))) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 2} (disj #{1 3 2} 3))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), cljs_core.Disj.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1} (disj #{1 3 2} 2 3))")}))
			}
			if cljs_core.Nil_(cljs_core.Disj.X_invoke_Arity2(nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (disj nil :foo))")}))
			}
			{
				var f_5062 = cljs_core.Memoize.X_invoke_Arity1(func(G__5063 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5063, 0, func() interface{} {
						return cljs_core.Rand.Arity0F()
					})
				}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn)
				_ = f_5062
				{
					f_5062.X_invoke_Arity0()
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return f_5062.X_invoke_Arity0()
				}(), func() interface{} {
					return f_5062.X_invoke_Arity0()
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= (f) (f))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})), nil) {
			} else {
				panic((&js.Error{("Assert failed: (= (find {} :a) nil)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (find {:a 1} :a) [:a 1])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})), nil) {
			} else {
				panic((&js.Error{("Assert failed: (= (find {:a 1} :b) nil)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (find {:a 1, :b 2} :a) [:a 1])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (find {:a 1, :b 2} :b) [:b 2])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})), nil) {
			} else {
				panic((&js.Error{("Assert failed: (= (find {:a 1, :b 2} :c) nil)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, nil), nil) {
			} else {
				panic((&js.Error{("Assert failed: (= (find {} nil) nil)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), nil), nil) {
			} else {
				panic((&js.Error{("Assert failed: (= (find {:a 1} nil) nil)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), nil), nil) {
			} else {
				panic((&js.Error{("Assert failed: (= (find {:a 1, :b 2} nil) nil)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(0)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (find [1 2 3] 0) [0 1])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(4), float64(2)).(float64), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 4 2) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(3), float64(2)).(float64), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 3 2) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(6), float64(4)).(float64), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 6 4) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(0), float64(5)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 0 5) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(42), float64(5)).(float64), float64(8)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 42 5) 8)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(42), float64(-5)).(float64), float64(-8)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 42 -5) -8)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(-42), float64(-5)).(float64), float64(8)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot -42 -5) 8)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(9), float64(3)).(float64), float64(3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 9 3) 3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(9), float64(-3)).(float64), float64(-3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 9 -3) -3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(-9), float64(3)).(float64), float64(-3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot -9 3) -3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(2), float64(-5)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 2 -5) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(-2), float64(5)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot -2 5) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(0), float64(3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 0 3) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(0), float64(-3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quot 0 -3) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(4), float64(2)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 4 2) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(3), float64(2)).(float64), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 3 2) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(6), float64(4)).(float64), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 6 4) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(0), float64(5)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 0 5) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(4.5, 2.0).(float64), 0.5) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 4.5 2.0) 0.5)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(42), float64(5)).(float64), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 42 5) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(9), float64(3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 9 3) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(9), float64(-3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 9 -3) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(-9), float64(3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod -9 3) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(-9), float64(-3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod -9 -3) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(0), float64(3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 0 3) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(3216478362187432), float64(432143214)).(float64), float64(120355456)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mod 3216478362187432 432143214) 120355456)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(4), float64(2)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (rem 4 2) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(0), float64(5)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (rem 0 5) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(4.5, 2.0).(float64), 0.5) {
			} else {
				panic((&js.Error{("Assert failed: (= (rem 4.5 2.0) 0.5)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(42), float64(5)).(float64), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (rem 42 5) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(2), float64(5)).(float64), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (rem 2 5) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(2), float64(-5)).(float64), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (rem 2 -5) 2)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(0), float64(3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (rem 0 3) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(9)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(6)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(5)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(4)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 10) (list 0 1 2 3 4 5 6 7 8 9))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity2(float64(10), float64(20)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(19)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(18)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(17)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(16)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(15)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(14)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(13)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(12)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(11)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(10))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 10 20) (list 10 11 12 13 14 15 16 17 18 19))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(10), float64(20), float64(2)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(18)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(16)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(14)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(12)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(10))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 10 20 2) (list 10 12 14 16 18))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Take.X_invoke_Arity2(float64(20), cljs_core.Range_.X_invoke_Arity0().(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(19)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(18)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(17)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(16)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(15)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(14)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(13)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(12)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(11)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(10)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(9)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(6)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(5)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(4)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (take 20 (range)) (list 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19))")}))
			}
			{
				var d_5064 = cljs_core.Group_by.X_invoke_Arity2(cljs_core.Second, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(6), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), float64(4), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), float64(2)}, nil}))
				_ = d_5064
				if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_5064, float64(1))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (count (get d 1)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_5064, float64(2))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (count (get d 2)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_5064, float64(4))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (count (get d 4)))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Merge.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(3), float64(4)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(5), float64(6)}, nil})}))) {
			} else {
				panic((&js.Error{("Assert failed: (= {1 2, 3 4, 5 6} (merge {1 2} {3 4} {5 6}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Merge.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(3), float64(4)}, nil}), nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= {1 2, 3 4} (merge {1 2} {3 4} nil))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), cljs_core.Frequencies.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:a 3, :b 2} (frequencies [:a :b :a :b :a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(6), float64(10), float64(15)}, nil}), cljs_core.Reductions.X_invoke_Arity2(cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 3 6 10 15] (reductions + [1 2 3 4 5]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__5065 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5065, 1, func(p1__4079_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__4079_SHARP_) {
						return p1__4079_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 3 5 7 9] (keep (fn* [p1__4079#] (if (odd? p1__4079#) p1__4079#)) [1 2 3 4 5 6 7 8 9 10]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(6), float64(8), float64(10)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__5066 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5066, 1, func(p1__4080_SHARP_ interface{}) interface{} {
					if cljs_core.Even_QMARK_.Arity1IB(p1__4080_SHARP_) {
						return p1__4080_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 4 6 8 10] (keep (fn* [p1__4080#] (if (even? p1__4080#) p1__4080#)) [1 2 3 4 5 6 7 8 9 10]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__5067 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5067, 2, func(p1__4081_SHARP_ interface{}, p2__4082_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__4081_SHARP_) {
						return p2__4082_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(11), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 3 5 7 9] (keep-indexed (fn* [p1__4081# p2__4082#] (if (odd? p1__4081#) p2__4082#)) [0 1 2 3 4 5 6 7 8 9 10]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(5)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__5068 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5068, 2, func(p1__4084_SHARP_ interface{}, p2__4083_SHARP_ interface{}) interface{} {
					if p2__4083_SHARP_.(float64) > float64(0) {
						return p1__4084_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-9), float64(0), float64(29), float64(-7), float64(45), float64(3), float64(-8)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 4 5] (keep-indexed (fn* [p1__4084# p2__4083#] (if (pos? p2__4083#) p1__4084#)) [-9 0 29 -7 45 3 -8]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil}), cljs_core.Map_indexed.X_invoke_Arity2(func(G__5069 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5069, 2, func(p1__4085_SHARP_ interface{}, p2__4086_SHARP_ interface{}) interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__4085_SHARP_, p2__4086_SHARP_}, nil})
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [[0 :a] [1 :b] [2 :c]] (map-indexed (fn* [p1__4085# p2__4086#] (vector p1__4085# p2__4086#)) [:a :b :c]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{"Foo", cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"foo", "FOO", "fOo"})).(*cljs_core.CljsCoreList), "Bar", cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"bar", "BAR", "BAr"})).(*cljs_core.CljsCoreList), "Baz", (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"baz"}, nil}), "Qux", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"qux", "quux"}, nil})}, nil}), cljs_core.Merge_with.X_invoke_ArityVariadic(cljs_core.Concat, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{"Foo", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "FOO"}, nil}), "Bar", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"bar", "BAR"}, nil}), "Baz", (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"baz"}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{"Foo", (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"fOo"}, nil}), "Bar", (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"BAr"}, nil}), "Qux", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"qux", "quux"}, nil})}, nil})}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote {\"Foo\" (\"foo\" \"FOO\" \"fOo\"), \"Bar\" (\"bar\" \"BAR\" \"BAr\"), \"Baz\" [\"baz\"], \"Qux\" [\"qux\" \"quux\"]}) (merge-with concat {\"Foo\" [\"foo\" \"FOO\"], \"Bar\" [\"bar\" \"BAR\"], \"Baz\" [\"baz\"]} {\"Foo\" [\"fOo\"], \"Bar\" [\"BAr\"], \"Qux\" [\"qux\" \"quux\"]}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(111), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(102), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(13)}, nil}), cljs_core.Merge_with.X_invoke_ArityVariadic(cljs_core.X_PLUS_, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(3)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(10), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(10)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(100), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(100)}, nil})}))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:a 111, :b 102, :c 13} (merge-with + {:a 1, :b 2, :c 3} {:a 10, :c 10} {:a 100, :b 100}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(102), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(13)}, nil}), cljs_core.Apply.X_invoke_Arity2(cljs_core.Merge_with, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(100)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(3)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(10)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:a 3, :b 102, :c 13} (apply merge-with [+ {:a 1, :b 100} {:a 1, :b 2, :c 3} {:a 1, :c 10}]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "e", Str: "e", X_hash: float64(-1273166571), X_meta: nil})}, nil}), cljs_core.Replace.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "d", Str: "d", X_hash: float64(-682293345), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "e", Str: "e", X_hash: float64(-1273166571), X_meta: nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2), float64(4)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote [a c e]) (replace (quote [a b c d e]) [0 2 4]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "zero", Fqn: "zero", X_hash: float64(-858964576)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "two", Fqn: "two", X_hash: float64(627606869)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "zero", Fqn: "zero", X_hash: float64(-858964576)})}, nil}), cljs_core.Replace.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "zero", Fqn: "zero", X_hash: float64(-858964576)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "two", Fqn: "two", X_hash: float64(627606869)})}, nil}), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(0), float64(2), float64(0)})).(*cljs_core.CljsCoreList))) {
			} else {
				panic((&js.Error{("Assert failed: (= [:one :zero :two :zero] (replace {0 :zero, 1 :one, 2 :two} (quote (1 0 2 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil})}, nil}), cljs_core.Split_at.X_invoke_Arity2(float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})).(cljs_core.CljsCoreIVector)) {
			} else {
				panic((&js.Error{("Assert failed: (= [[1 2] [3 4 5]] (split-at 2 [1 2 3 4 5]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil})}, nil}), cljs_core.Split_with.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.X_GT__EQ_, float64(3)).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})).(cljs_core.CljsCoreIVector)) {
			} else {
				panic((&js.Error{("Assert failed: (= [[1 2 3] [4 5]] (split-with (partial >= 3) [1 2 3 4 5]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(10000), cljs_core.Trampoline.X_invoke_ArityVariadic(func(f *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(f, 1, func(n interface{}) interface{} {
					if n.(float64) >= float64(10000) {
						return n
					} else {
						return func(G__5070 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__5070, 0, func() interface{} {
								return f.X_invoke_Arity1((n.(float64) + float64(1)))
							})
						}(&cljs_core.AFn{})
					}
				})
			}(&cljs_core.AFn{}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 10000 (trampoline (fn f [n] (if (>= n 10000) n (fn* [] (f (inc n))))) 0))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), cljs_core.Meta.X_invoke_Arity1(cljs_core.Vary_meta.X_invoke_Arity4(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Assoc, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:a 1} (meta (vary-meta [] assoc :a 1)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), cljs_core.Meta.X_invoke_Arity1(cljs_core.Vary_meta.X_invoke_Arity4(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})), cljs_core.Assoc, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:a 1, :b 2} (meta (vary-meta (with-meta [] {:b 2}) assoc :a 1)))")}))
			}
			cljs_core.Derive.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}))
			cljs_core.Derive.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "square", Fqn: "cljs.core-test/square", X_hash: float64(-1438469543)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), nil}, nil}, nil}), cljs_core.Parents.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= #{:cljs.core-test/shape} (parents :cljs.core-test/rect))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), nil, (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), nil}, nil}, nil}), cljs_core.Ancestors.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "square", Fqn: "cljs.core-test/square", X_hash: float64(-1438469543)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= #{:cljs.core-test/shape :cljs.core-test/rect} (ancestors :cljs.core-test/square))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), nil, (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "square", Fqn: "cljs.core-test/square", X_hash: float64(-1438469543)}), nil}, nil}, nil}), cljs_core.Descendants.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= #{:cljs.core-test/rect :cljs.core-test/square} (descendants :cljs.core-test/shape))")}))
			}
			if cljs_core.Isa_QMARK_.Arity2IIB(float64(42), float64(42)) == true {
			} else {
				panic((&js.Error{("Assert failed: (true? (isa? 42 42))")}))
			}
			if cljs_core.Isa_QMARK_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "square", Fqn: "cljs.core-test/square", X_hash: float64(-1438469543)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})) == true {
			} else {
				panic((&js.Error{("Assert failed: (true? (isa? :cljs.core-test/square :cljs.core-test/shape))")}))
			}
			cljs_core.Derive.X_invoke_Arity2(reflect.TypeOf((**cljs_core.CljsCorePersistentHashSet)(nil)).Elem(), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "collection", Fqn: "cljs.core-test/collection", X_hash: float64(-440981704)}))
			if cljs_core.Isa_QMARK_.Arity2IIB(reflect.TypeOf((**cljs_core.CljsCorePersistentHashSet)(nil)).Elem(), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "collection", Fqn: "cljs.core-test/collection", X_hash: float64(-440981704)})) == true {
			} else {
				panic((&js.Error{("Assert failed: (true? (isa? cljs.core/PersistentHashSet :cljs.core-test/collection))")}))
			}
			if cljs_core.Isa_QMARK_.Arity2IIB(reflect.TypeOf((**cljs_core.CljsCoreIndexedSeq)(nil)).Elem(), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "collection", Fqn: "cljs.core-test/collection", X_hash: float64(-440981704)})) == false {
			} else {
				panic((&js.Error{("Assert failed: (false? (isa? cljs.core/IndexedSeq :cljs.core-test/collection))")}))
			}
			if cljs_core.Isa_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "square", Fqn: "cljs.core-test/square", X_hash: float64(-1438469543)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil})) == true {
			} else {
				panic((&js.Error{("Assert failed: (true? (isa? [:cljs.core-test/square :cljs.core-test/rect] [:cljs.core-test/shape :cljs.core-test/shape]))")}))
			}
			Bar = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1070__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1071__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1072__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1073__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1074__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__
				return (&cljs_core.CljsCoreMultiFn{"bar", func(G__5071 *cljs_core.AFn, method_table__1070__auto__ *cljs_core.CljsCoreAtom, prefer_table__1071__auto__ *cljs_core.CljsCoreAtom, method_cache__1072__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1073__auto__ *cljs_core.CljsCoreAtom, hierarchy__1074__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5071, 2, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{x, y}, nil})
					})
				}(&cljs_core.AFn{}, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1074__auto__, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__})
			}()

			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}), func(G__5072 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5072, 2, func(x interface{}, y interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)})
				})
			}(&cljs_core.AFn{}))
			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})}, nil}), func(G__5073 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5073, 2, func(x interface{}, y interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "shape-rect", Fqn: "shape-rect", X_hash: float64(-613148403)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.Count.X_invoke_Arity1(cljs_core.Prefers.X_invoke_Arity1(Bar)).(float64) == float64(0) {
			} else {
				panic((&js.Error{("Assert failed: (zero? (count (prefers bar)))")}))
			}
			cljs_core.Prefer_method.X_invoke_Arity3(Bar, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})}, nil}))
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Count.X_invoke_Arity1(cljs_core.Prefers.X_invoke_Arity1(Bar)).(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (count (prefers bar)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)}), func() interface{} {
				var G__4629 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				var G__4630 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				_, _ = G__4629, G__4630
				return Bar.X_invoke_Arity2(G__4629, G__4630)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :rect-shape (bar :cljs.core-test/rect :cljs.core-test/rect))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)}), cljs_core.Apply.X_invoke_Arity2(Bar.X_get_method_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= :rect-shape (apply (-get-method bar [:cljs.core-test/rect :cljs.core-test/shape]) [:cljs.core-test/rect :cljs.core-test/shape]))")}))
			}
			Nested_dispatch = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1070__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1071__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1072__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1073__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1074__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__
				return (&cljs_core.CljsCoreMultiFn{"nested-dispatch", func(G__5074 *cljs_core.AFn, method_table__1070__auto__ *cljs_core.CljsCoreAtom, prefer_table__1071__auto__ *cljs_core.CljsCoreAtom, method_cache__1072__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1073__auto__ *cljs_core.CljsCoreAtom, hierarchy__1074__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5074, 1, func(m interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(m))
					})
				}(&cljs_core.AFn{}, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1074__auto__, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__})
			}()

			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), func(G__5075 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5075, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__5076 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5076, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__4631 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil})
				_ = G__4631
				return Nested_dispatch.X_invoke_Arity1(G__4631)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :nested-a (nested-dispatch {:a {:b :c}}))")}))
			}
			Nested_dispatch2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1070__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1071__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1072__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1073__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1074__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__
				return (&cljs_core.CljsCoreMultiFn{"nested-dispatch2", cljs_core.Ffirst, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1074__auto__, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__})
			}()

			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__5077 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5077, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__5078 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5078, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__4632 = (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})}, nil})
				_ = G__4632
				return Nested_dispatch2.X_invoke_Arity1(G__4632)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :nested-a (nested-dispatch2 [[:a :b]]))")}))
			}
			Foo1 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1070__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1071__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1072__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1073__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1074__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__
				return (&cljs_core.CljsCoreMultiFn{"foo1", func(G__5079 *cljs_core.AFn, method_table__1070__auto__ *cljs_core.CljsCoreAtom, prefer_table__1071__auto__ *cljs_core.CljsCoreAtom, method_cache__1072__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1073__auto__ *cljs_core.CljsCoreAtom, hierarchy__1074__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5079, 0, func(args__ ...interface{}) interface{} {
						var args = cljs_core.Seq.Arity1IQ(args__[0])
						_ = args
						return cljs_core.First.X_invoke_Arity1(args)
					})
				}(&cljs_core.AFn{}, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1074__auto__, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__})
			}()

			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__5080 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5080, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)})
				})
			}(&cljs_core.AFn{}))
			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__5081 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5081, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)}), func() interface{} {
				var G__4633 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
				_ = G__4633
				return Foo1.X_invoke_Arity1(G__4633)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :a-return (foo1 :a))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)}), func() interface{} {
				var G__4634 = float64(1)
				_ = G__4634
				return Foo1.X_invoke_Arity1(G__4634)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :default-return (foo1 1))")}))
			}
			Area = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1070__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1071__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1072__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1073__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1074__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__
				return (&cljs_core.CljsCoreMultiFn{"area", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Shape", Fqn: "Shape", X_hash: float64(-248980586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1074__auto__, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__})
			}()

			Rect = func(rect *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(rect, 2, func(wd interface{}, ht interface{}) interface{} {
					return (&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Shape", Fqn: "Shape", X_hash: float64(-248980586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Rect", Fqn: "Rect", X_hash: float64(-420704620)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wd", Fqn: "wd", X_hash: float64(-183204751)}), wd, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ht", Fqn: "ht", X_hash: float64(214115472)}), ht}, nil})
				})
			}(&cljs_core.AFn{})

			Circle = func(circle *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(circle, 1, func(radius interface{}) interface{} {
					return (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Shape", Fqn: "Shape", X_hash: float64(-248980586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Circle", Fqn: "Circle", X_hash: float64(-158896930)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}), radius}, nil})
				})
			}(&cljs_core.AFn{})

			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Rect", Fqn: "Rect", X_hash: float64(-420704620)}), func(G__5082 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5082, 1, func(r interface{}) interface{} {
					return ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wd", Fqn: "wd", X_hash: float64(-183204751)}).X_invoke_Arity1(r).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ht", Fqn: "ht", X_hash: float64(214115472)}).X_invoke_Arity1(r).(float64))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Circle", Fqn: "Circle", X_hash: float64(-158896930)}), func(G__5083 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5083, 1, func(c interface{}) interface{} {
					return (Math.PI * ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64)))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__5084 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5084, 1, func(x interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)})
				})
			}(&cljs_core.AFn{}))
			R = Rect.X_invoke_Arity2(float64(4), float64(13)).(cljs_core.CljsCoreIMap)

			C = Circle.X_invoke_Arity1(float64(12)).(cljs_core.CljsCoreIMap)

			if cljs_core.X_EQ_.Arity2IIB(float64(52), func() interface{} {
				var G__4635 = R
				_ = G__4635
				return Area.X_invoke_Arity1(G__4635)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 52 (area r))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)}), func() interface{} {
				var G__4636 = cljs_core.CljsCorePersistentArrayMap_EMPTY
				_ = G__4636
				return Area.X_invoke_Arity1(G__4636)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :oops (area {}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Count.X_invoke_Arity1(cljs_core.Methods.X_invoke_Arity1(Bar)).(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (count (methods bar)))")}))
			}
			cljs_core.Remove_method.X_invoke_Arity2(Bar, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}))
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Count.X_invoke_Arity1(cljs_core.Methods.X_invoke_Arity1(Bar)).(float64)) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (count (methods bar)))")}))
			}
			cljs_core.Remove_all_methods.X_invoke_Arity1(Bar)
			if cljs_core.Count.X_invoke_Arity1(cljs_core.Methods.X_invoke_Arity1(Bar)).(float64) == float64(0) {
			} else {
				panic((&js.Error{("Assert failed: (zero? (count (methods bar)))")}))
			}
			Apply_multi_test = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1070__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1071__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1072__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1073__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1074__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__
				return (&cljs_core.CljsCoreMultiFn{"apply-multi-test", func(G__5085 *cljs_core.AFn, method_table__1070__auto__ *cljs_core.CljsCoreAtom, prefer_table__1071__auto__ *cljs_core.CljsCoreAtom, method_cache__1072__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1073__auto__ *cljs_core.CljsCoreAtom, hierarchy__1074__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5085, 3, func(___ interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
						return float64(0)
					})
				}(&cljs_core.AFn{}, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1074__auto__, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__})
			}()

			Apply_multi_test.X_add_method_Arity3(float64(0), func(G__5086 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5086, 2, func(x interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})
				}, func(x interface{}, y interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "two", Fqn: "two", X_hash: float64(627606869)})
				}, func(x_y_r__ ...interface{}) interface{} {
					var x = x_y_r__[0]
					var y = x_y_r__[1]
					var r = cljs_core.Seq.Arity1IQ(x_y_r__[2])
					_, _, _ = x, y, r
					return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "three", Fqn: "three", X_hash: float64(-1651831795)}), r}, nil})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "three", Fqn: "three", X_hash: float64(-1651831795)}), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)})).(*cljs_core.CljsCoreList)}, nil}), cljs_core.Apply.X_invoke_Arity2(Apply_multi_test, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [:three (quote (2))] (apply apply-multi-test [0 1 2]))")}))
			}
			My_map_hierarchy = cljs_core.Atom.X_invoke_Arity1(cljs_core.Derive.X_invoke_Arity3(cljs_core.Derive.X_invoke_Arity3(cljs_core.Derive.X_invoke_Arity3(cljs_core.Make_hierarchy.X_invoke_Arity0().(cljs_core.CljsCoreIMap), cljs_core.Type_.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)})), cljs_core.Type_.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_EMPTY), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)})), cljs_core.Type_.X_invoke_Arity1(cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)}))).(*cljs_core.CljsCoreAtom)

			My_map_QMARK_ = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1070__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1071__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1072__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1073__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1074__auto__ = cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), My_map_hierarchy}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__
				return (&cljs_core.CljsCoreMultiFn{"my-map?", cljs_core.Type_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1074__auto__, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__})
			}()

			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)}), func(G__5087 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5087, 1, func(___ interface{}) interface{} {
					return true
				})
			}(&cljs_core.AFn{}))
			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__5088 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5088, 1, func(___ interface{}) interface{} {
					return false
				})
			}(&cljs_core.AFn{}))
			{
				var seq__4637_5089 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__4638_5090 interface{} = nil
				var count__4639_5091 = float64(0)
				var i__4640_5092 = float64(0)
				_, _, _, _ = seq__4637_5089, chunk__4638_5090, count__4639_5091, i__4640_5092
				for {
					if i__4640_5092 < count__4639_5091 {
						{
							var m_5093 = chunk__4638_5090.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4640_5092)
							_ = m_5093
							if cljs_core.Truth_(func() interface{} {
								var G__4641 = m_5093
								_ = G__4641
								return My_map_QMARK_.X_invoke_Arity1(G__4641)
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (my-map? m)")}))
							}
							seq__4637_5089, chunk__4638_5090, count__4639_5091, i__4640_5092 = seq__4637_5089, chunk__4638_5090, count__4639_5091, (i__4640_5092 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5094 = cljs_core.Seq.Arity1IQ(seq__4637_5089)
							_ = temp__4222__auto___5094
							if cljs_core.Truth_(temp__4222__auto___5094) {
								{
									var seq__4637_5095___1 = temp__4222__auto___5094
									_ = seq__4637_5095___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4637_5095___1) {
										{
											var c__960__auto___5096 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4637_5095___1)
											_ = c__960__auto___5096
											seq__4637_5089, chunk__4638_5090, count__4639_5091, i__4640_5092 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4637_5095___1), c__960__auto___5096, cljs_core.Count.X_invoke_Arity1(c__960__auto___5096).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_5097 = cljs_core.First.X_invoke_Arity1(seq__4637_5095___1)
											_ = m_5097
											if cljs_core.Truth_(func() interface{} {
												var G__4642 = m_5097
												_ = G__4642
												return My_map_QMARK_.X_invoke_Arity1(G__4642)
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (my-map? m)")}))
											}
											seq__4637_5089, chunk__4638_5090, count__4639_5091, i__4640_5092 = cljs_core.Next.Arity1IQ(seq__4637_5095___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
							}
						}
					}
					break
				}
			}
			{
				var seq__4643_5098 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentVector_EMPTY, float64(1), "asdf", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}))
				var chunk__4644_5099 interface{} = nil
				var count__4645_5100 = float64(0)
				var i__4646_5101 = float64(0)
				_, _, _, _ = seq__4643_5098, chunk__4644_5099, count__4645_5100, i__4646_5101
				for {
					if i__4646_5101 < count__4645_5100 {
						{
							var not_m_5102 = chunk__4644_5099.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4646_5101)
							_ = not_m_5102
							if cljs_core.Not.Arity1IB(func() interface{} {
								var G__4647 = not_m_5102
								_ = G__4647
								return My_map_QMARK_.X_invoke_Arity1(G__4647)
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (not (my-map? not-m))")}))
							}
							seq__4643_5098, chunk__4644_5099, count__4645_5100, i__4646_5101 = seq__4643_5098, chunk__4644_5099, count__4645_5100, (i__4646_5101 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5103 = cljs_core.Seq.Arity1IQ(seq__4643_5098)
							_ = temp__4222__auto___5103
							if cljs_core.Truth_(temp__4222__auto___5103) {
								{
									var seq__4643_5104___1 = temp__4222__auto___5103
									_ = seq__4643_5104___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4643_5104___1) {
										{
											var c__960__auto___5105 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4643_5104___1)
											_ = c__960__auto___5105
											seq__4643_5098, chunk__4644_5099, count__4645_5100, i__4646_5101 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4643_5104___1), c__960__auto___5105, cljs_core.Count.X_invoke_Arity1(c__960__auto___5105).(float64), float64(0)
											continue
										}
									} else {
										{
											var not_m_5106 = cljs_core.First.X_invoke_Arity1(seq__4643_5104___1)
											_ = not_m_5106
											if cljs_core.Not.Arity1IB(func() interface{} {
												var G__4648 = not_m_5106
												_ = G__4648
												return My_map_QMARK_.X_invoke_Arity1(G__4648)
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (not (my-map? not-m))")}))
											}
											seq__4643_5098, chunk__4644_5099, count__4645_5100, i__4646_5101 = cljs_core.Next.Arity1IQ(seq__4643_5104___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
							}
						}
					}
					break
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(10), float64(3)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(9)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(6)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 0 10 3) (list 0 3 6 9))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(10), float64(3)).(*cljs_core.CljsCoreRange)).(float64), float64(4)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count (range 0 10 3)) 4)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(-10), float64(-3)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(-9)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-6)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 0 -10 -3) (list 0 -3 -6 -9))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(-10), float64(-3)).(*cljs_core.CljsCoreRange)).(float64), float64(4)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count (range 0 -10 -3)) 4)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(-10), float64(10), float64(3)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(5)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-4)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-10))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range -10 10 3) (list -10 -7 -4 -1 2 5 8))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(-10), float64(10), float64(3)).(*cljs_core.CljsCoreRange)).(float64), float64(7)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count (range -10 10 3)) 7)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(1), float64(1)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 0 1 1) (list 0))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(-3), float64(-1)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(-2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 0 -3 -1) (list 0 -1 -2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(3), float64(0), float64(-1)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 3 0 -1) (list 3 2 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(10), float64(-1)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 0 10 -1) (list))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(1), float64(0)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 0 1 0) (list))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(10), float64(0), float64(1)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 10 0 1) (list))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(0), float64(0)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 0 0 0) (list))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(10), float64(-1)).(*cljs_core.CljsCoreRange)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count (range 0 10 -1)) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(1), float64(0)).(*cljs_core.CljsCoreRange)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count (range 0 1 0)) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(10), float64(0), float64(1)).(*cljs_core.CljsCoreRange)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count (range 10 0 1)) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(0), float64(0)).(*cljs_core.CljsCoreRange)).(float64), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count (range 0 0 0)) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.Range_.X_invoke_Arity3(float64(1), float64(0), float64(0)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= (take 3 (range 1 0 0)) (list 1 1 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.Range_.X_invoke_Arity3(float64(3), float64(1), float64(0)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
			} else {
				panic((&js.Error{("Assert failed: (= (take 3 (range 3 1 0)) (list 3 3 3))")}))
			}
			{
				var pv_5107 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange))
				_ = pv_5107
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity2(pv_5107, float64(96)), float64(96)) {
				} else {
					panic((&js.Error{("Assert failed: (= (nth pv 96) 96)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity3(pv_5107, float64(97), nil), nil) {
				} else {
					panic((&js.Error{("Assert failed: (= (nth pv 97 nil) nil)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4649 = float64(96)
					_ = G__4649
					return pv_5107.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4649)
				}(), float64(96)) {
				} else {
					panic((&js.Error{("Assert failed: (= (pv 96) 96)")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(cljs_core.CljsCorePersistentVector_EMPTY)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq []))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reverse.X_invoke_Arity1(pv_5107), cljs_core.Rseq.Arity1IQ(pv_5107)) {
				} else {
					panic((&js.Error{("Assert failed: (= (reverse pv) (rseq pv))")}))
				}
			}
			{
				var pv_5108 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(33)).(*cljs_core.CljsCoreRange))
				_ = pv_5108
				if cljs_core.X_EQ_.Arity2IIB(pv_5108, cljs_core.Conj.X_invoke_Arity2(cljs_core.Conj.X_invoke_Arity2(cljs_core.Pop.X_invoke_Arity1(cljs_core.Pop.X_invoke_Arity1(pv_5108)), float64(31)), float64(32))) {
				} else {
					panic((&js.Error{("Assert failed: (= pv (-> pv pop pop (conj 31) (conj 32)))")}))
				}
			}
			{
				var stack1_5109 = cljs_core.Pop.X_invoke_Arity1(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange)))
				var stack2_5110 = cljs_core.Pop.X_invoke_Arity1(stack1_5109)
				_, _ = stack1_5109, stack2_5110
				if cljs_core.X_EQ_.Arity2IIB(float64(95), cljs_core.Peek.X_invoke_Arity1(stack1_5109)) {
				} else {
					panic((&js.Error{("Assert failed: (= 95 (peek stack1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(94), cljs_core.Peek.X_invoke_Arity1(stack2_5110)) {
				} else {
					panic((&js.Error{("Assert failed: (= 94 (peek stack2))")}))
				}
			}
			{
				var sentinel_5111 = cljs_core.Rand.Arity0F()
				_ = sentinel_5111
				if reflect.DeepEqual(sentinel_5111, func() (return__5112 interface{}) {
					defer func() {
						if e4650 := recover(); e4650 != nil {
							if cljs_core.Value_(e4650).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var ___ = e4650
									_ = ___
									return__5112 = sentinel_5111
								}
							} else {
								panic(e4650)

							}
						}
					}()
					{
						return cljs_core.CljsCorePersistentVector_EMPTY.X_invoke_Arity1(float64(0))
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (identical? sentinel (try ([] 0) (catch js/Error _ sentinel)))")}))
				}
			}
			{
				var v1_5113 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				var v2_5114 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(5)).(*cljs_core.CljsCoreRange))
				var s_5115 = cljs_core.Subvec.X_invoke_Arity3(v1_5113, float64(2), float64(8)).(*cljs_core.CljsCoreSubvec)
				_, _, _ = v1_5113, v2_5114, s_5115
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(s_5115, cljs_core.Subvec.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2(v1_5113, float64(2)).(*cljs_core.CljsCoreSubvec), float64(0), float64(6)).(*cljs_core.CljsCoreSubvec), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Take.X_invoke_Arity2(float64(6), cljs_core.Drop.X_invoke_Arity2(float64(2), v1_5113).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= s (-> v1 (subvec 2) (subvec 0 6)) (->> v1 (drop 2) (take 6)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Count.X_invoke_Arity1(s_5115).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (count s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Pop.X_invoke_Arity1(s_5115)) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 4 5 6] (pop s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(7), cljs_core.Peek.X_invoke_Arity1(s_5115)) {
				} else {
					panic((&js.Error{("Assert failed: (= 7 (peek s))")}))
				}
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(1)}, nil}), cljs_core.Assoc.X_invoke_Arity3(s_5115, float64(6), float64(1)), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Conj.X_invoke_Arity2(s_5115, float64(1))}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 4 5 6 7 1] (assoc s 6 1) (conj s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(27), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, s_5115)) {
				} else {
					panic((&js.Error{("Assert failed: (= 27 (reduce + s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(s_5115, cljs_core.Vec.X_invoke_Arity1(s_5115)) {
				} else {
					panic((&js.Error{("Assert failed: (= s (vec s))")}))
				}
				{
					var m_5116 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), float64(1)}, nil})
					_ = m_5116
					if cljs_core.X_EQ_.Arity2IIB(m_5116, cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(s_5115, m_5116))) {
					} else {
						panic((&js.Error{("Assert failed: (= m (meta (with-meta s m)))")}))
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5117 interface{}) {
					defer func() {
						if e4651 := recover(); e4651 != nil {
							if cljs_core.Value_(e4651).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4651
									_ = e
									return__5117 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e4651)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_5114, float64(0), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 0 6) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5118 interface{}) {
					defer func() {
						if e4652 := recover(); e4652 != nil {
							if cljs_core.Value_(e4652).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4652
									_ = e
									return__5118 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e4652)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_5114, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5119 interface{}) {
					defer func() {
						if e4653 := recover(); e4653 != nil {
							if cljs_core.Value_(e4653).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4653
									_ = e
									return__5119 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e4653)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_5114, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5120 interface{}) {
					defer func() {
						if e4654 := recover(); e4654 != nil {
							if cljs_core.Value_(e4654).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4654
									_ = e
									return__5120 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e4654)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_5114, float64(3), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 3 6) (catch js/Error e :fail)))")}))
				}
				if reflect.DeepEqual(v1_5113, cljs_core.Subvec.X_invoke_Arity3(s_5115, float64(1), float64(4)).(*cljs_core.CljsCoreSubvec).V) {
				} else {
					panic((&js.Error{("Assert failed: (identical? v1 (.-v (subvec s 1 4)))")}))
				}
				{
					var sentinel_5121 = cljs_core.Rand.Arity0F()
					var s_5122___1 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2))
					_, _ = sentinel_5121, s_5122___1
					if reflect.DeepEqual(sentinel_5121, func() (return__5123 interface{}) {
						defer func() {
							if e4655 := recover(); e4655 != nil {
								if cljs_core.Value_(e4655).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var ___ = e4655
										_ = ___
										return__5123 = sentinel_5121
									}
								} else {
									panic(e4655)

								}
							}
						}()
						{
							{
								var G__4656 = float64(-1)
								_ = G__4656
								return s_5122___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4656)
							}
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (identical? sentinel (try (s -1) (catch js/Error _ sentinel)))")}))
					}
					if reflect.DeepEqual(sentinel_5121, func() (return__5124 interface{}) {
						defer func() {
							if e4657 := recover(); e4657 != nil {
								if cljs_core.Value_(e4657).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var ___ = e4657
										_ = ___
										return__5124 = sentinel_5121
									}
								} else {
									panic(e4657)

								}
							}
						}()
						{
							{
								var G__4658 = float64(1)
								_ = G__4658
								return s_5122___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4658)
							}
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (identical? sentinel (try (s 1) (catch js/Error _ sentinel)))")}))
					}
				}
				{
					var sv1_5125 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2)).(*cljs_core.CljsCoreSubvec)
					var sv2_5126 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(1)).(*cljs_core.CljsCoreSubvec)
					_, _ = sv1_5125, sv2_5126
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(sv1_5125), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(*cljs_core.CljsCoreList)) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq sv1) (quote (1)))")}))
					}
					if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(sv2_5126)) {
					} else {
						panic((&js.Error{("Assert failed: (nil? (rseq sv2))")}))
					}
				}
			}
			{
				var v1_5127 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(15), float64(48)).(*cljs_core.CljsCoreRange))
				var v2_5128 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(40), float64(57)).(*cljs_core.CljsCoreRange))
				var v1_5129___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Pop_BANG_.X_invoke_Arity1(cljs_core.Transient.X_invoke_Arity1(v1_5127)), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v2_5130___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(v2_5128), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v_5131 = cljs_core.Into.X_invoke_Arity2(v1_5129___1, v2_5130___1)
				_, _, _, _, _ = v1_5127, v2_5128, v1_5129___1, v2_5130___1, v_5131
				if cljs_core.X_EQ_.Arity2IIB(v_5131, cljs_core.Vec.X_invoke_Arity1(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(16), float64(47)).(*cljs_core.CljsCoreRange), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(41), float64(57)).(*cljs_core.CljsCoreRange), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})})).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (= v (vec (concat [:quux] (range 16 47) [:foo] [:quux] (range 41 57) [:bar])))")}))
				}
			}
			{
				var v_5132 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)
				var xs_5133 interface{} = cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)
				_, _ = v_5132, xs_5133
				for {
					{
						var temp__4220__auto___5134 = cljs_core.First.X_invoke_Arity1(xs_5133)
						_ = temp__4220__auto___5134
						if cljs_core.Truth_(temp__4220__auto___5134) {
							{
								var x_5135 = temp__4220__auto___5134
								_ = x_5135
								v_5132, xs_5133 = func() interface{} {
									var pred__4659 = func(G__5136 *cljs_core.AFn, v_5132 interface{}, xs_5133 interface{}, x_5135 interface{}, temp__4220__auto___5134 interface{}) *cljs_core.AFn {
										return cljs_core.Fn(G__5136, 2, func(p1__4087_SHARP_ interface{}, p2__4088_SHARP_ interface{}) interface{} {
											{
												var G__4662 = cljs_core.Mod.X_invoke_Arity2(p2__4088_SHARP_, float64(3)).(float64)
												_ = G__4662
												return p1__4087_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4662)
											}
										})
									}(&cljs_core.AFn{}, v_5132, xs_5133, x_5135, temp__4220__auto___5134)
									var expr__4660 = x_5135
									_, _ = pred__4659, expr__4660
									if cljs_core.Truth_(pred__4659.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(0), nil, float64(2), nil}, nil}, nil}), expr__4660)) {
										return cljs_core.Conj_BANG_.X_invoke_Arity2(v_5132, x_5135)
									} else {
										if cljs_core.Truth_(pred__4659.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), expr__4660)) {
											return cljs_core.Assoc_BANG_.X_invoke_Arity3(v_5132, cljs_core.Count.X_invoke_Arity1(v_5132).(float64), x_5135)
										} else {
											panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(expr__4660).(string))}))
										}
									}
								}(), cljs_core.Next.Arity1IQ(xs_5133)
								continue
							}
						} else {
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)), cljs_core.Persistent_BANG_.X_invoke_Arity1(v_5132)) {
							} else {
								panic((&js.Error{("Assert failed: (= (vec (range 100)) (persistent! v))")}))
							}
						}
					}
					break
				}
			}
			{
				var m1_5137 interface{} = cljs_core.CljsCorePersistentHashMap_EMPTY
				var m2_5138 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_EMPTY)
				var i_5139 = float64(0)
				_, _, _ = m1_5137, m2_5138, i_5139
				for {
					if i_5139 < float64(100) {
						m1_5137, m2_5138, i_5139 = cljs_core.Assoc.X_invoke_Arity3(m1_5137, i_5139, i_5139), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_5138, i_5139, i_5139), (i_5139 + float64(1))
						continue
					} else {
						{
							var m2_5140___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_5138)
							_ = m2_5140___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_5137).(float64), float64(100)) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m1) 100)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_5140___1).(float64), float64(100)) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m2) 100)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_5137, m2_5140___1) {
							} else {
								panic((&js.Error{("Assert failed: (= m1 m2)")}))
							}
							{
								var i_5141___1 = float64(0)
								_ = i_5141___1
								for {
									if i_5141___1 < float64(100) {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__4663 = i_5141___1
											_ = G__4663
											return m1_5137.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4663)
										}(), i_5141___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__4664 = i_5141___1
											_ = G__4664
											return m2_5140___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4664)
										}(), i_5141___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m2 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_5137, i_5141___1), i_5141___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_5140___1, i_5141___1), i_5141___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m2 i) i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_5137, i_5141___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m1 i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_5140___1, i_5141___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m2 i)")}))
										}
										i_5141___1 = (i_5141___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_5137))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range 100) (range 100)) (sort-by first (seq m1)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_5140___1))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range 100) (range 100)) (sort-by first (seq m2)))")}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_5137, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{("Assert failed: (not (contains? (dissoc m1 3) 3))")}))
							}
						}
					}
					break
				}
			}
			{
				var m_5142 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_5142
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5142).(float64), float64(7)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 7)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5142, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
				}
			}
			{
				var m_5143 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_5143
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5143).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5143, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var m_5144 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_5144
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5144).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5144, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var tm_5145 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_5145
				{
					var tm_5146___1 interface{} = tm_5145
					var ks_5147 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_5146___1, ks_5147
					for {
						{
							var temp__4220__auto___5148 = cljs_core.First.X_invoke_Arity1(ks_5147)
							_ = temp__4220__auto___5148
							if cljs_core.Truth_(temp__4220__auto___5148) {
								{
									var k_5149 = temp__4220__auto___5148
									_ = k_5149
									tm_5146___1, ks_5147 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_5146___1, k_5149), cljs_core.Next.Arity1IQ(ks_5147)
									continue
								}
							} else {
								{
									var m_5150 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5146___1)
									_ = m_5150
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5150).(float64), float64(7)) {
									} else {
										panic((&js.Error{("Assert failed: (= (count m) 7)")}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_5150, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
									} else {
										panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
									}
								}
							}
						}
						break
					}
				}
			}
			{
				var tm_5151 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_5151
				{
					var seq__4665_5152 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__4666_5153 interface{} = nil
					var count__4667_5154 = float64(0)
					var i__4668_5155 = float64(0)
					_, _, _, _ = seq__4665_5152, chunk__4666_5153, count__4667_5154, i__4668_5155
					for {
						if i__4668_5155 < count__4667_5154 {
							{
								var k_5156 = chunk__4666_5153.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4668_5155)
								_ = k_5156
								if cljs_core.X_EQ_.Arity2IIB(k_5156, cljs_core.Get.X_invoke_Arity2(tm_5151, k_5156)) {
								} else {
									panic((&js.Error{("Assert failed: (= k (get tm k))")}))
								}
								seq__4665_5152, chunk__4666_5153, count__4667_5154, i__4668_5155 = seq__4665_5152, chunk__4666_5153, count__4667_5154, (i__4668_5155 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___5157 = cljs_core.Seq.Arity1IQ(seq__4665_5152)
								_ = temp__4222__auto___5157
								if cljs_core.Truth_(temp__4222__auto___5157) {
									{
										var seq__4665_5158___1 = temp__4222__auto___5157
										_ = seq__4665_5158___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4665_5158___1) {
											{
												var c__960__auto___5159 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4665_5158___1)
												_ = c__960__auto___5159
												seq__4665_5152, chunk__4666_5153, count__4667_5154, i__4668_5155 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4665_5158___1), c__960__auto___5159, cljs_core.Count.X_invoke_Arity1(c__960__auto___5159).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_5160 = cljs_core.First.X_invoke_Arity1(seq__4665_5158___1)
												_ = k_5160
												if cljs_core.X_EQ_.Arity2IIB(k_5160, cljs_core.Get.X_invoke_Arity2(tm_5151, k_5160)) {
												} else {
													panic((&js.Error{("Assert failed: (= k (get tm k))")}))
												}
												seq__4665_5152, chunk__4666_5153, count__4667_5154, i__4668_5155 = cljs_core.Next.Arity1IQ(seq__4665_5158___1), nil, float64(0), float64(0)
												continue
											}
										}
									}
								} else {
								}
							}
						}
						break
					}
				}
				{
					var m_5161 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5151)
					_ = m_5161
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5162 float64) {
						defer func() {
							if e4669 := recover(); e4669 != nil {
								if cljs_core.Value_(e4669).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4669
										_ = e
										return__5162 = float64(2)
									}
								} else {
									panic(e4669)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_5151, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5163 float64) {
						defer func() {
							if e4670 := recover(); e4670 != nil {
								if cljs_core.Value_(e4670).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4670
										_ = e
										return__5163 = float64(2)
									}
								} else {
									panic(e4670)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_5151, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5164 float64) {
						defer func() {
							if e4671 := recover(); e4671 != nil {
								if cljs_core.Value_(e4671).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4671
										_ = e
										return__5164 = float64(2)
									}
								} else {
									panic(e4671)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5151)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (persistent! tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5165 float64) {
						defer func() {
							if e4672 := recover(); e4672 != nil {
								if cljs_core.Value_(e4672).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4672
										_ = e
										return__5165 = float64(2)
									}
								} else {
									panic(e4672)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_5151)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (count tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_5161, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
					}
				}
			}
			X__GT_FixedHash = func(__GT_FixedHash *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_FixedHash, 2, func(h interface{}, v interface{}) interface{} {
					return (&CljsCore_testFixedHash{h, v})
				})
			}(&cljs_core.AFn{})

			Fixed_hash_foo = (&CljsCore_testFixedHash{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})})

			Fixed_hash_bar = (&CljsCore_testFixedHash{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})})

			{
				var m_5166 = cljs_core.Assoc.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentHashMap_EMPTY, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_ = m_5166
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5166, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5166, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5166).(float64), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 2)")}))
				}
				{
					var m_5167___1 = cljs_core.Dissoc.X_invoke_Arity2(m_5166, Fixed_hash_foo)
					_ = m_5167___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5167___1, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_5167___1, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5167___1).(float64), float64(1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 1)")}))
					}
				}
			}
			{
				var m_5168 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_5169___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m_5168, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_, _ = m_5168, m_5169___1
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5169___1).(float64), float64(102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 102)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5169___1, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5169___1, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				{
					var m_5170___2 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_5169___1, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_5170___2
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5170___2, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_5170___2, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5170___2).(float64), float64(98)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 98)")}))
					}
				}
			}
			{
				var m_5171 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_5172___1 = cljs_core.Transient.X_invoke_Arity1(m_5171)
				var m_5173___2 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_5172___1, Fixed_hash_foo, float64(1))
				var m_5174___3 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_5173___2, Fixed_hash_bar, float64(2))
				var m_5175___4 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m_5174___3)
				_, _, _, _, _ = m_5171, m_5172___1, m_5173___2, m_5174___3, m_5175___4
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5175___4).(float64), float64(102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 102)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5175___4, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5175___4, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				{
					var m_5176___5 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_5175___4, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_5176___5
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5176___5, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_5176___5, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5176___5).(float64), float64(98)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 98)")}))
					}
				}
			}
			Array_map_conversion_threshold = cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD

			{
				var m1_5177 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_5178 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY)
				var i_5179 = float64(0)
				_, _, _ = m1_5177, m2_5178, i_5179
				for {
					if i_5179 < Array_map_conversion_threshold {
						m1_5177, m2_5178, i_5179 = cljs_core.Assoc.X_invoke_Arity3(m1_5177, i_5179, i_5179), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_5178, i_5179, i_5179), (i_5179 + float64(1))
						continue
					} else {
						{
							var m2_5180___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_5178)
							_ = m2_5180___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_5177).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m1) array-map-conversion-threshold)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_5180___1).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m2) array-map-conversion-threshold)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_5177, m2_5180___1) {
							} else {
								panic((&js.Error{("Assert failed: (= m1 m2)")}))
							}
							{
								var i_5181___1 = float64(0)
								_ = i_5181___1
								for {
									if i_5181___1 < Array_map_conversion_threshold {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__4673 = i_5181___1
											_ = G__4673
											return m1_5177.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4673)
										}(), i_5181___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__4674 = i_5181___1
											_ = G__4674
											return m2_5180___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4674)
										}(), i_5181___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m2 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_5177, i_5181___1), i_5181___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_5180___1, i_5181___1), i_5181___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m2 i) i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_5177, i_5181___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m1 i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_5180___1, i_5181___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m2 i)")}))
										}
										i_5181___1 = (i_5181___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_5177))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m1)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_5180___1))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m2)))")}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_5177, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{("Assert failed: (not (contains? (dissoc m1 3) 3))")}))
							}
						}
					}
					break
				}
			}
			{
				var m_5182 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_5182
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5182).(float64), float64(7)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 7)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5182, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
				}
			}
			{
				var m_5183 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_5183
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5183).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5183, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var m_5184 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_5184
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5184).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5184, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var tm_5185 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_5185
				{
					var tm_5186___1 interface{} = tm_5185
					var ks_5187 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_5186___1, ks_5187
					for {
						{
							var temp__4220__auto___5188 = cljs_core.First.X_invoke_Arity1(ks_5187)
							_ = temp__4220__auto___5188
							if cljs_core.Truth_(temp__4220__auto___5188) {
								{
									var k_5189 = temp__4220__auto___5188
									_ = k_5189
									tm_5186___1, ks_5187 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_5186___1, k_5189), cljs_core.Next.Arity1IQ(ks_5187)
									continue
								}
							} else {
								{
									var m_5190 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5186___1)
									_ = m_5190
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5190).(float64), float64(7)) {
									} else {
										panic((&js.Error{("Assert failed: (= (count m) 7)")}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_5190, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
									} else {
										panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
									}
								}
							}
						}
						break
					}
				}
			}
			{
				var tm_5191 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_5191
				{
					var seq__4675_5192 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__4676_5193 interface{} = nil
					var count__4677_5194 = float64(0)
					var i__4678_5195 = float64(0)
					_, _, _, _ = seq__4675_5192, chunk__4676_5193, count__4677_5194, i__4678_5195
					for {
						if i__4678_5195 < count__4677_5194 {
							{
								var k_5196 = chunk__4676_5193.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4678_5195)
								_ = k_5196
								if cljs_core.X_EQ_.Arity2IIB(k_5196, cljs_core.Get.X_invoke_Arity2(tm_5191, k_5196)) {
								} else {
									panic((&js.Error{("Assert failed: (= k (get tm k))")}))
								}
								seq__4675_5192, chunk__4676_5193, count__4677_5194, i__4678_5195 = seq__4675_5192, chunk__4676_5193, count__4677_5194, (i__4678_5195 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___5197 = cljs_core.Seq.Arity1IQ(seq__4675_5192)
								_ = temp__4222__auto___5197
								if cljs_core.Truth_(temp__4222__auto___5197) {
									{
										var seq__4675_5198___1 = temp__4222__auto___5197
										_ = seq__4675_5198___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4675_5198___1) {
											{
												var c__960__auto___5199 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4675_5198___1)
												_ = c__960__auto___5199
												seq__4675_5192, chunk__4676_5193, count__4677_5194, i__4678_5195 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4675_5198___1), c__960__auto___5199, cljs_core.Count.X_invoke_Arity1(c__960__auto___5199).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_5200 = cljs_core.First.X_invoke_Arity1(seq__4675_5198___1)
												_ = k_5200
												if cljs_core.X_EQ_.Arity2IIB(k_5200, cljs_core.Get.X_invoke_Arity2(tm_5191, k_5200)) {
												} else {
													panic((&js.Error{("Assert failed: (= k (get tm k))")}))
												}
												seq__4675_5192, chunk__4676_5193, count__4677_5194, i__4678_5195 = cljs_core.Next.Arity1IQ(seq__4675_5198___1), nil, float64(0), float64(0)
												continue
											}
										}
									}
								} else {
								}
							}
						}
						break
					}
				}
				{
					var m_5201 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5191)
					_ = m_5201
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5202 float64) {
						defer func() {
							if e4679 := recover(); e4679 != nil {
								if cljs_core.Value_(e4679).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4679
										_ = e
										return__5202 = float64(2)
									}
								} else {
									panic(e4679)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_5191, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5203 float64) {
						defer func() {
							if e4680 := recover(); e4680 != nil {
								if cljs_core.Value_(e4680).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4680
										_ = e
										return__5203 = float64(2)
									}
								} else {
									panic(e4680)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_5191, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5204 float64) {
						defer func() {
							if e4681 := recover(); e4681 != nil {
								if cljs_core.Value_(e4681).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4681
										_ = e
										return__5204 = float64(2)
									}
								} else {
									panic(e4681)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5191)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (persistent! tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5205 float64) {
						defer func() {
							if e4682 := recover(); e4682 != nil {
								if cljs_core.Value_(e4682).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4682
										_ = e
										return__5205 = float64(2)
									}
								} else {
									panic(e4682)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_5191)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (count tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_5201, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
					}
				}
			}
			{
				var m_5206 = cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))
				_ = m_5206
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5206).(float64), (float64(2) * Array_map_conversion_threshold)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) (* 2 array-map-conversion-threshold))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4683 = Array_map_conversion_threshold
					_ = G__4683
					return m_5206.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4683)
				}(), Array_map_conversion_threshold) {
				} else {
					panic((&js.Error{("Assert failed: (= (m array-map-conversion-threshold) array-map-conversion-threshold)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5206, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity2(func(G__5207 *cljs_core.AFn, m_5206 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5207, 1, func(p1__4089_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__4089_SHARP_, p1__4089_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_5206), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (= m (into (.-EMPTY cljs.core/PersistentHashMap) (map (fn* [p1__4089#] (vector p1__4089# p1__4089#)) (range (* 2 array-map-conversion-threshold)))))")}))
				}
			}
			{
				var m1_5208 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_5209 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var i_5210 = float64(0)
				_, _, _ = m1_5208, m2_5209, i_5210
				for {
					if i_5210 < float64(100) {
						m1_5208, m2_5209, i_5210 = cljs_core.Assoc.X_invoke_Arity3(m1_5208, i_5210, i_5210), cljs_core.Assoc.X_invoke_Arity3(m2_5209, ("foo"+cljs_core.Str.X_invoke_Arity1(i_5210).(string)), i_5210), (i_5210 + float64(1))
						continue
					} else {
						if cljs_core.X_EQ_.Arity2IIB(m1_5208, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{("Assert failed: (= m1 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (range 100) (range 100))))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(m2_5209, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Map_.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.Str, "foo").(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{("Assert failed: (= m2 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (map (partial str \"foo\") (range 100)) (range 100))))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_5208).(float64), float64(100)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count m1) 100)")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_5209).(float64), float64(100)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count m2) 100)")}))
						}
					}
					break
				}
			}
			{
				var i_5211 = float64(0)
				var m_5212 interface{} = cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(-1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))
				var result_5213 interface{} = cljs_core.CljsCorePersistentVector_EMPTY
				_, _, _ = i_5211, m_5212, result_5213
				for {
					if i_5211 <= (cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) {
						i_5211, m_5212, result_5213 = (i_5211 + float64(1)), cljs_core.Assoc.X_invoke_Arity3(m_5212, i_5211, i_5211), cljs_core.Conj.X_invoke_Arity2(result_5213, cljs_core.Meta.X_invoke_Arity1(m_5212))
						continue
					} else {
						{
							var n_5214 = ((cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) + float64(1))
							var expected_5215 = cljs_core.Repeat.X_invoke_Arity2(n_5214, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})).(*cljs_core.CljsCoreLazySeq)
							_, _ = n_5214, expected_5215
							if cljs_core.X_EQ_.Arity2IIB(result_5213, expected_5215) {
							} else {
								panic((&js.Error{("Assert failed: (= result expected)")}))
							}
						}
					}
					break
				}
			}
			{
				var s_5216 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_EMPTY)
				var i_5217 = float64(0)
				_, _ = s_5216, i_5217
				for {
					if i_5217 < float64(100) {
						s_5216, i_5217 = cljs_core.Conj_BANG_.X_invoke_Arity2(s_5216, i_5217), (i_5217 + float64(1))
						continue
					} else {
						{
							var s_5218___1 interface{} = s_5216
							var i_5219___1 = float64(0)
							_, _ = s_5218___1, i_5219___1
							for {
								if i_5219___1 < float64(100) {
									if cljs_core.Mod.X_invoke_Arity2(i_5219___1, float64(3)).(float64) == float64(0) {
										s_5218___1, i_5219___1 = cljs_core.Disj_BANG_.X_invoke_Arity2(s_5218___1, i_5219___1), (i_5219___1 + float64(1))
										continue
									} else {
										s_5218___1, i_5219___1 = s_5218___1, (i_5219___1 + float64(1))
										continue
									}
								} else {
									{
										var s_5220___2 = cljs_core.Persistent_BANG_.X_invoke_Arity1(s_5218___1)
										_ = s_5220___2
										if cljs_core.X_EQ_.Arity2IIB(s_5220___2, func() interface{} {
											var s___3 interface{} = cljs_core.CljsCorePersistentHashSet_EMPTY
											var xs interface{} = cljs_core.Remove.X_invoke_Arity2(func(G__5221 *cljs_core.AFn, s_5218___1 interface{}, i_5219___1 float64, s_5216 interface{}, i_5217 float64, s___3 interface{}, s_5220___2 interface{}) *cljs_core.AFn {
												return cljs_core.Fn(G__5221, 1, func(p1__4090_SHARP_ interface{}) interface{} {
													return (cljs_core.Mod.X_invoke_Arity2(p1__4090_SHARP_, float64(3)).(float64) == float64(0))
												})
											}(&cljs_core.AFn{}, s_5218___1, i_5219___1, s_5216, i_5217, s___3, s_5220___2), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
											_, _ = s___3, xs
											for {
												{
													var temp__4220__auto__ = cljs_core.First.X_invoke_Arity1(xs)
													_ = temp__4220__auto__
													if cljs_core.Truth_(temp__4220__auto__) {
														{
															var x = temp__4220__auto__
															_ = x
															s___3, xs = cljs_core.Conj.X_invoke_Arity2(s___3, x), cljs_core.Next.Arity1IQ(xs)
															continue
														}
													} else {
														return s___3
													}
												}
											}
										}()) {
										} else {
											panic((&js.Error{("Assert failed: (= s (loop [s #{} xs (remove (fn* [p1__4090#] (zero? (mod p1__4090# 3))) (range 100))] (if-let [x (first xs)] (recur (conj s x) (next xs)) s)))")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(s_5220___2, cljs_core.Set.X_invoke_Arity1(cljs_core.Remove.X_invoke_Arity2(func(G__5222 *cljs_core.AFn, s_5218___1 interface{}, i_5219___1 float64, s_5216 interface{}, i_5217 float64, s_5220___2 interface{}) *cljs_core.AFn {
											return cljs_core.Fn(G__5222, 1, func(p1__4091_SHARP_ interface{}) interface{} {
												return (cljs_core.Mod.X_invoke_Arity2(p1__4091_SHARP_, float64(3)).(float64) == float64(0))
											})
										}(&cljs_core.AFn{}, s_5218___1, i_5219___1, s_5216, i_5217, s_5220___2), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
										} else {
											panic((&js.Error{("Assert failed: (= s (set (remove (fn* [p1__4091#] (zero? (mod p1__4091# 3))) (range 100))))")}))
										}
									}
								}
								break
							}
						}
					}
					break
				}
			}
			{
				var m1_5223 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_5224 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var m2_5225 = cljs_core.Sorted_map_by.X_invoke_ArityVariadic(c2_5224, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{})).(*cljs_core.CljsCorePersistentTreeMap)
				_, _, _ = m1_5223, c2_5224, m2_5225
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeMap)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(m1_5223)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeMap (type m1))")}))
				}
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeMap)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(m2_5225)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeMap (type m2))")}))
				}
				if reflect.DeepEqual(cljs_core.Compare, cljs_core.Native_get_instance_field.X_invoke_Arity2(m1_5223, "Comp")) {
				} else {
					panic((&js.Error{("Assert failed: (identical? compare (.-comp m1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(m1_5223).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count m1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(m2_5225).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count m2))")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(m1_5223)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq m1))")}))
				}
				{
					var m1_5226___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m1_5223, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					var m2_5227___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m2_5225, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					_, _ = m1_5226___1, m2_5227___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_5226___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m1) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_5227___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m2) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_5226___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m1) (list [:bar 2] [:foo 1] [:quux 3]))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_5227___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m2) (list [:quux 3] [:foo 1] [:bar 2]))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_5226___1), cljs_core.Rseq.Arity1IQ(m2_5227___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m1) (rseq m2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_5227___1), cljs_core.Rseq.Arity1IQ(m1_5226___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m2) (rseq m1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m1_5226___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= (conj m1 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m1_5226___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count (conj m1 [:wibble 4])) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m2_5227___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= (conj m2 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m2_5227___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count (conj m2 [:wibble 4])) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m1_5226___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})).(cljs_core.CljsCoreICollection).X_conj_Arity2(nil)) {
					} else {
						panic((&js.Error{("Assert failed: (= (map key (assoc m1 nil 4)) (list nil :bar :foo :quux))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m2_5227___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(nil).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (map key (assoc m2 nil 4)) (list :quux :foo :bar nil))")}))
					}
				}
			}
			{
				var m_5228 = cljs_core.Apply.X_invoke_Arity2(cljs_core.Sorted_map, cljs_core.Mapcat.X_invoke_ArityVariadic(func(G__5231 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5231, 1, func(p1__4092_SHARP_ interface{}) interface{} {
						return cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(p1__4092_SHARP_).(cljs_core.CljsCoreICollection).X_conj_Arity2(p1__4092_SHARP_)
					})
				}(&cljs_core.AFn{}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_ArityVariadic(cljs_core.Partial.X_invoke_Arity2(cljs_core.Apply, cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(10)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(20), float64(30)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(50), float64(60)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(30), float64(40)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(40), float64(50)}, nil})}, nil})}))})))
				var s1_5229 = cljs_core.Map_.X_invoke_Arity2(func(G__5232 *cljs_core.AFn, m_5228 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5232, 1, func(p1__4093_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__4093_SHARP_, p1__4093_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_5228), cljs_core.Range_.X_invoke_Arity1(float64(60)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				var s2_5230 = cljs_core.Map_.X_invoke_Arity2(func(G__5233 *cljs_core.AFn, m_5228 interface{}, s1_5229 *cljs_core.CljsCoreLazySeq) *cljs_core.AFn {
					return cljs_core.Fn(G__5233, 1, func(p1__4094_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__4094_SHARP_, p1__4094_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_5228, s1_5229), cljs_core.Range_.X_invoke_Arity3(float64(59), float64(-1), float64(-1)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				_, _, _ = m_5228, s1_5229, s2_5230
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5228).(float64), float64(60)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 60)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m_5228), s1_5229) {
				} else {
					panic((&js.Error{("Assert failed: (= (seq m) s1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(m_5228), s2_5230) {
				} else {
					panic((&js.Error{("Assert failed: (= (rseq m) s2)")}))
				}
			}
			{
				var m_5234 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
				_ = m_5234
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m_5234, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(2), float64(3)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (dissoc m :foo) (hash-map :bar 2 :quux 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(m_5234, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))).(float64), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count (dissoc m :foo)) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(m_5234), cljs_core.Hash.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(1), float64(2), float64(3)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (hash m) (hash (hash-map :foo 1 :bar 2 :quux 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_5234, cljs_core.X_LT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m < :foo) (list [:bar 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_5234, cljs_core.X_LT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m <= :foo) (list [:bar 2] [:foo 1]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_5234, cljs_core.X_GT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m > :foo) (list [:quux 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_5234, cljs_core.X_GT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m >= :foo) (list [:foo 1] [:quux 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__5235 *cljs_core.AFn, m_5234 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5235, 1, func(p1__4095_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity2(func(G__5236 *cljs_core.AFn, m_5234 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__5236, 2, func(___ interface{}, x interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_5234), p1__4095_SHARP_)
					})
				}(&cljs_core.AFn{}, m_5234), m_5234).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__4095#] (reduce (fn [_ x] x) p1__4095#)) m) (list 2 1 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__5237 *cljs_core.AFn, m_5234 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5237, 1, func(p1__4096_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity3(func(G__5238 *cljs_core.AFn, m_5234 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__5238, 2, func(x interface{}, ___ interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_5234), float64(7), p1__4096_SHARP_)
					})
				}(&cljs_core.AFn{}, m_5234), m_5234).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__4096#] (reduce (fn [x _] x) 7 p1__4096#)) m) (list 7 7 7))")}))
				}
			}
			{
				var s1_5239 = cljs_core.Sorted_set.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_5240 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var s2_5241 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c2_5240, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c3_5242 = func(G__5245 *cljs_core.AFn, s1_5239 interface{}, c2_5240 cljs_core.CljsCoreIFn, s2_5241 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5245, 2, func(p1__4097_SHARP_ interface{}, p2__4098_SHARP_ interface{}) interface{} {
						return cljs_core.Compare.Arity2IIF(cljs_core.Quot.X_invoke_Arity2(p1__4097_SHARP_, float64(2)).(float64), cljs_core.Quot.X_invoke_Arity2(p2__4098_SHARP_, float64(2)).(float64))
					})
				}(&cljs_core.AFn{}, s1_5239, c2_5240, s2_5241)
				var s3_5243 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c3_5242, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var s4_5244 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(cljs_core.X_LT_, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				_, _, _, _, _, _ = s1_5239, c2_5240, s2_5241, c3_5242, s3_5243, s4_5244
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(s1_5239)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeSet (type s1))")}))
				}
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(s2_5241)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeSet (type s2))")}))
				}
				if reflect.DeepEqual(cljs_core.Compare, s1_5239.(cljs_core.CljsCoreISorted).X_comparator_Arity1()) {
				} else {
					panic((&js.Error{("Assert failed: (identical? compare (-comparator s1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(s1_5239).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count s1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(s2_5241).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count s2))")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(s1_5239)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq s1))")}))
				}
				{
					var s1_5246___1 = cljs_core.Conj.X_invoke_ArityVariadic(s1_5239, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s2_5247___1 = cljs_core.Conj.X_invoke_ArityVariadic(s2_5241, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s3_5248___1 = cljs_core.Conj.X_invoke_ArityVariadic(s3_5243, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(7), float64(8), float64(9)}))
					var s4_5249___1 = cljs_core.Conj.X_invoke_ArityVariadic(s4_5244, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					_, _, _, _ = s1_5246___1, s2_5247___1, s3_5248___1, s4_5249___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_5246___1), cljs_core.Hash.X_invoke_Arity1(s2_5247___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (hash s1) (hash s2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_5246___1), cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (hash s1) (hash #{1 3 2}))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_5246___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq s1) (list 1 2 3))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_5246___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq s1) (list 3 2 1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_5247___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq s2) (list 3 2 1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_5247___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq s2) (list 1 2 3))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_5246___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s1) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_5247___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s2) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s3_5248___1).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s3) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(s3_5248___1, float64(0)), float64(1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get s3 0) 1)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5248___1, cljs_core.X_GT_, float64(5)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 > 5) (list 7 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5248___1, cljs_core.X_GT_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 > 6) (list 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5248___1, cljs_core.X_GT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 6) (list 7 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5248___1, cljs_core.X_GT__EQ_, float64(12)), nil) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 12) nil)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5248___1, cljs_core.X_LT_, float64(0)), cljs_core.CljsCoreList_EMPTY) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 0) (list))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5248___1, cljs_core.X_LT_, float64(5)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 5) (list 1 2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5248___1, cljs_core.X_LT_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 6) (list 1 2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5248___1, cljs_core.X_LT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 <= 6) (list 1 2 7))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s3_5248___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 2 <= 6) (list 2 7))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s4_5249___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT_, float64(3)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s4 >= 2 < 3) (list 2))")}))
					}
					{
						var s1_5250___2 = cljs_core.Disj.X_invoke_Arity2(s1_5246___1, float64(2))
						var s2_5251___2 = cljs_core.Disj.X_invoke_Arity2(s2_5247___1, float64(2))
						_, _ = s1_5250___2, s2_5251___2
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_5250___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{("Assert failed: (= (seq s1) (list 1 3))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_5250___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{("Assert failed: (= (rseq s1) (list 3 1))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_5251___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{("Assert failed: (= (seq s2) (list 3 1))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_5251___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{("Assert failed: (= (rseq s2) (list 1 3))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_5250___2).(float64), float64(2)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count s1) 2)")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_5251___2).(float64), float64(2)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count s2) 2)")}))
						}
					}
				}
			}
			{
				X__GT_Person = func(__GT_Person *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_Person, 2, func(firstname interface{}, lastname interface{}) interface{} {
						return (&CljsCore_testPerson{firstname, lastname, nil, nil, nil})
					})
				}(&cljs_core.AFn{})

				Map__GT_Person = func(map__GT_Person *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(map__GT_Person, 1, func(G__4686 interface{}) interface{} {
						return (&CljsCore_testPerson{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}).X_invoke_Arity1(G__4686), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}).X_invoke_Arity1(G__4686), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__4686, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})})), nil})
					})
				}(&cljs_core.AFn{})

			}
			Fred = (&CljsCore_testPerson{"Fred", "Mertz", nil, nil, nil})

			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}).X_invoke_Arity1(Fred), "Fred") {
			} else {
				panic((&js.Error{("Assert failed: (= (:firstname fred) \"Fred\")")}))
			}
			Fred_too = (&CljsCore_testPerson{"Fred", "Mertz", nil, nil, nil})

			if cljs_core.X_EQ_.Arity2IIB(Fred, Fred_too) {
			} else {
				panic((&js.Error{("Assert failed: (= fred fred-too)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(Fred, nil) == false {
			} else {
				panic((&js.Error{("Assert failed: (false? (= fred nil))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, Fred) == false {
			} else {
				panic((&js.Error{("Assert failed: (false? (= nil fred))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(Map__GT_Person.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), "Fred", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Mertz"}, nil})).(*CljsCore_testPerson), Fred) {
			} else {
				panic((&js.Error{("Assert failed: (= (map->Person {:firstname \"Fred\", :lastname \"Mertz\"}) fred)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(X__GT_Person.X_invoke_Arity2("Fred", "Mertz").(*CljsCore_testPerson), Fred) {
			} else {
				panic((&js.Error{("Assert failed: (= (->Person \"Fred\" \"Mertz\") fred)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(Fred).(float64), float64(2)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count fred) 2)")}))
			}
			{
				X__GT_A = func(__GT_A *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_A, 0, func() interface{} {
						return (&CljsCore_testA{nil, nil, nil})
					})
				}(&cljs_core.AFn{})

				Map__GT_A = func(map__GT_A *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(map__GT_A, 1, func(G__4705 interface{}) interface{} {
						return (&CljsCore_testA{nil, cljs_core.Dissoc.X_invoke_Arity1(G__4705), nil})
					})
				}(&cljs_core.AFn{})

			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})}, nil}), cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&CljsCore_testA{nil, nil, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})}, nil})))) {
			} else {
				panic((&js.Error{("Assert failed: (= {:foo (quote bar)} (meta (with-meta (A.) {:foo (quote bar)})))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}).X_invoke_Arity1(cljs_core.Assoc.X_invoke_Arity3((&CljsCore_testA{nil, nil, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote bar) (:foo (assoc (A.) :foo (quote bar))))")}))
			}
			{
				X__GT_C = func(__GT_C *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_C, 3, func(a interface{}, b interface{}, c interface{}) interface{} {
						return (&CljsCore_testC{a, b, c, nil, nil, nil})
					})
				}(&cljs_core.AFn{})

				Map__GT_C = func(map__GT_C *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(map__GT_C, 1, func(G__4716 interface{}) interface{} {
						return (&CljsCore_testC{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__4716), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__4716), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}).X_invoke_Arity1(G__4716), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__4716, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})})), nil})
					})
				}(&cljs_core.AFn{})

			}
			Letters = (&CljsCore_testC{"a", "b", "c", nil, nil, nil})

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(Letters)), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (set (keys letters)) #{:c :b :a})")}))
			}
			More_letters = cljs_core.Assoc.X_invoke_ArityVariadic(Letters, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), "d", cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), "e", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), "f"}))

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(More_letters)), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(6), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (set (keys more-letters)) #{:e :c :b :d :f :a})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(More_letters, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})))), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(5), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (set (keys (dissoc more-letters :d))) #{:e :c :b :f :a})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(More_letters, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)})})))), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (set (keys (dissoc more-letters :d :e))) #{:c :b :f :a})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(More_letters, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)})})))), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (set (keys (dissoc more-letters :d :e :f))) #{:c :b :a})")}))
			}
			{
				var s_5252 = "abc"
				_ = s_5252
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_5252).Length) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (.-length s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_5252).Length) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (. s -length))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_((`` + cljs_core.Str.X_invoke_Arity1(float64(138)).(string))).Length) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (. (str 138) -length))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_("abc").Length) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (. \"abc\" -length))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_5252).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_("abc").Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring \"abc\" 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", func(G__5253 *cljs_core.AFn, s_5252 string) *cljs_core.AFn {
					return cljs_core.Fn(G__5253, 2, func(target4737 interface{}, start interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target4737, "Substring", []interface{}{start})
					})
				}(&cljs_core.AFn{}, s_5252).X_invoke_Arity2(s_5252, float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" ((memfn substring start) s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_5252).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s substring 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_5252).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s (substring 1)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_5252).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s (substring 1 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_5252).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring s 1 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_(s_5252).ToUpperCase()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" (. s (toUpperCase)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_("abc").ToUpperCase()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" (. \"abc\" (toUpperCase)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", func(G__5254 *cljs_core.AFn, s_5252 string) *cljs_core.AFn {
					return cljs_core.Fn(G__5254, 1, func(target4738 interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target4738, "ToUpperCase", []interface{}{})
					})
				}(&cljs_core.AFn{}, s_5252).X_invoke_Arity1(s_5252)) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" ((memfn toUpperCase) s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("BC", cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_5252).ToUpperCase(), "Substring", []interface{}{float64(1)})) {
				} else {
					panic((&js.Error{("Assert failed: (= \"BC\" (. (. s (toUpperCase)) substring 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_5252).ToUpperCase(), "Substring", []interface{}{float64(1)}), "Length")) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (.-length (. (. s (toUpperCase)) substring 1)))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(Fred, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wife", Fqn: "wife", X_hash: float64(2005722828)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ethel", Fqn: "ethel", X_hash: float64(-1460626342)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "friend", Fqn: "friend", X_hash: float64(-286879240)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ricky", Fqn: "ricky", X_hash: float64(-48928873)})}, nil})), Map__GT_Person.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), "Fred", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Mertz", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wife", Fqn: "wife", X_hash: float64(2005722828)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ethel", Fqn: "ethel", X_hash: float64(-1460626342)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "friend", Fqn: "friend", X_hash: float64(-286879240)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ricky", Fqn: "ricky", X_hash: float64(-48928873)})}, nil})).(*CljsCore_testPerson)) {
			} else {
				panic((&js.Error{("Assert failed: (= (conj fred {:wife :ethel, :friend :ricky}) (map->Person {:firstname \"Fred\", :lastname \"Mertz\", :wife :ethel, :friend :ricky}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(Fred, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Flintstone"}, nil})), Map__GT_Person.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), "Fred", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Flintstone"}, nil})).(*CljsCore_testPerson)) {
			} else {
				panic((&js.Error{("Assert failed: (= (conj fred {:lastname \"Flintstone\"}) (map->Person {:firstname \"Fred\", :lastname \"Flintstone\"}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Assoc.X_invoke_Arity3(Fred, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Flintstone"), Map__GT_Person.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), "Fred", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Flintstone"}, nil})).(*CljsCore_testPerson)) {
			} else {
				panic((&js.Error{("Assert failed: (= (assoc fred :lastname \"Flintstone\") (map->Person {:firstname \"Fred\", :lastname \"Flintstone\"}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Assoc.X_invoke_Arity3(Fred, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wife", Fqn: "wife", X_hash: float64(2005722828)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ethel", Fqn: "ethel", X_hash: float64(-1460626342)})), Map__GT_Person.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), "Fred", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Mertz", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wife", Fqn: "wife", X_hash: float64(2005722828)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ethel", Fqn: "ethel", X_hash: float64(-1460626342)})}, nil})).(*CljsCore_testPerson)) {
			} else {
				panic((&js.Error{("Assert failed: (= (assoc fred :wife :ethel) (map->Person {:firstname \"Fred\", :lastname \"Mertz\", :wife :ethel}))")}))
			}
			{
				X__GT_A2 = func(__GT_A2 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_A2, 1, func(x interface{}) interface{} {
						return (&CljsCore_testA2{x, nil, nil, nil})
					})
				}(&cljs_core.AFn{})

				Map__GT_A2 = func(map__GT_A2 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(map__GT_A2, 1, func(G__4741 interface{}) interface{} {
						return (&CljsCore_testA2{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__4741), nil, cljs_core.Dissoc.X_invoke_Arity2(G__4741, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
					})
				}(&cljs_core.AFn{})

			}
			{
				X__GT_B = func(__GT_B *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_B, 1, func(x interface{}) interface{} {
						return (&CljsCore_testB{x, nil, nil, nil})
					})
				}(&cljs_core.AFn{})

				Map__GT_B = func(map__GT_B *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(map__GT_B, 1, func(G__4756 interface{}) interface{} {
						return (&CljsCore_testB{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__4756), nil, cljs_core.Dissoc.X_invoke_Arity2(G__4756, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
					})
				}(&cljs_core.AFn{})

			}
			if cljs_core.Not_EQ_.Arity2IIB((&CljsCore_testA2{nil, nil, nil, nil}), (&CljsCore_testB{nil, nil, nil, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (not= (A2. nil) (B. nil))")}))
			}
			Foo = func(foo *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(foo, 1, func(this interface{}) interface{} {
					return this.(CljsCore_testIFoo).Foo_Arity1()
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(func() *CljsCore_testT4769 {
				X__GT_t4769 = func(__GT_t4769 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_t4769, 2, func(test_stuff___1 interface{}, meta4770 interface{}) interface{} {
						return (&CljsCore_testT4769{test_stuff___1, meta4770})
					})
				}(&cljs_core.AFn{})

				return (&CljsCore_testT4769{test_stuff, nil})
			}(), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (meta (with-meta (reify IFoo (foo [this] :foo)) {:foo :bar})) {:foo :bar})")}))
			}
			Foo2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1070__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1071__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1072__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1073__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1074__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__, hierarchy__1074__auto__
				return (&cljs_core.CljsCoreMultiFn{"foo2", cljs_core.Identity, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1074__auto__, method_table__1070__auto__, prefer_table__1071__auto__, method_cache__1072__auto__, cached_hierarchy__1073__auto__})
			}()

			Foo2.X_add_method_Arity3(float64(0), func(G__5255 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5255, 1, func(x interface{}) interface{} {
					return x
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB(Foo2, cljs_core.Ffirst.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{Foo2, float64(1)}, true, false).(*cljs_core.CljsCorePersistentArrayMap))) {
			} else {
				panic((&js.Error{("Assert failed: (= foo2 (ffirst {foo2 1}))")}))
			}
			Mutate = func(mutate *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(mutate, 1, func(this interface{}) interface{} {
					return this.(CljsCore_testIMutate).Mutate_Arity1()
				})
			}(&cljs_core.AFn{})

			X__GT_Mutate = func(__GT_Mutate *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_Mutate, 1, func(a interface{}) interface{} {
					return (&CljsCore_testMutate{a})
				})
			}(&cljs_core.AFn{})

			X__GT_FnLike = func(__GT_FnLike *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_FnLike, 0, func() interface{} {
					return (&CljsCore_testFnLike{})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&CljsCore_testFnLike{}).X_invoke_Arity0()) {
			} else {
				panic((&js.Error{("Assert failed: (= :a ((FnLike.)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&CljsCore_testFnLike{}).X_invoke_Arity1(float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= :b ((FnLike.) 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&CljsCore_testFnLike{}).X_invoke_Arity2(float64(1), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= :c ((FnLike.) 1 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), cljs_core.Map_.X_invoke_Arity2((&CljsCore_testFnLike{}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0), float64(0)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [:b :b :b] (map (FnLike.) [0 0 0]))")}))
			}
			X__GT_FnLikeB = func(__GT_FnLikeB *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_FnLikeB, 1, func(a interface{}) interface{} {
					return (&CljsCore_testFnLikeB{a})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(float64(1), (&CljsCore_testFnLikeB{float64(1)}).X_invoke_Arity0()) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 ((FnLikeB. 1)))")}))
			}
			{
				var g_5256 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}))}, true).(*cljs_core.CljsCorePersistentHashSet)
				var h_5257 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil})}, true).(*cljs_core.CljsCorePersistentHashSet)
				_, _ = g_5256, h_5257
				if cljs_core.X_EQ_.Arity2IIB(g_5256, h_5257) {
				} else {
					panic((&js.Error{("Assert failed: (= g h)")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})), cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (hash {:a 1, :b 2}) (hash {:b 2, :a 1}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, []interface{}{float64(1), float64(2)})), cljs_core.Hash.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, []interface{}{float64(2), float64(1)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (hash (hash-map :a 1 :b 2)) (hash (hash-map :b 2 :a 1)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "start", Fqn: "start", X_hash: float64(-355208981)}), float64(133), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end", Fqn: "end", X_hash: float64(-268185958)}), float64(134)}, nil})), cljs_core.Hash.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "start", Fqn: "start", X_hash: float64(-355208981)}), float64(133), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end", Fqn: "end", X_hash: float64(-268185958)}), float64(134)}, nil})))) {
			} else {
				panic((&js.Error{("Assert failed: (= (hash {:start 133, :end 134}) (hash (apply hash-map [:start 133 :end 134])))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})), cljs_core.Hash.X_invoke_Arity1(cljs_core.Keyword.X_invoke_Arity1("a"))) {
			} else {
				panic((&js.Error{("Assert failed: (= (hash :a) (hash (keyword \"a\")))")}))
			}
			X_get_first = func(_get_first *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(_get_first, 1, func(this interface{}) interface{} {
					return this.(CljsCore_testIHasFirst).X_get_first_Arity1()
				})
			}(&cljs_core.AFn{})

			X_find_first = func(_find_first *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(_find_first, 2, func(this interface{}, other interface{}) interface{} {
					return this.(CljsCore_testIFindsFirst).X_find_first_Arity2(other)
				})
			}(&cljs_core.AFn{})

			X__GT_First = func(__GT_First *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_First, 1, func(xs interface{}) interface{} {
					return (&CljsCore_testFirst{xs})
				})
			}(&cljs_core.AFn{})

			{
				var fv_5258 = (&CljsCore_testFirst{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})})
				var fs_5259 = (&CljsCore_testFirst{"asdf"})
				_, _ = fv_5258, fs_5259
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fv_5258.X_invoke_Arity0()
				}(), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (fv) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fs_5259.X_invoke_Arity0()
				}(), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((`` + cljs_core.Str.X_invoke_Arity1(fs_5259).(string)), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (str fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_5258.X_get_first_Arity1(), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-get-first fv) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fs_5259.X_get_first_Arity1(), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (-get-first fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_5258.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-find-first fv [1]) 1)")}))
				}
				if reflect.DeepEqual(func() interface{} {
					var G__4784 = float64(1)
					_ = G__4784
					return fv_5258.X_invoke_Arity1(G__4784)
				}(), fv_5258) {
				} else {
					panic((&js.Error{("Assert failed: (identical? (fv 1) fv)")}))
				}
			}
			X__GT_DestructuringWithLocals = func(__GT_DestructuringWithLocals *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_DestructuringWithLocals, 1, func(a interface{}) interface{} {
					return (&CljsCore_testDestructuringWithLocals{a})
				})
			}(&cljs_core.AFn{})

			{
				var t_5260 = (&CljsCore_testDestructuringWithLocals{float64(1)})
				_ = t_5260
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(1)}, nil}), t_5260.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 1] (-find-first t [2 3]))")}))
				}
			}
			{
				var x_5261 = float64(1)
				_ = x_5261
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4788 = x_5261
					_ = G__4788
					switch G__4788 {
					case float64(1):
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})

					default:
						panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_5261).(string))}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x 1 :one) :one)")}))
				}
			}
			{
				var x_5263 = float64(1)
				_ = x_5263
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4789 = x_5263
					_ = G__4789
					switch G__4789 {
					case float64(2):
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "two", Fqn: "two", X_hash: float64(627606869)})

					default:
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)})

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x 2 :two :default) :default)")}))
				}
			}
			{
				var x_5265 = float64(1)
				_ = x_5265
				if cljs_core.X_EQ_.Arity2IIB(func() (return__5266 interface{}) {
					defer func() {
						if e4790 := recover(); e4790 != nil {
							if cljs_core.Value_(e4790).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4790
									_ = e
									return__5266 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e4790)

							}
						}
					}()
					{
						{
							var G__4791 = x_5265
							_ = G__4791
							switch G__4791 {
							case float64(3):
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "three", Fqn: "three", X_hash: float64(-1651831795)})

							default:
								panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_5265).(string))}))

							}
						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (try (case x 3 :three) (catch js/Error e :fail)) :fail)")}))
				}
			}
			{
				var x_5268 = float64(1)
				_ = x_5268
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4792 = x_5268
					_ = G__4792
					switch G__4792 {
					case float64(1), float64(2), float64(3):
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})

					default:
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x (1 2 3) :ok :fail) :ok)")}))
				}
			}
			{
				var x_5270 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})
				_ = x_5270
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__4793 = x_5270
					_ = G__4793
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), G__4793) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
					} else {
						panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_5270).(string))}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x [:a :b] :ok) :ok)")}))
				}
			}
			{
				var a_5271 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = a_5271
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4794 = a_5271
					_ = G__4794
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__4794) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__4794) {
							return nil
						} else {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})

						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case a nil nil & :amp :none) :none)")}))
				}
			}
			{
				var a_5272 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil})
				_ = a_5272
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4795 = a_5272
					_ = G__4795
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__4795) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__4795) {
							return nil
						} else {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})

						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case a nil nil & :amp :none) :amp)")}))
				}
			}
			{
				var foo_5273 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = foo_5273
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__4796 = foo_5273
					_ = G__4796
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__4796) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__4796) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), G__4796) {
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
							} else {
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})

							}
						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case foo (a b c) :sym :none) :sym)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__4797 = foo_5273
					_ = G__4797
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "d", Str: "d", X_hash: float64(-682293345), X_meta: nil}), G__4797) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__4797) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__4797) {
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
							} else {
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})

							}
						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case foo (b c d) :sym :none) :none)")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF(false, false)) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (compare false false))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF(false, true)) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare false true))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF(true, false)) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (compare true false))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF(float64(0), float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare 0 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF(float64(-1), float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare -1 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF(float64(1), float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (compare 1 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF(float64(1), float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (compare 1 0))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF(float64(1), float64(-1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (compare 1 -1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF("cljs", "cljs")) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (compare \"cljs\" \"cljs\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "cljs", Fqn: "cljs", X_hash: float64(1492417629)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "cljs", Fqn: "cljs", X_hash: float64(1492417629)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (compare :cljs :cljs))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "cljs", Str: "cljs", X_hash: float64(-1162018140), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "cljs", Str: "cljs", X_hash: float64(-1162018140), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (compare (quote cljs) (quote cljs)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF("a", "b")) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare \"a\" \"b\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare :a :b))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare (quote a) (quote b)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: "b", Name: "a", Fqn: "b/a", X_hash: float64(-2123407436)}), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "a", Fqn: "c/a", X_hash: float64(-2123407439)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare :b/a :c/a))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare :c :a/b))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (compare :a/b :c))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: "b", Name: "a", Str: "b/a", X_hash: float64(-482875909), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "a", Str: "c/a", X_hash: float64(-482875912), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare (quote b/a) (quote c/a)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare (quote c) (quote a/b)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (compare (quote a/b) (quote c)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF("a", "c")) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare \"a\" \"c\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare :a :c))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare (quote a) (quote c)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(1)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare [1 2] [1 1 1]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(1)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare [1 2] [1 2 1]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare [1 1] [1 2]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (compare [1 2] [1 2]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (compare [1 2] [1 1]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (compare [1 1 1] [1 2]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(1)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (compare [1 1 2] [1 1 1]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(4)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec))) {
			} else {
				panic((&js.Error{("Assert failed: (= -1 (compare (subvec [1 2 3] 1) (subvec [1 2 4] 1)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec))) {
			} else {
				panic((&js.Error{("Assert failed: (= 0 (compare (subvec [1 2 3] 1) (subvec [1 2 3] 1)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(4)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec))) {
			} else {
				panic((&js.Error{("Assert failed: (= 1 (compare (subvec [1 2 4] 1) (subvec [1 2 3] 1)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(2), float64(1)})).(*cljs_core.CljsCoreList), cljs_core.Reverse.X_invoke_Arity1(cljs_core.Seq.Arity1IQ([]interface{}{float64(1), float64(2), float64(3)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (3 2 1)) (reverse (seq (array 1 2 3))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(2), float64(1)})).(*cljs_core.CljsCoreList), cljs_core.Reverse.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (3 2 1)) (reverse [1 2 3]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(4), float64(3), float64(2), float64(1)})).(*cljs_core.CljsCoreList), cljs_core.Cons.X_invoke_Arity2(float64(4), cljs_core.Reverse.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))).(*cljs_core.CljsCoreCons)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (4 3 2 1)) (cons 4 (reverse [1 2 3])))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Reverse.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})))) {
			} else {
				panic((&js.Error{("Assert failed: (= 6 (reduce + (reverse [1 2 3])))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(4), float64(3), float64(2)})).(*cljs_core.CljsCoreList), cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, cljs_core.Reverse.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (4 3 2)) (map inc (reverse [1 2 3])))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(4), float64(2)})).(*cljs_core.CljsCoreList), cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, cljs_core.Reverse.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (4 2)) (filter even? (reverse [1 2 3 4])))")}))
			}
			{
				var r_5274 = cljs_core.Range_.X_invoke_Arity1(float64(64)).(*cljs_core.CljsCoreRange)
				var v_5275 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, r_5274)
				_, _ = r_5274, v_5275
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_5275)), cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_5275))) {
				} else {
					panic((&js.Error{("Assert failed: (= (hash (seq v)) (hash (seq v)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Array_chunk.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)}).(*cljs_core.CljsCoreArrayChunk))) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (reduce + (array-chunk (array 1 2 3))))")}))
				}
				if cljs_core.Value_(cljs_core.Seq.Arity1IQ(v_5275)).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreChunkedSeq)(nil)).Elem()) {
				} else {
					panic((&js.Error{("Assert failed: (instance? ChunkedSeq (seq v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(r_5274, cljs_core.Seq.Arity1IQ(v_5275)) {
				} else {
					panic((&js.Error{("Assert failed: (= r (seq v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, r_5274).(*cljs_core.CljsCoreLazySeq), cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, v_5275).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (map inc r) (map inc v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, r_5274).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, v_5275).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (filter even? r) (filter even? v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, r_5274).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, v_5275).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (filter odd? r) (filter odd? v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic(r_5274, r_5274, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{r_5274})).(*cljs_core.CljsCoreLazySeq), cljs_core.Concat.X_invoke_ArityVariadic(v_5275, v_5275, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{v_5275})).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (concat r r r) (concat v v v))")}))
				}
				if cljs_core.Value_(cljs_core.Seq.Arity1IQ(v_5275)).Type().Implements(reflect.TypeOf((*cljs_core.CljsCoreIReduce)(nil)).Elem()) {
				} else {
					panic((&js.Error{("Assert failed: (satisfies? IReduce (seq v))")}))
				}
				if float64(2010) == cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_5275)))))).(float64) {
				} else {
					panic((&js.Error{("Assert failed: (== 2010 (reduce + (nnext (nnext (seq v)))))")}))
				}
				if float64(2020) == cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_5275)))))).(float64) {
				} else {
					panic((&js.Error{("Assert failed: (== 2020 (reduce + 10 (nnext (nnext (seq v)))))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(nil)) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next nil))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(cljs_core.Seq.Arity1IQ([]interface{}{float64(1)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next (seq (array 1))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.Seq.Arity1IQ([]interface{}{float64(1), float64(2), float64(3)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (2 3)) (next (seq (array 1 2 3))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(cljs_core.Reverse.X_invoke_Arity1(cljs_core.Seq.Arity1IQ([]interface{}{float64(1)})))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next (reverse (seq (array 1)))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(1)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.Reverse.X_invoke_Arity1(cljs_core.Seq.Arity1IQ([]interface{}{float64(1), float64(2), float64(3)})))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (2 1)) (next (reverse (seq (array 1 2 3)))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(cljs_core.Cons.X_invoke_Arity2(float64(1), nil).(*cljs_core.CljsCoreCons))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next (cons 1 nil)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.Cons.X_invoke_Arity2(float64(1), cljs_core.Cons.X_invoke_Arity2(float64(2), cljs_core.Cons.X_invoke_Arity2(float64(3), nil).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreCons))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (2 3)) (next (cons 1 (cons 2 (cons 3 nil)))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__5276 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5276, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), nil).(*cljs_core.CljsCoreCons)
				})
			}(&cljs_core.AFn{}), nil, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next (lazy-seq (cons 1 nil))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__5277 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5277, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), (&cljs_core.CljsCoreLazySeq{nil, func(G__5278 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__5278, 0, func() interface{} {
							return cljs_core.Cons.X_invoke_Arity2(float64(2), (&cljs_core.CljsCoreLazySeq{nil, func(G__5279 *cljs_core.AFn) *cljs_core.AFn {
								return cljs_core.Fn(G__5279, 0, func() interface{} {
									return cljs_core.Cons.X_invoke_Arity2(float64(3), nil).(*cljs_core.CljsCoreCons)
								})
							}(&cljs_core.AFn{}), nil, nil})).(*cljs_core.CljsCoreCons)
						})
					}(&cljs_core.AFn{}), nil, nil})).(*cljs_core.CljsCoreCons)
				})
			}(&cljs_core.AFn{}), nil, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (2 3)) (next (lazy-seq (cons 1 (lazy-seq (cons 2 (lazy-seq (cons 3 nil))))))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next (list 1)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1)))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (2 3)) (next (list 1 2 3)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next [1]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (2 3)) (next [1 2 3]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(2)).(*cljs_core.CljsCoreRange))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next (range 1 2)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(4)).(*cljs_core.CljsCoreRange))) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (2 3)) (next (range 1 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreUUID{"550e8400-e29b-41d4-a716-446655440000"}), (&cljs_core.CljsCoreUUID{"550e8400-e29b-41d4-a716-446655440000"})) {
			} else {
				panic((&js.Error{("Assert failed: (= (UUID. \"550e8400-e29b-41d4-a716-446655440000\") (UUID. \"550e8400-e29b-41d4-a716-446655440000\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCoreUUID{"550e8400-e29b-41d4-a716-446655440000"}), float64(42)}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCoreUUID{"550e8400-e29b-41d4-a716-446655440000"}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-at-all-found", Fqn: "not-at-all-found", X_hash: float64(-1887167218)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 42 (get {(UUID. \"550e8400-e29b-41d4-a716-446655440000\") 42} (UUID. \"550e8400-e29b-41d4-a716-446655440000\") :not-at-all-found))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-at-all-found", Fqn: "not-at-all-found", X_hash: float64(-1887167218)}), cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCoreUUID{"550e8400-e29b-41d4-a716-446655440000"}), float64(42)}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCoreUUID{"666e8400-e29b-41d4-a716-446655440000"}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-at-all-found", Fqn: "not-at-all-found", X_hash: float64(-1887167218)}))) {
			} else {
				panic((&js.Error{("Assert failed: (= :not-at-all-found (get {(UUID. \"550e8400-e29b-41d4-a716-446655440000\") 42} (UUID. \"666e8400-e29b-41d4-a716-446655440000\") :not-at-all-found))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (cljs.core/into (.-EMPTY cljs.core/PersistentQueue) [1]) (into (.-EMPTY cljs.core/PersistentQueue) [1]))")}))
			}
			if cljs_core.Not_EQ_.Arity2IIB(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (not= (cljs.core/into (.-EMPTY cljs.core/PersistentQueue) [1 2]) (into (.-EMPTY cljs.core/PersistentQueue) [1]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&js.Date{Millis: 1289585655666}), (&js.Date{Millis: 1289585655666})) {
			} else {
				panic((&js.Error{("Assert failed: (= #inst \"2010-11-12T18:14:15.666-00:00\" #inst \"2010-11-12T18:14:15.666-00:00\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}), (&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`})) {
			} else {
				panic((&js.Error{("Assert failed: (= #uuid \"550e8400-e29b-41d4-a716-446655440000\" #uuid \"550e8400-e29b-41d4-a716-446655440000\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}), float64(42)}, nil}), (&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}))) {
			} else {
				panic((&js.Error{("Assert failed: (= 42 (get {#uuid \"550e8400-e29b-41d4-a716-446655440000\" 42} #uuid \"550e8400-e29b-41d4-a716-446655440000\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(string), "1") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str 1) \"1\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(-1)})).(string), "-1") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str -1) \"-1\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{-1.5})).(string), "-1.5") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str -1.5) \"-1.5\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4)}, nil})})).(string), "[3 4]") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str [3 4]) \"[3 4]\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"foo"})).(string), "\"foo\"") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str \"foo\") \"\\\"foo\\\"\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hello", Fqn: "hello", X_hash: float64(-245025397)})})).(string), ":hello") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str :hello) \":hello\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "goodbye", Str: "goodbye", X_hash: float64(-1229580373), X_meta: nil})})).(string), "goodbye") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str (quote goodbye)) \"goodbye\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList)})).(string), "(7 8 9)") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str (quote (7 8 9))) \"(7 8 9)\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "deref", Str: "deref", X_hash: float64(1494944732), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})})).(*cljs_core.CljsCoreList)})).(string), "(deref foo)") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str (quote (deref foo))) \"(deref foo)\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "quote", Str: "quote", X_hash: float64(1377916282), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})})).(*cljs_core.CljsCoreList)})).(string), "(quote bar)") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str (quote (quote bar))) \"(quote bar)\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreSymbol{Ns: "foo", Name: "bar", Str: "foo/bar", X_hash: float64(254379989), X_meta: nil})})).(string), "foo/bar") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str (quote foo/bar)) \"foo/bar\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"a"})).(string), "\"a\"") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str \\a) \"\\\"a\\\"\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: "foo", Name: "bar", Fqn: "foo/bar", X_hash: float64(-1386151538)})})).(string), ":foo/bar") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str :foo/bar) \":foo/bar\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{nil})).(string), "nil") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str nil) \"nil\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{true})).(string), "true") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str true) \"true\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{false})).(string), "false") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str false) \"false\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"string"})).(string), "\"string\"") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str \"string\") \"\\\"string\\\"\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"\u00FC\u00F1\u00EE\u00E7\u00F3\u2202\u00A3", (&cljs_core.CljsCoreKeyword{Ns: "\u0E17\u0E14\u0E2A\u0E2D\u0E1A", Name: "\u4F60\u597D", Fqn: "\u0E17\u0E14\u0E2A\u0E2D\u0E1A/\u4F60\u597D", X_hash: float64(404745050)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "\u3053\u3093\u306B\u3061\u306F", Str: "\u3053\u3093\u306B\u3061\u306F", X_hash: float64(1271843387), X_meta: nil})}, nil})})).(string), "[\"\u00FC\u00F1\u00EE\u00E7\u00F3\u2202\u00A3\" :\u0E17\u0E14\u0E2A\u0E2D\u0E1A/\u4F60\u597D \u3053\u3093\u306B\u3061\u306F]") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str [\"\u00FC\u00F1\u00EE\u00E7\u00F3\u2202\u00A3\" :\u0E17\u0E14\u0E2A\u0E2D\u0E1A/\u4F60\u597D (quote \u3053\u3093\u306B\u3061\u306F)]) \"[\\\"\u00FC\u00F1\u00EE\u00E7\u00F3\u2202\u00A3\\\" :\u0E17\u0E14\u0E2A\u0E2D\u0E1A/\u4F60\u597D \u3053\u3093\u306B\u3061\u306F]\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"escape chars \t \r \n \\ \" \b \f"})).(string), "\"escape chars \\t \\r \\n \\\\ \\\" \\b \\f\"") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str \"escape chars \\t \\r \\n \\\\ \\\" \\b \\f\") \"\\\"escape chars \\\\t \\\\r \\\\n \\\\\\\\ \\\\\\\" \\\\b \\\\f\\\"\")")}))
			}
			{
				X__GT_PrintMe = func(__GT_PrintMe *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_PrintMe, 2, func(a interface{}, b interface{}) interface{} {
						return (&CljsCore_testPrintMe{a, b, nil, nil, nil})
					})
				}(&cljs_core.AFn{})

				Map__GT_PrintMe = func(map__GT_PrintMe *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(map__GT_PrintMe, 1, func(G__4800 interface{}) interface{} {
						return (&CljsCore_testPrintMe{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__4800), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__4800), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__4800, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})})), nil})
					})
				}(&cljs_core.AFn{})

			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&CljsCore_testPrintMe{float64(1), float64(2), nil, nil, nil})})).(string), "#cljs.core-test.PrintMe{:a 1, :b 2}") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str (PrintMe. 1 2)) \"#cljs.core-test.PrintMe{:a 1, :b 2}\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{"2010-11-12T13:14:15.666-05:00"})})).(string), "#inst \"2010-11-12T18:14:15.666-00:00\"") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str (js/Date. \"2010-11-12T13:14:15.666-05:00\")) \"#inst \\\"2010-11-12T18:14:15.666-00:00\\\"\")")}))
			}
			{
				var seq__4817_5280 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(13)).(*cljs_core.CljsCoreRange))
				var chunk__4830_5281 interface{} = nil
				var count__4831_5282 = float64(0)
				var i__4832_5283 = float64(0)
				_, _, _, _ = seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283
				for {
					if i__4832_5283 < count__4831_5282 {
						{
							var month_5284 = chunk__4830_5281.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4832_5283)
							_ = month_5284
							{
								var seq__4833_5285 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
								var chunk__4838_5286 interface{} = nil
								var count__4839_5287 = float64(0)
								var i__4840_5288 = float64(0)
								_, _, _, _ = seq__4833_5285, chunk__4838_5286, count__4839_5287, i__4840_5288
								for {
									if i__4840_5288 < count__4839_5287 {
										{
											var day_5289 = chunk__4838_5286.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4840_5288)
											_ = day_5289
											{
												var seq__4841_5290 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
												var chunk__4842_5291 interface{} = nil
												var count__4843_5292 = float64(0)
												var i__4844_5293 = float64(0)
												_, _, _, _ = seq__4841_5290, chunk__4842_5291, count__4843_5292, i__4844_5293
												for {
													if i__4844_5293 < count__4843_5292 {
														{
															var hour_5294 = chunk__4842_5291.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4844_5293)
															_ = hour_5294
															{
																var pad_5295 = func(G__5297 *cljs_core.AFn, seq__4841_5290 interface{}, chunk__4842_5291 interface{}, count__4843_5292 float64, i__4844_5293 float64, seq__4833_5285 interface{}, chunk__4838_5286 interface{}, count__4839_5287 float64, i__4840_5288 float64, seq__4817_5280 interface{}, chunk__4830_5281 interface{}, count__4831_5282 float64, i__4832_5283 float64, hour_5294 interface{}, day_5289 interface{}, month_5284 interface{}) *cljs_core.AFn {
																	return cljs_core.Fn(G__5297, 1, func(n interface{}) interface{} {
																		if n.(float64) < float64(10) {
																			return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																		} else {
																			return n
																		}
																	})
																}(&cljs_core.AFn{}, seq__4841_5290, chunk__4842_5291, count__4843_5292, i__4844_5293, seq__4833_5285, chunk__4838_5286, count__4839_5287, i__4840_5288, seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283, hour_5294, day_5289, month_5284)
																var inst_5296 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5295.X_invoke_Arity1(month_5284)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5295.X_invoke_Arity1(day_5289)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5295.X_invoke_Arity1(hour_5294)).(string) + ":14:15.666-00:00")
																_, _ = pad_5295, inst_5296
																if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5296})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5296).(string) + "\"")) {
																} else {
																	panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																}
															}
															seq__4841_5290, chunk__4842_5291, count__4843_5292, i__4844_5293 = seq__4841_5290, chunk__4842_5291, count__4843_5292, (i__4844_5293 + float64(1))
															continue
														}
													} else {
														{
															var temp__4222__auto___5298 = cljs_core.Seq.Arity1IQ(seq__4841_5290)
															_ = temp__4222__auto___5298
															if cljs_core.Truth_(temp__4222__auto___5298) {
																{
																	var seq__4841_5299___1 = temp__4222__auto___5298
																	_ = seq__4841_5299___1
																	if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4841_5299___1) {
																		{
																			var c__960__auto___5300 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4841_5299___1)
																			_ = c__960__auto___5300
																			seq__4841_5290, chunk__4842_5291, count__4843_5292, i__4844_5293 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4841_5299___1), c__960__auto___5300, cljs_core.Count.X_invoke_Arity1(c__960__auto___5300).(float64), float64(0)
																			continue
																		}
																	} else {
																		{
																			var hour_5301 = cljs_core.First.X_invoke_Arity1(seq__4841_5299___1)
																			_ = hour_5301
																			{
																				var pad_5302 = func(G__5304 *cljs_core.AFn, seq__4841_5290 interface{}, chunk__4842_5291 interface{}, count__4843_5292 float64, i__4844_5293 float64, seq__4833_5285 interface{}, chunk__4838_5286 interface{}, count__4839_5287 float64, i__4840_5288 float64, seq__4817_5280 interface{}, chunk__4830_5281 interface{}, count__4831_5282 float64, i__4832_5283 float64, hour_5301 interface{}, seq__4841_5299___1 interface{}, temp__4222__auto___5298 cljs_core.CljsCoreISeq, day_5289 interface{}, month_5284 interface{}) *cljs_core.AFn {
																					return cljs_core.Fn(G__5304, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__4841_5290, chunk__4842_5291, count__4843_5292, i__4844_5293, seq__4833_5285, chunk__4838_5286, count__4839_5287, i__4840_5288, seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283, hour_5301, seq__4841_5299___1, temp__4222__auto___5298, day_5289, month_5284)
																				var inst_5303 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5302.X_invoke_Arity1(month_5284)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5302.X_invoke_Arity1(day_5289)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5302.X_invoke_Arity1(hour_5301)).(string) + ":14:15.666-00:00")
																				_, _ = pad_5302, inst_5303
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5303})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5303).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__4841_5290, chunk__4842_5291, count__4843_5292, i__4844_5293 = cljs_core.Next.Arity1IQ(seq__4841_5299___1), nil, float64(0), float64(0)
																			continue
																		}
																	}
																}
															} else {
															}
														}
													}
													break
												}
											}
											seq__4833_5285, chunk__4838_5286, count__4839_5287, i__4840_5288 = seq__4833_5285, chunk__4838_5286, count__4839_5287, (i__4840_5288 + float64(1))
											continue
										}
									} else {
										{
											var temp__4222__auto___5305 = cljs_core.Seq.Arity1IQ(seq__4833_5285)
											_ = temp__4222__auto___5305
											if cljs_core.Truth_(temp__4222__auto___5305) {
												{
													var seq__4833_5306___1 = temp__4222__auto___5305
													_ = seq__4833_5306___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4833_5306___1) {
														{
															var c__960__auto___5307 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4833_5306___1)
															_ = c__960__auto___5307
															seq__4833_5285, chunk__4838_5286, count__4839_5287, i__4840_5288 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4833_5306___1), c__960__auto___5307, cljs_core.Count.X_invoke_Arity1(c__960__auto___5307).(float64), float64(0)
															continue
														}
													} else {
														{
															var day_5308 = cljs_core.First.X_invoke_Arity1(seq__4833_5306___1)
															_ = day_5308
															{
																var seq__4834_5309 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																var chunk__4835_5310 interface{} = nil
																var count__4836_5311 = float64(0)
																var i__4837_5312 = float64(0)
																_, _, _, _ = seq__4834_5309, chunk__4835_5310, count__4836_5311, i__4837_5312
																for {
																	if i__4837_5312 < count__4836_5311 {
																		{
																			var hour_5313 = chunk__4835_5310.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4837_5312)
																			_ = hour_5313
																			{
																				var pad_5314 = func(G__5316 *cljs_core.AFn, seq__4834_5309 interface{}, chunk__4835_5310 interface{}, count__4836_5311 float64, i__4837_5312 float64, seq__4833_5285 interface{}, chunk__4838_5286 interface{}, count__4839_5287 float64, i__4840_5288 float64, seq__4817_5280 interface{}, chunk__4830_5281 interface{}, count__4831_5282 float64, i__4832_5283 float64, hour_5313 interface{}, day_5308 interface{}, seq__4833_5306___1 interface{}, temp__4222__auto___5305 cljs_core.CljsCoreISeq, month_5284 interface{}) *cljs_core.AFn {
																					return cljs_core.Fn(G__5316, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__4834_5309, chunk__4835_5310, count__4836_5311, i__4837_5312, seq__4833_5285, chunk__4838_5286, count__4839_5287, i__4840_5288, seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283, hour_5313, day_5308, seq__4833_5306___1, temp__4222__auto___5305, month_5284)
																				var inst_5315 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5314.X_invoke_Arity1(month_5284)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5314.X_invoke_Arity1(day_5308)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5314.X_invoke_Arity1(hour_5313)).(string) + ":14:15.666-00:00")
																				_, _ = pad_5314, inst_5315
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5315})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5315).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__4834_5309, chunk__4835_5310, count__4836_5311, i__4837_5312 = seq__4834_5309, chunk__4835_5310, count__4836_5311, (i__4837_5312 + float64(1))
																			continue
																		}
																	} else {
																		{
																			var temp__4222__auto___5317___1 = cljs_core.Seq.Arity1IQ(seq__4834_5309)
																			_ = temp__4222__auto___5317___1
																			if cljs_core.Truth_(temp__4222__auto___5317___1) {
																				{
																					var seq__4834_5318___1 = temp__4222__auto___5317___1
																					_ = seq__4834_5318___1
																					if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4834_5318___1) {
																						{
																							var c__960__auto___5319 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4834_5318___1)
																							_ = c__960__auto___5319
																							seq__4834_5309, chunk__4835_5310, count__4836_5311, i__4837_5312 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4834_5318___1), c__960__auto___5319, cljs_core.Count.X_invoke_Arity1(c__960__auto___5319).(float64), float64(0)
																							continue
																						}
																					} else {
																						{
																							var hour_5320 = cljs_core.First.X_invoke_Arity1(seq__4834_5318___1)
																							_ = hour_5320
																							{
																								var pad_5321 = func(G__5323 *cljs_core.AFn, seq__4834_5309 interface{}, chunk__4835_5310 interface{}, count__4836_5311 float64, i__4837_5312 float64, seq__4833_5285 interface{}, chunk__4838_5286 interface{}, count__4839_5287 float64, i__4840_5288 float64, seq__4817_5280 interface{}, chunk__4830_5281 interface{}, count__4831_5282 float64, i__4832_5283 float64, hour_5320 interface{}, seq__4834_5318___1 interface{}, temp__4222__auto___5317___1 cljs_core.CljsCoreISeq, day_5308 interface{}, seq__4833_5306___1 interface{}, temp__4222__auto___5305 cljs_core.CljsCoreISeq, month_5284 interface{}) *cljs_core.AFn {
																									return cljs_core.Fn(G__5323, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__4834_5309, chunk__4835_5310, count__4836_5311, i__4837_5312, seq__4833_5285, chunk__4838_5286, count__4839_5287, i__4840_5288, seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283, hour_5320, seq__4834_5318___1, temp__4222__auto___5317___1, day_5308, seq__4833_5306___1, temp__4222__auto___5305, month_5284)
																								var inst_5322 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5321.X_invoke_Arity1(month_5284)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5321.X_invoke_Arity1(day_5308)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5321.X_invoke_Arity1(hour_5320)).(string) + ":14:15.666-00:00")
																								_, _ = pad_5321, inst_5322
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5322})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5322).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__4834_5309, chunk__4835_5310, count__4836_5311, i__4837_5312 = cljs_core.Next.Arity1IQ(seq__4834_5318___1), nil, float64(0), float64(0)
																							continue
																						}
																					}
																				}
																			} else {
																			}
																		}
																	}
																	break
																}
															}
															seq__4833_5285, chunk__4838_5286, count__4839_5287, i__4840_5288 = cljs_core.Next.Arity1IQ(seq__4833_5306___1), nil, float64(0), float64(0)
															continue
														}
													}
												}
											} else {
											}
										}
									}
									break
								}
							}
							seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283 = seq__4817_5280, chunk__4830_5281, count__4831_5282, (i__4832_5283 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5324 = cljs_core.Seq.Arity1IQ(seq__4817_5280)
							_ = temp__4222__auto___5324
							if cljs_core.Truth_(temp__4222__auto___5324) {
								{
									var seq__4817_5325___1 = temp__4222__auto___5324
									_ = seq__4817_5325___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4817_5325___1) {
										{
											var c__960__auto___5326 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4817_5325___1)
											_ = c__960__auto___5326
											seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4817_5325___1), c__960__auto___5326, cljs_core.Count.X_invoke_Arity1(c__960__auto___5326).(float64), float64(0)
											continue
										}
									} else {
										{
											var month_5327 = cljs_core.First.X_invoke_Arity1(seq__4817_5325___1)
											_ = month_5327
											{
												var seq__4818_5328 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
												var chunk__4823_5329 interface{} = nil
												var count__4824_5330 = float64(0)
												var i__4825_5331 = float64(0)
												_, _, _, _ = seq__4818_5328, chunk__4823_5329, count__4824_5330, i__4825_5331
												for {
													if i__4825_5331 < count__4824_5330 {
														{
															var day_5332 = chunk__4823_5329.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4825_5331)
															_ = day_5332
															{
																var seq__4826_5333 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																var chunk__4827_5334 interface{} = nil
																var count__4828_5335 = float64(0)
																var i__4829_5336 = float64(0)
																_, _, _, _ = seq__4826_5333, chunk__4827_5334, count__4828_5335, i__4829_5336
																for {
																	if i__4829_5336 < count__4828_5335 {
																		{
																			var hour_5337 = chunk__4827_5334.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4829_5336)
																			_ = hour_5337
																			{
																				var pad_5338 = func(G__5340 *cljs_core.AFn, seq__4826_5333 interface{}, chunk__4827_5334 interface{}, count__4828_5335 float64, i__4829_5336 float64, seq__4818_5328 interface{}, chunk__4823_5329 interface{}, count__4824_5330 float64, i__4825_5331 float64, seq__4817_5280 interface{}, chunk__4830_5281 interface{}, count__4831_5282 float64, i__4832_5283 float64, hour_5337 interface{}, day_5332 interface{}, month_5327 interface{}, seq__4817_5325___1 interface{}, temp__4222__auto___5324 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																					return cljs_core.Fn(G__5340, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__4826_5333, chunk__4827_5334, count__4828_5335, i__4829_5336, seq__4818_5328, chunk__4823_5329, count__4824_5330, i__4825_5331, seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283, hour_5337, day_5332, month_5327, seq__4817_5325___1, temp__4222__auto___5324)
																				var inst_5339 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5338.X_invoke_Arity1(month_5327)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5338.X_invoke_Arity1(day_5332)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5338.X_invoke_Arity1(hour_5337)).(string) + ":14:15.666-00:00")
																				_, _ = pad_5338, inst_5339
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5339})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5339).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__4826_5333, chunk__4827_5334, count__4828_5335, i__4829_5336 = seq__4826_5333, chunk__4827_5334, count__4828_5335, (i__4829_5336 + float64(1))
																			continue
																		}
																	} else {
																		{
																			var temp__4222__auto___5341___1 = cljs_core.Seq.Arity1IQ(seq__4826_5333)
																			_ = temp__4222__auto___5341___1
																			if cljs_core.Truth_(temp__4222__auto___5341___1) {
																				{
																					var seq__4826_5342___1 = temp__4222__auto___5341___1
																					_ = seq__4826_5342___1
																					if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4826_5342___1) {
																						{
																							var c__960__auto___5343 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4826_5342___1)
																							_ = c__960__auto___5343
																							seq__4826_5333, chunk__4827_5334, count__4828_5335, i__4829_5336 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4826_5342___1), c__960__auto___5343, cljs_core.Count.X_invoke_Arity1(c__960__auto___5343).(float64), float64(0)
																							continue
																						}
																					} else {
																						{
																							var hour_5344 = cljs_core.First.X_invoke_Arity1(seq__4826_5342___1)
																							_ = hour_5344
																							{
																								var pad_5345 = func(G__5347 *cljs_core.AFn, seq__4826_5333 interface{}, chunk__4827_5334 interface{}, count__4828_5335 float64, i__4829_5336 float64, seq__4818_5328 interface{}, chunk__4823_5329 interface{}, count__4824_5330 float64, i__4825_5331 float64, seq__4817_5280 interface{}, chunk__4830_5281 interface{}, count__4831_5282 float64, i__4832_5283 float64, hour_5344 interface{}, seq__4826_5342___1 interface{}, temp__4222__auto___5341___1 cljs_core.CljsCoreISeq, day_5332 interface{}, month_5327 interface{}, seq__4817_5325___1 interface{}, temp__4222__auto___5324 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																									return cljs_core.Fn(G__5347, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__4826_5333, chunk__4827_5334, count__4828_5335, i__4829_5336, seq__4818_5328, chunk__4823_5329, count__4824_5330, i__4825_5331, seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283, hour_5344, seq__4826_5342___1, temp__4222__auto___5341___1, day_5332, month_5327, seq__4817_5325___1, temp__4222__auto___5324)
																								var inst_5346 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5345.X_invoke_Arity1(month_5327)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5345.X_invoke_Arity1(day_5332)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5345.X_invoke_Arity1(hour_5344)).(string) + ":14:15.666-00:00")
																								_, _ = pad_5345, inst_5346
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5346})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5346).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__4826_5333, chunk__4827_5334, count__4828_5335, i__4829_5336 = cljs_core.Next.Arity1IQ(seq__4826_5342___1), nil, float64(0), float64(0)
																							continue
																						}
																					}
																				}
																			} else {
																			}
																		}
																	}
																	break
																}
															}
															seq__4818_5328, chunk__4823_5329, count__4824_5330, i__4825_5331 = seq__4818_5328, chunk__4823_5329, count__4824_5330, (i__4825_5331 + float64(1))
															continue
														}
													} else {
														{
															var temp__4222__auto___5348___1 = cljs_core.Seq.Arity1IQ(seq__4818_5328)
															_ = temp__4222__auto___5348___1
															if cljs_core.Truth_(temp__4222__auto___5348___1) {
																{
																	var seq__4818_5349___1 = temp__4222__auto___5348___1
																	_ = seq__4818_5349___1
																	if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4818_5349___1) {
																		{
																			var c__960__auto___5350 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4818_5349___1)
																			_ = c__960__auto___5350
																			seq__4818_5328, chunk__4823_5329, count__4824_5330, i__4825_5331 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4818_5349___1), c__960__auto___5350, cljs_core.Count.X_invoke_Arity1(c__960__auto___5350).(float64), float64(0)
																			continue
																		}
																	} else {
																		{
																			var day_5351 = cljs_core.First.X_invoke_Arity1(seq__4818_5349___1)
																			_ = day_5351
																			{
																				var seq__4819_5352 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																				var chunk__4820_5353 interface{} = nil
																				var count__4821_5354 = float64(0)
																				var i__4822_5355 = float64(0)
																				_, _, _, _ = seq__4819_5352, chunk__4820_5353, count__4821_5354, i__4822_5355
																				for {
																					if i__4822_5355 < count__4821_5354 {
																						{
																							var hour_5356 = chunk__4820_5353.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4822_5355)
																							_ = hour_5356
																							{
																								var pad_5357 = func(G__5359 *cljs_core.AFn, seq__4819_5352 interface{}, chunk__4820_5353 interface{}, count__4821_5354 float64, i__4822_5355 float64, seq__4818_5328 interface{}, chunk__4823_5329 interface{}, count__4824_5330 float64, i__4825_5331 float64, seq__4817_5280 interface{}, chunk__4830_5281 interface{}, count__4831_5282 float64, i__4832_5283 float64, hour_5356 interface{}, day_5351 interface{}, seq__4818_5349___1 interface{}, temp__4222__auto___5348___1 cljs_core.CljsCoreISeq, month_5327 interface{}, seq__4817_5325___1 interface{}, temp__4222__auto___5324 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																									return cljs_core.Fn(G__5359, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__4819_5352, chunk__4820_5353, count__4821_5354, i__4822_5355, seq__4818_5328, chunk__4823_5329, count__4824_5330, i__4825_5331, seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283, hour_5356, day_5351, seq__4818_5349___1, temp__4222__auto___5348___1, month_5327, seq__4817_5325___1, temp__4222__auto___5324)
																								var inst_5358 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5357.X_invoke_Arity1(month_5327)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5357.X_invoke_Arity1(day_5351)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5357.X_invoke_Arity1(hour_5356)).(string) + ":14:15.666-00:00")
																								_, _ = pad_5357, inst_5358
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5358})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5358).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__4819_5352, chunk__4820_5353, count__4821_5354, i__4822_5355 = seq__4819_5352, chunk__4820_5353, count__4821_5354, (i__4822_5355 + float64(1))
																							continue
																						}
																					} else {
																						{
																							var temp__4222__auto___5360___2 = cljs_core.Seq.Arity1IQ(seq__4819_5352)
																							_ = temp__4222__auto___5360___2
																							if cljs_core.Truth_(temp__4222__auto___5360___2) {
																								{
																									var seq__4819_5361___1 = temp__4222__auto___5360___2
																									_ = seq__4819_5361___1
																									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4819_5361___1) {
																										{
																											var c__960__auto___5362 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4819_5361___1)
																											_ = c__960__auto___5362
																											seq__4819_5352, chunk__4820_5353, count__4821_5354, i__4822_5355 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4819_5361___1), c__960__auto___5362, cljs_core.Count.X_invoke_Arity1(c__960__auto___5362).(float64), float64(0)
																											continue
																										}
																									} else {
																										{
																											var hour_5363 = cljs_core.First.X_invoke_Arity1(seq__4819_5361___1)
																											_ = hour_5363
																											{
																												var pad_5364 = func(G__5366 *cljs_core.AFn, seq__4819_5352 interface{}, chunk__4820_5353 interface{}, count__4821_5354 float64, i__4822_5355 float64, seq__4818_5328 interface{}, chunk__4823_5329 interface{}, count__4824_5330 float64, i__4825_5331 float64, seq__4817_5280 interface{}, chunk__4830_5281 interface{}, count__4831_5282 float64, i__4832_5283 float64, hour_5363 interface{}, seq__4819_5361___1 interface{}, temp__4222__auto___5360___2 cljs_core.CljsCoreISeq, day_5351 interface{}, seq__4818_5349___1 interface{}, temp__4222__auto___5348___1 cljs_core.CljsCoreISeq, month_5327 interface{}, seq__4817_5325___1 interface{}, temp__4222__auto___5324 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																													return cljs_core.Fn(G__5366, 1, func(n interface{}) interface{} {
																														if n.(float64) < float64(10) {
																															return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																														} else {
																															return n
																														}
																													})
																												}(&cljs_core.AFn{}, seq__4819_5352, chunk__4820_5353, count__4821_5354, i__4822_5355, seq__4818_5328, chunk__4823_5329, count__4824_5330, i__4825_5331, seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283, hour_5363, seq__4819_5361___1, temp__4222__auto___5360___2, day_5351, seq__4818_5349___1, temp__4222__auto___5348___1, month_5327, seq__4817_5325___1, temp__4222__auto___5324)
																												var inst_5365 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5364.X_invoke_Arity1(month_5327)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5364.X_invoke_Arity1(day_5351)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5364.X_invoke_Arity1(hour_5363)).(string) + ":14:15.666-00:00")
																												_, _ = pad_5364, inst_5365
																												if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5365})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5365).(string) + "\"")) {
																												} else {
																													panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																												}
																											}
																											seq__4819_5352, chunk__4820_5353, count__4821_5354, i__4822_5355 = cljs_core.Next.Arity1IQ(seq__4819_5361___1), nil, float64(0), float64(0)
																											continue
																										}
																									}
																								}
																							} else {
																							}
																						}
																					}
																					break
																				}
																			}
																			seq__4818_5328, chunk__4823_5329, count__4824_5330, i__4825_5331 = cljs_core.Next.Arity1IQ(seq__4818_5349___1), nil, float64(0), float64(0)
																			continue
																		}
																	}
																}
															} else {
															}
														}
													}
													break
												}
											}
											seq__4817_5280, chunk__4830_5281, count__4831_5282, i__4832_5283 = cljs_core.Next.Arity1IQ(seq__4817_5325___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
							}
						}
					}
					break
				}
			}
			{
				var uuid_str_5367 = "550e8400-e29b-41d4-a716-446655440000"
				var uuid_5368 = (&cljs_core.CljsCoreUUID{uuid_str_5367})
				_, _ = uuid_str_5367, uuid_5368
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{uuid_5368})).(string), ("#uuid \"" + cljs_core.Str.X_invoke_Arity1(uuid_str_5367).(string) + "\"")) {
				} else {
					panic((&js.Error{("Assert failed: (= (pr-str uuid) (str \"#uuid \\\"\" uuid-str \"\\\"\"))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest.Arity1IQ(cljs_core.Conj.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentQueue_EMPTY, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})))})).(string), "(2 3)") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str (rest (conj (.-EMPTY cljs.core/PersistentQueue) 1 2 3))) \"(2 3)\")")}))
			}
			X_bar = func(_bar *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(_bar, 2, func(this interface{}, x interface{}) interface{} {
					return this.(CljsCore_testIBar).X_bar_Arity2(x)
				})
			}(&cljs_core.AFn{})

			Baz = func(baz *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(baz, 1, func(f interface{}) interface{} {
					X__GT_t4850 = func(__GT_t4850 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t4850, 4, func(f___1 interface{}, baz___1 interface{}, test_stuff___1 interface{}, meta4851 interface{}) interface{} {
							return (&CljsCore_testT4850{f___1, baz___1, test_stuff___1, meta4851})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT4850{f, baz, test_stuff, nil})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(float64(2), Baz.X_invoke_Arity1(cljs_core.Inc).(*CljsCore_testT4850).X_bar_Arity2(float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (-bar (baz inc) 1))")}))
			}
			{
				var x_5369 = "original"
				_ = x_5369
				Original_closure_stmt = func(original_closure_stmt *cljs_core.AFn, x_5369 string) *cljs_core.AFn {
					return cljs_core.Fn(original_closure_stmt, 0, func() interface{} {
						return x_5369
					})
				}(&cljs_core.AFn{}, x_5369)

			}
			{
				var x_5370 = "overwritten"
				_ = x_5370
				if cljs_core.X_EQ_.Arity2IIB("original", Original_closure_stmt.X_invoke_Arity0().(string)) {
				} else {
					panic((&js.Error{("Assert failed: (= \"original\" (original-closure-stmt))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB("original", func() string {
				var x = "original"
				var oce = func(G__5371 *cljs_core.AFn, x string) *cljs_core.AFn {
					return cljs_core.Fn(G__5371, 0, func() interface{} {
						return x
					})
				}(&cljs_core.AFn{}, x)
				var x___1 = "overwritten"
				_, _, _ = x, oce, x___1
				return oce.X_invoke_Arity0().(string)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= \"original\" (let [x \"original\" oce (fn [] x) x \"overwritten\"] (oce)))")}))
			}
			{
				var x, y *cljs_core.AFn
				x = func(x *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(x, 0, func() interface{} {
						return "original"
					})
				}(&cljs_core.AFn{})
				y = func(y *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(y, 0, func() interface{} {
						return x.X_invoke_Arity0().(string)
					})
				}(&cljs_core.AFn{})
				_, _ = x, y
				{
					var x_5372___1 = func(G__5373 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__5373, 0, func() interface{} {
							return "overwritten"
						})
					}(&cljs_core.AFn{})
					_ = x_5372___1
					if cljs_core.X_EQ_.Arity2IIB("original", y.X_invoke_Arity0().(string)) {
					} else {
						panic((&js.Error{("Assert failed: (= \"original\" (y))")}))
					}
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reduce_kv.X_invoke_Arity3(cljs_core.Conj, cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}))), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (reduce-kv conj [] (sorted-map :foo 1 :bar 2)) [:bar 2 :foo 1])")}))
			}
			{
				var kvr_test *cljs_core.AFn
				kvr_test = func(kvr_test *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(kvr_test, 2, func(data interface{}, expect interface{}) interface{} {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)}), cljs_core.Reduce_kv.X_invoke_Arity3(func(G__5374 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__5374, 3, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
								return cljs_core.Reduced.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)})).(*cljs_core.CljsCoreReduced)
							})
						}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentVector_EMPTY, data)) {
						} else {
							panic((&js.Error{("Assert failed: (= :reduced (reduce-kv (fn [_ _ _] (reduced :reduced)) [] data))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sort.X_invoke_Arity1(expect), cljs_core.Sort.X_invoke_Arity1(cljs_core.Reduce_kv.X_invoke_Arity3(func(G__5375 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__5375, 3, func(r interface{}, k interface{}, v interface{}) interface{} {
								return cljs_core.Conj.X_invoke_Arity2(r, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{k, v}, nil}))
							})
						}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentVector_EMPTY, data))) {
							return nil
						} else {
							panic((&js.Error{("Assert failed: (= (sort expect) (sort (reduce-kv (fn [r k v] (-> r (conj [k v]))) [] data)))")}))
						}
					})
				}(&cljs_core.AFn{})
				_ = kvr_test
				kvr_test.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "k0", Fqn: "k0", X_hash: float64(218250735)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "k1", Fqn: "k1", X_hash: float64(952658428)})}, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v0", Fqn: "v0", X_hash: float64(-895937461)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v1", Fqn: "v1", X_hash: float64(513124261)})}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "k0", Fqn: "k0", X_hash: float64(218250735)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v0", Fqn: "v0", X_hash: float64(-895937461)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "k1", Fqn: "k1", X_hash: float64(952658428)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v1", Fqn: "v1", X_hash: float64(513124261)})}, nil})}, nil}))
				kvr_test.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "k0", Fqn: "k0", X_hash: float64(218250735)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v0", Fqn: "v0", X_hash: float64(-895937461)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "k1", Fqn: "k1", X_hash: float64(952658428)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v1", Fqn: "v1", X_hash: float64(513124261)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "k0", Fqn: "k0", X_hash: float64(218250735)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v0", Fqn: "v0", X_hash: float64(-895937461)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "k1", Fqn: "k1", X_hash: float64(952658428)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v1", Fqn: "v1", X_hash: float64(513124261)})}, nil})}, nil}))
				kvr_test.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v0", Fqn: "v0", X_hash: float64(-895937461)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v1", Fqn: "v1", X_hash: float64(513124261)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v0", Fqn: "v0", X_hash: float64(-895937461)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "v1", Fqn: "v1", X_hash: float64(513124261)})}, nil})}, nil}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "init", Fqn: "init", X_hash: float64(-1875481434)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "val", Fqn: "val", X_hash: float64(128701612)})}, nil}), cljs_core.Reduce_kv.X_invoke_Arity3(cljs_core.Assoc, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "init", Fqn: "init", X_hash: float64(-1875481434)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "val", Fqn: "val", X_hash: float64(128701612)})}, nil}), nil)) {
			} else {
				panic((&js.Error{("Assert failed: (= {:init :val} (reduce-kv assoc {:init :val} nil))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), func() (return__5376 interface{}) {
				defer func() {
					if e4855 := recover(); e4855 != nil {
						if cljs_core.Value_(e4855).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreExceptionInfo)(nil)).Elem()) {
							{
								var e = e4855
								_ = e
								return__5376 = cljs_core.Ex_data.X_invoke_Arity1(e)
							}
						} else {
							panic(e4855)

						}
					}
				}()
				{
					panic(cljs_core.Ex_info.X_invoke_Arity2("asdf", (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(*cljs_core.CljsCoreExceptionInfo))
				}
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= {:foo 1} (try (throw (ex-info \"asdf\" {:foo 1})) (catch ExceptionInfo e (ex-data e))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Assoc.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentArrayMap_EMPTY, float64(154618822656), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(261993005056), float64(1)})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(154618822656), float64(1), float64(261993005056), float64(1)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (assoc {} 154618822656 1 261993005056 1) {154618822656 1, 261993005056 1})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(1)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nothing-there", Fqn: "nothing-there", X_hash: float64(1302692654)})), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nothing-there", Fqn: "nothing-there", X_hash: float64(1302692654)})) {
			} else {
				panic((&js.Error{("Assert failed: (= (get-in {:a {:b 1}} [:a :b :c] :nothing-there) :nothing-there)")}))
			}
			if cljs_core.Nil_(cljs_core.Get_in.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (get-in {:foo {:bar 2}} [:foo :bar :baz]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "tag", Fqn: "tag", X_hash: float64(-1290361223)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "int", Str: "int", X_hash: float64(-100885395), X_meta: nil})}, nil}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "tag", Fqn: "tag", X_hash: float64(-1290361223)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "int", Str: "int", X_hash: float64(-100885395), X_meta: nil})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (meta (with-meta (quote foo) {:tag (quote int)})) {:tag (quote int)})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reduce_kv.X_invoke_Arity3(cljs_core.X_PLUS_, float64(0), cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, cljs_core.Range_.X_invoke_Arity1(float64(1000)).(*cljs_core.CljsCoreRange))), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Range_.X_invoke_Arity1(float64(1000)).(*cljs_core.CljsCoreRange))) {
			} else {
				panic((&js.Error{("Assert failed: (= (reduce-kv + 0 (apply hash-map (range 1000))) (reduce + (range 1000)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(1), float64(2)}, nil}), func(G__5377 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5377, 0, func(more__ ...interface{}) interface{} {
					var more = cljs_core.Seq.Arity1IQ(more__[0])
					_ = more
					return more
				})
			}(&cljs_core.AFn{}).X_invoke_Arity3(js.Undefined, float64(1), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= [js/undefined 1 2] ((fn [& more] more) js/undefined 1 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(4), float64(5)}, nil}), func(G__5378 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5378, 2, func(a_b_more__ ...interface{}) interface{} {
					var a = a_b_more__[0]
					var b = a_b_more__[1]
					var more = cljs_core.Seq.Arity1IQ(a_b_more__[2])
					_, _, _ = a, b, more
					return more
				})
			}(&cljs_core.AFn{}).X_invoke_Arity5(float64(1), float64(2), js.Undefined, float64(4), float64(5))) {
			} else {
				panic((&js.Error{("Assert failed: (= [js/undefined 4 5] ((fn [a b & more] more) 1 2 js/undefined 4 5))")}))
			}
			if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(float64(42), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "anything", Fqn: "anything", X_hash: float64(285898971)}))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (get 42 :anything))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity3(float64(42), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "anything", Fqn: "anything", X_hash: float64(285898971)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-found", Fqn: "not-found", X_hash: float64(-629079980)})), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-found", Fqn: "not-found", X_hash: float64(-629079980)})) {
			} else {
				panic((&js.Error{("Assert failed: (= (get 42 :anything :not-found) :not-found)")}))
			}
			if cljs_core.Nil_(cljs_core.First.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity3(cljs_core.Get, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "anything", Fqn: "anything", X_hash: float64(285898971)})}, nil})).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (first (map get [42] [:anything])))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.First.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity4(cljs_core.Get, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "anything", Fqn: "anything", X_hash: float64(285898971)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-found", Fqn: "not-found", X_hash: float64(-629079980)})}, nil})).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-found", Fqn: "not-found", X_hash: float64(-629079980)})) {
			} else {
				panic((&js.Error{("Assert failed: (= (first (map get [42] [:anything] [:not-found])) :not-found)")}))
			}
			{
				var fs_5379 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = fs_5379
				{
					var seq__4856_5380 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity1(float64(4)).(*cljs_core.CljsCoreRange))
					var chunk__4858_5381 interface{} = nil
					var count__4859_5382 = float64(0)
					var i__4860_5383 = float64(0)
					_, _, _, _ = seq__4856_5380, chunk__4858_5381, count__4859_5382, i__4860_5383
					for {
						if i__4860_5383 < count__4859_5382 {
							{
								var x_5384 = chunk__4858_5381.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4860_5383)
								_ = x_5384
								{
									var y_5385 = (x_5384.(float64) + float64(1))
									var f_5386 = func(G__5387 *cljs_core.AFn, seq__4856_5380 interface{}, chunk__4858_5381 interface{}, count__4859_5382 float64, i__4860_5383 float64, y_5385 float64, x_5384 interface{}, fs_5379 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
										return cljs_core.Fn(G__5387, 0, func() interface{} {
											return y_5385
										})
									}(&cljs_core.AFn{}, seq__4856_5380, chunk__4858_5381, count__4859_5382, i__4860_5383, y_5385, x_5384, fs_5379)
									_, _ = y_5385, f_5386
									cljs_core.Swap_BANG_.X_invoke_Arity3(fs_5379, cljs_core.Conj, f_5386)
								}
								seq__4856_5380, chunk__4858_5381, count__4859_5382, i__4860_5383 = seq__4856_5380, chunk__4858_5381, count__4859_5382, (i__4860_5383 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___5388 = cljs_core.Seq.Arity1IQ(seq__4856_5380)
								_ = temp__4222__auto___5388
								if cljs_core.Truth_(temp__4222__auto___5388) {
									{
										var seq__4856_5389___1 = temp__4222__auto___5388
										_ = seq__4856_5389___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4856_5389___1) {
											{
												var c__960__auto___5390 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4856_5389___1)
												_ = c__960__auto___5390
												seq__4856_5380, chunk__4858_5381, count__4859_5382, i__4860_5383 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4856_5389___1), c__960__auto___5390, cljs_core.Count.X_invoke_Arity1(c__960__auto___5390).(float64), float64(0)
												continue
											}
										} else {
											{
												var x_5391 = cljs_core.First.X_invoke_Arity1(seq__4856_5389___1)
												_ = x_5391
												{
													var y_5392 = (x_5391.(float64) + float64(1))
													var f_5393 = func(G__5394 *cljs_core.AFn, seq__4856_5380 interface{}, chunk__4858_5381 interface{}, count__4859_5382 float64, i__4860_5383 float64, y_5392 float64, x_5391 interface{}, seq__4856_5389___1 interface{}, temp__4222__auto___5388 cljs_core.CljsCoreISeq, fs_5379 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
														return cljs_core.Fn(G__5394, 0, func() interface{} {
															return y_5392
														})
													}(&cljs_core.AFn{}, seq__4856_5380, chunk__4858_5381, count__4859_5382, i__4860_5383, y_5392, x_5391, seq__4856_5389___1, temp__4222__auto___5388, fs_5379)
													_, _ = y_5392, f_5393
													cljs_core.Swap_BANG_.X_invoke_Arity3(fs_5379, cljs_core.Conj, f_5393)
												}
												seq__4856_5380, chunk__4858_5381, count__4859_5382, i__4860_5383 = cljs_core.Next.Arity1IQ(seq__4856_5389___1), nil, float64(0), float64(0)
												continue
											}
										}
									}
								} else {
								}
							}
						}
						break
					}
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__5395 *cljs_core.AFn, fs_5379 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__5395, 1, func(p1__4099_SHARP_ interface{}) interface{} {
						{
							return p1__4099_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
						}
					})
				}(&cljs_core.AFn{}, fs_5379), cljs_core.Deref.X_invoke_Arity1(fs_5379)).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__4099#] (p1__4099#)) (clojure.core/deref fs)) (quote (1 2 3 4)))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Char.X_invoke_Arity1(float64(65)), "A") {
			} else {
				panic((&js.Error{("Assert failed: (= (char 65) \\A)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Char.X_invoke_Arity1("A"), "A") {
			} else {
				panic((&js.Error{("Assert failed: (= (char \\A) \\A)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})), cljs_core.Hash.X_invoke_Arity1(cljs_core.Symbol.X_invoke_Arity1("foo"))) {
			} else {
				panic((&js.Error{("Assert failed: (= (hash (quote foo)) (hash (symbol \"foo\")))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: "foo", Name: "bar", Str: "foo/bar", X_hash: float64(254379989), X_meta: nil})), cljs_core.Hash.X_invoke_Arity1(cljs_core.Symbol.X_invoke_Arity2("foo", "bar").(*cljs_core.CljsCoreSymbol))) {
			} else {
				panic((&js.Error{("Assert failed: (= (hash (quote foo/bar)) (hash (symbol \"foo\" \"bar\")))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCoreLazySeq{nil, func(G__5396 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5396, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCoreLazySeq{nil, func(G__5397 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5397, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreLazySeq{nil, func(G__5398 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5398, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil})})).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (lazy-cat [1] [2] [3]) (quote (1 2 3)))")}))
			}
			X__GT_PositionalFactoryTest = func(__GT_PositionalFactoryTest *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_PositionalFactoryTest, 1, func(x interface{}) interface{} {
					return (&CljsCore_testPositionalFactoryTest{x})
				})
			}(&cljs_core.AFn{})

			if float64(1) == X__GT_PositionalFactoryTest.X_invoke_Arity1(float64(1)).(*CljsCore_testPositionalFactoryTest).X.(float64) {
			} else {
				panic((&js.Error{("Assert failed: (== 1 (.-x (->PositionalFactoryTest 1)))")}))
			}
			if cljs_core.Nil_((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "test", Fqn: "test", X_hash: float64(577538877)}).X_invoke_Arity1("test")) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (:test \"test\"))")}))
			}
			{
				var f_BANG_, g_BANG_ *cljs_core.AFn
				f_BANG_ = func(f_BANG_ *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(f_BANG_, 1, func(x interface{}) interface{} {
						cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"f"}))
						return x
					})
				}(&cljs_core.AFn{})
				g_BANG_ = func(g_BANG_ *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(g_BANG_, 1, func(x interface{}) interface{} {
						cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"g"}))
						return x
					})
				}(&cljs_core.AFn{})
				_, _ = f_BANG_, g_BANG_
				if cljs_core.X_EQ_.Arity2IIB("ffgfg", func() string {
					var sb__1130__auto__ = (&goog_string.StringBuffer{})
					_ = sb__1130__auto__
					{
						var _STAR_print_fn_STAR_4862_5399 = cljs_core.X_STAR_print_fn_STAR_
						_ = _STAR_print_fn_STAR_4862_5399
						func() {
							defer func() {
								cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_4862_5399

							}()
							{
								cljs_core.X_STAR_print_fn_STAR_ = func(G__5400 *cljs_core.AFn, _STAR_print_fn_STAR_4862_5399 interface{}, sb__1130__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
									return cljs_core.Fn(G__5400, 1, func(x__1131__auto__ interface{}) interface{} {
										return sb__1130__auto__.Append(x__1131__auto__)
									})
								}(&cljs_core.AFn{}, _STAR_print_fn_STAR_4862_5399, sb__1130__auto__)

								cljs_core.Value_(f_BANG_.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil}))).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreSymbol)(nil)).Elem())
								func(x, y float64) float64 {
									if x > y {
										return x
									} else {
										return y
									}
								}(f_BANG_.X_invoke_Arity1(float64(5)).(float64), g_BANG_.X_invoke_Arity1(float64(10)).(float64))
								func(x, y float64) float64 {
									if x < y {
										return x
									} else {
										return y
									}
								}(f_BANG_.X_invoke_Arity1(float64(5)).(float64), g_BANG_.X_invoke_Arity1(float64(10)).(float64))
							}
						}()
					}
					return (`` + cljs_core.Str.X_invoke_Arity1(sb__1130__auto__).(string))
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ffgfg\" (with-out-str (instance? Symbol (f! (quote foo))) (max (f! 5) (g! 10)) (min (f! 5) (g! 10))))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), cljs_core.Set.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 2} (set [1 2 2]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{float64(1), float64(2), float64(2)}, true)) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 2} (hash-set 1 2 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_set, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 2} (apply hash-set [1 2 2]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Last.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(cljs_core.Identity, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Range_.X_invoke_Arity1(float64(32)).(*cljs_core.CljsCoreRange))).(*cljs_core.CljsCoreLazySeq)), float64(31)) {
			} else {
				panic((&js.Error{("Assert failed: (= (last (map identity (into [] (range 32)))) 31)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashSet_EMPTY, cljs_core.Range_.X_invoke_Arity1(float64(32)).(*cljs_core.CljsCoreRange)), cljs_core.Set.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(cljs_core.Identity, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Range_.X_invoke_Arity1(float64(32)).(*cljs_core.CljsCoreRange))).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{("Assert failed: (= (into #{} (range 32)) (set (map identity (into [] (range 32)))))")}))
			}

			Foo580 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__5401 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5401, 0, func() interface{} {
					return nil
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), func(G__5402 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5402, 0, func() interface{} {
					{
						var G__4863 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
						_ = G__4863
						return Foo580.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4863)
					}
				})
			}(&cljs_core.AFn{})}, nil})

			if cljs_core.Nil_((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(Foo580).(cljs_core.CljsCoreIFn).X_invoke_Arity0().(cljs_core.CljsCoreIFn).X_invoke_Arity0()) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (((:b foo580))))")}))
			}
			if (cljs_core.First.X_invoke_Arity1(cljs_core.Filter.X_invoke_Arity2(func(G__5403 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5403, 1, func(p1__4100_SHARP_ interface{}) interface{} {
					return (p1__4100_SHARP_.(float64) == float64(9999))
				})
			}(&cljs_core.AFn{}), cljs_core.Range_.X_invoke_Arity0().(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)).(float64) == float64(9999)) {
			} else {
				panic((&js.Error{("Assert failed: (== (first (filter (fn* [p1__4100#] (== p1__4100# 9999)) (range))) 9999)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Concat.X_invoke_Arity2(nil, cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= () (concat nil []))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Concat.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= () (concat [] []))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("foobar", cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Concat.X_invoke_Arity2("foo", "bar").(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"foobar\" (apply str (concat \"foo\" \"bar\")))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{""})).(*cljs_core.CljsCoreList), cljs_core.Re_seq.X_invoke_Arity2((&js.RegExp{Pattern: `\s*`, Flags: ``}), "")) {
			} else {
				panic((&js.Error{("Assert failed: (= (quote (\"\")) (re-seq #\"\\s*\" \"\"))")}))
			}
			X__GT_KeywordTest = func(__GT_KeywordTest *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_KeywordTest, 0, func() interface{} {
					return (&CljsCore_testKeywordTest{})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1((&CljsCore_testKeywordTest{})), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nothing", Fqn: "nothing", X_hash: float64(-1022703296)})) {
			} else {
				panic((&js.Error{("Assert failed: (= (:a (KeywordTest.)) :nothing)")}))
			}
			{
				var a_5404 = func() *CljsCore_testT4864 {
					X__GT_t4864 = func(__GT_t4864 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t4864, 2, func(test_stuff___1 interface{}, meta4865 interface{}) interface{} {
							return (&CljsCore_testT4864{test_stuff___1, meta4865})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT4864{test_stuff, nil})
				}()
				var b_5405 = func() *CljsCore_testT4867 {
					X__GT_t4867 = func(__GT_t4867 *cljs_core.AFn, a_5404 *CljsCore_testT4864) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t4867, 3, func(a___1 interface{}, test_stuff___1 interface{}, meta4868 interface{}) interface{} {
							return (&CljsCore_testT4867{a___1, test_stuff___1, meta4868})
						})
					}(&cljs_core.AFn{}, a_5404)

					return (&CljsCore_testT4867{a_5404, test_stuff, nil})
				}()
				var s_5406 = cljs_core.Set.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(128)).(*cljs_core.CljsCoreRange))
				_, _, _ = a_5404, b_5405, s_5406
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_5406, a_5404, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_5405}))), a_5404)), a_5404), cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_5406, a_5404, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_5405}))), a_5404)), a_5404)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-> (conj s a b) transient (disj! a) persistent! (conj a)) (-> (conj s a b) transient (disj! a) persistent! (conj a)))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Juxt.X_invoke_Arity2(cljs_core.Namespace, cljs_core.Name).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1(cljs_core.Keyword.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a.b", Str: "a.b", X_hash: float64(-2098083151), X_meta: nil}))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "a.b"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (-> (quote a.b) keyword ((juxt namespace name))) [nil \"a.b\"])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Juxt.X_invoke_Arity2(cljs_core.Namespace, cljs_core.Name).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1(cljs_core.Keyword.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: "a.b", Name: "c", Str: "a.b/c", X_hash: float64(-122574001), X_meta: nil}))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a.b", "c"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (-> (quote a.b/c) keyword ((juxt namespace name))) [\"a.b\" \"c\"])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Juxt.X_invoke_Arity2(cljs_core.Namespace, cljs_core.Name).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1(cljs_core.Keyword.X_invoke_Arity1("a.b")), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "a.b"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (-> \"a.b\" keyword ((juxt namespace name))) [nil \"a.b\"])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Juxt.X_invoke_Arity2(cljs_core.Namespace, cljs_core.Name).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1(cljs_core.Keyword.X_invoke_Arity1("a.b/c")), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a.b", "c"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (-> \"a.b/c\" keyword ((juxt namespace name))) [\"a.b\" \"c\"])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Keyword.X_invoke_Arity1(float64(123)), nil) {
			} else {
				panic((&js.Error{("Assert failed: (= (keyword 123) nil)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Keyword.X_invoke_Arity1((&js.Date{})), nil) {
			} else {
				panic((&js.Error{("Assert failed: (= (keyword (js/Date.)) nil)")}))
			}
			Some_x = float64(1)

			Some_y = float64(1)

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{Some_y, Some_x}, true).(*cljs_core.CljsCorePersistentHashSet)).(float64), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count #{some-y some-x}) 1)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{Some_x, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), Some_y, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, true, false).(*cljs_core.CljsCorePersistentArrayMap)).(float64), float64(1)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count {some-x :foo, some-y :bar}) 1)")}))
			}
			if cljs_core.Value_([]interface{}{float64(1), float64(2), float64(3)}).Kind() == reflect.Slice {
			} else {
				panic((&js.Error{("Assert failed: (array? #js [1 2 3])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(len([]interface{}{float64(1), float64(2), float64(3)})), float64(3)) {
			} else {
				panic((&js.Error{("Assert failed: (= (alength #js [1 2 3]) 3)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ([]interface{}{float64(1), float64(2), float64(3)}), cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (seq #js [1 2 3]) (seq [1 2 3]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Vector, cljs_core.Drop_while.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.X_EQ_, float64(1)).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply vector (drop-while (partial = 1) [1 2 3])) [2 3])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.List, cljs_core.Drop_while.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.X_EQ_, float64(1)).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (apply list (drop-while (partial = 1) [1 2 3])) (quote (2 3)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Drop.X_invoke_Arity2(float64(1), []interface{}{float64(1), float64(2), float64(3)}).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(3), nil, float64(2), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (set (drop 1 #js [1 2 3])) #{3 2})")}))
			}
			if cljs_core.Nil_(cljs_core.First.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(cljs_core.Rest.Arity1IQ(cljs_core.Rest.Arity1IQ(cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange)))))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (first (rest (rest (rest (range 3))))))")}))
			}
			if cljs_core.Count.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(0)}, nil})}, true)).(float64) == cljs_core.Count.X_invoke_Arity1(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(0)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(0)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(0)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(2)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(4)}, nil}))).(float64) {
			} else {
				panic((&js.Error{("Assert failed: (== (count (hash-set [1 4] [2 4] [3 4] [0 3] [1 3] [2 3] [3 3] [0 2] [1 2] [2 2] [3 2] [4 2] [0 1] [1 1] [2 1] [3 1] [1 0] [2 0] [3 0])) (count (list [1 4] [2 4] [3 4] [0 3] [1 3] [2 3] [3 3] [0 2] [1 2] [2 2] [3 2] [4 2] [0 1] [1 1] [2 1] [3 1] [1 0] [2 0] [3 0])))")}))
			}
			X_woz = func(_woz *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(_woz, 1, func(this interface{}) interface{} {
					return this.(CljsCore_testIWoz).X_woz_Arity1()
				})
			}(&cljs_core.AFn{})

			Noz = cljs_core.CljsCorePersistentVector_EMPTY

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_ArityVariadic(cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)}))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (-> (transient []) (conj! 1 2) persistent!) [1 2])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_ArityVariadic(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil})), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)}))), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(3), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (-> (transient #{1 3 2}) (disj! 1 2) persistent!) #{3})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_ArityVariadic(cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (-> (transient {}) (assoc! :a 1 :b 2) persistent!) {:a 1, :b 2})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Dissoc_BANG_.X_invoke_ArityVariadic(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(3)}, nil})), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(3)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (-> (transient {:a 1, :b 2, :c 3}) (dissoc! :a :b) persistent!) {:c 3})")}))
			}
			{
				var seq__4870_5407 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__4871_5408 interface{} = nil
				var count__4872_5409 = float64(0)
				var i__4873_5410 = float64(0)
				_, _, _, _ = seq__4870_5407, chunk__4871_5408, count__4872_5409, i__4873_5410
				for {
					if i__4873_5410 < count__4872_5409 {
						{
							var n_5411 = chunk__4871_5408.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4873_5410)
							_ = n_5411
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5412 interface{}) {
								defer func() {
									if e4876 := recover(); e4876 != nil {
										if cljs_core.Value_(e4876).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4876
												_ = e
												return__5412 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4876)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5411, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5413 interface{}) {
								defer func() {
									if e4877 := recover(); e4877 != nil {
										if cljs_core.Value_(e4877).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4877
												_ = e
												return__5413 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4877)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_5411, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))")}))
							}
							seq__4870_5407, chunk__4871_5408, count__4872_5409, i__4873_5410 = seq__4870_5407, chunk__4871_5408, count__4872_5409, (i__4873_5410 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5414 = cljs_core.Seq.Arity1IQ(seq__4870_5407)
							_ = temp__4222__auto___5414
							if cljs_core.Truth_(temp__4222__auto___5414) {
								{
									var seq__4870_5415___1 = temp__4222__auto___5414
									_ = seq__4870_5415___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4870_5415___1) {
										{
											var c__960__auto___5416 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4870_5415___1)
											_ = c__960__auto___5416
											seq__4870_5407, chunk__4871_5408, count__4872_5409, i__4873_5410 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4870_5415___1), c__960__auto___5416, cljs_core.Count.X_invoke_Arity1(c__960__auto___5416).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_5417 = cljs_core.First.X_invoke_Arity1(seq__4870_5415___1)
											_ = n_5417
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5418 interface{}) {
												defer func() {
													if e4878 := recover(); e4878 != nil {
														if cljs_core.Value_(e4878).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4878
																_ = e
																return__5418 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4878)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5417, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5419 interface{}) {
												defer func() {
													if e4879 := recover(); e4879 != nil {
														if cljs_core.Value_(e4879).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4879
																_ = e
																return__5419 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4879)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_5417, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))")}))
											}
											seq__4870_5407, chunk__4871_5408, count__4872_5409, i__4873_5410 = cljs_core.Next.Arity1IQ(seq__4870_5415___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
							}
						}
					}
					break
				}
			}
			{
				var seq__4880_5420 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__4881_5421 interface{} = nil
				var count__4882_5422 = float64(0)
				var i__4883_5423 = float64(0)
				_, _, _, _ = seq__4880_5420, chunk__4881_5421, count__4882_5422, i__4883_5423
				for {
					if i__4883_5423 < count__4882_5422 {
						{
							var n_5424 = chunk__4881_5421.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4883_5423)
							_ = n_5424
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5425 interface{}) {
								defer func() {
									if e4886 := recover(); e4886 != nil {
										if cljs_core.Value_(e4886).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4886
												_ = e
												return__5425 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4886)

										}
									}
								}()
								{
									return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5424, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
							}
							seq__4880_5420, chunk__4881_5421, count__4882_5422, i__4883_5423 = seq__4880_5420, chunk__4881_5421, count__4882_5422, (i__4883_5423 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5426 = cljs_core.Seq.Arity1IQ(seq__4880_5420)
							_ = temp__4222__auto___5426
							if cljs_core.Truth_(temp__4222__auto___5426) {
								{
									var seq__4880_5427___1 = temp__4222__auto___5426
									_ = seq__4880_5427___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4880_5427___1) {
										{
											var c__960__auto___5428 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4880_5427___1)
											_ = c__960__auto___5428
											seq__4880_5420, chunk__4881_5421, count__4882_5422, i__4883_5423 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4880_5427___1), c__960__auto___5428, cljs_core.Count.X_invoke_Arity1(c__960__auto___5428).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_5429 = cljs_core.First.X_invoke_Arity1(seq__4880_5427___1)
											_ = n_5429
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5430 interface{}) {
												defer func() {
													if e4887 := recover(); e4887 != nil {
														if cljs_core.Value_(e4887).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4887
																_ = e
																return__5430 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4887)

														}
													}
												}()
												{
													return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5429, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
											}
											seq__4880_5420, chunk__4881_5421, count__4882_5422, i__4883_5423 = cljs_core.Next.Arity1IQ(seq__4880_5427___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
							}
						}
					}
					break
				}
			}
			{
				var map__4888_5431 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__4888_5432___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4888_5431) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4888_5431)
					} else {
						return map__4888_5431
					}
				}()
				var b_5433 = cljs_core.Get.X_invoke_Arity2(map__4888_5432___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a_5434 = cljs_core.Get.X_invoke_Arity2(map__4888_5432___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__4888_5431, map__4888_5432___1, b_5433, a_5434
				if cljs_core.X_EQ_.Arity2IIB(float64(1), a_5434) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), b_5433) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 b)")}))
				}
			}
			{
				var map__4889_5435 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__4889_5436___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4889_5435) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4889_5435)
					} else {
						return map__4889_5435
					}
				}()
				var d_5437 = cljs_core.Get.X_invoke_Arity2(map__4889_5436___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_5438 = cljs_core.Get.X_invoke_Arity2(map__4889_5436___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__4889_5435, map__4889_5436___1, d_5437, b_5438
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_5438) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_5437) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__4890_5439 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__4890_5440___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4890_5439) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4890_5439)
					} else {
						return map__4890_5439
					}
				}()
				var d_5441 = cljs_core.Get.X_invoke_Arity2(map__4890_5440___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_5442 = cljs_core.Get.X_invoke_Arity2(map__4890_5440___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__4890_5439, map__4890_5440___1, d_5441, b_5442
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_5442) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_5441) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__4891_5443 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}), float64(1), (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}), float64(2)}, nil})
				var map__4891_5444___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4891_5443) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4891_5443)
					} else {
						return map__4891_5443
					}
				}()
				var d_5445 = cljs_core.Get.X_invoke_Arity2(map__4891_5444___1, (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}))
				var b_5446 = cljs_core.Get.X_invoke_Arity2(map__4891_5444___1, (&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}))
				_, _, _, _ = map__4891_5443, map__4891_5444___1, d_5445, b_5446
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_5446) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_5445) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__4892_5447 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}), float64(2)}, nil})
				var map__4892_5448___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4892_5447) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4892_5447)
					} else {
						return map__4892_5447
					}
				}()
				var y_5449 = cljs_core.Get.X_invoke_Arity2(map__4892_5448___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}))
				var x_5450 = cljs_core.Get.X_invoke_Arity2(map__4892_5448___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}))
				_, _, _, _ = map__4892_5447, map__4892_5448___1, y_5449, x_5450
				if cljs_core.X_EQ_.Arity2IIB(x_5450, float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= x 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(y_5449, float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= y 2)")}))
				}
			}
			Cljs_739 = func(cljs_739 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(cljs_739, 2, func(arr interface{}, names interface{}) interface{} {
					for {
						{
							var name = cljs_core.First.X_invoke_Arity1(names)
							_ = name
							if cljs_core.Truth_(name) {
								arr, names = cljs_core.Conj.X_invoke_Arity2(arr, func(G__5451 *cljs_core.AFn, arr interface{}, names interface{}, name interface{}) *cljs_core.AFn {
									return cljs_core.Fn(G__5451, 0, func() interface{} {
										return cljs_core.Println.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{name}))
									})
								}(&cljs_core.AFn{}, arr, names, name)), cljs_core.Rest.Arity1IQ(names)
								continue
							} else {
								return arr
							}
						}
					}
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(func() string {
				var sb__1130__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1130__auto__
				{
					var _STAR_print_fn_STAR_4893_5452 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_4893_5452
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_4893_5452

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__5453 *cljs_core.AFn, _STAR_print_fn_STAR_4893_5452 interface{}, sb__1130__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__5453, 1, func(x__1131__auto__ interface{}) interface{} {
									return sb__1130__auto__.Append(x__1131__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_4893_5452, sb__1130__auto__)

							{
								var seq__4894_5454 interface{} = cljs_core.Seq.Arity1IQ(Cljs_739.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil})))
								var chunk__4895_5455 interface{} = nil
								var count__4896_5456 = float64(0)
								var i__4897_5457 = float64(0)
								_, _, _, _ = seq__4894_5454, chunk__4895_5455, count__4896_5456, i__4897_5457
								for {
									if i__4897_5457 < count__4896_5456 {
										{
											var fn_5458 = chunk__4895_5455.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4897_5457)
											_ = fn_5458
											{
												fn_5458.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
											}
											seq__4894_5454, chunk__4895_5455, count__4896_5456, i__4897_5457 = seq__4894_5454, chunk__4895_5455, count__4896_5456, (i__4897_5457 + float64(1))
											continue
										}
									} else {
										{
											var temp__4222__auto___5459 = cljs_core.Seq.Arity1IQ(seq__4894_5454)
											_ = temp__4222__auto___5459
											if cljs_core.Truth_(temp__4222__auto___5459) {
												{
													var seq__4894_5460___1 = temp__4222__auto___5459
													_ = seq__4894_5460___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4894_5460___1) {
														{
															var c__960__auto___5461 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4894_5460___1)
															_ = c__960__auto___5461
															seq__4894_5454, chunk__4895_5455, count__4896_5456, i__4897_5457 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4894_5460___1), c__960__auto___5461, cljs_core.Count.X_invoke_Arity1(c__960__auto___5461).(float64), float64(0)
															continue
														}
													} else {
														{
															var fn_5462 = cljs_core.First.X_invoke_Arity1(seq__4894_5460___1)
															_ = fn_5462
															{
																fn_5462.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
															}
															seq__4894_5454, chunk__4895_5455, count__4896_5456, i__4897_5457 = cljs_core.Next.Arity1IQ(seq__4894_5460___1), nil, float64(0), float64(0)
															continue
														}
													}
												}
											} else {
											}
										}
									}
									break
								}
							}
						}
					}()
				}
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1130__auto__).(string))
			}(), ":a\n:b\n:c\n:d\n") {
			} else {
				panic((&js.Error{("Assert failed: (= (with-out-str (doseq [fn (cljs-739 [] [:a :b :c :d])] (fn))) \":a\\n:b\\n:c\\n:d\\n\")")}))
			}
			{
				var seq__4898_5463 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__4899_5464 interface{} = nil
				var count__4900_5465 = float64(0)
				var i__4901_5466 = float64(0)
				_, _, _, _ = seq__4898_5463, chunk__4899_5464, count__4900_5465, i__4901_5466
				for {
					if i__4901_5466 < count__4900_5465 {
						{
							var n_5467 = chunk__4899_5464.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4901_5466)
							_ = n_5467
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5467)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get [1 2] n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5468 interface{}) {
								defer func() {
									if e4904 := recover(); e4904 != nil {
										if cljs_core.Value_(e4904).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4904
												_ = e
												return__5468 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4904)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5467)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5467, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get [1 2] n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5469 interface{}) {
								defer func() {
									if e4905 := recover(); e4905 != nil {
										if cljs_core.Value_(e4905).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4905
												_ = e
												return__5469 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4905)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5467, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5467)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (subvec [1 2] 1) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5470 interface{}) {
								defer func() {
									if e4906 := recover(); e4906 != nil {
										if cljs_core.Value_(e4906).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4906
												_ = e
												return__5470 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4906)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5467)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5467, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (subvec [1 2] 1) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5471 interface{}) {
								defer func() {
									if e4907 := recover(); e4907 != nil {
										if cljs_core.Value_(e4907).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4907
												_ = e
												return__5471 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4907)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5467, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5467)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (transient [1 2]) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5472 interface{}) {
								defer func() {
									if e4908 := recover(); e4908 != nil {
										if cljs_core.Value_(e4908).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4908
												_ = e
												return__5472 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4908)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5467)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5467, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (transient [1 2]) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5473 interface{}) {
								defer func() {
									if e4909 := recover(); e4909 != nil {
										if cljs_core.Value_(e4909).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4909
												_ = e
												return__5473 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4909)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5467, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5467)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (range 1 3) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5474 interface{}) {
								defer func() {
									if e4910 := recover(); e4910 != nil {
										if cljs_core.Value_(e4910).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4910
												_ = e
												return__5474 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4910)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5467)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5467, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (range 1 3) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5475 interface{}) {
								defer func() {
									if e4911 := recover(); e4911 != nil {
										if cljs_core.Value_(e4911).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4911
												_ = e
												return__5475 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4911)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5467, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))")}))
							}
							seq__4898_5463, chunk__4899_5464, count__4900_5465, i__4901_5466 = seq__4898_5463, chunk__4899_5464, count__4900_5465, (i__4901_5466 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5476 = cljs_core.Seq.Arity1IQ(seq__4898_5463)
							_ = temp__4222__auto___5476
							if cljs_core.Truth_(temp__4222__auto___5476) {
								{
									var seq__4898_5477___1 = temp__4222__auto___5476
									_ = seq__4898_5477___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4898_5477___1) {
										{
											var c__960__auto___5478 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4898_5477___1)
											_ = c__960__auto___5478
											seq__4898_5463, chunk__4899_5464, count__4900_5465, i__4901_5466 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4898_5477___1), c__960__auto___5478, cljs_core.Count.X_invoke_Arity1(c__960__auto___5478).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_5479 = cljs_core.First.X_invoke_Arity1(seq__4898_5477___1)
											_ = n_5479
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5479)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get [1 2] n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5480 interface{}) {
												defer func() {
													if e4912 := recover(); e4912 != nil {
														if cljs_core.Value_(e4912).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4912
																_ = e
																return__5480 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4912)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5479)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5479, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get [1 2] n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5481 interface{}) {
												defer func() {
													if e4913 := recover(); e4913 != nil {
														if cljs_core.Value_(e4913).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4913
																_ = e
																return__5481 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4913)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5479, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5479)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (subvec [1 2] 1) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5482 interface{}) {
												defer func() {
													if e4914 := recover(); e4914 != nil {
														if cljs_core.Value_(e4914).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4914
																_ = e
																return__5482 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4914)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5479)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5479, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (subvec [1 2] 1) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5483 interface{}) {
												defer func() {
													if e4915 := recover(); e4915 != nil {
														if cljs_core.Value_(e4915).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4915
																_ = e
																return__5483 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4915)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5479, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5479)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (transient [1 2]) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5484 interface{}) {
												defer func() {
													if e4916 := recover(); e4916 != nil {
														if cljs_core.Value_(e4916).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4916
																_ = e
																return__5484 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4916)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5479)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5479, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (transient [1 2]) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5485 interface{}) {
												defer func() {
													if e4917 := recover(); e4917 != nil {
														if cljs_core.Value_(e4917).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4917
																_ = e
																return__5485 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4917)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5479, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5479)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (range 1 3) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5486 interface{}) {
												defer func() {
													if e4918 := recover(); e4918 != nil {
														if cljs_core.Value_(e4918).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4918
																_ = e
																return__5486 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4918)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5479)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5479, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (range 1 3) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5487 interface{}) {
												defer func() {
													if e4919 := recover(); e4919 != nil {
														if cljs_core.Value_(e4919).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4919
																_ = e
																return__5487 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4919)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5479, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))")}))
											}
											seq__4898_5463, chunk__4899_5464, count__4900_5465, i__4901_5466 = cljs_core.Next.Arity1IQ(seq__4898_5477___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
							}
						}
					}
					break
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil})).X_rest_Arity1(), cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (= (-rest (rseq [0])) ())")}))
			}
			if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil})).(cljs_core.CljsCoreINext).X_next_Arity1()) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (-next (rseq [0])))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Rseq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil}))), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(0), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (set (rseq [0])) #{0})")}))
			}
			Cljs_780 = cljs_core.Atom.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList)}, nil}))}, nil})).(*cljs_core.CljsCoreAtom)

			cljs_core.Swap_BANG_.X_invoke_ArityVariadic(Cljs_780, cljs_core.Update_in, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}), cljs_core.Vary_meta, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Update_in, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}), cljs_core.Vec}))
			{
				var x_5488 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}).X_invoke_Arity1(cljs_core.Meta.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}).X_invoke_Arity1(cljs_core.Deref.X_invoke_Arity1(Cljs_780))))
				_ = x_5488
				if cljs_core.Vector_QMARK_.Arity1IB(x_5488) {
				} else {
					panic((&js.Error{("Assert failed: (vector? x)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(x_5488, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= x [1 2 3])")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3((&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}), "ToString", []interface{}{}), "550e8400-e29b-41d4-a716-446655440000") {
			} else {
				panic((&js.Error{("Assert failed: (= (.toString #uuid \"550e8400-e29b-41d4-a716-446655440000\") \"550e8400-e29b-41d4-a716-446655440000\")")}))
			}
			{
				var seq__4920_5489 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__4921_5490 interface{} = nil
				var count__4922_5491 = float64(0)
				var i__4923_5492 = float64(0)
				_, _, _, _ = seq__4920_5489, chunk__4921_5490, count__4922_5491, i__4923_5492
				for {
					if i__4923_5492 < count__4922_5491 {
						{
							var m_5493 = chunk__4921_5490.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4923_5492)
							_ = m_5493
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__5494 interface{}) {
								defer func() {
									if e4924 := recover(); e4924 != nil {
										if cljs_core.Value_(e4924).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var ___ = e4924
												_ = ___
												return__5494 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
											}
										} else {
											panic(e4924)

										}
									}
								}()
								{
									return cljs_core.Conj.X_invoke_Arity2(m_5493, "foo")
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5493, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m [:foo 1]))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5493, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m {:foo 1}))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5493, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m (list [:foo 1])))")}))
							}
							seq__4920_5489, chunk__4921_5490, count__4922_5491, i__4923_5492 = seq__4920_5489, chunk__4921_5490, count__4922_5491, (i__4923_5492 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5495 = cljs_core.Seq.Arity1IQ(seq__4920_5489)
							_ = temp__4222__auto___5495
							if cljs_core.Truth_(temp__4222__auto___5495) {
								{
									var seq__4920_5496___1 = temp__4222__auto___5495
									_ = seq__4920_5496___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4920_5496___1) {
										{
											var c__960__auto___5497 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4920_5496___1)
											_ = c__960__auto___5497
											seq__4920_5489, chunk__4921_5490, count__4922_5491, i__4923_5492 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4920_5496___1), c__960__auto___5497, cljs_core.Count.X_invoke_Arity1(c__960__auto___5497).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_5498 = cljs_core.First.X_invoke_Arity1(seq__4920_5496___1)
											_ = m_5498
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__5499 interface{}) {
												defer func() {
													if e4925 := recover(); e4925 != nil {
														if cljs_core.Value_(e4925).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var ___ = e4925
																_ = ___
																return__5499 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
															}
														} else {
															panic(e4925)

														}
													}
												}()
												{
													return cljs_core.Conj.X_invoke_Arity2(m_5498, "foo")
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5498, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m [:foo 1]))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5498, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m {:foo 1}))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5498, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m (list [:foo 1])))")}))
											}
											seq__4920_5489, chunk__4921_5490, count__4922_5491, i__4923_5492 = cljs_core.Next.Arity1IQ(seq__4920_5496___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
							}
						}
					}
					break
				}
			}
			{
				var seq__4926_5500 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Array_map, cljs_core.Hash_map, cljs_core.Sorted_map}, nil}))
				var chunk__4927_5501 interface{} = nil
				var count__4928_5502 = float64(0)
				var i__4929_5503 = float64(0)
				_, _, _, _ = seq__4926_5500, chunk__4927_5501, count__4928_5502, i__4929_5503
				for {
					if i__4929_5503 < count__4928_5502 {
						{
							var mt_5504 = chunk__4927_5501.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4929_5503)
							_ = mt_5504
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
								var G__4930 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
								var G__4931 = float64(1)
								_, _ = G__4930, G__4931
								return mt_5504.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__4930, G__4931)
							}(), func(make_seq *cljs_core.AFn, seq__4926_5500 interface{}, chunk__4927_5501 interface{}, count__4928_5502 float64, i__4929_5503 float64, mt_5504 interface{}) *cljs_core.AFn {
								return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
									if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
										X__GT_t4937 = func(__GT_t4937 *cljs_core.AFn, seq__4926_5500 interface{}, chunk__4927_5501 interface{}, count__4928_5502 float64, i__4929_5503 float64, mt_5504 interface{}) *cljs_core.AFn {
											return cljs_core.Fn(__GT_t4937, 9, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, i__4929___1 interface{}, count__4928___1 interface{}, chunk__4927___1 interface{}, seq__4926___1 interface{}, test_stuff___1 interface{}, meta4938 interface{}) interface{} {
												return (&CljsCore_testT4937{from_seq___1, make_seq___1, mt___1, i__4929___1, count__4928___1, chunk__4927___1, seq__4926___1, test_stuff___1, meta4938})
											})
										}(&cljs_core.AFn{}, seq__4926_5500, chunk__4927_5501, count__4928_5502, i__4929_5503, mt_5504)

										return (&CljsCore_testT4937{from_seq, make_seq, mt_5504, i__4929_5503, count__4928_5502, chunk__4927_5501, seq__4926_5500, test_stuff, nil})
									} else {
										return nil
									}
								})
							}(&cljs_core.AFn{}, seq__4926_5500, chunk__4927_5501, count__4928_5502, i__4929_5503, mt_5504).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))")}))
							}
							seq__4926_5500, chunk__4927_5501, count__4928_5502, i__4929_5503 = seq__4926_5500, chunk__4927_5501, count__4928_5502, (i__4929_5503 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5505 = cljs_core.Seq.Arity1IQ(seq__4926_5500)
							_ = temp__4222__auto___5505
							if cljs_core.Truth_(temp__4222__auto___5505) {
								{
									var seq__4926_5506___1 = temp__4222__auto___5505
									_ = seq__4926_5506___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4926_5506___1) {
										{
											var c__960__auto___5507 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4926_5506___1)
											_ = c__960__auto___5507
											seq__4926_5500, chunk__4927_5501, count__4928_5502, i__4929_5503 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4926_5506___1), c__960__auto___5507, cljs_core.Count.X_invoke_Arity1(c__960__auto___5507).(float64), float64(0)
											continue
										}
									} else {
										{
											var mt_5508 = cljs_core.First.X_invoke_Arity1(seq__4926_5506___1)
											_ = mt_5508
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
												var G__4942 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
												var G__4943 = float64(1)
												_, _ = G__4942, G__4943
												return mt_5508.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__4942, G__4943)
											}(), func(make_seq *cljs_core.AFn, seq__4926_5500 interface{}, chunk__4927_5501 interface{}, count__4928_5502 float64, i__4929_5503 float64, mt_5508 interface{}, seq__4926_5506___1 interface{}, temp__4222__auto___5505 cljs_core.CljsCoreISeq) *cljs_core.AFn {
												return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
													if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
														X__GT_t4949 = func(__GT_t4949 *cljs_core.AFn, seq__4926_5500 interface{}, chunk__4927_5501 interface{}, count__4928_5502 float64, i__4929_5503 float64, mt_5508 interface{}, seq__4926_5506___1 interface{}, temp__4222__auto___5505 cljs_core.CljsCoreISeq) *cljs_core.AFn {
															return cljs_core.Fn(__GT_t4949, 10, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, temp__4222__auto_____1 interface{}, i__4929___1 interface{}, count__4928___1 interface{}, chunk__4927___1 interface{}, seq__4926___2 interface{}, test_stuff___1 interface{}, meta4950 interface{}) interface{} {
																return (&CljsCore_testT4949{from_seq___1, make_seq___1, mt___1, temp__4222__auto_____1, i__4929___1, count__4928___1, chunk__4927___1, seq__4926___2, test_stuff___1, meta4950})
															})
														}(&cljs_core.AFn{}, seq__4926_5500, chunk__4927_5501, count__4928_5502, i__4929_5503, mt_5508, seq__4926_5506___1, temp__4222__auto___5505)

														return (&CljsCore_testT4949{from_seq, make_seq, mt_5508, temp__4222__auto___5505, i__4929_5503, count__4928_5502, chunk__4927_5501, seq__4926_5506___1, test_stuff, nil})
													} else {
														return nil
													}
												})
											}(&cljs_core.AFn{}, seq__4926_5500, chunk__4927_5501, count__4928_5502, i__4929_5503, mt_5508, seq__4926_5506___1, temp__4222__auto___5505).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))")}))
											}
											seq__4926_5500, chunk__4927_5501, count__4928_5502, i__4929_5503 = cljs_core.Next.Arity1IQ(seq__4926_5506___1), nil, float64(0), float64(0)
											continue
										}
									}
								}
							} else {
							}
						}
					}
					break
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_4954 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4954
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4954

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(1)

						return (`` + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(0)}, nil})).(string))
					}
				}()
			}(), "[1 ...]") {
			} else {
				panic((&js.Error{("Assert failed: (= (binding [*print-length* 1] (str [1 2 3 4 5 6 7 8 9 0])) \"[1 ...]\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_4955 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4955
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4955

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(2)

						return (`` + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(0)}, nil})).(string))
					}
				}()
			}(), "[1 2 ...]") {
			} else {
				panic((&js.Error{("Assert failed: (= (binding [*print-length* 2] (str [1 2 3 4 5 6 7 8 9 0])) \"[1 2 ...]\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_4956 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4956
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4956

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(10)

						return (`` + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(0)}, nil})).(string))
					}
				}()
			}(), "[1 2 3 4 5 6 7 8 9 0]") {
			} else {
				panic((&js.Error{("Assert failed: (= (binding [*print-length* 10] (str [1 2 3 4 5 6 7 8 9 0])) \"[1 2 3 4 5 6 7 8 9 0]\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_4957 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4957
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4957

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(10)

						return (`` + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar"}, nil})).(string))
					}
				}()
			}(), "{:foo \"bar\"}") {
			} else {
				panic((&js.Error{("Assert failed: (= (binding [*print-length* 10] (str {:foo \"bar\"})) \"{:foo \\\"bar\\\"}\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_4958 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4958
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4958

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(1)

						return (`` + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), "woz"}, nil})).(string))
					}
				}()
			}(), "{:foo \"bar\", ...}") {
			} else {
				panic((&js.Error{("Assert failed: (= (binding [*print-length* 1] (str {:foo \"bar\", :baz \"woz\"})) \"{:foo \\\"bar\\\", ...}\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_4959 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4959
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4959

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(10)

						return (`` + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), "woz"}, nil})).(string))
					}
				}()
			}(), "{:foo \"bar\", :baz \"woz\"}") {
			} else {
				panic((&js.Error{("Assert failed: (= (binding [*print-length* 10] (str {:foo \"bar\", :baz \"woz\"})) \"{:foo \\\"bar\\\", :baz \\\"woz\\\"}\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var x = "a"
				_ = x
				{
					var G__4960 = func() interface{} {
						if cljs_core.Value_(x).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
							return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(x), "Fqn")
						} else {
							return nil
						}
					}()
					_ = G__4960
					switch G__4960 {
					case "a":
						return float64(1)

					default:
						return "a"

					}
				}
			}(), "a") {
			} else {
				panic((&js.Error{("Assert failed: (= (let [x \"a\"] (case x :a 1 \"a\")) \"a\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("0atrue:key/wordsymb/olfalse[1 2 3 4]1234.56789", (`` + cljs_core.Str.X_invoke_Arity1(float64(0)).(string) + "a" + cljs_core.Str.X_invoke_Arity1(true).(string) + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: "key", Name: "word", Fqn: "key/word", X_hash: float64(-420295340)})).(string) + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: "symb", Name: "ol", Str: "symb/ol", X_hash: float64(-1700033577), X_meta: nil})).(string) + cljs_core.Str.X_invoke_Arity1(false).(string) + cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil})).(string) + cljs_core.Str.X_invoke_Arity1(1234.5678).(string) + cljs_core.Str.X_invoke_Arity1(float64(9)).(string))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"0atrue:key/wordsymb/olfalse[1 2 3 4]1234.56789\" (str 0 \"a\" true nil :key/word (quote symb/ol) false [1 2 3 4] 1234.5678 9))")}))
			}
			if cljs_core.Int_rotate_left.Arity2IIF(float64((cljs_core.Int32_(float64(2271560481))|cljs_core.Int32_(float64(0)))), float64(4)) == float64((cljs_core.Int32_(float64(1985229336)) | cljs_core.Int32_(float64(0)))) {
			} else {
				panic((&js.Error{("Assert failed: (== (int-rotate-left (bit-or 2271560481 0) 4) (bit-or 1985229336 0))")}))
			}
			if cljs_core.Int_rotate_left.Arity2IIF(float64((cljs_core.Int32_(float64(2271560481))|cljs_core.Int32_(float64(0)))), float64(8)) == float64((cljs_core.Int32_(float64(1698898311)) | cljs_core.Int32_(float64(0)))) {
			} else {
				panic((&js.Error{("Assert failed: (== (int-rotate-left (bit-or 2271560481 0) 8) (bit-or 1698898311 0))")}))
			}
			if cljs_core.Int_rotate_left.Arity2IIF(float64((cljs_core.Int32_(float64(2147483648))|cljs_core.Int32_(float64(0)))), float64(1)) == float64(1) {
			} else {
				panic((&js.Error{("Assert failed: (== (int-rotate-left (bit-or 2147483648 0) 1) 1)")}))
			}
			if cljs_core.Int_rotate_left.Arity2IIF(float64((cljs_core.Int32_(float64(2014458966))|cljs_core.Int32_(float64(0)))), float64(4)) == float64((cljs_core.Int32_(float64(2166572391)) | cljs_core.Int32_(float64(0)))) {
			} else {
				panic((&js.Error{("Assert failed: (== (int-rotate-left (bit-or 2014458966 0) 4) (bit-or 2166572391 0))")}))
			}
			if cljs_core.Int_rotate_left.Arity2IIF(float64((cljs_core.Int32_(float64(4294967295))|cljs_core.Int32_(float64(0)))), float64(4)) == float64((cljs_core.Int32_(float64(4294967295)) | cljs_core.Int32_(float64(0)))) {
			} else {
				panic((&js.Error{("Assert failed: (== (int-rotate-left (bit-or 4294967295 0) 4) (bit-or 4294967295 0))")}))
			}
			if cljs_core.Imul.Arity2IIF(float64(3), float64(3)) == float64(9) {
			} else {
				panic((&js.Error{("Assert failed: (== (imul 3 3) 9)")}))
			}
			if cljs_core.Imul.Arity2IIF(float64(-1), float64(8)) == float64(-8) {
			} else {
				panic((&js.Error{("Assert failed: (== (imul -1 8) -8)")}))
			}
			if cljs_core.Imul.Arity2IIF(float64(-2), float64(-2)) == float64(4) {
			} else {
				panic((&js.Error{("Assert failed: (== (imul -2 -2) 4)")}))
			}
			if cljs_core.Imul.Arity2IIF(float64(4294967295), float64(5)) == float64(-5) {
			} else {
				panic((&js.Error{("Assert failed: (== (imul 4294967295 5) -5)")}))
			}
			if cljs_core.Imul.Arity2IIF(float64(4294967294), float64(5)) == float64(-10) {
			} else {
				panic((&js.Error{("Assert failed: (== (imul 4294967294 5) -10)")}))
			}
			Case_recur = func(case_recur *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(case_recur, 1, func(value interface{}) interface{} {
					for {
						{
							var G__4962 = func() interface{} {
								if cljs_core.Value_(value).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
									return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(value), "Fqn")
								} else {
									return nil
								}
							}()
							_ = G__4962
							switch G__4962 {
							case "b":
								return float64(0)

							case "a":
								value = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
								continue

							default:
								panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(value).(string))}))

							}
						}
					}
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(Case_recur.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})), float64(0)) {
			} else {
				panic((&js.Error{("Assert failed: (= (case-recur :a) 0)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(clojure_set.Rename_keys.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "two"}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "z", Fqn: "z", X_hash: float64(-789527183)})}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "z", Fqn: "z", X_hash: float64(-789527183)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "two"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (set/rename-keys {:a \"one\", :b \"two\"} {:a :z}) {:z \"one\", :b \"two\"})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(clojure_set.Rename_keys.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "two"}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "z", Fqn: "z", X_hash: float64(-789527183)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "y", Fqn: "y", X_hash: float64(-1757859776)})}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "z", Fqn: "z", X_hash: float64(-789527183)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "two"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (set/rename-keys {:a \"one\", :b \"two\"} {:a :z, :c :y}) {:z \"one\", :b \"two\"})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(clojure_set.Rename_keys.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "two", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), "three"}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), "two", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), "three"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (set/rename-keys {:a \"one\", :b \"two\", :c \"three\"} {:a :b, :b :a}) {:a \"two\", :b \"one\", :c \"three\"})")}))
			}
			{
				var not_strings_5511 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, nil, float64(1), func(G__5512 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5512, 0, func() interface{} {
						return nil
					})
				}(&cljs_core.AFn{})}, nil})
				_ = not_strings_5511
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__5513 *cljs_core.AFn, not_strings_5511 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__5513, 1, func(p1__4101_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__5514 interface{}) {
							defer func() {
								if e4963 := recover(); e4963 != nil {
									if cljs_core.Value_(e4963).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e4963
											_ = ___
											return__5514 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e4963)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__4101_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_5511), not_strings_5511) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__4101#] (= :failed (try (re-find #\".\" p1__4101#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__5515 *cljs_core.AFn, not_strings_5511 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__5515, 1, func(p1__4102_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__5516 interface{}) {
							defer func() {
								if e4964 := recover(); e4964 != nil {
									if cljs_core.Value_(e4964).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e4964
											_ = ___
											return__5516 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e4964)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__4102_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_5511), not_strings_5511) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__4102#] (= :failed (try (re-matches #\".\" p1__4102#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__5517 *cljs_core.AFn, not_strings_5511 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__5517, 1, func(p1__4103_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__5518 interface{}) {
							defer func() {
								if e4965 := recover(); e4965 != nil {
									if cljs_core.Value_(e4965).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e4965
											_ = ___
											return__5518 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e4965)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__4103_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_5511), not_strings_5511) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__4103#] (= :failed (try (re-find #\"nomatch\" p1__4103#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__5519 *cljs_core.AFn, not_strings_5511 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__5519, 1, func(p1__4104_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__5520 interface{}) {
							defer func() {
								if e4966 := recover(); e4966 != nil {
									if cljs_core.Value_(e4966).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e4966
											_ = ___
											return__5520 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e4966)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__4104_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_5511), not_strings_5511) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__4104#] (= :failed (try (re-matches #\"nomatch\" p1__4104#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
			}
			if cljs_core.Truth_((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}).Equiv((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
			} else {
				panic((&js.Error{("Assert failed: (.equiv :foo :foo)")}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil}).Equiv((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (.equiv (quote foo) (quote foo))")}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}).Equiv((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (.equiv {:foo 1, :bar 2} {:foo 1, :bar 2})")}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}).Equiv((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (.equiv [1 2 3] [1 2 3])")}))
			}
			if cljs_core.Truth_(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList).Equiv(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList))) {
			} else {
				panic((&js.Error{("Assert failed: (.equiv (quote (1 2 3)) (quote (1 2 3)))")}))
			}
			if cljs_core.Truth_(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq).Equiv(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{("Assert failed: (.equiv (map inc [1 2 3]) (map inc [1 2 3]))")}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "cat", Fqn: "cat", X_hash: float64(-1457810207)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bird", Fqn: "bird", X_hash: float64(-1252014845)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "dog", Fqn: "dog", X_hash: float64(-1650861974)}), nil}, nil}, nil}).Equiv((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "cat", Fqn: "cat", X_hash: float64(-1457810207)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bird", Fqn: "bird", X_hash: float64(-1252014845)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "dog", Fqn: "dog", X_hash: float64(-1650861974)}), nil}, nil}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (.equiv #{:cat :bird :dog} #{:cat :bird :dog})")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), []interface{}{float64(1), float64(2), float64(3)}), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (map inc) (array 1 2 3)) (quote (2 3 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(func(G__5521 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5521, 1, func(p1__4105_SHARP_ interface{}) interface{} {
					return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(p1__4105_SHARP_, "ToUpperCase", []interface{}{})
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), "foo")), "FOO") {
			} else {
				panic((&js.Error{("Assert failed: (= (apply str (sequence (map (fn* [p1__4105#] (.toUpperCase p1__4105#))) \"foo\")) \"FOO\")")}))
			}
			if cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(float64) == cljs_core.Hash.X_invoke_Arity1(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))).(float64) {
			} else {
				panic((&js.Error{("Assert failed: (== (hash [1 2 3]) (hash (sequence (map inc) (range 3))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 3] (sequence (map inc) (range 3)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange)), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (map inc) (range 3)) [1 2 3])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Remove.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (remove even?) (range 10)) (quote (1 3 5 7 9)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Take.X_invoke_Arity1(float64(5)).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (take 5) (range 10)) (quote (0 1 2 3 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Take_while.X_invoke_Arity1(func(G__5522 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5522, 1, func(p1__4106_SHARP_ interface{}) interface{} {
					return (p1__4106_SHARP_.(float64) < float64(5))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (take-while (fn* [p1__4106#] (< p1__4106# 5))) (range 10)) (quote (0 1 2 3 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Drop.X_invoke_Arity1(float64(5)).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (drop 5) (range 10)) (quote (5 6 7 8 9)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Drop_while.X_invoke_Arity1(func(G__5523 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5523, 1, func(p1__4107_SHARP_ interface{}) interface{} {
					return (p1__4107_SHARP_.(float64) < float64(5))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (drop-while (fn* [p1__4107#] (< p1__4107# 5))) (range 10)) (quote (5 6 7 8 9)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Take_nth.X_invoke_Arity1(float64(2)).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(2), float64(4), float64(6), float64(8)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (take-nth 2) (range 10)) (quote (0 2 4 6 8)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Replace.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})).(cljs_core.CljsCoreIFn), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2)})).(*cljs_core.CljsCoreList)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (replace {:foo :bar}) (quote (:foo 1 :foo 2))) (quote (:bar 1 :bar 2)))")}))
			}
			{
				var ret_5524 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))
				_ = ret_5524
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_5524)) && (cljs_core.X_EQ_.Arity2IIB(ret_5524, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList))) {
				} else {
					panic((&js.Error{("Assert failed: (and (vector? ret) (= ret (quote (1 2 3))))")}))
				}
			}
			{
				var ret_5525 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				_ = ret_5525
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_5525)) && (cljs_core.X_EQ_.Arity2IIB(ret_5525, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(2), float64(4), float64(6), float64(8)})).(*cljs_core.CljsCoreList))) {
				} else {
					panic((&js.Error{("Assert failed: (and (vector? ret) (= ret (quote (0 2 4 6 8))))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (map inc (sequence (map inc) (range 3))) (quote (2 3 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Dedupe.X_invoke_Arity0().(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(2), float64(2), float64(3), float64(3)}, nil})), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (dedupe) [1 1 2 2 3 3]) (quote (1 2 3)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mapcat.X_invoke_ArityVariadic(cljs_core.Reverse, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(6), float64(5), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(9), float64(8), float64(7)}, nil})}, nil})})), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)) {
			} else {
				panic((&js.Error{("Assert failed: (= (mapcat reverse [[3 2 1 0] [6 5 4] [9 8 7]]) (range 10))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Reverse).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(6), float64(5), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(9), float64(8), float64(7)}, nil})}, nil})), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (mapcat reverse) [[3 2 1 0] [6 5 4] [9 8 7]]) (range 10))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(cljs_core.Iteration.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreIteration)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (seq (iteration (map inc) [1 2 3])) (quote (2 3 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Partition_by.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "split", Fqn: "split", X_hash: float64(-599435118)}), nil}, nil}, nil})).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "split", Fqn: "split", X_hash: float64(-599435118)}), float64(4), float64(5), float64(6)}, nil})), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "split", Fqn: "split", X_hash: float64(-599435118)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6)}, nil})})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (partition-by #{:split}) [1 2 3 :split 4 5 6]) (quote ([1 2 3] [:split] [4 5 6])))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Partition_all.X_invoke_Arity1(float64(3)).(cljs_core.CljsCoreIFn), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)})).(*cljs_core.CljsCoreList)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil})})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (partition-all 3) (quote (1 2 3 4 5))) (quote ([1 2 3] [4 5])))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Keep.X_invoke_Arity1(cljs_core.Identity).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), nil, float64(2), nil, float64(3)}, nil})), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (keep identity) [1 nil 2 nil 3]) (quote (1 2 3)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Keep_indexed.X_invoke_Arity2(cljs_core.Identity, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil})), cljs_core.Sequence.X_invoke_Arity2(cljs_core.Keep_indexed.X_invoke_Arity1(cljs_core.Identity).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (keep-indexed identity [:foo nil :bar nil :baz]) (sequence (keep-indexed identity) [:foo nil :bar nil :baz]))")}))
			}
			Xform = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Dedupe.X_invoke_Arity0().(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Partition_all.X_invoke_Arity1(float64(3)).(cljs_core.CljsCoreIFn), cljs_core.Partition_by.X_invoke_Arity1(func(G__5526 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5526, 1, func(p1__4108_SHARP_ interface{}) interface{} {
					return (cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, p1__4108_SHARP_).(float64) < float64(7))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Flatten).(cljs_core.CljsCoreIFn), cljs_core.Random_sample.X_invoke_Arity1(1.0).(cljs_core.CljsCoreIFn), cljs_core.Take_nth.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Keep.X_invoke_Arity1(func(G__5527 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5527, 1, func(p1__4109_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__4109_SHARP_) {
						return (p1__4109_SHARP_.(float64) * p1__4109_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Keep_indexed.X_invoke_Arity1(func(G__5528 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5528, 2, func(p1__4110_SHARP_ interface{}, p2__4111_SHARP_ interface{}) interface{} {
					if cljs_core.Even_QMARK_.Arity1IB(p1__4110_SHARP_) {
						return (p1__4110_SHARP_.(float64) * p2__4111_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Replace.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(2), "two", float64(6), "six", float64(18), "eighteen"}, nil})).(cljs_core.CljsCoreIFn), cljs_core.Take.X_invoke_Arity1(float64(11)).(cljs_core.CljsCoreIFn), cljs_core.Take_while.X_invoke_Arity1(func(G__5529 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5529, 1, func(p1__4112_SHARP_ interface{}) interface{} {
					return cljs_core.Not_EQ_.Arity2IIB(float64(300), p1__4112_SHARP_)
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Drop.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Drop_while.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Remove.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn)})).(cljs_core.CljsCoreIFn).(*cljs_core.AFn)

			Data = cljs_core.Vec.X_invoke_Arity1(cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(18)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(20)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(Xform, Data), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(36), float64(200), float64(10)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence xform data) (quote (36 200 10)))")}))
			}
			Xf = cljs_core.Map_.X_invoke_Arity1(func(G__5530 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5530, 2, func(p1__4113_SHARP_ interface{}, p2__4114_SHARP_ interface{}) interface{} {
					return (p1__4113_SHARP_.(float64) + p2__4114_SHARP_.(float64))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn).(*cljs_core.AFn)

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_ArityVariadic(Xf, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0)}, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence xf [0 0] [1 2]) [1 2])")}))
			}
			{
				var xs_5531 = (&cljs_core.CljsCorePersistentVector{nil, float64(21), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(44), float64(43), float64(42), float64(41), float64(40), float64(39), float64(38), float64(37), float64(36), float64(35), float64(34), float64(33), float64(32), float64(31), float64(30), float64(29), float64(28), float64(27), float64(26), float64(25), float64(24)}, nil})
				_ = xs_5531
				{
					var m_5532 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.Zipmap.X_invoke_Arity2(xs_5531, cljs_core.Repeat.X_invoke_Arity1(float64(1)).(*cljs_core.CljsCoreLazySeq)))
					var xs_5533___1 interface{} = xs_5531
					_, _ = m_5532, xs_5533___1
					for {
						{
							var temp__4220__auto___5534 = cljs_core.First.X_invoke_Arity1(xs_5533___1)
							_ = temp__4220__auto___5534
							if cljs_core.Truth_(temp__4220__auto___5534) {
								{
									var x_5535 = temp__4220__auto___5534
									_ = x_5535
									if cljs_core.Contains_QMARK_.Arity2IIB(m_5532, x_5535) {
										m_5532, xs_5533___1 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(m_5532, x_5535), cljs_core.Next.Arity1IQ(xs_5533___1)
										continue
									} else {
										panic(cljs_core.Ex_info.X_invoke_Arity2("CLJS-849 regression!", (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "m", Fqn: "m", X_hash: float64(1632677161)}), cljs_core.Persistent_BANG_.X_invoke_Arity1(m_5532), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "xs", Fqn: "xs", X_hash: float64(649443341)}), xs_5533___1}, nil})).(*cljs_core.CljsCoreExceptionInfo))
									}
								}
							} else {
							}
						}
						break
					}
				}
			}
			return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
		})
	}(&cljs_core.AFn{})

}

var Test_stuff *cljs_core.AFn

var Bar *cljs_core.CljsCoreMultiFn

var Nested_dispatch *cljs_core.CljsCoreMultiFn

var Nested_dispatch2 *cljs_core.CljsCoreMultiFn

var Foo1 *cljs_core.CljsCoreMultiFn

var Area *cljs_core.CljsCoreMultiFn

var Rect *cljs_core.AFn

var Circle *cljs_core.AFn

var R cljs_core.CljsCoreIMap

var C cljs_core.CljsCoreIMap

var Apply_multi_test *cljs_core.CljsCoreMultiFn

var My_map_hierarchy *cljs_core.CljsCoreAtom

var My_map_QMARK_ *cljs_core.CljsCoreMultiFn

type CljsCore_testFixedHash struct {
	H interface{}
	V interface{}
}

func (_ *CljsCore_testFixedHash) CljsCoreIEquiv__() {}

func (this *CljsCore_testFixedHash) X_equiv_Arity2(other interface{}) bool {
	return (cljs_core.Value_(other).Type().AssignableTo(reflect.TypeOf((**CljsCore_testFixedHash)(nil)).Elem())) && (cljs_core.X_EQ_.Arity2IIB(this.V, cljs_core.Native_get_instance_field.X_invoke_Arity2(other, "V")))
}

func (_ *CljsCore_testFixedHash) CljsCoreIHash__() {}

func (this *CljsCore_testFixedHash) X_hash_Arity1() interface{} {
	return this.H
}

var X__GT_FixedHash *cljs_core.AFn

var Fixed_hash_foo *CljsCore_testFixedHash

var Fixed_hash_bar *CljsCore_testFixedHash

var Array_map_conversion_threshold float64

type CljsCore_testPerson struct {
	Firstname interface{}
	Lastname  interface{}
	X__meta   interface{}
	X__extmap interface{}
	X__hash   interface{}
}

func (_ *CljsCore_testPerson) CljsCoreILookup__() {}

func (this__764__auto__ *CljsCore_testPerson) X_lookup_Arity2(k__765__auto__ interface{}) interface{} {
	return this__764__auto__.X_lookup_Arity3(k__765__auto__, nil)
}

func (this__766__auto__ *CljsCore_testPerson) X_lookup_Arity3(k4685 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__4688 = func() interface{} {
			if cljs_core.Value_(k4685).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k4685), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4688
		switch G__4688 {
		case "lastname":
			return this__766__auto__.Lastname

		case "firstname":
			return this__766__auto__.Firstname

		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k4685, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIPrintWithWriter__() {}

func (this__780__auto__ *CljsCore_testPerson) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__5537 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5537, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, cljs_core.Pr_writer, "", " ", "", opts__782__auto__, keyval__784__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__783__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, pr_pair__783__auto__, "#cljs.core-test.Person{", ", ", "}", opts__782__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), this__780__auto__.Firstname}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), this__780__auto__.Lastname}, nil})}, nil}), this__780__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testPerson) CljsCoreIMeta__() {}

func (this__762__auto__ *CljsCore_testPerson) X_meta_Arity1() interface{} {
	return this__762__auto__.X__meta
}

func (_ *CljsCore_testPerson) CljsCoreICloneable__() {}

func (this__758__auto__ *CljsCore_testPerson) X_clone_Arity1() interface{} {
	return (&CljsCore_testPerson{this__758__auto__.Firstname, this__758__auto__.Lastname, this__758__auto__.X__meta, this__758__auto__.X__extmap, this__758__auto__.X__hash})
}

func (_ *CljsCore_testPerson) CljsCoreICounted__() {}

func (this__768__auto__ *CljsCore_testPerson) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__768__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testPerson) CljsCoreIHash__() {}

func (this__759__auto__ *CljsCore_testPerson) X_hash_Arity1() interface{} {
	{
		var h__583__auto__ = this__759__auto__.X__hash
		_ = h__583__auto__
		if !(cljs_core.Nil_(h__583__auto__)) {
			return h__583__auto__
		} else {
			{
				var h__583__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__583__auto_____1
				this__759__auto__.X__hash = h__583__auto_____1

				return h__583__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIEquiv__() {}

func (this__760__auto__ *CljsCore_testPerson) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__165__auto__ = other__761__auto__
		_ = and__165__auto__
		if cljs_core.Truth_(and__165__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
		} else {
			return and__165__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testPerson) CljsCoreIRecord__() {}

func (_ *CljsCore_testPerson) CljsCoreIMap__() {}

func (this__775__auto__ *CljsCore_testPerson) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), nil}, nil}, nil}), k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testPerson{this__775__auto__.Firstname, this__775__auto__.Lastname, this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testPerson) CljsCoreIAssociative__() {}

func (this__771__auto__ *CljsCore_testPerson) X_assoc_Arity3(k__772__auto__ interface{}, G__4684 interface{}) interface{} {
	{
		var pred__4696 = cljs_core.Keyword_identical_QMARK_
		var expr__4697 = k__772__auto__
		_, _ = pred__4696, expr__4697
		if cljs_core.Truth_(func() interface{} {
			var G__4699 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)})
			var G__4700 = expr__4697
			_, _ = G__4699, G__4700
			return pred__4696.X_invoke_Arity2(G__4699, G__4700)
		}()) {
			return (&CljsCore_testPerson{G__4684, this__771__auto__.Lastname, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__4701 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})
				var G__4702 = expr__4697
				_, _ = G__4701, G__4702
				return pred__4696.X_invoke_Arity2(G__4701, G__4702)
			}()) {
				return (&CljsCore_testPerson{this__771__auto__.Firstname, G__4684, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPerson{this__771__auto__.Firstname, this__771__auto__.Lastname, this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__4684), nil})
			}
		}
	}
}

func (this__773__auto__ *CljsCore_testPerson) X_contains_key_QMARK__Arity2(k__774__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testPerson) CljsCoreISeqable__() {}

func (this__778__auto__ *CljsCore_testPerson) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), this__778__auto__.Firstname}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), this__778__auto__.Lastname}, nil})}, nil}), this__778__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testPerson) CljsCoreIWithMeta__() {}

func (this__763__auto__ *CljsCore_testPerson) X_with_meta_Arity2(G__4684 interface{}) interface{} {
	return (&CljsCore_testPerson{this__763__auto__.Firstname, this__763__auto__.Lastname, G__4684, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testPerson) CljsCoreICollection__() {}

func (this__769__auto__ *CljsCore_testPerson) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__769__auto__, entry__770__auto__)
	}
}

var X__GT_Person *cljs_core.AFn

var Map__GT_Person *cljs_core.AFn

var Fred *CljsCore_testPerson

var Fred_too *CljsCore_testPerson

type CljsCore_testA struct {
	X__meta   interface{}
	X__extmap interface{}
	X__hash   interface{}
}

func (_ *CljsCore_testA) CljsCoreILookup__() {}

func (this__764__auto__ *CljsCore_testA) X_lookup_Arity2(k__765__auto__ interface{}) interface{} {
	return this__764__auto__.X_lookup_Arity3(k__765__auto__, nil)
}

func (this__766__auto__ *CljsCore_testA) X_lookup_Arity3(k4704 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__4707 = k4704
		_ = G__4707
		switch G__4707 {
		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k4704, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testA) CljsCoreIPrintWithWriter__() {}

func (this__780__auto__ *CljsCore_testA) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__5539 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5539, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, cljs_core.Pr_writer, "", " ", "", opts__782__auto__, keyval__784__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__783__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, pr_pair__783__auto__, "#cljs.core-test.A{", ", ", "}", opts__782__auto__, cljs_core.Concat.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, this__780__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testA) CljsCoreIMeta__() {}

func (this__762__auto__ *CljsCore_testA) X_meta_Arity1() interface{} {
	return this__762__auto__.X__meta
}

func (_ *CljsCore_testA) CljsCoreICloneable__() {}

func (this__758__auto__ *CljsCore_testA) X_clone_Arity1() interface{} {
	return (&CljsCore_testA{this__758__auto__.X__meta, this__758__auto__.X__extmap, this__758__auto__.X__hash})
}

func (_ *CljsCore_testA) CljsCoreICounted__() {}

func (this__768__auto__ *CljsCore_testA) X_count_Arity1() float64 {
	return (float64(0) + cljs_core.Count.X_invoke_Arity1(this__768__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testA) CljsCoreIHash__() {}

func (this__759__auto__ *CljsCore_testA) X_hash_Arity1() interface{} {
	{
		var h__583__auto__ = this__759__auto__.X__hash
		_ = h__583__auto__
		if !(cljs_core.Nil_(h__583__auto__)) {
			return h__583__auto__
		} else {
			{
				var h__583__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__583__auto_____1
				this__759__auto__.X__hash = h__583__auto_____1

				return h__583__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testA) CljsCoreIEquiv__() {}

func (this__760__auto__ *CljsCore_testA) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__165__auto__ = other__761__auto__
		_ = and__165__auto__
		if cljs_core.Truth_(and__165__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
		} else {
			return and__165__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testA) CljsCoreIRecord__() {}

func (_ *CljsCore_testA) CljsCoreIMap__() {}

func (this__775__auto__ *CljsCore_testA) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testA{this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testA) CljsCoreIAssociative__() {}

func (this__771__auto__ *CljsCore_testA) X_assoc_Arity3(k__772__auto__ interface{}, G__4703 interface{}) interface{} {
	{
		var pred__4711 = cljs_core.Keyword_identical_QMARK_
		var expr__4712 = k__772__auto__
		_, _ = pred__4711, expr__4712
		return (&CljsCore_testA{this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__4703), nil})
	}
}

func (this__773__auto__ *CljsCore_testA) X_contains_key_QMARK__Arity2(k__774__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testA) CljsCoreISeqable__() {}

func (this__778__auto__ *CljsCore_testA) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, this__778__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testA) CljsCoreIWithMeta__() {}

func (this__763__auto__ *CljsCore_testA) X_with_meta_Arity2(G__4703 interface{}) interface{} {
	return (&CljsCore_testA{G__4703, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testA) CljsCoreICollection__() {}

func (this__769__auto__ *CljsCore_testA) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__769__auto__, entry__770__auto__)
	}
}

var X__GT_A *cljs_core.AFn

var Map__GT_A *cljs_core.AFn

type CljsCore_testC struct {
	A         interface{}
	B         interface{}
	C         interface{}
	X__meta   interface{}
	X__extmap interface{}
	X__hash   interface{}
}

func (_ *CljsCore_testC) CljsCoreILookup__() {}

func (this__764__auto__ *CljsCore_testC) X_lookup_Arity2(k__765__auto__ interface{}) interface{} {
	return this__764__auto__.X_lookup_Arity3(k__765__auto__, nil)
}

func (this__766__auto__ *CljsCore_testC) X_lookup_Arity3(k4715 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__4718 = func() interface{} {
			if cljs_core.Value_(k4715).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k4715), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4718
		switch G__4718 {
		case "c":
			return this__766__auto__.C

		case "b":
			return this__766__auto__.B

		case "a":
			return this__766__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k4715, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testC) CljsCoreIPrintWithWriter__() {}

func (this__780__auto__ *CljsCore_testC) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__5541 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5541, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, cljs_core.Pr_writer, "", " ", "", opts__782__auto__, keyval__784__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__783__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, pr_pair__783__auto__, "#cljs.core-test.C{", ", ", "}", opts__782__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__780__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__780__auto__.B}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), this__780__auto__.C}, nil})}, nil}), this__780__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testC) CljsCoreIMeta__() {}

func (this__762__auto__ *CljsCore_testC) X_meta_Arity1() interface{} {
	return this__762__auto__.X__meta
}

func (_ *CljsCore_testC) CljsCoreICloneable__() {}

func (this__758__auto__ *CljsCore_testC) X_clone_Arity1() interface{} {
	return (&CljsCore_testC{this__758__auto__.A, this__758__auto__.B, this__758__auto__.C, this__758__auto__.X__meta, this__758__auto__.X__extmap, this__758__auto__.X__hash})
}

func (_ *CljsCore_testC) CljsCoreICounted__() {}

func (this__768__auto__ *CljsCore_testC) X_count_Arity1() float64 {
	return (float64(3) + cljs_core.Count.X_invoke_Arity1(this__768__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testC) CljsCoreIHash__() {}

func (this__759__auto__ *CljsCore_testC) X_hash_Arity1() interface{} {
	{
		var h__583__auto__ = this__759__auto__.X__hash
		_ = h__583__auto__
		if !(cljs_core.Nil_(h__583__auto__)) {
			return h__583__auto__
		} else {
			{
				var h__583__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__583__auto_____1
				this__759__auto__.X__hash = h__583__auto_____1

				return h__583__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testC) CljsCoreIEquiv__() {}

func (this__760__auto__ *CljsCore_testC) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__165__auto__ = other__761__auto__
		_ = and__165__auto__
		if cljs_core.Truth_(and__165__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
		} else {
			return and__165__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testC) CljsCoreIRecord__() {}

func (_ *CljsCore_testC) CljsCoreIMap__() {}

func (this__775__auto__ *CljsCore_testC) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testC{this__775__auto__.A, this__775__auto__.B, this__775__auto__.C, this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testC) CljsCoreIAssociative__() {}

func (this__771__auto__ *CljsCore_testC) X_assoc_Arity3(k__772__auto__ interface{}, G__4714 interface{}) interface{} {
	{
		var pred__4728 = cljs_core.Keyword_identical_QMARK_
		var expr__4729 = k__772__auto__
		_, _ = pred__4728, expr__4729
		if cljs_core.Truth_(func() interface{} {
			var G__4731 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__4732 = expr__4729
			_, _ = G__4731, G__4732
			return pred__4728.X_invoke_Arity2(G__4731, G__4732)
		}()) {
			return (&CljsCore_testC{G__4714, this__771__auto__.B, this__771__auto__.C, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__4733 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__4734 = expr__4729
				_, _ = G__4733, G__4734
				return pred__4728.X_invoke_Arity2(G__4733, G__4734)
			}()) {
				return (&CljsCore_testC{this__771__auto__.A, G__4714, this__771__auto__.C, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
			} else {
				if cljs_core.Truth_(func() interface{} {
					var G__4735 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})
					var G__4736 = expr__4729
					_, _ = G__4735, G__4736
					return pred__4728.X_invoke_Arity2(G__4735, G__4736)
				}()) {
					return (&CljsCore_testC{this__771__auto__.A, this__771__auto__.B, G__4714, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
				} else {
					return (&CljsCore_testC{this__771__auto__.A, this__771__auto__.B, this__771__auto__.C, this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__4714), nil})
				}
			}
		}
	}
}

func (this__773__auto__ *CljsCore_testC) X_contains_key_QMARK__Arity2(k__774__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testC) CljsCoreISeqable__() {}

func (this__778__auto__ *CljsCore_testC) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__778__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__778__auto__.B}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), this__778__auto__.C}, nil})}, nil}), this__778__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testC) CljsCoreIWithMeta__() {}

func (this__763__auto__ *CljsCore_testC) X_with_meta_Arity2(G__4714 interface{}) interface{} {
	return (&CljsCore_testC{this__763__auto__.A, this__763__auto__.B, this__763__auto__.C, G__4714, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testC) CljsCoreICollection__() {}

func (this__769__auto__ *CljsCore_testC) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__769__auto__, entry__770__auto__)
	}
}

var X__GT_C *cljs_core.AFn

var Map__GT_C *cljs_core.AFn

var Letters *CljsCore_testC

var More_letters interface{}

type CljsCore_testA2 struct {
	X         interface{}
	X__meta   interface{}
	X__extmap interface{}
	X__hash   interface{}
}

func (_ *CljsCore_testA2) CljsCoreILookup__() {}

func (this__764__auto__ *CljsCore_testA2) X_lookup_Arity2(k__765__auto__ interface{}) interface{} {
	return this__764__auto__.X_lookup_Arity3(k__765__auto__, nil)
}

func (this__766__auto__ *CljsCore_testA2) X_lookup_Arity3(k4740 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__4743 = func() interface{} {
			if cljs_core.Value_(k4740).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k4740), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4743
		switch G__4743 {
		case "x":
			return this__766__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k4740, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIPrintWithWriter__() {}

func (this__780__auto__ *CljsCore_testA2) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__5543 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5543, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, cljs_core.Pr_writer, "", " ", "", opts__782__auto__, keyval__784__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__783__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, pr_pair__783__auto__, "#cljs.core-test.A2{", ", ", "}", opts__782__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__780__auto__.X}, nil})}, nil}), this__780__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testA2) CljsCoreIMeta__() {}

func (this__762__auto__ *CljsCore_testA2) X_meta_Arity1() interface{} {
	return this__762__auto__.X__meta
}

func (_ *CljsCore_testA2) CljsCoreICloneable__() {}

func (this__758__auto__ *CljsCore_testA2) X_clone_Arity1() interface{} {
	return (&CljsCore_testA2{this__758__auto__.X, this__758__auto__.X__meta, this__758__auto__.X__extmap, this__758__auto__.X__hash})
}

func (_ *CljsCore_testA2) CljsCoreICounted__() {}

func (this__768__auto__ *CljsCore_testA2) X_count_Arity1() float64 {
	return (float64(1) + cljs_core.Count.X_invoke_Arity1(this__768__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testA2) CljsCoreIHash__() {}

func (this__759__auto__ *CljsCore_testA2) X_hash_Arity1() interface{} {
	{
		var h__583__auto__ = this__759__auto__.X__hash
		_ = h__583__auto__
		if !(cljs_core.Nil_(h__583__auto__)) {
			return h__583__auto__
		} else {
			{
				var h__583__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__583__auto_____1
				this__759__auto__.X__hash = h__583__auto_____1

				return h__583__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIEquiv__() {}

func (this__760__auto__ *CljsCore_testA2) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__165__auto__ = other__761__auto__
		_ = and__165__auto__
		if cljs_core.Truth_(and__165__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
		} else {
			return and__165__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testA2) CljsCoreIRecord__() {}

func (_ *CljsCore_testA2) CljsCoreIMap__() {}

func (this__775__auto__ *CljsCore_testA2) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), nil}, nil}, nil}), k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testA2{this__775__auto__.X, this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testA2) CljsCoreIAssociative__() {}

func (this__771__auto__ *CljsCore_testA2) X_assoc_Arity3(k__772__auto__ interface{}, G__4739 interface{}) interface{} {
	{
		var pred__4749 = cljs_core.Keyword_identical_QMARK_
		var expr__4750 = k__772__auto__
		_, _ = pred__4749, expr__4750
		if cljs_core.Truth_(func() interface{} {
			var G__4752 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__4753 = expr__4750
			_, _ = G__4752, G__4753
			return pred__4749.X_invoke_Arity2(G__4752, G__4753)
		}()) {
			return (&CljsCore_testA2{G__4739, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testA2{this__771__auto__.X, this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__4739), nil})
		}
	}
}

func (this__773__auto__ *CljsCore_testA2) X_contains_key_QMARK__Arity2(k__774__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testA2) CljsCoreISeqable__() {}

func (this__778__auto__ *CljsCore_testA2) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__778__auto__.X}, nil})}, nil}), this__778__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testA2) CljsCoreIWithMeta__() {}

func (this__763__auto__ *CljsCore_testA2) X_with_meta_Arity2(G__4739 interface{}) interface{} {
	return (&CljsCore_testA2{this__763__auto__.X, G__4739, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testA2) CljsCoreICollection__() {}

func (this__769__auto__ *CljsCore_testA2) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__769__auto__, entry__770__auto__)
	}
}

var X__GT_A2 *cljs_core.AFn

var Map__GT_A2 *cljs_core.AFn

type CljsCore_testB struct {
	X         interface{}
	X__meta   interface{}
	X__extmap interface{}
	X__hash   interface{}
}

func (_ *CljsCore_testB) CljsCoreILookup__() {}

func (this__764__auto__ *CljsCore_testB) X_lookup_Arity2(k__765__auto__ interface{}) interface{} {
	return this__764__auto__.X_lookup_Arity3(k__765__auto__, nil)
}

func (this__766__auto__ *CljsCore_testB) X_lookup_Arity3(k4755 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__4758 = func() interface{} {
			if cljs_core.Value_(k4755).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k4755), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4758
		switch G__4758 {
		case "x":
			return this__766__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k4755, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testB) CljsCoreIPrintWithWriter__() {}

func (this__780__auto__ *CljsCore_testB) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__5545 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5545, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, cljs_core.Pr_writer, "", " ", "", opts__782__auto__, keyval__784__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__783__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, pr_pair__783__auto__, "#cljs.core-test.B{", ", ", "}", opts__782__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__780__auto__.X}, nil})}, nil}), this__780__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testB) CljsCoreIMeta__() {}

func (this__762__auto__ *CljsCore_testB) X_meta_Arity1() interface{} {
	return this__762__auto__.X__meta
}

func (_ *CljsCore_testB) CljsCoreICloneable__() {}

func (this__758__auto__ *CljsCore_testB) X_clone_Arity1() interface{} {
	return (&CljsCore_testB{this__758__auto__.X, this__758__auto__.X__meta, this__758__auto__.X__extmap, this__758__auto__.X__hash})
}

func (_ *CljsCore_testB) CljsCoreICounted__() {}

func (this__768__auto__ *CljsCore_testB) X_count_Arity1() float64 {
	return (float64(1) + cljs_core.Count.X_invoke_Arity1(this__768__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testB) CljsCoreIHash__() {}

func (this__759__auto__ *CljsCore_testB) X_hash_Arity1() interface{} {
	{
		var h__583__auto__ = this__759__auto__.X__hash
		_ = h__583__auto__
		if !(cljs_core.Nil_(h__583__auto__)) {
			return h__583__auto__
		} else {
			{
				var h__583__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__583__auto_____1
				this__759__auto__.X__hash = h__583__auto_____1

				return h__583__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testB) CljsCoreIEquiv__() {}

func (this__760__auto__ *CljsCore_testB) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__165__auto__ = other__761__auto__
		_ = and__165__auto__
		if cljs_core.Truth_(and__165__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
		} else {
			return and__165__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testB) CljsCoreIRecord__() {}

func (_ *CljsCore_testB) CljsCoreIMap__() {}

func (this__775__auto__ *CljsCore_testB) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), nil}, nil}, nil}), k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testB{this__775__auto__.X, this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testB) CljsCoreIAssociative__() {}

func (this__771__auto__ *CljsCore_testB) X_assoc_Arity3(k__772__auto__ interface{}, G__4754 interface{}) interface{} {
	{
		var pred__4764 = cljs_core.Keyword_identical_QMARK_
		var expr__4765 = k__772__auto__
		_, _ = pred__4764, expr__4765
		if cljs_core.Truth_(func() interface{} {
			var G__4767 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__4768 = expr__4765
			_, _ = G__4767, G__4768
			return pred__4764.X_invoke_Arity2(G__4767, G__4768)
		}()) {
			return (&CljsCore_testB{G__4754, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testB{this__771__auto__.X, this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__4754), nil})
		}
	}
}

func (this__773__auto__ *CljsCore_testB) X_contains_key_QMARK__Arity2(k__774__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testB) CljsCoreISeqable__() {}

func (this__778__auto__ *CljsCore_testB) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__778__auto__.X}, nil})}, nil}), this__778__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testB) CljsCoreIWithMeta__() {}

func (this__763__auto__ *CljsCore_testB) X_with_meta_Arity2(G__4754 interface{}) interface{} {
	return (&CljsCore_testB{this__763__auto__.X, G__4754, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testB) CljsCoreICollection__() {}

func (this__769__auto__ *CljsCore_testB) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__769__auto__, entry__770__auto__)
	}
}

var X__GT_B *cljs_core.AFn

var Map__GT_B *cljs_core.AFn

type CljsCore_testIFoo interface {
	CljsCore_testIFoo__()
	Foo_Arity1() interface{}
}

var Foo *cljs_core.AFn

type CljsCore_testT4769 struct {
	Test_stuff interface{}
	Meta4770   interface{}
}

func (_ *CljsCore_testT4769) CljsCore_testIFoo__() {}

func (this *CljsCore_testT4769) Foo_Arity1() interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
}

func (_ *CljsCore_testT4769) CljsCoreIMeta__() {}

func (_4771 *CljsCore_testT4769) X_meta_Arity1() interface{} {
	return _4771.Meta4770
}

func (_ *CljsCore_testT4769) CljsCoreIWithMeta__() {}

func (_4771 *CljsCore_testT4769) X_with_meta_Arity2(meta4770___1 interface{}) interface{} {
	return (&CljsCore_testT4769{_4771.Test_stuff, meta4770___1})
}

var X__GT_t4769 *cljs_core.AFn

var Foo2 *cljs_core.CljsCoreMultiFn

type CljsCore_testIMutate interface {
	CljsCore_testIMutate__()
	Mutate_Arity1() interface{}
}

var Mutate *cljs_core.AFn

type CljsCore_testMutate struct{ A interface{} }

func (_ *CljsCore_testMutate) CljsCore_testIMutate__() {}

func (___ *CljsCore_testMutate) Mutate_Arity1() interface{} {
	return func() interface{} {
		var return__5546 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})
		___.A = return__5546
		return return__5546
	}()
}

var X__GT_Mutate *cljs_core.AFn

type CljsCore_testFnLike struct{}

func (_ *CljsCore_testFnLike) CljsCoreIFn__() {}

func (___ *CljsCore_testFnLike) X_invoke_Arity0() interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
}

func (___ *CljsCore_testFnLike) X_invoke_Arity1(a interface{}) interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
}

func (___ *CljsCore_testFnLike) X_invoke_Arity2(a interface{}, b interface{}) interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})
}

func (this *CljsCore_testFnLike) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_FnLike *cljs_core.AFn

type CljsCore_testFnLikeB struct{ A interface{} }

func (_ *CljsCore_testFnLikeB) CljsCoreIFn__() {}

func (___ *CljsCore_testFnLikeB) X_invoke_Arity0() interface{} {
	return ___.A
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity1(a___1 interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 1"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity2(a___1 interface{}, b interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 2"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity3(a___1 interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity4(a___1 interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity5(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity6(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity7(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity8(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity9(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity10(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity11(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity12(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity13(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity14(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity15(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity16(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity17(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity18(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity19(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity20(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_FnLikeB *cljs_core.AFn

type CljsCore_testIHasFirst interface {
	CljsCore_testIHasFirst__()
	X_get_first_Arity1() interface{}
}

var X_get_first *cljs_core.AFn

type CljsCore_testIFindsFirst interface {
	CljsCore_testIFindsFirst__()
	X_find_first_Arity2(other interface{}) interface{}
}

var X_find_first *cljs_core.AFn

type CljsCore_testFirst struct{ Xs interface{} }

func (_ *CljsCore_testFirst) CljsCore_testIFindsFirst__() {}

func (___ *CljsCore_testFirst) X_find_first_Arity2(p__4772 interface{}) interface{} {
	{
		var vec__4774 = p__4772
		var x = cljs_core.Nth.X_invoke_Arity3(vec__4774, float64(0), nil)
		_, _ = vec__4774, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCore_testIHasFirst__() {}

func (p__4775 *CljsCore_testFirst) X_get_first_Arity1() interface{} {
	{
		var vec__4777 = p__4775
		var x = cljs_core.Nth.X_invoke_Arity3(vec__4777, float64(0), nil)
		_, _ = vec__4777, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCoreObject__() {}

func (p__4778 *CljsCore_testFirst) ToString() string {
	{
		var vec__4780 = p__4778
		var x = cljs_core.Nth.X_invoke_Arity3(vec__4780, float64(0), nil)
		_, _ = vec__4780, x
		return (`` + cljs_core.Str.X_invoke_Arity1(x).(string))
	}
}

func (this *CljsCore_testFirst) String() string {
	return this.ToString()
}

func (_ *CljsCore_testFirst) CljsCoreIFn__() {}

func (p__4781 *CljsCore_testFirst) X_invoke_Arity0() interface{} {
	{
		var vec__4783 = p__4781
		var x = cljs_core.Nth.X_invoke_Arity3(vec__4783, float64(0), nil)
		_, _ = vec__4783, x
		return x
	}
}

func (this *CljsCore_testFirst) X_invoke_Arity1(x interface{}) interface{} {
	return this
}

func (this *CljsCore_testFirst) X_invoke_Arity2(a interface{}, b interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 2"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity3(a interface{}, b interface{}, c interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 3"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity4(a interface{}, b interface{}, c interface{}, d interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 4"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity5(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 5"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity6(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 6"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity7(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 7"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity8(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 8"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity9(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 9"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity10(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 10"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity11(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 11"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity12(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 12"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity13(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 13"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity14(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 14"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity15(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 15"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity16(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 16"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity17(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 17"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, r interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

func (_ *CljsCore_testFirst) CljsCoreIIndexed__() {}

func (this *CljsCore_testFirst) X_nth_Arity2(i interface{}) interface{} {
	return cljs_core.Nth.X_invoke_Arity2(this.Xs, i)
}

func (this *CljsCore_testFirst) X_nth_Arity3(i interface{}, not_found interface{}) interface{} {
	return cljs_core.Nth.X_invoke_Arity3(this.Xs, i, not_found)
}

func (_ *CljsCore_testFirst) CljsCoreISeqable__() {}

func (this *CljsCore_testFirst) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(this.Xs)
}

var X__GT_First *cljs_core.AFn

type CljsCore_testDestructuringWithLocals struct{ A interface{} }

func (_ *CljsCore_testDestructuringWithLocals) CljsCore_testIFindsFirst__() {}

func (___ *CljsCore_testDestructuringWithLocals) X_find_first_Arity2(p__4785 interface{}) interface{} {
	{
		var vec__4787 = p__4785
		var x = cljs_core.Nth.X_invoke_Arity3(vec__4787, float64(0), nil)
		var y = cljs_core.Nth.X_invoke_Arity3(vec__4787, float64(1), nil)
		_, _, _ = vec__4787, x, y
		return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{x, y, ___.A}, nil})
	}
}

var X__GT_DestructuringWithLocals *cljs_core.AFn

type CljsCore_testPrintMe struct {
	A         interface{}
	B         interface{}
	X__meta   interface{}
	X__extmap interface{}
	X__hash   interface{}
}

func (_ *CljsCore_testPrintMe) CljsCoreILookup__() {}

func (this__764__auto__ *CljsCore_testPrintMe) X_lookup_Arity2(k__765__auto__ interface{}) interface{} {
	return this__764__auto__.X_lookup_Arity3(k__765__auto__, nil)
}

func (this__766__auto__ *CljsCore_testPrintMe) X_lookup_Arity3(k4799 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__4802 = func() interface{} {
			if cljs_core.Value_(k4799).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k4799), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4802
		switch G__4802 {
		case "b":
			return this__766__auto__.B

		case "a":
			return this__766__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k4799, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIPrintWithWriter__() {}

func (this__780__auto__ *CljsCore_testPrintMe) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__5548 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5548, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, cljs_core.Pr_writer, "", " ", "", opts__782__auto__, keyval__784__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__783__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__781__auto__, pr_pair__783__auto__, "#cljs.core-test.PrintMe{", ", ", "}", opts__782__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__780__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__780__auto__.B}, nil})}, nil}), this__780__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIMeta__() {}

func (this__762__auto__ *CljsCore_testPrintMe) X_meta_Arity1() interface{} {
	return this__762__auto__.X__meta
}

func (_ *CljsCore_testPrintMe) CljsCoreICloneable__() {}

func (this__758__auto__ *CljsCore_testPrintMe) X_clone_Arity1() interface{} {
	return (&CljsCore_testPrintMe{this__758__auto__.A, this__758__auto__.B, this__758__auto__.X__meta, this__758__auto__.X__extmap, this__758__auto__.X__hash})
}

func (_ *CljsCore_testPrintMe) CljsCoreICounted__() {}

func (this__768__auto__ *CljsCore_testPrintMe) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__768__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testPrintMe) CljsCoreIHash__() {}

func (this__759__auto__ *CljsCore_testPrintMe) X_hash_Arity1() interface{} {
	{
		var h__583__auto__ = this__759__auto__.X__hash
		_ = h__583__auto__
		if !(cljs_core.Nil_(h__583__auto__)) {
			return h__583__auto__
		} else {
			{
				var h__583__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__583__auto_____1
				this__759__auto__.X__hash = h__583__auto_____1

				return h__583__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIEquiv__() {}

func (this__760__auto__ *CljsCore_testPrintMe) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__165__auto__ = other__761__auto__
		_ = and__165__auto__
		if cljs_core.Truth_(and__165__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
		} else {
			return and__165__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIRecord__() {}

func (_ *CljsCore_testPrintMe) CljsCoreIMap__() {}

func (this__775__auto__ *CljsCore_testPrintMe) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testPrintMe{this__775__auto__.A, this__775__auto__.B, this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIAssociative__() {}

func (this__771__auto__ *CljsCore_testPrintMe) X_assoc_Arity3(k__772__auto__ interface{}, G__4798 interface{}) interface{} {
	{
		var pred__4810 = cljs_core.Keyword_identical_QMARK_
		var expr__4811 = k__772__auto__
		_, _ = pred__4810, expr__4811
		if cljs_core.Truth_(func() interface{} {
			var G__4813 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__4814 = expr__4811
			_, _ = G__4813, G__4814
			return pred__4810.X_invoke_Arity2(G__4813, G__4814)
		}()) {
			return (&CljsCore_testPrintMe{G__4798, this__771__auto__.B, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__4815 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__4816 = expr__4811
				_, _ = G__4815, G__4816
				return pred__4810.X_invoke_Arity2(G__4815, G__4816)
			}()) {
				return (&CljsCore_testPrintMe{this__771__auto__.A, G__4798, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPrintMe{this__771__auto__.A, this__771__auto__.B, this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__4798), nil})
			}
		}
	}
}

func (this__773__auto__ *CljsCore_testPrintMe) X_contains_key_QMARK__Arity2(k__774__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testPrintMe) CljsCoreISeqable__() {}

func (this__778__auto__ *CljsCore_testPrintMe) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__778__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__778__auto__.B}, nil})}, nil}), this__778__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testPrintMe) CljsCoreIWithMeta__() {}

func (this__763__auto__ *CljsCore_testPrintMe) X_with_meta_Arity2(G__4798 interface{}) interface{} {
	return (&CljsCore_testPrintMe{this__763__auto__.A, this__763__auto__.B, G__4798, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testPrintMe) CljsCoreICollection__() {}

func (this__769__auto__ *CljsCore_testPrintMe) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__770__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__769__auto__, entry__770__auto__)
	}
}

var X__GT_PrintMe *cljs_core.AFn

var Map__GT_PrintMe *cljs_core.AFn

type CljsCore_testIBar interface {
	CljsCore_testIBar__()
	X_bar_Arity2(x interface{}) interface{}
}

var X_bar *cljs_core.AFn

var Baz *cljs_core.AFn

type CljsCore_testT4850 struct {
	F          interface{}
	Baz        interface{}
	Test_stuff interface{}
	Meta4851   interface{}
}

func (_ *CljsCore_testT4850) CljsCore_testIBar__() {}

func (___ *CljsCore_testT4850) X_bar_Arity2(x interface{}) interface{} {
	{
		var G__4854 = x
		_ = G__4854
		return ___.F.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4854)
	}
}

func (_ *CljsCore_testT4850) CljsCoreIMeta__() {}

func (_4852 *CljsCore_testT4850) X_meta_Arity1() interface{} {
	return _4852.Meta4851
}

func (_ *CljsCore_testT4850) CljsCoreIWithMeta__() {}

func (_4852 *CljsCore_testT4850) X_with_meta_Arity2(meta4851___1 interface{}) interface{} {
	return (&CljsCore_testT4850{_4852.F, _4852.Baz, _4852.Test_stuff, meta4851___1})
}

var X__GT_t4850 *cljs_core.AFn

var Original_closure_stmt *cljs_core.AFn

type CljsCore_testPositionalFactoryTest struct{ X interface{} }

var X__GT_PositionalFactoryTest *cljs_core.AFn

var Foo580 interface{}

type CljsCore_testKeywordTest struct{}

func (_ *CljsCore_testKeywordTest) CljsCoreILookup__() {}

func (o *CljsCore_testKeywordTest) X_lookup_Arity2(k interface{}) interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nothing", Fqn: "nothing", X_hash: float64(-1022703296)})
}

func (o *CljsCore_testKeywordTest) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	return not_found
}

var X__GT_KeywordTest *cljs_core.AFn

type CljsCore_testT4864 struct {
	Test_stuff interface{}
	Meta4865   interface{}
}

func (_ *CljsCore_testT4864) CljsCoreIHash__() {}

func (___ *CljsCore_testT4864) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT4864) CljsCoreIMeta__() {}

func (_4866 *CljsCore_testT4864) X_meta_Arity1() interface{} {
	return _4866.Meta4865
}

func (_ *CljsCore_testT4864) CljsCoreIWithMeta__() {}

func (_4866 *CljsCore_testT4864) X_with_meta_Arity2(meta4865___1 interface{}) interface{} {
	return (&CljsCore_testT4864{_4866.Test_stuff, meta4865___1})
}

var X__GT_t4864 *cljs_core.AFn

type CljsCore_testT4867 struct {
	A          interface{}
	Test_stuff interface{}
	Meta4868   interface{}
}

func (_ *CljsCore_testT4867) CljsCoreIHash__() {}

func (___ *CljsCore_testT4867) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT4867) CljsCoreIMeta__() {}

func (_4869 *CljsCore_testT4867) X_meta_Arity1() interface{} {
	return _4869.Meta4868
}

func (_ *CljsCore_testT4867) CljsCoreIWithMeta__() {}

func (_4869 *CljsCore_testT4867) X_with_meta_Arity2(meta4868___1 interface{}) interface{} {
	return (&CljsCore_testT4867{_4869.A, _4869.Test_stuff, meta4868___1})
}

var X__GT_t4867 *cljs_core.AFn

var Some_x float64

var Some_y float64

type CljsCore_testIWoz interface {
	CljsCore_testIWoz__()
	X_woz_Arity1() interface{}
}

var X_woz *cljs_core.AFn

var Noz cljs_core.CljsCoreIVector

var Cljs_739 *cljs_core.AFn

var Cljs_780 *cljs_core.CljsCoreAtom

type CljsCore_testT4937 struct {
	From_seq    interface{}
	Make_seq    interface{}
	Mt          interface{}
	I__4929     interface{}
	Count__4928 interface{}
	Chunk__4927 interface{}
	Seq__4926   interface{}
	Test_stuff  interface{}
	Meta4938    interface{}
}

func (_ *CljsCore_testT4937) CljsCoreISeq__() {}

func (this *CljsCore_testT4937) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT4937) X_rest_Arity1() interface{} {
	{
		var G__4941 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__4941
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4941)
	}
}

func (_ *CljsCore_testT4937) CljsCoreISeqable__() {}

func (this *CljsCore_testT4937) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT4937) CljsCoreIMeta__() {}

func (_4939 *CljsCore_testT4937) X_meta_Arity1() interface{} {
	return _4939.Meta4938
}

func (_ *CljsCore_testT4937) CljsCoreIWithMeta__() {}

func (_4939 *CljsCore_testT4937) X_with_meta_Arity2(meta4938___1 interface{}) interface{} {
	return (&CljsCore_testT4937{_4939.From_seq, _4939.Make_seq, _4939.Mt, _4939.I__4929, _4939.Count__4928, _4939.Chunk__4927, _4939.Seq__4926, _4939.Test_stuff, meta4938___1})
}

var X__GT_t4937 *cljs_core.AFn

type CljsCore_testT4949 struct {
	From_seq           interface{}
	Make_seq           interface{}
	Mt                 interface{}
	Temp__4222__auto__ interface{}
	I__4929            interface{}
	Count__4928        interface{}
	Chunk__4927        interface{}
	Seq__4926          interface{}
	Test_stuff         interface{}
	Meta4950           interface{}
}

func (_ *CljsCore_testT4949) CljsCoreISeq__() {}

func (this *CljsCore_testT4949) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT4949) X_rest_Arity1() interface{} {
	{
		var G__4953 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__4953
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4953)
	}
}

func (_ *CljsCore_testT4949) CljsCoreISeqable__() {}

func (this *CljsCore_testT4949) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT4949) CljsCoreIMeta__() {}

func (_4951 *CljsCore_testT4949) X_meta_Arity1() interface{} {
	return _4951.Meta4950
}

func (_ *CljsCore_testT4949) CljsCoreIWithMeta__() {}

func (_4951 *CljsCore_testT4949) X_with_meta_Arity2(meta4950___1 interface{}) interface{} {
	return (&CljsCore_testT4949{_4951.From_seq, _4951.Make_seq, _4951.Mt, _4951.Temp__4222__auto__, _4951.I__4929, _4951.Count__4928, _4951.Chunk__4927, _4951.Seq__4926, _4951.Test_stuff, meta4950___1})
}

var X__GT_t4949 *cljs_core.AFn

var Case_recur *cljs_core.AFn

var Xform *cljs_core.AFn

var Data interface{}

var Xf *cljs_core.AFn

func Test_runner(t *testing.T) {
	assert.Equal(t, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), Test_stuff.X_invoke_Arity0().(*cljs_core.CljsCoreKeyword))
}
