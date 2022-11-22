package easy

import "sort"

// singleNumber 只出现一次的数字
/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
示例1：
输入：nums = [2,2,1]
输出：1
示例2：
输入：nums = [4,1,2,1,2]
输出：4
示例3：
输入：nums = [1]
输出：1
*/
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

// 位运算
func singleNumberOpt(nums []int) int {
	single := 0
	for _, v := range nums {
		single ^= v
	}
	return single
}
