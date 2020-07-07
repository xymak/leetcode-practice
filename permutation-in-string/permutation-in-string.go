package permutation_in_string

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").



示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False



注意：

    输入的字符串只包含小写字母
    两个字符串的长度都在 [1, 10,000] 之间

子串的意思是连续的
 */

func checkInclusion(s1 string, s2 string) bool {
	win := make(map[byte]int)
	need := make(map[byte]int)

	for i:=0;i<len(s1);i++ {
		need[s1[i]]++
	}

	match := 0
	left := 0
	right := 0

	for right < len(s2) {
		c := s2[right]
		right++

		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}

		for right - left >= len(s1) {
			if match == len(need) {
				return true
			}

			d := s2[left]
			if need[d] != 0 {
				if win[d] == need[d] {
					match--
				}
				win[d]--
			}

			left++
		}
	}

	return false
}