package _00096_unique_binary_search_trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock numTrees

var qs = []question{
	{
		param{5},
		answer{42},
	},
	{
		param{3},
		answer{5},
	},
	{
		param{1},
		answer{1},
	},
}

func Test_numTrees(t *testing.T) {
	t.Logf("~> LeetCode numTrees start")
	// do numTrees
	for _, q := range qs {
		a, p := q.answer, q.param
		res := numTrees(p.one)
		// verify numTrees
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode numTrees end")
}

func Benchmark_numTrees(b *testing.B) {
	// do numTrees
	for _, q := range qs {
		numTrees(q.param.one)
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
	one int
}
