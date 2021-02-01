### 解题思路

条件

- `1 <= candidates.length <= 30`
- `1 <= candidates[i] <= 200`
- candidate 中的每个元素都是独一无二的
- `1 <= target <= 500`

思路

- 排除特殊情况
- 对 candidates 排序，可以减少没必要的遍历
- 回溯 查找符合 target 要求的数组
  - 符合条件 index == target 就 append cur[] 并写入 res
  - 回溯出口是 index > target
  - 注意回溯的时候，index 不需要更新，因为 candidate 中的元素可以取多次

### 代码

```go
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	sort.Ints(candidates)
	res, cur := [][]int{}, []int{} // cur cache each find list
	rangeCombinationSum(&res, candidates, target, 0, cur, 0)
	return res
}

func rangeCombinationSum(res *[][]int, num []int, target, index int, cur []int, j int) {
	if index == target {
		newcopy := make([]int, len(cur))
		copy(newcopy, cur)
		*res = append(*res, newcopy)
		return
	}
	if index > target { // range outer
		return
	}
	for i := j; i < len(num); i++ {
		index = index + num[i]
		cur = append(cur, num[i])
		j = i                                               // mark index to next
		rangeCombinationSum(res, num, target, index, cur, j) // index not change, beacuse element can be taken multiple times
		index = index - num[i]
		cur = cur[:len(cur)-1]
	}
}

```
