package er_jin_zhi_zhong_1de_ge_shu_lcof

func hammingWeight(num uint32) int {
	sum := 0
	for i:=uint32(0);i<32;i++ {
		if (1 & (num >> i)) > 0 {
			sum ++
		}
	}
	return sum
}
