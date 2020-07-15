package da_yin_cong_1dao_zui_da_de_nwei_shu_lcof

func printNumbers(n int) []int {
	m := 1
	for i:=0;i<n;i++ {
		m *= 10
	}

	r := make([]int, m-1)
	for j:=0;j<m-1;j++ {
		r[j] = j+1
	}

	return r
}
