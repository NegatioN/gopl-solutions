package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"sync"
)

func echo(c net.Conn, shout string, delay time.Duration, waitgroup *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	waitgroup.Done()
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var waitgroup sync.WaitGroup
	for input.Scan() {
		waitgroup.Add(1)
		go echo(c, input.Text(), 4*time.Second, &waitgroup)
	}
	waitgroup.Wait()
	// NOTE: ignoring potential errors from input.Err()
	if tcpconn, ok := c.(*net.TCPConn); ok {
		tcpconn.CloseWrite()
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
