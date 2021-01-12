### 解题思路

对撞指针的思路，首尾 2个指针，每次移动后判断 长宽的乘积是否最大

- 定义变量 max start end 分别是 最大乘积 开始指针位置 最终指针位置
- 两个指针对撞，每次对撞算宽度
- 对撞谁小就位移，并且算出当前高度
- 算乘积，比较 max
- 直到 指针重合 返回 max 即可

### 代码

```go
func maxArea(height []int) int {
	max, s, e := 0, 0, len(height)-1

	for s < e { // crash to true
		width := e - s
		high := 0
		if height[s] < height[e] { // min height s
			high = height[s]
			s++
		} else { // min height e
			high = height[e]
			e--
		}
		volume := width * high // now volume
		if volume > max {
			max = volume
		}
	}
	return max
}
```
