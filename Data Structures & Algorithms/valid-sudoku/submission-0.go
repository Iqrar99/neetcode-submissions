func isValidSudoku(board [][]byte) bool {
	rNote := make([]map[byte]bool, 9)
	cNote := make([]map[byte]bool, 9)
	sqrs := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rNote[i] = make(map[byte]bool)
		cNote[i] = make(map[byte]bool)
		sqrs[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c< 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			val := board[r][c]
			sqrIdx := (r/3)*3 + c/3

			if rNote[r][val] || cNote[c][val] || sqrs[sqrIdx][val] {
				return false
			}

			rNote[r][val] = true
			cNote[c][val] = true
			sqrs[sqrIdx][val] = true
		}
	}

	return true
}
