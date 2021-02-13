package _00057_insert_interval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock insert

var qs = []question{
	{
		param{
			one: [][]int{
				{1, 3},
				{6, 9},
			},
			two: []int{2, 5},
		},
		answer{[][]int{
			{1, 5},
			{6, 9},
		}},
	},
	{
		param{
			one: [][]int{
				{1, 2},
				{3, 5},
				{6, 7},
				{8, 10},
				{12, 16},
			},
			two: []int{4, 8},
		},
		answer{[][]int{
			{1, 2},
			{3, 10},
			{12, 16},
		}},
	},
	{
		param{
			one: [][]int{},
			two: []int{5, 7},
		},
		answer{[][]int{
			{5, 7},
		}},
	},
	{
		param{
			one: [][]int{
				{1, 3},
				{6, 9},
			},
			two: []int{},
		},
		answer{[][]int{
			{1, 3},
			{6, 9},
		}},
	},
	{
		param{
			one: [][]int{
				{1, 5},
			},
			two: []int{2, 3},
		},
		answer{[][]int{
			{1, 5},
		}},
	},
	{
		param{
			one: [][]int{
				{1, 5},
			},
			two: []int{2, 7},
		},
		answer{[][]int{
			{1, 7},
		}},
	},
}

func Test_insert(t *testing.T) {
	t.Logf("~> LeetCode insert start")
	// do insert
	for _, q := range qs {
		a, p := q.answer, q.param
		res := insert(p.one, p.two)
		// verify insert
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode insert end")
}

func Benchmark_insert(b *testing.B) {
	// do insert
	for _, q := range qs {
		insert(q.param.one, q.param.two)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one [][]int
	two []int
}

// one first  answer
type answer struct {
	one [][]int
}
