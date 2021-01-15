### 解题思路

- 回溯法

步奏

- 先列出 数字键盘的编码二维数组

> tips: 仅包含数字 2-9 的字符串 按顺序放即可

- 按照回溯的写法，扫描整个字符串
- 按选优条件向前搜索，达到目标就返回，如果没有就后退
- 这道题没有特别的处理，按照模板写就行

### 代码

```go
var dict = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return res
	}
	letterBiz("", digits)
	return res
}

func letterBiz(result string, digits string) {
	if digits == "" { // exit
		res = append(res, result)
		return
	}
	k := digits[0:1] // recall biz
	digits = digits[1:]
	for i := 0; i < len(dict[k]); i++ {
		result += dict[k][i] // back
		letterBiz(result, digits)
		result = result[0 : len(result)-1]
	}
}

```
