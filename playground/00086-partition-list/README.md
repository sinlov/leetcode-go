### 解题思路

用额外的 before 和 head 来找对应的位置，做到正确的交换即可

### 代码

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
```
