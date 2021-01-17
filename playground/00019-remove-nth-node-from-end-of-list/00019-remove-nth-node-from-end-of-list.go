package _00019_remove_nth_node_from_end_of_list

import "github.com/sinlov/leetcode-go/structures"

// ListNode define
type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil { // invalid head
		return nil
	}
	if n <= 0 { // invalid n
		return head
	}

	closeP := head
	len := 0
	for closeP != nil { // count  head
		len++
		closeP = closeP.Next
	}

	if n > len { // grather so head retrun
		return head
	}

	if n == len { // equal is head.next
		closeP := head
		head = head.Next
		closeP.Next = nil
		return head
	}

	// less than
	closeP = head
	i := 0
	for closeP != nil {
		if i == len-n-1 { // move at some time between n
			deleteNode := closeP.Next
			closeP.Next = closeP.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		closeP = closeP.Next
	}
	return head
}
