package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Printf("ClojureScript to Go [go]\n")
	goGet := exec.Command("go", "get", "code.google.com/p/go.tools/cmd/goimports")
	if out, err := goGet.CombinedOutput(); err != nil {
		log.Fatal(string(out[:]))
	}
}
