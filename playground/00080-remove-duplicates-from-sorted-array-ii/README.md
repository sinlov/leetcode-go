### 解题思路

- 第 26 题的加强版
- 这题是去除重复的元素，使得每个元素最多出现两次

> 这里数组的删除并不是真的删除，只是将删除的元素移动到数组后面的空间内，然后返回数组实际剩余的元素个数

- 先记录数组长度 numLen
  - numLen == 0 直接 return
- 初始化当前标记 nowMark, last 都为 0, 0
- 循环 `last < numLen-1`
  - 每次初始化开始标记 startMark := -1
  - 循环 `nums[nowMark] == nums[last]` 也就是找不一样数字
    - 记录最新的 开始下标 和不一样数值的下标
  - 满足条件 `nowMark-startMark >= 2 && nums[nowMark-1] == nums[last] && nums[nowMark] != nums[last]` 也就是去重操作
      - last 增2个
    - 不满足条件
      - last 增1个
  - 确认是不是最后一个标记 nowMark == numLen-1
    - nums[nowMark] != nums[last-1] 满足这个条件还要去除一次
  - 最终返回 last + 1
- 不满足循环条件 `last < numLen-1` 也就是 nums 只有一个元素，直接返回 last + 1

### 代码

```go
func removeDuplicates(nums []int) int {
	numLen := len(nums)
	if numLen == 0 { // special case
		return 0
	}
	nowMark, last := 0, 0
	for last < numLen-1 {
		startMark := -1
		for nums[nowMark] == nums[last] { // range to mark repetition num
			if startMark == -1 || startMark > nowMark {
				startMark = nowMark
			}
			if nowMark == numLen-1 {
				break
			}
			nowMark++
		}
		if nowMark-startMark >= 2 &&
			nums[nowMark-1] == nums[last] &&
			nums[nowMark] != nums[last] { // this case find max 2 num
			nums[last+1] = nums[nowMark-1]
			nums[last+2] = nums[nowMark]
			last += 2
		} else {
			nums[last+1] = nums[nowMark]
			last++
		}
		if nowMark == numLen-1 { // last one find
			if nums[nowMark] != nums[last-1] {
				nums[last] = nums[nowMark]
			}
			return last + 1
		}
	}
	return last + 1 // final result is last + 1 when nums length 1
}
```
