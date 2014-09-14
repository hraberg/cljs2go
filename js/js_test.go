package js

import (
	"fmt"
	"math"
	"strings"
	"testing"
)
import (
	"github.com/hraberg/cljs.go/js/Math"
	"github.com/stretchr/testify/assert"
)

func Test_JS(t *testing.T) {
	assert.Equal(t, math.Inf(1), Infinity)
	assert.Equal(t, math.MaxFloat64, Number.MAX_VALUE)
	assert.Equal(t, 0.6046602879796196, Math.Random())
	assert.Equal(t, 3, Math.Ceil(2.6))
	assert.Equal(t, 2, Math.Floor(2.6))
	assert.Equal(t, 12, Math.Imul(2.3, 6.7))
	assert.Equal(t, "ABC", String.FromCharCode(65, 66, 67))
	assert.Nil(t, (&RegExp{"Hello", ""}).Exec("World"))
	assert.Equal(t, []interface{}{"Hello", "Hello"}, (&RegExp{"hello", "i"}).Exec("World Hello Hello"))
	assert.Equal(t, "HELLO World", JSString_("Hello World").Replace(&RegExp{"hello", "i"},
		func(match interface{}) interface{} {
			return strings.ToUpper(fmt.Sprint(match))
		},
	))
	assert.Equal(t, 6, JSString_("Hello World").Search(&RegExp{"world", "i"}))
	assert.Equal(t, "(?i)Hello", (&RegExp{"Hello", "i"}).String())

	date := Date{1407962432671}
	assert.Equal(t, 2014, date.GetUTCFullYear())
	assert.Equal(t, 7, date.GetUTCMonth())
	assert.Equal(t, 13, date.GetUTCDate())
	assert.Equal(t, 20, date.GetUTCHours())
	assert.Equal(t, 40, date.GetUTCMinutes())
	assert.Equal(t, 32, date.GetUTCSeconds())
	assert.Equal(t, 671, date.GetUTCMilliseconds())
	assert.Equal(t, "2014-08-13 21:40:32.000671 +0100 BST", date.String())

	assert.Equal(t, 3.14, ParseFloat("3.14"))
	assert.Equal(t, math.NaN(), ParseFloat(""))
	assert.Equal(t, 3, ParseInt("3", 10))
	assert.Equal(t, 10, ParseInt("a", 16))
	assert.Equal(t, math.NaN(), ParseInt("3.14", 10))
	assert.Equal(t, math.NaN(), ParseInt("x", 10))

	assert.Equal(t, "1", JSNumber(1).ToString())
	assert.Equal(t, "3.14", JSNumber(3.14).ToString())
	assert.Equal(t, "true", JSBoolean(true).ToString())
	assert.Equal(t, "", JSNil{}.ToString())
	assert.Equal(t, "Hello", JSString_("Hello").ToString())
	assert.Equal(t, "[Hello World]", JSArray_(&[]interface{}{"Hello", "World"}).ToString())

	assert.Equal(t, "l", (JSString_("Hello").CharAt(2)))
	assert.Equal(t, 108, (JSString_("Hello").CharCodeAt(2)))

	assert.Equal(t, []interface{}{"Hello", "World"}, JSString_("Hello World").Split(" "))
	assert.Equal(t, []interface{}{"Hello"}, JSString_("Hello World").Split(" ", 1.0))
	assert.Equal(t, []interface{}{"Hello", "World"}, JSString_("Hello World").Split(RegExp{"\\s+", ""}))
	assert.Equal(t, []interface{}{"F", "o", "o"}, JSString_("Foo").Split())
	assert.Equal(t, []interface{}{"Foo"}, JSString_("Foo").Split(1.0))

	assert.Equal(t, "HELLO", JSString_("Hello").ToUpperCase())
	assert.Equal(t, "hello", JSString_("Hello").ToLowerCase())

	assert.Equal(t, "llo", JSString_("Hello").Substring(2.0))
	assert.Equal(t, "ell", JSString_("Hello").Substring(1.0, 4.0))
	assert.Equal(t, "l", JSString_("Hello").Substring(3.0, 2.0))

	arr := []interface{}{"Hello", "Earth", "World", "!"}
	assert.Equal(t, []interface{}{"Earth", "World"}, JSArray_(&arr).Splice(1.0, 2.0, "Hyper", "Space"))
	assert.Equal(t, []interface{}{"Hello", "Hyper", "Space", "!"}, arr)

	assert.Equal(t, []interface{}{"Hyper", "Space"}, JSArray_(&arr).Slice(1, -1))
	assert.Equal(t, []interface{}{"Hello", "Hyper", "Space", "!"}, arr)

	arr = []interface{}{"Hello", "World"}
	assert.Equal(t, "World", JSArray_(&arr).Pop())
	assert.Equal(t, []interface{}{"Hello"}, arr)

	arr = []interface{}{"Hello", "World"}
	assert.Equal(t, 3, JSArray_(&arr).Push("Space"))
	assert.Equal(t, []interface{}{"Hello", "World", "Space"}, arr)

	var xs interface{} = []interface{}{"Hello", "World"}
	assert.Equal(t, 2, JSArray_(&xs).Length)
	assert.Equal(t, "World", JSArray_(&xs).Pop())
	assert.Equal(t, 1, JSArray_(&xs).Length)
	assert.Equal(t, []interface{}{"Hello"}, xs)
}
