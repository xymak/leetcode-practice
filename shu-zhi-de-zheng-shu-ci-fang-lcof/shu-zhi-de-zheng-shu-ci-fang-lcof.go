package shu_zhi_de_zheng_shu_ci_fang_lcof

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	nForNext := n / 2
	nForThis := n % 2
	this := 1.0
	if nForThis > 0 {
		this = x
	} else if nForThis < 0 {
		this = 1/x
	}

	next := 1.0
	if nForNext != 0 {
		next = myPow(x, nForNext)
	}

	return this * next * next
}
