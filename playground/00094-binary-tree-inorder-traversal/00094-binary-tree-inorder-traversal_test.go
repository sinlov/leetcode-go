package _00094_binary_tree_inorder_traversal

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock inorderTraversal

var qs = []question{
	{
		param{[]int{1, structures.NULL, 2, 3}},
		answer{[]int{1, 3, 2}},
	},
	{
		param{[]int{}},
		answer{[]int(nil)},
	},
	{
		param{[]int{1}},
		answer{[]int{1}},
	},
	{
		param{[]int{1, 2}},
		answer{[]int{2, 1}},
	},
	{
		param{[]int{1, structures.NULL, 2}},
		answer{[]int{1, 2}},
	},
}

func Test_inorderTraversal(t *testing.T) {
	t.Logf("~> LeetCode inorderTraversal start")
	// do inorderTraversal
	for _, q := range qs {
		a, p := q.answer, q.param
		treeMId := structures.Ints2TreeNode(p.one)
		res := inorderTraversal(treeMId)
		// verify inorderTraversal
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode inorderTraversal end")
}

func Benchmark_inorderTraversal(b *testing.B) {
	// do inorderTraversal
	for _, q := range qs {
		inorderTraversal(structures.Ints2TreeNode(q.param.one))
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
