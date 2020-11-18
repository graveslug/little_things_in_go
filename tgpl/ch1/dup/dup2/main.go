package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//we create an empty map with a key of string and a value of int and assign it to counts
	counts := make(map[string]int)
	//we take an argument from the terminal and assign it the files. The args takes the length of the input and its full length
	files := os.Args[1:]
	//take the len of the files and if its equal to zero
	if len(files) == 0 {
		//call the countLine function and give it the standard input and map.
		countLines(os.Stdin, counts)
	} else {
		//iterates through the files and tosses one value to the void
		for _, arg := range files {
			//returns to values one is an opened file (*os.File) and the other is os.Open built in error type.
			f, err := os.Open(arg)
			//if not nothing show error then continue
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			//takes the arguments (f) and passes the map to count the lines. Since the file is now opened we can count the lines inside the function.
			countLines(f, counts)
			//closes file
			f.Close()
		}
	}
	//Goes through everything adn checks the map for values greater than one to print every duplicate file.
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//a map is a reference to the data structure created by make when the map is passed the function it makes a copy of the reference so that the underlying datastructure can be seen by main  still.
func countLines(f *os.File, counts map[string]int) {
	//reads the lines and breaks it down into words or lines. In this case its lines.
	input := bufio.NewScanner(f)
	//scans the document and makes changes to the map and increments the value as needed for each duplicate line.
	for input.Scan() {
		counts[input.Text()]++
	}
	//ignoring error checking here
}
