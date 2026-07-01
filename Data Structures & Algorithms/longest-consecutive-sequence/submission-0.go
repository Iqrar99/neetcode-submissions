func longestConsecutive(nums []int) int {
	numSet := make(map[int]struct{})
	for _, n := range nums {
		numSet[n] = struct{}{}
	}

	maxStreak := 0
	for n := range numSet {
		if _, found := numSet[n-1]; !found {
			length := 1
			for {
				if _, exists := numSet[n+length]; exists {
					length++
				} else {
					break
				}
			}
			if length > maxStreak {
				maxStreak = length
			}
		}
	}
	
	return maxStreak
}
