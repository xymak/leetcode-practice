package longest_common_subsequence

import "testing"

type TestCase struct {
	Text1    string
	Text2    string
	Expected int
}

func TestLongestCommonSubsequence(t *testing.T) {
	aTestCase := []TestCase{
		{
			"abcde", "ace", 3,
		},
		{
			"abc", "abc", 3,
		},
		{
			"abc", "def", 0,
		},
	}

	for _, v := range aTestCase {
		actual := longestCommonSubsequence(v.Text1, v.Text2)
		if v.Expected != actual {
			t.Errorf("actual %v not match expected %v", actual, v.Expected)
		}
	}
}
