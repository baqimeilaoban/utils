package algorithmproblem

// climbStairs 爬楼梯
/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
示例1：
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
示例2：
输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
*/
// 动态规划 dp[i] = d[i-1] + dp[i-2] 到达第i阶必定要先到达i-1阶或者i-2阶
func climbStairs(n int) int {
	// 边界条件，以防panic
	if n == 1 {
		return 1
	}
	// 创建一个存放到达i阶需要多少步的数组
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
