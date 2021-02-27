package _00025_reverse_nodes_in_k_group

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock reverseKGroup

var qs = []question{
	{
		param{
			one: []int{1, 2, 3, 4, 5},
			two: 2,
		},
		answer{[]int{2, 1, 4, 3, 5}},
	},
	{
		param{
			one: []int{1, 2, 3, 4, 5},
			two: 3,
		},
		answer{[]int{3, 2, 1, 4, 5}},
	},
}

func Test_reverseKGroup(t *testing.T) {
	t.Logf("~> LeetCode reverseKGroup start")
	// do reverseKGroup
	for _, q := range qs {
		a, p := q.answer, q.param

		lN := structures.IntList2ListNode(p.one)
		mid := reverseKGroup(lN, p.two)
		res := structures.List2IntList(mid)
		// verify reverseKGroup
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode reverseKGroup end")
}

func Benchmark_reverseKGroup(b *testing.B) {
	// do reverseKGroup
	for _, q := range qs {
		reverseKGroup(structures.IntList2ListNode(q.param.one), q.param.two)
	}
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
