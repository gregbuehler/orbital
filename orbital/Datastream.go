package orbital

import (
	"fmt"
	"net/http"
)

// Datastream object
type Datastream struct {
	StreamID     string      `json:"Stream ID,required" description:"The ID of the Datastream"`
	Updated      string      `json:"Updated" description:"The time at which this Datastream was last updated"`
	Tags         []string    `json:"Tags" description:"Tagged metadata about the Datastream"`
	Units        string      `json:"Units" description:"The units of the Datastream, for example 'Celsius'"`
	UnitType     string      `json:"Unit type" description:"The type of the unit, for example 'basicSI'"`
	UnitSymbol   string      `json:"Unit symbol" description:"The symbol of the unit, for example 'C'"`
	MinValue     string      `json:"Min Value" description:"The minimum value since the last reset."`
	MaxValue     string      `json:"Max Value" description:"The maximum value since the last reset."`
	CurrentValue string      `json:"Current Value,required" description:"The current value of the Datastream"`
	Datapoints   []Datapoint `json:"Datapoints" description:"A collection of time-stamped values."`
}

// DatastreamListHandler lists datastreams belonging to a feed
func DatastreamListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DatastreamSimpleListHandler!\n\n%s", r.RequestURI)
}

// DatastreamCreateHandler creates a new datastream in a feed
func DatastreamCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DatastreamCreateHandler!\n\n%s", r.RequestURI)
}

// DatastreamDetailHandler details datastream contents
func DatastreamDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DatastreamDetailHandler!\n\n%s", r.RequestURI)
}

// DatastreamUpdateHandler updates datastream details
func DatastreamUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DatastreamUpdateHandler!\n\n%s", r.RequestURI)
}

// DatastreamDeleteHandler deletes a datastream
func DatastreamDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DatastreamDeleteHandler!\n\n%s", r.RequestURI)
}
