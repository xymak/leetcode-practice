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

func subsetsWithDup(nums []int) [][]int {
	var r [][]int
	var c []int

	backtrack(nums, 0, c, &r)
	return r
}

func backtrack(nums []int, f int, c []int, r *[][]int) {
	t := make([]int, len(c))
	copy(t, c)
	*r = append(*r, c)

	for i:=f;i<len(nums);i++ {
		//fmt.Println(c, nums[i])
		if i != f && nums[i] == nums[i-1] {
			continue
		}
		c = append(c, nums[i])
		newC := make([]int, len(c))
		copy(newC, c)
		backtrack(nums, i+1, newC, r)

		c = c[:len(c)-1]
	}

}
