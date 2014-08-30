package main

import "testing"

import "github.com/stretchr/testify/assert"

import . "github.com/hraberg/cljs.go/cljs/core"

func Benchmark_RecursiveDirectCall(t *testing.B) {
	fib := func(this *AFn) *AFn {
		return Fn(this, func(n interface{}) interface{} {
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.X_invoke_Arity1(n.(float64)-1.0).(float64) +
					this.X_invoke_Arity1(n.(float64)-2.0).(float64)
			}
		})
	}(&AFn{})
	assert.Equal(t, 832040, fib.X_invoke_Arity1(30.0))
}

func Benchmark_RecursiveDirectPrimitiveCall(t *testing.B) {
	fib := func(this *AFn) *AFn {
		return Fn(this, func(n float64) float64 {
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.Arity1FF(n-1.0) + this.Arity1FF(n-2.0)
			}
		})
	}(&AFn{})
	assert.Equal(t, 832040, fib.X_invoke_Arity1(30.0))
}

func Benchmark_RecursiveDispatch(t *testing.B) {
	fib := func(this *AFn) *AFn {
		return Fn(this, func(a interface{}) interface{} {
			var n = a.(float64)
			if n == 0.0 {
				return 0.0
			} else if n == 1.0 {
				return 1.0
			} else {
				return this.Call(n-1.0).(float64) + this.Call(n-2.0).(float64)
			}
		})
	}(&AFn{})
	assert.Equal(t, 832040, fib.Call(30.0))
}

func Benchmark_RecursiveGo(t *testing.B) {
	fib := func() func(float64) float64 {
		var this func(float64) float64
		this = func(n float64) float64 {
			if n == 0 {
				return 0
			} else if n == 1 {
				return 1
			} else {
				return this(n-1) + this(n-2)
			}
		}
		return this
	}()
	assert.Equal(t, 832040, fib(30))
}
