package structures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList2IntList(t *testing.T) {
	ast := assert.New(t)
	ast.Equal([]int{}, List2IntList(nil), "input nil，not return []int{}")

	one2three := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	ast.Equal([]int{1, 2, 3}, List2IntList(one2three), "not be []int")

	limit := 100
	overLimitList := IntList2ListNode(make([]int, limit+1))
	ast.Panics(func() {
		List2IntList(overLimitList)
	},
		"convert over %d, but not panic", limit,
	)
}

func TestIntList2ListNode(t *testing.T) {
	ast := assert.New(t)
	ast.Nil(IntList2ListNode([]int{}), "input []int{}，not return nil")

	listNode := IntList2ListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	i := 1
	for listNode != nil {
		ast.Equal(i, listNode.Val, "listNode.Val not equal")
		listNode = listNode.Next
		i++
	}
}

func TestListNode_GetNodeWith(t *testing.T) {
	ast := assert.New(t)
	listNode := IntList2ListNode([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	val := 10
	node := &ListNode{
		Val: val,
	}
	tail := listNode
	for tail.Next != nil {
		tail = tail.Next
	}

	tail.Next = node
	expected := node
	actual := listNode.GetNodeWith(val)
	ast.Equal(expected, actual, "GetNodeWith val: %d, error", val)
}

func TestIntList2ListWithCycle(t *testing.T) {
	ast := assert.New(t)
	intList := []int{1, 2, 3}
	listWithCycle := IntList2ListWithCycle(intList, -1)
	ast.Equal(intList, List2IntList(listWithCycle), "not some as IntList2ListWithCycle")

	errPos := 1
	withCycleOne := IntList2ListWithCycle(intList, errPos)
	ast.Panics(func() {
		List2IntList(withCycleOne)
	},
		"not panic with error pos %d when IntList2ListWithCycle", errPos,
	)
}
