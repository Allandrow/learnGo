package main

import "strings"

func isValidEmail(email string) bool {
	return strings.Contains(email, "@")
}

func isValidName(name string) bool {
	return len(name) > 2
}

func isValidAmount(value int16, limit int16) bool {
	return value > 0 && value <= limit
}