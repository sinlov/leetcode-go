package _00027_remove_element

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != count {
				nums[i], nums[count] = nums[count], nums[i]
			}
			count++
		}
	}
	return count
}
