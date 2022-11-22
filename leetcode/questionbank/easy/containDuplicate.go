package easy

// containsDuplicate 存在重复元素
/*
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
示例1：
输入：nums = [1,2,3,1]
输出：true
示例2：
输入：nums = [1,2,3,4]
输出：false
示例3：
输入：nums = [1,1,1,3,3,4,3,2,4,2]
输出：true
*/
func containsDuplicate(nums []int) bool {
	mp := make(map[int]int)
	for _, v := range nums {
		_, ok := mp[v]
		if ok {
			mp[v]++
		} else {
			mp[v] = 1
		}
	}
	for _, v := range mp {
		if v >= 2 {
			return true
		}
	}
	return false
}
