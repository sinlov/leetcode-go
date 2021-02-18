### 解题思路

> 求若干小问题的最大值，那么直接使用动态规划

- 使用 动态规划
- dp[i] 表示 [0,i] 区间内 各个子区间和的最大值
  - 状态转移
    - if dp[i-1] > 0 then dp[i] = nums[i] + dp[i-1]
    - if dp[i-1] ≤ 0 then dp[i] = nums[i]

### 代码

```go
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// Dynamic Programming
	dp := make([]int, len(nums))
	res := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = greater(res, dp[i])
	}

	return res
}

func greater(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

```
