package insert_into_a_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	now := root
	for {
		if now.Val > val {
			if now.Left == nil {
				now.Left = &TreeNode{
					Val: val,
				}
				break
			} else {
				now = now.Left
			}
		} else {
			if now.Right == nil {
				now.Right = &TreeNode{
					Val: val,
				}
				break
			} else {
				now = now.Right
			}

		}

	}

	return root
}
