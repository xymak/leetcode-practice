package evaluate_reverse_polish_notation

import (
	"strings"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	input := `["2","1","+","3","*"]`

	expected := 9
	splitInput := strings.Split(input, "\n")

	treeInputString := splitInput[0]

	treeInputString = treeInputString[1 : len(treeInputString)-1]
	treeInputArray := strings.Split(treeInputString, ",")
	for k, _ := range treeInputArray {
		treeInputArray[k] = strings.Replace(treeInputArray[k], `"`, "", 2)
	}

	actual := evalRPN(treeInputArray)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}
