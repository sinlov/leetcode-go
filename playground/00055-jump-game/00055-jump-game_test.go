package _00055_jump_game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock canJump

var qs = []question{
	{
		param{[]int{2, 3, 1, 1, 4}},
		answer{true},
	},
	{
		param{[]int{3, 2, 1, 0, 4}},
		answer{false},
	},
	{
		param{[]int{1}},
		answer{true},
	},
	{
		param{[]int{}},
		answer{false},
	},
}

func Test_canJump(t *testing.T) {
	t.Logf("~> LeetCode canJump start")
	// do canJump
	for _, q := range qs {
		a, p := q.answer, q.param
		res := canJump(p.one)
		// verify canJump
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode canJump end")
}

func Benchmark_canJump(b *testing.B) {
	// do canJump
	for _, q := range qs {
		canJump(q.param.one)
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
	one bool
}
