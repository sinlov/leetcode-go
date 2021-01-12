package _00011_container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	// mock maxArea
	qs := []question{
		{
			param{[]int{1, 2, 1}},
			answer{2},
		},
		{
			param{[]int{4, 3, 2, 1, 4}},
			answer{16},
		},
		{
			param{[]int{1, 1}},
			answer{1},
		},
		{
			param{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			answer{49},
		},
	}
	t.Logf("~> LeetCode maxArea start")
	// do maxArea
	for _, q := range qs {
		a, p := q.answer, q.param
		res := maxArea(p.one)
		// verify maxArea
		assert.Equal(t, a.one, res,
			"in [ %v ], out [%v]", p.one, res)
	}
	t.Logf("~> LeetCode maxArea end")
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
	one int
}
