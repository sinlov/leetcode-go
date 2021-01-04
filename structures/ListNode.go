package structures

import "fmt"

// ListNode is link node
type ListNode struct {
	Val  int
	Next *ListNode
}

// will panic when limit depth beyond so limit 100
func List2IntList(head *ListNode) []int {
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("list beyond %d, will be cycle. please check or change limit at function of structures.List2IntList.", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// int list convert to ListNode
func IntList2ListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

// GetNodeWith returns the first node with val as ListNode
func (l *ListNode) GetNodeWith(val int) *ListNode {
	res := l
	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}
	return res
}

// IntList2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos = -1, no cycle
func IntList2ListWithCycle(nums []int, pos int) *ListNode {
	head := IntList2ListNode(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
