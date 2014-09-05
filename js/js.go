package js

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

// This file contains a thin js runtime layer so ClojureScript itself can run with minimal modifications.

type Error struct {
	Message interface{}
}

type TypeError struct {
	Message interface{}
}

func (e *Error) Error() string {
	return fmt.Sprint(e.Message)
}

func (e *TypeError) Error() string {
	return fmt.Sprint(e.Message)
}

type JSObject map[string]interface{}
type JSNil struct{}
type JSBoolean bool
type JSNumber float64

type Date struct {
	Millis float64
}

func (this *Date) time() time.Time {
	return time.Unix(int64(this.Millis)/1000, 1000*(int64(this.Millis)%1000))
}

func (this *Date) GetUTCFullYear() float64 {
	return float64(this.time().UTC().Year())
}

func (this *Date) GetUTCMonth() float64 {
	return float64(this.time().UTC().Month() - 1)
}

func (this *Date) GetUTCDate() float64 {
	return float64(this.time().UTC().Day())
}

func (this *Date) GetUTCHours() float64 {
	return float64(this.time().UTC().Hour())
}

func (this *Date) GetUTCMinutes() float64 {
	return float64(this.time().UTC().Minute())
}

func (this *Date) GetUTCSeconds() float64 {
	return float64(this.time().UTC().Second())
}

func (this *Date) GetUTCMilliseconds() float64 {
	return float64(this.time().UTC().Nanosecond() / 1000)
}

func (this *Date) String() string {
	return this.time().String()
}

func (this *Date) ToString() string {
	return this.String()
}

type RegExp struct {
	Pattern interface{}
	Flags   interface{}
}

func (this *RegExp) compile() *regexp.Regexp {
	pattern, flags := this.Pattern.(string), this.Flags.(string)
	if len(flags) != 0 {
		pattern = "(?" + flags + ")" + pattern
	}
	return regexp.MustCompile(pattern)
}

func (this *RegExp) Exec(str string) []interface{} {
	if match := this.compile().FindAllString(string(str), -1); match != nil {
		strs := make([]interface{}, len(match))
		for i, v := range match {
			strs[i] = v
		}
		return strs
	}
	return nil
}

func (this *RegExp) String() string {
	return this.compile().String()
}

var Number = struct{ MAX_VALUE float64 }{math.MaxFloat64}

var Infinity = math.Inf(1)

func IsNaN(x float64) bool {
	return math.IsNaN(x)
}

func ParseFloat(str string) float64 {
	if val, ok := strconv.ParseFloat(str, 64); ok == nil {
		return val
	}
	return math.NaN()
}

func ParseInt(str string, radix float64) float64 {
	if val, ok := strconv.ParseInt(str, int(radix), 64); ok == nil {
		return float64(val)
	}
	return math.NaN()
}

var String = struct {
	FromCharCode func(...float64) string
}{func(num ...float64) string {
	var buffer bytes.Buffer
	for _, n := range num {
		_, _ = buffer.WriteRune(rune(int(n)))
	}
	return buffer.String()
}}

type JSString struct {
	Length float64
	str    string
}

func JSString_(str string) *JSString {
	return &JSString{float64(len(str)), str}
}

func (this *JSString) Replace(re *RegExp, f func(interface{}) interface{}) string {
	return re.compile().ReplaceAllStringFunc(this.String(),
		func(x string) string {
			return fmt.Sprint(f(x))
		})
}

func (this *JSString) Search(re *RegExp) float64 {
	match := re.compile().FindStringIndex(this.String())
	if match == nil {
		return -1
	}
	return float64(match[0])
}

func (this *JSString) CharAt(index float64) string {
	return string([]rune(this.String())[int(index)])
}

func (this *JSString) CharCodeAt(index float64) float64 {
	return float64([]rune(this.String())[int(index)])
}

func (this *JSString) Split(separator_limit ...interface{}) []interface{} {
	panic("Not implemented")
}

func (this *JSString) ToString() string {
	return this.str
}

func (this *JSString) String() string {
	return this.ToString()
}

func (this JSNumber) ToString() string {
	return fmt.Sprint(float64(this))
}

func (this JSNumber) String() string {
	return this.ToString()
}

func (this JSBoolean) ToString() string {
	return fmt.Sprint(bool(this))
}

func (this JSBoolean) String() string {
	return this.ToString()
}

func (this JSNil) ToString() string {
	return ""
}

func (this JSNil) String() string {
	return this.ToString()
}

type JSArray struct {
	Length float64
	array  interface{}
}

func JSArray_(a interface{}) *JSArray {
	var arr []interface{}
	switch a := a.(type) {
	case *[]interface{}:
		arr = *a
	case *interface{}:
		arr = (*a).([]interface{})
	default:
		panic("Unknown array type: " + fmt.Sprint(a))
	}
	return &JSArray{float64(len(arr)), a}
}

func (this *JSArray) arr() []interface{} {
	switch a := this.array.(type) {
	case *[]interface{}:
		return *a
	case *interface{}:
		return (*a).([]interface{})
	default:
		panic("Unknown array type: " + fmt.Sprint(a))
	}
}

func (this *JSArray) setArr(arr []interface{}) *JSArray {
	switch a := this.array.(type) {
	case *[]interface{}:
		*a = arr
	case *interface{}:
		*a = arr
	}
	this.Length = float64(len(arr))
	return this
}

// Untyped arguments, only used by ObjMap dissoc. We need to choose how to deal with this in general.
func (this *JSArray) Splice(index_, howMany_ interface{}, elements ...interface{}) []interface{} {
	index, howMany := index_.(float64), howMany_.(float64)
	if index < 0 {
		index = this.Length + index
	}
	removed := make([]interface{}, int(howMany))
	arr := this.arr()
	copy(removed, arr[int(index):int(index+howMany)])
	this.setArr(append(arr[:int(index)], append(elements, arr[int(index+howMany):]...)...))
	return removed
}

func (this *JSArray) Slice(index, end float64) []interface{} {
	arr := this.arr()
	if end < 0 {
		end = float64(len(arr)) + end
	}
	return arr[int(index):int(end)]
}

func (this *JSArray) Pop() interface{} {
	arr := this.arr()
	idx := len(arr) - 1
	this.setArr(arr[:idx])
	return arr[idx]
}

func (this *JSArray) Push(x interface{}) float64 {
	return this.setArr(append(this.arr(), x)).Length
}

func (this *JSArray) ToString() string {
	return fmt.Sprint(this.arr())
}

func (this *JSArray) String() string {
	return this.ToString()
}
