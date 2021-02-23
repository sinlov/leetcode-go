package _00070_climbing_stairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock climbStairs

var qs = []question{
	{
		param{4},
		answer{5},
	},
	{
		param{3},
		answer{3},
	},
	{
		param{2},
		answer{2},
	},
	{
		param{1},
		answer{1},
	},
}

func Test_climbStairs(t *testing.T) {
	t.Logf("~> LeetCode climbStairs start")
	// do climbStairs
	for _, q := range qs {
		a, p := q.answer, q.param
		res := climbStairs(p.one)
		// verify climbStairs
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode climbStairs end")
}

func Benchmark_climbStairs(b *testing.B) {
	// do climbStairs
	for _, q := range qs {
		climbStairs(q.param.one)
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
	one int
}
