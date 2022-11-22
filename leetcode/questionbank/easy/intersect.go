package easy

// intersect 两个数组的交集
/*
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
示例1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例2：
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]
*/
func intersect(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	for _, v := range nums1 {
		mp[v]++
	}
	var res []int
	for _, v := range nums2 {
		if mp[v] > 0 {
			res = append(res, v)
			mp[v]--
		}
	}
	return res
}
