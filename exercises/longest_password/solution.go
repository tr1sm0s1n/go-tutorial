package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Solution(S string) int {
	words := strings.Fields(S) // Split the string into words
	maxLength := -1            // Initialize the result as -1 for no valid password

	for _, word := range words {
		if isValidPassword(word) {
			if len(word) > maxLength {
				maxLength = len(word) // Update max length if this word is valid
			}
		}
	}

	return maxLength
}

// Helper function to check if a word is a valid password
func isValidPassword(word string) bool {
	hasLetter := false
	hasDigit := false

	for _, ch := range word {
		if unicode.IsLetter(ch) {
			hasLetter = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		} else {
			return false // Contains invalid characters
		}
	}

	// A valid password must have at least one letter and one digit
	return hasLetter && hasDigit
}

func main() {
	fmt.Println(Solution("test 5 a0A pass007 ?xy1")) // Output: 7 (pass007)

	fmt.Println(Solution("123 456 abc ?xy1")) // Output: -1 (no valid passwords)

	fmt.Println(Solution("a1 A2 Aa3 1234ab")) // Output: 6 (1234ab)

}
