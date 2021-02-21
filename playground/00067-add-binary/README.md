### 解题思路

- 先对齐字符串，长度更长的在前，短的，前面补0
- 设置 进位标识 carry = 0 输出 res
- 逆序循环，每次计算 sum = aInt + bInt + carry
  - sum == 3 进位 补1
  - sum == 2 进位 补0
  - 其他情况，直接占位，并将 carry 置为 0
- 最终如果 carry == 1 也就是还需要进位，那么进位 补1

### 代码

```go
func addBinary(a string, b string) string {
	// let len(a) > len(b)
	if len(a) < len(b) {
		a, b = b, a
	}
	var complementZeroLength = len(a) - len(b)
	if len(a) != len(b) { // complement zero
		for i := 0; i < complementZeroLength; i++ {
			b = "0" + b
		}
	}

	var carry = 0 // carry flat if 1 need carry on
	var res = ""
	for i := len(a) - 1; i >= 0; i-- { // reverse
		aInt := a[i] - '0'
		bInt := b[i] - '0'
		sum := int(aInt+bInt) + carry
		if sum == 3 { // need carry on
			res = "1" + res
			carry = 1
		} else if sum == 2 { // need carry on
			res = "0" + res
			carry = 1
		} else {
			res = strconv.Itoa(sum) + res
			carry = 0
		}
	}

	if carry == 1 { // last flat need to carry on
		res = "1" + res
	}

	return res
}
```
