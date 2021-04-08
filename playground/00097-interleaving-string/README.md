### 解题思路

DP 求解

- 排除特殊情况 s1 || s2 为空时，直接判断
- 排除 `len(s1) + len(s2) != len(s3)` 这种情况

- 初始化一个 `s1 + 1 : s2 + 1` 的矩阵，矩阵中 fasle 表示不符合，true 表示符合
- 动态方程求解为

```
  if s1[i-1] == s3[i+j-1] && s2[j-1] == s3[i+j-1]
    dp[i][j] = dp[i-1][j] || dp[i][j-1]

  if s1[i-1] == s3[i+j-1] && s2[j-1] != s3[i+j-1]
    dp[i][j] = dp[i-1][j]

  if s1[i-1] != s3[i+j-1] && s2[j-1] == s3[i+j-1]
    dp[i][j] = dp[i][j-1]
```

- 最终求解是否符合为 `dp[len(s1)][len(s2)]`

### 代码

```go
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 { // special case
		return s2 == s3
	}
	if len(s2) == 0 { // special case
		return s1 == s3
	}
	if len(s1)+len(s2) != len(s3) { // size not s1 + s2 = s3
		return false
	}

	dp := make([][]bool, len(s1)+1)  // dp len is s1+1
	for i := 0; i < len(s1)+1; i++ { // init dp as s1 : s2
		dp[i] = make([]bool, len(s2)+1) // length of dp[i] is len(s2)+1
	}

	dp[0][0] = true
	for i := 1; i < len(dp); i++ { // from 1 start update matrix column
		if s1[i-1] == s3[i-1] { // update status by s1
			dp[i][0] = dp[i-1][0]
		}
	}
	for j := 1; j < len(dp[0]); j++ { // from 1 start update matrix row
		if s2[j-1] == s3[j-1] { // update status by s2
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if dp[i-1][j] && s1[i-1] == s3[i+j-1] {
				dp[i][j] = true
			}
			if dp[i][j-1] && s2[j-1] == s3[i+j-1] {
				dp[i][j] = true
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
```
