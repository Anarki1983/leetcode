package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return nil
	}

	if head == head.Next {
		return head
	}

	slow := head
	fast := head

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		} else {
			fast = fast.Next
		}

		// detect cycle
		if slow == fast {
			slow2 := head

			// use two slow to detect cycle pos
			for slow.Next != nil && slow2.Next != nil {
				if slow == slow2 {
					return slow
				}

				slow = slow.Next
				slow2 = slow2.Next
			}
		}
	}

	return nil
}
