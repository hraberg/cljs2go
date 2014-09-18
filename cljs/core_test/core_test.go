// Compiled by ClojureScript to Go 0.0-2341
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(3), float64(2), float64(1), float64(0)}, nil}), func() *cljs_core.CljsCoreLazySeq {
				var i = float64(0)
				var j interface{} = cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)
				_, _ = i, j
				for {
					if i < float64(5) {
						i, j = (i + float64(1)), cljs_core.Conj.X_invoke_Arity2(j, func(G__4963 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__4963, 0, func() interface{} {
								return i
							})
						}(&cljs_core.AFn{}, i, j))
						continue
					} else {
						return cljs_core.Map_.X_invoke_Arity2(func(G__4964 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__4964, 1, func(p1__4078_SHARP_ interface{}) interface{} {
								{
									return p1__4078_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
								}
							})
						}(&cljs_core.AFn{}, i, j), j).(*cljs_core.CljsCoreLazySeq)
					}
				}
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [4 3 2 1 0] (loop [i 0 j ()] (if (< i 5) (recur (inc i) (conj j (fn [] i))) (map (fn* [p1__4078#] (p1__4078#)) j))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})}, nil}), cljs_core.Map_.X_invoke_Arity2(func(G__4965 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4965, 1, func(p1__4079_SHARP_ interface{}) interface{} {
					{
						return p1__4079_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
					}
				})
			}(&cljs_core.AFn{}), func() *cljs_core.CljsCoreLazySeq {
				var iter__923__auto__ = func(iter__4539 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(iter__4539, 1, func(s__4540 interface{}) interface{} {
						return (&cljs_core.CljsCoreLazySeq{nil, func(G__4966 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__4966, 0, func() interface{} {
								{
									var s__4540___1 interface{} = s__4540
									_ = s__4540___1
									for {
										{
											var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4540___1)
											_ = temp__4222__auto__
											if cljs_core.Truth_(temp__4222__auto__) {
												{
													var xs__4752__auto__ = temp__4222__auto__
													_ = xs__4752__auto__
													{
														var i = cljs_core.First.X_invoke_Arity1(xs__4752__auto__)
														_ = i
														{
															var iterys__919__auto__ = func(iter__4541 *cljs_core.AFn, s__4540___1 interface{}, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																return cljs_core.Fn(iter__4541, 1, func(s__4542 interface{}) interface{} {
																	return (&cljs_core.CljsCoreLazySeq{nil, func(G__4967 *cljs_core.AFn, s__4540___1 interface{}, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																		return cljs_core.Fn(G__4967, 0, func() interface{} {
																			{
																				var s__4542___1 interface{} = s__4542
																				_ = s__4542___1
																				for {
																					{
																						var temp__4222__auto_____1 = cljs_core.Seq.Arity1IQ(s__4542___1)
																						_ = temp__4222__auto_____1
																						if cljs_core.Truth_(temp__4222__auto_____1) {
																							{
																								var s__4542___2 = temp__4222__auto_____1
																								_ = s__4542___2
																								if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4542___2) {
																									{
																										var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4542___2)
																										var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																										var b__4544 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																										_, _, _ = c__921__auto__, size__922__auto__, b__4544
																										if func() bool {
																											var i__4543 = float64(0)
																											_ = i__4543
																											for {
																												if i__4543 < size__922__auto__ {
																													{
																														var j = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4543)
																														_ = j
																														cljs_core.Chunk_append.X_invoke_Arity2(b__4544, func(G__4968 *cljs_core.AFn, i__4543 float64, s__4540___1 interface{}, j interface{}, c__921__auto__ interface{}, size__922__auto__ float64, b__4544 *cljs_core.CljsCoreChunkBuffer, s__4542___2 interface{}, temp__4222__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																															return cljs_core.Fn(G__4968, 0, func() interface{} {
																																return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																															})
																														}(&cljs_core.AFn{}, i__4543, s__4540___1, j, c__921__auto__, size__922__auto__, b__4544, s__4542___2, temp__4222__auto_____1, i, xs__4752__auto__, temp__4222__auto__))
																														i__4543 = (i__4543 + float64(1))
																														continue
																													}
																												} else {
																													return true
																												}
																											}
																										}() {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4544), iter__4541.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4542___2)).(*cljs_core.CljsCoreLazySeq))
																										} else {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4544), nil)
																										}
																									}
																								} else {
																									{
																										var j = cljs_core.First.X_invoke_Arity1(s__4542___2)
																										_ = j
																										return cljs_core.Cons.X_invoke_Arity2(func(G__4969 *cljs_core.AFn, s__4540___1 interface{}, j interface{}, s__4542___2 interface{}, temp__4222__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																											return cljs_core.Fn(G__4969, 0, func() interface{} {
																												return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																											})
																										}(&cljs_core.AFn{}, s__4540___1, j, s__4542___2, temp__4222__auto_____1, i, xs__4752__auto__, temp__4222__auto__), iter__4541.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4542___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
																	}(&cljs_core.AFn{}, s__4540___1, i, xs__4752__auto__, temp__4222__auto__), nil, nil})
																})
															}(&cljs_core.AFn{}, s__4540___1, i, xs__4752__auto__, temp__4222__auto__)
															var fs__920__auto__ = cljs_core.Seq.Arity1IQ(iterys__919__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq))
															_, _ = iterys__919__auto__, fs__920__auto__
															if cljs_core.Truth_(fs__920__auto__) {
																return cljs_core.Concat.X_invoke_Arity2(fs__920__auto__, iter__4539.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4540___1)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
															} else {
																s__4540___1 = cljs_core.Rest.Arity1IQ(s__4540___1)
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
				_ = iter__923__auto__
				return iter__923__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})).(*cljs_core.CljsCoreLazySeq)
			}()).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [[1 1] [1 2] [1 3] [2 1] [2 2] [2 3]] (map (fn* [p1__4079#] (p1__4079#)) (for [i [1 2] j [1 2 3]] (fn [] [i j]))))")}))
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
				var sb__1124__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1124__auto__
				{
					var _STAR_print_fn_STAR_4550_4970 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_4550_4970
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_4550_4970

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__4971 *cljs_core.AFn, _STAR_print_fn_STAR_4550_4970 interface{}, sb__1124__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__4971, 1, func(x__1125__auto__ interface{}) interface{} {
									return sb__1124__auto__.Append(x__1125__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_4550_4970, sb__1124__auto__)

							cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)}))
							cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)}))
						}
					}()
				}
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1124__auto__).(string))
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= \"12\" (with-out-str (print 1) (print 2)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("12", func() string {
				var sb__1124__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1124__auto__
				{
					var _STAR_print_fn_STAR_4551_4972 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_4551_4972
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_4551_4972

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__4973 *cljs_core.AFn, _STAR_print_fn_STAR_4551_4972 interface{}, sb__1124__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__4973, 1, func(x__1125__auto__ interface{}) interface{} {
									return sb__1124__auto__.Append(x__1125__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_4551_4972, sb__1124__auto__)

							cljs_core.X_STAR_print_fn_STAR_.X_invoke_Arity1(float64(1))
							cljs_core.X_STAR_print_fn_STAR_.X_invoke_Arity1(float64(2))
						}
					}()
				}
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1124__auto__).(string))
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
				var jumble = func(G__4974 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4974, 2, func(a interface{}, b interface{}) interface{} {
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
				var jumble = func(G__4975 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4975, 2, func(a interface{}, b interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Apply.X_invoke_Arity2(func(G__4976 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4976, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return ((cljs_core.Nth.X_invoke_Arity2(args, float64(0)).(float64) + cljs_core.Nth.X_invoke_Arity2(args, float64(1)).(float64)) + cljs_core.Nth.X_invoke_Arity2(args, float64(2)).(float64))
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))) {
			} else {
				panic((&js.Error{("Assert failed: (= 3 (apply (fn [& args] (+ (nth args 0) (nth args 1) (nth args 2))) (iterate inc 0)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4977 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4977, 0, func(m__ ...interface{}) interface{} {
					var m = cljs_core.Seq.Arity1IQ(m__[0])
					_ = m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [0 1 2 3 4] (take 5 (apply (fn [& m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4978 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4978, 1, func(x_m__ ...interface{}) interface{} {
					var x = x_m__[0]
					var m = cljs_core.Seq.Arity1IQ(x_m__[1])
					_, _ = x, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (take 5 (apply (fn [x & m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4979 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4979, 2, func(x_y_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5), float64(6), float64(7)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4980 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4980, 3, func(x_y_z_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6), float64(7), float64(8)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4981 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4981, 4, func(x_y_z_a_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__4982 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__4982, 5, func(x_y_z_a_b_m__ ...interface{}) interface{} {
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
				var single_arity_non_variadic_4983 = func(G__4988 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4988, 3, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{})
				var multiple_arity_non_variadic_4984 = func(G__4989 *cljs_core.AFn, single_arity_non_variadic_4983 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4989, 3, func(x interface{}) interface{} {
						return x
					}, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{y, x}, nil})
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_4983)
				var single_arity_variadic_fixedargs_4985 = func(G__4990 *cljs_core.AFn, single_arity_non_variadic_4983 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_4984 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4990, 2, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{more, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_4983, multiple_arity_non_variadic_4984)
				var single_arity_variadic_nofixedargs_4986 = func(G__4991 *cljs_core.AFn, single_arity_non_variadic_4983 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_4984 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_4985 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4991, 0, func(more__ ...interface{}) interface{} {
						var more = cljs_core.Seq.Arity1IQ(more__[0])
						_ = more
						return more
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_4983, multiple_arity_non_variadic_4984, single_arity_variadic_fixedargs_4985)
				var multiple_arity_variadic_4987 = func(G__4992 *cljs_core.AFn, single_arity_non_variadic_4983 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_4984 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_4985 cljs_core.CljsCoreIFn, single_arity_variadic_nofixedargs_4986 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4992, 2, func(x interface{}) interface{} {
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
				}(&cljs_core.AFn{}, single_arity_non_variadic_4983, multiple_arity_non_variadic_4984, single_arity_variadic_fixedargs_4985, single_arity_variadic_nofixedargs_4986)
				_, _, _, _, _ = single_arity_non_variadic_4983, multiple_arity_non_variadic_4984, single_arity_variadic_fixedargs_4985, single_arity_variadic_nofixedargs_4986, multiple_arity_variadic_4987
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_non_variadic_4983, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic [1 2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_non_variadic_4983, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic 1 [2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_non_variadic_4983, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic 1 2 [3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_4984, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (apply multiple-arity-non-variadic [42]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_4984, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-non-variadic [1 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_4984, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-non-variadic 1 [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_4984, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic [1 2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_4984, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic 1 [2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_non_variadic_4984, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic 1 2 [3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_4985, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_fixedargs_4985, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_fixedargs_4985, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_fixedargs_4985, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_fixedargs_4985, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_4985, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 4 5] (take 3 (first (apply single-arity-variadic-fixedargs (iterate inc 1)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_4985, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (rest (apply single-arity-variadic-fixedargs (iterate inc 1))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_4986, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_nofixedargs_4986, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_nofixedargs_4986, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_nofixedargs_4986, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_nofixedargs_4986, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_4986, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (take 5 (apply single-arity-variadic-nofixedargs (iterate inc 1))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_4987, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (apply multiple-arity-variadic [42]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_4987, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-variadic [1 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_4987, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-variadic 1 [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_4987, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_4987, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_variadic_4987, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(multiple_arity_variadic_4987, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(multiple_arity_variadic_4987, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_4987, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 4 5] (take 3 (first (apply multiple-arity-variadic (iterate inc 1)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_4987, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (rest (apply multiple-arity-variadic (iterate inc 1))))")}))
				}
			}
			{
				var f1_4993 = func(f1 *cljs_core.AFn) *cljs_core.AFn {
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
				var f2_4994 = func(f2 *cljs_core.AFn, f1_4993 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(f2, 2, func(x interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					}, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return cljs_core.Apply.X_invoke_Arity3(f1_4993, y, more)
					})
				}(&cljs_core.AFn{}, f1_4993)
				_, _ = f1_4993, f2_4994
				if cljs_core.X_EQ_.Arity2IIB(float64(1), f2_4994.X_invoke_Arity2(float64(1), float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (f2 1 2))")}))
				}
			}
			{
				var f_4995 = func(G__4996 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__4996, 1, func() interface{} {
						return nil
					}, func(a_more__ ...interface{}) interface{} {
						var a = a_more__[0]
						var more = cljs_core.Seq.Arity1IQ(a_more__[1])
						_, _ = a, more
						return more
					})
				}(&cljs_core.AFn{})
				_ = f_4995
				if cljs_core.Nil_(f_4995.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (f :foo))")}))
				}
			}
			if cljs_core.Nil_(cljs_core.Array_seq.X_invoke_Arity2([]interface{}{float64(1)}, float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (array-seq (array 1) 1))")}))
			}
			{
				var f_4997 = func(G__5000 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5000, 1, func(x interface{}) interface{} {
						return (x.(float64) * float64(2))
					})
				}(&cljs_core.AFn{})
				var m_4998 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar"}, nil})
				var mf_4999 = cljs_core.With_meta.X_invoke_Arity2(f_4997, m_4998)
				_, _, _ = f_4997, m_4998, mf_4999
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(f_4997)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (meta f))")}))
				}
				if cljs_core.Fn_QMARK_.Arity1IB(mf_4999) {
				} else {
					panic((&js.Error{("Assert failed: (fn? mf)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__4552 = float64(2)
					_ = G__4552
					return mf_4999.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4552)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (mf 2))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Apply.X_invoke_Arity2(mf_4999, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (apply mf [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(mf_4999), m_4998) {
				} else {
					panic((&js.Error{("Assert failed: (= (meta mf) m)")}))
				}
			}
			{
				var a_5001 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_5001
				if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Deref.X_invoke_Arity1(a_5001)) {
				} else {
					panic((&js.Error{("Assert failed: (= 0 (deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity2(a_5001, cljs_core.Inc)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (swap! a inc))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(false, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_5001, float64(0), float64(42)))) {
				} else {
					panic((&js.Error{("Assert failed: (= false (compare-and-set! a 0 42))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(true, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_5001, float64(1), float64(7)))) {
				} else {
					panic((&js.Error{("Assert failed: (= true (compare-and-set! a 1 7))")}))
				}
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(a_5001)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (meta a))")}))
				}
				if cljs_core.Nil_(cljs_core.Get_validator.X_invoke_Arity1(a_5001)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (get-validator a))")}))
				}
			}
			{
				var a_5002 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_5002
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity3(a_5002, cljs_core.X_PLUS_, float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (swap! a + 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Swap_BANG_.X_invoke_Arity4(a_5002, cljs_core.X_PLUS_, float64(1), float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (swap! a + 1 2))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(10), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_5002, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 10 (swap! a + 1 2 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(20), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_5002, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(4)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 20 (swap! a + 1 2 3 4))")}))
				}
			}
			{
				var a_5003 = cljs_core.Atom.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "validator", Fqn: "validator", X_hash: float64(-1966190681)}), cljs_core.Coll_QMARK_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})})).(*cljs_core.CljsCoreAtom)
				_ = a_5003
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Coll_QMARK_, cljs_core.Get_validator.X_invoke_Arity1(a_5003)) {
				} else {
					panic((&js.Error{("Assert failed: (= coll? (get-validator a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_5003)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:a 1} (meta a))")}))
				}
				cljs_core.Alter_meta_BANG_.X_invoke_ArityVariadic(a_5003, cljs_core.Assoc, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}))
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_5003)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:a 1, :b 2} (meta a))")}))
				}
			}
			if cljs_core.Nil_(cljs_core.Empty.X_invoke_Arity1(nil)) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (empty nil))")}))
			}
			{
				var e_lazy_seq_5004 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCoreLazySeq{nil, func(G__5005 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5005, 0, func() interface{} {
						return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons)
					})
				}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_lazy_seq_5004
				if cljs_core.Seq_QMARK_.Arity1IB(e_lazy_seq_5004) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-lazy-seq)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_lazy_seq_5004) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-lazy-seq)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_lazy_seq_5004)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-lazy-seq))")}))
				}
			}
			{
				var e_list_5006 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_list_5006
				if cljs_core.Seq_QMARK_.Arity1IB(e_list_5006) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-list)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_list_5006) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-list)")}))
				}
			}
			{
				var e_elist_5007 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_elist_5007
				if cljs_core.Seq_QMARK_.Arity1IB(e_elist_5007) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-elist)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_elist_5007) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-elist)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), cljs_core.Get.X_invoke_Arity2(cljs_core.Meta.X_invoke_Arity1(e_elist_5007), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= :c (get (meta e-elist) :b))")}))
				}
			}
			{
				var e_cons_5008 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_cons_5008
				if cljs_core.Seq_QMARK_.Arity1IB(e_cons_5008) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-cons)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_cons_5008) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-cons)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_cons_5008)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-cons))")}))
				}
			}
			{
				var e_vec_5009 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_vec_5009
				if cljs_core.Vector_QMARK_.Arity1IB(e_vec_5009) {
				} else {
					panic((&js.Error{("Assert failed: (vector? e-vec)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_vec_5009) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-vec)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_vec_5009)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-vec))")}))
				}
			}
			{
				var e_omap_5010 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_omap_5010
				if cljs_core.Map_QMARK_.Arity1IB(e_omap_5010) {
				} else {
					panic((&js.Error{("Assert failed: (map? e-omap)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_omap_5010) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-omap)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_omap_5010)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-omap))")}))
				}
			}
			{
				var e_hmap_5011 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_hmap_5011
				if cljs_core.Map_QMARK_.Arity1IB(e_hmap_5011) {
				} else {
					panic((&js.Error{("Assert failed: (map? e-hmap)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_hmap_5011) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-hmap)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_hmap_5011)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-hmap))")}))
				}
			}
			{
				var a_5012 = cljs_core.Atom.X_invoke_Arity1(nil).(*cljs_core.CljsCoreAtom)
				_ = a_5012
				if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (try 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5013 interface{}) {
					defer func() {
						if e4553 := recover(); e4553 != nil {
							if cljs_core.Value_(e4553).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4553
									_ = e
									return__5013 = float64(2)
								}
							} else {
								panic(e4553)

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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5014 interface{}) {
					defer func() {
						if e4554 := recover(); e4554 != nil {
							if cljs_core.Value_(e4554).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4554
									_ = e
									return__5014 = float64(2)
								}
							} else {
								panic(e4554)

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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5015 interface{}) {
					defer func() {
						if e4555 := recover(); e4555 != nil {
							if cljs_core.Value_(e4555).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4555
									_ = e
									return__5015 = float64(2)
								}
							} else {
								{
									var e = e4555
									_ = e
									return__5015 = float64(3)
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
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__5016 interface{}) {
					defer func() {
						if e4556 := recover(); e4556 != nil {
							if cljs_core.Value_(e4556).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4556
									_ = e
									return__5016 = float64(2)
								}
							} else {
								{
									var e = e4556
									_ = e
									return__5016 = float64(3)
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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5017 interface{}) {
					defer func() {
						if e4557 := recover(); e4557 != nil {
							if cljs_core.Value_(e4557).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4557
									_ = e
									return__5017 = float64(3)
								}
							} else {
								{
									var e = e4557
									_ = e
									return__5017 = e
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
						cljs_core.Reset_BANG_.X_invoke_Arity2(a_5012, float64(42))
					}()
					{
						return float64(1)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (try 1 (finally (reset! a 42))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Deref.X_invoke_Arity1(a_5012)) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (deref a))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), cljs_core.Seq_(cljs_core.Nthnext.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)))) {
			} else {
				panic((&js.Error{("Assert failed: (= [3] (nthnext [1 2 3] 2))")}))
			}
			{
				var v_5018 = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})
				_ = v_5018
				if cljs_core.X_EQ_.Arity2IIB(v_5018, func() *cljs_core.CljsCoreLazySeq {
					var iter__923__auto__ = func(iter__4558 *cljs_core.AFn, v_5018 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__4558, 1, func(s__4559 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__5019 *cljs_core.AFn, v_5018 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__5019, 0, func() interface{} {
									{
										var s__4559___1 interface{} = s__4559
										_ = s__4559___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4559___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__4559___2 = temp__4222__auto__
														_ = s__4559___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4559___2) {
															{
																var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4559___2)
																var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																var b__4561 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__921__auto__, size__922__auto__, b__4561
																if func() bool {
																	var i__4560 = float64(0)
																	_ = i__4560
																	for {
																		if i__4560 < size__922__auto__ {
																			{
																				var e = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4560)
																				_ = e
																				cljs_core.Chunk_append.X_invoke_Arity2(b__4561, e)
																				i__4560 = (i__4560 + float64(1))
																				continue
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4561), iter__4558.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4559___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4561), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__4559___2)
																_ = e
																return cljs_core.Cons.X_invoke_Arity2(e, iter__4558.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4559___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_5018), nil, nil})
						})
					}(&cljs_core.AFn{}, v_5018)
					_ = iter__923__auto__
					return iter__923__auto__.X_invoke_Arity1(v_5018).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= v (for [e v] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(9)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__923__auto__ = func(iter__4564 *cljs_core.AFn, v_5018 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__4564, 1, func(s__4565 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__5020 *cljs_core.AFn, v_5018 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__5020, 0, func() interface{} {
									{
										var s__4565___1 interface{} = s__4565
										_ = s__4565___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4565___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__4565___2 = temp__4222__auto__
														_ = s__4565___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4565___2) {
															{
																var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4565___2)
																var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																var b__4567 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__921__auto__, size__922__auto__, b__4567
																if func() bool {
																	var i__4566 = float64(0)
																	_ = i__4566
																	for {
																		if i__4566 < size__922__auto__ {
																			{
																				var e = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4566)
																				_ = e
																				{
																					var m = (e.(float64) * e.(float64))
																					_ = m
																					cljs_core.Chunk_append.X_invoke_Arity2(b__4567, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																					i__4566 = (i__4566 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4567), iter__4564.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4565___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4567), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__4565___2)
																_ = e
																{
																	var m = (e.(float64) * e.(float64))
																	_ = m
																	return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__4564.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4565___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_5018), nil, nil})
						})
					}(&cljs_core.AFn{}, v_5018)
					_ = iter__923__auto__
					return iter__923__auto__.X_invoke_Arity1(v_5018).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [[1 1] [2 4] [3 9]] (for [e v :let [m (* e e)]] [e m]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__923__auto__ = func(iter__4570 *cljs_core.AFn, v_5018 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__4570, 1, func(s__4571 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__5021 *cljs_core.AFn, v_5018 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__5021, 0, func() interface{} {
									{
										var s__4571___1 interface{} = s__4571
										_ = s__4571___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4571___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__4571___2 = temp__4222__auto__
														_ = s__4571___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4571___2) {
															{
																var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4571___2)
																var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																var b__4573 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__921__auto__, size__922__auto__, b__4573
																if cljs_core.Truth_(func() interface{} {
																	var i__4572 = float64(0)
																	_ = i__4572
																	for {
																		if i__4572 < size__922__auto__ {
																			{
																				var e = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4572)
																				_ = e
																				if e.(float64) < float64(3) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__4573, e)
																					i__4572 = (i__4572 + float64(1))
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
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4573), iter__4570.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4571___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4573), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__4571___2)
																_ = e
																if e.(float64) < float64(3) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__4570.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4571___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_5018), nil, nil})
						})
					}(&cljs_core.AFn{}, v_5018)
					_ = iter__923__auto__
					return iter__923__auto__.X_invoke_Arity1(v_5018).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2] (for [e v :while (< e 3)] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__923__auto__ = func(iter__4576 *cljs_core.AFn, v_5018 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__4576, 1, func(s__4577 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__5022 *cljs_core.AFn, v_5018 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__5022, 0, func() interface{} {
									{
										var s__4577___1 interface{} = s__4577
										_ = s__4577___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4577___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__4577___2 = temp__4222__auto__
														_ = s__4577___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4577___2) {
															{
																var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4577___2)
																var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																var b__4579 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__921__auto__, size__922__auto__, b__4579
																if func() bool {
																	var i__4578 = float64(0)
																	_ = i__4578
																	for {
																		if i__4578 < size__922__auto__ {
																			{
																				var e = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4578)
																				_ = e
																				if e.(float64) > float64(2) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__4579, e)
																					i__4578 = (i__4578 + float64(1))
																					continue
																				} else {
																					i__4578 = (i__4578 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4579), iter__4576.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4577___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4579), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__4577___2)
																_ = e
																if e.(float64) > float64(2) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__4576.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4577___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
																} else {
																	s__4577___1 = cljs_core.Rest.Arity1IQ(s__4577___2)
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
							}(&cljs_core.AFn{}, v_5018), nil, nil})
						})
					}(&cljs_core.AFn{}, v_5018)
					_ = iter__923__auto__
					return iter__923__auto__.X_invoke_Arity1(v_5018).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [3] (for [e v :when (> e 2)] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__923__auto__ = func(iter__4582 *cljs_core.AFn, v_5018 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__4582, 1, func(s__4583 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__5023 *cljs_core.AFn, v_5018 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__5023, 0, func() interface{} {
									{
										var s__4583___1 interface{} = s__4583
										_ = s__4583___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__4583___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__4583___2 = temp__4222__auto__
														_ = s__4583___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__4583___2) {
															{
																var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__4583___2)
																var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																var b__4585 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__921__auto__, size__922__auto__, b__4585
																if cljs_core.Truth_(func() interface{} {
																	var i__4584 = float64(0)
																	_ = i__4584
																	for {
																		if i__4584 < size__922__auto__ {
																			{
																				var e = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4584)
																				_ = e
																				if e.(float64) < float64(3) {
																					{
																						var m = (e.(float64) * e.(float64))
																						_ = m
																						cljs_core.Chunk_append.X_invoke_Arity2(b__4585, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																						i__4584 = (i__4584 + float64(1))
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
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4585), iter__4582.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__4583___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__4585), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__4583___2)
																_ = e
																if e.(float64) < float64(3) {
																	{
																		var m = (e.(float64) * e.(float64))
																		_ = m
																		return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__4582.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__4583___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_5018), nil, nil})
						})
					}(&cljs_core.AFn{}, v_5018)
					_ = iter__923__auto__
					return iter__923__auto__.X_invoke_Arity1(v_5018).(*cljs_core.CljsCoreLazySeq)
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
				var a10_5024 = cljs_core.Partial.X_invoke_Arity2(cljs_core.X_PLUS_, float64(10)).(cljs_core.CljsCoreIFn)
				var a20_5025 = cljs_core.Partial.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), float64(10)).(cljs_core.CljsCoreIFn)
				var a21_5026 = cljs_core.Partial.X_invoke_Arity4(cljs_core.X_PLUS_, float64(10), float64(10), float64(1)).(cljs_core.CljsCoreIFn)
				var a22_5027 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)})).(cljs_core.CljsCoreIFn)
				var a23_5028 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(1)})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = a10_5024, a20_5025, a21_5026, a22_5027, a23_5028
				if cljs_core.X_EQ_.Arity2IIB(float64(110), func() interface{} {
					var G__4588 = float64(100)
					_ = G__4588
					return a10_5024.X_invoke_Arity1(G__4588)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 110 (a10 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(120), func() interface{} {
					var G__4589 = float64(100)
					_ = G__4589
					return a20_5025.X_invoke_Arity1(G__4589)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 120 (a20 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(121), func() interface{} {
					var G__4590 = float64(100)
					_ = G__4590
					return a21_5026.X_invoke_Arity1(G__4590)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 121 (a21 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(122), func() interface{} {
					var G__4591 = float64(100)
					_ = G__4591
					return a22_5027.X_invoke_Arity1(G__4591)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 122 (a22 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(123), func() interface{} {
					var G__4592 = float64(100)
					_ = G__4592
					return a23_5028.X_invoke_Arity1(G__4592)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 123 (a23 100))")}))
				}
			}
			{
				var n2_5029 = cljs_core.Comp.X_invoke_Arity2(cljs_core.First, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n3_5030 = cljs_core.Comp.X_invoke_Arity3(cljs_core.First, cljs_core.Rest, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n4_5031 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n5_5032 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n6_5033 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = n2_5029, n3_5030, n4_5031, n5_5032, n6_5033
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() interface{} {
					var G__4593 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__4593
					return n2_5029.X_invoke_Arity1(G__4593)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (n2 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() interface{} {
					var G__4594 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__4594
					return n3_5030.X_invoke_Arity1(G__4594)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (n3 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__4595 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__4595
					return n4_5031.X_invoke_Arity1(G__4595)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (n4 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(5), func() interface{} {
					var G__4596 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__4596
					return n5_5032.X_invoke_Arity1(G__4596)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 5 (n5 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), func() interface{} {
					var G__4597 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__4597
					return n6_5033.X_invoke_Arity1(G__4597)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (n6 [1 2 3 4 5 6 7]))")}))
				}
			}
			{
				var sf_5034 = cljs_core.Some_fn.X_invoke_Arity3(cljs_core.Number_QMARK_, cljs_core.Keyword_QMARK_, cljs_core.Symbol_QMARK_).(cljs_core.CljsCoreIFn)
				_ = sf_5034
				if cljs_core.Truth_(func() interface{} {
					var G__4598 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					var G__4599 = float64(1)
					_, _ = G__4598, G__4599
					return sf_5034.X_invoke_Arity2(G__4598, G__4599)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf :foo 1)")}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__4600 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					_ = G__4600
					return sf_5034.X_invoke_Arity1(G__4600)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf :foo)")}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__4601 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})
					var G__4602 = float64(1)
					_, _ = G__4601, G__4602
					return sf_5034.X_invoke_Arity2(G__4601, G__4602)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf (quote bar) 1)")}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__4603 = cljs_core.CljsCorePersistentVector_EMPTY
					var G__4604 = cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)
					_, _ = G__4603, G__4604
					return sf_5034.X_invoke_Arity2(G__4603, G__4604)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (not (sf [] ()))")}))
				}
			}
			{
				var ep_5035 = cljs_core.Every_pred.X_invoke_Arity2(cljs_core.Number_QMARK_, cljs_core.Zero_QMARK_).(cljs_core.CljsCoreIFn)
				_ = ep_5035
				if cljs_core.Truth_(func() interface{} {
					var G__4605 = float64(0)
					var G__4606 = float64(0)
					var G__4607 = float64(0)
					_, _, _ = G__4605, G__4606, G__4607
					return ep_5035.X_invoke_Arity3(G__4605, G__4606, G__4607)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (ep 0 0 0)")}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__4608 = float64(1)
					var G__4609 = float64(2)
					var G__4610 = float64(3)
					var G__4611 = float64(0)
					_, _, _, _ = G__4608, G__4609, G__4610, G__4611
					return ep_5035.X_invoke_Arity4(G__4608, G__4609, G__4610, G__4611)
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__4612 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4612, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4612, float64(1), nil)
				_, _, _ = vec__4612, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] [1 2]] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), func() cljs_core.CljsCoreISet {
				var vec__4613 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4613, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4613, float64(1), nil)
				_, _, _ = vec__4613, a, b
				return cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{a, b}, true).(*cljs_core.CljsCorePersistentHashSet)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 2} (let [[a b] [1 2]] #{a b}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__4614 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__4614___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4614) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4614)
					} else {
						return map__4614
					}
				}()
				var a = cljs_core.Get.X_invoke_Arity2(map__4614___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				var b = cljs_core.Get.X_invoke_Arity2(map__4614___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				_, _, _, _ = map__4614, map__4614___1, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (let [{a :a, b :b} {:a 1, :b 2}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__4615 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__4615___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4615) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4615)
					} else {
						return map__4615
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__4615___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__4615___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__4615, map__4615___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (let [{:keys [a b]} {:a 1, :b 2}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__4616 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4616, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4616, float64(1), nil)
				var v = vec__4616
				_, _, _, _ = vec__4616, a, b, v
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b, v}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 [1 2]] (let [[a b :as v] [1 2]] [a b v]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(42)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__4617 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__4617___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4617) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4617)
					} else {
						return map__4617
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity3(map__4617___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(42))
				var a = cljs_core.Get.X_invoke_Arity2(map__4617___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__4617, map__4617___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 42] (let [{:keys [a b], :or {b 42}} {:a 1}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), nil}, nil}), func() cljs_core.CljsCoreIVector {
				var map__4618 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__4618___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4618) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4618)
					} else {
						return map__4618
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__4618___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__4618___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__4618, map__4618___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 nil] (let [{:keys [a b], :or {c 42}} {:a 1}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__4619 = cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4619, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4619, float64(1), nil)
				_, _, _ = vec__4619, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] (quote (1 2))] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIMap {
				var vec__4620 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4620, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4620, float64(1), nil)
				_, _, _ = vec__4620, a, b
				return cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{a, b}, true, false).(*cljs_core.CljsCorePersistentArrayMap)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= {1 2} (let [[a b] [1 2]] {a b}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__4621 = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))
				var a = cljs_core.Nth.X_invoke_Arity3(vec__4621, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__4621, float64(1), nil)
				_, _, _ = vec__4621, a, b
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
				var a_5036 = cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{})
				_ = a_5036
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20), float64(30)}, nil}), cljs_core.Seq.Arity1IQ(func() []interface{} {
					var a__1042__auto__ = a_5036
					var ret = cljs_core.Aclone.X_invoke_Arity1(a__1042__auto__).([]interface{})
					_, _ = a__1042__auto__, ret
					{
						var i = float64(0)
						_ = i
						for {
							if i < float64(len(a__1042__auto__)) {
								ret[int(i)] = (float64(10) * (a_5036[int(i)]).(float64))
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
					var a__1048__auto__ = a_5036
					_ = a__1048__auto__
					{
						var i = float64(0)
						var ret = float64(0)
						_, _ = i, ret
						for {
							if i < float64(len(a__1048__auto__)) {
								i, ret = (i + float64(1)), (ret + (a_5036[int(i)]).(float64))
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_5036), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (seq a) (seq (to-array [1 2 3])))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), func() float64 { a_5036[int(float64(0))] = float64(42); return a_5036[int(float64(0))].(float64) }()) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (aset a 0 42))")}))
				}
				if cljs_core.Not_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_5036), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
				} else {
					panic((&js.Error{("Assert failed: (not= (seq a) (seq (to-array [1 2 3])))")}))
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
				var coll_5037 = (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})
				var shuffles_5038 = cljs_core.Filter.X_invoke_Arity2(func(G__5039 *cljs_core.AFn, coll_5037 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__5039, 1, func(p1__4080_SHARP_ interface{}) interface{} {
						return cljs_core.Not_EQ_.Arity2IIB(coll_5037, p1__4080_SHARP_)
					})
				}(&cljs_core.AFn{}, coll_5037), cljs_core.Take.X_invoke_Arity2(float64(100), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Shuffle, coll_5037).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
				_, _ = coll_5037, shuffles_5038
				if !(cljs_core.Empty_QMARK_.Arity1IB(shuffles_5038)) {
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
				var s_5040 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = s_5040
				{
					var n__1054__auto___5041 = float64(5)
					_ = n__1054__auto___5041
					{
						var n_5042 = float64(0)
						_ = n_5042
						for {
							if n_5042 < n__1054__auto___5041 {
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_5040, cljs_core.Conj, n_5042)
								n_5042 = (n_5042 + float64(1))
								continue
							} else {
							}
							break
						}
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Deref.X_invoke_Arity1(s_5040)) {
				} else {
					panic((&js.Error{("Assert failed: (= [0 1 2 3 4] (clojure.core/deref s))")}))
				}
			}
			{
				var v_5043 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})
				var s_5044 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)).(*cljs_core.CljsCoreAtom)
				_, _ = v_5043, s_5044
				{
					var seq__4622_5045 interface{} = cljs_core.Seq.Arity1IQ(v_5043)
					var chunk__4623_5046 interface{} = nil
					var count__4624_5047 = float64(0)
					var i__4625_5048 = float64(0)
					_, _, _, _ = seq__4622_5045, chunk__4623_5046, count__4624_5047, i__4625_5048
					for {
						if i__4625_5048 < count__4624_5047 {
							{
								var n_5049 = chunk__4623_5046.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4625_5048)
								_ = n_5049
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_5044, cljs_core.Conj, n_5049)
								seq__4622_5045, chunk__4623_5046, count__4624_5047, i__4625_5048 = seq__4622_5045, chunk__4623_5046, count__4624_5047, (i__4625_5048 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___5050 = cljs_core.Seq.Arity1IQ(seq__4622_5045)
								_ = temp__4222__auto___5050
								if cljs_core.Truth_(temp__4222__auto___5050) {
									{
										var seq__4622_5051___1 = temp__4222__auto___5050
										_ = seq__4622_5051___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4622_5051___1) {
											{
												var c__954__auto___5052 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4622_5051___1)
												_ = c__954__auto___5052
												seq__4622_5045, chunk__4623_5046, count__4624_5047, i__4625_5048 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4622_5051___1), c__954__auto___5052, cljs_core.Count.X_invoke_Arity1(c__954__auto___5052).(float64), float64(0)
												continue
											}
										} else {
											{
												var n_5053 = cljs_core.First.X_invoke_Arity1(seq__4622_5051___1)
												_ = n_5053
												cljs_core.Swap_BANG_.X_invoke_Arity3(s_5044, cljs_core.Conj, n_5053)
												seq__4622_5045, chunk__4623_5046, count__4624_5047, i__4625_5048 = cljs_core.Next.Arity1IQ(seq__4622_5051___1), nil, float64(0), float64(0)
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Deref.X_invoke_Arity1(s_5044), cljs_core.Reverse.X_invoke_Arity1(v_5043)) {
				} else {
					panic((&js.Error{("Assert failed: (= (clojure.core/deref s) (reverse v))")}))
				}
			}
			{
				var a_5054 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				var d_5055 = (&cljs_core.CljsCoreDelay{func(G__5056 *cljs_core.AFn, a_5054 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__5056, 0, func() interface{} {
						return cljs_core.Swap_BANG_.X_invoke_Arity2(a_5054, cljs_core.Inc)
					})
				}(&cljs_core.AFn{}, a_5054), nil})
				_, _ = a_5054, d_5055
				if cljs_core.Realized_QMARK_.Arity1IB(d_5055) == false {
				} else {
					panic((&js.Error{("Assert failed: (false? (realized? d))")}))
				}
				if cljs_core.Deref.X_invoke_Arity1(a_5054).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (clojure.core/deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_5055)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_5054)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref a))")}))
				}
				if cljs_core.Realized_QMARK_.Arity1IB(d_5055) == true {
				} else {
					panic((&js.Error{("Assert failed: (true? (realized? d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_5055)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_5054)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Force.X_invoke_Arity1(d_5055), cljs_core.Deref.X_invoke_Arity1(d_5055)) {
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
				var f_5057 = cljs_core.Memoize.X_invoke_Arity1(func(G__5058 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5058, 0, func() interface{} {
						return cljs_core.Rand.Arity0F()
					})
				}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn)
				_ = f_5057
				{
					f_5057.X_invoke_Arity0()
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return f_5057.X_invoke_Arity0()
				}(), func() interface{} {
					return f_5057.X_invoke_Arity0()
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
				var d_5059 = cljs_core.Group_by.X_invoke_Arity2(cljs_core.Second, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(6), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), float64(4), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), float64(2)}, nil}))
				_ = d_5059
				if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_5059, float64(1))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (count (get d 1)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_5059, float64(2))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (count (get d 2)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_5059, float64(4))).(float64)) {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__5060 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5060, 1, func(p1__4081_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__4081_SHARP_) {
						return p1__4081_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 3 5 7 9] (keep (fn* [p1__4081#] (if (odd? p1__4081#) p1__4081#)) [1 2 3 4 5 6 7 8 9 10]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(6), float64(8), float64(10)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__5061 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5061, 1, func(p1__4082_SHARP_ interface{}) interface{} {
					if cljs_core.Even_QMARK_.Arity1IB(p1__4082_SHARP_) {
						return p1__4082_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 4 6 8 10] (keep (fn* [p1__4082#] (if (even? p1__4082#) p1__4082#)) [1 2 3 4 5 6 7 8 9 10]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__5062 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5062, 2, func(p1__4083_SHARP_ interface{}, p2__4084_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__4083_SHARP_) {
						return p2__4084_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(11), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 3 5 7 9] (keep-indexed (fn* [p1__4083# p2__4084#] (if (odd? p1__4083#) p2__4084#)) [0 1 2 3 4 5 6 7 8 9 10]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(5)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__5063 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5063, 2, func(p1__4086_SHARP_ interface{}, p2__4085_SHARP_ interface{}) interface{} {
					if p2__4085_SHARP_.(float64) > float64(0) {
						return p1__4086_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-9), float64(0), float64(29), float64(-7), float64(45), float64(3), float64(-8)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 4 5] (keep-indexed (fn* [p1__4086# p2__4085#] (if (pos? p2__4085#) p1__4086#)) [-9 0 29 -7 45 3 -8]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil}), cljs_core.Map_indexed.X_invoke_Arity2(func(G__5064 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5064, 2, func(p1__4087_SHARP_ interface{}, p2__4088_SHARP_ interface{}) interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__4087_SHARP_, p2__4088_SHARP_}, nil})
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [[0 :a] [1 :b] [2 :c]] (map-indexed (fn* [p1__4087# p2__4088#] (vector p1__4087# p2__4088#)) [:a :b :c]))")}))
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
						return func(G__5065 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__5065, 0, func() interface{} {
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
				var method_table__1064__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1065__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1066__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1067__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1068__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__
				return (&cljs_core.CljsCoreMultiFn{"bar", func(G__5066 *cljs_core.AFn, method_table__1064__auto__ *cljs_core.CljsCoreAtom, prefer_table__1065__auto__ *cljs_core.CljsCoreAtom, method_cache__1066__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1067__auto__ *cljs_core.CljsCoreAtom, hierarchy__1068__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5066, 2, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{x, y}, nil})
					})
				}(&cljs_core.AFn{}, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}), func(G__5067 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5067, 2, func(x interface{}, y interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)})
				})
			}(&cljs_core.AFn{}))
			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})}, nil}), func(G__5068 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5068, 2, func(x interface{}, y interface{}) interface{} {
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
				var G__4627 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				var G__4628 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				_, _ = G__4627, G__4628
				return Bar.X_invoke_Arity2(G__4627, G__4628)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :rect-shape (bar :cljs.core-test/rect :cljs.core-test/rect))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)}), cljs_core.Apply.X_invoke_Arity2(Bar.X_get_method_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= :rect-shape (apply (-get-method bar [:cljs.core-test/rect :cljs.core-test/shape]) [:cljs.core-test/rect :cljs.core-test/shape]))")}))
			}
			Nested_dispatch = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1064__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1065__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1066__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1067__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1068__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__
				return (&cljs_core.CljsCoreMultiFn{"nested-dispatch", func(G__5069 *cljs_core.AFn, method_table__1064__auto__ *cljs_core.CljsCoreAtom, prefer_table__1065__auto__ *cljs_core.CljsCoreAtom, method_cache__1066__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1067__auto__ *cljs_core.CljsCoreAtom, hierarchy__1068__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5069, 1, func(m interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(m))
					})
				}(&cljs_core.AFn{}, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), func(G__5070 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5070, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__5071 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5071, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__4629 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil})
				_ = G__4629
				return Nested_dispatch.X_invoke_Arity1(G__4629)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :nested-a (nested-dispatch {:a {:b :c}}))")}))
			}
			Nested_dispatch2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1064__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1065__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1066__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1067__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1068__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__
				return (&cljs_core.CljsCoreMultiFn{"nested-dispatch2", cljs_core.Ffirst, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__5072 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5072, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__5073 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5073, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__4630 = (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})}, nil})
				_ = G__4630
				return Nested_dispatch2.X_invoke_Arity1(G__4630)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :nested-a (nested-dispatch2 [[:a :b]]))")}))
			}
			Foo1 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1064__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1065__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1066__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1067__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1068__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__
				return (&cljs_core.CljsCoreMultiFn{"foo1", func(G__5074 *cljs_core.AFn, method_table__1064__auto__ *cljs_core.CljsCoreAtom, prefer_table__1065__auto__ *cljs_core.CljsCoreAtom, method_cache__1066__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1067__auto__ *cljs_core.CljsCoreAtom, hierarchy__1068__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5074, 0, func(args__ ...interface{}) interface{} {
						var args = cljs_core.Seq.Arity1IQ(args__[0])
						_ = args
						return cljs_core.First.X_invoke_Arity1(args)
					})
				}(&cljs_core.AFn{}, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__5075 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5075, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)})
				})
			}(&cljs_core.AFn{}))
			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__5076 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5076, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)}), func() interface{} {
				var G__4631 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
				_ = G__4631
				return Foo1.X_invoke_Arity1(G__4631)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :a-return (foo1 :a))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)}), func() interface{} {
				var G__4632 = float64(1)
				_ = G__4632
				return Foo1.X_invoke_Arity1(G__4632)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :default-return (foo1 1))")}))
			}
			Area = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1064__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1065__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1066__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1067__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1068__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__
				return (&cljs_core.CljsCoreMultiFn{"area", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Shape", Fqn: "Shape", X_hash: float64(-248980586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
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

			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Rect", Fqn: "Rect", X_hash: float64(-420704620)}), func(G__5077 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5077, 1, func(r interface{}) interface{} {
					return ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wd", Fqn: "wd", X_hash: float64(-183204751)}).X_invoke_Arity1(r).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ht", Fqn: "ht", X_hash: float64(214115472)}).X_invoke_Arity1(r).(float64))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Circle", Fqn: "Circle", X_hash: float64(-158896930)}), func(G__5078 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5078, 1, func(c interface{}) interface{} {
					return (Math.PI * ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64)))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__5079 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5079, 1, func(x interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)})
				})
			}(&cljs_core.AFn{}))
			R = Rect.X_invoke_Arity2(float64(4), float64(13)).(cljs_core.CljsCoreIMap)

			C = Circle.X_invoke_Arity1(float64(12)).(cljs_core.CljsCoreIMap)

			if cljs_core.X_EQ_.Arity2IIB(float64(52), func() interface{} {
				var G__4633 = R
				_ = G__4633
				return Area.X_invoke_Arity1(G__4633)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 52 (area r))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)}), func() interface{} {
				var G__4634 = cljs_core.CljsCorePersistentArrayMap_EMPTY
				_ = G__4634
				return Area.X_invoke_Arity1(G__4634)
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
				var method_table__1064__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1065__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1066__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1067__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1068__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__
				return (&cljs_core.CljsCoreMultiFn{"apply-multi-test", func(G__5080 *cljs_core.AFn, method_table__1064__auto__ *cljs_core.CljsCoreAtom, prefer_table__1065__auto__ *cljs_core.CljsCoreAtom, method_cache__1066__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1067__auto__ *cljs_core.CljsCoreAtom, hierarchy__1068__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5080, 3, func(___ interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
						return float64(0)
					})
				}(&cljs_core.AFn{}, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			Apply_multi_test.X_add_method_Arity3(float64(0), func(G__5081 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5081, 2, func(x interface{}) interface{} {
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
				var method_table__1064__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1065__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1066__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1067__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1068__auto__ = cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), My_map_hierarchy}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__
				return (&cljs_core.CljsCoreMultiFn{"my-map?", cljs_core.Type_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)}), func(G__5082 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5082, 1, func(___ interface{}) interface{} {
					return true
				})
			}(&cljs_core.AFn{}))
			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__5083 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5083, 1, func(___ interface{}) interface{} {
					return false
				})
			}(&cljs_core.AFn{}))
			{
				var seq__4635_5084 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__4636_5085 interface{} = nil
				var count__4637_5086 = float64(0)
				var i__4638_5087 = float64(0)
				_, _, _, _ = seq__4635_5084, chunk__4636_5085, count__4637_5086, i__4638_5087
				for {
					if i__4638_5087 < count__4637_5086 {
						{
							var m_5088 = chunk__4636_5085.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4638_5087)
							_ = m_5088
							if cljs_core.Truth_(func() interface{} {
								var G__4639 = m_5088
								_ = G__4639
								return My_map_QMARK_.X_invoke_Arity1(G__4639)
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (my-map? m)")}))
							}
							seq__4635_5084, chunk__4636_5085, count__4637_5086, i__4638_5087 = seq__4635_5084, chunk__4636_5085, count__4637_5086, (i__4638_5087 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5089 = cljs_core.Seq.Arity1IQ(seq__4635_5084)
							_ = temp__4222__auto___5089
							if cljs_core.Truth_(temp__4222__auto___5089) {
								{
									var seq__4635_5090___1 = temp__4222__auto___5089
									_ = seq__4635_5090___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4635_5090___1) {
										{
											var c__954__auto___5091 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4635_5090___1)
											_ = c__954__auto___5091
											seq__4635_5084, chunk__4636_5085, count__4637_5086, i__4638_5087 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4635_5090___1), c__954__auto___5091, cljs_core.Count.X_invoke_Arity1(c__954__auto___5091).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_5092 = cljs_core.First.X_invoke_Arity1(seq__4635_5090___1)
											_ = m_5092
											if cljs_core.Truth_(func() interface{} {
												var G__4640 = m_5092
												_ = G__4640
												return My_map_QMARK_.X_invoke_Arity1(G__4640)
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (my-map? m)")}))
											}
											seq__4635_5084, chunk__4636_5085, count__4637_5086, i__4638_5087 = cljs_core.Next.Arity1IQ(seq__4635_5090___1), nil, float64(0), float64(0)
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
				var seq__4641_5093 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentVector_EMPTY, float64(1), "asdf", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}))
				var chunk__4642_5094 interface{} = nil
				var count__4643_5095 = float64(0)
				var i__4644_5096 = float64(0)
				_, _, _, _ = seq__4641_5093, chunk__4642_5094, count__4643_5095, i__4644_5096
				for {
					if i__4644_5096 < count__4643_5095 {
						{
							var not_m_5097 = chunk__4642_5094.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4644_5096)
							_ = not_m_5097
							if cljs_core.Not.Arity1IB(func() interface{} {
								var G__4645 = not_m_5097
								_ = G__4645
								return My_map_QMARK_.X_invoke_Arity1(G__4645)
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (not (my-map? not-m))")}))
							}
							seq__4641_5093, chunk__4642_5094, count__4643_5095, i__4644_5096 = seq__4641_5093, chunk__4642_5094, count__4643_5095, (i__4644_5096 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5098 = cljs_core.Seq.Arity1IQ(seq__4641_5093)
							_ = temp__4222__auto___5098
							if cljs_core.Truth_(temp__4222__auto___5098) {
								{
									var seq__4641_5099___1 = temp__4222__auto___5098
									_ = seq__4641_5099___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4641_5099___1) {
										{
											var c__954__auto___5100 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4641_5099___1)
											_ = c__954__auto___5100
											seq__4641_5093, chunk__4642_5094, count__4643_5095, i__4644_5096 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4641_5099___1), c__954__auto___5100, cljs_core.Count.X_invoke_Arity1(c__954__auto___5100).(float64), float64(0)
											continue
										}
									} else {
										{
											var not_m_5101 = cljs_core.First.X_invoke_Arity1(seq__4641_5099___1)
											_ = not_m_5101
											if cljs_core.Not.Arity1IB(func() interface{} {
												var G__4646 = not_m_5101
												_ = G__4646
												return My_map_QMARK_.X_invoke_Arity1(G__4646)
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (not (my-map? not-m))")}))
											}
											seq__4641_5093, chunk__4642_5094, count__4643_5095, i__4644_5096 = cljs_core.Next.Arity1IQ(seq__4641_5099___1), nil, float64(0), float64(0)
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
				var pv_5102 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange))
				_ = pv_5102
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity2(pv_5102, float64(96)), float64(96)) {
				} else {
					panic((&js.Error{("Assert failed: (= (nth pv 96) 96)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity3(pv_5102, float64(97), nil), nil) {
				} else {
					panic((&js.Error{("Assert failed: (= (nth pv 97 nil) nil)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4647 = float64(96)
					_ = G__4647
					return pv_5102.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4647)
				}(), float64(96)) {
				} else {
					panic((&js.Error{("Assert failed: (= (pv 96) 96)")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(cljs_core.CljsCorePersistentVector_EMPTY)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq []))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reverse.X_invoke_Arity1(pv_5102), cljs_core.Rseq.Arity1IQ(pv_5102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (reverse pv) (rseq pv))")}))
				}
			}
			{
				var pv_5103 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(33)).(*cljs_core.CljsCoreRange))
				_ = pv_5103
				if cljs_core.X_EQ_.Arity2IIB(pv_5103, cljs_core.Conj.X_invoke_Arity2(cljs_core.Conj.X_invoke_Arity2(cljs_core.Pop.X_invoke_Arity1(cljs_core.Pop.X_invoke_Arity1(pv_5103)), float64(31)), float64(32))) {
				} else {
					panic((&js.Error{("Assert failed: (= pv (-> pv pop pop (conj 31) (conj 32)))")}))
				}
			}
			{
				var stack1_5104 = cljs_core.Pop.X_invoke_Arity1(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange)))
				var stack2_5105 = cljs_core.Pop.X_invoke_Arity1(stack1_5104)
				_, _ = stack1_5104, stack2_5105
				if cljs_core.X_EQ_.Arity2IIB(float64(95), cljs_core.Peek.X_invoke_Arity1(stack1_5104)) {
				} else {
					panic((&js.Error{("Assert failed: (= 95 (peek stack1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(94), cljs_core.Peek.X_invoke_Arity1(stack2_5105)) {
				} else {
					panic((&js.Error{("Assert failed: (= 94 (peek stack2))")}))
				}
			}
			{
				var sentinel_5106 = cljs_core.Rand.Arity0F()
				_ = sentinel_5106
				if reflect.DeepEqual(sentinel_5106, func() (return__5107 interface{}) {
					defer func() {
						if e4648 := recover(); e4648 != nil {
							if cljs_core.Value_(e4648).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var ___ = e4648
									_ = ___
									return__5107 = sentinel_5106
								}
							} else {
								panic(e4648)

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
				var v1_5108 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				var v2_5109 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(5)).(*cljs_core.CljsCoreRange))
				var s_5110 = cljs_core.Subvec.X_invoke_Arity3(v1_5108, float64(2), float64(8)).(*cljs_core.CljsCoreSubvec)
				_, _, _ = v1_5108, v2_5109, s_5110
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(s_5110, cljs_core.Subvec.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2(v1_5108, float64(2)).(*cljs_core.CljsCoreSubvec), float64(0), float64(6)).(*cljs_core.CljsCoreSubvec), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Take.X_invoke_Arity2(float64(6), cljs_core.Drop.X_invoke_Arity2(float64(2), v1_5108).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= s (-> v1 (subvec 2) (subvec 0 6)) (->> v1 (drop 2) (take 6)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Count.X_invoke_Arity1(s_5110).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (count s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Pop.X_invoke_Arity1(s_5110)) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 4 5 6] (pop s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(7), cljs_core.Peek.X_invoke_Arity1(s_5110)) {
				} else {
					panic((&js.Error{("Assert failed: (= 7 (peek s))")}))
				}
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(1)}, nil}), cljs_core.Assoc.X_invoke_Arity3(s_5110, float64(6), float64(1)), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Conj.X_invoke_Arity2(s_5110, float64(1))}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 4 5 6 7 1] (assoc s 6 1) (conj s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(27), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, s_5110)) {
				} else {
					panic((&js.Error{("Assert failed: (= 27 (reduce + s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(s_5110, cljs_core.Vec.X_invoke_Arity1(s_5110)) {
				} else {
					panic((&js.Error{("Assert failed: (= s (vec s))")}))
				}
				{
					var m_5111 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), float64(1)}, nil})
					_ = m_5111
					if cljs_core.X_EQ_.Arity2IIB(m_5111, cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(s_5110, m_5111))) {
					} else {
						panic((&js.Error{("Assert failed: (= m (meta (with-meta s m)))")}))
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5112 interface{}) {
					defer func() {
						if e4649 := recover(); e4649 != nil {
							if cljs_core.Value_(e4649).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4649
									_ = e
									return__5112 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e4649)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_5109, float64(0), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 0 6) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5113 interface{}) {
					defer func() {
						if e4650 := recover(); e4650 != nil {
							if cljs_core.Value_(e4650).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4650
									_ = e
									return__5113 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e4650)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_5109, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5114 interface{}) {
					defer func() {
						if e4651 := recover(); e4651 != nil {
							if cljs_core.Value_(e4651).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4651
									_ = e
									return__5114 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e4651)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_5109, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5115 interface{}) {
					defer func() {
						if e4652 := recover(); e4652 != nil {
							if cljs_core.Value_(e4652).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4652
									_ = e
									return__5115 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e4652)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_5109, float64(3), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 3 6) (catch js/Error e :fail)))")}))
				}
				if reflect.DeepEqual(v1_5108, cljs_core.Subvec.X_invoke_Arity3(s_5110, float64(1), float64(4)).(*cljs_core.CljsCoreSubvec).V) {
				} else {
					panic((&js.Error{("Assert failed: (identical? v1 (.-v (subvec s 1 4)))")}))
				}
				{
					var sentinel_5116 = cljs_core.Rand.Arity0F()
					var s_5117___1 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2))
					_, _ = sentinel_5116, s_5117___1
					if reflect.DeepEqual(sentinel_5116, func() (return__5118 interface{}) {
						defer func() {
							if e4653 := recover(); e4653 != nil {
								if cljs_core.Value_(e4653).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var ___ = e4653
										_ = ___
										return__5118 = sentinel_5116
									}
								} else {
									panic(e4653)

								}
							}
						}()
						{
							{
								var G__4654 = float64(-1)
								_ = G__4654
								return s_5117___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4654)
							}
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (identical? sentinel (try (s -1) (catch js/Error _ sentinel)))")}))
					}
					if reflect.DeepEqual(sentinel_5116, func() (return__5119 interface{}) {
						defer func() {
							if e4655 := recover(); e4655 != nil {
								if cljs_core.Value_(e4655).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var ___ = e4655
										_ = ___
										return__5119 = sentinel_5116
									}
								} else {
									panic(e4655)

								}
							}
						}()
						{
							{
								var G__4656 = float64(1)
								_ = G__4656
								return s_5117___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4656)
							}
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (identical? sentinel (try (s 1) (catch js/Error _ sentinel)))")}))
					}
				}
				{
					var sv1_5120 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2)).(*cljs_core.CljsCoreSubvec)
					var sv2_5121 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(1)).(*cljs_core.CljsCoreSubvec)
					_, _ = sv1_5120, sv2_5121
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(sv1_5120), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(*cljs_core.CljsCoreList)) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq sv1) (quote (1)))")}))
					}
					if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(sv2_5121)) {
					} else {
						panic((&js.Error{("Assert failed: (nil? (rseq sv2))")}))
					}
				}
			}
			{
				var v1_5122 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(15), float64(48)).(*cljs_core.CljsCoreRange))
				var v2_5123 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(40), float64(57)).(*cljs_core.CljsCoreRange))
				var v1_5124___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Pop_BANG_.X_invoke_Arity1(cljs_core.Transient.X_invoke_Arity1(v1_5122)), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v2_5125___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(v2_5123), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v_5126 = cljs_core.Into.X_invoke_Arity2(v1_5124___1, v2_5125___1)
				_, _, _, _, _ = v1_5122, v2_5123, v1_5124___1, v2_5125___1, v_5126
				if cljs_core.X_EQ_.Arity2IIB(v_5126, cljs_core.Vec.X_invoke_Arity1(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(16), float64(47)).(*cljs_core.CljsCoreRange), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(41), float64(57)).(*cljs_core.CljsCoreRange), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})})).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (= v (vec (concat [:quux] (range 16 47) [:foo] [:quux] (range 41 57) [:bar])))")}))
				}
			}
			{
				var v_5127 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)
				var xs_5128 interface{} = cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)
				_, _ = v_5127, xs_5128
				for {
					{
						var temp__4220__auto___5129 = cljs_core.First.X_invoke_Arity1(xs_5128)
						_ = temp__4220__auto___5129
						if cljs_core.Truth_(temp__4220__auto___5129) {
							{
								var x_5130 = temp__4220__auto___5129
								_ = x_5130
								v_5127, xs_5128 = func() interface{} {
									var pred__4657 = func(G__5131 *cljs_core.AFn, v_5127 interface{}, xs_5128 interface{}, x_5130 interface{}, temp__4220__auto___5129 interface{}) *cljs_core.AFn {
										return cljs_core.Fn(G__5131, 2, func(p1__4089_SHARP_ interface{}, p2__4090_SHARP_ interface{}) interface{} {
											{
												var G__4660 = cljs_core.Mod.X_invoke_Arity2(p2__4090_SHARP_, float64(3)).(float64)
												_ = G__4660
												return p1__4089_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4660)
											}
										})
									}(&cljs_core.AFn{}, v_5127, xs_5128, x_5130, temp__4220__auto___5129)
									var expr__4658 = x_5130
									_, _ = pred__4657, expr__4658
									if cljs_core.Truth_(pred__4657.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(0), nil, float64(2), nil}, nil}, nil}), expr__4658)) {
										return cljs_core.Conj_BANG_.X_invoke_Arity2(v_5127, x_5130)
									} else {
										if cljs_core.Truth_(pred__4657.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), expr__4658)) {
											return cljs_core.Assoc_BANG_.X_invoke_Arity3(v_5127, cljs_core.Count.X_invoke_Arity1(v_5127).(float64), x_5130)
										} else {
											panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(expr__4658).(string))}))
										}
									}
								}(), cljs_core.Next.Arity1IQ(xs_5128)
								continue
							}
						} else {
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)), cljs_core.Persistent_BANG_.X_invoke_Arity1(v_5127)) {
							} else {
								panic((&js.Error{("Assert failed: (= (vec (range 100)) (persistent! v))")}))
							}
						}
					}
					break
				}
			}
			{
				var m1_5132 interface{} = cljs_core.CljsCorePersistentHashMap_EMPTY
				var m2_5133 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_EMPTY)
				var i_5134 = float64(0)
				_, _, _ = m1_5132, m2_5133, i_5134
				for {
					if i_5134 < float64(100) {
						m1_5132, m2_5133, i_5134 = cljs_core.Assoc.X_invoke_Arity3(m1_5132, i_5134, i_5134), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_5133, i_5134, i_5134), (i_5134 + float64(1))
						continue
					} else {
						{
							var m2_5135___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_5133)
							_ = m2_5135___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_5132).(float64), float64(100)) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m1) 100)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_5135___1).(float64), float64(100)) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m2) 100)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_5132, m2_5135___1) {
							} else {
								panic((&js.Error{("Assert failed: (= m1 m2)")}))
							}
							{
								var i_5136___1 = float64(0)
								_ = i_5136___1
								for {
									if i_5136___1 < float64(100) {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__4661 = i_5136___1
											_ = G__4661
											return m1_5132.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4661)
										}(), i_5136___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__4662 = i_5136___1
											_ = G__4662
											return m2_5135___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4662)
										}(), i_5136___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m2 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_5132, i_5136___1), i_5136___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_5135___1, i_5136___1), i_5136___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m2 i) i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_5132, i_5136___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m1 i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_5135___1, i_5136___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m2 i)")}))
										}
										i_5136___1 = (i_5136___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_5132))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range 100) (range 100)) (sort-by first (seq m1)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_5135___1))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range 100) (range 100)) (sort-by first (seq m2)))")}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_5132, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{("Assert failed: (not (contains? (dissoc m1 3) 3))")}))
							}
						}
					}
					break
				}
			}
			{
				var m_5137 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_5137
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5137).(float64), float64(7)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 7)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5137, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
				}
			}
			{
				var m_5138 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_5138
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5138).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5138, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var m_5139 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_5139
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5139).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5139, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var tm_5140 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_5140
				{
					var tm_5141___1 interface{} = tm_5140
					var ks_5142 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_5141___1, ks_5142
					for {
						{
							var temp__4220__auto___5143 = cljs_core.First.X_invoke_Arity1(ks_5142)
							_ = temp__4220__auto___5143
							if cljs_core.Truth_(temp__4220__auto___5143) {
								{
									var k_5144 = temp__4220__auto___5143
									_ = k_5144
									tm_5141___1, ks_5142 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_5141___1, k_5144), cljs_core.Next.Arity1IQ(ks_5142)
									continue
								}
							} else {
								{
									var m_5145 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5141___1)
									_ = m_5145
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5145).(float64), float64(7)) {
									} else {
										panic((&js.Error{("Assert failed: (= (count m) 7)")}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_5145, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var tm_5146 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_5146
				{
					var seq__4663_5147 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__4664_5148 interface{} = nil
					var count__4665_5149 = float64(0)
					var i__4666_5150 = float64(0)
					_, _, _, _ = seq__4663_5147, chunk__4664_5148, count__4665_5149, i__4666_5150
					for {
						if i__4666_5150 < count__4665_5149 {
							{
								var k_5151 = chunk__4664_5148.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4666_5150)
								_ = k_5151
								if cljs_core.X_EQ_.Arity2IIB(k_5151, cljs_core.Get.X_invoke_Arity2(tm_5146, k_5151)) {
								} else {
									panic((&js.Error{("Assert failed: (= k (get tm k))")}))
								}
								seq__4663_5147, chunk__4664_5148, count__4665_5149, i__4666_5150 = seq__4663_5147, chunk__4664_5148, count__4665_5149, (i__4666_5150 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___5152 = cljs_core.Seq.Arity1IQ(seq__4663_5147)
								_ = temp__4222__auto___5152
								if cljs_core.Truth_(temp__4222__auto___5152) {
									{
										var seq__4663_5153___1 = temp__4222__auto___5152
										_ = seq__4663_5153___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4663_5153___1) {
											{
												var c__954__auto___5154 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4663_5153___1)
												_ = c__954__auto___5154
												seq__4663_5147, chunk__4664_5148, count__4665_5149, i__4666_5150 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4663_5153___1), c__954__auto___5154, cljs_core.Count.X_invoke_Arity1(c__954__auto___5154).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_5155 = cljs_core.First.X_invoke_Arity1(seq__4663_5153___1)
												_ = k_5155
												if cljs_core.X_EQ_.Arity2IIB(k_5155, cljs_core.Get.X_invoke_Arity2(tm_5146, k_5155)) {
												} else {
													panic((&js.Error{("Assert failed: (= k (get tm k))")}))
												}
												seq__4663_5147, chunk__4664_5148, count__4665_5149, i__4666_5150 = cljs_core.Next.Arity1IQ(seq__4663_5153___1), nil, float64(0), float64(0)
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
					var m_5156 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5146)
					_ = m_5156
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5157 float64) {
						defer func() {
							if e4667 := recover(); e4667 != nil {
								if cljs_core.Value_(e4667).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4667
										_ = e
										return__5157 = float64(2)
									}
								} else {
									panic(e4667)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_5146, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5158 float64) {
						defer func() {
							if e4668 := recover(); e4668 != nil {
								if cljs_core.Value_(e4668).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4668
										_ = e
										return__5158 = float64(2)
									}
								} else {
									panic(e4668)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_5146, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5159 float64) {
						defer func() {
							if e4669 := recover(); e4669 != nil {
								if cljs_core.Value_(e4669).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4669
										_ = e
										return__5159 = float64(2)
									}
								} else {
									panic(e4669)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5146)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (persistent! tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5160 float64) {
						defer func() {
							if e4670 := recover(); e4670 != nil {
								if cljs_core.Value_(e4670).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4670
										_ = e
										return__5160 = float64(2)
									}
								} else {
									panic(e4670)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_5146)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (count tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_5156, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var m_5161 = cljs_core.Assoc.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentHashMap_EMPTY, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_ = m_5161
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5161, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5161, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5161).(float64), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 2)")}))
				}
				{
					var m_5162___1 = cljs_core.Dissoc.X_invoke_Arity2(m_5161, Fixed_hash_foo)
					_ = m_5162___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5162___1, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_5162___1, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5162___1).(float64), float64(1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 1)")}))
					}
				}
			}
			{
				var m_5163 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_5164___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m_5163, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_, _ = m_5163, m_5164___1
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5164___1).(float64), float64(102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 102)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5164___1, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5164___1, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				{
					var m_5165___2 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_5164___1, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_5165___2
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5165___2, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_5165___2, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5165___2).(float64), float64(98)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 98)")}))
					}
				}
			}
			{
				var m_5166 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_5167___1 = cljs_core.Transient.X_invoke_Arity1(m_5166)
				var m_5168___2 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_5167___1, Fixed_hash_foo, float64(1))
				var m_5169___3 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_5168___2, Fixed_hash_bar, float64(2))
				var m_5170___4 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m_5169___3)
				_, _, _, _, _ = m_5166, m_5167___1, m_5168___2, m_5169___3, m_5170___4
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5170___4).(float64), float64(102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 102)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5170___4, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5170___4, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				{
					var m_5171___5 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_5170___4, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_5171___5
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_5171___5, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_5171___5, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5171___5).(float64), float64(98)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 98)")}))
					}
				}
			}
			Array_map_conversion_threshold = cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD

			{
				var m1_5172 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_5173 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY)
				var i_5174 = float64(0)
				_, _, _ = m1_5172, m2_5173, i_5174
				for {
					if i_5174 < Array_map_conversion_threshold {
						m1_5172, m2_5173, i_5174 = cljs_core.Assoc.X_invoke_Arity3(m1_5172, i_5174, i_5174), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_5173, i_5174, i_5174), (i_5174 + float64(1))
						continue
					} else {
						{
							var m2_5175___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_5173)
							_ = m2_5175___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_5172).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m1) array-map-conversion-threshold)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_5175___1).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m2) array-map-conversion-threshold)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_5172, m2_5175___1) {
							} else {
								panic((&js.Error{("Assert failed: (= m1 m2)")}))
							}
							{
								var i_5176___1 = float64(0)
								_ = i_5176___1
								for {
									if i_5176___1 < Array_map_conversion_threshold {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__4671 = i_5176___1
											_ = G__4671
											return m1_5172.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4671)
										}(), i_5176___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__4672 = i_5176___1
											_ = G__4672
											return m2_5175___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4672)
										}(), i_5176___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m2 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_5172, i_5176___1), i_5176___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_5175___1, i_5176___1), i_5176___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m2 i) i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_5172, i_5176___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m1 i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_5175___1, i_5176___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m2 i)")}))
										}
										i_5176___1 = (i_5176___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_5172))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m1)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_5175___1))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m2)))")}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_5172, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{("Assert failed: (not (contains? (dissoc m1 3) 3))")}))
							}
						}
					}
					break
				}
			}
			{
				var m_5177 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_5177
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5177).(float64), float64(7)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 7)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5177, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
				}
			}
			{
				var m_5178 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_5178
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5178).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5178, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var m_5179 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_5179
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5179).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5179, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var tm_5180 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_5180
				{
					var tm_5181___1 interface{} = tm_5180
					var ks_5182 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_5181___1, ks_5182
					for {
						{
							var temp__4220__auto___5183 = cljs_core.First.X_invoke_Arity1(ks_5182)
							_ = temp__4220__auto___5183
							if cljs_core.Truth_(temp__4220__auto___5183) {
								{
									var k_5184 = temp__4220__auto___5183
									_ = k_5184
									tm_5181___1, ks_5182 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_5181___1, k_5184), cljs_core.Next.Arity1IQ(ks_5182)
									continue
								}
							} else {
								{
									var m_5185 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5181___1)
									_ = m_5185
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5185).(float64), float64(7)) {
									} else {
										panic((&js.Error{("Assert failed: (= (count m) 7)")}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_5185, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var tm_5186 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_5186
				{
					var seq__4673_5187 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__4674_5188 interface{} = nil
					var count__4675_5189 = float64(0)
					var i__4676_5190 = float64(0)
					_, _, _, _ = seq__4673_5187, chunk__4674_5188, count__4675_5189, i__4676_5190
					for {
						if i__4676_5190 < count__4675_5189 {
							{
								var k_5191 = chunk__4674_5188.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4676_5190)
								_ = k_5191
								if cljs_core.X_EQ_.Arity2IIB(k_5191, cljs_core.Get.X_invoke_Arity2(tm_5186, k_5191)) {
								} else {
									panic((&js.Error{("Assert failed: (= k (get tm k))")}))
								}
								seq__4673_5187, chunk__4674_5188, count__4675_5189, i__4676_5190 = seq__4673_5187, chunk__4674_5188, count__4675_5189, (i__4676_5190 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___5192 = cljs_core.Seq.Arity1IQ(seq__4673_5187)
								_ = temp__4222__auto___5192
								if cljs_core.Truth_(temp__4222__auto___5192) {
									{
										var seq__4673_5193___1 = temp__4222__auto___5192
										_ = seq__4673_5193___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4673_5193___1) {
											{
												var c__954__auto___5194 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4673_5193___1)
												_ = c__954__auto___5194
												seq__4673_5187, chunk__4674_5188, count__4675_5189, i__4676_5190 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4673_5193___1), c__954__auto___5194, cljs_core.Count.X_invoke_Arity1(c__954__auto___5194).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_5195 = cljs_core.First.X_invoke_Arity1(seq__4673_5193___1)
												_ = k_5195
												if cljs_core.X_EQ_.Arity2IIB(k_5195, cljs_core.Get.X_invoke_Arity2(tm_5186, k_5195)) {
												} else {
													panic((&js.Error{("Assert failed: (= k (get tm k))")}))
												}
												seq__4673_5187, chunk__4674_5188, count__4675_5189, i__4676_5190 = cljs_core.Next.Arity1IQ(seq__4673_5193___1), nil, float64(0), float64(0)
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
					var m_5196 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5186)
					_ = m_5196
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5197 float64) {
						defer func() {
							if e4677 := recover(); e4677 != nil {
								if cljs_core.Value_(e4677).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4677
										_ = e
										return__5197 = float64(2)
									}
								} else {
									panic(e4677)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_5186, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5198 float64) {
						defer func() {
							if e4678 := recover(); e4678 != nil {
								if cljs_core.Value_(e4678).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4678
										_ = e
										return__5198 = float64(2)
									}
								} else {
									panic(e4678)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_5186, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5199 float64) {
						defer func() {
							if e4679 := recover(); e4679 != nil {
								if cljs_core.Value_(e4679).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4679
										_ = e
										return__5199 = float64(2)
									}
								} else {
									panic(e4679)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_5186)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (persistent! tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__5200 float64) {
						defer func() {
							if e4680 := recover(); e4680 != nil {
								if cljs_core.Value_(e4680).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e4680
										_ = e
										return__5200 = float64(2)
									}
								} else {
									panic(e4680)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_5186)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (count tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_5196, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
					}
				}
			}
			{
				var m_5201 = cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))
				_ = m_5201
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5201).(float64), (float64(2) * Array_map_conversion_threshold)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) (* 2 array-map-conversion-threshold))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4681 = Array_map_conversion_threshold
					_ = G__4681
					return m_5201.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4681)
				}(), Array_map_conversion_threshold) {
				} else {
					panic((&js.Error{("Assert failed: (= (m array-map-conversion-threshold) array-map-conversion-threshold)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_5201, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity2(func(G__5202 *cljs_core.AFn, m_5201 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5202, 1, func(p1__4091_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__4091_SHARP_, p1__4091_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_5201), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (= m (into (.-EMPTY cljs.core/PersistentHashMap) (map (fn* [p1__4091#] (vector p1__4091# p1__4091#)) (range (* 2 array-map-conversion-threshold)))))")}))
				}
			}
			{
				var m1_5203 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_5204 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var i_5205 = float64(0)
				_, _, _ = m1_5203, m2_5204, i_5205
				for {
					if i_5205 < float64(100) {
						m1_5203, m2_5204, i_5205 = cljs_core.Assoc.X_invoke_Arity3(m1_5203, i_5205, i_5205), cljs_core.Assoc.X_invoke_Arity3(m2_5204, ("foo"+cljs_core.Str.X_invoke_Arity1(i_5205).(string)), i_5205), (i_5205 + float64(1))
						continue
					} else {
						if cljs_core.X_EQ_.Arity2IIB(m1_5203, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{("Assert failed: (= m1 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (range 100) (range 100))))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(m2_5204, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Map_.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.Str, "foo").(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{("Assert failed: (= m2 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (map (partial str \"foo\") (range 100)) (range 100))))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_5203).(float64), float64(100)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count m1) 100)")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_5204).(float64), float64(100)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count m2) 100)")}))
						}
					}
					break
				}
			}
			{
				var i_5206 = float64(0)
				var m_5207 interface{} = cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(-1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))
				var result_5208 interface{} = cljs_core.CljsCorePersistentVector_EMPTY
				_, _, _ = i_5206, m_5207, result_5208
				for {
					if i_5206 <= (cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) {
						i_5206, m_5207, result_5208 = (i_5206 + float64(1)), cljs_core.Assoc.X_invoke_Arity3(m_5207, i_5206, i_5206), cljs_core.Conj.X_invoke_Arity2(result_5208, cljs_core.Meta.X_invoke_Arity1(m_5207))
						continue
					} else {
						{
							var n_5209 = ((cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) + float64(1))
							var expected_5210 = cljs_core.Repeat.X_invoke_Arity2(n_5209, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})).(*cljs_core.CljsCoreLazySeq)
							_, _ = n_5209, expected_5210
							if cljs_core.X_EQ_.Arity2IIB(result_5208, expected_5210) {
							} else {
								panic((&js.Error{("Assert failed: (= result expected)")}))
							}
						}
					}
					break
				}
			}
			{
				var m1_5211 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_5212 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var m2_5213 = cljs_core.Sorted_map_by.X_invoke_ArityVariadic(c2_5212, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{})).(*cljs_core.CljsCorePersistentTreeMap)
				_, _, _ = m1_5211, c2_5212, m2_5213
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeMap)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(m1_5211)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeMap (type m1))")}))
				}
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeMap)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(m2_5213)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeMap (type m2))")}))
				}
				if reflect.DeepEqual(cljs_core.Compare, cljs_core.Native_get_instance_field.X_invoke_Arity2(m1_5211, "Comp")) {
				} else {
					panic((&js.Error{("Assert failed: (identical? compare (.-comp m1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(m1_5211).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count m1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(m2_5213).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count m2))")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(m1_5211)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq m1))")}))
				}
				{
					var m1_5214___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m1_5211, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					var m2_5215___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m2_5213, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					_, _ = m1_5214___1, m2_5215___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_5214___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m1) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_5215___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m2) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_5214___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m1) (list [:bar 2] [:foo 1] [:quux 3]))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_5215___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m2) (list [:quux 3] [:foo 1] [:bar 2]))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_5214___1), cljs_core.Rseq.Arity1IQ(m2_5215___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m1) (rseq m2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_5215___1), cljs_core.Rseq.Arity1IQ(m1_5214___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m2) (rseq m1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m1_5214___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= (conj m1 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m1_5214___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count (conj m1 [:wibble 4])) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m2_5215___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= (conj m2 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m2_5215___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count (conj m2 [:wibble 4])) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m1_5214___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})).(cljs_core.CljsCoreICollection).X_conj_Arity2(nil)) {
					} else {
						panic((&js.Error{("Assert failed: (= (map key (assoc m1 nil 4)) (list nil :bar :foo :quux))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m2_5215___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(nil).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (map key (assoc m2 nil 4)) (list :quux :foo :bar nil))")}))
					}
				}
			}
			{
				var m_5216 = cljs_core.Apply.X_invoke_Arity2(cljs_core.Sorted_map, cljs_core.Mapcat.X_invoke_ArityVariadic(func(G__5219 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5219, 1, func(p1__4092_SHARP_ interface{}) interface{} {
						return cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(p1__4092_SHARP_).(cljs_core.CljsCoreICollection).X_conj_Arity2(p1__4092_SHARP_)
					})
				}(&cljs_core.AFn{}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_ArityVariadic(cljs_core.Partial.X_invoke_Arity2(cljs_core.Apply, cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(10)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(20), float64(30)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(50), float64(60)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(30), float64(40)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(40), float64(50)}, nil})}, nil})}))})))
				var s1_5217 = cljs_core.Map_.X_invoke_Arity2(func(G__5220 *cljs_core.AFn, m_5216 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5220, 1, func(p1__4093_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__4093_SHARP_, p1__4093_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_5216), cljs_core.Range_.X_invoke_Arity1(float64(60)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				var s2_5218 = cljs_core.Map_.X_invoke_Arity2(func(G__5221 *cljs_core.AFn, m_5216 interface{}, s1_5217 *cljs_core.CljsCoreLazySeq) *cljs_core.AFn {
					return cljs_core.Fn(G__5221, 1, func(p1__4094_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__4094_SHARP_, p1__4094_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_5216, s1_5217), cljs_core.Range_.X_invoke_Arity3(float64(59), float64(-1), float64(-1)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				_, _, _ = m_5216, s1_5217, s2_5218
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_5216).(float64), float64(60)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 60)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m_5216), s1_5217) {
				} else {
					panic((&js.Error{("Assert failed: (= (seq m) s1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(m_5216), s2_5218) {
				} else {
					panic((&js.Error{("Assert failed: (= (rseq m) s2)")}))
				}
			}
			{
				var m_5222 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
				_ = m_5222
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m_5222, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(2), float64(3)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (dissoc m :foo) (hash-map :bar 2 :quux 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(m_5222, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))).(float64), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count (dissoc m :foo)) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(m_5222), cljs_core.Hash.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(1), float64(2), float64(3)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (hash m) (hash (hash-map :foo 1 :bar 2 :quux 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_5222, cljs_core.X_LT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m < :foo) (list [:bar 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_5222, cljs_core.X_LT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m <= :foo) (list [:bar 2] [:foo 1]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_5222, cljs_core.X_GT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m > :foo) (list [:quux 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_5222, cljs_core.X_GT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m >= :foo) (list [:foo 1] [:quux 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__5223 *cljs_core.AFn, m_5222 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5223, 1, func(p1__4095_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity2(func(G__5224 *cljs_core.AFn, m_5222 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__5224, 2, func(___ interface{}, x interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_5222), p1__4095_SHARP_)
					})
				}(&cljs_core.AFn{}, m_5222), m_5222).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__4095#] (reduce (fn [_ x] x) p1__4095#)) m) (list 2 1 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__5225 *cljs_core.AFn, m_5222 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5225, 1, func(p1__4096_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity3(func(G__5226 *cljs_core.AFn, m_5222 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__5226, 2, func(x interface{}, ___ interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_5222), float64(7), p1__4096_SHARP_)
					})
				}(&cljs_core.AFn{}, m_5222), m_5222).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__4096#] (reduce (fn [x _] x) 7 p1__4096#)) m) (list 7 7 7))")}))
				}
			}
			{
				var s1_5227 = cljs_core.Sorted_set.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_5228 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var s2_5229 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c2_5228, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c3_5230 = func(G__5233 *cljs_core.AFn, s1_5227 interface{}, c2_5228 cljs_core.CljsCoreIFn, s2_5229 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__5233, 2, func(p1__4097_SHARP_ interface{}, p2__4098_SHARP_ interface{}) interface{} {
						return cljs_core.Compare.Arity2IIF(cljs_core.Quot.X_invoke_Arity2(p1__4097_SHARP_, float64(2)).(float64), cljs_core.Quot.X_invoke_Arity2(p2__4098_SHARP_, float64(2)).(float64))
					})
				}(&cljs_core.AFn{}, s1_5227, c2_5228, s2_5229)
				var s3_5231 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c3_5230, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var s4_5232 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(cljs_core.X_LT_, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				_, _, _, _, _, _ = s1_5227, c2_5228, s2_5229, c3_5230, s3_5231, s4_5232
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(s1_5227)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeSet (type s1))")}))
				}
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(s2_5229)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeSet (type s2))")}))
				}
				if reflect.DeepEqual(cljs_core.Compare, s1_5227.(cljs_core.CljsCoreISorted).X_comparator_Arity1()) {
				} else {
					panic((&js.Error{("Assert failed: (identical? compare (-comparator s1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(s1_5227).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count s1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(s2_5229).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count s2))")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(s1_5227)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq s1))")}))
				}
				{
					var s1_5234___1 = cljs_core.Conj.X_invoke_ArityVariadic(s1_5227, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s2_5235___1 = cljs_core.Conj.X_invoke_ArityVariadic(s2_5229, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s3_5236___1 = cljs_core.Conj.X_invoke_ArityVariadic(s3_5231, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(7), float64(8), float64(9)}))
					var s4_5237___1 = cljs_core.Conj.X_invoke_ArityVariadic(s4_5232, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					_, _, _, _ = s1_5234___1, s2_5235___1, s3_5236___1, s4_5237___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_5234___1), cljs_core.Hash.X_invoke_Arity1(s2_5235___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (hash s1) (hash s2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_5234___1), cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (hash s1) (hash #{1 3 2}))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_5234___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq s1) (list 1 2 3))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_5234___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq s1) (list 3 2 1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_5235___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq s2) (list 3 2 1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_5235___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq s2) (list 1 2 3))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_5234___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s1) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_5235___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s2) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s3_5236___1).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s3) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(s3_5236___1, float64(0)), float64(1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get s3 0) 1)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5236___1, cljs_core.X_GT_, float64(5)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 > 5) (list 7 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5236___1, cljs_core.X_GT_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 > 6) (list 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5236___1, cljs_core.X_GT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 6) (list 7 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5236___1, cljs_core.X_LT_, float64(0)), cljs_core.CljsCoreList_EMPTY) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 0) (list))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5236___1, cljs_core.X_LT_, float64(5)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 5) (list 1 2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5236___1, cljs_core.X_LT_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 6) (list 1 2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_5236___1, cljs_core.X_LT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 <= 6) (list 1 2 7))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s3_5236___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 2 <= 6) (list 2 7))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s4_5237___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT_, float64(3)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s4 >= 2 < 3) (list 2))")}))
					}
					{
						var s1_5238___2 = cljs_core.Disj.X_invoke_Arity2(s1_5234___1, float64(2))
						var s2_5239___2 = cljs_core.Disj.X_invoke_Arity2(s2_5235___1, float64(2))
						_, _ = s1_5238___2, s2_5239___2
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_5238___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{("Assert failed: (= (seq s1) (list 1 3))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_5238___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{("Assert failed: (= (rseq s1) (list 3 1))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_5239___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{("Assert failed: (= (seq s2) (list 3 1))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_5239___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{("Assert failed: (= (rseq s2) (list 1 3))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_5238___2).(float64), float64(2)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count s1) 2)")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_5239___2).(float64), float64(2)) {
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
					return cljs_core.Fn(map__GT_Person, 1, func(G__4684 interface{}) interface{} {
						return (&CljsCore_testPerson{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}).X_invoke_Arity1(G__4684), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}).X_invoke_Arity1(G__4684), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__4684, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})})), nil})
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
					return cljs_core.Fn(map__GT_A, 1, func(G__4703 interface{}) interface{} {
						return (&CljsCore_testA{nil, cljs_core.Dissoc.X_invoke_Arity1(G__4703), nil})
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
					return cljs_core.Fn(map__GT_C, 1, func(G__4714 interface{}) interface{} {
						return (&CljsCore_testC{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__4714), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__4714), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}).X_invoke_Arity1(G__4714), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__4714, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})})), nil})
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
				var s_5240 = "abc"
				_ = s_5240
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_5240).Length) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (.-length s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_5240).Length) {
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
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_5240).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_("abc").Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring \"abc\" 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", func(G__5241 *cljs_core.AFn, s_5240 string) *cljs_core.AFn {
					return cljs_core.Fn(G__5241, 2, func(target4735 interface{}, start interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target4735, "Substring", []interface{}{start})
					})
				}(&cljs_core.AFn{}, s_5240).X_invoke_Arity2(s_5240, float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" ((memfn substring start) s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_5240).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s substring 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_5240).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s (substring 1)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_5240).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s (substring 1 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_5240).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring s 1 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_(s_5240).ToUpperCase()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" (. s (toUpperCase)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_("abc").ToUpperCase()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" (. \"abc\" (toUpperCase)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", func(G__5242 *cljs_core.AFn, s_5240 string) *cljs_core.AFn {
					return cljs_core.Fn(G__5242, 1, func(target4736 interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target4736, "ToUpperCase", []interface{}{})
					})
				}(&cljs_core.AFn{}, s_5240).X_invoke_Arity1(s_5240)) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" ((memfn toUpperCase) s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("BC", cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_5240).ToUpperCase(), "Substring", []interface{}{float64(1)})) {
				} else {
					panic((&js.Error{("Assert failed: (= \"BC\" (. (. s (toUpperCase)) substring 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_5240).ToUpperCase(), "Substring", []interface{}{float64(1)}), "Length")) {
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
					return cljs_core.Fn(map__GT_A2, 1, func(G__4739 interface{}) interface{} {
						return (&CljsCore_testA2{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__4739), nil, cljs_core.Dissoc.X_invoke_Arity2(G__4739, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
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
					return cljs_core.Fn(map__GT_B, 1, func(G__4754 interface{}) interface{} {
						return (&CljsCore_testB{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__4754), nil, cljs_core.Dissoc.X_invoke_Arity2(G__4754, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
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

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(func() *CljsCore_testT4767 {
				X__GT_t4767 = func(__GT_t4767 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_t4767, 2, func(test_stuff___1 interface{}, meta4768 interface{}) interface{} {
						return (&CljsCore_testT4767{test_stuff___1, meta4768})
					})
				}(&cljs_core.AFn{})

				return (&CljsCore_testT4767{test_stuff, nil})
			}(), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (meta (with-meta (reify IFoo (foo [this] :foo)) {:foo :bar})) {:foo :bar})")}))
			}
			Foo2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1064__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1065__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1066__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1067__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1068__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__
				return (&cljs_core.CljsCoreMultiFn{"foo2", cljs_core.Identity, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			Foo2.X_add_method_Arity3(float64(0), func(G__5243 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5243, 1, func(x interface{}) interface{} {
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
				var g_5244 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}))}, true).(*cljs_core.CljsCorePersistentHashSet)
				var h_5245 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil})}, true).(*cljs_core.CljsCorePersistentHashSet)
				_, _ = g_5244, h_5245
				if cljs_core.X_EQ_.Arity2IIB(g_5244, h_5245) {
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
				var fv_5246 = (&CljsCore_testFirst{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})})
				var fs_5247 = (&CljsCore_testFirst{"asdf"})
				_, _ = fv_5246, fs_5247
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fv_5246.X_invoke_Arity0()
				}(), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (fv) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fs_5247.X_invoke_Arity0()
				}(), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((`` + cljs_core.Str.X_invoke_Arity1(fs_5247).(string)), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (str fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_5246.X_get_first_Arity1(), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-get-first fv) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fs_5247.X_get_first_Arity1(), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (-get-first fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_5246.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-find-first fv [1]) 1)")}))
				}
				if reflect.DeepEqual(func() interface{} {
					var G__4782 = float64(1)
					_ = G__4782
					return fv_5246.X_invoke_Arity1(G__4782)
				}(), fv_5246) {
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
				var t_5248 = (&CljsCore_testDestructuringWithLocals{float64(1)})
				_ = t_5248
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(1)}, nil}), t_5248.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 1] (-find-first t [2 3]))")}))
				}
			}
			{
				var x_5249 = float64(1)
				_ = x_5249
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4786 = x_5249
					_ = G__4786
					switch G__4786 {
					case float64(1):
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})

					default:
						panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_5249).(string))}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x 1 :one) :one)")}))
				}
			}
			{
				var x_5251 = float64(1)
				_ = x_5251
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4787 = x_5251
					_ = G__4787
					switch G__4787 {
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
				var x_5253 = float64(1)
				_ = x_5253
				if cljs_core.X_EQ_.Arity2IIB(func() (return__5254 interface{}) {
					defer func() {
						if e4788 := recover(); e4788 != nil {
							if cljs_core.Value_(e4788).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e4788
									_ = e
									return__5254 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e4788)

							}
						}
					}()
					{
						{
							var G__4789 = x_5253
							_ = G__4789
							switch G__4789 {
							case float64(3):
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "three", Fqn: "three", X_hash: float64(-1651831795)})

							default:
								panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_5253).(string))}))

							}
						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (try (case x 3 :three) (catch js/Error e :fail)) :fail)")}))
				}
			}
			{
				var x_5256 = float64(1)
				_ = x_5256
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4790 = x_5256
					_ = G__4790
					switch G__4790 {
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
				var x_5258 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})
				_ = x_5258
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__4791 = x_5258
					_ = G__4791
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), G__4791) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
					} else {
						panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_5258).(string))}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x [:a :b] :ok) :ok)")}))
				}
			}
			{
				var a_5259 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = a_5259
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4792 = a_5259
					_ = G__4792
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__4792) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__4792) {
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
				var a_5260 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil})
				_ = a_5260
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__4793 = a_5260
					_ = G__4793
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__4793) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__4793) {
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
				var foo_5261 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = foo_5261
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__4794 = foo_5261
					_ = G__4794
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__4794) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__4794) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), G__4794) {
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
					var G__4795 = foo_5261
					_ = G__4795
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "d", Str: "d", X_hash: float64(-682293345), X_meta: nil}), G__4795) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__4795) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__4795) {
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
				var r_5262 = cljs_core.Range_.X_invoke_Arity1(float64(64)).(*cljs_core.CljsCoreRange)
				var v_5263 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, r_5262)
				_, _ = r_5262, v_5263
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_5263)), cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_5263))) {
				} else {
					panic((&js.Error{("Assert failed: (= (hash (seq v)) (hash (seq v)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Array_chunk.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)}).(*cljs_core.CljsCoreArrayChunk))) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (reduce + (array-chunk (array 1 2 3))))")}))
				}
				if cljs_core.Value_(cljs_core.Seq.Arity1IQ(v_5263)).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreChunkedSeq)(nil)).Elem()) {
				} else {
					panic((&js.Error{("Assert failed: (instance? ChunkedSeq (seq v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(r_5262, cljs_core.Seq.Arity1IQ(v_5263)) {
				} else {
					panic((&js.Error{("Assert failed: (= r (seq v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, r_5262).(*cljs_core.CljsCoreLazySeq), cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, v_5263).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (map inc r) (map inc v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, r_5262).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, v_5263).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (filter even? r) (filter even? v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, r_5262).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, v_5263).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (filter odd? r) (filter odd? v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic(r_5262, r_5262, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{r_5262})).(*cljs_core.CljsCoreLazySeq), cljs_core.Concat.X_invoke_ArityVariadic(v_5263, v_5263, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{v_5263})).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (concat r r r) (concat v v v))")}))
				}
				if cljs_core.Value_(cljs_core.Seq.Arity1IQ(v_5263)).Type().Implements(reflect.TypeOf((*cljs_core.CljsCoreIReduce)(nil)).Elem()) {
				} else {
					panic((&js.Error{("Assert failed: (satisfies? IReduce (seq v))")}))
				}
				if float64(2010) == cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_5263)))))).(float64) {
				} else {
					panic((&js.Error{("Assert failed: (== 2010 (reduce + (nnext (nnext (seq v)))))")}))
				}
				if float64(2020) == cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_5263)))))).(float64) {
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
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__5264 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5264, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), nil).(*cljs_core.CljsCoreCons)
				})
			}(&cljs_core.AFn{}), nil, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next (lazy-seq (cons 1 nil))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__5265 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5265, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), (&cljs_core.CljsCoreLazySeq{nil, func(G__5266 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__5266, 0, func() interface{} {
							return cljs_core.Cons.X_invoke_Arity2(float64(2), (&cljs_core.CljsCoreLazySeq{nil, func(G__5267 *cljs_core.AFn) *cljs_core.AFn {
								return cljs_core.Fn(G__5267, 0, func() interface{} {
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
					return cljs_core.Fn(map__GT_PrintMe, 1, func(G__4798 interface{}) interface{} {
						return (&CljsCore_testPrintMe{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__4798), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__4798), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__4798, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})})), nil})
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
				var seq__4815_5268 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(13)).(*cljs_core.CljsCoreRange))
				var chunk__4828_5269 interface{} = nil
				var count__4829_5270 = float64(0)
				var i__4830_5271 = float64(0)
				_, _, _, _ = seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271
				for {
					if i__4830_5271 < count__4829_5270 {
						{
							var month_5272 = chunk__4828_5269.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4830_5271)
							_ = month_5272
							{
								var seq__4831_5273 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
								var chunk__4836_5274 interface{} = nil
								var count__4837_5275 = float64(0)
								var i__4838_5276 = float64(0)
								_, _, _, _ = seq__4831_5273, chunk__4836_5274, count__4837_5275, i__4838_5276
								for {
									if i__4838_5276 < count__4837_5275 {
										{
											var day_5277 = chunk__4836_5274.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4838_5276)
											_ = day_5277
											{
												var seq__4839_5278 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
												var chunk__4840_5279 interface{} = nil
												var count__4841_5280 = float64(0)
												var i__4842_5281 = float64(0)
												_, _, _, _ = seq__4839_5278, chunk__4840_5279, count__4841_5280, i__4842_5281
												for {
													if i__4842_5281 < count__4841_5280 {
														{
															var hour_5282 = chunk__4840_5279.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4842_5281)
															_ = hour_5282
															{
																var pad_5283 = func(G__5285 *cljs_core.AFn, seq__4839_5278 interface{}, chunk__4840_5279 interface{}, count__4841_5280 float64, i__4842_5281 float64, seq__4831_5273 interface{}, chunk__4836_5274 interface{}, count__4837_5275 float64, i__4838_5276 float64, seq__4815_5268 interface{}, chunk__4828_5269 interface{}, count__4829_5270 float64, i__4830_5271 float64, hour_5282 interface{}, day_5277 interface{}, month_5272 interface{}) *cljs_core.AFn {
																	return cljs_core.Fn(G__5285, 1, func(n interface{}) interface{} {
																		if n.(float64) < float64(10) {
																			return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																		} else {
																			return n
																		}
																	})
																}(&cljs_core.AFn{}, seq__4839_5278, chunk__4840_5279, count__4841_5280, i__4842_5281, seq__4831_5273, chunk__4836_5274, count__4837_5275, i__4838_5276, seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271, hour_5282, day_5277, month_5272)
																var inst_5284 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5283.X_invoke_Arity1(month_5272)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5283.X_invoke_Arity1(day_5277)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5283.X_invoke_Arity1(hour_5282)).(string) + ":14:15.666-00:00")
																_, _ = pad_5283, inst_5284
																if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5284})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5284).(string) + "\"")) {
																} else {
																	panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																}
															}
															seq__4839_5278, chunk__4840_5279, count__4841_5280, i__4842_5281 = seq__4839_5278, chunk__4840_5279, count__4841_5280, (i__4842_5281 + float64(1))
															continue
														}
													} else {
														{
															var temp__4222__auto___5286 = cljs_core.Seq.Arity1IQ(seq__4839_5278)
															_ = temp__4222__auto___5286
															if cljs_core.Truth_(temp__4222__auto___5286) {
																{
																	var seq__4839_5287___1 = temp__4222__auto___5286
																	_ = seq__4839_5287___1
																	if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4839_5287___1) {
																		{
																			var c__954__auto___5288 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4839_5287___1)
																			_ = c__954__auto___5288
																			seq__4839_5278, chunk__4840_5279, count__4841_5280, i__4842_5281 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4839_5287___1), c__954__auto___5288, cljs_core.Count.X_invoke_Arity1(c__954__auto___5288).(float64), float64(0)
																			continue
																		}
																	} else {
																		{
																			var hour_5289 = cljs_core.First.X_invoke_Arity1(seq__4839_5287___1)
																			_ = hour_5289
																			{
																				var pad_5290 = func(G__5292 *cljs_core.AFn, seq__4839_5278 interface{}, chunk__4840_5279 interface{}, count__4841_5280 float64, i__4842_5281 float64, seq__4831_5273 interface{}, chunk__4836_5274 interface{}, count__4837_5275 float64, i__4838_5276 float64, seq__4815_5268 interface{}, chunk__4828_5269 interface{}, count__4829_5270 float64, i__4830_5271 float64, hour_5289 interface{}, seq__4839_5287___1 interface{}, temp__4222__auto___5286 cljs_core.CljsCoreISeq, day_5277 interface{}, month_5272 interface{}) *cljs_core.AFn {
																					return cljs_core.Fn(G__5292, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__4839_5278, chunk__4840_5279, count__4841_5280, i__4842_5281, seq__4831_5273, chunk__4836_5274, count__4837_5275, i__4838_5276, seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271, hour_5289, seq__4839_5287___1, temp__4222__auto___5286, day_5277, month_5272)
																				var inst_5291 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5290.X_invoke_Arity1(month_5272)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5290.X_invoke_Arity1(day_5277)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5290.X_invoke_Arity1(hour_5289)).(string) + ":14:15.666-00:00")
																				_, _ = pad_5290, inst_5291
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5291})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5291).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__4839_5278, chunk__4840_5279, count__4841_5280, i__4842_5281 = cljs_core.Next.Arity1IQ(seq__4839_5287___1), nil, float64(0), float64(0)
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
											seq__4831_5273, chunk__4836_5274, count__4837_5275, i__4838_5276 = seq__4831_5273, chunk__4836_5274, count__4837_5275, (i__4838_5276 + float64(1))
											continue
										}
									} else {
										{
											var temp__4222__auto___5293 = cljs_core.Seq.Arity1IQ(seq__4831_5273)
											_ = temp__4222__auto___5293
											if cljs_core.Truth_(temp__4222__auto___5293) {
												{
													var seq__4831_5294___1 = temp__4222__auto___5293
													_ = seq__4831_5294___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4831_5294___1) {
														{
															var c__954__auto___5295 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4831_5294___1)
															_ = c__954__auto___5295
															seq__4831_5273, chunk__4836_5274, count__4837_5275, i__4838_5276 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4831_5294___1), c__954__auto___5295, cljs_core.Count.X_invoke_Arity1(c__954__auto___5295).(float64), float64(0)
															continue
														}
													} else {
														{
															var day_5296 = cljs_core.First.X_invoke_Arity1(seq__4831_5294___1)
															_ = day_5296
															{
																var seq__4832_5297 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																var chunk__4833_5298 interface{} = nil
																var count__4834_5299 = float64(0)
																var i__4835_5300 = float64(0)
																_, _, _, _ = seq__4832_5297, chunk__4833_5298, count__4834_5299, i__4835_5300
																for {
																	if i__4835_5300 < count__4834_5299 {
																		{
																			var hour_5301 = chunk__4833_5298.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4835_5300)
																			_ = hour_5301
																			{
																				var pad_5302 = func(G__5304 *cljs_core.AFn, seq__4832_5297 interface{}, chunk__4833_5298 interface{}, count__4834_5299 float64, i__4835_5300 float64, seq__4831_5273 interface{}, chunk__4836_5274 interface{}, count__4837_5275 float64, i__4838_5276 float64, seq__4815_5268 interface{}, chunk__4828_5269 interface{}, count__4829_5270 float64, i__4830_5271 float64, hour_5301 interface{}, day_5296 interface{}, seq__4831_5294___1 interface{}, temp__4222__auto___5293 cljs_core.CljsCoreISeq, month_5272 interface{}) *cljs_core.AFn {
																					return cljs_core.Fn(G__5304, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__4832_5297, chunk__4833_5298, count__4834_5299, i__4835_5300, seq__4831_5273, chunk__4836_5274, count__4837_5275, i__4838_5276, seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271, hour_5301, day_5296, seq__4831_5294___1, temp__4222__auto___5293, month_5272)
																				var inst_5303 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5302.X_invoke_Arity1(month_5272)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5302.X_invoke_Arity1(day_5296)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5302.X_invoke_Arity1(hour_5301)).(string) + ":14:15.666-00:00")
																				_, _ = pad_5302, inst_5303
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5303})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5303).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__4832_5297, chunk__4833_5298, count__4834_5299, i__4835_5300 = seq__4832_5297, chunk__4833_5298, count__4834_5299, (i__4835_5300 + float64(1))
																			continue
																		}
																	} else {
																		{
																			var temp__4222__auto___5305___1 = cljs_core.Seq.Arity1IQ(seq__4832_5297)
																			_ = temp__4222__auto___5305___1
																			if cljs_core.Truth_(temp__4222__auto___5305___1) {
																				{
																					var seq__4832_5306___1 = temp__4222__auto___5305___1
																					_ = seq__4832_5306___1
																					if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4832_5306___1) {
																						{
																							var c__954__auto___5307 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4832_5306___1)
																							_ = c__954__auto___5307
																							seq__4832_5297, chunk__4833_5298, count__4834_5299, i__4835_5300 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4832_5306___1), c__954__auto___5307, cljs_core.Count.X_invoke_Arity1(c__954__auto___5307).(float64), float64(0)
																							continue
																						}
																					} else {
																						{
																							var hour_5308 = cljs_core.First.X_invoke_Arity1(seq__4832_5306___1)
																							_ = hour_5308
																							{
																								var pad_5309 = func(G__5311 *cljs_core.AFn, seq__4832_5297 interface{}, chunk__4833_5298 interface{}, count__4834_5299 float64, i__4835_5300 float64, seq__4831_5273 interface{}, chunk__4836_5274 interface{}, count__4837_5275 float64, i__4838_5276 float64, seq__4815_5268 interface{}, chunk__4828_5269 interface{}, count__4829_5270 float64, i__4830_5271 float64, hour_5308 interface{}, seq__4832_5306___1 interface{}, temp__4222__auto___5305___1 cljs_core.CljsCoreISeq, day_5296 interface{}, seq__4831_5294___1 interface{}, temp__4222__auto___5293 cljs_core.CljsCoreISeq, month_5272 interface{}) *cljs_core.AFn {
																									return cljs_core.Fn(G__5311, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__4832_5297, chunk__4833_5298, count__4834_5299, i__4835_5300, seq__4831_5273, chunk__4836_5274, count__4837_5275, i__4838_5276, seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271, hour_5308, seq__4832_5306___1, temp__4222__auto___5305___1, day_5296, seq__4831_5294___1, temp__4222__auto___5293, month_5272)
																								var inst_5310 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5309.X_invoke_Arity1(month_5272)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5309.X_invoke_Arity1(day_5296)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5309.X_invoke_Arity1(hour_5308)).(string) + ":14:15.666-00:00")
																								_, _ = pad_5309, inst_5310
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5310})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5310).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__4832_5297, chunk__4833_5298, count__4834_5299, i__4835_5300 = cljs_core.Next.Arity1IQ(seq__4832_5306___1), nil, float64(0), float64(0)
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
															seq__4831_5273, chunk__4836_5274, count__4837_5275, i__4838_5276 = cljs_core.Next.Arity1IQ(seq__4831_5294___1), nil, float64(0), float64(0)
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
							seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271 = seq__4815_5268, chunk__4828_5269, count__4829_5270, (i__4830_5271 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5312 = cljs_core.Seq.Arity1IQ(seq__4815_5268)
							_ = temp__4222__auto___5312
							if cljs_core.Truth_(temp__4222__auto___5312) {
								{
									var seq__4815_5313___1 = temp__4222__auto___5312
									_ = seq__4815_5313___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4815_5313___1) {
										{
											var c__954__auto___5314 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4815_5313___1)
											_ = c__954__auto___5314
											seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4815_5313___1), c__954__auto___5314, cljs_core.Count.X_invoke_Arity1(c__954__auto___5314).(float64), float64(0)
											continue
										}
									} else {
										{
											var month_5315 = cljs_core.First.X_invoke_Arity1(seq__4815_5313___1)
											_ = month_5315
											{
												var seq__4816_5316 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
												var chunk__4821_5317 interface{} = nil
												var count__4822_5318 = float64(0)
												var i__4823_5319 = float64(0)
												_, _, _, _ = seq__4816_5316, chunk__4821_5317, count__4822_5318, i__4823_5319
												for {
													if i__4823_5319 < count__4822_5318 {
														{
															var day_5320 = chunk__4821_5317.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4823_5319)
															_ = day_5320
															{
																var seq__4824_5321 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																var chunk__4825_5322 interface{} = nil
																var count__4826_5323 = float64(0)
																var i__4827_5324 = float64(0)
																_, _, _, _ = seq__4824_5321, chunk__4825_5322, count__4826_5323, i__4827_5324
																for {
																	if i__4827_5324 < count__4826_5323 {
																		{
																			var hour_5325 = chunk__4825_5322.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4827_5324)
																			_ = hour_5325
																			{
																				var pad_5326 = func(G__5328 *cljs_core.AFn, seq__4824_5321 interface{}, chunk__4825_5322 interface{}, count__4826_5323 float64, i__4827_5324 float64, seq__4816_5316 interface{}, chunk__4821_5317 interface{}, count__4822_5318 float64, i__4823_5319 float64, seq__4815_5268 interface{}, chunk__4828_5269 interface{}, count__4829_5270 float64, i__4830_5271 float64, hour_5325 interface{}, day_5320 interface{}, month_5315 interface{}, seq__4815_5313___1 interface{}, temp__4222__auto___5312 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																					return cljs_core.Fn(G__5328, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__4824_5321, chunk__4825_5322, count__4826_5323, i__4827_5324, seq__4816_5316, chunk__4821_5317, count__4822_5318, i__4823_5319, seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271, hour_5325, day_5320, month_5315, seq__4815_5313___1, temp__4222__auto___5312)
																				var inst_5327 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5326.X_invoke_Arity1(month_5315)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5326.X_invoke_Arity1(day_5320)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5326.X_invoke_Arity1(hour_5325)).(string) + ":14:15.666-00:00")
																				_, _ = pad_5326, inst_5327
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5327})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5327).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__4824_5321, chunk__4825_5322, count__4826_5323, i__4827_5324 = seq__4824_5321, chunk__4825_5322, count__4826_5323, (i__4827_5324 + float64(1))
																			continue
																		}
																	} else {
																		{
																			var temp__4222__auto___5329___1 = cljs_core.Seq.Arity1IQ(seq__4824_5321)
																			_ = temp__4222__auto___5329___1
																			if cljs_core.Truth_(temp__4222__auto___5329___1) {
																				{
																					var seq__4824_5330___1 = temp__4222__auto___5329___1
																					_ = seq__4824_5330___1
																					if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4824_5330___1) {
																						{
																							var c__954__auto___5331 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4824_5330___1)
																							_ = c__954__auto___5331
																							seq__4824_5321, chunk__4825_5322, count__4826_5323, i__4827_5324 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4824_5330___1), c__954__auto___5331, cljs_core.Count.X_invoke_Arity1(c__954__auto___5331).(float64), float64(0)
																							continue
																						}
																					} else {
																						{
																							var hour_5332 = cljs_core.First.X_invoke_Arity1(seq__4824_5330___1)
																							_ = hour_5332
																							{
																								var pad_5333 = func(G__5335 *cljs_core.AFn, seq__4824_5321 interface{}, chunk__4825_5322 interface{}, count__4826_5323 float64, i__4827_5324 float64, seq__4816_5316 interface{}, chunk__4821_5317 interface{}, count__4822_5318 float64, i__4823_5319 float64, seq__4815_5268 interface{}, chunk__4828_5269 interface{}, count__4829_5270 float64, i__4830_5271 float64, hour_5332 interface{}, seq__4824_5330___1 interface{}, temp__4222__auto___5329___1 cljs_core.CljsCoreISeq, day_5320 interface{}, month_5315 interface{}, seq__4815_5313___1 interface{}, temp__4222__auto___5312 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																									return cljs_core.Fn(G__5335, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__4824_5321, chunk__4825_5322, count__4826_5323, i__4827_5324, seq__4816_5316, chunk__4821_5317, count__4822_5318, i__4823_5319, seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271, hour_5332, seq__4824_5330___1, temp__4222__auto___5329___1, day_5320, month_5315, seq__4815_5313___1, temp__4222__auto___5312)
																								var inst_5334 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5333.X_invoke_Arity1(month_5315)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5333.X_invoke_Arity1(day_5320)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5333.X_invoke_Arity1(hour_5332)).(string) + ":14:15.666-00:00")
																								_, _ = pad_5333, inst_5334
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5334})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5334).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__4824_5321, chunk__4825_5322, count__4826_5323, i__4827_5324 = cljs_core.Next.Arity1IQ(seq__4824_5330___1), nil, float64(0), float64(0)
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
															seq__4816_5316, chunk__4821_5317, count__4822_5318, i__4823_5319 = seq__4816_5316, chunk__4821_5317, count__4822_5318, (i__4823_5319 + float64(1))
															continue
														}
													} else {
														{
															var temp__4222__auto___5336___1 = cljs_core.Seq.Arity1IQ(seq__4816_5316)
															_ = temp__4222__auto___5336___1
															if cljs_core.Truth_(temp__4222__auto___5336___1) {
																{
																	var seq__4816_5337___1 = temp__4222__auto___5336___1
																	_ = seq__4816_5337___1
																	if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4816_5337___1) {
																		{
																			var c__954__auto___5338 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4816_5337___1)
																			_ = c__954__auto___5338
																			seq__4816_5316, chunk__4821_5317, count__4822_5318, i__4823_5319 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4816_5337___1), c__954__auto___5338, cljs_core.Count.X_invoke_Arity1(c__954__auto___5338).(float64), float64(0)
																			continue
																		}
																	} else {
																		{
																			var day_5339 = cljs_core.First.X_invoke_Arity1(seq__4816_5337___1)
																			_ = day_5339
																			{
																				var seq__4817_5340 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																				var chunk__4818_5341 interface{} = nil
																				var count__4819_5342 = float64(0)
																				var i__4820_5343 = float64(0)
																				_, _, _, _ = seq__4817_5340, chunk__4818_5341, count__4819_5342, i__4820_5343
																				for {
																					if i__4820_5343 < count__4819_5342 {
																						{
																							var hour_5344 = chunk__4818_5341.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4820_5343)
																							_ = hour_5344
																							{
																								var pad_5345 = func(G__5347 *cljs_core.AFn, seq__4817_5340 interface{}, chunk__4818_5341 interface{}, count__4819_5342 float64, i__4820_5343 float64, seq__4816_5316 interface{}, chunk__4821_5317 interface{}, count__4822_5318 float64, i__4823_5319 float64, seq__4815_5268 interface{}, chunk__4828_5269 interface{}, count__4829_5270 float64, i__4830_5271 float64, hour_5344 interface{}, day_5339 interface{}, seq__4816_5337___1 interface{}, temp__4222__auto___5336___1 cljs_core.CljsCoreISeq, month_5315 interface{}, seq__4815_5313___1 interface{}, temp__4222__auto___5312 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																									return cljs_core.Fn(G__5347, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__4817_5340, chunk__4818_5341, count__4819_5342, i__4820_5343, seq__4816_5316, chunk__4821_5317, count__4822_5318, i__4823_5319, seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271, hour_5344, day_5339, seq__4816_5337___1, temp__4222__auto___5336___1, month_5315, seq__4815_5313___1, temp__4222__auto___5312)
																								var inst_5346 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5345.X_invoke_Arity1(month_5315)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5345.X_invoke_Arity1(day_5339)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5345.X_invoke_Arity1(hour_5344)).(string) + ":14:15.666-00:00")
																								_, _ = pad_5345, inst_5346
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5346})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5346).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__4817_5340, chunk__4818_5341, count__4819_5342, i__4820_5343 = seq__4817_5340, chunk__4818_5341, count__4819_5342, (i__4820_5343 + float64(1))
																							continue
																						}
																					} else {
																						{
																							var temp__4222__auto___5348___2 = cljs_core.Seq.Arity1IQ(seq__4817_5340)
																							_ = temp__4222__auto___5348___2
																							if cljs_core.Truth_(temp__4222__auto___5348___2) {
																								{
																									var seq__4817_5349___1 = temp__4222__auto___5348___2
																									_ = seq__4817_5349___1
																									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4817_5349___1) {
																										{
																											var c__954__auto___5350 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4817_5349___1)
																											_ = c__954__auto___5350
																											seq__4817_5340, chunk__4818_5341, count__4819_5342, i__4820_5343 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4817_5349___1), c__954__auto___5350, cljs_core.Count.X_invoke_Arity1(c__954__auto___5350).(float64), float64(0)
																											continue
																										}
																									} else {
																										{
																											var hour_5351 = cljs_core.First.X_invoke_Arity1(seq__4817_5349___1)
																											_ = hour_5351
																											{
																												var pad_5352 = func(G__5354 *cljs_core.AFn, seq__4817_5340 interface{}, chunk__4818_5341 interface{}, count__4819_5342 float64, i__4820_5343 float64, seq__4816_5316 interface{}, chunk__4821_5317 interface{}, count__4822_5318 float64, i__4823_5319 float64, seq__4815_5268 interface{}, chunk__4828_5269 interface{}, count__4829_5270 float64, i__4830_5271 float64, hour_5351 interface{}, seq__4817_5349___1 interface{}, temp__4222__auto___5348___2 cljs_core.CljsCoreISeq, day_5339 interface{}, seq__4816_5337___1 interface{}, temp__4222__auto___5336___1 cljs_core.CljsCoreISeq, month_5315 interface{}, seq__4815_5313___1 interface{}, temp__4222__auto___5312 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																													return cljs_core.Fn(G__5354, 1, func(n interface{}) interface{} {
																														if n.(float64) < float64(10) {
																															return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																														} else {
																															return n
																														}
																													})
																												}(&cljs_core.AFn{}, seq__4817_5340, chunk__4818_5341, count__4819_5342, i__4820_5343, seq__4816_5316, chunk__4821_5317, count__4822_5318, i__4823_5319, seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271, hour_5351, seq__4817_5349___1, temp__4222__auto___5348___2, day_5339, seq__4816_5337___1, temp__4222__auto___5336___1, month_5315, seq__4815_5313___1, temp__4222__auto___5312)
																												var inst_5353 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_5352.X_invoke_Arity1(month_5315)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_5352.X_invoke_Arity1(day_5339)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_5352.X_invoke_Arity1(hour_5351)).(string) + ":14:15.666-00:00")
																												_, _ = pad_5352, inst_5353
																												if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_5353})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_5353).(string) + "\"")) {
																												} else {
																													panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																												}
																											}
																											seq__4817_5340, chunk__4818_5341, count__4819_5342, i__4820_5343 = cljs_core.Next.Arity1IQ(seq__4817_5349___1), nil, float64(0), float64(0)
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
																			seq__4816_5316, chunk__4821_5317, count__4822_5318, i__4823_5319 = cljs_core.Next.Arity1IQ(seq__4816_5337___1), nil, float64(0), float64(0)
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
											seq__4815_5268, chunk__4828_5269, count__4829_5270, i__4830_5271 = cljs_core.Next.Arity1IQ(seq__4815_5313___1), nil, float64(0), float64(0)
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
				var uuid_str_5355 = "550e8400-e29b-41d4-a716-446655440000"
				var uuid_5356 = (&cljs_core.CljsCoreUUID{uuid_str_5355})
				_, _ = uuid_str_5355, uuid_5356
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{uuid_5356})).(string), ("#uuid \"" + cljs_core.Str.X_invoke_Arity1(uuid_str_5355).(string) + "\"")) {
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
					X__GT_t4848 = func(__GT_t4848 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t4848, 4, func(f___1 interface{}, baz___1 interface{}, test_stuff___1 interface{}, meta4849 interface{}) interface{} {
							return (&CljsCore_testT4848{f___1, baz___1, test_stuff___1, meta4849})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT4848{f, baz, test_stuff, nil})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(float64(2), Baz.X_invoke_Arity1(cljs_core.Inc).(*CljsCore_testT4848).X_bar_Arity2(float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (-bar (baz inc) 1))")}))
			}
			{
				var x_5357 = "original"
				_ = x_5357
				Original_closure_stmt = func(original_closure_stmt *cljs_core.AFn, x_5357 string) *cljs_core.AFn {
					return cljs_core.Fn(original_closure_stmt, 0, func() interface{} {
						return x_5357
					})
				}(&cljs_core.AFn{}, x_5357)

			}
			{
				var x_5358 = "overwritten"
				_ = x_5358
				if cljs_core.X_EQ_.Arity2IIB("original", Original_closure_stmt.X_invoke_Arity0().(string)) {
				} else {
					panic((&js.Error{("Assert failed: (= \"original\" (original-closure-stmt))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB("original", func() string {
				var x = "original"
				var oce = func(G__5359 *cljs_core.AFn, x string) *cljs_core.AFn {
					return cljs_core.Fn(G__5359, 0, func() interface{} {
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
					var x_5360___1 = func(G__5361 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__5361, 0, func() interface{} {
							return "overwritten"
						})
					}(&cljs_core.AFn{})
					_ = x_5360___1
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
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)}), cljs_core.Reduce_kv.X_invoke_Arity3(func(G__5362 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__5362, 3, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
								return cljs_core.Reduced.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)})).(*cljs_core.CljsCoreReduced)
							})
						}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentVector_EMPTY, data)) {
						} else {
							panic((&js.Error{("Assert failed: (= :reduced (reduce-kv (fn [_ _ _] (reduced :reduced)) [] data))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sort.X_invoke_Arity1(expect), cljs_core.Sort.X_invoke_Arity1(cljs_core.Reduce_kv.X_invoke_Arity3(func(G__5363 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__5363, 3, func(r interface{}, k interface{}, v interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), func() (return__5364 interface{}) {
				defer func() {
					if e4853 := recover(); e4853 != nil {
						if cljs_core.Value_(e4853).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreExceptionInfo)(nil)).Elem()) {
							{
								var e = e4853
								_ = e
								return__5364 = cljs_core.Ex_data.X_invoke_Arity1(e)
							}
						} else {
							panic(e4853)

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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(1), float64(2)}, nil}), func(G__5365 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5365, 0, func(more__ ...interface{}) interface{} {
					var more = cljs_core.Seq.Arity1IQ(more__[0])
					_ = more
					return more
				})
			}(&cljs_core.AFn{}).X_invoke_Arity3(js.Undefined, float64(1), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= [js/undefined 1 2] ((fn [& more] more) js/undefined 1 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(4), float64(5)}, nil}), func(G__5366 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5366, 2, func(a_b_more__ ...interface{}) interface{} {
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
				var fs_5367 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = fs_5367
				{
					var seq__4854_5368 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity1(float64(4)).(*cljs_core.CljsCoreRange))
					var chunk__4856_5369 interface{} = nil
					var count__4857_5370 = float64(0)
					var i__4858_5371 = float64(0)
					_, _, _, _ = seq__4854_5368, chunk__4856_5369, count__4857_5370, i__4858_5371
					for {
						if i__4858_5371 < count__4857_5370 {
							{
								var x_5372 = chunk__4856_5369.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4858_5371)
								_ = x_5372
								{
									var y_5373 = (x_5372.(float64) + float64(1))
									var f_5374 = func(G__5375 *cljs_core.AFn, seq__4854_5368 interface{}, chunk__4856_5369 interface{}, count__4857_5370 float64, i__4858_5371 float64, y_5373 float64, x_5372 interface{}, fs_5367 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
										return cljs_core.Fn(G__5375, 0, func() interface{} {
											return y_5373
										})
									}(&cljs_core.AFn{}, seq__4854_5368, chunk__4856_5369, count__4857_5370, i__4858_5371, y_5373, x_5372, fs_5367)
									_, _ = y_5373, f_5374
									cljs_core.Swap_BANG_.X_invoke_Arity3(fs_5367, cljs_core.Conj, f_5374)
								}
								seq__4854_5368, chunk__4856_5369, count__4857_5370, i__4858_5371 = seq__4854_5368, chunk__4856_5369, count__4857_5370, (i__4858_5371 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___5376 = cljs_core.Seq.Arity1IQ(seq__4854_5368)
								_ = temp__4222__auto___5376
								if cljs_core.Truth_(temp__4222__auto___5376) {
									{
										var seq__4854_5377___1 = temp__4222__auto___5376
										_ = seq__4854_5377___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4854_5377___1) {
											{
												var c__954__auto___5378 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4854_5377___1)
												_ = c__954__auto___5378
												seq__4854_5368, chunk__4856_5369, count__4857_5370, i__4858_5371 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4854_5377___1), c__954__auto___5378, cljs_core.Count.X_invoke_Arity1(c__954__auto___5378).(float64), float64(0)
												continue
											}
										} else {
											{
												var x_5379 = cljs_core.First.X_invoke_Arity1(seq__4854_5377___1)
												_ = x_5379
												{
													var y_5380 = (x_5379.(float64) + float64(1))
													var f_5381 = func(G__5382 *cljs_core.AFn, seq__4854_5368 interface{}, chunk__4856_5369 interface{}, count__4857_5370 float64, i__4858_5371 float64, y_5380 float64, x_5379 interface{}, seq__4854_5377___1 interface{}, temp__4222__auto___5376 cljs_core.CljsCoreISeq, fs_5367 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
														return cljs_core.Fn(G__5382, 0, func() interface{} {
															return y_5380
														})
													}(&cljs_core.AFn{}, seq__4854_5368, chunk__4856_5369, count__4857_5370, i__4858_5371, y_5380, x_5379, seq__4854_5377___1, temp__4222__auto___5376, fs_5367)
													_, _ = y_5380, f_5381
													cljs_core.Swap_BANG_.X_invoke_Arity3(fs_5367, cljs_core.Conj, f_5381)
												}
												seq__4854_5368, chunk__4856_5369, count__4857_5370, i__4858_5371 = cljs_core.Next.Arity1IQ(seq__4854_5377___1), nil, float64(0), float64(0)
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__5383 *cljs_core.AFn, fs_5367 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__5383, 1, func(p1__4099_SHARP_ interface{}) interface{} {
						{
							return p1__4099_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
						}
					})
				}(&cljs_core.AFn{}, fs_5367), cljs_core.Deref.X_invoke_Arity1(fs_5367)).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCoreLazySeq{nil, func(G__5384 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5384, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCoreLazySeq{nil, func(G__5385 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5385, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreLazySeq{nil, func(G__5386 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5386, 0, func() interface{} {
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
					var sb__1124__auto__ = (&goog_string.StringBuffer{})
					_ = sb__1124__auto__
					{
						var _STAR_print_fn_STAR_4860_5387 = cljs_core.X_STAR_print_fn_STAR_
						_ = _STAR_print_fn_STAR_4860_5387
						func() {
							defer func() {
								cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_4860_5387

							}()
							{
								cljs_core.X_STAR_print_fn_STAR_ = func(G__5388 *cljs_core.AFn, _STAR_print_fn_STAR_4860_5387 interface{}, sb__1124__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
									return cljs_core.Fn(G__5388, 1, func(x__1125__auto__ interface{}) interface{} {
										return sb__1124__auto__.Append(x__1125__auto__)
									})
								}(&cljs_core.AFn{}, _STAR_print_fn_STAR_4860_5387, sb__1124__auto__)

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
					return (`` + cljs_core.Str.X_invoke_Arity1(sb__1124__auto__).(string))
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
			if (cljs_core.First.X_invoke_Arity1(cljs_core.Filter.X_invoke_Arity2(func(G__5389 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5389, 1, func(p1__4100_SHARP_ interface{}) interface{} {
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
				var a_5390 = func() *CljsCore_testT4861 {
					X__GT_t4861 = func(__GT_t4861 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t4861, 2, func(test_stuff___1 interface{}, meta4862 interface{}) interface{} {
							return (&CljsCore_testT4861{test_stuff___1, meta4862})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT4861{test_stuff, nil})
				}()
				var b_5391 = func() *CljsCore_testT4864 {
					X__GT_t4864 = func(__GT_t4864 *cljs_core.AFn, a_5390 *CljsCore_testT4861) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t4864, 3, func(a___1 interface{}, test_stuff___1 interface{}, meta4865 interface{}) interface{} {
							return (&CljsCore_testT4864{a___1, test_stuff___1, meta4865})
						})
					}(&cljs_core.AFn{}, a_5390)

					return (&CljsCore_testT4864{a_5390, test_stuff, nil})
				}()
				var s_5392 = cljs_core.Set.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(128)).(*cljs_core.CljsCoreRange))
				_, _, _ = a_5390, b_5391, s_5392
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_5392, a_5390, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_5391}))), a_5390)), a_5390), cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_5392, a_5390, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_5391}))), a_5390)), a_5390)) {
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
				var seq__4867_5393 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__4868_5394 interface{} = nil
				var count__4869_5395 = float64(0)
				var i__4870_5396 = float64(0)
				_, _, _, _ = seq__4867_5393, chunk__4868_5394, count__4869_5395, i__4870_5396
				for {
					if i__4870_5396 < count__4869_5395 {
						{
							var n_5397 = chunk__4868_5394.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4870_5396)
							_ = n_5397
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5398 interface{}) {
								defer func() {
									if e4873 := recover(); e4873 != nil {
										if cljs_core.Value_(e4873).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4873
												_ = e
												return__5398 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4873)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5397, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5399 interface{}) {
								defer func() {
									if e4874 := recover(); e4874 != nil {
										if cljs_core.Value_(e4874).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4874
												_ = e
												return__5399 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4874)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_5397, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))")}))
							}
							seq__4867_5393, chunk__4868_5394, count__4869_5395, i__4870_5396 = seq__4867_5393, chunk__4868_5394, count__4869_5395, (i__4870_5396 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5400 = cljs_core.Seq.Arity1IQ(seq__4867_5393)
							_ = temp__4222__auto___5400
							if cljs_core.Truth_(temp__4222__auto___5400) {
								{
									var seq__4867_5401___1 = temp__4222__auto___5400
									_ = seq__4867_5401___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4867_5401___1) {
										{
											var c__954__auto___5402 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4867_5401___1)
											_ = c__954__auto___5402
											seq__4867_5393, chunk__4868_5394, count__4869_5395, i__4870_5396 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4867_5401___1), c__954__auto___5402, cljs_core.Count.X_invoke_Arity1(c__954__auto___5402).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_5403 = cljs_core.First.X_invoke_Arity1(seq__4867_5401___1)
											_ = n_5403
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5404 interface{}) {
												defer func() {
													if e4875 := recover(); e4875 != nil {
														if cljs_core.Value_(e4875).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4875
																_ = e
																return__5404 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4875)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5403, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5405 interface{}) {
												defer func() {
													if e4876 := recover(); e4876 != nil {
														if cljs_core.Value_(e4876).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4876
																_ = e
																return__5405 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4876)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_5403, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))")}))
											}
											seq__4867_5393, chunk__4868_5394, count__4869_5395, i__4870_5396 = cljs_core.Next.Arity1IQ(seq__4867_5401___1), nil, float64(0), float64(0)
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
				var seq__4877_5406 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__4878_5407 interface{} = nil
				var count__4879_5408 = float64(0)
				var i__4880_5409 = float64(0)
				_, _, _, _ = seq__4877_5406, chunk__4878_5407, count__4879_5408, i__4880_5409
				for {
					if i__4880_5409 < count__4879_5408 {
						{
							var n_5410 = chunk__4878_5407.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4880_5409)
							_ = n_5410
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5411 interface{}) {
								defer func() {
									if e4883 := recover(); e4883 != nil {
										if cljs_core.Value_(e4883).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4883
												_ = e
												return__5411 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4883)

										}
									}
								}()
								{
									return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5410, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
							}
							seq__4877_5406, chunk__4878_5407, count__4879_5408, i__4880_5409 = seq__4877_5406, chunk__4878_5407, count__4879_5408, (i__4880_5409 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5412 = cljs_core.Seq.Arity1IQ(seq__4877_5406)
							_ = temp__4222__auto___5412
							if cljs_core.Truth_(temp__4222__auto___5412) {
								{
									var seq__4877_5413___1 = temp__4222__auto___5412
									_ = seq__4877_5413___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4877_5413___1) {
										{
											var c__954__auto___5414 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4877_5413___1)
											_ = c__954__auto___5414
											seq__4877_5406, chunk__4878_5407, count__4879_5408, i__4880_5409 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4877_5413___1), c__954__auto___5414, cljs_core.Count.X_invoke_Arity1(c__954__auto___5414).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_5415 = cljs_core.First.X_invoke_Arity1(seq__4877_5413___1)
											_ = n_5415
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5416 interface{}) {
												defer func() {
													if e4884 := recover(); e4884 != nil {
														if cljs_core.Value_(e4884).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4884
																_ = e
																return__5416 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4884)

														}
													}
												}()
												{
													return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5415, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
											}
											seq__4877_5406, chunk__4878_5407, count__4879_5408, i__4880_5409 = cljs_core.Next.Arity1IQ(seq__4877_5413___1), nil, float64(0), float64(0)
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
				var map__4885_5417 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__4885_5418___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4885_5417) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4885_5417)
					} else {
						return map__4885_5417
					}
				}()
				var b_5419 = cljs_core.Get.X_invoke_Arity2(map__4885_5418___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a_5420 = cljs_core.Get.X_invoke_Arity2(map__4885_5418___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__4885_5417, map__4885_5418___1, b_5419, a_5420
				if cljs_core.X_EQ_.Arity2IIB(float64(1), a_5420) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), b_5419) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 b)")}))
				}
			}
			{
				var map__4886_5421 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__4886_5422___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4886_5421) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4886_5421)
					} else {
						return map__4886_5421
					}
				}()
				var d_5423 = cljs_core.Get.X_invoke_Arity2(map__4886_5422___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_5424 = cljs_core.Get.X_invoke_Arity2(map__4886_5422___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__4886_5421, map__4886_5422___1, d_5423, b_5424
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_5424) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_5423) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__4887_5425 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__4887_5426___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4887_5425) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4887_5425)
					} else {
						return map__4887_5425
					}
				}()
				var d_5427 = cljs_core.Get.X_invoke_Arity2(map__4887_5426___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_5428 = cljs_core.Get.X_invoke_Arity2(map__4887_5426___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__4887_5425, map__4887_5426___1, d_5427, b_5428
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_5428) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_5427) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__4888_5429 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}), float64(1), (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}), float64(2)}, nil})
				var map__4888_5430___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4888_5429) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4888_5429)
					} else {
						return map__4888_5429
					}
				}()
				var d_5431 = cljs_core.Get.X_invoke_Arity2(map__4888_5430___1, (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}))
				var b_5432 = cljs_core.Get.X_invoke_Arity2(map__4888_5430___1, (&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}))
				_, _, _, _ = map__4888_5429, map__4888_5430___1, d_5431, b_5432
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_5432) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_5431) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__4889_5433 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}), float64(2)}, nil})
				var map__4889_5434___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__4889_5433) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__4889_5433)
					} else {
						return map__4889_5433
					}
				}()
				var y_5435 = cljs_core.Get.X_invoke_Arity2(map__4889_5434___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}))
				var x_5436 = cljs_core.Get.X_invoke_Arity2(map__4889_5434___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}))
				_, _, _, _ = map__4889_5433, map__4889_5434___1, y_5435, x_5436
				if cljs_core.X_EQ_.Arity2IIB(x_5436, float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= x 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(y_5435, float64(2)) {
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
								arr, names = cljs_core.Conj.X_invoke_Arity2(arr, func(G__5437 *cljs_core.AFn, arr interface{}, names interface{}, name interface{}) *cljs_core.AFn {
									return cljs_core.Fn(G__5437, 0, func() interface{} {
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
				var sb__1124__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1124__auto__
				{
					var _STAR_print_fn_STAR_4890_5438 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_4890_5438
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_4890_5438

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__5439 *cljs_core.AFn, _STAR_print_fn_STAR_4890_5438 interface{}, sb__1124__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__5439, 1, func(x__1125__auto__ interface{}) interface{} {
									return sb__1124__auto__.Append(x__1125__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_4890_5438, sb__1124__auto__)

							{
								var seq__4891_5440 interface{} = cljs_core.Seq.Arity1IQ(Cljs_739.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil})))
								var chunk__4892_5441 interface{} = nil
								var count__4893_5442 = float64(0)
								var i__4894_5443 = float64(0)
								_, _, _, _ = seq__4891_5440, chunk__4892_5441, count__4893_5442, i__4894_5443
								for {
									if i__4894_5443 < count__4893_5442 {
										{
											var fn_5444 = chunk__4892_5441.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4894_5443)
											_ = fn_5444
											{
												fn_5444.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
											}
											seq__4891_5440, chunk__4892_5441, count__4893_5442, i__4894_5443 = seq__4891_5440, chunk__4892_5441, count__4893_5442, (i__4894_5443 + float64(1))
											continue
										}
									} else {
										{
											var temp__4222__auto___5445 = cljs_core.Seq.Arity1IQ(seq__4891_5440)
											_ = temp__4222__auto___5445
											if cljs_core.Truth_(temp__4222__auto___5445) {
												{
													var seq__4891_5446___1 = temp__4222__auto___5445
													_ = seq__4891_5446___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4891_5446___1) {
														{
															var c__954__auto___5447 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4891_5446___1)
															_ = c__954__auto___5447
															seq__4891_5440, chunk__4892_5441, count__4893_5442, i__4894_5443 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4891_5446___1), c__954__auto___5447, cljs_core.Count.X_invoke_Arity1(c__954__auto___5447).(float64), float64(0)
															continue
														}
													} else {
														{
															var fn_5448 = cljs_core.First.X_invoke_Arity1(seq__4891_5446___1)
															_ = fn_5448
															{
																fn_5448.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
															}
															seq__4891_5440, chunk__4892_5441, count__4893_5442, i__4894_5443 = cljs_core.Next.Arity1IQ(seq__4891_5446___1), nil, float64(0), float64(0)
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
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1124__auto__).(string))
			}(), ":a\n:b\n:c\n:d\n") {
			} else {
				panic((&js.Error{("Assert failed: (= (with-out-str (doseq [fn (cljs-739 [] [:a :b :c :d])] (fn))) \":a\\n:b\\n:c\\n:d\\n\")")}))
			}
			{
				var seq__4895_5449 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__4896_5450 interface{} = nil
				var count__4897_5451 = float64(0)
				var i__4898_5452 = float64(0)
				_, _, _, _ = seq__4895_5449, chunk__4896_5450, count__4897_5451, i__4898_5452
				for {
					if i__4898_5452 < count__4897_5451 {
						{
							var n_5453 = chunk__4896_5450.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4898_5452)
							_ = n_5453
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5453)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get [1 2] n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5454 interface{}) {
								defer func() {
									if e4901 := recover(); e4901 != nil {
										if cljs_core.Value_(e4901).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4901
												_ = e
												return__5454 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4901)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5453)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5453, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get [1 2] n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5455 interface{}) {
								defer func() {
									if e4902 := recover(); e4902 != nil {
										if cljs_core.Value_(e4902).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4902
												_ = e
												return__5455 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4902)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5453, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5453)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (subvec [1 2] 1) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5456 interface{}) {
								defer func() {
									if e4903 := recover(); e4903 != nil {
										if cljs_core.Value_(e4903).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4903
												_ = e
												return__5456 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4903)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5453)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5453, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (subvec [1 2] 1) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5457 interface{}) {
								defer func() {
									if e4904 := recover(); e4904 != nil {
										if cljs_core.Value_(e4904).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4904
												_ = e
												return__5457 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4904)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5453, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5453)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (transient [1 2]) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5458 interface{}) {
								defer func() {
									if e4905 := recover(); e4905 != nil {
										if cljs_core.Value_(e4905).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4905
												_ = e
												return__5458 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4905)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5453)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5453, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (transient [1 2]) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5459 interface{}) {
								defer func() {
									if e4906 := recover(); e4906 != nil {
										if cljs_core.Value_(e4906).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4906
												_ = e
												return__5459 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4906)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5453, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5453)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (range 1 3) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5460 interface{}) {
								defer func() {
									if e4907 := recover(); e4907 != nil {
										if cljs_core.Value_(e4907).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4907
												_ = e
												return__5460 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4907)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5453)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5453, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (range 1 3) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5461 interface{}) {
								defer func() {
									if e4908 := recover(); e4908 != nil {
										if cljs_core.Value_(e4908).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e4908
												_ = e
												return__5461 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e4908)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5453, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))")}))
							}
							seq__4895_5449, chunk__4896_5450, count__4897_5451, i__4898_5452 = seq__4895_5449, chunk__4896_5450, count__4897_5451, (i__4898_5452 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5462 = cljs_core.Seq.Arity1IQ(seq__4895_5449)
							_ = temp__4222__auto___5462
							if cljs_core.Truth_(temp__4222__auto___5462) {
								{
									var seq__4895_5463___1 = temp__4222__auto___5462
									_ = seq__4895_5463___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4895_5463___1) {
										{
											var c__954__auto___5464 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4895_5463___1)
											_ = c__954__auto___5464
											seq__4895_5449, chunk__4896_5450, count__4897_5451, i__4898_5452 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4895_5463___1), c__954__auto___5464, cljs_core.Count.X_invoke_Arity1(c__954__auto___5464).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_5465 = cljs_core.First.X_invoke_Arity1(seq__4895_5463___1)
											_ = n_5465
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5465)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get [1 2] n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5466 interface{}) {
												defer func() {
													if e4909 := recover(); e4909 != nil {
														if cljs_core.Value_(e4909).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4909
																_ = e
																return__5466 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4909)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5465)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5465, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get [1 2] n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5467 interface{}) {
												defer func() {
													if e4910 := recover(); e4910 != nil {
														if cljs_core.Value_(e4910).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4910
																_ = e
																return__5467 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4910)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_5465, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5465)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (subvec [1 2] 1) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5468 interface{}) {
												defer func() {
													if e4911 := recover(); e4911 != nil {
														if cljs_core.Value_(e4911).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4911
																_ = e
																return__5468 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4911)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5465)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5465, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (subvec [1 2] 1) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5469 interface{}) {
												defer func() {
													if e4912 := recover(); e4912 != nil {
														if cljs_core.Value_(e4912).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4912
																_ = e
																return__5469 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4912)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_5465, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5465)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (transient [1 2]) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5470 interface{}) {
												defer func() {
													if e4913 := recover(); e4913 != nil {
														if cljs_core.Value_(e4913).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4913
																_ = e
																return__5470 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4913)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5465)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5465, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (transient [1 2]) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5471 interface{}) {
												defer func() {
													if e4914 := recover(); e4914 != nil {
														if cljs_core.Value_(e4914).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4914
																_ = e
																return__5471 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4914)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_5465, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5465)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (range 1 3) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5472 interface{}) {
												defer func() {
													if e4915 := recover(); e4915 != nil {
														if cljs_core.Value_(e4915).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4915
																_ = e
																return__5472 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4915)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5465)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5465, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (range 1 3) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__5473 interface{}) {
												defer func() {
													if e4916 := recover(); e4916 != nil {
														if cljs_core.Value_(e4916).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e4916
																_ = e
																return__5473 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e4916)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_5465, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))")}))
											}
											seq__4895_5449, chunk__4896_5450, count__4897_5451, i__4898_5452 = cljs_core.Next.Arity1IQ(seq__4895_5463___1), nil, float64(0), float64(0)
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
				var x_5474 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}).X_invoke_Arity1(cljs_core.Meta.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}).X_invoke_Arity1(cljs_core.Deref.X_invoke_Arity1(Cljs_780))))
				_ = x_5474
				if cljs_core.Vector_QMARK_.Arity1IB(x_5474) {
				} else {
					panic((&js.Error{("Assert failed: (vector? x)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(x_5474, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= x [1 2 3])")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3((&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}), "ToString", []interface{}{}), "550e8400-e29b-41d4-a716-446655440000") {
			} else {
				panic((&js.Error{("Assert failed: (= (.toString #uuid \"550e8400-e29b-41d4-a716-446655440000\") \"550e8400-e29b-41d4-a716-446655440000\")")}))
			}
			{
				var seq__4917_5475 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__4918_5476 interface{} = nil
				var count__4919_5477 = float64(0)
				var i__4920_5478 = float64(0)
				_, _, _, _ = seq__4917_5475, chunk__4918_5476, count__4919_5477, i__4920_5478
				for {
					if i__4920_5478 < count__4919_5477 {
						{
							var m_5479 = chunk__4918_5476.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4920_5478)
							_ = m_5479
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__5480 interface{}) {
								defer func() {
									if e4921 := recover(); e4921 != nil {
										if cljs_core.Value_(e4921).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var ___ = e4921
												_ = ___
												return__5480 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
											}
										} else {
											panic(e4921)

										}
									}
								}()
								{
									return cljs_core.Conj.X_invoke_Arity2(m_5479, "foo")
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5479, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m [:foo 1]))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5479, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m {:foo 1}))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5479, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m (list [:foo 1])))")}))
							}
							seq__4917_5475, chunk__4918_5476, count__4919_5477, i__4920_5478 = seq__4917_5475, chunk__4918_5476, count__4919_5477, (i__4920_5478 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5481 = cljs_core.Seq.Arity1IQ(seq__4917_5475)
							_ = temp__4222__auto___5481
							if cljs_core.Truth_(temp__4222__auto___5481) {
								{
									var seq__4917_5482___1 = temp__4222__auto___5481
									_ = seq__4917_5482___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4917_5482___1) {
										{
											var c__954__auto___5483 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4917_5482___1)
											_ = c__954__auto___5483
											seq__4917_5475, chunk__4918_5476, count__4919_5477, i__4920_5478 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4917_5482___1), c__954__auto___5483, cljs_core.Count.X_invoke_Arity1(c__954__auto___5483).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_5484 = cljs_core.First.X_invoke_Arity1(seq__4917_5482___1)
											_ = m_5484
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__5485 interface{}) {
												defer func() {
													if e4922 := recover(); e4922 != nil {
														if cljs_core.Value_(e4922).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var ___ = e4922
																_ = ___
																return__5485 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
															}
														} else {
															panic(e4922)

														}
													}
												}()
												{
													return cljs_core.Conj.X_invoke_Arity2(m_5484, "foo")
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5484, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m [:foo 1]))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5484, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m {:foo 1}))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_5484, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m (list [:foo 1])))")}))
											}
											seq__4917_5475, chunk__4918_5476, count__4919_5477, i__4920_5478 = cljs_core.Next.Arity1IQ(seq__4917_5482___1), nil, float64(0), float64(0)
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
				var seq__4923_5486 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Array_map, cljs_core.Hash_map, cljs_core.Sorted_map}, nil}))
				var chunk__4924_5487 interface{} = nil
				var count__4925_5488 = float64(0)
				var i__4926_5489 = float64(0)
				_, _, _, _ = seq__4923_5486, chunk__4924_5487, count__4925_5488, i__4926_5489
				for {
					if i__4926_5489 < count__4925_5488 {
						{
							var mt_5490 = chunk__4924_5487.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__4926_5489)
							_ = mt_5490
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
								var G__4927 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
								var G__4928 = float64(1)
								_, _ = G__4927, G__4928
								return mt_5490.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__4927, G__4928)
							}(), func(make_seq *cljs_core.AFn, seq__4923_5486 interface{}, chunk__4924_5487 interface{}, count__4925_5488 float64, i__4926_5489 float64, mt_5490 interface{}) *cljs_core.AFn {
								return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
									if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
										X__GT_t4934 = func(__GT_t4934 *cljs_core.AFn, seq__4923_5486 interface{}, chunk__4924_5487 interface{}, count__4925_5488 float64, i__4926_5489 float64, mt_5490 interface{}) *cljs_core.AFn {
											return cljs_core.Fn(__GT_t4934, 9, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, i__4926___1 interface{}, count__4925___1 interface{}, chunk__4924___1 interface{}, seq__4923___1 interface{}, test_stuff___1 interface{}, meta4935 interface{}) interface{} {
												return (&CljsCore_testT4934{from_seq___1, make_seq___1, mt___1, i__4926___1, count__4925___1, chunk__4924___1, seq__4923___1, test_stuff___1, meta4935})
											})
										}(&cljs_core.AFn{}, seq__4923_5486, chunk__4924_5487, count__4925_5488, i__4926_5489, mt_5490)

										return (&CljsCore_testT4934{from_seq, make_seq, mt_5490, i__4926_5489, count__4925_5488, chunk__4924_5487, seq__4923_5486, test_stuff, nil})
									} else {
										return nil
									}
								})
							}(&cljs_core.AFn{}, seq__4923_5486, chunk__4924_5487, count__4925_5488, i__4926_5489, mt_5490).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))")}))
							}
							seq__4923_5486, chunk__4924_5487, count__4925_5488, i__4926_5489 = seq__4923_5486, chunk__4924_5487, count__4925_5488, (i__4926_5489 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___5491 = cljs_core.Seq.Arity1IQ(seq__4923_5486)
							_ = temp__4222__auto___5491
							if cljs_core.Truth_(temp__4222__auto___5491) {
								{
									var seq__4923_5492___1 = temp__4222__auto___5491
									_ = seq__4923_5492___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__4923_5492___1) {
										{
											var c__954__auto___5493 = cljs_core.Chunk_first.X_invoke_Arity1(seq__4923_5492___1)
											_ = c__954__auto___5493
											seq__4923_5486, chunk__4924_5487, count__4925_5488, i__4926_5489 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__4923_5492___1), c__954__auto___5493, cljs_core.Count.X_invoke_Arity1(c__954__auto___5493).(float64), float64(0)
											continue
										}
									} else {
										{
											var mt_5494 = cljs_core.First.X_invoke_Arity1(seq__4923_5492___1)
											_ = mt_5494
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
												var G__4939 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
												var G__4940 = float64(1)
												_, _ = G__4939, G__4940
												return mt_5494.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__4939, G__4940)
											}(), func(make_seq *cljs_core.AFn, seq__4923_5486 interface{}, chunk__4924_5487 interface{}, count__4925_5488 float64, i__4926_5489 float64, mt_5494 interface{}, seq__4923_5492___1 interface{}, temp__4222__auto___5491 cljs_core.CljsCoreISeq) *cljs_core.AFn {
												return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
													if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
														X__GT_t4946 = func(__GT_t4946 *cljs_core.AFn, seq__4923_5486 interface{}, chunk__4924_5487 interface{}, count__4925_5488 float64, i__4926_5489 float64, mt_5494 interface{}, seq__4923_5492___1 interface{}, temp__4222__auto___5491 cljs_core.CljsCoreISeq) *cljs_core.AFn {
															return cljs_core.Fn(__GT_t4946, 10, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, temp__4222__auto_____1 interface{}, i__4926___1 interface{}, count__4925___1 interface{}, chunk__4924___1 interface{}, seq__4923___2 interface{}, test_stuff___1 interface{}, meta4947 interface{}) interface{} {
																return (&CljsCore_testT4946{from_seq___1, make_seq___1, mt___1, temp__4222__auto_____1, i__4926___1, count__4925___1, chunk__4924___1, seq__4923___2, test_stuff___1, meta4947})
															})
														}(&cljs_core.AFn{}, seq__4923_5486, chunk__4924_5487, count__4925_5488, i__4926_5489, mt_5494, seq__4923_5492___1, temp__4222__auto___5491)

														return (&CljsCore_testT4946{from_seq, make_seq, mt_5494, temp__4222__auto___5491, i__4926_5489, count__4925_5488, chunk__4924_5487, seq__4923_5492___1, test_stuff, nil})
													} else {
														return nil
													}
												})
											}(&cljs_core.AFn{}, seq__4923_5486, chunk__4924_5487, count__4925_5488, i__4926_5489, mt_5494, seq__4923_5492___1, temp__4222__auto___5491).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))")}))
											}
											seq__4923_5486, chunk__4924_5487, count__4925_5488, i__4926_5489 = cljs_core.Next.Arity1IQ(seq__4923_5492___1), nil, float64(0), float64(0)
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
				var _STAR_print_length_STAR_4951 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4951
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4951

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
				var _STAR_print_length_STAR_4952 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4952
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4952

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
				var _STAR_print_length_STAR_4953 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4953
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4953

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
				var _STAR_print_length_STAR_4954 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4954
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4954

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
				var _STAR_print_length_STAR_4955 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4955
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4955

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
				var _STAR_print_length_STAR_4956 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_4956
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_4956

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
							var G__4958 = func() interface{} {
								if cljs_core.Value_(value).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
									return cljs_core.Native_get_instance_field.X_invoke_Arity2(value, "Fqn")
								} else {
									return nil
								}
							}()
							_ = G__4958
							switch G__4958 {
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
				var not_strings_5496 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, nil, float64(1), func(G__5497 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__5497, 0, func() interface{} {
						return nil
					})
				}(&cljs_core.AFn{})}, nil})
				_ = not_strings_5496
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__5498 *cljs_core.AFn, not_strings_5496 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__5498, 1, func(p1__4101_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__5499 interface{}) {
							defer func() {
								if e4959 := recover(); e4959 != nil {
									if cljs_core.Value_(e4959).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e4959
											_ = ___
											return__5499 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e4959)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__4101_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_5496), not_strings_5496) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__4101#] (= :failed (try (re-find #\".\" p1__4101#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__5500 *cljs_core.AFn, not_strings_5496 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__5500, 1, func(p1__4102_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__5501 interface{}) {
							defer func() {
								if e4960 := recover(); e4960 != nil {
									if cljs_core.Value_(e4960).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e4960
											_ = ___
											return__5501 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e4960)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__4102_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_5496), not_strings_5496) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__4102#] (= :failed (try (re-matches #\".\" p1__4102#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__5502 *cljs_core.AFn, not_strings_5496 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__5502, 1, func(p1__4103_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__5503 interface{}) {
							defer func() {
								if e4961 := recover(); e4961 != nil {
									if cljs_core.Value_(e4961).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e4961
											_ = ___
											return__5503 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e4961)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__4103_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_5496), not_strings_5496) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__4103#] (= :failed (try (re-find #\"nomatch\" p1__4103#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__5504 *cljs_core.AFn, not_strings_5496 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__5504, 1, func(p1__4104_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__5505 interface{}) {
							defer func() {
								if e4962 := recover(); e4962 != nil {
									if cljs_core.Value_(e4962).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e4962
											_ = ___
											return__5505 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e4962)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__4104_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_5496), not_strings_5496) {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(func(G__5506 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5506, 1, func(p1__4105_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Take_while.X_invoke_Arity1(func(G__5507 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5507, 1, func(p1__4106_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Drop_while.X_invoke_Arity1(func(G__5508 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5508, 1, func(p1__4107_SHARP_ interface{}) interface{} {
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
				var ret_5509 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))
				_ = ret_5509
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_5509)) && (cljs_core.X_EQ_.Arity2IIB(ret_5509, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList))) {
				} else {
					panic((&js.Error{("Assert failed: (and (vector? ret) (= ret (quote (1 2 3))))")}))
				}
			}
			{
				var ret_5510 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				_ = ret_5510
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_5510)) && (cljs_core.X_EQ_.Arity2IIB(ret_5510, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(2), float64(4), float64(6), float64(8)})).(*cljs_core.CljsCoreList))) {
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
			Xform = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Dedupe.X_invoke_Arity0().(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Partition_all.X_invoke_Arity1(float64(3)).(cljs_core.CljsCoreIFn), cljs_core.Partition_by.X_invoke_Arity1(func(G__5511 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5511, 1, func(p1__4108_SHARP_ interface{}) interface{} {
					return (cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, p1__4108_SHARP_).(float64) < float64(7))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Flatten).(cljs_core.CljsCoreIFn), cljs_core.Random_sample.X_invoke_Arity1(1.0).(cljs_core.CljsCoreIFn), cljs_core.Take_nth.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Keep.X_invoke_Arity1(func(G__5512 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5512, 1, func(p1__4109_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__4109_SHARP_) {
						return (p1__4109_SHARP_.(float64) * p1__4109_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Keep_indexed.X_invoke_Arity1(func(G__5513 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5513, 2, func(p1__4110_SHARP_ interface{}, p2__4111_SHARP_ interface{}) interface{} {
					if cljs_core.Even_QMARK_.Arity1IB(p1__4110_SHARP_) {
						return (p1__4110_SHARP_.(float64) * p2__4111_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Replace.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(2), "two", float64(6), "six", float64(18), "eighteen"}, nil})).(cljs_core.CljsCoreIFn), cljs_core.Take.X_invoke_Arity1(float64(11)).(cljs_core.CljsCoreIFn), cljs_core.Take_while.X_invoke_Arity1(func(G__5514 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5514, 1, func(p1__4112_SHARP_ interface{}) interface{} {
					return cljs_core.Not_EQ_.Arity2IIB(float64(300), p1__4112_SHARP_)
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Drop.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Drop_while.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Remove.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn)})).(*cljs_core.AFn)

			Data = cljs_core.Vec.X_invoke_Arity1(cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(18)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(20)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(Xform, Data), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(36), float64(200), float64(10)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence xform data) (quote (36 200 10)))")}))
			}
			Xf = cljs_core.Map_.X_invoke_Arity1(func(G__5515 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__5515, 2, func(p1__4113_SHARP_ interface{}, p2__4114_SHARP_ interface{}) interface{} {
					return (p1__4113_SHARP_.(float64) + p2__4114_SHARP_.(float64))
				})
			}(&cljs_core.AFn{})).(*cljs_core.AFn)

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_ArityVariadic(Xf, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0)}, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence xf [0 0] [1 2]) [1 2])")}))
			}
			{
				var xs = (&cljs_core.CljsCorePersistentVector{nil, float64(21), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(44), float64(43), float64(42), float64(41), float64(40), float64(39), float64(38), float64(37), float64(36), float64(35), float64(34), float64(33), float64(32), float64(31), float64(30), float64(29), float64(28), float64(27), float64(26), float64(25), float64(24)}, nil})
				_ = xs
				{
					var m interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.Zipmap.X_invoke_Arity2(xs, cljs_core.Repeat.X_invoke_Arity1(float64(1)).(*cljs_core.CljsCoreLazySeq)))
					var xs___1 interface{} = xs
					_, _ = m, xs___1
					for {
						{
							var temp__4220__auto__ = cljs_core.First.X_invoke_Arity1(xs___1)
							_ = temp__4220__auto__
							if cljs_core.Truth_(temp__4220__auto__) {
								{
									var x = temp__4220__auto__
									_ = x
									if cljs_core.Contains_QMARK_.Arity2IIB(m, x) {
										m, xs___1 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(m, x), cljs_core.Next.Arity1IQ(xs___1)
										continue
									} else {
										panic(cljs_core.Ex_info.X_invoke_Arity2("CLJS-849 regression!", (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "m", Fqn: "m", X_hash: float64(1632677161)}), cljs_core.Persistent_BANG_.X_invoke_Arity1(m), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "xs", Fqn: "xs", X_hash: float64(649443341)}), xs___1}, nil})).(*cljs_core.CljsCoreExceptionInfo))
									}
								}
							} else {
								return nil
							}
						}
					}
				}
			}
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

func (this__758__auto__ *CljsCore_testPerson) X_lookup_Arity2(k__759__auto__ interface{}) interface{} {
	return this__758__auto__.X_lookup_Arity3(k__759__auto__, nil)
}

func (this__760__auto__ *CljsCore_testPerson) X_lookup_Arity3(k4683 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__4686 = func() interface{} {
			if cljs_core.Value_(k4683).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k4683, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4686
		switch G__4686 {
		case "lastname":
			return this__760__auto__.Lastname

		case "firstname":
			return this__760__auto__.Firstname

		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k4683, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testPerson) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__5517 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5517, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, cljs_core.Pr_writer, "", " ", "", opts__776__auto__, keyval__778__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__777__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, pr_pair__777__auto__, "#cljs.core-test.Person{", ", ", "}", opts__776__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), this__774__auto__.Firstname}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), this__774__auto__.Lastname}, nil})}, nil}), this__774__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testPerson) CljsCoreIMeta__() {}

func (this__756__auto__ *CljsCore_testPerson) X_meta_Arity1() interface{} {
	return this__756__auto__.X__meta
}

func (_ *CljsCore_testPerson) CljsCoreICloneable__() {}

func (this__752__auto__ *CljsCore_testPerson) X_clone_Arity1() interface{} {
	return (&CljsCore_testPerson{this__752__auto__.Firstname, this__752__auto__.Lastname, this__752__auto__.X__meta, this__752__auto__.X__extmap, this__752__auto__.X__hash})
}

func (_ *CljsCore_testPerson) CljsCoreICounted__() {}

func (this__762__auto__ *CljsCore_testPerson) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__762__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testPerson) CljsCoreIHash__() {}

func (this__753__auto__ *CljsCore_testPerson) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = this__753__auto__.X__hash
		_ = h__577__auto__
		if !(cljs_core.Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__753__auto__).(float64)
				_ = h__577__auto_____1
				this__753__auto__.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIEquiv__() {}

func (this__754__auto__ *CljsCore_testPerson) X_equiv_Arity2(other__755__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__755__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__754__auto__), cljs_core.Type_.X_invoke_Arity1(other__755__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__754__auto__, other__755__auto__)))
		} else {
			return and__159__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testPerson) CljsCoreIRecord__() {}

func (_ *CljsCore_testPerson) CljsCoreIMap__() {}

func (this__769__auto__ *CljsCore_testPerson) X_dissoc_Arity2(k__770__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), nil}, nil}, nil}), k__770__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__769__auto__), this__769__auto__.X__meta), k__770__auto__)
	} else {
		return (&CljsCore_testPerson{this__769__auto__.Firstname, this__769__auto__.Lastname, this__769__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__769__auto__.X__extmap, k__770__auto__)), nil})
	}
}

func (_ *CljsCore_testPerson) CljsCoreIAssociative__() {}

func (this__765__auto__ *CljsCore_testPerson) X_assoc_Arity3(k__766__auto__ interface{}, G__4682 interface{}) interface{} {
	{
		var pred__4694 = cljs_core.Keyword_identical_QMARK_
		var expr__4695 = k__766__auto__
		_, _ = pred__4694, expr__4695
		if cljs_core.Truth_(func() interface{} {
			var G__4697 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)})
			var G__4698 = expr__4695
			_, _ = G__4697, G__4698
			return pred__4694.X_invoke_Arity2(G__4697, G__4698)
		}()) {
			return (&CljsCore_testPerson{G__4682, this__765__auto__.Lastname, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__4699 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})
				var G__4700 = expr__4695
				_, _ = G__4699, G__4700
				return pred__4694.X_invoke_Arity2(G__4699, G__4700)
			}()) {
				return (&CljsCore_testPerson{this__765__auto__.Firstname, G__4682, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPerson{this__765__auto__.Firstname, this__765__auto__.Lastname, this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__4682), nil})
			}
		}
	}
}

func (this__767__auto__ *CljsCore_testPerson) X_contains_key_QMARK__Arity2(k__768__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testPerson) CljsCoreISeqable__() {}

func (this__772__auto__ *CljsCore_testPerson) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), this__772__auto__.Firstname}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), this__772__auto__.Lastname}, nil})}, nil}), this__772__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testPerson) CljsCoreIWithMeta__() {}

func (this__757__auto__ *CljsCore_testPerson) X_with_meta_Arity2(G__4682 interface{}) interface{} {
	return (&CljsCore_testPerson{this__757__auto__.Firstname, this__757__auto__.Lastname, G__4682, this__757__auto__.X__extmap, this__757__auto__.X__hash})
}

func (_ *CljsCore_testPerson) CljsCoreICollection__() {}

func (this__763__auto__ *CljsCore_testPerson) X_conj_Arity2(entry__764__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__764__auto__) {
		return this__763__auto__.X_assoc_Arity3(entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__763__auto__, entry__764__auto__)
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

func (this__758__auto__ *CljsCore_testA) X_lookup_Arity2(k__759__auto__ interface{}) interface{} {
	return this__758__auto__.X_lookup_Arity3(k__759__auto__, nil)
}

func (this__760__auto__ *CljsCore_testA) X_lookup_Arity3(k4702 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__4705 = k4702
		_ = G__4705
		switch G__4705 {
		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k4702, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testA) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testA) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__5519 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5519, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, cljs_core.Pr_writer, "", " ", "", opts__776__auto__, keyval__778__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__777__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, pr_pair__777__auto__, "#cljs.core-test.A{", ", ", "}", opts__776__auto__, cljs_core.Concat.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, this__774__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testA) CljsCoreIMeta__() {}

func (this__756__auto__ *CljsCore_testA) X_meta_Arity1() interface{} {
	return this__756__auto__.X__meta
}

func (_ *CljsCore_testA) CljsCoreICloneable__() {}

func (this__752__auto__ *CljsCore_testA) X_clone_Arity1() interface{} {
	return (&CljsCore_testA{this__752__auto__.X__meta, this__752__auto__.X__extmap, this__752__auto__.X__hash})
}

func (_ *CljsCore_testA) CljsCoreICounted__() {}

func (this__762__auto__ *CljsCore_testA) X_count_Arity1() float64 {
	return (float64(0) + cljs_core.Count.X_invoke_Arity1(this__762__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testA) CljsCoreIHash__() {}

func (this__753__auto__ *CljsCore_testA) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = this__753__auto__.X__hash
		_ = h__577__auto__
		if !(cljs_core.Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__753__auto__).(float64)
				_ = h__577__auto_____1
				this__753__auto__.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testA) CljsCoreIEquiv__() {}

func (this__754__auto__ *CljsCore_testA) X_equiv_Arity2(other__755__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__755__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__754__auto__), cljs_core.Type_.X_invoke_Arity1(other__755__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__754__auto__, other__755__auto__)))
		} else {
			return and__159__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testA) CljsCoreIRecord__() {}

func (_ *CljsCore_testA) CljsCoreIMap__() {}

func (this__769__auto__ *CljsCore_testA) X_dissoc_Arity2(k__770__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, k__770__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__769__auto__), this__769__auto__.X__meta), k__770__auto__)
	} else {
		return (&CljsCore_testA{this__769__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__769__auto__.X__extmap, k__770__auto__)), nil})
	}
}

func (_ *CljsCore_testA) CljsCoreIAssociative__() {}

func (this__765__auto__ *CljsCore_testA) X_assoc_Arity3(k__766__auto__ interface{}, G__4701 interface{}) interface{} {
	{
		var pred__4709 = cljs_core.Keyword_identical_QMARK_
		var expr__4710 = k__766__auto__
		_, _ = pred__4709, expr__4710
		return (&CljsCore_testA{this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__4701), nil})
	}
}

func (this__767__auto__ *CljsCore_testA) X_contains_key_QMARK__Arity2(k__768__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testA) CljsCoreISeqable__() {}

func (this__772__auto__ *CljsCore_testA) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, this__772__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testA) CljsCoreIWithMeta__() {}

func (this__757__auto__ *CljsCore_testA) X_with_meta_Arity2(G__4701 interface{}) interface{} {
	return (&CljsCore_testA{G__4701, this__757__auto__.X__extmap, this__757__auto__.X__hash})
}

func (_ *CljsCore_testA) CljsCoreICollection__() {}

func (this__763__auto__ *CljsCore_testA) X_conj_Arity2(entry__764__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__764__auto__) {
		return this__763__auto__.X_assoc_Arity3(entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__763__auto__, entry__764__auto__)
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

func (this__758__auto__ *CljsCore_testC) X_lookup_Arity2(k__759__auto__ interface{}) interface{} {
	return this__758__auto__.X_lookup_Arity3(k__759__auto__, nil)
}

func (this__760__auto__ *CljsCore_testC) X_lookup_Arity3(k4713 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__4716 = func() interface{} {
			if cljs_core.Value_(k4713).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k4713, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4716
		switch G__4716 {
		case "c":
			return this__760__auto__.C

		case "b":
			return this__760__auto__.B

		case "a":
			return this__760__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k4713, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testC) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testC) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__5521 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5521, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, cljs_core.Pr_writer, "", " ", "", opts__776__auto__, keyval__778__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__777__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, pr_pair__777__auto__, "#cljs.core-test.C{", ", ", "}", opts__776__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__774__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__774__auto__.B}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), this__774__auto__.C}, nil})}, nil}), this__774__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testC) CljsCoreIMeta__() {}

func (this__756__auto__ *CljsCore_testC) X_meta_Arity1() interface{} {
	return this__756__auto__.X__meta
}

func (_ *CljsCore_testC) CljsCoreICloneable__() {}

func (this__752__auto__ *CljsCore_testC) X_clone_Arity1() interface{} {
	return (&CljsCore_testC{this__752__auto__.A, this__752__auto__.B, this__752__auto__.C, this__752__auto__.X__meta, this__752__auto__.X__extmap, this__752__auto__.X__hash})
}

func (_ *CljsCore_testC) CljsCoreICounted__() {}

func (this__762__auto__ *CljsCore_testC) X_count_Arity1() float64 {
	return (float64(3) + cljs_core.Count.X_invoke_Arity1(this__762__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testC) CljsCoreIHash__() {}

func (this__753__auto__ *CljsCore_testC) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = this__753__auto__.X__hash
		_ = h__577__auto__
		if !(cljs_core.Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__753__auto__).(float64)
				_ = h__577__auto_____1
				this__753__auto__.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testC) CljsCoreIEquiv__() {}

func (this__754__auto__ *CljsCore_testC) X_equiv_Arity2(other__755__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__755__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__754__auto__), cljs_core.Type_.X_invoke_Arity1(other__755__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__754__auto__, other__755__auto__)))
		} else {
			return and__159__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testC) CljsCoreIRecord__() {}

func (_ *CljsCore_testC) CljsCoreIMap__() {}

func (this__769__auto__ *CljsCore_testC) X_dissoc_Arity2(k__770__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__770__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__769__auto__), this__769__auto__.X__meta), k__770__auto__)
	} else {
		return (&CljsCore_testC{this__769__auto__.A, this__769__auto__.B, this__769__auto__.C, this__769__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__769__auto__.X__extmap, k__770__auto__)), nil})
	}
}

func (_ *CljsCore_testC) CljsCoreIAssociative__() {}

func (this__765__auto__ *CljsCore_testC) X_assoc_Arity3(k__766__auto__ interface{}, G__4712 interface{}) interface{} {
	{
		var pred__4726 = cljs_core.Keyword_identical_QMARK_
		var expr__4727 = k__766__auto__
		_, _ = pred__4726, expr__4727
		if cljs_core.Truth_(func() interface{} {
			var G__4729 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__4730 = expr__4727
			_, _ = G__4729, G__4730
			return pred__4726.X_invoke_Arity2(G__4729, G__4730)
		}()) {
			return (&CljsCore_testC{G__4712, this__765__auto__.B, this__765__auto__.C, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__4731 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__4732 = expr__4727
				_, _ = G__4731, G__4732
				return pred__4726.X_invoke_Arity2(G__4731, G__4732)
			}()) {
				return (&CljsCore_testC{this__765__auto__.A, G__4712, this__765__auto__.C, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
			} else {
				if cljs_core.Truth_(func() interface{} {
					var G__4733 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})
					var G__4734 = expr__4727
					_, _ = G__4733, G__4734
					return pred__4726.X_invoke_Arity2(G__4733, G__4734)
				}()) {
					return (&CljsCore_testC{this__765__auto__.A, this__765__auto__.B, G__4712, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
				} else {
					return (&CljsCore_testC{this__765__auto__.A, this__765__auto__.B, this__765__auto__.C, this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__4712), nil})
				}
			}
		}
	}
}

func (this__767__auto__ *CljsCore_testC) X_contains_key_QMARK__Arity2(k__768__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testC) CljsCoreISeqable__() {}

func (this__772__auto__ *CljsCore_testC) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__772__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__772__auto__.B}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), this__772__auto__.C}, nil})}, nil}), this__772__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testC) CljsCoreIWithMeta__() {}

func (this__757__auto__ *CljsCore_testC) X_with_meta_Arity2(G__4712 interface{}) interface{} {
	return (&CljsCore_testC{this__757__auto__.A, this__757__auto__.B, this__757__auto__.C, G__4712, this__757__auto__.X__extmap, this__757__auto__.X__hash})
}

func (_ *CljsCore_testC) CljsCoreICollection__() {}

func (this__763__auto__ *CljsCore_testC) X_conj_Arity2(entry__764__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__764__auto__) {
		return this__763__auto__.X_assoc_Arity3(entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__763__auto__, entry__764__auto__)
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

func (this__758__auto__ *CljsCore_testA2) X_lookup_Arity2(k__759__auto__ interface{}) interface{} {
	return this__758__auto__.X_lookup_Arity3(k__759__auto__, nil)
}

func (this__760__auto__ *CljsCore_testA2) X_lookup_Arity3(k4738 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__4741 = func() interface{} {
			if cljs_core.Value_(k4738).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k4738, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4741
		switch G__4741 {
		case "x":
			return this__760__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k4738, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testA2) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__5523 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5523, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, cljs_core.Pr_writer, "", " ", "", opts__776__auto__, keyval__778__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__777__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, pr_pair__777__auto__, "#cljs.core-test.A2{", ", ", "}", opts__776__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__774__auto__.X}, nil})}, nil}), this__774__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testA2) CljsCoreIMeta__() {}

func (this__756__auto__ *CljsCore_testA2) X_meta_Arity1() interface{} {
	return this__756__auto__.X__meta
}

func (_ *CljsCore_testA2) CljsCoreICloneable__() {}

func (this__752__auto__ *CljsCore_testA2) X_clone_Arity1() interface{} {
	return (&CljsCore_testA2{this__752__auto__.X, this__752__auto__.X__meta, this__752__auto__.X__extmap, this__752__auto__.X__hash})
}

func (_ *CljsCore_testA2) CljsCoreICounted__() {}

func (this__762__auto__ *CljsCore_testA2) X_count_Arity1() float64 {
	return (float64(1) + cljs_core.Count.X_invoke_Arity1(this__762__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testA2) CljsCoreIHash__() {}

func (this__753__auto__ *CljsCore_testA2) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = this__753__auto__.X__hash
		_ = h__577__auto__
		if !(cljs_core.Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__753__auto__).(float64)
				_ = h__577__auto_____1
				this__753__auto__.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIEquiv__() {}

func (this__754__auto__ *CljsCore_testA2) X_equiv_Arity2(other__755__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__755__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__754__auto__), cljs_core.Type_.X_invoke_Arity1(other__755__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__754__auto__, other__755__auto__)))
		} else {
			return and__159__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testA2) CljsCoreIRecord__() {}

func (_ *CljsCore_testA2) CljsCoreIMap__() {}

func (this__769__auto__ *CljsCore_testA2) X_dissoc_Arity2(k__770__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), nil}, nil}, nil}), k__770__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__769__auto__), this__769__auto__.X__meta), k__770__auto__)
	} else {
		return (&CljsCore_testA2{this__769__auto__.X, this__769__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__769__auto__.X__extmap, k__770__auto__)), nil})
	}
}

func (_ *CljsCore_testA2) CljsCoreIAssociative__() {}

func (this__765__auto__ *CljsCore_testA2) X_assoc_Arity3(k__766__auto__ interface{}, G__4737 interface{}) interface{} {
	{
		var pred__4747 = cljs_core.Keyword_identical_QMARK_
		var expr__4748 = k__766__auto__
		_, _ = pred__4747, expr__4748
		if cljs_core.Truth_(func() interface{} {
			var G__4750 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__4751 = expr__4748
			_, _ = G__4750, G__4751
			return pred__4747.X_invoke_Arity2(G__4750, G__4751)
		}()) {
			return (&CljsCore_testA2{G__4737, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testA2{this__765__auto__.X, this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__4737), nil})
		}
	}
}

func (this__767__auto__ *CljsCore_testA2) X_contains_key_QMARK__Arity2(k__768__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testA2) CljsCoreISeqable__() {}

func (this__772__auto__ *CljsCore_testA2) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__772__auto__.X}, nil})}, nil}), this__772__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testA2) CljsCoreIWithMeta__() {}

func (this__757__auto__ *CljsCore_testA2) X_with_meta_Arity2(G__4737 interface{}) interface{} {
	return (&CljsCore_testA2{this__757__auto__.X, G__4737, this__757__auto__.X__extmap, this__757__auto__.X__hash})
}

func (_ *CljsCore_testA2) CljsCoreICollection__() {}

func (this__763__auto__ *CljsCore_testA2) X_conj_Arity2(entry__764__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__764__auto__) {
		return this__763__auto__.X_assoc_Arity3(entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__763__auto__, entry__764__auto__)
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

func (this__758__auto__ *CljsCore_testB) X_lookup_Arity2(k__759__auto__ interface{}) interface{} {
	return this__758__auto__.X_lookup_Arity3(k__759__auto__, nil)
}

func (this__760__auto__ *CljsCore_testB) X_lookup_Arity3(k4753 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__4756 = func() interface{} {
			if cljs_core.Value_(k4753).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k4753, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4756
		switch G__4756 {
		case "x":
			return this__760__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k4753, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testB) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testB) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__5525 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5525, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, cljs_core.Pr_writer, "", " ", "", opts__776__auto__, keyval__778__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__777__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, pr_pair__777__auto__, "#cljs.core-test.B{", ", ", "}", opts__776__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__774__auto__.X}, nil})}, nil}), this__774__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testB) CljsCoreIMeta__() {}

func (this__756__auto__ *CljsCore_testB) X_meta_Arity1() interface{} {
	return this__756__auto__.X__meta
}

func (_ *CljsCore_testB) CljsCoreICloneable__() {}

func (this__752__auto__ *CljsCore_testB) X_clone_Arity1() interface{} {
	return (&CljsCore_testB{this__752__auto__.X, this__752__auto__.X__meta, this__752__auto__.X__extmap, this__752__auto__.X__hash})
}

func (_ *CljsCore_testB) CljsCoreICounted__() {}

func (this__762__auto__ *CljsCore_testB) X_count_Arity1() float64 {
	return (float64(1) + cljs_core.Count.X_invoke_Arity1(this__762__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testB) CljsCoreIHash__() {}

func (this__753__auto__ *CljsCore_testB) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = this__753__auto__.X__hash
		_ = h__577__auto__
		if !(cljs_core.Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__753__auto__).(float64)
				_ = h__577__auto_____1
				this__753__auto__.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testB) CljsCoreIEquiv__() {}

func (this__754__auto__ *CljsCore_testB) X_equiv_Arity2(other__755__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__755__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__754__auto__), cljs_core.Type_.X_invoke_Arity1(other__755__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__754__auto__, other__755__auto__)))
		} else {
			return and__159__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testB) CljsCoreIRecord__() {}

func (_ *CljsCore_testB) CljsCoreIMap__() {}

func (this__769__auto__ *CljsCore_testB) X_dissoc_Arity2(k__770__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), nil}, nil}, nil}), k__770__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__769__auto__), this__769__auto__.X__meta), k__770__auto__)
	} else {
		return (&CljsCore_testB{this__769__auto__.X, this__769__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__769__auto__.X__extmap, k__770__auto__)), nil})
	}
}

func (_ *CljsCore_testB) CljsCoreIAssociative__() {}

func (this__765__auto__ *CljsCore_testB) X_assoc_Arity3(k__766__auto__ interface{}, G__4752 interface{}) interface{} {
	{
		var pred__4762 = cljs_core.Keyword_identical_QMARK_
		var expr__4763 = k__766__auto__
		_, _ = pred__4762, expr__4763
		if cljs_core.Truth_(func() interface{} {
			var G__4765 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__4766 = expr__4763
			_, _ = G__4765, G__4766
			return pred__4762.X_invoke_Arity2(G__4765, G__4766)
		}()) {
			return (&CljsCore_testB{G__4752, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testB{this__765__auto__.X, this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__4752), nil})
		}
	}
}

func (this__767__auto__ *CljsCore_testB) X_contains_key_QMARK__Arity2(k__768__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testB) CljsCoreISeqable__() {}

func (this__772__auto__ *CljsCore_testB) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__772__auto__.X}, nil})}, nil}), this__772__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testB) CljsCoreIWithMeta__() {}

func (this__757__auto__ *CljsCore_testB) X_with_meta_Arity2(G__4752 interface{}) interface{} {
	return (&CljsCore_testB{this__757__auto__.X, G__4752, this__757__auto__.X__extmap, this__757__auto__.X__hash})
}

func (_ *CljsCore_testB) CljsCoreICollection__() {}

func (this__763__auto__ *CljsCore_testB) X_conj_Arity2(entry__764__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__764__auto__) {
		return this__763__auto__.X_assoc_Arity3(entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__763__auto__, entry__764__auto__)
	}
}

var X__GT_B *cljs_core.AFn

var Map__GT_B *cljs_core.AFn

type CljsCore_testIFoo interface {
	CljsCore_testIFoo__()
	Foo_Arity1() interface{}
}

var Foo *cljs_core.AFn

type CljsCore_testT4767 struct {
	Test_stuff interface{}
	Meta4768   interface{}
}

func (_ *CljsCore_testT4767) CljsCore_testIFoo__() {}

func (this *CljsCore_testT4767) Foo_Arity1() interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
}

func (_ *CljsCore_testT4767) CljsCoreIMeta__() {}

func (_4769 *CljsCore_testT4767) X_meta_Arity1() interface{} {
	return _4769.Meta4768
}

func (_ *CljsCore_testT4767) CljsCoreIWithMeta__() {}

func (_4769 *CljsCore_testT4767) X_with_meta_Arity2(meta4768___1 interface{}) interface{} {
	return (&CljsCore_testT4767{_4769.Test_stuff, meta4768___1})
}

var X__GT_t4767 *cljs_core.AFn

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
		var return__5526 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})
		___.A = return__5526
		return return__5526
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

func (___ *CljsCore_testFirst) X_find_first_Arity2(p__4770 interface{}) interface{} {
	{
		var vec__4772 = p__4770
		var x = cljs_core.Nth.X_invoke_Arity3(vec__4772, float64(0), nil)
		_, _ = vec__4772, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCore_testIHasFirst__() {}

func (p__4773 *CljsCore_testFirst) X_get_first_Arity1() interface{} {
	{
		var vec__4775 = p__4773
		var x = cljs_core.Nth.X_invoke_Arity3(vec__4775, float64(0), nil)
		_, _ = vec__4775, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCoreObject__() {}

func (p__4776 *CljsCore_testFirst) ToString() string {
	{
		var vec__4778 = p__4776
		var x = cljs_core.Nth.X_invoke_Arity3(vec__4778, float64(0), nil)
		_, _ = vec__4778, x
		return (`` + cljs_core.Str.X_invoke_Arity1(x).(string))
	}
}

func (this *CljsCore_testFirst) String() string {
	return this.ToString()
}

func (_ *CljsCore_testFirst) CljsCoreIFn__() {}

func (p__4779 *CljsCore_testFirst) X_invoke_Arity0() interface{} {
	{
		var vec__4781 = p__4779
		var x = cljs_core.Nth.X_invoke_Arity3(vec__4781, float64(0), nil)
		_, _ = vec__4781, x
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

func (___ *CljsCore_testDestructuringWithLocals) X_find_first_Arity2(p__4783 interface{}) interface{} {
	{
		var vec__4785 = p__4783
		var x = cljs_core.Nth.X_invoke_Arity3(vec__4785, float64(0), nil)
		var y = cljs_core.Nth.X_invoke_Arity3(vec__4785, float64(1), nil)
		_, _, _ = vec__4785, x, y
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

func (this__758__auto__ *CljsCore_testPrintMe) X_lookup_Arity2(k__759__auto__ interface{}) interface{} {
	return this__758__auto__.X_lookup_Arity3(k__759__auto__, nil)
}

func (this__760__auto__ *CljsCore_testPrintMe) X_lookup_Arity3(k4797 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__4800 = func() interface{} {
			if cljs_core.Value_(k4797).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k4797, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__4800
		switch G__4800 {
		case "b":
			return this__760__auto__.B

		case "a":
			return this__760__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k4797, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testPrintMe) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__5528 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__5528, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, cljs_core.Pr_writer, "", " ", "", opts__776__auto__, keyval__778__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__777__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__775__auto__, pr_pair__777__auto__, "#cljs.core-test.PrintMe{", ", ", "}", opts__776__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__774__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__774__auto__.B}, nil})}, nil}), this__774__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIMeta__() {}

func (this__756__auto__ *CljsCore_testPrintMe) X_meta_Arity1() interface{} {
	return this__756__auto__.X__meta
}

func (_ *CljsCore_testPrintMe) CljsCoreICloneable__() {}

func (this__752__auto__ *CljsCore_testPrintMe) X_clone_Arity1() interface{} {
	return (&CljsCore_testPrintMe{this__752__auto__.A, this__752__auto__.B, this__752__auto__.X__meta, this__752__auto__.X__extmap, this__752__auto__.X__hash})
}

func (_ *CljsCore_testPrintMe) CljsCoreICounted__() {}

func (this__762__auto__ *CljsCore_testPrintMe) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__762__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testPrintMe) CljsCoreIHash__() {}

func (this__753__auto__ *CljsCore_testPrintMe) X_hash_Arity1() interface{} {
	{
		var h__577__auto__ = this__753__auto__.X__hash
		_ = h__577__auto__
		if !(cljs_core.Nil_(h__577__auto__)) {
			return h__577__auto__
		} else {
			{
				var h__577__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__753__auto__).(float64)
				_ = h__577__auto_____1
				this__753__auto__.X__hash = h__577__auto_____1

				return h__577__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIEquiv__() {}

func (this__754__auto__ *CljsCore_testPrintMe) X_equiv_Arity2(other__755__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__755__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__754__auto__), cljs_core.Type_.X_invoke_Arity1(other__755__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__754__auto__, other__755__auto__)))
		} else {
			return and__159__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIRecord__() {}

func (_ *CljsCore_testPrintMe) CljsCoreIMap__() {}

func (this__769__auto__ *CljsCore_testPrintMe) X_dissoc_Arity2(k__770__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__770__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__769__auto__), this__769__auto__.X__meta), k__770__auto__)
	} else {
		return (&CljsCore_testPrintMe{this__769__auto__.A, this__769__auto__.B, this__769__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__769__auto__.X__extmap, k__770__auto__)), nil})
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIAssociative__() {}

func (this__765__auto__ *CljsCore_testPrintMe) X_assoc_Arity3(k__766__auto__ interface{}, G__4796 interface{}) interface{} {
	{
		var pred__4808 = cljs_core.Keyword_identical_QMARK_
		var expr__4809 = k__766__auto__
		_, _ = pred__4808, expr__4809
		if cljs_core.Truth_(func() interface{} {
			var G__4811 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__4812 = expr__4809
			_, _ = G__4811, G__4812
			return pred__4808.X_invoke_Arity2(G__4811, G__4812)
		}()) {
			return (&CljsCore_testPrintMe{G__4796, this__765__auto__.B, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__4813 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__4814 = expr__4809
				_, _ = G__4813, G__4814
				return pred__4808.X_invoke_Arity2(G__4813, G__4814)
			}()) {
				return (&CljsCore_testPrintMe{this__765__auto__.A, G__4796, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPrintMe{this__765__auto__.A, this__765__auto__.B, this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__4796), nil})
			}
		}
	}
}

func (this__767__auto__ *CljsCore_testPrintMe) X_contains_key_QMARK__Arity2(k__768__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testPrintMe) CljsCoreISeqable__() {}

func (this__772__auto__ *CljsCore_testPrintMe) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__772__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__772__auto__.B}, nil})}, nil}), this__772__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testPrintMe) CljsCoreIWithMeta__() {}

func (this__757__auto__ *CljsCore_testPrintMe) X_with_meta_Arity2(G__4796 interface{}) interface{} {
	return (&CljsCore_testPrintMe{this__757__auto__.A, this__757__auto__.B, G__4796, this__757__auto__.X__extmap, this__757__auto__.X__hash})
}

func (_ *CljsCore_testPrintMe) CljsCoreICollection__() {}

func (this__763__auto__ *CljsCore_testPrintMe) X_conj_Arity2(entry__764__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__764__auto__) {
		return this__763__auto__.X_assoc_Arity3(entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__764__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__763__auto__, entry__764__auto__)
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

type CljsCore_testT4848 struct {
	F          interface{}
	Baz        interface{}
	Test_stuff interface{}
	Meta4849   interface{}
}

func (_ *CljsCore_testT4848) CljsCore_testIBar__() {}

func (___ *CljsCore_testT4848) X_bar_Arity2(x interface{}) interface{} {
	{
		var G__4852 = x
		_ = G__4852
		return ___.F.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4852)
	}
}

func (_ *CljsCore_testT4848) CljsCoreIMeta__() {}

func (_4850 *CljsCore_testT4848) X_meta_Arity1() interface{} {
	return _4850.Meta4849
}

func (_ *CljsCore_testT4848) CljsCoreIWithMeta__() {}

func (_4850 *CljsCore_testT4848) X_with_meta_Arity2(meta4849___1 interface{}) interface{} {
	return (&CljsCore_testT4848{_4850.F, _4850.Baz, _4850.Test_stuff, meta4849___1})
}

var X__GT_t4848 *cljs_core.AFn

var Original_closure_stmt *cljs_core.AFn

type CljsCore_testPositionalFactoryTest struct{ X interface{} }

var X__GT_PositionalFactoryTest *cljs_core.AFn

type CljsCore_testKeywordTest struct{}

func (_ *CljsCore_testKeywordTest) CljsCoreILookup__() {}

func (o *CljsCore_testKeywordTest) X_lookup_Arity2(k interface{}) interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nothing", Fqn: "nothing", X_hash: float64(-1022703296)})
}

func (o *CljsCore_testKeywordTest) X_lookup_Arity3(k interface{}, not_found interface{}) interface{} {
	return not_found
}

var X__GT_KeywordTest *cljs_core.AFn

type CljsCore_testT4861 struct {
	Test_stuff interface{}
	Meta4862   interface{}
}

func (_ *CljsCore_testT4861) CljsCoreIHash__() {}

func (___ *CljsCore_testT4861) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT4861) CljsCoreIMeta__() {}

func (_4863 *CljsCore_testT4861) X_meta_Arity1() interface{} {
	return _4863.Meta4862
}

func (_ *CljsCore_testT4861) CljsCoreIWithMeta__() {}

func (_4863 *CljsCore_testT4861) X_with_meta_Arity2(meta4862___1 interface{}) interface{} {
	return (&CljsCore_testT4861{_4863.Test_stuff, meta4862___1})
}

var X__GT_t4861 *cljs_core.AFn

type CljsCore_testT4864 struct {
	A          interface{}
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
	return (&CljsCore_testT4864{_4866.A, _4866.Test_stuff, meta4865___1})
}

var X__GT_t4864 *cljs_core.AFn

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

type CljsCore_testT4934 struct {
	From_seq    interface{}
	Make_seq    interface{}
	Mt          interface{}
	I__4926     interface{}
	Count__4925 interface{}
	Chunk__4924 interface{}
	Seq__4923   interface{}
	Test_stuff  interface{}
	Meta4935    interface{}
}

func (_ *CljsCore_testT4934) CljsCoreISeq__() {}

func (this *CljsCore_testT4934) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT4934) X_rest_Arity1() interface{} {
	{
		var G__4938 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__4938
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4938)
	}
}

func (_ *CljsCore_testT4934) CljsCoreISeqable__() {}

func (this *CljsCore_testT4934) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT4934) CljsCoreIMeta__() {}

func (_4936 *CljsCore_testT4934) X_meta_Arity1() interface{} {
	return _4936.Meta4935
}

func (_ *CljsCore_testT4934) CljsCoreIWithMeta__() {}

func (_4936 *CljsCore_testT4934) X_with_meta_Arity2(meta4935___1 interface{}) interface{} {
	return (&CljsCore_testT4934{_4936.From_seq, _4936.Make_seq, _4936.Mt, _4936.I__4926, _4936.Count__4925, _4936.Chunk__4924, _4936.Seq__4923, _4936.Test_stuff, meta4935___1})
}

var X__GT_t4934 *cljs_core.AFn

type CljsCore_testT4946 struct {
	From_seq           interface{}
	Make_seq           interface{}
	Mt                 interface{}
	Temp__4222__auto__ interface{}
	I__4926            interface{}
	Count__4925        interface{}
	Chunk__4924        interface{}
	Seq__4923          interface{}
	Test_stuff         interface{}
	Meta4947           interface{}
}

func (_ *CljsCore_testT4946) CljsCoreISeq__() {}

func (this *CljsCore_testT4946) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT4946) X_rest_Arity1() interface{} {
	{
		var G__4950 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__4950
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__4950)
	}
}

func (_ *CljsCore_testT4946) CljsCoreISeqable__() {}

func (this *CljsCore_testT4946) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT4946) CljsCoreIMeta__() {}

func (_4948 *CljsCore_testT4946) X_meta_Arity1() interface{} {
	return _4948.Meta4947
}

func (_ *CljsCore_testT4946) CljsCoreIWithMeta__() {}

func (_4948 *CljsCore_testT4946) X_with_meta_Arity2(meta4947___1 interface{}) interface{} {
	return (&CljsCore_testT4946{_4948.From_seq, _4948.Make_seq, _4948.Mt, _4948.Temp__4222__auto__, _4948.I__4926, _4948.Count__4925, _4948.Chunk__4924, _4948.Seq__4923, _4948.Test_stuff, meta4947___1})
}

var X__GT_t4946 *cljs_core.AFn

var Case_recur *cljs_core.AFn

var Xform *cljs_core.AFn

var Data interface{}

var Xf *cljs_core.AFn

func Test_runner(t *testing.T) {
	Test_stuff.X_invoke_Arity0()
}
