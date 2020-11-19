package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//create an empty map with the key of string
	//and value of integer
	counts := make(map[string]int)
	//toss the index and use the argument
	//take the names in the whole length of theos.Args input
	for _, filename := range os.Args[1:] {
		//ioutil.ReadFile takes the whole file and reads it
		//it takes a variable declaration and an err
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		//once the file is read it is passed to the split function that is part of the strings package.
		//After splitting the files into newlines it increments each that is the same
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	//checks the map counts and if greater than 1 duplicate then it prints it
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
