package _00040_combination_sum_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock combinationSum2

var qs = []question{
	{
		param{
			one: []int{2, 5, 2, 1, 2},
			two: 5,
		},
		answer{[][]int{
			{1, 2, 2},
			{5},
		}},
	},
	{
		param{
			one: []int{10, 1, 2, 7, 6, 1, 5},
			two: 8,
		},
		answer{[][]int{
			{1, 1, 6},
			{1, 2, 5},
			{1, 7},
			{2, 6},
		}},
	},
	{
		param{
			one: []int{},
			two: 0,
		},
		answer{[][]int{}},
	},
}

func Test_combinationSum2(t *testing.T) {
	t.Logf("~> LeetCode combinationSum2 start")
	// do combinationSum2
	for _, q := range qs {
		a, p := q.answer, q.param
		res := combinationSum2(p.one, p.two)
		// verify combinationSum2
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode combinationSum2 end")
}

func Benchmark_combinationSum2(b *testing.B) {
	// do combinationSum2
	for _, q := range qs {
		combinationSum2(q.param.one, q.param.two)
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
	one [][]int
}
