package _00020_valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	// mock isValid
	qs := []question{
		{
			param{"(())]]"},
			answer{false},
		},
		{
			param{"(){[({[]})]}"},
			answer{true},
		},
		{
			param{"(}"},
			answer{false},
		},
		{
			param{")"},
			answer{false},
		},
		{
			param{"["},
			answer{false},
		},
		{
			param{"()[]{}"},
			answer{true},
		},
		{
			param{"()"},
			answer{true},
		},
		{
			param{""},
			answer{true},
		},
	}
	t.Logf("~> LeetCode isValid start")
	// do isValid
	for _, q := range qs {
		a, p := q.answer, q.param
		res := isValid(p.one)
		// verify isValid
		assert.Equal(t, a.one, res,
			"in [ %v ], out [%v]", p.one, res)
	}
	t.Logf("~> LeetCode isValid end")
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
	one bool
}
