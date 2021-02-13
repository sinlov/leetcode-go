### 解题思路

模拟移动

- 排除特殊情况
- 本身数组是已经排线好的，设置 left right 左右下标 作为二维数组列下标，循环二维数组，清洗数据
  - 对 目标下标1 < 插入下标0 移动 left
  - 对 目标下标0 > 插入下标1 设置 rigth 并跳出循环
  - 之间的情况需要精确处理
    - 不相交，就直接跳过
    - 覆盖则，直接替换
    - 相交 右侧侧， 替换 插入下标1 使用 遍历下标1
    - 相交 右侧侧， 替换 插入下标0 使用 遍历下标0
最终被清洗后，拼接 左侧 中间 右侧的数据，作为最终数据

### 代码

```go
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
```
