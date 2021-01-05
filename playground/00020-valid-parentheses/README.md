### 解题思路
使用栈思路
- 如果是左边括号 `([{`，就进栈
- 遇到右边括号 `)]}` 并且栈顶 `stack[len(stack)-1]` 为对应的左括号，那么栈顶出栈 `stack = stack[:len(stack)-1]`
- 如果遍历整个字符串，最终栈为空 `len(stack) == 0`，那么就是配对上的

### 代码

```go
func isValid(s string) bool {
	// empty is true
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') { // push
			stack = append(stack, v)
		} else {

			if (v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') || // valid ()
				(v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') || // valid [ ]
				(v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') { // valid { }
				stack = stack[:len(stack)-1]
			} else {
				return false // not valid
			}
		}
	}

	return len(stack) == 0 // final 0 is valid
}

```
