package _00095_unique_binary_search_trees_ii

import "github.com/sinlov/leetcode-go/structures"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structures.TreeNode

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return parseBSTres(1, n)
}

func parseBSTres(start, end int) []*TreeNode {
	tree := []*TreeNode{}
	if start > end {
		tree = append(tree, nil)
		return tree
	}
	for i := start; i <= end; i++ {
		left := parseBSTres(start, i-1)
		right := parseBSTres(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				tree = append(tree, root)
			}
		}
	}
	return tree
}
