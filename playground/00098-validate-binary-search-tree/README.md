### 解题思路

按照 `左中右` 的顺序输出到数组中

- 如果是 BST，则数组中的数字是从小到大有序的
- 如果出现逆序就不是 BST

### 代码

```go
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

```
