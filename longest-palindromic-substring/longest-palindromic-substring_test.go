package longest_palindromic_substring

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	input := `babad`

	expected := "bab"

	actual := longestPalindrome(input)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
