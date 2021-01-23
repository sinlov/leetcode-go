package _00024_swap_nodes_in_pairs

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_swapPairs(t *testing.T) {
	// mock swapPairs
	qs := []question{
		{
			param{[]int{1}},
			answer{[]int{1}},
		},
		{
			param{[]int{}},
			answer{[]int{}},
		},
		{
			param{[]int{1, 2, 3, 4}},
			answer{[]int{2, 1, 4, 3}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}},
			answer{[]int{2, 1, 4, 3, 5}},
		},
	}
	t.Logf("~> LeetCode swapPairs start")
	// do swapPairs
	for _, q := range qs {
		a, p := q.answer, q.param
		res := swapPairs(structures.IntList2ListNode(p.one))
		resIntArr := structures.List2IntList(res)
		// verify swapPairs
		assert.Equal(t, a.one, resIntArr,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, resIntArr, a.one)
	}
	t.Logf("~> LeetCode swapPairs end")
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
