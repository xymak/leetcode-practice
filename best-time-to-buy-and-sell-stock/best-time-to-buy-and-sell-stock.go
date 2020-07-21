package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	min := prices[0]
	max := 0

	for i := 1; i < n; i++ {
		if max < prices[i]-min {
			max = prices[i] - min
		}

		if prices[i] < min {
			min = prices[i]
		}
	}

	return max
}
