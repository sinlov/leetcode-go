package _00089_gray_code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock grayCode

var qs = []question{
	{
		param{2},
		answer{[]int{0, 1, 3, 2}},
	},
	{
		param{0},
		answer{[]int{0}},
	},
}

func Test_grayCode(t *testing.T) {
	t.Logf("~> LeetCode grayCode start")
	// do grayCode
	for _, q := range qs {
		a, p := q.answer, q.param
		res := grayCode(p.one)
		// verify grayCode
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode grayCode end")
}

func Benchmark_grayCode(b *testing.B) {
	// do grayCode
	for _, q := range qs {
		grayCode(q.param.one)
	}
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
	one []int
}
