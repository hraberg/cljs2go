
// +build ignore

package main

import cljs_core `github.com/hraberg/cljs.go/cljs/core`
import main_ns `.`

func main() {
	cljs_core.X_STAR_main_cli_fn_STAR_ = main_ns.X_main
	cljs_core.Main_()
}
