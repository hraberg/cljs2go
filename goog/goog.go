package goog

import "reflect"

func TypeOf(x interface{}) reflect.Type {
	return reflect.TypeOf(x)
}

func IsArray(x interface{}) bool {
	return TypeOf(x).String() == `js.JSArray`
}

func IsObject(x interface{}) bool {
	return TypeOf(x).String() == `js.JSObject`
}

func IsString(x interface{}) bool {
	return TypeOf(x).String() == `js.JSString`
}

func IsFunction(x interface{}) bool {
	return TypeOf(x).Kind() == reflect.Func
}

func GetUid(obj interface{}) float64 {
	return float64(reflect.ValueOf(&obj).Pointer())
}
