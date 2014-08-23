package string

import (
	"bytes"
	"fmt"
	"hash/fnv"
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

func (this *StringBuffer) ToString() string {
	return this.buffer().String()
}

func (this *StringBuffer) String() string {
	return this.ToString()
}

func (this *StringBuffer) Append(a1 interface{}) *StringBuffer {
	this.buffer().WriteString(fmt.Sprint(a1))
	return this
}

func HashCode(str string) float64 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return float64(h.Sum32())
}
