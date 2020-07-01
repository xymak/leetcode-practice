package find_minimum_in_rotated_sorted_array_ii

import (
	"strconv"
	"strings"
	"testing"
)

func TestFindMin(t *testing.T) {
	input := `[2,2,2,0,1,2]`

	expected := 0
	splitInput := strings.Split(input, "\n")

	inputString := splitInput[0]

	inputString = inputString[1 : len(inputString)-1]
	inputArray := strings.Split(inputString, ",")
	inputArrayInt := make([]int, len(inputArray), len(inputArray))
	for k, v := range inputArray {
		inputArrayInt[k], _ = strconv.Atoi(v)
	}

	actual := findMin(inputArrayInt)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
