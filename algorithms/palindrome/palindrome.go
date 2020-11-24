package palindrome

import (
	"fmt"
	"strings"
)

//Palidrome one makes use of the string package
func paldindrome(inputString string) bool {
	reverseString := reverse(inputString)
	return reverseString == inputString
}

//Reverse takes in input of string and reverses it
func reverse(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
		fmt.Println(b)
	}
	return b.String()
}

//Plaindrome 2, a zero package choice that compares both
//sides together as the string iterates through
//half of the length
func palindrome2(inputString string) bool {
	for i := 0; i < len(inputString)/2; i++ {
		if inputString[i] != inputString[len(inputString)-1-i] {
			return false
		}
	}
	return true
}

//Palindrome 3 which takes in account for runes
func paldindrome3(inputString string) bool {
	return inputString == reverseString(inputString)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
