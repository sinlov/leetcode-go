package _00019_remove_nth_node_from_end_of_list

import (
	"testing"

	"github.com/sinlov/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	// mock removeNthFromEnd
	qs := []question{
		{
			param{[]int{1, 2, 3, 4, 5}, 5},
			answer{[]int{2, 3, 4, 5}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}, 4},
			answer{[]int{1, 3, 4, 5}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}, 3},
			answer{[]int{1, 2, 4, 5}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}, 2},
			answer{[]int{1, 2, 3, 5}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}, 1},
			answer{[]int{1, 2, 3, 4}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}, 6},
			answer{[]int{1, 2, 3, 4, 5}},
		},
		{
			param{[]int{1, 2, 3, 4, 5}, -5},
			answer{[]int{1, 2, 3, 4, 5}},
		},
		{
			param{[]int{}, 5},
			answer{[]int{}},
		},
	}
	t.Logf("~> LeetCode removeNthFromEnd start")
	// do removeNthFromEnd
	for _, q := range qs {
		a, p := q.answer, q.param
		one := structures.IntList2ListNode(p.one)
		var res *ListNode
		if one == nil {
			res = removeNthFromEnd(nil, p.two)
		} else {
			res = removeNthFromEnd(one, p.two)
		}
		resList := structures.List2IntList(res)
		// verify removeNthFromEnd
		assert.Equal(t, a.one, resList,
			"fail: in [ %v , %v ], out [%v] , want [ %v ]", p.one, p.two, resList, a.one)
	}
	t.Logf("~> LeetCode removeNthFromEnd end")
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
