package utils

func IsAlpha(word string) bool {
	if word == "" {
		return false
	}

	for _, r := range word {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}

	return true
}

func IsNumeric(word string) bool {
	if word == "" {
		return false
	}

	for _, r := range word {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

func IsAlphanumeric(word string) bool {
	if word == "" {
		return false
	}
	
	for _, w := range word {
		if !IsAlpha(string(w)) && !IsNumeric(string(w)) {
			return false
		}
	}

	return true
}