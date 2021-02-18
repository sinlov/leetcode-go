package _00062_unique_paths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock uniquePaths

var qs = []question{
	{
		param{
			one: 3,
			two: 2,
		},
		answer{3},
	},
	{
		param{
			one: 3,
			two: 3,
		},
		answer{6},
	},
	{
		param{
			one: 3,
			two: 7,
		},
		answer{28},
	},
}

func Test_uniquePaths(t *testing.T) {
	t.Logf("~> LeetCode uniquePaths start")
	// do uniquePaths
	for _, q := range qs {
		a, p := q.answer, q.param
		res := uniquePaths(p.one, p.two)
		// verify uniquePaths
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode uniquePaths end")
}

func Benchmark_uniquePaths(b *testing.B) {
	// do uniquePaths
	for _, q := range qs {
		uniquePaths(q.param.one, q.param.two)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one int
	two int
}

// one first  answer
type answer struct {
	one int
}
