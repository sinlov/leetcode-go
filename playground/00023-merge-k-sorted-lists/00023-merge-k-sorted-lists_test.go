package _00023_merge_k_sorted_lists

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock mergeKLists

var qs = []question{
	{
		param{[][]int{
			{2, 3, 4},
			{5, 6, 7},
			{1, 4, 5},
		},
		},
		answer{[]int{1, 2, 3, 4, 4, 5, 5, 6, 7}},
	},
	{
		param{[][]int{
			{5, 6, 7},
			{2, 3, 4},
		},
		},
		answer{[]int{2, 3, 4, 5, 6, 7}},
	},
	{
		param{[][]int{
			{1, 3, 8},
			{1, 7},
		},
		},
		answer{[]int{1, 1, 3, 7, 8}},
	},
	{
		param{[][]int{
			{1},
			{1},
		},
		},
		answer{[]int{1, 1}},
	},
	{
		param{[][]int{
			{1, 2, 3},
		},
		},
		answer{[]int{1, 2, 3}},
	},
	{
		param{[][]int{}},
		answer{[]int{}},
	},
}

func Test_mergeKLists(t *testing.T) {
	t.Logf("~> LeetCode mergeKLists start")
	// do mergeKLists
	for _, q := range qs {
		a, p := q.answer, q.param
		var ls []*ListNode
		for _, l := range p.one {
			ls = append(ls, structures.IntList2ListNode(l))
		}
		mid := mergeKLists(ls)
		res := structures.List2IntList(mid)
		// verify mergeKLists
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode mergeKLists end")
}

func Benchmark_mergeKLists(b *testing.B) {
	// do mergeKLists
	for _, q := range qs {
		var ls []*ListNode
		for _, l := range q.param.one {
			ls = append(ls, structures.IntList2ListNode(l))
		}
		mergeKLists(ls)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one [][]int
}

// one first  answer
type answer struct {
	one []int
}
