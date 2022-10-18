package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var result []int

func inorderTraversal(root *TreeNode) []int {
	result = make([]int, 0, 3)

	traversal(root)

	return result
}

func traversal(node *TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil {
		traversal(node.Left)
	}

	result = append(result, node.Val)

	if node.Right != nil {
		traversal(node.Right)
	}
}
