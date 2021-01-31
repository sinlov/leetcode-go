package _00038_count_and_say

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock countAndSay

var qs = []question{
	{
		param{8},
		answer{"1113213211"},
	},
	{
		param{7},
		answer{"13112221"},
	},
	{
		param{6},
		answer{"312211"},
	},
	{
		param{5},
		answer{"111221"},
	},
	{
		param{4},
		answer{"1211"},
	},
	{
		param{1},
		answer{"1"},
	},
}

func Test_countAndSay(t *testing.T) {
	t.Logf("~> LeetCode countAndSay start")
	// do countAndSay
	for _, q := range qs {
		a, p := q.answer, q.param
		res := countAndSay(p.one)
		// verify countAndSay
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode countAndSay end")
}

func Benchmark_countAndSay(b *testing.B) {
	// do countAndSay
	for _, q := range qs {
		countAndSay(q.param.one)
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
	one string
}
