package single_number_ii

import (
	"strconv"
	"strings"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	input := `[-2,-2,1,1,-3,1,-3,-3,-4,-2]`

	expected := -4
	splitInput := strings.Split(input, "\n")

	inputString := splitInput[0]

	inputString = inputString[1 : len(inputString)-1]
	inputArray := strings.Split(inputString, ",")
	inputArrayInt := make([]int, len(inputArray), len(inputArray))
	for k, v := range inputArray {
		inputArrayInt[k], _ = strconv.Atoi(v)
	}



	actual := singleNumber(inputArrayInt)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
