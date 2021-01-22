package _00012_integer_to_roman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intToRoman(t *testing.T) {
	// mock intToRoman
	qs := []question{
		{
			param{1994},
			answer{"MCMXCIV"},
		},
		{
			param{511},
			answer{"DXI"},
		},
		{
			param{451},
			answer{"CDLI"},
		},
		{
			param{123},
			answer{"CXXIII"},
		},
		{
			param{58},
			answer{"LVIII"},
		},
		{
			param{45},
			answer{"XLV"},
		},
		{
			param{9},
			answer{"IX"},
		},
		{
			param{4},
			answer{"IV"},
		},
		{
			param{3},
			answer{"III"},
		},
	}
	t.Logf("~> LeetCode intToRoman start")
	// do intToRoman
	for _, q := range qs {
		a, p := q.answer, q.param
		res := intToRoman(p.one)
		// verify intToRoman
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode intToRoman end")
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
	one string
}
