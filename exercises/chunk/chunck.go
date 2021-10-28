package chunk

func Chunk(arr []int, size int) [][]int {
	var chunks [][]int
	for i := 0; i < len(arr); i += size {
		chunkSize := size
		if i+chunkSize >= len(arr) {
			chunkSize = len(arr) - i
		}
		chunks = append(chunks, arr[i:i+chunkSize])
	}
	return chunks
}
