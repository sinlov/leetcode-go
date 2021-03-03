### 解题思路

- 计数排序
分别记录 0 1 2 出现的，根据条件，不断记录即可

### 代码

```go
func sortColors(nums []int) {
	if len(nums) == 0 { //special case
		return
	}
	red, white, blue := 0, 0, 0 // the end of different colors
	for _, num := range nums {
		if num == 0 { // red color
			nums[blue] = 2
			blue++
			nums[white] = 1
			white++
			nums[red] = 0
			red++
		} else if num == 1 { // white color
			nums[blue] = 2
			blue++
			nums[white] = 1
			white++
		} else if num == 2 { // blue color
			blue++
		}
	}
}

```
