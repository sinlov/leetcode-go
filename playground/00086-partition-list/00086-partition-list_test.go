package _00086_partition_list

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock partition

var qs = []question{
	{
		param{
			one: []int{1, 4, 3, 2, 5, 2},
			two: 3,
		},
		answer{[]int{1, 2, 2, 4, 3, 5}},
	},
	{
		param{
			one: []int{2, 1},
			two: 2,
		},
		answer{[]int{1, 2}},
	},
}

func Test_partition(t *testing.T) {
	t.Logf("~> LeetCode partition start")
	// do partition
	for _, q := range qs {
		a, p := q.answer, q.param
		midLN := structures.IntList2ListNode(p.one)
		resLN := partition(midLN, p.two)
		res := structures.List2IntList(resLN)
		// verify partition
		assert.Equal(t, a.one, res,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, res, a.one)
	}
	t.Logf("~> LeetCode partition end")
}

func Benchmark_partition(b *testing.B) {
	// do partition
	for _, q := range qs {
		partition(structures.IntList2ListNode(q.param.one), q.param.two)
	}
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one []int
	two int
}

// one first  answer
type answer struct {
	one []int
}
