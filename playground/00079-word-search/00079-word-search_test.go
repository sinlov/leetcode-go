package _00079_word_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock exist

var qs = []question{
	{
		param{
			one: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			two: "ABCCED",
		},
		answer{true},
	},
	{
		param{
			one: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			two: "SEE",
		},
		answer{true},
	},
	{
		param{
			one: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			two: "ABCB",
		},
		answer{false},
	},
}

func Test_exist(t *testing.T) {
	t.Logf("~> LeetCode exist start")
	// do exist
	for _, q := range qs {
		a, p := q.answer, q.param
		res := exist(p.one, p.two)
		// verify exist
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode exist end")
}

func Benchmark_exist(b *testing.B) {
	// do exist
	for _, q := range qs {
		exist(q.param.one, q.param.two)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one [][]byte
	two string
}

// one first  answer
type answer struct {
	one bool
}
