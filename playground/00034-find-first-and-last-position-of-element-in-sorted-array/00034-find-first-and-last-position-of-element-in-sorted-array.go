package _00034_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 { // special case
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
