package utils

func StringInArray(aim string, list []string) bool {
	for _, b := range list {
		if b == aim {
			return true
		}
	}
	return false
}
