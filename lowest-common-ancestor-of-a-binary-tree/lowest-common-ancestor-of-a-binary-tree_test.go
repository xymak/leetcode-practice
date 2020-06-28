package lowest_common_ancestor_of_a_binary_tree

import (
	"strconv"
	"strings"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	input := `[3,5,1,6,2,0,8,null,null,7,4]
5
1`

	expected := 3
	splitInput := strings.Split(input, "\n")

	treeInputString := splitInput[0]
	pInputString := splitInput[1]
	qInputString := splitInput[2]

	treeInputString = treeInputString[1 : len(treeInputString)-1]
	//fmt.Println(treeInputString)
	treeInputArray := strings.Split(treeInputString, ",")
	treeInputArray = append([]string{"0"}, treeInputArray...)

	root := FormatTree(treeInputArray, 1)
	pInt, _ := strconv.Atoi(pInputString)
	qInt, _ := strconv.Atoi(qInputString)
	p := &TreeNode{
		Val:   pInt,
	}
	q := &TreeNode{
		Val:   qInt,
	}

	l := lowestCommonAncestor(root, p, q)

	actual := -9999
	if l != nil {
		actual = l.Val
	}

	if expected != actual {
		t.Errorf("actual %d not match expected %d", actual, expected)
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
			Val: val,
			Left: FormatTree(a, 2*i),
			Right: FormatTree(a, 2*i+1),
		}

		return root
	}
}
