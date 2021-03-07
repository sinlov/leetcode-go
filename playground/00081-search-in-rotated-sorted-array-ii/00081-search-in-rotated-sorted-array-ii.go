package _00081_search_in_rotated_sorted_array_ii

func search(nums []int, target int) bool {
	numLen := len(nums)
	if numLen == 0 { //special case
		return false
	}
	low, high := 0, len(nums)-1
	for low <= high { // binary search
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[low] { // at high
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // at low
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			}
			if nums[high] == nums[mid] {
				high--
			}
		}
	}
	return false
}
