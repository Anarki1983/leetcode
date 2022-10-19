package leetcode

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Middle of the Linked List.
// Memory Usage: 1.9 MB, less than 83.79% of Go online submissions for Middle of the Linked List.

// time complexity: O(N)
// space complexity: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	slow := head
	fast := head.Next

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return slow
		}
		fast = fast.Next
	}

	return slow
}
