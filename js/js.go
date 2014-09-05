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
	Pattern string
	Flags   string
}

func (this *RegExp) compile() *regexp.Regexp {
	pattern := this.Pattern
	if len(this.Flags) != 0 {
		pattern = "(?" + this.Flags + ")" + pattern
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
	array  *[]interface{}
}

func JSArray_(a *[]interface{}) *JSArray {
	return &JSArray{float64(len(*a)), a}
}

// Untyped arguments, only used by ObjMap dissoc. We need to choose how to deal with this in general.
func (this *JSArray) Splice(index_, howMany_ interface{}, elements ...interface{}) []interface{} {
	index, howMany := index_.(float64), howMany_.(float64)
	if index < 0 {
		index = this.Length + index
	}
	removed := make([]interface{}, int(howMany))
	copy(removed, (*this.array)[int(index):int(index+howMany)])
	*this.array = append((*this.array)[:int(index)], append(elements, (*this.array)[int(index+howMany):]...)...)
	this.Length = float64(len(*this.array))
	return removed
}

func (this *JSArray) Slice(index, end float64) []interface{} {
	if end < 0 {
		end = float64(len(*this.array)) + end
	}
	return (*this.array)[int(index):int(end)]
}

func (this *JSArray) Pop() interface{} {
	idx := len(*this.array) - 1
	x := (*this.array)[idx]
	*this.array = (*this.array)[:idx]
	this.Length -= 1
	return x
}

func (this *JSArray) Push(x interface{}) float64 {
	*this.array = append(*this.array, x)
	this.Length += 1
	return this.Length
}

func (this *JSArray) ToString() string {
	return fmt.Sprint(*this.array)
}

func (this *JSArray) String() string {
	return this.ToString()
}
