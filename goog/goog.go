package goog

import "reflect"

func TypeOf(x interface{}) reflect.Kind {
	return reflect.TypeOf(x).Kind()
}

func IsString(x interface{}) bool {
	return TypeOf(x) == reflect.String
}

func IsFunction(x interface{}) bool {
	return TypeOf(x) == reflect.Func
}

func GetUid(obj interface{}) float64 {
	return float64(reflect.ValueOf(&obj).Pointer())
}
