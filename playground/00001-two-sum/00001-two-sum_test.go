package _0001_two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	// mock twoSum
	qs := []question{
		{
			param{[]int{3, 2, 4}, 6},
			answer{[]int{1, 2}},
		},

		{
			param{[]int{3, 2, 4}, 5},
			answer{[]int{0, 1}},
		},

		{
			param{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			answer{[]int{1, 3}},
		},

		{
			param{[]int{0, 1}, 1},
			answer{[]int{0, 1}},
		},

		{
			param{[]int{0, 3}, 5},
			answer{[]int(nil)},
		},
		// 如需多个测试，可以复制上方元素。
	}
	t.Logf("~> LeetCode twoSum start")
	// do twoSum
	for _, q := range qs {
		a, p := q.answer, q.param
		// verify twoSum
		assert.Equal(t, a.one, twoSum(p.nums, p.target),
			"in [ %v ], out [%v]", p, twoSum(p.nums, p.target))
	}
	t.Logf("~> LeetCode twoSum end")
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	nums   []int
	target int
}

// one first  answer
type answer struct {
	one []int
}
