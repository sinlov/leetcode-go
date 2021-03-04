package _00078_subsets

func subsets(nums []int) [][]int {
	res := make([][]int, 1)

	for i := range nums { // top level
		for _, row := range res { // range res for append more
			clone := make([]int, len(row), len(row)+1)
			copy(clone, row) // copy clone
			clone = append(clone, nums[i])
			res = append(res, clone)
		}
	}
	return res
}
