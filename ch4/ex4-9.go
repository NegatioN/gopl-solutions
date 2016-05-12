package main


import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"path/filepath"
)

func main() {
	wordCounts := make(map[string]int)

	absPath, _ := filepath.Abs("inputData/textfile.txt")
	file, err := os.Open(absPath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) //set scanner to split on words, not lines

	for scanner.Scan(){
		wordCounts[scanner.Text()] += 1
	}

	//print words and wordcount
	for key, value := range wordCounts{
		fmt.Println(key + " : " + strconv.Itoa(value))
	}

}
