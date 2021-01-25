package _00031_next_permutation

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
