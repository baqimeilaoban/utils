package easy

// mergeTwoLists 合并两个有序链表
/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例2：
输入：l1 = [], l2 = []
输出：[]
示例3：
输入：l1 = [], l2 = [0]
输出：[0]
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dumy := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := dumy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmpNode := &ListNode{
				Val:  list1.Val,
				Next: nil,
			}
			list1 = list1.Next
			cur.Next = tmpNode
		} else {
			tmpNode := &ListNode{
				Val:  list2.Val,
				Next: nil,
			}
			list2 = list2.Next
			cur.Next = tmpNode
		}
		cur = cur.Next
	}
	if list1 != nil || list2 != nil {
		if list1 != nil {
			cur.Next = list1
		}
		if list2 != nil {
			cur.Next = list2
		}
	}
	return dumy.Next
}
