package goog

import "reflect"

func TypeOf(x interface{}) string {
	return reflect.TypeOf(x).String()
}

func IsArray(x interface{}) bool {
	return reflect.ValueOf(x).Kind() == reflect.Slice
}

func IsObject(x interface{}) bool {
	return reflect.ValueOf(x).Kind() == reflect.Map
}

func IsString(x interface{}) bool {
	return reflect.ValueOf(x).Kind() == reflect.String
}

func IsFunction(x interface{}) bool {
	return reflect.ValueOf(x).Kind() == reflect.Func
}

func GetUid(obj interface{}) float64 {
	return float64(reflect.ValueOf(&obj).Pointer())
}
