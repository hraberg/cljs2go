package object

func Create(keyvals ...interface{}) map[string]interface{} {
	obj := make(map[string]interface{}, len(keyvals)/2)
	for i := 0; i < len(keyvals); i++ {
		obj[keyvals[i].(string)] = keyvals[i+1]
		i++
	}
	return obj
}

func ForEach(obj map[string]interface{}, f func(k, v, obj interface{}) interface{}) interface{} {
	for k, v := range obj {
		f(k, v, obj)
	}
	return nil
}
