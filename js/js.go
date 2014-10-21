package js

import (
	"bytes"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
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

var Undefined interface{}

type Date struct {
	Millis interface{}
}

func (this *Date) time() time.Time {
	switch d := this.Millis.(type) {
	case time.Time:
		return d
	case string:
		for _, f := range []string{
			"2006-01-02T15:04:05.000-07:00",
			"2006-01-02T15:04:05.000",
			"2006-01-02T15:04:05",
			"2006-01-02T15:04",
			"2006-01-02T15",
			"2006-01-02",
			"2006-01",
			"2006"} {
			if t, ok := time.Parse(f, d); ok == nil {
				this.Millis = t
				break
			}
		}
	case float64:
		this.Millis = time.Unix(int64(d)/1000, 1000*1000*(int64(d)%1000))
	case int:
		this.Millis = time.Unix(int64(d)/1000, 1000*1000*(int64(d)%1000))
	case nil:
		this.Millis = time.Now()
	}
	if d, ok := this.Millis.(time.Time); ok {
		return d
	}
	panic(&Error{fmt.Sprint("Unknown date type: %v", this.Millis)})
}

func (this *Date) GetTime() float64 {
	return float64(this.time().UTC().UnixNano() / (1000 * 1000))
}

func (this *Date) ValueOf() float64 {
	return this.GetTime()
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
	return float64(this.time().UTC().Nanosecond() / (1000 * 1000))
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

var unicodeEscapeRegExp = regexp.MustCompile("\\\\u[[:xdigit:]]+")

func (this *RegExp) pattern() string {
	if p, ok := this.Pattern.(string); ok {
		return unicodeEscapeRegExp.ReplaceAllStringFunc(p,
			func(s string) string {
				return String.FromCharCode(ParseInt(s[2:], 16))
			})
	}
	return ""
}

func (this *RegExp) flags() string {
	if f, ok := this.Flags.(string); ok {
		return f
	}
	return ""
}

func (this *RegExp) compile() *regexp.Regexp {
	pattern := this.pattern()
	flags := this.flags()
	if len(flags) != 0 {
		pattern = "(?" + strings.Replace(flags, "g", "", -1) + ")" + pattern
	}
	return regexp.MustCompile(pattern)
}

func (this *RegExp) Exec(str string) []interface{} {
	if match := this.compile().FindStringSubmatchIndex(str); match != nil {
		strs := make([]interface{}, len(match)/2)
		for i := range strs {
			idx := i * 2
			if match[idx] != -1 {
				strs[i] = str[match[idx]:match[idx+1]]
			}
		}
		return strs
	}
	return nil
}

func (this *RegExp) String() string {
	pattern := this.pattern()
	if pattern == "" {
		pattern = "(?:)"
	}
	return "/" + pattern + "/" + this.flags()
}

var Number = struct{ MAX_VALUE float64 }{math.MaxFloat64}

var Infinity = math.Inf(1)
var NaN = math.NaN()

func IsNaN(x float64) bool {
	return math.IsNaN(x)
}

func ParseFloat(str interface{}) float64 {
	if str, ok := str.(string); ok {
		if val, err := strconv.ParseFloat(str, 64); err == nil {
			return val
		}
	}
	return math.NaN()
}

func ParseInt(str interface{}, radix float64) float64 {
	if str, ok := str.(string); ok {
		if val, err := strconv.ParseInt(str, int(radix), 64); err == nil {
			return float64(val)
		}
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

func (this *JSString) Replace(match interface{}, replacement interface{}) string {
	var re *RegExp
	if exp, ok := match.(*RegExp); ok {
		re = exp
	} else {
		re = &RegExp{regexp.QuoteMeta(fmt.Sprint(match)), ``}
	}
	global := strings.Contains(re.Flags.(string), "g")
	var f func(string) string
	if rf, ok := replacement.(func(interface{}) interface{}); ok {
		f = func(x string) string {
			return fmt.Sprint(rf(x))
		}
	} else {
		f = func(_ string) string {
			return fmt.Sprint(replacement)
		}
	}
	if global {
		return re.compile().ReplaceAllStringFunc(this.String(), f)
	} else {
		s := this.String()
		loc := re.compile().FindStringIndex(s)
		return s[:loc[0]] + f(s[loc[0]:loc[1]]) + s[loc[1]:]
	}
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

func (this *JSString) ToUpperCase() string {
	return strings.ToUpper(this.String())
}

func (this *JSString) ToLowerCase() string {
	return strings.ToLower(this.String())
}

func (this *JSString) IndexOf(x_limit ...interface{}) float64 {
	x := x_limit[0]
	limit := len(this.String())
	if len(x_limit) > 1 {
		limit = int(x_limit[1].(float64)) - 1
	}
	return float64(strings.Index(this.String()[:limit], fmt.Sprint(x)))
}

func (this *JSString) Substring(indexA_indexB ...float64) string {
	indexA := indexA_indexB[0]
	indexB := float64(len(this.String()))
	if len(indexA_indexB) > 1 {
		indexB = indexA_indexB[1]
	}
	if indexA > indexB {
		indexA, indexB = indexB, indexA
	}
	if indexA < 0 {
		indexA = 0
	}
	if indexB < 0 {
		indexB = 0
	}
	return string(this.String()[int(indexA):int(indexB)])
}

func (this *JSString) Split(separator_limit ...interface{}) []interface{} {
	var parts []string
	argc := len(separator_limit)
	if argc == 0 {
		parts = strings.Split(this.String(), "")
	}
	if argc > 0 {
		switch separator := separator_limit[0].(type) {
		case string:
			parts = strings.Split(this.String(), separator)
		case *RegExp:
			parts = separator.compile().Split(this.String(), -1)
		default:
			parts = []string{this.String()}
		}
	}
	limit := -1
	if argc == 2 {
		limit = int(separator_limit[1].(float64))
	}
	strs := make([]interface{}, len(parts))
	for i, v := range parts {
		strs[i] = v
	}
	if limit < 0 {
		return strs
	}
	return strs[:limit]
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

func (this *JSArray) Reverse() []interface{} {
	arr := this.arr()
	end := len(arr) - 1
	for i, v := range arr {
		if i > end/2 {
			break
		}
		arr[end-i], arr[i] = v, arr[end-i]
	}
	return arr
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

func (this *JSArray) Join(separator ...interface{}) string {
	sep := ","
	if len(separator) > 0 {
		sep = fmt.Sprint(separator[0])
	}
	arr := this.arr()
	ss := make([]string, len(arr))
	for i, v := range arr {
		ss[i] = fmt.Sprint(v)
	}
	return strings.Join(ss, sep)
}

func (this *JSArray) ToString() string {
	return fmt.Sprint(this.arr())
}

func (this *JSArray) String() string {
	return this.ToString()
}

// Base Type Protocols
// There are other extend-type calls at start of core.cljs we aren't dealing with.

func (_ JSNil) CljsCoreICounted__() {}

func (___ JSNil) X_count_Arity1() float64 {
	return 0
}
