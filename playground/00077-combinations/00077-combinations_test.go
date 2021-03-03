package _00077_combinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock combine

var qs = []question{
	{
		param{one: 4, two: 2},
		answer{[][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		}},
	},
	{
		param{one: 2, two: 4},
		answer{[][]int{}},
	},
	{
		param{one: -4, two: 2},
		answer{[][]int{}},
	},
}

func Test_combine(t *testing.T) {
	t.Logf("~> LeetCode combine start")
	// do combine
	for _, q := range qs {
		a, p := q.answer, q.param
		res := combine(p.one, p.two)
		// verify combine
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode combine end")
}

func Benchmark_combine(b *testing.B) {
	// do combine
	for _, q := range qs {
		combine(q.param.one, q.param.two)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one int
	two int
}

// one first  answer
type answer struct {
	one [][]int
}
