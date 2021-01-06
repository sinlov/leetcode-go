package _00021_merge_two_sorted_lists

import (
	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	// mock mergeTwoLists
	qs := []question{
		{
			param{
				one: []int{1},
				two: []int{9, 9, 9, 9},
			},
			answer{
				one: []int{1, 9, 9, 9, 9},
			},
		},
		{
			param{
				one: []int{1, 2, 3, 4},
				two: []int{5, 6, 7, 8},
			},
			answer{
				one: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
		},
		{
			param{
				one: []int{1, 2, 3, 4},
				two: []int{1, 2, 3, 4},
			},
			answer{
				one: []int{1, 1, 2, 2, 3, 3, 4, 4},
			},
		},
		{
			param{
				one: []int{2},
				two: []int{1},
			},
			answer{
				one: []int{1, 2},
			},
		},
		{
			param{
				one: []int{1},
				two: []int{1},
			},
			answer{
				one: []int{1, 1},
			},
		},
		{
			param{
				one: []int{2},
				two: []int{},
			},
			answer{
				one: []int{2},
			},
		},
		{
			param{
				one: []int{},
				two: []int{1},
			},
			answer{
				one: []int{1},
			},
		},
		{
			param{
				one: []int{},
				two: []int{},
			},
			answer{
				one: []int{},
			},
		},
	}
	t.Logf("~> LeetCode mergeTwoLists start")
	// do mergeTwoLists
	for _, q := range qs {
		a, p := q.answer, q.param

		res := mergeTwoLists(structures.IntList2ListNode(p.one), structures.IntList2ListNode(p.two))
		chk := structures.List2IntList(res)
		// verify mergeTwoLists
		assert.Equal(t, a.one, chk,
			"in [ %v, %v ], out [%v]", p.one, p.two, chk)
	}
	t.Logf("~> LeetCode mergeTwoLists end")
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
