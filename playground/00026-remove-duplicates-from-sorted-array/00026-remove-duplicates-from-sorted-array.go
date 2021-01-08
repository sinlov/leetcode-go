package _00026_remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count, summary := 0, 0
	for count < len(nums)-1 {
		for nums[count] == nums[summary] {
			summary++
			if summary == len(nums) {
				return count + 1
			}
		}
		nums[count+1] = nums[summary]
		count++
	}

	return count + 1
}
