package _00098_validate_binary_search_tree

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

func isValidBST(root *TreeNode) bool {
	arr := []int{}
	treeOrder(root, &arr)
	for i := 1; i < len(arr); i++ { // check is reverse
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func treeOrder(root *TreeNode, arr *[]int) { // left mid rgiht
	if root == nil {
		return
	}
	treeOrder(root.Left, arr)
	*arr = append(*arr, root.Val)
	treeOrder(root.Right, arr)
}
