### 解题思路

- i j 两个游标，从左到右遍历
- i 外层游标，0开始 
- 同时不满足 j 的长度 等于 needle，也不满足， i+j 的长度等于 haystack，并且当前 j 游标能匹配上 i + j 的 haystack 对应的字符时， i++
- j 内层游标每当一个字符匹配上，就 j++ 继续匹配
- 直到找到 j 的长度正好等于 needle 的长度时，正好 i 的值是目标
- 如果 i + j 的值已经找满 haystack 的长度时， 返回 -1 

### 代码

```go
func strStr(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}
```
