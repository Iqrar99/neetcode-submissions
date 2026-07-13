func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1
	m := (l+r) / 2
	
	for l <= r {
		if target == nums[m] {
			return m
		} else if target > nums[m] {
			l = m+1
			m = (l+r) / 2
		} else {
			r = m-1
			m = (l+r) / 2
		}
	}

	return -1
}
