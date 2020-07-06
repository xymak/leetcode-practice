package longest_increasing_subsequence

import (
	"strconv"
	"strings"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	//input := `[10,9,2,5,3,7,101,18]`
	input := `[10,9,2,5,3,4]`

	expected := 3
	splitInput := strings.Split(input, "\n")

	inputString := splitInput[0]

	inputString = inputString[1 : len(inputString)-1]
	inputArray := strings.Split(inputString, ",")
	inputArrayInt := make([]int, len(inputArray), len(inputArray))
	for k, v := range inputArray {
		inputArrayInt[k], _ = strconv.Atoi(v)
	}

	actual := lengthOfLIS(inputArrayInt)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
