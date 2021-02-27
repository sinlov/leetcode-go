### 解题思路

- k 个相邻的元素，移动一次链表
- 翻转链表
- 递归即可

### 代码

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverseTwoListNode(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverseTwoListNode(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}
```
