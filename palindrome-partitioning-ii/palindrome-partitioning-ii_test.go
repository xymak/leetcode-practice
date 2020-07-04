package palindrome_partitioning_ii

import (
	"testing"
)

func TestMinCut(t *testing.T) {
	input := `aab`
	expected := 1

	actual := minCut(input)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
