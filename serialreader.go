package main

import (
        "bufio"
	"io"
	"log"
        "strconv"

	"github.com/weirdsoul/scaley_beast/scalestate"
)

// ReadSerialLooping blocks forever and keeps reading from the specified reader.
// It uses scaleState to maintain a consistant view of the world.
func ReadSerialLooping(reader io.Reader, planeState *planestate.PlaneState) {
        bufReader := bufio.NewReader(reader)
	// Main receiver loop. It never stops.
	for {
		l, err := bufReader.ReadString('\n')
		if err != nil {
			log.Printf("bufReader.ReadString: %v", err)
		}
 		f, err := strconv.ParseFloat(l[0:len(l)-1], 32)
		if err != nil {
		        log.Printf("strconv.ParseFloat: %v", err)
		}
		log.Printf("weight=%v", f)

		var dataSet planestate.DataSet
		dataSet.Index = planestate.Weight
		
		dataSet.Values[0] = float32(f)
		planeState.UpdateControlData(dataSet)
	}
}
