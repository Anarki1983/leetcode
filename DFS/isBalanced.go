package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time complexity:  O(N)
// space complexity: O(1)

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	l, r := 0, 0
	l = getDepth(root.Left)
	if l == -1 {
		return false
	}

	r = getDepth(root.Right)
	if r == -1 {
		return false
	}

	if l > r {
		return l-r <= 1
	} else if r > l {
		return r-l <= 1
	}

	// l == r
	return true
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	if node.Left == nil && node.Right == nil {
		return 1
	}

	l, r := 0, 0
	if node.Left != nil {
		l = getDepth(node.Left)
		if l == -1 {
			return -1
		}
	}
	if node.Right != nil {
		r = getDepth(node.Right)
		if r == -1 {
			return -1
		}
	}

	// compare if l vs r balanced
	if l > r {
		if l-r > 1 {
			return -1
		} else {
			return l + 1
		}
	} else if r > l {
		if r-l > 1 {
			return -1
		} else {
			return r + 1
		}
	}

	// l == r
	return l + 1
}
