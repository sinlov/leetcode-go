package _00003_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var slidingWindow [512]int      // 0 <= s.length <= 5 * 10^4
	result, left, right := 0, 0, -1 // mark

	for left < len(s) {
		if right+1 < len(s) && slidingWindow[s[right+1]-'a'] == 0 { // char ASCII begins with a
			slidingWindow[s[right+1]-'a']++
			right++ // sliding right
		} else {
			slidingWindow[s[left]-'a']--
			left++ // sliding left
		}

		checkResult := right - left + 1
		if result <= checkResult { // update max when gather equal
			result = checkResult
		}
	}
	return result
}
