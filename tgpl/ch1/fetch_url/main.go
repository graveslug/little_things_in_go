package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// //use this cmd in terminal: go run main.go http://gopl.io
// func main() {
// 	for _, url := range os.Args[1:] {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}
// 		b, err := ioutil.ReadAll(resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("%s", b)
// 	}
// }

//Finished excersize as one whole
func main() {
	for _, url := range os.Args[1:] {
		//excer 2 here
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//excer 1 here
		_, err = io.Copy(os.Stdout, resp.Body)
		//excersize 3
		fmt.Printf("HTTP status: %d\n", resp.StatusCode)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

//excersize 1 call the function io.Copy(dst, src) reads from the src and writes to dst. Use it instead of the ioutil.ReadAll to copy the respobse body to os.Stdout without requiring a buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy
// func main() {
// 	for _, url := range os.Args[1:] {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}
// 		_, err = io.Copy(os.Stdout, resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 	}
// }
//excersize 2 Modify fetch to add the prefix http:// to each argument URL if it is missing. Try using strings.HasPrefix
// func main() {
// 	for _, url := range os.Args[1:] {
// 		if !strings.HasPrefix(url, "http://") {
// 			url = "http://" + url
// 		}
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}
// 		b, err := ioutil.ReadAll(resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("%s", b)
// 	}
// }
// //excersize 3 modify fetch to also print the HTTP status code found in resp.Status
// func main() {
// 	for _, url := range os.Args[1:] {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
// 			os.Exit(1)
// 		}
// 		b, err := ioutil.ReadAll(resp.Body)
// 		fmt.Printf("HTTP status %d\n", resp.StatusCode)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("%s", b)
// 	}
// }
