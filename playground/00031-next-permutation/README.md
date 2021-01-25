### 解题思路

- 排除特殊情况
- 重新排列组合第n-1个到最后一个可以得到更大的排列
- 从组合的最后一个数开始找到第一个比下标为n-1的大的数
- 将该数与下标为n-1的数交换位置
- 然后开始二分搜索，交换下标为n的数到最后一个数

### 代码

```go
func nextPermutation(nums []int) {
	numsLen := len(nums)
	if numsLen < 2 { // special case
		return
	}

	left, right := numsLen-2, numsLen-1
	for left >= 0 && nums[left] >= nums[right] { // swap
		left--
		right--
	}

	if left >= 0 { // left >= 0
		i := numsLen - 1
		for nums[i] <= nums[left] {
			i--
		}
		nums[i], nums[left] = nums[left], nums[i] // swap
	}

	start, end := right, numsLen-1
	for start < end { // binary search
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
```
