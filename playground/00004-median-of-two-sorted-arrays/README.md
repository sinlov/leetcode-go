### 解题思路

- 二分搜索

如何有效切分 数组1 ?

需要满足

1. 线左边的数都小于右边的数，也就是比较大小并换位
2. 满足下面的条件

```
nums1[midA-1] ≤ nums2[midB] && nums2[midB-1] ≤ nums1[midA]
```

如果这些条件都不满足，切分线就需要调整

- 如果 `nums1[midA] < nums2[midB-1]`，说明 midA 这条线划分出来左边的数小了，切分线应该右移
- 如果 `nums1[midA-1] > nums2[midB]`，说明 midA 这条线划分出来左边的数大了，切分线应该左移

找到切分线后

- `数组1` 在切分线两边的下标分别是 `midA - 1` 和 `midA`
- `数组2` 在切分线两边的下标分别是 `midB - 1` 和 `midB`

合并成最终数组时

- 如果数组长度是奇数，那么 中位数 就是 `max(nums1[midA-1], nums2[midB-1])`
- 如果数组长度是偶数，那么 中间位置的两个数

  依次是：`max(nums1[midA-1], nums2[midB-1])` 和 `min(nums1[midA], nums2[midB])`

  那么中位数就是 `(max(nums1[midA-1], nums2[midB-1]) + min(nums1[midA], nums2[midB])) / 2`

### 代码

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) { // use short array
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high { // binary search
		nums1Mid = low + (high-low)>>1 // left is mid -1, right is mid
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // move to left
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // move to right
			low = nums1Mid + 1
		} else { // find mid
			break
		}
	}

	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = greater(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}

	if (len(nums1)+len(nums2))&1 == 1 { // odd number
		return float64(midLeft)
	}

	// even number
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = lesser(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func greater(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lesser(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
```
