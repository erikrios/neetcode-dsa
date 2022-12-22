package search

func BinarySearch(target int, nums []int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2

		if nums[m] < target {
			l = m + 1
		} else if (nums[m]) > target {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}
