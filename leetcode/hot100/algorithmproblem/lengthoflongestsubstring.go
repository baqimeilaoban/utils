package algorithmproblem

// LengthOfLongestSubstring 无重复字符的最长子串 移动窗口方法
/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
示例一：
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例二：
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例3：
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
func LengthOfLongestSubstring(s string) int {
	m := make(map[byte]int) // 使用map的唯一性进行判断是否存在
	n := len(s)
	rk, res := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1]) // 删除map中的值
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		res = max(res, rk-i+1)
	}
	return res
}
