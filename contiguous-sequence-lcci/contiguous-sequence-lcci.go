package contiguous_sequence_lcci

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]

	m := dp[0]
	for i := 1; i < n; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		m = max(dp[i], m)
	}

	return m
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
