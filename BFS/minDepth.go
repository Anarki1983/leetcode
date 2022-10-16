package leetcode

// time complexity:  O(N)
// space complexity: O(w), w = max width of the tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// BFS
	queue := []*TreeNode{root}

	level := 1
	for len(queue) > 0 {
		sz := len(queue)

		for i := 0; i < sz; i++ {
			// pop left
			node := queue[0]
			queue = queue[1:]

			// find leaf node
			if node.Left == nil && node.Right == nil {
				return level
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}

	return 0
}
