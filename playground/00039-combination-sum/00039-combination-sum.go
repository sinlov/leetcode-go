package _00039_combination_sum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	sort.Ints(candidates)
	res, cur := [][]int{}, []int{} // cur cache each find list
	rangeCombinationSum(&res, candidates, target, 0, cur, 0)
	return res
}

func rangeCombinationSum(res *[][]int, num []int, target, index int, cur []int, j int) {
	if index == target {
		newcopy := make([]int, len(cur))
		copy(newcopy, cur)
		*res = append(*res, newcopy)
		return
	}
	if index > target { // range outer
		return
	}
	for i := j; i < len(num); i++ {
		index = index + num[i]
		cur = append(cur, num[i])
		j = i                                                // mark index to next
		rangeCombinationSum(res, num, target, index, cur, j) // index not change, beacuse element can be taken multiple times
		index = index - num[i]
		cur = cur[:len(cur)-1]
	}
}
