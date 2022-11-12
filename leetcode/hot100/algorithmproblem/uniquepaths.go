package algorithmproblem

// UniquePaths 不同路径
/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
示例1：
输入：m = 3, n = 7
输出：28
示例2：
输入：m = 3, n = 2
输出：3
解释：
从左上角开始，总共有 3 条路径可以到达右下角。
1. 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右
3. 向下 -> 向右 -> 向下
示例3：
输入：m = 7, n = 3
输出：28
示例4：
输入：m = 3, n = 3
输出：6
*/
// 动态规划 dp[m][n]=dp[m-1][n]+dp[m][n-1]
// 到达{m,n} 地点的路径只能是通过 {m-1,n}往下走一步，或者{m,n-1}往左走一步
func UniquePaths(m, n int) int {
	// 定义一个存放路径数的二维数组
	dp := make([][]int, m)
	for i := range dp {
		// 边界由于只有一条路径可通，因此边界值为1
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
