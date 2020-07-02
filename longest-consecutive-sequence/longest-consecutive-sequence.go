package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	isExist := make(map[int]bool)

	longestConsecutive := 0

	for _, v := range nums {
		isExist[v] = true
	}

	for _, v := range nums {

		if _, ok := isExist[v-1]; !ok {
			tmpConsecutive := 1
			for {
				_, ok := isExist[v+tmpConsecutive]

				if !ok {
					longestConsecutive = max(longestConsecutive, tmpConsecutive)
					break
				} else {
					tmpConsecutive++
				}
			}
		}

	}

	return longestConsecutive
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
