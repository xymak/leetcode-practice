package xuan_zhuan_shu_zu_de_zui_xiao_shu_zi_lcof

func minArray(numbers []int) int {

	l := len(numbers)
	if l == 0 {
		return -1
	}

	start := 0
	end := l - 1

	for {
		for start < end && numbers[end] == numbers[end-1] {
			end--
		}
		for start < end && numbers[start] == numbers[start+1] {
			start++
		}

		if start+1 >= end {
			break
		}
		mid := start + (end-start)/2
		if numbers[mid] < numbers[end] {
			end = mid
		} else if numbers[mid] == numbers[end] {
			start++
			end++
		} else {
			start = mid
		}
	}

	if numbers[start] > numbers[end] {
		return numbers[end]
	}
	return numbers[start]

}
