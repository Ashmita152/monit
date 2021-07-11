package server

import (
	"encoding/json"
	"net/http"

	"github.com/Ashmita152/testInternBoilerPlate/checker"
	"github.com/Ashmita152/testInternBoilerPlate/model"
)

type StatusHandler struct{}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{}
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		checker.Lock.Lock()

		regionStates := new(model.RegionStates)

		if model.RegionStatesMapVar != nil {
			for _, v := range model.RegionStatesMapVar {
				*regionStates = append(*regionStates, *v)
			}
		} else {
			w.Write([]byte("Error"))
		}

		data, err := json.Marshal(regionStates)
		if err != nil {
			w.Write([]byte("Error while marshalling data"))
		} else {
			w.Write(data)
		}

		checker.Lock.Unlock()
	}
}
