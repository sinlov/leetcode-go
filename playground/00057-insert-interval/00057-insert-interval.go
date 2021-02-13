package _00057_insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 { // special case
		return [][]int{newInterval}
	}
	if len(newInterval) == 0 { // special case
		return intervals
	}
	left, right := -1, len(intervals) // start search
	for i, val := range intervals {
		if val[1] < newInterval[0] { // move left to i
			left = i
		} else if val[0] > newInterval[1] { // move right to i and do break looper
			right = i
			break
		} else {
			if newInterval[0] < val[0] && newInterval[1] > val[1] { // next
				continue
			} else if newInterval[0] >= val[0] && newInterval[1] <= val[1] { // replace newInterval full
				newInterval = val
			} else if newInterval[0] < val[0] { // covery right
				newInterval[1] = val[1]
			} else if newInterval[1] > val[1] { // covery left
				newInterval[0] = val[0]
			}
		}
	}
	res := make([][]int, 0)
	res = append(res, intervals[:left+1]...) // append left
	res = append(res, newInterval)           // append mid
	res = append(res, intervals[right:]...)  // append right
	return res
}
