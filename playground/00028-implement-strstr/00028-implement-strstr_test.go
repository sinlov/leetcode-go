package _00028_implement_strstr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_strStr(t *testing.T) {
	// mock strStr
	qs := []question{
		{
			param{"abbcbabc", "abc"},
			answer{5},
		},
		{
			param{"", "abc"},
			answer{-1},
		},
		{
			param{"abab", "ab"},
			answer{0},
		},
		{
			param{"aaaaa", "bba"},
			answer{-1},
		},
		{
			param{"hello", "ll"},
			answer{2},
		},
	}
	t.Logf("~> LeetCode strStr start")
	// do strStr
	for _, q := range qs {
		a, p := q.answer, q.param
		res := strStr(p.one, p.two)
		// verify strStr
		assert.Equal(t, a.one, res,
			"in [ %v , %v ], out [%v]", p.one, p.two, res)
	}
	t.Logf("~> LeetCode strStr end")
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
	one int
}
