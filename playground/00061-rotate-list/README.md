### 解题思路

- 不能无脑循环，应该找出 O(n) 的复杂度的算法才行（禁止递归）
- 装载head 到 current 并计数 当前链表 length
  - k % length 为 0 则直接返回 head
  - 取模不为 0 则将 current 变成循环链表
    - 循环旋转，最终状态是确定的，利用链表的长度 取余 可以得到链表的最终旋转结果
    - 这个结果正好是 i 从 1 开始 i < (length - k%length) 旋转链表
  - 最终把循环链表指向输出，并将循环链表拆成普通链表输出

### 代码

```go
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
```
