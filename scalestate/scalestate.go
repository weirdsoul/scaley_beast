// Package planestate contains code for managing data state.
package planestate

import (
	"sync"
)

// The type of control data contained in a data set.
type ControlDataType int32

const (
        Weight ControlDataType = 0
)

// DataSet stores a single set of X-Plane control data.
type DataSet struct {
	// The index number of the data point.
	Index ControlDataType
	// The actual data values. Index determines their interpretation.
	// Unused values contain either zero or -999.
	Values [8]float32
}

// A sequence number is a monotonically increasing value that is updated with
// every state update. It is used to track which updates have been received.
type SequenceNumber int64

// A single piece of control data along with a timestamp.
type stampedControlData struct {
	// The sequence number of the last update for this data set.
	SeqNum SequenceNumber
	// The actual control data.
	ControlData DataSet
}

// PlaneState stores the current state of the plane as well as a timestamp for each
// control data value. This allows clients to quickly poll information that has been
// updated since the last query. Newer values will always replace older ones, so
// clients that can't process data quickly enough will only consider the most recent
// value for a specific parameter.
type PlaneState struct {
	// This map contains a timestamped entry for each control data type
	// in use, representing the most recent state.
	dataStateMap map[ControlDataType]stampedControlData
	// Contains sequence number of the last update. Clients first check this
	// value to see if anything has changed before going through the map.
	lastUpdate SequenceNumber
	// The mutex guarding this PlaneState instance.
	m *sync.RWMutex
	// The condition that is triggered when lastUpdate is updated.
	cond *sync.Cond
}

// NewPlaneState returns a new plane state instance with no current control data.
func NewPlaneState() *PlaneState {
	mutex := &sync.RWMutex{}
	return &PlaneState{
		lastUpdate:   0,
		dataStateMap: make(map[ControlDataType]stampedControlData),
		m:            mutex,
		cond:         sync.NewCond(mutex.RLocker()),
	}
}

// UpdateControlData updates the plane state with the specified piece of control
// data. Also increases lastUpdate. Returns the sequence number at which the
// update was performed.
func (p *PlaneState) UpdateControlData(controlData DataSet) SequenceNumber {
	p.m.Lock()
	defer p.m.Unlock()

	p.lastUpdate++
	newControlData := stampedControlData{
		SeqNum:      p.lastUpdate,
		ControlData: controlData,
	}

	p.dataStateMap[controlData.Index] = newControlData
	// Let the world know we have an update to make.
	p.cond.Broadcast()
	return p.lastUpdate
}

// GetControlDataSince retrieves all control data updates since the specified
// SequenceNumber. Returns the sequence number of the last update and a slice of
// DataSet instances, one for each type of control data that was updated.
// If blocking is true, the call will block until updates are available.
func (p *PlaneState) GetControlDataSince(seqNum SequenceNumber, blocking bool) (SequenceNumber, []DataSet) {
	p.m.RLock()
	defer p.m.RUnlock()

	for p.lastUpdate <= seqNum {
		if blocking {
			// Wait for updates.
			p.cond.Wait()
		} else {
			// Just return, we do not wait for any updates to occur.
			return p.lastUpdate, nil
		}
	}

	var newData []DataSet
	for _, controlData := range p.dataStateMap {
		// Pull all control data that has been updated since t.
		if controlData.SeqNum > seqNum {
			newData = append(newData, controlData.ControlData)
		}
	}
	return p.lastUpdate, newData
}
