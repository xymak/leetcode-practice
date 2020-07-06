package longest_increasing_subsequence

/*
定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。

说明:

    可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
    你算法的时间复杂度应该为 O(n2) 。

进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?


*/
func lengthOfLIS(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	dp := make([]int, l)
	dp[0] = 1

	maxLen := 1
	for i := 1; i < l; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {

			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
				maxLen = max(dp[i], maxLen)
			}
		}

	}

	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
