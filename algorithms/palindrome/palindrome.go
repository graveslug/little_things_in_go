package palindrome

import (
	"fmt"
	"strings"
)

func paldindrome(inputString string) bool {
	reverseString := Reverse(inputString)
	return reverseString == inputString
}

//Reverse takes in input of string and reverses it
func Reverse(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		b.WriteByte(s[i])
		fmt.Println(b)
	}
	return b.String()
}
