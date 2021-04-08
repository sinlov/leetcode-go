package _00097_interleaving_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock isInterleave

var qs = []question{
	{
		param{
			one:   "",
			two:   "a",
			three: "b",
		},
		answer{false},
	},
	{
		param{
			one:   "",
			two:   "a",
			three: "a",
		},
		answer{true},
	},
	{
		param{
			one:   "b",
			two:   "",
			three: "b",
		},
		answer{true},
	},
	{
		param{
			one:   "aabcc",
			two:   "bdbca",
			three: "aabdbcbcac",
		},
		answer{true},
	},
	{
		param{
			one:   "aabcc",
			two:   "dbbca",
			three: "aadbbcbcac",
		},
		answer{true},
	},
	{
		param{
			one:   "aabc",
			two:   "dbbca",
			three: "aadbbcbcac",
		},
		answer{false},
	},
}

func Test_isInterleave(t *testing.T) {
	t.Logf("~> LeetCode isInterleave start")
	// do isInterleave
	for _, q := range qs {
		a, p := q.answer, q.param
		res := isInterleave(p.one, p.two, p.three)
		// verify isInterleave
		assert.Equal(t, a.one, res,
			"fail: in [ %v %v %v ], out [%v] , want [ %v ]", p.one, p.two, p.three, res, a.one)
	}
	t.Logf("~> LeetCode isInterleave end")
}

func Benchmark_isInterleave(b *testing.B) {
	// do isInterleave
	for _, q := range qs {
		isInterleave(q.param.one, q.param.two, q.param.three)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one   string
	two   string
	three string
}

// one first  answer
type answer struct {
	one bool
}
