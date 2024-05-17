package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func handleClient(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	// Read data from the client
	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		now := time.Now()
		mtime := data[:len(data)-1]
		clientTime, err := time.Parse(time.RFC3339, string(mtime))
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			return
		}
		dur := now.Sub(clientTime)
		fmt.Println("latency:", dur.Milliseconds())
	}
}

func main() {
	// Start TCP server on port 8080
	port := os.Args[1]
	lport := ":" + port

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	listener, err := net.Listen("tcp", lport)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	fmt.Println("Server listening on port 6666")
	var wg sync.WaitGroup
	// Accept and handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		//fmt.Println("Accepted connection from:", conn.RemoteAddr())
		wg.Add(1)
		go handleClient(conn, &wg) // Handle client connection concurrently
	}
	<-sigs
	listener.Close()
	wg.Wait()
	fmt.Println("exit..")
}
