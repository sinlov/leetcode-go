### 解题思路

- 单循环，双指针+删除标志位
- 每次只删除前面的一个重复的元素，留一个用于下次遍历判重
- pre, back 指针的更新位置和值比较重要和巧妙

### 代码

```go
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
```
