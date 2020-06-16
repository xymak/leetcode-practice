package unique_paths

import "testing"

func TestUniquePaths(t *testing.T) {
	expected := 3
	actual := uniquePaths(3, 2)
	if expected != actual {
		t.Errorf("expeted %d, but actual %d", expected, actual)
	}
}

