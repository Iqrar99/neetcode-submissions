func productExceptSelf(nums []int) []int {
	prods := 1
	zeroCnt := 0
	for _, n := range nums {
		if n != 0 {
			prods *= n
		} else {
			zeroCnt++
		}
	}

	var res []int
	for _, n := range nums {
		switch {
		case zeroCnt > 1:
			res = append(res, 0)
		case zeroCnt == 1:
			if n != 0 {
				res = append(res, 0)
			} else {
				res = append(res, prods)
			}
		default:
			res = append(res, int(prods/n))
		}
	}

	return res
}
