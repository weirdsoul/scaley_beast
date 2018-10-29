// The instruments server binary contains the server side code to receive
// flight instrument data from UDP and export it via JSON.
// It also serves the main html file and all Javascript.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tarm/serial"
	"github.com/weirdsoul/scaley_beast/scalestate"
	"github.com/weirdsoul/scaley_beast/webservice"
)

var serialInterface = flag.String("serial", "/dev/cu.wchusbserial620", "serial port to listen on")
var httpPort = flag.Int("http_port", 8080, "port used for serving http")
var staticDir = flag.String("static_dir", "./client", "directory containing static content")

func main() {
	fmt.Println("Scaley beast - Version 0.1")
	fmt.Println("Copyright 2018 Andreas Eckleder\n")

	flag.Parse()

	log.Printf("Listening on serial port %v.\n", *serialInterface)

	planeState := planestate.NewPlaneState()

	c := &serial.Config{Name: *serialInterface,
		Baud: 115200}
	serial, err := serial.OpenPort(c)
	if err != nil {
		log.Printf("serial.OpenPort: %v", err)
		return
	}

	// Start receiver loop in the background. It never stops.
	go ReadSerialLooping(serial, planeState)

	log.Printf("Serving http on port %v.\n", *httpPort)

	// Start serving http.
	webservice.ServeHTTP(*httpPort, *staticDir, planeState)
}
