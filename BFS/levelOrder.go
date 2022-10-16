package leetcode

// Runtime: 2 ms, faster than 67.35% of Go online submissions for Binary Tree Level Order Traversal.
// Memory Usage: 3.2 MB, less than 5.72% of Go online submissions for Binary Tree Level Order Traversal.

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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// Create a queue and insert root
	queue := []*TreeNode{root}

	// Create result slice
	result := make([][]int, 0, 0)

	// Process as long as queue is not empty
	for len(queue) > 0 {
		// Get the current size or length of the queue.
		// This indicates the total number of nodes that are part of current level
		sz := len(queue)
		level := make([]int, 0, sz)
		for i := 0; i < sz; i++ {
			// Remove a node
			node := queue[0]
			queue = queue[1:]

			// Visit the node. Here visiting means collecting it into the output array
			level = append(level, node.Val)

			// Insert children of the node into the queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// level is filled with one level of nodes' values. Insert this into the final
		// result
		result = append(result, level)
	}

	return result
}
