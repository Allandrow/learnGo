package utils

import "unicode"

func ToUpper(str string) string {
	runes := []rune(str)

	for i := range runes {
		runes[i] = unicode.ToUpper(runes[i])
	}

	return string(runes)
}