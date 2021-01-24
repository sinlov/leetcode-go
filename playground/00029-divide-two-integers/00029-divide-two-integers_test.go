package _00029_divide_two_integers

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divide(t *testing.T) {
	// mock divide
	qs := []question{
		{
			param{math.MinInt32, -1},
			answer{math.MaxInt32},
		},
		{
			param{7, -3},
			answer{-2},
		},
		{
			param{10, 3},
			answer{3},
		},
	}
	t.Logf("~> LeetCode divide start")
	// do divide
	for _, q := range qs {
		a, p := q.answer, q.param
		res := divide(p.one, p.two)
		// verify divide
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode divide end")
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
	one int
}
