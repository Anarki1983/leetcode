package leetcode

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
// Memory Usage: 2.1 MB, less than 78.79% of Go online submissions for Swap Nodes in Pairs.

// time complexity: O(N)
// space complexity: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	currNode := head
	var prevNode *ListNode
	for currNode != nil && currNode.Next != nil {
		// swap
		tmp := currNode.Next
		currNode.Next = tmp.Next
		tmp.Next = currNode

		// concate to prev node
		if prevNode != nil {
			prevNode.Next = tmp
		}

		prevNode = currNode
		currNode = currNode.Next
	}

	return newHead
}
