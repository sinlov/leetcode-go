package _00066_plus_one

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock plusOne

var qs = []question{
	{
		param{[]int{1, 9, 9}},
		answer{[]int{2, 0, 0}},
	},
	{
		param{[]int{9, 9}},
		answer{[]int{1, 0, 0}},
	},
	{
		param{[]int{1, 2, 3}},
		answer{[]int{1, 2, 4}},
	},
	{
		param{[]int{0}},
		answer{[]int{1}},
	},
}

func Test_plusOne(t *testing.T) {
	t.Logf("~> LeetCode plusOne start")
	// do plusOne
	for _, q := range qs {
		a, p := q.answer, q.param
		res := plusOne(p.one)
		// verify plusOne
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode plusOne end")
}

func Benchmark_plusOne(b *testing.B) {
	// do plusOne
	for _, q := range qs {
		plusOne(q.param.one)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one []int
}

// one first  answer
type answer struct {
	one []int
}
