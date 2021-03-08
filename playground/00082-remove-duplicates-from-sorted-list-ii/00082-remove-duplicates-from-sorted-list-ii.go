package _00082_remove_duplicates_from_sorted_list_ii

import "github.com/sinlov/leetcode-go/structures"

// ListNode define
type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nilNode := &ListNode{Val: 0, Next: head}
	lastIsDel := false // flag of last delete

	head = nilNode                         // nil node
	pre, back := head.Next, head.Next.Next // double point of pre and back

	for head.Next != nil && head.Next.Next != nil { // delete only one element at one cycle, leaving one for the next iteration
		if pre.Val != back.Val && lastIsDel {
			head.Next = head.Next.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = false
			continue
		}

		if pre.Val == back.Val {
			head.Next = head.Next.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = true
		} else {
			head = head.Next
			pre, back = head.Next, head.Next.Next
			lastIsDel = false
		}
	}

	if lastIsDel && head.Next != nil { //special case as [1,1]
		head.Next = nil
	}
	return nilNode.Next
}
