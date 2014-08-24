package hello

import . "github.com/hraberg/cljs.go/cljs/core"

var _main = Fn(func(args ...interface{}) interface{} {
	Println.Invoke_ArityVariadic("Hello World")
	return nil
})

func init() {
	X_STAR_main_cli_fn_STAR_ = _main
}
