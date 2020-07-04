package palindrome_partitioning_ii

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回符合要求的最少分割次数。

示例:

输入: "aab"
输出: 1
解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。

*/

func minCut(s string) int {
	l := len(s)

	dp := make([]int, l)

	for i:=0;i<l;i++ {
		dp[i] = i
	}

	for j:=1;j<l;j++ {

		if isPalindrome(s, 0, j) {
			dp[j] = 0
			continue
		}

		for k:=0;k<j;k++ {
			if isPalindrome(s, k+1, j) {
				dp[j] = min(dp[k]+1, dp[j])
			}
		}
	}

	return dp[l-1]
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func min(a int, b int ) int {
	if a < b {
		return a
	} else {
		return b
	}
}