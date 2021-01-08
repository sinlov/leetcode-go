### 解题思路

> 注意: 输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的

- 已经是排序好的数组
- 那么删除并不是真的删除，将删除的元素移动到数组后面的空间内，并返回数组实际剩余的元素个数
- 最终判断题目的时候会读取数组剩余个数的元素进行输出

步骤:

- 处理空数组，这个基本情况
- 定义2个下标 计数和游标，count summary
- 循环到计数器到数组倒数1个位置
- 每次循环时
  - 游标和当前位置不断循环，每次游标增加，当找到相同的内容时，count加一
  - 将后面的值，不断循环，不再出现重复值时，返回count + 1
  - 循环到游标内容相等时，移动元素到紧挨着的位置（等效把重复的挤到后面）
  - 最终也是返回count + 1

### 代码

```go
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count, summary := 0, 0
	for count < len(nums)-1 {
		for nums[count] == nums[summary] {
			summary++
			if summary == len(nums) {
				return count + 1
			}
		}
		nums[count+1] = nums[summary]
		count++
	}

	return count + 1
}
```
