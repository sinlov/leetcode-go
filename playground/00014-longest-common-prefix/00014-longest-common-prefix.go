package _00014_longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	tmp := ""
	minStrCount := -1
	minStr := ""
	isSub := true
	for _, str := range strs {
		if minStrCount == -1 {
			minStrCount = len(str)
			minStr = str
		} else {
			if minStrCount > len(str) {
				minStrCount = len(str)
				minStr = str
			}
		}
	}
	if minStrCount == 0 {
		return ""
	}
	if minStrCount == 1 {
		for _, str := range strs {
			if !(str[0:1] == minStr) {
				isSub = false
				break
			} else {
				isSub = true
			}
		}
		if isSub {
			tmp = minStr
		}
	} else {
		for i := 1; i < minStrCount+1; i++ {
			substr := minStr[0:i]
			for _, str := range strs {
				if !(str[0:i] == substr) {
					isSub = false
					break
				} else {
					isSub = true
				}
			}
			if isSub {
				if len(substr) > len(tmp) {
					tmp = substr
				}
			}
		}
	}

	return tmp
}
