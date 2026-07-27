func minEatingSpeed(piles []int, h int) int {
	l, r := 1, -1
	for _, p := range piles {
		if r < p {
			r = p
		}
	}

	var kMin int
	var k int
	for l <= r {
		k = (l+r)/2
		hTotal := 0
		for _, p := range piles {
			hTotal += p/k
			if p % k > 0 {
				hTotal++
			}
		}
		if hTotal > h {
			l = k+1
		} else {
			kMin = k
			r = k-1
		}
	}

	return kMin
}
