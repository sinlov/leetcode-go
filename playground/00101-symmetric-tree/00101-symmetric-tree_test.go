package _00101_symmetric_tree

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock isSymmetric

var qs = []question{
	{
		param{[]int{1, 2, 3}},
		answer{false},
	},
	{
		param{[]int{1}},
		answer{true},
	},
	{
		param{[]int{1, 2, 2, 3, 4, 4, 3}},
		answer{true},
	},
	{
		param{[]int{1, 2, 2, structures.NULL, 3, structures.NULL, 3}},
		answer{false},
	},
	{
		param{[]int{}},
		answer{true},
	},
}

func Test_isSymmetric(t *testing.T) {
	t.Logf("~> LeetCode isSymmetric start")
	// do isSymmetric
	for _, q := range qs {
		a, p := q.answer, q.param
		treeMid := structures.Ints2TreeNode(p.one)
		res := isSymmetric(treeMid)
		// verify isSymmetric
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode isSymmetric end")
}

func Benchmark_isSymmetric(b *testing.B) {
	// do isSymmetric
	for _, q := range qs {
		isSymmetric(structures.Ints2TreeNode(q.param.one))
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
	one bool
}
