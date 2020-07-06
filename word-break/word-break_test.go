package word_break

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	input := `applepenapple`
	intputWordDict := []string{"apple", "pen"}

	expected := true

	actual := wordBreak(input, intputWordDict)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
