package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
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
	//rand.Seed(time.Now().UnixNano())
	//min := 1024
	//max := 48000
	//sourcePort := randomInRange(min, max)
	//localAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", sourcePort))
	//if err != nil {
	//	return
	//}
	//
	//// Create a dialer with the local address
	//dialer := &net.Dialer{
	//	LocalAddr: localAddr,
	//}
	//// Connect to TCP server
	//conn, err := dialer.Dial("tcp", serverAddr)
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Read response from server
	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}
	}

}
