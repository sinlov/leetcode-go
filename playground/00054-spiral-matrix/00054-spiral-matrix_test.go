package _00054_spiral_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock spiralOrder

var qs = []question{
	{
		param{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}},
		answer{[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	},
	{
		param{[][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}},
		answer{[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	},
	{
		param{[][]int{
			{},
		}},
		answer{[]int{}},
	},
	{
		param{[][]int{}},
		answer{[]int{}},
	},
}

func Test_spiralOrder(t *testing.T) {
	t.Logf("~> LeetCode spiralOrder start")
	// do spiralOrder
	for _, q := range qs {
		a, p := q.answer, q.param
		res := spiralOrder(p.one)
		// verify spiralOrder
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode spiralOrder end")
}

func Benchmark_spiralOrder(b *testing.B) {
	// do spiralOrder
	for _, q := range qs {
		spiralOrder(q.param.one)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one [][]int
}

// one first  answer
type answer struct {
	one []int
}
