### 解题思路

Manacher算法可以在线性时间复杂度内求出一个字符串的最长回文字串，达到了理论上的下界

#### Manacher 算法

将长度为奇数的回文串和长度为偶数的回文串一起考虑的算法
具体做法是：

- 在原字符串的每个相邻两个字符中间插入一个分隔符
- 同时在首尾也要添加一个分隔符
- 分隔符的要求是不在原串中出现

习惯上，分隔符用 `#`代替

```go
func toManacherStr(s string) string {
    ans := make([]rune, 0)
    ans = append(ans, '#')
    for i := 0; i < len(s); i++ {
        ans = append(ans, rune(s[i]), '#')
    }
    return string(ans)
}
```

### 解题步奏

- 先处理 空串
- 设置标识 当前检查的字符下标 id
- 设置标识 `mPoint`: middle point, 它是指当前有效回文字符串的中点字符的下标. mPoint的初始值是0, 它会逐渐右移
- 设置标识 `rPoint`：right point 它是指当前有效回文字符串的右端字符的下标+1，也就是回文半径
  - 从左端开始检查的
  - 利用回文字符串的特性, 减少不必要的检查
  - rPoint的初始值是0, 它会逐渐右移
- path 为一个数组，放置检查值
- 转换字符串成为 Manacher 字符串
- 遍历 Manacher 游标为 i
- 判断 `i < rPoint` 确认当前检查的字符是否落在了当前有效回文字符串内
  - true 则说明之前的检查已经找到了当前有效回文字符串, 并且其半径覆盖到了当前检查的字符, 那么 `path[i] = 1`
  - false 当前还没有找到回文字符串, 或者当前有效回文字符串的半径没有覆盖到当前检查的字符, 那么 更新检查
     - 更新方法为 判断 `path[2*id-i]` 是否 大于 `rPoint-i`
       - true 则把 `rPoint - i` 的值付给 `path[i]` 
       - false 则把 `path[2*id-i]` 给 `path[i]`
- 判断 如果回文找到 目标之外，则直接 `path[i] += 1`
- 判断 `path[i] > path[mPoint]` 则需要将 mPoint 右动位置
- 判断 `i+path[i] > rPoint` 则需要增加 rPoint 半径 `i + path[i]` ，并检查的字符下标 id 移动 
- 定位到 Manacher 字符串为 转换的字符串截取在 `mPoint-path[mPoint]+1` 和 `mPoint+path[mPoint]` 之间，注意半径计算
- 最后去掉 Manacher 字符串的分隔符

### 代码

```go
import "strings"

var (
	separator     = "#"
	separatorChar = []rune(separator)[0]
)

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	id, mPoint, rPoint := 0, 0, 0
	manacherStr := toManacherStr(s)

	var path = make([]int, len(manacherStr))
	for i, _ := range manacherStr {
		if i < rPoint {
			lg := path[2*id-i]
			if rPoint-i < lg {
				lg = rPoint - i
			}
			path[i] = lg
		} else {
			path[i] = 1
		}

		for i-path[i] >= 0 && i+path[i] < len(manacherStr) &&
			manacherStr[i-path[i]] == manacherStr[i+path[i]] {
			path[i] += 1
		}

		if path[i] > path[mPoint] {
			mPoint = i
		}

		if i+path[i] > rPoint {
			rPoint = i + path[i]
			id = i
		}
	}

	return strings.Replace(manacherStr[(mPoint-path[mPoint]+1):(mPoint+path[mPoint])], separator, "", -1)
}

func toManacherStr(s string) string {
	ans := make([]rune, 0)
	ans = append(ans, separatorChar)
	for i := 0; i < len(s); i++ {
		ans = append(ans, rune(s[i]), separatorChar)
	}
	return string(ans)
}
```
