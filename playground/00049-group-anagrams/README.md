### 解题思路

- 每个字符串转换成 runes， 然后对 runes 排序
- 排序的 runes 同样的记录在 map 中
- 遍历 map 给最后的返回数组即可

### 代码

```go
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	record := map[string][]string{}
	for _, str := range strs {
		sr := []rune(str)
		sort.Sort(sortRunes(sr))
		tmp := record[string(sr)]
		tmp = append(tmp,str)
		record[string(sr)] = tmp
	}

	for _, v := range record {
		res = append(res, v)
	}

	return res
}
```
