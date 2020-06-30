package _1_matrix

import "testing"

func TestIsValidBST(t *testing.T) {
	input := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}

	expected := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}

	actual := updateMatrix(input)

	for x, row := range expected {
		for y, _ := range row {
			if expected[x][y] != actual[x][y] {
				t.Errorf("actual %v not match expected %v", actual, expected)
			}
		}
	}

}
