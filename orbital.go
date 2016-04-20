package main

import (
	"fmt"
	"github.com/gregbuehler/orbital/orbital"
	"log"
	"net/http"

	"database/sql"
	"github.com/gregbuehler/orbital/orbital/database"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

var (
	Version   = "0.0.0"
	BuildTime = "???"
	Branch    = "???"
)

var (
	databaseAdapter          = "postgres"
	databaseConnectionString = "postgres://orbital:orbital@localhost/orbital?sslmode=disable"
)

func main() {
	fmt.Printf("orbital %s %s %s\n", Version, BuildTime, Branch)

	var err error
	database.DBCon, err = sql.Open(databaseAdapter, databaseConnectionString)
	if err != nil {
		panic("Database could not be initialized: " + err.Error())
	}

	if err = database.DBCon.Ping(); err != nil {
		panic("Database connection could not be established: " + err.Error())
	}

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
	router.HandleFunc("/v2/feeds", orbital.FeedSimpleListHandler).Methods("GET")
	router.HandleFunc("/v2/feeds/", orbital.FeedVerboseListHandler).Methods("GET")
	router.HandleFunc("/v2/feeds/{feed_id}", orbital.FeedDetailHandler).Methods("GET")
	router.HandleFunc("/v2/feeds/{feed_id}", orbital.FeedUpdateHandler).Methods("PUT")

	// Datastreams
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams", orbital.DatastreamListHandler).Methods("GET")
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams", orbital.DatastreamCreateHandler).Methods("POST")
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}", orbital.DatastreamDetailHandler).Methods("GET")
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}", orbital.DatastreamUpdateHandler).Methods("PUT")
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}", orbital.DatastreamDeleteHandler).Methods("DELETE")

	// Datapoints
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}/datapoints", orbital.DatapointDeleteHandler).Methods("DELETE")
	router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}/datapoints/{timestamp}", orbital.DatapointDeleteByTimestampHandler).Methods("DELETE")

	// Products
	router.HandleFunc("/v2/products", orbital.ProductListHandler).Methods("GET")
	router.HandleFunc("/v2/products", orbital.ProductCreateHandler).Methods("POST")
	router.HandleFunc("/v2/products/{product_id}", orbital.ProductDetailHandler).Methods("GET")
	router.HandleFunc("/v2/products/{product_id}", orbital.ProductUpdateHandler).Methods("PUT")
	router.HandleFunc("/v2/products/{product_id}", orbital.ProductDeleteHandler).Methods("DELETE")

	// Devices
	router.HandleFunc("/v2/products/{product_id}/devices", orbital.DeviceListHandler).Methods("GET")
	router.HandleFunc("/v2/products/{product_id}/devices", orbital.DeviceCreateHandler).Methods("POST")
	router.HandleFunc("/v2/products/{product_id}/devices/{serial_number}", orbital.DeviceDetailHandler).Methods("GET")
	router.HandleFunc("/v2/products/{product_id}/devices/{serial_number}", orbital.DeviceUpdateHandler).Methods("PUT")
	router.HandleFunc("/v2/products/{product_id}/devices/{serial_number}", orbital.DeviceDeleteHandler).Methods("DELETE")
	router.HandleFunc("/v2/devices/{activation_code}/activate", orbital.DeviceActivateHandler).Methods("GET")

	// Keys
	router.HandleFunc("/v2/keys", orbital.KeyListHandler).Methods("GET")
	router.HandleFunc("/v2/keys", orbital.KeyCreateHandler).Methods("POST")
	router.HandleFunc("/v2/keys/{key_id}", orbital.KeyDetailHandler).Methods("GET")
	router.HandleFunc("/v2/keys/{key_id}", orbital.KeyDeleteHandler).Methods("DELETE")

	// Triggers
	router.HandleFunc("/v2/triggers", orbital.TriggerListHandler).Methods("GET")
	router.HandleFunc("/v2/triggers", orbital.TriggerCreateHandler).Methods("POST")
	router.HandleFunc("/v2/triggers/{trigger_id}", orbital.TriggerDetailHandler).Methods("GET")
	router.HandleFunc("/v2/triggers/{trigger_id}", orbital.TriggerUpdateHandler).Methods("PUT")
	router.HandleFunc("/v2/triggers/{trigger_id}", orbital.TriggerDeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func NotImplementedYet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NotImplementedYet!\n\n%s", r.RequestURI)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index!")
}
