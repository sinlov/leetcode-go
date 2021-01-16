package _00018_4sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i+2; j-- {
			m, n := i+1, j-1
			for m < n {
				sum := nums[i] + nums[j] + nums[m] + nums[n]
				if sum == target {
					numArr := []int{nums[i], nums[m], nums[n], nums[j]}
					flag := true
					for _, v := range res {
						if v[0] == numArr[0] && v[1] == numArr[1] && v[2] == numArr[2] && v[3] == numArr[3] {
							flag = false
						}
					}
					if flag {
						res = append(res, numArr)
					}
					m, n = m+1, n-1
				} else if sum < target {
					m++
				} else {
					n--
				}
			}
		}

	}
	return res
}
