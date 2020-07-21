package best_time_to_buy_and_sell_stock_iii

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][3][2]int, n)

	dp[0][0][0] = 0
	dp[0][1][0] = -math.MaxInt32
	dp[0][2][0] = -math.MaxInt32
	dp[0][0][1] = -math.MaxInt32
	dp[0][1][1] = -prices[0]
	dp[0][2][1] = -math.MaxInt32

	for i := 1; i < n; i++ {
		dp[i][0][0] = dp[0][0][0]
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])

		dp[i][0][1] = 0
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
		dp[i][2][1] = max(dp[i-1][1][0]-prices[i], dp[i-1][2][1])

	}

	return max(dp[n-1][0][0], max(dp[n-1][2][0], dp[n-1][1][0]))
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
