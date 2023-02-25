package erro

func diff(first []string, second []string) []string {
	var diff []string
	for _, str := range first {
		if !contains(second, str) {
			diff = append(diff, str)
		}
	}
	return diff
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
