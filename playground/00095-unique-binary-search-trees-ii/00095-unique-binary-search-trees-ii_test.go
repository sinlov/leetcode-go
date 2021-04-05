package _00095_unique_binary_search_trees_ii

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock generateTrees

var qs = []question{
	{
		param{one: 3},
		answer{[][]int{
			{1, structures.NULL, 2, structures.NULL, 3},
			{1, structures.NULL, 3, 2},
			{2, 1, 3},
			{3, 1, structures.NULL, structures.NULL, 2},
			{3, 2, structures.NULL, 1},
		}},
	},
	{
		param{one: 1},
		answer{[][]int{
			{1},
		}},
	},
	{
		param{one: 0},
		answer{[][]int(nil)},
	},
}

func Test_generateTrees(t *testing.T) {
	t.Logf("~> LeetCode generateTrees start")
	// do generateTrees
	for _, q := range qs {
		a, p := q.answer, q.param
		resTrees := generateTrees(p.one)
		var res [][]int
		for _, t := range resTrees {
			item := structures.Tree2ints(t)
			res = append(res, item)
		}
		// verify generateTrees
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode generateTrees end")
}

func Benchmark_generateTrees(b *testing.B) {
	// do generateTrees
	for _, q := range qs {
		generateTrees(q.param.one)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one int
}

// one first  answer
type answer struct {
	one [][]int
}
