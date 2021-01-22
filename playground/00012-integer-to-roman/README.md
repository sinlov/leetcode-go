### 解题思路

罗马数字转换字符特征位表

| 罗马字符 | 特征值数值 |
|:--------|:---------|
| I       | 1        |
| IV      | 4        |
| V       | 5        |
| IX      | 9        |
| X       | 10       |
| XL      | 40       |
| L       | 50       |
| XC      | 90       |
| C       | 100      |
| CD      | 400      |
| D       | 500      |
| CM      | 900      |
| M       | 1000     |

- 按左不断递归，复制一个值和特征值比较
- 如果当前值大于等于则
  - 减去特征值
  - 并在输出串的右边加入当前特征值对应的罗马数字

### 代码

```go
func intToRoman(num int) string {
	res := ""
	copy := num
	for copy != 0 {
		if copy >= 1000 {
			copy = copy - 1000
			res += "M"
		} else if copy >= 900 {
			copy -= 900
			res += "CM"
		} else if copy >= 500 {
			copy -= 500
			res += "D"
		} else if copy >= 400 {
			copy -= 400
			res += "CD"
		} else if copy >= 100 {
			copy -= 100
			res += "C"
		} else if copy >= 90 {
			copy -= 90
			res += "XC"
		} else if copy >= 50 {
			copy -= 50
			res += "L"
		} else if copy >= 40 {
			copy -= 40
			res += "XL"
		} else if copy >= 10 {
			copy -= 10
			res += "X"
		} else if copy >= 9 {
			copy -= 9
			res += "IX"
		} else if copy >= 5 {
			copy -= 5
			res += "V"
		} else if copy >= 4 {
			copy -= 4
			res += "IV"
		} else if copy >= 1 {
			copy -= 1
			res += "I"
		}
	}
	return res
}
```
