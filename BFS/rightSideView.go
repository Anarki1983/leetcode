package leetcode

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Right Side View.
// Memory Usage: 2.3 MB, less than 76.70% of Go online submissions for Binary Tree Right Side View.

// time complexity:  O(N)
// space complexity: O(d), d = max depth of the tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	leftView := make([]int, 0, 0)

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		sz := len(queue)

		leftView = append(leftView, queue[0].Val)
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}
	}

	return leftView
}
