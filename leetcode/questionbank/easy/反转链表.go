package easy

// reverseList 反转链表
/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
示例1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例2：
输入：head = [1,2]
输出：[2,1]
示例3：
输入：head = []
输出：[]
*/
func reverseList(head *ListNode) *ListNode {
	// 先转为数组存起来
	var tmp []int
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	dumy := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := dumy
	for i := len(tmp) - 1; i >= 0; i-- {
		tmpNode := &ListNode{
			Val:  tmp[i],
			Next: nil,
		}
		cur.Next = tmpNode
		cur = cur.Next
	}
	return dumy.Next
}
