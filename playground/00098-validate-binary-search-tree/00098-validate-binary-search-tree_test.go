package _00098_validate_binary_search_tree

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock isValidBST

var qs = []question{
	{
		param{[]int{5, 1, 4, structures.NULL, structures.NULL, 3, 6}},
		answer{false},
	},
	{
		param{[]int{2, 1, 3}},
		answer{true},
	},
	{
		param{[]int{}},
		answer{true},
	},
}

func Test_isValidBST(t *testing.T) {
	t.Logf("~> LeetCode isValidBST start")
	// do isValidBST
	for _, q := range qs {
		a, p := q.answer, q.param
		tN := structures.Ints2TreeNode(p.one)
		res := isValidBST(tN)
		// verify isValidBST
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode isValidBST end")
}

func Benchmark_isValidBST(b *testing.B) {
	// do isValidBST
	for _, q := range qs {
		isValidBST(structures.Ints2TreeNode(q.param.one))
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
