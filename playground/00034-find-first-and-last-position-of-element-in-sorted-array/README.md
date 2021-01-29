### 解题思路

- 条件反射，二分搜索
- 排除特殊情况，减少搜索压力

### 代码

```go
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 { // special case
		return []int{-1, -1}
	}
	firstEqual := searchFirstEqual(nums, target)
	if firstEqual == -1 {
		return []int{-1, -1}
	}
	return []int{searchFirstEqual(nums, target), searchLastEqual(nums, target)}
}

// time complexity O(logn)
func searchLastEqual(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target { // last equal with target
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// time complexity O(logn)
func searchFirstEqual(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) { // frist equal with target
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}
```

- 如果有类似比较, 有二分查找方法

```go
// time complexity O(logn)
func searchFristGreater(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] < target { // first greater equal than target
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func searchLastLess(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] <= target {
			if (mid == len(nums)-1) || (nums[mid+1] > target) { // last less equal than target
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
```
