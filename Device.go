package main

import (
	"fmt"
	"net/http"
)

type Device struct {
	Serial         string `json:"serial,required" description:"The serial number of the device"`
	ActivatedAt    string `json:"activated_at,required" description:"The date at which the device was activated in ISO8601 format"`
	CreatedAt      string `json:"created_at,required" description:"The date at which the device was created in ISO8601 format"`
	FeedId         string `json:"Feed_id,required" description:"The id for the Feed that was created for this device"`
	ApiKey         string `json:"api_key,required" description:"The API key for the Feed that was created for this device."`
	ActivationCode string `json:"activation_code,required" description:"Activation code for registration"`
}

func DeviceListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceSimpleListHandler!\n\n%s", r.RequestURI)
}

func DeviceCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceCreateHandler!\n\n%s", r.RequestURI)
}

func DeviceDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceDetailHandler!\n\n%s", r.RequestURI)
}

func DeviceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceUpdateHandler!\n\n%s", r.RequestURI)
}

func DeviceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceDeleteHandler!\n\n%s", r.RequestURI)
}

func DeviceActivateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeviceActivateHandler!\n\n%s", r.RequestURI)
}
