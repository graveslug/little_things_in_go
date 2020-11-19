package main

import (
	"bufio"
	"fmt"
	"os"
)

//ex 1  Modify to print the names of all files in which each duplicated line occurs
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		//added ""
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			//refreshes the counts to prevent duplication with each iteration.
			counts := make(map[string]int)
			//added arg which takes the current file
			countLines(f, counts, arg)
			f.Close()
		}
	}
}

//added arg with a type string to pass the file name
func countLines(f *os.File, counts map[string]int, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//moved this out of main function and placed it into countLines. It cleans up the main function and has the countLines function print the statement. If this loop was in the main() it wouldn't recognize the arg string
	for line, n := range counts {
		if n > 1 {
			//prints file name

			fmt.Printf("\t\t%d\t%s\t%s\n", n, line, arg)
		}
	}
	//NOTE: ignoring potential erros from input.Err()
}

// func main() {
// 	//we create an empty map with a key of string and a value of int and assign it to counts
// 	counts := make(map[string]int)
// 	//we take an argument from the terminal and assign it the files. The args takes the length of the input and its full length
// 	files := os.Args[1:]
// 	//take the len of the files and if its equal to zero
// 	if len(files) == 0 {
// 		//call the countLine function and give it the standard input and map.
// 		countLines(os.Stdin, counts)
// 	} else {
// 		//iterates through the files and tosses one value to the void
// 		for _, arg := range files {
// 			//returns to values one is an opened file (*os.File) and the other is os.Open built in error type.
// 			f, err := os.Open(arg)
// 			//if not nothing show error then continue
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			//takes the arguments (f) and passes the map to count the lines. Since the file is now opened we can count the lines inside the function.
// 			countLines(f, counts)
// 			//closes file
// 			f.Close()
// 		}
// 	}
// 	//Goes through everything adn checks the map for values greater than one to print every duplicate file.
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// //a map is a reference to the data structure created by make when the map is passed the function it makes a copy of the reference so that the underlying datastructure can be seen by main  still.
// func countLines(f *os.File, counts map[string]int) {
// 	//reads the lines and breaks it down into words or lines. In this case its lines.
// 	input := bufio.NewScanner(f)
// 	//scans the document and makes changes to the map and increments the value as needed for each duplicate line.
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	//ignoring error checking here
// }
