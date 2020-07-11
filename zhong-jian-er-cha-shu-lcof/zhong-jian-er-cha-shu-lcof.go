package zhong_jian_er_cha_shu_lcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	i := 0
	for preorder[0] != inorder[i] {
		i++
	}

	thisNode := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:1+i], inorder[:i]),
		Right: buildTree(preorder[1+i:], inorder[i+1:]),
	}

	return thisNode
}
