package String

import "bytes"

func FromCharCode(num ...interface{}) interface{} {
	var buffer bytes.Buffer
	for _, n := range num {
		buffer.WriteRune(rune(n.(int)))
	}
	return buffer.String()
}
