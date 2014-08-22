package string

import (
	"bytes"
	"fmt"
	"hash/fnv"

	"github.com/hraberg/cljs.go/js"
)

type StringBuffer struct {
	Buffer interface{}
}

func (sb *StringBuffer) buffer() *bytes.Buffer {
	switch sb.Buffer.(type) {
	case *bytes.Buffer:
	case nil:
		sb.Buffer = &bytes.Buffer{}
	default:
		sb.Buffer = bytes.NewBufferString(fmt.Sprint(sb.Buffer))
	}
	return sb.Buffer.(*bytes.Buffer)
}

func (sb *StringBuffer) String() string {
	return sb.buffer().String()
}

func (sb *StringBuffer) Append(a1 interface{}) *StringBuffer {
	sb.buffer().WriteString(fmt.Sprint(a1))
	return sb
}

func HashCode(str js.JSString) float64 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return float64(h.Sum32())
}
