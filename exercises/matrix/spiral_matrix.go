package matrix

func SpiralMatrix(n int) [][]int {
	var result = make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	var (
		startColumn int
		endColumn   = n - 1
		startRow    int
		endRow      = n - 1
		counter     = 1
	)
	for counter <= n*n {
		for c := startColumn; c <= endColumn; c++ {
			result[startRow][c] = counter
			counter++
		}
		startRow++
		for r := startRow; r <= endRow; r++ {
			result[r][endColumn] = counter
			counter++
		}
		endColumn--

		for c := endColumn; c >= startColumn; c-- {
			result[endRow][c] = counter
			counter++
		}
		endRow--
		for r := endRow; r >= startRow; r-- {
			result[r][startColumn] = counter
			counter++
		}
		startColumn++
	}
	return result
}
