package palindrome_partitioning_ii

/*
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。

返回符合要求的最少分割次数。

示例:

输入: "aab"
输出: 1
解释: 进行一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。

其中一段为[0,j],

如果k为其中一个分隔点，[k+1, j]为回文
dp[j] = min(dp[k] + 1, dp[j])

*/

func minCut(s string) int {
	l := len(s)

	dp := make([]int, l)

	// 假设最长回文子串是1
	for i := 0; i < l; i++ {
		dp[i] = i
	}

	/*
		这里还可以用空间复杂度为n^2的最长回文子串的方法预先算好所有子串是否是回文。
		因为本身每个去判断的话，是没有做动态规划的。
		公式为dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
	*/
	dpIsPalindrome := make([][]bool, l)
	for i := l-1; i >= 0; i-- {
		dpIsPalindrome[i] = make([]bool, l)
		for j := l-1; j >= 0 && j >= i; j-- {
			dpIsPalindrome[i][j] = (i==j) || (j-i==1 && s[i] == s[j]) || (dpIsPalindrome[i+1][j-1] && s[i] == s[j])
		}
	}

	for j := 1; j < l; j++ {

		// 如果[0,j]是回文，那么从起点到这一段直接为0
		//if isPalindrome(s, 0, j) {
		if dpIsPalindrome[0][j] {
			dp[j] = 0
			continue
		}

		// 如果不是回文，则开始分隔
		for k := 0; k < j; k++ {
			// 分析后半段是回文的话，整一段的分个数等于前半段的分割数+1，并且与历史比较
			//if isPalindrome(s, k+1, j) {
			if dpIsPalindrome[k+1][j] {
				dp[j] = min(dp[k]+1, dp[j])
			}
		}
	}

	return dp[l-1]
}

//func isPalindrome(s string, i, j int) bool {
//	for i < j {
//		if s[i] != s[j] {
//			return false
//		}
//		i++
//		j--
//	}
//	return true
//}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
