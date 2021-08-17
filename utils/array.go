package utils

func StringInArray(aim string, list []string) bool {
	for _, i := range list {
		if i == aim {
			return true
		}
	}
	return false
}

func DeleteAimInArray(aim string, list []string) []string {
	var t []string
	for _, i := range list {
		if i != aim {
			t = append(t, aim)
		}
	}
	return t
}
