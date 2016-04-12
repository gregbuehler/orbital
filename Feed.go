package main

import (
	"fmt"
	"net/http"
)

type Feed struct {
	ID          string       `json:"ID" description:"The ID of the Feed"`
	Title       string       `json:"Title,required" description:"A descriptive name for the Feed"`
	Private     string       `json:"Private" description:"Whether or not the Feed is private to the creator of the Feed"`
	Tags        []string     `json:"Tags" description:"Tagged metadata about the Feed"`
	Description string       `json:"Description" description:"A description of the Feed."`
	Feed        string       `json:"Feed" description:"The URL of the Feed"`
	Status      string       `json:"Status" description:"Feed 'live' or 'frozen'"`
	Updated     string       `json:"Updated" description:"The time at which this Feed was last updated"`
	Created     string       `json:"Created" description:"The date the Feed was created"`
	Creator     string       `json:"Creator" description:"A URL referencing the creator of the Feed"`
	Version     string       `json:"Version" description:"Version of the data format Feed returned"`
	Location    string       `json:"Location" description:"Location metadata"`
	Website     string       `json:"Website" description:"The URL of a website which is relevant to this Feed"`
	Icon        string       `json:"Icon" description:"The URL of an icon which is relevant to this Feed"`
	UserLogin   string       `json:"User Login" description:"The user who owns the Feed"`
	Datastreams []Datastream `json:"Datastreams" description:"A collection of the Datastreams in this Feed"`
}

func FeedSimpleListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "FeedSimpleListHandler!\n\n%s", r.RequestURI)
}

func FeedVerboseListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "FeedVerboseListHandler!\n\n%s", r.RequestURI)
}

func FeedDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "FeedDetailHandler!\n\n%s", r.RequestURI)
}

func FeedUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "FeedUpdateHandler!\n\n%s", r.RequestURI)
}
