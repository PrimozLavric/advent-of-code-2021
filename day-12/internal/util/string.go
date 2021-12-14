package util

import "unicode"

// StringIsAllUppercase checks if the string contains only uppercase characters.
func StringIsAllUppercase(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// StringIsAllLowercase checks if the string contains only lowercase characters.
func StringIsAllLowercase(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
