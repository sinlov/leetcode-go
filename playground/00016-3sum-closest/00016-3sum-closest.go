package _00016_3sum_closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	res := 0
	lN := len(nums)
	if lN < 3 { // do less then 3
		return res
	}

	diffVal := math.MaxInt32    // init diff value
	sort.Ints(nums)             // sort
	for i := 0; i < lN-2; i++ { // only need range lN -2
		for j, k := i+1, lN-1; j < k; { // crash point
			sum := nums[i] + nums[j] + nums[k] // sum now
			if absInt(sum-target) < diffVal {  // each diff set by absOf target
				res, diffVal = sum, absInt(sum-target)
			}
			if sum == target {
				return res
			} else if sum > target {
				k-- // tail move
			} else {
				j++ // head move
			}
		}
	}

	return res
}

func absInt(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
