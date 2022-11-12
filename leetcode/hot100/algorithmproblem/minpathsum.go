package algorithmproblem

// MinPathSum 最小路径和
/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
示例1：
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例2：
输入：grid = [[1,2,3],[4,5,6]]
输出：12
*/
func MinPathSum(grid [][]int) int {
	// 获取网格的m与n
	rows, cols := len(grid), len(grid[0])
	// 声明存放最短路径的dp数组
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		// 初始化数组
		dp[i] = make([]int, cols)
	}
	// 第一步的长度固定为grid[0][0]
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		// 边界路径只能为grid[i][0]+grid[i-1][0]+...+grid[0][0]
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for j := 1; j < cols; j++ {
		// 同理
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			// 存放到达该路径的最短距离
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}
	return dp[rows-1][cols-1]
}
