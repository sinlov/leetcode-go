### 解题思路

- 排除特殊情况
- 设置递归出口
- 递归交换即可

### 代码

```go
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
```
