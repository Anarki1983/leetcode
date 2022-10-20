package leetcode

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := &findKthLargestHeap{}
	for _, num := range nums {
		heap.Push(h, num)
	}

	num := 0
	for i := 0; i < k; i++ {
		if h.Len() > 0 {
			num = heap.Pop(h).(int)
		} else {
			return -1
		}
	}

	return num
}

type findKthLargestHeap []int

func (h findKthLargestHeap) Len() int           { return len(h) }
func (h findKthLargestHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h findKthLargestHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *findKthLargestHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *findKthLargestHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
