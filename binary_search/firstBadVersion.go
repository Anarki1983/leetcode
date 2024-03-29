package leetcode

//Runtime: 0 ms, faster than 100.00% of Go online submissions for First Bad Version.
//Memory Usage: 1.9 MB, less than 87.79% of Go online submissions for First Bad Version.

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// time complexity: O(logN)
// space complexity: O(1)

func firstBadVersion(n int) int {
	// binary search
	// lo = 1
	// hi = n
	// mi = (lo + hi) / 2
	// for lo <= mi {
	//   if isBadVersion(mi) == true, bad version < mi
	//      if mi == 1 || isBadVersion(mi-1) == false
	//          return mi
	//      hi = mi-1
	//   else bad version > mi
	//      lo = mi+1
	// }
	// return 1

	lo := 1
	hi := n
	mi := (lo + hi) / 2
	for lo <= hi {
		if isBadVersion(mi) {
			if mi == 1 || !isBadVersion(mi-1) {
				return mi
			}

			hi = mi - 1
		} else {
			lo = mi + 1
		}

		mi = (lo + hi) / 2
	}

	return 1
}
