package _00080_remove_duplicates_from_sorted_array_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock removeDuplicates

var qs = []question{
	{
		param{[]int{1, 1, 1, 2, 2, 3}},
		answer{5},
	},
	{
		param{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
		answer{7},
	},
	{
		param{[]int{}},
		answer{0},
	},
	{
		param{[]int{1}},
		answer{1},
	},
}

func Test_removeDuplicates(t *testing.T) {
	t.Logf("~> LeetCode removeDuplicates start")
	// do removeDuplicates
	for _, q := range qs {
		a, p := q.answer, q.param
		res := removeDuplicates(p.one)
		// verify removeDuplicates
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode removeDuplicates end")
}

func Benchmark_removeDuplicates(b *testing.B) {
	// do removeDuplicates
	for _, q := range qs {
		removeDuplicates(q.param.one)
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
	one int
}
