package backpack

/*
在 n 个物品中挑选若干物品装入背包，最多能装多满？假设背包的大小为 m，每个物品的大小为 A[i]

样例

样例 1:
	输入:  [3,4,8,5], backpack size=10
	输出:  9

样例 2:
	输入:  [2,3,5,7], backpack size=12
	输出:  12

挑战

O(n x m) 的时间复杂度 and O(m) 空间复杂度
如果不知道如何优化空间O(n x m) 的空间复杂度也可以通过.

*/

func backPack(m int, A []int) int {
	// 前i个物品，是否能装满j

	/*
				0 1 2 3 4 5 6 7 8 9 10
			0   1 0 0 0 0 0 0 0 0 0 0
		3	1   1 0 0 1 0 0 0 0 0 0 0
		4	2   1 0 0 1 1 0 0 1 0 0 0
		8	3   1 0 0 1 1 0 0 1 1 0 0
		5	4   1 0 0 1 1 0 0 1 1 1 0
	*/
	// f[i][j] 前i个物品，j为能达到的重量
	// f[i][j] =f[i-1][j] 这一轮i不放入任何物品
	// 满足j-A[i-1]>=0条件，且上一轮是有状态的f[i-1][j-A[i-1]]==true

	dp := make([][]bool, len(A)+1)
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true

	for i := 1; i <= len(A); i++ {
		for j := 0; j <= m; j++ {
			// 这一轮i不放入任何物品
			dp[i][j] = dp[i-1][j]
			// 如果上一轮是有状态的，尝试放入
			if j >= A[i-1] && dp[i-1][j-A[i-1]] {
				dp[i][j] = true
			}
		}
	}

	for j := m; j >= 0; j-- {
		if dp[len(A)][j] {
			return j
		}
	}

	return 0
}
