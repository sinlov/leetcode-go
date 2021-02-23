### 解题思路

- 多少方法，第一反应 DP
  - 初始条件 第一 和 第二 都是 爬 1 层
  - 方程为 一个楼梯 可以由 `i-1` 和 `i-2` 的楼梯爬上来
    - `dp[i] = dp[i-1] + dp[i-2]`
- 最终 dp[n] 即是对应的方法数

### 代码

```go
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
```
