package minimum_window_substring

import (
	"math"
)

/*
给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字符的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"

说明：

    如果 S 中不存这样的子串，则返回空字符串 ""。
    如果 S 中存在这样的子串，我们保证它是唯一的答案。

*/

func minWindow(s string, t string) string {
	win := make(map[byte]int)
	need := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left := 0
	right := 0

	match := 0
	start := 0
	end := 0
	min := math.MaxInt64

	var c byte
	for right < len(s) {

		c = s[right]
		right++

		if need[c] != 0 {
			win[c]++

			if win[c] == need[c] {
				match++
			}
		}

		for match == len(need) {
			// 记录当前最小值
			if right-left < min {
				min = right-left
				start = left
				end = right
			}

			c = s[left]
			left++

			if need[c] != 0 {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}


	}

	if min == math.MaxInt64 {
		return ""
	} else {
		return s[start:end]
	}
}
