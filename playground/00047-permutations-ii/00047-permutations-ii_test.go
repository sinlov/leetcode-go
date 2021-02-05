package _00047_permutations_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock permuteUnique

var qs = []question{
	{
		param{[]int{1, 1, 2}},
		answer{[][]int{
			{1, 1, 2},
			{1, 2, 1},
			{2, 1, 1},
		}},
	},
	{
		param{[]int{}},
		answer{[][]int{}},
	},
}

func Test_permuteUnique(t *testing.T) {
	t.Logf("~> LeetCode permuteUnique start")
	// do permuteUnique
	for _, q := range qs {
		a, p := q.answer, q.param
		res := permuteUnique(p.one)
		// verify permuteUnique
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode permuteUnique end")
}

func Benchmark_permuteUnique(b *testing.B) {
	// do permuteUnique
	for _, q := range qs {
		permuteUnique(q.param.one)
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
	one [][]int
}
