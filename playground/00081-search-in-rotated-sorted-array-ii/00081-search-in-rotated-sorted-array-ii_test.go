package _00081_search_in_rotated_sorted_array_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock search

var qs = []question{
	{
		param{
			one: []int{4, 2, 1, 0, 9, 5, 6},
			two: 3,
		},
		answer{false},
	},
	{
		param{
			one: []int{4, 2, 1, 0, 9, 5, 6},
			two: 1,
		},
		answer{false},
	},
	{
		param{
			one: []int{4, 2, 1, 0, 9, 5, 6},
			two: 1,
		},
		answer{false},
	},
	{
		param{
			one: []int{6, 7, 9, 0, 1, 2, 3},
			two: 4,
		},
		answer{false},
	},
	{
		param{
			one: []int{2, 5, 6, 0, 0, 1, 2},
			two: 0,
		},
		answer{true},
	},
	{
		param{
			one: []int{2, 5, 6, 0, 0, 1, 2},
			two: 3,
		},
		answer{false},
	},
	{
		param{
			one: []int{},
			two: 3,
		},
		answer{false},
	},
}

func Test_search(t *testing.T) {
	t.Logf("~> LeetCode search start")
	// do search
	for _, q := range qs {
		a, p := q.answer, q.param
		res := search(p.one, p.two)
		// verify search
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode search end")
}

func Benchmark_search(b *testing.B) {
	// do search
	for _, q := range qs {
		search(q.param.one, q.param.two)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one []int
	two int
}

// one first  answer
type answer struct {
	one bool
}
