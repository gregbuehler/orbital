package orbital

import (
	"fmt"
	"net/http"
)

// Device object
type Device struct {
	Serial         string `json:"serial,required" description:"The serial number of the device"`
	ActivatedAt    string `json:"activated_at,required" description:"The date at which the device was activated in ISO8601 format"`
	CreatedAt      string `json:"created_at,required" description:"The date at which the device was created in ISO8601 format"`
	FeedID         string `json:"Feed_id,required" description:"The id for the Feed that was created for this device"`
	APIKey         string `json:"api_key,required" description:"The API key for the Feed that was created for this device."`
	ActivationCode string `json:"activation_code,required" description:"Activation code for registration"`
}

// DeviceListHandler lists devices belonging to a product
func DeviceListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceSimpleListHandler!\n\n%s", r.RequestURI)
}

// DeviceCreateHandler creates a device belonging to a product
func DeviceCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceCreateHandler!\n\n%s", r.RequestURI)
}

// DeviceDetailHandler details a device belonging to a product
func DeviceDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceDetailHandler!\n\n%s", r.RequestURI)
}

// DeviceUpdateHandler updates a device belonging to a product
func DeviceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceUpdateHandler!\n\n%s", r.RequestURI)
}

// DeviceDeleteHandler deletes a device belonging to a  product
func DeviceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceDeleteHandler!\n\n%s", r.RequestURI)
}

// DeviceActivateHandler activates a device
func DeviceActivateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceActivateHandler!\n\n%s", r.RequestURI)
}
