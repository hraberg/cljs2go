package string

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"regexp"
	"strings"
	"unicode"
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

func (this *StringBuffer) GetLength() float64 {
	return float64(this.buffer().Len())
}

func (this *StringBuffer) Append(a1 interface{}) *StringBuffer {
	_, _ = this.buffer().WriteString(fmt.Sprint(a1))
	return this
}

func HashCode(str string) float64 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(str))
	return float64(h.Sum32())
}

func RegExpEscape(s string) string {
	return regexp.QuoteMeta(s)
}

func Trim(str string) string {
	return strings.TrimSpace(str)
}

func TrimLeft(str string) string {
	return strings.TrimLeftFunc(str, unicode.IsSpace)
}

func TrimRight(str string) string {
	return strings.TrimRightFunc(str, unicode.IsSpace)
}

func IsEmptySafe(str interface{}) bool {
	switch str := str.(type) {
	case string:
		return IsBreakingWhitespace(str)
	case nil:
		return true
	default:
		return false
	}
}

func IsNumeric(str string) bool {
	return len(strings.FieldsFunc(str, unicode.IsNumber)) == 0
}

func IsBreakingWhitespace(str string) bool {
	return len(strings.FieldsFunc(str, unicode.IsSpace)) == 0
}

func Contains(str, subString string) bool {
	return strings.Contains(str, subString)
}

func BuildString(var_args ...interface{}) string {
	ss := make([]string, len(var_args))
	for i, v := range var_args {
		if v == nil {
			ss[i] = ``
		} else {
			ss[i] = fmt.Sprint(v)
		}
	}
	return strings.Join(ss, ``)
}
