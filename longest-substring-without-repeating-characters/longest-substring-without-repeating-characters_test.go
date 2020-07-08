package longest_substring_without_repeating_characters

import "testing"

type TestCase struct {
	Input    string
	Expected int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	input := []TestCase{
		{
			"abcabcbb", 3,
		},
		{
			"bbbbb", 1,
		},
		{
			"pwwkew", 3,
		},
	}

	for _, v := range input {
		actual := lengthOfLongestSubstring(v.Input)
		if v.Expected != actual {
			t.Errorf("actual %v not match expected %v", actual, v.Expected)
		}
	}
}
