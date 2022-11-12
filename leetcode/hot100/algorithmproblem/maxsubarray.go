package algorithmproblem

// MaxSubArray 最大子数组和
/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组 是数组中的一个连续部分。
示例1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
示例2：
输入：nums = [1]
输出：1
示例3：
输入：nums = [5,4,-1,7,8]
输出：23
*/
// f(i)=max{f(i-1)+nums[i], nums[i]}
func MaxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		res = max(nums[i], res)
	}
	return res
}
