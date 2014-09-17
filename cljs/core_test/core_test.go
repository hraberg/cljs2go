// Compiled by ClojureScript to Go 0.0-2322
// cljs.core-test

package core_test

import (
	"math"
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
						i, j = (i + float64(1)), cljs_core.Conj.X_invoke_Arity2(j, func(G__885 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__885, 0, func() interface{} {
								return i
							})
						}(&cljs_core.AFn{}, i, j))
						continue
					} else {
						return cljs_core.Map_.X_invoke_Arity2(func(G__886 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__886, 1, func(p1__56_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})}, nil}), cljs_core.Map_.X_invoke_Arity2(func(G__887 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__887, 1, func(p1__57_SHARP_ interface{}) interface{} {
					{
						return p1__57_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
					}
				})
			}(&cljs_core.AFn{}), func() *cljs_core.CljsCoreLazySeq {
				var iter__916__auto__ = func(iter__489 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(iter__489, 1, func(s__490 interface{}) interface{} {
						return (&cljs_core.CljsCoreLazySeq{nil, func(G__888 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__888, 0, func() interface{} {
								{
									var s__490___1 interface{} = s__490
									_ = s__490___1
									for {
										{
											var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__490___1)
											_ = temp__4222__auto__
											if cljs_core.Truth_(temp__4222__auto__) {
												{
													var xs__4752__auto__ = temp__4222__auto__
													_ = xs__4752__auto__
													{
														var i = cljs_core.First.X_invoke_Arity1(xs__4752__auto__)
														_ = i
														{
															var iterys__912__auto__ = func(iter__491 *cljs_core.AFn, s__490___1 interface{}, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																return cljs_core.Fn(iter__491, 1, func(s__492 interface{}) interface{} {
																	return (&cljs_core.CljsCoreLazySeq{nil, func(G__889 *cljs_core.AFn, s__490___1 interface{}, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																		return cljs_core.Fn(G__889, 0, func() interface{} {
																			{
																				var s__492___1 interface{} = s__492
																				_ = s__492___1
																				for {
																					{
																						var temp__4222__auto_____1 = cljs_core.Seq.Arity1IQ(s__492___1)
																						_ = temp__4222__auto_____1
																						if cljs_core.Truth_(temp__4222__auto_____1) {
																							{
																								var s__492___2 = temp__4222__auto_____1
																								_ = s__492___2
																								if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__492___2) {
																									{
																										var c__914__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__492___2)
																										var size__915__auto__ = cljs_core.Count.X_invoke_Arity1(c__914__auto__).(float64)
																										var b__494 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__915__auto__).(*cljs_core.CljsCoreChunkBuffer)
																										_, _, _ = c__914__auto__, size__915__auto__, b__494
																										if func() bool {
																											var i__493 = float64(0)
																											_ = i__493
																											for {
																												if i__493 < size__915__auto__ {
																													{
																														var j = c__914__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__493)
																														_ = j
																														cljs_core.Chunk_append.X_invoke_Arity2(b__494, func(G__890 *cljs_core.AFn, i__493 float64, s__490___1 interface{}, j interface{}, c__914__auto__ interface{}, size__915__auto__ float64, b__494 *cljs_core.CljsCoreChunkBuffer, s__492___2 interface{}, temp__4222__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																															return cljs_core.Fn(G__890, 0, func() interface{} {
																																return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																															})
																														}(&cljs_core.AFn{}, i__493, s__490___1, j, c__914__auto__, size__915__auto__, b__494, s__492___2, temp__4222__auto_____1, i, xs__4752__auto__, temp__4222__auto__))
																														i__493 = (i__493 + float64(1))
																														continue
																													}
																												} else {
																													return true
																												}
																											}
																										}() {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__494), iter__491.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__492___2)).(*cljs_core.CljsCoreLazySeq))
																										} else {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__494), nil)
																										}
																									}
																								} else {
																									{
																										var j = cljs_core.First.X_invoke_Arity1(s__492___2)
																										_ = j
																										return cljs_core.Cons.X_invoke_Arity2(func(G__891 *cljs_core.AFn, s__490___1 interface{}, j interface{}, s__492___2 interface{}, temp__4222__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4752__auto__ cljs_core.CljsCoreISeq, temp__4222__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																											return cljs_core.Fn(G__891, 0, func() interface{} {
																												return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																											})
																										}(&cljs_core.AFn{}, s__490___1, j, s__492___2, temp__4222__auto_____1, i, xs__4752__auto__, temp__4222__auto__), iter__491.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__492___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
																	}(&cljs_core.AFn{}, s__490___1, i, xs__4752__auto__, temp__4222__auto__), nil, nil})
																})
															}(&cljs_core.AFn{}, s__490___1, i, xs__4752__auto__, temp__4222__auto__)
															var fs__913__auto__ = cljs_core.Seq.Arity1IQ(iterys__912__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq))
															_, _ = iterys__912__auto__, fs__913__auto__
															if cljs_core.Truth_(fs__913__auto__) {
																return cljs_core.Concat.X_invoke_Arity2(fs__913__auto__, iter__489.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__490___1)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
															} else {
																s__490___1 = cljs_core.Rest.Arity1IQ(s__490___1)
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
				_ = iter__916__auto__
				return iter__916__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})).(*cljs_core.CljsCoreLazySeq)
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
			if cljs_core.X_EQ_.Arity2IIB(float64(42), math.Trunc(42.5)) {
			} else {
				panic((&js.Error{("Assert failed: (= 42 (int 42.5))")}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(math.Trunc(42.5)) {
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
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), math.Trunc(-1.5)) {
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
				var sb__1117__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1117__auto__
				{
					var _STAR_print_fn_STAR_500_892 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_500_892
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_500_892

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__893 *cljs_core.AFn, _STAR_print_fn_STAR_500_892 interface{}, sb__1117__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__893, 1, func(x__1118__auto__ interface{}) interface{} {
									return sb__1117__auto__.Append(x__1118__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_500_892, sb__1117__auto__)

							cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)}))
							cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)}))
						}
					}()
				}
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1117__auto__).(string))
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= \"12\" (with-out-str (print 1) (print 2)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("12", func() string {
				var sb__1117__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1117__auto__
				{
					var _STAR_print_fn_STAR_501_894 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_501_894
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_501_894

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__895 *cljs_core.AFn, _STAR_print_fn_STAR_501_894 interface{}, sb__1117__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__895, 1, func(x__1118__auto__ interface{}) interface{} {
									return sb__1117__auto__.Append(x__1118__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_501_894, sb__1117__auto__)

							cljs_core.X_STAR_print_fn_STAR_.X_invoke_Arity1(float64(1))
							cljs_core.X_STAR_print_fn_STAR_.X_invoke_Arity1(float64(2))
						}
					}()
				}
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1117__auto__).(string))
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
				var jumble = func(G__896 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__896, 2, func(a interface{}, b interface{}) interface{} {
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
				var jumble = func(G__897 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__897, 2, func(a interface{}, b interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Apply.X_invoke_Arity2(func(G__898 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__898, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return ((cljs_core.Nth.X_invoke_Arity2(args, float64(0)).(float64) + cljs_core.Nth.X_invoke_Arity2(args, float64(1)).(float64)) + cljs_core.Nth.X_invoke_Arity2(args, float64(2)).(float64))
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))) {
			} else {
				panic((&js.Error{("Assert failed: (= 3 (apply (fn [& args] (+ (nth args 0) (nth args 1) (nth args 2))) (iterate inc 0)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__899 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__899, 0, func(m__ ...interface{}) interface{} {
					var m = cljs_core.Seq.Arity1IQ(m__[0])
					_ = m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [0 1 2 3 4] (take 5 (apply (fn [& m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__900 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__900, 1, func(x_m__ ...interface{}) interface{} {
					var x = x_m__[0]
					var m = cljs_core.Seq.Arity1IQ(x_m__[1])
					_, _ = x, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (take 5 (apply (fn [x & m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__901 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__901, 2, func(x_y_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5), float64(6), float64(7)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__902 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__902, 3, func(x_y_z_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6), float64(7), float64(8)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__903 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__903, 4, func(x_y_z_a_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__904 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__904, 5, func(x_y_z_a_b_m__ ...interface{}) interface{} {
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
				var single_arity_non_variadic_905 = func(G__910 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__910, 3, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{})
				var multiple_arity_non_variadic_906 = func(G__911 *cljs_core.AFn, single_arity_non_variadic_905 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__911, 3, func(x interface{}) interface{} {
						return x
					}, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{y, x}, nil})
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_905)
				var single_arity_variadic_fixedargs_907 = func(G__912 *cljs_core.AFn, single_arity_non_variadic_905 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_906 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__912, 2, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{more, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_905, multiple_arity_non_variadic_906)
				var single_arity_variadic_nofixedargs_908 = func(G__913 *cljs_core.AFn, single_arity_non_variadic_905 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_906 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_907 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__913, 0, func(more__ ...interface{}) interface{} {
						var more = cljs_core.Seq.Arity1IQ(more__[0])
						_ = more
						return more
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_905, multiple_arity_non_variadic_906, single_arity_variadic_fixedargs_907)
				var multiple_arity_variadic_909 = func(G__914 *cljs_core.AFn, single_arity_non_variadic_905 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_906 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_907 cljs_core.CljsCoreIFn, single_arity_variadic_nofixedargs_908 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__914, 2, func(x interface{}) interface{} {
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
				}(&cljs_core.AFn{}, single_arity_non_variadic_905, multiple_arity_non_variadic_906, single_arity_variadic_fixedargs_907, single_arity_variadic_nofixedargs_908)
				_, _, _, _, _ = single_arity_non_variadic_905, multiple_arity_non_variadic_906, single_arity_variadic_fixedargs_907, single_arity_variadic_nofixedargs_908, multiple_arity_variadic_909
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_non_variadic_905, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic [1 2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_non_variadic_905, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic 1 [2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_non_variadic_905, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic 1 2 [3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_906, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (apply multiple-arity-non-variadic [42]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_906, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-non-variadic [1 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_906, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-non-variadic 1 [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_906, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic [1 2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_906, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic 1 [2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_non_variadic_906, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic 1 2 [3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_907, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_fixedargs_907, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_fixedargs_907, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_fixedargs_907, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_fixedargs_907, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_907, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 4 5] (take 3 (first (apply single-arity-variadic-fixedargs (iterate inc 1)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_907, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (rest (apply single-arity-variadic-fixedargs (iterate inc 1))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_908, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_nofixedargs_908, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_nofixedargs_908, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_nofixedargs_908, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_nofixedargs_908, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_908, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (take 5 (apply single-arity-variadic-nofixedargs (iterate inc 1))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_909, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (apply multiple-arity-variadic [42]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_909, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-variadic [1 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_909, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-variadic 1 [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_909, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_909, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_variadic_909, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(multiple_arity_variadic_909, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(multiple_arity_variadic_909, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_909, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 4 5] (take 3 (first (apply multiple-arity-variadic (iterate inc 1)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_909, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (rest (apply multiple-arity-variadic (iterate inc 1))))")}))
				}
			}
			{
				var f1_915 = func(f1 *cljs_core.AFn) *cljs_core.AFn {
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
				var f2_916 = func(f2 *cljs_core.AFn, f1_915 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(f2, 2, func(x interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					}, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return cljs_core.Apply.X_invoke_Arity3(f1_915, y, more)
					})
				}(&cljs_core.AFn{}, f1_915)
				_, _ = f1_915, f2_916
				if cljs_core.X_EQ_.Arity2IIB(float64(1), f2_916.X_invoke_Arity2(float64(1), float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (f2 1 2))")}))
				}
			}
			{
				var f_917 = func(G__918 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__918, 1, func() interface{} {
						return nil
					}, func(a_more__ ...interface{}) interface{} {
						var a = a_more__[0]
						var more = cljs_core.Seq.Arity1IQ(a_more__[1])
						_, _ = a, more
						return more
					})
				}(&cljs_core.AFn{})
				_ = f_917
				if cljs_core.Nil_(f_917.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (f :foo))")}))
				}
			}
			if cljs_core.Nil_(cljs_core.Array_seq.X_invoke_Arity2([]interface{}{float64(1)}, float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (array-seq (array 1) 1))")}))
			}
			{
				var f_919 = func(G__922 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__922, 1, func(x interface{}) interface{} {
						return (x.(float64) * float64(2))
					})
				}(&cljs_core.AFn{})
				var m_920 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar"}, nil})
				var mf_921 = cljs_core.With_meta.X_invoke_Arity2(f_919, m_920)
				_, _, _ = f_919, m_920, mf_921
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(f_919)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (meta f))")}))
				}
				if cljs_core.Fn_QMARK_.Arity1IB(mf_921) {
				} else {
					panic((&js.Error{("Assert failed: (fn? mf)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__502 = float64(2)
					_ = G__502
					return mf_921.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__502)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (mf 2))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Apply.X_invoke_Arity2(mf_921, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (apply mf [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(mf_921), m_920) {
				} else {
					panic((&js.Error{("Assert failed: (= (meta mf) m)")}))
				}
			}
			{
				var a_923 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_923
				if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Deref.X_invoke_Arity1(a_923)) {
				} else {
					panic((&js.Error{("Assert failed: (= 0 (deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity2(a_923, cljs_core.Inc)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (swap! a inc))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(false, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_923, float64(0), float64(42)))) {
				} else {
					panic((&js.Error{("Assert failed: (= false (compare-and-set! a 0 42))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(true, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_923, float64(1), float64(7)))) {
				} else {
					panic((&js.Error{("Assert failed: (= true (compare-and-set! a 1 7))")}))
				}
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(a_923)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (meta a))")}))
				}
				if cljs_core.Nil_(cljs_core.Get_validator.X_invoke_Arity1(a_923)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (get-validator a))")}))
				}
			}
			{
				var a_924 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_924
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity3(a_924, cljs_core.X_PLUS_, float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (swap! a + 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Swap_BANG_.X_invoke_Arity4(a_924, cljs_core.X_PLUS_, float64(1), float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (swap! a + 1 2))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(10), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_924, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 10 (swap! a + 1 2 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(20), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_924, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(4)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 20 (swap! a + 1 2 3 4))")}))
				}
			}
			{
				var a_925 = cljs_core.Atom.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "validator", Fqn: "validator", X_hash: float64(-1966190681)}), cljs_core.Coll_QMARK_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})})).(*cljs_core.CljsCoreAtom)
				_ = a_925
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Coll_QMARK_, cljs_core.Get_validator.X_invoke_Arity1(a_925)) {
				} else {
					panic((&js.Error{("Assert failed: (= coll? (get-validator a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_925)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:a 1} (meta a))")}))
				}
				cljs_core.Alter_meta_BANG_.X_invoke_ArityVariadic(a_925, cljs_core.Assoc, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}))
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_925)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:a 1, :b 2} (meta a))")}))
				}
			}
			if cljs_core.Nil_(cljs_core.Empty.X_invoke_Arity1(nil)) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (empty nil))")}))
			}
			{
				var e_lazy_seq_926 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCoreLazySeq{nil, func(G__927 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__927, 0, func() interface{} {
						return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons)
					})
				}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_lazy_seq_926
				if cljs_core.Seq_QMARK_.Arity1IB(e_lazy_seq_926) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-lazy-seq)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_lazy_seq_926) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-lazy-seq)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_lazy_seq_926)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-lazy-seq))")}))
				}
			}
			{
				var e_list_928 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_list_928
				if cljs_core.Seq_QMARK_.Arity1IB(e_list_928) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-list)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_list_928) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-list)")}))
				}
			}
			{
				var e_elist_929 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_elist_929
				if cljs_core.Seq_QMARK_.Arity1IB(e_elist_929) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-elist)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_elist_929) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-elist)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), cljs_core.Get.X_invoke_Arity2(cljs_core.Meta.X_invoke_Arity1(e_elist_929), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= :c (get (meta e-elist) :b))")}))
				}
			}
			{
				var e_cons_930 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_cons_930
				if cljs_core.Seq_QMARK_.Arity1IB(e_cons_930) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-cons)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_cons_930) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-cons)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_cons_930)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-cons))")}))
				}
			}
			{
				var e_vec_931 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_vec_931
				if cljs_core.Vector_QMARK_.Arity1IB(e_vec_931) {
				} else {
					panic((&js.Error{("Assert failed: (vector? e-vec)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_vec_931) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-vec)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_vec_931)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-vec))")}))
				}
			}
			{
				var e_omap_932 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_omap_932
				if cljs_core.Map_QMARK_.Arity1IB(e_omap_932) {
				} else {
					panic((&js.Error{("Assert failed: (map? e-omap)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_omap_932) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-omap)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_omap_932)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-omap))")}))
				}
			}
			{
				var e_hmap_933 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_hmap_933
				if cljs_core.Map_QMARK_.Arity1IB(e_hmap_933) {
				} else {
					panic((&js.Error{("Assert failed: (map? e-hmap)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_hmap_933) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-hmap)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_hmap_933)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-hmap))")}))
				}
			}
			{
				var a_934 = cljs_core.Atom.X_invoke_Arity1(nil).(*cljs_core.CljsCoreAtom)
				_ = a_934
				if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (try 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__935 interface{}) {
					defer func() {
						if e503 := recover(); e503 != nil {
							if func() bool { _, instanceof := e503.(*js.Error); return instanceof }() {
								{
									var e = e503
									_ = e
									return__935 = float64(2)
								}
							} else {
								panic(e503)

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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__936 interface{}) {
					defer func() {
						if e504 := recover(); e504 != nil {
							if func() bool { _, instanceof := e504.(*js.Error); return instanceof }() {
								{
									var e = e504
									_ = e
									return__936 = float64(2)
								}
							} else {
								panic(e504)

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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__937 interface{}) {
					defer func() {
						if e505 := recover(); e505 != nil {
							if func() bool { _, instanceof := e505.(*js.Error); return instanceof }() {
								{
									var e = e505
									_ = e
									return__937 = float64(2)
								}
							} else {
								{
									var e = e505
									_ = e
									return__937 = float64(3)
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
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__938 interface{}) {
					defer func() {
						if e506 := recover(); e506 != nil {
							if func() bool { _, instanceof := e506.(*js.Error); return instanceof }() {
								{
									var e = e506
									_ = e
									return__938 = float64(2)
								}
							} else {
								{
									var e = e506
									_ = e
									return__938 = float64(3)
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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__939 interface{}) {
					defer func() {
						if e507 := recover(); e507 != nil {
							if func() bool { _, instanceof := e507.(*js.Error); return instanceof }() {
								{
									var e = e507
									_ = e
									return__939 = float64(3)
								}
							} else {
								{
									var e = e507
									_ = e
									return__939 = e
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
						cljs_core.Reset_BANG_.X_invoke_Arity2(a_934, float64(42))
					}()
					{
						return float64(1)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (try 1 (finally (reset! a 42))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Deref.X_invoke_Arity1(a_934)) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (deref a))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), cljs_core.Seq_(cljs_core.Nthnext.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)))) {
			} else {
				panic((&js.Error{("Assert failed: (= [3] (nthnext [1 2 3] 2))")}))
			}
			{
				var v_940 = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})
				_ = v_940
				if cljs_core.X_EQ_.Arity2IIB(v_940, func() *cljs_core.CljsCoreLazySeq {
					var iter__916__auto__ = func(iter__508 *cljs_core.AFn, v_940 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__508, 1, func(s__509 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__941 *cljs_core.AFn, v_940 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__941, 0, func() interface{} {
									{
										var s__509___1 interface{} = s__509
										_ = s__509___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__509___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__509___2 = temp__4222__auto__
														_ = s__509___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__509___2) {
															{
																var c__914__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__509___2)
																var size__915__auto__ = cljs_core.Count.X_invoke_Arity1(c__914__auto__).(float64)
																var b__511 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__915__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__914__auto__, size__915__auto__, b__511
																if func() bool {
																	var i__510 = float64(0)
																	_ = i__510
																	for {
																		if i__510 < size__915__auto__ {
																			{
																				var e = c__914__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__510)
																				_ = e
																				cljs_core.Chunk_append.X_invoke_Arity2(b__511, e)
																				i__510 = (i__510 + float64(1))
																				continue
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__511), iter__508.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__509___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__511), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__509___2)
																_ = e
																return cljs_core.Cons.X_invoke_Arity2(e, iter__508.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__509___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_940), nil, nil})
						})
					}(&cljs_core.AFn{}, v_940)
					_ = iter__916__auto__
					return iter__916__auto__.X_invoke_Arity1(v_940).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= v (for [e v] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(9)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__916__auto__ = func(iter__514 *cljs_core.AFn, v_940 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__514, 1, func(s__515 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__942 *cljs_core.AFn, v_940 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__942, 0, func() interface{} {
									{
										var s__515___1 interface{} = s__515
										_ = s__515___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__515___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__515___2 = temp__4222__auto__
														_ = s__515___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__515___2) {
															{
																var c__914__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__515___2)
																var size__915__auto__ = cljs_core.Count.X_invoke_Arity1(c__914__auto__).(float64)
																var b__517 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__915__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__914__auto__, size__915__auto__, b__517
																if func() bool {
																	var i__516 = float64(0)
																	_ = i__516
																	for {
																		if i__516 < size__915__auto__ {
																			{
																				var e = c__914__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__516)
																				_ = e
																				{
																					var m = (e.(float64) * e.(float64))
																					_ = m
																					cljs_core.Chunk_append.X_invoke_Arity2(b__517, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																					i__516 = (i__516 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__517), iter__514.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__515___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__517), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__515___2)
																_ = e
																{
																	var m = (e.(float64) * e.(float64))
																	_ = m
																	return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__514.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__515___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_940), nil, nil})
						})
					}(&cljs_core.AFn{}, v_940)
					_ = iter__916__auto__
					return iter__916__auto__.X_invoke_Arity1(v_940).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [[1 1] [2 4] [3 9]] (for [e v :let [m (* e e)]] [e m]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__916__auto__ = func(iter__520 *cljs_core.AFn, v_940 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__520, 1, func(s__521 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__943 *cljs_core.AFn, v_940 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__943, 0, func() interface{} {
									{
										var s__521___1 interface{} = s__521
										_ = s__521___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__521___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__521___2 = temp__4222__auto__
														_ = s__521___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__521___2) {
															{
																var c__914__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__521___2)
																var size__915__auto__ = cljs_core.Count.X_invoke_Arity1(c__914__auto__).(float64)
																var b__523 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__915__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__914__auto__, size__915__auto__, b__523
																if cljs_core.Truth_(func() interface{} {
																	var i__522 = float64(0)
																	_ = i__522
																	for {
																		if i__522 < size__915__auto__ {
																			{
																				var e = c__914__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__522)
																				_ = e
																				if e.(float64) < float64(3) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__523, e)
																					i__522 = (i__522 + float64(1))
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
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__523), iter__520.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__521___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__523), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__521___2)
																_ = e
																if e.(float64) < float64(3) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__520.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__521___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_940), nil, nil})
						})
					}(&cljs_core.AFn{}, v_940)
					_ = iter__916__auto__
					return iter__916__auto__.X_invoke_Arity1(v_940).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2] (for [e v :while (< e 3)] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__916__auto__ = func(iter__526 *cljs_core.AFn, v_940 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__526, 1, func(s__527 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__944 *cljs_core.AFn, v_940 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__944, 0, func() interface{} {
									{
										var s__527___1 interface{} = s__527
										_ = s__527___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__527___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__527___2 = temp__4222__auto__
														_ = s__527___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__527___2) {
															{
																var c__914__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__527___2)
																var size__915__auto__ = cljs_core.Count.X_invoke_Arity1(c__914__auto__).(float64)
																var b__529 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__915__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__914__auto__, size__915__auto__, b__529
																if func() bool {
																	var i__528 = float64(0)
																	_ = i__528
																	for {
																		if i__528 < size__915__auto__ {
																			{
																				var e = c__914__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__528)
																				_ = e
																				if e.(float64) > float64(2) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__529, e)
																					i__528 = (i__528 + float64(1))
																					continue
																				} else {
																					i__528 = (i__528 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__529), iter__526.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__527___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__529), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__527___2)
																_ = e
																if e.(float64) > float64(2) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__526.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__527___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
																} else {
																	s__527___1 = cljs_core.Rest.Arity1IQ(s__527___2)
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
							}(&cljs_core.AFn{}, v_940), nil, nil})
						})
					}(&cljs_core.AFn{}, v_940)
					_ = iter__916__auto__
					return iter__916__auto__.X_invoke_Arity1(v_940).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [3] (for [e v :when (> e 2)] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__916__auto__ = func(iter__532 *cljs_core.AFn, v_940 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__532, 1, func(s__533 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__945 *cljs_core.AFn, v_940 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__945, 0, func() interface{} {
									{
										var s__533___1 interface{} = s__533
										_ = s__533___1
										for {
											{
												var temp__4222__auto__ = cljs_core.Seq.Arity1IQ(s__533___1)
												_ = temp__4222__auto__
												if cljs_core.Truth_(temp__4222__auto__) {
													{
														var s__533___2 = temp__4222__auto__
														_ = s__533___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__533___2) {
															{
																var c__914__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__533___2)
																var size__915__auto__ = cljs_core.Count.X_invoke_Arity1(c__914__auto__).(float64)
																var b__535 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__915__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__914__auto__, size__915__auto__, b__535
																if cljs_core.Truth_(func() interface{} {
																	var i__534 = float64(0)
																	_ = i__534
																	for {
																		if i__534 < size__915__auto__ {
																			{
																				var e = c__914__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__534)
																				_ = e
																				if e.(float64) < float64(3) {
																					{
																						var m = (e.(float64) * e.(float64))
																						_ = m
																						cljs_core.Chunk_append.X_invoke_Arity2(b__535, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																						i__534 = (i__534 + float64(1))
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
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__535), iter__532.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__533___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__535), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__533___2)
																_ = e
																if e.(float64) < float64(3) {
																	{
																		var m = (e.(float64) * e.(float64))
																		_ = m
																		return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__532.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__533___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_940), nil, nil})
						})
					}(&cljs_core.AFn{}, v_940)
					_ = iter__916__auto__
					return iter__916__auto__.X_invoke_Arity1(v_940).(*cljs_core.CljsCoreLazySeq)
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
				var a10_946 = cljs_core.Partial.X_invoke_Arity2(cljs_core.X_PLUS_, float64(10)).(cljs_core.CljsCoreIFn)
				var a20_947 = cljs_core.Partial.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), float64(10)).(cljs_core.CljsCoreIFn)
				var a21_948 = cljs_core.Partial.X_invoke_Arity4(cljs_core.X_PLUS_, float64(10), float64(10), float64(1)).(cljs_core.CljsCoreIFn)
				var a22_949 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)})).(cljs_core.CljsCoreIFn)
				var a23_950 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(1)})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = a10_946, a20_947, a21_948, a22_949, a23_950
				if cljs_core.X_EQ_.Arity2IIB(float64(110), func() interface{} {
					var G__538 = float64(100)
					_ = G__538
					return a10_946.X_invoke_Arity1(G__538)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 110 (a10 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(120), func() interface{} {
					var G__539 = float64(100)
					_ = G__539
					return a20_947.X_invoke_Arity1(G__539)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 120 (a20 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(121), func() interface{} {
					var G__540 = float64(100)
					_ = G__540
					return a21_948.X_invoke_Arity1(G__540)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 121 (a21 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(122), func() interface{} {
					var G__541 = float64(100)
					_ = G__541
					return a22_949.X_invoke_Arity1(G__541)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 122 (a22 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(123), func() interface{} {
					var G__542 = float64(100)
					_ = G__542
					return a23_950.X_invoke_Arity1(G__542)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 123 (a23 100))")}))
				}
			}
			{
				var n2_951 = cljs_core.Comp.X_invoke_Arity2(cljs_core.First, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n3_952 = cljs_core.Comp.X_invoke_Arity3(cljs_core.First, cljs_core.Rest, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n4_953 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n5_954 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n6_955 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = n2_951, n3_952, n4_953, n5_954, n6_955
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() interface{} {
					var G__543 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__543
					return n2_951.X_invoke_Arity1(G__543)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (n2 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() interface{} {
					var G__544 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__544
					return n3_952.X_invoke_Arity1(G__544)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (n3 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__545 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__545
					return n4_953.X_invoke_Arity1(G__545)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (n4 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(5), func() interface{} {
					var G__546 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__546
					return n5_954.X_invoke_Arity1(G__546)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 5 (n5 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), func() interface{} {
					var G__547 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__547
					return n6_955.X_invoke_Arity1(G__547)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (n6 [1 2 3 4 5 6 7]))")}))
				}
			}
			{
				var sf_956 = cljs_core.Some_fn.X_invoke_Arity3(cljs_core.Number_QMARK_, cljs_core.Keyword_QMARK_, cljs_core.Symbol_QMARK_).(cljs_core.CljsCoreIFn)
				_ = sf_956
				if cljs_core.Truth_(func() interface{} {
					var G__548 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					var G__549 = float64(1)
					_, _ = G__548, G__549
					return sf_956.X_invoke_Arity2(G__548, G__549)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf :foo 1)")}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__550 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					_ = G__550
					return sf_956.X_invoke_Arity1(G__550)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf :foo)")}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__551 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})
					var G__552 = float64(1)
					_, _ = G__551, G__552
					return sf_956.X_invoke_Arity2(G__551, G__552)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf (quote bar) 1)")}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__553 = cljs_core.CljsCorePersistentVector_EMPTY
					var G__554 = cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)
					_, _ = G__553, G__554
					return sf_956.X_invoke_Arity2(G__553, G__554)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (not (sf [] ()))")}))
				}
			}
			{
				var ep_957 = cljs_core.Every_pred.X_invoke_Arity2(cljs_core.Number_QMARK_, cljs_core.Zero_QMARK_).(cljs_core.CljsCoreIFn)
				_ = ep_957
				if cljs_core.Truth_(func() interface{} {
					var G__555 = float64(0)
					var G__556 = float64(0)
					var G__557 = float64(0)
					_, _, _ = G__555, G__556, G__557
					return ep_957.X_invoke_Arity3(G__555, G__556, G__557)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (ep 0 0 0)")}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__558 = float64(1)
					var G__559 = float64(2)
					var G__560 = float64(3)
					var G__561 = float64(0)
					_, _, _, _ = G__558, G__559, G__560, G__561
					return ep_957.X_invoke_Arity4(G__558, G__559, G__560, G__561)
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
				var vec__562 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__562, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__562, float64(1), nil)
				_, _, _ = vec__562, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] [1 2]] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), func() cljs_core.CljsCoreISet {
				var vec__563 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__563, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__563, float64(1), nil)
				_, _, _ = vec__563, a, b
				return cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{a, b}, true).(*cljs_core.CljsCorePersistentHashSet)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 2} (let [[a b] [1 2]] #{a b}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__564 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__564___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__564) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__564)
					} else {
						return map__564
					}
				}()
				var a = cljs_core.Get.X_invoke_Arity2(map__564___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				var b = cljs_core.Get.X_invoke_Arity2(map__564___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				_, _, _, _ = map__564, map__564___1, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (let [{a :a, b :b} {:a 1, :b 2}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__565 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__565___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__565) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__565)
					} else {
						return map__565
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__565___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__565___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__565, map__565___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (let [{:keys [a b]} {:a 1, :b 2}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__566 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__566, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__566, float64(1), nil)
				var v = vec__566
				_, _, _, _ = vec__566, a, b, v
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b, v}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 [1 2]] (let [[a b :as v] [1 2]] [a b v]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(42)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__567 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__567___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__567) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__567)
					} else {
						return map__567
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity3(map__567___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(42))
				var a = cljs_core.Get.X_invoke_Arity2(map__567___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__567, map__567___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 42] (let [{:keys [a b], :or {b 42}} {:a 1}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), nil}, nil}), func() cljs_core.CljsCoreIVector {
				var map__568 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__568___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__568) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__568)
					} else {
						return map__568
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__568___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__568___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__568, map__568___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 nil] (let [{:keys [a b], :or {c 42}} {:a 1}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__569 = cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)
				var a = cljs_core.Nth.X_invoke_Arity3(vec__569, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__569, float64(1), nil)
				_, _, _ = vec__569, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] (quote (1 2))] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIMap {
				var vec__570 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__570, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__570, float64(1), nil)
				_, _, _ = vec__570, a, b
				return cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{a, b}, true, false).(*cljs_core.CljsCorePersistentArrayMap)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= {1 2} (let [[a b] [1 2]] {a b}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__571 = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))
				var a = cljs_core.Nth.X_invoke_Arity3(vec__571, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__571, float64(1), nil)
				_, _, _ = vec__571, a, b
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
				var a_958 = cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{})
				_ = a_958
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20), float64(30)}, nil}), cljs_core.Seq.Arity1IQ(func() []interface{} {
					var a__1035__auto__ = a_958
					var ret = cljs_core.Aclone.X_invoke_Arity1(a__1035__auto__).([]interface{})
					_, _ = a__1035__auto__, ret
					{
						var i = float64(0)
						_ = i
						for {
							if i < float64(len(a__1035__auto__)) {
								ret[int(i)] = (float64(10) * (a_958[int(i)]).(float64))
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
					var a__1041__auto__ = a_958
					_ = a__1041__auto__
					{
						var i = float64(0)
						var ret = float64(0)
						_, _ = i, ret
						for {
							if i < float64(len(a__1041__auto__)) {
								i, ret = (i + float64(1)), (ret + (a_958[int(i)]).(float64))
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_958), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (seq a) (seq (to-array [1 2 3])))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), func() float64 { a_958[int(float64(0))] = float64(42); return a_958[int(float64(0))].(float64) }()) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (aset a 0 42))")}))
				}
				if cljs_core.Not_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_958), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
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
				var coll_959 = (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})
				var shuffles_960 = cljs_core.Filter.X_invoke_Arity2(func(G__961 *cljs_core.AFn, coll_959 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__961, 1, func(p1__58_SHARP_ interface{}) interface{} {
						return cljs_core.Not_EQ_.Arity2IIB(coll_959, p1__58_SHARP_)
					})
				}(&cljs_core.AFn{}, coll_959), cljs_core.Take.X_invoke_Arity2(float64(100), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Shuffle, coll_959).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
				_, _ = coll_959, shuffles_960
				if !(cljs_core.Empty_QMARK_.Arity1IB(shuffles_960)) {
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
				var s_962 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = s_962
				{
					var n__1047__auto___963 = float64(5)
					_ = n__1047__auto___963
					{
						var n_964 = float64(0)
						_ = n_964
						for {
							if n_964 < n__1047__auto___963 {
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_962, cljs_core.Conj, n_964)
								n_964 = (n_964 + float64(1))
								continue
							} else {
							}
							break
						}
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Deref.X_invoke_Arity1(s_962)) {
				} else {
					panic((&js.Error{("Assert failed: (= [0 1 2 3 4] (clojure.core/deref s))")}))
				}
			}
			{
				var v_965 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})
				var s_966 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)).(*cljs_core.CljsCoreAtom)
				_, _ = v_965, s_966
				{
					var seq__572_967 interface{} = cljs_core.Seq.Arity1IQ(v_965)
					var chunk__573_968 interface{} = nil
					var count__574_969 = float64(0)
					var i__575_970 = float64(0)
					_, _, _, _ = seq__572_967, chunk__573_968, count__574_969, i__575_970
					for {
						if i__575_970 < count__574_969 {
							{
								var n_971 = chunk__573_968.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__575_970)
								_ = n_971
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_966, cljs_core.Conj, n_971)
								seq__572_967, chunk__573_968, count__574_969, i__575_970 = seq__572_967, chunk__573_968, count__574_969, (i__575_970 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___972 = cljs_core.Seq.Arity1IQ(seq__572_967)
								_ = temp__4222__auto___972
								if cljs_core.Truth_(temp__4222__auto___972) {
									{
										var seq__572_973___1 = temp__4222__auto___972
										_ = seq__572_973___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__572_973___1) {
											{
												var c__947__auto___974 = cljs_core.Chunk_first.X_invoke_Arity1(seq__572_973___1)
												_ = c__947__auto___974
												seq__572_967, chunk__573_968, count__574_969, i__575_970 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__572_973___1), c__947__auto___974, cljs_core.Count.X_invoke_Arity1(c__947__auto___974).(float64), float64(0)
												continue
											}
										} else {
											{
												var n_975 = cljs_core.First.X_invoke_Arity1(seq__572_973___1)
												_ = n_975
												cljs_core.Swap_BANG_.X_invoke_Arity3(s_966, cljs_core.Conj, n_975)
												seq__572_967, chunk__573_968, count__574_969, i__575_970 = cljs_core.Next.Arity1IQ(seq__572_973___1), nil, float64(0), float64(0)
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Deref.X_invoke_Arity1(s_966), cljs_core.Reverse.X_invoke_Arity1(v_965)) {
				} else {
					panic((&js.Error{("Assert failed: (= (clojure.core/deref s) (reverse v))")}))
				}
			}
			{
				var a_976 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				var d_977 = (&cljs_core.CljsCoreDelay{func(G__978 *cljs_core.AFn, a_976 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__978, 0, func() interface{} {
						return cljs_core.Swap_BANG_.X_invoke_Arity2(a_976, cljs_core.Inc)
					})
				}(&cljs_core.AFn{}, a_976), nil})
				_, _ = a_976, d_977
				if cljs_core.Realized_QMARK_.Arity1IB(d_977) == false {
				} else {
					panic((&js.Error{("Assert failed: (false? (realized? d))")}))
				}
				if cljs_core.Deref.X_invoke_Arity1(a_976).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (clojure.core/deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_977)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_976)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref a))")}))
				}
				if cljs_core.Realized_QMARK_.Arity1IB(d_977) == true {
				} else {
					panic((&js.Error{("Assert failed: (true? (realized? d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_977)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_976)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Force.X_invoke_Arity1(d_977), cljs_core.Deref.X_invoke_Arity1(d_977)) {
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
				var f_979 = cljs_core.Memoize.X_invoke_Arity1(func(G__980 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__980, 0, func() interface{} {
						return cljs_core.Rand.Arity0F()
					})
				}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn)
				_ = f_979
				{
					f_979.X_invoke_Arity0()
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return f_979.X_invoke_Arity0()
				}(), func() interface{} {
					return f_979.X_invoke_Arity0()
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
				var d_981 = cljs_core.Group_by.X_invoke_Arity2(cljs_core.Second, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(6), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), float64(4), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), float64(2)}, nil}))
				_ = d_981
				if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_981, float64(1))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (count (get d 1)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_981, float64(2))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (count (get d 2)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_981, float64(4))).(float64)) {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__982 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__982, 1, func(p1__59_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(6), float64(8), float64(10)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__983 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__983, 1, func(p1__60_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__984 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__984, 2, func(p1__61_SHARP_ interface{}, p2__62_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(5)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__985 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__985, 2, func(p1__64_SHARP_ interface{}, p2__63_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil}), cljs_core.Map_indexed.X_invoke_Arity2(func(G__986 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__986, 2, func(p1__65_SHARP_ interface{}, p2__66_SHARP_ interface{}) interface{} {
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
						return func(G__987 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__987, 0, func() interface{} {
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
				var method_table__1057__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1058__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1059__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1060__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1061__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__
				return (&cljs_core.CljsCoreMultiFn{"bar", func(G__988 *cljs_core.AFn, method_table__1057__auto__ *cljs_core.CljsCoreAtom, prefer_table__1058__auto__ *cljs_core.CljsCoreAtom, method_cache__1059__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1060__auto__ *cljs_core.CljsCoreAtom, hierarchy__1061__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__988, 2, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{x, y}, nil})
					})
				}(&cljs_core.AFn{}, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1061__auto__, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__})
			}()

			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}), func(G__989 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__989, 2, func(x interface{}, y interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)})
				})
			}(&cljs_core.AFn{}))
			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})}, nil}), func(G__990 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__990, 2, func(x interface{}, y interface{}) interface{} {
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
				var G__577 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				var G__578 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				_, _ = G__577, G__578
				return Bar.X_invoke_Arity2(G__577, G__578)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :rect-shape (bar :cljs.core-test/rect :cljs.core-test/rect))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)}), cljs_core.Apply.X_invoke_Arity2(Bar.X_get_method_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= :rect-shape (apply (-get-method bar [:cljs.core-test/rect :cljs.core-test/shape]) [:cljs.core-test/rect :cljs.core-test/shape]))")}))
			}
			Nested_dispatch = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1057__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1058__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1059__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1060__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1061__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__
				return (&cljs_core.CljsCoreMultiFn{"nested-dispatch", func(G__991 *cljs_core.AFn, method_table__1057__auto__ *cljs_core.CljsCoreAtom, prefer_table__1058__auto__ *cljs_core.CljsCoreAtom, method_cache__1059__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1060__auto__ *cljs_core.CljsCoreAtom, hierarchy__1061__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__991, 1, func(m interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(m))
					})
				}(&cljs_core.AFn{}, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1061__auto__, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__})
			}()

			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), func(G__992 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__992, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__993 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__993, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__579 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil})
				_ = G__579
				return Nested_dispatch.X_invoke_Arity1(G__579)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :nested-a (nested-dispatch {:a {:b :c}}))")}))
			}
			Nested_dispatch2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1057__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1058__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1059__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1060__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1061__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__
				return (&cljs_core.CljsCoreMultiFn{"nested-dispatch2", cljs_core.Ffirst, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1061__auto__, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__})
			}()

			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__994 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__994, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__995 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__995, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__580 = (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})}, nil})
				_ = G__580
				return Nested_dispatch2.X_invoke_Arity1(G__580)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :nested-a (nested-dispatch2 [[:a :b]]))")}))
			}
			Foo1 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1057__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1058__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1059__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1060__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1061__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__
				return (&cljs_core.CljsCoreMultiFn{"foo1", func(G__996 *cljs_core.AFn, method_table__1057__auto__ *cljs_core.CljsCoreAtom, prefer_table__1058__auto__ *cljs_core.CljsCoreAtom, method_cache__1059__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1060__auto__ *cljs_core.CljsCoreAtom, hierarchy__1061__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__996, 0, func(args__ ...interface{}) interface{} {
						var args = cljs_core.Seq.Arity1IQ(args__[0])
						_ = args
						return cljs_core.First.X_invoke_Arity1(args)
					})
				}(&cljs_core.AFn{}, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1061__auto__, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__})
			}()

			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__997 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__997, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)})
				})
			}(&cljs_core.AFn{}))
			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__998 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__998, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)}), func() interface{} {
				var G__581 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
				_ = G__581
				return Foo1.X_invoke_Arity1(G__581)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :a-return (foo1 :a))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)}), func() interface{} {
				var G__582 = float64(1)
				_ = G__582
				return Foo1.X_invoke_Arity1(G__582)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :default-return (foo1 1))")}))
			}
			Area = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1057__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1058__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1059__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1060__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1061__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__
				return (&cljs_core.CljsCoreMultiFn{"area", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Shape", Fqn: "Shape", X_hash: float64(-248980586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1061__auto__, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__})
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

			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Rect", Fqn: "Rect", X_hash: float64(-420704620)}), func(G__999 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__999, 1, func(r interface{}) interface{} {
					return ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wd", Fqn: "wd", X_hash: float64(-183204751)}).X_invoke_Arity1(r).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ht", Fqn: "ht", X_hash: float64(214115472)}).X_invoke_Arity1(r).(float64))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Circle", Fqn: "Circle", X_hash: float64(-158896930)}), func(G__1000 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1000, 1, func(c interface{}) interface{} {
					return (Math.PI * ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64)))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1001 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1001, 1, func(x interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)})
				})
			}(&cljs_core.AFn{}))
			R = Rect.X_invoke_Arity2(float64(4), float64(13)).(cljs_core.CljsCoreIMap)

			C = Circle.X_invoke_Arity1(float64(12)).(cljs_core.CljsCoreIMap)

			if cljs_core.X_EQ_.Arity2IIB(float64(52), func() interface{} {
				var G__583 = R
				_ = G__583
				return Area.X_invoke_Arity1(G__583)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 52 (area r))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)}), func() interface{} {
				var G__584 = cljs_core.CljsCorePersistentArrayMap_EMPTY
				_ = G__584
				return Area.X_invoke_Arity1(G__584)
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
				var method_table__1057__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1058__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1059__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1060__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1061__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__
				return (&cljs_core.CljsCoreMultiFn{"apply-multi-test", func(G__1002 *cljs_core.AFn, method_table__1057__auto__ *cljs_core.CljsCoreAtom, prefer_table__1058__auto__ *cljs_core.CljsCoreAtom, method_cache__1059__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1060__auto__ *cljs_core.CljsCoreAtom, hierarchy__1061__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1002, 3, func(___ interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
						return float64(0)
					})
				}(&cljs_core.AFn{}, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1061__auto__, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__})
			}()

			Apply_multi_test.X_add_method_Arity3(float64(0), func(G__1003 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1003, 2, func(x interface{}) interface{} {
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
				var method_table__1057__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1058__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1059__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1060__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1061__auto__ = cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), My_map_hierarchy}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__
				return (&cljs_core.CljsCoreMultiFn{"my-map?", cljs_core.Type_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1061__auto__, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__})
			}()

			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)}), func(G__1004 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1004, 1, func(___ interface{}) interface{} {
					return true
				})
			}(&cljs_core.AFn{}))
			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1005 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1005, 1, func(___ interface{}) interface{} {
					return false
				})
			}(&cljs_core.AFn{}))
			{
				var seq__585_1006 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__586_1007 interface{} = nil
				var count__587_1008 = float64(0)
				var i__588_1009 = float64(0)
				_, _, _, _ = seq__585_1006, chunk__586_1007, count__587_1008, i__588_1009
				for {
					if i__588_1009 < count__587_1008 {
						{
							var m_1010 = chunk__586_1007.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__588_1009)
							_ = m_1010
							if cljs_core.Truth_(func() interface{} {
								var G__589 = m_1010
								_ = G__589
								return My_map_QMARK_.X_invoke_Arity1(G__589)
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (my-map? m)")}))
							}
							seq__585_1006, chunk__586_1007, count__587_1008, i__588_1009 = seq__585_1006, chunk__586_1007, count__587_1008, (i__588_1009 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1011 = cljs_core.Seq.Arity1IQ(seq__585_1006)
							_ = temp__4222__auto___1011
							if cljs_core.Truth_(temp__4222__auto___1011) {
								{
									var seq__585_1012___1 = temp__4222__auto___1011
									_ = seq__585_1012___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__585_1012___1) {
										{
											var c__947__auto___1013 = cljs_core.Chunk_first.X_invoke_Arity1(seq__585_1012___1)
											_ = c__947__auto___1013
											seq__585_1006, chunk__586_1007, count__587_1008, i__588_1009 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__585_1012___1), c__947__auto___1013, cljs_core.Count.X_invoke_Arity1(c__947__auto___1013).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_1014 = cljs_core.First.X_invoke_Arity1(seq__585_1012___1)
											_ = m_1014
											if cljs_core.Truth_(func() interface{} {
												var G__590 = m_1014
												_ = G__590
												return My_map_QMARK_.X_invoke_Arity1(G__590)
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (my-map? m)")}))
											}
											seq__585_1006, chunk__586_1007, count__587_1008, i__588_1009 = cljs_core.Next.Arity1IQ(seq__585_1012___1), nil, float64(0), float64(0)
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
				var seq__591_1015 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentVector_EMPTY, float64(1), "asdf", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}))
				var chunk__592_1016 interface{} = nil
				var count__593_1017 = float64(0)
				var i__594_1018 = float64(0)
				_, _, _, _ = seq__591_1015, chunk__592_1016, count__593_1017, i__594_1018
				for {
					if i__594_1018 < count__593_1017 {
						{
							var not_m_1019 = chunk__592_1016.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__594_1018)
							_ = not_m_1019
							if cljs_core.Not.Arity1IB(func() interface{} {
								var G__595 = not_m_1019
								_ = G__595
								return My_map_QMARK_.X_invoke_Arity1(G__595)
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (not (my-map? not-m))")}))
							}
							seq__591_1015, chunk__592_1016, count__593_1017, i__594_1018 = seq__591_1015, chunk__592_1016, count__593_1017, (i__594_1018 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1020 = cljs_core.Seq.Arity1IQ(seq__591_1015)
							_ = temp__4222__auto___1020
							if cljs_core.Truth_(temp__4222__auto___1020) {
								{
									var seq__591_1021___1 = temp__4222__auto___1020
									_ = seq__591_1021___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__591_1021___1) {
										{
											var c__947__auto___1022 = cljs_core.Chunk_first.X_invoke_Arity1(seq__591_1021___1)
											_ = c__947__auto___1022
											seq__591_1015, chunk__592_1016, count__593_1017, i__594_1018 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__591_1021___1), c__947__auto___1022, cljs_core.Count.X_invoke_Arity1(c__947__auto___1022).(float64), float64(0)
											continue
										}
									} else {
										{
											var not_m_1023 = cljs_core.First.X_invoke_Arity1(seq__591_1021___1)
											_ = not_m_1023
											if cljs_core.Not.Arity1IB(func() interface{} {
												var G__596 = not_m_1023
												_ = G__596
												return My_map_QMARK_.X_invoke_Arity1(G__596)
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (not (my-map? not-m))")}))
											}
											seq__591_1015, chunk__592_1016, count__593_1017, i__594_1018 = cljs_core.Next.Arity1IQ(seq__591_1021___1), nil, float64(0), float64(0)
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
				var pv_1024 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange))
				_ = pv_1024
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity2(pv_1024, float64(96)), float64(96)) {
				} else {
					panic((&js.Error{("Assert failed: (= (nth pv 96) 96)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity3(pv_1024, float64(97), nil), nil) {
				} else {
					panic((&js.Error{("Assert failed: (= (nth pv 97 nil) nil)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__597 = float64(96)
					_ = G__597
					return pv_1024.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__597)
				}(), float64(96)) {
				} else {
					panic((&js.Error{("Assert failed: (= (pv 96) 96)")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(cljs_core.CljsCorePersistentVector_EMPTY)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq []))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reverse.X_invoke_Arity1(pv_1024), cljs_core.Rseq.Arity1IQ(pv_1024)) {
				} else {
					panic((&js.Error{("Assert failed: (= (reverse pv) (rseq pv))")}))
				}
			}
			{
				var pv_1025 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(33)).(*cljs_core.CljsCoreRange))
				_ = pv_1025
				if cljs_core.X_EQ_.Arity2IIB(pv_1025, cljs_core.Conj.X_invoke_Arity2(cljs_core.Conj.X_invoke_Arity2(cljs_core.Pop.X_invoke_Arity1(cljs_core.Pop.X_invoke_Arity1(pv_1025)), float64(31)), float64(32))) {
				} else {
					panic((&js.Error{("Assert failed: (= pv (-> pv pop pop (conj 31) (conj 32)))")}))
				}
			}
			{
				var stack1_1026 = cljs_core.Pop.X_invoke_Arity1(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange)))
				var stack2_1027 = cljs_core.Pop.X_invoke_Arity1(stack1_1026)
				_, _ = stack1_1026, stack2_1027
				if cljs_core.X_EQ_.Arity2IIB(float64(95), cljs_core.Peek.X_invoke_Arity1(stack1_1026)) {
				} else {
					panic((&js.Error{("Assert failed: (= 95 (peek stack1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(94), cljs_core.Peek.X_invoke_Arity1(stack2_1027)) {
				} else {
					panic((&js.Error{("Assert failed: (= 94 (peek stack2))")}))
				}
			}
			{
				var sentinel_1028 = cljs_core.Rand.Arity0F()
				_ = sentinel_1028
				if reflect.DeepEqual(sentinel_1028, func() (return__1029 interface{}) {
					defer func() {
						if e598 := recover(); e598 != nil {
							if func() bool { _, instanceof := e598.(*js.Error); return instanceof }() {
								{
									var ___ = e598
									_ = ___
									return__1029 = sentinel_1028
								}
							} else {
								panic(e598)

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
				var v1_1030 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				var v2_1031 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(5)).(*cljs_core.CljsCoreRange))
				var s_1032 = cljs_core.Subvec.X_invoke_Arity3(v1_1030, float64(2), float64(8)).(*cljs_core.CljsCoreSubvec)
				_, _, _ = v1_1030, v2_1031, s_1032
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(s_1032, cljs_core.Subvec.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2(v1_1030, float64(2)).(*cljs_core.CljsCoreSubvec), float64(0), float64(6)).(*cljs_core.CljsCoreSubvec), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Take.X_invoke_Arity2(float64(6), cljs_core.Drop.X_invoke_Arity2(float64(2), v1_1030).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= s (-> v1 (subvec 2) (subvec 0 6)) (->> v1 (drop 2) (take 6)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Count.X_invoke_Arity1(s_1032).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (count s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Pop.X_invoke_Arity1(s_1032)) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 4 5 6] (pop s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(7), cljs_core.Peek.X_invoke_Arity1(s_1032)) {
				} else {
					panic((&js.Error{("Assert failed: (= 7 (peek s))")}))
				}
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(1)}, nil}), cljs_core.Assoc.X_invoke_Arity3(s_1032, float64(6), float64(1)), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Conj.X_invoke_Arity2(s_1032, float64(1))}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 4 5 6 7 1] (assoc s 6 1) (conj s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(27), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, s_1032)) {
				} else {
					panic((&js.Error{("Assert failed: (= 27 (reduce + s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(s_1032, cljs_core.Vec.X_invoke_Arity1(s_1032)) {
				} else {
					panic((&js.Error{("Assert failed: (= s (vec s))")}))
				}
				{
					var m_1033 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), float64(1)}, nil})
					_ = m_1033
					if cljs_core.X_EQ_.Arity2IIB(m_1033, cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(s_1032, m_1033))) {
					} else {
						panic((&js.Error{("Assert failed: (= m (meta (with-meta s m)))")}))
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1034 interface{}) {
					defer func() {
						if e599 := recover(); e599 != nil {
							if func() bool { _, instanceof := e599.(*js.Error); return instanceof }() {
								{
									var e = e599
									_ = e
									return__1034 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e599)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1031, float64(0), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 0 6) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1035 interface{}) {
					defer func() {
						if e600 := recover(); e600 != nil {
							if func() bool { _, instanceof := e600.(*js.Error); return instanceof }() {
								{
									var e = e600
									_ = e
									return__1035 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e600)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1031, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1036 interface{}) {
					defer func() {
						if e601 := recover(); e601 != nil {
							if func() bool { _, instanceof := e601.(*js.Error); return instanceof }() {
								{
									var e = e601
									_ = e
									return__1036 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e601)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1031, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1037 interface{}) {
					defer func() {
						if e602 := recover(); e602 != nil {
							if func() bool { _, instanceof := e602.(*js.Error); return instanceof }() {
								{
									var e = e602
									_ = e
									return__1037 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e602)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1031, float64(3), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 3 6) (catch js/Error e :fail)))")}))
				}
				if reflect.DeepEqual(v1_1030, cljs_core.Subvec.X_invoke_Arity3(s_1032, float64(1), float64(4)).(*cljs_core.CljsCoreSubvec).V) {
				} else {
					panic((&js.Error{("Assert failed: (identical? v1 (.-v (subvec s 1 4)))")}))
				}
				{
					var sentinel_1038 = cljs_core.Rand.Arity0F()
					var s_1039___1 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2))
					_, _ = sentinel_1038, s_1039___1
					if reflect.DeepEqual(sentinel_1038, func() (return__1040 interface{}) {
						defer func() {
							if e603 := recover(); e603 != nil {
								if func() bool { _, instanceof := e603.(*js.Error); return instanceof }() {
									{
										var ___ = e603
										_ = ___
										return__1040 = sentinel_1038
									}
								} else {
									panic(e603)

								}
							}
						}()
						{
							{
								var G__604 = float64(-1)
								_ = G__604
								return s_1039___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__604)
							}
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (identical? sentinel (try (s -1) (catch js/Error _ sentinel)))")}))
					}
					if reflect.DeepEqual(sentinel_1038, func() (return__1041 interface{}) {
						defer func() {
							if e605 := recover(); e605 != nil {
								if func() bool { _, instanceof := e605.(*js.Error); return instanceof }() {
									{
										var ___ = e605
										_ = ___
										return__1041 = sentinel_1038
									}
								} else {
									panic(e605)

								}
							}
						}()
						{
							{
								var G__606 = float64(1)
								_ = G__606
								return s_1039___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__606)
							}
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (identical? sentinel (try (s 1) (catch js/Error _ sentinel)))")}))
					}
				}
				{
					var sv1_1042 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2)).(*cljs_core.CljsCoreSubvec)
					var sv2_1043 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(1)).(*cljs_core.CljsCoreSubvec)
					_, _ = sv1_1042, sv2_1043
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(sv1_1042), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(*cljs_core.CljsCoreList)) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq sv1) (quote (1)))")}))
					}
					if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(sv2_1043)) {
					} else {
						panic((&js.Error{("Assert failed: (nil? (rseq sv2))")}))
					}
				}
			}
			{
				var v1_1044 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(15), float64(48)).(*cljs_core.CljsCoreRange))
				var v2_1045 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(40), float64(57)).(*cljs_core.CljsCoreRange))
				var v1_1046___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Pop_BANG_.X_invoke_Arity1(cljs_core.Transient.X_invoke_Arity1(v1_1044)), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v2_1047___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(v2_1045), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v_1048 = cljs_core.Into.X_invoke_Arity2(v1_1046___1, v2_1047___1)
				_, _, _, _, _ = v1_1044, v2_1045, v1_1046___1, v2_1047___1, v_1048
				if cljs_core.X_EQ_.Arity2IIB(v_1048, cljs_core.Vec.X_invoke_Arity1(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(16), float64(47)).(*cljs_core.CljsCoreRange), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(41), float64(57)).(*cljs_core.CljsCoreRange), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})})).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (= v (vec (concat [:quux] (range 16 47) [:foo] [:quux] (range 41 57) [:bar])))")}))
				}
			}
			{
				var v_1049 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)
				var xs_1050 interface{} = cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)
				_, _ = v_1049, xs_1050
				for {
					{
						var temp__4220__auto___1051 = cljs_core.First.X_invoke_Arity1(xs_1050)
						_ = temp__4220__auto___1051
						if cljs_core.Truth_(temp__4220__auto___1051) {
							{
								var x_1052 = temp__4220__auto___1051
								_ = x_1052
								v_1049, xs_1050 = func() interface{} {
									var pred__607 = func(G__1053 *cljs_core.AFn, v_1049 interface{}, xs_1050 interface{}, x_1052 interface{}, temp__4220__auto___1051 interface{}) *cljs_core.AFn {
										return cljs_core.Fn(G__1053, 2, func(p1__67_SHARP_ interface{}, p2__68_SHARP_ interface{}) interface{} {
											{
												var G__610 = cljs_core.Mod.X_invoke_Arity2(p2__68_SHARP_, float64(3)).(float64)
												_ = G__610
												return p1__67_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__610)
											}
										})
									}(&cljs_core.AFn{}, v_1049, xs_1050, x_1052, temp__4220__auto___1051)
									var expr__608 = x_1052
									_, _ = pred__607, expr__608
									if cljs_core.Truth_(pred__607.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(0), nil, float64(2), nil}, nil}, nil}), expr__608)) {
										return cljs_core.Conj_BANG_.X_invoke_Arity2(v_1049, x_1052)
									} else {
										if cljs_core.Truth_(pred__607.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), expr__608)) {
											return cljs_core.Assoc_BANG_.X_invoke_Arity3(v_1049, cljs_core.Count.X_invoke_Arity1(v_1049).(float64), x_1052)
										} else {
											panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(expr__608).(string))}))
										}
									}
								}(), cljs_core.Next.Arity1IQ(xs_1050)
								continue
							}
						} else {
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)), cljs_core.Persistent_BANG_.X_invoke_Arity1(v_1049)) {
							} else {
								panic((&js.Error{("Assert failed: (= (vec (range 100)) (persistent! v))")}))
							}
						}
					}
					break
				}
			}
			{
				var m1_1054 interface{} = cljs_core.CljsCorePersistentHashMap_EMPTY
				var m2_1055 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_EMPTY)
				var i_1056 = float64(0)
				_, _, _ = m1_1054, m2_1055, i_1056
				for {
					if i_1056 < float64(100) {
						m1_1054, m2_1055, i_1056 = cljs_core.Assoc.X_invoke_Arity3(m1_1054, i_1056, i_1056), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_1055, i_1056, i_1056), (i_1056 + float64(1))
						continue
					} else {
						{
							var m2_1057___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_1055)
							_ = m2_1057___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1054).(float64), float64(100)) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m1) 100)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1057___1).(float64), float64(100)) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m2) 100)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_1054, m2_1057___1) {
							} else {
								panic((&js.Error{("Assert failed: (= m1 m2)")}))
							}
							{
								var i_1058___1 = float64(0)
								_ = i_1058___1
								for {
									if i_1058___1 < float64(100) {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__611 = i_1058___1
											_ = G__611
											return m1_1054.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__611)
										}(), i_1058___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__612 = i_1058___1
											_ = G__612
											return m2_1057___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__612)
										}(), i_1058___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m2 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_1054, i_1058___1), i_1058___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_1057___1, i_1058___1), i_1058___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m2 i) i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_1054, i_1058___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m1 i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_1057___1, i_1058___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m2 i)")}))
										}
										i_1058___1 = (i_1058___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_1054))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range 100) (range 100)) (sort-by first (seq m1)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_1057___1))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range 100) (range 100)) (sort-by first (seq m2)))")}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_1054, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{("Assert failed: (not (contains? (dissoc m1 3) 3))")}))
							}
						}
					}
					break
				}
			}
			{
				var m_1059 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_1059
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1059).(float64), float64(7)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 7)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1059, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
				}
			}
			{
				var m_1060 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_1060
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1060).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1060, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var m_1061 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_1061
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1061).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1061, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var tm_1062 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_1062
				{
					var tm_1063___1 interface{} = tm_1062
					var ks_1064 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_1063___1, ks_1064
					for {
						{
							var temp__4220__auto___1065 = cljs_core.First.X_invoke_Arity1(ks_1064)
							_ = temp__4220__auto___1065
							if cljs_core.Truth_(temp__4220__auto___1065) {
								{
									var k_1066 = temp__4220__auto___1065
									_ = k_1066
									tm_1063___1, ks_1064 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1063___1, k_1066), cljs_core.Next.Arity1IQ(ks_1064)
									continue
								}
							} else {
								{
									var m_1067 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1063___1)
									_ = m_1067
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1067).(float64), float64(7)) {
									} else {
										panic((&js.Error{("Assert failed: (= (count m) 7)")}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_1067, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var tm_1068 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_1068
				{
					var seq__613_1069 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__614_1070 interface{} = nil
					var count__615_1071 = float64(0)
					var i__616_1072 = float64(0)
					_, _, _, _ = seq__613_1069, chunk__614_1070, count__615_1071, i__616_1072
					for {
						if i__616_1072 < count__615_1071 {
							{
								var k_1073 = chunk__614_1070.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__616_1072)
								_ = k_1073
								if cljs_core.X_EQ_.Arity2IIB(k_1073, cljs_core.Get.X_invoke_Arity2(tm_1068, k_1073)) {
								} else {
									panic((&js.Error{("Assert failed: (= k (get tm k))")}))
								}
								seq__613_1069, chunk__614_1070, count__615_1071, i__616_1072 = seq__613_1069, chunk__614_1070, count__615_1071, (i__616_1072 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___1074 = cljs_core.Seq.Arity1IQ(seq__613_1069)
								_ = temp__4222__auto___1074
								if cljs_core.Truth_(temp__4222__auto___1074) {
									{
										var seq__613_1075___1 = temp__4222__auto___1074
										_ = seq__613_1075___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__613_1075___1) {
											{
												var c__947__auto___1076 = cljs_core.Chunk_first.X_invoke_Arity1(seq__613_1075___1)
												_ = c__947__auto___1076
												seq__613_1069, chunk__614_1070, count__615_1071, i__616_1072 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__613_1075___1), c__947__auto___1076, cljs_core.Count.X_invoke_Arity1(c__947__auto___1076).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_1077 = cljs_core.First.X_invoke_Arity1(seq__613_1075___1)
												_ = k_1077
												if cljs_core.X_EQ_.Arity2IIB(k_1077, cljs_core.Get.X_invoke_Arity2(tm_1068, k_1077)) {
												} else {
													panic((&js.Error{("Assert failed: (= k (get tm k))")}))
												}
												seq__613_1069, chunk__614_1070, count__615_1071, i__616_1072 = cljs_core.Next.Arity1IQ(seq__613_1075___1), nil, float64(0), float64(0)
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
					var m_1078 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1068)
					_ = m_1078
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1079 float64) {
						defer func() {
							if e617 := recover(); e617 != nil {
								if func() bool { _, instanceof := e617.(*js.Error); return instanceof }() {
									{
										var e = e617
										_ = e
										return__1079 = float64(2)
									}
								} else {
									panic(e617)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1068, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1080 float64) {
						defer func() {
							if e618 := recover(); e618 != nil {
								if func() bool { _, instanceof := e618.(*js.Error); return instanceof }() {
									{
										var e = e618
										_ = e
										return__1080 = float64(2)
									}
								} else {
									panic(e618)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_1068, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1081 float64) {
						defer func() {
							if e619 := recover(); e619 != nil {
								if func() bool { _, instanceof := e619.(*js.Error); return instanceof }() {
									{
										var e = e619
										_ = e
										return__1081 = float64(2)
									}
								} else {
									panic(e619)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1068)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (persistent! tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1082 float64) {
						defer func() {
							if e620 := recover(); e620 != nil {
								if func() bool { _, instanceof := e620.(*js.Error); return instanceof }() {
									{
										var e = e620
										_ = e
										return__1082 = float64(2)
									}
								} else {
									panic(e620)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_1068)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (count tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_1078, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var m_1083 = cljs_core.Assoc.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentHashMap_EMPTY, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_ = m_1083
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1083, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1083, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1083).(float64), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 2)")}))
				}
				{
					var m_1084___1 = cljs_core.Dissoc.X_invoke_Arity2(m_1083, Fixed_hash_foo)
					_ = m_1084___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1084___1, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1084___1, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1084___1).(float64), float64(1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 1)")}))
					}
				}
			}
			{
				var m_1085 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_1086___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m_1085, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_, _ = m_1085, m_1086___1
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1086___1).(float64), float64(102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 102)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1086___1, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1086___1, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				{
					var m_1087___2 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_1086___1, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_1087___2
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1087___2, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1087___2, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1087___2).(float64), float64(98)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 98)")}))
					}
				}
			}
			{
				var m_1088 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_1089___1 = cljs_core.Transient.X_invoke_Arity1(m_1088)
				var m_1090___2 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_1089___1, Fixed_hash_foo, float64(1))
				var m_1091___3 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_1090___2, Fixed_hash_bar, float64(2))
				var m_1092___4 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m_1091___3)
				_, _, _, _, _ = m_1088, m_1089___1, m_1090___2, m_1091___3, m_1092___4
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1092___4).(float64), float64(102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 102)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1092___4, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1092___4, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				{
					var m_1093___5 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_1092___4, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_1093___5
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1093___5, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1093___5, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1093___5).(float64), float64(98)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 98)")}))
					}
				}
			}
			Array_map_conversion_threshold = cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD

			{
				var m1_1094 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_1095 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY)
				var i_1096 = float64(0)
				_, _, _ = m1_1094, m2_1095, i_1096
				for {
					if i_1096 < Array_map_conversion_threshold {
						m1_1094, m2_1095, i_1096 = cljs_core.Assoc.X_invoke_Arity3(m1_1094, i_1096, i_1096), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_1095, i_1096, i_1096), (i_1096 + float64(1))
						continue
					} else {
						{
							var m2_1097___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_1095)
							_ = m2_1097___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1094).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m1) array-map-conversion-threshold)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1097___1).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m2) array-map-conversion-threshold)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_1094, m2_1097___1) {
							} else {
								panic((&js.Error{("Assert failed: (= m1 m2)")}))
							}
							{
								var i_1098___1 = float64(0)
								_ = i_1098___1
								for {
									if i_1098___1 < Array_map_conversion_threshold {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__621 = i_1098___1
											_ = G__621
											return m1_1094.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__621)
										}(), i_1098___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__622 = i_1098___1
											_ = G__622
											return m2_1097___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__622)
										}(), i_1098___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m2 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_1094, i_1098___1), i_1098___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_1097___1, i_1098___1), i_1098___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m2 i) i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_1094, i_1098___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m1 i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_1097___1, i_1098___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m2 i)")}))
										}
										i_1098___1 = (i_1098___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_1094))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m1)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_1097___1))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m2)))")}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_1094, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{("Assert failed: (not (contains? (dissoc m1 3) 3))")}))
							}
						}
					}
					break
				}
			}
			{
				var m_1099 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_1099
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1099).(float64), float64(7)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 7)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1099, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
				}
			}
			{
				var m_1100 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_1100
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1100).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1100, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var m_1101 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_1101
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1101).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1101, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var tm_1102 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_1102
				{
					var tm_1103___1 interface{} = tm_1102
					var ks_1104 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_1103___1, ks_1104
					for {
						{
							var temp__4220__auto___1105 = cljs_core.First.X_invoke_Arity1(ks_1104)
							_ = temp__4220__auto___1105
							if cljs_core.Truth_(temp__4220__auto___1105) {
								{
									var k_1106 = temp__4220__auto___1105
									_ = k_1106
									tm_1103___1, ks_1104 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1103___1, k_1106), cljs_core.Next.Arity1IQ(ks_1104)
									continue
								}
							} else {
								{
									var m_1107 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1103___1)
									_ = m_1107
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1107).(float64), float64(7)) {
									} else {
										panic((&js.Error{("Assert failed: (= (count m) 7)")}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_1107, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var tm_1108 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_1108
				{
					var seq__623_1109 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__624_1110 interface{} = nil
					var count__625_1111 = float64(0)
					var i__626_1112 = float64(0)
					_, _, _, _ = seq__623_1109, chunk__624_1110, count__625_1111, i__626_1112
					for {
						if i__626_1112 < count__625_1111 {
							{
								var k_1113 = chunk__624_1110.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__626_1112)
								_ = k_1113
								if cljs_core.X_EQ_.Arity2IIB(k_1113, cljs_core.Get.X_invoke_Arity2(tm_1108, k_1113)) {
								} else {
									panic((&js.Error{("Assert failed: (= k (get tm k))")}))
								}
								seq__623_1109, chunk__624_1110, count__625_1111, i__626_1112 = seq__623_1109, chunk__624_1110, count__625_1111, (i__626_1112 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___1114 = cljs_core.Seq.Arity1IQ(seq__623_1109)
								_ = temp__4222__auto___1114
								if cljs_core.Truth_(temp__4222__auto___1114) {
									{
										var seq__623_1115___1 = temp__4222__auto___1114
										_ = seq__623_1115___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__623_1115___1) {
											{
												var c__947__auto___1116 = cljs_core.Chunk_first.X_invoke_Arity1(seq__623_1115___1)
												_ = c__947__auto___1116
												seq__623_1109, chunk__624_1110, count__625_1111, i__626_1112 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__623_1115___1), c__947__auto___1116, cljs_core.Count.X_invoke_Arity1(c__947__auto___1116).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_1117 = cljs_core.First.X_invoke_Arity1(seq__623_1115___1)
												_ = k_1117
												if cljs_core.X_EQ_.Arity2IIB(k_1117, cljs_core.Get.X_invoke_Arity2(tm_1108, k_1117)) {
												} else {
													panic((&js.Error{("Assert failed: (= k (get tm k))")}))
												}
												seq__623_1109, chunk__624_1110, count__625_1111, i__626_1112 = cljs_core.Next.Arity1IQ(seq__623_1115___1), nil, float64(0), float64(0)
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
					var m_1118 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1108)
					_ = m_1118
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1119 float64) {
						defer func() {
							if e627 := recover(); e627 != nil {
								if func() bool { _, instanceof := e627.(*js.Error); return instanceof }() {
									{
										var e = e627
										_ = e
										return__1119 = float64(2)
									}
								} else {
									panic(e627)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1108, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1120 float64) {
						defer func() {
							if e628 := recover(); e628 != nil {
								if func() bool { _, instanceof := e628.(*js.Error); return instanceof }() {
									{
										var e = e628
										_ = e
										return__1120 = float64(2)
									}
								} else {
									panic(e628)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_1108, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1121 float64) {
						defer func() {
							if e629 := recover(); e629 != nil {
								if func() bool { _, instanceof := e629.(*js.Error); return instanceof }() {
									{
										var e = e629
										_ = e
										return__1121 = float64(2)
									}
								} else {
									panic(e629)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1108)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (persistent! tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1122 float64) {
						defer func() {
							if e630 := recover(); e630 != nil {
								if func() bool { _, instanceof := e630.(*js.Error); return instanceof }() {
									{
										var e = e630
										_ = e
										return__1122 = float64(2)
									}
								} else {
									panic(e630)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_1108)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (count tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_1118, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
					}
				}
			}
			{
				var m_1123 = cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))
				_ = m_1123
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1123).(float64), (float64(2) * Array_map_conversion_threshold)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) (* 2 array-map-conversion-threshold))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__631 = Array_map_conversion_threshold
					_ = G__631
					return m_1123.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__631)
				}(), Array_map_conversion_threshold) {
				} else {
					panic((&js.Error{("Assert failed: (= (m array-map-conversion-threshold) array-map-conversion-threshold)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1123, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity2(func(G__1124 *cljs_core.AFn, m_1123 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1124, 1, func(p1__69_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__69_SHARP_, p1__69_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1123), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (= m (into (.-EMPTY cljs.core/PersistentHashMap) (map (fn* [p1__69#] (vector p1__69# p1__69#)) (range (* 2 array-map-conversion-threshold)))))")}))
				}
			}
			{
				var m1_1125 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_1126 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var i_1127 = float64(0)
				_, _, _ = m1_1125, m2_1126, i_1127
				for {
					if i_1127 < float64(100) {
						m1_1125, m2_1126, i_1127 = cljs_core.Assoc.X_invoke_Arity3(m1_1125, i_1127, i_1127), cljs_core.Assoc.X_invoke_Arity3(m2_1126, ("foo"+cljs_core.Str.X_invoke_Arity1(i_1127).(string)), i_1127), (i_1127 + float64(1))
						continue
					} else {
						if cljs_core.X_EQ_.Arity2IIB(m1_1125, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{("Assert failed: (= m1 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (range 100) (range 100))))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(m2_1126, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Map_.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.Str, "foo").(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{("Assert failed: (= m2 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (map (partial str \"foo\") (range 100)) (range 100))))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1125).(float64), float64(100)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count m1) 100)")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1126).(float64), float64(100)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count m2) 100)")}))
						}
					}
					break
				}
			}
			{
				var i_1128 = float64(0)
				var m_1129 interface{} = cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(-1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))
				var result_1130 interface{} = cljs_core.CljsCorePersistentVector_EMPTY
				_, _, _ = i_1128, m_1129, result_1130
				for {
					if i_1128 <= (cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) {
						i_1128, m_1129, result_1130 = (i_1128 + float64(1)), cljs_core.Assoc.X_invoke_Arity3(m_1129, i_1128, i_1128), cljs_core.Conj.X_invoke_Arity2(result_1130, cljs_core.Meta.X_invoke_Arity1(m_1129))
						continue
					} else {
						{
							var n_1131 = ((cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) + float64(1))
							var expected_1132 = cljs_core.Repeat.X_invoke_Arity2(n_1131, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})).(*cljs_core.CljsCoreLazySeq)
							_, _ = n_1131, expected_1132
							if cljs_core.X_EQ_.Arity2IIB(result_1130, expected_1132) {
							} else {
								panic((&js.Error{("Assert failed: (= result expected)")}))
							}
						}
					}
					break
				}
			}
			{
				var m1_1133 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_1134 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var m2_1135 = cljs_core.Sorted_map_by.X_invoke_ArityVariadic(c2_1134, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{})).(*cljs_core.CljsCorePersistentTreeMap)
				_, _, _ = m1_1133, c2_1134, m2_1135
				if reflect.DeepEqual(cljs_core.Compare, cljs_core.Native_get_instance_field.X_invoke_Arity2(m1_1133, "Comp")) {
				} else {
					panic((&js.Error{("Assert failed: (identical? compare (.-comp m1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(m1_1133).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count m1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(m2_1135).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count m2))")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(m1_1133)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq m1))")}))
				}
				{
					var m1_1136___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m1_1133, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					var m2_1137___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m2_1135, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					_, _ = m1_1136___1, m2_1137___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1136___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m1) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1137___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m2) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_1136___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m1) (list [:bar 2] [:foo 1] [:quux 3]))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_1137___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m2) (list [:quux 3] [:foo 1] [:bar 2]))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_1136___1), cljs_core.Rseq.Arity1IQ(m2_1137___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m1) (rseq m2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_1137___1), cljs_core.Rseq.Arity1IQ(m1_1136___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m2) (rseq m1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m1_1136___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= (conj m1 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m1_1136___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count (conj m1 [:wibble 4])) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m2_1137___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= (conj m2 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m2_1137___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count (conj m2 [:wibble 4])) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m1_1136___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})).(cljs_core.CljsCoreICollection).X_conj_Arity2(nil)) {
					} else {
						panic((&js.Error{("Assert failed: (= (map key (assoc m1 nil 4)) (list nil :bar :foo :quux))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m2_1137___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(nil).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (map key (assoc m2 nil 4)) (list :quux :foo :bar nil))")}))
					}
				}
			}
			{
				var m_1138 = cljs_core.Apply.X_invoke_Arity2(cljs_core.Sorted_map, cljs_core.Mapcat.X_invoke_ArityVariadic(func(G__1141 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1141, 1, func(p1__70_SHARP_ interface{}) interface{} {
						return cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(p1__70_SHARP_).(cljs_core.CljsCoreICollection).X_conj_Arity2(p1__70_SHARP_)
					})
				}(&cljs_core.AFn{}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_ArityVariadic(cljs_core.Partial.X_invoke_Arity2(cljs_core.Apply, cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(10)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(20), float64(30)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(50), float64(60)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(30), float64(40)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(40), float64(50)}, nil})}, nil})}))})))
				var s1_1139 = cljs_core.Map_.X_invoke_Arity2(func(G__1142 *cljs_core.AFn, m_1138 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1142, 1, func(p1__71_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__71_SHARP_, p1__71_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1138), cljs_core.Range_.X_invoke_Arity1(float64(60)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				var s2_1140 = cljs_core.Map_.X_invoke_Arity2(func(G__1143 *cljs_core.AFn, m_1138 interface{}, s1_1139 *cljs_core.CljsCoreLazySeq) *cljs_core.AFn {
					return cljs_core.Fn(G__1143, 1, func(p1__72_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__72_SHARP_, p1__72_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1138, s1_1139), cljs_core.Range_.X_invoke_Arity3(float64(59), float64(-1), float64(-1)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				_, _, _ = m_1138, s1_1139, s2_1140
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1138).(float64), float64(60)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 60)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m_1138), s1_1139) {
				} else {
					panic((&js.Error{("Assert failed: (= (seq m) s1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(m_1138), s2_1140) {
				} else {
					panic((&js.Error{("Assert failed: (= (rseq m) s2)")}))
				}
			}
			{
				var m_1144 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
				_ = m_1144
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m_1144, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(2), float64(3)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (dissoc m :foo) (hash-map :bar 2 :quux 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(m_1144, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))).(float64), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count (dissoc m :foo)) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(m_1144), cljs_core.Hash.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(1), float64(2), float64(3)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (hash m) (hash (hash-map :foo 1 :bar 2 :quux 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1144, cljs_core.X_LT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m < :foo) (list [:bar 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1144, cljs_core.X_LT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m <= :foo) (list [:bar 2] [:foo 1]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1144, cljs_core.X_GT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m > :foo) (list [:quux 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1144, cljs_core.X_GT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil})).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m >= :foo) (list [:foo 1] [:quux 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1145 *cljs_core.AFn, m_1144 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1145, 1, func(p1__73_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity2(func(G__1146 *cljs_core.AFn, m_1144 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__1146, 2, func(___ interface{}, x interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_1144), p1__73_SHARP_)
					})
				}(&cljs_core.AFn{}, m_1144), m_1144).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__73#] (reduce (fn [_ x] x) p1__73#)) m) (list 2 1 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1147 *cljs_core.AFn, m_1144 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1147, 1, func(p1__74_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity3(func(G__1148 *cljs_core.AFn, m_1144 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__1148, 2, func(x interface{}, ___ interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_1144), float64(7), p1__74_SHARP_)
					})
				}(&cljs_core.AFn{}, m_1144), m_1144).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__74#] (reduce (fn [x _] x) 7 p1__74#)) m) (list 7 7 7))")}))
				}
			}
			{
				var s1_1149 = cljs_core.Sorted_set.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_1150 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var s2_1151 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c2_1150, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c3_1152 = func(G__1155 *cljs_core.AFn, s1_1149 interface{}, c2_1150 cljs_core.CljsCoreIFn, s2_1151 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1155, 2, func(p1__75_SHARP_ interface{}, p2__76_SHARP_ interface{}) interface{} {
						return cljs_core.Compare.Arity2IIF(cljs_core.Quot.X_invoke_Arity2(p1__75_SHARP_, float64(2)).(float64), cljs_core.Quot.X_invoke_Arity2(p2__76_SHARP_, float64(2)).(float64))
					})
				}(&cljs_core.AFn{}, s1_1149, c2_1150, s2_1151)
				var s3_1153 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c3_1152, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var s4_1154 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(cljs_core.X_LT_, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				_, _, _, _, _, _ = s1_1149, c2_1150, s2_1151, c3_1152, s3_1153, s4_1154
				if reflect.DeepEqual(cljs_core.Compare, s1_1149.(cljs_core.CljsCoreISorted).X_comparator_Arity1()) {
				} else {
					panic((&js.Error{("Assert failed: (identical? compare (-comparator s1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(s1_1149).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count s1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(s2_1151).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count s2))")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(s1_1149)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq s1))")}))
				}
				{
					var s1_1156___1 = cljs_core.Conj.X_invoke_ArityVariadic(s1_1149, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s2_1157___1 = cljs_core.Conj.X_invoke_ArityVariadic(s2_1151, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s3_1158___1 = cljs_core.Conj.X_invoke_ArityVariadic(s3_1153, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(7), float64(8), float64(9)}))
					var s4_1159___1 = cljs_core.Conj.X_invoke_ArityVariadic(s4_1154, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					_, _, _, _ = s1_1156___1, s2_1157___1, s3_1158___1, s4_1159___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_1156___1), cljs_core.Hash.X_invoke_Arity1(s2_1157___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (hash s1) (hash s2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_1156___1), cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (hash s1) (hash #{1 3 2}))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_1156___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq s1) (list 1 2 3))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_1156___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq s1) (list 3 2 1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_1157___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq s2) (list 3 2 1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_1157___1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq s2) (list 1 2 3))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_1156___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s1) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_1157___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s2) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s3_1158___1).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s3) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(s3_1158___1, float64(0)), float64(1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get s3 0) 1)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1158___1, cljs_core.X_GT_, float64(5)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 > 5) (list 7 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1158___1, cljs_core.X_GT_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 > 6) (list 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1158___1, cljs_core.X_GT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 6) (list 7 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1158___1, cljs_core.X_LT_, float64(0)), cljs_core.CljsCoreList_EMPTY) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 0) (list))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1158___1, cljs_core.X_LT_, float64(5)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 5) (list 1 2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1158___1, cljs_core.X_LT_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 6) (list 1 2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1158___1, cljs_core.X_LT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 <= 6) (list 1 2 7))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s3_1158___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT__EQ_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 2 <= 6) (list 2 7))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s4_1159___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT_, float64(3)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s4 >= 2 < 3) (list 2))")}))
					}
					{
						var s1_1160___2 = cljs_core.Disj.X_invoke_Arity2(s1_1156___1, float64(2))
						var s2_1161___2 = cljs_core.Disj.X_invoke_Arity2(s2_1157___1, float64(2))
						_, _ = s1_1160___2, s2_1161___2
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_1160___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{("Assert failed: (= (seq s1) (list 1 3))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_1160___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{("Assert failed: (= (rseq s1) (list 3 1))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_1161___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{("Assert failed: (= (seq s2) (list 3 1))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_1161___2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3)).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{("Assert failed: (= (rseq s2) (list 1 3))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_1160___2).(float64), float64(2)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count s1) 2)")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_1161___2).(float64), float64(2)) {
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
					return cljs_core.Fn(map__GT_Person, 1, func(G__634 interface{}) interface{} {
						return (&CljsCore_testPerson{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}).X_invoke_Arity1(G__634), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}).X_invoke_Arity1(G__634), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__634, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})})), nil})
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
					return cljs_core.Fn(map__GT_A, 1, func(G__653 interface{}) interface{} {
						return (&CljsCore_testA{nil, cljs_core.Dissoc.X_invoke_Arity1(G__653), nil})
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
					return cljs_core.Fn(map__GT_C, 1, func(G__664 interface{}) interface{} {
						return (&CljsCore_testC{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__664), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__664), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}).X_invoke_Arity1(G__664), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__664, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})})), nil})
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
				var s_1162 = "abc"
				_ = s_1162
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_1162).Length) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (.-length s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_1162).Length) {
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
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1162).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_("abc").Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring \"abc\" 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", func(G__1163 *cljs_core.AFn, s_1162 string) *cljs_core.AFn {
					return cljs_core.Fn(G__1163, 2, func(target685 interface{}, start interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target685, "Substring", []interface{}{start})
					})
				}(&cljs_core.AFn{}, s_1162).X_invoke_Arity2(s_1162, float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" ((memfn substring start) s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1162).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s substring 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1162).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s (substring 1)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1162).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s (substring 1 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1162).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring s 1 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_(s_1162).ToUpperCase()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" (. s (toUpperCase)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_("abc").ToUpperCase()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" (. \"abc\" (toUpperCase)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", func(G__1164 *cljs_core.AFn, s_1162 string) *cljs_core.AFn {
					return cljs_core.Fn(G__1164, 1, func(target686 interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target686, "ToUpperCase", []interface{}{})
					})
				}(&cljs_core.AFn{}, s_1162).X_invoke_Arity1(s_1162)) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" ((memfn toUpperCase) s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("BC", cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_1162).ToUpperCase(), "Substring", []interface{}{float64(1)})) {
				} else {
					panic((&js.Error{("Assert failed: (= \"BC\" (. (. s (toUpperCase)) substring 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_1162).ToUpperCase(), "Substring", []interface{}{float64(1)}), "Length")) {
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
					return cljs_core.Fn(map__GT_A2, 1, func(G__689 interface{}) interface{} {
						return (&CljsCore_testA2{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__689), nil, cljs_core.Dissoc.X_invoke_Arity2(G__689, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
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
					return cljs_core.Fn(map__GT_B, 1, func(G__704 interface{}) interface{} {
						return (&CljsCore_testB{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__704), nil, cljs_core.Dissoc.X_invoke_Arity2(G__704, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
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

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(func() *CljsCore_testT717 {
				X__GT_t717 = func(__GT_t717 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_t717, 2, func(test_stuff___1 interface{}, meta718 interface{}) interface{} {
						return (&CljsCore_testT717{test_stuff___1, meta718})
					})
				}(&cljs_core.AFn{})

				return (&CljsCore_testT717{test_stuff, nil})
			}(), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (meta (with-meta (reify IFoo (foo [this] :foo)) {:foo :bar})) {:foo :bar})")}))
			}
			Foo2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1057__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1058__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1059__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1060__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1061__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__, hierarchy__1061__auto__
				return (&cljs_core.CljsCoreMultiFn{"foo2", cljs_core.Identity, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1061__auto__, method_table__1057__auto__, prefer_table__1058__auto__, method_cache__1059__auto__, cached_hierarchy__1060__auto__})
			}()

			Foo2.X_add_method_Arity3(float64(0), func(G__1165 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1165, 1, func(x interface{}) interface{} {
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
				var g_1166 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}))}, true).(*cljs_core.CljsCorePersistentHashSet)
				var h_1167 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil})}, true).(*cljs_core.CljsCorePersistentHashSet)
				_, _ = g_1166, h_1167
				if cljs_core.X_EQ_.Arity2IIB(g_1166, h_1167) {
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
				var fv_1168 = (&CljsCore_testFirst{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})})
				var fs_1169 = (&CljsCore_testFirst{"asdf"})
				_, _ = fv_1168, fs_1169
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fv_1168.X_invoke_Arity0()
				}(), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (fv) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fs_1169.X_invoke_Arity0()
				}(), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((`` + cljs_core.Str.X_invoke_Arity1(fs_1169).(string)), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (str fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_1168.X_get_first_Arity1(), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-get-first fv) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fs_1169.X_get_first_Arity1(), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (-get-first fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_1168.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-find-first fv [1]) 1)")}))
				}
				if reflect.DeepEqual(func() interface{} {
					var G__732 = float64(1)
					_ = G__732
					return fv_1168.X_invoke_Arity1(G__732)
				}(), fv_1168) {
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
				var t_1170 = (&CljsCore_testDestructuringWithLocals{float64(1)})
				_ = t_1170
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(1)}, nil}), t_1170.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 1] (-find-first t [2 3]))")}))
				}
			}
			{
				var x_1171 = float64(1)
				_ = x_1171
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__736 = x_1171
					_ = G__736
					switch G__736 {
					case float64(1):
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})

					default:
						panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_1171).(string))}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x 1 :one) :one)")}))
				}
			}
			{
				var x_1173 = float64(1)
				_ = x_1173
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__737 = x_1173
					_ = G__737
					switch G__737 {
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
				var x_1175 = float64(1)
				_ = x_1175
				if cljs_core.X_EQ_.Arity2IIB(func() (return__1176 interface{}) {
					defer func() {
						if e738 := recover(); e738 != nil {
							if func() bool { _, instanceof := e738.(*js.Error); return instanceof }() {
								{
									var e = e738
									_ = e
									return__1176 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e738)

							}
						}
					}()
					{
						{
							var G__739 = x_1175
							_ = G__739
							switch G__739 {
							case float64(3):
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "three", Fqn: "three", X_hash: float64(-1651831795)})

							default:
								panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_1175).(string))}))

							}
						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (try (case x 3 :three) (catch js/Error e :fail)) :fail)")}))
				}
			}
			{
				var x_1178 = float64(1)
				_ = x_1178
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__740 = x_1178
					_ = G__740
					switch G__740 {
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
				var x_1180 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})
				_ = x_1180
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__741 = x_1180
					_ = G__741
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), G__741) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
					} else {
						panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_1180).(string))}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x [:a :b] :ok) :ok)")}))
				}
			}
			{
				var a_1181 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = a_1181
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__742 = a_1181
					_ = G__742
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__742) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__742) {
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
				var a_1182 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil})
				_ = a_1182
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__743 = a_1182
					_ = G__743
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__743) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__743) {
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
				var foo_1183 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = foo_1183
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__744 = foo_1183
					_ = G__744
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__744) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__744) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), G__744) {
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
					var G__745 = foo_1183
					_ = G__745
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "d", Str: "d", X_hash: float64(-682293345), X_meta: nil}), G__745) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__745) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__745) {
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
				var r_1184 = cljs_core.Range_.X_invoke_Arity1(float64(64)).(*cljs_core.CljsCoreRange)
				var v_1185 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, r_1184)
				_, _ = r_1184, v_1185
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1185)), cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1185))) {
				} else {
					panic((&js.Error{("Assert failed: (= (hash (seq v)) (hash (seq v)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Array_chunk.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)}).(*cljs_core.CljsCoreArrayChunk))) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (reduce + (array-chunk (array 1 2 3))))")}))
				}
				if func() bool {
					_, instanceof := cljs_core.Seq.Arity1IQ(v_1185).(*cljs_core.CljsCoreChunkedSeq)
					return instanceof
				}() {
				} else {
					panic((&js.Error{("Assert failed: (instance? ChunkedSeq (seq v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(r_1184, cljs_core.Seq.Arity1IQ(v_1185)) {
				} else {
					panic((&js.Error{("Assert failed: (= r (seq v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, r_1184).(*cljs_core.CljsCoreLazySeq), cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, v_1185).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (map inc r) (map inc v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, r_1184).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, v_1185).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (filter even? r) (filter even? v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, r_1184).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, v_1185).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (filter odd? r) (filter odd? v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic(r_1184, r_1184, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{r_1184})).(*cljs_core.CljsCoreLazySeq), cljs_core.Concat.X_invoke_ArityVariadic(v_1185, v_1185, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{v_1185})).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (concat r r r) (concat v v v))")}))
				}
				if cljs_core.Truth_(cljs_core.Native_satisfies_QMARK_.X_invoke_Arity2((&cljs_core.CljsCoreSymbol{Ns: "cljs.core", Name: "IReduce", Str: "cljs.core/IReduce", X_hash: float64(-577837345), X_meta: nil}), cljs_core.Seq.Arity1IQ(v_1185))) {
				} else {
					panic((&js.Error{("Assert failed: (satisfies? IReduce (seq v))")}))
				}
				if float64(2010) == cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1185)))))).(float64) {
				} else {
					panic((&js.Error{("Assert failed: (== 2010 (reduce + (nnext (nnext (seq v)))))")}))
				}
				if float64(2020) == cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1185)))))).(float64) {
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
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__1186 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1186, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), nil).(*cljs_core.CljsCoreCons)
				})
			}(&cljs_core.AFn{}), nil, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next (lazy-seq (cons 1 nil))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__1187 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1187, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), (&cljs_core.CljsCoreLazySeq{nil, func(G__1188 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__1188, 0, func() interface{} {
							return cljs_core.Cons.X_invoke_Arity2(float64(2), (&cljs_core.CljsCoreLazySeq{nil, func(G__1189 *cljs_core.AFn) *cljs_core.AFn {
								return cljs_core.Fn(G__1189, 0, func() interface{} {
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
					return cljs_core.Fn(map__GT_PrintMe, 1, func(G__748 interface{}) interface{} {
						return (&CljsCore_testPrintMe{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__748), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__748), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__748, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})})), nil})
					})
				}(&cljs_core.AFn{})

			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&CljsCore_testPrintMe{float64(1), float64(2), nil, nil, nil})})).(string), "#cljs.core-test.PrintMe{:a 1, :b 2}") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str (PrintMe. 1 2)) \"#cljs.core-test.PrintMe{:a 1, :b 2}\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{float64(1289585655666)})})).(string), "#inst \"2010-11-12T18:14:15.666-00:00\"") {
			} else {
				panic((&js.Error{("Assert failed: (= (pr-str (js/Date. 1289585655666)) \"#inst \\\"2010-11-12T18:14:15.666-00:00\\\"\")")}))
			}
			{
				var uuid_str_1190 = "550e8400-e29b-41d4-a716-446655440000"
				var uuid_1191 = (&cljs_core.CljsCoreUUID{uuid_str_1190})
				_, _ = uuid_str_1190, uuid_1191
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{uuid_1191})).(string), ("#uuid \"" + cljs_core.Str.X_invoke_Arity1(uuid_str_1190).(string) + "\"")) {
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
					X__GT_t770 = func(__GT_t770 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t770, 4, func(f___1 interface{}, baz___1 interface{}, test_stuff___1 interface{}, meta771 interface{}) interface{} {
							return (&CljsCore_testT770{f___1, baz___1, test_stuff___1, meta771})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT770{f, baz, test_stuff, nil})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(float64(2), Baz.X_invoke_Arity1(cljs_core.Inc).(*CljsCore_testT770).X_bar_Arity2(float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (-bar (baz inc) 1))")}))
			}
			{
				var x_1192 = "original"
				_ = x_1192
				Original_closure_stmt = func(original_closure_stmt *cljs_core.AFn, x_1192 string) *cljs_core.AFn {
					return cljs_core.Fn(original_closure_stmt, 0, func() interface{} {
						return x_1192
					})
				}(&cljs_core.AFn{}, x_1192)

			}
			{
				var x_1193 = "overwritten"
				_ = x_1193
				if cljs_core.X_EQ_.Arity2IIB("original", Original_closure_stmt.X_invoke_Arity0().(string)) {
				} else {
					panic((&js.Error{("Assert failed: (= \"original\" (original-closure-stmt))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB("original", func() string {
				var x = "original"
				var oce = func(G__1194 *cljs_core.AFn, x string) *cljs_core.AFn {
					return cljs_core.Fn(G__1194, 0, func() interface{} {
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
					var x_1195___1 = func(G__1196 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__1196, 0, func() interface{} {
							return "overwritten"
						})
					}(&cljs_core.AFn{})
					_ = x_1195___1
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
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)}), cljs_core.Reduce_kv.X_invoke_Arity3(func(G__1197 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1197, 3, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
								return cljs_core.Reduced.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)})).(*cljs_core.CljsCoreReduced)
							})
						}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentVector_EMPTY, data)) {
						} else {
							panic((&js.Error{("Assert failed: (= :reduced (reduce-kv (fn [_ _ _] (reduced :reduced)) [] data))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sort.X_invoke_Arity1(expect), cljs_core.Sort.X_invoke_Arity1(cljs_core.Reduce_kv.X_invoke_Arity3(func(G__1198 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1198, 3, func(r interface{}, k interface{}, v interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), func() (return__1199 interface{}) {
				defer func() {
					if e775 := recover(); e775 != nil {
						if func() bool { _, instanceof := e775.(*cljs_core.CljsCoreExceptionInfo); return instanceof }() {
							{
								var e = e775
								_ = e
								return__1199 = cljs_core.Ex_data.X_invoke_Arity1(e)
							}
						} else {
							panic(e775)

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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(1), float64(2)}, nil}), func(G__1200 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1200, 0, func(more__ ...interface{}) interface{} {
					var more = cljs_core.Seq.Arity1IQ(more__[0])
					_ = more
					return more
				})
			}(&cljs_core.AFn{}).X_invoke_Arity3(js.Undefined, float64(1), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= [js/undefined 1 2] ((fn [& more] more) js/undefined 1 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(4), float64(5)}, nil}), func(G__1201 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1201, 2, func(a_b_more__ ...interface{}) interface{} {
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
				var fs_1202 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = fs_1202
				{
					var seq__776_1203 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity1(float64(4)).(*cljs_core.CljsCoreRange))
					var chunk__778_1204 interface{} = nil
					var count__779_1205 = float64(0)
					var i__780_1206 = float64(0)
					_, _, _, _ = seq__776_1203, chunk__778_1204, count__779_1205, i__780_1206
					for {
						if i__780_1206 < count__779_1205 {
							{
								var x_1207 = chunk__778_1204.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__780_1206)
								_ = x_1207
								{
									var y_1208 = (x_1207.(float64) + float64(1))
									var f_1209 = func(G__1210 *cljs_core.AFn, seq__776_1203 interface{}, chunk__778_1204 interface{}, count__779_1205 float64, i__780_1206 float64, y_1208 float64, x_1207 interface{}, fs_1202 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
										return cljs_core.Fn(G__1210, 0, func() interface{} {
											return y_1208
										})
									}(&cljs_core.AFn{}, seq__776_1203, chunk__778_1204, count__779_1205, i__780_1206, y_1208, x_1207, fs_1202)
									_, _ = y_1208, f_1209
									cljs_core.Swap_BANG_.X_invoke_Arity3(fs_1202, cljs_core.Conj, f_1209)
								}
								seq__776_1203, chunk__778_1204, count__779_1205, i__780_1206 = seq__776_1203, chunk__778_1204, count__779_1205, (i__780_1206 + float64(1))
								continue
							}
						} else {
							{
								var temp__4222__auto___1211 = cljs_core.Seq.Arity1IQ(seq__776_1203)
								_ = temp__4222__auto___1211
								if cljs_core.Truth_(temp__4222__auto___1211) {
									{
										var seq__776_1212___1 = temp__4222__auto___1211
										_ = seq__776_1212___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__776_1212___1) {
											{
												var c__947__auto___1213 = cljs_core.Chunk_first.X_invoke_Arity1(seq__776_1212___1)
												_ = c__947__auto___1213
												seq__776_1203, chunk__778_1204, count__779_1205, i__780_1206 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__776_1212___1), c__947__auto___1213, cljs_core.Count.X_invoke_Arity1(c__947__auto___1213).(float64), float64(0)
												continue
											}
										} else {
											{
												var x_1214 = cljs_core.First.X_invoke_Arity1(seq__776_1212___1)
												_ = x_1214
												{
													var y_1215 = (x_1214.(float64) + float64(1))
													var f_1216 = func(G__1217 *cljs_core.AFn, seq__776_1203 interface{}, chunk__778_1204 interface{}, count__779_1205 float64, i__780_1206 float64, y_1215 float64, x_1214 interface{}, seq__776_1212___1 interface{}, temp__4222__auto___1211 cljs_core.CljsCoreISeq, fs_1202 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
														return cljs_core.Fn(G__1217, 0, func() interface{} {
															return y_1215
														})
													}(&cljs_core.AFn{}, seq__776_1203, chunk__778_1204, count__779_1205, i__780_1206, y_1215, x_1214, seq__776_1212___1, temp__4222__auto___1211, fs_1202)
													_, _ = y_1215, f_1216
													cljs_core.Swap_BANG_.X_invoke_Arity3(fs_1202, cljs_core.Conj, f_1216)
												}
												seq__776_1203, chunk__778_1204, count__779_1205, i__780_1206 = cljs_core.Next.Arity1IQ(seq__776_1212___1), nil, float64(0), float64(0)
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1218 *cljs_core.AFn, fs_1202 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__1218, 1, func(p1__77_SHARP_ interface{}) interface{} {
						{
							return p1__77_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
						}
					})
				}(&cljs_core.AFn{}, fs_1202), cljs_core.Deref.X_invoke_Arity1(fs_1202)).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCoreLazySeq{nil, func(G__1219 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1219, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCoreLazySeq{nil, func(G__1220 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1220, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreLazySeq{nil, func(G__1221 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1221, 0, func() interface{} {
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
					var sb__1117__auto__ = (&goog_string.StringBuffer{})
					_ = sb__1117__auto__
					{
						var _STAR_print_fn_STAR_782_1222 = cljs_core.X_STAR_print_fn_STAR_
						_ = _STAR_print_fn_STAR_782_1222
						func() {
							defer func() {
								cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_782_1222

							}()
							{
								cljs_core.X_STAR_print_fn_STAR_ = func(G__1223 *cljs_core.AFn, _STAR_print_fn_STAR_782_1222 interface{}, sb__1117__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
									return cljs_core.Fn(G__1223, 1, func(x__1118__auto__ interface{}) interface{} {
										return sb__1117__auto__.Append(x__1118__auto__)
									})
								}(&cljs_core.AFn{}, _STAR_print_fn_STAR_782_1222, sb__1117__auto__)

								func() bool {
									_, instanceof := f_BANG_.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})).(*cljs_core.CljsCoreSymbol)
									return instanceof
								}()
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
					return (`` + cljs_core.Str.X_invoke_Arity1(sb__1117__auto__).(string))
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
			if (cljs_core.First.X_invoke_Arity1(cljs_core.Filter.X_invoke_Arity2(func(G__1224 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1224, 1, func(p1__78_SHARP_ interface{}) interface{} {
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
				var a_1225 = func() *CljsCore_testT783 {
					X__GT_t783 = func(__GT_t783 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t783, 2, func(test_stuff___1 interface{}, meta784 interface{}) interface{} {
							return (&CljsCore_testT783{test_stuff___1, meta784})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT783{test_stuff, nil})
				}()
				var b_1226 = func() *CljsCore_testT786 {
					X__GT_t786 = func(__GT_t786 *cljs_core.AFn, a_1225 *CljsCore_testT783) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t786, 3, func(a___1 interface{}, test_stuff___1 interface{}, meta787 interface{}) interface{} {
							return (&CljsCore_testT786{a___1, test_stuff___1, meta787})
						})
					}(&cljs_core.AFn{}, a_1225)

					return (&CljsCore_testT786{a_1225, test_stuff, nil})
				}()
				var s_1227 = cljs_core.Set.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(128)).(*cljs_core.CljsCoreRange))
				_, _, _ = a_1225, b_1226, s_1227
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_1227, a_1225, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_1226}))), a_1225)), a_1225), cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_1227, a_1225, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_1226}))), a_1225)), a_1225)) {
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
			if reflect.ValueOf([]interface{}{float64(1), float64(2), float64(3)}).Kind() == reflect.Slice {
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
			{
				var seq__789_1228 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__790_1229 interface{} = nil
				var count__791_1230 = float64(0)
				var i__792_1231 = float64(0)
				_, _, _, _ = seq__789_1228, chunk__790_1229, count__791_1230, i__792_1231
				for {
					if i__792_1231 < count__791_1230 {
						{
							var n_1232 = chunk__790_1229.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__792_1231)
							_ = n_1232
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1233 interface{}) {
								defer func() {
									if e795 := recover(); e795 != nil {
										if func() bool { _, instanceof := e795.(*js.Error); return instanceof }() {
											{
												var e = e795
												_ = e
												return__1233 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e795)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1232, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1234 interface{}) {
								defer func() {
									if e796 := recover(); e796 != nil {
										if func() bool { _, instanceof := e796.(*js.Error); return instanceof }() {
											{
												var e = e796
												_ = e
												return__1234 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e796)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_1232, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))")}))
							}
							seq__789_1228, chunk__790_1229, count__791_1230, i__792_1231 = seq__789_1228, chunk__790_1229, count__791_1230, (i__792_1231 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1235 = cljs_core.Seq.Arity1IQ(seq__789_1228)
							_ = temp__4222__auto___1235
							if cljs_core.Truth_(temp__4222__auto___1235) {
								{
									var seq__789_1236___1 = temp__4222__auto___1235
									_ = seq__789_1236___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__789_1236___1) {
										{
											var c__947__auto___1237 = cljs_core.Chunk_first.X_invoke_Arity1(seq__789_1236___1)
											_ = c__947__auto___1237
											seq__789_1228, chunk__790_1229, count__791_1230, i__792_1231 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__789_1236___1), c__947__auto___1237, cljs_core.Count.X_invoke_Arity1(c__947__auto___1237).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1238 = cljs_core.First.X_invoke_Arity1(seq__789_1236___1)
											_ = n_1238
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1239 interface{}) {
												defer func() {
													if e797 := recover(); e797 != nil {
														if func() bool { _, instanceof := e797.(*js.Error); return instanceof }() {
															{
																var e = e797
																_ = e
																return__1239 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e797)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1238, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1240 interface{}) {
												defer func() {
													if e798 := recover(); e798 != nil {
														if func() bool { _, instanceof := e798.(*js.Error); return instanceof }() {
															{
																var e = e798
																_ = e
																return__1240 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e798)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_1238, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))")}))
											}
											seq__789_1228, chunk__790_1229, count__791_1230, i__792_1231 = cljs_core.Next.Arity1IQ(seq__789_1236___1), nil, float64(0), float64(0)
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
				var seq__799_1241 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__800_1242 interface{} = nil
				var count__801_1243 = float64(0)
				var i__802_1244 = float64(0)
				_, _, _, _ = seq__799_1241, chunk__800_1242, count__801_1243, i__802_1244
				for {
					if i__802_1244 < count__801_1243 {
						{
							var n_1245 = chunk__800_1242.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__802_1244)
							_ = n_1245
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1246 interface{}) {
								defer func() {
									if e805 := recover(); e805 != nil {
										if func() bool { _, instanceof := e805.(*js.Error); return instanceof }() {
											{
												var e = e805
												_ = e
												return__1246 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e805)

										}
									}
								}()
								{
									return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1245, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
							}
							seq__799_1241, chunk__800_1242, count__801_1243, i__802_1244 = seq__799_1241, chunk__800_1242, count__801_1243, (i__802_1244 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1247 = cljs_core.Seq.Arity1IQ(seq__799_1241)
							_ = temp__4222__auto___1247
							if cljs_core.Truth_(temp__4222__auto___1247) {
								{
									var seq__799_1248___1 = temp__4222__auto___1247
									_ = seq__799_1248___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__799_1248___1) {
										{
											var c__947__auto___1249 = cljs_core.Chunk_first.X_invoke_Arity1(seq__799_1248___1)
											_ = c__947__auto___1249
											seq__799_1241, chunk__800_1242, count__801_1243, i__802_1244 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__799_1248___1), c__947__auto___1249, cljs_core.Count.X_invoke_Arity1(c__947__auto___1249).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1250 = cljs_core.First.X_invoke_Arity1(seq__799_1248___1)
											_ = n_1250
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1251 interface{}) {
												defer func() {
													if e806 := recover(); e806 != nil {
														if func() bool { _, instanceof := e806.(*js.Error); return instanceof }() {
															{
																var e = e806
																_ = e
																return__1251 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e806)

														}
													}
												}()
												{
													return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1250, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
											}
											seq__799_1241, chunk__800_1242, count__801_1243, i__802_1244 = cljs_core.Next.Arity1IQ(seq__799_1248___1), nil, float64(0), float64(0)
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
				var map__807_1252 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__807_1253___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__807_1252) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__807_1252)
					} else {
						return map__807_1252
					}
				}()
				var b_1254 = cljs_core.Get.X_invoke_Arity2(map__807_1253___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a_1255 = cljs_core.Get.X_invoke_Arity2(map__807_1253___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__807_1252, map__807_1253___1, b_1254, a_1255
				if cljs_core.X_EQ_.Arity2IIB(float64(1), a_1255) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), b_1254) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 b)")}))
				}
			}
			{
				var map__808_1256 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__808_1257___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__808_1256) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__808_1256)
					} else {
						return map__808_1256
					}
				}()
				var d_1258 = cljs_core.Get.X_invoke_Arity2(map__808_1257___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_1259 = cljs_core.Get.X_invoke_Arity2(map__808_1257___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__808_1256, map__808_1257___1, d_1258, b_1259
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1259) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1258) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__809_1260 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__809_1261___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__809_1260) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__809_1260)
					} else {
						return map__809_1260
					}
				}()
				var d_1262 = cljs_core.Get.X_invoke_Arity2(map__809_1261___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_1263 = cljs_core.Get.X_invoke_Arity2(map__809_1261___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__809_1260, map__809_1261___1, d_1262, b_1263
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1263) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1262) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__810_1264 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}), float64(1), (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}), float64(2)}, nil})
				var map__810_1265___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__810_1264) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__810_1264)
					} else {
						return map__810_1264
					}
				}()
				var d_1266 = cljs_core.Get.X_invoke_Arity2(map__810_1265___1, (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}))
				var b_1267 = cljs_core.Get.X_invoke_Arity2(map__810_1265___1, (&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}))
				_, _, _, _ = map__810_1264, map__810_1265___1, d_1266, b_1267
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1267) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1266) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__811_1268 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}), float64(2)}, nil})
				var map__811_1269___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__811_1268) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__811_1268)
					} else {
						return map__811_1268
					}
				}()
				var y_1270 = cljs_core.Get.X_invoke_Arity2(map__811_1269___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}))
				var x_1271 = cljs_core.Get.X_invoke_Arity2(map__811_1269___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}))
				_, _, _, _ = map__811_1268, map__811_1269___1, y_1270, x_1271
				if cljs_core.X_EQ_.Arity2IIB(x_1271, float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= x 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(y_1270, float64(2)) {
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
								arr, names = cljs_core.Conj.X_invoke_Arity2(arr, func(G__1272 *cljs_core.AFn, arr interface{}, names interface{}, name interface{}) *cljs_core.AFn {
									return cljs_core.Fn(G__1272, 0, func() interface{} {
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
				var sb__1117__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1117__auto__
				{
					var _STAR_print_fn_STAR_812_1273 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_812_1273
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_812_1273

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__1274 *cljs_core.AFn, _STAR_print_fn_STAR_812_1273 interface{}, sb__1117__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__1274, 1, func(x__1118__auto__ interface{}) interface{} {
									return sb__1117__auto__.Append(x__1118__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_812_1273, sb__1117__auto__)

							{
								var seq__813_1275 interface{} = cljs_core.Seq.Arity1IQ(Cljs_739.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil})))
								var chunk__814_1276 interface{} = nil
								var count__815_1277 = float64(0)
								var i__816_1278 = float64(0)
								_, _, _, _ = seq__813_1275, chunk__814_1276, count__815_1277, i__816_1278
								for {
									if i__816_1278 < count__815_1277 {
										{
											var fn_1279 = chunk__814_1276.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__816_1278)
											_ = fn_1279
											{
												fn_1279.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
											}
											seq__813_1275, chunk__814_1276, count__815_1277, i__816_1278 = seq__813_1275, chunk__814_1276, count__815_1277, (i__816_1278 + float64(1))
											continue
										}
									} else {
										{
											var temp__4222__auto___1280 = cljs_core.Seq.Arity1IQ(seq__813_1275)
											_ = temp__4222__auto___1280
											if cljs_core.Truth_(temp__4222__auto___1280) {
												{
													var seq__813_1281___1 = temp__4222__auto___1280
													_ = seq__813_1281___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__813_1281___1) {
														{
															var c__947__auto___1282 = cljs_core.Chunk_first.X_invoke_Arity1(seq__813_1281___1)
															_ = c__947__auto___1282
															seq__813_1275, chunk__814_1276, count__815_1277, i__816_1278 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__813_1281___1), c__947__auto___1282, cljs_core.Count.X_invoke_Arity1(c__947__auto___1282).(float64), float64(0)
															continue
														}
													} else {
														{
															var fn_1283 = cljs_core.First.X_invoke_Arity1(seq__813_1281___1)
															_ = fn_1283
															{
																fn_1283.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
															}
															seq__813_1275, chunk__814_1276, count__815_1277, i__816_1278 = cljs_core.Next.Arity1IQ(seq__813_1281___1), nil, float64(0), float64(0)
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
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1117__auto__).(string))
			}(), ":a\n:b\n:c\n:d\n") {
			} else {
				panic((&js.Error{("Assert failed: (= (with-out-str (doseq [fn (cljs-739 [] [:a :b :c :d])] (fn))) \":a\\n:b\\n:c\\n:d\\n\")")}))
			}
			{
				var seq__817_1284 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__818_1285 interface{} = nil
				var count__819_1286 = float64(0)
				var i__820_1287 = float64(0)
				_, _, _, _ = seq__817_1284, chunk__818_1285, count__819_1286, i__820_1287
				for {
					if i__820_1287 < count__819_1286 {
						{
							var n_1288 = chunk__818_1285.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__820_1287)
							_ = n_1288
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1288)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get [1 2] n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1289 interface{}) {
								defer func() {
									if e823 := recover(); e823 != nil {
										if func() bool { _, instanceof := e823.(*js.Error); return instanceof }() {
											{
												var e = e823
												_ = e
												return__1289 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e823)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1288)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1288, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get [1 2] n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1290 interface{}) {
								defer func() {
									if e824 := recover(); e824 != nil {
										if func() bool { _, instanceof := e824.(*js.Error); return instanceof }() {
											{
												var e = e824
												_ = e
												return__1290 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e824)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1288, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1288)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (subvec [1 2] 1) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1291 interface{}) {
								defer func() {
									if e825 := recover(); e825 != nil {
										if func() bool { _, instanceof := e825.(*js.Error); return instanceof }() {
											{
												var e = e825
												_ = e
												return__1291 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e825)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1288)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1288, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (subvec [1 2] 1) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1292 interface{}) {
								defer func() {
									if e826 := recover(); e826 != nil {
										if func() bool { _, instanceof := e826.(*js.Error); return instanceof }() {
											{
												var e = e826
												_ = e
												return__1292 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e826)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1288, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1288)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (transient [1 2]) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1293 interface{}) {
								defer func() {
									if e827 := recover(); e827 != nil {
										if func() bool { _, instanceof := e827.(*js.Error); return instanceof }() {
											{
												var e = e827
												_ = e
												return__1293 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e827)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1288)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1288, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (transient [1 2]) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1294 interface{}) {
								defer func() {
									if e828 := recover(); e828 != nil {
										if func() bool { _, instanceof := e828.(*js.Error); return instanceof }() {
											{
												var e = e828
												_ = e
												return__1294 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e828)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1288, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1288)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (range 1 3) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1295 interface{}) {
								defer func() {
									if e829 := recover(); e829 != nil {
										if func() bool { _, instanceof := e829.(*js.Error); return instanceof }() {
											{
												var e = e829
												_ = e
												return__1295 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e829)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1288)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1288, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (range 1 3) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1296 interface{}) {
								defer func() {
									if e830 := recover(); e830 != nil {
										if func() bool { _, instanceof := e830.(*js.Error); return instanceof }() {
											{
												var e = e830
												_ = e
												return__1296 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e830)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1288, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))")}))
							}
							seq__817_1284, chunk__818_1285, count__819_1286, i__820_1287 = seq__817_1284, chunk__818_1285, count__819_1286, (i__820_1287 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1297 = cljs_core.Seq.Arity1IQ(seq__817_1284)
							_ = temp__4222__auto___1297
							if cljs_core.Truth_(temp__4222__auto___1297) {
								{
									var seq__817_1298___1 = temp__4222__auto___1297
									_ = seq__817_1298___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__817_1298___1) {
										{
											var c__947__auto___1299 = cljs_core.Chunk_first.X_invoke_Arity1(seq__817_1298___1)
											_ = c__947__auto___1299
											seq__817_1284, chunk__818_1285, count__819_1286, i__820_1287 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__817_1298___1), c__947__auto___1299, cljs_core.Count.X_invoke_Arity1(c__947__auto___1299).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1300 = cljs_core.First.X_invoke_Arity1(seq__817_1298___1)
											_ = n_1300
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1300)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get [1 2] n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1301 interface{}) {
												defer func() {
													if e831 := recover(); e831 != nil {
														if func() bool { _, instanceof := e831.(*js.Error); return instanceof }() {
															{
																var e = e831
																_ = e
																return__1301 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e831)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1300)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1300, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get [1 2] n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1302 interface{}) {
												defer func() {
													if e832 := recover(); e832 != nil {
														if func() bool { _, instanceof := e832.(*js.Error); return instanceof }() {
															{
																var e = e832
																_ = e
																return__1302 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e832)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1300, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1300)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (subvec [1 2] 1) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1303 interface{}) {
												defer func() {
													if e833 := recover(); e833 != nil {
														if func() bool { _, instanceof := e833.(*js.Error); return instanceof }() {
															{
																var e = e833
																_ = e
																return__1303 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e833)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1300)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1300, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (subvec [1 2] 1) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1304 interface{}) {
												defer func() {
													if e834 := recover(); e834 != nil {
														if func() bool { _, instanceof := e834.(*js.Error); return instanceof }() {
															{
																var e = e834
																_ = e
																return__1304 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e834)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1300, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1300)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (transient [1 2]) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1305 interface{}) {
												defer func() {
													if e835 := recover(); e835 != nil {
														if func() bool { _, instanceof := e835.(*js.Error); return instanceof }() {
															{
																var e = e835
																_ = e
																return__1305 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e835)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1300)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1300, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (transient [1 2]) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1306 interface{}) {
												defer func() {
													if e836 := recover(); e836 != nil {
														if func() bool { _, instanceof := e836.(*js.Error); return instanceof }() {
															{
																var e = e836
																_ = e
																return__1306 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e836)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1300, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1300)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (range 1 3) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1307 interface{}) {
												defer func() {
													if e837 := recover(); e837 != nil {
														if func() bool { _, instanceof := e837.(*js.Error); return instanceof }() {
															{
																var e = e837
																_ = e
																return__1307 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e837)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1300)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1300, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (range 1 3) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1308 interface{}) {
												defer func() {
													if e838 := recover(); e838 != nil {
														if func() bool { _, instanceof := e838.(*js.Error); return instanceof }() {
															{
																var e = e838
																_ = e
																return__1308 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e838)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1300, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))")}))
											}
											seq__817_1284, chunk__818_1285, count__819_1286, i__820_1287 = cljs_core.Next.Arity1IQ(seq__817_1298___1), nil, float64(0), float64(0)
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
				var x_1309 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}).X_invoke_Arity1(cljs_core.Meta.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}).X_invoke_Arity1(cljs_core.Deref.X_invoke_Arity1(Cljs_780))))
				_ = x_1309
				if cljs_core.Vector_QMARK_.Arity1IB(x_1309) {
				} else {
					panic((&js.Error{("Assert failed: (vector? x)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(x_1309, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= x [1 2 3])")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3((&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}), "ToString", []interface{}{}), "550e8400-e29b-41d4-a716-446655440000") {
			} else {
				panic((&js.Error{("Assert failed: (= (.toString #uuid \"550e8400-e29b-41d4-a716-446655440000\") \"550e8400-e29b-41d4-a716-446655440000\")")}))
			}
			{
				var seq__839_1310 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__840_1311 interface{} = nil
				var count__841_1312 = float64(0)
				var i__842_1313 = float64(0)
				_, _, _, _ = seq__839_1310, chunk__840_1311, count__841_1312, i__842_1313
				for {
					if i__842_1313 < count__841_1312 {
						{
							var m_1314 = chunk__840_1311.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__842_1313)
							_ = m_1314
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__1315 interface{}) {
								defer func() {
									if e843 := recover(); e843 != nil {
										if func() bool { _, instanceof := e843.(*js.Error); return instanceof }() {
											{
												var ___ = e843
												_ = ___
												return__1315 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
											}
										} else {
											panic(e843)

										}
									}
								}()
								{
									return cljs_core.Conj.X_invoke_Arity2(m_1314, "foo")
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1314, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m [:foo 1]))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1314, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m {:foo 1}))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1314, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m (list [:foo 1])))")}))
							}
							seq__839_1310, chunk__840_1311, count__841_1312, i__842_1313 = seq__839_1310, chunk__840_1311, count__841_1312, (i__842_1313 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1316 = cljs_core.Seq.Arity1IQ(seq__839_1310)
							_ = temp__4222__auto___1316
							if cljs_core.Truth_(temp__4222__auto___1316) {
								{
									var seq__839_1317___1 = temp__4222__auto___1316
									_ = seq__839_1317___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__839_1317___1) {
										{
											var c__947__auto___1318 = cljs_core.Chunk_first.X_invoke_Arity1(seq__839_1317___1)
											_ = c__947__auto___1318
											seq__839_1310, chunk__840_1311, count__841_1312, i__842_1313 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__839_1317___1), c__947__auto___1318, cljs_core.Count.X_invoke_Arity1(c__947__auto___1318).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_1319 = cljs_core.First.X_invoke_Arity1(seq__839_1317___1)
											_ = m_1319
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__1320 interface{}) {
												defer func() {
													if e844 := recover(); e844 != nil {
														if func() bool { _, instanceof := e844.(*js.Error); return instanceof }() {
															{
																var ___ = e844
																_ = ___
																return__1320 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
															}
														} else {
															panic(e844)

														}
													}
												}()
												{
													return cljs_core.Conj.X_invoke_Arity2(m_1319, "foo")
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1319, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m [:foo 1]))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1319, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m {:foo 1}))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1319, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m (list [:foo 1])))")}))
											}
											seq__839_1310, chunk__840_1311, count__841_1312, i__842_1313 = cljs_core.Next.Arity1IQ(seq__839_1317___1), nil, float64(0), float64(0)
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
				var seq__845_1321 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Array_map, cljs_core.Hash_map, cljs_core.Sorted_map}, nil}))
				var chunk__846_1322 interface{} = nil
				var count__847_1323 = float64(0)
				var i__848_1324 = float64(0)
				_, _, _, _ = seq__845_1321, chunk__846_1322, count__847_1323, i__848_1324
				for {
					if i__848_1324 < count__847_1323 {
						{
							var mt_1325 = chunk__846_1322.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__848_1324)
							_ = mt_1325
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
								var G__849 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
								var G__850 = float64(1)
								_, _ = G__849, G__850
								return mt_1325.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__849, G__850)
							}(), func(make_seq *cljs_core.AFn, seq__845_1321 interface{}, chunk__846_1322 interface{}, count__847_1323 float64, i__848_1324 float64, mt_1325 interface{}) *cljs_core.AFn {
								return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
									if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
										X__GT_t856 = func(__GT_t856 *cljs_core.AFn, seq__845_1321 interface{}, chunk__846_1322 interface{}, count__847_1323 float64, i__848_1324 float64, mt_1325 interface{}) *cljs_core.AFn {
											return cljs_core.Fn(__GT_t856, 9, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, i__848___1 interface{}, count__847___1 interface{}, chunk__846___1 interface{}, seq__845___1 interface{}, test_stuff___1 interface{}, meta857 interface{}) interface{} {
												return (&CljsCore_testT856{from_seq___1, make_seq___1, mt___1, i__848___1, count__847___1, chunk__846___1, seq__845___1, test_stuff___1, meta857})
											})
										}(&cljs_core.AFn{}, seq__845_1321, chunk__846_1322, count__847_1323, i__848_1324, mt_1325)

										return (&CljsCore_testT856{from_seq, make_seq, mt_1325, i__848_1324, count__847_1323, chunk__846_1322, seq__845_1321, test_stuff, nil})
									} else {
										return nil
									}
								})
							}(&cljs_core.AFn{}, seq__845_1321, chunk__846_1322, count__847_1323, i__848_1324, mt_1325).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))")}))
							}
							seq__845_1321, chunk__846_1322, count__847_1323, i__848_1324 = seq__845_1321, chunk__846_1322, count__847_1323, (i__848_1324 + float64(1))
							continue
						}
					} else {
						{
							var temp__4222__auto___1326 = cljs_core.Seq.Arity1IQ(seq__845_1321)
							_ = temp__4222__auto___1326
							if cljs_core.Truth_(temp__4222__auto___1326) {
								{
									var seq__845_1327___1 = temp__4222__auto___1326
									_ = seq__845_1327___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__845_1327___1) {
										{
											var c__947__auto___1328 = cljs_core.Chunk_first.X_invoke_Arity1(seq__845_1327___1)
											_ = c__947__auto___1328
											seq__845_1321, chunk__846_1322, count__847_1323, i__848_1324 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__845_1327___1), c__947__auto___1328, cljs_core.Count.X_invoke_Arity1(c__947__auto___1328).(float64), float64(0)
											continue
										}
									} else {
										{
											var mt_1329 = cljs_core.First.X_invoke_Arity1(seq__845_1327___1)
											_ = mt_1329
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
												var G__861 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
												var G__862 = float64(1)
												_, _ = G__861, G__862
												return mt_1329.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__861, G__862)
											}(), func(make_seq *cljs_core.AFn, seq__845_1321 interface{}, chunk__846_1322 interface{}, count__847_1323 float64, i__848_1324 float64, mt_1329 interface{}, seq__845_1327___1 interface{}, temp__4222__auto___1326 cljs_core.CljsCoreISeq) *cljs_core.AFn {
												return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
													if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
														X__GT_t868 = func(__GT_t868 *cljs_core.AFn, seq__845_1321 interface{}, chunk__846_1322 interface{}, count__847_1323 float64, i__848_1324 float64, mt_1329 interface{}, seq__845_1327___1 interface{}, temp__4222__auto___1326 cljs_core.CljsCoreISeq) *cljs_core.AFn {
															return cljs_core.Fn(__GT_t868, 10, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, temp__4222__auto_____1 interface{}, i__848___1 interface{}, count__847___1 interface{}, chunk__846___1 interface{}, seq__845___2 interface{}, test_stuff___1 interface{}, meta869 interface{}) interface{} {
																return (&CljsCore_testT868{from_seq___1, make_seq___1, mt___1, temp__4222__auto_____1, i__848___1, count__847___1, chunk__846___1, seq__845___2, test_stuff___1, meta869})
															})
														}(&cljs_core.AFn{}, seq__845_1321, chunk__846_1322, count__847_1323, i__848_1324, mt_1329, seq__845_1327___1, temp__4222__auto___1326)

														return (&CljsCore_testT868{from_seq, make_seq, mt_1329, temp__4222__auto___1326, i__848_1324, count__847_1323, chunk__846_1322, seq__845_1327___1, test_stuff, nil})
													} else {
														return nil
													}
												})
											}(&cljs_core.AFn{}, seq__845_1321, chunk__846_1322, count__847_1323, i__848_1324, mt_1329, seq__845_1327___1, temp__4222__auto___1326).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))")}))
											}
											seq__845_1321, chunk__846_1322, count__847_1323, i__848_1324 = cljs_core.Next.Arity1IQ(seq__845_1327___1), nil, float64(0), float64(0)
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
				var _STAR_print_length_STAR_873 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_873
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_873

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
				var _STAR_print_length_STAR_874 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_874
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_874

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
				var _STAR_print_length_STAR_875 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_875
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_875

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
				var _STAR_print_length_STAR_876 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_876
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_876

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
				var _STAR_print_length_STAR_877 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_877
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_877

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
				var _STAR_print_length_STAR_878 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_878
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_878

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
							var G__880 = func() interface{} {
								if func() bool { _, instanceof := value.(*cljs_core.CljsCoreKeyword); return instanceof }() {
									return cljs_core.Native_get_instance_field.X_invoke_Arity2(value, "Fqn")
								} else {
									return nil
								}
							}()
							_ = G__880
							switch G__880 {
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
				var not_strings_1331 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, nil, float64(1), func(G__1332 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1332, 0, func() interface{} {
						return nil
					})
				}(&cljs_core.AFn{})}, nil})
				_ = not_strings_1331
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1333 *cljs_core.AFn, not_strings_1331 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1333, 1, func(p1__79_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1334 interface{}) {
							defer func() {
								if e881 := recover(); e881 != nil {
									if func() bool { _, instanceof := e881.(*js.TypeError); return instanceof }() {
										{
											var ___ = e881
											_ = ___
											return__1334 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e881)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__79_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1331), not_strings_1331) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__79#] (= :failed (try (re-find #\".\" p1__79#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1335 *cljs_core.AFn, not_strings_1331 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1335, 1, func(p1__80_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1336 interface{}) {
							defer func() {
								if e882 := recover(); e882 != nil {
									if func() bool { _, instanceof := e882.(*js.TypeError); return instanceof }() {
										{
											var ___ = e882
											_ = ___
											return__1336 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e882)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__80_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1331), not_strings_1331) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__80#] (= :failed (try (re-matches #\".\" p1__80#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1337 *cljs_core.AFn, not_strings_1331 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1337, 1, func(p1__81_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1338 interface{}) {
							defer func() {
								if e883 := recover(); e883 != nil {
									if func() bool { _, instanceof := e883.(*js.TypeError); return instanceof }() {
										{
											var ___ = e883
											_ = ___
											return__1338 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e883)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__81_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1331), not_strings_1331) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__81#] (= :failed (try (re-find #\"nomatch\" p1__81#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1339 *cljs_core.AFn, not_strings_1331 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1339, 1, func(p1__82_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1340 interface{}) {
							defer func() {
								if e884 := recover(); e884 != nil {
									if func() bool { _, instanceof := e884.(*js.TypeError); return instanceof }() {
										{
											var ___ = e884
											_ = ___
											return__1340 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e884)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__82_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1331), not_strings_1331) {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(func(G__1341 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1341, 1, func(p1__83_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Take_while.X_invoke_Arity1(func(G__1342 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1342, 1, func(p1__84_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Drop_while.X_invoke_Arity1(func(G__1343 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1343, 1, func(p1__85_SHARP_ interface{}) interface{} {
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
				var ret_1344 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))
				_ = ret_1344
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_1344)) && (cljs_core.X_EQ_.Arity2IIB(ret_1344, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList))) {
				} else {
					panic((&js.Error{("Assert failed: (and (vector? ret) (= ret (quote (1 2 3))))")}))
				}
			}
			{
				var ret_1345 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				_ = ret_1345
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_1345)) && (cljs_core.X_EQ_.Arity2IIB(ret_1345, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(2), float64(4), float64(6), float64(8)})).(*cljs_core.CljsCoreList))) {
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
			Xform = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Dedupe.X_invoke_Arity0().(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Partition_all.X_invoke_Arity1(float64(3)).(cljs_core.CljsCoreIFn), cljs_core.Partition_by.X_invoke_Arity1(func(G__1346 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1346, 1, func(p1__86_SHARP_ interface{}) interface{} {
					return (cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, p1__86_SHARP_).(float64) < float64(7))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Flatten).(cljs_core.CljsCoreIFn), cljs_core.Random_sample.X_invoke_Arity1(1.0).(cljs_core.CljsCoreIFn), cljs_core.Take_nth.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Keep.X_invoke_Arity1(func(G__1347 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1347, 1, func(p1__87_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__87_SHARP_) {
						return (p1__87_SHARP_.(float64) * p1__87_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Keep_indexed.X_invoke_Arity1(func(G__1348 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1348, 2, func(p1__88_SHARP_ interface{}, p2__89_SHARP_ interface{}) interface{} {
					if cljs_core.Even_QMARK_.Arity1IB(p1__88_SHARP_) {
						return (p1__88_SHARP_.(float64) * p2__89_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Replace.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(2), "two", float64(6), "six", float64(18), "eighteen"}, nil})).(cljs_core.CljsCoreIFn), cljs_core.Take.X_invoke_Arity1(float64(11)).(cljs_core.CljsCoreIFn), cljs_core.Take_while.X_invoke_Arity1(func(G__1349 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1349, 1, func(p1__90_SHARP_ interface{}) interface{} {
					return cljs_core.Not_EQ_.Arity2IIB(float64(300), p1__90_SHARP_)
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Drop.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Drop_while.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Remove.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn)})).(*cljs_core.AFn)

			Data = cljs_core.Vec.X_invoke_Arity1(cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(18)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(20)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(Xform, Data), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(36), float64(200), float64(10)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence xform data) (quote (36 200 10)))")}))
			}
			Xf = cljs_core.Map_.X_invoke_Arity1(func(G__1350 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1350, 2, func(p1__91_SHARP_ interface{}, p2__92_SHARP_ interface{}) interface{} {
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
	return (func() bool { _, instanceof := other.(*CljsCore_testFixedHash); return instanceof }()) && (cljs_core.X_EQ_.Arity2IIB(this.V, cljs_core.Native_get_instance_field.X_invoke_Arity2(other, "V")))
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

func (this__751__auto__ *CljsCore_testPerson) X_lookup_Arity2(k__752__auto__ interface{}) interface{} {
	return this__751__auto__.X_lookup_Arity3(k__752__auto__, nil)
}

func (this__753__auto__ *CljsCore_testPerson) X_lookup_Arity3(k633 interface{}, else__754__auto__ interface{}) interface{} {
	{
		var G__636 = func() interface{} {
			if func() bool { _, instanceof := k633.(*cljs_core.CljsCoreKeyword); return instanceof }() {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k633, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__636
		switch G__636 {
		case "lastname":
			return this__753__auto__.Lastname

		case "firstname":
			return this__753__auto__.Firstname

		default:
			return cljs_core.Get.X_invoke_Arity3(this__753__auto__.X__extmap, k633, else__754__auto__)

		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIPrintWithWriter__() {}

func (this__767__auto__ *CljsCore_testPerson) X_pr_writer_Arity3(writer__768__auto__ interface{}, opts__769__auto__ interface{}) interface{} {
	{
		var pr_pair__770__auto__ = func(G__1352 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1352, 3, func(keyval__771__auto__ interface{}, ___772__auto__ interface{}, ___772__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, cljs_core.Pr_writer, "", " ", "", opts__769__auto__, keyval__771__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__770__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, pr_pair__770__auto__, "#cljs.core-test.Person{", ", ", "}", opts__769__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), this__767__auto__.Firstname}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), this__767__auto__.Lastname}, nil})}, nil}), this__767__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testPerson) CljsCoreIMeta__() {}

func (this__749__auto__ *CljsCore_testPerson) X_meta_Arity1() interface{} {
	return this__749__auto__.X__meta
}

func (_ *CljsCore_testPerson) CljsCoreICloneable__() {}

func (this__745__auto__ *CljsCore_testPerson) X_clone_Arity1() interface{} {
	return (&CljsCore_testPerson{this__745__auto__.Firstname, this__745__auto__.Lastname, this__745__auto__.X__meta, this__745__auto__.X__extmap, this__745__auto__.X__hash})
}

func (_ *CljsCore_testPerson) CljsCoreICounted__() {}

func (this__755__auto__ *CljsCore_testPerson) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__755__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testPerson) CljsCoreIHash__() {}

func (this__746__auto__ *CljsCore_testPerson) X_hash_Arity1() interface{} {
	{
		var h__579__auto__ = this__746__auto__.X__hash
		_ = h__579__auto__
		if !(cljs_core.Nil_(h__579__auto__)) {
			return h__579__auto__
		} else {
			{
				var h__579__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__746__auto__).(float64)
				_ = h__579__auto_____1
				this__746__auto__.X__hash = h__579__auto_____1

				return h__579__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIEquiv__() {}

func (this__747__auto__ *CljsCore_testPerson) X_equiv_Arity2(other__748__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__748__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__747__auto__), cljs_core.Type_.X_invoke_Arity1(other__748__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__747__auto__, other__748__auto__)))
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

func (this__762__auto__ *CljsCore_testPerson) X_dissoc_Arity2(k__763__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), nil}, nil}, nil}), k__763__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__762__auto__), this__762__auto__.X__meta), k__763__auto__)
	} else {
		return (&CljsCore_testPerson{this__762__auto__.Firstname, this__762__auto__.Lastname, this__762__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__762__auto__.X__extmap, k__763__auto__)), nil})
	}
}

func (_ *CljsCore_testPerson) CljsCoreIAssociative__() {}

func (this__758__auto__ *CljsCore_testPerson) X_assoc_Arity3(k__759__auto__ interface{}, G__632 interface{}) interface{} {
	{
		var pred__644 = cljs_core.Keyword_identical_QMARK_
		var expr__645 = k__759__auto__
		_, _ = pred__644, expr__645
		if cljs_core.Truth_(func() interface{} {
			var G__647 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)})
			var G__648 = expr__645
			_, _ = G__647, G__648
			return pred__644.X_invoke_Arity2(G__647, G__648)
		}()) {
			return (&CljsCore_testPerson{G__632, this__758__auto__.Lastname, this__758__auto__.X__meta, this__758__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__649 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})
				var G__650 = expr__645
				_, _ = G__649, G__650
				return pred__644.X_invoke_Arity2(G__649, G__650)
			}()) {
				return (&CljsCore_testPerson{this__758__auto__.Firstname, G__632, this__758__auto__.X__meta, this__758__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPerson{this__758__auto__.Firstname, this__758__auto__.Lastname, this__758__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__758__auto__.X__extmap, k__759__auto__, G__632), nil})
			}
		}
	}
}

func (this__760__auto__ *CljsCore_testPerson) X_contains_key_QMARK__Arity2(k__761__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testPerson) CljsCoreISeqable__() {}

func (this__765__auto__ *CljsCore_testPerson) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), this__765__auto__.Firstname}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), this__765__auto__.Lastname}, nil})}, nil}), this__765__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testPerson) CljsCoreIWithMeta__() {}

func (this__750__auto__ *CljsCore_testPerson) X_with_meta_Arity2(G__632 interface{}) interface{} {
	return (&CljsCore_testPerson{this__750__auto__.Firstname, this__750__auto__.Lastname, G__632, this__750__auto__.X__extmap, this__750__auto__.X__hash})
}

func (_ *CljsCore_testPerson) CljsCoreICollection__() {}

func (this__756__auto__ *CljsCore_testPerson) X_conj_Arity2(entry__757__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__757__auto__) {
		return this__756__auto__.X_assoc_Arity3(entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__756__auto__, entry__757__auto__)
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

func (this__751__auto__ *CljsCore_testA) X_lookup_Arity2(k__752__auto__ interface{}) interface{} {
	return this__751__auto__.X_lookup_Arity3(k__752__auto__, nil)
}

func (this__753__auto__ *CljsCore_testA) X_lookup_Arity3(k652 interface{}, else__754__auto__ interface{}) interface{} {
	{
		var G__655 = k652
		_ = G__655
		switch G__655 {
		default:
			return cljs_core.Get.X_invoke_Arity3(this__753__auto__.X__extmap, k652, else__754__auto__)

		}
	}
}

func (_ *CljsCore_testA) CljsCoreIPrintWithWriter__() {}

func (this__767__auto__ *CljsCore_testA) X_pr_writer_Arity3(writer__768__auto__ interface{}, opts__769__auto__ interface{}) interface{} {
	{
		var pr_pair__770__auto__ = func(G__1354 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1354, 3, func(keyval__771__auto__ interface{}, ___772__auto__ interface{}, ___772__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, cljs_core.Pr_writer, "", " ", "", opts__769__auto__, keyval__771__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__770__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, pr_pair__770__auto__, "#cljs.core-test.A{", ", ", "}", opts__769__auto__, cljs_core.Concat.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, this__767__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testA) CljsCoreIMeta__() {}

func (this__749__auto__ *CljsCore_testA) X_meta_Arity1() interface{} {
	return this__749__auto__.X__meta
}

func (_ *CljsCore_testA) CljsCoreICloneable__() {}

func (this__745__auto__ *CljsCore_testA) X_clone_Arity1() interface{} {
	return (&CljsCore_testA{this__745__auto__.X__meta, this__745__auto__.X__extmap, this__745__auto__.X__hash})
}

func (_ *CljsCore_testA) CljsCoreICounted__() {}

func (this__755__auto__ *CljsCore_testA) X_count_Arity1() float64 {
	return (float64(0) + cljs_core.Count.X_invoke_Arity1(this__755__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testA) CljsCoreIHash__() {}

func (this__746__auto__ *CljsCore_testA) X_hash_Arity1() interface{} {
	{
		var h__579__auto__ = this__746__auto__.X__hash
		_ = h__579__auto__
		if !(cljs_core.Nil_(h__579__auto__)) {
			return h__579__auto__
		} else {
			{
				var h__579__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__746__auto__).(float64)
				_ = h__579__auto_____1
				this__746__auto__.X__hash = h__579__auto_____1

				return h__579__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testA) CljsCoreIEquiv__() {}

func (this__747__auto__ *CljsCore_testA) X_equiv_Arity2(other__748__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__748__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__747__auto__), cljs_core.Type_.X_invoke_Arity1(other__748__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__747__auto__, other__748__auto__)))
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

func (this__762__auto__ *CljsCore_testA) X_dissoc_Arity2(k__763__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, k__763__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__762__auto__), this__762__auto__.X__meta), k__763__auto__)
	} else {
		return (&CljsCore_testA{this__762__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__762__auto__.X__extmap, k__763__auto__)), nil})
	}
}

func (_ *CljsCore_testA) CljsCoreIAssociative__() {}

func (this__758__auto__ *CljsCore_testA) X_assoc_Arity3(k__759__auto__ interface{}, G__651 interface{}) interface{} {
	{
		var pred__659 = cljs_core.Keyword_identical_QMARK_
		var expr__660 = k__759__auto__
		_, _ = pred__659, expr__660
		return (&CljsCore_testA{this__758__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__758__auto__.X__extmap, k__759__auto__, G__651), nil})
	}
}

func (this__760__auto__ *CljsCore_testA) X_contains_key_QMARK__Arity2(k__761__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testA) CljsCoreISeqable__() {}

func (this__765__auto__ *CljsCore_testA) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, this__765__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testA) CljsCoreIWithMeta__() {}

func (this__750__auto__ *CljsCore_testA) X_with_meta_Arity2(G__651 interface{}) interface{} {
	return (&CljsCore_testA{G__651, this__750__auto__.X__extmap, this__750__auto__.X__hash})
}

func (_ *CljsCore_testA) CljsCoreICollection__() {}

func (this__756__auto__ *CljsCore_testA) X_conj_Arity2(entry__757__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__757__auto__) {
		return this__756__auto__.X_assoc_Arity3(entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__756__auto__, entry__757__auto__)
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

func (this__751__auto__ *CljsCore_testC) X_lookup_Arity2(k__752__auto__ interface{}) interface{} {
	return this__751__auto__.X_lookup_Arity3(k__752__auto__, nil)
}

func (this__753__auto__ *CljsCore_testC) X_lookup_Arity3(k663 interface{}, else__754__auto__ interface{}) interface{} {
	{
		var G__666 = func() interface{} {
			if func() bool { _, instanceof := k663.(*cljs_core.CljsCoreKeyword); return instanceof }() {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k663, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__666
		switch G__666 {
		case "c":
			return this__753__auto__.C

		case "b":
			return this__753__auto__.B

		case "a":
			return this__753__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__753__auto__.X__extmap, k663, else__754__auto__)

		}
	}
}

func (_ *CljsCore_testC) CljsCoreIPrintWithWriter__() {}

func (this__767__auto__ *CljsCore_testC) X_pr_writer_Arity3(writer__768__auto__ interface{}, opts__769__auto__ interface{}) interface{} {
	{
		var pr_pair__770__auto__ = func(G__1356 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1356, 3, func(keyval__771__auto__ interface{}, ___772__auto__ interface{}, ___772__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, cljs_core.Pr_writer, "", " ", "", opts__769__auto__, keyval__771__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__770__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, pr_pair__770__auto__, "#cljs.core-test.C{", ", ", "}", opts__769__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__767__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__767__auto__.B}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), this__767__auto__.C}, nil})}, nil}), this__767__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testC) CljsCoreIMeta__() {}

func (this__749__auto__ *CljsCore_testC) X_meta_Arity1() interface{} {
	return this__749__auto__.X__meta
}

func (_ *CljsCore_testC) CljsCoreICloneable__() {}

func (this__745__auto__ *CljsCore_testC) X_clone_Arity1() interface{} {
	return (&CljsCore_testC{this__745__auto__.A, this__745__auto__.B, this__745__auto__.C, this__745__auto__.X__meta, this__745__auto__.X__extmap, this__745__auto__.X__hash})
}

func (_ *CljsCore_testC) CljsCoreICounted__() {}

func (this__755__auto__ *CljsCore_testC) X_count_Arity1() float64 {
	return (float64(3) + cljs_core.Count.X_invoke_Arity1(this__755__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testC) CljsCoreIHash__() {}

func (this__746__auto__ *CljsCore_testC) X_hash_Arity1() interface{} {
	{
		var h__579__auto__ = this__746__auto__.X__hash
		_ = h__579__auto__
		if !(cljs_core.Nil_(h__579__auto__)) {
			return h__579__auto__
		} else {
			{
				var h__579__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__746__auto__).(float64)
				_ = h__579__auto_____1
				this__746__auto__.X__hash = h__579__auto_____1

				return h__579__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testC) CljsCoreIEquiv__() {}

func (this__747__auto__ *CljsCore_testC) X_equiv_Arity2(other__748__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__748__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__747__auto__), cljs_core.Type_.X_invoke_Arity1(other__748__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__747__auto__, other__748__auto__)))
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

func (this__762__auto__ *CljsCore_testC) X_dissoc_Arity2(k__763__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__763__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__762__auto__), this__762__auto__.X__meta), k__763__auto__)
	} else {
		return (&CljsCore_testC{this__762__auto__.A, this__762__auto__.B, this__762__auto__.C, this__762__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__762__auto__.X__extmap, k__763__auto__)), nil})
	}
}

func (_ *CljsCore_testC) CljsCoreIAssociative__() {}

func (this__758__auto__ *CljsCore_testC) X_assoc_Arity3(k__759__auto__ interface{}, G__662 interface{}) interface{} {
	{
		var pred__676 = cljs_core.Keyword_identical_QMARK_
		var expr__677 = k__759__auto__
		_, _ = pred__676, expr__677
		if cljs_core.Truth_(func() interface{} {
			var G__679 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__680 = expr__677
			_, _ = G__679, G__680
			return pred__676.X_invoke_Arity2(G__679, G__680)
		}()) {
			return (&CljsCore_testC{G__662, this__758__auto__.B, this__758__auto__.C, this__758__auto__.X__meta, this__758__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__681 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__682 = expr__677
				_, _ = G__681, G__682
				return pred__676.X_invoke_Arity2(G__681, G__682)
			}()) {
				return (&CljsCore_testC{this__758__auto__.A, G__662, this__758__auto__.C, this__758__auto__.X__meta, this__758__auto__.X__extmap, nil})
			} else {
				if cljs_core.Truth_(func() interface{} {
					var G__683 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})
					var G__684 = expr__677
					_, _ = G__683, G__684
					return pred__676.X_invoke_Arity2(G__683, G__684)
				}()) {
					return (&CljsCore_testC{this__758__auto__.A, this__758__auto__.B, G__662, this__758__auto__.X__meta, this__758__auto__.X__extmap, nil})
				} else {
					return (&CljsCore_testC{this__758__auto__.A, this__758__auto__.B, this__758__auto__.C, this__758__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__758__auto__.X__extmap, k__759__auto__, G__662), nil})
				}
			}
		}
	}
}

func (this__760__auto__ *CljsCore_testC) X_contains_key_QMARK__Arity2(k__761__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testC) CljsCoreISeqable__() {}

func (this__765__auto__ *CljsCore_testC) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__765__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__765__auto__.B}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), this__765__auto__.C}, nil})}, nil}), this__765__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testC) CljsCoreIWithMeta__() {}

func (this__750__auto__ *CljsCore_testC) X_with_meta_Arity2(G__662 interface{}) interface{} {
	return (&CljsCore_testC{this__750__auto__.A, this__750__auto__.B, this__750__auto__.C, G__662, this__750__auto__.X__extmap, this__750__auto__.X__hash})
}

func (_ *CljsCore_testC) CljsCoreICollection__() {}

func (this__756__auto__ *CljsCore_testC) X_conj_Arity2(entry__757__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__757__auto__) {
		return this__756__auto__.X_assoc_Arity3(entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__756__auto__, entry__757__auto__)
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

func (this__751__auto__ *CljsCore_testA2) X_lookup_Arity2(k__752__auto__ interface{}) interface{} {
	return this__751__auto__.X_lookup_Arity3(k__752__auto__, nil)
}

func (this__753__auto__ *CljsCore_testA2) X_lookup_Arity3(k688 interface{}, else__754__auto__ interface{}) interface{} {
	{
		var G__691 = func() interface{} {
			if func() bool { _, instanceof := k688.(*cljs_core.CljsCoreKeyword); return instanceof }() {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k688, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__691
		switch G__691 {
		case "x":
			return this__753__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__753__auto__.X__extmap, k688, else__754__auto__)

		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIPrintWithWriter__() {}

func (this__767__auto__ *CljsCore_testA2) X_pr_writer_Arity3(writer__768__auto__ interface{}, opts__769__auto__ interface{}) interface{} {
	{
		var pr_pair__770__auto__ = func(G__1358 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1358, 3, func(keyval__771__auto__ interface{}, ___772__auto__ interface{}, ___772__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, cljs_core.Pr_writer, "", " ", "", opts__769__auto__, keyval__771__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__770__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, pr_pair__770__auto__, "#cljs.core-test.A2{", ", ", "}", opts__769__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__767__auto__.X}, nil})}, nil}), this__767__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testA2) CljsCoreIMeta__() {}

func (this__749__auto__ *CljsCore_testA2) X_meta_Arity1() interface{} {
	return this__749__auto__.X__meta
}

func (_ *CljsCore_testA2) CljsCoreICloneable__() {}

func (this__745__auto__ *CljsCore_testA2) X_clone_Arity1() interface{} {
	return (&CljsCore_testA2{this__745__auto__.X, this__745__auto__.X__meta, this__745__auto__.X__extmap, this__745__auto__.X__hash})
}

func (_ *CljsCore_testA2) CljsCoreICounted__() {}

func (this__755__auto__ *CljsCore_testA2) X_count_Arity1() float64 {
	return (float64(1) + cljs_core.Count.X_invoke_Arity1(this__755__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testA2) CljsCoreIHash__() {}

func (this__746__auto__ *CljsCore_testA2) X_hash_Arity1() interface{} {
	{
		var h__579__auto__ = this__746__auto__.X__hash
		_ = h__579__auto__
		if !(cljs_core.Nil_(h__579__auto__)) {
			return h__579__auto__
		} else {
			{
				var h__579__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__746__auto__).(float64)
				_ = h__579__auto_____1
				this__746__auto__.X__hash = h__579__auto_____1

				return h__579__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIEquiv__() {}

func (this__747__auto__ *CljsCore_testA2) X_equiv_Arity2(other__748__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__748__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__747__auto__), cljs_core.Type_.X_invoke_Arity1(other__748__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__747__auto__, other__748__auto__)))
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

func (this__762__auto__ *CljsCore_testA2) X_dissoc_Arity2(k__763__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), nil}, nil}, nil}), k__763__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__762__auto__), this__762__auto__.X__meta), k__763__auto__)
	} else {
		return (&CljsCore_testA2{this__762__auto__.X, this__762__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__762__auto__.X__extmap, k__763__auto__)), nil})
	}
}

func (_ *CljsCore_testA2) CljsCoreIAssociative__() {}

func (this__758__auto__ *CljsCore_testA2) X_assoc_Arity3(k__759__auto__ interface{}, G__687 interface{}) interface{} {
	{
		var pred__697 = cljs_core.Keyword_identical_QMARK_
		var expr__698 = k__759__auto__
		_, _ = pred__697, expr__698
		if cljs_core.Truth_(func() interface{} {
			var G__700 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__701 = expr__698
			_, _ = G__700, G__701
			return pred__697.X_invoke_Arity2(G__700, G__701)
		}()) {
			return (&CljsCore_testA2{G__687, this__758__auto__.X__meta, this__758__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testA2{this__758__auto__.X, this__758__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__758__auto__.X__extmap, k__759__auto__, G__687), nil})
		}
	}
}

func (this__760__auto__ *CljsCore_testA2) X_contains_key_QMARK__Arity2(k__761__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testA2) CljsCoreISeqable__() {}

func (this__765__auto__ *CljsCore_testA2) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__765__auto__.X}, nil})}, nil}), this__765__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testA2) CljsCoreIWithMeta__() {}

func (this__750__auto__ *CljsCore_testA2) X_with_meta_Arity2(G__687 interface{}) interface{} {
	return (&CljsCore_testA2{this__750__auto__.X, G__687, this__750__auto__.X__extmap, this__750__auto__.X__hash})
}

func (_ *CljsCore_testA2) CljsCoreICollection__() {}

func (this__756__auto__ *CljsCore_testA2) X_conj_Arity2(entry__757__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__757__auto__) {
		return this__756__auto__.X_assoc_Arity3(entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__756__auto__, entry__757__auto__)
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

func (this__751__auto__ *CljsCore_testB) X_lookup_Arity2(k__752__auto__ interface{}) interface{} {
	return this__751__auto__.X_lookup_Arity3(k__752__auto__, nil)
}

func (this__753__auto__ *CljsCore_testB) X_lookup_Arity3(k703 interface{}, else__754__auto__ interface{}) interface{} {
	{
		var G__706 = func() interface{} {
			if func() bool { _, instanceof := k703.(*cljs_core.CljsCoreKeyword); return instanceof }() {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k703, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__706
		switch G__706 {
		case "x":
			return this__753__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__753__auto__.X__extmap, k703, else__754__auto__)

		}
	}
}

func (_ *CljsCore_testB) CljsCoreIPrintWithWriter__() {}

func (this__767__auto__ *CljsCore_testB) X_pr_writer_Arity3(writer__768__auto__ interface{}, opts__769__auto__ interface{}) interface{} {
	{
		var pr_pair__770__auto__ = func(G__1360 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1360, 3, func(keyval__771__auto__ interface{}, ___772__auto__ interface{}, ___772__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, cljs_core.Pr_writer, "", " ", "", opts__769__auto__, keyval__771__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__770__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, pr_pair__770__auto__, "#cljs.core-test.B{", ", ", "}", opts__769__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__767__auto__.X}, nil})}, nil}), this__767__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testB) CljsCoreIMeta__() {}

func (this__749__auto__ *CljsCore_testB) X_meta_Arity1() interface{} {
	return this__749__auto__.X__meta
}

func (_ *CljsCore_testB) CljsCoreICloneable__() {}

func (this__745__auto__ *CljsCore_testB) X_clone_Arity1() interface{} {
	return (&CljsCore_testB{this__745__auto__.X, this__745__auto__.X__meta, this__745__auto__.X__extmap, this__745__auto__.X__hash})
}

func (_ *CljsCore_testB) CljsCoreICounted__() {}

func (this__755__auto__ *CljsCore_testB) X_count_Arity1() float64 {
	return (float64(1) + cljs_core.Count.X_invoke_Arity1(this__755__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testB) CljsCoreIHash__() {}

func (this__746__auto__ *CljsCore_testB) X_hash_Arity1() interface{} {
	{
		var h__579__auto__ = this__746__auto__.X__hash
		_ = h__579__auto__
		if !(cljs_core.Nil_(h__579__auto__)) {
			return h__579__auto__
		} else {
			{
				var h__579__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__746__auto__).(float64)
				_ = h__579__auto_____1
				this__746__auto__.X__hash = h__579__auto_____1

				return h__579__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testB) CljsCoreIEquiv__() {}

func (this__747__auto__ *CljsCore_testB) X_equiv_Arity2(other__748__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__748__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__747__auto__), cljs_core.Type_.X_invoke_Arity1(other__748__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__747__auto__, other__748__auto__)))
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

func (this__762__auto__ *CljsCore_testB) X_dissoc_Arity2(k__763__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), nil}, nil}, nil}), k__763__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__762__auto__), this__762__auto__.X__meta), k__763__auto__)
	} else {
		return (&CljsCore_testB{this__762__auto__.X, this__762__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__762__auto__.X__extmap, k__763__auto__)), nil})
	}
}

func (_ *CljsCore_testB) CljsCoreIAssociative__() {}

func (this__758__auto__ *CljsCore_testB) X_assoc_Arity3(k__759__auto__ interface{}, G__702 interface{}) interface{} {
	{
		var pred__712 = cljs_core.Keyword_identical_QMARK_
		var expr__713 = k__759__auto__
		_, _ = pred__712, expr__713
		if cljs_core.Truth_(func() interface{} {
			var G__715 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__716 = expr__713
			_, _ = G__715, G__716
			return pred__712.X_invoke_Arity2(G__715, G__716)
		}()) {
			return (&CljsCore_testB{G__702, this__758__auto__.X__meta, this__758__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testB{this__758__auto__.X, this__758__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__758__auto__.X__extmap, k__759__auto__, G__702), nil})
		}
	}
}

func (this__760__auto__ *CljsCore_testB) X_contains_key_QMARK__Arity2(k__761__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testB) CljsCoreISeqable__() {}

func (this__765__auto__ *CljsCore_testB) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__765__auto__.X}, nil})}, nil}), this__765__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testB) CljsCoreIWithMeta__() {}

func (this__750__auto__ *CljsCore_testB) X_with_meta_Arity2(G__702 interface{}) interface{} {
	return (&CljsCore_testB{this__750__auto__.X, G__702, this__750__auto__.X__extmap, this__750__auto__.X__hash})
}

func (_ *CljsCore_testB) CljsCoreICollection__() {}

func (this__756__auto__ *CljsCore_testB) X_conj_Arity2(entry__757__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__757__auto__) {
		return this__756__auto__.X_assoc_Arity3(entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__756__auto__, entry__757__auto__)
	}
}

var X__GT_B *cljs_core.AFn

var Map__GT_B *cljs_core.AFn

type CljsCore_testIFoo interface {
	CljsCore_testIFoo__()
	Foo_Arity1() interface{}
}

func init() {
	cljs_core.RegisterProtocol_("cljs.core-test/IFoo", (*CljsCore_testIFoo)(nil))
}

var Foo *cljs_core.AFn

type CljsCore_testT717 struct {
	Test_stuff interface{}
	Meta718    interface{}
}

func (_ *CljsCore_testT717) CljsCore_testIFoo__() {}

func (this *CljsCore_testT717) Foo_Arity1() interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
}

func (_ *CljsCore_testT717) CljsCoreIMeta__() {}

func (_719 *CljsCore_testT717) X_meta_Arity1() interface{} {
	return _719.Meta718
}

func (_ *CljsCore_testT717) CljsCoreIWithMeta__() {}

func (_719 *CljsCore_testT717) X_with_meta_Arity2(meta718___1 interface{}) interface{} {
	return (&CljsCore_testT717{_719.Test_stuff, meta718___1})
}

var X__GT_t717 *cljs_core.AFn

var Foo2 *cljs_core.CljsCoreMultiFn

type CljsCore_testIMutate interface {
	CljsCore_testIMutate__()
	Mutate_Arity1() interface{}
}

func init() {
	cljs_core.RegisterProtocol_("cljs.core-test/IMutate", (*CljsCore_testIMutate)(nil))
}

var Mutate *cljs_core.AFn

type CljsCore_testMutate struct{ A interface{} }

func (_ *CljsCore_testMutate) CljsCore_testIMutate__() {}

func (___ *CljsCore_testMutate) Mutate_Arity1() interface{} {
	return func() interface{} {
		var return__1361 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})
		___.A = return__1361
		return return__1361
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

func (this *CljsCore_testFnLike) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCore_testFnLike) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
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

func (this *CljsCore_testFnLikeB) X_invoke_Arity18(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity19(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCore_testFnLikeB) X_invoke_Arity20(a___1 interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 20"}))
}

var X__GT_FnLikeB *cljs_core.AFn

type CljsCore_testIHasFirst interface {
	CljsCore_testIHasFirst__()
	X_get_first_Arity1() interface{}
}

func init() {
	cljs_core.RegisterProtocol_("cljs.core-test/IHasFirst", (*CljsCore_testIHasFirst)(nil))
}

var X_get_first *cljs_core.AFn

type CljsCore_testIFindsFirst interface {
	CljsCore_testIFindsFirst__()
	X_find_first_Arity2(other interface{}) interface{}
}

func init() {
	cljs_core.RegisterProtocol_("cljs.core-test/IFindsFirst", (*CljsCore_testIFindsFirst)(nil))
}

var X_find_first *cljs_core.AFn

type CljsCore_testFirst struct{ Xs interface{} }

func (_ *CljsCore_testFirst) CljsCore_testIFindsFirst__() {}

func (___ *CljsCore_testFirst) X_find_first_Arity2(p__720 interface{}) interface{} {
	{
		var vec__722 = p__720
		var x = cljs_core.Nth.X_invoke_Arity3(vec__722, float64(0), nil)
		_, _ = vec__722, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCore_testIHasFirst__() {}

func (p__723 *CljsCore_testFirst) X_get_first_Arity1() interface{} {
	{
		var vec__725 = p__723
		var x = cljs_core.Nth.X_invoke_Arity3(vec__725, float64(0), nil)
		_, _ = vec__725, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCoreObject__() {}

func (p__726 *CljsCore_testFirst) ToString() string {
	{
		var vec__728 = p__726
		var x = cljs_core.Nth.X_invoke_Arity3(vec__728, float64(0), nil)
		_, _ = vec__728, x
		return (`` + cljs_core.Str.X_invoke_Arity1(x).(string))
	}
}

func (this *CljsCore_testFirst) String() string {
	return this.ToString()
}

func (_ *CljsCore_testFirst) CljsCoreIFn__() {}

func (p__729 *CljsCore_testFirst) X_invoke_Arity0() interface{} {
	{
		var vec__731 = p__729
		var x = cljs_core.Nth.X_invoke_Arity3(vec__731, float64(0), nil)
		_, _ = vec__731, x
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

func (this *CljsCore_testFirst) X_invoke_Arity18(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 18"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity19(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}) interface{} {
	panic((&js.Error{"Invalid arity: 19"}))
}

func (this *CljsCore_testFirst) X_invoke_Arity20(a interface{}, b interface{}, c interface{}, d interface{}, e interface{}, f interface{}, g interface{}, h interface{}, i interface{}, j interface{}, k interface{}, l interface{}, m interface{}, n interface{}, o interface{}, p interface{}, q interface{}, s interface{}, t interface{}, rest interface{}) interface{} {
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

func (___ *CljsCore_testDestructuringWithLocals) X_find_first_Arity2(p__733 interface{}) interface{} {
	{
		var vec__735 = p__733
		var x = cljs_core.Nth.X_invoke_Arity3(vec__735, float64(0), nil)
		var y = cljs_core.Nth.X_invoke_Arity3(vec__735, float64(1), nil)
		_, _, _ = vec__735, x, y
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

func (this__751__auto__ *CljsCore_testPrintMe) X_lookup_Arity2(k__752__auto__ interface{}) interface{} {
	return this__751__auto__.X_lookup_Arity3(k__752__auto__, nil)
}

func (this__753__auto__ *CljsCore_testPrintMe) X_lookup_Arity3(k747 interface{}, else__754__auto__ interface{}) interface{} {
	{
		var G__750 = func() interface{} {
			if func() bool { _, instanceof := k747.(*cljs_core.CljsCoreKeyword); return instanceof }() {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(k747, "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__750
		switch G__750 {
		case "b":
			return this__753__auto__.B

		case "a":
			return this__753__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__753__auto__.X__extmap, k747, else__754__auto__)

		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIPrintWithWriter__() {}

func (this__767__auto__ *CljsCore_testPrintMe) X_pr_writer_Arity3(writer__768__auto__ interface{}, opts__769__auto__ interface{}) interface{} {
	{
		var pr_pair__770__auto__ = func(G__1363 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1363, 3, func(keyval__771__auto__ interface{}, ___772__auto__ interface{}, ___772__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, cljs_core.Pr_writer, "", " ", "", opts__769__auto__, keyval__771__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__770__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__768__auto__, pr_pair__770__auto__, "#cljs.core-test.PrintMe{", ", ", "}", opts__769__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__767__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__767__auto__.B}, nil})}, nil}), this__767__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIMeta__() {}

func (this__749__auto__ *CljsCore_testPrintMe) X_meta_Arity1() interface{} {
	return this__749__auto__.X__meta
}

func (_ *CljsCore_testPrintMe) CljsCoreICloneable__() {}

func (this__745__auto__ *CljsCore_testPrintMe) X_clone_Arity1() interface{} {
	return (&CljsCore_testPrintMe{this__745__auto__.A, this__745__auto__.B, this__745__auto__.X__meta, this__745__auto__.X__extmap, this__745__auto__.X__hash})
}

func (_ *CljsCore_testPrintMe) CljsCoreICounted__() {}

func (this__755__auto__ *CljsCore_testPrintMe) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__755__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testPrintMe) CljsCoreIHash__() {}

func (this__746__auto__ *CljsCore_testPrintMe) X_hash_Arity1() interface{} {
	{
		var h__579__auto__ = this__746__auto__.X__hash
		_ = h__579__auto__
		if !(cljs_core.Nil_(h__579__auto__)) {
			return h__579__auto__
		} else {
			{
				var h__579__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__746__auto__).(float64)
				_ = h__579__auto_____1
				this__746__auto__.X__hash = h__579__auto_____1

				return h__579__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIEquiv__() {}

func (this__747__auto__ *CljsCore_testPrintMe) X_equiv_Arity2(other__748__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__748__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__747__auto__), cljs_core.Type_.X_invoke_Arity1(other__748__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__747__auto__, other__748__auto__)))
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

func (this__762__auto__ *CljsCore_testPrintMe) X_dissoc_Arity2(k__763__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__763__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__762__auto__), this__762__auto__.X__meta), k__763__auto__)
	} else {
		return (&CljsCore_testPrintMe{this__762__auto__.A, this__762__auto__.B, this__762__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__762__auto__.X__extmap, k__763__auto__)), nil})
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIAssociative__() {}

func (this__758__auto__ *CljsCore_testPrintMe) X_assoc_Arity3(k__759__auto__ interface{}, G__746 interface{}) interface{} {
	{
		var pred__758 = cljs_core.Keyword_identical_QMARK_
		var expr__759 = k__759__auto__
		_, _ = pred__758, expr__759
		if cljs_core.Truth_(func() interface{} {
			var G__761 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__762 = expr__759
			_, _ = G__761, G__762
			return pred__758.X_invoke_Arity2(G__761, G__762)
		}()) {
			return (&CljsCore_testPrintMe{G__746, this__758__auto__.B, this__758__auto__.X__meta, this__758__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__763 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__764 = expr__759
				_, _ = G__763, G__764
				return pred__758.X_invoke_Arity2(G__763, G__764)
			}()) {
				return (&CljsCore_testPrintMe{this__758__auto__.A, G__746, this__758__auto__.X__meta, this__758__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPrintMe{this__758__auto__.A, this__758__auto__.B, this__758__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__758__auto__.X__extmap, k__759__auto__, G__746), nil})
			}
		}
	}
}

func (this__760__auto__ *CljsCore_testPrintMe) X_contains_key_QMARK__Arity2(k__761__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testPrintMe) CljsCoreISeqable__() {}

func (this__765__auto__ *CljsCore_testPrintMe) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__765__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__765__auto__.B}, nil})}, nil}), this__765__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testPrintMe) CljsCoreIWithMeta__() {}

func (this__750__auto__ *CljsCore_testPrintMe) X_with_meta_Arity2(G__746 interface{}) interface{} {
	return (&CljsCore_testPrintMe{this__750__auto__.A, this__750__auto__.B, G__746, this__750__auto__.X__extmap, this__750__auto__.X__hash})
}

func (_ *CljsCore_testPrintMe) CljsCoreICollection__() {}

func (this__756__auto__ *CljsCore_testPrintMe) X_conj_Arity2(entry__757__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__757__auto__) {
		return this__756__auto__.X_assoc_Arity3(entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), entry__757__auto__.(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__756__auto__, entry__757__auto__)
	}
}

var X__GT_PrintMe *cljs_core.AFn

var Map__GT_PrintMe *cljs_core.AFn

type CljsCore_testIBar interface {
	CljsCore_testIBar__()
	X_bar_Arity2(x interface{}) interface{}
}

func init() {
	cljs_core.RegisterProtocol_("cljs.core-test/IBar", (*CljsCore_testIBar)(nil))
}

var X_bar *cljs_core.AFn

var Baz *cljs_core.AFn

type CljsCore_testT770 struct {
	F          interface{}
	Baz        interface{}
	Test_stuff interface{}
	Meta771    interface{}
}

func (_ *CljsCore_testT770) CljsCore_testIBar__() {}

func (___ *CljsCore_testT770) X_bar_Arity2(x interface{}) interface{} {
	{
		var G__774 = x
		_ = G__774
		return ___.F.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__774)
	}
}

func (_ *CljsCore_testT770) CljsCoreIMeta__() {}

func (_772 *CljsCore_testT770) X_meta_Arity1() interface{} {
	return _772.Meta771
}

func (_ *CljsCore_testT770) CljsCoreIWithMeta__() {}

func (_772 *CljsCore_testT770) X_with_meta_Arity2(meta771___1 interface{}) interface{} {
	return (&CljsCore_testT770{_772.F, _772.Baz, _772.Test_stuff, meta771___1})
}

var X__GT_t770 *cljs_core.AFn

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

type CljsCore_testT783 struct {
	Test_stuff interface{}
	Meta784    interface{}
}

func (_ *CljsCore_testT783) CljsCoreIHash__() {}

func (___ *CljsCore_testT783) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT783) CljsCoreIMeta__() {}

func (_785 *CljsCore_testT783) X_meta_Arity1() interface{} {
	return _785.Meta784
}

func (_ *CljsCore_testT783) CljsCoreIWithMeta__() {}

func (_785 *CljsCore_testT783) X_with_meta_Arity2(meta784___1 interface{}) interface{} {
	return (&CljsCore_testT783{_785.Test_stuff, meta784___1})
}

var X__GT_t783 *cljs_core.AFn

type CljsCore_testT786 struct {
	A          interface{}
	Test_stuff interface{}
	Meta787    interface{}
}

func (_ *CljsCore_testT786) CljsCoreIHash__() {}

func (___ *CljsCore_testT786) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT786) CljsCoreIMeta__() {}

func (_788 *CljsCore_testT786) X_meta_Arity1() interface{} {
	return _788.Meta787
}

func (_ *CljsCore_testT786) CljsCoreIWithMeta__() {}

func (_788 *CljsCore_testT786) X_with_meta_Arity2(meta787___1 interface{}) interface{} {
	return (&CljsCore_testT786{_788.A, _788.Test_stuff, meta787___1})
}

var X__GT_t786 *cljs_core.AFn

var Some_x float64

var Some_y float64

type CljsCore_testIWoz interface {
	CljsCore_testIWoz__()
	X_woz_Arity1() interface{}
}

func init() {
	cljs_core.RegisterProtocol_("cljs.core-test/IWoz", (*CljsCore_testIWoz)(nil))
}

var X_woz *cljs_core.AFn

var Noz cljs_core.CljsCoreIVector

var Cljs_739 *cljs_core.AFn

var Cljs_780 *cljs_core.CljsCoreAtom

type CljsCore_testT856 struct {
	From_seq   interface{}
	Make_seq   interface{}
	Mt         interface{}
	I__848     interface{}
	Count__847 interface{}
	Chunk__846 interface{}
	Seq__845   interface{}
	Test_stuff interface{}
	Meta857    interface{}
}

func (_ *CljsCore_testT856) CljsCoreISeq__() {}

func (this *CljsCore_testT856) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT856) X_rest_Arity1() interface{} {
	{
		var G__860 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__860
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__860)
	}
}

func (_ *CljsCore_testT856) CljsCoreISeqable__() {}

func (this *CljsCore_testT856) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT856) CljsCoreIMeta__() {}

func (_858 *CljsCore_testT856) X_meta_Arity1() interface{} {
	return _858.Meta857
}

func (_ *CljsCore_testT856) CljsCoreIWithMeta__() {}

func (_858 *CljsCore_testT856) X_with_meta_Arity2(meta857___1 interface{}) interface{} {
	return (&CljsCore_testT856{_858.From_seq, _858.Make_seq, _858.Mt, _858.I__848, _858.Count__847, _858.Chunk__846, _858.Seq__845, _858.Test_stuff, meta857___1})
}

var X__GT_t856 *cljs_core.AFn

type CljsCore_testT868 struct {
	From_seq           interface{}
	Make_seq           interface{}
	Mt                 interface{}
	Temp__4222__auto__ interface{}
	I__848             interface{}
	Count__847         interface{}
	Chunk__846         interface{}
	Seq__845           interface{}
	Test_stuff         interface{}
	Meta869            interface{}
}

func (_ *CljsCore_testT868) CljsCoreISeq__() {}

func (this *CljsCore_testT868) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT868) X_rest_Arity1() interface{} {
	{
		var G__872 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__872
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__872)
	}
}

func (_ *CljsCore_testT868) CljsCoreISeqable__() {}

func (this *CljsCore_testT868) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT868) CljsCoreIMeta__() {}

func (_870 *CljsCore_testT868) X_meta_Arity1() interface{} {
	return _870.Meta869
}

func (_ *CljsCore_testT868) CljsCoreIWithMeta__() {}

func (_870 *CljsCore_testT868) X_with_meta_Arity2(meta869___1 interface{}) interface{} {
	return (&CljsCore_testT868{_870.From_seq, _870.Make_seq, _870.Mt, _870.Temp__4222__auto__, _870.I__848, _870.Count__847, _870.Chunk__846, _870.Seq__845, _870.Test_stuff, meta869___1})
}

var X__GT_t868 *cljs_core.AFn

var Case_recur *cljs_core.AFn

var Xform *cljs_core.AFn

var Data interface{}

var Xf *cljs_core.AFn

func Test_runner(t *testing.T) {
	Test_stuff.X_invoke_Arity0()
}
