package longest_palindromic_subsequence

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i:=0;i<n;i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	/*
	     0 1 2 3 4 5
	   0
	   1
	   2
	   3
	   4         1 1
	   5           1
	*/

	for i:=n-2;i>=0;i-- {
		for j:=i+1;j<=n-1;j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1],dp[i+1][j])
			}
		}
	}

	return dp[0][n-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
