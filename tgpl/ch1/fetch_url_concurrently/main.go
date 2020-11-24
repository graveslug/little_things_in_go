//Fetchall fetches URLS in parallel and reports their times and sizes
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660) //Opens file and creates one if there isn't one. Appends if ones already created
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Fprintln(file, <-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.fs %7d %s", secs, nbytes, url)
}

//cmd to run app
// go run main.go http://google.com http://facebook.com http://youtube.com http://baidu.com http://yahoo.com http://amazon.com http://qq.com http://twitter.com http://taobao.com http://live.com http://google.co.in

// //Heavily commented version
// func main() {
// 	// time.Now() makes use of the time package and the now function which starts a timer for our function.
// 	start := time.Now()
// 	//Opens a file when called and appends the information to it.
// 	//if no file exists it creates it.
// 	//Errors else file
// 	//idk what the number is for. Error log maybe?
// 	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
// 	//closes file to prevent leaking
// 	defer file.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(3)
// 	}
// 	//we use make to create an empty input which here it is a channel that takes an argument of a string
// 	ch := make(chan string)
// 	// this range function starts a new goroutine for each argument. It tosses the index. We only need the URL.
// 	for _, url := range os.Args[1:] {
// 		// this makes the go routine and calls the fetch function with the argument of the current URL and ch
// 		go fetch(url, ch)
// 	}
// 	for range os.Args[1:] {
// 		//goes through the arguments and prints each channel with all its information
// 		//appends information tp each channel to the file
// 		fmt.Fprintln(file, <-ch)
// 	}
// 	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
// }

// func fetch(url string, ch chan<- string) {
// 	//Times the function using the time package with the Now() function that starts a timer
// 	start := time.Now()
// 	//uses the http package with the Get() function to retrieve the URL that we passed over as a string
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		//if errors we will pass the error message to the channel
// 		ch <- fmt.Sprint(err)
// 		return
// 	}
// 	//Copy() copies from the src (resp.Body) to the dst (ioutil.Discard) until either EOF is reached on src or an error occurs. It returns the number of bytes copied and the first error encountered while copying, if any.
// 	//Discard is an io.Writer on which all Write calls succed without doing anything. Basically does nothing.
// 	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
// 	//Close the file to prevent leaking.
// 	resp.Body.Close()
// 	if err != nil {
// 		//pass the URL with the Err message to the channel. Most likely looks like an EOF error due to how Copy() works. %s prints it as a string w/ the argument of url. %v prints it as its original value of the err.
// 		ch <- fmt.Sprintf("while reading %s: %v", url, err)
// 		return
// 	}
// 	//Takes the start time and returns the duration as a floating point number of seconds
// 	secs := time.Since(start).Seconds()
// 	//%.2fs is print the number as a floating point number with a decimal place of .2 and print it as a string represting it as seconds. %7d stands for print as a decimal integer with a max of seven digits repersenting bytes. %s repersents a string to print out the url as a string
// 	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
// }
