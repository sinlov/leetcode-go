## 解题思路

逆序链表，就是数据结构中的 ListNode

- 要求`从低位开始相加`，得出结果也`逆序输出`，返回值是逆序结果链表的头结点
- 一定注意的是`进位问题`

甚至这种情况

```
In: (9 -> 9 -> 9 -> 9 -> 9) + (1 -> )
Out: (0 -> 0 -> 0 -> 0 -> 0 -> 1)
```

处理时

- 先建立一个虚拟头结点
- 这个虚拟头结点的 `Next current`, `current` 指向真正的 head 这样 `最终 head 不需要单独处理`
- 且最终获取的是 head.next, carry 是前一位的（逆序就是进位的前一位），一旦进位，不为0
- 循环终止条件应该是 l1 != nil l2 != nil 并且 carry 不为 0

## 代码

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	n1, n2, carry := 0, 0, 0
	// new visual current head
	current := head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		// decimal computation
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	return head.Next
}
```