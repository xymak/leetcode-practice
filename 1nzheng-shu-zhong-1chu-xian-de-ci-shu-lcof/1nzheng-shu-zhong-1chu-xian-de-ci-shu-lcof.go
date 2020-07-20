package _nzheng_shu_zhong_1chu_xian_de_ci_shu_lcof

func countDigitOne(n int) int {

	index := 1
	count := 0

	high := n
	low := 0
	current := 0

	for high > 0 {
		high /= 10
		current = (n / index) % 10
		low = n - (n / index) * index

		if current == 0 {
			count += high * index
		} else if current == 1 {
			count += high * index + low + 1
		} else if current > 1 {
			count += (high+1) * index
		}

		index *= 10
	}

	return count

}
