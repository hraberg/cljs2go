package js

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
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

var Undefined interface{} = nil

type Object interface {
	ToString() JSString
	Equiv(other interface{}) bool
}

type JSObject map[JSString]interface{}

type JSBoolean bool
type JSArray []interface{}

var Array = struct {
	IsArray func(interface{}) bool
}{func(x interface{}) bool {
	return reflect.TypeOf(x).Kind() == reflect.Slice
}}

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

func (this *RegExp) Exec(str JSString) []JSString {
	if match := this.compile().FindAllString(string(str), -1); match != nil {
		strs := make([]JSString, len(match))
		for i, v := range match {
			strs[i] = JSString(v)
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

func IsNAN(x float64) bool {
	return math.IsNaN(x)
}

func ParseFloat(str JSString) float64 {
	if val, ok := strconv.ParseFloat(string(str), 64); ok == nil {
		return val
	}
	return math.NaN()
}

func ParseInt(str JSString, radix float64) float64 {
	if val, ok := strconv.ParseInt(string(str), int(radix), 64); ok == nil {
		return float64(val)
	}
	return math.NaN()
}

var Console = struct {
	Log func(...interface{}) interface{}
}{func(obj ...interface{}) interface{} {
	fmt.Println(obj...)
	return nil
}}

var String = struct {
	FromCharCode func(...float64) JSString
}{func(num ...float64) JSString {
	var buffer bytes.Buffer
	for _, n := range num {
		buffer.WriteRune(rune(int(n)))
	}
	return JSString(buffer.String())
}}

type JSString string

func (this JSString) Replace(re *RegExp, f func(interface{}) interface{}) JSString {
	return JSString(re.compile().ReplaceAllStringFunc(this.String(),
		func(x string) string {
			return fmt.Sprint(f(JSString(x)))
		}))
}

func (this JSString) Search(re *RegExp) float64 {
	match := re.compile().FindStringIndex(this.String())
	if match == nil {
		return -1
	}
	return float64(match[0])
}

func (this JSString) CharAt(index float64) JSString {
	return JSString([]rune(this.String())[int(index)])
}

func (this JSString) CharCodeAt(index float64) float64 {
	return float64([]rune(this.String())[int(index)])
}

func (this JSString) String() string {
	return string(this)
}
