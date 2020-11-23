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
	defer file.Close()                                                                //In Go it is common practice to defer and close the file after opening to prevent leaking
	if err != nil {
		fmt.Println(err)
		os.Exit(3) //in Go when we exit we print a status message. 0 stands for success while 1-125 means there was an error
	}
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //starts a go routine
	}
	for range os.Args[1:] {
		fmt.Fprintln(file, <-ch) //recieve from a channel ch and file
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
	resp.Body.Close() //don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.fs %7d %s", secs, nbytes, url)
}

//cmd to run app
// go run main.go http://google.com http://facebook.com http://youtube.com http://baidu.com http://yahoo.com http://amazon.com http://qq.com http://twitter.com http://taobao.com http://live.com http://google.co.in
