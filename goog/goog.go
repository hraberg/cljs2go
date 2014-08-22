package goog

import (
	"reflect"

	"github.com/hraberg/cljs.go/js"
)

func TypeOf(x interface{}) js.JSString {
	return js.JSString(reflect.TypeOf(x).String())
}

func IsArray(x interface{}) bool {
	return TypeOf(x) == `js.JSArray`
}

func IsObject(x interface{}) bool {
	return TypeOf(x) == `js.JSObject`
}

func IsString(x interface{}) bool {
	return TypeOf(x) == `js.JSString`
}

func IsFunction(x interface{}) bool {
	return reflect.TypeOf(x).Kind() == reflect.Func
}

func GetUid(obj interface{}) float64 {
	return float64(reflect.ValueOf(&obj).Pointer())
}
