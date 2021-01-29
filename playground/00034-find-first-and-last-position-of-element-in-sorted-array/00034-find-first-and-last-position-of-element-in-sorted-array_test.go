package _00034_find_first_and_last_position_of_element_in_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	// mock searchRange
	qs := []question{
		{
			param{
				one: []int{1, 2, 3, 5, 7, 7, 7, 8, 9, 10, 11},
				two: 7,
			},
			answer{[]int{4, 6}},
		},
		{
			param{
				one: []int{1, 2, 3, 5, 7, 7, 8, 8, 9, 10, 11},
				two: 5,
			},
			answer{[]int{3, 3}},
		},
		{
			param{
				one: []int{1, 2, 3},
				two: 0,
			},
			answer{[]int{-1, -1}},
		},
		{
			param{
				one: []int{1, 2, 3},
				two: 4,
			},
			answer{[]int{-1, -1}},
		},
		{
			param{
				one: []int{},
				two: 0,
			},
			answer{[]int{-1, -1}},
		},
		{
			param{
				one: []int{5, 7, 7, 8, 8, 10},
				two: 8,
			},
			answer{[]int{3, 4}},
		},
	}
	t.Logf("~> LeetCode searchRange start")
	// do searchRange
	for _, q := range qs {
		a, p := q.answer, q.param
		res := searchRange(p.one, p.two)
		// verify searchRange
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode searchRange end")
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
	one []int
}
