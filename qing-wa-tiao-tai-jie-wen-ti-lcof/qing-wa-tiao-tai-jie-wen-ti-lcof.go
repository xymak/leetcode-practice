package qing_wa_tiao_tai_jie_wen_ti_lcof

func numWays(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	if n == 0 {
		return 1
	}

	dp[1] = 1
	if n == 1 {
		return 1
	}

	for i:= 2;i<=n;i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 1000000007
	}

	return dp[n]
}
