package _00069_sqrtx

func mySqrt(x int) int {
	if x == 0 { // special case
		return 0
	}

	left, right := 1, x
	res := 0
	for left <= right {
		mid := left + ((right - left) >> 1) // mid
		if mid < x/mid {                    // to right
			left = mid + 1
			res = mid
		} else if mid == x/mid {
			return mid
		} else { // to left
			right = mid - 1
		}
	}
	return res
}
