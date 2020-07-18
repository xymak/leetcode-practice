package shu_de_zi_jie_gou_lcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return false
	}

	var q []*TreeNode
	q = append(q, A)

	for {

		lenOfQ := len(q)
		if lenOfQ == 0 {
			return false
		}

		for i := 0; i < lenOfQ; i++ {
			if q[i].Val == B.Val {
				if is(q[i].Left, B.Left) && is(q[i].Right, B.Right) {
					return true
				}
			}

			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}

		}

		q = q[lenOfQ:]

	}

	return false
}

func is(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	} else if A == nil {
		return false
	}

	if A.Val == B.Val {
		if is(A.Left, B.Left) && is(A.Right, B.Right) {
			return true
		}
	}

	return false

}
