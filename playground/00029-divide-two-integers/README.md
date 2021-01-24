### 解题思路

- 二分搜索，把商作为二分搜索的目标
- 商的范围在 [0, dividend]
- 满足条件

```
  (商 + 1） * 除数 > 被除数 && 商 * 除数 <= 被除数
    or
  (商+1)* 除数 >= 被除数 && 商 * 除数 < 被除数
```

换成位运算，就是二分搜索 除数 极限逼近 被除数 前，搜索到一个值(商)，满足小于等于被除数按除数位移的逼近

### 代码

```go
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	res := 0 // quotient start at 0

	symbol := -1
	if dividend > 0 && divisor > 0 || dividend < 0 && divisor < 0 { // final symbol
		symbol = 1
	}

	absDivd, absDivs := abs(dividend), abs(divisor)

	for absDivd >= absDivs { // dichotomization
		tmp := absDivs
		mid := 1
		for tmp<<1 <= absDivd { // out (tmp << 1) <= dividend
			tmp <<= 1
			mid <<= 1
		}
		absDivd -= tmp
		res += mid
	}

	return symbol * res // final add symbol
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```
