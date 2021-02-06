package _00048_rotate_image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock rotate

var qs = []question{
	{
		param{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}},
		answer{[][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}},
	},
	{
		param{[][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}},
		answer{[][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		}},
	},
	{
		param{[][]int{
			{1, 2},
			{3, 4},
		}},
		answer{[][]int{
			{3, 1},
			{4, 2},
		}},
	},
	{
		param{[][]int{}},
		answer{[][]int{}},
	},
	{
		param{[][]int{
			{1},
		}},
		answer{[][]int{
			{1},
		}},
	},
}

func Test_rotate(t *testing.T) {
	t.Logf("~> LeetCode rotate start")
	// do rotate
	for _, q := range qs {
		a, p := q.answer, q.param
		rotate(p.one)
		// verify rotate
		assert.Equal(t, a.one, p.one,
			"fail: out [%v] , want [ %v ]", p.one, a.one)
	}
	t.Logf("~> LeetCode rotate end")
}

func Benchmark_rotate(b *testing.B) {
	// do rotate
	for _, q := range qs {
		rotate(q.param.one)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one [][]int
}

// one first  answer
type answer struct {
	one [][]int
}
