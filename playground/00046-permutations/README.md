### 解题思路

- DFS 深度优先搜索

### 代码

```go
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	res := [][]int{}                // final result
	cur := []int{}                  // current cache
	used := make([]bool, len(nums)) // mark used
	searchPermute(&res, &used, 0, nums, cur)
	return res
}

func searchPermute(res *[][]int, used *[]bool, index int, nums []int, cur []int) {
	if index == len(nums) { // search out
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			(*used)[i] = true
			cur = append(cur, nums[i])
			searchPermute(res, used, index+1, nums, cur) // next index search
			cur = cur[:len(cur)-1]
			(*used)[i] = false
		}
	}
}
```
