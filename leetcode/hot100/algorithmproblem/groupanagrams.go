package algorithmproblem

import "sort"

// GroupAnagrams 字母异位分词
/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。
示例1：
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例2：
输入: strs = [""]
输出: [[""]]
示例3：
输入: strs = ["a"]
输出: [["a"]]
*/
func GroupAnagrams(strs []string) [][]string {
	var res [][]string
	mp := make(map[string][]string, 0)
	for _, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		bsStr := string(bs)
		mp[bsStr] = append(mp[bsStr], str)
	}
	for _, strings := range mp {
		res = append(res, strings)
	}
	return res
}
