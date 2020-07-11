package fei_bo_na_qi_shu_lie_lcof

func fib(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	if n == 0 {
		return 0
	}
	dp[1] = 1
	if n == 1 {
		return 1
	}

	for i:=2;i<=n;i++{
		dp[i] = (dp[i-1]+dp[i-2]) % 1000000007
	}

	return dp[n]
}
