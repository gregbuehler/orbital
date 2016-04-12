package main

import (
	"fmt"
	"net/http"
)

// A Datapoint object
type Datapoint struct {
	Value     string `json:"value,required" description:"datapoint value"`
	Timestamp string `json:"at,required" description:"datapoint timestamp"`
}

func DatapointDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DatapointSimpleDeleteHandler!\n\n%s", r.RequestURI)
}

func DatapointDeleteByTimestampHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DatastreamSimpleDeleteByTimestampHandler!\n\n%s", r.RequestURI)
}
