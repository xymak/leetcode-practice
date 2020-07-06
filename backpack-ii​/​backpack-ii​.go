package _backpack_ii_

/*
有 n 个物品和一个大小为 m 的背包. 给定数组 A 表示每个物品的大小和数组 V 表示每个物品的价值.

问最多能装入背包的总价值是多大?

    A[i], V[i], n, m 均为整数
    你不能将物品进行切分
    你所挑选的要装入背包的物品的总大小不能超过 m
    每个物品只能取一次

样例

样例 1:

输入: m = 10, A = [2, 3, 5, 7], V = [1, 5, 2, 4]
输出: 9
解释: 装入 A[1] 和 A[3] 可以得到最大价值, V[1] + V[3] = 9

样例 2:

输入: m = 10, A = [2, 3, 8], V = [2, 5, 8]
输出: 10
解释: 装入 A[0] 和 A[2] 可以得到最大价值, V[0] + V[2] = 10

挑战

O(nm) 空间复杂度可以通过, 不过你可以尝试 O(m) 空间复杂度吗?

*/

func backPackII(m int, A []int, V []int) int {
	// f[i][j] 前i个物品，j为能达到的重量
	// f[i][j] =f[i-1][j] 这一轮i不放入任何物品
	// 满足j-A[i-1]>=0条件，且上一轮是有状态的f[i-1][j-A[i-1]]==true

	dp := make([][]int, len(A)+1)
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, m+1)
	}
	//dp[0][0] = true

	for i := 1; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			// 这一轮i不放入任何物品
			dp[i][j] = dp[i-1][j]
			// 尝试放入
			if j >= A[i-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-A[i-1]]+V[i-1])
			}
		}
	}

	return dp[len(A)][m]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}