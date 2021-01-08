package _00026_remove_duplicates_from_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	// mock removeDuplicates
	qs := []question{
		{
			param{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			answer{5},
		},
		{
			param{[]int{1, 1, 2}},
			answer{2},
		},
		{
			param{[]int{0, 0, 0, 0}},
			answer{1},
		},
		{
			param{[]int{1}},
			answer{1},
		},
		{
			param{[]int{}},
			answer{0},
		},
	}
	t.Logf("~> LeetCode removeDuplicates start")
	// do removeDuplicates
	for _, q := range qs {
		a, p := q.answer, q.param
		res := removeDuplicates(p.one)
		// verify removeDuplicates
		assert.Equal(t, a.one, res,
			"in [ %v ], out [%v]", p.one, res)
	}
	t.Logf("~> LeetCode removeDuplicates end")
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one []int
}

// one first  answer
type answer struct {
	one int
}
