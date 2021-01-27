package _00033_search_in_rotated_sorted_array

func search(nums []int, target int) int {
	if len(nums) == 0 { // length zero can not found
		return -1 // can not found
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
