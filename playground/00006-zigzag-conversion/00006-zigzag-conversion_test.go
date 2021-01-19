package _00006_zigzag_conversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_convert(t *testing.T) {
	// mock convert
	qs := []question{
		{
			param{
				one: "A",
				two: 1,
			},
			answer{"A"},
		},
		{
			param{
				one: "PAYPALISHIRING",
				two: 4,
			},
			answer{"PINALSIGYAHRPI"},
		},
		{
			param{
				one: "PAYPALISHIRING",
				two: 3,
			},
			answer{"PAHNAPLSIIGYIR"},
		},
	}
	t.Logf("~> LeetCode convert start")
	// do convert
	for _, q := range qs {
		a, p := q.answer, q.param
		res := convert(p.one, p.two)
		// verify convert
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode convert end")
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one string
	two int
}

// one first  answer
type answer struct {
	one string
}
