package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// //EX 1 not a very efficient way to build strings from the terminal as its a quadratic equation. Which means that the size of which the input increases by the size of the input.
// //command line arguments are supplied within the os packages
// func main() {
// 	//decalres two strings [s] and [sep] with a type of string
// 	var s, sep string
// 	//iterates through the length of the argument and increments by each index
// 	for i := 1; i < len(os.Args); i++ {
// 		//add s + sep + os.Args current index to build the string piece by piece and assign it to the s variable
// 		s += sep + os.Args[i]
// 		//change sep back to an empty string so we do not double add the values over time that way each value is fresh
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }
// //EX 2 This still isn't very efficient at large data.
// func main() {
// 	//short hand declaration with a type of string thats empty
// 	s, sep := "", ""
// 	//iterates through the os.Args using the range which tales a string or slice taking a index and the value of that index and tosses the index but uses the element
// 	for _, arg := range os.Args[1:] {
// 		//adds them together
// 		s += sep + arg
// 		//resets sep
// 		sep = " "
// 	}
// 	//prints value
// 	fmt.Println(s)
// }

//v3 A simpler and more efficient way of handling large data is to use the Join function in the strings package of the golang library.

// func main() {
// 	//prints out as a string
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// }

//EXTRA
//excersizing the usage of os.Args
//1 Modify to print os.Args[0] to see the name of the command that invoked it.
//2Modify the program to print the index and value of each of its arguments one per line
//3 benchmark the 3 different versions above to see the performance.
// func main() {
// 	//v3 way using as comparison to the others
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// 	//e 1
// 	fmt.Println(os.Args[0])
// 	//e 2
// 	s, sep := "", ""
// 	for i, args := range os.Args[1:] {
// 		s += sep + args
// 		fmt.Println(i, s)
// 		sep = " "
// 	}
// }

//Benchmarker

func main() {
	v1Echo()
	v2Echo()
	v3Echo()
}

//chances are pretty high that these below functions are over complicating the benchmark.
func v1Echo() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func v2Echo() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func v3Echo() {
	start := time.Now()
	strings.Join(os.Args[1:], " ")
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
