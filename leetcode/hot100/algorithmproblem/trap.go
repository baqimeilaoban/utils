package algorithmproblem

// Trap 接雨水
/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例2：
输入：height = [4,2,0,3,2,5]
输出：9
*/
func Trap(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}