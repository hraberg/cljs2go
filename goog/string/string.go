package string

import (
	"fmt"
	"hash/fnv"

	"github.com/hraberg/cljs.go/js"
)

type StringBuffer struct {
	Buffer interface{}
}

func (sb *StringBuffer) String() string {
	switch sb.Buffer.(type) {
	case nil:
		sb.Buffer = js.JSString("")
	case string:
		sb.Buffer = js.JSString(sb.Buffer.(string))
	}
	return string(sb.Buffer.(js.JSString))
}

func (sb *StringBuffer) Append(a1 interface{}) *StringBuffer {
	sb.Buffer = js.JSString(sb.String() + fmt.Sprint(a1))
	return sb
}

func HashCode(str js.JSString) float64 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return float64(h.Sum32())
}
