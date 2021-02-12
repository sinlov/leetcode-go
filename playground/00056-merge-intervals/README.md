### 解题思路

- 按照 `区间起点进行排序`
- 从区间起点小的开始扫描，依次合并每个有重叠的区间

### 代码

```go
func merge(intervals [][]int) [][]int {
	length := len(intervals)
	if length == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
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
```
