package unique_paths_ii

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	o := make([][]int, len(obstacleGrid))
	h := len(obstacleGrid)
	w := len(obstacleGrid[0])

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if o[i] == nil {
				o[i] = make([]int, w)
			}
			o[i][j] = 1
		}
	}

	if obstacleGrid[0][0] == 1 {
		o[0][0] = 0
	} else {
		o[0][0] = 1
	}

	for i := 1; i < h; i++ {
		if obstacleGrid[i][0] == 1 || o[i-1][0] == 0 {
			o[i][0] = 0
		}
	}

	for j := 1; j < w; j++ {
		if obstacleGrid[0][j] == 1 || o[0][j-1] == 0 {
			o[0][j] = 0
		}
	}

	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if obstacleGrid[i][j] == 1 {
				o[i][j] = 0
			} else {
				o[i][j] = o[i-1][j] + o[i][j-1]
			}
		}
	}

	return o[h-1][w-1]
}
