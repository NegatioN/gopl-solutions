package main

import (
	"golang.org/x/net/html"
	"fmt"
	"net/http"
)


/**
Modify forEachNode so that the pre and post functions return a boolean result indicating whether to continue the traversal.
Use it to write a function ElementByID with the following signature that finds the first HTML element with the specified id
attribute. The function should stop the travesal as soon as a match is found.

 */

var shouldEvaluate bool
func main() {
	fmt.Println("ENTER")
	n := ElementByID(getHtml("http://www.finn.no"), "meta")
	fmt.Println(n.Data)
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool){
	shouldEvaluate = true
	if pre != nil {
		shouldEvaluate = !pre(n,id)
	}
	fmt.Println(shouldEvaluate)

	for c := n.FirstChild; c != nil && shouldEvaluate; c = c.NextSibling{
		forEachNode(c, id, pre, post)
	}

	if post != nil {
		post(n, id)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	forEachNode(doc,id,  findElement, nil)
	return doc
}
func findElement(n *html.Node, id string) bool {
	return n.Data == id && n.Type == html.ElementNode
}



func getHtml(url string) *html.Node {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil
	}

	return doc
}