package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"

	"github.com/weirdsoul/scaley_beast/scalestate"
)

const receiverBufferSize = 2048

// ReadUDPLooping blocks forever and keeps reading from the specified UDP port.
// It uses planeState to maintain a consistant view of the world.
func ReadUDPLooping(sock *net.UDPConn, planeState *planestate.PlaneState) {
	buffer := make([]byte, receiverBufferSize)
	// Main receiver loop. It never stops.
	for {
		n, _, err := sock.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("sock.ReadFromUDP: %v", err)
		}
		if (n-5)%36 == 0 && string(buffer[:4]) == "DATA" {
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
				// This looks like a valid X-Plane data packet. Update the control state.
				planeState.UpdateControlData(dataSet)
			}
		} else {
			log.Printf("Ignoring packet with header '%s'", string(buffer[:4]))
		}
	}
}
