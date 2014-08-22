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
	assert.Nil(t, RegExp{"Hello", ""}.Exec("World"))
	assert.Equal(t, []JSString{"Hello", "Hello"}, RegExp{"hello", "i"}.Exec("World Hello Hello"))
	assert.Equal(t, "HELLO World", (JSString("Hello World").Replace(RegExp{"hello", "i"},
		func(match interface{}) interface{} {
			return strings.ToUpper(fmt.Sprint(match))
		},
	)))
	assert.Equal(t, 6, (JSString("Hello World").Search(RegExp{"world", "i"})))
	assert.Equal(t, "(?i)Hello", (RegExp{"Hello", "i"}).String())

	date := Date{1407962432671}
	assert.Equal(t, 2014, date.GetUTCFullYear())
	assert.Equal(t, 7, date.GetUTCMonth())
	assert.Equal(t, 13, date.GetUTCDate())
	assert.Equal(t, 20, date.GetUTCHours())
	assert.Equal(t, 40, date.GetUTCMinutes())
	assert.Equal(t, 32, date.GetUTCSeconds())
	assert.Equal(t, 671, date.GetUTCMilliseconds())
	assert.Equal(t, "2014-08-13 21:40:32.000671 +0100 BST", date.String())

	assert.Equal(t, 3.14, ParseFloat(JSString("3.14")))
	assert.Equal(t, math.NaN(), ParseFloat(""))
	assert.Equal(t, 3, ParseInt("3", 10))
	assert.Equal(t, 10, ParseInt("a", 16))
	assert.Equal(t, math.NaN(), ParseInt("3.14", 10))
	assert.Equal(t, math.NaN(), ParseInt("x", 10))

	assert.Equal(t, "l", (JSString("Hello").CharAt(2)))
	assert.Equal(t, 108, (JSString("Hello").CharCodeAt(2)))
	assert.True(t, Array.IsArray(JSArray{"Hello", "World"}))
	assert.False(t, Array.IsArray("Hello"))
}
