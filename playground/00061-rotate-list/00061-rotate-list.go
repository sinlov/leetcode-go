package _00061_rotate_list

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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 { // special case
		return head
	}

	current := head
	length := 1
	for current.Next != nil { // load head add head
		length++
		current = current.Next
	}
	if (k % length) == 0 { // turn circle so return head
		return head
	}

	current.Next = head // let current to circular linked list
	current = head

	for i := 1; i < (length - k%length); i++ { // let link to final
		current = current.Next
	}
	res := current.Next
	current.Next = nil // cut circular linked list
	return res
}
