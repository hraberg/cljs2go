// Compiled by ClojureScript to Go 0.0-2411
// cljs.keyword-test

package keyword_test

import (
	"strings"
	"testing"

	cljs_core "github.com/hraberg/cljs2go/cljs/core"
	"github.com/hraberg/cljs2go/js"
	"github.com/stretchr/testify/assert"
)

func init() {
	Test_keyword = func(test_keyword *cljs_core.AFn) *cljs_core.AFn {
		return cljs_core.Fn(test_keyword, 0, func() interface{} {
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: "cljs.keyword-test", Name: "bar", Fqn: "cljs.keyword-test/bar", X_hash: float64(757060530)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.keyword-test", Name: "bar", Fqn: "cljs.keyword-test/bar", X_hash: float64(757060530)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :cljs.keyword-test/bar :cljs.keyword-test/bar)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: "cljs.keyword-other", Name: "foo", Fqn: "cljs.keyword-other/foo", X_hash: float64(-1299575836)}), (&cljs_core.CljsCoreKeyword{Ns: "cljs.keyword-other", Name: "foo", Fqn: "cljs.keyword-other/foo", X_hash: float64(-1299575836)})) {
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :cljs.keyword-other/foo :cljs.keyword-other/foo)").(string)}, ``)}))
			}
			if cljs_core.X_EQ_.Arity2IIB((&cljs_core.CljsCoreKeyword{Ns: "clojure.core", Name: "foo", Fqn: "clojure.core/foo", X_hash: float64(2105943067)}), (&cljs_core.CljsCoreKeyword{Ns: "clojure.core", Name: "foo", Fqn: "clojure.core/foo", X_hash: float64(2105943067)})) {
				return nil
			} else {
				panic((&js.Error{strings.Join([]string{cljs_core.Str.X_invoke_Arity1("Assert failed: ").(string), cljs_core.Str.X_invoke_Arity1("(= :clojure.core/foo :clojure.core/foo)").(string)}, ``)}))
			}
		})
	}(&cljs_core.AFn{})

}

var Test_keyword *cljs_core.AFn

func Test_runner(t *testing.T) {
	Test_keyword.X_invoke_Arity0()
	assert.True(t, true)
}
