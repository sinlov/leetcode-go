### 解题思路

- 构造一个新的头结点指向当前的头
- 找到第一个需要反转的结点的前一个结点 p
  - 循环次数用 n-m 来控制
    - 从这个结点开始，依次把后面的结点用“头插”法，插入到 p 结点的后面

### 代码

```go
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left >= right { // speace case head nil or left ge right
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	prev := newHead
	for count := 0; prev.Next != nil && count < left-1; count++ {
		prev = prev.Next
	}
	current := prev.Next
	for i := 0; i < right-left; i++ {
		tmp := prev.Next
		prev.Next = current.Next
		current.Next = current.Next.Next
		prev.Next.Next = tmp
	}
	return newHead.Next
}
```
