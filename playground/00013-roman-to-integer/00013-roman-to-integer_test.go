package _0013_roman_to_integer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	// mock romanToInt
	qs := []question{
		{
			param{"III"},
			answer{3},
		},
		{
			param{"IV"},
			answer{4},
		},
		{
			param{"IX"},
			answer{9},
		},
		{
			param{"LVIII"},
			answer{58},
		},
		{
			param{"MCMXCIV"},
			answer{1994},
		},
		{
			param{"MCMXICIVI"},
			answer{2014},
		},
	}
	t.Logf("~> LeetCode romanToInt start")
	// do romanToInt
	for _, q := range qs {
		a, p := q.answer, q.param
		//t.Logf("in [ %v ], out [%v]", p.one, romanToInt(p.one))
		// verify romanToInt
		assert.Equal(t, a.one, romanToInt(p.one),
			"in [ %v ], out [%v]", p.one, romanToInt(p.one))
	}
	t.Logf("~> LeetCode romanToInt end")
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
