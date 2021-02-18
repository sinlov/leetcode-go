package _00063_unique_paths_ii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock uniquePathsWithObstacles

var qs = []question{
	{
		param{[][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}},
		answer{2},
	},
	{
		param{[][]int{
			{0, 0, 0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0},
		}},
		answer{13},
	},
	{
		param{[][]int{
			{0, 1},
			{0, 0},
		}},
		answer{1},
	},
	{
		param{[][]int{
			{1, 0},
			{0, 0},
		}},
		answer{0},
	},
	{
		param{[][]int{}},
		answer{0},
	},
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	t.Logf("~> LeetCode uniquePathsWithObstacles start")
	// do uniquePathsWithObstacles
	for _, q := range qs {
		a, p := q.answer, q.param
		res := uniquePathsWithObstacles(p.one)
		// verify uniquePathsWithObstacles
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode uniquePathsWithObstacles end")
}

func Benchmark_uniquePathsWithObstacles(b *testing.B) {
	// do uniquePathsWithObstacles
	for _, q := range qs {
		uniquePathsWithObstacles(q.param.one)
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
	one int
}
