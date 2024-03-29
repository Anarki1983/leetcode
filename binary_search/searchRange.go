package leetcode

func searchRange(nums []int, target int) []int {
	// binary search for target index
	// binary search for left bound of target index
	// binary search for right bound of target index
	targetIndex := binarySearch(nums, target, 0, len(nums)-1)
	//fmt.Printf("targetIndex: %d\n", targetIndex)

	if targetIndex == -1 {
		return []int{-1, -1}
	}

	if len(nums) == 1 {
		return []int{0, 0}
	}

	leftBoundIndex := searchLeftBound(nums, target, 0, targetIndex)
	//fmt.Printf("leftBoundIndex: %d\n", leftBoundIndex)

	rightBoundIndex := searchRightBound(nums, target, targetIndex, len(nums)-1)
	//fmt.Printf("rightBoundIndex: %d\n", rightBoundIndex)

	return []int{leftBoundIndex, rightBoundIndex}
}

func binarySearch(nums []int, target int, left, right int) int {
	if left == right {
		if target == nums[left] {
			return left
		}

		return -1
	}

	//fmt.Printf("binarySearch %d, %d\n", left, right)
	for left <= right {
		if left == right {
			if target == nums[left] {
				return left
			}

			return -1
		}

		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		} else if nums[mid] > target { // target at left side
			right = mid - 1
		} else { // target at right side
			left = mid + 1
		}
	}

	return -1
}

func searchLeftBound(nums []int, target int, left, right int) int {
	//fmt.Printf("searchLeftBound %d, %d\n", left, right)
	if left == right {
		if target == nums[left] {
			return left
		}

		return -1
	}

	for left <= right {
		if left == right {
			if target == nums[left] {
				return left
			}

			return -1
		}
		mid := (left + right) / 2
		//fmt.Printf("searchLeftBound %d, %d, %d\n", left, right, mid)

		if target == nums[mid] {
			// 已經在最左邊了
			if mid == 0 {
				return 0
			}

			// 如果左邊剛好不是target, 找到左邊界
			if target != nums[mid-1] {
				//fmt.Println("左邊剛好不是target, 找到左邊界", mid)
				return mid
			}

			// 繼續往左邊找
			right = mid - 1
			//fmt.Printf("繼續往左邊找, left: %d, right: %d", left, right)
		} else if target > nums[mid] { // 跳過頭了, 往右邊找
			//fmt.Println("跳過頭了, 往右邊找")
			if target == nums[mid+1] {
				//fmt.Println("右邊剛好是target, 找到左邊界", mid+1)
				return mid + 1
			}

			left = mid + 1
		}
	}

	return -1
}

func searchRightBound(nums []int, target int, left, right int) int {
	//fmt.Printf("searchRightBound %d, %d\n", left, right)
	if left == right {
		if target == nums[right] {
			return right
		}

		return -1
	}

	for left <= right {
		if left == right {
			if target == nums[right] {
				return right
			}

			return -1
		}

		mid := (left + right) / 2
		//fmt.Printf("searchRightBound %d, %d, %d\n", left, right, mid)

		if target == nums[mid] {
			// 已經在最右邊了
			if mid == len(nums)-1 {
				return len(nums) - 1
			}

			// 如果右邊剛好不是target, 找到右邊界
			if target != nums[mid+1] {
				//fmt.Println("右邊剛好不是target, 找到右邊界", mid)
				return mid
			}

			// 繼續往右邊找
			left = mid + 1
			//fmt.Printf("繼續往右邊找, left: %d, right: %d\n", left, right)
		} else if target < nums[mid] { // 跳過頭了, 往左邊找
			//fmt.Println("跳過頭了, 往左邊找")
			if target == nums[mid-1] {
				//fmt.Println("左邊剛好是target, 找到右邊界", mid-1)
				return mid - 1
			}

			right = mid - 1
			//fmt.Printf("繼續往左邊找, left: %d, right: %d\n", left, right)
		}
	}

	return -1
}
