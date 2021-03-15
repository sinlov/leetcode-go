package _00090_subsets_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock subsetsWithDup

var qs = []question{
	{
		param{[]int{1, 2, 2}},
		answer{[][]int{
			{},
			{1},
			{2},
			{1, 2},
			{2, 2},
			{1, 2, 2},
		}},
	},
}

func Test_subsetsWithDup(t *testing.T) {
	t.Logf("~> LeetCode subsetsWithDup start")
	// do subsetsWithDup
	for _, q := range qs {
		a, p := q.answer, q.param
		res := subsetsWithDup(p.one)
		// verify subsetsWithDup
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode subsetsWithDup end")
}

func Benchmark_subsetsWithDup(b *testing.B) {
	// do subsetsWithDup
	for _, q := range qs {
		subsetsWithDup(q.param.one)
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
