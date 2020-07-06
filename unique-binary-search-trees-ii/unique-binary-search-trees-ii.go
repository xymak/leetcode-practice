package unique_binary_search_trees_ii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func generate(start int, end int) []*TreeNode {
	var r []*TreeNode

	if start > end {
		r = append(r, nil)
		return r
	}
	for i := start; i <= end; i++ {
		aLeft := generate(start, i-1)
		aRight := generate(i+1, end)
		// fmt.Println(aLeft, aRight)

		for _, lv := range aLeft {
			for _, rv := range aRight {
				r = append(r, &TreeNode{
					Val:   i,
					Left:  lv,
					Right: rv,
				})
			}
		}

	}

	return r
}
