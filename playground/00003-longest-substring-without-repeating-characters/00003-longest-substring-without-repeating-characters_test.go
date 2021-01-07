package _00003_longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	// mock lengthOfLongestSubstring
	qs := []question{
		{
			param{"abcabcbb"},
			answer{3},
		},
		{
			param{"pwwkew"},
			answer{3},
		},
		{
			param{"bbbbb"},
			answer{1},
		},
		{
			param{""},
			answer{0},
		},
	}
	t.Logf("~> LeetCode lengthOfLongestSubstring start")
	// do lengthOfLongestSubstring
	for _, q := range qs {
		a, p := q.answer, q.param
		res := lengthOfLongestSubstring(p.one)
		// verify lengthOfLongestSubstring
		assert.Equal(t, a.one, res,
			"in [ %v ], out [%v]", p.one, res)
	}
	t.Logf("~> LeetCode lengthOfLongestSubstring end")
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
