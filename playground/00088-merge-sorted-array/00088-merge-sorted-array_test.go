package _00088_merge_sorted_array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock merge

var qs = []question{
	{
		param{
			one:   []int{10, 13, 15, 27, 30, 52},
			two:   4,
			three: []int{1, 3},
			four:  2,
		},
		answer{[]int{1, 3, 10, 13, 15, 27}},
	},
	{
		param{
			one:   []int{1, 2, 3, 0, 0, 0},
			two:   3,
			three: []int{2, 5, 6},
			four:  3,
		},
		answer{[]int{1, 2, 2, 3, 5, 6}},
	},
	{
		param{
			one:   []int{1},
			two:   1,
			three: []int{},
			four:  0,
		},
		answer{[]int{1}},
	},
	{
		param{
			one:   []int{1},
			two:   0,
			three: []int{},
			four:  0,
		},
		answer{[]int{1}},
	},
}

func Test_merge(t *testing.T) {
	t.Logf("~> LeetCode merge start")
	// do merge
	for _, q := range qs {
		a, p := q.answer, q.param
		merge(p.one, p.two, p.three, p.four)
		// verify merge
		assert.Equal(t, a.one, p.one,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, p.one, a.one)
	}
	t.Logf("~> LeetCode merge end")
}

func Benchmark_merge(b *testing.B) {
	// do merge
	for _, q := range qs {
		merge(q.param.one, q.param.two, q.param.three, q.param.four)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one   []int
	two   int
	three []int
	four  int
}

// one first  answer
type answer struct {
	one []int
}
