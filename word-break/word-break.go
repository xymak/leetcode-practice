package word_break

/*
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

    拆分时可以重复使用字典中的单词。
    你可以假设字典中没有重复的单词。

示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。

示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。

示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

这题跟​palindrome-partitioning-ii​有点像
*/

func wordBreak(s string, wordDict []string) bool {
	mWordDict := make(map[string]bool)

	maxLengthOfWord := 0
	for _, word := range wordDict {
		mWordDict[word] = true
		maxLengthOfWord = max(maxLengthOfWord, len(word))
	}

	/*
		对于一个字符串[0,i]，j为其中一个分隔点
		状态转移公式为：f[i] = f[j] + InDict(j+1,i)
	*/

	l := len(s)
	dp := make(map[int]bool, l+1)

	//maxS := 0

	dp[0] = true
	for i, _ := range s {
		//fmt.Println("i:", i, s[0:i+1], mWordDict[s[0:i+1]])
		//if _, ok := mWordDict[s[0:i+1]]; ok {
		//	dp[i+1] = true
		//	continue
		//}

		l := 0
		if i-maxLengthOfWord > 0 {
			l = i - maxLengthOfWord
		}
		for j := l; j <= i; j++ {
			_, isInDict := mWordDict[s[j:i+1]]
			dp[i+1] = dp[i+1] || (dp[j] && isInDict)
			//fmt.Println("j", j, i, s[j:i+1], isInDict, dp[j], dp[j+1])
		}
	}

	return dp[l]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
