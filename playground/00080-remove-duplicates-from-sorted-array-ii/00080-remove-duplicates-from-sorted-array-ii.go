package _00080_remove_duplicates_from_sorted_array_ii

func removeDuplicates(nums []int) int {
	numLen := len(nums)
	if numLen == 0 { // special case
		return 0
	}
	nowMark, last := 0, 0
	for last < numLen-1 {
		startMark := -1
		for nums[nowMark] == nums[last] { // range to mark repetition num
			if startMark == -1 || startMark > nowMark {
				startMark = nowMark
			}
			if nowMark == numLen-1 {
				break
			}
			nowMark++
		}
		if nowMark-startMark >= 2 &&
			nums[nowMark-1] == nums[last] &&
			nums[nowMark] != nums[last] { // this case find max 2 num
			nums[last+1] = nums[nowMark-1]
			nums[last+2] = nums[nowMark]
			last += 2
		} else {
			nums[last+1] = nums[nowMark]
			last++
		}
		if nowMark == numLen-1 { // last one find
			if nums[nowMark] != nums[last-1] {
				nums[last] = nums[nowMark]
			}
			return last + 1
		}
	}
	return last + 1 // final result is last + 1 when nums length 1
}
