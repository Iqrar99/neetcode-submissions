func evalRPN(tokens []string) int {
	var stack []int
	for _, t := range tokens {
		if !isOperand(t) {
			conv, _ := strconv.Atoi(t)
			stack = append(stack, conv)
		} else {
			res := calc(stack[len(stack)-2], stack[len(stack)-1], t)
			stack = append(stack[:len(stack)-2], res)
		}
	}

	return stack[len(stack)-1]
}

func isOperand(t string) bool {
	if t == "+" || t == "-" || t == "*" || t == "/" {
		return true
	}
	return false
}

func calc(a, b int, operand string) int {
	switch operand {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
	}
	return 0
}