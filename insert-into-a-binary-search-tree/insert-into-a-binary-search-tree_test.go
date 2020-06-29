package insert_into_a_binary_search_tree

import (
	"strconv"
	"strings"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	input := `[4,2,7,1,3]
5`

	expected := "[4,2,7,1,3,5]"
	splitInput := strings.Split(input, "\n")

	treeInputString := splitInput[0]
	valInputString := splitInput[1]

	treeInputString = treeInputString[1 : len(treeInputString)-1]
	treeInputArray := strings.Split(treeInputString, ",")
	treeInputArray = append([]string{"0"}, treeInputArray...)

	root := FormatTree(treeInputArray, 1)
	valInt, _ := strconv.Atoi(valInputString)

	actual := insertIntoBST(root, valInt)

	if FormatTreeToString(actual) != expected {
		t.Errorf("actual %v not match expected %v", FormatTreeToString(actual), expected)
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

func FormatTreeToString(root *TreeNode) string {
	var q []string

	bfs(root, &q)
	return "[" + strings.Join(q, ",") + "]"
}

func bfs(root *TreeNode, q *[]string) {

	var p []*TreeNode

	if root == nil {
		return
	}

	p = append(p, root)

	for {
		if len(p) == 0 {
			break
		}

		lenOfP := len(p)
		for i := 0; i < lenOfP; i++ {
			v := p[i]
			*q = append(*q, strconv.Itoa(v.Val))
			if v.Left != nil {
				p = append(p, v.Left)
			}
			if v.Right != nil {
				p = append(p, v.Right)
			}
		}
		p = p[lenOfP:]
	}

}
