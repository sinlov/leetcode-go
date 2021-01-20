package _00008_string_to_integer_atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi(t *testing.T) {
	// mock myAtoi
	qs := []question{
		{
			param{"4193 with words"},
			answer{4193},
		},
		{
			param{"+91283472332"},
			answer{2147483647},
		},
		{
			param{"-91283472332"},
			answer{-2147483648},
		},
		{
			param{"words and 987"},
			answer{0},
		},
		{
			param{"   -42"},
			answer{-42},
		},
		{
			param{"42"},
			answer{42},
		},
		{
			param{""},
			answer{0},
		},
	}
	t.Logf("~> LeetCode myAtoi start")
	// do myAtoi
	for _, q := range qs {
		a, p := q.answer, q.param
		res := myAtoi(p.one)
		// verify myAtoi
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode myAtoi end")
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
