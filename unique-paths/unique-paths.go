package unique_paths

func uniquePaths(m int, n int) int {
	// d := make([][]int, m)
	// for k, _ := range d {
	//     d[k] = make([]int, n)
	// }
	// for p:=0;p<m;p++{
	//     d[p][0] = 1
	// }
	// for q:=0;q<n;q++{
	//     d[0][q] = 1
	// }
	// for i:=1; i<m; i++ {
	//     for j:=1;j<n;j++ {
	//         d[i][j] = d[i][j-1] + d[i-1][j]
	//     }
	// }

	// return d[m-1][n-1]

	d := make([]int, m)
	for p := 0; p < m; p++ {
		d[p] = 1
	}

	for i:= 1; i < n; i++ {
		for j:=1; j<m; j++ {
			d[j] = d[j-1] + d[j]
		}
	}

	return d[m-1]
}
