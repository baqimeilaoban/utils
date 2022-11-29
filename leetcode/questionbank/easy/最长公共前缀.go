package easy

// longestCommonPrefix 最长公共前缀
/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例1：
输入：strs = ["flower","flow","flight"]
输出："fl"
示例2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

func lcp(s1, s2 string) string {
	minLen := min(len(s1), len(s2))
	index := 0
	for index < minLen && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}
