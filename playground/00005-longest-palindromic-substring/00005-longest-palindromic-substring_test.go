package _00005_longest_palindromic_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	// mock longestPalindrome
	qs := []question{
		{
			param{"abctenet"},
			answer{"tenet"},
		},
		{
			param{"tenet"},
			answer{"tenet"},
		},
		{
			param{"ac"},
			answer{"a"},
		},
		{
			param{"a"},
			answer{"a"},
		},
		{
			param{"cbbd"},
			answer{"bb"},
		},
		{
			param{"babad"},
			answer{"bab"},
		},
		{
			param{""},
			answer{""},
		},
	}
	t.Logf("~> LeetCode longestPalindrome start")
	// do longestPalindrome
	for _, q := range qs {
		a, p := q.answer, q.param
		res := longestPalindrome(p.one)
		// verify longestPalindrome
		assert.Equal(t, a.one, res,
			"in [ %v ], out [%v]", p.one, res)
	}
	t.Logf("~> LeetCode longestPalindrome end")
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
	one string
}
