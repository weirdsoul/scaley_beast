package planestate

import (
	"testing"
	"time"
)

func TestPlaneStateBasic(t *testing.T) {
	state := NewPlaneState()

	// Create a control data state entry for speed.
	seqNum := state.UpdateControlData(DataSet{
		Index:  Speed,
		Values: [8]float32{120, 0, 0, 0, 0, 0, 0, 0},
	})
	state.UpdateControlData(DataSet{
		Index:  RPM,
		Values: [8]float32{2400, 0, 0, 0, 0, 0, 0, 0},
	})
	// We claim to have retrieved speed already.
	_, dataSet := state.GetControlDataSince(seqNum, false)
	if len(dataSet) != 1 {
		t.Errorf("GetControlDataSince: len(dataSet)=%v, expected: 1",
			len(dataSet))
	}
	if dataSet[0].Index != RPM {
		t.Errorf("dataSet[0].Index: %v", dataSet[0].Index)
	}
	// Now we add a new value for speed.
	state.UpdateControlData(DataSet{
		Index:  Speed,
		Values: [8]float32{125, 0, 0, 0, 0, 0, 0, 0},
	})
	// Querying with the same sequence number should now return two updates.
	seqNum, dataSet = state.GetControlDataSince(seqNum, false)
	if len(dataSet) != 2 {
		t.Errorf("GetControlDataSince: len(dataSet)=%v, expected: 2",
			len(dataSet))
	}
	if dataSet[0].Index != Speed {
		t.Errorf("dataSet[0].Index: %v", dataSet[0].Index)
	}
	// Also, the speed should have been updated.
	if dataSet[0].Values[0] != 125 {
		t.Errorf("dataSet[0].Values[0]: %v, expected: 125",
			dataSet[0].Values[0])
	}

	// Finally, querying with the update sequence number should show
	// nothing new, and the same seqNum should be returned.
	newSeq, dataSet := state.GetControlDataSince(seqNum, false)
	if len(dataSet) > 0 {
		t.Errorf("GetControlDataSince: len(dataSet)=%v, expected: 0",
			len(dataSet))
	}
	if newSeq != seqNum {
		t.Errorf("newSeq(%v) == seqNum(%v)", newSeq, seqNum)
	}
}

func TestPlaneStateBlocking(t *testing.T) {
	state := NewPlaneState()
	go func() {
		// Sleep for a bit before we update the data.
		time.Sleep(time.Duration(100) * time.Millisecond)
		state.UpdateControlData(DataSet{
			Index:  RPM,
			Values: [8]float32{2400, 0, 0, 0, 0, 0, 0, 0},
		})
	}()
	// We want to block until we receive data.
	_, dataSet := state.GetControlDataSince(0, true)
	if len(dataSet) != 1 {
		t.Errorf("GetControlDataSince: len(dataSet)=%v, expected: 1",
			len(dataSet))
	}
	if dataSet[0].Index != RPM {
		t.Errorf("dataSet[0].Index: %v", dataSet[0].Index)
	}
}
