package greedy

import "sort"

// K 次取反后最大化的数组和
/*
给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：
选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
重复这个过程恰好 k 次。可以多次选择同一个下标 i 。
以这种方式修改数组后，返回数组 可能的最大和 。
示例1：
输入：nums = [4,2,3], k = 1
输出：5
解释：选择下标 1 ，nums 变为 [4,-2,3] 。
示例2：
输入：nums = [3,-1,0,2], k = 3
输出：6
解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。
示例3：
输入：nums = [2,-3,-1,5,-4], k = 2
输出：13
解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。
*/
// 局部最优：让绝对值大的负数变为正数，当前数值变为最大
// 整体最优：整个数组和达到最大
// 第一步：将数组按照绝对值大小进行排序
// 第二步：从前向后进行遍历，遇到负数将其变为正数，同时k--
// 第三步：如果K还⼤于0，那么反复转变数值最⼩的元素，将K⽤完
// 第四步：求和
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] *= -1
			k--
		}
	}
	if k%2 == 1 {
		sort.Ints(nums)
		nums[0] *= -1
	}
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	return ans
}
