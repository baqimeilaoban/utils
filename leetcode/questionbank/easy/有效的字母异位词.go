package easy

// isAnagram 有效的字母异位词
/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词
示例1：
输入: s = "anagram", t = "nagaram"
输出: true
示例2：
输入: s = "rat", t = "car"
输出: false
*/
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp1, mp2 := make(map[byte]int, 0), make(map[byte]int, 0)
	bs := []byte(s)
	bt := []byte(t)
	for _, v := range bs {
		mp1[v-'a']++
	}
	for _, v := range bt {
		mp2[v-'a']++
	}
	for i, v := range mp1 {
		value, ok := mp2[i]
		if !ok {
			return false
		}
		if value != v {
			return false
		}
	}
	return true
}
