package easy

// firstUniqChar 字符串中的第一个唯一字符
/*
给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。
示例1：
输入: s = "leetcode"
输出: 0
示例2：
输入: s = "loveleetcode"
输出: 2
示例3：
输入: s = "aabb"
输出: -1
*/
func firstUniqChar(s string) int {
	ctn := [26]int{}
	for _, v := range s {
		ctn[v-'a']++
	}
	for i, v := range s {
		if ctn[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
