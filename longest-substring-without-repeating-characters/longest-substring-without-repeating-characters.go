package longest_substring_without_repeating_characters

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

*/

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	left := 0
	right := 0
	maxLen := 0

	win := make(map[byte]int)

	for right < len(s) {
		c := s[right]
		right++

		win[c]++

		for win[c] > 1 {
			d := s[left]
			win[d]--
			left++
		}

		maxLen = max(right-left, maxLen)

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
