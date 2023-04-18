package code

/*
1026. 节点与其祖先之间的最大差值
给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。

（如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func dfs(root *TreeNode, min_val, max_val int) (res int) {
	if root == nil {
		return
	}
	res = max(abs(max_val-root.Val), abs(root.Val-min_val))
	min_val = min(min_val, root.Val)
	max_val = max(max_val, root.Val)
	res = max(res, dfs(root.Left, min_val, max_val))
	res = max(res, dfs(root.Right, min_val, max_val))
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
