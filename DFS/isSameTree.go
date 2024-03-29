package leetcode

// time complexity:  O(N)
// space complexity: O(1)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// compare root
	if p == nil && q == nil {
		return true
	}

	if p != nil && q == nil {
		return false
	}
	if p == nil && q != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	// compare left
	if !isSameTree(p.Left, q.Left) {
		return false
	}

	// compare right
	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
