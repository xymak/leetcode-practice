package shu_zu_zhong_zhong_fu_de_shu_zi_lcof

func findRepeatNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
		if m[v] > 1 {
			return v
		}
	}

	return 0
}
