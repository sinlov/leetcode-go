package _00055_jump_game

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 { // special case
		return false
	}

	if n == 1 { // special case
		return true
	}

	max := 0
	for i, v := range nums {
		if i > max {
			return false
		}
		max = greater(max, i+v)
	}
	return true
}

func greater(a, b int) int {
	if a < b {
		return b
	}
	return a
}
