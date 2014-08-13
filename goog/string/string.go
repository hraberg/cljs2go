package string

import "hash/fnv"
import "bytes"

type StringBuffer struct {
	bytes.Buffer
}

func (sb *StringBuffer) ToString() string {
	return sb.String()
}

func (sb *StringBuffer) Append(a1 string) StringBuffer {
	sb.WriteString(a1)
	return *sb
}

func HashCode(str string) float64 {
	h := fnv.New32a()
	h.Write([]byte(str))
	return float64(h.Sum32())
}
