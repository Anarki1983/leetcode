package leetcode

/*
*
  - Definition for a binary tree node.

// Runtime: 4 ms, faster than 91.60% of Go online submissions for Path Sum.
// Memory Usage: 4.6 MB, less than 93.37% of Go online submissions for Path Sum.

// time complexity:  O(N)
// space complexity: O(1)

  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return hasPathSumDfs(root, targetSum, 0)
}

func hasPathSumDfs(node *TreeNode, targetSum, parentSum int) bool {
	if node.Left == nil && node.Right == nil {
		if parentSum+node.Val == targetSum {
			return true
		}
	}

	if node.Left != nil {
		if hasPathSumDfs(node.Left, targetSum, parentSum+node.Val) {
			return true
		}
	}

	if node.Right != nil {
		if hasPathSumDfs(node.Right, targetSum, parentSum+node.Val) {
			return true
		}
	}

	return false
}
