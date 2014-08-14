package Math

import (
	"math"
	"math/rand"
)

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Random() float64 {
	return rand.Float64()
}

func Imul(a, b float64) float64 {
	return float64(int64(a) * int64(b))
}
