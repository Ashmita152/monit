package checker

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/Ashmita152/testInternBoilerPlate/model"
)

var Lock = &sync.Mutex{}

func PollEndpoints(endpoints []string) {
	var states = model.NewSingletonRegionStates()

	for {
		for _, endpoint := range endpoints {
			go PollEndpoint(endpoint, states)
		}
		time.Sleep(time.Second * 60)
	}
}

func PollEndpoint(endpoint string, states model.RegionStatesMap) {
	Lock.Lock()

	response, err := http.Get("https://" + endpoint)
	if err != nil {
		log.Printf("Error while polling %s: %s\n", endpoint, err)
	}

	// if part handles the first insertion
	// else part handles all the other insertions
	if _, ok := states[endpoint]; !ok {
		states[endpoint] = &model.RegionState{
			Domain:    endpoint,
			Status:    response.StatusCode,
			Timestamp: time.Now().UTC(),
		}
	} else {
		prevState := model.PrevState{
			Status:    states[endpoint].Status,
			Timestamp: states[endpoint].Timestamp,
		}
		states[endpoint].PrevState = append(states[endpoint].PrevState, prevState)
		states[endpoint].Status = response.StatusCode
		states[endpoint].Timestamp = time.Now().UTC()
	}
	Lock.Unlock()
}
