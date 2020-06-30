package _1_matrix

type QueueElement struct {
	X int
	Y int
}

func updateMatrix(matrix [][]int) [][]int {
	var q []QueueElement

	for x, row := range matrix {

		for y, _ := range row {
			if matrix[x][y] == 0 {
				q = append(q, QueueElement{X: x, Y: y})
			} else {
				matrix[x][y] = -1
			}
		}
	}

	for {
		lenOfQ := len(q)
		if len(q) == 0 {
			break
		}

		for i := lenOfQ - 1; i >= 0; i-- {
			x := q[i].X
			y := q[i].Y

			if x > 0 && matrix[x-1][y] == -1 {
				matrix[x-1][y] = matrix[x][y] + 1
				q = append(q, QueueElement{x - 1, y})
			}

			if x < len(matrix)-1 && matrix[x+1][y] == -1 {
				matrix[x+1][y] = matrix[x][y] + 1
				q = append(q, QueueElement{x + 1, y})
			}

			if y < len(matrix[0])-1 && matrix[x][y+1] == -1 {
				matrix[x][y+1] = matrix[x][y] + 1
				q = append(q, QueueElement{x, y + 1})
			}

			if y > 0 && matrix[x][y-1] == -1 {
				matrix[x][y-1] = matrix[x][y] + 1
				q = append(q, QueueElement{x, y - 1})
			}

		}

		q = q[lenOfQ:]
	}

	return matrix

}
