package jump_game_ii

import (
	"strconv"
	"strings"
	"testing"
)

func TestJump(t *testing.T) {
	input := `[2,3,1,1,4]`

	expected := 2
	splitInput := strings.Split(input, "\n")

	inputString := splitInput[0]

	inputString = inputString[1 : len(inputString)-1]
	inputArray := strings.Split(inputString, ",")
	inputArrayInt := make([]int, len(inputArray), len(inputArray))
	for k, v := range inputArray {
		inputArrayInt[k], _ = strconv.Atoi(v)
	}

	actual := jump(inputArrayInt)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}