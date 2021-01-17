### 解题思路

- 排除不用删除的情况
- 2 个指针，一个指针距离前一个指针 n 个距离
- 同时移动 2 个指针，2 个指针都移动相同的距离
- 当一个指针移动到了终点，那么前一个指针就是倒数第 n 个节点

### 代码

```go
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
```
