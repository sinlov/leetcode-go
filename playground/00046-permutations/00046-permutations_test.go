package _00046_permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock permute

var qs = []question{
	{
		param{[]int{1, 2, 3}},
		answer{[][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
	},
	{
		param{[]int{}},
		answer{[][]int{}},
	},
}

func Test_permute(t *testing.T) {
	t.Logf("~> LeetCode permute start")
	// do permute
	for _, q := range qs {
		a, p := q.answer, q.param
		res := permute(p.one)
		// verify permute
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode permute end")
}

func Benchmark_permute(b *testing.B) {
	// do permute
	for _, q := range qs {
		permute(q.param.one)
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
