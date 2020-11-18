package main

//dup1 prints the text of each line that appears
//more than once in the standard library,
//preceded by its count (line)
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//make creates an empty map of which we then map
	//a set of key value pairs of which the key
	//is a string and the value is an integer

	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	//Note: ignoring the potential errors from input,Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
