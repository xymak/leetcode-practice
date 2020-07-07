package permutation_in_string

import "testing"

func TestCheckInclusion(t *testing.T) {
	input1 := "ad"
	input2 := "eidbaooo"

	expected := false

	actual := checkInclusion(input1, input2)
	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
