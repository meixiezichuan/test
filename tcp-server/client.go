package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:20985")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Start a loop to continuously read data from the server
	for {
		// Read data from the server
		buffer := make([]byte, 1024)
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data:", err)
			return
		}

		// Print the received data
		fmt.Println("Received data from server:", string(buffer[:bytesRead]))
	}
}
