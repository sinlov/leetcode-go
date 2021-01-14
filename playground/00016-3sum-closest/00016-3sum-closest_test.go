package _00016_3sum_closest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_threeSumClosest(t *testing.T) {
	// mock threeSumClosest
	qs := []question{
		{
			param{[]int{-1, 0, 2, 1, 45}, 3},
			answer{3},
		},
		{
			param{[]int{-1, 0, 1, 1, 45}, 3},
			answer{2},
		},
		{
			param{[]int{0, 0, 0}, 1},
			answer{0},
		},
		{
			param{[]int{1, 1, -1}, 0},
			answer{1},
		},
		{
			param{[]int{-1, 2, 1, -4}, 1},
			answer{2},
		},
		{
			param{[]int{-1}, 1},
			answer{0},
		},
	}
	t.Logf("~> LeetCode threeSumClosest start")
	// do threeSumClosest
	for _, q := range qs {
		a, p := q.answer, q.param
		res := threeSumClosest(p.one, p.two)
		// verify threeSumClosest
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode threeSumClosest end")
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one []int
	two int
}

// one first  answer
type answer struct {
	one int
}
