package object

import "github.com/hraberg/cljs.go/js"

func Create(keyvals ...interface{}) js.JSObject {
	obj := make(js.JSObject, len(keyvals)/2)
	for i := 0; i < len(keyvals); i++ {
		obj[keyvals[i].(js.JSString)] = keyvals[i+1]
		i++
	}
	return obj
}

func ForEach(obj js.JSObject, f func(k, v, obj interface{}) interface{}) interface{} {
	for k, v := range obj {
		f(k, v, obj)
	}
	return nil
}
