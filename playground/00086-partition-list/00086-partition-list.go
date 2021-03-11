package _00086_partition_list

import "github.com/sinlov/leetcode-go/structures"

// ListNode define
type ListNode = structures.ListNode

func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{Val: 0, Next: nil} // before
	before := beforeHead
	afterHead := &ListNode{Val: 0, Next: nil} // after
	after := afterHead

	for head != nil {
		if head.Val < x { // less than change before
			before.Next = head
			before = before.Next
		} else { // other change after
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}
