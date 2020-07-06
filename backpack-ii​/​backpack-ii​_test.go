package _backpack_ii_

import (
	"testing"
)

type TestCase struct {
	A        []int
	M        int
	V        []int
	Expected int
}

func TestBackPackII(t *testing.T) {
	aTestCase := []TestCase{
		{
			[]int{3, 4, 8, 5}, 10, []int{1, 5, 2, 4}, 9,
		},
		{
			[]int{2, 3, 8}, 10, []int{2, 5, 8}, 10,
		},
	}

	for _, v := range aTestCase {
		actual := backPackII(v.M, v.A, v.V)
		if v.Expected != actual {
			t.Errorf("actual %v not match expected %v", actual, v.Expected)
		}
	}
}
