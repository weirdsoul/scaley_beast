package main

import (
	"sync"
	"time"
)

// A single piece of control data along with a timestamp.
type stampedControlData struct {
	// The time at which this piece of data was produced.
	TimeStamp time.Time
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
	// Contains the timestamp of the last update. Clients first check this
	// timestamp to see if anything has changed before going through the map.
	lastUpdate time.Time
	// The mutex guarding this PlaneState instance.
	m sync.RWMutex
}

// NewPlaneState returns a new plane state instance with no current control data.
func NewPlaneState() *PlaneState {
	return &PlaneState{
		lastUpdate:   time.Now(),
		dataStateMap: make(map[ControlDataType]stampedControlData),
	}
}

// UpdateControlData updates the plane state with the specified piece of control
// data. Also updates the lastUpdate timestamp.
func (p *PlaneState) UpdateControlData(controlData DataSet) {
	t := time.Now()
	newControlData := stampedControlData{
		TimeStamp:   t,
		ControlData: controlData,
	}

	p.m.Lock()
	defer p.m.Unlock()

	p.dataStateMap[controlData.Index] = newControlData
	p.lastUpdate = t
}

// GetControlDataSince retrieves all control data updates since the specified
// timestamp. Returns a slice of DataSet instances, one for each type of control
// data that was updated. Returns nil if no updates occurred.
func (p PlaneState) GetControlDataSince(t time.Time) []DataSet {
	p.m.RLock()
	defer p.m.RUnlock()

	if !p.lastUpdate.After(t) {
		return nil
	}

	var newData []DataSet
	for _, controlData := range p.dataStateMap {
		// Pull all control data that has been updated since t.
		if controlData.TimeStamp.After(t) {
			newData = append(newData, controlData.ControlData)
		}
	}
	return newData
}
