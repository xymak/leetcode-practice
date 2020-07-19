package lian_xu_zi_shu_zu_de_zui_da_he_lcof

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
