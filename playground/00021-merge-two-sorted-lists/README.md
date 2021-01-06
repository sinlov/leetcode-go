### 解题思路

- 最简单情况，l1 空，则合并结果就是 l2, 反之为 l1
- 比较时，如果 l1 当前值 小于 l2 当前值， 那么插入链表 l2 左侧, 返回 l1 
- 如果 l1 大于等于 l2 的值，那么插入 l2 右侧, 返回 l2

### 代码

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// simple merge
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val { // merge left
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	// merge right
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2 // return l2
}
```
