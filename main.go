package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func ping(host string){
	start := time.Now()
	conn , err := net.Dial("ip4:icmp", host)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	defer conn.Close()

	// Send ping request
	_ , err = conn.Write([]byte("ping"))
	if err != nil {
		fmt.Printf("Error sending ping: %s\n", err)
		return
	}

	// Wait for response
	_, err = conn.Read(make([]byte, 512))
	if err != nil {
		fmt.Printf("Error reading response: %s\n", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("Ping to %s took %s", host , elapsed)
}

func main(){
	if len(os.Args) < 2{
	    fmt.Println("Usage: snal <host>")
		return
	}

	host := os.Args[1]
	ping(host)
}