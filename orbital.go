package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	Version   = "0.0.0"
	BuildTime = "???"
	Branch    = "???"
)

func main() {
	fmt.Printf("orbital %s %s %s\n", Version, BuildTime, Branch)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)

	/*
	 *
	 *    Orbital App API
	 *
	 */

	// Users
	router.HandleFunc("/users/{username}", NotImplementedYet).Methods("GET")

	/*
	 *
	 *    Orbital Data API
	 *
	 */

	// Feeds
	router.HandleFunc("/v2/feeds", FeedSimpleListHandler).Methods("GET")
	router.HandleFunc("/v2/feeds/", FeedVerboseListHandler).Methods("GET")
	router.HandleFunc("/v2/feeds/{feed_id}", FeedDetailHandler).Methods("GET")
	router.HandleFunc("/v2/feeds/{feed_id}", FeedUpdateHandler).Methods("PUT")

	// Datastreams
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams", DatastreamListHandler).Methods("GET")
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams", DatastreamCreateHandler).Methods("POST")
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}", DatastreamDetailHandler).Methods("GET")
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}", DatastreamUpdateHandler).Methods("PUT")
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}", DatastreamDeleteHandler).Methods("DELETE")

	// Datapoints
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}/datapoints", DatapointDeleteHandler).Methods("DELETE")
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}/datapoints/{timestamp}", DatapointDeleteByTimestampHandler).Methods("DELETE")

	// Products
	router.HandleFunc("/v2/products", ProductListHandler).Methods("GET")
	router.HandleFunc("/v2/products", ProductCreateHandler).Methods("POST")
	router.HandleFunc("/v2/products/{product_id}", ProductDetailHandler).Methods("GET")
	router.HandleFunc("/v2/products/{product_id}", ProductUpdateHandler).Methods("PUT")
	router.HandleFunc("/v2/products/{product_id}", ProductDeleteHandler).Methods("DELETE")

	// Devices
	router.HandleFunc("/v2/products/{product_id}/devices", DeviceListHandler).Methods("GET")
	router.HandleFunc("/v2/products/{product_id}/devices", DeviceCreateHandler).Methods("POST")
	router.HandleFunc("/v2/products/{product_id}/devices/{serial_number}", DeviceDetailHandler).Methods("GET")
	router.HandleFunc("/v2/products/{product_id}/devices/{serial_number}", DeviceUpdateHandler).Methods("PUT")
	router.HandleFunc("/v2/products/{product_id}/devices/{serial_number}", DeviceDeleteHandler).Methods("DELETE")
	router.HandleFunc("/v2/devices/{activation_code}/activate", DeviceActivateHandler).Methods("GET")

	// Keys
	router.HandleFunc("/v2/keys", KeyListHandler).Methods("GET")
	router.HandleFunc("/v2/keys", KeyCreateHandler).Methods("POST")
	router.HandleFunc("/v2/keys/{key_id}", KeyDetailHandler).Methods("GET")
	router.HandleFunc("/v2/keys/{key_id}", KeyDeleteHandler).Methods("DELETE")

	// Triggers
	router.HandleFunc("/v2/triggers", TriggerListHandler).Methods("GET")
	router.HandleFunc("/v2/triggers", TriggerCreateHandler).Methods("POST")
	router.HandleFunc("/v2/triggers/{trigger_id}", TriggerDetailHandler).Methods("GET")
	router.HandleFunc("/v2/triggers/{trigger_id}", TriggerUpdateHandler).Methods("PUT")
	router.HandleFunc("/v2/triggers/{trigger_id}", TriggerDeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func NotImplementedYet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NotImplementedYet!\n\n%s", r.RequestURI)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index!")
}
