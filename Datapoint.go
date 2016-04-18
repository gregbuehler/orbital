package main

import (
	"fmt"
	"net/http"
)

// Datapoint object
type Datapoint struct {
	Value     string `json:"value,required" description:"datapoint value"`
	Timestamp string `json:"at,required" description:"datapoint timestamp"`
}

// DatapointDeleteHandler deletes all datapoints from a datastream
func DatapointDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DatapointSimpleDeleteHandler!\n\n%s", r.RequestURI)
}

// DatapointDeleteByTimestampHandler deletes a datapoint from a datastream by timestamp
func DatapointDeleteByTimestampHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DatastreamSimpleDeleteByTimestampHandler!\n\n%s", r.RequestURI)
}
