### 解题思路

- 动态规划求解

- dp[n] 代表翻译长度为 n 个字符的字符串的方法总数
  - 由于题目中的数字可能出现 0，0 不能翻译成任何字母，所以出现 0 要跳过

- dp[0] 代表空字符串，只有一种翻译方法, dp[0] = 1
- dp[1] 需要考虑原字符串是否是 0 开头的
  - 如果是 0 开头的，dp[1] = 0
  - 如果不是 0 开头的，dp[1] = 1

- 状态转移方程

```
if (1 ≤ s[i-1 : i] ≤ 9)
  dp[i] += dp[i-1]

if (10 ≤ s[i-2 : i] ≤ 26)
  dp[i] += dp[i-2]
```

- 最终结果是 dp[n]

### 代码

```go
func numDecodings(s string) int {
	if len(s) == 0 { // special case
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[:1] == "0" { // is start as 0
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 2; i <= len(s); i++ {
		lastNum, _ := strconv.Atoi(s[i-1 : i])
		if lastNum >= 1 && lastNum <= 9 {
			dp[i] += dp[i-1]
		}
		lastNum, _ = strconv.Atoi(s[i-2 : i])
		if lastNum >= 10 && lastNum <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)] // dp[n]
}

```
