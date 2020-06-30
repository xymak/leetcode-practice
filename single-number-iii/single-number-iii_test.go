package single_number_iii

import (
	"strconv"
	"strings"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	input := `[1,2,1,3,2,5]`

	expected := []int{3, 5}
	splitInput := strings.Split(input, "\n")

	inputString := splitInput[0]

	inputString = inputString[1 : len(inputString)-1]
	inputArray := strings.Split(inputString, ",")
	inputArrayInt := make([]int, len(inputArray), len(inputArray))
	for k, v := range inputArray {
		inputArrayInt[k], _ = strconv.Atoi(v)
	}

	actual := singleNumber(inputArrayInt)

	if !((expected[0] == actual[0] && expected[1] == actual[1]) || (expected[0] == actual[1] && expected[1] == actual[0])) {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
