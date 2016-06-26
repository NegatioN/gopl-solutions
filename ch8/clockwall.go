package main

import (
	"strings"
	"os"
	"net"
	"io"
	"sync"
	"fmt"
)

type host struct {
	name    string
	address string
	port    string
}

func main() {
	hosts := make([]host, 0)
	if len(os.Args) > 1 {
		for i, str := range os.Args {
			if i == 0 {
				continue
			}
			split1 := strings.Split(str, "=")
			name := split1[0]
			split2 := strings.Split(split1[1], ":")
			address := split2[0]
			port := split2[1]
			hosts = append(hosts, host{name, address, port})
		}
	}
	output := make(map[host]string)
	var waitgroup sync.WaitGroup
	var lock sync.RWMutex
	for _, host := range hosts {
		waitgroup.Add(1)
		go getServerTime(host, output, &waitgroup,lock)
	}
	waitgroup.Wait()
	printTable(output)
}

func getServerTime(myhost host, outputMap map[host]string, waitgroup *sync.WaitGroup, mapLock sync.RWMutex){
	defer waitgroup.Done()

	hoststr := myhost.address + ":" + myhost.port
	conn, err := net.Dial("tcp", hoststr)
	b := make([]byte, 1024)
	n, err := conn.Read(b)
	if err == io.EOF {
		os.Exit(0)
	}
	if err == nil && n > 0 {
		mapLock.Lock()
		outputMap[myhost] = string(b[:len(b)])
		mapLock.Unlock()
	}
}

func printTable(outputMap map[host]string){
	var finalOutput string
	var counter int
	for host, output := range outputMap{
		counter+=1
		if counter % 4 == 0{
			finalOutput += "\n"
		}
		finalOutput += "\t" + host.name + " - " + output
	}
	fmt.Print(finalOutput)
}