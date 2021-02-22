### 解题思路

- 牛顿迭代法 [https://en.wikipedia.org/wiki/Integer_square_root](https://en.wikipedia.org/wiki/Integer_square_root)
- Quake III 游戏引擎中有一种比 STL 的 sqrt 快 4 倍的实现 [https://en.wikipedia.org/wiki/Fast_inverse_square_root](https://en.wikipedia.org/wiki/Fast_inverse_square_root)

- 根号 x 的取值范围一定在 [0,x] 之间，这个区间内的值是
  - 递增有序的
  - 有边界的
  - 可以用下标访问的
- 那么满足 二分搜索的条件

### 代码

```go
func mySqrt(x int) int {
	if x == 0 { // special case
		return 0
	}

	left, right := 1, x
	res := 0
	for left <= right {
		mid := left + ((right - left) >> 1) // mid
		if mid < x/mid {                    // to right
			left = mid + 1
			res = mid
		} else if mid == x/mid {
			return mid
		} else { // to left
			right = mid - 1
		}
	}
	return res
}

```
