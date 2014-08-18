package string

import "hash/fnv"

type StringBuffer struct {
	Buffer interface{}
}

func (sb *StringBuffer) String() string {
	if sb.Buffer == nil {
		sb.Buffer = ""
	}
	return sb.Buffer.(string)
}

func (sb *StringBuffer) Append(a1 interface{}) *StringBuffer {
	sb.Buffer = sb.String() + a1.(string)
	return sb
}

func HashCode(str string) float64 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return float64(h.Sum32())
}
