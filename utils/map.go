package utils

func GetMapKeys(mp map[string]string) []string {
	j := 0
	keys := make([]string, len(mp))
	for i := range mp {
		keys[j] = i
		j++
	}
	return keys
}
