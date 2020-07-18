package er_cha_shu_zhong_he_wei_mou_yi_zhi_de_lu_jing_lcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var r [][]int

	if root == nil {
		return r
	}

	dfs(root, &r, []int{}, sum)

	return r
}

func dfs(root *TreeNode, r *[][]int, path []int, sum int) {

	path = append(path, root.Val)

	if root.Left != nil {
		newPath := make([]int, len(path))
		copy(newPath, path)
		dfs(root.Left, r, newPath, sum)
	}

	if root.Right != nil {
		newPath := make([]int, len(path))
		copy(newPath, path)
		dfs(root.Right, r, newPath, sum)
	}

	if root.Left == nil && root.Right == nil {
		s := sumPath(path)

		if s == sum {
			*r = append(*r, path)
		}
	}
}

func sumPath(path []int) int {
	s := 0
	for i := 0; i < len(path); i++ {
		s += path[i]
	}

	return s
}
