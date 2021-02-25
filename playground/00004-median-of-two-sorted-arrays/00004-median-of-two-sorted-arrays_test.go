package _00004_median_of_two_sorted_arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock findMedianSortedArrays

var qs = []question{
	{
		param{
			one: []int{7, 9, 10},
			two: []int{1, 3, 4},
		},
		answer{5.5000},
	},
	{
		param{
			one: []int{7, 9, 10},
			two: []int{1, 3, 4, 6},
		},
		answer{6.0000},
	},
	{
		param{
			one: []int{1, 3},
			two: []int{2},
		},
		answer{2.0000},
	},
	{
		param{
			one: []int{1, 2},
			two: []int{3, 4},
		},
		answer{2.5000},
	},
	{
		param{
			one: []int{1, 1},
			two: []int{1, 2},
		},
		answer{1.00000},
	},
	{
		param{
			one: []int{0, 0},
			two: []int{0, 0},
		},
		answer{0.00000},
	},
	{
		param{
			one: []int{},
			two: []int{1},
		},
		answer{1.00000},
	},
	{
		param{
			one: []int{2},
			two: []int{},
		},
		answer{2.00000},
	},
}

func Test_findMedianSortedArrays(t *testing.T) {
	t.Logf("~> LeetCode findMedianSortedArrays start")
	// do findMedianSortedArrays
	for _, q := range qs {
		a, p := q.answer, q.param
		res := findMedianSortedArrays(p.one, p.two)
		// verify findMedianSortedArrays
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode findMedianSortedArrays end")
}

func Benchmark_findMedianSortedArrays(b *testing.B) {
	// do findMedianSortedArrays
	for _, q := range qs {
		findMedianSortedArrays(q.param.one, q.param.two)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one []int
	two []int
}

// one first  answer
type answer struct {
	one float64
}
