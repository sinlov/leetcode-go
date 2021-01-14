### 解题思路

- 遇到合 第一反应 对撞指针
- 并且排序对最值问题有时间优化效果
- 当然数组有重复，遍历过程需要排除

求解过程

- 先排除长度小于 3 情况
- 对数组排序
- 初始化一个接近值
- 然后使用对撞指针计算
  - 每次计算先刷新 输出结果合接近值，并比较3数合和目标的大小
    - 大于目标，尾指针移动
    - 小于目标，投指针移动
    - 等于目标，就直接输出
  - 实在找不到就输出3数合的值

### 代码

```go
func threeSumClosest(nums []int, target int) int {
	res := 0
	lN := len(nums)
	if lN < 3 { // do less then 3
		return res
	}

	diffVal := math.MaxInt32    // init diff value
	sort.Ints(nums)             // sort
	for i := 0; i < lN-2; i++ { // only need range lN -2
		for j, k := i+1, lN-1; j < k; { // crash point
			sum := nums[i] + nums[j] + nums[k] // sum now
			if absInt(sum-target) < diffVal {  // each diff set by absOf target
				res, diffVal = sum, absInt(sum-target)
			}
			if sum == target {
				return res
			} else if sum > target {
				k-- // tail move
			} else {
				j++ // head move
			}
		}
	}

	return res
}

func absInt(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```
