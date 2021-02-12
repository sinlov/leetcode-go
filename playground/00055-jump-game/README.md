### 解题思路

- 作为 `起跳点` 的格子可以跳跃的距离是 `n`，那么表示后面 `n` 个格子都可以作为 `起跳点`
- 对每一个能作为 `起跳点` 的格子都尝试跳一次
  - 不断搜索 跳到最远的距离 `max`
    - 如果可以一直跳到最后，就成功
    - 中间有一个点比 `max` 还要大，说明在这个点和 max 中间连不上, 那么就达不到

### 代码

```go
func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 { // special case
		return false
	}

	if n == 1 { // special case
		return true
	}

	max := 0
	for i, v := range nums {
		if i > max {
			return false
		}
		max = greater(max, i+v)
	}
	return true
}

func greater(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```
