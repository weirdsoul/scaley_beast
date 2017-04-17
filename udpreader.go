package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"

	"github.com/weirdsoul/browser_instruments/planestate"
)

const receiverBufferSize = 2048

// ReadUDPLooping blocks forever and keeps reading from the specified UDP port.
// It maintains a consistant view of the world which can then be polled for
// the next messages to be sent via JSON.
func ReadUDPLooping(sock *net.UDPConn) {
	buffer := make([]byte, receiverBufferSize)
	// Main receiver loop. It never stops.
	for {
		n, senderAddr, err := sock.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("sock.ReadFromUDP: %v", err)
		}
		log.Printf("Received %v bytes of data from %v", n, *senderAddr)
		if (n-5)%36 == 0 && string(buffer[:4]) == "DATA" {
			log.Printf("This looks like an X-Plane DATA packet.")
			r := bytes.NewReader(buffer[5:n])
			for {
				var dataSet planestate.DataSet
				if err = binary.Read(r, binary.LittleEndian, &dataSet); err != nil {
					if err != io.EOF {
						// Only print unexpected errors.
						log.Printf("binary.Read: %v", err)
					}
					break
				}
				log.Printf("Message with index %v: %v", dataSet.Index, dataSet.Values)
			}
		} else {
			log.Printf("Ignoring packet with header '%s'", string(buffer[:4]))
		}
	}
}
