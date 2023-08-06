package core

func Split(s string) []string {
	chunkSize := 1000
	var result []string

	for i := 0; i < (len(s)+chunkSize-1)/chunkSize; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(s) {
			end = len(s)
		}
		result = append(result, s[start:end])
	}

	return result
}
