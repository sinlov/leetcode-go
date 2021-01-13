### 解题思路

答案中不可以包含重复的三元组, 故去重和排序也可以降低复杂度
遇到组合，第一反应对撞指针解法，当然需要跨过重复的，降低没必要的计算

步奏：
- 定义输出 sort 排序
- 先排除肯定不会出现满足条件的三元组情况
- 定义双指针 索引
- 从第二个开始遍历
 - 每次遍历重置双指针
 - 跳过重复的值
 - 双指针碰撞
   - 指针位移的时候，头指针，和尾指针跳过重复的值
   - 计算每次碰撞的三元组合
    - 如果三元组合 等于0 就记录在输出
    - 如果三元组合大于 0 因为已经排序过了，只需要尾指针位移
    - 如果三元组合小于 0 那么就是头指针位移
- 输出最终结果

### 代码

```go
func threeSum(nums []int) [][]int {
	res := make([][]int, 0) // result
	length := len(nums)

	if length < 3 { //ignore less then 3
		return res
	}
	sort.Ints(nums)                                        // sort by go std pkg
	if length > 0 && (nums[0] > 0 || nums[length-1] < 0) { // ignore must be 3-sum to 0， which sorted
		return res
	}

	start, end, index := 0, 0, 0 // double point and index

	for index = 1; index < length-1; index++ {
		start, end = 0, length-1                       // each range reset
		if index > 1 && nums[index] == nums[index-1] { // skip repetition, which is sorted
			start = index - 1
		}
		for start < index && end > index { // point crash
			if start > 0 && nums[start] == nums[start-1] { // repetition number
				start++ // start to next
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] { //repetition number
				end-- // end to previous
				continue
			}
			sumNumber := nums[start] + nums[end] + nums[index] // 3-sum
			if sumNumber == 0 {                                //
				res = append(res, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if sumNumber > 0 { // to previous
				end--
			} else { // to next
				start++
			}
		}
	}

	return res
}
```
