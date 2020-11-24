package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Get the directory name and if there is no argument we ask for one and terminate the program
	//IE: go run main.go name_of_file
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a director")
		return
	}
	//reads the directory path from the argument. In this program we would type go run main.go files/ and if found it will print the files out
	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int
	//counts the file size reserving space by counting the total length of everyfile in the folder and writes the exact number of bytes to reserve.
	for _, file := range files {
		if file.Size() == 0 {
			total += len(file.Name()) + 1
		}
	}
	fmt.Printf("Total required space: %d bytes.\n", total)

	//Creates an splice array of bytes
	names := make([]byte, 0, total)
	///loops through files and retrieves all empty ones.
	//also translates files from bytes to a string via file.Name
	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
		}

	}
	//writes data to a file that takes the names parameter and writes them to a document. Creates a text file with the name of all empty text files of X directory.
	//Unix Octal 0644 means: User Read/Write; Group Read; Other Read.
	err = ioutil.WriteFile("out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", names)
}
