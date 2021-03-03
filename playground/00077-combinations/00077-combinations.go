package _00077_combinations

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n { // special case
		return [][]int{}
	}

	cache, res := []int{}, [][]int{}
	generateCombinations(n, k, 1, cache, &res)
	return res
}

func generateCombinations(n, k, start int, c []int, res *[][]int) {
	if len(c) == k { // recursion out
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := start; i <= n-(k-len(c))+1; i++ { // n - (k - c.size()) + 1
		c = append(c, i)
		generateCombinations(n, k, i+1, c, res)
		c = c[:len(c)-1]
	}
	return
}
