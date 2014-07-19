package main

import (
	"fmt"
)

var plus_one = func(x int) int {
	return x + 1
}

func main() {
	fmt.Printf("ClojureScript to Go [go]\n")
	plus_one(1)
}
