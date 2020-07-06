package edit_distance

import "testing"

type TestCase struct {
	Text1    string
	Text2    string
	Expected int
}

func TestMinDistance(t *testing.T) {
	aTestCase := []TestCase{
		{
			"horse", "ros", 3,
		},
		{
			"intention", "execution", 5,
		},
	}

	for _, v := range aTestCase {
		actual := minDistance(v.Text1, v.Text2)
		if v.Expected != actual {
			t.Errorf("actual %v not match expected %v", actual, v.Expected)
		}
	}
}