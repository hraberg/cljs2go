// Compiled by ClojureScript to Go 0.0-2411
// cljs.core-test

package core_test

import (
	"reflect"
	"strings"
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 1)").(string)}, ``)}))
			}
			if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 1 1)").(string)}, ``)}))
			}
			if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(1)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 1 1 1)").(string)}, ``)}))
			}
			if !(cljs_core.X_EQ_.Arity2IIB(float64(1), float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (= 1 2))").(string)}, ``)}))
			}
			if !(cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (= 1 2 1))").(string)}, ``)}))
			}
			if !(cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (= 1 1 2))").(string)}, ``)}))
			}
			if !(cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(1)})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (= 1 1 2 1))").(string)}, ``)}))
			}
			if !(cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(float64(1), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (= 1 1 1 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (+) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.CljsCorePersistentVector_EMPTY), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply + []) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (+ 1) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply + [1]) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) + float64(1)), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (+ 1 1) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply + [1 1]) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(((float64(1) + float64(2)) + float64(3)), float64(6)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (+ 1 2 3) 6)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})), float64(6)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply + [1 2 3]) 6)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((-float64(1)), float64(-1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (- 1) -1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X___, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(-1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply - [1]) -1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) - float64(1)), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (- 1 1) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X___, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply - [1 1]) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(((float64(3) - float64(2)) - float64(1)), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (- 3 2 1) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X___, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil})), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply - [3 2 1]) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (*) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_STAR_, cljs_core.CljsCorePersistentVector_EMPTY), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply * []) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (* 2) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_STAR_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil})), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply * [2]) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(2) * float64(3)), float64(6)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (* 2 3) 6)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_STAR_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})), float64(6)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply * [2 3]) 6)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) / float64(2)), 0.5) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (/ 2) 0.5)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_SLASH_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil})), 0.5) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply / [2]) 0.5)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(6) / float64(2)), float64(3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (/ 6 2) 3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_SLASH_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(6), float64(2)}, nil})), float64(3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply / [6 2]) 3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(((float64(6) / float64(3)) / float64(2)), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (/ 6 3 2) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_SLASH_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(6), float64(3), float64(2)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply / [6 3 2]) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(true, true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (< 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply < [1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) < float64(2)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (< 1 2) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply < [1 2]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) < float64(1)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (< 1 1) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply < [1 1]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(2) < float64(1)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (< 2 1) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply < [2 1]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) < float64(2)) && (float64(2) < float64(3)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (< 1 2 3) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply < [1 2 3]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) < float64(1)) && (float64(1) < float64(3)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (< 1 1 3) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(3)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply < [1 1 3]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) < float64(1)) && (float64(1) < float64(1)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (< 3 1 1) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply < [3 1 1]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(true, true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (<= 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply <= [1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) <= float64(1)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (<= 1 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply <= [1 1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) <= float64(2)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (<= 1 2) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply <= [1 2]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(2) <= float64(1)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (<= 2 1) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply <= [2 1]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) <= float64(2)) && (float64(2) <= float64(3)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (<= 1 2 3) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply <= [1 2 3]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) <= float64(1)) && (float64(1) <= float64(3)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (<= 1 1 3) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(3)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply <= [1 1 3]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) <= float64(1)) && (float64(1) <= float64(1)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (<= 3 1 1) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_LT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply <= [3 1 1]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(true, true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (> 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply > [1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(2) > float64(1)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (> 2 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply > [2 1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) > float64(1)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (> 1 1) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply > [1 1]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) > float64(2)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (> 1 2) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply > [1 2]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) > float64(2)) && (float64(2) > float64(1)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (> 3 2 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply > [3 2 1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) > float64(1)) && (float64(1) > float64(1)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (> 3 1 1) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply > [3 1 1]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) > float64(1)) && (float64(1) > float64(3)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (> 1 1 3) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(3)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply > [1 1 3]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(true, true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (>= 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply >= [1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(2) >= float64(1)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (>= 2 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply >= [2 1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) >= float64(1)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (>= 1 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply >= [1 1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) >= float64(2)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (>= 1 2) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply >= [1 2]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) >= float64(2)) && (float64(2) >= float64(1)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (>= 3 2 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply >= [3 2 1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) >= float64(1)) && (float64(1) >= float64(1)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (>= 3 1 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply >= [3 1 1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(3) >= float64(1)) && (float64(1) >= float64(2)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (>= 3 1 2) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(2)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply >= [3 1 2]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) >= float64(1)) && (float64(1) >= float64(3)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (>= 1 1 3) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.X_GT__EQ_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(3)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply >= [1 1 3]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) - float64(1)), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (dec 1) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Dec, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply dec [1]) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(0) + float64(1)), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (inc 0) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Inc, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply inc [0]) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(0) == float64(0)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (zero? 0) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Zero_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply zero? [0]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) == float64(0)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (zero? 1) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Zero_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply zero? [1]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(-11) == float64(0)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (zero? -11) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Zero_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-11)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply zero? [-11]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(0) > float64(0)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pos? 0) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Pos_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply pos? [0]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(1) > float64(0)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pos? 1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Pos_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply pos? [1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(-1) > float64(0)), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pos? -1) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Pos_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-1)}, nil})), false) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply pos? [-1]) false)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((float64(-1) < float64(0)), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (neg? -1) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Neg_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-1)}, nil})), true) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply neg? [-1]) true)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (max 1) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply max [1]) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(float64(1), float64(2)), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (max 1 2) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply max [1 2]) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x > y {
					return x
				} else {
					return y
				}
			}(float64(2), float64(1)), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (max 2 1) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply max [2 1]) 2)").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (max 1 2 3) 3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})), float64(3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply max [1 2 3]) 3)").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (max 1 3 2) 3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Max, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(2)}, nil})), float64(3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply max [1 3 2]) 3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (min 1) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Min, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply min [1]) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(float64(1), float64(2)), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (min 1 2) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Min, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply min [1 2]) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func(x, y float64) float64 {
				if x < y {
					return x
				} else {
					return y
				}
			}(float64(2), float64(1)), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (min 2 1) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Min, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply min [2 1]) 1)").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (min 1 2 3) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Min, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply min [1 2 3]) 1)").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (min 2 1 3) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Min, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1), float64(3)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply min [3 1 3]) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(4), float64(2)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 4 2) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Mod, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(2)}, nil})), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply mod [4 2]) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(3), float64(2)).(float64), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 3 2) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Mod, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2)}, nil})), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply mod [3 2]) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(-2), float64(5)).(float64), float64(3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod -2 5) 3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(3), float64(2), float64(1), float64(0)}, nil}), func() interface{} {
				var i = float64(0)
				var j interface{} = cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)
				_, _ = i, j
				for {
					if i < float64(5) {
						i, j = (i + float64(1)), cljs_core.Conj.X_invoke_Arity2(j, func(G__957 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__957, 0, func() interface{} {
								return i
							})
						}(&cljs_core.AFn{}, i, j))
						continue
					} else {
						return cljs_core.Map_.X_invoke_Arity2(func(G__958 *cljs_core.AFn, i float64, j interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__958, 1, func(p1__56_SHARP_ interface{}) interface{} {
								{
									return p1__56_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
								}
							})
						}(&cljs_core.AFn{}, i, j), j).(*cljs_core.CljsCoreLazySeq)
					}
				}
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [4 3 2 1 0] (loop [i 0 j ()] (if (< i 5) (recur (inc i) (conj j (fn [] i))) (map (fn* [p1__56#] (p1__56#)) j))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})}, nil}), cljs_core.Map_.X_invoke_Arity2(func(G__959 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__959, 1, func(p1__57_SHARP_ interface{}) interface{} {
					{
						return p1__57_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
					}
				})
			}(&cljs_core.AFn{}), func() *cljs_core.CljsCoreLazySeq {
				var iter__939__auto__ = func(iter__526 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(iter__526, 1, func(s__527 interface{}) interface{} {
						return (&cljs_core.CljsCoreLazySeq{nil, func(G__960 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__960, 0, func() interface{} {
								{
									var s__527___1 interface{} = s__527
									_ = s__527___1
									for {
										{
											var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__527___1)
											_ = temp__4388__auto__
											if cljs_core.Truth_(temp__4388__auto__) {
												{
													var xs__4940__auto__ = temp__4388__auto__
													_ = xs__4940__auto__
													{
														var i = cljs_core.First.X_invoke_Arity1(xs__4940__auto__)
														_ = i
														{
															var iterys__935__auto__ = func(iter__528 *cljs_core.AFn, s__527___1 interface{}, i interface{}, xs__4940__auto__ cljs_core.CljsCoreISeq, temp__4388__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																return cljs_core.Fn(iter__528, 1, func(s__529 interface{}) interface{} {
																	return (&cljs_core.CljsCoreLazySeq{nil, func(G__961 *cljs_core.AFn, s__527___1 interface{}, i interface{}, xs__4940__auto__ cljs_core.CljsCoreISeq, temp__4388__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																		return cljs_core.Fn(G__961, 0, func() interface{} {
																			{
																				var s__529___1 interface{} = s__529
																				_ = s__529___1
																				for {
																					{
																						var temp__4388__auto_____1 = cljs_core.Seq.Arity1IQ(s__529___1)
																						_ = temp__4388__auto_____1
																						if cljs_core.Truth_(temp__4388__auto_____1) {
																							{
																								var s__529___2 = temp__4388__auto_____1
																								_ = s__529___2
																								if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__529___2) {
																									{
																										var c__937__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__529___2)
																										var size__938__auto__ = cljs_core.Count.X_invoke_Arity1(c__937__auto__).(float64)
																										var b__531 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__938__auto__).(*cljs_core.CljsCoreChunkBuffer)
																										_, _, _ = c__937__auto__, size__938__auto__, b__531
																										if func() bool {
																											var i__530 = float64(0)
																											_ = i__530
																											for {
																												if i__530 < size__938__auto__ {
																													{
																														var j = cljs_core.Decorate_(c__937__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__530)
																														_ = j
																														cljs_core.Chunk_append.X_invoke_Arity2(b__531, func(G__962 *cljs_core.AFn, i__530 float64, s__527___1 interface{}, j interface{}, c__937__auto__ interface{}, size__938__auto__ float64, b__531 *cljs_core.CljsCoreChunkBuffer, s__529___2 interface{}, temp__4388__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4940__auto__ cljs_core.CljsCoreISeq, temp__4388__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																															return cljs_core.Fn(G__962, 0, func() interface{} {
																																return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																															})
																														}(&cljs_core.AFn{}, i__530, s__527___1, j, c__937__auto__, size__938__auto__, b__531, s__529___2, temp__4388__auto_____1, i, xs__4940__auto__, temp__4388__auto__))
																														i__530 = (i__530 + float64(1))
																														continue
																													}
																												} else {
																													return true
																												}
																											}
																										}() {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__531), iter__528.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__529___2)).(*cljs_core.CljsCoreLazySeq))
																										} else {
																											return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__531), nil)
																										}
																									}
																								} else {
																									{
																										var j = cljs_core.First.X_invoke_Arity1(s__529___2)
																										_ = j
																										return cljs_core.Cons.X_invoke_Arity2(func(G__963 *cljs_core.AFn, s__527___1 interface{}, j interface{}, s__529___2 interface{}, temp__4388__auto_____1 cljs_core.CljsCoreISeq, i interface{}, xs__4940__auto__ cljs_core.CljsCoreISeq, temp__4388__auto__ cljs_core.CljsCoreISeq) *cljs_core.AFn {
																											return cljs_core.Fn(G__963, 0, func() interface{} {
																												return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{i, j}, nil})
																											})
																										}(&cljs_core.AFn{}, s__527___1, j, s__529___2, temp__4388__auto_____1, i, xs__4940__auto__, temp__4388__auto__), iter__528.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__529___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
																	}(&cljs_core.AFn{}, s__527___1, i, xs__4940__auto__, temp__4388__auto__), nil, nil})
																})
															}(&cljs_core.AFn{}, s__527___1, i, xs__4940__auto__, temp__4388__auto__)
															var fs__936__auto__ = cljs_core.Seq.Arity1IQ(iterys__935__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq))
															_, _ = iterys__935__auto__, fs__936__auto__
															if cljs_core.Truth_(fs__936__auto__) {
																return cljs_core.Concat.X_invoke_Arity2(fs__936__auto__, iter__526.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__527___1)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
															} else {
																s__527___1 = cljs_core.Rest.Arity1IQ(s__527___1)
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
				_ = iter__939__auto__
				return iter__939__auto__.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})).(*cljs_core.CljsCoreLazySeq)
			}()).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[1 1] [1 2] [1 3] [2 1] [2 2] [2 3]] (map (fn* [p1__57#] (p1__57#)) (for [i [1 2] j [1 2 3]] (fn [] [i j]))))").(string)}, ``)}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(integer? 0)").(string)}, ``)}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(float64(42)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(integer? 42)").(string)}, ``)}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(float64(-42)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(integer? -42)").(string)}, ``)}))
			}
			if !(cljs_core.Integer_QMARK_.Arity1IB("")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (integer? \"\"))").(string)}, ``)}))
			}
			if !(cljs_core.Integer_QMARK_.Arity1IB(1.0E308)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (integer? 1.0E308))").(string)}, ``)}))
			}
			if !(cljs_core.Integer_QMARK_.Arity1IB(js.Infinity)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (integer? js/Infinity))").(string)}, ``)}))
			}
			if !(cljs_core.Integer_QMARK_.Arity1IB((-js.Infinity))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (integer? (- js/Infinity)))").(string)}, ``)}))
			}
			if !(cljs_core.Integer_QMARK_.Arity1IB(js.NaN)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (integer? js/NaN))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(42), float64(cljs_core.Int32_(42.5))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 42 (int 42.5))").(string)}, ``)}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(float64(cljs_core.Int32_(42.5))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(integer? (int 42.5))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Long.X_invoke_Arity1(42.5).(float64)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 42 (long 42.5))").(string)}, ``)}))
			}
			if cljs_core.Integer_QMARK_.Arity1IB(cljs_core.Long.X_invoke_Arity1(42.5).(float64)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(integer? (long 42.5))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), float64(cljs_core.Int32_(-1.5))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (int -1.5))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-9), cljs_core.Long.X_invoke_Arity1(-9.8).(float64)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -9 (long -9.8))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (:b {:a 1, :b 2}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}).X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 ((quote b) (quote {:a 1, b 2})))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}).X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 ({:a 1, :b 2} :b))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), float64(1), float64(2), float64(2)}, nil}).X_invoke_Arity1(float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 ({1 1, 2 2} 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(1)}, nil}), float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (:a {:b 1} 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (:a {} 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(1)}, nil}).X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 ({:b 1} :a 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.CljsCorePersistentArrayMap_EMPTY.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 ({} :a 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (:a {}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1("")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (:a \"\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity2("", float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (:a \"\" 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}).X_invoke_Arity1(float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (#{1 3 2} 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Apply.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), float64(2)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (apply :a (quote [{:a 1, a 2}])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Apply.X_invoke_Arity2((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (apply (quote a) (quote [{a 1, :b 2}])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Apply.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (apply {:a 1} [:a]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Apply.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (apply {:a 1} [:b 2]))").(string)}, ``)}))
			}
			if cljs_core.Nil_(cljs_core.Namespace.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "/", Str: "/", X_hash: float64(-1371932971), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (namespace (quote /)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("/", cljs_core.Name.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "/", Str: "/", X_hash: float64(-1371932971), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"/\" (name (quote /)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("keyword", cljs_core.Name.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "keyword", Fqn: "keyword", X_hash: float64(811389747)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"keyword\" (name :keyword))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(":hello", strings.Join([]string{cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hello", Fqn: "hello", X_hash: float64(-245025397)})).(string)}, ``)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \":hello\" (str :hello))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("hello", strings.Join([]string{cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "hello", Str: "hello", X_hash: float64(1395506130), X_meta: nil})).(string)}, ``)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"hello\" (str (quote hello)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("hello:world", strings.Join([]string{cljs_core.Str.X_invoke_Arity1("hello").(string), cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "world", Fqn: "world", X_hash: float64(-418292623)})).(string)}, ``)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"hello:world\" (str \"hello\" :world))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(":helloworld", strings.Join([]string{cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hello", Fqn: "hello", X_hash: float64(-245025397)})).(string), cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "world", Str: "world", X_hash: float64(1222238904), X_meta: nil})).(string)}, ``)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \":helloworld\" (str :hello (quote world)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), cljs_core.Symbol.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote a) (symbol (quote a)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Keyword.X_invoke_Arity1("a")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :a (keyword \"a\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Keyword.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :a (keyword (quote a)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), cljs_core.Keyword.X_invoke_Arity2((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil})).(*cljs_core.CljsCoreKeyword)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :a/b (keyword (quote a) (quote b)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Keyword.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :a (keyword :a))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), cljs_core.Get.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), float64(4), float64(5)}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:a :b} (get {[1 2 3] {:a :b}, 4 5} [1 2 3]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil}), float64(0))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :a (nth [:a :b :c :d] 0))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil}), 0.1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :a (nth [:a :b :c :d] 0.1))").(string)}, ``)}))
			}
			if !(cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), nil}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (= {:a :b, :c nil} {:a :b, :d nil}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3)), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (list 3 2 1) [3 2 1])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Seq.Arity1IQ([]interface{}{float64(3), float64(2), float64(1)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3 2 1] (seq (array 3 2 1)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(9), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Next.Arity1IQ(cljs_core.Seq.Arity1IQ([]interface{}{float64(1), float64(2), float64(3), float64(4)})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 9 (reduce + (next (seq (array 1 2 3 4)))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Rest.Arity1IQ(nil)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= () (rest nil))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Seq.Arity1IQ([]interface{}{})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (seq (array)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Seq.Arity1IQ("")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (seq \"\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Seq.Arity1IQ(cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (seq []))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Seq.Arity1IQ(cljs_core.CljsCorePersistentArrayMap_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (seq {}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Rest.Arity1IQ(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= () (rest ()))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Rest.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= () (rest [1]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Rest.Arity1IQ([]interface{}{float64(1)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= () (rest (array 1)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"x", "y"}, nil}), cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"x", "y"}, nil})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {\"x\" \"y\"} (meta []))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), cljs_core.Dissoc.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:a :b} (dissoc {:a :b, :c :d} :c))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("\"asdf\" \"asdf\"", cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"asdf", "asdf"})).(string)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"\\\"asdf\\\" \\\"asdf\\\"\" (pr-str \"asdf\" \"asdf\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("[1 true {:a 2, :b #\"x\\\"y\"} #js [3 4]]", cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), true, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&js.RegExp{Pattern: `x\"y`, Flags: ``})}, nil}), []interface{}{float64(3), float64(4)}}, nil})})).(string)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"[1 true {:a 2, :b #\\\"x\\\\\\\"y\\\"} #js [3 4]]\" (pr-str [1 true {:a 2, :b #\"x\\\"y\"} (array 3 4)]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("\"asdf\"\n", cljs_core.Prn_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"asdf"})).(string)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"\\\"asdf\\\"\\n\" (prn-str \"asdf\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("[1 true {:a 2, :b 42} #js [3 4]]\n", cljs_core.Prn_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), true, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(42)}, nil}), []interface{}{float64(3), float64(4)}}, nil})})).(string)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"[1 true {:a 2, :b 42} #js [3 4]]\\n\" (prn-str [1 true {:a 2, :b 42} (array 3 4)]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("asdf", cljs_core.Print_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"asdf"})).(string)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"asdf\" (print-str \"asdf\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("asdf\n", cljs_core.Println_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"asdf"})).(string)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"asdf\\n\" (println-str \"asdf\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("", cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{})).(string)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"\" (pr-str))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("\n", cljs_core.Prn_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{})).(string)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"\\n\" (prn-str))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("12", func() string {
				var sb__1140__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1140__auto__
				{
					var _STAR_print_fn_STAR_537_964 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_537_964
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_537_964

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__965 *cljs_core.AFn, _STAR_print_fn_STAR_537_964 interface{}, sb__1140__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__965, 1, func(x__1141__auto__ interface{}) interface{} {
									return sb__1140__auto__.Append(x__1141__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_537_964, sb__1140__auto__)

							cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)}))
							cljs_core.Print.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)}))
						}
					}()
				}
				return strings.Join([]string{cljs_core.Str.X_invoke_Arity1(sb__1140__auto__).(string)}, ``)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"12\" (with-out-str (print 1) (print 2)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("12", func() string {
				var sb__1140__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1140__auto__
				{
					var _STAR_print_fn_STAR_538_966 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_538_966
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_538_966

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__967 *cljs_core.AFn, _STAR_print_fn_STAR_538_966 interface{}, sb__1140__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__967, 1, func(x__1141__auto__ interface{}) interface{} {
									return sb__1140__auto__.Append(x__1141__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_538_966, sb__1140__auto__)

							cljs_core.X_STAR_print_fn_STAR_.X_invoke_Arity1(float64(1))
							cljs_core.X_STAR_print_fn_STAR_.X_invoke_Arity1(float64(2))
						}
					}()
				}
				return strings.Join([]string{cljs_core.Str.X_invoke_Arity1(sb__1140__auto__).(string)}, ``)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"12\" (with-out-str (*print-fn* 1) (*print-fn* 2)))").(string)}, ``)}))
			}
			if !(cljs_core.X_EQ_.Arity2IIB("one", "two")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (= \"one\" \"two\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Count.X_invoke_Arity1("abc").(float64)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (count \"abc\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Count.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(4)}).(float64)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (count (array 1 2 3 4)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("c", cljs_core.Nth.X_invoke_Arity2("abc", float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"c\" (nth \"abc\" 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("quux", cljs_core.Nth.X_invoke_Arity3("abc", float64(3), "quux")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"quux\" (nth \"abc\" 3 \"quux\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Nth.X_invoke_Arity2([]interface{}{float64(1), float64(2), float64(3), float64(4)}, float64(0))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (nth (array 1 2 3 4) 0))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("val", cljs_core.Nth.X_invoke_Arity3([]interface{}{float64(1), float64(2), float64(3), float64(4)}, float64(4), "val")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"val\" (nth (array 1 2 3 4) 4 \"val\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("b", cljs_core.Get.X_invoke_Arity2("abc", float64(1))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"b\" (get \"abc\" 1))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("harriet", cljs_core.Get.X_invoke_Arity3("abcd", float64(4), "harriet")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"harriet\" (get \"abcd\" 4 \"harriet\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity2([]interface{}{float64(1), float64(2), float64(3), float64(4)}, float64(3))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get (array 1 2 3 4) 3))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("zot", cljs_core.Get.X_invoke_Arity3([]interface{}{float64(1), float64(2), float64(3), float64(4)}, float64(4), "zot")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"zot\" (get (array 1 2 3 4) 4 \"zot\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(10), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, []interface{}{float64(1), float64(2), float64(3), float64(4)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 10 (reduce + (array 1 2 3 4)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(20), cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), []interface{}{float64(1), float64(2), float64(3), float64(4)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 20 (reduce + 10 (array 1 2 3 4)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("cabd", func() interface{} {
				var jumble = func(G__968 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__968, 2, func(a interface{}, b interface{}) interface{} {
						return strings.Join([]string{cljs_core.Str.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Reverse.X_invoke_Arity1(strings.Join([]string{cljs_core.Str.X_invoke_Arity1(a).(string)}, ``)))).(string), cljs_core.Str.X_invoke_Arity1(b).(string)}, ``)
					})
				}(&cljs_core.AFn{})
				_ = jumble
				return cljs_core.Reduce.X_invoke_Arity2(jumble, "abcd")
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"cabd\" (let [jumble (fn [a b] (str (apply str (reverse (str a))) b))] (reduce jumble \"abcd\")))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("cafrogbd", func() interface{} {
				var jumble = func(G__969 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__969, 2, func(a interface{}, b interface{}) interface{} {
						return strings.Join([]string{cljs_core.Str.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Reverse.X_invoke_Arity1(strings.Join([]string{cljs_core.Str.X_invoke_Arity1(a).(string)}, ``)))).(string), cljs_core.Str.X_invoke_Arity1(b).(string)}, ``)
					})
				}(&cljs_core.AFn{})
				_ = jumble
				return cljs_core.Reduce.X_invoke_Arity3(jumble, "frog", "abcd")
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"cafrogbd\" (let [jumble (fn [a b] (str (apply str (reverse (str a))) b))] (reduce jumble \"frog\" \"abcd\")))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0), float64(1), float64(0), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) & cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(0)) & cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(1)) & cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(42)) & cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(41)) & cljs_core.Int32_(float64(1))))}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [0 0 1 0 1] [(bit-and 1 0) (bit-and 0 0) (bit-and 1 1) (bit-and 42 1) (bit-and 41 1)])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(0), float64(1), float64(43), float64(41)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) | cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(0)) | cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(1)) | cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(42)) | cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(41)) | cljs_core.Int32_(float64(1))))}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 0 1 43 41] [(bit-or 1 0) (bit-or 0 0) (bit-or 1 1) (bit-or 42 1) (bit-or 41 1)])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(0), float64(0), float64(42), float64(40)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) &^ cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(0)) &^ cljs_core.Int32_(float64(0)))), float64((cljs_core.Int32_(float64(1)) &^ cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(42)) &^ cljs_core.Int32_(float64(1)))), float64((cljs_core.Int32_(float64(41)) &^ cljs_core.Int32_(float64(1))))}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 0 0 42 40] [(bit-and-not 1 0) (bit-and-not 0 0) (bit-and-not 1 1) (bit-and-not 42 1) (bit-and-not 41 1)])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2), float64(968), float64(16649), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) &^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(0))))), float64((cljs_core.Int32_(float64(2)) &^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(0))))), float64((cljs_core.Int32_(float64(1000)) &^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(5))))), float64((cljs_core.Int32_(float64(16713)) &^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(6))))), float64((cljs_core.Int32_(float64(1024)) &^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(10)))))}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [0 2 968 16649 0] [(bit-clear 1 0) (bit-clear 2 0) (bit-clear 1000 5) (bit-clear 16713 6) (bit-clear 1024 10)])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0), float64(992), float64(18761), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) ^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(0))))), float64((cljs_core.Int32_(float64(2)) ^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(1))))), float64((cljs_core.Int32_(float64(1000)) ^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(3))))), float64((cljs_core.Int32_(float64(16713)) ^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(11))))), float64((cljs_core.Int32_(float64(1024)) ^ (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(10)))))}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [0 0 992 18761 0] [(bit-flip 1 0) (bit-flip 2 1) (bit-flip 1000 3) (bit-flip 16713 11) (bit-flip 1024 10)])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-2), float64(-3), float64(999), float64(-16714), float64(-1025)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((^cljs_core.Int32_(float64(1)))), float64((^cljs_core.Int32_(float64(2)))), float64((^cljs_core.Int32_(float64(-1000)))), float64((^cljs_core.Int32_(float64(16713)))), float64((^cljs_core.Int32_(float64(1024))))}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [-2 -3 999 -16714 -1025] [(bit-not 1) (bit-not 2) (bit-not -1000) (bit-not 16713) (bit-not 1024)])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(1000), float64(18761), float64(1024)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64((cljs_core.Int32_(float64(1)) | (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(0))))), float64((cljs_core.Int32_(float64(2)) | (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(1))))), float64((cljs_core.Int32_(float64(1000)) | (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(3))))), float64((cljs_core.Int32_(float64(16713)) | (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(11))))), float64((cljs_core.Int32_(float64(1024)) | (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(10)))))}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 1000 18761 1024] [(bit-set 1 0) (bit-set 2 1) (bit-set 1000 3) (bit-set 16713 11) (bit-set 1024 10)])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true, true, false, true}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(cljs_core.Int32_(float64(1)) & (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(0)))) != 0, (cljs_core.Int32_(float64(2)) & (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(1)))) != 0, (cljs_core.Int32_(float64(1000)) & (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(3)))) != 0, (cljs_core.Int32_(float64(16713)) & (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(11)))) != 0, (cljs_core.Int32_(float64(1024)) & (cljs_core.Int32_(1) << cljs_core.UInt32_(float64(10)))) != 0}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [true true true false true] [(bit-test 1 0) (bit-test 2 1) (bit-test 1000 3) (bit-test 16713 11) (bit-test 1024 10)])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, true, false, false, false}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true == true, false == true, false == false, true == false, js.Undefined == true, js.Undefined == false}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [true false true false false false] [(true? true) (true? false) (false? false) (false? true) (true? js/undefined) (false? js/undefined)])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, nil)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0 (apply + nil))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.CljsCoreList_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0 (apply + (list)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (apply + (list 1)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Apply.X_invoke_Arity3(cljs_core.X_PLUS_, float64(1), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (apply + 1 (list 2)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(7), cljs_core.Apply.X_invoke_Arity4(cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(4)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 7 (apply + 1 2 (list 4)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(15), cljs_core.Apply.X_invoke_Arity5(cljs_core.X_PLUS_, float64(1), float64(2), float64(4), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 15 (apply + 1 2 4 (list 8)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(31), cljs_core.Apply.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(1), float64(2), float64(4), float64(8), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(16))}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 31 (apply + 1 2 4 8 (list 16)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(63), cljs_core.Apply.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(1), float64(2), float64(4), float64(8), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(16), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(32))}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 63 (apply + 1 2 4 8 16 (list 32)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(127), cljs_core.Apply.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(1), float64(2), float64(4), float64(8), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(16), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(64))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(32))}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 127 (apply + 1 2 4 8 16 (list 32 64)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(4950), cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Take.X_invoke_Arity2(float64(100), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4950 (apply + (take 100 (iterate inc 0))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Apply.X_invoke_Arity2(cljs_core.List, cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= () (apply list []))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), cljs_core.Apply.X_invoke_Arity2(cljs_core.List, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3] (apply list [1 2 3]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Apply.X_invoke_Arity2(cljs_core.Apply, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 6 (apply apply [+ [1 2 3]]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Apply.X_invoke_Arity2(func(G__970 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__970, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return ((cljs_core.Nth.X_invoke_Arity2(args, float64(0)).(float64) + cljs_core.Nth.X_invoke_Arity2(args, float64(1)).(float64)) + cljs_core.Nth.X_invoke_Arity2(args, float64(2)).(float64))
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (apply (fn [& args] (+ (nth args 0) (nth args 1) (nth args 2))) (iterate inc 0)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__971 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__971, 0, func(m__ ...interface{}) interface{} {
					var m = cljs_core.Seq.Arity1IQ(m__[0])
					_ = m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [0 1 2 3 4] (take 5 (apply (fn [& m] m) (iterate inc 0))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__972 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__972, 1, func(x_m__ ...interface{}) interface{} {
					var x = x_m__[0]
					var m = cljs_core.Seq.Arity1IQ(x_m__[1])
					_, _ = x, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3 4 5] (take 5 (apply (fn [x & m] m) (iterate inc 0))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__973 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__973, 2, func(x_y_m__ ...interface{}) interface{} {
					var x = x_y_m__[0]
					var y = x_y_m__[1]
					var m = cljs_core.Seq.Arity1IQ(x_y_m__[2])
					_, _, _ = x, y, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 3 4 5 6] (take 5 (apply (fn [x y & m] m) (iterate inc 0))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5), float64(6), float64(7)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__974 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__974, 3, func(x_y_z_m__ ...interface{}) interface{} {
					var x = x_y_z_m__[0]
					var y = x_y_z_m__[1]
					var z = x_y_z_m__[2]
					var m = cljs_core.Seq.Arity1IQ(x_y_z_m__[3])
					_, _, _, _ = x, y, z, m
					return m
				})
			}(&cljs_core.AFn{}), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(0)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3 4 5 6 7] (take 5 (apply (fn [x y z & m] m) (iterate inc 0))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6), float64(7), float64(8)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__975 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__975, 4, func(x_y_z_a_m__ ...interface{}) interface{} {
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [4 5 6 7 8] (take 5 (apply (fn [x y z a & m] m) (iterate inc 0))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(func(G__976 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__976, 5, func(x_y_z_a_b_m__ ...interface{}) interface{} {
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [5 6 7 8 9] (take 5 (apply (fn [x y z a b & m] m) (iterate inc 0))))").(string)}, ``)}))
			}
			{
				var single_arity_non_variadic_977 = func(G__982 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__982, 3, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{})
				var multiple_arity_non_variadic_978 = func(G__983 *cljs_core.AFn, single_arity_non_variadic_977 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__983, 3, func(x interface{}) interface{} {
						return x
					}, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{y, x}, nil})
					}, func(x interface{}, y interface{}, z interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{z, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_977)
				var single_arity_variadic_fixedargs_979 = func(G__984 *cljs_core.AFn, single_arity_non_variadic_977 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_978 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__984, 2, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{more, y, x}, nil})
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_977, multiple_arity_non_variadic_978)
				var single_arity_variadic_nofixedargs_980 = func(G__985 *cljs_core.AFn, single_arity_non_variadic_977 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_978 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_979 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__985, 0, func(more__ ...interface{}) interface{} {
						var more = cljs_core.Seq.Arity1IQ(more__[0])
						_ = more
						return more
					})
				}(&cljs_core.AFn{}, single_arity_non_variadic_977, multiple_arity_non_variadic_978, single_arity_variadic_fixedargs_979)
				var multiple_arity_variadic_981 = func(G__986 *cljs_core.AFn, single_arity_non_variadic_977 cljs_core.CljsCoreIFn, multiple_arity_non_variadic_978 cljs_core.CljsCoreIFn, single_arity_variadic_fixedargs_979 cljs_core.CljsCoreIFn, single_arity_variadic_nofixedargs_980 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(G__986, 2, func(x interface{}) interface{} {
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
				}(&cljs_core.AFn{}, single_arity_non_variadic_977, multiple_arity_non_variadic_978, single_arity_variadic_fixedargs_979, single_arity_variadic_nofixedargs_980)
				_, _, _, _, _ = single_arity_non_variadic_977, multiple_arity_non_variadic_978, single_arity_variadic_fixedargs_979, single_arity_variadic_nofixedargs_980, multiple_arity_variadic_981
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_non_variadic_977, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3 2 1] (apply single-arity-non-variadic [1 2 3]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_non_variadic_977, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3 2 1] (apply single-arity-non-variadic 1 [2 3]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_non_variadic_977, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3 2 1] (apply single-arity-non-variadic 1 2 [3]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_978, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 42 (apply multiple-arity-non-variadic [42]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_978, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 1] (apply multiple-arity-non-variadic [1 2]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_978, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 1] (apply multiple-arity-non-variadic 1 [2]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_non_variadic_978, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3 2 1] (apply multiple-arity-non-variadic [1 2 3]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_non_variadic_978, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3 2 1] (apply multiple-arity-non-variadic 1 [2 3]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_non_variadic_978, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3 2 1] (apply multiple-arity-non-variadic 1 2 [3]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_979, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs [1 2 3 4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_fixedargs_979, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 [2 3 4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_fixedargs_979, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 [3 4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_fixedargs_979, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 [4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_fixedargs_979, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[3 4 5] 2 1] (apply single-arity-variadic-fixedargs 1 2 3 4 [5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_979, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3 4 5] (take 3 (first (apply single-arity-variadic-fixedargs (iterate inc 1)))))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_fixedargs_979, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 1] (rest (apply single-arity-variadic-fixedargs (iterate inc 1))))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_980, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs [1 2 3 4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity3(single_arity_variadic_nofixedargs_980, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 [2 3 4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity4(single_arity_variadic_nofixedargs_980, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 [3 4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_Arity5(single_arity_variadic_nofixedargs_980, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 [4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(single_arity_variadic_nofixedargs_980, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3 4 5] (apply single-arity-variadic-nofixedargs 1 2 3 4 [5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(5), cljs_core.Apply.X_invoke_Arity2(single_arity_variadic_nofixedargs_980, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3 4 5] (take 5 (apply single-arity-variadic-nofixedargs (iterate inc 1))))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_981, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 42 (apply multiple-arity-variadic [42]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_981, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 1] (apply multiple-arity-variadic [1 2]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_981, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 1] (apply multiple-arity-variadic 1 [2]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_981, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[3 4 5] 2 1] (apply multiple-arity-variadic [1 2 3 4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity3(multiple_arity_variadic_981, float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 [2 3 4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity4(multiple_arity_variadic_981, float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 [3 4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_Arity5(multiple_arity_variadic_981, float64(1), float64(2), float64(3), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 [4 5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), float64(2), float64(1)}, nil}), cljs_core.Apply.X_invoke_ArityVariadic(multiple_arity_variadic_981, float64(1), float64(2), float64(3), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5)}, nil})}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[3 4 5] 2 1] (apply multiple-arity-variadic 1 2 3 4 [5]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil}), cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.First.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_981, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3 4 5] (take 3 (first (apply multiple-arity-variadic (iterate inc 1)))))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), cljs_core.Rest.Arity1IQ(cljs_core.Apply.X_invoke_Arity2(multiple_arity_variadic_981, cljs_core.Iterate.X_invoke_Arity2(cljs_core.Inc, float64(1)).(*cljs_core.CljsCoreCons)))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 1] (rest (apply multiple-arity-variadic (iterate inc 1))))").(string)}, ``)}))
				}
			}
			{
				var f1_987 = func(f1 *cljs_core.AFn) *cljs_core.AFn {
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
				var f2_988 = func(f2 *cljs_core.AFn, f1_987 cljs_core.CljsCoreIFn) *cljs_core.AFn {
					return cljs_core.Fn(f2, 2, func(x interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					}, func(x_y_more__ ...interface{}) interface{} {
						var x = x_y_more__[0]
						var y = x_y_more__[1]
						var more = cljs_core.Seq.Arity1IQ(x_y_more__[2])
						_, _, _ = x, y, more
						return cljs_core.Apply.X_invoke_Arity3(f1_987, y, more)
					})
				}(&cljs_core.AFn{}, f1_987)
				_, _ = f1_987, f2_988
				if cljs_core.X_EQ_.Arity2IIB(float64(1), f2_988.X_invoke_Arity2(float64(1), float64(2))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (f2 1 2))").(string)}, ``)}))
				}
			}
			{
				var f_989 = func(G__990 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__990, 1, func() interface{} {
						return nil
					}, func(a_more__ ...interface{}) interface{} {
						var a = a_more__[0]
						var more = cljs_core.Seq.Arity1IQ(a_more__[1])
						_, _ = a, more
						return more
					})
				}(&cljs_core.AFn{})
				_ = f_989
				if cljs_core.Nil_(f_989.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (f :foo))").(string)}, ``)}))
				}
			}
			if cljs_core.Nil_(cljs_core.Array_seq.X_invoke_Arity2([]interface{}{float64(1)}, float64(1))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (array-seq (array 1) 1))").(string)}, ``)}))
			}
			{
				var f_991 = func(G__994 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__994, 1, func(x interface{}) interface{} {
						return (x.(float64) * float64(2))
					})
				}(&cljs_core.AFn{})
				var m_992 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar"}, nil})
				var mf_993 = cljs_core.With_meta.X_invoke_Arity2(f_991, m_992)
				_, _, _ = f_991, m_992, mf_993
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(f_991)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (meta f))").(string)}, ``)}))
				}
				if cljs_core.Fn_QMARK_.Arity1IB(mf_993) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(fn? mf)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__539 = float64(2)
					_ = G__539
					return mf_993.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__539)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (mf 2))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Apply.X_invoke_Arity2(mf_993, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (apply mf [2]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(mf_993), m_992) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (meta mf) m)").(string)}, ``)}))
				}
			}
			{
				var a_995 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_995
				if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Deref.X_invoke_Arity1(a_995)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0 (deref a))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity2(a_995, cljs_core.Inc)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (swap! a inc))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(false, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_995, float64(0), float64(42)))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= false (compare-and-set! a 0 42))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(true, cljs_core.Truth_(cljs_core.Compare_and_set_BANG_.X_invoke_Arity3(a_995, float64(1), float64(7)))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= true (compare-and-set! a 1 7))").(string)}, ``)}))
				}
				if cljs_core.Nil_(cljs_core.Meta.X_invoke_Arity1(a_995)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (meta a))").(string)}, ``)}))
				}
				if cljs_core.Nil_(cljs_core.Get_validator.X_invoke_Arity1(a_995)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get-validator a))").(string)}, ``)}))
				}
			}
			{
				var a_996 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				_ = a_996
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Swap_BANG_.X_invoke_Arity3(a_996, cljs_core.X_PLUS_, float64(1))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (swap! a + 1))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Swap_BANG_.X_invoke_Arity4(a_996, cljs_core.X_PLUS_, float64(1), float64(2))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (swap! a + 1 2))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(10), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_996, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 10 (swap! a + 1 2 3))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(20), cljs_core.Swap_BANG_.X_invoke_ArityVariadic(a_996, cljs_core.X_PLUS_, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(4)}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 20 (swap! a + 1 2 3 4))").(string)}, ``)}))
				}
			}
			{
				var a_997 = cljs_core.Atom.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "validator", Fqn: "validator", X_hash: float64(-1966190681)}), cljs_core.Coll_QMARK_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "meta", Fqn: "meta", X_hash: float64(1499536964)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})})).(*cljs_core.CljsCoreAtom)
				_ = a_997
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Coll_QMARK_, cljs_core.Get_validator.X_invoke_Arity1(a_997)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= coll? (get-validator a))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_997)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:a 1} (meta a))").(string)}, ``)}))
				}
				cljs_core.Alter_meta_BANG_.X_invoke_ArityVariadic(a_997, cljs_core.Assoc, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}))
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), cljs_core.Meta.X_invoke_Arity1(a_997)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:a 1, :b 2} (meta a))").(string)}, ``)}))
				}
			}
			if cljs_core.Nil_(cljs_core.Empty.X_invoke_Arity1(nil)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (empty nil))").(string)}, ``)}))
			}
			{
				var e_lazy_seq_998 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCoreLazySeq{nil, func(G__999 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__999, 0, func() interface{} {
						return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons)
					})
				}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_lazy_seq_998
				if cljs_core.Seq_QMARK_.Arity1IB(e_lazy_seq_998) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(seq? e-lazy-seq)").(string)}, ``)}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_lazy_seq_998) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(empty? e-lazy-seq)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_lazy_seq_998)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:b :c} (meta e-lazy-seq))").(string)}, ``)}))
				}
			}
			{
				var e_list_1000 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_list_1000
				if cljs_core.Seq_QMARK_.Arity1IB(e_list_1000) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(seq? e-list)").(string)}, ``)}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_list_1000) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(empty? e-list)").(string)}, ``)}))
				}
			}
			{
				var e_elist_1001 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_elist_1001
				if cljs_core.Seq_QMARK_.Arity1IB(e_elist_1001) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(seq? e-elist)").(string)}, ``)}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_elist_1001) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(empty? e-elist)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), cljs_core.Get.X_invoke_Arity2(cljs_core.Meta.X_invoke_Arity1(e_elist_1001), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :c (get (meta e-elist) :b))").(string)}, ``)}))
				}
			}
			{
				var e_cons_1002 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil).(*cljs_core.CljsCoreCons), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_cons_1002
				if cljs_core.Seq_QMARK_.Arity1IB(e_cons_1002) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(seq? e-cons)").(string)}, ``)}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_cons_1002) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(empty? e-cons)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_cons_1002)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:b :c} (meta e-cons))").(string)}, ``)}))
				}
			}
			{
				var e_vec_1003 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_vec_1003
				if cljs_core.Vector_QMARK_.Arity1IB(e_vec_1003) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(vector? e-vec)").(string)}, ``)}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_vec_1003) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(empty? e-vec)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_vec_1003)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:b :c} (meta e-vec))").(string)}, ``)}))
				}
			}
			{
				var e_omap_1004 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_omap_1004
				if cljs_core.Map_QMARK_.Arity1IB(e_omap_1004) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(map? e-omap)").(string)}, ``)}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_omap_1004) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(empty? e-omap)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_omap_1004)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:b :c} (meta e-omap))").(string)}, ``)}))
				}
			}
			{
				var e_hmap_1005 = cljs_core.Empty.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)})}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})))
				_ = e_hmap_1005
				if cljs_core.Map_QMARK_.Arity1IB(e_hmap_1005) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(map? e-hmap)").(string)}, ``)}))
				}
				if cljs_core.Empty_QMARK_.Arity1IB(e_hmap_1005) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(empty? e-hmap)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), cljs_core.Meta.X_invoke_Arity1(e_hmap_1005)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:b :c} (meta e-hmap))").(string)}, ``)}))
				}
			}
			{
				var a_1006 = cljs_core.Atom.X_invoke_Arity1(nil).(*cljs_core.CljsCoreAtom)
				_ = a_1006
				if cljs_core.X_EQ_.Arity2IIB(float64(1), float64(1)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (try 1))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1007 interface{}) {
					defer func() {
						if e540 := recover(); e540 != nil {
							if cljs_core.Value_(e540).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e540
									_ = e
									return__1007 = float64(2)
								}
							} else {
								panic(e540)

							}
						}
					}()
					{
						panic((&js.Error{}))
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try 1 (throw (js/Error.)) (catch js/Error e 2)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1008 interface{}) {
					defer func() {
						if e541 := recover(); e541 != nil {
							if cljs_core.Value_(e541).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e541
									_ = e
									return__1008 = float64(2)
								}
							} else {
								panic(e541)

							}
						}
					}()
					{
						panic((&js.Error{}))
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try 1 (throw (js/Error.)) (catch js/Error e 1 2)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1009 interface{}) {
					defer func() {
						if e542 := recover(); e542 != nil {
							if cljs_core.Value_(e542).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e542
									_ = e
									return__1009 = float64(2)
								}
							} else {
								{
									var e = e542
									_ = e
									return__1009 = float64(3)
								}

							}
						}
					}()
					{
						panic((&js.Error{}))
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try 1 (throw (js/Error.)) (catch js/Error e 2) (catch :default e 3)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() (return__1010 interface{}) {
					defer func() {
						if e543 := recover(); e543 != nil {
							if cljs_core.Value_(e543).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e543
									_ = e
									return__1010 = float64(2)
								}
							} else {
								{
									var e = e543
									_ = e
									return__1010 = float64(3)
								}

							}
						}
					}()
					{
						panic(true)
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (try 1 (throw true) (catch js/Error e 2) (catch :default e 3)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1011 interface{}) {
					defer func() {
						if e544 := recover(); e544 != nil {
							if cljs_core.Value_(e544).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e544
									_ = e
									return__1011 = float64(3)
								}
							} else {
								{
									var e = e544
									_ = e
									return__1011 = e
								}

							}
						}
					}()
					{
						panic(float64(2))
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try 1 (throw 2) (catch js/Error e 3) (catch :default e e)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), func() float64 {
					defer func() {
						cljs_core.Reset_BANG_.X_invoke_Arity2(a_1006, float64(42))
					}()
					{
						return float64(1)
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (try 1 (finally (reset! a 42))))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Deref.X_invoke_Arity1(a_1006)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 42 (deref a))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), cljs_core.Seq_(cljs_core.Nthnext.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3] (nthnext [1 2 3] 2))").(string)}, ``)}))
			}
			{
				var v_1012 = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})
				_ = v_1012
				if cljs_core.X_EQ_.Arity2IIB(v_1012, func() *cljs_core.CljsCoreLazySeq {
					var iter__939__auto__ = func(iter__545 *cljs_core.AFn, v_1012 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__545, 1, func(s__546 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1013 *cljs_core.AFn, v_1012 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1013, 0, func() interface{} {
									{
										var s__546___1 interface{} = s__546
										_ = s__546___1
										for {
											{
												var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__546___1)
												_ = temp__4388__auto__
												if cljs_core.Truth_(temp__4388__auto__) {
													{
														var s__546___2 = temp__4388__auto__
														_ = s__546___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__546___2) {
															{
																var c__937__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__546___2)
																var size__938__auto__ = cljs_core.Count.X_invoke_Arity1(c__937__auto__).(float64)
																var b__548 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__938__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__937__auto__, size__938__auto__, b__548
																if func() bool {
																	var i__547 = float64(0)
																	_ = i__547
																	for {
																		if i__547 < size__938__auto__ {
																			{
																				var e = cljs_core.Decorate_(c__937__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__547)
																				_ = e
																				cljs_core.Chunk_append.X_invoke_Arity2(b__548, e)
																				i__547 = (i__547 + float64(1))
																				continue
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__548), iter__545.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__546___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__548), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__546___2)
																_ = e
																return cljs_core.Cons.X_invoke_Arity2(e, iter__545.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__546___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_1012), nil, nil})
						})
					}(&cljs_core.AFn{}, v_1012)
					_ = iter__939__auto__
					return iter__939__auto__.X_invoke_Arity1(v_1012).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= v (for [e v] e))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(9)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__939__auto__ = func(iter__551 *cljs_core.AFn, v_1012 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__551, 1, func(s__552 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1014 *cljs_core.AFn, v_1012 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1014, 0, func() interface{} {
									{
										var s__552___1 interface{} = s__552
										_ = s__552___1
										for {
											{
												var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__552___1)
												_ = temp__4388__auto__
												if cljs_core.Truth_(temp__4388__auto__) {
													{
														var s__552___2 = temp__4388__auto__
														_ = s__552___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__552___2) {
															{
																var c__937__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__552___2)
																var size__938__auto__ = cljs_core.Count.X_invoke_Arity1(c__937__auto__).(float64)
																var b__554 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__938__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__937__auto__, size__938__auto__, b__554
																if func() bool {
																	var i__553 = float64(0)
																	_ = i__553
																	for {
																		if i__553 < size__938__auto__ {
																			{
																				var e = cljs_core.Decorate_(c__937__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__553)
																				_ = e
																				{
																					var m = (e.(float64) * e.(float64))
																					_ = m
																					cljs_core.Chunk_append.X_invoke_Arity2(b__554, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																					i__553 = (i__553 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__554), iter__551.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__552___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__554), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__552___2)
																_ = e
																{
																	var m = (e.(float64) * e.(float64))
																	_ = m
																	return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__551.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__552___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_1012), nil, nil})
						})
					}(&cljs_core.AFn{}, v_1012)
					_ = iter__939__auto__
					return iter__939__auto__.X_invoke_Arity1(v_1012).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[1 1] [2 4] [3 9]] (for [e v :let [m (* e e)]] [e m]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__939__auto__ = func(iter__557 *cljs_core.AFn, v_1012 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__557, 1, func(s__558 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1015 *cljs_core.AFn, v_1012 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1015, 0, func() interface{} {
									{
										var s__558___1 interface{} = s__558
										_ = s__558___1
										for {
											{
												var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__558___1)
												_ = temp__4388__auto__
												if cljs_core.Truth_(temp__4388__auto__) {
													{
														var s__558___2 = temp__4388__auto__
														_ = s__558___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__558___2) {
															{
																var c__937__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__558___2)
																var size__938__auto__ = cljs_core.Count.X_invoke_Arity1(c__937__auto__).(float64)
																var b__560 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__938__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__937__auto__, size__938__auto__, b__560
																if cljs_core.Truth_(func() interface{} {
																	var i__559 = float64(0)
																	_ = i__559
																	for {
																		if i__559 < size__938__auto__ {
																			{
																				var e = cljs_core.Decorate_(c__937__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__559)
																				_ = e
																				if e.(float64) < float64(3) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__560, e)
																					i__559 = (i__559 + float64(1))
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
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__560), iter__557.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__558___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__560), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__558___2)
																_ = e
																if e.(float64) < float64(3) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__557.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__558___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_1012), nil, nil})
						})
					}(&cljs_core.AFn{}, v_1012)
					_ = iter__939__auto__
					return iter__939__auto__.X_invoke_Arity1(v_1012).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2] (for [e v :while (< e 3)] e))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__939__auto__ = func(iter__563 *cljs_core.AFn, v_1012 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__563, 1, func(s__564 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1016 *cljs_core.AFn, v_1012 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1016, 0, func() interface{} {
									{
										var s__564___1 interface{} = s__564
										_ = s__564___1
										for {
											{
												var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__564___1)
												_ = temp__4388__auto__
												if cljs_core.Truth_(temp__4388__auto__) {
													{
														var s__564___2 = temp__4388__auto__
														_ = s__564___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__564___2) {
															{
																var c__937__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__564___2)
																var size__938__auto__ = cljs_core.Count.X_invoke_Arity1(c__937__auto__).(float64)
																var b__566 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__938__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__937__auto__, size__938__auto__, b__566
																if func() bool {
																	var i__565 = float64(0)
																	_ = i__565
																	for {
																		if i__565 < size__938__auto__ {
																			{
																				var e = cljs_core.Decorate_(c__937__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__565)
																				_ = e
																				if e.(float64) > float64(2) {
																					cljs_core.Chunk_append.X_invoke_Arity2(b__566, e)
																					i__565 = (i__565 + float64(1))
																					continue
																				} else {
																					i__565 = (i__565 + float64(1))
																					continue
																				}
																			}
																		} else {
																			return true
																		}
																	}
																}() {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__566), iter__563.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__564___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__566), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__564___2)
																_ = e
																if e.(float64) > float64(2) {
																	return cljs_core.Cons.X_invoke_Arity2(e, iter__563.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__564___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
																} else {
																	s__564___1 = cljs_core.Rest.Arity1IQ(s__564___2)
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
							}(&cljs_core.AFn{}, v_1012), nil, nil})
						})
					}(&cljs_core.AFn{}, v_1012)
					_ = iter__939__auto__
					return iter__939__auto__.X_invoke_Arity1(v_1012).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [3] (for [e v :when (> e 2)] e))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil})}, nil}), func() *cljs_core.CljsCoreLazySeq {
					var iter__939__auto__ = func(iter__569 *cljs_core.AFn, v_1012 cljs_core.CljsCoreIVector) *cljs_core.AFn {
						return cljs_core.Fn(iter__569, 1, func(s__570 interface{}) interface{} {
							return (&cljs_core.CljsCoreLazySeq{nil, func(G__1017 *cljs_core.AFn, v_1012 cljs_core.CljsCoreIVector) *cljs_core.AFn {
								return cljs_core.Fn(G__1017, 0, func() interface{} {
									{
										var s__570___1 interface{} = s__570
										_ = s__570___1
										for {
											{
												var temp__4388__auto__ = cljs_core.Seq.Arity1IQ(s__570___1)
												_ = temp__4388__auto__
												if cljs_core.Truth_(temp__4388__auto__) {
													{
														var s__570___2 = temp__4388__auto__
														_ = s__570___2
														if cljs_core.Chunked_seq_QMARK_.Arity1IB(s__570___2) {
															{
																var c__937__auto__ = cljs_core.Chunk_first.X_invoke_Arity1(s__570___2)
																var size__938__auto__ = cljs_core.Count.X_invoke_Arity1(c__937__auto__).(float64)
																var b__572 = cljs_core.Chunk_buffer.X_invoke_Arity1(size__938__auto__).(*cljs_core.CljsCoreChunkBuffer)
																_, _, _ = c__937__auto__, size__938__auto__, b__572
																if cljs_core.Truth_(func() interface{} {
																	var i__571 = float64(0)
																	_ = i__571
																	for {
																		if i__571 < size__938__auto__ {
																			{
																				var e = cljs_core.Decorate_(c__937__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__571)
																				_ = e
																				if e.(float64) < float64(3) {
																					{
																						var m = (e.(float64) * e.(float64))
																						_ = m
																						cljs_core.Chunk_append.X_invoke_Arity2(b__572, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}))
																						i__571 = (i__571 + float64(1))
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
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__572), iter__569.X_invoke_Arity1(cljs_core.Chunk_rest.X_invoke_Arity1(s__570___2)).(*cljs_core.CljsCoreLazySeq))
																} else {
																	return cljs_core.Chunk_cons.X_invoke_Arity2(cljs_core.Chunk.X_invoke_Arity1(b__572), nil)
																}
															}
														} else {
															{
																var e = cljs_core.First.X_invoke_Arity1(s__570___2)
																_ = e
																if e.(float64) < float64(3) {
																	{
																		var m = (e.(float64) * e.(float64))
																		_ = m
																		return cljs_core.Cons.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{e, m}, nil}), iter__569.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(s__570___2)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreCons)
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
							}(&cljs_core.AFn{}, v_1012), nil, nil})
						})
					}(&cljs_core.AFn{}, v_1012)
					_ = iter__939__auto__
					return iter__939__auto__.X_invoke_Arity1(v_1012).(*cljs_core.CljsCoreLazySeq)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[1 1] [2 4]] (for [e v :while (< e 3) :let [m (* e e)]] [e m]))").(string)}, ``)}))
				}
			}
			if cljs_core.Not_EQ_.Arity2IIB(float64(1), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not= 1 2)").(string)}, ``)}))
			}
			if !(cljs_core.Not_EQ_.Arity2IIB(float64(1), float64(1))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (not= 1 1))").(string)}, ``)}))
			}
			if cljs_core.Not.Arity1IB(cljs_core.Not_empty.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (not-empty []))").(string)}, ``)}))
			}
			if cljs_core.Boolean.Arity1IB(cljs_core.Not_empty.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(boolean (not-empty [1 2 3]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("joel", cljs_core.Min_key.X_invoke_ArityVariadic(cljs_core.Count, "joel", "tom servo", cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"crooooooooow"}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"joel\" (min-key count \"joel\" \"tom servo\" \"crooooooooow\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("crooooooooow", cljs_core.Max_key.X_invoke_ArityVariadic(cljs_core.Count, "joel", "tom servo", cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"crooooooooow"}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"crooooooooow\" (max-key count \"joel\" \"tom servo\" \"crooooooooow\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Partition_all.X_invoke_Arity2(float64(4), (&cljs_core.CljsCorePersistentVector{nil, float64(9), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9)}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7), float64(8)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(9)}, nil})}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (partition-all 4 [1 2 3 4 5 6 7 8 9]) [[1 2 3 4] [5 6 7 8] [9]])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Partition_all.X_invoke_Arity3(float64(4), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(9), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9)}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5), float64(6)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7), float64(8)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(7), float64(8), float64(9)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(9)}, nil})}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (partition-all 4 2 [1 2 3 4 5 6 7 8 9]) [[1 2 3 4] [3 4 5 6] [5 6 7 8] [7 8 9] [9]])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true}, nil}), cljs_core.Take_while.X_invoke_Arity2(cljs_core.True_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true, float64(2), float64(3), float64(4)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [true true] (take-while true? [true true 2 3 4]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{false, false, false}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true}, nil})}, nil}), cljs_core.Partition_by.X_invoke_Arity2(cljs_core.True_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, true, false, false, false, true, true}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[true true] [false false false] [true true]] (partition-by true? [true true false false false true true]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2), float64(4), float64(6), float64(8), float64(10)}, nil}), cljs_core.Take_nth.X_invoke_Arity2(float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(11), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [0 2 4 6 8 10] (take-nth 2 [0 1 2 3 4 5 6 7 8 9 10]))").(string)}, ``)}))
			}
			{
				var a10_1018 = cljs_core.Partial.X_invoke_Arity2(cljs_core.X_PLUS_, float64(10)).(cljs_core.CljsCoreIFn)
				var a20_1019 = cljs_core.Partial.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), float64(10)).(cljs_core.CljsCoreIFn)
				var a21_1020 = cljs_core.Partial.X_invoke_Arity4(cljs_core.X_PLUS_, float64(10), float64(10), float64(1)).(cljs_core.CljsCoreIFn)
				var a22_1021 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)})).(cljs_core.CljsCoreIFn)
				var a23_1022 = cljs_core.Partial.X_invoke_ArityVariadic(cljs_core.X_PLUS_, float64(10), float64(5), float64(4), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(1)})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = a10_1018, a20_1019, a21_1020, a22_1021, a23_1022
				if cljs_core.X_EQ_.Arity2IIB(float64(110), func() interface{} {
					var G__575 = float64(100)
					_ = G__575
					return a10_1018.X_invoke_Arity1(G__575)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 110 (a10 100))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(120), func() interface{} {
					var G__576 = float64(100)
					_ = G__576
					return a20_1019.X_invoke_Arity1(G__576)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 120 (a20 100))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(121), func() interface{} {
					var G__577 = float64(100)
					_ = G__577
					return a21_1020.X_invoke_Arity1(G__577)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 121 (a21 100))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(122), func() interface{} {
					var G__578 = float64(100)
					_ = G__578
					return a22_1021.X_invoke_Arity1(G__578)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 122 (a22 100))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(123), func() interface{} {
					var G__579 = float64(100)
					_ = G__579
					return a23_1022.X_invoke_Arity1(G__579)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 123 (a23 100))").(string)}, ``)}))
				}
			}
			{
				var n2_1023 = cljs_core.Comp.X_invoke_Arity2(cljs_core.First, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n3_1024 = cljs_core.Comp.X_invoke_Arity3(cljs_core.First, cljs_core.Rest, cljs_core.Rest).(cljs_core.CljsCoreIFn)
				var n4_1025 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n5_1026 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				var n6_1027 = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.First, cljs_core.Rest, cljs_core.Rest, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest, cljs_core.Rest, cljs_core.Rest})).(cljs_core.CljsCoreIFn)
				_, _, _, _, _ = n2_1023, n3_1024, n4_1025, n5_1026, n6_1027
				if cljs_core.X_EQ_.Arity2IIB(float64(2), func() interface{} {
					var G__580 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__580
					return n2_1023.X_invoke_Arity1(G__580)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (n2 [1 2 3 4 5 6 7]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), func() interface{} {
					var G__581 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__581
					return n3_1024.X_invoke_Arity1(G__581)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (n3 [1 2 3 4 5 6 7]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(4), func() interface{} {
					var G__582 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__582
					return n4_1025.X_invoke_Arity1(G__582)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (n4 [1 2 3 4 5 6 7]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(5), func() interface{} {
					var G__583 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__583
					return n5_1026.X_invoke_Arity1(G__583)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 5 (n5 [1 2 3 4 5 6 7]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), func() interface{} {
					var G__584 = (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7)}, nil})
					_ = G__584
					return n6_1027.X_invoke_Arity1(G__584)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 6 (n6 [1 2 3 4 5 6 7]))").(string)}, ``)}))
				}
			}
			{
				var sf_1028 = cljs_core.Some_fn.X_invoke_Arity3(cljs_core.Number_QMARK_, cljs_core.Keyword_QMARK_, cljs_core.Symbol_QMARK_).(cljs_core.CljsCoreIFn)
				_ = sf_1028
				if cljs_core.Truth_(func() interface{} {
					var G__585 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					var G__586 = float64(1)
					_, _ = G__585, G__586
					return sf_1028.X_invoke_Arity2(G__585, G__586)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(sf :foo 1)").(string)}, ``)}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__587 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
					_ = G__587
					return sf_1028.X_invoke_Arity1(G__587)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(sf :foo)").(string)}, ``)}))
				}
				if cljs_core.Truth_(func() interface{} {
					var G__588 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})
					var G__589 = float64(1)
					_, _ = G__588, G__589
					return sf_1028.X_invoke_Arity2(G__588, G__589)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(sf (quote bar) 1)").(string)}, ``)}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__590 = cljs_core.CljsCorePersistentVector_EMPTY
					var G__591 = cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)
					_, _ = G__590, G__591
					return sf_1028.X_invoke_Arity2(G__590, G__591)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (sf [] ()))").(string)}, ``)}))
				}
			}
			{
				var ep_1029 = cljs_core.Every_pred.X_invoke_Arity2(cljs_core.Number_QMARK_, cljs_core.Zero_QMARK_).(cljs_core.CljsCoreIFn)
				_ = ep_1029
				if cljs_core.Truth_(func() interface{} {
					var G__592 = float64(0)
					var G__593 = float64(0)
					var G__594 = float64(0)
					_, _, _ = G__592, G__593, G__594
					return ep_1029.X_invoke_Arity3(G__592, G__593, G__594)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(ep 0 0 0)").(string)}, ``)}))
				}
				if cljs_core.Not.Arity1IB(func() interface{} {
					var G__595 = float64(1)
					var G__596 = float64(2)
					var G__597 = float64(3)
					var G__598 = float64(0)
					_, _, _, _ = G__595, G__596, G__597, G__598
					return ep_1029.X_invoke_Arity4(G__595, G__596, G__597, G__598)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (ep 1 2 3 0))").(string)}, ``)}))
				}
			}
			if cljs_core.Truth_(cljs_core.Complement.X_invoke_Arity1(cljs_core.Number_QMARK_).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("((complement number?) :foo)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})}, nil}), cljs_core.Juxt.X_invoke_Arity3(cljs_core.First, cljs_core.Rest, cljs_core.Seq).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 [2 3] [1 2 3]] ((juxt first rest seq) [1 2 3]))").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 5 (max 1 2 3 4 5))").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 5 (max 5 4 3 2 1))").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 5.5 (max 1 2 3 4 5 5.5))").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (min 5 4 3 2 1))").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (min 1 2 3 4 5))").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0.5 (min 5 4 3 0.5 2 1))").(string)}, ``)}))
			}
			if cljs_core.Truth_(cljs_core.Set.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(set [])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, cljs_core.Set.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{} (set []))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, cljs_core.CljsCorePersistentHashSet_EMPTY) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{} (hash-set))").(string)}, ``)}))
			}
			if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentHashSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? cljs.core/PersistentHashSet (type (hash-set)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"foo", nil}, nil}, nil}), cljs_core.Set.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo"}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{\"foo\"} (set [\"foo\"]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"foo", nil}, nil}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"foo", nil}, nil}), nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{\"foo\"} (hash-set \"foo\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{1 3 2} #{1 3 2})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(7), float64(8)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6)}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(9), float64(10)}, true).(*cljs_core.CljsCorePersistentHashSet), cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(7), float64(8)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6)}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(9), float64(10)}, true).(*cljs_core.CljsCorePersistentHashSet)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{{7 8} [4 5 6] #{1 3 2} 9 10} #{{7 8} [4 5 6] #{1 3 2} 9 10})").(string)}, ``)}))
			}
			if !(cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{nil, float64(0), cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashSet_EMPTY}, true).(*cljs_core.CljsCorePersistentHashSet), cljs_core.CljsCorePersistentHashSet_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (= #{nil 0 [] {} #{}} #{}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{nil, float64(0), cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashSet_EMPTY}, true).(*cljs_core.CljsCorePersistentHashSet)).(float64), float64(5)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count #{nil 0 [] {} #{}}) 5)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), float64(1)), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (conj #{1} 1) #{1})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), float64(2)), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (conj #{1} 2) #{1 2})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{float64(1), nil, float64(4), nil, float64(3), nil, float64(2), nil}, nil}, nil}).X_empty_Arity1()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{} (-empty #{1 4 3 2}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(5), []interface{}{float64(1), nil, float64(4), nil, float64(3), nil, float64(2), nil, float64(5), nil}, nil}, nil})), float64(15)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (reduce + #{1 4 3 2 5}) 15)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{float64(1), nil, float64(4), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(4))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get #{1 4 3 2} 4))").(string)}, ``)}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{float64(1), nil, float64(4), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(4)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? #{1 4 3 2} 4)").(string)}, ``)}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{nil, float64(0), cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashSet_EMPTY}, true).(*cljs_core.CljsCorePersistentHashSet), cljs_core.CljsCorePersistentArrayMap_EMPTY) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? #{nil 0 [] {} #{}} {})").(string)}, ``)}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})}, true).(*cljs_core.CljsCorePersistentHashSet), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? #{[1 2 3]} [1 2 3])").(string)}, ``)}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}).X_disjoin_Arity2(float64(3)), float64(3))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? (-disjoin #{1 3 2} 3) 3))").(string)}, ``)}))
			}
			if float64(-1) < float64(0) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(neg? -1)").(string)}, ``)}))
			}
			if !(float64(1) < float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (neg? 1))").(string)}, ``)}))
			}
			if -1.765 < float64(0) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(neg? -1.765)").(string)}, ``)}))
			}
			if !(float64(0) < float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (neg? 0))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, true, false, true, false, true, false}, nil}), cljs_core.Map_.X_invoke_Arity2(cljs_core.Integer_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), 1.00001, float64(2023), cljs_core.CljsCorePersistentVector_EMPTY, (float64(88) - float64(1001991881)), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(0), "0"}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [true false true false true false true false] (map integer? [1 1.00001 2023 [] (- 88 1001991881) :foo 0 \"0\"]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, true, false, true, false}, nil}), cljs_core.Map_.X_invoke_Arity2(cljs_core.Odd_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(-1), float64(0)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [true false true false true false] (map odd? [1 2 3 4 -1 0]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, true, false, true, true}, nil}), cljs_core.Map_.X_invoke_Arity2(cljs_core.Even_QMARK_, (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(-2), float64(0)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [true false true false true true] (map even? [2 3 4 5 -2 0]))").(string)}, ``)}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? {:a 1, :b 2} :a)").(string)}, ``)}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "z", Fqn: "z", X_hash: float64(-789527183)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? {:a 1, :b 2} :z))").(string)}, ``)}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil}), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? [5 6 7] 1)").(string)}, ``)}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil}), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? [5 6 7] 2)").(string)}, ``)}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil}), float64(3))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? [5 6 7] 3))").(string)}, ``)}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil})).([]interface{}), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? (to-array [5 6 7]) 1)").(string)}, ``)}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil})).([]interface{}), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? (to-array [5 6 7]) 2)").(string)}, ``)}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(6), float64(7)}, nil})).([]interface{}), float64(3))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? (to-array [5 6 7]) 3))").(string)}, ``)}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB(nil, float64(42))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? nil 42))").(string)}, ``)}))
			}
			if cljs_core.Contains_QMARK_.Arity2IIB("f", float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? \"f\" 0)").(string)}, ``)}))
			}
			if !(cljs_core.Contains_QMARK_.Arity2IIB("f", float64(55))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? \"f\" 55))").(string)}, ``)}))
			}
			if cljs_core.Truth_(cljs_core.Distinct_QMARK_.X_invoke_ArityVariadic(float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(distinct? 1 2 3)").(string)}, ``)}))
			}
			if !(cljs_core.Truth_(cljs_core.Distinct_QMARK_.X_invoke_ArityVariadic(float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(1)})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (distinct? 1 2 3 1))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)).(*cljs_core.CljsCoreLazySeq), cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct ()) ())").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(*cljs_core.CljsCoreList)).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct (quote (1))) (quote (1)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(1), float64(1), float64(1)})).(*cljs_core.CljsCoreList)).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct (quote (1 2 3 1 1 1))) (quote (1 2 3)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(1), float64(2)}, nil})).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [1 1 1 2]) (quote (1 2)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(1), float64(2)}, nil})).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [1 2 1 2]) (quote (1 2)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1("a").(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct \"a\") [\"a\"])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1("abcabab").(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a", "b", "c"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct \"abcabab\") [\"a\" \"b\" \"c\"])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"abc", "abc"}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"abc"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [\"abc\" \"abc\"]) [\"abc\"])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, nil}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [nil nil]) [nil])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{0.0, 0.0}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{0.0}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [0.0 0.0]) [0.0])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "sym", Str: "sym", X_hash: float64(195671222), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "sym", Str: "sym", X_hash: float64(195671222), X_meta: nil})}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "sym", Str: "sym", X_hash: float64(195671222), X_meta: nil})}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [(quote sym) (quote sym)]) (quote [sym]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "kw", Fqn: "kw", X_hash: float64(1158308175)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "kw", Fqn: "kw", X_hash: float64(1158308175)})}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "kw", Fqn: "kw", X_hash: float64(1158308175)})}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [:kw :kw]) [:kw])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42), float64(42)}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [42 42]) [42])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.CljsCorePersistentVector_EMPTY}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentVector_EMPTY}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [[] []]) [[]])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [(quote (1 2)) (quote (1 2))]) (quote [(1 2)]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [() ()]) [()])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [[1 2] [1 2]]) [[1 2]])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [{:a 1, :b 2} {:a 1, :b 2}]) [{:a 1, :b 2}])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentArrayMap_EMPTY}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [{} {}]) [{}])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil})}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil})}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [#{1 2} #{1 2}]) [#{1 2}])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Distinct.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentHashSet_EMPTY, cljs_core.CljsCorePersistentHashSet_EMPTY}, nil})).(*cljs_core.CljsCoreLazySeq), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentHashSet_EMPTY}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (distinct [#{} #{}]) [#{}])").(string)}, ``)}))
			}
			{
				var r1_1030 = (&js.RegExp{Pattern: `foo`, Flags: ``})
				var r2_1031 = cljs_core.Re_pattern.X_invoke_Arity1(r1_1030)
				_, _ = r1_1030, r2_1031
				if cljs_core.X_EQ_.Arity2IIB(r1_1030, r2_1031) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= r1 r2)").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_find.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo"), "foo bar foo baz foo zot"), "foo") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (re-find (re-pattern \"foo\") \"foo bar foo baz foo zot\") \"foo\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_find.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("f(.)o"), "foo bar foo baz foo zot"), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (re-find (re-pattern \"f(.)o\") \"foo bar foo baz foo zot\") [\"foo\" \"o\"])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo"), "foo"), "foo") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (re-matches (re-pattern \"foo\") \"foo\") \"foo\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo"), "foo bar foo baz foo zot"), nil) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (re-matches (re-pattern \"foo\") \"foo bar foo baz foo zot\") nil)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo.*"), "foo bar foo baz foo zot"), "foo bar foo baz foo zot") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (re-matches (re-pattern \"foo.*\") \"foo bar foo baz foo zot\") \"foo bar foo baz foo zot\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_seq.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("foo"), "foo bar foo baz foo zot"), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2("foo")).(cljs_core.CljsCoreICollection).X_conj_Arity2("foo")).(cljs_core.CljsCoreICollection).X_conj_Arity2("foo")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (re-seq (re-pattern \"foo\") \"foo bar foo baz foo zot\") (list \"foo\" \"foo\" \"foo\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_seq.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("f(.)o"), "foo bar foo baz foo zot"), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "o"}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (re-seq (re-pattern \"f(.)o\") \"foo bar foo baz foo zot\") (list [\"foo\" \"o\"] [\"foo\" \"o\"] [\"foo\" \"o\"]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Re_matches.X_invoke_Arity2(cljs_core.Re_pattern.X_invoke_Arity1("(?i)foo"), "Foo"), "Foo") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (re-matches (re-pattern \"(?i)foo\") \"Foo\") \"Foo\")").(string)}, ``)}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{"#\"(?:)\"", nil, "#\"\"", nil}, nil}, nil}).X_invoke_Arity1(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.RegExp{Pattern: ``, Flags: ``})})).(string))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(#{\"#\\\"(?:)\\\"\" \"#\\\"\\\"\"} (pr-str #\"\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__599 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__599, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__599, float64(1), nil)
				_, _, _ = vec__599, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 1] (let [[a b] [1 2]] [b a]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), func() cljs_core.CljsCoreISet {
				var vec__600 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__600, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__600, float64(1), nil)
				_, _, _ = vec__600, a, b
				return cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{a, b}, true).(*cljs_core.CljsCorePersistentHashSet)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{1 2} (let [[a b] [1 2]] #{a b}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__601 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__601___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__601) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__601)
					} else {
						return map__601
					}
				}()
				var a = cljs_core.Get.X_invoke_Arity2(map__601___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				var b = cljs_core.Get.X_invoke_Arity2(map__601___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				_, _, _, _ = map__601, map__601___1, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2] (let [{a :a, b :b} {:a 1, :b 2}] [a b]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__602 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__602___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__602) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__602)
					} else {
						return map__602
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__602___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__602___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__602, map__602___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2] (let [{:keys [a b]} {:a 1, :b 2}] [a b]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__603 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__603, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__603, float64(1), nil)
				var v = vec__603
				_, _, _, _ = vec__603, a, b, v
				return (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b, v}, nil})
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 [1 2]] (let [[a b :as v] [1 2]] [a b v]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(42)}, nil}), func() cljs_core.CljsCoreIVector {
				var map__604 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__604___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__604) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__604)
					} else {
						return map__604
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity3(map__604___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(42))
				var a = cljs_core.Get.X_invoke_Arity2(map__604___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__604, map__604___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 42] (let [{:keys [a b], :or {b 42}} {:a 1}] [a b]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), nil}, nil}), func() cljs_core.CljsCoreIVector {
				var map__605 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})
				var map__605___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__605) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__605)
					} else {
						return map__605
					}
				}()
				var b = cljs_core.Get.X_invoke_Arity2(map__605___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a = cljs_core.Get.X_invoke_Arity2(map__605___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__605, map__605___1, b, a
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{a, b}, nil})
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 nil] (let [{:keys [a b], :or {c 42}} {:a 1}] [a b]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__606 = cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2)})).(*cljs_core.CljsCoreList)
				var a = cljs_core.Nth.X_invoke_Arity3(vec__606, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__606, float64(1), nil)
				_, _, _ = vec__606, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 1] (let [[a b] (quote (1 2))] [b a]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), func() cljs_core.CljsCoreIMap {
				var vec__607 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})
				var a = cljs_core.Nth.X_invoke_Arity3(vec__607, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__607, float64(1), nil)
				_, _, _ = vec__607, a, b
				return cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{a, b}, true, false).(*cljs_core.CljsCorePersistentArrayMap)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {1 2} (let [[a b] [1 2]] {a b}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), func() cljs_core.CljsCoreIVector {
				var vec__608 = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))
				var a = cljs_core.Nth.X_invoke_Arity3(vec__608, float64(0), nil)
				var b = cljs_core.Nth.X_invoke_Arity3(vec__608, float64(1), nil)
				_, _, _ = vec__608, a, b
				return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{b, a}, nil})
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 1] (let [[a b] (seq [1 2])] [b a]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(1)}, nil})}, nil})}, nil}), cljs_core.Update_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(0)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}), cljs_core.Inc)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo {:bar {:baz 1}}} (update-in {:foo {:bar {:baz 0}}} [:foo :bar :baz] inc))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(10)}, nil}), cljs_core.Update_in.X_invoke_Arity4((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}), cljs_core.X_PLUS_, float64(7))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1, :bar 2, :baz 10} (update-in {:foo 1, :bar 2, :baz 3} [:baz] + 7))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(3)}, nil})}, nil}), cljs_core.Update_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}), cljs_core.Inc)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [{:foo 1, :bar 2} {:foo 1, :bar 3}] (update-in [{:foo 1, :bar 2} {:foo 1, :bar 2}] [1 :bar] inc))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(3)}, nil})}, nil})}, nil}), cljs_core.Update_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}), cljs_core.Inc)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [{:foo {:bar 2}} {:foo {:bar 3}}] (update-in [{:foo {:bar 2}} {:foo {:bar 2}}] [1 :foo :bar] inc))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(100)}, nil})}, nil})}, nil}), cljs_core.Assoc_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(0)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}), float64(100))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo {:bar {:baz 100}}} (assoc-in {:foo {:bar {:baz 0}}} [:foo :bar :baz] 100))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(100)}, nil}), cljs_core.Assoc_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}), float64(100))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1, :bar 2, :baz 100} (assoc-in {:foo 1, :bar 2, :baz 3} [:baz] 100))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(100)}, nil})}, nil})}, nil})}, nil}), cljs_core.Assoc_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}), float64(100))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [{:foo [{:bar 2} {:baz 3}]} {:foo [{:bar 2} {:baz 100}]}] (assoc-in [{:foo [{:bar 2} {:baz 3}]} {:foo [{:bar 2} {:baz 3}]}] [1 :foo 1 :baz] 100))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(100)}, nil})}, nil}), cljs_core.Assoc_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}), float64(100))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [{:foo 1, :bar 2} {:foo 1, :bar 100}] (assoc-in [{:foo 1, :bar 2} {:foo 1, :bar 2}] [1 :bar] 100))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Get_in.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (get-in {:foo 1, :bar 2} [:foo]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Get_in.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (get-in {:foo {:bar 2}} [:foo :bar]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Get_in.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (get-in [{:foo 1} {:foo 2}] [0 :foo]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get_in.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(1)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "buzz", Fqn: "buzz", X_hash: float64(298530064)}), float64(2)}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "buzz", Fqn: "buzz", X_hash: float64(298530064)}), float64(4)}, nil})}, nil})}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "buzz", Fqn: "buzz", X_hash: float64(298530064)})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get-in [{:foo 1, :bar [{:baz 1} {:buzz 2}]} {:foo 3, :bar [{:baz 3} {:buzz 4}]}] [1 :bar 1 :buzz]))").(string)}, ``)}))
			}
			{
				var a_1032 = cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{})
				_ = a_1032
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20), float64(30)}, nil}), cljs_core.Seq.Arity1IQ(func() []interface{} {
					var a__1058__auto__ = a_1032
					var ret = cljs_core.Aclone.X_invoke_Arity1(a__1058__auto__).([]interface{})
					_, _ = a__1058__auto__, ret
					{
						var i = float64(0)
						_ = i
						for {
							if i < float64(len(a__1058__auto__)) {
								ret[int(i)] = (float64(10) * (a_1032[int(i)]).(float64))
								i = (i + float64(1))
								continue
							} else {
								return ret
							}
						}
					}
				}())) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [10 20 30] (seq (amap a i ret (* 10 (aget a i)))))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), func() float64 {
					var a__1064__auto__ = a_1032
					_ = a__1064__auto__
					{
						var i = float64(0)
						var ret = float64(0)
						_, _ = i, ret
						for {
							if i < float64(len(a__1064__auto__)) {
								i, ret = (i + float64(1)), (ret + (a_1032[int(i)]).(float64))
								continue
							} else {
								return ret
							}
						}
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 6 (areduce a i ret 0 (+ ret (aget a i))))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_1032), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq a) (seq (to-array [1 2 3])))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(42), func() float64 { a_1032[int(float64(0))] = float64(42); return a_1032[int(float64(0))].(float64) }()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 42 (aset a 0 42))").(string)}, ``)}))
				}
				if cljs_core.Not_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(a_1032), cljs_core.Seq.Arity1IQ(cljs_core.To_array.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).([]interface{}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not= (seq a) (seq (to-array [1 2 3])))").(string)}, ``)}))
				}
			}
			{
				var a_1033 = []interface{}{[]interface{}{float64(1), float64(2), float64(3)}, []interface{}{float64(4), float64(5), float64(6)}}
				_ = a_1033
				if cljs_core.X_EQ_.Arity2IIB((a_1033[int(float64(0))].([]interface{})[int(float64(1))]), float64(2)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (aget a 0 1) 2)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity3(cljs_core.Aget, a_1033, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil})), float64(2)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply aget a [0 1]) 2)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((a_1033[int(float64(1))].([]interface{})[int(float64(1))]), float64(5)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (aget a 1 1) 5)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity3(cljs_core.Aget, a_1033, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil})), float64(5)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply aget a [1 1]) 5)").(string)}, ``)}))
				}
				a_1033[int(float64(0))].([]interface{})[int(float64(0))] = "foo"
				if cljs_core.X_EQ_.Arity2IIB((a_1033[int(float64(0))].([]interface{})[int(float64(0))]), "foo") {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (aget a 0 0) \"foo\")").(string)}, ``)}))
				}
				cljs_core.Apply.X_invoke_Arity3(cljs_core.Aset, a_1033, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0), "bar"}, nil}))
				if cljs_core.X_EQ_.Arity2IIB((a_1033[int(float64(0))].([]interface{})[int(float64(0))]), "bar") {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (aget a 0 0) \"bar\")").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Sort.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(3), float64(1), float64(4), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3 4 5] (sort [5 3 1 4 2]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil}), cljs_core.Sort.X_invoke_Arity2(cljs_core.X_LT_, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(3), float64(1), float64(4), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3 4 5] (sort < [5 3 1 4 2]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(4), float64(3), float64(2), float64(1)}, nil}), cljs_core.Sort.X_invoke_Arity2(cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(3), float64(1), float64(4), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [5 4 3 2 1] (sort > [5 3 1 4 2]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), "foo"}, nil}), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.Count, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "a", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [\"a\" [1 2] \"foo\"] (sort-by count [\"foo\" \"a\" [1 2]]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), "a"}, nil}), cljs_core.Sort_by.X_invoke_Arity3(cljs_core.Count, cljs_core.X_GT_, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "a", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [\"foo\" [1 2] \"a\"] (sort-by count > [\"foo\" \"a\" [1 2]]))").(string)}, ``)}))
			}
			{
				var coll_1034 = (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})
				var shuffles_1035 = cljs_core.Filter.X_invoke_Arity2(func(G__1036 *cljs_core.AFn, coll_1034 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1036, 1, func(p1__58_SHARP_ interface{}) interface{} {
						return cljs_core.Not_EQ_.Arity2IIB(coll_1034, p1__58_SHARP_)
					})
				}(&cljs_core.AFn{}, coll_1034), cljs_core.Take.X_invoke_Arity2(float64(100), cljs_core.Iterate.X_invoke_Arity2(cljs_core.Shuffle, coll_1034).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)
				_, _ = coll_1034, shuffles_1035
				if !(cljs_core.Empty_QMARK_.Arity1IB(shuffles_1035)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (empty? shuffles))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Last.X_invoke_Arity1(nil)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (last nil))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Last.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (last [1 2 3]))").(string)}, ``)}))
			}
			{
				var s_1037 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = s_1037
				{
					var n__1070__auto___1038 = float64(5)
					_ = n__1070__auto___1038
					{
						var n_1039 = float64(0)
						_ = n_1039
						for {
							if n_1039 < n__1070__auto___1038 {
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_1037, cljs_core.Conj, n_1039)
								n_1039 = (n_1039 + float64(1))
								continue
							} else {
							}
							break
						}
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Deref.X_invoke_Arity1(s_1037)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [0 1 2 3 4] (clojure.core/deref s))").(string)}, ``)}))
				}
			}
			{
				var v_1040 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})
				var s_1041 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY)).(*cljs_core.CljsCoreAtom)
				_, _ = v_1040, s_1041
				{
					var seq__609_1042 interface{} = cljs_core.Seq.Arity1IQ(v_1040)
					var chunk__610_1043 interface{} = nil
					var count__611_1044 = float64(0)
					var i__612_1045 = float64(0)
					_, _, _, _ = seq__609_1042, chunk__610_1043, count__611_1044, i__612_1045
					for {
						if i__612_1045 < count__611_1044 {
							{
								var n_1046 = cljs_core.Decorate_(chunk__610_1043).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__612_1045)
								_ = n_1046
								cljs_core.Swap_BANG_.X_invoke_Arity3(s_1041, cljs_core.Conj, n_1046)
								seq__609_1042, chunk__610_1043, count__611_1044, i__612_1045 = seq__609_1042, chunk__610_1043, count__611_1044, (i__612_1045 + float64(1))
								continue
							}
						} else {
							{
								var temp__4388__auto___1047 = cljs_core.Seq.Arity1IQ(seq__609_1042)
								_ = temp__4388__auto___1047
								if cljs_core.Truth_(temp__4388__auto___1047) {
									{
										var seq__609_1048___1 = temp__4388__auto___1047
										_ = seq__609_1048___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__609_1048___1) {
											{
												var c__970__auto___1049 = cljs_core.Chunk_first.X_invoke_Arity1(seq__609_1048___1)
												_ = c__970__auto___1049
												seq__609_1042, chunk__610_1043, count__611_1044, i__612_1045 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__609_1048___1), c__970__auto___1049, cljs_core.Count.X_invoke_Arity1(c__970__auto___1049).(float64), float64(0)
												continue
											}
										} else {
											{
												var n_1050 = cljs_core.First.X_invoke_Arity1(seq__609_1048___1)
												_ = n_1050
												cljs_core.Swap_BANG_.X_invoke_Arity3(s_1041, cljs_core.Conj, n_1050)
												seq__609_1042, chunk__610_1043, count__611_1044, i__612_1045 = cljs_core.Next.Arity1IQ(seq__609_1048___1), nil, float64(0), float64(0)
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Deref.X_invoke_Arity1(s_1041), cljs_core.Reverse.X_invoke_Arity1(v_1040)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (clojure.core/deref s) (reverse v))").(string)}, ``)}))
				}
			}
			{
				var a_1051 = cljs_core.Atom.X_invoke_Arity1(float64(0)).(*cljs_core.CljsCoreAtom)
				var d_1052 = (&cljs_core.CljsCoreDelay{func(G__1053 *cljs_core.AFn, a_1051 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__1053, 0, func() interface{} {
						return cljs_core.Swap_BANG_.X_invoke_Arity2(a_1051, cljs_core.Inc)
					})
				}(&cljs_core.AFn{}, a_1051), nil})
				_, _ = a_1051, d_1052
				if cljs_core.Realized_QMARK_.Arity1IB(d_1052) == false {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(false? (realized? d))").(string)}, ``)}))
				}
				if cljs_core.Deref.X_invoke_Arity1(a_1051).(float64) == float64(0) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(zero? (clojure.core/deref a))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_1052)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (clojure.core/deref d))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_1051)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (clojure.core/deref a))").(string)}, ``)}))
				}
				if cljs_core.Realized_QMARK_.Arity1IB(d_1052) == true {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(true? (realized? d))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(d_1052)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (clojure.core/deref d))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Deref.X_invoke_Arity1(a_1051)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (clojure.core/deref a))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Force.X_invoke_Arity1(d_1052), cljs_core.Deref.X_invoke_Arity1(d_1052)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (force d) (clojure.core/deref d))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Force.X_invoke_Arity1(float64(1))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (force 1))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Assoc.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentArrayMap_EMPTY, float64(1), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(4)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {1 2, 3 4} (assoc {} 1 2 3 4))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), cljs_core.Assoc.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, float64(1), float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {1 2} (assoc {} 1 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42), float64(2)}, nil}), cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(0), float64(42))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [42 2] (assoc [1 2] 0 42))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Dissoc.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {} (dissoc {1 2, 3 4} 1 3))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), cljs_core.Dissoc.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), float64(3))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {1 2} (dissoc {1 2, 3 4} 3))").(string)}, ``)}))
			}
			if cljs_core.Nil_(cljs_core.Dissoc.X_invoke_Arity2(nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (dissoc nil :foo))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), cljs_core.Disj.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{1 3 2} (disj #{1 3 2}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), cljs_core.Disj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(3))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{1 2} (disj #{1 3 2} 3))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), cljs_core.Disj.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}), float64(2), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{1} (disj #{1 3 2} 2 3))").(string)}, ``)}))
			}
			if cljs_core.Nil_(cljs_core.Disj.X_invoke_Arity2(nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (disj nil :foo))").(string)}, ``)}))
			}
			{
				var f_1054 = cljs_core.Memoize.X_invoke_Arity1(func(G__1055 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1055, 0, func() interface{} {
						return cljs_core.Rand.Arity0F()
					})
				}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn)
				_ = f_1054
				{
					f_1054.X_invoke_Arity0()
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return f_1054.X_invoke_Arity0()
				}(), func() interface{} {
					return f_1054.X_invoke_Arity0()
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (f) (f))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})), nil) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (find {} :a) nil)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (find {:a 1} :a) [:a 1])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})), nil) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (find {:a 1} :b) nil)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (find {:a 1, :b 2} :a) [:a 1])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (find {:a 1, :b 2} :b) [:b 2])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})), nil) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (find {:a 1, :b 2} :c) nil)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, nil), nil) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (find {} nil) nil)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), nil), nil) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (find {:a 1} nil) nil)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), nil), nil) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (find {:a 1, :b 2} nil) nil)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Find.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(0)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (find [1 2 3] 0) [0 1])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(4), float64(2)).(float64), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 4 2) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(3), float64(2)).(float64), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 3 2) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(6), float64(4)).(float64), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 6 4) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(0), float64(5)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 0 5) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(42), float64(5)).(float64), float64(8)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 42 5) 8)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(42), float64(-5)).(float64), float64(-8)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 42 -5) -8)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(-42), float64(-5)).(float64), float64(8)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot -42 -5) 8)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(9), float64(3)).(float64), float64(3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 9 3) 3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(9), float64(-3)).(float64), float64(-3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 9 -3) -3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(-9), float64(3)).(float64), float64(-3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot -9 3) -3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(2), float64(-5)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 2 -5) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(-2), float64(5)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot -2 5) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(0), float64(3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 0 3) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Quot.X_invoke_Arity2(float64(0), float64(-3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quot 0 -3) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(4), float64(2)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 4 2) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(3), float64(2)).(float64), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 3 2) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(6), float64(4)).(float64), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 6 4) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(0), float64(5)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 0 5) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(4.5, 2.0).(float64), 0.5) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 4.5 2.0) 0.5)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(42), float64(5)).(float64), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 42 5) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(9), float64(3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 9 3) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(9), float64(-3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 9 -3) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(-9), float64(3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod -9 3) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(-9), float64(-3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod -9 -3) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(0), float64(3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 0 3) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mod.X_invoke_Arity2(float64(3216478362187432), float64(432143214)).(float64), float64(120355456)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mod 3216478362187432 432143214) 120355456)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(4), float64(2)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rem 4 2) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(0), float64(5)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rem 0 5) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(4.5, 2.0).(float64), 0.5) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rem 4.5 2.0) 0.5)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(42), float64(5)).(float64), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rem 42 5) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(2), float64(5)).(float64), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rem 2 5) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(2), float64(-5)).(float64), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rem 2 -5) 2)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rem.X_invoke_Arity2(float64(0), float64(3)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rem 0 3) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(9))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(8))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(6))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(5))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(4))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 10) (list 0 1 2 3 4 5 6 7 8 9))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity2(float64(10), float64(20)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(19))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(18))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(17))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(16))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(15))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(14))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(13))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(12))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(11))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(10))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 10 20) (list 10 11 12 13 14 15 16 17 18 19))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(10), float64(20), float64(2)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(18))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(16))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(14))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(12))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(10))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 10 20 2) (list 10 12 14 16 18))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Take.X_invoke_Arity2(float64(20), cljs_core.Range_.X_invoke_Arity0().(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(19))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(18))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(17))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(16))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(15))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(14))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(13))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(12))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(11))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(10))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(9))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(8))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(6))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(5))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(4))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (take 20 (range)) (list 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19))").(string)}, ``)}))
			}
			{
				var d_1056 = cljs_core.Group_by.X_invoke_Arity2(cljs_core.Second, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(6), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), float64(4), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), float64(2)}, nil}))
				_ = d_1056
				if cljs_core.X_EQ_.Arity2IIB(float64(3), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_1056, float64(1))).(float64)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (count (get d 1)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_1056, float64(2))).(float64)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (count (get d 2)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Count.X_invoke_Arity1(cljs_core.Get.X_invoke_Arity2(d_1056, float64(4))).(float64)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (count (get d 4)))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Merge.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(3), float64(4)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(5), float64(6)}, nil})}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {1 2, 3 4, 5 6} (merge {1 2} {3 4} {5 6}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}), cljs_core.Merge.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(3), float64(4)}, nil}), nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {1 2, 3 4} (merge {1 2} {3 4} nil))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), cljs_core.Frequencies.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:a 3, :b 2} (frequencies [:a :b :a :b :a]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(6), float64(10), float64(15)}, nil}), cljs_core.Reductions.X_invoke_Arity2(cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 3 6 10 15] (reductions + [1 2 3 4 5]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__1057 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1057, 1, func(p1__59_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__59_SHARP_) {
						return p1__59_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 3 5 7 9] (keep (fn* [p1__59#] (if (odd? p1__59#) p1__59#)) [1 2 3 4 5 6 7 8 9 10]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(6), float64(8), float64(10)}, nil}), cljs_core.Keep.X_invoke_Arity2(func(G__1058 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1058, 1, func(p1__60_SHARP_ interface{}) interface{} {
					if cljs_core.Even_QMARK_.Arity1IB(p1__60_SHARP_) {
						return p1__60_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 4 6 8 10] (keep (fn* [p1__60#] (if (even? p1__60#) p1__60#)) [1 2 3 4 5 6 7 8 9 10]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__1059 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1059, 2, func(p1__61_SHARP_ interface{}, p2__62_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__61_SHARP_) {
						return p2__62_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(11), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 3 5 7 9] (keep-indexed (fn* [p1__61# p2__62#] (if (odd? p1__61#) p2__62#)) [0 1 2 3 4 5 6 7 8 9 10]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4), float64(5)}, nil}), cljs_core.Keep_indexed.X_invoke_Arity2(func(G__1060 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1060, 2, func(p1__64_SHARP_ interface{}, p2__63_SHARP_ interface{}) interface{} {
					if p2__63_SHARP_.(float64) > float64(0) {
						return p1__64_SHARP_
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(-9), float64(0), float64(29), float64(-7), float64(45), float64(3), float64(-8)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 4 5] (keep-indexed (fn* [p1__64# p2__63#] (if (pos? p2__63#) p1__64#)) [-9 0 29 -7 45 3 -8]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil}), cljs_core.Map_indexed.X_invoke_Arity2(func(G__1061 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1061, 2, func(p1__65_SHARP_ interface{}, p2__66_SHARP_ interface{}) interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__65_SHARP_, p2__66_SHARP_}, nil})
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[0 :a] [1 :b] [2 :c]] (map-indexed (fn* [p1__65# p2__66#] (vector p1__65# p2__66#)) [:a :b :c]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{"Foo", cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"foo", "FOO", "fOo"})).(*cljs_core.CljsCoreList), "Bar", cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"bar", "BAR", "BAr"})).(*cljs_core.CljsCoreList), "Baz", (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"baz"}, nil}), "Qux", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"qux", "quux"}, nil})}, nil}), cljs_core.Merge_with.X_invoke_ArityVariadic(cljs_core.Concat, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{"Foo", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"foo", "FOO"}, nil}), "Bar", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"bar", "BAR"}, nil}), "Baz", (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"baz"}, nil})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{"Foo", (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"fOo"}, nil}), "Bar", (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"BAr"}, nil}), "Qux", (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"qux", "quux"}, nil})}, nil})}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote {\"Foo\" (\"foo\" \"FOO\" \"fOo\"), \"Bar\" (\"bar\" \"BAR\" \"BAr\"), \"Baz\" [\"baz\"], \"Qux\" [\"qux\" \"quux\"]}) (merge-with concat {\"Foo\" [\"foo\" \"FOO\"], \"Bar\" [\"bar\" \"BAR\"], \"Baz\" [\"baz\"]} {\"Foo\" [\"fOo\"], \"Bar\" [\"BAr\"], \"Qux\" [\"qux\" \"quux\"]}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(111), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(102), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(13)}, nil}), cljs_core.Merge_with.X_invoke_ArityVariadic(cljs_core.X_PLUS_, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(3)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(10), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(10)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(100), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(100)}, nil})}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:a 111, :b 102, :c 13} (merge-with + {:a 1, :b 2, :c 3} {:a 10, :c 10} {:a 100, :b 100}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(102), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(13)}, nil}), cljs_core.Apply.X_invoke_Arity2(cljs_core.Merge_with, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.X_PLUS_, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(100)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(3)}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(10)}, nil})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:a 3, :b 102, :c 13} (apply merge-with [+ {:a 1, :b 100} {:a 1, :b 2, :c 3} {:a 1, :c 10}]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "e", Str: "e", X_hash: float64(-1273166571), X_meta: nil})}, nil}), cljs_core.Replace.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "d", Str: "d", X_hash: float64(-682293345), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "e", Str: "e", X_hash: float64(-1273166571), X_meta: nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2), float64(4)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote [a c e]) (replace (quote [a b c d e]) [0 2 4]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "zero", Fqn: "zero", X_hash: float64(-858964576)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "two", Fqn: "two", X_hash: float64(627606869)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "zero", Fqn: "zero", X_hash: float64(-858964576)})}, nil}), cljs_core.Replace.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "zero", Fqn: "zero", X_hash: float64(-858964576)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "two", Fqn: "two", X_hash: float64(627606869)})}, nil}), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(0), float64(2), float64(0)})).(*cljs_core.CljsCoreList))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [:one :zero :two :zero] (replace {0 :zero, 1 :one, 2 :two} (quote (1 0 2 0))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4), float64(5)}, nil})}, nil}), cljs_core.Split_at.X_invoke_Arity2(float64(2), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})).(cljs_core.CljsCoreIVector)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[1 2] [3 4 5]] (split-at 2 [1 2 3 4 5]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil})}, nil}), cljs_core.Split_with.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.X_GT__EQ_, float64(3)).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)}, nil})).(cljs_core.CljsCoreIVector)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [[1 2 3] [4 5]] (split-with (partial >= 3) [1 2 3 4 5]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(10000), cljs_core.Trampoline.X_invoke_ArityVariadic(func(f *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(f, 1, func(n interface{}) interface{} {
					if n.(float64) >= float64(10000) {
						return n
					} else {
						return func(G__1062 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1062, 0, func() interface{} {
								return f.X_invoke_Arity1((n.(float64) + float64(1)))
							})
						}(&cljs_core.AFn{})
					}
				})
			}(&cljs_core.AFn{}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 10000 (trampoline (fn f [n] (if (>= n 10000) n (fn* [] (f (inc n))))) 0))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}), cljs_core.Meta.X_invoke_Arity1(cljs_core.Vary_meta.X_invoke_Arity4(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Assoc, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:a 1} (meta (vary-meta [] assoc :a 1)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil}), cljs_core.Meta.X_invoke_Arity1(cljs_core.Vary_meta.X_invoke_Arity4(cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})), cljs_core.Assoc, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:a 1, :b 2} (meta (vary-meta (with-meta [] {:b 2}) assoc :a 1)))").(string)}, ``)}))
			}
			cljs_core.Derive.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}))
			cljs_core.Derive.X_invoke_Arity2((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "square", Fqn: "cljs.core-test/square", X_hash: float64(-1438469543)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), nil}, nil}, nil}), cljs_core.Parents.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{:cljs.core-test/shape} (parents :cljs.core-test/rect))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), nil, (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), nil}, nil}, nil}), cljs_core.Ancestors.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "square", Fqn: "cljs.core-test/square", X_hash: float64(-1438469543)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{:cljs.core-test/shape :cljs.core-test/rect} (ancestors :cljs.core-test/square))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), nil, (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "square", Fqn: "cljs.core-test/square", X_hash: float64(-1438469543)}), nil}, nil}, nil}), cljs_core.Descendants.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{:cljs.core-test/rect :cljs.core-test/square} (descendants :cljs.core-test/shape))").(string)}, ``)}))
			}
			if cljs_core.Isa_QMARK_.Arity2IIB(float64(42), float64(42)) == true {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(true? (isa? 42 42))").(string)}, ``)}))
			}
			if cljs_core.Isa_QMARK_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "square", Fqn: "cljs.core-test/square", X_hash: float64(-1438469543)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})) == true {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(true? (isa? :cljs.core-test/square :cljs.core-test/shape))").(string)}, ``)}))
			}
			cljs_core.Derive.X_invoke_Arity2(reflect.TypeOf((**cljs_core.CljsCorePersistentHashSet)(nil)).Elem(), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "collection", Fqn: "cljs.core-test/collection", X_hash: float64(-440981704)}))
			if cljs_core.Isa_QMARK_.Arity2IIB(reflect.TypeOf((**cljs_core.CljsCorePersistentHashSet)(nil)).Elem(), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "collection", Fqn: "cljs.core-test/collection", X_hash: float64(-440981704)})) == true {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(true? (isa? cljs.core/PersistentHashSet :cljs.core-test/collection))").(string)}, ``)}))
			}
			if cljs_core.Isa_QMARK_.Arity2IIB(reflect.TypeOf((**cljs_core.CljsCoreIndexedSeq)(nil)).Elem(), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "collection", Fqn: "cljs.core-test/collection", X_hash: float64(-440981704)})) == false {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(false? (isa? cljs.core/IndexedSeq :cljs.core-test/collection))").(string)}, ``)}))
			}
			if cljs_core.Isa_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "square", Fqn: "cljs.core-test/square", X_hash: float64(-1438469543)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil})) == true {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(true? (isa? [:cljs.core-test/square :cljs.core-test/rect] [:cljs.core-test/shape :cljs.core-test/shape]))").(string)}, ``)}))
			}
			Bar = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1080__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1084__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__
				return (&cljs_core.CljsCoreMultiFn{cljs_core.Symbol.X_invoke_Arity2("cljs.core-test", "bar").(*cljs_core.CljsCoreSymbol), func(G__1063 *cljs_core.AFn, method_table__1080__auto__ *cljs_core.CljsCoreAtom, prefer_table__1081__auto__ *cljs_core.CljsCoreAtom, method_cache__1082__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1083__auto__ *cljs_core.CljsCoreAtom, hierarchy__1084__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1063, 2, func(x interface{}, y interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{x, y}, nil})
					})
				}(&cljs_core.AFn{}, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1084__auto__, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__})
			}()

			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}), func(G__1064 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1064, 2, func(x interface{}, y interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)})
				})
			}(&cljs_core.AFn{}))
			Bar.X_add_method_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})}, nil}), func(G__1065 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1065, 2, func(x interface{}, y interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "shape-rect", Fqn: "shape-rect", X_hash: float64(-613148403)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.Count.X_invoke_Arity1(cljs_core.Prefers.X_invoke_Arity1(Bar)).(float64) == float64(0) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(zero? (count (prefers bar)))").(string)}, ``)}))
			}
			cljs_core.Prefer_method.X_invoke_Arity3(Bar, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})}, nil}))
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Count.X_invoke_Arity1(cljs_core.Prefers.X_invoke_Arity1(Bar)).(float64)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (count (prefers bar)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)}), func() interface{} {
				var G__614 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				var G__615 = (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)})
				_, _ = G__614, G__615
				return Bar.X_invoke_Arity2(G__614, G__615)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :rect-shape (bar :cljs.core-test/rect :cljs.core-test/rect))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "rect-shape", Fqn: "rect-shape", X_hash: float64(-618116442)}), cljs_core.Apply.X_invoke_Arity2(Bar.X_get_method_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :rect-shape (apply (-get-method bar [:cljs.core-test/rect :cljs.core-test/shape]) [:cljs.core-test/rect :cljs.core-test/shape]))").(string)}, ``)}))
			}
			Nested_dispatch = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1080__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1084__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__
				return (&cljs_core.CljsCoreMultiFn{cljs_core.Symbol.X_invoke_Arity2("cljs.core-test", "nested-dispatch").(*cljs_core.CljsCoreSymbol), func(G__1066 *cljs_core.AFn, method_table__1080__auto__ *cljs_core.CljsCoreAtom, prefer_table__1081__auto__ *cljs_core.CljsCoreAtom, method_cache__1082__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1083__auto__ *cljs_core.CljsCoreAtom, hierarchy__1084__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1066, 1, func(m interface{}) interface{} {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(m))
					})
				}(&cljs_core.AFn{}, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1084__auto__, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__})
			}()

			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), func(G__1067 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1067, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1068 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1068, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__616 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil})}, nil})
				_ = G__616
				return Nested_dispatch.X_invoke_Arity1(G__616)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :nested-a (nested-dispatch {:a {:b :c}}))").(string)}, ``)}))
			}
			Nested_dispatch2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1080__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1084__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__
				return (&cljs_core.CljsCoreMultiFn{cljs_core.Symbol.X_invoke_Arity2("cljs.core-test", "nested-dispatch2").(*cljs_core.CljsCoreSymbol), cljs_core.Ffirst, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1084__auto__, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__})
			}()

			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__1069 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1069, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)})
				})
			}(&cljs_core.AFn{}))
			Nested_dispatch2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1070 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1070, 1, func(m interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-default", Fqn: "nested-default", X_hash: float64(187449106)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nested-a", Fqn: "nested-a", X_hash: float64(-411458151)}), func() interface{} {
				var G__617 = (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})}, nil})
				_ = G__617
				return Nested_dispatch2.X_invoke_Arity1(G__617)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :nested-a (nested-dispatch2 [[:a :b]]))").(string)}, ``)}))
			}
			Foo1 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1080__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1084__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__
				return (&cljs_core.CljsCoreMultiFn{cljs_core.Symbol.X_invoke_Arity2("cljs.core-test", "foo1").(*cljs_core.CljsCoreSymbol), func(G__1071 *cljs_core.AFn, method_table__1080__auto__ *cljs_core.CljsCoreAtom, prefer_table__1081__auto__ *cljs_core.CljsCoreAtom, method_cache__1082__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1083__auto__ *cljs_core.CljsCoreAtom, hierarchy__1084__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1071, 0, func(args__ ...interface{}) interface{} {
						var args = cljs_core.Seq.Arity1IQ(args__[0])
						_ = args
						return cljs_core.First.X_invoke_Arity1(args)
					})
				}(&cljs_core.AFn{}, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1084__auto__, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__})
			}()

			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__1072 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1072, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)})
				})
			}(&cljs_core.AFn{}))
			Foo1.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1073 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1073, 0, func(args__ ...interface{}) interface{} {
					var args = cljs_core.Seq.Arity1IQ(args__[0])
					_ = args
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a-return", Fqn: "a-return", X_hash: float64(-1900750605)}), func() interface{} {
				var G__618 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
				_ = G__618
				return Foo1.X_invoke_Arity1(G__618)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :a-return (foo1 :a))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default-return", Fqn: "default-return", X_hash: float64(413143669)}), func() interface{} {
				var G__619 = float64(1)
				_ = G__619
				return Foo1.X_invoke_Arity1(G__619)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :default-return (foo1 1))").(string)}, ``)}))
			}
			Area = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1080__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1084__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__
				return (&cljs_core.CljsCoreMultiFn{cljs_core.Symbol.X_invoke_Arity2("cljs.core-test", "area").(*cljs_core.CljsCoreSymbol), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Shape", Fqn: "Shape", X_hash: float64(-248980586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1084__auto__, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__})
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

			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Rect", Fqn: "Rect", X_hash: float64(-420704620)}), func(G__1074 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1074, 1, func(r interface{}) interface{} {
					return ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wd", Fqn: "wd", X_hash: float64(-183204751)}).X_invoke_Arity1(r).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ht", Fqn: "ht", X_hash: float64(214115472)}).X_invoke_Arity1(r).(float64))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "Circle", Fqn: "Circle", X_hash: float64(-158896930)}), func(G__1075 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1075, 1, func(c interface{}) interface{} {
					return (Math.PI * ((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64) * (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "radius", Fqn: "radius", X_hash: float64(-2073122258)}).X_invoke_Arity1(c).(float64)))
				})
			}(&cljs_core.AFn{}))
			Area.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1076 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1076, 1, func(x interface{}) interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)})
				})
			}(&cljs_core.AFn{}))
			R = Rect.X_invoke_Arity2(float64(4), float64(13)).(cljs_core.CljsCoreIMap)

			C = Circle.X_invoke_Arity1(float64(12)).(cljs_core.CljsCoreIMap)

			if cljs_core.X_EQ_.Arity2IIB(float64(52), func() interface{} {
				var G__620 = R
				_ = G__620
				return Area.X_invoke_Arity1(G__620)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 52 (area r))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "oops", Fqn: "oops", X_hash: float64(-1033827083)}), func() interface{} {
				var G__621 = cljs_core.CljsCorePersistentArrayMap_EMPTY
				_ = G__621
				return Area.X_invoke_Arity1(G__621)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :oops (area {}))").(string)}, ``)}))
			}
			Foo2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1080__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1084__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__
				return (&cljs_core.CljsCoreMultiFn{cljs_core.Symbol.X_invoke_Arity2("cljs.core-test", "foo2").(*cljs_core.CljsCoreSymbol), func(G__1077 *cljs_core.AFn, method_table__1080__auto__ *cljs_core.CljsCoreAtom, prefer_table__1081__auto__ *cljs_core.CljsCoreAtom, method_cache__1082__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1083__auto__ *cljs_core.CljsCoreAtom, hierarchy__1084__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1077, 0, func() interface{} {
						return nil
					})
				}(&cljs_core.AFn{}, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1084__auto__, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__})
			}()

			Foo2.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1078 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1078, 0, func() interface{} {
					return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), func() interface{} {
				return Foo2.X_invoke_Arity0()
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :foo (foo2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Count.X_invoke_Arity1(cljs_core.Methods.X_invoke_Arity1(Bar)).(float64)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (count (methods bar)))").(string)}, ``)}))
			}
			cljs_core.Remove_method.X_invoke_Arity2(Bar, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "rect", Fqn: "cljs.core-test/rect", X_hash: float64(1940896440)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "shape", Fqn: "cljs.core-test/shape", X_hash: float64(-118750990)})}, nil}))
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Count.X_invoke_Arity1(cljs_core.Methods.X_invoke_Arity1(Bar)).(float64)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (count (methods bar)))").(string)}, ``)}))
			}
			cljs_core.Remove_all_methods.X_invoke_Arity1(Bar)
			if cljs_core.Count.X_invoke_Arity1(cljs_core.Methods.X_invoke_Arity1(Bar)).(float64) == float64(0) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(zero? (count (methods bar)))").(string)}, ``)}))
			}
			Apply_multi_test = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1080__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1084__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__
				return (&cljs_core.CljsCoreMultiFn{cljs_core.Symbol.X_invoke_Arity2("cljs.core-test", "apply-multi-test").(*cljs_core.CljsCoreSymbol), func(G__1079 *cljs_core.AFn, method_table__1080__auto__ *cljs_core.CljsCoreAtom, prefer_table__1081__auto__ *cljs_core.CljsCoreAtom, method_cache__1082__auto__ *cljs_core.CljsCoreAtom, cached_hierarchy__1083__auto__ *cljs_core.CljsCoreAtom, hierarchy__1084__auto__ interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1079, 3, func(___ interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}) interface{} {
						return float64(0)
					}, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
						return float64(0)
					})
				}(&cljs_core.AFn{}, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1084__auto__, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__})
			}()

			Apply_multi_test.X_add_method_Arity3(float64(0), func(G__1080 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1080, 2, func(x interface{}) interface{} {
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [:three (quote (2))] (apply apply-multi-test [0 1 2]))").(string)}, ``)}))
			}
			No_dispatch_value = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1080__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1084__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__
				return (&cljs_core.CljsCoreMultiFn{cljs_core.Symbol.X_invoke_Arity2("cljs.core-test", "no-dispatch-value").(*cljs_core.CljsCoreSymbol), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "test", Fqn: "test", X_hash: float64(577538877)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1084__auto__, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__})
			}()

			func() {
				defer func() {
					if e622 := recover(); e622 != nil {
						if cljs_core.Value_(e622).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
							{
								var e_1082 = e622
								_ = e_1082
								if cljs_core.Not_EQ_.Arity2IIB(float64(-1), cljs_core.Native_invoke_instance_method.X_invoke_Arity3(cljs_core.Native_get_instance_field.X_invoke_Arity2(e_1082, "Message"), "IndexOf", []interface{}{"cljs.core-test/no-dispatch-value"})) {
								} else {
									panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not= -1 (.indexOf (.-message e) \"cljs.core-test/no-dispatch-value\"))").(string)}, ``)}))
								}
							}
						} else {
							panic(e622)

						}
					}
				}()
				{
					{
						var G__623_1083 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "test", Fqn: "test", X_hash: float64(577538877)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "test", Fqn: "test", X_hash: float64(577538877)})}, nil})
						_ = G__623_1083
						No_dispatch_value.X_invoke_Arity1(G__623_1083)
					}
				}
			}()
			My_map_hierarchy = cljs_core.Atom.X_invoke_Arity1(cljs_core.Derive.X_invoke_Arity3(cljs_core.Derive.X_invoke_Arity3(cljs_core.Derive.X_invoke_Arity3(cljs_core.Make_hierarchy.X_invoke_Arity0().(cljs_core.CljsCoreIMap), cljs_core.Type_.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)})), cljs_core.Type_.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_EMPTY), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)})), cljs_core.Type_.X_invoke_Arity1(cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))), (&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)}))).(*cljs_core.CljsCoreAtom)

			My_map_QMARK_ = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1080__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1084__auto__ = cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), My_map_hierarchy}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__
				return (&cljs_core.CljsCoreMultiFn{cljs_core.Symbol.X_invoke_Arity2("cljs.core-test", "my-map?").(*cljs_core.CljsCoreSymbol), cljs_core.Type_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1084__auto__, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__})
			}()

			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: "cljs.core-test", Name: "map", Fqn: "cljs.core-test/map", X_hash: float64(-1007238055)}), func(G__1084 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1084, 1, func(___ interface{}) interface{} {
					return true
				})
			}(&cljs_core.AFn{}))
			My_map_QMARK_.X_add_method_Arity3((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), func(G__1085 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1085, 1, func(___ interface{}) interface{} {
					return false
				})
			}(&cljs_core.AFn{}))
			{
				var seq__624_1086 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__625_1087 interface{} = nil
				var count__626_1088 = float64(0)
				var i__627_1089 = float64(0)
				_, _, _, _ = seq__624_1086, chunk__625_1087, count__626_1088, i__627_1089
				for {
					if i__627_1089 < count__626_1088 {
						{
							var m_1090 = cljs_core.Decorate_(chunk__625_1087).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__627_1089)
							_ = m_1090
							if cljs_core.Truth_(func() interface{} {
								var G__628 = m_1090
								_ = G__628
								return My_map_QMARK_.X_invoke_Arity1(G__628)
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(my-map? m)").(string)}, ``)}))
							}
							seq__624_1086, chunk__625_1087, count__626_1088, i__627_1089 = seq__624_1086, chunk__625_1087, count__626_1088, (i__627_1089 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1091 = cljs_core.Seq.Arity1IQ(seq__624_1086)
							_ = temp__4388__auto___1091
							if cljs_core.Truth_(temp__4388__auto___1091) {
								{
									var seq__624_1092___1 = temp__4388__auto___1091
									_ = seq__624_1092___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__624_1092___1) {
										{
											var c__970__auto___1093 = cljs_core.Chunk_first.X_invoke_Arity1(seq__624_1092___1)
											_ = c__970__auto___1093
											seq__624_1086, chunk__625_1087, count__626_1088, i__627_1089 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__624_1092___1), c__970__auto___1093, cljs_core.Count.X_invoke_Arity1(c__970__auto___1093).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_1094 = cljs_core.First.X_invoke_Arity1(seq__624_1092___1)
											_ = m_1094
											if cljs_core.Truth_(func() interface{} {
												var G__629 = m_1094
												_ = G__629
												return My_map_QMARK_.X_invoke_Arity1(G__629)
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(my-map? m)").(string)}, ``)}))
											}
											seq__624_1086, chunk__625_1087, count__626_1088, i__627_1089 = cljs_core.Next.Arity1IQ(seq__624_1092___1), nil, float64(0), float64(0)
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
				var seq__630_1095 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentVector_EMPTY, float64(1), "asdf", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}))
				var chunk__631_1096 interface{} = nil
				var count__632_1097 = float64(0)
				var i__633_1098 = float64(0)
				_, _, _, _ = seq__630_1095, chunk__631_1096, count__632_1097, i__633_1098
				for {
					if i__633_1098 < count__632_1097 {
						{
							var not_m_1099 = cljs_core.Decorate_(chunk__631_1096).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__633_1098)
							_ = not_m_1099
							if cljs_core.Not.Arity1IB(func() interface{} {
								var G__634 = not_m_1099
								_ = G__634
								return My_map_QMARK_.X_invoke_Arity1(G__634)
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (my-map? not-m))").(string)}, ``)}))
							}
							seq__630_1095, chunk__631_1096, count__632_1097, i__633_1098 = seq__630_1095, chunk__631_1096, count__632_1097, (i__633_1098 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1100 = cljs_core.Seq.Arity1IQ(seq__630_1095)
							_ = temp__4388__auto___1100
							if cljs_core.Truth_(temp__4388__auto___1100) {
								{
									var seq__630_1101___1 = temp__4388__auto___1100
									_ = seq__630_1101___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__630_1101___1) {
										{
											var c__970__auto___1102 = cljs_core.Chunk_first.X_invoke_Arity1(seq__630_1101___1)
											_ = c__970__auto___1102
											seq__630_1095, chunk__631_1096, count__632_1097, i__633_1098 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__630_1101___1), c__970__auto___1102, cljs_core.Count.X_invoke_Arity1(c__970__auto___1102).(float64), float64(0)
											continue
										}
									} else {
										{
											var not_m_1103 = cljs_core.First.X_invoke_Arity1(seq__630_1101___1)
											_ = not_m_1103
											if cljs_core.Not.Arity1IB(func() interface{} {
												var G__635 = not_m_1103
												_ = G__635
												return My_map_QMARK_.X_invoke_Arity1(G__635)
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (my-map? not-m))").(string)}, ``)}))
											}
											seq__630_1095, chunk__631_1096, count__632_1097, i__633_1098 = cljs_core.Next.Arity1IQ(seq__630_1101___1), nil, float64(0), float64(0)
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 0 10 3) (list 0 3 6 9))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(10), float64(3)).(*cljs_core.CljsCoreRange)).(float64), float64(4)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count (range 0 10 3)) 4)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(-10), float64(-3)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(-9))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-6))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 0 -10 -3) (list 0 -3 -6 -9))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(-10), float64(-3)).(*cljs_core.CljsCoreRange)).(float64), float64(4)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count (range 0 -10 -3)) 4)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(-10), float64(10), float64(3)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(5))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-4))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-10))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range -10 10 3) (list -10 -7 -4 -1 2 5 8))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(-10), float64(10), float64(3)).(*cljs_core.CljsCoreRange)).(float64), float64(7)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count (range -10 10 3)) 7)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(1), float64(1)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 0 1 1) (list 0))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(-3), float64(-1)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(-2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(-1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(0))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 0 -3 -1) (list 0 -1 -2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(3), float64(0), float64(-1)).(*cljs_core.CljsCoreRange), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 3 0 -1) (list 3 2 1))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(10), float64(-1)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 0 10 -1) (list))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(1), float64(0)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 0 1 0) (list))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(10), float64(0), float64(1)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 10 0 1) (list))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(0), float64(0)).(*cljs_core.CljsCoreRange), cljs_core.CljsCoreList_EMPTY) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (range 0 0 0) (list))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(10), float64(-1)).(*cljs_core.CljsCoreRange)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count (range 0 10 -1)) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(1), float64(0)).(*cljs_core.CljsCoreRange)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count (range 0 1 0)) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(10), float64(0), float64(1)).(*cljs_core.CljsCoreRange)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count (range 10 0 1)) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity3(float64(0), float64(0), float64(0)).(*cljs_core.CljsCoreRange)).(float64), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count (range 0 0 0)) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.Range_.X_invoke_Arity3(float64(1), float64(0), float64(0)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (take 3 (range 1 0 0)) (list 1 1 1))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Take.X_invoke_Arity2(float64(3), cljs_core.Range_.X_invoke_Arity3(float64(3), float64(1), float64(0)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (take 3 (range 3 1 0)) (list 3 3 3))").(string)}, ``)}))
			}
			{
				var pv_1104 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange))
				_ = pv_1104
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity2(pv_1104, float64(96)), float64(96)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (nth pv 96) 96)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Nth.X_invoke_Arity3(pv_1104, float64(97), nil), nil) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (nth pv 97 nil) nil)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__636 = float64(96)
					_ = G__636
					return pv_1104.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__636)
				}(), float64(96)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pv 96) 96)").(string)}, ``)}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(cljs_core.CljsCorePersistentVector_EMPTY)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (rseq []))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reverse.X_invoke_Arity1(pv_1104), cljs_core.Rseq.Arity1IQ(pv_1104)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (reverse pv) (rseq pv))").(string)}, ``)}))
				}
			}
			{
				var pv_1105 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(33)).(*cljs_core.CljsCoreRange))
				_ = pv_1105
				if cljs_core.X_EQ_.Arity2IIB(pv_1105, cljs_core.Conj.X_invoke_Arity2(cljs_core.Conj.X_invoke_Arity2(cljs_core.Pop.X_invoke_Arity1(cljs_core.Pop.X_invoke_Arity1(pv_1105)), float64(31)), float64(32))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= pv (-> pv pop pop (conj 31) (conj 32)))").(string)}, ``)}))
				}
			}
			{
				var stack1_1106 = cljs_core.Pop.X_invoke_Arity1(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(97)).(*cljs_core.CljsCoreRange)))
				var stack2_1107 = cljs_core.Pop.X_invoke_Arity1(stack1_1106)
				_, _ = stack1_1106, stack2_1107
				if cljs_core.X_EQ_.Arity2IIB(float64(95), cljs_core.Peek.X_invoke_Arity1(stack1_1106)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 95 (peek stack1))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(94), cljs_core.Peek.X_invoke_Arity1(stack2_1107)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 94 (peek stack2))").(string)}, ``)}))
				}
			}
			{
				var sentinel_1108 = cljs_core.Rand.Arity0F()
				_ = sentinel_1108
				if reflect.DeepEqual(sentinel_1108, func() (return__1109 interface{}) {
					defer func() {
						if e637 := recover(); e637 != nil {
							if cljs_core.Value_(e637).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var ___ = e637
									_ = ___
									return__1109 = sentinel_1108
								}
							} else {
								panic(e637)

							}
						}
					}()
					{
						return cljs_core.CljsCorePersistentVector_EMPTY.X_invoke_Arity1(float64(0))
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? sentinel (try ([] 0) (catch js/Error _ sentinel)))").(string)}, ``)}))
				}
			}
			{
				var v1_1110 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				var v2_1111 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(5)).(*cljs_core.CljsCoreRange))
				var s_1112 = cljs_core.Subvec.X_invoke_Arity3(v1_1110, float64(2), float64(8)).(*cljs_core.CljsCoreSubvec)
				_, _, _ = v1_1110, v2_1111, s_1112
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic(s_1112, cljs_core.Subvec.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2(v1_1110, float64(2)).(*cljs_core.CljsCoreSubvec), float64(0), float64(6)).(*cljs_core.CljsCoreSubvec), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Take.X_invoke_Arity2(float64(6), cljs_core.Drop.X_invoke_Arity2(float64(2), v1_1110).(*cljs_core.CljsCoreLazySeq)).(*cljs_core.CljsCoreLazySeq)}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= s (-> v1 (subvec 2) (subvec 0 6)) (->> v1 (drop 2) (take 6)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Count.X_invoke_Arity1(s_1112).(float64)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 6 (count s))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6)}, nil}), cljs_core.Pop.X_invoke_Arity1(s_1112)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 3 4 5 6] (pop s))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(7), cljs_core.Peek.X_invoke_Arity1(s_1112)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 7 (peek s))").(string)}, ``)}))
				}
				if cljs_core.Truth_(cljs_core.X_EQ_.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(1)}, nil}), cljs_core.Assoc.X_invoke_Arity3(s_1112, float64(6), float64(1)), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Conj.X_invoke_Arity2(s_1112, float64(1))}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 3 4 5 6 7 1] (assoc s 6 1) (conj s 1))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(27), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, s_1112)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 27 (reduce + s))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(s_1112, cljs_core.Vec.X_invoke_Arity1(s_1112)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= s (vec s))").(string)}, ``)}))
				}
				{
					var m_1113 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), float64(1)}, nil})
					_ = m_1113
					if cljs_core.X_EQ_.Arity2IIB(m_1113, cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(s_1112, m_1113))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m (meta (with-meta s m)))").(string)}, ``)}))
					}
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1114 interface{}) {
					defer func() {
						if e638 := recover(); e638 != nil {
							if cljs_core.Value_(e638).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e638
									_ = e
									return__1114 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e638)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1111, float64(0), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (subvec v2 0 6) (catch js/Error e :fail)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1115 interface{}) {
					defer func() {
						if e639 := recover(); e639 != nil {
							if cljs_core.Value_(e639).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e639
									_ = e
									return__1115 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e639)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1111, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1116 interface{}) {
					defer func() {
						if e640 := recover(); e640 != nil {
							if cljs_core.Value_(e640).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e640
									_ = e
									return__1116 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e640)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1111, float64(6), float64(10)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (subvec v2 6 10) (catch js/Error e :fail)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1117 interface{}) {
					defer func() {
						if e641 := recover(); e641 != nil {
							if cljs_core.Value_(e641).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e641
									_ = e
									return__1117 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e641)

							}
						}
					}()
					{
						return cljs_core.Subvec.X_invoke_Arity3(v2_1111, float64(3), float64(6)).(*cljs_core.CljsCoreSubvec)
					}
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (subvec v2 3 6) (catch js/Error e :fail)))").(string)}, ``)}))
				}
				if reflect.DeepEqual(v1_1110, cljs_core.Subvec.X_invoke_Arity3(s_1112, float64(1), float64(4)).(*cljs_core.CljsCoreSubvec).V) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? v1 (.-v (subvec s 1 4)))").(string)}, ``)}))
				}
				{
					var sentinel_1118 = cljs_core.Rand.Arity0F()
					var s_1119___1 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2))
					_, _ = sentinel_1118, s_1119___1
					if reflect.DeepEqual(sentinel_1118, func() (return__1120 interface{}) {
						defer func() {
							if e642 := recover(); e642 != nil {
								if cljs_core.Value_(e642).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var ___ = e642
										_ = ___
										return__1120 = sentinel_1118
									}
								} else {
									panic(e642)

								}
							}
						}()
						{
							{
								var G__643 = float64(-1)
								_ = G__643
								return s_1119___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__643)
							}
						}
					}()) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? sentinel (try (s -1) (catch js/Error _ sentinel)))").(string)}, ``)}))
					}
					if reflect.DeepEqual(sentinel_1118, func() (return__1121 interface{}) {
						defer func() {
							if e644 := recover(); e644 != nil {
								if cljs_core.Value_(e644).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var ___ = e644
										_ = ___
										return__1121 = sentinel_1118
									}
								} else {
									panic(e644)

								}
							}
						}()
						{
							{
								var G__645 = float64(1)
								_ = G__645
								return s_1119___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__645)
							}
						}
					}()) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? sentinel (try (s 1) (catch js/Error _ sentinel)))").(string)}, ``)}))
					}
				}
				{
					var sv1_1122 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(2)).(*cljs_core.CljsCoreSubvec)
					var sv2_1123 = cljs_core.Subvec.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(3)}, nil}), float64(1), float64(1)).(*cljs_core.CljsCoreSubvec)
					_, _ = sv1_1122, sv2_1123
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(sv1_1122), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(*cljs_core.CljsCoreList)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rseq sv1) (quote (1)))").(string)}, ``)}))
					}
					if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(sv2_1123)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (rseq sv2))").(string)}, ``)}))
					}
				}
			}
			{
				var v1_1124 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(15), float64(48)).(*cljs_core.CljsCoreRange))
				var v2_1125 = cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity2(float64(40), float64(57)).(*cljs_core.CljsCoreRange))
				var v1_1126___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Pop_BANG_.X_invoke_Arity1(cljs_core.Transient.X_invoke_Arity1(v1_1124)), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v2_1127___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(v2_1125), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})), float64(0), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})))
				var v_1128 = cljs_core.Into.X_invoke_Arity2(v1_1126___1, v2_1127___1)
				_, _, _, _, _ = v1_1124, v2_1125, v1_1126___1, v2_1127___1, v_1128
				if cljs_core.X_EQ_.Arity2IIB(v_1128, cljs_core.Vec.X_invoke_Arity1(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(16), float64(47)).(*cljs_core.CljsCoreRange), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), cljs_core.Range_.X_invoke_Arity2(float64(41), float64(57)).(*cljs_core.CljsCoreRange), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})})).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= v (vec (concat [:quux] (range 16 47) [:foo] [:quux] (range 41 57) [:bar])))").(string)}, ``)}))
				}
			}
			{
				var v_1129 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)
				var xs_1130 interface{} = cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)
				_, _ = v_1129, xs_1130
				for {
					{
						var temp__4386__auto___1131 = cljs_core.First.X_invoke_Arity1(xs_1130)
						_ = temp__4386__auto___1131
						if cljs_core.Truth_(temp__4386__auto___1131) {
							{
								var x_1132 = temp__4386__auto___1131
								_ = x_1132
								v_1129, xs_1130 = func() interface{} {
									var pred__646 = func(G__1133 *cljs_core.AFn, v_1129 interface{}, xs_1130 interface{}, x_1132 interface{}, temp__4386__auto___1131 interface{}) *cljs_core.AFn {
										return cljs_core.Fn(G__1133, 2, func(p1__67_SHARP_ interface{}, p2__68_SHARP_ interface{}) interface{} {
											{
												var G__649 = cljs_core.Mod.X_invoke_Arity2(p2__68_SHARP_, float64(3)).(float64)
												_ = G__649
												return p1__67_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__649)
											}
										})
									}(&cljs_core.AFn{}, v_1129, xs_1130, x_1132, temp__4386__auto___1131)
									var expr__647 = x_1132
									_, _ = pred__646, expr__647
									if cljs_core.Truth_(pred__646.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(0), nil, float64(2), nil}, nil}, nil}), expr__647)) {
										return cljs_core.Conj_BANG_.X_invoke_Arity2(v_1129, x_1132)
									} else {
										if cljs_core.Truth_(pred__646.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(1), nil}, nil}, nil}), expr__647)) {
											return cljs_core.Assoc_BANG_.X_invoke_Arity3(v_1129, cljs_core.Count.X_invoke_Arity1(v_1129).(float64), x_1132)
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("No matching clause: ").(string), cljs_core.Str.X_invoke_Arity1(expr__647).(string)}, ``)}))
										}
									}
								}(), cljs_core.Next.Arity1IQ(xs_1130)
								continue
							}
						} else {
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Vec.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)), cljs_core.Persistent_BANG_.X_invoke_Arity1(v_1129)) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (vec (range 100)) (persistent! v))").(string)}, ``)}))
							}
						}
					}
					break
				}
			}
			{
				var m1_1134 interface{} = cljs_core.CljsCorePersistentHashMap_EMPTY
				var m2_1135 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_EMPTY)
				var i_1136 = float64(0)
				_, _, _ = m1_1134, m2_1135, i_1136
				for {
					if i_1136 < float64(100) {
						m1_1134, m2_1135, i_1136 = cljs_core.Assoc.X_invoke_Arity3(m1_1134, i_1136, i_1136), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_1135, i_1136, i_1136), (i_1136 + float64(1))
						continue
					} else {
						{
							var m2_1137___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_1135)
							_ = m2_1137___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1134).(float64), float64(100)) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m1) 100)").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1137___1).(float64), float64(100)) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m2) 100)").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_1134, m2_1137___1) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m1 m2)").(string)}, ``)}))
							}
							{
								var i_1138___1 = float64(0)
								_ = i_1138___1
								for {
									if i_1138___1 < float64(100) {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__650 = i_1138___1
											_ = G__650
											return m1_1134.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__650)
										}(), i_1138___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (m1 i) i)").(string)}, ``)}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__651 = i_1138___1
											_ = G__651
											return m2_1137___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__651)
										}(), i_1138___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (m2 i) i)").(string)}, ``)}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_1134, i_1138___1), i_1138___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m1 i) i)").(string)}, ``)}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_1137___1, i_1138___1), i_1138___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m2 i) i)").(string)}, ``)}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_1134, i_1138___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? m1 i)").(string)}, ``)}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_1137___1, i_1138___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? m2 i)").(string)}, ``)}))
										}
										i_1138___1 = (i_1138___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_1134))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map vector (range 100) (range 100)) (sort-by first (seq m1)))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_1137___1))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map vector (range 100) (range 100)) (sort-by first (seq m2)))").(string)}, ``)}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_1134, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? (dissoc m1 3) 3))").(string)}, ``)}))
							}
						}
					}
					break
				}
			}
			{
				var m_1139 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_1139
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1139).(float64), float64(7)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 7)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1139, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})").(string)}, ``)}))
				}
			}
			{
				var m_1140 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_1140
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1140).(float64), float64(11)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 11)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1140, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})").(string)}, ``)}))
				}
			}
			{
				var m_1141 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_1141
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1141).(float64), float64(11)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 11)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1141, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})").(string)}, ``)}))
				}
			}
			{
				var tm_1142 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_1142
				{
					var tm_1143___1 interface{} = tm_1142
					var ks_1144 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_1143___1, ks_1144
					for {
						{
							var temp__4386__auto___1145 = cljs_core.First.X_invoke_Arity1(ks_1144)
							_ = temp__4386__auto___1145
							if cljs_core.Truth_(temp__4386__auto___1145) {
								{
									var k_1146 = temp__4386__auto___1145
									_ = k_1146
									tm_1143___1, ks_1144 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1143___1, k_1146), cljs_core.Next.Arity1IQ(ks_1144)
									continue
								}
							} else {
								{
									var m_1147 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1143___1)
									_ = m_1147
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1147).(float64), float64(7)) {
									} else {
										panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 7)").(string)}, ``)}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_1147, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
									} else {
										panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})").(string)}, ``)}))
									}
								}
							}
						}
						break
					}
				}
			}
			{
				var tm_1148 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_1148
				{
					var seq__652_1149 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__653_1150 interface{} = nil
					var count__654_1151 = float64(0)
					var i__655_1152 = float64(0)
					_, _, _, _ = seq__652_1149, chunk__653_1150, count__654_1151, i__655_1152
					for {
						if i__655_1152 < count__654_1151 {
							{
								var k_1153 = cljs_core.Decorate_(chunk__653_1150).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__655_1152)
								_ = k_1153
								if cljs_core.X_EQ_.Arity2IIB(k_1153, cljs_core.Get.X_invoke_Arity2(tm_1148, k_1153)) {
								} else {
									panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= k (get tm k))").(string)}, ``)}))
								}
								seq__652_1149, chunk__653_1150, count__654_1151, i__655_1152 = seq__652_1149, chunk__653_1150, count__654_1151, (i__655_1152 + float64(1))
								continue
							}
						} else {
							{
								var temp__4388__auto___1154 = cljs_core.Seq.Arity1IQ(seq__652_1149)
								_ = temp__4388__auto___1154
								if cljs_core.Truth_(temp__4388__auto___1154) {
									{
										var seq__652_1155___1 = temp__4388__auto___1154
										_ = seq__652_1155___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__652_1155___1) {
											{
												var c__970__auto___1156 = cljs_core.Chunk_first.X_invoke_Arity1(seq__652_1155___1)
												_ = c__970__auto___1156
												seq__652_1149, chunk__653_1150, count__654_1151, i__655_1152 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__652_1155___1), c__970__auto___1156, cljs_core.Count.X_invoke_Arity1(c__970__auto___1156).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_1157 = cljs_core.First.X_invoke_Arity1(seq__652_1155___1)
												_ = k_1157
												if cljs_core.X_EQ_.Arity2IIB(k_1157, cljs_core.Get.X_invoke_Arity2(tm_1148, k_1157)) {
												} else {
													panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= k (get tm k))").(string)}, ``)}))
												}
												seq__652_1149, chunk__653_1150, count__654_1151, i__655_1152 = cljs_core.Next.Arity1IQ(seq__652_1155___1), nil, float64(0), float64(0)
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
					var m_1158 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1148)
					_ = m_1158
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1159 float64) {
						defer func() {
							if e656 := recover(); e656 != nil {
								if cljs_core.Value_(e656).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e656
										_ = e
										return__1159 = float64(2)
									}
								} else {
									panic(e656)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1148, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1160 float64) {
						defer func() {
							if e657 := recover(); e657 != nil {
								if cljs_core.Value_(e657).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e657
										_ = e
										return__1160 = float64(2)
									}
								} else {
									panic(e657)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_1148, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1161 float64) {
						defer func() {
							if e658 := recover(); e658 != nil {
								if cljs_core.Value_(e658).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e658
										_ = e
										return__1161 = float64(2)
									}
								} else {
									panic(e658)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1148)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try (persistent! tm) 1 (catch js/Error e 2)))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1162 float64) {
						defer func() {
							if e659 := recover(); e659 != nil {
								if cljs_core.Value_(e659).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e659
										_ = e
										return__1162 = float64(2)
									}
								} else {
									panic(e659)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_1148)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try (count tm) 1 (catch js/Error e 2)))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_1158, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})").(string)}, ``)}))
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
				var m_1163 = cljs_core.Assoc.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentHashMap_EMPTY, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_ = m_1163
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1163, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m fixed-hash-foo) 1)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1163, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m fixed-hash-bar) 2)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1163).(float64), float64(2)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 2)").(string)}, ``)}))
				}
				{
					var m_1164___1 = cljs_core.Dissoc.X_invoke_Arity2(m_1163, Fixed_hash_foo)
					_ = m_1164___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1164___1, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m fixed-hash-bar) 2)").(string)}, ``)}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1164___1, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? m fixed-hash-foo))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1164___1).(float64), float64(1)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 1)").(string)}, ``)}))
					}
				}
			}
			{
				var m_1165 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_1166___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m_1165, Fixed_hash_foo, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{Fixed_hash_bar, float64(2)}))
				_, _ = m_1165, m_1166___1
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1166___1).(float64), float64(102)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 102)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1166___1, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m fixed-hash-foo) 1)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1166___1, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m fixed-hash-bar) 2)").(string)}, ``)}))
				}
				{
					var m_1167___2 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_1166___1, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_1167___2
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1167___2, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m fixed-hash-bar) 2)").(string)}, ``)}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1167___2, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? m fixed-hash-foo))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1167___2).(float64), float64(98)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 98)").(string)}, ``)}))
					}
				}
			}
			{
				var m_1168 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Zipmap.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)))
				var m_1169___1 = cljs_core.Transient.X_invoke_Arity1(m_1168)
				var m_1170___2 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_1169___1, Fixed_hash_foo, float64(1))
				var m_1171___3 = cljs_core.Assoc_BANG_.X_invoke_Arity3(m_1170___2, Fixed_hash_bar, float64(2))
				var m_1172___4 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m_1171___3)
				_, _, _, _, _ = m_1168, m_1169___1, m_1170___2, m_1171___3, m_1172___4
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1172___4).(float64), float64(102)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 102)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1172___4, Fixed_hash_foo), float64(1)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m fixed-hash-foo) 1)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1172___4, Fixed_hash_bar), float64(2)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m fixed-hash-bar) 2)").(string)}, ``)}))
				}
				{
					var m_1173___5 = cljs_core.Dissoc.X_invoke_ArityVariadic(m_1172___4, float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7), Fixed_hash_foo}))
					_ = m_1173___5
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m_1173___5, Fixed_hash_bar), float64(2)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m fixed-hash-bar) 2)").(string)}, ``)}))
					}
					if !(cljs_core.Contains_QMARK_.Arity2IIB(m_1173___5, Fixed_hash_foo)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? m fixed-hash-foo))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1173___5).(float64), float64(98)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 98)").(string)}, ``)}))
					}
				}
			}
			Array_map_conversion_threshold = cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD

			{
				var m1_1174 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_1175 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY)
				var i_1176 = float64(0)
				_, _, _ = m1_1174, m2_1175, i_1176
				for {
					if i_1176 < Array_map_conversion_threshold {
						m1_1174, m2_1175, i_1176 = cljs_core.Assoc.X_invoke_Arity3(m1_1174, i_1176, i_1176), cljs_core.Assoc_BANG_.X_invoke_Arity3(m2_1175, i_1176, i_1176), (i_1176 + float64(1))
						continue
					} else {
						{
							var m2_1177___1 = cljs_core.Persistent_BANG_.X_invoke_Arity1(m2_1175)
							_ = m2_1177___1
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1174).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m1) array-map-conversion-threshold)").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1177___1).(float64), Array_map_conversion_threshold) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m2) array-map-conversion-threshold)").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB(m1_1174, m2_1177___1) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m1 m2)").(string)}, ``)}))
							}
							{
								var i_1178___1 = float64(0)
								_ = i_1178___1
								for {
									if i_1178___1 < Array_map_conversion_threshold {
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__660 = i_1178___1
											_ = G__660
											return m1_1174.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__660)
										}(), i_1178___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (m1 i) i)").(string)}, ``)}))
										}
										if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
											var G__661 = i_1178___1
											_ = G__661
											return m2_1177___1.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__661)
										}(), i_1178___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (m2 i) i)").(string)}, ``)}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m1_1174, i_1178___1), i_1178___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m1 i) i)").(string)}, ``)}))
										}
										if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(m2_1177___1, i_1178___1), i_1178___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get m2 i) i)").(string)}, ``)}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m1_1174, i_1178___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? m1 i)").(string)}, ``)}))
										}
										if cljs_core.Contains_QMARK_.Arity2IIB(m2_1177___1, i_1178___1) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(contains? m2 i)").(string)}, ``)}))
										}
										i_1178___1 = (i_1178___1 + float64(1))
										continue
									} else {
									}
									break
								}
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m1_1174))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m1)))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(Array_map_conversion_threshold).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Sort_by.X_invoke_Arity2(cljs_core.First, cljs_core.Seq.Arity1IQ(m2_1177___1))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map vector (range array-map-conversion-threshold) (range array-map-conversion-threshold)) (sort-by first (seq m2)))").(string)}, ``)}))
							}
							if !(cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m1_1174, float64(3)), float64(3))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not (contains? (dissoc m1 3) 3))").(string)}, ``)}))
							}
						}
					}
					break
				}
			}
			{
				var m_1179 = cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)}))
				_ = m_1179
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1179).(float64), float64(7)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 7)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1179, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})").(string)}, ``)}))
				}
			}
			{
				var m_1180 = cljs_core.Conj.X_invoke_Arity2(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))
				_ = m_1180
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1180).(float64), float64(11)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 11)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1180, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})").(string)}, ``)}))
				}
			}
			{
				var m_1181 = cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))
				_ = m_1181
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1181).(float64), float64(11)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 11)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1181, cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2), float64(9), float64(5), float64(8)}, []interface{}{float64(0), float64(7), float64(1), float64(4), float64(6), float64(3), float64(1), float64(2), float64(9), float64(5), float64(8)}).(*cljs_core.CljsCorePersistentHashMap)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m {0 0, 7 7, 1 1, 4 4, 6 6, 3 3, :foo 1, 2 2, 9 9, 5 5, 8 8})").(string)}, ``)}))
				}
			}
			{
				var tm_1182 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)))
				_ = tm_1182
				{
					var tm_1183___1 interface{} = tm_1182
					var ks_1184 interface{} = (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(5), float64(7)}, nil})
					_, _ = tm_1183___1, ks_1184
					for {
						{
							var temp__4386__auto___1185 = cljs_core.First.X_invoke_Arity1(ks_1184)
							_ = temp__4386__auto___1185
							if cljs_core.Truth_(temp__4386__auto___1185) {
								{
									var k_1186 = temp__4386__auto___1185
									_ = k_1186
									tm_1183___1, ks_1184 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1183___1, k_1186), cljs_core.Next.Arity1IQ(ks_1184)
									continue
								}
							} else {
								{
									var m_1187 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1183___1)
									_ = m_1187
									if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1187).(float64), float64(7)) {
									} else {
										panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 7)").(string)}, ``)}))
									}
									if cljs_core.X_EQ_.Arity2IIB(m_1187, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
									} else {
										panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})").(string)}, ``)}))
									}
								}
							}
						}
						break
					}
				}
			}
			{
				var tm_1188 = cljs_core.Transient.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)), float64(3), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(7)})))
				_ = tm_1188
				{
					var seq__662_1189 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1), float64(2), float64(4), float64(6), float64(8), float64(9)}, nil}))
					var chunk__663_1190 interface{} = nil
					var count__664_1191 = float64(0)
					var i__665_1192 = float64(0)
					_, _, _, _ = seq__662_1189, chunk__663_1190, count__664_1191, i__665_1192
					for {
						if i__665_1192 < count__664_1191 {
							{
								var k_1193 = cljs_core.Decorate_(chunk__663_1190).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__665_1192)
								_ = k_1193
								if cljs_core.X_EQ_.Arity2IIB(k_1193, cljs_core.Get.X_invoke_Arity2(tm_1188, k_1193)) {
								} else {
									panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= k (get tm k))").(string)}, ``)}))
								}
								seq__662_1189, chunk__663_1190, count__664_1191, i__665_1192 = seq__662_1189, chunk__663_1190, count__664_1191, (i__665_1192 + float64(1))
								continue
							}
						} else {
							{
								var temp__4388__auto___1194 = cljs_core.Seq.Arity1IQ(seq__662_1189)
								_ = temp__4388__auto___1194
								if cljs_core.Truth_(temp__4388__auto___1194) {
									{
										var seq__662_1195___1 = temp__4388__auto___1194
										_ = seq__662_1195___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__662_1195___1) {
											{
												var c__970__auto___1196 = cljs_core.Chunk_first.X_invoke_Arity1(seq__662_1195___1)
												_ = c__970__auto___1196
												seq__662_1189, chunk__663_1190, count__664_1191, i__665_1192 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__662_1195___1), c__970__auto___1196, cljs_core.Count.X_invoke_Arity1(c__970__auto___1196).(float64), float64(0)
												continue
											}
										} else {
											{
												var k_1197 = cljs_core.First.X_invoke_Arity1(seq__662_1195___1)
												_ = k_1197
												if cljs_core.X_EQ_.Arity2IIB(k_1197, cljs_core.Get.X_invoke_Arity2(tm_1188, k_1197)) {
												} else {
													panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= k (get tm k))").(string)}, ``)}))
												}
												seq__662_1189, chunk__663_1190, count__664_1191, i__665_1192 = cljs_core.Next.Arity1IQ(seq__662_1195___1), nil, float64(0), float64(0)
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
					var m_1198 = cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1188)
					_ = m_1198
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1199 float64) {
						defer func() {
							if e666 := recover(); e666 != nil {
								if cljs_core.Value_(e666).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e666
										_ = e
										return__1199 = float64(2)
									}
								} else {
									panic(e666)

								}
							}
						}()
						{
							cljs_core.Dissoc_BANG_.X_invoke_Arity2(tm_1188, float64(1))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try (dissoc! tm 1) 1 (catch js/Error e 2)))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1200 float64) {
						defer func() {
							if e667 := recover(); e667 != nil {
								if cljs_core.Value_(e667).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e667
										_ = e
										return__1200 = float64(2)
									}
								} else {
									panic(e667)

								}
							}
						}()
						{
							cljs_core.Assoc_BANG_.X_invoke_Arity3(tm_1188, float64(10), float64(10))
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try (assoc! tm 10 10) 1 (catch js/Error e 2)))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1201 float64) {
						defer func() {
							if e668 := recover(); e668 != nil {
								if cljs_core.Value_(e668).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e668
										_ = e
										return__1201 = float64(2)
									}
								} else {
									panic(e668)

								}
							}
						}()
						{
							cljs_core.Persistent_BANG_.X_invoke_Arity1(tm_1188)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try (persistent! tm) 1 (catch js/Error e 2)))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(float64(2), func() (return__1202 float64) {
						defer func() {
							if e669 := recover(); e669 != nil {
								if cljs_core.Value_(e669).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
									{
										var e = e669
										_ = e
										return__1202 = float64(2)
									}
								} else {
									panic(e669)

								}
							}
						}()
						{
							cljs_core.Count.X_invoke_Arity1(tm_1188)
							return float64(1)
						}
					}()) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (try (count tm) 1 (catch js/Error e 2)))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(m_1198, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(7), []interface{}{float64(0), float64(0), float64(1), float64(1), float64(2), float64(2), float64(4), float64(4), float64(6), float64(6), float64(8), float64(8), float64(9), float64(9)}, nil})) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m {0 0, 1 1, 2 2, 4 4, 6 6, 8 8, 9 9})").(string)}, ``)}))
					}
				}
			}
			{
				var m_1203 = cljs_core.Apply.X_invoke_Arity3(cljs_core.Assoc, cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))
				_ = m_1203
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1203).(float64), (float64(2) * Array_map_conversion_threshold)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) (* 2 array-map-conversion-threshold))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__670 = Array_map_conversion_threshold
					_ = G__670
					return m_1203.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__670)
				}(), Array_map_conversion_threshold) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (m array-map-conversion-threshold) array-map-conversion-threshold)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(m_1203, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity2(func(G__1204 *cljs_core.AFn, m_1203 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1204, 1, func(p1__69_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__69_SHARP_, p1__69_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1203), cljs_core.Range_.X_invoke_Arity1((float64(2)*Array_map_conversion_threshold)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m (into (.-EMPTY cljs.core/PersistentHashMap) (map (fn* [p1__69#] (vector p1__69# p1__69#)) (range (* 2 array-map-conversion-threshold)))))").(string)}, ``)}))
				}
			}
			{
				var m1_1205 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var m2_1206 interface{} = cljs_core.CljsCorePersistentArrayMap_EMPTY
				var i_1207 = float64(0)
				_, _, _ = m1_1205, m2_1206, i_1207
				for {
					if i_1207 < float64(100) {
						m1_1205, m2_1206, i_1207 = cljs_core.Assoc.X_invoke_Arity3(m1_1205, i_1207, i_1207), cljs_core.Assoc.X_invoke_Arity3(m2_1206, strings.Join([]string{cljs_core.Str.X_invoke_Arity1("foo").(string), cljs_core.Str.X_invoke_Arity1(i_1207).(string)}, ``), i_1207), (i_1207 + float64(1))
						continue
					} else {
						if cljs_core.X_EQ_.Arity2IIB(m1_1205, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m1 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (range 100) (range 100))))").(string)}, ``)}))
						}
						if cljs_core.X_EQ_.Arity2IIB(m2_1206, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Map_.X_invoke_Arity3(cljs_core.Vector, cljs_core.Map_.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.Str, "foo").(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= m2 (into (.-EMPTY cljs.core/PersistentHashMap) (map vector (map (partial str \"foo\") (range 100)) (range 100))))").(string)}, ``)}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1205).(float64), float64(100)) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m1) 100)").(string)}, ``)}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1206).(float64), float64(100)) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m2) 100)").(string)}, ``)}))
						}
					}
					break
				}
			}
			{
				var i_1208 = float64(0)
				var m_1209 interface{} = cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(-1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))
				var result_1210 interface{} = cljs_core.CljsCorePersistentVector_EMPTY
				_, _, _ = i_1208, m_1209, result_1210
				for {
					if i_1208 <= (cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) {
						i_1208, m_1209, result_1210 = (i_1208 + float64(1)), cljs_core.Assoc.X_invoke_Arity3(m_1209, i_1208, i_1208), cljs_core.Conj.X_invoke_Arity2(result_1210, cljs_core.Meta.X_invoke_Arity1(m_1209))
						continue
					} else {
						{
							var n_1211 = ((cljs_core.CljsCorePersistentArrayMap_HASHMAP_THRESHOLD + float64(2)) + float64(1))
							var expected_1212 = cljs_core.Repeat.X_invoke_Arity2(n_1211, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})).(*cljs_core.CljsCoreLazySeq)
							_, _ = n_1211, expected_1212
							if cljs_core.X_EQ_.Arity2IIB(result_1210, expected_1212) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= result expected)").(string)}, ``)}))
							}
						}
					}
					break
				}
			}
			{
				var s_1213 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_EMPTY)
				var i_1214 = float64(0)
				_, _ = s_1213, i_1214
				for {
					if i_1214 < float64(100) {
						s_1213, i_1214 = cljs_core.Conj_BANG_.X_invoke_Arity2(s_1213, i_1214), (i_1214 + float64(1))
						continue
					} else {
						{
							var s_1215___1 interface{} = s_1213
							var i_1216___1 = float64(0)
							_, _ = s_1215___1, i_1216___1
							for {
								if i_1216___1 < float64(100) {
									if cljs_core.Mod.X_invoke_Arity2(i_1216___1, float64(3)).(float64) == float64(0) {
										s_1215___1, i_1216___1 = cljs_core.Disj_BANG_.X_invoke_Arity2(s_1215___1, i_1216___1), (i_1216___1 + float64(1))
										continue
									} else {
										s_1215___1, i_1216___1 = s_1215___1, (i_1216___1 + float64(1))
										continue
									}
								} else {
									{
										var s_1217___2 = cljs_core.Persistent_BANG_.X_invoke_Arity1(s_1215___1)
										_ = s_1217___2
										if cljs_core.X_EQ_.Arity2IIB(s_1217___2, func() interface{} {
											var s___3 interface{} = cljs_core.CljsCorePersistentHashSet_EMPTY
											var xs interface{} = cljs_core.Remove.X_invoke_Arity2(func(G__1218 *cljs_core.AFn, s_1215___1 interface{}, i_1216___1 float64, s_1213 interface{}, i_1214 float64, s___3 interface{}, s_1217___2 interface{}) *cljs_core.AFn {
												return cljs_core.Fn(G__1218, 1, func(p1__70_SHARP_ interface{}) interface{} {
													return (cljs_core.Mod.X_invoke_Arity2(p1__70_SHARP_, float64(3)).(float64) == float64(0))
												})
											}(&cljs_core.AFn{}, s_1215___1, i_1216___1, s_1213, i_1214, s___3, s_1217___2), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
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
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= s (loop [s #{} xs (remove (fn* [p1__70#] (zero? (mod p1__70# 3))) (range 100))] (if-let [x (first xs)] (recur (conj s x) (next xs)) s)))").(string)}, ``)}))
										}
										if cljs_core.X_EQ_.Arity2IIB(s_1217___2, cljs_core.Set.X_invoke_Arity1(cljs_core.Remove.X_invoke_Arity2(func(G__1219 *cljs_core.AFn, s_1215___1 interface{}, i_1216___1 float64, s_1213 interface{}, i_1214 float64, s_1217___2 interface{}) *cljs_core.AFn {
											return cljs_core.Fn(G__1219, 1, func(p1__71_SHARP_ interface{}) interface{} {
												return (cljs_core.Mod.X_invoke_Arity2(p1__71_SHARP_, float64(3)).(float64) == float64(0))
											})
										}(&cljs_core.AFn{}, s_1215___1, i_1216___1, s_1213, i_1214, s_1217___2), cljs_core.Range_.X_invoke_Arity1(float64(100)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))) {
										} else {
											panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= s (set (remove (fn* [p1__71#] (zero? (mod p1__71# 3))) (range 100))))").(string)}, ``)}))
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
				var m1_1220 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_1221 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var m2_1222 = cljs_core.Sorted_map_by.X_invoke_ArityVariadic(c2_1221, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{})).(*cljs_core.CljsCorePersistentTreeMap)
				_, _, _ = m1_1220, c2_1221, m2_1222
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeMap)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(m1_1220)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? cljs.core/PersistentTreeMap (type m1))").(string)}, ``)}))
				}
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeMap)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(m2_1222)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? cljs.core/PersistentTreeMap (type m2))").(string)}, ``)}))
				}
				if reflect.DeepEqual(cljs_core.Compare, cljs_core.Native_get_instance_field.X_invoke_Arity2(m1_1220, "Comp")) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? compare (.-comp m1))").(string)}, ``)}))
				}
				if cljs_core.Count.X_invoke_Arity1(m1_1220).(float64) == float64(0) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(zero? (count m1))").(string)}, ``)}))
				}
				if cljs_core.Count.X_invoke_Arity1(m2_1222).(float64) == float64(0) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(zero? (count m2))").(string)}, ``)}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(m1_1220)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (rseq m1))").(string)}, ``)}))
				}
				{
					var m1_1223___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m1_1220, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					var m2_1224___1 = cljs_core.Assoc.X_invoke_ArityVariadic(m2_1222, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
					_, _ = m1_1223___1, m2_1224___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m1_1223___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m1) 3)").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m2_1224___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m2) 3)").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_1223___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq m1) (list [:bar 2] [:foo 1] [:quux 3]))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_1224___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq m2) (list [:quux 3] [:foo 1] [:bar 2]))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m1_1223___1), cljs_core.Rseq.Arity1IQ(m2_1224___1)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq m1) (rseq m2))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m2_1224___1), cljs_core.Rseq.Arity1IQ(m1_1223___1)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq m2) (rseq m1))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m1_1223___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (conj m1 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m1_1223___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count (conj m1 [:wibble 4])) 4)").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(m2_1224___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil})) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (conj m2 [:wibble 4]) {:foo 1, :bar 2, :quux 3, :wibble 4})").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Conj.X_invoke_Arity2(m2_1224___1, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wibble", Fqn: "wibble", X_hash: float64(33319396)}), float64(4)}, nil}))).(float64), float64(4)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count (conj m2 [:wibble 4])) 4)").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m1_1223___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}))).(cljs_core.CljsCoreICollection).X_conj_Arity2(nil)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map key (assoc m1 nil 4)) (list nil :bar :foo :quux))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Key, cljs_core.Assoc.X_invoke_Arity3(m2_1224___1, nil, float64(4))).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(nil)).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map key (assoc m2 nil 4)) (list :quux :foo :bar nil))").(string)}, ``)}))
					}
				}
			}
			{
				var m_1225 = cljs_core.Apply.X_invoke_Arity2(cljs_core.Sorted_map, cljs_core.Mapcat.X_invoke_ArityVariadic(func(G__1228 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1228, 1, func(p1__72_SHARP_ interface{}) interface{} {
						return cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(p1__72_SHARP_)).(cljs_core.CljsCoreICollection).X_conj_Arity2(p1__72_SHARP_)
					})
				}(&cljs_core.AFn{}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_ArityVariadic(cljs_core.Partial.X_invoke_Arity2(cljs_core.Apply, cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(10)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(20), float64(30)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(20)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(50), float64(60)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(30), float64(40)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(40), float64(50)}, nil})}, nil})}))})))
				var s1_1226 = cljs_core.Map_.X_invoke_Arity2(func(G__1229 *cljs_core.AFn, m_1225 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1229, 1, func(p1__73_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__73_SHARP_, p1__73_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1225), cljs_core.Range_.X_invoke_Arity1(float64(60)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				var s2_1227 = cljs_core.Map_.X_invoke_Arity2(func(G__1230 *cljs_core.AFn, m_1225 interface{}, s1_1226 *cljs_core.CljsCoreLazySeq) *cljs_core.AFn {
					return cljs_core.Fn(G__1230, 1, func(p1__74_SHARP_ interface{}) interface{} {
						return (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{p1__74_SHARP_, p1__74_SHARP_}, nil})
					})
				}(&cljs_core.AFn{}, m_1225, s1_1226), cljs_core.Range_.X_invoke_Arity3(float64(59), float64(-1), float64(-1)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)
				_, _, _ = m_1225, s1_1226, s2_1227
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(m_1225).(float64), float64(60)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count m) 60)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m_1225), s1_1226) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq m) s1)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(m_1225), s2_1227) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rseq m) s2)").(string)}, ``)}))
				}
			}
			{
				var m_1231 = cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}))
				_ = m_1231
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Dissoc.X_invoke_Arity2(m_1231, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(2), float64(3)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (dissoc m :foo) (hash-map :bar 2 :quux 3))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(m_1231, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))).(float64), float64(2)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count (dissoc m :foo)) 2)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(m_1231), cljs_core.Hash.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)})}, []interface{}{float64(1), float64(2), float64(3)}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (hash m) (hash (hash-map :foo 1 :bar 2 :quux 3)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1231, cljs_core.X_LT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq m < :foo) (list [:bar 2]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1231, cljs_core.X_LT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq m <= :foo) (list [:bar 2] [:foo 1]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1231, cljs_core.X_GT_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq m > :foo) (list [:quux 3]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(m_1231, cljs_core.X_GT__EQ_, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "quux", Fqn: "quux", X_hash: float64(-2106357800)}), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq m >= :foo) (list [:foo 1] [:quux 3]))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1232 *cljs_core.AFn, m_1231 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1232, 1, func(p1__75_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity2(func(G__1233 *cljs_core.AFn, m_1231 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__1233, 2, func(___ interface{}, x interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_1231), p1__75_SHARP_)
					})
				}(&cljs_core.AFn{}, m_1231), m_1231).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map (fn* [p1__75#] (reduce (fn [_ x] x) p1__75#)) m) (list 2 1 3))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1234 *cljs_core.AFn, m_1231 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1234, 1, func(p1__76_SHARP_ interface{}) interface{} {
						return cljs_core.Reduce.X_invoke_Arity3(func(G__1235 *cljs_core.AFn, m_1231 interface{}) *cljs_core.AFn {
							return cljs_core.Fn(G__1235, 2, func(x interface{}, ___ interface{}) interface{} {
								return x
							})
						}(&cljs_core.AFn{}, m_1231), float64(7), p1__76_SHARP_)
					})
				}(&cljs_core.AFn{}, m_1231), m_1231).(*cljs_core.CljsCoreLazySeq), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map (fn* [p1__76#] (reduce (fn [x _] x) 7 p1__76#)) m) (list 7 7 7))").(string)}, ``)}))
				}
			}
			{
				var s1_1236 = cljs_core.Sorted_set.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c2_1237 = cljs_core.Comp.X_invoke_Arity2(cljs_core.X___, cljs_core.Compare).(cljs_core.CljsCoreIFn)
				var s2_1238 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c2_1237, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var c3_1239 = func(G__1242 *cljs_core.AFn, s1_1236 interface{}, c2_1237 cljs_core.CljsCoreIFn, s2_1238 interface{}) *cljs_core.AFn {
					return cljs_core.Fn(G__1242, 2, func(p1__77_SHARP_ interface{}, p2__78_SHARP_ interface{}) interface{} {
						return cljs_core.Compare.Arity2IIF(cljs_core.Quot.X_invoke_Arity2(p1__77_SHARP_, float64(2)).(float64), cljs_core.Quot.X_invoke_Arity2(p2__78_SHARP_, float64(2)).(float64))
					})
				}(&cljs_core.AFn{}, s1_1236, c2_1237, s2_1238)
				var s3_1240 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(c3_1239, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				var s4_1241 = cljs_core.Sorted_set_by.X_invoke_ArityVariadic(cljs_core.X_LT_, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))
				_, _, _, _, _, _ = s1_1236, c2_1237, s2_1238, c3_1239, s3_1240, s4_1241
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(s1_1236)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? cljs.core/PersistentTreeSet (type s1))").(string)}, ``)}))
				}
				if reflect.DeepEqual(reflect.TypeOf((**cljs_core.CljsCorePersistentTreeSet)(nil)).Elem(), cljs_core.Type_.X_invoke_Arity1(s2_1238)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? cljs.core/PersistentTreeSet (type s2))").(string)}, ``)}))
				}
				if reflect.DeepEqual(cljs_core.Compare, cljs_core.Decorate_(s1_1236).(cljs_core.CljsCoreISorted).X_comparator_Arity1()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? compare (-comparator s1))").(string)}, ``)}))
				}
				if cljs_core.Count.X_invoke_Arity1(s1_1236).(float64) == float64(0) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(zero? (count s1))").(string)}, ``)}))
				}
				if cljs_core.Count.X_invoke_Arity1(s2_1238).(float64) == float64(0) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(zero? (count s2))").(string)}, ``)}))
				}
				if cljs_core.Nil_(cljs_core.Rseq.Arity1IQ(s1_1236)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (rseq s1))").(string)}, ``)}))
				}
				{
					var s1_1243___1 = cljs_core.Conj.X_invoke_ArityVariadic(s1_1236, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s2_1244___1 = cljs_core.Conj.X_invoke_ArityVariadic(s2_1238, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					var s3_1245___1 = cljs_core.Conj.X_invoke_ArityVariadic(s3_1240, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(7), float64(8), float64(9)}))
					var s4_1246___1 = cljs_core.Conj.X_invoke_ArityVariadic(s4_1241, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)}))
					_, _, _, _ = s1_1243___1, s2_1244___1, s3_1245___1, s4_1246___1
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_1243___1), cljs_core.Hash.X_invoke_Arity1(s2_1244___1)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (hash s1) (hash s2))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(s1_1243___1), cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil}))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (hash s1) (hash #{1 3 2}))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_1243___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq s1) (list 1 2 3))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_1243___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rseq s1) (list 3 2 1))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_1244___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq s2) (list 3 2 1))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_1244___1), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rseq s2) (list 1 2 3))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_1243___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count s1) 3)").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_1244___1).(float64), float64(3)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count s2) 3)").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s3_1245___1).(float64), float64(4)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count s3) 4)").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity2(s3_1245___1, float64(0)), float64(1)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get s3 0) 1)").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1245___1, cljs_core.X_GT_, float64(5)), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq s3 > 5) (list 7 8))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1245___1, cljs_core.X_GT_, float64(6)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq s3 > 6) (list 8))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1245___1, cljs_core.X_GT__EQ_, float64(6)), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(8))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(7))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq s3 >= 6) (list 7 8))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1245___1, cljs_core.X_GT__EQ_, float64(12)), nil) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq s3 >= 12) nil)").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1245___1, cljs_core.X_LT_, float64(0)), cljs_core.CljsCoreList_EMPTY) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq s3 < 0) (list))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1245___1, cljs_core.X_LT_, float64(5)), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq s3 < 5) (list 1 2))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1245___1, cljs_core.X_LT_, float64(6)), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq s3 < 6) (list 1 2))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity3(s3_1245___1, cljs_core.X_LT__EQ_, float64(6)), cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq s3 <= 6) (list 1 2 7))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s3_1245___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT__EQ_, float64(6)), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(7))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq s3 >= 2 <= 6) (list 2 7))").(string)}, ``)}))
					}
					if cljs_core.X_EQ_.Arity2IIB(cljs_core.Subseq.X_invoke_Arity5(s4_1246___1, cljs_core.X_GT__EQ_, float64(2), cljs_core.X_LT_, float64(3)), cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(2))) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (subseq s4 >= 2 < 3) (list 2))").(string)}, ``)}))
					}
					{
						var s1_1247___2 = cljs_core.Disj.X_invoke_Arity2(s1_1243___1, float64(2))
						var s2_1248___2 = cljs_core.Disj.X_invoke_Arity2(s2_1244___1, float64(2))
						_, _ = s1_1247___2, s2_1248___2
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s1_1247___2), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq s1) (list 1 3))").(string)}, ``)}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s1_1247___2), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rseq s1) (list 3 1))").(string)}, ``)}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(s2_1248___2), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(3))) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq s2) (list 3 1))").(string)}, ``)}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Rseq.Arity1IQ(s2_1248___2), cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1))) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (rseq s2) (list 1 3))").(string)}, ``)}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s1_1247___2).(float64), float64(2)) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count s1) 2)").(string)}, ``)}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(s2_1248___2).(float64), float64(2)) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count s2) 2)").(string)}, ``)}))
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
					return cljs_core.Fn(map__GT_Person, 1, func(G__673 interface{}) interface{} {
						return (&CljsCore_testPerson{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}).X_invoke_Arity1(G__673), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}).X_invoke_Arity1(G__673), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__673, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})})), nil})
					})
				}(&cljs_core.AFn{})

			}
			Fred = (&CljsCore_testPerson{"Fred", "Mertz", nil, nil, nil})

			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}).X_invoke_Arity1(Fred), "Fred") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (:firstname fred) \"Fred\")").(string)}, ``)}))
			}
			Fred_too = (&CljsCore_testPerson{"Fred", "Mertz", nil, nil, nil})

			if cljs_core.X_EQ_.Arity2IIB(Fred, Fred_too) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= fred fred-too)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(Fred, nil) == false {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(false? (= fred nil))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, Fred) == false {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(false? (= nil fred))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(Map__GT_Person.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), "Fred", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Mertz"}, nil})).(*CljsCore_testPerson), Fred) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map->Person {:firstname \"Fred\", :lastname \"Mertz\"}) fred)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(X__GT_Person.X_invoke_Arity2("Fred", "Mertz").(*CljsCore_testPerson), Fred) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (->Person \"Fred\" \"Mertz\") fred)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(Fred).(float64), float64(2)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count fred) 2)").(string)}, ``)}))
			}
			{
				X__GT_A = func(__GT_A *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_A, 0, func() interface{} {
						return (&CljsCore_testA{nil, nil, nil})
					})
				}(&cljs_core.AFn{})

				Map__GT_A = func(map__GT_A *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(map__GT_A, 1, func(G__692 interface{}) interface{} {
						return (&CljsCore_testA{nil, cljs_core.Dissoc.X_invoke_Arity1(G__692), nil})
					})
				}(&cljs_core.AFn{})

			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})}, nil}), cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&CljsCore_testA{nil, nil, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})}, nil})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo (quote bar)} (meta (with-meta (A.) {:foo (quote bar)})))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}).X_invoke_Arity1(cljs_core.Assoc.X_invoke_Arity3((&CljsCore_testA{nil, nil, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote bar) (:foo (assoc (A.) :foo (quote bar))))").(string)}, ``)}))
			}
			{
				X__GT_C = func(__GT_C *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_C, 3, func(a interface{}, b interface{}, c interface{}) interface{} {
						return (&CljsCore_testC{a, b, c, nil, nil, nil})
					})
				}(&cljs_core.AFn{})

				Map__GT_C = func(map__GT_C *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(map__GT_C, 1, func(G__703 interface{}) interface{} {
						return (&CljsCore_testC{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__703), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__703), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}).X_invoke_Arity1(G__703), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__703, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})})), nil})
					})
				}(&cljs_core.AFn{})

			}
			Letters = (&CljsCore_testC{"a", "b", "c", nil, nil, nil})

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(Letters)), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (set (keys letters)) #{:c :b :a})").(string)}, ``)}))
			}
			More_letters = cljs_core.Assoc.X_invoke_ArityVariadic(Letters, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), "d", cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), "e", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), "f"}))

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(More_letters)), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(6), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (set (keys more-letters)) #{:e :c :b :d :f :a})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(More_letters, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})))), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(5), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (set (keys (dissoc more-letters :d))) #{:e :c :b :f :a})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(More_letters, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)})})))), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (set (keys (dissoc more-letters :d :e))) #{:c :b :f :a})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Keys.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_ArityVariadic(More_letters, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "f", Fqn: "f", X_hash: float64(-1597136552)})})))), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (set (keys (dissoc more-letters :d :e :f))) #{:c :b :a})").(string)}, ``)}))
			}
			{
				var s_1249 = "abc"
				_ = s_1249
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_1249).Length) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (.-length s))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(s_1249).Length) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (. s -length))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_(strings.Join([]string{cljs_core.Str.X_invoke_Arity1(float64(138)).(string)}, ``)).Length) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (. (str 138) -length))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(3), js.JSString_("abc").Length) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 3 (. \"abc\" -length))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1249).Substring(float64(1))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"bc\" (.substring s 1))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_("abc").Substring(float64(1))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"bc\" (.substring \"abc\" 1))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", func(G__1250 *cljs_core.AFn, s_1249 string) *cljs_core.AFn {
					return cljs_core.Fn(G__1250, 2, func(target724 interface{}, start interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target724, "Substring", []interface{}{start})
					})
				}(&cljs_core.AFn{}, s_1249).X_invoke_Arity2(s_1249, float64(1))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"bc\" ((memfn substring start) s 1))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1249).Substring(float64(1))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"bc\" (. s substring 1))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1249).Substring(float64(1))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"bc\" (. s (substring 1)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1249).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"bc\" (. s (substring 1 3)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("bc", js.JSString_(s_1249).Substring(float64(1), float64(3))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"bc\" (.substring s 1 3))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_(s_1249).ToUpperCase()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"ABC\" (. s (toUpperCase)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", js.JSString_("abc").ToUpperCase()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"ABC\" (. \"abc\" (toUpperCase)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("ABC", func(G__1251 *cljs_core.AFn, s_1249 string) *cljs_core.AFn {
					return cljs_core.Fn(G__1251, 1, func(target725 interface{}) interface{} {
						return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(target725, "ToUpperCase", []interface{}{})
					})
				}(&cljs_core.AFn{}, s_1249).X_invoke_Arity1(s_1249)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"ABC\" ((memfn toUpperCase) s))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB("BC", cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_1249).ToUpperCase(), "Substring", []interface{}{float64(1)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"BC\" (. (. s (toUpperCase)) substring 1))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Native_invoke_instance_method.X_invoke_Arity3(js.JSString_(s_1249).ToUpperCase(), "Substring", []interface{}{float64(1)}), "Length")) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (.-length (. (. s (toUpperCase)) substring 1)))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(Fred, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wife", Fqn: "wife", X_hash: float64(2005722828)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ethel", Fqn: "ethel", X_hash: float64(-1460626342)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "friend", Fqn: "friend", X_hash: float64(-286879240)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ricky", Fqn: "ricky", X_hash: float64(-48928873)})}, nil})), Map__GT_Person.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(4), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), "Fred", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Mertz", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wife", Fqn: "wife", X_hash: float64(2005722828)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ethel", Fqn: "ethel", X_hash: float64(-1460626342)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "friend", Fqn: "friend", X_hash: float64(-286879240)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ricky", Fqn: "ricky", X_hash: float64(-48928873)})}, nil})).(*CljsCore_testPerson)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (conj fred {:wife :ethel, :friend :ricky}) (map->Person {:firstname \"Fred\", :lastname \"Mertz\", :wife :ethel, :friend :ricky}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(Fred, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Flintstone"}, nil})), Map__GT_Person.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), "Fred", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Flintstone"}, nil})).(*CljsCore_testPerson)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (conj fred {:lastname \"Flintstone\"}) (map->Person {:firstname \"Fred\", :lastname \"Flintstone\"}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Assoc.X_invoke_Arity3(Fred, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Flintstone"), Map__GT_Person.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), "Fred", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Flintstone"}, nil})).(*CljsCore_testPerson)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (assoc fred :lastname \"Flintstone\") (map->Person {:firstname \"Fred\", :lastname \"Flintstone\"}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Assoc.X_invoke_Arity3(Fred, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wife", Fqn: "wife", X_hash: float64(2005722828)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ethel", Fqn: "ethel", X_hash: float64(-1460626342)})), Map__GT_Person.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), "Fred", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), "Mertz", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "wife", Fqn: "wife", X_hash: float64(2005722828)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ethel", Fqn: "ethel", X_hash: float64(-1460626342)})}, nil})).(*CljsCore_testPerson)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (assoc fred :wife :ethel) (map->Person {:firstname \"Fred\", :lastname \"Mertz\", :wife :ethel}))").(string)}, ``)}))
			}
			{
				X__GT_A2 = func(__GT_A2 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_A2, 1, func(x interface{}) interface{} {
						return (&CljsCore_testA2{x, nil, nil, nil})
					})
				}(&cljs_core.AFn{})

				Map__GT_A2 = func(map__GT_A2 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(map__GT_A2, 1, func(G__728 interface{}) interface{} {
						return (&CljsCore_testA2{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__728), nil, cljs_core.Dissoc.X_invoke_Arity2(G__728, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
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
					return cljs_core.Fn(map__GT_B, 1, func(G__743 interface{}) interface{} {
						return (&CljsCore_testB{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}).X_invoke_Arity1(G__743), nil, cljs_core.Dissoc.X_invoke_Arity2(G__743, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})), nil})
					})
				}(&cljs_core.AFn{})

			}
			if cljs_core.Not_EQ_.Arity2IIB((&CljsCore_testA2{nil, nil, nil, nil}), (&CljsCore_testB{nil, nil, nil, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not= (A2. nil) (B. nil))").(string)}, ``)}))
			}
			Foo = func(foo *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(foo, 1, func(this interface{}) interface{} {
					return cljs_core.Decorate_(this).(CljsCore_testIFoo).Foo_Arity1()
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(func() *CljsCore_testT756 {
				X__GT_t756 = func(__GT_t756 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_t756, 2, func(test_stuff___1 interface{}, meta757 interface{}) interface{} {
						return (&CljsCore_testT756{test_stuff___1, meta757})
					})
				}(&cljs_core.AFn{})

				return (&CljsCore_testT756{test_stuff, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(5), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-column", Fqn: "end-column", X_hash: float64(1425389514)}), float64(61), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-line", Fqn: "end-line", X_hash: float64(1837326455)}), float64(1524), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "column", Fqn: "column", X_hash: float64(2078222095)}), float64(31), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "line", Fqn: "line", X_hash: float64(212345235)}), float64(1524), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "file", Fqn: "file", X_hash: float64(-1269645878)}), "/home/hraberg/go/src/github.com/hraberg/cljs2go/test/cljs/core_test.cljs"}, nil})})
			}(), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (meta (with-meta (reify IFoo (foo [this] :foo)) {:foo :bar})) {:foo :bar})").(string)}, ``)}))
			}
			Foo2 = func() *cljs_core.CljsCoreMultiFn {
				var method_table__1080__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var prefer_table__1081__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var method_cache__1082__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var cached_hierarchy__1083__auto__ = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY).(*cljs_core.CljsCoreAtom)
				var hierarchy__1084__auto__ = cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_EMPTY, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hierarchy", Fqn: "hierarchy", X_hash: float64(-1053470341)}), cljs_core.Get_global_hierarchy.X_invoke_Arity0())
				_, _, _, _, _ = method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__, hierarchy__1084__auto__
				return (&cljs_core.CljsCoreMultiFn{cljs_core.Symbol.X_invoke_Arity2("cljs.core-test", "foo2").(*cljs_core.CljsCoreSymbol), cljs_core.Identity, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)}), hierarchy__1084__auto__, method_table__1080__auto__, prefer_table__1081__auto__, method_cache__1082__auto__, cached_hierarchy__1083__auto__})
			}()

			Foo2.X_add_method_Arity3(float64(0), func(G__1252 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1252, 1, func(x interface{}) interface{} {
					return x
				})
			}(&cljs_core.AFn{}))
			if cljs_core.X_EQ_.Arity2IIB(Foo2, cljs_core.Ffirst.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{Foo2, float64(1)}, true, false).(*cljs_core.CljsCorePersistentArrayMap))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= foo2 (ffirst {foo2 1}))").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :a ((FnLike.)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&CljsCore_testFnLike{}).X_invoke_Arity1(float64(1))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :b ((FnLike.) 1))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&CljsCore_testFnLike{}).X_invoke_Arity2(float64(1), float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :c ((FnLike.) 1 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), cljs_core.Map_.X_invoke_Arity2((&CljsCore_testFnLike{}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0), float64(0)}, nil})).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [:b :b :b] (map (FnLike.) [0 0 0]))").(string)}, ``)}))
			}
			X__GT_FnLikeB = func(__GT_FnLikeB *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_FnLikeB, 1, func(a interface{}) interface{} {
					return (&CljsCore_testFnLikeB{a})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(float64(1), (&CljsCore_testFnLikeB{float64(1)}).X_invoke_Arity0()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 ((FnLikeB. 1)))").(string)}, ``)}))
			}
			{
				var g_1253 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{cljs_core.Conj.X_invoke_Arity2((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}))}, true).(*cljs_core.CljsCorePersistentHashSet)
				var h_1254 = cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "alt", Fqn: "alt", X_hash: float64(-3214426)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "2", Fqn: "2", X_hash: float64(-1645882217)}), nil}, nil}, nil})}, true).(*cljs_core.CljsCorePersistentHashSet)
				_, _ = g_1253, h_1254
				if cljs_core.X_EQ_.Arity2IIB(g_1253, h_1254) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= g h)").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})), cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (hash {:a 1, :b 2}) (hash {:b 2, :a 1}))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, []interface{}{float64(1), float64(2)})), cljs_core.Hash.X_invoke_Arity1(cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, []interface{}{float64(2), float64(1)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (hash (hash-map :a 1 :b 2)) (hash (hash-map :b 2 :a 1)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "start", Fqn: "start", X_hash: float64(-355208981)}), float64(133), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end", Fqn: "end", X_hash: float64(-268185958)}), float64(134)}, nil})), cljs_core.Hash.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "start", Fqn: "start", X_hash: float64(-355208981)}), float64(133), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end", Fqn: "end", X_hash: float64(-268185958)}), float64(134)}, nil})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (hash {:start 133, :end 134}) (hash (apply hash-map [:start 133 :end 134])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})), cljs_core.Hash.X_invoke_Arity1(cljs_core.Keyword.X_invoke_Arity1("a"))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (hash :a) (hash (keyword \"a\")))").(string)}, ``)}))
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
				var fv_1255 = (&CljsCore_testFirst{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})})
				var fs_1256 = (&CljsCore_testFirst{"asdf"})
				_, _ = fv_1255, fs_1256
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fv_1255.X_invoke_Arity0()
				}(), float64(1)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (fv) 1)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					return fs_1256.X_invoke_Arity0()
				}(), "a") {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (fs) \\a)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(strings.Join([]string{cljs_core.Str.X_invoke_Arity1(fs_1256).(string)}, ``), "a") {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (str fs) \\a)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_1255.X_get_first_Arity1(), float64(1)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-get-first fv) 1)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fs_1256.X_get_first_Arity1(), "a") {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-get-first fs) \\a)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(fv_1255.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), float64(1)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-find-first fv [1]) 1)").(string)}, ``)}))
				}
				if reflect.DeepEqual(func() interface{} {
					var G__771 = float64(1)
					_ = G__771
					return fv_1255.X_invoke_Arity1(G__771)
				}(), fv_1255) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(identical? (fv 1) fv)").(string)}, ``)}))
				}
			}
			X__GT_DestructuringWithLocals = func(__GT_DestructuringWithLocals *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_DestructuringWithLocals, 1, func(a interface{}) interface{} {
					return (&CljsCore_testDestructuringWithLocals{a})
				})
			}(&cljs_core.AFn{})

			{
				var t_1257 = (&CljsCore_testDestructuringWithLocals{float64(1)})
				_ = t_1257
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3), float64(1)}, nil}), t_1257.X_find_first_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [2 3 1] (-find-first t [2 3]))").(string)}, ``)}))
				}
			}
			{
				var x_1258 = float64(1)
				_ = x_1258
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__775 = x_1258
					_ = G__775
					switch G__775 {
					case float64(1):
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})

					default:
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("No matching clause: ").(string), cljs_core.Str.X_invoke_Arity1(x_1258).(string)}, ``)}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "one", Fqn: "one", X_hash: float64(935007904)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (case x 1 :one) :one)").(string)}, ``)}))
				}
			}
			{
				var x_1260 = float64(1)
				_ = x_1260
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__776 = x_1260
					_ = G__776
					switch G__776 {
					case float64(2):
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "two", Fqn: "two", X_hash: float64(627606869)})

					default:
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)})

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "default", Fqn: "default", X_hash: float64(-1987822328)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (case x 2 :two :default) :default)").(string)}, ``)}))
				}
			}
			{
				var x_1262 = float64(1)
				_ = x_1262
				if cljs_core.X_EQ_.Arity2IIB(func() (return__1263 interface{}) {
					defer func() {
						if e777 := recover(); e777 != nil {
							if cljs_core.Value_(e777).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
								{
									var e = e777
									_ = e
									return__1263 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
								}
							} else {
								panic(e777)

							}
						}
					}()
					{
						{
							var G__778 = x_1262
							_ = G__778
							switch G__778 {
							case float64(3):
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "three", Fqn: "three", X_hash: float64(-1651831795)})

							default:
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("No matching clause: ").(string), cljs_core.Str.X_invoke_Arity1(x_1262).(string)}, ``)}))

							}
						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (try (case x 3 :three) (catch js/Error e :fail)) :fail)").(string)}, ``)}))
				}
			}
			{
				var x_1265 = float64(1)
				_ = x_1265
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__779 = x_1265
					_ = G__779
					switch G__779 {
					case float64(1), float64(2), float64(3):
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})

					default:
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (case x (1 2 3) :ok :fail) :ok)").(string)}, ``)}))
				}
			}
			{
				var x_1267 = (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil})
				_ = x_1267
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__780 = x_1267
					_ = G__780
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}, nil}), G__780) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("No matching clause: ").(string), cljs_core.Str.X_invoke_Arity1(x_1267).(string)}, ``)}))

					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (case x [:a :b] :ok) :ok)").(string)}, ``)}))
				}
			}
			{
				var a_1268 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = a_1268
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__781 = a_1268
					_ = G__781
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__781) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__781) {
							return nil
						} else {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})

						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (case a nil nil & :amp :none) :none)").(string)}, ``)}))
				}
			}
			{
				var a_1269 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil})
				_ = a_1269
				if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
					var G__782 = a_1269
					_ = G__782
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "&", Str: "&", X_hash: float64(-2144855648), X_meta: nil}), G__782) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB(nil, G__782) {
							return nil
						} else {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})

						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "amp", Fqn: "amp", X_hash: float64(271690571)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (case a nil nil & :amp :none) :amp)").(string)}, ``)}))
				}
			}
			{
				var foo_1270 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil})
				_ = foo_1270
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__783 = foo_1270
					_ = G__783
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__783) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__783) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), G__783) {
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
							} else {
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})

							}
						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (case foo (a b c) :sym :none) :sym)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(func() *cljs_core.CljsCoreKeyword {
					var G__784 = foo_1270
					_ = G__784
					if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "d", Str: "d", X_hash: float64(-682293345), X_meta: nil}), G__784) {
						return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
					} else {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), G__784) {
							return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
						} else {
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}), G__784) {
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "sym", Fqn: "sym", X_hash: float64(-1444860305)})
							} else {
								return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})

							}
						}
					}
				}(), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "none", Fqn: "none", X_hash: float64(1333468478)})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (case foo (b c d) :sym :none) :none)").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF(false, false)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0 (compare false false))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF(false, true)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare false true))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF(true, false)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (compare true false))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF(float64(0), float64(1))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare 0 1))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF(float64(-1), float64(1))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare -1 1))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF(float64(1), float64(1))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0 (compare 1 1))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF(float64(1), float64(0))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (compare 1 0))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF(float64(1), float64(-1))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (compare 1 -1))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF("cljs", "cljs")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0 (compare \"cljs\" \"cljs\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "cljs", Fqn: "cljs", X_hash: float64(1492417629)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "cljs", Fqn: "cljs", X_hash: float64(1492417629)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0 (compare :cljs :cljs))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "cljs", Str: "cljs", X_hash: float64(-1162018140), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "cljs", Str: "cljs", X_hash: float64(-1162018140), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0 (compare (quote cljs) (quote cljs)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF("a", "b")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare \"a\" \"b\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare :a :b))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "b", Str: "b", X_hash: float64(-1172211299), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare (quote a) (quote b)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: "b", Name: "a", Fqn: "b/a", X_hash: float64(-2123407436)}), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "a", Fqn: "c/a", X_hash: float64(-2123407439)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare :b/a :c/a))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare :c :a/b))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (compare :a/b :c))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: "b", Name: "a", Str: "b/a", X_hash: float64(-482875909), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "a", Str: "c/a", X_hash: float64(-482875912), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare (quote b/a) (quote c/a)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare (quote c) (quote a/b)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (compare (quote a/b) (quote c)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF("a", "c")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare \"a\" \"c\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare :a :c))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a", Str: "a", X_hash: float64(-482876059), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "c", Str: "c", X_hash: float64(-122660552), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare (quote a) (quote c)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(1)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare [1 2] [1 1 1]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(1)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare [1 2] [1 2 1]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare [1 1] [1 2]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0 (compare [1 2] [1 2]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (compare [1 2] [1 1]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (compare [1 1 1] [1 2]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(1)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (compare [1 1 2] [1 1 1]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), cljs_core.Compare.Arity2IIF(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(4)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 (compare (subvec [1 2 3] 1) (subvec [1 2 4] 1)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(0), cljs_core.Compare.Arity2IIF(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 0 (compare (subvec [1 2 3] 1) (subvec [1 2 3] 1)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(1), cljs_core.Compare.Arity2IIF(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(4)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 (compare (subvec [1 2 4] 1) (subvec [1 2 3] 1)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(2), float64(1)})).(*cljs_core.CljsCoreList), cljs_core.Reverse.X_invoke_Arity1(cljs_core.Seq.Arity1IQ([]interface{}{float64(1), float64(2), float64(3)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (3 2 1)) (reverse (seq (array 1 2 3))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(3), float64(2), float64(1)})).(*cljs_core.CljsCoreList), cljs_core.Reverse.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (3 2 1)) (reverse [1 2 3]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(4), float64(3), float64(2), float64(1)})).(*cljs_core.CljsCoreList), cljs_core.Cons.X_invoke_Arity2(float64(4), cljs_core.Reverse.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))).(*cljs_core.CljsCoreCons)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (4 3 2 1)) (cons 4 (reverse [1 2 3])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Reverse.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 6 (reduce + (reverse [1 2 3])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(4), float64(3), float64(2)})).(*cljs_core.CljsCoreList), cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, cljs_core.Reverse.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (4 3 2)) (map inc (reverse [1 2 3])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(4), float64(2)})).(*cljs_core.CljsCoreList), cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, cljs_core.Reverse.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil}))).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (4 2)) (filter even? (reverse [1 2 3 4])))").(string)}, ``)}))
			}
			{
				var r_1271 = cljs_core.Range_.X_invoke_Arity1(float64(64)).(*cljs_core.CljsCoreRange)
				var v_1272 = cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, r_1271)
				_, _ = r_1271, v_1272
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1272)), cljs_core.Hash.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1272))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (hash (seq v)) (hash (seq v)))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(6), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Array_chunk.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)}).(*cljs_core.CljsCoreArrayChunk))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 6 (reduce + (array-chunk (array 1 2 3))))").(string)}, ``)}))
				}
				if cljs_core.Value_(cljs_core.Seq.Arity1IQ(v_1272)).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreChunkedSeq)(nil)).Elem()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(instance? ChunkedSeq (seq v))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(r_1271, cljs_core.Seq.Arity1IQ(v_1272)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= r (seq v))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, r_1271).(*cljs_core.CljsCoreLazySeq), cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, v_1272).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map inc r) (map inc v))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, r_1271).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Even_QMARK_, v_1272).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (filter even? r) (filter even? v))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, r_1271).(*cljs_core.CljsCoreLazySeq), cljs_core.Filter.X_invoke_Arity2(cljs_core.Odd_QMARK_, v_1272).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (filter odd? r) (filter odd? v))").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic(r_1271, r_1271, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{r_1271})).(*cljs_core.CljsCoreLazySeq), cljs_core.Concat.X_invoke_ArityVariadic(v_1272, v_1272, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{v_1272})).(*cljs_core.CljsCoreLazySeq)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (concat r r r) (concat v v v))").(string)}, ``)}))
				}
				if cljs_core.DecoratedValue_(cljs_core.Seq.Arity1IQ(v_1272)).Type().Implements(reflect.TypeOf((*cljs_core.CljsCoreIReduce)(nil)).Elem()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(satisfies? IReduce (seq v))").(string)}, ``)}))
				}
				if float64(2010) == cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1272)))))).(float64) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== 2010 (reduce + (nnext (nnext (seq v)))))").(string)}, ``)}))
				}
				if float64(2020) == cljs_core.Reduce.X_invoke_Arity3(cljs_core.X_PLUS_, float64(10), cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq_(cljs_core.Nnext.X_invoke_Arity1(cljs_core.Seq.Arity1IQ(v_1272)))))).(float64) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== 2020 (reduce + 10 (nnext (nnext (seq v)))))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(nil)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (next nil))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(cljs_core.Seq.Arity1IQ([]interface{}{float64(1)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (next (seq (array 1))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.Seq.Arity1IQ([]interface{}{float64(1), float64(2), float64(3)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (2 3)) (next (seq (array 1 2 3))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(cljs_core.Reverse.X_invoke_Arity1(cljs_core.Seq.Arity1IQ([]interface{}{float64(1)})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (next (reverse (seq (array 1)))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(1)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.Reverse.X_invoke_Arity1(cljs_core.Seq.Arity1IQ([]interface{}{float64(1), float64(2), float64(3)})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (2 1)) (next (reverse (seq (array 1 2 3)))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(cljs_core.Cons.X_invoke_Arity2(float64(1), nil).(*cljs_core.CljsCoreCons))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (next (cons 1 nil)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.Cons.X_invoke_Arity2(float64(1), cljs_core.Cons.X_invoke_Arity2(float64(2), cljs_core.Cons.X_invoke_Arity2(float64(3), nil).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreCons)).(*cljs_core.CljsCoreCons))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (2 3)) (next (cons 1 (cons 2 (cons 3 nil)))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__1273 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1273, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), nil).(*cljs_core.CljsCoreCons)
				})
			}(&cljs_core.AFn{}), nil, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (next (lazy-seq (cons 1 nil))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ((&cljs_core.CljsCoreLazySeq{nil, func(G__1274 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1274, 0, func() interface{} {
					return cljs_core.Cons.X_invoke_Arity2(float64(1), (&cljs_core.CljsCoreLazySeq{nil, func(G__1275 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__1275, 0, func() interface{} {
							return cljs_core.Cons.X_invoke_Arity2(float64(2), (&cljs_core.CljsCoreLazySeq{nil, func(G__1276 *cljs_core.AFn) *cljs_core.AFn {
								return cljs_core.Fn(G__1276, 0, func() interface{} {
									return cljs_core.Cons.X_invoke_Arity2(float64(3), nil).(*cljs_core.CljsCoreCons)
								})
							}(&cljs_core.AFn{}), nil, nil})).(*cljs_core.CljsCoreCons)
						})
					}(&cljs_core.AFn{}), nil, nil})).(*cljs_core.CljsCoreCons)
				})
			}(&cljs_core.AFn{}), nil, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (2 3)) (next (lazy-seq (cons 1 (lazy-seq (cons 2 (lazy-seq (cons 3 nil))))))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(1)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (next (list 1)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2(float64(3))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(2))).(cljs_core.CljsCoreICollection).X_conj_Arity2(float64(1)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (2 3)) (next (list 1 2 3)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (next [1]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (2 3)) (next [1 2 3]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(nil, cljs_core.Next.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(2)).(*cljs_core.CljsCoreRange))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= nil (next (range 1 2)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList), cljs_core.Next.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(4)).(*cljs_core.CljsCoreRange))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (2 3)) (next (range 1 4)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreUUID{"550e8400-e29b-41d4-a716-446655440000"}), (&cljs_core.CljsCoreUUID{"550e8400-e29b-41d4-a716-446655440000"})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (UUID. \"550e8400-e29b-41d4-a716-446655440000\") (UUID. \"550e8400-e29b-41d4-a716-446655440000\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCoreUUID{"550e8400-e29b-41d4-a716-446655440000"}), float64(42)}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCoreUUID{"550e8400-e29b-41d4-a716-446655440000"}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-at-all-found", Fqn: "not-at-all-found", X_hash: float64(-1887167218)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 42 (get {(UUID. \"550e8400-e29b-41d4-a716-446655440000\") 42} (UUID. \"550e8400-e29b-41d4-a716-446655440000\") :not-at-all-found))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-at-all-found", Fqn: "not-at-all-found", X_hash: float64(-1887167218)}), cljs_core.Get.X_invoke_Arity3(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{(&cljs_core.CljsCoreUUID{"550e8400-e29b-41d4-a716-446655440000"}), float64(42)}, true, false).(*cljs_core.CljsCorePersistentArrayMap), (&cljs_core.CljsCoreUUID{"666e8400-e29b-41d4-a716-446655440000"}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-at-all-found", Fqn: "not-at-all-found", X_hash: float64(-1887167218)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :not-at-all-found (get {(UUID. \"550e8400-e29b-41d4-a716-446655440000\") 42} (UUID. \"666e8400-e29b-41d4-a716-446655440000\") :not-at-all-found))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})), cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (cljs.core/into (.-EMPTY cljs.core/PersistentQueue) [1]) (into (.-EMPTY cljs.core/PersistentQueue) [1]))").(string)}, ``)}))
			}
			if cljs_core.Not_EQ_.Arity2IIB(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentQueue_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(not= (cljs.core/into (.-EMPTY cljs.core/PersistentQueue) [1 2]) (into (.-EMPTY cljs.core/PersistentQueue) [1]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&js.Date{Millis: 1289585655666}), (&js.Date{Millis: 1289585655666})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #inst \"2010-11-12T18:14:15.666-00:00\" #inst \"2010-11-12T18:14:15.666-00:00\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}), (&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #uuid \"550e8400-e29b-41d4-a716-446655440000\" #uuid \"550e8400-e29b-41d4-a716-446655440000\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(42), cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}), float64(42)}, nil}), (&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 42 (get {#uuid \"550e8400-e29b-41d4-a716-446655440000\" 42} #uuid \"550e8400-e29b-41d4-a716-446655440000\"))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1)})).(string), "1") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str 1) \"1\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(-1)})).(string), "-1") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str -1) \"-1\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{-1.5})).(string), "-1.5") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str -1.5) \"-1.5\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4)}, nil})})).(string), "[3 4]") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str [3 4]) \"[3 4]\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"foo"})).(string), "\"foo\"") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str \"foo\") \"\\\"foo\\\"\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "hello", Fqn: "hello", X_hash: float64(-245025397)})})).(string), ":hello") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str :hello) \":hello\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "goodbye", Str: "goodbye", X_hash: float64(-1229580373), X_meta: nil})})).(string), "goodbye") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (quote goodbye)) \"goodbye\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList)})).(string), "(7 8 9)") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (quote (7 8 9))) \"(7 8 9)\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "deref", Str: "deref", X_hash: float64(1494944732), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})})).(*cljs_core.CljsCoreList)})).(string), "(deref foo)") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (quote (deref foo))) \"(deref foo)\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreSymbol{Ns: nil, Name: "quote", Str: "quote", X_hash: float64(1377916282), X_meta: nil}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "bar", Str: "bar", X_hash: float64(254284943), X_meta: nil})})).(*cljs_core.CljsCoreList)})).(string), "(quote bar)") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (quote (quote bar))) \"(quote bar)\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreSymbol{Ns: "foo", Name: "bar", Str: "foo/bar", X_hash: float64(254379989), X_meta: nil})})).(string), "foo/bar") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (quote foo/bar)) \"foo/bar\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"a"})).(string), "\"a\"") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str \\a) \"\\\"a\\\"\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: "foo", Name: "bar", Fqn: "foo/bar", X_hash: float64(-1386151538)})})).(string), ":foo/bar") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str :foo/bar) \":foo/bar\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{nil})).(string), "nil") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str nil) \"nil\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{true})).(string), "true") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str true) \"true\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{false})).(string), "false") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str false) \"false\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"string"})).(string), "\"string\"") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str \"string\") \"\\\"string\\\"\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"\u00FC\u00F1\u00EE\u00E7\u00F3\u2202\u00A3", (&cljs_core.CljsCoreKeyword{Ns: "\u0E17\u0E14\u0E2A\u0E2D\u0E1A", Name: "\u4F60\u597D", Fqn: "\u0E17\u0E14\u0E2A\u0E2D\u0E1A/\u4F60\u597D", X_hash: float64(404745050)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "\u3053\u3093\u306B\u3061\u306F", Str: "\u3053\u3093\u306B\u3061\u306F", X_hash: float64(1271843387), X_meta: nil})}, nil})})).(string), "[\"\u00FC\u00F1\u00EE\u00E7\u00F3\u2202\u00A3\" :\u0E17\u0E14\u0E2A\u0E2D\u0E1A/\u4F60\u597D \u3053\u3093\u306B\u3061\u306F]") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str [\"\u00FC\u00F1\u00EE\u00E7\u00F3\u2202\u00A3\" :\u0E17\u0E14\u0E2A\u0E2D\u0E1A/\u4F60\u597D (quote \u3053\u3093\u306B\u3061\u306F)]) \"[\\\"\u00FC\u00F1\u00EE\u00E7\u00F3\u2202\u00A3\\\" :\u0E17\u0E14\u0E2A\u0E2D\u0E1A/\u4F60\u597D \u3053\u3093\u306B\u3061\u306F]\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{"escape chars \t \r \n \\ \" \b \f"})).(string), "\"escape chars \\t \\r \\n \\\\ \\\" \\b \\f\"") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str \"escape chars \\t \\r \\n \\\\ \\\" \\b \\f\") \"\\\"escape chars \\\\t \\\\r \\\\n \\\\\\\\ \\\\\\\" \\\\b \\\\f\\\"\")").(string)}, ``)}))
			}
			{
				X__GT_PrintMe = func(__GT_PrintMe *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(__GT_PrintMe, 2, func(a interface{}, b interface{}) interface{} {
						return (&CljsCore_testPrintMe{a, b, nil, nil, nil})
					})
				}(&cljs_core.AFn{})

				Map__GT_PrintMe = func(map__GT_PrintMe *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(map__GT_PrintMe, 1, func(G__787 interface{}) interface{} {
						return (&CljsCore_testPrintMe{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1(G__787), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(G__787), nil, cljs_core.Dissoc.X_invoke_ArityVariadic(G__787, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})})), nil})
					})
				}(&cljs_core.AFn{})

			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&CljsCore_testPrintMe{float64(1), float64(2), nil, nil, nil})})).(string), "#cljs.core-test.PrintMe{:a 1, :b 2}") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (PrintMe. 1 2)) \"#cljs.core-test.PrintMe{:a 1, :b 2}\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{"2010-11-12T13:14:15.666-05:00"})})).(string), "#inst \"2010-11-12T18:14:15.666-00:00\"") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (js/Date. \"2010-11-12T13:14:15.666-05:00\")) \"#inst \\\"2010-11-12T18:14:15.666-00:00\\\"\")").(string)}, ``)}))
			}
			{
				var seq__804_1277 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(13)).(*cljs_core.CljsCoreRange))
				var chunk__817_1278 interface{} = nil
				var count__818_1279 = float64(0)
				var i__819_1280 = float64(0)
				_, _, _, _ = seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280
				for {
					if i__819_1280 < count__818_1279 {
						{
							var month_1281 = cljs_core.Decorate_(chunk__817_1278).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__819_1280)
							_ = month_1281
							{
								var seq__820_1282 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
								var chunk__825_1283 interface{} = nil
								var count__826_1284 = float64(0)
								var i__827_1285 = float64(0)
								_, _, _, _ = seq__820_1282, chunk__825_1283, count__826_1284, i__827_1285
								for {
									if i__827_1285 < count__826_1284 {
										{
											var day_1286 = cljs_core.Decorate_(chunk__825_1283).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__827_1285)
											_ = day_1286
											{
												var seq__828_1287 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
												var chunk__829_1288 interface{} = nil
												var count__830_1289 = float64(0)
												var i__831_1290 = float64(0)
												_, _, _, _ = seq__828_1287, chunk__829_1288, count__830_1289, i__831_1290
												for {
													if i__831_1290 < count__830_1289 {
														{
															var hour_1291 = cljs_core.Decorate_(chunk__829_1288).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__831_1290)
															_ = hour_1291
															{
																var pad_1292 = func(G__1294 *cljs_core.AFn, seq__828_1287 interface{}, chunk__829_1288 interface{}, count__830_1289 float64, i__831_1290 float64, seq__820_1282 interface{}, chunk__825_1283 interface{}, count__826_1284 float64, i__827_1285 float64, seq__804_1277 interface{}, chunk__817_1278 interface{}, count__818_1279 float64, i__819_1280 float64, hour_1291 interface{}, day_1286 interface{}, month_1281 interface{}) *cljs_core.AFn {
																	return cljs_core.Fn(G__1294, 1, func(n interface{}) interface{} {
																		if n.(float64) < float64(10) {
																			return strings.Join([]string{cljs_core.Str.X_invoke_Arity1("0").(string), cljs_core.Str.X_invoke_Arity1(n).(string)}, ``)
																		} else {
																			return n
																		}
																	})
																}(&cljs_core.AFn{}, seq__828_1287, chunk__829_1288, count__830_1289, i__831_1290, seq__820_1282, chunk__825_1283, count__826_1284, i__827_1285, seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280, hour_1291, day_1286, month_1281)
																var inst_1293 = strings.Join([]string{cljs_core.Str.X_invoke_Arity1("2010-").(string), cljs_core.Str.X_invoke_Arity1(pad_1292.X_invoke_Arity1(month_1281)).(string), cljs_core.Str.X_invoke_Arity1("-").(string), cljs_core.Str.X_invoke_Arity1(pad_1292.X_invoke_Arity1(day_1286)).(string), cljs_core.Str.X_invoke_Arity1("T").(string), cljs_core.Str.X_invoke_Arity1(pad_1292.X_invoke_Arity1(hour_1291)).(string), cljs_core.Str.X_invoke_Arity1(":14:15.666-00:00").(string)}, ``)
																_, _ = pad_1292, inst_1293
																if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1293})})).(string), strings.Join([]string{cljs_core.Str.X_invoke_Arity1("#inst \"").(string), cljs_core.Str.X_invoke_Arity1(inst_1293).(string), cljs_core.Str.X_invoke_Arity1("\"").(string)}, ``)) {
																} else {
																	panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))").(string)}, ``)}))
																}
															}
															seq__828_1287, chunk__829_1288, count__830_1289, i__831_1290 = seq__828_1287, chunk__829_1288, count__830_1289, (i__831_1290 + float64(1))
															continue
														}
													} else {
														{
															var temp__4388__auto___1295 = cljs_core.Seq.Arity1IQ(seq__828_1287)
															_ = temp__4388__auto___1295
															if cljs_core.Truth_(temp__4388__auto___1295) {
																{
																	var seq__828_1296___1 = temp__4388__auto___1295
																	_ = seq__828_1296___1
																	if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__828_1296___1) {
																		{
																			var c__970__auto___1297 = cljs_core.Chunk_first.X_invoke_Arity1(seq__828_1296___1)
																			_ = c__970__auto___1297
																			seq__828_1287, chunk__829_1288, count__830_1289, i__831_1290 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__828_1296___1), c__970__auto___1297, cljs_core.Count.X_invoke_Arity1(c__970__auto___1297).(float64), float64(0)
																			continue
																		}
																	} else {
																		{
																			var hour_1298 = cljs_core.First.X_invoke_Arity1(seq__828_1296___1)
																			_ = hour_1298
																			{
																				var pad_1299 = func(G__1301 *cljs_core.AFn, seq__828_1287 interface{}, chunk__829_1288 interface{}, count__830_1289 float64, i__831_1290 float64, seq__820_1282 interface{}, chunk__825_1283 interface{}, count__826_1284 float64, i__827_1285 float64, seq__804_1277 interface{}, chunk__817_1278 interface{}, count__818_1279 float64, i__819_1280 float64, hour_1298 interface{}, seq__828_1296___1 interface{}, temp__4388__auto___1295 cljs_core.CljsCoreISeq, day_1286 interface{}, month_1281 interface{}) *cljs_core.AFn {
																					return cljs_core.Fn(G__1301, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return strings.Join([]string{cljs_core.Str.X_invoke_Arity1("0").(string), cljs_core.Str.X_invoke_Arity1(n).(string)}, ``)
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__828_1287, chunk__829_1288, count__830_1289, i__831_1290, seq__820_1282, chunk__825_1283, count__826_1284, i__827_1285, seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280, hour_1298, seq__828_1296___1, temp__4388__auto___1295, day_1286, month_1281)
																				var inst_1300 = strings.Join([]string{cljs_core.Str.X_invoke_Arity1("2010-").(string), cljs_core.Str.X_invoke_Arity1(pad_1299.X_invoke_Arity1(month_1281)).(string), cljs_core.Str.X_invoke_Arity1("-").(string), cljs_core.Str.X_invoke_Arity1(pad_1299.X_invoke_Arity1(day_1286)).(string), cljs_core.Str.X_invoke_Arity1("T").(string), cljs_core.Str.X_invoke_Arity1(pad_1299.X_invoke_Arity1(hour_1298)).(string), cljs_core.Str.X_invoke_Arity1(":14:15.666-00:00").(string)}, ``)
																				_, _ = pad_1299, inst_1300
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1300})})).(string), strings.Join([]string{cljs_core.Str.X_invoke_Arity1("#inst \"").(string), cljs_core.Str.X_invoke_Arity1(inst_1300).(string), cljs_core.Str.X_invoke_Arity1("\"").(string)}, ``)) {
																				} else {
																					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))").(string)}, ``)}))
																				}
																			}
																			seq__828_1287, chunk__829_1288, count__830_1289, i__831_1290 = cljs_core.Next.Arity1IQ(seq__828_1296___1), nil, float64(0), float64(0)
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
											seq__820_1282, chunk__825_1283, count__826_1284, i__827_1285 = seq__820_1282, chunk__825_1283, count__826_1284, (i__827_1285 + float64(1))
											continue
										}
									} else {
										{
											var temp__4388__auto___1302 = cljs_core.Seq.Arity1IQ(seq__820_1282)
											_ = temp__4388__auto___1302
											if cljs_core.Truth_(temp__4388__auto___1302) {
												{
													var seq__820_1303___1 = temp__4388__auto___1302
													_ = seq__820_1303___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__820_1303___1) {
														{
															var c__970__auto___1304 = cljs_core.Chunk_first.X_invoke_Arity1(seq__820_1303___1)
															_ = c__970__auto___1304
															seq__820_1282, chunk__825_1283, count__826_1284, i__827_1285 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__820_1303___1), c__970__auto___1304, cljs_core.Count.X_invoke_Arity1(c__970__auto___1304).(float64), float64(0)
															continue
														}
													} else {
														{
															var day_1305 = cljs_core.First.X_invoke_Arity1(seq__820_1303___1)
															_ = day_1305
															{
																var seq__821_1306 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																var chunk__822_1307 interface{} = nil
																var count__823_1308 = float64(0)
																var i__824_1309 = float64(0)
																_, _, _, _ = seq__821_1306, chunk__822_1307, count__823_1308, i__824_1309
																for {
																	if i__824_1309 < count__823_1308 {
																		{
																			var hour_1310 = cljs_core.Decorate_(chunk__822_1307).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__824_1309)
																			_ = hour_1310
																			{
																				var pad_1311 = func(G__1313 *cljs_core.AFn, seq__821_1306 interface{}, chunk__822_1307 interface{}, count__823_1308 float64, i__824_1309 float64, seq__820_1282 interface{}, chunk__825_1283 interface{}, count__826_1284 float64, i__827_1285 float64, seq__804_1277 interface{}, chunk__817_1278 interface{}, count__818_1279 float64, i__819_1280 float64, hour_1310 interface{}, day_1305 interface{}, seq__820_1303___1 interface{}, temp__4388__auto___1302 cljs_core.CljsCoreISeq, month_1281 interface{}) *cljs_core.AFn {
																					return cljs_core.Fn(G__1313, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return strings.Join([]string{cljs_core.Str.X_invoke_Arity1("0").(string), cljs_core.Str.X_invoke_Arity1(n).(string)}, ``)
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__821_1306, chunk__822_1307, count__823_1308, i__824_1309, seq__820_1282, chunk__825_1283, count__826_1284, i__827_1285, seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280, hour_1310, day_1305, seq__820_1303___1, temp__4388__auto___1302, month_1281)
																				var inst_1312 = strings.Join([]string{cljs_core.Str.X_invoke_Arity1("2010-").(string), cljs_core.Str.X_invoke_Arity1(pad_1311.X_invoke_Arity1(month_1281)).(string), cljs_core.Str.X_invoke_Arity1("-").(string), cljs_core.Str.X_invoke_Arity1(pad_1311.X_invoke_Arity1(day_1305)).(string), cljs_core.Str.X_invoke_Arity1("T").(string), cljs_core.Str.X_invoke_Arity1(pad_1311.X_invoke_Arity1(hour_1310)).(string), cljs_core.Str.X_invoke_Arity1(":14:15.666-00:00").(string)}, ``)
																				_, _ = pad_1311, inst_1312
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1312})})).(string), strings.Join([]string{cljs_core.Str.X_invoke_Arity1("#inst \"").(string), cljs_core.Str.X_invoke_Arity1(inst_1312).(string), cljs_core.Str.X_invoke_Arity1("\"").(string)}, ``)) {
																				} else {
																					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))").(string)}, ``)}))
																				}
																			}
																			seq__821_1306, chunk__822_1307, count__823_1308, i__824_1309 = seq__821_1306, chunk__822_1307, count__823_1308, (i__824_1309 + float64(1))
																			continue
																		}
																	} else {
																		{
																			var temp__4388__auto___1314___1 = cljs_core.Seq.Arity1IQ(seq__821_1306)
																			_ = temp__4388__auto___1314___1
																			if cljs_core.Truth_(temp__4388__auto___1314___1) {
																				{
																					var seq__821_1315___1 = temp__4388__auto___1314___1
																					_ = seq__821_1315___1
																					if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__821_1315___1) {
																						{
																							var c__970__auto___1316 = cljs_core.Chunk_first.X_invoke_Arity1(seq__821_1315___1)
																							_ = c__970__auto___1316
																							seq__821_1306, chunk__822_1307, count__823_1308, i__824_1309 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__821_1315___1), c__970__auto___1316, cljs_core.Count.X_invoke_Arity1(c__970__auto___1316).(float64), float64(0)
																							continue
																						}
																					} else {
																						{
																							var hour_1317 = cljs_core.First.X_invoke_Arity1(seq__821_1315___1)
																							_ = hour_1317
																							{
																								var pad_1318 = func(G__1320 *cljs_core.AFn, seq__821_1306 interface{}, chunk__822_1307 interface{}, count__823_1308 float64, i__824_1309 float64, seq__820_1282 interface{}, chunk__825_1283 interface{}, count__826_1284 float64, i__827_1285 float64, seq__804_1277 interface{}, chunk__817_1278 interface{}, count__818_1279 float64, i__819_1280 float64, hour_1317 interface{}, seq__821_1315___1 interface{}, temp__4388__auto___1314___1 cljs_core.CljsCoreISeq, day_1305 interface{}, seq__820_1303___1 interface{}, temp__4388__auto___1302 cljs_core.CljsCoreISeq, month_1281 interface{}) *cljs_core.AFn {
																									return cljs_core.Fn(G__1320, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return strings.Join([]string{cljs_core.Str.X_invoke_Arity1("0").(string), cljs_core.Str.X_invoke_Arity1(n).(string)}, ``)
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__821_1306, chunk__822_1307, count__823_1308, i__824_1309, seq__820_1282, chunk__825_1283, count__826_1284, i__827_1285, seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280, hour_1317, seq__821_1315___1, temp__4388__auto___1314___1, day_1305, seq__820_1303___1, temp__4388__auto___1302, month_1281)
																								var inst_1319 = strings.Join([]string{cljs_core.Str.X_invoke_Arity1("2010-").(string), cljs_core.Str.X_invoke_Arity1(pad_1318.X_invoke_Arity1(month_1281)).(string), cljs_core.Str.X_invoke_Arity1("-").(string), cljs_core.Str.X_invoke_Arity1(pad_1318.X_invoke_Arity1(day_1305)).(string), cljs_core.Str.X_invoke_Arity1("T").(string), cljs_core.Str.X_invoke_Arity1(pad_1318.X_invoke_Arity1(hour_1317)).(string), cljs_core.Str.X_invoke_Arity1(":14:15.666-00:00").(string)}, ``)
																								_, _ = pad_1318, inst_1319
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1319})})).(string), strings.Join([]string{cljs_core.Str.X_invoke_Arity1("#inst \"").(string), cljs_core.Str.X_invoke_Arity1(inst_1319).(string), cljs_core.Str.X_invoke_Arity1("\"").(string)}, ``)) {
																								} else {
																									panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))").(string)}, ``)}))
																								}
																							}
																							seq__821_1306, chunk__822_1307, count__823_1308, i__824_1309 = cljs_core.Next.Arity1IQ(seq__821_1315___1), nil, float64(0), float64(0)
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
															seq__820_1282, chunk__825_1283, count__826_1284, i__827_1285 = cljs_core.Next.Arity1IQ(seq__820_1303___1), nil, float64(0), float64(0)
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
							seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280 = seq__804_1277, chunk__817_1278, count__818_1279, (i__819_1280 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1321 = cljs_core.Seq.Arity1IQ(seq__804_1277)
							_ = temp__4388__auto___1321
							if cljs_core.Truth_(temp__4388__auto___1321) {
								{
									var seq__804_1322___1 = temp__4388__auto___1321
									_ = seq__804_1322___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__804_1322___1) {
										{
											var c__970__auto___1323 = cljs_core.Chunk_first.X_invoke_Arity1(seq__804_1322___1)
											_ = c__970__auto___1323
											seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__804_1322___1), c__970__auto___1323, cljs_core.Count.X_invoke_Arity1(c__970__auto___1323).(float64), float64(0)
											continue
										}
									} else {
										{
											var month_1324 = cljs_core.First.X_invoke_Arity1(seq__804_1322___1)
											_ = month_1324
											{
												var seq__805_1325 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(29)).(*cljs_core.CljsCoreRange))
												var chunk__810_1326 interface{} = nil
												var count__811_1327 = float64(0)
												var i__812_1328 = float64(0)
												_, _, _, _ = seq__805_1325, chunk__810_1326, count__811_1327, i__812_1328
												for {
													if i__812_1328 < count__811_1327 {
														{
															var day_1329 = cljs_core.Decorate_(chunk__810_1326).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__812_1328)
															_ = day_1329
															{
																var seq__813_1330 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																var chunk__814_1331 interface{} = nil
																var count__815_1332 = float64(0)
																var i__816_1333 = float64(0)
																_, _, _, _ = seq__813_1330, chunk__814_1331, count__815_1332, i__816_1333
																for {
																	if i__816_1333 < count__815_1332 {
																		{
																			var hour_1334 = cljs_core.Decorate_(chunk__814_1331).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__816_1333)
																			_ = hour_1334
																			{
																				var pad_1335 = func(G__1337 *cljs_core.AFn, seq__813_1330 interface{}, chunk__814_1331 interface{}, count__815_1332 float64, i__816_1333 float64, seq__805_1325 interface{}, chunk__810_1326 interface{}, count__811_1327 float64, i__812_1328 float64, seq__804_1277 interface{}, chunk__817_1278 interface{}, count__818_1279 float64, i__819_1280 float64, hour_1334 interface{}, day_1329 interface{}, month_1324 interface{}, seq__804_1322___1 interface{}, temp__4388__auto___1321 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																					return cljs_core.Fn(G__1337, 1, func(n interface{}) interface{} {
																						if n.(float64) < float64(10) {
																							return strings.Join([]string{cljs_core.Str.X_invoke_Arity1("0").(string), cljs_core.Str.X_invoke_Arity1(n).(string)}, ``)
																						} else {
																							return n
																						}
																					})
																				}(&cljs_core.AFn{}, seq__813_1330, chunk__814_1331, count__815_1332, i__816_1333, seq__805_1325, chunk__810_1326, count__811_1327, i__812_1328, seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280, hour_1334, day_1329, month_1324, seq__804_1322___1, temp__4388__auto___1321)
																				var inst_1336 = strings.Join([]string{cljs_core.Str.X_invoke_Arity1("2010-").(string), cljs_core.Str.X_invoke_Arity1(pad_1335.X_invoke_Arity1(month_1324)).(string), cljs_core.Str.X_invoke_Arity1("-").(string), cljs_core.Str.X_invoke_Arity1(pad_1335.X_invoke_Arity1(day_1329)).(string), cljs_core.Str.X_invoke_Arity1("T").(string), cljs_core.Str.X_invoke_Arity1(pad_1335.X_invoke_Arity1(hour_1334)).(string), cljs_core.Str.X_invoke_Arity1(":14:15.666-00:00").(string)}, ``)
																				_, _ = pad_1335, inst_1336
																				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1336})})).(string), strings.Join([]string{cljs_core.Str.X_invoke_Arity1("#inst \"").(string), cljs_core.Str.X_invoke_Arity1(inst_1336).(string), cljs_core.Str.X_invoke_Arity1("\"").(string)}, ``)) {
																				} else {
																					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))").(string)}, ``)}))
																				}
																			}
																			seq__813_1330, chunk__814_1331, count__815_1332, i__816_1333 = seq__813_1330, chunk__814_1331, count__815_1332, (i__816_1333 + float64(1))
																			continue
																		}
																	} else {
																		{
																			var temp__4388__auto___1338___1 = cljs_core.Seq.Arity1IQ(seq__813_1330)
																			_ = temp__4388__auto___1338___1
																			if cljs_core.Truth_(temp__4388__auto___1338___1) {
																				{
																					var seq__813_1339___1 = temp__4388__auto___1338___1
																					_ = seq__813_1339___1
																					if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__813_1339___1) {
																						{
																							var c__970__auto___1340 = cljs_core.Chunk_first.X_invoke_Arity1(seq__813_1339___1)
																							_ = c__970__auto___1340
																							seq__813_1330, chunk__814_1331, count__815_1332, i__816_1333 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__813_1339___1), c__970__auto___1340, cljs_core.Count.X_invoke_Arity1(c__970__auto___1340).(float64), float64(0)
																							continue
																						}
																					} else {
																						{
																							var hour_1341 = cljs_core.First.X_invoke_Arity1(seq__813_1339___1)
																							_ = hour_1341
																							{
																								var pad_1342 = func(G__1344 *cljs_core.AFn, seq__813_1330 interface{}, chunk__814_1331 interface{}, count__815_1332 float64, i__816_1333 float64, seq__805_1325 interface{}, chunk__810_1326 interface{}, count__811_1327 float64, i__812_1328 float64, seq__804_1277 interface{}, chunk__817_1278 interface{}, count__818_1279 float64, i__819_1280 float64, hour_1341 interface{}, seq__813_1339___1 interface{}, temp__4388__auto___1338___1 cljs_core.CljsCoreISeq, day_1329 interface{}, month_1324 interface{}, seq__804_1322___1 interface{}, temp__4388__auto___1321 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																									return cljs_core.Fn(G__1344, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return strings.Join([]string{cljs_core.Str.X_invoke_Arity1("0").(string), cljs_core.Str.X_invoke_Arity1(n).(string)}, ``)
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__813_1330, chunk__814_1331, count__815_1332, i__816_1333, seq__805_1325, chunk__810_1326, count__811_1327, i__812_1328, seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280, hour_1341, seq__813_1339___1, temp__4388__auto___1338___1, day_1329, month_1324, seq__804_1322___1, temp__4388__auto___1321)
																								var inst_1343 = strings.Join([]string{cljs_core.Str.X_invoke_Arity1("2010-").(string), cljs_core.Str.X_invoke_Arity1(pad_1342.X_invoke_Arity1(month_1324)).(string), cljs_core.Str.X_invoke_Arity1("-").(string), cljs_core.Str.X_invoke_Arity1(pad_1342.X_invoke_Arity1(day_1329)).(string), cljs_core.Str.X_invoke_Arity1("T").(string), cljs_core.Str.X_invoke_Arity1(pad_1342.X_invoke_Arity1(hour_1341)).(string), cljs_core.Str.X_invoke_Arity1(":14:15.666-00:00").(string)}, ``)
																								_, _ = pad_1342, inst_1343
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1343})})).(string), strings.Join([]string{cljs_core.Str.X_invoke_Arity1("#inst \"").(string), cljs_core.Str.X_invoke_Arity1(inst_1343).(string), cljs_core.Str.X_invoke_Arity1("\"").(string)}, ``)) {
																								} else {
																									panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))").(string)}, ``)}))
																								}
																							}
																							seq__813_1330, chunk__814_1331, count__815_1332, i__816_1333 = cljs_core.Next.Arity1IQ(seq__813_1339___1), nil, float64(0), float64(0)
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
															seq__805_1325, chunk__810_1326, count__811_1327, i__812_1328 = seq__805_1325, chunk__810_1326, count__811_1327, (i__812_1328 + float64(1))
															continue
														}
													} else {
														{
															var temp__4388__auto___1345___1 = cljs_core.Seq.Arity1IQ(seq__805_1325)
															_ = temp__4388__auto___1345___1
															if cljs_core.Truth_(temp__4388__auto___1345___1) {
																{
																	var seq__805_1346___1 = temp__4388__auto___1345___1
																	_ = seq__805_1346___1
																	if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__805_1346___1) {
																		{
																			var c__970__auto___1347 = cljs_core.Chunk_first.X_invoke_Arity1(seq__805_1346___1)
																			_ = c__970__auto___1347
																			seq__805_1325, chunk__810_1326, count__811_1327, i__812_1328 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__805_1346___1), c__970__auto___1347, cljs_core.Count.X_invoke_Arity1(c__970__auto___1347).(float64), float64(0)
																			continue
																		}
																	} else {
																		{
																			var day_1348 = cljs_core.First.X_invoke_Arity1(seq__805_1346___1)
																			_ = day_1348
																			{
																				var seq__806_1349 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(23)).(*cljs_core.CljsCoreRange))
																				var chunk__807_1350 interface{} = nil
																				var count__808_1351 = float64(0)
																				var i__809_1352 = float64(0)
																				_, _, _, _ = seq__806_1349, chunk__807_1350, count__808_1351, i__809_1352
																				for {
																					if i__809_1352 < count__808_1351 {
																						{
																							var hour_1353 = cljs_core.Decorate_(chunk__807_1350).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__809_1352)
																							_ = hour_1353
																							{
																								var pad_1354 = func(G__1356 *cljs_core.AFn, seq__806_1349 interface{}, chunk__807_1350 interface{}, count__808_1351 float64, i__809_1352 float64, seq__805_1325 interface{}, chunk__810_1326 interface{}, count__811_1327 float64, i__812_1328 float64, seq__804_1277 interface{}, chunk__817_1278 interface{}, count__818_1279 float64, i__819_1280 float64, hour_1353 interface{}, day_1348 interface{}, seq__805_1346___1 interface{}, temp__4388__auto___1345___1 cljs_core.CljsCoreISeq, month_1324 interface{}, seq__804_1322___1 interface{}, temp__4388__auto___1321 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																									return cljs_core.Fn(G__1356, 1, func(n interface{}) interface{} {
																										if n.(float64) < float64(10) {
																											return strings.Join([]string{cljs_core.Str.X_invoke_Arity1("0").(string), cljs_core.Str.X_invoke_Arity1(n).(string)}, ``)
																										} else {
																											return n
																										}
																									})
																								}(&cljs_core.AFn{}, seq__806_1349, chunk__807_1350, count__808_1351, i__809_1352, seq__805_1325, chunk__810_1326, count__811_1327, i__812_1328, seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280, hour_1353, day_1348, seq__805_1346___1, temp__4388__auto___1345___1, month_1324, seq__804_1322___1, temp__4388__auto___1321)
																								var inst_1355 = strings.Join([]string{cljs_core.Str.X_invoke_Arity1("2010-").(string), cljs_core.Str.X_invoke_Arity1(pad_1354.X_invoke_Arity1(month_1324)).(string), cljs_core.Str.X_invoke_Arity1("-").(string), cljs_core.Str.X_invoke_Arity1(pad_1354.X_invoke_Arity1(day_1348)).(string), cljs_core.Str.X_invoke_Arity1("T").(string), cljs_core.Str.X_invoke_Arity1(pad_1354.X_invoke_Arity1(hour_1353)).(string), cljs_core.Str.X_invoke_Arity1(":14:15.666-00:00").(string)}, ``)
																								_, _ = pad_1354, inst_1355
																								if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1355})})).(string), strings.Join([]string{cljs_core.Str.X_invoke_Arity1("#inst \"").(string), cljs_core.Str.X_invoke_Arity1(inst_1355).(string), cljs_core.Str.X_invoke_Arity1("\"").(string)}, ``)) {
																								} else {
																									panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))").(string)}, ``)}))
																								}
																							}
																							seq__806_1349, chunk__807_1350, count__808_1351, i__809_1352 = seq__806_1349, chunk__807_1350, count__808_1351, (i__809_1352 + float64(1))
																							continue
																						}
																					} else {
																						{
																							var temp__4388__auto___1357___2 = cljs_core.Seq.Arity1IQ(seq__806_1349)
																							_ = temp__4388__auto___1357___2
																							if cljs_core.Truth_(temp__4388__auto___1357___2) {
																								{
																									var seq__806_1358___1 = temp__4388__auto___1357___2
																									_ = seq__806_1358___1
																									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__806_1358___1) {
																										{
																											var c__970__auto___1359 = cljs_core.Chunk_first.X_invoke_Arity1(seq__806_1358___1)
																											_ = c__970__auto___1359
																											seq__806_1349, chunk__807_1350, count__808_1351, i__809_1352 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__806_1358___1), c__970__auto___1359, cljs_core.Count.X_invoke_Arity1(c__970__auto___1359).(float64), float64(0)
																											continue
																										}
																									} else {
																										{
																											var hour_1360 = cljs_core.First.X_invoke_Arity1(seq__806_1358___1)
																											_ = hour_1360
																											{
																												var pad_1361 = func(G__1363 *cljs_core.AFn, seq__806_1349 interface{}, chunk__807_1350 interface{}, count__808_1351 float64, i__809_1352 float64, seq__805_1325 interface{}, chunk__810_1326 interface{}, count__811_1327 float64, i__812_1328 float64, seq__804_1277 interface{}, chunk__817_1278 interface{}, count__818_1279 float64, i__819_1280 float64, hour_1360 interface{}, seq__806_1358___1 interface{}, temp__4388__auto___1357___2 cljs_core.CljsCoreISeq, day_1348 interface{}, seq__805_1346___1 interface{}, temp__4388__auto___1345___1 cljs_core.CljsCoreISeq, month_1324 interface{}, seq__804_1322___1 interface{}, temp__4388__auto___1321 cljs_core.CljsCoreISeq) *cljs_core.AFn {
																													return cljs_core.Fn(G__1363, 1, func(n interface{}) interface{} {
																														if n.(float64) < float64(10) {
																															return strings.Join([]string{cljs_core.Str.X_invoke_Arity1("0").(string), cljs_core.Str.X_invoke_Arity1(n).(string)}, ``)
																														} else {
																															return n
																														}
																													})
																												}(&cljs_core.AFn{}, seq__806_1349, chunk__807_1350, count__808_1351, i__809_1352, seq__805_1325, chunk__810_1326, count__811_1327, i__812_1328, seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280, hour_1360, seq__806_1358___1, temp__4388__auto___1357___2, day_1348, seq__805_1346___1, temp__4388__auto___1345___1, month_1324, seq__804_1322___1, temp__4388__auto___1321)
																												var inst_1362 = strings.Join([]string{cljs_core.Str.X_invoke_Arity1("2010-").(string), cljs_core.Str.X_invoke_Arity1(pad_1361.X_invoke_Arity1(month_1324)).(string), cljs_core.Str.X_invoke_Arity1("-").(string), cljs_core.Str.X_invoke_Arity1(pad_1361.X_invoke_Arity1(day_1348)).(string), cljs_core.Str.X_invoke_Arity1("T").(string), cljs_core.Str.X_invoke_Arity1(pad_1361.X_invoke_Arity1(hour_1360)).(string), cljs_core.Str.X_invoke_Arity1(":14:15.666-00:00").(string)}, ``)
																												_, _ = pad_1361, inst_1362
																												if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&js.Date{inst_1362})})).(string), strings.Join([]string{cljs_core.Str.X_invoke_Arity1("#inst \"").(string), cljs_core.Str.X_invoke_Arity1(inst_1362).(string), cljs_core.Str.X_invoke_Arity1("\"").(string)}, ``)) {
																												} else {
																													panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (js/Date. inst)) (str \"#inst \\\"\" inst \"\\\"\"))").(string)}, ``)}))
																												}
																											}
																											seq__806_1349, chunk__807_1350, count__808_1351, i__809_1352 = cljs_core.Next.Arity1IQ(seq__806_1358___1), nil, float64(0), float64(0)
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
																			seq__805_1325, chunk__810_1326, count__811_1327, i__812_1328 = cljs_core.Next.Arity1IQ(seq__805_1346___1), nil, float64(0), float64(0)
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
											seq__804_1277, chunk__817_1278, count__818_1279, i__819_1280 = cljs_core.Next.Arity1IQ(seq__804_1322___1), nil, float64(0), float64(0)
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
				var uuid_str_1364 = "550e8400-e29b-41d4-a716-446655440000"
				var uuid_1365 = (&cljs_core.CljsCoreUUID{uuid_str_1364})
				_, _ = uuid_str_1364, uuid_1365
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{uuid_1365})).(string), strings.Join([]string{cljs_core.Str.X_invoke_Arity1("#uuid \"").(string), cljs_core.Str.X_invoke_Arity1(uuid_str_1364).(string), cljs_core.Str.X_invoke_Arity1("\"").(string)}, ``)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str uuid) (str \"#uuid \\\"\" uuid-str \"\\\"\"))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Pr_str.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Rest.Arity1IQ(cljs_core.Conj.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentQueue_EMPTY, float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})))})).(string), "(2 3)") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (pr-str (rest (conj (.-EMPTY cljs.core/PersistentQueue) 1 2 3))) \"(2 3)\")").(string)}, ``)}))
			}
			X_bar = func(_bar *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(_bar, 2, func(this interface{}, x interface{}) interface{} {
					return cljs_core.Decorate_(this).(CljsCore_testIBar).X_bar_Arity2(x)
				})
			}(&cljs_core.AFn{})

			Baz = func(baz *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(baz, 1, func(f interface{}) interface{} {
					X__GT_t837 = func(__GT_t837 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t837, 4, func(f___1 interface{}, baz___1 interface{}, test_stuff___1 interface{}, meta838 interface{}) interface{} {
							return (&CljsCore_testT837{f___1, baz___1, test_stuff___1, meta838})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT837{f, baz, test_stuff, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(5), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-column", Fqn: "end-column", X_hash: float64(1425389514)}), float64(16), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-line", Fqn: "end-line", X_hash: float64(1837326455)}), float64(1837), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "column", Fqn: "column", X_hash: float64(2078222095)}), float64(5), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "line", Fqn: "line", X_hash: float64(212345235)}), float64(1834), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "file", Fqn: "file", X_hash: float64(-1269645878)}), "/home/hraberg/go/src/github.com/hraberg/cljs2go/test/cljs/core_test.cljs"}, nil})})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(float64(2), Baz.X_invoke_Arity1(cljs_core.Inc).(*CljsCore_testT837).X_bar_Arity2(float64(1))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 (-bar (baz inc) 1))").(string)}, ``)}))
			}
			{
				var x_1366 = "original"
				_ = x_1366
				Original_closure_stmt = func(original_closure_stmt *cljs_core.AFn, x_1366 string) *cljs_core.AFn {
					return cljs_core.Fn(original_closure_stmt, 0, func() interface{} {
						return x_1366
					})
				}(&cljs_core.AFn{}, x_1366)

			}
			{
				var x_1367 = "overwritten"
				_ = x_1367
				if cljs_core.X_EQ_.Arity2IIB("original", Original_closure_stmt.X_invoke_Arity0().(string)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"original\" (original-closure-stmt))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB("original", func() string {
				var x = "original"
				var oce = func(G__1368 *cljs_core.AFn, x string) *cljs_core.AFn {
					return cljs_core.Fn(G__1368, 0, func() interface{} {
						return x
					})
				}(&cljs_core.AFn{}, x)
				var x___1 = "overwritten"
				_, _, _ = x, oce, x___1
				return oce.X_invoke_Arity0().(string)
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"original\" (let [x \"original\" oce (fn [] x) x \"overwritten\"] (oce)))").(string)}, ``)}))
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
					var x_1369___1 = func(G__1370 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(G__1370, 0, func() interface{} {
							return "overwritten"
						})
					}(&cljs_core.AFn{})
					_ = x_1369___1
					if cljs_core.X_EQ_.Arity2IIB("original", y.X_invoke_Arity0().(string)) {
					} else {
						panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"original\" (y))").(string)}, ``)}))
					}
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reduce_kv.X_invoke_Arity3(cljs_core.Conj, cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}))), (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (reduce-kv conj [] (sorted-map :foo 1 :bar 2)) [:bar 2 :foo 1])").(string)}, ``)}))
			}
			{
				var kvr_test *cljs_core.AFn
				kvr_test = func(kvr_test *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(kvr_test, 2, func(data interface{}, expect interface{}) interface{} {
						if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)}), cljs_core.Reduce_kv.X_invoke_Arity3(func(G__1371 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1371, 3, func(___ interface{}, ______1 interface{}, ______2 interface{}) interface{} {
								return cljs_core.Reduced.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "reduced", Fqn: "reduced", X_hash: float64(1465210961)})).(*cljs_core.CljsCoreReduced)
							})
						}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentVector_EMPTY, data)) {
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :reduced (reduce-kv (fn [_ _ _] (reduced :reduced)) [] data))").(string)}, ``)}))
						}
						if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sort.X_invoke_Arity1(expect), cljs_core.Sort.X_invoke_Arity1(cljs_core.Reduce_kv.X_invoke_Arity3(func(G__1372 *cljs_core.AFn) *cljs_core.AFn {
							return cljs_core.Fn(G__1372, 3, func(r interface{}, k interface{}, v interface{}) interface{} {
								return cljs_core.Conj.X_invoke_Arity2(r, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{k, v}, nil}))
							})
						}(&cljs_core.AFn{}), cljs_core.CljsCorePersistentVector_EMPTY, data))) {
							return nil
						} else {
							panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sort expect) (sort (reduce-kv (fn [r k v] (-> r (conj [k v]))) [] data)))").(string)}, ``)}))
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:init :val} (reduce-kv assoc {:init :val} nil))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), func() (return__1373 interface{}) {
				defer func() {
					if e842 := recover(); e842 != nil {
						if cljs_core.Value_(e842).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreExceptionInfo)(nil)).Elem()) {
							{
								var e = e842
								_ = e
								return__1373 = cljs_core.Ex_data.X_invoke_Arity1(e)
							}
						} else {
							panic(e842)

						}
					}
				}()
				{
					panic(cljs_core.Ex_info.X_invoke_Arity2("asdf", (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})).(*cljs_core.CljsCoreExceptionInfo))
				}
			}()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1} (try (throw (ex-info \"asdf\" {:foo 1})) (catch ExceptionInfo e (ex-data e))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Assoc.X_invoke_ArityVariadic(cljs_core.CljsCorePersistentArrayMap_EMPTY, float64(154618822656), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(261993005056), float64(1)})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(154618822656), float64(1), float64(261993005056), float64(1)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (assoc {} 154618822656 1 261993005056 1) {154618822656 1, 261993005056 1})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get_in.X_invoke_Arity3((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(1)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})}, nil}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nothing-there", Fqn: "nothing-there", X_hash: float64(1302692654)})), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nothing-there", Fqn: "nothing-there", X_hash: float64(1302692654)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get-in {:a {:b 1}} [:a :b :c] :nothing-there) :nothing-there)").(string)}, ``)}))
			}
			if cljs_core.Nil_(cljs_core.Get_in.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get-in {:foo {:bar 2}} [:foo :bar :baz]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "tag", Fqn: "tag", X_hash: float64(-1290361223)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "int", Str: "int", X_hash: float64(-100885395), X_meta: nil})}, nil}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "tag", Fqn: "tag", X_hash: float64(-1290361223)}), (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "int", Str: "int", X_hash: float64(-100885395), X_meta: nil})}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (meta (with-meta (quote foo) {:tag (quote int)})) {:tag (quote int)})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Reduce_kv.X_invoke_Arity3(cljs_core.X_PLUS_, float64(0), cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, cljs_core.Range_.X_invoke_Arity1(float64(1000)).(*cljs_core.CljsCoreRange))), cljs_core.Reduce.X_invoke_Arity2(cljs_core.X_PLUS_, cljs_core.Range_.X_invoke_Arity1(float64(1000)).(*cljs_core.CljsCoreRange))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (reduce-kv + 0 (apply hash-map (range 1000))) (reduce + (range 1000)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(1), float64(2)}, nil}), func(G__1374 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1374, 0, func(more__ ...interface{}) interface{} {
					var more = cljs_core.Seq.Arity1IQ(more__[0])
					_ = more
					return more
				})
			}(&cljs_core.AFn{}).X_invoke_Arity3(js.Undefined, float64(1), float64(2))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [js/undefined 1 2] ((fn [& more] more) js/undefined 1 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{js.Undefined, float64(4), float64(5)}, nil}), func(G__1375 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1375, 2, func(a_b_more__ ...interface{}) interface{} {
					var a = a_b_more__[0]
					var b = a_b_more__[1]
					var more = cljs_core.Seq.Arity1IQ(a_b_more__[2])
					_, _, _ = a, b, more
					return more
				})
			}(&cljs_core.AFn{}).X_invoke_Arity5(float64(1), float64(2), js.Undefined, float64(4), float64(5))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [js/undefined 4 5] ((fn [a b & more] more) 1 2 js/undefined 4 5))").(string)}, ``)}))
			}
			if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(float64(42), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "anything", Fqn: "anything", X_hash: float64(285898971)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get 42 :anything))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Get.X_invoke_Arity3(float64(42), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "anything", Fqn: "anything", X_hash: float64(285898971)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-found", Fqn: "not-found", X_hash: float64(-629079980)})), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-found", Fqn: "not-found", X_hash: float64(-629079980)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (get 42 :anything :not-found) :not-found)").(string)}, ``)}))
			}
			if cljs_core.Nil_(cljs_core.First.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity3(cljs_core.Get, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "anything", Fqn: "anything", X_hash: float64(285898971)})}, nil})).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (first (map get [42] [:anything])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.First.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity4(cljs_core.Get, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(42)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "anything", Fqn: "anything", X_hash: float64(285898971)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-found", Fqn: "not-found", X_hash: float64(-629079980)})}, nil})).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "not-found", Fqn: "not-found", X_hash: float64(-629079980)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (first (map get [42] [:anything] [:not-found])) :not-found)").(string)}, ``)}))
			}
			{
				var fs_1376 = cljs_core.Atom.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreAtom)
				_ = fs_1376
				{
					var seq__843_1377 interface{} = cljs_core.Seq.Arity1IQ(cljs_core.Range_.X_invoke_Arity1(float64(4)).(*cljs_core.CljsCoreRange))
					var chunk__845_1378 interface{} = nil
					var count__846_1379 = float64(0)
					var i__847_1380 = float64(0)
					_, _, _, _ = seq__843_1377, chunk__845_1378, count__846_1379, i__847_1380
					for {
						if i__847_1380 < count__846_1379 {
							{
								var x_1381 = cljs_core.Decorate_(chunk__845_1378).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__847_1380)
								_ = x_1381
								{
									var y_1382 = (x_1381.(float64) + float64(1))
									var f_1383 = func(G__1384 *cljs_core.AFn, seq__843_1377 interface{}, chunk__845_1378 interface{}, count__846_1379 float64, i__847_1380 float64, y_1382 float64, x_1381 interface{}, fs_1376 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
										return cljs_core.Fn(G__1384, 0, func() interface{} {
											return y_1382
										})
									}(&cljs_core.AFn{}, seq__843_1377, chunk__845_1378, count__846_1379, i__847_1380, y_1382, x_1381, fs_1376)
									_, _ = y_1382, f_1383
									cljs_core.Swap_BANG_.X_invoke_Arity3(fs_1376, cljs_core.Conj, f_1383)
								}
								seq__843_1377, chunk__845_1378, count__846_1379, i__847_1380 = seq__843_1377, chunk__845_1378, count__846_1379, (i__847_1380 + float64(1))
								continue
							}
						} else {
							{
								var temp__4388__auto___1385 = cljs_core.Seq.Arity1IQ(seq__843_1377)
								_ = temp__4388__auto___1385
								if cljs_core.Truth_(temp__4388__auto___1385) {
									{
										var seq__843_1386___1 = temp__4388__auto___1385
										_ = seq__843_1386___1
										if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__843_1386___1) {
											{
												var c__970__auto___1387 = cljs_core.Chunk_first.X_invoke_Arity1(seq__843_1386___1)
												_ = c__970__auto___1387
												seq__843_1377, chunk__845_1378, count__846_1379, i__847_1380 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__843_1386___1), c__970__auto___1387, cljs_core.Count.X_invoke_Arity1(c__970__auto___1387).(float64), float64(0)
												continue
											}
										} else {
											{
												var x_1388 = cljs_core.First.X_invoke_Arity1(seq__843_1386___1)
												_ = x_1388
												{
													var y_1389 = (x_1388.(float64) + float64(1))
													var f_1390 = func(G__1391 *cljs_core.AFn, seq__843_1377 interface{}, chunk__845_1378 interface{}, count__846_1379 float64, i__847_1380 float64, y_1389 float64, x_1388 interface{}, seq__843_1386___1 interface{}, temp__4388__auto___1385 cljs_core.CljsCoreISeq, fs_1376 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
														return cljs_core.Fn(G__1391, 0, func() interface{} {
															return y_1389
														})
													}(&cljs_core.AFn{}, seq__843_1377, chunk__845_1378, count__846_1379, i__847_1380, y_1389, x_1388, seq__843_1386___1, temp__4388__auto___1385, fs_1376)
													_, _ = y_1389, f_1390
													cljs_core.Swap_BANG_.X_invoke_Arity3(fs_1376, cljs_core.Conj, f_1390)
												}
												seq__843_1377, chunk__845_1378, count__846_1379, i__847_1380 = cljs_core.Next.Arity1IQ(seq__843_1386___1), nil, float64(0), float64(0)
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
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(func(G__1392 *cljs_core.AFn, fs_1376 *cljs_core.CljsCoreAtom) *cljs_core.AFn {
					return cljs_core.Fn(G__1392, 1, func(p1__79_SHARP_ interface{}) interface{} {
						{
							return p1__79_SHARP_.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
						}
					})
				}(&cljs_core.AFn{}, fs_1376), cljs_core.Deref.X_invoke_Arity1(fs_1376)).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map (fn* [p1__79#] (p1__79#)) (clojure.core/deref fs)) (quote (1 2 3 4)))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Char.X_invoke_Arity1(float64(65)), "A") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (char 65) \\A)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Char.X_invoke_Arity1("A"), "A") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (char \\A) \\A)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})), cljs_core.Hash.X_invoke_Arity1(cljs_core.Symbol.X_invoke_Arity1("foo"))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (hash (quote foo)) (hash (symbol \"foo\")))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: "foo", Name: "bar", Str: "foo/bar", X_hash: float64(254379989), X_meta: nil})), cljs_core.Hash.X_invoke_Arity1(cljs_core.Symbol.X_invoke_Arity2("foo", "bar").(*cljs_core.CljsCoreSymbol))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (hash (quote foo/bar)) (hash (symbol \"foo\" \"bar\")))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Concat.X_invoke_ArityVariadic((&cljs_core.CljsCoreLazySeq{nil, func(G__1393 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1393, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), (&cljs_core.CljsCoreLazySeq{nil, func(G__1394 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1394, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreLazySeq{nil, func(G__1395 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1395, 0, func() interface{} {
					return (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3)}, nil})
				})
			}(&cljs_core.AFn{}), nil, nil})})).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (lazy-cat [1] [2] [3]) (quote (1 2 3)))").(string)}, ``)}))
			}
			X__GT_PositionalFactoryTest = func(__GT_PositionalFactoryTest *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_PositionalFactoryTest, 1, func(x interface{}) interface{} {
					return (&CljsCore_testPositionalFactoryTest{x})
				})
			}(&cljs_core.AFn{})

			if float64(1) == X__GT_PositionalFactoryTest.X_invoke_Arity1(float64(1)).(*CljsCore_testPositionalFactoryTest).X.(float64) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== 1 (.-x (->PositionalFactoryTest 1)))").(string)}, ``)}))
			}
			if cljs_core.Nil_((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "test", Fqn: "test", X_hash: float64(577538877)}).X_invoke_Arity1("test")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (:test \"test\"))").(string)}, ``)}))
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
					var sb__1140__auto__ = (&goog_string.StringBuffer{})
					_ = sb__1140__auto__
					{
						var _STAR_print_fn_STAR_849_1396 = cljs_core.X_STAR_print_fn_STAR_
						_ = _STAR_print_fn_STAR_849_1396
						func() {
							defer func() {
								cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_849_1396

							}()
							{
								cljs_core.X_STAR_print_fn_STAR_ = func(G__1397 *cljs_core.AFn, _STAR_print_fn_STAR_849_1396 interface{}, sb__1140__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
									return cljs_core.Fn(G__1397, 1, func(x__1141__auto__ interface{}) interface{} {
										return sb__1140__auto__.Append(x__1141__auto__)
									})
								}(&cljs_core.AFn{}, _STAR_print_fn_STAR_849_1396, sb__1140__auto__)

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
					return strings.Join([]string{cljs_core.Str.X_invoke_Arity1(sb__1140__auto__).(string)}, ``)
				}()) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"ffgfg\" (with-out-str (instance? Symbol (f! (quote foo))) (max (f! 5) (g! 10)) (min (f! 5) (g! 10))))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), cljs_core.Set.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{1 2} (set [1 2 2]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{float64(1), float64(2), float64(2)}, true)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{1 2} (hash-set 1 2 2))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(1), nil, float64(2), nil}, nil}, nil}), cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_set, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= #{1 2} (apply hash-set [1 2 2]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Last.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(cljs_core.Identity, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Range_.X_invoke_Arity1(float64(32)).(*cljs_core.CljsCoreRange))).(*cljs_core.CljsCoreLazySeq)), float64(31)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (last (map identity (into [] (range 32)))) 31)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentHashSet_EMPTY, cljs_core.Range_.X_invoke_Arity1(float64(32)).(*cljs_core.CljsCoreRange)), cljs_core.Set.X_invoke_Arity1(cljs_core.Map_.X_invoke_Arity2(cljs_core.Identity, cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Range_.X_invoke_Arity1(float64(32)).(*cljs_core.CljsCoreRange))).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (into #{} (range 32)) (set (map identity (into [] (range 32)))))").(string)}, ``)}))
			}

			Foo580 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), func(G__1398 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1398, 0, func() interface{} {
					return nil
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), func(G__1399 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1399, 0, func() interface{} {
					{
						var G__850 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
						_ = G__850
						return Foo580.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__850)
					}
				})
			}(&cljs_core.AFn{})}, nil})

			if cljs_core.Nil_((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}).X_invoke_Arity1(Foo580).(cljs_core.CljsCoreIFn).X_invoke_Arity0().(cljs_core.CljsCoreIFn).X_invoke_Arity0()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (((:b foo580))))").(string)}, ``)}))
			}
			if (cljs_core.First.X_invoke_Arity1(cljs_core.Filter.X_invoke_Arity2(func(G__1400 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1400, 1, func(p1__80_SHARP_ interface{}) interface{} {
					return (p1__80_SHARP_.(float64) == float64(9999))
				})
			}(&cljs_core.AFn{}), cljs_core.Range_.X_invoke_Arity0().(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq)).(float64) == float64(9999)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (first (filter (fn* [p1__80#] (== p1__80# 9999)) (range))) 9999)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Concat.X_invoke_Arity2(nil, cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= () (concat nil []))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreIEmptyList(cljs_core.CljsCoreList_EMPTY), cljs_core.Concat.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.CljsCorePersistentVector_EMPTY).(*cljs_core.CljsCoreLazySeq)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= () (concat [] []))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("foobar", cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Concat.X_invoke_Arity2("foo", "bar").(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"foobar\" (apply str (concat \"foo\" \"bar\")))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{""})).(*cljs_core.CljsCoreList), cljs_core.Re_seq.X_invoke_Arity2((&js.RegExp{Pattern: `\s*`, Flags: ``}), "")) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (quote (\"\")) (re-seq #\"\\s*\" \"\"))").(string)}, ``)}))
			}
			X__GT_KeywordTest = func(__GT_KeywordTest *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(__GT_KeywordTest, 0, func() interface{} {
					return (&CljsCore_testKeywordTest{})
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}).X_invoke_Arity1((&CljsCore_testKeywordTest{})), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "nothing", Fqn: "nothing", X_hash: float64(-1022703296)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (:a (KeywordTest.)) :nothing)").(string)}, ``)}))
			}
			{
				var a_1401 = func() *CljsCore_testT851 {
					X__GT_t851 = func(__GT_t851 *cljs_core.AFn) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t851, 2, func(test_stuff___1 interface{}, meta852 interface{}) interface{} {
							return (&CljsCore_testT851{test_stuff___1, meta852})
						})
					}(&cljs_core.AFn{})

					return (&CljsCore_testT851{test_stuff, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(5), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-column", Fqn: "end-column", X_hash: float64(1425389514)}), float64(39), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-line", Fqn: "end-line", X_hash: float64(1837326455)}), float64(2029), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "column", Fqn: "column", X_hash: float64(2078222095)}), float64(11), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "line", Fqn: "line", X_hash: float64(212345235)}), float64(2029), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "file", Fqn: "file", X_hash: float64(-1269645878)}), "/home/hraberg/go/src/github.com/hraberg/cljs2go/test/cljs/core_test.cljs"}, nil})})
				}()
				var b_1402 = func() *CljsCore_testT854 {
					X__GT_t854 = func(__GT_t854 *cljs_core.AFn, a_1401 *CljsCore_testT851) *cljs_core.AFn {
						return cljs_core.Fn(__GT_t854, 3, func(a___1 interface{}, test_stuff___1 interface{}, meta855 interface{}) interface{} {
							return (&CljsCore_testT854{a___1, test_stuff___1, meta855})
						})
					}(&cljs_core.AFn{}, a_1401)

					return (&CljsCore_testT854{a_1401, test_stuff, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(5), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-column", Fqn: "end-column", X_hash: float64(1425389514)}), float64(39), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-line", Fqn: "end-line", X_hash: float64(1837326455)}), float64(2030), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "column", Fqn: "column", X_hash: float64(2078222095)}), float64(11), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "line", Fqn: "line", X_hash: float64(212345235)}), float64(2030), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "file", Fqn: "file", X_hash: float64(-1269645878)}), "/home/hraberg/go/src/github.com/hraberg/cljs2go/test/cljs/core_test.cljs"}, nil})})
				}()
				var s_1403 = cljs_core.Set.X_invoke_Arity1(cljs_core.Range_.X_invoke_Arity1(float64(128)).(*cljs_core.CljsCoreRange))
				_, _, _ = a_1401, b_1402, s_1403
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_1403, a_1401, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_1402}))), a_1401)), a_1401), cljs_core.Conj.X_invoke_Arity2(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1(cljs_core.Conj.X_invoke_ArityVariadic(s_1403, a_1401, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{b_1402}))), a_1401)), a_1401)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-> (conj s a b) transient (disj! a) persistent! (conj a)) (-> (conj s a b) transient (disj! a) persistent! (conj a)))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Juxt.X_invoke_Arity2(cljs_core.Namespace, cljs_core.Name).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1(cljs_core.Keyword.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "a.b", Str: "a.b", X_hash: float64(-2098083151), X_meta: nil}))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "a.b"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-> (quote a.b) keyword ((juxt namespace name))) [nil \"a.b\"])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Juxt.X_invoke_Arity2(cljs_core.Namespace, cljs_core.Name).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1(cljs_core.Keyword.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: "a.b", Name: "c", Str: "a.b/c", X_hash: float64(-122574001), X_meta: nil}))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a.b", "c"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-> (quote a.b/c) keyword ((juxt namespace name))) [\"a.b\" \"c\"])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Juxt.X_invoke_Arity2(cljs_core.Namespace, cljs_core.Name).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1(cljs_core.Keyword.X_invoke_Arity1("a.b")), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "a.b"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-> \"a.b\" keyword ((juxt namespace name))) [nil \"a.b\"])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Juxt.X_invoke_Arity2(cljs_core.Namespace, cljs_core.Name).(cljs_core.CljsCoreIFn).(cljs_core.CljsCoreIFn).X_invoke_Arity1(cljs_core.Keyword.X_invoke_Arity1("a.b/c")), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a.b", "c"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-> \"a.b/c\" keyword ((juxt namespace name))) [\"a.b\" \"c\"])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Keyword.X_invoke_Arity1(float64(123)), nil) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (keyword 123) nil)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Keyword.X_invoke_Arity1((&js.Date{})), nil) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (keyword (js/Date.)) nil)").(string)}, ``)}))
			}
			Some_x = float64(1)

			Some_y = float64(1)

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{Some_y, Some_x}, true).(*cljs_core.CljsCorePersistentHashSet)).(float64), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count #{some-y some-x}) 1)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Count.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_FromArray.X_invoke_Arity3([]interface{}{Some_x, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), Some_y, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, true, false).(*cljs_core.CljsCorePersistentArrayMap)).(float64), float64(1)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (count {some-x :foo, some-y :bar}) 1)").(string)}, ``)}))
			}
			if cljs_core.Value_([]interface{}{float64(1), float64(2), float64(3)}).Kind() == reflect.Slice {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(array? #js [1 2 3])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(len([]interface{}{float64(1), float64(2), float64(3)})), float64(3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (alength #js [1 2 3]) 3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ([]interface{}{float64(1), float64(2), float64(3)}), cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq #js [1 2 3]) (seq [1 2 3]))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Vector, cljs_core.Drop_while.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.X_EQ_, float64(1)).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply vector (drop-while (partial = 1) [1 2 3])) [2 3])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.List, cljs_core.Drop_while.X_invoke_Arity2(cljs_core.Partial.X_invoke_Arity2(cljs_core.X_EQ_, float64(1)).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply list (drop-while (partial = 1) [1 2 3])) (quote (2 3)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Drop.X_invoke_Arity2(float64(1), []interface{}{float64(1), float64(2), float64(3)}).(*cljs_core.CljsCoreLazySeq)), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{float64(3), nil, float64(2), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (set (drop 1 #js [1 2 3])) #{3 2})").(string)}, ``)}))
			}
			if cljs_core.Nil_(cljs_core.First.X_invoke_Arity1(cljs_core.Rest.Arity1IQ(cljs_core.Rest.Arity1IQ(cljs_core.Rest.Arity1IQ(cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange)))))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (first (rest (rest (rest (range 3))))))").(string)}, ``)}))
			}
			if cljs_core.Count.X_invoke_Arity1(cljs_core.CljsCorePersistentHashSet_FromArray.X_invoke_Arity2([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(0)}, nil})}, true)).(float64) == cljs_core.Count.X_invoke_Arity1(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.Decorate_(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(0)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(0)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(0)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(1)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(2)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(3)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(4)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(4)}, nil}))).(cljs_core.CljsCoreICollection).X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(4)}, nil}))).(float64) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (count (hash-set [1 4] [2 4] [3 4] [0 3] [1 3] [2 3] [3 3] [0 2] [1 2] [2 2] [3 2] [4 2] [0 1] [1 1] [2 1] [3 1] [1 0] [2 0] [3 0])) (count (list [1 4] [2 4] [3 4] [0 3] [1 3] [2 3] [3 3] [0 2] [1 2] [2 2] [3 2] [4 2] [0 1] [1 1] [2 1] [3 1] [1 0] [2 0] [3 0])))").(string)}, ``)}))
			}
			X_woz = func(_woz *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(_woz, 1, func(this interface{}) interface{} {
					return cljs_core.Decorate_(this).(CljsCore_testIWoz).X_woz_Arity1()
				})
			}(&cljs_core.AFn{})

			Noz = cljs_core.CljsCorePersistentVector_EMPTY

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Conj_BANG_.X_invoke_ArityVariadic(cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)}))), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-> (transient []) (conj! 1 2) persistent!) [1 2])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Disj_BANG_.X_invoke_ArityVariadic(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(1), nil, float64(3), nil, float64(2), nil}, nil}, nil})), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2)}))), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(3), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-> (transient #{1 3 2}) (disj! 1 2) persistent!) #{3})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Assoc_BANG_.X_invoke_ArityVariadic(cljs_core.Transient.X_invoke_Arity1(cljs_core.CljsCorePersistentArrayMap_EMPTY), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-> (transient {}) (assoc! :a 1 :b 2) persistent!) {:a 1, :b 2})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Persistent_BANG_.X_invoke_Arity1(cljs_core.Dissoc_BANG_.X_invoke_ArityVariadic(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(3)}, nil})), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})}))), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), float64(3)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-> (transient {:a 1, :b 2, :c 3}) (dissoc! :a :b) persistent!) {:c 3})").(string)}, ``)}))
			}
			{
				var seq__857_1404 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__858_1405 interface{} = nil
				var count__859_1406 = float64(0)
				var i__860_1407 = float64(0)
				_, _, _, _ = seq__857_1404, chunk__858_1405, count__859_1406, i__860_1407
				for {
					if i__860_1407 < count__859_1406 {
						{
							var n_1408 = cljs_core.Decorate_(chunk__858_1405).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__860_1407)
							_ = n_1408
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1409 interface{}) {
								defer func() {
									if e863 := recover(); e863 != nil {
										if cljs_core.Value_(e863).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e863
												_ = e
												return__1409 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e863)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1408, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1410 interface{}) {
								defer func() {
									if e864 := recover(); e864 != nil {
										if cljs_core.Value_(e864).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e864
												_ = e
												return__1410 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e864)

										}
									}
								}()
								{
									return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_1408, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							seq__857_1404, chunk__858_1405, count__859_1406, i__860_1407 = seq__857_1404, chunk__858_1405, count__859_1406, (i__860_1407 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1411 = cljs_core.Seq.Arity1IQ(seq__857_1404)
							_ = temp__4388__auto___1411
							if cljs_core.Truth_(temp__4388__auto___1411) {
								{
									var seq__857_1412___1 = temp__4388__auto___1411
									_ = seq__857_1412___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__857_1412___1) {
										{
											var c__970__auto___1413 = cljs_core.Chunk_first.X_invoke_Arity1(seq__857_1412___1)
											_ = c__970__auto___1413
											seq__857_1404, chunk__858_1405, count__859_1406, i__860_1407 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__857_1412___1), c__970__auto___1413, cljs_core.Count.X_invoke_Arity1(c__970__auto___1413).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1414 = cljs_core.First.X_invoke_Arity1(seq__857_1412___1)
											_ = n_1414
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1415 interface{}) {
												defer func() {
													if e865 := recover(); e865 != nil {
														if cljs_core.Value_(e865).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e865
																_ = e
																return__1415 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e865)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1414, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (assoc [1 2] n 4) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1416 interface{}) {
												defer func() {
													if e866 := recover(); e866 != nil {
														if cljs_core.Value_(e866).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e866
																_ = e
																return__1416 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e866)

														}
													}
												}()
												{
													return cljs_core.Assoc.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), float64(2)).(*cljs_core.CljsCoreSubvec), n_1414, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (assoc (subvec [1 2 3] 2) n 4) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											seq__857_1404, chunk__858_1405, count__859_1406, i__860_1407 = cljs_core.Next.Arity1IQ(seq__857_1412___1), nil, float64(0), float64(0)
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
				var seq__867_1417 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__868_1418 interface{} = nil
				var count__869_1419 = float64(0)
				var i__870_1420 = float64(0)
				_, _, _, _ = seq__867_1417, chunk__868_1418, count__869_1419, i__870_1420
				for {
					if i__870_1420 < count__869_1419 {
						{
							var n_1421 = cljs_core.Decorate_(chunk__868_1418).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__870_1420)
							_ = n_1421
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1422 interface{}) {
								defer func() {
									if e873 := recover(); e873 != nil {
										if cljs_core.Value_(e873).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e873
												_ = e
												return__1422 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e873)

										}
									}
								}()
								{
									return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1421, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							seq__867_1417, chunk__868_1418, count__869_1419, i__870_1420 = seq__867_1417, chunk__868_1418, count__869_1419, (i__870_1420 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1423 = cljs_core.Seq.Arity1IQ(seq__867_1417)
							_ = temp__4388__auto___1423
							if cljs_core.Truth_(temp__4388__auto___1423) {
								{
									var seq__867_1424___1 = temp__4388__auto___1423
									_ = seq__867_1424___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__867_1424___1) {
										{
											var c__970__auto___1425 = cljs_core.Chunk_first.X_invoke_Arity1(seq__867_1424___1)
											_ = c__970__auto___1425
											seq__867_1417, chunk__868_1418, count__869_1419, i__870_1420 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__867_1424___1), c__970__auto___1425, cljs_core.Count.X_invoke_Arity1(c__970__auto___1425).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1426 = cljs_core.First.X_invoke_Arity1(seq__867_1424___1)
											_ = n_1426
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1427 interface{}) {
												defer func() {
													if e874 := recover(); e874 != nil {
														if cljs_core.Value_(e874).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e874
																_ = e
																return__1427 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e874)

														}
													}
												}()
												{
													return cljs_core.Assoc_BANG_.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1426, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (assoc! (transient [1 2]) n 4) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											seq__867_1417, chunk__868_1418, count__869_1419, i__870_1420 = cljs_core.Next.Arity1IQ(seq__867_1424___1), nil, float64(0), float64(0)
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
				var map__875_1428 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), float64(2)}, nil})
				var map__875_1429___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__875_1428) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__875_1428)
					} else {
						return map__875_1428
					}
				}()
				var b_1430 = cljs_core.Get.X_invoke_Arity2(map__875_1429___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}))
				var a_1431 = cljs_core.Get.X_invoke_Arity2(map__875_1429___1, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}))
				_, _, _, _ = map__875_1428, map__875_1429___1, b_1430, a_1431
				if cljs_core.X_EQ_.Arity2IIB(float64(1), a_1431) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 a)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), b_1430) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 b)").(string)}, ``)}))
				}
			}
			{
				var map__876_1432 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__876_1433___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__876_1432) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__876_1432)
					} else {
						return map__876_1432
					}
				}()
				var d_1434 = cljs_core.Get.X_invoke_Arity2(map__876_1433___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_1435 = cljs_core.Get.X_invoke_Arity2(map__876_1433___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__876_1432, map__876_1433___1, d_1434, b_1435
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1435) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 b)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1434) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 d)").(string)}, ``)}))
				}
			}
			{
				var map__877_1436 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}), float64(2)}, nil})
				var map__877_1437___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__877_1436) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__877_1436)
					} else {
						return map__877_1436
					}
				}()
				var d_1438 = cljs_core.Get.X_invoke_Arity2(map__877_1437___1, (&cljs_core.CljsCoreKeyword{Ns: "c", Name: "d", Fqn: "c/d", X_hash: float64(1972142513)}))
				var b_1439 = cljs_core.Get.X_invoke_Arity2(map__877_1437___1, (&cljs_core.CljsCoreKeyword{Ns: "a", Name: "b", Fqn: "a/b", X_hash: float64(1482224565)}))
				_, _, _, _ = map__877_1436, map__877_1437___1, d_1438, b_1439
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1439) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 b)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1438) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 d)").(string)}, ``)}))
				}
			}
			{
				var map__878_1440 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}), float64(1), (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}), float64(2)}, nil})
				var map__878_1441___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__878_1440) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__878_1440)
					} else {
						return map__878_1440
					}
				}()
				var d_1442 = cljs_core.Get.X_invoke_Arity2(map__878_1441___1, (&cljs_core.CljsCoreSymbol{Ns: "c", Name: "d", Str: "c/d", X_hash: float64(-682293256), X_meta: nil}))
				var b_1443 = cljs_core.Get.X_invoke_Arity2(map__878_1441___1, (&cljs_core.CljsCoreSymbol{Ns: "a", Name: "b", Str: "a/b", X_hash: float64(-1172211204), X_meta: nil}))
				_, _, _, _ = map__878_1440, map__878_1441___1, d_1442, b_1443
				if cljs_core.X_EQ_.Arity2IIB(float64(1), b_1443) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1 b)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(float64(2), d_1442) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 2 d)").(string)}, ``)}))
				}
			}
			{
				var map__879_1444 = (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}), float64(2)}, nil})
				var map__879_1445___1 = func() interface{} {
					if cljs_core.Seq_QMARK_.Arity1IB(map__879_1444) {
						return cljs_core.Apply.X_invoke_Arity2(cljs_core.Hash_map, map__879_1444)
					} else {
						return map__879_1444
					}
				}()
				var y_1446 = cljs_core.Get.X_invoke_Arity2(map__879_1445___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "y", Fqn: "clojure.string/y", X_hash: float64(1821360795)}))
				var x_1447 = cljs_core.Get.X_invoke_Arity2(map__879_1445___1, (&cljs_core.CljsCoreKeyword{Ns: "clojure.string", Name: "x", Fqn: "clojure.string/x", X_hash: float64(1710944900)}))
				_, _, _, _ = map__879_1444, map__879_1445___1, y_1446, x_1447
				if cljs_core.X_EQ_.Arity2IIB(x_1447, float64(1)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= x 1)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(y_1446, float64(2)) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= y 2)").(string)}, ``)}))
				}
			}
			Cljs_739 = func(cljs_739 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(cljs_739, 2, func(arr interface{}, names interface{}) interface{} {
					for {
						{
							var name = cljs_core.First.X_invoke_Arity1(names)
							_ = name
							if cljs_core.Truth_(name) {
								arr, names = cljs_core.Conj.X_invoke_Arity2(arr, func(G__1448 *cljs_core.AFn, arr interface{}, names interface{}, name interface{}) *cljs_core.AFn {
									return cljs_core.Fn(G__1448, 0, func() interface{} {
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
				var sb__1140__auto__ = (&goog_string.StringBuffer{})
				_ = sb__1140__auto__
				{
					var _STAR_print_fn_STAR_880_1449 = cljs_core.X_STAR_print_fn_STAR_
					_ = _STAR_print_fn_STAR_880_1449
					func() {
						defer func() {
							cljs_core.X_STAR_print_fn_STAR_ = _STAR_print_fn_STAR_880_1449

						}()
						{
							cljs_core.X_STAR_print_fn_STAR_ = func(G__1450 *cljs_core.AFn, _STAR_print_fn_STAR_880_1449 interface{}, sb__1140__auto__ *goog_string.StringBuffer) *cljs_core.AFn {
								return cljs_core.Fn(G__1450, 1, func(x__1141__auto__ interface{}) interface{} {
									return sb__1140__auto__.Append(x__1141__auto__)
								})
							}(&cljs_core.AFn{}, _STAR_print_fn_STAR_880_1449, sb__1140__auto__)

							{
								var seq__881_1451 interface{} = cljs_core.Seq.Arity1IQ(Cljs_739.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)})}, nil})))
								var chunk__882_1452 interface{} = nil
								var count__883_1453 = float64(0)
								var i__884_1454 = float64(0)
								_, _, _, _ = seq__881_1451, chunk__882_1452, count__883_1453, i__884_1454
								for {
									if i__884_1454 < count__883_1453 {
										{
											var fn_1455 = cljs_core.Decorate_(chunk__882_1452).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__884_1454)
											_ = fn_1455
											{
												fn_1455.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
											}
											seq__881_1451, chunk__882_1452, count__883_1453, i__884_1454 = seq__881_1451, chunk__882_1452, count__883_1453, (i__884_1454 + float64(1))
											continue
										}
									} else {
										{
											var temp__4388__auto___1456 = cljs_core.Seq.Arity1IQ(seq__881_1451)
											_ = temp__4388__auto___1456
											if cljs_core.Truth_(temp__4388__auto___1456) {
												{
													var seq__881_1457___1 = temp__4388__auto___1456
													_ = seq__881_1457___1
													if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__881_1457___1) {
														{
															var c__970__auto___1458 = cljs_core.Chunk_first.X_invoke_Arity1(seq__881_1457___1)
															_ = c__970__auto___1458
															seq__881_1451, chunk__882_1452, count__883_1453, i__884_1454 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__881_1457___1), c__970__auto___1458, cljs_core.Count.X_invoke_Arity1(c__970__auto___1458).(float64), float64(0)
															continue
														}
													} else {
														{
															var fn_1459 = cljs_core.First.X_invoke_Arity1(seq__881_1457___1)
															_ = fn_1459
															{
																fn_1459.(cljs_core.CljsCoreIFn).X_invoke_Arity0()
															}
															seq__881_1451, chunk__882_1452, count__883_1453, i__884_1454 = cljs_core.Next.Arity1IQ(seq__881_1457___1), nil, float64(0), float64(0)
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
				return strings.Join([]string{cljs_core.Str.X_invoke_Arity1(sb__1140__auto__).(string)}, ``)
			}(), ":a\n:b\n:c\n:d\n") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (with-out-str (doseq [fn (cljs-739 [] [:a :b :c :d])] (fn))) \":a\\n:b\\n:c\\n:d\\n\")").(string)}, ``)}))
			}
			{
				var seq__885_1460 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(8), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{nil, "-1", "", "0", "1", false, true, true}, nil}))
				var chunk__886_1461 interface{} = nil
				var count__887_1462 = float64(0)
				var i__888_1463 = float64(0)
				_, _, _, _ = seq__885_1460, chunk__886_1461, count__887_1462, i__888_1463
				for {
					if i__888_1463 < count__887_1462 {
						{
							var n_1464 = cljs_core.Decorate_(chunk__886_1461).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__888_1463)
							_ = n_1464
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1464)) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get [1 2] n))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1465 interface{}) {
								defer func() {
									if e891 := recover(); e891 != nil {
										if cljs_core.Value_(e891).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e891
												_ = e
												return__1465 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e891)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1464)
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth [1 2] n) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1464, float64(4))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get [1 2] n 4))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1466 interface{}) {
								defer func() {
									if e892 := recover(); e892 != nil {
										if cljs_core.Value_(e892).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e892
												_ = e
												return__1466 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e892)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1464, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1464)) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get (subvec [1 2] 1) n))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1467 interface{}) {
								defer func() {
									if e893 := recover(); e893 != nil {
										if cljs_core.Value_(e893).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e893
												_ = e
												return__1467 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e893)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1464)
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1464, float64(4))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get (subvec [1 2] 1) n 4))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1468 interface{}) {
								defer func() {
									if e894 := recover(); e894 != nil {
										if cljs_core.Value_(e894).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e894
												_ = e
												return__1468 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e894)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1464, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1464)) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get (transient [1 2]) n))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1469 interface{}) {
								defer func() {
									if e895 := recover(); e895 != nil {
										if cljs_core.Value_(e895).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e895
												_ = e
												return__1469 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e895)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1464)
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1464, float64(4))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get (transient [1 2]) n 4))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1470 interface{}) {
								defer func() {
									if e896 := recover(); e896 != nil {
										if cljs_core.Value_(e896).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e896
												_ = e
												return__1470 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e896)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1464, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1464)) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get (range 1 3) n))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1471 interface{}) {
								defer func() {
									if e897 := recover(); e897 != nil {
										if cljs_core.Value_(e897).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e897
												_ = e
												return__1471 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e897)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1464)
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1464, float64(4))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get (range 1 3) n 4))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1472 interface{}) {
								defer func() {
									if e898 := recover(); e898 != nil {
										if cljs_core.Value_(e898).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var e = e898
												_ = e
												return__1472 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
											}
										} else {
											panic(e898)

										}
									}
								}()
								{
									return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1464, float64(4))
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))").(string)}, ``)}))
							}
							seq__885_1460, chunk__886_1461, count__887_1462, i__888_1463 = seq__885_1460, chunk__886_1461, count__887_1462, (i__888_1463 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1473 = cljs_core.Seq.Arity1IQ(seq__885_1460)
							_ = temp__4388__auto___1473
							if cljs_core.Truth_(temp__4388__auto___1473) {
								{
									var seq__885_1474___1 = temp__4388__auto___1473
									_ = seq__885_1474___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__885_1474___1) {
										{
											var c__970__auto___1475 = cljs_core.Chunk_first.X_invoke_Arity1(seq__885_1474___1)
											_ = c__970__auto___1475
											seq__885_1460, chunk__886_1461, count__887_1462, i__888_1463 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__885_1474___1), c__970__auto___1475, cljs_core.Count.X_invoke_Arity1(c__970__auto___1475).(float64), float64(0)
											continue
										}
									} else {
										{
											var n_1476 = cljs_core.First.X_invoke_Arity1(seq__885_1474___1)
											_ = n_1476
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1476)) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get [1 2] n))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1477 interface{}) {
												defer func() {
													if e899 := recover(); e899 != nil {
														if cljs_core.Value_(e899).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e899
																_ = e
																return__1477 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e899)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1476)
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth [1 2] n) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1476, float64(4))) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get [1 2] n 4))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1478 interface{}) {
												defer func() {
													if e900 := recover(); e900 != nil {
														if cljs_core.Value_(e900).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e900
																_ = e
																return__1478 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e900)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), n_1476, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth [1 2] n 4) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1476)) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get (subvec [1 2] 1) n))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1479 interface{}) {
												defer func() {
													if e901 := recover(); e901 != nil {
														if cljs_core.Value_(e901).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e901
																_ = e
																return__1479 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e901)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1476)
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (subvec [1 2] 1) n) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1476, float64(4))) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get (subvec [1 2] 1) n 4))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1480 interface{}) {
												defer func() {
													if e902 := recover(); e902 != nil {
														if cljs_core.Value_(e902).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e902
																_ = e
																return__1480 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e902)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Subvec.X_invoke_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}), float64(1)).(*cljs_core.CljsCoreSubvec), n_1476, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (subvec [1 2] 1) n 4) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1476)) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get (transient [1 2]) n))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1481 interface{}) {
												defer func() {
													if e903 := recover(); e903 != nil {
														if cljs_core.Value_(e903).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e903
																_ = e
																return__1481 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e903)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1476)
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (transient [1 2]) n) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1476, float64(4))) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get (transient [1 2]) n 4))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1482 interface{}) {
												defer func() {
													if e904 := recover(); e904 != nil {
														if cljs_core.Value_(e904).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e904
																_ = e
																return__1482 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e904)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Transient.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})), n_1476, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (transient [1 2]) n 4) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											if cljs_core.Nil_(cljs_core.Get.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1476)) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (get (range 1 3) n))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1483 interface{}) {
												defer func() {
													if e905 := recover(); e905 != nil {
														if cljs_core.Value_(e905).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e905
																_ = e
																return__1483 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e905)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1476)
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (range 1 3) n) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB(float64(4), cljs_core.Get.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1476, float64(4))) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4 (get (range 1 3) n 4))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)}), func() (return__1484 interface{}) {
												defer func() {
													if e906 := recover(); e906 != nil {
														if cljs_core.Value_(e906).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var e = e906
																_ = e
																return__1484 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "fail", Fqn: "fail", X_hash: float64(1706214930)})
															}
														} else {
															panic(e906)

														}
													}
												}()
												{
													return cljs_core.Nth.X_invoke_Arity3(cljs_core.Range_.X_invoke_Arity2(float64(1), float64(3)).(*cljs_core.CljsCoreRange), n_1476, float64(4))
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :fail (try (nth (range 1 3) n 4) (catch js/Error e :fail)))").(string)}, ``)}))
											}
											seq__885_1460, chunk__886_1461, count__887_1462, i__888_1463 = cljs_core.Next.Arity1IQ(seq__885_1474___1), nil, float64(0), float64(0)
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
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (-rest (rseq [0])) ())").(string)}, ``)}))
			}
			if cljs_core.Nil_(cljs_core.Decorate_(cljs_core.Rseq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil}))).(cljs_core.CljsCoreINext).X_next_Arity1()) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(nil? (-next (rseq [0])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Set.X_invoke_Arity1(cljs_core.Rseq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0)}, nil}))), (&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{float64(0), nil}, nil}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (set (rseq [0])) #{0})").(string)}, ``)}))
			}
			Cljs_780 = cljs_core.Atom.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), cljs_core.With_meta.X_invoke_Arity2(cljs_core.CljsCorePersistentVector_EMPTY, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList)}, nil}))}, nil})).(*cljs_core.CljsCoreAtom)

			cljs_core.Swap_BANG_.X_invoke_ArityVariadic(Cljs_780, cljs_core.Update_in, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}), cljs_core.Vary_meta, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Update_in, (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil}), cljs_core.Vec}))
			{
				var x_1485 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}).X_invoke_Arity1(cljs_core.Meta.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}).X_invoke_Arity1(cljs_core.Deref.X_invoke_Arity1(Cljs_780))))
				_ = x_1485
				if cljs_core.Vector_QMARK_.Arity1IB(x_1485) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(vector? x)").(string)}, ``)}))
				}
				if cljs_core.X_EQ_.Arity2IIB(x_1485, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= x [1 2 3])").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Native_invoke_instance_method.X_invoke_Arity3((&cljs_core.CljsCoreUUID{Uuid: `550e8400-e29b-41d4-a716-446655440000`}), "ToString", []interface{}{}), "550e8400-e29b-41d4-a716-446655440000") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (.toString #uuid \"550e8400-e29b-41d4-a716-446655440000\") \"550e8400-e29b-41d4-a716-446655440000\")").(string)}, ``)}))
			}
			{
				var seq__907_1486 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.CljsCorePersistentArrayMap_EMPTY, cljs_core.CljsCorePersistentHashMap_EMPTY, cljs_core.Sorted_map.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{}))}, nil}))
				var chunk__908_1487 interface{} = nil
				var count__909_1488 = float64(0)
				var i__910_1489 = float64(0)
				_, _, _, _ = seq__907_1486, chunk__908_1487, count__909_1488, i__910_1489
				for {
					if i__910_1489 < count__909_1488 {
						{
							var m_1490 = cljs_core.Decorate_(chunk__908_1487).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__910_1489)
							_ = m_1490
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__1491 interface{}) {
								defer func() {
									if e911 := recover(); e911 != nil {
										if cljs_core.Value_(e911).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
											{
												var ___ = e911
												_ = ___
												return__1491 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
											}
										} else {
											panic(e911)

										}
									}
								}()
								{
									return cljs_core.Conj.X_invoke_Arity2(m_1490, "foo")
								}
							}()) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1490, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1} (conj m [:foo 1]))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1490, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1} (conj m {:foo 1}))").(string)}, ``)}))
							}
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1490, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1} (conj m (list [:foo 1])))").(string)}, ``)}))
							}
							seq__907_1486, chunk__908_1487, count__909_1488, i__910_1489 = seq__907_1486, chunk__908_1487, count__909_1488, (i__910_1489 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1492 = cljs_core.Seq.Arity1IQ(seq__907_1486)
							_ = temp__4388__auto___1492
							if cljs_core.Truth_(temp__4388__auto___1492) {
								{
									var seq__907_1493___1 = temp__4388__auto___1492
									_ = seq__907_1493___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__907_1493___1) {
										{
											var c__970__auto___1494 = cljs_core.Chunk_first.X_invoke_Arity1(seq__907_1493___1)
											_ = c__970__auto___1494
											seq__907_1486, chunk__908_1487, count__909_1488, i__910_1489 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__907_1493___1), c__970__auto___1494, cljs_core.Count.X_invoke_Arity1(c__970__auto___1494).(float64), float64(0)
											continue
										}
									} else {
										{
											var m_1495 = cljs_core.First.X_invoke_Arity1(seq__907_1493___1)
											_ = m_1495
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), func() (return__1496 interface{}) {
												defer func() {
													if e912 := recover(); e912 != nil {
														if cljs_core.Value_(e912).Type().AssignableTo(reflect.TypeOf((**js.Error)(nil)).Elem()) {
															{
																var ___ = e912
																_ = ___
																return__1496 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
															}
														} else {
															panic(e912)

														}
													}
												}()
												{
													return cljs_core.Conj.X_invoke_Arity2(m_1495, "foo")
												}
											}()) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :ok (try (conj m \"foo\") (catch js/Error _ :ok)))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1495, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1} (conj m [:foo 1]))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1495, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}))) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1} (conj m {:foo 1}))").(string)}, ``)}))
											}
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil}), cljs_core.Conj.X_invoke_Arity2(m_1495, cljs_core.CljsCoreList_EMPTY.X_conj_Arity2((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1)}, nil})))) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1} (conj m (list [:foo 1])))").(string)}, ``)}))
											}
											seq__907_1486, chunk__908_1487, count__909_1488, i__910_1489 = cljs_core.Next.Arity1IQ(seq__907_1493___1), nil, float64(0), float64(0)
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
				var seq__913_1497 interface{} = cljs_core.Seq.Arity1IQ((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{cljs_core.Array_map, cljs_core.Hash_map, cljs_core.Sorted_map}, nil}))
				var chunk__914_1498 interface{} = nil
				var count__915_1499 = float64(0)
				var i__916_1500 = float64(0)
				_, _, _, _ = seq__913_1497, chunk__914_1498, count__915_1499, i__916_1500
				for {
					if i__916_1500 < count__915_1499 {
						{
							var mt_1501 = cljs_core.Decorate_(chunk__914_1498).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(i__916_1500)
							_ = mt_1501
							if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
								var G__917 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
								var G__918 = float64(1)
								_, _ = G__917, G__918
								return mt_1501.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__917, G__918)
							}(), func(make_seq *cljs_core.AFn, seq__913_1497 interface{}, chunk__914_1498 interface{}, count__915_1499 float64, i__916_1500 float64, mt_1501 interface{}) *cljs_core.AFn {
								return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
									if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
										X__GT_t924 = func(__GT_t924 *cljs_core.AFn, seq__913_1497 interface{}, chunk__914_1498 interface{}, count__915_1499 float64, i__916_1500 float64, mt_1501 interface{}) *cljs_core.AFn {
											return cljs_core.Fn(__GT_t924, 9, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, i__916___1 interface{}, count__915___1 interface{}, chunk__914___1 interface{}, seq__913___1 interface{}, test_stuff___1 interface{}, meta925 interface{}) interface{} {
												return (&CljsCore_testT924{from_seq___1, make_seq___1, mt___1, i__916___1, count__915___1, chunk__914___1, seq__913___1, test_stuff___1, meta925})
											})
										}(&cljs_core.AFn{}, seq__913_1497, chunk__914_1498, count__915_1499, i__916_1500, mt_1501)

										return (&CljsCore_testT924{from_seq, make_seq, mt_1501, i__916_1500, count__915_1499, chunk__914_1498, seq__913_1497, test_stuff, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(5), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-column", Fqn: "end-column", X_hash: float64(1425389514)}), float64(71), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-line", Fqn: "end-line", X_hash: float64(1837326455)}), float64(2231), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "column", Fqn: "column", X_hash: float64(2078222095)}), float64(27), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "line", Fqn: "line", X_hash: float64(212345235)}), float64(2226), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "file", Fqn: "file", X_hash: float64(-1269645878)}), "/home/hraberg/go/src/github.com/hraberg/cljs2go/test/cljs/core_test.cljs"}, nil})})
									} else {
										return nil
									}
								})
							}(&cljs_core.AFn{}, seq__913_1497, chunk__914_1498, count__915_1499, i__916_1500, mt_1501).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
							} else {
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))").(string)}, ``)}))
							}
							seq__913_1497, chunk__914_1498, count__915_1499, i__916_1500 = seq__913_1497, chunk__914_1498, count__915_1499, (i__916_1500 + float64(1))
							continue
						}
					} else {
						{
							var temp__4388__auto___1502 = cljs_core.Seq.Arity1IQ(seq__913_1497)
							_ = temp__4388__auto___1502
							if cljs_core.Truth_(temp__4388__auto___1502) {
								{
									var seq__913_1503___1 = temp__4388__auto___1502
									_ = seq__913_1503___1
									if cljs_core.Chunked_seq_QMARK_.Arity1IB(seq__913_1503___1) {
										{
											var c__970__auto___1504 = cljs_core.Chunk_first.X_invoke_Arity1(seq__913_1503___1)
											_ = c__970__auto___1504
											seq__913_1497, chunk__914_1498, count__915_1499, i__916_1500 = cljs_core.Chunk_rest.X_invoke_Arity1(seq__913_1503___1), c__970__auto___1504, cljs_core.Count.X_invoke_Arity1(c__970__auto___1504).(float64), float64(0)
											continue
										}
									} else {
										{
											var mt_1505 = cljs_core.First.X_invoke_Arity1(seq__913_1503___1)
											_ = mt_1505
											if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil}), cljs_core.Conj.X_invoke_Arity2(func() interface{} {
												var G__929 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
												var G__930 = float64(1)
												_, _ = G__929, G__930
												return mt_1505.(cljs_core.CljsCoreIFn).X_invoke_Arity2(G__929, G__930)
											}(), func(make_seq *cljs_core.AFn, seq__913_1497 interface{}, chunk__914_1498 interface{}, count__915_1499 float64, i__916_1500 float64, mt_1505 interface{}, seq__913_1503___1 interface{}, temp__4388__auto___1502 cljs_core.CljsCoreISeq) *cljs_core.AFn {
												return cljs_core.Fn(make_seq, 1, func(from_seq interface{}) interface{} {
													if cljs_core.Truth_(cljs_core.Seq.Arity1IQ(from_seq)) {
														X__GT_t936 = func(__GT_t936 *cljs_core.AFn, seq__913_1497 interface{}, chunk__914_1498 interface{}, count__915_1499 float64, i__916_1500 float64, mt_1505 interface{}, seq__913_1503___1 interface{}, temp__4388__auto___1502 cljs_core.CljsCoreISeq) *cljs_core.AFn {
															return cljs_core.Fn(__GT_t936, 10, func(from_seq___1 interface{}, make_seq___1 interface{}, mt___1 interface{}, temp__4388__auto_____1 interface{}, i__916___1 interface{}, count__915___1 interface{}, chunk__914___1 interface{}, seq__913___2 interface{}, test_stuff___1 interface{}, meta937 interface{}) interface{} {
																return (&CljsCore_testT936{from_seq___1, make_seq___1, mt___1, temp__4388__auto_____1, i__916___1, count__915___1, chunk__914___1, seq__913___2, test_stuff___1, meta937})
															})
														}(&cljs_core.AFn{}, seq__913_1497, chunk__914_1498, count__915_1499, i__916_1500, mt_1505, seq__913_1503___1, temp__4388__auto___1502)

														return (&CljsCore_testT936{from_seq, make_seq, mt_1505, temp__4388__auto___1502, i__916_1500, count__915_1499, chunk__914_1498, seq__913_1503___1, test_stuff, (&cljs_core.CljsCorePersistentArrayMap{nil, float64(5), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-column", Fqn: "end-column", X_hash: float64(1425389514)}), float64(71), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "end-line", Fqn: "end-line", X_hash: float64(1837326455)}), float64(2231), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "column", Fqn: "column", X_hash: float64(2078222095)}), float64(27), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "line", Fqn: "line", X_hash: float64(212345235)}), float64(2226), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "file", Fqn: "file", X_hash: float64(-1269645878)}), "/home/hraberg/go/src/github.com/hraberg/cljs2go/test/cljs/core_test.cljs"}, nil})})
													} else {
														return nil
													}
												})
											}(&cljs_core.AFn{}, seq__913_1497, chunk__914_1498, count__915_1499, i__916_1500, mt_1505, seq__913_1503___1, temp__4388__auto___1502).X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), float64(3)}, nil})}, nil})))) {
											} else {
												panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo 1, :bar 2, :baz 3} (conj (mt :foo 1) ((fn make-seq [from-seq] (when (seq from-seq) (reify ISeqable (-seq [this] this) ISeq (-first [this] (first from-seq)) (-rest [this] (make-seq (rest from-seq)))))) [[:bar 2] [:baz 3]])))").(string)}, ``)}))
											}
											seq__913_1497, chunk__914_1498, count__915_1499, i__916_1500 = cljs_core.Next.Arity1IQ(seq__913_1503___1), nil, float64(0), float64(0)
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
				var _STAR_print_length_STAR_941 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_941
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_941

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(1)

						return strings.Join([]string{cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(0)}, nil})).(string)}, ``)
					}
				}()
			}(), "[1 ...]") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (binding [*print-length* 1] (str [1 2 3 4 5 6 7 8 9 0])) \"[1 ...]\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_942 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_942
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_942

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(2)

						return strings.Join([]string{cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(0)}, nil})).(string)}, ``)
					}
				}()
			}(), "[1 2 ...]") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (binding [*print-length* 2] (str [1 2 3 4 5 6 7 8 9 0])) \"[1 2 ...]\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_943 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_943
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_943

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(10)

						return strings.Join([]string{cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(10), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(0)}, nil})).(string)}, ``)
					}
				}()
			}(), "[1 2 3 4 5 6 7 8 9 0]") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (binding [*print-length* 10] (str [1 2 3 4 5 6 7 8 9 0])) \"[1 2 3 4 5 6 7 8 9 0]\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_944 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_944
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_944

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(10)

						return strings.Join([]string{cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar"}, nil})).(string)}, ``)
					}
				}()
			}(), "{:foo \"bar\"}") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (binding [*print-length* 10] (str {:foo \"bar\"})) \"{:foo \\\"bar\\\"}\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_945 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_945
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_945

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(1)

						return strings.Join([]string{cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), "woz"}, nil})).(string)}, ``)
					}
				}()
			}(), "{:foo \"bar\", ...}") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (binding [*print-length* 1] (str {:foo \"bar\", :baz \"woz\"})) \"{:foo \\\"bar\\\", ...}\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var _STAR_print_length_STAR_946 = cljs_core.X_STAR_print_length_STAR_
				_ = _STAR_print_length_STAR_946
				return func() string {
					defer func() {
						cljs_core.X_STAR_print_length_STAR_ = _STAR_print_length_STAR_946

					}()
					{
						cljs_core.X_STAR_print_length_STAR_ = float64(10)

						return strings.Join([]string{cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), "bar", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)}), "woz"}, nil})).(string)}, ``)
					}
				}()
			}(), "{:foo \"bar\", :baz \"woz\"}") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (binding [*print-length* 10] (str {:foo \"bar\", :baz \"woz\"})) \"{:foo \\\"bar\\\", :baz \\\"woz\\\"}\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var x = "a"
				_ = x
				{
					var G__947 = func() interface{} {
						if cljs_core.Value_(x).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
							return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(x), "Fqn")
						} else {
							return nil
						}
					}()
					_ = G__947
					switch G__947 {
					case "a":
						return float64(1)

					default:
						return "a"

					}
				}
			}(), "a") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (let [x \"a\"] (case x :a 1 \"a\")) \"a\")").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB("0atrue:key/wordsymb/olfalse[1 2 3 4]1234.56789", strings.Join([]string{cljs_core.Str.X_invoke_Arity1(float64(0)).(string), cljs_core.Str.X_invoke_Arity1("a").(string), cljs_core.Str.X_invoke_Arity1(true).(string), cljs_core.Str.X_invoke_Arity1(nil).(string), cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: "key", Name: "word", Fqn: "key/word", X_hash: float64(-420295340)})).(string), cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCoreSymbol{Ns: "symb", Name: "ol", Str: "symb/ol", X_hash: float64(-1700033577), X_meta: nil})).(string), cljs_core.Str.X_invoke_Arity1(false).(string), cljs_core.Str.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), float64(4)}, nil})).(string), cljs_core.Str.X_invoke_Arity1(1234.5678).(string), cljs_core.Str.X_invoke_Arity1(float64(9)).(string)}, ``)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= \"0atrue:key/wordsymb/olfalse[1 2 3 4]1234.56789\" (str 0 \"a\" true nil :key/word (quote symb/ol) false [1 2 3 4] 1234.5678 9))").(string)}, ``)}))
			}
			if cljs_core.Int_rotate_left.Arity2IIF(float64((cljs_core.Int32_(float64(2271560481))|cljs_core.Int32_(float64(0)))), float64(4)) == float64((cljs_core.Int32_(float64(1985229336)) | cljs_core.Int32_(float64(0)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (int-rotate-left (bit-or 2271560481 0) 4) (bit-or 1985229336 0))").(string)}, ``)}))
			}
			if cljs_core.Int_rotate_left.Arity2IIF(float64((cljs_core.Int32_(float64(2271560481))|cljs_core.Int32_(float64(0)))), float64(8)) == float64((cljs_core.Int32_(float64(1698898311)) | cljs_core.Int32_(float64(0)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (int-rotate-left (bit-or 2271560481 0) 8) (bit-or 1698898311 0))").(string)}, ``)}))
			}
			if cljs_core.Int_rotate_left.Arity2IIF(float64((cljs_core.Int32_(float64(2147483648))|cljs_core.Int32_(float64(0)))), float64(1)) == float64(1) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (int-rotate-left (bit-or 2147483648 0) 1) 1)").(string)}, ``)}))
			}
			if cljs_core.Int_rotate_left.Arity2IIF(float64((cljs_core.Int32_(float64(2014458966))|cljs_core.Int32_(float64(0)))), float64(4)) == float64((cljs_core.Int32_(float64(2166572391)) | cljs_core.Int32_(float64(0)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (int-rotate-left (bit-or 2014458966 0) 4) (bit-or 2166572391 0))").(string)}, ``)}))
			}
			if cljs_core.Int_rotate_left.Arity2IIF(float64((cljs_core.Int32_(float64(4294967295))|cljs_core.Int32_(float64(0)))), float64(4)) == float64((cljs_core.Int32_(float64(4294967295)) | cljs_core.Int32_(float64(0)))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (int-rotate-left (bit-or 4294967295 0) 4) (bit-or 4294967295 0))").(string)}, ``)}))
			}
			if cljs_core.Imul.Arity2IIF(float64(3), float64(3)) == float64(9) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (imul 3 3) 9)").(string)}, ``)}))
			}
			if cljs_core.Imul.Arity2IIF(float64(-1), float64(8)) == float64(-8) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (imul -1 8) -8)").(string)}, ``)}))
			}
			if cljs_core.Imul.Arity2IIF(float64(-2), float64(-2)) == float64(4) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (imul -2 -2) 4)").(string)}, ``)}))
			}
			if cljs_core.Imul.Arity2IIF(float64(4294967295), float64(5)) == float64(-5) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (imul 4294967295 5) -5)").(string)}, ``)}))
			}
			if cljs_core.Imul.Arity2IIF(float64(4294967294), float64(5)) == float64(-10) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (imul 4294967294 5) -10)").(string)}, ``)}))
			}
			Case_recur = func(case_recur *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(case_recur, 1, func(value interface{}) interface{} {
					for {
						{
							var G__949 = func() interface{} {
								if cljs_core.Value_(value).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
									return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(value), "Fqn")
								} else {
									return nil
								}
							}()
							_ = G__949
							switch G__949 {
							case "b":
								return float64(0)

							case "a":
								value = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
								continue

							default:
								panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("No matching clause: ").(string), cljs_core.Str.X_invoke_Arity1(value).(string)}, ``)}))

							}
						}
					}
				})
			}(&cljs_core.AFn{})

			if cljs_core.X_EQ_.Arity2IIB(Case_recur.X_invoke_Arity1((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})), float64(0)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (case-recur :a) 0)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(clojure_set.Rename_keys.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "two"}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "z", Fqn: "z", X_hash: float64(-789527183)})}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "z", Fqn: "z", X_hash: float64(-789527183)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "two"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (set/rename-keys {:a \"one\", :b \"two\"} {:a :z}) {:z \"one\", :b \"two\"})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(clojure_set.Rename_keys.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "two"}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "z", Fqn: "z", X_hash: float64(-789527183)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "y", Fqn: "y", X_hash: float64(-1757859776)})}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "z", Fqn: "z", X_hash: float64(-789527183)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "two"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (set/rename-keys {:a \"one\", :b \"two\"} {:a :z, :c :y}) {:z \"one\", :b \"two\"})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(clojure_set.Rename_keys.X_invoke_Arity2((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "two", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), "three"}, nil}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, nil})), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), "two", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), "one", (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), "three"}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (set/rename-keys {:a \"one\", :b \"two\", :c \"three\"} {:a :b, :b :a}) {:a \"two\", :b \"one\", :c \"three\"})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})}, nil}), cljs_core.Keys.X_invoke_Arity1(cljs_core.Apply.X_invoke_Arity2(cljs_core.Array_map, (&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2)}, nil})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [:foo] (keys (apply array-map [:foo 1 :foo 2])))").(string)}, ``)}))
			}
			{
				var not_strings_1508 = (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{true, false, nil, float64(1), func(G__1509 *cljs_core.AFn) *cljs_core.AFn {
					return cljs_core.Fn(G__1509, 0, func() interface{} {
						return nil
					})
				}(&cljs_core.AFn{})}, nil})
				_ = not_strings_1508
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1510 *cljs_core.AFn, not_strings_1508 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1510, 1, func(p1__81_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1511 interface{}) {
							defer func() {
								if e950 := recover(); e950 != nil {
									if cljs_core.Value_(e950).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e950
											_ = ___
											return__1511 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e950)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__81_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1508), not_strings_1508) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(every? (fn* [p1__81#] (= :failed (try (re-find #\".\" p1__81#) (catch js/TypeError _ :failed)))) not-strings)").(string)}, ``)}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1512 *cljs_core.AFn, not_strings_1508 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1512, 1, func(p1__82_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1513 interface{}) {
							defer func() {
								if e951 := recover(); e951 != nil {
									if cljs_core.Value_(e951).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e951
											_ = ___
											return__1513 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e951)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `.`, Flags: ``}), p1__82_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1508), not_strings_1508) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(every? (fn* [p1__82#] (= :failed (try (re-matches #\".\" p1__82#) (catch js/TypeError _ :failed)))) not-strings)").(string)}, ``)}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1514 *cljs_core.AFn, not_strings_1508 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1514, 1, func(p1__83_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1515 interface{}) {
							defer func() {
								if e952 := recover(); e952 != nil {
									if cljs_core.Value_(e952).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e952
											_ = ___
											return__1515 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e952)

									}
								}
							}()
							{
								return cljs_core.Re_find.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__83_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1508), not_strings_1508) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(every? (fn* [p1__83#] (= :failed (try (re-find #\"nomatch\" p1__83#) (catch js/TypeError _ :failed)))) not-strings)").(string)}, ``)}))
				}
				if cljs_core.Every_QMARK_.Arity2IIB(func(G__1516 *cljs_core.AFn, not_strings_1508 cljs_core.CljsCoreIVector) *cljs_core.AFn {
					return cljs_core.Fn(G__1516, 1, func(p1__84_SHARP_ interface{}) interface{} {
						return cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)}), func() (return__1517 interface{}) {
							defer func() {
								if e953 := recover(); e953 != nil {
									if cljs_core.Value_(e953).Type().AssignableTo(reflect.TypeOf((**js.TypeError)(nil)).Elem()) {
										{
											var ___ = e953
											_ = ___
											return__1517 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "failed", Fqn: "failed", X_hash: float64(-1397425762)})
										}
									} else {
										panic(e953)

									}
								}
							}()
							{
								return cljs_core.Re_matches.X_invoke_Arity2((&js.RegExp{Pattern: `nomatch`, Flags: ``}), p1__84_SHARP_)
							}
						}())
					})
				}(&cljs_core.AFn{}, not_strings_1508), not_strings_1508) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(every? (fn* [p1__84#] (= :failed (try (re-matches #\"nomatch\" p1__84#) (catch js/TypeError _ :failed)))) not-strings)").(string)}, ``)}))
				}
			}
			if cljs_core.Truth_((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}).Equiv((&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(.equiv :foo :foo)").(string)}, ``)}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil}).Equiv((&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(.equiv (quote foo) (quote foo))").(string)}, ``)}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}).Equiv((&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(.equiv {:foo 1, :bar 2} {:foo 1, :bar 2})").(string)}, ``)}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}).Equiv((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(.equiv [1 2 3] [1 2 3])").(string)}, ``)}))
			}
			if cljs_core.Truth_(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList).Equiv(cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(.equiv (quote (1 2 3)) (quote (1 2 3)))").(string)}, ``)}))
			}
			if cljs_core.Truth_(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq).Equiv(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreLazySeq))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(.equiv (map inc [1 2 3]) (map inc [1 2 3]))").(string)}, ``)}))
			}
			if cljs_core.Truth_((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "cat", Fqn: "cat", X_hash: float64(-1457810207)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bird", Fqn: "bird", X_hash: float64(-1252014845)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "dog", Fqn: "dog", X_hash: float64(-1650861974)}), nil}, nil}, nil}).Equiv((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "cat", Fqn: "cat", X_hash: float64(-1457810207)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bird", Fqn: "bird", X_hash: float64(-1252014845)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "dog", Fqn: "dog", X_hash: float64(-1650861974)}), nil}, nil}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(.equiv #{:cat :bird :dog} #{:cat :bird :dog})").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), []interface{}{float64(1), float64(2), float64(3)}), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (map inc) (array 1 2 3)) (quote (2 3 4)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Apply.X_invoke_Arity2(cljs_core.Str, cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(func(G__1518 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1518, 1, func(p1__85_SHARP_ interface{}) interface{} {
					return cljs_core.Native_invoke_instance_method.X_invoke_Arity3(p1__85_SHARP_, "ToUpperCase", []interface{}{})
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), "foo")), "FOO") {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (apply str (sequence (map (fn* [p1__85#] (.toUpperCase p1__85#))) \"foo\")) \"FOO\")").(string)}, ``)}))
			}
			if cljs_core.Hash.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(float64) == cljs_core.Hash.X_invoke_Arity1(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))).(float64) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(== (hash [1 2 3]) (hash (sequence (map inc) (range 3))))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= [1 2 3] (sequence (map inc) (range 3)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange)), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (map inc) (range 3)) [1 2 3])").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Remove.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(3), float64(5), float64(7), float64(9)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (remove even?) (range 10)) (quote (1 3 5 7 9)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Take.X_invoke_Arity1(float64(5)).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (take 5) (range 10)) (quote (0 1 2 3 4)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Take_while.X_invoke_Arity1(func(G__1519 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1519, 1, func(p1__86_SHARP_ interface{}) interface{} {
					return (p1__86_SHARP_.(float64) < float64(5))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(1), float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (take-while (fn* [p1__86#] (< p1__86# 5))) (range 10)) (quote (0 1 2 3 4)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Drop.X_invoke_Arity1(float64(5)).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (drop 5) (range 10)) (quote (5 6 7 8 9)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Drop_while.X_invoke_Arity1(func(G__1520 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1520, 1, func(p1__87_SHARP_ interface{}) interface{} {
					return (p1__87_SHARP_.(float64) < float64(5))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(5), float64(6), float64(7), float64(8), float64(9)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (drop-while (fn* [p1__87#] (< p1__87# 5))) (range 10)) (quote (5 6 7 8 9)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Take_nth.X_invoke_Arity1(float64(2)).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(2), float64(4), float64(6), float64(8)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (take-nth 2) (range 10)) (quote (0 2 4 6 8)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Replace.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)})}, nil})).(cljs_core.CljsCoreIFn), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), float64(2)})).(*cljs_core.CljsCoreList)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(1), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), float64(2)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (replace {:foo :bar}) (quote (:foo 1 :foo 2))) (quote (:bar 1 :bar 2)))").(string)}, ``)}))
			}
			{
				var ret_1521 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))
				_ = ret_1521
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_1521)) && (cljs_core.X_EQ_.Arity2IIB(ret_1521, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(and (vector? ret) (= ret (quote (1 2 3))))").(string)}, ``)}))
				}
			}
			{
				var ret_1522 = cljs_core.Into.X_invoke_Arity3(cljs_core.CljsCorePersistentVector_EMPTY, cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange))
				_ = ret_1522
				if (cljs_core.Vector_QMARK_.Arity1IB(ret_1522)) && (cljs_core.X_EQ_.Arity2IIB(ret_1522, cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(0), float64(2), float64(4), float64(6), float64(8)})).(*cljs_core.CljsCoreList))) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(and (vector? ret) (= ret (quote (0 2 4 6 8))))").(string)}, ``)}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Map_.X_invoke_Arity2(cljs_core.Inc, cljs_core.Sequence.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Range_.X_invoke_Arity1(float64(3)).(*cljs_core.CljsCoreRange))).(*cljs_core.CljsCoreLazySeq), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (map inc (sequence (map inc) (range 3))) (quote (2 3 4)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Dedupe.X_invoke_Arity0().(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(6), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1), float64(2), float64(2), float64(3), float64(3)}, nil})), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (dedupe) [1 1 2 2 3 3]) (quote (1 2 3)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Mapcat.X_invoke_ArityVariadic(cljs_core.Reverse, cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(6), float64(5), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(9), float64(8), float64(7)}, nil})}, nil})})), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (mapcat reverse [[3 2 1 0] [6 5 4] [9 8 7]]) (range 10))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Reverse).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(4), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(2), float64(1), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(6), float64(5), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(9), float64(8), float64(7)}, nil})}, nil})), cljs_core.Range_.X_invoke_Arity1(float64(10)).(*cljs_core.CljsCoreRange)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (mapcat reverse) [[3 2 1 0] [6 5 4] [9 8 7]]) (range 10))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(cljs_core.Eduction.X_invoke_Arity2(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil})).(*cljs_core.CljsCoreEduction)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(2), float64(3), float64(4)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq (eduction (map inc) [1 2 3])) (quote (2 3 4)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Partition_by.X_invoke_Arity1((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "split", Fqn: "split", X_hash: float64(-599435118)}), nil}, nil}, nil})).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(7), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "split", Fqn: "split", X_hash: float64(-599435118)}), float64(4), float64(5), float64(6)}, nil})), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "split", Fqn: "split", X_hash: float64(-599435118)})}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5), float64(6)}, nil})})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (partition-by #{:split}) [1 2 3 :split 4 5 6]) (quote ([1 2 3] [:split] [4 5 6])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Partition_all.X_invoke_Arity1(float64(3)).(cljs_core.CljsCoreIFn), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3), float64(4), float64(5)})).(*cljs_core.CljsCoreList)), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(5)}, nil})})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (partition-all 3) (quote (1 2 3 4 5))) (quote ([1 2 3] [4 5])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(cljs_core.Keep.X_invoke_Arity1(cljs_core.Identity).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), nil, float64(2), nil, float64(3)}, nil})), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(1), float64(2), float64(3)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence (keep identity) [1 nil 2 nil 3]) (quote (1 2 3)))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Keep_indexed.X_invoke_Arity2(cljs_core.Identity, (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil})), cljs_core.Sequence.X_invoke_Arity2(cljs_core.Keep_indexed.X_invoke_Arity1(cljs_core.Identity).(cljs_core.CljsCoreIFn), (&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "bar", Fqn: "bar", X_hash: float64(-1386246584)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "baz", Fqn: "baz", X_hash: float64(-1134894686)})}, nil}))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (keep-indexed identity [:foo nil :bar nil :baz]) (sequence (keep-indexed identity) [:foo nil :bar nil :baz]))").(string)}, ``)}))
			}
			Xform = cljs_core.Comp.X_invoke_ArityVariadic(cljs_core.Map_.X_invoke_Arity1(cljs_core.Inc).(cljs_core.CljsCoreIFn), cljs_core.Filter.X_invoke_Arity1(cljs_core.Even_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Dedupe.X_invoke_Arity0().(cljs_core.CljsCoreIFn), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Range_).(cljs_core.CljsCoreIFn), cljs_core.Partition_all.X_invoke_Arity1(float64(3)).(cljs_core.CljsCoreIFn), cljs_core.Partition_by.X_invoke_Arity1(func(G__1523 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1523, 1, func(p1__88_SHARP_ interface{}) interface{} {
					return (cljs_core.Apply.X_invoke_Arity2(cljs_core.X_PLUS_, p1__88_SHARP_).(float64) < float64(7))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Mapcat.X_invoke_Arity1(cljs_core.Flatten).(cljs_core.CljsCoreIFn), cljs_core.Random_sample.X_invoke_Arity1(1.0).(cljs_core.CljsCoreIFn), cljs_core.Take_nth.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Keep.X_invoke_Arity1(func(G__1524 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1524, 1, func(p1__89_SHARP_ interface{}) interface{} {
					if cljs_core.Odd_QMARK_.Arity1IB(p1__89_SHARP_) {
						return (p1__89_SHARP_.(float64) * p1__89_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Keep_indexed.X_invoke_Arity1(func(G__1525 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1525, 2, func(p1__90_SHARP_ interface{}, p2__91_SHARP_ interface{}) interface{} {
					if cljs_core.Even_QMARK_.Arity1IB(p1__90_SHARP_) {
						return (p1__90_SHARP_.(float64) * p2__91_SHARP_.(float64))
					} else {
						return nil
					}
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Replace.X_invoke_Arity1((&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{float64(2), "two", float64(6), "six", float64(18), "eighteen"}, nil})).(cljs_core.CljsCoreIFn), cljs_core.Take.X_invoke_Arity1(float64(11)).(cljs_core.CljsCoreIFn), cljs_core.Take_while.X_invoke_Arity1(func(G__1526 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1526, 1, func(p1__92_SHARP_ interface{}) interface{} {
					return cljs_core.Not_EQ_.Arity2IIB(float64(300), p1__92_SHARP_)
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn), cljs_core.Drop.X_invoke_Arity1(float64(1)).(cljs_core.CljsCoreIFn), cljs_core.Drop_while.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn), cljs_core.Remove.X_invoke_Arity1(cljs_core.String_QMARK_).(cljs_core.CljsCoreIFn)})).(cljs_core.CljsCoreIFn).(*cljs_core.AFn)

			Data = cljs_core.Vec.X_invoke_Arity1(cljs_core.Interleave.X_invoke_Arity2(cljs_core.Range_.X_invoke_Arity1(float64(18)).(*cljs_core.CljsCoreRange), cljs_core.Range_.X_invoke_Arity1(float64(20)).(*cljs_core.CljsCoreRange)).(*cljs_core.CljsCoreLazySeq))

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_Arity2(Xform, Data), cljs_core.List.X_invoke_ArityVariadic(cljs_core.Array_seq.X_invoke_Arity1([]interface{}{float64(36), float64(200), float64(10)})).(*cljs_core.CljsCoreList)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence xform data) (quote (36 200 10)))").(string)}, ``)}))
			}
			Xf = cljs_core.Map_.X_invoke_Arity1(func(G__1527 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1527, 2, func(p1__93_SHARP_ interface{}, p2__94_SHARP_ interface{}) interface{} {
					return (p1__93_SHARP_.(float64) + p2__94_SHARP_.(float64))
				})
			}(&cljs_core.AFn{})).(cljs_core.CljsCoreIFn).(*cljs_core.AFn)

			if cljs_core.X_EQ_.Arity2IIB(cljs_core.Sequence.X_invoke_ArityVariadic(Xf, (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0)}, nil}), cljs_core.Array_seq.X_invoke_Arity1([]interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})})), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (sequence xf [0 0] [1 2]) [1 2])").(string)}, ``)}))
			}
			{
				var xs_1528 = (&cljs_core.CljsCorePersistentVector{nil, float64(21), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(44), float64(43), float64(42), float64(41), float64(40), float64(39), float64(38), float64(37), float64(36), float64(35), float64(34), float64(33), float64(32), float64(31), float64(30), float64(29), float64(28), float64(27), float64(26), float64(25), float64(24)}, nil})
				_ = xs_1528
				{
					var m_1529 interface{} = cljs_core.Transient.X_invoke_Arity1(cljs_core.Zipmap.X_invoke_Arity2(xs_1528, cljs_core.Repeat.X_invoke_Arity1(float64(1)).(*cljs_core.CljsCoreLazySeq)))
					var xs_1530___1 interface{} = xs_1528
					_, _ = m_1529, xs_1530___1
					for {
						{
							var temp__4386__auto___1531 = cljs_core.First.X_invoke_Arity1(xs_1530___1)
							_ = temp__4386__auto___1531
							if cljs_core.Truth_(temp__4386__auto___1531) {
								{
									var x_1532 = temp__4386__auto___1531
									_ = x_1532
									if cljs_core.Contains_QMARK_.Arity2IIB(m_1529, x_1532) {
										m_1529, xs_1530___1 = cljs_core.Dissoc_BANG_.X_invoke_Arity2(m_1529, x_1532), cljs_core.Next.Arity1IQ(xs_1530___1)
										continue
									} else {
										panic(cljs_core.Ex_info.X_invoke_Arity2("CLJS-849 regression!", (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "m", Fqn: "m", X_hash: float64(1632677161)}), cljs_core.Persistent_BANG_.X_invoke_Arity1(m_1529), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "xs", Fqn: "xs", X_hash: float64(649443341)}), xs_1530___1}, nil})).(*cljs_core.CljsCoreExceptionInfo))
									}
								}
							} else {
							}
						}
						break
					}
				}
			}
			{
				var m_1533 = cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10), float64(11), float64(12), float64(13), float64(14), float64(15)}, []interface{}{float64(0), float64(1), float64(2), float64(3), float64(4), float64(5), float64(6), float64(7), float64(8), float64(9), float64(10), float64(11), float64(12), float64(13), float64(14), float64(15)})
				_ = m_1533
				if cljs_core.X_EQ_.Arity2IIB(cljs_core.Seq.Arity1IQ(m_1533), (&cljs_core.CljsCorePersistentVector{nil, float64(16), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{(&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(0), float64(0)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(1)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(2), float64(2)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(3), float64(3)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(4), float64(4)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(5), float64(5)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(6), float64(6)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(7), float64(7)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(8), float64(8)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(9), float64(9)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(10), float64(10)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(11), float64(11)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(12), float64(12)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(13), float64(13)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(14), float64(14)}, nil}), (&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(15), float64(15)}, nil})}, nil})) {
				} else {
					panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (seq m) [[0 0] [1 1] [2 2] [3 3] [4 4] [5 5] [6 6] [7 7] [8 8] [9 9] [10 10] [11 11] [12 12] [13 13] [14 14] [15 15]])").(string)}, ``)}))
				}
			}
			Test_map = cljs_core.CljsCorePersistentHashMap_FromArrays.X_invoke_Arity2([]interface{}{(&js.Date{Millis: 1386306000000}), (&js.Date{Millis: 1387429200000}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "e", Fqn: "e", X_hash: float64(1381269198)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "g", Fqn: "g", X_hash: float64(1738089905)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "j", Fqn: "j", X_hash: float64(-1397974765)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "h", Fqn: "h", X_hash: float64(1109658740)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "d", Fqn: "d", X_hash: float64(1972142424)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "i", Fqn: "i", X_hash: float64(-1386841315)}), (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})}, []interface{}{float64(6), float64(3), float64(5), float64(7), float64(10), float64(8), float64(2), float64(4), float64(9), float64(1)}).(*cljs_core.CljsCorePersistentHashMap)

			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var G__954 = (&js.Date{Millis: 1387429200000})
				_ = G__954
				return Test_map.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__954)
			}(), float64(3)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (test-map #inst \"2013-12-19T05:00:00.000-00:00\") 3)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(func() interface{} {
				var G__955 = (&js.Date{Millis: 1386306000000})
				_ = G__955
				return Test_map.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__955)
			}(), float64(6)) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= (test-map #inst \"2013-12-06T05:00:00.000-00:00\") 6)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), true}, nil}), cljs_core.Meta.X_invoke_Arity1(cljs_core.With_meta.X_invoke_Arity2(func(G__1534 *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(G__1534, 0, func() interface{} {
					return nil
				})
			}(&cljs_core.AFn{}), (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)}), true}, nil})))) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= {:foo true} (meta (fn [])))").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(float64(-1), -1.0) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= -1 -1N)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(9.007199254740996E15, 9.007199254740996E15) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 9.007199254740996E15 9007199254740995N)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(1.5, 1.5) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 1.5 1.5M)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB(4.9E-324, 4.9E-324) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= 4.9E-324 5E-324M)").(string)}, ``)}))
			}
			Var_test = func(var_test *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(var_test, 2, func(a interface{}, b interface{}) interface{} {
					return (a.(float64) + b.(float64))
				})
			}(&cljs_core.AFn{})

			Var_test_test = func(var_test_test *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(var_test_test, 2, func(a interface{}, b interface{}) interface{} {
					return (a.(float64) + b.(float64))
				})
			}(&cljs_core.AFn{})

			Var_test_self_call = func(var_test_self_call *cljs_core.AFn) *cljs_core.AFn {
				return cljs_core.Fn(var_test_self_call, 2, func(a interface{}, b interface{}) interface{} {
					return (a.(float64) + b.(float64))
				})
			}(&cljs_core.AFn{})

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

var No_dispatch_value *cljs_core.CljsCoreMultiFn

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

func (this__766__auto__ *CljsCore_testPerson) X_lookup_Arity3(k672 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__675 = func() interface{} {
			if cljs_core.Value_(k672).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k672), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__675
		switch G__675 {
		case "lastname":
			return this__766__auto__.Lastname

		case "firstname":
			return this__766__auto__.Firstname

		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k672, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIPrintWithWriter__() {}
func (this__780__auto__ *CljsCore_testPerson) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__1536 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1536, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
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
		var h__582__auto__ = this__759__auto__.X__hash
		_ = h__582__auto__
		if !(cljs_core.Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__582__auto_____1
				this__759__auto__.X__hash = h__582__auto_____1

				return h__582__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testPerson) CljsCoreIEquiv__() {}
func (this__760__auto__ *CljsCore_testPerson) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__761__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
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
func (_ *CljsCore_testPerson) CljsCoreIMap__()    {}
func (this__775__auto__ *CljsCore_testPerson) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)}), nil}, nil}, nil}), k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testPerson{this__775__auto__.Firstname, this__775__auto__.Lastname, this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testPerson) CljsCoreIAssociative__() {}
func (this__771__auto__ *CljsCore_testPerson) X_assoc_Arity3(k__772__auto__ interface{}, G__671 interface{}) interface{} {
	{
		var pred__683 = cljs_core.Keyword_identical_QMARK_
		var expr__684 = k__772__auto__
		_, _ = pred__683, expr__684
		if cljs_core.Truth_(func() interface{} {
			var G__686 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "firstname", Fqn: "firstname", X_hash: float64(1659984849)})
			var G__687 = expr__684
			_, _ = G__686, G__687
			return pred__683.X_invoke_Arity2(G__686, G__687)
		}()) {
			return (&CljsCore_testPerson{G__671, this__771__auto__.Lastname, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__688 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "lastname", Fqn: "lastname", X_hash: float64(-265181465)})
				var G__689 = expr__684
				_, _ = G__688, G__689
				return pred__683.X_invoke_Arity2(G__688, G__689)
			}()) {
				return (&CljsCore_testPerson{this__771__auto__.Firstname, G__671, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPerson{this__771__auto__.Firstname, this__771__auto__.Lastname, this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__671), nil})
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
func (this__763__auto__ *CljsCore_testPerson) X_with_meta_Arity2(G__671 interface{}) interface{} {
	return (&CljsCore_testPerson{this__763__auto__.Firstname, this__763__auto__.Lastname, G__671, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testPerson) CljsCoreICollection__() {}
func (this__769__auto__ *CljsCore_testPerson) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
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

func (this__766__auto__ *CljsCore_testA) X_lookup_Arity3(k691 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__694 = k691
		_ = G__694
		switch G__694 {
		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k691, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testA) CljsCoreIPrintWithWriter__() {}
func (this__780__auto__ *CljsCore_testA) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__1538 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1538, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
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
		var h__582__auto__ = this__759__auto__.X__hash
		_ = h__582__auto__
		if !(cljs_core.Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__582__auto_____1
				this__759__auto__.X__hash = h__582__auto_____1

				return h__582__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testA) CljsCoreIEquiv__() {}
func (this__760__auto__ *CljsCore_testA) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__761__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
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
func (_ *CljsCore_testA) CljsCoreIMap__()    {}
func (this__775__auto__ *CljsCore_testA) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB(cljs_core.CljsCorePersistentHashSet_EMPTY, k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testA{this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testA) CljsCoreIAssociative__() {}
func (this__771__auto__ *CljsCore_testA) X_assoc_Arity3(k__772__auto__ interface{}, G__690 interface{}) interface{} {
	{
		var pred__698 = cljs_core.Keyword_identical_QMARK_
		var expr__699 = k__772__auto__
		_, _ = pred__698, expr__699
		return (&CljsCore_testA{this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__690), nil})
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
func (this__763__auto__ *CljsCore_testA) X_with_meta_Arity2(G__690 interface{}) interface{} {
	return (&CljsCore_testA{G__690, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testA) CljsCoreICollection__() {}
func (this__769__auto__ *CljsCore_testA) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
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

func (this__766__auto__ *CljsCore_testC) X_lookup_Arity3(k702 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__705 = func() interface{} {
			if cljs_core.Value_(k702).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k702), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__705
		switch G__705 {
		case "c":
			return this__766__auto__.C

		case "b":
			return this__766__auto__.B

		case "a":
			return this__766__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k702, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testC) CljsCoreIPrintWithWriter__() {}
func (this__780__auto__ *CljsCore_testC) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__1540 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1540, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
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
		var h__582__auto__ = this__759__auto__.X__hash
		_ = h__582__auto__
		if !(cljs_core.Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__582__auto_____1
				this__759__auto__.X__hash = h__582__auto_____1

				return h__582__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testC) CljsCoreIEquiv__() {}
func (this__760__auto__ *CljsCore_testC) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__761__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
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
func (_ *CljsCore_testC) CljsCoreIMap__()    {}
func (this__775__auto__ *CljsCore_testC) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testC{this__775__auto__.A, this__775__auto__.B, this__775__auto__.C, this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testC) CljsCoreIAssociative__() {}
func (this__771__auto__ *CljsCore_testC) X_assoc_Arity3(k__772__auto__ interface{}, G__701 interface{}) interface{} {
	{
		var pred__715 = cljs_core.Keyword_identical_QMARK_
		var expr__716 = k__772__auto__
		_, _ = pred__715, expr__716
		if cljs_core.Truth_(func() interface{} {
			var G__718 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__719 = expr__716
			_, _ = G__718, G__719
			return pred__715.X_invoke_Arity2(G__718, G__719)
		}()) {
			return (&CljsCore_testC{G__701, this__771__auto__.B, this__771__auto__.C, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__720 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__721 = expr__716
				_, _ = G__720, G__721
				return pred__715.X_invoke_Arity2(G__720, G__721)
			}()) {
				return (&CljsCore_testC{this__771__auto__.A, G__701, this__771__auto__.C, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
			} else {
				if cljs_core.Truth_(func() interface{} {
					var G__722 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "c", Fqn: "c", X_hash: float64(-1763192079)})
					var G__723 = expr__716
					_, _ = G__722, G__723
					return pred__715.X_invoke_Arity2(G__722, G__723)
				}()) {
					return (&CljsCore_testC{this__771__auto__.A, this__771__auto__.B, G__701, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
				} else {
					return (&CljsCore_testC{this__771__auto__.A, this__771__auto__.B, this__771__auto__.C, this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__701), nil})
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
func (this__763__auto__ *CljsCore_testC) X_with_meta_Arity2(G__701 interface{}) interface{} {
	return (&CljsCore_testC{this__763__auto__.A, this__763__auto__.B, this__763__auto__.C, G__701, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testC) CljsCoreICollection__() {}
func (this__769__auto__ *CljsCore_testC) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
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

func (this__766__auto__ *CljsCore_testA2) X_lookup_Arity3(k727 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__730 = func() interface{} {
			if cljs_core.Value_(k727).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k727), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__730
		switch G__730 {
		case "x":
			return this__766__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k727, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIPrintWithWriter__() {}
func (this__780__auto__ *CljsCore_testA2) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__1542 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1542, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
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
		var h__582__auto__ = this__759__auto__.X__hash
		_ = h__582__auto__
		if !(cljs_core.Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__582__auto_____1
				this__759__auto__.X__hash = h__582__auto_____1

				return h__582__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testA2) CljsCoreIEquiv__() {}
func (this__760__auto__ *CljsCore_testA2) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__761__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
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
func (_ *CljsCore_testA2) CljsCoreIMap__()    {}
func (this__775__auto__ *CljsCore_testA2) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), nil}, nil}, nil}), k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testA2{this__775__auto__.X, this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testA2) CljsCoreIAssociative__() {}
func (this__771__auto__ *CljsCore_testA2) X_assoc_Arity3(k__772__auto__ interface{}, G__726 interface{}) interface{} {
	{
		var pred__736 = cljs_core.Keyword_identical_QMARK_
		var expr__737 = k__772__auto__
		_, _ = pred__736, expr__737
		if cljs_core.Truth_(func() interface{} {
			var G__739 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__740 = expr__737
			_, _ = G__739, G__740
			return pred__736.X_invoke_Arity2(G__739, G__740)
		}()) {
			return (&CljsCore_testA2{G__726, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testA2{this__771__auto__.X, this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__726), nil})
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
func (this__763__auto__ *CljsCore_testA2) X_with_meta_Arity2(G__726 interface{}) interface{} {
	return (&CljsCore_testA2{this__763__auto__.X, G__726, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testA2) CljsCoreICollection__() {}
func (this__769__auto__ *CljsCore_testA2) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
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

func (this__766__auto__ *CljsCore_testB) X_lookup_Arity3(k742 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__745 = func() interface{} {
			if cljs_core.Value_(k742).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k742), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__745
		switch G__745 {
		case "x":
			return this__766__auto__.X

		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k742, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testB) CljsCoreIPrintWithWriter__() {}
func (this__780__auto__ *CljsCore_testB) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__1544 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1544, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
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
		var h__582__auto__ = this__759__auto__.X__hash
		_ = h__582__auto__
		if !(cljs_core.Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__582__auto_____1
				this__759__auto__.X__hash = h__582__auto_____1

				return h__582__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testB) CljsCoreIEquiv__() {}
func (this__760__auto__ *CljsCore_testB) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__761__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
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
func (_ *CljsCore_testB) CljsCoreIMap__()    {}
func (this__775__auto__ *CljsCore_testB) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)}), nil}, nil}, nil}), k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testB{this__775__auto__.X, this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testB) CljsCoreIAssociative__() {}
func (this__771__auto__ *CljsCore_testB) X_assoc_Arity3(k__772__auto__ interface{}, G__741 interface{}) interface{} {
	{
		var pred__751 = cljs_core.Keyword_identical_QMARK_
		var expr__752 = k__772__auto__
		_, _ = pred__751, expr__752
		if cljs_core.Truth_(func() interface{} {
			var G__754 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "x", Fqn: "x", X_hash: float64(2099068185)})
			var G__755 = expr__752
			_, _ = G__754, G__755
			return pred__751.X_invoke_Arity2(G__754, G__755)
		}()) {
			return (&CljsCore_testB{G__741, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
		} else {
			return (&CljsCore_testB{this__771__auto__.X, this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__741), nil})
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
func (this__763__auto__ *CljsCore_testB) X_with_meta_Arity2(G__741 interface{}) interface{} {
	return (&CljsCore_testB{this__763__auto__.X, G__741, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testB) CljsCoreICollection__() {}
func (this__769__auto__ *CljsCore_testB) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
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

type CljsCore_testT756 struct {
	Test_stuff interface{}
	Meta757    interface{}
}

func (_ *CljsCore_testT756) CljsCore_testIFoo__() {}
func (this *CljsCore_testT756) Foo_Arity1() interface{} {
	return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "foo", Fqn: "foo", X_hash: float64(1268894036)})
}

func (_ *CljsCore_testT756) CljsCoreIMeta__() {}
func (_758 *CljsCore_testT756) X_meta_Arity1() interface{} {
	return _758.Meta757
}

func (_ *CljsCore_testT756) CljsCoreIWithMeta__() {}
func (_758 *CljsCore_testT756) X_with_meta_Arity2(meta757___1 interface{}) interface{} {
	return (&CljsCore_testT756{_758.Test_stuff, meta757___1})
}

var X__GT_t756 *cljs_core.AFn

type CljsCore_testIMutate interface {
	CljsCore_testIMutate__()
	Mutate_Arity1() interface{}
}

var Mutate *cljs_core.AFn

type CljsCore_testMutate struct{ A interface{} }

func (_ *CljsCore_testMutate) CljsCore_testIMutate__() {}
func (___ *CljsCore_testMutate) Mutate_Arity1() interface{} {
	return func() interface{} {
		var return__1545 = (&cljs_core.CljsCoreSymbol{Ns: nil, Name: "foo", Str: "foo", X_hash: float64(-1385541733), X_meta: nil})
		___.A = return__1545
		return return__1545
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
func (___ *CljsCore_testFirst) X_find_first_Arity2(p__759 interface{}) interface{} {
	{
		var vec__761 = p__759
		var x = cljs_core.Nth.X_invoke_Arity3(vec__761, float64(0), nil)
		_, _ = vec__761, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCore_testIHasFirst__() {}
func (p__762 *CljsCore_testFirst) X_get_first_Arity1() interface{} {
	{
		var vec__764 = p__762
		var x = cljs_core.Nth.X_invoke_Arity3(vec__764, float64(0), nil)
		_, _ = vec__764, x
		return x
	}
}

func (_ *CljsCore_testFirst) CljsCoreObject__() {}
func (p__765 *CljsCore_testFirst) ToString() string {
	{
		var vec__767 = p__765
		var x = cljs_core.Nth.X_invoke_Arity3(vec__767, float64(0), nil)
		_, _ = vec__767, x
		return strings.Join([]string{cljs_core.Str.X_invoke_Arity1(x).(string)}, ``)
	}
}

func (this *CljsCore_testFirst) String() string {
	return this.ToString()
}

func (_ *CljsCore_testFirst) CljsCoreIFn__() {}
func (p__768 *CljsCore_testFirst) X_invoke_Arity0() interface{} {
	{
		var vec__770 = p__768
		var x = cljs_core.Nth.X_invoke_Arity3(vec__770, float64(0), nil)
		_, _ = vec__770, x
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
func (___ *CljsCore_testDestructuringWithLocals) X_find_first_Arity2(p__772 interface{}) interface{} {
	{
		var vec__774 = p__772
		var x = cljs_core.Nth.X_invoke_Arity3(vec__774, float64(0), nil)
		var y = cljs_core.Nth.X_invoke_Arity3(vec__774, float64(1), nil)
		_, _, _ = vec__774, x, y
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

func (this__766__auto__ *CljsCore_testPrintMe) X_lookup_Arity3(k786 interface{}, else__767__auto__ interface{}) interface{} {
	{
		var G__789 = func() interface{} {
			if cljs_core.Value_(k786).Type().AssignableTo(reflect.TypeOf((**cljs_core.CljsCoreKeyword)(nil)).Elem()) {
				return cljs_core.Native_get_instance_field.X_invoke_Arity2(cljs_core.Keyword.X_invoke_Arity1(k786), "Fqn")
			} else {
				return nil
			}
		}()
		_ = G__789
		switch G__789 {
		case "b":
			return this__766__auto__.B

		case "a":
			return this__766__auto__.A

		default:
			return cljs_core.Get.X_invoke_Arity3(this__766__auto__.X__extmap, k786, else__767__auto__)

		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIPrintWithWriter__() {}
func (this__780__auto__ *CljsCore_testPrintMe) X_pr_writer_Arity3(writer__781__auto__ interface{}, opts__782__auto__ interface{}) interface{} {
	{
		var pr_pair__783__auto__ = func(G__1547 *cljs_core.AFn) *cljs_core.AFn {
			return cljs_core.Fn(G__1547, 3, func(keyval__784__auto__ interface{}, ___785__auto__ interface{}, ___785__auto_____1 interface{}) interface{} {
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
		var h__582__auto__ = this__759__auto__.X__hash
		_ = h__582__auto__
		if !(cljs_core.Nil_(h__582__auto__)) {
			return h__582__auto__
		} else {
			{
				var h__582__auto_____1 = cljs_core.Hash_imap.X_invoke_Arity1(this__759__auto__).(float64)
				_ = h__582__auto_____1
				this__759__auto__.X__hash = h__582__auto_____1

				return h__582__auto_____1
			}
		}
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIEquiv__() {}
func (this__760__auto__ *CljsCore_testPrintMe) X_equiv_Arity2(other__761__auto__ interface{}) bool {
	if cljs_core.Truth_(func() interface{} {
		var and__159__auto__ = other__761__auto__
		_ = and__159__auto__
		if cljs_core.Truth_(and__159__auto__) {
			return (reflect.DeepEqual(cljs_core.Type_.X_invoke_Arity1(this__760__auto__), cljs_core.Type_.X_invoke_Arity1(other__761__auto__))) && (cljs_core.Truth_(cljs_core.Equiv_map.X_invoke_Arity2(this__760__auto__, other__761__auto__)))
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
func (_ *CljsCore_testPrintMe) CljsCoreIMap__()    {}
func (this__775__auto__ *CljsCore_testPrintMe) X_dissoc_Arity2(k__776__auto__ interface{}) interface{} {
	if cljs_core.Contains_QMARK_.Arity2IIB((&cljs_core.CljsCorePersistentHashSet{nil, &cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{(&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)}), nil, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)}), nil}, nil}, nil}), k__776__auto__) {
		return cljs_core.Dissoc.X_invoke_Arity2(cljs_core.With_meta.X_invoke_Arity2(cljs_core.Into.X_invoke_Arity2(cljs_core.CljsCorePersistentArrayMap_EMPTY, this__775__auto__), this__775__auto__.X__meta), k__776__auto__)
	} else {
		return (&CljsCore_testPrintMe{this__775__auto__.A, this__775__auto__.B, this__775__auto__.X__meta, cljs_core.Not_empty.X_invoke_Arity1(cljs_core.Dissoc.X_invoke_Arity2(this__775__auto__.X__extmap, k__776__auto__)), nil})
	}
}

func (_ *CljsCore_testPrintMe) CljsCoreIAssociative__() {}
func (this__771__auto__ *CljsCore_testPrintMe) X_assoc_Arity3(k__772__auto__ interface{}, G__785 interface{}) interface{} {
	{
		var pred__797 = cljs_core.Keyword_identical_QMARK_
		var expr__798 = k__772__auto__
		_, _ = pred__797, expr__798
		if cljs_core.Truth_(func() interface{} {
			var G__800 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})
			var G__801 = expr__798
			_, _ = G__800, G__801
			return pred__797.X_invoke_Arity2(G__800, G__801)
		}()) {
			return (&CljsCore_testPrintMe{G__785, this__771__auto__.B, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
		} else {
			if cljs_core.Truth_(func() interface{} {
				var G__802 = (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "b", Fqn: "b", X_hash: float64(1482224470)})
				var G__803 = expr__798
				_, _ = G__802, G__803
				return pred__797.X_invoke_Arity2(G__802, G__803)
			}()) {
				return (&CljsCore_testPrintMe{this__771__auto__.A, G__785, this__771__auto__.X__meta, this__771__auto__.X__extmap, nil})
			} else {
				return (&CljsCore_testPrintMe{this__771__auto__.A, this__771__auto__.B, this__771__auto__.X__meta, cljs_core.Assoc.X_invoke_Arity3(this__771__auto__.X__extmap, k__772__auto__, G__785), nil})
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
func (this__763__auto__ *CljsCore_testPrintMe) X_with_meta_Arity2(G__785 interface{}) interface{} {
	return (&CljsCore_testPrintMe{this__763__auto__.A, this__763__auto__.B, G__785, this__763__auto__.X__extmap, this__763__auto__.X__hash})
}

func (_ *CljsCore_testPrintMe) CljsCoreICollection__() {}
func (this__769__auto__ *CljsCore_testPrintMe) X_conj_Arity2(entry__770__auto__ interface{}) interface{} {
	if cljs_core.Vector_QMARK_.Arity1IB(entry__770__auto__) {
		return this__769__auto__.X_assoc_Arity3(cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(0)), cljs_core.Decorate_(entry__770__auto__).(cljs_core.CljsCoreIIndexed).X_nth_Arity2(float64(1)))
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

type CljsCore_testT837 struct {
	F          interface{}
	Baz        interface{}
	Test_stuff interface{}
	Meta838    interface{}
}

func (_ *CljsCore_testT837) CljsCore_testIBar__() {}
func (___ *CljsCore_testT837) X_bar_Arity2(x interface{}) interface{} {
	{
		var G__841 = x
		_ = G__841
		return ___.F.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__841)
	}
}

func (_ *CljsCore_testT837) CljsCoreIMeta__() {}
func (_839 *CljsCore_testT837) X_meta_Arity1() interface{} {
	return _839.Meta838
}

func (_ *CljsCore_testT837) CljsCoreIWithMeta__() {}
func (_839 *CljsCore_testT837) X_with_meta_Arity2(meta838___1 interface{}) interface{} {
	return (&CljsCore_testT837{_839.F, _839.Baz, _839.Test_stuff, meta838___1})
}

var X__GT_t837 *cljs_core.AFn

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

type CljsCore_testT851 struct {
	Test_stuff interface{}
	Meta852    interface{}
}

func (_ *CljsCore_testT851) CljsCoreIHash__() {}
func (___ *CljsCore_testT851) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT851) CljsCoreIMeta__() {}
func (_853 *CljsCore_testT851) X_meta_Arity1() interface{} {
	return _853.Meta852
}

func (_ *CljsCore_testT851) CljsCoreIWithMeta__() {}
func (_853 *CljsCore_testT851) X_with_meta_Arity2(meta852___1 interface{}) interface{} {
	return (&CljsCore_testT851{_853.Test_stuff, meta852___1})
}

var X__GT_t851 *cljs_core.AFn

type CljsCore_testT854 struct {
	A          interface{}
	Test_stuff interface{}
	Meta855    interface{}
}

func (_ *CljsCore_testT854) CljsCoreIHash__() {}
func (___ *CljsCore_testT854) X_hash_Arity1() interface{} {
	return float64(42)
}

func (_ *CljsCore_testT854) CljsCoreIMeta__() {}
func (_856 *CljsCore_testT854) X_meta_Arity1() interface{} {
	return _856.Meta855
}

func (_ *CljsCore_testT854) CljsCoreIWithMeta__() {}
func (_856 *CljsCore_testT854) X_with_meta_Arity2(meta855___1 interface{}) interface{} {
	return (&CljsCore_testT854{_856.A, _856.Test_stuff, meta855___1})
}

var X__GT_t854 *cljs_core.AFn

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

type CljsCore_testT924 struct {
	From_seq   interface{}
	Make_seq   interface{}
	Mt         interface{}
	I__916     interface{}
	Count__915 interface{}
	Chunk__914 interface{}
	Seq__913   interface{}
	Test_stuff interface{}
	Meta925    interface{}
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
	return (&CljsCore_testT924{_926.From_seq, _926.Make_seq, _926.Mt, _926.I__916, _926.Count__915, _926.Chunk__914, _926.Seq__913, _926.Test_stuff, meta925___1})
}

var X__GT_t924 *cljs_core.AFn

type CljsCore_testT936 struct {
	From_seq           interface{}
	Make_seq           interface{}
	Mt                 interface{}
	Temp__4388__auto__ interface{}
	I__916             interface{}
	Count__915         interface{}
	Chunk__914         interface{}
	Seq__913           interface{}
	Test_stuff         interface{}
	Meta937            interface{}
}

func (_ *CljsCore_testT936) CljsCoreISeq__() {}
func (this *CljsCore_testT936) X_first_Arity1() interface{} {
	return cljs_core.First.X_invoke_Arity1(this.From_seq)
}

func (this *CljsCore_testT936) X_rest_Arity1() interface{} {
	{
		var G__940 = cljs_core.Rest.Arity1IQ(this.From_seq)
		_ = G__940
		return this.Make_seq.(cljs_core.CljsCoreIFn).X_invoke_Arity1(G__940)
	}
}

func (_ *CljsCore_testT936) CljsCoreISeqable__() {}
func (this *CljsCore_testT936) X_seq_Arity1() interface{} {
	return this
}

func (_ *CljsCore_testT936) CljsCoreIMeta__() {}
func (_938 *CljsCore_testT936) X_meta_Arity1() interface{} {
	return _938.Meta937
}

func (_ *CljsCore_testT936) CljsCoreIWithMeta__() {}
func (_938 *CljsCore_testT936) X_with_meta_Arity2(meta937___1 interface{}) interface{} {
	return (&CljsCore_testT936{_938.From_seq, _938.Make_seq, _938.Mt, _938.Temp__4388__auto__, _938.I__916, _938.Count__915, _938.Chunk__914, _938.Seq__913, _938.Test_stuff, meta937___1})
}

var X__GT_t936 *cljs_core.AFn

var Case_recur *cljs_core.AFn

var Xform *cljs_core.AFn

var Data interface{}

var Xf *cljs_core.AFn

var Test_map cljs_core.CljsCoreIMap

// A docstring
var Var_test *cljs_core.AFn

// A docstring
var Var_test_test *cljs_core.AFn

// A docstring
var Var_test_self_call *cljs_core.AFn

func Test_runner(t *testing.T) {
	assert.Equal(t, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), Test_stuff.X_invoke_Arity0().(*cljs_core.CljsCoreKeyword))
}
