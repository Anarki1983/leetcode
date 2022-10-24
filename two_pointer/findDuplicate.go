package leetcode

func findDuplicate(nums []int) int {
	return twoPointerMethod(nums)
}

// sort
// Time Complexity: O(NlogN)
// Space Complexity: O(1)

// hashMap
// Time Complexity: O(N)
// Space Complexity: O(N)

// swap
// Time Complexity: O(N)
// Space Complexity: O(1)
func swapMethod(nums []int) int {
	for nums[0] != nums[nums[0]] {
		nums[0], nums[nums[0]] = nums[nums[0]], nums[0]
	}

	return nums[0]
}

// two pointer
// 這個方法有些 edge case 會跑不出來
// Time Complexity: O(N)
// Space Complexity: O(1)
func twoPointerMethod(nums []int) int {
	fast := nums[nums[0]]
	slow := nums[0]

	// phase 1, find intersect of two pointer
	for fast != slow {
		fast = nums[nums[fast]] // fast = fast.Next.Next
		slow = nums[slow]       // slow = slow.Next
	}

	// phase 2, use two slow pointer to find the "entrance" of the cycle
	slow2 := nums[0]
	for slow != slow2 {
		slow = nums[slow]
		slow2 = nums[slow2]
	}

	return slow
}
