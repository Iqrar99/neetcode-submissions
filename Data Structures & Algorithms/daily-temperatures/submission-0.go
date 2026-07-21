func dailyTemperatures(t []int) []int {
	stackIdx := []int{}
	stackIdx = append(stackIdx, 0)
	res := make([]int, len(t))
	
	for i := 1; i < len(t); i++ {
		for len(stackIdx) > 0 && t[i] > t[top(stackIdx)] {
			res[top(stackIdx)] = i - top(stackIdx)
			stackIdx = pop(stackIdx)
		}

		stackIdx = append(stackIdx, i)
	}

	return res
}

func top(stack []int) int {
	if len(stack) > 0 {
		return stack[len(stack)-1]
	}
	return -1
}

func pop(stack []int) []int {
	if len(stack) > 0 {
		return stack[:len(stack)-1]
	}
	return stack
}
