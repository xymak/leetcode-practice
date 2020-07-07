package subsets

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

*/

func subsets(nums []int) [][]int {
	output := &([][]int{})
	var current []int

	backtrack(0, current, nums, output)

	return *output
}

func backtrack(first int, current []int, nums []int, output *[][]int) {
	t := make([]int, len(current))
	copy(t, current)
	*output = append(*output, t)

	for i := first; i < len(nums); i++ {
		// 放入
		current = append(current, nums[i])
		backtrack(i+1, current, nums, output)

		// 回溯
		lenOfCurrent := len(current)
		current = current[:lenOfCurrent-1]
	}
}
