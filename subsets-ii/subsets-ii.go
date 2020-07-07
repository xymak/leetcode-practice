package subsets_rii

import "sort"

/*
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

*/

func subsetsWithDup(nums []int) [][]int {
	var r [][]int
	var c []int

	sort.Ints(nums)

	backtrack(nums, 0, c, &r)
	return r
}

func backtrack(nums []int, f int, c []int, r *[][]int) {
	t := make([]int, len(c))
	copy(t, c)
	*r = append(*r, c)

	for i:=f;i<len(nums);i++ {
		if i != f && nums[i] == nums[i-1] {
			continue
		}
		c = append(c, nums[i])
		backtrack(nums, i+1, c, r)

		c = c[:len(c)-1]
	}

}
