package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftDepth := 0
	isB := true

	if root.Left != nil {
		leftDepth, isB = isBalancedAndDepth(root.Left)
		if !isB {
			return false
		}
	}

	rightDepth := 0
	if root.Right != nil {
		rightDepth, isB = isBalancedAndDepth(root.Right)
		if !isB {
			return false
		}
	}

	if (leftDepth-rightDepth) <= 1 && (leftDepth-rightDepth) >= -1 {
		return true
	} else {
		return false
	}
}

func isBalancedAndDepth(root *TreeNode) (d int, isBalanced bool) {

	if root == nil {
		return 0, true
	}

	leftDepth := 0

	if root.Left != nil {
		leftDepth, isBalanced = isBalancedAndDepth(root.Left)
		if !isBalanced {
			return 0, isBalanced
		}
	}

	rightDepth := 0
	if root.Right != nil {
		rightDepth, isBalanced = isBalancedAndDepth(root.Right)
		if !isBalanced {
			return 0, isBalanced
		}
	}

	if leftDepth > rightDepth {
		d = leftDepth
	} else {
		d = rightDepth
	}

	if (leftDepth-rightDepth) <= 1 && (leftDepth-rightDepth) >= -1 {
		isBalanced = true
	} else {
		isBalanced = false
	}

	d = d + 1
	return d, isBalanced
}
