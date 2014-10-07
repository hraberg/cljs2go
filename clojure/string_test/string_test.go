// Compiled by ClojureScript to Go 0.0-2356
// clojure.string-test

package string_test

import (
	"testing"

	cljs_core "github.com/hraberg/cljs.go/cljs/core"
	clojure_string "github.com/hraberg/cljs.go/clojure/string"
	"github.com/hraberg/cljs.go/js"
	"github.com/stretchr/testify/assert"
)

func init() {
	Test_string = func(test_string *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test_string, 0, func() interface{} {
			if cljs_core.X_EQ_.Arity2IIB("", clojure_string.Reverse.X_invoke_Arity1("")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\" (s/reverse \"\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("tab", clojure_string.Reverse.X_invoke_Arity1("bat")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"tab\" (s/reverse \"bat\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("faabar", clojure_string.Replace.X_invoke_Arity3("foobar", "o", "a")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"faabar\" (s/replace \"foobar\" \\o \\a))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("barbarbar", clojure_string.Replace.X_invoke_Arity3("foobarfoo", "foo", "bar")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"barbarbar\" (s/replace \"foobarfoo\" \"foo\" \"bar\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("FOObarFOO", clojure_string.Replace.X_invoke_Arity3("foobarfoo", (&js.RegExp{Pattern: `foo`, Flags: ``}), clojure_string.Upper_case)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"FOObarFOO\" (s/replace \"foobarfoo\" #\"foo\" s/upper-case))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("barbar)foo", clojure_string.Replace.X_invoke_Arity3("foo(bar)foo", "foo(", "bar")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"barbar)foo\" (s/replace \"foo(bar)foo\" \"foo(\" \"bar\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("", clojure_string.Join.X_invoke_Arity1(nil)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\" (s/join nil))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("", clojure_string.Join.X_invoke_Arity1(cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\" (s/join []))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("1", clojure_string.Join.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"1\" (s/join [1]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("12", clojure_string.Join.X_invoke_Arity1((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"12\" (s/join [1 2]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("1,2,3", clojure_string.Join.X_invoke_Arity2(",", (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"1,2,3\" (s/join \\, [1 2 3]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("", clojure_string.Join.X_invoke_Arity2(",", cljs_core.CljsCorePersistentVector_EMPTY)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\" (s/join \\, []))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("1 and-a 2 and-a 3", clojure_string.Join.X_invoke_Arity2(" and-a ", (&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{float64(1), float64(2), float64(3)}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"1 and-a 2 and-a 3\" (s/join \" and-a \" [1 2 3]))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("FOOBAR", clojure_string.Upper_case.X_invoke_Arity1("Foobar")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"FOOBAR\" (s/upper-case \"Foobar\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("foobar", clojure_string.Lower_case.X_invoke_Arity1("FooBar")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"foobar\" (s/lower-case \"FooBar\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("Foobar", clojure_string.Capitalize.X_invoke_Arity1("foobar")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"Foobar\" (s/capitalize \"foobar\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("Foobar", clojure_string.Capitalize.X_invoke_Arity1("FOOBAR")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"Foobar\" (s/capitalize \"FOOBAR\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a", "b"}, nil}), clojure_string.Split.X_invoke_Arity2("a-b", (&js.RegExp{Pattern: `-`, Flags: ``}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"a\" \"b\"] (s/split \"a-b\" #\"-\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a", "b", "c"}, nil}), clojure_string.Split.X_invoke_Arity3("a-b-c", (&js.RegExp{Pattern: `-`, Flags: ``}), float64(-1))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"a\" \"b\" \"c\"] (s/split \"a-b-c\" #\"-\" -1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a", "b", "c"}, nil}), clojure_string.Split.X_invoke_Arity3("a-b-c", (&js.RegExp{Pattern: `-`, Flags: ``}), float64(0))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"a\" \"b\" \"c\"] (s/split \"a-b-c\" #\"-\" 0))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a-b-c"}, nil}), clojure_string.Split.X_invoke_Arity3("a-b-c", (&js.RegExp{Pattern: `-`, Flags: ``}), float64(1))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"a-b-c\"] (s/split \"a-b-c\" #\"-\" 1))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(2), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a", "b-c"}, nil}), clojure_string.Split.X_invoke_Arity3("a-b-c", (&js.RegExp{Pattern: `-`, Flags: ``}), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"a\" \"b-c\"] (s/split \"a-b-c\" #\"-\" 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a", "b", "c"}, nil}), clojure_string.Split.X_invoke_Arity3("a-b-c", (&js.RegExp{Pattern: `-`, Flags: ``}), float64(3))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"a\" \"b\" \"c\"] (s/split \"a-b-c\" #\"-\" 3))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a", "b", "c"}, nil}), clojure_string.Split.X_invoke_Arity3("a-b-c", (&js.RegExp{Pattern: `-`, Flags: ``}), float64(4))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"a\" \"b\" \"c\"] (s/split \"a-b-c\" #\"-\" 4))")}))
			}
			if cljs_core.Vector_QMARK_.Arity1IB(clojure_string.Split.X_invoke_Arity2("abc", (&js.RegExp{Pattern: `-`, Flags: ``}))) {
			} else {
				panic((&js.Error{("Assert failed: (vector? (s/split \"abc\" #\"-\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a-b-c"}, nil}), clojure_string.Split.X_invoke_Arity3("a-b-c", (&js.RegExp{Pattern: `x`, Flags: ``}), float64(2))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"a-b-c\"] (s/split \"a-b-c\" #\"x\" 2))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(5), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"", "a", "b", "c", ""}, nil}), clojure_string.Split.X_invoke_Arity3("abc", cljs_core.Re_pattern.X_invoke_Arity1("").(*js.RegExp), float64(5))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"\" \"a\" \"b\" \"c\" \"\"] (s/split \"abc\" (re-pattern \"\") 5))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(1), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"a"}, nil}), clojure_string.Split.X_invoke_Arity2("ab", (&js.RegExp{Pattern: `b`, Flags: ``}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [\"a\"] (s/split \"ab\" #\"b\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCorePersistentVector_EMPTY, clojure_string.Split.X_invoke_Arity2("ab", (&js.RegExp{Pattern: `ab`, Flags: ``}))) {
			} else {
				panic((&js.Error{("Assert failed: (= [] (s/split \"ab\" #\"ab\"))")}))
			}
			{
				var result_4048 = clojure_string.Split_lines.X_invoke_Arity1("one\ntwo\r\nthree")
				_ = result_4048
				if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCorePersistentVector{nil, float64(3), float64(5), cljs_core.CljsCorePersistentVector_EMPTY_NODE, []interface{}{"one", "two", "three"}, nil}), result_4048) {
				} else {
					panic((&js.Error{("Assert failed: (= [\"one\" \"two\" \"three\"] result)")}))
				}
				if cljs_core.Vector_QMARK_.Arity1IB(result_4048) {
				} else {
					panic((&js.Error{("Assert failed: (vector? result)")}))
				}
			}
			if cljs_core.X_EQ_.Arity2IIB(cljs_core.CljsCoreList_EMPTY.X_conj_Arity2("foo"), clojure_string.Split_lines.X_invoke_Arity1("foo")) {
			} else {
				panic((&js.Error{("Assert failed: (= (list \"foo\") (s/split-lines \"foo\"))")}))
			}
			if cljs_core.Truth_(clojure_string.Blank_QMARK_.X_invoke_Arity1(nil)) {
			} else {
				panic((&js.Error{("Assert failed: (s/blank? nil)")}))
			}
			if cljs_core.Truth_(clojure_string.Blank_QMARK_.X_invoke_Arity1("")) {
			} else {
				panic((&js.Error{("Assert failed: (s/blank? \"\")")}))
			}
			if cljs_core.Truth_(clojure_string.Blank_QMARK_.X_invoke_Arity1(" ")) {
			} else {
				panic((&js.Error{("Assert failed: (s/blank? \" \")")}))
			}
			if cljs_core.Truth_(clojure_string.Blank_QMARK_.X_invoke_Arity1(" \t \n  \r ")) {
			} else {
				panic((&js.Error{("Assert failed: (s/blank? \" \\t \\n  \\r \")")}))
			}
			if cljs_core.Not.Arity1IB(clojure_string.Blank_QMARK_.X_invoke_Arity1("  foo  ")) {
			} else {
				panic((&js.Error{("Assert failed: (not (s/blank? \"  foo  \"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("&lt;foo&amp;bar&gt;", clojure_string.Escape.X_invoke_Arity2("<foo&bar>", (&cljs_core.CljsCorePersistentArrayMap{nil, float64(3), []interface{}{"&", "&amp;", "<", "&lt;", ">", "&gt;"}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"&lt;foo&amp;bar&gt;\" (s/escape \"<foo&bar>\" {\\& \"&amp;\", \\< \"&lt;\", \\> \"&gt;\"}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(" \\\"foo\\\" ", clojure_string.Escape.X_invoke_Arity2(" \"foo\" ", (&cljs_core.CljsCorePersistentArrayMap{nil, float64(1), []interface{}{"\"", "\\\""}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \" \\\\\\\"foo\\\\\\\" \" (s/escape \" \\\"foo\\\" \" {\\\" \"\\\\\\\"\"}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("faabor", clojure_string.Escape.X_invoke_Arity2("foobar", (&cljs_core.CljsCorePersistentArrayMap{nil, float64(2), []interface{}{"a", "o", "o", "a"}, nil}))) {
			} else {
				panic((&js.Error{("Assert failed: (= \"faabor\" (s/escape \"foobar\" {\\a \\o, \\o \\a}))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("barbarfoo", clojure_string.Replace_first.X_invoke_Arity3("foobarfoo", "foo", "bar")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"barbarfoo\" (s/replace-first \"foobarfoo\" \"foo\" \"bar\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("barbarfoo", clojure_string.Replace_first.X_invoke_Arity3("foobarfoo", (&js.RegExp{Pattern: `foo`, Flags: ``}), "bar")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"barbarfoo\" (s/replace-first \"foobarfoo\" #\"foo\" \"bar\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("z.ology", clojure_string.Replace_first.X_invoke_Arity3("zoology", "o", ".")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"z.ology\" (s/replace-first \"zoology\" \\o \\.))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("FOObarfoo", clojure_string.Replace_first.X_invoke_Arity3("foobarfoo", (&js.RegExp{Pattern: `foo`, Flags: ``}), clojure_string.Upper_case)) {
			} else {
				panic((&js.Error{("Assert failed: (= \"FOObarfoo\" (s/replace-first \"foobarfoo\" #\"foo\" s/upper-case))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("foo ", clojure_string.Triml.X_invoke_Arity1(" foo ")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"foo \" (s/triml \" foo \"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("", clojure_string.Triml.X_invoke_Arity1("   ")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\" (s/triml \"   \"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB(" foo", clojure_string.Trimr.X_invoke_Arity1(" foo ")) {
			} else {
				panic((&js.Error{("Assert failed: (= \" foo\" (s/trimr \" foo \"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("", clojure_string.Trimr.X_invoke_Arity1("   ")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\" (s/trimr \"   \"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("foo", clojure_string.Trim.X_invoke_Arity1("  foo  \r\n")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"foo\" (s/trim \"  foo  \\r\\n\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("foo", clojure_string.Trim_newline.X_invoke_Arity1("foo\n")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"foo\" (s/trim-newline \"foo\\n\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("foo", clojure_string.Trim_newline.X_invoke_Arity1("foo\r\n")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"foo\" (s/trim-newline \"foo\\r\\n\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("foo", clojure_string.Trim_newline.X_invoke_Arity1("foo")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"foo\" (s/trim-newline \"foo\"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("foo\r ", clojure_string.Trim_newline.X_invoke_Arity1("foo\r ")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"foo\\r \" (s/trim-newline \"foo\\r \"))")}))
			}
			if cljs_core.X_EQ_.Arity2IIB("", clojure_string.Trim_newline.X_invoke_Arity1("")) {
			} else {
				panic((&js.Error{("Assert failed: (= \"\" (s/trim-newline \"\"))")}))
			}
			return (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)})
		})
	}(&cljs_core.AFn{})

}

var Test_string *cljs_core.AFn

func Test_runner(t *testing.T) {
	assert.Equal(t, (&cljs_core.CljsCoreKeyword{Ns: nil, Name: "ok", Fqn: "ok", X_hash: float64(967785236)}), Test_string.X_invoke_Arity0().(*cljs_core.CljsCoreKeyword))
}
