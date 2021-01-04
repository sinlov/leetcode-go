package _00002_add_two_numbers

import (
	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	// mock addTwoNumbers
	qs := []question{
		{
			param{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
			answer{[]int{8, 9, 9, 9, 0, 0, 0, 1}},
		},
		{
			param{[]int{9, 9, 9, 9, 9}, []int{1}},
			answer{[]int{0, 0, 0, 0, 0, 1}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			answer{[]int{2, 4, 6, 8, 0, 1}},
		},
		{
			param{[]int{9}, []int{1}},
			answer{[]int{0, 1}},
		},
		{
			param{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			answer{[]int{2, 4, 6, 8}},
		},
		{
			param{[]int{1, 8, 2}, []int{7, 1}},
			answer{[]int{8, 9, 2}},
		},
		{
			param{[]int{2, 4, 3}, []int{5, 6, 4}},
			answer{[]int{7, 0, 8}},
		},
		{
			param{[]int{1}, []int{1}},
			answer{[]int{2}},
		},
		{
			param{[]int{}, []int{}},
			answer{[]int{}},
		},
	}
	t.Logf("~> LeetCode addTwoNumbers start")
	// do addTwoNumbers
	for _, q := range qs {
		a, p := q.answer, q.param
		// verify addTwoNumbers
		outListNode := addTwoNumbers(structures.IntList2ListNode(p.one), structures.IntList2ListNode(p.two))
		actionList := structures.List2IntList(outListNode)
		assert.Equal(t, a.one, actionList,
			"in [ %v ], [ %v ], out [%v]", p.one, p.two, actionList)
	}
	t.Logf("~> LeetCode addTwoNumbers end")
}

type question struct {
	param
	answer
}

// one first param
type param struct {
	one []int
	two []int
}

// one first  answer
type answer struct {
	one []int
}
