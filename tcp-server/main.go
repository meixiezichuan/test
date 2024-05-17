package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	// Read data from the client
        reader := bufio.NewReader(conn)
	for {
	    data, err := reader.ReadString('\n')
	    if err != nil {
	    	fmt.Println("Error reading:", err)
	    	return
	    }
	    fmt.Printf("Received data: %s\n", data)

	    // Send response back to the client
	    conn.Write([]byte("Hello from server\n"))
	}
}

func main() {
	// Start TCP server on port 8080
	port := os.Args[1]
	lport := ":"+ port

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	listener, err := net.Listen("tcp", lport)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on port 6666")

	// Accept and handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Accepted connection from:", conn.RemoteAddr())
		go handleClient(conn) // Handle client connection concurrently
	}
	<-sigs
}

