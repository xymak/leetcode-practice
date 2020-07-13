package jian_sheng_zi_lcof

func cuttingRope(n int) int {
	dp := make([]int, n+1)

	if n <= 2 {
		return 1
	}

	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 0; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max((i-j)*j, j*dp[i-j]))
		}

	}

	return dp[n]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
