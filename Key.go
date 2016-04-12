package main

import (
	"fmt"
	"net/http"
)

type Key struct {
	ApiKey        string `json:"api_key,required" description:"The actual api key token to be used by the client application"`
	Label         string `json:"label,required" description:"A label by which the key can be referenced	"`
	ExpiresAt     string `json:"expires_at" description:"Expiration date for the key after which it won’t work"`
	PrivateAccess string `json:"private_access" description:"Flag that indicates whether this key can access private resources belonging to the user"`
	Permissions   string `json:"permissions,required" description:"Collection of Permission objects controlling the access level"`
}

type KeyPermissions struct {
	AccessMethods   []string      `json:"access_methods,required" description:"An array indicating what type of access the key has"`
	SourceIP        string        `json:"source_ip" description:"Restricts access to requests that originate from this IP address"`
	Referer         string        `json:"referer" description:"The referer domain"`
	MinimumInterval string        `json:"minimum_interval" description:"Creates a key that can only request data with a certain resolution"`
	Label           string        `json:"label" description:"Optional label for identifying the permission set"`
	Resources       []KeyResource `json:"resources" description:"Resource objects restrict access to specific Feeds or Datastreams"`
}

type KeyResource struct {
	FeedId       string `json:"Feed_id" description:"Reference to a specific Feed id"`
	DatastreamId string `json:"" description:"Reference to a specific Datastream id within a Feed"`
}

func KeyListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "KeySimpleListHandler!\n\n%s", r.RequestURI)
}

func KeyCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "KeyCreateHandler!\n\n%s", r.RequestURI)
}

func KeyDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "KeyDetailHandler!\n\n%s", r.RequestURI)
}

func KeyDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "KeyDeleteHandler!\n\n%s", r.RequestURI)
}
