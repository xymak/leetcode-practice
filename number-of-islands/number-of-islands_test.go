package number_of_islands

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	input := `[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]`
	expected := 1

	var inputOfArrayString [][]string
	json.Unmarshal([]byte(input), &inputOfArrayString)
	inputOfArrayByte := make([][]byte, len(inputOfArrayString))
	for krow, row := range inputOfArrayString {
		inputOfArrayByte[krow] = []byte(strings.Join(row, ""))
	}

	actual := numIslands(inputOfArrayByte)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
