package main

//dup1 prints the text of each line that appears
//more than once in the standard library,
//preceded by its count (line)
//for this to work input into terminal repeated lines /then hit ctrl-d and this will bring up the input of the repeated counted lines
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
		//each time dup reads a line of input the line is used as a key into the map and the corresponding value is increamented.
		//another way to look at this is like this:
		//line := input.Text()
		//counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	//Note: ignoring the potential errors from input,Err()
	for line, n := range counts {
		if n > 1 {
			//%d = decimal integer
			//\t = tab space
			//%s = string
			//\n = newline
			// n value which is the current value
			//line is the word
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
