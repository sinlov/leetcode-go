package _00056_merge_intervals

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock merge

var qs = []question{
	{
		param{[][]int{
			{4, 5},
			{2, 4},
			{4, 6},
			{3, 4},
			{0, 0},
			{1, 1},
			{3, 5},
			{2, 2},
		}},
		answer{[][]int{
			{0, 0},
			{1, 1},
			{2, 6},
		}},
	},
	{
		param{[][]int{
			{1, 4},
			{0, 4},
		}},
		answer{[][]int{
			{0, 4},
		}},
	},
	{
		param{[][]int{
			{1, 4},
			{4, 5},
		}},
		answer{[][]int{
			{1, 5},
		}},
	},
	{
		param{[][]int{
			{1, 3},
			{2, 6},
			{8, 10},
			{15, 18},
		}},
		answer{[][]int{
			{1, 6},
			{8, 10},
			{15, 18},
		}},
	},
	{
		param{[][]int{}},
		answer{[][]int{}},
	},
}

func Test_merge(t *testing.T) {
	t.Logf("~> LeetCode merge start")
	// do merge
	for _, q := range qs {
		a, p := q.answer, q.param
		res := merge(p.one)
		// verify merge
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode merge end")
}

func Benchmark_merge(b *testing.B) {
	// do merge
	for _, q := range qs {
		merge(q.param.one)
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
	one [][]int
}
