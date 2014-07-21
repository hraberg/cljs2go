package main

import "fmt"

func double(x interface{}) float64 {
	switch x.(type) {
	case int:
		return float64(x.(int))
	case int64:
		return float64(x.(int64))
	default:
		return x.(float64)
	}
}

func long(x interface{}) int64 {
	switch x.(type) {
	case int:
		return int64(x.(int))
	case float64:
		return int64(x.(float64))
	default:
		return x.(int64)
	}
}

func plus_one(x interface{}) interface{} {
	return (double(x) + (1))
}

func main() {
	fmt.Printf("ClojureScript to Go [go] %v\n", plus_one(1))
}
