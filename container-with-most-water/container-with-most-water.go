package container_with_most_water

func maxArea(height []int) int {
	m := 0
	l := 0
	r := len(height) - 1

	for {
		if l >= r {
			break
		}
		m = max(m, min(height[l], height[r])*(r-l))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return m
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
