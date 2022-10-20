package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var nodes []*TreeNode

func flatten(root *TreeNode) {
	// inorder traversal and push node to stack
	nodes = make([]*TreeNode, 0, 0)
	inOrderTraversal(root)

	for i := range nodes {
		nodes[i].Left = nil
		nodes[i].Right = nil
		if i+1 < len(nodes) {
			nodes[i].Right = nodes[i+1]
		}
	}
}

func inOrderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	// push node to stack
	nodes = append(nodes, node)

	// push node.Left to stack
	inOrderTraversal(node.Left)

	// push node.Right to stack
	inOrderTraversal(node.Right)
}
