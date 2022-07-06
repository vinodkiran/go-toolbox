package net

import (
	"fmt"
	"net"
	"sort"
	"strconv"
)

func worker(host string, ports, results chan int) {
	for port := range ports {
		if _, err := ping(host, port); err == nil {
			results <- port
		} else {
			results <- 0
		}
	}
}

func ping(host string, port int) (int, error) {
	conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	if err != nil {
		return 0, err
	}
	conn.Close()
	return port, nil
}

func ScanPorts(host string) {
	startPort := 2
	endPort := 1024
	ports := make(chan int, 100)
	results := make(chan int)

	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(host, ports, results)
	}

	go func() {
		for i := startPort; i <= endPort; i++ {
			ports <- i
		}
	}()

	for i := startPort; i <= endPort; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openPorts)
	fmt.Printf("Open ports on %s : %v\n", host, openPorts)
}
