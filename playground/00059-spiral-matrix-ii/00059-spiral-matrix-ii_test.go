package _00059_spiral_matrix_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock generateMatrix

var qs = []question{
	{
		param{3},
		answer{[][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		}},
	},
	{
		param{1},
		answer{[][]int{
			{1},
		}},
	},
}

func Test_generateMatrix(t *testing.T) {
	t.Logf("~> LeetCode generateMatrix start")
	// do generateMatrix
	for _, q := range qs {
		a, p := q.answer, q.param
		res := generateMatrix(p.one)
		// verify generateMatrix
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode generateMatrix end")
}

func Benchmark_generateMatrix(b *testing.B) {
	// do generateMatrix
	for _, q := range qs {
		generateMatrix(q.param.one)
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
