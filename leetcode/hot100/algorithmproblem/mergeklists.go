package algorithmproblem

// MergeKLists 合并k个升序链表
/*
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
示例1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例2：
输入：lists = []
输出：[]
示例3：
输入：lists = [[]]
输出：[]
*/
// 遍历链表数组
func MergeKLists(lists []*ListNode) *ListNode {
	// 首先判断是否链表数组有效
	if lists == nil && len(lists) == 0 {
		return nil
	}
	var res *ListNode
	for _, list := range lists {
		res = mergeTwoList(res, list)
	}
	return res
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	dumy := &ListNode{}
	p := dumy
	// 记住是for循环
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dumy.Next
}
