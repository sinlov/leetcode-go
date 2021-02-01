package _00039_combination_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock combinationSum

var qs = []question{
	{
		param{
			one: []int{2, 3, 4, 6, 7, 8},
			two: 9,
		},
		answer{[][]int{
			{2, 2, 2, 3},
			{2, 3, 4},
			{2, 7},
			{3, 3, 3},
			{3, 6},
		}},
	},
	{
		param{
			one: []int{2, 3, 6, 7},
			two: 7,
		},
		answer{[][]int{
			{2, 2, 3},
			{7},
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

func Test_combinationSum(t *testing.T) {
	t.Logf("~> LeetCode combinationSum start")
	// do combinationSum
	for _, q := range qs {
		a, p := q.answer, q.param
		res := combinationSum(p.one, p.two)
		// verify combinationSum
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode combinationSum end")
}

func Benchmark_combinationSum(b *testing.B) {
	// do combinationSum
	for _, q := range qs {
		combinationSum(q.param.one, q.param.two)
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
