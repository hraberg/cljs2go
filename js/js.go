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

func (e Error) Error() string {
	return fmt.Sprint(e.Message)
}

func (e TypeError) Error() string {
	return fmt.Sprint(e.Message)
}

type Object interface {
	ToString() string
	Equiv(other interface{}) bool
}
type Boolean bool
type Array []interface{}

type Date struct {
	Millis float64
}

func (this Date) time() time.Time {
	return time.Unix(int64(this.Millis)/1000, 1000*(int64(this.Millis)%1000))
}

func (this Date) GetUTCFullYear() float64 {
	return float64(this.time().UTC().Year())
}

func (this Date) GetUTCMonth() float64 {
	return float64(this.time().UTC().Month() - 1)
}

func (this Date) GetUTCDate() float64 {
	return float64(this.time().UTC().Day())
}

func (this Date) GetUTCHours() float64 {
	return float64(this.time().UTC().Hour())
}

func (this Date) GetUTCMinutes() float64 {
	return float64(this.time().UTC().Minute())
}

func (this Date) GetUTCSeconds() float64 {
	return float64(this.time().UTC().Second())
}

func (this Date) GetUTCMilliseconds() float64 {
	return float64(this.time().UTC().Nanosecond() / 1000)
}

func (this Date) String() string {
	return this.time().String()
}

func (this Date) ToString() string {
	return this.String()
}

type RegExp struct {
	Pattern string
	Flags   string
}

func (this RegExp) compile() *regexp.Regexp {
	pattern := this.Pattern
	if len(this.Flags) != 0 {
		pattern = "(?" + this.Flags + ")" + pattern
	}
	return regexp.MustCompile(pattern)
}

func (this RegExp) Exec(str string) []string {
	return this.compile().FindAllString(str, -1)
}

func (this RegExp) String() string {
	return this.compile().String()
}

var Number = struct{ MAX_VALUE float64 }{math.MaxFloat64}

var Infinity = math.Inf(1)

func IsNAN(x float64) bool {
	return math.IsNaN(x)
}

func ParseFloat(string string) float64 {
	if val, ok := strconv.ParseFloat(string, 64); ok == nil {
		return val
	}
	return math.NaN()
}

func ParseInt(string string, radix float64) float64 {
	if val, ok := strconv.ParseInt(string, int(radix), 64); ok == nil {
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
	FromCharCode func(...interface{}) interface{}
}{func(num ...interface{}) interface{} {
	var buffer bytes.Buffer
	for _, n := range num {
		buffer.WriteRune(rune(n.(int)))
	}
	return buffer.String()
}}

type JSString string

func (this JSString) Replace(re RegExp, f func(string) string) string {
	return re.compile().ReplaceAllStringFunc(string(this), f)
}

func (this JSString) Search(re RegExp) float64 {
	match := re.compile().FindStringIndex(string(this))
	if match == nil {
		return -1
	}
	return float64(match[0])
}

func (this JSString) CharAt(index float64) string {
	return string([]rune(string(this))[int(index)])
}

func (this JSString) CharCodeAt(index float64) float64 {
	return float64([]rune(string(this))[int(index)])
}

func (this JSString) String() string {
	return string(this)
}
