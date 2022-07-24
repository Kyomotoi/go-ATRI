package utils

func GetSliceByRange(d [][]string, start int, end int) [][]string {
	if start < 0 {
		start = 0
	}
	if end > len(d) {
		end = len(d)
	}
	return d[start:end]
}
