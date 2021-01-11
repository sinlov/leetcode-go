package _00005_longest_palindromic_substring

import "strings"

var (
	separator     = "#"
	separatorChar = []rune(separator)[0]
)

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	id, mPoint, rPoint := 0, 0, 0
	manacherStr := toManacherStr(s)

	var path = make([]int, len(manacherStr))
	for i, _ := range manacherStr {
		if i < rPoint {
			lg := path[2*id-i]
			if rPoint-i < lg {
				lg = rPoint - i
			}
			path[i] = lg
		} else {
			path[i] = 1
		}

		for i-path[i] >= 0 && i+path[i] < len(manacherStr) &&
			manacherStr[i-path[i]] == manacherStr[i+path[i]] {
			path[i] += 1
		}

		if path[i] > path[mPoint] {
			mPoint = i
		}

		if i+path[i] > rPoint {
			rPoint = i + path[i]
			id = i
		}
	}

	return strings.Replace(manacherStr[(mPoint-path[mPoint]+1):(mPoint+path[mPoint])], separator, "", -1)
}

func toManacherStr(s string) string {
	ans := make([]rune, 0)
	ans = append(ans, separatorChar)
	for i := 0; i < len(s); i++ {
		ans = append(ans, rune(s[i]), separatorChar)
	}
	return string(ans)
}
