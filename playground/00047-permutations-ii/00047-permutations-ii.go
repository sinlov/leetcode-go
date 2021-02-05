package _00047_permutations_ii

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 { // special case
		return [][]int{}
	}

	sort.Ints(nums) // for remove duplicates

	res := [][]int{}                // final result
	cur := []int{}                  // current cache
	used := make([]bool, len(nums)) // mark used
	searchPermuteUnique(&res, &used, 0, nums, cur)
	return res
}

func searchPermuteUnique(res *[][]int, used *[]bool, index int, nums []int, cur []int) {
	if index == len(nums) { // search out
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] { // remove duplicates
				continue
			}
			(*used)[i] = true
			cur = append(cur, nums[i])
			searchPermuteUnique(res, used, index+1, nums, cur) // next index search
			cur = cur[:len(cur)-1]
			(*used)[i] = false
		}
	}
}
