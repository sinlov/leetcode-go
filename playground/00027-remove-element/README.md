### 解题思路

- 还是先处理长度 0
- 加入移除后的计数，用于返回
- 删除并不是真的删除，将删除的元素移动到数组后面的空间内，并返回数组实际剩余的元素个数

### 代码

```go
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != count {
				nums[i], nums[count] = nums[count], nums[i]
			}
			count++
		}
	}
	return count
}
```
