package leetcode

// Runtime: 2 ms, faster than 47.79% of Go online submissions for Invert Binary Tree.
// Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Invert Binary Tree.

// time complexity:  O(N)
// space complexity: O(1)

func invertTree(root *TreeNode) *TreeNode {
	invertTreeDfs(root)

	return root
}

func invertTreeDfs(node *TreeNode) {
	if node == nil {
		return
	}

	invertTreeDfs(node.Left)
	invertTreeDfs(node.Right)

	node.Left, node.Right = node.Right, node.Left
}
