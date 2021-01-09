package _00027_remove_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeElement(t *testing.T) {
	// mock removeElement
	qs := []question{
		{
			param{[]int{0, 1, 0, 3, 0, 0, 0, 0, 1, 11}, 0},
			answer{4},
		},
		{
			param{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			answer{5},
		},
		{
			param{[]int{3, 2, 2, 3}, 3},
			answer{2},
		},
		{
			param{[]int{1}, 1},
			answer{0},
		},
		{
			param{[]int{}, 0},
			answer{0},
		},
	}
	t.Logf("~> LeetCode removeElement start")
	// do removeElement
	for _, q := range qs {
		a, p := q.answer, q.param
		res := removeElement(p.one, p.two)
		// verify removeElement
		assert.Equal(t, a.one, res,
			"in [ %v , %v ], out [%v]", p.one, p.two, res)
	}
	t.Logf("~> LeetCode removeElement end")
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
