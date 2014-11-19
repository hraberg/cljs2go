package goog

import "reflect"

func TypeOf(x interface{}) string {
	if x == nil {
		return reflect.ValueOf(&x).Elem().Type().String()
	}
	return reflect.ValueOf(x).Type().String()
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
