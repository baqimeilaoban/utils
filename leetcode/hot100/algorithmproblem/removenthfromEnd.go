package algorithmproblem

// RemoveNthFromEnd 删除链表的倒数第 N 个结点
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
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	// 先获取长度
	len := func(head *ListNode) int {
		res := 0
		for head != nil {
			res++
			head = head.Next
		}
		return res
	}(head)
	dumy := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := dumy
	for i := 0; i < len-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dumy.Next
}

// RemoveNthFromEndOpt 使用栈的方式处理
func RemoveNthFromEndOpt(head *ListNode, n int) *ListNode {
	var nodes []*ListNode
	dumy := &ListNode{
		Val:  0,
		Next: head,
	}
	// 进栈
	for node := dumy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prve := nodes[len(nodes)-1-n]
	prve.Next = prve.Next.Next
	return dumy.Next
}
