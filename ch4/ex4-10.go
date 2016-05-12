package main

//mod issues to report the results in age categories. less than a month old, less than a year old and more than a year old.

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("Less than a month old:")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 <= 30 {
			printIssue(item)
		}
	}

	fmt.Println("Less than a year old:")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 <= 365 {
			printIssue(item)
		}
	}

	fmt.Println("More than a year old:")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 > 365 {
			printIssue(item)
		}
	}
}

func printIssue(issue *github.Issue){
	fmt.Printf("#%-5d %9.9s %.55s\n",
		issue.Number, issue.User.Login, issue.Title)
}



