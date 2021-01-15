package _00017_letter_combinations_of_a_phone_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinations(t *testing.T) {
	// mock letterCombinations
	qs := []question{
		{
			param{"635"},
			answer{[]string{
				"mdj", "mdk", "mdl", "mej", "mek", "mel", "mfj", "mfk", "mfl", "ndj", "ndk", "ndl", "nej", "nek", "nel", "nfj", "nfk", "nfl", "odj", "odk", "odl", "oej", "oek", "oel", "ofj", "ofk", "ofl",
			}},
		},
		{
			param{"23"},
			answer{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},
		{
			param{""},
			answer{[]string{}},
		},
	}
	t.Logf("~> LeetCode letterCombinations start")
	// do letterCombinations
	for _, q := range qs {
		a, p := q.answer, q.param
		res := letterCombinations(p.one)
		// verify letterCombinations
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode letterCombinations end")
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
	one []string
}
