package best_time_to_buy_and_sell_stock_with_cooldown

/*
	持有一支股票，对应的「累计最大收益」记为 f[i][0]
	不持有任何股票，并且处于冷冻期中，对应的「累计最大收益」记为 f[i][1]
	不持有任何股票，并且不处于冷冻期中，对应的「累计最大收益」记为 f[i][2]

	f[i][0]=max(f[i−1][0],f[i−1][2]−prices[i])
	f[i][1]=f[i−1][0]+prices[i]
	f[i][2]=max(f[i−1][1],f[i−1][2])
 */

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 3)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}

	return max(dp[n-1][1], dp[n-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
