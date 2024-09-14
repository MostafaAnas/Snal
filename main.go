package main

import (
	"fmt"
	"os"

	probing "github.com/prometheus-community/pro-bing"
)

func ping(host string) {
	
	pinger, err := probing.NewPinger(host)
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(true) 
	// TODO: paramartize count
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Print(stats)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: snal <host>")
		return
	}

	host := os.Args[1]
	ping(host)
}
