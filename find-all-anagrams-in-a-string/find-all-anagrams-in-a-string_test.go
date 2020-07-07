package find_all_anagrams_in_a_string

import "testing"

type TestCase struct {
	S        string
	P        string
	Expected []int
}

func TestFindAnagrams(t *testing.T) {
	input := []TestCase{
		{
			"cbaebabacd",
			"abc",
			[]int{0,6},
		},
		{
			"abab",
			"ab",
			[]int{0,1,2},
		},
	}

	for _, v := range input {
		actual := findAnagrams(v.S, v.P)
		if len(actual) != len(v.Expected) {
			t.Errorf("not match, actual %v, expected %v", actual, v.Expected)
		}
		for k, _ := range actual {
			if actual[k] != v.Expected[k] {
				t.Errorf("not match, actual %v, expected %v", actual, v.Expected)
			}
		}
	}
}
