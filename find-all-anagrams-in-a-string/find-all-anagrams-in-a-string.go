package find_all_anagrams_in_a_string

/*
给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：

    字母异位词指字母相同，但排列不同的字符串。
    不考虑答案输出的顺序。

示例 1:

输入:
s: "cbaebabacd" p: "abc"

输出:
[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。

 示例 2:

输入:
s: "abab" p: "ab"

输出:
[0, 1, 2]

解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
*/

func findAnagrams(s string, p string) []int {
	var r []int

	win := make(map[byte]int)
	need := make(map[byte]int)

	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	match := 0
	left := 0
	right := 0

	for right < len(s) {
		c := s[right]
		right++

		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}

		for right-left >= len(p) {
			if match == len(need) {
				r = append(r, left)
			}

			d := s[left]
			if need[d] != 0 {
				if win[d] == need[d] {
					match--
				}
				win[d]--
			}

			left++
		}
	}

	return r
}
