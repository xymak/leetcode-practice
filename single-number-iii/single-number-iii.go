package single_number_iii

func singleNumber(nums []int) []int {
	diff := 0

	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}

	result := []int{diff, diff}
	diff = diff&(diff-1) ^ diff

	for j := 0; j < len(nums); j++ {
		if diff&nums[j] > 0 {
			result[0] ^= nums[j]
		} else {
			result[1] ^= nums[j]
		}
	}

	return result
}
