package greedy

import "fmt"

// partitionLabels 划分字母区间
/*
字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。
示例1：
输入：S = "ababcbacadefegdehijhklij"
输出：[9,7,8]
解释：
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
*/
func partitionLabels(s string) []int {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	fmt.Println(lastPos)
	var partition []int
	start, end := 0, 0
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return partition
}
