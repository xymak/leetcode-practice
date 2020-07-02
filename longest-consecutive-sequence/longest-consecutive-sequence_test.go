package longest_consecutive_sequence

import (
	"strconv"
	"strings"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	input := `[100,4,200,1,3,2]`

	expected := 4
	splitInput := strings.Split(input, "\n")

	treeInputString := splitInput[0]

	inputString := treeInputString[1 : len(treeInputString)-1]
	inputArrayString := strings.Split(inputString, ",")
	inputArrayInt := make([]int, len(inputArrayString), len(inputArrayString))
	for k, v := range inputArrayString {
		inputArrayInt[k], _ = strconv.Atoi(v)
	}

	actual := longestConsecutive(inputArrayInt)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
