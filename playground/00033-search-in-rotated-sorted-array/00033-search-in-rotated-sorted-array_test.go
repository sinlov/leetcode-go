package _00033_search_in_rotated_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	// mock search
	qs := []question{
		{
			param{
				one: []int{4, 2, 1, 0, 9, 5, 6},
				two: 1,
			},
			answer{-1},
		},
		{
			param{
				one: []int{4, 2, 1, 0, 9, 5, 6},
				two: 0,
			},
			answer{3},
		},
		{
			param{
				one: []int{4, 2, 1, 0, 5, 7, 8},
				two: 7,
			},
			answer{5},
		},
		{
			param{
				one: []int{},
				two: 0,
			},
			answer{-1},
		},
		{
			param{
				one: []int{1},
				two: 0,
			},
			answer{-1},
		},
		{
			param{
				one: []int{4, 5, 6, 7, 0, 1, 2},
				two: 3,
			},
			answer{-1},
		},
		{
			param{
				one: []int{4, 5, 6, 7, 0, 1, 2},
				two: 0,
			},
			answer{4},
		},
	}
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
	one int
}
