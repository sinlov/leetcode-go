package _00067_add_binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock addBinary

var qs = []question{
	{
		param{
			one: "1",
			two: "11",
		},
		answer{"100"},
	},
	{
		param{
			one: "1111",
			two: "111",
		},
		answer{"10110"},
	},
	{
		param{
			one: "1010",
			two: "1011",
		},
		answer{"10101"},
	},
}

func Test_addBinary(t *testing.T) {
	t.Logf("~> LeetCode addBinary start")
	// do addBinary
	for _, q := range qs {
		a, p := q.answer, q.param
		res := addBinary(p.one, p.two)
		// verify addBinary
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode addBinary end")
}

func Benchmark_addBinary(b *testing.B) {
	// do addBinary
	for _, q := range qs {
		addBinary(q.param.one, q.param.two)
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
