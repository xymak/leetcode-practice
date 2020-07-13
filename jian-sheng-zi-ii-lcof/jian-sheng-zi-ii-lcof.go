package jian_sheng_zi_ii_lcof

func cuttingRope(n int) int {
	dp := make([]int, n+1)

	if n <= 2 {
		return 1
	}

	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	dp[3] = 2
	dp[4] = 4
	dp[5] = 6
	dp[6] = 9


	for i := 7; i <= n; i++ {
		dp[i] = (dp[i-3] * 3) % 1000000007
	}

	return dp[n]
}