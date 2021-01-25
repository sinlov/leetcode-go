package _00031_next_permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextPermutation(t *testing.T) {
	// mock nextPermutation
	qs := []question{
		{
			param{[]int{3, 5, 1, 4, 7, 6, 2}},
			answer{[]int{3, 5, 1, 6, 2, 4, 7}},
		},
		{
			param{[]int{1}},
			answer{[]int{1}},
		},
		{
			param{[]int{1, 1, 5}},
			answer{[]int{1, 5, 1}},
		},
		{
			param{[]int{3, 2, 1}},
			answer{[]int{1, 2, 3}},
		},
		{
			param{[]int{1, 2, 3}},
			answer{[]int{1, 3, 2}},
		},
	}
	t.Logf("~> LeetCode nextPermutation start")
	// do nextPermutation
	for _, q := range qs {
		a, p := q.answer, q.param
		nextPermutation(p.one)
		// verify nextPermutation
		assert.Equal(t, a.one, p.one,
			"fail: in after [%v] , want [ %v ]", p.one, a.one)
	}
	t.Logf("~> LeetCode nextPermutation end")
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
	one []int
}
