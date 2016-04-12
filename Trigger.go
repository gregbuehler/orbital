package main

import (
	"fmt"
	"net/http"
)

var TriggerTypes = [...]string{
	"gt",     // greater than
	"gte",    // greater than, or equal to
	"lt",     // less than
	"lte",    // less than, or equal to
	"eq",     // equal to
	"change", // any change
	"frozen", // no updates for 15m
	"live",   // upon update if frozen
}

type Trigger struct {
	Url            string `json:"Url,required" description:"the URL to be POSTed to when the trigger threshold is met"`
	TriggerType    string `json:"Trigger_type,required" description:"the type of evaluation to be done by the trigger"`
	ThresholdValue string `json:"Threshold_value" description:"the threshold value to be evaluated against"`
	FeedId         string `json:"Environment_id,required" description:"Feed ID for the Feed that contains the Datastream to be monitored"`
	StreamId       string `json:"Stream_id,required" description:"string name of the Datastream to be monitored"`
}

func TriggerListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TriggerSimpleListHandler!\n\n%s", r.RequestURI)
}

func TriggerCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TriggerCreateHandler!\n\n%s", r.RequestURI)
}

func TriggerDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TriggerDetailHandler!\n\n%s", r.RequestURI)
}

func TriggerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TriggerUpdateHandler!\n\n%s", r.RequestURI)
}

func TriggerDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TriggerDeleteHandler!\n\n%s", r.RequestURI)
}
