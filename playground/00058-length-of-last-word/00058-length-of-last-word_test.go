package _00058_length_of_last_word

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock lengthOfLastWord

var qs = []question{
	{
		param{"a"},
		answer{1},
	},
	{
		param{"a "},
		answer{1},
	},
	{
		param{"Hello World"},
		answer{5},
	},
	{
		param{""},
		answer{0},
	},
}

func Test_lengthOfLastWord(t *testing.T) {
	t.Logf("~> LeetCode lengthOfLastWord start")
	// do lengthOfLastWord
	for _, q := range qs {
		a, p := q.answer, q.param
		res := lengthOfLastWord(p.one)
		// verify lengthOfLastWord
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode lengthOfLastWord end")
}

func Benchmark_lengthOfLastWord(b *testing.B) {
	// do lengthOfLastWord
	for _, q := range qs {
		lengthOfLastWord(q.param.one)
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
