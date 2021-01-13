package _00015_3sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSum(t *testing.T) {
	// mock threeSum
	qs := []question{
		{
			param{[]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}},
			answer{[][]int{
				{-4, -2, 6},
				{-2, -2, 4},
				{-4, 0, 4},
				{-2, 0, 2},
				{-4, 1, 3},
				{-4, 2, 2},
			}},
		},
		{
			param{[]int{-1, 0, 1, 2, -1, -4}},
			answer{[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			}},
		},
		{
			param{[]int{0, 0, 0}},
			answer{[][]int{{0, 0, 0}}},
		},
		{
			param{[]int{0}},
			answer{make([][]int, 0)},
		},
		{
			param{[]int{}},
			answer{make([][]int, 0)},
		},
	}
	t.Logf("~> LeetCode threeSum start")
	// do threeSum
	for _, q := range qs {
		a, p := q.answer, q.param
		res := threeSum(p.one)
		// verify threeSum
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v], want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode threeSum end")
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
	one [][]int
}
