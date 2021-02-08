### 解题思路

- 递归的方式，不断的将 n/2 分下去
- 注意 n 的正负数，n 的奇偶性二分问题

### 代码

```go
func myPow(x float64, n int) float64 {
	if n == 0 { // special case
		return 1
	}

	if n == 1 { // special case
		return x
	}

	if n < 0 { // if n ltz just change x
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 { // even number
		return tmp * tmp
	}
	return tmp * tmp * x // odd number
}
```
