package leetcode

import "math"

// Runtime: 24 ms, faster than 67.33% of Go online submissions for Count Complete Tree Nodes.
// Memory Usage: 7 MB, less than 94.00% of Go online submissions for Count Complete Tree Nodes.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := getLeftHeight(root.Left)
	rh := getRightHeight(root.Right)

	// perfact binary tree
	if lh == rh {
		return int(math.Pow(2, float64(lh+1))) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func getLeftHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + getLeftHeight(node.Left)
}

func getRightHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + getRightHeight(node.Right)
}
