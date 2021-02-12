package _00056_merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length == 0 {
		return intervals
	}
	sort.Sort(intervalSS(intervals))
	res := make([][]int, 0)

	tmp := []int{}
	for _, s := range intervals {
		if len(res) == 0 || tmp[1] < s[0] {
			res = append(res, s)
			tmp = s
		} else if tmp[1] < s[1] {
			tmp[1] = s[1]
		}
	}

	return res
}

type intervalSS [][]int

func (s intervalSS) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s intervalSS) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s intervalSS) Len() int {
	return len(s)
}
