package core

import (
	"fmt"
	"reflect"

	"testing"
)
import (
	gstring "github.com/hraberg/cljs2go/goog/string"
	"github.com/hraberg/cljs2go/js"
	"github.com/hraberg/cljs2go/js/Math"
	"github.com/stretchr/testify/assert"
)

func Test_Main(t *testing.T) {
	var mainArgs []interface{}
	X_STAR_main_cli_fn_STAR_ = Fn(func(args__ ...interface{}) interface{} {
		mainArgs = args__
		return nil
	})
	Main_()
	assert.Equal(t, 1, len(mainArgs))
}

var Baz = Fn(1, func(args ...interface{}) interface{} {
	x := args[0]
	xs := args[1]
	_ = x
	return Into_array.X_invoke_Arity1(xs)
}, func(x interface{}) interface{} {
	return x
})

func Test_Invoke(t *testing.T) {
	assert.True(t, Native_satisfies_QMARK_.X_invoke_Arity2(reflect.TypeOf((*CljsCoreIFn)(nil)).Elem(), Baz).(bool))
	PanicsWith(t, "Invalid arity: 0", func() { Baz.Call() })
	PanicsWith(t, "Invalid arity: 0", func() { Baz.X_invoke_Arity0() })
	assert.Equal(t, 1, Baz.MaxFixedArity)
	assert.True(t, Baz.isVariadic())
	assert.Equal(t, "Hello", Baz.X_invoke_Arity1("Hello"))
	assert.Equal(t, "Hello", X_invoke.X_invoke_Arity2(Baz, "Hello"))
	assert.Equal(t, []interface{}{"World"}, Baz.Call("Hello", Array_seq.X_invoke_Arity1([]interface{}{"World"})))
	assert.Equal(t, []interface{}{"World"}, Baz.X_invoke_ArityVariadic("Hello", Array_seq.X_invoke_Arity1([]interface{}{"World"})))
	assert.Equal(t, []interface{}{"World"}, X_invoke.X_invoke_ArityVariadic(Baz, "Hello", Array_seq.X_invoke_Arity1([]interface{}{"World"})))
	assert.Equal(t, []interface{}{"World"}, X_invoke.Call(Baz, "Hello", Array_seq.X_invoke_Arity1([]interface{}{"World"})))
	assert.Equal(t, []interface{}{"World"}, Apply.X_invoke_Arity2(Baz, Array_seq.X_invoke_Arity1([]interface{}{"Hello", "World"})))
	assert.Equal(t, []interface{}{"World"}, Apply.X_invoke_Arity3(Baz, "Hello", Array_seq.X_invoke_Arity1([]interface{}{"World"})))

	assert.Equal(t, []interface{}{"Another", "Brave", "And", "New", "World"},
		Apply.X_invoke_ArityVariadic(Baz, "Hello", "Another", "Brave", "And", Array_seq.X_invoke_Arity1([]interface{}{"New", Array_seq.X_invoke_Arity1([]interface{}{"World"})})))

	assert.Equal(t, 0, Apply.X_invoke_Arity2(X_PLUS_, Array_seq.X_invoke_Arity1([]interface{}{})))
	assert.Equal(t, 0, Apply.X_invoke_Arity2(X_PLUS_, CljsCorePersistentVector_EMPTY))
	assert.Equal(t, 5, Apply.X_invoke_Arity2(X_PLUS_, Array_seq.X_invoke_Arity1([]interface{}{5.0})))
	assert.Equal(t, 5, Apply.X_invoke_Arity2(X_PLUS_, Array_seq.X_invoke_Arity1([]interface{}{2.0, 3.0})))
	assert.Equal(t, 1, Apply.X_invoke_Arity3(X_PLUS_, 1.0, CljsCorePersistentVector_EMPTY))
	assert.Equal(t, 6, Apply.X_invoke_Arity3(X_PLUS_, 1.0, Array_seq.X_invoke_Arity1([]interface{}{2.0, 3.0})))
	assert.Equal(t, 8, Apply.X_invoke_Arity4(X_PLUS_, 1.0, 2.0, Array_seq.X_invoke_Arity1([]interface{}{5.0})))
	assert.Equal(t, 15, Apply.X_invoke_ArityVariadic(X_PLUS_, 1.0, 2.0, 3.0, 4.0, Array_seq.X_invoke_Arity1([]interface{}{Array_seq.X_invoke_Arity1([]interface{}{5.0})})))

	assert.Equal(t, 3, Apply.X_invoke_Arity3(X_PLUS_, float64(1), CljsCoreList_EMPTY.X_conj_Arity2(float64(2))))
	assert.Equal(t, 1, Count.X_invoke_Arity1(Apply.X_invoke_Arity3(List, 1.0, CljsCorePersistentVector_EMPTY)))
	assert.Equal(t, 1, Count.X_invoke_Arity1(Apply.X_invoke_Arity2(List, Array_seq.X_invoke_Arity1([]interface{}{5.0}))))
	assert.Equal(t, 0, Count.X_invoke_Arity1(Apply.X_invoke_Arity2(List, CljsCorePersistentVector_EMPTY)))

	assert.Equal(t, []interface{}{"World"}, Baz.X_invoke_Arity2("Hello", "World"))
	assert.Equal(t, []interface{}{"World", "Space"}, Baz.X_invoke_Arity3("Hello", "World", "Space"))

	assert.Equal(t, "", Str.X_invoke_Arity0())
	assert.Equal(t, "Hello", Str.X_invoke_Arity1("Hello"))
	assert.Equal(t, "1", Str.X_invoke_Arity1(1.0))
	assert.Equal(t, "HelloClojureWorld", Str.X_invoke_ArityVariadic("Hello", Array_seq.X_invoke_Arity1([]interface{}{"Clojure", "World"})))
}

func Test_Int32BitOperations(t *testing.T) {
	assert.Equal(t, -2023406815.0, Bit_or.X_invoke_Arity2(float64(0x87654321), 0.0))
	assert.Equal(t, 1985229336.0, Bit_or.X_invoke_Arity2(float64(0x76543218), 0.0))
	assert.Equal(t, 1985229328.0, Bit_shift_left.X_invoke_Arity2(float64(0x87654321), 4.0))
	assert.Equal(t, 141972530.0, Unsigned_bit_shift_right.X_invoke_Arity2(float64(0x87654321), 4.0))
	assert.Equal(t, 8.0, Unsigned_bit_shift_right.X_invoke_Arity2(float64(0x87654321), -4.0))

	assert.Equal(t, 8.0, Unsigned_bit_shift_right.X_invoke_Arity2(Bit_or.X_invoke_Arity2(float64(0x87654321), 0.0), -4.0))
	assert.Equal(t, 141972530.0, Unsigned_bit_shift_right.X_invoke_Arity2(Bit_or.X_invoke_Arity2(float64(0x87654321), 0.0), 4.0))
	assert.Equal(t, 1985229336.0, Int_rotate_left.X_invoke_Arity2(Bit_or.X_invoke_Arity2(float64(0x87654321), 0.0), 4.0))
	assert.Equal(t, 1985229336.0, Bit_or.X_invoke_Arity2(float64(0x76543218), 0.0))

	assert.Equal(t, Int_rotate_left.X_invoke_Arity2(Bit_or.X_invoke_Arity2(float64(0x87654321), 0.0), 4.0),
		Bit_or.X_invoke_Arity2(float64(0x76543218), 0.0))

	assert.Equal(t, -2023406815, Int32_(float64(int(float64(2271560481))|int(float64(0)))), 0)

	assert.Equal(t, Hash.X_invoke_Arity1((&CljsCoreKeyword{Ns: nil, Name: "a", Fqn: "a", X_hash: float64(-2123407586)})),
		Hash.X_invoke_Arity1(Keyword.X_invoke_Arity1("a")))
}

func Test_PrimitiveFn(t *testing.T) {
	fib := func(this *AFn) *AFn {
		return Fn(this, func(n float64) float64 {
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.Arity1FF(n-1.0) + this.Arity1FF(n-2.0)
			}
		})
	}(&AFn{})
	PanicsWith(t, "Invalid arity: 0", func() { fib.Call() })
	assert.Equal(t, 832040, fib.X_invoke_Arity1(30.0))

	odd := Fn(func(n interface{}) bool {
		return int(n.(float64))%2 != 0
	})
	assert.NotNil(t, odd.Arity1IB)
	assert.NotNil(t, odd.Arity1)
	assert.True(t, odd.X_invoke_Arity1(1.0).(bool))
	assert.False(t, odd.Arity1IB(2.0))

	assert.True(t, Not.X_invoke_Arity1(false).(bool))
	assert.False(t, Not.Arity1IB(true))
}

func Test_Protocols(t *testing.T) {
	symbol := Symbol.X_invoke_Arity2("foo", "bar")

	assert.True(t, symbol.(CljsCoreObject).Equiv(Symbol.X_invoke_Arity2("foo", "bar")))
	assert.True(t, Native_satisfies_QMARK_.X_invoke_Arity2(reflect.TypeOf((*CljsCoreObject)(nil)).Elem(), symbol).(bool))
	assert.True(t, Native_satisfies_QMARK_.X_invoke_Arity2(reflect.TypeOf((*CljsCoreINamed)(nil)).Elem(), symbol).(bool))
	assert.True(t, Native_satisfies_QMARK_.X_invoke_Arity2(reflect.TypeOf((*CljsCoreIFn)(nil)).Elem(), symbol).(bool))
	assert.Equal(t, "foo", symbol.(CljsCoreINamed).X_namespace_Arity1())
	assert.Equal(t, "foo", X_namespace.X_invoke_Arity1(symbol))
	PanicsWith(t, "Invalid arity: 0", func() { symbol.(*CljsCoreSymbol).X_invoke_Arity0() })
	PanicsWith(t, "Invalid arity: 0", func() { X_invoke.X_invoke_Arity1(symbol) })

	assert.Equal(t, "foo", Namespace.X_invoke_Arity1(symbol))
	PanicsWith(t, "Doesn't support namespace: 2", func() { Namespace.X_invoke_Arity1(2.0) })

	assert.Equal(t, "bar", Name.X_invoke_Arity1(symbol))
	PanicsWith(t, "Doesn't support name: 2", func() { Name.X_invoke_Arity1(2.0) })
	assert.Equal(t, "baz", Name.X_invoke_Arity1("baz"))

	assert.Equal(t, "foo", X_invoke.X_invoke_Arity2(X_namespace, symbol))
	assert.Equal(t, "bar", symbol.(CljsCoreINamed).X_name_Arity1())
	assert.Equal(t, "foo/bar", symbol.(fmt.Stringer).String())

	assert.Equal(t, 0, Decorate_(nil).(CljsCoreICounted).X_count_Arity1())
}

func Test_InteropViaReflection(t *testing.T) {
	sb := &gstring.StringBuffer{"foo"}
	assert.Equal(t, "foo", Native_get_instance_field.X_invoke_Arity2(sb, "Buffer"),
		"(.-buffer sb)")
	Native_set_instance_field.X_invoke_Arity3(sb, "Buffer", "bar")
	assert.Equal(t, "bar", sb.Buffer,
		"(set! (.-buffer sb) \"bar\")")
	assert.Equal(t, sb, Native_invoke_instance_method.X_invoke_Arity3(sb, "Append", []interface{}{"baz"}))
	assert.Equal(t, "barbaz", sb.String(),
		"(.append sb \"baz\")")

	assert.Equal(t, "H", Native_invoke_instance_method.X_invoke_Arity3("Hello", "CharAt", []interface{}{0.0}))

	assert.Equal(t, js.Number.MAX_VALUE, Native_get_instance_field.X_invoke_Arity2(js.Number, "MAX_VALUE"),
		"(.-MAX-VALUE js/Number)")

	assert.Equal(t, 3, Native_invoke_func.X_invoke_Arity2(Math.Floor, []interface{}{3.14}),
		"(Math/floor 3.14)")
	assert.Equal(t, "3.14", Native_invoke_func.X_invoke_Arity2(fmt.Sprint, []interface{}{3.14}),
		"(fmt/Sprint 3.14)")
	assert.Equal(t, 3.14, Native_invoke_func.X_invoke_Arity2(js.ParseFloat, []interface{}{"3.14"}),
		"(js/parseFloat \"3.14\")")

	assert.Equal(t, "ABC", Native_invoke_func.X_invoke_Arity2(js.String.FromCharCode, []interface{}{65.0, 66.0, 67.0}),
		"(.fromCharCode js/String 65 66 67)")
}

func PanicsWith(t *testing.T, message string, f assert.PanicTestFunc) {
	assert.Equal(t, message, func() (message string) {
		defer func() { message = fmt.Sprint(recover()) }()
		f()
		assert.Fail(t, "should panic")
		return
	}())
}
