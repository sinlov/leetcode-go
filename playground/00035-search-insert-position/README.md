### 解题思路

- 二分搜索
- 满足低或者高的情况，分开搜索
- 满足二分中间的小标，等于数组最大值或者下标的后一位值大于等于目标，那么返回 min 下标+1

### 代码

```go
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)>>1 // mid
		if nums[mid] >= target {   // move to low
			high = mid - 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) { // out for at this
				return mid + 1
			}
			low = mid + 1 // move to high
		}
	}
	return 0
}
```
