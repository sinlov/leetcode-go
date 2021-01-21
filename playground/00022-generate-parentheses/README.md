### 解题思路

- 不需要判断括号是否匹配的问题
- 排除特殊情况
- 直接 DFS 遍历生成

### 代码

```go
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}

	var res = []string{}
	dfsGenerateParenthesis(&res, n, n, "")
	return res
}

func dfsGenerateParenthesis(res *[]string, lIndex, rIndex int, word string) {
	if lIndex == 0 && rIndex == 0 { // dfs outter
		*res = append(*res, word)
		return
	}

	if lIndex > 0 { // left search
		dfsGenerateParenthesis(res, lIndex-1, rIndex, word+"(")
	}
	if rIndex > 0 && lIndex < rIndex { // right search
		dfsGenerateParenthesis(res, lIndex, rIndex-1, word+")")
	}
}

```
