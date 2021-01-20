### 解题思路

- 排除特殊情况后
- 先对字符串左去除空格
- 设置标识位是否为正数
- 对去除左空格的字符串，第一个位置进行 字符串比较是否为 - 或者 + 或者不是数字符号直接输出
- 然后对 1 后面每轮循环判断
  - 如果不是数字，直接输出
  - 如果正数越界，输出 math.MaxInt32
  - 如果负数越界，输出 math.MinInt32
  - 满足条件，根据符号十进制进位计算

### 代码

```go
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	strpStr := stripLeftWhitespace(s)
	positive := true
	res := 0
	for i, char := range strpStr {
		num := int(char - '0')
		if i == 0 { // check symbol
			if char == '-' {
				positive = false
			} else if char == '+' {
				positive = true
			} else if num >= 0 && num <= 9 {
				res = res*10 + num
			} else {
				return res
			}
		} else {
			if num < 0 || num > 9 { // not number
				break
			}
			if positive && res*10+num > math.MaxInt32 { // max
				return math.MaxInt32
			}
			if !positive && res*10-num < math.MinInt32 {// min
				return math.MinInt32
			}
			if positive {
				res = res*10 + num
			} else {
				res = res*10 - num
			}
		}
	}
	return res
}

func stripLeftWhitespace(str string) string {
	res := ""
	headNotWhitespace := true
	for _, char := range str {
		if headNotWhitespace {
			if char != ' ' {
				headNotWhitespace = false
				res = res + string(char)
			}
		} else {
			res = res + string(char)
		}
	}
	return res
}
```
