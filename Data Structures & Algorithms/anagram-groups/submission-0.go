func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)

	for _, s := range strs {
		var cnt [26]int
		for _, c := range s {
			cnt[c-'a']++
		}
		res[cnt] = append(res[cnt], s)
	}

	var result [][]string
	for _, group := range res {
		result = append(result, group)
	}
	return result
}
