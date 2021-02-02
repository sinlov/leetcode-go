package _00040_combination_sum_ii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 { // special case
		return [][]int{}
	}

	sort.Ints(candidates) // sort to skip

	res, cur := [][]int{}, []int{} // cur cache each find list
	rangeCombinationSum(&res, candidates, target, 0, cur)
	return res
}

func rangeCombinationSum(res *[][]int, nums []int, target, index int, cur []int) {
	if target == 0 {
		newcopy := make([]int, len(cur))
		copy(newcopy, cur)
		*res = append(*res, newcopy)
		return
	}

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] { // skip at repeatable number
			continue
		}
		if target >= nums[i] {
			cur = append(cur, nums[i])
			rangeCombinationSum(res, nums, target-nums[i], i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
}
