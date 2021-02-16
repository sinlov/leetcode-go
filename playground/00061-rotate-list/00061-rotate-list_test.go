package _00061_rotate_list

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

// mock rotateRight

var qs = []question{
	{
		param{
			one: []int{1, 2, 3, 4, 5},
			two: 2,
		},
		answer{[]int{4, 5, 1, 2, 3}},
	},
	{
		param{
			one: []int{1, 2, 3, 4, 5},
			two: 3,
		},
		answer{[]int{3, 4, 5, 1, 2}},
	},
	{
		param{
			one: []int{1, 2, 3, 4, 5},
			two: 5,
		},
		answer{[]int{1, 2, 3, 4, 5}},
	},
	{
		param{
			one: []int{0, 1, 2},
			two: 4,
		},
		answer{[]int{2, 0, 1}},
	},
	{
		param{
			one: []int{},
			two: 2,
		},
		answer{[]int{}},
	},
}

func Test_rotateRight(t *testing.T) {
	t.Logf("~> LeetCode rotateRight start")
	// do rotateRight
	for _, q := range qs {
		a, p := q.answer, q.param
		head := structures.IntList2ListNode(p.one)
		res := rotateRight(head, p.two)
		resInt := structures.List2IntList(res)
		// verify rotateRight
		assert.Equal(t, a.one, resInt,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, resInt, a.one)
	}
	t.Logf("~> LeetCode rotateRight end")
}

func Benchmark_rotateRight(b *testing.B) {
	// do rotateRight
	for _, q := range qs {
		head := structures.IntList2ListNode(q.param.one)
		rotateRight(head, q.param.two)
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
