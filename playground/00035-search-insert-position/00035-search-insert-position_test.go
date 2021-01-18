package _00035_search_insert_position

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInsert(t *testing.T) {
	// mock searchInsert
	qs := []question{
		{
			param{
				one: []int{1, 3, 5, 6},
				two: 5,
			},
			answer{2},
		},
		{
			param{
				one: []int{1, 3, 5, 6},
				two: 2,
			},
			answer{1},
		},
		{
			param{
				one: []int{1, 3, 5, 6},
				two: 7,
			},
			answer{4},
		},
		{
			param{
				one: []int{1, 3, 5, 6},
				two: 0,
			},
			answer{0},
		},
	}
	t.Logf("~> LeetCode searchInsert start")
	// do searchInsert
	for _, q := range qs {
		a, p := q.answer, q.param
		res := searchInsert(p.one, p.two)
		// verify searchInsert
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode searchInsert end")
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
