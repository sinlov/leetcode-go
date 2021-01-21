package _00022_generate_parentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateParenthesis(t *testing.T) {
	// mock generateParenthesis
	qs := []question{
		{
			param{0},
			answer{[]string{""}},
		},
		{
			param{1},
			answer{[]string{"()"}},
		},
		{
			param{2},
			answer{[]string{"(())", "()()"}},
		},
		{
			param{3},
			answer{[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()"},
			},
		},
	}
	t.Logf("~> LeetCode generateParenthesis start")
	// do generateParenthesis
	for _, q := range qs {
		a, p := q.answer, q.param
		res := generateParenthesis(p.one)
		// verify generateParenthesis
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode generateParenthesis end")
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
	one []string
}
