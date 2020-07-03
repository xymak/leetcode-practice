package jump_game_ii

/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例:

输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

说明:

假设你总是可以到达数组的最后一个位置。
*/

func jump(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = 0

	for i := 0; i < l; i++ {
		for j := 1; j <= nums[i] && j+i < l; j++ {
			if dp[j+i] == 0 {
				dp[j+i] = dp[i]+1
			} else {
				dp[j+i] = min(dp[i+j], dp[i]+1)
			}

		}
	}

	return dp[l-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}