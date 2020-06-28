package validate_binary_search_tree

import (
	"strconv"
	"strings"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	input := `[2,1,3]`

	expected := true
	splitInput := strings.Split(input, "\n")

	treeInputString := splitInput[0]

	treeInputString = treeInputString[1 : len(treeInputString)-1]
	treeInputArray := strings.Split(treeInputString, ",")
	treeInputArray = append([]string{"0"}, treeInputArray...)

	root := FormatTree(treeInputArray, 1)

	actual := isValidBST(root)

	if expected != actual {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
}

func FormatTree(a []string, i int) *TreeNode {
	if i >= len(a) {
		return nil
	}

	var val int
	if a[i] == "null" {
		return nil
	} else {
		val, _ = strconv.Atoi(a[i])
		root := &TreeNode{
			Val:   val,
			Left:  FormatTree(a, 2*i),
			Right: FormatTree(a, 2*i+1),
		}

		return root
	}
}
