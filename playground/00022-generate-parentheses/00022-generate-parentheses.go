package _00022_generate_parentheses

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
