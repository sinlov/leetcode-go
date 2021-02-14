### 解题思路

- 逆向遍历
  - 找到第一个不是空格的下标 tail ，然后再找到第一个是空格的下标 head
- tail - head 就是要的结果

### 代码

```go
func lengthOfLastWord(s string) int {
	if len(s) == 0 { // speace case
		return 0
	}
	tail := len(s) - 1
	for tail >= 0 && s[tail] == ' ' { // find out tail not space
		tail--
	}
	head := tail
	for head >= 0 && s[head] != ' ' { // find out head space
		head--
	}
	return tail - head
}
```
