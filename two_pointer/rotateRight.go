package leetcode

// time complexity: O(N)
// space complexity: O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// find kth node from end
	slow := head
	fast := head

	// move fast k times
	for i := 0; i < k; i++ {
		fast = fast.Next
		if fast == nil {
			fast = head
		}
	}

	// move fast & slow until fast.Next = tail
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	fast.Next = head
	head = slow.Next
	slow.Next = nil

	return head
}
