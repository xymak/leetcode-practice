package binary_tree_level_order_traversal_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	n := depthOfTree(root)

	o := make([][]int, n)

	dfs(root, 0, o, n)

	return o
}

func dfs(root *TreeNode, depth int, o [][]int, n int) {
	if root == nil {
		return
	}

	o[n-1-depth] = append(o[n-1-depth], root.Val)

	dfs(root.Left, depth+1, o, n)
	dfs(root.Right, depth+1, o, n)
}

func depthOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depthOfTree(root.Left), depthOfTree(root.Right)) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}