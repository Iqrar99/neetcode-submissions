func maxArea(heights []int) int {
	result := 0
	l, r := 0, len(heights) - 1

	for l < r {
		area := min(heights[l], heights[r]) * (r-l)
		result = max(result, area)
		if heights[l] <= heights[r] {
			l++
		} else {
			r--
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}