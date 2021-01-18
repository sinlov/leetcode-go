package _00035_search_insert_position

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
