package algorithmproblem

// MaxDepth 二叉树的最大深度
/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。
示例1：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/
func MaxDepth(root *TreeNode) int {
	treeMax = 0
	maxDepthDfs(root, 0)
	return treeMax
}

func maxDepthDfs(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	depth++
	if root.Left == nil && root.Right == nil {
		if depth > treeMax {
			treeMax = depth
		}
	}
	// 左子树遍历
	maxDepthDfs(root.Left, depth)
	// 右子树遍历
	maxDepthDfs(root.Right, depth)
}

var (
	treeMax int
)
