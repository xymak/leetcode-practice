package backpack

import "testing"

type TestCase struct {
	A        []int
	M        int
	Expected int
}

func TestBackPack(t *testing.T) {
	aTestCase := []TestCase{
		{
			[]int{3, 4, 8, 5}, 10, 9,
		},
		{
			[]int{2, 3, 5, 7}, 12, 12,
		},
	}

	for _, v := range aTestCase {
		actual := backPack(v.M, v.A)
		if v.Expected != actual {
			t.Errorf("actual %v not match expected %v", actual, v.Expected)
		}
	}
}
