package _00053_maximum_subarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock maxSubArray

var qs = []question{
	{
		param{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
		answer{6},
	},
	{
		param{[]int{1}},
		answer{1},
	},
	{
		param{[]int{0}},
		answer{0},
	},
	{
		param{[]int{}},
		answer{0},
	},
}

func Test_maxSubArray(t *testing.T) {
	t.Logf("~> LeetCode maxSubArray start")
	// do maxSubArray
	for _, q := range qs {
		a, p := q.answer, q.param
		res := maxSubArray(p.one)
		// verify maxSubArray
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode maxSubArray end")
}

func Benchmark_maxSubArray(b *testing.B) {
	// do maxSubArray
	for _, q := range qs {
		maxSubArray(q.param.one)
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
	one int
}
