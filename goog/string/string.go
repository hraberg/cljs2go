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

func (this *StringBuffer) buffer() *bytes.Buffer {
	switch this.Buffer.(type) {
	case *bytes.Buffer:
	case nil:
		this.Buffer = &bytes.Buffer{}
	default:
		this.Buffer = bytes.NewBufferString(fmt.Sprint(this.Buffer))
	}
	return this.Buffer.(*bytes.Buffer)
}

func (this *StringBuffer) Equiv(other interface{}) bool {
	if ptr, instanceof := other.(*StringBuffer); instanceof {
		other = *ptr
	}
	return *this == other
}

func (this *StringBuffer) ToString() js.JSString {
	return js.JSString(this.buffer().String())
}

func (this *StringBuffer) String() string {
	return string(this.ToString())
}

func (this *StringBuffer) Append(a1 interface{}) *StringBuffer {
	this.buffer().WriteString(fmt.Sprint(a1))
	return this
}

func HashCode(str js.JSString) float64 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return float64(h.Sum32())
}
