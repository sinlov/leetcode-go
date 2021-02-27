package _00023_merge_k_sorted_lists

import "github.com/sinlov/leetcode-go/structures"

// ListNode define
type ListNode = structures.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 { // special case
		return nil
	}
	if length == 1 { // special case
		return lists[0]
	}
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	return mergeAscendingTwoListsToOne(left, right)
}

func mergeAscendingTwoListsToOne(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeAscendingTwoListsToOne(l1.Next, l2)
		return l1
	}
	l2.Next = mergeAscendingTwoListsToOne(l1, l2.Next)
	return l2
}
