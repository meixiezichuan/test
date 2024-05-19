package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var interal = 10 * time.Millisecond // ms
var duration = 5 * time.Minute      //minitue
func handleClient(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	// Read data from the client
	count := duration.Milliseconds() / interal.Milliseconds()
	reader := bufio.NewReader(conn)
	var i int64
	for i = 0; i < count; i++ {
		time.Sleep(interal)
		startTime := time.Now().UnixNano()
		mTime := strconv.FormatInt(startTime, 10) + "\n"
		_, err := conn.Write([]byte(mTime))
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		now := time.Now().UnixNano()
		mtime := data[:len(data)-1]
		startTime, err = strconv.ParseInt(mtime, 10, 64)
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			return
		}
		//fmt.Println("clientTime:", clientTime, "now", now)
		dur := (now - startTime) / 1000 // RTT
		// microsecond
		fmt.Println("latency:", dur/2)
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

	//fmt.Println("Server listening on port 6666")
	var wg sync.WaitGroup
	//count := 0
	//tc := 10000
	//total := os.Getenv("TOTAL")
	//tc, err = strconv.Atoi(total)
	//if err != nil {
	//	fmt.Println("Error get TOTAL:", err)
	//}
	// Accept and handle incoming connections
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			//fmt.Println("Accepted connection from:", conn.RemoteAddr())
			wg.Add(1)
			//count++
			//if count == tc {
			//	flag = true
			//	fmt.Println("Accpet all connection:", tc)
			//}
			go handleClient(conn, &wg) // Handle client connection concurrently
		}
	}()
	<-sigs
	listener.Close()
	wg.Wait()
	fmt.Println("exit..")
}
