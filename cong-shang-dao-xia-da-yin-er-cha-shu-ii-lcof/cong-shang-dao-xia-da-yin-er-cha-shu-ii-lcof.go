package cong_shang_dao_xia_da_yin_er_cha_shu_ii_lcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var r [][]int
	var q []*TreeNode

	if root == nil {
		return r
	}

	q = append(q, root)
	for {
		lenOfQ := len(q)
		if lenOfQ == 0 {
			break
		}

		var l []int
		for i := 0; i < lenOfQ; i++ {
			l = append(l, q[i].Val)

			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		r = append(r, l)

		q = q[lenOfQ:]
	}

	return r
}
