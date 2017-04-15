// The instruments server binary contains the server side code to receive
// flight instrument data from UDP and export it via JSON.
// It also serves the main html file and all Javascript.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var port = flag.Int("port", 49042, "udp port to listen on")

func main() {
	fmt.Println("Instruments server - Version 0.1")
	fmt.Println("Copyright 2017 Andreas Eckleder\n")

	flag.Parse()

	log.Printf("Listening on port %v.\n", *port)

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatalf("net.ResolveUDPAddr: %v", err)
	}
	sock, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("net.ListenUDP: %v", err)
	}
	// Start receiver loop in the background. It never stops.
	go ReadUDPLooping(sock)

	for {
	}
}
