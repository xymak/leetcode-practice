package li_wu_de_zui_da_jie_zhi_lcof

func maxValue(grid [][]int) int {
	m := len(grid[0])
	n := len(grid)

	for i := 1; i < m; i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}
	for j := 1; j < n; j++ {
		grid[j][0] = grid[j-1][0] + grid[j][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			grid[i][j] = grid[i][j] + max(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[n-1][m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
