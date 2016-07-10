package apistructs

import (
	"encoding/json"
	"strings"
)

type APIPatroniState struct {
	DatabaseSystemIdentifier string `json:"database_system_identifier"`
	PostmasterStartTime      string `json:"postmaster_start_time"`
	Xlog                     struct {
		ReceivedLocation  int64       `json:"received_location"`
		ReplayedTimestamp interface{} `json:"replayed_timestamp"`
		Paused            bool        `json:"paused"`
		ReplayedLocation  int64       `json:"replayed_location"`
	} `json:"xlog"`
	Patroni struct {
		Scope   string `json:"scope"`
		Version string `json:"version"`
	} `json:"patroni"`
	State         string `json:"state"`
	Role          string `json:"role"`
	ServerVersion int    `json:"server_version"`
}

func NewAPIPatroniState(jsonStream string) (apiPatroniState *APIPatroniState, err error) {
	apiPatroniState = &APIPatroniState{}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	err = dec.Decode(apiPatroniState)

	return
}
