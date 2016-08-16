// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
	"flag"
)

type CrawlURL struct{
	url string
	depth int
}


func crawl(url string, depth int) []CrawlURL {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	crawlUrls := make([]CrawlURL, 100)
	for i := 0; i < len(list); i+=1 {
		crawlUrls[i] = CrawlURL{list[i], depth+1}
	}
	return crawlUrls
}

//!+
func main() {
	depth := flag.Int("depth", 1, "depth of links to serach per given url")




	worklist := make(chan []CrawlURL)  // lists of URLs, may have duplicates
	unseenLinks := make(chan CrawlURL) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() {
		crawlUrls := make([]CrawlURL, 100)
		for i := 1; i < len(os.Args[1:]) -1; i+=1 {
			fmt.Println(os.Args[i])
			crawlUrls[i] = CrawlURL{os.Args[i], 1}
		}
		worklist <- crawlUrls }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if link.depth <= *depth {
					fmt.Println(link.url)
					foundCrawlUrls := crawl(link.url, link.depth)
					go func() { worklist <- foundCrawlUrls}()
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

//!-