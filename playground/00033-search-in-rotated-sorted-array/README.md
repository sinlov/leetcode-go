### 解题思路

- nums 中的每个值都 独一无二, 肯定会在某个点上旋转
- 那么 二分搜索就可以直接找，找到就返回下标， 没找到返回 -1
  - 找到目标的条件为 `nums[mid] == target`
  - 如果 mid 落在了前一段数值比较大的区间内 ，则一定有 `nums[mid] > nums[low]`
    - target 在  nums[low] 到 nums[mid] 之间， 则分到 low 的部分
    - 反之则，二分到 high 部分
  - 如果 mid 落在了后一段数值比较小的区间内，则一定有 `nums[mid] < nums[high]`
    - target 在 nums[mid] 到 nums[higt] 之间，则分到 higt 部分
    - 反之，则二分到 low 部分
  - `nums[low] == nums[mid]` low 移动
  - `nums[high] == nums[mid]` high 移动
- 还是不满足 `nums[mid] == target` 返回 -1

### 代码

```go
func search(nums []int, target int) int {
	if len(nums) == 0 { // length zero can not found
		return -1
	}

	low, high := 0, len(nums)-1 // binary search start
	for low <= high {
		mid := (high-low)>>1 + low // each mid index
		if nums[mid] == target {   // find out target index
			return mid
		} else if nums[mid] > nums[low] { // index part move
			if nums[low] <= target && target < nums[mid] { // move low part
				high = mid - 1
			} else { // move to hihgt part
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // index move
			if nums[mid] < target && target <= nums[high] { // move low when target gratter
				low = mid + 1
			} else { // move hight target less
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] { // move low at mid
				low++
			}
			if nums[high] == nums[mid] { // move high at mid
				high--
			}
		}
	}

	return -1 // not found
}
```
