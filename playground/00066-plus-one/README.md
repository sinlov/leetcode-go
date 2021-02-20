### 解题思路

- 记录当前进位 carry = 1
- 从尾部开始扫描
  - 每次循环判断是否需要进位 carry == 0 就直接退出循环
  - 不需要进位，digits[i] += carry， 并且将 carry = 0
  - 需要进位，当前下标值置为 0, 并且进位保持 carry = 1

- 如果 下标 0 还需要进位，需要在 0 下标前插入 1

### 代码

```go
func plusOne(digits []int) []int {
	length := len(digits)
	carry := 1 // default carry 1
	for i := length - 1; i >= 0; i-- {
		if digits[i]+carry > 9 { // need carry on
			digits[i] = 0
			carry = 1
		} else {
			digits[i] += carry
			carry = 0
		}
	}
	if digits[0] == 0 && carry == 1 { // subscript 0 need carry on
		digits = append([]int{1}, digits...)
	}
	return digits
}
```
