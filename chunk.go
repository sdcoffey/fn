package fn

// Chunk takes a slice of T and returns a slice of slices of T, each of length chunkSize.
// If len(items) % chunkSize != 0, the last slice will be shorter than chunkSize.
func Chunk[T any](items []T, chunkSize int) [][]T {
	result := make([][]T, 0, len(items)/chunkSize+1)
	for i := 0; i < len(items); i += chunkSize {
		end := i + chunkSize
		if end > len(items) {
			end = len(items)
		}
		result = append(result, items[i:end])
	}
	return result
}

// ChunkWhile takes a slice of T and returns a slice of slices of T, each of which is a contiguous chunk of items
// for which `chunker` returns true.
// `chunker` is called with the item before and the item after the current item.
func ChunkWhile[T any](items []T, chunker func(eltBefore, eltAfter T) bool) [][]T {
	result := make([][]T, 0, len(items))
	if len(items) == 0 {
		return result
	}
	chunk := make([]T, 0, len(items))
	chunk = append(chunk, items[0])
	for i := 1; i < len(items); i++ {
		if chunker(items[i-1], items[i]) {
			chunk = append(chunk, items[i])
		} else {
			result = append(result, chunk)
			chunk = make([]T, 0, len(items))
			chunk = append(chunk, items[i])
		}
	}
	result = append(result, chunk)
	return result
}
