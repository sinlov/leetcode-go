package _00015_3sum

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0) // result
	length := len(nums)

	if length < 3 { //ignore less then 3
		return res
	}
	sort.Ints(nums)                         // sort by go std pkg
	if length > 0 && (nums[length-1] < 0) { // ignore must be 3-sum to 0， which sorted max number less than 0
		return res
	}

	start, end, index := 0, 0, 0 // double point and index

	for index = 1; index < length-1; index++ {
		start, end = 0, length-1                       // each range reset
		if index > 1 && nums[index] == nums[index-1] { // skip repetition, which is sorted
			start = index - 1
		}
		for start < index && end > index { // point crash
			if start > 0 && nums[start] == nums[start-1] { // repetition number
				start++ // start to next
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] { //repetition number
				end-- // end to previous
				continue
			}
			sumNumber := nums[start] + nums[end] + nums[index] // 3-sum
			if sumNumber == 0 {                                // mark to res
				res = append(res, []int{nums[start], nums[index], nums[end]})
				start++ // each crash point move all
				end--
			} else if sumNumber > 0 { // to previous
				end--
			} else { // to next
				start++
			}
		}
	}

	return res
}
