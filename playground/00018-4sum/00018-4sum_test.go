package _00018_4sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fourSum(t *testing.T) {
	// mock fourSum
	qs := []question{
		{
			param{[]int{1, 0, -1, 0, -2, 2, 0, 0, 0, 0}, 0},
			answer{[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}, {0, 0, 0, 0}}},
		},
		{
			param{[]int{0, 1, 5, 0, 1, 5, 5, -4}, 11},
			answer{[][]int{{-4, 5, 5, 5}, {0, 1, 5, 5}}},
		},
		{
			param{[]int{1, 1, 1, 1}, 4},
			answer{[][]int{{1, 1, 1, 1}}},
		},
	}
	t.Logf("~> LeetCode fourSum start")
	// do fourSum
	for _, q := range qs {
		a, p := q.answer, q.param
		res := fourSum(p.one, p.two)
		// verify fourSum
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode fourSum end")
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
	one [][]int
}
