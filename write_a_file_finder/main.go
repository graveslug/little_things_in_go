package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a director")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	var total int

	for _, file := range files {
		if file.Size() == 0 {
			total += len(file.Name()) + 1
		}
	}

	fmt.Printf("Total required space: %d bytes.\n", total)

	names := make([]byte, 0, total)

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
		}

	}

	err = ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", names)
}
