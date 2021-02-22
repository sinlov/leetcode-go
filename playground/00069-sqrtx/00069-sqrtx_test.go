package _00069_sqrtx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock mySqrt

var qs = []question{
	{
		param{8},
		answer{2},
	},
	{
		param{9},
		answer{3},
	},
	{
		param{4},
		answer{2},
	},
	{
		param{1},
		answer{1},
	},
	{
		param{0},
		answer{0},
	},
}

func Test_mySqrt(t *testing.T) {
	t.Logf("~> LeetCode mySqrt start")
	// do mySqrt
	for _, q := range qs {
		a, p := q.answer, q.param
		res := mySqrt(p.one)
		// verify mySqrt
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode mySqrt end")
}

func Benchmark_mySqrt(b *testing.B) {
	// do mySqrt
	for _, q := range qs {
		mySqrt(q.param.one)
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
