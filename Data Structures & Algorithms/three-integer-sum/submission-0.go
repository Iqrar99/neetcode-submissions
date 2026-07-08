func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for ptr := range(nums) {
		start := nums[ptr]
		if start > 0 {
			break
		}
		if ptr > 0 && start == nums[ptr-1] {
			continue
		}

		l, r := ptr+1, len(nums)-1
		for l < r {
			sum := start + nums[l] + nums[r]
			switch true {
			case sum > 0:
				r--
			case sum < 0:
				l++
			default:
				result = append(result, []int{start, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
			
		}
	}

	return result
}
