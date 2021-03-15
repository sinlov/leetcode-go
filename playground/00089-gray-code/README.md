### 解题思路

格雷码 生成规则

- 以二进制为0值的格雷码为第零项
  - 第一次改变最右边的位元
  - 第二次改变右起第一个为1的位元的左边位元
  - 第三、四次方法同第一、二次
  - 直到编码完成

使用递归求解

### 代码

```go
func grayCode(n int) []int {
	if n == 0 { // special case
		return []int{0}
	}
	res := []int{}
	num := make([]int, n)
	generateGrayCode(int(1<<uint(n)), 0, &num, &res)
	return res
}

func generateGrayCode(n, step int, num *[]int, res *[]int) {
	if n == 0 {
		return
	}
	*res = append(*res, convertBinary(*num))

	if step%2 == 0 {
		(*num)[len(*num)-1] = flipGrayCode((*num)[len(*num)-1])
	} else {
		index := len(*num) - 1
		for ; index >= 0; index-- {
			if (*num)[index] == 1 {
				break
			}
		}
		if index == 0 {
			(*num)[len(*num)-1] = flipGrayCode((*num)[len(*num)-1])
		} else {
			(*num)[index-1] = flipGrayCode((*num)[index-1])
		}
	}
	generateGrayCode(n-1, step+1, num, res)
	return
}

func convertBinary(num []int) int {
	res, rad := 0, 1
	for i := len(num) - 1; i >= 0; i-- {
		res += num[i] * rad
		rad *= 2
	}
	return res
}

// filp code
func flipGrayCode(num int) int {
	if num == 0 {
		return 1
	}
	return 0
}
```
