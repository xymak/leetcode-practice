package largest_rectangle_in_histogram

import (
	"strconv"
	"strings"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	input := `[2,1,5,6,2,3]`

	expected := 10
	splitInput := strings.Split(input, "\n")

	treeInputString := splitInput[0]

	inputString := treeInputString[1 : len(treeInputString)-1]
	inputArrayString := strings.Split(inputString, ",")
	inputArrayInt := make([]int, len(inputArrayString), len(inputArrayString))
	for k, v := range inputArrayString {
		inputArrayInt[k], _ = strconv.Atoi(v)
	}

	actual := largestRectangleArea(inputArrayInt)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
