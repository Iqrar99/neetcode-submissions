func topKFrequent(nums []int, k int) []int {
	numCnt := make(map[int]int)
	for _, n := range nums {
		numCnt[n]++
	}

	var pairs [][2]int
	for k, v := range numCnt {
		pairs = append(pairs, [2]int{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	var res []int
	for i := 0; i < k; i++ {
		res = append(res, pairs[i][0])
	}
	return res
}
