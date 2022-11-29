package easy

// removeNthFromEnd 删除链表的倒数第N个节点
/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
示例1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例2：
输入：head = [1], n = 1
输出：[]
示例3：
输入：head = [1,2], n = 1
输出：[1]
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	tmp := head
	for tmp != nil {
		length++
		tmp = tmp.Next
	}
	dumy := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := dumy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dumy.Next
}
