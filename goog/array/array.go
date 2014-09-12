package array

import (
	"fmt"
	"math/rand"
	"sort"
)

//func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }

func DefaultCompare(x, y interface{}) interface{} {
	switch t := y.(type) {
	case string:
		if y.(string) > x.(string) {
			return -1.0
		} else if y.(string) == x.(string) {
			return 0.0
		} else {
			return 1.0
		}
	case float64:
		return x.(float64) - y.(float64)
	default:
		panic(fmt.Sprintf("Cannot compare %v", t))
	}
}

type JSComparator struct {
	a    []interface{}
	comp func(a, b interface{}) interface{}
}

func (this JSComparator) Len() int           { return len(this.a) }
func (this JSComparator) Swap(i, j int)      { this.a[i], this.a[j] = this.a[j], this.a[i] }
func (this JSComparator) Less(i, j int) bool { return this.comp(this.a[i], this.a[j]).(float64) < 0 }

func StableSort(a []interface{}, comp func(a, b interface{}) interface{}) interface{} {
	sort.Sort(JSComparator{a, comp})
	return nil
}

func Shuffle(a []interface{}) interface{} {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return nil
}
