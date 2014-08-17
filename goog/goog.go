package goog

import "reflect"

func TypeOf(x interface{}) reflect.Type {
	return reflect.TypeOf(x)
}

func IsString(x interface{}) bool {
	return TypeOf(x).Kind() == reflect.String
}

func IsFunction(x interface{}) bool {
	return TypeOf(x).Kind() == reflect.Func
}

func GetUid(obj interface{}) float64 {
	return float64(reflect.ValueOf(&obj).Pointer())
}
