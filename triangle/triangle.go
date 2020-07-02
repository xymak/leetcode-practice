package triangle

func minimumTotal(triangle [][]int) int {
	var preRow []int
	var thisRow []int

	preRow = triangle[0]

	lineNum := len(triangle)

	preRow = make([]int, len(triangle[lineNum-1]))
	thisRow = make([]int, len(triangle[lineNum-1]))

	preRow[0] = triangle[0][0]

	for i := 1; i < lineNum; i++ {
		for j, v := range triangle[i] {
			if j == 0 {
				thisRow[j] = preRow[0] + v
			} else if j == len(triangle[i])-1 {
				thisRow[j] = preRow[len(triangle[i])-1-1] + v
			} else {
				thisRow[j] = min(preRow[j-1]+v, preRow[j]+v)
			}
		}
		copy(preRow, thisRow)
	}

	return min(preRow...)
}

func min(a ...int) int {
	min := a[0]
	l := len(a)

	for i := 1; i < l; i++ {
		if a[i] < min {
			min = a[i]
		}
	}

	return min

}
