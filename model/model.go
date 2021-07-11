package model

import (
	"sync"
	"time"
)

var singletonOnce sync.Once

type RegionStates []RegionState

type RegionStatesMap map[string]*RegionState

type RegionState struct {
	Domain    string      `json:"domain"`
	Status    int         `json:"status"`
	Timestamp time.Time   `json:"timestamp"`
	PrevState []PrevState `json:"prevState"`
}

type PrevState struct {
	Status    int       `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

var RegionStatesMapVar RegionStatesMap

func NewSingletonRegionStates() RegionStatesMap {
	singletonOnce.Do(func() {
		RegionStatesMapVar = make(RegionStatesMap)
	})
	return RegionStatesMapVar
}
