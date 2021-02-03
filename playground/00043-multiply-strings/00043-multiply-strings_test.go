package _00043_multiply_strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock multiply

var qs = []question{
	{
		param{
			one: "123",
			two: "456",
		},
		answer{"56088"},
	},
	{
		param{
			one: "3",
			two: "10",
		},
		answer{"30"},
	},
	{
		param{
			one: "2",
			two: "3",
		},
		answer{"6"},
	},
	{
		param{
			one: "0",
			two: "3",
		},
		answer{"0"},
	},
}

func Test_multiply(t *testing.T) {
	t.Logf("~> LeetCode multiply start")
	// do multiply
	for _, q := range qs {
		a, p := q.answer, q.param
		res := multiply(p.one, p.two)
		// verify multiply
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode multiply end")
}

func Benchmark_multiply(b *testing.B) {
	// do multiply
	for _, q := range qs {
		multiply(q.param.one, q.param.two)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one string
	two string
}

// one first  answer
type answer struct {
	one string
}
