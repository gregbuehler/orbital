package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/gorilla/mux"
)

var (
    Version = "0.0.0"
    BuildTime = "???"
    Branch = "???"
)

func main() {
  fmt.Println("orbital %s %s %s", Version, BuildTime, Branch)

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)

  // Feeds
  router.HandleFunc("/v2/feeds", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/feeds/", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/feeds/{feed_id}", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/feeds/{feed_id}", NotImplementedYet).Methods("PUT")

  // Datastreams
  router.HandleFunc("/v2/feeds/{feed_id}/datastreams", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/feeds/{feed_id}/datastreams", NotImplementedYet).Methods("POST")
  router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}", NotImplementedYet).Methods("PUT")
  router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}", NotImplementedYet).Methods("DELETE")

  // Datapoints
  router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}/datapoints", NotImplementedYet).Methods("DELETE")
  router.HandleFunc("/v2/feeds/{feed_id}/datastreams/{datastream_id}/datapoints/{timestamp}", NotImplementedYet).Methods("DELETE")

  // Products
  router.HandleFunc("/v2/products", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/products", NotImplementedYet).Methods("POST")
  router.HandleFunc("/v2/products/{product_id}", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/products/{product_id}", NotImplementedYet).Methods("PUT")
  router.HandleFunc("/v2/products/{product_id}", NotImplementedYet).Methods("DELETE")

  // Devices
  router.HandleFunc("/v2/products/{product_id}/devices", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/products/{product_id}/devices", NotImplementedYet).Methods("POST")
  router.HandleFunc("/v2/products/{product_id}/devices/{serial_number}", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/products/{product_id}/devices/{serial_number}", NotImplementedYet).Methods("PUT")
  router.HandleFunc("/v2/products/{product_id}/devices/{serial_number}", NotImplementedYet).Methods("DELETE")
  router.HandleFunc("/v2/devices/{activation_code}/activate", NotImplementedYet).Methods("GET")

  // Keys
  router.HandleFunc("/v2/keys", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/keys", NotImplementedYet).Methods("POST")
  router.HandleFunc("/v2/keys/{key_id}", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/keys/{key_id}", NotImplementedYet).Methods("DELETE")

  // Triggers
  router.HandleFunc("/v2/triggers", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/triggers", NotImplementedYet).Methods("POST")
  router.HandleFunc("/v2/triggers/{trigger_id}", NotImplementedYet).Methods("GET")
  router.HandleFunc("/v2/triggers/{trigger_id}", NotImplementedYet).Methods("PUT")
  router.HandleFunc("/v2/triggers/{trigger_id}", NotImplementedYet).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8080", router))
}

func NotImplementedYet(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "NotImplementedYet!")
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Index!")
}
