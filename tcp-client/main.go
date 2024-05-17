package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func randomInRange(min, max int) int {
    return rand.Intn(max-min+1) + min
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port message\n", os.Args[0])
		os.Exit(1)
	}

	serverAddr := os.Args[1]
	rand.Seed(time.Now().UnixNano())
	min := 40000
        max := 48000
        sourcePort := randomInRange(min, max)
	localAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", sourcePort))
	if err != nil {
		return
	}

	// Create a dialer with the local address
	dialer := &net.Dialer{
		LocalAddr: localAddr,
	}
	// Connect to TCP server
	conn, err := dialer.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	message := fmt.Sprintf("hello from %s\n", localAddr.String())
	// Send data to server
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	fmt.Println("Sent:", message)
        startTime := time.Now()
        data := make([]byte, 1024) // 1KB message
        for i := 0; i < 1023; i++ {
            data[i] = 'a'
        }

        // 最后一个元素赋值为 '\n'
        data[1023] = '\n'
        //for i := 0; i < 10000; i++ { // Send 10MB of data
        //    _, err = conn.Write(data)
        //    if err != nil {
        //        log.Fatalf("Error sending data: %v", err)
        //    }
        //}

	// Read response from server
	reader := bufio.NewReader(conn)
	for {
	     _, err = conn.Write([]byte("hello\n"))
            if err != nil {
                    fmt.Println("Error sending data:", err)
                    return
            }
	    fmt.Println("Sent: hello")

            // Set a deadline for reading from the connection
            conn.SetReadDeadline(time.Now().Add(60 * time.Second)) // 5 seconds timeout

            // Wait and read the response
            response, err := reader.ReadString('\n')
            if err != nil {
                    fmt.Println("Error read data:", err)
            }
	    fmt.Println("Received:", response)
	}
        duration := time.Since(startTime)
        fmt.Println("Sent 10MB of data in %v", duration)
}

