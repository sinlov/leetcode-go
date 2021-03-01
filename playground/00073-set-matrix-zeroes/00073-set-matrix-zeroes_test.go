package _00073_set_matrix_zeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock setZeroes

var qs = []question{
	{
		param{[][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}},
		answer{[][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		}},
	},
	{
		param{[][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}},
		answer{[][]int{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		}},
	},
}

func Test_setZeroes(t *testing.T) {
	t.Logf("~> LeetCode setZeroes start")
	// do setZeroes
	for _, q := range qs {
		a, p := q.answer, q.param
		setZeroes(p.one)
		// verify setZeroes
		assert.Equal(t, a.one, p.one,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, p.one, a.one)
	}
	t.Logf("~> LeetCode setZeroes end")
}

func Benchmark_setZeroes(b *testing.B) {
	// do setZeroes
	for _, q := range qs {
		setZeroes(q.param.one)
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
