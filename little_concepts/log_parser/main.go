package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	//Asks for users input ie: go run main.go fool < passage.txt
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a search word.")
	}
	query := args[0]

	rx := regexp.MustCompile(`[^a-z]+`)

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		rx.ReplaceAllString(word, "")

		if len(word) > 2 {
			words[word] = true
		}
	}

	for word := range words {
		fmt.Println(word)
	}

	//queries user input
	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}
	fmt.Printf("The input does not contain %q.\n", query)
}
