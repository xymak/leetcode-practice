package minimum_path_sum

func minPathSum(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])

	for i := 1; i < h; i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}

	for j := 1; j < w; j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			grid[i][j] = min(grid[i][j]+grid[i-1][j], grid[i][j]+grid[i][j-1])
		}
	}

	return grid[h-1][w-1]
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
