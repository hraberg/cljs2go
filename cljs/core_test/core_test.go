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
						i, j = (i + float64(1)), cljs_core.Conj.X_invoke_Arity2(j, func(G__941 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__941, 0, func() interface{} {
								return i
							})
						}(&cljs_core.AFn{}, i, j))
						continue
					} else {
						return cljs_core.Map_.X_invoke_Arity2(func(G__942 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__942, 1, func(p1__56_SHARP_ interface{}) interface{} {
								{
									return p1__56_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
								}
							})
						}(&cljs_core.AFn{}, i, j), j).(*cljs_core.CljsCoreLazySeq)
					}
				}
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [4 3 2 1 0] (loop [i 0 j ()] (if (< i 5) (recur (inc i) (conj j (fn [] i))) (map (fn* [p1__56#] (p1__56#)) j))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})}, nil}), cljs_core.Map_.X_invoke_Arity2(func(G__943 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__943, 1, func(p1__57_SHARP_ interface{}) interface{} {
					{
						return p1__57_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
					}
				})
			}(&cljs_core.AFn{}), func() *cljs_core.CljsCoreLazySeq {
				var iter__923__auto__ = func(iter__517 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(iter__517, 1, func(s__518 interface{}) interface{} {
						return (&cljs_core.CljsCoreLazySeq{nil, func(G__944 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__944, 0, func() interface{} {
								{
									var s__518___1 interface{} = s__518
									_ = s__518___1
									for {
										{
											var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__518___1)
											_ = temp__4222__auto__
											if cljs_core.Truth_(temp__4222__auto__) {
												{
													var xs__4752__auto__ = temp__4222__auto__
													_ = xs__4752__auto__
													{
														var i = cljs_core.First.X_invoke_Arity1(xs__4752__auto__)
														_ = i
														{
															var iterys__919__auto__ = func(iter__519 *cljs_core.AFn, s__518___1 interface{}, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																return cljs_core.Fn(iter__519, 1, func(s__520 interface{}) interface{} {
																	return (&cljs_core.CljsCoreLazySeq{nil, func(G__945 *cljs_core.AFn, s__518___1 interface{}, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																		return cljs_core.Fn(G__945, 0, func() interface{} {
																			{
																				var s__520___1 interface{} = s__520
																				_ = s__520___1
																				for {
																					{
																						var temp__4222__auto_____1 = cljs_core.Seq.Arity1IQ(s__520___1)
																						_ = temp__4222__auto_____1
																						if cljs_core.Truth_(temp__4222__auto_____1) {
																							{
																								var s__520___2 = temp__4222__auto_____1
																								_ = s__520___2
																								if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__520___2) {
																									{
																										var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__520___2)
																										var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																										var b__522 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																										_, _, _ = c__921__auto__, size__922__auto__, b__522
																										if func() bool {
																											var i__521 = float64(0)
																											_ = i__521
																											for {
																												if i__521 < size__922__auto__ {
																													{
																														var j = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__521)
																														_ = j
																														cljs_core.Chunk_append.X_invoke_Arity2(b__522, func(G__946 *cljs_core.AFn, i__521 float64, s__518___1 interface{}, j interface{}, c__921__auto__ interface{}, size__922__auto__ float64, b__522 *cljs_core.CljsCoreChunkBuffer, s__520___2 interface{}, temp__4222__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																															return cljs_core.Fn(G__946, 0, func() interface{} {
																																return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																															})
																														}(&cljs_core.AFn{}, i__521, s__518___1, j, c__921__auto__, size__922__auto__, b__522, s__520___2, temp__4222__auto_____1, i, xs__4752__auto__, temp__4222__auto__))
																														i__521 = (i__521 + float64(1))
																														continue
																													}
																												} else {
																													return true
																												}
																											}
																										}() {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__522), iter__519.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__520___2)).(*cljs_core.CljsCoreLazySeq))
																										} else {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__522), nil)
																										}
																									}
																								} else {
																									{
																										var j = cljs_core.First.X_invoke_Arity1(s__520___2)
																										_ = j
																										return cljs_core.Cons.X_invoke_Arity2(func(G__947 *cljs_core.AFn, s__518___1 interface{}, j interface{}, s__520___2 interface{}, temp__4222__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																											return cljs_core.Fn(G__947, 0, func() interface{} {
																												return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																											})
																										}(&cljs_core.AFn{}, s__518___1, j, s__520___2, temp__4222__auto_____1, i, xs__4752__auto__, temp__4222__auto__), iter__519.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__520___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
																	}(&cljs_core.AFn{}, s__518___1, i, xs__4752__auto__, temp__4222__auto__), nil, nil})
																})
															}(&cljs_core.AFn{}, s__518___1, i, xs__4752__auto__, temp__4222__auto__)
															var fs__920__auto__ = cljs_core.Seq.Arity1IQ(iterys__919__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq))
															_, _ = iterys__919__auto__, fs__920__auto__
															if cljs_core.Truth_(fs__920__auto__) {
																return cljs_core.Concat.X_invoke_Arity2(fs__920__auto__, iter__517.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__518___1)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
															} else {
																s__518___1 = cljs_core.Rest.Arity1IQ(s__518___1)
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
				panic((&js.Error{("Assert failed: (= [[1 1] [1 2] [1 3] [2 1] [2 2] [2 3]] (map (fn* [p1__57#] (p1__57#)) (for [i [1 2] j [1 2 3]] (fn [] [i j]))))")}))
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
					var _STAR_print_fn_STAR_528_948 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_528_948
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_528_948

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__949 *cljs_core.AFn, _STAR_print_fn_STAR_528_948 interface{}, sb__1124__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__949, 1, func(x__1125__auto__ interface{}) interface{} {
									return sb__1124__auto__.Append(x__1125__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_528_948, sb__1124__auto__)

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
					var _STAR_print_fn_STAR_529_950 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_529_950
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_529_950

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__951 *cljs_core.AFn, _STAR_print_fn_STAR_529_950 interface{}, sb__1124__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__951, 1, func(x__1125__auto__ interface{}) interface{} {
									return sb__1124__auto__.Append(x__1125__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_529_950, sb__1124__auto__)

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
				var jumble = func(G__952 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__952, 2, func(a interface{}, b interface{}) interface{} {
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
				var jumble = func(G__953 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__953, 2, func(a interface{}, b interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Apply.X_invoke_Arity2(func(G__954 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__954, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return ((cljs_core.Nth.X_invoke_Arity2(args, float64(0)).(float64) + cljs_core.Nth.X_invoke_Arity2(args, float64(1)).(float64)) + cljs_core.Nth.X_invoke_Arity2(args, float64(2)).(float64))
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))) {
			} else {
				panic((&js.Error{("Assert failed: (= 3 (apply (fn [& args] (+ (nth args 0) (nth args 1) (nth args 2))) (iterate inc 0)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__955 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__955, 0, func(m__ ...interface{}) interface{} {
					var m = cljs_core.Seq.Arity1IQ(m__[0])
					_ = m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [0 1 2 3 4] (take 5 (apply (fn [& m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__956 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__956, 1, func(x_m__ ...interface{}) interface{} {
					var x = x_m__[0]
					var m = cljs_core.Seq.Arity1IQ(x_m__[1])
					_, _ = x, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (take 5 (apply (fn [x & m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__957 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__957, 2, func(x_y_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5), float64(6), float64(7)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__958 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__958, 3, func(x_y_z_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6), float64(7), float64(8)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__959 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__959, 4, func(x_y_z_a_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__960 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__960, 5, func(x_y_z_a_b_m__ ...interface{}) interface{} {
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
				var single_arity_non_variadic_961 = func(G__966 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__966, 3, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{})
				var multiple_arity_non_variadic_962 = func(G__967 *cljs_core.AFn, single_arity_non_variadic_961 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__967, 3, func(x interface{}) interface{} {
						return x
					}, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{y, x}, nil})
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_961)
				var single_arity_variadic_fixedargs_963 = func(G__968 *cljs_core.AFn, single_arity_non_variadic_961 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_962 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__968, 2, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{more, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_961, multiple_arity_non_variadic_962)
				var single_arity_variadic_nofixedargs_964 = func(G__969 *cljs_core.AFn, single_arity_non_variadic_961 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_962 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_963 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__969, 0, func(more__ ...interface{}) interface{} {
						var more = cljs_core.Seq.Arity1IQ(more__[0])
						_ = more
						return more
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_961, multiple_arity_non_variadic_962, single_arity_variadic_fixedargs_963)
				var multiple_arity_variadic_965 = func(G__970 *cljs_core.AFn, single_arity_non_variadic_961 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_962 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_963 cljs_core.CljsCoreIFn, single_arity_variadic_nofixedargs_964 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__970, 2, func(x interface{}) interface{} {
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
				}(&cljs_core.AFn{}, single_arity_non_variadic_961, multiple_arity_non_variadic_962, single_arity_variadic_fixedargs_963, single_arity_variadic_nofixedargs_964)
				_, _, _, _, _ = single_arity_non_variadic_961, multiple_arity_non_variadic_962, single_arity_variadic_fixedargs_963, single_arity_variadic_nofixedargs_964, multiple_arity_variadic_965
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_non_variadic_961, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic [1 2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_non_variadic_961, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic 1 [2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_non_variadic_961, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic 1 2 [3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_962, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (apply multiple-arity-non-variadic [42]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_962, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-non-variadic [1 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_962, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-non-variadic 1 [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_962, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic [1 2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_962, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic 1 [2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_non_variadic_962, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic 1 2 [3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_963, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_fixedargs_963, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_fixedargs_963, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_fixedargs_963, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_fixedargs_963, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_963, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 4 5] (take 3 (first (apply single-arity-variadic-fixedargs (iterate inc 1)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_963, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (rest (apply single-arity-variadic-fixedargs (iterate inc 1))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_964, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_nofixedargs_964, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_nofixedargs_964, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_nofixedargs_964, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_nofixedargs_964, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_964, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (take 5 (apply single-arity-variadic-nofixedargs (iterate inc 1))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_965, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (apply multiple-arity-variadic [42]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_965, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-variadic [1 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_965, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-variadic 1 [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_965, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_965, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_variadic_965, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(multiple_arity_variadic_965, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(multiple_arity_variadic_965, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_965, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 4 5] (take 3 (first (apply multiple-arity-variadic (iterate inc 1)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_965, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (rest (apply multiple-arity-variadic (iterate inc 1))))")}))
				}
			}
			{
				var f1_971 = func(f1 *cljs_core.AFn) *cljs_core.AFn {
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
				var f2_972 = func(f2 *cljs_core.AFn, f1_971 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(f2, 2, func(x interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					}, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return cljs_core.Apply.X_invoke_Arity3(f1_971, y, more)
					})
				}(&cljs_core.AFn{}, f1_971)
				_, _ = f1_971, f2_972
				if cljs_core.X_EQ_.Arity2IIB(float64(1), f2_972.X_invoke_Arity2(float64(1), float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (f2 1 2))")}))
				}
			}
			{
				var f_973 = func(G__974 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__974, 1, func() interface{} {
						return nil
					}, func(a_more__ ...interface{}) interface{} {
						var a = a_more__[0]
						var more = cljs_core.Seq.Arity1IQ(a_more__[1])
						_, _ = a, more
						return more
					})
				}(&cljs_core.AFn{})
				_ = f_973
				if cljs_core.Nil_(f_973.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (f :foo))")}))
				}
			}
			if cljs_core.Nil_(cljs_core.Array_seq.X_invoke_Arity2([]interface{}{float64(1)}, float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (array-seq (array 1) 1))")}))
			}
			{
				var f_975 = func(G__978 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__978, 1, func(x interface{}) interface{} {
						return (x.(float64) * float64(2))
					})
				}(&cljs_core.AFn{})
				var m_976 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar"}, nil})
				var mf_977 = cljs_core.With_meta.X_invoke_Arity2(f_975, m_976)
				_, _, _ = f_975, m_976, mf_977
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(f_975)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (meta f))")}))
				}
				if cljs_core.Fn_QMARK_.Arity1IB(mf_977) {
				} else {
					panic((&js.Error{("Assert failed: (fn? mf)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__530 = float64(2)
					_ = G__530
					return mf_977.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__530)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (mf 2))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Apply.X_invoke_Arity2(mf_977, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (apply mf [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(mf_977), m_976) {
				} else {
					panic((&js.Error{("Assert failed: (= (meta mf) m)")}))
				}
			}
			{
				var a_979 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_979
				if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Deref.X_invoke_Arity1(a_979)) {
				} else {
					panic((&js.Error{("Assert failed: (= 0 (deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity2(a_979, cljs_core.Inc)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (swap! a inc))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(false, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_979, float64(0), float64(42)))) {
				} else {
					panic((&js.Error{("Assert failed: (= false (compare-and-set! a 0 42))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(true, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_979, float64(1), float64(7)))) {
				} else {
					panic((&js.Error{("Assert failed: (= true (compare-and-set! a 1 7))")}))
				}
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(a_979)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (meta a))")}))
				}
				if cljs_core.Nil_(cljs_core.Get_validator.X_invoke_Arity1(a_979)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (get-validator a))")}))
				}
			}
			{
				var a_980 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_980
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity3(a_980, cljs_core.X_PLUS_, float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (swap! a + 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Swap_BANG_.X_invoke_Arity4(a_980, cljs_core.X_PLUS_, float64(1), float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (swap! a + 1 2))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(10), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_980, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 10 (swap! a + 1 2 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(20), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_980, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(4)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 20 (swap! a + 1 2 3 4))")}))
				}
			}
			{
				var a_981 = cljs_core.Atom.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "validator", Fqn: "validator", X_hash: float64(-1966190681)}), cljs_core.Coll_QMARK_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})})).(*cljs_core.CljsCoreAtom)
				_ = a_981
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Coll_QMARK_, cljs_core.Get_validator.X_invoke_Arity1(a_981)) {
				} else {
					panic((&js.Error{("Assert failed: (= coll? (get-validator a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_981)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:a 1} (meta a))")}))
				}
				cljs_core.Alter_meta_BANG_.X_invoke_ArityVariadic(a_981, cljs_core.Assoc, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}))
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_981)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:a 1, :b 2} (meta a))")}))
				}
			}
			if cljs_core.Nil_(cljs_core.Empty.X_invoke_Arity1(nil)) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (empty nil))")}))
			}
			{
				var e_lazy_seq_982 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCoreLazySeq{nil, func(G__983 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__983, 0, func() interface{} {
						return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons)
					})
				}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_lazy_seq_982
				if cljs_core.Seq_QMARK_.Arity1IB(e_lazy_seq_982) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-lazy-seq)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_lazy_seq_982) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-lazy-seq)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_lazy_seq_982)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-lazy-seq))")}))
				}
			}
			{
				var e_list_984 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_list_984
				if cljs_core.Seq_QMARK_.Arity1IB(e_list_984) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-list)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_list_984) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-list)")}))
				}
			}
			{
				var e_elist_985 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_elist_985
				if cljs_core.Seq_QMARK_.Arity1IB(e_elist_985) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-elist)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_elist_985) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-elist)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), cljs_core.Get.X_invoke_Arity2(cljs_core.Meta.X_invoke_Arity1(e_elist_985), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= :c (get (meta e-elist) :b))")}))
				}
			}
			{
				var e_cons_986 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_cons_986
				if cljs_core.Seq_QMARK_.Arity1IB(e_cons_986) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-cons)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_cons_986) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-cons)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_cons_986)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-cons))")}))
				}
			}
			{
				var e_vec_987 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_vec_987
				if cljs_core.Vector_QMARK_.Arity1IB(e_vec_987) {
				} else {
					panic((&js.Error{("Assert failed: (vector? e-vec)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_vec_987) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-vec)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_vec_987)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-vec))")}))
				}
			}
			{
				var e_omap_988 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_omap_988
				if cljs_core.Map_QMARK_.Arity1IB(e_omap_988) {
				} else {
					panic((&js.Error{("Assert failed: (map? e-omap)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_omap_988) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-omap)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_omap_988)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-omap))")}))
				}
			}
			{
				var e_hmap_989 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_hmap_989
				if cljs_core.Map_QMARK_.Arity1IB(e_hmap_989) {
				} else {
					panic((&js.Error{("Assert failed: (map? e-hmap)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_hmap_989) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-hmap)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_hmap_989)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-hmap))")}))
				}
			}
			{
				var a_990 = cljs_core.Atom.X_invoke_Arity1(nil).(*cljs_core.CljsCoreAtom)
				_ = a_990
				if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (try 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__991 interface{}) {
					defer func() {
						if e531 := recover(); e531 != nil {
							if cljs_core.Value_(e531).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e531
									_ = e
									return__991 = float64(2)
								}
							} else {
								panic(e531)

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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__992 interface{}) {
					defer func() {
						if e532 := recover(); e532 != nil {
							if cljs_core.Value_(e532).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e532
									_ = e
									return__992 = float64(2)
								}
							} else {
								panic(e532)

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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__993 interface{}) {
					defer func() {
						if e533 := recover(); e533 != nil {
							if cljs_core.Value_(e533).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e533
									_ = e
									return__993 = float64(2)
								}
							} else {
								{
									var e = e533
									_ = e
									return__993 = float64(3)
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
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__994 interface{}) {
					defer func() {
						if e534 := recover(); e534 != nil {
							if cljs_core.Value_(e534).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e534
									_ = e
									return__994 = float64(2)
								}
							} else {
								{
									var e = e534
									_ = e
									return__994 = float64(3)
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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__995 interface{}) {
					defer func() {
						if e535 := recover(); e535 != nil {
							if cljs_core.Value_(e535).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e535
									_ = e
									return__995 = float64(3)
								}
							} else {
								{
									var e = e535
									_ = e
									return__995 = e
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
						cljs_core.Reset_BANG_.X_invoke_Arity2(a_990, float64(42))
					}()
					{
						return float64(1)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (try 1 (finally (reset! a 42))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Deref.X_invoke_Arity1(a_990)) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (deref a))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), cljs_core.Seq_(cljs_core.Nthnext.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)))) {
			} else {
				panic((&js.Error{("Assert failed: (= [3] (nthnext [1 2 3] 2))")}))
			}
			{
				var v_996 = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})
				_ = v_996
				if cljs_core.X_EQ_.Arity2IIB(v_996, func() *cljs_core.CljsCoreLazySeq {
					var iter__923__auto__ = func(iter__536 *cljs_core.AFn, v_996 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__536, 1, func(s__537 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__997 *cljs_core.AFn, v_996 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__997, 0, func() interface{} {
									{
										var s__537___1 interface{} = s__537
										_ = s__537___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__537___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__537___2 = temp__4222__auto__
														_ = s__537___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__537___2) {
															{
																var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__537___2)
																var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																var b__539 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__921__auto__, size__922__auto__, b__539
																if func() bool {
																	var i__538 = float64(0)
																	_ = i__538
																	for {
																		if i__538 < size__922__auto__ {
																			{
																				var e = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__538)
																				_ = e
																				cljs_core.Chunk_append.X_invoke_Arity2(b__539, e)
																				i__538 = (i__538 + float64(1))
																				continue
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__539), iter__536.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__537___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__539), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__537___2)
																_ = e
																return cljs_core.Cons.X_invoke_Arity2(e, iter__536.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__537___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_996), nil, nil})
						})
					}(&cljs_core.AFn{}, v_996)
					_ = iter__923__auto__
					return iter__923__auto__.X_invoke_Arity1(v_996).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= v (for [e v] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(9)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__923__auto__ = func(iter__542 *cljs_core.AFn, v_996 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__542, 1, func(s__543 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__998 *cljs_core.AFn, v_996 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__998, 0, func() interface{} {
									{
										var s__543___1 interface{} = s__543
										_ = s__543___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__543___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__543___2 = temp__4222__auto__
														_ = s__543___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__543___2) {
															{
																var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__543___2)
																var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																var b__545 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__921__auto__, size__922__auto__, b__545
																if func() bool {
																	var i__544 = float64(0)
																	_ = i__544
																	for {
																		if i__544 < size__922__auto__ {
																			{
																				var e = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__544)
																				_ = e
																				{
																					var m = (e.(float64) * e.(float64))
																					_ = m
																					cljs_core.Chunk_append.X_invoke_Arity2(b__545, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																					i__544 = (i__544 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__545), iter__542.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__543___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__545), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__543___2)
																_ = e
																{
																	var m = (e.(float64) * e.(float64))
																	_ = m
																	return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__542.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__543___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_996), nil, nil})
						})
					}(&cljs_core.AFn{}, v_996)
					_ = iter__923__auto__
					return iter__923__auto__.X_invoke_Arity1(v_996).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [[1 1] [2 4] [3 9]] (for [e v :let [m (* e e)]] [e m]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__923__auto__ = func(iter__548 *cljs_core.AFn, v_996 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__548, 1, func(s__549 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__999 *cljs_core.AFn, v_996 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__999, 0, func() interface{} {
									{
										var s__549___1 interface{} = s__549
										_ = s__549___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__549___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__549___2 = temp__4222__auto__
														_ = s__549___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__549___2) {
															{
																var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__549___2)
																var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																var b__551 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__921__auto__, size__922__auto__, b__551
																if cljs_core.Truth_(func() interface{} {
																	var i__550 = float64(0)
																	_ = i__550
																	for {
																		if i__550 < size__922__auto__ {
																			{
																				var e = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__550)
																				_ = e
																				if e.(float64) < float64(3) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__551, e)
																					i__550 = (i__550 + float64(1))
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
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__551), iter__548.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__549___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__551), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__549___2)
																_ = e
																if e.(float64) < float64(3) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__548.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__549___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_996), nil, nil})
						})
					}(&cljs_core.AFn{}, v_996)
					_ = iter__923__auto__
					return iter__923__auto__.X_invoke_Arity1(v_996).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2] (for [e v :while (< e 3)] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__923__auto__ = func(iter__554 *cljs_core.AFn, v_996 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__554, 1, func(s__555 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1000 *cljs_core.AFn, v_996 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1000, 0, func() interface{} {
									{
										var s__555___1 interface{} = s__555
										_ = s__555___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__555___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__555___2 = temp__4222__auto__
														_ = s__555___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__555___2) {
															{
																var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__555___2)
																var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																var b__557 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__921__auto__, size__922__auto__, b__557
																if func() bool {
																	var i__556 = float64(0)
																	_ = i__556
																	for {
																		if i__556 < size__922__auto__ {
																			{
																				var e = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__556)
																				_ = e
																				if e.(float64) > float64(2) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__557, e)
																					i__556 = (i__556 + float64(1))
																					continue
																				} else {
																					i__556 = (i__556 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__557), iter__554.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__555___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__557), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__555___2)
																_ = e
																if e.(float64) > float64(2) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__554.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__555___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
																} else {
																	s__555___1 = cljs_core.Rest.Arity1IQ(s__555___2)
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
							}(&cljs_core.AFn{}, v_996), nil, nil})
						})
					}(&cljs_core.AFn{}, v_996)
					_ = iter__923__auto__
					return iter__923__auto__.X_invoke_Arity1(v_996).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [3] (for [e v :when (> e 2)] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__923__auto__ = func(iter__560 *cljs_core.AFn, v_996 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__560, 1, func(s__561 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1001 *cljs_core.AFn, v_996 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1001, 0, func() interface{} {
									{
										var s__561___1 interface{} = s__561
										_ = s__561___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__561___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__561___2 = temp__4222__auto__
														_ = s__561___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__561___2) {
															{
																var c__921__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__561___2)
																var size__922__auto__ = cljs_core.Count.X_invoke_Arity1(c__921__auto__).(float64)
																var b__563 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__922__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__921__auto__, size__922__auto__, b__563
																if cljs_core.Truth_(func() interface{} {
																	var i__562 = float64(0)
																	_ = i__562
																	for {
																		if i__562 < size__922__auto__ {
																			{
																				var e = c__921__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__562)
																				_ = e
																				if e.(float64) < float64(3) {
																					{
																						var m = (e.(float64) * e.(float64))
																						_ = m
																						cljs_core.Chunk_append.X_invoke_Arity2(b__563, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																						i__562 = (i__562 + float64(1))
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
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__563), iter__560.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__561___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__563), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__561___2)
																_ = e
																if e.(float64) < float64(3) {
																	{
																		var m = (e.(float64) * e.(float64))
																		_ = m
																		return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__560.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__561___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_996), nil, nil})
						})
					}(&cljs_core.AFn{}, v_996)
					_ = iter__923__auto__
					return iter__923__auto__.X_invoke_Arity1(v_996).(*cljs_core.CljsCoreLazySeq)
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
				var a10_1002 = cljs_core.Partial.X_invoke_Arity2(cljs_core.X_PLUS_, float64(10)).(cljs_core.CljsCoreIFn)
				var a20_1003 = cljs_core.Partial.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), float64(10)).(cljs_core.CljsCoreIFn)
				var a21_1004 = cljs_core.Partial.X_invoke_Arity4(cljs_core.X_PLUS_, float64(10), float64(10), float64(1)).(cljs_core.CljsCoreIFn)
				var a22_1005 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)})).(cljs_core.CljsCoreIFn)
				var a23_1006 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(1)})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = a10_1002, a20_1003, a21_1004, a22_1005, a23_1006
				if cljs_core.X_EQ_.Arity2IIB(float64(110), func() interface{} {
					var G__566 = float64(100)
					_ = G__566
					return a10_1002.X_invoke_Arity1(G__566)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 110 (a10 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(120), func() interface{} {
					var G__567 = float64(100)
					_ = G__567
					return a20_1003.X_invoke_Arity1(G__567)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 120 (a20 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(121), func() interface{} {
					var G__568 = float64(100)
					_ = G__568
					return a21_1004.X_invoke_Arity1(G__568)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 121 (a21 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(122), func() interface{} {
					var G__569 = float64(100)
					_ = G__569
					return a22_1005.X_invoke_Arity1(G__569)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 122 (a22 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(123), func() interface{} {
					var G__570 = float64(100)
					_ = G__570
					return a23_1006.X_invoke_Arity1(G__570)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 123 (a23 100))")}))
				}
			}
			{
				var n2_1007 = cljs_core.Comp.X_invoke_Arity2(cljs_core.First, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n3_1008 = cljs_core.Comp.X_invoke_Arity3(cljs_core.First, cljs_core.Rest, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n4_1009 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n5_1010 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n6_1011 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = n2_1007, n3_1008, n4_1009, n5_1010, n6_1011
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() interface{} {
					var G__571 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__571
					return n2_1007.X_invoke_Arity1(G__571)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (n2 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() interface{} {
					var G__572 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__572
					return n3_1008.X_invoke_Arity1(G__572)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (n3 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__573 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__573
					return n4_1009.X_invoke_Arity1(G__573)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (n4 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(5), func() interface{} {
					var G__574 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__574
					return n5_1010.X_invoke_Arity1(G__574)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 5 (n5 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), func() interface{} {
					var G__575 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__575
					return n6_1011.X_invoke_Arity1(G__575)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (n6 [1 2 3 4 5 6 7]))")}))
				}
			}
			{
				var sf_1012 = cljs_core.Some_fn.X_invoke_Arity3(cljs_core.Number_QMARK_, cljs_core.Keyword_QMARK_, cljs_core.Symbol_QMARK_).(cljs_core.CljsCoreIFn)
				_ = sf_1012
				if cljs_core.Truth_(func() interface{} {
					var G__576 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					var G__577 = float64(1)
					_, _ = G__576, G__577
					return sf_1012.X_invoke_Arity2(G__576, G__577)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf :foo 1)")}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__578 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					_ = G__578
					return sf_1012.X_invoke_Arity1(G__578)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf :foo)")}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__579 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})
					var G__580 = float64(1)
					_, _ = G__579, G__580
					return sf_1012.X_invoke_Arity2(G__579, G__580)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf (quote bar) 1)")}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__581 = cljs_core.CljsCorePersistentVector_EMPTY
					var G__582 = cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)
					_, _ = G__581, G__582
					return sf_1012.X_invoke_Arity2(G__581, G__582)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (not (sf [] ()))")}))
				}
			}
			{
				var ep_1013 = cljs_core.Every_pred.X_invoke_Arity2(cljs_core.Number_QMARK_, cljs_core.Zero_QMARK_).(cljs_core.CljsCoreIFn)
				_ = ep_1013
				if cljs_core.Truth_(func() interface{} {
					var G__583 = float64(0)
					var G__584 = float64(0)
					var G__585 = float64(0)
					_, _, _ = G__583, G__584, G__585
					return ep_1013.X_invoke_Arity3(G__583, G__584, G__585)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (ep 0 0 0)")}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__586 = float64(1)
					var G__587 = float64(2)
					var G__588 = float64(3)
					var G__589 = float64(0)
					_, _, _, _ = G__586, G__587, G__588, G__589
					return ep_1013.X_invoke_Arity4(G__586, G__587, G__588, G__589)
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
				var vec__590 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__590, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__590, float64(1), nil)
				_, _, _ = vec__590, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] [1 2]] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), func() cljs_core.CljsCoreISet {
				var vec__591 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__591, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__591, float64(1), nil)
				_, _, _ = vec__591, a, b
				return cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{a, b}, true).(*cljs_core.CljsCorePersistentHashSet)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 2} (let [[a b] [1 2]] #{a b}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__592 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__592___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__592) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__592)
					} else {
						return map__592
					}
				}()
				var a = cljs_core.Get.X_invoke_Arity2(map__592___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				var b = cljs_core.Get.X_invoke_Arity2(map__592___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				_, _, _, _ = map__592, map__592___1, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (let [{a :a, b :b} {:a 1, :b 2}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__593 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__593___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__593) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__593)
					} else {
						return map__593
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__593___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__593___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__593, map__593___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (let [{:keys [a b]} {:a 1, :b 2}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__594 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__594, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__594, float64(1), nil)
				var v = vec__594
				_, _, _, _ = vec__594, a, b, v
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b, v}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 [1 2]] (let [[a b :as v] [1 2]] [a b v]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(42)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__595 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__595___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__595) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__595)
					} else {
						return map__595
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity3(map__595___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(42))
				var a = cljs_core.Get.X_invoke_Arity2(map__595___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__595, map__595___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 42] (let [{:keys [a b], :or {b 42}} {:a 1}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), nil}, nil}), func() cljs_core.CljsCoreIVector {
				var map__596 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__596___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__596) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__596)
					} else {
						return map__596
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__596___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__596___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__596, map__596___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 nil] (let [{:keys [a b], :or {c 42}} {:a 1}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__597 = cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)
				var a = cljs_core.Nth.X_invoke_Arity3(vec__597, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__597, float64(1), nil)
				_, _, _ = vec__597, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] (quote (1 2))] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIMap {
				var vec__598 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__598, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__598, float64(1), nil)
				_, _, _ = vec__598, a, b
				return cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{a, b}, true, false).(*cljs_core.CljsCorePersistentArrayMap)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= {1 2} (let [[a b] [1 2]] {a b}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__599 = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))
				var a = cljs_core.Nth.X_invoke_Arity3(vec__599, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__599, float64(1), nil)
				_, _, _ = vec__599, a, b
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
				var a_1014 = cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{})
				_ = a_1014
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20), float64(30)}, nil}), cljs_core.Seq.Arity1IQ(func() []interface{} {
					var a__1042__auto__ = a_1014
					var ret = cljs_core.Aclone.X_invoke_Arity1(a__1042__auto__).([]interface{})
					_, _ = a__1042__auto__, ret
					{
						var i = float64(0)
						_ = i
						for {
							if i < float64(len(a__1042__auto__)) {
								ret[int(i)] = (float64(10) * (a_1014[int(i)]).(float64))
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
					var a__1048__auto__ = a_1014
					_ = a__1048__auto__
					{
						var i = float64(0)
						var ret = float64(0)
						_, _ = i, ret
						for {
							if i < float64(len(a__1048__auto__)) {
								i, ret = (i + float64(1)), (ret + (a_1014[int(i)]).(float64))
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_1014), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (seq a) (seq (to-array [1 2 3])))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), func() float64 { a_1014[int(float64(0))] = float64(42); return a_1014[int(float64(0))].(float64) }()) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (aset a 0 42))")}))
				}
				if cljs_core.Not_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_1014), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
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
				var coll_1015 = (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})
				var shuffles_1016 = cljs_core.Filter.X_invoke_Arity2(func(G__1017 *cljs_core.AFn, coll_1015 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1017, 1, func(p1__58_SHARP_ interface{}) interface{} {
						return cljs_core.Not_EQ_.Arity2IIB(coll_1015, p1__58_SHARP_)
					})
				}(&cljs_core.AFn{}, coll_1015), cljs_core.Take.X_invoke_Arity2(float64(100), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Shuffle, coll_1015).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
				_, _ = coll_1015, shuffles_1016
				if !(cljs_core.Empty_QMARK_.Arity1IB(shuffles_1016)) {
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
				var s_1018 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = s_1018
				{
					var n__1054__auto___1019 = float64(5)
					_ = n__1054__auto___1019
					{
						var n_1020 = float64(0)
						_ = n_1020
						for {
							if n_1020 < n__1054__auto___1019 {
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_1018, cljs_core.Conj, n_1020)
								n_1020 = (n_1020 + float64(1))
								continue
							} else {
							}
							break
						}
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Deref.X_invoke_Arity1(s_1018)) {
				} else {
					panic((&js.Error{("Assert failed: (= [0 1 2 3 4] (clojure.core/deref s))")}))
				}
			}
			{
				var v_1021 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})
				var s_1022 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)).(*cljs_core.CljsCoreAtom)
				_, _ = v_1021, s_1022
				{
					var seq__600_1023 interface{} = cljs_core.Seq.Arity1IQ(v_1021)
					var chunk__601_1024 interface{} = nil
					var count__602_1025 = float64(0)
					var i__603_1026 = float64(0)
					_, _, _, _ = seq__600_1023, chunk__601_1024, count__602_1025, i__603_1026
					for {
						if i__603_1026 < count__602_1025 {
							{
								var n_1027 = chunk__601_1024.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__603_1026)
								_ = n_1027
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_1022, cljs_core.Conj, n_1027)
								seq__600_1023, chunk__601_1024, count__602_1025, i__603_1026 = seq__600_1023, chunk__601_1024, count__602_1025, (i__603_1026 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___1028 = cljs_core.Seq.Arity1IQ(seq__600_1023)
								_ = temp__4222__auto___1028
								if cljs_core.Truth_(temp__4222__auto___1028) {
									{
										var seq__600_1029___1 = temp__4222__auto___1028
										_ = seq__600_1029___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__600_1029___1) {
											{
												var c__954__auto___1030 = cljs_core.Chunk_first.X_invoke_Arity1(seq__600_1029___1)
												_ = c__954__auto___1030
												seq__600_1023, chunk__601_1024, count__602_1025, i__603_1026 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__600_1029___1), c__954__auto___1030, cljs_core.Count.X_invoke_Arity1(c__954__auto___1030).(float64), float64(0)
												continue
											}
										} else {
											{
												var n_1031 = cljs_core.First.X_invoke_Arity1(seq__600_1029___1)
												_ = n_1031
												cljs_core.Swap_BANG_.X_invoke_Arity3(s_1022, cljs_core.Conj, n_1031)
												seq__600_1023, chunk__601_1024, count__602_1025, i__603_1026 = cljs_core.Next.Arity1IQ(seq__600_1029___1), nil, float64(0), float64(0)
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Deref.X_invoke_Arity1(s_1022), cljs_core.Reverse.X_invoke_Arity1(v_1021)) {
				} else {
					panic((&js.Error{("Assert failed: (= (clojure.core/deref s) (reverse v))")}))
				}
			}
			{
				var a_1032 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				var d_1033 = (&cljs_core.CljsCoreDelay{func(G__1034 *cljs_core.AFn, a_1032 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__1034, 0, func() interface{} {
						return cljs_core.Swap_BANG_.X_invoke_Arity2(a_1032, cljs_core.Inc)
					})
				}(&cljs_core.AFn{}, a_1032), nil})
				_, _ = a_1032, d_1033
				if cljs_core.Realized_QMARK_.Arity1IB(d_1033) == false {
				} else {
					panic((&js.Error{("Assert failed: (false? (realized? d))")}))
				}
				if cljs_core.Deref.X_invoke_Arity1(a_1032).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (clojure.core/deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_1033)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_1032)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref a))")}))
				}
				if cljs_core.Realized_QMARK_.Arity1IB(d_1033) == true {
				} else {
					panic((&js.Error{("Assert failed: (true? (realized? d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_1033)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_1032)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Force.X_invoke_Arity1(d_1033), cljs_core.Deref.X_invoke_Arity1(d_1033)) {
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
				var f_1035 = cljs_core.Memoize.X_invoke_Arity1(func(G__1036 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1036, 0, func() interface{} {
						return cljs_core.Rand.Arity0F()
					})
				}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn)
				_ = f_1035
				{
					f_1035.X_invoke_Arity0()
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return f_1035.X_invoke_Arity0()
				}(), func() interface{} {
					return f_1035.X_invoke_Arity0()
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
				var d_1037 = cljs_core.Group_by.X_invoke_Arity2(cljs_core.Second, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(6), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), float64(4), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), float64(2)}, nil}))
				_ = d_1037
				if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_1037, float64(1))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (count (get d 1)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_1037, float64(2))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (count (get d 2)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_1037, float64(4))).(float64)) {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__1038 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1038, 1, func(p1__59_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__59_SHARP_) {
						return p1__59_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 3 5 7 9] (keep (fn* [p1__59#] (if (odd? p1__59#) p1__59#)) [1 2 3 4 5 6 7 8 9 10]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(6), float64(8), float64(10)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__1039 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1039, 1, func(p1__60_SHARP_ interface{}) interface{} {
					if cljs_core.Even_QMARK_.Arity1IB(p1__60_SHARP_) {
						return p1__60_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 4 6 8 10] (keep (fn* [p1__60#] (if (even? p1__60#) p1__60#)) [1 2 3 4 5 6 7 8 9 10]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__1040 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1040, 2, func(p1__61_SHARP_ interface{}, p2__62_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__61_SHARP_) {
						return p2__62_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(11), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 3 5 7 9] (keep-indexed (fn* [p1__61# p2__62#] (if (odd? p1__61#) p2__62#)) [0 1 2 3 4 5 6 7 8 9 10]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(5)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__1041 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1041, 2, func(p1__64_SHARP_ interface{}, p2__63_SHARP_ interface{}) interface{} {
					if p2__63_SHARP_.(float64) > float64(0) {
						return p1__64_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-9), float64(0), float64(29), float64(-7), float64(45), float64(3), float64(-8)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 4 5] (keep-indexed (fn* [p1__64# p2__63#] (if (pos? p2__63#) p1__64#)) [-9 0 29 -7 45 3 -8]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil}), cljs_core.Map_indexed.X_invoke_Arity2(func(G__1042 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1042, 2, func(p1__65_SHARP_ interface{}, p2__66_SHARP_ interface{}) interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__65_SHARP_, p2__66_SHARP_}, nil})
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [[0 :a] [1 :b] [2 :c]] (map-indexed (fn* [p1__65# p2__66#] (vector p1__65# p2__66#)) [:a :b :c]))")}))
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
						return func(G__1043 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1043, 0, func() interface{} {
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
			Bar = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1064__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1065__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1066__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1067__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1068__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__
				return (&cljs_core.CljsCoreMultiFn{"bar", func(G__1044 *cljs_core.AFn, method_table__1064__auto__ *cljs_core.CljsCoreAtom, prefer_table__1065__auto__ *cljs_core.CljsCoreAtom, method_cache__1066__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1067__auto__ *cljs_core.CljsCoreAtom, hierarchy__1068__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1044, 2, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{x, y}, nil})
					})
				}(&cljs_core.AFn{}, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}), func(G__1045 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1045, 2, func(x interface{}, y interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)})
				})
			}(&cljs_core.AFn{}))
			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})}, nil}), func(G__1046 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1046, 2, func(x interface{}, y interface{}) interface{} {
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
				var G__605 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				var G__606 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				_, _ = G__605, G__606
				return Bar.X_invoke_Arity2(G__605, G__606)
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
				return (&cljs_core.CljsCoreMultiFn{"nested-dispatch", func(G__1047 *cljs_core.AFn, method_table__1064__auto__ *cljs_core.CljsCoreAtom, prefer_table__1065__auto__ *cljs_core.CljsCoreAtom, method_cache__1066__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1067__auto__ *cljs_core.CljsCoreAtom, hierarchy__1068__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1047, 1, func(m interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(m))
					})
				}(&cljs_core.AFn{}, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), func(G__1048 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1048, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1049 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1049, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__607 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil})
				_ = G__607
				return Nested_dispatch.X_invoke_Arity1(G__607)
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

			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__1050 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1050, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1051 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1051, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__608 = (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})}, nil})
				_ = G__608
				return Nested_dispatch2.X_invoke_Arity1(G__608)
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
				return (&cljs_core.CljsCoreMultiFn{"foo1", func(G__1052 *cljs_core.AFn, method_table__1064__auto__ *cljs_core.CljsCoreAtom, prefer_table__1065__auto__ *cljs_core.CljsCoreAtom, method_cache__1066__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1067__auto__ *cljs_core.CljsCoreAtom, hierarchy__1068__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1052, 0, func(args__ ...interface{}) interface{} {
						var args = cljs_core.Seq.Arity1IQ(args__[0])
						_ = args
						return cljs_core.First.X_invoke_Arity1(args)
					})
				}(&cljs_core.AFn{}, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__1053 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1053, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)})
				})
			}(&cljs_core.AFn{}))
			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1054 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1054, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)}), func() interface{} {
				var G__609 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
				_ = G__609
				return Foo1.X_invoke_Arity1(G__609)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :a-return (foo1 :a))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)}), func() interface{} {
				var G__610 = float64(1)
				_ = G__610
				return Foo1.X_invoke_Arity1(G__610)
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

			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Rect", Fqn: "Rect", X_hash: float64(-420704620)}), func(G__1055 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1055, 1, func(r interface{}) interface{} {
					return ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wd", Fqn: "wd", X_hash: float64(-183204751)}).X_invoke_Arity1(r).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ht", Fqn: "ht", X_hash: float64(214115472)}).X_invoke_Arity1(r).(float64))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Circle", Fqn: "Circle", X_hash: float64(-158896930)}), func(G__1056 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1056, 1, func(c interface{}) interface{} {
					return (Math.PI * ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64)))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1057 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1057, 1, func(x interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)})
				})
			}(&cljs_core.AFn{}))
			R = Rect.X_invoke_Arity2(float64(4), float64(13)).(cljs_core.CljsCoreIMap)

			C = Circle.X_invoke_Arity1(float64(12)).(cljs_core.CljsCoreIMap)

			if cljs_core.X_EQ_.Arity2IIB(float64(52), func() interface{} {
				var G__611 = R
				_ = G__611
				return Area.X_invoke_Arity1(G__611)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 52 (area r))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)}), func() interface{} {
				var G__612 = cljs_core.CljsCorePersistentArrayMap_EMPTY
				_ = G__612
				return Area.X_invoke_Arity1(G__612)
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
				return (&cljs_core.CljsCoreMultiFn{"apply-multi-test", func(G__1058 *cljs_core.AFn, method_table__1064__auto__ *cljs_core.CljsCoreAtom, prefer_table__1065__auto__ *cljs_core.CljsCoreAtom, method_cache__1066__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1067__auto__ *cljs_core.CljsCoreAtom, hierarchy__1068__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1058, 3, func(___ interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
						return float64(0)
					})
				}(&cljs_core.AFn{}, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__, hierarchy__1068__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1068__auto__, method_table__1064__auto__, prefer_table__1065__auto__, method_cache__1066__auto__, cached_hierarchy__1067__auto__})
			}()

			Apply_multi_test.X_add_method_Arity3(float64(0), func(G__1059 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1059, 2, func(x interface{}) interface{} {
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

			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)}), func(G__1060 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1060, 1, func(___ interface{}) interface{} {
					return true
				})
			}(&cljs_core.AFn{}))
			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1061 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1061, 1, func(___ interface{}) interface{} {
					return false
				})
			}(&cljs_core.AFn{}))
			{
				var seq__613_1062 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__614_1063 interface{} = nil
				var count__615_1064 = float64(0)
				var i__616_1065 = float64(0)
				_, _, _, _ = seq__613_1062, chunk__614_1063, count__615_1064, i__616_1065
				for {
					if i__616_1065 < count__615_1064 {
						{
							var m_1066 = chunk__614_1063.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__616_1065)
							_ = m_1066
							if cljs_core.Truth_(func() interface{} {
								var G__617 = m_1066
								_ = G__617
								return My_map_QMARK_.X_invoke_Arity1(G__617)
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (my-map? m)")}))
							}
							seq__613_1062, chunk__614_1063, count__615_1064, i__616_1065 = seq__613_1062, chunk__614_1063, count__615_1064, (i__616_1065 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1067 = cljs_core.Seq.Arity1IQ(seq__613_1062)
							_ = temp__4222__auto___1067
							if cljs_core.Truth_(temp__4222__auto___1067) {
								{
									var seq__613_1068___1 = temp__4222__auto___1067
									_ = seq__613_1068___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__613_1068___1) {
										{
											var c__954__auto___1069 = cljs_core.Chunk_first.X_invoke_Arity1(seq__613_1068___1)
											_ = c__954__auto___1069
											seq__613_1062, chunk__614_1063, count__615_1064, i__616_1065 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__613_1068___1), c__954__auto___1069, cljs_core.Count.X_invoke_Arity1(c__954__auto___1069).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_1070 = cljs_core.First.X_invoke_Arity1(seq__613_1068___1)
											_ = m_1070
											if cljs_core.Truth_(func() interface{} {
												var G__618 = m_1070
												_ = G__618
												return My_map_QMARK_.X_invoke_Arity1(G__618)
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (my-map? m)")}))
											}
											seq__613_1062, chunk__614_1063, count__615_1064, i__616_1065 = cljs_core.Next.Arity1IQ(seq__613_1068___1), nil, float64(0), float64(0)
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
				var seq__619_1071 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentVector_EMPTY, float64(1), "asdf", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}))
				var chunk__620_1072 interface{} = nil
				var count__621_1073 = float64(0)
				var i__622_1074 = float64(0)
				_, _, _, _ = seq__619_1071, chunk__620_1072, count__621_1073, i__622_1074
				for {
					if i__622_1074 < count__621_1073 {
						{
							var not_m_1075 = chunk__620_1072.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__622_1074)
							_ = not_m_1075
							if cljs_core.Not.Arity1IB(func() interface{} {
								var G__623 = not_m_1075
								_ = G__623
								return My_map_QMARK_.X_invoke_Arity1(G__623)
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (not (my-map? not-m))")}))
							}
							seq__619_1071, chunk__620_1072, count__621_1073, i__622_1074 = seq__619_1071, chunk__620_1072, count__621_1073, (i__622_1074 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1076 = cljs_core.Seq.Arity1IQ(seq__619_1071)
							_ = temp__4222__auto___1076
							if cljs_core.Truth_(temp__4222__auto___1076) {
								{
									var seq__619_1077___1 = temp__4222__auto___1076
									_ = seq__619_1077___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__619_1077___1) {
										{
											var c__954__auto___1078 = cljs_core.Chunk_first.X_invoke_Arity1(seq__619_1077___1)
											_ = c__954__auto___1078
											seq__619_1071, chunk__620_1072, count__621_1073, i__622_1074 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__619_1077___1), c__954__auto___1078, cljs_core.Count.X_invoke_Arity1(c__954__auto___1078).(float64), float64(0)
											continue
										}
									} else {
										{
											var not_m_1079 = cljs_core.First.X_invoke_Arity1(seq__619_1077___1)
											_ = not_m_1079
											if cljs_core.Not.Arity1IB(func() interface{} {
												var G__624 = not_m_1079
												_ = G__624
												return My_map_QMARK_.X_invoke_Arity1(G__624)
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (not (my-map? not-m))")}))
											}
											seq__619_1071, chunk__620_1072, count__621_1073, i__622_1074 = cljs_core.Next.Arity1IQ(seq__619_1077___1), nil, float64(0), float64(0)
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
				var pv_1080 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange))
				_ = pv_1080
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity2(pv_1080, float64(96)), float64(96)) {
				} else {
					panic((&js.Error{("Assert failed: (= (nth pv 96) 96)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity3(pv_1080, float64(97), nil), nil) {
				} else {
					panic((&js.Error{("Assert failed: (= (nth pv 97 nil) nil)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__625 = float64(96)
					_ = G__625
					return pv_1080.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__625)
				}(), float64(96)) {
				} else {
					panic((&js.Error{("Assert failed: (= (pv 96) 96)")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(cljs_core.CljsCorePersistentVector_EMPTY)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq []))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reverse.X_invoke_Arity1(pv_1080), cljs_core.Rseq.Arity1IQ(pv_1080)) {
				} else {
					panic((&js.Error{("Assert failed: (= (reverse pv) (rseq pv))")}))
				}
			}
			{
				var pv_1081 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(33)).(*cljs_core.CljsCoreRange))
				_ = pv_1081
				if cljs_core.X_EQ_.Arity2IIB(pv_1081, cljs_core.Conj.X_invoke_Arity2(cljs_core.Conj.X_invoke_Arity2(cljs_core.Pop.X_invoke_Arity1(cljs_core.Pop.X_invoke_Arity1(pv_1081)), float64(31)), float64(32))) {
				} else {
					panic((&js.Error{("Assert failed: (= pv (-> pv pop pop (conj 31) (conj 32)))")}))
				}
			}
			{
				var stack1_1082 = cljs_core.Pop.X_invoke_Arity1(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange)))
				var stack2_1083 = cljs_core.Pop.X_invoke_Arity1(stack1_1082)
				_, _ = stack1_1082, stack2_1083
				if cljs_core.X_EQ_.Arity2IIB(float64(95), cljs_core.Peek.X_invoke_Arity1(stack1_1082)) {
				} else {
					panic((&js.Error{("Assert failed: (= 95 (peek stack1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(94), cljs_core.Peek.X_invoke_Arity1(stack2_1083)) {
				} else {
					panic((&js.Error{("Assert failed: (= 94 (peek stack2))")}))
				}
			}
			{
				var sentinel_1084 = cljs_core.Rand.Arity0F()
				_ = sentinel_1084
				if reflect.DeepEqual(sentinel_1084, func() (return__1085 interface{}) {
					defer func() {
						if e626 := recover(); e626 != nil {
							if cljs_core.Value_(e626).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var ___ = e626
									_ = ___
									return__1085 = sentinel_1084
								}
							} else {
								panic(e626)

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
				var v1_1086 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				var v2_1087 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(5)).(*cljs_core.CljsCoreRange))
				var s_1088 = cljs_core.Subvec.X_invoke_Arity3(v1_1086, float64(2), float64(8)).(*cljs_core.CljsCoreSubvec)
				_, _, _ = v1_1086, v2_1087, s_1088
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(s_1088, cljs_core.Subvec.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2(v1_1086, float64(2)).(*cljs_core.CljsCoreSubvec), float64(0), float64(6)).(*cljs_core.CljsCoreSubvec), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Take.X_invoke_Arity2(float64(6), cljs_core.Drop.X_invoke_Arity2(float64(2), v1_1086).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= s (-> v1 (subvec 2) (subvec 0 6)) (->> v1 (drop 2) (take 6)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Count.X_invoke_Arity1(s_1088).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (count s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Pop.X_invoke_Arity1(s_1088)) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 4 5 6] (pop s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(7), cljs_core.Peek.X_invoke_Arity1(s_1088)) {
				} else {
					panic((&js.Error{("Assert failed: (= 7 (peek s))")}))
				}
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(1)}, nil}), cljs_core.Assoc.X_invoke_Arity3(s_1088, float64(6), float64(1)), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Conj.X_invoke_Arity2(s_1088, float64(1))}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 4 5 6 7 1] (assoc s 6 1) (conj s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(27), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, s_1088)) {
				} else {
					panic((&js.Error{("Assert failed: (= 27 (reduce + s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(s_1088, cljs_core.Vec.X_invoke_Arity1(s_1088)) {
				} else {
					panic((&js.Error{("Assert failed: (= s (vec s))")}))
				}
				{
					var m_1089 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), float64(1)}, nil})
					_ = m_1089
					if cljs_core.X_EQ_.Arity2IIB(m_1089, cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(s_1088, m_1089))) {
					} else {
						panic((&js.Error{("Assert failed: (= m (meta (with-meta s m)))")}))
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1090 interface{}) {
					defer func() {
						if e627 := recover(); e627 != nil {
							if cljs_core.Value_(e627).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e627
									_ = e
									return__1090 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e627)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1087, float64(0), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 0 6) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1091 interface{}) {
					defer func() {
						if e628 := recover(); e628 != nil {
							if cljs_core.Value_(e628).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e628
									_ = e
									return__1091 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e628)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1087, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1092 interface{}) {
					defer func() {
						if e629 := recover(); e629 != nil {
							if cljs_core.Value_(e629).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e629
									_ = e
									return__1092 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e629)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1087, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1093 interface{}) {
					defer func() {
						if e630 := recover(); e630 != nil {
							if cljs_core.Value_(e630).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e630
									_ = e
									return__1093 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e630)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1087, float64(3), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 3 6) (catch js/Error e :fail)))")}))
				}
				if reflect.DeepEqual(v1_1086, cljs_core.Subvec.X_invoke_Arity3(s_1088, float64(1), float64(4)).(*cljs_core.CljsCoreSubvec).V) {
				} else {
					panic((&js.Error{("Assert failed: (identical? v1 (.-v (subvec s 1 4)))")}))
				}
				{
					var sentinel_1094 = cljs_core.Rand.Arity0F()
					var s_1095___1 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2))
					_, _ = sentinel_1094, s_1095___1
					if reflect.DeepEqual(sentinel_1094, func() (return__1096 interface{}) {
						defer func() {
							if e631 := recover(); e631 != nil {
								if cljs_core.Value_(e631).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var ___ = e631
										_ = ___
										return__1096 = sentinel_1094
									}
								} else {
									panic(e631)

								}
							}
						}()
						{
							{
								var G__632 = float64(-1)
								_ = G__632
								return s_1095___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__632)
							}
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (identical? sentinel (try (s -1) (catch js/Error _ sentinel)))")}))
					}
					if reflect.DeepEqual(sentinel_1094, func() (return__1097 interface{}) {
						defer func() {
							if e633 := recover(); e633 != nil {
								if cljs_core.Value_(e633).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var ___ = e633
										_ = ___
										return__1097 = sentinel_1094
									}
								} else {
									panic(e633)

								}
							}
						}()
						{
							{
								var G__634 = float64(1)
								_ = G__634
								return s_1095___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__634)
							}
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (identical? sentinel (try (s 1) (catch js/Error _ sentinel)))")}))
					}
				}
				{
					var sv1_1098 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2)).(*cljs_core.CljsCoreSubvec)
					var sv2_1099 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(1)).(*cljs_core.CljsCoreSubvec)
					_, _ = sv1_1098, sv2_1099
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(sv1_1098), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(*cljs_core.CljsCoreList)) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq sv1) (quote (1)))")}))
					}
					if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(sv2_1099)) {
					} else {
						panic((&js.Error{("Assert failed: (nil? (rseq sv2))")}))
					}
				}
			}
			{
				var v1_1100 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(15), float64(48)).(*cljs_core.CljsCoreRange))
				var v2_1101 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(40), float64(57)).(*cljs_core.CljsCoreRange))
				var v1_1102___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Pop_BANG_.X_invoke_Arity1(cljs_core.Transient.X_invoke_Arity1(v1_1100)), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v2_1103___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(v2_1101), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v_1104 = cljs_core.Into.X_invoke_Arity2(v1_1102___1, v2_1103___1)
				_, _, _, _, _ = v1_1100, v2_1101, v1_1102___1, v2_1103___1, v_1104
				if cljs_core.X_EQ_.Arity2IIB(v_1104, cljs_core.Vec.X_invoke_Arity1(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(16), float64(47)).(*cljs_core.CljsCoreRange), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(41), float64(57)).(*cljs_core.CljsCoreRange), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})})).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (= v (vec (concat [:quux] (range 16 47) [:foo] [:quux] (range 41 57) [:bar])))")}))
				}
			}
			{
				var v_1105 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)
				var xs_1106 interface{} = cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)
				_, _ = v_1105, xs_1106
				for {
					{
						var temp__4220__auto___1107 = cljs_core.First.X_invoke_Arity1(xs_1106)
						_ = temp__4220__auto___1107
						if cljs_core.Truth_(temp__4220__auto___1107) {
							{
								var x_1108 = temp__4220__auto___1107
								_ = x_1108
								v_1105, xs_1106 = func() interface{} {
									var pred__635 = func(G__1109 *cljs_core.AFn, v_1105 interface{}, xs_1106 interface{}, x_1108 interface{}, temp__4220__auto___1107 interface{}) *cljs_core.AFn {
										return cljs_core.Fn(G__1109, 2, func(p1__67_SHARP_ interface{}, p2__68_SHARP_ interface{}) interface{} {
											{
												var G__638 = cljs_core.Mod.X_invoke_Arity2(p2__68_SHARP_, float64(3)).(float64)
												_ = G__638
												return p1__67_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__638)
											}
										})
									}(&cljs_core.AFn{}, v_1105, xs_1106, x_1108, temp__4220__auto___1107)
									var expr__636 = x_1108
									_, _ = pred__635, expr__636
									if cljs_core.Truth_(pred__635.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(0), nil, float64(2), nil}, nil}, nil}), expr__636)) {
										return cljs_core.Conj_BANG_.X_invoke_Arity2(v_1105, x_1108)
									} else {
										if cljs_core.Truth_(pred__635.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), expr__636)) {
											return cljs_core.Assoc_BANG_.X_invoke_Arity3(v_1105, cljs_core.Count.X_invoke_Arity1(v_1105).(float64), x_1108)
										} else {
											panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(expr__636).(string))}))
										}
									}
								}(), cljs_core.Next.Arity1IQ(xs_1106)
								continue
							}
						} else {
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)), cljs_core.Persistent_BANG_.X_invoke_Arity1(v_1105)) {
							} else {
								panic((&js.Error{("Assert failed: (= (vec (range 100)) (persistent! v))")}))
							}
						}
					}
					break
				}
			}
			{
				var m1_1110 interface{} = cljs_core.CljsCorePersistentHashMap_EMPTY
				var m2_1111 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_EMPTY)
				var i_1112 = float64(0)
				_, _, _ = m1_1110, m2_1111, i_1112
				for {
					if i_1112 < float64(100) {
						m1_1110, m2_1111, i_1112 = cljs_core.Assoc.X_invoke_Arity3(m1_1110, i_1112, i_1112), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_1111, i_1112, i_1112), (i_1112 + float64(1))
						continue
					} else {
						{
							var m2_1113___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_1111)
							_ = m2_1113___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1110).(float64), float64(100)) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m1) 100)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1113___1).(float64), float64(100)) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m2) 100)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_1110, m2_1113___1) {
							} else {
								panic((&js.Error{("Assert failed: (= m1 m2)")}))
							}
							{
								var i_1114___1 = float64(0)
								_ = i_1114___1
								for {
									if i_1114___1 < float64(100) {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__639 = i_1114___1
											_ = G__639
											return m1_1110.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__639)
										}(), i_1114___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__640 = i_1114___1
											_ = G__640
											return m2_1113___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__640)
										}(), i_1114___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m2 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_1110, i_1114___1), i_1114___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_1113___1, i_1114___1), i_1114___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m2 i) i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_1110, i_1114___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m1 i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_1113___1, i_1114___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m2 i)")}))
										}
										i_1114___1 = (i_1114___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_1110))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range 100) (range 100)) (sort-by first (seq m1)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_1113___1))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range 100) (range 100)) (sort-by first (seq m2)))")}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_1110, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{("Assert failed: (not (contains? (dissoc m1 3) 3))")}))
							}
						}
					}
					break
				}
			}
			{
				var m_1115 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_1115
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1115).(float64), float64(7)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 7)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1115, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
				}
			}
			{
				var m_1116 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_1116
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1116).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1116, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var m_1117 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_1117
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1117).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1117, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var tm_1118 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_1118
				{
					var tm_1119___1 interface{} = tm_1118
					var ks_1120 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_1119___1, ks_1120
					for {
						{
							var temp__4220__auto___1121 = cljs_core.First.X_invoke_Arity1(ks_1120)
							_ = temp__4220__auto___1121
							if cljs_core.Truth_(temp__4220__auto___1121) {
								{
									var k_1122 = temp__4220__auto___1121
									_ = k_1122
									tm_1119___1, ks_1120 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1119___1, k_1122), cljs_core.Next.Arity1IQ(ks_1120)
									continue
								}
							} else {
								{
									var m_1123 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1119___1)
									_ = m_1123
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1123).(float64), float64(7)) {
									} else {
										panic((&js.Error{("Assert failed: (= (count m) 7)")}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_1123, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var tm_1124 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_1124
				{
					var seq__641_1125 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__642_1126 interface{} = nil
					var count__643_1127 = float64(0)
					var i__644_1128 = float64(0)
					_, _, _, _ = seq__641_1125, chunk__642_1126, count__643_1127, i__644_1128
					for {
						if i__644_1128 < count__643_1127 {
							{
								var k_1129 = chunk__642_1126.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__644_1128)
								_ = k_1129
								if cljs_core.X_EQ_.Arity2IIB(k_1129, cljs_core.Get.X_invoke_Arity2(tm_1124, k_1129)) {
								} else {
									panic((&js.Error{("Assert failed: (= k (get tm k))")}))
								}
								seq__641_1125, chunk__642_1126, count__643_1127, i__644_1128 = seq__641_1125, chunk__642_1126, count__643_1127, (i__644_1128 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___1130 = cljs_core.Seq.Arity1IQ(seq__641_1125)
								_ = temp__4222__auto___1130
								if cljs_core.Truth_(temp__4222__auto___1130) {
									{
										var seq__641_1131___1 = temp__4222__auto___1130
										_ = seq__641_1131___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__641_1131___1) {
											{
												var c__954__auto___1132 = cljs_core.Chunk_first.X_invoke_Arity1(seq__641_1131___1)
												_ = c__954__auto___1132
												seq__641_1125, chunk__642_1126, count__643_1127, i__644_1128 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__641_1131___1), c__954__auto___1132, cljs_core.Count.X_invoke_Arity1(c__954__auto___1132).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_1133 = cljs_core.First.X_invoke_Arity1(seq__641_1131___1)
												_ = k_1133
												if cljs_core.X_EQ_.Arity2IIB(k_1133, cljs_core.Get.X_invoke_Arity2(tm_1124, k_1133)) {
												} else {
													panic((&js.Error{("Assert failed: (= k (get tm k))")}))
												}
												seq__641_1125, chunk__642_1126, count__643_1127, i__644_1128 = cljs_core.Next.Arity1IQ(seq__641_1131___1), nil, float64(0), float64(0)
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
					var m_1134 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1124)
					_ = m_1134
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1135 float64) {
						defer func() {
							if e645 := recover(); e645 != nil {
								if cljs_core.Value_(e645).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e645
										_ = e
										return__1135 = float64(2)
									}
								} else {
									panic(e645)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1124, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1136 float64) {
						defer func() {
							if e646 := recover(); e646 != nil {
								if cljs_core.Value_(e646).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e646
										_ = e
										return__1136 = float64(2)
									}
								} else {
									panic(e646)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_1124, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1137 float64) {
						defer func() {
							if e647 := recover(); e647 != nil {
								if cljs_core.Value_(e647).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e647
										_ = e
										return__1137 = float64(2)
									}
								} else {
									panic(e647)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1124)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (persistent! tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1138 float64) {
						defer func() {
							if e648 := recover(); e648 != nil {
								if cljs_core.Value_(e648).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e648
										_ = e
										return__1138 = float64(2)
									}
								} else {
									panic(e648)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_1124)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (count tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_1134, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var m_1139 = cljs_core.Assoc.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentHashMap_EMPTY, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_ = m_1139
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1139, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1139, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1139).(float64), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 2)")}))
				}
				{
					var m_1140___1 = cljs_core.Dissoc.X_invoke_Arity2(m_1139, Fixed_hash_foo)
					_ = m_1140___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1140___1, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1140___1, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1140___1).(float64), float64(1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 1)")}))
					}
				}
			}
			{
				var m_1141 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_1142___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m_1141, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_, _ = m_1141, m_1142___1
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1142___1).(float64), float64(102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 102)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1142___1, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1142___1, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				{
					var m_1143___2 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_1142___1, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_1143___2
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1143___2, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1143___2, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1143___2).(float64), float64(98)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 98)")}))
					}
				}
			}
			{
				var m_1144 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_1145___1 = cljs_core.Transient.X_invoke_Arity1(m_1144)
				var m_1146___2 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_1145___1, Fixed_hash_foo, float64(1))
				var m_1147___3 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_1146___2, Fixed_hash_bar, float64(2))
				var m_1148___4 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m_1147___3)
				_, _, _, _, _ = m_1144, m_1145___1, m_1146___2, m_1147___3, m_1148___4
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1148___4).(float64), float64(102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 102)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1148___4, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1148___4, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				{
					var m_1149___5 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_1148___4, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_1149___5
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1149___5, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1149___5, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1149___5).(float64), float64(98)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 98)")}))
					}
				}
			}
			Array_map_conversion_threshold = cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD

			{
				var m1_1150 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_1151 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY)
				var i_1152 = float64(0)
				_, _, _ = m1_1150, m2_1151, i_1152
				for {
					if i_1152 < Array_map_conversion_threshold {
						m1_1150, m2_1151, i_1152 = cljs_core.Assoc.X_invoke_Arity3(m1_1150, i_1152, i_1152), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_1151, i_1152, i_1152), (i_1152 + float64(1))
						continue
					} else {
						{
							var m2_1153___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_1151)
							_ = m2_1153___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1150).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m1) array-map-conversion-threshold)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1153___1).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m2) array-map-conversion-threshold)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_1150, m2_1153___1) {
							} else {
								panic((&js.Error{("Assert failed: (= m1 m2)")}))
							}
							{
								var i_1154___1 = float64(0)
								_ = i_1154___1
								for {
									if i_1154___1 < Array_map_conversion_threshold {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__649 = i_1154___1
											_ = G__649
											return m1_1150.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__649)
										}(), i_1154___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__650 = i_1154___1
											_ = G__650
											return m2_1153___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__650)
										}(), i_1154___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m2 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_1150, i_1154___1), i_1154___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_1153___1, i_1154___1), i_1154___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m2 i) i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_1150, i_1154___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m1 i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_1153___1, i_1154___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m2 i)")}))
										}
										i_1154___1 = (i_1154___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_1150))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m1)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_1153___1))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m2)))")}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_1150, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{("Assert failed: (not (contains? (dissoc m1 3) 3))")}))
							}
						}
					}
					break
				}
			}
			{
				var m_1155 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_1155
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1155).(float64), float64(7)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 7)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1155, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
				}
			}
			{
				var m_1156 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_1156
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1156).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1156, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var m_1157 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_1157
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1157).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1157, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var tm_1158 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_1158
				{
					var tm_1159___1 interface{} = tm_1158
					var ks_1160 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_1159___1, ks_1160
					for {
						{
							var temp__4220__auto___1161 = cljs_core.First.X_invoke_Arity1(ks_1160)
							_ = temp__4220__auto___1161
							if cljs_core.Truth_(temp__4220__auto___1161) {
								{
									var k_1162 = temp__4220__auto___1161
									_ = k_1162
									tm_1159___1, ks_1160 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1159___1, k_1162), cljs_core.Next.Arity1IQ(ks_1160)
									continue
								}
							} else {
								{
									var m_1163 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1159___1)
									_ = m_1163
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1163).(float64), float64(7)) {
									} else {
										panic((&js.Error{("Assert failed: (= (count m) 7)")}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_1163, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var tm_1164 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_1164
				{
					var seq__651_1165 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__652_1166 interface{} = nil
					var count__653_1167 = float64(0)
					var i__654_1168 = float64(0)
					_, _, _, _ = seq__651_1165, chunk__652_1166, count__653_1167, i__654_1168
					for {
						if i__654_1168 < count__653_1167 {
							{
								var k_1169 = chunk__652_1166.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__654_1168)
								_ = k_1169
								if cljs_core.X_EQ_.Arity2IIB(k_1169, cljs_core.Get.X_invoke_Arity2(tm_1164, k_1169)) {
								} else {
									panic((&js.Error{("Assert failed: (= k (get tm k))")}))
								}
								seq__651_1165, chunk__652_1166, count__653_1167, i__654_1168 = seq__651_1165, chunk__652_1166, count__653_1167, (i__654_1168 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___1170 = cljs_core.Seq.Arity1IQ(seq__651_1165)
								_ = temp__4222__auto___1170
								if cljs_core.Truth_(temp__4222__auto___1170) {
									{
										var seq__651_1171___1 = temp__4222__auto___1170
										_ = seq__651_1171___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__651_1171___1) {
											{
												var c__954__auto___1172 = cljs_core.Chunk_first.X_invoke_Arity1(seq__651_1171___1)
												_ = c__954__auto___1172
												seq__651_1165, chunk__652_1166, count__653_1167, i__654_1168 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__651_1171___1), c__954__auto___1172, cljs_core.Count.X_invoke_Arity1(c__954__auto___1172).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_1173 = cljs_core.First.X_invoke_Arity1(seq__651_1171___1)
												_ = k_1173
												if cljs_core.X_EQ_.Arity2IIB(k_1173, cljs_core.Get.X_invoke_Arity2(tm_1164, k_1173)) {
												} else {
													panic((&js.Error{("Assert failed: (= k (get tm k))")}))
												}
												seq__651_1165, chunk__652_1166, count__653_1167, i__654_1168 = cljs_core.Next.Arity1IQ(seq__651_1171___1), nil, float64(0), float64(0)
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
					var m_1174 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1164)
					_ = m_1174
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1175 float64) {
						defer func() {
							if e655 := recover(); e655 != nil {
								if cljs_core.Value_(e655).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e655
										_ = e
										return__1175 = float64(2)
									}
								} else {
									panic(e655)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1164, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1176 float64) {
						defer func() {
							if e656 := recover(); e656 != nil {
								if cljs_core.Value_(e656).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e656
										_ = e
										return__1176 = float64(2)
									}
								} else {
									panic(e656)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_1164, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1177 float64) {
						defer func() {
							if e657 := recover(); e657 != nil {
								if cljs_core.Value_(e657).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e657
										_ = e
										return__1177 = float64(2)
									}
								} else {
									panic(e657)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1164)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (persistent! tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1178 float64) {
						defer func() {
							if e658 := recover(); e658 != nil {
								if cljs_core.Value_(e658).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e658
										_ = e
										return__1178 = float64(2)
									}
								} else {
									panic(e658)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_1164)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (count tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_1174, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
					}
				}
			}
			{
				var m_1179 = cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))
				_ = m_1179
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1179).(float64), (float64(2) * Array_map_conversion_threshold)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) (* 2 array-map-conversion-threshold))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__659 = Array_map_conversion_threshold
					_ = G__659
					return m_1179.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__659)
				}(), Array_map_conversion_threshold) {
				} else {
					panic((&js.Error{("Assert failed: (= (m array-map-conversion-threshold) array-map-conversion-threshold)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1179, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity2(func(G__1180 *cljs_core.AFn, m_1179 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1180, 1, func(p1__69_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__69_SHARP_, p1__69_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1179), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (= m (into (.-EMPTY cljs.core/PersistentHashMap) (map (fn* [p1__69#] (vector p1__69# p1__69#)) (range (* 2 array-map-conversion-threshold)))))")}))
				}
			}
			{
				var m1_1181 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_1182 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var i_1183 = float64(0)
				_, _, _ = m1_1181, m2_1182, i_1183
				for {
					if i_1183 < float64(100) {
						m1_1181, m2_1182, i_1183 = cljs_core.Assoc.X_invoke_Arity3(m1_1181, i_1183, i_1183), cljs_core.Assoc.X_invoke_Arity3(m2_1182, ("foo"+cljs_core.Str.X_invoke_Arity1(i_1183).(string)), i_1183), (i_1183 + float64(1))
						continue
					} else {
						if cljs_core.X_EQ_.Arity2IIB(m1_1181, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{("Assert failed: (= m1 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (range 100) (range 100))))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(m2_1182, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Map_.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.Str, "foo").(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{("Assert failed: (= m2 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (map (partial str \"foo\") (range 100)) (range 100))))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1181).(float64), float64(100)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count m1) 100)")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1182).(float64), float64(100)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count m2) 100)")}))
						}
					}
					break
				}
			}
			{
				var i_1184 = float64(0)
				var m_1185 interface{} = cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(-1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))
				var result_1186 interface{} = cljs_core.CljsCorePersistentVector_EMPTY
				_, _, _ = i_1184, m_1185, result_1186
				for {
					if i_1184 <= (cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) {
						i_1184, m_1185, result_1186 = (i_1184 + float64(1)), cljs_core.Assoc.X_invoke_Arity3(m_1185, i_1184, i_1184), cljs_core.Conj.X_invoke_Arity2(result_1186, cljs_core.Meta.X_invoke_Arity1(m_1185))
						continue
					} else {
						{
							var n_1187 = ((cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) + float64(1))
							var expected_1188 = cljs_core.Repeat.X_invoke_Arity2(n_1187, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})).(*cljs_core.CljsCoreLazySeq)
							_, _ = n_1187, expected_1188
							if cljs_core.X_EQ_.Arity2IIB(result_1186, expected_1188) {
							} else {
								panic((&js.Error{("Assert failed: (= result expected)")}))
							}
						}
					}
					break
				}
			}
			{
				var m1_1189 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_1190 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var m2_1191 = cljs_core.Sorted_map_by.X_invoke_ArityVariadic(c2_1190, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{})).(*cljs_core.CljsCorePersistentTreeMap)
				_, _, _ = m1_1189, c2_1190, m2_1191
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeMap)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(m1_1189)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeMap (type m1))")}))
				}
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeMap)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(m2_1191)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeMap (type m2))")}))
				}
				if reflect.DeepEqual(cljs_core.Compare, cljs_core.Native_get_instance_field.X_invoke_Arity2(m1_1189, "Comp")) {
				} else {
					panic((&js.Error{("Assert failed: (identical? compare (.-comp m1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(m1_1189).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count m1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(m2_1191).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count m2))")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(m1_1189)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq m1))")}))
				}
				{
					var m1_1192___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m1_1189, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					var m2_1193___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m2_1191, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					_, _ = m1_1192___1, m2_1193___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1192___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m1) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1193___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m2) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_1192___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m1) (list [:bar 2] [:foo 1] [:quux 3]))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_1193___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m2) (list [:quux 3] [:foo 1] [:bar 2]))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_1192___1), cljs_core.Rseq.Arity1IQ(m2_1193___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m1) (rseq m2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_1193___1), cljs_core.Rseq.Arity1IQ(m1_1192___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m2) (rseq m1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m1_1192___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= (conj m1 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m1_1192___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count (conj m1 [:wibble 4])) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m2_1193___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= (conj m2 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m2_1193___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count (conj m2 [:wibble 4])) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m1_1192___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})).(cljs_core.CljsCoreICollection).X_conj_Arity2(nil)) {
					} else {
						panic((&js.Error{("Assert failed: (= (map key (assoc m1 nil 4)) (list nil :bar :foo :quux))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m2_1193___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(nil).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (map key (assoc m2 nil 4)) (list :quux :foo :bar nil))")}))
					}
				}
			}
			{
				var m_1194 = cljs_core.Apply.X_invoke_Arity2(cljs_core.Sorted_map, cljs_core.Mapcat.X_invoke_ArityVariadic(func(G__1197 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1197, 1, func(p1__70_SHARP_ interface{}) interface{} {
						return cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(p1__70_SHARP_).(cljs_core.CljsCoreICollection).X_conj_Arity2(p1__70_SHARP_)
					})
				}(&cljs_core.AFn{}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_ArityVariadic(cljs_core.Partial.X_invoke_Arity2(cljs_core.Apply, cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(10)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(20), float64(30)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(50), float64(60)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(30), float64(40)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(40), float64(50)}, nil})}, nil})}))})))
				var s1_1195 = cljs_core.Map_.X_invoke_Arity2(func(G__1198 *cljs_core.AFn, m_1194 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1198, 1, func(p1__71_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__71_SHARP_, p1__71_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1194), cljs_core.Range_.X_invoke_Arity1(float64(60)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				var s2_1196 = cljs_core.Map_.X_invoke_Arity2(func(G__1199 *cljs_core.AFn, m_1194 interface{}, s1_1195 *cljs_core.CljsCoreLazySeq) *cljs_core.AFn {
					return cljs_core.Fn(G__1199, 1, func(p1__72_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__72_SHARP_, p1__72_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1194, s1_1195), cljs_core.Range_.X_invoke_Arity3(float64(59), float64(-1), float64(-1)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				_, _, _ = m_1194, s1_1195, s2_1196
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1194).(float64), float64(60)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 60)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m_1194), s1_1195) {
				} else {
					panic((&js.Error{("Assert failed: (= (seq m) s1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(m_1194), s2_1196) {
				} else {
					panic((&js.Error{("Assert failed: (= (rseq m) s2)")}))
				}
			}
			{
				var m_1200 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
				_ = m_1200
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m_1200, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(2), float64(3)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (dissoc m :foo) (hash-map :bar 2 :quux 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(m_1200, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))).(float64), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count (dissoc m :foo)) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(m_1200), cljs_core.Hash.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(1), float64(2), float64(3)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (hash m) (hash (hash-map :foo 1 :bar 2 :quux 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1200, cljs_core.X_LT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m < :foo) (list [:bar 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1200, cljs_core.X_LT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m <= :foo) (list [:bar 2] [:foo 1]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1200, cljs_core.X_GT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m > :foo) (list [:quux 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1200, cljs_core.X_GT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m >= :foo) (list [:foo 1] [:quux 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1201 *cljs_core.AFn, m_1200 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1201, 1, func(p1__73_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity2(func(G__1202 *cljs_core.AFn, m_1200 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__1202, 2, func(___ interface{}, x interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_1200), p1__73_SHARP_)
					})
				}(&cljs_core.AFn{}, m_1200), m_1200).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__73#] (reduce (fn [_ x] x) p1__73#)) m) (list 2 1 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1203 *cljs_core.AFn, m_1200 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1203, 1, func(p1__74_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity3(func(G__1204 *cljs_core.AFn, m_1200 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__1204, 2, func(x interface{}, ___ interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_1200), float64(7), p1__74_SHARP_)
					})
				}(&cljs_core.AFn{}, m_1200), m_1200).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__74#] (reduce (fn [x _] x) 7 p1__74#)) m) (list 7 7 7))")}))
				}
			}
			{
				var s1_1205 = cljs_core.Sorted_set.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_1206 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var s2_1207 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c2_1206, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c3_1208 = func(G__1211 *cljs_core.AFn, s1_1205 interface{}, c2_1206 cljs_core.CljsCoreIFn, s2_1207 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1211, 2, func(p1__75_SHARP_ interface{}, p2__76_SHARP_ interface{}) interface{} {
						return cljs_core.Compare.Arity2IIF(cljs_core.Quot.X_invoke_Arity2(p1__75_SHARP_, float64(2)).(float64), cljs_core.Quot.X_invoke_Arity2(p2__76_SHARP_, float64(2)).(float64))
					})
				}(&cljs_core.AFn{}, s1_1205, c2_1206, s2_1207)
				var s3_1209 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c3_1208, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var s4_1210 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(cljs_core.X_LT_, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				_, _, _, _, _, _ = s1_1205, c2_1206, s2_1207, c3_1208, s3_1209, s4_1210
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(s1_1205)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeSet (type s1))")}))
				}
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(s2_1207)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeSet (type s2))")}))
				}
				if reflect.DeepEqual(cljs_core.Compare, s1_1205.(cljs_core.CljsCoreISorted).X_comparator_Arity1()) {
				} else {
					panic((&js.Error{("Assert failed: (identical? compare (-comparator s1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(s1_1205).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count s1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(s2_1207).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count s2))")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(s1_1205)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq s1))")}))
				}
				{
					var s1_1212___1 = cljs_core.Conj.X_invoke_ArityVariadic(s1_1205, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s2_1213___1 = cljs_core.Conj.X_invoke_ArityVariadic(s2_1207, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s3_1214___1 = cljs_core.Conj.X_invoke_ArityVariadic(s3_1209, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(7), float64(8), float64(9)}))
					var s4_1215___1 = cljs_core.Conj.X_invoke_ArityVariadic(s4_1210, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					_, _, _, _ = s1_1212___1, s2_1213___1, s3_1214___1, s4_1215___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_1212___1), cljs_core.Hash.X_invoke_Arity1(s2_1213___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (hash s1) (hash s2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_1212___1), cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (hash s1) (hash #{1 3 2}))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_1212___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq s1) (list 1 2 3))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_1212___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq s1) (list 3 2 1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_1213___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq s2) (list 3 2 1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_1213___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq s2) (list 1 2 3))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_1212___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s1) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_1213___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s2) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s3_1214___1).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s3) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(s3_1214___1, float64(0)), float64(1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get s3 0) 1)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1214___1, cljs_core.X_GT_, float64(5)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 > 5) (list 7 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1214___1, cljs_core.X_GT_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 > 6) (list 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1214___1, cljs_core.X_GT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 6) (list 7 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1214___1, cljs_core.X_LT_, float64(0)), cljs_core.CljsCoreList_EMPTY) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 0) (list))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1214___1, cljs_core.X_LT_, float64(5)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 5) (list 1 2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1214___1, cljs_core.X_LT_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 6) (list 1 2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1214___1, cljs_core.X_LT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 <= 6) (list 1 2 7))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s3_1214___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 2 <= 6) (list 2 7))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s4_1215___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT_, float64(3)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s4 >= 2 < 3) (list 2))")}))
					}
					{
						var s1_1216___2 = cljs_core.Disj.X_invoke_Arity2(s1_1212___1, float64(2))
						var s2_1217___2 = cljs_core.Disj.X_invoke_Arity2(s2_1213___1, float64(2))
						_, _ = s1_1216___2, s2_1217___2
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_1216___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{("Assert failed: (= (seq s1) (list 1 3))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_1216___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{("Assert failed: (= (rseq s1) (list 3 1))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_1217___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{("Assert failed: (= (seq s2) (list 3 1))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_1217___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{("Assert failed: (= (rseq s2) (list 1 3))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_1216___2).(float64), float64(2)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count s1) 2)")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_1217___2).(float64), float64(2)) {
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
					return cljs_core.Fn(map__GT_Person, 1, func(G__662 interface{}) interface{} {
						return (&CljsCore_testPerson{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}).X_invoke_Arity1(G__662), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}).X_invoke_Arity1(G__662), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__662, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})})), nil})
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
					return cljs_core.Fn(map__GT_A, 1, func(G__681 interface{}) interface{} {
						return (&CljsCore_testA{nil, cljs_core.Dissoc.X_invoke_Arity1(G__681), nil})
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
					return cljs_core.Fn(map__GT_C, 1, func(G__692 interface{}) interface{} {
						return (&CljsCore_testC{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__692), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__692), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}).X_invoke_Arity1(G__692), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__692, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})})), nil})
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
				var s_1218 = "abc"
				_ = s_1218
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_1218).Length) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (.-length s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_1218).Length) {
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
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1218).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_("abc").Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring \"abc\" 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", func(G__1219 *cljs_core.AFn, s_1218 string) *cljs_core.AFn {
					return cljs_core.Fn(G__1219, 2, func(target713 interface{}, start interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target713, "Substring", []interface{}{start})
					})
				}(&cljs_core.AFn{}, s_1218).X_invoke_Arity2(s_1218, float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" ((memfn substring start) s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1218).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s substring 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1218).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s (substring 1)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1218).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s (substring 1 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1218).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring s 1 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_(s_1218).ToUpperCase()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" (. s (toUpperCase)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_("abc").ToUpperCase()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" (. \"abc\" (toUpperCase)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", func(G__1220 *cljs_core.AFn, s_1218 string) *cljs_core.AFn {
					return cljs_core.Fn(G__1220, 1, func(target714 interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target714, "ToUpperCase", []interface{}{})
					})
				}(&cljs_core.AFn{}, s_1218).X_invoke_Arity1(s_1218)) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" ((memfn toUpperCase) s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("BC", cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_1218).ToUpperCase(), "Substring", []interface{}{float64(1)})) {
				} else {
					panic((&js.Error{("Assert failed: (= \"BC\" (. (. s (toUpperCase)) substring 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_1218).ToUpperCase(), "Substring", []interface{}{float64(1)}), "Length")) {
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
					return cljs_core.Fn(map__GT_A2, 1, func(G__717 interface{}) interface{} {
						return (&CljsCore_testA2{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__717), nil, cljs_core.Dissoc.X_invoke_Arity2(G__717, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
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
					return cljs_core.Fn(map__GT_B, 1, func(G__732 interface{}) interface{} {
						return (&CljsCore_testB{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__732), nil, cljs_core.Dissoc.X_invoke_Arity2(G__732, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
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

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(func() *CljsCore_testT745 {
				X__GT_t745 = func(__GT_t745 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_t745, 2, func(test_stuff___1 interface{}, meta746 interface{}) interface{} {
						return (&CljsCore_testT745{test_stuff___1, meta746})
					})
				}(&cljs_core.AFn{})

				return (&CljsCore_testT745{test_stuff, nil})
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

			Foo2.X_add_method_Arity3(float64(0), func(G__1221 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1221, 1, func(x interface{}) interface{} {
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
				var g_1222 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}))}, true).(*cljs_core.CljsCorePersistentHashSet)
				var h_1223 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil})}, true).(*cljs_core.CljsCorePersistentHashSet)
				_, _ = g_1222, h_1223
				if cljs_core.X_EQ_.Arity2IIB(g_1222, h_1223) {
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
				var fv_1224 = (&CljsCore_testFirst{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})})
				var fs_1225 = (&CljsCore_testFirst{"asdf"})
				_, _ = fv_1224, fs_1225
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fv_1224.X_invoke_Arity0()
				}(), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (fv) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fs_1225.X_invoke_Arity0()
				}(), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((`` + cljs_core.Str.X_invoke_Arity1(fs_1225).(string)), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (str fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_1224.X_get_first_Arity1(), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-get-first fv) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fs_1225.X_get_first_Arity1(), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (-get-first fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_1224.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-find-first fv [1]) 1)")}))
				}
				if reflect.DeepEqual(func() interface{} {
					var G__760 = float64(1)
					_ = G__760
					return fv_1224.X_invoke_Arity1(G__760)
				}(), fv_1224) {
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
				var t_1226 = (&CljsCore_testDestructuringWithLocals{float64(1)})
				_ = t_1226
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(1)}, nil}), t_1226.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 1] (-find-first t [2 3]))")}))
				}
			}
			{
				var x_1227 = float64(1)
				_ = x_1227
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__764 = x_1227
					_ = G__764
					switch G__764 {
					case float64(1):
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})

					default:
						panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_1227).(string))}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x 1 :one) :one)")}))
				}
			}
			{
				var x_1229 = float64(1)
				_ = x_1229
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__765 = x_1229
					_ = G__765
					switch G__765 {
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
				var x_1231 = float64(1)
				_ = x_1231
				if cljs_core.X_EQ_.Arity2IIB(func() (return__1232 interface{}) {
					defer func() {
						if e766 := recover(); e766 != nil {
							if cljs_core.Value_(e766).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e766
									_ = e
									return__1232 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e766)

							}
						}
					}()
					{
						{
							var G__767 = x_1231
							_ = G__767
							switch G__767 {
							case float64(3):
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "three", Fqn: "three", X_hash: float64(-1651831795)})

							default:
								panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_1231).(string))}))

							}
						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (try (case x 3 :three) (catch js/Error e :fail)) :fail)")}))
				}
			}
			{
				var x_1234 = float64(1)
				_ = x_1234
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__768 = x_1234
					_ = G__768
					switch G__768 {
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
				var x_1236 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})
				_ = x_1236
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__769 = x_1236
					_ = G__769
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), G__769) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
					} else {
						panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_1236).(string))}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x [:a :b] :ok) :ok)")}))
				}
			}
			{
				var a_1237 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = a_1237
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__770 = a_1237
					_ = G__770
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__770) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__770) {
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
				var a_1238 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil})
				_ = a_1238
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__771 = a_1238
					_ = G__771
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__771) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__771) {
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
				var foo_1239 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = foo_1239
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__772 = foo_1239
					_ = G__772
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__772) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__772) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), G__772) {
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
					var G__773 = foo_1239
					_ = G__773
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "d", Str: "d", X_hash: float64(-682293345), X_meta: nil}), G__773) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__773) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__773) {
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
				var r_1240 = cljs_core.Range_.X_invoke_Arity1(float64(64)).(*cljs_core.CljsCoreRange)
				var v_1241 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, r_1240)
				_, _ = r_1240, v_1241
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1241)), cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1241))) {
				} else {
					panic((&js.Error{("Assert failed: (= (hash (seq v)) (hash (seq v)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Array_chunk.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)}).(*cljs_core.CljsCoreArrayChunk))) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (reduce + (array-chunk (array 1 2 3))))")}))
				}
				if cljs_core.Value_(cljs_core.Seq.Arity1IQ(v_1241)).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreChunkedSeq)(nil)).Elem()) {
				} else {
					panic((&js.Error{("Assert failed: (instance? ChunkedSeq (seq v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(r_1240, cljs_core.Seq.Arity1IQ(v_1241)) {
				} else {
					panic((&js.Error{("Assert failed: (= r (seq v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, r_1240).(*cljs_core.CljsCoreLazySeq), cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, v_1241).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (map inc r) (map inc v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, r_1240).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, v_1241).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (filter even? r) (filter even? v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, r_1240).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, v_1241).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (filter odd? r) (filter odd? v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic(r_1240, r_1240, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{r_1240})).(*cljs_core.CljsCoreLazySeq), cljs_core.Concat.X_invoke_ArityVariadic(v_1241, v_1241, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{v_1241})).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (concat r r r) (concat v v v))")}))
				}
				if cljs_core.Value_(cljs_core.Seq.Arity1IQ(v_1241)).Type().Implements(reflect.TypeOf((*cljs_core.CljsCoreIReduce)(nil)).Elem()) {
				} else {
					panic((&js.Error{("Assert failed: (satisfies? IReduce (seq v))")}))
				}
				if float64(2010) == cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1241)))))).(float64) {
				} else {
					panic((&js.Error{("Assert failed: (== 2010 (reduce + (nnext (nnext (seq v)))))")}))
				}
				if float64(2020) == cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1241)))))).(float64) {
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
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__1242 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1242, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), nil).(*cljs_core.CljsCoreCons)
				})
			}(&cljs_core.AFn{}), nil, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next (lazy-seq (cons 1 nil))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__1243 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1243, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), (&cljs_core.CljsCoreLazySeq{nil, func(G__1244 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__1244, 0, func() interface{} {
							return cljs_core.Cons.X_invoke_Arity2(float64(2), (&cljs_core.CljsCoreLazySeq{nil, func(G__1245 *cljs_core.AFn) *cljs_core.AFn {
								return cljs_core.Fn(G__1245, 0, func() interface{} {
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
					return cljs_core.Fn(map__GT_PrintMe, 1, func(G__776 interface{}) interface{} {
						return (&CljsCore_testPrintMe{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__776), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__776), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__776, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})})), nil})
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
				var seq__793_1246 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(13)).(*cljs_core.CljsCoreRange))
				var chunk__806_1247 interface{} = nil
				var count__807_1248 = float64(0)
				var i__808_1249 = float64(0)
				_, _, _, _ = seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249
				for {
					if i__808_1249 < count__807_1248 {
						{
							var month_1250 = chunk__806_1247.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__808_1249)
							_ = month_1250
							{
								var seq__809_1251 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
								var chunk__814_1252 interface{} = nil
								var count__815_1253 = float64(0)
								var i__816_1254 = float64(0)
								_, _, _, _ = seq__809_1251, chunk__814_1252, count__815_1253, i__816_1254
								for {
									if i__816_1254 < count__815_1253 {
										{
											var day_1255 = chunk__814_1252.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__816_1254)
											_ = day_1255
											{
												var seq__817_1256 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
												var chunk__818_1257 interface{} = nil
												var count__819_1258 = float64(0)
												var i__820_1259 = float64(0)
												_, _, _, _ = seq__817_1256, chunk__818_1257, count__819_1258, i__820_1259
												for {
													if i__820_1259 < count__819_1258 {
														{
															var hour_1260 = chunk__818_1257.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__820_1259)
															_ = hour_1260
															{
																var pad_1261 = func(G__1263 *cljs_core.AFn, seq__817_1256 interface{}, chunk__818_1257 interface{}, count__819_1258 float64, i__820_1259 float64, seq__809_1251 interface{}, chunk__814_1252 interface{}, count__815_1253 float64, i__816_1254 float64, seq__793_1246 interface{}, chunk__806_1247 interface{}, count__807_1248 float64, i__808_1249 float64, hour_1260 interface{}, day_1255 interface{}, month_1250 interface{}) *cljs_core.AFn {
																	return cljs_core.Fn(G__1263, 1, func(n interface{}) interface{} {
																		if n.(float64) < float64(10) {
																			return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																		} else {
																			return n
																		}
																	})
																}(&cljs_core.AFn{}, seq__817_1256, chunk__818_1257, count__819_1258, i__820_1259, seq__809_1251, chunk__814_1252, count__815_1253, i__816_1254, seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249, hour_1260, day_1255, month_1250)
																var inst_1262 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1261.X_invoke_Arity1(month_1250)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1261.X_invoke_Arity1(day_1255)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1261.X_invoke_Arity1(hour_1260)).(string) + ":14:15.666-00:00")
																_, _ = pad_1261, inst_1262
																if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1262})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1262).(string) + "\"")) {
																} else {
																	panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																}
															}
															seq__817_1256, chunk__818_1257, count__819_1258, i__820_1259 = seq__817_1256, chunk__818_1257, count__819_1258, (i__820_1259 + float64(1))
															continue
														}
													} else {
														{
															var temp__4222__auto___1264 = cljs_core.Seq.Arity1IQ(seq__817_1256)
															_ = temp__4222__auto___1264
															if cljs_core.Truth_(temp__4222__auto___1264) {
																{
																	var seq__817_1265___1 = temp__4222__auto___1264
																	_ = seq__817_1265___1
																	if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__817_1265___1) {
																		{
																			var c__954__auto___1266 = cljs_core.Chunk_first.X_invoke_Arity1(seq__817_1265___1)
																			_ = c__954__auto___1266
																			seq__817_1256, chunk__818_1257, count__819_1258, i__820_1259 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__817_1265___1), c__954__auto___1266, cljs_core.Count.X_invoke_Arity1(c__954__auto___1266).(float64), float64(0)
																			continue
																		}
																	} else {
																		{
																			var hour_1267 = cljs_core.First.X_invoke_Arity1(seq__817_1265___1)
																			_ = hour_1267
																			{
																				var pad_1268 = func(G__1270 *cljs_core.AFn, seq__817_1256 interface{}, chunk__818_1257 interface{}, count__819_1258 float64, i__820_1259 float64, seq__809_1251 interface{}, chunk__814_1252 interface{}, count__815_1253 float64, i__816_1254 float64, seq__793_1246 interface{}, chunk__806_1247 interface{}, count__807_1248 float64, i__808_1249 float64, hour_1267 interface{}, seq__817_1265___1 interface{}, temp__4222__auto___1264 cljs_core.CljsCoreISeq, day_1255 interface{}, month_1250 interface{}) *cljs_core.AFn {
																					return cljs_core.Fn(G__1270, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__817_1256, chunk__818_1257, count__819_1258, i__820_1259, seq__809_1251, chunk__814_1252, count__815_1253, i__816_1254, seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249, hour_1267, seq__817_1265___1, temp__4222__auto___1264, day_1255, month_1250)
																				var inst_1269 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1268.X_invoke_Arity1(month_1250)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1268.X_invoke_Arity1(day_1255)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1268.X_invoke_Arity1(hour_1267)).(string) + ":14:15.666-00:00")
																				_, _ = pad_1268, inst_1269
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1269})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1269).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__817_1256, chunk__818_1257, count__819_1258, i__820_1259 = cljs_core.Next.Arity1IQ(seq__817_1265___1), nil, float64(0), float64(0)
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
											seq__809_1251, chunk__814_1252, count__815_1253, i__816_1254 = seq__809_1251, chunk__814_1252, count__815_1253, (i__816_1254 + float64(1))
											continue
										}
									} else {
										{
											var temp__4222__auto___1271 = cljs_core.Seq.Arity1IQ(seq__809_1251)
											_ = temp__4222__auto___1271
											if cljs_core.Truth_(temp__4222__auto___1271) {
												{
													var seq__809_1272___1 = temp__4222__auto___1271
													_ = seq__809_1272___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__809_1272___1) {
														{
															var c__954__auto___1273 = cljs_core.Chunk_first.X_invoke_Arity1(seq__809_1272___1)
															_ = c__954__auto___1273
															seq__809_1251, chunk__814_1252, count__815_1253, i__816_1254 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__809_1272___1), c__954__auto___1273, cljs_core.Count.X_invoke_Arity1(c__954__auto___1273).(float64), float64(0)
															continue
														}
													} else {
														{
															var day_1274 = cljs_core.First.X_invoke_Arity1(seq__809_1272___1)
															_ = day_1274
															{
																var seq__810_1275 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																var chunk__811_1276 interface{} = nil
																var count__812_1277 = float64(0)
																var i__813_1278 = float64(0)
																_, _, _, _ = seq__810_1275, chunk__811_1276, count__812_1277, i__813_1278
																for {
																	if i__813_1278 < count__812_1277 {
																		{
																			var hour_1279 = chunk__811_1276.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__813_1278)
																			_ = hour_1279
																			{
																				var pad_1280 = func(G__1282 *cljs_core.AFn, seq__810_1275 interface{}, chunk__811_1276 interface{}, count__812_1277 float64, i__813_1278 float64, seq__809_1251 interface{}, chunk__814_1252 interface{}, count__815_1253 float64, i__816_1254 float64, seq__793_1246 interface{}, chunk__806_1247 interface{}, count__807_1248 float64, i__808_1249 float64, hour_1279 interface{}, day_1274 interface{}, seq__809_1272___1 interface{}, temp__4222__auto___1271 cljs_core.CljsCoreISeq, month_1250 interface{}) *cljs_core.AFn {
																					return cljs_core.Fn(G__1282, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__810_1275, chunk__811_1276, count__812_1277, i__813_1278, seq__809_1251, chunk__814_1252, count__815_1253, i__816_1254, seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249, hour_1279, day_1274, seq__809_1272___1, temp__4222__auto___1271, month_1250)
																				var inst_1281 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1280.X_invoke_Arity1(month_1250)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1280.X_invoke_Arity1(day_1274)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1280.X_invoke_Arity1(hour_1279)).(string) + ":14:15.666-00:00")
																				_, _ = pad_1280, inst_1281
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1281})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1281).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__810_1275, chunk__811_1276, count__812_1277, i__813_1278 = seq__810_1275, chunk__811_1276, count__812_1277, (i__813_1278 + float64(1))
																			continue
																		}
																	} else {
																		{
																			var temp__4222__auto___1283___1 = cljs_core.Seq.Arity1IQ(seq__810_1275)
																			_ = temp__4222__auto___1283___1
																			if cljs_core.Truth_(temp__4222__auto___1283___1) {
																				{
																					var seq__810_1284___1 = temp__4222__auto___1283___1
																					_ = seq__810_1284___1
																					if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__810_1284___1) {
																						{
																							var c__954__auto___1285 = cljs_core.Chunk_first.X_invoke_Arity1(seq__810_1284___1)
																							_ = c__954__auto___1285
																							seq__810_1275, chunk__811_1276, count__812_1277, i__813_1278 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__810_1284___1), c__954__auto___1285, cljs_core.Count.X_invoke_Arity1(c__954__auto___1285).(float64), float64(0)
																							continue
																						}
																					} else {
																						{
																							var hour_1286 = cljs_core.First.X_invoke_Arity1(seq__810_1284___1)
																							_ = hour_1286
																							{
																								var pad_1287 = func(G__1289 *cljs_core.AFn, seq__810_1275 interface{}, chunk__811_1276 interface{}, count__812_1277 float64, i__813_1278 float64, seq__809_1251 interface{}, chunk__814_1252 interface{}, count__815_1253 float64, i__816_1254 float64, seq__793_1246 interface{}, chunk__806_1247 interface{}, count__807_1248 float64, i__808_1249 float64, hour_1286 interface{}, seq__810_1284___1 interface{}, temp__4222__auto___1283___1 cljs_core.CljsCoreISeq, day_1274 interface{}, seq__809_1272___1 interface{}, temp__4222__auto___1271 cljs_core.CljsCoreISeq, month_1250 interface{}) *cljs_core.AFn {
																									return cljs_core.Fn(G__1289, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__810_1275, chunk__811_1276, count__812_1277, i__813_1278, seq__809_1251, chunk__814_1252, count__815_1253, i__816_1254, seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249, hour_1286, seq__810_1284___1, temp__4222__auto___1283___1, day_1274, seq__809_1272___1, temp__4222__auto___1271, month_1250)
																								var inst_1288 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1287.X_invoke_Arity1(month_1250)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1287.X_invoke_Arity1(day_1274)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1287.X_invoke_Arity1(hour_1286)).(string) + ":14:15.666-00:00")
																								_, _ = pad_1287, inst_1288
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1288})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1288).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__810_1275, chunk__811_1276, count__812_1277, i__813_1278 = cljs_core.Next.Arity1IQ(seq__810_1284___1), nil, float64(0), float64(0)
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
															seq__809_1251, chunk__814_1252, count__815_1253, i__816_1254 = cljs_core.Next.Arity1IQ(seq__809_1272___1), nil, float64(0), float64(0)
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
							seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249 = seq__793_1246, chunk__806_1247, count__807_1248, (i__808_1249 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1290 = cljs_core.Seq.Arity1IQ(seq__793_1246)
							_ = temp__4222__auto___1290
							if cljs_core.Truth_(temp__4222__auto___1290) {
								{
									var seq__793_1291___1 = temp__4222__auto___1290
									_ = seq__793_1291___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__793_1291___1) {
										{
											var c__954__auto___1292 = cljs_core.Chunk_first.X_invoke_Arity1(seq__793_1291___1)
											_ = c__954__auto___1292
											seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__793_1291___1), c__954__auto___1292, cljs_core.Count.X_invoke_Arity1(c__954__auto___1292).(float64), float64(0)
											continue
										}
									} else {
										{
											var month_1293 = cljs_core.First.X_invoke_Arity1(seq__793_1291___1)
											_ = month_1293
											{
												var seq__794_1294 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
												var chunk__799_1295 interface{} = nil
												var count__800_1296 = float64(0)
												var i__801_1297 = float64(0)
												_, _, _, _ = seq__794_1294, chunk__799_1295, count__800_1296, i__801_1297
												for {
													if i__801_1297 < count__800_1296 {
														{
															var day_1298 = chunk__799_1295.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__801_1297)
															_ = day_1298
															{
																var seq__802_1299 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																var chunk__803_1300 interface{} = nil
																var count__804_1301 = float64(0)
																var i__805_1302 = float64(0)
																_, _, _, _ = seq__802_1299, chunk__803_1300, count__804_1301, i__805_1302
																for {
																	if i__805_1302 < count__804_1301 {
																		{
																			var hour_1303 = chunk__803_1300.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__805_1302)
																			_ = hour_1303
																			{
																				var pad_1304 = func(G__1306 *cljs_core.AFn, seq__802_1299 interface{}, chunk__803_1300 interface{}, count__804_1301 float64, i__805_1302 float64, seq__794_1294 interface{}, chunk__799_1295 interface{}, count__800_1296 float64, i__801_1297 float64, seq__793_1246 interface{}, chunk__806_1247 interface{}, count__807_1248 float64, i__808_1249 float64, hour_1303 interface{}, day_1298 interface{}, month_1293 interface{}, seq__793_1291___1 interface{}, temp__4222__auto___1290 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																					return cljs_core.Fn(G__1306, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__802_1299, chunk__803_1300, count__804_1301, i__805_1302, seq__794_1294, chunk__799_1295, count__800_1296, i__801_1297, seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249, hour_1303, day_1298, month_1293, seq__793_1291___1, temp__4222__auto___1290)
																				var inst_1305 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1304.X_invoke_Arity1(month_1293)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1304.X_invoke_Arity1(day_1298)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1304.X_invoke_Arity1(hour_1303)).(string) + ":14:15.666-00:00")
																				_, _ = pad_1304, inst_1305
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1305})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1305).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__802_1299, chunk__803_1300, count__804_1301, i__805_1302 = seq__802_1299, chunk__803_1300, count__804_1301, (i__805_1302 + float64(1))
																			continue
																		}
																	} else {
																		{
																			var temp__4222__auto___1307___1 = cljs_core.Seq.Arity1IQ(seq__802_1299)
																			_ = temp__4222__auto___1307___1
																			if cljs_core.Truth_(temp__4222__auto___1307___1) {
																				{
																					var seq__802_1308___1 = temp__4222__auto___1307___1
																					_ = seq__802_1308___1
																					if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__802_1308___1) {
																						{
																							var c__954__auto___1309 = cljs_core.Chunk_first.X_invoke_Arity1(seq__802_1308___1)
																							_ = c__954__auto___1309
																							seq__802_1299, chunk__803_1300, count__804_1301, i__805_1302 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__802_1308___1), c__954__auto___1309, cljs_core.Count.X_invoke_Arity1(c__954__auto___1309).(float64), float64(0)
																							continue
																						}
																					} else {
																						{
																							var hour_1310 = cljs_core.First.X_invoke_Arity1(seq__802_1308___1)
																							_ = hour_1310
																							{
																								var pad_1311 = func(G__1313 *cljs_core.AFn, seq__802_1299 interface{}, chunk__803_1300 interface{}, count__804_1301 float64, i__805_1302 float64, seq__794_1294 interface{}, chunk__799_1295 interface{}, count__800_1296 float64, i__801_1297 float64, seq__793_1246 interface{}, chunk__806_1247 interface{}, count__807_1248 float64, i__808_1249 float64, hour_1310 interface{}, seq__802_1308___1 interface{}, temp__4222__auto___1307___1 cljs_core.CljsCoreISeq, day_1298 interface{}, month_1293 interface{}, seq__793_1291___1 interface{}, temp__4222__auto___1290 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																									return cljs_core.Fn(G__1313, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__802_1299, chunk__803_1300, count__804_1301, i__805_1302, seq__794_1294, chunk__799_1295, count__800_1296, i__801_1297, seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249, hour_1310, seq__802_1308___1, temp__4222__auto___1307___1, day_1298, month_1293, seq__793_1291___1, temp__4222__auto___1290)
																								var inst_1312 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1311.X_invoke_Arity1(month_1293)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1311.X_invoke_Arity1(day_1298)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1311.X_invoke_Arity1(hour_1310)).(string) + ":14:15.666-00:00")
																								_, _ = pad_1311, inst_1312
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1312})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1312).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__802_1299, chunk__803_1300, count__804_1301, i__805_1302 = cljs_core.Next.Arity1IQ(seq__802_1308___1), nil, float64(0), float64(0)
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
															seq__794_1294, chunk__799_1295, count__800_1296, i__801_1297 = seq__794_1294, chunk__799_1295, count__800_1296, (i__801_1297 + float64(1))
															continue
														}
													} else {
														{
															var temp__4222__auto___1314___1 = cljs_core.Seq.Arity1IQ(seq__794_1294)
															_ = temp__4222__auto___1314___1
															if cljs_core.Truth_(temp__4222__auto___1314___1) {
																{
																	var seq__794_1315___1 = temp__4222__auto___1314___1
																	_ = seq__794_1315___1
																	if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__794_1315___1) {
																		{
																			var c__954__auto___1316 = cljs_core.Chunk_first.X_invoke_Arity1(seq__794_1315___1)
																			_ = c__954__auto___1316
																			seq__794_1294, chunk__799_1295, count__800_1296, i__801_1297 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__794_1315___1), c__954__auto___1316, cljs_core.Count.X_invoke_Arity1(c__954__auto___1316).(float64), float64(0)
																			continue
																		}
																	} else {
																		{
																			var day_1317 = cljs_core.First.X_invoke_Arity1(seq__794_1315___1)
																			_ = day_1317
																			{
																				var seq__795_1318 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																				var chunk__796_1319 interface{} = nil
																				var count__797_1320 = float64(0)
																				var i__798_1321 = float64(0)
																				_, _, _, _ = seq__795_1318, chunk__796_1319, count__797_1320, i__798_1321
																				for {
																					if i__798_1321 < count__797_1320 {
																						{
																							var hour_1322 = chunk__796_1319.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__798_1321)
																							_ = hour_1322
																							{
																								var pad_1323 = func(G__1325 *cljs_core.AFn, seq__795_1318 interface{}, chunk__796_1319 interface{}, count__797_1320 float64, i__798_1321 float64, seq__794_1294 interface{}, chunk__799_1295 interface{}, count__800_1296 float64, i__801_1297 float64, seq__793_1246 interface{}, chunk__806_1247 interface{}, count__807_1248 float64, i__808_1249 float64, hour_1322 interface{}, day_1317 interface{}, seq__794_1315___1 interface{}, temp__4222__auto___1314___1 cljs_core.CljsCoreISeq, month_1293 interface{}, seq__793_1291___1 interface{}, temp__4222__auto___1290 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																									return cljs_core.Fn(G__1325, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__795_1318, chunk__796_1319, count__797_1320, i__798_1321, seq__794_1294, chunk__799_1295, count__800_1296, i__801_1297, seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249, hour_1322, day_1317, seq__794_1315___1, temp__4222__auto___1314___1, month_1293, seq__793_1291___1, temp__4222__auto___1290)
																								var inst_1324 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1323.X_invoke_Arity1(month_1293)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1323.X_invoke_Arity1(day_1317)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1323.X_invoke_Arity1(hour_1322)).(string) + ":14:15.666-00:00")
																								_, _ = pad_1323, inst_1324
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1324})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1324).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__795_1318, chunk__796_1319, count__797_1320, i__798_1321 = seq__795_1318, chunk__796_1319, count__797_1320, (i__798_1321 + float64(1))
																							continue
																						}
																					} else {
																						{
																							var temp__4222__auto___1326___2 = cljs_core.Seq.Arity1IQ(seq__795_1318)
																							_ = temp__4222__auto___1326___2
																							if cljs_core.Truth_(temp__4222__auto___1326___2) {
																								{
																									var seq__795_1327___1 = temp__4222__auto___1326___2
																									_ = seq__795_1327___1
																									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__795_1327___1) {
																										{
																											var c__954__auto___1328 = cljs_core.Chunk_first.X_invoke_Arity1(seq__795_1327___1)
																											_ = c__954__auto___1328
																											seq__795_1318, chunk__796_1319, count__797_1320, i__798_1321 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__795_1327___1), c__954__auto___1328, cljs_core.Count.X_invoke_Arity1(c__954__auto___1328).(float64), float64(0)
																											continue
																										}
																									} else {
																										{
																											var hour_1329 = cljs_core.First.X_invoke_Arity1(seq__795_1327___1)
																											_ = hour_1329
																											{
																												var pad_1330 = func(G__1332 *cljs_core.AFn, seq__795_1318 interface{}, chunk__796_1319 interface{}, count__797_1320 float64, i__798_1321 float64, seq__794_1294 interface{}, chunk__799_1295 interface{}, count__800_1296 float64, i__801_1297 float64, seq__793_1246 interface{}, chunk__806_1247 interface{}, count__807_1248 float64, i__808_1249 float64, hour_1329 interface{}, seq__795_1327___1 interface{}, temp__4222__auto___1326___2 cljs_core.CljsCoreISeq, day_1317 interface{}, seq__794_1315___1 interface{}, temp__4222__auto___1314___1 cljs_core.CljsCoreISeq, month_1293 interface{}, seq__793_1291___1 interface{}, temp__4222__auto___1290 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																													return cljs_core.Fn(G__1332, 1, func(n interface{}) interface{} {
																														if n.(float64) < float64(10) {
																															return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																														} else {
																															return n
																														}
																													})
																												}(&cljs_core.AFn{}, seq__795_1318, chunk__796_1319, count__797_1320, i__798_1321, seq__794_1294, chunk__799_1295, count__800_1296, i__801_1297, seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249, hour_1329, seq__795_1327___1, temp__4222__auto___1326___2, day_1317, seq__794_1315___1, temp__4222__auto___1314___1, month_1293, seq__793_1291___1, temp__4222__auto___1290)
																												var inst_1331 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1330.X_invoke_Arity1(month_1293)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1330.X_invoke_Arity1(day_1317)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1330.X_invoke_Arity1(hour_1329)).(string) + ":14:15.666-00:00")
																												_, _ = pad_1330, inst_1331
																												if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1331})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1331).(string) + "\"")) {
																												} else {
																													panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																												}
																											}
																											seq__795_1318, chunk__796_1319, count__797_1320, i__798_1321 = cljs_core.Next.Arity1IQ(seq__795_1327___1), nil, float64(0), float64(0)
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
																			seq__794_1294, chunk__799_1295, count__800_1296, i__801_1297 = cljs_core.Next.Arity1IQ(seq__794_1315___1), nil, float64(0), float64(0)
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
											seq__793_1246, chunk__806_1247, count__807_1248, i__808_1249 = cljs_core.Next.Arity1IQ(seq__793_1291___1), nil, float64(0), float64(0)
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
				var uuid_str_1333 = "550e8400-e29b-41d4-a716-446655440000"
				var uuid_1334 = (&cljs_core.CljsCoreUUID{uuid_str_1333})
				_, _ = uuid_str_1333, uuid_1334
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{uuid_1334})).(string), ("#uuid \"" + cljs_core.Str.X_invoke_Arity1(uuid_str_1333).(string) + "\"")) {
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
					X__GT_t826 = func(__GT_t826 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t826, 4, func(f___1 interface{}, baz___1 interface{}, test_stuff___1 interface{}, meta827 interface{}) interface{} {
							return (&CljsCore_testT826{f___1, baz___1, test_stuff___1, meta827})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT826{f, baz, test_stuff, nil})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(float64(2), Baz.X_invoke_Arity1(cljs_core.Inc).(*CljsCore_testT826).X_bar_Arity2(float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (-bar (baz inc) 1))")}))
			}
			{
				var x_1335 = "original"
				_ = x_1335
				Original_closure_stmt = func(original_closure_stmt *cljs_core.AFn, x_1335 string) *cljs_core.AFn {
					return cljs_core.Fn(original_closure_stmt, 0, func() interface{} {
						return x_1335
					})
				}(&cljs_core.AFn{}, x_1335)

			}
			{
				var x_1336 = "overwritten"
				_ = x_1336
				if cljs_core.X_EQ_.Arity2IIB("original", Original_closure_stmt.X_invoke_Arity0().(string)) {
				} else {
					panic((&js.Error{("Assert failed: (= \"original\" (original-closure-stmt))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB("original", func() string {
				var x = "original"
				var oce = func(G__1337 *cljs_core.AFn, x string) *cljs_core.AFn {
					return cljs_core.Fn(G__1337, 0, func() interface{} {
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
					var x_1338___1 = func(G__1339 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__1339, 0, func() interface{} {
							return "overwritten"
						})
					}(&cljs_core.AFn{})
					_ = x_1338___1
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
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)}), cljs_core.Reduce_kv.X_invoke_Arity3(func(G__1340 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1340, 3, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
								return cljs_core.Reduced.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)})).(*cljs_core.CljsCoreReduced)
							})
						}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentVector_EMPTY, data)) {
						} else {
							panic((&js.Error{("Assert failed: (= :reduced (reduce-kv (fn [_ _ _] (reduced :reduced)) [] data))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sort.X_invoke_Arity1(expect), cljs_core.Sort.X_invoke_Arity1(cljs_core.Reduce_kv.X_invoke_Arity3(func(G__1341 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1341, 3, func(r interface{}, k interface{}, v interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), func() (return__1342 interface{}) {
				defer func() {
					if e831 := recover(); e831 != nil {
						if cljs_core.Value_(e831).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreExceptionInfo)(nil)).Elem()) {
							{
								var e = e831
								_ = e
								return__1342 = cljs_core.Ex_data.X_invoke_Arity1(e)
							}
						} else {
							panic(e831)

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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(1), float64(2)}, nil}), func(G__1343 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1343, 0, func(more__ ...interface{}) interface{} {
					var more = cljs_core.Seq.Arity1IQ(more__[0])
					_ = more
					return more
				})
			}(&cljs_core.AFn{}).X_invoke_Arity3(js.Undefined, float64(1), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= [js/undefined 1 2] ((fn [& more] more) js/undefined 1 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(4), float64(5)}, nil}), func(G__1344 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1344, 2, func(a_b_more__ ...interface{}) interface{} {
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
				var fs_1345 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = fs_1345
				{
					var seq__832_1346 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity1(float64(4)).(*cljs_core.CljsCoreRange))
					var chunk__834_1347 interface{} = nil
					var count__835_1348 = float64(0)
					var i__836_1349 = float64(0)
					_, _, _, _ = seq__832_1346, chunk__834_1347, count__835_1348, i__836_1349
					for {
						if i__836_1349 < count__835_1348 {
							{
								var x_1350 = chunk__834_1347.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__836_1349)
								_ = x_1350
								{
									var y_1351 = (x_1350.(float64) + float64(1))
									var f_1352 = func(G__1353 *cljs_core.AFn, seq__832_1346 interface{}, chunk__834_1347 interface{}, count__835_1348 float64, i__836_1349 float64, y_1351 float64, x_1350 interface{}, fs_1345 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
										return cljs_core.Fn(G__1353, 0, func() interface{} {
											return y_1351
										})
									}(&cljs_core.AFn{}, seq__832_1346, chunk__834_1347, count__835_1348, i__836_1349, y_1351, x_1350, fs_1345)
									_, _ = y_1351, f_1352
									cljs_core.Swap_BANG_.X_invoke_Arity3(fs_1345, cljs_core.Conj, f_1352)
								}
								seq__832_1346, chunk__834_1347, count__835_1348, i__836_1349 = seq__832_1346, chunk__834_1347, count__835_1348, (i__836_1349 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___1354 = cljs_core.Seq.Arity1IQ(seq__832_1346)
								_ = temp__4222__auto___1354
								if cljs_core.Truth_(temp__4222__auto___1354) {
									{
										var seq__832_1355___1 = temp__4222__auto___1354
										_ = seq__832_1355___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__832_1355___1) {
											{
												var c__954__auto___1356 = cljs_core.Chunk_first.X_invoke_Arity1(seq__832_1355___1)
												_ = c__954__auto___1356
												seq__832_1346, chunk__834_1347, count__835_1348, i__836_1349 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__832_1355___1), c__954__auto___1356, cljs_core.Count.X_invoke_Arity1(c__954__auto___1356).(float64), float64(0)
												continue
											}
										} else {
											{
												var x_1357 = cljs_core.First.X_invoke_Arity1(seq__832_1355___1)
												_ = x_1357
												{
													var y_1358 = (x_1357.(float64) + float64(1))
													var f_1359 = func(G__1360 *cljs_core.AFn, seq__832_1346 interface{}, chunk__834_1347 interface{}, count__835_1348 float64, i__836_1349 float64, y_1358 float64, x_1357 interface{}, seq__832_1355___1 interface{}, temp__4222__auto___1354 cljs_core.CljsCoreISeq, fs_1345 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
														return cljs_core.Fn(G__1360, 0, func() interface{} {
															return y_1358
														})
													}(&cljs_core.AFn{}, seq__832_1346, chunk__834_1347, count__835_1348, i__836_1349, y_1358, x_1357, seq__832_1355___1, temp__4222__auto___1354, fs_1345)
													_, _ = y_1358, f_1359
													cljs_core.Swap_BANG_.X_invoke_Arity3(fs_1345, cljs_core.Conj, f_1359)
												}
												seq__832_1346, chunk__834_1347, count__835_1348, i__836_1349 = cljs_core.Next.Arity1IQ(seq__832_1355___1), nil, float64(0), float64(0)
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1361 *cljs_core.AFn, fs_1345 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__1361, 1, func(p1__77_SHARP_ interface{}) interface{} {
						{
							return p1__77_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
						}
					})
				}(&cljs_core.AFn{}, fs_1345), cljs_core.Deref.X_invoke_Arity1(fs_1345)).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__77#] (p1__77#)) (clojure.core/deref fs)) (quote (1 2 3 4)))")}))
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCoreLazySeq{nil, func(G__1362 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1362, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCoreLazySeq{nil, func(G__1363 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1363, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreLazySeq{nil, func(G__1364 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1364, 0, func() interface{} {
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
						var _STAR_print_fn_STAR_838_1365 = cljs_core.X_STAR_print_fn_STAR_
						_ = _STAR_print_fn_STAR_838_1365
						func() {
							defer func() {
								cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_838_1365

							}()
							{
								cljs_core.X_STAR_print_fn_STAR_ = func(G__1366 *cljs_core.AFn, _STAR_print_fn_STAR_838_1365 interface{}, sb__1124__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
									return cljs_core.Fn(G__1366, 1, func(x__1125__auto__ interface{}) interface{} {
										return sb__1124__auto__.Append(x__1125__auto__)
									})
								}(&cljs_core.AFn{}, _STAR_print_fn_STAR_838_1365, sb__1124__auto__)

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
			if (cljs_core.First.X_invoke_Arity1(cljs_core.Filter.X_invoke_Arity2(func(G__1367 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1367, 1, func(p1__78_SHARP_ interface{}) interface{} {
					return (p1__78_SHARP_.(float64) == float64(9999))
				})
			}(&cljs_core.AFn{}), cljs_core.Range_.X_invoke_Arity0().(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)).(float64) == float64(9999)) {
			} else {
				panic((&js.Error{("Assert failed: (== (first (filter (fn* [p1__78#] (== p1__78# 9999)) (range))) 9999)")}))
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
				var a_1368 = func() *CljsCore_testT839 {
					X__GT_t839 = func(__GT_t839 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t839, 2, func(test_stuff___1 interface{}, meta840 interface{}) interface{} {
							return (&CljsCore_testT839{test_stuff___1, meta840})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT839{test_stuff, nil})
				}()
				var b_1369 = func() *CljsCore_testT842 {
					X__GT_t842 = func(__GT_t842 *cljs_core.AFn, a_1368 *CljsCore_testT839) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t842, 3, func(a___1 interface{}, test_stuff___1 interface{}, meta843 interface{}) interface{} {
							return (&CljsCore_testT842{a___1, test_stuff___1, meta843})
						})
					}(&cljs_core.AFn{}, a_1368)

					return (&CljsCore_testT842{a_1368, test_stuff, nil})
				}()
				var s_1370 = cljs_core.Set.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(128)).(*cljs_core.CljsCoreRange))
				_, _, _ = a_1368, b_1369, s_1370
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_1370, a_1368, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_1369}))), a_1368)), a_1368), cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_1370, a_1368, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_1369}))), a_1368)), a_1368)) {
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
				var seq__845_1371 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__846_1372 interface{} = nil
				var count__847_1373 = float64(0)
				var i__848_1374 = float64(0)
				_, _, _, _ = seq__845_1371, chunk__846_1372, count__847_1373, i__848_1374
				for {
					if i__848_1374 < count__847_1373 {
						{
							var n_1375 = chunk__846_1372.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__848_1374)
							_ = n_1375
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1376 interface{}) {
								defer func() {
									if e851 := recover(); e851 != nil {
										if cljs_core.Value_(e851).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e851
												_ = e
												return__1376 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e851)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1375, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1377 interface{}) {
								defer func() {
									if e852 := recover(); e852 != nil {
										if cljs_core.Value_(e852).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e852
												_ = e
												return__1377 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e852)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_1375, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))")}))
							}
							seq__845_1371, chunk__846_1372, count__847_1373, i__848_1374 = seq__845_1371, chunk__846_1372, count__847_1373, (i__848_1374 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1378 = cljs_core.Seq.Arity1IQ(seq__845_1371)
							_ = temp__4222__auto___1378
							if cljs_core.Truth_(temp__4222__auto___1378) {
								{
									var seq__845_1379___1 = temp__4222__auto___1378
									_ = seq__845_1379___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__845_1379___1) {
										{
											var c__954__auto___1380 = cljs_core.Chunk_first.X_invoke_Arity1(seq__845_1379___1)
											_ = c__954__auto___1380
											seq__845_1371, chunk__846_1372, count__847_1373, i__848_1374 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__845_1379___1), c__954__auto___1380, cljs_core.Count.X_invoke_Arity1(c__954__auto___1380).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1381 = cljs_core.First.X_invoke_Arity1(seq__845_1379___1)
											_ = n_1381
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1382 interface{}) {
												defer func() {
													if e853 := recover(); e853 != nil {
														if cljs_core.Value_(e853).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e853
																_ = e
																return__1382 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e853)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1381, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1383 interface{}) {
												defer func() {
													if e854 := recover(); e854 != nil {
														if cljs_core.Value_(e854).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e854
																_ = e
																return__1383 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e854)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_1381, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))")}))
											}
											seq__845_1371, chunk__846_1372, count__847_1373, i__848_1374 = cljs_core.Next.Arity1IQ(seq__845_1379___1), nil, float64(0), float64(0)
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
				var seq__855_1384 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__856_1385 interface{} = nil
				var count__857_1386 = float64(0)
				var i__858_1387 = float64(0)
				_, _, _, _ = seq__855_1384, chunk__856_1385, count__857_1386, i__858_1387
				for {
					if i__858_1387 < count__857_1386 {
						{
							var n_1388 = chunk__856_1385.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__858_1387)
							_ = n_1388
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1389 interface{}) {
								defer func() {
									if e861 := recover(); e861 != nil {
										if cljs_core.Value_(e861).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e861
												_ = e
												return__1389 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e861)

										}
									}
								}()
								{
									return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1388, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
							}
							seq__855_1384, chunk__856_1385, count__857_1386, i__858_1387 = seq__855_1384, chunk__856_1385, count__857_1386, (i__858_1387 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1390 = cljs_core.Seq.Arity1IQ(seq__855_1384)
							_ = temp__4222__auto___1390
							if cljs_core.Truth_(temp__4222__auto___1390) {
								{
									var seq__855_1391___1 = temp__4222__auto___1390
									_ = seq__855_1391___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__855_1391___1) {
										{
											var c__954__auto___1392 = cljs_core.Chunk_first.X_invoke_Arity1(seq__855_1391___1)
											_ = c__954__auto___1392
											seq__855_1384, chunk__856_1385, count__857_1386, i__858_1387 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__855_1391___1), c__954__auto___1392, cljs_core.Count.X_invoke_Arity1(c__954__auto___1392).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1393 = cljs_core.First.X_invoke_Arity1(seq__855_1391___1)
											_ = n_1393
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1394 interface{}) {
												defer func() {
													if e862 := recover(); e862 != nil {
														if cljs_core.Value_(e862).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e862
																_ = e
																return__1394 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e862)

														}
													}
												}()
												{
													return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1393, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
											}
											seq__855_1384, chunk__856_1385, count__857_1386, i__858_1387 = cljs_core.Next.Arity1IQ(seq__855_1391___1), nil, float64(0), float64(0)
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
				var map__863_1395 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__863_1396___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__863_1395) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__863_1395)
					} else {
						return map__863_1395
					}
				}()
				var b_1397 = cljs_core.Get.X_invoke_Arity2(map__863_1396___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a_1398 = cljs_core.Get.X_invoke_Arity2(map__863_1396___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__863_1395, map__863_1396___1, b_1397, a_1398
				if cljs_core.X_EQ_.Arity2IIB(float64(1), a_1398) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), b_1397) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 b)")}))
				}
			}
			{
				var map__864_1399 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__864_1400___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__864_1399) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__864_1399)
					} else {
						return map__864_1399
					}
				}()
				var d_1401 = cljs_core.Get.X_invoke_Arity2(map__864_1400___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_1402 = cljs_core.Get.X_invoke_Arity2(map__864_1400___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__864_1399, map__864_1400___1, d_1401, b_1402
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1402) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1401) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__865_1403 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__865_1404___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__865_1403) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__865_1403)
					} else {
						return map__865_1403
					}
				}()
				var d_1405 = cljs_core.Get.X_invoke_Arity2(map__865_1404___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_1406 = cljs_core.Get.X_invoke_Arity2(map__865_1404___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__865_1403, map__865_1404___1, d_1405, b_1406
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1406) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1405) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__866_1407 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}), float64(1), (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}), float64(2)}, nil})
				var map__866_1408___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__866_1407) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__866_1407)
					} else {
						return map__866_1407
					}
				}()
				var d_1409 = cljs_core.Get.X_invoke_Arity2(map__866_1408___1, (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}))
				var b_1410 = cljs_core.Get.X_invoke_Arity2(map__866_1408___1, (&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}))
				_, _, _, _ = map__866_1407, map__866_1408___1, d_1409, b_1410
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1410) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1409) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__867_1411 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}), float64(2)}, nil})
				var map__867_1412___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__867_1411) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__867_1411)
					} else {
						return map__867_1411
					}
				}()
				var y_1413 = cljs_core.Get.X_invoke_Arity2(map__867_1412___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}))
				var x_1414 = cljs_core.Get.X_invoke_Arity2(map__867_1412___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}))
				_, _, _, _ = map__867_1411, map__867_1412___1, y_1413, x_1414
				if cljs_core.X_EQ_.Arity2IIB(x_1414, float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= x 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(y_1413, float64(2)) {
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
								arr, names = cljs_core.Conj.X_invoke_Arity2(arr, func(G__1415 *cljs_core.AFn, arr interface{}, names interface{}, name interface{}) *cljs_core.AFn {
									return cljs_core.Fn(G__1415, 0, func() interface{} {
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
					var _STAR_print_fn_STAR_868_1416 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_868_1416
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_868_1416

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__1417 *cljs_core.AFn, _STAR_print_fn_STAR_868_1416 interface{}, sb__1124__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__1417, 1, func(x__1125__auto__ interface{}) interface{} {
									return sb__1124__auto__.Append(x__1125__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_868_1416, sb__1124__auto__)

							{
								var seq__869_1418 interface{} = cljs_core.Seq.Arity1IQ(Cljs_739.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil})))
								var chunk__870_1419 interface{} = nil
								var count__871_1420 = float64(0)
								var i__872_1421 = float64(0)
								_, _, _, _ = seq__869_1418, chunk__870_1419, count__871_1420, i__872_1421
								for {
									if i__872_1421 < count__871_1420 {
										{
											var fn_1422 = chunk__870_1419.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__872_1421)
											_ = fn_1422
											{
												fn_1422.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
											}
											seq__869_1418, chunk__870_1419, count__871_1420, i__872_1421 = seq__869_1418, chunk__870_1419, count__871_1420, (i__872_1421 + float64(1))
											continue
										}
									} else {
										{
											var temp__4222__auto___1423 = cljs_core.Seq.Arity1IQ(seq__869_1418)
											_ = temp__4222__auto___1423
											if cljs_core.Truth_(temp__4222__auto___1423) {
												{
													var seq__869_1424___1 = temp__4222__auto___1423
													_ = seq__869_1424___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__869_1424___1) {
														{
															var c__954__auto___1425 = cljs_core.Chunk_first.X_invoke_Arity1(seq__869_1424___1)
															_ = c__954__auto___1425
															seq__869_1418, chunk__870_1419, count__871_1420, i__872_1421 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__869_1424___1), c__954__auto___1425, cljs_core.Count.X_invoke_Arity1(c__954__auto___1425).(float64), float64(0)
															continue
														}
													} else {
														{
															var fn_1426 = cljs_core.First.X_invoke_Arity1(seq__869_1424___1)
															_ = fn_1426
															{
																fn_1426.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
															}
															seq__869_1418, chunk__870_1419, count__871_1420, i__872_1421 = cljs_core.Next.Arity1IQ(seq__869_1424___1), nil, float64(0), float64(0)
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
				var seq__873_1427 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__874_1428 interface{} = nil
				var count__875_1429 = float64(0)
				var i__876_1430 = float64(0)
				_, _, _, _ = seq__873_1427, chunk__874_1428, count__875_1429, i__876_1430
				for {
					if i__876_1430 < count__875_1429 {
						{
							var n_1431 = chunk__874_1428.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__876_1430)
							_ = n_1431
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1431)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get [1 2] n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1432 interface{}) {
								defer func() {
									if e879 := recover(); e879 != nil {
										if cljs_core.Value_(e879).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e879
												_ = e
												return__1432 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e879)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1431)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1431, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get [1 2] n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1433 interface{}) {
								defer func() {
									if e880 := recover(); e880 != nil {
										if cljs_core.Value_(e880).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e880
												_ = e
												return__1433 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e880)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1431, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1431)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (subvec [1 2] 1) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1434 interface{}) {
								defer func() {
									if e881 := recover(); e881 != nil {
										if cljs_core.Value_(e881).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e881
												_ = e
												return__1434 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e881)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1431)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1431, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (subvec [1 2] 1) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1435 interface{}) {
								defer func() {
									if e882 := recover(); e882 != nil {
										if cljs_core.Value_(e882).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e882
												_ = e
												return__1435 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e882)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1431, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1431)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (transient [1 2]) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1436 interface{}) {
								defer func() {
									if e883 := recover(); e883 != nil {
										if cljs_core.Value_(e883).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e883
												_ = e
												return__1436 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e883)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1431)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1431, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (transient [1 2]) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1437 interface{}) {
								defer func() {
									if e884 := recover(); e884 != nil {
										if cljs_core.Value_(e884).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e884
												_ = e
												return__1437 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e884)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1431, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1431)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (range 1 3) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1438 interface{}) {
								defer func() {
									if e885 := recover(); e885 != nil {
										if cljs_core.Value_(e885).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e885
												_ = e
												return__1438 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e885)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1431)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1431, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (range 1 3) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1439 interface{}) {
								defer func() {
									if e886 := recover(); e886 != nil {
										if cljs_core.Value_(e886).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e886
												_ = e
												return__1439 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e886)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1431, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))")}))
							}
							seq__873_1427, chunk__874_1428, count__875_1429, i__876_1430 = seq__873_1427, chunk__874_1428, count__875_1429, (i__876_1430 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1440 = cljs_core.Seq.Arity1IQ(seq__873_1427)
							_ = temp__4222__auto___1440
							if cljs_core.Truth_(temp__4222__auto___1440) {
								{
									var seq__873_1441___1 = temp__4222__auto___1440
									_ = seq__873_1441___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__873_1441___1) {
										{
											var c__954__auto___1442 = cljs_core.Chunk_first.X_invoke_Arity1(seq__873_1441___1)
											_ = c__954__auto___1442
											seq__873_1427, chunk__874_1428, count__875_1429, i__876_1430 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__873_1441___1), c__954__auto___1442, cljs_core.Count.X_invoke_Arity1(c__954__auto___1442).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1443 = cljs_core.First.X_invoke_Arity1(seq__873_1441___1)
											_ = n_1443
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1443)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get [1 2] n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1444 interface{}) {
												defer func() {
													if e887 := recover(); e887 != nil {
														if cljs_core.Value_(e887).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e887
																_ = e
																return__1444 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e887)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1443)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1443, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get [1 2] n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1445 interface{}) {
												defer func() {
													if e888 := recover(); e888 != nil {
														if cljs_core.Value_(e888).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e888
																_ = e
																return__1445 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e888)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1443, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1443)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (subvec [1 2] 1) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1446 interface{}) {
												defer func() {
													if e889 := recover(); e889 != nil {
														if cljs_core.Value_(e889).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e889
																_ = e
																return__1446 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e889)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1443)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1443, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (subvec [1 2] 1) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1447 interface{}) {
												defer func() {
													if e890 := recover(); e890 != nil {
														if cljs_core.Value_(e890).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e890
																_ = e
																return__1447 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e890)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1443, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1443)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (transient [1 2]) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1448 interface{}) {
												defer func() {
													if e891 := recover(); e891 != nil {
														if cljs_core.Value_(e891).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e891
																_ = e
																return__1448 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e891)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1443)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1443, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (transient [1 2]) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1449 interface{}) {
												defer func() {
													if e892 := recover(); e892 != nil {
														if cljs_core.Value_(e892).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e892
																_ = e
																return__1449 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e892)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1443, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1443)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (range 1 3) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1450 interface{}) {
												defer func() {
													if e893 := recover(); e893 != nil {
														if cljs_core.Value_(e893).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e893
																_ = e
																return__1450 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e893)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1443)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1443, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (range 1 3) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1451 interface{}) {
												defer func() {
													if e894 := recover(); e894 != nil {
														if cljs_core.Value_(e894).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e894
																_ = e
																return__1451 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e894)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1443, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))")}))
											}
											seq__873_1427, chunk__874_1428, count__875_1429, i__876_1430 = cljs_core.Next.Arity1IQ(seq__873_1441___1), nil, float64(0), float64(0)
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
				var x_1452 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}).X_invoke_Arity1(cljs_core.Meta.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}).X_invoke_Arity1(cljs_core.Deref.X_invoke_Arity1(Cljs_780))))
				_ = x_1452
				if cljs_core.Vector_QMARK_.Arity1IB(x_1452) {
				} else {
					panic((&js.Error{("Assert failed: (vector? x)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(x_1452, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= x [1 2 3])")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3((&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}), "ToString", []interface{}{}), "550e8400-e29b-41d4-a716-446655440000") {
			} else {
				panic((&js.Error{("Assert failed: (= (.toString #uuid \"550e8400-e29b-41d4-a716-446655440000\") \"550e8400-e29b-41d4-a716-446655440000\")")}))
			}
			{
				var seq__895_1453 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__896_1454 interface{} = nil
				var count__897_1455 = float64(0)
				var i__898_1456 = float64(0)
				_, _, _, _ = seq__895_1453, chunk__896_1454, count__897_1455, i__898_1456
				for {
					if i__898_1456 < count__897_1455 {
						{
							var m_1457 = chunk__896_1454.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__898_1456)
							_ = m_1457
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__1458 interface{}) {
								defer func() {
									if e899 := recover(); e899 != nil {
										if cljs_core.Value_(e899).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var ___ = e899
												_ = ___
												return__1458 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
											}
										} else {
											panic(e899)

										}
									}
								}()
								{
									return cljs_core.Conj.X_invoke_Arity2(m_1457, "foo")
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1457, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m [:foo 1]))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1457, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m {:foo 1}))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1457, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m (list [:foo 1])))")}))
							}
							seq__895_1453, chunk__896_1454, count__897_1455, i__898_1456 = seq__895_1453, chunk__896_1454, count__897_1455, (i__898_1456 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1459 = cljs_core.Seq.Arity1IQ(seq__895_1453)
							_ = temp__4222__auto___1459
							if cljs_core.Truth_(temp__4222__auto___1459) {
								{
									var seq__895_1460___1 = temp__4222__auto___1459
									_ = seq__895_1460___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__895_1460___1) {
										{
											var c__954__auto___1461 = cljs_core.Chunk_first.X_invoke_Arity1(seq__895_1460___1)
											_ = c__954__auto___1461
											seq__895_1453, chunk__896_1454, count__897_1455, i__898_1456 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__895_1460___1), c__954__auto___1461, cljs_core.Count.X_invoke_Arity1(c__954__auto___1461).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_1462 = cljs_core.First.X_invoke_Arity1(seq__895_1460___1)
											_ = m_1462
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__1463 interface{}) {
												defer func() {
													if e900 := recover(); e900 != nil {
														if cljs_core.Value_(e900).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var ___ = e900
																_ = ___
																return__1463 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
															}
														} else {
															panic(e900)

														}
													}
												}()
												{
													return cljs_core.Conj.X_invoke_Arity2(m_1462, "foo")
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1462, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m [:foo 1]))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1462, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m {:foo 1}))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1462, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m (list [:foo 1])))")}))
											}
											seq__895_1453, chunk__896_1454, count__897_1455, i__898_1456 = cljs_core.Next.Arity1IQ(seq__895_1460___1), nil, float64(0), float64(0)
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
				var seq__901_1464 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Array_map, cljs_core.Hash_map, cljs_core.Sorted_map}, nil}))
				var chunk__902_1465 interface{} = nil
				var count__903_1466 = float64(0)
				var i__904_1467 = float64(0)
				_, _, _, _ = seq__901_1464, chunk__902_1465, count__903_1466, i__904_1467
				for {
					if i__904_1467 < count__903_1466 {
						{
							var mt_1468 = chunk__902_1465.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__904_1467)
							_ = mt_1468
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
								var G__905 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
								var G__906 = float64(1)
								_, _ = G__905, G__906
								return mt_1468.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__905, G__906)
							}(), func(make_seq *cljs_core.AFn, seq__901_1464 interface{}, chunk__902_1465 interface{}, count__903_1466 float64, i__904_1467 float64, mt_1468 interface{}) *cljs_core.AFn {
								return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
									if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
										X__GT_t912 = func(__GT_t912 *cljs_core.AFn, seq__901_1464 interface{}, chunk__902_1465 interface{}, count__903_1466 float64, i__904_1467 float64, mt_1468 interface{}) *cljs_core.AFn {
											return cljs_core.Fn(__GT_t912, 9, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, i__904___1 interface{}, count__903___1 interface{}, chunk__902___1 interface{}, seq__901___1 interface{}, test_stuff___1 interface{}, meta913 interface{}) interface{} {
												return (&CljsCore_testT912{from_seq___1, make_seq___1, mt___1, i__904___1, count__903___1, chunk__902___1, seq__901___1, test_stuff___1, meta913})
											})
										}(&cljs_core.AFn{}, seq__901_1464, chunk__902_1465, count__903_1466, i__904_1467, mt_1468)

										return (&CljsCore_testT912{from_seq, make_seq, mt_1468, i__904_1467, count__903_1466, chunk__902_1465, seq__901_1464, test_stuff, nil})
									} else {
										return nil
									}
								})
							}(&cljs_core.AFn{}, seq__901_1464, chunk__902_1465, count__903_1466, i__904_1467, mt_1468).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))")}))
							}
							seq__901_1464, chunk__902_1465, count__903_1466, i__904_1467 = seq__901_1464, chunk__902_1465, count__903_1466, (i__904_1467 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1469 = cljs_core.Seq.Arity1IQ(seq__901_1464)
							_ = temp__4222__auto___1469
							if cljs_core.Truth_(temp__4222__auto___1469) {
								{
									var seq__901_1470___1 = temp__4222__auto___1469
									_ = seq__901_1470___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__901_1470___1) {
										{
											var c__954__auto___1471 = cljs_core.Chunk_first.X_invoke_Arity1(seq__901_1470___1)
											_ = c__954__auto___1471
											seq__901_1464, chunk__902_1465, count__903_1466, i__904_1467 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__901_1470___1), c__954__auto___1471, cljs_core.Count.X_invoke_Arity1(c__954__auto___1471).(float64), float64(0)
											continue
										}
									} else {
										{
											var mt_1472 = cljs_core.First.X_invoke_Arity1(seq__901_1470___1)
											_ = mt_1472
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
												var G__917 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
												var G__918 = float64(1)
												_, _ = G__917, G__918
												return mt_1472.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__917, G__918)
											}(), func(make_seq *cljs_core.AFn, seq__901_1464 interface{}, chunk__902_1465 interface{}, count__903_1466 float64, i__904_1467 float64, mt_1472 interface{}, seq__901_1470___1 interface{}, temp__4222__auto___1469 cljs_core.CljsCoreISeq) *cljs_core.AFn {
												return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
													if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
														X__GT_t924 = func(__GT_t924 *cljs_core.AFn, seq__901_1464 interface{}, chunk__902_1465 interface{}, count__903_1466 float64, i__904_1467 float64, mt_1472 interface{}, seq__901_1470___1 interface{}, temp__4222__auto___1469 cljs_core.CljsCoreISeq) *cljs_core.AFn {
															return cljs_core.Fn(__GT_t924, 10, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, temp__4222__auto_____1 interface{}, i__904___1 interface{}, count__903___1 interface{}, chunk__902___1 interface{}, seq__901___2 interface{}, test_stuff___1 interface{}, meta925 interface{}) interface{} {
																return (&CljsCore_testT924{from_seq___1, make_seq___1, mt___1, temp__4222__auto_____1, i__904___1, count__903___1, chunk__902___1, seq__901___2, test_stuff___1, meta925})
															})
														}(&cljs_core.AFn{}, seq__901_1464, chunk__902_1465, count__903_1466, i__904_1467, mt_1472, seq__901_1470___1, temp__4222__auto___1469)

														return (&CljsCore_testT924{from_seq, make_seq, mt_1472, temp__4222__auto___1469, i__904_1467, count__903_1466, chunk__902_1465, seq__901_1470___1, test_stuff, nil})
													} else {
														return nil
													}
												})
											}(&cljs_core.AFn{}, seq__901_1464, chunk__902_1465, count__903_1466, i__904_1467, mt_1472, seq__901_1470___1, temp__4222__auto___1469).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))")}))
											}
											seq__901_1464, chunk__902_1465, count__903_1466, i__904_1467 = cljs_core.Next.Arity1IQ(seq__901_1470___1), nil, float64(0), float64(0)
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
				var _STAR_print_length_STAR_929 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_929
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_929

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
				var _STAR_print_length_STAR_930 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_930
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_930

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
				var _STAR_print_length_STAR_931 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_931
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_931

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
				var _STAR_print_length_STAR_932 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_932
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_932

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
				var _STAR_print_length_STAR_933 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_933
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_933

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
				var _STAR_print_length_STAR_934 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_934
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_934

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
							var G__936 = func() interface{} {
								if cljs_core.Value_(value).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
									return cljs_core.Native_get_instance_field.X_invoke_Arity2(value, "Fqn")
								} else {
									return nil
								}
							}()
							_ = G__936
							switch G__936 {
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
				var not_strings_1474 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, nil, float64(1), func(G__1475 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1475, 0, func() interface{} {
						return nil
					})
				}(&cljs_core.AFn{})}, nil})
				_ = not_strings_1474
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1476 *cljs_core.AFn, not_strings_1474 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1476, 1, func(p1__79_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1477 interface{}) {
							defer func() {
								if e937 := recover(); e937 != nil {
									if cljs_core.Value_(e937).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e937
											_ = ___
											return__1477 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e937)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__79_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1474), not_strings_1474) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__79#] (= :failed (try (re-find #\".\" p1__79#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1478 *cljs_core.AFn, not_strings_1474 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1478, 1, func(p1__80_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1479 interface{}) {
							defer func() {
								if e938 := recover(); e938 != nil {
									if cljs_core.Value_(e938).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e938
											_ = ___
											return__1479 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e938)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__80_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1474), not_strings_1474) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__80#] (= :failed (try (re-matches #\".\" p1__80#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1480 *cljs_core.AFn, not_strings_1474 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1480, 1, func(p1__81_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1481 interface{}) {
							defer func() {
								if e939 := recover(); e939 != nil {
									if cljs_core.Value_(e939).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e939
											_ = ___
											return__1481 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e939)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__81_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1474), not_strings_1474) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__81#] (= :failed (try (re-find #\"nomatch\" p1__81#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1482 *cljs_core.AFn, not_strings_1474 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1482, 1, func(p1__82_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1483 interface{}) {
							defer func() {
								if e940 := recover(); e940 != nil {
									if cljs_core.Value_(e940).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e940
											_ = ___
											return__1483 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e940)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__82_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1474), not_strings_1474) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__82#] (= :failed (try (re-matches #\"nomatch\" p1__82#) (catch js/TypeError _ :failed)))) not-strings)")}))
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(func(G__1484 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1484, 1, func(p1__83_SHARP_ interface{}) interface{} {
					return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(p1__83_SHARP_, "ToUpperCase", []interface{}{})
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), "foo")), "FOO") {
			} else {
				panic((&js.Error{("Assert failed: (= (apply str (sequence (map (fn* [p1__83#] (.toUpperCase p1__83#))) \"foo\")) \"FOO\")")}))
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Take_while.X_invoke_Arity1(func(G__1485 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1485, 1, func(p1__84_SHARP_ interface{}) interface{} {
					return (p1__84_SHARP_.(float64) < float64(5))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (take-while (fn* [p1__84#] (< p1__84# 5))) (range 10)) (quote (0 1 2 3 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Drop.X_invoke_Arity1(float64(5)).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (drop 5) (range 10)) (quote (5 6 7 8 9)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Drop_while.X_invoke_Arity1(func(G__1486 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1486, 1, func(p1__85_SHARP_ interface{}) interface{} {
					return (p1__85_SHARP_.(float64) < float64(5))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (drop-while (fn* [p1__85#] (< p1__85# 5))) (range 10)) (quote (5 6 7 8 9)))")}))
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
				var ret_1487 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))
				_ = ret_1487
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_1487)) && (cljs_core.X_EQ_.Arity2IIB(ret_1487, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList))) {
				} else {
					panic((&js.Error{("Assert failed: (and (vector? ret) (= ret (quote (1 2 3))))")}))
				}
			}
			{
				var ret_1488 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				_ = ret_1488
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_1488)) && (cljs_core.X_EQ_.Arity2IIB(ret_1488, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(2), float64(4), float64(6), float64(8)})).(*cljs_core.CljsCoreList))) {
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
			Xform = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Dedupe.X_invoke_Arity0().(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Partition_all.X_invoke_Arity1(float64(3)).(cljs_core.CljsCoreIFn), cljs_core.Partition_by.X_invoke_Arity1(func(G__1489 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1489, 1, func(p1__86_SHARP_ interface{}) interface{} {
					return (cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, p1__86_SHARP_).(float64) < float64(7))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Flatten).(cljs_core.CljsCoreIFn), cljs_core.Random_sample.X_invoke_Arity1(1.0).(cljs_core.CljsCoreIFn), cljs_core.Take_nth.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Keep.X_invoke_Arity1(func(G__1490 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1490, 1, func(p1__87_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__87_SHARP_) {
						return (p1__87_SHARP_.(float64) * p1__87_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Keep_indexed.X_invoke_Arity1(func(G__1491 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1491, 2, func(p1__88_SHARP_ interface{}, p2__89_SHARP_ interface{}) interface{} {
					if cljs_core.Even_QMARK_.Arity1IB(p1__88_SHARP_) {
						return (p1__88_SHARP_.(float64) * p2__89_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Replace.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(2), "two", float64(6), "six", float64(18), "eighteen"}, nil})).(cljs_core.CljsCoreIFn), cljs_core.Take.X_invoke_Arity1(float64(11)).(cljs_core.CljsCoreIFn), cljs_core.Take_while.X_invoke_Arity1(func(G__1492 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1492, 1, func(p1__90_SHARP_ interface{}) interface{} {
					return cljs_core.Not_EQ_.Arity2IIB(float64(300), p1__90_SHARP_)
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Drop.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Drop_while.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Remove.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn)})).(*cljs_core.AFn)

			Data = cljs_core.Vec.X_invoke_Arity1(cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(18)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(20)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(Xform, Data), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(36), float64(200), float64(10)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence xform data) (quote (36 200 10)))")}))
			}
			Xf = cljs_core.Map_.X_invoke_Arity1(func(G__1493 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1493, 2, func(p1__91_SHARP_ interface{}, p2__92_SHARP_ interface{}) interface{} {
					return (p1__91_SHARP_.(float64) + p2__92_SHARP_.(float64))
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

func (this__760__auto__ *CljsCore_testPerson) X_lookup_Arity3(k661 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__664 = func() interface{} {
			if cljs_core.Value_(k661).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k661, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__664
		switch G__664 {
		case "lastname":
			return this__760__auto__.Lastname

		case "firstname":
			return this__760__auto__.Firstname

		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k661, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testPerson) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__1495 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1495, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
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

func (this__765__auto__ *CljsCore_testPerson) X_assoc_Arity3(k__766__auto__ interface{}, G__660 interface{}) interface{} {
	{
		var pred__672 = cljs_core.Keyword_identical_QMARK_
		var expr__673 = k__766__auto__
		_, _ = pred__672, expr__673
		if cljs_core.Truth_(func() interface{} {
			var G__675 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)})
			var G__676 = expr__673
			_, _ = G__675, G__676
			return pred__672.X_invoke_Arity2(G__675, G__676)
		}()) {
			return (&CljsCore_testPerson{G__660, this__765__auto__.Lastname, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__677 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})
				var G__678 = expr__673
				_, _ = G__677, G__678
				return pred__672.X_invoke_Arity2(G__677, G__678)
			}()) {
				return (&CljsCore_testPerson{this__765__auto__.Firstname, G__660, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPerson{this__765__auto__.Firstname, this__765__auto__.Lastname, this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__660), nil})
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

func (this__757__auto__ *CljsCore_testPerson) X_with_meta_Arity2(G__660 interface{}) interface{} {
	return (&CljsCore_testPerson{this__757__auto__.Firstname, this__757__auto__.Lastname, G__660, this__757__auto__.X__extmap, this__757__auto__.X__hash})
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

func (this__760__auto__ *CljsCore_testA) X_lookup_Arity3(k680 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__683 = k680
		_ = G__683
		switch G__683 {
		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k680, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testA) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testA) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__1497 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1497, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
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

func (this__765__auto__ *CljsCore_testA) X_assoc_Arity3(k__766__auto__ interface{}, G__679 interface{}) interface{} {
	{
		var pred__687 = cljs_core.Keyword_identical_QMARK_
		var expr__688 = k__766__auto__
		_, _ = pred__687, expr__688
		return (&CljsCore_testA{this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__679), nil})
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

func (this__757__auto__ *CljsCore_testA) X_with_meta_Arity2(G__679 interface{}) interface{} {
	return (&CljsCore_testA{G__679, this__757__auto__.X__extmap, this__757__auto__.X__hash})
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

func (this__760__auto__ *CljsCore_testC) X_lookup_Arity3(k691 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__694 = func() interface{} {
			if cljs_core.Value_(k691).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k691, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__694
		switch G__694 {
		case "c":
			return this__760__auto__.C

		case "b":
			return this__760__auto__.B

		case "a":
			return this__760__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k691, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testC) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testC) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__1499 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1499, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
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

func (this__765__auto__ *CljsCore_testC) X_assoc_Arity3(k__766__auto__ interface{}, G__690 interface{}) interface{} {
	{
		var pred__704 = cljs_core.Keyword_identical_QMARK_
		var expr__705 = k__766__auto__
		_, _ = pred__704, expr__705
		if cljs_core.Truth_(func() interface{} {
			var G__707 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__708 = expr__705
			_, _ = G__707, G__708
			return pred__704.X_invoke_Arity2(G__707, G__708)
		}()) {
			return (&CljsCore_testC{G__690, this__765__auto__.B, this__765__auto__.C, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__709 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__710 = expr__705
				_, _ = G__709, G__710
				return pred__704.X_invoke_Arity2(G__709, G__710)
			}()) {
				return (&CljsCore_testC{this__765__auto__.A, G__690, this__765__auto__.C, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
			} else {
				if cljs_core.Truth_(func() interface{} {
					var G__711 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})
					var G__712 = expr__705
					_, _ = G__711, G__712
					return pred__704.X_invoke_Arity2(G__711, G__712)
				}()) {
					return (&CljsCore_testC{this__765__auto__.A, this__765__auto__.B, G__690, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
				} else {
					return (&CljsCore_testC{this__765__auto__.A, this__765__auto__.B, this__765__auto__.C, this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__690), nil})
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

func (this__757__auto__ *CljsCore_testC) X_with_meta_Arity2(G__690 interface{}) interface{} {
	return (&CljsCore_testC{this__757__auto__.A, this__757__auto__.B, this__757__auto__.C, G__690, this__757__auto__.X__extmap, this__757__auto__.X__hash})
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

func (this__760__auto__ *CljsCore_testA2) X_lookup_Arity3(k716 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__719 = func() interface{} {
			if cljs_core.Value_(k716).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k716, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__719
		switch G__719 {
		case "x":
			return this__760__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k716, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testA2) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__1501 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1501, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
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

func (this__765__auto__ *CljsCore_testA2) X_assoc_Arity3(k__766__auto__ interface{}, G__715 interface{}) interface{} {
	{
		var pred__725 = cljs_core.Keyword_identical_QMARK_
		var expr__726 = k__766__auto__
		_, _ = pred__725, expr__726
		if cljs_core.Truth_(func() interface{} {
			var G__728 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__729 = expr__726
			_, _ = G__728, G__729
			return pred__725.X_invoke_Arity2(G__728, G__729)
		}()) {
			return (&CljsCore_testA2{G__715, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testA2{this__765__auto__.X, this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__715), nil})
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

func (this__757__auto__ *CljsCore_testA2) X_with_meta_Arity2(G__715 interface{}) interface{} {
	return (&CljsCore_testA2{this__757__auto__.X, G__715, this__757__auto__.X__extmap, this__757__auto__.X__hash})
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

func (this__760__auto__ *CljsCore_testB) X_lookup_Arity3(k731 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__734 = func() interface{} {
			if cljs_core.Value_(k731).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k731, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__734
		switch G__734 {
		case "x":
			return this__760__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k731, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testB) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testB) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__1503 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1503, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
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

func (this__765__auto__ *CljsCore_testB) X_assoc_Arity3(k__766__auto__ interface{}, G__730 interface{}) interface{} {
	{
		var pred__740 = cljs_core.Keyword_identical_QMARK_
		var expr__741 = k__766__auto__
		_, _ = pred__740, expr__741
		if cljs_core.Truth_(func() interface{} {
			var G__743 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__744 = expr__741
			_, _ = G__743, G__744
			return pred__740.X_invoke_Arity2(G__743, G__744)
		}()) {
			return (&CljsCore_testB{G__730, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testB{this__765__auto__.X, this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__730), nil})
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

func (this__757__auto__ *CljsCore_testB) X_with_meta_Arity2(G__730 interface{}) interface{} {
	return (&CljsCore_testB{this__757__auto__.X, G__730, this__757__auto__.X__extmap, this__757__auto__.X__hash})
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

type CljsCore_testT745 struct {
	Test_stuff interface{}
	Meta746    interface{}
}

func (_ *CljsCore_testT745) CljsCore_testIFoo__() {}

func (this *CljsCore_testT745) Foo_Arity1() interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
}

func (_ *CljsCore_testT745) CljsCoreIMeta__() {}

func (_747 *CljsCore_testT745) X_meta_Arity1() interface{} {
	return _747.Meta746
}

func (_ *CljsCore_testT745) CljsCoreIWithMeta__() {}

func (_747 *CljsCore_testT745) X_with_meta_Arity2(meta746___1 interface{}) interface{} {
	return (&CljsCore_testT745{_747.Test_stuff, meta746___1})
}

var X__GT_t745 *cljs_core.AFn

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
		var return__1504 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})
		___.A = return__1504
		return return__1504
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

func (___ *CljsCore_testFirst) X_find_first_Arity2(p__748 interface{}) interface{} {
	{
		var vec__750 = p__748
		var x = cljs_core.Nth.X_invoke_Arity3(vec__750, float64(0), nil)
		_, _ = vec__750, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCore_testIHasFirst__() {}

func (p__751 *CljsCore_testFirst) X_get_first_Arity1() interface{} {
	{
		var vec__753 = p__751
		var x = cljs_core.Nth.X_invoke_Arity3(vec__753, float64(0), nil)
		_, _ = vec__753, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCoreObject__() {}

func (p__754 *CljsCore_testFirst) ToString() string {
	{
		var vec__756 = p__754
		var x = cljs_core.Nth.X_invoke_Arity3(vec__756, float64(0), nil)
		_, _ = vec__756, x
		return (`` + cljs_core.Str.X_invoke_Arity1(x).(string))
	}
}

func (this *CljsCore_testFirst) String() string {
	return this.ToString()
}

func (_ *CljsCore_testFirst) CljsCoreIFn__() {}

func (p__757 *CljsCore_testFirst) X_invoke_Arity0() interface{} {
	{
		var vec__759 = p__757
		var x = cljs_core.Nth.X_invoke_Arity3(vec__759, float64(0), nil)
		_, _ = vec__759, x
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

func (___ *CljsCore_testDestructuringWithLocals) X_find_first_Arity2(p__761 interface{}) interface{} {
	{
		var vec__763 = p__761
		var x = cljs_core.Nth.X_invoke_Arity3(vec__763, float64(0), nil)
		var y = cljs_core.Nth.X_invoke_Arity3(vec__763, float64(1), nil)
		_, _, _ = vec__763, x, y
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

func (this__760__auto__ *CljsCore_testPrintMe) X_lookup_Arity3(k775 interface{}, else__761__auto__ interface{}) interface{} {
	{
		var G__778 = func() interface{} {
			if cljs_core.Value_(k775).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k775, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__778
		switch G__778 {
		case "b":
			return this__760__auto__.B

		case "a":
			return this__760__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__760__auto__.X__extmap, k775, else__761__auto__)

		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIPrintWithWriter__() {}

func (this__774__auto__ *CljsCore_testPrintMe) X_pr_writer_Arity3(writer__775__auto__ interface{}, opts__776__auto__ interface{}) interface{} {
	{
		var pr_pair__777__auto__ = func(G__1506 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1506, 3, func(keyval__778__auto__ interface{}, ___779__auto__ interface{}, ___779__auto_____1 interface{}) interface{} {
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

func (this__765__auto__ *CljsCore_testPrintMe) X_assoc_Arity3(k__766__auto__ interface{}, G__774 interface{}) interface{} {
	{
		var pred__786 = cljs_core.Keyword_identical_QMARK_
		var expr__787 = k__766__auto__
		_, _ = pred__786, expr__787
		if cljs_core.Truth_(func() interface{} {
			var G__789 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__790 = expr__787
			_, _ = G__789, G__790
			return pred__786.X_invoke_Arity2(G__789, G__790)
		}()) {
			return (&CljsCore_testPrintMe{G__774, this__765__auto__.B, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__791 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__792 = expr__787
				_, _ = G__791, G__792
				return pred__786.X_invoke_Arity2(G__791, G__792)
			}()) {
				return (&CljsCore_testPrintMe{this__765__auto__.A, G__774, this__765__auto__.X__meta, this__765__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPrintMe{this__765__auto__.A, this__765__auto__.B, this__765__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__765__auto__.X__extmap, k__766__auto__, G__774), nil})
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

func (this__757__auto__ *CljsCore_testPrintMe) X_with_meta_Arity2(G__774 interface{}) interface{} {
	return (&CljsCore_testPrintMe{this__757__auto__.A, this__757__auto__.B, G__774, this__757__auto__.X__extmap, this__757__auto__.X__hash})
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

type CljsCore_testT826 struct {
	F          interface{}
	Baz        interface{}
	Test_stuff interface{}
	Meta827    interface{}
}

func (_ *CljsCore_testT826) CljsCore_testIBar__() {}

func (___ *CljsCore_testT826) X_bar_Arity2(x interface{}) interface{} {
	{
		var G__830 = x
		_ = G__830
		return ___.F.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__830)
	}
}

func (_ *CljsCore_testT826) CljsCoreIMeta__() {}

func (_828 *CljsCore_testT826) X_meta_Arity1() interface{} {
	return _828.Meta827
}

func (_ *CljsCore_testT826) CljsCoreIWithMeta__() {}

func (_828 *CljsCore_testT826) X_with_meta_Arity2(meta827___1 interface{}) interface{} {
	return (&CljsCore_testT826{_828.F, _828.Baz, _828.Test_stuff, meta827___1})
}

var X__GT_t826 *cljs_core.AFn

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

type CljsCore_testT839 struct {
	Test_stuff interface{}
	Meta840    interface{}
}

func (_ *CljsCore_testT839) CljsCoreIHash__() {}

func (___ *CljsCore_testT839) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT839) CljsCoreIMeta__() {}

func (_841 *CljsCore_testT839) X_meta_Arity1() interface{} {
	return _841.Meta840
}

func (_ *CljsCore_testT839) CljsCoreIWithMeta__() {}

func (_841 *CljsCore_testT839) X_with_meta_Arity2(meta840___1 interface{}) interface{} {
	return (&CljsCore_testT839{_841.Test_stuff, meta840___1})
}

var X__GT_t839 *cljs_core.AFn

type CljsCore_testT842 struct {
	A          interface{}
	Test_stuff interface{}
	Meta843    interface{}
}

func (_ *CljsCore_testT842) CljsCoreIHash__() {}

func (___ *CljsCore_testT842) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT842) CljsCoreIMeta__() {}

func (_844 *CljsCore_testT842) X_meta_Arity1() interface{} {
	return _844.Meta843
}

func (_ *CljsCore_testT842) CljsCoreIWithMeta__() {}

func (_844 *CljsCore_testT842) X_with_meta_Arity2(meta843___1 interface{}) interface{} {
	return (&CljsCore_testT842{_844.A, _844.Test_stuff, meta843___1})
}

var X__GT_t842 *cljs_core.AFn

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

type CljsCore_testT912 struct {
	From_seq   interface{}
	Make_seq   interface{}
	Mt         interface{}
	I__904     interface{}
	Count__903 interface{}
	Chunk__902 interface{}
	Seq__901   interface{}
	Test_stuff interface{}
	Meta913    interface{}
}

func (_ *CljsCore_testT912) CljsCoreISeq__() {}

func (this *CljsCore_testT912) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT912) X_rest_Arity1() interface{} {
	{
		var G__916 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__916
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__916)
	}
}

func (_ *CljsCore_testT912) CljsCoreISeqable__() {}

func (this *CljsCore_testT912) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT912) CljsCoreIMeta__() {}

func (_914 *CljsCore_testT912) X_meta_Arity1() interface{} {
	return _914.Meta913
}

func (_ *CljsCore_testT912) CljsCoreIWithMeta__() {}

func (_914 *CljsCore_testT912) X_with_meta_Arity2(meta913___1 interface{}) interface{} {
	return (&CljsCore_testT912{_914.From_seq, _914.Make_seq, _914.Mt, _914.I__904, _914.Count__903, _914.Chunk__902, _914.Seq__901, _914.Test_stuff, meta913___1})
}

var X__GT_t912 *cljs_core.AFn

type CljsCore_testT924 struct {
	From_seq           interface{}
	Make_seq           interface{}
	Mt                 interface{}
	Temp__4222__auto__ interface{}
	I__904             interface{}
	Count__903         interface{}
	Chunk__902         interface{}
	Seq__901           interface{}
	Test_stuff         interface{}
	Meta925            interface{}
}

func (_ *CljsCore_testT924) CljsCoreISeq__() {}

func (this *CljsCore_testT924) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT924) X_rest_Arity1() interface{} {
	{
		var G__928 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__928
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__928)
	}
}

func (_ *CljsCore_testT924) CljsCoreISeqable__() {}

func (this *CljsCore_testT924) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT924) CljsCoreIMeta__() {}

func (_926 *CljsCore_testT924) X_meta_Arity1() interface{} {
	return _926.Meta925
}

func (_ *CljsCore_testT924) CljsCoreIWithMeta__() {}

func (_926 *CljsCore_testT924) X_with_meta_Arity2(meta925___1 interface{}) interface{} {
	return (&CljsCore_testT924{_926.From_seq, _926.Make_seq, _926.Mt, _926.Temp__4222__auto__, _926.I__904, _926.Count__903, _926.Chunk__902, _926.Seq__901, _926.Test_stuff, meta925___1})
}

var X__GT_t924 *cljs_core.AFn

var Case_recur *cljs_core.AFn

var Xform *cljs_core.AFn

var Data interface{}

var Xf *cljs_core.AFn

func Test_runner(t *testing.T) {
	Test_stuff.X_invoke_Arity0()
}
