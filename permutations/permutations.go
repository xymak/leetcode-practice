package permutations

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

 */

func permute(nums []int) [][]int {
	output := &([][]int{})
	current := []int{}
	backtrack(nums, current, output)

	return *output
}

func backtrack(nums []int, current []int, output *[][]int) {

	for i:=0;i<len(nums);i++ {
		t := []int{}
		t = append(t,nums[:i]...)
		t = append(t,nums[i+1:]...)
		current = append(current,nums[i])
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