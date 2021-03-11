package _00083_remove_duplicates_from_sorted_list

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock deleteDuplicates

var qs = []question{
	{
		param{[]int{1, 1, 2}},
		answer{[]int{1, 2}},
	},
	{
		param{[]int{1, 1, 2, 3, 3}},
		answer{[]int{1, 2, 3}},
	},
	{
		param{[]int{1}},
		answer{[]int{1}},
	},
	{
		param{[]int{}},
		answer{[]int{}},
	},
}

func Test_deleteDuplicates(t *testing.T) {
	t.Logf("~> LeetCode deleteDuplicates start")
	// do deleteDuplicates
	for _, q := range qs {
		a, p := q.answer, q.param
		mid := structures.IntList2ListNode(p.one)
		resLN := deleteDuplicates(mid)
		res := structures.List2IntList(resLN)
		// verify deleteDuplicates
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode deleteDuplicates end")
}

func Benchmark_deleteDuplicates(b *testing.B) {
	// do deleteDuplicates
	for _, q := range qs {
		deleteDuplicates((structures.IntList2ListNode(q.param.one)))
	}
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
