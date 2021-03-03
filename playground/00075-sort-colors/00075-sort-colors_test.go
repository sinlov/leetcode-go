package _00075_sort_colors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock sortColors

var qs = []question{
	{
		param{[]int{2, 0, 2, 1, 1, 0}},
		answer{[]int{0, 0, 1, 1, 2, 2}},
	},
	{
		param{[]int{2, 0, 1}},
		answer{[]int{0, 1, 2}},
	},
	{
		param{[]int{1}},
		answer{[]int{1}},
	},
	{
		param{[]int{0}},
		answer{[]int{0}},
	},
	{
		param{[]int{}},
		answer{[]int{}},
	},
}

func Test_sortColors(t *testing.T) {
	t.Logf("~> LeetCode sortColors start")
	// do sortColors
	for _, q := range qs {
		a, p := q.answer, q.param
		sortColors(p.one)
		// verify sortColors
		assert.Equal(t, a.one, p.one,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, p.one, a.one)
	}
	t.Logf("~> LeetCode sortColors end")
}

func Benchmark_sortColors(b *testing.B) {
	// do sortColors
	for _, q := range qs {
		sortColors(q.param.one)
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
	one []int
}
