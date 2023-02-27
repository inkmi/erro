package internal

func isIntInSlice(i int, s []int) bool {
	for vi := range s {
		if s[vi] == i {
			return true
		}
	}
	return false
}

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
