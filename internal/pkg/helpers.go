package pkg

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func StringNotEmpty(s string) bool {
	return s != ""
}
