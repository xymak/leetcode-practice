package best_time_to_buy_and_sell_stock_iv

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n <= 1 || k == 0 {
		return 0
	}

	if k >= n/2 {
		return maxProfitInf(prices)
	}

	dp := make([][2]int, k+1)

	dp[0][0] = 0
	dp[1][1] = -prices[0]
	dp[0][1] = -math.MaxInt32
	dp[1][0] = -math.MaxInt32

	for i := 2; i <= k; i++ {
		dp[i][0] = -math.MaxInt32
		dp[i][1] = -math.MaxInt32
	}

	for i := 1; i < n; i++ {
		for j := k; j >= 0; j-- {
			dp[j][0] = max(dp[j][0], dp[j][1]+prices[i])
			if j > 0 {
				dp[j][1] = max(dp[j][1], dp[j-1][0]-prices[i])
			}

		}
	}

	var a []int
	for j := 0; j <= k; j++ {
		a = append(a, dp[j][0])
	}
	return max(a...)
}

func max(a ...int) int {
	m := -math.MaxInt32
	for _, v := range a {
		if v > m {
			m = v
		}
	}
	return m
}

func maxProfitInf(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[n-1][1]
}
