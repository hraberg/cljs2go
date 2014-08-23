package goog

import "testing"

import (
	goog_array "github.com/hraberg/cljs.go/goog/array"
	goog_object "github.com/hraberg/cljs.go/goog/object"
	goog_string "github.com/hraberg/cljs.go/goog/string"
	"github.com/stretchr/testify/assert"
)

func Test_Goog(t *testing.T) {
	is := []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}
	goog_array.Shuffle(is)
	goog_array.StableSort(is, func(a, b interface{}) interface{} { return a.(float64) - b.(float64) })
	assert.Equal(t, []interface{}{5.0, 4.0, 3.0, 2.0, 1.0}, is)
	goog_array.Shuffle(is)
	goog_array.StableSort(is, goog_array.DefaultCompare)
	assert.Equal(t, []interface{}{1.0, 2.0, 3.0, 4.0, 5.0}, is)

	ss := []interface{}{"foo", "bar"}
	assert.True(t, IsArray(ss))
	assert.False(t, IsObject(ss))
	goog_array.StableSort(ss, goog_array.DefaultCompare)
	assert.Equal(t, []interface{}{"bar", "foo"}, ss)

	obj := goog_object.Create("foo", 2, "bar", 3)
	assert.True(t, IsObject(obj))
	assert.False(t, IsArray(obj))
	copy := map[string]interface{}{}
	goog_object.ForEach(obj, func(k, v, o interface{}) interface{} {
		assert.Equal(t, obj, o)
		assert.Equal(t, v, o.(map[string]interface{})[k.(string)])
		copy[k.(string)] = v
		return nil
	})
	assert.Equal(t, obj, copy)

	sb := goog_string.StringBuffer{}
	assert.Equal(t, "Hello JavaScript World", sb.Append("Hello Java").Append("Script World").String())
	assert.Equal(t, "Hello JavaScript World", sb.String())
	assert.Equal(t, 3.012568359e+09, (goog_string.HashCode("Hello World")))

	s := "Hello World"
	assert.False(t, IsObject(s))
	assert.True(t, IsString(s))
	assert.Equal(t, "string", TypeOf(s))

	assert.True(t, IsFunction(func() {}))

	x, y := "Hello", "Hello"
	assert.Equal(t, x, y)
	assert.NotEqual(t, GetUid(x), GetUid(y))
}
