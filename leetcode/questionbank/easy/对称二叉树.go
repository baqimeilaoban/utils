package easy

/*
给你一个二叉树的根节点 root ， 检查它是否轴对称。
示例1：
输入：root = [1,2,2,3,4,4,3]
输出：true
示例2：
输入：root = [1,2,2,null,3,null,3]
输出：false
*/
func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
