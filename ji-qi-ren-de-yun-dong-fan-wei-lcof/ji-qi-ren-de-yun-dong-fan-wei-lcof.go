package ji_qi_ren_de_yun_dong_fan_wei_lcof

func movingCount(m int, n int, k int) int {
	var q [][]int
	visited := make([][]bool, m)
	for i:=0;i<m;i++ {
		visited[i] = make([]bool, n)
	}

	q = append(q, []int{0,0})
	visited[0][0] = true

	sum := 0
	for {
		lenOfQ := len(q)
		if lenOfQ == 0 {
			break
		}

		for i:=0;i<lenOfQ;i++ {
			x := q[i][0]
			y := q[i][1]

			// fmt.Println(x,y,k, isOk(x, y, k))
			if isOk(x, y, k) {
				sum++
				if x < m-1 && !visited[x+1][y] {
					q = append(q, []int{x+1, y})
					visited[x+1][y] = true
				}
				if y < n-1 && !visited[x][y+1] {
					q = append(q, []int{x, y+1})
					visited[x][y+1] = true
				}
			}
		}

		q = q[lenOfQ:]
	}

	return sum
}

func isOk(m, n, k int) bool {
	return (m / 100 + m % 100 / 10 + m % 10 + n / 100  + n % 100 / 10 + n % 10) <= k
}
