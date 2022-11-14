package greedy

// jump 跳跃游戏2
/*
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。
示例1：
输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
示例2：
输入: nums = [2,3,0,1,4]
输出: 2
*/
// 局部最优：当前可移动距离尽可能多走，如果还没到达终点，步数再加一
// 全局最优：一步尽可能多走，从而达到最小步数
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	curDistance := 0
	ans := 0
	nextDistance := 0
	for i := 0; i < len(nums); i++ {
		nextDistance = max(nums[i]+i, nextDistance)
		if i == curDistance {
			if curDistance != len(nums)-1 {
				ans++
				curDistance = nextDistance
				if nextDistance >= len(nums)-1 {
					break
				}
			} else {
				break
			}
		}
	}
	return ans
}
