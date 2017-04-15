// The instruments server binary contains the server side code to receive
// flight instrument data from UDP and export it via JSON.
// It also serves the main html file and all Javascript.
package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
)

const receiverBufferSize = 2048

// The type of control data contained in a data set.
type ControlDataType int32

const (
	Speed ControlDataType = 3
	RPM   ControlDataType = 37
)

// DataSet stores a single set of X-Plane control data.
type DataSet struct {
	// The index number of the data point.
	Index ControlDataType
	// The actual data values. Index determines their interpretation.
	// Unused values contain either zero or -999.
	Values [8]float32
}

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
				var dataSet DataSet
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
