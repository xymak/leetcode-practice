package er_wei_shu_zu_zhong_de_cha_zhao_lcof

func findNumberIn2DArray(matrix [][]int, target int) bool {
	h := len(matrix)
	if h == 0 {
		return false
	}

	w := len(matrix[0])
	if w == 0 {
		return false
	}

	x:= w-1
	y:= 0

	for x >= 0 && y < h {
		if matrix[y][x] > target {
			x--
		} else if matrix[y][x] < target {
			y++
		} else {
			return true
		}
	}

	return false
}