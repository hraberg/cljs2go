package main

import (
	"log"
	"os/exec"
)

import . "github.com/hraberg/cljs.go/cljs/core"

var _main = Fn(func(args ...interface{}) interface{} {
	Println.Invoke_ArityVariadic("ClojureScript to Go [go]")
	goGet := exec.Command("go", "get", "code.google.com/p/go.tools/cmd/goimports")
	if out, err := goGet.CombinedOutput(); err != nil {
		log.Fatal(string(out[:]))
	}
	return nil
})

func init() {
	X_STAR_main_cli_fn_STAR_ = _main
}

func main() {
	Main()
}
