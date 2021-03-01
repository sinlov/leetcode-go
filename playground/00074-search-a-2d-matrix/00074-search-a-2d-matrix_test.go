package _00074_search_a_2d_matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock searchMatrix

var qs = []question{
	{
		param{one: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
			two: 34,
		},
		answer{true},
	},
	{
		param{one: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
			two: 3,
		},
		answer{true},
	},
	{
		param{one: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 60},
		},
			two: 12,
		},
		answer{false},
	},
	{
		param{one: [][]int{},
			two: 3,
		},
		answer{false},
	},
}

func Test_searchMatrix(t *testing.T) {
	t.Logf("~> LeetCode searchMatrix start")
	// do searchMatrix
	for _, q := range qs {
		a, p := q.answer, q.param
		res := searchMatrix(p.one, p.two)
		// verify searchMatrix
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode searchMatrix end")
}

func Benchmark_searchMatrix(b *testing.B) {
	// do searchMatrix
	for _, q := range qs {
		searchMatrix(q.param.one, q.param.two)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one [][]int
	two int
}

// one first  answer
type answer struct {
	one bool
}
