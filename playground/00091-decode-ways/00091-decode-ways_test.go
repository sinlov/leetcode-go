package _00091_decode_ways

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock numDecodings

var qs = []question{
	{
		param{"12"},
		answer{2},
	},
	{
		param{"226"},
		answer{3},
	},
	{
		param{""},
		answer{0},
	},
	{
		param{"06"},
		answer{0},
	},
}

func Test_numDecodings(t *testing.T) {
	t.Logf("~> LeetCode numDecodings start")
	// do numDecodings
	for _, q := range qs {
		a, p := q.answer, q.param
		res := numDecodings(p.one)
		// verify numDecodings
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode numDecodings end")
}

func Benchmark_numDecodings(b *testing.B) {
	// do numDecodings
	for _, q := range qs {
		numDecodings(q.param.one)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one string
}

// one first  answer
type answer struct {
	one int
}
