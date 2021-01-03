package _00014_longest_common_prefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	// mock longestCommonPrefix
	qs := []question{
		{
			param{
				[]string{"", "b"},
			},
			answer{""},
		},
		{
			param{
				[]string{"ca", "a"},
			},
			answer{""},
		},
		{
			param{
				[]string{"vuhoexabc", "uthhxqzqwjfuvbabc", "utysicygqjzyjvyaigyabc"},
			},
			answer{""},
		},
		{
			param{
				[]string{"aflower", "bflow", "cfloight"},
			},
			answer{""},
		},
		{
			param{
				[]string{"ab", "a"},
			},
			answer{"a"},
		},
		{
			param{
				[]string{"a"},
			},
			answer{"a"},
		},
		{
			param{
				[]string{"flower", "flow", "flight"},
			},
			answer{"fl"},
		},
	}
	t.Logf("~> LeetCode longestCommonPrefix start")
	// do longestCommonPrefix
	for _, q := range qs {
		a, p := q.answer, q.param
		// verify longestCommonPrefix
		assert.Equal(t, a.one, longestCommonPrefix(p.one),
			"in [ %v ], out [%v]", p.one, longestCommonPrefix(p.one))
	}
	t.Logf("~> LeetCode longestCommonPrefix end")
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one []string
}

// one first  answer
type answer struct {
	one string
}
