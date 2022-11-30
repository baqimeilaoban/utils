package easy

// isPalindromeList 回文链表
/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
示例1：
输入：head = [1,2,2,1]
输出：true
示例2：
输入：head = [1,2]
输出：fals
*/
func isPalindromeList(head *ListNode) bool {
	var tmp []int
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	i, j := 0, len(tmp)-1
	for i < j {
		if tmp[i] != tmp[j] {
			return false
		}
		i++
		j--
	}
	return true
}
