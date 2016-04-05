package main

import "fmt"
import "strconv"

func main() {
	argArr := [4]string {"name", "birth", "turkey", "games"}
	for id, arg := range argArr {
		fmt.Println("arg nr:" + strconv.Itoa(id) + " : " + arg)
	}
}