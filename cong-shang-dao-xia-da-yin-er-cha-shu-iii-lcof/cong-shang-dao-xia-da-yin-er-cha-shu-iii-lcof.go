package cong_shang_dao_xia_da_yin_er_cha_shu_iii_lcof

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
	var round int
	for {
		lenOfQ := len(q)
		if lenOfQ == 0 {
			break
		}

		l := make([]int, lenOfQ)
		round++

		for i := 0; i < lenOfQ; i++ {
			if round%2 == 0 {
				l[lenOfQ-1-i] = q[i].Val
			} else {
				l[i] = q[i].Val
			}

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
