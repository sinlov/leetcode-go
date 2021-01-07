### 解题思路

先处理空串

`滑动窗口` 思路
- 滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界
- 出现了重复字符，就需要缩小左边界，直到重复的字符移出了左边界
- 以此类推

长度计数

- 每次移动需要计算当前长度
- 判断是否需要更新最大长度
- 最终返回计算的长度值即可

### 代码

```go
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

```
