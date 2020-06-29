package number_of_islands

func numIslands(grid [][]byte) int {
	var c int
	for krow, row := range grid {
		for kcolumn, _ := range row {
			if dfs(grid, krow, kcolumn) > 0 {
				c++
			}
		}
	}

	return c
}

func dfs(grid [][]byte, i int, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}

	if grid[i][j] == '1' {
		grid[i][j] = 0
		return dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1) + 1
	}

	return 0
}
