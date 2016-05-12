// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Create("myfile.txt")
	defer file.Close()

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		file.WriteString(<-ch + "\n") // receive from channel ch and write to file.
	}
	file.WriteString(strconv.FormatFloat(time.Since(start).Seconds(), 'f', 6, 64) + "s elapsed\n")
}

func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "http://www."){
		url = "http://" + url
	}
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-