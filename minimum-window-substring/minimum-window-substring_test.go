package minimum_window_substring

import "testing"

func TestMinWindow(t *testing.T) {
	input1 := "ADOBECODEBANC"
	input2 := "ABC"

	expected := "BANC"

	actual := minWindow(input1, input2)

	if expected != actual {
		t.Errorf("actual %s not match expected %s", actual, expected)
	}
}