package coin_change

import "testing"

type TestCase struct {
	Coins    []int
	Amount   int
	Expected int
}

func TestCoinChange(t *testing.T) {
	aTestCase := []TestCase{
		{
			[]int{1, 2, 5}, 11, 3,
		},
		{
			[]int{2}, 3, -1,
		},
	}

	for _, v := range aTestCase {
		actual := coinChange(v.Coins, v.Amount)
		if v.Expected != actual {
			t.Errorf("actual %v not match expected %v", actual, v.Expected)
		}
	}
}
