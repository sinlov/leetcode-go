package _00090_subsets_ii

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	sort.Ints(nums) // to remove duplicates

	for k := 0; k <= len(nums); k++ {
		generateSubsetsWithDup(nums, k, 0, c, &res)
	}
	return res
}

func generateSubsetsWithDup(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k { // dim recursion outer len(c) == k
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}

	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		// fmt.Printf("i = %v start = %v c = %v\n", i, start, c)
		if i > start && nums[i] == nums[i-1] { // duplicate numbers will be fetched this time. The next loop may duplicate numbers
			continue
		}
		c = append(c, nums[i])
		generateSubsetsWithDup(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}
