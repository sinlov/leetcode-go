### 解题思路

- 类似 78 题，多一个条件,数组中的数字会出现重复
- 方法还是 DFS

### 代码

```go
func subsetsWithDup(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	sort.Ints(nums) // to remove duplicates

	for k := 0; k <= len(nums); k++ {
		generateSubsetsWithDup(nums, k, 0, c, &res)
	}
	return res
}

func generateSubsetsWithDup(nums []int, k, start int, c []int, res *[][]int) {
	if len(c) == k { // dim recursion outer len(c) == k
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}

	for i := start; i < len(nums)-(k-len(c))+1; i++ {
		if i > start && nums[i] == nums[i-1] { // duplicate numbers will be fetched this time. The next loop may duplicate numbers
			continue
		}
		c = append(c, nums[i])
		generateSubsetsWithDup(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}
```
