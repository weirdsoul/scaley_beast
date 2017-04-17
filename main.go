// The instruments server binary contains the server side code to receive
// flight instrument data from UDP and export it via JSON.
// It also serves the main html file and all Javascript.
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/weirdsoul/browser_instruments/planestate"
	"github.com/weirdsoul/browser_instruments/webservice"
)

var udpPort = flag.Int("udp_port", 49042, "udp port to listen on")
var httpPort = flag.Int("http_port", 8080, "port used for serving http")
var staticDir = flag.String("static_dir", "./static_html", "directory containing static content")

func main() {
	fmt.Println("Instruments server - Version 0.1")
	fmt.Println("Copyright 2017 Andreas Eckleder\n")

	flag.Parse()

	log.Printf("Listening on port %v.\n", *udpPort)

	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%v", *udpPort))
	if err != nil {
		log.Fatalf("net.ResolveUDPAddr: %v", err)
	}
	sock, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("net.ListenUDP: %v", err)
	}
	planeState := planestate.NewPlaneState()

	// Start receiver loop in the background. It never stops.
	go ReadUDPLooping(sock, planeState)

	log.Printf("Serving http on port %v.\n", *httpPort)

	// Start serving http.
	webservice.ServeHTTP(*httpPort, *staticDir, planeState)
}
