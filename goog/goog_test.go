package goog

import "testing"

import (
	goog_array "github.com/hraberg/cljs.go/goog/array"
	goog_object "github.com/hraberg/cljs.go/goog/object"
	goog_string "github.com/hraberg/cljs.go/goog/string"
	"github.com/hraberg/cljs.go/js"
	"github.com/stretchr/testify/assert"
)

func Test_Goog(t *testing.T) {
	is := js.JSArray{1.0, 2.0, 3.0, 4.0, 5.0}
	goog_array.Shuffle(is)
	goog_array.StableSort(is, func(a, b interface{}) interface{} { return a.(float64) - b.(float64) })
	assert.Equal(t, js.JSArray{5.0, 4.0, 3.0, 2.0, 1.0}, is)
	goog_array.Shuffle(is)
	goog_array.StableSort(is, goog_array.DefaultCompare)
	assert.Equal(t, js.JSArray{1.0, 2.0, 3.0, 4.0, 5.0}, is)

	ss := js.JSArray{"foo", "bar"}
	goog_array.StableSort(ss, goog_array.DefaultCompare)
	assert.Equal(t, js.JSArray{"bar", "foo"}, ss)

	obj := goog_object.Create("foo", 2, "bar", 3)
	copy := js.JSObject{}
	goog_object.ForEach(obj, func(k, v, o interface{}) interface{} {
		assert.Equal(t, obj, o)
		assert.Equal(t, v, o.(js.JSObject)[k.(string)])
		copy[k.(string)] = v
		return nil
	})
	assert.Equal(t, obj, copy)

	sb := goog_string.StringBuffer{}

	assert.Equal(t, "Hello JavaScript World", sb.Append("Hello Java").Append("Script World").String())
	assert.Equal(t, "Hello JavaScript World", sb.String())
	assert.Equal(t, 3.012568359e+09, (goog_string.HashCode("Hello World")))
}
