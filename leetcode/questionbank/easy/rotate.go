package easy

// rotate 旋转数组
/*
示例1：
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]
示例2：
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]
*/
// 额外数组
func rotate(nums []int, k int) {
	numsTmp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numsTmp[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, numsTmp)
}
