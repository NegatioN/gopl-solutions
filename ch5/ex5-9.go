package main

import (
	"strings"
	"fmt"
)


/**


 */

func main() {
	fmt.Println(expand("hello, $foo", exchangeOutput))
}


func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func exchangeOutput(s string) string{
	return "somestring"
}
