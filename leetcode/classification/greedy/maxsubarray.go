package greedy

// maxSubArray 最大子序和
/*
给定⼀个整数数组 nums ，找到⼀个具有最⼤和的连续⼦数组（⼦数组最少包含⼀个元
素），返回其最⼤和。
示例1：
输⼊: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续⼦数组 [4,-1,2,1] 的和最⼤，为 6。
示例2：
输入：nums = [1]
输出：1
*/
// 暴力法 会超时
func maxSubArray(nums []int) int {
	res := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		count = 0
		for j := i; j < len(nums); j++ {
			count += nums[j]
			res = max(res, count)
		}
	}
	return res
}

// 局部最优：当前“连续和”为负数时立刻放弃，从下一个元素重新计算连续和，因为负数加上下一个元素“连续和”会越来越小
// 全局最优：选取最大“连续和”
func maxSubArrayOpt(nums []int) int {
	res := nums[0]
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if res < count {
			res = count
		}
		// 若是count为负数时，应当重新开始计数
		if count < 0 {
			count = 0
		}
	}
	return res
}

// maxSubArrayOpt2 动态规划
func maxSubArrayOpt2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		// 状态转移方程
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
