package fibonacci_number

import "testing"

func TestFib(t *testing.T) {

}

type TestCase struct {
	Input    int
	Expected int
}

func TestBackPack(t *testing.T) {
	aTestCase := []TestCase{
		{
			0, 0,
		},
		{
			2, 1,
		},
		{
			3, 2,
		},
		{
			4, 3,
		},
	}

	for _, v := range aTestCase {
		actual := fib(v.Input)
		if v.Expected != actual {
			t.Errorf("actual %v not match expected %v", actual, v.Expected)
		}
	}
}
