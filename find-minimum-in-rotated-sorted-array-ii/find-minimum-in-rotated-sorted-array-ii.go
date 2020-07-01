package find_minimum_in_rotated_sorted_array_ii

func findMin(nums []int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}

	start := 0
	end := l - 1

	for {
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}

		if start+1 >= end {
			break
		}
		mid := start + (end-start)/2
		if nums[mid] < nums[end] {
			end = mid
		} else if nums[mid] == nums[end] {
			start++
			end++
		} else {
			start = mid
		}
	}

	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]
}
