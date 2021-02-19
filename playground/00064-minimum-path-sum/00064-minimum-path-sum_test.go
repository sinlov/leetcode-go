package _00064_minimum_path_sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock minPathSum
var qs = []question{
	{
		param{[][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}},
		answer{7},
	},
	{
		param{[][]int{
			{1, 2, 3},
			{4, 5, 6},
		}},
		answer{12},
	},
	{
		param{[][]int{
			{},
		}},
		answer{0},
	},
	{
		param{[][]int{}},
		answer{0},
	},
}

func Test_minPathSum(t *testing.T) {
	t.Logf("~> LeetCode minPathSum start")
	// do minPathSum
	for _, q := range qs {
		a, p := q.answer, q.param
		res := minPathSum(p.one)
		// verify minPathSum
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode minPathSum end")
}

func Benchmark_minPathSum(b *testing.B) {
	// do minPathSum
	for _, q := range qs {
		minPathSum(q.param.one)
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
	one int
}
