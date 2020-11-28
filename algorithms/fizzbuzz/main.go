package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) string {
	var result string

	for i := 1; i <= n; i++ {
		xOfThree := i%3 == 0
		xOfFive := i%5 == 0

		if xOfThree && xOfFive {
			result += "fizzbuzz"
		} else if xOfThree {
			result += "fizz"
		} else if xOfFive {
			result += "buzz"
		} else {
			result += strconv.Itoa(i) // because the question wants us to return a string we must convert the integer to a string using the strconv package that contains Itoa
		}
	}
	return result
}

func main() {
	//change the numbers here to test
	fmt.Println(fizzbuzz(30))
}
