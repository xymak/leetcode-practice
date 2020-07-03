package climbing_stairs

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	input := 3

	expected := 3

	actual := climbStairs(input)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
