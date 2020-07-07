package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	形成中序遍历的数据，中序数据必须节点从小到大排列
 */

func isValidBST(root *TreeNode) bool {
	s := []int{}

	inorder(root, &s)

	l := len(s)
	if l <= 1 {
		return true
	}

	for i := 0; i < l-1; i++ {
		if s[i] >= s[i+1] {
			return false
		}
	}

	return true
}

func inorder(root *TreeNode, s *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		inorder(root.Left, s)
	}

	*s = append(*s, root.Val)

	if root.Right != nil {
		inorder(root.Right, s)
	}

}
