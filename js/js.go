package js

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"time"
)

// This file contains a thin js runtime layer so ClojureScript itself can run with minimal modifications.

type Error struct {
	Message string
}

type TypeError struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func (e TypeError) Error() string {
	return e.Message
}

type Array []interface{}

type Date struct {
	Millis int64
}

func (this Date) time() time.Time {
	return time.Unix(this.Millis/1000, 1000*(this.Millis%1000))
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
	var pattern = this.Pattern
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

type NumberConstructor struct {
	MAX_VALUE float64
}

var Number = NumberConstructor{math.MaxFloat64}

type MathConstructor struct{}

var Math = MathConstructor{}

func (_ MathConstructor) Floor(x float64) float64 {
	return math.Floor(x)
}

func (_ MathConstructor) Ceil(x float64) float64 {
	return math.Ceil(x)
}

func (_ MathConstructor) Random() float64 {
	return rand.Float64()
}

func (_ MathConstructor) Imul(a, b float64) float64 {
	return float64(int64(a) * int64(b))
}

var Infinity = math.Inf(1)

func IsNAN(x float64) bool {
	return math.IsNaN(x)
}

type ConsoleConstructor struct{}

func (_ ConsoleConstructor) Log(obj ...interface{}) interface{} {
	fmt.Println(obj...)
	return nil
}

var Console = ConsoleConstructor{}

type StringConstructor struct{}

func (_ StringConstructor) FromCharCode(num ...interface{}) interface{} {
	var buffer bytes.Buffer
	for _, n := range num {
		buffer.WriteRune(rune(n.(int)))
	}
	return buffer.String()
}

var String = StringConstructor{}

type JSString string

func (this JSString) Replace(re RegExp, f func(string) string) string {
	return re.compile().ReplaceAllStringFunc(string(this), f)
}

func (this JSString) Search(re RegExp) float64 {
	var match = re.compile().FindStringIndex(string(this))
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
