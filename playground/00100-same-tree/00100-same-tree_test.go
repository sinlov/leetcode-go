package _00100_same_tree

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock isSameTree

var qs = []question{
	{
		param{
			one: []int{1, 2, 1},
			two: []int{1, 1, 2},
		},
		answer{false},
	},
	{
		param{
			one: []int{1, 2},
			two: []int{1, structures.NULL, 2},
		},
		answer{false},
	},
	{
		param{
			one: []int{1, 2, 3},
			two: []int{1, 2, 3},
		},
		answer{true},
	},
	{
		param{
			one: []int{1},
			two: []int{},
		},
		answer{false},
	},
	{
		param{
			one: []int{},
			two: []int{},
		},
		answer{true},
	},
}

func Test_isSameTree(t *testing.T) {
	t.Logf("~> LeetCode isSameTree start")
	// do isSameTree
	for _, q := range qs {
		a, p := q.answer, q.param
		pTree := structures.Ints2TreeNode(p.one)
		qTree := structures.Ints2TreeNode(p.two)
		res := isSameTree(pTree, qTree)
		// verify isSameTree
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode isSameTree end")
}

func Benchmark_isSameTree(b *testing.B) {
	// do isSameTree
	for _, q := range qs {
		isSameTree(
			structures.Ints2TreeNode(q.param.one),
			structures.Ints2TreeNode(q.param.two))
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one []int
	two []int
}

// one first  answer
type answer struct {
	one bool
}
