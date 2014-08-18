package string

import "hash/fnv"
import "bytes"

type StringBuffer struct {
	Buffer interface{}
}

func (sb *StringBuffer) String() string {
	return sb.Buffer.(string)
}

func (sb *StringBuffer) Append(a1 interface{}) *StringBuffer {
	if sb.Buffer == nil {
		sb.Buffer = ""
	}
	buffer := bytes.NewBufferString(sb.Buffer.(string))
	buffer.WriteString(a1.(string))
	sb.Buffer = buffer.String()
	return sb
}

func HashCode(str string) float64 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return float64(h.Sum32())
}
