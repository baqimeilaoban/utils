package easy

// strStr 实现strStr()
/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
如果 needle 不是 haystack 的一部分，则返回  -1 。
示例1：
输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例2：
输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。
*/
func strStr(hayStack, needle string) int {
	n, m := len(hayStack), len(needle)
outer:
	for i := 0; i+m <= n; i++ {
		for j := range needle {
			if hayStack[i+j] != needle[j] {
				continue outer
			}
		}
		return i
	}
	return -1
}
