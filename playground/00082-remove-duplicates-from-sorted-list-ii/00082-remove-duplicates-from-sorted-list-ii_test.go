package _00082_remove_duplicates_from_sorted_list_ii

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock deleteDuplicates

var qs = []question{
	{
		param{[]int{1, 2, 3, 3, 4, 4, 5}},
		answer{[]int{1, 2, 5}},
	},
	{
		param{[]int{1, 1, 1, 2, 3}},
		answer{[]int{2, 3}},
	},
	{
		param{[]int{}},
		answer{[]int{}},
	},
	{
		param{[]int{1, 1, 1}},
		answer{[]int{}},
	},
}

func Test_deleteDuplicates(t *testing.T) {
	t.Logf("~> LeetCode deleteDuplicates start")
	// do deleteDuplicates
	for _, q := range qs {
		a, p := q.answer, q.param
		inLN := structures.IntList2ListNode(p.one)
		resMid := deleteDuplicates(inLN)
		res := structures.List2IntList(resMid)
		// verify deleteDuplicates
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode deleteDuplicates end")
}

func Benchmark_deleteDuplicates(b *testing.B) {
	// do deleteDuplicates
	for _, q := range qs {
		deleteDuplicates(structures.IntList2ListNode(q.param.one))
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
