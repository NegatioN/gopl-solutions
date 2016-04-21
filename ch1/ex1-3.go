package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"
)

func main() {
	concatables := make([]string, 900000000)
	for i := 0; i < 1000; i++{
		concatables[i] = "abby"
	}

	stringer := ""
	start_concat := time.Now()
	for _, string_at_index := range concatables{
		stringer += string_at_index
	}
	secs_concat := time.Since(start_concat).Seconds()
	start_join := time.Now()
	strings.Join(concatables, " ")
	secs_join := time.Since(start_join).Seconds()



	fmt.Println("Concat took: " + formatFloat(secs_concat) + " seconds")
	fmt.Println("Join took: " + formatFloat(secs_join) + " seconds")
	fmt.Println("Concat took: " + formatFloat((secs_concat/secs_join)*100) + " percent of join's time.")

}

func formatFloat(vari float64)string{
	return strconv.FormatFloat(vari, 'f', 6, 64)
}
