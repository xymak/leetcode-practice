package longest_palindromic_substring

/*
	这题碰见很多次，七牛、流利说都考了这题，下面写的思路是
	有一段子串i到j
	dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
	本身空间复杂度是是要n^2的
	因为公式是先要知道i+1行的结果，所以i遍历的时候倒序，j是先要知道j-1的，如果是采用同一行，那么会跟本行冲突，所以也要倒序
*/

func longestPalindrome(s string) string {
	l := len(s)
	maxS := ""
	dp := make([]bool,l)

	for i := l-1 ;i>=0;i-- {
		for j := l-1; j>=i; j-- {
			// 长度为1 || 长度为2 且 两边相等 || 去掉两边的子串为回文 且 两边相等
			dp[j] = (i==j) || (j-i==1 && s[i] == s[j]) || (dp[j-1] && s[i] == s[j])
			if dp[j] && len(maxS) <= j-i+1 {
				maxS = s[i:j+1]
			}
		}
	}

	return maxS
}
