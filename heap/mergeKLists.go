package leetcode

import "container/heap"

// time complexity: O(N+KlogK)
// space complexity: O(K)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	h := &MergeKListsHeap{}
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	if h.Len() == 0 {
		return nil
	}

	head := heap.Pop(h).(*ListNode)
	curr := head

	for {
		if curr.Next != nil {
			heap.Push(h, curr.Next)
		}

		if h.Len() == 0 {
			break
		}

		node := heap.Pop(h).(*ListNode)
		curr.Next = node
		curr = curr.Next
	}

	return head
}

type MergeKListsHeap []*ListNode

func (h MergeKListsHeap) Len() int           { return len(h) }
func (h MergeKListsHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MergeKListsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MergeKListsHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MergeKListsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
