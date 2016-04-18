package main

import (
	"fmt"
	"net/http"
)

// Key object
type Key struct {
	APIKey        string `json:"api_key,required" description:"The actual api key token to be used by the client application"`
	Label         string `json:"label,required" description:"A label by which the key can be referenced	"`
	ExpiresAt     string `json:"expires_at" description:"Expiration date for the key after which it won’t work"`
	PrivateAccess string `json:"private_access" description:"Flag that indicates whether this key can access private resources belonging to the user"`
	Permissions   string `json:"permissions,required" description:"Collection of Permission objects controlling the access level"`
}

// KeyPermissions object
type KeyPermissions struct {
	AccessMethods   []string      `json:"access_methods,required" description:"An array indicating what type of access the key has"`
	SourceIP        string        `json:"source_ip" description:"Restricts access to requests that originate from this IP address"`
	Referer         string        `json:"referer" description:"The referer domain"`
	MinimumInterval string        `json:"minimum_interval" description:"Creates a key that can only request data with a certain resolution"`
	Label           string        `json:"label" description:"Optional label for identifying the permission set"`
	Resources       []KeyResource `json:"resources" description:"Resource objects restrict access to specific Feeds or Datastreams"`
}

// KeyResource object
type KeyResource struct {
	FeedID       string `json:"Feed_id" description:"Reference to a specific Feed id"`
	DatastreamID string `json:"" description:"Reference to a specific Datastream id within a Feed"`
}

// KeyListHandler lists keys
func KeyListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "KeySimpleListHandler!\n\n%s", r.RequestURI)
}

// KeyCreateHandler creates keys
func KeyCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "KeyCreateHandler!\n\n%s", r.RequestURI)
}

// KeyDetailHandler details a specific key
func KeyDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "KeyDetailHandler!\n\n%s", r.RequestURI)
}

// KeyDeleteHandler deletes a specific key
func KeyDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "KeyDeleteHandler!\n\n%s", r.RequestURI)
}
