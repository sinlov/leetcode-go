package _00024_swap_nodes_in_pairs

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
func swapPairs(head *ListNode) *ListNode {
	if head == nil { // nil out
		return nil
	}
	if head.Next == nil { // recursion out
		return head
	}
	behind, previous := head, head.Next    // cache now
	behind.Next = swapPairs(previous.Next) // swap
	previous.Next = behind
	return previous
}
