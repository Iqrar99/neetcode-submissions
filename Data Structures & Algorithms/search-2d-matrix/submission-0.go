func searchMatrix(matrix [][]int, target int) bool {
	oL, oR := 0, len(matrix) - 1
	for oL <= oR {
		oM := (oL+oR)/2

		currSlice := matrix[oM]
		if target >= currSlice[0] && target <= currSlice[len(currSlice)-1] {
			iL, iR := 0, len(currSlice) - 1
			for iL <= iR {
				iM := (iL+iR)/2
				if currSlice[iM] == target {
					return true
				} else if currSlice[iM] < target {
					iL = iM+1
				} else {
					iR = iM-1
				}
			}
			return false

		} else if target < currSlice[0] {
			oR = oM-1
	
		} else if target > currSlice[len(currSlice)-1] {
			oL = oM+1
		}
	}

	return false
}
