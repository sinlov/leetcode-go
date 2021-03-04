package _00078_subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock subsets

var qs = []question{
	{
		param{[]int{
			1, 2, 3,
		}},
		answer{[][]int{
			[]int(nil),
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		}},
	},
	{
		param{[]int{0}},
		answer{[][]int{
			[]int(nil),
			{0},
		}},
	},
}

func Test_subsets(t *testing.T) {
	t.Logf("~> LeetCode subsets start")
	// do subsets
	for _, q := range qs {
		a, p := q.answer, q.param
		res := subsets(p.one)
		// verify subsets
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode subsets end")
}

func Benchmark_subsets(b *testing.B) {
	// do subsets
	for _, q := range qs {
		subsets(q.param.one)
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
