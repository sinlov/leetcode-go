package _00050_powx_n

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mock myPow

var qs = []question{
	{
		param{
			one: 2.00000,
			two: 10,
		},
		answer{1024.00000},
	},
	{
		param{
			one: 2.10000,
			two: 3,
		},
		answer{9.26100},
	},
	{
		param{
			one: 2.00000,
			two: -2,
		},
		answer{0.25000},
	},
	{
		param{
			one: 2.00000,
			two: 1,
		},
		answer{2.00000},
	},
	{
		param{
			one: 2.00000,
			two: 0,
		},
		answer{1},
	},
}

func Test_myPow(t *testing.T) {
	t.Logf("~> LeetCode myPow start")
	// do myPow
	for _, q := range qs {
		a, p := q.answer, q.param
		tmp := myPow(p.one, p.two)
		res, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", tmp), 64)
		// verify myPow
		assert.Equal(t, a.one, res,
			"fail: in [ %v ], out [%v] , want [ %v ]", p.one, res, a.one)
	}
	t.Logf("~> LeetCode myPow end")
}

func Benchmark_myPow(b *testing.B) {
	// do myPow
	for _, q := range qs {
		myPow(q.param.one, q.param.two)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one float64
	two int
}

// one first  answer
type answer struct {
	one float64
}
