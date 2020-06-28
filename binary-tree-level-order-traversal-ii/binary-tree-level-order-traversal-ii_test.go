package binary_tree_level_order_traversal_ii

import (
	"strconv"
	"strings"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
	input := `[3,9,20,null,null,15,7]`

	var expected [][]int
	expected = [][]int{
		{15, 7},
		{9, 20},
		{3},
	}

	splitInput := strings.Split(input, "\n")

	treeInputString := splitInput[0]

	treeInputString = treeInputString[1 : len(treeInputString)-1]
	treeInputArray := strings.Split(treeInputString, ",")
	treeInputArray = append([]string{"0"}, treeInputArray...)

	root := FormatTree(treeInputArray, 1)

	actual := levelOrderBottom(root)

	if len(expected) != len(actual) {
		t.Errorf("actual %v not match expected %v", actual, expected)
	}
	for k, _ := range expected {
		if len(expected[k]) != len(actual[k]) {
			t.Errorf("actual %v not match expected %v", actual, expected)
		}
		for ck, _ := range expected[k] {
			if expected[k][ck] != actual[k][ck] {
				t.Errorf("actual %v not match expected %v", actual, expected)
			}
		}
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
