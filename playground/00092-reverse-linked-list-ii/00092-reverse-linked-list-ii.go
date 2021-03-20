package _00092_reverse_linked_list_ii

import "github.com/sinlov/leetcode-go/structures"

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left >= right { // speace case head nil or left ge right
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	prev := newHead
	for count := 0; prev.Next != nil && count < left-1; count++ {
		prev = prev.Next
	}
	// if prev.Next == nil {
	// 	return head
	// }
	current := prev.Next
	for i := 0; i < right-left; i++ {
		tmp := prev.Next
		prev.Next = current.Next
		current.Next = current.Next.Next
		prev.Next.Next = tmp
	}
	return newHead.Next
}
