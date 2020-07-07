package permutations_ii

import "sort"

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

*/

func permuteUnique(nums []int) [][]int {
	output := &([][]int{})
	current := []int{}

	sort.Ints(nums)
	backtrack(nums, current, output)

	return *output
}

func backtrack(nums []int, current []int, output *[][]int) {

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		t := []int{}
		t = append(t, nums[:i]...)
		t = append(t, nums[i+1:]...)
		current = append(current, nums[i])
		if len(t) == 0 {
			c := make([]int, len(current))
			copy(c, current)
			*output = append(*output, c)
		} else {
			backtrack(t, current, output)
			current = current[:len(current)-1]
		}
	}
}
