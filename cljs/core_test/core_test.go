// Compiled by ClojureScript to Go 0.0-2371
// cljs.core-test

package core_test

import (
	"reflect"
	"testing"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	clojure_set "github.com/hraberg/cljs2go/clojure/set"
	goog_string "github.com/hraberg/cljs2go/goog/string"
	"github.com/hraberg/cljs2go/js"
	"github.com/hraberg/cljs2go/js/Math"
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
						i, j = (i + float64(1)), cljs_core.Conj.X_invoke_Arity2(j, func(G__947 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__947, 0, func() interface{} {
								return i
							})
						}(&cljs_core.AFn{}, i, j))
						continue
					} else {
						return cljs_core.Map_.X_invoke_Arity2(func(G__948 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__948, 1, func(p1__56_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})}, nil}), cljs_core.Map_.X_invoke_Arity2(func(G__949 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__949, 1, func(p1__57_SHARP_ interface{}) interface{} {
					{
						return p1__57_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
					}
				})
			}(&cljs_core.AFn{}), func() *cljs_core.CljsCoreLazySeq {
				var iter__940__auto__ = func(iter__521 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(iter__521, 1, func(s__522 interface{}) interface{} {
						return (&cljs_core.CljsCoreLazySeq{nil, func(G__950 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__950, 0, func() interface{} {
								{
									var s__522___1 interface{} = s__522
									_ = s__522___1
									for {
										{
											var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__522___1)
											_ = temp__4388__auto__
											if cljs_core.Truth_(temp__4388__auto__) {
												{
													var xs__4940__auto__ = temp__4388__auto__
													_ = xs__4940__auto__
													{
														var i = cljs_core.First.X_invoke_Arity1(xs__4940__auto__)
														_ = i
														{
															var iterys__936__auto__ = func(iter__523 *cljs_core.AFn, s__522___1 interface{}, i interface{}, xs__4940__auto__ cljs_core.CljsCoreISeq, temp__4388__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																return cljs_core.Fn(iter__523, 1, func(s__524 interface{}) interface{} {
																	return (&cljs_core.CljsCoreLazySeq{nil, func(G__951 *cljs_core.AFn, s__522___1 interface{}, i interface{}, xs__4940__auto__ cljs_core.CljsCoreISeq, temp__4388__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																		return cljs_core.Fn(G__951, 0, func() interface{} {
																			{
																				var s__524___1 interface{} = s__524
																				_ = s__524___1
																				for {
																					{
																						var temp__4388__auto_____1 = cljs_core.Seq.Arity1IQ(s__524___1)
																						_ = temp__4388__auto_____1
																						if cljs_core.Truth_(temp__4388__auto_____1) {
																							{
																								var s__524___2 = temp__4388__auto_____1
																								_ = s__524___2
																								if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__524___2) {
																									{
																										var c__938__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__524___2)
																										var size__939__auto__ = cljs_core.Count.X_invoke_Arity1(c__938__auto__).(float64)
																										var b__526 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__939__auto__).(*cljs_core.CljsCoreChunkBuffer)
																										_, _, _ = c__938__auto__, size__939__auto__, b__526
																										if func() bool {
																											var i__525 = float64(0)
																											_ = i__525
																											for {
																												if i__525 < size__939__auto__ {
																													{
																														var j = cljs_core.Decorate_(c__938__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__525)
																														_ = j
																														cljs_core.Chunk_append.X_invoke_Arity2(b__526, func(G__952 *cljs_core.AFn, i__525 float64, s__522___1 interface{}, j interface{}, c__938__auto__ interface{}, size__939__auto__ float64, b__526 *cljs_core.CljsCoreChunkBuffer, s__524___2 interface{}, temp__4388__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4940__auto__ cljs_core.CljsCoreISeq, temp__4388__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																															return cljs_core.Fn(G__952, 0, func() interface{} {
																																return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																															})
																														}(&cljs_core.AFn{}, i__525, s__522___1, j, c__938__auto__, size__939__auto__, b__526, s__524___2, temp__4388__auto_____1, i, xs__4940__auto__, temp__4388__auto__))
																														i__525 = (i__525 + float64(1))
																														continue
																													}
																												} else {
																													return true
																												}
																											}
																										}() {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__526), iter__523.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__524___2)).(*cljs_core.CljsCoreLazySeq))
																										} else {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__526), nil)
																										}
																									}
																								} else {
																									{
																										var j = cljs_core.First.X_invoke_Arity1(s__524___2)
																										_ = j
																										return cljs_core.Cons.X_invoke_Arity2(func(G__953 *cljs_core.AFn, s__522___1 interface{}, j interface{}, s__524___2 interface{}, temp__4388__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4940__auto__ cljs_core.CljsCoreISeq, temp__4388__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																											return cljs_core.Fn(G__953, 0, func() interface{} {
																												return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																											})
																										}(&cljs_core.AFn{}, s__522___1, j, s__524___2, temp__4388__auto_____1, i, xs__4940__auto__, temp__4388__auto__), iter__523.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__524___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
																	}(&cljs_core.AFn{}, s__522___1, i, xs__4940__auto__, temp__4388__auto__), nil, nil})
																})
															}(&cljs_core.AFn{}, s__522___1, i, xs__4940__auto__, temp__4388__auto__)
															var fs__937__auto__ = cljs_core.Seq.Arity1IQ(iterys__936__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq))
															_, _ = iterys__936__auto__, fs__937__auto__
															if cljs_core.Truth_(fs__937__auto__) {
																return cljs_core.Concat.X_invoke_Arity2(fs__937__auto__, iter__521.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__522___1)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
															} else {
																s__522___1 = cljs_core.Rest.Arity1IQ(s__522___1)
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
				_ = iter__940__auto__
				return iter__940__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})).(*cljs_core.CljsCoreLazySeq)
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3)), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil})) {
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
				var sb__1141__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1141__auto__
				{
					var _STAR_print_fn_STAR_532_954 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_532_954
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_532_954

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__955 *cljs_core.AFn, _STAR_print_fn_STAR_532_954 interface{}, sb__1141__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__955, 1, func(x__1142__auto__ interface{}) interface{} {
									return sb__1141__auto__.Append(x__1142__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_532_954, sb__1141__auto__)

							cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)}))
							cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)}))
						}
					}()
				}
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1141__auto__).(string))
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= \"12\" (with-out-str (print 1) (print 2)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("12", func() string {
				var sb__1141__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1141__auto__
				{
					var _STAR_print_fn_STAR_533_956 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_533_956
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_533_956

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__957 *cljs_core.AFn, _STAR_print_fn_STAR_533_956 interface{}, sb__1141__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__957, 1, func(x__1142__auto__ interface{}) interface{} {
									return sb__1141__auto__.Append(x__1142__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_533_956, sb__1141__auto__)

							cljs_core.X_STAR_print_fn_STAR_.X_invoke_Arity1(float64(1))
							cljs_core.X_STAR_print_fn_STAR_.X_invoke_Arity1(float64(2))
						}
					}()
				}
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1141__auto__).(string))
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
				var jumble = func(G__958 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__958, 2, func(a interface{}, b interface{}) interface{} {
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
				var jumble = func(G__959 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__959, 2, func(a interface{}, b interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB(float64(127), cljs_core.Apply.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(1), float64(2), float64(4), float64(8), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(16), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(64))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(32))}))) {
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
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Apply.X_invoke_Arity2(func(G__960 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__960, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return ((cljs_core.Nth.X_invoke_Arity2(args, float64(0)).(float64) + cljs_core.Nth.X_invoke_Arity2(args, float64(1)).(float64)) + cljs_core.Nth.X_invoke_Arity2(args, float64(2)).(float64))
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))) {
			} else {
				panic((&js.Error{("Assert failed: (= 3 (apply (fn [& args] (+ (nth args 0) (nth args 1) (nth args 2))) (iterate inc 0)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__961 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__961, 0, func(m__ ...interface{}) interface{} {
					var m = cljs_core.Seq.Arity1IQ(m__[0])
					_ = m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [0 1 2 3 4] (take 5 (apply (fn [& m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__962 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__962, 1, func(x_m__ ...interface{}) interface{} {
					var x = x_m__[0]
					var m = cljs_core.Seq.Arity1IQ(x_m__[1])
					_, _ = x, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (take 5 (apply (fn [x & m] m) (iterate inc 0))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__963 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__963, 2, func(x_y_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5), float64(6), float64(7)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__964 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__964, 3, func(x_y_z_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6), float64(7), float64(8)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__965 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__965, 4, func(x_y_z_a_m__ ...interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__966 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__966, 5, func(x_y_z_a_b_m__ ...interface{}) interface{} {
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
				var single_arity_non_variadic_967 = func(G__972 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__972, 3, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{})
				var multiple_arity_non_variadic_968 = func(G__973 *cljs_core.AFn, single_arity_non_variadic_967 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__973, 3, func(x interface{}) interface{} {
						return x
					}, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{y, x}, nil})
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_967)
				var single_arity_variadic_fixedargs_969 = func(G__974 *cljs_core.AFn, single_arity_non_variadic_967 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_968 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__974, 2, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{more, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_967, multiple_arity_non_variadic_968)
				var single_arity_variadic_nofixedargs_970 = func(G__975 *cljs_core.AFn, single_arity_non_variadic_967 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_968 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_969 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__975, 0, func(more__ ...interface{}) interface{} {
						var more = cljs_core.Seq.Arity1IQ(more__[0])
						_ = more
						return more
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_967, multiple_arity_non_variadic_968, single_arity_variadic_fixedargs_969)
				var multiple_arity_variadic_971 = func(G__976 *cljs_core.AFn, single_arity_non_variadic_967 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_968 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_969 cljs_core.CljsCoreIFn, single_arity_variadic_nofixedargs_970 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__976, 2, func(x interface{}) interface{} {
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
				}(&cljs_core.AFn{}, single_arity_non_variadic_967, multiple_arity_non_variadic_968, single_arity_variadic_fixedargs_969, single_arity_variadic_nofixedargs_970)
				_, _, _, _, _ = single_arity_non_variadic_967, multiple_arity_non_variadic_968, single_arity_variadic_fixedargs_969, single_arity_variadic_nofixedargs_970, multiple_arity_variadic_971
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_non_variadic_967, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic [1 2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_non_variadic_967, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic 1 [2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_non_variadic_967, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply single-arity-non-variadic 1 2 [3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_968, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (apply multiple-arity-non-variadic [42]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_968, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-non-variadic [1 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_968, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-non-variadic 1 [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_968, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic [1 2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_968, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic 1 [2 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_non_variadic_968, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 2 1] (apply multiple-arity-non-variadic 1 2 [3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_969, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_fixedargs_969, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_fixedargs_969, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_fixedargs_969, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_fixedargs_969, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_969, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 4 5] (take 3 (first (apply single-arity-variadic-fixedargs (iterate inc 1)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_969, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (rest (apply single-arity-variadic-fixedargs (iterate inc 1))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_970, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_nofixedargs_970, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_nofixedargs_970, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_nofixedargs_970, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_nofixedargs_970, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_970, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2 3 4 5] (take 5 (apply single-arity-variadic-nofixedargs (iterate inc 1))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_971, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (apply multiple-arity-variadic [42]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_971, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-variadic [1 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_971, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (apply multiple-arity-variadic 1 [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_971, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic [1 2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_971, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 [2 3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_variadic_971, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 [3 4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(multiple_arity_variadic_971, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 [4 5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(multiple_arity_variadic_971, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 4 [5]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_971, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= [3 4 5] (take 3 (first (apply multiple-arity-variadic (iterate inc 1)))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_971, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 1] (rest (apply multiple-arity-variadic (iterate inc 1))))")}))
				}
			}
			{
				var f1_977 = func(f1 *cljs_core.AFn) *cljs_core.AFn {
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
				var f2_978 = func(f2 *cljs_core.AFn, f1_977 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(f2, 2, func(x interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					}, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return cljs_core.Apply.X_invoke_Arity3(f1_977, y, more)
					})
				}(&cljs_core.AFn{}, f1_977)
				_, _ = f1_977, f2_978
				if cljs_core.X_EQ_.Arity2IIB(float64(1), f2_978.X_invoke_Arity2(float64(1), float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (f2 1 2))")}))
				}
			}
			{
				var f_979 = func(G__980 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__980, 1, func() interface{} {
						return nil
					}, func(a_more__ ...interface{}) interface{} {
						var a = a_more__[0]
						var more = cljs_core.Seq.Arity1IQ(a_more__[1])
						_, _ = a, more
						return more
					})
				}(&cljs_core.AFn{})
				_ = f_979
				if cljs_core.Nil_(f_979.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (f :foo))")}))
				}
			}
			if cljs_core.Nil_(cljs_core.Array_seq.X_invoke_Arity2([]interface{}{float64(1)}, float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (array-seq (array 1) 1))")}))
			}
			{
				var f_981 = func(G__984 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__984, 1, func(x interface{}) interface{} {
						return (x.(float64) * float64(2))
					})
				}(&cljs_core.AFn{})
				var m_982 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar"}, nil})
				var mf_983 = cljs_core.With_meta.X_invoke_Arity2(f_981, m_982)
				_, _, _ = f_981, m_982, mf_983
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(f_981)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (meta f))")}))
				}
				if cljs_core.Fn_QMARK_.Arity1IB(mf_983) {
				} else {
					panic((&js.Error{("Assert failed: (fn? mf)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__534 = float64(2)
					_ = G__534
					return mf_983.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__534)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (mf 2))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Apply.X_invoke_Arity2(mf_983, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (apply mf [2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(mf_983), m_982) {
				} else {
					panic((&js.Error{("Assert failed: (= (meta mf) m)")}))
				}
			}
			{
				var a_985 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_985
				if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Deref.X_invoke_Arity1(a_985)) {
				} else {
					panic((&js.Error{("Assert failed: (= 0 (deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity2(a_985, cljs_core.Inc)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (swap! a inc))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(false, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_985, float64(0), float64(42)))) {
				} else {
					panic((&js.Error{("Assert failed: (= false (compare-and-set! a 0 42))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(true, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_985, float64(1), float64(7)))) {
				} else {
					panic((&js.Error{("Assert failed: (= true (compare-and-set! a 1 7))")}))
				}
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(a_985)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (meta a))")}))
				}
				if cljs_core.Nil_(cljs_core.Get_validator.X_invoke_Arity1(a_985)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (get-validator a))")}))
				}
			}
			{
				var a_986 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_986
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity3(a_986, cljs_core.X_PLUS_, float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (swap! a + 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Swap_BANG_.X_invoke_Arity4(a_986, cljs_core.X_PLUS_, float64(1), float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (swap! a + 1 2))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(10), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_986, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 10 (swap! a + 1 2 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(20), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_986, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(4)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= 20 (swap! a + 1 2 3 4))")}))
				}
			}
			{
				var a_987 = cljs_core.Atom.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "validator", Fqn: "validator", X_hash: float64(-1966190681)}), cljs_core.Coll_QMARK_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})})).(*cljs_core.CljsCoreAtom)
				_ = a_987
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Coll_QMARK_, cljs_core.Get_validator.X_invoke_Arity1(a_987)) {
				} else {
					panic((&js.Error{("Assert failed: (= coll? (get-validator a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_987)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:a 1} (meta a))")}))
				}
				cljs_core.Alter_meta_BANG_.X_invoke_ArityVariadic(a_987, cljs_core.Assoc, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}))
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_987)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:a 1, :b 2} (meta a))")}))
				}
			}
			if cljs_core.Nil_(cljs_core.Empty.X_invoke_Arity1(nil)) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (empty nil))")}))
			}
			{
				var e_lazy_seq_988 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCoreLazySeq{nil, func(G__989 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__989, 0, func() interface{} {
						return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons)
					})
				}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_lazy_seq_988
				if cljs_core.Seq_QMARK_.Arity1IB(e_lazy_seq_988) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-lazy-seq)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_lazy_seq_988) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-lazy-seq)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_lazy_seq_988)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-lazy-seq))")}))
				}
			}
			{
				var e_list_990 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_list_990
				if cljs_core.Seq_QMARK_.Arity1IB(e_list_990) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-list)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_list_990) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-list)")}))
				}
			}
			{
				var e_elist_991 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_elist_991
				if cljs_core.Seq_QMARK_.Arity1IB(e_elist_991) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-elist)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_elist_991) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-elist)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), cljs_core.Get.X_invoke_Arity2(cljs_core.Meta.X_invoke_Arity1(e_elist_991), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= :c (get (meta e-elist) :b))")}))
				}
			}
			{
				var e_cons_992 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_cons_992
				if cljs_core.Seq_QMARK_.Arity1IB(e_cons_992) {
				} else {
					panic((&js.Error{("Assert failed: (seq? e-cons)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_cons_992) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-cons)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_cons_992)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-cons))")}))
				}
			}
			{
				var e_vec_993 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_vec_993
				if cljs_core.Vector_QMARK_.Arity1IB(e_vec_993) {
				} else {
					panic((&js.Error{("Assert failed: (vector? e-vec)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_vec_993) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-vec)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_vec_993)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-vec))")}))
				}
			}
			{
				var e_omap_994 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_omap_994
				if cljs_core.Map_QMARK_.Arity1IB(e_omap_994) {
				} else {
					panic((&js.Error{("Assert failed: (map? e-omap)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_omap_994) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-omap)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_omap_994)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-omap))")}))
				}
			}
			{
				var e_hmap_995 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_hmap_995
				if cljs_core.Map_QMARK_.Arity1IB(e_hmap_995) {
				} else {
					panic((&js.Error{("Assert failed: (map? e-hmap)")}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_hmap_995) {
				} else {
					panic((&js.Error{("Assert failed: (empty? e-hmap)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_hmap_995)) {
				} else {
					panic((&js.Error{("Assert failed: (= {:b :c} (meta e-hmap))")}))
				}
			}
			{
				var a_996 = cljs_core.Atom.X_invoke_Arity1(nil).(*cljs_core.CljsCoreAtom)
				_ = a_996
				if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (try 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__997 interface{}) {
					defer func() {
						if e535 := recover(); e535 != nil {
							if cljs_core.Value_(e535).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e535
									_ = e
									return__997 = float64(2)
								}
							} else {
								panic(e535)

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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__998 interface{}) {
					defer func() {
						if e536 := recover(); e536 != nil {
							if cljs_core.Value_(e536).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e536
									_ = e
									return__998 = float64(2)
								}
							} else {
								panic(e536)

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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__999 interface{}) {
					defer func() {
						if e537 := recover(); e537 != nil {
							if cljs_core.Value_(e537).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e537
									_ = e
									return__999 = float64(2)
								}
							} else {
								{
									var e = e537
									_ = e
									return__999 = float64(3)
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
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__1000 interface{}) {
					defer func() {
						if e538 := recover(); e538 != nil {
							if cljs_core.Value_(e538).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e538
									_ = e
									return__1000 = float64(2)
								}
							} else {
								{
									var e = e538
									_ = e
									return__1000 = float64(3)
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
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1001 interface{}) {
					defer func() {
						if e539 := recover(); e539 != nil {
							if cljs_core.Value_(e539).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e539
									_ = e
									return__1001 = float64(3)
								}
							} else {
								{
									var e = e539
									_ = e
									return__1001 = e
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
						cljs_core.Reset_BANG_.X_invoke_Arity2(a_996, float64(42))
					}()
					{
						return float64(1)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (try 1 (finally (reset! a 42))))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Deref.X_invoke_Arity1(a_996)) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (deref a))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), cljs_core.Seq_(cljs_core.Nthnext.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)))) {
			} else {
				panic((&js.Error{("Assert failed: (= [3] (nthnext [1 2 3] 2))")}))
			}
			{
				var v_1002 = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})
				_ = v_1002
				if cljs_core.X_EQ_.Arity2IIB(v_1002, func() *cljs_core.CljsCoreLazySeq {
					var iter__940__auto__ = func(iter__540 *cljs_core.AFn, v_1002 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__540, 1, func(s__541 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1003 *cljs_core.AFn, v_1002 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1003, 0, func() interface{} {
									{
										var s__541___1 interface{} = s__541
										_ = s__541___1
										for {
											{
												var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__541___1)
												_ = temp__4388__auto__
												if cljs_core.Truth_(temp__4388__auto__) {
													{
														var s__541___2 = temp__4388__auto__
														_ = s__541___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__541___2) {
															{
																var c__938__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__541___2)
																var size__939__auto__ = cljs_core.Count.X_invoke_Arity1(c__938__auto__).(float64)
																var b__543 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__939__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__938__auto__, size__939__auto__, b__543
																if func() bool {
																	var i__542 = float64(0)
																	_ = i__542
																	for {
																		if i__542 < size__939__auto__ {
																			{
																				var e = cljs_core.Decorate_(c__938__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__542)
																				_ = e
																				cljs_core.Chunk_append.X_invoke_Arity2(b__543, e)
																				i__542 = (i__542 + float64(1))
																				continue
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__543), iter__540.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__541___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__543), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__541___2)
																_ = e
																return cljs_core.Cons.X_invoke_Arity2(e, iter__540.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__541___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_1002), nil, nil})
						})
					}(&cljs_core.AFn{}, v_1002)
					_ = iter__940__auto__
					return iter__940__auto__.X_invoke_Arity1(v_1002).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= v (for [e v] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(9)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__940__auto__ = func(iter__546 *cljs_core.AFn, v_1002 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__546, 1, func(s__547 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1004 *cljs_core.AFn, v_1002 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1004, 0, func() interface{} {
									{
										var s__547___1 interface{} = s__547
										_ = s__547___1
										for {
											{
												var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__547___1)
												_ = temp__4388__auto__
												if cljs_core.Truth_(temp__4388__auto__) {
													{
														var s__547___2 = temp__4388__auto__
														_ = s__547___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__547___2) {
															{
																var c__938__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__547___2)
																var size__939__auto__ = cljs_core.Count.X_invoke_Arity1(c__938__auto__).(float64)
																var b__549 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__939__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__938__auto__, size__939__auto__, b__549
																if func() bool {
																	var i__548 = float64(0)
																	_ = i__548
																	for {
																		if i__548 < size__939__auto__ {
																			{
																				var e = cljs_core.Decorate_(c__938__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__548)
																				_ = e
																				{
																					var m = (e.(float64) * e.(float64))
																					_ = m
																					cljs_core.Chunk_append.X_invoke_Arity2(b__549, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																					i__548 = (i__548 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__549), iter__546.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__547___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__549), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__547___2)
																_ = e
																{
																	var m = (e.(float64) * e.(float64))
																	_ = m
																	return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__546.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__547___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_1002), nil, nil})
						})
					}(&cljs_core.AFn{}, v_1002)
					_ = iter__940__auto__
					return iter__940__auto__.X_invoke_Arity1(v_1002).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [[1 1] [2 4] [3 9]] (for [e v :let [m (* e e)]] [e m]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__940__auto__ = func(iter__552 *cljs_core.AFn, v_1002 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__552, 1, func(s__553 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1005 *cljs_core.AFn, v_1002 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1005, 0, func() interface{} {
									{
										var s__553___1 interface{} = s__553
										_ = s__553___1
										for {
											{
												var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__553___1)
												_ = temp__4388__auto__
												if cljs_core.Truth_(temp__4388__auto__) {
													{
														var s__553___2 = temp__4388__auto__
														_ = s__553___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__553___2) {
															{
																var c__938__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__553___2)
																var size__939__auto__ = cljs_core.Count.X_invoke_Arity1(c__938__auto__).(float64)
																var b__555 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__939__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__938__auto__, size__939__auto__, b__555
																if cljs_core.Truth_(func() interface{} {
																	var i__554 = float64(0)
																	_ = i__554
																	for {
																		if i__554 < size__939__auto__ {
																			{
																				var e = cljs_core.Decorate_(c__938__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__554)
																				_ = e
																				if e.(float64) < float64(3) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__555, e)
																					i__554 = (i__554 + float64(1))
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
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__555), iter__552.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__553___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__555), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__553___2)
																_ = e
																if e.(float64) < float64(3) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__552.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__553___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_1002), nil, nil})
						})
					}(&cljs_core.AFn{}, v_1002)
					_ = iter__940__auto__
					return iter__940__auto__.X_invoke_Arity1(v_1002).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [1 2] (for [e v :while (< e 3)] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__940__auto__ = func(iter__558 *cljs_core.AFn, v_1002 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__558, 1, func(s__559 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1006 *cljs_core.AFn, v_1002 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1006, 0, func() interface{} {
									{
										var s__559___1 interface{} = s__559
										_ = s__559___1
										for {
											{
												var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__559___1)
												_ = temp__4388__auto__
												if cljs_core.Truth_(temp__4388__auto__) {
													{
														var s__559___2 = temp__4388__auto__
														_ = s__559___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__559___2) {
															{
																var c__938__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__559___2)
																var size__939__auto__ = cljs_core.Count.X_invoke_Arity1(c__938__auto__).(float64)
																var b__561 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__939__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__938__auto__, size__939__auto__, b__561
																if func() bool {
																	var i__560 = float64(0)
																	_ = i__560
																	for {
																		if i__560 < size__939__auto__ {
																			{
																				var e = cljs_core.Decorate_(c__938__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__560)
																				_ = e
																				if e.(float64) > float64(2) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__561, e)
																					i__560 = (i__560 + float64(1))
																					continue
																				} else {
																					i__560 = (i__560 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__561), iter__558.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__559___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__561), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__559___2)
																_ = e
																if e.(float64) > float64(2) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__558.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__559___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
																} else {
																	s__559___1 = cljs_core.Rest.Arity1IQ(s__559___2)
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
							}(&cljs_core.AFn{}, v_1002), nil, nil})
						})
					}(&cljs_core.AFn{}, v_1002)
					_ = iter__940__auto__
					return iter__940__auto__.X_invoke_Arity1(v_1002).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= [3] (for [e v :when (> e 2)] e))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__940__auto__ = func(iter__564 *cljs_core.AFn, v_1002 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__564, 1, func(s__565 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1007 *cljs_core.AFn, v_1002 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1007, 0, func() interface{} {
									{
										var s__565___1 interface{} = s__565
										_ = s__565___1
										for {
											{
												var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__565___1)
												_ = temp__4388__auto__
												if cljs_core.Truth_(temp__4388__auto__) {
													{
														var s__565___2 = temp__4388__auto__
														_ = s__565___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__565___2) {
															{
																var c__938__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__565___2)
																var size__939__auto__ = cljs_core.Count.X_invoke_Arity1(c__938__auto__).(float64)
																var b__567 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__939__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__938__auto__, size__939__auto__, b__567
																if cljs_core.Truth_(func() interface{} {
																	var i__566 = float64(0)
																	_ = i__566
																	for {
																		if i__566 < size__939__auto__ {
																			{
																				var e = cljs_core.Decorate_(c__938__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__566)
																				_ = e
																				if e.(float64) < float64(3) {
																					{
																						var m = (e.(float64) * e.(float64))
																						_ = m
																						cljs_core.Chunk_append.X_invoke_Arity2(b__567, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																						i__566 = (i__566 + float64(1))
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
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__567), iter__564.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__565___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__567), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__565___2)
																_ = e
																if e.(float64) < float64(3) {
																	{
																		var m = (e.(float64) * e.(float64))
																		_ = m
																		return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__564.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__565___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_1002), nil, nil})
						})
					}(&cljs_core.AFn{}, v_1002)
					_ = iter__940__auto__
					return iter__940__auto__.X_invoke_Arity1(v_1002).(*cljs_core.CljsCoreLazySeq)
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{false, false, false}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true}, nil})}, nil}), cljs_core.Partition_by.X_invoke_Arity2(cljs_core.True_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true, false, false, false, true, true}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [[true true] [false false false] [true true]] (partition-by true? [true true false false false true true]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2), float64(4), float64(6), float64(8), float64(10)}, nil}), cljs_core.Take_nth.X_invoke_Arity2(float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(11), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{("Assert failed: (= [0 2 4 6 8 10] (take-nth 2 [0 1 2 3 4 5 6 7 8 9 10]))")}))
			}
			{
				var a10_1008 = cljs_core.Partial.X_invoke_Arity2(cljs_core.X_PLUS_, float64(10)).(cljs_core.CljsCoreIFn)
				var a20_1009 = cljs_core.Partial.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), float64(10)).(cljs_core.CljsCoreIFn)
				var a21_1010 = cljs_core.Partial.X_invoke_Arity4(cljs_core.X_PLUS_, float64(10), float64(10), float64(1)).(cljs_core.CljsCoreIFn)
				var a22_1011 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)})).(cljs_core.CljsCoreIFn)
				var a23_1012 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(1)})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = a10_1008, a20_1009, a21_1010, a22_1011, a23_1012
				if cljs_core.X_EQ_.Arity2IIB(float64(110), func() interface{} {
					var G__570 = float64(100)
					_ = G__570
					return a10_1008.X_invoke_Arity1(G__570)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 110 (a10 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(120), func() interface{} {
					var G__571 = float64(100)
					_ = G__571
					return a20_1009.X_invoke_Arity1(G__571)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 120 (a20 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(121), func() interface{} {
					var G__572 = float64(100)
					_ = G__572
					return a21_1010.X_invoke_Arity1(G__572)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 121 (a21 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(122), func() interface{} {
					var G__573 = float64(100)
					_ = G__573
					return a22_1011.X_invoke_Arity1(G__573)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 122 (a22 100))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(123), func() interface{} {
					var G__574 = float64(100)
					_ = G__574
					return a23_1012.X_invoke_Arity1(G__574)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 123 (a23 100))")}))
				}
			}
			{
				var n2_1013 = cljs_core.Comp.X_invoke_Arity2(cljs_core.First, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n3_1014 = cljs_core.Comp.X_invoke_Arity3(cljs_core.First, cljs_core.Rest, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n4_1015 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n5_1016 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n6_1017 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = n2_1013, n3_1014, n4_1015, n5_1016, n6_1017
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() interface{} {
					var G__575 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__575
					return n2_1013.X_invoke_Arity1(G__575)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (n2 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() interface{} {
					var G__576 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__576
					return n3_1014.X_invoke_Arity1(G__576)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (n3 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__577 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__577
					return n4_1015.X_invoke_Arity1(G__577)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 4 (n4 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(5), func() interface{} {
					var G__578 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__578
					return n5_1016.X_invoke_Arity1(G__578)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 5 (n5 [1 2 3 4 5 6 7]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), func() interface{} {
					var G__579 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__579
					return n6_1017.X_invoke_Arity1(G__579)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (n6 [1 2 3 4 5 6 7]))")}))
				}
			}
			{
				var sf_1018 = cljs_core.Some_fn.X_invoke_Arity3(cljs_core.Number_QMARK_, cljs_core.Keyword_QMARK_, cljs_core.Symbol_QMARK_).(cljs_core.CljsCoreIFn)
				_ = sf_1018
				if cljs_core.Truth_(func() interface{} {
					var G__580 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					var G__581 = float64(1)
					_, _ = G__580, G__581
					return sf_1018.X_invoke_Arity2(G__580, G__581)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf :foo 1)")}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__582 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					_ = G__582
					return sf_1018.X_invoke_Arity1(G__582)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf :foo)")}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__583 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})
					var G__584 = float64(1)
					_, _ = G__583, G__584
					return sf_1018.X_invoke_Arity2(G__583, G__584)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (sf (quote bar) 1)")}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__585 = cljs_core.CljsCorePersistentVector_EMPTY
					var G__586 = cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)
					_, _ = G__585, G__586
					return sf_1018.X_invoke_Arity2(G__585, G__586)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (not (sf [] ()))")}))
				}
			}
			{
				var ep_1019 = cljs_core.Every_pred.X_invoke_Arity2(cljs_core.Number_QMARK_, cljs_core.Zero_QMARK_).(cljs_core.CljsCoreIFn)
				_ = ep_1019
				if cljs_core.Truth_(func() interface{} {
					var G__587 = float64(0)
					var G__588 = float64(0)
					var G__589 = float64(0)
					_, _, _ = G__587, G__588, G__589
					return ep_1019.X_invoke_Arity3(G__587, G__588, G__589)
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (ep 0 0 0)")}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__590 = float64(1)
					var G__591 = float64(2)
					var G__592 = float64(3)
					var G__593 = float64(0)
					_, _, _, _ = G__590, G__591, G__592, G__593
					return ep_1019.X_invoke_Arity4(G__590, G__591, G__592, G__593)
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
			{
				var r1_1020 = (&js.RegExp{Pattern: `foo`, Flags: ``})
				var r2_1021 = cljs_core.Re_pattern.X_invoke_Arity1(r1_1020)
				_, _ = r1_1020, r2_1021
				if cljs_core.X_EQ_.Arity2IIB(r1_1020, r2_1021) {
				} else {
					panic((&js.Error{("Assert failed: (= r1 r2)")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_find.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo"), "foo bar foo baz foo zot"), "foo") {
			} else {
				panic((&js.Error{("Assert failed: (= (re-find (re-pattern \"foo\") \"foo bar foo baz foo zot\") \"foo\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_find.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("f(.)o"), "foo bar foo baz foo zot"), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (re-find (re-pattern \"f(.)o\") \"foo bar foo baz foo zot\") [\"foo\" \"o\"])")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo"), "foo"), "foo") {
			} else {
				panic((&js.Error{("Assert failed: (= (re-matches (re-pattern \"foo\") \"foo\") \"foo\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo"), "foo bar foo baz foo zot"), nil) {
			} else {
				panic((&js.Error{("Assert failed: (= (re-matches (re-pattern \"foo\") \"foo bar foo baz foo zot\") nil)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo.*"), "foo bar foo baz foo zot"), "foo bar foo baz foo zot") {
			} else {
				panic((&js.Error{("Assert failed: (= (re-matches (re-pattern \"foo.*\") \"foo bar foo baz foo zot\") \"foo bar foo baz foo zot\")")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_seq.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo"), "foo bar foo baz foo zot"), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2("foo")).(cljs_core.CljsCoreICollection).X_conj_Arity2("foo")).(cljs_core.CljsCoreICollection).X_conj_Arity2("foo")) {
			} else {
				panic((&js.Error{("Assert failed: (= (re-seq (re-pattern \"foo\") \"foo bar foo baz foo zot\") (list \"foo\" \"foo\" \"foo\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_seq.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("f(.)o"), "foo bar foo baz foo zot"), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= (re-seq (re-pattern \"f(.)o\") \"foo bar foo baz foo zot\") (list [\"foo\" \"o\"] [\"foo\" \"o\"] [\"foo\" \"o\"]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("(?i)foo"), "Foo"), "Foo") {
			} else {
				panic((&js.Error{("Assert failed: (= (re-matches (re-pattern \"(?i)foo\") \"Foo\") \"Foo\")")}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{"#\"(?:)\"", nil, "#\"\"", nil}, nil}, nil}).X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.RegExp{Pattern: ``, Flags: ``})})).(string))) {
			} else {
				panic((&js.Error{("Assert failed: (#{\"#\\\"(?:)\\\"\" \"#\\\"\\\"\"} (pr-str #\"\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__594 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__594, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__594, float64(1), nil)
				_, _, _ = vec__594, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] [1 2]] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), func() cljs_core.CljsCoreISet {
				var vec__595 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__595, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__595, float64(1), nil)
				_, _, _ = vec__595, a, b
				return cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{a, b}, true).(*cljs_core.CljsCorePersistentHashSet)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= #{1 2} (let [[a b] [1 2]] #{a b}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__596 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__596___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__596) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__596)
					} else {
						return map__596
					}
				}()
				var a = cljs_core.Get.X_invoke_Arity2(map__596___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				var b = cljs_core.Get.X_invoke_Arity2(map__596___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				_, _, _, _ = map__596, map__596___1, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (let [{a :a, b :b} {:a 1, :b 2}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__597 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__597___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__597) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__597)
					} else {
						return map__597
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__597___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__597___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__597, map__597___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2] (let [{:keys [a b]} {:a 1, :b 2}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__598 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__598, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__598, float64(1), nil)
				var v = vec__598
				_, _, _, _ = vec__598, a, b, v
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b, v}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 2 [1 2]] (let [[a b :as v] [1 2]] [a b v]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(42)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__599 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__599___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__599) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__599)
					} else {
						return map__599
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity3(map__599___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(42))
				var a = cljs_core.Get.X_invoke_Arity2(map__599___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__599, map__599___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 42] (let [{:keys [a b], :or {b 42}} {:a 1}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), nil}, nil}), func() cljs_core.CljsCoreIVector {
				var map__600 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__600___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__600) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__600)
					} else {
						return map__600
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__600___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__600___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__600, map__600___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [1 nil] (let [{:keys [a b], :or {c 42}} {:a 1}] [a b]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__601 = cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)
				var a = cljs_core.Nth.X_invoke_Arity3(vec__601, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__601, float64(1), nil)
				_, _, _ = vec__601, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= [2 1] (let [[a b] (quote (1 2))] [b a]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIMap {
				var vec__602 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__602, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__602, float64(1), nil)
				_, _, _ = vec__602, a, b
				return cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{a, b}, true, false).(*cljs_core.CljsCorePersistentArrayMap)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= {1 2} (let [[a b] [1 2]] {a b}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__603 = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))
				var a = cljs_core.Nth.X_invoke_Arity3(vec__603, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__603, float64(1), nil)
				_, _, _ = vec__603, a, b
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
				var a_1022 = cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{})
				_ = a_1022
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20), float64(30)}, nil}), cljs_core.Seq.Arity1IQ(func() []interface{} {
					var a__1059__auto__ = a_1022
					var ret = cljs_core.Aclone.X_invoke_Arity1(a__1059__auto__).([]interface{})
					_, _ = a__1059__auto__, ret
					{
						var i = float64(0)
						_ = i
						for {
							if i < float64(len(a__1059__auto__)) {
								ret[int(i)] = (float64(10) * (a_1022[int(i)]).(float64))
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
					var a__1065__auto__ = a_1022
					_ = a__1065__auto__
					{
						var i = float64(0)
						var ret = float64(0)
						_, _ = i, ret
						for {
							if i < float64(len(a__1065__auto__)) {
								i, ret = (i + float64(1)), (ret + (a_1022[int(i)]).(float64))
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_1022), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (seq a) (seq (to-array [1 2 3])))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), func() float64 { a_1022[int(float64(0))] = float64(42); return a_1022[int(float64(0))].(float64) }()) {
				} else {
					panic((&js.Error{("Assert failed: (= 42 (aset a 0 42))")}))
				}
				if cljs_core.Not_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_1022), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
				} else {
					panic((&js.Error{("Assert failed: (not= (seq a) (seq (to-array [1 2 3])))")}))
				}
			}
			{
				var a_1023 = []interface{}{[]interface{}{float64(1), float64(2), float64(3)}, []interface{}{float64(4), float64(5), float64(6)}}
				_ = a_1023
				if cljs_core.X_EQ_.Arity2IIB((a_1023[int(float64(0))].([]interface{})[int(float64(1))]), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (aget a 0 1) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity3(cljs_core.Aget, a_1023, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil})), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (apply aget a [0 1]) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((a_1023[int(float64(1))].([]interface{})[int(float64(1))]), float64(5)) {
				} else {
					panic((&js.Error{("Assert failed: (= (aget a 1 1) 5)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity3(cljs_core.Aget, a_1023, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), float64(5)) {
				} else {
					panic((&js.Error{("Assert failed: (= (apply aget a [1 1]) 5)")}))
				}
				a_1023[int(float64(0))].([]interface{})[int(float64(0))] = "foo"
				if cljs_core.X_EQ_.Arity2IIB((a_1023[int(float64(0))].([]interface{})[int(float64(0))]), "foo") {
				} else {
					panic((&js.Error{("Assert failed: (= (aget a 0 0) \"foo\")")}))
				}
				cljs_core.Apply.X_invoke_Arity3(cljs_core.Aset, a_1023, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0), "bar"}, nil}))
				if cljs_core.X_EQ_.Arity2IIB((a_1023[int(float64(0))].([]interface{})[int(float64(0))]), "bar") {
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
				var coll_1024 = (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})
				var shuffles_1025 = cljs_core.Filter.X_invoke_Arity2(func(G__1026 *cljs_core.AFn, coll_1024 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1026, 1, func(p1__58_SHARP_ interface{}) interface{} {
						return cljs_core.Not_EQ_.Arity2IIB(coll_1024, p1__58_SHARP_)
					})
				}(&cljs_core.AFn{}, coll_1024), cljs_core.Take.X_invoke_Arity2(float64(100), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Shuffle, coll_1024).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
				_, _ = coll_1024, shuffles_1025
				if !(cljs_core.Empty_QMARK_.Arity1IB(shuffles_1025)) {
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
				var s_1027 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = s_1027
				{
					var n__1071__auto___1028 = float64(5)
					_ = n__1071__auto___1028
					{
						var n_1029 = float64(0)
						_ = n_1029
						for {
							if n_1029 < n__1071__auto___1028 {
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_1027, cljs_core.Conj, n_1029)
								n_1029 = (n_1029 + float64(1))
								continue
							} else {
							}
							break
						}
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Deref.X_invoke_Arity1(s_1027)) {
				} else {
					panic((&js.Error{("Assert failed: (= [0 1 2 3 4] (clojure.core/deref s))")}))
				}
			}
			{
				var v_1030 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})
				var s_1031 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)).(*cljs_core.CljsCoreAtom)
				_, _ = v_1030, s_1031
				{
					var seq__604_1032 interface{} = cljs_core.Seq.Arity1IQ(v_1030)
					var chunk__605_1033 interface{} = nil
					var count__606_1034 = float64(0)
					var i__607_1035 = float64(0)
					_, _, _, _ = seq__604_1032, chunk__605_1033, count__606_1034, i__607_1035
					for {
						if i__607_1035 < count__606_1034 {
							{
								var n_1036 = cljs_core.Decorate_(chunk__605_1033).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__607_1035)
								_ = n_1036
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_1031, cljs_core.Conj, n_1036)
								seq__604_1032, chunk__605_1033, count__606_1034, i__607_1035 = seq__604_1032, chunk__605_1033, count__606_1034, (i__607_1035 + float64(1))
								continue
							}
						} else {
							{
								var temp__4388__auto___1037 = cljs_core.Seq.Arity1IQ(seq__604_1032)
								_ = temp__4388__auto___1037
								if cljs_core.Truth_(temp__4388__auto___1037) {
									{
										var seq__604_1038___1 = temp__4388__auto___1037
										_ = seq__604_1038___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__604_1038___1) {
											{
												var c__971__auto___1039 = cljs_core.Chunk_first.X_invoke_Arity1(seq__604_1038___1)
												_ = c__971__auto___1039
												seq__604_1032, chunk__605_1033, count__606_1034, i__607_1035 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__604_1038___1), c__971__auto___1039, cljs_core.Count.X_invoke_Arity1(c__971__auto___1039).(float64), float64(0)
												continue
											}
										} else {
											{
												var n_1040 = cljs_core.First.X_invoke_Arity1(seq__604_1038___1)
												_ = n_1040
												cljs_core.Swap_BANG_.X_invoke_Arity3(s_1031, cljs_core.Conj, n_1040)
												seq__604_1032, chunk__605_1033, count__606_1034, i__607_1035 = cljs_core.Next.Arity1IQ(seq__604_1038___1), nil, float64(0), float64(0)
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Deref.X_invoke_Arity1(s_1031), cljs_core.Reverse.X_invoke_Arity1(v_1030)) {
				} else {
					panic((&js.Error{("Assert failed: (= (clojure.core/deref s) (reverse v))")}))
				}
			}
			{
				var a_1041 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				var d_1042 = (&cljs_core.CljsCoreDelay{func(G__1043 *cljs_core.AFn, a_1041 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__1043, 0, func() interface{} {
						return cljs_core.Swap_BANG_.X_invoke_Arity2(a_1041, cljs_core.Inc)
					})
				}(&cljs_core.AFn{}, a_1041), nil})
				_, _ = a_1041, d_1042
				if cljs_core.Realized_QMARK_.Arity1IB(d_1042) == false {
				} else {
					panic((&js.Error{("Assert failed: (false? (realized? d))")}))
				}
				if cljs_core.Deref.X_invoke_Arity1(a_1041).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (clojure.core/deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_1042)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_1041)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref a))")}))
				}
				if cljs_core.Realized_QMARK_.Arity1IB(d_1042) == true {
				} else {
					panic((&js.Error{("Assert failed: (true? (realized? d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_1042)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref d))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_1041)) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 (clojure.core/deref a))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Force.X_invoke_Arity1(d_1042), cljs_core.Deref.X_invoke_Arity1(d_1042)) {
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
				var f_1044 = cljs_core.Memoize.X_invoke_Arity1(func(G__1045 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1045, 0, func() interface{} {
						return cljs_core.Rand.Arity0F()
					})
				}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn)
				_ = f_1044
				{
					f_1044.X_invoke_Arity0()
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return f_1044.X_invoke_Arity0()
				}(), func() interface{} {
					return f_1044.X_invoke_Arity0()
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(9))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(8))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(6))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(5))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(4))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 10) (list 0 1 2 3 4 5 6 7 8 9))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity2(float64(10), float64(20)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(19))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(18))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(17))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(16))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(15))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(14))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(13))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(12))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(11))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(10))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 10 20) (list 10 11 12 13 14 15 16 17 18 19))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(10), float64(20), float64(2)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(18))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(16))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(14))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(12))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(10))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 10 20 2) (list 10 12 14 16 18))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Take.X_invoke_Arity2(float64(20), cljs_core.Range_.X_invoke_Arity0().(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(19))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(18))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(17))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(16))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(15))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(14))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(13))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(12))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(11))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(10))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(9))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(8))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(6))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(5))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(4))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (take 20 (range)) (list 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19))")}))
			}
			{
				var d_1046 = cljs_core.Group_by.X_invoke_Arity2(cljs_core.Second, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(6), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), float64(4), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), float64(2)}, nil}))
				_ = d_1046
				if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_1046, float64(1))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (count (get d 1)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_1046, float64(2))).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 (count (get d 2)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_1046, float64(4))).(float64)) {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__1047 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1047, 1, func(p1__59_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(6), float64(8), float64(10)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__1048 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1048, 1, func(p1__60_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__1049 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1049, 2, func(p1__61_SHARP_ interface{}, p2__62_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(5)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__1050 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1050, 2, func(p1__64_SHARP_ interface{}, p2__63_SHARP_ interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil}), cljs_core.Map_indexed.X_invoke_Arity2(func(G__1051 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1051, 2, func(p1__65_SHARP_ interface{}, p2__66_SHARP_ interface{}) interface{} {
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
						return func(G__1052 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1052, 0, func() interface{} {
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
				var method_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1084__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1085__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__
				return (&cljs_core.CljsCoreMultiFn{"bar", func(G__1053 *cljs_core.AFn, method_table__1081__auto__ *cljs_core.CljsCoreAtom, prefer_table__1082__auto__ *cljs_core.CljsCoreAtom, method_cache__1083__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1084__auto__ *cljs_core.CljsCoreAtom, hierarchy__1085__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1053, 2, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{x, y}, nil})
					})
				}(&cljs_core.AFn{}, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1085__auto__, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__})
			}()

			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}), func(G__1054 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1054, 2, func(x interface{}, y interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)})
				})
			}(&cljs_core.AFn{}))
			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})}, nil}), func(G__1055 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1055, 2, func(x interface{}, y interface{}) interface{} {
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
				var G__609 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				var G__610 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				_, _ = G__609, G__610
				return Bar.X_invoke_Arity2(G__609, G__610)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :rect-shape (bar :cljs.core-test/rect :cljs.core-test/rect))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)}), cljs_core.Apply.X_invoke_Arity2(Bar.X_get_method_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= :rect-shape (apply (-get-method bar [:cljs.core-test/rect :cljs.core-test/shape]) [:cljs.core-test/rect :cljs.core-test/shape]))")}))
			}
			Nested_dispatch = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1084__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1085__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__
				return (&cljs_core.CljsCoreMultiFn{"nested-dispatch", func(G__1056 *cljs_core.AFn, method_table__1081__auto__ *cljs_core.CljsCoreAtom, prefer_table__1082__auto__ *cljs_core.CljsCoreAtom, method_cache__1083__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1084__auto__ *cljs_core.CljsCoreAtom, hierarchy__1085__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1056, 1, func(m interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(m))
					})
				}(&cljs_core.AFn{}, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1085__auto__, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__})
			}()

			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), func(G__1057 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1057, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1058 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1058, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__611 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil})
				_ = G__611
				return Nested_dispatch.X_invoke_Arity1(G__611)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :nested-a (nested-dispatch {:a {:b :c}}))")}))
			}
			Nested_dispatch2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1084__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1085__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__
				return (&cljs_core.CljsCoreMultiFn{"nested-dispatch2", cljs_core.Ffirst, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1085__auto__, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__})
			}()

			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__1059 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1059, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1060 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1060, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__612 = (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})}, nil})
				_ = G__612
				return Nested_dispatch2.X_invoke_Arity1(G__612)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :nested-a (nested-dispatch2 [[:a :b]]))")}))
			}
			Foo1 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1084__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1085__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__
				return (&cljs_core.CljsCoreMultiFn{"foo1", func(G__1061 *cljs_core.AFn, method_table__1081__auto__ *cljs_core.CljsCoreAtom, prefer_table__1082__auto__ *cljs_core.CljsCoreAtom, method_cache__1083__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1084__auto__ *cljs_core.CljsCoreAtom, hierarchy__1085__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1061, 0, func(args__ ...interface{}) interface{} {
						var args = cljs_core.Seq.Arity1IQ(args__[0])
						_ = args
						return cljs_core.First.X_invoke_Arity1(args)
					})
				}(&cljs_core.AFn{}, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1085__auto__, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__})
			}()

			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__1062 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1062, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)})
				})
			}(&cljs_core.AFn{}))
			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1063 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1063, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)}), func() interface{} {
				var G__613 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
				_ = G__613
				return Foo1.X_invoke_Arity1(G__613)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :a-return (foo1 :a))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)}), func() interface{} {
				var G__614 = float64(1)
				_ = G__614
				return Foo1.X_invoke_Arity1(G__614)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :default-return (foo1 1))")}))
			}
			Area = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1084__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1085__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__
				return (&cljs_core.CljsCoreMultiFn{"area", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Shape", Fqn: "Shape", X_hash: float64(-248980586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1085__auto__, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__})
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

			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Rect", Fqn: "Rect", X_hash: float64(-420704620)}), func(G__1064 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1064, 1, func(r interface{}) interface{} {
					return ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wd", Fqn: "wd", X_hash: float64(-183204751)}).X_invoke_Arity1(r).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ht", Fqn: "ht", X_hash: float64(214115472)}).X_invoke_Arity1(r).(float64))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Circle", Fqn: "Circle", X_hash: float64(-158896930)}), func(G__1065 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1065, 1, func(c interface{}) interface{} {
					return (Math.PI * ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64)))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1066 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1066, 1, func(x interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)})
				})
			}(&cljs_core.AFn{}))
			R = Rect.X_invoke_Arity2(float64(4), float64(13)).(cljs_core.CljsCoreIMap)

			C = Circle.X_invoke_Arity1(float64(12)).(cljs_core.CljsCoreIMap)

			if cljs_core.X_EQ_.Arity2IIB(float64(52), func() interface{} {
				var G__615 = R
				_ = G__615
				return Area.X_invoke_Arity1(G__615)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= 52 (area r))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)}), func() interface{} {
				var G__616 = cljs_core.CljsCorePersistentArrayMap_EMPTY
				_ = G__616
				return Area.X_invoke_Arity1(G__616)
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :oops (area {}))")}))
			}
			Foo2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1084__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1085__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__
				return (&cljs_core.CljsCoreMultiFn{"foo2", func(G__1067 *cljs_core.AFn, method_table__1081__auto__ *cljs_core.CljsCoreAtom, prefer_table__1082__auto__ *cljs_core.CljsCoreAtom, method_cache__1083__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1084__auto__ *cljs_core.CljsCoreAtom, hierarchy__1085__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1067, 0, func() interface{} {
						return nil
					})
				}(&cljs_core.AFn{}, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1085__auto__, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__})
			}()

			Foo2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1068 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1068, 0, func() interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), func() interface{} {
				return Foo2.X_invoke_Arity0()
			}()) {
			} else {
				panic((&js.Error{("Assert failed: (= :foo (foo2))")}))
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
				var method_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1084__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1085__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__
				return (&cljs_core.CljsCoreMultiFn{"apply-multi-test", func(G__1069 *cljs_core.AFn, method_table__1081__auto__ *cljs_core.CljsCoreAtom, prefer_table__1082__auto__ *cljs_core.CljsCoreAtom, method_cache__1083__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1084__auto__ *cljs_core.CljsCoreAtom, hierarchy__1085__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1069, 3, func(___ interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
						return float64(0)
					})
				}(&cljs_core.AFn{}, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1085__auto__, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__})
			}()

			Apply_multi_test.X_add_method_Arity3(float64(0), func(G__1070 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1070, 2, func(x interface{}) interface{} {
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
				var method_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1084__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1085__auto__ = cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), My_map_hierarchy}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__
				return (&cljs_core.CljsCoreMultiFn{"my-map?", cljs_core.Type_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1085__auto__, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__})
			}()

			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)}), func(G__1071 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1071, 1, func(___ interface{}) interface{} {
					return true
				})
			}(&cljs_core.AFn{}))
			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1072 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1072, 1, func(___ interface{}) interface{} {
					return false
				})
			}(&cljs_core.AFn{}))
			{
				var seq__617_1073 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__618_1074 interface{} = nil
				var count__619_1075 = float64(0)
				var i__620_1076 = float64(0)
				_, _, _, _ = seq__617_1073, chunk__618_1074, count__619_1075, i__620_1076
				for {
					if i__620_1076 < count__619_1075 {
						{
							var m_1077 = cljs_core.Decorate_(chunk__618_1074).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__620_1076)
							_ = m_1077
							if cljs_core.Truth_(func() interface{} {
								var G__621 = m_1077
								_ = G__621
								return My_map_QMARK_.X_invoke_Arity1(G__621)
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (my-map? m)")}))
							}
							seq__617_1073, chunk__618_1074, count__619_1075, i__620_1076 = seq__617_1073, chunk__618_1074, count__619_1075, (i__620_1076 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1078 = cljs_core.Seq.Arity1IQ(seq__617_1073)
							_ = temp__4388__auto___1078
							if cljs_core.Truth_(temp__4388__auto___1078) {
								{
									var seq__617_1079___1 = temp__4388__auto___1078
									_ = seq__617_1079___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__617_1079___1) {
										{
											var c__971__auto___1080 = cljs_core.Chunk_first.X_invoke_Arity1(seq__617_1079___1)
											_ = c__971__auto___1080
											seq__617_1073, chunk__618_1074, count__619_1075, i__620_1076 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__617_1079___1), c__971__auto___1080, cljs_core.Count.X_invoke_Arity1(c__971__auto___1080).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_1081 = cljs_core.First.X_invoke_Arity1(seq__617_1079___1)
											_ = m_1081
											if cljs_core.Truth_(func() interface{} {
												var G__622 = m_1081
												_ = G__622
												return My_map_QMARK_.X_invoke_Arity1(G__622)
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (my-map? m)")}))
											}
											seq__617_1073, chunk__618_1074, count__619_1075, i__620_1076 = cljs_core.Next.Arity1IQ(seq__617_1079___1), nil, float64(0), float64(0)
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
				var seq__623_1082 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentVector_EMPTY, float64(1), "asdf", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}))
				var chunk__624_1083 interface{} = nil
				var count__625_1084 = float64(0)
				var i__626_1085 = float64(0)
				_, _, _, _ = seq__623_1082, chunk__624_1083, count__625_1084, i__626_1085
				for {
					if i__626_1085 < count__625_1084 {
						{
							var not_m_1086 = cljs_core.Decorate_(chunk__624_1083).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__626_1085)
							_ = not_m_1086
							if cljs_core.Not.Arity1IB(func() interface{} {
								var G__627 = not_m_1086
								_ = G__627
								return My_map_QMARK_.X_invoke_Arity1(G__627)
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (not (my-map? not-m))")}))
							}
							seq__623_1082, chunk__624_1083, count__625_1084, i__626_1085 = seq__623_1082, chunk__624_1083, count__625_1084, (i__626_1085 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1087 = cljs_core.Seq.Arity1IQ(seq__623_1082)
							_ = temp__4388__auto___1087
							if cljs_core.Truth_(temp__4388__auto___1087) {
								{
									var seq__623_1088___1 = temp__4388__auto___1087
									_ = seq__623_1088___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__623_1088___1) {
										{
											var c__971__auto___1089 = cljs_core.Chunk_first.X_invoke_Arity1(seq__623_1088___1)
											_ = c__971__auto___1089
											seq__623_1082, chunk__624_1083, count__625_1084, i__626_1085 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__623_1088___1), c__971__auto___1089, cljs_core.Count.X_invoke_Arity1(c__971__auto___1089).(float64), float64(0)
											continue
										}
									} else {
										{
											var not_m_1090 = cljs_core.First.X_invoke_Arity1(seq__623_1088___1)
											_ = not_m_1090
											if cljs_core.Not.Arity1IB(func() interface{} {
												var G__628 = not_m_1090
												_ = G__628
												return My_map_QMARK_.X_invoke_Arity1(G__628)
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (not (my-map? not-m))")}))
											}
											seq__623_1082, chunk__624_1083, count__625_1084, i__626_1085 = cljs_core.Next.Arity1IQ(seq__623_1088___1), nil, float64(0), float64(0)
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(10), float64(3)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(9))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(6))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 0 10 3) (list 0 3 6 9))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(10), float64(3)).(*cljs_core.CljsCoreRange)).(float64), float64(4)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count (range 0 10 3)) 4)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(-10), float64(-3)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(-9))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-6))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 0 -10 -3) (list 0 -3 -6 -9))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(-10), float64(-3)).(*cljs_core.CljsCoreRange)).(float64), float64(4)) {
			} else {
				panic((&js.Error{("Assert failed: (= (count (range 0 -10 -3)) 4)")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(-10), float64(10), float64(3)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(5))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-4))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-10))) {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(-3), float64(-1)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(-2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= (range 0 -3 -1) (list 0 -1 -2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(3), float64(0), float64(-1)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.Range_.X_invoke_Arity3(float64(1), float64(0), float64(0)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= (take 3 (range 1 0 0)) (list 1 1 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.Range_.X_invoke_Arity3(float64(3), float64(1), float64(0)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
			} else {
				panic((&js.Error{("Assert failed: (= (take 3 (range 3 1 0)) (list 3 3 3))")}))
			}
			{
				var pv_1091 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange))
				_ = pv_1091
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity2(pv_1091, float64(96)), float64(96)) {
				} else {
					panic((&js.Error{("Assert failed: (= (nth pv 96) 96)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity3(pv_1091, float64(97), nil), nil) {
				} else {
					panic((&js.Error{("Assert failed: (= (nth pv 97 nil) nil)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__629 = float64(96)
					_ = G__629
					return pv_1091.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__629)
				}(), float64(96)) {
				} else {
					panic((&js.Error{("Assert failed: (= (pv 96) 96)")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(cljs_core.CljsCorePersistentVector_EMPTY)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq []))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reverse.X_invoke_Arity1(pv_1091), cljs_core.Rseq.Arity1IQ(pv_1091)) {
				} else {
					panic((&js.Error{("Assert failed: (= (reverse pv) (rseq pv))")}))
				}
			}
			{
				var pv_1092 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(33)).(*cljs_core.CljsCoreRange))
				_ = pv_1092
				if cljs_core.X_EQ_.Arity2IIB(pv_1092, cljs_core.Conj.X_invoke_Arity2(cljs_core.Conj.X_invoke_Arity2(cljs_core.Pop.X_invoke_Arity1(cljs_core.Pop.X_invoke_Arity1(pv_1092)), float64(31)), float64(32))) {
				} else {
					panic((&js.Error{("Assert failed: (= pv (-> pv pop pop (conj 31) (conj 32)))")}))
				}
			}
			{
				var stack1_1093 = cljs_core.Pop.X_invoke_Arity1(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange)))
				var stack2_1094 = cljs_core.Pop.X_invoke_Arity1(stack1_1093)
				_, _ = stack1_1093, stack2_1094
				if cljs_core.X_EQ_.Arity2IIB(float64(95), cljs_core.Peek.X_invoke_Arity1(stack1_1093)) {
				} else {
					panic((&js.Error{("Assert failed: (= 95 (peek stack1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(94), cljs_core.Peek.X_invoke_Arity1(stack2_1094)) {
				} else {
					panic((&js.Error{("Assert failed: (= 94 (peek stack2))")}))
				}
			}
			{
				var sentinel_1095 = cljs_core.Rand.Arity0F()
				_ = sentinel_1095
				if reflect.DeepEqual(sentinel_1095, func() (return__1096 interface{}) {
					defer func() {
						if e630 := recover(); e630 != nil {
							if cljs_core.Value_(e630).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var ___ = e630
									_ = ___
									return__1096 = sentinel_1095
								}
							} else {
								panic(e630)

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
				var v1_1097 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				var v2_1098 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(5)).(*cljs_core.CljsCoreRange))
				var s_1099 = cljs_core.Subvec.X_invoke_Arity3(v1_1097, float64(2), float64(8)).(*cljs_core.CljsCoreSubvec)
				_, _, _ = v1_1097, v2_1098, s_1099
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(s_1099, cljs_core.Subvec.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2(v1_1097, float64(2)).(*cljs_core.CljsCoreSubvec), float64(0), float64(6)).(*cljs_core.CljsCoreSubvec), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Take.X_invoke_Arity2(float64(6), cljs_core.Drop.X_invoke_Arity2(float64(2), v1_1097).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= s (-> v1 (subvec 2) (subvec 0 6)) (->> v1 (drop 2) (take 6)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Count.X_invoke_Arity1(s_1099).(float64)) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (count s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Pop.X_invoke_Arity1(s_1099)) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 4 5 6] (pop s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(7), cljs_core.Peek.X_invoke_Arity1(s_1099)) {
				} else {
					panic((&js.Error{("Assert failed: (= 7 (peek s))")}))
				}
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(1)}, nil}), cljs_core.Assoc.X_invoke_Arity3(s_1099, float64(6), float64(1)), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Conj.X_invoke_Arity2(s_1099, float64(1))}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 4 5 6 7 1] (assoc s 6 1) (conj s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(27), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, s_1099)) {
				} else {
					panic((&js.Error{("Assert failed: (= 27 (reduce + s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(s_1099, cljs_core.Vec.X_invoke_Arity1(s_1099)) {
				} else {
					panic((&js.Error{("Assert failed: (= s (vec s))")}))
				}
				{
					var m_1100 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), float64(1)}, nil})
					_ = m_1100
					if cljs_core.X_EQ_.Arity2IIB(m_1100, cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(s_1099, m_1100))) {
					} else {
						panic((&js.Error{("Assert failed: (= m (meta (with-meta s m)))")}))
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1101 interface{}) {
					defer func() {
						if e631 := recover(); e631 != nil {
							if cljs_core.Value_(e631).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e631
									_ = e
									return__1101 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e631)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1098, float64(0), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 0 6) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1102 interface{}) {
					defer func() {
						if e632 := recover(); e632 != nil {
							if cljs_core.Value_(e632).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e632
									_ = e
									return__1102 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e632)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1098, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1103 interface{}) {
					defer func() {
						if e633 := recover(); e633 != nil {
							if cljs_core.Value_(e633).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e633
									_ = e
									return__1103 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e633)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1098, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1104 interface{}) {
					defer func() {
						if e634 := recover(); e634 != nil {
							if cljs_core.Value_(e634).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e634
									_ = e
									return__1104 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e634)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1098, float64(3), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{("Assert failed: (= :fail (try (subvec v2 3 6) (catch js/Error e :fail)))")}))
				}
				if reflect.DeepEqual(v1_1097, cljs_core.Subvec.X_invoke_Arity3(s_1099, float64(1), float64(4)).(*cljs_core.CljsCoreSubvec).V) {
				} else {
					panic((&js.Error{("Assert failed: (identical? v1 (.-v (subvec s 1 4)))")}))
				}
				{
					var sentinel_1105 = cljs_core.Rand.Arity0F()
					var s_1106___1 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2))
					_, _ = sentinel_1105, s_1106___1
					if reflect.DeepEqual(sentinel_1105, func() (return__1107 interface{}) {
						defer func() {
							if e635 := recover(); e635 != nil {
								if cljs_core.Value_(e635).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var ___ = e635
										_ = ___
										return__1107 = sentinel_1105
									}
								} else {
									panic(e635)

								}
							}
						}()
						{
							{
								var G__636 = float64(-1)
								_ = G__636
								return s_1106___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__636)
							}
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (identical? sentinel (try (s -1) (catch js/Error _ sentinel)))")}))
					}
					if reflect.DeepEqual(sentinel_1105, func() (return__1108 interface{}) {
						defer func() {
							if e637 := recover(); e637 != nil {
								if cljs_core.Value_(e637).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var ___ = e637
										_ = ___
										return__1108 = sentinel_1105
									}
								} else {
									panic(e637)

								}
							}
						}()
						{
							{
								var G__638 = float64(1)
								_ = G__638
								return s_1106___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__638)
							}
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (identical? sentinel (try (s 1) (catch js/Error _ sentinel)))")}))
					}
				}
				{
					var sv1_1109 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2)).(*cljs_core.CljsCoreSubvec)
					var sv2_1110 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(1)).(*cljs_core.CljsCoreSubvec)
					_, _ = sv1_1109, sv2_1110
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(sv1_1109), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(*cljs_core.CljsCoreList)) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq sv1) (quote (1)))")}))
					}
					if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(sv2_1110)) {
					} else {
						panic((&js.Error{("Assert failed: (nil? (rseq sv2))")}))
					}
				}
			}
			{
				var v1_1111 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(15), float64(48)).(*cljs_core.CljsCoreRange))
				var v2_1112 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(40), float64(57)).(*cljs_core.CljsCoreRange))
				var v1_1113___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Pop_BANG_.X_invoke_Arity1(cljs_core.Transient.X_invoke_Arity1(v1_1111)), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v2_1114___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(v2_1112), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v_1115 = cljs_core.Into.X_invoke_Arity2(v1_1113___1, v2_1114___1)
				_, _, _, _, _ = v1_1111, v2_1112, v1_1113___1, v2_1114___1, v_1115
				if cljs_core.X_EQ_.Arity2IIB(v_1115, cljs_core.Vec.X_invoke_Arity1(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(16), float64(47)).(*cljs_core.CljsCoreRange), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(41), float64(57)).(*cljs_core.CljsCoreRange), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})})).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (= v (vec (concat [:quux] (range 16 47) [:foo] [:quux] (range 41 57) [:bar])))")}))
				}
			}
			{
				var v_1116 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)
				var xs_1117 interface{} = cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)
				_, _ = v_1116, xs_1117
				for {
					{
						var temp__4386__auto___1118 = cljs_core.First.X_invoke_Arity1(xs_1117)
						_ = temp__4386__auto___1118
						if cljs_core.Truth_(temp__4386__auto___1118) {
							{
								var x_1119 = temp__4386__auto___1118
								_ = x_1119
								v_1116, xs_1117 = func() interface{} {
									var pred__639 = func(G__1120 *cljs_core.AFn, v_1116 interface{}, xs_1117 interface{}, x_1119 interface{}, temp__4386__auto___1118 interface{}) *cljs_core.AFn {
										return cljs_core.Fn(G__1120, 2, func(p1__67_SHARP_ interface{}, p2__68_SHARP_ interface{}) interface{} {
											{
												var G__642 = cljs_core.Mod.X_invoke_Arity2(p2__68_SHARP_, float64(3)).(float64)
												_ = G__642
												return p1__67_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__642)
											}
										})
									}(&cljs_core.AFn{}, v_1116, xs_1117, x_1119, temp__4386__auto___1118)
									var expr__640 = x_1119
									_, _ = pred__639, expr__640
									if cljs_core.Truth_(pred__639.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(0), nil, float64(2), nil}, nil}, nil}), expr__640)) {
										return cljs_core.Conj_BANG_.X_invoke_Arity2(v_1116, x_1119)
									} else {
										if cljs_core.Truth_(pred__639.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), expr__640)) {
											return cljs_core.Assoc_BANG_.X_invoke_Arity3(v_1116, cljs_core.Count.X_invoke_Arity1(v_1116).(float64), x_1119)
										} else {
											panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(expr__640).(string))}))
										}
									}
								}(), cljs_core.Next.Arity1IQ(xs_1117)
								continue
							}
						} else {
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)), cljs_core.Persistent_BANG_.X_invoke_Arity1(v_1116)) {
							} else {
								panic((&js.Error{("Assert failed: (= (vec (range 100)) (persistent! v))")}))
							}
						}
					}
					break
				}
			}
			{
				var m1_1121 interface{} = cljs_core.CljsCorePersistentHashMap_EMPTY
				var m2_1122 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_EMPTY)
				var i_1123 = float64(0)
				_, _, _ = m1_1121, m2_1122, i_1123
				for {
					if i_1123 < float64(100) {
						m1_1121, m2_1122, i_1123 = cljs_core.Assoc.X_invoke_Arity3(m1_1121, i_1123, i_1123), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_1122, i_1123, i_1123), (i_1123 + float64(1))
						continue
					} else {
						{
							var m2_1124___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_1122)
							_ = m2_1124___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1121).(float64), float64(100)) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m1) 100)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1124___1).(float64), float64(100)) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m2) 100)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_1121, m2_1124___1) {
							} else {
								panic((&js.Error{("Assert failed: (= m1 m2)")}))
							}
							{
								var i_1125___1 = float64(0)
								_ = i_1125___1
								for {
									if i_1125___1 < float64(100) {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__643 = i_1125___1
											_ = G__643
											return m1_1121.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__643)
										}(), i_1125___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__644 = i_1125___1
											_ = G__644
											return m2_1124___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__644)
										}(), i_1125___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m2 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_1121, i_1125___1), i_1125___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_1124___1, i_1125___1), i_1125___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m2 i) i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_1121, i_1125___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m1 i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_1124___1, i_1125___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m2 i)")}))
										}
										i_1125___1 = (i_1125___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_1121))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range 100) (range 100)) (sort-by first (seq m1)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_1124___1))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range 100) (range 100)) (sort-by first (seq m2)))")}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_1121, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{("Assert failed: (not (contains? (dissoc m1 3) 3))")}))
							}
						}
					}
					break
				}
			}
			{
				var m_1126 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_1126
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1126).(float64), float64(7)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 7)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1126, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
				}
			}
			{
				var m_1127 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_1127
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1127).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1127, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var m_1128 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_1128
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1128).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1128, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var tm_1129 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_1129
				{
					var tm_1130___1 interface{} = tm_1129
					var ks_1131 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_1130___1, ks_1131
					for {
						{
							var temp__4386__auto___1132 = cljs_core.First.X_invoke_Arity1(ks_1131)
							_ = temp__4386__auto___1132
							if cljs_core.Truth_(temp__4386__auto___1132) {
								{
									var k_1133 = temp__4386__auto___1132
									_ = k_1133
									tm_1130___1, ks_1131 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1130___1, k_1133), cljs_core.Next.Arity1IQ(ks_1131)
									continue
								}
							} else {
								{
									var m_1134 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1130___1)
									_ = m_1134
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1134).(float64), float64(7)) {
									} else {
										panic((&js.Error{("Assert failed: (= (count m) 7)")}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_1134, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var tm_1135 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_1135
				{
					var seq__645_1136 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__646_1137 interface{} = nil
					var count__647_1138 = float64(0)
					var i__648_1139 = float64(0)
					_, _, _, _ = seq__645_1136, chunk__646_1137, count__647_1138, i__648_1139
					for {
						if i__648_1139 < count__647_1138 {
							{
								var k_1140 = cljs_core.Decorate_(chunk__646_1137).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__648_1139)
								_ = k_1140
								if cljs_core.X_EQ_.Arity2IIB(k_1140, cljs_core.Get.X_invoke_Arity2(tm_1135, k_1140)) {
								} else {
									panic((&js.Error{("Assert failed: (= k (get tm k))")}))
								}
								seq__645_1136, chunk__646_1137, count__647_1138, i__648_1139 = seq__645_1136, chunk__646_1137, count__647_1138, (i__648_1139 + float64(1))
								continue
							}
						} else {
							{
								var temp__4388__auto___1141 = cljs_core.Seq.Arity1IQ(seq__645_1136)
								_ = temp__4388__auto___1141
								if cljs_core.Truth_(temp__4388__auto___1141) {
									{
										var seq__645_1142___1 = temp__4388__auto___1141
										_ = seq__645_1142___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__645_1142___1) {
											{
												var c__971__auto___1143 = cljs_core.Chunk_first.X_invoke_Arity1(seq__645_1142___1)
												_ = c__971__auto___1143
												seq__645_1136, chunk__646_1137, count__647_1138, i__648_1139 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__645_1142___1), c__971__auto___1143, cljs_core.Count.X_invoke_Arity1(c__971__auto___1143).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_1144 = cljs_core.First.X_invoke_Arity1(seq__645_1142___1)
												_ = k_1144
												if cljs_core.X_EQ_.Arity2IIB(k_1144, cljs_core.Get.X_invoke_Arity2(tm_1135, k_1144)) {
												} else {
													panic((&js.Error{("Assert failed: (= k (get tm k))")}))
												}
												seq__645_1136, chunk__646_1137, count__647_1138, i__648_1139 = cljs_core.Next.Arity1IQ(seq__645_1142___1), nil, float64(0), float64(0)
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
					var m_1145 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1135)
					_ = m_1145
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1146 float64) {
						defer func() {
							if e649 := recover(); e649 != nil {
								if cljs_core.Value_(e649).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e649
										_ = e
										return__1146 = float64(2)
									}
								} else {
									panic(e649)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1135, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1147 float64) {
						defer func() {
							if e650 := recover(); e650 != nil {
								if cljs_core.Value_(e650).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e650
										_ = e
										return__1147 = float64(2)
									}
								} else {
									panic(e650)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_1135, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1148 float64) {
						defer func() {
							if e651 := recover(); e651 != nil {
								if cljs_core.Value_(e651).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e651
										_ = e
										return__1148 = float64(2)
									}
								} else {
									panic(e651)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1135)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (persistent! tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1149 float64) {
						defer func() {
							if e652 := recover(); e652 != nil {
								if cljs_core.Value_(e652).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e652
										_ = e
										return__1149 = float64(2)
									}
								} else {
									panic(e652)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_1135)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (count tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_1145, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var m_1150 = cljs_core.Assoc.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentHashMap_EMPTY, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_ = m_1150
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1150, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1150, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1150).(float64), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 2)")}))
				}
				{
					var m_1151___1 = cljs_core.Dissoc.X_invoke_Arity2(m_1150, Fixed_hash_foo)
					_ = m_1151___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1151___1, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1151___1, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1151___1).(float64), float64(1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 1)")}))
					}
				}
			}
			{
				var m_1152 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_1153___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m_1152, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_, _ = m_1152, m_1153___1
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1153___1).(float64), float64(102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 102)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1153___1, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1153___1, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				{
					var m_1154___2 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_1153___1, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_1154___2
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1154___2, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1154___2, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1154___2).(float64), float64(98)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 98)")}))
					}
				}
			}
			{
				var m_1155 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_1156___1 = cljs_core.Transient.X_invoke_Arity1(m_1155)
				var m_1157___2 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_1156___1, Fixed_hash_foo, float64(1))
				var m_1158___3 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_1157___2, Fixed_hash_bar, float64(2))
				var m_1159___4 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m_1158___3)
				_, _, _, _, _ = m_1155, m_1156___1, m_1157___2, m_1158___3, m_1159___4
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1159___4).(float64), float64(102)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 102)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1159___4, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-foo) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1159___4, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
				}
				{
					var m_1160___5 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_1159___4, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_1160___5
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1160___5, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get m fixed-hash-bar) 2)")}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1160___5, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{("Assert failed: (not (contains? m fixed-hash-foo))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1160___5).(float64), float64(98)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m) 98)")}))
					}
				}
			}
			Array_map_conversion_threshold = cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD

			{
				var m1_1161 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_1162 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY)
				var i_1163 = float64(0)
				_, _, _ = m1_1161, m2_1162, i_1163
				for {
					if i_1163 < Array_map_conversion_threshold {
						m1_1161, m2_1162, i_1163 = cljs_core.Assoc.X_invoke_Arity3(m1_1161, i_1163, i_1163), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_1162, i_1163, i_1163), (i_1163 + float64(1))
						continue
					} else {
						{
							var m2_1164___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_1162)
							_ = m2_1164___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1161).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m1) array-map-conversion-threshold)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1164___1).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{("Assert failed: (= (count m2) array-map-conversion-threshold)")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_1161, m2_1164___1) {
							} else {
								panic((&js.Error{("Assert failed: (= m1 m2)")}))
							}
							{
								var i_1165___1 = float64(0)
								_ = i_1165___1
								for {
									if i_1165___1 < Array_map_conversion_threshold {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__653 = i_1165___1
											_ = G__653
											return m1_1161.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__653)
										}(), i_1165___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__654 = i_1165___1
											_ = G__654
											return m2_1164___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__654)
										}(), i_1165___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (m2 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_1161, i_1165___1), i_1165___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m1 i) i)")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_1164___1, i_1165___1), i_1165___1) {
										} else {
											panic((&js.Error{("Assert failed: (= (get m2 i) i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_1161, i_1165___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m1 i)")}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_1164___1, i_1165___1) {
										} else {
											panic((&js.Error{("Assert failed: (contains? m2 i)")}))
										}
										i_1165___1 = (i_1165___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_1161))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m1)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_1164___1))) {
							} else {
								panic((&js.Error{("Assert failed: (= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m2)))")}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_1161, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{("Assert failed: (not (contains? (dissoc m1 3) 3))")}))
							}
						}
					}
					break
				}
			}
			{
				var m_1166 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_1166
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1166).(float64), float64(7)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 7)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1166, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
				}
			}
			{
				var m_1167 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_1167
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1167).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1167, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var m_1168 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_1168
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1168).(float64), float64(11)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 11)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1168, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{("Assert failed: (= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})")}))
				}
			}
			{
				var tm_1169 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_1169
				{
					var tm_1170___1 interface{} = tm_1169
					var ks_1171 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_1170___1, ks_1171
					for {
						{
							var temp__4386__auto___1172 = cljs_core.First.X_invoke_Arity1(ks_1171)
							_ = temp__4386__auto___1172
							if cljs_core.Truth_(temp__4386__auto___1172) {
								{
									var k_1173 = temp__4386__auto___1172
									_ = k_1173
									tm_1170___1, ks_1171 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1170___1, k_1173), cljs_core.Next.Arity1IQ(ks_1171)
									continue
								}
							} else {
								{
									var m_1174 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1170___1)
									_ = m_1174
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1174).(float64), float64(7)) {
									} else {
										panic((&js.Error{("Assert failed: (= (count m) 7)")}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_1174, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
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
				var tm_1175 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_1175
				{
					var seq__655_1176 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__656_1177 interface{} = nil
					var count__657_1178 = float64(0)
					var i__658_1179 = float64(0)
					_, _, _, _ = seq__655_1176, chunk__656_1177, count__657_1178, i__658_1179
					for {
						if i__658_1179 < count__657_1178 {
							{
								var k_1180 = cljs_core.Decorate_(chunk__656_1177).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__658_1179)
								_ = k_1180
								if cljs_core.X_EQ_.Arity2IIB(k_1180, cljs_core.Get.X_invoke_Arity2(tm_1175, k_1180)) {
								} else {
									panic((&js.Error{("Assert failed: (= k (get tm k))")}))
								}
								seq__655_1176, chunk__656_1177, count__657_1178, i__658_1179 = seq__655_1176, chunk__656_1177, count__657_1178, (i__658_1179 + float64(1))
								continue
							}
						} else {
							{
								var temp__4388__auto___1181 = cljs_core.Seq.Arity1IQ(seq__655_1176)
								_ = temp__4388__auto___1181
								if cljs_core.Truth_(temp__4388__auto___1181) {
									{
										var seq__655_1182___1 = temp__4388__auto___1181
										_ = seq__655_1182___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__655_1182___1) {
											{
												var c__971__auto___1183 = cljs_core.Chunk_first.X_invoke_Arity1(seq__655_1182___1)
												_ = c__971__auto___1183
												seq__655_1176, chunk__656_1177, count__657_1178, i__658_1179 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__655_1182___1), c__971__auto___1183, cljs_core.Count.X_invoke_Arity1(c__971__auto___1183).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_1184 = cljs_core.First.X_invoke_Arity1(seq__655_1182___1)
												_ = k_1184
												if cljs_core.X_EQ_.Arity2IIB(k_1184, cljs_core.Get.X_invoke_Arity2(tm_1175, k_1184)) {
												} else {
													panic((&js.Error{("Assert failed: (= k (get tm k))")}))
												}
												seq__655_1176, chunk__656_1177, count__657_1178, i__658_1179 = cljs_core.Next.Arity1IQ(seq__655_1182___1), nil, float64(0), float64(0)
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
					var m_1185 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1175)
					_ = m_1185
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1186 float64) {
						defer func() {
							if e659 := recover(); e659 != nil {
								if cljs_core.Value_(e659).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e659
										_ = e
										return__1186 = float64(2)
									}
								} else {
									panic(e659)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1175, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1187 float64) {
						defer func() {
							if e660 := recover(); e660 != nil {
								if cljs_core.Value_(e660).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e660
										_ = e
										return__1187 = float64(2)
									}
								} else {
									panic(e660)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_1175, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1188 float64) {
						defer func() {
							if e661 := recover(); e661 != nil {
								if cljs_core.Value_(e661).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e661
										_ = e
										return__1188 = float64(2)
									}
								} else {
									panic(e661)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1175)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (persistent! tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1189 float64) {
						defer func() {
							if e662 := recover(); e662 != nil {
								if cljs_core.Value_(e662).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e662
										_ = e
										return__1189 = float64(2)
									}
								} else {
									panic(e662)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_1175)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{("Assert failed: (= 2 (try (count tm) 1 (catch js/Error e 2)))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_1185, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})")}))
					}
				}
			}
			{
				var m_1190 = cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))
				_ = m_1190
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1190).(float64), (float64(2) * Array_map_conversion_threshold)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) (* 2 array-map-conversion-threshold))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__663 = Array_map_conversion_threshold
					_ = G__663
					return m_1190.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__663)
				}(), Array_map_conversion_threshold) {
				} else {
					panic((&js.Error{("Assert failed: (= (m array-map-conversion-threshold) array-map-conversion-threshold)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1190, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity2(func(G__1191 *cljs_core.AFn, m_1190 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1191, 1, func(p1__69_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__69_SHARP_, p1__69_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1190), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{("Assert failed: (= m (into (.-EMPTY cljs.core/PersistentHashMap) (map (fn* [p1__69#] (vector p1__69# p1__69#)) (range (* 2 array-map-conversion-threshold)))))")}))
				}
			}
			{
				var m1_1192 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_1193 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var i_1194 = float64(0)
				_, _, _ = m1_1192, m2_1193, i_1194
				for {
					if i_1194 < float64(100) {
						m1_1192, m2_1193, i_1194 = cljs_core.Assoc.X_invoke_Arity3(m1_1192, i_1194, i_1194), cljs_core.Assoc.X_invoke_Arity3(m2_1193, ("foo"+cljs_core.Str.X_invoke_Arity1(i_1194).(string)), i_1194), (i_1194 + float64(1))
						continue
					} else {
						if cljs_core.X_EQ_.Arity2IIB(m1_1192, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{("Assert failed: (= m1 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (range 100) (range 100))))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(m2_1193, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Map_.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.Str, "foo").(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{("Assert failed: (= m2 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (map (partial str \"foo\") (range 100)) (range 100))))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1192).(float64), float64(100)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count m1) 100)")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1193).(float64), float64(100)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count m2) 100)")}))
						}
					}
					break
				}
			}
			{
				var i_1195 = float64(0)
				var m_1196 interface{} = cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(-1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))
				var result_1197 interface{} = cljs_core.CljsCorePersistentVector_EMPTY
				_, _, _ = i_1195, m_1196, result_1197
				for {
					if i_1195 <= (cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) {
						i_1195, m_1196, result_1197 = (i_1195 + float64(1)), cljs_core.Assoc.X_invoke_Arity3(m_1196, i_1195, i_1195), cljs_core.Conj.X_invoke_Arity2(result_1197, cljs_core.Meta.X_invoke_Arity1(m_1196))
						continue
					} else {
						{
							var n_1198 = ((cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) + float64(1))
							var expected_1199 = cljs_core.Repeat.X_invoke_Arity2(n_1198, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})).(*cljs_core.CljsCoreLazySeq)
							_, _ = n_1198, expected_1199
							if cljs_core.X_EQ_.Arity2IIB(result_1197, expected_1199) {
							} else {
								panic((&js.Error{("Assert failed: (= result expected)")}))
							}
						}
					}
					break
				}
			}
			{
				var s_1200 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_EMPTY)
				var i_1201 = float64(0)
				_, _ = s_1200, i_1201
				for {
					if i_1201 < float64(100) {
						s_1200, i_1201 = cljs_core.Conj_BANG_.X_invoke_Arity2(s_1200, i_1201), (i_1201 + float64(1))
						continue
					} else {
						{
							var s_1202___1 interface{} = s_1200
							var i_1203___1 = float64(0)
							_, _ = s_1202___1, i_1203___1
							for {
								if i_1203___1 < float64(100) {
									if cljs_core.Mod.X_invoke_Arity2(i_1203___1, float64(3)).(float64) == float64(0) {
										s_1202___1, i_1203___1 = cljs_core.Disj_BANG_.X_invoke_Arity2(s_1202___1, i_1203___1), (i_1203___1 + float64(1))
										continue
									} else {
										s_1202___1, i_1203___1 = s_1202___1, (i_1203___1 + float64(1))
										continue
									}
								} else {
									{
										var s_1204___2 = cljs_core.Persistent_BANG_.X_invoke_Arity1(s_1202___1)
										_ = s_1204___2
										if cljs_core.X_EQ_.Arity2IIB(s_1204___2, func() interface{} {
											var s___3 interface{} = cljs_core.CljsCorePersistentHashSet_EMPTY
											var xs interface{} = cljs_core.Remove.X_invoke_Arity2(func(G__1205 *cljs_core.AFn, s_1202___1 interface{}, i_1203___1 float64, s_1200 interface{}, i_1201 float64, s___3 interface{}, s_1204___2 interface{}) *cljs_core.AFn {
												return cljs_core.Fn(G__1205, 1, func(p1__70_SHARP_ interface{}) interface{} {
													return (cljs_core.Mod.X_invoke_Arity2(p1__70_SHARP_, float64(3)).(float64) == float64(0))
												})
											}(&cljs_core.AFn{}, s_1202___1, i_1203___1, s_1200, i_1201, s___3, s_1204___2), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
											_, _ = s___3, xs
											for {
												{
													var temp__4386__auto__ = cljs_core.First.X_invoke_Arity1(xs)
													_ = temp__4386__auto__
													if cljs_core.Truth_(temp__4386__auto__) {
														{
															var x = temp__4386__auto__
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
											panic((&js.Error{("Assert failed: (= s (loop [s #{} xs (remove (fn* [p1__70#] (zero? (mod p1__70# 3))) (range 100))] (if-let [x (first xs)] (recur (conj s x) (next xs)) s)))")}))
										}
										if cljs_core.X_EQ_.Arity2IIB(s_1204___2, cljs_core.Set.X_invoke_Arity1(cljs_core.Remove.X_invoke_Arity2(func(G__1206 *cljs_core.AFn, s_1202___1 interface{}, i_1203___1 float64, s_1200 interface{}, i_1201 float64, s_1204___2 interface{}) *cljs_core.AFn {
											return cljs_core.Fn(G__1206, 1, func(p1__71_SHARP_ interface{}) interface{} {
												return (cljs_core.Mod.X_invoke_Arity2(p1__71_SHARP_, float64(3)).(float64) == float64(0))
											})
										}(&cljs_core.AFn{}, s_1202___1, i_1203___1, s_1200, i_1201, s_1204___2), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
										} else {
											panic((&js.Error{("Assert failed: (= s (set (remove (fn* [p1__71#] (zero? (mod p1__71# 3))) (range 100))))")}))
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
				var m1_1207 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_1208 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var m2_1209 = cljs_core.Sorted_map_by.X_invoke_ArityVariadic(c2_1208, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{})).(*cljs_core.CljsCorePersistentTreeMap)
				_, _, _ = m1_1207, c2_1208, m2_1209
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeMap)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(m1_1207)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeMap (type m1))")}))
				}
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeMap)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(m2_1209)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeMap (type m2))")}))
				}
				if reflect.DeepEqual(cljs_core.Compare, cljs_core.Native_get_instance_field.X_invoke_Arity2(m1_1207, "Comp")) {
				} else {
					panic((&js.Error{("Assert failed: (identical? compare (.-comp m1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(m1_1207).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count m1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(m2_1209).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count m2))")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(m1_1207)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq m1))")}))
				}
				{
					var m1_1210___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m1_1207, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					var m2_1211___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m2_1209, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					_, _ = m1_1210___1, m2_1211___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1210___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m1) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1211___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count m2) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_1210___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m1) (list [:bar 2] [:foo 1] [:quux 3]))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_1211___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m2) (list [:quux 3] [:foo 1] [:bar 2]))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_1210___1), cljs_core.Rseq.Arity1IQ(m2_1211___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m1) (rseq m2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_1211___1), cljs_core.Rseq.Arity1IQ(m1_1210___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq m2) (rseq m1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m1_1210___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= (conj m1 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m1_1210___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count (conj m1 [:wibble 4])) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m2_1211___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{("Assert failed: (= (conj m2 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m2_1211___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count (conj m2 [:wibble 4])) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m1_1210___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}))).(cljs_core.CljsCoreICollection).X_conj_Arity2(nil)) {
					} else {
						panic((&js.Error{("Assert failed: (= (map key (assoc m1 nil 4)) (list nil :bar :foo :quux))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m2_1211___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(nil)).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (map key (assoc m2 nil 4)) (list :quux :foo :bar nil))")}))
					}
				}
			}
			{
				var m_1212 = cljs_core.Apply.X_invoke_Arity2(cljs_core.Sorted_map, cljs_core.Mapcat.X_invoke_ArityVariadic(func(G__1215 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1215, 1, func(p1__72_SHARP_ interface{}) interface{} {
						return cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(p1__72_SHARP_)).(cljs_core.CljsCoreICollection).X_conj_Arity2(p1__72_SHARP_)
					})
				}(&cljs_core.AFn{}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_ArityVariadic(cljs_core.Partial.X_invoke_Arity2(cljs_core.Apply, cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(10)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(20), float64(30)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(50), float64(60)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(30), float64(40)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(40), float64(50)}, nil})}, nil})}))})))
				var s1_1213 = cljs_core.Map_.X_invoke_Arity2(func(G__1216 *cljs_core.AFn, m_1212 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1216, 1, func(p1__73_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__73_SHARP_, p1__73_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1212), cljs_core.Range_.X_invoke_Arity1(float64(60)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				var s2_1214 = cljs_core.Map_.X_invoke_Arity2(func(G__1217 *cljs_core.AFn, m_1212 interface{}, s1_1213 *cljs_core.CljsCoreLazySeq) *cljs_core.AFn {
					return cljs_core.Fn(G__1217, 1, func(p1__74_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__74_SHARP_, p1__74_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1212, s1_1213), cljs_core.Range_.X_invoke_Arity3(float64(59), float64(-1), float64(-1)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				_, _, _ = m_1212, s1_1213, s2_1214
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1212).(float64), float64(60)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count m) 60)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m_1212), s1_1213) {
				} else {
					panic((&js.Error{("Assert failed: (= (seq m) s1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(m_1212), s2_1214) {
				} else {
					panic((&js.Error{("Assert failed: (= (rseq m) s2)")}))
				}
			}
			{
				var m_1218 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
				_ = m_1218
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m_1218, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(2), float64(3)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (dissoc m :foo) (hash-map :bar 2 :quux 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(m_1218, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))).(float64), float64(2)) {
				} else {
					panic((&js.Error{("Assert failed: (= (count (dissoc m :foo)) 2)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(m_1218), cljs_core.Hash.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(1), float64(2), float64(3)}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (hash m) (hash (hash-map :foo 1 :bar 2 :quux 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1218, cljs_core.X_LT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m < :foo) (list [:bar 2]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1218, cljs_core.X_LT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m <= :foo) (list [:bar 2] [:foo 1]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1218, cljs_core.X_GT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m > :foo) (list [:quux 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1218, cljs_core.X_GT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= (subseq m >= :foo) (list [:foo 1] [:quux 3]))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1219 *cljs_core.AFn, m_1218 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1219, 1, func(p1__75_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity2(func(G__1220 *cljs_core.AFn, m_1218 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__1220, 2, func(___ interface{}, x interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_1218), p1__75_SHARP_)
					})
				}(&cljs_core.AFn{}, m_1218), m_1218).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__75#] (reduce (fn [_ x] x) p1__75#)) m) (list 2 1 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1221 *cljs_core.AFn, m_1218 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1221, 1, func(p1__76_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity3(func(G__1222 *cljs_core.AFn, m_1218 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__1222, 2, func(x interface{}, ___ interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_1218), float64(7), p1__76_SHARP_)
					})
				}(&cljs_core.AFn{}, m_1218), m_1218).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__76#] (reduce (fn [x _] x) 7 p1__76#)) m) (list 7 7 7))")}))
				}
			}
			{
				var s1_1223 = cljs_core.Sorted_set.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_1224 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var s2_1225 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c2_1224, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c3_1226 = func(G__1229 *cljs_core.AFn, s1_1223 interface{}, c2_1224 cljs_core.CljsCoreIFn, s2_1225 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1229, 2, func(p1__77_SHARP_ interface{}, p2__78_SHARP_ interface{}) interface{} {
						return cljs_core.Compare.Arity2IIF(cljs_core.Quot.X_invoke_Arity2(p1__77_SHARP_, float64(2)).(float64), cljs_core.Quot.X_invoke_Arity2(p2__78_SHARP_, float64(2)).(float64))
					})
				}(&cljs_core.AFn{}, s1_1223, c2_1224, s2_1225)
				var s3_1227 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c3_1226, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var s4_1228 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(cljs_core.X_LT_, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				_, _, _, _, _, _ = s1_1223, c2_1224, s2_1225, c3_1226, s3_1227, s4_1228
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(s1_1223)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeSet (type s1))")}))
				}
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(s2_1225)) {
				} else {
					panic((&js.Error{("Assert failed: (identical? cljs.core/PersistentTreeSet (type s2))")}))
				}
				if reflect.DeepEqual(cljs_core.Compare, cljs_core.Decorate_(s1_1223).(cljs_core.CljsCoreISorted).X_comparator_Arity1()) {
				} else {
					panic((&js.Error{("Assert failed: (identical? compare (-comparator s1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(s1_1223).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count s1))")}))
				}
				if cljs_core.Count.X_invoke_Arity1(s2_1225).(float64) == float64(0) {
				} else {
					panic((&js.Error{("Assert failed: (zero? (count s2))")}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(s1_1223)) {
				} else {
					panic((&js.Error{("Assert failed: (nil? (rseq s1))")}))
				}
				{
					var s1_1230___1 = cljs_core.Conj.X_invoke_ArityVariadic(s1_1223, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s2_1231___1 = cljs_core.Conj.X_invoke_ArityVariadic(s2_1225, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s3_1232___1 = cljs_core.Conj.X_invoke_ArityVariadic(s3_1227, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(7), float64(8), float64(9)}))
					var s4_1233___1 = cljs_core.Conj.X_invoke_ArityVariadic(s4_1228, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					_, _, _, _ = s1_1230___1, s2_1231___1, s3_1232___1, s4_1233___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_1230___1), cljs_core.Hash.X_invoke_Arity1(s2_1231___1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (hash s1) (hash s2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_1230___1), cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}))) {
					} else {
						panic((&js.Error{("Assert failed: (= (hash s1) (hash #{1 3 2}))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_1230___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq s1) (list 1 2 3))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_1230___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq s1) (list 3 2 1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_1231___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{("Assert failed: (= (seq s2) (list 3 2 1))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_1231___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (rseq s2) (list 1 2 3))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_1230___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s1) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_1231___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s2) 3)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s3_1232___1).(float64), float64(4)) {
					} else {
						panic((&js.Error{("Assert failed: (= (count s3) 4)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(s3_1232___1, float64(0)), float64(1)) {
					} else {
						panic((&js.Error{("Assert failed: (= (get s3 0) 1)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1232___1, cljs_core.X_GT_, float64(5)), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 > 5) (list 7 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1232___1, cljs_core.X_GT_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 > 6) (list 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1232___1, cljs_core.X_GT__EQ_, float64(6)), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 6) (list 7 8))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1232___1, cljs_core.X_GT__EQ_, float64(12)), nil) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 12) nil)")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1232___1, cljs_core.X_LT_, float64(0)), cljs_core.CljsCoreList_EMPTY) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 0) (list))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1232___1, cljs_core.X_LT_, float64(5)), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 5) (list 1 2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1232___1, cljs_core.X_LT_, float64(6)), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 < 6) (list 1 2))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1232___1, cljs_core.X_LT__EQ_, float64(6)), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 <= 6) (list 1 2 7))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s3_1232___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT__EQ_, float64(6)), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s3 >= 2 <= 6) (list 2 7))")}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s4_1233___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT_, float64(3)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{("Assert failed: (= (subseq s4 >= 2 < 3) (list 2))")}))
					}
					{
						var s1_1234___2 = cljs_core.Disj.X_invoke_Arity2(s1_1230___1, float64(2))
						var s2_1235___2 = cljs_core.Disj.X_invoke_Arity2(s2_1231___1, float64(2))
						_, _ = s1_1234___2, s2_1235___2
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_1234___2), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{("Assert failed: (= (seq s1) (list 1 3))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_1234___2), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{("Assert failed: (= (rseq s1) (list 3 1))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_1235___2), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{("Assert failed: (= (seq s2) (list 3 1))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_1235___2), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{("Assert failed: (= (rseq s2) (list 1 3))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_1234___2).(float64), float64(2)) {
						} else {
							panic((&js.Error{("Assert failed: (= (count s1) 2)")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_1235___2).(float64), float64(2)) {
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
					return cljs_core.Fn(map__GT_Person, 1, func(G__666 interface{}) interface{} {
						return (&CljsCore_testPerson{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}).X_invoke_Arity1(G__666), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}).X_invoke_Arity1(G__666), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__666, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})})), nil})
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
					return cljs_core.Fn(map__GT_A, 1, func(G__685 interface{}) interface{} {
						return (&CljsCore_testA{nil, cljs_core.Dissoc.X_invoke_Arity1(G__685), nil})
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
					return cljs_core.Fn(map__GT_C, 1, func(G__696 interface{}) interface{} {
						return (&CljsCore_testC{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__696), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__696), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}).X_invoke_Arity1(G__696), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__696, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})})), nil})
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
				var s_1236 = "abc"
				_ = s_1236
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_1236).Length) {
				} else {
					panic((&js.Error{("Assert failed: (= 3 (.-length s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_1236).Length) {
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
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1236).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_("abc").Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring \"abc\" 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", func(G__1237 *cljs_core.AFn, s_1236 string) *cljs_core.AFn {
					return cljs_core.Fn(G__1237, 2, func(target717 interface{}, start interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target717, "Substring", []interface{}{start})
					})
				}(&cljs_core.AFn{}, s_1236).X_invoke_Arity2(s_1236, float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" ((memfn substring start) s 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1236).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s substring 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1236).Substring(float64(1))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s (substring 1)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1236).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (. s (substring 1 3)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1236).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{("Assert failed: (= \"bc\" (.substring s 1 3))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_(s_1236).ToUpperCase()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" (. s (toUpperCase)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_("abc").ToUpperCase()) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" (. \"abc\" (toUpperCase)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", func(G__1238 *cljs_core.AFn, s_1236 string) *cljs_core.AFn {
					return cljs_core.Fn(G__1238, 1, func(target718 interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target718, "ToUpperCase", []interface{}{})
					})
				}(&cljs_core.AFn{}, s_1236).X_invoke_Arity1(s_1236)) {
				} else {
					panic((&js.Error{("Assert failed: (= \"ABC\" ((memfn toUpperCase) s))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB("BC", cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_1236).ToUpperCase(), "Substring", []interface{}{float64(1)})) {
				} else {
					panic((&js.Error{("Assert failed: (= \"BC\" (. (. s (toUpperCase)) substring 1))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_1236).ToUpperCase(), "Substring", []interface{}{float64(1)}), "Length")) {
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
					return cljs_core.Fn(map__GT_A2, 1, func(G__721 interface{}) interface{} {
						return (&CljsCore_testA2{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__721), nil, cljs_core.Dissoc.X_invoke_Arity2(G__721, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
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
					return cljs_core.Fn(map__GT_B, 1, func(G__736 interface{}) interface{} {
						return (&CljsCore_testB{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__736), nil, cljs_core.Dissoc.X_invoke_Arity2(G__736, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
					})
				}(&cljs_core.AFn{})

			}
			if cljs_core.Not_EQ_.Arity2IIB((&CljsCore_testA2{nil, nil, nil, nil}), (&CljsCore_testB{nil, nil, nil, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (not= (A2. nil) (B. nil))")}))
			}
			Foo = func(foo *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(foo, 1, func(this interface{}) interface{} {
					return cljs_core.Decorate_(this).(CljsCore_testIFoo).Foo_Arity1()
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(func() *CljsCore_testT749 {
				X__GT_t749 = func(__GT_t749 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_t749, 2, func(test_stuff___1 interface{}, meta750 interface{}) interface{} {
						return (&CljsCore_testT749{test_stuff___1, meta750})
					})
				}(&cljs_core.AFn{})

				return (&CljsCore_testT749{test_stuff, nil})
			}(), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (meta (with-meta (reify IFoo (foo [this] :foo)) {:foo :bar})) {:foo :bar})")}))
			}
			Foo2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1084__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1085__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__, hierarchy__1085__auto__
				return (&cljs_core.CljsCoreMultiFn{"foo2", cljs_core.Identity, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1085__auto__, method_table__1081__auto__, prefer_table__1082__auto__, method_cache__1083__auto__, cached_hierarchy__1084__auto__})
			}()

			Foo2.X_add_method_Arity3(float64(0), func(G__1239 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1239, 1, func(x interface{}) interface{} {
					return x
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB(Foo2, cljs_core.Ffirst.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{Foo2, float64(1)}, true, false).(*cljs_core.CljsCorePersistentArrayMap))) {
			} else {
				panic((&js.Error{("Assert failed: (= foo2 (ffirst {foo2 1}))")}))
			}
			Mutate = func(mutate *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(mutate, 1, func(this interface{}) interface{} {
					return cljs_core.Decorate_(this).(CljsCore_testIMutate).Mutate_Arity1()
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
				var g_1240 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}))}, true).(*cljs_core.CljsCorePersistentHashSet)
				var h_1241 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil})}, true).(*cljs_core.CljsCorePersistentHashSet)
				_, _ = g_1240, h_1241
				if cljs_core.X_EQ_.Arity2IIB(g_1240, h_1241) {
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
					return cljs_core.Decorate_(this).(CljsCore_testIHasFirst).X_get_first_Arity1()
				})
			}(&cljs_core.AFn{})

			X_find_first = func(_find_first *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(_find_first, 2, func(this interface{}, other interface{}) interface{} {
					return cljs_core.Decorate_(this).(CljsCore_testIFindsFirst).X_find_first_Arity2(other)
				})
			}(&cljs_core.AFn{})

			X__GT_First = func(__GT_First *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_First, 1, func(xs interface{}) interface{} {
					return (&CljsCore_testFirst{xs})
				})
			}(&cljs_core.AFn{})

			{
				var fv_1242 = (&CljsCore_testFirst{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})})
				var fs_1243 = (&CljsCore_testFirst{"asdf"})
				_, _ = fv_1242, fs_1243
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fv_1242.X_invoke_Arity0()
				}(), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (fv) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fs_1243.X_invoke_Arity0()
				}(), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB((`` + cljs_core.Str.X_invoke_Arity1(fs_1243).(string)), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (str fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_1242.X_get_first_Arity1(), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-get-first fv) 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fs_1243.X_get_first_Arity1(), "a") {
				} else {
					panic((&js.Error{("Assert failed: (= (-get-first fs) \\a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_1242.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= (-find-first fv [1]) 1)")}))
				}
				if reflect.DeepEqual(func() interface{} {
					var G__764 = float64(1)
					_ = G__764
					return fv_1242.X_invoke_Arity1(G__764)
				}(), fv_1242) {
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
				var t_1244 = (&CljsCore_testDestructuringWithLocals{float64(1)})
				_ = t_1244
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(1)}, nil}), t_1244.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{("Assert failed: (= [2 3 1] (-find-first t [2 3]))")}))
				}
			}
			{
				var x_1245 = float64(1)
				_ = x_1245
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__768 = x_1245
					_ = G__768
					switch G__768 {
					case float64(1):
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})

					default:
						panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_1245).(string))}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x 1 :one) :one)")}))
				}
			}
			{
				var x_1247 = float64(1)
				_ = x_1247
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__769 = x_1247
					_ = G__769
					switch G__769 {
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
				var x_1249 = float64(1)
				_ = x_1249
				if cljs_core.X_EQ_.Arity2IIB(func() (return__1250 interface{}) {
					defer func() {
						if e770 := recover(); e770 != nil {
							if cljs_core.Value_(e770).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e770
									_ = e
									return__1250 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e770)

							}
						}
					}()
					{
						{
							var G__771 = x_1249
							_ = G__771
							switch G__771 {
							case float64(3):
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "three", Fqn: "three", X_hash: float64(-1651831795)})

							default:
								panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_1249).(string))}))

							}
						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (try (case x 3 :three) (catch js/Error e :fail)) :fail)")}))
				}
			}
			{
				var x_1252 = float64(1)
				_ = x_1252
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__772 = x_1252
					_ = G__772
					switch G__772 {
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
				var x_1254 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})
				_ = x_1254
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__773 = x_1254
					_ = G__773
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), G__773) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
					} else {
						panic((&js.Error{("No matching clause: " + cljs_core.Str.X_invoke_Arity1(x_1254).(string))}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
				} else {
					panic((&js.Error{("Assert failed: (= (case x [:a :b] :ok) :ok)")}))
				}
			}
			{
				var a_1255 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = a_1255
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__774 = a_1255
					_ = G__774
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__774) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__774) {
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
				var a_1256 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil})
				_ = a_1256
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__775 = a_1256
					_ = G__775
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__775) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__775) {
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
				var foo_1257 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = foo_1257
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__776 = foo_1257
					_ = G__776
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__776) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__776) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), G__776) {
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
					var G__777 = foo_1257
					_ = G__777
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "d", Str: "d", X_hash: float64(-682293345), X_meta: nil}), G__777) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__777) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__777) {
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
				var r_1258 = cljs_core.Range_.X_invoke_Arity1(float64(64)).(*cljs_core.CljsCoreRange)
				var v_1259 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, r_1258)
				_, _ = r_1258, v_1259
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1259)), cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1259))) {
				} else {
					panic((&js.Error{("Assert failed: (= (hash (seq v)) (hash (seq v)))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Array_chunk.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)}).(*cljs_core.CljsCoreArrayChunk))) {
				} else {
					panic((&js.Error{("Assert failed: (= 6 (reduce + (array-chunk (array 1 2 3))))")}))
				}
				if cljs_core.Value_(cljs_core.Seq.Arity1IQ(v_1259)).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreChunkedSeq)(nil)).Elem()) {
				} else {
					panic((&js.Error{("Assert failed: (instance? ChunkedSeq (seq v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(r_1258, cljs_core.Seq.Arity1IQ(v_1259)) {
				} else {
					panic((&js.Error{("Assert failed: (= r (seq v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, r_1258).(*cljs_core.CljsCoreLazySeq), cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, v_1259).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (map inc r) (map inc v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, r_1258).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, v_1259).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (filter even? r) (filter even? v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, r_1258).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, v_1259).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (filter odd? r) (filter odd? v))")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic(r_1258, r_1258, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{r_1258})).(*cljs_core.CljsCoreLazySeq), cljs_core.Concat.X_invoke_ArityVariadic(v_1259, v_1259, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{v_1259})).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{("Assert failed: (= (concat r r r) (concat v v v))")}))
				}
				if cljs_core.DecoratedValue_(cljs_core.Seq.Arity1IQ(v_1259)).Type().Implements(reflect.TypeOf((*cljs_core.CljsCoreIReduce)(nil)).Elem()) {
				} else {
					panic((&js.Error{("Assert failed: (satisfies? IReduce (seq v))")}))
				}
				if float64(2010) == cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1259)))))).(float64) {
				} else {
					panic((&js.Error{("Assert failed: (== 2010 (reduce + (nnext (nnext (seq v)))))")}))
				}
				if float64(2020) == cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1259)))))).(float64) {
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
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__1260 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1260, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), nil).(*cljs_core.CljsCoreCons)
				})
			}(&cljs_core.AFn{}), nil, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= nil (next (lazy-seq (cons 1 nil))))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__1261 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1261, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), (&cljs_core.CljsCoreLazySeq{nil, func(G__1262 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__1262, 0, func() interface{} {
							return cljs_core.Cons.X_invoke_Arity2(float64(2), (&cljs_core.CljsCoreLazySeq{nil, func(G__1263 *cljs_core.AFn) *cljs_core.AFn {
								return cljs_core.Fn(G__1263, 0, func() interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1)))) {
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
					return cljs_core.Fn(map__GT_PrintMe, 1, func(G__780 interface{}) interface{} {
						return (&CljsCore_testPrintMe{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__780), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__780), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__780, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})})), nil})
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
				var seq__797_1264 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(13)).(*cljs_core.CljsCoreRange))
				var chunk__810_1265 interface{} = nil
				var count__811_1266 = float64(0)
				var i__812_1267 = float64(0)
				_, _, _, _ = seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267
				for {
					if i__812_1267 < count__811_1266 {
						{
							var month_1268 = cljs_core.Decorate_(chunk__810_1265).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__812_1267)
							_ = month_1268
							{
								var seq__813_1269 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
								var chunk__818_1270 interface{} = nil
								var count__819_1271 = float64(0)
								var i__820_1272 = float64(0)
								_, _, _, _ = seq__813_1269, chunk__818_1270, count__819_1271, i__820_1272
								for {
									if i__820_1272 < count__819_1271 {
										{
											var day_1273 = cljs_core.Decorate_(chunk__818_1270).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__820_1272)
											_ = day_1273
											{
												var seq__821_1274 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
												var chunk__822_1275 interface{} = nil
												var count__823_1276 = float64(0)
												var i__824_1277 = float64(0)
												_, _, _, _ = seq__821_1274, chunk__822_1275, count__823_1276, i__824_1277
												for {
													if i__824_1277 < count__823_1276 {
														{
															var hour_1278 = cljs_core.Decorate_(chunk__822_1275).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__824_1277)
															_ = hour_1278
															{
																var pad_1279 = func(G__1281 *cljs_core.AFn, seq__821_1274 interface{}, chunk__822_1275 interface{}, count__823_1276 float64, i__824_1277 float64, seq__813_1269 interface{}, chunk__818_1270 interface{}, count__819_1271 float64, i__820_1272 float64, seq__797_1264 interface{}, chunk__810_1265 interface{}, count__811_1266 float64, i__812_1267 float64, hour_1278 interface{}, day_1273 interface{}, month_1268 interface{}) *cljs_core.AFn {
																	return cljs_core.Fn(G__1281, 1, func(n interface{}) interface{} {
																		if n.(float64) < float64(10) {
																			return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																		} else {
																			return n
																		}
																	})
																}(&cljs_core.AFn{}, seq__821_1274, chunk__822_1275, count__823_1276, i__824_1277, seq__813_1269, chunk__818_1270, count__819_1271, i__820_1272, seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267, hour_1278, day_1273, month_1268)
																var inst_1280 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1279.X_invoke_Arity1(month_1268)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1279.X_invoke_Arity1(day_1273)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1279.X_invoke_Arity1(hour_1278)).(string) + ":14:15.666-00:00")
																_, _ = pad_1279, inst_1280
																if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1280})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1280).(string) + "\"")) {
																} else {
																	panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																}
															}
															seq__821_1274, chunk__822_1275, count__823_1276, i__824_1277 = seq__821_1274, chunk__822_1275, count__823_1276, (i__824_1277 + float64(1))
															continue
														}
													} else {
														{
															var temp__4388__auto___1282 = cljs_core.Seq.Arity1IQ(seq__821_1274)
															_ = temp__4388__auto___1282
															if cljs_core.Truth_(temp__4388__auto___1282) {
																{
																	var seq__821_1283___1 = temp__4388__auto___1282
																	_ = seq__821_1283___1
																	if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__821_1283___1) {
																		{
																			var c__971__auto___1284 = cljs_core.Chunk_first.X_invoke_Arity1(seq__821_1283___1)
																			_ = c__971__auto___1284
																			seq__821_1274, chunk__822_1275, count__823_1276, i__824_1277 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__821_1283___1), c__971__auto___1284, cljs_core.Count.X_invoke_Arity1(c__971__auto___1284).(float64), float64(0)
																			continue
																		}
																	} else {
																		{
																			var hour_1285 = cljs_core.First.X_invoke_Arity1(seq__821_1283___1)
																			_ = hour_1285
																			{
																				var pad_1286 = func(G__1288 *cljs_core.AFn, seq__821_1274 interface{}, chunk__822_1275 interface{}, count__823_1276 float64, i__824_1277 float64, seq__813_1269 interface{}, chunk__818_1270 interface{}, count__819_1271 float64, i__820_1272 float64, seq__797_1264 interface{}, chunk__810_1265 interface{}, count__811_1266 float64, i__812_1267 float64, hour_1285 interface{}, seq__821_1283___1 interface{}, temp__4388__auto___1282 cljs_core.CljsCoreISeq, day_1273 interface{}, month_1268 interface{}) *cljs_core.AFn {
																					return cljs_core.Fn(G__1288, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__821_1274, chunk__822_1275, count__823_1276, i__824_1277, seq__813_1269, chunk__818_1270, count__819_1271, i__820_1272, seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267, hour_1285, seq__821_1283___1, temp__4388__auto___1282, day_1273, month_1268)
																				var inst_1287 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1286.X_invoke_Arity1(month_1268)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1286.X_invoke_Arity1(day_1273)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1286.X_invoke_Arity1(hour_1285)).(string) + ":14:15.666-00:00")
																				_, _ = pad_1286, inst_1287
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1287})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1287).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__821_1274, chunk__822_1275, count__823_1276, i__824_1277 = cljs_core.Next.Arity1IQ(seq__821_1283___1), nil, float64(0), float64(0)
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
											seq__813_1269, chunk__818_1270, count__819_1271, i__820_1272 = seq__813_1269, chunk__818_1270, count__819_1271, (i__820_1272 + float64(1))
											continue
										}
									} else {
										{
											var temp__4388__auto___1289 = cljs_core.Seq.Arity1IQ(seq__813_1269)
											_ = temp__4388__auto___1289
											if cljs_core.Truth_(temp__4388__auto___1289) {
												{
													var seq__813_1290___1 = temp__4388__auto___1289
													_ = seq__813_1290___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__813_1290___1) {
														{
															var c__971__auto___1291 = cljs_core.Chunk_first.X_invoke_Arity1(seq__813_1290___1)
															_ = c__971__auto___1291
															seq__813_1269, chunk__818_1270, count__819_1271, i__820_1272 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__813_1290___1), c__971__auto___1291, cljs_core.Count.X_invoke_Arity1(c__971__auto___1291).(float64), float64(0)
															continue
														}
													} else {
														{
															var day_1292 = cljs_core.First.X_invoke_Arity1(seq__813_1290___1)
															_ = day_1292
															{
																var seq__814_1293 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																var chunk__815_1294 interface{} = nil
																var count__816_1295 = float64(0)
																var i__817_1296 = float64(0)
																_, _, _, _ = seq__814_1293, chunk__815_1294, count__816_1295, i__817_1296
																for {
																	if i__817_1296 < count__816_1295 {
																		{
																			var hour_1297 = cljs_core.Decorate_(chunk__815_1294).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__817_1296)
																			_ = hour_1297
																			{
																				var pad_1298 = func(G__1300 *cljs_core.AFn, seq__814_1293 interface{}, chunk__815_1294 interface{}, count__816_1295 float64, i__817_1296 float64, seq__813_1269 interface{}, chunk__818_1270 interface{}, count__819_1271 float64, i__820_1272 float64, seq__797_1264 interface{}, chunk__810_1265 interface{}, count__811_1266 float64, i__812_1267 float64, hour_1297 interface{}, day_1292 interface{}, seq__813_1290___1 interface{}, temp__4388__auto___1289 cljs_core.CljsCoreISeq, month_1268 interface{}) *cljs_core.AFn {
																					return cljs_core.Fn(G__1300, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__814_1293, chunk__815_1294, count__816_1295, i__817_1296, seq__813_1269, chunk__818_1270, count__819_1271, i__820_1272, seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267, hour_1297, day_1292, seq__813_1290___1, temp__4388__auto___1289, month_1268)
																				var inst_1299 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1298.X_invoke_Arity1(month_1268)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1298.X_invoke_Arity1(day_1292)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1298.X_invoke_Arity1(hour_1297)).(string) + ":14:15.666-00:00")
																				_, _ = pad_1298, inst_1299
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1299})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1299).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__814_1293, chunk__815_1294, count__816_1295, i__817_1296 = seq__814_1293, chunk__815_1294, count__816_1295, (i__817_1296 + float64(1))
																			continue
																		}
																	} else {
																		{
																			var temp__4388__auto___1301___1 = cljs_core.Seq.Arity1IQ(seq__814_1293)
																			_ = temp__4388__auto___1301___1
																			if cljs_core.Truth_(temp__4388__auto___1301___1) {
																				{
																					var seq__814_1302___1 = temp__4388__auto___1301___1
																					_ = seq__814_1302___1
																					if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__814_1302___1) {
																						{
																							var c__971__auto___1303 = cljs_core.Chunk_first.X_invoke_Arity1(seq__814_1302___1)
																							_ = c__971__auto___1303
																							seq__814_1293, chunk__815_1294, count__816_1295, i__817_1296 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__814_1302___1), c__971__auto___1303, cljs_core.Count.X_invoke_Arity1(c__971__auto___1303).(float64), float64(0)
																							continue
																						}
																					} else {
																						{
																							var hour_1304 = cljs_core.First.X_invoke_Arity1(seq__814_1302___1)
																							_ = hour_1304
																							{
																								var pad_1305 = func(G__1307 *cljs_core.AFn, seq__814_1293 interface{}, chunk__815_1294 interface{}, count__816_1295 float64, i__817_1296 float64, seq__813_1269 interface{}, chunk__818_1270 interface{}, count__819_1271 float64, i__820_1272 float64, seq__797_1264 interface{}, chunk__810_1265 interface{}, count__811_1266 float64, i__812_1267 float64, hour_1304 interface{}, seq__814_1302___1 interface{}, temp__4388__auto___1301___1 cljs_core.CljsCoreISeq, day_1292 interface{}, seq__813_1290___1 interface{}, temp__4388__auto___1289 cljs_core.CljsCoreISeq, month_1268 interface{}) *cljs_core.AFn {
																									return cljs_core.Fn(G__1307, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__814_1293, chunk__815_1294, count__816_1295, i__817_1296, seq__813_1269, chunk__818_1270, count__819_1271, i__820_1272, seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267, hour_1304, seq__814_1302___1, temp__4388__auto___1301___1, day_1292, seq__813_1290___1, temp__4388__auto___1289, month_1268)
																								var inst_1306 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1305.X_invoke_Arity1(month_1268)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1305.X_invoke_Arity1(day_1292)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1305.X_invoke_Arity1(hour_1304)).(string) + ":14:15.666-00:00")
																								_, _ = pad_1305, inst_1306
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1306})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1306).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__814_1293, chunk__815_1294, count__816_1295, i__817_1296 = cljs_core.Next.Arity1IQ(seq__814_1302___1), nil, float64(0), float64(0)
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
															seq__813_1269, chunk__818_1270, count__819_1271, i__820_1272 = cljs_core.Next.Arity1IQ(seq__813_1290___1), nil, float64(0), float64(0)
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
							seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267 = seq__797_1264, chunk__810_1265, count__811_1266, (i__812_1267 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1308 = cljs_core.Seq.Arity1IQ(seq__797_1264)
							_ = temp__4388__auto___1308
							if cljs_core.Truth_(temp__4388__auto___1308) {
								{
									var seq__797_1309___1 = temp__4388__auto___1308
									_ = seq__797_1309___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__797_1309___1) {
										{
											var c__971__auto___1310 = cljs_core.Chunk_first.X_invoke_Arity1(seq__797_1309___1)
											_ = c__971__auto___1310
											seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__797_1309___1), c__971__auto___1310, cljs_core.Count.X_invoke_Arity1(c__971__auto___1310).(float64), float64(0)
											continue
										}
									} else {
										{
											var month_1311 = cljs_core.First.X_invoke_Arity1(seq__797_1309___1)
											_ = month_1311
											{
												var seq__798_1312 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
												var chunk__803_1313 interface{} = nil
												var count__804_1314 = float64(0)
												var i__805_1315 = float64(0)
												_, _, _, _ = seq__798_1312, chunk__803_1313, count__804_1314, i__805_1315
												for {
													if i__805_1315 < count__804_1314 {
														{
															var day_1316 = cljs_core.Decorate_(chunk__803_1313).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__805_1315)
															_ = day_1316
															{
																var seq__806_1317 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																var chunk__807_1318 interface{} = nil
																var count__808_1319 = float64(0)
																var i__809_1320 = float64(0)
																_, _, _, _ = seq__806_1317, chunk__807_1318, count__808_1319, i__809_1320
																for {
																	if i__809_1320 < count__808_1319 {
																		{
																			var hour_1321 = cljs_core.Decorate_(chunk__807_1318).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__809_1320)
																			_ = hour_1321
																			{
																				var pad_1322 = func(G__1324 *cljs_core.AFn, seq__806_1317 interface{}, chunk__807_1318 interface{}, count__808_1319 float64, i__809_1320 float64, seq__798_1312 interface{}, chunk__803_1313 interface{}, count__804_1314 float64, i__805_1315 float64, seq__797_1264 interface{}, chunk__810_1265 interface{}, count__811_1266 float64, i__812_1267 float64, hour_1321 interface{}, day_1316 interface{}, month_1311 interface{}, seq__797_1309___1 interface{}, temp__4388__auto___1308 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																					return cljs_core.Fn(G__1324, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__806_1317, chunk__807_1318, count__808_1319, i__809_1320, seq__798_1312, chunk__803_1313, count__804_1314, i__805_1315, seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267, hour_1321, day_1316, month_1311, seq__797_1309___1, temp__4388__auto___1308)
																				var inst_1323 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1322.X_invoke_Arity1(month_1311)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1322.X_invoke_Arity1(day_1316)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1322.X_invoke_Arity1(hour_1321)).(string) + ":14:15.666-00:00")
																				_, _ = pad_1322, inst_1323
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1323})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1323).(string) + "\"")) {
																				} else {
																					panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																				}
																			}
																			seq__806_1317, chunk__807_1318, count__808_1319, i__809_1320 = seq__806_1317, chunk__807_1318, count__808_1319, (i__809_1320 + float64(1))
																			continue
																		}
																	} else {
																		{
																			var temp__4388__auto___1325___1 = cljs_core.Seq.Arity1IQ(seq__806_1317)
																			_ = temp__4388__auto___1325___1
																			if cljs_core.Truth_(temp__4388__auto___1325___1) {
																				{
																					var seq__806_1326___1 = temp__4388__auto___1325___1
																					_ = seq__806_1326___1
																					if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__806_1326___1) {
																						{
																							var c__971__auto___1327 = cljs_core.Chunk_first.X_invoke_Arity1(seq__806_1326___1)
																							_ = c__971__auto___1327
																							seq__806_1317, chunk__807_1318, count__808_1319, i__809_1320 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__806_1326___1), c__971__auto___1327, cljs_core.Count.X_invoke_Arity1(c__971__auto___1327).(float64), float64(0)
																							continue
																						}
																					} else {
																						{
																							var hour_1328 = cljs_core.First.X_invoke_Arity1(seq__806_1326___1)
																							_ = hour_1328
																							{
																								var pad_1329 = func(G__1331 *cljs_core.AFn, seq__806_1317 interface{}, chunk__807_1318 interface{}, count__808_1319 float64, i__809_1320 float64, seq__798_1312 interface{}, chunk__803_1313 interface{}, count__804_1314 float64, i__805_1315 float64, seq__797_1264 interface{}, chunk__810_1265 interface{}, count__811_1266 float64, i__812_1267 float64, hour_1328 interface{}, seq__806_1326___1 interface{}, temp__4388__auto___1325___1 cljs_core.CljsCoreISeq, day_1316 interface{}, month_1311 interface{}, seq__797_1309___1 interface{}, temp__4388__auto___1308 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																									return cljs_core.Fn(G__1331, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__806_1317, chunk__807_1318, count__808_1319, i__809_1320, seq__798_1312, chunk__803_1313, count__804_1314, i__805_1315, seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267, hour_1328, seq__806_1326___1, temp__4388__auto___1325___1, day_1316, month_1311, seq__797_1309___1, temp__4388__auto___1308)
																								var inst_1330 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1329.X_invoke_Arity1(month_1311)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1329.X_invoke_Arity1(day_1316)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1329.X_invoke_Arity1(hour_1328)).(string) + ":14:15.666-00:00")
																								_, _ = pad_1329, inst_1330
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1330})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1330).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__806_1317, chunk__807_1318, count__808_1319, i__809_1320 = cljs_core.Next.Arity1IQ(seq__806_1326___1), nil, float64(0), float64(0)
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
															seq__798_1312, chunk__803_1313, count__804_1314, i__805_1315 = seq__798_1312, chunk__803_1313, count__804_1314, (i__805_1315 + float64(1))
															continue
														}
													} else {
														{
															var temp__4388__auto___1332___1 = cljs_core.Seq.Arity1IQ(seq__798_1312)
															_ = temp__4388__auto___1332___1
															if cljs_core.Truth_(temp__4388__auto___1332___1) {
																{
																	var seq__798_1333___1 = temp__4388__auto___1332___1
																	_ = seq__798_1333___1
																	if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__798_1333___1) {
																		{
																			var c__971__auto___1334 = cljs_core.Chunk_first.X_invoke_Arity1(seq__798_1333___1)
																			_ = c__971__auto___1334
																			seq__798_1312, chunk__803_1313, count__804_1314, i__805_1315 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__798_1333___1), c__971__auto___1334, cljs_core.Count.X_invoke_Arity1(c__971__auto___1334).(float64), float64(0)
																			continue
																		}
																	} else {
																		{
																			var day_1335 = cljs_core.First.X_invoke_Arity1(seq__798_1333___1)
																			_ = day_1335
																			{
																				var seq__799_1336 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																				var chunk__800_1337 interface{} = nil
																				var count__801_1338 = float64(0)
																				var i__802_1339 = float64(0)
																				_, _, _, _ = seq__799_1336, chunk__800_1337, count__801_1338, i__802_1339
																				for {
																					if i__802_1339 < count__801_1338 {
																						{
																							var hour_1340 = cljs_core.Decorate_(chunk__800_1337).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__802_1339)
																							_ = hour_1340
																							{
																								var pad_1341 = func(G__1343 *cljs_core.AFn, seq__799_1336 interface{}, chunk__800_1337 interface{}, count__801_1338 float64, i__802_1339 float64, seq__798_1312 interface{}, chunk__803_1313 interface{}, count__804_1314 float64, i__805_1315 float64, seq__797_1264 interface{}, chunk__810_1265 interface{}, count__811_1266 float64, i__812_1267 float64, hour_1340 interface{}, day_1335 interface{}, seq__798_1333___1 interface{}, temp__4388__auto___1332___1 cljs_core.CljsCoreISeq, month_1311 interface{}, seq__797_1309___1 interface{}, temp__4388__auto___1308 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																									return cljs_core.Fn(G__1343, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__799_1336, chunk__800_1337, count__801_1338, i__802_1339, seq__798_1312, chunk__803_1313, count__804_1314, i__805_1315, seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267, hour_1340, day_1335, seq__798_1333___1, temp__4388__auto___1332___1, month_1311, seq__797_1309___1, temp__4388__auto___1308)
																								var inst_1342 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1341.X_invoke_Arity1(month_1311)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1341.X_invoke_Arity1(day_1335)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1341.X_invoke_Arity1(hour_1340)).(string) + ":14:15.666-00:00")
																								_, _ = pad_1341, inst_1342
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1342})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1342).(string) + "\"")) {
																								} else {
																									panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																								}
																							}
																							seq__799_1336, chunk__800_1337, count__801_1338, i__802_1339 = seq__799_1336, chunk__800_1337, count__801_1338, (i__802_1339 + float64(1))
																							continue
																						}
																					} else {
																						{
																							var temp__4388__auto___1344___2 = cljs_core.Seq.Arity1IQ(seq__799_1336)
																							_ = temp__4388__auto___1344___2
																							if cljs_core.Truth_(temp__4388__auto___1344___2) {
																								{
																									var seq__799_1345___1 = temp__4388__auto___1344___2
																									_ = seq__799_1345___1
																									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__799_1345___1) {
																										{
																											var c__971__auto___1346 = cljs_core.Chunk_first.X_invoke_Arity1(seq__799_1345___1)
																											_ = c__971__auto___1346
																											seq__799_1336, chunk__800_1337, count__801_1338, i__802_1339 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__799_1345___1), c__971__auto___1346, cljs_core.Count.X_invoke_Arity1(c__971__auto___1346).(float64), float64(0)
																											continue
																										}
																									} else {
																										{
																											var hour_1347 = cljs_core.First.X_invoke_Arity1(seq__799_1345___1)
																											_ = hour_1347
																											{
																												var pad_1348 = func(G__1350 *cljs_core.AFn, seq__799_1336 interface{}, chunk__800_1337 interface{}, count__801_1338 float64, i__802_1339 float64, seq__798_1312 interface{}, chunk__803_1313 interface{}, count__804_1314 float64, i__805_1315 float64, seq__797_1264 interface{}, chunk__810_1265 interface{}, count__811_1266 float64, i__812_1267 float64, hour_1347 interface{}, seq__799_1345___1 interface{}, temp__4388__auto___1344___2 cljs_core.CljsCoreISeq, day_1335 interface{}, seq__798_1333___1 interface{}, temp__4388__auto___1332___1 cljs_core.CljsCoreISeq, month_1311 interface{}, seq__797_1309___1 interface{}, temp__4388__auto___1308 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																													return cljs_core.Fn(G__1350, 1, func(n interface{}) interface{} {
																														if n.(float64) < float64(10) {
																															return ("0" + cljs_core.Str.X_invoke_Arity1(n).(string))
																														} else {
																															return n
																														}
																													})
																												}(&cljs_core.AFn{}, seq__799_1336, chunk__800_1337, count__801_1338, i__802_1339, seq__798_1312, chunk__803_1313, count__804_1314, i__805_1315, seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267, hour_1347, seq__799_1345___1, temp__4388__auto___1344___2, day_1335, seq__798_1333___1, temp__4388__auto___1332___1, month_1311, seq__797_1309___1, temp__4388__auto___1308)
																												var inst_1349 = ("2010-" + cljs_core.Str.X_invoke_Arity1(pad_1348.X_invoke_Arity1(month_1311)).(string) + "-" + cljs_core.Str.X_invoke_Arity1(pad_1348.X_invoke_Arity1(day_1335)).(string) + "T" + cljs_core.Str.X_invoke_Arity1(pad_1348.X_invoke_Arity1(hour_1347)).(string) + ":14:15.666-00:00")
																												_, _ = pad_1348, inst_1349
																												if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1349})})).(string), ("#inst \"" + cljs_core.Str.X_invoke_Arity1(inst_1349).(string) + "\"")) {
																												} else {
																													panic((&js.Error{("Assert failed: (= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))")}))
																												}
																											}
																											seq__799_1336, chunk__800_1337, count__801_1338, i__802_1339 = cljs_core.Next.Arity1IQ(seq__799_1345___1), nil, float64(0), float64(0)
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
																			seq__798_1312, chunk__803_1313, count__804_1314, i__805_1315 = cljs_core.Next.Arity1IQ(seq__798_1333___1), nil, float64(0), float64(0)
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
											seq__797_1264, chunk__810_1265, count__811_1266, i__812_1267 = cljs_core.Next.Arity1IQ(seq__797_1309___1), nil, float64(0), float64(0)
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
				var uuid_str_1351 = "550e8400-e29b-41d4-a716-446655440000"
				var uuid_1352 = (&cljs_core.CljsCoreUUID{uuid_str_1351})
				_, _ = uuid_str_1351, uuid_1352
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{uuid_1352})).(string), ("#uuid \"" + cljs_core.Str.X_invoke_Arity1(uuid_str_1351).(string) + "\"")) {
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
					return cljs_core.Decorate_(this).(CljsCore_testIBar).X_bar_Arity2(x)
				})
			}(&cljs_core.AFn{})

			Baz = func(baz *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(baz, 1, func(f interface{}) interface{} {
					X__GT_t830 = func(__GT_t830 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t830, 4, func(f___1 interface{}, baz___1 interface{}, test_stuff___1 interface{}, meta831 interface{}) interface{} {
							return (&CljsCore_testT830{f___1, baz___1, test_stuff___1, meta831})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT830{f, baz, test_stuff, nil})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(float64(2), Baz.X_invoke_Arity1(cljs_core.Inc).(*CljsCore_testT830).X_bar_Arity2(float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= 2 (-bar (baz inc) 1))")}))
			}
			{
				var x_1353 = "original"
				_ = x_1353
				Original_closure_stmt = func(original_closure_stmt *cljs_core.AFn, x_1353 string) *cljs_core.AFn {
					return cljs_core.Fn(original_closure_stmt, 0, func() interface{} {
						return x_1353
					})
				}(&cljs_core.AFn{}, x_1353)

			}
			{
				var x_1354 = "overwritten"
				_ = x_1354
				if cljs_core.X_EQ_.Arity2IIB("original", Original_closure_stmt.X_invoke_Arity0().(string)) {
				} else {
					panic((&js.Error{("Assert failed: (= \"original\" (original-closure-stmt))")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB("original", func() string {
				var x = "original"
				var oce = func(G__1355 *cljs_core.AFn, x string) *cljs_core.AFn {
					return cljs_core.Fn(G__1355, 0, func() interface{} {
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
					var x_1356___1 = func(G__1357 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__1357, 0, func() interface{} {
							return "overwritten"
						})
					}(&cljs_core.AFn{})
					_ = x_1356___1
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
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)}), cljs_core.Reduce_kv.X_invoke_Arity3(func(G__1358 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1358, 3, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
								return cljs_core.Reduced.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)})).(*cljs_core.CljsCoreReduced)
							})
						}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentVector_EMPTY, data)) {
						} else {
							panic((&js.Error{("Assert failed: (= :reduced (reduce-kv (fn [_ _ _] (reduced :reduced)) [] data))")}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sort.X_invoke_Arity1(expect), cljs_core.Sort.X_invoke_Arity1(cljs_core.Reduce_kv.X_invoke_Arity3(func(G__1359 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1359, 3, func(r interface{}, k interface{}, v interface{}) interface{} {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), func() (return__1360 interface{}) {
				defer func() {
					if e835 := recover(); e835 != nil {
						if cljs_core.Value_(e835).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreExceptionInfo)(nil)).Elem()) {
							{
								var e = e835
								_ = e
								return__1360 = cljs_core.Ex_data.X_invoke_Arity1(e)
							}
						} else {
							panic(e835)

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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(1), float64(2)}, nil}), func(G__1361 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1361, 0, func(more__ ...interface{}) interface{} {
					var more = cljs_core.Seq.Arity1IQ(more__[0])
					_ = more
					return more
				})
			}(&cljs_core.AFn{}).X_invoke_Arity3(js.Undefined, float64(1), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= [js/undefined 1 2] ((fn [& more] more) js/undefined 1 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(4), float64(5)}, nil}), func(G__1362 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1362, 2, func(a_b_more__ ...interface{}) interface{} {
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
				var fs_1363 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = fs_1363
				{
					var seq__836_1364 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity1(float64(4)).(*cljs_core.CljsCoreRange))
					var chunk__838_1365 interface{} = nil
					var count__839_1366 = float64(0)
					var i__840_1367 = float64(0)
					_, _, _, _ = seq__836_1364, chunk__838_1365, count__839_1366, i__840_1367
					for {
						if i__840_1367 < count__839_1366 {
							{
								var x_1368 = cljs_core.Decorate_(chunk__838_1365).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__840_1367)
								_ = x_1368
								{
									var y_1369 = (x_1368.(float64) + float64(1))
									var f_1370 = func(G__1371 *cljs_core.AFn, seq__836_1364 interface{}, chunk__838_1365 interface{}, count__839_1366 float64, i__840_1367 float64, y_1369 float64, x_1368 interface{}, fs_1363 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
										return cljs_core.Fn(G__1371, 0, func() interface{} {
											return y_1369
										})
									}(&cljs_core.AFn{}, seq__836_1364, chunk__838_1365, count__839_1366, i__840_1367, y_1369, x_1368, fs_1363)
									_, _ = y_1369, f_1370
									cljs_core.Swap_BANG_.X_invoke_Arity3(fs_1363, cljs_core.Conj, f_1370)
								}
								seq__836_1364, chunk__838_1365, count__839_1366, i__840_1367 = seq__836_1364, chunk__838_1365, count__839_1366, (i__840_1367 + float64(1))
								continue
							}
						} else {
							{
								var temp__4388__auto___1372 = cljs_core.Seq.Arity1IQ(seq__836_1364)
								_ = temp__4388__auto___1372
								if cljs_core.Truth_(temp__4388__auto___1372) {
									{
										var seq__836_1373___1 = temp__4388__auto___1372
										_ = seq__836_1373___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__836_1373___1) {
											{
												var c__971__auto___1374 = cljs_core.Chunk_first.X_invoke_Arity1(seq__836_1373___1)
												_ = c__971__auto___1374
												seq__836_1364, chunk__838_1365, count__839_1366, i__840_1367 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__836_1373___1), c__971__auto___1374, cljs_core.Count.X_invoke_Arity1(c__971__auto___1374).(float64), float64(0)
												continue
											}
										} else {
											{
												var x_1375 = cljs_core.First.X_invoke_Arity1(seq__836_1373___1)
												_ = x_1375
												{
													var y_1376 = (x_1375.(float64) + float64(1))
													var f_1377 = func(G__1378 *cljs_core.AFn, seq__836_1364 interface{}, chunk__838_1365 interface{}, count__839_1366 float64, i__840_1367 float64, y_1376 float64, x_1375 interface{}, seq__836_1373___1 interface{}, temp__4388__auto___1372 cljs_core.CljsCoreISeq, fs_1363 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
														return cljs_core.Fn(G__1378, 0, func() interface{} {
															return y_1376
														})
													}(&cljs_core.AFn{}, seq__836_1364, chunk__838_1365, count__839_1366, i__840_1367, y_1376, x_1375, seq__836_1373___1, temp__4388__auto___1372, fs_1363)
													_, _ = y_1376, f_1377
													cljs_core.Swap_BANG_.X_invoke_Arity3(fs_1363, cljs_core.Conj, f_1377)
												}
												seq__836_1364, chunk__838_1365, count__839_1366, i__840_1367 = cljs_core.Next.Arity1IQ(seq__836_1373___1), nil, float64(0), float64(0)
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1379 *cljs_core.AFn, fs_1363 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__1379, 1, func(p1__79_SHARP_ interface{}) interface{} {
						{
							return p1__79_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
						}
					})
				}(&cljs_core.AFn{}, fs_1363), cljs_core.Deref.X_invoke_Arity1(fs_1363)).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
				} else {
					panic((&js.Error{("Assert failed: (= (map (fn* [p1__79#] (p1__79#)) (clojure.core/deref fs)) (quote (1 2 3 4)))")}))
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCoreLazySeq{nil, func(G__1380 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1380, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCoreLazySeq{nil, func(G__1381 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1381, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreLazySeq{nil, func(G__1382 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1382, 0, func() interface{} {
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
					var sb__1141__auto__ = (&goog_string.StringBuffer{})
					_ = sb__1141__auto__
					{
						var _STAR_print_fn_STAR_842_1383 = cljs_core.X_STAR_print_fn_STAR_
						_ = _STAR_print_fn_STAR_842_1383
						func() {
							defer func() {
								cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_842_1383

							}()
							{
								cljs_core.X_STAR_print_fn_STAR_ = func(G__1384 *cljs_core.AFn, _STAR_print_fn_STAR_842_1383 interface{}, sb__1141__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
									return cljs_core.Fn(G__1384, 1, func(x__1142__auto__ interface{}) interface{} {
										return sb__1141__auto__.Append(x__1142__auto__)
									})
								}(&cljs_core.AFn{}, _STAR_print_fn_STAR_842_1383, sb__1141__auto__)

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
					return (`` + cljs_core.Str.X_invoke_Arity1(sb__1141__auto__).(string))
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

			Foo580 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__1385 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1385, 0, func() interface{} {
					return nil
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), func(G__1386 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1386, 0, func() interface{} {
					{
						var G__843 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
						_ = G__843
						return Foo580.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__843)
					}
				})
			}(&cljs_core.AFn{})}, nil})

			if cljs_core.Nil_((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(Foo580).(cljs_core.CljsCoreIFn).X_invoke_Arity0().(cljs_core.CljsCoreIFn).X_invoke_Arity0()) {
			} else {
				panic((&js.Error{("Assert failed: (nil? (((:b foo580))))")}))
			}
			if (cljs_core.First.X_invoke_Arity1(cljs_core.Filter.X_invoke_Arity2(func(G__1387 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1387, 1, func(p1__80_SHARP_ interface{}) interface{} {
					return (p1__80_SHARP_.(float64) == float64(9999))
				})
			}(&cljs_core.AFn{}), cljs_core.Range_.X_invoke_Arity0().(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)).(float64) == float64(9999)) {
			} else {
				panic((&js.Error{("Assert failed: (== (first (filter (fn* [p1__80#] (== p1__80# 9999)) (range))) 9999)")}))
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
				var a_1388 = func() *CljsCore_testT844 {
					X__GT_t844 = func(__GT_t844 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t844, 2, func(test_stuff___1 interface{}, meta845 interface{}) interface{} {
							return (&CljsCore_testT844{test_stuff___1, meta845})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT844{test_stuff, nil})
				}()
				var b_1389 = func() *CljsCore_testT847 {
					X__GT_t847 = func(__GT_t847 *cljs_core.AFn, a_1388 *CljsCore_testT844) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t847, 3, func(a___1 interface{}, test_stuff___1 interface{}, meta848 interface{}) interface{} {
							return (&CljsCore_testT847{a___1, test_stuff___1, meta848})
						})
					}(&cljs_core.AFn{}, a_1388)

					return (&CljsCore_testT847{a_1388, test_stuff, nil})
				}()
				var s_1390 = cljs_core.Set.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(128)).(*cljs_core.CljsCoreRange))
				_, _, _ = a_1388, b_1389, s_1390
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_1390, a_1388, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_1389}))), a_1388)), a_1388), cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_1390, a_1388, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_1389}))), a_1388)), a_1388)) {
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
			if cljs_core.Count.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(0)}, nil})}, true)).(float64) == cljs_core.Count.X_invoke_Arity1(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(0)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(0)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(0)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(4)}, nil}))).(float64) {
			} else {
				panic((&js.Error{("Assert failed: (== (count (hash-set [1 4] [2 4] [3 4] [0 3] [1 3] [2 3] [3 3] [0 2] [1 2] [2 2] [3 2] [4 2] [0 1] [1 1] [2 1] [3 1] [1 0] [2 0] [3 0])) (count (list [1 4] [2 4] [3 4] [0 3] [1 3] [2 3] [3 3] [0 2] [1 2] [2 2] [3 2] [4 2] [0 1] [1 1] [2 1] [3 1] [1 0] [2 0] [3 0])))")}))
			}
			X_woz = func(_woz *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(_woz, 1, func(this interface{}) interface{} {
					return cljs_core.Decorate_(this).(CljsCore_testIWoz).X_woz_Arity1()
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
				var seq__850_1391 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__851_1392 interface{} = nil
				var count__852_1393 = float64(0)
				var i__853_1394 = float64(0)
				_, _, _, _ = seq__850_1391, chunk__851_1392, count__852_1393, i__853_1394
				for {
					if i__853_1394 < count__852_1393 {
						{
							var n_1395 = cljs_core.Decorate_(chunk__851_1392).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__853_1394)
							_ = n_1395
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1396 interface{}) {
								defer func() {
									if e856 := recover(); e856 != nil {
										if cljs_core.Value_(e856).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e856
												_ = e
												return__1396 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e856)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1395, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1397 interface{}) {
								defer func() {
									if e857 := recover(); e857 != nil {
										if cljs_core.Value_(e857).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e857
												_ = e
												return__1397 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e857)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_1395, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))")}))
							}
							seq__850_1391, chunk__851_1392, count__852_1393, i__853_1394 = seq__850_1391, chunk__851_1392, count__852_1393, (i__853_1394 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1398 = cljs_core.Seq.Arity1IQ(seq__850_1391)
							_ = temp__4388__auto___1398
							if cljs_core.Truth_(temp__4388__auto___1398) {
								{
									var seq__850_1399___1 = temp__4388__auto___1398
									_ = seq__850_1399___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__850_1399___1) {
										{
											var c__971__auto___1400 = cljs_core.Chunk_first.X_invoke_Arity1(seq__850_1399___1)
											_ = c__971__auto___1400
											seq__850_1391, chunk__851_1392, count__852_1393, i__853_1394 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__850_1399___1), c__971__auto___1400, cljs_core.Count.X_invoke_Arity1(c__971__auto___1400).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1401 = cljs_core.First.X_invoke_Arity1(seq__850_1399___1)
											_ = n_1401
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1402 interface{}) {
												defer func() {
													if e858 := recover(); e858 != nil {
														if cljs_core.Value_(e858).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e858
																_ = e
																return__1402 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e858)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1401, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1403 interface{}) {
												defer func() {
													if e859 := recover(); e859 != nil {
														if cljs_core.Value_(e859).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e859
																_ = e
																return__1403 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e859)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_1401, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))")}))
											}
											seq__850_1391, chunk__851_1392, count__852_1393, i__853_1394 = cljs_core.Next.Arity1IQ(seq__850_1399___1), nil, float64(0), float64(0)
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
				var seq__860_1404 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__861_1405 interface{} = nil
				var count__862_1406 = float64(0)
				var i__863_1407 = float64(0)
				_, _, _, _ = seq__860_1404, chunk__861_1405, count__862_1406, i__863_1407
				for {
					if i__863_1407 < count__862_1406 {
						{
							var n_1408 = cljs_core.Decorate_(chunk__861_1405).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__863_1407)
							_ = n_1408
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1409 interface{}) {
								defer func() {
									if e866 := recover(); e866 != nil {
										if cljs_core.Value_(e866).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e866
												_ = e
												return__1409 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e866)

										}
									}
								}()
								{
									return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1408, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
							}
							seq__860_1404, chunk__861_1405, count__862_1406, i__863_1407 = seq__860_1404, chunk__861_1405, count__862_1406, (i__863_1407 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1410 = cljs_core.Seq.Arity1IQ(seq__860_1404)
							_ = temp__4388__auto___1410
							if cljs_core.Truth_(temp__4388__auto___1410) {
								{
									var seq__860_1411___1 = temp__4388__auto___1410
									_ = seq__860_1411___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__860_1411___1) {
										{
											var c__971__auto___1412 = cljs_core.Chunk_first.X_invoke_Arity1(seq__860_1411___1)
											_ = c__971__auto___1412
											seq__860_1404, chunk__861_1405, count__862_1406, i__863_1407 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__860_1411___1), c__971__auto___1412, cljs_core.Count.X_invoke_Arity1(c__971__auto___1412).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1413 = cljs_core.First.X_invoke_Arity1(seq__860_1411___1)
											_ = n_1413
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1414 interface{}) {
												defer func() {
													if e867 := recover(); e867 != nil {
														if cljs_core.Value_(e867).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e867
																_ = e
																return__1414 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e867)

														}
													}
												}()
												{
													return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1413, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
											}
											seq__860_1404, chunk__861_1405, count__862_1406, i__863_1407 = cljs_core.Next.Arity1IQ(seq__860_1411___1), nil, float64(0), float64(0)
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
				var map__868_1415 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__868_1416___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__868_1415) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__868_1415)
					} else {
						return map__868_1415
					}
				}()
				var b_1417 = cljs_core.Get.X_invoke_Arity2(map__868_1416___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a_1418 = cljs_core.Get.X_invoke_Arity2(map__868_1416___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__868_1415, map__868_1416___1, b_1417, a_1418
				if cljs_core.X_EQ_.Arity2IIB(float64(1), a_1418) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 a)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), b_1417) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 b)")}))
				}
			}
			{
				var map__869_1419 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__869_1420___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__869_1419) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__869_1419)
					} else {
						return map__869_1419
					}
				}()
				var d_1421 = cljs_core.Get.X_invoke_Arity2(map__869_1420___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_1422 = cljs_core.Get.X_invoke_Arity2(map__869_1420___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__869_1419, map__869_1420___1, d_1421, b_1422
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1422) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1421) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__870_1423 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__870_1424___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__870_1423) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__870_1423)
					} else {
						return map__870_1423
					}
				}()
				var d_1425 = cljs_core.Get.X_invoke_Arity2(map__870_1424___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_1426 = cljs_core.Get.X_invoke_Arity2(map__870_1424___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__870_1423, map__870_1424___1, d_1425, b_1426
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1426) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1425) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__871_1427 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}), float64(1), (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}), float64(2)}, nil})
				var map__871_1428___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__871_1427) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__871_1427)
					} else {
						return map__871_1427
					}
				}()
				var d_1429 = cljs_core.Get.X_invoke_Arity2(map__871_1428___1, (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}))
				var b_1430 = cljs_core.Get.X_invoke_Arity2(map__871_1428___1, (&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}))
				_, _, _, _ = map__871_1427, map__871_1428___1, d_1429, b_1430
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1430) {
				} else {
					panic((&js.Error{("Assert failed: (= 1 b)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1429) {
				} else {
					panic((&js.Error{("Assert failed: (= 2 d)")}))
				}
			}
			{
				var map__872_1431 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}), float64(2)}, nil})
				var map__872_1432___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__872_1431) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__872_1431)
					} else {
						return map__872_1431
					}
				}()
				var y_1433 = cljs_core.Get.X_invoke_Arity2(map__872_1432___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}))
				var x_1434 = cljs_core.Get.X_invoke_Arity2(map__872_1432___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}))
				_, _, _, _ = map__872_1431, map__872_1432___1, y_1433, x_1434
				if cljs_core.X_EQ_.Arity2IIB(x_1434, float64(1)) {
				} else {
					panic((&js.Error{("Assert failed: (= x 1)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(y_1433, float64(2)) {
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
								arr, names = cljs_core.Conj.X_invoke_Arity2(arr, func(G__1435 *cljs_core.AFn, arr interface{}, names interface{}, name interface{}) *cljs_core.AFn {
									return cljs_core.Fn(G__1435, 0, func() interface{} {
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
				var sb__1141__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1141__auto__
				{
					var _STAR_print_fn_STAR_873_1436 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_873_1436
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_873_1436

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__1437 *cljs_core.AFn, _STAR_print_fn_STAR_873_1436 interface{}, sb__1141__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__1437, 1, func(x__1142__auto__ interface{}) interface{} {
									return sb__1141__auto__.Append(x__1142__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_873_1436, sb__1141__auto__)

							{
								var seq__874_1438 interface{} = cljs_core.Seq.Arity1IQ(Cljs_739.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil})))
								var chunk__875_1439 interface{} = nil
								var count__876_1440 = float64(0)
								var i__877_1441 = float64(0)
								_, _, _, _ = seq__874_1438, chunk__875_1439, count__876_1440, i__877_1441
								for {
									if i__877_1441 < count__876_1440 {
										{
											var fn_1442 = cljs_core.Decorate_(chunk__875_1439).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__877_1441)
											_ = fn_1442
											{
												fn_1442.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
											}
											seq__874_1438, chunk__875_1439, count__876_1440, i__877_1441 = seq__874_1438, chunk__875_1439, count__876_1440, (i__877_1441 + float64(1))
											continue
										}
									} else {
										{
											var temp__4388__auto___1443 = cljs_core.Seq.Arity1IQ(seq__874_1438)
											_ = temp__4388__auto___1443
											if cljs_core.Truth_(temp__4388__auto___1443) {
												{
													var seq__874_1444___1 = temp__4388__auto___1443
													_ = seq__874_1444___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__874_1444___1) {
														{
															var c__971__auto___1445 = cljs_core.Chunk_first.X_invoke_Arity1(seq__874_1444___1)
															_ = c__971__auto___1445
															seq__874_1438, chunk__875_1439, count__876_1440, i__877_1441 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__874_1444___1), c__971__auto___1445, cljs_core.Count.X_invoke_Arity1(c__971__auto___1445).(float64), float64(0)
															continue
														}
													} else {
														{
															var fn_1446 = cljs_core.First.X_invoke_Arity1(seq__874_1444___1)
															_ = fn_1446
															{
																fn_1446.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
															}
															seq__874_1438, chunk__875_1439, count__876_1440, i__877_1441 = cljs_core.Next.Arity1IQ(seq__874_1444___1), nil, float64(0), float64(0)
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
				return (`` + cljs_core.Str.X_invoke_Arity1(sb__1141__auto__).(string))
			}(), ":a\n:b\n:c\n:d\n") {
			} else {
				panic((&js.Error{("Assert failed: (= (with-out-str (doseq [fn (cljs-739 [] [:a :b :c :d])] (fn))) \":a\\n:b\\n:c\\n:d\\n\")")}))
			}
			{
				var seq__878_1447 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__879_1448 interface{} = nil
				var count__880_1449 = float64(0)
				var i__881_1450 = float64(0)
				_, _, _, _ = seq__878_1447, chunk__879_1448, count__880_1449, i__881_1450
				for {
					if i__881_1450 < count__880_1449 {
						{
							var n_1451 = cljs_core.Decorate_(chunk__879_1448).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__881_1450)
							_ = n_1451
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1451)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get [1 2] n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1452 interface{}) {
								defer func() {
									if e884 := recover(); e884 != nil {
										if cljs_core.Value_(e884).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e884
												_ = e
												return__1452 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e884)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1451)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1451, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get [1 2] n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1453 interface{}) {
								defer func() {
									if e885 := recover(); e885 != nil {
										if cljs_core.Value_(e885).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e885
												_ = e
												return__1453 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e885)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1451, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1451)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (subvec [1 2] 1) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1454 interface{}) {
								defer func() {
									if e886 := recover(); e886 != nil {
										if cljs_core.Value_(e886).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e886
												_ = e
												return__1454 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e886)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1451)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1451, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (subvec [1 2] 1) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1455 interface{}) {
								defer func() {
									if e887 := recover(); e887 != nil {
										if cljs_core.Value_(e887).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e887
												_ = e
												return__1455 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e887)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1451, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1451)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (transient [1 2]) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1456 interface{}) {
								defer func() {
									if e888 := recover(); e888 != nil {
										if cljs_core.Value_(e888).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e888
												_ = e
												return__1456 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e888)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1451)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1451, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (transient [1 2]) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1457 interface{}) {
								defer func() {
									if e889 := recover(); e889 != nil {
										if cljs_core.Value_(e889).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e889
												_ = e
												return__1457 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e889)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1451, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1451)) {
							} else {
								panic((&js.Error{("Assert failed: (nil? (get (range 1 3) n))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1458 interface{}) {
								defer func() {
									if e890 := recover(); e890 != nil {
										if cljs_core.Value_(e890).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e890
												_ = e
												return__1458 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e890)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1451)
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1451, float64(4))) {
							} else {
								panic((&js.Error{("Assert failed: (= 4 (get (range 1 3) n 4))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1459 interface{}) {
								defer func() {
									if e891 := recover(); e891 != nil {
										if cljs_core.Value_(e891).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e891
												_ = e
												return__1459 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e891)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1451, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))")}))
							}
							seq__878_1447, chunk__879_1448, count__880_1449, i__881_1450 = seq__878_1447, chunk__879_1448, count__880_1449, (i__881_1450 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1460 = cljs_core.Seq.Arity1IQ(seq__878_1447)
							_ = temp__4388__auto___1460
							if cljs_core.Truth_(temp__4388__auto___1460) {
								{
									var seq__878_1461___1 = temp__4388__auto___1460
									_ = seq__878_1461___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__878_1461___1) {
										{
											var c__971__auto___1462 = cljs_core.Chunk_first.X_invoke_Arity1(seq__878_1461___1)
											_ = c__971__auto___1462
											seq__878_1447, chunk__879_1448, count__880_1449, i__881_1450 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__878_1461___1), c__971__auto___1462, cljs_core.Count.X_invoke_Arity1(c__971__auto___1462).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1463 = cljs_core.First.X_invoke_Arity1(seq__878_1461___1)
											_ = n_1463
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1463)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get [1 2] n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1464 interface{}) {
												defer func() {
													if e892 := recover(); e892 != nil {
														if cljs_core.Value_(e892).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e892
																_ = e
																return__1464 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e892)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1463)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1463, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get [1 2] n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1465 interface{}) {
												defer func() {
													if e893 := recover(); e893 != nil {
														if cljs_core.Value_(e893).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e893
																_ = e
																return__1465 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e893)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1463, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1463)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (subvec [1 2] 1) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1466 interface{}) {
												defer func() {
													if e894 := recover(); e894 != nil {
														if cljs_core.Value_(e894).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e894
																_ = e
																return__1466 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e894)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1463)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1463, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (subvec [1 2] 1) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1467 interface{}) {
												defer func() {
													if e895 := recover(); e895 != nil {
														if cljs_core.Value_(e895).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e895
																_ = e
																return__1467 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e895)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1463, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1463)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (transient [1 2]) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1468 interface{}) {
												defer func() {
													if e896 := recover(); e896 != nil {
														if cljs_core.Value_(e896).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e896
																_ = e
																return__1468 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e896)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1463)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1463, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (transient [1 2]) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1469 interface{}) {
												defer func() {
													if e897 := recover(); e897 != nil {
														if cljs_core.Value_(e897).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e897
																_ = e
																return__1469 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e897)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1463, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))")}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1463)) {
											} else {
												panic((&js.Error{("Assert failed: (nil? (get (range 1 3) n))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1470 interface{}) {
												defer func() {
													if e898 := recover(); e898 != nil {
														if cljs_core.Value_(e898).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e898
																_ = e
																return__1470 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e898)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1463)
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1463, float64(4))) {
											} else {
												panic((&js.Error{("Assert failed: (= 4 (get (range 1 3) n 4))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1471 interface{}) {
												defer func() {
													if e899 := recover(); e899 != nil {
														if cljs_core.Value_(e899).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e899
																_ = e
																return__1471 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e899)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1463, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))")}))
											}
											seq__878_1447, chunk__879_1448, count__880_1449, i__881_1450 = cljs_core.Next.Arity1IQ(seq__878_1461___1), nil, float64(0), float64(0)
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
			if cljs_core.Nil_(cljs_core.Decorate_(cljs_core.Rseq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil}))).(cljs_core.CljsCoreINext).X_next_Arity1()) {
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
				var x_1472 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}).X_invoke_Arity1(cljs_core.Meta.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}).X_invoke_Arity1(cljs_core.Deref.X_invoke_Arity1(Cljs_780))))
				_ = x_1472
				if cljs_core.Vector_QMARK_.Arity1IB(x_1472) {
				} else {
					panic((&js.Error{("Assert failed: (vector? x)")}))
				}
				if cljs_core.X_EQ_.Arity2IIB(x_1472, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})) {
				} else {
					panic((&js.Error{("Assert failed: (= x [1 2 3])")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3((&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}), "ToString", []interface{}{}), "550e8400-e29b-41d4-a716-446655440000") {
			} else {
				panic((&js.Error{("Assert failed: (= (.toString #uuid \"550e8400-e29b-41d4-a716-446655440000\") \"550e8400-e29b-41d4-a716-446655440000\")")}))
			}
			{
				var seq__900_1473 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__901_1474 interface{} = nil
				var count__902_1475 = float64(0)
				var i__903_1476 = float64(0)
				_, _, _, _ = seq__900_1473, chunk__901_1474, count__902_1475, i__903_1476
				for {
					if i__903_1476 < count__902_1475 {
						{
							var m_1477 = cljs_core.Decorate_(chunk__901_1474).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__903_1476)
							_ = m_1477
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__1478 interface{}) {
								defer func() {
									if e904 := recover(); e904 != nil {
										if cljs_core.Value_(e904).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var ___ = e904
												_ = ___
												return__1478 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
											}
										} else {
											panic(e904)

										}
									}
								}()
								{
									return cljs_core.Conj.X_invoke_Arity2(m_1477, "foo")
								}
							}()) {
							} else {
								panic((&js.Error{("Assert failed: (= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1477, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m [:foo 1]))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1477, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m {:foo 1}))")}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1477, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1} (conj m (list [:foo 1])))")}))
							}
							seq__900_1473, chunk__901_1474, count__902_1475, i__903_1476 = seq__900_1473, chunk__901_1474, count__902_1475, (i__903_1476 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1479 = cljs_core.Seq.Arity1IQ(seq__900_1473)
							_ = temp__4388__auto___1479
							if cljs_core.Truth_(temp__4388__auto___1479) {
								{
									var seq__900_1480___1 = temp__4388__auto___1479
									_ = seq__900_1480___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__900_1480___1) {
										{
											var c__971__auto___1481 = cljs_core.Chunk_first.X_invoke_Arity1(seq__900_1480___1)
											_ = c__971__auto___1481
											seq__900_1473, chunk__901_1474, count__902_1475, i__903_1476 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__900_1480___1), c__971__auto___1481, cljs_core.Count.X_invoke_Arity1(c__971__auto___1481).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_1482 = cljs_core.First.X_invoke_Arity1(seq__900_1480___1)
											_ = m_1482
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__1483 interface{}) {
												defer func() {
													if e905 := recover(); e905 != nil {
														if cljs_core.Value_(e905).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var ___ = e905
																_ = ___
																return__1483 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
															}
														} else {
															panic(e905)

														}
													}
												}()
												{
													return cljs_core.Conj.X_invoke_Arity2(m_1482, "foo")
												}
											}()) {
											} else {
												panic((&js.Error{("Assert failed: (= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1482, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m [:foo 1]))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1482, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m {:foo 1}))")}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1482, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1} (conj m (list [:foo 1])))")}))
											}
											seq__900_1473, chunk__901_1474, count__902_1475, i__903_1476 = cljs_core.Next.Arity1IQ(seq__900_1480___1), nil, float64(0), float64(0)
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
				var seq__906_1484 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Array_map, cljs_core.Hash_map, cljs_core.Sorted_map}, nil}))
				var chunk__907_1485 interface{} = nil
				var count__908_1486 = float64(0)
				var i__909_1487 = float64(0)
				_, _, _, _ = seq__906_1484, chunk__907_1485, count__908_1486, i__909_1487
				for {
					if i__909_1487 < count__908_1486 {
						{
							var mt_1488 = cljs_core.Decorate_(chunk__907_1485).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__909_1487)
							_ = mt_1488
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
								var G__910 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
								var G__911 = float64(1)
								_, _ = G__910, G__911
								return mt_1488.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__910, G__911)
							}(), func(make_seq *cljs_core.AFn, seq__906_1484 interface{}, chunk__907_1485 interface{}, count__908_1486 float64, i__909_1487 float64, mt_1488 interface{}) *cljs_core.AFn {
								return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
									if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
										X__GT_t917 = func(__GT_t917 *cljs_core.AFn, seq__906_1484 interface{}, chunk__907_1485 interface{}, count__908_1486 float64, i__909_1487 float64, mt_1488 interface{}) *cljs_core.AFn {
											return cljs_core.Fn(__GT_t917, 9, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, i__909___1 interface{}, count__908___1 interface{}, chunk__907___1 interface{}, seq__906___1 interface{}, test_stuff___1 interface{}, meta918 interface{}) interface{} {
												return (&CljsCore_testT917{from_seq___1, make_seq___1, mt___1, i__909___1, count__908___1, chunk__907___1, seq__906___1, test_stuff___1, meta918})
											})
										}(&cljs_core.AFn{}, seq__906_1484, chunk__907_1485, count__908_1486, i__909_1487, mt_1488)

										return (&CljsCore_testT917{from_seq, make_seq, mt_1488, i__909_1487, count__908_1486, chunk__907_1485, seq__906_1484, test_stuff, nil})
									} else {
										return nil
									}
								})
							}(&cljs_core.AFn{}, seq__906_1484, chunk__907_1485, count__908_1486, i__909_1487, mt_1488).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
							} else {
								panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))")}))
							}
							seq__906_1484, chunk__907_1485, count__908_1486, i__909_1487 = seq__906_1484, chunk__907_1485, count__908_1486, (i__909_1487 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1489 = cljs_core.Seq.Arity1IQ(seq__906_1484)
							_ = temp__4388__auto___1489
							if cljs_core.Truth_(temp__4388__auto___1489) {
								{
									var seq__906_1490___1 = temp__4388__auto___1489
									_ = seq__906_1490___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__906_1490___1) {
										{
											var c__971__auto___1491 = cljs_core.Chunk_first.X_invoke_Arity1(seq__906_1490___1)
											_ = c__971__auto___1491
											seq__906_1484, chunk__907_1485, count__908_1486, i__909_1487 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__906_1490___1), c__971__auto___1491, cljs_core.Count.X_invoke_Arity1(c__971__auto___1491).(float64), float64(0)
											continue
										}
									} else {
										{
											var mt_1492 = cljs_core.First.X_invoke_Arity1(seq__906_1490___1)
											_ = mt_1492
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
												var G__922 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
												var G__923 = float64(1)
												_, _ = G__922, G__923
												return mt_1492.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__922, G__923)
											}(), func(make_seq *cljs_core.AFn, seq__906_1484 interface{}, chunk__907_1485 interface{}, count__908_1486 float64, i__909_1487 float64, mt_1492 interface{}, seq__906_1490___1 interface{}, temp__4388__auto___1489 cljs_core.CljsCoreISeq) *cljs_core.AFn {
												return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
													if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
														X__GT_t929 = func(__GT_t929 *cljs_core.AFn, seq__906_1484 interface{}, chunk__907_1485 interface{}, count__908_1486 float64, i__909_1487 float64, mt_1492 interface{}, seq__906_1490___1 interface{}, temp__4388__auto___1489 cljs_core.CljsCoreISeq) *cljs_core.AFn {
															return cljs_core.Fn(__GT_t929, 10, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, temp__4388__auto_____1 interface{}, i__909___1 interface{}, count__908___1 interface{}, chunk__907___1 interface{}, seq__906___2 interface{}, test_stuff___1 interface{}, meta930 interface{}) interface{} {
																return (&CljsCore_testT929{from_seq___1, make_seq___1, mt___1, temp__4388__auto_____1, i__909___1, count__908___1, chunk__907___1, seq__906___2, test_stuff___1, meta930})
															})
														}(&cljs_core.AFn{}, seq__906_1484, chunk__907_1485, count__908_1486, i__909_1487, mt_1492, seq__906_1490___1, temp__4388__auto___1489)

														return (&CljsCore_testT929{from_seq, make_seq, mt_1492, temp__4388__auto___1489, i__909_1487, count__908_1486, chunk__907_1485, seq__906_1490___1, test_stuff, nil})
													} else {
														return nil
													}
												})
											}(&cljs_core.AFn{}, seq__906_1484, chunk__907_1485, count__908_1486, i__909_1487, mt_1492, seq__906_1490___1, temp__4388__auto___1489).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
											} else {
												panic((&js.Error{("Assert failed: (= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))")}))
											}
											seq__906_1484, chunk__907_1485, count__908_1486, i__909_1487 = cljs_core.Next.Arity1IQ(seq__906_1490___1), nil, float64(0), float64(0)
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
				var _STAR_print_length_STAR_934 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_934
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_934

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
				var _STAR_print_length_STAR_935 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_935
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_935

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
				var _STAR_print_length_STAR_936 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_936
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_936

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
				var _STAR_print_length_STAR_937 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_937
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_937

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
				var _STAR_print_length_STAR_938 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_938
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_938

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
				var _STAR_print_length_STAR_939 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_939
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_939

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
					var G__940 = func() interface{} {
						if cljs_core.Value_(x).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
							return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(x), "Fqn")
						} else {
							return nil
						}
					}()
					_ = G__940
					switch G__940 {
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
							var G__942 = func() interface{} {
								if cljs_core.Value_(value).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
									return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(value), "Fqn")
								} else {
									return nil
								}
							}()
							_ = G__942
							switch G__942 {
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
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}), cljs_core.Keys.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(cljs_core.Array_map, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2)}, nil})))) {
			} else {
				panic((&js.Error{("Assert failed: (= [:foo] (keys (apply array-map [:foo 1 :foo 2])))")}))
			}
			{
				var not_strings_1495 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, nil, float64(1), func(G__1496 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1496, 0, func() interface{} {
						return nil
					})
				}(&cljs_core.AFn{})}, nil})
				_ = not_strings_1495
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1497 *cljs_core.AFn, not_strings_1495 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1497, 1, func(p1__81_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1498 interface{}) {
							defer func() {
								if e943 := recover(); e943 != nil {
									if cljs_core.Value_(e943).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e943
											_ = ___
											return__1498 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e943)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__81_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1495), not_strings_1495) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__81#] (= :failed (try (re-find #\".\" p1__81#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1499 *cljs_core.AFn, not_strings_1495 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1499, 1, func(p1__82_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1500 interface{}) {
							defer func() {
								if e944 := recover(); e944 != nil {
									if cljs_core.Value_(e944).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e944
											_ = ___
											return__1500 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e944)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__82_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1495), not_strings_1495) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__82#] (= :failed (try (re-matches #\".\" p1__82#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1501 *cljs_core.AFn, not_strings_1495 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1501, 1, func(p1__83_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1502 interface{}) {
							defer func() {
								if e945 := recover(); e945 != nil {
									if cljs_core.Value_(e945).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e945
											_ = ___
											return__1502 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e945)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__83_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1495), not_strings_1495) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__83#] (= :failed (try (re-find #\"nomatch\" p1__83#) (catch js/TypeError _ :failed)))) not-strings)")}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1503 *cljs_core.AFn, not_strings_1495 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1503, 1, func(p1__84_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1504 interface{}) {
							defer func() {
								if e946 := recover(); e946 != nil {
									if cljs_core.Value_(e946).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e946
											_ = ___
											return__1504 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e946)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__84_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1495), not_strings_1495) {
				} else {
					panic((&js.Error{("Assert failed: (every? (fn* [p1__84#] (= :failed (try (re-matches #\"nomatch\" p1__84#) (catch js/TypeError _ :failed)))) not-strings)")}))
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(func(G__1505 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1505, 1, func(p1__85_SHARP_ interface{}) interface{} {
					return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(p1__85_SHARP_, "ToUpperCase", []interface{}{})
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), "foo")), "FOO") {
			} else {
				panic((&js.Error{("Assert failed: (= (apply str (sequence (map (fn* [p1__85#] (.toUpperCase p1__85#))) \"foo\")) \"FOO\")")}))
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Take_while.X_invoke_Arity1(func(G__1506 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1506, 1, func(p1__86_SHARP_ interface{}) interface{} {
					return (p1__86_SHARP_.(float64) < float64(5))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (take-while (fn* [p1__86#] (< p1__86# 5))) (range 10)) (quote (0 1 2 3 4)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Drop.X_invoke_Arity1(float64(5)).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (drop 5) (range 10)) (quote (5 6 7 8 9)))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Drop_while.X_invoke_Arity1(func(G__1507 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1507, 1, func(p1__87_SHARP_ interface{}) interface{} {
					return (p1__87_SHARP_.(float64) < float64(5))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence (drop-while (fn* [p1__87#] (< p1__87# 5))) (range 10)) (quote (5 6 7 8 9)))")}))
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
				var ret_1508 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))
				_ = ret_1508
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_1508)) && (cljs_core.X_EQ_.Arity2IIB(ret_1508, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList))) {
				} else {
					panic((&js.Error{("Assert failed: (and (vector? ret) (= ret (quote (1 2 3))))")}))
				}
			}
			{
				var ret_1509 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				_ = ret_1509
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_1509)) && (cljs_core.X_EQ_.Arity2IIB(ret_1509, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(2), float64(4), float64(6), float64(8)})).(*cljs_core.CljsCoreList))) {
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
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(cljs_core.Eduction.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreEduction)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (seq (eduction (map inc) [1 2 3])) (quote (2 3 4)))")}))
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
			Xform = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Dedupe.X_invoke_Arity0().(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Partition_all.X_invoke_Arity1(float64(3)).(cljs_core.CljsCoreIFn), cljs_core.Partition_by.X_invoke_Arity1(func(G__1510 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1510, 1, func(p1__88_SHARP_ interface{}) interface{} {
					return (cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, p1__88_SHARP_).(float64) < float64(7))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Flatten).(cljs_core.CljsCoreIFn), cljs_core.Random_sample.X_invoke_Arity1(1.0).(cljs_core.CljsCoreIFn), cljs_core.Take_nth.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Keep.X_invoke_Arity1(func(G__1511 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1511, 1, func(p1__89_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__89_SHARP_) {
						return (p1__89_SHARP_.(float64) * p1__89_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Keep_indexed.X_invoke_Arity1(func(G__1512 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1512, 2, func(p1__90_SHARP_ interface{}, p2__91_SHARP_ interface{}) interface{} {
					if cljs_core.Even_QMARK_.Arity1IB(p1__90_SHARP_) {
						return (p1__90_SHARP_.(float64) * p2__91_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Replace.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(2), "two", float64(6), "six", float64(18), "eighteen"}, nil})).(cljs_core.CljsCoreIFn), cljs_core.Take.X_invoke_Arity1(float64(11)).(cljs_core.CljsCoreIFn), cljs_core.Take_while.X_invoke_Arity1(func(G__1513 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1513, 1, func(p1__92_SHARP_ interface{}) interface{} {
					return cljs_core.Not_EQ_.Arity2IIB(float64(300), p1__92_SHARP_)
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Drop.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Drop_while.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Remove.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn)})).(cljs_core.CljsCoreIFn).(*cljs_core.AFn)

			Data = cljs_core.Vec.X_invoke_Arity1(cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(18)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(20)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(Xform, Data), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(36), float64(200), float64(10)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence xform data) (quote (36 200 10)))")}))
			}
			Xf = cljs_core.Map_.X_invoke_Arity1(func(G__1514 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1514, 2, func(p1__93_SHARP_ interface{}, p2__94_SHARP_ interface{}) interface{} {
					return (p1__93_SHARP_.(float64) + p2__94_SHARP_.(float64))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn).(*cljs_core.AFn)

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_ArityVariadic(Xf, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0)}, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})) {
			} else {
				panic((&js.Error{("Assert failed: (= (sequence xf [0 0] [1 2]) [1 2])")}))
			}
			{
				var xs_1515 = (&cljs_core.CljsCorePersistentVector{nil, float64(21), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(44), float64(43), float64(42), float64(41), float64(40), float64(39), float64(38), float64(37), float64(36), float64(35), float64(34), float64(33), float64(32), float64(31), float64(30), float64(29), float64(28), float64(27), float64(26), float64(25), float64(24)}, nil})
				_ = xs_1515
				{
					var m_1516 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.Zipmap.X_invoke_Arity2(xs_1515, cljs_core.Repeat.X_invoke_Arity1(float64(1)).(*cljs_core.CljsCoreLazySeq)))
					var xs_1517___1 interface{} = xs_1515
					_, _ = m_1516, xs_1517___1
					for {
						{
							var temp__4386__auto___1518 = cljs_core.First.X_invoke_Arity1(xs_1517___1)
							_ = temp__4386__auto___1518
							if cljs_core.Truth_(temp__4386__auto___1518) {
								{
									var x_1519 = temp__4386__auto___1518
									_ = x_1519
									if cljs_core.Contains_QMARK_.Arity2IIB(m_1516, x_1519) {
										m_1516, xs_1517___1 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(m_1516, x_1519), cljs_core.Next.Arity1IQ(xs_1517___1)
										continue
									} else {
										panic(cljs_core.Ex_info.X_invoke_Arity2("CLJS-849 regression!", (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "m", Fqn: "m", X_hash: float64(1632677161)}), cljs_core.Persistent_BANG_.X_invoke_Arity1(m_1516), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "xs", Fqn: "xs", X_hash: float64(649443341)}), xs_1517___1}, nil})).(*cljs_core.CljsCoreExceptionInfo))
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

var Foo2 *cljs_core.CljsCoreMultiFn

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
func (this__775__auto__ *CljsCore_testPerson) X_lookup_Arity2(k__776__auto__ interface{}) interface{} {
	return this__775__auto__.X_lookup_Arity3(k__776__auto__, nil)
}

func (this__777__auto__ *CljsCore_testPerson) X_lookup_Arity3(k665 interface{}, else__778__auto__ interface{}) interface{} {
	{
		var G__668 = func() interface{} {
			if cljs_core.Value_(k665).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k665), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__668
		switch G__668 {
		case "lastname":
			return this__777__auto__.Lastname

		case "firstname":
			return this__777__auto__.Firstname

		default:
			return cljs_core.Get.X_invoke_Arity3(this__777__auto__.X__extmap, k665, else__778__auto__)

		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIPrintWithWriter__() {}
func (this__791__auto__ *CljsCore_testPerson) X_pr_writer_Arity3(writer__792__auto__ interface{}, opts__793__auto__ interface{}) interface{} {
	{
		var pr_pair__794__auto__ = func(G__1521 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1521, 3, func(keyval__795__auto__ interface{}, ___796__auto__ interface{}, ___796__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, cljs_core.Pr_writer, "", " ", "", opts__793__auto__, keyval__795__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__794__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, pr_pair__794__auto__, "#cljs.core-test.Person{", ", ", "}", opts__793__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), this__791__auto__.Firstname}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), this__791__auto__.Lastname}, nil})}, nil}), this__791__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testPerson) CljsCoreIMeta__() {}
func (this__773__auto__ *CljsCore_testPerson) X_meta_Arity1() interface{} {
	return this__773__auto__.X__meta
}

func (_ *CljsCore_testPerson) CljsCoreICloneable__() {}
func (this__769__auto__ *CljsCore_testPerson) X_clone_Arity1() interface{} {
	return (&CljsCore_testPerson{this__769__auto__.Firstname, this__769__auto__.Lastname, this__769__auto__.X__meta, this__769__auto__.X__extmap, this__769__auto__.X__hash})
}

func (_ *CljsCore_testPerson) CljsCoreICounted__() {}
func (this__779__auto__ *CljsCore_testPerson) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__779__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testPerson) CljsCoreIHash__() {}
func (this__770__auto__ *CljsCore_testPerson) X_hash_Arity1() interface{} {
	{
		var h__593__auto__ = this__770__auto__.X__hash
		_ = h__593__auto__
		if !(cljs_core.Nil_(h__593__auto__)) {
			return h__593__auto__
		} else {
			{
				var h__593__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__770__auto__).(float64)
				_ = h__593__auto_____1
				this__770__auto__.X__hash = h__593__auto_____1

				return h__593__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIEquiv__() {}
func (this__771__auto__ *CljsCore_testPerson) X_equiv_Arity2(other__772__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__170__auto__ = other__772__auto__
		_ = and__170__auto__
		if cljs_core.Truth_(and__170__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__771__auto__), cljs_core.Type_.X_invoke_Arity1(other__772__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__771__auto__, other__772__auto__)))
		} else {
			return and__170__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testPerson) CljsCoreIRecord__() {}
func (_ *CljsCore_testPerson) CljsCoreIMap__()    {}
func (this__786__auto__ *CljsCore_testPerson) X_dissoc_Arity2(k__787__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), nil}, nil}, nil}), k__787__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__786__auto__), this__786__auto__.X__meta), k__787__auto__)
	} else {
		return (&CljsCore_testPerson{this__786__auto__.Firstname, this__786__auto__.Lastname, this__786__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__786__auto__.X__extmap, k__787__auto__)), nil})
	}
}

func (_ *CljsCore_testPerson) CljsCoreIAssociative__() {}
func (this__782__auto__ *CljsCore_testPerson) X_assoc_Arity3(k__783__auto__ interface{}, G__664 interface{}) interface{} {
	{
		var pred__676 = cljs_core.Keyword_identical_QMARK_
		var expr__677 = k__783__auto__
		_, _ = pred__676, expr__677
		if cljs_core.Truth_(func() interface{} {
			var G__679 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)})
			var G__680 = expr__677
			_, _ = G__679, G__680
			return pred__676.X_invoke_Arity2(G__679, G__680)
		}()) {
			return (&CljsCore_testPerson{G__664, this__782__auto__.Lastname, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__681 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})
				var G__682 = expr__677
				_, _ = G__681, G__682
				return pred__676.X_invoke_Arity2(G__681, G__682)
			}()) {
				return (&CljsCore_testPerson{this__782__auto__.Firstname, G__664, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPerson{this__782__auto__.Firstname, this__782__auto__.Lastname, this__782__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__782__auto__.X__extmap, k__783__auto__, G__664), nil})
			}
		}
	}
}

func (this__784__auto__ *CljsCore_testPerson) X_contains_key_QMARK__Arity2(k__785__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testPerson) CljsCoreISeqable__() {}
func (this__789__auto__ *CljsCore_testPerson) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), this__789__auto__.Firstname}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), this__789__auto__.Lastname}, nil})}, nil}), this__789__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testPerson) CljsCoreIWithMeta__() {}
func (this__774__auto__ *CljsCore_testPerson) X_with_meta_Arity2(G__664 interface{}) interface{} {
	return (&CljsCore_testPerson{this__774__auto__.Firstname, this__774__auto__.Lastname, G__664, this__774__auto__.X__extmap, this__774__auto__.X__hash})
}

func (_ *CljsCore_testPerson) CljsCoreICollection__() {}
func (this__780__auto__ *CljsCore_testPerson) X_conj_Arity2(entry__781__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__781__auto__) {
		return this__780__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__780__auto__, entry__781__auto__)
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
func (this__775__auto__ *CljsCore_testA) X_lookup_Arity2(k__776__auto__ interface{}) interface{} {
	return this__775__auto__.X_lookup_Arity3(k__776__auto__, nil)
}

func (this__777__auto__ *CljsCore_testA) X_lookup_Arity3(k684 interface{}, else__778__auto__ interface{}) interface{} {
	{
		var G__687 = k684
		_ = G__687
		switch G__687 {
		default:
			return cljs_core.Get.X_invoke_Arity3(this__777__auto__.X__extmap, k684, else__778__auto__)

		}
	}
}

func (_ *CljsCore_testA) CljsCoreIPrintWithWriter__() {}
func (this__791__auto__ *CljsCore_testA) X_pr_writer_Arity3(writer__792__auto__ interface{}, opts__793__auto__ interface{}) interface{} {
	{
		var pr_pair__794__auto__ = func(G__1523 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1523, 3, func(keyval__795__auto__ interface{}, ___796__auto__ interface{}, ___796__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, cljs_core.Pr_writer, "", " ", "", opts__793__auto__, keyval__795__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__794__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, pr_pair__794__auto__, "#cljs.core-test.A{", ", ", "}", opts__793__auto__, cljs_core.Concat.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, this__791__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testA) CljsCoreIMeta__() {}
func (this__773__auto__ *CljsCore_testA) X_meta_Arity1() interface{} {
	return this__773__auto__.X__meta
}

func (_ *CljsCore_testA) CljsCoreICloneable__() {}
func (this__769__auto__ *CljsCore_testA) X_clone_Arity1() interface{} {
	return (&CljsCore_testA{this__769__auto__.X__meta, this__769__auto__.X__extmap, this__769__auto__.X__hash})
}

func (_ *CljsCore_testA) CljsCoreICounted__() {}
func (this__779__auto__ *CljsCore_testA) X_count_Arity1() float64 {
	return (float64(0) + cljs_core.Count.X_invoke_Arity1(this__779__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testA) CljsCoreIHash__() {}
func (this__770__auto__ *CljsCore_testA) X_hash_Arity1() interface{} {
	{
		var h__593__auto__ = this__770__auto__.X__hash
		_ = h__593__auto__
		if !(cljs_core.Nil_(h__593__auto__)) {
			return h__593__auto__
		} else {
			{
				var h__593__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__770__auto__).(float64)
				_ = h__593__auto_____1
				this__770__auto__.X__hash = h__593__auto_____1

				return h__593__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testA) CljsCoreIEquiv__() {}
func (this__771__auto__ *CljsCore_testA) X_equiv_Arity2(other__772__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__170__auto__ = other__772__auto__
		_ = and__170__auto__
		if cljs_core.Truth_(and__170__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__771__auto__), cljs_core.Type_.X_invoke_Arity1(other__772__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__771__auto__, other__772__auto__)))
		} else {
			return and__170__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testA) CljsCoreIRecord__() {}
func (_ *CljsCore_testA) CljsCoreIMap__()    {}
func (this__786__auto__ *CljsCore_testA) X_dissoc_Arity2(k__787__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, k__787__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__786__auto__), this__786__auto__.X__meta), k__787__auto__)
	} else {
		return (&CljsCore_testA{this__786__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__786__auto__.X__extmap, k__787__auto__)), nil})
	}
}

func (_ *CljsCore_testA) CljsCoreIAssociative__() {}
func (this__782__auto__ *CljsCore_testA) X_assoc_Arity3(k__783__auto__ interface{}, G__683 interface{}) interface{} {
	{
		var pred__691 = cljs_core.Keyword_identical_QMARK_
		var expr__692 = k__783__auto__
		_, _ = pred__691, expr__692
		return (&CljsCore_testA{this__782__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__782__auto__.X__extmap, k__783__auto__, G__683), nil})
	}
}

func (this__784__auto__ *CljsCore_testA) X_contains_key_QMARK__Arity2(k__785__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testA) CljsCoreISeqable__() {}
func (this__789__auto__ *CljsCore_testA) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, this__789__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testA) CljsCoreIWithMeta__() {}
func (this__774__auto__ *CljsCore_testA) X_with_meta_Arity2(G__683 interface{}) interface{} {
	return (&CljsCore_testA{G__683, this__774__auto__.X__extmap, this__774__auto__.X__hash})
}

func (_ *CljsCore_testA) CljsCoreICollection__() {}
func (this__780__auto__ *CljsCore_testA) X_conj_Arity2(entry__781__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__781__auto__) {
		return this__780__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__780__auto__, entry__781__auto__)
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
func (this__775__auto__ *CljsCore_testC) X_lookup_Arity2(k__776__auto__ interface{}) interface{} {
	return this__775__auto__.X_lookup_Arity3(k__776__auto__, nil)
}

func (this__777__auto__ *CljsCore_testC) X_lookup_Arity3(k695 interface{}, else__778__auto__ interface{}) interface{} {
	{
		var G__698 = func() interface{} {
			if cljs_core.Value_(k695).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k695), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__698
		switch G__698 {
		case "c":
			return this__777__auto__.C

		case "b":
			return this__777__auto__.B

		case "a":
			return this__777__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__777__auto__.X__extmap, k695, else__778__auto__)

		}
	}
}

func (_ *CljsCore_testC) CljsCoreIPrintWithWriter__() {}
func (this__791__auto__ *CljsCore_testC) X_pr_writer_Arity3(writer__792__auto__ interface{}, opts__793__auto__ interface{}) interface{} {
	{
		var pr_pair__794__auto__ = func(G__1525 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1525, 3, func(keyval__795__auto__ interface{}, ___796__auto__ interface{}, ___796__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, cljs_core.Pr_writer, "", " ", "", opts__793__auto__, keyval__795__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__794__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, pr_pair__794__auto__, "#cljs.core-test.C{", ", ", "}", opts__793__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__791__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__791__auto__.B}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), this__791__auto__.C}, nil})}, nil}), this__791__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testC) CljsCoreIMeta__() {}
func (this__773__auto__ *CljsCore_testC) X_meta_Arity1() interface{} {
	return this__773__auto__.X__meta
}

func (_ *CljsCore_testC) CljsCoreICloneable__() {}
func (this__769__auto__ *CljsCore_testC) X_clone_Arity1() interface{} {
	return (&CljsCore_testC{this__769__auto__.A, this__769__auto__.B, this__769__auto__.C, this__769__auto__.X__meta, this__769__auto__.X__extmap, this__769__auto__.X__hash})
}

func (_ *CljsCore_testC) CljsCoreICounted__() {}
func (this__779__auto__ *CljsCore_testC) X_count_Arity1() float64 {
	return (float64(3) + cljs_core.Count.X_invoke_Arity1(this__779__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testC) CljsCoreIHash__() {}
func (this__770__auto__ *CljsCore_testC) X_hash_Arity1() interface{} {
	{
		var h__593__auto__ = this__770__auto__.X__hash
		_ = h__593__auto__
		if !(cljs_core.Nil_(h__593__auto__)) {
			return h__593__auto__
		} else {
			{
				var h__593__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__770__auto__).(float64)
				_ = h__593__auto_____1
				this__770__auto__.X__hash = h__593__auto_____1

				return h__593__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testC) CljsCoreIEquiv__() {}
func (this__771__auto__ *CljsCore_testC) X_equiv_Arity2(other__772__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__170__auto__ = other__772__auto__
		_ = and__170__auto__
		if cljs_core.Truth_(and__170__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__771__auto__), cljs_core.Type_.X_invoke_Arity1(other__772__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__771__auto__, other__772__auto__)))
		} else {
			return and__170__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testC) CljsCoreIRecord__() {}
func (_ *CljsCore_testC) CljsCoreIMap__()    {}
func (this__786__auto__ *CljsCore_testC) X_dissoc_Arity2(k__787__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__787__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__786__auto__), this__786__auto__.X__meta), k__787__auto__)
	} else {
		return (&CljsCore_testC{this__786__auto__.A, this__786__auto__.B, this__786__auto__.C, this__786__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__786__auto__.X__extmap, k__787__auto__)), nil})
	}
}

func (_ *CljsCore_testC) CljsCoreIAssociative__() {}
func (this__782__auto__ *CljsCore_testC) X_assoc_Arity3(k__783__auto__ interface{}, G__694 interface{}) interface{} {
	{
		var pred__708 = cljs_core.Keyword_identical_QMARK_
		var expr__709 = k__783__auto__
		_, _ = pred__708, expr__709
		if cljs_core.Truth_(func() interface{} {
			var G__711 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__712 = expr__709
			_, _ = G__711, G__712
			return pred__708.X_invoke_Arity2(G__711, G__712)
		}()) {
			return (&CljsCore_testC{G__694, this__782__auto__.B, this__782__auto__.C, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__713 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__714 = expr__709
				_, _ = G__713, G__714
				return pred__708.X_invoke_Arity2(G__713, G__714)
			}()) {
				return (&CljsCore_testC{this__782__auto__.A, G__694, this__782__auto__.C, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
			} else {
				if cljs_core.Truth_(func() interface{} {
					var G__715 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})
					var G__716 = expr__709
					_, _ = G__715, G__716
					return pred__708.X_invoke_Arity2(G__715, G__716)
				}()) {
					return (&CljsCore_testC{this__782__auto__.A, this__782__auto__.B, G__694, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
				} else {
					return (&CljsCore_testC{this__782__auto__.A, this__782__auto__.B, this__782__auto__.C, this__782__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__782__auto__.X__extmap, k__783__auto__, G__694), nil})
				}
			}
		}
	}
}

func (this__784__auto__ *CljsCore_testC) X_contains_key_QMARK__Arity2(k__785__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testC) CljsCoreISeqable__() {}
func (this__789__auto__ *CljsCore_testC) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__789__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__789__auto__.B}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), this__789__auto__.C}, nil})}, nil}), this__789__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testC) CljsCoreIWithMeta__() {}
func (this__774__auto__ *CljsCore_testC) X_with_meta_Arity2(G__694 interface{}) interface{} {
	return (&CljsCore_testC{this__774__auto__.A, this__774__auto__.B, this__774__auto__.C, G__694, this__774__auto__.X__extmap, this__774__auto__.X__hash})
}

func (_ *CljsCore_testC) CljsCoreICollection__() {}
func (this__780__auto__ *CljsCore_testC) X_conj_Arity2(entry__781__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__781__auto__) {
		return this__780__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__780__auto__, entry__781__auto__)
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
func (this__775__auto__ *CljsCore_testA2) X_lookup_Arity2(k__776__auto__ interface{}) interface{} {
	return this__775__auto__.X_lookup_Arity3(k__776__auto__, nil)
}

func (this__777__auto__ *CljsCore_testA2) X_lookup_Arity3(k720 interface{}, else__778__auto__ interface{}) interface{} {
	{
		var G__723 = func() interface{} {
			if cljs_core.Value_(k720).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k720), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__723
		switch G__723 {
		case "x":
			return this__777__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__777__auto__.X__extmap, k720, else__778__auto__)

		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIPrintWithWriter__() {}
func (this__791__auto__ *CljsCore_testA2) X_pr_writer_Arity3(writer__792__auto__ interface{}, opts__793__auto__ interface{}) interface{} {
	{
		var pr_pair__794__auto__ = func(G__1527 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1527, 3, func(keyval__795__auto__ interface{}, ___796__auto__ interface{}, ___796__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, cljs_core.Pr_writer, "", " ", "", opts__793__auto__, keyval__795__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__794__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, pr_pair__794__auto__, "#cljs.core-test.A2{", ", ", "}", opts__793__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__791__auto__.X}, nil})}, nil}), this__791__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testA2) CljsCoreIMeta__() {}
func (this__773__auto__ *CljsCore_testA2) X_meta_Arity1() interface{} {
	return this__773__auto__.X__meta
}

func (_ *CljsCore_testA2) CljsCoreICloneable__() {}
func (this__769__auto__ *CljsCore_testA2) X_clone_Arity1() interface{} {
	return (&CljsCore_testA2{this__769__auto__.X, this__769__auto__.X__meta, this__769__auto__.X__extmap, this__769__auto__.X__hash})
}

func (_ *CljsCore_testA2) CljsCoreICounted__() {}
func (this__779__auto__ *CljsCore_testA2) X_count_Arity1() float64 {
	return (float64(1) + cljs_core.Count.X_invoke_Arity1(this__779__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testA2) CljsCoreIHash__() {}
func (this__770__auto__ *CljsCore_testA2) X_hash_Arity1() interface{} {
	{
		var h__593__auto__ = this__770__auto__.X__hash
		_ = h__593__auto__
		if !(cljs_core.Nil_(h__593__auto__)) {
			return h__593__auto__
		} else {
			{
				var h__593__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__770__auto__).(float64)
				_ = h__593__auto_____1
				this__770__auto__.X__hash = h__593__auto_____1

				return h__593__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIEquiv__() {}
func (this__771__auto__ *CljsCore_testA2) X_equiv_Arity2(other__772__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__170__auto__ = other__772__auto__
		_ = and__170__auto__
		if cljs_core.Truth_(and__170__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__771__auto__), cljs_core.Type_.X_invoke_Arity1(other__772__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__771__auto__, other__772__auto__)))
		} else {
			return and__170__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testA2) CljsCoreIRecord__() {}
func (_ *CljsCore_testA2) CljsCoreIMap__()    {}
func (this__786__auto__ *CljsCore_testA2) X_dissoc_Arity2(k__787__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), nil}, nil}, nil}), k__787__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__786__auto__), this__786__auto__.X__meta), k__787__auto__)
	} else {
		return (&CljsCore_testA2{this__786__auto__.X, this__786__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__786__auto__.X__extmap, k__787__auto__)), nil})
	}
}

func (_ *CljsCore_testA2) CljsCoreIAssociative__() {}
func (this__782__auto__ *CljsCore_testA2) X_assoc_Arity3(k__783__auto__ interface{}, G__719 interface{}) interface{} {
	{
		var pred__729 = cljs_core.Keyword_identical_QMARK_
		var expr__730 = k__783__auto__
		_, _ = pred__729, expr__730
		if cljs_core.Truth_(func() interface{} {
			var G__732 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__733 = expr__730
			_, _ = G__732, G__733
			return pred__729.X_invoke_Arity2(G__732, G__733)
		}()) {
			return (&CljsCore_testA2{G__719, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testA2{this__782__auto__.X, this__782__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__782__auto__.X__extmap, k__783__auto__, G__719), nil})
		}
	}
}

func (this__784__auto__ *CljsCore_testA2) X_contains_key_QMARK__Arity2(k__785__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testA2) CljsCoreISeqable__() {}
func (this__789__auto__ *CljsCore_testA2) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__789__auto__.X}, nil})}, nil}), this__789__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testA2) CljsCoreIWithMeta__() {}
func (this__774__auto__ *CljsCore_testA2) X_with_meta_Arity2(G__719 interface{}) interface{} {
	return (&CljsCore_testA2{this__774__auto__.X, G__719, this__774__auto__.X__extmap, this__774__auto__.X__hash})
}

func (_ *CljsCore_testA2) CljsCoreICollection__() {}
func (this__780__auto__ *CljsCore_testA2) X_conj_Arity2(entry__781__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__781__auto__) {
		return this__780__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__780__auto__, entry__781__auto__)
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
func (this__775__auto__ *CljsCore_testB) X_lookup_Arity2(k__776__auto__ interface{}) interface{} {
	return this__775__auto__.X_lookup_Arity3(k__776__auto__, nil)
}

func (this__777__auto__ *CljsCore_testB) X_lookup_Arity3(k735 interface{}, else__778__auto__ interface{}) interface{} {
	{
		var G__738 = func() interface{} {
			if cljs_core.Value_(k735).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k735), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__738
		switch G__738 {
		case "x":
			return this__777__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__777__auto__.X__extmap, k735, else__778__auto__)

		}
	}
}

func (_ *CljsCore_testB) CljsCoreIPrintWithWriter__() {}
func (this__791__auto__ *CljsCore_testB) X_pr_writer_Arity3(writer__792__auto__ interface{}, opts__793__auto__ interface{}) interface{} {
	{
		var pr_pair__794__auto__ = func(G__1529 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1529, 3, func(keyval__795__auto__ interface{}, ___796__auto__ interface{}, ___796__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, cljs_core.Pr_writer, "", " ", "", opts__793__auto__, keyval__795__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__794__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, pr_pair__794__auto__, "#cljs.core-test.B{", ", ", "}", opts__793__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__791__auto__.X}, nil})}, nil}), this__791__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testB) CljsCoreIMeta__() {}
func (this__773__auto__ *CljsCore_testB) X_meta_Arity1() interface{} {
	return this__773__auto__.X__meta
}

func (_ *CljsCore_testB) CljsCoreICloneable__() {}
func (this__769__auto__ *CljsCore_testB) X_clone_Arity1() interface{} {
	return (&CljsCore_testB{this__769__auto__.X, this__769__auto__.X__meta, this__769__auto__.X__extmap, this__769__auto__.X__hash})
}

func (_ *CljsCore_testB) CljsCoreICounted__() {}
func (this__779__auto__ *CljsCore_testB) X_count_Arity1() float64 {
	return (float64(1) + cljs_core.Count.X_invoke_Arity1(this__779__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testB) CljsCoreIHash__() {}
func (this__770__auto__ *CljsCore_testB) X_hash_Arity1() interface{} {
	{
		var h__593__auto__ = this__770__auto__.X__hash
		_ = h__593__auto__
		if !(cljs_core.Nil_(h__593__auto__)) {
			return h__593__auto__
		} else {
			{
				var h__593__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__770__auto__).(float64)
				_ = h__593__auto_____1
				this__770__auto__.X__hash = h__593__auto_____1

				return h__593__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testB) CljsCoreIEquiv__() {}
func (this__771__auto__ *CljsCore_testB) X_equiv_Arity2(other__772__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__170__auto__ = other__772__auto__
		_ = and__170__auto__
		if cljs_core.Truth_(and__170__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__771__auto__), cljs_core.Type_.X_invoke_Arity1(other__772__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__771__auto__, other__772__auto__)))
		} else {
			return and__170__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testB) CljsCoreIRecord__() {}
func (_ *CljsCore_testB) CljsCoreIMap__()    {}
func (this__786__auto__ *CljsCore_testB) X_dissoc_Arity2(k__787__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), nil}, nil}, nil}), k__787__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__786__auto__), this__786__auto__.X__meta), k__787__auto__)
	} else {
		return (&CljsCore_testB{this__786__auto__.X, this__786__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__786__auto__.X__extmap, k__787__auto__)), nil})
	}
}

func (_ *CljsCore_testB) CljsCoreIAssociative__() {}
func (this__782__auto__ *CljsCore_testB) X_assoc_Arity3(k__783__auto__ interface{}, G__734 interface{}) interface{} {
	{
		var pred__744 = cljs_core.Keyword_identical_QMARK_
		var expr__745 = k__783__auto__
		_, _ = pred__744, expr__745
		if cljs_core.Truth_(func() interface{} {
			var G__747 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__748 = expr__745
			_, _ = G__747, G__748
			return pred__744.X_invoke_Arity2(G__747, G__748)
		}()) {
			return (&CljsCore_testB{G__734, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testB{this__782__auto__.X, this__782__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__782__auto__.X__extmap, k__783__auto__, G__734), nil})
		}
	}
}

func (this__784__auto__ *CljsCore_testB) X_contains_key_QMARK__Arity2(k__785__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testB) CljsCoreISeqable__() {}
func (this__789__auto__ *CljsCore_testB) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), this__789__auto__.X}, nil})}, nil}), this__789__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testB) CljsCoreIWithMeta__() {}
func (this__774__auto__ *CljsCore_testB) X_with_meta_Arity2(G__734 interface{}) interface{} {
	return (&CljsCore_testB{this__774__auto__.X, G__734, this__774__auto__.X__extmap, this__774__auto__.X__hash})
}

func (_ *CljsCore_testB) CljsCoreICollection__() {}
func (this__780__auto__ *CljsCore_testB) X_conj_Arity2(entry__781__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__781__auto__) {
		return this__780__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__780__auto__, entry__781__auto__)
	}
}

var X__GT_B *cljs_core.AFn

var Map__GT_B *cljs_core.AFn

type CljsCore_testIFoo interface {
	CljsCore_testIFoo__()
	Foo_Arity1() interface{}
}

var Foo *cljs_core.AFn

type CljsCore_testT749 struct {
	Test_stuff interface{}
	Meta750    interface{}
}

func (_ *CljsCore_testT749) CljsCore_testIFoo__() {}
func (this *CljsCore_testT749) Foo_Arity1() interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
}

func (_ *CljsCore_testT749) CljsCoreIMeta__() {}
func (_751 *CljsCore_testT749) X_meta_Arity1() interface{} {
	return _751.Meta750
}

func (_ *CljsCore_testT749) CljsCoreIWithMeta__() {}
func (_751 *CljsCore_testT749) X_with_meta_Arity2(meta750___1 interface{}) interface{} {
	return (&CljsCore_testT749{_751.Test_stuff, meta750___1})
}

var X__GT_t749 *cljs_core.AFn

type CljsCore_testIMutate interface {
	CljsCore_testIMutate__()
	Mutate_Arity1() interface{}
}

var Mutate *cljs_core.AFn

type CljsCore_testMutate struct{ A interface{} }

func (_ *CljsCore_testMutate) CljsCore_testIMutate__() {}
func (___ *CljsCore_testMutate) Mutate_Arity1() interface{} {
	return func() interface{} {
		var return__1530 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})
		___.A = return__1530
		return return__1530
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
func (___ *CljsCore_testFirst) X_find_first_Arity2(p__752 interface{}) interface{} {
	{
		var vec__754 = p__752
		var x = cljs_core.Nth.X_invoke_Arity3(vec__754, float64(0), nil)
		_, _ = vec__754, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCore_testIHasFirst__() {}
func (p__755 *CljsCore_testFirst) X_get_first_Arity1() interface{} {
	{
		var vec__757 = p__755
		var x = cljs_core.Nth.X_invoke_Arity3(vec__757, float64(0), nil)
		_, _ = vec__757, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCoreObject__() {}
func (p__758 *CljsCore_testFirst) ToString() string {
	{
		var vec__760 = p__758
		var x = cljs_core.Nth.X_invoke_Arity3(vec__760, float64(0), nil)
		_, _ = vec__760, x
		return (`` + cljs_core.Str.X_invoke_Arity1(x).(string))
	}
}

func (this *CljsCore_testFirst) String() string {
	return this.ToString()
}

func (_ *CljsCore_testFirst) CljsCoreIFn__() {}
func (p__761 *CljsCore_testFirst) X_invoke_Arity0() interface{} {
	{
		var vec__763 = p__761
		var x = cljs_core.Nth.X_invoke_Arity3(vec__763, float64(0), nil)
		_, _ = vec__763, x
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
func (___ *CljsCore_testDestructuringWithLocals) X_find_first_Arity2(p__765 interface{}) interface{} {
	{
		var vec__767 = p__765
		var x = cljs_core.Nth.X_invoke_Arity3(vec__767, float64(0), nil)
		var y = cljs_core.Nth.X_invoke_Arity3(vec__767, float64(1), nil)
		_, _, _ = vec__767, x, y
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
func (this__775__auto__ *CljsCore_testPrintMe) X_lookup_Arity2(k__776__auto__ interface{}) interface{} {
	return this__775__auto__.X_lookup_Arity3(k__776__auto__, nil)
}

func (this__777__auto__ *CljsCore_testPrintMe) X_lookup_Arity3(k779 interface{}, else__778__auto__ interface{}) interface{} {
	{
		var G__782 = func() interface{} {
			if cljs_core.Value_(k779).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k779), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__782
		switch G__782 {
		case "b":
			return this__777__auto__.B

		case "a":
			return this__777__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__777__auto__.X__extmap, k779, else__778__auto__)

		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIPrintWithWriter__() {}
func (this__791__auto__ *CljsCore_testPrintMe) X_pr_writer_Arity3(writer__792__auto__ interface{}, opts__793__auto__ interface{}) interface{} {
	{
		var pr_pair__794__auto__ = func(G__1532 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1532, 3, func(keyval__795__auto__ interface{}, ___796__auto__ interface{}, ___796__auto_____1 interface{}) interface{} {
				return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, cljs_core.Pr_writer, "", " ", "", opts__793__auto__, keyval__795__auto__)
			})
		}(&cljs_core.AFn{})
		_ = pr_pair__794__auto__
		return cljs_core.Pr_sequential_writer.X_invoke_Arity7(writer__792__auto__, pr_pair__794__auto__, "#cljs.core-test.PrintMe{", ", ", "}", opts__793__auto__, cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__791__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__791__auto__.B}, nil})}, nil}), this__791__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIMeta__() {}
func (this__773__auto__ *CljsCore_testPrintMe) X_meta_Arity1() interface{} {
	return this__773__auto__.X__meta
}

func (_ *CljsCore_testPrintMe) CljsCoreICloneable__() {}
func (this__769__auto__ *CljsCore_testPrintMe) X_clone_Arity1() interface{} {
	return (&CljsCore_testPrintMe{this__769__auto__.A, this__769__auto__.B, this__769__auto__.X__meta, this__769__auto__.X__extmap, this__769__auto__.X__hash})
}

func (_ *CljsCore_testPrintMe) CljsCoreICounted__() {}
func (this__779__auto__ *CljsCore_testPrintMe) X_count_Arity1() float64 {
	return (float64(2) + cljs_core.Count.X_invoke_Arity1(this__779__auto__.X__extmap).(float64))
}

func (_ *CljsCore_testPrintMe) CljsCoreIHash__() {}
func (this__770__auto__ *CljsCore_testPrintMe) X_hash_Arity1() interface{} {
	{
		var h__593__auto__ = this__770__auto__.X__hash
		_ = h__593__auto__
		if !(cljs_core.Nil_(h__593__auto__)) {
			return h__593__auto__
		} else {
			{
				var h__593__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__770__auto__).(float64)
				_ = h__593__auto_____1
				this__770__auto__.X__hash = h__593__auto_____1

				return h__593__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIEquiv__() {}
func (this__771__auto__ *CljsCore_testPrintMe) X_equiv_Arity2(other__772__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__170__auto__ = other__772__auto__
		_ = and__170__auto__
		if cljs_core.Truth_(and__170__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__771__auto__), cljs_core.Type_.X_invoke_Arity1(other__772__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__771__auto__, other__772__auto__)))
		} else {
			return and__170__auto__
		}
	}()) {
		return true
	} else {
		return false
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIRecord__() {}
func (_ *CljsCore_testPrintMe) CljsCoreIMap__()    {}
func (this__786__auto__ *CljsCore_testPrintMe) X_dissoc_Arity2(k__787__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__787__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__786__auto__), this__786__auto__.X__meta), k__787__auto__)
	} else {
		return (&CljsCore_testPrintMe{this__786__auto__.A, this__786__auto__.B, this__786__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__786__auto__.X__extmap, k__787__auto__)), nil})
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIAssociative__() {}
func (this__782__auto__ *CljsCore_testPrintMe) X_assoc_Arity3(k__783__auto__ interface{}, G__778 interface{}) interface{} {
	{
		var pred__790 = cljs_core.Keyword_identical_QMARK_
		var expr__791 = k__783__auto__
		_, _ = pred__790, expr__791
		if cljs_core.Truth_(func() interface{} {
			var G__793 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__794 = expr__791
			_, _ = G__793, G__794
			return pred__790.X_invoke_Arity2(G__793, G__794)
		}()) {
			return (&CljsCore_testPrintMe{G__778, this__782__auto__.B, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__795 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__796 = expr__791
				_, _ = G__795, G__796
				return pred__790.X_invoke_Arity2(G__795, G__796)
			}()) {
				return (&CljsCore_testPrintMe{this__782__auto__.A, G__778, this__782__auto__.X__meta, this__782__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPrintMe{this__782__auto__.A, this__782__auto__.B, this__782__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__782__auto__.X__extmap, k__783__auto__, G__778), nil})
			}
		}
	}
}

func (this__784__auto__ *CljsCore_testPrintMe) X_contains_key_QMARK__Arity2(k__785__auto__ interface{}) bool {
	panic((&js.Error{"Not implemented."}))
}

func (_ *CljsCore_testPrintMe) CljsCoreISeqable__() {}
func (this__789__auto__ *CljsCore_testPrintMe) X_seq_Arity1() interface{} {
	return cljs_core.Seq.Arity1IQ(cljs_core.Concat.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), this__789__auto__.A}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), this__789__auto__.B}, nil})}, nil}), this__789__auto__.X__extmap).(*cljs_core.CljsCoreLazySeq))
}

func (_ *CljsCore_testPrintMe) CljsCoreIWithMeta__() {}
func (this__774__auto__ *CljsCore_testPrintMe) X_with_meta_Arity2(G__778 interface{}) interface{} {
	return (&CljsCore_testPrintMe{this__774__auto__.A, this__774__auto__.B, G__778, this__774__auto__.X__extmap, this__774__auto__.X__hash})
}

func (_ *CljsCore_testPrintMe) CljsCoreICollection__() {}
func (this__780__auto__ *CljsCore_testPrintMe) X_conj_Arity2(entry__781__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__781__auto__) {
		return this__780__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__781__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
	} else {
		return cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_conj, this__780__auto__, entry__781__auto__)
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

type CljsCore_testT830 struct {
	F          interface{}
	Baz        interface{}
	Test_stuff interface{}
	Meta831    interface{}
}

func (_ *CljsCore_testT830) CljsCore_testIBar__() {}
func (___ *CljsCore_testT830) X_bar_Arity2(x interface{}) interface{} {
	{
		var G__834 = x
		_ = G__834
		return ___.F.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__834)
	}
}

func (_ *CljsCore_testT830) CljsCoreIMeta__() {}
func (_832 *CljsCore_testT830) X_meta_Arity1() interface{} {
	return _832.Meta831
}

func (_ *CljsCore_testT830) CljsCoreIWithMeta__() {}
func (_832 *CljsCore_testT830) X_with_meta_Arity2(meta831___1 interface{}) interface{} {
	return (&CljsCore_testT830{_832.F, _832.Baz, _832.Test_stuff, meta831___1})
}

var X__GT_t830 *cljs_core.AFn

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

type CljsCore_testT844 struct {
	Test_stuff interface{}
	Meta845    interface{}
}

func (_ *CljsCore_testT844) CljsCoreIHash__() {}
func (___ *CljsCore_testT844) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT844) CljsCoreIMeta__() {}
func (_846 *CljsCore_testT844) X_meta_Arity1() interface{} {
	return _846.Meta845
}

func (_ *CljsCore_testT844) CljsCoreIWithMeta__() {}
func (_846 *CljsCore_testT844) X_with_meta_Arity2(meta845___1 interface{}) interface{} {
	return (&CljsCore_testT844{_846.Test_stuff, meta845___1})
}

var X__GT_t844 *cljs_core.AFn

type CljsCore_testT847 struct {
	A          interface{}
	Test_stuff interface{}
	Meta848    interface{}
}

func (_ *CljsCore_testT847) CljsCoreIHash__() {}
func (___ *CljsCore_testT847) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT847) CljsCoreIMeta__() {}
func (_849 *CljsCore_testT847) X_meta_Arity1() interface{} {
	return _849.Meta848
}

func (_ *CljsCore_testT847) CljsCoreIWithMeta__() {}
func (_849 *CljsCore_testT847) X_with_meta_Arity2(meta848___1 interface{}) interface{} {
	return (&CljsCore_testT847{_849.A, _849.Test_stuff, meta848___1})
}

var X__GT_t847 *cljs_core.AFn

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

type CljsCore_testT917 struct {
	From_seq   interface{}
	Make_seq   interface{}
	Mt         interface{}
	I__909     interface{}
	Count__908 interface{}
	Chunk__907 interface{}
	Seq__906   interface{}
	Test_stuff interface{}
	Meta918    interface{}
}

func (_ *CljsCore_testT917) CljsCoreISeq__() {}
func (this *CljsCore_testT917) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT917) X_rest_Arity1() interface{} {
	{
		var G__921 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__921
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__921)
	}
}

func (_ *CljsCore_testT917) CljsCoreISeqable__() {}
func (this *CljsCore_testT917) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT917) CljsCoreIMeta__() {}
func (_919 *CljsCore_testT917) X_meta_Arity1() interface{} {
	return _919.Meta918
}

func (_ *CljsCore_testT917) CljsCoreIWithMeta__() {}
func (_919 *CljsCore_testT917) X_with_meta_Arity2(meta918___1 interface{}) interface{} {
	return (&CljsCore_testT917{_919.From_seq, _919.Make_seq, _919.Mt, _919.I__909, _919.Count__908, _919.Chunk__907, _919.Seq__906, _919.Test_stuff, meta918___1})
}

var X__GT_t917 *cljs_core.AFn

type CljsCore_testT929 struct {
	From_seq           interface{}
	Make_seq           interface{}
	Mt                 interface{}
	Temp__4388__auto__ interface{}
	I__909             interface{}
	Count__908         interface{}
	Chunk__907         interface{}
	Seq__906           interface{}
	Test_stuff         interface{}
	Meta930            interface{}
}

func (_ *CljsCore_testT929) CljsCoreISeq__() {}
func (this *CljsCore_testT929) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT929) X_rest_Arity1() interface{} {
	{
		var G__933 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__933
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__933)
	}
}

func (_ *CljsCore_testT929) CljsCoreISeqable__() {}
func (this *CljsCore_testT929) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT929) CljsCoreIMeta__() {}
func (_931 *CljsCore_testT929) X_meta_Arity1() interface{} {
	return _931.Meta930
}

func (_ *CljsCore_testT929) CljsCoreIWithMeta__() {}
func (_931 *CljsCore_testT929) X_with_meta_Arity2(meta930___1 interface{}) interface{} {
	return (&CljsCore_testT929{_931.From_seq, _931.Make_seq, _931.Mt, _931.Temp__4388__auto__, _931.I__909, _931.Count__908, _931.Chunk__907, _931.Seq__906, _931.Test_stuff, meta930___1})
}

var X__GT_t929 *cljs_core.AFn

var Case_recur *cljs_core.AFn

var Xform *cljs_core.AFn

var Data interface{}

var Xf *cljs_core.AFn

func Test_runner(t *testing.T) {
	assert.Equal(t, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), Test_stuff.X_invoke_Arity0().(*cljs_core.CljsCoreKeyword))
}
