package maximal_square

func maximalSquare(matrix [][]byte) int {
	max := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j] = int(matrix[i][j] - "0"[0])
			if dp[i][j] == 1 {
				max = 1
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}

	return max * max

}

func min(a ...int) int {
	m := 99999
	for _, v := range a {
		if m > v {
			m = v
		}
	}

	return m
}
